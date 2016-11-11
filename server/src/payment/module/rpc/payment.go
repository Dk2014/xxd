package rpc

import (
	"payment/database"
	"rpc_common"
)

type Args_PaymentNotify struct {
	rpc_common.RPCArgTag
	PaymentId int64   // 内部订单ID
	Pid       int64   // 玩家ID
	Money     int64   //增加一级货币数量
	Amount    float64 //充值货币
	ProductId string  //产品ID
	IP        string  //支付IP地址
}

type Reply_PaymentNotify struct {
	//Success   bool
	PaymentId int64 // 返回消息
}

func paymentNotiryCallback(reply *Reply_PaymentNotify) {
	database.DeleteRecordById(reply.PaymentId, database.TABLE_SOHA_DELIVER)
}
