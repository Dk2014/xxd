package rpc

import (
	"core/fail"
	"game_server/api/protocol/town_api"
	"game_server/config"
	"game_server/dat/meditation_dat"
	"game_server/mdb"
	"game_server/module"
	. "game_server/rpc"
)

type Args_GetClubHouseMeditation struct {
	RPCArgTag
	PlayerId int64
	Mode     int8 //模式1、查询，2、开始打坐，3、结束打坐
}

type Reply_GetClubHouseMeditation struct {
	IsMeditation      bool //是否在集会所打坐
	InCliqueClubhouse bool //是否在集会所
}

func (this *RemoteServe) GetClubHouseMeditation(args *Args_GetClubHouseMeditation, reply *Reply_GetClubHouseMeditation) error {
	return Remote.Serve(mdb.RPC_Remote_GetClubHouseMeditation, args, mdb.TRANS_TAG_RPC_Serve_GetClubHouseMeditation, func() error {
		reply.IsMeditation = false
		reply.InCliqueClubhouse = false
		if session, ok := module.Player.GetPlayerOnline(args.PlayerId); ok {
			state := module.State(session)
			reply.InCliqueClubhouse = state.InCliqueClubhouse
			db := state.Database
			playerCliqueInfo := db.Lookup.PlayerGlobalCliqueInfo(args.PlayerId)
			if playerCliqueInfo == nil || playerCliqueInfo.CliqueId <= 0 {
				return nil
			}
			if args.Mode == meditation_dat.CLUBHOUSE_MEDITATION_START {
				state.IsMeditation = true
				module.CliqueRPC.BroadcastClubhouse(playerCliqueInfo.CliqueId, &town_api.UpdateTownPlayerMeditationState_Out{
					PlayerId:        args.PlayerId,
					MeditationState: true,
				})
			} else if args.Mode == meditation_dat.CLUBHOUSE_MEDITATION_STOP {
				state.IsMeditation = false
				module.CliqueRPC.BroadcastClubhouse(playerCliqueInfo.CliqueId, &town_api.UpdateTownPlayerMeditationState_Out{
					PlayerId:        args.PlayerId,
					MeditationState: false,
				})
			}
			reply.IsMeditation = state.IsMeditation
		}
		return nil
	})
}

func RemoteGetClubHouseMeditation(playerId int64, mode int8) {
	args := &Args_GetClubHouseMeditation{PlayerId: playerId, Mode: mode}
	reply := &Reply_GetClubHouseMeditation{}
	Remote.Call(config.ServerCfg.GlobalServerId, mdb.RPC_Remote_GetClubHouseMeditation, args, reply, mdb.TRANS_TAG_RPC_Call_GetClubHouseMeditation, func(err error) {
		fail.When(err != nil, err)
	})
}
