package event_dat

import (
	"core/mysql"
)

var (
	eventsRichmanClubAwards []*EventsRichmanClubAwards
)

type EventsRichmanClubAwards struct {
	RequireVipLevel  int16 // 所需的vip等级
	RequireVipCount  int16 // 所需的vip相应人数
	AwardVipLevel1   int16 // 能领奖的vip等级1
	AwardVipItem1Id  int16 // 能领奖的vip的奖励物品1 ID
	AwardVipItem1Num int16 // 能领奖的vip的奖励物品1数量 默认为1
	AwardVipLevel2   int16 // 能领奖的vip等级2
	AwardVipItem2Id  int16 // 能领奖的vip的奖励物品2 ID
	AwardVipItem2Num int16 // 能领奖的vip的奖励物品2数量 默认为1
	AwardVipLevel3   int16 // 能领奖的vip等级3
	AwardVipItem3Id  int16 // 能领奖的vip的奖励物品3 ID
	AwardVipItem3Num int16 // 能领奖的vip的奖励物品3数量 默认为1
	AwardVipLevel4   int16 // 能领奖的vip等级4
	AwardVipItem4Id  int16 // 能领奖的vip的奖励物品4 ID
	AwardVipItem4Num int16 // 能领奖的vip的奖励物品4数量 默认为1
	AwardVipLevel5   int16 // 能领奖的vip等级5
	AwardVipItem5Id  int16 // 能领奖的vip的奖励物品5 ID
	AwardVipItem5Num int16 // 能领奖的vip的奖励物品5数量 默认为1
}

func loadEventsRichmanClubAwards(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM events_richman_club_awards ORDER BY `require_vip_level` ASC, `require_vip_count` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iRequireVipLevel := res.Map("require_vip_level")
	iRequireVipCount := res.Map("require_vip_count")
	iAwardVipLevel1 := res.Map("award_vip_level1")
	iAwardVipItem1Id := res.Map("award_vip_item1_id")
	iAwardVipItem1Num := res.Map("award_vip_item1_num")
	iAwardVipLevel2 := res.Map("award_vip_level2")
	iAwardVipItem2Id := res.Map("award_vip_item2_id")
	iAwardVipItem2Num := res.Map("award_vip_item2_num")
	iAwardVipLevel3 := res.Map("award_vip_level3")
	iAwardVipItem3Id := res.Map("award_vip_item3_id")
	iAwardVipItem3Num := res.Map("award_vip_item3_num")
	iAwardVipLevel4 := res.Map("award_vip_level4")
	iAwardVipItem4Id := res.Map("award_vip_item4_id")
	iAwardVipItem4Num := res.Map("award_vip_item4_num")
	iAwardVipLevel5 := res.Map("award_vip_level5")
	iAwardVipItem5Id := res.Map("award_vip_item5_id")
	iAwardVipItem5Num := res.Map("award_vip_item5_num")

	eventsRichmanClubAwards = make([]*EventsRichmanClubAwards, 0)
	for _, row := range res.Rows {
		eventsRichmanClubAwards = append(eventsRichmanClubAwards, &EventsRichmanClubAwards{
			RequireVipLevel:  row.Int16(iRequireVipLevel),
			RequireVipCount:  row.Int16(iRequireVipCount),
			AwardVipLevel1:   row.Int16(iAwardVipLevel1),
			AwardVipItem1Id:  row.Int16(iAwardVipItem1Id),
			AwardVipItem1Num: row.Int16(iAwardVipItem1Num),
			AwardVipLevel2:   row.Int16(iAwardVipLevel2),
			AwardVipItem2Id:  row.Int16(iAwardVipItem2Id),
			AwardVipItem2Num: row.Int16(iAwardVipItem2Num),
			AwardVipLevel3:   row.Int16(iAwardVipLevel3),
			AwardVipItem3Id:  row.Int16(iAwardVipItem3Id),
			AwardVipItem3Num: row.Int16(iAwardVipItem3Num),
			AwardVipLevel4:   row.Int16(iAwardVipLevel4),
			AwardVipItem4Id:  row.Int16(iAwardVipItem4Id),
			AwardVipItem4Num: row.Int16(iAwardVipItem4Num),
			AwardVipLevel5:   row.Int16(iAwardVipLevel5),
			AwardVipItem5Id:  row.Int16(iAwardVipItem5Id),
			AwardVipItem5Num: row.Int16(iAwardVipItem5Num),
		})
	}
}

type EventsRichmanExt struct {
	StartUnixTime   int64
	EndUnixTime     int64
	DisposeUnixTime int64
	IsRelative      int8
	LTitle          string
	RTitle          string
	Content         string
	Tag             int8
	Weight          int16
}

// ############### 对外接口实现 coding here ###############

//按顺序来获取奖励内容 顺序从1开始
func GetRichmanAwardBySeq(column int) (award *EventsRichmanClubAwards) {
	if column <= len(eventsRichmanClubAwards) {
		award = eventsRichmanClubAwards[column-1]
	}
	return
}

//根据行号  列号获取sequence  column和row都是从1开始算
func GetRichmanRealSequence(column, row int) int {
	seq := 0
	for i := 0; i < column-1; i++ {
		awardsRow := eventsRichmanClubAwards[i]
		if awardsRow.AwardVipItem1Id > 0 {
			seq++
		}
		if awardsRow.AwardVipItem2Id > 0 {
			seq++
		}
		if awardsRow.AwardVipItem3Id > 0 {
			seq++
		}
		if awardsRow.AwardVipItem4Id > 0 {
			seq++
		}
		if awardsRow.AwardVipItem5Id > 0 {
			seq++
		}
	}
	return seq + row
}
