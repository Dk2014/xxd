package battle

import (
	//"core/fail"
	"core/net"
	"game_server/battle"
	"game_server/module"
)

/*
Battle_PVE 用于完全自动战斗，不允许有任何操作
*/
type Battle_PVE struct {
	module.BasicBattle
	battleState *battle.BattleState
	notAward    bool // 对战斗胜利和失败不作处理

	battleStatus int // 战场状态
	//boatId       int64
}

func (this *Battle_PVE) GetBattle() *battle.BattleState {
	return this.battleState
}

func (this *Battle_PVE) InitRound(session *net.Session) {
	this.battleState.AutoNextRound_v2(func(results []*battle.FightResult, status int, nowRound int) (notEnd bool) {
		switch this.battleState.BatType {
		case battle.BT_ARENA:
			if status == battle.NOT_END && nowRound+1 >= 20 {
				status = battle.DEF_WIN
			}
		}
		this.battleStatus = status
		rsp := GetNextRoundResponse(results, status, nowRound, this.battleState, nil, nil)
		session.Send(rsp)
		switch status {
		case battle.ATK_WIN:
			this.DoWin(session)
			return false
		case battle.DEF_WIN:
			this.DoLose(session)
			return false
		}
		return true
	})
}

func (this *Battle_PVE) DoWin(session *net.Session) {
	//state := module.State(session)
	switch this.battleState.BatType {
	case battle.BT_ARENA:
		module.Arena.BattleDoWin(session)
		//case battle.BT_DRIVING_SWORD_BF_LEVEL:
		//module.DrivingSword.VisitingAward(session)
		// case battle.BT_HIJACK_BOAT:
		// 	//rpc 通知互动服 劫持战斗胜利
		// 	rpc.RemoteEscortBoatHijackBattleWin(state.PlayerId, this.boatId)
		// case battle.BT_RECOVER_BOAT:
		// 	//rpc 通知互动服 夺回战斗胜利
		// 	rpc.RemoteEscortBoatRecoverBattleWin(state.PlayerId, this.boatId)
	}

	this.battleState = nil
	module.State(session).Battle = nil
}

func (this *Battle_PVE) DoLose(session *net.Session) {
	//state := module.State(session)
	switch this.battleState.BatType {
	case battle.BT_ARENA:
		module.Arena.BattleDoLose(session)
		// case battle.BT_RECOVER_BOAT:
		// 	//rpc 通知互动服 夺回战斗失败
		// 	rpc.RemoteEscortBoatRecoverBattleLose(state.PlayerId, this.boatId)
	}

	this.battleState = nil
	module.State(session).Battle = nil
}

func (this *Battle_PVE) LeaveBattle(session *net.Session) {
	//state := module.State(session)
	switch this.battleState.BatType {
	case battle.BT_ARENA:
		module.Arena.BattleDoLose(session)
		// case battle.BT_RECOVER_BOAT:
		// 	//rpc 通知互动服 夺回战斗失败
		// 	rpc.RemoteEscortBoatRecoverBattleLose(state.PlayerId, this.boatId)
	}

	this.battleState = nil
	module.State(session).Battle = nil
}
