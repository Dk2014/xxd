package clique_dat

import (
	"core/fail"
	"core/mysql"
)

var mapCliqueDailyQuest map[int16]*CliqueDailyQuest // 帮派每日任务

type QuestAwardInfo struct {
	AwardCliqueContri int64 // 奖励帮贡
	AwardExp          int32 // 奖励经验
	AwardIngot        int32 // 奖励元宝
	AwardCoins        int64 // 奖励铜钱
	AwardPhysical     int8  // 奖励体力
	AwardItem1Id      int16 // 奖励物品1
	AwardItem1Num     int16 // 奖励物品1数量
	AwardItem2Id      int16 // 奖励物品2
	AwardItem2Num     int16 // 奖励物品2数量
	AwardItem3Id      int16 // 奖励物品3
	AwardItem3Num     int16 // 奖励物品3数量
	AwardItem4Id      int16 // 奖励物品4
	AwardItem4Num     int16 // 奖励物品4数量
}

type CliqueDailyQuest struct {
	Id              int16          // 任务ID
	Name            string         // 任务标题
	Desc            string         // 简介
	RequireMinLevel int32          // 要求玩家最低等级
	RequireMaxLevel int32          // 要求玩家最高等级
	RequireCount    int16          // 需要数量
	QuestAward      QuestAwardInfo //奖励
	LevelType       int8           // 关卡类型; -1 无; 0-区域关卡;1-资源关卡;2-通天塔;8-难度关卡;9-伙伴关卡;10-灵宠关卡;11-魂侍关卡
	LevelSubType    int8           // 关卡子类型(-1--无;1--铜钱关卡;2--经验关卡)
	Class           int16          // 任务类别
	Order           int32          // 显示优先级
	//AwardIngot   int32          // 奖励元宝
}

func loadCliqueDailyQuest(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM clique_daily_quest ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iId := res.Map("id")
	iName := res.Map("name")
	iDesc := res.Map("desc")
	iRequireMinLevel := res.Map("require_min_level")
	iRequireMaxLevel := res.Map("require_max_level")
	iRequireCount := res.Map("require_count")
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
	iLevelType := res.Map("level_type")
	iLevelSubType := res.Map("level_sub_type")
	iClass := res.Map("class")
	iOrder := res.Map("order")
	iAwardIngot := res.Map("award_ingot")

	var pri_id int16
	mapCliqueDailyQuest = map[int16]*CliqueDailyQuest{}
	for _, row := range res.Rows {

		pri_id = row.Int16(iId)
		mapCliqueDailyQuest[pri_id] = &CliqueDailyQuest{
			//Id:                pri_id,
			Name:            row.Str(iName),
			Desc:            row.Str(iDesc),
			RequireMinLevel: row.Int32(iRequireMinLevel),
			RequireMaxLevel: row.Int32(iRequireMaxLevel),
			RequireCount:    row.Int16(iRequireCount),
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
			LevelType:    row.Int8(iLevelType),
			LevelSubType: row.Int8(iLevelSubType),
			Class:        row.Int16(iClass),
			Order:        row.Int32(iOrder),
		}
	}
}

func GetCliqueDailyQuestByLevel(level int32) map[int16]*CliqueDailyQuest {
	questMap := map[int16]*CliqueDailyQuest{}

	for id, quest := range mapCliqueDailyQuest {
		if level >= quest.RequireMinLevel && level <= quest.RequireMaxLevel {
			questMap[id] = quest
		}
	}

	return questMap
}

func GetDailyCliqueQuestByQuestId(id int16) *CliqueDailyQuest {
	quest, ok := mapCliqueDailyQuest[id]
	fail.When(!ok, "can't found DailyCliqueQuest")
	return quest
}
