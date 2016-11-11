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

func paymentHandler(w http.ResponseWriter, req *http.Request) {
	//0. 数据校验
	payment := &database.PaymentLog{}
	payment.OrderId = req.FormValue("order_id")
	payment.Amount = req.FormValue("amount")
	payment.PayStatus = req.FormValue("pay_status")
	payment.PayTime = req.FormValue("pay_time")
	payment.UserId = req.FormValue("user_id")
	payment.OrderType = req.FormValue("order_type")
	payment.ProductId = req.FormValue("product_id")
	payment.ProductName = req.FormValue("product_name")
	payment.PrivateData = req.FormValue("pivate_data")
	payment.ChannelNumber = req.FormValue("channel_number")
	payment.Sign = req.FormValue("sign")
	payment.Source = req.FormValue("source")
	payment.EnchancedSign = req.FormValue("enchanced_sign")

	//TODO product count
	//payment.ProductCount = req.FormValue("product_count")

	serverId, err := strconv.ParseInt(req.FormValue("server_id"), 10, 64)
	fail.When(err != nil, err)
	payment.ServerId = serverId

	gameUserId, err := strconv.ParseInt(req.FormValue("game_user_id"), 10, 64)
	fail.When(err != nil, err)
	payment.GameUserId = gameUserId

	fail.When(!validate(req), "validate failed")
	//1. 保存支付信息
	dup := database.NewPaymentLog(payment)
	if !dup {
		//2. 告知 通知线程
		module.Postman.NewPayment(payment)
	}
	io.WriteString(w, "ok")
}
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
