package quest_dat

import (
	"core/mysql"
)

var (
	mapQuestStartAward map[int32]*QuestStartAward
)

type QuestStartAward struct {
	Id       int32  // 主键
	Name     string // 奖励名称
	Stars    int32  // 所需星星数量
	Sign     string // 标识
	Ingot    int32  // 奖励元宝
	Coin     int64  // 奖励铜钱
	Heart    int32  // 奖励爱心
	Item1    int32  // 奖励物品1
	Item1Num int32  // 奖励物品1数量
	Item2    int32  // 奖励物品2
	Item2Num int32  // 奖励物品2数量
	Item3    int32  // 奖励物品3
	Item3Num int32  // 奖励物品3数量
	Item4    int32  // 奖励物品4
	Item4Num int32  // 奖励物品4数量
	Item5    int32  // 奖励物品5
	Item5Num int32  // 奖励物品5数量
}

func loadQuestStartAward(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM quest_start_award ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iName := res.Map("name")
	iStars := res.Map("stars")
	iSign := res.Map("sign")
	iIngot := res.Map("ingot")
	iCoin := res.Map("coin")
	iHeart := res.Map("heart")
	iItem1 := res.Map("item1")
	iItem1Num := res.Map("item1_num")
	iItem2 := res.Map("item2")
	iItem2Num := res.Map("item2_num")
	iItem3 := res.Map("item3")
	iItem3Num := res.Map("item3_num")
	iItem4 := res.Map("item4")
	iItem4Num := res.Map("item4_num")
	iItem5 := res.Map("item5")
	iItem5Num := res.Map("item5_num")

	var pri_id int32
	mapQuestStartAward = map[int32]*QuestStartAward{}
	for _, row := range res.Rows {
		pri_id = row.Int32(iStars)
		mapQuestStartAward[pri_id] = &QuestStartAward{
			Name:     row.Str(iName),
			Stars:    pri_id,
			Sign:     row.Str(iSign),
			Ingot:    row.Int32(iIngot),
			Coin:     row.Int64(iCoin),
			Heart:    row.Int32(iHeart),
			Item1:    row.Int32(iItem1),
			Item1Num: row.Int32(iItem1Num),
			Item2:    row.Int32(iItem2),
			Item2Num: row.Int32(iItem2Num),
			Item3:    row.Int32(iItem3),
			Item3Num: row.Int32(iItem3Num),
			Item4:    row.Int32(iItem4),
			Item4Num: row.Int32(iItem4Num),
			Item5:    row.Int32(iItem5),
			Item5Num: row.Int32(iItem5Num),
		}
	}
}

// ############### 对外接口实现 coding here ###############
func GetStarAwardByStarsNum(stars int32) *QuestStartAward {
	var result *QuestStartAward
	if award, ok := mapQuestStartAward[stars]; ok {
		result = award
	}
	return result
}
