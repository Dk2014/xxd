package town_dat

import (
	"core/fail"
	"core/mysql"
)

var (
	mapTownNpc      map[int16][]int32       // map[town_id][]npc_id
	mapTown         map[int16]*Town         // map[town_id]lock
	mapTownNextTown map[int16]int16         // map[town_id]即将开启的城镇ID
	mapNpcTalk      map[int32]*NpcTalk      //NPC ID -> NPC Talk
	mapTownNpcItem  map[int32][]int16       //map[商人NPC_ID] []可购买物品ID
	mapTownTreasure map[int16]*TownTreasure //城镇宝箱数据
	clubHouseDat    *Town                   //帮派集会所

	firstTownId int16
)

type Award struct {
	ItemId  int16 // 奖励物品ID
	ItemNum int16 // 奖励物品数量
}

type NpcTalk struct {
	TownId int16            // 相关城镇
	Awards map[int16]*Award //对话奖励 任务ID -> 奖励内容
}

func loadNpcTalk(db *mysql.Connection) {
	sql := "SELECT * FROM npc_talk "
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}

	iNpcId := res.Map("npc_id")
	iTownId := res.Map("town_id")
	iType := res.Map("type")
	iQuestId := res.Map("quest_id")
	iAwardItemId := res.Map("award_item_id")
	iAwardItemNum := res.Map("award_item_num")

	var pri_id int32

	mapNpcTalk = map[int32]*NpcTalk{}
	for _, row := range res.Rows {
		pri_id = row.Int32(iNpcId)
		var npcTalk *NpcTalk
		npcTalk, _ = mapNpcTalk[pri_id]
		if npcTalk == nil {
			npcTalk = &NpcTalk{
				TownId: row.Int16(iTownId),
			}
			mapNpcTalk[pri_id] = npcTalk
		}

		ItemId := row.Int16(iAwardItemId)
		ItemNum := row.Int16(iAwardItemNum)
		questId := row.Int16(iQuestId)
		if ItemId <= 0 || ItemNum <= 0 {
			continue
		}

		if npcTalk.Awards == nil {
			npcTalk.Awards = map[int16]*Award{}
		}

		if row.Int8(iType) == NPC_TALK_TYPE_FIRST_TALK {
			npcTalk.Awards[FIRST_TALK_QUEST_ID] = &Award{
				ItemId:  ItemId,
				ItemNum: ItemNum,
			}
		} else if questId > 0 {
			npcTalk.Awards[questId] = &Award{
				ItemId:  ItemId,
				ItemNum: ItemNum,
			}
		}
	}
}

func loadTownNpcItem(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM town_npc_item ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iTownNpcId := res.Map("town_npc_id")
	iItemId := res.Map("item_id")

	var pri_id int32
	mapTownNpcItem = map[int32][]int16{}
	for _, row := range res.Rows {
		pri_id = row.Int32(iTownNpcId)
		mapTownNpcItem[pri_id] = append(mapTownNpcItem[pri_id], row.Int16(iItemId))
	}
}

func Load(db *mysql.Connection) {
	loadTownNpc(db)
	loadTown(db)
	loadNpcTalk(db)
	loadTownNpcItem(db)
	loadTownTreasures(db)
}

type Town struct {
	Lock   int32 // 解锁权值
	StartX int16 // 出生点x轴坐标
	StartY int16 // 出生点y轴坐标
}

func loadTown(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM town ORDER BY `lock` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iId := res.Map("id")
	iLock := res.Map("lock")
	iStartX := res.Map("start_x")
	iStartY := res.Map("start_y")

	var pri_id, lastTownId int16
	mapTown = map[int16]*Town{}
	mapTownNextTown = map[int16]int16{}

	for _, row := range res.Rows {
		if row.Int16(iId) == -1 {
			clubHouseDat = &Town{
				StartX: int16(row.Int32(iStartX)),
				StartY: int16(row.Int32(iStartY)),
			}
			continue
		}
		pri_id = row.Int16(iId)
		mapTown[pri_id] = &Town{
			Lock:   row.Int32(iLock),
			StartX: int16(row.Int32(iStartX)),
			StartY: int16(row.Int32(iStartY)),
		}

		if lastTownId > 0 {
			mapTownNextTown[lastTownId] = pri_id
		}

		lastTownId = pri_id

		if firstTownId == 0 && pri_id > 0 {
			firstTownId = pri_id
		}
	}
	// 下一个城镇为0表示最后一个城镇
	mapTownNextTown[lastTownId] = 0
}

func loadTownNpc(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM town_npc ORDER BY `town_id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iId := res.Map("id")
	iTownId := res.Map("town_id")

	var id int32
	var townId int16

	mapTownNpc = map[int16][]int32{}

	for _, row := range res.Rows {
		id = row.Int32(iId)
		townId = row.Int16(iTownId)

		if _, ok := mapTownNpc[townId]; !ok {
			mapTownNpc[townId] = []int32{}
		}

		mapTownNpc[townId] = append(mapTownNpc[townId], id)
	}
}

type TownTreasure struct {
	TownId        int16 // 城镇ID
	AwardExp      int32 // 奖励经验
	AwardIngot    int32 // 奖励元宝
	AwardCoins    int64 // 奖励铜钱
	AwardPhysical int8  // 奖励体力
	AwardItem1Id  int16 // 奖励物品1
	AwardItem1Num int16 // 奖励物品1数量
	AwardItem2Id  int16 // 奖励物品2
	AwardItem2Num int16 // 奖励物品2数量
	AwardItem3Id  int16 // 奖励物品3
	AwardItem3Num int16 // 奖励物品3数量
	AwardItem4Id  int16 // 奖励物品4
	AwardItem4Num int16 // 奖励物品4数量
}

func loadTownTreasures(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM `town_treasure` ORDER BY `town_id`, `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iTownId := res.Map("town_id")
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

	mapTownTreasure = map[int16]*TownTreasure{}

	for _, row := range res.Rows {
		town_id := row.Int16(iTownId)

		mapTownTreasure[town_id] = &TownTreasure{
			TownId:        town_id,
			AwardExp:      row.Int32(iAwardExp),
			AwardIngot:    row.Int32(iAwardIngot),
			AwardCoins:    row.Int64(iAwardCoins),
			AwardPhysical: row.Int8(iAwardPhysical),
			AwardItem1Id:  row.Int16(iAwardItem1Id),
			AwardItem1Num: row.Int16(iAwardItem1Num),
			AwardItem2Id:  row.Int16(iAwardItem2Id),
			AwardItem2Num: row.Int16(iAwardItem2Num),
			AwardItem3Id:  row.Int16(iAwardItem3Id),
			AwardItem3Num: row.Int16(iAwardItem3Num),
			AwardItem4Id:  row.Int16(iAwardItem4Id),
			AwardItem4Num: row.Int16(iAwardItem4Num),
		}
	}
}

// ############### 对外接口实现 coding here ###############

func ExistNPCIdWithTownId(townId int16, npcId int32) bool {
	fail.When(townId < 0, "townid must greate than 0")
	for _, id := range mapTownNpc[townId] {
		if id == npcId {
			return true
		}
	}

	return false
}

func GetInitTownIdAndLock() (int16, int32) {
	town, ok := mapTown[firstTownId]
	fail.When(!ok, "not found first town id")
	return firstTownId, town.Lock
}

func GetTownWithTownId(townId int16) *Town {
	fail.When(townId < 0, "townid must greate than 0")
	town, ok := mapTown[townId]
	fail.When(!ok, "not found town dat")
	return town
}

func GetJiHuiSuo() *Town {
	return clubHouseDat
}

func GetTownIdWithLock(foundlock int32) int16 {
	var townId int16

	for id, town := range mapTown {
		if foundlock == town.Lock {
			townId = id
			break
		}
	}

	fail.When(townId == 0, "can't found town id with foundlock")
	return townId
}

func GetTownIdListWithLock(maxLock int32) []int16 {
	list := []int16{}
	for id, town := range mapTown {
		if maxLock >= town.Lock {
			list = append(list, id)
		}
	}

	fail.When(len(list) == 0, "can't get town id list")
	return list
}

func GetNextTownIdWithCurrentTownId(townId int16) int16 {
	fail.When(townId < 0, "townid must greate than 0")
	id, ok := mapTownNextTown[townId]
	fail.When(!ok, "not found next town id")
	return id
}

func GetNPCTalkAward(npcId int32) *NpcTalk {
	npcTalk, ok := mapNpcTalk[npcId]
	fail.When(!ok, "NPC对话奖励：次NPC未配置")
	return npcTalk
}

func GetTownNpcItem(npcId int32) []int16 {
	return mapTownNpcItem[npcId]
}

func GetTownTreasure(townId int16) (*TownTreasure, bool) {
	fail.When(townId < 0, "townid must greate than 0")
	treasure, ok := mapTownTreasure[townId]
	return treasure, ok
}
