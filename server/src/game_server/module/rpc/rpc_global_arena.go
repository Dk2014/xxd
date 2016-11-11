package rpc

import (
	"core/fail"
	"core/time"
	"game_server/api/protocol/arena_api"
	"game_server/config"
	"game_server/dat/player_dat"
	"game_server/global"
	"game_server/mdb"
	"game_server/module"
	. "game_server/rpc"
	"math/rand"
	goTime "time"
)

//获取玩家排名
type Args_GlobalArenaGetArenaRank struct {
	RPCArgTag
	PlayerId int64
}
type Reply_GlobalArenaGetArenaRank struct {
	PlayerRank int32
}

func (this *RemoteServe) GetArenaRank(args *Args_GlobalArenaGetArenaRank, reply *Reply_GlobalArenaGetArenaRank) error {
	return Remote.Serve(mdb.RPC_Remote_GetArenaRank, args, mdb.TRANS_TAG_RPC_Serve_GetArenaRank, func() error {
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(args.PlayerId, func(db *mdb.Database) {
				reply.PlayerRank = getPlayerArenaRankInGlobalServer(db, args.PlayerId)
			})
		})
		return nil
	})
}

func RemoteGetArenaRank(playerId int64, callback func(*Reply_GlobalArenaGetArenaRank)) {
	args := &Args_GlobalArenaGetArenaRank{PlayerId: playerId}
	reply := &Reply_GlobalArenaGetArenaRank{}

	Remote.Call(config.ServerCfg.GlobalServerId, mdb.RPC_Remote_GetArenaRank, args, reply, mdb.TRANS_TAG_RPC_Call_GetArenaRank, func(err error) {
		fail.When(err != nil, err)
		callback(reply)
	})
}

/*
	获取玩家排行列表
*/
type Args_GlobalArenaEnterArena struct {
	RPCArgTag
	PlayerId int64
}

type Reply_GlobalArenaEnterArena struct {
	PlayerRank   int32
	BoxAwardTime int64
	Ranks        []arena_api.Enter_Out_Ranks
}

func (this *RemoteServe) EnterArena(args *Args_GlobalArenaEnterArena, reply *Reply_GlobalArenaEnterArena) error {
	return Remote.Serve(mdb.RPC_Remote_EnterArena, args, mdb.TRANS_TAG_RPC_Serve_EnterArena, func() error {
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(args.PlayerId, func(db *mdb.Database) {
				reply.PlayerRank = getPlayerArenaRankInGlobalServer(db, args.PlayerId)
				reply.BoxAwardTime = time.GetNextTodayPointHour(player_dat.RESET_ARENA_AWARD_BOX_IN_HOUR, goTime.Now())
				reply.Ranks = module.ArenaRPC.GetPlayerRankWithRank(db, reply.PlayerRank)
			})
		})
		return nil
	})
}

func RemoteEnterArena(playerId int64, callback func(*Reply_GlobalArenaEnterArena)) {
	args := &Args_GlobalArenaEnterArena{PlayerId: playerId}
	reply := &Reply_GlobalArenaEnterArena{}

	Remote.Call(config.ServerCfg.GlobalServerId, mdb.RPC_Remote_EnterArena, args, reply, mdb.TRANS_TAG_RPC_Call_EnterArena, func(err error) {
		fail.When(err != nil, err)
		callback(reply)
	})
}

/*
	获取挑战玩家的信息
*/

type Args_GlobalArenaGetTargetPlayerRank struct {
	RPCArgTag
	PlayerId       int64
	TargetPlayerId int64
}

type Reply_GlobalArenaGetTargetPlayerRank struct {
	PlayerRank       int32
	TargetPlayerRank int32
	TargetPlayerInfo *global.PlayerInfo
}

func (this *RemoteServe) GetTargetPlayerRank(args *Args_GlobalArenaGetTargetPlayerRank, reply *Reply_GlobalArenaGetTargetPlayerRank) error {
	return Remote.Serve(mdb.RPC_Remote_GetTargetPlayerRank, args, mdb.TRANS_TAG_RPC_Serve_GetTargetPlayerRank, func() error {
		reply.PlayerRank = module.ArenaRPC.GetPlayerRank(args.PlayerId)
		reply.TargetPlayerRank = module.ArenaRPC.GetPlayerRank(args.TargetPlayerId)
		reply.TargetPlayerInfo = global.GetPlayerInfo(args.TargetPlayerId)
		return nil
	})
}

func RemoteGetTargetPlayerRank(playerId, targetPlayerId int64, callback func(*Reply_GlobalArenaGetTargetPlayerRank)) {
	args := &Args_GlobalArenaGetTargetPlayerRank{PlayerId: playerId, TargetPlayerId: targetPlayerId}
	reply := &Reply_GlobalArenaGetTargetPlayerRank{}

	Remote.Call(config.ServerCfg.GlobalServerId, mdb.RPC_Remote_GetTargetPlayerRank, args, reply, mdb.TRANS_TAG_RPC_Call_GetTargetPlayerRank, func(err error) {
		fail.When(err != nil, err)
		callback(reply)
	})
}

/*
	设置玩家排名
*/
type Args_GlobalArenaSwapPlayerRank struct {
	RPCArgTag
	PlayerId       int64
	PlayerRank     int32
	TargetPlayerId int64
	TargetRank     int32
}

type Reply_GlobalArenaSwapPlayerRank struct {
	Result bool
}

func (this *RemoteServe) SwapPlayerRank(args *Args_GlobalArenaSwapPlayerRank, reply *Reply_GlobalArenaSwapPlayerRank) error {
	return Remote.Serve(mdb.RPC_Remote_SwapPlayerRank, args, mdb.TRANS_TAG_RPC_Serve_SwapPlayerRank, func() error {
		// 对手排名落后于玩家则不交换排名
		if args.TargetRank > args.PlayerRank {
			return nil
		}

		playerRank := module.ArenaRPC.GetPlayerRank(args.PlayerId)
		targetRank := module.ArenaRPC.GetPlayerRank(args.TargetPlayerId)
		// 排名已经被其他玩家抢走了
		if targetRank != args.TargetRank || playerRank != args.PlayerRank {
			reply.Result = false
		} else {
			reply.Result = true

			mdb.GlobalExecute(func(db *mdb.Database) {
				module.ArenaRPC.UpdatePlayerRank(db, args.PlayerId, args.TargetRank)
				module.ArenaRPC.UpdatePlayerRank(db, args.TargetPlayerId, args.PlayerRank)
			})
		}

		return nil
	})
}

func RemoteSwapPlayerRank(playerId, targetPlayerId int64, playerRank, targetRank int32, callback func(*Reply_GlobalArenaSwapPlayerRank)) {
	args := &Args_GlobalArenaSwapPlayerRank{PlayerId: playerId, TargetPlayerId: targetPlayerId, PlayerRank: playerRank, TargetRank: targetRank}
	reply := &Reply_GlobalArenaSwapPlayerRank{}

	Remote.Call(config.ServerCfg.GlobalServerId, mdb.RPC_Remote_SwapPlayerRank, args, reply, mdb.TRANS_TAG_RPC_Call_SwapPlayerRank, func(err error) {
		fail.When(err != nil, err)
		callback(reply)
	})
}

/*
	为玩家添加排名
*/
type Args_GlobalArenaAddPlayerRank struct {
	RPCArgTag
	PlayerId int64
}

type Reply_GlobalArenaAddPlayerRank struct {
}

// 互动服上的玩家排名
func getPlayerArenaRankInGlobalServer(db *mdb.Database, pid int64) (rank int32) {
	rank = module.ArenaRPC.GetPlayerRank(pid)
	if rank == 0 {
		rank = module.ArenaRPC.AddPlayerRank(db, pid)
	}

	return
}

/*
	从比武场排名前后随机出一名玩家用做云海御剑拜访用
*/
type Args_GetTargetFighterForDrivingSwordVisting struct {
	RPCArgTag
	PlayerId   int64
	PlayerUsed map[int64]int64
}

type Reply_GetTargetFighterForDrivingSwordVisting struct {
	TargetPid        int64
	TargetPlayerInfo *global.PlayerInfo
	// Fighters         [][]byte //*battle.Fighter
	// BattlePlayerInfo *battle.BattlePlayerInfo
}

func (this *RemoteServe) GetTargetFighterForDrivingSwordVisting(args *Args_GetTargetFighterForDrivingSwordVisting, reply *Reply_GetTargetFighterForDrivingSwordVisting) error {
	return Remote.Serve(mdb.RPC_Remote_GetTargetFighterForDrivingSwordVisting, args, mdb.TRANS_TAG_RPC_Serve_GetTargetFighterForDrivingSwordVisting, func() error {
		//fighters := make([]*battle.Fighter, module.ALL_FIGHTER_POS_NUM)
		var targetPid int64
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(args.PlayerId, func(db *mdb.Database) {
				// 查找可随机得玩家数组
				players := make([]int64, 0)
				count := 10
				totalTimes := 2 * 2 * count //防止死循环，最大查找次数
				side := true
				playerRank := module.ArenaRPC.GetPlayerRank(args.PlayerId)
				var before int32 = playerRank //排名往前找的索引
				var after int32 = playerRank  //排名往后找的索引
				for count > 0 && totalTimes > 0 {
					if side {
						// 往前找
						before = before - 1
						if target := db.Lookup.GlobalArenaRank(before); target != nil {
							if _, ok := args.PlayerUsed[target.Pid]; !ok {
								players = append(players, target.Pid)
								count = count - 1
							}
						}
					} else {
						// 往后找
						after = after + 1
						if target := db.Lookup.GlobalArenaRank(after); target != nil {
							if _, ok := args.PlayerUsed[target.Pid]; !ok {
								players = append(players, target.Pid)
								count = count - 1
							}
						}
					}
					side = !side
					totalTimes = totalTimes - 1
				}
				// 接下来从找到的玩家数组中随机一个
				total := len(players)

				if total > 0 {
					targetIndex := rand.Intn(total)
					targetPid = players[targetIndex]
				} else {
					// 极端情况 找不到玩家 便打玩家自己
					targetPid = args.PlayerId
				}

			})
		})
		reply.TargetPid = targetPid
		reply.TargetPlayerInfo = global.GetPlayerInfo(targetPid)
		return nil
	})
	return nil
}

func RemoteGetTargetFighterForDrivingSwordVisting(playerId int64, playerUsed map[int64]int64, callback func(*Reply_GetTargetFighterForDrivingSwordVisting)) {
	args := &Args_GetTargetFighterForDrivingSwordVisting{PlayerId: playerId, PlayerUsed: playerUsed}
	reply := &Reply_GetTargetFighterForDrivingSwordVisting{}

	Remote.Call(config.ServerCfg.GlobalServerId, mdb.RPC_Remote_GetTargetFighterForDrivingSwordVisting, args, reply, mdb.TRANS_TAG_RPC_Call_GetTargetFighterForDrivingSwordVisting, func(err error) {
		fail.When(err != nil, err)
		callback(reply)
	})
}
