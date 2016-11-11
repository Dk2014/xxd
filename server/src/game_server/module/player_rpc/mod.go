package player_rpc

import "game_server/module"

type PlayerRPCMod struct {
}

func init() {
	module.PlayerRPC = PlayerRPCMod{}
}

func (mod PlayerRPCMod) LoadGlobal() {
	LoadGlobal()
}
