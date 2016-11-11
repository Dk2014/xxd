package item_dat

import (
	"core/mysql"
)

var (
	mapEquipmentResonance   []*EquipmentResonance
	emptyEquipmentResonance = &EquipmentResonance{}
)

type EquipmentResonance struct {
	RequireLevel int16   // 要求等级
	Health       int     // 生命 - health
	Attack       float64 // 普攻 - attack
	Defence      float64 // 普防 - defence
}

func loadEquipmentResonance(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM equipment_resonance ORDER BY `require_level` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iRequireLevel := res.Map("require_level")
	iHealth := res.Map("health")
	iAttack := res.Map("attack")
	iDefence := res.Map("defence")

	mapEquipmentResonance = nil
	for _, row := range res.Rows {
		mapEquipmentResonance = append(mapEquipmentResonance, &EquipmentResonance{
			RequireLevel: row.Int16(iRequireLevel),
			Health:       int(row.Int32(iHealth)),
			Attack:       float64(row.Int32(iAttack)),
			Defence:      float64(row.Int32(iDefence)),
		})
	}
}

func GetEquipmentResonance(level int16) (dat *EquipmentResonance) {
	dat = emptyEquipmentResonance
	for _, resonanceDat := range mapEquipmentResonance {
		if level >= resonanceDat.RequireLevel {
			dat = resonanceDat
		} else {
			break
		}
	}
	return dat
}
