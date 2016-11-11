package rpc

import (
	"core/log"
	"fmt"
	"game_server/api/protocol/notify_api"
	"game_server/dat/player_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/xdlog"
	"rpc_common"
)

func (this *RemoteServe) IAPPurchaseDeliver(args *rpc_common.Args_IAPPurchaseDeliver, reply *rpc_common.Reply_IAPPurchaseDeliver) error {
	return Remote.Serve(mdb.RPC_Remote_IAPPurchaseDeliver, args, mdb.TRANS_TAG_RPC_Serve_IAPPurchaseDeliver, func() error {
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(args.Pid, func(db *mdb.Database) {
				log.Infof("PaymentNotify : pid %d, source %d, orderid %d, money %d\n", args.Pid, args.Source, args.OrderId, args.Money)
				dup := false
				db.Select.PlayerIapOrder(func(row *mdb.PlayerIapOrderRow) {
					if row.Source() == args.Source && row.OrderId() == args.OrderId {
						dup = true
						row.Break()
					}
				})
				reply.Source = args.Source
				reply.OrderId = args.OrderId
				if !dup {
					db.Insert.PlayerIapOrder(&mdb.PlayerIapOrder{
						Pid:     args.Pid,
						Source:  args.Source,
						OrderId: args.OrderId,
					})
					money := args.Money
					presentMoney := args.PresentMoney
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
						playerVIPInfo := db.Lookup.PlayerVip(db.PlayerId())
						if playerVIPInfo.Ingot == 0 {
							if session, ok := module.Player.GetPlayerOnline(db.PlayerId()); ok {
								//台湾版首充赠送额度等于充值赠送的2两倍 否则使用 args.PresentMoney
								presentMoney = money * 2
								session.Send(&notify_api.NotifyFirstRechargeState_Out{
									NeverRecharge: false,
								})
							}
						}
						module.VIP.UpdateIngot(db, playerVIPInfo.Ingot+args.Money)
						if args.Source == rpc_common.IAP_SOURCE_APPSTORE {
							xdlog.PlayerChargeLog(db, args.IP, fmt.Sprintf("%d", args.OrderId), args.PlatformOrderId, "appstore", args.TwdMoney, args.Money, "twd")
						} else if args.Source == rpc_common.IAP_SOURCE_GOOGLE_PLAY {
							xdlog.PlayerChargeLog(db, args.IP, fmt.Sprintf("%d", args.OrderId), args.PlatformOrderId, "googleplay", args.TwdMoney, args.Money, "twd")
						} else if args.Source == rpc_common.IAP_SOURCE_WEGAMES_PLATFORM {
							xdlog.PlayerChargeLog(db, args.IP, fmt.Sprintf("%d", args.OrderId), args.PlatformOrderId, "wegames", args.TwdMoney, args.Money, "twd")
						}
					}
					if presentMoney > 0 {
						module.Player.IncMoney(db, nil, presentMoney, player_dat.INGOT, 0 /*无需tlog*/, xdlog.ET_CHARGE_PRESENT, "")
					}
				}
			})
		})
		return nil
	})
}
