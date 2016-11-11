package quest_dat

import (
	"core/fail"
	"core/mysql"
	util "core/time"
	"game_server/dat/player_dat"
	"strconv"
	"strings"
	"time"
)

var (
	mapQuest                      map[int16]*Quest
	mapDailyQuest                 map[int16]*DailyQuest
	firstQuestId                  int16
	mapQuestAwardMissionLevelLock map[int16]int32 //quest id -> lock 领取到该任务时应该获得的关卡权值
)

func Load(db *mysql.Connection) {
	loadQuest(db)
	loadDailyQuest(db)
	loadExtendQuest(db)
	loadAdditionQuest(db)
	loadQuestStartAward(db)
}

type Quest struct {
	Type                  int8  // 任务类型
	RequireLevel          int32 // 要求玩家等级
	TownId                int16 // 城镇ID
	TownNpcId             int32 // 完成任务所需的城镇NPC ID
	MissionLevelId        int32 // 完成任务所需的关卡ID
	EnemyNum              int32 // 敌人组数
	EnemyId               int16 // 敌人ID
	AwardExp              int32 // 奖励经验
	AwardCoins            int64 // 奖励铜钱
	AwardItem1Id          int16 // 奖励物品1
	AwardItem1Num         int16 // 奖励物品1数量
	AwardItem2Id          int16 // 奖励物品2
	AwardItem2Num         int16 // 奖励物品2数量
	AwardItem3Id          int16 // 奖励物品3
	AwardItem3Num         int16 // 奖励物品3数量
	AwardItem4Id          int16 // 奖励物品4
	AwardItem4Num         int16 // 奖励物品4数量
	AwardFuncKey          int16 // 奖励功能权值
	AwardRoleId           int8  // 奖励角色ID
	AwardRoleLevel        int8  // 奖励角色等级
	AwardMissionLevelLock int32 // 奖励关卡权值
	AwardTownKey          int32 // 奖励城镇权值
	AwardPhysical         int8  // 奖励体力

	Order       int32 //任务顺序
	NextQuestId int16 // 下一个任务
}

func loadQuest(db *mysql.Connection) {
	//order by `order` 务必不能去掉
	res, err := db.ExecuteFetch([]byte("SELECT * FROM quest ORDER BY `order` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iId := res.Map("id")
	iType := res.Map("type")
	iRequireLevel := res.Map("require_level")
	iTownId := res.Map("town_id")
	iTownNpcId := res.Map("town_npc_id")
	iMissionLevelId := res.Map("mission_level_id")
	iEnemyNum := res.Map("enemy_num")
	iEnemyId := res.Map("enemy_id")
	iAwardExp := res.Map("award_exp")
	iAwardCoins := res.Map("award_coins")
	iAwardItem1Id := res.Map("award_item1_id")
	iAwardItem1Num := res.Map("award_item1_num")
	iAwardItem2Id := res.Map("award_item2_id")
	iAwardItem2Num := res.Map("award_item2_num")
	iAwardItem3Id := res.Map("award_item3_id")
	iAwardItem3Num := res.Map("award_item3_num")
	iAwardItem4Id := res.Map("award_item4_id")
	iAwardItem4Num := res.Map("award_item4_num")
	iAwardFuncKey := res.Map("award_func_key")
	iAwardRoleId := res.Map("award_role_id")
	iAwardRoleLevel := res.Map("award_role_level")
	iAwardMissionLevelLock := res.Map("award_mission_level_lock")
	iAwardTownKey := res.Map("award_town_key")
	iAwardPhysical := res.Map("award_physical")
	iOrder := res.Map("order")

	var awardMissionLevelLock int32
	var pri_id, lastQuestId int16
	mapQuest = map[int16]*Quest{}
	mapQuestAwardMissionLevelLock = map[int16]int32{}
	for _, row := range res.Rows {
		pri_id = row.Int16(iId)
		mapQuest[pri_id] = &Quest{
			Type:                  row.Int8(iType),
			RequireLevel:          row.Int32(iRequireLevel),
			TownId:                row.Int16(iTownId),
			TownNpcId:             row.Int32(iTownNpcId),
			MissionLevelId:        row.Int32(iMissionLevelId),
			EnemyNum:              row.Int32(iEnemyNum),
			EnemyId:               row.Int16(iEnemyId),
			AwardExp:              row.Int32(iAwardExp),
			AwardCoins:            row.Int64(iAwardCoins),
			AwardItem1Id:          row.Int16(iAwardItem1Id),
			AwardItem1Num:         row.Int16(iAwardItem1Num),
			AwardItem2Id:          row.Int16(iAwardItem2Id),
			AwardItem2Num:         row.Int16(iAwardItem2Num),
			AwardItem3Id:          row.Int16(iAwardItem3Id),
			AwardItem3Num:         row.Int16(iAwardItem3Num),
			AwardItem4Id:          row.Int16(iAwardItem4Id),
			AwardItem4Num:         row.Int16(iAwardItem4Num),
			AwardFuncKey:          row.Int16(iAwardFuncKey),
			AwardRoleId:           row.Int8(iAwardRoleId),
			AwardRoleLevel:        row.Int8(iAwardRoleLevel),
			AwardMissionLevelLock: row.Int32(iAwardMissionLevelLock),
			AwardTownKey:          row.Int32(iAwardTownKey),
			AwardPhysical:         row.Int8(iAwardPhysical),
			Order:                 row.Int32(iOrder),
		}

		mapQuestAwardMissionLevelLock[pri_id] = awardMissionLevelLock
		if mapQuest[pri_id].AwardMissionLevelLock > awardMissionLevelLock {
			awardMissionLevelLock = mapQuest[pri_id].AwardMissionLevelLock
		}
		//最后一个主线任务的 NextQuestId是 0，当玩家完成了最后一个主线任务时，需要查询领取下一个主线任务的时候的关卡权值
		mapQuestAwardMissionLevelLock[0] = awardMissionLevelLock

		// 设置任务关系链，最后一条任务NextQuestId为0代码主线任务结束
		if lastQuestId > 0 {
			mapQuest[lastQuestId].NextQuestId = pri_id
		}
		lastQuestId = pri_id

		if firstQuestId == 0 {
			firstQuestId = pri_id
		}
	}
}

type DailyQuest struct {
	Class           int16          // 类别
	RequireMinLevel int32          // 要求玩家最低等级
	RequireMaxLevel int32          // 要求玩家最高等级
	RequireOpenDay  []time.Weekday // 开放日
	RequireCount    int16          // 需要数量
	AwardExp        int32          // 奖励经验
	AwardIngot      int32          // 奖励元宝
	AwardCoins      int64          // 奖励铜钱
	AwardPhysical   int8           // 奖励体力
	AwardItem1Id    int16          // 奖励物品1
	AwardItem1Num   int16          // 奖励物品1数量
	AwardItem2Id    int16          // 奖励物品2
	AwardItem2Num   int16          // 奖励物品2数量
	AwardItem3Id    int16          // 奖励物品3
	AwardItem3Num   int16          // 奖励物品3数量
	AwardItem4Id    int16          // 奖励物品4
	AwardItem4Num   int16          // 奖励物品4数量
	LevelType       int8           // 关卡类型; -1 无; 0-区域关卡;1-资源关卡;2-通天塔;8-难度关卡;9-伙伴关卡;10-灵宠关卡;11-魂侍关卡
	LevelSubType    int8           // 关卡子类型(-1--无;1--铜钱关卡;2--经验关卡)
	AwardStars      int8           // 奖励星数
}

func loadDailyQuest(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM daily_quest ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iId := res.Map("id")
	iRequireMinLevel := res.Map("require_min_level")
	iRequireMaxLevel := res.Map("require_max_level")
	iRequireOpenDay := res.Map("require_open_day")
	iRequireCount := res.Map("require_count")
	iAwardExp := res.Map("award_exp")
	iAwardIngot := res.Map("award_ingot")
	iAwardCoins := res.Map("award_coins")
	iAwardPhysical := res.Map("award_physical")
	iAwardItem1Id := res.Map("award_item1_id")
	iAwardItem1Num := res.Map("award_item1_num")
	iAwardItem2Id := res.Map("award_item2_id")
	iAwardItem2Num := res.Map("award_item2_num")
	iAwardItem3Id := res.Map("award_item3_id")
	iAwardItem3Num := res.Map("award_item3_num")
	iAwardItem4Id := res.Map("award_item4_id")
	iAwardItem4Num := res.Map("award_item4_num")
	iLevelType := res.Map("level_type")
	iLevelSubType := res.Map("level_sub_type")
	iClass := res.Map("class")
	iAwardStars := res.Map("award_stars")

	var pri_id, class int16

	mapDailyQuest = map[int16]*DailyQuest{}

	for _, row := range res.Rows {
		pri_id = row.Int16(iId)
		class = row.Int16(iClass)
		openDay := []time.Weekday{}

		openDayStr := strings.Trim(row.Str(iRequireOpenDay), " ")
		if openDayStr != "" {
			openDayList := strings.Split(openDayStr, " ")
			for _, val := range openDayList {
				day, err := strconv.Atoi(val)
				if err != nil {
					panic(err)
				}
				openDay = append(openDay, time.Weekday(day))
			}
		}

		quest := &DailyQuest{
			Class:           class,
			RequireMinLevel: row.Int32(iRequireMinLevel),
			RequireMaxLevel: row.Int32(iRequireMaxLevel),
			RequireOpenDay:  openDay,
			RequireCount:    row.Int16(iRequireCount),
			AwardExp:        row.Int32(iAwardExp),
			AwardIngot:      row.Int32(iAwardIngot),
			AwardCoins:      row.Int64(iAwardCoins),
			AwardPhysical:   row.Int8(iAwardPhysical),
			AwardItem1Id:    row.Int16(iAwardItem1Id),
			AwardItem1Num:   row.Int16(iAwardItem1Num),
			AwardItem2Id:    row.Int16(iAwardItem2Id),
			AwardItem2Num:   row.Int16(iAwardItem2Num),
			AwardItem3Id:    row.Int16(iAwardItem3Id),
			AwardItem3Num:   row.Int16(iAwardItem3Num),
			AwardItem4Id:    row.Int16(iAwardItem4Id),
			AwardItem4Num:   row.Int16(iAwardItem4Num),
			LevelType:       row.Int8(iLevelType),
			LevelSubType:    row.Int8(iLevelSubType),
			AwardStars:      row.Int8(iAwardStars),
		}

		mapDailyQuest[pri_id] = quest
	}
}

// ############### 对外接口实现 coding here ###############

func GetQuestById(id int16) *Quest {
	quest, ok := mapQuest[id]
	fail.When(!ok, "can't found quest")
	return quest
}

func GetQuestsByOrder(order int32) (quests []*Quest) {
	for _, quest := range mapQuest {
		if quest.Order <= order {
			quests = append(quests, quest)
		}
	}
	return quests
}

func GetInitQuest() int16 {
	return firstQuestId
}

func GetDailyQuestWithQuestId(id int16) *DailyQuest {
	quest, ok := mapDailyQuest[id]
	fail.When(!ok, "can't found DailyQuest")
	return quest
}

func GetDailyQuestWithLevel(roleLevel int32) map[int16]*DailyQuest {
	q := map[int16]*DailyQuest{}
	today := util.GetNowWeekByExpectHour(player_dat.RESET_DAILY_QUEST_IN_HOUR)

	for id, quest := range mapDailyQuest {
		if roleLevel >= quest.RequireMinLevel && roleLevel <= quest.RequireMaxLevel {
			if len(quest.RequireOpenDay) > 0 {
				for _, openDay := range quest.RequireOpenDay {
					if today == openDay {
						q[id] = quest
						break
					}
				}
			} else {
				q[id] = quest
			}
		}
	}

	return q
}
