package module

import (
	"github.com/dogenzaka/go-iap/appstore"
	//androidpublisher "google.golang.org/api/androidpublisher/v2"
	"payment/config"
	"payment/database"
)

var (
	Web            iWebModule
	Postman        iPostman
	RPC            iRPC
	GoogleVerify   iGoogleVerifyMod
	AppStoreVerify iAppStoreVeirfyMod
	Wegames        iWegamesMod
)

type iWebModule interface {
	Start(addr, command string, tlsCfg config.TlsConfig)
}

type iPostman interface {
	Init(newPaymentChanSize int64, retryConfig map[int8]int64, retryEntryLimit int64, flag uint)
	Retry(record *database.PendingQueue)
	//NewPayment(log *database.PaymentLog)
	NewPayment(log interface{})
	Stop()
}

type iRPC interface {
	Init()
	//any sdk payment notify
	PaymentNotify(paymentId, gameUserId int64, productId string, amount float64, ip string)
	//in app purchase notify
	IAPNotify(source int8, orderId int64, platformOrderId string, currencyType string, pid int64, productId string, ip string)
	//wegames purchase notify
	WegamesPurchaseNotify(orderId int64, platformOrderId string, server_id int, platformUid string, twdMoney float64, money, presentMoney int64, isMonthCard bool, item string)
}

type iGoogleVerifyMod interface {
	//Init(defaultPkg string, timeoutSecond int, googleTokenPath string)
	//VerifyProduct(productId, token string) (*androidpublisher.ProductPurchase, error)
	InitOpenSSL(publicKeyPath string)
	OpenSSLVerify(data string, b64signature string) bool
}

type iAppStoreVeirfyMod interface {
	Init(timeoutSecond int)
	Verify6(receiptStr string, pid int64) (appstore.IAPResponse6, []byte, error)
}

type iWegamesMod interface {
	PushGooglePlay([]*database.GooglePlayPendingQueue)
	PushAppStore([]*database.AppStorePendingQueue)
}
