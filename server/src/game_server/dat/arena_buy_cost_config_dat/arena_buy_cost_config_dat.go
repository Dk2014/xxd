package arena_buy_cost_config_dat

import (
	"core/fail"
	"core/mysql"
)

var (
	mapArenaBuyCostConfig map[int32]int32
)

func Load(db *mysql.Connection) {
	loadArenaBuyCostConfig(db)
}

type ArenaBuyCostConfig struct {
	Times int32 // 购买次数
	Cost  int32 // 购买所需元宝
}

func loadArenaBuyCostConfig(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM arena_buy_cost_config ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iTimes := res.Map("times")
	iCost := res.Map("cost")

	mapArenaBuyCostConfig = make(map[int32]int32)
	for _, row := range res.Rows {
		mapArenaBuyCostConfig[row.Int32(iTimes)] = row.Int32(iCost)
	}
}

// ############### 对外接口实现 coding here ###############
//获取当天当次购买比武场次数所需元宝
func GetCost(times int32) int32 {
	times += 1 //次数从1,times为已购买次数
	cost, ok := mapArenaBuyCostConfig[times]
	if !ok {
		for max_times, max_cost := range mapArenaBuyCostConfig {
			if times > max_times { //超过设置上限，则按最大消耗元宝来
				cost = max_cost
			} else {
				fail.When(true, "arena buy cost config missing value ")
			}
		}
	}
	return cost
}
