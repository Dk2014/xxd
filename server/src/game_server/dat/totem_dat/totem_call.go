package totem_dat

import (
	"core/mysql"
)

var (
	mapTotemCallCostConfig map[int16]int32 //times -> price
	maxTottemCallPrice     int32
)

func loadTotemCallCostConfig(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM totem_call_cost_config ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iTimes := res.Map("times")
	iCost := res.Map("cost")

	var price int32
	mapTotemCallCostConfig = map[int16]int32{}
	maxTottemCallPrice = 0
	for _, row := range res.Rows {
		price = row.Int32(iCost)
		mapTotemCallCostConfig[row.Int16(iTimes)] = price
		if maxTottemCallPrice < price {
			maxTottemCallPrice = price
		}
	}
}

func GetTotemCallCost(times int16) (price int32) {
	if price, ok := mapTotemCallCostConfig[times]; ok {
		return price
	} else {
		return maxTottemCallPrice
	}
}
