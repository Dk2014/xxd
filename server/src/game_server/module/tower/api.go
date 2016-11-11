package tower

import (
	"core/net"
	"game_server/api/protocol/tower_api"
	"game_server/module"
)

func init() {
	tower_api.SetInHandler(TowerAPI{})
}

type TowerAPI struct{}

func (api TowerAPI) GetInfo(session *net.Session, in *tower_api.GetInfo_In) {
	state := module.State(session)
	session.Send(&tower_api.GetInfo_Out{
		FloorNum: getPlayerTowerInfo(state),
		Friends:  getPlayerFriends(state),
	})
}

func (api TowerAPI) UseLadder(session *net.Session, in *tower_api.UseLadder_In) {
	session.Send(&tower_api.UseLadder_Out{
		FloorNum: useLadder(module.State(session)),
	})
}
