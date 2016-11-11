package rpc

import (
	"game_server/mdb"
)

/*
	封玩家操作
*/
type Args_XdgmLockPlayer struct {
	RPCArgTag
	Pid       int64
	BlockTime int64
}

type Reply_XdgmLockPlayer struct {
}

func RemoteXdgmLockPlyaer(pid int64, sid int, blocktime int64, callback func(*Reply_XdgmLockPlayer, error)) {
	reply := &Reply_XdgmLockPlayer{}
	args := &Args_XdgmLockPlayer{
		Pid:       pid,
		BlockTime: blocktime,
	}
	Remote.Call(sid, mdb.RPC_Remote_XdgmLockPlayer, args, reply, func(err error) {
		callback(reply, err)
	})
}

/*
	封玩家操作
*/
type Args_XdgmGagPlayer struct {
	RPCArgTag
	Pid       int64
	BlockTime int64
}

type Reply_XdgmGagPlayer struct {
}

func RemoteXdgmGagPlyaer(pid int64, sid int, blocktime int64, callback func(*Reply_XdgmGagPlayer, error)) {
	reply := &Reply_XdgmGagPlayer{}
	args := &Args_XdgmGagPlayer{
		Pid:       pid,
		BlockTime: blocktime,
	}
	Remote.Call(sid, mdb.RPC_Remote_XdgmGagPlayer, args, reply, func(err error) {
		callback(reply, err)
	})
}

/*
	增加角色VIP经验
*/
type Args_XdgmAddVipExp struct {
	RPCArgTag
	Pid   int64
	Value int32
}

type Reply_XdgmAddVipExp struct {
}

func RemoteXdgmAddVipExp(pid int64, sid int, value int32, callback func(*Reply_XdgmAddVipExp, error)) {
	reply := &Reply_XdgmAddVipExp{}
	args := &Args_XdgmAddVipExp{
		Pid:   pid,
		Value: value,
	}
	Remote.Call(sid, mdb.RPC_Remote_XdgmAddVipExp, args, reply, func(err error) {
		callback(reply, err)
	})
}
