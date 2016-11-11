package clique_dat

import (
	"core/fail"
	"core/mysql"
	"fmt"
	"math/rand"
)

var (
	mapCliqueBoat     map[int16]*CliqueBoat
	ingotBoat         *CliqueBoat
	randomBoatRateSum int32
)

type CliqueBoat struct {
	Id                      int16  // 主键
	Name                    string // 镖船名称
	Sign                    string // 资源标识
	Rate                    int8   // 概率
	EscortTime              int16  // 运送时间（单位分钟）
	SelectCostIngot         int64  // 直接选择话费元宝（0则不可直接选择）
	AwardCoins              int64  // 奖励铜钱
	AwardFame               int32  // 奖励声望
	AwardCliqueContrib      int64  // 奖励贡献
	HijackLoseCoins         int64  // 抢劫损失铜钱
	HijackLoseFame          int32  // 抢劫损失声望
	HijackLoseCliqueContrib int64  // 抢劫损失贡献
}

func loadCliqueBoat(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM clique_boat ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iId := res.Map("id")
	iName := res.Map("name")
	iSign := res.Map("sign")
	iRate := res.Map("rate")
	iEscortTime := res.Map("escort_time")
	iSelectCostIngot := res.Map("select_cost_ingot")
	iAwardCoins := res.Map("award_coins")
	iAwardFame := res.Map("award_fame")
	iAwardCliqueContrib := res.Map("award_clique_contrib")
	iHijackLoseCoins := res.Map("hijack_lose_coins")
	iHijackLoseFame := res.Map("hijack_lose_fame")
	iHijackLoseCliqueContrib := res.Map("hijack_lose_clique_contrib")

	var pri_id int16
	mapCliqueBoat = map[int16]*CliqueBoat{}
	randomBoatRateSum = 0
	for _, row := range res.Rows {
		pri_id = row.Int16(iId)
		mapCliqueBoat[pri_id] = &CliqueBoat{
			Id:                      pri_id,
			Name:                    row.Str(iName),
			Sign:                    row.Str(iSign),
			Rate:                    row.Int8(iRate),
			EscortTime:              row.Int16(iEscortTime),
			SelectCostIngot:         row.Int64(iSelectCostIngot),
			AwardCoins:              row.Int64(iAwardCoins),
			AwardFame:               row.Int32(iAwardFame),
			AwardCliqueContrib:      row.Int64(iAwardCliqueContrib),
			HijackLoseCoins:         row.Int64(iHijackLoseCoins),
			HijackLoseFame:          row.Int32(iHijackLoseFame),
			HijackLoseCliqueContrib: row.Int64(iHijackLoseCliqueContrib),
		}
		if row.Int64(iSelectCostIngot) > 0 {
			ingotBoat = mapCliqueBoat[pri_id]
		}
		randomBoatRateSum += int32(row.Int8(iRate))
	}
}

// ############### 对外接口实现 coding here ###############

func GetBoatById(id int16) *CliqueBoat {
	boat, ok := mapCliqueBoat[id]
	fail.When(!ok, fmt.Sprintf("unrecognized boat id %d ", id))
	return boat
}

func GetIngotBoat() *CliqueBoat {
	return ingotBoat
}

func GetRandomBoatId() int16 {
	rate := rand.Int31n(randomBoatRateSum) + 1
	for _, boat := range mapCliqueBoat {
		if boat.Rate <= 0 {
			continue
		}
		if rate <= int32(boat.Rate) {
			return boat.Id
		}
		rate -= int32(boat.Rate)
	}
	panic("随机不到可用的船")
}
