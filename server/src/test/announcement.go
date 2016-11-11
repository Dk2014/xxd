package main

import (
	"fmt"
	"time"

	"client_test"
	"game_server/api/protocol/notify_api"
	"game_server/api/protocol/announcement_api"
)

func main() {
	client := client_test.NewClient("127.0.0.1:8080")
	client.Player_FromPlatformLogin("annc01", "annc01", 1, "", 0, false)
	client_rpc := client_test.NewClient("127.0.0.1:9081")
	//client_rpc.Player_FromPlatformLogin("annc01", "annc01", 1, "", 0, false)

	//注册回调函数
	client.OutNotify_SendAnnouncement = func(out *notify_api.SendAnnouncement_Out) {
		fmt.Println("==============")
		fmt.Printf("notify_api.SendAnnouncement:\n%#v\n", out)
		fmt.Println("==============")
	}
	
	client.OutAnnouncement_GetList = func(out *announcement_api.GetList_Out) {
		fmt.Println("==============")
		fmt.Printf("announcement_api.GetList\n%#v\n", out)
		fmt.Println("==============")
	}

	client_rpc.OutNotify_SendAnnouncement = func(out *notify_api.SendAnnouncement_Out) {
		fmt.Println("==============")
		fmt.Printf("rpc:notify_api.SendAnnouncement:\n%#v\n", out)
		fmt.Println("==============")
	}
	

	client_rpc.OutAnnouncement_GetList = func(out *announcement_api.GetList_Out) {
		fmt.Println("==============")
		fmt.Printf("rpc:announcement_api.GetList\n%#v\n", out)
		fmt.Println("==============")
	}


	//调用api.CreateAnnouncement发送创建一条公告
	client.Debug_CreateAnnouncement()

	//调用announcement.GetList
	//client_rpc.Announcement_GetList()
	
	b := time.Tick(5 * time.Second)
	<-b

	//client.Announcement_GetList()
	c := time.Tick(3 * time.Second)
	<-c



}
