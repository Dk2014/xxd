package event_dat

import (
	"core/mysql"
	"game_server/dat/mail_dat"
)

type LevelUpAward struct {
	Require_level int16
	Awards        []*mail_dat.Attachment
}

var (
	listEventsLevelUp []*LevelUpAward
)

func loadEventsLevelUp(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM events_level_up ORDER BY `require_level` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iRequireLevel := res.Map("require_level")
	iIngot := res.Map("ingot")
	iItem1Id := res.Map("item1_id")
	iItem1Num := res.Map("item1_num")
	iItem2Id := res.Map("item2_id")
	iItem2Num := res.Map("item2_num")
	iItem3Id := res.Map("item3_id")
	iItem3Num := res.Map("item3_num")
	iCoin := res.Map("coins")

	var pri_require_level int16
	listEventsLevelUp = make([]*LevelUpAward, 0)

	for _, row := range res.Rows {
		pri_require_level = row.Int16(iRequireLevel)
		attachs := makeAttachment(row.Int64(iCoin), row.Int16(iIngot),
			row.Int16(iItem1Id), row.Int16(iItem1Num),
			row.Int16(iItem2Id), row.Int16(iItem2Num),
			row.Int16(iItem3Id), row.Int16(iItem3Num), 0, 0, 0, 0)

		if len(attachs) > 0 {
			listEventsLevelUp = append(listEventsLevelUp, &LevelUpAward{
				Require_level: pri_require_level,
				Awards:        attachs,
			})
		}

	}
}

// ############### 对外接口实现 coding here ###############
func GetEventLevelUpAward(level int16) (attaches []*mail_dat.Attachment, ok bool) {
	for index := 0; index < len(listEventsLevelUp); index++ {
		level_up_awards := listEventsLevelUp[index]
		if level_up_awards.Require_level == level {
			attaches = level_up_awards.Awards
			ok = true
			break
		}
	}
	return
}

func GetEventLevelUpAwards() []*LevelUpAward {
	return listEventsLevelUp
}

type EventsLevelUp struct {
	RequireLevel int16 // 需要等级
	Ingot        int16 // 奖励元宝
	Coin         int64 // 奖励铜钱
	Item1Id      int16 // 物品1
	Item1Num     int16 // 物品1数量
	Item2Id      int16 // 物品2
	Item2Num     int16 // 物品2数量
	Item3Id      int16 // 物品3
	Item3Num     int16 // 物品3数量
}

type EventsLevelUpExt struct {
	StartUnixTime   int64
	EndUnixTime     int64
	DisposeUnixTime int64
	IsRelative      int8
	LTitle          string
	RTitle          string
	Content         string
	Tag             int8
	Weight          int16
	List            []*EventsLevelUp
}

// 升级运营活动数据配置
func LoadEventsLevelUp(list []*EventsLevelUp) {
	for _, item := range list {
		attachs := makeAttachment(item.Coin, item.Ingot,
			item.Item1Id, item.Item1Num,
			item.Item2Id, item.Item2Num,
			item.Item3Id, item.Item3Num, 0, 0, 0, 0)

		if len(attachs) > 0 {
			for index := 0; index < len(listEventsLevelUp); index++ {
				if listEventsLevelUp[index].Require_level == item.RequireLevel {
					listEventsLevelUp[index].Awards = attachs
				}
			}
		}
	}
}

func GetMaxWeightInLevel() int32 {
	return int32(listEventsLevelUp[len(listEventsLevelUp)-1].Require_level)
}

func GetNextLevel(now int32) (next int32) {
	for index := 0; index < len(listEventsLevelUp); index++ {
		if listEventsLevelUp[index].Require_level > int16(now) {
			next = int32(listEventsLevelUp[index].Require_level)
			break
		}
	}
	return
}

func GetPlayerLevelAward(awarded int32) (newAwarded, nextAward int32, awards []*mail_dat.Attachment) {
	var index int
	for index = 0; index < len(listEventsLevelUp); index++ {
		if listEventsLevelUp[index].Require_level > int16(awarded) {
			newAwarded = int32(listEventsLevelUp[index].Require_level)
			awards = listEventsLevelUp[index].Awards
			break
		}
	}
	if index < len(listEventsLevelUp)-1 {
		nextAward = int32(listEventsLevelUp[index+1].Require_level)
	}
	return
}
