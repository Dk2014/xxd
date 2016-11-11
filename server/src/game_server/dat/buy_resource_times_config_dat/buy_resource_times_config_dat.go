package buy_resource_times_config_dat

import (
	"core/fail"
	"core/mysql"
)

var (
	mapBuyResourceTimesConfig map[int32]int32
)

func Load(db *mysql.Connection) {
	loadBuyResourceTimesConfig(db)
}

type BuyResourceTimesConfig struct {
	Id    int32 //
	Times int32 // 购买次数
	Cost  int32 // 购买所需元宝
}

func loadBuyResourceTimesConfig(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM buy_resource_times_config ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iTimes := res.Map("times")
	iCost := res.Map("cost")

	mapBuyResourceTimesConfig = make(map[int32]int32)
	for _, row := range res.Rows {
		mapBuyResourceTimesConfig[row.Int32(iTimes)] = row.Int32(iCost)
	}
}

// ############### 对外接口实现 coding here ###############
//按照当天购买体力次数获取所需消耗元宝
func GetCost(times int32) int32 {
	times += 1 //次数从1开始算起
	cost, ok := mapBuyResourceTimesConfig[times]
	if !ok {
		for max_times, max_cost := range mapBuyResourceTimesConfig {
			if times > max_times { //超过设置上限，则按最大消耗元宝来
				cost = max_cost
			} else {
				fail.When(true, "buy resource times cost config missing value ")
			}
		}
	}
	return cost
}
