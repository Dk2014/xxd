package ghost_dat

import (
	"core/mysql"
	"math"
)

var (
	mapGhostLevel       []*GhostLevel
	mapGhostFruitCost   []int32
	mapGhostFruitExpMod []int64
)

type GhostLevel struct {
	Level        int16 // 魂侍等级
	Exp          int64 // 升级所需经验
	NeedFruitNum int32 // 所需影界果实数量
	MinAddExp    int64 // 最小经验加值
	MaxAddExp    int64 // 最大经验加值
}

func loadGhostLevel(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM ghost_level ORDER BY `level` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iLevel := res.Map("level")
	iExp := res.Map("exp")
	iNeedFruitNum := res.Map("need_fruit_num")
	iMinAddExp := res.Map("min_add_exp")
	iMaxAddExp := res.Map("max_add_exp")

	mapGhostLevel = []*GhostLevel{}
	for _, row := range res.Rows {
		mapGhostLevel = append(mapGhostLevel, &GhostLevel{
			Level:        row.Int16(iLevel),
			Exp:          row.Int64(iExp),
			NeedFruitNum: row.Int32(iNeedFruitNum),
			MinAddExp:    row.Int64(iMinAddExp),
			MaxAddExp:    row.Int64(iMaxAddExp),
		})
	}

	mapGhostFruitCost = make([]int32, len(mapGhostLevel))
	mapGhostFruitExpMod = make([]int64, len(mapGhostLevel))
	preExp := int64(0)
	for i := 0; i < len(mapGhostLevel); i++ {
		mapGhostFruitExpMod[i] = preExp
		level := mapGhostLevel[i]
		if level.Exp <= preExp {
			mapGhostFruitCost[i] = 0
			preExp -= level.Exp
		} else {
			aveExp := (level.MinAddExp + level.MaxAddExp) / 2
			leftExp := level.Exp - preExp
			times := int32(math.Ceil(float64(leftExp) / float64(aveExp)))
			mapGhostFruitCost[i] = times * level.NeedFruitNum
			preExp = int64(times)*aveExp - leftExp
		}
		if i > 0 {
			mapGhostFruitCost[i] += mapGhostFruitCost[i-1]
		}
	}
}

// ############### 对外接口实现 coding here ###############

func GetGhostLevel(level int16) (ghostLevel *GhostLevel) {
	ghostLevel = mapGhostLevel[level-1]
	return ghostLevel
}

func GetGhostFruitCost(level int16, exp int64) int32 {
	totalFruit := int32(0)
	if level > 1 {
		totalFruit = mapGhostFruitCost[level-2]
	} else if level == 1 {
		return 1 // 只升一级的情况固定1个返还
	}
	expMod := mapGhostFruitExpMod[level-1]
	levelInfo := GetGhostLevel(level)
	aveExp := (levelInfo.MinAddExp + levelInfo.MaxAddExp) / 2
	if exp > expMod {
		exp -= expMod
		totalFruit += int32(exp/aveExp) * levelInfo.NeedFruitNum
	}
	return totalFruit
}

//见策划档 魂侍 魂侍被动技
func GetGhostShieldDiscount(star int8) (discount float64) {
	if star >= 5 {
		return 1.0
	}
	if star >= 4 {
		return 0.75
	}
	if star >= 3 {
		return 0.5
	}
	if star >= 2 {
		return 0.25
	}
	panic("魂侍护盾：错误的魂侍星级")
}
