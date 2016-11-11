package chest_dat

import (
	"core/fail"
	"core/mysql"
)

var (
	mapChest         map[int8][]*Chest
	mapFixChestAward map[int8]map[int32][]*Chest
)

func Load(db *mysql.Connection) {
	loadChest(db)
	loadChestItem(db)
	loadFateBox(db)
}

type Chest struct {
	Id            int32 // 主键
	Type          int8  // 类型:1 - 青铜宝箱, 2 - 神龙宝箱
	Quality       int8  // 宝箱品质
	Probability   int8  // 概率（%）
	FixAwardCount int32 //固定抽奖次数
}

func loadChest(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM chest ORDER BY `type`, `fix_award_count`, `quality` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iId := res.Map("id")
	iType := res.Map("type")
	iQuality := res.Map("quality")
	iProbability := res.Map("probability")
	iFixAwardCount := res.Map("fix_award_count")

	var typeId int8
	var fixAwardCount int32

	mapChest = map[int8][]*Chest{}
	mapFixChestAward = map[int8]map[int32][]*Chest{}

	for _, row := range res.Rows {
		typeId = row.Int8(iType)
		fixAwardCount = int32(row.Int8(iFixAwardCount))
		if mapFixChestAward[typeId] == nil {
			mapFixChestAward[typeId] = map[int32][]*Chest{}
		}
		if fixAwardCount > 0 {
			mapFixChestAward[typeId][fixAwardCount] = append(mapFixChestAward[typeId][fixAwardCount], &Chest{
				Id:          row.Int32(iId),
				Type:        typeId,
				Probability: 100,
			})
		} else {
			mapChest[typeId] = append(mapChest[typeId], &Chest{
				Id:          row.Int32(iId),
				Type:        typeId,
				Quality:     row.Int8(iQuality),
				Probability: row.Int8(iProbability),
			})
		}
	}
}

// ############### 对外接口实现 coding here ###############

func GetChestsByType(typeId int8) (chests []*Chest) {
	chests = mapChest[typeId]
	fail.When(chests == nil, "wrong type id")
	return chests
}

func GetChestFixAward(typeId int8, times int32) (chests []*Chest) {
	chestsConfig := mapFixChestAward[typeId]
	fail.When(chestsConfig == nil, "wrong type id")
	chests = chestsConfig[times]
	return chests
}
