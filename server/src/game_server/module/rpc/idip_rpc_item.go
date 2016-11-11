package rpc

import (
	"core/fail"
	"fmt"
	"game_server/dat/item_dat"
	"game_server/mdb"
	"game_server/module"
	. "game_server/rpc"
	"game_server/tlog"
)

/*

	发放道具

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

func (this *RemoteServe) IdipSendItem(args *Args_IdipSendItem, reply *Reply_IdipSendItem) error {
	return Remote.Serve(mdb.RPC_Remote_IdipSendItem, args, mdb.TRANS_TAG_RPC_Serve_IdipSendItem, func() error {
		pid, ok := module.Player.GetPlayerByUsername(string(args.OpenId))
		if !ok {
			reply.ErrMsg = "query empty"
			return nil
		}
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(pid, func(db *mdb.Database) {
				itemInfo := item_dat.GetItem(args.ItemId)
				switch itemInfo.TypeId {
				case item_dat.TYPE_CLOTHES:
					fail.When(args.ItemNum > 1, fmt.Sprintf("%d just allow 1 num", itemInfo.Id))
				case item_dat.TYPE_WEAPON:
					fail.When(args.ItemNum > 1, fmt.Sprintf("%d just allow 1 num", itemInfo.Id))
				case item_dat.TYPE_SHOE:
					fail.When(args.ItemNum > 1, fmt.Sprintf("%d just allow 1 num", itemInfo.Id))
				case item_dat.TYPE_FASHION:
					fail.When(args.ItemNum > 1, fmt.Sprintf("%d just allow 1 num", itemInfo.Id))
				case item_dat.TYPE_CHEAT:
					fail.When(args.ItemNum > 1, fmt.Sprintf("%d just allow 1 num", itemInfo.Id))
				case item_dat.TYPE_ACCESSORIES:
					fail.When(args.ItemNum > 1, fmt.Sprintf("%d just allow 1 num", itemInfo.Id))
				case item_dat.TYPE_RESOURCE:
					fail.When(true, "can't send resource")
				default:
					fail.When(args.ItemNum > 999, fmt.Sprintf("%d just allow 999 num", itemInfo.Id))
				}
				module.Item.AddItem(db, itemInfo.Id, args.ItemNum, tlog.IFR_IDIP_RPC_SEND_ITEM, 0, "")
				reply.Result = 0
				reply.RetMsg = "success"
			})
		})
		return nil
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

func (this *RemoteServe) IdipDelItem(args *Args_IdipDelItem, reply *Reply_IdipDelItem) error {
	return Remote.Serve(mdb.RPC_Remote_IdipDelItem, args, mdb.TRANS_TAG_RPC_Serve_IdipDelItem, func() error {
		pid, ok := module.Player.GetPlayerByUsername(string(args.OpenId))
		if !ok {
			reply.ErrMsg = "query empty"
			return nil
		}
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(pid, func(db *mdb.Database) {
				item_info := item_dat.GetItem(args.ItemId)
				module.Item.DelItemByItemId(db, item_info.Id, args.ItemNum, tlog.IFR_IDIP_RPC_DEL_ITEM, 0)
				reply.Result = 0
				reply.RetMsg = "success"
			})
		})
		return nil
	})
}
