package rpc

import (
	"core/fail"
	"game_server/dat/mail_dat"
	"game_server/mdb"
	"game_server/module"
	. "game_server/rpc"
)

type Args_SendHeartToPlatformFriend struct {
	RPCArgTag
	Nickname  string
	FriendPid int64
}

type Reply_SendHeartToPlatformFriend struct {
}

//游戏服作为RPC服务器，为本游戏服的玩家添加爱心受赠邮件
func (main *RemoteServe) SendHeartToPlatformFriend(args *Args_SendHeartToPlatformFriend, reply *Reply_SendHeartToPlatformFriend) error {
	return Remote.Serve(mdb.RPC_Remote_SendHeartToPlatformFriend, args, mdb.TRANS_TAG_RPC_Serve_SendHeartToPlatformFriend, func() error {
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(args.FriendPid, func(db *mdb.Database) {
				module.Mail.SendMail(db, mail_dat.MailHeart{
					Who: args.Nickname,
				})
			})
		})
		return nil
	})

}

//RPC客户端，暴露给游戏服内部使用的API
func RemoteSendHeartToPlatformFriend(pid int64, friendPid int64, nickname string, cb func()) {
	args := &Args_SendHeartToPlatformFriend{
		Nickname:  nickname,
		FriendPid: friendPid,
	}
	serverId, _ := module.GetServerIdWithPlayerId(friendPid)
	Remote.Call(serverId, mdb.RPC_Remote_SendHeartToPlatformFriend, args, &Reply_SendHeartToPlatformFriend{}, mdb.TRANS_TAG_RPC_Call_SendHeartToPlatformFriend, func(err error) {
		fail.When(err != nil, err)
		if cb == nil {
			return
		}
		cb()
	})
}
