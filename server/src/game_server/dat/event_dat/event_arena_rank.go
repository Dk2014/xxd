package event_dat

import (
	"core/fail"
	"core/mysql"
	"game_server/dat/mail_dat"
)

var (
	mapEventsArenaRank map[int32]*EventDefaultAward
)

func loadEventsArenaRank(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM events_arena_rank_awards ORDER BY `require_arena_rank` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iRequireArenaRank := res.Map("require_arena_rank")
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

	var pri_require_arena_rank int32
	mapEventsArenaRank = map[int32]*EventDefaultAward{}

	for _, row := range res.Rows {
		pri_require_arena_rank = row.Int32(iRequireArenaRank)
		mapEventsArenaRank[pri_require_arena_rank] = &EventDefaultAward{
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
		}

	}
}

// ############### 对外接口实现 coding here ###############
func GetEventArenaRankAward(arena_rank int32) (awards *EventDefaultAward, ok bool) {
	awards, ok = mapEventsArenaRank[arena_rank]
	return
}

func GetEventArenaRankAwards() map[int32]*EventDefaultAward {
	return mapEventsArenaRank
}

type EventsArenaRank struct {
	RequireArenaRank int32 // 需要比武场排名活动
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

type EventsArenaRankExt struct {
	StartUnixTime   int64
	EndUnixTime     int64
	DisposeUnixTime int64
	IsRelative      int8
	LTitle          string
	RTitle          string
	Content         string
	Tag             int8
	Weight          int16
	List            []*EventsArenaRank
}

// 比武场排名活动运营活动数据配置
func LoadEventsArenaRank(list []*EventsArenaRank) {
	for _, item := range list {
		mapEventsArenaRank[item.RequireArenaRank] = &EventDefaultAward{
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

func GetMaxWeightInArenaRank() int32 {
	var max int32 = 0
	for arena_rank, _ := range mapEventsArenaRank {
		if arena_rank >= max {
			max = arena_rank
		}
	}
	return max
}

func GetNextArenaRank(now int32) (next int32) {
	for arena_rank, _ := range mapEventsArenaRank {
		if arena_rank <= now {
			continue
		}
		if next == 0 {
			next = arena_rank
		} else if next > arena_rank {
			next = arena_rank
		}
	}
	return
}

func GetPlayerArenaRankAward(awarded, maxAward int32) (newAwarded, nextAward int32, awards *EventDefaultAward) {
	fail.When(awarded > maxAward, "player event awarded status error")
	newAwarded = GetNextArenaRank(awarded)
	if newAwarded <= maxAward {
		awards = mapEventsArenaRank[newAwarded]
		nextAward = GetNextArenaRank(newAwarded)
	}
	return
}

//根据比武场排名获取比武场奖励名次
func GetAwardRankByRank(rank int32) int32 {
	var min int32 = 0
	for arena_rank, _ := range mapEventsArenaRank {
		if rank > arena_rank {
			continue
		} else {
			if min == 0 {
				min = arena_rank
			} else {
				if min > arena_rank {
					min = arena_rank
				}
			}
		}
	}
	return min
}

func MakeArenaRankAwardAttachment(award *EventDefaultAward) (attachs []*mail_dat.Attachment) {
	attachs = makeAttachment(int64(award.Coin), award.Ingot, award.Item1Id, award.Item1Num, award.Item2Id, award.Item2Num, award.Item3Id, award.Item3Num, award.Item4Id, award.Item4Num, award.Item5Id, award.Item5Num)
	return
}
