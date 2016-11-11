package trader_dat

import (
	"core/fail"
	"core/mysql"
)

func Load(db *mysql.Connection) {
	loadTraderGrid(db)
	loadTraderGridConfig(db)
	loadTraderRefreshPrice(db)
	loadTraderSchedule(db)
	loadTraderGridConfigEquiement(db)
}

//获取箱子消费信息：货币类型、所属商人ID
func GridInfo(id int32) *TraderGrid {
	grid := mapTraderGrid[id]
	fail.When(grid == nil, "trader grid no found")
	return grid
}

//输入格子ID，返回配置的商品结构体列表
func GridConfig(id int32) []*TraderGridConfig {
	configs, ok := mapTraderGridConfig[id]
	fail.When(!ok || configs == nil, "grid is not configured")
	return configs
}

//刷新随机商人价格
func RefreshPrice(traderId, t int16) int64 {
	fail.When(t < 1, "wrong time")
	priceMap, ok := mapTraderRefreshPrice[traderId]
	fail.When(!ok, "trader is not configured")
	return priceMap[t]
}

func TraderGridIds(traderId int16) []int32 {
	ids := mapTraderGridIds[traderId]
	fail.When(ids == nil, "trader grid no found")
	return ids
}

func Refreshable(traderId int16) bool {
	for _, id := range RefreshableTrader {
		if id == traderId {
			return true
		}
	}
	return false
}

//等级装备类型货物的真实装备类型
func GetGridConfigEquiement(configId int32, level int16) *TraderGridConfigEquiement {
	equipments, ok := mapTraderGridConfigEquiement[configId]
	fail.When(!ok, "没有配置装备")
	var eqm *TraderGridConfigEquiement
	for _, e := range equipments {
		if e.MinLevel <= level && level <= e.MaxLevel {
			return e
		}
		eqm = e
	}
	return eqm
}
