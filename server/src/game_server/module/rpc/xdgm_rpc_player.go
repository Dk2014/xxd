package rpc

import (
	"core/fail"
	"core/time"
	"game_server/mdb"
	"game_server/module"
	. "game_server/rpc"
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

func (this *RemoteServe) XdgmLockPlayer(args *Args_XdgmLockPlayer, reply *Reply_XdgmLockPlayer) error {
	return Remote.Serve(mdb.RPC_Remote_XdgmLockPlayer, args, mdb.TRANS_TAG_RPC_Serve_XdgmLockPlayer, func() error {
		fail.When(!mdb.CheckPlayer(args.Pid), "pid does not exists")
		mdb.GlobalExecute(func(globalDb *mdb.Database) {
			globalDb.AgentExecute(args.Pid, func(db *mdb.Database) {
				data := db.Lookup.PlayerState(db.PlayerId())
				nowtime := time.GetNowTime()
				if args.BlockTime == 0 {
					data.BanEndTime = -1
				} else if args.BlockTime == -1 {
					data.BanEndTime = 0
				} else {
					data.BanStartTime = nowtime
					data.BanEndTime = nowtime + args.BlockTime
				}
				db.Update.PlayerState(data)
				if args.BlockTime != 0 {
					p_session, ok := module.Player.GetPlayerOnline(args.Pid)
					if ok {
						p_session.Close()
					}
				}
			})
		})
		return nil
	})
}

/*
	禁言玩家操作
*/
type Args_XdgmGagPlayer struct {
	RPCArgTag
	Pid       int64
	BlockTime int64
}

type Reply_XdgmGagPlayer struct {
}

func (this *RemoteServe) XdgmGagPlayer(args *Args_XdgmGagPlayer, reply *Reply_XdgmGagPlayer) error {
	return Remote.Serve(mdb.RPC_Remote_XdgmGagPlayer, args, mdb.TRANS_TAG_RPC_Serve_XdgmGagPlayer, func() error {
		fail.When(!mdb.CheckPlayer(args.Pid), "pid does not exists")
		mdb.GlobalExecute(func(globalDb *mdb.Database) {
			if args.BlockTime < 0 {
				module.ChatRPC.SetBanState(globalDb, 99999999)
			} else {
				module.ChatRPC.SetBanState(globalDb, args.BlockTime)
			}
		})
		return nil
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

func (this *RemoteServe) XdgmAddVipExp(args *Args_XdgmAddVipExp, reply *Reply_XdgmAddVipExp) error {
	return Remote.Serve(mdb.RPC_Remote_XdgmAddVipExp, args, mdb.TRANS_TAG_RPC_Serve_XdgmAddVipExp, func() error {
		fail.When(!mdb.CheckPlayer(args.Pid), "pid does not exists")
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(args.Pid, func(db *mdb.Database) {
				module.VIP.AddVipExp(db, args.Value)
			})
		})
		return nil
	})
}
