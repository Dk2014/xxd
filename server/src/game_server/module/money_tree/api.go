package money_tree

import (
	"core/fail"
	"core/net"
	"core/time"
	"game_server/api/protocol/money_tree_api"
	"game_server/dat/money_tree_dat"
	"game_server/dat/player_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/tlog"
	"game_server/xdlog"
	"math/rand"
)

func init() {
	money_tree_api.SetInHandler(MoneyTreeAPI{})
}

type MoneyTreeAPI struct{}

func (api MoneyTreeAPI) GetTreeStatus(session *net.Session, in *money_tree_api.GetTreeStatus_In) {
	state := module.State(session)

	db := state.Database
	record := db.Lookup.PlayerMoneyTree(state.PlayerId)
	if record == nil { // new for money tree
		record = newMoneyTree(db)
	}

	out := &money_tree_api.GetTreeStatus_Out{}
	out.LastTime = record.LastWavedTime
	out.Money = record.Total
	out.Times = record.WavedTimes
	out.Remind = money_tree_dat.WAVE_TIMES_NEED_TREE - record.WavedTimesTotal
	if record.LastWavedTime > 0 && !time.IsToday(record.LastWavedTime) {
		// a new day
		out.Times = 0
		record.WavedTimes = 0
		db.Update.PlayerMoneyTree(record)
	}
	if record.WavedTimesTotal == money_tree_dat.WAVE_TIMES_NEED_TREE {
		out.Status = 2 // 默认CD中，或者满足可以领钱
	} else if time.GetNowTime()-record.LastWavedTime >= money_tree_dat.WAVE_GAP_SECONDS {
		out.Status = 1
	}
	session.Send(out)
}

func (api MoneyTreeAPI) GetTreeMoney(session *net.Session, in *money_tree_api.GetTreeMoney_In) {
	state := module.State(session)

	out := &money_tree_api.GetTreeMoney_Out{}
	out.Code = 1
	record := state.Database.Lookup.PlayerMoneyTree(state.PlayerId)
	fail.When(record == nil, "can not find player's money tree")
	fail.When(record.WavedTimesTotal != money_tree_dat.WAVE_TIMES_NEED_TREE, "this tree is not ready for achieving")
	if record.WavedTimesTotal == money_tree_dat.WAVE_TIMES_NEED_TREE {
		// can be awaded
		out.Code = 0
		// TODO 稍后添加tlog的原因信息
		module.Player.IncMoney(state.Database, state.MoneyState, int64(record.Total), player_dat.COINS, tlog.MFR_EVENT_CENTER, xdlog.ET_MONEY_TREE, "")

		// reinit
		record.Total = 0
		record.WavedTimesTotal = 0
		state.Database.Update.PlayerMoneyTree(record)
	}
	session.Send(out)
}

func (api MoneyTreeAPI) WaveTree(session *net.Session, in *money_tree_api.WaveTree_In) {
	state := module.State(session)
	record := state.Database.Lookup.PlayerMoneyTree(state.PlayerId)
	fail.When(record == nil, "can not find player's money tree")

	out := &money_tree_api.WaveTree_Out{}

	// check
	fail.When(record.WavedTimes >= money_tree_dat.WAVE_TIMES_PER_DAY, "wave times is full today")
	fail.When(record.WavedTimesTotal >= money_tree_dat.WAVE_TIMES_NEED_TREE, "wave times of this tree is over")
	fail.When(record.LastWavedTime+money_tree_dat.WAVE_GAP_SECONDS > time.GetNowTime(), "wave time gap still exists")

	coins := rand.Int31n(money_tree_dat.MAX_MONEY-money_tree_dat.MIN_MONEY+1) + money_tree_dat.MIN_MONEY // 本次摇下来的钱数
	record.Total += coins
	record.WavedTimes += 1
	record.WavedTimesTotal += 1
	record.LastWavedTime = time.GetNowTime()
	state.Database.Update.PlayerMoneyTree(record)

	out.Remaind = money_tree_dat.WAVE_TIMES_NEED_TREE - record.WavedTimesTotal
	out.Money = coins
	if record.WavedTimesTotal == money_tree_dat.WAVE_TIMES_NEED_TREE {
		out.Status = 2 // 默认CD中，或者满足可以领钱
	} else if time.GetNowTime()-record.LastWavedTime >= money_tree_dat.WAVE_GAP_SECONDS {
		out.Status = 1
	}
	session.Send(out)
}

func newMoneyTree(db *mdb.Database) *mdb.PlayerMoneyTree {
	record := &mdb.PlayerMoneyTree{}

	// init fields
	record.Pid = db.PlayerId()
	db.Insert.PlayerMoneyTree(record)
	return record
}
