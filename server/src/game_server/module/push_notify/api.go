package push_notify

import (
	"core/net"
	"game_server/api/protocol/push_notify_api"
	"game_server/dat/push_notify_dat"
	"game_server/module"
	"game_server/module/rpc"
)

func init() {
	push_notify_api.SetInHandler(PushNotifyAPI{})
}

type PushNotifyAPI struct{}

func (api PushNotifyAPI) PushInfo(session *net.Session, in *push_notify_api.PushInfo_In) {
	state := module.State(session)
	out := &push_notify_api.PushInfo_Out{}
	options := state.Database.Lookup.PlayerPushNotifySwitch(state.PlayerId)
	bitMask := uint64(options.Options)

	for _, id := range push_notify_dat.DefaultNotificationID {
		if (bitMask & (1 << uint64(id))) != 0 {
			out.EffectNotification = append(out.EffectNotification, push_notify_api.PushInfo_Out_EffectNotification{
				NotificationId: id,
			})
		}
	}

	session.Send(out)
}

func (api PushNotifyAPI) PushNotificationSwitch(session *net.Session, in *push_notify_api.PushNotificationSwitch_In) {
	state := module.State(session)
	options := state.Database.Lookup.PlayerPushNotifySwitch(state.PlayerId)
	bitMask := uint64(options.Options)

	//TODO 检查客户端传过来的 ID 是否合法

	//打开
	if in.TurnOn && (bitMask&(1<<uint64(in.NotificationId))) == 0 {
		bitMask += (1 << uint64(in.NotificationId))
		options.Options = int64(bitMask)
		state.Database.Update.PlayerPushNotifySwitch(options)
		rpc.RemoteUpdatePlayerPushNotificationOptions(state.PlayerId, options.Options)
		return
	}

	//关闭
	if !in.TurnOn && (bitMask&(1<<uint64(in.NotificationId))) != 0 {
		bitMask -= (1 << uint64(in.NotificationId))
		options.Options = int64(bitMask)
		state.Database.Update.PlayerPushNotifySwitch(options)
		rpc.RemoteUpdatePlayerPushNotificationOptions(state.PlayerId, options.Options)
		return
	}
}
