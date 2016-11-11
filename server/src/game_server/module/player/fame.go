package player

import (
	"core/fail"
	"game_server/api/protocol/notify_api"
	"game_server/dat/player_dat"
	"game_server/mdb"
	"game_server/module"
)

/*
//!!!!!即将被废弃
func switchSystemFamePointer(playerFame *mdb.PlayerFame, system int16) (fameVal *int32) {
	switch system {
	case player_dat.ARENA_SYSTEM:
		return &playerFame.ArenaFame
	case player_dat.MULT_LEVEL_SYSTEM:
		return &playerFame.MultLevelFame
	}
	//fail.When(true, "未知系统类型")
	return nil
}

func getSystemFame(db *mdb.Database, system int16) int32 {
	playerFame := db.Lookup.PlayerFame(db.PlayerId())
	famePtr := switchSystemFamePointer(playerFame, system)
	if famePtr == nil {
		return 0
	}
	return *famePtr
}
*/

func addFame(db *mdb.Database, val int32) {
	fail.When(val <= 0, "invalid fame")
	playerFame := db.Lookup.PlayerFame(db.PlayerId())
	playerFame.Fame += val
	requiredFame, nextLv := player_dat.NextLevelRequiredFame(playerFame.Level + 1)
	var fameLvUp bool
	for nextLv && playerFame.Fame >= requiredFame {
		fameLvUp = true
		playerFame.Level += 1

		requiredFame, nextLv = player_dat.NextLevelRequiredFame(playerFame.Level + 1)
	}
	//更新所有伙伴技能
	if fameLvUp {
		db.Select.PlayerRole(func(row *mdb.PlayerRoleRow) {
			module.Skill.UpdateSkill(db, row.RoleId(), int16(row.FriendshipLevel()), playerFame.Level, row.Level())
		})
	}
	if session, online := module.Player.GetPlayerOnline(db.PlayerId()); online {
		session.Send(&notify_api.FameChange_Out{
			Fame: playerFame.Fame,
		})
	}

	db.Update.PlayerFame(playerFame)
}
