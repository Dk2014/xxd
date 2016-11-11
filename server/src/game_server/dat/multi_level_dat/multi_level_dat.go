package multi_level_dat

import (
	"core/fail"
	"core/mysql"
)

var (
	mapMultiLevel       map[int16]*MultiLevel
	firstMultiLevelLock int32
)

func Load(db *mysql.Connection) {
	loadMultiLevel(db)
}

type MultiLevel struct {
	Id                int16
	Name              string
	RequireLevel      int16 // 主角等级要求
	AwardExp          int32 // 奖励经验
	AwardFame         int32 // 奖励声望
	AwardCoin         int64 // 奖励铜钱
	AwardRelationship int32 // 奖励友情
	AwardItem1Id      int32 // 奖励物品1 id
	AwardItem1Num     int32 // 物品1数量
	AwardItem2Id      int32 // 奖励物品2 id
	AwardItem2Num     int32 // 物品2数量
	AwardItem3Id      int32 // 奖励物品3 id
	AwardItem3Num     int32 // 物品3数量
	AwardLock         int32
	Lock              int32
}

func loadMultiLevel(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM multi_level ORDER BY `lock` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iId := res.Map("id")
	iRequireLevel := res.Map("require_level")
	iAwardExp := res.Map("award_exp")
	iAwardFame := res.Map("award_fame")
	iAwardCoin := res.Map("award_coin")
	iAwardRelationship := res.Map("award_relationship")
	iAwardItem1Id := res.Map("award_item1_id")
	iAwardItem1Num := res.Map("award_item1_num")
	iAwardItem2Id := res.Map("award_item2_id")
	iAwardItem2Num := res.Map("award_item2_num")
	iAwardItem3Id := res.Map("award_item3_id")
	iAwardItem3Num := res.Map("award_item3_num")
	iAwardLock := res.Map("award_lock")
	iName := res.Map("name")
	iLock := res.Map("lock")

	var pri_id int16
	firstMultiLevelLock = -1
	mapMultiLevel = map[int16]*MultiLevel{}
	for _, row := range res.Rows {
		pri_id = row.Int16(iId)
		mapMultiLevel[pri_id] = &MultiLevel{
			Id:                pri_id,
			RequireLevel:      row.Int16(iRequireLevel),
			AwardExp:          row.Int32(iAwardExp),
			AwardFame:         row.Int32(iAwardFame),
			AwardCoin:         row.Int64(iAwardCoin),
			AwardRelationship: row.Int32(iAwardRelationship),
			AwardItem1Id:      row.Int32(iAwardItem1Id),
			AwardItem1Num:     row.Int32(iAwardItem1Num),
			AwardItem2Id:      row.Int32(iAwardItem2Id),
			AwardItem2Num:     row.Int32(iAwardItem2Num),
			AwardItem3Id:      row.Int32(iAwardItem3Id),
			AwardItem3Num:     row.Int32(iAwardItem3Num),
			AwardLock:         row.Int32(iAwardLock),
			Lock:              row.Int32(iLock),
			Name:              row.Str(iName),
		}

		if firstMultiLevelLock == -1 {
			firstMultiLevelLock = row.Int32(iLock)
		}
	}
}

// ############### 对外接口实现 coding here ###############

func GetMultiLevelById(id int16) *MultiLevel {
	level, ok := mapMultiLevel[id]
	fail.When(!ok, "incorrect multi level id")
	return level
}

func GetFirstMultiLevelLock() int32 {
	return firstMultiLevelLock
}
