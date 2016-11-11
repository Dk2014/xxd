package player

import (
	"core/fail"
	"core/net"
	"core/time"
	"game_server/dat/heart_dat"
	"game_server/dat/heart_draw_dat"
	"game_server/dat/physical_dat"
	"game_server/dat/push_notify_dat"
	"game_server/dat/quest_dat"
	"game_server/dat/role_dat"
	"game_server/dat/skill_dat"
	"game_server/dat/town_dat"
	"game_server/mdb"
	"game_server/module"
)

func InitPlayer(session *net.Session, nickname string, roleId int8) (ret bool, pid int64) {
	if nickname == "" {
		return
	}

	if _, ret = GetPlayerIdByNickname(nickname); ret {
		ret = false
		return
	}

	state := module.State(session)
	player := state.Database.Lookup.Player(state.PlayerId)
	if player == nil {
		return
	}

	main_role_id := initMainRole(state.Database, roleId)

	player.Nick = nickname
	player.MainRoleId = main_role_id
	state.Database.Update.Player(player)

	// 更多初始化 coding here!
	initPlayerTown(state.Database)
	initPlayerQuest(state.Database)
	initPlayerPhysical(state.Database)
	initPlayerFormation(state.Database, roleId)
	initPlayerTeamInfo(state.Database)
	initPlayerMission(state.Database)
	initPlayerMissionLevel(state.Database)
	initPlayerInfo(state.Database)
	initPlayerFunKey(state.Database)
	initPlayerSkill(state.Database, roleId)
	initPlayerEquipment(state.Database, roleId)
	initPlayerItemBuyback(state.Database)
	initPlayerHeart(state.Database)
	initPlayerFightNum(state.Database)
	initPlayerHeartDraw(state.Database)
	initPlayerCoinsInfo(state.Database)
	initPlayerVIP(state.Database)
	initPlayerHardLevel(state.Database)
	initPlayerDailySignInState(state.Database)
	initPlayerOperated(state.Database)
	initPlayerState(state.Database)
	initPlayerLoginAwardRecord(state.Database)
	initPlayerFashion(state.Database)
	initPlayerEventFightPower(state.Database)
	initPlayerEventAwardRecord(state.Database)
	initPlayerPushNotification(state.Database)
	initPlayerActivity(state.Database)
	initPlayerFame(state.Database)
	initPlayerSealeBook(state.Database)
	ret = true
	pid = state.PlayerId

	state.RoleId = roleId
	return
}

func initPlayerEventFightPower(db *mdb.Database) {
	db.Insert.PlayerEventsFightPower(&mdb.PlayerEventsFightPower{
		Pid:  db.PlayerId(),
		Lock: 0,
	})
}

func initPlayerOperated(db *mdb.Database) {
	db.Insert.PlayerIsOperated(&mdb.PlayerIsOperated{
		Pid:          db.PlayerId(),
		OperateValue: 0,
	})
}

func initPlayerHeartDraw(db *mdb.Database) {
	db.Insert.PlayerHeartDraw(&mdb.PlayerHeartDraw{
		Pid:      db.PlayerId(),
		DrawType: heart_draw_dat.DRAW_TYPE_WHEEL,
		DailyNum: 0,
		DrawTime: 0,
	})

	db.Insert.PlayerHeartDraw(&mdb.PlayerHeartDraw{
		Pid:      db.PlayerId(),
		DrawType: heart_draw_dat.DRAW_TYPE_CARD,
		DailyNum: 0,
		DrawTime: 0,
	})
}

func initPlayerQuest(db *mdb.Database) {
	db.Insert.PlayerQuest(&mdb.PlayerQuest{
		Pid:     db.PlayerId(),
		QuestId: quest_dat.GetInitQuest(),
		State:   quest_dat.QUEST_STATUS_ALREADY_RECEIVE,
	})
}

func initPlayerPhysical(db *mdb.Database) {
	db.Insert.PlayerPhysical(&mdb.PlayerPhysical{
		Pid:           db.PlayerId(),
		Value:         physical_dat.MAX_PHYSICAL_VALUE_BY_TIME,
		UpdateTime:    time.GetNowTime(),
		BuyCount:      0,
		BuyUpdateTime: 0,
		DailyCount:    0,
	})
}

func initPlayerCoinsInfo(db *mdb.Database) {
	db.Insert.PlayerCoins(&mdb.PlayerCoins{
		Pid: db.PlayerId(),
	})
}

func initMainRole(db *mdb.Database, roleId int8) int64 {
	// 校验主角ID的合法性
	fail.When(role_dat.IsMainRole(roleId) == false, "not main role id")

	//Role增加新字段务必在 module/role/role_init 和 和此处给定初始值
	role := &mdb.PlayerRole{
		Pid:             db.PlayerId(),
		RoleId:          roleId,
		Level:           1,
		FriendshipLevel: 1,
		Exp:             0,
		CoopId:          0,
		Status:          0, //正常
	}

	db.Insert.PlayerRole(role)

	return role.Id
}

func initPlayerTown(db *mdb.Database) {
	townId, lock := town_dat.GetInitTownIdAndLock()
	town := town_dat.GetTownWithTownId(townId)

	db.Insert.PlayerTown(&mdb.PlayerTown{
		Pid:    db.PlayerId(),
		TownId: townId,
		Lock:   lock,
		AtPosX: town.StartX,
		AtPosY: town.StartY,
	})
}

// 玩家默认布阵, 战术类型
func initPlayerFormation(db *mdb.Database, roleId int8) {
	db.Insert.PlayerFormation(&mdb.PlayerFormation{
		Pid:  db.PlayerId(),
		Pos0: -1,
		Pos1: roleId,
		Pos2: -1,
		Pos3: -1,
		Pos4: -1,
		Pos5: -1,
		Pos6: -1,
		Pos7: -1,
		Pos8: -1,
	})
}

func initPlayerTeamInfo(db *mdb.Database) {
	db.Insert.PlayerTeamInfo(&mdb.PlayerTeamInfo{
		Pid:          db.PlayerId(),
		Relationship: 0,
		HealthLv:     0,
		AttackLv:     0,
		DefenceLv:    0,
	})
}

func initPlayerMission(db *mdb.Database) {
	db.Insert.PlayerMission(&mdb.PlayerMission{
		Pid:      db.PlayerId(),
		Key:      0,
		MaxOrder: 0,
	})
}

func initPlayerMissionLevel(db *mdb.Database) {
	db.Insert.PlayerMissionLevel(&mdb.PlayerMissionLevel{
		Pid:     db.PlayerId(),
		Lock:    0,
		MaxLock: 0,
	})
}

func initPlayerInfo(db *mdb.Database) {
	db.Insert.PlayerInfo(&mdb.PlayerInfo{
		Pid:            db.PlayerId(),
		Ingot:          0,
		Coins:          0,
		NewMailNum:     0,
		FirstLoginTime: time.GetNowTime(),
	})
}

func initPlayerFunKey(db *mdb.Database) {
	db.Insert.PlayerFuncKey(&mdb.PlayerFuncKey{
		Pid:       db.PlayerId(),
		Key:       0,
		PlayedKey: 0,
		UniqueKey: 0,
	})
}

func initPlayerSkill(db *mdb.Database, roleId int8) {
	var skillId1 int16 = 0
	var skillId2 int16 = 0

	roleInfo := role_dat.GetRoleInfo(roleId)

	if roleInfo.SkillID1 != 0 {
		skillInfo1 := skill_dat.GetSkillInfo(roleInfo.SkillID1)
		db.Insert.PlayerSkill(&mdb.PlayerSkill{
			Pid:        db.PlayerId(),
			RoleId:     roleId,
			SkillId:    skillInfo1.Id,
			SkillTrnlv: 1,
		})
		skillId1 = skillInfo1.Id
	}

	if roleInfo.SkillID2 != 0 {
		skillInfo2 := skill_dat.GetSkillInfo(roleInfo.SkillID2)
		db.Insert.PlayerSkill(&mdb.PlayerSkill{
			Pid:        db.PlayerId(),
			RoleId:     roleId,
			SkillId:    skillInfo2.Id,
			SkillTrnlv: 1,
		})
		skillId2 = skillInfo2.Id
	}

	db.Insert.PlayerUseSkill(&mdb.PlayerUseSkill{
		Pid:      db.PlayerId(),
		RoleId:   roleId,
		SkillId1: skillId1,
		SkillId2: skillId2,
		SkillId3: 0,
		SkillId4: 0,
	})
}

func initPlayerEquipment(db *mdb.Database, roleId int8) {
	db.Insert.PlayerEquipment(&mdb.PlayerEquipment{
		Pid:           db.PlayerId(),
		RoleId:        roleId,
		WeaponId:      0,
		ClothesId:     0,
		AccessoriesId: 0,
		ShoeId:        0,
	})
}

func initPlayerItemBuyback(db *mdb.Database) {
	db.Insert.PlayerItemBuyback(&mdb.PlayerItemBuyback{
		Pid: db.PlayerId(),
	})
}

func initPlayerHeart(db *mdb.Database) {
	db.Insert.PlayerHeart(&mdb.PlayerHeart{
		Pid:        db.PlayerId(),
		Value:      heart_dat.HEART_INIT_VAL,
		UpdateTime: time.GetNowTime(),
	})
}

func initPlayerFightNum(db *mdb.Database) {
	db.Insert.PlayerFightNum(&mdb.PlayerFightNum{
		Pid: db.PlayerId(),
	})
}

func initPlayerVIP(db *mdb.Database) {
	db.Insert.PlayerVip(&mdb.PlayerVip{
		Pid: db.PlayerId(),
	})
}

func initPlayerHardLevel(db *mdb.Database) {
	db.Insert.PlayerHardLevel(&mdb.PlayerHardLevel{
		Pid: db.PlayerId(),
	})
}

func initPlayerDailySignInState(db *mdb.Database) {
	db.Insert.PlayerDailySignInState(&mdb.PlayerDailySignInState{
		Pid: db.PlayerId(),
	})
}

func initPlayerState(db *mdb.Database) {
	db.Insert.PlayerState(&mdb.PlayerState{
		Pid:          db.PlayerId(),
		BanStartTime: 0,
		BanEndTime:   -1,
	})
}

func initPlayerLoginAwardRecord(db *mdb.Database) {
	db.Insert.PlayerLoginAwardRecord(&mdb.PlayerLoginAwardRecord{
		Pid: db.PlayerId(),
	})
}

func initPlayerFashion(db *mdb.Database) {
	db.Insert.PlayerFashionState(&mdb.PlayerFashionState{
		Pid: db.PlayerId(),
	})
}

func initPlayerEventAwardRecord(db *mdb.Database) {
	db.Insert.PlayerEventAwardRecord(&mdb.PlayerEventAwardRecord{
		Pid: db.PlayerId(),
	})
}

func initPlayerSealeBook(db *mdb.Database) {
	db.Insert.PlayerSealedbook(&mdb.PlayerSealedbook{
		Pid: db.PlayerId(),
	})
}

func initPlayerPushNotification(db *mdb.Database) {
	var options int64
	for _, id := range push_notify_dat.DefaultNotificationID {
		options += int64(1 << uint64(id))
	}
	db.Insert.PlayerPushNotifySwitch(&mdb.PlayerPushNotifySwitch{
		Pid:     db.PlayerId(),
		Options: options,
	})
}

func initPlayerActivity(db *mdb.Database) {
	db.Insert.PlayerActivity(&mdb.PlayerActivity{
		Pid: db.PlayerId(),
	})
}

func initPlayerFame(db *mdb.Database) {
	db.Insert.PlayerFame(&mdb.PlayerFame{
		Pid:   db.PlayerId(),
		Level: 1,
	})
}
