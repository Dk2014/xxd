package trader_dat

import (
	"core/mysql"
	"core/i18l"
	
)

const( 
	YINGHAIJISHI = 1
	XUNYOUSHANGREN = 2
	HEISHILAODA = 3
)
const( 
	COINS_TYPE = 0
	INGOT_TYPE = 1
	HEART_TYPE = 2
	DRAGON_COIN_TYPE = 3
)
const( 
	ITEM = 0
	SWORD_SOUL = 1
	GHOST = 2
	PET = 3
	HEART = 4
	COINS = 5
	INGOT = 6
	PHYSICAL = 7
	EQUIPMENT = 8
)
const( 
	NOTIFY_IN_ADVANCE = 300
	YINGHAISHANGREN_REQUIRE_LEVEL = 14
)


var RefreshableTrader = []int16{1,3,}

var YingHaiJiShiRefresh = []int64{43200,64800,75600,}

var XunYouShangRenRefresh = []int64{36000,79200,}

var HeiShiLaoDaRefresh = []int64{}





//出现时间变量
var (
	MapTraderSchedule map[int16][]*TraderSchedule //出现时间配置
)

//初始化加随机商人出现时间和刷新时间
//每个商人默认会有一个24小时出现永不过期的配置
func loadTraderSchedule(db *mysql.Connection) {
	MapTraderSchedule = make(map[int16][]*TraderSchedule)
	//瀛海集市配置
	MapTraderSchedule[1] = make([]*TraderSchedule, 0, 1)
	MapTraderSchedule[1] = append(MapTraderSchedule[1], &TraderSchedule{Expire: 0, Showup: 0, Disappear: 86400})
	//巡游商人配置
	//巡游商人出现时间需要特别处理
	MapTraderSchedule[2] = make([]*TraderSchedule, 0, 2)
	MapTraderSchedule[2] = append(MapTraderSchedule[2], &TraderSchedule{Expire: 0, Showup: 36000, Disappear: 79200})
	//黑市老大配置
	MapTraderSchedule[3] = make([]*TraderSchedule, 0, 1)
	MapTraderSchedule[3] = append(MapTraderSchedule[3], &TraderSchedule{Expire: 0, Showup: 0, Disappear: 86400})
}
//根据商人ID查询刷新时间计划
func TraderRefreshSchedule(traderId int16) []int64{
	switch traderId {
	case 1 :
		return YingHaiJiShiRefresh
	case 2 :
		return XunYouShangRenRefresh
	case 3 :
		return HeiShiLaoDaRefresh
	default:
		panic("undefine trader")
	}
}
//根据商人ID查询商人名称
func GetTraderNameById(traderId int16) string{
	switch traderId {
	case 1 :
		return i18l.T.Tran("瀛海集市")
	case 2 :
		return i18l.T.Tran("巡游商人")
	case 3 :
		return i18l.T.Tran("黑市老大")
	default:
		panic("undefine trader")
	}
}

