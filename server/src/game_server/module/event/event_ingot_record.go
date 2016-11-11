package event

import (
	"game_server/api/protocol/notify_api"
	"game_server/dat/event_dat"
	"game_server/mdb"
	"game_server/module"
)

func event_ingot_change(db *mdb.Database, val int64, isIn bool) {
	var eventInfo *event_dat.EventCenter
	if isIn {
		eventInfo, _ = event_dat.GetEventInfoById(event_dat.EVENT_TOTAL_RECHARGE)
	} else {
		// TODO 累计消费活动
		eventInfo, _ = event_dat.GetEventInfoById(event_dat.EVENT_TOTAL_CONSUME)
	}
	if !event_dat.CheckEventTime(eventInfo, event_dat.NOT_END) {
		return //不在活动期间
	}

	record := db.Lookup.PlayerEventsIngotRecord(db.PlayerId())
	if record == nil {
		record = &mdb.PlayerEventsIngotRecord{
			Pid: db.PlayerId(),
		}
		db.Insert.PlayerEventsIngotRecord(record)
	}

	var old int64

	if isIn { /* 充值元宝 */
		if eventInfo.End != record.IngotInEndTime {
			record.IngotIn = val
			record.IngotInEndTime = eventInfo.End
		} else {
			old = record.IngotIn
			record.IngotIn += val
		}
		levelVal, _ := event_dat.GetEventTotalRechargeLevel(int32(record.IngotIn))
		if int64(levelVal) > old {
			if session, ok := module.Player.GetPlayerOnline(db.PlayerId()); ok {
				session.Send(&notify_api.SendEventCenterChange_Out{})
			}
		}
	} else { /* 消耗元宝 */
		if eventInfo.End != record.IngotOutEndTime {
			record.IngotOut = val
			record.IngotOutEndTime = eventInfo.End
		} else {
			old = record.IngotOut
			record.IngotOut += val
		}
		if old/int64(event_dat.GetTotalConsumeRadix()) < record.IngotOut/int64(event_dat.GetTotalConsumeRadix()) {
			if session, ok := module.Player.GetPlayerOnline(db.PlayerId()); ok {
				session.Send(&notify_api.SendEventCenterChange_Out{})
			}
		}
	}
	db.Update.PlayerEventsIngotRecord(record)
}

func get_event_ingot(db *mdb.Database, isIn bool) (val int64) {
	record := db.Lookup.PlayerEventsIngotRecord(db.PlayerId())
	if record != nil {
		var eventInfo *event_dat.EventCenter
		if isIn {
			eventInfo, _ = event_dat.GetEventInfoById(event_dat.EVENT_TOTAL_RECHARGE)
			if record.IngotInEndTime != eventInfo.End {
				record.IngotIn = 0
				record.IngotInEndTime = eventInfo.End
				db.Update.PlayerEventsIngotRecord(record)
			} else {
				val = record.IngotIn
			}
		} else {
			eventInfo, _ = event_dat.GetEventInfoById(event_dat.EVENT_TOTAL_CONSUME)
			if record.IngotOutEndTime != eventInfo.End {
				record.IngotOut = 0
				record.IngotOutEndTime = eventInfo.End
				db.Update.PlayerEventsIngotRecord(record)
			} else {
				val = record.IngotOut
			}
		}
	}
	return
}
