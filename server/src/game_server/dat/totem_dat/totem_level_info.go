package totem_dat

import (
	"core/fail"
	"core/mysql"
	"fmt"
)

var (
	mapTotemLevelInfo map[int8]map[int8]*TotemLevelInfo //quality -> level -> data
)

type TotemLevelInfo struct {
	Health          int32   // 生命 - health
	Attack          float64 // 普攻 - attack
	Defence         float64 // 普防 - defence
	Cultivation     float64 // 内力 - cultivation
	RockRuneRate    int8    // 分解的石符文概率
	RockRuneNum     int16   // 分解的石符文数量
	JadeRuneRate    int8    // 分解的玉符文概率
	JadeRuneNum     int16   // 分解的玉符文数量
	UpgradeNeedRock int16   // 升级所需石符文概率
	UpgradeNeedJade int16   // 升级所需玉符文概率
}

func loadTotemLevelInfo(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM totem_level_info ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iQuality := res.Map("quality")
	iLevel := res.Map("level")
	iHealth := res.Map("health")
	iAttack := res.Map("attack")
	iDefence := res.Map("defence")
	iCultivation := res.Map("cultivation")
	iRockRuneRate := res.Map("rock_rune_rate")
	iRockRuneNum := res.Map("rock_rune_num")
	iJadeRuneRate := res.Map("jade_rune_rate")
	iJadeRuneNum := res.Map("jade_rune_num")
	iUpgradeNeedRock := res.Map("upgrade_need_rock")
	iUpgradeNeedJade := res.Map("upgrade_need_jade")

	var quality, level int8
	mapTotemLevelInfo = map[int8]map[int8]*TotemLevelInfo{}
	for _, row := range res.Rows {
		quality = row.Int8(iQuality)
		level = row.Int8(iLevel)
		if mapTotemLevelInfo[quality] == nil {
			mapTotemLevelInfo[quality] = map[int8]*TotemLevelInfo{}
		}
		mapTotemLevelInfo[quality][level] = &TotemLevelInfo{
			Health:          row.Int32(iHealth),
			Attack:          float64(row.Int32(iAttack)),
			Defence:         float64(row.Int32(iDefence)),
			Cultivation:     float64(row.Int32(iCultivation)),
			RockRuneRate:    row.Int8(iRockRuneRate),
			RockRuneNum:     row.Int16(iRockRuneNum),
			JadeRuneRate:    row.Int8(iJadeRuneRate),
			JadeRuneNum:     row.Int16(iJadeRuneNum),
			UpgradeNeedRock: row.Int16(iUpgradeNeedRock),
			UpgradeNeedJade: row.Int16(iUpgradeNeedJade),
		}
	}
}

func GetTotemLevelInfo(quality, level int8) *TotemLevelInfo {
	totemLevelInfoQuality := mapTotemLevelInfo[quality]
	fail.When(len(totemLevelInfoQuality) < 1, fmt.Sprintf("没有找到该品质级的阵印 %d", quality))
	totemLevelInfo, ok := totemLevelInfoQuality[level]
	if !ok {
		fail.When(true, fmt.Sprintf("没有找到该等级的阵印 %d", level))
	}
	return totemLevelInfo
}
