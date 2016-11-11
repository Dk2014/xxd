package item_dat

import (
	"core/fail"
	"core/mysql"
)

var (
	mapDragonBallConfig map[int16][]*DragonBallConfig
)

type DragonBallConfig struct {
	ItemId  int16 // 目标物品ID
	ItemNum int16 // 目标物品数量
	Rate    int16 // 目标物品概率
}

func loadDragonBallConfig(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM item_dragon_ball_config ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iSourceItem := res.Map("source_item")
	iItemId := res.Map("item_id")
	iItemNum := res.Map("item_num")
	iRate := res.Map("rate")

	var pri_id int16
	mapDragonBallConfig = map[int16][]*DragonBallConfig{}
	for _, row := range res.Rows {
		pri_id = row.Int16(iSourceItem)
		mapDragonBallConfig[pri_id] = append(mapDragonBallConfig[pri_id], &DragonBallConfig{
			ItemId:  row.Int16(iItemId),
			ItemNum: row.Int16(iItemNum),
			Rate:    row.Int16(iRate),
		})
	}
}

func GetDragonBallConfig(sourceItem int16) []*DragonBallConfig {
	configs, ok := mapDragonBallConfig[sourceItem]
	fail.When(!ok, "没有配置龙珠兑换物品")
	return configs
}
