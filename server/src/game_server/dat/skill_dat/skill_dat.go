package skill_dat

import (
	"core/fail"
	"core/mysql"
)

var (
	g_SkillInfoMap           map[int16]*SkillInfo
	g_RoleSkills             map[int8][]*SkillInfo
	g_GhostSkillMap          map[int8]*SkillInfo //TODO role id 字段可能需要改为 int16 避免溢出
	g_CheatSkillMap          map[int16]*SkillInfo
	g_SkillTrainingCost      []int64
	g_SkillTrainingTotalCost []int64
)

type SkillInfo struct {
	Id                      int16
	Type                    int8
	ChildKind               int8
	ChildType               int8
	RoleId                  int8
	RequiredLevel           int16
	RequiredFameLevel       int16
	RequiredFriendshipLevel int16
	CanAddLevel             int8
	ParentSkillId           int16
	CheatId                 int16
	AutoLearnLevel          int8
}

func Load(db *mysql.Connection) {
	loadSkillInfo(db)
	loadSkillContent(db)
	loadSkillTrainingCost(db)
}

func loadSkillInfo(db *mysql.Connection) {
	g_SkillInfoMap = make(map[int16]*SkillInfo)
	g_RoleSkills = make(map[int8][]*SkillInfo)
	g_GhostSkillMap = map[int8]*SkillInfo{}
	g_CheatSkillMap = make(map[int16]*SkillInfo)

	// 不加载type＝8的魂侍第二绝招
	res, err := db.ExecuteFetch([]byte("select * from skill where type not in (8) order by role_id, required_level"), -1)
	if err != nil {
		panic(err)
	}

	iId := res.Map("id")
	iType := res.Map("type")
	iChildKind := res.Map("child_kind")
	iChildType := res.Map("child_type")
	iRoleId := res.Map("role_id")
	iRequiredLevel := res.Map("required_level")
	iRequiredFameLevel := res.Map("required_fame_level")
	iRequiredFriendshipLevel := res.Map("required_friendship_level")
	iCanAddLevel := res.Map("can_add_level")
	iParentSkillId := res.Map("parent_skill_id")
	icheatId := res.Map("cheat_id")
	iautoLearnLevel := res.Map("auto_learn_level")

	for _, row := range res.Rows {
		skill := &SkillInfo{
			Id:                      row.Int16(iId),
			Type:                    row.Int8(iType),
			ChildKind:               row.Int8(iChildKind),
			ChildType:               row.Int8(iChildType),
			RoleId:                  row.Int8(iRoleId),
			RequiredLevel:           row.Int16(iRequiredLevel),
			RequiredFameLevel:       row.Int16(iRequiredFameLevel),
			RequiredFriendshipLevel: row.Int16(iRequiredFriendshipLevel),
			CanAddLevel:             row.Int8(iCanAddLevel),
			ParentSkillId:           row.Int16(iParentSkillId),
			CheatId:                 row.Int16(icheatId),
			AutoLearnLevel:          row.Int8(iautoLearnLevel),
		}

		g_SkillInfoMap[row.Int16(iId)] = skill

		if row.Int8(iType) == OWNER_ROLE {
			g_RoleSkills[row.Int8(iRoleId)] = append(g_RoleSkills[row.Int8(iRoleId)], skill)
		}

		if row.Int16(icheatId) > 0 {
			g_CheatSkillMap[row.Int16(icheatId)] = skill
		}

		if row.Int8(iType) == OWNER_GHOST {
			g_GhostSkillMap[row.Int8(iRoleId)] = skill
		}
	}
}

func loadSkillTrainingCost(db *mysql.Connection) {
	g_SkillTrainingCost = []int64{}
	g_SkillTrainingTotalCost = []int64{}

	res, err := db.ExecuteFetch([]byte("select * from `skill_training_cost` order by `level` asc;"), -1)
	if err != nil {
		panic(err)
	}

	iCost := res.Map("cost")

	var totalCost int64 = 0

	for _, row := range res.Rows {
		g_SkillTrainingCost = append(g_SkillTrainingCost, row.Int64(iCost))
		totalCost += g_SkillTrainingCost[len(g_SkillTrainingCost)-1]
		g_SkillTrainingTotalCost = append(g_SkillTrainingTotalCost, totalCost)
	}
}

func GetSkillInfo(id int16) *SkillInfo {
	v, ok := g_SkillInfoMap[id]
	fail.When(!ok, "GetSkillInfo wrong id")
	return v
}

func GetSkillByRoleIdWithLevel(roleId int8, friendshipLevel, fameLevel, level int16) []*SkillInfo {
	skills, ok := g_RoleSkills[roleId]
	fail.When(!ok, "GetSkillByRoleIdWithLevel wrong role id")

	var canAddSkill []*SkillInfo
	for _, skill := range skills {
		if skill.AutoLearnLevel == SKILL_CAN_AUTO_LEARN && skill.RequiredFriendshipLevel <= friendshipLevel && skill.RequiredFameLevel <= fameLevel && skill.RequiredLevel <= level {
			canAddSkill = append(canAddSkill, skill)
		}
	}

	return canAddSkill
}

func GetGhostSkillByGhostId(ghostId int8) (skillInfo *SkillInfo) {
	skillInfo = g_GhostSkillMap[ghostId]
	fail.When(skillInfo == nil, "wrong RequiredStar")
	return skillInfo
}

func GetSkillTypeByPos(pos int8) int8 {
	switch pos {
	case POS_SKILL_1:
		return SKILL_KIND_FOR_POS_1
	case POS_SKILL_2:
		return SKILL_KIND_FOR_POS_2
	case POS_SKILL_3:
		return SKILL_KIND_FOR_POS_3
	case POS_SKILL_4:
		return SKILL_KIND_FOR_POS_4
	}
	panic("unreachable")
}

func GetSkillByCheatId(cheatId int16) (skillInfo *SkillInfo) {
	skill := g_CheatSkillMap[cheatId]
	return skill
}

func GetSkillTrainingCost(level int16) int64 {
	return g_SkillTrainingCost[level-1]
}

func GetSkillTrainingTotalCost(level int16) int64 {
	return g_SkillTrainingTotalCost[level-1]
}
