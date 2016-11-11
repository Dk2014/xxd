package main

import (
	"fmt"
	"time"

	"client_test"
	"game_server/api/protocol/daily_sign_in_api"
)

func main() {
	client := client_test.NewClient("127.0.0.1:8080")
	client.Player_FromPlatformLogin(1, "sign_info13", "sign_info13", 2, "", 0, false)

	//注册回调函数
	client.OutDailySignIn_Info = func(out *daily_sign_in_api.Info_Out) {
		fmt.Println("\n==============")
		fmt.Printf("daily_sign_in_api.InfoOut:\n%#v\n", out)
		fmt.Println("==============\n")
	}

	client.OutDailySignIn_Sign = func(out *daily_sign_in_api.Sign_Out) {
		fmt.Println("\n==============")
		fmt.Printf("daily_sign_in_api.SignOut:\n%#v\n", out)
		fmt.Println("==============\n")
	}

	client.DailySignIn_Info()
	client.DailySignIn_Sign()
	client.DailySignIn_Info()

	b := time.Tick(5 * time.Second)
	<-b
}
