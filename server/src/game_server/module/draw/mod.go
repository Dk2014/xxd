package draw

import (
	"game_server/mdb"
	"game_server/module"
)

func init() {
	module.Draw = DrawMod{}
}

type DrawMod struct{}

func (mod DrawMod) OpenFuncForChestDraw(db *mdb.Database) {
	db.Insert.PlayerChestState(&mdb.PlayerChestState{
		Pid:             db.PlayerId(),
		IsFirstCoinTen:  1,
		IsFirstIngotTen: 1,
	})
}

func (mod DrawMod) OpenFuncForFateBox(db *mdb.Database) {
	db.Insert.PlayerFateBoxState(&mdb.PlayerFateBoxState{
		Pid:  db.PlayerId(),
		Lock: 0,
	})
}
