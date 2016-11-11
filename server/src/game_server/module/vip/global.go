package vip

import "game_server/global"

func getVIPTotal() int64 {
	return int64(global.GetPlayerVIPTotal())
}

func getLevelVipTotal() []int32 {
	return global.GetPlayerVIP()
}
