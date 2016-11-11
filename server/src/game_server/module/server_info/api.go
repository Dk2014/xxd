package server_info

import (
	"core/net"
	"game_server/api/protocol/server_info_api"
	"game_server/config"
	"game_server/module"
	"game_server/module/event"
	"game_server/module/rpc"
)

func init() {
	server_info_api.SetInHandler(ServerInfoAPI{})
}

type ServerInfoAPI struct {
}

func (this ServerInfoAPI) GetVersion(session *net.Session, in *server_info_api.GetVersion_In) {
	session.Send(&server_info_api.GetVersion_Out{
		Version: version,
	})
}

func (this ServerInfoAPI) GetApiCount(session *net.Session, in *server_info_api.GetApiCount_In) {
	state := module.State(session)
	session.Send(&server_info_api.GetApiCount_Out{
		CountIn:  state.RspQ.ReqCounter,
		CountOut: state.RspQ.RspCounter,
	})
}

func (this ServerInfoAPI) SearchPlayerRole(session *net.Session, in *server_info_api.SearchPlayerRole_In) {
	out := &server_info_api.SearchPlayerRole_Out{}
	_, out.Result = module.Player.GetPlayerByUsername(string(in.Openid))
	session.Send(out)
}

func (this ServerInfoAPI) UpdateAccessToken(session *net.Session, in *server_info_api.UpdateAccessToken_In) {
	state := module.State(session)
	state.MoneyState.Openkey = string(in.Token)
	state.MoneyState.Pfkey = string(in.Pfkey)
}

func (this ServerInfoAPI) UpdateEventData(session *net.Session, in *server_info_api.UpdateEventData_In) {
	if in.Version < module.Event.GetVersion() {
		session.Send(&server_info_api.UpdateEventData_Out{
			Json: event.EventConfigRawData,
		})
	}
}

func (this ServerInfoAPI) TssData(session *net.Session, in *server_info_api.TssData_In) {
	state := module.State(session)
	if state.PlayerId > 0 {
		player := state.Database.Lookup.Player(state.PlayerId)
		rpc.RemoteTssRecvData(player.User, player.Id, int(config.ServerCfg.XGSdk.PlatformType), in.Data)
	}
}
