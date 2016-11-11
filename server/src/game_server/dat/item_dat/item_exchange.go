package item_dat

import (
	"core/mysql"
)

var (
	mapItemExchange map[int64]*ItemExchange
)

type ItemExchange struct {
	Id           int64 // 主键ID
	TargetItemId int16 // 目标物品id
	ItemId       int16 // 物品id
	ItemNum      int16 // 物品数量
}

func loadItemExchange(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM item_exchange ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iId := res.Map("id")
	iTargetItemId := res.Map("target_item_id")
	iItemId := res.Map("item_id")
	iItemNum := res.Map("item_num")

	var pri_id int64
	mapItemExchange = map[int64]*ItemExchange{}
	for _, row := range res.Rows {
		pri_id = row.Int64(iId)
		mapItemExchange[pri_id] = &ItemExchange{
			Id:           pri_id,
			TargetItemId: row.Int16(iTargetItemId),
			ItemId:       row.Int16(iItemId),
			ItemNum:      row.Int16(iItemNum),
		}
	}
}

func FetchItemExchange(itemId int16, cb func(itemExchange *ItemExchange)) (res bool) {
	res = false
	for _, v := range mapItemExchange {
		if itemId == v.TargetItemId {
			cb(v)
			res = true
			break
		}
	}
	return
}
