package mission

import (
	"core/fail"
	"core/net"
	"game_server/api/protocol/battle_api"
	"game_server/api/protocol/mission_api"
	"game_server/battle"
	"game_server/dat/chest_dat"
	"game_server/dat/mission_dat"
	"game_server/dat/player_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/xdlog"
)

func init() {
	module.Mission = MissionMod{}
	module.DisposeEvent.Regisiter(DisposeHandler)
}

type MissionMod struct {
}

func (mod MissionMod) HardLevelLose(state *module.SessionState, defendSide *battle.SideInfo) {
	////获取战场信息
	//battle := state.Battle.GetBattle()

	////查找出难度关卡的状态
	//var hardLevelState *mdb.PlayerHardLevelState
	//state.Database.Select.PlayerHardLevelState(func(row *mdb.PlayerHardLevelStateRow) {
	//  if row.MissionLevelId() == state.MissionLevelState.LevelId {
	//      hardLevelState = row.GoObject()
	//      row.Break()
	//  }
	//})
	//needInsert := false
	//if hardLevelState == nil {
	//  hardLevelState = &mdb.PlayerHardLevelState{
	//      Pid:            state.PlayerId,
	//      MissionLevelId: state.MissionLevelState.LevelId,
	//      BattleTime:     0,
	//  }
	//  needInsert = true
	//}

	//if !coreTime.IsToday(hardLevelState.BattleTime) {
	//  hardLevelState.BattleTime = time.Now().Unix()
	//  hardLevelState.Round = int32(battle.GetRounds())
	//} else {
	//  hardLevelState.Round += int32(battle.GetRounds())
	//}
	//hardLevelState.Round++
	//module.Battle.DumpHardLevelEnemyHealth(defendSide, hardLevelState)
	//if needInsert {
	//  state.Database.Insert.PlayerHardLevelState(hardLevelState)
	//} else {
	//  state.Database.Update.PlayerHardLevelState(hardLevelState)
	//}
}

func (mod MissionMod) StartLevelBattle(battleType int8, enemy_id int32, session *net.Session) {
	state := module.State(session)

	fail.When(state.MissionLevelState == nil, "MissionLevelState is nil")

	levelEnemy := mission_dat.GetMissionLevelEnemyById(int32(enemy_id))
	fail.When(levelEnemy.MissionLevelId != state.MissionLevelState.LevelId, "incorrect enemy_id")
	fail.When(state.MissionLevelState.IsPassEnemy(enemy_id), "enemy is pass")
	fail.When(state.MissionLevelState.PassBoss, "level boss is pass")

	// 要开打了，先存下这次战斗的地图，用于客户端断线重连回来
	state.MissionLevelState.LastFightingMap = levelEnemy.ShadedMissionId

	state.Battle = module.Battle.NewBattleForLevel(session, battleType, enemy_id, false)
}

func (mod MissionMod) GetMissionLevelRecord(db *mdb.Database, levelId int32) (levelRecord *mdb.PlayerMissionLevelRecord) {
	db.Select.PlayerMissionLevelRecord(func(row *mdb.PlayerMissionLevelRecordRow) {
		if row.MissionLevelId() == levelId {
			levelRecord = row.GoObject()
			row.Break()
		}
	})

	return
}

func (mod MissionMod) GetHardLevelRecordById(db *mdb.Database, levelId int32) (levelRecord *mdb.PlayerHardLevelRecord) {
	db.Select.PlayerHardLevelRecord(func(row *mdb.PlayerHardLevelRecordRow) {
		if row.LevelId() == levelId {
			levelRecord = row.GoObject()
			row.Break()
		}
	})
	return
}

func (mod MissionMod) NotifyCatchBattlePet(session *net.Session, petId int32) {
	session.Send(&mission_api.NotifyCatchBattlePet_Out{
		PetId: petId,
	})
}

func (mod MissionMod) OpenFuncForExtendLevel(db *mdb.Database, funcId int32) {
	newRecord := false

	playerExtendLevel := db.Lookup.PlayerExtendLevel(db.PlayerId())
	if playerExtendLevel == nil {
		newRecord = true
		playerExtendLevel = &mdb.PlayerExtendLevel{
			Pid:           db.PlayerId(),
			CoinPassTime:  -1,
			ExpPassTime:   -1,
			GhostPassTime: -1,
			PetPassTime:   -1,
			BuddyPassTime: -1,
			CoinDailyNum:  -1,
			ExpDailyNum:   -1,
			BuddyDailyNum: -1,
			PetDailyNum:   -1,
			GhostDailyNum: -1,
		}
	}

	switch funcId {
	case player_dat.FUNC_RESOURCE_LEVEL:
		playerExtendLevel.CoinDailyNum = 0
		playerExtendLevel.CoinPassTime = 0

		playerExtendLevel.ExpPassTime = 0
		playerExtendLevel.ExpDailyNum = 0

	case player_dat.FUNC_ACTIVE_LEVLE:
		playerExtendLevel.BuddyDailyNum = 0
		playerExtendLevel.BuddyPassTime = 0
		playerExtendLevel.RolePos = 0
		playerExtendLevel.BuddyPos = 0
		playerExtendLevel.RandBuddyRoleId = randRoleIdForBuddyLevel(db)

		playerExtendLevel.PetDailyNum = 0
		playerExtendLevel.PetPassTime = 0

		playerExtendLevel.GhostDailyNum = 0
		playerExtendLevel.GhostPassTime = 0

	default:
		fail.When(true, "error extend level type")
	}

	if newRecord {
		db.Insert.PlayerExtendLevel(playerExtendLevel)
	} else {
		db.Update.PlayerExtendLevel(playerExtendLevel)
	}
}

func (mod MissionMod) AwardCatchedBattlePet(state *module.SessionState, xdEventType int32) {
	fail.When(state.MissionLevelState.CatchedBattlePetId == 0, "not catch battle pet")
	awardCatchedBattlePet(state, xdEventType)
}

// 正常的关卡奖励
func (mod MissionMod) AwardLevel(state *module.SessionState, xdEventType int32) {
	awardLevel(state, xdEventType)
}

// 服务端直接奖励
func (mod MissionMod) ServerDirectAward(state *module.SessionState,
	playerLevel *mdb.PlayerMissionLevel, levelId int32, boxOpenCount int8, xdEventType int32) {

	boxState := &module.MissionLevelBoxState{}
	levelInfo := mission_dat.GetMissionLevelById(levelId)

	// 先奖励宝箱
	if levelInfo.AwardBox {
		boxState.AwardedList = make(map[int8]int64)
		boxState.OpenedBoxPos = make([]int8, 5)
		boxState.AwardOpenCount = boxOpenCount

		for i := int8(0); i < boxOpenCount; i++ {
			realOpenBox(state, levelInfo.Id, i+1, boxState, nil, xdEventType)
		}
	}

	realAwardLevel(state, playerLevel, levelInfo, boxState.AwardExp,
		boxState.AwardCoin, boxState.AwardMultExp, boxState.AwardMultCoin, xdEventType)
}

// 离开关卡
func (mod MissionMod) LeaveMissionLevel(state *module.SessionState) {
	if state.MissionLevelState != nil {
		if state.Battle != nil {
			state.Battle.LeaveBattle(nil)
			state.Battle = nil
		}
		state.MissionLevelState = nil
	}
}

func (mod MissionMod) SetMissionLevelLock(state *module.SessionState, newLock int32, playerLevel *mdb.PlayerMissionLevel, isNotify bool) {
	if playerLevel == nil {
		playerLevel = state.Database.Lookup.PlayerMissionLevel(state.PlayerId)
	}

	fail.When(playerLevel.MaxLock >= newLock, "关卡权值已获得")
	//Lock是一个冗余的字段，目前仅用到 MaxLock
	playerLevel.Lock = newLock
	playerLevel.MaxLock = newLock

	state.Database.Update.PlayerMissionLevel(playerLevel)

	if isNotify {
		if session, ok := module.Player.GetPlayerOnline(state.PlayerId); ok {
			module.Notify.SendMissionLevelLockChanged(session, playerLevel.MaxLock, playerLevel.AwardLock)
		}
	}
}

// 检查玩家是否需要重建关卡
func (mod MissionMod) CheckIsNeedRebuildMissionLevel(session *net.Session) {
	state := module.State(session)
	if !module.RevertMissionLevelStateForPlayer(state) {
		return
	}

	passEnemys := make([]mission_api.Rebuild_Out_Pass, len(state.MissionLevelState.PassEnemyIds))
	for i, id := range state.MissionLevelState.PassEnemyIds {
		passEnemys[i].EnemyId = id
	}

	rsp := &mission_api.Rebuild_Out{
		LevelType:       mission_api.BattleType(state.MissionLevelState.LevelType),
		LevelId:         state.MissionLevelState.LevelId,
		ReliveIngot:     state.MissionLevelState.NextReliveCostIngot,
		Pass:            passEnemys,
		TotalRound:      int16(state.MissionLevelState.TotalRound),
		LastFightingMap: state.MissionLevelState.LastFightingMap,
	}

	if state.MissionLevelState.LevelType == battle.BT_BUDDY_LEVEL {
		playerExtendLevel := state.Database.Lookup.PlayerExtendLevel(state.PlayerId)
		rsp.BuddyRoleId = playerExtendLevel.RandBuddyRoleId
		rsp.MainRolePos = playerExtendLevel.RolePos
		rsp.BuddyPos = playerExtendLevel.BuddyPos
	}

	if len(state.MissionLevelState.SmallBoxList) > 0 {
		for boxId, _ := range state.MissionLevelState.SmallBoxList {
			rsp.Smallbox = append(rsp.Smallbox, mission_api.Rebuild_Out_Smallbox{
				BoxId: boxId,
			})
		}
	}
	if len(state.MissionLevelState.MengYaoList) > 0 {
		for myId, _ := range state.MissionLevelState.MengYaoList {
			rsp.MengYao = append(rsp.MengYao, mission_api.Rebuild_Out_MengYao{
				MyId: myId,
			})
		}
	}
	if len(state.MissionLevelState.ShadowList) > 0 {
		for shadedId, _ := range state.MissionLevelState.ShadowList {
			rsp.Shadow = append(rsp.Shadow, mission_api.Rebuild_Out_Shadow{
				ShadedId: shadedId,
			})
		}
	}

	session.Send(rsp)

	module.Notify.RoleBattleStatusChange(session)
}

func (mod MissionMod) CountMissoinStar(db *mdb.Database, missionId int16) (star int16) {
	db.Select.PlayerMissionLevelRecord(func(row *mdb.PlayerMissionLevelRecordRow) {
		levelDat := mission_dat.GetMissionLevelById(row.MissionLevelId())
		if levelDat.MissionId == missionId {
			levelStar := mission_dat.CalLevelStarByRound(row.MissionLevelId(), row.Round())
			star += int16(levelStar)
		}

	})
	return
}

func (mod MissionMod) StartFateBoxLevelBattle(session *net.Session, enemy_id int32) {
	state := module.State(session)
	levelEnemy := mission_dat.GetMissionLevelEnemyById(enemy_id)
	missionLevel := mission_dat.GetMissionLevelById(levelEnemy.MissionLevelId)
	fail.When(missionLevel.ParentType != int8(battle_api.BATTLE_TYPE_FATE_BOX), "关卡敌人与关卡类型不匹配")
	fateBoxDat := chest_dat.GetFateBoxByMissionLevelId(missionLevel.Id)
	mainRole := module.Role.GetMainRole(state.Database)
	fail.When(mainRole.Level < fateBoxDat.Level, "挑战命锁关卡失败：等级不够")
	fateBoxState := state.Database.Lookup.PlayerFateBoxState(state.PlayerId)
	fail.When(fateBoxState.Lock < fateBoxDat.RequireLock, "挑战命锁关卡失败：命锁权值不够")

	module.Town.LeaveTown(session)
	state.MissionLevelState = module.NewMissionLevelState(battle.BT_FATE_BOX_LEVEL, missionLevel.Id)
	state.MissionLevelState.LevelId = missionLevel.Id
	state.MissionLevelState.LevelType = battle.BT_FATE_BOX_LEVEL
	state.MissionLevelState.EnemyId = enemy_id
	state.MissionLevelState.LoadBuddySkill(state)
	state.Battle = module.Battle.NewBattleForLevel(session, battle.BT_FATE_BOX_LEVEL, enemy_id, false)
}

func autoOpenlevelBox(state *module.SessionState, xdEventType int32) {
	for i := state.MissionLevelState.BoxState.AwardOpenCount; i > 0; i-- {
		// 找未开启的宝箱开启
		for pos, v := range state.MissionLevelState.BoxState.OpenedBoxPos {
			if v == 0 {
				OpenLevelBox(state, int8(pos)+1, xdEventType)
				break
			}
		}
	}
}

// 会话销毁时执行
func DisposeHandler(state *module.SessionState) {
	// 玩家存在关卡状态
	if state.MissionLevelState != nil {
		/*
		   系统自动为玩家抽取区域关卡通关宝箱规则：
		   *）当玩家已经抽过一次或以上的宝箱时
		   *）如果奖励了玩家宝箱抽取次数后，从未抽取过宝箱，那么恢复关卡状态，让玩家自行抽取
		*/

		levelType := state.MissionLevelState.LevelType
		var boxCount int8
		if levelType == battle.BT_MISSION_LEVEL || levelType == battle.BT_HARD_LEVEL {
			//boxCount = mission_dat.MAX_LEVEL_BOX_AWARD_COUNT
			boxCount = mission_dat.CalLevelStarByRound(state.MissionLevelState.LevelId, int8(state.MissionLevelState.TotalRound))
		} else {
			boxCount = mission_dat.LEVEL_BOX_AWARD_COUNT
		}

		if state.MissionLevelState.BoxState.AwardOpenCount > 0 &&
			state.MissionLevelState.BoxState.AwardOpenCount < boxCount {
			autoOpenlevelBox(state, xdlog.ET_DISPOSE_EVENT)
		} else {
			switch state.MissionLevelState.LevelType {
			case battle.BT_MISSION_LEVEL, battle.BT_BUDDY_LEVEL, battle.BT_GHOST_LEVEL, battle.BT_PET_LEVEL, battle.BT_RESOURCE_LEVEL:
				if mission_dat.CheckMissionLevel(state.MissionLevelState.LevelId) {
					module.SaveMissionLevelStateForPlayer(state)
				}
			default:
				if state.MissionLevelState.BoxState.AwardOpenCount > 0 {
					autoOpenlevelBox(state, xdlog.ET_DISPOSE_EVENT)
				}
			}
		}
	}
}
