package physical_buy_cost_config_dat

import (
	"core/mysql"
)

var (
	mapPhysicalBuyCostConfig map[int32]int32
)

func Load(db *mysql.Connection) {
	loadPhysicalBuyCostConfig(db)
}

type PhysicalBuyCostConfig struct {
	Times int32 // 购买次数
	Cost  int32 // 购买所需元宝
}

func loadPhysicalBuyCostConfig(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT `times`,`cost` FROM physical_buy_cost_config ORDER BY `times` DESC"), -1)
	if err != nil {
		panic(err)
	}

	iTimes := res.Map("times")
	iCost := res.Map("cost")

	mapPhysicalBuyCostConfig = make(map[int32]int32)
	for _, row := range res.Rows {
		mapPhysicalBuyCostConfig[row.Int32(iTimes)] = row.Int32(iCost)
	}
}

// ############### 对外接口实现 coding here ###############
//按照当天购买体力次数获取所需消耗元宝
func GetCost(times int32) int32 {
	times += 1 //次数从1开始算起
	cost, ok := mapPhysicalBuyCostConfig[times]
	if !ok {
		for _, max_cost := range mapPhysicalBuyCostConfig {
			if max_cost > cost { //超过设置上限，则按最大消耗元宝来
				cost = max_cost
			}
		}
	}
	return cost
}
