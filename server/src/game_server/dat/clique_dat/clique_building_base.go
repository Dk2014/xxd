package clique_dat

import (
	"core/fail"
	"core/mysql"
	"fmt"
)

var (
	arrayCliqueCenterBuildingLevelInfo []*CliqueCenterBuildingLevelInfo
)

type CliqueCenterBuildingLevelInfo struct {
	Id            int32  // ID
	Level         int16  // 等级
	MaxMember     int16  // 成员数量
	LevelupCoins  int64  // 升级消耗金钱
	DailyMaxCoins int64  // 每日贡献上限
	Desc          string // 对应等级描述
	CanUpgrade    bool   //是否可继续升级
}

func loadCliqueCenterBuildingLevelInfo(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM clique_center_building_level_info ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iLevel := res.Map("level")
	iMaxMember := res.Map("max_member")
	iLevelupCoins := res.Map("levelup_coins")
	iDailyMaxCoins := res.Map("daily_max_coins")

	arrayCliqueCenterBuildingLevelInfo = nil
	for _, row := range res.Rows {
		if len(arrayCliqueCenterBuildingLevelInfo) > 0 {
			arrayCliqueCenterBuildingLevelInfo[len(arrayCliqueCenterBuildingLevelInfo)-1].CanUpgrade = true
		}
		arrayCliqueCenterBuildingLevelInfo = append(arrayCliqueCenterBuildingLevelInfo, &CliqueCenterBuildingLevelInfo{
			Level:         row.Int16(iLevel),
			MaxMember:     row.Int16(iMaxMember),
			LevelupCoins:  row.Int64(iLevelupCoins),
			DailyMaxCoins: row.Int64(iDailyMaxCoins),
		})
	}
}

// ############### 对外接口实现 coding here ###############

func GetCenterBuildingLevelInfo(level int16) *CliqueCenterBuildingLevelInfo {
	fail.When(level <= 0 || int(level) > len(arrayCliqueCenterBuildingLevelInfo), fmt.Sprintf("未配置的总舵等级%s", level))
	return arrayCliqueCenterBuildingLevelInfo[level-1]
}
