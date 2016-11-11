package driving_sword_dat

import (
	"core/mysql"
)

var (
	mapDrivingSwordTreasureContent map[ /*cloud_id*/ int16]map[ /*event_id*/ int8][]*DrivingSwordTreasureContent
	mapTreasureCount               map[int16]int
)

type DrivingSwordTreasureContent struct {
	//Id         int32 //
	//TreasureId int32 // 云海宝箱id
	Order         int8  // 奖励顺序
	AwardItem     int16 // 奖励品
	AwardNum      int32 // 奖励品数量
	AwardCoinsNum int32 // 奖励品铜币数量
}

func loadDrivingSwordTreasureContent(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM `driving_sword_treasure_content` LEFT JOIN `driving_sword_treasure` ON `driving_sword_treasure_content`.`treasure_id` = `driving_sword_treasure`.`id` ORDER BY `driving_sword_treasure`.`cloud_id`,`driving_sword_treasure`.`event_id`,`order` ASC;"), -1)
	if err != nil {
		panic(err)
	}

	iCloudId := res.Map("cloud_id")
	iEventId := res.Map("event_id")
	iOrder := res.Map("order")
	iAwardItem := res.Map("award_item")
	iAwardNum := res.Map("award_num")
	iAwardCoinsNum := res.Map("award_coins")

	mapDrivingSwordTreasureContent = map[int16]map[int8][]*DrivingSwordTreasureContent{}

	for _, row := range res.Rows {
		cloud_id := row.Int16(iCloudId)
		event_id := row.Int8(iEventId)
		cloudTreasure, exist := mapDrivingSwordTreasureContent[cloud_id]
		if !exist {
			cloudTreasure = map[int8][]*DrivingSwordTreasureContent{}
		}
		cloudTreasure[event_id] = append(cloudTreasure[event_id], &DrivingSwordTreasureContent{
			Order:         row.Int8(iOrder),
			AwardItem:     row.Int16(iAwardItem),
			AwardNum:      row.Int32(iAwardNum),
			AwardCoinsNum: row.Int32(iAwardCoinsNum),
		})
		mapDrivingSwordTreasureContent[cloud_id] = cloudTreasure
	}
}

func loadDrivingSwordTreasureCount(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("select `cloud_id`, count(1) as `count` from `driving_sword_treasure` group by cloud_id;"), -1)
	if err != nil {
		panic(err)
	}

	iCloudId := res.Map("cloud_id")
	iCount := res.Map("count")

	mapTreasureCount = map[int16]int{}
	for _, row := range res.Rows {
		cloud_id := row.Int16(iCloudId)
		mapTreasureCount[cloud_id] = int(row.Int8(iCount))
	}
}

//获取对应阶段宝箱的奖品
func GetTreasureOrderAward(cloud_id int16, event_id int8, order int8) (awardItem int16, awardNum, awardCoins int32) {
	if awardlist, ok := mapDrivingSwordTreasureContent[cloud_id][event_id]; ok {
		for _, award := range awardlist {
			if award.Order == order {
				awardItem = award.AwardItem
				awardNum = award.AwardNum
				awardCoins = award.AwardCoinsNum
				return
			}
		}
	}

	return
}

//获取某层的宝箱总数
func CountTreasureByCloud(cloud int16) int {
	return mapTreasureCount[cloud]
}
