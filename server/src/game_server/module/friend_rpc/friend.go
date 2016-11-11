package friend_rpc

import (
	"core/fail"
	"core/time"
	"fmt"
	"game_server/api/protocol/friend_api"
	"game_server/dat/friend_dat"
	"game_server/dat/mail_dat"
	"game_server/dat/quest_dat"
	"game_server/global"
	"game_server/mdb"
	"game_server/module"
	"game_server/module/rpc"
	"game_server/tlog"
)

func GetFriendList(state *module.SessionState) (out *friend_api.GetFriendList_Out) {
	out = new(friend_api.GetFriendList_Out)
	playerFriendState := state.Database.Lookup.PlayerGlobalFriendState(state.PlayerId)
	fail.When(playerFriendState == nil, "GetFriendList playerFriendState nil")
	out.CancelListenCount = playerFriendState.DeleteDayCount
	out.PlatformFriendNum = playerFriendState.PlatformFriendNum
	out.Friends = []friend_api.GetFriendList_Out_Friends{}

	state.Database.Select.PlayerGlobalFriend(func(row *mdb.PlayerGlobalFriendRow) {
		/*
			TODO 需要review一下：
			约束条件：下面添加好友的逻辑要求 如果  A 和 B 之间的好友关系要没没有记录、要么必定有两条记录。
			这里可能导致 (A,B) 这条记录删除了，但是 (B,A) 记录还在，
			那么 B 添加 A 为好友是不满足约束条件，导致添加好友失败
			可以定期通过数据升级脚本的形式来删除冗余好友关系
				//删除没消息的陌生人
				if row.FriendMode() == int8(friend_api.FRIEND_MODE_STRANGE) && row.BlockMode() == 0 {
					noChat := true
					state.Database.Select.PlayerGlobalFriendChat(func(row2 *mdb.PlayerGlobalFriendChatRow) {
						if row.FriendPid() == row2.FriendPid() {
							noChat = false
							row2.Break()
						}
					})
					if noChat {
						state.Database.Delete.PlayerGlobalFriend(row.GoObject())
					}
				}
		*/

		friendInfo := global.GetPlayerInfo(row.FriendPid())

		out.Friends = append(out.Friends,
			friend_api.GetFriendList_Out_Friends{
				Pid:          row.FriendPid(),
				RoleId:       row.FriendRoleId(),
				Nickname:     []byte(row.FriendNick()),
				Level:        friendInfo.RoleLevel,
				FightNum:     friendInfo.FightNum,
				FriendMode:   friend_api.FriendMode(row.FriendMode()),
				BlockMode:    row.BlockMode(),
				LastChatTime: row.LastChatTime(),
				FriendTime:   row.FriendTime(),
			})
	})

	return
}

func ListenByPid(state *module.SessionState, pid int64) (out *friend_api.ListenByNick_Out) {
	out = new(friend_api.ListenByNick_Out)

	//添加自己,返回失败
	if state.PlayerId == pid {
		out.Result = friend_api.ADD_RESULT_FAILED_ADD_SELF
		return
	}

	count := 0
	var playerFriend *mdb.PlayerGlobalFriend
	state.Database.Select.PlayerGlobalFriend(func(row *mdb.PlayerGlobalFriendRow) {
		if playerFriend == nil && row.FriendPid() == pid {
			playerFriend = row.GoObject()
		}
		if row.FriendMode() == int8(friend_api.FRIEND_MODE_FRIEND) {
			count++
		}
	})

	if count >= friend_dat.FRIEND_MAX_NUM {
		out.Result = friend_api.ADD_RESULT_FAILED_ADD_FULL
		var playerFriend2 *mdb.PlayerGlobalFriend
		state.Database.AgentExecute(pid, func(db *mdb.Database) {
			db.Select.PlayerGlobalFriend(func(row *mdb.PlayerGlobalFriendRow) {
				if row.FriendPid() == state.PlayerId {
					playerFriend2 = row.GoObject()
					row.Break()
				}
			})
		})
		if playerFriend2 != nil && playerFriend != nil {
			playerFriend.FriendMode = int8(friend_api.FRIEND_MODE_STRANGE)
			playerFriend2.FriendMode = int8(friend_api.FRIEND_MODE_STRANGE)
			state.Database.Update.PlayerGlobalFriend(playerFriend)
			state.Database.AgentExecute(pid, func(db *mdb.Database) {
				db.Update.PlayerGlobalFriend(playerFriend2)
			})
		}
		return
	}

	//添加已关注玩家,返回失败
	if playerFriend != nil && (playerFriend.FriendMode == int8(friend_api.FRIEND_MODE_LISTEN_ONLY) ||
		playerFriend.FriendMode == int8(friend_api.FRIEND_MODE_FRIEND)) {
		out.Result = friend_api.ADD_RESULT_FAILED_ADD_FOLLOW
		return
	}

	count2 := 0
	var playerFriend2 *mdb.PlayerGlobalFriend
	state.Database.AgentExecute(pid, func(db *mdb.Database) {
		db.Select.PlayerGlobalFriend(func(row *mdb.PlayerGlobalFriendRow) {
			if playerFriend2 == nil && row.FriendPid() == state.PlayerId {
				playerFriend2 = row.GoObject()
			}
			if row.FriendMode() == int8(friend_api.FRIEND_MODE_FRIEND) {
				count2++
			}
		})
	})

	//对方玩家好友数量已满,返回失败
	if count2 >= friend_dat.FRIEND_MAX_NUM {
		if playerFriend != nil && playerFriend2 != nil {
			playerFriend.FriendMode = int8(friend_api.FRIEND_MODE_STRANGE)
			playerFriend2.FriendMode = int8(friend_api.FRIEND_MODE_STRANGE)
			state.Database.Update.PlayerGlobalFriend(playerFriend)
			state.Database.AgentExecute(pid, func(db *mdb.Database) {
				db.Update.PlayerGlobalFriend(playerFriend2)
			})
		}
		out.Result = friend_api.ADD_RESULT_FAILED_TARGET_FULL
		return
	}

	//添加或更改双方玩家的关注关系
	if playerFriend != nil {
		fail.When(playerFriend2 == nil, "ListenByPid playerFriend2 nil")
		if playerFriend.FriendMode == int8(friend_api.FRIEND_MODE_STRANGE) {
			playerFriend.FriendMode = int8(friend_api.FRIEND_MODE_LISTEN_ONLY)
			playerFriend2.FriendMode = int8(friend_api.FRIEND_MODE_LISTENED_ONLY)
		} else if playerFriend.FriendMode == int8(friend_api.FRIEND_MODE_LISTENED_ONLY) {
			playerFriend.FriendMode = int8(friend_api.FRIEND_MODE_FRIEND)
			playerFriend2.FriendMode = int8(friend_api.FRIEND_MODE_FRIEND)
		}

		now := time.GetNowTime()
		playerFriend.FriendTime = now
		playerFriend2.FriendTime = now

		state.Database.Update.PlayerGlobalFriend(playerFriend)
		state.Database.AgentExecute(pid, func(db *mdb.Database) {
			db.Update.PlayerGlobalFriend(playerFriend2)
		})
	} else {
		fail.When(playerFriend2 != nil, "ListenByPid playerFriend2 not nil")
		playerFriend, playerFriend2 = insertFriend(state, pid, 1)
	}

	// 通知被关注玩家状态变化
	if session, ok := module.Player.GetPlayerOnline(pid); ok {
		session.Send(&friend_api.NotifyListenedState_Out{
			Pid:   state.PlayerId,
			Nick:  state.PlayerNick,
			State: friend_api.LISTEND_STATE_LISTEND,
		})
	}

	friendInfo := global.GetPlayerInfo(pid)
	out.Result = friend_api.ADD_RESULT_SUCCEED
	out.RoleId = playerFriend.FriendRoleId
	out.Nickname = []byte(playerFriend.FriendNick)
	out.Level = friendInfo.RoleLevel
	out.FightNum = friendInfo.FightNum
	return
}

func CancelListen(state *module.SessionState, pid int64) (result bool) {
	var playerFriend *mdb.PlayerGlobalFriend
	state.Database.Select.PlayerGlobalFriend(func(row *mdb.PlayerGlobalFriendRow) {
		if row.FriendPid() == pid {
			playerFriend = row.GoObject()
			row.Break()
		}
	})
	fail.When(playerFriend == nil, "CancelListen playerFriend nil")

	//仅当删除好友时进行统计
	if playerFriend.FriendMode == int8(friend_api.FRIEND_MODE_FRIEND) {
		playerFriendState := state.Database.Lookup.PlayerGlobalFriendState(state.PlayerId)
		fail.When(playerFriendState == nil, "CancelListen playerFriendState nil")

		if !time.IsToday(playerFriendState.DeleteTime) {
			playerFriendState.DeleteDayCount = 0
		}

		fail.When(playerFriendState.DeleteDayCount >= friend_dat.FRIEND_DELETE_MAX_DAY_COUNT, "CancelListen wrong count")
		if playerFriendState.DeleteDayCount >= friend_dat.FRIEND_DELETE_MAX_DAY_COUNT {
			return false
		}

		playerFriendState.DeleteDayCount++
		playerFriendState.DeleteTime = time.GetNowTime()
		state.Database.Update.PlayerGlobalFriendState(playerFriendState)
	}

	playerFriend.FriendMode = int8(friend_api.FRIEND_MODE_STRANGE)
	state.Database.Update.PlayerGlobalFriend(playerFriend)

	result = true
	state.Database.AgentExecute(pid, func(db *mdb.Database) {
		var playerFriend *mdb.PlayerGlobalFriend
		db.Select.PlayerGlobalFriend(func(row *mdb.PlayerGlobalFriendRow) {
			if row.FriendPid() == state.PlayerId {
				playerFriend = row.GoObject()
				row.Break()
			}
		})
		fail.When(playerFriend == nil, "CancelListen playerFriend nil")
		if playerFriend == nil {
			result = false
			return
		}
		playerFriend.FriendMode = int8(friend_api.FRIEND_MODE_STRANGE)
		db.Update.PlayerGlobalFriend(playerFriend)

		// 通知被关注玩家状态变化
		if session, ok := module.Player.GetPlayerOnline(pid); ok {
			session.Send(&friend_api.NotifyListenedState_Out{
				Pid:   state.PlayerId,
				Nick:  state.PlayerNick,
				State: friend_api.LISTEND_STATE_CANCEL,
			})
		}
	})

	return result
}

func Block(state *module.SessionState, pid int64) {
	var playerFriend *mdb.PlayerGlobalFriend
	state.Database.Select.PlayerGlobalFriend(func(row *mdb.PlayerGlobalFriendRow) {
		if row.FriendPid() == pid {
			playerFriend = row.GoObject()
			row.Break()
		}
	})
	playerFriend.BlockMode = 1
	state.Database.Update.PlayerGlobalFriend(playerFriend)
}

func CancelBlock(state *module.SessionState, pid int64) {
	var playerFriend *mdb.PlayerGlobalFriend
	state.Database.Select.PlayerGlobalFriend(func(row *mdb.PlayerGlobalFriendRow) {
		if row.FriendPid() == pid {
			playerFriend = row.GoObject()
			row.Break()
		}
	})
	playerFriend.BlockMode = 0
	state.Database.Update.PlayerGlobalFriend(playerFriend)
}

//添加好友关系
//optType 1:添加关注 2:添加陌生人
//返回值: 主动关注者,被关注者
func insertFriend(state *module.SessionState, pid int64, optType int8) (playerFriend, playerFriend2 *mdb.PlayerGlobalFriend) {
	var mode1, mode2 int8

	player := global.GetPlayerInfo(state.PlayerId)
	otherPlayer := global.GetPlayerInfo(pid)

	if optType == 1 {
		mode1 = int8(friend_api.FRIEND_MODE_LISTEN_ONLY)
		mode2 = int8(friend_api.FRIEND_MODE_LISTENED_ONLY)
	} else {
		mode1 = int8(friend_api.FRIEND_MODE_STRANGE)
		mode2 = int8(friend_api.FRIEND_MODE_STRANGE)
	}

	nowTime := time.GetNowTime()
	state.Database.AgentExecute(pid, func(db *mdb.Database) {
		playerFriend2 = &mdb.PlayerGlobalFriend{
			Pid:          pid,
			FriendPid:    state.PlayerId,
			FriendNick:   string(player.PlayerNick),
			FriendRoleId: player.RoleId,
			FriendMode:   mode2,
			FriendTime:   nowTime,
		}
		db.Insert.PlayerGlobalFriend(playerFriend2)
	})

	playerFriend = &mdb.PlayerGlobalFriend{
		Pid:          state.PlayerId,
		FriendPid:    pid,
		FriendNick:   string(otherPlayer.PlayerNick),
		FriendRoleId: otherPlayer.RoleId,
		FriendMode:   mode1,
		FriendTime:   nowTime,
	}
	state.Database.Insert.PlayerGlobalFriend(playerFriend)
	return
}

func CurrentPlatformFriendNum(state *module.SessionState, friendNum int32) {
	playerFriendState := state.Database.Lookup.PlayerGlobalFriendState(state.PlayerId)
	if friendNum <= playerFriendState.PlatformFriendNum {
		return
	}
	awards := friend_dat.GetPlatformFriendAward(playerFriendState.PlatformFriendNum, friendNum)
	playerFriendState.PlatformFriendNum = friendNum
	state.Database.Update.PlayerGlobalFriendState(playerFriendState)

	for _, award := range awards {
		rpc.RemoteMailSend(state.PlayerId, mail_dat.MailPlatformFriendAward{
			FriendNum: fmt.Sprintf("%d", friendNum),
			ItemNum:   fmt.Sprintf("%d", award.Num),
			ItemName:  award.Name,
			Attachments: []*mail_dat.Attachment{
				&mail_dat.Attachment{
					AttachmentType: award.AwardType,
					ItemId:         award.AwardId,
					ItemNum:        int64(award.Num),
				},
			},
		})
	}
}

func SendHeart(state *module.SessionState, friendType friend_api.FriendType, nickname string, friendPid int64) {
	if sendHeartRecord := module.Friend.GetSendHeartRecord(state.Database, friendPid); sendHeartRecord != nil {
		//fail.When(sendHeartRecord.SendHeartTime+friend_dat.FRIEND_SEND_HEART_INTERVAL > time.GetNowTime(), "赠送爱心CD中")
		if sendHeartRecord.SendHeartTime+friend_dat.FRIEND_SEND_HEART_INTERVAL > time.GetNowTime() {
			return
		}
	}
	module.Quest.RefreshDailyQuest(state, quest_dat.DAILY_QUEST_CLASS_SEND_FRIEND_HEART)
	tlog.PlayerSystemModuleFlowLog(state.Database, tlog.SMT_FRIEND)
	if friendType == friend_api.FRIEND_TYPE_GAME_FRIEND {
		rpc.RemoteSendHeartToGameFriend(state.PlayerId, friendPid, func(isFriend bool) {
			//fail.When(!isFriend, "双方不是好友")
			if !isFriend {
				return
			}
			mdb.GlobalExecute(func(globalDB *mdb.Database) {
				globalDB.AgentExecute(state.PlayerId, func(agentDB *mdb.Database) {
					module.Friend.SendHeart(agentDB, friendPid)
					rpc.RemoteMailSend(friendPid, mail_dat.MailHeart{
						Who: nickname,
					})
					tlog.PlayerSendHeartFlowLog(agentDB, friendPid)
				})
			})
		})
	} else {
		rpc.RemoteSendHeartToPlatformFriend(state.PlayerId, friendPid, nickname, func() {
			mdb.GlobalExecute(func(globalDB *mdb.Database) {
				globalDB.AgentExecute(state.PlayerId, func(agentDB *mdb.Database) {
					module.Friend.SendHeart(state.Database, friendPid)
					tlog.PlayerSendHeartFlowLog(state.Database, friendPid)
				})
			})
		})
	}
}

func SendHeartToAllFriends(state *module.SessionState, friends *friend_api.SendHeartToAllFriends_In) {
	var (
		excludePids      = []int64{}
		CanSend     bool = true
	)
	switch friends.FriendType {
	case friend_api.FRIEND_TYPE_GAME_FRIEND:
		state.Database.Select.PlayerSendHeartRecord(func(row *mdb.PlayerSendHeartRecordRow) {
			if time.GetNowTime() < row.SendHeartTime()+friend_dat.FRIEND_SEND_HEART_INTERVAL { //if can not send heart
				excludePids = append(excludePids, row.FriendPid())
			}
		})
		rpc.RemoteSendHeartToAllGameFriends(state.PlayerId, excludePids, func(valuedFriendsList []int64) {
			if valuedFriendsList == nil {
				return
			}
			mdb.GlobalExecute(func(globalDB *mdb.Database) {
				globalDB.AgentExecute(state.PlayerId, func(agentDB *mdb.Database) {
					for _, friendPid := range valuedFriendsList {
						if state.PlayerId == friendPid && len(valuedFriendsList) == 1 {
							return
						}
						if CanSend {
							module.Quest.RefreshDailyQuest(state, quest_dat.DAILY_QUEST_CLASS_SEND_FRIEND_HEART)
							tlog.PlayerSystemModuleFlowLog(state.Database, tlog.SMT_FRIEND)
							CanSend = false
						}

						if state.PlayerId == friendPid {
							continue
						}
						module.Friend.SendHeart(agentDB, friendPid)
						playerInfo := agentDB.Lookup.Player(state.PlayerId)
						rpc.RemoteMailSend(friendPid, mail_dat.MailHeart{
							Who: playerInfo.Nick,
						})
						tlog.PlayerSendHeartFlowLog(agentDB, friendPid)
					}
				})
			})
		})
	case friend_api.FRIEND_TYPE_PLATFORM_FRIEND:
		for _, friend := range friends.Friends {
			if state.PlayerId == friend.Pid {
				continue
			}
			if sendHeartRecord := module.Friend.GetSendHeartRecord(state.Database, friend.Pid); sendHeartRecord != nil {
				if sendHeartRecord.SendHeartTime+friend_dat.FRIEND_SEND_HEART_INTERVAL > time.GetNowTime() {
					continue
				}
			}
			if CanSend {
				module.Quest.RefreshDailyQuest(state, quest_dat.DAILY_QUEST_CLASS_SEND_FRIEND_HEART)
				tlog.PlayerSystemModuleFlowLog(state.Database, tlog.SMT_FRIEND)
				CanSend = false
			}
			module.Friend.SendHeart(state.Database, friend.Pid)
			tlog.PlayerSendHeartFlowLog(state.Database, friend.Pid)
			rpc.RemoteSendHeartToPlatformFriend(state.PlayerId, friend.Pid, string(friend.Nickname), nil)
		}
	}
}
func GetSendHeartList(state *module.SessionState, out *friend_api.GetSendHeartList_Out) {
	nowTimestamp := time.GetNowTime()
	state.Database.Select.PlayerSendHeartRecord(func(row *mdb.PlayerSendHeartRecordRow) {
		timestamp := row.SendHeartTime()
		if timestamp+friend_dat.FRIEND_SEND_HEART_INTERVAL <= nowTimestamp {
			return
		}
		out.Friends = append(out.Friends, friend_api.GetSendHeartList_Out_Friends{
			Pid:           row.FriendPid(),
			SendHeartTime: timestamp,
		})
	})
}

func GetPlayerByFacebook(state *module.SessionState, in *friend_api.GetPlayerByFacebook_In, out *friend_api.GetPlayerByFacebook_Out) {
	for _, v := range in.FbInfos {
		var pid int64
		pid, ok := module.Player.GetPlayerByUsername(string(v.PlatUid))
		if ok {
			mdb.GlobalExecute(func(globalDB *mdb.Database) {
				globalDB.AgentExecute(pid, func(agentDB *mdb.Database) {
					mainRoleInfo := module.Role.GetMainRole(agentDB)
					playerInfo := module.Player.GetPlayer(agentDB)
					playerFightNum := agentDB.Lookup.PlayerFightNum(pid)
					out.Friends = append(out.Friends, friend_api.GetPlayerByFacebook_Out_Friends{
						FightNum: playerFightNum.FightNum,
						Level:    mainRoleInfo.Level,
						Nickname: []byte(playerInfo.Nick),
						Pid:      pid,
						RoleId:   mainRoleInfo.RoleId,
						Name:     v.Name,
						Id:       v.Id,
					})
				})
			})
		}
	}
}
