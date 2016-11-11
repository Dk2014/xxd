package main

import (
	"fmt"
	"time"

	"client_test"
	"game_server/api/protocol/notify_api"
	"game_server/api/protocol/sword_soul_api"
)

func main() {
	client := client_test.NewClient("127.0.0.1:8080")
	client.Player_FromPlatformLogin(1, "swordbug002", "庄狱让", 1, "", 0, false)

	//注册回调函数
	client.OutSwordSoul_Info = func(out *sword_soul_api.Info_Out) {
		fmt.Println("\n==============")
		fmt.Printf("sward_api.Info_Out:\n%#v\n", out)
		fmt.Println("==============\n")
	}

	client.OutNotify_SendSwordSoulDrawNumChange = func(out *notify_api.SendSwordSoulDrawNumChange_Out) {
		fmt.Println("\n==============")
		fmt.Printf("notify_api.SendSwordSoulDrawNumChange_Out:\n%#v\n", out)
		fmt.Println("==============\n")
	}

	client.OutSwordSoul_Draw = func(out *sword_soul_api.Draw_Out) {
		fmt.Println("\n==============")
		fmt.Printf("sward_api.Draw_Out:\n%#v\n", out)
		fmt.Println("==============\n")
	}

	//开启剑心功能权值
	//client.Debug_OpenFunc(1500)

	//查看剑心信息
	client.SwordSoul_Info()

	//拔剑一次
	//client.SwordSoul_Draw(

	//查看拔剑次数

	//把拔剑次数恢复时间设置得短

	c := time.Tick(15 * time.Second)
	<-c
}
