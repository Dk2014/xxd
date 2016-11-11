package coins_exchange_dat

import (
	"core/mysql"
)

func Load(db *mysql.Connection) {
	loadCoinsExchange(db)
}

//获取第unique_key次购买的信息
func GetCoinsExchangeInfo(unique_key int16) *CoinsExchange {
	return mapCoinsExchange[unique_key]
}
