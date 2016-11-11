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

type json_sconsume_handler struct{}

func init() {
	JsonEventHandlers[event_dat.JSON_EVENT_SINGLE_CONSUME] = &json_sconsume_handler{}

}

func (handler *json_sconsume_handler) GetJsonEventStatus(session *net.Session, page int32) event_api.GetEvents_Out_Events {
	state := module.State(session)
	var result event_api.GetEvents_Out_Events

	if jsonEventInfo, ok := event_dat.GetJsonEventInfoById(event_dat.JSON_EVENT_SINGLE_CONSUME, page); ok {
		result = event_api.GetEvents_Out_Events{
			EventId: event_dat.JSON_EVENT_SINGLE_CONSUME,
			Page:    page,
			//PlayerProcess: jsonEventRecord.Process, //元宝消费数
			//Process:       jsonEventRecord.Awarded, //已领取数
		}
		config := jsonEventInfo.List[0] // 始终取第一项作为当前活动的配置
		jsonEventRecord, exists := state.JsonEventsState.GetStatus(event_dat.JSON_EVENT_SINGLE_CONSUME, page)
		if exists {
			result.PlayerProcess = jsonEventRecord.Process
			result.Process = jsonEventRecord.Awarded
			awardCNT := jsonEventRecord.Process/config.Grade - jsonEventRecord.Awarded
			if awardCNT > 0 {
				result.IsAward = true
			}
		}
	}

	return result
}

func (handler *json_sconsume_handler) GetJsonEventAward(session *net.Session, page int32, params *event_api.GetEventAward_In) {
	// 单条消耗领奖
	state := module.State(session)
	out := &event_api.GetEventAward_Out{}
	out.Result = 3
	if jsonEventInfo, ok := event_dat.GetJsonEventInfoById(event_dat.JSON_EVENT_SINGLE_CONSUME, page); ok {
		jsonEventRecord, exists := state.JsonEventsState.GetStatus(event_dat.JSON_EVENT_SINGLE_CONSUME, page)
		if exists {
			config := jsonEventInfo.List[0] // 始终取第一项作为当前活动的配置
			awardCNT := jsonEventRecord.Process/config.Grade - jsonEventRecord.Awarded

			if awardCNT > 0 {
				jsonEventRecord.Awarded += 1
				out.Award = jsonEventRecord.Awarded //给客户端计算
				state.JsonEventsState.ChangeStatus(event_dat.JSON_EVENT_SINGLE_CONSUME, page, jsonEventRecord.Process, jsonEventRecord.Awarded)
				state.JsonEventsState.Save(state.Database)
				addAwardContents(state, config.Award, xdlog.ET_EVENT_CENTER_JSON_SINGLE_CONSUME, fmt.Sprintf("%d", page))

				if awardCNT-1 > 0 { //还可以继续领取
					out.Result = 1
				} else {
					out.Result = 0
				}
			} else { //没有可以领取的了
				out.Result = 0
			}
		}
	}

	session.Send(out)
}
