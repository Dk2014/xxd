// 角色配置信息
package role_dat

import (
	"core/mysql"
)

var (
	g_RoleInfos map[int8]*RoleInfo
)

// 角色信息
type RoleInfo struct {
	ID          int8
	Name        string // 角色名称
	Type        int8   // 类型：1.主角，2.伙伴
	IsSpecial   int8   // 是否特殊伙伴 0不是 1是
	SkillID1    int16  // 默认绝招1
	SkillID2    int16  // 默认绝招2
	Level       int16  // 初始等级
	MissionLock int32  // 解锁副本权值
}

func loadRoleInfos(db *mysql.Connection) {
	sql := "select * from role"
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}

	g_RoleInfos = make(map[int8]*RoleInfo)

	iId := res.Map("id")
	iName := res.Map("name")
	iType := res.Map("type")
	iIsSpecial := res.Map("is_special")
	iIskillId1 := res.Map("skill_id1")
	iIskillId2 := res.Map("skill_id2")
	iLevel := res.Map("buddy_level")
	iMissionLock := res.Map("mission_lock")

	for _, row := range res.Rows {
		infoID := row.Int8(iId)
		g_RoleInfos[infoID] = &RoleInfo{
			ID:          row.Int8(iId),
			Name:        row.Str(iName),
			Type:        row.Int8(iType),
			IsSpecial:   row.Int8(iIsSpecial),
			SkillID1:    row.Int16(iIskillId1),
			SkillID2:    row.Int16(iIskillId2),
			Level:       row.Int16(iLevel),
			MissionLock: row.Int32(iMissionLock),
		}
	}
}
