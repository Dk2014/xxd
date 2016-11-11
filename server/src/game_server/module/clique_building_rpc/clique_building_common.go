package clique_building_rpc

import (
	"core/i18l"
	"core/time"
	"game_server/dat/channel_dat"
	"game_server/dat/clique_dat"
	"game_server/dat/player_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/module/rpc"
	"strconv"
)

func upgradeLevelNotice(buildingID int32, cliqueID int64, level int16) {

	var buildingNmae string
	builingInfo, ok := clique_dat.GetCliqueBuildingByID(buildingID)

	if ok {
		buildingNmae = builingInfo.Name
	}

	module.CliqueRPC.AddCliqueNews(cliqueID,
		channel_dat.MessageCliqueBuildingLevelUP{
			BuildType: channel_dat.ParamString{ // 帮派建筑名称
				Content: i18l.T.Tran(buildingNmae),
			},
			CliqueLevel: channel_dat.ParamString{
				Content: strconv.Itoa(int(level)),
			},
		},
	)
	return
}

func donateBuilding(db *mdb.Database, money int64, moneyFlow int32) int64 {
	playerCliqueInfo := db.Lookup.PlayerGlobalCliqueInfo(db.PlayerId())
	if playerCliqueInfo.CliqueId <= 0 {
		return 0
	}
	if !time.IsToday(playerCliqueInfo.DonateCoinsTime) {
		playerCliqueInfo.DailyDonateCoins = 0
	}
	playerCliqueInfo.DailyDonateCoins += money
	playerCliqueInfo.DonateCoinsTime = time.GetNowTime()
	db.Update.PlayerGlobalCliqueInfo(playerCliqueInfo)

	//玩家声望
	rpc.RemoteAddFame(db.PlayerId(), money/clique_dat.CLIQUE_FAME_RATE, player_dat.ARENA_SYSTEM, nil)
	module.CliqueRPC.AddPlayerCliqueContrib(db, money/clique_dat.CLIQUE_CONTRIB_RATE)

	return playerCliqueInfo.DailyDonateCoins
}
