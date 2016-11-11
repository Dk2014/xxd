package sword_soul

const (
	//FIXME 常量移动到 dat

	DRAW_OUT_NULL       = 0 // 拔剑没有获得东西
	DRAW_OUT_SWORD_SOUL = 1 // 拔剑获得剑心
	DRAW_OUT_QIANLONG   = 2 // 拔剑获得 QIANLONG 剑心

	// Deprecated
	AWARD_COIN = 500 // 拔剑没有获得东西，奖励的铜钱

	BAG_FULL = -1 // 背包已经满

	EQUIPPED = 1 // 已经装备
	INBAG    = 0 // 未装备,在背包

	FULL_LEVEL = 20 // 剑心满级

	FIRST_DRAW_AWARD_SWORD_SOUL = 2 //头两次拔剑奖励的剑心
)

// 拔剑开出下一个箱子的概率
// Deprecated
var NextBoxProbability = []float64{0.4, 0.4, 0.3, 0.4, 0}

// 铜钱拔剑开出 QIANLONG 到概率
// Deprecated
var QianlongProbabilityByCoin = []float64{0, 0, 0, 0.2, 0.3}

// 铜钱拔剑获得剑心的概率
// Deprecated
var SwordSoulProbabilityWithQualityByCoin = []map[int8]float64{
	{ // A
		0: 0.75, // quality: probability
		2: 0.25,
		3: 0,
		4: 0,
		5: 0,
	},
	{ // B
		0: 0.6,
		2: 0.35,
		3: 0.05,
		4: 0,
		5: 0,
	},
	{ // C
		0: 0.39,
		2: 0.45,
		3: 0.15,
		4: 0.1,
		5: 0,
	},
	{ // D
		0: 0,
		2: 0,
		3: 0.6,
		4: 0.2,
		5: 0,
	},
	{ // E
		0: 0,
		2: 0,
		3: 0.3,
		4: 0.3,
		5: 0.1,
	},
}
