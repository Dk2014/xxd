package rpc

import (
	"core/log"
	"fmt"
	"game_server/dat/payments_rule_dat"
	"game_server/dat/player_dat"
	"game_server/mdb"
	"game_server/module"
	. "game_server/rpc"
	"game_server/xdlog"
)

type Args_PaymentNotify struct {
	RPCArgTag
	PaymentId int64   // 内部订单ID
	Pid       int64   // 玩家ID
	Money     int64   // 增加一级货币数量
	Amount    float64 // 充值金额
	ProductId string  //产品ID
	IP        string  //支付IP地址
}

type Reply_PaymentNotify struct {
	//Success   bool
	PaymentId int64 // 返回消息
}

//
func (this *RemoteServe) PaymentNotify(args *Args_PaymentNotify, reply *Reply_PaymentNotify) error {
	return Remote.Serve(mdb.RPC_Remote_PaymentNotify, args, mdb.TRANS_TAG_RPC_Serve_PaymentNotify, func() error {
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(args.Pid, func(db *mdb.Database) {
				//只有tx渠道需要 moneyState
				log.Infof("PaymentNotify : pid %d, payment_id %d, money %d\n", args.Pid, args.PaymentId, args.Money)
				dup := false
				var present int64
				db.Select.PlayerAnySdkOrder(func(row *mdb.PlayerAnySdkOrderRow) {
					if row.OrderId() == args.PaymentId {
						dup = true
						row.Break()
					}
				})
				if !dup {
					present = payments_rule_dat.GetPresent(args.Money)
					db.Insert.PlayerAnySdkOrder(&mdb.PlayerAnySdkOrder{
						Pid:     args.Pid,
						OrderId: args.PaymentId,
						Present: present,
					})
					module.Player.IncMoney(db, nil, args.Money, player_dat.INGOT, 0 /*无需tlog*/, xdlog.ET_CHARGE, "")
					if present > 0 {
						module.Player.IncMoney(db, nil, present, player_dat.INGOT, 0 /*无需tlog*/, xdlog.ET_CHARGE_PRESENT, "")
					}
					playerVIPInfo := db.Lookup.PlayerVip(args.Pid)
					module.VIP.UpdateIngot(db, playerVIPInfo.Ingot+args.Money)
					playerinfo := db.Lookup.Player(args.Pid)
					var paytype string
					if playerinfo.Cid == 1 {
						paytype = "ios"
					} else if playerinfo.Cid == 2 {
						paytype = "android"
					} else {
						paytype = "unknown"
					}
					//FIXME 越南版专有代码 支付金额 = args.Money*100 vnd
					xdlog.PlayerChargeLog(db, args.IP, fmt.Sprintf("%d", args.PaymentId), "", paytype, float64(args.Money*100), args.Money+present, "vnd")
				}
				reply.PaymentId = args.PaymentId
			})
		})
		return nil
	})
}
