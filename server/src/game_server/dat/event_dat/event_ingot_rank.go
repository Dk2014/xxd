package event_dat

import (
	//	. "game_server/config"
	"core/fail"
)

type EventsIngotRankConfigStruct struct {
	StartUnixTime   int64
	EndUnixTime     int64
	DisposeUnixTime int64
	IsRelative      int8
	Weight          int64
	Type            int8
	Tag             int8
	LTitle          string
	RTitle          string
	Content         string
	List            []string
}

var eventIngotRankConfig *EventsIngotRankConfigStruct

func LoadEventInGotRank(info *EventsIngotRankConfigStruct) {
	if info != nil {
		eventIngotRankConfig = &EventsIngotRankConfigStruct{
			StartUnixTime:   info.StartUnixTime,
			EndUnixTime:     info.EndUnixTime,
			DisposeUnixTime: info.DisposeUnixTime,
			IsRelative:      info.IsRelative,
			Weight:          info.Weight,
			Type:            info.Type,
			Tag:             info.Tag,
			LTitle:          info.LTitle,
			RTitle:          info.RTitle,
			Content:         info.Content,
			List:            info.List,
		}
	}
}

func GetEventIngotRankConfig() *EventsIngotRankConfigStruct {
	if eventIngotRankConfig.StartUnixTime > eventIngotRankConfig.EndUnixTime {
		fail.When(true, "Error:time set error!")
	}
	return eventIngotRankConfig
}
