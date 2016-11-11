package rpc

import (
	"core/time"

	"core/fail"

	"game_server/api/protocol/notify_api"
	"game_server/config"
	"game_server/mdb"
	"game_server/module"
	. "game_server/rpc"
)

type RemoteGlobalAnnouncement int

type Args_CreateAnnouncement struct {
	RPCArgTag
	Id         int32
	TplId      int32
	Parameters string
	Content    string
	Liftime    int64
	SendTime   int64
	ExpireTime int64
	Freq       int32 // 滚动频率
}

type Reply_CreateAnnouncement struct{}

//RPC服务实现。在互动服务器增加一条记录然后发送
func (this *RemoteServe) CreateAnnouncement(args *Args_CreateAnnouncement, reply *Reply_CreateAnnouncement) error {
	return Remote.Serve(mdb.RPC_Remote_CreateAnnouncement, args, mdb.TRANS_TAG_RPC_Serve_CreateAnnouncement, func() error {
		newAnnc := new(mdb.GlobalAnnouncement)
		newAnnc.TplId = args.TplId
		newAnnc.Parameters = args.Parameters
		newAnnc.Content = args.Content
		newAnnc.ExpireTime = args.ExpireTime
		newAnnc.SendTime = args.SendTime
		newAnnc.SpacingTime = int64(args.Freq)

		//互动服务器的数据库操作需要使用全局事务
		mdb.GlobalExecute(func(db *mdb.Database) {
			//插入数据
			db.Insert.GlobalAnnouncement(newAnnc)
		})

		module.PostContent(&module.PostUnit{
			Type:      module.GLOBAL_POST_UNIT_TYPE_ANNOUNCEMENT,
			StartTime: args.SendTime,
			EndTime:   args.ExpireTime,
			Response: &notify_api.SendAnnouncement_Out{
				Id:          newAnnc.Id,
				TplId:       args.TplId,
				ExpireTime:  args.ExpireTime,
				Parameters:  []byte(args.Parameters),
				Content:     []byte(args.Content),
				SpacingTime: args.Freq,
			},
			Id: newAnnc.Id,
		}, args.SendTime-time.GetNowTime())
		return nil
	})
}

//RPC客户端
func RemoteGlobalAnnouncementCreate(tplId int32, paramters string, content string, sendtime int64, expiretime int64, freq int32) {

	if time.GetNowTime() > expiretime {
		return
	}

	args := new(Args_CreateAnnouncement)
	args.TplId = tplId
	args.Parameters = paramters
	args.Content = content
	args.SendTime = sendtime
	args.ExpireTime = expiretime
	args.Freq = freq

	reply := new(Reply_CreateAnnouncement)
	Remote.Call(config.ServerCfg.GlobalServerId, mdb.RPC_Remote_CreateAnnouncement, args, reply, mdb.TRANS_TAG_RPC_Call_CreateAnnouncement, func(err error) {
		fail.When(err != nil, err)
	})
}
