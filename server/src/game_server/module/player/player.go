package player

import (
	"core/fail"
	"core/log"
	"core/net"
	"core/time"
	"core/valid_str"
	"game_server/api/protocol/notify_api"
	"game_server/api/protocol/player_api"
	"game_server/config"
	"game_server/dat/chest_dat"
	"game_server/dat/item_dat"
	"game_server/dat/mail_dat"
	"game_server/dat/player_dat"
	"game_server/dat/quest_dat"
	"game_server/dat/vip_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/module/rpc"
	"game_server/tencent"
	"game_server/tlog"
	"game_server/xdlog"
	goTime "time"
)

func FromLoginPlayer(session *net.Session, username, nickname, hashcode []byte, platformType, roleId int8, unixtime int64, restore bool, in *player_api.FromPlatformLogin_In, out *player_api.FromPlatformLogin_Out) (pid int64) {
	fail.When(config.ServerCfg.EnableGlobalServer, "forbidden login")

	out.Status = player_api.LOGIN_STATUS_FAILED

	nickLen := len(nickname)
	fail.When(nickLen == 0 || nickLen > 50, "player nick length error. (empty or too long)")
	err := valid_str.StripEmoji(string(nickname))
	fail.When(err != nil, err)
	// 登陆验证需要关闭debug接口模式
	if !config.ServerCfg.IsEnableDebugAPI() {
		fail.When(int(in.Gsid) != config.ServerCfg.ServerId, "登陆gs-srv请求的gsid错误")
		fail.When(!platformLoginTest(platformType, roleId, username, nickname, hashcode, unixtime, in.WegamesPlatformUid), "FromLoginPlayer login failed. error hashcode")
	}
	user := string(username)

	state := module.State(session)

	state.Ip = session.RemoteAdd().String()

	// 支付系统需要
	state.MoneyState = &module.MoneyState{
		Openid:       string(username),
		Openkey:      string(in.Openkey),
		PayToken:     string(in.PayToken),
		Pfkey:        string(in.Pfkey),
		Zoneid:       int(in.Zoneid),
		Pf:           string(in.Pf),
		PlatformType: int8(in.PlatformType),
	}

	// tolg需要
	state.TlogClientVersion = string(in.ClientVersion)
	state.TlogSystemHardware = string(in.SystemHardware)
	state.TlogTelecomOper = string(in.TelecomOper)
	state.TlogNetwork = string(in.Network)

	// 异步检查支付token是否有效
	TencentPayTest(state.MoneyState, func(rsp *tencent.GET_BALANCE_RSP) {
		if rsp.Ret != 0 {
			log.Errorf("(gs)TencentPayTest faild. %v, %v, %v ,%v ,%v, %v, %v\n",
				state.MoneyState.Openid,
				state.MoneyState.Openkey,
				state.MoneyState.PayToken,
				state.MoneyState.Pfkey,
				state.MoneyState.Zoneid,
				state.MoneyState.PlatformType,
				rsp,
			)
			session.Send(&player_api.FromPlatformLogin_Out{Status: player_api.LOGIN_STATUS_INVALID_PAYTOKEN})
		}
	})

	var ok, firstLogin bool
	if pid, ok = GetPlayerIdByUsername(user); !ok {
		// 新玩家初次登陆
		state.Database.Init()
		player := &mdb.Player{
			User: user,
			Cid:  int64(in.ChannelId),
		}
		state.Database.Insert.Player(player)

		pid = player.Id
		state.PlayerId = player.Id

		if ok, _ = InitPlayer(session, string(nickname), roleId); !ok {
			fail.When(true, "[InitPlayer] never happened.")
		}

		firstLogin = true
		state.Database.AddTLog(tlog.NewPlayerRegister(user, state.TlogClientVersion, "" /*SystemSoftware*/, state.TlogSystemHardware, state.TlogTelecomOper, state.TlogNetwork,
			0, 0, 0.0, in.Channel, "" /*CpuHardware*/, 0, "" /*GLRender*/, "" /*GLVersion*/, "" /*DeviceId*/))
		xdlog.PlayerCreateLog(state.Database, state.Ip, int32(in.ChannelId))
	} else {
		// 处理玩家存在多处登陆的情况
		if oldSession, ok := module.Player.GetPlayerOnline(pid); ok {
			if oldSession != session {
				// 当之前的连接成功断开后通知新连接可以登陆了
				oldState := module.State(oldSession)
				oldState.SendReLoginCallback = SendReLogin(session)

				out.Status = player_api.LOGIN_STATUS_WAIT_CLOSE

				// 尽量让旧的客户端收到登出通知,5s后才断开
				oldSession.Send(&notify_api.SendLogout_Out{})
				goTime.AfterFunc(5*goTime.Second, func() {
					defer func() {
						err := recover()
						if err != nil {
							log.Debugf("%d Close old session err %#v\n", pid, err)
						}

					}()
					oldSession.Close()
				})

				// 两台设备互踢，清理旧玩家状态的计数器，(新玩家状态可能restore旧玩家状态)
				oldState.RspQ.ReqCounter = 0
				oldState.RspQ.RspCounter = 0
			}
			return
		}

		// 客户端发起恢复会话状态请求
		if restore {
			newState := module.ReStoreSessionState(pid)
			if newState != nil {
				newState.MoneyState = state.MoneyState
				session.State = newState

				// 重发上次未成功发送的数据
				if !newState.RspQ.ReSend(session, in.RecvCount) {
					out.Status = player_api.LOGIN_STATUS_RELOGIN
					return
				}

				LoginDone(session, in.Channel, out)
				if out.Status == player_api.LOGIN_STATUS_SUCCEED {
					out.Status = player_api.LOGIN_STATUS_RESTORE_SUCCEED
				}
				return
			}
		}
	}

	state.PlayerId = pid
	state.CId = int32(in.ChannelId)
	state.RoleId = roleId
	state.PlayerNick = nickname
	state.WegamesPlatformUid = in.WegamesPlatformUid

	// 保证在TryDropSessionState之前mount玩家数据
	state.Database.Mount(pid, int32(in.ChannelId))

	// 丢弃不想恢复但已保存的会话状态
	module.TryDropSessionState(pid, state)

	LoginDone(session, in.Channel, out)

	// 登陆成功后发送欢迎邮件
	if firstLogin && out.Status == player_api.LOGIN_STATUS_SUCCEED {
		// 发送新手礼包邮件(取消)
		//rpc.RemoteMailSend(pid, mail_dat.MailNewbieGift{})
	}
	return
}

func LoginDone(session *net.Session, channel int32, out *player_api.FromPlatformLogin_Out) {
	state := module.State(session)

	// 检查账号是否被冻结
	playerState := state.Database.Lookup.PlayerState(state.PlayerId)
	if playerState.BanEndTime == 0 || playerState.BanEndTime > time.GetNowTime() {
		out.Status = player_api.LOGIN_STATUS_BAN
		out.BanTime = playerState.BanEndTime
		return
	}

	SetPlayerOnline(state.CId, state.PlayerId, session)
	playerInfo := state.Database.Lookup.PlayerInfo(state.PlayerId)

	//发送每日赠送爱心邮件邮件
	if !time.IsToday(playerInfo.LastLoginTime) {
		vipHeart := module.VIP.PrivilegeTimes(state, vip_dat.AIXINFULI)
		if vipHeart > 0 {
			rpc.RemoteMailSend(state.PlayerId, mail_dat.MailVIPHeart{
				Attachments: []*mail_dat.Attachment{
					&mail_dat.Attachment{
						AttachmentType: mail_dat.ATTACHMENT_HEART,
						ItemId:         0,
						ItemNum:        int64(vipHeart),
					},
				},
			})
		}
	}

	if !time.IsToday(playerInfo.LastLoginTime) {
		module.Quest.RefreshLoginExtendQuest(state, time.IsToday(playerInfo.LastLoginTime+86400))
	}

	//登陆成功存储登录时间
	playerInfo.LastLoginTime = time.GetNowTime()
	state.Database.Update.PlayerInfo(playerInfo)

	//加载QQ会员服务状态
	state.TencentState = &module.TencentState{
		QqVipStatus: module.Player.QQVipStatus(state),
	}
	//QQ特权礼包发放检查
	module.Event.CheckQQVIPGift(state.Database, state.TencentState.QqVipStatus)

	role := module.Role.GetMainRole(state.Database)
	//刷新玩家等级相关功能权值
	module.Player.RefreshPlayerLevelFuncKey(state.Database, role.Level)

	module.Quest.StartDailyQuestTimer(state)
	module.SwordSoul.SendSwordSoulWhenMaxDrawNum(session)
	//删除过期时装
	module.Fashion.LoginUpdateFashion(state.Database)
	//登录刷新打坐状态
	module.Meditation.LoginUpdateMeditationState(session)

	//加载玩家的所有活动状态
	module.InitPlayerEventsStatus(session)

	//加载玩家所有的json活动状态
	module.InitPlayerJsonEventsStatus(session)

	//加载玩家所有的天书记录
	module.InitPlayerSealedBookRecord(session)

	//修复玩家可能丢失的功能权值
	updatePlayerFuncKey(state)

	//台湾wegames每日月卡奖励
	//checkMonthCardAward(state.Database)

	if len(state.WegamesPlatformUid) > 0 {
		playerWegamesUid := state.Database.Lookup.PlayerWegamesUid(state.PlayerId)
		if playerWegamesUid == nil {
			playerWegamesUid = &mdb.PlayerWegamesUid{
				Pid:                state.PlayerId,
				WegamesPlatformUid: string(state.WegamesPlatformUid),
			}
			state.Database.Insert.PlayerWegamesUid(playerWegamesUid)
		}
	}

	// 登陆成功. 通知客户端可以连接互动服务器
	session.Send(&player_api.NotifyGlobalServerInfo_Out{
		Gsid: int32(config.ServerCfg.GlobalServerId),
		Addr: []byte(config.ServerCfg.GlobalServerAddr),
	})

	player := state.Database.Lookup.Player(state.PlayerId)
	rpc.RemoteTssUserLogin(player.User, player.Nick, "", player.Id, int(config.ServerCfg.XGSdk.PlatformType), 0)

	out.Status = player_api.LOGIN_STATUS_SUCCEED

	state.Database.AddTLog(tlog.NewPlayerLogin(player.User, int32(role.Level), 0, state.TlogClientVersion, "" /*SystemSoftware*/, state.TlogSystemHardware,
		state.TlogTelecomOper, state.TlogNetwork, 0, 0, 0.0, channel, "" /*CpuHardware*/, 0, "" /*GLRender*/, "" /*GLVersion*/, "" /*DeviceId*/))
	xdlog.PlayerLoginLog(state.Database, state.Ip, xdlog.LOGIN)
}

func GetPlayerInfo(session *net.Session, info *player_api.Info_Out) bool {
	state := module.State(session)
	if _, ok := module.Player.GetPlayerOnline(state.PlayerId); !ok {
		return false
	}

	town := state.Database.Lookup.PlayerTown(state.PlayerId)
	player := state.Database.Lookup.Player(state.PlayerId)
	role := state.Database.Lookup.PlayerRole(player.MainRoleId)

	mission := state.Database.Lookup.PlayerMission(player.Id)
	mission_level := state.Database.Lookup.PlayerMissionLevel(player.Id)
	playerInfo := state.Database.Lookup.PlayerInfo(state.PlayerId)
	playerPhysical := state.Database.Lookup.PlayerPhysical(state.PlayerId)
	playerFuncKey := state.Database.Lookup.PlayerFuncKey(state.PlayerId)
	playerHeart := state.Database.Lookup.PlayerHeart(state.PlayerId)
	playerCoinsInfo := state.Database.Lookup.PlayerCoins(state.PlayerId)
	playerVIPInfo := state.Database.Lookup.PlayerVip(state.PlayerId)
	playerHardLevelLock := state.Database.Lookup.PlayerHardLevel(state.PlayerId)
	playerFame := state.Database.Lookup.PlayerFame(state.PlayerId)

	// 更新体力
	module.Physical.LoginUpdatePhysical(state, playerPhysical)
	// 更新爱心数
	module.Heart.LoginUpdateHeartNum(state, playerHeart)
	//更新购买铜币次数
	playerCoinsInfo = updateCoinsBuyCount(state, playerCoinsInfo)

	// 更新玩家每日任务
	module.Quest.UpdatePlayerDailyQuest(state.Database, role.Level)

	info.Nickname = []byte(player.Nick)
	info.TownId = town.TownId
	info.TownLock = town.Lock

	info.RoleId = role.RoleId
	info.RoleLv = role.Level
	info.RoleExp = role.Exp
	info.MissionKey = int32(module.Item.GetItemNum(state.Database, item_dat.ITEM_MISSION_KEY_ID)) /*mission.Key (光明钥匙已改为物品)*/
	info.Physical = playerPhysical.Value
	info.PhysicalBuyCount = int16(playerPhysical.DailyCount)
	info.CoinsBuyCount = playerCoinsInfo.DailyCount
	info.BatchBoughtCoins = int16(playerCoinsInfo.BatchBought)
	info.CornucopiaCount = module.Item.CornucopiaCount(state.Database)
	info.VipLevel = playerVIPInfo.Level
	info.Fame = playerFame.Fame

	info.FuncKey = playerFuncKey.Key
	info.PlayedKey = playerFuncKey.PlayedKey

	info.MissionMaxOrder = mission.MaxOrder
	info.MissionLevelMaxLock = mission_level.MaxLock
	info.MissionLevelAwardLock = mission_level.AwardLock
	info.HardLevelLock = playerHardLevelLock.Lock

	info.Ingot = module.Player.GetIngot(state.Database, state.MoneyState)
	info.Coins = playerInfo.Coins
	info.Fragments = playerInfo.SwordSoulFragment
	info.HeartNum = playerHeart.Value

	playerQuest := state.Database.Lookup.PlayerQuest(state.PlayerId)
	info.QuestId = playerQuest.QuestId
	info.QuestState = playerQuest.State

	info.MonthCardExpireTimestamp = getMontCardExpireTimestamp(state.Database)
	if playerVIPInfo.Ingot == 0 {
		info.NeverRecharge = true
	}

	if playerQuest.State == quest_dat.QUEST_STATUS_AWARD {
		info.QuestId = quest_dat.QUEST_ID_NO_MORE
		info.QuestState = quest_dat.QUEST_STATUS_ALREADY_RECEIVE //客户端强行要求使用这个常量
	}

	info.ArenaReportNum = playerInfo.NewArenaReportNum

	if playerInfo.NewArenaReportNum > 0 {
		playerInfo.NewArenaReportNum = 0
		state.Database.Update.PlayerInfo(playerInfo)
	}

	info.QqVipStatus = state.TencentState.QqVipStatus
	info.ServerOpenTime = config.ServerCfg.ServerOpenTime

	if module.Player.IsOpenFunc(state.Database, player_dat.FUNC_FATE_BOX) {
		playerFateBoxState := state.Database.Lookup.PlayerFateBoxState(state.PlayerId)

		//权值需要从高到低，才能遍历到所有宝箱时间
		switch playerFateBoxState.Lock {
		case chest_dat.HunyuanBoxLock:
			info.NextFreeHunyuanBoxTimestamp = playerFateBoxState.HunyuanBoxFreeDrawTimestamp + chest_dat.FREE_FATE_BOX_CD
			fallthrough
		case chest_dat.SunBoxLock:
			info.NextFreeSunBoxTimestamp = playerFateBoxState.SunBoxFreeDrawTimestamp + chest_dat.FREE_FATE_BOX_CD
			fallthrough
		case chest_dat.MoonBoxLock:
			info.NextFreeMoonBoxTimestamp = playerFateBoxState.MoonBoxFreeDrawTimestamp + chest_dat.FREE_FATE_BOX_CD
			fallthrough
		case chest_dat.StarBoxLock:
			info.NextFreeStarBoxTimestamp = playerFateBoxState.StarBoxFreeDrawTimestamp + chest_dat.FREE_FATE_BOX_CD

		}
	}

	// 检查是否需要重构关卡
	module.Mission.CheckIsNeedRebuildMissionLevel(session)

	// 尝试重建状态
	module.TryRebuildState(session)

	return true
}

//增加新功能是，老玩家可能错过了获取功能权值的时机（目前是任务）
//故需在登录是检查是否需要为已通过任务，但未开启功能权值的玩家做修复
func updatePlayerFuncKey(state *module.SessionState) {
	playerQuest := state.Database.Lookup.PlayerQuest(state.PlayerId)
	module.Quest.RefreshExtendQuest(state.Database, playerQuest.QuestId)

	quest := quest_dat.GetQuestById(playerQuest.QuestId)
	if playerQuest.State == quest_dat.QUEST_STATUS_AWARD {
		// 玩家完成了所有任务的状态, 尝试去更新新的任务
		module.Quest.RefreshQuest(state, quest_dat.QUEST_TYPE_NONE, xdlog.ET_QUEST)
	} else if playerQuest.State == quest_dat.QUEST_STATUS_NO_RECEIVE {
		//客户端因为意外在升级时没有刷新任务，那么在登录时刷新
		mainRoleLv := int32(module.Role.GetMainRole(state.Database).Level)
		if mainRoleLv >= quest.RequireLevel {
			module.Quest.RefreshQuest(state, quest_dat.QUEST_TYPE_NONE, xdlog.ET_QUEST)
		}
	}

	//目前只有任务回产出功能权值
	maxLock := player_dat.GetMaxLockByQuestOrder(quest.Order)

	module.Player.RefreshPlayerFuncKey(state.Database, maxLock)

	//修复玩家关卡权值
	var missionLevelLock int32
	if playerQuest.State == quest_dat.QUEST_STATUS_AWARD {
		missionLevelLock = quest_dat.GetMissioLevelLockByMainQuest(quest.NextQuestId)
	} else {
		missionLevelLock = quest_dat.GetMissioLevelLockByMainQuest(playerQuest.QuestId)
	}
	playerMissionLevel := state.Database.Lookup.PlayerMissionLevel(state.PlayerId)
	if playerMissionLevel.Lock < missionLevelLock {
		module.Mission.SetMissionLevelLock(state, missionLevelLock, playerMissionLevel, true)
	}

}
