package battle

import (
	"time"

	"core/net"
	"game_server/battle"
	"game_server/dat/mission_dat"
	"game_server/dat/role_dat"
	"game_server/dat/skill_dat"
	"game_server/mdb"
	"game_server/module"
)

type BattleMod struct {
}

func init() {
	module.Battle = BattleMod{}
}

// 创建关卡战场
func (this BattleMod) NewBattleForLevel(session *net.Session, battleType int8, enemyId int32, notAward bool) module.IBattle {
	state := module.State(session)

	state.MissionLevelState.EnemyId = enemyId
	state.MissionLevelState.LevelType = battleType

	battleMissionLevel := &Battle_MissionLevel{notAward: notAward}

	var attackSide, defendSide *battle.SideInfo
	attackSide = newAttackerSide(state, battleType)
	defendSide = battleMissionLevel.enemyDefendSide(battleType, enemyId, state.MissionLevelState.DefendEnemy)

	battleMissionLevel.battleState = battle.Start(int(battleType), attackSide, defendSide, nil)
	battleMissionLevel.battleState.RegisterAddMonsterHook(commomSummonHandlerFacotry(state.PlayerId))

	// 玩家是复活的，我们不需要为玩家导入信息，并且需要重新导出玩家信息
	if state.MissionLevelState.ReliveState == module.LEVEL_RELIVE_STATE_RELIVING {
		battleMissionLevel.importGhostInfo(state)
		battleMissionLevel.saveAttackerInfo(state)
	} else {
		battleMissionLevel.importAttackInfo(state)
	}

	// 构建战场成功，说明玩家成功复活，重置复活状态
	state.MissionLevelState.ReliveState = module.LEVEL_RELIVE_STATE_NORMAL
	skillUsage := make(map[int64]map[int]int)
	skillUsage[state.PlayerId] = state.MissionLevelState.SkillReleaseNum

	ghostUsage := make(map[int64]map[int16]int8)
	ghostUsage[state.PlayerId] = state.MissionLevelState.UsedGhost
	rsp := GetStartBattleResponse(battleMissionLevel.battleState, state.MissionLevelState, skillUsage, ghostUsage)
	session.Send(rsp)

	return battleMissionLevel
}

func (this BattleMod) NewBattleForPetVirtualEnv(session *net.Session, missionLvId int32) module.IBattle {
	state := module.State(session)
	//实现IBattle接口
	battleMissionLevel := &Battle_MissionLevel{}
	attackSide := newAttackerSide(state, battle.BT_PVE_LEVEL)

	enemyIds := mission_dat.GetEnemyIdByMissionLevelId(missionLvId)
	defendSide := battleMissionLevel.enemyDefendSide(battle.BT_PVE_LEVEL, enemyIds[0], nil)
	//一次最多加载2组
	for i := 1; i < len(enemyIds) && i < mission_dat.PVE_LEVEL_INIT_ENEMY_GROUP_NUM; i++ {
		defendSide.Groups = append(defendSide.Groups, module.NewEnemyFighterGroup(enemyIds[i]))
	}

	//构造战斗核心
	battleMissionLevel.battleState = battle.Start(battle.BT_PVE_LEVEL, attackSide, defendSide, nil)
	battleMissionLevel.battleState.RegisterAddMonsterHook(commomSummonHandlerFacotry(state.PlayerId))

	skillUsage := make(map[int64]map[int]int)
	skillUsage[state.PlayerId] = make(map[int]int)
	skillUsage[state.PlayerId] = state.MissionLevelState.SkillReleaseNum

	rsp := GetStartBattleResponse(battleMissionLevel.battleState, state.MissionLevelState, skillUsage, nil)
	rsp.TotalGroup = int8(len(enemyIds))
	session.Send(rsp)
	return battleMissionLevel
}

func (this BattleMod) NewBattleForDriving(session *net.Session, missionLvId int32) module.IBattle {
	state := module.State(session)
	//实现IBattle接口
	battleMissionLevel := &Battle_MissionLevel{}
	attackSide := newAttackerSide(state, battle.BT_DRIVING_LEVEL)

	enemyIds := mission_dat.GetEnemyIdByMissionLevelId(missionLvId)
	defendSide := battleMissionLevel.enemyDefendSide(battle.BT_DRIVING_LEVEL, enemyIds[0], nil)
	//一次最多加载2组
	for i := 1; i < len(enemyIds) && i < mission_dat.PVE_LEVEL_INIT_ENEMY_GROUP_NUM; i++ {
		defendSide.Groups = append(defendSide.Groups, module.NewEnemyFighterGroup(enemyIds[i]))
	}

	//构造战斗核心
	battleMissionLevel.battleState = battle.Start(battle.BT_DRIVING_LEVEL, attackSide, defendSide, nil)
	battleMissionLevel.battleState.RegisterAddMonsterHook(commomSummonHandlerFacotry(state.PlayerId))

	skillUsage := make(map[int64]map[int]int)
	skillUsage[state.PlayerId] = make(map[int]int)
	skillUsage[state.PlayerId] = state.MissionLevelState.SkillReleaseNum

	rsp := GetStartBattleResponse(battleMissionLevel.battleState, state.MissionLevelState, skillUsage, nil)
	rsp.TotalGroup = int8(len(enemyIds))
	session.Send(rsp)
	return battleMissionLevel
}

func (this BattleMod) NewBattleForRainbowLevel(session *net.Session, enemyId int32) module.IBattle {
	state := module.State(session)

	//实现IBattle接口
	battleMissionLevel := &Battle_MissionLevel{}

	var attackSide, defendSide *battle.SideInfo
	attackSide = newAttackerSide(state, battle.BT_RAINBOW_LEVEL)
	defendSide = battleMissionLevel.enemyDefendSide(battle.BT_RAINBOW_LEVEL, enemyId, state.MissionLevelState.DefendEnemy)

	//构造战斗核心
	battleMissionLevel.battleState = battle.Start(battle.BT_RAINBOW_LEVEL, attackSide, defendSide, nil)
	battleMissionLevel.battleState.RegisterAddMonsterHook(commomSummonHandlerFacotry(state.PlayerId))

	//从 MissionLevetState 导入玩家数据
	battleMissionLevel.importAttackInfo(state)

	skillUsage := make(map[int64]map[int]int)
	skillUsage[state.PlayerId] = state.MissionLevelState.SkillReleaseNum
	ghostUsage := make(map[int64]map[int16]int8)
	ghostUsage[state.PlayerId] = state.MissionLevelState.UsedGhost

	rsp := GetStartBattleResponse(battleMissionLevel.battleState, state.MissionLevelState, skillUsage, ghostUsage)
	session.Send(rsp)
	return battleMissionLevel
}

func (this BattleMod) NewBattleRoom(battleType int, atkSide, defSide *battle.SideInfo, playerChannel *net.Channel, onBattleRoom module.IBattleRoom, roundTime time.Duration) {
	NewBattleRoom(battleType, onBattleRoom, atkSide, defSide, nil, playerChannel, roundTime, false)
}

// func (this BattleMod) NewBattleForEscortBoat(session *net.Session, battleType int, attackerSide, defendSide *battle.SideInfo, boatId int64) module.IBattle {
// 	pve := &Battle_PVE{notAward: false, forbidOP: false, boatId: boatId}
// 	pve.battleState = battle.Start(battleType, attackerSide, defendSide, nil)

// 	state := module.State(session)
// 	skillUsage := make(map[int64]map[int]int)
// 	skillUsage[state.PlayerId] = make(map[int]int)
// 	for _, f := range attackerSide.Fighters {
// 		if f != nil && !role_dat.IsMainRole(int8(f.RoleId)) {
// 			for _, skill := range f.Skills {
// 				if skill != nil {
// 					skillData := skill_dat.GetSkillInfo(int16(skill.SkillId))
// 					if skillData.ChildKind == skill_dat.SKILL_KIND_ULTIMATE {
// 						skillContent := skill_dat.GetSkillContent(int16(skill.SkillId))
// 						skillUsage[state.PlayerId][skill.SkillId] = int(skillContent.ReleaseNum)
// 					}
// 				}
// 			}
// 		}
// 	}
// 	rsp := GetStartBattleResponse(pve.battleState, nil, skillUsage, nil)
// 	session.Send(rsp)

// 	return pve
// }

func (this BattleMod) NewBattleForPVE(session *net.Session, battleType int, attackerSide, defendSide *battle.SideInfo, forbidOP, notAward bool) module.IBattle {
	state := module.State(session)
	pve := &Battle_PVE{notAward: notAward}

	pve.battleState = battle.Start(battleType, attackerSide, defendSide, nil)
	pve.battleState.RegisterAddMonsterHook(commomSummonHandlerFacotry(state.PlayerId))
	skillUsage := make(map[int64]map[int]int)
	skillUsage[state.PlayerId] = make(map[int]int)
	for _, f := range attackerSide.Fighters {
		if f != nil && !role_dat.IsMainRole(int8(f.RoleId)) {
			for _, skill := range f.Skills {
				if skill != nil {
					skillData := skill_dat.GetSkillInfo(int16(skill.SkillId))
					if skillData.ChildKind == skill_dat.SKILL_KIND_ULTIMATE {
						skillContent := skill_dat.GetSkillContent(int16(skill.SkillId))
						skillUsage[state.PlayerId][skill.SkillId] = int(skillContent.ReleaseNum)
					}
				}
			}
		}
	}
	rsp := GetStartBattleResponse(pve.battleState, nil, skillUsage, nil)
	session.Send(rsp)

	return pve
}

//云海拜访
func (this BattleMod) NewBattleForVisiting(session *net.Session, attackSide, defendSide *battle.SideInfo) module.IBattle {
	state := module.State(session)
	//实现IBattle接口
	battleMissionLevel := &Battle_MissionLevel{}

	//构造战斗核心
	battleMissionLevel.battleState = battle.Start(battle.BT_DRIVING_SWORD_BF_LEVEL, attackSide, defendSide, nil)

	skillUsage := make(map[int64]map[int]int)
	skillUsage[state.PlayerId] = make(map[int]int)
	skillUsage[state.PlayerId] = state.MissionLevelState.SkillReleaseNum

	rsp := GetStartBattleResponse(battleMissionLevel.battleState, state.MissionLevelState, skillUsage, nil)
	session.Send(rsp)
	return battleMissionLevel
}

func (this BattleMod) NewBattleForHiJackBoat(session *net.Session, battleType int, attackSide, defendSide *battle.SideInfo, boatId int64) module.IBattle {
	state := module.State(session)
	//实现IBattle接口
	battleMissionLevel := &Battle_MissionLevel{}
	battleMissionLevel.boatId = boatId

	//构造战斗核心
	battleMissionLevel.battleState = battle.Start(battle.BT_HIJACK_BOAT, attackSide, defendSide, nil)

	skillUsage := make(map[int64]map[int]int)
	skillUsage[state.PlayerId] = make(map[int]int)
	skillUsage[state.PlayerId] = state.MissionLevelState.SkillReleaseNum

	rsp := GetStartBattleResponse(battleMissionLevel.battleState, state.MissionLevelState, skillUsage, nil)
	session.Send(rsp)
	return battleMissionLevel
}

func (this BattleMod) NewBattleForRecoverBoat(session *net.Session, battleType int, attackSide, defendSide *battle.SideInfo, boatId int64) module.IBattle {
	state := module.State(session)
	//实现IBattle接口
	battleMissionLevel := &Battle_MissionLevel{}
	battleMissionLevel.boatId = boatId

	//构造战斗核心
	battleMissionLevel.battleState = battle.Start(battle.BT_RECOVER_BOAT, attackSide, defendSide, nil)

	skillUsage := make(map[int64]map[int]int)
	skillUsage[state.PlayerId] = make(map[int]int)
	skillUsage[state.PlayerId] = state.MissionLevelState.SkillReleaseNum

	rsp := GetStartBattleResponse(battleMissionLevel.battleState, state.MissionLevelState, skillUsage, nil)
	session.Send(rsp)
	return battleMissionLevel
}

func newAttackerSide(state *module.SessionState, battleType int8) (attackSide *battle.SideInfo) {
	switch battleType {
	case battle.BT_BUDDY_LEVEL:
		attackSide, _ = module.NewBattleSideWithPlayerDatabaseAndPlayerForm(state.Database, getBuddyLevelForm(state), false, false, false)

	case battle.BT_PET_LEVEL:
		state.MissionLevelState.MaxRound = mission_dat.EXTEND_LEVEL_PET_MAX_ROUND_NUM
		attackSide, _ = module.NewBattleSideWithPlayerDatabaseAndPlayerForm(state.Database, getPetLevelForm(state), false, false, false)

	default:
		attackSide, _ = module.NewBattleSideWithPlayerDatabase(state.Database, false, false, false)
	}
	return
}

func getBuddyLevelForm(state *module.SessionState) *mdb.PlayerFormation {
	playerExtendLevel := state.Database.Lookup.PlayerExtendLevel(state.PlayerId)

	if playerExtendLevel.RolePos == 0 && playerExtendLevel.BuddyPos == 0 {
		playerExtendLevel.RolePos = 2
		playerExtendLevel.BuddyPos = 5
	}

	form := &mdb.PlayerFormation{
		Pid:  state.PlayerId,
		Pos0: module.NONE_ROLE_ID,
		Pos1: module.NONE_ROLE_ID,
		Pos2: module.NONE_ROLE_ID,
		Pos3: module.NONE_ROLE_ID,
		Pos4: module.NONE_ROLE_ID,
		Pos5: module.NONE_ROLE_ID,
		Pos6: module.NONE_ROLE_ID,
		Pos7: module.NONE_ROLE_ID,
		Pos8: module.NONE_ROLE_ID,
	}

	setPos := func(pos, roleId int8) {
		switch pos {
		case 1:
			form.Pos0 = roleId
		case 2:
			form.Pos1 = roleId
		case 3:
			form.Pos2 = roleId
		case 4:
			form.Pos3 = roleId
		case 5:
			form.Pos4 = roleId
		case 6:
			form.Pos5 = roleId
		case 7:
			form.Pos6 = roleId
		case 8:
			form.Pos7 = roleId
		case 9:
			form.Pos8 = roleId
		}
	}

	setPos(playerExtendLevel.RolePos, state.RoleId)
	setPos(playerExtendLevel.BuddyPos, playerExtendLevel.RandBuddyRoleId)
	return form
}

func getPetLevelForm(state *module.SessionState) *mdb.PlayerFormation {
	form := &mdb.PlayerFormation{
		Pid:  state.PlayerId,
		Pos0: module.NONE_ROLE_ID,
		Pos1: module.NONE_ROLE_ID,
		Pos2: module.NONE_ROLE_ID,
		Pos3: module.NONE_ROLE_ID,
		Pos4: state.RoleId,
		Pos5: module.NONE_ROLE_ID,
		Pos6: module.NONE_ROLE_ID,
		Pos7: module.NONE_ROLE_ID,
		Pos8: module.NONE_ROLE_ID,
	}

	return form
}
