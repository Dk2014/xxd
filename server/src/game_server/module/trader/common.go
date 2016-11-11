package trader

import (
	"core/fail"
	coreTime "core/time"
	"game_server/dat/trader_dat"
	"game_server/mdb"
	"game_server/module"
)

//检查货物是否已经被被刷新
func isGoodsExpired(traderId int16, updateTime int64) bool {
	isExpired := false
	isToday := coreTime.IsToday(updateTime)
	todayZeroTime := coreTime.GetTodayZero() //当日零点的时刻
	nowUnix := coreTime.GetNowTime()
	refreshSchedule := trader_dat.TraderRefreshSchedule(traderId)
	for _, scheduleTiming := range refreshSchedule {
		if updateTime < scheduleTiming+todayZeroTime &&
			(scheduleTiming+todayZeroTime <= nowUnix || !isToday) {
			isExpired = true
			break
		}
	}
	return isExpired
}

//更新（或初始化）玩家的刷新状态，只维护手工刷新的数据
func updateRefreshState(state *module.SessionState, traderId int16) *mdb.PlayerTraderRefreshState {
	refreshState := getRefreshState(state, traderId)
	if refreshState == nil {
		refreshState = new(mdb.PlayerTraderRefreshState)
		refreshState.Pid = state.PlayerId
		refreshState.TraderId = traderId
		refreshState.LastUpdateTime = coreTime.GetNowTime()
		state.Database.Insert.PlayerTraderRefreshState(refreshState)
	} else if !coreTime.IsToday(refreshState.LastUpdateTime) {
		refreshState.LastUpdateTime = coreTime.GetNowTime()
		refreshState.RefreshNum = 0
		state.Database.Update.PlayerTraderRefreshState(refreshState)
	}
	return refreshState
}

//刷新指定商人旗下的格子
func refreshStoreState(state *module.SessionState, traderId int16) []*mdb.PlayerTraderStoreState {
	playerGrids := getGridsByTraderId(state, traderId)
	mainRoleLevel := module.Role.GetMainRole(state.Database).Level
	for _, playerGrid := range playerGrids {
		playerGrid.Stock = 1
		playerGrid.GoodsType, playerGrid.ItemId, playerGrid.Num, playerGrid.Cost = randomGoodsByGridId(playerGrid.GridId, mainRoleLevel)
		state.Database.Update.PlayerTraderStoreState(playerGrid)
	}
	return playerGrids
}

//自动刷新
func updateStoreState(state *module.SessionState, traderId int16, playerRefreshState *mdb.PlayerTraderRefreshState) []*mdb.PlayerTraderStoreState {

	//计算是否需要刷新商店
	shouldAutoRefresh := isGoodsExpired(traderId, playerRefreshState.AutoUpdateTime)

	var playerGrids []*mdb.PlayerTraderStoreState
	state.Database.Select.PlayerTraderStoreState(func(row *mdb.PlayerTraderStoreStateRow) {
		if row.TraderId() == traderId {
			playerGrids = append(playerGrids, row.GoObject())
		}
	})
	gridIds := trader_dat.TraderGridIds(traderId)

	nowUnix := coreTime.GetNowTime()
	mainRoleLevel := module.Role.GetMainRole(state.Database).Level
	if len(playerGrids) == 0 {
		//第一次使用商店
		for _, gridId := range gridIds {
			playerGrid := &mdb.PlayerTraderStoreState{
				Pid:      state.PlayerId,
				TraderId: traderId,
				GridId:   gridId,
				Stock:    1,
			}
			playerGrid.GoodsType, playerGrid.ItemId, playerGrid.Num, playerGrid.Cost = randomGoodsByGridId(playerGrid.GridId, mainRoleLevel)
			state.Database.Insert.PlayerTraderStoreState(playerGrid)
			playerGrids = append(playerGrids, playerGrid)
		}
		playerRefreshState.AutoUpdateTime = nowUnix
		state.Database.Update.PlayerTraderRefreshState(playerRefreshState)
	} else if len(playerGrids) == len(gridIds) && shouldAutoRefresh {
		//更新
		for _, playerGrid := range playerGrids {
			playerGrid.Stock = 1
			playerGrid.GoodsType, playerGrid.ItemId, playerGrid.Num, playerGrid.Cost = randomGoodsByGridId(playerGrid.GridId, mainRoleLevel)
			state.Database.Update.PlayerTraderStoreState(playerGrid)
		}
		playerRefreshState.AutoUpdateTime = nowUnix
		state.Database.Update.PlayerTraderRefreshState(playerRefreshState)
	} else if shouldAutoRefresh {
		// 商人格子数有变化，需要手动清理这个号
		panic("玩家随机商店格子状态记录与商人格子数不一致")
	}
	return playerGrids
}

func checkOpen(state *module.SessionState, traderId int16) {
	var open bool
	switch traderId {
	case trader_dat.YINGHAIJISHI:
		open = checkOpenYingHai(state)
	case trader_dat.XUNYOUSHANGREN:
		open = checkOpenXunYou(state)
	default:
		panic("undefine trader id")
	}
	fail.When(!open, "store not open")
}

//检查巡游商人是否开放
func checkOpenXunYou(state *module.SessionState) bool {
	schedules := trader_dat.TraderScheduleInfo(trader_dat.XUNYOUSHANGREN)
	//今天零点是多少秒
	zeroTimeToday := coreTime.GetTodayZero()
	now := coreTime.GetNowTime()
	//此刻是今日的第几秒
	secondToday := now - zeroTimeToday
	open := false
	for _, schedule := range schedules {
		if schedule.Expire == 0 || schedule.Expire > now {
			if schedule.Showup < secondToday && secondToday < schedule.Disappear {
				open = true
				break
			}
		}
	}
	return open
}

//检查是否开放瀛海集市
func checkOpenYingHai(state *module.SessionState) bool {
	return module.Role.GetMainRole(state.Database).Level >= trader_dat.YINGHAISHANGREN_REQUIRE_LEVEL
}

//获取玩家刷新状态记录
func getRefreshState(state *module.SessionState, traderId int16) *mdb.PlayerTraderRefreshState {
	var refreshState *mdb.PlayerTraderRefreshState
	state.Database.Select.PlayerTraderRefreshState(func(row *mdb.PlayerTraderRefreshStateRow) {
		if row.TraderId() == traderId {
			refreshState = row.GoObject()
		}
	})
	return refreshState
}

//根据商店格子Id获取玩家格子状态
func getGridById(state *module.SessionState, gridId int32) (grid *mdb.PlayerTraderStoreState) {
	state.Database.Select.PlayerTraderStoreState(func(row *mdb.PlayerTraderStoreStateRow) {
		if row.GridId() == gridId {
			grid = row.GoObject()
			row.Break()
		}
	})
	return grid
}

//获取某个商人的所有格子状态
func getGridsByTraderId(state *module.SessionState, traderId int16) (grids []*mdb.PlayerTraderStoreState) {
	state.Database.Select.PlayerTraderStoreState(func(row *mdb.PlayerTraderStoreStateRow) {
		if row.TraderId() == traderId {
			grids = append(grids, row.GoObject())
		}
	})
	return grids
}
