// 角色等级信息

package role_dat

import (
	"core/mysql"
)

var (
	g_RoleLevelInfos map[int8][]*RoleLevelInfo // key: role_id
)

// 角色等级信息
type RoleLevelInfo struct {
	RoleID              int8    // 角色模板ID
	Level               int16   // 等级
	Health              int32   // 生命
	Attack              float64 // 普攻
	Defence             float64 // 普防
	Cultivation         float64 // 内力
	Speed               float64 // 速度
	Critial             float64 // 暴击
	Block               float64 // 格挡
	Hit                 float64 // 命中
	Dodge               float64 // 闪避
	CritialHurt         float64 // 暴击伤害,必杀
	Sleep               int32   // 睡眠抗性
	Dizziness           int32   // 眩晕抗性
	Random              int32   // 混乱抗性
	DisableSkill        int32   // 封魔抗性
	Poisoning           int32   // 中毒抗性
	MaxPower            int32   // 精气上限
	InitPower           int32   // 初始精气
	SunderMaxValue      int32   // 护甲值
	SunderHurtRate      int32   // 破甲前起始的伤害转换率（百分比)
	SunderEndHurtRate   int32   // 破甲后的伤害转换率（百分比）
	SunderRoundNum      int8    // 破甲持续回合数
	SunderDizziness     int8    // 破甲后眩晕回合数
	SunderEndDefendRate int32   // 破甲后减防（百分比）
}

func loadPlayerRoles(db *mysql.Connection) {
	sql := "select * from role_level order by `role_id`, `level`"
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}

	g_RoleLevelInfos = make(map[int8][]*RoleLevelInfo, 5)

	iRoleID := res.Map("role_id")
	iLevel := res.Map("level")
	iHealth := res.Map("health")
	iAttack := res.Map("attack")
	iDefence := res.Map("defence")
	iCultivation := res.Map("cultivation")
	iSpeed := res.Map("speed")
	iCritial := res.Map("critial")
	iBlock := res.Map("block")
	iHit := res.Map("hit")
	iDodge := res.Map("dodge")
	iCritialHurt := res.Map("critial_hurt")
	iSleep := res.Map("sleep")
	iDizziness := res.Map("dizziness")
	iRandom := res.Map("random")
	iDisableSkill := res.Map("disable_skill")
	iPoisoning := res.Map("poisoning")
	iMaxPower := res.Map("max_power")
	iInitPower := res.Map("init_power")
	iSunderMaxValue := res.Map("sunder_max_value")
	iSunderHurtRate := res.Map("sunder_hurt_rate")
	iSunderEndHurtRate := res.Map("sunder_end_hurt_rate")
	iSunderRoundNum := res.Map("sunder_round_num")
	iSunderDizziness := res.Map("sunder_dizziness")
	iSunderEndDefendRate := res.Map("sunder_end_defend_rate")

	for _, row := range res.Rows {
		roleID := row.Int8(iRoleID)

		g_RoleLevelInfos[roleID] = append(g_RoleLevelInfos[roleID], &RoleLevelInfo{
			RoleID:              row.Int8(iRoleID),
			Level:               row.Int16(iLevel),
			Health:              row.Int32(iHealth),
			Attack:              row.Float64(iAttack),
			Defence:             row.Float64(iDefence),
			Cultivation:         row.Float64(iCultivation),
			Speed:               row.Float64(iSpeed),
			Critial:             row.Float64(iCritial),
			Block:               row.Float64(iBlock),
			Hit:                 row.Float64(iHit),
			Dodge:               row.Float64(iDodge),
			CritialHurt:         row.Float64(iCritialHurt),
			Sleep:               row.Int32(iSleep),
			Dizziness:           row.Int32(iDizziness),
			Random:              row.Int32(iRandom),
			DisableSkill:        row.Int32(iDisableSkill),
			Poisoning:           row.Int32(iPoisoning),
			MaxPower:            row.Int32(iMaxPower),
			InitPower:           row.Int32(iInitPower),
			SunderMaxValue:      row.Int32(iSunderMaxValue),
			SunderHurtRate:      row.Int32(iSunderHurtRate),
			SunderEndHurtRate:   row.Int32(iSunderEndHurtRate),
			SunderRoundNum:      row.Int8(iSunderRoundNum),
			SunderDizziness:     row.Int8(iSunderDizziness),
			SunderEndDefendRate: row.Int32(iSunderEndDefendRate),
		})
	}
}
