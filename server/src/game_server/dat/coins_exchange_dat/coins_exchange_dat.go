package coins_exchange_dat

import (
	"core/mysql"
)

var (
	mapCoinsExchange map[int16]*CoinsExchange
)


type CoinsExchange struct {
	UniqueKey int16 // 第几次兑换
	Ingot int64 // 消耗元宝
	Coins int64 // 获得铜币
}


func loadCoinsExchange(db *mysql.Connection) {	
	res, err := db.ExecuteFetch([]byte("SELECT * FROM coins_exchange ORDER BY `unique_key` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iUniqueKey := res.Map("unique_key")
	iIngot := res.Map("ingot")
	iCoins := res.Map("coins")


	var pri_id int16
	mapCoinsExchange = map[int16]*CoinsExchange{}
	for _, row := range res.Rows {
		pri_id = row.Int16(iUniqueKey)
		mapCoinsExchange[pri_id] = &CoinsExchange{
			UniqueKey : pri_id,
			Ingot : row.Int64(iIngot),
			Coins : row.Int64(iCoins),

		}
	}
}
