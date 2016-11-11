package ghost_dat

import (
	"core/fail"
	"core/mysql"
)

var (
	mapGhostTrainPrice   []*GhostTrainPrice
	mapGhostUpgradePrice map[int8]int64 //quality => price

)

type GhostTrainPrice struct {
	Times int32
	Price int64
}

func loadGhostUpgradePrice(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM ghost_upgrade_price ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iQuality := res.Map("quality")
	iCost := res.Map("cost")

	mapGhostUpgradePrice = map[int8]int64{}
	for _, row := range res.Rows {
		mapGhostUpgradePrice[row.Int8(iQuality)] = int64(row.Int32(iCost))
	}
}

func loadGhostTrainPrice(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM ghost_train_price ORDER BY `times` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iTimes := res.Map("times")
	iCost := res.Map("cost")
	mapGhostTrainPrice = []*GhostTrainPrice{}
	for _, row := range res.Rows {
		mapGhostTrainPrice = append(mapGhostTrainPrice, &GhostTrainPrice{
			Times: row.Int32(iTimes),
			Price: int64(row.Int64(iCost)),
		})
	}
}

//获取魂力果实单价
func GetGhostTrainPrice(times int32) (price int64) {
	for _, data := range mapGhostTrainPrice {
		if data.Times > times {
			break
		}
		price = data.Price
	}
	return price
}

//获取魂侍升级碎片价格
func GetGhostUpgradePrice(quality int8) int64 {
	price, ok := mapGhostUpgradePrice[quality]
	fail.When(!ok, "未配置的魂侍品质")
	return price
}
