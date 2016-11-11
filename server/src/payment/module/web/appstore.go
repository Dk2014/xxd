package web

import (
	"core/debug"
	"core/fail"
	"core/log"
	"core/time"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/dogenzaka/go-iap/appstore"
	"io/ioutil"
	"net/http"
	"payment/database"
	"payment/module"
	"strconv"
)

const (
	APP_STORE_PAYMENT_STAUTS_OK       = 0 //成功
	APP_STORE_PAYMENT_STAUTS_DUP      = 1 //重复
	APP_STORE_PAYMENT_STAUTS_REJECTED = 2 //苹果验证失败
	APP_STORE_PAYMENT_STAUTS_AGAIN    = 3 //苹果验证发生意外无法完成验证
	APP_STORE_PAYMENT_STAUTS_INVALID  = 4 //非法请求
)

type IAPReceipt struct {
	ProductId      string `json:"product_id"`
	Quantity       int    `json:"quantity"`
	TransactionId  string `json:"transaction_id"`
	PurchaseDateMs string `json:"purchase_date_ms"`
}

type IAPResponse struct {
	Status  int        `json:"status"`
	Receipt IAPReceipt `json:"receipt"`
}

type AppStorePaymentReq struct {
	IP       string              `json:"ip"`
	Receipt  appstore.IAPRequest `json:"receipt"`
	Pid      int64               `json:"pid"`
	Nickname string              `json:"nickname"`
	OpenId   string              `json:"openid"`
}

type AppStorePaymentResp struct {
	Status int8
}

func appStorePaymentHandler(w http.ResponseWriter, req *http.Request) {
	fail.When(req.Body == nil, "request body is nil")
	rawReq, err := ioutil.ReadAll(req.Body)
	fail.When(err != nil, err)

	var paymentData AppStorePaymentReq
	err = json.Unmarshal(rawReq, &paymentData)
	fail.When(err != nil, err)
	paymentLog := &database.AppStorePaymentLog{
		Receipt:      paymentData.Receipt.ReceiptData,
		ReceiptHash:  fmt.Sprintf("%x", md5.Sum([]byte(paymentData.Receipt.ReceiptData))),
		IP:           paymentData.IP,
		GameUserId:   paymentData.Pid,
		OpenId:       paymentData.OpenId,
		Nickname:     paymentData.Nickname,
		Status:       database.PAYMENT_STATUS_CONFIRMING,
		VerifyResult: []byte(""),
		Try:          1,
		TryTimestamp: time.GetNowTime(),
	}
	isDup := database.NewAppStorePaymentLog(paymentLog)
	if isDup {
		log.Debugf("App Store dup payment request receipt hash [%s] game user id [%d]", paymentLog.ReceiptHash, paymentLog.GameUserId)
		jsonResponse(w, AppStorePaymentResp{Status: APP_STORE_PAYMENT_STAUTS_DUP})
		return
	}
	result, rawRsp, err := module.AppStoreVerify.Verify6(paymentLog.Receipt, paymentLog.GameUserId)

	if err != nil {
		//验证过程发生意外
		log.Infof("App Store  payment validation encounter error : id [%d] receipt hash [%s] game user id [%d] err [%v]", paymentLog.Id, paymentLog.ReceiptHash, paymentLog.GameUserId, err)
		jsonResponse(w, AppStorePaymentResp{Status: APP_STORE_PAYMENT_STAUTS_AGAIN})
		return
	}
	log.Debugf("App Store receipt verify result %s\n", debug.Print(0, false, true, "    ", nil, result))
	if result.Status == 0 {
		purchaseDate, err := strconv.ParseInt(result.Receipt.PurchaseDateMs, 10, 64)
		if err != nil {
			log.Warnf("App Store parse purchase date error: id [%d] receipt hash [%s] game user id [%d] err [%v]", paymentLog.Id, paymentLog.ReceiptHash, paymentLog.GameUserId, err)

		}
		paymentLog.PurchaseDateMs = purchaseDate
		paymentLog.TransactionId = result.Receipt.TransactionId
		paymentLog.ProductId = result.Receipt.ProductId
		paymentLog.Quantity = 1
		paymentLog.Status = database.PAYMENT_STATUS_CONFIRMED
		paymentLog.VerifyResult = rawRsp
		database.AppStoreConfirmPayment(paymentLog)
		module.Postman.NewPayment(paymentLog)
		jsonResponse(w, AppStorePaymentResp{Status: APP_STORE_PAYMENT_STAUTS_OK})
	} else {
		paymentLog.VerifyResult = rawRsp
		paymentLog.Status = database.PAYMENT_STATUS_REJECTED
		database.AppStoreConfirmPayment(paymentLog)
		log.Infof("App Store reject payment for [%s] request id [%d] receipt hash [%s] game user id [%d]", appstore.HandleError(result.Status), paymentLog.Id, paymentLog.ReceiptHash, paymentLog.GameUserId)
		jsonResponse(w, AppStorePaymentResp{Status: APP_STORE_PAYMENT_STAUTS_REJECTED})
	}
}
