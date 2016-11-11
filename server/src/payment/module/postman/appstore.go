package postman

import (
	"core/log"
	"fmt"
	"payment/database"
	"payment/module"
	"reflect"
	"rpc_common"
	"strconv"
)

//将已确认的订单通知游戏服
type AppStoreDeliverJob struct{}

func (job AppStoreDeliverJob) Process(data interface{}) {
	payment, ok := data.(*database.AppStorePaymentLog)
	if !ok || payment == nil {
		panic(fmt.Errorf("TypeError: except *database.AppStorePendingQueue, have %v", reflect.TypeOf(data)))
	}
	database.ProcessRecord(payment.Id, database.TABLE_WEGAMES_APP_STORE_DELIVER)
	module.RPC.IAPNotify(rpc_common.IAP_SOURCE_APPSTORE, payment.Id, payment.TransactionId, "", payment.GameUserId, payment.ProductId, payment.IP)
}

func (job AppStoreDeliverJob) OnExit() {
}

//重试需要重新发货的订单
type AppStoreRedeliverJob struct{}

func (job AppStoreRedeliverJob) Process(data interface{}) {
	tried_time, ok := data.(int8)
	if !ok {
		panic(fmt.Errorf("TypeError: except int8, have %v", reflect.TypeOf(data)))
		return
	}

	delivering := database.AppStoreFetchDelivering(tried_time, 1000)
	ids := make([]int64, len(delivering))
	for _, pending := range delivering {
		ids = append(ids, pending.Id)
		module.RPC.IAPNotify(rpc_common.IAP_SOURCE_APPSTORE, pending.Id, pending.TransactionId, "", pending.GameUserId, pending.ProductId, pending.IP)
	}
	database.BatchProcessRecord(ids, database.TABLE_WEGAMES_APP_STORE_DELIVER)
}

func (job AppStoreRedeliverJob) OnExit() {
}

//重试在验单过程发生意外的未确认订单
type AppStoreReconfirmJob struct{}

func (job AppStoreReconfirmJob) Process(data interface{}) {
	tried_time, ok := data.(int8)
	if !ok {
		panic(fmt.Errorf("TypeError: except int8, have %v", reflect.TypeOf(data)))
	}

	confirmings := database.AppStoreFetchConfirming(tried_time, 100)
	paymentLogIds := make([]int64, 0, len(confirmings))
	for _, paymentLog := range confirmings {
		paymentLogIds = append(paymentLogIds, paymentLog.Id)

		result, rawRsp, err := module.AppStoreVerify.Verify6(paymentLog.Receipt, paymentLog.GameUserId)
		if err != nil {
			log.Infof("App Store  payment validation encounter error : id [%d] receipt hash [%s] game user id [%d] err [%v]", paymentLog.Id, paymentLog.ReceiptHash, paymentLog.GameUserId, err)
			continue
		}
		if result.Status == 0 {
			purchaseDate, _ := strconv.ParseInt(result.Receipt.PurchaseDateMs, 10, 64)
			paymentLog.PurchaseDateMs = purchaseDate
			paymentLog.TransactionId = result.Receipt.TransactionId
			paymentLog.ProductId = result.Receipt.ProductId
			paymentLog.Quantity = 1
			paymentLog.Status = database.PAYMENT_STATUS_CONFIRMED
			paymentLog.VerifyResult = rawRsp
			database.AppStoreConfirmPayment(paymentLog)
			module.Postman.NewPayment(paymentLog)
		} else {
			paymentLog.VerifyResult = rawRsp
			paymentLog.Status = database.PAYMENT_STATUS_REJECTED
			database.AppStoreConfirmPayment(paymentLog)
		}
	}
	database.BatchProcessRecord(paymentLogIds, database.TABLE_WEGAMES_APP_STORE_LOG)
}

func (job AppStoreReconfirmJob) OnExit() {
}

type WegamePushAppStorePayment struct{}

func (job WegamePushAppStorePayment) Process(data interface{}) {
	deliveredPayment := database.AppStoreFetchDelivered(20)
	log.Debug("[WegamePushAppStorePayment]")
	if len(deliveredPayment) > 0 {
		//delivredPayments = database.AppStoreFetchDelivered(20)
		module.Wegames.PushAppStore(deliveredPayment)
	}
}

func (job WegamePushAppStorePayment) OnExit() {
}
