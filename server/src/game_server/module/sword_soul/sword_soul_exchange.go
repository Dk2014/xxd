package sword_soul

import (
	"core/net"
	"game_server/api/protocol/sword_soul_api"
	"game_server/dat/player_dat"
	"game_server/dat/sword_soul_dat"
	"game_server/module"
	"game_server/tlog"
	"game_server/xdlog"
)

func Exchange(session *net.Session, in *sword_soul_api.Exchange_In, out *sword_soul_api.Exchange_Out) {
	swordSoulId := in.SwordSoulId
	out.Id = exchangeByFragment(session, swordSoulId)
}

// 使用碎片兑换
func exchangeByFragment(session *net.Session, swordSoulId int16) (playerSwordSoulId int64) {
	state := module.State(session)
	swordSoul := sword_soul_dat.GetSwordSoul(swordSoulId)

	// 操作背包
	module.Player.DecSwordSoulFragment(state.Database, int64(swordSoul.FragmentNum), player_dat.SWORDSOULFRAGMENT, tlog.MFR_SWORD_SOUL_EXCHANGE, xdlog.ET_EXCHANGE_SWORD_SOUL)
	tlog.PlayerSystemModuleFlowLog(state.Database, tlog.SMT_SWORD_SOUL)

	// 增加剑心
	playerSwordSoulId = addSwordSoul(state, swordSoulId, tlog.IFR_SWORD_SOUL_EXCHANGE)

	return playerSwordSoulId
}
