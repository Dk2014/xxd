package rpc

import (
	"game_server/mdb"
)

/*
	增删铜钱
*/
type Args_IdipUpdateMoney struct {
	RPCArgTag
	OpenId string
	RoleId int64 //角色ID
	Value  int64 /* 数量（正加负减） */
	Mtype  int   //货币类型
}

type Reply_IdipUpdateMoney struct {
	Result int32  // 结果：（0）成功，（其他）失败
	RetMsg string // 返回消息
	ErrMsg string // 错误信息
}

func RemoteIdipUpdateMoney(openId string, sid int, roleId, value int64, mtype int, callback func(*Reply_IdipUpdateMoney, error)) {
	reply := &Reply_IdipUpdateMoney{}
	args := &Args_IdipUpdateMoney{OpenId: openId, RoleId: roleId, Value: value, Mtype: mtype}
	Remote.Call(sid, mdb.RPC_Remote_IdipUpdateMoney, args, reply, func(err error) {
		callback(reply, err)
	})
}

/*
	增角色经验
*/
type Args_IdipUpdateExp struct {
	RPCArgTag
	OpenId string
	RoleId int64 //角色ID
	Value  int64 /* 数量（正加负减） */
}

type Reply_IdipUpdateExp struct {
	Result int32  // 结果：（0）成功，（其他）失败
	RetMsg string // 返回消息
	ErrMsg string // 错误信息
}

func RemoteIdipUpdateExp(openId string, sid int, roleId, value int64, callback func(*Reply_IdipUpdateExp, error)) {
	reply := &Reply_IdipUpdateExp{}
	args := &Args_IdipUpdateExp{OpenId: openId, RoleId: roleId, Value: value}
	Remote.Call(sid, mdb.RPC_Remote_IdipUpdateExp, args, reply, func(err error) {
		callback(reply, err)
	})
}

/*
	增加角色体力
*/
type Args_IdipUpdatePhysical struct {
	RPCArgTag
	OpenId string
	RoleId int64 //角色ID
	Value  int64 /* 数量（正加负减） */
}

type Reply_IdipUpdatePhysical struct {
	Result int32  // 结果：（0）成功，（其他）失败
	RetMsg string // 返回消息
	ErrMsg string // 错误信息
}

func RemoteIdipUpdatePhysical(openId string, sid int, roleId, value int64, callback func(*Reply_IdipUpdatePhysical, error)) {
	reply := &Reply_IdipUpdatePhysical{}
	args := &Args_IdipUpdatePhysical{OpenId: openId, RoleId: roleId, Value: value}
	Remote.Call(sid, mdb.RPC_Remote_IdipUpdatePhysical, args, reply, func(err error) {
		callback(reply, err)
	})
}

/*
	增加爱心
*/
type Args_IdipUpdateHeart struct {
	RPCArgTag
	OpenId string
	RoleId int64 //角色ID
	Value  int64 /* 数量（正加负减） */
}

type Reply_IdipUpdateHeart struct {
	Result int32  // 结果：（0）成功，（其他）失败
	RetMsg string // 返回消息
	ErrMsg string // 错误信息
}

func RemoteIdipUpdateHeart(openId string, sid int, roleId, value int64, callback func(*Reply_IdipUpdateHeart, error)) {
	reply := &Reply_IdipUpdateHeart{}
	args := &Args_IdipUpdateHeart{OpenId: openId, RoleId: roleId, Value: value}
	Remote.Call(sid, mdb.RPC_Remote_IdipUpdateHeart, args, reply, func(err error) {
		callback(reply, err)
	})
}

/*
	封号
*/
type Args_IdipBanUser struct {
	RPCArgTag
	OpenId  string // openid
	BanTime uint32 // 封号时长0 永久，**秒）
}

type Reply_IdipBanUser struct {
	Result int32  // 结果：（0）成功，（其他）失败
	RetMsg string // 返回消息
	ErrMsg string // 错误信息
}

func RemoteIdipBanUser(openId string, sid int, banTime uint32, callback func(*Reply_IdipBanUser, error)) {
	reply := &Reply_IdipBanUser{}
	args := &Args_IdipBanUser{OpenId: openId, BanTime: banTime}
	Remote.Call(sid, mdb.RPC_Remote_IdipBanUser, args, reply, func(err error) {
		callback(reply, err)
	})
}

/*
	解封
*/
type Args_IdipUnBanUser struct {
	RPCArgTag
	OpenId string // openid
}

type Reply_IdipUnBanUser struct {
	Result int32  // 结果：（0）成功，（其他）失败
	RetMsg string // 返回消息
	ErrMsg string // 错误信息
}

func RemoteIdipUnBanUser(openId string, sid int, callback func(*Reply_IdipUnBanUser, error)) {
	reply := &Reply_IdipUnBanUser{}
	args := &Args_IdipUnBanUser{OpenId: openId}
	Remote.Call(sid, mdb.RPC_Remote_IdipUnBanUser, args, reply, func(err error) {
		callback(reply, err)
	})
}

/*
	封号(AQ)
*/
type Args_IdipAqBanUser struct {
	RPCArgTag
	OpenId string // openid
	Time   uint32 // 封号时长0 永久，**秒）
}

type Reply_IdipAqBanUser struct {
	Result int32  // 结果：（0）成功，（其他）失败
	RetMsg string // 返回消息
	ErrMsg string // 错误信息
}

func RemoteIdipAqBanUser(openId string, sid int, banTime uint32, callback func(*Reply_IdipAqBanUser, error)) {
	reply := &Reply_IdipAqBanUser{}
	args := &Args_IdipAqBanUser{OpenId: openId, Time: banTime}
	Remote.Call(sid, mdb.RPC_Remote_IdipAqBanUser, args, reply, func(err error) {
		callback(reply, err)
	})
}

/*
	解除惩罚(AQ)
*/
type Args_IdipAqRelievePunish struct {
	RPCArgTag
	OpenId     string // openid
	RelieveBan uint8  // 解除封号（0 否，1 是）
}

type Reply_IdipAqRelievePunish struct {
	Result int32  // 结果：（0）成功，（其他）失败
	RetMsg string // 返回消息
	ErrMsg string // 错误信息
}

func RemoteIdipAqRelievePunish(openid string, sid int, relieveban uint8, callback func(*Reply_IdipAqRelievePunish, error)) {
	reply := &Reply_IdipAqRelievePunish{}
	args := &Args_IdipAqRelievePunish{OpenId: openid, RelieveBan: relieveban}
	Remote.Call(sid, mdb.RPC_Remote_IdipAqRelievePunish, args, reply, func(err error) {
		callback(reply, err)
	})
}

/*
	清零数据
*/
type Args_IdipCleanData struct {
	RPCArgTag
	OpenId        string //openid
	ClearCoin     uint8  //铜钱清零：否（0），是（1）
	ClearPhysical uint8  //体力清零：否（0），是（1）
	ClearHeart    uint8  //爱心清零：否（0），是（1）
}

type Reply_IdipCleanData struct {
	Result int32  // 结果：（0）成功，（其他）失败
	RetMsg string // 返回消息
	ErrMsg string // 错误信息
}

func RemoteIdipCleanData(openid string, sid int, clearcoin, clearphysical, clerheart uint8, callback func(*Reply_IdipCleanData, error)) {
	reply := &Reply_IdipCleanData{}
	args := &Args_IdipCleanData{OpenId: openid, ClearCoin: clearcoin, ClearPhysical: clearphysical, ClearHeart: clerheart}
	Remote.Call(sid, mdb.RPC_Remote_IdipCleanData, args, reply, func(err error) {
		callback(reply, err)
	})
}

/*
	设置宠物激活请求
*/
type Args_IdipAddPet struct {
	RPCArgTag
	OpenId string //openid
	PetId  int16  //宠物ID
}

type Reply_IdipAddPet struct {
	Result int32  // 结果：（0）成功，（其他）失败
	RetMsg string // 返回消息
	ErrMsg string // 错误信息
}

func RemoteIdipAddPet(openid string, sid int, petid int16, callback func(*Reply_IdipAddPet, error)) {
	reply := &Reply_IdipAddPet{}
	args := &Args_IdipAddPet{OpenId: openid, PetId: petid}
	Remote.Call(sid, mdb.RPC_Remote_IdipAddPet, args, reply, func(err error) {
		callback(reply, err)
	})
}
