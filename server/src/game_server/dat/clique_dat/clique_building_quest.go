package clique_dat

import (
	"core/fail"
	"core/mysql"
)

var (
	mapCliqueBuildingQuest map[int16]*CliqueBuildingQuest
)

type CliqueBuildingQuest struct {
	Id                   int16          // 任务ID
	Name                 string         // 任务标题
	Desc                 string         // 简介
	RequireBuildingLevel int32          // 要求建筑等级
	QuestAward           QuestAwardInfo //奖励
	Class                int16          // 任务类别
	Order                int32          // 显示优先级

}

func loadCliqueBuildingQuest(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM clique_building_quest ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iId := res.Map("id")
	iName := res.Map("name")
	iDesc := res.Map("desc")
	iRequireBuildingLevel := res.Map("require_building_level")
	iAwardExp := res.Map("award_exp")
	iAwardCliqueContri := res.Map("award_clique_contri")
	iAwardCoins := res.Map("award_coins")
	iAwardPhysical := res.Map("award_physical")
	iAwardItem1Id := res.Map("award_item1_id")
	iAwardItem1Num := res.Map("award_item1_num")
	iAwardItem2Id := res.Map("award_item2_id")
	iAwardItem2Num := res.Map("award_item2_num")
	iAwardItem3Id := res.Map("award_item3_id")
	iAwardItem3Num := res.Map("award_item3_num")
	iAwardItem4Id := res.Map("award_item4_id")
	iAwardItem4Num := res.Map("award_item4_num")
	iClass := res.Map("class")
	iOrder := res.Map("order")
	iAwardIngot := res.Map("award_ingot")

	var pri_id int16
	mapCliqueBuildingQuest = map[int16]*CliqueBuildingQuest{}
	for _, row := range res.Rows {
		pri_id = row.Int16(iId)
		mapCliqueBuildingQuest[pri_id] = &CliqueBuildingQuest{
			Id:                   pri_id,
			Name:                 row.Str(iName),
			Desc:                 row.Str(iDesc),
			RequireBuildingLevel: row.Int32(iRequireBuildingLevel),
			QuestAward: QuestAwardInfo{
				AwardExp:          row.Int32(iAwardExp),
				AwardCliqueContri: row.Int64(iAwardCliqueContri),
				AwardCoins:        row.Int64(iAwardCoins),
				AwardPhysical:     row.Int8(iAwardPhysical),
				AwardItem1Id:      row.Int16(iAwardItem1Id),
				AwardItem1Num:     row.Int16(iAwardItem1Num),
				AwardItem2Id:      row.Int16(iAwardItem2Id),
				AwardItem2Num:     row.Int16(iAwardItem2Num),
				AwardItem3Id:      row.Int16(iAwardItem3Id),
				AwardItem3Num:     row.Int16(iAwardItem3Num),
				AwardItem4Id:      row.Int16(iAwardItem4Id),
				AwardItem4Num:     row.Int16(iAwardItem4Num),
				AwardIngot:        row.Int32(iAwardIngot),
			},
			Class: row.Int16(iClass),
			Order: row.Int32(iOrder),
		}
	}
}

// ############### 对外接口实现 coding here ###############

func GetCliqueBuildingQuestByLevel(level int32, class int16) (*CliqueBuildingQuest, bool) {
	for _, quest := range mapCliqueBuildingQuest {
		if level == quest.RequireBuildingLevel && class == quest.Class {
			return quest, true
		}
	}

	return nil, false
}

func GetCliqueBuildingQuestByQuestId(id int16) *CliqueBuildingQuest {
	quest, ok := mapCliqueBuildingQuest[id]
	fail.When(!ok, "can't found GetCliqueBuildingQuestByQuestId")
	return quest
}
