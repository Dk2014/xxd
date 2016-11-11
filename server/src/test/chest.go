package main

import (
	"client_test"
	"fmt"
	"game_server/api/protocol/debug_api"
	"game_server/api/protocol/draw_api"
	"time"
)

func main() {
	client := client_test.NewClient("127.0.0.1:8080")
	client.Player_FromPlatformLogin("test1", "test1", 1, "", 0, false)

	client.OutDebug_OpenFunc = func(out *debug_api.OpenFunc_Out) {
		fmt.Println("==============")
		fmt.Printf("debug_api.OpenFunc_Out:\n%#v\n", out)
		fmt.Println("==============")
	}

	client.Debug_OpenFunc(1600)

	client.OutDraw_GetChestInfo = func(out *draw_api.GetChestInfo_Out) {
		fmt.Println("==============")
		fmt.Printf("draw_api.GetChestInfo_Out:\n%#v\n", out)
		fmt.Println("==============")
	}

	client.Draw_GetChestInfo()

	client.Draw_DrawChest(draw_api.CHEST_TYPE_COIN_FREE)

	client.Draw_GetChestInfo()

	c := time.Tick(3 * time.Second)
	<-c
}
