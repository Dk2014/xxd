package monster_property_addition_dat

import (
	"core/fail"
	"core/mysql"
)

var (
	mapMonsterPropertyAddition map[int32]map[int32]*MonsterPropertyAddition
	emptyMonsterPropertyAdd    MonsterPropertyAddition
)

func Load(db *mysql.Connection) {
	loadMonsterPropertyAddition(db)
}

type MonsterPropertyAddition struct {
	MonsterId           int32   // 怪物ID
	Level               int32   // 等级
	Health              int32   // 生命
	Attack              int32   // 攻击
	Defence             int32   // 防御
	Cultivation         int32   // 内力
	Speed               int32   // 速度
	SunderMaxValue      int32   // 护甲上限
	SunderMinHurtRate   int32   // 破甲前百分比
	SunderEndHurtRate   int32   // 破甲后百分比
	SunderEndDefendRate int32   // 破甲后减防百分比
	SunderAttack        int32   // 攻击破甲值
	Critial             float32 // 暴击%
	Block               float32 // 格挡%
	Hit                 float32 // 命中%
	Dodge               float32 // 闪避%
	CritialHurt         int32   // 必杀%
	Toughness           float32 // 韧性%
	Destroy             float32 // 破击%
	Sleep               int32   // 睡眠抗性%
	Dizziness           int32   // 眩晕抗性%
	Random              int32   // 混乱抗性%
	DisableSkill        int32   // 封魔抗性%
	Poisoning           int32   // 中毒抗性%
}

func loadMonsterPropertyAddition(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM monster_property_addition ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iMonsterId := res.Map("monster_id")
	iLevel := res.Map("level")
	iHealth := res.Map("health")
	iAttack := res.Map("attack")
	iDefence := res.Map("defence")
	iCultivation := res.Map("cultivation")
	iSpeed := res.Map("speed")
	iSunderMaxValue := res.Map("sunder_max_value")
	iSunderMinHurtRate := res.Map("sunder_min_hurt_rate")
	iSunderEndHurtRate := res.Map("sunder_end_hurt_rate")
	iSunderEndDefendRate := res.Map("sunder_end_defend_rate")
	iSunderAttack := res.Map("sunder_attack")
	iCritial := res.Map("critial")
	iBlock := res.Map("block")
	iHit := res.Map("hit")
	iDodge := res.Map("dodge")
	iCritialHurt := res.Map("critial_hurt")
	iToughness := res.Map("toughness")
	iDestroy := res.Map("destroy")
	iSleep := res.Map("sleep")
	iDizziness := res.Map("dizziness")
	iRandom := res.Map("random")
	iDisableSkill := res.Map("disable_skill")
	iPoisoning := res.Map("poisoning")

	var indexId int32
	var indexLevel int32
	mapMonsterPropertyAddition = map[int32]map[int32]*MonsterPropertyAddition{}
	for _, row := range res.Rows {
		indexId = row.Int32(iMonsterId)
		indexLevel = row.Int32(iLevel)
		if mapMonsterPropertyAddition[indexId] == nil {
			mapMonsterPropertyAddition[indexId] = map[int32]*MonsterPropertyAddition{}
		}
		if indexLevel > 0 {

			mapMonsterPropertyAddition[indexId][indexLevel] = &MonsterPropertyAddition{
				MonsterId:           row.Int32(iMonsterId),
				Level:               row.Int32(iLevel),
				Health:              row.Int32(iHealth),
				Attack:              row.Int32(iAttack),
				Defence:             row.Int32(iDefence),
				Cultivation:         row.Int32(iCultivation),
				Speed:               row.Int32(iSpeed),
				SunderMaxValue:      row.Int32(iSunderMaxValue),
				SunderMinHurtRate:   row.Int32(iSunderMinHurtRate),
				SunderEndHurtRate:   row.Int32(iSunderEndHurtRate),
				SunderEndDefendRate: row.Int32(iSunderEndDefendRate),
				SunderAttack:        row.Int32(iSunderAttack),
				Critial:             row.Float32(iCritial),
				Block:               row.Float32(iBlock),
				Hit:                 row.Float32(iHit),
				Dodge:               row.Float32(iDodge),
				CritialHurt:         row.Int32(iCritialHurt),
				Toughness:           row.Float32(iToughness),
				Destroy:             row.Float32(iDestroy),
				Sleep:               row.Int32(iSleep),
				Dizziness:           row.Int32(iDizziness),
				Random:              row.Int32(iRandom),
				DisableSkill:        row.Int32(iDisableSkill),
				Poisoning:           row.Int32(iPoisoning),
			}

		} else {
			fail.When(true, "wrong level")
		}
	}
}

// ############### 对外接口实现 coding here ###############

func GetMonsterPropertyAdd(monsterId int32, level int32) (monster *MonsterPropertyAddition) {
	monsterConfig := mapMonsterPropertyAddition[monsterId]
	if monsterConfig == nil {
		return &emptyMonsterPropertyAdd
	}
	if monster, ok := monsterConfig[level]; ok {
		return monster
	}
	return &emptyMonsterPropertyAdd
}
