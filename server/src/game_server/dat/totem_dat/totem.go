package totem_dat

import (
	"core/fail"
	"core/mysql"
	"fmt"
)

var (
	mapTotemQuality    map[int16]int8   //ID -> quality
	mapTotemsByQuality map[int8][]int16 //quality -> []ID
	mapTotem           map[int16]*Totem //ID -> totem
)

func Load(db *mysql.Connection) {
	loadTotem(db)
	loadTotemLevelInfo(db)
	loadTotemSkill(db)
	loadTotemCallCostConfig(db)
}

type Totem struct {
	Id      int16  // 主键
	Name    string // 阵印名称
	Quality int8   // 阵印品质
}

func (totem *Totem) IsRareItem() bool {
	return totem.Quality == TOTEM_QUALITY_GOLDEN
}

func loadTotem(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM totem ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iId := res.Map("id")
	iName := res.Map("name")
	iQuality := res.Map("quality")

	var pri_id int16
	var quality int8
	mapTotemQuality = map[int16]int8{}
	mapTotemsByQuality = map[int8][]int16{}
	mapTotem = map[int16]*Totem{}
	for _, row := range res.Rows {
		pri_id = row.Int16(iId)
		quality = row.Int8(iQuality)

		mapTotemQuality[pri_id] = quality
		mapTotemsByQuality[quality] = append(mapTotemsByQuality[quality], pri_id)
		mapTotem[pri_id] = &Totem{
			Id:      pri_id,
			Name:    row.Str(iName),
			Quality: quality,
		}
	}
}

// ############### 对外接口实现 coding here ###############

func GetTotemQualiyById(totemId int16) int8 {
	if quality, ok := mapTotemQuality[totemId]; !ok {
		fail.When(true, fmt.Sprintf("没有找到图腾配置 %d", totemId))
	} else {
		return quality
	}
	panic("unreachable")
}

func GetTotemsByQuality(quality int8) []int16 {
	totems := mapTotemsByQuality[quality]
	fail.When(len(totems) < 1, fmt.Sprintf("没有找到该品质的阵印 %d", quality))
	return totems
}

func GetTotemById(id int16) (totem *Totem) {
	totem = mapTotem[id]
	fail.When(totem == nil, fmt.Sprintf("找不到阵印 %d ", id))
	return totem
}
