package totem_dat

import (
	"core/fail"
	"core/mysql"
	"fmt"
)

var (
	mapTotemSkill map[int16][]int16 //totem_id -> []skill_id
)

func loadTotemSkill(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT id,role_id FROM skill WHERE `type`=9 ORDER BY `role_id` ASC"), -1)
	fail.When(err != nil, err)
	iId := res.Map("id")
	iRoleId := res.Map("role_id")
	mapTotemSkill = map[int16][]int16{}
	var role_id int16
	for _, row := range res.Rows {
		role_id = int16(row.Int8(iRoleId))
		mapTotemSkill[role_id] = append(mapTotemSkill[role_id], row.Int16(iId))
	}
	for totemId, skills := range mapTotemSkill {
		fail.When(len(skills) < 3, fmt.Sprintf("阵印技能配置错误 %d ", totemId))
	}
}

func GetSkillByTotemId(totemId int16) (skills []int16) {
	skills = mapTotemSkill[totemId]
	return skills
}
