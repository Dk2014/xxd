package postman

import (
	"core/log"
	"fmt"
	"payment/database"
	"payment/module"
	"reflect"
	"strconv"
)

type WegamesPlatformDeliverJob struct{}

func (job WegamesPlatformDeliverJob) Process(data interface{}) {
	log.Info("new wegames platform deliver job")
	payment, ok := data.(*database.WegamesPaymentLog)
	if payment == nil || !ok {
		panic(fmt.Errorf("TypeError: except *database.WegamesPaymentLog, have %v", reflect.TypeOf(data)))
		return
	}
	database.ProcessRecord(payment.Id, database.TABLE_WEGAMES_PLATFORM_DELIVER)
	twdMoney, err := strconv.ParseFloat(payment.TwdMoney, 64)
	if err != nil {
		log.Errorf("WegamesPlatformDeliverJob parase twd_money error [%v] id [%d] value [%v]\n", err, payment.Id, payment.TwdMoney)
		return
	}

	module.RPC.WegamesPurchaseNotify(payment.Id, payment.OrderId, payment.ServerCode, payment.PlatformUid, twdMoney, payment.GameMoney, payment.PresentGameMoney, payment.OtherItem == "month", payment.VirtualItems)
}

func (job WegamesPlatformDeliverJob) OnExit() {
}

type WegamesPlatformRedeliverJob struct{}

func (job WegamesPlatformRedeliverJob) Process(data interface{}) {
	log.Info("wegames platform redeliver job")
	triedTime, ok := data.(int8)
	if !ok {
		panic(fmt.Errorf("TypeError: except int8, have %v", reflect.TypeOf(data)))
	}
	pendings := database.WegamesFetchPending(triedTime, 100)
	ids := make([]int64, len(pendings))
	for len(pendings) > 0 {
		ids = ids[:]
		for _, record := range pendings {
			ids = append(ids, record.Id)
			twdMoney, err := strconv.ParseFloat(record.TwdMoney, 64)
			if err != nil {
				log.Errorf("WegamesPlatformRedeliverJob parase twd_money error [%v] id [%d]value [%v]\n", err, record.Id, record.TwdMoney)
				continue
			}
			module.RPC.WegamesPurchaseNotify(record.Id, record.OrderId, record.ServerCode, record.PlatformUid, twdMoney, record.GameMoney, record.PresentGameMoney, record.OtherItem == "month", record.VirtualItems)
		}
		pendings = database.WegamesFetchPending(triedTime, 100)
		database.BatchProcessRecord(ids, database.TABLE_WEGAMES_PLATFORM_DELIVER)
	}

}

func (job WegamesPlatformRedeliverJob) OnExit() {
}
