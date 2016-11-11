package friend_dat

import (
	"core/mysql"
)

var (
	platformFriendAwardConfigs []*PlatformFriendAward
)

func Load(db *mysql.Connection) {
	loadPlatformFriendAward(db)
}

type PlatformFriendAward struct {
	RequireFriendNum int32  // 平台好友数
	AwardType        int8   // 奖励类型 0-物品 1-剑心 2-魂侍 3-灵宠契约球 4-爱心 5-铜币 6-元宝 7-体力
	AwardId          int16  // 物品ID
	Num              int32  // 物品数量
	Name             string //物品名称
}

func loadPlatformFriendAward(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM platform_friend_award ORDER BY `require_friend_num` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iRequireFriendNum := res.Map("require_friend_num")
	iAwardType := res.Map("award_type")
	iAwardId := res.Map("award_id")
	iNum := res.Map("num")
	iName := res.Map("name")

	platformFriendAwardConfigs = []*PlatformFriendAward{}
	for _, row := range res.Rows {
		platformFriendAwardConfigs = append(platformFriendAwardConfigs, &PlatformFriendAward{
			RequireFriendNum: row.Int32(iRequireFriendNum),
			AwardType:        row.Int8(iAwardType),
			AwardId:          row.Int16(iAwardId),
			Num:              row.Int32(iNum),
			Name:             row.Str(iName),
		})
	}
}
