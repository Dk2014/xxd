package mission_dat

import (
	"core/fail"
	"core/mysql"
)

var (
	mapMission map[int16]*Mission
	//mapTownMission map[int16][]int16 // town_id->list mission_id
	mapMissonLevelTown map[int16]int16

	mapMissionLevel map[int32]*MissionLevel
	//mapMissionLevelID     map[int16][]int32
	mapLeveIDMisson       map[int32]int16
	mapSpecialLevel       map[int8]map[int16]int32          // map[ParentType][ParentId][关卡ID]
	mapMissionLevelByItem map[int16]map[int32]*MissionLevel //map[ItemId]MissionLevel

	mapMissionEnemy    map[int32]*MissionEnemy
	mapEnemyRole       map[int32]*EnemyRole
	mapMissionLevelBox map[int32][]*MissionLevelBox // map[关卡ID][品质顺序]宝箱数据
	mapExtendLevel     map[int8]map[int16][]int32   // map[类型]map[等级上限][]关卡ID

	mapMissionLevelSmallBox      map[int32][]*MissionLevelSmallBox
	mapMissionLevelSmallBoxItems map[int32][]*MissionLevelSmallBoxItems

	mapMissionLevelMengYao map[int32][]*MissionLevelMengYao

	mapShadedMission       map[int32]*ShadedMission
	mapMissionId2ShadedIds map[int32][]int32 //mission_level下含的影之间隙

	mapMissionLevelEnemy map[int32][]int32 //map[关卡ID][]int32{怪物组ID}

	mapHardLevel map[int16]*HardLevel //难度关卡

	mapLevelStar       map[int32]*LevelStar    //关卡评价
	mapPveLevel        = map[int16]*PveLevel{} //灵宠幻境关卡信息 id -> dat
	mapPveLevelByFloor = map[int16]*PveLevel{} //灵宠幻境关卡信息 floor -> dat

	mapRandAwardBox   = map[int32][]*RandAwardBox{}   //随机奖品保箱
	mapTownStarAwards = map[int16][]*TownStarAwards{} //城镇评星奖励
)

func Load(db *mysql.Connection) {
	loadMission(db)
	loadMissionLevel(db)
	loadMissionEnemy(db)
	loadEnemyRole(db)
	loadMissionLevelBox(db)
	loadExtendLevel(db)
	loadMissionLevelSmallBox(db)
	loadMissionLevelSmallBoxItems(db)
	loadMissionLevelMengYao(db)
	loadShadedMission(db)
	loadHardLevel(db)
	loadLevelStar(db)
	loadPveLevel(db)
	loadRandAwardBox(db)
	loadTownStarAwards(db)
}

type Mission struct {
	Id     int16  // 区域ID
	TownId int16  // 城镇ID
	Keys   int32  // 开启钥匙数
	Name   string // 区域名称
	Sign   string // 资源标识
	Order  int8   // 区域开启顺序
}

func loadMission(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM mission ORDER BY `id`,`order` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iId := res.Map("id")
	iTownId := res.Map("town_id")
	iKeys := res.Map("keys")
	iName := res.Map("name")
	iSign := res.Map("sign")
	iOrder := res.Map("order")

	var pri_id int16
	mapMission = map[int16]*Mission{}
	//mapTownMission = make(map[int16][]int16, 0)
	mapMissonLevelTown = make(map[int16]int16, 0)
	for _, row := range res.Rows {
		pri_id = row.Int16(iId)
		mapMission[pri_id] = &Mission{
			Id:     pri_id,
			TownId: row.Int16(iTownId),
			Keys:   row.Int32(iKeys),
			Name:   row.Str(iName),
			Sign:   row.Str(iSign),
			Order:  row.Int8(iOrder),
		}
		// town ->list[mission_id]
		//mapTownMission[row.Int16(iTownId)] = append(mapTownMission[row.Int16(iTownId)], pri_id)
		mapMissonLevelTown[pri_id] = row.Int16(iTownId)
	}
}

type LevelStar struct {
	LevelId        int32 // 关卡ID
	TwoStarScore   int32 // 两星要求分数
	ThreeStarScore int32 // 三星要求分数
	TwoStarRound   int8  // 两星要求回合
	ThreeStarRound int8  // 三星要求回合
}

func loadLevelStar(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM level_star ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iLevelId := res.Map("level_id")
	// iTwoStarScore := res.Map("two_star_score")
	// iThreeStarScore := res.Map("three_star_score")
	iTwoStarRound := res.Map("two_star_round")
	iThreeStarRound := res.Map("three_star_round")

	var pri_id int32
	mapLevelStar = map[int32]*LevelStar{}
	for _, row := range res.Rows {
		pri_id = row.Int32(iLevelId)
		mapLevelStar[pri_id] = &LevelStar{
			LevelId: pri_id,
			// TwoStarScore:   row.Int32(iTwoStarScore),
			// ThreeStarScore: row.Int32(iThreeStarScore),
			TwoStarRound:   row.Int8(iTwoStarRound),
			ThreeStarRound: row.Int8(iThreeStarRound),
		}
	}
}
func CalLevelStar(levelId int32) *LevelStar {
	return mapLevelStar[levelId]
}
func CalLevelStarByRound(levelId int32, round int8) int8 {
	lvStar, ok := mapLevelStar[levelId]
	if !ok {
		return -1
	}
	if round == 0 {
		return 0
	}

	if round < lvStar.ThreeStarRound {
		return THREE_STAR
	} else if round < lvStar.TwoStarRound {
		return TWO_STAR
	}

	return ONE_STAR
}

type MissionLevel struct {
	Id                int32  // 区域关卡ID
	MissionId         int16  // 区域ID
	Lock              int32  // 关卡开启的权值
	Name              string // 关卡名称
	Type              int8   // 关卡类型(0--普通;1--精英;2--Boss)
	DailyNum          int8   // 允许每天进入次数,0表示不限制
	Physical          int8   // 每次进入消耗的体力
	BoxX              int32  // 宝箱x坐标
	BoxY              int32  // 宝箱y坐标
	AwardItem         int16  //奖励物品ID
	AwardItemNum      int16  //奖励物品数量
	AwardKey          int32  // 奖励钥匙数
	AwardExp          int32  // 奖励经验
	EnterY            int32  // 出生点y坐标
	EnterX            int32  // 出生点x坐标
	Sign              string // 资源标识
	SignWar           string // 关卡战斗资源标识
	Music             string // 音乐资源标识
	AwardLock         int32  // 通关奖励权值
	AwardCoin         int64  // 奖励铜钱
	AwardRelationship int32  // 奖励友情
	AwardBox          bool   // 是否奖励宝箱

	ParentType int8  // 关联关卡类型(0--无;1--资源关卡;)
	ParentId   int16 // 关联关卡的外键
	SubType    int8
}

func loadMissionLevel(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM mission_level ORDER BY `mission_id`,`lock`,`order` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iId := res.Map("id")
	iMissionId := res.Map("mission_id")
	iLock := res.Map("lock")
	iName := res.Map("name")
	iType := res.Map("type")
	iDailyNum := res.Map("daily_num")
	iPhysical := res.Map("physical")
	iBoxX := res.Map("box_x")
	iBoxY := res.Map("box_y")
	iAwardItem := res.Map("award_item")
	iAwardItemNum := res.Map("award_item_num")
	iAwardKey := res.Map("award_key")
	iAwardExp := res.Map("award_exp")
	iEnterY := res.Map("enter_y")
	iEnterX := res.Map("enter_x")
	iSign := res.Map("sign")
	iSignWar := res.Map("sign_war")
	iMusic := res.Map("music")
	iAwardLock := res.Map("award_lock")
	iAwardCoin := res.Map("award_coin")
	iAwardRelationship := res.Map("award_relationship")
	iAwardBox := res.Map("award_box")

	iParentType := res.Map("parent_type")
	iParentId := res.Map("parent_id")
	iSubType := res.Map("sub_type")

	var pri_id int32
	var parentType int8
	var parentId int16

	mapMissionLevel = map[int32]*MissionLevel{}
	mapSpecialLevel = map[int8]map[int16]int32{}
	mapMissionLevelByItem = map[int16]map[int32]*MissionLevel{}
	//mapMissionLevelID = make(map[int16][]int32, 0)
	mapLeveIDMisson = make(map[int32]int16, 0)
	var missionLevel *MissionLevel
	for _, row := range res.Rows {
		pri_id = row.Int32(iId)
		parentId = row.Int16(iParentId)
		parentType = row.Int8(iParentType)

		missionLevel = &MissionLevel{
			Id:                pri_id,
			MissionId:         row.Int16(iMissionId),
			Lock:              row.Int32(iLock),
			Name:              row.Str(iName),
			Type:              row.Int8(iType),
			DailyNum:          row.Int8(iDailyNum),
			Physical:          row.Int8(iPhysical),
			BoxX:              row.Int32(iBoxX),
			BoxY:              row.Int32(iBoxY),
			AwardItem:         row.Int16(iAwardItem),
			AwardItemNum:      row.Int16(iAwardItemNum),
			AwardKey:          row.Int32(iAwardKey),
			AwardExp:          row.Int32(iAwardExp),
			EnterY:            row.Int32(iEnterY),
			EnterX:            row.Int32(iEnterX),
			Sign:              row.Str(iSign),
			SignWar:           row.Str(iSignWar),
			Music:             row.Str(iMusic),
			AwardLock:         row.Int32(iAwardLock),
			AwardCoin:         row.Int64(iAwardCoin),
			AwardRelationship: row.Int32(iAwardRelationship),
			AwardBox:          row.Bool(iAwardBox),
			SubType:           row.Int8(iSubType),

			ParentType: parentType,
			ParentId:   parentId,
		}
		mapMissionLevel[pri_id] = missionLevel
		if parentType == 0 {
			if mapMissionLevelByItem[missionLevel.AwardItem] == nil {
				mapMissionLevelByItem[missionLevel.AwardItem] = map[int32]*MissionLevel{}
			}
			mapMissionLevelByItem[missionLevel.AwardItem][pri_id] = missionLevel
		}

		if parentType > 0 {
			if _, ok := mapSpecialLevel[parentType]; !ok {
				mapSpecialLevel[parentType] = map[int16]int32{}
			}
			mapSpecialLevel[parentType][parentId] = pri_id
		}

		//mapMissionLevelID[row.Int16(iMissionId)] = append(mapMissionLevelID[row.Int16(iMissionId)], pri_id)
		mapLeveIDMisson[pri_id] = row.Int16(iMissionId)
	}
}

// loadExtendLevel要在loadMissionLevel之后调用
func loadExtendLevel(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM extend_level ORDER BY `level_type`,`max_level` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iId := res.Map("id")
	iLevelType := res.Map("level_type")
	iMaxLevel := res.Map("max_level")

	var id, maxLevel int16
	var levelType int8
	mapExtendLevel = map[int8]map[int16][]int32{}

	for _, row := range res.Rows {
		id = row.Int16(iId)
		maxLevel = row.Int16(iMaxLevel)
		levelType = row.Int8(iLevelType)

		if _, ok := mapExtendLevel[levelType]; !ok {
			mapExtendLevel[levelType] = make(map[int16][]int32)
		}

		levelList := []int32{}
		for levelId, levelInfo := range mapMissionLevel {
			if levelInfo.ParentType == levelType && levelInfo.ParentId == id {
				levelList = append(levelList, levelId)
			}
		}
		mapExtendLevel[levelType][maxLevel] = levelList
	}
}

type MissionEnemy struct {
	Id              int32  //
	MissionLevelId  int32  // 副本关卡id
	MonsterNum      int8   // 怪物数量
	EnterX          int32  // 出生点x坐标
	EnterY          int32  // 出生点y坐标
	Monster1Id      int32  // 怪物1 ID
	Monster1Chance  int8   // 出现概率
	Monster2Id      int32  // 怪物2 ID
	Monster2Chance  int8   // 出现概率
	Monster3Id      int32  // 怪物3 ID
	Monster4Id      int32  // 怪物4 ID
	Monster4Chance  int8   // 出现概率
	Monster5Id      int32  // 怪物5 ID
	Monster5Chance  int8   // 出现概率
	Monster3Chance  int8   // 出现概率
	IsBoss          bool   // boss id
	Talk            string // 副本对话（怪物旁白）
	BossDir         int8   // 怪物朝向(0--左;1--右)
	BestRound       int8   // 最好的通关回合数
	ShadedMissionId int32  // 所属影之间隙
}

func loadMissionEnemy(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM mission_enemy ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iId := res.Map("id")
	iMissionLevelId := res.Map("mission_level_id")
	iMonsterNum := res.Map("monster_num")
	iEnterX := res.Map("enter_x")
	iEnterY := res.Map("enter_y")
	iMonster1Id := res.Map("monster1_id")
	iMonster1Chance := res.Map("monster1_chance")
	iMonster2Id := res.Map("monster2_id")
	iMonster2Chance := res.Map("monster2_chance")
	iMonster3Id := res.Map("monster3_id")
	iMonster4Id := res.Map("monster4_id")
	iMonster4Chance := res.Map("monster4_chance")
	iMonster5Id := res.Map("monster5_id")
	iMonster5Chance := res.Map("monster5_chance")
	iMonster3Chance := res.Map("monster3_chance")
	iIsBoss := res.Map("is_boss")
	iBossDir := res.Map("boss_dir")
	iBestRound := res.Map("best_round")
	iShadedMissionId := res.Map("shaded_mission_id")

	var pri_id, levelId int32
	mapMissionEnemy = map[int32]*MissionEnemy{}
	mapMissionLevelEnemy = map[int32][]int32{}
	for _, row := range res.Rows {
		pri_id = row.Int32(iId)
		levelId = row.Int32(iMissionLevelId)
		mapMissionLevelEnemy[levelId] = append(mapMissionLevelEnemy[levelId], pri_id)
		mapMissionEnemy[pri_id] = &MissionEnemy{
			Id:              pri_id,
			MissionLevelId:  levelId,
			MonsterNum:      row.Int8(iMonsterNum),
			EnterX:          row.Int32(iEnterX),
			EnterY:          row.Int32(iEnterY),
			Monster1Id:      row.Int32(iMonster1Id),
			Monster1Chance:  row.Int8(iMonster1Chance),
			Monster2Id:      row.Int32(iMonster2Id),
			Monster2Chance:  row.Int8(iMonster2Chance),
			Monster3Id:      row.Int32(iMonster3Id),
			Monster4Id:      row.Int32(iMonster4Id),
			Monster4Chance:  row.Int8(iMonster4Chance),
			Monster5Id:      row.Int32(iMonster5Id),
			Monster5Chance:  row.Int8(iMonster5Chance),
			Monster3Chance:  row.Int8(iMonster3Chance),
			IsBoss:          row.Bool(iIsBoss),
			BossDir:         row.Int8(iBossDir),
			BestRound:       row.Int8(iBestRound),
			ShadedMissionId: row.Int32(iShadedMissionId),
		}
	}
}

type EnemyRole struct {
	Id                  int32   // 角色ID
	Name                string  // 角色名称
	Sign                string  // 资源标识
	Level               int32   // 等级 - level
	Prop                int8    // 种族
	Health              int32   // 生命 - health
	Cultivation         int32   // 内力 - cultivation
	Speed               int32   // 速度 - speed
	Attack              int32   // 普攻 - attack
	Defence             int32   // 普防 - defence
	Dodge               float32 // 闪避 - dodge
	Hit                 float32 // 命中 - hit
	Block               float32 // 格挡 - block
	Critial             float32 // 暴击 - critial
	Toughness           float32 // 韧性
	Destroy             float32 // 破击
	CritialHurt         int32   // 必杀 – critial hurt
	Sleep               int32   // 睡眠抗性
	Dizziness           int32   // 眩晕抗性
	Random              int32   // 混乱抗性
	DisableSkill        int32   // 封魔抗性
	Poisoning           int32   // 中毒抗性
	SkillId             int16   // 绝招ID
	SkillForce          int32   // 绝招威力
	Skill2Id            int16   // 绝招2 ID
	Skill2Force         int32   // 绝招2 威力
	SunderMaxValue      int32   // 护甲值
	SunderMinHurtRate   int32   // 破甲前起始的伤害转换率（百分比）
	SunderEndHurtRate   int32   // 破甲后的伤害转换率（百分比）
	SunderEndDefendRate int32   // 破甲后减防（百分比）
	SunderAttack        int32   // 攻击破甲值
	SkillWait           int8    // 绝招蓄力回合
	ReleaseNum          int8    // 释放次数
	RecoverRoundNum     int8    // 恢复回合数
	CommonAttackNum     int8    // 入场普通攻击次数
	IsBoss              bool
	UseRhythm           int32
}

func loadEnemyRole(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM enemy_role ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iId := res.Map("id")
	iName := res.Map("name")
	iSign := res.Map("sign")
	iLevel := res.Map("level")
	iProp := res.Map("prop")
	iHealth := res.Map("health")
	iCultivation := res.Map("cultivation")
	iSpeed := res.Map("speed")
	iAttack := res.Map("attack")
	iDefence := res.Map("defence")
	iDodge := res.Map("dodge")
	iHit := res.Map("hit")
	iBlock := res.Map("block")
	iCritial := res.Map("critial")
	iToughness := res.Map("toughness")
	iDestroy := res.Map("destroy")
	iCritialHurt := res.Map("critial_hurt")
	iSleep := res.Map("sleep")
	iDizziness := res.Map("dizziness")
	iRandom := res.Map("random")
	iDisableSkill := res.Map("disable_skill")
	iPoisoning := res.Map("poisoning")
	iSkillId := res.Map("skill_id")
	iSkillForce := res.Map("skill_force")
	iSkill2Id := res.Map("skill2_id")
	iSkill2Force := res.Map("skill2_force")
	iSunderMaxValue := res.Map("sunder_max_value")
	iSunderMinHurtRate := res.Map("sunder_min_hurt_rate")
	iSunderEndHurtRate := res.Map("sunder_end_hurt_rate")
	iSunderEndDefendRate := res.Map("sunder_end_defend_rate")
	iSunderAttack := res.Map("sunder_attack")
	iSkillWait := res.Map("skill_wait")
	iReleaseNum := res.Map("release_num")
	iRecoverRoundNum := res.Map("recover_round_num")
	iCommonAttackNum := res.Map("common_attack_num")
	iIsBoss := res.Map("is_boss")

	var pri_id int32
	mapEnemyRole = map[int32]*EnemyRole{}
	for _, row := range res.Rows {
		pri_id = row.Int32(iId)

		rawProp := row.Int8(iProp)
		prop := rawProp / 2 // 恰好两个原始种族相邻并归并为一个

		mapEnemyRole[pri_id] = &EnemyRole{
			Id:                  pri_id,
			Name:                row.Str(iName),
			Sign:                row.Str(iSign),
			Level:               row.Int32(iLevel),
			Prop:                prop,
			Health:              row.Int32(iHealth),
			Cultivation:         row.Int32(iCultivation),
			Speed:               row.Int32(iSpeed),
			Attack:              row.Int32(iAttack),
			Defence:             row.Int32(iDefence),
			Dodge:               row.Float32(iDodge),
			Hit:                 row.Float32(iHit),
			Block:               row.Float32(iBlock),
			Critial:             row.Float32(iCritial),
			Toughness:           row.Float32(iToughness),
			Destroy:             row.Float32(iDestroy),
			CritialHurt:         row.Int32(iCritialHurt),
			Sleep:               row.Int32(iSleep),
			Dizziness:           row.Int32(iDizziness),
			Random:              row.Int32(iRandom),
			DisableSkill:        row.Int32(iDisableSkill),
			Poisoning:           row.Int32(iPoisoning),
			SkillId:             row.Int16(iSkillId),
			SkillForce:          row.Int32(iSkillForce),
			Skill2Id:            row.Int16(iSkill2Id),
			Skill2Force:         row.Int32(iSkill2Force),
			SunderMaxValue:      row.Int32(iSunderMaxValue),
			SunderMinHurtRate:   row.Int32(iSunderMinHurtRate),
			SunderEndHurtRate:   row.Int32(iSunderEndHurtRate),
			SunderEndDefendRate: row.Int32(iSunderEndDefendRate),
			SunderAttack:        row.Int32(iSunderAttack),
			SkillWait:           row.Int8(iSkillWait),
			ReleaseNum:          row.Int8(iReleaseNum),
			RecoverRoundNum:     row.Int8(iRecoverRoundNum),
			CommonAttackNum:     row.Int8(iCommonAttackNum),
			IsBoss:              row.Bool(iIsBoss),
			UseRhythm:           row.Int32(iReleaseNum) + row.Int32(iRecoverRoundNum),
		}
	}
}

type MissionLevelBox struct {
	Id             int64 // 主键ID
	MissionLevelId int32 // 关卡id
	Order          int8  // 品质顺序
	AwardType      int8  // 奖励类型(0--铜钱;1--道具;2--装备)
	AwardChance    int8  // 奖励概率
	AwardNum       int32 // 奖励数量
	ItemId         int32 // 物品ID(物品奖励填写)
	MustInFirst    bool  //是否为第一次通关必获奖励
}

func loadMissionLevelBox(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM mission_level_box ORDER BY `mission_level_id`, `order` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iId := res.Map("id")
	iMissionLevelId := res.Map("mission_level_id")
	iOrder := res.Map("order")
	iAwardType := res.Map("award_type")
	iAwardChance := res.Map("award_chance")
	iAwardNum := res.Map("award_num")
	iItemId := res.Map("item_id")
	iMustInFirst := res.Map("must_in_first")

	var levelId int32
	var order int8
	mapMissionLevelBox = map[int32][]*MissionLevelBox{}
	for _, row := range res.Rows {
		levelId = row.Int32(iMissionLevelId)
		order = row.Int8(iOrder)

		if _, ok := mapMissionLevelBox[levelId]; !ok {
			mapMissionLevelBox[levelId] = make([]*MissionLevelBox, 5)
		}
		mapMissionLevelBox[levelId][order-1] = &MissionLevelBox{
			Id:             row.Int64(iId),
			MissionLevelId: levelId,
			Order:          order,
			AwardType:      row.Int8(iAwardType),
			AwardChance:    row.Int8(iAwardChance),
			AwardNum:       row.Int32(iAwardNum),
			ItemId:         row.Int32(iItemId),
			MustInFirst:    row.Bool(iMustInFirst),
		}
	}
}

type MissionLevelSmallBox struct {
	Id          int32
	ItemsCount  int8
	Probability int8 // 出现概率
}

func loadMissionLevelSmallBox(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM mission_level_small_box ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iId := res.Map("id")
	iMissionLevelId := res.Map("mission_level_id")
	iProbability := res.Map("probability")
	iItemsCount := res.Map("items_kind")

	mapMissionLevelSmallBox = make(map[int32][]*MissionLevelSmallBox)
	for _, row := range res.Rows {
		missionLevelId := row.Int32(iMissionLevelId)
		mapMissionLevelSmallBox[missionLevelId] = append(mapMissionLevelSmallBox[missionLevelId], &MissionLevelSmallBox{
			Id:          row.Int32(iId),
			Probability: row.Int8(iProbability),
			ItemsCount:  row.Int8(iItemsCount),
		})
	}
}

type MissionLevelSmallBoxItems struct {
	BoxItemId   int32 // 唯一ID
	ItemId      int32 // 物品ID
	Probability int8  // 出现概率
	ItemNumber  int32 // 奖励数量
	AwardType   int8  // 奖励类型(0--铜钱;1--道具;2--装备)
}

func loadMissionLevelSmallBoxItems(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM mission_level_small_box_items ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iId := res.Map("id")
	iSmallBoxId := res.Map("small_box_id")
	iItemId := res.Map("item_id")
	iProbability := res.Map("probability")
	iItemNumber := res.Map("item_number")
	iAwardType := res.Map("award_type")

	mapMissionLevelSmallBoxItems = make(map[int32][]*MissionLevelSmallBoxItems)
	for _, row := range res.Rows {
		smallBoxId := row.Int32(iSmallBoxId)
		mapMissionLevelSmallBoxItems[smallBoxId] = append(mapMissionLevelSmallBoxItems[smallBoxId], &MissionLevelSmallBoxItems{
			BoxItemId:   row.Int32(iId),
			ItemId:      row.Int32(iItemId),
			Probability: row.Int8(iProbability),
			ItemNumber:  row.Int32(iItemNumber),
			AwardType:   row.Int8(iAwardType),
		})
	}
}

type MissionLevelMengYao struct {
	Id          int32 //关卡梦妖id
	Probability int8  //关卡梦妖出现几率
	Effect      int8  //关卡梦妖增效效果
}

func loadMissionLevelMengYao(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM mission_level_recovery_meng_yao ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}
	mId := res.Map("id")
	mLevelId := res.Map("mission_level_id")
	mProbability := res.Map("probability")
	mEffect := res.Map("my_effect")

	mapMissionLevelMengYao = make(map[int32][]*MissionLevelMengYao)
	for _, row := range res.Rows {
		levelId := row.Int32(mLevelId)
		mapMissionLevelMengYao[levelId] = append(mapMissionLevelMengYao[levelId], &MissionLevelMengYao{
			Id:          row.Int32(mId),
			Probability: row.Int8(mProbability),
			Effect:      row.Int8(mEffect),
		})
	}
}

type ShadedMission struct {
	Id             int32 // 影之间隙id
	MissionLevelId int32 // 所属关卡id
	Order          int8  // 位序
	AwardExp       int32 // 奖励经验
	AwardCoin      int32 // 奖励铜钱
	AwardItem1     int16 // 奖励物品1
	AwardItem1Num  int16 // 奖励物品1数量
	AwardItem2     int16 // 奖励物品2
	AwardItem2Num  int16 // 奖励物品2数量
	AwardItem3     int16 // 奖励物品3
	AwardItem3Num  int16 // 奖励物品3数量
}

func loadShadedMission(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("select * from `shaded_mission`;"), -1)
	if err != nil {
		panic(err)
	}
	iId := res.Map("id")
	iMissionLevelId := res.Map("mission_level_id")
	iOrder := res.Map("order")
	iAwardExp := res.Map("award_exp")
	iAwardCoin := res.Map("award_coin")
	iAwardItem1 := res.Map("award_item1")
	iAwardItem1Num := res.Map("award_item1_num")
	iAwardItem2 := res.Map("award_item2")
	iAwardItem2Num := res.Map("award_item2_num")
	iAwardItem3 := res.Map("award_item3")
	iAwardItem3Num := res.Map("award_item3_num")

	mapShadedMission = make(map[int32]*ShadedMission)
	mapMissionId2ShadedIds = make(map[int32][]int32)
	for _, row := range res.Rows {
		shadedId := row.Int32(iId)
		missionId := row.Int32(iMissionLevelId)
		order := row.Int8(iOrder)
		mapShadedMission[shadedId] = &ShadedMission{
			Id:             shadedId,
			MissionLevelId: missionId,
			Order:          order,
			AwardExp:       row.Int32(iAwardExp),
			AwardCoin:      row.Int32(iAwardCoin),
			AwardItem1:     row.Int16(iAwardItem1),
			AwardItem1Num:  row.Int16(iAwardItem1Num),
			AwardItem2:     row.Int16(iAwardItem2),
			AwardItem2Num:  row.Int16(iAwardItem2Num),
			AwardItem3:     row.Int16(iAwardItem3),
			AwardItem3Num:  row.Int16(iAwardItem3Num),
		}
		if _, had := mapMissionId2ShadedIds[missionId]; !had {
			mapMissionId2ShadedIds[missionId] = []int32{}
		}
		mapMissionId2ShadedIds[missionId] = append(mapMissionId2ShadedIds[missionId], shadedId)
	}
}

type PveLevel struct {
	Id            int16 //ID
	Floor         int16 // 关卡层数
	AwardItem     int16 // 首次通关奖励物品ID
	AwardNum      int16 // 首次通关奖励物品数量
	MosterNum     int16 // 怪物数量
	BasicAwardNum int16 // 基础奖励
	AwardFactor   int16 // 奖励系数
	Level         int16 // 要求等级
}

func loadPveLevel(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM pve_level ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iId := res.Map("id")
	iFloor := res.Map("floor")
	iAwardItem := res.Map("award_item")
	iAwardNum := res.Map("award_num")
	iMosterNum := res.Map("moster_num")
	iBasicAwardNum := res.Map("basic_award_num")
	iAwardFactor := res.Map("award_factor")
	iLevel := res.Map("level")

	var pri_id, floor int16
	mapPveLevel = map[int16]*PveLevel{}
	mapPveLevelByFloor = map[int16]*PveLevel{}
	for _, row := range res.Rows {
		pri_id = row.Int16(iId)
		floor = row.Int16(iFloor)
		mapPveLevel[pri_id] = &PveLevel{
			Id:            pri_id,
			Floor:         floor,
			AwardItem:     row.Int16(iAwardItem),
			AwardNum:      row.Int16(iAwardNum),
			MosterNum:     row.Int16(iMosterNum),
			BasicAwardNum: row.Int16(iBasicAwardNum),
			AwardFactor:   row.Int16(iAwardFactor),
			Level:         row.Int16(iLevel),
		}
		mapPveLevelByFloor[floor] = mapPveLevel[pri_id]
	}
}

type HardLevel struct {
	Id                 int16  //
	TownId             int16  // 所在城镇ID
	MissionLevelLock   int32  //要求的区域关卡权值
	HardLevelLock      int32  // 进入的难度关卡权值
	AwardHardLevelLock int32  //奖励的难度关卡权值
	Desc               string // 关卡描述
}

func loadHardLevel(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM hard_level ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iId := res.Map("id")
	iMissionLevelLock := res.Map("mission_level_lock")
	iDesc := res.Map("desc")
	iTownId := res.Map("town_id")
	iHardLevelLock := res.Map("hard_level_lock")
	iAwardHardLevelLock := res.Map("award_hard_level_lock")

	var pri_id int16
	mapHardLevel = map[int16]*HardLevel{}
	for _, row := range res.Rows {
		pri_id = row.Int16(iId)
		mapHardLevel[pri_id] = &HardLevel{
			Id:                 pri_id,
			MissionLevelLock:   row.Int32(iMissionLevelLock),
			Desc:               row.Str(iDesc),
			TownId:             row.Int16(iTownId),
			HardLevelLock:      row.Int32(iHardLevelLock),
			AwardHardLevelLock: row.Int32(iAwardHardLevelLock),
		}
	}
}

type RandAwardBox struct {
	Id             int64 //
	MissionLevelId int32 // 关卡id
	Order          int8  // 品质顺序
	AwardType      int8  // 奖励类型(0--铜钱;1--道具;2--装备)
	AwardChance    int8  // 奖励概率
	AwardNum       int32 // 奖励数量
	ItemId         int32 // 物品ID(物品奖励填写)
	MustInFirst    bool  //是否为第一次通关必获奖励
}

func loadRandAwardBox(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM random_award_box ORDER BY `mission_level_id`, `order` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iId := res.Map("id")
	iMissionLevelId := res.Map("mission_level_id")
	iOrder := res.Map("order")
	iAwardType := res.Map("award_type")
	iAwardChance := res.Map("award_chance")
	iAwardNum := res.Map("award_num")
	iItemId := res.Map("item_id")
	iMustInFirst := res.Map("must_in_first")
	var levelId int32
	var order int8
	mapRandAwardBox = map[int32][]*RandAwardBox{}
	for _, row := range res.Rows {
		levelId = row.Int32(iMissionLevelId)
		order = row.Int8(iOrder)

		mapRandAwardBox[levelId] = append(mapRandAwardBox[levelId], &RandAwardBox{
			Id:             row.Int64(iId),
			MissionLevelId: levelId,
			Order:          order,
			AwardType:      row.Int8(iAwardType),
			AwardChance:    row.Int8(iAwardChance),
			AwardNum:       row.Int32(iAwardNum),
			ItemId:         row.Int32(iItemId),
			MustInFirst:    row.Bool(iMustInFirst),
		})
	}
}

/*
var (
	mapTownStarAwards map[int32]*TownStarAwards
)
*/
type TownStarAwards struct {
	//Id int32 // ID
	TownId    int16 // 城镇ID
	BoxType   int8  // 宝箱类型 1:铜 2:银 3:金
	Totalstar int16 // 通关回合数
	Ingot     int16 // 奖励元宝
	Coins     int32 // 奖励铜钱
	Heart     int16 // 奖励爱心
	Item1Id   int16 // 物品1
	Item1Num  int16 // 物品1数量
	Item2Id   int16 // 物品2
	Item2Num  int16 // 物品2数量
	Item3Id   int16 // 物品3
	Item3Num  int16 // 物品3数量
	Item4Id   int16 // 物品4
	Item4Num  int16 // 物品4数量
	Item5Id   int16 // 物品5
	Item5Num  int16 // 物品5数量
}

func loadTownStarAwards(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM town_star_awards ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	//iId := res.Map("id")
	iTwonId := res.Map("town_id")
	iBoxtype := res.Map("box_type")
	iTotalstar := res.Map("totalstar")
	iIngot := res.Map("ingot")
	iCoins := res.Map("coins")
	iHeart := res.Map("heart")
	iItem1Id := res.Map("item1_id")
	iItem1Num := res.Map("item1_num")
	iItem2Id := res.Map("item2_id")
	iItem2Num := res.Map("item2_num")
	iItem3Id := res.Map("item3_id")
	iItem3Num := res.Map("item3_num")
	iItem4Id := res.Map("item4_id")
	iItem4Num := res.Map("item4_num")
	iItem5Id := res.Map("item5_id")
	iItem5Num := res.Map("item5_num")

	mapTownStarAwards = map[int16][]*TownStarAwards{}
	for _, row := range res.Rows {
		key := row.Int16(iTwonId)
		mapTownStarAwards[key] = append(mapTownStarAwards[key], &TownStarAwards{
			Totalstar: row.Int16(iTotalstar),
			BoxType:   row.Int8(iBoxtype),
			Ingot:     row.Int16(iIngot),
			Coins:     row.Int32(iCoins),
			Heart:     row.Int16(iHeart),
			Item1Id:   row.Int16(iItem1Id),
			Item1Num:  row.Int16(iItem1Num),
			Item2Id:   row.Int16(iItem2Id),
			Item2Num:  row.Int16(iItem2Num),
			Item3Id:   row.Int16(iItem3Id),
			Item3Num:  row.Int16(iItem3Num),
			Item4Id:   row.Int16(iItem4Id),
			Item4Num:  row.Int16(iItem4Num),
			Item5Id:   row.Int16(iItem5Id),
			Item5Num:  row.Int16(iItem5Num),
		})
	}
}

// ############### 对外接口实现 coding here ###############

func GetMissionById(id int16) *Mission {
	v, ok := mapMission[id]
	fail.When(!ok, "incorrect mission id")
	return v
}

func GetMissionLevelById(id int32) *MissionLevel {
	v, ok := mapMissionLevel[id]
	fail.When(!ok, "incorrect mission level id")
	return v
}

func GetMissionLevelByLock(lock int32) *MissionLevel {
	for key, value := range mapMissionLevel {
		if value.Lock == lock {
			return mapMissionLevel[key]
		}
	}
	return nil
}

func GetMissionLevelEnemyById(id int32) *MissionEnemy {
	v, ok := mapMissionEnemy[id]
	fail.When(!ok, "incorrect enemy id")
	return v
}

func GetEnemyRole(id int32) *EnemyRole {
	v, ok := mapEnemyRole[id]
	fail.When(!ok, "incorrect enemy id")
	return v
}

func GetLevelBoxByLevelId(id int32) []*MissionLevelBox {
	v, ok := mapMissionLevelBox[id]
	fail.When(!ok, "incorrect level box id")
	return v
}

func GetMustLevelBoxByLevelId(id int32) []*MissionLevelBox {
	v := GetLevelBoxByLevelId(id)
	result := make([]*MissionLevelBox, 0)
	for _, val := range v {
		if val.MustInFirst {
			result = append(result, val)
		}
	}
	return result
}

func GetMustRandomBoxByLevelId(id int32) []*RandAwardBox {
	v := GetRandomBoxByLevelId(id)
	result := make([]*RandAwardBox, 0)
	for _, val := range v {
		if val.MustInFirst {
			result = append(result, val)
		}
	}
	return result
}

func GetExtendLevelRequireLevel(levelType int8, levelId int32) int16 {
	for level, ids := range mapExtendLevel[levelType] {
		for _, id := range ids {
			if id == levelId {
				return level
			}
		}
	}
	fail.When(true, "GetResourceLevelRequireLevel error")
	return 0
}

func GetSpecialLevelId(parentType int8, parentId int16) int32 {
	id, ok := mapSpecialLevel[parentType][parentId]
	fail.When(!ok, "can't get special level id")
	return id
}

func GetMissionSmallBox(missionId int32) ([]*MissionLevelSmallBox, bool) {
	box, ok := mapMissionLevelSmallBox[missionId]
	return box, ok
}

func GetSmailBoxItems(boxId int32) []*MissionLevelSmallBoxItems {
	items, ok := mapMissionLevelSmallBoxItems[boxId]
	fail.When(!ok, "模板数据错误: MissionLevelSmallBoxItems")
	return items
}

func GetMissionMengYao(missionId int32) ([]*MissionLevelMengYao, bool) {
	my, ok := mapMissionLevelMengYao[missionId]
	return my, ok
}

func GetShadedMission(shadedId int32) (*ShadedMission, bool) {
	shaded, ok := mapShadedMission[shadedId]
	return shaded, ok
}

func GetShadedIdsByMissionId(missionId int32) []int32 {
	return mapMissionId2ShadedIds[missionId]
}

func GetHardLevelInfo(id int16) *HardLevel {
	return mapHardLevel[id]
}

func GetEnemyIdByMissionLevelId(id int32) []int32 {
	ids, ok := mapMissionLevelEnemy[id]
	fail.When(!ok, "没有找到关卡")
	return ids
}

func GetPetVirtualEnvLevel(id int16) (dat *PveLevel) {
	dat = mapPveLevel[id]
	fail.When(dat == nil, "找不到灵宠幻境关卡")
	return dat
}

func GetPetVirtualEnvLevelByFloor(floor int16) (dat *PveLevel) {
	dat = mapPveLevelByFloor[floor]
	fail.When(dat == nil, "找不到灵宠幻境关卡")
	return dat
}

// func CalLevelStarByScore(levelId int32, score int32) int8 {
// 	lvStar, ok := mapLevelStar[levelId]
// 	if !ok {
// 		return -1
// 	}
//
// 	if score >= lvStar.ThreeStarScore {
// 		return THREE_STAR
// 	} else if score >= lvStar.TwoStarScore {
// 		return TWO_STAR
// 	} else if score > 0 {
// 		return ONE_STAR
// 	}
// 	panic("undefine score")
// }

func GetMissionLevelByItemId(itemId int16) map[int32]*MissionLevel {
	return mapMissionLevelByItem[itemId]
}

func GetRandomBoxByLevelId(id int32) []*RandAwardBox {
	v, ok := mapRandAwardBox[id]
	fail.When(!ok, "incorrect level box id")
	return v
}

func GetTwonIDByMission(missionID int16) int16 {
	townID, ok := mapMissonLevelTown[missionID]
	if !ok {
		return 0
	}
	return townID
}

func GetMissionIDByLevelID(levelID int32) int16 {
	missionID, ok := mapLeveIDMisson[levelID]
	if !ok {
		return 0
	}
	return missionID
}

func GetMissionStarAwards(townID int16, boxType int8) (awards *TownStarAwards) {
	awardlists, ok := mapTownStarAwards[townID]
	if !ok {
		return nil
	}
	for _, v := range awardlists {
		if boxType == v.BoxType {
			return v
		}
	}
	return nil
}

func CheckMissionStarAwards(townID int16, boxType int8) int16 {
	awardlists, ok := mapTownStarAwards[townID]
	if !ok {
		return 0
	}
	for _, v := range awardlists {
		if boxType == v.BoxType {
			return v.Totalstar
		}
	}

	return 0
}

func CheckMissionLevel(levelId int32) bool {
	_, exist := mapMissionLevel[levelId]
	return exist
}
