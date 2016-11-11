package web

import (
	//"core/debug"
	"core/fail"
	"core/log"
	//"core/time"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"payment/database"
	"payment/module"
	"strconv"
)

/*
{"orderId":"GPA.1393-0207-1293-88807","packageName":"com.wegames.rog.and","productId":"com.wegames.rog.and_004","purchaseTime":1438321346374,"purchaseState":0,"developerPayload":"558435942793220","purchaseToken":"pejjldegfhejheakdocnlbon.AO-J1Ow0l2pi2Sb0qGVPS5KeA0dzZ3uz7EmrdYwJ3Kh3pEhF0223LJeO0bWOr7Kn-6FsLYSuCqyZjZDvQ5Dt5Mt-gN1snz-wkM9hw-rn-dlQQ2Hd5ykjTEfLHFnhWK8wPfAjnePnAe1S"}
*/

type GooglePlayPurchaseData struct {
	OrderId          string `json:"orderId"`
	PackageName      string `json:"packageName"`
	ProductId        string `json:"productId"`
	PurchaseDateMs   int64  `json:"purchaseTime"`
	PurchaseState    int32  `json:"purchaseState"`
	DeveloperPayload string `json:"developerPayload"`
	PurchaseToken    string `json:"purchaseToken"`
}

type GooglePlayPaymentReq struct {
	Nickname   string                 `json:"nickname"`
	OpenId     string                 `json:"openid"`
	IP         string                 `json:"ip"`
	Signature  string                 `json:"signature"`
	SignedData GooglePlayPurchaseData `json:"signed_data"`
}

func googlePlayPaymentHandler(w http.ResponseWriter, req *http.Request) {
	fail.When(req.Body == nil, "request body is nil")
	rawReq, err := ioutil.ReadAll(req.Body)
	fail.When(err != nil, err)

	var paymentData GooglePlayPaymentReq
	err = json.Unmarshal(rawReq, &paymentData)
	if err != nil {
		log.Debugf("googlePlayPaymentHandler: unmaarshal error %v req %s", err, rawReq)
	}
	fail.When(err != nil, err)

	signedData, err := json.Marshal(paymentData.SignedData)
	fail.When(err != nil, err)

	if !module.GoogleVerify.OpenSSLVerify(string(signedData), paymentData.Signature) {
		log.Infof("GoogleVerify fail signedData %s \n%s", signedData, paymentData.SignedData)
		jsonResponse(w, AppStorePaymentResp{Status: APP_STORE_PAYMENT_STAUTS_REJECTED})
		return
	}

	pid, err := strconv.ParseInt(paymentData.SignedData.DeveloperPayload, 10, 64)
	fail.When(err != nil, err)

	paymentLog := &database.GooglePlayPaymentLog{
		Token:         paymentData.SignedData.PurchaseToken,
		TokenHash:     fmt.Sprintf("%x", md5.Sum([]byte(paymentData.SignedData.PurchaseToken))),
		Status:        database.PAYMENT_STATUS_CONFIRMED,
		IP:            paymentData.IP,
		GameUserId:    pid,
		OpenId:        paymentData.OpenId,
		Nickname:      paymentData.Nickname,
		OriginReq:     rawReq,
		TransactionId: paymentData.SignedData.OrderId,
		ProductId:     paymentData.SignedData.ProductId,
	}

	isDup := database.NewGooglePlayPaymentLog(paymentLog)
	if isDup {
		log.Debugf("google play dup payment request token hash [%s] game user id [%d]", paymentLog.TokenHash, paymentLog.GameUserId)
		jsonResponse(w, AppStorePaymentResp{Status: APP_STORE_PAYMENT_STAUTS_DUP})
		return
	}
	module.Postman.NewPayment(paymentLog)
	jsonResponse(w, AppStorePaymentResp{Status: APP_STORE_PAYMENT_STAUTS_OK})
}
