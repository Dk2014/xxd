package quest_dat

import (
	"core/fail"
	"fmt"
)

func GetExtentQuestById(id int32) *ExtendQuest {
	return mapExtendQuest[id]
}

//func GetAllExtendQuest() map[int32]*ExtendQuest {
//	return mapExtendQuest
//}

func GetExtendQuestByMainQuest(mainQuestOrder int32) (quests []*ExtendQuest) {
	for _, quest := range mapExtendQuest {
		if mainQuestOrder >= quest.MainQuestOrder {
			quests = append(quests, quest)
		}
	}
	return quests
}

//区域评星伙伴招募任务
func GetExtendQuestByMissionId(missionId int16) *ExtendQuest {
	return mapExtendQuestByMissionId[missionId]
}

//获取支线任务
func GetAdditionQuestById(questId int32) (quest *AdditionQuest) {
	quest = mapAdditionQuest[questId]
	fail.When(quest == nil, fmt.Sprintf("找不到指定的支线任务 %d\n", questId))
	return quest
}

func GetAvailableAdditionQuestBySerialLock(serial_number int32, lock int16) (quests []*AdditionQuest) {
	if _, ok := mapAvailableAdditionQuest[serial_number]; ok {
		return mapAvailableAdditionQuest[serial_number][lock]
	}
	return nil
}

func GetAvailableAdditionQuestByLevel(level int16) (quests []*AdditionQuest) {
	return mapAvailableAdditionQuestByLevel[level]
}

func GetAvailableAdditionQuestByMainQuest(questId int16) (quest []*AdditionQuest) {
	return mapAvailableAdditionQuestByMainQuest[questId]
}

func GetInitAdditionQuest() []*AdditionQuest {
	return mapInitAdditionQuest
}

func GetAdditionQuestByMission(missionId int16) *AdditionQuest {
	return mapAdditionQuestByMission[missionId]
}

//获取关卡掉落类型的任务
func GetCollectionAdditionQuestByMissionLevel(levelId int32) map[int32]*AdditionQuest {
	return mapAdditionQuestByMissionLevelForQuestItem[levelId]
}

func GetAdditionQuestByMissionLevel(levelId int32) map[int32]*AdditionQuest {
	return mapAdditionQuestByMissionLevel[levelId]
}

func GetAdditionQuestByMissionEnemy(enemyGroupId int32) map[int32]*AdditionQuest {
	return mapAdditionQuestByMissionEnemy[enemyGroupId]
}

func GetAdditionQuestByRole(roleId int8) *AdditionQuest {
	return mapAdditionQuestByRole[roleId]
}

func GetMissioLevelLockByMainQuest(questId int16) int32 {
	lock, ok := mapQuestAwardMissionLevelLock[questId]
	fail.When(!ok, "找不到制定主线任务")
	return lock
}
