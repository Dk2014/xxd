package event

import (
	"core/net"
	"game_server/api/protocol/event_api"
	"game_server/global"
)

func getGroupBuyInfo(session *net.Session) {
	count := global.GetGroupBuyCount()
	out := &event_api.InfoGroupBuy_Out{}
	out.Count = count
	session.Send(out)
}
