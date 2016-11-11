/*
玩家相关通讯协议
*/
package player

import (
	"core/net"
	"core/time"
	"game_server/api/protocol/player_api"
	"game_server/dat/player_dat"
	"game_server/dat/vip_dat"
	"game_server/module"
	"game_server/module/player_rpc"
	"game_server/tlog"
	"game_server/xdlog"
)

func init() {
	player_api.SetInHandler(PlayerAPI{})
}

type PlayerAPI struct {
}

func (api PlayerAPI) Info(session *net.Session, in *player_api.Info_In) {
	info := &player_api.Info_Out{}
	if GetPlayerInfo(session, info) {
		session.Send(info)
	}
}

func (api PlayerAPI) BuyPhysical(session *net.Session, in *player_api.BuyPhysical_In) {
	rsp := &player_api.BuyPhysical_Out{}
	rsp.PhysicalBuyCount, _ = module.Physical.BuyPhysical(session)
	session.Send(rsp)
}

func (api PlayerAPI) BuyCoins(session *net.Session, in *player_api.BuyCoins_In) {
	state := module.State(session)
	playerCoinsInfo := state.Database.Lookup.PlayerCoins(state.PlayerId)
	//批量购买是一种VIP特权
	if in.Number > 1 {
		module.VIP.CheckPrivilege(state, vip_dat.PILIANGDUIHUAN)
	}

	playerCoinsInfo = updateCoinsBuyCount(state, playerCoinsInfo)

	var ingotDec, coinsInc int64
	out := new(player_api.BuyCoins_Out)
	for i := int16(0); i < in.Number; i++ {
		ingot, coins, crit := buyCoins(state, playerCoinsInfo)
		out.Result = append(out.Result, player_api.BuyCoins_Out_Result{
			Ingot: ingot,
			Coins: coins,
			Crit:  crit,
		})
		ingotDec += ingot
		coinsInc += coins
	}
	module.Player.DecMoney(state.Database, state.MoneyState, ingotDec, player_dat.INGOT, tlog.MFR_BUY_COINS, xdlog.ET_BUY_CONINS)
	module.Player.IncMoney(state.Database, state.MoneyState, coinsInc, player_dat.COINS, tlog.MFR_BUY_COINS, xdlog.ET_BUY_CONINS, "")
	if in.Number > 1 {
		playerCoinsInfo.BatchBought += 1
	}

	state.Database.Update.PlayerCoins(playerCoinsInfo)

	session.Send(out)
}

func (api PlayerAPI) Relogin(session *net.Session, in *player_api.Relogin_In) {
}

func (api PlayerAPI) SetPlayKey(session *net.Session, in *player_api.SetPlayKey_In) {
	setFuncPlaykey(module.State(session).Database, in.Lock)
}

func (api PlayerAPI) GetTime(session *net.Session, in *player_api.GetTime_In) {
	session.Send(&player_api.GetTime_Out{time.GetNowTime()})
}

func (api PlayerAPI) FromPlatformLogin(session *net.Session, in *player_api.FromPlatformLogin_In) {
	out := &player_api.FromPlatformLogin_Out{}
	out.PlayerId = FromLoginPlayer(session, in.Username, in.Nickname, in.Hashcode, int8(in.PlatformType), in.RoleId, in.Unixtime, in.Restore, in, out)
	session.Send(out)
}

func (api PlayerAPI) GlobalLogin(session *net.Session, in *player_api.GlobalLogin_In) {
	rsp := &player_api.GlobalLogin_Out{}
	rsp.Result, rsp.Paytoken = loginGlobal(session, in.Pid, in.Time, in.Openid, in.Nick, in.Hash, in.RoleId, in.RoleLevel, in.FightNum, in)
	session.Send(rsp)
}

func (api PlayerAPI) CrossLoginGameServer(session *net.Session, in *player_api.CrossLoginGameServer_In) {
	session.Send(&player_api.CrossLoginGameServer_Out{
		Result: CrossLogin(session, in.Pid, in.Time, in.Openid, in.Nick, in.Hash, in.RoleId, in.RoleLevel, in.WegamesPlatformUid),
	})
}

func (api PlayerAPI) GetLoginInfo(session *net.Session, in *player_api.GetLoginInfo_In) {
	rsp := &player_api.GetLoginInfo_Out{}
	rsp.Hash, rsp.Time = getLoginInfo(module.State(session))
	session.Send(rsp)
}

func (api PlayerAPI) GetIngot(session *net.Session, in *player_api.GetIngot_In) {
	state := module.State(session)
	out := &player_api.GetIngot_Out{}
	out.Ingot = module.Player.GetIngot(state.Database, state.MoneyState)
	session.Send(out)
}

//查询各系统累计获得声望
func (api PlayerAPI) SystemFame(session *net.Session, in *player_api.SystemFame_In) {
	panic("废弃")
	/*
		session.Send(&player_api.SystemFame_Out{
			Fame:   getSystemFame(state.Database, in.System),
			System: in.System,
		})
	*/
}

func (api PlayerAPI) GetRanks(session *net.Session, in *player_api.GetRanks_In) {
	if in.Flag == player_api.RANK_TYPE_INGOT {
		player_rpc.GetRanksForIngotRank(session, in)
	} else {
		player_rpc.GetRanks(session, in)
	}
}
