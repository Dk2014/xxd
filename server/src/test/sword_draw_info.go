package main

import (
	"fmt"
	"time"

	"client_test"
	"game_server/api/protocol/sword_soul_api"
)

func main() {
	client := client_test.NewClient("127.0.0.1:8080")
	client.Player_FromPlatformLogin(1, "python3.1", "施语娜", 2, "", 0, false)

	//注册回调函数
	client.OutSwordSoul_Info = func(out *sword_soul_api.Info_Out) {
		fmt.Println("\n==============")
		fmt.Printf("sword_soul_api.InfoOut:\n%#v\n", out)
		fmt.Println("==============\n")
	}

	client.SwordSoul_Info()

	b := time.Tick(5 * time.Second)
	<-b
}
