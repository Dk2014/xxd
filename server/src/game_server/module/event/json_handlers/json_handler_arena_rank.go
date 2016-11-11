// json配置的十连抽活动处理模块
package json_handlers

import (
	"core/net"
	"fmt"
	"game_server/api/protocol/event_api"
	"game_server/api/protocol/notify_api"
	"game_server/dat/event_dat"
	"game_server/dat/player_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/xdlog"
)

type json_arena_rank_handler struct{}

func init() {
	JsonEventHandlers[event_dat.JSON_EVENT_ARENA_RANK] = &json_arena_rank_handler{}
}

func (handler *json_arena_rank_handler) GetJsonEventStatus(session *net.Session, page int32) event_api.GetEvents_Out_Events {
	state := module.State(session)
	var result event_api.GetEvents_Out_Events

	if jsonEventInfo, ok := event_dat.GetJsonEventInfoById(event_dat.JSON_EVENT_ARENA_RANK, page); ok {
		result = event_api.GetEvents_Out_Events{
			EventId: event_dat.JSON_EVENT_ARENA_RANK,
			Page:    page,
			//PlayerProcess: , 	//排名
			//Process:       0,	//0 or index
		}

		jsonEventRecord, exists := state.JsonEventsState.GetStatus(event_dat.JSON_EVENT_ARENA_RANK, page)

		//一旦登录，就刷新其比武场名次,减少领奖时rpc次数
		if jsonEventInfo.CheckStatus(event_dat.NOT_END) {
			if !exists {
				handler.rankHandler(session, page, false)
			}
		} else { //领奖时间
			if !exists { //没有名次，就rpc上更新到数据库
				handler.rankHandler(session, page, true)
			}

			if jsonEventRecord != nil {
				result.PlayerProcess = jsonEventRecord.Process

				_, jsonEventAward := jsonEventInfo.GetGradeByIndex(int32(jsonEventRecord.Process))
				//不是所有排名都有奖励
				if jsonEventAward != nil {
					//没有领取过
					if jsonEventRecord.Awarded != jsonEventAward.Grade {
						result.IsAward = true
					} else {
						result.EventId = 0
					}
					result.Process = jsonEventAward.Grade //index
				} else { //有排名，没奖 不显示了
					result.EventId = 0
				}
			}
		}
	}

	return result
}

func (handler *json_arena_rank_handler) GetJsonEventAward(session *net.Session, page int32, params *event_api.GetEventAward_In) {
	state := module.State(session)
	out := &event_api.GetEventAward_Out{}
	out.Result = 3

	if jsonEventInfo, ok := event_dat.GetJsonEventInfoById(event_dat.JSON_EVENT_ARENA_RANK, page); ok {
		if jsonEventInfo.CheckStatus(event_dat.NOT_DISPOSE) && module.Player.IsOpenFunc(state.Database, player_dat.FUNC_ARENA) {

			jsonEventRecord, exists := state.JsonEventsState.GetStatus(event_dat.JSON_EVENT_ARENA_RANK, page)
			if exists { //有记录，才有奖励
				_, jsonEventAward := jsonEventInfo.GetGradeByIndex(jsonEventRecord.Process)

				if jsonEventAward != nil {
					//没有领取过
					if jsonEventRecord.Awarded != jsonEventAward.Grade {
						state.JsonEventsState.ChangeStatus(event_dat.JSON_EVENT_ARENA_RANK, page, jsonEventRecord.Process, jsonEventAward.Grade)
						state.JsonEventsState.Save(state.Database)
						addAwardContents(state, jsonEventAward.Award, xdlog.ET_EVENT_CENTER_JSON_ARENA_RANK, fmt.Sprintf("%d", page))
						out.Result = 2
						out.Award = jsonEventAward.Grade
					}
				}
			}
		}
	}

	session.Send(out)
}

func (handler *json_arena_rank_handler) rankHandler(session *net.Session, page int32, send bool) {

	state := module.State(session)
	if module.Player.IsOpenFunc(state.Database, player_dat.FUNC_ARENA) {
		module.Arena.GetPlayerArenaRank(state.Database, func(rank int32) {
			state.JsonEventsState.ChangeStatus(event_dat.JSON_EVENT_ARENA_RANK, page, rank, 0)
			mdb.GlobalExecute(func(globalDB *mdb.Database) {
				globalDB.AgentExecute(state.PlayerId, func(agenDB *mdb.Database) {
					state.JsonEventsState.Save(agenDB)
				})
			})
			if send {
				session.Send(&notify_api.SendEventCenterChange_Out{})
			}
		})
	}
}
