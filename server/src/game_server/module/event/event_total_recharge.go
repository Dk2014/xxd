package event

import (
	"core/fail"
	"core/net"
	"game_server/api/protocol/event_api"
	"game_server/dat/event_dat"
	"game_server/module"
	"game_server/xdlog"
)

func getEventTotalRechargeAward(session *net.Session, order int16) {
	state := module.State(session)
	eventId := int16(event_dat.EVENT_TOTAL_RECHARGE)
	out := &event_api.GetEventTotalAward_Out{}

	eventInfo, _ := event_dat.GetEventInfoById(eventId)
	if !event_dat.CheckEventTime(eventInfo, event_dat.NOT_END) {
		out.Result = 3
		session.Send(out)
		return
	}

	eventRecord := state.EventsState.GetPlayerEventInfoById(eventId)
	if eventRecord.Awarded&(1<<uint(order)) != 0 { //已经领过
		out.Result = 3
		session.Send(out)
		return
	}

	info := event_dat.GetEventTotalRechargeByOrder(order)
	fail.When(info == nil, "Award does not exists")

	total := module.Event.GetEventsIngot(state.Database, true)
	if total < int64(info.RequireTotal) { //没达到
		out.Result = 3
		session.Send(out)
		return
	}

	eventRecord.Awarded += 1 << uint(order)
	state.EventsState.AddEventAwardState(state.Database, eventId, 0, eventRecord.Awarded)
	//开始添加奖励
	addAwardContents(state, info.Awards, xdlog.ET_EVENT_CENTER_TOTAL_RECHARGE)
	if eventRecord.Awarded == (1<<uint(event_dat.CountEventTotalRecharge()) - 1) {
		out.Result = 2
	}
	session.Send(out)
}
