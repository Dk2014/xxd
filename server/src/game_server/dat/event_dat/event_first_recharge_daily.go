package event_dat

import (
	"core/mysql"
	"sort"
)

var (
	listEventsFirstRechargeDaily []*EventsFirstRechargeDaily
)

type byFirstRechargeDaily []*EventsFirstRechargeDaily

func (f byFirstRechargeDaily) Len() int           { return len(f) }
func (f byFirstRechargeDaily) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
func (f byFirstRechargeDaily) Less(i, j int) bool { return f[i].RequireDay < f[j].RequireDay }

/*
func Load(db *mysql.Connection) {
	loadEventsFirstRechargeDaily(db)
}
*/
type EventsFirstRechargeDaily struct {
	RequireDay int16 // 索引天数
	EventAward *EventDefaultAward
}

type EventsFRDAward struct {
	RequireDay int16
	Ingot      int16 // 奖励元宝
	Coin       int32 // 奖励铜钱
	Heart      int16 // 奖励爱心
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

type EventsFirstRechargeDailyExt struct {
	StartUnixTime   int64
	EndUnixTime     int64
	DisposeUnixTime int64
	IsRelative      int8
	LTitle          string
	RTitle          string
	Content         string
	Tag             int8
	Weight          int16
	List            []*EventsFRDAward
}

func loadEventsFirstRechargeDaily(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM event_first_recharge_daily ORDER BY `require_day` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iRequireDay := res.Map("require_day")
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

	//var pri_id int32
	listEventsFirstRechargeDaily = make([]*EventsFirstRechargeDaily, 0)
	for _, row := range res.Rows {

		listEventsFirstRechargeDaily = append(listEventsFirstRechargeDaily, &EventsFirstRechargeDaily{
			RequireDay: row.Int16(iRequireDay),
			EventAward: &EventDefaultAward{
				Ingot:    row.Int16(iIngot),
				Coin:     int32(row.Int16(iCoins)),
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
func LoadEventsFirstRechargeDaily(list []*EventsFRDAward) {

	var isAdd bool
	for _, item := range list {

		if item.RequireDay > 0 {
			isAdd = true
			for index := 0; index < len(listEventsFirstRechargeDaily); index++ {

				if listEventsFirstRechargeDaily[index].RequireDay == item.RequireDay {
					listEventsFirstRechargeDaily[index] = &EventsFirstRechargeDaily{
						RequireDay: item.RequireDay,
						EventAward: &EventDefaultAward{
							Ingot:    item.Ingot,
							Coin:     int32(item.Coin),
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
						},
					}

					isAdd = false
				}
			}

			if isAdd {
				listEventsFirstRechargeDaily = append(listEventsFirstRechargeDaily, &EventsFirstRechargeDaily{
					RequireDay: item.RequireDay,
					EventAward: &EventDefaultAward{
						Ingot:    item.Ingot,
						Coin:     int32(item.Coin),
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
					},
				})
			}
		}

	}

	//清除掉空数据
	tempList := make([]*EventsFirstRechargeDaily, 0)
	for _, v := range listEventsFirstRechargeDaily {
		if IsVaildEventDefaultAward(v.EventAward) {
			tempList = append(tempList, v)
		}
	}

	if tempList != nil {
		sort.Sort(byFirstRechargeDaily(tempList))
	}

	listEventsFirstRechargeDaily = tempList
}

func FirstRechargeDailyReward(day int) (e *EventsFirstRechargeDaily) {
	if (day > len(listEventsFirstRechargeDaily)) || (day <= 0) {
		return
	}

	for _, v := range listEventsFirstRechargeDaily {
		if day == int(v.RequireDay) {
			e = v
		}
	}

	return
}

func GetFirstRechargeDailyRewards() []*EventsFirstRechargeDaily {
	return listEventsFirstRechargeDaily
}
