package clique_escort_rpc

import (
	"core/net"
	"game_server/api/protocol/clique_escort_api"
	"game_server/module"
)

type CliqueEscortAPI struct{}

func init() {
	clique_escort_api.SetInHandler(CliqueEscortAPI{})
}

//玩家基本运镖信息
func (this CliqueEscortAPI) EscortInfo(session *net.Session, in *clique_escort_api.EscortInfo_In) {
	escortInfo(session)
}

//使用元宝获得宝船
func (this CliqueEscortAPI) GetIngotBoat(session *net.Session, in *clique_escort_api.GetIngotBoat_In) {
	getIngotBoat(session)
}

//开始运镖
func (this CliqueEscortAPI) StartEscort(session *net.Session, in *clique_escort_api.StartEscort_In) {
	startEscort(session)
}

//劫持镖船
func (this CliqueEscortAPI) HijackBoat(session *net.Session, in *clique_escort_api.HijackBoat_In) {
	hijackBoat(session, in.BoatId)
}

//夺回标船
func (this CliqueEscortAPI) RecoverBoat(session *net.Session, in *clique_escort_api.RecoverBoat_In) {
	recoverBoat(session, in.BoatId)
}

//标船列表
func (this CliqueEscortAPI) ListBoats(session *net.Session, in *clique_escort_api.ListBoats_In) {
	listBoats(session)
}

//获取随机镖船
func (this CliqueEscortAPI) GetRandomBoat(session *net.Session, in *clique_escort_api.GetRandomBoat_In) {
	getRandomBoat(session)
}

//领取劫镖奖励
func (this CliqueEscortAPI) TakeHijackAward(session *net.Session, in *clique_escort_api.TakeHijackAward_In) {
	takeHijackAward(session)
}

//领取送奖励
func (this CliqueEscortAPI) TakeEscortAward(session *net.Session, in *clique_escort_api.TakeEscortAward_In) {
	takeEscortAward(session)
}

func (this CliqueEscortAPI) GetCliqueBoatMessages(session *net.Session, in *clique_escort_api.GetCliqueBoatMessages_In) {
	getCliqueBoatMessages(session)
}

func (this CliqueEscortAPI) ReadCliqueBoatMessage(session *net.Session, in *clique_escort_api.ReadCliqueBoatMessage_In) {
	state := module.State(session)
	msg := state.Database.Lookup.PlayerGlobalCliqueEscortMessage(in.Id)
	if msg != nil {
		state.Database.Delete.PlayerGlobalCliqueEscortMessage(msg)
	}
}
