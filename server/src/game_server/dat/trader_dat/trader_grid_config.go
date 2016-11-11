package trader_dat

import (
	"core/mysql"
)

var (
	mapTraderGridConfig          map[int32][]*TraderGridConfig
	mapTraderGridConfigEquiement map[int32][]*TraderGridConfigEquiement
)

type TraderGridConfig struct {
	Id          int32 // 配置ID
	GoodsType   int8  // 物品类型
	ItemId      int16 // 物品ID
	Num         int16 // 物品数量
	Cost        int64 // 价格
	Probability int8  // 出现概率（％）
}

func loadTraderGridConfig(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM trader_grid_config ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iId := res.Map("id")
	iGridId := res.Map("grid_id")
	iGoodsType := res.Map("goods_type")
	iItemId := res.Map("item_id")
	iNum := res.Map("num")
	iCost := res.Map("cost")
	iProbability := res.Map("probability")

	var pri_id int32
	mapTraderGridConfig = map[int32][]*TraderGridConfig{}
	for _, row := range res.Rows {
		pri_id = row.Int32(iGridId)
		mapTraderGridConfig[pri_id] = append(mapTraderGridConfig[pri_id], &TraderGridConfig{
			Id:          row.Int32(iId),
			GoodsType:   row.Int8(iGoodsType),
			ItemId:      row.Int16(iItemId),
			Num:         row.Int16(iNum),
			Cost:        row.Int64(iCost),
			Probability: row.Int8(iProbability),
		})
	}
}

type TraderGridConfigEquiement struct {
	//ConfigId int32 // 配置ID
	Cost     int64 // 价格
	ItemId   int16 // 物品ID
	MinLevel int16 // 等级下限
	MaxLevel int16 // 等级上限
}

func loadTraderGridConfigEquiement(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM trader_grid_config_equiement ORDER BY `config_id`,`min_level` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iConfigId := res.Map("config_id")
	iCost := res.Map("cost")
	iItemId := res.Map("item_id")
	iMinLevel := res.Map("min_level")
	iMaxLevel := res.Map("max_level")

	var pri_id int32
	mapTraderGridConfigEquiement = map[int32][]*TraderGridConfigEquiement{}
	for _, row := range res.Rows {
		pri_id = row.Int32(iConfigId)
		mapTraderGridConfigEquiement[pri_id] = append(mapTraderGridConfigEquiement[pri_id], &TraderGridConfigEquiement{
			Cost:     row.Int64(iCost),
			ItemId:   row.Int16(iItemId),
			MinLevel: row.Int16(iMinLevel),
			MaxLevel: row.Int16(iMaxLevel),
		})
	}
}
