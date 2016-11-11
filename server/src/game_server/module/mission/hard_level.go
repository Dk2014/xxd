package mission

import (
	"core/fail"
	"core/net"
	"core/time"
	"game_server/api/protocol/mission_api"
	"game_server/battle"
	"game_server/dat/buy_boss_level_times_config_dat"
	"game_server/dat/buy_hard_level_times_config_dat"
	"game_server/dat/mission_dat"
	"game_server/dat/player_dat"
	"game_server/dat/town_dat"
	"game_server/dat/vip_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/tlog"
	"game_server/xdlog"
)

func getHardLevel(state *module.SessionState, rsp *mission_api.GetHardLevel_Out) {
	tlog.PlayerSystemModuleFlowLog(state.Database, tlog.SMT_QUEST)
	var remain_buy_times int16
	state.Database.Select.PlayerHardLevelRecord(func(row *mdb.PlayerHardLevelRecordRow) {

		dailyNum := row.DailyNum()
		buyTimes := row.BuyTimes()
		if !time.IsInPointHour(player_dat.RESET_HARD_LEVEL_TIMES_IN_HOUR, row.LastEnterTime()) {
			dailyNum = 0
		}
		if !time.IsInPointHour(player_dat.RESET_HARD_LEVEL_TIMES_IN_HOUR, row.BuyUpdateTime()) {
			buyTimes = 0
		}

		maxBuyNum := module.VIP.PrivilegeTimes(state, vip_dat.GOUMAISHENYUAN)
		if buyTimes >= maxBuyNum {
			remain_buy_times = 0
		} else {
			remain_buy_times = int16(maxBuyNum - buyTimes)
		}

		rsp.Levels = append(rsp.Levels, mission_api.GetHardLevel_Out_Levels{
			LevelId:        row.LevelId(),
			DailyNum:       dailyNum,
			RoundNum:       row.Round(),
			BuyTimes:       buyTimes,
			RemainBuyTimes: remain_buy_times,
		})
	})
}

func buyHardLevelTimes(session *net.Session, hardLevelId int32, rsp *mission_api.BuyHardLevelTimes_Out) {
	state := module.State(session)
	//	hardLevelRecordId := mission_dat.GetSpecialLevelId(mission_dat.LEVEL_TYPE_HARD, int16(hardLevelId))
	state.Database.Select.PlayerHardLevelRecord(func(row *mdb.PlayerHardLevelRecordRow) {
		if row.LevelId() == hardLevelId {
			hardLevelRecord := row.GoObject()
			if !time.IsInPointHour(player_dat.RESET_BUY_HARD_LEVEL_TIMES_IN_HOUR, hardLevelRecord.BuyUpdateTime) {
				hardLevelRecord.BuyTimes = 0
			}
			maxBuyNum := module.VIP.PrivilegeTimes(state, vip_dat.GOUMAISHENYUAN)
			fail.When(hardLevelRecord.BuyTimes >= maxBuyNum, "购买深渊关卡次数已达到今日上限！")
			// 符合购买条件
			buyHardLevelCost := int64(buy_hard_level_times_config_dat.GetCost(int64(hardLevelRecord.BuyTimes + 1)))
			fail.When(!module.Player.CheckMoney(state, buyHardLevelCost, player_dat.INGOT), "元宝不足！")
			module.Player.DecMoney(state.Database, state.MoneyState, buyHardLevelCost, player_dat.INGOT, tlog.MFR_BUY_HARD_LEVEL_TIMES, xdlog.ET_BUY_HARD_LEVEL_TIMES)
			hardLevelRecord.BuyTimes += 1
			hardLevelRecord.BuyUpdateTime = time.GetNowTime()
			state.Database.Update.PlayerHardLevelRecord(hardLevelRecord)

			rsp.Result = mission_api.OUT_RESULT_SUCCEED
			row.Break()
		}
	})
	session.Send(rsp)
}

func buyBossLevelTimes(session *net.Session, LevelId int32, rsp *mission_api.BuyBossLevelTimes_Out) {
	state := module.State(session)
	var IsBossLevel bool = false
	mission_enemys := mission_dat.GetEnemyIdByMissionLevelId(LevelId)
	for _, enemy_id := range mission_enemys {
		missionEnemy := mission_dat.GetMissionLevelEnemyById(int32(enemy_id))
		if missionEnemy.IsBoss == true {
			IsBossLevel = true
			break
		}
	}
	fail.When(!IsBossLevel, "关卡id错误，不是boss关卡")
	state.Database.Select.PlayerMissionLevelRecord(func(row *mdb.PlayerMissionLevelRecordRow) {
		if row.MissionLevelId() == LevelId {
			boss_level_record := row.GoObject()
			if !time.IsInPointHour(player_dat.RESET_BUY_BOSS_LEVEL_TIMES_IN_HOUR, boss_level_record.BuyUpdateTime) {
				boss_level_record.BuyTimes = 0
			}

			maxBuyNum := int32(module.VIP.PrivilegeTimes(state, int16(vip_dat.GOUMAIBOSSGUANQIA)))
			fail.When(boss_level_record.BuyTimes >= maxBuyNum, "boss关卡购买次数已达到今日上限！")
			buyBossLevelCost := buy_boss_level_times_config_dat.GetCost(int64(boss_level_record.BuyTimes + 1))
			fail.When(!module.Player.CheckMoney(state, buyBossLevelCost, player_dat.INGOT), "元宝不足！")
			module.Player.DecMoney(state.Database, state.MoneyState, buyBossLevelCost, player_dat.INGOT, tlog.MFR_BUY_BOSS_LEVEL_TIMES, xdlog.ET_BUY_BOSS_LEVEL_TIMES)
			boss_level_record.BuyTimes += 1
			boss_level_record.BuyUpdateTime = time.GetNowTime()

			state.Database.Update.PlayerMissionLevelRecord(boss_level_record)

			rsp.Result = mission_api.OUT_RESULT_SUCCEED
			row.Break()
		}
	})
	session.Send(rsp)
}

func enterHardLevel(session *net.Session, levelId int32) bool {
	state := module.State(session)
	//关卡信息
	levelInfo := mission_dat.GetMissionLevelById(levelId)
	//难度关卡信息:关卡描述、出现的区域、要求区域关卡功能权值和难度关卡功能权值
	hardLevelInfo := mission_dat.GetHardLevelInfo(levelInfo.ParentId)

	//玩家城镇信息与关卡所在城镇信息
	playerTown := state.Database.Lookup.PlayerTown(state.PlayerId)
	townInfo := town_dat.GetTownWithTownId(hardLevelInfo.TownId)

	//检查玩家进入的关卡与玩家所在城镇是否一样
	fail.When(playerTown.Lock < townInfo.Lock, "关卡所在城镇未开启")

	//检查功能权值
	playerLevel := state.Database.Lookup.PlayerMissionLevel(state.PlayerId)
	fail.When(playerLevel.MaxLock < hardLevelInfo.MissionLevelLock, "hard level not open require higher mission level lock")
	fail.When(playerLevel.AwardLock < hardLevelInfo.MissionLevelLock, "hard level not open require higher mission level award lock")

	playerHardLevel := state.Database.Lookup.PlayerHardLevel(state.PlayerId)
	fail.When(playerHardLevel.Lock < hardLevelInfo.HardLevelLock, "hard level not open require higher hard level lock")

	//检查体力
	if levelInfo.Physical > 0 {
		fail.When(!module.Physical.CheckGE(state, int16(levelInfo.Physical)), "physical not enough when enter hard level")
	}

	levelRecord := module.Mission.GetHardLevelRecordById(state.Database, levelId)
	//检查进入次数
	if levelInfo.DailyNum > 0 && levelRecord != nil {
		// 重置每日进入次数
		if !time.IsInPointHour(player_dat.RESET_HARD_LEVEL_TIMES_IN_HOUR, levelRecord.LastEnterTime) {
			levelRecord.DailyNum = 0
		}
		if !time.IsInPointHour(player_dat.RESET_BUY_HARD_LEVEL_TIMES_IN_HOUR, levelRecord.BuyUpdateTime) {
			levelRecord.BuyTimes = 0
		}
		fail.When(levelRecord.DailyNum-int8(levelRecord.BuyTimes) >= levelInfo.DailyNum, "hard level dailyNum is full")
	}

	var newRecord bool
	if levelRecord == nil {
		newRecord = true
		levelRecord = &mdb.PlayerHardLevelRecord{
			Pid:     state.PlayerId,
			LevelId: levelId,
		}
	}
	levelRecord.LastEnterTime = time.GetNowTime()
	if newRecord {
		state.Database.Insert.PlayerHardLevelRecord(levelRecord)
	} else {
		state.Database.Update.PlayerHardLevelRecord(levelRecord)
	}

	//进入难度关卡先默认扣除1点体力，其余的通关后再扣除
	if levelInfo.Physical > 0 {
		module.Physical.Decrease(state.Database, mission_dat.LEVEL_HARD_PHYSICAL_MIN, tlog.PFR_HARD_LEVEL)
	}
	doEnterLevel(session, levelInfo, battle.BT_HARD_LEVEL)
	tlog.PlayerSystemModuleFlowLog(state.Database, tlog.SMT_HARD_LEVEL)
	tlog.PlayerMissionFlowLog(state.Database, levelId, tlog.ENTER)
	xdlog.MissionLog(state.Database, levelId, xdlog.MA_ENTER)

	return true
}
