package main

import (
	"fmt"
	"time"

	"client_test"
	"game_server/api/protocol/role_api"
)

func main() {
	client := client_test.NewClient("127.0.0.1:8080")
	client.Player_FromPlatformLogin("fighta", "fighta", 1, "", 0, false)

	//注册回调函数
	client.OutRole_GetFightNum = func(out *role_api.GetFightNum_Out) {
		fmt.Println("==============")
		fmt.Printf("role_api.GetFightNum_Out:\n%#v\n", out)
		fmt.Println("==============")
	}

	client.Role_GetFightNum(role_api.FIGHTNUM_TYPE_ALL)

	b := time.Tick(5 * time.Second)
	<-b
}
