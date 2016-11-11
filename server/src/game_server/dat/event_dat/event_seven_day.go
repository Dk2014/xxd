package event_dat

import (
	"core/mysql"
	"game_server/dat/mail_dat"
)

type SevenDayAward struct {
	Require_day int32
	Awards      *EventDefaultAward
}

var (
	listEventsSevenDay []*SevenDayAward
)

func loadEventsSevenDay(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM events_seven_day_awards ORDER BY `require_day` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iRequireDay := res.Map("require_day")
	iIngot := res.Map("ingot")
	iCoins := res.Map("coins")
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

	var pri_require_day int32
	listEventsSevenDay = make([]*SevenDayAward, 0)

	for _, row := range res.Rows {
		pri_require_day = row.Int32(iRequireDay)
		listEventsSevenDay = append(listEventsSevenDay, &SevenDayAward{
			Require_day: pri_require_day,
			Awards: &EventDefaultAward{
				Ingot:    row.Int16(iIngot),
				Coin:     row.Int32(iCoins),
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
func GetEventSevenDayAward(day int32) (awards *EventDefaultAward, ok bool) {
	for index := 0; index < len(listEventsSevenDay); index++ {
		day_awards := listEventsSevenDay[index]
		if day_awards.Require_day == day {
			awards = day_awards.Awards
			ok = true
			break
		}
	}
	return
}

func GetEventSevenDayAwards() []*SevenDayAward {
	return listEventsSevenDay
}

type EventsSevenDay struct {
	RequireDay int32 // 需要新手七天乐
	Ingot      int16 // 奖励元宝
	Coin       int32 // 奖励铜钱
	Item1Id    int16 // 物品1
	Item1Num   int16 // 物品1数量
	Item2Id    int16 // 物品2
	Item2Num   int16 // 物品2数量
	Item3Id    int16 // 物品3
	Item3Num   int16 // 物品3数量
	Item4Id    int16 // 物品4
	Item4Num   int16 // 物品4数量
	Item5Id    int16 // 物品5
	Item5Num   int16 // 物品5数量
}

type EventsSevenDayExt struct {
	StartUnixTime   int64
	EndUnixTime     int64
	DisposeUnixTime int64
	IsRelative      int8
	LTitle          string
	RTitle          string
	Content         string
	Tag             int8
	Weight          int16
	List            []*EventsSevenDay
}

// 新手七天乐运营活动数据配置
func LoadEventsSevenDay(list []*EventsSevenDay) {
	for _, item := range list {
		if item.RequireDay > 0 {
			for index := 0; index < len(listEventsSevenDay); index++ {
				if listEventsSevenDay[index].Require_day == item.RequireDay {
					listEventsSevenDay[index].Awards = &EventDefaultAward{
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

func GetMaxWeightInSevenDay() int32 {
	return listEventsSevenDay[len(listEventsSevenDay)-1].Require_day
}

func GetNextSevenDay(now int32) (next int32) {
	for index := 0; index < len(listEventsSevenDay); index++ {
		if listEventsSevenDay[index].Require_day > now {
			next = listEventsSevenDay[index].Require_day
			break
		}
	}
	return
}

func GetPlayerSevenDayAward(awarded int32) (newAwarded, nextAward int32, awards *EventDefaultAward) {
	var index int
	for index = 0; index < len(listEventsSevenDay); index++ {
		if listEventsSevenDay[index].Require_day > awarded {
			newAwarded = listEventsSevenDay[index].Require_day
			awards = listEventsSevenDay[index].Awards
			break
		}
	}
	if index < len(listEventsSevenDay)-1 {
		nextAward = int32(listEventsSevenDay[index+1].Require_day)
	}
	return
}

func MailEventSevenLoginAward(day int32) (attachs []*mail_dat.Attachment) {
	if awards, ok := GetEventSevenDayAward(day); ok {
		attachs = makeAttachment(
			int64(awards.Coin), awards.Ingot, awards.Item1Id, awards.Item1Num, awards.Item2Id, awards.Item2Num, awards.Item3Id, awards.Item3Num, awards.Item4Id, awards.Item4Num, awards.Item5Id, awards.Item5Num)

	}
	return
}
