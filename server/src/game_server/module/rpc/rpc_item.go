package rpc

import (
	"game_server/mdb"
	"game_server/module"
	. "game_server/rpc"
	"game_server/tlog"
	"game_server/xdlog"
)

type Args_DeleteItem struct {
	RPCArgTag
	Pid    int64
	ItemId int16
	Num    int16
}

type Reply_DeleteItem struct {
}

//游戏服务器作为RPC服务器，本服玩家通过该服务增加邮件
func (mail *RemoteServe) DeleteItem(args *Args_DeleteItem, reply *Reply_DeleteItem) error {
	return Remote.Serve(mdb.RPC_Remote_DeleteItem, args, mdb.TRANS_TAG_RPC_Serve_DeleteItem, func() error {
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(args.Pid, func(db *mdb.Database) {
				module.Item.DelItemByItemId(db, args.ItemId, args.Num, tlog.IFR_ITEM_ROLE_USE, xdlog.ET_ITEM_ROLE_USE)
			})
		})
		return nil
	})
}

func RemoteDeleteItem(pid int64, itemId int16, num int16) {
	args := &Args_DeleteItem{
		Pid:    pid,
		ItemId: itemId,
		Num:    num,
	}
	serverId, _ := module.GetServerIdWithPlayerId(pid)
	Remote.Call(serverId, mdb.RPC_Remote_DeleteItem, args, &Reply_DeleteItem{}, mdb.TRANS_TAG_RPC_Call_DeleteItem, nil)
}
