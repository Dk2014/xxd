package item_dat

import (
	"core/mysql"
)

var (
	mapSealedBook map[int8][]*SealedBook
)

type SealedBook struct {
	//Id          int16 // 物品ID
	//Type        int8  // 类型(1- - 伙伴；1 - 魂侍；5 - 怪物；7 - 魂侍技1； 8 - 魂侍技2)
	ItemId      int16 // 物品id
	Health      int32 // 生命
	Attack      int32 // 攻击
	Defense     int32 // 防御
	Cultivation int32 // 内力
	Coins       int32 //激活消耗铜钱
}

func loadSealedBook(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM sealed_book ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	//iId := res.Map("id")
	iType := res.Map("type")
	iItemId := res.Map("item_id")
	iHealth := res.Map("health")
	iAttack := res.Map("attack")
	iDefense := res.Map("defense")
	iCultivation := res.Map("cultivation")
	iCoins := res.Map("coins")

	mapSealedBook = map[int8][]*SealedBook{}
	for _, row := range res.Rows {
		bookType := row.Int8(iType)

		mapSealedBook[bookType] = append(mapSealedBook[bookType], &SealedBook{
			ItemId:      row.Int16(iItemId),
			Health:      row.Int32(iHealth),
			Attack:      row.Int32(iAttack),
			Defense:     row.Int32(iDefense),
			Cultivation: row.Int32(iCultivation),
			Coins:       row.Int32(iCoins),
		})
	}
}

// ############### 对外接口实现 coding here ###############

func GetSealedBookInfo(itemType int8, itemID int16) *SealedBook {
	value, ok := mapSealedBook[itemType]
	if !ok {
		return nil
	}
	for _, itemList := range value {
		if itemID == itemList.ItemId {
			return itemList
		}
	}
	return nil
}

/*
func GetSealedBookType(itemID int16) (int8, bool) {
	for itemtype, itemidList := range mapSealedBook {
		for _, itemids := range itemidList {
			if (itemID == itemids.ItemId) &&
				(itemtype == STEALDBOOK_TYPE_EQUIPMENTS || itemtype == STEALDBOOK_TYPE_NOEQUIPMENTS || itemtype == STEALDBOOK_TYPE_BATTLETOOLS) {
				return itemtype, true
			}
		}
	}

	return 0, false
}
*/
