package battle_pet

import (
	"core/fail"
	//"core/net"
	"core/time"
	"game_server/api/protocol/battle_pet_api"
	"game_server/battle"
	"game_server/dat/battle_pet_dat"
	"game_server/dat/item_dat"
	"game_server/dat/player_dat"
	"game_server/dat/quest_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/tlog"
	"game_server/xdlog"
	"math/rand"
)

func getPetInfo(state *module.SessionState, rsp *battle_pet_api.GetPetInfo_Out) {
	rsp.Pet = []battle_pet_api.GetPetInfo_Out_Pet{}
	rsp.Set = []battle_pet_api.GetPetInfo_Out_Set{}

	fail.When(!module.Player.IsOpenFunc(state.Database, player_dat.FUNC_BATTLE_PET), "灵宠功能未开启")

	//playerBattlePetState := getPlayerBattlePetState(state.Database)
	//rsp.UpgradeTimes = playerBattlePetState.UpgradeByIngotNum

	var calledPet bool

	state.Database.Select.PlayerBattlePet(func(row *mdb.PlayerBattlePetRow) {
		if state.MissionLevelState != nil && state.MissionLevelState.CalledBattlePet != nil {
			_, calledPet = state.MissionLevelState.CalledBattlePet[row.BattlePetId()]
		} else {
			calledPet = false
		}

		//pet := battle_pet_dat.GetBattlePetWithEmemyId(int32(row.BattlePetId()))
		rsp.Pet = append(rsp.Pet, battle_pet_api.GetPetInfo_Out_Pet{
			PetId:      row.BattlePetId(),
			Level:      int16(row.Level()),
			Exp:        row.Exp(),
			SkillLevel: int16(row.SkillLevel()),
			Called:     calledPet,
		})
	})

	state.Database.Select.PlayerBattlePetGrid(func(row *mdb.PlayerBattlePetGridRow) {
		rsp.Set = append(rsp.Set, battle_pet_api.GetPetInfo_Out_Set{
			GridNum: row.GridId(),
			PetId:   row.BattlePetId(),
			//Exp:     row.Exp(),
			//Level:   row.Level(),
		})
	})
}

func addPet(db *mdb.Database, petId, itemFlowReason, xdEventType int32) {
	var playerPet *mdb.PlayerBattlePet
	petDat := battle_pet_dat.GetBattlePetWithEnemyId(petId)
	db.Select.PlayerBattlePet(func(row *mdb.PlayerBattlePetRow) {
		if row.BattlePetId() == petDat.PetId {
			playerPet = row.GoObject()
			row.Break()
		}
	})

	if playerPet == nil { //新获一个灵宠
		playerPet = &mdb.PlayerBattlePet{
			Pid:         db.PlayerId(),
			BattlePetId: petDat.PetId,
			//Star:        petDat.Star,
			//ParentPetId: petDat.ParentPetId,
			Level:      1,
			SkillLevel: 1,
		}
		db.Insert.PlayerBattlePet(playerPet)

		//for sealedbook
		if session, exist := module.Player.GetPlayerOnline(db.PlayerId()); exist {
			state := module.State(session)
			if _, result := state.GetSealedBookRecord().FindRecord(item_dat.STEALDBOOK_TYPE_PETS, int16(petId)); !result {
				sealedbook := item_dat.GetSealedBookInfo(item_dat.STEALDBOOK_TYPE_PETS, int16(petId))
				if sealedbook != nil {
					state.GetSealedBookRecord().AddRecord(item_dat.STEALDBOOK_TYPE_PETS, int16(petId), item_dat.STEALDBOOK_HAVING, db)
				}
			}
		}

	} else {
		petSoulNum := int16(battle_pet_dat.DUPLICATED_PET2SOUL_NUM)
		module.Item.AddItem(db, battle_pet_dat.ITEM_BATTLE_PET_SOUL, petSoulNum, itemFlowReason, xdEventType, "")
	}
	tlog.PlayerAddPetFlowLog(db, petId)
}

func failIfPetGridNotOpened(mainLevel int16, gridNum int8) {
	failStr := "main role level is less than the level grid number required"
	switch gridNum {
	case 1:
	case 2:
		fail.When(battle_pet_dat.OPEN_GRID_LEVEL1 > mainLevel, failStr)
	case 3:
		fail.When(battle_pet_dat.OPEN_GRID_LEVEL2 > mainLevel, failStr)
	case 4:
		fail.When(battle_pet_dat.OPEN_GRID_LEVEL3 > mainLevel, failStr)
	case 5:
		fail.When(battle_pet_dat.OPEN_GRID_LEVEL4 > mainLevel, failStr)
	default:
	}
}

func setPet(state *module.SessionState, gridNum int8, petId int32) {
	fail.When(gridNum > battle_pet_dat.PET_GRID5_NUM || gridNum < battle_pet_dat.PET_GRID1_NUM, "set pet grid error")
	failIfPetGridNotOpened(module.Role.GetMainRole(state.Database).Level, gridNum)

	// 灵宠关卡、彩虹关卡可以操作灵宠
	if state.MissionLevelState != nil {
		switch state.MissionLevelState.LevelType {
		case battle.BT_PET_LEVEL:
		case battle.BT_RAINBOW_LEVEL:
		default:
			fail.When(state.MissionLevelState != nil, "can't setPet in level")
		}
	}

	var targetGrid *mdb.PlayerBattlePetGrid
	var petNum int
	state.Database.Select.PlayerBattlePetGrid(func(row *mdb.PlayerBattlePetGridRow) {
		if row.BattlePetId() > 0 {
			petNum++
		}
		if gridNum == row.GridId() {
			targetGrid = row.GoObject()
		}
		if petId > 0 && petId == row.BattlePetId() { //操纵的灵宠已装备状态
			fail.When(true, "目标灵宠已装备")
		}
	})

	fail.When(targetGrid == nil, "错误的灵宠格子ID")

	if petId > 0 { //装载灵宠
		var foundTargetPet bool
		state.Database.Select.PlayerBattlePet(func(row *mdb.PlayerBattlePetRow) {
			if row.BattlePetId() == petId {
				foundTargetPet = true
				row.Break()
			}
		})
		fail.When(!foundTargetPet, "找不到要装备的灵宠")
	} else if petId == 0 { //卸载灵宠
		// 灵宠关卡限制检查；至少要装配两个灵宠
		if state.MissionLevelState != nil && state.MissionLevelState.LevelType == battle.BT_PET_LEVEL {
			fail.When(petNum <= 2, "pet-level MUST set two battle-pet")
		}
	}

	targetGrid.BattlePetId = petId
	state.Database.Update.PlayerBattlePetGrid(targetGrid)

	tlog.PlayerSystemModuleFlowLog(state.Database, tlog.SMT_PET)
}

func setPetSwap(state *module.SessionState, fromGridNum int8, toGridNum int8) {
	failIfPetGridNotOpened(module.Role.GetMainRole(state.Database).Level, fromGridNum)
	failIfPetGridNotOpened(module.Role.GetMainRole(state.Database).Level, toGridNum)
	// 灵宠关卡、彩虹关卡可以装备灵宠
	if state.MissionLevelState != nil {
		switch state.MissionLevelState.LevelType {
		case battle.BT_PET_LEVEL:
		case battle.BT_RAINBOW_LEVEL:
		default:
			fail.When(state.MissionLevelState != nil, "can't setPet in level")
		}
	}
	var srcGrid, dstGrid *mdb.PlayerBattlePetGrid
	state.Database.Select.PlayerBattlePetGrid(func(row *mdb.PlayerBattlePetGridRow) {
		if row.GridId() == fromGridNum {
			srcGrid = row.GoObject()
		} else if row.GridId() == toGridNum {
			dstGrid = row.GoObject()
		}
	})

	dstGrid.BattlePetId, srcGrid.BattlePetId = srcGrid.BattlePetId, dstGrid.BattlePetId

	state.Database.Update.PlayerBattlePetGrid(srcGrid)
	state.Database.Update.PlayerBattlePetGrid(dstGrid)
}

func getPlayerBattlePetState(db *mdb.Database) *mdb.PlayerBattlePetState {
	playerBattlePetState := db.Lookup.PlayerBattlePetState(db.PlayerId())
	if !time.IsInPointHour(player_dat.RESET_BATTLE_PET_UPGRADE_INFO_IN_HOUR, playerBattlePetState.UpgradeByIngotTime) {
		playerBattlePetState.UpgradeByIngotNum = 0
		db.Update.PlayerBattlePetState(playerBattlePetState)
	}
	return playerBattlePetState
}

func findSweetyPetByPetId(db *mdb.Database, petId int32) (sweety *mdb.PlayerBattlePet, found bool) {
	found = false

	// 查找灵宠信息
	db.Select.PlayerBattlePet(func(row *mdb.PlayerBattlePetRow) {
		if row.BattlePetId() == petId {
			found = true
			sweety = row.GoObject()
			row.Break()
		}
	})

	return
}

func upgradePet(state *module.SessionState, petId int32) ( /*exp*/ int64 /*level*/, int16) {
	db := state.Database

	sweetyPet, foundPet := findSweetyPetByPetId(db, petId)

	// 检查是否存在
	fail.When(!foundPet, "Didn't find the training pet specified by client side")
	// 灵宠等级不能超过最大等级
	fail.When(sweetyPet.Level >= battle_pet_dat.MAX_BATTLE_PET_LEVEL, "the battle pet level is already the maximum level")
	// 灵宠等级不能超越主角等级
	mainRole := module.Role.GetMainRole(db)
	fail.When(int16(sweetyPet.Level) >= mainRole.Level, "the pets can not be levelup over main role level")

	petLevelExp := battle_pet_dat.GetPetLevelExpInfo(int16(sweetyPet.Level))

	// 扣除灵魄
	module.Item.DelItemByItemId(db, battle_pet_dat.ITEM_BATTLE_PET_SOUL, int16(petLevelExp.NeedSoulNum), tlog.IFR_BATTLE_PET_UPGRADE, xdlog.ET_BATTLE_PET_UPGRADE)

	// 计算经验增量
	add_exp := petLevelExp.MinAddExp + rand.Int63n(petLevelExp.MaxAddExp-petLevelExp.MinAddExp+1)

	// 升级
	sweetyPet.Exp += add_exp
	for sweetyPet.Exp >= petLevelExp.Exp && int16(sweetyPet.Level) < mainRole.Level {
		sweetyPet.Exp -= petLevelExp.Exp
		sweetyPet.Level++
		petLevelExp = battle_pet_dat.GetPetLevelExpInfo(int16(sweetyPet.Level))
	}
	db.Update.PlayerBattlePet(sweetyPet)

	//刷新每日任务
	module.Quest.RefreshDailyQuest(state, quest_dat.DAILY_QUEST_CLASS_LEVELUP_PET)

	return sweetyPet.Exp, int16(sweetyPet.Level)
}

func trainingPetSkill(state *module.SessionState, petId int32) {
	db := state.Database
	sweetyPet, foundPet := findSweetyPetByPetId(db, petId)

	// 检查是否存在
	fail.When(!foundPet, "Didn't find the training pet specified by client side")
	// 技能不能超过灵宠等级
	fail.When(sweetyPet.SkillLevel >= sweetyPet.Level, "the skill level of sweety pets can not be bigger than level of them")

	skillStuff := battle_pet_dat.GetPetSkillStuff(int16(sweetyPet.SkillLevel))
	// 扣钱
	module.Player.DecMoney(db, state.MoneyState, skillStuff.CostCoins, player_dat.COINS, tlog.MFR_PET_SKILL_LEVELUP, xdlog.ET_PET_SKILL)

	// 升级
	sweetyPet.SkillLevel++
	db.Update.PlayerBattlePet(sweetyPet)
}
