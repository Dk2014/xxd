package push_notify

import (
	"game_server/config"
	"game_server/dat/push_notify_dat"
	"game_server/global"
	"game_server/mdb"
	"game_server/module"
	"game_server/tencent"
)

type PushNotifyMod struct{}

func init() {
	module.PushNotify = PushNotifyMod{}
}

func (mod PushNotifyMod) EnabledPushNotify(db *mdb.Database, notificationId int32) (enabled bool) {
	var bitMask uint64
	if config.ServerCfg.EnableGlobalServer {
		playerInfo := global.GetPlayerInfo(db.PlayerId())
		bitMask = uint64(playerInfo.PushNotificationOptions)
	} else {
		options := db.Lookup.PlayerPushNotifySwitch(db.PlayerId())
		bitMask = uint64(options.Options)
	}

	return (bitMask & (1 << uint64(notificationId))) != 0
}

func (mod PushNotifyMod) SendNotification(pid int64, notificationId int32) {
	notification := push_notify_dat.GetPushNotificationById(notificationId)
	tencent.SendNotificationToSingleAccount(pid, notification.Name, notification.Content)
}
