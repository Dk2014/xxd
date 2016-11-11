package meditation

import (
	"core/fail"
	"core/time"
	//"game_server/dat/item_dat"
	"game_server/dat/meditation_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/module/rpc"
	"game_server/tlog"
)

func meditationInfo(state *module.SessionState) (int32, int32) {
	meditationState := state.Database.Lookup.PlayerMeditationState(state.PlayerId)
	meditateDuration := int32(time.GetNowTime() - meditationState.StartTimestamp)
	if meditateDuration > meditation_dat.MAX_MEDITATION_TIME {
		meditateDuration = meditation_dat.MAX_MEDITATION_TIME
	}

	if meditationState.StartTimestamp != meditation_dat.MEDITATION_STOP && meditateDuration > 0 {
		return meditateDuration, meditationState.AccumulateTime + int32(meditateDuration)
	} else {
		return 0, meditationState.AccumulateTime
	}
}

func startMeditation(db *mdb.Database, timestamp int64, inClub bool) {
	fail.When(timestamp == meditation_dat.MEDITATION_STOP, "参数错误：错误的开始打坐时间")
	meditationState := db.Lookup.PlayerMeditationState(db.PlayerId())

	//if meditationState.StartTimestamp != meditation_dat.MEDITATION_STOP && meditationState.StartTimestamp >= timestamp {
	if meditationState.StartTimestamp != meditation_dat.MEDITATION_STOP {
		//已在打坐状态
		return
	}
	if inClub {
		rpc.RemoteGetClubHouseMeditation(db.PlayerId(), meditation_dat.CLUBHOUSE_MEDITATION_START)
	}
	meditationState.StartTimestamp = timestamp
	db.Update.PlayerMeditationState(meditationState)
}

func stopMeditation(db *mdb.Database, inClub bool) {
	if inClub {
		rpc.RemoteGetClubHouseMeditation(db.PlayerId(), meditation_dat.CLUBHOUSE_MEDITATION_STOP)
	}
	meditationState := db.Lookup.PlayerMeditationState(db.PlayerId())
	now := time.GetNowTime()
	if meditationState.StartTimestamp != meditation_dat.MEDITATION_STOP && meditationState.StartTimestamp < now {
		meditateDuration := int32(now - meditationState.StartTimestamp)
		if meditateDuration > meditation_dat.MAX_MEDITATION_TIME {
			meditateDuration = meditation_dat.MAX_MEDITATION_TIME
		}
		mainRoleLv := module.Role.GetMainRole(db).Level
		awardExp := (meditateDuration / meditation_dat.EXP_UNIT_TIME) * int32(meditation_dat.GetMetationExpByLevel(mainRoleLv))

		keyAccumulateTime := meditationState.AccumulateTime + meditateDuration
		//awardKeys := keyAccumulateTime / meditation_dat.KEY_UNIT_TIME

		meditationState.StartTimestamp = meditation_dat.MEDITATION_STOP
		meditationState.AccumulateTime = keyAccumulateTime % meditation_dat.KEY_UNIT_TIME
		db.Update.PlayerMeditationState(meditationState)

		//if awardKeys > 0 {
		//	module.Item.AddItem(db, item_dat.ITEM_MISSION_KEY_ID, int16(awardKeys), tlog.IFR_MEDITATION)
		//}

		if awardExp > 0 {
			mainRole := module.Role.GetMainRole(db)
			module.Role.AddRoleExp(db, mainRole.RoleId, int64(awardExp), mainRole.RoleId, tlog.EFT_MEDITATION)
		}
	}
}
