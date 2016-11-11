package buy_hard_level_times_config_dat

import (
	"core/fail"
	"core/mysql"
)

var (
	mapBuyHardLevelTimesConfig map[int64]int64
)

func Load(db *mysql.Connection) {
	loadBuyHardLevelTimesConfig(db)
}

type BuyHardLevelTimesConfig struct {
	Times int64 // 购买次数
	Cost  int64 // 购买价格
}

func loadBuyHardLevelTimesConfig(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM buy_hard_level_times_config ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iTimes := res.Map("times")
	iCost := res.Map("cost")

	var pri_times int64
	mapBuyHardLevelTimesConfig = map[int64]int64{}
	for _, row := range res.Rows {
		pri_times = row.Int64(iTimes)
		mapBuyHardLevelTimesConfig[pri_times] = row.Int64(iCost)
	}
}

// ############### 对外接口实现 coding here ###############

func GetCost(times int64) int64 {
	cost, err := mapBuyHardLevelTimesConfig[times]
	if !err {
		fail.When(true, "wrong times")
	}
	return cost
}
