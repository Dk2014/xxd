package database

const (
	TABLE_SOHA_DELIVER = iota + 1
	TABLE_WEGAMES_APP_STORE_LOG
	TABLE_WEGAMES_APP_STORE_DELIVER
	TABLE_WEGAMES_GOOGLE_PLAY_LOG
	TABLE_WEGAMES_GOOGLE_PLAY_DELIVER
	TABLE_WEGAMES_PLATFORM_DELIVER
)

func TableId2TableName(tableId int) string {
	switch tableId {
	case TABLE_SOHA_DELIVER:
		//越南支付发货队列
		return "pending_queue"
	case TABLE_WEGAMES_APP_STORE_DELIVER:
		//app store 发货队列
		return "app_store_pending_queue"
	case TABLE_WEGAMES_APP_STORE_LOG:
		//app store 日志
		return "app_store_payment_log"
	case TABLE_WEGAMES_GOOGLE_PLAY_LOG:
		//google play 日志
		return "google_play_payment_log"
	case TABLE_WEGAMES_GOOGLE_PLAY_DELIVER:
		//google play 发货队列
		return "google_play_pending_queue"
	case TABLE_WEGAMES_PLATFORM_DELIVER:
		// wegames 平台充值 发货队列
		return "wegames_pending_queue"
	}
	panic("unknow table")
}

const (
	PAYMENT_STATUS_CONFIRMING = 0
	PAYMENT_STATUS_REJECTED   = 1
	PAYMENT_STATUS_CONFIRMED  = 2
)
