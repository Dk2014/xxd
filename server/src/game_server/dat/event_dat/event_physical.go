package event_dat

import "core/mysql"

type PhysicalAward struct {
	Require_physical int32
	Awards           *EventDefaultAward
}

var (
	listEventsPhysical []*PhysicalAward
)

func loadEventsPhysical(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM events_physical_awards ORDER BY `require_physical` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iRequirePhysical := res.Map("require_physical")
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

	var pri_require_physical int32
	listEventsPhysical = make([]*PhysicalAward, 0)

	for _, row := range res.Rows {
		pri_require_physical = row.Int32(iRequirePhysical)
		listEventsPhysical = append(listEventsPhysical, &PhysicalAward{
			Require_physical: pri_require_physical,
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
func GetEventPhysicalAward(physical int32) (awards *EventDefaultAward, ok bool) {
	for index := 0; index < len(listEventsPhysical); index++ {
		physical_awards := listEventsPhysical[index]
		if physical_awards.Require_physical == physical {
			awards = physical_awards.Awards
			ok = true
			break
		}
	}
	return
}

func GetEventPhysicalAwards() []*PhysicalAward {
	return listEventsPhysical
}

type EventsPhysical struct {
	RequirePhysical int32 // 需要活跃度活动
	Ingot           int16 // 奖励元宝
	Coin            int32 // 奖励铜钱
	Item1Id         int16 // 物品1
	Item1Num        int16 // 物品1数量
	Item2Id         int16 // 物品2
	Item2Num        int16 // 物品2数量
	Item3Id         int16 // 物品3
	Item3Num        int16 // 物品3数量
	Item4Id         int16 // 物品4
	Item4Num        int16 // 物品4数量
	Item5Id         int16 // 物品5
	Item5Num        int16 // 物品5数量
}

type EventsPhysicalExt struct {
	StartUnixTime   int64
	EndUnixTime     int64
	DisposeUnixTime int64
	IsRelative      int8
	LTitle          string
	RTitle          string
	Content         string
	Tag             int8
	Weight          int16
	List            []*EventsPhysical
}

// 活跃度活动运营活动数据配置
func LoadEventsPhysical(list []*EventsPhysical) {
	for _, item := range list {
		if item.RequirePhysical > 0 {
			for index := 0; index < len(listEventsPhysical); index++ {
				if listEventsPhysical[index].Require_physical == item.RequirePhysical {
					listEventsPhysical[index].Awards = &EventDefaultAward{
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

func GetMaxWeightInPhysical() int32 {
	return listEventsPhysical[len(listEventsPhysical)-1].Require_physical
}

func GetNextPhysical(now int32) (next int32) {
	for index := 0; index < len(listEventsPhysical); index++ {
		if listEventsPhysical[index].Require_physical > now {
			next = listEventsPhysical[index].Require_physical
			break
		}
	}
	return
}

func GetPlayerPhysicalAward(awarded int32) (newAwarded, nextAward int32, awards *EventDefaultAward) {
	var index int
	for index = 0; index < len(listEventsPhysical); index++ {
		if listEventsPhysical[index].Require_physical > awarded {
			newAwarded = listEventsPhysical[index].Require_physical
			awards = listEventsPhysical[index].Awards
			break
		}
	}
	if index < len(listEventsPhysical)-1 {
		nextAward = int32(listEventsPhysical[index+1].Require_physical)
	}
	return
}
