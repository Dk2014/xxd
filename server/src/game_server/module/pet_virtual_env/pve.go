package pet_virtual_env

import (
	"core/fail"
	"game_server/battle"
	"game_server/dat/mission_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/tlog"
	"math"
)

func getPveLevelDatByEnemyId(enemyId int32) (*mission_dat.PveLevel, *mission_dat.MissionLevel) {
	missionEnemy := mission_dat.GetMissionLevelEnemyById(enemyId)
	missionLevel := mission_dat.GetMissionLevelById(missionEnemy.MissionLevelId)
	fail.When(missionLevel.ParentType != battle.BT_PVE_LEVEL, "关卡类型错误")
	pveDat := mission_dat.GetPetVirtualEnvLevel(missionLevel.ParentId)
	return pveDat, missionLevel
}

func checkPveLevel(db *mdb.Database, lvDat *mission_dat.PveLevel) {
	mailRole := module.Role.GetMainRole(db)
	fail.When(mailRole.Level < lvDat.Level, "等级未达到关卡要求")
	playerPVEstate := db.Lookup.PlayerPveState(db.PlayerId())
	fail.When(lvDat.Floor > playerPVEstate.MaxAwardedFloor+1, "需要先完成前置关卡")
}

func givePVERegularAward(state *module.SessionState, xdEventType int32) {
	if state.PetVirtualEnvState.EnemyNum > 0 {
		pveDat := mission_dat.GetPetVirtualEnvLevel(state.PetVirtualEnvState.PveLevelId)
		awardNum := pveDat.BasicAwardNum
		awardNum += int16(math.Ceil(float64(state.PetVirtualEnvState.EnemyNum) * (float64(pveDat.AwardFactor+100) / 100)))
		module.Item.AddItem(state.Database, pveDat.AwardItem, awardNum, tlog.IFR_PET_VIRTUAL, xdEventType, "")
	}
}
