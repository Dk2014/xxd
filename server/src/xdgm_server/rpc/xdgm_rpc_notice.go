package rpc

import (
	"game_server/mdb"
)

/*
   发送游戏跑马灯
*/
type Args_XdgmSendGameLamp struct {
	RPCArgTag
	Content   string //跑马灯内容
	BeginTime int64  //开始时间
	EndTime   int64  //结束时间
	Interval  int32  //跑马灯间隔
}

type Reply_XdgmSendGameLamp struct {
}

func RemoteXdgmSendGameLamp(sid int, content string, begintime, endtime int64, interval int32, callback func(*Reply_XdgmSendGameLamp, error)) {
	reply := &Reply_XdgmSendGameLamp{}
	args := &Args_XdgmSendGameLamp{
		Content:   content,
		BeginTime: begintime,
		EndTime:   endtime,
		Interval:  interval,
	}
	Remote.Call(sid, mdb.RPC_Remote_XdgmSendGameLamp, args, reply, func(err error) {
		callback(reply, err)
	})
}

/*
	查询游戏跑马灯
*/
type XdGameLampNoticeInfo struct {
	BeginTime     int64  `json:"begin_time"` // 公告生效时间
	EndTime       int64  `json:"end_time"`   // 公告结束时间
	Interval      int32  `json:"interval"`   // 滚动频率
	NoticeId      int64  `json:"notice_id"`  // 公告ID
	NoticeContent string `json:"content"`    // 公告内容

}

type Args_XdgmSearchGameLamp struct {
	RPCArgTag
}

type Reply_XdgmSearchGameLamp struct {
	XdGameLampNoticeList_count uint32                 `json:"count"`      // 走马灯公告信息列表的最大数量
	XdGameLampNoticeList       []XdGameLampNoticeInfo `json:"list"`       // 走马灯公告信息列表
	Channel_id                 int                    `json:"channel_id"` //渠道id
	Server_id                  int                    `json:"server_id"`  //服务器id
}

func RemoteXdgmSearchGameLamp(sid int, callback func(*Reply_XdgmSearchGameLamp, error)) {
	reply := &Reply_XdgmSearchGameLamp{}
	args := &Args_XdgmSearchGameLamp{}
	Remote.Call(sid, mdb.RPC_Remote_XdgmSearchGameLamp, args, reply, func(err error) {
		callback(reply, err)
	})
}

/*
	删除跑马灯
*/
type Args_XdgmDelGameLamp struct {
	RPCArgTag
	NoticeId int64 //跑马灯Id
}

type Reply_XdgmDelGameLamp struct {
}

func RemoteXdgmDelGameLamp(sid int, noticeId int64, callback func(*Reply_XdgmDelGameLamp, error)) {
	reply := &Reply_XdgmDelGameLamp{}
	args := &Args_XdgmDelGameLamp{
		NoticeId: noticeId,
	}
	Remote.Call(sid, mdb.RPC_Remote_XdgmDelGameLamp, args, reply, func(err error) {
		callback(reply, err)
	})
}
