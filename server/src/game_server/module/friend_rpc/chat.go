package friend_rpc

import (
	"core/fail"
	"core/time"
	"game_server/api/protocol/friend_api"
	"game_server/api/protocol/notify_api"
	"game_server/dat/friend_dat"
	"game_server/dat/push_notify_dat"
	"game_server/global"
	"game_server/mdb"
	"game_server/module"
	"game_server/tlog"
)

func Chat(state *module.SessionState, in *friend_api.Chat_In) (banned bool) {
	if !module.ChatRPC.CanChat(state.Database) {
		return true
	}
	//陌生人对话要先添加好友关系
	var playerFriend, senderPlayerFriend *mdb.PlayerGlobalFriend
	have := false
	state.Database.Select.PlayerGlobalFriend(func(row *mdb.PlayerGlobalFriendRow) {
		if row.FriendPid() == in.Pid {
			have = true
			senderPlayerFriend = row.GoObject()
			row.Break()
		}
	})

	// 不存在好友关系，首次聊天
	if !have {
		senderPlayerFriend, playerFriend = insertFriend(state, in.Pid, 2)
	} else {
		// 检查对方是否屏蔽发送者的消息
		state.Database.AgentExecute(in.Pid, func(db *mdb.Database) {
			db.Select.PlayerGlobalFriend(func(row *mdb.PlayerGlobalFriendRow) {
				if row.FriendPid() == state.PlayerId {
					playerFriend = row.GoObject()
					row.Break()
				}
			})
		})
	}
	tlog.PlayerSystemModuleFlowLog(state.Database, tlog.SMT_FRIEND_CHAT)
	now := time.GetNowTime()
	state.Database.Insert.PlayerGlobalFriendChat(&mdb.PlayerGlobalFriendChat{
		Pid:       state.PlayerId,
		FriendPid: in.Pid,
		Mode:      friend_dat.CHAT_MESSAGE_MODE_SEND,
		SendTime:  now,
		Message:   string(in.Message),
	})

	senderPlayerFriend.LastChatTime = now
	state.Database.Update.PlayerGlobalFriend(senderPlayerFriend)

	// 对方屏蔽了消息
	if playerFriend.BlockMode == 1 {
		return
	}

	state.Database.AgentExecute(in.Pid, func(db *mdb.Database) {
		db.Insert.PlayerGlobalFriendChat(&mdb.PlayerGlobalFriendChat{
			Pid:       in.Pid,
			FriendPid: state.PlayerId,
			Mode:      friend_dat.CHAT_MESSAGE_MODE_RECEIVE,
			SendTime:  now,
			Message:   string(in.Message),
		})

		playerFriend.LastChatTime = now
		db.Update.PlayerGlobalFriend(playerFriend)
	})

	if session, ok := module.Player.GetPlayerOnline(in.Pid); ok {
		playerInfo := global.GetPlayerInfo(state.PlayerId)
		session.Send(&notify_api.Chat_Out{
			Pid:      playerInfo.PlayerId,
			RoleId:   playerInfo.RoleId,
			Nickname: playerInfo.PlayerNick,
			Level:    playerInfo.RoleLevel,
			FightNum: 0,
			Message:  in.Message,
		})
	} else {
		state.Database.AgentExecute(in.Pid, func(db *mdb.Database) {
			playerFriendState := db.Lookup.PlayerGlobalFriendState(state.PlayerId)
			fail.When(playerFriendState == nil, "Chat playerFriendState nil")
			playerFriendState.ExistOfflineChat = 1
			db.Update.PlayerGlobalFriendState(playerFriendState)
			if module.PushNotify.EnabledPushNotify(db, push_notify_dat.PRIVATEMESSAGE) {
				module.PushNotify.SendNotification(in.Pid, push_notify_dat.PRIVATEMESSAGE)
			}
		})
	}
	return
}

func GetChatHistory(state *module.SessionState, in *friend_api.GetChatHistory_In) (out *friend_api.GetChatHistory_Out) {
	out = &friend_api.GetChatHistory_Out{}
	out.Messages = []friend_api.GetChatHistory_Out_Messages{}

	now := time.GetNowTime()
	state.Database.Select.PlayerGlobalFriendChat(func(row *mdb.PlayerGlobalFriendChatRow) {
		if now-row.SendTime() >= friend_dat.CHAT_MESSAGE_STORE_TIME {
			state.Database.Delete.PlayerGlobalFriendChat(row.GoObject())
			return
		}

		if row.FriendPid() == in.Pid {
			out.Messages = append(out.Messages, friend_api.GetChatHistory_Out_Messages{
				Mode:     friend_api.MsgMode(row.Mode()),
				Id:       row.Id(),
				SendTime: row.SendTime(),
				Message:  []byte(row.Message()),
			})
		}
	})

	return
}

func GetOfflineMessages(state *module.SessionState) (out *friend_api.GetOfflineMessages_Out, ret bool) {
	out = &friend_api.GetOfflineMessages_Out{
		Chats:    []friend_api.GetOfflineMessages_Out_Chats{},
		Listener: []friend_api.GetOfflineMessages_Out_Listener{},
	}

	playerInfo := global.GetPlayerInfo(state.PlayerId)
	if playerInfo.LastOfflineTime == 0 {
		return
	}

	// 玩家离线被关注记录
	state.Database.Select.PlayerGlobalFriend(func(row *mdb.PlayerGlobalFriendRow) {
		if row.BlockMode() == 0 && row.FriendMode() > friend_dat.FRIEND_MODE_LISTEN_ONLY && playerInfo.LastOfflineTime < row.FriendTime() {
			out.Listener = append(out.Listener, friend_api.GetOfflineMessages_Out_Listener{
				Pid:  row.FriendPid(),
				Nick: []byte(row.FriendNick()),
			})
		}
	})

	//检查玩家是否有离线消息状态
	playerFriendState := state.Database.Lookup.PlayerGlobalFriendState(state.PlayerId)
	fail.When(playerFriendState == nil, "GetOfflineMessages playerFriendState nil")
	if playerFriendState.ExistOfflineChat == 0 {
		return
	}

	records := make(map[int64]bool)
	state.Database.Select.PlayerGlobalFriendChat(func(row *mdb.PlayerGlobalFriendChatRow) {
		if row.Mode() == friend_dat.CHAT_MESSAGE_MODE_RECEIVE && row.SendTime() >= playerInfo.LastOfflineTime && row.SendTime() < playerInfo.LastLoginTime {
			if _, ok := records[row.FriendPid()]; ok {
				return
			}

			records[row.FriendPid()] = true
			friendInfo := global.GetPlayerInfo(row.FriendPid())

			out.Chats = append(out.Chats, friend_api.GetOfflineMessages_Out_Chats{
				Pid:      friendInfo.PlayerId,
				Nickname: friendInfo.PlayerNick,
				RoleId:   friendInfo.RoleId,
				Level:    friendInfo.RoleLevel,
			})
		}
	})

	ret = true
	playerFriendState.ExistOfflineChat = 0
	state.Database.Update.PlayerGlobalFriendState(playerFriendState)
	return
}

func CleanChatHistory(state *module.SessionState, pid int64) {
	state.Database.Select.PlayerGlobalFriendChat(func(row *mdb.PlayerGlobalFriendChatRow) {
		if row.FriendPid() == pid {
			state.Database.Delete.PlayerGlobalFriendChat(row.GoObject())
		}
	})

	state.Database.Select.PlayerGlobalFriend(func(row *mdb.PlayerGlobalFriendRow) {
		if row.FriendPid() == pid {
			playerFriend := row.GoObject()
			playerFriend.LastChatTime = 0
			state.Database.Update.PlayerGlobalFriend(playerFriend)
			row.Break()
		}
	})
}
