package event_dat

import (
	"core/mysql"
)

var (
	listEventsTenDrawAwards []*EventsTenDrawAwards
)

type EventsTenDrawAwards struct {
	RequireTimes int16              // 需要次数
	Award        *EventDefaultAward // 所获奖励
}

func loadEventsTenDrawAwards(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM events_ten_draw_awards ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iRequireTimes := res.Map("require_times")
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

	listEventsTenDrawAwards = make([]*EventsTenDrawAwards, 0)
	for _, row := range res.Rows {
		listEventsTenDrawAwards = append(listEventsTenDrawAwards, &EventsTenDrawAwards{
			RequireTimes: row.Int16(iRequireTimes),
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
		})
	}
}

func IsAccessNewTenDrawAward(times int16) (exists bool) {
	max := listEventsTenDrawAwards[len(listEventsTenDrawAwards)-1].RequireTimes
	if times > max {
		return
	}
	for _, awardItem := range listEventsTenDrawAwards {
		if awardItem.RequireTimes == times {
			exists = true
			break
		} else if awardItem.RequireTimes > times {
			break
		}
	}
	return
}

func GetNextTenDraw(now int32) (next int32) {
	for index := 0; index < len(listEventsTenDrawAwards); index++ {
		if int32(listEventsTenDrawAwards[index].RequireTimes) > now {
			next = int32(listEventsTenDrawAwards[index].RequireTimes)
			break
		}
	}
	return
}

func GetEventTenDrawAward(times int32) (awards *EventDefaultAward, ok bool) {
	for index := 0; index < len(listEventsTenDrawAwards); index++ {
		ten_draw_awards := listEventsTenDrawAwards[index]
		if int32(ten_draw_awards.RequireTimes) == times {
			awards = ten_draw_awards.Award
			ok = true
			break
		}
	}
	return
}

func GetEventTenDrawAwards() []*EventsTenDrawAwards {
	return listEventsTenDrawAwards
}

func GetMaxTenDrawAwardTimes() int32 {
	return int32(listEventsTenDrawAwards[len(listEventsTenDrawAwards)-1].RequireTimes)
}

type EventsTenDraw struct {
	RequireTimes int32 // 需要十连抽次数
	Ingot        int16 // 奖励元宝
	Coin         int32 // 奖励铜钱
	Heart        int16 // 奖励爱心
	Item1Id      int16 // 物品1
	Item1Num     int16 // 物品1数量
	Item2Id      int16 // 物品2
	Item2Num     int16 // 物品2数量
	Item3Id      int16 // 物品3
	Item3Num     int16 // 物品3数量
	Item4Id      int16 // 物品4
	Item4Num     int16 // 物品4数量
	Item5Id      int16 // 物品5
	Item5Num     int16 // 物品5数量
}

type EventsTenDrawExt struct {
	StartUnixTime   int64
	EndUnixTime     int64
	DisposeUnixTime int64
	IsRelative      int8
	Type            int8
	LTitle          string
	RTitle          string
	Content         string
	Tag             int8
	Weight          int16
	List            []*EventsTenDraw
}

// 十连抽奖励活动运营活动数据配置
func LoadEventsTenDraw(list []*EventsTenDraw) {
	for _, item := range list {
		if item.RequireTimes > 0 {
			for index := 0; index < len(listEventsTenDrawAwards); index++ {
				if int32(listEventsTenDrawAwards[index].RequireTimes) == item.RequireTimes {
					listEventsTenDrawAwards[index].Award = &EventDefaultAward{
						Ingot:    item.Ingot,
						Coin:     item.Coin,
						Item1Id:  item.Item1Id,
						Item1Num: item.Item1Num,
						Item2Id:  item.Item2Id,
						Item2Num: item.Item2Num,
						Item3Id:  item.Item3Id,
						Item3Num: item.Item3Num,
						Item4Id:  item.Item4Id,
						Item4Num: item.Item4Num,
						Item5Id:  item.Item5Id,
						Item5Num: item.Item5Num,
					}
				}
			}
		}

	}
}
