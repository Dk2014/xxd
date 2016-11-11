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
	client_test.FromPlatformLogin(client, 1, "auto010", "auto010", 1, false)

	client.OutMission_AutoFightLevel = func(out *mission_api.AutoFightLevel_Out) {
		fmt.Println("\n==============")
		fmt.Printf("mission.AutoFightLevel:\n")
		bytes, _ := json.MarshalIndent(out, "", "\t")
		fmt.Printf(string(bytes))
		fmt.Println("\n==============\n")
	}

	client.OutPlayer_FromPlatformLogin = func(out *player_api.FromPlatformLogin_Out) {
		client.Debug_OpenFunc(1500)
		client.Debug_AddItem(250, 100)
		//关卡类型 普通的关卡0 难度关卡 8
		//关卡ID
		//扫荡次数
		client.Mission_AutoFightLevel(0, 19, 2)
	}

	b := time.Tick(5 * time.Second)
	<-b
}