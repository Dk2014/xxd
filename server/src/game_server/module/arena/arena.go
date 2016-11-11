package arena

import (
	"core/fail"
	"core/net"
	"core/time"
	"encoding/json"
	"game_server/api/protocol/arena_api"
	"game_server/battle"
	"game_server/dat/arena_award_box_dat"
	"game_server/dat/player_dat"
	"game_server/dat/vip_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/module/rpc"
	"game_server/tlog"
	"game_server/xdlog"
	"math"
)

func enter(session *net.Session) {
	state := module.State(session)
	playerArena := state.Database.Lookup.PlayerArena(state.PlayerId)

	if !time.IsInPointHour(player_dat.RESET_ARENA_TIMES_IN_HOUR, playerArena.BattleTime) {
		playerArena.DailyNum = 0
		playerArena.DailyFame = 0
		playerArena.DailyAwardItem = 0
		playerArena.BuyTimes = 0
	}
	state.Database.Update.PlayerArena(playerArena)

	rpc.RemoteEnterArena(state.PlayerId, func(reply *rpc.Reply_GlobalArenaEnterArena) {
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(state.PlayerId, func(db *mdb.Database) {
				module.Arena.InitPlayerArenaRank(db, state.PlayerId)
				playerArenaRank := db.Lookup.PlayerArenaRank(state.PlayerId)
				if module.Arena.RefreshAwardBox(playerArenaRank, reply.PlayerRank) {
					db.Update.PlayerArenaRank(playerArenaRank)
				}
				boxs := []arena_api.Enter_Out_AwardBox{}
				if playerArenaRank.Rank > 0 {
					boxs = append(boxs, arena_api.Enter_Out_AwardBox{Num: 0, Rank: playerArenaRank.Rank})
				}

				if playerArenaRank.Rank1 > 0 {
					boxs = append(boxs, arena_api.Enter_Out_AwardBox{Num: 1, Rank: playerArenaRank.Rank1})
				}

				if playerArenaRank.Rank2 > 0 {
					boxs = append(boxs, arena_api.Enter_Out_AwardBox{Num: 2, Rank: playerArenaRank.Rank2})
				}

				if playerArenaRank.Rank3 > 0 {
					boxs = append(boxs, arena_api.Enter_Out_AwardBox{Num: 3, Rank: playerArenaRank.Rank3})
				}
				session.Send(&arena_api.Enter_Out{
					Rank:           reply.PlayerRank,
					AwardBoxTime:   reply.BoxAwardTime,
					DailyNum:       playerArena.DailyNum,
					WinTimes:       playerArena.WinTimes,
					NewRecordNum:   playerArena.NewRecordCount,
					FailedCdTime:   playerArena.FailedCdTime,
					DailyAwardItem: playerArena.DailyAwardItem,
					DailyFame:      playerArena.DailyFame,
					Ranks:          reply.Ranks,
					AwardBox:       boxs,
					BuyTimes:       playerArena.BuyTimes,
				})
			})
		})
	})
}

func cleanCDTime(session *net.Session) {
	state := module.State(session)
	tlog.PlayerSystemModuleFlowLog(state.Database, tlog.SMT_ARENA)
	playerVIPInfo := state.Database.Lookup.PlayerVip(state.PlayerId)
	fail.When(!vip_dat.HaveVIPPrivilege(playerVIPInfo.Level, vip_dat.BIWUCHANGTEQUAN), "no privilege")
	playerArena := state.Database.Lookup.PlayerArena(state.PlayerId)

	nowTime := time.GetNowTime()
	fail.When(nowTime > playerArena.FailedCdTime, "incorrect call cleanCDTime. no cd time")

	// 秒换分
	dtMin := math.Ceil((float64(playerArena.FailedCdTime - nowTime)) / 60.0)
	// 分换元宝
	costIngot := dtMin * arena_award_box_dat.CD_TIME_ONE_MIN_COST_INGOT
	if costIngot > 0 {
		module.Player.DecMoney(state.Database, state.MoneyState, int64(costIngot), player_dat.INGOT, tlog.MFR_ARENA_CD, xdlog.ET_ARENA_CD)
		playerArena.FailedCdTime = 0
		state.Database.Update.PlayerArena(playerArena)
	}

	session.Send(&arena_api.CleanFailedCdTime_Out{
		FailedCdTime: playerArena.FailedCdTime,
	})
}

func sendRecords(session *net.Session) {
	state := module.State(session)
	playerArena := state.Database.Lookup.PlayerArena(state.PlayerId)

	playerArena.NewRecordCount = 0
	state.Database.Update.PlayerArena(playerArena)

	var recoredLimit int8 = arena_award_box_dat.MAX_RESERVED_RECORD

	rsp := &arena_api.GetRecords_Out{}
	state.Database.Select.PlayerArenaRecord(func(row *mdb.PlayerArenaRecordRow) {
		if recoredLimit <= 0 {
			row.Break()
			return
		}
		recoredLimit--
		rsp.Records = append(rsp.Records, arena_api.GetRecords_Out_Records{
			Mode:           row.Mode(),
			Time:           row.RecordTime(),
			OldRank:        row.OldRank(),
			NewRank:        row.NewRank(),
			FightNum:       row.FightNum(),
			TargetPid:      row.TargetPid(),
			TargetNick:     []byte(row.TargetNick()),
			TargetOldRank:  row.TargetOldRank(),
			TargetNewRank:  row.TargetNewRank(),
			TargetFightNum: row.TargetFightNum(),
		})
	})
	session.Send(rsp)
}

func startBattle(session *net.Session, targetPlayerId int64, targetPlayerRank int32) {
	state := module.State(session)
	tlog.PlayerSystemModuleFlowLog(state.Database, tlog.SMT_ARENA)

	playerArena := state.Database.Lookup.PlayerArena(state.PlayerId)
	if !time.IsInPointHour(player_dat.RESET_ARENA_TIMES_IN_HOUR, playerArena.BattleTime) {
		playerArena.DailyNum = 0
		playerArena.DailyFame = 0
		playerArena.DailyAwardItem = 0
		playerArena.BuyTimes = 0
	}

	fail.When(playerArena.DailyNum >= module.VIP.GetPrivilegeTimesByDB(state.Database, vip_dat.BIWUCHANGCISHU)+arena_award_box_dat.MAX_DAILY_NUM, "arena daily_num full ")
	fail.When(playerArena.FailedCdTime > time.GetNowTime(), "arena failed cd time")

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
			battlePlayerInfo.Auto = true
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

		rpc.RemoteGetTargetPlayerRank(state.PlayerId, targetPlayerId, func(reply *rpc.Reply_GlobalArenaGetTargetPlayerRank) {
			mdb.GlobalExecute(func(globalDB *mdb.Database) {
				globalDB.AgentExecute(state.PlayerId, func(db *mdb.Database) {
					attackSide, fightNum := module.NewBattleSideWithPlayerDatabase(db, false, true, true)
					for _, f := range attackSide.Fighters {
						if f != nil {
							f.MaxHealth *= 2
							f.Health = f.MaxHealth
						}
					}
					state.Battle = module.Battle.NewBattleForPVE(session, battle.BT_ARENA, attackSide, defendSide, true, false)
					state.ArenaState = &module.PlayerArenaState{
						PlayerRank:       reply.PlayerRank,
						PlayerFightNum:   fightNum,
						TargetPlayerRank: targetPlayerRank,
						TargetInfo:       reply.TargetPlayerInfo,
					}

					playerArena.DailyNum += 1
					playerArena.BattleTime = time.GetNowTime()
					db.Update.PlayerArena(playerArena)
				})
			})
		})
	})
}
