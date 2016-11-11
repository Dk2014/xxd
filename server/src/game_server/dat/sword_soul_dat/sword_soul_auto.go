package sword_soul_dat


const( 
	MAX_DRAW_NUM = 100
	MAX_INGOT_DRAW_NUM = 20
	RECOVERY_VALUE = 1
	RECOVERY_TIME = 300
	BAG_ON_BODY_NUM = 9
	BAG_START = 0
	BAG_END = 35
	BAG_NUM = 36
	BOX_A = 0
	BOX_B = 1
	BOX_C = 2
	BOX_D = 3
	BOX_E = 4
	BOX_NUM = 5
	QIAN_LONG_ID = 37
	RUBBISH_AWARD_COIN = 500
)

type BoxConfig struct {
	Key  int8
	Coin  int64
}

var BoxConfigs = []BoxConfig{
	BoxConfig{1,1000,},
	BoxConfig{2,2000,},
	BoxConfig{4,3000,},
	BoxConfig{8,4000,},
	BoxConfig{16,8000,},
}


func GetOpenedPosNum(level int16) (res int) {
	if level <= 9 {
		return 1
	}
	if level >= 10 && level <= 19 {
		return 2
	}
	if level >= 20 && level <= 29 {
		return 3
	}
	if level >= 30 && level <= 39 {
		return 4
	}
	if level >= 40 && level <= 49 {
		return 5
	}
	if level >= 50 && level <= 59 {
		return 6
	}
	if level >= 60 && level <= 69 {
		return 7
	}
	if level >= 70 && level <= 79 {
		return 8
	}
	if level >= 80 {
		return 9
	}
	panic("GetOpenedPosNum err")
	return 1
}
func SwordSoulMaxLevel(level int16) (res int8) {
	if level <= 20 {
		return 4
	}
	if level >= 21 && level <= 40 {
		return 8
	}
	if level >= 41 && level <= 60 {
		return 12
	}
	if level >= 61 && level <= 80 {
		return 16
	}
	if level >= 81 {
		return 20
	}
	panic("SwordSoulMaxLevel err")
	return 1
}
func GetSwordDrawPriceInIngot(times int64) (res int64) {
	if times >= 1 && times <= 2 {
		return 50
	}
	if times >= 3 && times <= 5 {
		return 75
	}
	if times >= 6 && times <= 10 {
		return 100
	}
	if times >= 11 && times <= 20 {
		return 200
	}
	panic("GetSwordDrawPriceInIngot err")
	return 0
}
func GetSwordDrawPriceInIngotCoin(times int64) (res int64) {
	if times >= 1 && times <= 2 {
		return 4000
	}
	if times >= 3 && times <= 5 {
		return 4000
	}
	if times >= 6 && times <= 10 {
		return 4000
	}
	if times >= 11 && times <= 20 {
		return 4000
	}
	panic("GetSwordDrawPriceInIngotCoin err")
	return 0
}



const (
	TYPE_ATTACK  =  1
	TYPE_DEFENCE  =  2
	TYPE_HEALTH  =  3
	TYPE_SPEED  =  4
	TYPE_CULTIVATION  =  5
	TYPE_HIT_LEVEL  =  6
	TYPE_CRITICAL_LEVEL =  7
	TYPE_BLOCK_LEVEL  =  8
	TYPE_DESTROY_LEVEL  =  9
	TYPE_TENACITY_LEVEL =  10
	TYPE_DODGE_LEVEL =  11
	TYPE_SUNDER_MAX_VALUE =  12
	TYPE_EXP =  13
	TYPE_NUM = 13
)
const (
	QUALITY_NULL =  0
	QUALITY_SPECIAL =  1
	QUALITY_FINE =  2
	QUALITY_EXCELLENT =  3
	QUALITY_LEGEND =  4
	QUALITY_ARTIFACT =  5
	QUALITY_ONLY =  6
	QUALITY_NUM = 7
)
var NextBoxProbability = []float64{float64(40)/100, float64(40)/100, float64(30)/100, float64(40)/100, float64(0)/100, }
var NextBoxProbability_VIP = []float64{float64(40)/100, float64(0)/100, }
// Order: rubbish, exp, green, blue, purple, yello
var SwordSoulProbability = [][]float64{
	{float64(75)/100, float64(0)/100, float64(25)/100, float64(0)/100, float64(0)/100, float64(0)/100, },
	{float64(60)/100, float64(0)/100, float64(35)/100, float64(5)/100, float64(0)/100, float64(0)/100, },
	{float64(39)/100, float64(0)/100, float64(45)/100, float64(15)/100, float64(1)/100, float64(0)/100, },
	{float64(0)/100, float64(20)/100, float64(0)/100, float64(60)/100, float64(20)/100, float64(0)/100, },
	{float64(0)/100, float64(30)/100, float64(0)/100, float64(30)/100, float64(37)/100, float64(3)/100, },
}
var SwordSoulProbability_VIP = [][]float64{
	{float64(0)/100, float64(70)/100, float64(0)/100, float64(0)/100, float64(30)/100, float64(0)/100, },
	{float64(0)/100, float64(50)/100, float64(0)/100, float64(0)/100, float64(47)/100, float64(3)/100, },
}

