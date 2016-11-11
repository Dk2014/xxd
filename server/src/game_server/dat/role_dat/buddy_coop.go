package role_dat

import (
	"core/fail"
	"core/mysql"
	"fmt"
)

var (
	arrayBuddyCooperation    []*BuddyCooperation
	arrayMainRoleCooperation []*MainRoleCooperation
	dummyMainRoleCooperation = &MainRoleCooperation{}
	//buddyCoopeartionPartner map[int8]map[int8]int8 // 主伙伴 -> 协力伙伴
)

type BuddyCooperation struct {
	Id                     int16
	RequireFriendshipLevel int32 // 要求羁绊等级
	RoleId1                int8  // 伙伴1
	RoleId2                int8  // 伙伴2
	Health                 int32 // 生命 - health
	Attack                 int32 // 普攻 - attack
	Defence                int32 // 普防 - defence
	SkillLevel             int32 // 技能熟练度
	Speed                  int32 // 速度
	Cultivation            int32 // 内力
	Sunder                 int32 // 护甲
	DodgeLevel             int32 // 闪避等级
	HitLevel               int32 // 命中等级
	BlockLevel             int32 // 格挡等级
	TenacityLevel          int32 // 韧性等级
	DestroyLevel           int32 // 破击等级
	CriticalLevel          int32 // 暴击等级
	GhostPower             int32 // 初始魂力
	CriticalHurtLevel      int32 // 必杀等级
}

func loadBuddyCooperation(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM buddy_cooperation "), -1)
	if err != nil {
		panic(err)
	}

	iId := res.Map("id")
	iRequireFriendshipLevel := res.Map("require_friendship_level")
	iRoleId1 := res.Map("role_id1")
	iRoleId2 := res.Map("role_id2")
	iHealth := res.Map("health")
	iAttack := res.Map("attack")
	iDefence := res.Map("defence")
	iSkillLevel := res.Map("skill_level")
	iSpeed := res.Map("speed")
	iCultivation := res.Map("cultivation")
	iSunder := res.Map("sunder")
	iDodgeLevel := res.Map("dodge_level")
	iHitLevel := res.Map("hit_level")
	iBlockLevel := res.Map("block_level")
	iTenacityLevel := res.Map("tenacity_level")
	iDestroyLevel := res.Map("destroy_level")
	iCriticalLevel := res.Map("critical_level")
	iCriticalHurtLevel := res.Map("critical_hurt_level")
	iGhostPower := res.Map("ghost_power")

	var role1, role2 int8
	arrayBuddyCooperation = nil
	for _, row := range res.Rows {
		role1 = row.Int8(iRoleId1)
		role2 = row.Int8(iRoleId2)

		arrayBuddyCooperation = append(arrayBuddyCooperation, &BuddyCooperation{
			Id: row.Int16(iId),
			RequireFriendshipLevel: int32(row.Int16(iRequireFriendshipLevel)),
			RoleId1:                role1,
			RoleId2:                role2,
			Health:                 row.Int32(iHealth),
			Attack:                 row.Int32(iAttack),
			Defence:                row.Int32(iDefence),
			SkillLevel:             row.Int32(iSkillLevel),
			Speed:                  row.Int32(iSpeed),
			Cultivation:            row.Int32(iCultivation),
			Sunder:                 row.Int32(iSunder),
			DodgeLevel:             row.Int32(iDodgeLevel),
			HitLevel:               row.Int32(iHitLevel),
			BlockLevel:             row.Int32(iBlockLevel),
			TenacityLevel:          row.Int32(iTenacityLevel),
			DestroyLevel:           row.Int32(iDestroyLevel),
			CriticalLevel:          row.Int32(iCriticalLevel),
			CriticalHurtLevel:      row.Int32(iCriticalHurtLevel),
			GhostPower:             row.Int32(iGhostPower),
		})
	}
}

type MainRoleCooperation struct {
	BuddyNum    int8  // 要求伙伴人数
	Health      int32 // 生命 - health
	Attack      int32 // 普攻 - attack
	Defence     int32 // 普防 - defence
	Speed       int32 // 速度
	Cultivation int32 // 内力
}

func loadMainRoleCooperation(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM main_role_cooperation ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iBuddyNum := res.Map("buddy_num")
	iHealth := res.Map("health")
	iAttack := res.Map("attack")
	iDefence := res.Map("defence")
	iSpeed := res.Map("speed")
	iCultivation := res.Map("cultivation")

	arrayMainRoleCooperation = nil
	for _, row := range res.Rows {
		arrayMainRoleCooperation = append(arrayMainRoleCooperation, &MainRoleCooperation{
			BuddyNum:    row.Int8(iBuddyNum),
			Health:      row.Int32(iHealth),
			Attack:      row.Int32(iAttack),
			Defence:     row.Int32(iDefence),
			Speed:       row.Int32(iSpeed),
			Cultivation: row.Int32(iCultivation),
		})
	}
}

// ############### 对外接口实现 coding here ###############

func GetBuddyCooperation(coopId int16) *BuddyCooperation {
	for _, coop := range arrayBuddyCooperation {
		if coop.Id == coopId {
			return coop
		}

	}
	fail.When(true, fmt.Sprintf("unknow buddy cooperation %d ", coopId))
	return nil
}

func GetBuddyCooperationGroup(role int8) (coopIds []int16) {
	for _, coop := range arrayBuddyCooperation {
		if coop.RoleId1 == role || coop.RoleId2 == role {
			coopIds = append(coopIds, coop.Id)
		}

	}
	return coopIds
}

func GetMainRoleCooperation(roleNum int8) (retCoop *MainRoleCooperation) {
	retCoop = dummyMainRoleCooperation
	for _, coop := range arrayMainRoleCooperation {
		if roleNum >= coop.BuddyNum {
			retCoop = coop
		} else {
			break
		}

	}
	return retCoop
}
