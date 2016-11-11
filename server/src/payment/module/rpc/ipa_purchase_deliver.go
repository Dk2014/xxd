package rpc

import (
	"core/fail"
	"core/log"
	"payment/config"
	"payment/database"
	"rpc_common"
)

//in app purchase notify
func (mod RPCMod) IAPNotify(source int8, orderId int64, platformOrderId string, currencyType string, pid int64, productId string, ip string) {
	productCfg, ok := config.CFG.ProductInfo[productId]
	if !ok {
		log.Warnf("unkonw product [%s], paymentId %d gameUserId %d", productId, orderId, pid)
		return
	}
	args := &rpc_common.Args_IAPPurchaseDeliver{
		Source:          source,
		OrderId:         orderId,
		PlatformOrderId: platformOrderId,
		IsMonthCard:     productCfg.IsMonthCard,
		CurrencyType:    currencyType,
		Pid:             pid,
		ProductId:       productId,
		Money:           productCfg.GameMoney,
		PresentMoney:    productCfg.PresentGameMoney,
		//Items
		TwdMoney: productCfg.Cost,
		IP:       ip,
	}
	reply := &rpc_common.Reply_IAPPurchaseDeliver{}
	g_RPC.Call(int(pid>>32), "RemoteServe.IAPPurchaseDeliver", args, reply, func(err error) {
		if err == nil {
			switch reply.Source {
			case rpc_common.IAP_SOURCE_APPSTORE:
				database.BatchDeliverRecord([]int64{reply.OrderId}, database.TABLE_WEGAMES_APP_STORE_DELIVER)
			case rpc_common.IAP_SOURCE_GOOGLE_PLAY:
				database.BatchDeliverRecord([]int64{reply.OrderId}, database.TABLE_WEGAMES_GOOGLE_PLAY_DELIVER)
			default:
				fail.When(true, "undefine source type")
			}
		} else {
			fail.When(true, err)
		}
	})
}

func (mod RPCMod) WegamesPurchaseNotify(orderId int64, platformOrderId string, serverId int, platformUid string, twdMoney float64, money int64, presentMoney int64, isMonthCard bool, item string) {
	var (
		pid int64
		err error
	)

	queryPidArg := &rpc_common.Args_QueryPid{
		OpenId: platformUid,
	}
	queryPidReply := &rpc_common.Reply_QueryPid{}
	//互动服规则 区ID * 10 + 9
	g_RPC.Call(serverId*10+9, "RemoteServe.QueryPid", queryPidArg, queryPidReply, func(err2 error) {
		err = err2
		pid = queryPidReply.Pid
	})
	if err != nil {
		log.Errorf("pid query order id [%d], platformOrderId [%s] serverId [%d] error [%v]", orderId, platformOrderId, serverId, err)
		return
	}
	args := &rpc_common.Args_IAPPurchaseDeliver{
		Source:          rpc_common.IAP_SOURCE_WEGAMES_PLATFORM,
		OrderId:         orderId,
		PlatformOrderId: platformOrderId,
		IsMonthCard:     isMonthCard,
		Pid:             pid,
		TwdMoney:        twdMoney,
		Money:           money,
		PresentMoney:    presentMoney,
		Items:           item,
		//IP: ip,
	}
	reply := &rpc_common.Reply_WegamesPaymentNotify{}
	// pid 高32位为玩家所在服务器的唯一ID
	g_RPC.Call(int(pid>>32), "RemoteServe.IAPPurchaseDeliver", args, reply, func(err error) {
		if err == nil {
			database.DeleteRecordById(reply.OrderId, database.TABLE_WEGAMES_PLATFORM_DELIVER)
		} else {
			fail.When(true, err)
		}
	})
}
