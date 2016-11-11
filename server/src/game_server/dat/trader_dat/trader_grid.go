package trader_dat

import (
	"core/mysql"
)

var (
	mapTraderGrid    map[int32]*TraderGrid
	mapTraderGridIds map[int16][]int32
)

type TraderGrid struct {
	Id        int32 // ID
	TraderId  int16 // 随机商人ID
	MoneyType int8  // 货币类型
}

func loadTraderGrid(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM trader_grid ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iId := res.Map("id")
	iTraderId := res.Map("trader_id")
	iMoneyType := res.Map("money_type")

	var pri_id int32
	var trader_id int16
	mapTraderGrid = map[int32]*TraderGrid{}
	mapTraderGridIds = map[int16][]int32{}
	for _, row := range res.Rows {
		pri_id = row.Int32(iId)
		trader_id = row.Int16(iTraderId)
		mapTraderGrid[pri_id] = &TraderGrid{
			Id:        pri_id,
			TraderId:  trader_id,
			MoneyType: row.Int8(iMoneyType),
		}
		mapTraderGridIds[trader_id] = append(mapTraderGridIds[trader_id], pri_id)
	}
}
