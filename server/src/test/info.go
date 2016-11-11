package main

import (
	"fmt"
	"time"

	"client_test"
	"game_server/api/protocol/player_api"
)

func main() {
	client := client_test.NewClient("127.0.0.1:8080")
	client.Player_FromPlatformLogin("sworda07", "王羽妖", 2, "", 0, false)

	//注册回调函数
	client.OutPlayer_Info = func(out *player_api.Info_Out) {
		fmt.Println("\n==============")
		fmt.Printf("player_api.InfoOut:\n%#v\n", out)
		fmt.Println("==============\n")
	}

	//调用调试接口获取设置一定数量的元宝
	//client.Debug_SetIngot(1000)

	//获取用户信息，期待里面会有新的字段
	client.Player_Info()

	b := time.Tick(5 * time.Second)
	<-b
}
