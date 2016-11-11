package main

import (
	"fmt"
	"time"

	"client_test"
	"game_server/api/protocol/player_api"
	"game_server/api/protocol/debug_api"
	"game_server/api/protocol/mission_api"
	"game_server/api/protocol/battle_api"
)


func main() {
	client := client_test.NewClient("127.0.0.1:8080")
	client.Player_FromPlatformLogin("hard24", "hard24", 1, "", 0, false)

	//注册回调函数
	client.OutPlayer_Info = func(out *player_api.Info_Out) {
		fmt.Println("\n==============")
		fmt.Printf("player_api.InfoOut:\n%#v\n", out)
		fmt.Println("==============\n")
	}

	client.OutMission_GetHardLevel = func(out *mission_api.GetHardLevel_Out) {
		fmt.Println("\n==============")
		fmt.Printf("mission_api.GetHardLevel:\n%#v\n", out)
		fmt.Println("==============\n")
	}



	//获取用户信息，期待里面会有新的字段
	client.Player_Info()
	//进入难度关卡 select * from mission_level where parent_type=8;
	client.Mission_EnterHardLevel(67)
	//开始战斗 select * from mission_enemy where mission_level_id=64;
	client.Battle_StartBattle(battle_api.BATTLE_TYPE_HARD, 204)

	b := time.Tick(5 * time.Second)
	<-b
}
