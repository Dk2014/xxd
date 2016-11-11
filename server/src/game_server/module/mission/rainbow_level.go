package mission

import (
	"core/fail"
	"core/net"
	"game_server/api/protocol/mission_api"
	"game_server/battle"
	"game_server/dat/mission_dat"
	"game_server/dat/rainbow_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/tlog"
	"game_server/xdlog"
)

func enterRainbowLevel(session *net.Session, levelId int32, out *mission_api.EnterRainbowLevel_Out) {
	state := module.State(session)

	//查询玩家在当前关卡的状态，
	playerRainbowLevelState := state.Database.Lookup.PlayerRainbowLevel(state.PlayerId)

	if playerRainbowLevelState.Status == rainbow_dat.RAINBOW_LEVEL_STATUS_NO_MORE {
		//可能存在增加关卡的情况，进入时也需要更新玩家关卡信息。
		fail.When(playerRainbowLevelState.Segment > rainbow_dat.MaxRainbowLevelSegment, "没有跟多彩虹关卡")
		playerRainbowLevelState.Segment += 1
		playerRainbowLevelState.Order = 1
		playerRainbowLevelState.Status = rainbow_dat.RAINBOW_LEVEL_STATUS_NEVER_PASS
		state.Database.Update.PlayerRainbowLevel(playerRainbowLevelState)
	}
	if playerRainbowLevelState.MaxOpenSegment < playerRainbowLevelState.Segment {
		playerRainbowLevelState.MaxOpenSegment = playerRainbowLevelState.Segment
		state.Database.Update.PlayerRainbowLevel(playerRainbowLevelState)
	}

	//根据服务器数据查询当前彩虹关卡的ID
	levelIdFromState := rainbow_dat.GetRainbowLevelId(playerRainbowLevelState.Segment, playerRainbowLevelState.Order)
	fail.When(levelIdFromState != levelId, "彩虹关卡尚未开启或者已经通过")

	//查找关卡信息
	levelInfo := mission_dat.GetMissionLevelById(levelId)

	// 先清理之前的关卡数据
	module.Town.LeaveTown(session)
	module.Mission.LeaveMissionLevel(state)

	//加载彩虹关卡状态
	if state.RainbowLevelState == nil {
		state.RevertRainbowLevelState()
	}

	state.RainbowLevelState.Update(state)

	//构造 MissionLevel
	doEnterLevel(session, levelInfo, battle.BT_RAINBOW_LEVEL)
	tlog.PlayerSystemModuleFlowLog(state.Database, tlog.SMT_RAINBOW)
	state.MissionLevelState.MaxRound = rainbow_dat.MAX_BATTLE_ROUND
	//进入彩虹关卡时，把彩虹关卡的的状态信息同步到MissionLevelState
	state.MissionLevelState.AttackerInfo = state.RainbowLevelState.AttackerInfo
	state.MissionLevelState.UsedGhost = state.RainbowLevelState.UsedGhost
	state.MissionLevelState.CalledBattlePet = state.RainbowLevelState.CalledBattlePet

	state.RainbowLevelState.AwardCoin = levelInfo.AwardCoin
	state.RainbowLevelState.AwardExp = levelInfo.AwardExp
	state.RainbowLevelState.AwardRelationship = levelInfo.AwardRelationship

	//加载伙伴进阶技能
	for skillId, restReleaseNum := range state.RainbowLevelState.SkillReleaseNum {
		state.MissionLevelState.SkillReleaseNum[skillId] = restReleaseNum
	}

	for ghostId, _ := range state.RainbowLevelState.UsedGhost {
		out.UsedGhost = append(out.UsedGhost, mission_api.EnterRainbowLevel_Out_UsedGhost{
			GhostId: ghostId,
		})
	}
	for petId, _ := range state.RainbowLevelState.CalledBattlePet {
		out.CalledPet = append(out.CalledPet, mission_api.EnterRainbowLevel_Out_CalledPet{
			PetId: petId,
		})
	}
	tlog.PlayerMissionFlowLog(state.Database, levelId, tlog.ENTER)
	xdlog.MissionLog(state.Database, levelId, xdlog.MA_ENTER)
}

func findCalledPetGridInRainbowLevel(state *module.SessionState) (calledGrids []int8) {

	state.Database.Select.PlayerBattlePetGrid(func(row *mdb.PlayerBattlePetGridRow) {
		if _, called := state.RainbowLevelState.CalledBattlePet[row.BattlePetId()]; called {
			calledGrids = append(calledGrids, row.GridId())
		}
	})

	return
}
