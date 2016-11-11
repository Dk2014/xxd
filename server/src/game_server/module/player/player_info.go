package player

import (
	"core/fail"
	"game_server/config"
	"game_server/dat/player_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/tencent"
	"game_server/tlog"
	"game_server/xdlog"
)

func enableTencentMoney() bool {
	return config.ServerCfg.MoneySdk.MoneySdkAddr != ""
}

func TencentPayTest(moneyState *module.MoneyState, callback func(*tencent.GET_BALANCE_RSP)) {
	if enableTencentMoney() {
		tencent.GetIngotTest(moneyState, callback)
	}
}

func (mod PlayerMod) CheckMoney(state *module.SessionState, num int64, mtype int) bool {
	if mtype == player_dat.COINS || !enableTencentMoney() {
		playerInfo := state.Database.Lookup.PlayerInfo(state.PlayerId)
		return checkMoney(playerInfo, num, mtype)
	} else {
		return tencent.CheckIngot(state.MoneyState, num)
	}
}

func (mod PlayerMod) GetIngot(db *mdb.Database, moneyState *module.MoneyState) (ingot int64) {
	var amtIngot int64
	//var presentingot int64
	if !enableTencentMoney() {
		//走非腾讯货币，不再更新vip等级，想要vip通过debug接口接入
		ingot = db.Lookup.PlayerInfo(db.PlayerId()).Ingot
	} else {
		ingot = tencent.GetIngot(moneyState)
		amtIngot = tencent.GetSaveAmt(moneyState)
		if amtIngot > 0 {
			module.VIP.UpdateIngot(db, int64(amtIngot))
		}
	}
	return
}

func (mod PlayerMod) IncMoney(db *mdb.Database, moneyState *module.MoneyState, num int64, mtype int, moneyFlowReason, xdEventType int32, xdEventParam string) {
	fail.When(mtype != player_dat.COINS && mtype != player_dat.INGOT, "money wrong type")
	fail.When(num < 0, "negative num")

	var aftermoney, imoney, imoneytype, xdlogmoneytype int32
	var moneytype int8
	var value int64

	imoney = int32(num)
	playerInfo := db.Lookup.PlayerInfo(db.PlayerId())

	//当操作铜钱，则不用管是否腾讯货币，直接处理
	if mtype == player_dat.COINS {
		playerInfo.Coins += num
		aftermoney = int32(playerInfo.Coins)
		imoneytype = tlog.MT_COIN
		xdlogmoneytype = xdlog.MT_COIN
		moneytype = int8(player_dat.COINS)
		value = playerInfo.Coins
		db.Update.PlayerInfo(playerInfo)
	} else if mtype == player_dat.INGOT { //操作元宝，先判断是不是走腾讯虚拟币
		if !enableTencentMoney() { //操作元宝，先判断是不是走腾讯虚拟币
			playerInfo.Ingot += num
			aftermoney = int32(playerInfo.Ingot)
			imoneytype = tlog.MT_INGOT
			xdlogmoneytype = xdlog.MT_INGOT
			moneytype = int8(player_dat.INGOT)
			value = playerInfo.Ingot
			db.Update.PlayerInfo(playerInfo)
		} else { //走腾讯，则操作调用接口
			ingot := tencent.IncIngot(moneyState, player_dat.TPRESENT_DEFAULT_DISCOUNTID, player_dat.TPRESENT_DEFAULT_GIFTID, num)
			imoneytype = tlog.MT_INGOT
			xdlogmoneytype = xdlog.MT_INGOT
			aftermoney = int32(ingot)
			moneytype = int8(player_dat.INGOT)
			value = ingot
		}
	}

	xdlog.IncomeLog(db, imoney, aftermoney, xdlogmoneytype, xdEventType, xdEventParam)
	moneyFlow(db, aftermoney, imoney, tlog.ADD, imoneytype, moneyFlowReason)

	//通知玩家货币变动
	if session, ok := module.Player.GetPlayerOnline(db.PlayerId()); ok {
		module.Notify.SendMoneyChange(session, moneytype, value)
	}
}

func (mod PlayerMod) IncSwordSoulFragment(db *mdb.Database, num int64, ftype int, moneyFlowReason, xdEventType int32) {
	fail.When(ftype != player_dat.SWORDSOULFRAGMENT, "money wrong type")
	fail.When(num < 0, "negative num")

	var afterfragment, ifragment, ifragmenttype int32
	var fragmenttype int8
	var value int64

	ifragment = int32(num)
	playerInfo := db.Lookup.PlayerInfo(db.PlayerId())

	//增加剑心碎片数量
	playerInfo.SwordSoulFragment += num
	afterfragment = int32(playerInfo.SwordSoulFragment)
	ifragmenttype = tlog.MT_SWORD_SOUL_FRAGMENT
	fragmenttype = int8(player_dat.SWORDSOULFRAGMENT)
	value = playerInfo.SwordSoulFragment
	db.Update.PlayerInfo(playerInfo)

	moneyFlow(db, afterfragment, ifragment, tlog.ADD, ifragmenttype, moneyFlowReason)
	xdlog.IncomeLog(db, ifragment, afterfragment, xdlog.MT_SWORD_SOUL_FRAGMENT, xdEventType, "")

	//通知玩家货币变动
	if session, ok := module.Player.GetPlayerOnline(db.PlayerId()); ok {
		module.Notify.SendMoneyChange(session, fragmenttype, value)
	}
}

//判断用户是否是会员
func (mod PlayerMod) QQVipStatus(state *module.SessionState) int16 {
	if config.ServerCfg.MSDKApiAddr != "" {
		return tencent.QQVipStatus(state.MoneyState)
	}
	return 0
}

func (mod PlayerMod) DecMoney(db *mdb.Database, moneyState *module.MoneyState, num int64, mtype int, moneyFlowReason, xdEventType int32) {
	var aftermoney, imoney, imoneytype, xdlogmoneytype int32
	var moneytype int8
	var value int64

	imoney = int32(num)
	playerInfo := db.Lookup.PlayerInfo(db.PlayerId())

	//当操作铜钱，则不用管是否腾讯货币，直接处理
	if mtype == player_dat.COINS {
		fail.When(!checkMoney(playerInfo, num, mtype), "money not enough")
		playerInfo.Coins -= num
		aftermoney = int32(playerInfo.Coins)
		imoneytype = tlog.MT_COIN
		xdlogmoneytype = xdlog.MT_COIN
		moneytype = player_dat.COINS
		value = playerInfo.Coins
		db.Update.PlayerInfo(playerInfo)
	} else if mtype == player_dat.INGOT { //操作元宝，先判断是不是走腾讯虚拟币
		if !enableTencentMoney() { //不走腾讯，则操作数据库控制元宝
			fail.When(!checkMoney(playerInfo, num, mtype), "money not enough")
			playerInfo.Ingot -= num
			aftermoney = int32(playerInfo.Ingot)
			imoneytype = tlog.MT_INGOT
			xdlogmoneytype = xdlog.MT_INGOT
			moneytype = player_dat.INGOT
			value = playerInfo.Ingot
			db.Update.PlayerInfo(playerInfo)
		} else { //走腾讯，则操作调用接口
			ingot := tencent.DecIngot(moneyState, num)
			imoneytype = tlog.MT_INGOT
			xdlogmoneytype = xdlog.MT_INGOT
			aftermoney = int32(ingot)
			moneytype = player_dat.INGOT
			value = ingot
		}

		if moneyFlowReason != tlog.MFR_CLIQUE_DONATE_BANK_TRADE {
			module.Event.UpdateEventsIngot(db, num, false)
			module.Event.UpdateJsonEventTotalConsume(db, int32(num))
			//单条消耗奖励json
			module.Event.UpdateJsonEventSingleConsume(db, int32(num))
			module.Event.ChangeNewYearConsumeStatus(db, int32(num))
		}

	}

	xdlog.ConsumeLog(db, imoney, aftermoney, xdlogmoneytype, xdEventType)
	moneyFlow(db, aftermoney, imoney, tlog.REDUCE, imoneytype, moneyFlowReason)

	//通知玩家货币变动
	if session, ok := module.Player.GetPlayerOnline(db.PlayerId()); ok {
		module.Notify.SendMoneyChange(session, moneytype, value)
	}
}

func (mod PlayerMod) DecSwordSoulFragment(db *mdb.Database, num int64, ftype int, moneyFlowReason, xdEventType int32) {
	var afterfragment, ifragment, ifragmenttype int32
	var fragmenttype int8
	var value int64

	ifragment = int32(num)
	playerInfo := db.Lookup.PlayerInfo(db.PlayerId())

	fail.When(!checkFragment(playerInfo, num, ftype), "money not enough")
	playerInfo.SwordSoulFragment -= num
	afterfragment = int32(playerInfo.SwordSoulFragment)
	ifragmenttype = tlog.MT_SWORD_SOUL_FRAGMENT
	fragmenttype = player_dat.SWORDSOULFRAGMENT
	value = playerInfo.SwordSoulFragment
	db.Update.PlayerInfo(playerInfo)

	moneyFlow(db, afterfragment, ifragment, tlog.REDUCE, ifragmenttype, moneyFlowReason)
	xdlog.ConsumeLog(db, ifragment, afterfragment, xdlog.MT_SWORD_SOUL_FRAGMENT, xdEventType)

	//通知玩家货币变动
	if session, ok := module.Player.GetPlayerOnline(db.PlayerId()); ok {
		module.Notify.SendMoneyChange(session, fragmenttype, value)
	}
}

func checkMoney(playerInfo *mdb.PlayerInfo, num int64, mtype int) bool {
	fail.When(mtype != player_dat.COINS && mtype != player_dat.INGOT, "money wrong type")
	fail.When(num < 0, "negative num")
	if mtype == player_dat.COINS {
		return playerInfo.Coins >= num
	}
	return playerInfo.Ingot >= num
}

func checkFragment(playerInfo *mdb.PlayerInfo, num int64, ftype int) bool {
	fail.When(ftype != player_dat.SWORDSOULFRAGMENT, "fragment wrong type")
	fail.When(num < 0, "negative num")
	return playerInfo.SwordSoulFragment >= num
}

//游戏货币流水表
func moneyFlow(db *mdb.Database, aftermoney, imoney, addorreduce, imoneytype int32, moneyFlowReason int32) {
	openid := module.Player.GetPlayer(db).User
	role_vip := module.VIP.VIPInfo(db)
	db.AddTLog(tlog.NewMoneyFlow(0, openid, int32(module.Role.GetMainRole(db).Level), aftermoney, imoney, moneyFlowReason, 0, addorreduce, imoneytype, int32(role_vip.Level)))
}

//获取月卡信息，并分析是否需要赠送
func checekMonthCard(db *mdb.Database, row *mdb.PlayerMonthCardInfo, moneyState *module.MoneyState, begintime, endtime, opendays, ingot int64) (addingot int64) {
	buytimes := opendays / player_dat.MONTH_CARD_PER_TIMES
	if row == nil { //之前没记录过月卡信息
		addingot = int64(buytimes) * ingot
		module.Player.IncMoney(db, moneyState, addingot, player_dat.INGOT, tlog.MFR_MONTH_CARD, xdlog.ET_MONTH_CARD, "")
		db.Insert.PlayerMonthCardInfo(&mdb.PlayerMonthCardInfo{
			Starttime:    begintime,
			Endtime:      endtime,
			Pid:          db.PlayerId(),
			Presenttotal: addingot,
			Buytimes:     buytimes,
		})
	} else if buytimes > row.Buytimes { //如果记录的购买次数和算出来的不一样，说明新的月卡已购买
		addingot = ingot * int64(buytimes-row.Buytimes)
		module.Player.IncMoney(db, moneyState, addingot, player_dat.INGOT, tlog.MFR_MONTH_CARD, xdlog.ET_MONTH_CARD, "")
		row.Starttime = begintime
		row.Endtime = endtime
		row.Buytimes = buytimes
		row.Presenttotal += addingot
		db.Update.PlayerMonthCardInfo(row)
	}
	return
}
