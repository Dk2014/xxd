package player

import (
	"game_server/api/protocol/player_api"
	"game_server/dat/player_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/module/player_common"
)

func (mod PlayerMod) ForceUpdateLevel() {
	updateLevel.update()
}
func (mod PlayerMod) ForceUpdateFightNum() {
	updateFightNum.update()
}

func (mod PlayerMod) AddUpdateLevel(db *mdb.Database, newLevel int16, timestamp int64) {
	if newLevel < player_dat.RANK_OPEN_MIN_LEVEL {
		return
	}
	playerInfo := db.Lookup.Player(db.PlayerId())
	playerFightNum := db.Lookup.PlayerFightNum(db.PlayerId())

	data := player_common.NewPlayerRankData()
	data.Pid = db.PlayerId()
	data.Nick = playerInfo.Nick
	data.Values = append(data.Values, player_api.GetRanks_Out_Ranks_Values{
		Flag:  player_api.RANK_TYPE_LEVEL,
		Value: int64(newLevel),
	})
	data.Values = append(data.Values, player_api.GetRanks_Out_Ranks_Values{
		Flag:  player_api.RANK_TYPE_FIGHTNUM,
		Value: int64(playerFightNum.FightNum),
	})
	data.Timestamp = timestamp

	updateLevel.AddData(data)
}

func (mod PlayerMod) AddUpdateFightNum(db *mdb.Database, fightNum int32) {
	playerInfo := db.Lookup.Player(db.PlayerId())
	playerRole := module.Role.GetMainRole(db)
	if playerRole.Level < player_dat.RANK_OPEN_MIN_LEVEL {
		return
	}

	data := player_common.NewPlayerRankData()
	data.Pid = db.PlayerId()
	data.Nick = playerInfo.Nick
	data.Values = append(data.Values, player_api.GetRanks_Out_Ranks_Values{
		Flag:  player_api.RANK_TYPE_FIGHTNUM,
		Value: int64(fightNum),
	})
	data.Values = append(data.Values, player_api.GetRanks_Out_Ranks_Values{
		Flag:  player_api.RANK_TYPE_LEVEL,
		Value: int64(playerRole.Level),
	})

	updateFightNum.AddData(data)
}

func (mod PlayerMod) AddUpdateMainRoleFightNum(db *mdb.Database, fightNum int32) {
	playerInfo := db.Lookup.Player(db.PlayerId())
	playerRole := module.Role.GetMainRole(db)
	if playerRole.Level < player_dat.RANK_OPEN_MIN_LEVEL {
		return
	}

	data := player_common.NewPlayerRankData()
	data.Pid = db.PlayerId()
	data.Nick = playerInfo.Nick
	data.Values = append(data.Values, player_api.GetRanks_Out_Ranks_Values{
		Flag:  player_api.RANK_TYPE_MAINROLE_FIGHTNUM,
		Value: int64(fightNum),
	})
	data.Values = append(data.Values, player_api.GetRanks_Out_Ranks_Values{
		Flag:  player_api.RANK_TYPE_LEVEL,
		Value: int64(playerRole.Level),
	})

	updateMainRoleFightNum.AddData(data)
}

func (mod PlayerMod) AddUpdateInGot(db *mdb.Database, timestamp int64) {
	playerInfo := db.Lookup.Player(db.PlayerId())
	playerRole := module.Role.GetMainRole(db)

	playerAddIngotRecord := db.Lookup.PlayerAddIngotRecord(db.PlayerId())
	data := player_common.NewPlayerRankData()
	data.Pid = db.PlayerId()
	data.Nick = playerInfo.Nick
	data.Timestamp = timestamp
	data.Values = append(data.Values, player_api.GetRanks_Out_Ranks_Values{
		Flag:  player_api.RANK_TYPE_INGOT,
		Value: playerAddIngotRecord.Ingot,
	})
	data.Values = append(data.Values, player_api.GetRanks_Out_Ranks_Values{
		Flag:  player_api.RANK_TYPE_LEVEL,
		Value: int64(playerRole.Level),
	})

	updateIngot.AddData(data)
}
