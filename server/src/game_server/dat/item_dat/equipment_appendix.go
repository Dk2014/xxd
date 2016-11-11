package item_dat

import (
	"core/fail"
	"core/mysql"
)

var (
	mapEquipmentAppendix map[int32]*EquipmentAppendix
)

type EquipmentAppendix struct {
	Level         int32 // 等级
	Health        int32 // 生命
	Cultivation   int32 // 内力
	Speed         int32 // 速度
	Attack        int32 // 攻击
	Defence       int32 // 防御
	DodgeLevel    int32 // 闪避
	HitLevel      int32 // 命中
	BlockLevel    int32 // 格挡
	CriticalLevel int32 // 暴击
	TenacityLevel int32 // 韧性
	DestroyLevel  int32 // 破击
}

func loadEquipmentAppendix(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM equipment_appendix ORDER BY `level` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iLevel := res.Map("level")
	iHealth := res.Map("health")
	iCultivation := res.Map("cultivation")
	iSpeed := res.Map("speed")
	iAttack := res.Map("attack")
	iDefence := res.Map("defence")
	iDodgeLevel := res.Map("dodge_level")
	iHitLevel := res.Map("hit_level")
	iBlockLevel := res.Map("block_level")
	iCriticalLevel := res.Map("critical_level")
	iTenacityLevel := res.Map("tenacity_level")
	iDestroyLevel := res.Map("destroy_level")

	mapEquipmentAppendix = map[int32]*EquipmentAppendix{}

	var level int32
	for _, row := range res.Rows {
		level = row.Int32(iLevel)

		equipmentAppendix := &EquipmentAppendix{
			Level:         level,
			Health:        row.Int32(iHealth),
			Cultivation:   row.Int32(iCultivation),
			Speed:         row.Int32(iSpeed),
			Attack:        row.Int32(iAttack),
			Defence:       row.Int32(iDefence),
			DodgeLevel:    row.Int32(iDodgeLevel),
			HitLevel:      row.Int32(iHitLevel),
			BlockLevel:    row.Int32(iBlockLevel),
			CriticalLevel: row.Int32(iCriticalLevel),
			TenacityLevel: row.Int32(iTenacityLevel),
			DestroyLevel:  row.Int32(iDestroyLevel),
		}
		mapEquipmentAppendix[level] = equipmentAppendix
	}
}

// 装备生成时附加属性
func GetEquipmentAppendix(level int32) (lower *EquipmentAppendix, upper *EquipmentAppendix) {
	lower = mapEquipmentAppendix[level-1]
	fail.When(lower == nil, "wrong level")
	upper = mapEquipmentAppendix[level]
	fail.When(upper == nil, "wrong level")
	return lower, upper
}

// 装备重铸生成时附加属性
func GetEquipmentRecastAppendix(level int32) (lower *EquipmentAppendix, upper *EquipmentAppendix) {
	lower = mapEquipmentAppendix[level]
	fail.When(lower == nil, "wrong level")
	upper = mapEquipmentAppendix[level+2]
	fail.When(upper == nil, "wrong level")
	return lower, upper
}

//传入参数为特殊物品ID的负值
func GetSpecialEquipmentAppendix(level int32) *EquipmentAppendix {
	// 传入后再转为负值
	v, ok := mapEquipmentAppendix[-level]
	fail.When(!ok, "GetSpecialEquipmentAppendix wrong level")
	return v
}
