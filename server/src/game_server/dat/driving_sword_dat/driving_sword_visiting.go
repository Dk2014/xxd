package driving_sword_dat

import (
	"core/mysql"
	"fmt"
)

type DrivingSwordAward struct {
	Item1    int16
	Item1Num int32
	Item2    int16
	Item2Num int32
	Item3    int16
	Item3Num int32
	Coins    int64
}

var mapDrivingSwordVistingAwards map[string]*DrivingSwordAward
var mapVisitingCount map[int16]int

func loadDrivingSwordVistingAwards(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM driving_sword_visiting"), -1)
	if err != nil {
		panic(err)
	}

	iCloudId := res.Map("cloud_id")
	iEventId := res.Map("event_id")
	iAwardItem1 := res.Map("award_item1")
	iAwardItemNum1 := res.Map("award_num1")
	iAwardItem2 := res.Map("award_item2")
	iAwardItemNum2 := res.Map("award_num2")
	iAwardItem3 := res.Map("award_item3")
	iAwardItemNum3 := res.Map("award_num3")
	iAwardCoins := res.Map("award_coin_num")

	mapDrivingSwordVistingAwards = make(map[string]*DrivingSwordAward)
	mapVisitingCount = make(map[int16]int)

	for _, row := range res.Rows {
		cloud_id := row.Int16(iCloudId)
		mapDrivingSwordVistingAwards[fmt.Sprintf("%d_%d", cloud_id, row.Int8(iEventId))] = &DrivingSwordAward{
			Item1:    row.Int16(iAwardItem1),
			Item1Num: row.Int32(iAwardItemNum1),
			Item2:    row.Int16(iAwardItem2),
			Item2Num: row.Int32(iAwardItemNum2),
			Item3:    row.Int16(iAwardItem3),
			Item3Num: row.Int32(iAwardItemNum3),
			Coins:    row.Int64(iAwardCoins),
		}
		mapVisitingCount[cloud_id]++
	}
}

func GetDrivingSwordVistingAward(cloudId int32, eventId int8) *DrivingSwordAward {
	key := fmt.Sprintf("%d_%d", cloudId, eventId)
	return mapDrivingSwordVistingAwards[key]
}

//云层拜访总数
func CountVisitingByCloud(cloud_id int16) int {
	return mapVisitingCount[cloud_id]
}
