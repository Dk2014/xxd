package global

import (
	"game_server/mdb"
)

func NewPlayerMount(db *mdb.Database, info *PlayerInfo) {
	// 为新玩家分配一个数据切片，并挂载到database
	mdb.NewPlayerTablesWithDatabase(db, info.PlayerId, info.CId)

	playerInfoDataTable.addPlayerInfo(info)

	InitPlayerData(db)
}

/*
	新玩家在互动服初始化数据
*/
func CheckPlayerHaveInitData(db *mdb.Database) bool {
	return db.Lookup.PlayerGlobalFriendState(db.PlayerId()) != nil
}

func InitPlayerData(db *mdb.Database) {
	// 好友系统
	db.Insert.PlayerGlobalFriendState(&mdb.PlayerGlobalFriendState{
		Pid: db.PlayerId(),
	})
	//全局邮件
	db.Insert.PlayerGlobalMailState(&mdb.PlayerGlobalMailState{
		Pid:          db.PlayerId(),
		MaxTimestamp: 0,
	})
}
