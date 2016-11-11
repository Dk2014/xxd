package mission

import (
	"core/fail"
	"core/log"
	"core/net"
	"core/time"
	"game_server/api/protocol/mission_api"
	"game_server/battle"
	"game_server/dat/battle_pet_dat"
	"game_server/dat/chest_dat"
	"game_server/dat/event_dat"
	"game_server/dat/item_dat"
	"game_server/dat/mission_dat"
	"game_server/dat/player_dat"
	"game_server/dat/quest_dat"
	"game_server/dat/role_dat"
	"game_server/dat/skill_dat"
	//	"game_server/dat/vip_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/tlog"
	"game_server/xdlog"
	"math/rand"
	goTime "time"
)

func doEnterLevel(session *net.Session, level_info *mission_dat.MissionLevel, levelType int8) {
	module.Town.LeaveTown(session)

	state := module.State(session)

	// 先清理之前的关卡数据
	module.Mission.LeaveMissionLevel(state)

	state.MissionLevelState = module.NewMissionLevelState(levelType, level_info.Id)

	state.MissionLevelState.AwardBox = level_info.AwardBox

	// 统计主地图怪物数
	for _, enemy_id := range mission_dat.GetEnemyIdByMissionLevelId(state.MissionLevelState.LevelId) {
		enemy := mission_dat.GetMissionLevelEnemyById(enemy_id)
		if enemy.ShadedMissionId == 0 {
			state.MissionLevelState.MainMonsterLeft++
		}
	}

	//  加载玩家所有角色，计算战斗数据加成并保存
	state.MissionLevelState.LoadFighterAttribute(state)

	//加载伙伴技能使用次数
	state.MissionLevelState.LoadBuddySkill(state)

	module.Quest.RefreshQuest(state, quest_dat.QUEST_TYPE_MISSION, xdlog.ET_QUEST)
}

func OpenLevelBox(state *module.SessionState, boxPos int8, xdEventType int32) (boxId int64) {
	// 检查是否杀光怪物
	fail.When(!state.MissionLevelState.HasKilledThoseMustDie(), "there are still some creatures left behind")

	// 检查开宝箱次数
	fail.When(state.MissionLevelState.BoxState.AwardOpenCount < 1, "can't open box")
	// 检查对应宝箱是否已经开过
	fail.When(state.MissionLevelState.BoxState.OpenedBoxPos[boxPos-1] > 0, "box have already opend")

	boxId = realOpenBox(state, state.MissionLevelState.LevelId, boxPos, &state.MissionLevelState.BoxState, nil, xdEventType)

	// 开宝箱结束后，奖励关卡
	if state.MissionLevelState.BoxState.AwardOpenCount == 0 {
		awardLevel(state, xdEventType)
		// 通关后清理关卡状态
		module.Mission.LeaveMissionLevel(state)
	}
	return
}

func awardCatchedBattlePet(state *module.SessionState, xdEventType int32) {
	fail.When(state.MissionLevelState.CatchedBattlePetId == 0, "not found pet")
	module.BattlePet.AddPet(state.Database, state.MissionLevelState.CatchedBattlePetId, tlog.IFR_MISSION_LEVEL_CATCH_PET, xdEventType)
	state.MissionLevelState.CatchedBattlePetId = 0
}

func awardLevel(state *module.SessionState, xdEventType int32) {
	levelInfo := mission_dat.GetMissionLevelById(state.MissionLevelState.LevelId)
	playerLevel := state.Database.Lookup.PlayerMissionLevel(state.PlayerId)

	realAwardLevel(state, playerLevel, levelInfo,
		state.MissionLevelState.BoxState.AwardExp,
		state.MissionLevelState.BoxState.AwardCoin,
		state.MissionLevelState.BoxState.AwardMultExp,
		state.MissionLevelState.BoxState.AwardMultCoin, xdEventType)

	// 重置关卡宝箱奖励加成数据
	state.MissionLevelState.BoxState.AwardExp = 0
	state.MissionLevelState.BoxState.AwardCoin = 0
	state.MissionLevelState.BoxState.AwardMultCoin = 0
	state.MissionLevelState.BoxState.AwardMultExp = 0

	switch state.MissionLevelState.LevelType {
	case battle.BT_MISSION_LEVEL:
		// 更新区域关卡记录
		levelRecord := module.Mission.GetMissionLevelRecord(state.Database, levelInfo.Id)
		if state.MissionLevelState.BossScore > levelRecord.Score {
			levelRecord.Score = state.MissionLevelState.BossScore
		}
		if levelRecord.Round == 0 || levelRecord.Round > int8(state.MissionLevelState.TotalRound) {
			oldStar := mission_dat.CalLevelStarByRound(levelRecord.MissionLevelId, int8(levelRecord.Round))
			newStar := mission_dat.CalLevelStarByRound(levelRecord.MissionLevelId, int8(state.MissionLevelState.TotalRound))
			if newStar > oldStar {
				module.Quest.RefreshQuestForMissionStarChange(state.Database, levelInfo.MissionId, int16(newStar-oldStar))
			}
			levelRecord.Round = int8(state.MissionLevelState.TotalRound)
		}

		// 累计关卡进入次数
		if levelInfo.DailyNum > 0 && !state.MissionLevelState.HaveReduceLevelDialyNum {
			levelRecord.DailyNum += 1
		}
		//第一次通关更新award_lock字段
		if levelInfo.Lock > playerLevel.AwardLock {
			playerLevel.AwardLock = levelInfo.Lock
			state.Database.Update.PlayerMissionLevel(playerLevel)
		}
		state.Database.Update.PlayerMissionLevelRecord(levelRecord)
		// 扣除剩余的体力
		if levelInfo.Physical > mission_dat.LEVEL_PHYSICAL_MIN {
			module.Physical.Decrease(state.Database, int16(levelInfo.Physical)-mission_dat.LEVEL_PHYSICAL_MIN, tlog.PFR_MISSION_LEVEL)
		}

		if session, ok := module.Player.GetPlayerOnline(state.PlayerId); ok {
			playerMission := state.Database.Lookup.PlayerMission(state.PlayerId)

			module.Notify.SendPlayerKeyChanged(session, 0, playerMission.MaxOrder)
			module.Notify.SendMissionLevelLockChanged(session, playerLevel.MaxLock, playerLevel.AwardLock)
		}

	case battle.BT_RESOURCE_LEVEL:
		// 资源关卡要成功通关后才累计进入次数
		extendLevel := mission_dat.GetExtendLevelRequireLevel(state.MissionLevelState.LevelType, state.MissionLevelState.LevelId)

		playerExtendLevel := state.Database.Lookup.PlayerExtendLevel(state.PlayerId)
		if levelInfo.SubType == mission_dat.RESOURCE_COIN_LEVEL {
			playerExtendLevel.CoinPassTime = time.GetNowTime()
			playerExtendLevel.CoinDailyNum += 1
			if playerExtendLevel.CoinsMaxlevel < extendLevel {
				playerExtendLevel.CoinsMaxlevel = extendLevel
			}

			module.Quest.RefreshDailyQuest(state, quest_dat.DAILY_QUEST_CLASS_RESOURCE_COIN_LEVEL)

		} else {
			playerExtendLevel.ExpPassTime = time.GetNowTime()
			playerExtendLevel.ExpDailyNum += 1
			if playerExtendLevel.ExpMaxlevel < extendLevel {
				playerExtendLevel.ExpMaxlevel = extendLevel
			}

			module.Quest.RefreshDailyQuest(state, quest_dat.DAILY_QUEST_CLASS_RESOURCE_EXP_LEVEL)
		}

		state.Database.Update.PlayerExtendLevel(playerExtendLevel)

		// 扣除剩余的体力
		if levelInfo.Physical > mission_dat.LEVEL_PHYSICAL_MIN {
			module.Physical.Decrease(state.Database, int16(levelInfo.Physical)-mission_dat.LEVEL_PHYSICAL_MIN, tlog.PFR_EXTEND_LEVEL)
		}

	case battle.BT_BUDDY_LEVEL:
		playerExtendLevel := state.Database.Lookup.PlayerExtendLevel(state.PlayerId)
		playerExtendLevel.BuddyPassTime = time.GetNowTime()
		playerExtendLevel.BuddyDailyNum += 1
		playerExtendLevel.RolePos = 0
		playerExtendLevel.BuddyPos = 0
		playerExtendLevel.RandBuddyRoleId = randRoleIdForBuddyLevel(state.Database)
		state.Database.Update.PlayerExtendLevel(playerExtendLevel)
		module.Quest.RefreshDailyQuest(state, quest_dat.DAILY_QUEST_CLASS_BUDDY_LEVEL)

		// 扣除剩余的体力
		if levelInfo.Physical > mission_dat.LEVEL_PHYSICAL_MIN {
			module.Physical.Decrease(state.Database, int16(levelInfo.Physical)-mission_dat.LEVEL_PHYSICAL_MIN, tlog.PFR_EXTEND_LEVEL)
		}

	case battle.BT_PET_LEVEL:
		playerExtendLevel := state.Database.Lookup.PlayerExtendLevel(state.PlayerId)
		playerExtendLevel.PetPassTime = time.GetNowTime()
		playerExtendLevel.PetDailyNum += 1
		state.Database.Update.PlayerExtendLevel(playerExtendLevel)
		module.Quest.RefreshDailyQuest(state, quest_dat.DAILY_QUEST_CLASS_PET_LEVEL)

		// 扣除剩余的体力
		if levelInfo.Physical > mission_dat.LEVEL_PHYSICAL_MIN {
			module.Physical.Decrease(state.Database, int16(levelInfo.Physical)-mission_dat.LEVEL_PHYSICAL_MIN, tlog.PFR_EXTEND_LEVEL)
		}

	case battle.BT_GHOST_LEVEL:
		playerExtendLevel := state.Database.Lookup.PlayerExtendLevel(state.PlayerId)
		playerExtendLevel.GhostPassTime = time.GetNowTime()
		playerExtendLevel.GhostDailyNum += 1
		state.Database.Update.PlayerExtendLevel(playerExtendLevel)
		module.Quest.RefreshDailyQuest(state, quest_dat.DAILY_QUEST_CLASS_GHOST_LEVEL)

		// 扣除剩余的体力
		if levelInfo.Physical > mission_dat.LEVEL_PHYSICAL_MIN {
			module.Physical.Decrease(state.Database, int16(levelInfo.Physical)-mission_dat.LEVEL_PHYSICAL_MIN, tlog.PFR_EXTEND_LEVEL)
		}

	case battle.BT_HARD_LEVEL:
		//难度关卡的数值
		hardLevelDat := mission_dat.GetHardLevelInfo(levelInfo.ParentId)
		//玩家难度关卡的记录
		playerHardLevelRecord := module.Mission.GetHardLevelRecordById(state.Database, levelInfo.Id)
		//如果有今日次数限制则更新
		if levelInfo.DailyNum > 0 {
			playerHardLevelRecord.DailyNum += 1
		}
		//更新最后进入时间
		playerHardLevelRecord.LastEnterTime = time.GetNowTime()
		//记录历史最高分
		if playerHardLevelRecord.Round == 0 || playerHardLevelRecord.Round > int8(state.MissionLevelState.TotalRound) {
			playerHardLevelRecord.Round = int8(state.MissionLevelState.TotalRound)
		}
		state.Database.Update.PlayerHardLevelRecord(playerHardLevelRecord)

		var notifyLockChange int32

		//获得新的难度关卡功能权值
		playerHardLevel := state.Database.Lookup.PlayerHardLevel(state.PlayerId)
		if hardLevelDat.AwardHardLevelLock > playerHardLevel.Lock {
			playerHardLevel.Lock = hardLevelDat.AwardHardLevelLock
			state.Database.Update.PlayerHardLevel(playerHardLevel)

			notifyLockChange = hardLevelDat.AwardHardLevelLock
		}
		//第一次通关判断 更新award_lock
		if hardLevelDat.HardLevelLock > playerHardLevel.AwardLock {
			playerHardLevel.AwardLock = hardLevelDat.HardLevelLock
			state.Database.Update.PlayerHardLevel(playerHardLevel)
		}

		module.Quest.RefreshDailyQuest(state, quest_dat.DAILY_QUEST_CLASS_HARD_LEVEL)

		// 扣除剩余的体力
		if levelInfo.Physical > mission_dat.LEVEL_HARD_PHYSICAL_MIN {
			module.Physical.Decrease(state.Database, int16(levelInfo.Physical)-mission_dat.LEVEL_HARD_PHYSICAL_MIN, tlog.PFR_HARD_LEVEL)
		}
		if session, ok := module.Player.GetPlayerOnline(state.PlayerId); ok {

			if notifyLockChange > 0 {
				module.Notify.SendHardLevelLockChange(session, notifyLockChange)
			}
			playerMission := state.Database.Lookup.PlayerMission(state.PlayerId)
			module.Notify.SendPlayerKeyChanged(session, 0, playerMission.MaxOrder)
		}
	case battle.BT_FATE_BOX_LEVEL:
		fateBoxDat := chest_dat.GetFateBoxByMissionLevelId(state.MissionLevelState.LevelId)
		fateBoxState := state.Database.Lookup.PlayerFateBoxState(state.PlayerId)
		if fateBoxState.Lock < fateBoxDat.AwardLock {
			fateBoxState.Lock = fateBoxDat.AwardLock
			state.Database.Update.PlayerFateBoxState(fateBoxState)
		}

	}
	tlog.PlayerMissionFlowLog(state.Database, state.MissionLevelState.LevelId, tlog.FINISH)
	xdlog.MissionLog(state.Database, state.MissionLevelState.LevelId, xdlog.MA_FINISH)
}

func realOpenBox(state *module.SessionState, levelId int32, boxPos int8, boxState *module.MissionLevelBoxState, delayAwardItem map[int16]int16, xdEventType int32) (boxId int64) {
	if boxState.AwardOpenCount == 0 {
		return
	}

	/*
		开启规则：
		品质数字越小，品质档次越高
		如果随机到一个宝箱，但该局游戏已经获得过该宝箱时，则取次一档的宝箱作为当次奖励
		如果没有可获取的次档宝箱，则取高一档可获取的宝箱。

	*/
	levelBoxs := mission_dat.GetLevelBoxByLevelId(levelId)

	var order int8
	if state.MissionLevelState != nil {
		//levelType = state.MissionLevelState.LevelType
		isFirst := false
		if state.MissionLevelState.LevelType == mission_dat.LEVEL_TYPE_HARD {
			playerHardLevel := state.Database.Lookup.PlayerHardLevel(state.PlayerId)
			missionLevelInfo := mission_dat.GetMissionLevelById(levelId)
			hardMissionInfo := mission_dat.GetHardLevelInfo(missionLevelInfo.ParentId)
			if hardMissionInfo.AwardHardLevelLock > playerHardLevel.AwardLock {
				isFirst = true
			}
		}

		if state.MissionLevelState.LevelType == mission_dat.LEVEL_TYPE_MISSION {
			playerMissionLevel := state.Database.Lookup.PlayerMissionLevel(state.PlayerId)
			missionLevelInfo := mission_dat.GetMissionLevelById(levelId)
			if missionLevelInfo.Lock > playerMissionLevel.AwardLock {
				isFirst = true
			}

		}
		if state.MissionLevelState.LevelType == mission_dat.EXTEND_LEVEL_TYPE_RESOURCE {
			levelInfo := mission_dat.GetMissionLevelById(levelId)
			playerExtendLevel := state.Database.Lookup.PlayerExtendLevel(state.PlayerId)
			if levelInfo.SubType == mission_dat.RESOURCE_COIN_LEVEL {
				if playerExtendLevel.CoinPassTime <= 0 {
					isFirst = true
				}
			} else {
				if playerExtendLevel.ExpPassTime <= 0 {
					isFirst = true
				}
			}
		}

		if isFirst {
			//难度关卡判断第一次通关的固定奖励
			musts := mission_dat.GetMustLevelBoxByLevelId(state.MissionLevelState.LevelId)
			if len(musts) > 0 {
				awarded := state.MissionLevelState.BoxState.AwardedList
				for _, must := range musts {
					if _, ok := awarded[must.Order]; !ok {
						order = must.Order
						break
					}
				}
			}
		}

	}
	var chance, randChance int
	if order == 0 { //order != 0 是因为首次奖励已经计算好了必中物品
		for _, box := range levelBoxs {
			if _, ok := boxState.AwardedList[box.Order]; ok {
				chance += int(box.AwardChance)
				continue
			}
		}

		randChance = rand.Intn(100-chance) + 1

		// 检查随机的概率是否在某个品质的概率中，存在就抽中宝箱了
		for _, box := range levelBoxs {
			if _, ok := boxState.AwardedList[box.Order]; ok { //跳过重复的奖励
				continue
			}

			if randChance <= int(box.AwardChance) {
				order = box.Order
				break
			}

			randChance -= int(box.AwardChance)
		}
	}
	// 开启宝箱
	box := levelBoxs[order-1]

	boxId = box.Id

	boxState.AwardedList[order] = boxId
	boxState.AwardedItemType = box.AwardType
	boxState.OpenedBoxPos[boxPos-1] = order

	switch box.AwardType {
	case mission_dat.LEVEL_BOX_AWARD_EQUIMENT:
		if delayAwardItem != nil {
			delayAwardItem[int16(box.ItemId)] += int16(box.AwardNum)
		} else {
			module.Item.AddItem(state.Database, int16(box.ItemId), int16(box.AwardNum), tlog.IFR_MISSION_LEVEL_REAL_OPEN_BOX, xdEventType, "")
		}
	case mission_dat.LEVEL_BOX_AWARD_ITEM:
		if delayAwardItem != nil {
			delayAwardItem[int16(box.ItemId)] += int16(box.AwardNum)
		} else {
			module.Item.AddItem(state.Database, int16(box.ItemId), int16(box.AwardNum), tlog.IFR_MISSION_LEVEL_REAL_OPEN_BOX, xdEventType, "")
		}
	case mission_dat.LEVEL_BOX_AWARD_PET_BALL:
		//TODO 灵宠契约球废弃 直接奖励灵宠
		//pet := battle_pet_dat.GetBattlePetWithItemId(int32(box.ItemId))
		//module.BattlePet.AddPet(state.Database, int16(pet.PetId), int16(box.AwardNum))
	case mission_dat.LEVEL_BOX_AWARD_EXP:
		boxState.AwardExp += box.AwardNum
	case mission_dat.LEVEL_BOX_AWARD_COIN:
		boxState.AwardCoin += int64(box.AwardNum)
	case mission_dat.LEVEL_BOX_AWARD_MULTIPLE_COIN:
		boxState.AwardMultCoin += int64(box.AwardNum)
	case mission_dat.LEVEL_BOX_AWARD_MULTIPLE_EXP:
		boxState.AwardMultExp += box.AwardNum
	}

	boxState.AwardOpenCount -= 1

	return
}

func realAwardLevel(state *module.SessionState, playerLevel *mdb.PlayerMissionLevel,
	levelInfo *mission_dat.MissionLevel, extraExp int32, extraCoin int64, multExp int32, multCoin int64, xdEventType int32) {

	// 奖励钥匙(改用物品道具)
	if levelInfo.AwardKey > 0 {
		module.Item.AddItem(state.Database, item_dat.ITEM_MISSION_KEY_ID, int16(levelInfo.AwardKey), tlog.IFR_MISSION_LEVEL_AWARD, xdEventType, "")
	}

	// 奖励关卡权值
	if levelInfo.AwardLock > playerLevel.MaxLock {
		module.Mission.SetMissionLevelLock(state, levelInfo.AwardLock, playerLevel, false)
	}

	// 奖励友情
	if levelInfo.AwardRelationship > 0 {
		module.Team.IncRelationship(state.Database, levelInfo.AwardRelationship)
	}

	if levelInfo.AwardItemNum > 0 {
		module.Item.AddItem(state.Database, levelInfo.AwardItem, levelInfo.AwardItemNum, tlog.IFR_MISSION_LEVEL_AWARD, xdEventType, "")
	}

	awardExp := levelInfo.AwardExp + extraExp
	awardCoin := levelInfo.AwardCoin + extraCoin
	if multCoin > 0 {
		awardCoin *= multCoin
	}

	if multExp > 0 {
		awardExp *= multExp
	}

	//eventsInfo := event_dat.GetEventsInfo()
	multiypEventInfo, _ := event_dat.GetEventInfoById(event_dat.EVENT_MULTIPY_CONFIG)
	qqvipEventInfo, _ := event_dat.GetEventInfoById(event_dat.EVENT_QQVIP_ADDITION)
	isDuringEventMultipy := event_dat.CheckEventTime(multiypEventInfo, event_dat.NOT_END)
	isDuringEventQQVIP := event_dat.CheckEventTime(qqvipEventInfo, event_dat.NOT_END)

	// 奖励铜钱
	if awardCoin > 0 {
		var times float32 = 1
		if levelInfo.ParentType == battle.BT_MISSION_LEVEL {
			if isDuringEventMultipy {
				times *= module.GetMulitipyByKey(event_dat.MISSION_COINS) //普通关卡 翻倍奖励活动
			}
			if isDuringEventQQVIP {
				if (state.TencentState.QqVipStatus & 2) > 0 {
					times *= module.GetMulitipyByKey(event_dat.SUPERQQ_COINS) //qq超级会员加成
				} else if (state.TencentState.QqVipStatus & 1) > 0 {
					times *= module.GetMulitipyByKey(event_dat.QQVIP_COINS) //qq会员加成
				}
			}
		}
		module.Player.IncMoney(state.Database, state.MoneyState, int64(times*float32(awardCoin)), player_dat.COINS, tlog.MFR_MISSION_LEVEL_AWARD, xdEventType, "")
	}

	// 上阵角色奖励经验（包括主角）
	if awardExp > 0 {
		var times float32 = 1
		if levelInfo.ParentType == battle.BT_MISSION_LEVEL {
			if isDuringEventMultipy {
				times *= module.GetMulitipyByKey(event_dat.MISSION_COINS) //普通关卡 翻倍奖励活动
			}
			if isDuringEventQQVIP {
				if (state.TencentState.QqVipStatus & 2) > 0 {
					times *= module.GetMulitipyByKey(event_dat.SUPERQQ_EXPS) //qq超级会员加成
				} else if (state.TencentState.QqVipStatus & 1) > 0 {
					times *= module.GetMulitipyByKey(event_dat.QQVIP_EXPS) //qq会员加成
				}
			}
		}
		module.Role.AddFormRoleExp(state, int64(times*float32(awardExp)), tlog.EFT_MISSION_LEVEL_AWARD)
	}
}

func useItem(session *net.Session, roleId int, itemId, xdEventType int32) {
	state := module.State(session)
	module.Item.DelItemByItemId(state.Database, int16(itemId), 1, tlog.IFR_MISSION_LEVEL_USE_ITEM, xdEventType)

	attr := state.MissionLevelState.AttackerInfo.Fighters[roleId]
	buffs := state.MissionLevelState.AttackerInfo.Buffs[roleId]

	// 如果是使用buff道具，就用调用getFighter()接口获得战斗对象
	getFighter := func() *battle.Fighter {
		// 构造获得一个没有数据加成的fighter对象
		fighters := make([]*battle.Fighter, module.ALL_FIGHTER_POS_NUM)
		inFormRoleInfos := []*module.InFormRoleInfo{&module.InFormRoleInfo{RoleId: int8(roleId)}}
		module.GetBattleBiz.SetFighters(state.Database, inFormRoleInfos, fighters, false, false, module.FIGHT_FOR_EMPTY)
		return fighters[0]
	}

	buffs = battle.UseLevelItem(getFighter, itemId, &attr, buffs)
	if buffs != nil {
		state.MissionLevelState.AttackerInfo.Buffs[roleId] = buffs
	}

	state.MissionLevelState.AttackerInfo.Fighters[roleId] = attr

	module.Notify.RoleBattleStatusChange(session)
}

func checkOpenExtendLevel(levelType int8, subType int8) bool {
	var openDay []int
	var expectHour int

	switch levelType {
	case mission_dat.EXTEND_LEVEL_TYPE_RESOURCE:
		// 资源关卡区别铜钱和经验类的关卡，分别有不一样的开放时间
		if subType == mission_dat.RESOURCE_COIN_LEVEL {
			openDay = mission_dat.ExtendLevelCoinOpenDay
			expectHour = player_dat.RESET_COIN_LEVEL_TIMES_IN_HOUR

		} else if subType == mission_dat.RESOURCE_EXP_LEVEL {
			openDay = mission_dat.ExtendLevelExpOpenDay
			expectHour = player_dat.RESET_EXP_LEVEL_TIMES_IN_HOUR
		}

	case mission_dat.EXTEND_LEVEL_TYPE_BUDDY:
		openDay = mission_dat.ExtendLevelBuddyOpenDay
		expectHour = player_dat.RESET_BUDDY_LEVEL_TIMES_IN_HOUR

	case mission_dat.EXTEND_LEVEL_TYPE_PET:
		openDay = mission_dat.ExtendLevelPetOpenDay
		expectHour = player_dat.RESET_PET_LEVEL_TIMES_IN_HOUR

	case mission_dat.EXTEND_LEVEL_TYPE_GHOST:
		openDay = mission_dat.ExtendLevelGhostOpenDay
		expectHour = player_dat.RESET_GHOST_LEVEL_TIMES_IN_HOUR
	}

	weekDay := time.GetNowWeekByExpectHour(expectHour)

	for _, day := range openDay {
		if weekDay == goTime.Weekday(day) {
			return true
		}
	}

	return false
}

func getExtendLevelFuncId(levelType int8) (id int32) {
	switch levelType {
	case mission_dat.EXTEND_LEVEL_TYPE_RESOURCE:
		id = player_dat.FUNC_RESOURCE_LEVEL
	case mission_dat.EXTEND_LEVEL_TYPE_BUDDY:
		id = player_dat.FUNC_ACTIVE_LEVLE
	case mission_dat.EXTEND_LEVEL_TYPE_PET:
		id = player_dat.FUNC_ACTIVE_LEVLE
	case mission_dat.EXTEND_LEVEL_TYPE_GHOST:
		id = player_dat.FUNC_ACTIVE_LEVLE
	default:
		fail.When(true, "error extend level type")
	}
	return
}

func randRoleIdForBuddyLevel(db *mdb.Database) int8 {
	var buddyRoleId []int8
	db.Select.PlayerRole(func(row *mdb.PlayerRoleRow) {
		if !role_dat.IsMainRole(row.RoleId()) && row.Status() == role_dat.ROLE_STATUS_NOMAL {
			buddyRoleId = append(buddyRoleId, row.RoleId())
		}
	})

	c := len(buddyRoleId)
	fail.When(c == 0, "randRoleIdForBuddyLevel. no buddy in team")

	idx := rand.Intn(c)
	return buddyRoleId[idx]
}

func autoConfigPetForPetLevel(db *mdb.Database) bool {
	var eqpPetNum int //已装备的灵宠数量
	var emptyGrids []*mdb.PlayerBattlePetGrid
	eqpPetSet := make(map[int32]int8)

	db.Select.PlayerBattlePetGrid(func(row *mdb.PlayerBattlePetGridRow) {
		if row.BattlePetId() > 0 {
			eqpPetSet[row.BattlePetId()] = 1
			eqpPetNum++
		} else {
			emptyGrids = append(emptyGrids, row.GoObject())
		}
		if eqpPetNum >= battle_pet_dat.MIN_PET_NUM_IN_PET_LEVEL {
			row.Break()
		}
	})
	if eqpPetNum >= battle_pet_dat.MIN_PET_NUM_IN_PET_LEVEL {
		return true
	}

	var gridIndex int = 0
	db.Select.PlayerBattlePet(func(row *mdb.PlayerBattlePetRow) {
		if _, eqp := eqpPetSet[row.BattlePetId()]; !eqp {
			emptyGrids[gridIndex].BattlePetId = row.BattlePetId()
			db.Update.PlayerBattlePetGrid(emptyGrids[gridIndex])
			gridIndex++
			eqpPetNum++
			eqpPetSet[row.BattlePetId()] = 1
		}
		if eqpPetNum >= battle_pet_dat.MIN_PET_NUM_IN_PET_LEVEL {
			row.Break()
		}
	})
	return eqpPetNum >= battle_pet_dat.MIN_PET_NUM_IN_PET_LEVEL
}

func checkGhostLevelConfig(db *mdb.Database, mainRoleId int8) (ret bool) {
	db.Select.PlayerGhostEquipment(func(row *mdb.PlayerGhostEquipmentRow) {
		if row.RoleId() == mainRoleId {
			if row.Pos1() > 0 {
				ret = true
			} else {
				ret = false
			}
			row.Break()
		}
	})
	return
}

func enterExtendLevel(session *net.Session, levelType int8, levelId int32) bool {
	state := module.State(session)

	funcId := getExtendLevelFuncId(levelType)
	if funcId != player_dat.FUNC_RESOURCE_LEVEL {
		module.Player.MustOpenFunc(state.Database, funcId)
	} else {
		fail.When(!module.Player.IsOpenFunc(state.Database, player_dat.FUNC_RESOURCE_LEVEL), "未开启资源关卡")
	}

	levelInfo := mission_dat.GetMissionLevelById(levelId)
	fail.When(!checkOpenExtendLevel(levelType, levelInfo.SubType), "extend-level is not opened")

	roleLevel := module.Role.GetMainRole(state.Database).Level
	fail.When(roleLevel < mission_dat.GetExtendLevelRequireLevel(levelType, levelId), "require role level")

	playerExtendLevel := state.Database.Lookup.PlayerExtendLevel(state.PlayerId)

	var dailyNum int8
	var maxDailyNum int8

	switch levelType {
	case mission_dat.EXTEND_LEVEL_TYPE_RESOURCE:
		if levelInfo.SubType == mission_dat.RESOURCE_COIN_LEVEL {
			dailyNum = playerExtendLevel.CoinDailyNum
			maxDailyNum = mission_dat.EXTEND_LEVEL_RESOURCE_MAX_DAILY_NUM + int8(playerExtendLevel.CoinsBuyNum)
		} else {
			dailyNum = playerExtendLevel.ExpDailyNum
			maxDailyNum = mission_dat.EXTEND_LEVEL_RESOURCE_MAX_DAILY_NUM + int8(playerExtendLevel.ExpBuyNum)
		}
		tlog.PlayerSystemModuleFlowLog(state.Database, tlog.SMT_MISSION_RESOURCE)

	case mission_dat.EXTEND_LEVEL_TYPE_BUDDY:
		dailyNum = playerExtendLevel.BuddyDailyNum
		maxDailyNum = mission_dat.EXTEND_LEVEL_BUDDY_MAX_DAILY_NUM
		tlog.PlayerSystemModuleFlowLog(state.Database, tlog.SMT_MISSION_RESOURCE2)

	case mission_dat.EXTEND_LEVEL_TYPE_PET:
		dailyNum = playerExtendLevel.PetDailyNum
		maxDailyNum = mission_dat.EXTEND_LEVEL_PET_MAX_DAILY_NUM
		tlog.PlayerSystemModuleFlowLog(state.Database, tlog.SMT_MISSION_RESOURCE2)

	case mission_dat.EXTEND_LEVEL_TYPE_GHOST:
		dailyNum = playerExtendLevel.GhostDailyNum
		maxDailyNum = mission_dat.EXTEND_LEVEL_GHOST_MAX_DAILY_NUM
		tlog.PlayerSystemModuleFlowLog(state.Database, tlog.SMT_MISSION_RESOURCE2)
	}

	// 检查已进入过的资源关卡今日次数是否已满
	fail.When(dailyNum >= maxDailyNum, "daily num full")

	// 检查体力
	if levelInfo.Physical > 0 {
		fail.When(!module.Physical.CheckGE(state, int16(levelInfo.Physical)), "physical not enough when enter resource level")
	}

	switch levelType {
	case mission_dat.EXTEND_LEVEL_TYPE_PET:
		if !autoConfigPetForPetLevel(state.Database) {
			return false
		}
	case mission_dat.EXTEND_LEVEL_TYPE_GHOST:
		if !checkGhostLevelConfig(state.Database, state.RoleId) {
			return false
		}
	}

	// 扣除体力
	if levelInfo.Physical > 0 {
		// 先至少扣除3点体力,通关后再扣除剩余体力
		module.Physical.Decrease(state.Database, mission_dat.LEVEL_PHYSICAL_MIN, tlog.PFR_EXTEND_LEVEL)
	}

	doEnterLevel(session, levelInfo, levelType)

	tlog.PlayerMissionFlowLog(state.Database, levelId, tlog.ENTER)
	xdlog.MissionLog(state.Database, levelId, xdlog.MA_ENTER)

	return true
}

func getExtendLevelInfo(state *module.SessionState, levelType mission_api.ExtendType, rsp *mission_api.GetExtendLevelInfo_Out) {
	playerExtendLevel := state.Database.Lookup.PlayerExtendLevel(state.PlayerId)
	fail.When(playerExtendLevel == nil, "not opened extend level")
	var needUpdate bool
	if levelType == mission_api.EXTEND_TYPE_RESOURCE {

		if playerExtendLevel.CoinPassTime >= 0 && !time.IsInPointHour(player_dat.RESET_COIN_LEVEL_TIMES_IN_HOUR, playerExtendLevel.CoinPassTime) {
			needUpdate = true
			playerExtendLevel.CoinDailyNum = 0
		}

		if playerExtendLevel.ExpPassTime >= 0 && !time.IsInPointHour(player_dat.RESET_EXP_LEVEL_TIMES_IN_HOUR, playerExtendLevel.ExpPassTime) {
			needUpdate = true
			playerExtendLevel.ExpDailyNum = 0
		}

		if playerExtendLevel.CoinsBuyTime >= 0 && !time.IsInPointHour(player_dat.RESET_BUY_RESOURCE_LEVEL_TIMES_IN_HOUR, playerExtendLevel.CoinsBuyTime) {
			needUpdate = true
			playerExtendLevel.CoinsBuyNum = 0
		}

		if playerExtendLevel.ExpBuyTime >= 0 && !time.IsInPointHour(player_dat.RESET_BUY_RESOURCE_LEVEL_TIMES_IN_HOUR, playerExtendLevel.ExpBuyTime) {
			needUpdate = true
			playerExtendLevel.ExpBuyNum = 0
		}

		rsp.Info = make([]mission_api.GetExtendLevelInfo_Out_Info, 2)
		rsp.Info[0] = mission_api.GetExtendLevelInfo_Out_Info{
			LevelType:    mission_api.EXTEND_LEVEL_TYPE_RESOURCE,
			LevelSubType: mission_api.EXTEND_LEVEL_SUB_TYPE_COIN,
			DailyNum:     playerExtendLevel.CoinDailyNum,
			MaxLevel:     playerExtendLevel.CoinsMaxlevel,
			BuyNum:       playerExtendLevel.CoinsBuyNum,
		}

		rsp.Info[1] = mission_api.GetExtendLevelInfo_Out_Info{
			LevelType:    mission_api.EXTEND_LEVEL_TYPE_RESOURCE,
			LevelSubType: mission_api.EXTEND_LEVEL_SUB_TYPE_EXP,
			DailyNum:     playerExtendLevel.ExpDailyNum,
			MaxLevel:     playerExtendLevel.ExpMaxlevel,
			BuyNum:       playerExtendLevel.ExpBuyNum,
		}

	} else if levelType == mission_api.EXTEND_TYPE_ACTIVITY {
		if playerExtendLevel.BuddyPassTime > 0 && !time.IsInPointHour(player_dat.RESET_BUDDY_LEVEL_TIMES_IN_HOUR, playerExtendLevel.BuddyPassTime) {
			needUpdate = true
			playerExtendLevel.BuddyDailyNum = 0
			playerExtendLevel.RolePos = 0
			playerExtendLevel.BuddyPos = 0
			// 隔天重置随机的伙伴角色
			playerExtendLevel.RandBuddyRoleId = randRoleIdForBuddyLevel(state.Database)
		}

		if playerExtendLevel.PetPassTime >= 0 && !time.IsInPointHour(player_dat.RESET_PET_LEVEL_TIMES_IN_HOUR, playerExtendLevel.PetPassTime) {
			needUpdate = true
			playerExtendLevel.PetDailyNum = 0
		}

		if playerExtendLevel.GhostPassTime >= 0 && !time.IsInPointHour(player_dat.RESET_GHOST_LEVEL_TIMES_IN_HOUR, playerExtendLevel.GhostPassTime) {
			needUpdate = true
			playerExtendLevel.GhostDailyNum = 0
		}

		rsp.Info = make([]mission_api.GetExtendLevelInfo_Out_Info, 3)
		rsp.Info[0] = mission_api.GetExtendLevelInfo_Out_Info{
			LevelType:    mission_api.EXTEND_LEVEL_TYPE_BUDDY,
			LevelSubType: mission_api.EXTEND_LEVEL_SUB_TYPE_NONE,
			DailyNum:     playerExtendLevel.BuddyDailyNum,
		}

		rsp.Info[1] = mission_api.GetExtendLevelInfo_Out_Info{
			LevelType:    mission_api.EXTEND_LEVEL_TYPE_PET,
			LevelSubType: mission_api.EXTEND_LEVEL_SUB_TYPE_NONE,
			DailyNum:     playerExtendLevel.PetDailyNum,
		}

		rsp.Info[2] = mission_api.GetExtendLevelInfo_Out_Info{
			LevelType:    mission_api.EXTEND_LEVEL_TYPE_GHOST,
			LevelSubType: mission_api.EXTEND_LEVEL_SUB_TYPE_NONE,
			DailyNum:     playerExtendLevel.GhostDailyNum,
		}
	} else {
		fail.When(true, "getExtendLevelInfo. error level_type")
	}

	if needUpdate {
		state.Database.Update.PlayerExtendLevel(playerExtendLevel)
	}
}

func openSmallBox(state *module.SessionState, boxId int32, rsp *mission_api.OpenSmallBox_Out, xdEventType int32) {
	fail.When(len(state.MissionLevelState.SmallBoxList) == 0, "can't open small box")
	found := false

	var awardCount int8
	var foundBoxId int32

	for foundBoxId, awardCount = range state.MissionLevelState.SmallBoxList {
		if foundBoxId == boxId {
			found = true
			delete(state.MissionLevelState.SmallBoxList, foundBoxId)
			break
		}
	}
	fail.When(!found, "not found small box")

	items := mission_dat.GetSmailBoxItems(foundBoxId)
	for i := int8(0); i < awardCount; i++ {
		randNum := int8(rand.Intn(100) + 1)
		chance := int8(0)
		for _, item := range items {
			if randNum > chance && randNum <= chance+item.Probability {
				switch item.AwardType {
				case mission_dat.LEVEL_BOX_AWARD_COIN:
					module.Player.IncMoney(state.Database, state.MoneyState, int64(item.ItemNumber), player_dat.COINS, tlog.MFR_MISSION_LEVEL_SMALL_BOX, xdEventType, "")
				case mission_dat.LEVEL_BOX_AWARD_ITEM:
					module.Item.AddItem(state.Database, int16(item.ItemId), int16(item.ItemNumber), tlog.IFR_MISSION_LEVEL_SMALL_BOX, xdEventType, "")
				}
				rsp.Items = append(rsp.Items, mission_api.OpenSmallBox_Out_Items{
					BoxItemId: item.BoxItemId,
				})
			}
			chance += item.Probability
		}
	}

	fail.When(len(rsp.Items) == 0, "[openSmallBox] can't award item")
	return
}

func openMengYao(state *module.SessionState, mengYaoId int32, rsp *mission_api.OpenMengYao_Out) bool {
	if len(state.MissionLevelState.MengYaoList) == 0 {
		log.Error("don't have any mengyao in this level ")
		return false
	}

	found := false

	var effect int8
	var foundMengYaoId int32

	for foundMengYaoId, effect = range state.MissionLevelState.MengYaoList {
		if foundMengYaoId == mengYaoId {
			found = true
			delete(state.MissionLevelState.MengYaoList, foundMengYaoId)
			break
		}
	}
	fail.When(!found, "can't find the mengyao")
	switch effect {
	case mission_dat.RECOVER_BATTLE_PET:
		state.MissionLevelState.CalledBattlePet = make(map[int32]int8)
	case mission_dat.RECOVER_MEMBERS_HELATH:
		for _, roleId := range module.Role.GetFormRoleId(state) {
			fighterAttr := state.MissionLevelState.AttackerInfo.Fighters[int(roleId)]
			fighterAttr.Health = fighterAttr.MaxHealth
			state.MissionLevelState.AttackerInfo.Fighters[int(roleId)] = fighterAttr
		}
	case mission_dat.RECOVER_SKILL_RELEASE:
		for skillId, release := range state.MissionLevelState.SkillReleaseNum {
			skillContent := skill_dat.GetSkillContent(int16(skillId))
			if release < int(skillContent.ReleaseNum) {
				state.MissionLevelState.SkillReleaseNum[skillId] = int(skillContent.ReleaseNum)
			}
		}
	default:
		fail.When(true, "can't understand the effect type")
	}
	rsp.MengYaoId = mengYaoId
	return true
}

func doOpenShadedBox(state *module.SessionState, shadedId, xdEventType int32) {
	if creatureCount, exist := state.MissionLevelState.ShadowList[shadedId]; !exist || creatureCount > 0 {
		fail.When(true, "You have no right to open this shaded box, it doesn't exist or already opened or still have alive guarders")
	}
	delete(state.MissionLevelState.ShadowList, shadedId)
	shadow, _ := mission_dat.GetShadedMission(shadedId)
	missionLevelRec := module.Mission.GetMissionLevelRecord(state.Database, shadow.MissionLevelId)
	missionLevelRec.EmptyShadowBits |= 1 << uint(shadow.Order-1)
	state.Database.Update.PlayerMissionLevelRecord(missionLevelRec)

	shadedTreasure, _ := mission_dat.GetShadedMission(shadedId)
	if shadedTreasure.AwardExp > 0 {
		module.Role.AddFormRoleExp(state, int64(shadedTreasure.AwardExp), tlog.EFT_SHADED_AWARD)
	}
	if shadedTreasure.AwardCoin > 0 {
		module.Player.IncMoney(state.Database, state.MoneyState, int64(shadedTreasure.AwardCoin), player_dat.COINS, tlog.MFR_SHADED_AWARD, xdEventType, "")
	}
	if shadedTreasure.AwardItem1 > 0 && shadedTreasure.AwardItem1Num > 0 {
		module.Item.AddItem(state.Database, shadedTreasure.AwardItem1, int16(shadedTreasure.AwardItem1Num), tlog.IFR_SHADED_ITEM_GET, xdEventType, "")
	}
	if shadedTreasure.AwardItem2 > 0 && shadedTreasure.AwardItem2Num > 0 {
		module.Item.AddItem(state.Database, shadedTreasure.AwardItem2, int16(shadedTreasure.AwardItem2Num), tlog.IFR_SHADED_ITEM_GET, xdEventType, "")
	}
	if shadedTreasure.AwardItem3 > 0 && shadedTreasure.AwardItem3Num > 0 {
		module.Item.AddItem(state.Database, shadedTreasure.AwardItem3, int16(shadedTreasure.AwardItem3Num), tlog.IFR_SHADED_ITEM_GET, xdEventType, "")
	}
}

func openRandomlBox(state *module.SessionState, levelId int32, randomBox *module.MissionLevelRandomBoxState, xdEventType int32) (boxId int64) {
	// 是否清理掉了需要清理的怪物
	if !state.MissionLevelState.HasKilledThoseMustDie() {
		return
	}

	//levelType := state.MissionLevelState.LevelType
	if state.MissionLevelState.LevelType == mission_dat.LEVEL_TYPE_HARD {
		playerHardLevel := state.Database.Lookup.PlayerHardLevel(state.PlayerId)
		missionLevelInfo := mission_dat.GetMissionLevelById(levelId)
		hardMissionInfo := mission_dat.GetHardLevelInfo(missionLevelInfo.ParentId)
		if hardMissionInfo.HardLevelLock >= playerHardLevel.AwardLock {
			state.MissionLevelState.RandomBox.FirstOpen = true
		}
	}

	boxId = getRandomAwardBox(state, levelId, randomBox, xdEventType)

	return
}

func getRandomAwardBox(state *module.SessionState, levelId int32, randomBox *module.MissionLevelRandomBoxState, xdEventType int32) (boxId int64) {
	// 已经开过了
	if randomBox.RandomBoxOpenCount < 1 {
		return
	}

	var randChance int
	levelBoxs := mission_dat.GetRandomBoxByLevelId(levelId)
	randChance = rand.Intn(100) + 1
	var order int8

	if randomBox.FirstOpen {
		//难度关卡判断第一次通关的固定奖励
		musts := mission_dat.GetMustRandomBoxByLevelId(state.MissionLevelState.LevelId)
		if len(musts) > 0 {
			order = musts[0].Order
		}
		randomBox.FirstOpen = false
	}

	if order == 0 {
		for _, box := range levelBoxs {
			if randChance <= int(box.AwardChance) {
				order = box.Order
				break
			}
			randChance -= int(box.AwardChance)
		}
	}

	// 开启宝箱
	var index int8
	if order > 0 {
		index = order - 1
	}

	boxId = levelBoxs[index].Id

	// 随机宝箱的奖励   0 '铜钱', 1 *'道具', 2 '装备'
	if levelBoxs[index].AwardType == 0 {
		module.Player.IncMoney(state.Database, state.MoneyState, int64(levelBoxs[index].AwardNum), player_dat.COINS, tlog.MFR_MISSION_LEVEL_AWARD, xdEventType, "")
	} else {
		module.Item.AddItem(state.Database, int16(levelBoxs[index].ItemId), int16(levelBoxs[index].AwardNum), tlog.MFR_MISSION_LEVEL_AWARD, xdEventType, "")
	}

	randomBox.RandomBoxOpenCount -= 1

	return
}

func awardAdditionalQuestItem(session *net.Session, rsp *mission_api.EvaluateLevel_Out, xdEventType int32) {
	rsp.AdditionalQuestItems = []mission_api.EvaluateLevel_Out_AdditionalQuestItems{}
	state := module.State(session)

	// 检查是否杀光怪物
	fail.When(!state.MissionLevelState.HasKilledThoseMustDie(), "there are still some creatures left behind")

	// 概率奖励支线任务物品
	var items map[int16]int16
	if state.MissionLevelState.HasBeenEvaluated {
		items = state.MissionLevelState.ItemsAwardedByAddQuest
	} else {
		items := module.Quest.RefreshQuestForPassMissionLevel(state.Database, state.MissionLevelState.LevelId, xdEventType)
		state.MissionLevelState.ItemsAwardedByAddQuest = items
	}
	for item_id, item_cnt := range items {
		rsp.AdditionalQuestItems = append(rsp.AdditionalQuestItems, mission_api.EvaluateLevel_Out_AdditionalQuestItems{
			ItemId:  item_id,
			ItemCnt: item_cnt,
		})
	}
}
