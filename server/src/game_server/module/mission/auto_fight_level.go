package mission

import (
	"core/fail"
	"core/net"
	"core/time"
	"game_server/api/protocol/mission_api"
	"game_server/dat/battle_pet_dat"
	"game_server/dat/mission_dat"
	"game_server/dat/player_dat"
	"game_server/dat/quest_dat"
	"game_server/dat/vip_dat"
	"game_server/module"
	"game_server/tlog"
	"game_server/xdlog"
	"math/rand"
)

//关卡扫荡
func autoFightLevel(session *net.Session, levelType int8, levelId int32, times int8, xdEventType int32) {
	out := new(mission_api.AutoFightLevel_Out)
	state := module.State(session)

	module.Player.MustOpenFunc(state.Database, player_dat.FUNC_DIRECTLY_AWARD)

	levelInfo := mission_dat.GetMissionLevelById(levelId)

	fail.When(levelInfo.ParentType != levelType, "关卡类型不一致")

	var levelRound int8 = 0

	//获取关卡星级和已进入次数
	switch levelInfo.ParentType {
	case mission_dat.LEVEL_TYPE_MISSION:
		//检查关卡是否已经开启
		playerLevel := state.Database.Lookup.PlayerMissionLevel(state.PlayerId)
		fail.When(playerLevel.AwardLock < levelInfo.Lock, "未开启或未通关，不可扫荡")

		playerMissionLevelRecord := module.Mission.GetMissionLevelRecord(state.Database, levelId)
		if !time.IsInPointHour(player_dat.RESET_MAIN_LEVEL_TIMES_IN_HOUR, playerMissionLevelRecord.LastEnterTime) {
			playerMissionLevelRecord.DailyNum = 0
		}
		playerMissionLevelRecord.LastEnterTime = time.GetNowTime()
		playerMissionLevelRecord.DailyNum += times

		//判断是不是boss关卡
		var isBossLevel bool = false
		mission_enemys := mission_dat.GetEnemyIdByMissionLevelId(playerMissionLevelRecord.MissionLevelId)
		for _, enemy_id := range mission_enemys {
			missionEnemy := mission_dat.GetMissionLevelEnemyById(int32(enemy_id))
			if missionEnemy.IsBoss == true {
				isBossLevel = true
				break
			}
		}
		if isBossLevel {
			if levelInfo.DailyNum > 0 {
				fail.When(playerMissionLevelRecord.DailyNum > levelInfo.DailyNum+int8(playerMissionLevelRecord.BuyTimes), "超过每日进入上限")
			}
		} else {
			if levelInfo.DailyNum > 0 {
				fail.When(playerMissionLevelRecord.DailyNum > levelInfo.DailyNum, "超过每日进入上限")
			}
		}

		levelRound = playerMissionLevelRecord.Round
		state.Database.Update.PlayerMissionLevelRecord(playerMissionLevelRecord)
	case mission_dat.LEVEL_TYPE_HARD:
		playerHardLevelRecord := module.Mission.GetHardLevelRecordById(state.Database, levelId)
		fail.When(playerHardLevelRecord == nil, "hard level not open")
		if !time.IsInPointHour(player_dat.RESET_HARD_LEVEL_TIMES_IN_HOUR, playerHardLevelRecord.LastEnterTime) {
			playerHardLevelRecord.DailyNum = 0
		}
		if !time.IsInPointHour(player_dat.RESET_BUY_HARD_LEVEL_TIMES_IN_HOUR, playerHardLevelRecord.BuyUpdateTime) {
			playerHardLevelRecord.BuyTimes = 0
		}
		playerHardLevelRecord.LastEnterTime = time.GetNowTime()
		playerHardLevelRecord.DailyNum += times
		levelRound = playerHardLevelRecord.Round
		fail.When(playerHardLevelRecord.DailyNum-int8(playerHardLevelRecord.BuyTimes) > levelInfo.DailyNum, "超过每日进入上限")
		state.Database.Update.PlayerHardLevelRecord(playerHardLevelRecord)

	case mission_dat.EXTEND_LEVEL_TYPE_RESOURCE:

		module.VIP.CheckPrivilege(state, vip_dat.ZIYUANGUANQIASAODANG)
		funcId := getExtendLevelFuncId(levelType)
		if funcId != player_dat.FUNC_RESOURCE_LEVEL {
			module.Player.MustOpenFunc(state.Database, funcId)
		} else {
			fail.When(!module.Player.IsOpenFunc(state.Database, player_dat.FUNC_RESOURCE_LEVEL), "未开启资源关卡")
		}

		levelInfo := mission_dat.GetMissionLevelById(levelId)
		fail.When(!checkOpenExtendLevel(levelType, levelInfo.SubType), "extend-level is not opened")

		roleLevel := module.Role.GetMainRole(state.Database).Level
		extendLevel := mission_dat.GetExtendLevelRequireLevel(levelType, levelId)
		fail.When(roleLevel < extendLevel, "require role level")

		playerExtendLevel := state.Database.Lookup.PlayerExtendLevel(state.PlayerId)
		fail.When(playerExtendLevel == nil, "not opened resource level")

		var dailyNum, maxDailyNum int8
		if levelInfo.SubType == mission_dat.RESOURCE_COIN_LEVEL {
			dailyNum = playerExtendLevel.CoinDailyNum

			playerExtendLevel.CoinPassTime = time.GetNowTime()
			playerExtendLevel.CoinDailyNum += 1
			if playerExtendLevel.CoinsMaxlevel < extendLevel {
				playerExtendLevel.CoinsMaxlevel = extendLevel
			}
			maxDailyNum = mission_dat.EXTEND_LEVEL_RESOURCE_MAX_DAILY_NUM + int8(playerExtendLevel.CoinsBuyNum)

		} else {
			dailyNum = playerExtendLevel.ExpDailyNum

			playerExtendLevel.ExpPassTime = time.GetNowTime()
			playerExtendLevel.ExpDailyNum += 1
			if playerExtendLevel.ExpMaxlevel < extendLevel {
				playerExtendLevel.ExpMaxlevel = extendLevel
			}
			maxDailyNum = mission_dat.EXTEND_LEVEL_RESOURCE_MAX_DAILY_NUM + int8(playerExtendLevel.ExpBuyNum)
		}
		// 检查已进入过的资源关卡今日次数是否已满
		fail.When(dailyNum >= maxDailyNum, "daily num full")

		state.Database.Update.PlayerExtendLevel(playerExtendLevel)
	default:
		panic("未定义的扫荡关卡类型")
	}

	if levelInfo.ParentType != mission_dat.EXTEND_LEVEL_TYPE_RESOURCE {
		//检查星级是否满足
		levelStar := mission_dat.CalLevelStarByRound(levelInfo.Id, levelRound)
		fail.When(levelStar != mission_dat.THREE_STAR, "关卡评价不满足")
	}
	//扣体力
	module.Physical.Decrease(state.Database, int16(times*levelInfo.Physical), tlog.PFR_MISSION_LEVEL)

	//扫荡可以完成消灭怪物类任务
	enemyGroups := mission_dat.GetEnemyIdByMissionLevelId(levelInfo.Id)
	for _, enemyGroup := range enemyGroups {
		for i := int8(0); i < times; i++ {
			module.Quest.RefreshQuestForBeatEnemyGroup(state.Database, enemyGroup)
		}
	}

	awardItems := make(map[int16]int16) //关卡奖励

	for i := int8(0); i < times; i++ {
		var result mission_api.AutoFightLevel_Out_Result
		autoFightOpenLevelBox(session, levelInfo, &result, awardItems, xdEventType)

		autoFightOpenRandomlBox(state, levelInfo, &result, xdEventType) //随机宝箱

		tlog.PlayerMissionFlowLog(state.Database, levelId, tlog.AUTO)
		xdlog.MissionLog(state.Database, levelId, xdlog.MA_AUTO)

		// 刷新任务
		var classType int16
		if levelInfo.ParentType == mission_dat.LEVEL_TYPE_HARD {
			classType = quest_dat.DAILY_QUEST_CLASS_HARD_LEVEL
		} else if levelInfo.ParentType == mission_dat.LEVEL_TYPE_MISSION && levelInfo.Type == mission_dat.BOSS {
			classType = quest_dat.DAILY_QUEST_CLASS_MISSION_LEVEL_BOSS
		} else if levelInfo.ParentType == mission_dat.EXTEND_LEVEL_TYPE_RESOURCE {
			if levelInfo.SubType == mission_dat.RESOURCE_COIN_LEVEL {
				classType = quest_dat.DAILY_QUEST_CLASS_RESOURCE_COIN_LEVEL
			} else {
				classType = quest_dat.DAILY_QUEST_CLASS_RESOURCE_EXP_LEVEL
			}
		}
		if classType > 0 {
			module.Quest.RefreshDailyQuest(state, classType)
		}

		//支线奖励
		for itemId, num := range module.Quest.RefreshQuestForPassMissionLevel(state.Database, levelInfo.Id, xdEventType) {
			result.AdditionQuestItem = append(result.AdditionQuestItem, mission_api.AutoFightLevel_Out_Result_AdditionQuestItem{
				ItemId:  itemId,
				ItemNum: num,
			})
		}

		out.Result = append(out.Result, result)
	}
	if len(awardItems) > 0 {
		module.Item.BatchAddItem(state.Database, awardItems, tlog.IFR_MISSION_LEVEL_REAL_OPEN_BOX, xdEventType)
	}

	session.Send(out)
}

func autoFightOpenSmallBox(state *module.SessionState, levelInfo *mission_dat.MissionLevel, out_result *mission_api.AutoFightLevel_Out_Result, xdEventType int32) {
	missionLevelState := new(module.MissionLevelState)
	initSmallBoxAndSetMissionState(levelInfo.Id, missionLevelState)

	for foundBoxId, awardCount := range missionLevelState.SmallBoxList {
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
					out_result.SmallBox = append(out_result.SmallBox, mission_api.AutoFightLevel_Out_Result_SmallBox{item.BoxItemId})
				}
				chance += item.Probability
			}
		}
	}
}

func autoFightOpenLevelBox(session *net.Session, levelInfo *mission_dat.MissionLevel, out_result *mission_api.AutoFightLevel_Out_Result, awardItems map[int16]int16, xdEventType int32) {
	state := module.State(session)
	playerLevel := state.Database.Lookup.PlayerMissionLevel(state.PlayerId)
	boxState := &module.MissionLevelBoxState{}
	if levelInfo.AwardBox {
		boxState.AwardedList = make(map[int8]int64)
		boxState.OpenedBoxPos = make([]int8, 5)
		if levelInfo.ParentType == mission_dat.LEVEL_TYPE_MISSION {
			boxState.AwardOpenCount = mission_dat.MAX_LEVEL_BOX_AWARD_COUNT //主线关卡给3个奖励
		} else {
			boxState.AwardOpenCount = mission_dat.LEVEL_BOX_AWARD_COUNT
		}
		for i := int8(0); i < mission_dat.LEVEL_BOX_AWARD_COUNT; i++ {
			realOpenBox(state, levelInfo.Id, i+1, boxState, awardItems, xdEventType)
		}
	}
	realAwardLevel(state, playerLevel, levelInfo, boxState.AwardExp,
		boxState.AwardCoin, boxState.AwardMultExp, boxState.AwardMultCoin, xdEventType)

	for _, id := range boxState.AwardedList {
		out_result.LevelBox = append(out_result.LevelBox, mission_api.AutoFightLevel_Out_Result_LevelBox{id})
	}
	playerMission := state.Database.Lookup.PlayerMission(state.PlayerId)
	module.Notify.SendPlayerKeyChanged(session, 0, playerMission.MaxOrder)
}

func autoFightCatchBattlePet(state *module.SessionState, levelInfo *mission_dat.MissionLevel, out_battle *[]mission_api.AutoFightLevel_Out_Result_BattlePet, ballConsume *int16, ballNum int16, xdEventType int32) {
	enemyIds := mission_dat.GetEnemyIdByMissionLevelId(levelInfo.Id)
	for _, enemyId := range enemyIds {
		if levelPet, exist := battle_pet_dat.GetLevelBattlePet(enemyId); exist {
			if int8(rand.Int31n(100)+1) >= levelPet.Rate {
				//灵宠不出现
				continue
			}
			var petCatchEvent mission_api.AutoFightLevel_Out_Result_BattlePet
			petCatchEvent.PetId = levelPet.Pet.PetId
			i := int8(0)
			for *ballConsume < ballNum && i < levelPet.LiveRound {
				i++
				(*ballConsume)++
				petCatchEvent.ConsumeBalls++
				if rand.Int31n(100)+1 <= mission_dat.ITEM_PET_BALL_RATE {
					petCatchEvent.Catched = true
					module.BattlePet.AddPet(state.Database, petCatchEvent.PetId, tlog.IFR_MISSION_LEVEL_CATCH_PET, xdEventType)
					break
				}
			}
			*out_battle = append(*out_battle, petCatchEvent)
		}
	}
}

func autoFightOpenRandomlBox(state *module.SessionState, levelInfo *mission_dat.MissionLevel, out_result *mission_api.AutoFightLevel_Out_Result, xdEventType int32) {
	randomBoxState := &module.MissionLevelRandomBoxState{}

	if levelInfo.ParentType == mission_dat.LEVEL_TYPE_HARD {
		randomBoxState.RandomBoxOpenCount = mission_dat.RANDDOM_BOX_AWARD_COUNT
	}

	if id := getRandomAwardBox(state, levelInfo.Id, randomBoxState, xdEventType); id > 0 {
		out_result.RandomAwardBox = append(out_result.RandomAwardBox, mission_api.AutoFightLevel_Out_Result_RandomAwardBox{id})
	}
}
