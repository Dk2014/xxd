package sword_soul

import (
	"core/fail"
	"core/net"
	"game_server/api/protocol/sword_soul_api"
	"game_server/dat/sword_soul_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/tlog"
)

func Info(session *net.Session, out *sword_soul_api.Info_Out) {
	state := module.State(session)
	db := state.Database

	db.Select.PlayerSwordSoul(func(row *mdb.PlayerSwordSoulRow) {
		out.SwordSouls = append(out.SwordSouls, sword_soul_api.Info_Out_SwordSouls{
			Id:          row.Id(),
			SwordSoulId: row.SwordSoulId(),
			Exp:         row.Exp(),
			Level:       row.Level(),
		})
	})

	db.Select.PlayerSwordSoulEquipment(func(row *mdb.PlayerSwordSoulEquipmentRow) {
		out.RoleEquip = append(out.RoleEquip, sword_soul_api.Info_Out_RoleEquip{
			RoleId: row.RoleId(),
			Pos1Id: row.Pos1(),
			Pos2Id: row.Pos2(),
			Pos3Id: row.Pos3(),
			Pos4Id: row.Pos4(),
			Pos5Id: row.Pos5(),
			Pos6Id: row.Pos6(),
			Pos7Id: row.Pos7(),
			Pos8Id: row.Pos8(),
			Pos9Id: row.Pos9(),
		})
	})

	playerSwordSoulState := db.Lookup.PlayerSwordSoulState(state.PlayerId)
	if fixSwordSoulIngotNum(playerSwordSoulState) {
		db.Update.PlayerSwordSoulState(playerSwordSoulState)
	}
	//updateSwordSoulNum(state, playerSwordSoulState)

	out.BoxState = playerSwordSoulState.BoxState

	out.Num = playerSwordSoulState.Num
	out.IngotNum = playerSwordSoulState.IngotNum
	out.CdTime = playerSwordSoulState.UpdateTime + sword_soul_dat.RECOVERY_TIME
	if out.Num >= sword_soul_dat.MAX_DRAW_NUM {
		out.CdTime = 0
	}

	return
}

// 装备
func Equip(session *net.Session, in *sword_soul_api.Equip_In) {
	state := module.State(session)
	tlog.PlayerSystemModuleFlowLog(state.Database, tlog.SMT_SWORD_SOUL)

	playerSwordSoulId := in.FromId
	roleId := in.RoleId
	equipPos := in.EquipPos

	roleLevel := module.Role.GetBuddyRoleInTeam(state.Database, roleId).Level
	fail.When(!isEquipPosOpen(roleLevel, equipPos), "equip pos is not openned")

	equipSwordSoulByPos(state, roleId, equipPos, playerSwordSoulId)
}

func isEquipPosOpen(level int16, equipPos sword_soul_api.EquipPos) bool {
	return sword_soul_dat.GetOpenedPosNum(level) >= int(equipPos+1)
}

func equipSwordSoulByPos(state *module.SessionState, roleId int8, equipPos sword_soul_api.EquipPos, playerSwordSoulId int64) {
	db := state.Database

	playerSwordSoul := db.Lookup.PlayerSwordSoul(playerSwordSoulId)

	swordSoul := sword_soul_dat.GetSwordSoul(playerSwordSoul.SwordSoulId)
	fail.When(swordSoul.TypeId == sword_soul_dat.TYPE_EXP, "不能装备经验剑心")

	var playerSwordSoulEquipment *mdb.PlayerSwordSoulEquipment
	db.Select.PlayerSwordSoulEquipment(func(row *mdb.PlayerSwordSoulEquipmentRow) {
		if row.RoleId() == roleId {
			playerSwordSoulEquipment = row.GoObject()
			row.Break()
		}
	})

	// 检查是否有装备同类型的剑心
	typeBit := int64(1 << uint(swordSoul.TypeId))
	fail.When(playerSwordSoulEquipment.TypeBits&typeBit == typeBit, "already equip same type sword soul")

	// 标记类型
	playerSwordSoulEquipment.TypeBits += typeBit

	// 标记装备
	playerSwordSoul.Pos = EQUIPPED

	// 设置用户装备位信息
	setEquipIdByPos(playerSwordSoulEquipment, equipPos, playerSwordSoulId)

	// 更新用户数据
	db.Update.PlayerSwordSoul(playerSwordSoul)
	db.Update.PlayerSwordSoulEquipment(playerSwordSoulEquipment)
}

// 卸下
func Unequip(session *net.Session, in *sword_soul_api.Unequip_In) {
	state := module.State(session)
	db := state.Database

	tlog.PlayerSystemModuleFlowLog(db, tlog.SMT_SWORD_SOUL)

	roleId := in.RoleId
	playerSwordSoulId := in.FromId

	playerSwordSoul := db.Lookup.PlayerSwordSoul(playerSwordSoulId)

	swordSoul := sword_soul_dat.GetSwordSoul(playerSwordSoul.SwordSoulId)

	playerSwordSoul.Pos = INBAG

	var playerSwordSoulEquipment *mdb.PlayerSwordSoulEquipment
	db.Select.PlayerSwordSoulEquipment(func(row *mdb.PlayerSwordSoulEquipmentRow) {
		if row.RoleId() == roleId {
			playerSwordSoulEquipment = row.GoObject()
			row.Break()
		}
	})

	// 移除类型标记
	typeBit := int64(1 << uint(swordSoul.TypeId))
	fail.When(playerSwordSoulEquipment.TypeBits&typeBit != typeBit, "no this sword soul type")

	playerSwordSoulEquipment.TypeBits -= typeBit

	// 设置用户装备位id
	setEquipIdById(playerSwordSoulEquipment, playerSwordSoulId, 0)

	// 更新用户数据
	db.Update.PlayerSwordSoul(playerSwordSoul)
	db.Update.PlayerSwordSoulEquipment(playerSwordSoulEquipment)
}

func setEquipIdByPos(playerSwordSoulEquipment *mdb.PlayerSwordSoulEquipment, equipPos sword_soul_api.EquipPos, playerSwordSoulId int64) {
	switch equipPos {
	case sword_soul_api.EQUIP_POS_POS1:
		playerSwordSoulEquipment.Pos1 = playerSwordSoulId
	case sword_soul_api.EQUIP_POS_POS2:
		playerSwordSoulEquipment.Pos2 = playerSwordSoulId
	case sword_soul_api.EQUIP_POS_POS3:
		playerSwordSoulEquipment.Pos3 = playerSwordSoulId
	case sword_soul_api.EQUIP_POS_POS4:
		playerSwordSoulEquipment.Pos4 = playerSwordSoulId
	case sword_soul_api.EQUIP_POS_POS5:
		playerSwordSoulEquipment.Pos5 = playerSwordSoulId
	case sword_soul_api.EQUIP_POS_POS6:
		playerSwordSoulEquipment.Pos6 = playerSwordSoulId
	case sword_soul_api.EQUIP_POS_POS7:
		playerSwordSoulEquipment.Pos7 = playerSwordSoulId
	case sword_soul_api.EQUIP_POS_POS8:
		playerSwordSoulEquipment.Pos8 = playerSwordSoulId
	case sword_soul_api.EQUIP_POS_POS9:
		playerSwordSoulEquipment.Pos9 = playerSwordSoulId
	}
}

func setEquipIdById(playerSwordSoulEquipment *mdb.PlayerSwordSoulEquipment, curPlayerSwordSoulId int64, newPlayerSwordSoulId int64) {
	switch curPlayerSwordSoulId {
	case playerSwordSoulEquipment.Pos1:
		playerSwordSoulEquipment.Pos1 = newPlayerSwordSoulId
	case playerSwordSoulEquipment.Pos2:
		playerSwordSoulEquipment.Pos2 = newPlayerSwordSoulId
	case playerSwordSoulEquipment.Pos3:
		playerSwordSoulEquipment.Pos3 = newPlayerSwordSoulId
	case playerSwordSoulEquipment.Pos4:
		playerSwordSoulEquipment.Pos4 = newPlayerSwordSoulId
	case playerSwordSoulEquipment.Pos5:
		playerSwordSoulEquipment.Pos5 = newPlayerSwordSoulId
	case playerSwordSoulEquipment.Pos6:
		playerSwordSoulEquipment.Pos6 = newPlayerSwordSoulId
	case playerSwordSoulEquipment.Pos7:
		playerSwordSoulEquipment.Pos7 = newPlayerSwordSoulId
	case playerSwordSoulEquipment.Pos8:
		playerSwordSoulEquipment.Pos8 = newPlayerSwordSoulId
	case playerSwordSoulEquipment.Pos9:
		playerSwordSoulEquipment.Pos9 = newPlayerSwordSoulId
	}
}

// 装备拦内的移动到空的装备拦格子
func EquipPosChange(session *net.Session, in *sword_soul_api.EquipPosChange_In) {
	state := module.State(session)
	db := state.Database

	tlog.PlayerSystemModuleFlowLog(db, tlog.SMT_SWORD_SOUL)

	roleId := in.RoleId
	playerSwordSoulId := in.FromId
	toEquipPos := in.ToPos

	roleLevel := module.Role.GetBuddyRoleInTeam(state.Database, roleId).Level
	fail.When(!isEquipPosOpen(roleLevel, toEquipPos), "equip pos is not openned")

	var playerSwordSoulEquipment *mdb.PlayerSwordSoulEquipment
	db.Select.PlayerSwordSoulEquipment(func(row *mdb.PlayerSwordSoulEquipmentRow) {
		if row.RoleId() == roleId {
			playerSwordSoulEquipment = row.GoObject()
			row.Break()
		}
	})

	equips := getEquips(playerSwordSoulEquipment)

	fromEquipPos := getEquipPosById(playerSwordSoulEquipment, playerSwordSoulId)

	// 交换位置
	*equips[toEquipPos], *equips[fromEquipPos] = *equips[fromEquipPos], *equips[toEquipPos]

	// 更新用户装备数据
	db.Update.PlayerSwordSoulEquipment(playerSwordSoulEquipment)
}

func getEquipPosById(playerSwordSoulEquipment *mdb.PlayerSwordSoulEquipment, playerSwordSoulId int64) (equipPos sword_soul_api.EquipPos) {
	switch playerSwordSoulId {
	case playerSwordSoulEquipment.Pos1:
		return sword_soul_api.EQUIP_POS_POS1
	case playerSwordSoulEquipment.Pos2:
		return sword_soul_api.EQUIP_POS_POS2
	case playerSwordSoulEquipment.Pos3:
		return sword_soul_api.EQUIP_POS_POS3
	case playerSwordSoulEquipment.Pos4:
		return sword_soul_api.EQUIP_POS_POS4
	case playerSwordSoulEquipment.Pos5:
		return sword_soul_api.EQUIP_POS_POS5
	case playerSwordSoulEquipment.Pos6:
		return sword_soul_api.EQUIP_POS_POS6
	case playerSwordSoulEquipment.Pos7:
		return sword_soul_api.EQUIP_POS_POS7
	case playerSwordSoulEquipment.Pos8:
		return sword_soul_api.EQUIP_POS_POS8
	case playerSwordSoulEquipment.Pos9:
		return sword_soul_api.EQUIP_POS_POS9
	}
	fail.When(true, "this playerSwordSoul not equipped")
	return
}

// 背包是否已满
func IsBagFull(session *net.Session, out *sword_soul_api.IsBagFull_Out) {
	state := module.State(session)
	out.IsFull = firstEmptyPos(state) == BAG_FULL
}

// 返回背包空位数量
func EmptyPosNum(session *net.Session, out *sword_soul_api.EmptyPosNum_Out) {
	state := module.State(session)
	out.EmptyPosNum = emptyPosNum(state)
}

func Swap(session *net.Session, in *sword_soul_api.Swap_In) {
	state := module.State(session)
	db := state.Database
	tlog.PlayerSystemModuleFlowLog(db, tlog.SMT_SWORD_SOUL)

	fromId := in.FromId
	toId := in.ToId
	roleId := in.RoleId

	fromPlayerSwordSoul := db.Lookup.PlayerSwordSoul(fromId)
	toPlayerSwordSoul := db.Lookup.PlayerSwordSoul(toId)

	fromSwordSoul := sword_soul_dat.GetSwordSoul(fromPlayerSwordSoul.SwordSoulId)
	toSwordSoul := sword_soul_dat.GetSwordSoul(toPlayerSwordSoul.SwordSoulId)

	fromBit := int64(1 << uint(fromSwordSoul.TypeId))
	toBit := int64(1 << uint(toSwordSoul.TypeId))

	var playerSwordSoulEquipment *mdb.PlayerSwordSoulEquipment
	db.Select.PlayerSwordSoulEquipment(func(row *mdb.PlayerSwordSoulEquipmentRow) {
		if row.RoleId() == roleId {
			playerSwordSoulEquipment = row.GoObject()
			row.Break()
		}
	})

	switch {
	case fromPlayerSwordSoul.Pos == EQUIPPED && toPlayerSwordSoul.Pos == EQUIPPED:
		//装备位置内交换
		equips := getEquips(playerSwordSoulEquipment)

		fromEquipPos := getEquipPosById(playerSwordSoulEquipment, fromId)
		toEquipPos := getEquipPosById(playerSwordSoulEquipment, toId)

		// 交换位置
		*equips[toEquipPos], *equips[fromEquipPos] = *equips[fromEquipPos], *equips[toEquipPos]
		// 更新用户装备数据
		db.Update.PlayerSwordSoulEquipment(playerSwordSoulEquipment)

	case fromPlayerSwordSoul.Pos != EQUIPPED && toPlayerSwordSoul.Pos == EQUIPPED:
		//将背包的的剑心和装备位置上的剑心交换
		fail.When(fromSwordSoul.TypeId == sword_soul_dat.TYPE_EXP, "不能装备经验剑心")

		//去装装备位目标剑心标记
		playerSwordSoulEquipment.TypeBits -= toBit
		fail.When(playerSwordSoulEquipment.TypeBits&fromBit == fromBit, "其他装备位已装备该类型剑心")
		playerSwordSoulEquipment.TypeBits += fromBit

		equips := getEquips(playerSwordSoulEquipment)

		toEquipPos := getEquipPosById(playerSwordSoulEquipment, toId)

		toPlayerSwordSoul.Pos = INBAG
		*equips[toEquipPos] = fromId
		fromPlayerSwordSoul.Pos = EQUIPPED

		// 更新用户装备数据
		db.Update.PlayerSwordSoulEquipment(playerSwordSoulEquipment)
		db.Update.PlayerSwordSoul(fromPlayerSwordSoul)
		db.Update.PlayerSwordSoul(toPlayerSwordSoul)

	case fromPlayerSwordSoul.Pos == EQUIPPED && toPlayerSwordSoul.Pos != EQUIPPED:
		//装备的和背包的剑心交换
		fail.When(toSwordSoul.TypeId == sword_soul_dat.TYPE_EXP, "不能装备经验剑心")

		playerSwordSoulEquipment.TypeBits -= fromBit
		fail.When(playerSwordSoulEquipment.TypeBits&toBit == toBit, "其他装备位已装备该类型剑心")
		playerSwordSoulEquipment.TypeBits += toBit

		equips := getEquips(playerSwordSoulEquipment)

		fromEquipPos := getEquipPosById(playerSwordSoulEquipment, fromId)

		fromPlayerSwordSoul.Pos = INBAG
		*equips[fromEquipPos] = toId
		toPlayerSwordSoul.Pos = EQUIPPED

		// 更新用户装备数据
		db.Update.PlayerSwordSoulEquipment(playerSwordSoulEquipment)
		db.Update.PlayerSwordSoul(fromPlayerSwordSoul)
		db.Update.PlayerSwordSoul(toPlayerSwordSoul)

	case fromPlayerSwordSoul.Pos != EQUIPPED && toPlayerSwordSoul.Pos != EQUIPPED:
		//背包内两个剑心交换
		//废弃
		//fromPlayerSwordSoul.Pos, toPlayerSwordSoul.Pos = toPlayerSwordSoul.Pos, fromPlayerSwordSoul.Pos
		//db.Update.PlayerSwordSoul(fromPlayerSwordSoul)
		//db.Update.PlayerSwordSoul(toPlayerSwordSoul)

	}

}

func getEquips(playerSwordSoulEquipment *mdb.PlayerSwordSoulEquipment) (equips []*int64) {

	equips = []*int64{
		&playerSwordSoulEquipment.Pos1,
		&playerSwordSoulEquipment.Pos2,
		&playerSwordSoulEquipment.Pos3,
		&playerSwordSoulEquipment.Pos4,
		&playerSwordSoulEquipment.Pos5,
		&playerSwordSoulEquipment.Pos6,
		&playerSwordSoulEquipment.Pos7,
		&playerSwordSoulEquipment.Pos8,
		&playerSwordSoulEquipment.Pos9,
	}

	return equips
}
