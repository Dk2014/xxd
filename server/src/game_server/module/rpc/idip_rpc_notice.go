package rpc

import (
	"game_server/api/protocol/notify_api"
	"game_server/mdb"
	"game_server/module"
	. "game_server/rpc"
	"strconv"
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

func (this *RemoteServe) IdipUpdateGameLamp(args *Args_IdipUpdateGameLamp, reply *Reply_IdipUpdateGameLamp) error {
	return Remote.Serve(mdb.RPC_Remote_IdipUpdateGameLamp, args, mdb.TRANS_TAG_RPC_Serve_IdipUpdateGameLamp, func() error {
		RemoteGlobalAnnouncementCreate(0, "", args.LampContent, int64(args.BeginTime), int64(args.EndTime), args.Freq)
		reply.Result = 0
		reply.RetMsg = "success"
		return nil
	})
}

/*
	查询走马灯公告
*/
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

func (this *RemoteServe) IdipGameLampInfo(args *Args_IdipGameLampInfo, reply *Reply_IdipGameLampInfo) error {
	return Remote.Serve(mdb.RPC_Remote_IdipGameLampInfo, args, mdb.TRANS_TAG_RPC_Serve_IdipGameLampInfo, func() error {
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.Select.GlobalAnnouncement(func(row *mdb.GlobalAnnouncementRow) {
				if row == nil {
					row.Break()
				}
				reply.GameLampNoticeList_count++
				rowGO := row.GoObject()
				reply.GameLampNoticeList = append(reply.GameLampNoticeList, SGameLampNoticeInfo{
					BeginTime:     strconv.FormatInt(rowGO.SendTime, 10),
					EndTime:       strconv.FormatInt(rowGO.ExpireTime, 10),
					Freq:          int32(rowGO.SpacingTime),
					NoticeId:      strconv.FormatInt(rowGO.Id, 10),
					NoticeContent: rowGO.Content,
				})
			})
		})
		return nil
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

func (this *RemoteServe) IdipDelNotice(args *Args_IdipDelNotice, reply *Reply_IdipDelNotice) error {
	return Remote.Serve(mdb.RPC_Remote_IdipDelNotice, args, mdb.TRANS_TAG_RPC_Serve_IdipDelNotice, func() error {
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
		reply.Result = 0
		reply.RetMsg = "success"
		return nil
	})
}
