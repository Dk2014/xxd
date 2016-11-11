package item_dat

import (
	"core/fail"
	"core/mysql"
)

var (
	mapItemPurchaseLimit map[int16]int16 //ItemId -> limit
)

func loadPurchaseLimit(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM purchase_limit ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iItemId := res.Map("item_id")
	iNum := res.Map("num")

	mapItemPurchaseLimit = map[int16]int16{}
	for _, row := range res.Rows {
		mapItemPurchaseLimit[row.Int16(iItemId)] = row.Int16(iNum)
	}
}

//物品购买次数限制
func PurchaseLimitation(itemId int16) int16 {
	num, ok := mapItemPurchaseLimit[itemId]
	fail.When(!ok, "商品未配置购买次数限制，不允许购买")
	return num
}
