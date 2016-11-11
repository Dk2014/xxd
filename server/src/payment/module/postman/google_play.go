package postman

import (
	//"core/debug"
	"core/log"
	//"encoding/json"
	"fmt"
	"payment/database"
	"payment/module"
	"reflect"
	"rpc_common"
)

//已确认的订单发货
type GooglePlayDeliverJob struct{}

func (job GooglePlayDeliverJob) Process(data interface{}) {
	payment, ok := data.(*database.GooglePlayPaymentLog)
	if !ok || payment == nil {
		panic(fmt.Errorf("TypeError: except *database.GooglePlayPendingQueue, have %v", reflect.TypeOf(data)))
	}
	database.ProcessRecord(payment.Id, database.TABLE_WEGAMES_GOOGLE_PLAY_DELIVER)
	module.RPC.IAPNotify(rpc_common.IAP_SOURCE_GOOGLE_PLAY, payment.Id, payment.TransactionId, "", payment.GameUserId, payment.ProductId, payment.IP)
}

func (job GooglePlayDeliverJob) OnExit() {
}

//重试需要重新发货的订单
type GooglePlayRedeliverJob struct{}

func (job GooglePlayRedeliverJob) Process(data interface{}) {
	tried_time, ok := data.(int8)
	if !ok {
		panic(fmt.Errorf("TypeError: except int8, have %v", reflect.TypeOf(data)))
	}

	delivering := database.GooglePlayFetchDelivering(tried_time, 1000)
	ids := make([]int64, len(delivering))
	for _, pending := range delivering {
		ids = append(ids, pending.Id)
		module.RPC.IAPNotify(rpc_common.IAP_SOURCE_GOOGLE_PLAY, pending.Id, pending.TransactionId, "", pending.GameUserId, pending.ProductId, pending.IP)
	}
	database.BatchProcessRecord(ids, database.TABLE_WEGAMES_GOOGLE_PLAY_DELIVER)
}

func (job GooglePlayRedeliverJob) OnExit() {
}

/*
//重试在验单过程发生意外的未确认订单
type GooglePlayReconfirmJob struct{}

func (job GooglePlayReconfirmJob) Process(data interface{}) {
	tried_time, ok := data.(int8)
	if !ok {
		panic(fmt.Errorf("TypeError: except int8, have %v", reflect.TypeOf(data)))
		return
	}

	confirmings := database.GooglePlayFetchConfirming(tried_time, 100)
	paymentLogIds := make([]int64, 0, len(confirmings))
	for _, paymentLog := range confirmings {
		paymentLogIds = append(paymentLogIds, paymentLog.Id)

		result, err := module.GoogleVerify.VerifyProduct(paymentLog.ProductId, paymentLog.Token)
		if err != nil {
			log.Infof("google play payment validation encounter error : id [%d] token hash [%s] game user id [%d] err [%v]", paymentLog.Id, paymentLog.TokenHash, paymentLog.GameUserId, err)
			continue
		}
		rawRsp, err := json.Marshal(result)
		if err != nil {
			log.Infof("json encode error: result %s", debug.Print(0, false, true, "    ", nil, result))
		}
		if result.PurchaseState == 0 {
			paymentLog.PurchaseDateMs = result.PurchaseTimeMillis
			paymentLog.Quantity = 1
			paymentLog.Status = database.PAYMENT_STATUS_CONFIRMED
			paymentLog.VerifyResult = rawRsp
			database.GooglePlayConfirmPayment(paymentLog)
			module.Postman.NewPayment(paymentLog)
		} else {
			log.Infof("google play reject payment for [%s] request id [%d] token hash [%s] game user id [%d]", result.PurchaseState, paymentLog.Id, paymentLog.TokenHash, paymentLog.GameUserId)
			paymentLog.VerifyResult = rawRsp
			paymentLog.Status = database.PAYMENT_STATUS_REJECTED
			database.GooglePlayConfirmPayment(paymentLog)
		}
	}
	database.BatchProcessRecord(paymentLogIds, database.TABLE_WEGAMES_GOOGLE_PLAY_LOG)
}

func (job GooglePlayReconfirmJob) OnExit() {
}
*/

type WegamePushGooglePayment struct{}

func (job WegamePushGooglePayment) Process(data interface{}) {
	log.Debug("[WegamePushGooglePayment]")
	deliveredPayments := database.GooglePlayFetchDelivered(20)
	if len(deliveredPayments) > 0 {
		module.Wegames.PushGooglePlay(deliveredPayments)
	}
}

func (job WegamePushGooglePayment) OnExit() {
}
