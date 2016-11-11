package main

import (
	"client_test"
	"fmt"
	"game_server/api/protocol/battle_api"
	"game_server/api/protocol/player_api"
)

func main() {
	// 建立三个客户端
	c1 := client_test.NewClient("127.0.0.1:8080")
	c2 := client_test.NewClient("127.0.0.1:8080")
	c3 := client_test.NewClient("127.0.0.1:8080")

	// 两个客户端登陆
	c2.Player_FromPlatformLogin("l21", "喻旗瑞", 1, "", 0, false)
	c3.Player_FromPlatformLogin("l99", "任笑若", 2, "", 0, false)

	// 设置c1登陆成功后创建多人关卡房间，并让c2，c3加入房间，最后c1开启战斗
	c1.OutPlayer_FromPlatformLogin = func(out *player_api.FromPlatformLogin_Out) {
		c1.MultiLevel_CreateRoom(1)

		c2.MultiLevel_EnterRoom(out.PlayerId)
		c3.MultiLevel_EnterRoom(out.PlayerId)

		c1.Battle_StartBattle(battle_api.BATTLE_TYPE_MULTILEVEL, 1)
	}

	c1.OutBattle_StartBattle = func(out *battle_api.StartBattle_Out) {
		fmt.Println("OutBattle_StartBattle")
		fmt.Println(out)
	}

	c1.Player_FromPlatformLogin("l20", "史善宛", 2, "", 0, false)

	exitChan := make(chan bool)
	<-exitChan
}
