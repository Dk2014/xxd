// package item_dat

// import (
// 	"core/fail"
// 	"core/mysql"
// )

// var (
// 	mapEquipmentRefine      map[int8][]*EquipmentRefine
// 	mapEquipmentRefineLevel map[int8]map[int8]*EquipmentRefineLevel
// )

// type EquipmentRefine struct {
// 	Level              int32 // 等级下限
// 	Quality            int8  // 品质
// 	FragmentNum        int16 // 需要部位碎片数量
// 	BlueCrystalNum     int16 // 需要蓝色结晶数量
// 	PurpleCrystalNum   int16 // 需要紫色结晶数量
// 	GoldenCrystalNum   int16 // 需要金色结晶数量
// 	OrangeCrystalNum   int16 // 需要橙色结晶数量
// 	Level1ConsumeCoin  int64 // 精练到1级消耗的铜钱
// 	Level2ConsumeCoin  int64 // 精练到2级消耗的铜钱
// 	Level3ConsumeCoin  int64 // 精练到3级消耗的铜钱
// 	Level4ConsumeCoin  int64 // 精练到4级消耗的铜钱
// 	Level5ConsumeCoin  int64 // 精练到5级消耗的铜钱
// 	Level6ConsumeCoin  int64 // 精练到6级消耗的铜钱
// 	Level7ConsumeCoin  int64 // 精练到7级消耗的铜钱
// 	Level8ConsumeCoin  int64 // 精练到8级消耗的铜钱
// 	Level9ConsumeCoin  int64 // 精练到9级消耗的铜钱
// 	Level10ConsumeCoin int64 // 精练到10级消耗的铜钱
// }

// type EquipmentRefineLevel struct {
// 	Level       int8  // 精练级别
// 	Quality     int8  // 品质
// 	Probability int8  // 精练成功概率
// 	GainPct     int32 // 增益百分比
// }

// func loadEquipmentRefine(db *mysql.Connection) {
// 	res, err := db.ExecuteFetch([]byte("SELECT * FROM equipment_refine ORDER BY `quality`,`level` ASC"), -1)
// 	if err != nil {
// 		panic(err)
// 	}

// 	iLevel := res.Map("level")
// 	iQuality := res.Map("quality")
// 	iFragmentNum := res.Map("fragment_num")
// 	iBlueCrystalNum := res.Map("blue_crystal_num")
// 	iPurpleCrystalNum := res.Map("purple_crystal_num")
// 	iGoldenCrystalNum := res.Map("golden_crystal_num")
// 	iOrangeCrystalNum := res.Map("orange_crystal_num")
// 	iLevel1ConsumeCoin := res.Map("level1_consume_coin")
// 	iLevel2ConsumeCoin := res.Map("level2_consume_coin")
// 	iLevel3ConsumeCoin := res.Map("level3_consume_coin")
// 	iLevel4ConsumeCoin := res.Map("level4_consume_coin")
// 	iLevel5ConsumeCoin := res.Map("level5_consume_coin")
// 	iLevel6ConsumeCoin := res.Map("level6_consume_coin")
// 	iLevel7ConsumeCoin := res.Map("level7_consume_coin")
// 	iLevel8ConsumeCoin := res.Map("level8_consume_coin")
// 	iLevel9ConsumeCoin := res.Map("level9_consume_coin")
// 	iLevel10ConsumeCoin := res.Map("level10_consume_coin")

// 	mapEquipmentRefine = map[int8][]*EquipmentRefine{}
// 	for _, row := range res.Rows {
// 		quality := row.Int8(iQuality)
// 		if mapEquipmentRefine[quality] == nil {
// 			mapEquipmentRefine[quality] = []*EquipmentRefine{}
// 		}
// 		mapEquipmentRefine[quality] = append(mapEquipmentRefine[quality], &EquipmentRefine{
// 			Level:              row.Int32(iLevel),
// 			Quality:            quality,
// 			FragmentNum:        row.Int16(iFragmentNum),
// 			BlueCrystalNum:     row.Int16(iBlueCrystalNum),
// 			PurpleCrystalNum:   row.Int16(iPurpleCrystalNum),
// 			GoldenCrystalNum:   row.Int16(iGoldenCrystalNum),
// 			OrangeCrystalNum:   row.Int16(iOrangeCrystalNum),
// 			Level1ConsumeCoin:  row.Int64(iLevel1ConsumeCoin),
// 			Level2ConsumeCoin:  row.Int64(iLevel2ConsumeCoin),
// 			Level3ConsumeCoin:  row.Int64(iLevel3ConsumeCoin),
// 			Level4ConsumeCoin:  row.Int64(iLevel4ConsumeCoin),
// 			Level5ConsumeCoin:  row.Int64(iLevel5ConsumeCoin),
// 			Level6ConsumeCoin:  row.Int64(iLevel6ConsumeCoin),
// 			Level7ConsumeCoin:  row.Int64(iLevel7ConsumeCoin),
// 			Level8ConsumeCoin:  row.Int64(iLevel8ConsumeCoin),
// 			Level9ConsumeCoin:  row.Int64(iLevel9ConsumeCoin),
// 			Level10ConsumeCoin: row.Int64(iLevel10ConsumeCoin),
// 		})
// 	}
// }

// func loadEquipmentRefineLevel(db *mysql.Connection) {
// 	res, err := db.ExecuteFetch([]byte("SELECT * FROM equipment_refine_level ORDER BY `id` ASC"), -1)
// 	if err != nil {
// 		panic(err)
// 	}

// 	iLevel := res.Map("level")
// 	iQuality := res.Map("quality")
// 	iProbability := res.Map("probability")
// 	iGainPct := res.Map("gain_pct")

// 	mapEquipmentRefineLevel = map[int8]map[int8]*EquipmentRefineLevel{}
// 	for _, row := range res.Rows {

// 		quality := row.Int8(iQuality)
// 		if mapEquipmentRefineLevel[quality] == nil {
// 			mapEquipmentRefineLevel[quality] = map[int8]*EquipmentRefineLevel{}
// 		}

// 		level := row.Int8(iLevel)
// 		mapEquipmentRefineLevel[quality][level] = &EquipmentRefineLevel{
// 			Level:       level,
// 			Quality:     quality,
// 			Probability: row.Int8(iProbability),
// 			GainPct:     row.Int32(iGainPct),
// 		}
// 	}
// }

// // ############### 对外接口实现 coding here ###############

// func GetEquipmentRefine(quality int8, level int32) (ret *EquipmentRefine) {
// 	equipmentRefines := mapEquipmentRefine[quality]
// 	fail.When(equipmentRefines == nil, "no this quality EquipmentRefines")
// 	for i, equipmentRefine := range equipmentRefines {
// 		if equipmentRefine.Level > level {
// 			ret = equipmentRefines[i-1]
// 			break
// 		}
// 	}
// 	return ret
// }

// func GetEquipmentRefineLevel(quality int8, level int8) *EquipmentRefineLevel {
// 	equipmentRefineLevelMap := mapEquipmentRefineLevel[quality]
// 	fail.When(equipmentRefineLevelMap == nil, "no this quality EquipmentRefineLevelMap")
// 	equipmentRefineLevel := equipmentRefineLevelMap[level]
// 	fail.When(equipmentRefineLevel == nil, "no this quality EquipmentRefineLevel")
// 	return equipmentRefineLevel
// }
package item_dat

import (
	"core/fail"
	"core/mysql"
	"fmt"
)

var (
	mapEquipmentRefineNew map[string]*EquipmentRefineNew
)

type EquipmentRefineNew struct {
	EquType    int8  // 装备类型(武器,防具等)
	EquKind    int8  // 装备种类(鹰扬,玄武等)
	BaseVal    int32 // 装备基础强度
	BasePrice  int32 // 装备基础价格
	IncreVal   int32 // 强化单位提升属性
	IncrePrice int32 // 强化单位价格
}

func loadEquipmentRefineNew(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM equipment_refine_new ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iEquType := res.Map("equ_type")
	iEquKind := res.Map("equ_kind")
	iBaseVal := res.Map("base_val")
	iBasePrice := res.Map("base_price")
	iIncreVal := res.Map("incre_val")
	iIncrePrice := res.Map("incre_price")

	var pri_id string
	mapEquipmentRefineNew = map[string]*EquipmentRefineNew{}
	for _, row := range res.Rows {
		pri_id = fmt.Sprintf("%d_%d", row.Int8(iEquType), row.Int8(iEquKind))
		mapEquipmentRefineNew[pri_id] = &EquipmentRefineNew{
			EquType:    row.Int8(iEquType),
			EquKind:    row.Int8(iEquKind),
			BaseVal:    row.Int32(iBaseVal),
			BasePrice:  row.Int32(iBasePrice),
			IncreVal:   row.Int32(iIncreVal),
			IncrePrice: row.Int32(iIncrePrice),
		}
	}
}

// ############### 对外接口实现 coding here ###############
func GetEquipmentRefineDat(typeId, kindId int8) *EquipmentRefineNew {
	key := fmt.Sprintf("%d_%d", typeId, kindId)
	finalRefineDat := mapEquipmentRefineNew[key]
	fail.When(finalRefineDat == nil, "can't find the refine dat")
	return finalRefineDat
}
