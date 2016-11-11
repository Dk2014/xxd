package clique_dat

// 帮派建筑物 钱庄
import (
	"core/mysql"
)

type CliqueBuildingBankInfo struct {
	//Level             int16
	//SilverCouponCoins int64
	//SilverCouponNum   int8
	//SilverCouponSold  int64
	//GoldCouponCoins   int64
	//GoldCouponNum     int8
	//GoldCouponSold    int64
	//Upgrade           int64
	//Desc              string
	Level             int16 // 钱庄等级
	SilverCouponCoins int64
	SilverCouponNum   int16 // 银劵限购
	SilverCouponSold  int64 // 银劵售价
	GoldCouponCoins   int64 // 金劵底价
	GoldCouponNum     int16 // 金劵限购
	GoldCouponSold    int64 // 金劵售价
	Upgrade           int64 // 钱庄升级所需铜钱
}

var mapCliqueBuildingBankInfo map[int16]*CliqueBuildingBankInfo

func loadCliqueBuildingBankDat(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM `clique_building_bank` ORDER BY `level` ASC"), -1)
	if err != nil {
		panic(err)
	}
	iLevel := res.Map("level")
	iSilverCouponCoins := res.Map("silver_coupon_coins")
	iSilverCouponNum := res.Map("silver_coupon_num")
	iSilverCouponSold := res.Map("silver_coupon_sold")
	iGoldCouponCoins := res.Map("gold_coupon_coins")
	iGoldCouponNum := res.Map("gold_coupon_num")
	iGoldCouponSold := res.Map("gold_coupon_sold")
	iUpgrade := res.Map("upgrade")
	//iDesc := res.Map("desc")

	mapCliqueBuildingBankInfo = make(map[int16]*CliqueBuildingBankInfo)

	for _, row := range res.Rows {
		pri_level := row.Int16(iLevel)
		mapCliqueBuildingBankInfo[pri_level] = &CliqueBuildingBankInfo{
			Level:             pri_level,
			SilverCouponCoins: int64(row.Int32(iSilverCouponCoins)),
			SilverCouponNum:   row.Int16(iSilverCouponNum),
			SilverCouponSold:  int64(row.Int32(iSilverCouponSold)),
			GoldCouponCoins:   int64(row.Int32(iGoldCouponCoins)),
			GoldCouponNum:     row.Int16(iGoldCouponNum),
			GoldCouponSold:    int64(row.Int32(iGoldCouponSold)),
			Upgrade:           row.Int64(iUpgrade),
			//Desc:              row.Str(iDesc),
		}
	}
}

// ############### 对外接口实现 coding here ###############
func GetCliqueBuildingBankInfo(level int16) *CliqueBuildingBankInfo {
	if info, ok := mapCliqueBuildingBankInfo[level]; ok {
		return info
	}
	return nil
}
