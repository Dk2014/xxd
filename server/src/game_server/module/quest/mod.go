package quest

import (
	"core/log"
	"core/net"
	"game_server/dat/quest_dat"
	"game_server/dat/town_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/tlog"
)

type QuestMod struct {
}

func init() {
	module.Quest = QuestMod{}
}

//获得新伙伴
func (mod QuestMod) RefreshQuestForNewBuddy(db *mdb.Database, roleId int8) {
	refreshAdditionQuestForNewBuddy(db, roleId)
}

//消灭怪物组
func (mod QuestMod) RefreshQuestForBeatEnemyGroup(db *mdb.Database, enemyId int32) {
	refreshAdditionQuestForBeatEnemyGruop(db, enemyId)
}

//通关关卡
func (mod QuestMod) RefreshQuestForPassMissionLevel(db *mdb.Database, missionLevelId, xdEventType int32) map[int16]int16 {
	refreshAdditionQuestForPassMissionLevel(db, missionLevelId)
	return awardAdditionQuestItem(db, missionLevelId, xdEventType)
}

func (mod QuestMod) RefreshQuestForMissionStarChange(db *mdb.Database, missionId int16, incStar int16) {
	if incStar <= 0 {
		return
	}
	//伙伴招募任务
	refresMissionLevelExtendQuest(db, missionId, incStar)
	//支线任务
	refreshAdditionQuestForMissionStarChange(db, missionId, incStar)
}

//支线任务进度刷新：连续登陆
func (mod QuestMod) RefreshLoginExtendQuest(state *module.SessionState, continueLogin bool) {
	var extendQuestDat *quest_dat.ExtendQuest
	state.Database.Select.PlayerExtendQuest(func(row *mdb.PlayerExtendQuestRow) {
		extendQuestDat = quest_dat.GetExtentQuestById(row.QuestId())
		if extendQuestDat.Type == quest_dat.EXTEND_QUEST_TYPE_LOGIN && row.State() == quest_dat.EXTEND_QUEST_STATUS_NONE {
			playerExtQuest := row.GoObject()
			if continueLogin {
				playerExtQuest.Progress++
				if playerExtQuest.Progress >= extendQuestDat.RequiredProgress {
					playerExtQuest.State = quest_dat.EXTEND_QUEST_STATUS_CAN_AWARD
					playerExtQuest.Progress = extendQuestDat.RequiredProgress
				}
				state.Database.Update.PlayerExtendQuest(playerExtQuest)
			} else {
				playerExtQuest.Progress = 1
				state.Database.Update.PlayerExtendQuest(playerExtQuest)
			}
		}
	})
}

func (mod QuestMod) RefreshQuestOnMainQuestChange(db *mdb.Database, mainQuestId int16) {
	newAdditionQuests := checkNewAdditionQuestForMainQuestChange(db, mainQuestId)
	notifyNewAdditionQuest(db, newAdditionQuests)
}

func (mod QuestMod) RefreshQuestOnLevelUp(db *mdb.Database, oldLv, newLv int16) {
	newAdditionQuests := checkNewAdditionQuestForLevelUp(db, oldLv, newLv)
	notifyNewAdditionQuest(db, newAdditionQuests)
}

func (mod QuestMod) RefreshQuest(state *module.SessionState, action int, xdEventType int32) {
	if state.PlayerId == 0 {
		log.Errorf("未成功登陆请求RefreshQuest接口。action(%d) \n", action)
		return
	}

	playerQuest := state.Database.Lookup.PlayerQuest(state.PlayerId)
	questInfo := quest_dat.GetQuestById(playerQuest.QuestId)

	// 任务全部做完
	if questInfo.NextQuestId == 0 && playerQuest.State == quest_dat.QUEST_STATUS_AWARD {
		return
	}

	// 存在指定了任务类型(action)情况下，如果指定的任务类型和玩家当前的任务类型不同就不处理
	if action != quest_dat.QUEST_TYPE_NONE && int(questInfo.Type) != action {
		return
	}

	var missionLevelId, enemyId int32
	if state.MissionLevelState != nil {
		enemyId = state.MissionLevelState.EnemyId
		missionLevelId = state.MissionLevelState.LevelId
	}

	oldQuestId := playerQuest.QuestId
	oldQuestState := playerQuest.State

	switch playerQuest.State {

	// 未接任务
	case quest_dat.QUEST_STATUS_NO_RECEIVE:
		mainRole := module.Role.GetMainRole(state.Database)
		if mainRole.Level >= int16(questInfo.RequireLevel) {
			playerQuest.State = quest_dat.QUEST_STATUS_ALREADY_RECEIVE
		}

	case quest_dat.QUEST_STATUS_AWARD:
		setNextQuest(state, playerQuest, questInfo)

	// 已接任务
	case quest_dat.QUEST_STATUS_ALREADY_RECEIVE:
		switch questInfo.Type {

		// 进入城镇
		case quest_dat.QUEST_TYPE_TOWN:
			if state.TownId == questInfo.TownId {
				playerQuest.State = quest_dat.QUEST_STATUS_FINISH
			}

		// 城镇npc
		case quest_dat.QUEST_TYPE_TOWN_NPC:
			if town_dat.ExistNPCIdWithTownId(state.TownId, questInfo.TownNpcId) {
				playerQuest.State = quest_dat.QUEST_STATUS_FINISH
			}

		// 消灭怪物, 由客户端表现
		case quest_dat.QUEST_TYPE_MISSION_ENEMY:
			if missionLevelId == questInfo.MissionLevelId && questInfo.EnemyId > 0 && questInfo.EnemyNum == enemyId {
				playerQuest.State = quest_dat.QUEST_STATUS_FINISH
			}

		// 进入关卡
		case quest_dat.QUEST_TYPE_MISSION:
			if missionLevelId == questInfo.MissionLevelId {
				playerQuest.State = quest_dat.QUEST_STATUS_FINISH
			}

		// 关卡通关
		case quest_dat.QUEST_TYPE_WIN_MISSION:
			if missionLevelId == questInfo.MissionLevelId {
				playerQuest.State = quest_dat.QUEST_STATUS_FINISH
			}
		}

	// 完成任务
	case quest_dat.QUEST_STATUS_FINISH:
		awardQuest(state, questInfo, xdEventType)
		tlog.PlayerQuestFlowLog(state.Database, int32(playerQuest.QuestId), tlog.QT_NORMAL, tlog.QUEST_FINISH)
		playerQuest.State = quest_dat.QUEST_STATUS_AWARD
		// 任务已奖励

		setNextQuest(state, playerQuest, questInfo)
	}

	// 只有在任务id或状态发生改变时才通知
	if oldQuestId != playerQuest.QuestId || oldQuestState != playerQuest.State {
		state.Database.Update.PlayerQuest(playerQuest)
		if oldQuestId != playerQuest.QuestId {
			//支线任务
			newAdditionQuests := checkNewAdditionQuestForMainQuestChange(state.Database, playerQuest.QuestId)
			notifyNewAdditionQuest(state.Database, newAdditionQuests)
		}
		tlog.PlayerQuestFlowLog(state.Database, int32(playerQuest.QuestId), tlog.QT_NORMAL, tlog.QUEST_ACCEPT)
		if session, ok := module.Player.GetPlayerOnline(state.PlayerId); ok {
			if questInfo.NextQuestId == quest_dat.QUEST_ID_NO_MORE && playerQuest.State == quest_dat.QUEST_STATUS_AWARD {
				module.Notify.SendQuestChange(session, quest_dat.QUEST_ID_NO_MORE, quest_dat.QUEST_STATUS_ALREADY_RECEIVE)
			} else {
				module.Notify.SendQuestChange(session, playerQuest.QuestId, playerQuest.State)
			}

		}
	}
}

func (mod QuestMod) RefreshExtendQuest(db *mdb.Database, mainQuestId int16) {
	takeExtendQuest(db, mainQuestId)
}

func (mod QuestMod) TakeAdditionQuest4Debug(session *net.Session, questId int32) {
	takeAdditionQuest(session, questId)
}
