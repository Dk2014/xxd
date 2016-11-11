package trader

import (
	//"game_server/dat/battle_pet_dat"
	"game_server/dat/ghost_dat"
	"game_server/dat/mail_dat"
	"game_server/dat/player_dat"
	"game_server/dat/sword_soul_dat"
	"game_server/dat/trader_dat"
	"game_server/module"
	"game_server/module/rpc"
	"game_server/tlog"
)

func decMoney(state *module.SessionState, moneyType int8, cost int64, xdEventType int32) {
	switch moneyType {
	case trader_dat.COINS_TYPE:
		module.Player.DecMoney(state.Database, state.MoneyState, cost, player_dat.COINS, tlog.MFR_TRADER_BUY, xdEventType)
	case trader_dat.INGOT_TYPE:
		module.Player.DecMoney(state.Database, state.MoneyState, cost, player_dat.INGOT, tlog.MFR_TRADER_BUY, xdEventType)
	case trader_dat.HEART_TYPE:
		module.Heart.DecHeart(state.Database, int16(cost))
	case trader_dat.DRAGON_COIN_TYPE:
		//消耗龙币。
		module.Item.DelItemByItemId(state.Database, 270, int16(cost), tlog.IFR_TRADER_BUY, xdEventType)
	default:
		panic("undefine money type")
	}
}

func addGoods(state *module.SessionState, goodsType int8, itemId int16, num int16, source string, xdEventType int32) {
	switch goodsType {
	case trader_dat.ITEM:
		module.Item.AddItem(state.Database, itemId, num, tlog.IFR_TRADER_BUY, xdEventType, "")
	case trader_dat.SWORD_SOUL:
		addSwordSoul(state, itemId, source, tlog.IFR_TRADER_BUY)
	case trader_dat.GHOST:
		addGhost(state, itemId, source, tlog.IFR_TRADER_BUY, xdEventType)
	case trader_dat.HEART:
		module.Heart.IncHeart(state, num)
	case trader_dat.COINS:
		module.Player.IncMoney(state.Database, state.MoneyState, int64(num), player_dat.COINS, tlog.MFR_TRADER_BUY, xdEventType, "")
	case trader_dat.INGOT:
		module.Player.IncMoney(state.Database, state.MoneyState, int64(num), player_dat.INGOT, tlog.MFR_TRADER_BUY, xdEventType, "")
	case trader_dat.PHYSICAL:
		module.Physical.AwardIncrease(state, num, tlog.PFR_TRADER_BUY)
	case trader_dat.PET:
		//如果怪物数量超过 int16 最大值则可能溢出
		addPet(state, int32(itemId), source, tlog.IFR_TRADER_BUY, xdEventType)
	default:
		panic("unsupport goods type")
	}
}

func addGhost(state *module.SessionState, ghostId int16, source string, itemReason, xdEventType int32) {
	ghost := ghost_dat.GetGhost(ghostId)
	funcName := player_dat.GetFuncById(player_dat.FUNC_GHOST).Name
	if !module.Player.IsOpenFunc(state.Database, player_dat.FUNC_GHOST) {
		rpc.RemoteMailSend(state.PlayerId, mail_dat.MailPurchaseTips{
			Source:   source,
			ItemName: ghost.Name,
			Func:     funcName,
		})
	}
	module.Ghost.AddGhost(state, ghostId, itemReason, xdEventType)
}

func addSwordSoul(state *module.SessionState, swordSoulId int16, source string, itemReason int32) {
	swordSoul := sword_soul_dat.GetSwordSoul(swordSoulId)
	funcName := player_dat.GetFuncById(player_dat.FUNC_SWORD_SOUL).Name
	if !module.Player.IsOpenFunc(state.Database, player_dat.FUNC_SWORD_SOUL) {
		rpc.RemoteMailSend(state.PlayerId, mail_dat.MailPurchaseTips{
			Source:   source,
			ItemName: swordSoul.Name,
			Func:     funcName,
		})
	}
	module.SwordSoul.AddSwordSoul(state, swordSoulId, itemReason)
}

func addPet(state *module.SessionState, enemyId int32, source string, itemFlowReason, xdEventType int32) {
	module.BattlePet.AddPet(state.Database, enemyId, itemFlowReason, xdEventType)
}
