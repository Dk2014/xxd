package main

import (
	"fmt"
	"time"

	"client_test"
	coreTime "core/time"
	"game_server/api/protocol/trader_api"
)

func main() {
	client := client_test.NewClient("127.0.0.1:8080")
	client.Player_FromPlatformLogin(1, "trader04", "trader04", 2, "", 0, false)

	//注册回调函数
	var traderIds []int16
	client.OutTrader_Info = func(out *trader_api.Info_Out) {
		fmt.Println("\n==============")
		fmt.Printf("trader_api.InfoOut:\n%#v\n", out)
		now := time.Now().Unix()
		todayZero := coreTime.GetTodayZero()
		for _, schedules := range out.Schedule {
			for _, schedule := range schedules.During {
				if schedule.Expire < 0 || schedule.Expire > now {
					if schedule.Showup+todayZero < now && now < schedule.Disappear+todayZero {
						traderIds = append(traderIds, schedules.TraderId)
						break
					}
				}
			}
			fmt.Printf("\n\n")
		}
		fmt.Printf("valid traders :%v\n", traderIds)
		for _, traderId := range traderIds {
			client.Trader_StoreState(traderId)
			client.Trader_Refresh(traderId)
			client.Trader_StoreState(traderId)
		}
		fmt.Println("==============\n")
	}

	client.OutTrader_StoreState = func(out *trader_api.StoreState_Out) {
		fmt.Println("\n==============")
		fmt.Printf("trader_api.Store_State:\n%#v\n", out)
		fmt.Println("==============\n")
	}

	client.Trader_Info()

	b := time.Tick(5 * time.Second)
	<-b
}
