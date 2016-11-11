package item_dat

import (
	"core/fail"
	"core/mysql"
)

var (
	mapItemType map[int8]*ItemType
)

type ItemType struct {
	Id          int8  // 类型ID
	MaxNumInPos int16 // 每个位置最多可堆叠的数量
}

func loadItemType(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM item_type ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iId := res.Map("id")
	iMaxNumInPos := res.Map("max_num_in_pos")

	var pri_id int8
	mapItemType = map[int8]*ItemType{}
	for _, row := range res.Rows {
		pri_id = row.Int8(iId)
		mapItemType[pri_id] = &ItemType{
			Id:          pri_id,
			MaxNumInPos: row.Int16(iMaxNumInPos),
		}
	}
}

func GetItemType(id int32) *ItemType {
	v, ok := mapItemType[int8(id)]
	fail.When(!ok, "item type wrong id")
	return v
}
