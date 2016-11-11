package rpc

import (
	"core/time"
	"game_server/api/protocol/arena_api"
	"game_server/dat/arena_award_box_dat"
	"game_server/dat/event_dat"
	"game_server/dat/push_notify_dat"
	"game_server/mdb"
	"game_server/module"
	. "game_server/rpc"
	"game_server/tlog"
)

/*
	玩家比武场数据初始化
*/
type Args_ArenaInitPlayerArenaRank struct {
	RPCArgTag
	PlayerId int64
}

type Reply_ArenaInitPlayerArenaRank struct {
}

func (this *RemoteServe) ArenaInitPlayerArenaRank(args *Args_ArenaInitPlayerArenaRank, reply *Reply_ArenaInitPlayerArenaRank) error {
	return Remote.Serve(mdb.RPC_Remote_ArenaInitPlayerArenaRank, args, mdb.TRANS_TAG_RPC_Serve_ArenaInitPlayerArenaRank, func() error {
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(args.PlayerId, func(db *mdb.Database) {
				if db.Lookup.PlayerArenaRank(args.PlayerId) == nil {
					db.Insert.PlayerArenaRank(&mdb.PlayerArenaRank{
						Pid:   args.PlayerId,
						Rank:  0,
						Rank1: 0,
						Rank2: 0,
						Rank3: 0,
						Time:  time.GetNowTime(),
					})
				}
			})
		})
		return nil
	})
}

func RemoteInitPlayerArenaRank(pid int64) {
	args := &Args_ArenaInitPlayerArenaRank{PlayerId: pid}
	serverId, _ := module.GetServerIdWithPlayerId(pid)
	Remote.Call(serverId, mdb.RPC_Remote_ArenaInitPlayerArenaRank, args, &Reply_ArenaInitPlayerArenaRank{}, mdb.TRANS_TAG_RPC_Call_ArenaInitPlayerArenaRank, nil)
}

/*
	添加比武场战报
*/
type Args_ArenaAddRecord struct {
	RPCArgTag
	Record *mdb.PlayerArenaRecord
}

type Reply_ArenaAddRecord struct {
}

func (this *RemoteServe) AddRecord(args *Args_ArenaAddRecord, reply *Reply_ArenaAddRecord) error {
	return Remote.Serve(mdb.RPC_Remote_AddRecord, args, mdb.TRANS_TAG_RPC_Serve_AddRecord, func() error {
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(args.Record.Pid, func(db *mdb.Database) {
				db.Insert.PlayerArenaRecord(args.Record)

				tlog.PlayerPvpFlowLog(db, int32(args.Record.OldRank), int32(args.Record.NewRank), int32(args.Record.Mode))

				playerArena := db.Lookup.PlayerArena(db.PlayerId())
				// 挑战方胜利
				if args.Record.Mode == arena_award_box_dat.ARENA_ATTACKED_FAILED {
					if playerArena.WinTimes >= 0 {
						playerArena.WinTimes = -1
					}
				} else {
					if playerArena.WinTimes < 0 {
						playerArena.WinTimes = 1
					} else {
						playerArena.WinTimes += 1
					}
				}

				playerArena.NewRecordCount += 1
				db.Update.PlayerArena(playerArena)

				// 同步玩家比武场趋势到互动服
				RemoteUpdatePlayerArenaTrendWin(playerArena.Pid, playerArena.WinTimes)
				//rpc 更新JSON比武场活动
				module.Event.UpdateJsonEventArenaRank(db, args.Record.NewRank)

				if s, ok := module.Player.GetPlayerOnline(db.PlayerId()); ok {
					s.Send(&arena_api.NotifyArenaInfo_Out{
						NotifyType: arena_api.NotifyArenaMode(args.Record.Mode),
						Pid:        args.Record.TargetPid,
						Nick:       []byte(args.Record.TargetNick),
						Num:        args.Record.TargetNewRank,
					})
					//更新比武场活动
					state := module.State(s)
					module.UpdateEventArenaRank(state, args.Record.NewRank)

				} else {
					// 记录离线玩家战报数，在登陆时客户端获取
					playerInfo := db.Lookup.PlayerInfo(db.PlayerId())
					playerInfo.NewArenaReportNum += 1
					db.Update.PlayerInfo(playerInfo)
					if module.PushNotify.EnabledPushNotify(db, push_notify_dat.ARENAATTACK) {
						module.PushNotify.SendNotification(db.PlayerId(), push_notify_dat.ARENAATTACK)
					}

					//更新离线玩家的比武场活动信息
					event, _ := event_dat.GetEventInfoById(event_dat.EVENT_ARENA_RANK_AWARDS)
					if event_dat.CheckEventTime(event, event_dat.NOT_END) {
						playerEventsInfo := db.Lookup.PlayerEventAwardRecord(db.PlayerId())
						eventInfoList := module.NewEventInfoList()
						eventInfoList.Decode(playerEventsInfo.RecordBytes)
						if eventInfo := eventInfoList.GetPlayerEventInfoById(event_dat.EVENT_ARENA_RANK_AWARDS); eventInfo.Awarded > 0 {
							eventInfoList.ClearState(db, event_dat.EVENT_ARENA_RANK_AWARDS)
						}
						eventInfoList.UpdateMax(db, event_dat.EVENT_ARENA_RANK_AWARDS, args.Record.NewRank)
					}
				}

			})
		})

		return nil
	})
}

func RemoteArenaAddRecord(record *mdb.PlayerArenaRecord) {
	args := &Args_ArenaAddRecord{Record: record}
	serverId, _ := module.GetServerIdWithPlayerId(record.Pid)
	Remote.Call(serverId, mdb.RPC_Remote_AddRecord, args, &Reply_ArenaAddRecord{}, mdb.TRANS_TAG_RPC_Call_AddRecord, nil)
}
