package clique_quest_rpc

import (
	"core/fail"
	"core/net"
	"game_server/api/protocol/clique_quest_api"
	"game_server/dat/clique_dat"
	"game_server/mdb"
	"game_server/module"
)

func getCliqueBuildingQuest(session *net.Session) {
	state := module.State(session)

	playerCliqueInfo := state.Database.Lookup.PlayerGlobalCliqueInfo(state.PlayerId)
	if playerCliqueInfo == nil || playerCliqueInfo.CliqueId <= 0 {
		session.Send(&clique_quest_api.GetCliqueBuildingQuest_Out{})
		return
	}
	cliqueInfo := module.CliqueRPC.CliqueInfoLookUp(state.Database, playerCliqueInfo.CliqueId)
	if cliqueInfo == nil {
		session.Send(&clique_quest_api.GetCliqueBuildingQuest_Out{})
		return
	}
	out := getBuildingQuest(cliqueInfo, state.Database)

	session.Send(out)

	return
}

func awardCliqueBuildingQuest(session *net.Session, questID int16) {
	state := module.State(session)
	out := &clique_quest_api.AwardCliqueBuildingQuest_Out{}
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

	var playercliqueQuest *mdb.PlayerGlobalCliqueBuildingQuest
	state.Database.Select.PlayerGlobalCliqueBuildingQuest(func(row *mdb.PlayerGlobalCliqueBuildingQuestRow) {
		if row.QuestId() == questID {
			playercliqueQuest = row.GoObject()
			row.Break()
		}
	})

	fail.When(playercliqueQuest == nil, "[awardCliqueBuildingQuest] no this questID")
	if playercliqueQuest.AwardStatus != clique_dat.CLIQUE_QUEST_STATUS_CAN_AWARD {
		out.Result = clique_dat.CLIQUE_AWARD_STATUS_FAILED
		session.Send(out)
		return
	}

	quest := clique_dat.GetCliqueBuildingQuestByQuestId(questID)

	var buildingLevel int32

	switch quest.Class {
	case clique_dat.CLIQUE_BUILDING_ZONGCI:
		buildingLevel = int32(cliqueInfo.TempleBuildingLevel)
		if buildingLevel < quest.RequireBuildingLevel {
			goto failed
		}
	case clique_dat.CLIQUE_BUILDING_HUICHUNTANG:
		buildingLevel = int32(cliqueInfo.HealthBuildingLevel)
		if buildingLevel < quest.RequireBuildingLevel {
			goto failed
		}
	case clique_dat.CLIQUE_BUILDING_SHENBINGTANG:
		buildingLevel = int32(cliqueInfo.AttackBuildingLevel)

		if buildingLevel < quest.RequireBuildingLevel {
			goto failed
		}
	case clique_dat.CLIQUE_BUILDING_JINGANGTANG:
		buildingLevel = int32(cliqueInfo.DefenseBuildingLevel)
		if buildingLevel < quest.RequireBuildingLevel {
			goto failed
		}
	case clique_dat.CLIQUE_BUILDING_ZONGDUO:
		buildingLevel = int32(cliqueInfo.CenterBuildingLevel)
		if buildingLevel < quest.RequireBuildingLevel {
			goto failed
		}
	case clique_dat.CLIQUE_BUILDING_QIANZHUANG:
		buildingLevel = int32(cliqueInfo.BankBuildingLevel)
		if buildingLevel < quest.RequireBuildingLevel {
			goto failed
		}
	default:
		goto failed
	}

	award(state, &quest.QuestAward)
	playercliqueQuest.AwardStatus = clique_dat.CLIQUE_QUEST_STATUS_RECIVED_AWARD
	state.Database.Update.PlayerGlobalCliqueBuildingQuest(playercliqueQuest)
	//是否还能领取下一个奖
	/*if quest.RequireBuildingLevel+1 <= buildingLevel {

		state.Database.Delete.PlayerGlobalCliqueBuildingQuest(playercliqueQuest)

		if cliqueQuest, ok := clique_dat.GetCliqueBuildingQuestByLevel(int32(quest.RequireBuildingLevel+1), quest.Class); ok {
			state.Database.Insert.PlayerGlobalCliqueBuildingQuest(&mdb.PlayerGlobalCliqueBuildingQuest{
				Pid:          state.Database.PlayerId(),
				QuestId:      cliqueQuest.Id,
				AwardStatus:  clique_dat.CLIQUE_QUEST_STATUS_CAN_AWARD,
				BuildingType: cliqueQuest.Class,
			})
			out.Result = clique_dat.CLIQUE_AWARD_STATUS_CONTINUE
		}
	} else {
		state.Database.Update.PlayerGlobalCliqueBuildingQuest(playercliqueQuest)
	}
	*/
	session.Send(out)
	return

failed:
	fail.When(true, "[awardCliqueBuildingQuest] quest.Class err")
}

func getBuildingQuest(cliqueInfo *mdb.GlobalClique, db *mdb.Database) *clique_quest_api.GetCliqueBuildingQuest_Out {
	result := &clique_quest_api.GetCliqueBuildingQuest_Out{}
	var playerQuest *mdb.PlayerGlobalCliqueBuildingQuest
	var haveQuest bool

	db.Select.PlayerGlobalCliqueBuildingQuest(func(row *mdb.PlayerGlobalCliqueBuildingQuestRow) {
		haveQuest = true
	})

	if !haveQuest {
		addCliqueBuildingQuest(cliqueInfo, db, db.PlayerId())
	}

	db.Select.PlayerGlobalCliqueBuildingQuest(func(row *mdb.PlayerGlobalCliqueBuildingQuestRow) {
		playerQuest = row.GoObject()
		questinfo := clique_dat.GetCliqueBuildingQuestByQuestId(playerQuest.QuestId)
		buildingLevel, _ := getBuildingLevelandCoinsByBuildingType(cliqueInfo, playerQuest.BuildingType)

		// 如果当前等级大于记录等级，则先领奖
		if questinfo.RequireBuildingLevel < int32(buildingLevel) {

			if playerQuest.AwardStatus == clique_dat.CLIQUE_QUEST_STATUS_NO_AWARD {
				playerQuest.AwardStatus = clique_dat.CLIQUE_QUEST_STATUS_CAN_AWARD
				db.Update.PlayerGlobalCliqueBuildingQuest(playerQuest)
			}

			if playerQuest.AwardStatus == clique_dat.CLIQUE_QUEST_STATUS_RECIVED_AWARD {
				if questinfo.RequireBuildingLevel+1 < clique_dat.CLIQUE_BUILDING_MAX_LEVEL {

					if cliqueQuest, ok := clique_dat.GetCliqueBuildingQuestByLevel(questinfo.RequireBuildingLevel+1, playerQuest.BuildingType); ok {
						playerQuest.QuestId = cliqueQuest.Id
						//playerQuest.AwardStatus = clique_dat.CLIQUE_QUEST_STATUS_CAN_AWARD
					}

					if questinfo.RequireBuildingLevel+1 < int32(buildingLevel) { //下一级也是可以领简奖
						playerQuest.AwardStatus = clique_dat.CLIQUE_QUEST_STATUS_CAN_AWARD
					} else { //下一还不能领取，但是状态要更新为不能领取
						playerQuest.AwardStatus = clique_dat.CLIQUE_QUEST_STATUS_NO_AWARD
					}

					db.Update.PlayerGlobalCliqueBuildingQuest(playerQuest)

				}

			}
		}
		_, conins := getBuildingLevelandCoinsByBuildingType(cliqueInfo, playerQuest.BuildingType)
		result.Quest = append(result.Quest,
			clique_quest_api.GetCliqueBuildingQuest_Out_Quest{
				Id:          playerQuest.QuestId,
				AwardState:  playerQuest.AwardStatus,
				DonateCoins: conins,
			})
	})

	return result
}

func getBuildingLevelandCoinsByBuildingType(cliqueInfo *mdb.GlobalClique, buildingType int16) (int32, int64) {
	switch buildingType {
	case clique_dat.CLIQUE_BUILDING_ZONGCI:
		return int32(cliqueInfo.TempleBuildingLevel), cliqueInfo.TempleBuildingCoins
	case clique_dat.CLIQUE_BUILDING_HUICHUNTANG:
		return int32(cliqueInfo.HealthBuildingLevel), cliqueInfo.HealthBuildingCoins
	case clique_dat.CLIQUE_BUILDING_SHENBINGTANG:
		return int32(cliqueInfo.AttackBuildingLevel), cliqueInfo.AttackBuildingCoins
	case clique_dat.CLIQUE_BUILDING_JINGANGTANG:
		return int32(cliqueInfo.DefenseBuildingLevel), cliqueInfo.DefenseBuildingCoins
	case clique_dat.CLIQUE_BUILDING_ZONGDUO:
		return int32(cliqueInfo.CenterBuildingLevel), cliqueInfo.CenterBuildingCoins
	case clique_dat.CLIQUE_BUILDING_QIANZHUANG:
		return int32(cliqueInfo.BankBuildingLevel), cliqueInfo.BankBuildingCoins
	default:
		return 0, 0
	}
}

func addCliqueBuildingQuest(cliqueInfo *mdb.GlobalClique, db *mdb.Database, pid int64) {

	var (
		templeQuest  *mdb.PlayerGlobalCliqueBuildingQuest
		healthQuest  *mdb.PlayerGlobalCliqueBuildingQuest
		attackQuest  *mdb.PlayerGlobalCliqueBuildingQuest
		defenseQuest *mdb.PlayerGlobalCliqueBuildingQuest
		centerQuest  *mdb.PlayerGlobalCliqueBuildingQuest
		bankQuest    *mdb.PlayerGlobalCliqueBuildingQuest
	)

	if quest, ok := clique_dat.GetCliqueBuildingQuestByLevel(int32(cliqueInfo.TempleBuildingLevel), clique_dat.CLIQUE_BUILDING_ZONGCI); ok {
		templeQuest = &mdb.PlayerGlobalCliqueBuildingQuest{
			QuestId:      quest.Id,
			Pid:          pid,
			BuildingType: clique_dat.CLIQUE_BUILDING_ZONGCI,
			AwardStatus:  clique_dat.CLIQUE_QUEST_STATUS_NO_AWARD,
		}
	} else if quest, ok := clique_dat.GetCliqueBuildingQuestByLevel(int32(cliqueInfo.TempleBuildingLevel-1), clique_dat.CLIQUE_BUILDING_ZONGCI); ok {
		templeQuest = &mdb.PlayerGlobalCliqueBuildingQuest{
			QuestId:      quest.Id,
			Pid:          pid,
			BuildingType: clique_dat.CLIQUE_BUILDING_ZONGCI,
			AwardStatus:  clique_dat.CLIQUE_QUEST_STATUS_RECIVED_AWARD,
		}
	}
	if quest, ok := clique_dat.GetCliqueBuildingQuestByLevel(int32(cliqueInfo.HealthBuildingLevel), clique_dat.CLIQUE_BUILDING_HUICHUNTANG); ok {

		healthQuest = &mdb.PlayerGlobalCliqueBuildingQuest{
			QuestId:      quest.Id,
			Pid:          pid,
			BuildingType: clique_dat.CLIQUE_BUILDING_HUICHUNTANG,
			AwardStatus:  clique_dat.CLIQUE_QUEST_STATUS_NO_AWARD,
		}
	} else if quest, ok := clique_dat.GetCliqueBuildingQuestByLevel(int32(cliqueInfo.HealthBuildingLevel-1), clique_dat.CLIQUE_BUILDING_HUICHUNTANG); ok {
		healthQuest = &mdb.PlayerGlobalCliqueBuildingQuest{
			QuestId:      quest.Id,
			Pid:          pid,
			BuildingType: clique_dat.CLIQUE_BUILDING_HUICHUNTANG,
			AwardStatus:  clique_dat.CLIQUE_QUEST_STATUS_RECIVED_AWARD,
		}
	}
	if quest, ok := clique_dat.GetCliqueBuildingQuestByLevel(int32(cliqueInfo.AttackBuildingLevel), clique_dat.CLIQUE_BUILDING_SHENBINGTANG); ok {

		attackQuest = &mdb.PlayerGlobalCliqueBuildingQuest{
			QuestId:      quest.Id,
			Pid:          pid,
			BuildingType: clique_dat.CLIQUE_BUILDING_SHENBINGTANG,
			AwardStatus:  clique_dat.CLIQUE_QUEST_STATUS_NO_AWARD,
		}
	} else if quest, ok := clique_dat.GetCliqueBuildingQuestByLevel(int32(cliqueInfo.AttackBuildingLevel-1), clique_dat.CLIQUE_BUILDING_SHENBINGTANG); ok {
		attackQuest = &mdb.PlayerGlobalCliqueBuildingQuest{
			QuestId:      quest.Id,
			Pid:          pid,
			BuildingType: clique_dat.CLIQUE_BUILDING_SHENBINGTANG,
			AwardStatus:  clique_dat.CLIQUE_QUEST_STATUS_RECIVED_AWARD,
		}
	}
	if quest, ok := clique_dat.GetCliqueBuildingQuestByLevel(int32(cliqueInfo.DefenseBuildingLevel), clique_dat.CLIQUE_BUILDING_JINGANGTANG); ok {

		defenseQuest = &mdb.PlayerGlobalCliqueBuildingQuest{
			QuestId:      quest.Id,
			Pid:          pid,
			BuildingType: clique_dat.CLIQUE_BUILDING_JINGANGTANG,
			AwardStatus:  clique_dat.CLIQUE_QUEST_STATUS_NO_AWARD,
		}
	} else if quest, ok := clique_dat.GetCliqueBuildingQuestByLevel(int32(cliqueInfo.DefenseBuildingLevel-1), clique_dat.CLIQUE_BUILDING_JINGANGTANG); ok {
		defenseQuest = &mdb.PlayerGlobalCliqueBuildingQuest{
			QuestId:      quest.Id,
			Pid:          pid,
			BuildingType: clique_dat.CLIQUE_BUILDING_JINGANGTANG,
			AwardStatus:  clique_dat.CLIQUE_QUEST_STATUS_RECIVED_AWARD,
		}
	}
	if quest, ok := clique_dat.GetCliqueBuildingQuestByLevel(int32(cliqueInfo.CenterBuildingLevel), clique_dat.CLIQUE_BUILDING_ZONGDUO); ok {
		centerQuest = &mdb.PlayerGlobalCliqueBuildingQuest{
			QuestId:      quest.Id,
			Pid:          pid,
			BuildingType: clique_dat.CLIQUE_BUILDING_ZONGDUO,
			AwardStatus:  clique_dat.CLIQUE_QUEST_STATUS_NO_AWARD,
		}
	} else if quest, ok := clique_dat.GetCliqueBuildingQuestByLevel(int32(cliqueInfo.CenterBuildingLevel-1), clique_dat.CLIQUE_BUILDING_ZONGDUO); ok {
		centerQuest = &mdb.PlayerGlobalCliqueBuildingQuest{
			QuestId:      quest.Id,
			Pid:          pid,
			BuildingType: clique_dat.CLIQUE_BUILDING_ZONGDUO,
			AwardStatus:  clique_dat.CLIQUE_QUEST_STATUS_RECIVED_AWARD,
		}
	}
	if quest, ok := clique_dat.GetCliqueBuildingQuestByLevel(int32(cliqueInfo.BankBuildingLevel), clique_dat.CLIQUE_BUILDING_QIANZHUANG); ok {
		bankQuest = &mdb.PlayerGlobalCliqueBuildingQuest{
			QuestId:      quest.Id,
			Pid:          pid,
			BuildingType: clique_dat.CLIQUE_BUILDING_QIANZHUANG,
			AwardStatus:  clique_dat.CLIQUE_QUEST_STATUS_NO_AWARD,
		}
	} else if quest, ok := clique_dat.GetCliqueBuildingQuestByLevel(int32(cliqueInfo.BankBuildingLevel-1), clique_dat.CLIQUE_BUILDING_QIANZHUANG); ok {
		bankQuest = &mdb.PlayerGlobalCliqueBuildingQuest{
			QuestId:      quest.Id,
			Pid:          pid,
			BuildingType: clique_dat.CLIQUE_BUILDING_QIANZHUANG,
			AwardStatus:  clique_dat.CLIQUE_QUEST_STATUS_RECIVED_AWARD,
		}
	}

	db.AgentExecute(pid, func(agentDB *mdb.Database) {
		if templeQuest != nil {
			agentDB.Insert.PlayerGlobalCliqueBuildingQuest(templeQuest)
		}
		if healthQuest != nil {
			agentDB.Insert.PlayerGlobalCliqueBuildingQuest(healthQuest)
		}
		if attackQuest != nil {
			agentDB.Insert.PlayerGlobalCliqueBuildingQuest(attackQuest)
		}
		if defenseQuest != nil {
			agentDB.Insert.PlayerGlobalCliqueBuildingQuest(defenseQuest)
		}
		if centerQuest != nil {
			agentDB.Insert.PlayerGlobalCliqueBuildingQuest(centerQuest)
		}
		if bankQuest != nil {
			agentDB.Insert.PlayerGlobalCliqueBuildingQuest(bankQuest)
		}
	})
	return
}

func cleanCliqueBuildingQuest(db *mdb.Database, pid int64) {
	db.AgentExecute(pid, func(agentDB *mdb.Database) {
		agentDB.Select.PlayerGlobalCliqueBuildingQuest(func(row *mdb.PlayerGlobalCliqueBuildingQuestRow) {
			agentDB.Delete.PlayerGlobalCliqueBuildingQuest(row.GoObject())
		})
		agentDB.Select.PlayerGlobalCliqueDailyQuest(func(row *mdb.PlayerGlobalCliqueDailyQuestRow) {
			//已完成的每日任务在推出帮派时不需要清理，避免刷任务
			if row.AwardStatus() != clique_dat.CLIQUE_QUEST_STATUS_RECIVED_AWARD {
				agentDB.Delete.PlayerGlobalCliqueDailyQuest(row.GoObject())
			}
		})
	})
}
