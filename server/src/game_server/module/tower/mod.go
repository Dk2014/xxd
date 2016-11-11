package tower

import (
	"core/fail"
	"core/net"
	"core/time"
	"game_server/battle"
	"game_server/dat/mission_dat"
	//"game_server/dat/player_dat"
	"game_server/dat/tower_level_dat"
	"game_server/mdb"
	"game_server/module"
	goTime "time"
)

const (
	BATTLE_STATE_NONE   = 0
	BATTLE_STATE_LOSE   = 1
	BATTLE_STATE_PASSED = 2 // 通过全部楼层
)

func init() {
	module.Tower = TowerMod{}
}

type TowerMod struct {
}

//废弃
func (mod TowerMod) StartTownBattle(session *net.Session, battleId int32) {
	module.Town.LeaveTown(session)

	state := module.State(session)
	//module.Player.MustOpenLevelFunc(state.Database, player_dat.LEVEL_FUNC_JI_XIAN_GUAN_KA)

	levelEnemy := mission_dat.GetMissionLevelEnemyById(battleId)
	levelInfo := mission_dat.GetMissionLevelById(levelEnemy.MissionLevelId)

	fail.When(levelInfo.ParentType != battle.BT_TOWER_LEVEL, "incorrect tower level id")

	playerTower := state.Database.Lookup.PlayerTowerLevel(state.PlayerId)
	fail.When(!time.IsToday(playerTower.OpenTime), "incorrect tower battle request")

	towerId := tower_level_dat.GetTowerIdByFloor(playerTower.Floor)

	fail.When(playerTower.Floor >= tower_level_dat.MAX_FLOOR_NUM, "tower is passed")
	fail.When(towerId != levelInfo.ParentId, "incorrect tower id")

	state.MissionLevelState = module.NewMissionLevelState()
	state.MissionLevelState.LevelId = levelInfo.Id
	state.MissionLevelState.AwardBox = levelInfo.AwardBox
	state.MissionLevelState.MaxRound = tower_level_dat.MAX_BATTLE_ROUND

	module.Mission.StartLevelBattle(battle.BT_TOWER_LEVEL, battleId, session)
}

func (mod TowerMod) BattleWin(state *module.SessionState) {
	playerTower := state.Database.Lookup.PlayerTowerLevel(state.PlayerId)

	playerTower.Floor += 1
	playerTower.OpenTime = goTime.Now().Unix()

	if playerTower.Floor >= tower_level_dat.MAX_FLOOR_NUM {
		playerTower.BattleState = BATTLE_STATE_PASSED
	} else {
		playerTower.BattleState = BATTLE_STATE_NONE
	}

	state.Database.Update.PlayerTowerLevel(playerTower)
	module.Mission.LeaveMissionLevel(state)
}

func (mod TowerMod) BattleLose(state *module.SessionState) {
	playerTower := state.Database.Lookup.PlayerTowerLevel(state.PlayerId)
	playerTower.BattleState = BATTLE_STATE_LOSE
	state.Database.Update.PlayerTowerLevel(playerTower)
	module.Mission.LeaveMissionLevel(state)
}

func (mod TowerMod) OpenFunc(db *mdb.Database) {
	db.Insert.PlayerTowerLevel(&mdb.PlayerTowerLevel{
		Pid:         db.PlayerId(),
		Floor:       1,
		BattleState: BATTLE_STATE_NONE,
		OpenTime:    goTime.Now().Unix(),
	})
}
