package channel_rpc

import (
	"core/fail"
	"core/net"
	"core/time"
	"game_server/api/protocol/channel_api"
	"game_server/dat/channel_dat"
	"game_server/dat/player_dat"
	"game_server/global"
	"game_server/mdb"
	"game_server/module"
	"game_server/module/rpc"
	"game_server/tlog"
	"game_server/xdlog"
)

type ChannelAPI struct{}

func init() {
	channel_api.SetInHandler(ChannelAPI{})
}

func (this ChannelAPI) GetLatestWorldChannelMessage(session *net.Session, in *channel_api.GetLatestWorldChannelMessage_In) {
	out := &channel_api.GetLatestWorldChannelMessage_Out{}
	msgs := module.GetLatestWorldChannelMessage()
	for _, msg := range msgs {
		out.Messages = append(out.Messages, channel_api.GetLatestWorldChannelMessage_Out_Messages{
			Pid:        msg.Pid,
			Nickname:   []byte(msg.Nickname),
			MsgType:    channel_api.MessageType(msg.MsgType),
			Timestamp:  msg.Timestamp,
			Parameters: []byte(msg.Parameters),
			TplId:      msg.TplId,
		})
	}
	session.Send(out)

}

func (this ChannelAPI) AddWorldChat(session *net.Session, in *channel_api.AddWorldChat_In) {
	state := module.State(session)
	if !module.ChatRPC.CanChat(state.Database) {
		session.Send(&channel_api.AddWorldChat_Out{
			Banned: true,
		})
		return
	}

	size := 0
	for _ = range string(in.Content) {
		size++
		fail.When(size > channel_dat.WORLD_CHAT_MAX_CONTENT_LEN, "超过聊天信息超过最大长度")
	}
	fail.When(size == 0, "空消息")

	worldChatState := getWorldChatState(state.Database)
	now := time.GetNowTime()
	//检查CD时间
	fail.When(worldChatState.Timestamp+channel_dat.WORLD_CHAT_CD_TIME > now, "CD冷却中")

	worldChatState.DailyNum += 1
	addMsg := func() {
		msgTpl := channel_dat.MessageCommonChat{
			Content: channel_dat.ParamString{string(in.Content)},
		}
		module.AddWorldChannelMessage(&module.Message{
			Pid:       state.PlayerId,
			MsgType:   module.CHANNEL_CHAT,
			Nickname:  state.PlayerNick,
			Timestamp: now,
			//Content:   in.Content,
			Parameters: msgTpl.GetParameters(),
			TplId:      msgTpl.GetTplId(),
		})
		worldChatState.Timestamp = now
		state.Database.Update.PlayerGlobalWorldChatState(worldChatState)
		session.Send(&channel_api.AddWorldChat_Out{})
	}
	if worldChatState.DailyNum > channel_dat.WORLD_CHAT_DAILY_FREE_NUM {
		rpc.RemoteDecMoney(state.PlayerId, channel_dat.WORLD_CHAT_COST, player_dat.INGOT, tlog.MFR_WORLD_CHAT_COST, xdlog.ET_WORLD_CHAT, addMsg)
	} else {
		addMsg()
	}
}

func (this ChannelAPI) WorldChannelInfo(session *net.Session, in *channel_api.WorldChannelInfo_In) {
	state := module.State(session)
	worldChatState := getWorldChatState(state.Database)
	out := &channel_api.WorldChannelInfo_Out{
		Timestamp: worldChatState.Timestamp,
		DailyNum:  worldChatState.DailyNum,
	}
	session.Send(out)
}

func (this ChannelAPI) AddCliqueChat(session *net.Session, in *channel_api.AddCliqueChat_In) {
	state := module.State(session)
	if !module.ChatRPC.CanChat(state.Database) {
		session.Send(&channel_api.AddCliqueChat_Out{
			Banned: true,
		})
		return
	}
	size := 0
	for _ = range string(in.Content) {
		size++
		fail.When(size > channel_dat.WORLD_CHAT_MAX_CONTENT_LEN, "超过聊天信息超过最大长度")
	}
	fail.When(size == 0, "空消息")

	playerCliqueInfo := state.Database.Lookup.PlayerGlobalCliqueInfo(state.PlayerId)
	fail.When(playerCliqueInfo == nil || playerCliqueInfo.CliqueId <= 0, "帮派信息不存在或帮派id不存在")
	if playerCliqueInfo != nil && playerCliqueInfo.CliqueId > 0 {
		msgTpl := channel_dat.MessageCommonChat{
			Content: channel_dat.ParamString{string(in.Content)},
		}
		module.CliqueRPC.AddCliqueChat(playerCliqueInfo.CliqueId, state.PlayerId, state.PlayerNick, msgTpl)
		session.Send(&channel_api.AddCliqueChat_Out{})
	}

}

func (this ChannelAPI) GetLatestCliqueMessages(session *net.Session, in *channel_api.GetLatestCliqueMessages_In) {
	state := module.State(session)
	playerCliqueInfo := state.Database.Lookup.PlayerGlobalCliqueInfo(state.PlayerId)
	if playerCliqueInfo != nil && playerCliqueInfo.CliqueId > 0 {
		out := &channel_api.GetLatestCliqueMessages_Out{}
		messages := module.CliqueRPC.GetLatestMessages(playerCliqueInfo.CliqueId)
		for i, _ := range messages {
			out.Messages = append(out.Messages, channel_api.GetLatestCliqueMessages_Out_Messages{
				Message: messages[i],
			})
		}
		session.Send(out)
	}
}

func getWorldChatState(db *mdb.Database) (worldChatState *mdb.PlayerGlobalWorldChatState) {
	worldChatState = db.Lookup.PlayerGlobalWorldChatState(db.PlayerId())
	if worldChatState == nil {
		playerInfo := global.GetPlayerInfo(db.PlayerId())
		fail.When(playerInfo.RoleLevel < channel_dat.WORLD_CHAT_SERVER_OPEN_LEVEL, "等级不足")
		worldChatState = &mdb.PlayerGlobalWorldChatState{
			Pid: db.PlayerId(),
		}
		db.Insert.PlayerGlobalWorldChatState(worldChatState)
	}
	if !time.IsToday(worldChatState.Timestamp) {
		worldChatState.DailyNum = 0
	}
	return worldChatState
}
