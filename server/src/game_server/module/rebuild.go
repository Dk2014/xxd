package module

import (
	"core/net"
)

func TryRebuildState(session *net.Session) {

	state := State(session)
	db := state.Database
	playerId := state.PlayerId

	// 装备重铸状态重建
	playerItemRecastState := db.Lookup.PlayerItemRecastState(playerId)
	if playerItemRecastState != nil {
		Notify.SendItemRecastStateRebuild(session)
	}

}
