package ghost_dat

import (
	"core/fail"
	"core/mysql"
)

var (
	mapGhostBaptize map[int8]*GhostBaptize
)

type GhostBaptize struct {
	Id            int8 // 魂侍洗炼ID
	Star          int8 // 魂侍星级
	MaxAddGrowth  int8 // 最大增加成长值
	Probablity1   int8 // 概率1
	MinAddGrowth1 int8 // 最小随机添加成长值1
	MaxAddGrowth1 int8 // 最大随机添加成长值1
	Probablity2   int8 // 概率2
	MinAddGrowth2 int8 // 最小随机添加成长值2
	MaxAddGrowth2 int8 // 最大随机添加成长值2
}

func loadGhostBaptize(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM ghost_baptize ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iId := res.Map("id")
	iStar := res.Map("star")
	iMaxAddGrowth := res.Map("max_add_growth")
	iProbablity1 := res.Map("probablity1")
	iMinAddGrowth1 := res.Map("min_add_growth1")
	iMaxAddGrowth1 := res.Map("max_add_growth1")
	iProbablity2 := res.Map("probablity2")
	iMinAddGrowth2 := res.Map("min_add_growth2")
	iMaxAddGrowth2 := res.Map("max_add_growth2")

	var pri_id int8
	mapGhostBaptize = map[int8]*GhostBaptize{}
	for _, row := range res.Rows {
		pri_id = row.Int8(iStar)
		mapGhostBaptize[pri_id] = &GhostBaptize{
			Id:            row.Int8(iId),
			Star:          pri_id,
			MaxAddGrowth:  row.Int8(iMaxAddGrowth),
			Probablity1:   row.Int8(iProbablity1),
			MinAddGrowth1: row.Int8(iMinAddGrowth1),
			MaxAddGrowth1: row.Int8(iMaxAddGrowth1),
			Probablity2:   row.Int8(iProbablity2),
			MinAddGrowth2: row.Int8(iMinAddGrowth2),
			MaxAddGrowth2: row.Int8(iMaxAddGrowth2),
		}
	}
}

// ############### 对外接口实现 coding here ###############
func GetGhostBaptize(star int8) (ghostBaptize *GhostBaptize) {
	ghostBaptize = mapGhostBaptize[star]
	fail.When(ghostBaptize == nil, "ghost_baptize::GetGhostBaptize 模板数据不存在")
	return ghostBaptize
}
