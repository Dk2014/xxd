package ghost

import (
	"game_server/api/protocol/ghost_api"
)

const (
	DAILY_DRAW_NUM = 3

	DRAW_CONSUME = 10000

	BAG_FULL = -1 // 背包已经满

	LEVEL_FOR_INIT_GHOST_POWER_LEVEL1 = 30 // 初始魂力一级
	LEVEL1_ADD_INIT_GHOST_POWER       = 5  // 初始魂力一级增加的魂力

	LEVEL_FOR_INIT_GHOST_POWER_LEVEL2 = 60 // 初始魂力二级
	LEVEL2_ADD_INIT_GHOST_POWER       = 10 // 初始魂力二级增加的魂力

	GHOST_YONGQIZHIWEI = 1 // 勇气之卫

	POS2_NEED_ROLE_LEVEL = 20
	POS3_NEED_ROLE_LEVEL = 40
	POS4_NEED_ROLE_LEVEL = 60
)

func GetGhostGridRequireLevel(pos ghost_api.EquipPos) int16 {
	switch pos {
	case ghost_api.EQUIP_POS_POS2:
		return POS2_NEED_ROLE_LEVEL
	case ghost_api.EQUIP_POS_POS3:
		return POS3_NEED_ROLE_LEVEL
	case ghost_api.EQUIP_POS_POS4:
		return POS4_NEED_ROLE_LEVEL
	}
	return 0
}
