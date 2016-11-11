package ghost_dat

import (
	"core/fail"
	"core/mysql"
)

func Load(db *mysql.Connection) {
	loadGhost(db)
	loadGhostLevel(db)
	loadGhostStar(db)
	loadGhostTrainPrice(db)
	loadGhostUpgradePrice(db)
	loadGHostSkillTrainPrice(db)
	loadGhostRelationship(db)
	loadGhostBaptize(db)
}

var (
	mapGhost            map[int16]*Ghost
	mapGhostsWithTownId map[int8][]*Ghost
)

type Ghost struct {
	Id         int16  // 主键
	Name       string //魂侍名
	TownId     int8   // 城镇id（影界id）
	FragmentId int16  // 对应碎片物品id
	InitStar   int8   // 初始星级
	UniqueKey  int16  // 每个影界中魂侍的唯一标记
	Health     int32  // 生命
	Attack     int32  // 攻击
	Defence    int32  // 防御
	Quality    int8   //品质
}

type GhostAddData struct {
	Health  int32 // 生命
	Attack  int32 // 攻击
	Defence int32 // 防御
	Growth  int16 // 成长值
}

func loadGhost(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM ghost ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iId := res.Map("id")
	iName := res.Map("name")
	iTownId := res.Map("town_id")
	iFragmentId := res.Map("fragment_id")
	iInitStar := res.Map("init_star")
	iUniqueKey := res.Map("unique_key")
	iHealth := res.Map("health")
	iAttack := res.Map("attack")
	iDefence := res.Map("defence")
	iQuality := res.Map("quality")

	var pri_id int16
	mapGhost = map[int16]*Ghost{}
	mapGhostsWithTownId = map[int8][]*Ghost{}
	for _, row := range res.Rows {
		pri_id = row.Int16(iId)

		ghost := &Ghost{
			Id:         pri_id,
			Name:       row.Str(iName),
			TownId:     row.Int8(iTownId),
			FragmentId: row.Int16(iFragmentId),
			InitStar:   row.Int8(iInitStar),
			UniqueKey:  row.Int16(iUniqueKey),
			Health:     row.Int32(iHealth),
			Attack:     row.Int32(iAttack),
			Defence:    row.Int32(iDefence),
			Quality:    row.Int8(iQuality),
		}

		mapGhost[pri_id] = ghost

		if mapGhostsWithTownId[row.Int8(iTownId)] == nil {
			mapGhostsWithTownId[row.Int8(iTownId)] = []*Ghost{}
		}

		mapGhostsWithTownId[row.Int8(iTownId)] = append(mapGhostsWithTownId[row.Int8(iTownId)], ghost)
	}
}

// ############### 对外接口实现 coding here ###############

// 根据id取魂侍信息
func GetGhost(id int16) (ghost *Ghost) {
	ghost = mapGhost[id]
	fail.When(ghost == nil, "GetGhostInfoForId 模板数据不存在")
	return ghost
}

func GetGhostsByTownId(townId int8) (ghosts []*Ghost) {
	ghosts = mapGhostsWithTownId[townId]
	fail.When(ghosts == nil, "wrong town id")
	return ghosts
}
