package rpc

import (
	"game_server/mdb"
)

type JsonEventAward struct {
	Grade int32 //奖励所需的档位
	Award *EventDefaultAward
}

// ps: 通常活动的奖励这个结构已满足使用
type EventDefaultAward struct {
	Ingot    int16 // 奖励元宝
	Coin     int32 // 奖励铜钱
	Heart    int16 // 奖励爱心
	Item1Id  int16 // 物品1
	Item1Num int16 // 物品1数量
	Item2Id  int16 // 物品2
	Item2Num int16 // 物品2数量
	Item3Id  int16 // 物品3
	Item3Num int16 // 物品3数量
	Item4Id  int16 // 物品4
	Item4Num int16 // 物品4数量
	Item5Id  int16 // 物品5
	Item5Num int16 // 物品5数量
}

type EventInfo struct {
	EventSign       int16
	Page            int32
	LTitle          string
	RTitle          string
	Content         string
	StartUnixTime   int64
	EndUnixTime     int64
	DisposeUnixTime int64
	IsRelative      bool
	Weight          int16
	Tag             int8
}

type Args_XdgmGetNormalEventInfo struct {
	RPCArgTag
	Limit  int16
	Offset int16
}

type Reply_XdgmGetNormalEventInfo struct {
	Total  int          `json:"total"`
	Events []*EventInfo `json:"events"`
}

func XdgmGetNormalEventInfo(sid int, offset, limit int16, callback func(*Reply_XdgmGetNormalEventInfo, error)) {
	reply := &Reply_XdgmGetNormalEventInfo{}
	args := &Args_XdgmGetNormalEventInfo{
		Offset: offset,
		Limit:  limit,
	}
	Remote.Call(sid, mdb.RPC_Remote_XdgmGetNormalEventInfo, args, reply, func(err error) {
		callback(reply, err)
	})
}

type Args_XdgmGetJsonEventInfo struct {
	RPCArgTag
	Limit  int16
	Offset int16
}

type Reply_XdgmGetJsonEventInfo struct {
	Total  int          `json:"total"`
	Events []*EventInfo `json:"events"`
}

func XdgmGetJsonEventInfo(sid int, offset, limit int16, callback func(*Reply_XdgmGetJsonEventInfo, error)) {
	reply := &Reply_XdgmGetJsonEventInfo{}
	args := &Args_XdgmGetJsonEventInfo{
		Offset: offset,
		Limit:  limit,
	}
	Remote.Call(sid, mdb.RPC_Remote_XdgmGetJsonEventInfo, args, reply, func(err error) {
		callback(reply, err)
	})
}

type Args_XdgmUpdateEventAwards struct {
	RPCArgTag
	EventId   int16
	AwardsRaw string
}
type Reply_XdgmUpdateEventAwards struct {
}

func XdgmUpdateEventAwards(sid int, eventId int16, awardsRaw string, callback func(*Reply_XdgmUpdateEventAwards, error)) {
	reply := &Reply_XdgmUpdateEventAwards{}
	args := &Args_XdgmUpdateEventAwards{
		EventId:   eventId,
		AwardsRaw: awardsRaw,
	}
	Remote.Call(sid, mdb.RPC_Remote_XdgmUpdateEventAwards, args, reply, func(err error) {
		callback(reply, err)
	})
}

type Args_XdgmGetEventAwardInfo struct {
	RPCArgTag
	EventSign int16
	Page      int32
}

type Reply_XdgmGetEventAwardInfo struct {
	Awards []*JsonEventAward
}

func XdgmGetEventAwardInfo(sid int, sign int16, page int32, callback func(*Reply_XdgmGetEventAwardInfo, error)) {
	reply := &Reply_XdgmGetEventAwardInfo{}
	args := &Args_XdgmGetEventAwardInfo{
		EventSign: sign,
		Page:      page,
	}
	Remote.Call(sid, mdb.RPC_Remote_XdgmGetEventAwardInfo, args, reply, func(err error) {
		callback(reply, err)
	})
}

type Args_XdgmUpdateNormalEventInfo struct {
	RPCArgTag
	ServerId      int
	EventsInfoRaw string
}
type Reply_XdgmUpdateNormalEventInfo struct {
	Message string `json:"message"`
}

func XdgmUpdateNormalEventInfo(sid int, eventsRaw string, callback func(*Reply_XdgmUpdateNormalEventInfo, error)) {
	reply := &Reply_XdgmUpdateNormalEventInfo{}
	args := &Args_XdgmUpdateNormalEventInfo{
		ServerId:      sid,
		EventsInfoRaw: eventsRaw,
	}
	Remote.Call(sid, mdb.RPC_Remote_XdgmUpdateNormalEventInfo, args, reply, func(err error) {
		callback(reply, err)
	})
}

type EventAnnounce struct {
	Ltitle        string
	Rtitle        string
	Content       string
	StartUnixTime int64
	EndUnixTime   int64
	Weight        int16
	Tag           int8
	Jump          int8
	IsRelative    bool
}

type Args_XdgmGetTextEvents struct {
	RPCArgTag
	ServerId int
}
type Reply_XdgmGetTextEvents struct {
	Events []*EventAnnounce `json:"events"`
}

func XdgmGetTextEvents(sid int, callback func(*Reply_XdgmGetTextEvents, error)) {
	reply := &Reply_XdgmGetTextEvents{}
	args := &Args_XdgmGetTextEvents{
		ServerId: sid,
	}
	Remote.Call(sid, mdb.RPC_Remote_XdgmGetTextEvents, args, reply, func(err error) {
		callback(reply, err)
	})
}
