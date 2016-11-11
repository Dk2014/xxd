package event_dat

import (
	"core/fail"
	"core/mysql"
)

type ShareAward struct {
	Require_times int32
	Awards        *EventDefaultAward
}

var (
	listEventsShare []*ShareAward
)

func loadEventsShare(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM events_share_awards ORDER BY `require_times` ASC"), -1)
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

	var pri_require_times int32
	listEventsShare = make([]*ShareAward, 0)

	for _, row := range res.Rows {
		pri_require_times = row.Int32(iRequireTimes)
		listEventsShare = append(listEventsShare, &ShareAward{
			Require_times: pri_require_times,
			Awards: &EventDefaultAward{
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

// ############### 对外接口实现 coding here ###############
func GetEventShareAward(times int32) (awards *EventDefaultAward, ok bool) {
	for index := 0; index < len(listEventsShare); index++ {
		share_awards := listEventsShare[index]
		if share_awards.Require_times == times {
			awards = share_awards.Awards
			ok = true
			break
		}
	}
	return
}

func GetEventShareAwards() []*ShareAward {
	return listEventsShare
}

type EventsShare struct {
	RequireTimes int32 // 需要分享送好礼
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

type EventsShareExt struct {
	StartUnixTime   int64
	EndUnixTime     int64
	DisposeUnixTime int64
	IsRelative      int8
	LTitle          string
	RTitle          string
	Content         string
	Tag             int8
	Weight          int16
	List            []*EventsShare
}

// 分享送好礼运营活动数据配置
func LoadEventsShare(list []*EventsShare) {
	for _, item := range list {
		if item.RequireTimes > 0 {
			for index := 0; index < len(listEventsShare); index++ {
				if listEventsShare[index].Require_times == item.RequireTimes {
					listEventsShare[index].Awards = &EventDefaultAward{
						Ingot:    item.Ingot,
						Coin:     item.Coin,
						Heart:    item.Heart,
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

func GetMaxWeightInShare() int32 {
	return listEventsShare[len(listEventsShare)-1].Require_times
}

func GetNextShare(now int32) (next int32) {
	for index := 0; index < len(listEventsShare); index++ {
		if listEventsShare[index].Require_times > now {
			next = listEventsShare[index].Require_times
			break
		}
	}
	return
}

func GetPlayerShareAward(awarded, maxAward int32) (newAwarded, nextAward int32, awards *EventDefaultAward) {
	fail.When(awarded > maxAward, "player event awarded status error")
	var index int
	for index = 0; index < len(listEventsShare); index++ {
		if listEventsShare[index].Require_times > awarded {
			newAwarded = listEventsShare[index].Require_times
			awards = listEventsShare[index].Awards
			break
		}
	}
	if index < len(listEventsShare)-1 {
		nextAward = int32(listEventsShare[index+1].Require_times)
	}
	return
}
