package friend_rpc

import (
	"core/time"
	"game_server/mdb"
	"game_server/module"
)

func init() {
	module.ChatRPC = ChatModRPC{}
}

type ChatModRPC struct{}

func (mod ChatModRPC) SetBanState(db *mdb.Database, banTime int64) {
	var banUntil int64
	if banTime <= 0 {
		banUntil = 0
	} else {
		banUntil = time.GetNowTime() + banTime
	}
	chatState := db.Lookup.PlayerGlobalChatState(db.PlayerId())
	if chatState == nil {
		db.Insert.PlayerGlobalChatState(&mdb.PlayerGlobalChatState{
			Pid:      db.PlayerId(),
			BanUntil: banUntil,
		})
	} else {
		chatState.BanUntil = banUntil
		db.Update.PlayerGlobalChatState(chatState)
	}
}

func (mod ChatModRPC) CanChat(db *mdb.Database) bool {
	chatState := db.Lookup.PlayerGlobalChatState(db.PlayerId())
	nowTime := time.GetNowTime()
	if chatState == nil || chatState.BanUntil <= nowTime {
		return true
	}
	return false
}
