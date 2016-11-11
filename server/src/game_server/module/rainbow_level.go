package module

import (
	"bytes"
	"core/log"
	"encoding/gob"
	"game_server/battle"
	"game_server/dat/role_dat"
	"game_server/dat/skill_dat"
	"game_server/mdb"
)

type RainbowLevelState struct {
	AwardBoxIndex   []int8                   //宝箱奖励索引
	CalledBattlePet map[int32]int8           //使用过的灵宠 battle_pet_id -> whild_card
	UsedGhost       map[int16]int8           //使用技能的魂侍 ghost_id -> whild_card
	SkillReleaseNum map[int]int              //伙伴技能释放次数 skill_id -> release_num
	AttackerInfo    battle.ExportFighterInfo // 战场攻方信息

	AwardExp          int32
	AwardCoin         int64
	AwardRelationship int32
}

func NewMissionLevelStateForRainbowLevel() *MissionLevelState {
	return &MissionLevelState{
		ReliveState: LEVEL_RELIVE_STATE_NORMAL,

		AttackerInfo: battle.ExportFighterInfo{
			Buffs:     make(map[int][]battle.Buff),
			Skills:    make(map[int][]battle.SkillInfo),
			Fighters:  make(map[int]battle.FighterAttribute), // 保存玩家所有角色的一些属性，为战斗提供
			GhostInfo: make(map[int64]battle.PlayerGhostInfo),
		},
		SkillReleaseNum: make(map[int]int),
		UsedGhost:       make(map[int16]int8),
		CalledBattlePet: make(map[int32]int8),
	}
}

func NewRainbowLevelState() *RainbowLevelState {
	return &RainbowLevelState{
		AttackerInfo: battle.ExportFighterInfo{
			Buffs:     make(map[int][]battle.Buff),
			Skills:    make(map[int][]battle.SkillInfo),
			Fighters:  make(map[int]battle.FighterAttribute),
			GhostInfo: make(map[int64]battle.PlayerGhostInfo),
		},
		SkillReleaseNum: make(map[int]int),
		UsedGhost:       make(map[int16]int8),
		CalledBattlePet: make(map[int32]int8),
	}
}

//初始化伙伴技能使用次数，保存在RainbowLevelState中
func (this *RainbowLevelState) LoadBuddySkill(state *SessionState) {
	state.Database.Select.PlayerSkill(func(row *mdb.PlayerSkillRow) {
		if role_dat.IsMainRole(row.RoleId()) {
			return
		}
		skillId := row.SkillId()
		skillData := skill_dat.GetSkillInfo(skillId)
		if skillData.ChildKind == skill_dat.SKILL_KIND_ULTIMATE {
			skillContent := skill_dat.GetSkillContent(skillId)
			if _, ok := this.SkillReleaseNum[int(skillId)]; !ok {
				//没有记录的新技能
				this.SkillReleaseNum[int(skillId)] = int(skillContent.ReleaseNum)
			}
		}
	})
}

func (this *RainbowLevelState) LoadFighterAttribute(state *SessionState) {
	fighters := make([]*battle.Fighter, ALL_FIGHTER_POS_NUM)
	inFormRoleInfos := make([]*InFormRoleInfo, 1, ALL_FIGHTER_POS_NUM)

	// 为玩家角色构造form
	pos := 0
	state.Database.Select.PlayerRole(func(row *mdb.PlayerRoleRow) {
		inFormRoleInfos = GetBattleBiz.setFormRoleInfo(row.RoleId(), POS_TYPE_MAIN_FORM, pos, false, inFormRoleInfos)
		pos += 1
	})

	// 角色数据加成计算（fighters里面存放的是最新的数据）
	GetBattleBiz.SetFighters(state.Database, inFormRoleInfos, fighters, false, false, FIGHT_FOR_ALL)

	for _, f := range fighters {
		if f == nil {
			continue
		}
		if _, ok := this.AttackerInfo.Fighters[f.RoleId]; !ok {
			//新伙伴使用最新的数据
			this.AttackerInfo.Fighters[f.RoleId] = battle.FighterAttribute{
				SunderValue: f.SunderMaxValue,
				Power:       f.Power,
				Health:      f.MaxHealth,
			}
		} else {
			//精气 魂力 恢复数据库保存的值，
			//生命、护甲最大值如果比之前多，则在保存的值基础上加上增多的值
			//生命、护甲最大值如果比之前少，则使用保存的值，但要不能超过当前上限
			fighterInfo := this.AttackerInfo.Fighters[f.RoleId]
			maxHealthDelt := f.MaxHealth - this.AttackerInfo.Fighters[f.RoleId].MaxHealth
			if maxHealthDelt > 0 {
				fighterInfo.Health += maxHealthDelt
			}
			if this.AttackerInfo.Fighters[f.RoleId].Health > f.MaxHealth {
				fighterInfo.Health = f.MaxHealth
			}

			maxSunderDelt := f.SunderMaxValue - this.AttackerInfo.Fighters[f.RoleId].SunderMaxValue
			if maxSunderDelt > 0 {
				fighterInfo.SunderValue += maxSunderDelt
			}
			if this.AttackerInfo.Fighters[f.RoleId].SunderValue > f.SunderMaxValue {
				fighterInfo.SunderValue = f.SunderMaxValue
			}
			this.AttackerInfo.Fighters[f.RoleId] = fighterInfo
		}
	}
}

//重置彩虹关卡状态
func (this *RainbowLevelState) Reset(state *SessionState) {
	this.CalledBattlePet = make(map[int32]int8)
	this.UsedGhost = make(map[int16]int8)
	//this.UsedGhost = this.UsedGhost[:]
	this.SkillReleaseNum = make(map[int]int, len(this.SkillReleaseNum))
	this.AttackerInfo = battle.ExportFighterInfo{
		Buffs:    make(map[int][]battle.Buff),
		Skills:   make(map[int][]battle.SkillInfo),
		Fighters: make(map[int]battle.FighterAttribute),
	}

	this.LoadBuddySkill(state)
	this.LoadFighterAttribute(state)
}

func (this *RainbowLevelState) SyncMissionLevelState(state *SessionState) {
	//同步魂侍、灵宠信息：灵宠和魂侍使用情况在MissionLevelState和RainbowLevelState公用同一个Map

	//同步战场导出信息
	this.AttackerInfo = state.MissionLevelState.AttackerInfo

	//同步伙伴进阶技能
	for skillId, restReleaseNum := range state.MissionLevelState.SkillReleaseNum {
		this.SkillReleaseNum[skillId] = restReleaseNum
	}
}

//更新彩虹关卡状态
func (this *RainbowLevelState) Update(state *SessionState) {
	this.LoadBuddySkill(state)
	this.LoadFighterAttribute(state)
}

func (state *SessionState) SaveRainbowLevelState() {

	var stateBin bytes.Buffer
	encoder := gob.NewEncoder(&stateBin)
	err := encoder.Encode(state.RainbowLevelState)
	if err != nil {
		log.Errorf("player %v encode rainbow_level_state error: %v", state.PlayerId, err)
		return
	}

	rainbowStateBin := state.Database.Lookup.PlayerRainbowLevelStateBin(state.PlayerId)
	rainbowStateBin.Bin = stateBin.Bytes()

	state.Database.Update.PlayerRainbowLevelStateBin(rainbowStateBin)
}

// 恢复玩家上一次离线时所在的关卡状态；返回true表示恢复成功，反之false
func (state *SessionState) RevertRainbowLevelState() {
	row := state.Database.Lookup.PlayerRainbowLevelStateBin(state.PlayerId)
	if len(row.Bin) == 0 {
		//提前return避免使用 gob.NewDecoder
		state.RainbowLevelState = NewRainbowLevelState()
		return
	}

	var rainbowState *RainbowLevelState
	stateBin := bytes.NewBuffer(row.Bin)
	decoder := gob.NewDecoder(stateBin)

	err := decoder.Decode(&rainbowState)
	if err != nil {
		//反序列化失败属于内部错误，此时把玩家的彩虹关卡状态战斗信息重置
		log.Errorf("player %v decode rainbow_level_state error: %v", state.PlayerId, err)
		state.RainbowLevelState = NewRainbowLevelState()
	} else {
		state.RainbowLevelState = rainbowState
	}
}
