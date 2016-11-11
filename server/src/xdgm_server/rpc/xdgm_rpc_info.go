package rpc

import (
	"game_server/api/protocol/role_api"
	"game_server/mdb"
)

/*
	玩家当前个人信息
*/
type Args_XdgmGetUserinfo struct {
	RPCArgTag
	Pid int64
}

type MailsInfo struct {
	MailTitle   string // 邮件标题
	MailContent string // 邮件内容
	SendTime    int64  // 邮件发送时间
	ItemDetail  []MailItemxdgm
}

type MailItemxdgm struct {
	ItemId  int16 // 邮件赠送道具ID
	ItemNum int64 // 邮件赠送道具数量
}

type Reply_XdgmGetUserinfo struct {
	RoleName      string      // 角色名称
	Pid           int64       // 玩家ID
	Level         int16       // 当前等级
	Vip           int16       // 当前VIP等级
	Exp           int64       // 当前经验
	Coin          int64       // 当前铜钱
	Ingot         int64       // 当前元宝数量
	Physical      int16       // 当前体力值
	RegisterTime  int64       // 注册时间
	LastLoginTime int64       // 玩家最后登录时间
	Mails         []MailsInfo //邮件信息
	PlayerInfo    []role_api.PlayerInfo_Roles
}

func RemoteXdgmGetUserInfo(pid int64, sid int, callback func(*Reply_XdgmGetUserinfo, error)) {
	reply := &Reply_XdgmGetUserinfo{}
	args := &Args_XdgmGetUserinfo{Pid: pid}
	Remote.Call(sid, mdb.RPC_Remote_XdgmGetUserInfo, args, reply, func(err error) {
		callback(reply, err)
	})
}

/*
	玩家比武场排名
*/
type Args_XdgmGetUserRank struct {
	RPCArgTag
	Pid int64
}

type Reply_XdgmGetUserRank struct {
	Rank int32
}

func RemoteXdgmGetUserRank(pid int64, sid int, callback func(*Reply_XdgmGetUserRank, error)) {
	reply := &Reply_XdgmGetUserRank{}
	args := &Args_XdgmGetUserRank{Pid: pid}
	Remote.Call(sid, mdb.RPC_Remote_XdgmGetUserRank, args, reply, func(err error) {
		callback(reply, err)
	})
}
