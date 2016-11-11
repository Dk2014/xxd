package clique_rpc

import (
	"core/time"
	"game_server/api/protocol/clique_api"
	"game_server/dat/clique_dat"
	"game_server/dat/mail_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/module/rpc"
	"sort"
)

type byApply []*JoinApply

func (a byApply) Len() int           { return len(a) }
func (a byApply) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byApply) Less(i, j int) bool { return a[i].ArenaRank < a[j].ArenaRank }

func resetPlayerCliqueContrib(db *mdb.Database) {
	//record := db.Lookup.PlayerGlobalCliqueContrib(db.PlayerId())
	//if record != nil {
	//	record.Contrib = 0
	//	db.Update.PlayerGlobalCliqueContrib(record)
	//}

}

//重置玩家帮派信息
func resetPlayerCliqueState(player *mdb.PlayerGlobalCliqueInfo, playerBuilding *mdb.PlayerGlobalCliqueBuilding, cliqueInfo *mdb.GlobalClique) {
	player.CliqueId = 0
	player.JoinTime = 0
	player.Contrib = 0
	player.TotalContrib = 0
	player.DonateCoinsTime = 0
	player.DailyDonateCoins = 0
	//player.ContribClearTime = 0
	playerBuilding.DonateCoinsCenterBuilding = 0
	playerBuilding.DonateCoinsTempleBuilding = 0
	playerBuilding.DonateCoinsHealthBuilding = 0
	playerBuilding.DonateCoinsDefenseBuilding = 0
	playerBuilding.DonateCoinsAttackBuilding = 0
	playerBuilding.DonateCoinsBankBuilding = 0

	if playerBuilding.SilverExchangeNum > 0 || playerBuilding.GoldExchangeNum > 0 {
		bankDatInfo := clique_dat.GetCliqueBuildingBankInfo(cliqueInfo.BankBuildingLevel)
		var attachments []*mail_dat.Attachment
		if playerBuilding.SilverExchangeNum > 0 {
			attachments = append(attachments, &mail_dat.Attachment{
				AttachmentType: mail_dat.ATTACHMENT_COINS,
				ItemId:         0,
				ItemNum:        int64(playerBuilding.SilverExchangeNum) * bankDatInfo.SilverCouponCoins,
			})
		}
		if playerBuilding.GoldExchangeNum > 0 {
			attachments = append(attachments, &mail_dat.Attachment{
				AttachmentType: mail_dat.ATTACHMENT_INGOT,
				ItemId:         0,
				ItemNum:        int64(playerBuilding.GoldExchangeNum) * bankDatInfo.GoldCouponCoins,
			})
		}
		rpc.RemoteMailSend(player.Pid, &mail_dat.Mailcliqueleavebank{
			Attachments: attachments,
		})
	}
	playerBuilding.SilverExchangeNum = 0
	playerBuilding.SilverExchangeTime = 0
	playerBuilding.GoldExchangeNum = 0
	playerBuilding.GlodExchangeTime = 0
}

//扣除玩家建筑捐献
func withdrawPlayerCliqueContrib(playerCliqueBuilding *mdb.PlayerGlobalCliqueBuilding, cliqueInfo *mdb.GlobalClique) {
	cliqueInfo.CenterBuildingCoins -= playerCliqueBuilding.DonateCoinsCenterBuilding
	cliqueInfo.TempleBuildingCoins -= playerCliqueBuilding.DonateCoinsTempleBuilding
	cliqueInfo.HealthBuildingCoins -= playerCliqueBuilding.DonateCoinsHealthBuilding
	cliqueInfo.AttackBuildingCoins -= playerCliqueBuilding.DonateCoinsAttackBuilding
	cliqueInfo.DefenseBuildingCoins -= playerCliqueBuilding.DonateCoinsDefenseBuilding
	cliqueInfo.BankBuildingCoins -= playerCliqueBuilding.DonateCoinsBankBuilding
}

func isCliqueAuthority(state *module.SessionState, pid int64, cliqueID int64) bool {
	cliqueInfo := cliqueInfoLookUp(state.Database, cliqueID)
	if cliqueInfo == nil {
		return false
	}

	if isOwer(cliqueInfo, pid) || isManger(cliqueInfo, pid) {
		return true
	}

	return false
}

func isOwer(cliqueInfo *mdb.GlobalClique, pid int64) bool {

	if cliqueInfo.OwnerPid == pid {
		return true
	}
	return false
}

func isManger(cliqueInfo *mdb.GlobalClique, pid int64) bool {

	if cliqueInfo.MangerPid1 == pid || cliqueInfo.MangerPid2 == pid {
		return true
	}

	return false
}

func getCliqueApplys(cliqueId int64, limit, offset int16, listApply *clique_api.ListApply_Out) {
	cliqueCache := CacheGetCliqueInfo(cliqueId)
	applys := make([]*JoinApply, 0)
	for _, member := range cliqueCache.JoinApplies {
		applys = append(applys, member)
	}

	length := len(applys)
	if length > 0 {
		sort.Sort(byApply(applys))
	}

	if length >= int(limit+offset) { //足够多
		for i := offset; i < limit; i++ {
			listApply.Players = append(listApply.Players, clique_api.ListApply_Out_Players{
				Pid:       applys[i].Pid,
				Nick:      []byte(applys[i].Nick),
				Level:     applys[i].Level,
				ArenaRank: int64(applys[i].ArenaRank),
				Timestamp: applys[i].Timestamp,
			})
		}
	} else {
		for i := int(offset); i < length; i++ {
			listApply.Players = append(listApply.Players, clique_api.ListApply_Out_Players{
				Pid:       applys[i].Pid,
				Nick:      []byte(applys[i].Nick),
				Level:     applys[i].Level,
				ArenaRank: int64(applys[i].ArenaRank),
				Timestamp: applys[i].Timestamp,
			})
		}
	}

	return
}

func getPlayerGlobalCliqueInfo(db *mdb.Database) *mdb.PlayerGlobalCliqueInfo {
	playerCliqueInfo := db.Lookup.PlayerGlobalCliqueInfo(db.PlayerId())
	if playerCliqueInfo == nil {
		playerCliqueInfo = &mdb.PlayerGlobalCliqueInfo{
			Pid: db.PlayerId(),
		}
		db.Insert.PlayerGlobalCliqueInfo(playerCliqueInfo)

		playerCliqueBuilding := &mdb.PlayerGlobalCliqueBuilding{
			Pid: db.PlayerId(),
		}
		db.Insert.PlayerGlobalCliqueBuilding(playerCliqueBuilding)
	}
	return playerCliqueInfo
}

//查询帮派信息接口封装
//实现查询时，如果满足清零条件直接清零
func cliqueInfoLookUp(db *mdb.Database, cliqueId int64) (cliqueInfo *mdb.GlobalClique) {
	cliqueInfo = db.Lookup.GlobalClique(cliqueId)
	if cliqueInfo == nil {
		return nil
	}
	if !time.IsThisWeek(cliqueInfo.ContribClearTime) {
		cliqueInfo.Contrib = 0
		cliqueInfo.ContribClearTime = time.GetNowTime()
		//无需更新
	}
	return cliqueInfo
}

//离开帮派操作统一操作
func leaveCliqueCleanup(db *mdb.Database) {
	module.CliqueQuestRPC.CleanCliqueBuildingQuest(db, db.PlayerId())
}

//加入帮派操作统一操作
func joinCliqueCleanup(db *mdb.Database, playerCliqueInfo *mdb.PlayerGlobalCliqueInfo, cliqueInfo *mdb.GlobalClique) {
	//加入帮派时把镖船迁移到新帮派
	module.CliqueEscortRPC.MigrateBoatToNewClique(db, playerCliqueInfo.CliqueId)
	//帮派任务
	module.CliqueQuestRPC.AddCliqueBuildingQuest(cliqueInfo, db, db.PlayerId())
}
