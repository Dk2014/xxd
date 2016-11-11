package arena

import (
	"core/net"
	"game_server/api/protocol/arena_api"
	"game_server/module"
)

type ArenaAPI struct {
}

func init() {
	arena_api.SetInHandler(ArenaAPI{})
}

func (api ArenaAPI) Enter(session *net.Session, in *arena_api.Enter_In) {
	enter(session)
}

func (api ArenaAPI) GetRecords(session *net.Session, in *arena_api.GetRecords_In) {
	sendRecords(session)
}

func (api ArenaAPI) StartBattle(session *net.Session, in *arena_api.StartBattle_In) {
	startBattle(session, in.PlayerId, in.PlayerRank)
}

func (api ArenaAPI) CleanFailedCdTime(session *net.Session, in *arena_api.CleanFailedCdTime_In) {
	cleanCDTime(session)
}

func (api ArenaAPI) GetTopRank(session *net.Session, in *arena_api.GetTopRank_In) {
	// 互动服务器处理
	module.ArenaRPC.SendTopRank(session)
}

func (api ArenaAPI) AwardBox(session *net.Session, in *arena_api.AwardBox_In) {
	// 互动服务器处理
	module.ArenaRPC.GetAwardBox(session, in.Num)
}
