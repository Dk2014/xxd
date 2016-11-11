package web

import (
	"core/fail"
	"fmt"
	"io"
	"net/http"
	"payment/database"
	"payment/module"
	"strconv"
)

func wegamePlatformPaymentHandler(w http.ResponseWriter, req *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			io.WriteString(w, fmt.Sprintf(`{"status": -1, "msg": "%v"}`, err))
			panic(err)
		}
	}()
	//0. 数据校验
	fail.When(!wegamesValidate(req), "validate failed")

	payment := &database.WegamesPaymentLog{}
	payment.OrderId = req.FormValue("wg_order_id")
	payment.PlatformUid = req.FormValue("wg_platform_uid")
	payment.PayAmount = req.FormValue("wg_pay_amount")
	payment.TwdMoney = req.FormValue("wg_twd_money")
	payment.GameCode = req.FormValue("wg_game_code")
	payment.OtherItem = req.FormValue("wg_other_item")

	serverCode, err := strconv.ParseInt(req.FormValue("wg_server_code"), 10, 64)
	fail.When(err != nil, err)
	payment.ServerCode = int(serverCode)

	gameMoney, err := strconv.ParseInt(req.FormValue("wg_game_money"), 10, 64)
	fail.When(err != nil, err)
	payment.GameMoney = gameMoney

	presentGameMoney, err := strconv.ParseInt(req.FormValue("wg_present_game_money"), 10, 64)
	fail.When(err != nil, err)
	payment.PresentGameMoney = presentGameMoney

	payment.VirtualItems = req.FormValue("wg_virtual_items")
	timestamp, err := strconv.ParseInt(req.FormValue("wg_time"), 10, 64)
	fail.When(err != nil, err)
	payment.Time = timestamp
	payment.Sign = req.FormValue("wg_sign")

	//1. 保存支付信息
	dup := database.WegamesNewPaymentLog(payment)
	if !dup {
		//2. 告知 通知线程
		module.Postman.NewPayment(payment)
		io.WriteString(w, `{"status": 1, "msg": "ok"}`)
		return
	} else {
		io.WriteString(w, `{"status": 2, "msg": "dup"}`)
	}
}

/*
func retryHandler(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.ParseInt(req.FormValue("id"), 10, 64)
	fail.When(err != nil, err)
	pq, ok := database.FetchPendingById(id)
	if ok {
		w.Write([]byte("ok"))
		module.Postman.Retry(pq)
	} else {
		w.Write([]byte(fmt.Sprintf("query fail id [%s]", req.FormValue("id"))))
	}
}
*/
