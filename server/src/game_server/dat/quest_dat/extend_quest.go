package quest_dat

import (
	"core/mysql"
)

var (
	mapExtendQuest            map[int32]*ExtendQuest
	mapExtendQuestByMissionId map[int16]*ExtendQuest
	//mapExtendQuestByMainQuest map[int32][]*ExtendQuest
)

type ExtendQuest struct {
	Id               int32 // ID
	Type             int8  // 任务类型 1--通关区域评星 2--连续登录 3--元宝购买
	RelatedNpc       int32 // 关联NPC
	RequiredQuest    int16 // 前置主线任务
	RelatedMission   int16 // 关联主线区域
	RequiredProgress int16 // 要求进度
	AwardItem1       int16 // 奖励物品1
	AwardNum1        int16 // 奖励数量1
	AwardItem2       int16 // 奖励物品2
	AwardNum2        int16 // 奖励数量2
	AwardItem3       int16 // 奖励物品3
	AwardNum3        int16 // 奖励数量3
	AwardExp         int32 // 奖励经验
	AwardCoins       int32 // 奖励铜钱
	MainQuestOrder   int32 //主线任务顺序
}

func loadExtendQuest(db *mysql.Connection) {
	sql := `select eq.id as id, 
	eq.type as type, 
	eq.required_quest as required_quest, 
	eq.related_mission as related_mission,
	eq.related_npc as related_npc,
	eq.required_progress as required_progress,
	eq.award_item_1 as award_item_1,
	eq.award_num_1 as award_num_1,
	eq.award_item_2 as award_item_2,
	eq.award_num_2 as award_num_2,
	eq.award_item_3 as award_item_3,
	eq.award_num_3 as award_num_3,
	eq.award_exp as award_exp,
	eq.award_coins as award_coins,
	mq.order as main_quest_order 
	from extend_quest  eq left join quest mq 
	on eq.required_quest=mq.id
	order by mq.order;`

	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}

	iId := res.Map("id")
	iType := res.Map("type")
	iRelatedNpc := res.Map("related_npc")
	iRelatedMission := res.Map("related_mission")
	iRequiredProgress := res.Map("required_progress")
	iAwardItem1 := res.Map("award_item_1")
	iAwardNum1 := res.Map("award_num_1")
	iAwardItem2 := res.Map("award_item_2")
	iAwardNum2 := res.Map("award_num_2")
	iAwardItem3 := res.Map("award_item_3")
	iAwardNum3 := res.Map("award_num_3")
	iAwardExp := res.Map("award_exp")
	iAwardCoins := res.Map("award_coins")
	iRequiredQuest := res.Map("required_quest")
	iMainQuestOrder := res.Map("main_quest_order")

	var pri_id int32
	mapExtendQuest = map[int32]*ExtendQuest{}
	mapExtendQuestByMissionId = map[int16]*ExtendQuest{}
	var extendQuest *ExtendQuest
	for _, row := range res.Rows {
		pri_id = row.Int32(iId)
		extendQuest = &ExtendQuest{
			Id:               pri_id,
			Type:             row.Int8(iType),
			RelatedNpc:       row.Int32(iRelatedNpc),
			RequiredQuest:    row.Int16(iRequiredQuest),
			RelatedMission:   row.Int16(iRelatedMission),
			RequiredProgress: row.Int16(iRequiredProgress),
			AwardItem1:       row.Int16(iAwardItem1),
			AwardNum1:        row.Int16(iAwardNum1),
			AwardItem2:       row.Int16(iAwardItem2),
			AwardNum2:        row.Int16(iAwardNum2),
			AwardItem3:       row.Int16(iAwardItem3),
			AwardNum3:        row.Int16(iAwardNum3),
			AwardExp:         row.Int32(iAwardExp),
			AwardCoins:       row.Int32(iAwardCoins),
			MainQuestOrder:   row.Int32(iMainQuestOrder),
		}
		mapExtendQuest[pri_id] = extendQuest
		mapExtendQuestByMissionId[extendQuest.RelatedMission] = extendQuest
		//mapExtendQuestByMainQuest[pri_id] = append(mapExtendQuestByMainQuest[pri_id], extendQuest)
	}
}
