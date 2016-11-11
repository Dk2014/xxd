// json配置的春节活动处理模块
package json_handlers

import (
	"core/net"
	"encoding/json"
	"fmt"
	"game_server/api/protocol/event_api"
	. "game_server/config"
	"game_server/dat/event_dat"
	"game_server/dat/player_dat"
	"game_server/module"
	"game_server/tlog"
	"game_server/xdlog"
	"strconv"
	"strings"
	gotime "time"
)

type json_new_year_handler struct{}

func init() {
	JsonEventHandlers[event_dat.JSON_EVENT_NEW_YEAR] = &json_new_year_handler{}
}

func (handler *json_new_year_handler) GetJsonEventStatus(session *net.Session, page int32) event_api.GetEvents_Out_Events {
	var result event_api.GetEvents_Out_Events
	// 只对新年红包活动进行检测，过期且有没领取的奖励则发送奖励邮件，其他情况不做任何处理
	if eventInfo, exists := event_dat.GetJsonEventInfoById(event_dat.JSON_EVENT_NEW_YEAR, event_dat.EVENT_NEW_YEAR_INVAILD_PAGE); exists {
		if eventInfo.IsRelative || eventInfo.EndUnixTime > ServerCfg.ServerOpenTime { // 红包活动领奖开始后开的区 不下发红包活动
			result = event_api.GetEvents_Out_Events{
				EventId: event_dat.JSON_EVENT_NEW_YEAR,
				Page:    event_dat.EVENT_NEW_YEAR_INVAILD_PAGE,
			}
		}
	}
	return result
}

func (handler *json_new_year_handler) GetJsonEventAward(session *net.Session, page int32, params *event_api.GetEventAward_In) {
	//新年红包
	state := module.State(session)
	out := &event_api.GetEventAward_Out{}
	out.Result = 3
	if eventinfo, exists := event_dat.GetJsonEventInfoById(event_dat.JSON_EVENT_NEW_YEAR, event_dat.EVENT_NEW_YEAR_INVAILD_PAGE); exists && !eventinfo.CheckStatus(event_dat.NOT_END) && eventinfo.CheckStatus(event_dat.NOT_DISPOSE) {
		day_order := params.Param1 // param1用来存放要领红包的天数 从1开始
		var status int32 = 0
		jsonEventRecord, existss := state.JsonEventsState.GetStatus(event_dat.JSON_EVENT_NEW_YEAR, event_dat.EVENT_NEW_YEAR_INVAILD_PAGE)
		if existss {
			status = jsonEventRecord.Awarded
		}
		if status&(1<<uint(day_order-1)) == 0 {
			// 没有领取过
			record := state.Database.Lookup.PlayerNewYearConsumeRecord(state.PlayerId)
			if record != nil {
				realRecord := make(map[string]int)
				json.Unmarshal([]byte(record.ConsumeRecord), &realRecord)
				key := fmt.Sprintf("%d-%d", gotime.Now().Year(), day_order)
				if val, ok := realRecord[key]; ok && val > 0 {
					if val < 10 {
						val = 10
					}
					module.Player.IncMoney(state.Database, state.MoneyState, int64(float32(val)*event_dat.EVENT_NEW_YEAR_RATE), player_dat.INGOT, tlog.MFR_EVENT_CENTER, xdlog.ET_EVENT_CENTER_JSON_NEW_YEAR, fmt.Sprintf("%d", page))
					// 更新领取状态
					state.JsonEventsState.ChangeStatus(event_dat.JSON_EVENT_NEW_YEAR, event_dat.EVENT_NEW_YEAR_INVAILD_PAGE, 0, status+1<<uint(day_order-1))
					state.JsonEventsState.Save(state.Database)
					out.Result = 1
				}
			}
		}
	}

	session.Send(out)
}

func _is_over(process map[string]int, awarded int32) int {
	var result int
	var year = gotime.Now().Year()
	for key, val := range process {
		if val > 0 && val < 10 {
			val = 10
		}
		key_slice := strings.Split(key, "-")
		if key_slice[0] == fmt.Sprintf("%d", year) {
			day_order, _ := strconv.Atoi(key_slice[1])
			if awarded&(1<<uint(day_order-1)) == 0 {
				result += val
			}
		}
	}
	return result
}
