package player_rpc

import (
	"core/mysql"
	"game_server/api/protocol/player_api"
	"game_server/config"
	"game_server/dat/event_dat"
	"game_server/dat/player_dat"
	"game_server/module/player_common"
)

var (
	playerFightNumRankTable         = player_common.NewPlayerRankTable()
	playerLevelRankTable            = player_common.NewPlayerRankTable()
	playerMainRoleFightNumRankTable = player_common.NewPlayerRankTable()
	playerIngotRankTable            = player_common.NewPlayerRankTable()

	//控制显示排行榜玩家数目
	desplay_num                      = 0
	rankTypeFlag player_api.RankType = player_api.RANK_TYPE_NULL
)

func LoadGlobal() {
	db, err1 := mysql.Connect(config.GetDBConfig())
	if err1 != nil {
		panic(err1)
	}
	defer db.Close()

	// fight num
	loadLevelAndFightNum(db)
	loadIngot(db)
	//	loadMainRoleFightNum(db)
}

func loadLevelAndFightNum(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte(`SELECT A.id as pid, A.nick, B.level, B.timestamp, C.fight_num FROM player AS A
		LEFT JOIN player_role AS B  ON A.id = B.pid and (B.role_id = 1 or B.role_id = 2)
		LEFT JOIN player_fight_num AS C ON A.id = C.pid`), -1)
	if err != nil {
		panic(err)
	}

	iPid := res.Map("pid")
	iNick := res.Map("nick")
	iFightNum := res.Map("fight_num")
	iLevel := res.Map("level")
	iTimestamp := res.Map("timestamp")

	datas1 := []*player_common.PlayerRankData{}
	datas2 := []*player_common.PlayerRankData{}
	for _, row := range res.Rows {
		level := row.Int16(iLevel)
		if level < player_dat.RANK_OPEN_MIN_LEVEL {
			continue
		}

		timestamp := row.Int64(iTimestamp)
		data1 := player_common.NewPlayerRankData()
		data1.Timestamp = timestamp
		data1.Pid = row.Int64(iPid)
		data1.Nick = row.Str(iNick)
		data1.Values = append(data1.Values, player_api.GetRanks_Out_Ranks_Values{
			Flag:  player_api.RANK_TYPE_LEVEL,
			Value: int64(row.Int16(iLevel)),
		})
		data1.Values = append(data1.Values, player_api.GetRanks_Out_Ranks_Values{
			Flag:  player_api.RANK_TYPE_FIGHTNUM,
			Value: int64(row.Int32(iFightNum)),
		})
		datas1 = append(datas1, data1)

		data2 := player_common.NewPlayerRankData()
		data2.Pid = row.Int64(iPid)
		data2.Nick = row.Str(iNick)
		data2.Values = append(data2.Values, player_api.GetRanks_Out_Ranks_Values{
			Flag:  player_api.RANK_TYPE_FIGHTNUM,
			Value: int64(row.Int32(iFightNum)),
		})
		data2.Values = append(data2.Values, player_api.GetRanks_Out_Ranks_Values{
			Flag:  player_api.RANK_TYPE_LEVEL,
			Value: int64(row.Int16(iLevel)),
		})
		datas2 = append(datas2, data2)
	}
	playerLevelRankTable.Update(datas1)
	playerFightNumRankTable.Update(datas2)
}

func loadMainRoleFightNum(db *mysql.Connection) {
	// TODO
}

func loadIngot(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte(`select  A.pid, A.ingot,A.timestamp, B.level, C.nick from player_add_ingot_record as A left join player_role as B on A.pid = B.id and (B.role_id=1 or B.role_id=2) left join player as C on A.pid=C.id`), -1)
	//res, err := db.ExecuteFetch([]byte(`select  A.id, C.ingot, A.nick, B.level, C.timestamp from player as A left join player_role as B on A.id=B.pid and (B.role_id=1 or B.role_id=2) left join player_add_ingot_record as C on A.id=C.pid`), -1)
	if err != nil {
		panic(err)
	}

	iPid := res.Map("pid")
	iIngot := res.Map("ingot")
	iNick := res.Map("nick")
	iLevel := res.Map("level")
	iTimestamp := res.Map("timestamp")

	eventIngotRankConfig := event_dat.GetEventIngotRankConfig()
	openIngotRankTime := eventIngotRankConfig.StartUnixTime
	endIngotRankTime := eventIngotRankConfig.EndUnixTime
	datas := []*player_common.PlayerRankData{}

	for _, row := range res.Rows {
		timestamp := row.Int64(iTimestamp)
		if timestamp <= openIngotRankTime || timestamp >= endIngotRankTime {
			continue
		}

		data := player_common.NewPlayerRankData()
		data.Timestamp = timestamp
		data.Pid = row.Int64(iPid)
		data.Nick = row.Str(iNick)
		data.Values = append(data.Values, player_api.GetRanks_Out_Ranks_Values{
			Flag:  player_api.RANK_TYPE_INGOT,
			Value: int64(row.Int64(iIngot)),
		})
		data.Values = append(data.Values, player_api.GetRanks_Out_Ranks_Values{
			Flag:  player_api.RANK_TYPE_LEVEL,
			Value: int64(row.Int16(iLevel)),
		})
		datas = append(datas, data)
	}
	playerIngotRankTable.Update(datas)
}
func UpdateGlobal(rankType player_api.RankType, datas []*player_common.PlayerRankData) {
	switch rankType {

	case player_api.RANK_TYPE_FIGHTNUM:
		playerFightNumRankTable.Update(datas)

	case player_api.RANK_TYPE_LEVEL:
		playerLevelRankTable.Update(datas)

	case player_api.RANK_TYPE_MAINROLE_FIGHTNUM:
		playerMainRoleFightNumRankTable.Update(datas)

	case player_api.RANK_TYPE_INGOT:
		playerIngotRankTable.Update(datas)
	}
}
