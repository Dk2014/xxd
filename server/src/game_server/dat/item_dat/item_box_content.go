package item_dat

import (
	"core/fail"
	"core/mysql"
)

var (
	mapItemBoxContent map[int16][]*ItemBoxContent
)

type ItemBoxContent struct {
	ItemId      int16  // 物品宝箱的ID
	Type        int8   // 类型，0铜钱，1元宝，2物品
	Mode        int8   // 随机方式，0直接获得，1概率数量，2概率获得
	GetItemId   int16  // 得到的物品ID
	ItemIdSet   string // 随机的物品ID集
	MinNum      int32  // 最少数量
	MaxNum      int32  // 最多数量
	Probability int8   // 概率
}

func loadItemBoxContent(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM item_box_content ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iItemId := res.Map("item_id")
	iType := res.Map("type")
	iMode := res.Map("mode")
	iGetItemId := res.Map("get_item_id")
	iItemIdSet := res.Map("item_id_set")
	iMinNum := res.Map("min_num")
	iMaxNum := res.Map("max_num")
	iProbability := res.Map("probability")

	mapItemBoxContent = map[int16][]*ItemBoxContent{}
	for _, row := range res.Rows {

		itemId := row.Int16(iItemId)

		if mapItemBoxContent[itemId] == nil {
			mapItemBoxContent[itemId] = []*ItemBoxContent{}
		}

		mapItemBoxContent[itemId] = append(mapItemBoxContent[itemId], &ItemBoxContent{
			ItemId:      itemId,
			Type:        row.Int8(iType),
			Mode:        row.Int8(iMode),
			GetItemId:   row.Int16(iGetItemId),
			ItemIdSet:   row.Str(iItemIdSet),
			MinNum:      row.Int32(iMinNum),
			MaxNum:      row.Int32(iMaxNum),
			Probability: row.Int8(iProbability),
		})
	}
}

// ############### 对外接口实现 coding here ###############

func GetItemBoxContent(itemId int16) (itemBoxContents []*ItemBoxContent) {
	itemBoxContents = mapItemBoxContent[itemId]
	fail.When(itemBoxContents == nil, "Wrong item id for ItemBoxContent")
	return itemBoxContents
}
