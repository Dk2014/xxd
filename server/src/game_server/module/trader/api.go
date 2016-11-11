package trader

import (
	"core/net"
	"game_server/api/protocol/trader_api"
	"game_server/xdlog"
)

func init() {
	trader_api.SetInHandler(TraderAPI{})
}

type TraderAPI struct{}

func (api TraderAPI) Info(session *net.Session, in *trader_api.Info_In) {
	out := new(trader_api.Info_Out)
	info(session, out)
	session.Send(out)
}

func (api TraderAPI) StoreState(session *net.Session, in *trader_api.StoreState_In) {
	out := new(trader_api.StoreState_Out)
	storeSate(session, in.TraderId, out)
	session.Send(out)
}

func (api TraderAPI) Buy(session *net.Session, in *trader_api.Buy_In) {
	out := new(trader_api.Buy_Out)
	buy(session, in.GridId, xdlog.ET_TRADER, out)
	session.Send(out)
}

func (api TraderAPI) Refresh(session *net.Session, in *trader_api.Refresh_In) {
	out := new(trader_api.Refresh_Out)
	refresh(session, in.TraderId, out)
	session.Send(out)
}
