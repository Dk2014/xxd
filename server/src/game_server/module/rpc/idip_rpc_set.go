package rpc

import (
	"game_server/mdb"
	"game_server/module"
	. "game_server/rpc"
	"game_server/tlog"
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

func (this *RemoteServe) IdipSetSoulLevel(args *Args_IdipSetSoulLevel, reply *Reply_IdipSetSoulLevel) error {
	return Remote.Serve(mdb.RPC_Remote_IdipSetSoulLevel, args, mdb.TRANS_TAG_RPC_Serve_IdipSetSoulLevel, func() error {
		pid, ok := module.Player.GetPlayerByUsername(string(args.OpenId))
		if !ok {
			reply.ErrMsg = "query empty"
			return nil
		}
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(pid, func(db *mdb.Database) {
				module.Ghost.SetLevel(db, int64(args.SoulId), int16(args.Value))
				reply.Result = 0
				reply.RetMsg = "success"
			})
		})
		return nil
	})
}

/*
	设置剑心等级
*/
type Args_IdipSetSwordLevel struct {
	RPCArgTag
	OpenId string
	RoleId uint64 // 角色ID
	Pos    uint64 // 剑心位置
	Value  int32  // 等级设置：填1则表示1级，2则表示2级
}

type Reply_IdipSetSwordLevel struct {
	Result int32  // 结果：（0）成功，（其他）失败
	RetMsg string // 返回消息
	ErrMsg string // 错误信息
}

func (this *RemoteServe) IdipSetSwordLevel(args *Args_IdipSetSwordLevel, reply *Reply_IdipSetSwordLevel) error {
	return Remote.Serve(mdb.RPC_Remote_IdipSetSwordLevel, args, mdb.TRANS_TAG_RPC_Serve_IdipSetSwordLevel, func() error {
		pid, ok := module.Player.GetPlayerByUsername(string(args.OpenId))
		if !ok {
			reply.ErrMsg = "query empty"
			return nil
		}
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(pid, func(db *mdb.Database) {
				module.SwordSoul.SetSwordSoulLevel(db, int8(args.RoleId), int8(args.Pos), int8(args.Value))
				reply.Result = 0
				reply.RetMsg = "success"
			})
		})
		return nil
	})
}

/*
	增加vip经验
*/
type Args_IdipAddVipExp struct {
	RPCArgTag
	Data  map[string]interface{} //map 结构体 key=>value
	Cmdid uint32
}

type Reply_IdipAddVipExp struct {
	Result int32  // 结果：（0）成功，（其他）失败
	RetMsg string // 返回消息
	ErrMsg string // 错误信息
}

func (this *RemoteServe) IdipAddVipExp(args *Args_IdipAddVipExp, reply *Reply_IdipAddVipExp) error {
	return Remote.Serve(mdb.RPC_Remote_IdipAddVipExp, args, mdb.TRANS_TAG_RPC_Serve_IdipAddVipExp, func() error {
		pid, ok := module.Player.GetPlayerByUsername(args.Data["OpenId"].(string))
		if !ok {
			reply.ErrMsg = "query empty"
			return nil
		}
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(pid, func(db *mdb.Database) {
				module.VIP.AddVipExp(db, int32(args.Data["Value"].(uint32)))
				reply.Result = 0
				reply.RetMsg = "success"
				db.AddTLog(tlog.NewIDIPFLOW(
					args.Data["AreaId"].(uint32),
					uint32(args.Data["PlatId"].(uint8)),
					args.Data["Partition"].(uint32),
					args.Data["OpenId"].(string),
					0,
					args.Data["Value"].(uint32),
					tlog.IDIP_VIP,
					args.Data["Serial"].(string),
					args.Data["Source"].(uint32),
					args.Cmdid,
				))
			})
		})
		return nil
	})
}
