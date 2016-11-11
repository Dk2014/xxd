package quest

import (
	"core/fail"
	"core/time"
	"game_server/api/protocol/quest_api"
	"game_server/dat/player_dat"
	"game_server/dat/quest_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/tlog"
)

//根据玩家任务状态来刷领取扩展任务
//TODO 做一些优化避免每次接主线任务都会触发
func takeExtendQuest(db *mdb.Database, mainQuestId int16) {
	mainQuest := quest_dat.GetQuestById(mainQuestId)
	availableExtendQuest := quest_dat.GetExtendQuestByMainQuest(mainQuest.Order)
	recivedQeustSet := map[int32]int8{}
	db.Select.PlayerExtendQuest(func(row *mdb.PlayerExtendQuestRow) {
		recivedQeustSet[row.QuestId()] = 1
	})
	var progress int16
	var status int8
	for _, quest := range availableExtendQuest {
		progress = 0
		status = quest_dat.EXTEND_QUEST_STATUS_NONE
		if _, recived := recivedQeustSet[quest.Id]; !recived {
			if quest.Type == quest_dat.EXTEND_QUEST_TYPE_LOGIN {
				progress = 1
			} else if quest.Type == quest_dat.EXTEND_QUEST_TYPE_MISSION_START {
				progress = module.Mission.CountMissoinStar(db, quest.RelatedMission)
			}
			if progress >= quest.RequiredProgress {
				status = quest_dat.EXTEND_QUEST_STATUS_CAN_AWARD
			}
			db.Insert.PlayerExtendQuest(&mdb.PlayerExtendQuest{
				Pid:      db.PlayerId(),
				State:    status,
				QuestId:  quest.Id,
				Progress: progress,
			})
		}
	}
}

//获取当前任务状态
func getExtendQuestInfo(state *module.SessionState, npcId int32, out *quest_api.GetExtendQuestInfoByNpcId_Out) {
	//玩家连续登录
	playerInfo := state.Database.Lookup.PlayerInfo(state.PlayerId)
	lastLoginTime := playerInfo.LastLoginTime
	var continueLoginCount int16 = 0
	for lastLoginTime < time.GetNowTime() && !time.IsToday(lastLoginTime) {
		lastLoginTime += 86400
		continueLoginCount++
	}

	state.Database.Select.PlayerExtendQuest(func(row *mdb.PlayerExtendQuestRow) {
		questDat := quest_dat.GetExtentQuestById(row.QuestId())
		if questDat.RelatedNpc != npcId {
			return
		}
		quest := row.GoObject()
		if questDat.Type == quest_dat.EXTEND_QUEST_TYPE_LOGIN && quest.State == quest_dat.EXTEND_QUEST_STATUS_NONE {
			//此处无需保存，只需返回一个正确的结果给客户端就可以了
			quest.Progress += continueLoginCount
			if quest.Progress > questDat.RequiredProgress {
				quest.Progress = questDat.RequiredProgress
				quest.State = quest_dat.EXTEND_QUEST_STATUS_CAN_AWARD
			}
		}

		out.Quest = append(out.Quest, quest_api.GetExtendQuestInfoByNpcId_Out_Quest{
			Id:       quest.QuestId,
			Progress: quest.Progress,
			State:    quest.State,
		})
	})
}

//领取扩展任务奖励
func takeExtendQuestAward(state *module.SessionState, questId, xdEventType int32) {
	var playerExtQuest *mdb.PlayerExtendQuest
	state.Database.Select.PlayerExtendQuest(func(row *mdb.PlayerExtendQuestRow) {
		if row.QuestId() == questId {
			playerExtQuest = row.GoObject()
			row.Break()
		}
	})
	fail.When(playerExtQuest == nil, "未接受扩展任务")

	extendQuestDat := quest_dat.GetExtentQuestById(questId)
	switch extendQuestDat.Type {
	case quest_dat.EXTEND_QUEST_TYPE_PAY_INGOT:
		//元宝购买
		module.Player.DecMoney(state.Database, state.MoneyState, int64(extendQuestDat.RequiredProgress), player_dat.INGOT, tlog.MFR_EXTEND_QUEST, xdEventType)
		playerExtQuest.State = quest_dat.EXTEND_QUEST_STATUS_CAN_AWARD
	case quest_dat.EXTEND_QUEST_TYPE_LOGIN:
		//玩家连续登录
		playerInfo := state.Database.Lookup.PlayerInfo(state.PlayerId)
		lastLoginTime := playerInfo.LastLoginTime
		var continueLoginCount int16 = 0
		for lastLoginTime < time.GetNowTime() && !time.IsToday(lastLoginTime) {
			lastLoginTime += 86400
			continueLoginCount++
		}
		playerExtQuest.Progress += continueLoginCount
		if playerExtQuest.Progress >= extendQuestDat.RequiredProgress {
			playerExtQuest.Progress = extendQuestDat.RequiredProgress
			playerExtQuest.State = quest_dat.EXTEND_QUEST_STATUS_CAN_AWARD
		}
	}
	fail.When(playerExtQuest.State != quest_dat.EXTEND_QUEST_STATUS_CAN_AWARD, "未完成扩展任务")

	playerExtQuest.State = quest_dat.EXTEND_QUEST_STATUS_AWARDED
	state.Database.Update.PlayerExtendQuest(playerExtQuest)

	if extendQuestDat.AwardNum1 > 0 {
		module.Item.AddItem(state.Database, extendQuestDat.AwardItem1, extendQuestDat.AwardNum1, tlog.IFR_EXTEND_QUEST, xdEventType, "")
	}
	if extendQuestDat.AwardNum2 > 0 {
		module.Item.AddItem(state.Database, extendQuestDat.AwardItem2, extendQuestDat.AwardNum1, tlog.IFR_EXTEND_QUEST, xdEventType, "")
	}
	if extendQuestDat.AwardNum3 > 0 {
		module.Item.AddItem(state.Database, extendQuestDat.AwardItem3, extendQuestDat.AwardNum1, tlog.IFR_EXTEND_QUEST, xdEventType, "")
	}
	// 所有上阵角色增加经验
	if extendQuestDat.AwardExp > int32(0) {
		module.Role.AddFormRoleExp(state, int64(extendQuestDat.AwardExp), tlog.EFT_EXTEND_QUEST)
	}

	if extendQuestDat.AwardCoins > 0 {
		module.Player.IncMoney(state.Database, state.MoneyState, int64(extendQuestDat.AwardCoins), player_dat.COINS, tlog.MFR_EXTEND_QUEST, xdEventType, "")
	}
}

func refresMissionLevelExtendQuest(db *mdb.Database, missoinId int16, progress int16) {
	extendQuestDat := quest_dat.GetExtendQuestByMissionId(missoinId)
	if extendQuestDat == nil {
		return
	}
	var playerExtQuest *mdb.PlayerExtendQuest
	db.Select.PlayerExtendQuest(func(row *mdb.PlayerExtendQuestRow) {
		if row.QuestId() == extendQuestDat.Id {
			if row.State() == quest_dat.EXTEND_QUEST_STATUS_NONE {
				playerExtQuest = row.GoObject()
			}
			row.Break()
		}
	})
	if playerExtQuest != nil {
		playerExtQuest.Progress += progress
		if playerExtQuest.Progress >= extendQuestDat.RequiredProgress {
			playerExtQuest.Progress = extendQuestDat.RequiredProgress
			playerExtQuest.State = quest_dat.EXTEND_QUEST_STATUS_CAN_AWARD
		}
		db.Update.PlayerExtendQuest(playerExtQuest)
	}
}
