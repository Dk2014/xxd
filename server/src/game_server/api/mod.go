package api

import (
	"core/net"
	"game_server/module"
)

func init() {
	module.API = APIMod{}
}

type APIMod struct {
}

func (this APIMod) Broadcast(sessions net.SessionList, response net.Response) {
	g_ApiServer.Broadcast(sessions, response)
}
