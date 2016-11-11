package rpc_common

const (
	IAP_SOURCE_APPSTORE         = 1
	IAP_SOURCE_GOOGLE_PLAY      = 2
	IAP_SOURCE_WEGAMES_PLATFORM = 3
)

type Args_IAPPurchaseDeliver struct {
	RPCArgTag
	Source          int8    //订单来源见 IAP__SOURCE_
	OrderId         int64   //内部订单ID
	PlatformOrderId string  //平台订单ID
	IsMonthCard     bool    //月卡
	CurrencyType    string  //货币类型
	TwdMoney        float64 //等价台币
	Pid             int64   //内部玩家ID
	ProductId       string  //产品ID
	Money           int64   //增加一级货币
	PresentMoney    int64   //赠送一级货币
	Items           string  //额外物品 json 格式  [{"type":"a","id":"001","num":"2"},{"type":"b","id":"002","num":"1"}]
	IP              string
}

type Reply_IAPPurchaseDeliver struct {
	Source  int8
	OrderId int64
}
