package item_dat

import (
	"core/fail"
	"core/mysql"
)

var (
	mapItemReflectConfig map[int16][]*ItemReflectConfig
)

type ItemReflectConfig struct {
	ItemId       int16 // 物品ID
	AwardCoinMin int32 // 最小奖励铜钱
	AwardCoinMax int32 // 最大奖励铜钱
}

func loadItemReflectConfig(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM `item_reflect_config` ORDER BY `item_id`, `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iItemId := res.Map("item_id")
	iAwardCoinMin := res.Map("award_coin_min")
	iAwardCoinMax := res.Map("award_coin_max")

	mapItemReflectConfig = map[int16][]*ItemReflectConfig{}
	for _, row := range res.Rows {

		itemId := row.Int16(iItemId)

		if mapItemReflectConfig[itemId] == nil {
			mapItemReflectConfig[itemId] = []*ItemReflectConfig{}
		}

		mapItemReflectConfig[itemId] = append(mapItemReflectConfig[itemId], &ItemReflectConfig{
			ItemId:       itemId,
			AwardCoinMin: row.Int32(iAwardCoinMin),
			AwardCoinMax: row.Int32(iAwardCoinMax),
		})
	}
}

func GetItemReflectConfig(itemId int16) (config []*ItemReflectConfig) {
	config = mapItemReflectConfig[itemId]
	fail.When(config == nil, "Wrong item id for ItemReflectConfig")
	return
}
