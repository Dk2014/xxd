package clique_building_rpc

import (
	"core/fail"
	"core/net"
	"core/time"
	"game_server/api/protocol/clique_building_api"
	"game_server/dat/clique_dat"
	"game_server/dat/mail_dat"
	"game_server/dat/player_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/module/rpc"
	"game_server/tlog"
	"game_server/xdlog"
	"strconv"
)

func cliqueTempleInfo(session *net.Session, in *clique_building_api.CliqueTempleInfo_In) {
	state := module.State(session)
	playerCliqueInfo := state.Database.Lookup.PlayerGlobalCliqueInfo(state.PlayerId)
	out := &clique_building_api.CliqueTempleInfo_Out{}

	if playerCliqueInfo == nil || playerCliqueInfo.CliqueId <= 0 {
		session.Send(out)
		return
	}

	cliqueInfo := module.CliqueRPC.CliqueInfoLookUp(state.Database, playerCliqueInfo.CliqueId)

	if cliqueInfo == nil {
		session.Send(out)
		return
	}
	playerCliqueBuilding := state.Database.Lookup.PlayerGlobalCliqueBuilding(state.PlayerId)

	if time.IsToday(playerCliqueBuilding.WorshipTime) {
		out.Isworship = true
		out.WorshipType = playerCliqueBuilding.WorshipType
	}
	if time.IsToday(cliqueInfo.WorshipTime) {
		out.Worshipcnt = cliqueInfo.WorshipCnt
	}
	if !time.IsToday(playerCliqueInfo.DonateCoinsTime) {
		playerCliqueInfo.DailyDonateCoins = 0
	}

	out.TempleBuildingCoins = int32(cliqueInfo.TempleBuildingCoins)
	out.TempleBuildingLevel = cliqueInfo.TempleBuildingLevel
	out.Totaldonatecoins = playerCliqueInfo.DailyDonateCoins

	session.Send(out)

	return
}

func cliqueTempleWorship(session *net.Session, in *clique_building_api.CliqueTempleWorship_In) {
	state := module.State(session)
	playerCliqueInfo := state.Database.Lookup.PlayerGlobalCliqueInfo(state.PlayerId)
	out := &clique_building_api.CliqueTempleWorship_Out{}

	if playerCliqueInfo == nil || playerCliqueInfo.CliqueId <= 0 {
		session.Send(out)
		return
	}
	cliqueInfo := module.CliqueRPC.CliqueInfoLookUp(state.Database, playerCliqueInfo.CliqueId)

	if cliqueInfo == nil || cliqueInfo.TempleBuildingLevel <= 0 {
		session.Send(out)
		return
	}

	if !time.IsToday(cliqueInfo.WorshipTime) {
		cliqueInfo.WorshipCnt = 0
	}

	playerCliqueBuilding := state.Database.Lookup.PlayerGlobalCliqueBuilding(state.PlayerId)

	//一天只能上一次
	if time.IsToday(playerCliqueBuilding.WorshipTime) {
		session.Send(out)
		return
	}

	var (
		worshipType int64
		feeType     int
	)

	switch in.WorshipType {
	case clique_building_api.ANCESTRAL_HALL_WORSHIP_WHITESANDALWOOD:
		worshipType = clique_dat.WHITESANDALWOOD
		feeType = player_dat.COINS
		//加属性
	case clique_building_api.ANCESTRAL_HALL_WORSHIP_STORAX:
		worshipType = clique_dat.STORAX
		feeType = player_dat.INGOT
	case clique_building_api.ANCESTRAL_HALL_WORSHIP_DAYS:
		worshipType = clique_dat.DAYS
		feeType = player_dat.INGOT
	default:
		fail.When(true, "wrong   clique   Worship type ")
	}
	//上香要扣钱
	rpc.RemoteDecMoney(state.PlayerId, worshipType, feeType, tlog.MFR_CLIQUE_DONATE_TEMPLE_WORKSHIP, xdlog.ET_CLIQUE_TEMPLE, func() {
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(state.PlayerId, func(agentDB *mdb.Database) {
				playerCliq := agentDB.Lookup.PlayerGlobalCliqueInfo(state.PlayerId)
				if playerCliq == nil || playerCliq.CliqueId <= 0 {
					session.Send(out)
					return
				}
				cliqueInfo = module.CliqueRPC.CliqueInfoLookUp(state.Database, playerCliq.CliqueId)
				if cliqueInfo == nil {
					session.Send(out)
					return
				}
				//更新状态
				playerCliqueBuilding := agentDB.Lookup.PlayerGlobalCliqueBuilding(state.PlayerId)
				playerCliqueBuilding.WorshipType = int8(in.WorshipType)
				playerCliqueBuilding.WorshipTime = time.GetNowTime()
				agentDB.Update.PlayerGlobalCliqueBuilding(playerCliqueBuilding)

				ancestra := clique_dat.GetCliqueTempleByType(int8(in.WorshipType), cliqueInfo.TempleBuildingLevel)
				if ancestra != nil {
					if ancestra.Fame > 0 {
						// 加声望   player_dat.ARENA_SYSTEM:
						rpc.RemoteAddFame(state.PlayerId, int64(ancestra.Fame), player_dat.ARENA_SYSTEM, nil)
					}
					if ancestra.Contrib > 0 {
						module.CliqueRPC.AddPlayerCliqueContrib(agentDB, int64(ancestra.Contrib))
					}
				}

				if time.IsToday(cliqueInfo.WorshipTime) {
					cliqueInfo.WorshipCnt += 1
					cliqueInfo.WorshipTime = time.GetNowTime()
					//满8次，成员获取祈福奖
					mailCoins := strconv.Itoa(int(cliqueInfo.TempleBuildingLevel) * int(clique_dat.PRAY_BASE))

					if cliqueInfo.WorshipCnt == clique_dat.WORSHCOUNT {
						//发送奖励
						pidList := module.CliqueRPC.CliqueInfoListPid(playerCliqueInfo.CliqueId)
						for _, pid := range pidList {
							rpc.RemoteMailSend(pid, &mail_dat.Mailzongcishangxiangqifujiang{
								Num: mailCoins,
								Attachments: []*mail_dat.Attachment{&mail_dat.Attachment{
									AttachmentType: mail_dat.ATTACHMENT_COINS,
									ItemId:         0,
									ItemNum:        int64(cliqueInfo.TempleBuildingLevel) * clique_dat.PRAY_BASE,
								}},
							})
						}
					}
				} else {
					cliqueInfo.WorshipCnt = 1
					cliqueInfo.WorshipTime = time.GetNowTime()
				}

				agentDB.Update.GlobalClique(cliqueInfo)
				out.Success = true
				//更新帮派每日任务 ，上香
				module.CliqueQuestRPC.UpDatePlayerCliqueDailyQuest(agentDB, clique_dat.CLIQUE_DAYLY_QUEST_WORSHIP)

				session.Send(out)
			})
		})
	})

	return
}

func cliqueTempleDonate(session *net.Session, in *clique_building_api.CliqueTempleDonate_In) {
	state := module.State(session)
	out := &clique_building_api.CliqueTempleDonate_Out{
		Result: clique_building_api.CLIQUE_BUILDING_DONATE_RESULT_FAILED,
	}
	if in.Money <= 0 {
		fail.When(true, "cliqueTempleDonate  find in.money <= 0 ")
	}

	playerCliqueInfo := state.Database.Lookup.PlayerGlobalCliqueInfo(state.PlayerId)
	if playerCliqueInfo == nil || playerCliqueInfo.CliqueId <= 0 {
		session.Send(out)
		return
	}
	cliqueInfo := module.CliqueRPC.CliqueInfoLookUp(state.Database, playerCliqueInfo.CliqueId)
	if cliqueInfo == nil {
		session.Send(out)
		return
	}
	// 帮派总舵信息
	baseInfo := clique_dat.GetCenterBuildingLevelInfo(cliqueInfo.CenterBuildingLevel)
	//clique.TempleBuildingCoins += in.Money
	fee := clique_dat.GetUpgradeFeeByLevel(cliqueInfo.TempleBuildingLevel)

	if cliqueInfo.TempleBuildingLevel > cliqueInfo.CenterBuildingLevel ||
		fee == nil ||
		fee.UpgradeFee <= 0 ||
		(cliqueInfo.TempleBuildingCoins+in.Money >= int64(fee.UpgradeFee) && cliqueInfo.TempleBuildingLevel+1 > cliqueInfo.CenterBuildingLevel) {
		session.Send(out)
		return
	}

	if !time.IsToday(playerCliqueInfo.DonateCoinsTime) {
		playerCliqueInfo.DailyDonateCoins = 0
	}
	//捐得超过限制了
	if playerCliqueInfo.DailyDonateCoins+in.Money > baseInfo.DailyMaxCoins {
		session.Send(out)
		return
	}

	rpc.RemoteDecMoney(state.PlayerId, in.Money, player_dat.COINS, tlog.MFR_CLIQUE_DONATE_TEMPLE, xdlog.ET_CLIQUE_DONATE, func() {
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(state.PlayerId, func(agentDB *mdb.Database) {
				playerCliq := agentDB.Lookup.PlayerGlobalCliqueInfo(state.PlayerId)
				if playerCliq == nil || playerCliq.CliqueId <= 0 || playerCliq.CliqueId != cliqueInfo.Id {
					session.Send(out)
					return
				}
				clique := module.CliqueRPC.CliqueInfoLookUp(state.Database, playerCliq.CliqueId)
				if clique == nil {
					session.Send(out)
					return
				}

				clique.TempleBuildingCoins += in.Money
				fee := clique_dat.GetUpgradeFeeByLevel(clique.TempleBuildingLevel)
				if fee.UpgradeFee <= 0 {
					session.Send(out)
					return
				}
				//如果费用够升级，就立马升一级
				if clique.TempleBuildingCoins >= int64(fee.UpgradeFee) {
					if clique.TempleBuildingLevel+1 <= clique.CenterBuildingLevel {
						clique.TempleBuildingLevel += 1
						clique.TempleBuildingCoins -= int64(fee.UpgradeFee)
						//升级就发动态
						upgradeLevelNotice(clique_dat.CLIQUE_BUILDING_ZONGCI, clique.Id, clique.TempleBuildingLevel)
					}
				}
				contrib := in.Money / clique_dat.CLIQUE_CONTRIB_RATE
				clique.Contrib += contrib
				clique.TotalDonateCoins += in.Money
				agentDB.Update.GlobalClique(clique)

				playerCliqueBuilding := agentDB.Lookup.PlayerGlobalCliqueBuilding(state.PlayerId)
				playerCliqueBuilding.DonateCoinsTempleBuilding += in.Money
				agentDB.Update.PlayerGlobalCliqueBuilding(playerCliqueBuilding)

				out.Totaldonatecoins = donateBuilding(agentDB, in.Money, tlog.MFR_CLIQUE_DONATE_TEMPLE)

				out.TempleBuildingCoins = int32(clique.TempleBuildingCoins)
				out.TempleBuildingLevel = clique.TempleBuildingLevel

				out.Result = clique_building_api.CLIQUE_BUILDING_DONATE_RESULT_SUCCESS
				//更新帮派每日任务 ，铜钱捐赠
				module.CliqueQuestRPC.UpDatePlayerCliqueDailyQuest(agentDB, clique_dat.CLIQUE_DAYLY_QUEST_DONATE)
				session.Send(out)

			})
		})
	})

	return
}
