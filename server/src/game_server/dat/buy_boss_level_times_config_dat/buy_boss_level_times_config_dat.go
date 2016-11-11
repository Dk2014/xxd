package buy_boss_level_times_config_dat

import (
	"core/fail"
	"core/mysql"
)

var (
	mapBuyBossLevelTimesConfig map[int64]int64
)

func Load(db *mysql.Connection) {
	loadBuyBossLevelTimesConfig(db)
}

type BuyBossLevelTimesConfig struct {
	Times int64 // 购买次数
	Cost  int64 // 购买价格
}

func loadBuyBossLevelTimesConfig(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM buy_boss_level_times_config ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iTimes := res.Map("times")
	iCost := res.Map("cost")

	var pri_times int64
	mapBuyBossLevelTimesConfig = map[int64]int64{}
	for _, row := range res.Rows {
		pri_times = row.Int64(iTimes)
		mapBuyBossLevelTimesConfig[pri_times] = row.Int64(iCost)
	}
}

// ############### 对外接口实现 coding here ###############

func GetCost(times int64) int64 {
	cost, err := mapBuyBossLevelTimesConfig[times]
	if !err {
		fail.When(true, "wrong times")
	}
	return cost
}
