package rpc

import (
	"game_server/mdb"
)

type Args_XdgmReloadGiftCode struct {
	RPCArgTag
}

type Reply_XdgmReloadGiftCode struct {
}

func RemoteXdgmReloadGiftCode(sid int, callback func(*Reply_XdgmReloadGiftCode, error)) {
	reply := &Reply_XdgmReloadGiftCode{}
	args := &Args_XdgmReloadGiftCode{}
	Remote.Call(sid, mdb.RPC_Remote_XdgmReloadGiftCode, args, reply, func(err error) {
		callback(reply, err)
	})
}
