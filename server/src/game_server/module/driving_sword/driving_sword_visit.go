package driving_sword

import (
	"core/fail"
	"core/net"
	"encoding/json"
	"game_server/api/protocol/driving_sword_api"
	"game_server/battle"
	"game_server/dat/driving_sword_dat"
	"game_server/dat/player_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/module/rpc"
	"game_server/tlog"
)

type targetStatusInfo struct {
	Pid             int64
	Nick            string
	RoleId          int8
	Level           int16
	FriendshipLevel int16
	FashionId       int16
	FightNum        int32
}

func startBattleVisting(session *net.Session) {
	state := module.State(session)

	currentInfo := state.Database.Lookup.PlayerDrivingSwordInfo(state.PlayerId)

	//构造玩家的信息
	attackSide, _ := module.NewBattleSideWithPlayerDatabase(state.Database, false, false, false)
	for _, f := range attackSide.Fighters {
		if f != nil {
			f.MaxHealth *= 2
			f.Health = f.MaxHealth
		}
	}

	var defendSide *battle.SideInfo
	var targetSideInfoBytes []byte
	state.Database.Select.PlayerDrivingSwordEventVisiting(func(row *mdb.PlayerDrivingSwordEventVisitingRow) {
		if row.X() == currentInfo.CurrentX && row.Y() == currentInfo.CurrentY && row.CloudId() == currentInfo.CurrentCloud {
			if row.Status() == 0 && row.TargetSideState() != nil {
				targetSideInfoBytes = row.TargetSideState()
			}
		}
	})
	replys := make([]*rpc.Reply_NewPlayerFighter, 0)
	//容错处理 对于json解析错误或者记录信息为空
	err := json.Unmarshal(targetSideInfoBytes, &replys)
	if targetSideInfoBytes == nil || err != nil {
		callback := func(dside *battle.SideInfo) {
			//TODO 最好是给客户端刷新一下数据，然后客户端在发起一次战斗
			state.MissionLevelState = module.NewMissionLevelState(battle.BT_DRIVING_SWORD_BF_LEVEL, 0)
			//state.MissionLevelState.MaxRound = mission_dat.PVE_LEVEL_MAX_ROUND
			state.MissionLevelState.LoadBuddySkill(state)
			state.MissionLevelState.LoadFighterAttribute(state)
			state.MissionLevelState.LevelType = battle.BT_DRIVING_SWORD_BF_LEVEL
			state.Battle = module.Battle.NewBattleForVisiting(session, attackSide, defendSide)
		}
		visitingMoutian(session, false, callback)
		return
	} else {
		fighters := make([]*battle.Fighter, module.ALL_FIGHTER_POS_NUM)
		var battlePlayerInfo *battle.BattlePlayerInfo
		battleTotemInfo := [5]*battle.TotemInfo{}

		for _, reply := range replys {
			json.Unmarshal(reply.TotemInfo, &battleTotemInfo)
			battlePlayerInfo = reply.BattlePlayerInfo
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

		clientPid := battlePlayerInfo.PlayerId
		if clientPid == state.PlayerId {
			clientPid = 1
		}
		battlePlayerInfo.PlayerId = clientPid
		for _, fighter := range fighters {
			if fighter != nil {
				fighter.PlayerId = clientPid
			}
		}
		defendSide = &battle.SideInfo{
			Groups:    [][]*battle.Fighter{fighters},
			Fighters:  fighters,
			Players:   []*battle.BattlePlayerInfo{battlePlayerInfo},
			TotemInfo: battleTotemInfo,
		}
	}

	//state.MissionLevelState = module.NewMissionLevelState()
	state.MissionLevelState = module.NewMissionLevelState(battle.BT_DRIVING_SWORD_BF_LEVEL, 0)
	//state.MissionLevelState.MaxRound = mission_dat.PVE_LEVEL_MAX_ROUND
	state.MissionLevelState.LoadBuddySkill(state)
	state.MissionLevelState.LoadFighterAttribute(state)
	state.Battle = module.Battle.NewBattleForVisiting(session, attackSide, defendSide)
}

func visitingMoutian(session *net.Session, isRequest bool, callback func(*battle.SideInfo)) (defendSide *battle.SideInfo) {
	state := module.State(session)
	row := state.Database.Lookup.PlayerDrivingSwordInfo(state.PlayerId)
	current_cloud := row.CurrentCloud
	current_x := row.CurrentX
	current_y := row.CurrentY
	playerUsed := make(map[int64]int64)
	state.Database.Select.PlayerDrivingSwordEventVisiting(func(row *mdb.PlayerDrivingSwordEventVisitingRow) {
		if row.CloudId() == current_cloud && row.TargetPid() > 0 {
			playerUsed[row.TargetPid()] = row.TargetPid()
		}

	})
	var record *mdb.PlayerDrivingSwordEventVisiting
	state.Database.Select.PlayerDrivingSwordEventVisiting(func(row *mdb.PlayerDrivingSwordEventVisitingRow) {
		if row.CloudId() == current_cloud && row.X() == current_x && row.Y() == current_y {
			record = row.GoObject()
			rpc.RemoteGetTargetFighterForDrivingSwordVisting(state.PlayerId, playerUsed, func(reply *rpc.Reply_GetTargetFighterForDrivingSwordVisting) {
				record.TargetPid = reply.TargetPid
				args := []*rpc.Args_NewPlayerFighter{{Pid: reply.TargetPid, AutoFight: true}}
				rpc.RemoteNewPlayerFighter(args, func(newFighterReplys []*rpc.Reply_NewPlayerFighter, errs []error) {
					for _, err := range errs {
						fail.When(err != nil, errs)
					}
					bytes, err := json.Marshal(newFighterReplys)
					fail.When(err != nil, err)
					record.TargetSideState = bytes
					fighters := make([]*battle.Fighter, module.ALL_FIGHTER_POS_NUM)
					var battlePlayerInfo *battle.BattlePlayerInfo

					var mainRoleFashion int16
					var mainRoleFriendshipLevel int16
					for _, newFighterReply := range newFighterReplys {
						battlePlayerInfo = newFighterReply.BattlePlayerInfo
						for _, raw := range newFighterReply.Fighters {
							fighter := new(battle.Fighter)
							err := json.Unmarshal(raw, fighter)
							fail.When(err != nil, err)
							if fighter.Kind == battle.FK_PLAYER {
								mainRoleFashion = fighter.FashionId
								mainRoleFriendshipLevel = fighter.FriendshipLevel
							}
							fighters[fighter.Position-1] = fighter
							if fighter != nil {
								fighter.MaxHealth *= 2
								fighter.Health = fighter.MaxHealth
							}
						}
					}

					defendSide = &battle.SideInfo{
						Groups:   [][]*battle.Fighter{fighters},
						Fighters: fighters,
						Players:  []*battle.BattlePlayerInfo{battlePlayerInfo},
					}
					if callback != nil {
						callback(defendSide)
					}

					targetInfo := targetStatusInfo{}
					targetInfo.Nick = string(reply.TargetPlayerInfo.PlayerNick)
					targetInfo.Pid = reply.TargetPid
					targetInfo.RoleId = reply.TargetPlayerInfo.RoleId
					targetInfo.Level = reply.TargetPlayerInfo.RoleLevel
					targetInfo.FightNum = reply.TargetPlayerInfo.FightNum
					targetInfo.FriendshipLevel = mainRoleFriendshipLevel
					targetInfo.FashionId = mainRoleFashion

					bytess, _ := json.Marshal(targetInfo)
					record.TargetStatus = string(bytess)
					state.Database.Update.PlayerDrivingSwordEventVisiting(record)

					if isRequest {
						out := &driving_sword_api.VisitMountain_Out{}
						out.Nick = []byte(targetInfo.Nick)
						out.Pid = reply.TargetPid
						out.RoleId = targetInfo.RoleId
						out.Level = targetInfo.Level
						out.FriendshipLevel = targetInfo.FriendshipLevel
						out.FashionId = targetInfo.FashionId
						out.FightNum = targetInfo.FightNum
						session.Send(out)
					}

				})
			})
		}
	})
	return
}

func getDrivingSwordAward(session *net.Session, award *driving_sword_dat.DrivingSwordAward, xdEventType int32) {
	state := module.State(session)
	if award.Item1Num > 0 {
		module.Item.AddItem(state.Database, award.Item1, int16(award.Item1Num), 1, xdEventType, "")
	}
	if award.Item2Num > 0 {
		module.Item.AddItem(state.Database, award.Item2, int16(award.Item2Num), 1, xdEventType, "")
	}
	if award.Item3Num > 0 {
		module.Item.AddItem(state.Database, award.Item3, int16(award.Item3Num), 1, xdEventType, "")
	}
	if award.Coins > 0 {
		module.Player.IncMoney(state.Database, state.MoneyState, award.Coins, player_dat.COINS, tlog.MFR_BUY_COINS, xdEventType, "")
	}
}
