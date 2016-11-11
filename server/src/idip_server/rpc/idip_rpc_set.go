package rpc

import (
	"game_server/mdb"
)

/*
	设置魂侍等级
*/
type Args_IdipSetSoulLevel struct {
	RPCArgTag
	OpenId string
	SoulId uint64 // 魂侍ID
	Value  int32  // 等级设置：填1则表示1级，2则表示2级
}

type Reply_IdipSetSoulLevel struct {
	Result int32  // 结果：（0）成功，（其他）失败
	RetMsg string // 返回消息
	ErrMsg string // 错误信息
}

func RemoteIdipSetSoulLevel(openId string, sid int, soulid uint64, value int32, callback func(*Reply_IdipSetSoulLevel, error)) {
	reply := &Reply_IdipSetSoulLevel{}
	args := &Args_IdipSetSoulLevel{OpenId: openId, SoulId: soulid, Value: value}
	Remote.Call(sid, mdb.RPC_Remote_IdipSetSoulLevel, args, reply, func(err error) {
		callback(reply, err)
	})
}

/*
	设置剑心等级
*/
type Args_IdipSetSwordLevel struct {
	RPCArgTag
	OpenId string
	RoleId uint64 // 角色ID
	Pos    uint32 // 剑心位置ID
	Value  int32  // 等级设置：填1则表示1级，2则表示2级
}

type Reply_IdipSetSwordLevel struct {
	Result int32  // 结果：（0）成功，（其他）失败
	RetMsg string // 返回消息
	ErrMsg string // 错误信息
}

func RemoteIdipSetSwordLevel(openId string, sid int, roleid uint64, pos uint32, value int32, callback func(*Reply_IdipSetSwordLevel, error)) {
	reply := &Reply_IdipSetSwordLevel{}
	args := &Args_IdipSetSwordLevel{OpenId: openId, RoleId: roleid, Pos: pos, Value: value}
	Remote.Call(sid, mdb.RPC_Remote_IdipSetSwordLevel, args, reply, func(err error) {
		callback(reply, err)
	})
}

/*
	增加vip经验
*/
type Args_IdipAddVipExp struct {
	RPCArgTag
	Data  map[string]interface{}
	Cmdid uint32
}

type Reply_IdipAddVipExp struct {
	Result int32  // 结果：（0）成功，（其他）失败
	RetMsg string // 返回消息
	ErrMsg string // 错误信息
}

func RemoteIdipAddVipExp(data map[string]interface{}, cmdid uint32, sid int, callback func(*Reply_IdipAddVipExp, error)) {
	reply := &Reply_IdipAddVipExp{}
	args := &Args_IdipAddVipExp{Data: data, Cmdid: cmdid}
	Remote.Call(sid, mdb.RPC_Remote_IdipAddVipExp, args, reply, func(err error) {
		callback(reply, err)
	})
}
