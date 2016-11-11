package postman

import (
	"core/debug"
	"core/fail"
	"core/log"
	"payment/database"
	"payment/module"
	"reflect"
	"strconv"
)

const (
	POSTMAN_FLAG_SOHA            = 1 << 0
	POSTMAN_FLAG_WEGAME_APPLE    = 1 << 1
	POSTMAN_FLAG_WEGAME_GOOGLE   = 1 << 2
	POSTMAN_FLAG_WEGAME_PLATFORM = 1 << 3
)

type PostmanMod struct{}

func init() {
	module.Postman = PostmanMod{}
}

func (mod PostmanMod) Init(newPaymentChanSize int64, retryConfig map[int8]int64, retryLimit int64, flag uint) {
	fail.When(len(retryConfig) == 0, "retryCofig size zero")
	for retried_time, duration := range retryConfig {
		fail.When(retried_time < 0, "invalid retry time")
		fail.When(duration <= 0, "invalid retry duration")
	}
	g_postman = &Postman{
		retryTickChan:   make(chan int8, 1),
		retryConfig:     retryConfig,
		retryEntryLimit: retryLimit,
	}
	if flag&POSTMAN_FLAG_SOHA != 0 {
		g_postman.sohaNewPaymentChan = make(chan *database.PaymentLog, newPaymentChanSize)
		//是否需要启动时把数据库里面的所有订单都处理一遍
		go g_postman.newPaymentLoop()
		//soha订单重新发货
		go g_postman.retryLoop()
	}

	if flag&POSTMAN_FLAG_WEGAME_APPLE != 0 {
		g_postman.appStoreNewPaymentChan = make(chan interface{}, newPaymentChanSize)
		//app store 订单发货
		go EventDrivenJob("App Store deliver goods", g_postman.appStoreNewPaymentChan, AppStoreDeliverJob{})

		//app sotre 订单发货失败重试
		go RetryCronJob("App Store redilver goods", retryConfig, AppStoreRedeliverJob{})

		//app store 未完成校验的订单从新校验
		go RetryCronJob("App Store reconfirm purchasements", retryConfig, AppStoreReconfirmJob{})

		//发送到 wegames
		go RetryCronJob("App Store push purchasement to wegames", map[int8]int64{0: 10}, WegamePushAppStorePayment{})
	}

	if flag&POSTMAN_FLAG_WEGAME_GOOGLE != 0 {
		g_postman.googlePlayNewPaymentChan = make(chan interface{}, newPaymentChanSize)
		//google play 订单发货
		go EventDrivenJob("Google Play deliver goods", g_postman.googlePlayNewPaymentChan, GooglePlayDeliverJob{})

		//google play 未完成校验的订单从新校验
		//go RetryCronJob("Google reconfirm purchase", retryConfig, GooglePlayReconfirmJob{})

		//google play 发货失败订单重试
		go RetryCronJob("Google Play deliver goods", retryConfig, GooglePlayRedeliverJob{})

		//发送到 wegames
		go RetryCronJob("Google Play push purchasement to wegames", map[int8]int64{0: 10}, WegamePushGooglePayment{})
	}
	if flag&POSTMAN_FLAG_WEGAME_PLATFORM != 0 {
		g_postman.wegamesPlatformPaymentChan = make(chan interface{}, newPaymentChanSize)
		go EventDrivenJob("Wegames Platform deliver goods", g_postman.wegamesPlatformPaymentChan, WegamesPlatformDeliverJob{})
		go RetryCronJob("Wegames Platform redeliver goods", retryConfig, WegamesPlatformRedeliverJob{})
	}

}

func (mod PostmanMod) Retry(record *database.PendingQueue) {
	amount, _ := strconv.ParseFloat(record.Amount, 64)
	module.RPC.PaymentNotify(record.Id, record.GameUserId, record.ProductId, amount, record.IP)
}

func (mod PostmanMod) NewPayment(payment interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Debugf("new payment arg error : arg type %v, arg %s\n", reflect.TypeOf(payment), debug.Print(0, false, true, "    ", nil, payment))
			panic(err)
		}
	}()
	switch payment.(type) {
	case *database.PaymentLog:
		val := payment.(*database.PaymentLog)
		if val == nil {
			panic("nil param")
		}
		select {
		case g_postman.sohaNewPaymentChan <- val:
		default:
			//TODO 来到这里服务器基本要崩掉了 需要做一些日志
			go func() {
				g_postman.sohaNewPaymentChan <- val
			}()
		}
	case *database.AppStorePaymentLog:
		val := payment.(*database.AppStorePaymentLog)
		if val == nil {
			panic("nil param")
		}
		select {
		case g_postman.appStoreNewPaymentChan <- val:
		default:
			go func() {
				g_postman.appStoreNewPaymentChan <- val
			}()
		}
	case *database.GooglePlayPaymentLog:
		val := payment.(*database.GooglePlayPaymentLog)
		if val == nil {
			panic("nil param")
		}
		select {
		case g_postman.googlePlayNewPaymentChan <- val:
		default:
			go func() {
				g_postman.googlePlayNewPaymentChan <- val
			}()
		}
	case *database.WegamesPaymentLog:
		val := payment.(*database.WegamesPaymentLog)
		if val == nil {
			panic("nil param")
		}
		select {
		case g_postman.wegamesPlatformPaymentChan <- val:
		default:
			go func() {
				g_postman.wegamesPlatformPaymentChan <- val
			}()
		}
	default:
		panic("unexcept type")
	}
}

func (mod PostmanMod) Stop() {
	//TODO 是否需要 graceful shutdown?
}
