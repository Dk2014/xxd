package rainbow_buy_cost_config_dat

import (
	"core/fail"
	"core/mysql"
)

var (
	mapRainbowBuyCostConfig map[int32]int32
)

func Load(db *mysql.Connection) {
	loadRainbowBuyCostConfig(db)
}

type RainbowBuyCostConfig struct {
	Times int32 // 购买次数
	Cost  int32 // 购买所需元宝
}

func loadRainbowBuyCostConfig(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM rainbow_buy_cost_config ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iTimes := res.Map("times")
	iCost := res.Map("cost")

	mapRainbowBuyCostConfig = make(map[int32]int32)
	for _, row := range res.Rows {
		mapRainbowBuyCostConfig[row.Int32(iTimes)] = row.Int32(iCost)
	}
}

// ############### 对外接口实现 coding here ###############
//获取当天当次购买彩虹关卡扫荡次数所需元宝
func GetCost(times int32) int32 {
	times += 1 //次数从1,times为已购买次数
	cost, ok := mapRainbowBuyCostConfig[times]
	if !ok {
		for max_times, max_cost := range mapRainbowBuyCostConfig {
			if times > max_times { //超过设置上限，则按最大消耗元宝来
				cost = max_cost
			} else {
				fail.When(true, "arena buy cost config missing value ")
			}
		}
	}
	return cost
}
