package module

import (
	"core/time"
	"game_server/dat/player_dat"
	"game_server/mdb"
)

func preOperate(player_acitivity *mdb.PlayerActivity) (isRefresh bool) {
	//每日清零时间点
	if !time.IsInPointHour(player_dat.RESET_PLAYER_ACTIVITY_IN_HOUR, player_acitivity.LastUpdate) {
		player_acitivity.Activity = 0
		isRefresh = true
	}
	return
}

func GetPlayerActivity(db *mdb.Database) (activity int32, isRefresh bool) {
	player_acitivity := db.Lookup.PlayerActivity(db.PlayerId())
	isRefresh = preOperate(player_acitivity)
	activity = player_acitivity.Activity
	if isRefresh {
		player_acitivity.LastUpdate = time.GetNowTime()
		db.Update.PlayerActivity(player_acitivity)
	}
	return
}

func AddPlayerActivity(db *mdb.Database, increVal int32) (activity int32, isRefresh bool) { //activity为添加之后得新的活跃度
	player_acitivity := db.Lookup.PlayerActivity(db.PlayerId())
	isRefresh = preOperate(player_acitivity)
	player_acitivity.Activity += increVal
	if player_acitivity.Activity > player_dat.MAX_PLAYER_ACTIVITY { //到达活跃度上限
		player_acitivity.Activity = player_dat.MAX_PLAYER_ACTIVITY
	}
	player_acitivity.LastUpdate = time.GetNowTime()
	db.Update.PlayerActivity(player_acitivity)
	activity = player_acitivity.Activity
	return
}
