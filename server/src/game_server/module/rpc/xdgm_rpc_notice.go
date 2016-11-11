package rpc

import (
	"game_server/api/protocol/notify_api"
	"game_server/mdb"
	"game_server/module"
	. "game_server/rpc"
)

/*
   更新游戏内走马灯
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

func (this *RemoteServe) XdgmSendGameLamp(args *Args_XdgmSendGameLamp, reply *Reply_XdgmSendGameLamp) error {
	return Remote.Serve(mdb.RPC_Remote_XdgmSendGameLamp, args, mdb.TRANS_TAG_RPC_Serve_XdgmSendGameLamp, func() error {
		RemoteGlobalAnnouncementCreate(0, "", args.Content, args.BeginTime, args.EndTime, args.Interval)
		return nil
	})
}

/*
	查询游戏跑马灯
*/
type XdGameLampNoticeInfo struct {
	BeginTime     int64  // 公告生效时间
	EndTime       int64  // 公告结束时间
	Interval      int32  // 滚动频率
	NoticeId      int64  // 公告ID
	NoticeContent string // 公告内容

}

type Args_XdgmSearchGameLamp struct {
	RPCArgTag
}

type Reply_XdgmSearchGameLamp struct {
	XdGameLampNoticeList_count uint32                 // 走马灯公告信息列表的最大数量
	XdGameLampNoticeList       []XdGameLampNoticeInfo // 走马灯公告信息列表
}

func (this *RemoteServe) XdgmSearchGameLamp(args *Args_XdgmSearchGameLamp, reply *Reply_XdgmSearchGameLamp) error {
	return Remote.Serve(mdb.RPC_Remote_XdgmSearchGameLamp, args, mdb.TRANS_TAG_RPC_Serve_XdgmSearchGameLamp, func() error {
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.Select.GlobalAnnouncement(func(row *mdb.GlobalAnnouncementRow) {
				if row == nil {
					row.Break()
				}
				reply.XdGameLampNoticeList_count++
				rowGO := row.GoObject()
				reply.XdGameLampNoticeList = append(reply.XdGameLampNoticeList, XdGameLampNoticeInfo{
					BeginTime:     rowGO.SendTime,
					EndTime:       rowGO.ExpireTime,
					Interval:      int32(rowGO.SpacingTime),
					NoticeId:      rowGO.Id,
					NoticeContent: rowGO.Content,
				})
			})
		})
		return nil
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

func (this *RemoteServe) XdgmDelGameLamp(args *Args_XdgmDelGameLamp, reply *Reply_XdgmDelGameLamp) error {
	return Remote.Serve(mdb.RPC_Remote_XdgmDelGameLamp, args, mdb.TRANS_TAG_RPC_Serve_XdgmDelGameLamp, func() error {
		//互动服务器的数据库操作需要使用全局事务
		var deleteAnn = new(mdb.GlobalAnnouncement)
		mdb.GlobalExecute(func(db *mdb.Database) {
			//删除数据
			deleteAnn = db.Lookup.GlobalAnnouncement(args.NoticeId)
			if deleteAnn != nil {
				db.Delete.GlobalAnnouncement(deleteAnn)
			}
		})
		module.DeleteContent(&module.PostUnit{
			Type: module.GLOBAL_DELETE_UNIT_TYPE_ANNOUNCEMENT,
			Response: &notify_api.DeleteAnnouncement_Out{
				Id: deleteAnn.Id,
			},
			Id:   deleteAnn.Id,
			Used: true,
		})
		return nil
	})
}
