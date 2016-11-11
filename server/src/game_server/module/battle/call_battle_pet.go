package battle

import (
	"core/fail"
	"core/net"
	"game_server/battle"
	"game_server/dat/battle_pet_dat"
	"game_server/mdb"
	"game_server/module"
)

func callBattlePet(session *net.Session, gridNum int8) bool {
	state := module.State(session)
	fail.When(gridNum > battle_pet_dat.PET_GRID5_NUM || gridNum < battle_pet_dat.PET_GRID1_NUM, "set pet grid error")
	fail.When(state.MissionLevelState == nil || state.Battle == nil, "can not call battle pet")

	var grid *mdb.PlayerBattlePetGrid
	state.Database.Select.PlayerBattlePetGrid(func(row *mdb.PlayerBattlePetGridRow) {
		if row.GridId() == gridNum {
			grid = row.GoObject()
		}
	})

	fail.When(grid == nil || grid.BattlePetId <= 0, "not found player battle pet")

	_, petCalled := state.MissionLevelState.CalledBattlePet[grid.BattlePetId]
	fail.When(petCalled, "battle pet already call")

	foundPet := false
	var battlePet *mdb.PlayerBattlePet
	//查找玩家的灵宠信息
	state.Database.Select.PlayerBattlePet(func(row *mdb.PlayerBattlePetRow) {
		if row.BattlePetId() == grid.BattlePetId {
			foundPet = true
			battlePet = row.GoObject()
			row.Break()
		}
	})
	fail.When(!foundPet, "Didn't find pet in a specified grid")

	if state.Battle.CallBattlePet(session, grid.BattlePetId, int16(battlePet.Level), int16(battlePet.SkillLevel)) {
		if state.Battle.GetBattle().BatType != battle.BT_PET_LEVEL {
			// 灵宠关卡不限制召唤次数
			state.MissionLevelState.CalledBattlePet[grid.BattlePetId] = 1
		}
		return true
	}
	return false
}

func callBattlePetInMultiLevel(session *net.Session, gridNum int8) (success bool) {
	// 屏蔽掉多人关卡灵宠召唤
	return false

	state := module.State(session)
	battlePetId, ok := state.MultiLevelState.BattlePetInfo[gridNum]
	fail.When(!ok, "找不到可召唤的灵宠信息")
	//TODO CallBattlePet 的最后一个参数 需要在构建多人关卡的时候通过RPC穿过来
	//放在  state.MultiLevelState 里面。 暂时给非法值 -1 在多人关卡开启灵宠召唤时修复这个非法参数
	//原来最后一个参数是grid level，现在被我改成pet level -- DHZ 20141228
	if success = state.Battle.CallBattlePet(session, battlePetId, -1, -1); success {
		//记录灵宠已被使用
		delete(state.MultiLevelState.BattlePetInfo, gridNum)
	}
	return success
}
