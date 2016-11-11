package role_dat

import (
	"core/mysql"
)

var (
	mapRoleLevelExp [MAX_ROLE_LEVEL]int64
)

func loadRoleLevelExp(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM role_level_exp ORDER BY `level` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iLevel := res.Map("level")
	iExp := res.Map("exp")

	defer func() {
		if err := recover(); err != nil {
			panic(err)
		}
	}()

	for _, row := range res.Rows {
		mapRoleLevelExp[row.Int32(iLevel)-1] = row.Int64(iExp)
	}
}
