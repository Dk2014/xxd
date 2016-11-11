package vip

import (
	"core/net"

	"game_server/api/protocol/vip_api"
	"game_server/module"
)

type VIP_API struct{}

func init() {
	vip_api.SetInHandler(VIP_API{})
}

func (this VIP_API) Info(session *net.Session, in *vip_api.Info_In) {
	state := module.State(session)
	playerVIPInfo := state.Database.Lookup.PlayerVip(in.PlayerId)

	out := &vip_api.Info_Out{}
	out.Level = playerVIPInfo.Level
	out.Ingot = playerVIPInfo.Ingot + playerVIPInfo.PresentExp
	out.CardId = []byte(playerVIPInfo.CardId)
	session.Send(out)
}

func (this VIP_API) GetTotal(session *net.Session, in *vip_api.GetTotal_In) {
	out := &vip_api.GetTotal_Out{}
	out.Total = getVIPTotal()
	session.Send(out)
}

func (this VIP_API) VipLevelTotal(session *net.Session, in *vip_api.VipLevelTotal_In) {
	out := &vip_api.VipLevelTotal_Out{}
	vip_level_arr := getLevelVipTotal()
	for level, total := range vip_level_arr {
		if level == 0 {
			continue
		}
		out.VipLevelArr = append(out.VipLevelArr, vip_api.VipLevelTotal_Out_VipLevelArr{
			VipLevel: int16(level),
			Total:    total,
		})
	}
	session.Send(out)
}

func (this VIP_API) BuyTimes(session *net.Session, in *vip_api.BuyTimes_In) {
	state := module.State(session)
	out := &vip_api.BuyTimes_Out{}
	result := buytimes(state, in.Buytype)
	out.Result = result
	session.Send(out)
}
