package role

import (
	"core/fail"
	"game_server/dat/player_dat"
	"game_server/dat/role_dat"
	"game_server/mdb"
	"game_server/module"
)

// 添加新的伙伴角色
func addNewBuddyRole(state *module.SessionState, roleId int8, level int16) (status int8) {
	fail.When(role_dat.IsMainRole(roleId), "main role only have one")
	if findRoleByRoleId(state.Database, roleId, false) != nil {
		return
	}

	total_role_using := 0
	state.Database.Select.PlayerRole(func(row *mdb.PlayerRoleRow) {
		if row.Status() == role_dat.ROLE_STATUS_NOMAL {
			total_role_using++
		}
	})

	if total_role_using >= role_dat.MAX_ROLE_NUM_USING {
		status = role_dat.ROLE_STATUS_ININN
	}

	//Role增加新字段务必在 module/player/player_init 和 和此处给定初始值
	state.Database.Insert.PlayerRole(&mdb.PlayerRole{
		Pid:             state.PlayerId,
		RoleId:          roleId,
		Level:           level,
		FriendshipLevel: 1,
		Exp:             0,
		CoopId:          0,
		Status:          int16(status),
	})

	module.Skill.InitRoleSkill(state, roleId)    // 添加初始技能
	module.Item.InitRoleEquipment(state, roleId) // 初始化角色装备

	// 新添加的伙伴等级高于1级，可能会有新的技能添加
	if level > 1 {
		playerFame := state.Database.Lookup.PlayerFame(state.PlayerId)
		module.Skill.UpdateSkill(state.Database, roleId, 1 /*羁绊最小等级*/, playerFame.Level, int16(level))
	}

	// 初始化角色魂侍装备
	if module.Player.IsFuncInited(state.Database, player_dat.FUNC_GHOST) {
		module.Ghost.InitRoleGhostEquipment(state.Database, roleId)
	}

	// 初始化角色剑心装备
	if module.Player.IsFuncInited(state.Database, player_dat.FUNC_SWORD_SOUL) {
		module.SwordSoul.InitRoleSwordSoulEquipment(state, roleId)
	}
	return
}
