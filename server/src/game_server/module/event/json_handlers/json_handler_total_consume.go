// json配置的累计消费活动处理模块
package json_handlers

import (
	"core/net"
	"fmt"
	"game_server/api/protocol/event_api"
	"game_server/dat/event_dat"
	"game_server/module"
	"game_server/xdlog"
)

type json_total_consume_handler struct{}

func init() {
	JsonEventHandlers[event_dat.JSON_EVENT_TOTAL_CONSUME] = &json_total_consume_handler{}
}

func (handler *json_total_consume_handler) GetJsonEventStatus(session *net.Session, page int32) event_api.GetEvents_Out_Events {
	state := module.State(session)
	var result event_api.GetEvents_Out_Events
	if jsonEventInfo, ok := event_dat.GetJsonEventInfoById(event_dat.JSON_EVENT_TOTAL_CONSUME, page); ok {
		jsonEventRecord, exists := state.JsonEventsState.GetStatus(event_dat.JSON_EVENT_TOTAL_CONSUME, page)
		if !exists || jsonEventRecord.Awarded < int32(1<<uint(len(jsonEventInfo.List)))-1 {
			result = event_api.GetEvents_Out_Events{
				EventId: event_dat.JSON_EVENT_TOTAL_CONSUME,
				Page:    page,
			}
			if exists {
				result.PlayerProcess = jsonEventRecord.Process
				result.Process = jsonEventRecord.Awarded
				maxAwardIndex, _ := jsonEventInfo.GetMaxCanAwardGrade(jsonEventRecord.Process)
				if jsonEventInfo.List[0].Grade <= jsonEventRecord.Process &&
					jsonEventRecord.Awarded < int32(1<<uint(maxAwardIndex+1))-1 {
					result.IsAward = true
				}
			}
		}
	}
	return result
}

func (handler *json_total_consume_handler) GetJsonEventAward(session *net.Session, page int32, params *event_api.GetEventAward_In) {
	state := module.State(session)
	out := &event_api.GetEventAward_Out{}
	out.Result = 3
	jsonEventRecord, _ := state.JsonEventsState.GetStatus(event_dat.JSON_EVENT_TOTAL_CONSUME, page)
	jsonEventInfo, ok := event_dat.GetJsonEventInfoById(event_dat.JSON_EVENT_TOTAL_CONSUME, page)
	if ok {
		wantAwardIndex := int(params.Param1)
		if wantAwardIndex < len(jsonEventInfo.List) &&
			jsonEventInfo.List[wantAwardIndex].Grade <= jsonEventRecord.Process &&
			jsonEventRecord.Awarded&(1<<uint(wantAwardIndex)) == 0 {

			state.JsonEventsState.ChangeStatus(event_dat.JSON_EVENT_TOTAL_CONSUME, page, jsonEventRecord.Process, jsonEventRecord.Awarded+(1<<uint(wantAwardIndex)))
			state.JsonEventsState.Save(state.Database)
			addAwardContents(state, jsonEventInfo.List[wantAwardIndex].Award, xdlog.ET_EVENT_CENTER_JSON_TOTAL_CONSUME, fmt.Sprintf("%d", page))
			out.Award = jsonEventRecord.Awarded
			if jsonEventRecord.Awarded == 1<<uint(len(jsonEventInfo.List))-1 {
				out.Result = 2
			} else {
				maxAwardIndex, _ := jsonEventInfo.GetMaxCanAwardGrade(jsonEventRecord.Process)
				if jsonEventRecord.Awarded < int32(1<<uint(maxAwardIndex+1))-1 {
					out.Result = 1
				} else {
					out.Result = 0
				}
			}

		}
	}
	session.Send(out)
}
