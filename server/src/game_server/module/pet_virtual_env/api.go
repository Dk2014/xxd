package pet_virtual_env

import (
	"core/fail"
	"core/net"
	"core/time"
	"game_server/api/protocol/pet_virtual_env_api"
	"game_server/dat/mission_dat"
	"game_server/dat/player_dat"
	"game_server/dat/quest_dat"
	"game_server/module"
	"game_server/tlog"
	"game_server/xdlog"
	"math"
)

type PetVirtualEnvAPI struct{}

func init() {
	pet_virtual_env_api.SetInHandler(PetVirtualEnvAPI{})
}

func (api PetVirtualEnvAPI) Info(session *net.Session, in *pet_virtual_env_api.Info_In) {
	state := module.State(session)
	fail.When(!module.Player.IsOpenFunc(state.Database, player_dat.FUNC_BATTLE_PET), "灵宠功能未开启")
	playerPveState := state.Database.Lookup.PlayerPveState(state.PlayerId)
	if !time.IsInPointHour(player_dat.RESET_PVE_LEVEL_IN_HOUR, playerPveState.EnterTime) {
		playerPveState.DailyNum = 0
	}
	state.Database.Update.PlayerPveState(playerPveState)
	session.Send(&pet_virtual_env_api.Info_Out{
		DailyNum:              playerPveState.DailyNum,
		MaxFloor:              playerPveState.MaxPassedFloor,
		MaxAwardedFloor:       playerPveState.MaxAwardedFloor,
		UnpassedFloorEnemyNum: playerPveState.UnpassedFloorEnemyNum,
	})
}

func (api PetVirtualEnvAPI) TakeAward(session *net.Session, in *pet_virtual_env_api.TakeAward_In) {
	state := module.State(session)
	fail.When(!module.Player.IsOpenFunc(state.Database, player_dat.FUNC_BATTLE_PET), "灵宠功能未开启")
	playerPveState := state.Database.Lookup.PlayerPveState(state.PlayerId)
	fail.When(playerPveState.MaxAwardedFloor >= playerPveState.MaxPassedFloor, "灵宠环境奖励不可领取")
	pveLvDat := mission_dat.GetPetVirtualEnvLevelByFloor(playerPveState.MaxPassedFloor)
	module.Item.AddItem(state.Database, pveLvDat.AwardItem, pveLvDat.AwardNum, tlog.IFR_PET_VIRTUAL, xdlog.ET_PET_VIRTUAL, "")
	playerPveState.MaxAwardedFloor = playerPveState.MaxPassedFloor
	state.Database.Update.PlayerPveState(playerPveState)
	//由于需要领取奖励之后才能进入下一关卡，因此发送空的返回给客户端，以便告知客户端领取奖励成功
	session.Send(&pet_virtual_env_api.TakeAward_Out{})
}

func (api PetVirtualEnvAPI) AutoFight(session *net.Session, in *pet_virtual_env_api.AutoFight_In) {
	state := module.State(session)

	playerPveState := state.Database.Lookup.PlayerPveState(state.PlayerId)
	if !time.IsInPointHour(player_dat.RESET_PVE_LEVEL_IN_HOUR, playerPveState.EnterTime) {
		playerPveState.DailyNum = 0
	}
	fail.When(playerPveState.DailyNum >= mission_dat.PVE_LEVEL_DAILY_NUM, "灵宠环境达到每天最大进入次数")

	pveDat := mission_dat.GetPetVirtualEnvLevelByFloor(in.Floor)
	fail.When(pveDat.Floor > playerPveState.MaxPassedFloor, "尚未通关不能扫荡")

	awardNum := pveDat.BasicAwardNum
	awardNum += int16(math.Ceil(float64(pveDat.MosterNum) * (float64(pveDat.AwardFactor+100) / 100)))
	module.Item.AddItem(state.Database, pveDat.AwardItem, awardNum, tlog.IFR_PET_VIRTUAL, xdlog.ET_PET_VIRTUAL_AUTO_FIGHT, "")

	playerPveState.EnterTime = time.GetNowTime()
	playerPveState.DailyNum += 1
	state.Database.Update.PlayerPveState(playerPveState)

	//刷新每日任务状态
	module.Quest.RefreshDailyQuest(state, quest_dat.DAILY_QUEST_CLASS_PET_VIRENV)
}
