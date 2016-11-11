package sword_soul

import (
	"core/fail"
	"core/net"
	"fmt"
	"game_server/api/protocol/sword_soul_api"
	"game_server/dat/role_dat"
	"game_server/dat/sword_soul_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/tlog"
	"strconv"
	"strings"
)

func Upgrade(session *net.Session, in *sword_soul_api.Upgrade_In, out *sword_soul_api.Upgrade_Out) {
	state := module.State(session)
	db := state.Database

	targetId := in.TargetId

	var totalExp, beforlevel int32
	var swordlist []string
	swordarr := make(map[int32]int32)

	targetPlayerSwordSoul := db.Lookup.PlayerSwordSoul(targetId)
	targetSwordSoul := sword_soul_dat.GetSwordSoul(targetPlayerSwordSoul.SwordSoulId)
	if targetSwordSoul.TypeId == sword_soul_dat.TYPE_EXP {
		totalExp += targetPlayerSwordSoul.Exp
	} else {
		totalExp += calcTotalExp(targetSwordSoul.Quality, targetPlayerSwordSoul.Level, targetPlayerSwordSoul.Exp)
	}
	for _, srcSwordSoul := range in.SwordSouls {
		playerSwordSoulId := srcSwordSoul.Id
		playerSwordSoul := db.Lookup.PlayerSwordSoul(playerSwordSoulId)
		swordSoul := sword_soul_dat.GetSwordSoul(playerSwordSoul.SwordSoulId)
		totalExp += calcTotalExp(swordSoul.Quality, playerSwordSoul.Level, playerSwordSoul.Exp)
		swordarr[int32(playerSwordSoul.SwordSoulId)] += 1
		db.Delete.PlayerSwordSoul(playerSwordSoul)
	}

	beforlevel = int32(targetPlayerSwordSoul.Level)
	if targetSwordSoul.TypeId == sword_soul_dat.TYPE_EXP {
		targetPlayerSwordSoul.Exp = totalExp
	} else {
		maxLevel := maxLevelForPlayer(module.Role.GetMainRole(db).Level)
		level, exp := calcLevelAndExp(targetSwordSoul.Quality, totalExp, maxLevel)
		targetPlayerSwordSoul.Level = level
		targetPlayerSwordSoul.Exp = exp
	}
	db.Update.PlayerSwordSoul(targetPlayerSwordSoul)

	for k, v := range swordarr {
		swordlist = append(swordlist, strconv.FormatInt(int64(k), 10)+" "+strconv.FormatInt(int64(v), 10))
	}

	tlog.PlayerSwordUpgradeFlowLog(db, int32(targetPlayerSwordSoul.SwordSoulId), beforlevel, int32(targetPlayerSwordSoul.Level), strings.Join(swordlist, " "))

	out.Id = targetPlayerSwordSoul.Id
	out.SwordSoulId = targetPlayerSwordSoul.SwordSoulId
	out.Exp = targetPlayerSwordSoul.Exp
	out.Level = targetPlayerSwordSoul.Level
	tlog.PlayerSystemModuleFlowLog(db, tlog.SMT_SWORD_SOUL)
}

func setswordsoullevel(db *mdb.Database, roleid int8, swordpos int8, swordlevel int8) {
	var roleinfo *mdb.PlayerRole = nil
	db.Select.PlayerRole(func(row *mdb.PlayerRoleRow) {
		if row.RoleId() == roleid {
			roleinfo = row.GoObject()
			row.Break()
		}
	})
	fail.When(roleinfo == nil, fmt.Sprintf("role:%d does not exist", roleid))
	maxlevel := sword_soul_dat.SwordSoulMaxLevel(roleinfo.Level)
	if swordlevel > maxlevel {
		swordlevel = maxlevel
	}
	hasEquiped := false

	db.Select.PlayerSwordSoulEquipment(func(row *mdb.PlayerSwordSoulEquipmentRow) {
		if row.RoleId() == roleid {
			hasEquiped = true
			swordid := getswordbypos(row.GoObject(), swordpos)
			fail.When(swordid == 0, fmt.Sprintf("sword pos %d not found in role:%d", swordpos, roleid))
			playerSword := db.Lookup.PlayerSwordSoul(swordid)
			fail.When(playerSword.Level > swordlevel, "set level must great than player sword current level")
			playerSword.Level = swordlevel
			playerSword.Exp = 0
			db.Update.PlayerSwordSoul(playerSword)
			row.Break()
		}
	})
	fail.When(!hasEquiped, "no sword equiped")
}

func calcTotalExp(qualityId int8, level int8, exp int32) (totalExp int32) {
	quality := sword_soul_dat.GetQuality(qualityId)
	totalExp = quality.LevelFullExp[level-1] + exp
	return totalExp
}

func calcLevelAndExp(qualityId int8, totalExp int32, maxLevel int8) (level int8, exp int32) {
	quality := sword_soul_dat.GetQuality(qualityId)

	for i, fullExp := range quality.LevelFullExp {

		if maxLevel == int8(i) {
			level = maxLevel
			exp = totalExp - quality.LevelFullExp[i-1]
			return level, exp
		}

		if fullExp > totalExp {
			level = int8(i)
			if level == FULL_LEVEL {
				exp = 0
			} else {
				exp = totalExp - quality.LevelFullExp[i-1]
			}
			return level, exp
		}
	}

	// 其余情况为满级
	level = FULL_LEVEL
	exp = 0

	return level, exp
}

func maxLevelForPlayer(mainRoleLevel int16) (maxLevel int8) {

	if mainRoleLevel <= 20 {
		maxLevel = 4
	} else if mainRoleLevel <= 40 {
		maxLevel = 8
	} else if mainRoleLevel <= 60 {
		maxLevel = 12
	} else if mainRoleLevel <= 80 {
		maxLevel = 16
	} else if mainRoleLevel <= role_dat.MAX_ROLE_LEVEL {
		maxLevel = 20
	} else {
		panic("undefine level")
	}
	return maxLevel
}

func getswordbypos(playerswordequip *mdb.PlayerSwordSoulEquipment, swordpos int8) int64 {
	switch swordpos {
	case 1:
		return playerswordequip.Pos1
	case 2:
		return playerswordequip.Pos2
	case 3:
		return playerswordequip.Pos3
	case 4:
		return playerswordequip.Pos4
	case 5:
		return playerswordequip.Pos5
	case 6:
		return playerswordequip.Pos6
	case 7:
		return playerswordequip.Pos7
	case 8:
		return playerswordequip.Pos8
	case 9:
		return playerswordequip.Pos9
	default:
		return 0
	}
}
