package totem

import (
	"core/net"
	"game_server/api/protocol/notify_api"
	"game_server/api/protocol/totem_api"
	"game_server/module"
)

func init() {
	totem_api.SetInHandler(TotemAPI{})
}

type TotemAPI struct{}

// 获取玩家所有阵印信息
func (this TotemAPI) Info(session *net.Session, in *totem_api.Info_In) {
	out := &totem_api.Info_Out{}
	state := module.State(session)
	getTotemInfo(state.Database, out)
	session.Send(out)
}

// 装备阵印
func (this TotemAPI) Equip(session *net.Session, in *totem_api.Equip_In) {
	//out := &totem_api.Equip_Out{}
	state := module.State(session)
	equipTotem(state.Database, in.TotemId, in.EquipPos)
	//session.Send(out)
}

// 卸载一个阵印
func (this TotemAPI) Unequip(session *net.Session, in *totem_api.Unequip_In) {
	out := &totem_api.Unequip_Out{}
	state := module.State(session)
	unequipTotem(state.Database, in.EquipPos)
	session.Send(out)
}

//装备栏和背包互相拖
//已装备阵印和未装备阵印交换位置
func (this TotemAPI) Swap(session *net.Session, in *totem_api.Swap_In) {
	out := &totem_api.Swap_Out{}
	state := module.State(session)
	swapTotem(state.Database, in.EquipedId, in.InbagId)
	session.Send(out)
}

//装备栏内拖动
// 装备的阵印位置改变
func (this TotemAPI) EquipPosChange(session *net.Session, in *totem_api.EquipPosChange_In) {
	out := &totem_api.EquipPosChange_Out{}
	state := module.State(session)
	totemEquipPosChange(state.Database, int8(in.FromPos), int8(in.ToPos))
	session.Send(out)
}

//铭刻
func (this TotemAPI) Upgrade(session *net.Session, in *totem_api.Upgrade_In) {
	out := &totem_api.Upgrade_Out{}
	state := module.State(session)

	rockNum, jadeNum, ok := upgradeTotem(state.Database, in.TargetId, out)
	session.Send(&notify_api.NotifyRuneChange_Out{
		RockRuneNum: rockNum,
		JadeRuneNum: jadeNum,
	})
	out.Ok = ok
	session.Send(out)
}

// 阵印召唤
func (this TotemAPI) CallTotem(session *net.Session, in *totem_api.CallTotem_In) {
	//out := &totem_api.CallTotem_Out{}
	callTotem(session, in.CallType)
	//session.Send(out)
}

// 阵印分解
func (this TotemAPI) Decompose(session *net.Session, in *totem_api.Decompose_In) {
	out := &totem_api.Decompose_In{}
	state := module.State(session)
	rockNum, jadeNum := decomposeTotem(state.Database, in.TotemId)
	session.Send(&notify_api.NotifyRuneChange_Out{
		RockRuneNum: rockNum,
		JadeRuneNum: jadeNum,
	})
	session.Send(out)
}

func (this TotemAPI) IsBagFull(session *net.Session, in *totem_api.IsBagFull_In) {
	state := module.State(session)
	session.Send(&totem_api.IsBagFull_Out{
		Full: isTotemBagFull(state.Database),
	})
}
