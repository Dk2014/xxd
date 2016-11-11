package main

import (
	"fmt"
	"time"

	"client_test"
	"game_server/api/protocol/sword_soul_api"
)

func main() {
	client := client_test.NewClient("127.0.0.1:8080")
	client.Player_FromPlatformLogin("sword001", "施奉勤", 1, "", 0, false)

	//注册回调函数
	client.OutSwordSoul_Info = func(out *sword_soul_api.Info_Out) {
		fmt.Println("\n==============")
		fmt.Printf("sward_api.Info_Out:\n%#v\n", out)
		fmt.Printf("now:\n%d\n", time.Now().Unix())
		fmt.Println("==============\n")
	}

	client.SwordSoul_Info()
	client.Debug_ResetSwordSoulDrawCd()
	b := time.Tick(2 * time.Second)
	<-b
	client.SwordSoul_Info()

	c := time.Tick(5 * time.Second)
	<-c
}
