package event_dat

import "core/mysql"

type QQVIPGiftAward struct {
	Require_login_days int32
	Awards             *EventDefaultAward
}

var (
	listEventsQQVIPGift []*QQVIPGiftAward
)

func loadEventsQQVIPGift(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM events_qqvip_gift_awards ORDER BY `require_login_days` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iRequireLoginDays := res.Map("require_login_days")
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

	var pri_require_login_days int32
	listEventsQQVIPGift = make([]*QQVIPGiftAward, 0)

	for _, row := range res.Rows {
		pri_require_login_days = row.Int32(iRequireLoginDays)
		listEventsQQVIPGift = append(listEventsQQVIPGift, &QQVIPGiftAward{
			Require_login_days: pri_require_login_days,
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
func GetEventQQVIPGiftAward(login_days int32) (days int32, awards *EventDefaultAward) {
	for index := 0; index < len(listEventsQQVIPGift); index++ {
		login_days_awards := listEventsQQVIPGift[index]
		awards = login_days_awards.Awards
		days = login_days_awards.Require_login_days
		if login_days_awards.Require_login_days == login_days {
			//查到就不往下找了
			break
		} else if login_days_awards.Require_login_days > login_days {
			//第一次超越了则返回上一个，而且不再继续找下去
			awards = listEventsQQVIPGift[index-1].Awards
			days = listEventsQQVIPGift[index-1].Require_login_days
			break
		}
	}
	return
}

type EventsQQVIPGift struct {
	RequireLoginDays int32 // 需要QQ特权赠送物品
	Ingot            int16 // 奖励元宝
	Coin             int32 // 奖励铜钱
	Item1Id          int16 // 物品1
	Item1Num         int16 // 物品1数量
	Item2Id          int16 // 物品2
	Item2Num         int16 // 物品2数量
	Item3Id          int16 // 物品3
	Item3Num         int16 // 物品3数量
	Item4Id          int16 // 物品4
	Item4Num         int16 // 物品4数量
	Item5Id          int16 // 物品5
	Item5Num         int16 // 物品5数量
}

type EventsQQVIPGiftExt struct {
	StartUnixTime   int64
	EndUnixTime     int64
	DisposeUnixTime int64
	IsRelative      int8
	LTitle          string
	RTitle          string
	Content         string
	Tag             int8
	Weight          int16
	List            []*EventsQQVIPGift
	Addition        map[string]float32
}

// QQ特权赠送物品运营活动数据配置
func LoadEventsQQVIPGift(list []*EventsQQVIPGift) {
	for _, item := range list {
		if item.RequireLoginDays > 0 {
			for index := 0; index < len(listEventsQQVIPGift); index++ {
				if listEventsQQVIPGift[index].Require_login_days == item.RequireLoginDays {
					listEventsQQVIPGift[index].Awards = &EventDefaultAward{
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

func GetMaxWeightInQQVIPGift() int32 {
	return listEventsQQVIPGift[len(listEventsQQVIPGift)-1].Require_login_days
}

func GetNextQQVIPGift(now int32) (next int32) {
	for index := 0; index < len(listEventsQQVIPGift); index++ {
		if listEventsQQVIPGift[index].Require_login_days > now {
			next = listEventsQQVIPGift[index].Require_login_days
			break
		}
	}
	return
}

func GetPlayerQQVIPGiftAward(awarded int32) (newAwarded, nextAward int32, awards *EventDefaultAward) {
	var index int
	for index = 0; index < len(listEventsQQVIPGift); index++ {
		if listEventsQQVIPGift[index].Require_login_days > awarded {
			newAwarded = listEventsQQVIPGift[index].Require_login_days
			awards = listEventsQQVIPGift[index].Awards
			break
		}
	}
	if index < len(listEventsQQVIPGift)-1 {
		nextAward = int32(listEventsQQVIPGift[index+1].Require_login_days)
	}
	return
}
