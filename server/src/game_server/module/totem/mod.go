package totem

import (
	"game_server/api/protocol/notify_api"
	"game_server/dat/item_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/tlog"
	"game_server/xdlog"
)

func init() {
	module.Totem = TotemMod{}
}

type TotemMod struct{}

func (mod TotemMod) AddTotem(db *mdb.Database, totemId int16) {
	if isTotemBagFull(db) {
		sendMailWithTotem(db, []int16{totemId})
		return
	}
	playerTotemId, skillId := addTotem(db, totemId)
	if session, online := module.Player.GetPlayerOnline(db.PlayerId()); online {
		session.Send(&notify_api.NotifyNewTotem_Out{
			Id:      playerTotemId,
			TotemId: totemId,
			Skill:   skillId,
		})
	}
}

func (mod TotemMod) OpenFuncForTotem(db *mdb.Database) {
	db.Insert.PlayerTotemInfo(&mdb.PlayerTotemInfo{
		Pid: db.PlayerId(),
	})
	module.Item.AddItem(db, item_dat.ITEM_BONE, 5, tlog.IFR_MAIL_TAKE_ATTACHMENT, xdlog.ET_OPEN_FUNC, "")
}
