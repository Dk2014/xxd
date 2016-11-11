package rpc

import (
	"game_server/mdb"
)

/*
	更新游戏内走马灯
*/
type Args_IdipUpdateGameLamp struct {
	RPCArgTag
	LampContent string // 走马灯内容
	BeginTime   uint64 // 开始时间
	EndTime     uint64 // 结束时间
	Freq        int32  // 滚动频率
}

type Reply_IdipUpdateGameLamp struct {
	Result int32  // 结果：（0）成功，（其他）失败
	RetMsg string // 返回消息
}

func RemoteIdipUpdateGameLamp(sid int, lampContent string, beginTime, endTime uint64, freq int32, callback func(*Reply_IdipUpdateGameLamp, error)) {
	reply := &Reply_IdipUpdateGameLamp{}
	args := &Args_IdipUpdateGameLamp{LampContent: lampContent, BeginTime: beginTime, EndTime: endTime, Freq: freq}
	Remote.Call(sid, mdb.RPC_Remote_IdipUpdateGameLamp, args, reply, func(err error) {
		callback(reply, err)
	})

}

/*
	查询走马灯公告
*/
// 走马灯公告信息对象
type SGameLampNoticeInfo struct {
	BeginTime     string // 公告生效时间
	EndTime       string // 公告结束时间
	Freq          int32  // 滚动频率
	NoticeId      string // 公告ID
	NoticeContent string // 公告内容

}

type Args_IdipGameLampInfo struct {
	RPCArgTag
}

type Reply_IdipGameLampInfo struct {
	GameLampNoticeList_count uint32                // 走马灯公告信息列表的最大数量
	GameLampNoticeList       []SGameLampNoticeInfo // 走马灯公告信息列表
}

func RemoteIdipGameLampInfo(sid int, callback func(*Reply_IdipGameLampInfo, error)) {
	reply := &Reply_IdipGameLampInfo{}
	args := &Args_IdipGameLampInfo{}
	Remote.Call(sid, mdb.RPC_Remote_IdipGameLampInfo, args, reply, func(err error) {
		callback(reply, err)
	})
}

/*
	删除公告
*/
type Args_IdipDelNotice struct {
	RPCArgTag
	NoticeId int64 // 公告ID
}

type Reply_IdipDelNotice struct {
	Result int32  // 结果：（0）成功，（其他）失败
	RetMsg string // 返回消息
}

func RemoteIdipDelNotice(sid int, noticeid int64, callback func(*Reply_IdipDelNotice, error)) {
	reply := &Reply_IdipDelNotice{}
	args := &Args_IdipDelNotice{NoticeId: noticeid}
	Remote.Call(sid, mdb.RPC_Remote_IdipDelNotice, args, reply, func(err error) {
		callback(reply, err)
	})
}
