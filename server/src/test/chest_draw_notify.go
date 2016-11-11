package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"client_test"
	"game_server/api/protocol/notify_api"
	"game_server/api/protocol/player_api"
)

func main() {
	client := client_test.NewClient("127.0.0.1:8080")
	name := fmt.Sprint("chest_draw_notify_%d", rand.Int())
	client_test.FromPlatformLogin(client, 1, name, name, 1, false)

	client.OutPlayer_FromPlatformLogin = func(out *player_api.FromPlatformLogin_Out) {
		client.Draw_DrawChest(1)
		client.Draw_DrawChest(4)
	}

	client.OutNotify_SendChestDrawReady = func(out *notify_api.SendChestDrawReady_Out) {
		fmt.Println("\n==============")
		fmt.Printf("notify.SendChestDrawReady:\n")
		bytes, _ := json.MarshalIndent(out, "", "\t")
		fmt.Printf(string(bytes))
		fmt.Println("==============\n")
	}

	b := time.Tick(30 * time.Second)
	<-b
}
