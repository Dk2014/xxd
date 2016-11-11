package awaken

import (
	"core/net"
	"game_server/api/protocol/awaken_api"
	"game_server/module"
)

func init() {
	awaken_api.SetInHandler(AwakenAPI{})
}

type AwakenAPI struct {
}

// 觉醒信息
func (api AwakenAPI) AwakenInfo(session *net.Session, in *awaken_api.AwakenInfo_In) {
	state := module.State(session)
	role_id := in.RoleId
	out := awakenInfo(state, role_id)

	session.Send(out)
}

// 觉醒升级
func (api AwakenAPI) LevelupAttr(session *net.Session, in *awaken_api.LevelupAttr_In) {
	role_id := in.RoleId
	attr_impl := in.AttrImpl
	succeed := levelupAttr(session, role_id, attr_impl)

	if succeed {
		session.Send(&awaken_api.LevelupAttr_Out{
			RoleId:   role_id,
			AttrImpl: attr_impl,
		})
	} else {
		session.Send(&awaken_api.LevelupAttr_Out{
			RoleId:   -1,
			AttrImpl: -1,
		})
	}
}
