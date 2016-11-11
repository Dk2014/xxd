package item_dat

import (
	"core/fail"
	"core/mysql"
)

var (
	mapEquipmentRecast map[int8][]*EquipmentRecast
)

type EquipmentRecast struct {
	Level            int32 // 等级下限
	Quality          int8  // 品质
	FragmentNum      int16 // 需要部位碎片数量
	BlueCrystalNum   int16 // 需要蓝色结晶数量
	PurpleCrystalNum int16 // 需要紫色结晶数量
	GoldenCrystalNum int16 // 需要金色结晶数量
	OrangeCrystalNum int16 // 需要橙色结晶数量
	ConsumeCoin      int64 // 消耗的铜钱
}

func loadEquipmentRecast(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM equipment_recast ORDER BY `quality`,`level` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iLevel := res.Map("level")
	iQuality := res.Map("quality")
	iFragmentNum := res.Map("fragment_num")
	iBlueCrystalNum := res.Map("blue_crystal_num")
	iPurpleCrystalNum := res.Map("purple_crystal_num")
	iGoldenCrystalNum := res.Map("golden_crystal_num")
	iOrangeCrystalNum := res.Map("orange_crystal_num")
	iConsumeCoin := res.Map("consume_coin")

	mapEquipmentRecast = map[int8][]*EquipmentRecast{}
	for _, row := range res.Rows {

		quality := row.Int8(iQuality)
		if mapEquipmentRecast[quality] == nil {
			mapEquipmentRecast[quality] = []*EquipmentRecast{}
		}

		mapEquipmentRecast[quality] = append(mapEquipmentRecast[quality], &EquipmentRecast{
			Level:            row.Int32(iLevel),
			Quality:          quality,
			FragmentNum:      row.Int16(iFragmentNum),
			BlueCrystalNum:   row.Int16(iBlueCrystalNum),
			PurpleCrystalNum: row.Int16(iPurpleCrystalNum),
			GoldenCrystalNum: row.Int16(iGoldenCrystalNum),
			OrangeCrystalNum: row.Int16(iOrangeCrystalNum),
			ConsumeCoin:      row.Int64(iConsumeCoin),
		})
	}
}

// ############### 对外接口实现 coding here ###############

func GetEquipmentRecast(quality int8, level int32) (ret *EquipmentRecast) {
	equipmentRecasts := mapEquipmentRecast[quality]
	fail.When(equipmentRecasts == nil, "no this quality EquipmentRecasts")
	for i, equipmentRecast := range equipmentRecasts {
		if equipmentRecast.Level > level {
			ret = equipmentRecasts[i-1]
			break
		}
	}
	return ret
}
