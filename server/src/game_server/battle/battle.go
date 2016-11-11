package battle

import (
	"sort"
)

// 战场状态
type BattleState struct {
	/*
		runtimeAddMonsterHook 负责
		1. 通知客户端
		2. 往战场中增加怪物
	*/
	runtimeAddMonsterHook func(b *BattleState, result *FightResult) bool

	order             *FightOrder    // 出手顺序
	targets           []*Fighter     // 攻击目标缓存，用于优化查找攻击目标的性能，避免重复创建
	needResort        bool           // 是否需要重新排序
	records           []iRecord      // 战报记录
	result            []*FightResult // 优化性能用
	needRecord        bool           // 需要记录
	round             int            // 回合数
	BatType           int            // 战斗类型 ，BT_开头的常量
	Attackers         *SideInfo      // 攻方数据
	Defenders         *SideInfo      // 守方数据
	enableGhostSkill  bool           // 使用魂侍技能
	useItemRound      bool           // 当前回合使用了战斗道具
	summoner          *Fighter       // 当前回合的召唤者
	launchBattleTotem bool           //检查阵印

	callTotemRound   int   //召唤阵印回合
	currentActionPid int64 //当前可行动玩家
}

type PlayerGhostInfo struct {
	GhostPower     int //魂力
	InitGhostPower int //初始魂力
	EnableGhost    bool
}

type BattlePlayerInfo struct {
	PlayerId  int64           // 玩家ID
	GhostInfo PlayerGhostInfo //魂侍信息
	Auto      bool            //自动战斗

	Speed        int // 主角速度，觉得操作顺序
	PrepareRound int //以准备回合数

	//下面字段废弃
	JobIndex int          // 职业策略索引
	Jobs     [][]*JobInfo // 伙伴职业选择方案
}

// 攻守方信息
type SideInfo struct {
	live       int                 // 存活计数器
	lastLive   int                 // 出手前存活数量，用于伙伴绝招的条件判断
	groupIndex int                 // 当前参与战斗的组索引
	deadNums   int                 //我放死亡角色计数
	TotemInfo  [5]*TotemInfo       //阵印信息
	Groups     [][]*Fighter        // 战斗小组，一组等于一个完整阵型，一组打完再上下一组
	Fighters   []*Fighter          // 角色信息，按角色站位放置数据，没有角色的站位上数据为nil
	Players    []*BattlePlayerInfo // 每个玩家的信息，例如一个玩家的不同角色共享魂力
}

//阵印信息
type TotemInfo struct {
	Id      int16
	SkillId int16
	Level   int16
	used    bool
}

func (this *SideInfo) AnyRoleBadlyHurt() bool {
	for _, fighter := range this.Fighters {
		if fighter != nil {
			if !fighter.IsBattlePet && fighter.Health > MIN_HEALTH && fighter.BadlyHurt() {
				return true
			}
		}
	}
	return false
}

func (this *SideInfo) GetLive() int {
	n := 0

	for _, f := range this.Fighters {
		if f != nil && f.Health > MIN_HEALTH {
			n += 1
		}
	}

	return n
}

//获取阵亡角色数
func (this *SideInfo) GetDead() int {
	return this.deadNums
}

func (this *SideInfo) CurrentGroupIndex() int {
	return this.groupIndex
}

//动态增加一战斗分组
func (b *BattleState) RunTimeAddFightersGroup(fighterType int, fighters []*Fighter) {
	var sideInfo *SideInfo
	if fighterType == FT_ATK {
		sideInfo = b.Attackers
	} else {
		sideInfo = b.Defenders
	}
	for _, fighter := range fighters {
		if fighter != nil {
			fighter.side = sideInfo
			fighter.init(b, fighterType)
		}
	}
	//TODO 这里好像漏了对 BattlePlayerInfo 做初始化 参考  defendSideInit 或者  attackSideInit
	sideInfo.Groups = append(sideInfo.Groups, fighters)
}

func attackSideInit(b *BattleState, attackSideInfo *SideInfo) {
	for _, oneAttacker := range attackSideInfo.Players {
		oneAttacker.GhostInfo.EnableGhost = false
		if oneAttacker.PlayerId > 0 {
			for _, fighter := range attackSideInfo.Fighters {
				if fighter != nil && fighter.PlayerId == oneAttacker.PlayerId {
					fighter.playerInfo = oneAttacker
					fighter.GhostInfo = &(oneAttacker.GhostInfo)
					if len(fighter.Ghosts) > 0 {
						oneAttacker.GhostInfo.EnableGhost = true
					}
					oneAttacker.GhostInfo.InitGhostPower += fighter.InitGhostPower
					if fighter.Kind == FK_PLAYER {
						oneAttacker.Speed = int(fighter.Speed)
					}
				}
			}
		}
		oneAttacker.PrepareRound = b.GetRounds() - 1
		oneAttacker.GhostInfo.GhostPower = oneAttacker.GhostInfo.InitGhostPower
		if oneAttacker.GhostInfo.GhostPower > FULL_GHOST_POWER {
			oneAttacker.GhostInfo.GhostPower = FULL_GHOST_POWER
		}
	}

	for _, attackers := range attackSideInfo.Groups {
		for _, f := range attackers {
			if f == nil {
				continue
			}

			f.side = b.Attackers

			f.init(b, FT_ATK)
		}
	}

	b.UpdateAttackerLiveCount()
}

func defendSideInit(b *BattleState, defendSideInfo *SideInfo) {
	for _, oneDefender := range defendSideInfo.Players {
		oneDefender.GhostInfo.EnableGhost = false
		if oneDefender.PlayerId > 0 {
			for _, fighter := range defendSideInfo.Fighters {
				if fighter != nil && fighter.PlayerId == oneDefender.PlayerId {
					fighter.playerInfo = oneDefender
					fighter.GhostInfo = &(oneDefender.GhostInfo)
					if len(fighter.Ghosts) > 0 {
						oneDefender.GhostInfo.EnableGhost = true
					}
					oneDefender.GhostInfo.InitGhostPower += fighter.InitGhostPower

					if fighter.Kind == FK_PLAYER {
						oneDefender.Speed = int(fighter.Speed)
					}
				}
			}
		}
		oneDefender.PrepareRound = b.GetRounds() - 1
		oneDefender.GhostInfo.GhostPower = oneDefender.GhostInfo.InitGhostPower
		if oneDefender.GhostInfo.GhostPower > FULL_GHOST_POWER {
			oneDefender.GhostInfo.GhostPower = FULL_GHOST_POWER
		}

	}
	for _, defenders := range defendSideInfo.Groups {
		for _, f := range defenders {
			if f == nil {
				continue
			}

			f.side = b.Defenders

			f.init(b, FT_DEF)
		}
	}
	b.UpdateDefenderLiveCount()
}

func orderInit(b *BattleState, resetFighted bool) {
	b.order.fighters = make([]*Fighter, 0, 30)
	b.order.Append(b.Defenders.Fighters, resetFighted)
	b.order.Append(b.Attackers.Fighters, resetFighted)
	sort.Sort(b.order)
}

// 开始一场战斗
func Start(
	batType int, // 战场类型
	attackSideInfo, defendSideInfo *SideInfo, // 攻守方信息
	lastInitFunc func(), // 最后初始化方法
) *BattleState {
	attackSideInfo.Fighters = attackSideInfo.Groups[0]
	defendSideInfo.Fighters = defendSideInfo.Groups[0]

	b := &BattleState{
		BatType:        batType,
		order:          new(FightOrder),
		targets:        make([]*Fighter, 0, 15),
		Attackers:      attackSideInfo,
		Defenders:      defendSideInfo,
		result:         make([]*FightResult, 2),
		callTotemRound: -1,
	}

	attackSideInit(b, attackSideInfo)
	defendSideInit(b, defendSideInfo)

	if lastInitFunc != nil {
		lastInitFunc()
	}

	orderInit(b, true)

	return b
}

// 重新为战场设置攻击方
func (b *BattleState) ResetAttackersSide(attackSide *SideInfo) {
	b.Attackers = attackSide
	attackSideInit(b, attackSide)
	orderInit(b, true)
}

// 在战场动态添加一个战斗对象（召唤）
func (b *BattleState) RuntimeAddFighter(fighterType int, f *Fighter) {
	if fighterType == FT_ATK {
		f.side = b.Attackers
	} else {
		f.side = b.Defenders
	}

	f.side.Fighters[f.Position-1] = f
	f.init(b, fighterType)

	if fighterType == FT_ATK {
		b.UpdateAttackerLiveCount()
	} else {
		b.UpdateDefenderLiveCount()
	}

	orderInit(b, false)
}

// 启用战都过程纪录
func (b *BattleState) EnableRecord() {
	b.needRecord = true
}

func (b *BattleState) UpdateAttackerLiveCount() {
	b.Attackers.live = 0
	for _, f := range b.Attackers.Fighters {
		if f != nil && f.Health > MIN_HEALTH {
			b.Attackers.live++
		}
	}
}

func (b *BattleState) UpdateDefenderLiveCount() {
	b.Defenders.live = 0
	for _, f := range b.Defenders.Fighters {
		if f != nil && f.Health > MIN_HEALTH {
			b.Defenders.live++
		}
	}
}

// 判断一个玩家是否是攻方
/*func (b *BattleState) IsAttacker(playerId int64) bool {
	for _, attackerId := range b.Attackers.playerIds {
		if attackerId == playerId {
			return true
		}
	}

	return false
}*/

func (b *BattleState) GetAttackerPlayerIds() (playerIds []int64) {
	playerIds = make([]int64, 0, len(b.Attackers.Players))

	for _, player := range b.Attackers.Players {
		playerIds = append(playerIds, player.PlayerId)
	}

	return
}

func (b *BattleState) IsDefender(playerId int64) bool {
	for _, player := range b.Defenders.Players {
		if player.PlayerId == playerId {
			return true
		}
	}

	return false
}

// 判断战场是否是由玩家主角先出手
func (b *BattleState) GetFirstFighter() *Fighter {
	var fastest *Fighter

	for _, f := range b.Attackers.Fighters {
		if f != nil && (fastest == nil || fastest.Speed < f.Speed) {
			fastest = f
		}
	}

	for _, f := range b.Defenders.Fighters {
		if f != nil && (fastest == nil || fastest.Speed < f.Speed) {
			fastest = f
		}
	}

	return fastest
}

//获取在 *nowRound* 里面应该出手的一名玩家id
//外部需要维护一个 当前回合数
func (b *BattleState) GetNextPlayer(nowRound int) int64 {
	var (
		speed int
		pid   int64
		///nowRound = b.GetRounds()
	)
	for _, player := range b.Attackers.Players {
		if player.PrepareRound >= nowRound {
			continue
		}
		if pid <= 0 || player.Speed > speed {
			pid = player.PlayerId
			speed = player.Speed
		}
	}
	for _, player := range b.Defenders.Players {
		if player.PrepareRound >= nowRound {
			continue
		}
		if pid <= 0 || player.Speed > speed {
			pid = player.PlayerId
			speed = player.Speed
		}
	}
	return pid
}

func (b *BattleState) SetTotemState() {
	b.launchBattleTotem = true
}

// 设置主角使用技能和是否自动战斗
func (b *BattleState) SetPlayerState(skillIndex int, itemId int32, autoFight bool, playerId int64, isAttacker bool, position int, jobIndex int, useSwordSoul bool, useGhostSkillPosition int, useGhostSkillId int16, launchBattleTotem bool) bool {
	b.launchBattleTotem = launchBattleTotem
	if launchBattleTotem {
		return false
	}

	side := b.Attackers

	if !isAttacker {
		side = b.Defenders
	}

	fighter := side.Fighters[position-1]

	if useGhostSkillPosition > 0 {
		useGhostSkillFighter := side.Fighters[useGhostSkillPosition-1]
		useGhostSkillFighter.useGhostSkillPosition = useGhostSkillPosition
		useGhostSkillFighter.useGhostSkillId = useGhostSkillId
		b.enableGhostSkill = true
	}

	if fighter.Kind == FK_PLAYER && fighter.PlayerId == playerId {
		fighter.useSkillIndex = skillIndex
		fighter.AutoFight = autoFight
		fighter.useItemId = itemId
		if itemId > 0 {
			b.useItemRound = true
		}

		// 三回合后且没使用过
		if useSwordSoul && fighter.SwordSoulValue > 0 {
			fighter.useSwordSoul = true
		} else {
			fighter.useSwordSoul = false
		}

		return true
	}

	return false
}

func (b *BattleState) SetAuto(playerId int64, isAttacker, autoFight bool) bool {
	side := b.Attackers

	if !isAttacker {
		side = b.Defenders
	}

	for _, fighter := range side.Fighters {
		if fighter != nil && fighter.Kind == FK_PLAYER && fighter.PlayerId == playerId {
			fighter.AutoFight = autoFight
			return true
		}
	}

	return false
}

// 计算攻方总的存活角色数量
func (b *BattleState) GetLiveAttackerNum() (result int) {
	for _, attackers := range b.Attackers.Groups {
		for _, attacker := range attackers {
			if attacker != nil && attacker.Health > MIN_HEALTH {
				result += 1
			}
		}
	}
	return
}

// 获取一批战斗数据，不一定是一回合的完整数据
func (b *BattleState) NextRound() (result []*FightResult, status, nowRound int) {
	var (
		order        = b.order
		needRestart  = false
		maxGoToCount = 0
		start        = 0
	)

	result = b.result[0:0]

START:

	maxGoToCount++
	if maxGoToCount > 50 {
		panic("BattleState::NextRound() --> maxGoToCount > 50")
	}

	var (
		i int
		//pri *Fighter
		cur           *Fighter
		useBattleItem bool
		useGhost      bool
		useTotem      bool
	)

	nowRound = b.GetRounds()

	//清理应该消失的灵宠
	if b.summoner == nil {
		for i = 0; i < len(order.fighters); i++ {
			cur = order.fighters[i]
			// 检查是否是灵宠
			if !cur.IsBattlePet || cur.Health <= MIN_HEALTH {
				continue
			}
			/*
				灵宠消失条件：
				灵宠已经超过了出场回合数
			*/
			if nowRound-cur.BattlePetLiveStartRound >= cur.BattlePetLiveRound {
				setDisappearWithBattlePet(cur)
			}
		}
	} else if b.summoner != nil {
		// 召唤者会将一个回合切分成两部分（召唤后下发数据；客户端请求NextRound继续当前回合的计算）
		// 召唤时会记录召唤者，所以在第二部分时从召唤者之后开始计算战报
		for i = 0; i < len(order.fighters); i++ {
			if order.fighters[i] == b.summoner {
				start = i + 1 // +1是从召唤者下一个战斗对象开始计算战报
				break
			}
		}

		b.summoner = nil
	}

	// 所有人按顺序出手，当前后两个出手的角色类型不一样时暂停出手，以保证玩家感受到的是连续攻击，并且有操控主角的机会
	for i = start; i < len(order.fighters); i++ {
		cur = order.fighters[i]
		//伙伴可以在出手后放魂侍，在需要放魂侍的时候已出手的fighter也需要作为释放魂侍的候选走下去检查
		//出手过的玩家不需要再次出手，使用魂侍技能或者释放阵营的情况例外
		if cur.fighted && (!b.enableGhostSkill && !b.launchBattleTotem) {
			continue
		}

		// 到主角等待操作
		if len(result) > 0 && cur.Kind == FK_PLAYER && !cur.AutoFight {
			break
		}

		// 挪到break 之后避免玩家死了操作无效
		if cur.Health <= MIN_HEALTH {
			continue
		}

		b.Attackers.lastLive = b.Attackers.live
		b.Defenders.lastLive = b.Defenders.live

		if b.launchBattleTotem {
			if len(cur.side.TotemInfo) > nowRound && cur.side.TotemInfo[nowRound] != nil && !cur.side.TotemInfo[nowRound].used {
				result = append(result, cur.fight())
				b.launchBattleTotem = false
				cur.side.TotemInfo[nowRound].used = true
				useTotem = true
				break
			}
			continue
		}

		if b.enableGhostSkill {
			if cur.useGhostSkillPosition > 0 && cur.Ftype == FT_ATK {
				fightResult := cur.fight()
				result = append(result, fightResult)
				//连锁魂侍技能释放
				if cur.Health > 0 && b.Defenders.live > 0 && cur.Ghosts[0].RelationGhost != nil {
					// 获得连锁魂侍一次出手结果 copy the result first
					ghost := cur.Ghosts[0] // 临时保留原本的魂侍
					cur.Ghosts[0] = cur.Ghosts[0].RelationGhost
					cur.useGhostSkillId = int16(cur.Ghosts[0].GhostSkillId)
					result = append(result, cur.fight())
					cur.Ghosts[0] = ghost
				}

				useGhost = true
				cur.useGhostSkillPosition = 0
				cur.useGhostSkillId = 0
				if b.BatType != BT_GHOST_LEVEL {
					cur.UsedGhostSkill = true
				}
				b.enableGhostSkill = false
				break
			}
			continue
		}

		if b.useItemRound {
			if cur.useItemId > 0 {
				result = append(result, cur.fight())
				cur.useItemId = 0
				b.useItemRound = false
				useBattleItem = true
				break
			}
			continue
		}

		result = append(result, cur.fight())

		//如果这次是主动的魂侍攻击那么不算入一回合
		if result[len(result)-1].UseGhostSkill {
			if b.Defenders.live > 0 && cur.Ghosts[0].RelationGhost != nil {
				// 连锁魂侍出手
				ghost := cur.Ghosts[0] // 临时保留原本的魂侍
				cur.Ghosts[0] = cur.Ghosts[0].RelationGhost
				cur.useGhostSkillId = int16(cur.Ghosts[0].GhostSkillId)
				result = append(result, cur.fight())
				cur.Ghosts[0] = ghost
			}
			if b.BatType != BT_GHOST_LEVEL {
				cur.UsedGhostSkill = true
			}
			cur.useGhostSkillPosition = 0
			cur.useGhostSkillId = 0
			useGhost = true
			b.enableGhostSkill = false
			break
		}
		// 防止重新排序后又出手
		cur.fighted = true

		if b.Attackers.live <= 0 || b.Defenders.live <= 0 {
			break
		}

		// 速度buff 技能后需要resort
		if b.needResort {
			order.Reset()
			order.Append(b.Defenders.Fighters, false)
			order.Append(b.Attackers.Fighters, false)

			sort.Sort(order)

			b.needResort = false
			goto START
		}
		CallEnemysArr := result[len(result)-1].CallEnemys
		if len(CallEnemysArr[0]) > 0 || len(CallEnemysArr[1]) > 0 {
			b.summoner = cur
			break //本次出手促发召唤，暂时中断出手
		}
	}

	if b.Attackers.live <= 0 {
		b.Attackers.groupIndex++

		if b.Attackers.groupIndex == len(b.Attackers.Groups) {
			status = DEF_WIN
		} else {
			b.Attackers.Fighters = b.Attackers.Groups[b.Attackers.groupIndex]
			b.UpdateAttackerLiveCount()
			b.needResort = true
			status = ATK_NEXT
		}
	} else if b.Defenders.live <= 0 {
		b.Defenders.groupIndex++
		if b.Defenders.groupIndex == len(b.Defenders.Groups) {
			status = ATK_WIN
		} else {
			b.Defenders.Fighters = b.Defenders.Groups[b.Defenders.groupIndex]
			b.UpdateDefenderLiveCount()
			b.needResort = true
			status = DEF_NEXT
		}
	} else {
		status = NOT_END

		// 这一轮该出手的角色，在上一轮已经被击败，立即开始下一轮计算，防止返回空数据给客户端
		if !b.enableGhostSkill && !useTotem && !b.useItemRound && len(result) == 0 {
			// 已经重新计算一次就不再重新计算，防止死循环
			needRestart = !needRestart
			//当前回合中是有促发怪物召唤的，且召唤的怪物在当前的出手fighter之前的情况不需要restart
			if start > 0 {
				needRestart = false
			}
		}
	}

	// 必须放在排序前，否则回合值会变化
	nowRound = b.GetRounds()

	var nextGroup = status == DEF_NEXT || status == ATK_NEXT
	// 所有人都出手过，或者速度变化 或者切下一组就算新回合
	var newRound = nextGroup || (i == len(order.fighters) && !useGhost && !useBattleItem && !useTotem)

	if newRound || status == ATK_WIN || status == DEF_WIN {
		for i := 0; i < len(order.fighters); i++ {
			fighter := order.fighters[i]
			if fighter != nil && fighter.Health > MIN_HEALTH {
				//所有人刷新buff
				fighter.Buffs.UpdateOnTurnEnd(fighter)
				//主角增加精气
				if fighter.Kind == FK_PLAYER {
					fighter.AddPower(RECOVER_POWER)
				}
			}
		}
	}

	if (status == NOT_END || nextGroup) &&
		(newRound || b.needResort) {
		// 增加回合数
		if newRound {
			b.round++
		}

		order.Reset()
		order.Append(b.Defenders.Fighters, newRound)
		order.Append(b.Attackers.Fighters, newRound)

		sort.Sort(order)

		b.needResort = false
	}

	if needRestart {
		goto START
	}
	if b.summoner != nil {
		status = TRIGGER_CALL_ENEMYS
	}
	return
}

// 获取当前战斗轮数
func (b *BattleState) GetRounds() int {
	return b.round
}

// 自动计算
func (b *BattleState) AutoNextRound(maxRound int, callback func(result []*FightResult, status, nowRound int)) {
	for {
		result, status, nowRound := b.NextRound()

		if nowRound >= maxRound && status == NOT_END {
			// 计算血量 判定输赢
			atkHealths, defHealths := 0, 0
			for _, fighter := range b.Attackers.Fighters {
				if fighter != nil && fighter.Health > MIN_HEALTH {
					atkHealths += fighter.Health
				}
			}

			for _, fighter := range b.Defenders.Fighters {
				if fighter != nil && fighter.Health > MIN_HEALTH {
					defHealths += fighter.Health
				}
			}

			status = ATK_WIN
			if atkHealths < defHealths {
				status = DEF_WIN
			}
		}

		callback(result, status, nowRound)

		if status == ATK_WIN || status == DEF_WIN {
			return
		}
	}
}

// 获取下一个出手玩家id,没有主角在场返回0
func (b *BattleState) GetNextFighterPlayerId() int64 {
	for i := 0; i < len(b.order.fighters); i++ {
		var cur = b.order.fighters[i]

		if !cur.fighted && cur.Health > MIN_HEALTH && (cur.Kind == FK_PLAYER || cur.Kind == FK_BUDDY) {
			return cur.PlayerId
		}
	}

	return 0
}

// 清除概率
func (b *BattleState) ClearRandom() {
	for _, fighter := range b.Attackers.Fighters {
		if fighter != nil {
			fighter.Critial = 0
			fighter.CritialLevel = 0
			fighter.Block = 0
			fighter.BlockLevel = 0
			fighter.Dodge = 0
			fighter.DodgeLevel = 0
			fighter.Hit = 100
		}
	}

	for _, fighter := range b.Defenders.Fighters {
		if fighter != nil {
			fighter.Critial = 0
			fighter.CritialLevel = 0
			fighter.Block = 0
			fighter.BlockLevel = 0
			fighter.Dodge = 0
			fighter.DodgeLevel = 0
			fighter.Hit = 100
		}
	}
}

// 检查我方是否有伙伴死亡
func (b *BattleState) HaveBuddyDeadExcludePet() bool {
	for _, fighter := range b.Attackers.Fighters {
		if fighter != nil && !fighter.IsBattlePet && fighter.Health <= MIN_HEALTH && fighter.Kind == FK_BUDDY {
			return true
		}
	}

	return false
}

func (b *BattleState) NextRound_v2() (result []*FightResult, status, nowRound int) {
	var (
		i     int
		start int
		cur   *Fighter
		order = b.order
	)
	result = b.result[0:0]
START:
	if b.summoner != nil {
		for i = 0; i < len(order.fighters); i++ {
			if order.fighters[i] == b.summoner {
				start = i + 1 // +1是从召唤者下一个战斗对象开始计算战报
				break
			}
		}
		b.summoner = nil
	}

	for i = start; i < len(order.fighters); i++ {
		cur = order.fighters[i]
		if cur.Health <= MIN_HEALTH {
			continue
		}

		b.Attackers.lastLive = b.Attackers.live
		b.Defenders.lastLive = b.Defenders.live
		result = append(result, cur.fight())
		cur.fighted = true
		//这里有个个坑：
		//1. 只能boss才能召唤
		//2. 假设召唤成功，那么召唤者必定存活
		CallEnemysArr := result[len(result)-1].CallEnemys
		callInfos := CallEnemysArr[0]
		callInfos = append(callInfos, CallEnemysArr[1]...)
		if b.runtimeAddMonsterHook != nil && len(callInfos) > 0 && b.runtimeAddMonsterHook(b, result[len(result)-1]) {
			b.summoner = cur
			goto START
			//break //本次出手促发召唤，暂时中断出手
		}

		if b.Attackers.live <= 0 || b.Defenders.live <= 0 {
			break
		}
	}
	status, nowRound = b.handleRoundEnd(i)
	return result, status, nowRound
}

func (b *BattleState) SetSkill(pid int64, isAttacker bool, posIdx int8, skillIdx int8) {
	if isAttacker {
		if fighter := b.Attackers.Fighters[posIdx-1]; fighter != nil && fighter.PlayerId == pid {
			fighter.useSkillIndex = int(skillIdx)
		}
	} else {
		if fighter := b.Defenders.Fighters[posIdx-1]; fighter != nil && fighter.PlayerId == pid {
			fighter.useSkillIndex = int(skillIdx)
		}
	}
}

//回合开始，所有在线用户请求阵印，调用者需要通知 firstPid 可以战斗
func (b *BattleState) RequireTotem() (result []*FightResult, status, nowRound int, firstPid int64) {
	nowRound = b.GetRounds()
	if b.callTotemRound >= nowRound {
		return
	}
	var (
		order = b.order
		cur   *Fighter
	)
	result = b.result[0:0]
	b.launchBattleTotem = true
	for i := 0; i < len(order.fighters); i++ {
		cur = order.fighters[i]
		cur.cacheIdx = 0
		if cur.Kind == FK_PLAYER {
			//释放阵印需要以来 Fighter.fight 方法，实际上这个fighter 可能已经死亡
			//后续可能考虑把 阵印
			if len(cur.side.TotemInfo) > nowRound && cur.side.TotemInfo[nowRound] != nil && !cur.side.TotemInfo[nowRound].used {
				result = append(result, cur.fight())
				cur.side.TotemInfo[nowRound].used = true
			}
			if firstPid <= 0 {
				firstPid = cur.PlayerId
			}
		}
	}
	b.launchBattleTotem = false
	b.currentActionPid = firstPid

	return result, WAITING, nowRound, firstPid
}

func (b *BattleState) UseItem(pid int64, isAttacker bool, posIdx int8, itemId int32) (result []*FightResult, status, nowRound int) {
	if b.currentActionPid != pid {
		return nil, 0, 0
	}
	side := b.Attackers
	result = b.result[0:0]

	if !isAttacker {
		side = b.Defenders
	}
	if fighter := side.Fighters[posIdx-1]; fighter != nil && fighter.Kind == FK_PLAYER && fighter.PlayerId == pid && fighter.Health > MIN_HEALTH /*TODO && fighter.canFighter()/*/ {
		fighter.cacheIdx = 0
		fighter.useItemId = itemId
		b.useItemRound = true
		result = append(result, fighter.fight())
		b.useItemRound = false
		fighter.useItemId = 0
	}
	status, nowRound = b.handleRoundEnd(-1)
	return result, status, nowRound
}

func (b *BattleState) AutoUseGhost(pid int64) (result []*FightResult, status, nowRound int) {
	if b.currentActionPid != pid {
		return nil, 0, 0
	}
	var cur *Fighter
	result = b.result[0:0]
	for i := 0; i < len(b.order.fighters); i++ {
		cur = b.order.fighters[i]
		if cur.Health > MIN_HEALTH && cur.PlayerId == pid && cur.canReleaseGhostSkill() {

			cur.cacheIdx = 0
			cur.useGhostSkillId = int16(cur.Ghosts[0].GhostSkillId)
			result = append(result, cur.fight())
			//连锁魂侍技能释放
			if cur.Health > 0 && b.Defenders.live > 0 && b.Attackers.live > 0 && cur.Ghosts[0].RelationGhost != nil {
				ghost := cur.Ghosts[0] // 临时保留原本的魂侍
				cur.Ghosts[0] = cur.Ghosts[0].RelationGhost
				cur.useGhostSkillId = int16(cur.Ghosts[0].GhostSkillId)
				result = append(result, cur.fight())
				cur.Ghosts[0] = ghost
			}
			cur.useGhostSkillId = 0
			if b.BatType != BT_GHOST_LEVEL {
				cur.UsedGhostSkill = true
			}
			break
		}
	}
	status, nowRound = b.handleRoundEnd(-1)
	return result, status, nowRound
}

func (b *BattleState) UseGhost(pid int64, isAttacker bool, posIdx int8) (result []*FightResult, status, nowRound int) {
	if b.currentActionPid != pid {
		return nil, 0, 0
	}
	side := b.Attackers
	otherSide := b.Defenders
	result = b.result[0:0]

	if !isAttacker {
		side = b.Defenders
		otherSide = b.Attackers
	}
	if fighter := side.Fighters[posIdx-1]; fighter != nil && fighter.Health > MIN_HEALTH && fighter.canReleaseGhostSkill() {
		fighter.cacheIdx = 0
		fighter.useGhostSkillId = int16(fighter.Ghosts[0].GhostSkillId)
		result = append(result, fighter.fight())
		//连锁魂侍技能释放
		if fighter.Health > 0 && otherSide.live > 0 && fighter.Ghosts[0].RelationGhost != nil {
			ghost := fighter.Ghosts[0] // 临时保留原本的魂侍
			fighter.Ghosts[0] = fighter.Ghosts[0].RelationGhost
			fighter.useGhostSkillId = int16(fighter.Ghosts[0].GhostSkillId)
			result = append(result, fighter.fight())
			fighter.Ghosts[0] = ghost
		}
		if b.BatType != BT_GHOST_LEVEL {
			fighter.UsedGhostSkill = true
		}
		fighter.useGhostSkillId = 0
		status, nowRound := b.handleRoundEnd(-1)
		return result, status, nowRound
	}
	return nil, b.getStatus(), b.GetRounds()
}

func (b *BattleState) SetAuto_v2(pid int64) {
	for _, player := range b.Attackers.Players {
		if player.PlayerId == pid {
			player.Auto = true
			return
		}
	}
	for _, player := range b.Defenders.Players {
		if player.PlayerId == pid {
			player.Auto = true
			return
		}
	}
}

func (b *BattleState) IsAuto_v2(pid int64) bool {
	for _, player := range b.Attackers.Players {
		if player.PlayerId == pid {
			return player.Auto
		}
	}
	for _, player := range b.Defenders.Players {
		if player.PlayerId == pid {
			return player.Auto
		}
	}
	return false
}

func (b *BattleState) CancelAuto_v2(pid int64) (joinRound int) {
	for _, player := range b.Attackers.Players {
		if player.PlayerId == pid {
			player.Auto = false
			return
		}
	}
	for _, player := range b.Defenders.Players {
		if player.PlayerId == pid {
			player.Auto = false
			return
		}
	}
	return b.GetRounds() + 1
}

//如果玩家准入ready 状态，如果没有其他玩家在等待，那么开始计算战报
//TODO PrepareReady 不计算战报比较符合语义
func (b *BattleState) PrepareReady(pid int64) (result []*FightResult, status, nowRound int) {
	var (
		anyoneNotReady bool = false
		playerSpeed    int
		nextPlayer     int64
	)
	nowRound = b.GetRounds()
	for _, player := range b.Attackers.Players {
		if player.PlayerId == pid {
			player.PrepareRound = nowRound
			continue
		}
		if player.PrepareRound < nowRound {
			anyoneNotReady = true
			if player.Speed > playerSpeed {
				playerSpeed = player.Speed
				nextPlayer = player.PlayerId
			}
		}
	}
	for _, player := range b.Defenders.Players {
		if player.PlayerId == pid {
			player.PrepareRound = nowRound
			continue
		}
		if player.PrepareRound < nowRound {
			anyoneNotReady = true
			if player.Speed > playerSpeed {
				playerSpeed = player.Speed
				nextPlayer = player.PlayerId
			}
		}
	}
	b.currentActionPid = nextPlayer
	if anyoneNotReady {
		return nil, WAITING, nowRound
	}
	//所有人都准备好了那么开始计算战报
	result, status, nowRound = b.NextRound_v2()
	return result, status, nowRound
}

//每次产生有效战报后需要调用 handleRoundEnd
//curFighterIdx 是当前最后一个出手玩家索引
//魂侍技能 战斗道具 阵印 等不计算回合输的出手传 -1
func (b *BattleState) handleRoundEnd(curFighterIdx int) (status, nowRound int) {
	nowRound = b.GetRounds()
	var newRound = curFighterIdx == len(b.order.fighters) || b.Attackers.live <= 0 || b.Defenders.live <= 0
	if newRound {
		//每个玩家刷新buff 并且 主角增加精气
		for i := 0; i < len(b.order.fighters); i++ {
			fighter := b.order.fighters[i]
			if fighter != nil && fighter.Health > MIN_HEALTH {
				//所有人刷新buff
				fighter.Buffs.UpdateOnTurnEnd(fighter)
				//主角增加精气
				if fighter.Kind == FK_PLAYER {
					fighter.AddPower(RECOVER_POWER)
				}
				if fighter.IsBattlePet && (nowRound-fighter.BattlePetLiveStartRound+1 >= fighter.BattlePetLiveRound) {
					setDisappearWithBattlePet(fighter)
				}
			}
		}
		b.round++
		b.currentActionPid = 0
		b.order.Reset()
		b.order.Append(b.Defenders.Fighters, true)
		b.order.Append(b.Attackers.Fighters, true)
		sort.Sort(b.order)
	}
	status = b.getStatus()
	if status == ATK_NEXT || status == DEF_NEXT {
		for _, player := range b.Attackers.Players {
			if player.PrepareRound < b.round-1 {
				player.PrepareRound = b.round - 1
			}
		}
		for _, player := range b.Defenders.Players {
			if player.PrepareRound < b.round-1 {
				player.PrepareRound = b.round - 1
			}
		}
	}
	if b.needResort {
		orderInit(b, true)
	}
	return status, nowRound
}

func (b *BattleState) getStatus() (status int) {
	if b.Attackers.live <= 0 {
		b.currentActionPid = 0
		b.Attackers.groupIndex++

		if b.Attackers.groupIndex == len(b.Attackers.Groups) {
			status = DEF_WIN
		} else {
			b.Attackers.Fighters = b.Attackers.Groups[b.Attackers.groupIndex]
			b.UpdateAttackerLiveCount()
			b.needResort = true
			status = ATK_NEXT
		}
	} else if b.Defenders.live <= 0 {
		b.currentActionPid = 0
		b.Defenders.groupIndex++
		if b.Defenders.groupIndex == len(b.Defenders.Groups) {
			status = ATK_WIN
		} else {
			b.Defenders.Fighters = b.Defenders.Groups[b.Defenders.groupIndex]
			b.UpdateDefenderLiveCount()
			b.needResort = true
			status = DEF_NEXT
		}
	} else if b.currentActionPid > 0 {
		status = WAITING
	} else {
		status = NOT_END
	}
	return status
}

//战斗中怪物使用召唤技能召唤出新的怪物，在这个钩子函数里面实现通知以及修改战场等操作
func (b *BattleState) RegisterAddMonsterHook(hook func(*BattleState, *FightResult) bool) {
	b.runtimeAddMonsterHook = hook
}

func (b *BattleState) AutoNextRound_v2(callback func(result []*FightResult, status, nowRound int) bool) {
	var (
		nextPid      int64
		results      []*FightResult
		status       int = b.getStatus()
		nowRound     int = 0
		currentRound int = 0
	)
	//多个战斗小组支持和最大回合数支持 在callback中实现
	for {
		//外层for循环代表一个回合，需要获取保留当前回合，
		//因为 nowRound 在在本回合结束就会加一，而GetNextPlayer 需要一个正确的回合数
		currentRound = b.GetRounds()
		if status == ATK_WIN || status == DEF_WIN {
			break
		}
		//RequireTotem 计算阵印战报 并且负责设置 battle.currentActionPid
		results, status, nowRound, _ = b.RequireTotem()
		if len(results) > 0 {
			if !callback(results, status, nowRound) {
				return
			}
		}
		//假设nextPid == 0 那么意味着所有玩家在 currentRound 回合已经行动
		for nextPid = b.GetNextPlayer(currentRound); nextPid > 0; nextPid = b.GetNextPlayer(currentRound) {
			results, status, nowRound = b.AutoUseGhost(nextPid)
			if len(results) > 0 {
				if !callback(results, status, nowRound) {
					return
				}
				continue
			}
			results, status, nowRound = b.PrepareReady(nextPid)
			if len(results) > 0 {
				if !callback(results, status, nowRound) {
					return
				}
			}
		}
	}

}
