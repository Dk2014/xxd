package module

import (
	"core/net"
	"game_server/battle"
	"game_server/dat/battle_pet_dat"
	"game_server/dat/enemy_deploy_dat"
	"game_server/dat/mission_dat"
	"game_server/dat/monster_property_addition_dat"
	"game_server/dat/role_dat"
	"game_server/dat/totem_dat"
	"game_server/mdb"
	"math/rand"
)

type NextRoundParams struct {
	Session    *net.Session
	PlayerId   int64
	SkillIndex int
	AutoFight  bool
	IsAttacker bool
	Position   int
	JobIndex   int

	IsSystemSet  bool
	SendNum      int16
	UseSwordSoul bool

	UseGhostSkillPosition int
	UseGhostSkillId       int16

	UseItemId int16

	LaunchBattleTotem bool
}

type IBattle interface {
	GetBattle() *battle.BattleState
	LeaveBattle(session *net.Session)

	NextRound(*NextRoundParams)
	Escape(session *net.Session)

	// 复活
	Relive(session *net.Session)

	// 召唤灵宠
	CallBattlePet(session *net.Session, petId int32, petLevel int16, petSkillLv int16) bool

	//召唤怪物
	//CallNewEnemy(enemy battle.CallInfo) *battle.Fighter

	//使用伙伴技能
	UseBuddySkill(session *net.Session, posIndex int8, skillIndex int8)

	//V2接口
	//不保证线程安全，操作需要加锁或者用channel串行化

	//设置角色技能
	SetSkill(session *net.Session, posIdx int8, skillIdx int8)

	//战场中所有在线并且不在自动战斗状态的人都会请求
	InitRound(session *net.Session)

	UseGhost(session *net.Session, isAttacker bool, posIdx int8)
	UseItem(session *net.Session, isAttacker bool, posIdx int8, itemId int16)
	SetAuto(session *net.Session)
	CancelAuto(session *net.Session)
	PrepareReady(session *net.Session, isAuto bool)
}

type IBattleRoom interface {
	OnStartReady()
	OnNextRound(result []*battle.FightResult, status, nowRound int) // 处理每轮回调
	OnAllLeave()                                                    // 玩家都掉线
}

var (
	GetBattleBiz = BattleBiz{}
)

// 为玩家角色初始化战场队伍
func NewBattleSideWithPlayerDatabase(db *mdb.Database, mainRoleOnly, autoFight, isCountFightersNum bool) (*battle.SideInfo, int32) {
	playerForm := db.Lookup.PlayerFormation(db.PlayerId())
	return NewBattleSideWithPlayerDatabaseAndPlayerForm(db, playerForm, mainRoleOnly, autoFight, isCountFightersNum)
}

func NewBattleSideWithPlayerDatabaseAndPlayerForm(db *mdb.Database, playerForm *mdb.PlayerFormation, mainRoleOnly, autoFight, isCountFightersNum bool) (battleSideInfo *battle.SideInfo, totalFighterNum int32) {
	inFormRoleInfos, _ := GetBattleBiz.GetInFormRoleForBattle(db, playerForm, mainRoleOnly)

	fighters := make([]*battle.Fighter, ALL_FIGHTER_POS_NUM)

	_, totalFighterNum = GetBattleBiz.SetFighters(db, inFormRoleInfos, fighters, autoFight, isCountFightersNum, FIGHT_FOR_ALL)

	battleSideInfo = new(battle.SideInfo)
	battleSideInfo.Groups = [][]*battle.Fighter{fighters}
	battleSideInfo.Fighters = fighters
	battleSideInfo.Players = []*battle.BattlePlayerInfo{
		&battle.BattlePlayerInfo{
			PlayerId: db.PlayerId(),
			Auto:     autoFight,
		},
	}
	battleSideInfo.TotemInfo = NewBattleTotemInfo(db)
	return
}

// 为怪物初始化战场队伍
func NewBattleSideWithEnemyDeployForm(battleType int8, parentId int32) *battle.SideInfo {
	enemyForms := enemy_deploy_dat.GetEnemyDeployForm(battleType, parentId)

	enemies := make([][]*battle.Fighter, 0, len(enemyForms))

	for _, form := range enemyForms {
		fighters := make([]*battle.Fighter, ALL_FIGHTER_POS_NUM)

		for i, enemy_id := range form {
			if enemy_id == 0 {
				continue
			}
			//如果需要为怪物进行属性加成，请修改NewFighterForEnemy第二个参数
			fighters[i] = NewFighterForEnemy(enemy_id, 0, i+1)
		}

		enemies = append(enemies, fighters)
	}

	return &battle.SideInfo{
		Groups:   enemies,
		Fighters: enemies[0],
	}
}

// 为怪物创建战斗对象
func NewFighterForEnemy(enemy_id int32, enemy_level int32, pos int) *battle.Fighter {
	enemy := mission_dat.GetEnemyRole(enemy_id)
	propertyAdd := monster_property_addition_dat.GetMonsterPropertyAdd(enemy_id, enemy_level)
	return &battle.Fighter{
		Kind:                battle.FK_ENEMY,
		Position:            pos,
		IsBoss:              enemy.IsBoss,
		Level:               int(enemy.Level),                                                 // 等级
		Prop:                enemy.Prop,                                                       // 种族
		Speed:               float64(enemy.Speed + propertyAdd.Speed),                         // 速度
		Dodge:               float64(enemy.Dodge + propertyAdd.Dodge),                         // 闪避
		Health:              int(enemy.Health + propertyAdd.Health),                           // 生命
		MaxHealth:           int(enemy.Health + propertyAdd.Health),                           // 最大生命
		Attack:              float64(enemy.Attack + propertyAdd.Attack),                       // 攻击
		Defend:              float64(enemy.Defence + propertyAdd.Defence),                     // 防御
		Critial:             float64(enemy.Critial + propertyAdd.Critial),                     // 暴击
		Hit:                 float64(enemy.Hit + propertyAdd.Hit),                             // 命中
		Block:               float64(enemy.Block + propertyAdd.Block),                         // 格挡
		CritialHurt:         float64(enemy.CritialHurt + propertyAdd.CritialHurt),             // 必杀
		Tenacity:            float64(enemy.Toughness + propertyAdd.Toughness),                 // 韧性
		Destroy:             float64(enemy.Destroy + propertyAdd.Destroy),                     // 破击
		Sleep:               float64(enemy.Sleep + propertyAdd.Sleep),                         // 睡眠抗性
		Dizziness:           float64(enemy.Dizziness + propertyAdd.Dizziness),                 // 眩晕抗性
		Random:              float64(enemy.Random + propertyAdd.Random),                       // 混乱抗性
		DisableSkill:        float64(enemy.DisableSkill + propertyAdd.DisableSkill),           // 封魔抗性
		Poisoning:           float64(enemy.Poisoning + propertyAdd.Poisoning),                 // 中毒抗性
		SunderMaxValue:      int(enemy.SunderMaxValue + propertyAdd.SunderMaxValue),           // 护甲值
		SunderMinHurtRate:   int(enemy.SunderMinHurtRate + propertyAdd.SunderMinHurtRate),     // 破甲起始伤害转换率
		SunderEndHurtRate:   int(enemy.SunderEndHurtRate + propertyAdd.SunderEndHurtRate),     // 破甲后的伤害转换率
		SunderEndDefendRate: int(enemy.SunderEndDefendRate + propertyAdd.SunderEndDefendRate), // 破甲后减防
		SunderAttack:        int(enemy.SunderAttack + propertyAdd.SunderAttack),               // 攻击破甲值
		RoleId:              int(enemy.Id),
		SkillWait:           int(enemy.SkillWait),
		SkillInfos: []*battle.SkillInfo{
			&battle.SkillInfo{
				SkillId:       int(enemy.SkillId),
				SkillForce:    float64(enemy.SkillForce),
				SkillId2:      int(enemy.Skill2Id),
				SkillForce2:   float64(enemy.Skill2Force),
				Rhythm:        int(enemy.RecoverRoundNum + 1 - enemy.CommonAttackNum),
				UseRhythm:     int(enemy.UseRhythm),
				RecoverRhythm: int(enemy.RecoverRoundNum),
			},
		},
	}
}

func NewBattleTotemInfo(db *mdb.Database) (battleTotemInfo [5]*battle.TotemInfo) {
	if playerTotemInfo := db.Lookup.PlayerTotemInfo(db.PlayerId()); playerTotemInfo != nil {
		totemEquipmentPosData := []int64{
			playerTotemInfo.Pos1,
			playerTotemInfo.Pos2,
			playerTotemInfo.Pos3,
			playerTotemInfo.Pos4,
			playerTotemInfo.Pos5,
		}
		for i, tid := range totemEquipmentPosData {
			if tid != totem_dat.TOTEM_POS_EMPTY {
				playerTotem := db.Lookup.PlayerTotem(tid)
				battleTotemInfo[i] = &battle.TotemInfo{
					Id:      playerTotem.TotemId,
					SkillId: playerTotem.SkillId,
					Level:   int16(playerTotem.Level),
				}
			}
		}
	}

	return battleTotemInfo
}

// 创建玩家角色战斗对象
func newFighterForPlayerRole(player_id int64, position int, role *mdb.PlayerRole, autoFight, isMainRole bool) *battle.Fighter {
	roleInfo := role_dat.GetRoleLevelInfo(role.RoleId, role.Level)

	fighter := &battle.Fighter{
		Kind:                battle.FK_BUDDY,                // 是否是主角
		PlayerId:            player_id,                      // 玩家ID
		RoleId:              int(role.RoleId),               // 角色ID
		Level:               int(role.Level),                // 等级
		Prop:                battle.FP_HUMAN,                // 种族
		FriendshipLevel:     int16(role.FriendshipLevel),    // 羁绊等级
		Position:            position,                       // 站位ID
		Health:              int(roleInfo.Health),           // 生命
		MaxHealth:           int(roleInfo.Health),           // 生命值
		Speed:               roleInfo.Speed,                 // 速度
		Attack:              roleInfo.Attack,                // 普通攻击
		Defend:              roleInfo.Defence,               // 普通防御
		Block:               roleInfo.Block,                 // 格挡
		Critial:             roleInfo.Critial,               // 暴击
		CritialHurt:         roleInfo.CritialHurt,           // 必杀
		Dodge:               roleInfo.Dodge,                 // 闪避
		Hit:                 roleInfo.Hit,                   // 命中
		Cultivation:         roleInfo.Cultivation,           // 修为
		Sleep:               float64(roleInfo.Sleep),        // 睡眠抗性
		Dizziness:           float64(roleInfo.Dizziness),    // 眩晕抗性
		Random:              float64(roleInfo.Random),       // 混乱抗性
		DisableSkill:        float64(roleInfo.DisableSkill), // 封魔抗性
		Poisoning:           float64(roleInfo.Poisoning),    // 中毒抗性
		SunderMaxValue:      int(roleInfo.SunderMaxValue),
		SunderMinHurtRate:   int(roleInfo.SunderHurtRate),
		SunderEndHurtRate:   int(roleInfo.SunderEndHurtRate),
		SunderEndDefendRate: int(roleInfo.SunderEndDefendRate), // 破甲后减防
	}

	if isMainRole {
		fighter.Kind = battle.FK_PLAYER
		fighter.AutoFight = autoFight
		fighter.Power = int(roleInfo.InitPower)   // 初始精气
		fighter.MaxPower = int(roleInfo.MaxPower) // 精气上限
	}

	return fighter
}

// 创建灵宠战斗对象
func NewFighterForBattlePet(player_id int64, petId int32, pos int, roundAttackNum, liveRound int8, nowRound int, petLevel int16, petSkillLv int16) *battle.Fighter {
	//如果需要为怪物进行属性加成，请修改NewFighterForEnemy第二个参数
	fighter := NewFighterForEnemy(petId, 0, pos)
	fighter.IsBattlePet = true
	fighter.BattlePetLiveRound = int(liveRound)
	fighter.BattlePetLiveStartRound = nowRound
	fighter.BattlePetPerRoundAttackNum = roundAttackNum

	// 玩家的灵宠
	if player_id > 0 {
		fighter.Kind = battle.FK_BUDDY
		fighter.Level = int(petLevel)
		fighter.PlayerId = player_id
		//playerPetDat := battle_pet_dat.GetBattlePetWithEnemyId(petId)
		//gridAttribute := battle_pet_dat.GetBattlePetGridAttribute(gridId, gridLv)
		petLevelInfo := battle_pet_dat.GetPetLevelInfo(petId, petLevel)
		fighter.NaturalSkillLv = petSkillLv
		fighter.Health = int(petLevelInfo.Health)      //int(playerPetDat.Health) + int(gridAttribute.Health)
		fighter.MaxHealth = int(petLevelInfo.Health)   //int(playerPetDat.Health) + int(gridAttribute.Health)
		fighter.Speed = float64(petLevelInfo.Speed)    //float64(playerPetDat.Speed) + float64(gridAttribute.Speed)
		fighter.Defend = float64(petLevelInfo.Defence) //float64(playerPetDat.Defence) + float64(gridAttribute.Defence)
		fighter.Attack = float64(petLevelInfo.Attack)  //float64(playerPetDat.Attack) + float64(gridAttribute.Attack)
		fighter.SunderMaxValue = int(petLevelInfo.SunderMaxValue)
		fighter.SunderMinHurtRate = int(petLevelInfo.SunderMinHurtRate)
		fighter.SunderEndHurtRate = int(petLevelInfo.SunderEndHurtRate)
		fighter.SunderEndDefendRate = int(petLevelInfo.SunderEndDefendRate)
	}

	return fighter
}

var (
	randomPos = []int{5, 4, 3, 2, 1, 10, 9, 8, 7, 14, 13} //怪物随机布阵从这里随机
	reserve   = []int{6, 12, 15, 11}                      //随机可位置用完了从这里按照依次使用
)

//随机产生怪物阵形
func genEnemyFormPos() (posForm []int) {
	// 随机关卡怪物
	/*
		阵型位置随机
		随机到的怪物允许放在如下3*5位置
		|15|14|13|12|11|
		|10|9 | 8| 7| 6|
		|5 |4 | 3| 2| 1|
		用0~14的数作为索引取位置。
		优先在 randomPos 里面随机
		最后考虑使用 reserve 的位置
	*/
	posNum := rand.Perm(len(randomPos))
	for i := 0; i < len(posNum); i++ {
		posForm = append(posForm, randomPos[posNum[i]])
	}
	posForm = append(posForm, reserve...)
	return posForm
}

//TODO 支持随机出怪和按照布阵出怪两种形式
//加载战场中的一组怪物
func NewEnemyFighterGroup(enemyId int32) []*battle.Fighter {
	levelEnemy := mission_dat.GetMissionLevelEnemyById(enemyId)
	var fighters []*battle.Fighter

	// 随机关卡怪物
	/*
		阵型位置随机
		随机到的怪物允许放在如下3*3位置
		|14|13|12|
		|9 | 8| 7|
		|4 | 3| 2|
		用0~8的数作为索引取位置。
	*/

	fighters = make([]*battle.Fighter, ALL_FIGHTER_POS_NUM)

	var pos, posIdx int
	var chance, randChance, i int8

	posIdx = 0
	posForm := genEnemyFormPos()
	for i = 0; i < levelEnemy.MonsterNum; i++ {
		chance = int8(0)
		randChance = int8(rand.Intn(100) + 1)
		if levelEnemy.Monster1Id > 0 && levelEnemy.Monster1Chance > 0 {
			if randChance >= chance && randChance <= chance+levelEnemy.Monster1Chance {
				pos = posForm[posIdx]
				//如果需要为怪物进行属性加成，请修改NewFighterForEnemy第二个参数
				fighters[pos-1] = NewFighterForEnemy(levelEnemy.Monster1Id, 0, pos)
				posIdx++
				continue
			}
			chance += levelEnemy.Monster1Chance
		}

		if levelEnemy.Monster2Id > 0 && levelEnemy.Monster2Chance > 0 {
			if randChance >= chance && randChance <= chance+levelEnemy.Monster2Chance {
				pos = posForm[posIdx]
				//如果需要为怪物进行属性加成，请修改NewFighterForEnemy第二个参数
				fighters[pos-1] = NewFighterForEnemy(levelEnemy.Monster2Id, 0, pos)
				posIdx++
				continue
			}
			chance += levelEnemy.Monster2Chance
		}

		if levelEnemy.Monster3Id > 0 && levelEnemy.Monster3Chance > 0 {
			if randChance >= chance && randChance <= chance+levelEnemy.Monster3Chance {
				pos = posForm[posIdx]
				//如果需要为怪物进行属性加成，请修改NewFighterForEnemy第二个参数
				fighters[pos-1] = NewFighterForEnemy(levelEnemy.Monster3Id, 0, pos)
				posIdx++
				continue
			}
			chance += levelEnemy.Monster3Chance
		}

		if levelEnemy.Monster4Id > 0 && levelEnemy.Monster4Chance > 0 {
			if randChance >= chance && randChance <= chance+levelEnemy.Monster4Chance {
				pos = posForm[posIdx]
				//如果需要为怪物进行属性加成，请修改NewFighterForEnemy第二个参数
				fighters[pos-1] = NewFighterForEnemy(levelEnemy.Monster4Id, 0, pos)
				posIdx++
				continue
			}
			chance += levelEnemy.Monster4Chance
		}

		if levelEnemy.Monster5Id > 0 && levelEnemy.Monster5Chance > 0 {
			if randChance >= chance && randChance <= chance+levelEnemy.Monster5Chance {
				pos = posForm[posIdx]
				//如果需要为怪物进行属性加成，请修改NewFighterForEnemy第二个参数
				fighters[pos-1] = NewFighterForEnemy(levelEnemy.Monster5Id, 0, pos)
				posIdx++
				continue
			}
			chance += levelEnemy.Monster5Chance
		}
	}
	return fighters
}
