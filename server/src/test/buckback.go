package main

import (
	"fmt"
	"time"

	"client_test"
	"game_server/api/protocol/item_api"
)

func main() {
	client := client_test.NewClient("127.0.0.1:8080")
	client.Player_FromPlatformLogin(1, "buybackbug11", "孔烟真", 2, "", 0, false)

	//注册回调函数
	client.OutItem_GetAllItems = func(out *item_api.GetAllItems_Out) {
		fmt.Println("\n==============")
		fmt.Printf("item_api.GetAllItems_Out:\n%#v\n", out)
		fmt.Println("==============\n")
	}

	client.OutItem_BuyItemBack = func(out *item_api.BuyItemBack_Out) {
		fmt.Println("\n==============")
		fmt.Printf("item_api.BuyItemBack_Out:\n%#v\n", out)
		fmt.Println("==============\n")
	}

	client.Item_GetAllItems()
	client.Item_BuyItemBack(4294967473)
	client.Item_GetAllItems()

	b := time.Tick(5 * time.Second)
	<-b
}
