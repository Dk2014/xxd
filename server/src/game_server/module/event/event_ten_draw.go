package event

import (
	"core/net"
	"game_server/api/protocol/event_api"
	"game_server/dat/event_dat"
	"game_server/module"
)

func getTenDrawTimes(session *net.Session) {
	state := module.State(session)
	eventState := state.EventsState.GetPlayerEventInfoById(event_dat.EVENT_TEN_DRAW)
	times := eventState.MaxAward
	session.Send(&event_api.GetEventTenDrawTimes_Out{
		Times: times,
	})
}
