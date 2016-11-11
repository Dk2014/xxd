package main

import (
	"fmt"
	"time"

	"client_test"
	"game_server/api/protocol/mission_api"
)

func main() {
	client := client_test.NewClient("127.0.0.1:8080")
	client.Player_FromPlatformLogin("hard93", "秦剧乘", 1, "", 0, false)

	//注册回调函数

	client.OutMission_GetHardLevel = func(out *mission_api.GetHardLevel_Out) {
		fmt.Println("\n==============")
		fmt.Printf("mission_api.GetHardLevel:\n%#v\n", out)
		fmt.Println("==============\n")
	}

	//获取用户信息，期待里面会有新的字段
	client.Mission_GetHardLevel()

	b := time.Tick(5 * time.Second)
	<-b
}
