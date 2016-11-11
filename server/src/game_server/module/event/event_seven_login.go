package event

import (
	"core/fail"
	"core/net"
	"core/time"
	"game_server/api/protocol/event_api"
	"game_server/dat/event_dat"
	"game_server/dat/mail_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/module/rpc"
	"game_server/xdlog"
)

func getSevenInfo(session *net.Session) {
	eventId := int16(event_dat.EVENT_SEVEN_DAY_AWARDS)
	eventInfo, _ := event_dat.GetEventInfoById(eventId)
	out := &event_api.GetSevenInfo_Out{}
	if !event_dat.CheckEventTime(eventInfo, event_dat.NOT_END) {
		out.Status = 2 //活动已经结束
		session.Send(out)
		return
	}
	state := module.State(session)
	eventRecord := state.EventsState.GetPlayerEventInfoById(eventId)
	playerLevel := state.Database.Lookup.PlayerMissionLevel(state.PlayerId)
	max_day := event_dat.GetMaxWeightInSevenDay() //最大奖励到的天数 目前为7
	//max_day := event_dat.MAX_LOGIN_AWARDS_DAY //最大领取奖励的天数为5，后面的都改邮件发送
	if playerLevel.MaxLock < event_dat.SRD_LOCK {
		out.Status = 1 //活动未开启
		session.Send(out)
		return
	} else if eventRecord.LastUpdated == 0 {
		//需要开启
		state.EventsState.UpdateMax(state.Database, eventId, 1<<uint(max_day))
		eventRecord = state.EventsState.GetPlayerEventInfoById(eventId)
	}
	startTimeUnix := eventRecord.LastUpdated //活动开始时间
	fail.When(startTimeUnix == 0, "seven day event does not open")

	gap := int32(time.GetNowDay()) - int32(time.GetNowDayFromUnix(startTimeUnix)) + 1 //从第一天开始
	if gap > int32(max_day) {
		if gap <= int32(event_dat.GetMaxWeightInSevenDay()) && eventRecord.Awarded < 1<<uint(gap-1) {
			//发邮件
			attachs := event_dat.MailEventSevenLoginAward(gap)
			mail := new(mail_dat.MailQiriXinShouLi)
			mail.Attachments = append(mail.Attachments, attachs...)
			rpc.RemoteMailSend(state.PlayerId, mail)
			state.EventsState.List[eventId].Awarded += int32(1 << uint(gap-1))
			//更新至数据库
			state.Database.Update.PlayerEventAwardRecord(&mdb.PlayerEventAwardRecord{
				Pid:             state.Database.PlayerId(),
				RecordBytes:     state.EventsState.Encode(),
				JsonEventRecord: state.JsonEventsState.Encode(),
			})
		}
		out.Status = 2
		session.Send(out)
		return
	}

	out.Day = int16(gap)
	out.Schedule = eventRecord.Awarded
	session.Send(out)
}

func getSevenAward(session *net.Session) {
	state := module.State(session)
	eventId := int16(event_dat.EVENT_SEVEN_DAY_AWARDS)
	eventRecord := state.EventsState.GetPlayerEventInfoById(eventId)
	out := &event_api.GetSevenAward_Out{}

	startTimeUnix := eventRecord.LastUpdated //活动开始时间
	fail.When(startTimeUnix == 0, "seven day event does not open")
	gap := int32(time.GetNowDay()) - int32(time.GetNowDayFromUnix(startTimeUnix)) + 1 //从第一天开始
	award, ok := event_dat.GetEventSevenDayAward(gap)
	if !ok {
		out.Result = 1
		session.Send(out)
		return
	}
	today := int32(1 << uint(gap-1)) //今日对应的权值
	if eventRecord.Awarded >= today {
		out.Result = 1
		session.Send(out)
		return
	}
	addAwardContents(state, award, xdlog.ET_EVENT_CENTER_SEVEN_DAY_AWARDS)
	state.EventsState.List[eventId].Awarded += today
	//更新至数据库
	state.Database.Update.PlayerEventAwardRecord(&mdb.PlayerEventAwardRecord{
		Pid:             state.Database.PlayerId(),
		RecordBytes:     state.EventsState.Encode(),
		JsonEventRecord: state.JsonEventsState.Encode(),
	})
	out.Result = 0
	session.Send(out)
}
