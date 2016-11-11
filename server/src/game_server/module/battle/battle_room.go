package battle

import (
	"core/debug"
	"core/fail"
	"core/log"
	"core/net"
	//"fmt"
	"game_server/api/protocol/battle_api"
	"game_server/battle"
	"game_server/dat/battle_pet_dat"
	//"game_server/mdb"
	"game_server/module"
	"sync"
	//"sync/atomic"
	"time"
)

const (
	BATTLE_READY_TIMEOUT = 10 // 战斗准备超时时间
)

// 战场房间状态
const (
	ROOM_STATE_READY = iota + 1 // 加载完战斗界面
	ROOM_STATE_END              // 战斗结束
)

// 结束判定方式设定
const (
	ROOM_END_BY_NORMAL = iota
	ROOM_END_BY_HEALTH
	ROOM_END_BY_SIDE_OFFLINE
)

const (
	CMD_CALLBACK = iota
	CMD_LEAVE_ROOM
	CMD_NEXT_ROUND
	CMD_START_READY

	// V2 additional
	CMD_INIT_ROUND
	CMD_USE_GHOST
	CMD_USE_ITEM
	CMD_ESCAPE
	CMD_CALL_PET
	CMD_SET_SKILL
	CMD_PREPARE_READY
	CMD_SET_AUTO
	CMD_CANCEL_AUTO
)

type battleCMD struct {
	cmd      int
	param    interface{}
	callback func()
}

type BattleRoomPlayer struct {
	playerId   int64
	isAttacker bool // 是否是攻击方
	isOnline   bool // 是否在线

	status  int         // 玩家状态，战斗播放状态，进入当前轮数，当前轮数操作中
	opTimer *time.Timer // 玩家进入操作状态的超时器
}

const (
	ROOMP_STAT_FIGHT   = iota // 玩家操作完成，进入战斗状态
	ROOMP_STAT_INITED         // 玩家前一回合战斗播放完成，进入新一回合
	ROOMP_STAT_OPERATE        // 玩家进入操作状态
	ROOMP_STAT_OFFLINE        // 玩家离线状态
)

var (
	battleRooms      = make(map[int64]*BattleRoom, 1000)
	battleRoomId     = int64(1)
	battleRoomsMutex sync.Mutex
)

const (
	skillAnimationDelay = 2   //攻击延时
	ghostAnimationDelay = 2.5 //魂侍出场延时
)

type BattleRoom struct {
	module.BasicBattle
	battleId      int64               // 战斗id
	battle        *battle.BattleState // 基础战斗类
	state         int8                // 战场房间状态
	roundTime     time.Duration       // 每轮时长配置
	presetEndType int                 // 预设结束方式
	endType       int                 // 实际结束方式

	isAuto bool //自动战斗

	startReadyTimer *time.Timer // 开始战斗前加载超时计算器

	roundTimer *time.Timer // 每轮战斗超时机制
	endTimer   *time.Timer // 战斗时间限定

	// 按状态做回调
	OnBattleRoomEvent module.IBattleRoom

	players  map[int64]*BattleRoomPlayer // 战斗过程中的玩家状态
	autoPids []int64                     // 被托管的玩家id

	playerChannel *net.Channel // 参赛者

	playerNum int8 // 参与战斗人数（冗余）
	readyNum  int8 // 进入战斗场景人数（冗余）
	onlineNum int8 // 在线人数（冗余）

	RoundSendNum int16
	lastSendTime time.Time

	//cmdChan 被 sendCmdFromClient 和 doClose 使用
	//他们不是串行的， doClose 关闭 cmdChan ，可能导致 sendCmdFromClient 往关闭的channel写
	//加一个读写所保护 cmdChan 在 sendCmdFromClient 被使用的时候不被关闭
	chanLock sync.RWMutex
	isClosed bool
	cmdChan  chan *battleCMD

	skillUsage map[int64]map[int]int //伙伴技能

	totalAnimationDelay float64 //魂侍技能以及战斗道具播放延时
}

// 初始战斗房间，返回房间及是否双方都掉线
func NewBattleRoom(battleType int, onRoomEvent module.IBattleRoom, atkSide, defSide *battle.SideInfo, battleLastInitFunc func(), playerChannel *net.Channel, roundTime time.Duration, recordBattle bool) (
	br *BattleRoom, allIsSupend bool) {

	battleRoomsMutex.Lock()
	defer battleRoomsMutex.Unlock()

	// 构造战斗房间中各个玩家状态实例
	playerNum := int8(len(atkSide.Players) + len(defSide.Players))
	players := make(map[int64]*BattleRoomPlayer, playerNum)

	skillUsage := make(map[int64]map[int]int)
	var pid int64
	for _, attacker := range atkSide.Players {
		pid = attacker.PlayerId
		players[pid] = &BattleRoomPlayer{playerId: pid, isAttacker: true}
		skillUsage[pid] = make(map[int]int)

		//伙伴技能初始化
		for _, fighter := range atkSide.Fighters {
			if fighter != nil && fighter.PlayerId == pid {
				for _, skillInfo := range fighter.SkillInfos {
					if skillInfo != nil && skillInfo.MaxReleaseNum > 0 {
						skillUsage[pid][skillInfo.SkillId] = skillInfo.MaxReleaseNum
					}
				}
			}
		}
	}

	for _, defender := range defSide.Players {
		pid = defender.PlayerId
		players[pid] = &BattleRoomPlayer{playerId: pid}
	}

	// 构造战斗房间的实例
	b := battle.Start(battleType, atkSide, defSide, battleLastInitFunc)
	br = &BattleRoom{
		battleId:          battleRoomId,
		battle:            b,
		roundTime:         roundTime,
		players:           players,
		autoPids:          make([]int64, 0, playerNum),
		playerChannel:     playerChannel,
		onlineNum:         0,
		playerNum:         playerNum,
		OnBattleRoomEvent: onRoomEvent,
		cmdChan:           make(chan *battleCMD, 10),
		skillUsage:        skillUsage,
		isAuto:            true,
	}

	br.startRecvCmd()
	battleRooms[battleRoomId] = br
	battleRoomId++

	var state *module.SessionState
	playerChannel.Fetch(func(session *net.Session) {
		state = module.State(session)
		players[state.PlayerId].isOnline = true

		state.Battle = br
		br.onlineNum++
	})

	if recordBattle {
		b.EnableRecord()
	}

	module.API.Broadcast(br.playerChannel, GetStartBattleResponse(br.battle, nil, br.skillUsage, nil))

	// 特殊情况：所有玩家都不在线
	if br.onlineNum == 0 {
		allIsSupend = true
		return
	}

	return
}

func (this *BattleRoom) startRecvCmd() {
	go func() {
		var cmd *battleCMD

		defer func() {
			if err := recover(); err != nil {
				log.Errorf(`battle room recv cmd
Error = %v
Cmd	  = %v
Stack =
%s`,
					err,
					cmd,
					debug.Stack(1, "    "),
				)
			}
		}()

		for {
			select {
			case cmd, ok := <-this.cmdChan:
				if cmd == nil || !ok {
					return
				}
				switch cmd.cmd {
				case CMD_CALLBACK:
					cmd.callback()
				case CMD_LEAVE_ROOM:
					this.cmdLeaveBattle(cmd.param.(*net.Session))
				case CMD_NEXT_ROUND:
					panic("废弃")
				case CMD_START_READY:
					panic("废弃")

				// V2 additional
				case CMD_INIT_ROUND:
					//fmt.Println("CMD_INIT_ROUND", cmd.param.(int64))
					this.cmdInitRound(cmd.param.(int64))
				case CMD_USE_GHOST:
					//fmt.Println("CMD_USE_GHOST")
					this.cmdUseGhost(cmd.param.(*UseGhostParams))
				case CMD_USE_ITEM:
					//fmt.Println("CMD_USE_ITEM", cmd.param.(int64))
					this.cmdUseItem(cmd.param.(*UseItemParams))
				case CMD_ESCAPE:
					this.cmdEscape(cmd.param.(int64))
				case CMD_CALL_PET:
					this.cmdCallPet(cmd.param.(*CallPetParams))
				case CMD_SET_SKILL:
					this.cmdSetSkill(cmd.param.(*SetSkillParams))
				case CMD_PREPARE_READY:
					//fmt.Println("CMD_PREPARE_READY", cmd.param.(int64))
					this.cmdPrepareReady(cmd.param.(int64))
				case CMD_SET_AUTO:
					this.cmdSetAuto(cmd.param.(int64))
				case CMD_CANCEL_AUTO:
					this.cmdCancelAuto(cmd.param.(int64))
				}
			}
			if this.state == ROOM_STATE_END {
				return
			}
		}
	}()
}

// 预设结束类型
func (this *BattleRoom) doPresetEndType(endType int) {
	this.presetEndType = endType
}

// 获取结束类型
func (this *BattleRoom) getEndType() int {
	return this.endType
}

func (this *BattleRoom) GetBattle() *battle.BattleState {
	return this.battle
}

func (this *BattleRoom) Relive(session *net.Session) {

}

func (this *BattleRoom) UseBuddySkill(session *net.Session, posIndex int8, skillIndex int8) {
	state := module.State(session)
	side := this.battle.Attackers
	buddy := side.Fighters[posIndex-1]
	fail.When(buddy == nil, "指定位置不存在伙伴")
	fail.When(buddy.PlayerId != state.PlayerId, "操作不属于自己的伙伴")
	buddy.SetBuddySkill(int(skillIndex))
}

func (this *BattleRoom) sendCmdFromClient(cmd *battleCMD) {
	this.chanLock.RLock()
	defer this.chanLock.RUnlock()
	if this.isClosed {
		return
	}
	defer func() {
		if err := recover(); err != nil {
			log.Errorf(`module.battle.battle_room.sendCmdFromClient
		Error = %v
		Stack = 
		%s`,
				err,
				debug.Stack(1, "    "),
			)
		}
	}()
	select {
	case this.cmdChan <- cmd:
	default:
		//go func() {
		//	this.cmdChan <- cmd
		//}()
		go this.sendCmdFromClient(cmd)
	}
}

func (this *BattleRoom) handleBattleResult(results []*battle.FightResult, status, nowRound int) {
	this.RoundSendNum++
	this.lastSendTime = time.Now()
	// 广播会进行战斗的封包
	module.API.Broadcast(this.playerChannel, GetNextRoundResponse(results, status, nowRound, this.battle, nil /*this.autoPids*/, nil))

	this.endType = ROOM_END_BY_NORMAL

	// 执行外部逻辑
	this.OnBattleRoomEvent.OnNextRound(results, status, nowRound)

	// 战斗结束对会话的值才重置（拷贝）
	if status == battle.ATK_WIN || status == battle.DEF_WIN {
		this.doClose()
	}
}

func (this *BattleRoom) InitRound(session *net.Session) {
	this.sendCmdFromClient(&battleCMD{
		cmd:   CMD_INIT_ROUND,
		param: module.State(session).PlayerId,
	})
}

func (this *BattleRoom) cmdInitRound(pid int64) {
	// 检查玩家状态
	if this.players[pid].status != ROOMP_STAT_FIGHT {
		return
	}

	// 设置玩家状态
	this.players[pid].status = ROOMP_STAT_INITED

	// 检查进入当前回合玩家数
	playerCount := 0
	initedCount := 0
	for _, player := range this.players {
		playerCount++
		if player.status == ROOMP_STAT_INITED || player.status == ROOMP_STAT_OFFLINE {
			initedCount++
		}
	}

	// 进入回合玩家数过半，开始
	if playerCount/2 <= initedCount {

		// 将所有玩家强制置为准备完成
		for _, player := range this.players {
			if player.status != ROOMP_STAT_INITED && player.status != ROOMP_STAT_OFFLINE {
				player.status = ROOMP_STAT_INITED
			}
		}

		// 回合开始，图腾出场
		results, status, currentRound, _ := this.battle.RequireTotem()
		if len(results) > 0 {
			this.handleBattleResult(results, status, currentRound)
		}

		// 修改第一位玩家进入操作状态
		nextPid := this.battle.GetNextPlayer(currentRound)
		for ; nextPid > 0 && this.players[nextPid].status == ROOMP_STAT_OFFLINE; nextPid = this.battle.GetNextPlayer(currentRound) {
			// 跳过离线玩家
			//自动使用魂侍
			results, status, nowRound := this.battle.AutoUseGhost(nextPid)
			if len(results) > 0 {
				this.handleBattleResult(results, status, nowRound)
				continue
			}
			results, status, nowRound = this.battle.PrepareReady(nextPid)
			if len(results) > 0 {
				this.handleBattleResult(results, status, nowRound)
			}
		}
		if nextPid > 0 {
			this.players[nextPid].status = ROOMP_STAT_OPERATE
			// 设置操作超时器
			this.players[nextPid].opTimer = time.AfterFunc(this.roundTime, func() {
				this.sendCmdFromClient(&battleCMD{
					cmd:   CMD_PREPARE_READY,
					param: nextPid,
				})
			})
			module.API.Broadcast(this.playerChannel,
				&battle_api.NotifyReady_Out{
					Pid: nextPid,
				})

		}
	}
}

func (this *BattleRoom) UseGhost(session *net.Session, isAttacker bool, posIdx int8) {
	//fmt.Println("UseGhost Interface", module.State(session).PlayerId)
	this.sendCmdFromClient(&battleCMD{
		cmd: CMD_USE_GHOST,
		param: &UseGhostParams{
			Pid:        module.State(session).PlayerId,
			IsAttacker: isAttacker,
			PosIdx:     posIdx,
		},
	})
}

type UseGhostParams struct {
	Pid        int64
	IsAttacker bool
	PosIdx     int8
}

func (this *BattleRoom) cmdUseGhost(params *UseGhostParams) {
	//fmt.Println("cmdUseGhost", params.Pid, this.players[params.Pid].status)
	// 检查玩家状态
	if this.players[params.Pid].status != ROOMP_STAT_OPERATE {
		return
	}

	// 使用魂侍
	results, status, nowRound := this.battle.UseGhost(params.Pid, params.IsAttacker, params.PosIdx)
	if len(results) > 0 {
		this.handleBattleResult(results, status, nowRound)
	}
}

func (this *BattleRoom) UseItem(session *net.Session, isAttacker bool, posIdx int8, itemId int16) {
	this.sendCmdFromClient(&battleCMD{
		cmd: CMD_USE_ITEM,
		param: &UseItemParams{
			Pid:        module.State(session).PlayerId,
			IsAttacker: isAttacker,
			PosIdx:     posIdx,
			ItemId:     itemId,
		},
	})
}

type UseItemParams struct {
	Pid        int64
	IsAttacker bool
	PosIdx     int8
	ItemId     int16
}

func (this *BattleRoom) cmdUseItem(params *UseItemParams) {
	// 检查玩家状态
	if this.players[params.Pid].status != ROOMP_STAT_OPERATE {
		return
	}

	// 使用道具
	results, status, nowRound := this.battle.UseItem(params.Pid, params.IsAttacker, params.PosIdx, int32(params.ItemId))
	if len(results) > 0 {
		this.handleBattleResult(results, status, nowRound)
	}
}

func (this *BattleRoom) SetSkill(session *net.Session, posIdx int8, skillIdx int8) {
	this.sendCmdFromClient(&battleCMD{
		cmd: CMD_SET_SKILL,
		param: &SetSkillParams{
			Pid:      module.State(session).PlayerId,
			PosIdx:   posIdx,
			SkillIdx: skillIdx,
		},
	})
}

type SetSkillParams struct {
	Pid      int64
	PosIdx   int8
	SkillIdx int8
}

func (this *BattleRoom) cmdSetSkill(params *SetSkillParams) {
	// 检查玩家状态
	if this.players[params.Pid].status != ROOMP_STAT_OPERATE {
		return
	}

	// 查一下玩家的立场，为兼容旧接口，这段为临时代码
	isAttacker := (this.battle.Attackers.Fighters[params.PosIdx-1].PlayerId == params.Pid)

	// 设置技能
	this.battle.SetSkill(params.Pid, isAttacker, params.PosIdx, params.SkillIdx)
}

func (this *BattleRoom) PrepareReady(session *net.Session, __isAuto bool) {
	//fmt.Println("PrepareReady recive", module.State(session).PlayerId)
	this.sendCmdFromClient(&battleCMD{
		cmd:   CMD_PREPARE_READY,
		param: module.State(session).PlayerId,
	})
}

func (this *BattleRoom) cmdPrepareReady(pid int64) {
	//fmt.Println("cmdPrepareReady pid arg ", pid)
	// 检查玩家状态
	if this.players[pid].status != ROOMP_STAT_OPERATE {
		return
	}

	// 修改玩家状态
	this.players[pid].status = ROOMP_STAT_FIGHT
	// 取消操作超时
	this.players[pid].opTimer.Stop()
	this.players[pid].opTimer = nil
	//fmt.Println("cmdPrepareReady player ready ", pid)
	var (
		results  []*battle.FightResult
		status   int
		nowRound int
	)

	oldRound := this.battle.GetRounds()
	if this.isAuto {
		//自动使用魂侍
		results, status, nowRound = this.battle.AutoUseGhost(pid)
		if len(results) > 0 {
			this.handleBattleResult(results, status, nowRound)
		}

	}
	nextPid := this.battle.GetNextPlayer(oldRound)
	if nextPid <= 0 {
		return
	}

	results, status, nowRound = this.battle.PrepareReady(pid)
	if len(results) > 0 {
		this.handleBattleResult(results, status, nowRound)
	}
	nextPid = this.battle.GetNextPlayer(oldRound)

	for nextPid > 0 && nextPid != pid && this.players[nextPid].status == ROOMP_STAT_OFFLINE {
		//fmt.Println("cmdPrepareReady for offline player", nextPid)
		// 刚离线，未自动战斗的玩家
		if !this.battle.IsAuto_v2(nextPid) {
			this.battle.SetAuto_v2(nextPid)
		}
		//自动使用魂侍
		results, status, nowRound = this.battle.AutoUseGhost(nextPid)
		if len(results) > 0 {
			this.handleBattleResult(results, status, nowRound)
			nextPid = this.battle.GetNextPlayer(oldRound)
			continue
		}
		//自动使用准备
		results, status, nowRound = this.battle.PrepareReady(nextPid)
		if len(results) > 0 {
			this.handleBattleResult(results, status, nowRound)
		}
		nextPid = this.battle.GetNextPlayer(oldRound)
	}
	//fmt.Println("cmdPrepareReady next player ", nextPid)
	if nextPid > 0 {
		// 修改下一位玩家状态
		this.players[nextPid].status = ROOMP_STAT_OPERATE
		// 设置操作超时器
		this.players[nextPid].opTimer = time.AfterFunc(this.roundTime, func() {
			this.sendCmdFromClient(&battleCMD{
				cmd:   CMD_PREPARE_READY,
				param: nextPid,
			})
		})

		module.API.Broadcast(this.playerChannel,
			&battle_api.NotifyReady_Out{
				Pid: nextPid,
			})
	}
}

func (this *BattleRoom) SetAuto(session *net.Session) {
	this.sendCmdFromClient(&battleCMD{
		cmd:   CMD_SET_AUTO,
		param: module.State(session).PlayerId,
	})
}

func (this *BattleRoom) cmdSetAuto(pid int64) {
	// 检查玩家状态
	if this.players[pid].status != ROOMP_STAT_OPERATE {
		return
	}

	this.battle.SetAuto_v2(pid)
}

func (this *BattleRoom) CancelAuto(session *net.Session) {
	this.sendCmdFromClient(&battleCMD{
		cmd:   CMD_CANCEL_AUTO,
		param: module.State(session).PlayerId,
	})
}

func (this *BattleRoom) cmdCancelAuto(pid int64) {
	// 检查玩家状态
	if this.players[pid].status != ROOMP_STAT_OPERATE {
		return
	}

	round := this.battle.CancelAuto_v2(pid)
	this.playerChannel.FetchOne(func(session *net.Session) {
		session.Send(&battle_api.CancelAuto_Out{
			Round: int16(round),
		})
	}, uint64(pid))
}

// 设置下一回数据
func (this *BattleRoom) NextRound(params *module.NextRoundParams) {
	// 屏蔽道具使用
	if params.UseItemId > 0 {
		return
	}

	this.sendCmdFromClient(&battleCMD{cmd: CMD_NEXT_ROUND, param: params})
}

// 自动战斗特定回合，最后判断血量确定输赢
func (this *BattleRoom) autoBattle(round int) {
	//TODO 最大回合实现
	this.battle.AutoNextRound_v2(func(result []*battle.FightResult, status, nowRound int) bool {
		if len(result) > 0 {
			this.handleBattleResult(result, status, nowRound)
		}
		if status == battle.ATK_WIN || status == battle.DEF_WIN {
			return false
		}
		return true
	})

}

// 关闭战斗房间
func (this *BattleRoom) doClose() {
	if this.state == ROOM_STATE_END {
		return
	}

	this.state = ROOM_STATE_END

	this.playerChannel.Fetch(func(session *net.Session) {
		this.battleClear(session)
	})

	for _, player := range this.players {
		if player.opTimer != nil {
			player.opTimer.Stop()
		}
	}
	this.chanLock.Lock()
	this.isClosed = true
	close(this.cmdChan)
	this.chanLock.Unlock()

	if this.endTimer != nil {
		this.endTimer.Stop()
		this.endTimer = nil
	}

	delete(battleRooms, this.battleId)
}

func (this *BattleRoom) battleClear(session *net.Session) {
	module.State(session).Battle = nil
}

func (this *BattleRoom) LeaveBattle(session *net.Session) {
	this.sendCmdFromClient(&battleCMD{cmd: CMD_LEAVE_ROOM, param: session})
}

// 掉线处理
func (this *BattleRoom) cmdLeaveBattle(session *net.Session) {
	state := module.State(session)
	playerId := state.PlayerId

	// 清除战场关联的玩家session
	this.battleClear(session)

	this.playerChannel.Exit(session)

	playerState, exist := this.players[playerId]

	// 观战玩家
	if !exist {
		return
	}

	playerState.isOnline = false
	playerState.status = ROOMP_STAT_OFFLINE
	this.onlineNum--

	// 所有玩家都掉线则战斗结束
	if this.onlineNum == 0 {
		this.OnBattleRoomEvent.OnAllLeave()

		this.endType = ROOM_END_BY_NORMAL
		this.doClose()
		return
	}

	// 按没在线判定输赢
	if this.presetEndType == ROOM_END_BY_SIDE_OFFLINE {
		var (
			isAttacker = playerState.isAttacker
			hadOnline  = false
		)

		for _, player := range this.players {
			if player.isAttacker == isAttacker && player.isOnline {
				hadOnline = true
				break
			}
		}

		if !hadOnline {
			// 设置输赢
			var status int = battle.ATK_WIN
			if isAttacker {
				status = battle.DEF_WIN
			}
			this.setEnd(status, ROOM_END_BY_SIDE_OFFLINE)
		}
	}
}

func battleRoomPrintFilter(typeName string, column string) bool {
	return column == "State"
}

// 直接设置输赢
func (this *BattleRoom) setEnd(status, endByType int) {
	if this.state == ROOM_STATE_END {
		return
	}

	// 战斗结果设定
	module.API.Broadcast(this.playerChannel, &battle_api.End_Out{Status: battle_api.RoundStatus(status)})

	this.endType = endByType

	var (
		result   []*battle.FightResult
		nowRound int = this.battle.GetRounds()
	)

	this.OnBattleRoomEvent.OnNextRound(result, status, nowRound)

	this.doClose()
}

func (this *BattleRoom) Escape(session *net.Session) {
}

func (this *BattleRoom) cmdEscape(pid int64) {
	// TODO
}

func (this *BattleRoom) CallBattlePet(session *net.Session, petId int32, petLevel int16, petSkillLv int16) bool {
	// 屏蔽召唤灵宠
	return false

	battlePet := battle_pet_dat.GetBattlePetWithEnemyId(petId)
	pos := module.GetBattlePetPos(battlePet.LivePos, this.battle.Attackers.Fighters)
	if pos < 0 {
		return false
	}

	this.sendCmdFromClient(&battleCMD{callback: func() {

		playerId := module.State(session).PlayerId
		var playerFighter *battle.Fighter
		for _, f := range this.battle.Attackers.Fighters {
			if f != nil && f.Kind == battle.FK_PLAYER && f.PlayerId == playerId {
				fail.When(f.Power < int(battlePet.CostPower), "召唤灵宠精气不足")
				playerFighter = f
				break
			}
		}

		fail.When(playerFighter == nil, "player fighter not found")
		playerFighter.Power -= int(battlePet.CostPower)

		petFighter := module.NewFighterForBattlePet(playerId, battlePet.PetId, pos, battlePet.RoundAttack, battlePet.LiveRound, this.battle.GetRounds(), petLevel, petSkillLv)
		// 添加灵宠到战场
		this.battle.RuntimeAddFighter(battle.FT_ATK, petFighter)

		rsp := GetCallBattlePetResponse(playerFighter, petFighter)

		module.API.Broadcast(this.playerChannel, rsp)
	}, /*end of callback function literals*/
	}) /*end of battleCMD struct literals*/
	return true
}

type CallPetParams struct {
}

func (this *BattleRoom) cmdCallPet(params *CallPetParams) {
	// TODO
}

func (this *BattleRoom) CallNewEnemy(enemy battle.CallInfo) *battle.Fighter {
	return nil
}
