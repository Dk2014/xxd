package rainbow_dat

import (
	"core/fail"
	"core/i18l"
	"core/mysql"
)

var (
	mapRainbowLevelIds          map[int16][]int32 //map[segment] => []int32{/*mission level id/*}
	MaxRainbowLevelSegment      int16
	mapRainbowLevelAward        map[int32][]*RainbowLevelAward //mission_level_id -> []award
	mapRainbowSegmentAward      map[int16]int16                //map[segment_order] -> award_ghost_id
	mapRainbowAutoFightBoxAward map[int32][]*RainbowLevelAward // 扫荡宝箱
)

type RainbowLevelAward struct {
	AwardType   int8  // 奖励类型(0--铜钱;1--道具;2--装备 3--经验 4--经验倍数 5--铜钱倍数 6--恢复伙伴技能 7--恢复魂侍技能 8--恢复灵宠状态 9--主角精气 10--百分比血量 11-增加魂力)
	AwardChance int8  // 奖励概率
	AwardNum    int32 // 奖励数量
	ItemId      int32 // 物品ID(物品奖励填写)
	Order       int8  // 品质顺序
}

func Load(db *mysql.Connection) {
	loadRainbowLevelId(db)
	loadRainbowLevelAward(db)
	loadRainbowLevel(db)
}

func loadRainbowLevelAward(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM `rainbow_level_award` order by `order` ASC"), -1)
	if err != nil {
		panic(err)
	}
	iMissionLevelId := res.Map("mission_level_id")
	iAwardType := res.Map("award_type")
	iAwardChance := res.Map("award_chance")
	iAwardNum := res.Map("award_num")
	iItemId := res.Map("item_id")
	iOrder := res.Map("order")
	iAutoFightBox := res.Map("autofight_box")

	var pri_id int32
	var autoFightBox int8

	mapRainbowLevelAward = map[int32][]*RainbowLevelAward{}
	// 扫荡宝箱的唯一标识是存储在mission_level_id字段中。含义是彩虹桥段数
	mapRainbowAutoFightBoxAward = map[int32][]*RainbowLevelAward{}

	for _, row := range res.Rows {
		pri_id = row.Int32(iMissionLevelId)
		autoFightBox = row.Int8(iAutoFightBox)
		awardRs := &RainbowLevelAward{
			AwardType:   row.Int8(iAwardType),
			AwardChance: row.Int8(iAwardChance),
			AwardNum:    row.Int32(iAwardNum),
			ItemId:      row.Int32(iItemId),
			Order:       row.Int8(iOrder),
		}

		if autoFightBox == 1 {
			mapRainbowAutoFightBoxAward[pri_id] = append(mapRainbowAutoFightBoxAward[pri_id], awardRs)
		} else {
			mapRainbowLevelAward[pri_id] = append(mapRainbowLevelAward[pri_id], awardRs)
		}
	}
}

func loadRainbowLevelId(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT `id`, `parent_id` FROM mission_level WHERE `parent_type`=12 ORDER BY `parent_id`,`order` ASC"), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iParentId := res.Map("parent_id")

	var segment int16
	mapRainbowLevelIds = make(map[int16][]int32)
	for _, row := range res.Rows {
		segment = row.Int16(iParentId)
		if segment > MaxRainbowLevelSegment {
			MaxRainbowLevelSegment = segment
		}
		mapRainbowLevelIds[segment] = append(mapRainbowLevelIds[segment], row.Int32(iId))
	}
}

func loadRainbowLevel(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM rainbow_level ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iSegment := res.Map("segment")
	iAwardId := res.Map("award_id")

	mapRainbowSegmentAward = map[int16]int16{}
	for _, row := range res.Rows {
		mapRainbowSegmentAward[row.Int16(iSegment)] = row.Int16(iAwardId)
	}
}

//根据彩虹桥段数返回段内所有关卡ID
func GetRainbowLevelId(segment int16, order int8) int32 {
	ids, ok := mapRainbowLevelIds[segment]
	fail.When(!ok, "找不到彩虹关卡数据")
	return ids[order-1]
}

func GetRainbowLevelAward(levelId int32) (awards []*RainbowLevelAward) {
	return mapRainbowLevelAward[levelId]
}

func GetRainbowAutoFightBox(segment int32) []*RainbowLevelAward {
	award, ok := mapRainbowAutoFightBoxAward[segment]
	fail.When(!ok, "扫荡宝箱不存在")
	return award
}

func GetRainbowSegmentAward(segment int16) (awardId int16) {
	awardId, configured := mapRainbowSegmentAward[segment]
	fail.When(!configured, "彩虹关卡段通关奖励未配置")
	return awardId
}

func SegmentNum2SegmentName(segment int16) string {
	switch segment {
	case 1:
		return i18l.T.Tran("彩虹桥一")
	case 2:
		return i18l.T.Tran("彩虹桥二")
	case 3:
		return i18l.T.Tran("彩虹桥三")
	case 4:
		return i18l.T.Tran("彩虹桥四")
	case 5:
		return i18l.T.Tran("彩虹桥五")
	case 6:
		return i18l.T.Tran("彩虹桥六")
	case 7:
		return i18l.T.Tran("彩虹桥七")
	case 8:
		return i18l.T.Tran("彩虹桥八")
	case 9:
		return i18l.T.Tran("彩虹桥九")
	case 10:
		return i18l.T.Tran("彩虹桥十")
	}
	return i18l.T.Tran("彩虹桥")
}
