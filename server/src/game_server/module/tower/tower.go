package tower

import (
	"core/fail"
	"core/time"
	"game_server/api/protocol/tower_api"
	"game_server/battle"
	_ "game_server/dat/friend_dat"
	"game_server/dat/mission_dat"
	"game_server/dat/player_dat"
	_ "game_server/dat/role_dat"
	"game_server/dat/tower_level_dat"
	_ "game_server/mdb"
	"game_server/module"
	"game_server/tlog"
	"math/rand"
)

func getPlayerTowerInfo(state *module.SessionState) int16 {
	playerTower := state.Database.Lookup.PlayerTowerLevel(state.PlayerId)

	// 隔天重置楼层
	if !time.IsToday(playerTower.OpenTime) {
		playerTower.Floor = 1
		playerTower.OpenTime = time.GetNowTime()
		state.Database.Update.PlayerTowerLevel(playerTower)
	}

	return playerTower.Floor
}

func useLadder(state *module.SessionState) int16 {
	fail.When(!module.Player.CheckMoney(state, tower_level_dat.USE_LADDER_COINS_PRICE, player_dat.COINS), "not enough coins use ladder")

	randNum := int16(rand.Intn(9) + 1)
	playerTower := state.Database.Lookup.PlayerTowerLevel(state.PlayerId)
	fail.When(!time.IsToday(playerTower.OpenTime), "incorrect useLadder")

	startFloor := playerTower.Floor
	playerTower.Floor += randNum
	playerTower.OpenTime = time.GetNowTime()

	// 超过或等于最高层，保留最高层，让玩家自己通过
	if playerTower.Floor >= tower_level_dat.MAX_FLOOR_NUM {
		playerTower.Floor = tower_level_dat.MAX_FLOOR_NUM - 1
	}

	var levelId int32
	var id int16
	playerLevel := state.Database.Lookup.PlayerMissionLevel(state.PlayerId)

	for floor := startFloor; floor < playerTower.Floor; floor++ {
		id = tower_level_dat.GetTowerIdByFloor(floor)
		levelId = mission_dat.GetSpecialLevelId(battle.BT_TOWER_LEVEL, id)
		module.Mission.ServerDirectAward(state, playerLevel, levelId, mission_dat.LEVEL_BOX_AWARD_COUNT)
	}

	module.Player.DecMoney(state.Database, state.MoneyState, tower_level_dat.USE_LADDER_COINS_PRICE, player_dat.COINS, tlog.MFR_USE_LADDER)

	state.Database.Update.PlayerTowerLevel(playerTower)

	return playerTower.Floor
}

func getPlayerFriends(state *module.SessionState) []tower_api.GetInfo_Out_Friends {
	friends := []tower_api.GetInfo_Out_Friends{}

	// var roleId int8
	// var roleLevel, floor int16
	// var playerTower *mdb.PlayerTowerLevel

	// state.Database.Select.PlayerFriend(func(row *mdb.PlayerFriendRow) {
	// 	// 超上限就不再取好友数据
	// 	if len(friends) > tower_level_dat.MAX_RANK_TOTAL {
	// 		row.Break()
	// 	}

	// 	if row.FriendMode() != friend_dat.FRIEND_MODE_FRIEND {
	// 		return
	// 	}

	// 	mdb.AgentExecute(row.FriendPid(), func(db *mdb.Database) {
	// 		// 找好友楼层
	// 		playerTower = db.Lookup.PlayerTowerLevel(db.PlayerId())

	// 		// 找好友主角信息
	// 		db.Select.PlayerRole(func(row2 *mdb.PlayerRoleRow) {
	// 			if role_dat.IsMainRole(row2.RoleId()) {
	// 				roleId = row2.RoleId()
	// 				roleLevel = row2.Level()
	// 				row2.Break()
	// 			}
	// 		})
	// 	})

	// 	// 好友可能没有开启通天塔
	// 	if playerTower == nil {
	// 		floor = 0
	// 	} else {
	// 		floor = playerTower.Floor
	// 		// 重置楼层
	// 		if !time.IsToday(playerTower.OpenTime) {
	// 			floor = 1
	// 		}
	// 	}

	// 	friends = append(friends, tower_api.GetInfo_Out_Friends{
	// 		RoleId:   roleId,
	// 		Nickname: []byte(row.FriendNick()),
	// 		Level:    roleLevel,
	// 		Floor:    floor,
	// 	})
	// })

	return friends
}
