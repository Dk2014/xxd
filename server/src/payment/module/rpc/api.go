package rpc

import (
	//"payment/database"
	"core/log"
	"payment/config"
	"payment/module"
)

var (
	g_RPC *RPC
)

func init() {
	module.RPC = RPCMod{}
}

type RPCMod struct{}

func (mod RPCMod) Init() {
	g_RPC = NewRPC()
	g_RPC.RefreshRPCServerList()
}

func (mod RPCMod) PaymentNotify(paymentId, gameUserId int64, productId string, amount float64, ip string) {
	productCfg, ok := config.CFG.ProductInfo[productId]
	if !ok {
		log.Warnf("unkonw product [%s], paymentId %d gameUserId %d", productId, paymentId, gameUserId)
		return
	}

	reply := &Reply_PaymentNotify{}
	args := &Args_PaymentNotify{
		PaymentId: paymentId,
		Pid:       gameUserId,
		Money:     productCfg.GameMoney,
		Amount:    amount,
		ProductId: productId,
		IP:        ip,
	}
	//TODO 根据 封装一下获取游戏服ID的方法
	g_RPC.Call(int(gameUserId>>32), "RemoteServe.PaymentNotify", args, reply, func(err error) {
		if err == nil {
			paymentNotiryCallback(reply)
		} else {
			log.Errorf("Payment rpc Error:%v", err)
		}
	})
}
