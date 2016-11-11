package rpc

import (
	"core/log"
	"game_server/dat/player_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/xdlog"
	"rpc_common"
)

func (this *RemoteServe) WegamesPaymentNotify(args *rpc_common.Args_WegamesPaymentNotify, reply *rpc_common.Reply_WegamesPaymentNotify) error {
	return Remote.Serve(mdb.RPC_Remote_WegamesPaymentNotify, args, mdb.TRANS_TAG_RPC_Serve_WegamesPaymentNotify, func() error {
		panic("废弃")
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(args.Pid, func(db *mdb.Database) {
				log.Infof("PaymentNotify : pid %d, source %d, orderid %d, money %d\n", args.Pid, "wegames platform", args.OrderId, args.Money)
				dup := false
				db.Select.PlayerIapOrder(func(row *mdb.PlayerIapOrderRow) {
					if row.Source() == rpc_common.IAP_SOURCE_WEGAMES_PLATFORM && row.OrderId() == args.OrderId {
						dup = true
						row.Break()
					}
				})
				if !dup {
					db.Insert.PlayerIapOrder(&mdb.PlayerIapOrder{
						Pid:     args.Pid,
						Source:  rpc_common.IAP_SOURCE_WEGAMES_PLATFORM,
						OrderId: args.OrderId,
					})
					money := args.Money
					presentMoney := args.PresentMoney
					// 如果是月卡并且玩家尚不能购买月卡 则转换为
					if args.IsMonthCard {
						if activeMonCard := module.Player.ActiveMonthCard(db); !activeMonCard {
							money = 300
							presentMoney = 25
						} else {
							money = 0
							presentMoney = 0
						}
					}
					if money > 0 {
						module.Player.IncMoney(db, nil, money, player_dat.INGOT, 0 /*无需tlog*/, xdlog.ET_CHARGE, "")
					}
					if presentMoney > 0 {
						module.Player.IncMoney(db, nil, presentMoney, player_dat.INGOT, 0 /*无需tlog*/, xdlog.ET_CHARGE_PRESENT, "")
					}
					playerVIPInfo := db.Lookup.PlayerVip(db.PlayerId())
					module.VIP.UpdateIngot(db, playerVIPInfo.Ingot+args.Money)
					//wegames 平台充值不提供 IP 地址信息
					xdlog.PlayerChargeLog(db, "0.0.0.0.", "", args.PlatformOrderId, "ios", args.TwdMoney, money, "twd")
				}
			})
		})
		return nil
	})
}
