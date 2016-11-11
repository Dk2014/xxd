package fashion_dat

import (
	"core/fail"
	"core/mysql"
)

var (
	mapFashion       map[int16]*Fashion //fashion id -> fashion data
	mapItemToFashion map[int16]int16    //item id -> fashion id
)

func Load(db *mysql.Connection) {
	loadFashion(db)
	loadFashionExchange(db)
}

type Fashion struct {
	Health      int32 // 生命
	Speed       int32 // 速度
	Cultivation int32 // 内力
	Attack      int32 // 攻击
	Defence     int32 // 防御
	Level       int32 //要求等级
}

func loadFashion(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM fashion ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iId := res.Map("id")
	iHealth := res.Map("health")
	iSpeed := res.Map("speed")
	iCultivation := res.Map("cultivation")
	iAttack := res.Map("attack")
	iDefence := res.Map("defence")
	iLevel := res.Map("level")

	var pri_id int16
	mapFashion = map[int16]*Fashion{}
	for _, row := range res.Rows {
		pri_id = row.Int16(iId)
		mapFashion[pri_id] = &Fashion{
			Health:      row.Int32(iHealth),
			Speed:       row.Int32(iSpeed),
			Cultivation: row.Int32(iCultivation),
			Attack:      row.Int32(iAttack),
			Defence:     row.Int32(iDefence),
			Level:       row.Int32(iLevel),
		}

	}
}

func loadFashionExchange(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM fashion_exchange ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iFashionId := res.Map("fashion_id")
	iItemId := res.Map("item_id")

	mapItemToFashion = map[int16]int16{}
	for _, row := range res.Rows {
		mapItemToFashion[row.Int16(iItemId)] = row.Int16(iFashionId)
	}
}

// ############### 对外接口实现 coding here ###############

func GetFashionData(id int16) *Fashion {
	fashion, ok := mapFashion[id]
	fail.When(!ok, "找不到时装数据")
	return fashion
}

func GetFashionIdByItemId(itemId int16) int16 {
	fashionId, ok := mapItemToFashion[itemId]
	fail.When(!ok, "找不到关联的时装")
	return fashionId
}
