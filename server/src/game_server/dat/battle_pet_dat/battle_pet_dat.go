package battle_pet_dat

import (
	"core/fail"
	"core/mysql"
	//"game_server/dat/mission_dat"
)

const (
	GRID_NOT_OPENED = -1
	GRID_OPENED     = 0
)

var (
	mapBattlePet      map[int16]*BattlePet
	mapLevelBattlePet map[int32]*LevelBattlePet
	mapPetIdTable     map[int32]int16
	arrPetLevelExp    []*PetLevelExpInfo
	mapPetSkillStuff  []*PetSkillStuff
	mapPetLevelInfo   map[int32][]*PetLevelInfo
)

func Load(db *mysql.Connection) {
	loadBattlePet(db)
	loadLevelBattlePet(db)
	//loadBattlePetGridLevel(db)
	//loadBattlePetGridAttribute(db)
	//loadBattlePetSoulExchange(db)
	//loadBattlePetGridUpgradePrice(db)
	loadPetLevelExpInfo(db)
	loadPetSkillStuff(db)
	loadPetLevelInfo(db)
}

type BattlePet struct {
	PetId       int32 // 灵宠ID(enemy_role)
	RoundAttack int8  // 单回合行动次数
	CostPower   int8  // 召唤时消耗精气
	LiveRound   int8  // 召唤后存活回合数
	LivePos     int8  // 召唤后出现的位置(1-前排；2-后排；3-左侧)
	Quality     int8  // 品质
	//Level       int16 //TODO 删掉这东西 看看那里有使用
	//ParentPetId int32 // 父灵宠ID(enemy_role)
	//Health      int32 // 生命
	//Attack      int32 // 攻击
	//Defence     int32 // 防御
	//Speed       int32 // 速度
	//Force       int32 // 绝招威力
	//Star        int8  // 星级 [1,5]
}

func loadBattlePet(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM battle_pet ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iId := res.Map("id")
	iPetId := res.Map("pet_id")
	iRoundAttack := res.Map("round_attack")
	iCostPower := res.Map("cost_power")
	iLiveRound := res.Map("live_round")
	iLivePos := res.Map("live_pos")
	//iParentPetId := res.Map("parent_pet_id")
	//iHealth := res.Map("health")
	//iAttack := res.Map("attack")
	//iDefence := res.Map("defence")
	//iSpeed := res.Map("speed")
	//iForce := res.Map("force")
	//iStar := res.Map("star")
	iQuality := res.Map("quality")

	var pri_id int16
	var enemyId int32
	mapBattlePet = map[int16]*BattlePet{}
	mapPetIdTable = map[int32]int16{}
	for _, row := range res.Rows {
		pri_id = row.Int16(iId)
		enemyId = row.Int32(iPetId)

		mapPetIdTable[enemyId] = pri_id

		mapBattlePet[pri_id] = &BattlePet{
			PetId:       enemyId,
			RoundAttack: row.Int8(iRoundAttack),
			CostPower:   row.Int8(iCostPower),
			LiveRound:   row.Int8(iLiveRound),
			LivePos:     row.Int8(iLivePos),
			//Level:       int16(mission_dat.GetEnemyRole(enemyId).Level),
			//ParentPetId: row.Int32(iParentPetId),
			//Health:      row.Int32(iHealth),
			//Attack:      row.Int32(iAttack),
			//Defence:     row.Int32(iDefence),
			//Speed:       row.Int32(iSpeed),
			//Force:       row.Int32(iForce),
			//Star:        row.Int8(iStar),
			Quality: row.Int8(iQuality),
		}
	}
}

type LevelBattlePet struct {
	Pet       *BattlePet // 灵宠
	Rate      int8       // 出现概率%
	LiveRound int8       // 出现后存活回合数
}

func loadLevelBattlePet(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM level_battle_pet ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iMissionEnemyId := res.Map("mission_enemy_id")
	iBattlePetId := res.Map("battle_pet_id")
	iRate := res.Map("rate")
	iLiveRound := res.Map("live_round")

	var missionEnemyId int32
	mapLevelBattlePet = map[int32]*LevelBattlePet{}
	for _, row := range res.Rows {
		missionEnemyId = row.Int32(iMissionEnemyId)
		mapLevelBattlePet[missionEnemyId] = &LevelBattlePet{
			Pet:       mapBattlePet[row.Int16(iBattlePetId)],
			Rate:      row.Int8(iRate),
			LiveRound: row.Int8(iLiveRound),
		}
	}
}

type PetLevelExpInfo struct {
	Level       int16 // 宠物等级
	Exp         int64 // 升级所需经验
	NeedSoulNum int32 // 所需灵魄数量
	MinAddExp   int64 // 最小经验加值
	MaxAddExp   int64 // 最大经验加值
}

func loadPetLevelExpInfo(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM `battle_pet_exp` ORDER BY `level` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iLevel := res.Map("level")
	iExp := res.Map("exp")
	iNeedSoulNum := res.Map("need_soul_num")
	iMinAddExp := res.Map("min_add_exp")
	iMaxAddExp := res.Map("max_add_exp")

	arrPetLevelExp = []*PetLevelExpInfo{}
	for _, row := range res.Rows {
		arrPetLevelExp = append(arrPetLevelExp, &PetLevelExpInfo{
			Level:       row.Int16(iLevel),
			Exp:         row.Int64(iExp),
			NeedSoulNum: row.Int32(iNeedSoulNum),
			MinAddExp:   row.Int64(iMinAddExp),
			MaxAddExp:   row.Int64(iMaxAddExp),
		})
	}
}

type PetSkillStuff struct {
	Level     int16 // 宠物等级
	CostCoins int64 // 花费铜钱
}

func loadPetSkillStuff(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM `battle_pet_skill_training` ORDER BY `level` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iLevel := res.Map("level")
	iCostCoins := res.Map("cost_coins")

	mapPetSkillStuff = []*PetSkillStuff{}
	for _, row := range res.Rows {
		level := row.Int16(iLevel)
		mapPetSkillStuff = append(mapPetSkillStuff, &PetSkillStuff{
			Level:     level,
			CostCoins: row.Int64(iCostCoins),
		})
	}
}

type PetLevelInfo struct {
	PetId               int32 // 宠物ID(enemy role)
	Level               int16 // 宠物等级
	Health              int32 // 生命
	Speed               int32 // 速度
	Attack              int32 // 普攻
	Defence             int32 // 普防
	SunderMaxValue      int32 // 护甲值
	SunderMinHurtRate   int32 // 破甲前起始的伤害转换率(百分比)
	SunderEndHurtRate   int32 // 破甲后的伤害转换率(百分比)
	SunderEndDefendRate int32 // 婆家后减防(百分比)
}

func loadPetLevelInfo(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM `battle_pet_level_info` ORDER BY `pet_id`, `level` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iPetId := res.Map("pet_id")
	iLevel := res.Map("level")
	iHealth := res.Map("health")
	iSpeed := res.Map("speed")
	iAttack := res.Map("attack")
	iDefence := res.Map("defence")
	iSunderMaxValue := res.Map("sunder_max_value")
	iSunderMinHurtRate := res.Map("sunder_min_hurt_rate")
	iSunderEndHurtRate := res.Map("sunder_end_hurt_rate")
	iSunderEndDefendRate := res.Map("sunder_end_defend_rate")

	mapPetLevelInfo = map[int32][]*PetLevelInfo{}
	for _, row := range res.Rows {
		petId := row.Int32(iPetId)
		level := row.Int16(iLevel)
		mapPetLevelInfo[petId] = append(mapPetLevelInfo[petId], &PetLevelInfo{
			PetId:               petId,
			Level:               level,
			Health:              row.Int32(iHealth),
			Speed:               row.Int32(iSpeed),
			Attack:              row.Int32(iAttack),
			Defence:             row.Int32(iDefence),
			SunderMaxValue:      row.Int32(iSunderMaxValue),
			SunderMinHurtRate:   row.Int32(iSunderMinHurtRate),
			SunderEndHurtRate:   row.Int32(iSunderEndHurtRate),
			SunderEndDefendRate: row.Int32(iSunderEndDefendRate),
		})
	}
}

// ############### 对外接口实现 coding here ###############

//根据灵宠ID获取灵宠信息
func GetBattlePetWithPetId(id int16) *BattlePet {
	pet, ok := mapBattlePet[id]
	fail.When(!ok, "can not found battlePet")
	return pet
}

//根据怪物ID获取灵宠信息
func GetBattlePetWithEnemyId(enemyId int32) *BattlePet {
	id, ok := mapPetIdTable[enemyId]
	fail.When(!ok, "can not found pet id")

	var pet *BattlePet
	pet, ok = mapBattlePet[id]
	fail.When(!ok, "can not found battlePet")
	return pet
}

//根据怪物ID获取灵宠战场信息（存活回合出现几率等)
func GetLevelBattlePet(enemyId int32) (*LevelBattlePet, bool) {
	levelPet, ok := mapLevelBattlePet[enemyId]
	return levelPet, ok
}

//获取宠物等级经验信息
func GetPetLevelExpInfo(level int16) *PetLevelExpInfo {
	return arrPetLevelExp[level-1]
}

//获取灵宠技能等级信息
func GetPetSkillStuff(level int16) *PetSkillStuff {
	return mapPetSkillStuff[level-1]
}

//获取宠物等级信息
func GetPetLevelInfo(petId int32, level int16) *PetLevelInfo {
	return mapPetLevelInfo[petId][level-1]
}
