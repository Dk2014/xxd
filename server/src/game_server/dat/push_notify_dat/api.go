package push_notify_dat

import (
	"core/fail"
)

var (
	mapPushNotify map[int32]*PushNotify
)

type PushNotify struct {
	Content string // 内容
	Name    string // 推送通知名称
}

// ############### 对外接口实现 coding here ###############

func GetPushNotificationById(id int32) *PushNotify {
	data, exist := mapPushNotify[id]
	fail.When(!exist, "找不到推送通知数据")
	return data
}
