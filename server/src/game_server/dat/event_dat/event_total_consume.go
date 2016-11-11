package event_dat

import (
	"core/mysql"
)

var (
	EventsTotalConsumeConfig *EventsTotalConsume
)

type EventsTotalConsume struct {
	RequireCost int16 // 需要消耗
	Award       *EventDefaultAward
}

func loadEventsTotalConsume(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM events_total_consume ORDER BY `id` DESC"), -1)
	if err != nil {
		panic(err)
	}

	iRequireCost := res.Map("require_cost")
	iIngot := res.Map("ingot")
	iCoins := res.Map("coins")
	iHeart := res.Map("heart")
	iItem1Id := res.Map("item1_id")
	iItem1Num := res.Map("item1_num")
	iItem2Id := res.Map("item2_id")
	iItem2Num := res.Map("item2_num")
	iItem3Id := res.Map("item3_id")
	iItem3Num := res.Map("item3_num")
	iItem4Id := res.Map("item4_id")
	iItem4Num := res.Map("item4_num")
	iItem5Id := res.Map("item5_id")
	iItem5Num := res.Map("item5_num")

	for _, row := range res.Rows {
		EventsTotalConsumeConfig = &EventsTotalConsume{
			RequireCost: row.Int16(iRequireCost),
			Award: &EventDefaultAward{
				Ingot:    row.Int16(iIngot),
				Coin:     row.Int32(iCoins),
				Heart:    row.Int16(iHeart),
				Item1Id:  row.Int16(iItem1Id),
				Item1Num: row.Int16(iItem1Num),
				Item2Id:  row.Int16(iItem2Id),
				Item2Num: row.Int16(iItem2Num),
				Item3Id:  row.Int16(iItem3Id),
				Item3Num: row.Int16(iItem3Num),
				Item4Id:  row.Int16(iItem4Id),
				Item4Num: row.Int16(iItem4Num),
				Item5Id:  row.Int16(iItem5Id),
				Item5Num: row.Int16(iItem5Num),
			},
		}

		break
	}
}

// ############### 对外接口实现 coding here ###############

func GetTotalConsumeRadix() int16 {
	return EventsTotalConsumeConfig.RequireCost
}

func GetTotalConsumeAward() *EventDefaultAward {
	return EventsTotalConsumeConfig.Award
}

func GetTotalConsumeAwards() *EventsTotalConsume {
	return EventsTotalConsumeConfig
}

type EventsTotalConsumeExt struct {
	StartUnixTime   int64
	EndUnixTime     int64
	DisposeUnixTime int64
	IsRelative      int8
	LTitle          string
	RTitle          string
	Content         string
	Tag             int8
	Weight          int16
	List            []*EventsTotalConsumeAward
}

type EventsTotalConsumeAward struct {
	RequireCost int16
	Ingot       int16 // 奖励元宝
	Coin        int32 // 奖励铜钱
	Heart       int16 // 奖励爱心
	Item1Id     int16 // 物品1
	Item1Num    int16 // 物品1数量
	Item2Id     int16 // 物品2
	Item2Num    int16 // 物品2数量
	Item3Id     int16 // 物品3
	Item3Num    int16 // 物品3数量
	Item4Id     int16 // 物品4
	Item4Num    int16 // 物品4数量
	Item5Id     int16 // 物品5
	Item5Num    int16 // 物品5数量
}

func LoadEventTotalConsume(config []*EventsTotalConsumeAward) {
	if len(config) > 0 {
		EventsTotalConsumeConfig.RequireCost = config[0].RequireCost
		EventsTotalConsumeConfig.Award = &EventDefaultAward{
			Ingot:    config[0].Ingot,
			Coin:     config[0].Coin,
			Heart:    config[0].Heart,
			Item1Id:  config[0].Item1Id,
			Item1Num: config[0].Item1Num,
			Item2Id:  config[0].Item2Id,
			Item2Num: config[0].Item2Num,
			Item3Id:  config[0].Item3Id,
			Item3Num: config[0].Item3Num,
			Item4Id:  config[0].Item4Id,
			Item4Num: config[0].Item4Num,
			Item5Id:  config[0].Item5Id,
			Item5Num: config[0].Item5Num,
		}
	}
}
