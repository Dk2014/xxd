package notify

import (
	"core/net"
	"game_server/api/protocol/notify_api"
)

type NotifyAPI struct {
}

var (
	api = NotifyAPI{}
)

func init() {
	notify_api.SetInHandler(api)
}

func (api NotifyAPI) PlayerKeyChanged(session *net.Session, in *notify_api.PlayerKeyChanged_In) {
}
