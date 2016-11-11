package town

import (
	"core/log"
	"core/net"
	"game_server/api/protocol/town_api"
	"game_server/module"
)

/*
城镇相关全局变量和常量
*/

func init() {
	module.Town = TownMod{}
}

type TownMod struct {
}

func (mod TownMod) LeaveTown(session *net.Session) {
	state := module.State(session)

	if state.TownChannel != nil {
		townChan := state.TownChannel

		g_TownChannelMgr.removePlayer(session)

		module.API.Broadcast(townChan, &town_api.Leave_Out{
			PlayerId: state.PlayerId,
		})
	}

	if state.TownId > 0 {
		playerTown := state.Database.Lookup.PlayerTown(state.PlayerId)
		playerTown.TownId = state.TownId
		playerTown.AtPosX = state.LastTownX
		playerTown.AtPosY = state.LastTownY
		state.Database.Update.PlayerTown(playerTown)
	}

	state.TownId, state.LastTownX, state.LastTownY = 0, 0, 0
}

func (mod TownMod) SetTownLock(state *module.SessionState, newLock int32) {
	playerTown := state.Database.Lookup.PlayerTown(state.PlayerId)
	if playerTown.Lock >= newLock {
		log.Errorf("set town lock errr old loack :%d new lock :%d", playerTown.Lock, newLock)
		return
	}

	playerTown.Lock = newLock
	state.Database.Update.PlayerTown(playerTown)

	if session, ok := module.Player.GetPlayerOnline(state.PlayerId); ok {
		module.Notify.SendTownLockChange(session, newLock)
	}
}
