package rpc

import (
	"game_server/api/protocol/player_api"
	"game_server/config"
	"game_server/mdb"
	"game_server/module/player_common"
	"game_server/module/player_rpc"
	. "game_server/rpc"
)

type Args_PlayerUpdateRankDatas struct {
	RPCArgTag
	RankType player_api.RankType
	Datas    []*player_common.PlayerRankData
}

type Reply_PlayerUpdateRankDatas struct {
}

func (srv *RemoteServe) PlayerUpdateRankDatas(args *Args_PlayerUpdateRankDatas, reply *Reply_PlayerUpdateRankDatas) error {
	return Remote.Serve(mdb.RPC_Remote_PlayerUpdateRankDatas, args, mdb.TRANS_TAG_RPC_Serve_PlayerUpdateRankDatas, func() error {
		player_rpc.UpdateGlobal(args.RankType, args.Datas)
		return nil
	})
}

func RemotePlayerUpdateRankDatas(rankType player_api.RankType, datas []*player_common.PlayerRankData) {
	args := &Args_PlayerUpdateRankDatas{
		RankType: rankType,
		Datas:    datas,
	}
	Remote.Call(config.ServerCfg.GlobalServerId, mdb.RPC_Remote_PlayerUpdateRankDatas, args, &Reply_PlayerUpdateRankDatas{}, mdb.TRANS_TAG_RPC_Call_PlayerUpdateRankDatas, nil)
}
