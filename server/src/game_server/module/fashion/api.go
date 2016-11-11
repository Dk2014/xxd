package fashion

import (
	"core/net"
	"game_server/api/protocol/fashion_api"
)

type FashionAPI struct{}

func init() {
	fashion_api.SetInHandler(FashionAPI{})
}

func (this FashionAPI) FashionInfo(session *net.Session, in *fashion_api.FashionInfo_In) {
	out := &fashion_api.FashionInfo_Out{}
	fashionInfo(session, out)
	session.Send(out)
}

func (this FashionAPI) DressFashion(session *net.Session, in *fashion_api.DressFashion_In) {
	out := &fashion_api.DressFashion_Out{}
	out.DressCdTime = dressFashion(session, in)
	session.Send(out)
}
