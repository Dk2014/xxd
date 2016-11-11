package trader_dat

import (
	"core/mysql"
)

var (
	mapTraderRefreshPrice map[int16]map[int16]int64 //trader -> time -> price
)

func loadTraderRefreshPrice(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM trader_refresh_price ORDER BY `time` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iPrice := res.Map("price")
	iTraderId := res.Map("trader_id")
	iTime := res.Map("time")

	mapTraderRefreshPrice = make(map[int16]map[int16]int64, 2)
	for _, row := range res.Rows {
		traderId := row.Int16(iTraderId)
		if mapTraderRefreshPrice[traderId] == nil {
			mapTraderRefreshPrice[traderId] = make(map[int16]int64, 0)
		}
		mapTraderRefreshPrice[traderId][row.Int16(iTime)] = int64(row.Int32(iPrice))
	}
}
