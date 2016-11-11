package main

import (
	"fmt"
	"time"

	"client_test"
	"game_server/api/protocol/player_api"
	"game_server/api/protocol/trader_api"
)

func main() {
	client := client_test.NewClient("127.0.0.1:8080")
	client_test.FromPlatformLogin(client, 1, "trader_state02", "store_state", 2, false)

	client.OutTrader_StoreState = func(out *trader_api.StoreState_Out) {
		fmt.Println("\n==============")
		fmt.Printf("timing:%v\n", time.Now())
		fmt.Printf("trader_api.Store_State:\n%#v\n", out)
		fmt.Println("==============\n")
	}

	client.OutPlayer_FromPlatformLogin = func(out *player_api.FromPlatformLogin_Out) {
		go func() {
			for true {
				client.Trader_StoreState(3)
				time.Sleep(400 * time.Millisecond)
			}
		}()
	}

	b := time.Tick(25 * time.Second)
	<-b
}
