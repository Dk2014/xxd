package battle

import (
	"core/fail"
	"core/net"
	"game_server/api/protocol/battle_api"
	"game_server/battle"
	"game_server/dat/player_dat"
	"game_server/module"
	"game_server/module/rpc"
	"game_server/tlog"
	"game_server/xdlog"
)

type BattleAPI struct {
}

func init() {
	battle_api.SetInHandler(BattleAPI{})
}

func (api BattleAPI) StartBattle(session *net.Session, in *battle_api.StartBattle_In) {
	state := module.State(session)
	db := state.Database
	switch in.BattleType {
	// 区域关卡
	case battle_api.BATTLE_TYPE_MISSION, battle_api.BATTLE_TYPE_ARENA:
		module.Mission.StartLevelBattle(battle.BT_MISSION_LEVEL, int32(in.BattleId), session)

	// 资源关卡
	case battle_api.BATTLE_TYPE_RESOURCE:
		fail.When(!module.Player.IsOpenFunc(db, player_dat.FUNC_RESOURCE_LEVEL), "未开启资源关卡")
		module.Mission.StartLevelBattle(battle.BT_RESOURCE_LEVEL, int32(in.BattleId), session)

	// 多人关卡
	case battle_api.BATTLE_TYPE_MULTILEVEL:
		state := module.State(session)
		rpc.RemoteStartMultiLevel(state.PlayerId, func() {
			module.MultiLevel.StartBattle(state, int32(in.BattleId))
		})

	//难度关卡
	case battle_api.BATTLE_TYPE_HARD:
		fail.When(!module.Player.IsOpenFunc(db, player_dat.FUNC_HARD_LEVEL), "未开启难度关卡")
		module.Mission.StartLevelBattle(battle.BT_HARD_LEVEL, int32(in.BattleId), session)

	// 伙伴关卡（活动关卡)
	case battle_api.BATTLE_TYPE_BUDDY:
		module.Player.MustOpenFunc(db, player_dat.FUNC_ACTIVE_LEVLE)
		module.Mission.StartLevelBattle(battle.BT_BUDDY_LEVEL, int32(in.BattleId), session)

	// 灵宠关卡
	case battle_api.BATTLE_TYPE_PET:
		module.Player.MustOpenFunc(db, player_dat.FUNC_ACTIVE_LEVLE)
		module.Mission.StartLevelBattle(battle.BT_PET_LEVEL, int32(in.BattleId), session)

	// 魂侍关卡
	case battle_api.BATTLE_TYPE_GHOST:
		module.Player.MustOpenFunc(db, player_dat.FUNC_ACTIVE_LEVLE)
		module.Mission.StartLevelBattle(battle.BT_GHOST_LEVEL, int32(in.BattleId), session)

	//彩虹关卡
	case battle_api.BATTLE_TYPE_RAINBOW:
		fail.When(!module.Player.IsOpenFunc(db, player_dat.FUNC_RAINBOW_LEVEL), "未开彩虹关卡")
		module.Rainbow.StartRainbowLevel(session, int32(in.BattleId))

	//灵宠幻境PVE
	case battle_api.BATTLE_TYPE_PVE:
		fail.When(!module.Player.IsOpenFunc(db, player_dat.FUNC_BATTLE_PET), "灵宠功能未开启")
		module.PetVirtualEnv.StartPetVirtualEnvLevel(session, int32(in.BattleId))
	case battle_api.BATTLE_TYPE_FATE_BOX:
		//检查功能权值
		fail.When(!module.Player.IsOpenFunc(db, player_dat.FUNC_FATE_BOX), "命锁宝箱功能未开启")
		module.Mission.StartFateBoxLevelBattle(session, int32(in.BattleId))

	//仙山探险
	case battle_api.BATTLE_TYPE_DRIVING_EXPLORING:
		module.DrivingSword.StartExploreLevel(session)

	default:
		fail.When(true, "error battle type")
	}
}

func (api BattleAPI) NextRound(session *net.Session, in *battle_api.NextRound_In) {
	panic("废弃")
	state := module.State(session)

	if state.Battle.GetBattle().BatType == battle.BT_MULTI_LEVEL {
		if in.UseItem > 0 {
			num, ok := state.MultiLevelState.BattleItemInfo[in.UseItem]
			fail.When(!ok || num < 1, "没有可用的战斗道具")
			state.MultiLevelState.BattleItemInfo[in.UseItem] -= 1
			rpc.RemoteDeleteItem(state.PlayerId, in.UseItem, 1)
		}
	}

	state.Battle.NextRound(&module.NextRoundParams{
		Session:               session,
		PlayerId:              state.PlayerId,
		SkillIndex:            int(in.UseSkill),
		AutoFight:             in.AutoFight,
		SendNum:               in.SendNum,
		IsAttacker:            in.IsAttacker,
		Position:              int(in.Position),
		JobIndex:              int(in.JobIndex),
		UseItemId:             in.UseItem,
		UseSwordSoul:          in.UseSwordSoul,
		UseGhostSkillPosition: int(in.UseGhostSkillPosition),
		UseGhostSkillId:       int16(in.UseGhostSkillId),
		LaunchBattleTotem:     in.UseTotem,
	})
}

func (api BattleAPI) Escape(session *net.Session, in *battle_api.Escape_In) {
	state := module.State(session)
	state.Battle.Escape(session)
}

func (api BattleAPI) StartReady(session *net.Session, in *battle_api.StartReady_In) {
	/*
		TODO 废弃
		state := module.State(session)
		if state.Battle != nil {
			state.Battle.(*BattleRoom).StartReady(state.PlayerId)
		}
	*/
}

func (api BattleAPI) CallBattlePet(session *net.Session, in *battle_api.CallBattlePet_In) {
	state := module.State(session)
	var success bool
	if state.Battle.GetBattle().BatType == battle.BT_MULTI_LEVEL { //在多人关卡？
		success = callBattlePetInMultiLevel(session, in.GridNum)
	} else {
		success = callBattlePet(session, in.GridNum)
	}
	if !success {
		session.Send(&battle_api.CallBattlePet_Out{
			Success: false,
		})
	}
}

func (api BattleAPI) UseBuddySkill(session *net.Session, in *battle_api.UseBuddySkill_In) {
	state := module.State(session)
	if state.Battle != nil {
		state.Battle.UseBuddySkill(session, in.Pos, in.UseSkill)
		session.Send(&battle_api.UseBuddySkill_Out{
			Pos:      in.Pos,
			UseSkill: in.UseSkill,
		})
	}
}

func (api BattleAPI) StartBattleForHijackBoat(session *net.Session, in *battle_api.StartBattleForHijackBoat_In) {
	startBattleForHijackBoat(session, in.BoatId, in.Pid)
}

func (api BattleAPI) StartBattleForRecoverBoat(session *net.Session, in *battle_api.StartBattleForRecoverBoat_In) {
	startBattleForRecoverBoat(session, in.BoatId, in.Pid)
}

func (this BattleAPI) RoundReady(session *net.Session, in *battle_api.RoundReady_In) {
	state := module.State(session)
	if state.Battle != nil {
		state.Battle.PrepareReady(session, in.IsAuto)
	}
}

func (this BattleAPI) InitRound(session *net.Session, in *battle_api.InitRound_In) {
	state := module.State(session)
	if state.Battle != nil {
		state.Battle.InitRound(session)
	}
}

func (this BattleAPI) SetAuto(session *net.Session, in *battle_api.SetAuto_In) {
	state := module.State(session)
	if state.Battle != nil {
		state.Battle.SetAuto(session)
	}
}

func (this BattleAPI) CancelAuto(session *net.Session, in *battle_api.CancelAuto_In) {
	state := module.State(session)
	if state.Battle != nil {
		state.Battle.CancelAuto(session)
	}
}

func (this BattleAPI) SetSkill(session *net.Session, in *battle_api.SetSkill_In) {
	state := module.State(session)
	if state.Battle != nil {
		state.Battle.SetSkill(session, in.PosIdx, in.SkillIdx)
	}
}

func (this BattleAPI) UseItem(session *net.Session, in *battle_api.UseItem_In) {
	state := module.State(session)
	if state.Battle != nil {
		if state.Battle.GetBattle().BatType == battle.BT_MULTI_LEVEL {
			num, ok := state.MultiLevelState.BattleItemInfo[in.ItemId]
			fail.When(!ok || num < 1, "没有可用的战斗道具")
			state.MultiLevelState.BattleItemInfo[in.ItemId] -= 1
			rpc.RemoteDeleteItem(state.PlayerId, in.ItemId, 1)
		} else {
			module.Item.DelItemByItemId(state.Database, in.ItemId, 1, tlog.IFR_BATTLE_ITEM, xdlog.ET_MISSION_LEVEL_USE_ITEM)
		}
		state.Battle.UseItem(session, in.IsAttacker, in.Position, in.ItemId)
	}

}

func (this BattleAPI) UseGhost(session *net.Session, in *battle_api.UseGhost_In) {
	state := module.State(session)
	if state.Battle != nil {
		state.Battle.UseGhost(session, in.IsAttacker, in.Position)
	}

}

func (this BattleAPI) BattleReconnect(session *net.Session, in *battle_api.BattleReconnect_In) {
}
