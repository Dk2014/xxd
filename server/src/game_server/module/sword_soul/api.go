package sword_soul

import (
	"core/net"
	"game_server/api/protocol/sword_soul_api"
)

func init() {
	sword_soul_api.SetInHandler(SwordSoulAPI{})
}

type SwordSoulAPI struct {
}

func (this SwordSoulAPI) Info(session *net.Session, in *sword_soul_api.Info_In) {
	out := &sword_soul_api.Info_Out{}
	Info(session, out)
	session.Send(out)
}

func (this SwordSoulAPI) Draw(session *net.Session, in *sword_soul_api.Draw_In) {
	out := &sword_soul_api.Draw_Out{}
	Draw(session, in, out)
	session.Send(out)
}

func (this SwordSoulAPI) Upgrade(session *net.Session, in *sword_soul_api.Upgrade_In) {
	out := &sword_soul_api.Upgrade_Out{}
	Upgrade(session, in, out)
	session.Send(out)
}

func (this SwordSoulAPI) Exchange(session *net.Session, in *sword_soul_api.Exchange_In) {
	out := &sword_soul_api.Exchange_Out{}
	Exchange(session, in, out)
	session.Send(out)
}

func (this SwordSoulAPI) Equip(session *net.Session, in *sword_soul_api.Equip_In) {
	Equip(session, in)
	session.Send(&sword_soul_api.Equip_Out{})
}

func (this SwordSoulAPI) Unequip(session *net.Session, in *sword_soul_api.Unequip_In) {
	Unequip(session, in)
	session.Send(&sword_soul_api.Unequip_Out{})
}

func (this SwordSoulAPI) EquipPosChange(session *net.Session, in *sword_soul_api.EquipPosChange_In) {
	EquipPosChange(session, in)
	session.Send(&sword_soul_api.EquipPosChange_Out{})
}

//func (this SwordSoulAPI) BagPosChange(session *net.Session, in *sword_soul_api.BagPosChange_In) {
//	BagPosChange(session, in)
//	session.Send(&sword_soul_api.BagPosChange_Out{})
//}

func (this SwordSoulAPI) Swap(session *net.Session, in *sword_soul_api.Swap_In) {
	out := &sword_soul_api.Swap_Out{}
	Swap(session, in)
	session.Send(out)
}

func (this SwordSoulAPI) IsBagFull(session *net.Session, in *sword_soul_api.IsBagFull_In) {
	out := &sword_soul_api.IsBagFull_Out{}
	IsBagFull(session, out)
	session.Send(out)
}

func (this SwordSoulAPI) EmptyPosNum(session *net.Session, in *sword_soul_api.EmptyPosNum_In) {
	out := &sword_soul_api.EmptyPosNum_Out{}
	EmptyPosNum(session, out)
	session.Send(out)
}
