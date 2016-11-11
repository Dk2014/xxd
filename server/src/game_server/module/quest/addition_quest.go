package quest

import (
	"core/fail"
	"core/net"
	"core/time"
	"game_server/api/protocol/quest_api"
	"game_server/dat/quest_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/tlog"
	"game_server/xdlog"
	"math/rand"
)

func dailyQuestState2QuestState(state int8) int8 {
	switch state {
	case quest_dat.DAILY_QUEST_STATUS_NONE:
		return quest_dat.ADDITION_QUEST_STATE_RECIVED
	case quest_dat.DAILY_QUEST_STATUS_CAN_AWARD:
		return quest_dat.ADDITION_QUEST_STATE_NOT_AWARD
	case quest_dat.DAILY_QUEST_STATUS_HAVE_AWARD:
		return quest_dat.ADDITION_QUEST_STATE_AWARDED
	}
	panic("undefine state")
}

//返回任务面板需要显示的任务
func getPannelQuestInfo(session *net.Session) {
	state := module.State(session)
	out := &quest_api.GetPannelQuestInfo_Out{}
	state.Database.Select.PlayerDailyQuest(func(row *mdb.PlayerDailyQuestRow) {
		out.Quest = append(out.Quest, quest_api.GetPannelQuestInfo_Out_Quest{
			QuestClass: quest_dat.QUEST_CLASS_DAILY,
			Id:         int32(row.QuestId()),
			Progress:   row.FinishCount(),
			State:      dailyQuestState2QuestState(row.AwardStatus()),
		})
	})
	state.Database.Select.PlayerAdditionQuest(func(row *mdb.PlayerAdditionQuestRow) {
		if row.State() == quest_dat.ADDITION_QUEST_STATE_RECIVED || row.State() == quest_dat.ADDITION_QUEST_STATE_NOT_AWARD {
			out.Quest = append(out.Quest, quest_api.GetPannelQuestInfo_Out_Quest{
				QuestClass: quest_dat.QUEST_CLASS_ADDITION,
				Id:         row.QuestId(),
				Progress:   row.Progress(),
				State:      row.State(),
			})
		}
	})
	// 每日任务星数奖励信息
	stars_info := state.Database.Lookup.PlayerDailyQuestStarAwardInfo(state.PlayerId)
	if stars_info != nil {
		if !time.IsToday(stars_info.Lastupdatetime) {
			stars_info.Awarded = ""
			stars_info.Lastupdatetime = time.GetNowTime()
			stars_info.Stars = 0
			state.Database.Update.PlayerDailyQuestStarAwardInfo(stars_info)
		}
		out.CurStars = stars_info.Stars
		out.Awarded = []byte(stars_info.Awarded)
	}

	session.Send(out)
}

//获取当前已领取和可领取的任务
func getAdditionQuest(session *net.Session) {
	state := module.State(session)
	out := &quest_api.GetAdditionQuest_Out{}
	var questState int8
	additionQuestCanRecive := map[int32]*quest_dat.AdditionQuest{} //quest_id -> _
	var recivedQeust []*mdb.PlayerAdditionQuest
	for _, quest := range quest_dat.GetInitAdditionQuest() {
		additionQuestCanRecive[quest.SerialNumber] = quest
	}
	state.Database.Select.PlayerAdditionQuest(func(row *mdb.PlayerAdditionQuestRow) {
		questState = row.State()
		//获取可领取的任务
		for _, quest := range quest_dat.GetAvailableAdditionQuestBySerialLock(row.SerialNumber(), row.Lock()) {
			additionQuestCanRecive[quest.SerialNumber] = quest
		}

		if questState == quest_dat.ADDITION_QUEST_STATE_NOT_RECIVE {
			return
		}
		recivedQeust = append(recivedQeust, row.GoObject())
	})

	//过滤掉 已领取的任务
	for _, quest := range recivedQeust {
		if availableQuestDat, ok := additionQuestCanRecive[quest.SerialNumber]; ok {
			currentAdditionQuestDat := quest_dat.GetAdditionQuestById(quest.QuestId)
			if currentAdditionQuestDat.RequireLock >= availableQuestDat.RequireLock {
				delete(additionQuestCanRecive, quest.SerialNumber)
			}
		}
		if quest.State == quest_dat.ADDITION_QUEST_STATE_AWARDED {
			continue
		}
		out.Quest = append(out.Quest, quest_api.GetAdditionQuest_Out_Quest{
			QuestId:  quest.QuestId,
			Progress: quest.Progress,
			State:    quest.State,
		})
	}

	mainRole := module.Role.GetMainRole(state.Database)
	playerQuest := state.Database.Lookup.PlayerQuest(state.PlayerId)
	playerQuestDat := quest_dat.GetQuestById(playerQuest.QuestId)

	for _, additionQuestDat := range additionQuestCanRecive {
		if mainRole.Level < additionQuestDat.RequireLevel {
			continue
		}
		if additionQuestDat.ShowupMainQuest > 0 {
			showupMainQuestDat := quest_dat.GetQuestById(additionQuestDat.ShowupMainQuest)
			if playerQuestDat.Order < showupMainQuestDat.Order {
				continue
			}
		}
		if additionQuestDat.DisappearMainQuest > 0 {
			disappearMainQuestDat := quest_dat.GetQuestById(additionQuestDat.DisappearMainQuest)
			if playerQuestDat.Order >= disappearMainQuestDat.Order {
				continue
			}
		}
		out.Quest = append(out.Quest, quest_api.GetAdditionQuest_Out_Quest{
			QuestId:  additionQuestDat.Id,
			Progress: 0,
			State:    quest_dat.ADDITION_QUEST_STATE_NOT_RECIVE,
		})
	}
	session.Send(out)
}

func takeAdditionQuest(session *net.Session, questId int32) {
	state := module.State(session)
	//检查玩家是否满足领取任务的条件
	additionQuestDat := quest_dat.GetAdditionQuestById(questId)
	//1. 等级
	mainRole := module.Role.GetMainRole(state.Database)
	fail.When(mainRole.Level < additionQuestDat.RequireLevel, "领取支线任务：等级不足")

	//2  主线任务状态
	playerQuest := state.Database.Lookup.PlayerQuest(state.PlayerId)
	playerQuestDat := quest_dat.GetQuestById(playerQuest.QuestId)
	if additionQuestDat.ShowupMainQuest > 0 {
		showupMainQuestDat := quest_dat.GetQuestById(additionQuestDat.ShowupMainQuest)
		fail.When(playerQuestDat.Order < showupMainQuestDat.Order, "支线任务：尚未开放")
	}
	if additionQuestDat.DisappearMainQuest > 0 {
		disappearMainQuestDat := quest_dat.GetQuestById(additionQuestDat.DisappearMainQuest)
		fail.When(playerQuestDat.Order >= disappearMainQuestDat.Order, "支线任务：已过期")
	}

	var currentAdditionQuest *mdb.PlayerAdditionQuest
	var count int = 0 //接受任务计数
	state.Database.Select.PlayerAdditionQuest(func(row *mdb.PlayerAdditionQuestRow) {
		if row.State() != quest_dat.ADDITION_QUEST_STATE_NOT_RECIVE &&
			row.State() != quest_dat.ADDITION_QUEST_STATE_AWARDED {
			//能领有限个的任务 已放弃的和已奖励的任务不统计
			count++
		}
		if row.SerialNumber() == additionQuestDat.SerialNumber {
			currentAdditionQuest = row.GoObject()
		}
		if count >= quest_dat.MAX_ADDITION_QUEST_NUM {
			row.Break()
		}
	})
	if count >= quest_dat.MAX_ADDITION_QUEST_NUM {
		session.Send(&quest_api.TakeAdditionQuest_Out{
			Success: false,
		})
		return
	}

	//3.1检查是否完成改系列支线的前置任务
	if additionQuestDat.RequireLock > 0 {
		fail.When(currentAdditionQuest.Lock < additionQuestDat.RequireLock, "支线任务：权值不满足")
	}

	//3.2避免领取已完成的任务
	if currentAdditionQuest != nil {
		if additionQuestDat.AwardLock == 0 && additionQuestDat.RequireLock == 0 { //孤立？
			fail.When(currentAdditionQuest.State != quest_dat.ADDITION_QUEST_STATE_NOT_RECIVE, "支线任务：不可领取")
		} else {
			fail.When(currentAdditionQuest.Lock >= additionQuestDat.AwardLock, "支线任务：不可领取")
		}
	}

	//4. 插入或更新任务
	//有些任务是领取后立即完成或者领取有需要立即统计进度的
	var progress int16 = 0
	var questState int8 = quest_dat.ADDITION_QUEST_STATE_RECIVED
	switch additionQuestDat.Type {
	case quest_dat.ADDITION_QUEST_TYPE_MISSION_STAR:
		progress = module.Mission.CountMissoinStar(state.Database, additionQuestDat.MissionId)
		if progress >= additionQuestDat.RequiredProgress {
			questState = quest_dat.ADDITION_QUEST_STATE_NOT_AWARD
		}
	case quest_dat.ADDITION_QUEST_TYPE_RECUIT_BUDDY:
		if module.Role.GetBuddyRole(state.Database, additionQuestDat.RoleId) != nil {
			questState = quest_dat.ADDITION_QUEST_STATE_NOT_AWARD
		}
	}

	if currentAdditionQuest != nil {
		currentAdditionQuest.QuestId = questId
		currentAdditionQuest.State = questState
		currentAdditionQuest.Progress = progress
		state.Database.Update.PlayerAdditionQuest(currentAdditionQuest)
	} else {
		currentAdditionQuest = &mdb.PlayerAdditionQuest{
			Pid:          state.PlayerId,
			SerialNumber: additionQuestDat.SerialNumber,
			QuestId:      questId,
			Progress:     progress,
			State:        questState,
		}
		state.Database.Insert.PlayerAdditionQuest(currentAdditionQuest)
	}
	session.Send(&quest_api.TakeAdditionQuest_Out{
		Success: true,
	})
	sendNotifyWithPlayerAdditionQuest(state.Database, []*mdb.PlayerAdditionQuest{currentAdditionQuest})
}

func giveUpAdditionQuest(session *net.Session, questId int32) {
	state := module.State(session)
	var additionQuest *mdb.PlayerAdditionQuest
	state.Database.Select.PlayerAdditionQuest(func(row *mdb.PlayerAdditionQuestRow) {
		if row.QuestId() == questId {
			additionQuest = row.GoObject()
			row.Break()
		}
	})
	fail.When(additionQuest == nil, "任务未领取")
	additionQuest.State = quest_dat.ADDITION_QUEST_STATE_NOT_RECIVE
	additionQuest.Progress = 0
	state.Database.Update.PlayerAdditionQuest(additionQuest)
	sendNotifyWithPlayerAdditionQuest(state.Database, []*mdb.PlayerAdditionQuest{additionQuest})
}

func takeAdditionQuestAward(session *net.Session, questId, xdEventType int32) {
	state := module.State(session)
	additionQuest := getAdditionQuestByQuestId(state.Database, questId)
	fail.When(additionQuest == nil, "未领取任务")
	additionQuestDat := quest_dat.GetAdditionQuestById(questId)
	switch additionQuestDat.Type {
	case quest_dat.ADDITION_QUEST_TYPE_CONVERSION:
		fail.When(additionQuest.State != quest_dat.ADDITION_QUEST_STATE_NOT_AWARD, "任务未完成")
	case quest_dat.ADDITION_QUEST_TYPE_KILL_ENEMY:
		fail.When(additionQuest.State != quest_dat.ADDITION_QUEST_STATE_NOT_AWARD, "任务未完成")
	case quest_dat.ADDITION_QUEST_TYPE_PASS_LEVEL:
		fail.When(additionQuest.State != quest_dat.ADDITION_QUEST_STATE_NOT_AWARD, "任务未完成")
	case quest_dat.ADDITION_QUEST_TYPE_COLLECT_ITEM:
		//删除物品
		module.Item.DelItemByItemId(state.Database, additionQuestDat.QuestItemId, additionQuestDat.RequiredProgress, tlog.IFR_ADDQUEST_NEEDED, xdEventType)
	case quest_dat.ADDITION_QUEST_TYPE_SHOW_ITEM:
		//计算物品数量
		itemNum := module.Item.GetItemNum(state.Database, additionQuestDat.RequireItemId)
		fail.When(itemNum < additionQuestDat.RequiredProgress, "任务未完成")
	case quest_dat.ADDITION_QUEST_TYPE_MISSION_STAR:
		//区域评星
		fail.When(module.Mission.CountMissoinStar(state.Database, additionQuestDat.MissionId) < additionQuestDat.RequiredProgress, "任务未完成")
	case quest_dat.ADDITION_QUEST_TYPE_RECUIT_BUDDY:
		//招募伙伴
		fail.When(module.Role.GetBuddyRole(state.Database, additionQuestDat.RoleId) == nil, "任务未完成")
	default:
		fail.When(true, "undeine addition quest type")
	}
	module.Item.Award(state, additionQuestDat, tlog.IFR_ADDQUEST_AWARD, tlog.MFR_ADDQUEST_AWARD, tlog.EFT_ADDQUEST_AWARD, xdlog.ET_ADDITION_QUEST)
	additionQuest.State = quest_dat.ADDITION_QUEST_STATE_AWARDED
	if additionQuestDat.AwardLock > additionQuest.Lock {
		additionQuest.Lock = additionQuestDat.AwardLock
	}
	state.Database.Update.PlayerAdditionQuest(additionQuest)
	sendNotifyWithPlayerAdditionQuest(state.Database, []*mdb.PlayerAdditionQuest{additionQuest})

	newAdditionQuests := checkNewAdditionQuest(state.Database, additionQuest.SerialNumber, additionQuest.Lock)
	notifyNewAdditionQuest(state.Database, newAdditionQuests)

	//支线任务奖励星数
	// stars := quest_dat.ADDITION_QUEST_AWARDE_STARS
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

func refreshAdditionQuest(session *net.Session, questId int32) {
	state := module.State(session)
	additionQuest := getAdditionQuestByQuestId(state.Database, questId)
	fail.When(additionQuest == nil, "未了领取任务")
	additionQuestDat := quest_dat.GetAdditionQuestById(questId)
	if additionQuestDat.Type == quest_dat.ADDITION_QUEST_TYPE_CONVERSION {
		additionQuest.State = quest_dat.ADDITION_QUEST_STATE_NOT_AWARD
		state.Database.Update.PlayerAdditionQuest(additionQuest)
	}
	sendNotifyWithPlayerAdditionQuest(state.Database, []*mdb.PlayerAdditionQuest{additionQuest})
}

func getAdditionQuestByQuestId(db *mdb.Database, questId int32) (quest *mdb.PlayerAdditionQuest) {
	db.Select.PlayerAdditionQuest(func(row *mdb.PlayerAdditionQuestRow) {
		if row.QuestId() == questId {
			quest = row.GoObject()
			row.Break()
		}
	})
	return quest
}

func checkNewAdditionQuestForLevelUp(db *mdb.Database, oldLevel, newLevel int16) []*quest_dat.AdditionQuest {
	var availableQuest []*quest_dat.AdditionQuest
	for oldLevel < newLevel {
		oldLevel++
		quests := quest_dat.GetAvailableAdditionQuestByLevel(oldLevel)
		if len(quests) == 0 {
			continue
		}
		quests = filterAvailableAdditionQuest(db, quests)
		availableQuest = append(availableQuest, quests...)
	}
	return availableQuest
}

func checkNewAdditionQuestForMainQuestChange(db *mdb.Database, mainQuestId int16) []*quest_dat.AdditionQuest {
	quests := quest_dat.GetAvailableAdditionQuestByMainQuest(mainQuestId)
	if len(quests) == 0 {
		return nil
	}
	quests = filterAvailableAdditionQuest(db, quests)
	return quests
}

//但支线任务权值发生变化时刷新任务
func checkNewAdditionQuest(db *mdb.Database, serialNum int32, lock int16) []*quest_dat.AdditionQuest {
	quests := quest_dat.GetAvailableAdditionQuestBySerialLock(serialNum, lock)
	if len(quests) == 0 {
		return nil
	}
	quests = filterAvailableAdditionQuest(db, quests)
	return quests
}

func getAdditionQuestLock(db *mdb.Database) (map[int32]int16, map[int32]int8) {
	questLock := map[int32]int16{}
	isoAdditionQuest := map[int32]int8{}
	db.Select.PlayerAdditionQuest(func(row *mdb.PlayerAdditionQuestRow) {
		if row.Lock() > 0 {
			questLock[row.SerialNumber()] = row.Lock()
		} else {
			questDat := quest_dat.GetAdditionQuestById(row.QuestId())
			if questDat.AwardLock == 0 && questDat.RequireLock == 0 { //孤立支线?
				if row.State() != quest_dat.ADDITION_QUEST_STATE_NOT_RECIVE {
					isoAdditionQuest[row.QuestId()] = 1
				}
			}
		}
	})
	return questLock, isoAdditionQuest
}

func filterAvailableAdditionQuest(db *mdb.Database, quests []*quest_dat.AdditionQuest) (availableQuest []*quest_dat.AdditionQuest) {
	playerQuest := db.Lookup.PlayerQuest(db.PlayerId())
	mainQuestDat := quest_dat.GetQuestById(playerQuest.QuestId)
	mainRole := module.Role.GetMainRole(db)
	questLock, isoQuest := getAdditionQuestLock(db)
	for _, quest := range quests {
		if quest.RequireLevel > mainRole.Level {
			continue
		}

		if quest.ShowupMainQuest > 0 {
			showupMainQuestDat := quest_dat.GetQuestById(quest.ShowupMainQuest)
			if mainQuestDat.Order < showupMainQuestDat.Order {
				continue
			}
		}
		if quest.DisappearMainQuest > 0 {
			disappearMainQuestDat := quest_dat.GetQuestById(quest.DisappearMainQuest)
			if mainQuestDat.Order >= disappearMainQuestDat.Order {
				continue
			}
		}
		//任务需要的支线权值不足 或 任务已完成
		if quest.RequireLock > 0 && quest.AwardLock > 0 {
			if quest.RequireLock > questLock[quest.RequireSerialNumber] ||
				quest.AwardLock <= questLock[quest.RequireSerialNumber] {
				continue
			}
		} else {
			//过滤已领取的孤立支线任务
			if _, recived := isoQuest[quest.Id]; recived {
				continue
			}
		}
		availableQuest = append(availableQuest, quest)
	}
	return
}

func notifyNewAdditionQuest(db *mdb.Database, quests []*quest_dat.AdditionQuest) {
	if len(quests) <= 0 {
		return
	}
	if session, ok := module.Player.GetPlayerOnline(db.PlayerId()); ok {
		rsp := &quest_api.GetAdditionQuest_Out{}
		for _, quest := range quests {
			rsp.Quest = append(rsp.Quest, quest_api.GetAdditionQuest_Out_Quest{
				QuestId:  quest.Id,
				Progress: 0,
				State:    quest_dat.ADDITION_QUEST_STATE_NOT_RECIVE,
			})
		}
		session.Send(rsp)
	}
}

//奖励任务掉落物品
func awardAdditionQuestItem(db *mdb.Database, missionLevelId, xdEventType int32) (awardItems map[int16]int16) {
	additionQuestDatMap := quest_dat.GetCollectionAdditionQuestByMissionLevel(missionLevelId)
	if len(additionQuestDatMap) == 0 {
		return
	}
	awardItems = map[int16]int16{}
	db.Select.PlayerAdditionQuest(func(row *mdb.PlayerAdditionQuestRow) {
		if row.State() == quest_dat.ADDITION_QUEST_STATE_RECIVED {
			if questDat, ok := additionQuestDatMap[row.QuestId()]; ok {
				if module.Item.GetItemNum(db, questDat.QuestItemId) >= questDat.RequiredProgress {
					return
				}
				chance := rand.Int31n(100) + 1
				if chance <= int32(questDat.QuestItemRate) {
					awardItems[questDat.QuestItemId] += questDat.QuestItemNum
					module.Item.AddItem(db, questDat.QuestItemId, questDat.QuestItemNum, tlog.IFR_ADDQUEST_ITEM_GET, xdEventType, "")
				}
			}

		}
	})
	return awardItems
}

func refreshAdditionQuestForPassMissionLevel(db *mdb.Database, missionLevelId int32) {
	additionQuestDatMap := quest_dat.GetAdditionQuestByMissionLevel(missionLevelId)
	if len(additionQuestDatMap) == 0 {
		return
	}
	var effectedQuest []*mdb.PlayerAdditionQuest
	db.Select.PlayerAdditionQuest(func(row *mdb.PlayerAdditionQuestRow) {
		if row.State() == quest_dat.ADDITION_QUEST_STATE_RECIVED {
			if questDat, ok := additionQuestDatMap[row.QuestId()]; ok {
				additionQuest := row.GoObject()
				additionQuest.Progress++
				if additionQuest.Progress >= questDat.RequiredProgress {
					additionQuest.Progress = questDat.RequiredProgress
					additionQuest.State = quest_dat.ADDITION_QUEST_STATE_NOT_AWARD
				}
				db.Update.PlayerAdditionQuest(additionQuest)
				effectedQuest = append(effectedQuest, additionQuest)
			}

		}
	})
	sendNotifyWithPlayerAdditionQuest(db, effectedQuest)
}

func refreshAdditionQuestForBeatEnemyGruop(db *mdb.Database, enemyGroupId int32) {
	additionQuestDatMap := quest_dat.GetAdditionQuestByMissionEnemy(enemyGroupId)
	if len(additionQuestDatMap) == 0 {
		return
	}
	var effectedQuest []*mdb.PlayerAdditionQuest
	db.Select.PlayerAdditionQuest(func(row *mdb.PlayerAdditionQuestRow) {
		if row.State() == quest_dat.ADDITION_QUEST_STATE_RECIVED {
			if questDat, ok := additionQuestDatMap[row.QuestId()]; ok {
				additionQuest := row.GoObject()
				additionQuest.Progress++
				if additionQuest.Progress >= questDat.RequiredProgress {
					additionQuest.Progress = questDat.RequiredProgress
					additionQuest.State = quest_dat.ADDITION_QUEST_STATE_NOT_AWARD
				}
				db.Update.PlayerAdditionQuest(additionQuest)
				effectedQuest = append(effectedQuest, additionQuest)
			}

		}
	})
	sendNotifyWithPlayerAdditionQuest(db, effectedQuest)
}

func refreshAdditionQuestForNewBuddy(db *mdb.Database, roleId int8) {
	additionQuestDat := quest_dat.GetAdditionQuestByRole(roleId)
	if additionQuestDat == nil {
		return
	}
	db.Select.PlayerAdditionQuest(func(row *mdb.PlayerAdditionQuestRow) {
		if row.QuestId() == additionQuestDat.Id {
			row.Break()
			additionQuest := row.GoObject()
			if additionQuest.State == quest_dat.ADDITION_QUEST_STATE_RECIVED {
				additionQuest.State = quest_dat.ADDITION_QUEST_STATE_NOT_AWARD
				db.Update.PlayerAdditionQuest(additionQuest)
				sendNotifyWithPlayerAdditionQuest(db, []*mdb.PlayerAdditionQuest{additionQuest})
			}

		}
	})
}

func refreshAdditionQuestForMissionStarChange(db *mdb.Database, missionId, progress int16) {
	fail.When(progress <= 0, "invalid progress")
	additionQuestDat := quest_dat.GetAdditionQuestByMission(missionId)
	if additionQuestDat == nil {
		return
	}
	var additionQuest *mdb.PlayerAdditionQuest
	db.Select.PlayerAdditionQuest(func(row *mdb.PlayerAdditionQuestRow) {
		if row.QuestId() == additionQuestDat.Id {
			additionQuest = row.GoObject()
			row.Break()
		}
	})
	if additionQuest == nil ||
		additionQuest.State != quest_dat.ADDITION_QUEST_STATE_RECIVED {
		return
	}
	additionQuest.Progress += progress
	if additionQuest.Progress >= additionQuestDat.RequiredProgress {
		additionQuest.Progress = additionQuestDat.RequiredProgress
		additionQuest.State = quest_dat.ADDITION_QUEST_STATE_NOT_AWARD
	}
	db.Update.PlayerAdditionQuest(additionQuest)
	if session, ok := module.Player.GetPlayerOnline(db.PlayerId()); ok {
		rsp := &quest_api.GetAdditionQuest_Out{}
		rsp.Quest = append(rsp.Quest, quest_api.GetAdditionQuest_Out_Quest{
			QuestId:  additionQuest.QuestId,
			Progress: additionQuest.Progress,
			State:    additionQuest.State,
		})
		session.Send(rsp)
	}
}

func sendNotifyWithPlayerAdditionQuest(db *mdb.Database, additionQuests []*mdb.PlayerAdditionQuest) {
	if len(additionQuests) == 0 {
		return
	}
	if session, ok := module.Player.GetPlayerOnline(db.PlayerId()); ok {
		rsp := &quest_api.GetAdditionQuest_Out{}
		for _, quest := range additionQuests {
			rsp.Quest = append(rsp.Quest, quest_api.GetAdditionQuest_Out_Quest{
				QuestId:  quest.QuestId,
				Progress: quest.Progress,
				State:    quest.State,
			})
		}
		session.Send(rsp)
	}
}
