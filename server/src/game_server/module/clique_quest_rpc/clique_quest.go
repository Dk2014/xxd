package clique_quest_rpc

import (
	"core/fail"
	"core/net"
	"core/time"
	"game_server/api/protocol/clique_quest_api"
	"game_server/dat/clique_dat"
	"game_server/dat/player_dat"
	"game_server/global"
	"game_server/mdb"
	"game_server/module"
	"game_server/module/rpc"
	//"game_server/tlog"
)

func getCliqueDailyQuest(session *net.Session) {
	state := module.State(session)
	out := &clique_quest_api.GetCliqueDailyQuest_Out{}

	playerCliqueInfo := state.Database.Lookup.PlayerGlobalCliqueInfo(state.PlayerId)
	if playerCliqueInfo == nil || playerCliqueInfo.CliqueId <= 0 {
		session.Send(out)
		return
	}
	cliqueInfo := module.CliqueRPC.CliqueInfoLookUp(state.Database, playerCliqueInfo.CliqueId)
	if cliqueInfo == nil {
		session.Send(out)
		return
	}

	//检查是否需要重新加载任务
	playerInfo := global.GetPlayerInfo(state.PlayerId)
	refreshCliqueDailyQuest(state.Database, playerInfo.RoleLevel)

	state.Database.Select.PlayerGlobalCliqueDailyQuest(func(row *mdb.PlayerGlobalCliqueDailyQuestRow) {
		if row.AwardStatus() != clique_dat.CLIQUE_QUEST_STATUS_RECIVED_AWARD {
			out.Quest = append(out.Quest, clique_quest_api.GetCliqueDailyQuest_Out_Quest{
				Id:          row.QuestId(),
				FinishCount: row.FinishCount(),
				AwardState:  row.AwardStatus(),
			})
		}
	})
	session.Send(out)
	return
}

func awardCliqueDailyQuest(session *net.Session, questId int16) {
	state := module.State(session)
	out := &clique_quest_api.AwardCliqueDailyQuest_Out{}

	playerCliqueInfo := state.Database.Lookup.PlayerGlobalCliqueInfo(state.PlayerId)
	if playerCliqueInfo == nil || playerCliqueInfo.CliqueId <= 0 {
		out.Result = clique_dat.CLIQUE_AWARD_STATUS_FAILED
		session.Send(out)
		return
	}
	cliqueInfo := module.CliqueRPC.CliqueInfoLookUp(state.Database, playerCliqueInfo.CliqueId)
	if cliqueInfo == nil {
		out.Result = clique_dat.CLIQUE_AWARD_STATUS_FAILED
		session.Send(out)
		return
	}

	var playerCliqueQuest *mdb.PlayerGlobalCliqueDailyQuest

	state.Database.Select.PlayerGlobalCliqueDailyQuest(func(row *mdb.PlayerGlobalCliqueDailyQuestRow) {
		if row.QuestId() == questId {
			playerCliqueQuest = row.GoObject()
			row.Break()
		}
	})

	fail.When(playerCliqueQuest.AwardStatus != clique_dat.CLIQUE_QUEST_STATUS_CAN_AWARD, "[awardDailyQuest] not finish")

	quest := clique_dat.GetDailyCliqueQuestByQuestId(questId)
	if playerCliqueQuest.FinishCount != quest.RequireCount {
		out.Result = clique_dat.CLIQUE_AWARD_STATUS_FAILED
		session.Send(out)
		return
	}

	playerCliqueQuest.AwardStatus = clique_dat.CLIQUE_QUEST_STATUS_RECIVED_AWARD
	state.Database.Update.PlayerGlobalCliqueDailyQuest(playerCliqueQuest)

	questInfo := clique_dat.GetDailyCliqueQuestByQuestId(playerCliqueQuest.QuestId)
	award(state, &questInfo.QuestAward)
	out.Result = clique_dat.CLIQUE_AWARD_STATUS_SUCCESS
	session.Send(out)
	return
}

func award(state *module.SessionState, questAward *clique_dat.QuestAwardInfo) {
	// 增加经验
	if questAward.AwardExp > int32(0) {
		rpc.RemoteAddAwardExp(state.PlayerId, questAward.AwardExp)
	}
	//增加帮贡
	if questAward.AwardCliqueContri > int64(0) {
		module.CliqueRPC.AddPlayerCliqueContrib(state.Database, int64(questAward.AwardCliqueContri))
	}
}

func upDatePlayerCliqueDailyQuest(db *mdb.Database, class int16) {
	//检查是否可以重新加载任务
	playerInfo := global.GetPlayerInfo(db.PlayerId())
	refreshCliqueDailyQuest(db, playerInfo.RoleLevel)

	var playerQuest *mdb.PlayerGlobalCliqueDailyQuest
	db.Select.PlayerGlobalCliqueDailyQuest(func(row *mdb.PlayerGlobalCliqueDailyQuestRow) {
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
	if playerQuest.AwardStatus == clique_dat.CLIQUE_QUEST_STATUS_CAN_AWARD {
		return
	}

	quest := clique_dat.GetDailyCliqueQuestByQuestId(playerQuest.QuestId)
	playerQuest.FinishCount += 1

	if quest.RequireCount == playerQuest.FinishCount {
		playerQuest.AwardStatus = clique_dat.CLIQUE_QUEST_STATUS_CAN_AWARD
	}

	db.Update.PlayerGlobalCliqueDailyQuest(playerQuest)
	if session, ok := module.Player.GetPlayerOnline(db.PlayerId()); ok {
		session.Send(&clique_quest_api.NotifyCliqueDailyChange_Out{
			Id:          playerQuest.QuestId,
			FinishCount: playerQuest.FinishCount,
			AwardState:  playerQuest.AwardStatus,
		})
	}

	return
}

func refreshCliqueDailyQuest(db *mdb.Database, level int16) {
	//没有加入帮派，就没有必要刷新了
	playerCliqueInfo := db.Lookup.PlayerGlobalCliqueInfo(db.PlayerId())
	if playerCliqueInfo == nil || playerCliqueInfo.CliqueId <= 0 {
		return
	}
	cliqueDailyQuest := clique_dat.GetCliqueDailyQuestByLevel(int32(level))

	haveQuest := map[int16]bool{}

	var (
		delQuest    []*mdb.PlayerGlobalCliqueDailyQuest
		playerQuest *mdb.PlayerGlobalCliqueDailyQuest
	)

	//删除不是今天的每日任务;更新今天的每日任务
	db.Select.PlayerGlobalCliqueDailyQuest(func(row *mdb.PlayerGlobalCliqueDailyQuestRow) {
		playerQuest = row.GoObject()
		if !time.IsInPointHour(player_dat.RESET_DAILY_QUEST_IN_HOUR, playerQuest.LastUpdateTime) { //任务过期？
			if _, ok := cliqueDailyQuest[playerQuest.QuestId]; !ok { //没有这个任务
				delQuest = append(delQuest, playerQuest)
			} else {
				haveQuest[playerQuest.Class] = true //刷新这个任务
				playerQuest.FinishCount = 0
				playerQuest.LastUpdateTime = time.GetNowTime()
				playerQuest.AwardStatus = clique_dat.CLIQUE_QUEST_STATUS_NO_AWARD
				db.Update.PlayerGlobalCliqueDailyQuest(playerQuest)
			}
		} else {
			haveQuest[playerQuest.Class] = true
		}
	})

	for _, quest := range delQuest {
		db.Delete.PlayerGlobalCliqueDailyQuest(quest)
	}
	// 添加每日新任务
	for id, quest := range cliqueDailyQuest {
		if _, ok := haveQuest[quest.Class]; !ok {

			db.Insert.PlayerGlobalCliqueDailyQuest(&mdb.PlayerGlobalCliqueDailyQuest{
				Pid:            db.PlayerId(),
				QuestId:        id,
				FinishCount:    0,
				LastUpdateTime: time.GetNowTime(),
				AwardStatus:    clique_dat.CLIQUE_QUEST_STATUS_NO_AWARD,
				Class:          quest.Class,
			})
			//tlog.PlayerQuestFlowLog(db, int32(id), tlog.QT_DAILY, tlog.QUEST_ACCEPT)
		}
	}

	return
}
