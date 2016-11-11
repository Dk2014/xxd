package item

import (
	"game_server/mdb"
)

func tryGetPlayerItemAppendix(db *mdb.Database, playerItemAppendixId int64) (playerItemAppendix *mdb.PlayerItemAppendix) {
	if playerItemAppendixId != 0 {
		playerItemAppendix = db.Lookup.PlayerItemAppendix(playerItemAppendixId)
	} else {
		playerItemAppendix = &mdb.PlayerItemAppendix{}
	}
	return playerItemAppendix
}
