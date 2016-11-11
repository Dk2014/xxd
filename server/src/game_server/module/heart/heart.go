package heart

import (
	"core/fail"
	"core/time"
	"game_server/dat/heart_dat"
	"game_server/mdb"
	"game_server/module"
	gotime "time"
)

func init() {
	module.Heart = HeartMod{}
}

type HeartMod struct {
}

func (this HeartMod) GetHeartDailyNum(state *module.SessionState) int32 {
	playerHeart := state.Database.Lookup.PlayerHeart(state.PlayerId)
	return getAddDayCount(playerHeart)
}

func (this HeartMod) LoginUpdateHeartNum(state *module.SessionState, playerHeart *mdb.PlayerHeart) {
	//更新每日获取从好友处获得的爱心数量
	playerHeart.AddDayCount = getAddDayCount(playerHeart)
	if playerHeart.Value >= heart_dat.HEART_UPPER {
		return
	}
	if playerHeart.Value < 0 {
		playerHeart.Value = 0
	}

	oldValue := playerHeart.Value
	nowTime := time.GetNowTime()
	passTime := nowTime - playerHeart.UpdateTime
	deltaT := passTime % heart_dat.HEART_RECOVER_TIME_INTERVAL
	todayZero := time.GetTodayZero()

	recoverValue := int64(playerHeart.Value) + (passTime / heart_dat.HEART_RECOVER_TIME_INTERVAL * heart_dat.HEART_RECOVER_NUM)
	if recoverValue > heart_dat.HEART_UPPER {
		recoverValue = heart_dat.HEART_UPPER
	}
	playerHeart.Value = int16(recoverValue)

	//上次恢复时间不在今天
	if playerHeart.UpdateTime < todayZero {
		//这一次下线的时间内，属于今日的恢复量
		maxRecoverToday := heart_dat.HEART_UPPER - oldValue - int16((todayZero-playerHeart.UpdateTime)/heart_dat.HEART_RECOVER_TIME_INTERVAL*heart_dat.HEART_RECOVER_NUM)
		if maxRecoverToday <= 0 {
			//昨天已经恢复满了，不算入今日恢复
			playerHeart.RecoverDayCount = 0
		} else {
			//今日以前的时间未能使爱心恢复满，则使用今日的时间来恢复爱心
			//这里会有一些误差，假设玩家在23:55分开始恢复爱心次日00:05分获得爱心，这个爱心将即不属于今天也不属于昨天
			playerHeart.RecoverDayCount = int16((nowTime - todayZero) / heart_dat.HEART_RECOVER_TIME_INTERVAL * heart_dat.HEART_RECOVER_NUM)
			if playerHeart.RecoverDayCount > maxRecoverToday {
				playerHeart.RecoverDayCount = maxRecoverToday
			}
		}
	} else {
		playerHeart.RecoverDayCount += playerHeart.Value - oldValue
	}

	playerHeart.UpdateTime = nowTime - deltaT

	state.Database.Update.PlayerHeart(playerHeart)
	startTimer(state, playerHeart, gotime.Duration(heart_dat.HEART_RECOVER_TIME_INTERVAL-deltaT))
}

//还有赠送爱心需要特殊处理
func (this HeartMod) IncHeartFromFriend(state *module.SessionState, val int16) {
	playerHeart := state.Database.Lookup.PlayerHeart(state.PlayerId)
	playerHeart.AddDayCount = getAddDayCount(playerHeart)
	fail.When(playerHeart.AddDayCount >= heart_dat.HEART_DAY_COUNT_MAX, "IncHeart wrong count")

	increase(state, playerHeart, val)
	playerHeart.AddDayCount++
	playerHeart.AddTime = time.GetNowTime()

	if session, ok := module.Player.GetPlayerOnline(state.PlayerId); ok {
		module.Notify.SendHeartChange(session, playerHeart.Value)
	}
	state.Database.Update.PlayerHeart(playerHeart)
}

//增加爱心
func (this HeartMod) IncHeart(state *module.SessionState, val int16) {
	playerHeart := state.Database.Lookup.PlayerHeart(state.PlayerId)

	increase(state, playerHeart, val)

	if session, ok := module.Player.GetPlayerOnline(state.PlayerId); ok {
		module.Notify.SendHeartChange(session, playerHeart.Value)
	}
	state.Database.Update.PlayerHeart(playerHeart)
}

func (this HeartMod) DecHeart(db *mdb.Database, val int16) {
	playerHeart := db.Lookup.PlayerHeart(db.PlayerId())
	fail.When(playerHeart.Value < val, "DecHeart heart not enough")

	oldValue := playerHeart.Value

	playerHeart.Value -= val

	if session, ok := module.Player.GetPlayerOnline(db.PlayerId()); ok {
		module.Notify.SendHeartChange(session, playerHeart.Value)
		//只有在原先超过10个 减少后少于10个才需要启动计时器
		state := module.State(session)
		if oldValue >= heart_dat.HEART_UPPER && playerHeart.Value < heart_dat.HEART_UPPER {
			startTimer(state, playerHeart, heart_dat.HEART_RECOVER_TIME_INTERVAL)
		}

	}
	if oldValue >= heart_dat.HEART_UPPER && playerHeart.Value < heart_dat.HEART_UPPER {
		playerHeart.UpdateTime = time.GetNowTime()
	}
	db.Update.PlayerHeart(playerHeart)
}

func increase(state *module.SessionState, playerHeart *mdb.PlayerHeart, val int16) {
	playerHeart.Value += val

	if playerHeart.Value >= heart_dat.HEART_UPPER {
		stopTimer(state)
	}
}

func increaseByTimer(state *module.SessionState) {
	playerHeart := state.Database.Lookup.PlayerHeart(state.PlayerId)
	oldValue := playerHeart.Value

	increase(state, playerHeart, heart_dat.HEART_RECOVER_NUM)
	if playerHeart.Value > heart_dat.HEART_UPPER {
		playerHeart.Value = heart_dat.HEART_UPPER
	}

	if time.IsToday(playerHeart.UpdateTime) {
		playerHeart.RecoverDayCount += (playerHeart.Value - oldValue)
	} else {
		playerHeart.RecoverDayCount = playerHeart.Value - oldValue
	}
	playerHeart.UpdateTime = time.GetNowTime()

	state.Database.Update.PlayerHeart(playerHeart)

	if session, ok := module.Player.GetPlayerOnline(state.PlayerId); ok {
		module.Notify.SendHeartChange(session, playerHeart.Value)
	}
	startTimer(state, playerHeart, heart_dat.HEART_RECOVER_TIME_INTERVAL)
}

func startTimer(state *module.SessionState, playerHeart *mdb.PlayerHeart, delaySecs gotime.Duration) {
	if playerHeart.Value >= heart_dat.HEART_UPPER {
		return
	}
	//if session, ok := module.Player.GetPlayerOnline(state.PlayerId); ok {
	//	module.Notify.SendHeartRecoverTime(session, int64(delaySecs)+time.GetNowTime())
	//}
	state.TimerMgr.Start(module.TIMER_HEART, delaySecs*gotime.Second, increaseByTimer)
}

func stopTimer(state *module.SessionState) {
	state.TimerMgr.Stop(module.TIMER_HEART)
}

//获得本日可领取爱心次数
func getAddDayCount(playerHeart *mdb.PlayerHeart) (count int32) {
	count = playerHeart.AddDayCount
	fail.When(count < 0, "getAddDayCount wrong count")

	if !time.IsToday(playerHeart.AddTime) {
		count = 0
	}
	return
}
