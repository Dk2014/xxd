package battle

import (
	"core/fail"
	"core/net"
	"encoding/json"
	"game_server/battle"
	"game_server/module"
	"game_server/module/rpc"
)

func startBattleForHijackBoat(session *net.Session, boatId, targetPlayerId int64) {
	state := module.State(session)
	args := []*rpc.Args_NewPlayerFighter{{Pid: targetPlayerId, AutoFight: true}}
	rpc.RemoteNewPlayerFighter(args, func(replys []*rpc.Reply_NewPlayerFighter, errs []error) {
		for _, err := range errs {
			fail.When(err != nil, errs)
		}

		fighters := make([]*battle.Fighter, module.ALL_FIGHTER_POS_NUM)
		var battlePlayerInfo *battle.BattlePlayerInfo
		var battleTotemInfo [5]*battle.TotemInfo

		for _, reply := range replys {
			battlePlayerInfo = reply.BattlePlayerInfo
			err := json.Unmarshal(reply.TotemInfo, &battleTotemInfo)
			fail.When(err != nil, err)
			for _, raw := range reply.Fighters {
				fighter := new(battle.Fighter)
				err := json.Unmarshal(raw, fighter)
				fail.When(err != nil, err)
				fighters[fighter.Position-1] = fighter
				if fighter != nil {
					fighter.MaxHealth *= 2
					fighter.Health = fighter.MaxHealth
				}
			}
		}
		defendSide := &battle.SideInfo{
			Groups:    [][]*battle.Fighter{fighters},
			Fighters:  fighters,
			Players:   []*battle.BattlePlayerInfo{battlePlayerInfo},
			TotemInfo: battleTotemInfo,
		}
		attackSide, _ := module.NewBattleSideWithPlayerDatabase(state.Database, false, false, true)
		for _, f := range attackSide.Fighters {
			if f != nil {
				f.MaxHealth *= 2
				f.Health = f.MaxHealth
			}
		}
		state.MissionLevelState = module.NewMissionLevelState(battle.BT_HIJACK_BOAT, 0)
		state.MissionLevelState.LoadBuddySkill(state)
		state.MissionLevelState.LoadFighterAttribute(state)
		//state.Battle = module.Battle.NewBattleForVisiting(session, attackSide, defendSide)
		state.Battle = module.Battle.NewBattleForHiJackBoat(session, battle.BT_HIJACK_BOAT, attackSide, defendSide, boatId)
	})
}

func startBattleForRecoverBoat(session *net.Session, boatId, targetPlayerId int64) {
	state := module.State(session)
	args := []*rpc.Args_NewPlayerFighter{{Pid: targetPlayerId, AutoFight: true}}
	rpc.RemoteNewPlayerFighter(args, func(replys []*rpc.Reply_NewPlayerFighter, errs []error) {
		for _, err := range errs {
			fail.When(err != nil, errs)
		}

		fighters := make([]*battle.Fighter, module.ALL_FIGHTER_POS_NUM)
		var battlePlayerInfo *battle.BattlePlayerInfo
		var battleTotemInfo [5]*battle.TotemInfo

		for _, reply := range replys {
			battlePlayerInfo = reply.BattlePlayerInfo
			err := json.Unmarshal(reply.TotemInfo, &battleTotemInfo)
			fail.When(err != nil, err)
			for _, raw := range reply.Fighters {
				fighter := new(battle.Fighter)
				err := json.Unmarshal(raw, fighter)
				fail.When(err != nil, err)
				fighters[fighter.Position-1] = fighter
				if fighter != nil {
					fighter.MaxHealth *= 2
					fighter.Health = fighter.MaxHealth
				}
			}
		}
		defendSide := &battle.SideInfo{
			Groups:    [][]*battle.Fighter{fighters},
			Fighters:  fighters,
			Players:   []*battle.BattlePlayerInfo{battlePlayerInfo},
			TotemInfo: battleTotemInfo,
		}
		attackSide, _ := module.NewBattleSideWithPlayerDatabase(state.Database, false, false, true)
		for _, f := range attackSide.Fighters {
			if f != nil {
				f.MaxHealth *= 2
				f.Health = f.MaxHealth
			}
		}
		state.MissionLevelState = module.NewMissionLevelState(battle.BT_RECOVER_BOAT, 0)
		state.MissionLevelState.LoadBuddySkill(state)
		state.MissionLevelState.LoadFighterAttribute(state)
		//state.Battle = module.Battle.NewBattleForVisiting(session, attackSide, defendSide)
		state.Battle = module.Battle.NewBattleForRecoverBoat(session, battle.BT_RECOVER_BOAT, attackSide, defendSide, boatId)
	})
}

func init() {
	module.PrepareStoreEvent.Regisiter(PerpareStoreHandlerForEscort)
}
func PerpareStoreHandlerForEscort(session *net.Session) {
	state := module.State(session)
	if state.Battle != nil && state.Battle.GetBattle() != nil {
		if state.Battle.GetBattle().BatType == battle.BT_RECOVER_BOAT {
			if missionLvBattle, ok := state.Battle.(*Battle_MissionLevel); ok && missionLvBattle != nil {
				rpc.RemoteEscortBoatRecoverBattleLose(state.PlayerId, missionLvBattle.boatId)
			}
		}
	}
}
