package enemy_deploy_dat

import (
	"core/fail"
	"core/mysql"
)

var (
	mapEnemyDeployForm map[int8]map[int32][][]int32 // map[BattleType][ParentId][form_idx][0~14]
)

func Load(db *mysql.Connection) {
	loadEnemyDeployForm(db)
}

type EnemyDeployForm struct {
	Id         int32 //
	ParentId   int32 // 关联此阵法的某表唯一ID
	BattleType int8  // 战场类型(0--关卡;)
	Pos1       int32 // 位置1上的敌人ID
	Pos2       int32 // 位置2上的敌人ID
	Pos3       int32 // 位置3上的敌人ID
	Pos4       int32 // 位置4上的敌人ID
	Pos5       int32 // 位置5上的敌人ID
	Pos6       int32 // 位置6上的敌人ID
	Pos7       int32 // 位置7上的敌人ID
	Pos8       int32 // 位置8上的敌人ID
	Pos9       int32 // 位置9上的敌人ID
	Pos10      int32 // 位置10上的敌人ID
	Pos11      int32 // 位置11上的敌人ID
	Pos12      int32 // 位置12上的敌人ID
	Pos13      int32 // 位置13上的敌人ID
	Pos14      int32 // 位置14上的敌人ID
	Pos15      int32 // 位置15上的敌人ID
}

func loadEnemyDeployForm(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM enemy_deploy_form ORDER BY `battle_type`, `parent_id`"), -1)
	if err != nil {
		panic(err)
	}

	iParentId := res.Map("parent_id")
	iBattleType := res.Map("battle_type")
	iPos1 := res.Map("pos1")
	iPos2 := res.Map("pos2")
	iPos3 := res.Map("pos3")
	iPos4 := res.Map("pos4")
	iPos5 := res.Map("pos5")
	iPos6 := res.Map("pos6")
	iPos7 := res.Map("pos7")
	iPos8 := res.Map("pos8")
	iPos9 := res.Map("pos9")
	iPos10 := res.Map("pos10")
	iPos11 := res.Map("pos11")
	iPos12 := res.Map("pos12")
	iPos13 := res.Map("pos13")
	iPos14 := res.Map("pos14")
	iPos15 := res.Map("pos15")

	var parent_id int32
	var battle_type int8
	var ok bool

	mapEnemyDeployForm = make(map[int8]map[int32][][]int32)

	for _, row := range res.Rows {
		parent_id = row.Int32(iParentId)
		battle_type = row.Int8(iBattleType)

		if _, ok = mapEnemyDeployForm[battle_type]; !ok {
			mapEnemyDeployForm[battle_type] = make(map[int32][][]int32)
		}

		mapEnemyDeployForm[battle_type][parent_id] = append(mapEnemyDeployForm[battle_type][parent_id], []int32{
			row.Int32(iPos1),
			row.Int32(iPos2),
			row.Int32(iPos3),
			row.Int32(iPos4),
			row.Int32(iPos5),
			row.Int32(iPos6),
			row.Int32(iPos7),
			row.Int32(iPos8),
			row.Int32(iPos9),
			row.Int32(iPos10),
			row.Int32(iPos11),
			row.Int32(iPos12),
			row.Int32(iPos13),
			row.Int32(iPos14),
			row.Int32(iPos15),
		})
	}
}

// ############### 对外接口实现 coding here ###############

func GetEnemyDeployForm(battleType int8, parentId int32) [][]int32 {
	form, ok := mapEnemyDeployForm[battleType][parentId]
	fail.When(!ok, "can't get enemy form")
	return form
}
