package ghost

import (
	"core/net"
	"game_server/api/protocol/ghost_api"
	"game_server/module"
)

func init() {
	ghost_api.SetInHandler(GhostAPI{})
}

type GhostAPI struct {
}

// 获取玩家所有魂侍
func (this GhostAPI) Info(session *net.Session, in *ghost_api.Info_In) {
	out := &ghost_api.Info_Out{}
	Info(session, out)
	session.Send(out)
}

// 装备魂侍
func (this GhostAPI) Equip(session *net.Session, in *ghost_api.Equip_In) {
	out := &ghost_api.Equip_Out{}
	Equip(session, in)
	session.Send(out)
}

// 卸载一个魂侍
func (this GhostAPI) Unequip(session *net.Session, in *ghost_api.Unequip_In) {
	out := &ghost_api.Unequip_Out{}
	Unequip(session, in)
	session.Send(out)
}

// 装备魂侍和魂侍替换位置
func (this GhostAPI) Swap(session *net.Session, in *ghost_api.Swap_In) {
	out := &ghost_api.Swap_Out{}
	Swap(session, in)
	session.Send(out)
}

// 装备的魂侍位置改变
func (this GhostAPI) EquipPosChange(session *net.Session, in *ghost_api.EquipPosChange_In) {
	out := &ghost_api.EquipPosChange_Out{}
	EquipPosChange(session, in)
	session.Send(out)
}

// 魂侍培养
func (this GhostAPI) Train(session *net.Session, in *ghost_api.Train_In) {
	out := &ghost_api.Train_Out{}
	Train(session, in, out)
	session.Send(out)
}

// 魂侍升星
func (this GhostAPI) Upgrade(session *net.Session, in *ghost_api.Upgrade_In) {
	out := &ghost_api.Upgrade_Out{}
	Upgrade(session, in)
	session.Send(out)
}

// 合成
func (this GhostAPI) Compose(session *net.Session, in *ghost_api.Compose_In) {
	out := &ghost_api.Compose_Out{}
	Compose(session, in, out)
	session.Send(out)
}

//训练技能
func (this GhostAPI) TrainSkill(session *net.Session, in *ghost_api.TrainSkill_In) {
	state := module.State(session)
	trainGhostSkill(state, in.Id)
	out := &ghost_api.TrainSkill_Out{}
	session.Send(out)
}

//魂侍洗点
func (this GhostAPI) FlushAttr(session *net.Session, in *ghost_api.FlushAttr_In) {
	state := module.State(session)
	flush_time := flushGhostAttr(state, in.Id)
	session.Send(&ghost_api.FlushAttr_Out{
		FlushTime: flush_time,
	})
}

func (this GhostAPI) RelationGhost(session *net.Session, in *ghost_api.RelationGhost_In) {
	state := module.State(session)
	out := &ghost_api.RelationGhost_Out{}
	relation(state, in.Master, in.Slave)
	session.Send(out)
}

func (this GhostAPI) Baptize(session *net.Session, in *ghost_api.Baptize_In) {
	out := &ghost_api.Baptize_Out{}
	Baptize(session, in, out)
	session.Send(out)
}
