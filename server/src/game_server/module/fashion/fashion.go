package fashion

import (
	"core/fail"
	"core/net"
	"core/time"
	"game_server/api/protocol/fashion_api"
	"game_server/api/protocol/town_api"
	"game_server/dat/fashion_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/module/rpc"
	"game_server/tlog"
)

func fashionInfo(session *net.Session, out *fashion_api.FashionInfo_Out) {
	state := module.State(session)
	playerFashionState := state.Database.Lookup.PlayerFashionState(state.PlayerId)
	out.DressCdTime = playerFashionState.UpdateTime + fashion_dat.FASHION_OPERA_CD
	out.DressedFashionId = playerFashionState.DressedFashionId
	state.Database.Select.PlayerFashion(func(row *mdb.PlayerFashionRow) {
		out.Fashions = append(out.Fashions, fashion_api.FashionInfo_Out_Fashions{
			FashionId:  row.FashionId(),
			ExpireTime: row.ExpireTime(),
		})
	})
}

func dressFashion(session *net.Session, in *fashion_api.DressFashion_In) int64 {
	state := module.State(session)
	playerFashionState := state.Database.Lookup.PlayerFashionState(state.PlayerId)
	now := time.GetNowTime()

	fail.When(now-playerFashionState.UpdateTime < fashion_dat.FASHION_OPERA_CD, "换装冷却中")

	if in.FashionId > 0 {
		var playerFashion *mdb.PlayerFashion
		state.Database.Select.PlayerFashion(func(row *mdb.PlayerFashionRow) {
			if row.FashionId() == in.FashionId {
				playerFashion = row.GoObject()
				row.Break()
			}
		})
		fail.When(playerFashion == nil, "玩家目标时装未获得")
	}
	playerFashionState.DressedFashionId = in.FashionId
	playerFashionState.UpdateTime = now
	state.Database.Update.PlayerFashionState(playerFashionState)
	if state.TownChannel != nil {
		module.API.Broadcast(state.TownChannel, &town_api.UpdateTownPlayer_Out{
			PlayerId:  state.PlayerId,
			FashionId: playerFashionState.DressedFashionId,
		})
	}
	rpc.RemoteUpdatePlayerFashion(state.PlayerId, int64(playerFashionState.DressedFashionId))

	tlog.PlayerSystemModuleFlowLog(state.Database, tlog.SMT_FASHION)
	if in.InClubhouse {
		rpc.RemoteUpdatClubHousePlayer(state.PlayerId, in.FashionId)
	}
	return now + fashion_dat.FASHION_OPERA_CD
}
