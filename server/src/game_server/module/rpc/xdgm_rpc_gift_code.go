package rpc

import (
	"game_server/config"
	"game_server/global"
	"game_server/mdb"
	. "game_server/rpc"
)

type Args_XdgmReloadGiftCode struct {
	RPCArgTag
}

type Reply_XdgmReloadGiftCode struct {
}

func (srv *RemoteServe) XdgmReloadGiftCode(args *Args_XdgmReloadGiftCode, reply *Reply_XdgmReloadGiftCode) error {
	return Remote.Serve(mdb.RPC_Remote_XdgmReloadGiftCode, args, mdb.TRANS_TAG_RPC_Serve_XdgmReloadGiftCode, func() error {
		mdb.GlobalExecute(func(gdb *mdb.Database) {
			global.LoadGiftCodeFromDB(gdb, int(config.ServerCfg.ServerId/10), config.GetGiftCodeDBConfig())
		})
		return nil
	})
}
