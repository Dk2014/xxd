package rpc

import (
	"core/fail"
	"core/log"
	"encoding/json"
	"errors"
	"game_server/api/protocol/role_api"
	"game_server/battle"
	. "game_server/config"
	"game_server/dat/arena_award_box_dat"
	"game_server/dat/player_dat"
	"game_server/mdb"
	"game_server/module"
	. "game_server/rpc"
	"game_server/tlog"
	"game_server/xdlog"
)

/*

	获取玩家信息

*/

type Args_GetInfo struct {
	RPCArgTag
	PlayerId      int64
	OtherPlayerId int64
}

type Reply_GetInfo struct {
	PlayerInfo *role_api.GetPlayerInfo_Out
}

func (this *RemoteServe) GetPlayerInfo(args *Args_GetInfo, reply *Reply_GetInfo) error {
	return Remote.Serve(mdb.RPC_Remote_GetPlayerInfo, args, mdb.TRANS_TAG_RPC_Serve_GetPlayerInfo, func() error {
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(args.OtherPlayerId, func(db *mdb.Database) {
				reply.PlayerInfo = &role_api.GetPlayerInfo_Out{}
				module.Role.GetOtherPlayerInfo(args.PlayerId, db, &reply.PlayerInfo.Player)
			})
		})
		return nil
	})
}

func RemoteGetInfo(state *module.SessionState, remotePlayerId int64, callback func(*Reply_GetInfo, error)) {
	reply := &Reply_GetInfo{}
	serverId, _ := module.GetServerIdWithPlayerId(remotePlayerId)

	args := &Args_GetInfo{PlayerId: state.PlayerId, OtherPlayerId: remotePlayerId}
	Remote.Call(serverId, mdb.RPC_Remote_GetPlayerInfo, args, reply, mdb.TRANS_TAG_RPC_Call_GetPlayerInfo, func(err error) {
		callback(reply, err)
	})
}

/*
	通过平台openid 和 游戏服id 获取玩家信息
*/
type Args_GetInfoWithOpenId struct {
	RPCArgTag
	OpenId       []byte
	PlayerId     int64
	GameServerId int
}

type Reply_GetInfoWithOpenId struct {
	PlayerInfo *role_api.GetPlayerInfoWithOpenid_Out
}

func (this *RemoteServe) GetPlayerInfoWithOpenId(args *Args_GetInfoWithOpenId, reply *Reply_GetInfoWithOpenId) error {
	return Remote.Serve(mdb.RPC_Remote_GetPlayerInfoWithOpenId, args, mdb.TRANS_TAG_RPC_Serve_GetPlayerInfoWithOpenId, func() error {
		pid, ok := module.Player.GetPlayerByUsername(string(args.OpenId))
		if !ok {
			return errors.New("[GetPlayerInfoWithOpenId] can't found player with OpenId")
		}

		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(pid, func(db *mdb.Database) {
				reply.PlayerInfo = &role_api.GetPlayerInfoWithOpenid_Out{}
				module.Role.GetOtherPlayerInfo(args.PlayerId, db, &reply.PlayerInfo.Player)
			})
		})
		return nil
	})
}
func RemoteGetInfoWithOpenId(state *module.SessionState, OpenId []byte, GameServerId int, callback func(*Reply_GetInfoWithOpenId, error)) {
	reply := &Reply_GetInfoWithOpenId{}
	args := &Args_GetInfoWithOpenId{OpenId: OpenId, GameServerId: GameServerId, PlayerId: state.PlayerId}
	Remote.Call(args.GameServerId, mdb.RPC_Remote_GetPlayerInfoWithOpenId, args, reply, mdb.TRANS_TAG_RPC_Call_GetPlayerInfoWithOpenId, func(err error) {
		callback(reply, err)
	})
}

/*

	获取玩家战场数据

*/

type Args_NewPlayerFighter struct {
	RPCArgTag
	Pid          int64
	FormRoleInfo []*module.InFormRoleInfo
	Tactical     int8
	AutoFight    bool
}

type Reply_NewPlayerFighter struct {
	Fighters         [][]byte //*battle.Fighter
	TotemInfo        []byte   // [5]*battle.Fighter
	BattlePlayerInfo *battle.BattlePlayerInfo
}

func setFighter(db *mdb.Database, fighters []*battle.Fighter, arg *Args_NewPlayerFighter, reply *Reply_NewPlayerFighter) {
	module.GetBattleBiz.SetFighters(db, arg.FormRoleInfo, fighters, false, false, module.FIGHT_FOR_ALL)

	reply.BattlePlayerInfo = &battle.BattlePlayerInfo{
		PlayerId: arg.Pid,
		JobIndex: int(arg.Tactical),
	}

}

func (this *RemoteServe) NewPlayerFighter(args *Args_NewPlayerFighter, reply *Reply_NewPlayerFighter) error {
	return Remote.Serve(mdb.RPC_Remote_NewPlayerFighter, args, mdb.TRANS_TAG_RPC_Serve_NewPlayerFighter, func() error {
		fighters := make([]*battle.Fighter, module.ALL_FIGHTER_POS_NUM)

		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(args.Pid, func(db *mdb.Database) {
				if args.FormRoleInfo == nil {
					side, _ := module.NewBattleSideWithPlayerDatabase(db, false, args.AutoFight, false)
					fighters = side.Fighters
					//side.Groups = nil
					//side.Fighters = nil
					reply.BattlePlayerInfo = side.Players[0]
				} else {
					setFighter(db, fighters, args, reply)
				}
				battleTotemInfo := module.NewBattleTotemInfo(db)
				rawTotemInfo, err := json.Marshal(&battleTotemInfo)
				fail.When(err != nil, err)
				reply.TotemInfo = rawTotemInfo
			})
		})

		for _, fighter := range fighters {
			if fighter != nil {
				raw, err := json.Marshal(fighter)
				fail.When(err != nil, err)
				reply.Fighters = append(reply.Fighters, raw)
			}
		}

		return nil
	})
}

func RemoteNewPlayerFighter(args []*Args_NewPlayerFighter, callback func([]*Reply_NewPlayerFighter, []error)) {
	var replys []*Reply_NewPlayerFighter

	var remoteArgs []RPCArg
	var remoteSid []int
	var remoteReplys []interface{}
	var remoteCount int

	for _, arg := range args {
		serverId, _ := module.GetServerIdWithPlayerId(arg.Pid)
		remoteCount++
		remoteArgs = append(remoteArgs, arg)
		remoteSid = append(remoteSid, serverId)
		remoteReplys = append(remoteReplys, &Reply_NewPlayerFighter{})
	}

	Remote.BatchCall(mdb.RPC_Remote_NewPlayerFighter, remoteSid, remoteArgs, remoteReplys, mdb.TRANS_TAG_RPC_Call_NewPlayerFighter, func(errs []error) {
		for i := 0; i < remoteCount; i++ {
			replys = append(replys, remoteReplys[i].(*Reply_NewPlayerFighter))
		}

		callback(replys, errs)
	})
}

/*
	比武场宝箱奖励
*/

type Args_AwardArenaBox struct {
	RPCArgTag
	PlayerId   int64
	Num        int8
	Rank       int32
	Moneystate *module.MoneyState
}

type Reply_AwardArenaBox struct {
}

func (this *RemoteServe) AwardArenaBox(args *Args_AwardArenaBox, reply *Reply_AwardArenaBox) error {
	return Remote.Serve(mdb.RPC_Remote_AwardArenaBox, args, mdb.TRANS_TAG_RPC_Serve_AwardArenaBox, func() error {
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(args.PlayerId, func(db *mdb.Database) {
				playerArenaBox := db.Lookup.PlayerArenaRank(args.PlayerId)
				tlog.PlayerSystemModuleFlowLog(db, tlog.SMT_ARENA)
				// 领取宝箱的时候，先更新宝箱状态，避免发生0点的边界问题
				if module.Arena.RefreshAwardBox(playerArenaBox, args.Rank) {
					db.Update.PlayerArenaRank(playerArenaBox)
				}

				var rank int32
				var costIngot int64

				switch args.Num {
				case 0:
					fail.When(playerArenaBox.Rank == 0, "arena box alreay open")
					rank = playerArenaBox.Rank
					playerArenaBox.Rank = 0
				case 1:
					fail.When(playerArenaBox.Rank1 == 0, "arena box alreay open")
					costIngot = arena_award_box_dat.GETBACK_AWARD_BOX_DAY1_COST_INGOT

					rank = playerArenaBox.Rank1
					playerArenaBox.Rank1 = 0
				case 2:
					fail.When(playerArenaBox.Rank2 == 0, "arena box alreay open")
					costIngot = arena_award_box_dat.GETBACK_AWARD_BOX_DAY2_COST_INGOT

					rank = playerArenaBox.Rank2
					playerArenaBox.Rank2 = 0
				case 3:
					fail.When(playerArenaBox.Rank3 == 0, "arena box alreay open")
					costIngot = arena_award_box_dat.GETBACK_AWARD_BOX_DAY3_COST_INGOT

					rank = playerArenaBox.Rank3
					playerArenaBox.Rank3 = 0
				}

				if costIngot > 0 {
					session, ok := module.Player.GetPlayerOnline(args.PlayerId)
					fail.When(!ok, "AwardArenaBox player offline")

					state := module.State(session)
					fail.When(!module.Player.CheckMoney(state, costIngot, player_dat.INGOT), "get back arena box ingot not enough")
					module.Player.DecMoney(db, state.MoneyState, costIngot, player_dat.INGOT, tlog.MFR_ARENA_AWARD_BOX, xdlog.ET_ARENA_AWARD_BOX)
				}

				db.Update.PlayerArenaRank(playerArenaBox)

				box := arena_award_box_dat.GetAwardWithRank(rank)
				if box.Coins > 0 {
					module.Player.IncMoney(db, args.Moneystate, int64(box.Coins), player_dat.COINS, tlog.MFR_ARENA_AWARD_BOX, xdlog.ET_ARENA_AWARD_BOX, "")
				}

				if box.Ingot > 0 {
					module.Player.IncMoney(db, args.Moneystate, int64(box.Ingot), player_dat.INGOT, tlog.MFR_ARENA_AWARD_BOX, xdlog.ET_ARENA_AWARD_BOX, "")
				}
				if box.Fame > 0 {
					module.Player.AddFame(db, box.Fame)
				}

				if box.ItemId > 0 && box.ItemNum > 0 {
					module.Item.AddItem(db, box.ItemId, box.ItemNum, tlog.IFR_ARENA_AWARD_BOX, xdlog.ET_ARENA_AWARD_BOX, "")
				}

				if box.Item2Id > 0 && box.Item2Num > 0 {
					module.Item.AddItem(db, box.Item2Id, box.Item2Num, tlog.IFR_ARENA_AWARD_BOX, xdlog.ET_ARENA_AWARD_BOX, "")
				}
				tlog.PlayerPvpFlowLog(db, 0, 0, tlog.ARENABOX)
			})
		})
		return nil
	})
}

func RemoteAwardArenaBox(playerId int64, num int8, moneystate *module.MoneyState, rank int32, callback func()) {
	reply := new(Reply_AwardArenaBox)
	args := &Args_AwardArenaBox{
		PlayerId:   playerId,
		Num:        num,
		Rank:       rank,
		Moneystate: moneystate,
	}

	serverId, _ := module.GetServerIdWithPlayerId(playerId)
	Remote.Call(serverId, mdb.RPC_Remote_AwardArenaBox, args, reply, mdb.TRANS_TAG_RPC_Call_AwardArenaBox, func(err error) {
		fail.When(err != nil, err)
		callback()
	})
}

/*
	获取各个游戏进程的在线玩家数(由互动服务器调用)
*/
type Args_GetOnlineNumber struct {
	RPCArgTag
}

type Reply_GetOnlineNumber struct {
	Count map[int32]int64
}

func (this *RemoteServe) GetOnlineNumber(args *Args_GetOnlineNumber, reply *Reply_GetOnlineNumber) error {
	return Remote.Serve(mdb.RPC_Remote_GetOnlineNumber, args, mdb.TRANS_TAG_RPC_Serve_GetOnlineNumber, func() error {
		reply.Count = module.Player.OnlinePlayerNum()
		return nil
	})
}

func RemoteGetOnlineNumber(callback func(map[int32]int64)) {
	var remoteArgs []RPCArg
	var remoteSid []int
	var remoteReplys []interface{}

	for _, gsid := range Remote.GetRPCServerIds() {
		if gsid/10 != ServerCfg.ServerId/10 || gsid == ServerCfg.ServerId {
			continue
		}

		remoteArgs = append(remoteArgs, &Args_GetOnlineNumber{})
		remoteSid = append(remoteSid, gsid)
		remoteReplys = append(remoteReplys, &Reply_GetOnlineNumber{})
	}

	Remote.BatchCall(mdb.RPC_Remote_GetOnlineNumber, remoteSid, remoteArgs, remoteReplys, mdb.TRANS_TAG_RPC_Call_GetOnlineNumber, func(errs []error) {
		count := make(map[int32]int64, 0)
		for _, replay := range remoteReplys {
			for k, v := range replay.(*Reply_GetOnlineNumber).Count {
				count[k] += v
			}
		}
		callback(count)
	})
}

//
type Args_DecMoney struct {
	RPCArgTag
	Pid             int64
	Num             int64
	MType           int
	MoneyFlowReason int32
	XdEventType     int32
}

type Reply_DecMoney struct{}

func (srv *RemoteServe) DecMoney(args *Args_DecMoney, reply *Reply_DecMoney) error {
	return Remote.Serve(mdb.RPC_Remote_DecMoney, args, mdb.TRANS_TAG_RPC_Serve_DecMoney, func() error {
		session, ok := module.Player.GetPlayerOnline(args.Pid)
		fail.When(!ok, "找不到玩家会话")
		state := module.State(session)
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(args.Pid, func(db *mdb.Database) {
				module.Player.DecMoney(db, state.MoneyState, args.Num, args.MType, args.MoneyFlowReason, args.XdEventType)
			})
		})
		return nil
	})
}

func RemoteDecMoney(pid int64, num int64, mtype int, moneyFlowReason, xdEventType int32, callback func()) {
	reply := new(Reply_DecMoney)
	args := &Args_DecMoney{
		Pid:             pid,
		Num:             num,
		MType:           mtype,
		MoneyFlowReason: moneyFlowReason,
		XdEventType:     xdEventType,
	}
	serverId, _ := module.GetServerIdWithPlayerId(pid)
	Remote.Call(serverId, mdb.RPC_Remote_DecMoney, args, reply, mdb.TRANS_TAG_RPC_Call_DecMoney, func(err error) {
		if err == nil {
			callback()
		} else {
			log.Errorf("RemoteDecMoney %v\n", err)
		}
	})
}

type Args_IncMoney struct {
	RPCArgTag
	Pid             int64
	Num             int64
	MType           int
	MoneyFlowReason int32
	XdEventType     int32
}

type Reply_IncMoney struct{}

func (srv *RemoteServe) IncMoney(args *Args_IncMoney, reply *Reply_IncMoney) error {
	return Remote.Serve(mdb.RPC_Remote_IncMoney, args, mdb.TRANS_TAG_RPC_Serve_IncMoney, func() error {
		session, ok := module.Player.GetPlayerOnline(args.Pid)
		fail.When(!ok, "找不到玩家会话")
		state := module.State(session)
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(args.Pid, func(db *mdb.Database) {
				module.Player.IncMoney(db, state.MoneyState, args.Num, args.MType, args.MoneyFlowReason, args.XdEventType, "")
			})
		})
		return nil
	})
}

func RemoteIncMoney(pid int64, num int64, mtype int, moneyFlowReason, xdEventType int32, callback func()) {
	reply := new(Reply_IncMoney)
	args := &Args_IncMoney{
		Pid:             pid,
		Num:             num,
		MType:           mtype,
		MoneyFlowReason: moneyFlowReason,
		XdEventType:     xdEventType,
	}
	serverId, _ := module.GetServerIdWithPlayerId(pid)
	Remote.Call(serverId, mdb.RPC_Remote_IncMoney, args, reply, mdb.TRANS_TAG_RPC_Call_IncMoney, func(err error) {
		if err == nil {
			if callback != nil {
				callback()
			}
		} else {
			log.Errorf("RemoteIncMoney %v\n", err)
		}
	})
}

type Args_AddFame struct {
	RPCArgTag
	Pid    int64
	Num    int64
	System int32
}

type Reply_AddFame struct{}

func (srv *RemoteServe) AddFame(args *Args_AddFame, reply *Reply_AddFame) error {
	return Remote.Serve(mdb.RPC_Remote_AddFame, args, mdb.TRANS_TAG_RPC_Serve_AddFame, func() error {
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(args.Pid, func(db *mdb.Database) {
				module.Player.AddFame(db, int32(args.Num))
			})
		})
		return nil
	})
}

func RemoteAddFame(pid int64, num int64, system int32, callback func(error)) {
	reply := new(Reply_AddFame)
	args := &Args_AddFame{
		Pid:    pid,
		Num:    num,
		System: system,
	}
	serverId, _ := module.GetServerIdWithPlayerId(pid)
	Remote.Call(serverId, mdb.RPC_Remote_AddFame, args, reply, mdb.TRANS_TAG_RPC_Call_AddFame, func(err error) {
		if err != nil {
			log.Errorf("RemoteAddFame %v\n", err)
		}
		if callback != nil {
			callback(err)
		}

	})
}
