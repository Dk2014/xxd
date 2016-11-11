package quest

import (
	"game_server/dat/player_dat"
	"game_server/dat/quest_dat"
	"game_server/dat/team_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/tlog"
)

func awardQuest(state *module.SessionState, questInfo *quest_dat.Quest, xdEventType int32) {
	// 所有上阵角色增加经验
	if questInfo.AwardExp > int32(0) {
		module.Role.AddFormRoleExp(state, int64(questInfo.AwardExp), tlog.EFT_QUEST_AWARD)
	}

	if questInfo.AwardCoins > 0 {
		module.Player.IncMoney(state.Database, state.MoneyState, questInfo.AwardCoins, player_dat.COINS, tlog.MFR_QUEST_AWARD, xdEventType, "")
	}

	if questInfo.AwardPhysical > 0 {
		module.Physical.AwardIncrease(state, int16(questInfo.AwardPhysical), tlog.PFR_QUEST_AWARD)
	}

	if questInfo.AwardFuncKey > 0 {
		module.Player.RefreshPlayerFuncKey(state.Database, questInfo.AwardFuncKey)
	}

	if questInfo.AwardTownKey > 0 {
		module.Town.SetTownLock(state, questInfo.AwardTownKey)
	}

	if questInfo.AwardRoleId > int8(0) {
		//TODO questInfo.AwardRoleLevel 应该是int16类型
		module.Role.AddBuddyRole(state, questInfo.AwardRoleId, int16(questInfo.AwardRoleLevel))
		// 如果奖励朱媛媛，默认上阵到4号位（剧情需要）
		/*
			if questInfo.AwardRoleId == 4 {
				playerForm := state.Database.Lookup.PlayerFormation(state.PlayerId)
				arrangeBuddyRole(playerForm, questInfo.AwardRoleId, 4)
				state.Database.Update.PlayerFormation(playerForm)
				// 奖励袁铭志默认上阵
			} else if questInfo.AwardRoleId == 3 {
				playerForm := state.Database.Lookup.PlayerFormation(state.PlayerId)
				playerForm.Pos0 = int8(module.Role.GetMainRole(state.Database).RoleId) //主角在0号位
				playerForm.Pos1 = team_dat.POS_NO_ROLE
				playerForm.Pos2 = 3 //袁铭志放在2号位
				playerForm.Pos3 = team_dat.POS_NO_ROLE
				playerForm.Pos4 = 4 //朱媛媛在4号位
				playerForm.Pos5 = team_dat.POS_NO_ROLE
				playerForm.Pos6 = team_dat.POS_NO_ROLE
				playerForm.Pos7 = team_dat.POS_NO_ROLE
				playerForm.Pos8 = team_dat.POS_NO_ROLE
				state.Database.Update.PlayerFormation(playerForm)
			}
		*/
	}

	if questInfo.AwardMissionLevelLock > 0 {
		playerLevel := state.Database.Lookup.PlayerMissionLevel(state.PlayerId)
		if playerLevel.Lock < questInfo.AwardMissionLevelLock {
			module.Mission.SetMissionLevelLock(state, questInfo.AwardMissionLevelLock, playerLevel, true)
		}
	}

	if questInfo.AwardItem1Id > int16(0) {
		module.Item.AddItem(state.Database, questInfo.AwardItem1Id, questInfo.AwardItem1Num, tlog.IFR_QUEST_AWARD, xdEventType, "")
	}

	if questInfo.AwardItem2Id > int16(0) {
		module.Item.AddItem(state.Database, questInfo.AwardItem2Id, questInfo.AwardItem2Num, tlog.IFR_QUEST_AWARD, xdEventType, "")
	}

	if questInfo.AwardItem3Id > int16(0) {
		module.Item.AddItem(state.Database, questInfo.AwardItem3Id, questInfo.AwardItem3Num, tlog.IFR_QUEST_AWARD, xdEventType, "")
	}

	if questInfo.AwardItem4Id > int16(0) {
		module.Item.AddItem(state.Database, questInfo.AwardItem4Id, questInfo.AwardItem4Num, tlog.IFR_QUEST_AWARD, xdEventType, "")
	}
	//主线任务完成奖励星数
	// stars := quest_dat.MAIN_QUEST_AWARDE_STARS
	// playerStarsInfo := state.Database.Lookup.PlayerDailyQuestStarAwardInfo(state.PlayerId)
	// if playerStarsInfo != nil {

	// 	if !time.IsToday(playerStarsInfo.Lastupdatetime) {
	// 		playerStarsInfo.Awarded = ""
	// 		playerStarsInfo.Lastupdatetime = time.GetNowTime()
	// 		playerStarsInfo.Stars = 0
	// 	}
	// 	playerStarsInfo.Stars += int32(stars)
	// 	state.Database.Update.PlayerDailyQuestStarAwardInfo(playerStarsInfo)
	// } else {
	// 	playerStarsInfo := &mdb.PlayerDailyQuestStarAwardInfo{
	// 		Pid:            state.PlayerId,
	// 		Stars:          int32(stars),
	// 		Lastupdatetime: time.GetNowTime(),
	// 	}
	// 	state.Database.Insert.PlayerDailyQuestStarAwardInfo(playerStarsInfo)
	// }
}

func setNextQuest(state *module.SessionState, playerQuest *mdb.PlayerQuest, questInfo *quest_dat.Quest) {
	// 完成所有主线任务
	if questInfo.NextQuestId == quest_dat.QUEST_ID_NO_MORE {
		playerQuest.State = quest_dat.QUEST_STATUS_AWARD
		takeExtendQuest(state.Database, playerQuest.QuestId)
		return
	}

	var newQuest *quest_dat.Quest
	mainRole := module.Role.GetMainRole(state.Database)

	// 已有新的任务，但是是未接状态
	newQuest = quest_dat.GetQuestById(questInfo.NextQuestId)
	playerQuest.QuestId = questInfo.NextQuestId

	if mainRole.Level < int16(newQuest.RequireLevel) {
		playerQuest.State = quest_dat.QUEST_STATUS_NO_RECEIVE
	} else {
		playerQuest.State = quest_dat.QUEST_STATUS_ALREADY_RECEIVE
		//伙伴招募任务
		takeExtendQuest(state.Database, playerQuest.QuestId)
	}
}

//由于剧情需要，需要把某些奖励的伙伴角色自动上阵
//自动上阵的伙伴有默认的位置，如果这个位置被占用
//则需要为占用其位置的角色挑选其他空位
//无反射版本
func arrangeBuddyRole(form *mdb.PlayerFormation, roleId, pos int8) {
	var originRole int8
	switch pos {
	case 0:
		originRole = form.Pos0
		form.Pos0 = roleId
	case 1:
		originRole = form.Pos1
		form.Pos1 = roleId
	case 2:
		originRole = form.Pos2
		form.Pos2 = roleId
	case 3:
		originRole = form.Pos3
		form.Pos3 = roleId
	case 4:
		originRole = form.Pos4
		form.Pos4 = roleId
	case 5:
		originRole = form.Pos5
		form.Pos5 = roleId
	case 6:
		originRole = form.Pos6
		form.Pos6 = roleId
	case 7:
		originRole = form.Pos7
		form.Pos7 = roleId
	case 8:
		originRole = form.Pos8
		form.Pos8 = roleId
	default:
		panic("invalid formation position")
	}
	if originRole == team_dat.POS_NO_ROLE {
		return
	}
	if form.Pos0 == team_dat.POS_NO_ROLE {
		form.Pos0 = originRole
	} else if form.Pos1 == team_dat.POS_NO_ROLE {
		form.Pos1 = originRole
	} else if form.Pos2 == team_dat.POS_NO_ROLE {
		form.Pos2 = originRole
	} else if form.Pos3 == team_dat.POS_NO_ROLE {
		form.Pos3 = originRole
	} else if form.Pos4 == team_dat.POS_NO_ROLE {
		form.Pos4 = originRole
	} else if form.Pos5 == team_dat.POS_NO_ROLE {
		form.Pos5 = originRole
	} else if form.Pos6 == team_dat.POS_NO_ROLE {
		form.Pos6 = originRole
	} else if form.Pos7 == team_dat.POS_NO_ROLE {
		form.Pos7 = originRole
	} else if form.Pos8 == team_dat.POS_NO_ROLE {
		form.Pos8 = originRole
	}
}
