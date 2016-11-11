package event

import (
	"core/net"
	"core/time"
	"encoding/json"
	"fmt"
	"game_server/api/protocol/notify_api"
	. "game_server/config"
	"game_server/dat/event_dat"
	"game_server/dat/mail_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/module/rpc"
	gotime "time"
)

func init() {
	module.Event = EventMod{}
}

type EventMod struct {
}

func (mod EventMod) LevelUp(db *mdb.Database, level, newLevel int16) {
	var max int16
	eventInfo, _ := event_dat.GetEventInfoById(event_dat.EVENT_LEVEL_AWARD)
	if event_dat.CheckEventTime(eventInfo, event_dat.NOT_END) {
		for i := level + 1; i <= newLevel; i++ {
			_, ok := event_dat.GetEventLevelUpAward(i)
			if ok {
				max = i
			}
		}
		//更新等级活动状态记录
		if max > 0 {
			if session, ok := module.Player.GetPlayerOnline(db.PlayerId()); ok {
				session.Send(&notify_api.SendEventCenterChange_Out{}) //发送活动中心通知
				state := module.State(session)
				state.EventsState.UpdateMax(state.Database, event_dat.EVENT_LEVEL_AWARD, int32(max))
			}
		}
	}
}

func (mod EventMod) UpdatePlayerLoginAwardInfo(db *mdb.Database) {
	//屏蔽七天累计登录奖励
	return
	updateLoginAwardInfo(db)
}

func (mod EventMod) UpdateEventActivityStatus(session *net.Session, decreVal int16) {
	state := module.State(session)
	activit, isRefresh := module.AddPlayerActivity(state.Database, int32(decreVal)) //添加活跃度

	eventInfo, _ := event_dat.GetEventInfoById(event_dat.EVENT_PHYSICAL_AWARDS)
	if event_dat.CheckEventTime(eventInfo, event_dat.NOT_END) {
		state := module.State(session)
		if isRefresh {
			state.EventsState.ClearState(state.Database, event_dat.EVENT_PHYSICAL_AWARDS)
		}
		eventInfo := state.EventsState.GetPlayerEventInfoById(event_dat.EVENT_PHYSICAL_AWARDS)
		nextLevel := event_dat.GetNextPhysical(eventInfo.MaxAward)
		if nextLevel <= activit && nextLevel > 0 { //nextLevel在超过活跃度最大限制时返回始终为0
			state.EventsState.UpdateMax(state.Database, event_dat.EVENT_PHYSICAL_AWARDS, nextLevel)
			session.Send(&notify_api.SendEventCenterChange_Out{})
		}
	}
}

func (mod EventMod) GetVersion() int32 {
	return eventExt.Version
}

func (mod EventMod) GetMonthCardPresentIngot() int64 {
	return int64(eventExt.Config.MonthCardIngot)
}

func (mod EventMod) UpdateEventsIngot(db *mdb.Database, val int64, isIn bool) {
	event_ingot_change(db, val, isIn)
}

func (mod EventMod) GetEventsIngot(db *mdb.Database, isIn bool) int64 {
	return get_event_ingot(db, isIn)
}

func (mod EventMod) UpdateFirstRecharge(db *mdb.Database) {

	eventInfo, _ := event_dat.GetEventInfoById(event_dat.EVENT_FIRST_RECHARGE_DAILY)
	//活动是否过期
	if event_dat.CheckEventTime(eventInfo, event_dat.NOT_END) {

		if session, ok := module.Player.GetPlayerOnline(db.PlayerId()); ok {
			state := module.State(session)
			eventRecord := state.EventsState.GetPlayerEventInfoById(event_dat.EVENT_FIRST_RECHARGE_DAILY)
			//如果今天没有更新过
			if !time.IsToday(eventRecord.LastUpdated) {
				state.EventsState.UpdateAwardedTime(state.Database, event_dat.EVENT_FIRST_RECHARGE_DAILY)
				session.Send(&notify_api.SendEventCenterChange_Out{})
			}
		}
	}
}

func (mod EventMod) UpdateJsonEventTenDraw(session *net.Session) {
	state := module.State(session)
	isChanged := false
	if events, ok := event_dat.JsonEvents[event_dat.JSON_EVENT_TEN_DRAW]; ok && len(events) > 0 {
		for page, eventInfo := range events {
			if eventInfo.CheckStatus(event_dat.NOT_END) {
				status, exists := state.JsonEventsState.GetStatus(event_dat.JSON_EVENT_TEN_DRAW, page)
				if exists {
					state.JsonEventsState.ChangeStatus(event_dat.JSON_EVENT_TEN_DRAW, page, status.Process+1, status.Awarded)
				} else {
					state.JsonEventsState.ChangeStatus(event_dat.JSON_EVENT_TEN_DRAW, page, 1, 0)
				}
				isChanged = true
			}
		}
	}
	if isChanged {
		state.JsonEventsState.Save(state.Database)
		session.Send(&notify_api.SendEventCenterChange_Out{})
	}
}

func (mod EventMod) UpdateJsonEventTotalRecharge(db *mdb.Database, val int32) {
	if session, ok := module.Player.GetPlayerOnline(db.PlayerId()); ok {
		state := module.State(session)
		isChanged := false
		if events, ok := event_dat.JsonEvents[event_dat.JSON_EVENT_TOTAL_RECHARGE]; ok && len(events) > 0 {
			for page, eventInfo := range events {
				if eventInfo.CheckStatus(event_dat.NOT_END) {
					status, exists := state.JsonEventsState.GetStatus(event_dat.JSON_EVENT_TOTAL_RECHARGE, page)
					if exists {
						state.JsonEventsState.ChangeStatus(event_dat.JSON_EVENT_TOTAL_RECHARGE, page, status.Process+val, status.Awarded)
					} else {
						state.JsonEventsState.ChangeStatus(event_dat.JSON_EVENT_TOTAL_RECHARGE, page, val, 0)
					}
					isChanged = true
				}
			}
		}
		if isChanged {
			state.JsonEventsState.Save(state.Database)
			if session, ok := module.Player.GetPlayerOnline(db.PlayerId()); ok {
				session.Send(&notify_api.SendEventCenterChange_Out{})
			}
		}
	}
}

func (mod EventMod) UpdateJsonEventTotalConsume(db *mdb.Database, val int32) {
	if session, ok := module.Player.GetPlayerOnline(db.PlayerId()); ok {
		state := module.State(session)
		isChanged := false
		if events, ok := event_dat.JsonEvents[event_dat.JSON_EVENT_TOTAL_CONSUME]; ok && len(events) > 0 {
			for page, eventInfo := range events {
				if eventInfo.CheckStatus(event_dat.NOT_END) {
					status, exists := state.JsonEventsState.GetStatus(event_dat.JSON_EVENT_TOTAL_CONSUME, page)
					if exists {
						state.JsonEventsState.ChangeStatus(event_dat.JSON_EVENT_TOTAL_CONSUME, page, status.Process+val, status.Awarded)
					} else {
						state.JsonEventsState.ChangeStatus(event_dat.JSON_EVENT_TOTAL_CONSUME, page, val, 0)
					}
					isChanged = true
				}
			}
		}
		// /*if isChanged {
		// 	state.JsonEventsState.Save(state.Database)
		// }
		// */
		// //累计单挑消费
		// if events, ok := event_dat.JsonEvents[event_dat.JSON_EVENT_SINGLE_CONSUME]; ok && len(events) > 0 {
		// 	for page, eventInfo := range events {
		// 		if eventInfo.CheckStatus(event_dat.NOT_END) {
		// 			status, _ := state.JsonEventsState.GetStatus(event_dat.JSON_EVENT_SINGLE_CONSUME, page)
		// 			state.JsonEventsState.ChangeStatus(event_dat.JSON_EVENT_SINGLE_CONSUME, page, status.Process+val, status.Awarded)
		// 			isChanged = true
		// 		}
		// 	}
		// }

		if isChanged {
			if session, ok := module.Player.GetPlayerOnline(db.PlayerId()); ok {
				session.Send(&notify_api.SendEventCenterChange_Out{})
			}
			state.JsonEventsState.Save(state.Database)
		}

	}
}

func (mod EventMod) UpdateJsonEventSingleConsume(db *mdb.Database, val int32) {
	//累计单挑消费
	if session, ok := module.Player.GetPlayerOnline(db.PlayerId()); ok {
		state := module.State(session)
		isChanged := false
		if events, ok := event_dat.JsonEvents[event_dat.JSON_EVENT_SINGLE_CONSUME]; ok && len(events) > 0 {
			for page, eventInfo := range events {
				if eventInfo.CheckStatus(event_dat.NOT_END) {
					status, exists := state.JsonEventsState.GetStatus(event_dat.JSON_EVENT_SINGLE_CONSUME, page)
					if !exists {
						status = &module.JsonEventRecordItem{
							Type: int16(event_dat.JSON_EVENT_SINGLE_CONSUME),
							Page: page,
						}
					}

					state.JsonEventsState.ChangeStatus(event_dat.JSON_EVENT_SINGLE_CONSUME, page, status.Process+val, status.Awarded)
					isChanged = true
				}
			}
		}

		if isChanged {
			if session, ok := module.Player.GetPlayerOnline(db.PlayerId()); ok {
				session.Send(&notify_api.SendEventCenterChange_Out{})
			}
			state.JsonEventsState.Save(state.Database)
		}
	}
}

func (mod EventMod) UpdateJsonFirstRecharge(db *mdb.Database) {

	if session, ok := module.Player.GetPlayerOnline(db.PlayerId()); ok {
		state := module.State(session)
		isChanged := false
		if events, ok := event_dat.JsonEvents[event_dat.JSON_EVENT_FIRST_RECHAGE_DAILY]; ok && len(events) > 0 {
			for page, eventInfo := range events {
				if eventInfo.CheckStatus(event_dat.NOT_END) {
					status, exists := state.JsonEventsState.GetStatus(event_dat.JSON_EVENT_FIRST_RECHAGE_DAILY, page)
					if !exists {
						status = &module.JsonEventRecordItem{
							Type: int16(event_dat.JSON_EVENT_FIRST_RECHAGE_DAILY),
							Page: page,
						}
					}

					if !time.IsToday(status.LastUpdated) { //今天没有更新过才算首冲
						state.JsonEventsState.ChangeStatus(event_dat.JSON_EVENT_FIRST_RECHAGE_DAILY, page, status.Process, status.Awarded)
						isChanged = true
					}
				}
			}
		}
		if isChanged {
			state.JsonEventsState.Save(state.Database)
		}
	}
}

func (mod EventMod) UpdateJsonEventArenaRank(db *mdb.Database, rank int32) {
	if events, ok := event_dat.JsonEvents[event_dat.JSON_EVENT_ARENA_RANK]; ok && len(events) > 0 {
		for page, eventInfo := range events {
			if eventInfo.CheckStatus(event_dat.NOT_END) {
				if session, ok := module.Player.GetPlayerOnline(db.PlayerId()); ok {
					isChanged := false
					state := module.State(session)
					status, exists := state.JsonEventsState.GetStatus(event_dat.JSON_EVENT_ARENA_RANK, page)
					if !exists {
						status = &module.JsonEventRecordItem{
							Type: int16(event_dat.JSON_EVENT_ARENA_RANK),
							Page: page,
						}
					}
					if rank != status.Process {
						state.JsonEventsState.ChangeStatus(event_dat.JSON_EVENT_ARENA_RANK, page, rank, status.Awarded)
						isChanged = true
					}
					if isChanged {
						mdb.GlobalExecute(func(globalDB *mdb.Database) {
							globalDB.AgentExecute(state.PlayerId, func(agentDB *mdb.Database) {
								state.JsonEventsState.Save(agentDB)
							})
						})
					}
				} else { //不在线的情况
					if events, ok := event_dat.JsonEvents[event_dat.JSON_EVENT_ARENA_RANK]; ok && len(events) > 0 {
						newEvent := &module.JsonEventRecord{}
						newEvent.Load(db)
						newEvent.ChangeStatus(event_dat.JSON_EVENT_ARENA_RANK, page, rank, 0)
						mdb.GlobalExecute(func(globalDB *mdb.Database) {
							globalDB.AgentExecute(db.PlayerId(), func(agentDB *mdb.Database) {
								newEvent.Save(agentDB)
							})
						})
					}
				}
			}
		}
	}
}

func (mod EventMod) ChangeNewYearConsumeStatus(db *mdb.Database, ingot int32) {
	// TODO 根据活动时间来确定第几天的消费
	event_info, exists := event_dat.GetJsonEventInfoById(event_dat.JSON_EVENT_NEW_YEAR, event_dat.EVENT_NEW_YEAR_INVAILD_PAGE)
	if exists && event_info.CheckStatus(event_dat.NOT_END) {
		// 春节活动期间的消费
		startTime := event_info.StartUnixTime
		if event_info.IsRelative {
			startTime += ServerCfg.ServerOpenTime
		}
		day_order := time.GetNowDay() - time.GetNowDayFromUnix(startTime) + 1
		key := fmt.Sprintf("%d-%d", gotime.Now().Year(), day_order)
		record := db.Lookup.PlayerNewYearConsumeRecord(db.PlayerId())
		if record != nil {
			//执行更新
			eventRecord := make(map[string]int)
			json.Unmarshal([]byte(record.ConsumeRecord), &eventRecord)
			if _, ok := eventRecord[key]; ok {
				eventRecord[key] += int(ingot)
			} else {
				eventRecord[key] = int(ingot)
			}
			bytes, _ := json.Marshal(eventRecord)
			record.ConsumeRecord = string(bytes)
			db.Update.PlayerNewYearConsumeRecord(record)
		} else {
			//执行插入
			eventRecord := make(map[string]int)
			eventRecord[key] = int(ingot)
			bytes, _ := json.Marshal(eventRecord)
			record = &mdb.PlayerNewYearConsumeRecord{
				Pid: db.PlayerId(),
			}
			record.ConsumeRecord = string(bytes)
			db.Insert.PlayerNewYearConsumeRecord(record)
		}
	}
}

func (mod EventMod) CheckQQVIPGift(db *mdb.Database, status int16) {
	eventInfo, _ := event_dat.GetEventInfoById(event_dat.EVENT_QQVIP_ADDITION)
	if !event_dat.CheckEventTime(eventInfo, event_dat.NOT_END) {
		return
	}
	record := db.Lookup.PlayerQqVipGift(db.PlayerId())
	var mail_qqvip, mail_sqqvip bool
	if status > 0 {
		if record != nil {
			//执行检查
			if (status&(1<<event_dat.SUPER_QQ_VIP)) > 0 && record.Surper == 0 {
				record.Surper = 1
				mail_sqqvip = true
			}
			if (status&(1<<event_dat.QQ_VIP)) > 0 && record.Qqvip == 0 && record.Surper == 0 {
				record.Qqvip = 1
				mail_qqvip = true
			}
			db.Update.PlayerQqVipGift(record)
		} else {
			//需要发放
			record = new(mdb.PlayerQqVipGift)
			record.Pid = db.PlayerId()
			if (status&(1<<event_dat.SUPER_QQ_VIP)) > 0 && record.Surper == 0 {
				record.Surper = 1
				mail_sqqvip = true
			}
			if (status&(1<<event_dat.QQ_VIP)) > 0 && record.Qqvip == 0 {
				record.Qqvip = 1
				mail_qqvip = true
			}
			db.Insert.PlayerQqVipGift(record)
		}

		if mail_sqqvip {
			rpc.RemoteMailSend(db.PlayerId(), mail_dat.MailQQSvip{})
		} else if mail_qqvip {
			rpc.RemoteMailSend(db.PlayerId(), mail_dat.MailQQVip{})
		}
	}
}
