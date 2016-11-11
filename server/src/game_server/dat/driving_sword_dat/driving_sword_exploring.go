package driving_sword_dat

import (
	"core/mysql"
)

var (
	mapDrivingSwordExploring map[int16][]*ExploringAwardList
	mapExploringCount        map[int16]int
)

/*
func Load(db *mysql.Connection) {
	loadDrivingSwordExploring(db)
}
*/
type ExploringAwardList struct {
	Id              int32 //
	EventId         int8  // 探险山id
	AwardItem1      int16 // 奖励品1
	AwardItem2      int16 // 奖励品2
	AwardItem3      int16 // 奖励品3
	AwardNum1       int32 // 奖励品1数量
	AwardNum2       int32 // 奖励品2数量
	AwardNum3       int32 // 奖励品3数量
	AwardCoinNum    int32 // 奖励铜币数量
	GarrisonItem    int16 // 驻守奖励品
	GarrisonNum     int32 // 驻守奖励品数量
	GarrisonCoinNum int32 // 驻守奖励铜币数量
}

func loadDrivingSwordExploring(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM driving_sword_exploring ORDER BY `cloud_id`, `event_id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iId := res.Map("id")
	iCloudId := res.Map("cloud_id")
	iEventId := res.Map("event_id")
	iAwardItem1 := res.Map("award_item1")
	iAwardItem2 := res.Map("award_item2")
	iAwardItem3 := res.Map("award_item3")
	iAwardNum1 := res.Map("award_num1")
	iAwardNum2 := res.Map("award_num2")
	iAwardNum3 := res.Map("award_num3")
	iAwardCoinNum := res.Map("award_coin_num")
	iGarrisonItem := res.Map("garrison_item")
	iGarrisonNum := res.Map("garrison_num")
	iGarrisonCoinNum := res.Map("garrison_coin_num")

	mapDrivingSwordExploring = make(map[int16][]*ExploringAwardList, 0)
	mapExploringCount = make(map[int16]int, 0)
	for _, row := range res.Rows {
		cloud_id := row.Int16(iCloudId)
		mapDrivingSwordExploring[cloud_id] = append(mapDrivingSwordExploring[cloud_id], &ExploringAwardList{
			Id: row.Int32(iId),
			//CloudId:         row.Int32(iCloudId),
			EventId:         row.Int8(iEventId),
			AwardItem1:      row.Int16(iAwardItem1),
			AwardItem2:      row.Int16(iAwardItem2),
			AwardItem3:      row.Int16(iAwardItem3),
			AwardNum1:       row.Int32(iAwardNum1),
			AwardNum2:       row.Int32(iAwardNum2),
			AwardNum3:       row.Int32(iAwardNum3),
			AwardCoinNum:    row.Int32(iAwardCoinNum),
			GarrisonItem:    row.Int16(iGarrisonItem),
			GarrisonNum:     row.Int32(iGarrisonNum),
			GarrisonCoinNum: row.Int32(iGarrisonCoinNum),
		})
		mapExploringCount[cloud_id]++
	}
}

// ############### 对外接口实现 coding here ###############
//获取对应阶段宝箱的奖品
func GetDrivingSwordExploring(cloudID int16, eventID int8) (exploringList *ExploringAwardList) {
	if awrarlist, ok := mapDrivingSwordExploring[cloudID]; ok {
		for _, award := range awrarlist {
			if award.EventId == eventID {
				exploringList = award
				return
			}

		}
	}

	return
}

//云层探险总数
func CountExploringByCloud(cloud_id int16) int {
	return mapExploringCount[cloud_id]
}
