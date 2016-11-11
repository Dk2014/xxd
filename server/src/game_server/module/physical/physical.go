package physical

import (
	"core/fail"
	"core/net"
	"core/time"
	"game_server/dat/event_dat"
	"game_server/dat/physical_buy_cost_config_dat"
	"game_server/dat/physical_dat"
	"game_server/dat/player_dat"
	"game_server/dat/vip_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/tlog"
	"game_server/xdlog"
	goTime "time"
)

/*
	玩家体力系统
*/

func init() {
	module.Physical = PhysicalMod{}
}

type PhysicalMod struct {
}

// 登陆增长体力（根据最后一次体力更新时间计算）
func (this PhysicalMod) LoginUpdatePhysical(state *module.SessionState, playerPhysical *mdb.PlayerPhysical) {
	updateDailyBuyCount(playerPhysical)
	if playerPhysical.Value >= physical_dat.MAX_PHYSICAL_VALUE_BY_TIME {
		return
	}
	if playerPhysical.Value < 0 {
		playerPhysical.Value = 0
	}

	nowTime := time.GetNowTime()
	passTime := nowTime - playerPhysical.UpdateTime
	deltaT := passTime % physical_dat.PHYSICAL_RECOVER_SECOND

	var physicalValue int64 //避免长时间不登陆回复体力溢出
	physicalValue = int64(playerPhysical.Value) + passTime/physical_dat.PHYSICAL_RECOVER_SECOND*physical_dat.PER_RECOVER_PHYSICAL_VALUE
	// 离线体力增长不允许超出上限
	if physicalValue > physical_dat.MAX_PHYSICAL_VALUE_BY_TIME {
		physicalValue = physical_dat.MAX_PHYSICAL_VALUE_BY_TIME
	}

	playerPhysical.Value = int16(physicalValue)
	playerPhysical.UpdateTime = nowTime - deltaT

	state.Database.Update.PlayerPhysical(playerPhysical)
	tlog.PlayerPhysicalFlowLog(state.Database, int32(passTime/physical_dat.PHYSICAL_RECOVER_SECOND*physical_dat.PER_RECOVER_PHYSICAL_VALUE), int32(playerPhysical.Value), tlog.ADD, tlog.PFR_LOGIN)
	startTimer(state, playerPhysical, goTime.Duration(physical_dat.PHYSICAL_RECOVER_SECOND-deltaT))
}

/*
	购买体力
	- 每日购买有次数限制
	- 满体力或超过上限则无法购买
*/
func (this PhysicalMod) BuyPhysical(session *net.Session) (int16, bool) {
	state := module.State(session)
	playerPhysical := state.Database.Lookup.PlayerPhysical(state.PlayerId)
	updateDailyBuyCount(playerPhysical)

	if playerPhysical.Value >= physical_dat.MAX_PHYSICAL_VALUE {
		return playerPhysical.DailyCount, false
	}

	// 购买次数上限
	maxBuyCount := int16(module.VIP.PrivilegeTimes(state, vip_dat.GOUMAITILI))

	fail.When(playerPhysical.DailyCount >= maxBuyCount, "购买体力次数达到每日上限")
	buy_physical_price := int64(physical_buy_cost_config_dat.GetCost(int32(playerPhysical.DailyCount)))
	fail.When(!module.Player.CheckMoney(state, buy_physical_price, player_dat.INGOT), "ingot not enough")

	playerPhysical.DailyCount += 1
	playerPhysical.BuyCount += 1
	playerPhysical.BuyUpdateTime = time.GetNowTime()
	// 增加购买的体力
	increase(state, playerPhysical, physical_dat.BUY_PHYSICAL_VALUE, tlog.PFR_BUY_PHYSICAL)

	state.Database.Update.PlayerPhysical(playerPhysical)

	module.Player.DecMoney(state.Database, state.MoneyState, buy_physical_price, player_dat.INGOT, tlog.MFR_BUY_PHYSICAL, xdlog.ET_BUY_PHYSICAL)

	module.Notify.SendPhysicalChange(session, playerPhysical.Value)
	return playerPhysical.DailyCount, true
}

func (this PhysicalMod) Decrease(db *mdb.Database, decVal int16, reason int32) bool {
	playerPhysical := db.Lookup.PlayerPhysical(db.PlayerId())
	fail.When(playerPhysical.Value < decVal, "physical not enough")

	playerPhysical.UpdateTime = time.GetNowTime()
	playerPhysical.Value -= decVal
	db.Update.PlayerPhysical(playerPhysical)

	tlog.PlayerPhysicalFlowLog(db, int32(decVal), int32(playerPhysical.Value), tlog.REDUCE, reason)
	session, ok := module.Player.GetPlayerOnline(db.PlayerId())
	if ok {
		module.Event.UpdateEventActivityStatus(session, decVal) //活跃度活动更新状态
		module.Notify.SendPhysicalChange(session, playerPhysical.Value)
		state := module.State(session)
		startTimer(state, playerPhysical, physical_dat.PHYSICAL_RECOVER_SECOND)
	} else {
		playerEventRecordBytes := db.Lookup.PlayerEventAwardRecord(db.PlayerId()).RecordBytes
		playerEventRecord := module.NewEventInfoList()
		playerEventRecord.Decode(playerEventRecordBytes)

		activit, isRefresh := module.AddPlayerActivity(db, int32(decVal)) //添加活跃度

		eventInfo, _ := event_dat.GetEventInfoById(event_dat.EVENT_PHYSICAL_AWARDS)
		if event_dat.CheckEventTime(eventInfo, event_dat.NOT_END) {
			if isRefresh {
				playerEventRecord.ClearState(db, event_dat.EVENT_PHYSICAL_AWARDS)
			}
			eventInfo := playerEventRecord.GetPlayerEventInfoById(event_dat.EVENT_PHYSICAL_AWARDS)
			nextLevel := event_dat.GetNextPhysical(eventInfo.MaxAward)
			if nextLevel <= activit && nextLevel > 0 { //nextLevel在超过活跃度最大限制时返回始终为0
				playerEventRecord.UpdateMax(db, event_dat.EVENT_PHYSICAL_AWARDS, nextLevel)
			}
		}

	}
	return true
}

// 如果体力值大于等于返回true，反之false
func (this PhysicalMod) CheckGE(state *module.SessionState, val int16) bool {
	playerPhysical := state.Database.Lookup.PlayerPhysical(state.PlayerId)
	return (playerPhysical.Value >= val)
}

func (this PhysicalMod) AwardIncrease(state *module.SessionState, val int16, pfrReason int32) {
	playerPhysical := state.Database.Lookup.PlayerPhysical(state.PlayerId)

	increase(state, playerPhysical, val, pfrReason)

	state.Database.Update.PlayerPhysical(playerPhysical)

	if session, ok := module.Player.GetPlayerOnline(state.PlayerId); ok {
		module.Notify.SendPhysicalChange(session, playerPhysical.Value)
	}
}

func increase(state *module.SessionState, playerPhysical *mdb.PlayerPhysical, incVal int16, reason int32) {
	playerPhysical.Value += incVal
	playerPhysical.UpdateTime = time.GetNowTime()

	if reason == tlog.PFR_TIMMER {
		// 定时增长体力不允许超出上限
		tlog.PlayerPhysicalFlowLog(state.Database, int32(incVal), int32(physical_dat.MAX_PHYSICAL_VALUE_BY_TIME), tlog.ADD, reason)
	} else {
		tlog.PlayerPhysicalFlowLog(state.Database, int32(incVal), int32(playerPhysical.Value), tlog.ADD, reason)
	}

	// 最终体力值大于自动体力恢复上限，那么不再进行体力恢复
	if playerPhysical.Value >= physical_dat.MAX_PHYSICAL_VALUE_BY_TIME {
		stopTimer(state)
	}
}

func updateDailyBuyCount(playerPhysical *mdb.PlayerPhysical) {
	// 如果最后一次购买时间不是当天，那么重置购买次数
	if !time.IsInPointHour(player_dat.RESET_BUY_PHYSICAL_TIMES_IN_HOUR, playerPhysical.BuyUpdateTime) {
		playerPhysical.DailyCount = 0
	}
}

func increaseByTimer(state *module.SessionState) {
	playerPhysical := state.Database.Lookup.PlayerPhysical(state.PlayerId)

	// 定时恢复1点体力
	increase(state, playerPhysical, physical_dat.PER_RECOVER_PHYSICAL_VALUE, tlog.PFR_TIMMER)

	// 定时增长体力不允许超出上限
	if playerPhysical.Value > physical_dat.MAX_PHYSICAL_VALUE_BY_TIME {
		playerPhysical.Value = physical_dat.MAX_PHYSICAL_VALUE_BY_TIME
	}

	state.Database.Update.PlayerPhysical(playerPhysical)

	if session, ok := module.Player.GetPlayerOnline(state.PlayerId); ok {
		module.Notify.SendPhysicalChange(session, playerPhysical.Value)
	}

	startTimer(state, playerPhysical, physical_dat.PHYSICAL_RECOVER_SECOND)
}

// 在线玩家定时恢复体力
func startTimer(state *module.SessionState, playerPhysical *mdb.PlayerPhysical, delaySecs goTime.Duration) {
	// 最终体力值小于允许体力上限，那么就开始恢复体力
	if playerPhysical.Value >= physical_dat.MAX_PHYSICAL_VALUE_BY_TIME {
		return
	}

	//通知体力下次体力恢复时间
	if session, ok := module.Player.GetPlayerOnline(state.PlayerId); ok {
		module.Notify.SendPhysicalRecoverTime(session, time.GetNowTime()+int64(delaySecs))
	}

	state.TimerMgr.Start(module.TIMER_PHYSICAL, delaySecs*goTime.Second, increaseByTimer)
}

func stopTimer(state *module.SessionState) {
	state.TimerMgr.Stop(module.TIMER_PHYSICAL)
}
