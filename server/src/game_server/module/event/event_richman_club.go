package event

import (
	"core/net"
	"game_server/api/protocol/event_api"
	"game_server/dat/event_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/module/rpc"
	"game_server/tlog"
	"game_server/xdlog"
)

func getEventRichmanClubStatus(session *net.Session) {
	state := module.State(session)
	eventId := int16(event_dat.EVENT_RICHMAN_CLUB)
	eventInfo, _ := event_dat.GetEventInfoById(eventId)

	out := &event_api.GetRichmanClubInfo_Out{}
	if !event_dat.CheckEventTime(eventInfo, event_dat.NOT_END) {
		//活动已结束
		out.Status = 1
		session.Send(out)
		return
	}

	eventRecord := state.EventsState.GetPlayerEventInfoById(eventId)
	out.Schedule = eventRecord.Awarded
	session.Send(out)
}

func getEventRichmanClubAward(session *net.Session, column, order int8) {
	state := module.State(session)
	eventId := int16(event_dat.EVENT_RICHMAN_CLUB)
	eventInfo, _ := event_dat.GetEventInfoById(eventId)

	out := &event_api.GetRichmanClubAward_Out{}
	if !event_dat.CheckEventTime(eventInfo, event_dat.NOT_END) {
		//活动已结束
		out.Result = 1
		session.Send(out)
		return
	}
	eventRecord := state.EventsState.GetPlayerEventInfoById(eventId)
	// TODO 验证
	// 1.是否已领取
	// 2.VIP等级是否达到
	// 3.目标VIP条件是否已达到
	seq := event_dat.GetRichmanRealSequence(int(column), int(order))
	rpc.RemoteGetGlobalLevelVipCount(func(countByLevel []int32) {

		if (1<<uint(seq-1))&eventRecord.Awarded > 0 {
			//验证是否已领取
			out.Result = 1
			session.Send(out)
			return
		}
		awardsRow := event_dat.GetRichmanAwardBySeq(int(column))
		total := 0
		for i := int(awardsRow.RequireVipLevel); i < len(countByLevel); i++ {
			total += int(countByLevel[i])
		}
		if int(awardsRow.RequireVipCount) > total {
			//验证目标VIP是否已达到
			out.Result = 1
			session.Send(out)
			return
		}
		var needVipLevel, itemId, num int16
		switch order {
		case 1:
			needVipLevel = awardsRow.AwardVipLevel1
			itemId = awardsRow.AwardVipItem1Id
			num = awardsRow.AwardVipItem1Num
		case 2:
			needVipLevel = awardsRow.AwardVipLevel2
			itemId = awardsRow.AwardVipItem2Id
			num = awardsRow.AwardVipItem2Num
		case 3:
			needVipLevel = awardsRow.AwardVipLevel3
			itemId = awardsRow.AwardVipItem3Id
			num = awardsRow.AwardVipItem3Num
		case 4:
			needVipLevel = awardsRow.AwardVipLevel4
			itemId = awardsRow.AwardVipItem4Id
			num = awardsRow.AwardVipItem4Num
		case 5:
			needVipLevel = awardsRow.AwardVipLevel5
			itemId = awardsRow.AwardVipItem5Id
			num = awardsRow.AwardVipItem5Num
		}
		if state.Database.Lookup.PlayerVip(state.PlayerId).Level < needVipLevel {
			//验证玩家是否已达到所需vip等级
			out.Result = 1
			session.Send(out)
			return
		}

		//获取奖励 普通物品
		if eventRecord.LastUpdated == 0 {
			mdb.GlobalExecute(func(globalDB *mdb.Database) {
				globalDB.AgentExecute(state.PlayerId, func(agenDB *mdb.Database) {
					state.EventsState.AddEventAwardState(agenDB, eventId, 0, 1<<uint(seq-1))
				})
			})
		} else {
			mdb.GlobalExecute(func(globalDB *mdb.Database) {
				globalDB.AgentExecute(state.PlayerId, func(agenDB *mdb.Database) {
					state.EventsState.UpdateAwarded(agenDB, eventId, eventRecord.Awarded+(1<<uint(seq-1)))
				})
			})
		}
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(state.PlayerId, func(agenDB *mdb.Database) {
				module.Item.AddItem(agenDB, itemId, num, tlog.IFR_EVENT_CENTER, xdlog.ET_EVENT_CENTER, "")
			})
		})
		session.Send(out)

	})

}
