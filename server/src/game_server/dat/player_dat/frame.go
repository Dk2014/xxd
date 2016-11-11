package player_dat

import (
	"core/fail"
	"core/mysql"
)

var (
	mapFameLevel map[int16]int32 //level -> required_fame
	mapMaxFame   map[int16]int32 //system -> max_fame
)

func loadFameSystem(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM fame_system ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iId := res.Map("id")
	iMaxFame := res.Map("max_fame")

	mapMaxFame = map[int16]int32{}
	for _, row := range res.Rows {
		mapMaxFame[row.Int16(iId)] = row.Int32(iMaxFame)
	}
}

func loadFameLevel(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM fame_level ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iLevel := res.Map("level")
	iRequiredFame := res.Map("required_fame")

	mapFameLevel = map[int16]int32{}
	for _, row := range res.Rows {
		mapFameLevel[row.Int16(iLevel)] = row.Int32(iRequiredFame)
	}
}

func NextLevelRequiredFame(nextLevel int16) (fame int32, haveNextLv bool) {
	fame, haveNextLv = mapFameLevel[nextLevel]
	return fame, haveNextLv
}

func MaxFameForSystem(system int16) int32 {
	maxFame, ok := mapMaxFame[system]
	fail.When(!ok, "未定义的系统")
	return maxFame
}
