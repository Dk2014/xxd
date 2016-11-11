package mission

import (
	"core/fail"
	"core/log"
	"core/net"
	coreTime "core/time"
	"game_server/api/protocol/mission_api"
	"game_server/battle"
	"game_server/dat/item_dat"
	"game_server/dat/mission_dat"
	"game_server/dat/player_dat"
	"game_server/dat/vip_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/tlog"
	"game_server/xdlog"
	"math/rand"
)

func MissionOpen(session *net.Session, mission_id int16, xdEventType int32) (mission_api.OutResult, bool) {
	state := module.State(session)
	// 检查开启区域，玩家是否在正确的城镇开启区域
	player_town := state.Database.Lookup.PlayerTown(state.PlayerId)
	mission_info := mission_dat.GetMissionById(mission_id)

	// 检查是否有足够的钥匙数量
	if !module.Item.CheckItemNum(state, item_dat.ITEM_MISSION_KEY_ID, int16(mission_info.Keys)) {
		log.Error("MissionOpen key not enough")
		return mission_api.OUT_RESULT_FAILED, false
	}

	player_mission := state.Database.Lookup.PlayerMission(state.PlayerId)
	// 检查是否按顺序开启区域
	if player_mission.MaxOrder+1 != mission_info.Order {
		return mission_api.OUT_RESULT_FAILED, true
	}

	if mission_info.Keys > 0 {
		module.Item.DelItemByItemId(state.Database, item_dat.ITEM_MISSION_KEY_ID, int16(mission_info.Keys), tlog.IFR_MISSION_OPEN, xdEventType)
	}

	player_mission.MaxOrder = mission_info.Order

	state.Database.Update.PlayerMission(player_mission)

	state.Database.Insert.PlayerMissionRecord(&mdb.PlayerMissionRecord{
		Pid:       state.PlayerId,
		TownId:    player_town.TownId,
		MissionId: mission_id,
		OpenTime:  coreTime.GetNowTime(),
	})

	module.Notify.SendPlayerKeyChanged(session, 0, player_mission.MaxOrder)

	return mission_api.OUT_RESULT_SUCCEED, true
}

func GetMissionLevel(state *module.SessionState, mission_id int16, rsp *mission_api.GetMissionLevel_Out) (ret bool) {
	ret = false
	// 检查玩家区域是否开启
	state.Database.Select.PlayerMissionRecord(func(row *mdb.PlayerMissionRecordRow) {
		if row.MissionId() == mission_id {
			ret = true
			row.Break()
		}
	})
	fail.When(ret == false, "mission is not open")

	var (
		dailyNum         int8
		remain_buy_times int16 = 0
		buyTimes         int32 = 0
		isBossLevel      bool  = false
	)
	rsp.Levels = []mission_api.GetMissionLevel_Out_Levels{}
	state.Database.Select.PlayerMissionLevelRecord(func(row *mdb.PlayerMissionLevelRecordRow) {
		if row.MissionId() == mission_id {
			dailyNum = row.DailyNum()
			buyTimes = row.BuyTimes()
			if !coreTime.IsInPointHour(player_dat.RESET_MAIN_LEVEL_TIMES_IN_HOUR, row.LastEnterTime()) {
				dailyNum = 0
			}
			if !coreTime.IsInPointHour(player_dat.RESET_BUY_BOSS_LEVEL_TIMES_IN_HOUR, row.BuyUpdateTime()) {
				buyTimes = 0
			}

			mission_enemys := mission_dat.GetEnemyIdByMissionLevelId(row.MissionLevelId())
			for _, enemy_id := range mission_enemys {
				missionEnemy := mission_dat.GetMissionLevelEnemyById(enemy_id)
				if missionEnemy.IsBoss == true {
					maxBuyNum := module.VIP.PrivilegeTimes(state, vip_dat.GOUMAIBOSSGUANQIA)
					if buyTimes >= int32(maxBuyNum) {
						remain_buy_times = 0
					} else {
						remain_buy_times = maxBuyNum - int16(buyTimes)
					}
					isBossLevel = true
					break
				}
			}
			if !isBossLevel {
				remain_buy_times = -1 //普通关卡不限购买次数
			}
			emptyShadowBits := row.EmptyShadowBits()
			shadowAmount := 0
			for _, shadedId := range mission_dat.GetShadedIdsByMissionId(row.MissionLevelId()) {
				shadow, _ := mission_dat.GetShadedMission(shadedId)
				if emptyShadowBits&(0x1<<uint(shadow.Order-1)) <= 0 {
					shadowAmount++
				}
			}
			rsp.Levels = append(rsp.Levels, mission_api.GetMissionLevel_Out_Levels{
				LevelId:        row.MissionLevelId(),
				RoundNum:       row.Round(),
				DailyNum:       dailyNum,
				WaitingShadows: int8(shadowAmount),
				RemainBuyTimes: remain_buy_times,
			})
		}
	})

	return
}

func EnterMissionLevel(session *net.Session, mission_level_id int32, rsp *mission_api.EnterLevel_Out) (ret mission_api.OutResult) {
	ret = mission_api.OUT_RESULT_FAILED

	state := module.State(session)

	// 检查玩家进入的关卡城镇是否是玩家当前所在城镇
	mission_level_info := mission_dat.GetMissionLevelById(mission_level_id)
	mission_info := mission_dat.GetMissionById(mission_level_info.MissionId)

	// 检查当前区域是否开启（如果关卡开启区域应该是正常开启的）
	// 检查关卡是否开启
	playerLevel := state.Database.Lookup.PlayerMissionLevel(state.PlayerId)
	fail.When(playerLevel.MaxLock < mission_level_info.Lock, "mission level not open")

	// 检查体力
	if mission_level_info.Physical > 0 {
		fail.When(!module.Physical.CheckGE(state, int16(mission_level_info.Physical)), "physical not enough when enter mission level")
	}

	levelRecord := module.Mission.GetMissionLevelRecord(state.Database, mission_level_id)
	// 关卡进入有次数限制
	if mission_level_info.DailyNum > 0 && levelRecord != nil {
		// 重置每日进入次数
		if !coreTime.IsInPointHour(player_dat.RESET_MAIN_LEVEL_TIMES_IN_HOUR, levelRecord.LastEnterTime) {
			levelRecord.DailyNum = 0
		}
		fail.When(levelRecord.DailyNum >= mission_level_info.DailyNum, "mission level dailyNum is full")
	}

	newRecord := false
	nowTime := coreTime.GetNowTime()
	if levelRecord == nil {
		newRecord = true
		levelRecord = &mdb.PlayerMissionLevelRecord{
			Pid:            state.PlayerId,
			MissionId:      mission_info.Id,
			MissionLevelId: mission_level_id,
			OpenTime:       nowTime,
		}
	}

	levelRecord.LastEnterTime = nowTime
	if newRecord {
		state.Database.Insert.PlayerMissionLevelRecord(levelRecord)
	} else {
		state.Database.Update.PlayerMissionLevelRecord(levelRecord)
	}

	emptyShadows := levelRecord.EmptyShadowBits

	// 扣除体力
	if mission_level_info.Physical > 0 {
		// 先扣掉至少需要的体力，通关后再扣除剩余的
		module.Physical.Decrease(state.Database, mission_dat.LEVEL_PHYSICAL_MIN, tlog.PFR_MISSION_LEVEL)
	}

	doEnterLevel(session, mission_level_info, battle.BT_MISSION_LEVEL)
	ret = mission_api.OUT_RESULT_SUCCEED

	tlog.PlayerMissionFlowLog(state.Database, mission_level_id, tlog.ENTER)
	xdlog.MissionLog(state.Database, mission_level_id, xdlog.MA_ENTER)

	// 随机小宝箱
	smallBoxids := initSmallBoxAndSetMissionState(mission_level_id, state.MissionLevelState)
	for _, id := range smallBoxids {
		rsp.Smallbox = append(rsp.Smallbox, mission_api.EnterLevel_Out_Smallbox{
			BoxId: id,
		})
	}
	mys := initMengYaoAndSetMissionState(mission_level_id, state.MissionLevelState)
	for _, id := range mys {
		rsp.MengYao = append(rsp.MengYao, mission_api.EnterLevel_Out_MengYao{
			MyId: id,
		})
	}
	// 残存影之间隙
	shadows := initShadowsAndSetMissionState(mission_level_id, emptyShadows, state.MissionLevelState)
	for _, id := range shadows {
		rsp.Shadow = append(rsp.Shadow, mission_api.EnterLevel_Out_Shadow{
			ShadedId: id,
		})
	}
	return
}

func initMengYaoAndSetMissionState(levelId int32, missionLevelState *module.MissionLevelState) (ids []int32) {
	if missionLevelState.MengYaoList == nil {
		missionLevelState.MengYaoList = make(map[int32]int8, 1)
	}
	if mengYao, ok := mission_dat.GetMissionMengYao(levelId); ok {
		randNum := int8(rand.Intn(100) + 1)
		chance := int8(0)
		for _, my := range mengYao {
			if randNum > chance && randNum <= chance+my.Probability {
				missionLevelState.MengYaoList[my.Id] = my.Effect
				ids = append(ids, my.Id)
			}
			chance += my.Probability
		}
	}

	return
}

func initSmallBoxAndSetMissionState(levelId int32, missionLevelState *module.MissionLevelState) (ids []int32) {
	if missionLevelState.SmallBoxList == nil {
		missionLevelState.SmallBoxList = make(map[int32]int8, 1)
	}
	if smallBox, ok := mission_dat.GetMissionSmallBox(levelId); ok {
		randNum := int8(rand.Intn(100) + 1)
		chance := int8(0)
		for _, box := range smallBox {
			if randNum > chance && randNum <= chance+box.Probability {
				missionLevelState.SmallBoxList[box.Id] = box.ItemsCount
				ids = append(ids, box.Id)
			}
			chance += box.Probability
		}
	}
	return
}

func initShadowsAndSetMissionState(levelId int32, emptyShadows int16, missionLevelState *module.MissionLevelState) (ids []int32) {
	if missionLevelState.ShadowList == nil {
		missionLevelState.ShadowList = make(map[int32]int, 1)
	}
	for _, shadedId := range mission_dat.GetShadedIdsByMissionId(levelId) {
		shadow, _ := mission_dat.GetShadedMission(shadedId)
		if emptyShadows&(0x1<<uint(shadow.Order-1)) <= 0 {
			missionLevelState.ShadowList[shadow.Id] = 0
			ids = append(ids, shadow.Id)
		}
	}
	enemyIds := mission_dat.GetEnemyIdByMissionLevelId(levelId)
	for _, enemyId := range enemyIds {
		enemy := mission_dat.GetMissionLevelEnemyById(enemyId)
		shadedId := enemy.ShadedMissionId
		if count, exist := missionLevelState.ShadowList[shadedId]; exist {
			missionLevelState.ShadowList[shadedId] = count + 1
		}
	}
	return
}

func getMissionLevelByItemId(session *net.Session, itemId int16, rsp *mission_api.GetMissionLevelByItemId_Out) {
	missionLevelDatByItem := mission_dat.GetMissionLevelByItemId(itemId)
	if missionLevelDatByItem == nil || len(missionLevelDatByItem) == 0 {
		return
	}
	state := module.State(session)
	var minIndex int = 0
	var lock int32 = 1 << 30
	var levels [3]*mdb.PlayerMissionLevelRecord
	var levelsDat [3]*mission_dat.MissionLevel
	state.Database.Select.PlayerMissionLevelRecord(func(row *mdb.PlayerMissionLevelRecordRow) {
		missionLevelDat, ok := missionLevelDatByItem[row.MissionLevelId()]
		if !ok {
			return
		}
		if missionLevelDat.DailyNum > 0 && row.DailyNum() >= missionLevelDat.DailyNum &&
			coreTime.IsInPointHour(player_dat.RESET_MAIN_LEVEL_TIMES_IN_HOUR,
				row.LastEnterTime()) {
			return
		}

		for i, lv := range levels {
			if lv == nil { //如果有空位则使用空位
				levelsDat[i] = missionLevelDat
				levels[i] = row.GoObject()
				return
			} else if levelsDat[i].Lock < lock { //计算权值最小关卡
				minIndex = i
				lock = levelsDat[i].Lock
			}
		}
		if missionLevelDat.Lock > levelsDat[minIndex].Lock {
			levelsDat[minIndex] = missionLevelDat
			levels[minIndex] = row.GoObject()
		}

	})
	var dailyNum int8
	for _, lv := range levels {
		if lv != nil {
			if !coreTime.IsInPointHour(player_dat.RESET_MAIN_LEVEL_TIMES_IN_HOUR, lv.LastEnterTime) {
				dailyNum = 0
			} else {
				dailyNum = lv.DailyNum
			}

			rsp.Levels = append(rsp.Levels, mission_api.GetMissionLevelByItemId_Out_Levels{
				LevelId:  lv.MissionLevelId,
				RoundNum: lv.Round,
				DailyNum: dailyNum,
			})
		}
	}
}

/*
func GetMissionToallStar(session *net.Session, townID int16) total ,awardLevel int32{
	state := module.State(session)

	state.Select.PlayerMissionLevelRecord(func(row *mdb.PlayerMissionLevelRecordRow) {
		missionID := mission_dat.GetMissionIDByLevelID(row.MissionLevelId())
		townid := mission_dat.GetTwonIDByMission()
		if townid == townID {
			total += row.Round()
		}

	})


	return
}

func GetMissionStartAward() {

}

*/
