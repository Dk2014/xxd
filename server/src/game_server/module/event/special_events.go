package event

import (
	"core/time"
	"game_server/api/protocol/event_api"
	"game_server/dat/event_dat"
)

func specials_event_handler(out *event_api.GetEvents_Out) {
	out.Specials = make([]event_api.GetEvents_Out_Specials, 0)
	/*
		"InGotRank": {
			"StartUnixTime": 0,
			"EndUnixTime": 0,
			"DisposeUnixTime": 0,
			"IsRelative": 0,
			"Weight": -1,
			"Type":0,
			"Tag": -1,
			"LTitle": "",
			"RTitle": "",
			"Content": "",
			"List": []
		},
	*/
	event_ingot_rank_info := event_dat.GetEventIngotRankConfig()
	if event_ingot_rank_info != nil {
		if event_ingot_rank_info.StartUnixTime <= time.GetNowTime() && event_ingot_rank_info.EndUnixTime >= time.GetNowTime() {
			out.Specials = append(out.Specials, event_api.GetEvents_Out_Specials{
				Sign: []byte("InGotRank"),
			})
		}
	}
	// TODO add more
}
