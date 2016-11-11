package push_notify_dat

import (
	"core/i18l"
	
)

const( 
	UNSCHEDULED_NOTIFY = -1
	PAUSED_NOTIFY = -2
)
const( 
	AFTERNOONPHYSICAL = 1
	NIGHTPHYSICAL = 2
	SEASHOP = 3
	MAXPHYSICAL = 4
	PRIVATEMESSAGE = 5
	ARENAATTACK = 6
)


var DefaultNotificationID = []int32{1,2,}





//自动生产请勿修改
func Load(){
	mapPushNotify = make(map[int32]*PushNotify)
	mapPushNotify[4] = &PushNotify {
		Content: i18l.T.Tran("我们已经为大侠回复满体力了！赶快来继续探险吧！"),
		Name: i18l.T.Tran("体力回复满通知"),
	}

	mapPushNotify[5] = &PushNotify {
		Content: i18l.T.Tran("您有一位好友在游戏里给你发送了消息哦！快去看看他想悄悄告诉你什么！"),
		Name: i18l.T.Tran("私聊消息通知"),
	}

	mapPushNotify[6] = &PushNotify {
		Content: i18l.T.Tran("您在比武场遭到了挑战！快去看看是谁那么胆大包天！"),
		Name: i18l.T.Tran("比武场被攻击通知"),
	}

}

