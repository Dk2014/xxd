// json配置的十连抽活动处理模块
package json_handlers

import (
	"core/net"
	"core/time"
	"fmt"
	"game_server/api/protocol/event_api"
	. "game_server/config"
	"game_server/dat/event_dat"
	"game_server/module"
	"game_server/xdlog"
)

type json_frd_handler struct{}

func init() {
	JsonEventHandlers[event_dat.JSON_EVENT_FIRST_RECHAGE_DAILY] = &json_frd_handler{}
}

func (handler *json_frd_handler) NowStartDays(info *event_dat.JsonEvent) int32 {
	var dt = info.StartUnixTime
	if info.IsRelative { //是相对时间
		dt += ServerCfg.ServerOpenTime
	}

	now := time.GetNowTime()
	index := time.GetNowDayFromUnix(now) - time.GetNowDayFromUnix(dt)
	var timespan = time.GetNowDayFromUnix(info.EndUnixTime) - time.GetNowDayFromUnix(info.StartUnixTime)
	if index > timespan {
		index = event_dat.INVALIDINDEX
	}

	return int32(index + 1)
}

func (handler *json_frd_handler) GetJsonEventStatus(session *net.Session, page int32) event_api.GetEvents_Out_Events {

	state := module.State(session)
	var result event_api.GetEvents_Out_Events

	if jsonEventInfo, ok := event_dat.GetJsonEventInfoById(event_dat.JSON_EVENT_FIRST_RECHAGE_DAILY, page); ok {
		//查询今天是否有对应的奖励
		result = event_api.GetEvents_Out_Events{
			EventId: event_dat.JSON_EVENT_FIRST_RECHAGE_DAILY,
			Page:    page,
		}
		awardIndex := handler.NowStartDays(jsonEventInfo)
		if awardIndex == event_dat.INVALIDINDEX {
			result.EventId = 0 // 无效活动 不添加到返回列表
		} else {
			result.Process = awardIndex
		}

		jsonEventRecord, exists := state.JsonEventsState.GetStatus(event_dat.JSON_EVENT_FIRST_RECHAGE_DAILY, page)

		// 活动记录存在 且 今天充值了 且今天还没领奖 且 奖励奖项有效
		if exists && time.IsToday(jsonEventRecord.LastUpdated) &&
			jsonEventRecord.Awarded != awardIndex &&
			len(jsonEventInfo.List) >= int(awardIndex) {

			result.IsAward = true
		}

	}

	return result
}

func (handler *json_frd_handler) GetJsonEventAward(session *net.Session, page int32, params *event_api.GetEventAward_In) {
	//每日首冲奖励
	state := module.State(session)
	out := &event_api.GetEventAward_Out{}
	out.Result = 3
	jsonEventRecord, exists := state.JsonEventsState.GetStatus(event_dat.JSON_EVENT_FIRST_RECHAGE_DAILY, page)
	if exists && time.IsToday(jsonEventRecord.LastUpdated) {
		if jsonEventInfo, ok := event_dat.GetJsonEventInfoById(event_dat.JSON_EVENT_FIRST_RECHAGE_DAILY, page); ok {

			if jsonEventInfo.CheckStatus(event_dat.NOT_DISPOSE) {
				awardIndex := handler.NowStartDays(jsonEventInfo)
				_, newAward := jsonEventInfo.GetGradeByIndex(awardIndex)

				if jsonEventRecord.Awarded != int32(awardIndex) &&
					awardIndex != event_dat.INVALIDINDEX && newAward != nil {
					//有奖，更新数据库
					jsonEventRecord.Awarded = int32(awardIndex)
					out.Award = jsonEventRecord.Awarded

					state.JsonEventsState.ChangeStatus(event_dat.JSON_EVENT_FIRST_RECHAGE_DAILY, page, jsonEventRecord.Process, jsonEventRecord.Awarded)
					state.JsonEventsState.Save(state.Database)
					//领奖
					addAwardContents(state, newAward.Award, xdlog.ET_EVENT_CENTER_JSON_FIRST_RECHARGE_DAILY, fmt.Sprintf("%d", page))
					out.Result = 0
				}
			}
		}
	}

	session.Send(out)
}
