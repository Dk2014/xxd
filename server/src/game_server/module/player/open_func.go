package player

import (
	"game_server/dat/player_dat"
	"game_server/mdb"
	"game_server/module"
)

//设置已播放开启动画的功能开放权值
func setFuncPlaykey(db *mdb.Database, playLock int16) {
	playerKey := db.Lookup.PlayerFuncKey(db.PlayerId())
	if playLock > playerKey.Key {
		playLock = playerKey.Key
	}
	playerKey.PlayedKey = playLock
	db.Update.PlayerFuncKey(playerKey)
}

func setLevelFuncKey(db *mdb.Database, level int16) {
	functions := player_dat.GetLevelFuncs(level)
	playerFuncKey := db.Lookup.PlayerFuncKey(db.PlayerId())
	needUpdate := false
	for _, function := range functions {
		diffUniqueKey := function.UniqueKey & playerFuncKey.UniqueKey
		if diffUniqueKey == 0 {
			openFunc(db, function.Id)
			playerFuncKey.UniqueKey |= function.UniqueKey
			needUpdate = true
		}
	}
	if needUpdate {
		db.Update.PlayerFuncKey(playerFuncKey)
	}
}

//设置功能开启
func setFuncKey(db *mdb.Database, newFuncKey int16) {
	playerFuncKey := db.Lookup.PlayerFuncKey(db.PlayerId())
	playerFuncKey.Key = newFuncKey
	if playerFuncKey.PlayedKey > newFuncKey {
		playerFuncKey.PlayedKey = newFuncKey
	}

	if newFuncKey <= 0 {
		db.Update.PlayerFuncKey(playerFuncKey)
		return
	}

	lockFuncs := player_dat.GetLockFuncs(newFuncKey)
	for _, function := range lockFuncs {
		if (function.UniqueKey & playerFuncKey.UniqueKey) == 0 {
			playerFuncKey.UniqueKey |= function.UniqueKey
			openFunc(db, function.Id)
		}
	}

	// 发送功能开启通知
	session, isOnline := module.Player.GetPlayerOnline(db.PlayerId())
	if isOnline {
		module.Notify.SendFuncKeyChange(session, newFuncKey)
	}

	db.Update.PlayerFuncKey(playerFuncKey)
}

func openFunc(db *mdb.Database, funcId int32) {
	switch funcId {
	case player_dat.FUNC_MULTI_LEVEL:
		//多人关卡
		module.MultiLevel.OpenFunc(db)
	case player_dat.FUNC_ACTIVE_LEVLE:
		//活动关卡（俠之试炼）
		module.Mission.OpenFuncForExtendLevel(db, player_dat.FUNC_ACTIVE_LEVLE)
	case player_dat.FUNC_MEDITATION:
		//打坐
		module.Meditation.OpenFunc(db)
	case player_dat.FUNC_ARENA:
		module.Arena.OpenFunc(db)
	case player_dat.FUNC_RESOURCE_LEVEL:
		//资源关卡
		module.Mission.OpenFuncForExtendLevel(db, player_dat.FUNC_RESOURCE_LEVEL)
	case player_dat.FUNC_GHOST:
		module.Ghost.OpenFuncForGhost(db)

	case player_dat.FUNC_SWORD_SOUL:
		module.SwordSoul.OpenFuncForSwordSoul(db)

	case player_dat.FUNC_BATTLE_PET:
		module.BattlePet.OpenFunc(db)
		//灵宠幻境
		module.PetVirtualEnv.OpenFuncForPetVirtualEnv(db)
	case player_dat.FUNC_FATE_BOX:
		module.Draw.OpenFuncForFateBox(db)
	case player_dat.FUNC_RAINBOW_LEVEL:
		//极限关卡（彩虹桥）
		module.Rainbow.OpenFunc(db)
	case player_dat.FUNC_TOTEM:
		module.Totem.OpenFuncForTotem(db)
	case player_dat.FUNC_DRIVING_SWORD:
		//云海御剑
		module.DrivingSword.OpenFunc(db)
	}
}
