package battle_pet_dat

/*

import (
	"core/mysql"
)

var (
	mapBattlePetSoulExchange map[int8]map[int8]int16 // quality -> start -> soul num
)

func loadBattlePetSoulExchange(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM battle_pet_soul_exchange ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iQuality := res.Map("quality")
	iStar := res.Map("star")
	iSoulNum := res.Map("soul_num")

	var quality, star int8
	mapBattlePetSoulExchange = map[int8]map[int8]int16{}
	for _, row := range res.Rows {
		quality = row.Int8(iQuality)
		star = row.Int8(iStar)
		if mapBattlePetSoulExchange[quality] == nil {
			mapBattlePetSoulExchange[quality] = map[int8]int16{}
		}
		mapBattlePetSoulExchange[quality][star] = row.Int16(iSoulNum)
	}
}

//根据灵宠（怪物)ID获取对应兑换的灵宠魂魄数量
func GetBattlePetSoulNum(enemyId int32) int16 {
	petDat := GetBattlePetWithEnemyId(enemyId)
	return mapBattlePetSoulExchange[petDat.Quality][petDat.Star]
}
*/
