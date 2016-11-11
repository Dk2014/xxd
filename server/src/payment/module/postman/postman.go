package postman

import (
	"core/log"
	"payment/database"
	"payment/module"
	"strconv"
	"time"
)

var (
	g_postman *Postman
)

type Postman struct {
	sohaNewPaymentChan         chan *database.PaymentLog
	appStoreNewPaymentChan     chan interface{} //*database.AppStorePaymentLog
	googlePlayNewPaymentChan   chan interface{} //*database.GooglePlayPaymentLog
	wegamesPlatformPaymentChan chan interface{} //*database.WegamesPaymentLog
	retryTickChan              chan int8
	retryConfig                map[int8]int64 //tried time -> duration
	retryEntryLimit            int64
}

func (p *Postman) newPaymentLoop() {
LOOP:
	for {
		select {
		case payment, ok := <-p.sohaNewPaymentChan:
			if !ok {
				break LOOP
			}
			//日志处理次数+1
			database.ProcessRecord(payment.Id, database.TABLE_SOHA_DELIVER)
			//通过RPC通知游戏服
			//module.RPC.PaymentNotify(payment)
			amount, _ := strconv.ParseFloat(payment.Amount, 64)
			module.RPC.PaymentNotify(payment.Id, payment.GameUserId, payment.ProductId, amount, payment.PrivateData)
		}
	}
}

//soha订单通知游戏服时发生意外，在此重试
func (p *Postman) retryLoop() {
	for tried_time, duration := range p.retryConfig {
		retryTicker(tried_time, duration, p.retryTickChan)
	}
LOOP:
	for {
		select {
		case tried_time, ok := <-p.retryTickChan:
			if !ok {
				break LOOP
			}
			//如果需要 scale up 考虑 (payment_log.id % total_srv_num) == this_payment_id
			pendings := database.FetchPending(tried_time, 100)
			ids := make([]int64, len(pendings))
			for len(pendings) > 0 {
				ids = ids[:]
				for _, record := range pendings {
					//通过RPC通知游戏服
					amount, _ := strconv.ParseFloat(record.Amount, 64)
					module.RPC.PaymentNotify(record.Id, record.GameUserId, record.ProductId, amount, record.IP)
					ids = append(ids, record.Id)
				}
				database.BatchProcessRecord(ids, database.TABLE_SOHA_DELIVER)
				pendings = database.FetchPending(tried_time, 100)
			}
		}
	}
}

func retryTickerGeneric(tried_time interface{}, duration int64, ch chan interface{}) {
	time.AfterFunc(time.Second*time.Duration(duration), func() {
		select {
		case ch <- tried_time:
		default:
			//如果channel堵塞那么放弃这次 tick
			//TODO 记录日志并且统计
		}
		retryTickerGeneric(tried_time, duration, ch)
	})
}

func retryTicker(tried_time int8, duration int64, ch chan int8) {
	time.AfterFunc(time.Second*time.Duration(duration), func() {
		select {
		case ch <- tried_time:
		default:
			//如果channel堵塞那么放弃这次 tick
			//TODO 记录日志并且统计
		}
		retryTicker(tried_time, duration, ch)
	})
}

type Job interface {
	Process(interface{})
	OnExit()
}

func RetryCronJob(name string, tickerConfig map[int8]int64, job Job) {
	tickerChan := make(chan interface{}, 1)
	for tried_time, duration := range tickerConfig {
		retryTickerGeneric(tried_time, duration, tickerChan)
	}
	EventDrivenJob(name, tickerChan, job)
}

func EventDrivenJob(name string, tickChan chan interface{}, job Job) {
	for {
		select {
		case data, ok := <-tickChan:
			if !ok {
				job.OnExit()
				return
			}
			log.Infof("Running [%s]", name)
			job.Process(data)
		}
	}
}
