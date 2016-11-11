package ghost_dat

import (
	"core/fail"
	"core/mysql"
	"fmt"
)

var (
	mapGhostStar map[int8]map[int8]*GhostStar // quality -> star -> GHostStar
)

type GhostStar struct {
	Star            int8  // 星级
	NeedFragmentNum int16 // 需要碎片数量
	Growth          int16 // 成长值
	Quality         int8  // 颜色
	Costs           int64 //费用
	Health          int32 // 生命
	Attack          int32 // 攻击
	Defence         int32 // 防御
}

func loadGhostStar(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM ghost_star ORDER BY `star` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iStar := res.Map("star")
	iNeedFragmentNum := res.Map("need_fragment_num")
	iGrowth := res.Map("growth")
	iQuality := res.Map("quality")
	iCosts := res.Map("costs")
	iHealth := res.Map("health")
	iAttack := res.Map("attack")
	iDefence := res.Map("defence")

	mapGhostStar = map[int8]map[int8]*GhostStar{}
	var quality, star int8
	for _, row := range res.Rows {
		star = row.Int8(iStar)
		quality = row.Int8(iQuality)
		if mapGhostStar[quality] == nil {
			mapGhostStar[quality] = map[int8]*GhostStar{}
		}
		mapGhostStar[quality][star] = &GhostStar{
			Star:            row.Int8(iStar),
			NeedFragmentNum: row.Int16(iNeedFragmentNum),
			Growth:          row.Int16(iGrowth),
			Quality:         row.Int8(iQuality),
			Costs:           row.Int64(iCosts),
			Health:          row.Int32(iHealth),
			Attack:          row.Int32(iAttack),
			Defence:         row.Int32(iDefence),
		}
	}
}

// ############### 对外接口实现 coding here ###############
func GetGhostStar(quality, star int8) (ghostStar *GhostStar) {
	ghostStarByQuality := mapGhostStar[quality]
	fail.When(ghostStarByQuality == nil, fmt.Sprintf("找不到数值 quality %d ", quality))

	ghostStar = ghostStarByQuality[star]
	fail.When(ghostStar == nil, fmt.Sprintf("找不到数值 quality %d star %d", quality, star))
	return ghostStar
}
