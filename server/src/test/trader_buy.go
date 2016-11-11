package main

import (
	"fmt"
	"time"

	"client_test"
	"game_server/api/protocol/notify_api"
	"game_server/api/protocol/trader_api"
)

func main() {
	client := client_test.NewClient("127.0.0.1:8080")
	client.Player_FromPlatformLogin(1, "trader06", "trader06", 2, "", 0, false)

	//注册回调函数
	client.OutNotify_ItemChange = func(out *notify_api.ItemChange_Out) {
		fmt.Println("\n==============")
		fmt.Printf("notify_api.ItemChange:\n%#v\n", out)
		fmt.Println("==============\n")
	}

	client.OutTrader_Buy = func(out *trader_api.Buy_Out) {
		fmt.Println("\n==============")
		fmt.Printf("trader_api.Buy:\n%#v\n", out)
		fmt.Println("==============\n")
	}

	client.Debug_SetIngot(1000)
	client.Debug_SetCoins(100000)

	client.Trader_Buy(1)

	b := time.Tick(5 * time.Second)
	<-b
}
