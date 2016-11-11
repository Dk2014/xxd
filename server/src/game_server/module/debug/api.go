package debug

import (
	"fmt"
	"time"

	"core/debug"
	"core/fail"
	"core/log"
	"core/net"
	core_time "core/time"
	"game_server/api/protocol/battle_api"
	"game_server/api/protocol/debug_api"
	"game_server/api/protocol/notify_api"
	"game_server/battle"
	"game_server/config"
	"game_server/dat/announcement_dat"
	"game_server/dat/battle_pet_dat"
	"game_server/dat/channel_dat"
	"game_server/dat/clique_dat"
	"game_server/dat/daily_sign_in_dat"
	"game_server/dat/driving_sword_dat"
	"game_server/dat/event_dat"
	"game_server/dat/item_dat"
	"game_server/dat/mail_dat"
	"game_server/dat/mission_dat"
	"game_server/dat/multi_level_dat"
	"game_server/dat/player_dat"
	"game_server/dat/quest_dat"
	"game_server/dat/rainbow_dat"
	"game_server/dat/role_dat"
	"game_server/dat/team_dat"
	"game_server/dat/town_dat"
	"game_server/dat/trader_dat"
	"game_server/dat/vip_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/module/rpc"
	"game_server/tencent"
	"game_server/tlog"
)

func init() {
	debug_api.SetInHandler(DebugAPI{})
}

type DebugAPI struct{}

func (debug DebugAPI) AddBuddy(session *net.Session, in *debug_api.AddBuddy_In) {
	state := module.State(session)
	role_dat.GetRoleLevelInfo(in.RoleId, 10)
	module.Role.AddBuddyRole(state, in.RoleId, int16(1))
	session.Send(&debug_api.AddBuddy_Out{})
}

func (debug DebugAPI) AddItem(session *net.Session, in *debug_api.AddItem_In) {
	state := module.State(session)
	itemId := in.ItemId
	item := item_dat.GetItem(itemId)
	/*
		//重置送爱心时间
		state.Database.Select.PlayerSendHeartRecord(func(row *mdb.PlayerSendHeartRecordRow) {
			sendHeartRecord := row.GoObject()
			sendHeartRecord.SendHeartTime -= 86400
			state.Database.Update.PlayerSendHeartRecord(sendHeartRecord)
		})
	*/
	nowTime := core_time.GetNowTime()
	fmt.Println("nowTime=", nowTime)
	switch item.TypeId {
	case item_dat.TYPE_WEAPON, item_dat.TYPE_CLOTHES, item_dat.TYPE_SHOE, item_dat.TYPE_ACCESSORIES:
		// 装备
		for i := 0; i < int(in.Number); i++ {
			module.Item.AddItem(state.Database, in.ItemId, 1, tlog.IFR_DEBUG_ADD_ITEM, 0, "")
		}
	default:
		module.Item.AddItem(state.Database, in.ItemId, in.Number, tlog.IFR_DEBUG_ADD_ITEM, 0, "")
	}

	session.Send(&debug_api.AddItem_Out{})
}

func (debug DebugAPI) SetRoleLevel(session *net.Session, in *debug_api.SetRoleLevel_In) {
	fail.When(in.Level <= 0 || in.Level > role_dat.MAX_ROLE_LEVEL, "角色等级超过上限")
	state := module.State(session)
	var role *mdb.PlayerRole

	state.Database.Select.PlayerRole(func(row *mdb.PlayerRoleRow) {
		if row.RoleId() == in.RoleId {
			role = row.GoObject()
			row.Break()
		}
	})
	if role.Level >= in.Level {
		session.Send(&debug_api.SetRoleLevel_Out{})
		return
	}
	//如果是主角的话
	if in.RoleId == 1 || in.RoleId == 2 {
		//如果等级大过25且没有伙伴则添加燕无名
		if in.Level >= 25 {
			roleNum := 0
			state.Database.Select.PlayerRole(func(row *mdb.PlayerRoleRow) {
				roleNum++
			})
			if roleNum == 1 {
				module.Role.AddBuddyRole(state, 6, 1)
				playerForm := state.Database.Lookup.PlayerFormation(state.PlayerId)
				playerForm.Pos0 = int8(module.Role.GetMainRole(state.Database).RoleId) //主角在0号位
				playerForm.Pos1 = team_dat.POS_NO_ROLE
				playerForm.Pos2 = 6
				playerForm.Pos3 = team_dat.POS_NO_ROLE
				playerForm.Pos4 = team_dat.POS_NO_ROLE
				playerForm.Pos5 = team_dat.POS_NO_ROLE
				state.Database.Update.PlayerFormation(playerForm)
			}
		}
		//刷新等级功能
		module.Player.RefreshPlayerLevelFuncKey(state.Database, in.Level)
		//更新等级活动的状态
		module.Event.LevelUp(state.Database, 1, in.Level)
	}

	playerFame := state.Database.Lookup.PlayerFame(state.PlayerId)
	module.Skill.UpdateSkill(state.Database, role.RoleId, int16(role.FriendshipLevel), playerFame.Level, in.Level)
	role.Level = in.Level
	state.Database.Update.PlayerRole(role)
	module.Notify.SendRoleExpChanged(session, in.RoleId, 0, role.Exp, role.Level)
	session.Send(&debug_api.SetRoleLevel_Out{})
}

func (debug DebugAPI) SetCoins(session *net.Session, in *debug_api.SetCoins_In) {
	state := module.State(session)
	module.Player.IncMoney(state.Database, state.MoneyState, in.Number, player_dat.COINS, -1, 0, "")
	session.Send(&debug_api.SetCoins_Out{})
}

func (debug DebugAPI) SetIngot(session *net.Session, in *debug_api.SetIngot_In) {
	state := module.State(session)
	module.Player.IncMoney(state.Database, state.MoneyState, in.Number, player_dat.INGOT, -1, 0, "")
	session.Send(&debug_api.SetIngot_Out{})
}

func (debug DebugAPI) AddGhost(session *net.Session, in *debug_api.AddGhost_In) {
	state := module.State(session)
	module.Ghost.AddGhost(state, in.GhostId, -1, 0)
	session.Send(&debug_api.AddGhost_Out{})
}

func (debug DebugAPI) SetPlayerPhysical(session *net.Session, in *debug_api.SetPlayerPhysical_In) {
	state := module.State(session)
	pp := state.Database.Lookup.PlayerPhysical(state.PlayerId)
	pp.Value = in.Physical

	state.Database.Update.PlayerPhysical(pp)
	module.Notify.SendPhysicalChange(session, pp.Value)
	session.Send(&debug_api.AddGhost_Out{})
}

func (debug DebugAPI) ResetLevelEnterCount(session *net.Session, in *debug_api.ResetLevelEnterCount_In) {
	state := module.State(session)
	var levelRecord *mdb.PlayerMissionLevelRecord
	state.Database.Select.PlayerMissionLevelRecord(func(row *mdb.PlayerMissionLevelRecordRow) {
		if row.MissionLevelId() == in.LevelId {
			levelRecord = row.GoObject()
			row.Break()
		}
	})
	if levelRecord != nil {
		levelRecord.LastEnterTime = 0
		state.Database.Update.PlayerMissionLevelRecord(levelRecord)
	}
	session.Send(&debug_api.ResetLevelEnterCount_Out{})
}

func (debug DebugAPI) AddExp(session *net.Session, in *debug_api.AddExp_In) {
	state := module.State(session)
	module.Role.AddRoleExp(state.Database, in.RoleId, in.AddExp, state.RoleId, -1)
	session.Send(&debug_api.AddBuddy_Out{})
}

func (debug DebugAPI) OpenGhostMission(session *net.Session, in *debug_api.OpenGhostMission_In) {
	// TODO
	//state := module.State(session)
	//module.Ghost.OpenGhostMissionById(state, in.MissionId)
	//session.Send(&debug_api.OpenGhostMission_Out{})
}

func (debug DebugAPI) SendMail(session *net.Session, in *debug_api.SendMail_In) {
	state := module.State(session)
	mail := mail_dat.NewMailTemplete(in.MailId)
	rpc.RemoteMailSend(state.PlayerId, mail)
	session.Send(&debug_api.SendMail_Out{})
}

func (debug DebugAPI) ClearMail(session *net.Session, in *debug_api.ClearMail_In) {
	state := module.State(session)

	state.Database.Select.PlayerMail(func(row *mdb.PlayerMailRow) {
		if row.HaveAttachment() == 1 {
			state.Database.Select.PlayerMailAttachment(func(attachRow *mdb.PlayerMailAttachmentRow) {
				if attachRow.PlayerMailId() == row.Id() {
					state.Database.Delete.PlayerMailAttachment(attachRow.GoObject())
				}
			})
		}
		state.Database.Delete.PlayerMail(row.GoObject())
	})
}

func (debug DebugAPI) OpenMissionLevel(session *net.Session, in *debug_api.OpenMissionLevel_In) {
	state := module.State(session)

	levelInfo := mission_dat.GetMissionLevelById(in.LevelId)
	mission_info := mission_dat.GetMissionById(levelInfo.MissionId)

	playerLevel := state.Database.Lookup.PlayerMissionLevel(state.PlayerId)
	if playerLevel.MaxLock < levelInfo.Lock {
		module.Mission.SetMissionLevelLock(state, levelInfo.Lock, nil, true)
	}

	player_mission := state.Database.Lookup.PlayerMission(state.PlayerId)
	if player_mission.MaxOrder < mission_info.Order {
		player_mission.MaxOrder = mission_info.Order
		state.Database.Update.PlayerMission(player_mission)
		module.Notify.SendPlayerKeyChanged(session, 0, player_mission.MaxOrder)
	}

	alreadyOpen := false
	state.Database.Select.PlayerMissionRecord(func(row *mdb.PlayerMissionRecordRow) {
		if row.MissionId() == levelInfo.MissionId {
			alreadyOpen = true
			row.Break()
		}
	})

	if !alreadyOpen {
		state.Database.Insert.PlayerMissionRecord(&mdb.PlayerMissionRecord{
			Pid:       state.PlayerId,
			TownId:    mission_info.TownId,
			MissionId: levelInfo.MissionId,
			OpenTime:  time.Now().Unix(),
		})
	}
	levelRecord := module.Mission.GetMissionLevelRecord(state.Database, in.LevelId)
	if levelRecord == nil {
		state.Database.Insert.PlayerMissionLevelRecord(&mdb.PlayerMissionLevelRecord{
			Pid:            state.PlayerId,
			MissionId:      levelInfo.MissionId,
			MissionLevelId: in.LevelId,
		})
	}
}

func (debug DebugAPI) StartBattle(session *net.Session, in *debug_api.StartBattle_In) {
	state := module.State(session)
	state.MissionLevelState = module.NewMissionLevelState(0, 0)
	state.MissionLevelState.LoadBuddySkill(state)
	fail.When(in.BattleType != int8(battle_api.BATTLE_TYPE_MISSION), "wront battle type")
	state.Battle = module.Battle.NewBattleForLevel(session, battle.BT_MISSION_LEVEL, in.EnemyId, true)
}

func (debug DebugAPI) ListenByName(session *net.Session, in *debug_api.ListenByName_In) {
	state := module.State(session)
	module.Friend.ListenByName(state, string(in.Name))
}

func (debug DebugAPI) OpenQuest(session *net.Session, in *debug_api.OpenQuest_In) {
	state := module.State(session)
	playerQuest := state.Database.Lookup.PlayerQuest(state.PlayerId)

	playerQuest.QuestId = in.QuestId
	playerQuest.State = quest_dat.QUEST_STATUS_ALREADY_RECEIVE

	module.Notify.SendQuestChange(session, playerQuest.QuestId, playerQuest.State)
	state.Database.Update.PlayerQuest(playerQuest)
}

func (debug DebugAPI) OpenFunc(session *net.Session, in *debug_api.OpenFunc_In) {
	module.Player.RefreshPlayerFuncKey(module.State(session).Database, in.Lock)
}

func (debug DebugAPI) AddSwordSoul(session *net.Session, in *debug_api.AddSwordSoul_In) {
	state := module.State(session)
	module.SwordSoul.AddSwordSoul(state, in.SwordSoulId, -1)
}

func (debug DebugAPI) AddBattlePet(session *net.Session, in *debug_api.AddBattlePet_In) {
	state := module.State(session)
	petDat := battle_pet_dat.GetBattlePetWithPetId(in.PetId)
	module.BattlePet.AddPet(state.Database, petDat.PetId, tlog.IFR_DEBUG_ADD_ITEM, 0)
}

//重置进入多人关卡次数
func (debug DebugAPI) ResetMultiLevelEnterCount(session *net.Session, in *debug_api.ResetMultiLevelEnterCount_In) {
	state := module.State(session)
	playerMultiLevel := state.Database.Lookup.PlayerMultiLevelInfo(state.PlayerId)
	playerMultiLevel.DailyNum = 0
	state.Database.Update.PlayerMultiLevelInfo(playerMultiLevel)
}

//开启指定多人关卡
func (debug DebugAPI) OpenMultiLevel(session *net.Session, in *debug_api.OpenMultiLevel_In) {
	state := module.State(session)

	//获取玩家主角
	mainRole := module.Role.GetMainRole(state.Database)

	//获取多人关卡信息
	multiLevelInfo := multi_level_dat.GetMultiLevelById(in.LevelId)
	fail.When(multiLevelInfo == nil, "multi_level_missoin not exist")

	if mainRole.Level < multiLevelInfo.RequireLevel {
		mainRole.Level = multiLevelInfo.RequireLevel
	}
	state.Database.Update.PlayerRole(mainRole)

	//获取玩家多人关卡信息
	playerMultiLevelInfo := state.Database.Lookup.PlayerMultiLevelInfo(state.PlayerId)
	var needInsert bool
	if playerMultiLevelInfo == nil {
		needInsert = true
		playerMultiLevelInfo = &mdb.PlayerMultiLevelInfo{
			Pid: state.PlayerId,
		}
	}

	playerMultiLevelInfo.Lock = multiLevelInfo.Lock
	if needInsert {
		state.Database.Insert.PlayerMultiLevelInfo(playerMultiLevelInfo)
	} else {
		state.Database.Update.PlayerMultiLevelInfo(playerMultiLevelInfo)
	}
}

//产生一条公告
func (debug DebugAPI) CreateAnnouncement(session *net.Session, in *debug_api.CreateAnnouncement_In) {
	newAnc := new(announcement_dat.AnnounceTestAnnounce3)
	now := time.Now().Unix()
	rpc.RemoteGlobalAnnouncementCreate(newAnc.GetAnnouncementTplId(), newAnc.GetParameters(), "", now+5, now+100, 60)
}

func (debug DebugAPI) OpenAllPetGrid(session *net.Session, in *debug_api.OpenAllPetGrid_In) {
	//废弃
	// 临时用来做每日首冲活动测试 可随时去除
	state := module.State(session)
	eventId := int16(event_dat.EVENT_FIRST_RECHARGE_DAILY)
	eventInfo, _ := event_dat.GetEventInfoById(eventId)
	eventInfo.Start -= 3600 * 24
	eventInfo.End -= 3600 * 24
	eventInfo.Dispose -= 3600 * 24
	state.EventsState.ClearState(state.Database, eventId)
}

func (debug DebugAPI) AddHeart(session *net.Session, in *debug_api.AddHeart_In) {
	state := module.State(session)
	module.Heart.IncHeart(state, in.Number)
	session.Send(&debug_api.AddHeart_Out{})
}

func (debug DebugAPI) ResetHardLevelEnterCount(session *net.Session, in *debug_api.ResetHardLevelEnterCount_In) {
	state := module.State(session)
	var records []*mdb.PlayerHardLevelRecord
	state.Database.Select.PlayerHardLevelRecord(func(row *mdb.PlayerHardLevelRecordRow) {
		records = append(records, row.GoObject())
	})
	for _, record := range records {
		record.DailyNum = 0
		state.Database.Update.PlayerHardLevelRecord(record)
	}
}

func (debug DebugAPI) OpenHardLevel(session *net.Session, in *debug_api.OpenHardLevel_In) {
	state := module.State(session)
	levelDat := mission_dat.GetMissionLevelById(in.LevelId)
	hardLevelDat := mission_dat.GetHardLevelInfo(levelDat.ParentId)
	playerHardLevel := state.Database.Lookup.PlayerHardLevel(state.PlayerId)
	playerHardLevel.Lock = hardLevelDat.HardLevelLock
	state.Database.Update.PlayerHardLevel(playerHardLevel)
	module.Notify.SendHardLevelLockChange(session, hardLevelDat.HardLevelLock)
}

func (debug DebugAPI) SetVipLevel(session *net.Session, in *debug_api.SetVipLevel_In) {
	state := module.State(session)
	ingot, exist := vip_dat.GetVIPLevelInfo(in.Level)
	if !exist {
		fail.When(true, "VIP level not exist")
	}
	module.VIP.UpdateIngot(state.Database, ingot)
	session.Send(&debug_api.SetIngot_Out{})
}

var defaultExtendLevelCoinLevelOpenDay = mission_dat.ExtendLevelCoinOpenDay
var defaultExtendLevelExpLevelOpenDay = mission_dat.ExtendLevelExpOpenDay
var defaultExtendLevelBuddyOpenDay = mission_dat.ExtendLevelBuddyOpenDay
var defaultExtendLevelPetOpenDay = mission_dat.ExtendLevelPetOpenDay
var defaultExtendLevelGhostOpenDay = mission_dat.ExtendLevelGhostOpenDay

func (debug DebugAPI) SetResourceLevelOpenDay(session *net.Session, in *debug_api.SetResourceLevelOpenDay_In) {
	switch in.LevelType {
	case mission_dat.RESOURCE_COIN_LEVEL:
		mission_dat.ExtendLevelCoinOpenDay = append(mission_dat.ExtendLevelCoinOpenDay, int(in.OpenDay))
	case mission_dat.RESOURCE_EXP_LEVEL:
		mission_dat.ExtendLevelExpOpenDay = append(mission_dat.ExtendLevelExpOpenDay, int(in.OpenDay))
	case mission_dat.EXTEND_LEVEL_TYPE_BUDDY:
		mission_dat.ExtendLevelBuddyOpenDay = append(mission_dat.ExtendLevelBuddyOpenDay, int(in.OpenDay))
	case mission_dat.EXTEND_LEVEL_TYPE_PET:
		mission_dat.ExtendLevelPetOpenDay = append(mission_dat.ExtendLevelPetOpenDay, int(in.OpenDay))
	case mission_dat.EXTEND_LEVEL_TYPE_GHOST:
		mission_dat.ExtendLevelGhostOpenDay = append(mission_dat.ExtendLevelGhostOpenDay, int(in.OpenDay))
	}
}

func (debug DebugAPI) ResetResourceLevelOpenDay(session *net.Session, in *debug_api.ResetResourceLevelOpenDay_In) {
	mission_dat.ExtendLevelCoinOpenDay = defaultExtendLevelCoinLevelOpenDay
	mission_dat.ExtendLevelExpOpenDay = defaultExtendLevelExpLevelOpenDay
	mission_dat.ExtendLevelBuddyOpenDay = defaultExtendLevelBuddyOpenDay
	mission_dat.ExtendLevelPetOpenDay = defaultExtendLevelPetOpenDay
	mission_dat.ExtendLevelGhostOpenDay = defaultExtendLevelGhostOpenDay
}

func (debug DebugAPI) ResetArenaDailyCount(session *net.Session, in *debug_api.ResetArenaDailyCount_In) {
	state := module.State(session)
	playerArena := state.Database.Lookup.PlayerArena(state.PlayerId)
	playerArena.DailyNum = 0
	state.Database.Update.PlayerArena(playerArena)
	session.Send(&debug_api.ResetArenaDailyCount_Out{})
}

func (debug DebugAPI) ResetSwordSoulDrawCd(session *net.Session, in *debug_api.ResetSwordSoulDrawCd_In) {
	//state := module.State(session)
	//playerSwordSoulState := state.Database.Lookup.PlayerSwordSoulState(state.PlayerId)
	//playerSwordSoulState.UpdateTime = time.Now().Unix() - sword_soul_dat.CD_TIME
	//state.Database.Update.PlayerSwordSoulState(playerSwordSoulState)
}

func (debug DebugAPI) SetFirstLoginTime(session *net.Session, in *debug_api.SetFirstLoginTime_In) {
	state := module.State(session)
	playerInfo := state.Database.Lookup.PlayerInfo(state.PlayerId)
	playerInfo.FirstLoginTime = in.Timestamp
	state.Database.Update.PlayerInfo(playerInfo)
}

//将玩家首次登录日期
//同时签到记录里面的日期需要回退一天
func (debug DebugAPI) EarlierFirstLoginTime(session *net.Session, in *debug_api.EarlierFirstLoginTime_In) {
	state := module.State(session)
	playerInfo := state.Database.Lookup.PlayerInfo(state.PlayerId)
	playerInfo.FirstLoginTime -= 86400
	state.Database.Update.PlayerInfo(playerInfo)
	playerSignInState := state.Database.Lookup.PlayerDailySignInState(state.PlayerId)
	if playerSignInState.UpdateTime > 0 {
		playerSignInState.UpdateTime -= 86400
		state.Database.Update.PlayerDailySignInState(playerSignInState)
	}
	daily_sign_in_dat.ServerOpenTime--
	session.Send(&debug_api.EarlierFirstLoginTime_Out{})
}

//从配置文件重新加载服务器开服时间
func (debug DebugAPI) ResetServerOpenTime(session *net.Session, in *debug_api.ResetServerOpenTime_In) {
	daily_sign_in_dat.ServerOpenTime = core_time.GetNowDayFromUnix(config.ServerCfg.ServerOpenTime)
	session.Send(&debug_api.ResetServerOpenTime_Out{})
}

//清空商人出现时间
func (debug DebugAPI) ClearTraderSchedule(session *net.Session, in *debug_api.ClearTraderSchedule_In) {
	trader_dat.MapTraderSchedule[in.TraderId] = []*trader_dat.TraderSchedule{}
	session.Send(&debug_api.ClearTraderSchedule_Out{})
}

//添加商人出现时间
func (debug DebugAPI) AddTraderSchedule(session *net.Session, in *debug_api.AddTraderSchedule_In) {
	traderId := in.TraderId
	schedule := &trader_dat.TraderSchedule{
		Expire:    in.Expire,
		Showup:    in.Showup,
		Disappear: in.Disappear,
	}
	trader_dat.MapTraderSchedule[traderId] = append(trader_dat.MapTraderSchedule[traderId], schedule)
	session.Send(&debug_api.AddTraderSchedule_Out{})
}

//添加商人刷新时间
func (debuf DebugAPI) AddTraderRefreshTime(session *net.Session, in *debug_api.AddTraderRefreshTime_In) {
	switch in.TraderId {
	case 1:
		trader_dat.YingHaiJiShiRefresh = append(trader_dat.YingHaiJiShiRefresh, in.Timing)
	case 2:
		trader_dat.XunYouShangRenRefresh = append(trader_dat.XunYouShangRenRefresh, in.Timing)
	case 3:
		trader_dat.HeiShiLaoDaRefresh = append(trader_dat.HeiShiLaoDaRefresh, in.Timing)
	default:
		panic("undefine trader")
	}
	session.Send(&debug_api.AddTraderRefreshTime_Out{})
}

//清空商人刷新时间
func (debug DebugAPI) ClearTraderRefreshTime(session *net.Session, in *debug_api.ClearTraderRefreshTime_In) {
	switch in.TraderId {
	case 1:
		trader_dat.YingHaiJiShiRefresh = []int64{}
	case 2:
		trader_dat.XunYouShangRenRefresh = []int64{}
	case 3:
		trader_dat.HeiShiLaoDaRefresh = []int64{}
	default:
		panic("undefine trader")
	}
	session.Send(&debug_api.ClearTraderRefreshTime_Out{})
}

//开启城镇
func (debug DebugAPI) OpenTown(session *net.Session, in *debug_api.OpenTown_In) {
	state := module.State(session)
	town := town_dat.GetTownWithTownId(in.TownId)
	module.Town.SetTownLock(state, town.Lock)
	session.Send(&debug_api.OpenTown_Out{})
}

//增加测试全局邮件（互动）
func (debug DebugAPI) AddGlobalMail(session *net.Session, in *debug_api.AddGlobalMail_In) {
	state := module.State(session)
	now := time.Now().Unix()
	sendTime := now + in.SendDelay
	title := fmt.Sprintf("测试的全局邮件公告")
	content := fmt.Sprintf("在%v产生，%d秒后发送", time.Now(), in.SendDelay)
	attachments := []*mail_dat.Attachment{&mail_dat.Attachment{3, 0, 5}, &mail_dat.Attachment{0, 303, 5}, &mail_dat.Attachment{1, 0, 20000}, &mail_dat.Attachment{2, 0, 200}}
	var expireTime int64
	if in.ExpireDelay > 1 {
		expireTime = now + in.ExpireDelay
	} else {
		expireTime = in.ExpireDelay
	}
	module.MailRPC.AddGlobalMail(state.Database, title, content, attachments, sendTime, expireTime, 1, 0, 0, 0, 0)
}

//产生没有模版的公告
func (debug DebugAPI) CreateAnnouncementWithoutTpl(session *net.Session, in *debug_api.CreateAnnouncementWithoutTpl_In) {
	now := time.Now().Unix()
	rpc.RemoteGlobalAnnouncementCreate(0, "", "没有模版的公公告", now+5, now+100, 30)
}

////设置累计登录天数
//func (debug DebugAPI) SetLoginDay(session *net.Session, in *debug_api.SetLoginDay_In) {
//	state := module.State(session)
//	playerAwardInfo := state.Database.Lookup.PlayerLoginAwardRecord(state.PlayerId)
//	playerAwardInfo.ActiveDays = in.Days
//	state.Database.Update.PlayerLoginAwardRecord(playerAwardInfo)
//	session.Send(&debug_api.SetLoginDay_Out{})
//}

////重置七天登录奖励
//func (debug DebugAPI) ResetLoginAward(session *net.Session, in *debug_api.ResetLoginAward_In) {
//	state := module.State(session)
//	playerAwardInfo := state.Database.Lookup.PlayerLoginAwardRecord(state.PlayerId)
//	playerAwardInfo.Record = 0
//	state.Database.Update.PlayerLoginAwardRecord(playerAwardInfo)
//	session.Send(&debug_api.ResetLoginAward_Out{})
//}

//设置累计登录天数
func (debug DebugAPI) SetLoginDay(session *net.Session, in *debug_api.SetLoginDay_In) {
	state := module.State(session)
	eventRecord := state.EventsState.GetPlayerEventInfoById(event_dat.EVENT_SEVEN_DAY_AWARDS)
	eventRecord.LastUpdated -= 24 * 3600 * int64(in.Days)
	state.EventsState.List[event_dat.EVENT_SEVEN_DAY_AWARDS] = &module.EventInfo{
		EventId:     event_dat.EVENT_SEVEN_DAY_AWARDS,
		EndUnixTime: eventRecord.EndUnixTime,
		Awarded:     eventRecord.Awarded,
		MaxAward:    eventRecord.MaxAward,
		LastUpdated: eventRecord.LastUpdated,
	}
	//更新至数据库
	state.Database.Update.PlayerEventAwardRecord(&mdb.PlayerEventAwardRecord{
		Pid:             state.Database.PlayerId(),
		RecordBytes:     state.EventsState.Encode(),
		JsonEventRecord: state.JsonEventsState.Encode(),
	})
	session.Send(&debug_api.ResetLoginAward_Out{})
}

//重置七天登录奖励
func (debug DebugAPI) ResetLoginAward(session *net.Session, in *debug_api.ResetLoginAward_In) {
	state := module.State(session)
	state.EventsState.ClearState(state.Database, event_dat.EVENT_SEVEN_DAY_AWARDS)
	session.Send(&debug_api.ResetLoginAward_Out{})
}

//重置玩家得获奖lock
func (debug DebugAPI) RestPlayerAwardLock(session *net.Session, in *debug_api.RestPlayerAwardLock_In) {
	state := module.State(session)
	playerAwardLockMissionLevelInfo := state.Database.Lookup.PlayerMissionLevel(state.PlayerId)
	playerAwardLockHardLevelInfo := state.Database.Lookup.PlayerHardLevel(state.PlayerId)
	playerAwardLockHardLevelInfo.AwardLock = 0
	playerAwardLockMissionLevelInfo.AwardLock = 0
	state.Database.Update.PlayerMissionLevel(playerAwardLockMissionLevelInfo)
	state.Database.Update.PlayerHardLevel(playerAwardLockHardLevelInfo)
	session.Send(&debug_api.RestPlayerAwardLock_Out{})
}

func (debug DebugAPI) ResetRainbowLevel(session *net.Session, in *debug_api.ResetRainbowLevel_In) {
	state := module.State(session)
	rainbowLevel := state.Database.Lookup.PlayerRainbowLevel(state.PlayerId)
	rainbowLevel.Status = rainbow_dat.RAINBOW_LEVEL_STATUS_NEVER_PASS
	rainbowLevel.Order = rainbow_dat.INIT_LEVEL_ORDER
	rainbowLevel.ResetTimestamp = 0
	state.Database.Update.PlayerRainbowLevel(rainbowLevel)

	state.RainbowLevelState = module.NewRainbowLevelState()
	state.SaveRainbowLevelState()
}

func (debug DebugAPI) MonthCard(session *net.Session, in *debug_api.MonthCard_In) {
	// state := module.State(session)
	// db := state.Database
	// playerMonthCard := db.Lookup.PlayerMonthCardInfo(db.PlayerId())
	// start := core_time.GetDayFirstTime(time.Now(), 1)
	// if playerMonthCard == nil {
	// 	db.Insert.PlayerMonthCardInfo(&mdb.PlayerMonthCardInfo{
	// 		Pid:       db.PlayerId(),
	// 		Starttime: start,
	// 		Endtime:   start + 30*3600*24,
	// 		Buytimes:  1,
	// 	})
	// 	session.Send(&notify_api.NotifyMonthCardOpen_Out{})
	// } else if playerMonthCard.Endtime-start <= 3*3600*24 {
	// 	playerMonthCard.Buytimes++
	// 	if playerMonthCard.Starttime < start {
	// 		playerMonthCard.Starttime = start
	// 	}
	// 	playerMonthCard.Endtime += 30 * 3600 * 24
	// 	db.Update.PlayerMonthCardInfo(playerMonthCard)
	// 	session.Send(&notify_api.NotifyMonthCardRenewal_Out{})
	// }
	// session.Send(&debug_api.MonthCard_Out{})
	// 临时用来测试每日登陆活动，可随时删除
	state := module.State(session)
	eventId := int16(event_dat.EVENT_TOTAL_LOGIN)
	eventInfo, _ := event_dat.GetEventInfoById(eventId)
	eventRecord := state.EventsState.GetPlayerEventInfoById(eventId)
	state.EventsState.UpdateMax(state.Database, eventId, eventRecord.MaxAward+1)
	if eventRecord.MaxAward-eventRecord.Awarded > 1 {
		state.EventsState.List[eventId] = &module.EventInfo{
			EventId:     eventInfo.Id,
			EndUnixTime: eventInfo.End,
			Awarded:     0,
			MaxAward:    1,
			LastUpdated: core_time.GetNowTime(),
		}
		state.Database.Update.PlayerEventAwardRecord(&mdb.PlayerEventAwardRecord{
			Pid:             state.PlayerId,
			RecordBytes:     state.EventsState.Encode(),
			JsonEventRecord: state.JsonEventsState.Encode(),
		})
	}
}

func (debug DebugAPI) SetRainbowLevel(session *net.Session, in *debug_api.SetRainbowLevel_In) {
	state := module.State(session)
	rainbowLevel := state.Database.Lookup.PlayerRainbowLevel(state.PlayerId)
	fail.When(in.Segment < rainbow_dat.INIT_SEGMENT_NUM, "segment 参数不正确")
	fail.When(in.Order < rainbow_dat.INIT_LEVEL_ORDER || in.Order > rainbow_dat.LEVEL_NUM_PER_SEGMENT, "order 参数不正确")

	rainbowLevel.Segment = in.Segment
	rainbowLevel.Order = in.Order
	rainbowLevel.Status = rainbow_dat.RAINBOW_LEVEL_STATUS_NEVER_PASS
	state.Database.Update.PlayerRainbowLevel(rainbowLevel)

	state.RainbowLevelState = module.NewRainbowLevelState()
	state.SaveRainbowLevelState()
}

func (debug DebugAPI) SendPushNotification(session *net.Session, in *debug_api.SendPushNotification_In) {
	state := module.State(session)
	tencent.SendNotificationToSingleAccount(state.PlayerId, "信鸽推送标题", "信鸽推送内容\n信鸽推送内容")
}

func (debug DebugAPI) ResetPetVirtualEnv(session *net.Session, in *debug_api.ResetPetVirtualEnv_In) {
	state := module.State(session)
	playerPveState := state.Database.Lookup.PlayerPveState(state.PlayerId)
	playerPveState.DailyNum = 0
	state.Database.Update.PlayerPveState(playerPveState)
}

func (debug DebugAPI) AddFame(session *net.Session, in *debug_api.AddFame_In) {
	state := module.State(session)
	module.Player.AddFame(state.Database, in.Val)
}

func (debug DebugAPI) AddWorldChatMessage(session *net.Session, in *debug_api.AddWorldChatMessage_In) {
	state := module.State(session)
	now := time.Now().Unix()
	for x := int16(0); x < in.Num; x++ {
		module.AddWorldChannelMessage(&module.Message{
			Pid:     state.PlayerId,
			MsgType: module.CHANNEL_CHAT,
			//Nickname:  state.PlayerNick,
			Timestamp: now,
			Content:   []byte(fmt.Sprintf("%d", x)),
		})
	}
}

// func (debug DebugAPI) AddTotalLoginDays(session *net.Session, in *debug_api.AddTotalLoginDays_In) {
// 	state := module.State(session)
// 	eventId := int16(event_dat.EVENT_TOTAL_LOGIN)
// 	eventInfo, _ := event_dat.GetEventInfoById(eventId)
// 	eventRecord := state.EventsState.GetPlayerEventInfoById(eventId)
// 	state.EventsState.UpdateMax(state.Database, eventId, eventRecord.MaxAward+1)
// 	if eventRecord.MaxAward-eventRecord.Awarded > 1 {
// 		state.EventsState.List[eventId] = &module.EventInfo{
// 			EventId:     eventInfo.Id,
// 			EndUnixTime: eventInfo.End,
// 			Awarded:     0,
// 			MaxAward:    1,
// 			LastUpdated: core_time.GetNowTime(),
// 		}
// 		state.Database.Update.PlayerEventAwardRecord(&mdb.PlayerEventAwardRecord{
// 			Pid:         state.PlayerId,
// 			RecordBytes: state.EventsState.Encode(),
// 		})
// 	}
// }

// func (debug DebugAPI) ChangeFirstDailyRecharge(session *net.Session, in *debug_api.ChangeFirstDailyRecharge_In) {
// 	state := module.State(session)
// 	eventId := int16(event_dat.EVENT_FIRST_RECHARGE_DAILY)
// 	eventInfo, _ := event_dat.GetEventInfoById(eventId)
// 	eventInfo.Start -= 3600 * 24
// 	eventInfo.End -= 3600 * 24
// 	eventInfo.Dispose -= 3600 * 24
// 	state.EventsState.ClearState(state.Database, eventId)
// }

func (debug DebugAPI) EnterSandbox(session *net.Session, in *debug_api.EnterSandbox_In) {
	//进入沙盒模式
	state := module.State(session)
	state.SandBoxMode()
}

func (debug DebugAPI) SandboxRollback(session *net.Session, in *debug_api.SandboxRollback_In) {
	//在沙盒状态退出游戏就会回滚
}

func (debug DebugAPI) ExitSandbox(session *net.Session, in *debug_api.ExitSandbox_In) {
	state := module.State(session)
	state.ExitSandbox()
}

func (debug DebugAPI) ResetShadedMissions(session *net.Session, in *debug_api.ResetShadedMissions_In) {
	state := module.State(session)
	level_id := in.LevelId
	missionLevelRec := module.Mission.GetMissionLevelRecord(state.Database, level_id)
	if missionLevelRec != nil {
		missionLevelRec.EmptyShadowBits = 0
		state.Database.Update.PlayerMissionLevelRecord(missionLevelRec)
	}
}

func (debug DebugAPI) CleanCornucopia(session *net.Session, in *debug_api.CleanCornucopia_In) {
	state := module.State(session)
	cornucopiaInfo := state.Database.Lookup.PlayerCornucopia(state.PlayerId)
	cornucopiaInfo.OpenTime = 0
	cornucopiaInfo.DailyCount = 0
	state.Database.Update.PlayerCornucopia(cornucopiaInfo)
}

func (debug DebugAPI) AddTotem(session *net.Session, in *debug_api.AddTotem_In) {
	state := module.State(session)
	module.Totem.AddTotem(state.Database, in.TotemId)
}

func (debug DebugAPI) AddRune(session *net.Session, in *debug_api.AddRune_In) {
	state := module.State(session)
	playerTotemInfo := state.Database.Lookup.PlayerTotemInfo(state.PlayerId)
	playerTotemInfo.JadeRuneNum += in.JadeNum
	playerTotemInfo.RockRuneNum += in.RockNum
	state.Database.Update.PlayerTotemInfo(playerTotemInfo)
	session.Send(&notify_api.NotifyRuneChange_Out{
		RockRuneNum: playerTotemInfo.RockRuneNum,
		JadeRuneNum: playerTotemInfo.JadeRuneNum,
	})
}

func (debug DebugAPI) SendRareItemMessage(session *net.Session, in *debug_api.SendRareItemMessage_In) {
	state := module.State(session)
	rpc.RemoteAddWorldChannelTplMessage(state.PlayerId, []channel_dat.MessageTpl{
		channel_dat.MessageDrawSwordSoul{
			Player: channel_dat.ParamPlayer{state.PlayerNick, state.PlayerId},
			Item:   channel_dat.ParamItem{7, 1},
		},
	})
	rpc.RemoteAddWorldChannelTplMessage(state.PlayerId, []channel_dat.MessageTpl{
		channel_dat.MessageRainbowLevelGhost{
			Player: channel_dat.ParamPlayer{state.PlayerNick, state.PlayerId},
			Item:   channel_dat.ParamItem{6, 1},
			Level:  channel_dat.ParamString{"测试关卡"},
		},
	})
}

func (debug DebugAPI) AddSwordDrivingAction(session *net.Session, in *debug_api.AddSwordDrivingAction_In) {
	state := module.State(session)
	db := state.Database
	module.DrivingSword.IncActionPoint(db, in.Point)
}

func (debug DebugAPI) ResetDrivingSwordData(session *net.Session, in *debug_api.ResetDrivingSwordData_In) {
	state := module.State(session)
	module.DrivingSword.ResetCloudData(state, in.Cloud)
}

func (debug_api DebugAPI) AddSwordSoulFragment(session *net.Session, in *debug_api.AddSwordSoulFragment_In) {
	state := module.State(session)
	module.Player.IncSwordSoulFragment(state.Database, in.Number, player_dat.SWORDSOULFRAGMENT, -1, 0)
}

func (debug DebugAPI) ResetMoneyTreeStatus(session *net.Session, in *debug_api.ResetMoneyTreeStatus_In) {
	state := module.State(session)
	record := state.Database.Lookup.PlayerMoneyTree(state.PlayerId)
	if record != nil {
		record.LastWavedTime = 0
		record.Total = 0
		record.WavedTimes = 0
		record.WavedTimesTotal = 0
		state.Database.Update.PlayerMoneyTree(record)
	}
}

func (debug DebugAPI) ResetTodayMoneyTree(session *net.Session, in *debug_api.ResetTodayMoneyTree_In) {
	state := module.State(session)
	record := state.Database.Lookup.PlayerMoneyTree(state.PlayerId)
	if record != nil && record.LastWavedTime > 0 {
		record.LastWavedTime -= 3600 * 24
		state.Database.Update.PlayerMoneyTree(record)
	}
}

func (debug_api DebugAPI) CleanSwordSoulIngotDrawNums(session *net.Session, in *debug_api.CleanSwordSoulIngotDrawNums_In) {
	state := module.State(session)
	playerSwordSoulState := state.Database.Lookup.PlayerSwordSoulState(state.PlayerId)
	playerSwordSoulState.IngotNum = 0
	state.Database.Update.PlayerSwordSoulState(playerSwordSoulState)
}

func (debug_api DebugAPI) PunchDrivingSwordCloud(session *net.Session, in *debug_api.PunchDrivingSwordCloud_In) {
	state := module.State(session)
	db := state.Database
	i := 0
	for i < 100 {
		i++

		drivingInfo := state.Database.Lookup.PlayerDrivingSwordInfo(state.PlayerId)

		except_obstacle := driving_sword_dat.CountObstacleByCloud(drivingInfo.CurrentCloud)
		except_exploring := driving_sword_dat.CountExploringByCloud(drivingInfo.CurrentCloud)
		except_visiting := driving_sword_dat.CountVisitingByCloud(drivingInfo.CurrentCloud)
		except_treasure := driving_sword_dat.CountTreasureByCloud(drivingInfo.CurrentCloud)
		except_teleport := driving_sword_dat.CountTeleportByCloud(drivingInfo.CurrentCloud)
		except_hole := 1

		obstacle_num := 0
		exploring_num := 0
		visiting_num := 0
		treasure_num := 0
		teleport_num := 0
		hole_num := 0

		module.DrivingSword.ResetCloudData(state, drivingInfo.CurrentCloud)
		module.DrivingSword.PunchCloudData(state)
		db.Select.PlayerDrivingSwordEvent(func(row *mdb.PlayerDrivingSwordEventRow) {
			if row.CloudId() == drivingInfo.CurrentCloud {
				switch row.EventType() {
				case 0:
					hole_num++
				case 1:
					teleport_num++
				case 2:
					obstacle_num++
				case 3:
					teleport_num++
				}
			}
		})
		db.Select.PlayerDrivingSwordEventExploring(func(row *mdb.PlayerDrivingSwordEventExploringRow) {
			if row.CloudId() == drivingInfo.CurrentCloud {
				exploring_num++
			}
		})
		db.Select.PlayerDrivingSwordEventVisiting(func(row *mdb.PlayerDrivingSwordEventVisitingRow) {
			if row.CloudId() == drivingInfo.CurrentCloud {
				visiting_num++
			}
		})
		db.Select.PlayerDrivingSwordEventTreasure(func(row *mdb.PlayerDrivingSwordEventTreasureRow) {
			if row.CloudId() == drivingInfo.CurrentCloud {
				treasure_num++
			}
		})
		if !(except_obstacle == obstacle_num &&
			except_exploring == exploring_num &&
			except_visiting == visiting_num &&
			except_treasure == treasure_num &&
			except_teleport == teleport_num &&
			except_hole == hole_num) {
			fmt.Println("===========")
			fmt.Printf("except obstacle %d, real %d; except_exploring %d, real  %d, except_visiting %d, real %d, except_treasure %d, real %d; except_teleport %d, real %d; except_hole %d, real %d\n",
				except_obstacle, obstacle_num,
				except_exploring, exploring_num,
				except_visiting, visiting_num,
				except_treasure, treasure_num,
				except_teleport, teleport_num,
				except_hole, hole_num)
			fmt.Println("===========")
			break
		}

	}
}

func (debug DebugAPI) ClearCliqueDailyDonate(session *net.Session, in *debug_api.ClearCliqueDailyDonate_In) {
	state := module.State(session)
	playerCliqueInfo := state.Database.Lookup.PlayerGlobalCliqueInfo(state.PlayerId)
	if playerCliqueInfo.CliqueId > 0 {
		playerCliqueInfo.DonateCoinsTime = 1
		state.Database.Update.PlayerGlobalCliqueInfo(playerCliqueInfo)
	}
}

func (debug DebugAPI) SetCliqueContrib(session *net.Session, in *debug_api.SetCliqueContrib_In) {
	state := module.State(session)
	playerCliqueInfo := state.Database.Lookup.PlayerGlobalCliqueInfo(state.PlayerId)
	if playerCliqueInfo.CliqueId > 0 {
		playerCliqueInfo.Contrib = in.Contrib
		state.Database.Update.PlayerGlobalCliqueInfo(playerCliqueInfo)
	}
}

func (debug DebugAPI) RefreshCliqueWorship(session *net.Session, in *debug_api.RefreshCliqueWorship_In) {
	state := module.State(session)
	playerCliqueInfo := state.Database.Lookup.PlayerGlobalCliqueInfo(state.PlayerId)
	if playerCliqueInfo.CliqueId > 0 {
		playerCliqueBuilding := state.Database.Lookup.PlayerGlobalCliqueBuilding(state.PlayerId)
		playerCliqueBuilding.WorshipTime = 1
		state.Database.Update.PlayerGlobalCliqueBuilding(playerCliqueBuilding)
	}
}

func (debug DebugAPI) CliqueEscortRecoverBattleWin(session *net.Session, in *debug_api.CliqueEscortRecoverBattleWin_In) {
	state := module.State(session)
	module.CliqueEscortRPC.RecoverBattleWin(state.Database, state.PlayerId, in.BoatId)
}

func (debug DebugAPI) CliqueEscortHijackBattleWin(session *net.Session, in *debug_api.CliqueEscortHijackBattleWin_In) {
	state := module.State(session)
	module.CliqueEscortRPC.HijackBattleWin(state.Database, state.PlayerId, in.BoatId)
}

func (debug DebugAPI) CliqueEscortNotifyMessage(session *net.Session, in *debug_api.CliqueEscortNotifyMessage_In) {
	state := module.State(session)
	msg2 := channel_dat.MessageBoatHijackingFinished{
		Hijacker: channel_dat.ParamPlayer{
			Nick: state.PlayerNick,
			Pid:  state.PlayerId,
		},
		Coins:   channel_dat.ParamString{"1000"},
		Fame:    channel_dat.ParamString{"1000"},
		Contrib: channel_dat.ParamString{"1000"},
	}
	module.CliqueEscortRPC.NewBoatMessage(state.Database, msg2)
	module.CliqueEscortRPC.NewBoatMessage(state.Database, channel_dat.MessageBoatEscortFinished{})
	module.CliqueEscortRPC.NewBoatMessage(state.Database, channel_dat.MessageBoatHijackFinished{})
	module.CliqueEscortRPC.NewBoatMessage(state.Database, channel_dat.MessageBoatRecovered{
		Fighter: channel_dat.ParamPlayer{
			Nick: state.PlayerNick,
			Pid:  state.PlayerId,
		},
	})
	module.CliqueEscortRPC.NewBoatMessage(state.Database, channel_dat.MessageBoatRecoveredByHero{
		Fighter: channel_dat.ParamPlayer{
			Nick: state.PlayerNick,
			Pid:  state.PlayerId,
		},
	})
	module.CliqueEscortRPC.NewBoatMessage(state.Database, channel_dat.MessageBoatHijacking{
		Hijacker: channel_dat.ParamPlayer{
			Nick: state.PlayerNick,
			Pid:  state.PlayerId,
		},
		HijackerClique: channel_dat.ParamClique{
			Name: "随便帮派",
			Id:   10001,
		},
	})
}

func (debug DebugAPI) CliqueEscortNotifyDailyQuest(session *net.Session, in *debug_api.CliqueEscortNotifyDailyQuest_In) {
	state := module.State(session)
	var playerQuest *mdb.PlayerGlobalCliqueDailyQuest

	state.Database.Select.PlayerGlobalCliqueDailyQuest(func(row *mdb.PlayerGlobalCliqueDailyQuestRow) {
		playerQuest = row.GoObject()
		playerQuest.AwardStatus = clique_dat.CLIQUE_QUEST_STATUS_NO_AWARD
		playerQuest.FinishCount = 0
		state.Database.Update.PlayerGlobalCliqueDailyQuest(playerQuest)
	})

	return
}

func (debug DebugAPI) SetCliqueBuildingLevel(session *net.Session, in *debug_api.SetCliqueBuildingLevel_In) {
	fail.When(in.Level < 0 || in.Level > 12, "level error")
	state := module.State(session)
	playerCliqueInfo := state.Database.Lookup.PlayerGlobalCliqueInfo(state.PlayerId)
	if playerCliqueInfo.CliqueId > 0 {
		clique := state.Database.Lookup.GlobalClique(playerCliqueInfo.CliqueId)
		switch in.Building {
		case clique_dat.CLIQUE_BUILDING_ZONGCI:
			clique.TempleBuildingLevel = in.Level
		case clique_dat.CLIQUE_BUILDING_HUICHUNTANG:
			clique.HealthBuildingLevel = in.Level
		case clique_dat.CLIQUE_BUILDING_SHENBINGTANG:
			clique.AttackBuildingLevel = in.Level
		case clique_dat.CLIQUE_BUILDING_JINGANGTANG:
			clique.DefenseBuildingLevel = in.Level
		case clique_dat.CLIQUE_BUILDING_QIANZHUANG:
			clique.BankBuildingLevel = in.Level
		case clique_dat.CLIQUE_BUILDING_ZONGDUO:
			clique.CenterBuildingLevel = in.Level
		}
		state.Database.Update.GlobalClique(clique)
	}
}

func (debug DebugAPI) SetCliqueBuildingMoney(session *net.Session, in *debug_api.SetCliqueBuildingMoney_In) {

	fail.When(in.Money < 0 || in.Money >= 100000000, "money error")
	state := module.State(session)
	playerCliqueInfo := state.Database.Lookup.PlayerGlobalCliqueInfo(state.PlayerId)
	if playerCliqueInfo.CliqueId > 0 {
		clique := state.Database.Lookup.GlobalClique(playerCliqueInfo.CliqueId)
		switch in.Building {
		case clique_dat.CLIQUE_BUILDING_ZONGCI:
			clique.TempleBuildingCoins = in.Money
		case clique_dat.CLIQUE_BUILDING_HUICHUNTANG:
			clique.HealthBuildingCoins = in.Money
		case clique_dat.CLIQUE_BUILDING_SHENBINGTANG:
			clique.AttackBuildingCoins = in.Money
		case clique_dat.CLIQUE_BUILDING_JINGANGTANG:
			clique.DefenseBuildingCoins = in.Money
		case clique_dat.CLIQUE_BUILDING_QIANZHUANG:
			clique.BankBuildingCoins = in.Money
		case clique_dat.CLIQUE_BUILDING_ZONGDUO:
			clique.CenterBuildingCoins = in.Money
		}
		state.Database.Update.GlobalClique(clique)
	}
}
func (this DebugAPI) EscortBench(session *net.Session, in *debug_api.EscortBench_In) {
	state := module.State(session)
	state.Database.Select.GlobalClique(func(row *mdb.GlobalCliqueRow) {
		memberList := module.CliqueRPC.CliqueInfoListPid(row.Id())
		for _, pid := range memberList {
			state.Database.AgentExecute(pid, func(agentDB *mdb.Database) {
				module.CliqueEscortRPC.ForceEscort(agentDB)
			})
		}
	})
}

func (this DebugAPI) ResetCliqueEscortDailyNum(session *net.Session, in *debug_api.ResetCliqueEscortDailyNum_In) {
	state := module.State(session)
	escortInfo := state.Database.Lookup.PlayerGlobalCliqueEscort(state.PlayerId)
	if escortInfo != nil {
		escortInfo.DailyEscortNum = 0
		state.Database.Update.PlayerGlobalCliqueEscort(escortInfo)
	}
	hijackInfo := state.Database.Lookup.PlayerGlobalCliqueHijack(state.PlayerId)
	if hijackInfo != nil {
		hijackInfo.DailyHijackNum = 0
		state.Database.Update.PlayerGlobalCliqueHijack(hijackInfo)
	}
}

func (this DebugAPI) SetMissionStarMax(session *net.Session, in *debug_api.SetMissionStarMax_In) {
	state := module.State(session)

	playerId := state.PlayerId
	playerQuest := state.Database.Lookup.PlayerQuest(playerId)
	questNow := quest_dat.GetQuestById(playerQuest.QuestId)
	quests := quest_dat.GetQuestsByOrder(questNow.Order)

	for _, quest := range quests {
		missionLevelId := quest.MissionLevelId
		missionId := mission_dat.GetMissionIDByLevelID(missionLevelId)
		levelStar := mission_dat.CalLevelStar(missionLevelId)

		if levelStar != nil {
			var playerMissionLevelRecord *mdb.PlayerMissionLevelRecord
			state.Database.Select.PlayerMissionLevelRecord(func(row *mdb.PlayerMissionLevelRecordRow) {
				if row.MissionLevelId() == missionLevelId {
					playerMissionLevelRecord = row.GoObject()
					row.Break()
				}
			})

			if playerMissionLevelRecord != nil {
				playerMissionLevelRecord.Round = levelStar.ThreeStarRound
				state.Database.Update.PlayerMissionLevelRecord(playerMissionLevelRecord)
			} else {
				playerMissionLevelRecord = &mdb.PlayerMissionLevelRecord{}
				playerMissionLevelRecord.Pid = playerId
				playerMissionLevelRecord.MissionId = missionId
				playerMissionLevelRecord.MissionLevelId = missionLevelId
				playerMissionLevelRecord.Round = levelStar.ThreeStarRound
				playerMissionLevelRecord.OpenTime = time.Now().Unix()
				playerMissionLevelRecord.LastEnterTime = time.Now().Unix()
				state.Database.Insert.PlayerMissionLevelRecord(playerMissionLevelRecord)
			}

			// update quest
			oldStar := mission_dat.CalLevelStarByRound(playerMissionLevelRecord.MissionLevelId, int8(playerMissionLevelRecord.Round))
			newStar := mission_dat.CalLevelStarByRound(playerMissionLevelRecord.MissionLevelId, int8(levelStar.ThreeStarRound))
			if newStar > oldStar {
				module.Quest.RefreshQuestForMissionStarChange(state.Database, missionId, int16(newStar-oldStar))
			}
		}
	}
}

func (this DebugAPI) TakeAdditionQuest(session *net.Session, in *debug_api.TakeAdditionQuest_In) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("Request Failed\n%s\n", debug.Stack(1, "    "))
			session.Send(&debug_api.TakeAdditionQuest_Out{
				Msg: []byte(fmt.Sprintf("%v", err)),
			})
		}
	}()
	state := module.State(session)
	additionQuestDat := quest_dat.GetAdditionQuestById(in.FirstQuestId)
	state.Database.Select.PlayerAdditionQuest(func(row *mdb.PlayerAdditionQuestRow) {
		if row.SerialNumber() > 0 && row.SerialNumber() == additionQuestDat.SerialNumber {
			state.Database.Delete.PlayerAdditionQuest(row.GoObject())
			row.Break()
		}
		if row.SerialNumber() == 0 && row.QuestId() == in.FirstQuestId {
			state.Database.Delete.PlayerAdditionQuest(row.GoObject())
			row.Break()
		}
	})
	module.Quest.TakeAdditionQuest4Debug(session, in.FirstQuestId)
}
func (this DebugAPI) CliqueBankCd(session *net.Session, in *debug_api.CliqueBankCd_In) {
	state := module.State(session)
	playerCliqueBuilding := state.Database.Lookup.PlayerGlobalCliqueBuilding(state.PlayerId)
	if playerCliqueBuilding.GoldExchangeNum > 0 {
		playerCliqueBuilding.SilverExchangeTime = core_time.GetNowTime() - clique_dat.CLIQUE_BUILDING_BANK_SOLD_TIMESPAN + 10
	}
	if playerCliqueBuilding.SilverExchangeNum > 0 {
		playerCliqueBuilding.SilverExchangeTime = core_time.GetNowTime() - clique_dat.CLIQUE_BUILDING_BANK_SOLD_TIMESPAN + 10
	}
	state.Database.Update.PlayerGlobalCliqueBuilding(playerCliqueBuilding)
}

func (this DebugAPI) CleanRechargeInfo(session *net.Session, in *debug_api.CleanRechargeInfo_In) {
	state := module.State(session)
	playerMonthCard := state.Database.Lookup.PlayerMonthCard(state.PlayerId)
	if playerMonthCard != nil {
		playerMonthCard.ExpireTimestamp = 0
		playerMonthCard.AwardTimestamp = 0
		state.Database.Update.PlayerMonthCard(playerMonthCard)
	}
	playerInfo := state.Database.Lookup.PlayerInfo(state.PlayerId)
	playerInfo.Ingot = 0
	state.Database.Update.PlayerInfo(playerInfo)
	playerVipInfo := state.Database.Lookup.PlayerVip(state.PlayerId)
	playerVipInfo.Ingot = 0
	playerVipInfo.CardId = ""
	playerVipInfo.Level = 0
	state.Database.Update.PlayerVip(playerVipInfo)

}
