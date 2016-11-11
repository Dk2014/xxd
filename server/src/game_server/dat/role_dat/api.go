package role_dat

import (
	"core/fail"
	"core/mysql"
)

func Load(db *mysql.Connection) {
	loadRecruitBuddy(db)
	loadPlayerRoles(db)
	loadRoleInfos(db)
	loadRoleLevelExp(db)
	loadRoleFriendship(db)
	loadBuddyCooperation(db)
	loadMainRoleCooperation(db)
}

// 是否主角
func IsMainRole(roleId int8) bool {
	return roleId == PLAYER_BOY_ROLE_ID || roleId == PLAYER_GIRL_ROLE_ID
}

// 获取角色信息
func GetRoleInfo(roleId int8) *RoleInfo {
	info, ok := g_RoleInfos[roleId]
	fail.When(!ok, "获取角色信息错误，模板数据不存在")
	return info
}

// 获取角色等级信息
func GetRoleLevelInfo(roleID int8, level int16) *RoleLevelInfo {
	// 主角2(girl)的数据和主角1(boy)使用同一份
	if roleID == PLAYER_GIRL_ROLE_ID {
		return g_RoleLevelInfos[PLAYER_BOY_ROLE_ID][level-1]
	}
	return g_RoleLevelInfos[roleID][level-1]
}

// 根据等级获取经验
func GetRoleLevelExp(level int16) int64 {
	return mapRoleLevelExp[level-1]
}

// 根据角色及羁绊等级获得羁绊信息
func GetRoleFriendship(roleID int8, friendshipLevel int16) *RoleFriendshipStuff {
	return g_RoleFriendshipStuffes[roleID][friendshipLevel-1 /*参考该数组的生成注释*/]
}

func GetRecruitBuddyDat(roleId int8) *RecruitBuddy {
	return mapRecruitBuddy[roleId]
}
