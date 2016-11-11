package tower_level_dat

import (
	"core/fail"
	"core/mysql"
)

var (
	mapTowerLevel map[int16]int16 // map[楼层][id]
)

func Load(db *mysql.Connection) {
	loadTowerLevel(db)
}

func loadTowerLevel(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM tower_level ORDER BY `floor` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iId := res.Map("id")
	iFloor := res.Map("floor")

	mapTowerLevel = make(map[int16]int16)

	for _, row := range res.Rows {
		mapTowerLevel[row.Int16(iFloor)] = row.Int16(iId)
	}
}

// ############### 对外接口实现 coding here ###############

func GetTowerIdByFloor(floor int16) int16 {
	id, ok := mapTowerLevel[floor]
	fail.When(!ok, "floor error")
	return id
}
