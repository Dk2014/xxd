package driving_sword

import (
	"core/fail"
	"core/net"
	"game_server/api/protocol/driving_sword_api"
	"game_server/dat/driving_sword_dat"
	"game_server/dat/player_dat"
	"game_server/module"
)

func openMountainTreasure(session *net.Session, xdEventType int32) {

	state := module.State(session)
	drivingSwordInfo := state.Database.Lookup.PlayerDrivingSwordInfo(state.Database.PlayerId())
	fail.When(drivingSwordInfo == nil, "no driving sword information")
	playerTreasure := lookupPlayerDrivingTreasureEvent(state.Database, drivingSwordInfo.CurrentCloud, uint8(drivingSwordInfo.CurrentX), uint8(drivingSwordInfo.CurrentY))

	fail.When(playerTreasure == nil, "openMountainTreasure  player event type err")
	playerTreasure.Progress += 1
	item, num, awardCoins := driving_sword_dat.GetTreasureOrderAward(drivingSwordInfo.CurrentCloud, playerTreasure.DataId, playerTreasure.Progress)
	fail.When(item <= 0 && awardCoins <= 0, "openMountainTreasure  player award null")

	if item > 0 && num > 0 {
		module.Item.AddItem(state.Database, item, int16(num), 0, xdEventType, "")
	}
	if awardCoins > 0 {
		module.Player.IncMoney(state.Database, state.MoneyState, int64(awardCoins), player_dat.COINS, -1 /*TODO tlog*/, xdEventType, "")
	}

	state.Database.Update.PlayerDrivingSwordEventTreasure(playerTreasure)
	session.Send(&driving_sword_api.MountainTreasureOpen_Out{})
}
