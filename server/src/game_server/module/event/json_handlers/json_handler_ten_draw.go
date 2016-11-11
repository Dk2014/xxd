// json配置的十连抽活动处理模块
package json_handlers

import (
	"core/net"
	"fmt"
	"game_server/api/protocol/event_api"
	"game_server/dat/event_dat"
	"game_server/module"
	"game_server/xdlog"
)

type json_ten_draw_handler struct{}

func init() {
	JsonEventHandlers[event_dat.JSON_EVENT_TEN_DRAW] = &json_ten_draw_handler{}
}

func (handler *json_ten_draw_handler) GetJsonEventStatus(session *net.Session, page int32) event_api.GetEvents_Out_Events {
	state := module.State(session)
	var result event_api.GetEvents_Out_Events
	if jsonEventInfo, ok := event_dat.GetJsonEventInfoById(event_dat.JSON_EVENT_TEN_DRAW, page); ok {
		jsonEventRecord, exists := state.JsonEventsState.GetStatus(event_dat.JSON_EVENT_TEN_DRAW, page)
		result = event_api.GetEvents_Out_Events{
			EventId: event_dat.JSON_EVENT_TEN_DRAW,
			Page:    page,
		}
		if exists {
			result.PlayerProcess = jsonEventRecord.Process
			result.Process = jsonEventRecord.Awarded

			_, nextGrade := jsonEventInfo.GetNextGrade(jsonEventRecord.Awarded)
			if nextGrade == nil {
				result.EventId = 0
			} else {
				if jsonEventRecord.Awarded < jsonEventRecord.Process && jsonEventRecord.Process >= nextGrade.Grade {
					result.IsAward = true
				}
			}
		}
	}
	return result
}

func (handler *json_ten_draw_handler) GetJsonEventAward(session *net.Session, page int32, params *event_api.GetEventAward_In) {
	// 十连抽领奖活动逻辑
	state := module.State(session)
	out := &event_api.GetEventAward_Out{}
	jsonEventRecord, exists := state.JsonEventsState.GetStatus(event_dat.JSON_EVENT_TEN_DRAW, page)
	jsonEventInfo, ok := event_dat.GetJsonEventInfoById(event_dat.JSON_EVENT_TEN_DRAW, page)
	if ok && exists {
		_, jsonAward := jsonEventInfo.GetNextGrade(jsonEventRecord.Awarded)
		if jsonAward != nil && jsonAward.Grade <= jsonEventRecord.Process {
			state.JsonEventsState.ChangeStatus(event_dat.JSON_EVENT_TEN_DRAW, page, jsonEventRecord.Process, jsonAward.Grade)
			state.JsonEventsState.Save(state.Database)
			addAwardContents(state, jsonAward.Award, xdlog.ET_EVENT_CENTER_JSON_TEN_DRAW, fmt.Sprintf("%d", page))
			_, nextAward := jsonEventInfo.GetNextGrade(jsonAward.Grade)
			if nextAward == nil {
				out.Result = 2
			} else {
				out.Award = jsonAward.Grade
				if jsonEventRecord.Process >= nextAward.Grade {
					out.Result = 1
				} else {
					out.Result = 0
				}
			}
		} else {
			out.Result = 3
		}
	} else {
		out.Result = 3
	}
	session.Send(out)
}
