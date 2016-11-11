package quest_dat

import (
	"core/mysql"
)

var (
	mapAdditionQuest                           map[int32]*AdditionQuest
	mapAvailableAdditionQuest                  map[int32]map[int16][]*AdditionQuest // (serial_id, lock) -> []quest
	mapInitAdditionQuest                       []*AdditionQuest                     // 没有前置支线任务的支线任务
	mapAvailableAdditionQuestByLevel           map[int16][]*AdditionQuest           //level -> []quest
	mapAvailableAdditionQuestByMainQuest       map[int16][]*AdditionQuest           //main_quest_id -> []quest
	mapAdditionQuestByMission                  map[int16]*AdditionQuest
	mapAdditionQuestByRole                     map[int8]*AdditionQuest
	mapAdditionQuestByMissionLevel             map[int32]map[int32]*AdditionQuest
	mapAdditionQuestByMissionLevelForQuestItem map[int32]map[int32]*AdditionQuest //level -> []quest
	mapAdditionQuestByMissionEnemy             map[int32]map[int32]*AdditionQuest
)

type AdditionQuest struct {
	Id           int32 // ID
	SerialNumber int32 // 任务系列
	//Name               string // 任务名称
	//Description        string // 任务描述
	ShowupMainQuest         int16  // 出现主线任务
	DisappearMainQuest      int16  // 消失主线任务
	RequireLock             int16  // 要求支线权值
	AwardLock               int16  // 奖励支线权值
	PublishNpc              int32  // 任务发布NPC
	Type                    int8   // 任务类型 1-NPC对话 2-消灭敌人 3-通关关卡 4-收集物品 5-展示物品
	NpcId                   int32  // 对话类任务NPC
	MissionLevelId          int32  // 关联区域关卡ID
	MissionEnemyId          int32  // 关联关卡怪物组ID
	QuestItemId             int16  // 掉落物品ID（任务关卡掉落）
	QuestItemNum            int16  // 掉落物品数量
	QuestItemRate           int16  // 掉落物品
	RequireItemType         int8   // 需求物品类型
	RequireItemId           int16  // 需求物品ID
	RequiredProgress        int16  // 要求进度（物品数量通关次数等）
	AwardItem1              int16  // 奖励物品1
	AwardNum1               int16  // 奖励数量1
	AwardItem2              int16  // 奖励物品2
	AwardNum2               int16  // 奖励数量2
	AwardEquip1             int16  // 奖励装备1
	AwardEquipNum1          int16  // 奖励装备1
	AwardExp                int32  // 奖励经验
	AwardCoins              int32  // 奖励铜钱
	ConversionRecivingQuest string // 领取任务对话
	ConversionRecivedQuest  string // 任务中对话
	ConversionFinishQuest   string // 完成任务对话
	EnemyId                 int16  // 击杀怪物
	RequireSerialNumber     int32  // 前置任务系列
	RequireLevel            int16  // 要求等级
	RoleId                  int8   // 任务伙伴ID
	MissionId               int16  // 任务区域关卡ID

}

func (quest *AdditionQuest) Coins() int32 {
	return quest.AwardCoins
}

func (quest *AdditionQuest) Ingot() int32 {
	return 0
}

func (quest *AdditionQuest) Item() (award map[int16]int16) {
	award = map[int16]int16{}
	if quest.AwardNum1 > 0 {
		award[quest.AwardItem1] = quest.AwardNum1
	}
	if quest.AwardNum2 > 0 {
		award[quest.AwardItem2] = quest.AwardNum2
	}
	if quest.AwardEquipNum1 > 0 {
		award[quest.AwardEquip1] = quest.AwardEquipNum1
	}
	return award
}

func (quest *AdditionQuest) Exp() int32 {
	return quest.AwardExp
}

func loadAdditionQuest(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM addition_quest ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iId := res.Map("id")
	iSerialNumber := res.Map("serial_number")
	//iName := res.Map("name")
	iShowupMainQuest := res.Map("showup_main_quest")
	iDisappearMainQuest := res.Map("disappear_main_quest")
	iRequireLock := res.Map("require_lock")
	iAwardLock := res.Map("award_lock")
	iPublishNpc := res.Map("publish_npc")
	iType := res.Map("type")
	iNpcId := res.Map("npc_id")
	iMissionLevelId := res.Map("mission_level_id")
	iMissionEnemyId := res.Map("mission_enemy_id")
	iQuestItemId := res.Map("quest_item_id")
	iQuestItemNum := res.Map("quest_item_num")
	iQuestItemRate := res.Map("quest_item_rate")
	iRequireItemType := res.Map("require_item_type")
	iRequireItemId := res.Map("require_item_id")
	iRequiredProgress := res.Map("required_progress")
	iAwardItem1 := res.Map("award_item_1")
	iAwardNum1 := res.Map("award_num_1")
	iAwardItem2 := res.Map("award_item_2")
	iAwardNum2 := res.Map("award_num_2")
	iAwardEquip1 := res.Map("award_equip_1")
	iAwardEquipNum1 := res.Map("award_equip_num_1")
	iAwardExp := res.Map("award_exp")
	iAwardCoins := res.Map("award_coins")
	iRequireSerialNumber := res.Map("require_serial_number")
	iRequireLevel := res.Map("require_level")
	iRoleId := res.Map("role_id")
	iMissionId := res.Map("mission_id")

	var pri_id, serial_number int32
	var lock int16
	mapAvailableAdditionQuest = map[int32]map[int16][]*AdditionQuest{}
	mapAvailableAdditionQuestByMainQuest = map[int16][]*AdditionQuest{}
	mapAvailableAdditionQuestByLevel = map[int16][]*AdditionQuest{}

	mapAdditionQuest = map[int32]*AdditionQuest{}
	mapAdditionQuestByMission = map[int16]*AdditionQuest{}
	mapAdditionQuestByRole = map[int8]*AdditionQuest{}
	mapAdditionQuestByMissionLevel = map[int32]map[int32]*AdditionQuest{}
	mapAdditionQuestByMissionLevelForQuestItem = map[int32]map[int32]*AdditionQuest{}
	mapAdditionQuestByMissionEnemy = map[int32]map[int32]*AdditionQuest{}
	mapInitAdditionQuest = []*AdditionQuest{}

	for _, row := range res.Rows {
		pri_id = row.Int32(iId)
		serial_number = row.Int32(iRequireSerialNumber)
		lock = row.Int16(iRequireLock)
		quest := &AdditionQuest{
			Id:                  pri_id,
			SerialNumber:        row.Int32(iSerialNumber),
			ShowupMainQuest:     row.Int16(iShowupMainQuest),
			DisappearMainQuest:  row.Int16(iDisappearMainQuest),
			RequireLock:         row.Int16(iRequireLock),
			AwardLock:           row.Int16(iAwardLock),
			Type:                row.Int8(iType),
			NpcId:               row.Int32(iNpcId),
			MissionLevelId:      row.Int32(iMissionLevelId),
			PublishNpc:          row.Int32(iPublishNpc),
			MissionEnemyId:      row.Int32(iMissionEnemyId),
			QuestItemId:         row.Int16(iQuestItemId),
			QuestItemNum:        row.Int16(iQuestItemNum),
			QuestItemRate:       row.Int16(iQuestItemRate),
			RequireItemType:     row.Int8(iRequireItemType),
			RequireItemId:       row.Int16(iRequireItemId),
			RequiredProgress:    row.Int16(iRequiredProgress),
			AwardItem1:          row.Int16(iAwardItem1),
			AwardNum1:           row.Int16(iAwardNum1),
			AwardItem2:          row.Int16(iAwardItem2),
			AwardNum2:           row.Int16(iAwardNum2),
			AwardEquip1:         row.Int16(iAwardEquip1),
			AwardEquipNum1:      row.Int16(iAwardEquipNum1),
			AwardExp:            row.Int32(iAwardExp),
			AwardCoins:          row.Int32(iAwardCoins),
			RequireSerialNumber: row.Int32(iRequireSerialNumber),
			RequireLevel:        row.Int16(iRequireLevel),
			RoleId:              row.Int8(iRoleId),
			MissionId:           row.Int16(iMissionId),
		}
		mapAdditionQuest[pri_id] = quest
		if lock > 0 {
			if mapAvailableAdditionQuest[serial_number] == nil {
				mapAvailableAdditionQuest[serial_number] = map[int16][]*AdditionQuest{}
			}
			mapAvailableAdditionQuest[serial_number][lock] = append(mapAvailableAdditionQuest[serial_number][lock], quest)
		} else {
			mapInitAdditionQuest = append(mapInitAdditionQuest, quest)
		}
		if quest.ShowupMainQuest > 0 {
			mapAvailableAdditionQuestByMainQuest[quest.ShowupMainQuest] = append(mapAvailableAdditionQuestByMainQuest[quest.ShowupMainQuest], quest)
		}
		if quest.RequireLevel > 0 {
			mapAvailableAdditionQuestByLevel[quest.RequireLevel] = append(mapAvailableAdditionQuestByLevel[quest.RequireLevel], quest)
		}
		switch quest.Type {
		case ADDITION_QUEST_TYPE_MISSION_STAR:
			mapAdditionQuestByMission[quest.MissionId] = quest
		case ADDITION_QUEST_TYPE_RECUIT_BUDDY:
			mapAdditionQuestByRole[quest.RoleId] = quest
		case ADDITION_QUEST_TYPE_KILL_ENEMY:
			if mapAdditionQuestByMissionEnemy[quest.MissionEnemyId] == nil {
				mapAdditionQuestByMissionEnemy[quest.MissionEnemyId] = map[int32]*AdditionQuest{}
			}
			mapAdditionQuestByMissionEnemy[quest.MissionEnemyId][quest.Id] = quest
		case ADDITION_QUEST_TYPE_COLLECT_ITEM:
			if mapAdditionQuestByMissionLevelForQuestItem[quest.MissionLevelId] == nil {
				mapAdditionQuestByMissionLevelForQuestItem[quest.MissionLevelId] = map[int32]*AdditionQuest{}
			}
			mapAdditionQuestByMissionLevelForQuestItem[quest.MissionLevelId][quest.Id] = quest
		case ADDITION_QUEST_TYPE_PASS_LEVEL:
			if mapAdditionQuestByMissionLevel[quest.MissionLevelId] == nil {
				mapAdditionQuestByMissionLevel[quest.MissionLevelId] = map[int32]*AdditionQuest{}
			}
			mapAdditionQuestByMissionLevel[quest.MissionLevelId][quest.Id] = quest
		}
	}
}
