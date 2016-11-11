package driving_sword

import (
	"core/fail"
	"core/net"
	"core/time"
	"game_server/api/protocol/driving_sword_api"
	"game_server/dat/driving_sword_dat"
	"game_server/dat/player_dat"
	"game_server/dat/vip_dat"
	"game_server/module"
	"game_server/xdlog"
)

func buyActionTimes(session *net.Session) {
	state := module.State(session)
	drivingSwordInfo := state.Database.Lookup.PlayerDrivingSwordInfo(state.Database.PlayerId())
	fail.When(drivingSwordInfo == nil, "no driving sword information")
	if !time.IsInPointHour(player_dat.RESET_DRIVING_SWORD_ACTION_TIMES_IN_HOUR, drivingSwordInfo.ActionRefreshTime) {
		drivingSwordInfo.DailyActionBought = 0
	}

	playerVIP := module.VIP.VIPInfo(state.Database)
	times := vip_dat.GetVIPPrivilegeTime(playerVIP.Level, vip_dat.YUNHAITEQUAN)
	fail.When(times == 0 || times <= int16(drivingSwordInfo.DailyActionBought), "can not buy times")
	costIngot := driving_sword_dat.GetCost(int32(drivingSwordInfo.DailyActionBought))
	module.Player.DecMoney(state.Database, state.MoneyState, int64(costIngot), player_dat.INGOT, 0, xdlog.ET_DRIVING_ACTION_TIMES)
	drivingSwordInfo.DailyActionBought += 1
	drivingSwordInfo.AllowedAction += driving_sword_dat.DAILY_ADDITIONAL_ACTION_COUNT
	drivingSwordInfo.ActionBuyTime = time.GetNowTime()
	state.Database.Update.PlayerDrivingSwordInfo(drivingSwordInfo)

	session.Send(&driving_sword_api.BuyAllowedAction_Out{})
}
