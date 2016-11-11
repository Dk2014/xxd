package rpc

import (
	"game_server/mdb"
)

/*

	发放道具请求

*/
type Args_IdipSendItem struct {
	RPCArgTag
	OpenId  string
	RoleId  int64 // 角色ID
	ItemId  int16 // 道具ID
	ItemNum int16 // 数量
}

type Reply_IdipSendItem struct {
	Result int32  // 结果：（0）成功，（其他）失败
	RetMsg string // 返回消息
	ErrMsg string // 错误信息
}

func RemoteIdipSendItem(openId string, sid int, roleId int64, itemId, itemNum int16, callback func(*Reply_IdipSendItem, error)) {
	reply := &Reply_IdipSendItem{}

	args := &Args_IdipSendItem{OpenId: openId, RoleId: roleId, ItemId: itemId, ItemNum: itemNum}
	Remote.Call(sid, mdb.RPC_Remote_IdipSendItem, args, reply, func(err error) {
		callback(reply, err)
	})

}

/*
	删除道具请求
*/
type Args_IdipDelItem struct {
	RPCArgTag
	OpenId  string
	ItemId  int16 // 道具ID
	ItemNum int16 // 数量
}

type Reply_IdipDelItem struct {
	Result int32  // 结果：（0）成功，（其他）失败
	RetMsg string // 返回消息
	ErrMsg string // 错误信息
}

func RemoteIdipDelItem(openId string, sid int, itemId, itemNum int16, callback func(*Reply_IdipDelItem, error)) {
	reply := &Reply_IdipDelItem{}
	args := &Args_IdipDelItem{OpenId: openId, ItemId: itemId, ItemNum: itemNum}
	Remote.Call(sid, mdb.RPC_Remote_IdipDelItem, args, reply, func(err error) {
		callback(reply, err)
	})
}
