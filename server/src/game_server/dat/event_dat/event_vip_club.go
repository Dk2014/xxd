package event_dat

import (
	"core/fail"
	"core/mysql"
)

type VipClubAward struct {
	Require_vip_count int32
	Awards            *EventDefaultAward
}

var (
	listEventsVipClub []*VipClubAward
)

func loadEventsVipClub(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM events_vip_club_awards ORDER BY `require_vip_count` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iRequireVipCount := res.Map("require_vip_count")
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

	var pri_require_vip_count int32
	listEventsVipClub = make([]*VipClubAward, 0)

	for _, row := range res.Rows {
		pri_require_vip_count = row.Int32(iRequireVipCount)
		listEventsVipClub = append(listEventsVipClub, &VipClubAward{
			Require_vip_count: pri_require_vip_count,
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
func GetEventVipClubAward(vip_count int32) (awards *EventDefaultAward, ok bool) {
	for index := 0; index < len(listEventsVipClub); index++ {
		vip_club_awards := listEventsVipClub[index]
		if vip_club_awards.Require_vip_count == vip_count {
			awards = vip_club_awards.Awards
			ok = true
			break
		}
	}
	return
}

func GetEventVipClubAwards() []*VipClubAward {
	return listEventsVipClub
}

type EventsVipClub struct {
	RequireVipCount int32 // 需要VIP俱乐部
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

type EventsVipClubExt struct {
	StartUnixTime   int64
	EndUnixTime     int64
	DisposeUnixTime int64
	IsRelative      int8
	LTitle          string
	RTitle          string
	Content         string
	Tag             int8
	Weight          int16
	List            []*EventsVipClub
}

// VIP俱乐部运营活动数据配置
func LoadEventsVipClub(list []*EventsVipClub) {
	for _, item := range list {
		if item.RequireVipCount > 0 {
			for index := 0; index < len(listEventsVipClub); index++ {
				if listEventsVipClub[index].Require_vip_count == item.RequireVipCount {
					listEventsVipClub[index].Awards = &EventDefaultAward{
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

func GetMaxWeightInVipCount() int32 {
	return listEventsVipClub[len(listEventsVipClub)-1].Require_vip_count
}

func GetNextVipCount(now int32) (next int32) {
	for index := 0; index < len(listEventsVipClub); index++ {
		if listEventsVipClub[index].Require_vip_count > now {
			next = listEventsVipClub[index].Require_vip_count
			break
		}
	}
	return
}

func GetPlayerVipClubAward(awarded, maxAward int32) (newAwarded, nextAward int32, awards *EventDefaultAward) {
	fail.When(awarded > maxAward, "player event awarded status error")
	var index int
	for index = 0; index < len(listEventsVipClub); index++ {
		if listEventsVipClub[index].Require_vip_count > awarded {
			newAwarded = listEventsVipClub[index].Require_vip_count
			awards = listEventsVipClub[index].Awards
			break
		}
	}
	if index < len(listEventsVipClub)-1 {
		nextAward = int32(listEventsVipClub[index+1].Require_vip_count)
	}
	return
}
