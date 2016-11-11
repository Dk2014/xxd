package team

import (
	"core/fail"
	"game_server/dat/quest_dat"
	"game_server/dat/role_dat"
	"game_server/dat/team_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/tlog"
)

func getFormPosArray(formationInfo *mdb.PlayerFormation) (posArray [9]*int8) {
	posArray[0] = &formationInfo.Pos0
	posArray[1] = &formationInfo.Pos1
	posArray[2] = &formationInfo.Pos2
	posArray[3] = &formationInfo.Pos3
	posArray[4] = &formationInfo.Pos4
	posArray[5] = &formationInfo.Pos5
	posArray[6] = &formationInfo.Pos6
	posArray[7] = &formationInfo.Pos7
	posArray[8] = &formationInfo.Pos8
	return posArray
}

// 角色上阵
func upFormation(state *module.SessionState, roleId, pos int8) {
	fail.When(!checkPos(pos), "error: position not in range ")

	buddyRole := module.Role.GetBuddyRoleInTeam(state.Database, roleId)
	fail.When(buddyRole == nil, "未获得角色")

	formationInfo := state.Database.Lookup.PlayerFormation(state.PlayerId)
	posArray := getFormPosArray(formationInfo)
	var inFormNum int8
	fail.When(*posArray[pos] != team_dat.POS_NO_ROLE, "目标位置已有角色")
	mainRoleLv := module.Role.GetMainRole(state.Database).Level
	maxInFormNum := team_dat.GetMaxInFormRoleNum(mainRoleLv)
	for _, posPtr := range posArray {
		fail.When(*posPtr == roleId, "角色已上阵")
		if posPtr != nil && *posPtr != team_dat.POS_NO_ROLE {
			inFormNum++
		}
		fail.When(inFormNum >= maxInFormNum, "上阵人数达上限，无法上阵")
	}

	setPosRole(formationInfo, pos, roleId)
	state.Database.Update.PlayerFormation(formationInfo)
	tlog.PlayerSystemModuleFlowLog(state.Database, tlog.SMT_TEAM)
}

// 角色下阵
func downFormation(db *mdb.Database, pos int8) {
	fail.When(!checkPos(pos), "error: position not in range ")
	formationInfo := db.Lookup.PlayerFormation(db.PlayerId())

	fail.When(role_dat.IsMainRole(getRoleIdByPos(formationInfo, pos)),
		"main role can't downFormation")

	posArray := getFormPosArray(formationInfo)
	if *posArray[pos] == team_dat.POS_NO_ROLE {
		return
	}

	module.Role.BreakBuddyCoop(db, *posArray[pos])
	*posArray[pos] = team_dat.POS_NO_ROLE
	db.Update.PlayerFormation(formationInfo)
	tlog.PlayerSystemModuleFlowLog(db, tlog.SMT_TEAM)
}

// 交换两个角色布阵位置, 空格子要交换
func swapFormation(state *module.SessionState, posFrom, posTo int8) {
	fail.When(posFrom == posTo, "error: posFrom equal posTo")
	fail.When(!checkPos(posFrom) || !checkPos(posTo), "error: position not in range")

	formationInfo := state.Database.Lookup.PlayerFormation(state.PlayerId)

	posFromRoleId := getRoleIdByPos(formationInfo, posFrom)
	posToRoleId := getRoleIdByPos(formationInfo, posTo)

	setPosRole(formationInfo, posFrom, posToRoleId)
	setPosRole(formationInfo, posTo, posFromRoleId)

	state.Database.Update.PlayerFormation(formationInfo)
	tlog.PlayerSystemModuleFlowLog(state.Database, tlog.SMT_TEAM)
}

//在阵上和不在阵上的两个角色交换
func replaceFormation(state *module.SessionState, roleId, pos int8) {
	downFormation(state.Database, pos)
	upFormation(state, roleId, pos)
}

// 设置某个角色到某个位置
func setPosRole(formationInfo *mdb.PlayerFormation, pos, roleId int8) {
	switch pos {
	case team_dat.POS0:
		formationInfo.Pos0 = roleId
	case team_dat.POS1:
		formationInfo.Pos1 = roleId
	case team_dat.POS2:
		formationInfo.Pos2 = roleId
	case team_dat.POS3:
		formationInfo.Pos3 = roleId
	case team_dat.POS4:
		formationInfo.Pos4 = roleId
	case team_dat.POS5:
		formationInfo.Pos5 = roleId
	case team_dat.POS6:
		formationInfo.Pos6 = roleId
	case team_dat.POS7:
		formationInfo.Pos7 = roleId
	case team_dat.POS8:
		formationInfo.Pos8 = roleId
	default:
		return
	}
}

// 根据位置获取角色id 返回-1表示该位置没有角色
func getRoleIdByPos(formationInfo *mdb.PlayerFormation, pos int8) int8 {
	switch pos {
	case team_dat.POS0:
		return formationInfo.Pos0
	case team_dat.POS1:
		return formationInfo.Pos1
	case team_dat.POS2:
		return formationInfo.Pos2
	case team_dat.POS3:
		return formationInfo.Pos3
	case team_dat.POS4:
		return formationInfo.Pos4
	case team_dat.POS5:
		return formationInfo.Pos5
	case team_dat.POS6:
		return formationInfo.Pos6
	case team_dat.POS7:
		return formationInfo.Pos7
	case team_dat.POS8:
		return formationInfo.Pos8
	default:
		return team_dat.POS_NO_ROLE
	}
}

// 检测位置是否合法
func checkPos(pos int8) bool {
	if pos < 0 || pos > 8 {
		return false
	}
	return true
}

// 伙伴配合训练
func trainingTeamship(state *module.SessionState, attrInd int8) {
	teamInfo := state.Database.Lookup.PlayerTeamInfo(state.PlayerId)
	switch attrInd {
	case team_dat.TEAMSHIP_HEALTH_IND:
		teamInfo.Relationship = trainSpecTeamship(state, teamInfo.HealthLv, teamInfo.Relationship)
		teamInfo.HealthLv++
	case team_dat.TEAMSHIP_ATTACK_IND:
		teamInfo.Relationship = trainSpecTeamship(state, teamInfo.AttackLv, teamInfo.Relationship)
		teamInfo.AttackLv++
	case team_dat.TEAMSHIP_DEFENCE_IND:
		teamInfo.Relationship = trainSpecTeamship(state, teamInfo.DefenceLv, teamInfo.Relationship)
		teamInfo.DefenceLv++
	default:
		fail.When(true, "the teamship training index is invalid")
	}
	module.Quest.RefreshDailyQuest(state, quest_dat.DAILY_QUEST_CLASS_TEAMSHIP_TRAINING)
	state.Database.Update.PlayerTeamInfo(teamInfo)
}

// 按等级检查操作可行性及扣除消耗
func trainSpecTeamship(state *module.SessionState, level int16, relationship int32) int32 {
	// 检查伙伴配合等级是否达到最大
	fail.When(level >= team_dat.TEAMSHIP_MAX_LEVEL, "the teamship level is already the maximum level")

	// 伙伴配合等级必须不大于主角等级
	fail.When(level >= module.Role.GetMainRole(state.Database).Level, "the teamship level shouldn't be greater than main role level")

	// 检查友情量是否充足
	teamshipStuff := team_dat.GetTeamshipStuff(level)
	fail.When(teamshipStuff.NeedsRelationship > relationship, "the amount of relationship is not enough")

	return relationship - teamshipStuff.NeedsRelationship
}
