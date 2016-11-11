package quest

import (
	"core/fail"
	"core/net"
	"core/time"
	"fmt"
	"game_server/api/protocol/quest_api"
	"game_server/dat/item_dat"
	"game_server/dat/player_dat"
	"game_server/dat/quest_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/tlog"
	"strings"
	goTime "time"
)

func (mod QuestMod) RefreshDailyQuest(state *module.SessionState, class int16) {
	// 凌晨需要更新每日任务
	if state.RefreshDailyQuestAtWeeHours {
		state.RefreshDailyQuestAtWeeHours = false
		roleLevel := module.Role.GetMainRole(state.Database).Level
		module.Quest.UpdatePlayerDailyQuest(state.Database, roleLevel)
	}

	var playerQuest *mdb.PlayerDailyQuest
	state.Database.Select.PlayerDailyQuest(func(row *mdb.PlayerDailyQuestRow) {
		if row.Class() == class {
			playerQuest = row.GoObject()
			row.Break()
		}
	})

	// 玩家还不存在此日常任务
	if playerQuest == nil {
		return
	}

	// 此日常任务是已完成状态，等待领取奖励
	if playerQuest.AwardStatus == quest_dat.DAILY_QUEST_STATUS_CAN_AWARD {
		return
	}

	quest := quest_dat.GetDailyQuestWithQuestId(playerQuest.QuestId)
	playerQuest.FinishCount += 1
	playerQuest.LastUpdateTime = time.GetNowTime()

	if quest.RequireCount == playerQuest.FinishCount {
		playerQuest.AwardStatus = quest_dat.DAILY_QUEST_STATUS_CAN_AWARD
	}

	state.Database.Update.PlayerDailyQuest(playerQuest)
	if session, ok := module.Player.GetPlayerOnline(state.PlayerId); ok {
		session.Send(&quest_api.NotifyDailyChange_Out{
			Id:          playerQuest.QuestId,
			FinishCount: playerQuest.FinishCount,
			AwardState:  dailyQuestState2QuestState(playerQuest.AwardStatus),
		})
	}
}

// 更新玩家日常任务（登陆、升级、每日凌晨）
func (mod QuestMod) UpdatePlayerDailyQuest(db *mdb.Database, roleLevel int16) {
	dailyQuest := quest_dat.GetDailyQuestWithLevel(int32(roleLevel))

	haveQuest := map[int16]bool{}
	var delQuest []*mdb.PlayerDailyQuest
	var playerQuest *mdb.PlayerDailyQuest

	//	删除不是今天的每日任务;更新今天的每日任务
	db.Select.PlayerDailyQuest(func(row *mdb.PlayerDailyQuestRow) {
		playerQuest = row.GoObject()

		obsolete := (playerQuest.LastUpdateTime > 0 && !time.IsInPointHour(player_dat.RESET_DAILY_QUEST_IN_HOUR, playerQuest.LastUpdateTime))

		if _, ok := dailyQuest[playerQuest.QuestId]; !ok && (playerQuest.LastUpdateTime == 0 || obsolete) {
			delQuest = append(delQuest, playerQuest)

		} else {
			haveQuest[playerQuest.Class] = true
			if obsolete {
				playerQuest.FinishCount = 0
				playerQuest.LastUpdateTime = 0
				playerQuest.AwardStatus = quest_dat.DAILY_QUEST_STATUS_NONE
				db.Update.PlayerDailyQuest(playerQuest)
			}
		}
	})

	for _, quest := range delQuest {
		db.Delete.PlayerDailyQuest(quest)
	}

	// 添加每日新任务
	for id, quest := range dailyQuest {
		if _, ok := haveQuest[quest.Class]; !ok {
			db.Insert.PlayerDailyQuest(&mdb.PlayerDailyQuest{
				Pid:            db.PlayerId(),
				QuestId:        id,
				FinishCount:    0,
				LastUpdateTime: 0,
				AwardStatus:    quest_dat.DAILY_QUEST_STATUS_NONE,
				Class:          quest.Class,
			})
			tlog.PlayerQuestFlowLog(db, int32(id), tlog.QT_DAILY, tlog.QUEST_ACCEPT)
		}
	}
}

func (mod QuestMod) StartDailyQuestTimer(state *module.SessionState) {
	t := goTime.Now()
	st := t.Unix()
	et := time.GetNextTodayPointHour(player_dat.RESET_DAILY_QUEST_IN_HOUR, t)
	state.TimerMgr.Start(module.TIMER_DAILY_QUEST, goTime.Duration(et-st)*goTime.Second, func(s *module.SessionState) {
		s.RefreshDailyQuestAtWeeHours = true
	})
}

func awardDailyQuest(session *net.Session, id int16, xdEventType int32) {
	var playerQuest *mdb.PlayerDailyQuest
	state := module.State(session)
	tlog.PlayerSystemModuleFlowLog(state.Database, tlog.SMT_QUEST)
	state.Database.Select.PlayerDailyQuest(func(row *mdb.PlayerDailyQuestRow) {
		if row.QuestId() == id {
			playerQuest = row.GoObject()
			row.Break()
		}
	})

	fail.When(playerQuest == nil, "[awardDailyQuest] error id")
	fail.When(playerQuest.AwardStatus != quest_dat.DAILY_QUEST_STATUS_CAN_AWARD, "[awardDailyQuest] not finish")

	playerQuest.AwardStatus = quest_dat.DAILY_QUEST_STATUS_HAVE_AWARD
	state.Database.Update.PlayerDailyQuest(playerQuest)
	tlog.PlayerQuestFlowLog(state.Database, int32(id), tlog.QT_DAILY, tlog.QUEST_FINISH)

	questInfo := quest_dat.GetDailyQuestWithQuestId(playerQuest.QuestId)

	// 所有上阵角色增加经验
	if questInfo.AwardExp > int32(0) {
		module.Role.AddFormRoleExp(state, int64(questInfo.AwardExp), tlog.EFT_DAILY_QUEST)
	}

	if questInfo.AwardCoins > 0 {
		module.Player.IncMoney(state.Database, state.MoneyState, questInfo.AwardCoins, player_dat.COINS, tlog.MFR_DAILY_QUEST, xdEventType, "")
	}

	if questInfo.AwardPhysical > 0 {
		module.Physical.AwardIncrease(state, int16(questInfo.AwardPhysical), tlog.PFR_DAILY_QUEST)
	}

	if questInfo.AwardIngot > 0 {
		module.Player.IncMoney(state.Database, state.MoneyState, int64(questInfo.AwardIngot), player_dat.INGOT, tlog.MFR_DAILY_QUEST, xdEventType, "")
	}

	haveAddKey := false

	if questInfo.AwardItem1Id > int16(0) {
		module.Item.AddItem(state.Database, questInfo.AwardItem1Id, questInfo.AwardItem1Num, tlog.IFR_DAILY_QUEST, xdEventType, "")
		if questInfo.AwardItem1Id == item_dat.ITEM_MISSION_KEY_ID {
			haveAddKey = true
		}
	}

	if questInfo.AwardItem2Id > int16(0) {
		module.Item.AddItem(state.Database, questInfo.AwardItem2Id, questInfo.AwardItem2Num, tlog.IFR_DAILY_QUEST, xdEventType, "")
		if questInfo.AwardItem2Id == item_dat.ITEM_MISSION_KEY_ID {
			haveAddKey = true
		}
	}

	if questInfo.AwardItem3Id > int16(0) {
		module.Item.AddItem(state.Database, questInfo.AwardItem3Id, questInfo.AwardItem3Num, tlog.IFR_DAILY_QUEST, xdEventType, "")
		if questInfo.AwardItem3Id == item_dat.ITEM_MISSION_KEY_ID {
			haveAddKey = true
		}
	}

	if questInfo.AwardItem4Id > int16(0) {
		module.Item.AddItem(state.Database, questInfo.AwardItem4Id, questInfo.AwardItem4Num, tlog.IFR_DAILY_QUEST, xdEventType, "")
		if questInfo.AwardItem4Id == item_dat.ITEM_MISSION_KEY_ID {
			haveAddKey = true
		}
	}

	// 奖励获得星数
	stars := questInfo.AwardStars
	playerStarsInfo := state.Database.Lookup.PlayerDailyQuestStarAwardInfo(state.PlayerId)
	if playerStarsInfo != nil {

		if !time.IsToday(playerStarsInfo.Lastupdatetime) {
			playerStarsInfo.Awarded = ""
			playerStarsInfo.Lastupdatetime = time.GetNowTime()
			playerStarsInfo.Stars = 0
		}
		playerStarsInfo.Stars += int32(stars)
		state.Database.Update.PlayerDailyQuestStarAwardInfo(playerStarsInfo)
	} else {
		playerStarsInfo := &mdb.PlayerDailyQuestStarAwardInfo{
			Pid:            state.PlayerId,
			Stars:          int32(stars),
			Lastupdatetime: time.GetNowTime(),
		}
		state.Database.Insert.PlayerDailyQuestStarAwardInfo(playerStarsInfo)
	}

	session.Send(&quest_api.NotifyDailyChange_Out{
		Id:          playerQuest.QuestId,
		FinishCount: playerQuest.FinishCount,
		AwardState:  dailyQuestState2QuestState(playerQuest.AwardStatus),
	})

	if haveAddKey {
		player_mission := state.Database.Lookup.PlayerMission(state.PlayerId)
		module.Notify.SendPlayerKeyChanged(session, 0, player_mission.MaxOrder)
	}
}

func getDailyInfo(session *net.Session) {
	state := module.State(session)
	rsp := &quest_api.GetDailyInfo_Out{}

	state.Database.Select.PlayerDailyQuest(func(row *mdb.PlayerDailyQuestRow) {
		rsp.Quest = append(rsp.Quest, quest_api.GetDailyInfo_Out_Quest{
			Id:          row.QuestId(),
			FinishCount: row.FinishCount(),
			AwardState:  dailyQuestState2QuestState(row.AwardStatus()),
		})
	})

	session.Send(rsp)
}

func takeQuestStarsAwaded(session *net.Session, starsLevel int32) {
	state := module.State(session)
	out := &quest_api.TakeQuestStarsAwaded_Out{}
	playerDailyQuestAwardedInfo := state.Database.Lookup.PlayerDailyQuestStarAwardInfo(state.PlayerId)
	if playerDailyQuestAwardedInfo != nil {
		if playerDailyQuestAwardedInfo.Stars < starsLevel {
			return
		}
		awardedStr := playerDailyQuestAwardedInfo.Awarded // 已领取的奖励用|联立在一起
		awardedArr := strings.Split(awardedStr, "|")
		for _, awarded := range awardedArr {
			if awarded == fmt.Sprintf("%d", starsLevel) {
				return // 已领取
			}
		}
		award := quest_dat.GetStarAwardByStarsNum(starsLevel)
		if award != nil {
			if award.Ingot > 0 {
				module.Player.IncMoney(state.Database, state.MoneyState, int64(award.Ingot), player_dat.INGOT, 0, 0, "")
			}
			if award.Coin > 0 {
				module.Player.IncMoney(state.Database, state.MoneyState, award.Coin, player_dat.COINS, 0, 0, "")
			}
			if award.Heart > 0 {
				module.Heart.IncHeart(state, int16(award.Heart))
			}
			if award.Item1 > 0 && award.Item1Num > 0 {
				module.Item.AddItem(state.Database, int16(award.Item1), int16(award.Item1Num), 0, 0, "")
			}
			if award.Item2 > 0 && award.Item2Num > 0 {
				module.Item.AddItem(state.Database, int16(award.Item2), int16(award.Item2Num), 0, 0, "")
			}
			if award.Item3 > 0 && award.Item3Num > 0 {
				module.Item.AddItem(state.Database, int16(award.Item3), int16(award.Item3Num), 0, 0, "")
			}
			if award.Item4 > 0 && award.Item4Num > 0 {
				module.Item.AddItem(state.Database, int16(award.Item4), int16(award.Item4Num), 0, 0, "")
			}
			if award.Item5 > 0 && award.Item5Num > 0 {
				module.Item.AddItem(state.Database, int16(award.Item5), int16(award.Item5Num), 0, 0, "")
			}
			//更新记录
			if playerDailyQuestAwardedInfo.Awarded == "" {
				playerDailyQuestAwardedInfo.Awarded = fmt.Sprintf("%d", starsLevel)
			} else {
				playerDailyQuestAwardedInfo.Awarded += fmt.Sprintf("|%d", starsLevel)
			}
			playerDailyQuestAwardedInfo.Lastupdatetime = time.GetNowTime()
			state.Database.Update.PlayerDailyQuestStarAwardInfo(playerDailyQuestAwardedInfo)
			out.Result = true
		}
	}
	session.Send(out)
}
