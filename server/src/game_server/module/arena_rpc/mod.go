package arena_rpc

import (
	"core/net"
	"game_server/api/protocol/arena_api"
	"game_server/global"
	"game_server/mdb"
	"game_server/module"
	"game_server/module/rpc"
)

func init() {
	module.ArenaRPC = ArenaRPCMod{}
}

type ArenaRPCMod struct {
}

func (mod ArenaRPCMod) SendTopRank(session *net.Session) {
	state := module.State(session)
	rsp := &arena_api.GetTopRank_Out{}

	for i := int32(1); i <= 50; i++ {
		rank := state.Database.Lookup.GlobalArenaRank(i)
		if rank == nil {
			break
		}

		playerInfo := global.GetPlayerInfo(rank.Pid)

		// 前三名给详细信息
		if i < 4 {
			rsp.Top3 = append(rsp.Top3, arena_api.GetTopRank_Out_Top3{
				Openid: playerInfo.Openid,
				RoleId: playerInfo.RoleId,
			})
		}

		trend := int8(playerInfo.ArenaTrendWin)
		if playerInfo.ArenaTrendWin > 0 {
			trend = 1
		}

		rsp.Ranks = append(rsp.Ranks, arena_api.GetTopRank_Out_Ranks{
			Pid:   rank.Pid,
			Rank:  rank.Rank,
			Nick:  playerInfo.PlayerNick,
			Level: playerInfo.RoleLevel,
			Trend: trend,
		})
	}

	session.Send(rsp)
}

func (mod ArenaRPCMod) GetPlayerRank(pid int64) int32 {
	return rankTable.getPlayerRank(pid)
}

func (mod ArenaRPCMod) AddPlayerRank(db *mdb.Database, pid int64) int32 {
	return addPlayerRank(db, pid)
}

func (mod ArenaRPCMod) UpdatePlayerRank(db *mdb.Database, pid int64, rank int32) {
	updatePlayerRank(db, pid, rank)
}

func (mod ArenaRPCMod) GetPlayerRankWithRank(db *mdb.Database, rank int32) []arena_api.Enter_Out_Ranks {
	return getPlayerRankWithRank(db, rank)
}

func (mod ArenaRPCMod) GetAwardBox(session *net.Session, num int8) {
	state := module.State(session)
	rank := getPlayerRank(state.PlayerId)

	rpc.RemoteAwardArenaBox(state.PlayerId, num, state.MoneyState, rank, func() {
		session.Send(&arena_api.AwardBox_Out{
			Result: true,
		})
	})
}
