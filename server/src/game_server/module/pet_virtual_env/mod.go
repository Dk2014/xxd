package pet_virtual_env

import (
	"core/fail"
	"core/net"
	"core/time"
	"game_server/api/protocol/pet_virtual_env_api"
	"game_server/battle"
	"game_server/dat/mission_dat"
	"game_server/dat/player_dat"
	"game_server/dat/quest_dat"
	"game_server/mdb"
	"game_server/module"
)

func init() {
	module.PetVirtualEnv = PetVirtualEnvMod{}
}

type PetVirtualEnvMod struct{}

func (mod PetVirtualEnvMod) BattleWin(session *net.Session, xdEventType int32) {
	state := module.State(session)

	playerPveState := state.Database.Lookup.PlayerPveState(state.PlayerId)
	if playerPveState.MaxPassedFloor < state.PetVirtualEnvState.Floor {
		playerPveState.MaxPassedFloor = state.PetVirtualEnvState.Floor
		playerPveState.UnpassedFloorEnemyNum = 0
		state.Database.Update.PlayerPveState(playerPveState)
	}

	session.Send(&pet_virtual_env_api.PveKills_Out{
		Num: state.PetVirtualEnvState.EnemyNum,
	})
	givePVERegularAward(state, xdEventType)

	state.PetVirtualEnvState = nil
	playerPveState.DailyNum += 1
	state.Database.Update.PlayerPveState(playerPveState)

	//刷新每日任务状态
	module.Quest.RefreshDailyQuest(state, quest_dat.DAILY_QUEST_CLASS_PET_VIRENV)
}

func (mod PetVirtualEnvMod) StartPetVirtualEnvLevel(session *net.Session, enemyId int32) {
	state := module.State(session)
	playerPveState := state.Database.Lookup.PlayerPveState(state.PlayerId)
	if !time.IsInPointHour(player_dat.RESET_PVE_LEVEL_IN_HOUR, playerPveState.EnterTime) {
		playerPveState.DailyNum = 0
	}
	fail.When(playerPveState.DailyNum >= mission_dat.PVE_LEVEL_DAILY_NUM, "灵宠环境达到每天最大进入次数")

	//根据 enemy ID 反查 关卡
	pveDat, missionLevelDat := getPveLevelDatByEnemyId(enemyId)
	//检查是否满足进入条件
	checkPveLevel(state.Database, pveDat)
	//cache关卡信息
	state.PetVirtualEnvState = &module.PetVirtualEnvState{
		Floor:      pveDat.Floor,
		PveLevelId: pveDat.Id,
	}
	//构造 mission level state
	state.MissionLevelState = module.NewMissionLevelState(battle.BT_PVE_LEVEL, 0)
	state.MissionLevelState.MaxRound = mission_dat.PVE_LEVEL_MAX_ROUND
	state.MissionLevelState.LoadBuddySkill(state)
	state.MissionLevelState.LoadFighterAttribute(state)
	state.MissionLevelState.EnemyId = enemyId

	//构造战场
	state.Battle = module.Battle.NewBattleForPetVirtualEnv(session, missionLevelDat.Id)
	//设置灵宠关卡状态
	enemyIds := mission_dat.GetEnemyIdByMissionLevelId(missionLevelDat.Id)
	if len(enemyIds) > mission_dat.PVE_LEVEL_INIT_ENEMY_GROUP_NUM {
		state.PetVirtualEnvState.MaxLoadedIndex = mission_dat.PVE_LEVEL_INIT_ENEMY_GROUP_NUM - 1
	} else {
		state.PetVirtualEnvState.MaxLoadedIndex = len(enemyIds) - 1
	}
	state.PetVirtualEnvState.EnemyIds = enemyIds

	//设置时间
	playerPveState.EnterTime = time.GetNowTime()
	state.Database.Update.PlayerPveState(playerPveState)
}

func (mod PetVirtualEnvMod) BattleLose(session *net.Session, xdEventType int32) {
	state := module.State(session)
	playerPveState := state.Database.Lookup.PlayerPveState(state.PlayerId)
	if playerPveState.UnpassedFloorEnemyNum < state.PetVirtualEnvState.EnemyNum {
		playerPveState.UnpassedFloorEnemyNum = state.PetVirtualEnvState.EnemyNum
		state.Database.Update.PlayerPveState(playerPveState)
	}
	session.Send(&pet_virtual_env_api.PveKills_Out{
		Num: state.PetVirtualEnvState.EnemyNum,
	})
	givePVERegularAward(state, xdEventType)
	state.PetVirtualEnvState = nil
	playerPveState.DailyNum += 1
	state.Database.Update.PlayerPveState(playerPveState)

	//刷新每日任务状态
	module.Quest.RefreshDailyQuest(state, quest_dat.DAILY_QUEST_CLASS_PET_VIRENV)
}

func (mod PetVirtualEnvMod) OpenFuncForPetVirtualEnv(db *mdb.Database) {
	db.Insert.PlayerPveState(&mdb.PlayerPveState{
		Pid: db.PlayerId(),
	})
}
func (mod PetVirtualEnvMod) WarnupNextEnemyGroup(state *module.SessionState, battleState *battle.BattleState) (fighters []*battle.Fighter) {
	currendIndex := battleState.Defenders.CurrentGroupIndex()
	if currendIndex < state.PetVirtualEnvState.MaxLoadedIndex ||
		state.PetVirtualEnvState.MaxLoadedIndex == len(state.PetVirtualEnvState.EnemyIds)-1 {
		return nil
	}
	//战场中的怪物组已是最后一组时，检查是否有更多
	state.PetVirtualEnvState.MaxLoadedIndex += 1
	enemyId := state.PetVirtualEnvState.EnemyIds[state.PetVirtualEnvState.MaxLoadedIndex]
	return module.NewEnemyFighterGroup(enemyId)
}
