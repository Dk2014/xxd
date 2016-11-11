package module

import (
	"core/net"
	"game_server/battle"
)

type BasicBattle struct{}

func (b BasicBattle) GetBattle() *battle.BattleState {
	panic("not implement")
}

func (b BasicBattle) LeaveBattle(session *net.Session) {
	panic("not implement")
}

func (b BasicBattle) NextRound(_ *NextRoundParams) {
	panic("not implement")
}

func (b BasicBattle) Escape(session *net.Session) {
	panic("not implement")
}

func (b BasicBattle) Relive(session *net.Session) {
	panic("not implement")
}

func (b BasicBattle) CallBattlePet(session *net.Session, petId int32, petLevel int16, petSkillLv int16) bool {
	panic("not implement")
}

func (b BasicBattle) CallNewEnemy(enemy battle.CallInfo) *battle.Fighter {
	panic("not implement")
}

func (b BasicBattle) UseBuddySkill(session *net.Session, posIndex int8, skillIndex int8) {
	panic("not implement")
}

func (b BasicBattle) SetSkill(session *net.Session, posIdx int8, skillIdx int8) {
	panic("not implement")
}

func (b BasicBattle) InitRound(session *net.Session) {
	panic("not implement")
	/*
		1. send or broadcast totem result
		2. maintin online state for every player. notify online player it is his turn or send prepare_ready to battleState

		for single player:
		battleState := b.GetBattle()
		result, status, nowRound, firstPid := battleState.RequireTotem()
		session.Send(result)
		session.Send(firstPid)

		for multi player battle:
		battleState := b.GetBattle()
		result, status, nowRound, firstPid := battleState.RequireTotem()
		if len(result) > 0 {
			self.battleChannel.Broadcast(result)
		}
		if isOnline(firstPid) {
			self.battleChannel.Broadcast(firstPid)
		} else {
			self.PrepareReady(firstPid)
		}

	*/
}

func (b BasicBattle) UseGhost(session *net.Session, isAttacker bool, posIdx int8) {
	panic("not implement")
}

func (b BasicBattle) UseItem(session *net.Session, isAttacker bool, posIdx int8, itemId int16) {
	panic("not implement")
}

func (b BasicBattle) SetAuto(session *net.Session) {
	panic("not implement")
}

func (b BasicBattle) CancelAuto(session *net.Session) {
	panic("not implement")
}

func (b BasicBattle) PrepareReady(session *net.Session, isAuto bool) {
	panic("not implement")
}
