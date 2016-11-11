package clique_dat

import (
	"core/mysql"
)

var (
	mapCliqueBuilding map[int32]*CliqueBuilding
)

type CliqueBuilding struct {
	//Id      int32  // 标示ID
	Name    string // 建筑名称
	Biaoshi string // 建筑标识
	Desc    string // 建筑描述
	Order   int16  // 优先权排序
}

func loadCliqueBuilding(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM clique_building ORDER BY `order` "), -1)
	if err != nil {
		panic(err)
	}

	iId := res.Map("id")
	iName := res.Map("name")
	iBiaoshi := res.Map("biaoshi")
	iDesc := res.Map("desc")
	iOrder := res.Map("order")

	var pri_id int32
	mapCliqueBuilding = map[int32]*CliqueBuilding{}
	for _, row := range res.Rows {
		pri_id = row.Int32(iId)
		mapCliqueBuilding[pri_id] = &CliqueBuilding{
			//Id:      pri_id,
			Name:    row.Str(iName),
			Biaoshi: row.Str(iBiaoshi),
			Desc:    row.Str(iDesc),
			Order:   row.Int16(iOrder),
		}
	}
}

// ############### 对外接口实现 coding here ###############

func GetCliqueBuildingByID(id int32) (*CliqueBuilding, bool) {
	if content, ok := mapCliqueBuilding[id]; ok {
		return content, true
	}
	return nil, false
}
