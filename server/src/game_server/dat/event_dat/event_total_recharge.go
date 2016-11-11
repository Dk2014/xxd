package event_dat

import (
	"core/mysql"
)

var (
	listEventsTotalRechargeAwards []*EventsTotalRechargeAwards
)

type EventsTotalRechargeAwards struct {
	RequireTotal int16 // 需要额度
	Awards       *EventDefaultAward
}

func loadEventsTotalRechargeAwards(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM events_total_recharge_awards ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iRequireTotal := res.Map("require_total")
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

	listEventsTotalRechargeAwards = make([]*EventsTotalRechargeAwards, 0)
	for _, row := range res.Rows {
		listEventsTotalRechargeAwards = append(listEventsTotalRechargeAwards, &EventsTotalRechargeAwards{
			RequireTotal: row.Int16(iRequireTotal),
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

//充值返利活动
type EventTotalRechargeExt struct {
	StartUnixTime   int64
	EndUnixTime     int64
	DisposeUnixTime int64
	IsRelative      int8
	LTitle          string
	RTitle          string
	Content         string
	Tag             int8
	Weight          int16
	List            []*EventsTotalRechargeAwardsExt
}

type EventsTotalRechargeAwardsExt struct {
	RequireTotal int16 // 需要额度
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

func LoadEventsTotalRecharge(list []*EventsTotalRechargeAwardsExt) {
	for _, item := range list {
		if item.RequireTotal > 0 {
			for index := 0; index < len(listEventsTotalRechargeAwards); index++ {
				if listEventsTotalRechargeAwards[index].RequireTotal == item.RequireTotal {
					listEventsTotalRechargeAwards[index].Awards = &EventDefaultAward{
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

func GetEventTotalRechargeByOrder(order int16) *EventsTotalRechargeAwards {
	if int(order) < len(listEventsTotalRechargeAwards) {
		return listEventsTotalRechargeAwards[int(order)]
	}
	return nil
}

func CountEventTotalRecharge() int {
	return len(listEventsTotalRechargeAwards)
}

func GetEventTotalRechargeAwards() []*EventsTotalRechargeAwards {
	return listEventsTotalRechargeAwards
}

func GetEventTotalRechargeLevel(total int32) (result int16, index int) {
	index = -1
	for key, item := range listEventsTotalRechargeAwards {
		if total >= int32(item.RequireTotal) {
			result = item.RequireTotal
			index = key
		} else {
			break
		}
	}
	return
}
