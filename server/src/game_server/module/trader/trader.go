package trader

import (
	"core/fail"
	"core/net"
	"core/time"
	"game_server/api/protocol/trader_api"
	"game_server/dat/item_dat"
	"game_server/dat/player_dat"
	"game_server/dat/trader_dat"
	"game_server/dat/vip_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/tlog"
	"game_server/xdlog"
)

//获取巡游商人出现时间计划
func info(session *net.Session, out *trader_api.Info_Out) {
	now := time.GetNowTime()
	schedules := trader_dat.MapTraderSchedule[trader_dat.XUNYOUSHANGREN]
	for _, schedule := range schedules {
		if schedule.Expire == 0 || schedule.Expire > now {
			out.During = append(out.During, trader_api.Info_Out_During{
				Expire:    schedule.Expire,
				Showup:    schedule.Showup,
				Disappear: schedule.Disappear,
			})
		}
	}
}

//获取玩家商店信息
func storeSate(session *net.Session, traderId int16, out *trader_api.StoreState_Out) {
	state := module.State(session)
	checkOpen(state, traderId)

	playerRefreshState := updateRefreshState(state, traderId)
	out.RefreshNum = playerRefreshState.RefreshNum

	grids := updateStoreState(state, traderId, playerRefreshState)
	for _, grid := range grids {
		out.Goods = append(out.Goods, trader_api.StoreState_Out_Goods{
			GridId:    grid.GridId,
			GoodsType: grid.GoodsType,
			ItemId:    grid.ItemId,
			Cost:      grid.Cost,
			Num:       grid.Num,
			Stock:     grid.Stock,
		})
	}
}

//购买指定格子内的物品
func buy(session *net.Session, gridId, xdEventType int32, out *trader_api.Buy_Out) {
	state := module.State(session)
	playerGrid := getGridById(state, gridId)
	//检查是否在开放时间
	checkOpen(state, playerGrid.TraderId)

	//检查货物是否已经被刷新
	playerRefreshState := getRefreshState(state, playerGrid.TraderId)
	if isGoodsExpired(playerGrid.TraderId, playerRefreshState.AutoUpdateTime) {
		out.Expired = true
		return
	}

	//检查库存
	fail.When(playerGrid.Stock < 1, "sold out!")

	moneyType := trader_dat.GridInfo(gridId).MoneyType

	traderName := trader_dat.GetTraderNameById(playerGrid.TraderId)
	decMoney(state, moneyType, playerGrid.Cost, xdEventType)
	addGoods(state, playerGrid.GoodsType, playerGrid.ItemId, playerGrid.Num, traderName, xdEventType)

	if playerGrid.GoodsType == trader_dat.ITEM {
		item := item_dat.GetItem(playerGrid.ItemId)
		switch moneyType {
		case trader_dat.COINS_TYPE:
			module.Item.ItemMoneyFlow(state.Database, item, int32(playerGrid.Num), int32(playerGrid.Cost), tlog.MT_COIN)
			tlog.PlayerTradeBuyFlowLog(state.Database, int32(item.Id), int32(playerGrid.Num), int32(playerGrid.Cost), tlog.MT_COIN)
		case trader_dat.INGOT_TYPE:
			module.Item.ItemMoneyFlow(state.Database, item, int32(playerGrid.Num), int32(playerGrid.Cost), tlog.MT_INGOT)
			tlog.PlayerTradeBuyFlowLog(state.Database, int32(item.Id), int32(playerGrid.Num), int32(playerGrid.Cost), tlog.MT_INGOT)
		case trader_dat.HEART_TYPE:
			module.Item.ItemMoneyFlow(state.Database, item, int32(playerGrid.Num), int32(playerGrid.Cost), tlog.MT_HEART)
			tlog.PlayerTradeBuyFlowLog(state.Database, int32(item.Id), int32(playerGrid.Num), int32(playerGrid.Cost), tlog.MT_HEART)
		}
	}

	playerGrid.Stock -= 1
	state.Database.Update.PlayerTraderStoreState(playerGrid)
	out.Expired = false
	tlog.PlayerSystemModuleFlowLog(state.Database, tlog.SMT_TRADE)
}

//商店刷新
func refresh(session *net.Session, traderId int16, out *trader_api.Refresh_Out) {
	state := module.State(session)
	fail.When(!trader_dat.Refreshable(traderId), "grader is not refreshable")
	checkOpen(state, traderId)

	playerRefreshState := updateRefreshState(state, traderId)

	var refreshTime int16
	switch traderId {
	case trader_dat.YINGHAIJISHI:
		refreshTime = module.VIP.PrivilegeTimes(state, vip_dat.YINGHAIJISHI)
	}
	fail.When(playerRefreshState.RefreshNum >= refreshTime, "reach max refresh num")
	refreshPrice := trader_dat.RefreshPrice(traderId, playerRefreshState.RefreshNum+1)
	module.Player.DecMoney(state.Database, state.MoneyState, refreshPrice, player_dat.INGOT, tlog.MFR_TRADER_REFRESH, xdlog.ET_TRADER)

	var playerGrids []*mdb.PlayerTraderStoreState
	playerGrids = refreshStoreState(state, traderId)
	playerRefreshState.RefreshNum++
	playerRefreshState.LastUpdateTime = time.GetNowTime()
	state.Database.Update.PlayerTraderRefreshState(playerRefreshState)
	for _, playerGrid := range playerGrids {
		out.Goods = append(out.Goods, trader_api.Refresh_Out_Goods{
			GridId:    playerGrid.GridId,
			GoodsType: playerGrid.GoodsType,
			ItemId:    playerGrid.ItemId,
			Cost:      playerGrid.Cost,
			Num:       playerGrid.Num,
			Stock:     playerGrid.Stock,
		})
	}
}
