package chest_dat

import (
	"core/fail"
	"core/mysql"
)

var (
	mapChestItem map[int32][]*ChestItem
)

type ChestItem struct {
	ChestId int32 // 宝箱id
	Type    int8  // 类型
	ItemId  int16 // 物品
	ItemNum int32 // 数量
}

func loadChestItem(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM chest_item ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iChestId := res.Map("chest_id")
	iType := res.Map("type")
	iItemId := res.Map("item_id")
	iItemNum := res.Map("item_num")

	var chestId int32
	mapChestItem = map[int32][]*ChestItem{}

	for _, row := range res.Rows {
		chestId = row.Int32(iChestId)

		if mapChestItem[chestId] == nil {
			mapChestItem[chestId] = []*ChestItem{}
		}

		mapChestItem[chestId] = append(mapChestItem[chestId], &ChestItem{
			ChestId: chestId,
			Type:    row.Int8(iType),
			ItemId:  row.Int16(iItemId),
			ItemNum: row.Int32(iItemNum),
		})
	}
}

// ############### 对外接口实现 coding here ###############

func GetChestItems(chestId int32) (chestItem []*ChestItem) {
	chestItem = mapChestItem[chestId]
	fail.When(chestItem == nil, "wrong chest id")
	return chestItem
}
