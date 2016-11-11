package ghost_dat

import (
	"core/mysql"
)

var (
	mapGhostSkillTrainPrice      map[int16]int64
	mapGhostSkillTrainTotalPrice map[int16]int64
)

func loadGHostSkillTrainPrice(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM ghost_skill_train_price  ORDER BY `level` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iLevel := res.Map("level")
	iCost := res.Map("cost")

	mapGhostSkillTrainPrice = map[int16]int64{}
	for _, row := range res.Rows {
		mapGhostSkillTrainPrice[row.Int16(iLevel)] = int64(row.Int32(iCost))
	}

	mapGhostSkillTrainTotalPrice = map[int16]int64{}
	for level, _ := range mapGhostSkillTrainPrice {
		calcSkillTrainTotalPrice(level)
	}
}

func calcSkillTrainTotalPrice(level int16) int64 {
	if level == 0 {
		return 0
	}
	if rtn, ok := mapGhostSkillTrainTotalPrice[level]; ok {
		return rtn
	}
	preTotalPrice, exist := mapGhostSkillTrainTotalPrice[level-1]
	if !exist {
		preTotalPrice = calcSkillTrainTotalPrice(level - 1)
	}
	mapGhostSkillTrainTotalPrice[level] = preTotalPrice + mapGhostSkillTrainPrice[level]
	return mapGhostSkillTrainTotalPrice[level]
}

// ############### 对外接口实现 coding here ###############

func GetGhostSkillTrainPrice(level int16) int64 {
	return mapGhostSkillTrainPrice[level]
}

func GetGhostSkillTrainTotalPrice(level int16) int64 {
	return mapGhostSkillTrainTotalPrice[level]
}
