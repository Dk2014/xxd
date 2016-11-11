package chest_dat

import (
	"core/mysql"
)

var (
	mapFateBox                 map[int32]*FateBox
	mapMissionLevelIdToFateBox map[int32]*FateBox
)

type FateBox struct {
	Id          int32 // id
	Level       int16 // 要求等级
	RequireLock int32 // 命锁宝箱权值
	AwardLock   int32 // 奖励命锁宝箱权值
	FixedProp   int16 //固定道具
}

func loadFateBox(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte(`select 
	f.id as id ,f.level as level, f.require_lock as require_lock, f.award_lock as award_lock,  m.id as mission_level_id, f.fixed_prop as fixed_prop 
	from fate_box f 
	left join mission_level m 
	on m.parent_type=14 and m.parent_id=f.id;`),
		-1)
	if err != nil {
		panic(err)
	}

	iId := res.Map("id")
	iLevel := res.Map("level")
	iRequireLock := res.Map("require_lock")
	iAwardLock := res.Map("award_lock")
	iMissionLevelId := res.Map("mission_level_id")
	iFixedProp := res.Map("fixed_prop")

	var pri_id int32
	mapFateBox = map[int32]*FateBox{}
	mapMissionLevelIdToFateBox = map[int32]*FateBox{}
	for _, row := range res.Rows {
		pri_id = row.Int32(iId)
		fateBox := &FateBox{
			Id:          pri_id,
			Level:       row.Int16(iLevel),
			RequireLock: row.Int32(iRequireLock),
			AwardLock:   row.Int32(iAwardLock),
			FixedProp:   row.Int16(iFixedProp),
		}
		mapFateBox[pri_id] = fateBox
		mapMissionLevelIdToFateBox[row.Int32(iMissionLevelId)] = fateBox
	}
}

// ############### 对外接口实现 coding here ###############

func GetFateBoxById(id int32) *FateBox {
	return mapFateBox[id]
}

func GetFateBoxByMissionLevelId(missionLevelId int32) *FateBox {
	return mapMissionLevelIdToFateBox[missionLevelId]
}
