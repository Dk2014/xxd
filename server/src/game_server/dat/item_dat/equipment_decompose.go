package item_dat

import (
	"core/mysql"
)

var (
	mapEquipmentDecomposes map[int8][]*EquipmentDecompose // quality -> []*EquipmentDecompose
)

type EquipmentDecompose struct {
	Level         int32 // 等级下限
	FragmentNum   int16 // 获得部位碎片数量
	CrystalNum    int16 // 获得结晶数量
	CrystalId     int16 // 获得结晶ID
	DragonBall    int16 // 获得龙珠ID
	DragonBallNum int16 // 获得数量
}

func loadEquipmentDecompose(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM equipment_decompose ORDER BY `quality`, `level` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iQuality := res.Map("quality")
	iLevel := res.Map("level")
	iFragmentNum := res.Map("fragment_num")
	iCrystalNum := res.Map("crystal_num")
	iCrystalId := res.Map("crystal_id")
	iDragonBall := res.Map("dragon_ball")
	iDragonBallNum := res.Map("dragon_ball_num")

	var quality int8
	mapEquipmentDecomposes = map[int8][]*EquipmentDecompose{}
	for _, row := range res.Rows {
		quality = row.Int8(iQuality)
		mapEquipmentDecomposes[quality] = append(mapEquipmentDecomposes[quality], &EquipmentDecompose{
			Level:         row.Int32(iLevel),
			FragmentNum:   row.Int16(iFragmentNum),
			CrystalNum:    row.Int16(iCrystalNum),
			CrystalId:     row.Int16(iCrystalId),
			DragonBall:    row.Int16(iDragonBall),
			DragonBallNum: row.Int16(iDragonBallNum),
		})
	}
}

// ############### 对外接口实现 coding here ###############

func GetEquipmentDecompose(quality int8, level int32) (equipmentDecompose *EquipmentDecompose) {
	decomposeDat := mapEquipmentDecomposes[quality]
	for _, mapEquipmentDecompose := range decomposeDat {
		if mapEquipmentDecompose.Level >= level {
			equipmentDecompose = mapEquipmentDecompose
			break
		}
	}
	return equipmentDecompose
}
