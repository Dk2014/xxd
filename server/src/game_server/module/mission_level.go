package module

import (
	"bytes"
	"core/log"
	"encoding/gob"
	"game_server/battle"
	"game_server/dat/mission_dat"
	"game_server/dat/role_dat"
	"game_server/dat/skill_dat"
	"game_server/mdb"
)

// 关卡通关宝箱状态
type MissionLevelBoxState struct {
	AwardOpenCount  int8           // 允许开启次数, 初始值 0
	AwardedList     map[int8]int64 // 已获得的奖励 map[品质顺序][宝箱ID]
	AwardedItemType int8           //奖励类型：深渊关卡有奖励保护机制，两次翻牌必定获得不同类型的物品
	OpenedBoxPos    []int8         // 已开启宝箱的序号1~5(从左到右)[0~4][品质顺序]
	// 宝箱奖励
	AwardExp      int32
	AwardCoin     int64
	AwardMultExp  int32
	AwardMultCoin int64
}

// 关卡通关随机奖励宝箱状态
type MissionLevelRandomBoxState struct {
	RandomBoxOpenCount int8 // 允许开启次数, 初始值 0
	//RandomAwardedList  []int64 // 已获得的奖励
	FirstOpen bool
}

type EnemyState struct {
	EnemyId int32
	Health  int
}

const (
	LEVEL_RELIVE_STATE_NORMAL = iota
	LEVEL_RELIVE_STATE_RELIVING
	LEVEL_RELIVE_STATE_CAN_RELIVE
)

// 玩家关卡状态
// 注意：关卡状态会被保存在数据库text字段，允许最大大小2^16=65535bytes
type MissionLevelState struct {
	// 状态变量
	LevelType int8  // 关卡类型
	LevelId   int32 // 当前关卡ID
	EnemyId   int32 // 当前关卡遭遇的怪物ID 初始值为0
	MaxRound  int   // 允许最大回合数

	HaveReduceLevelDialyNum bool // 标记是否已经累计了区域关卡的进入

	TotalRound  int   // 整个关卡的总回合数
	BossScore   int32 // boss战得分
	PassBoss    bool  // 是否打败Boss
	ReliveRound int   //与同一个怪物敌人战斗累计战败复活的回合数，在与在每次DoWin时清空

	PassEnemyIds       []int32        // 当前已通过的怪物组id
	SmallBoxList       map[int32]int8 // 出现的小宝箱列表map[宝箱ID]奖励个数
	MengYaoList        map[int32]int8 //出现的梦妖列表[关卡梦妖ID]效果ID
	ShadowList         map[int32]int  // [影之间隙ID]残存怪物数
	MainMonsterLeft    int            // 主场景残存怪物数
	LastFightingMap    int32          // 最近一次战斗场景
	CatchedBattlePetId int32          // 捕捉到的灵宠ID

	CalledBattlePet map[int32]int8 //使用过的灵宠 battle_pet_id -> whild_card
	UsedGhost       map[int16]int8 //使用技能的魂侍 ghost_id -> whild_card

	// 关卡相关的状态类型
	AwardBox         bool
	HasBeenEvaluated bool                 // 已进入过结算流程
	BoxState         MissionLevelBoxState // 通关宝箱状态
	ReliveState      uint8                // 关卡复活状态(可复活，已复活 ...)

	NextReliveCostIngot int32 // 下一次复活需要的元宝数

	AttackerInfo battle.ExportFighterInfo // 战场攻方信息
	DefendEnemy  map[int]EnemyState       // [pos][enemyId]玩家在战斗中逃跑时保存对战的怪物信息，以便再次相遇

	SkillReleaseNum map[int]int //伙伴技能释放次数 skill_id -> release_num

	RandomBox              MissionLevelRandomBoxState // 随机宝箱状态
	EventPartedStatus      map[string]int16           // 关卡内活动参加状态
	ItemsAwardedByAddQuest map[int16]int16            // 玩家从支线任务获得的物品
}

func NewMissionLevelState(levelType int8, levelId int32) *MissionLevelState {
	return &MissionLevelState{
		LevelType:    levelType,
		LevelId:      levelId,
		PassEnemyIds: []int32{},

		SmallBoxList: make(map[int32]int8),
		MengYaoList:  make(map[int32]int8),
		ShadowList:   make(map[int32]int),

		BoxState: MissionLevelBoxState{
			AwardedList:     make(map[int8]int64),
			OpenedBoxPos:    make([]int8, 5),
			AwardedItemType: mission_dat.LEVEL_BOX_AWARD_NIL,
		},

		ReliveState: LEVEL_RELIVE_STATE_NORMAL,

		AttackerInfo: battle.ExportFighterInfo{
			Buffs:     make(map[int][]battle.Buff),
			Skills:    make(map[int][]battle.SkillInfo),
			Fighters:  make(map[int]battle.FighterAttribute),  // 保存玩家所有角色的一些属性，为战斗提供
			GhostInfo: make(map[int64]battle.PlayerGhostInfo), //玩家魂侍战斗相关信息

		},
		SkillReleaseNum: make(map[int]int),
		UsedGhost:       make(map[int16]int8),
		CalledBattlePet: make(map[int32]int8),

		RandomBox: MissionLevelRandomBoxState{
		//RandomAwardedList: []int64{},
		},

		EventPartedStatus: make(map[string]int16),
	}
}

//初始化伙伴技能使用次数，保存在MissionLevelState中
func (this *MissionLevelState) LoadBuddySkill(state *SessionState) {
	state.Database.Select.PlayerSkill(func(row *mdb.PlayerSkillRow) {
		if role_dat.IsMainRole(row.RoleId()) {
			return
		}
		skillId := row.SkillId()
		skillData := skill_dat.GetSkillInfo(skillId)
		if skillData.ChildKind == skill_dat.SKILL_KIND_ULTIMATE {
			skillContent := skill_dat.GetSkillContent(skillId)
			this.SkillReleaseNum[int(skillId)] = int(skillContent.ReleaseNum)
		}
	})
}

func (this *MissionLevelState) LoadFighterAttribute(state *SessionState) {
	fighters := make([]*battle.Fighter, ALL_FIGHTER_POS_NUM)
	inFormRoleInfos := make([]*InFormRoleInfo, 1, ALL_FIGHTER_POS_NUM)

	// 为玩家角色构造form
	pos := 0
	state.Database.Select.PlayerRole(func(row *mdb.PlayerRoleRow) {
		inFormRoleInfos = GetBattleBiz.setFormRoleInfo(row.RoleId(), POS_TYPE_MAIN_FORM, pos, false, inFormRoleInfos)
		pos += 1
	})

	// 角色数据加成计算
	GetBattleBiz.SetFighters(state.Database, inFormRoleInfos, fighters, false, false, FIGHT_FOR_ALL)

	for _, f := range fighters {
		if f == nil {
			continue
		}
		this.AttackerInfo.Fighters[f.RoleId] = battle.FighterAttribute{
			UsedGhostSkill: f.UsedGhostSkill,
			SunderValue:    f.SunderMaxValue,
			Power:          f.Power,
			Health:         f.MaxHealth,

			MaxHealth:      f.MaxHealth,
			SunderMaxValue: f.SunderMaxValue,
		}
	}

}

func (this *MissionLevelState) HasKilledThoseMustDie() bool {
	return this.MainMonsterLeft <= 0
}

func (this *MissionLevelState) MarkPassEnemy() {
	if this.EnemyId <= 0 {
		return
	}
	this.PassEnemyIds = append(this.PassEnemyIds, this.EnemyId)
	enemy := mission_dat.GetMissionLevelEnemyById(this.EnemyId)
	if enemy.ShadedMissionId == 0 {
		this.MainMonsterLeft--
	} else {
		if count, exist := this.ShadowList[enemy.ShadedMissionId]; exist {
			this.ShadowList[enemy.ShadedMissionId] = count - 1
		}
	}
}

func (this *MissionLevelState) IsPassEnemy(enemyId int32) bool {
	for _, id := range this.PassEnemyIds {
		if id == enemyId {
			return true
		}
	}
	return false
}

// 获取在关卡中角色的buff
func GetBuffInMissionLevelWithRoleId(levelState *MissionLevelState, roleId int) []battle.Buff {
	return levelState.AttackerInfo.Buffs[roleId]
}

func SaveMissionLevelStateForPlayer(state *SessionState) {
	if state.MissionLevelState == nil {
		return
	}

	// 如果玩家在战斗中完全掉线，则不保存关卡信息
	if state.Battle != nil {
		return
	}

	var stateBin bytes.Buffer
	encoder := gob.NewEncoder(&stateBin)
	err := encoder.Encode(state.MissionLevelState)
	if err != nil {
		log.Errorf("player %v encode mission-level-state error: %v", state.PlayerId, err)
		return
	}

	state.Database.Insert.PlayerMissionLevelStateBin(&mdb.PlayerMissionLevelStateBin{
		Pid: state.PlayerId,
		Bin: stateBin.Bytes(),
	})
}

// 恢复玩家上一次离线时所在的关卡状态；返回true表示恢复成功，反之false
func RevertMissionLevelStateForPlayer(state *SessionState) (ret bool) {
	row := state.Database.Lookup.PlayerMissionLevelStateBin(state.PlayerId)
	if row == nil {
		return
	}

	var levelState *MissionLevelState

	stateBin := bytes.NewBuffer(row.Bin)
	decoder := gob.NewDecoder(stateBin)

	err := decoder.Decode(&levelState)
	if err != nil {
		log.Errorf("player %v decode mission-level-state error: %v", state.PlayerId, err)
	} else {
		ret = true
		state.MissionLevelState = levelState
	}

	// 读取成功后删除状态记录
	state.Database.Delete.PlayerMissionLevelStateBin(row)
	return
}
