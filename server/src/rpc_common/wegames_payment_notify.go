package rpc_common

type Args_WegamesPaymentNotify struct {
	RPCArgTag
	OrderId         int64   //内部订单ID
	PlatformOrderId string  //平台订单ID
	IsMonthCard     bool    //月卡
	TwdMoney        float64 //等价台币
	Pid             int64   //内部玩家ID
	Money           int64   //增加一级货币
	PresentMoney    int64   //赠送一级货币
	Items           string  //额外物品 json 格式  [{"type":"a","id":"001","num":"2"},{"type":"b","id":"002","num":"1"}]
}

type Reply_WegamesPaymentNotify struct {
	OrderId int64
}
