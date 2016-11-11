package player

import (
	"math/rand"

	"core/fail"
	"core/time"

	"game_server/dat/coins_exchange_dat"
	"game_server/dat/player_dat"
	"game_server/dat/vip_dat"
	"game_server/mdb"
	"game_server/module"
)

//购买铜币
func buyCoins(state *module.SessionState, playerCoinsInfo *mdb.PlayerCoins) (ingot, coins int64, crit bool) {
	//检查购买次数
	var maxBuyCount int16
	maxBuyCount = int16(module.VIP.PrivilegeTimes(state, vip_dat.GOUMAITONGQIAN))

	fail.When(playerCoinsInfo.DailyCount >= maxBuyCount, "DailyCount is full")
	exchangeDat := coins_exchange_dat.GetCoinsExchangeInfo(playerCoinsInfo.DailyCount + 1)

	playerCoinsInfo.DailyCount++
	playerCoinsInfo.BuyTime = time.GetNowTime()

	find := false
	for i, max := 0, len(player_dat.INGOT_TO_COINS_DOUBLE_NUMS); i < max; i++ {
		if player_dat.INGOT_TO_COINS_DOUBLE_NUMS[i] == int(playerCoinsInfo.DailyCount) {
			find = true
			break
		}
	}

	//计算是否暴击产生双倍铜币
	if find {
		crit = true
	} else {
		crit = rand.Int63()%100 < coins_exchange_dat.INGOT_TO_CONIS_CRIT_RATE
	}
	ingot = exchangeDat.Ingot
	coins = exchangeDat.Coins
	if crit {
		coins *= 2
	}
	return
}

//更新购买铜币次数
func updateCoinsBuyCount(state *module.SessionState, playerCoinsInfo *mdb.PlayerCoins) *mdb.PlayerCoins {
	if !time.IsToday(playerCoinsInfo.BuyTime) {
		playerCoinsInfo.DailyCount = 0
		playerCoinsInfo.BatchBought = 0
		state.Database.Update.PlayerCoins(playerCoinsInfo)
	}
	return playerCoinsInfo
}
