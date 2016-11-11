package main

import (
	"encoding/json"
	"fmt"
	"time"

	"client_test"
	"game_server/api/protocol/mission_api"
	"game_server/api/protocol/player_api"
)

func main() {
	client := client_test.NewClient("127.0.0.1:8080")
	client_test.FromPlatformLogin(client, 1, "auto001", "auto001", 2, false)

	client.OutMission_AutoFightLevel = func(out *mission_api.AutoFightLevel_Out) {
		fmt.Println("\n==============")
		fmt.Printf("mission.AutoFightLevel:\n")
		bytes, _ := json.MarshalIndent(out, "", "\t")
		fmt.Printf(string(bytes))
		fmt.Println("==============\n")
	}

	client.OutPlayer_FromPlatformLogin = func(out *player_api.FromPlatformLogin_Out) {
		//关卡类型 普通的关卡0 难度关卡 8
		//关卡ID
		//扫荡次数
		client.Mission_AutoFightLevel(8, 62, 1)
	}

	b := time.Tick(3 * time.Second)
	<-b
}
