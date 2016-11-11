package main

import (
	"fmt"
	"time"

	"client_test"
	"game_server/api/protocol/player_api"
	"game_server/api/protocol/notify_api"
)

func main() {
	client := client_test.NewClient("127.0.0.1:8080")
	client.Player_FromPlatformLogin("coins4", "coins4", 1, "", 0, false)

	//注册回调函数
	client.OutPlayer_Info = func(out *player_api.Info_Out) {
		fmt.Println("==============")
		fmt.Printf("player_api.InfoOut:\n%#v\n", out)
		fmt.Println("==============")
	}
	client.OutPlayer_BuyCoins = func(out *player_api.BuyCoins_Out) {
		fmt.Println("==============")
		fmt.Printf("player_api.BuyCoinsOut:\n%#v\n", out)
		fmt.Println("==============")
	}

	client.OutNotify_MoneyChange = func(out *notify_api.MoneyChange_Out) {
		fmt.Println("==============")
		fmt.Printf("notify_api.MoneyChange_Out:\n%#v\n", out)
		fmt.Println("==============")
	}

	//获取用户信息，期待里面会有新的字段
	client.Player_Info()

	////调用调试接口获取设置一定数量的元宝
	client.Debug_SetIngot(1000)

	//再次查看用户信息
	client.Player_Info()

	////用元宝购买铜币
	client.Player_BuyCoins(5)
	client.Player_Info()


	b := time.Tick(5 * time.Second)
	<-b
}
