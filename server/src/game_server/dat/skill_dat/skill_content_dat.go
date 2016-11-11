package skill_dat

import (
	"core/fail"
	"core/mysql"
	"fmt"
)

var (
	mapSkillContent map[int16]*SkillContent
)

type SkillContent struct {
	Id         int64 // 主键ID
	SkillId    int16 // 绝招ID
	ReleaseNum int32 // 释放次数
}

func loadSkillContent(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM skill_content ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iId := res.Map("id")
	iSkillId := res.Map("skill_id")
	iReleaseNum := res.Map("release_num")

	mapSkillContent = map[int16]*SkillContent{}
	var pri_id int64
	var skillId int16
	for _, row := range res.Rows {
		pri_id = row.Int64(iId)
		skillId = row.Int16(iSkillId)
		mapSkillContent[skillId] = &SkillContent{
			Id:         pri_id,
			SkillId:    skillId,
			ReleaseNum: row.Int32(iReleaseNum),
		}
	}
}

func GetSkillContent(skillId int16) *SkillContent {
	v, ok := mapSkillContent[skillId]
	fail.When(!ok, fmt.Sprintf("skill content wrong id %d", skillId))
	return v
}
