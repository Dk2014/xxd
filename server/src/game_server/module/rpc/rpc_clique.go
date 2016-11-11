package rpc

import (
	"game_server/api/protocol/clique_api"
	"game_server/config"
	"game_server/dat/clique_dat"
	"game_server/global"
	"game_server/mdb"
	"game_server/module"
	. "game_server/rpc"
	//"game_server/tlog"
)

type Args_UpdateCliqueKongfuAttrib struct {
	RPCArgTag
	Pid    int64
	Attrib map[int8]int32
}

type Reply_UpdateCliqueKongfuAttrib struct {
}

//游戏服务器作为RPC服务器，本服玩家通过该服务增加邮件
func (mail *RemoteServe) UpdateCliqueKongfuAttrib(args *Args_UpdateCliqueKongfuAttrib, reply *Reply_UpdateCliqueKongfuAttrib) error {
	return Remote.Serve(mdb.RPC_Remote_UpdateCliqueKongfuAttrib, args, mdb.TRANS_TAG_RPC_Serve_UpdateCliqueKongfuAttrib, func() error {
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(args.Pid, func(db *mdb.Database) {
				insert := false
				playerCliqueKongfuAttrib := db.Lookup.PlayerCliqueKongfuAttrib(args.Pid)
				if playerCliqueKongfuAttrib == nil {
					insert = true
					playerCliqueKongfuAttrib = &mdb.PlayerCliqueKongfuAttrib{
						Pid: args.Pid,
					}
				}
				for attribType, value := range args.Attrib {
					switch attribType {
					case clique_dat.KONGFU_ATTRIB_HEALTH:
						playerCliqueKongfuAttrib.Health = value
					case clique_dat.KONGFU_ATTRIB_ATTACK:
						playerCliqueKongfuAttrib.Attack = value
					case clique_dat.KONGFU_ATTRIB_DEFANCE:
						playerCliqueKongfuAttrib.Defence = value
					}
				}
				if insert {
					db.Insert.PlayerCliqueKongfuAttrib(playerCliqueKongfuAttrib)
				} else {
					db.Update.PlayerCliqueKongfuAttrib(playerCliqueKongfuAttrib)
				}
			})
		})
		return nil
	})
}

func RemoteUpdateCliqueKongfuAttrib(pid int64, attrib map[int8]int32) {
	args := &Args_UpdateCliqueKongfuAttrib{
		Pid:    pid,
		Attrib: attrib,
	}
	serverId, _ := module.GetServerIdWithPlayerId(pid)
	Remote.Call(serverId, mdb.RPC_Remote_UpdateCliqueKongfuAttrib, args, &Reply_UpdateCliqueKongfuAttrib{}, mdb.TRANS_TAG_RPC_Call_UpdateCliqueKongfuAttrib, nil)
}

type Args_AddAwardExp struct {
	RPCArgTag
	Pid         int64
	AwardExpNum int32
}

type Reply_AddAwardExp struct {
}

//游戏服务器作为RPC服务器，本服玩家通过该服务增加邮件
func (mail *RemoteServe) AddAwardExp(args *Args_AddAwardExp, reply *Reply_AddAwardExp) error {
	return Remote.Serve(mdb.RPC_Remote_AddAwardExp, args, mdb.TRANS_TAG_RPC_Call_AddAwardExp, func() error {
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(args.Pid, func(db *mdb.Database) {
				if session, ok := module.Player.GetPlayerOnline(args.Pid); ok {
					state := module.State(session)
					module.Role.AddFormRoleExp(state, int64(args.AwardExpNum), 0)
				}
			})
		})
		return nil
	})

}

func RemoteAddAwardExp(pid int64, AwardExp int32) {
	args := &Args_AddAwardExp{
		Pid:         pid,
		AwardExpNum: AwardExp,
	}
	serverId, _ := module.GetServerIdWithPlayerId(pid)
	Remote.Call(serverId, mdb.RPC_Remote_AddAwardExp, args, &Reply_AddAwardExp{}, mdb.TRANS_TAG_RPC_Call_AddAwardExp, nil)

}

type Args_EscortBoatHijackBattleWin struct {
	RPCArgTag
	AttackerPid int64
	BoatId      int64
}

type Reply_EscortBoatHijackBattleWin struct {
}

func (srv *RemoteServe) EscortBoatHijackBattleWin(args *Args_EscortBoatHijackBattleWin, reply *Reply_EscortBoatHijackBattleWin) error {
	return Remote.Serve(mdb.RPC_Remote_EscortBoatHijackBattleWin, args, mdb.TRANS_TAG_RPC_Serve_EscortBoatHijackBattleWin, func() error {
		//劫持战斗胜利
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			module.CliqueEscortRPC.HijackBattleWin(globalDB, args.AttackerPid, args.BoatId)
		})
		return nil
	})
}

func RemoteEscortBoatHijackBattleWin(attackerPid, boatId int64) {
	args := &Args_EscortBoatHijackBattleWin{
		AttackerPid: attackerPid,
		BoatId:      boatId,
	}
	Remote.Call(config.ServerCfg.GlobalServerId, mdb.RPC_Remote_EscortBoatHijackBattleWin, args, &Reply_EscortBoatHijackBattleWin{}, mdb.TRANS_TAG_RPC_Call_EscortBoatHijackBattleWin, nil)
}

type Args_EscortBoatRecoverBattleWin struct {
	RPCArgTag
	AttackerPid int64
	BoatId      int64
}

type Reply_EscortBoatRecoverBattleWin struct {
}

func (srv *RemoteServe) EscortBoatRecoverBattleWin(args *Args_EscortBoatRecoverBattleWin, reply *Reply_EscortBoatRecoverBattleWin) error {
	return Remote.Serve(mdb.RPC_Remote_EscortBoatRecoverBattleWin, args, mdb.TRANS_TAG_RPC_Serve_EscortBoatRecoverBattleWin, func() error {
		//劫持战斗胜利
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			module.CliqueEscortRPC.RecoverBattleWin(globalDB, args.AttackerPid, args.BoatId)
		})
		return nil
	})
}

func RemoteEscortBoatRecoverBattleWin(attackerPid, boatId int64) {
	args := &Args_EscortBoatRecoverBattleWin{
		AttackerPid: attackerPid,
		BoatId:      boatId,
	}
	Remote.Call(config.ServerCfg.GlobalServerId, mdb.RPC_Remote_EscortBoatRecoverBattleWin, args, &Reply_EscortBoatRecoverBattleWin{}, mdb.TRANS_TAG_RPC_Call_EscortBoatRecoverBattleWin, nil)
}

type Args_EscortBoatRecoverBattleLose struct {
	RPCArgTag
	AttackerPid int64
	BoatId      int64
}

type Reply_EscortBoatRecoverBattleLose struct {
}

func (srv *RemoteServe) EscortBoatRecoverBattleLose(args *Args_EscortBoatRecoverBattleLose, reply *Reply_EscortBoatRecoverBattleLose) error {
	return Remote.Serve(mdb.RPC_Remote_EscortBoatRecoverBattleLose, args, mdb.TRANS_TAG_RPC_Serve_EscortBoatRecoverBattleLose, func() error {
		//劫持战斗胜利
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			module.CliqueEscortRPC.RecoverBattleLose(globalDB, args.AttackerPid, args.BoatId)
		})
		return nil
	})
}

func RemoteEscortBoatRecoverBattleLose(attackerPid, boatId int64) {
	args := &Args_EscortBoatRecoverBattleLose{
		AttackerPid: attackerPid,
		BoatId:      boatId,
	}
	Remote.Call(config.ServerCfg.GlobalServerId, mdb.RPC_Remote_EscortBoatRecoverBattleLose, args, &Reply_EscortBoatRecoverBattleLose{}, mdb.TRANS_TAG_RPC_Call_EscortBoatRecoverBattleLose, nil)
}

type Args_UpdatClubHousePlayer struct {
	RPCArgTag
	Pid              int64
	DressedFashionId int16
}

type Reply_UpdatClubHousePlayer struct {
}

func (sev *RemoteServe) UpdateClubHousePlayer(args *Args_UpdatClubHousePlayer, reply *Reply_UpdatClubHousePlayer) error {
	return Remote.Serve(mdb.RPC_Remote_EscortBoatRecoverBattleLose, args, mdb.TRANS_TAG_RPC_Serve_EscortBoatRecoverBattleLose, func() error {
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(args.Pid, func(db *mdb.Database) {
				if session, ok := module.Player.GetPlayerOnline(args.Pid); ok {
					state := module.State(session)
					playerCliqueInfo := state.Database.Lookup.PlayerGlobalCliqueInfo(state.PlayerId)
					if playerCliqueInfo == nil || playerCliqueInfo.CliqueId <= 0 {
						return
					}
					if !state.InCliqueClubhouse {
						return
					}
					playerInfo := global.GetPlayerInfo(args.Pid)
					module.CliqueRPC.BroadcastClubhouse(playerCliqueInfo.CliqueId, &clique_api.NotifyUpdatePlayer_Out{
						PlayerId:          args.Pid,
						FashionId:         args.DressedFashionId,
						InMeditationState: state.IsMeditation,
					})
					playerInfo.FashionId = args.DressedFashionId
					global.UpdatePlayerInfo(args.Pid, playerInfo)
				}
			})
		})
		return nil
	})
}

func RemoteUpdatClubHousePlayer(pid int64, dressedFashionId int16) {
	args := &Args_UpdatClubHousePlayer{Pid: pid, DressedFashionId: dressedFashionId}
	Remote.Call(config.ServerCfg.GlobalServerId, mdb.RPC_Remote_UpdateClubHousePlayer, args, &Reply_UpdatClubHousePlayer{}, mdb.TRANS_TAG_RPC_Call_UpdateClubHousePlayer, nil)
}
