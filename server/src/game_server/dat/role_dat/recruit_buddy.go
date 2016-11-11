package role_dat

import (
	"core/mysql"
)

var (
	mapRecruitBuddy map[int8]*RecruitBuddy
)

type RecruitBuddy struct {
	RoleId         int8   // 角色ID
	InitLevel      int16  // 初始等级
	Description    string // 描述
	FavouriteItem  int16  // 喜好品ID
	FavouriteCount int16  // 喜好品需求量
	QuestId        int16  // 开启任务
}

func loadRecruitBuddy(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM recruit_buddy ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iRoleId := res.Map("role_id")
	iInitLevel := res.Map("init_level")
	iDescription := res.Map("description")
	iFavouriteItem := res.Map("favourite_item")
	iFavouriteCount := res.Map("favourite_count")
	iQuestId := res.Map("quest_id")

	var pri_id int8
	mapRecruitBuddy = map[int8]*RecruitBuddy{}
	for _, row := range res.Rows {
		pri_id = row.Int8(iRoleId)
		mapRecruitBuddy[pri_id] = &RecruitBuddy{
			RoleId:         pri_id,
			InitLevel:      row.Int16(iInitLevel),
			Description:    row.Str(iDescription),
			FavouriteItem:  row.Int16(iFavouriteItem),
			FavouriteCount: row.Int16(iFavouriteCount),
			QuestId:        row.Int16(iQuestId),
		}
	}
}

func GetRecruitBuddys() map[int8]*RecruitBuddy {
	return mapRecruitBuddy
}
