package event_dat

import (
	"core/fail"
	"core/mysql"
)

var (
	mapLoginAward map[int32][]*LoginAward
)

type LoginAward struct {
	RequrieActiveDays int32 // 累计活跃天数
	AwardType         int8  // 奖励类型
	AwardId           int32 // 奖励内容外键
	AwardNum          int32 // 奖励数量
}

func loadLoginAward(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM login_award ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iRequrieActiveDays := res.Map("requrie_active_days")
	iAwardType := res.Map("award_type")
	iAwardId := res.Map("award_id")
	iAwardNum := res.Map("award_num")

	var pri_id int32
	mapLoginAward = map[int32][]*LoginAward{}
	for _, row := range res.Rows {
		pri_id = row.Int32(iRequrieActiveDays)
		mapLoginAward[pri_id] = append(mapLoginAward[pri_id], &LoginAward{
			RequrieActiveDays: pri_id,
			AwardType:         row.Int8(iAwardType),
			AwardId:           row.Int32(iAwardId),
			AwardNum:          row.Int32(iAwardNum),
		})
	}
}

func GetLoginAwards(loginDays int32) []*LoginAward {
	awards, ok := mapLoginAward[loginDays]
	fail.When(!ok || len(awards) == 0, "登录奖励未配置")
	return awards
}

//单个累计天数的奖励设置
type EventsLoginPerDays struct {
	RequrieActiveDays int32
	AwardsList        []*LoginAward
}
type EventsLoginExt struct {
	StartUnixTime   int64
	EndUnixTime     int64
	DisposeUnixTime int64
	IsRelative      int8
	LTitle          string
	RTitle          string
	Content         string
	Tag             int8
	Weight          int16
	List            []*EventsLoginPerDays
}

// 升级运营活动数据配置
func LoadEventsLogin(list []*EventsLoginPerDays) {
	for _, item := range list {
		if item.RequrieActiveDays > 0 {
			mapLoginAward[item.RequrieActiveDays] = make([]*LoginAward, 0)
			for _, award := range item.AwardsList {
				pri_id := item.RequrieActiveDays
				mapLoginAward[pri_id] = append(mapLoginAward[pri_id], &LoginAward{
					RequrieActiveDays: pri_id,
					AwardType:         award.AwardType,
					AwardId:           award.AwardId,
					AwardNum:          award.AwardNum,
				})
			}
		}
	}
}

func GetMaxWeightInLogin() (max int32) {
	max = 0
	for days, _ := range mapLoginAward {
		if days > max {
			max = days
		}
	}
	return
}
