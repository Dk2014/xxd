package clique_building_rpc

import (
	"core/net"
	"core/time"
	"game_server/api/protocol/clique_building_api"
	"game_server/dat/clique_dat"
	"game_server/dat/player_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/module/rpc"
	"game_server/tlog"
	"game_server/xdlog"
)

func cliqueBuildingBaseDonate(session *net.Session, money int64) {
	state := module.State(session)
	out := &clique_building_api.CliqueBaseDonate_Out{}
	playerCliqueInfo := state.Database.Lookup.PlayerGlobalCliqueInfo(state.PlayerId)
	if playerCliqueInfo == nil || playerCliqueInfo.CliqueId <= 0 {
		out.Result = clique_building_api.CLIQUE_BUILDING_DONATE_RESULT_FAILED
		session.Send(out)
		return
	}
	cliqueInfo := module.CliqueRPC.CliqueInfoLookUp(state.Database, playerCliqueInfo.CliqueId)
	if cliqueInfo == nil {
		out.Result = clique_building_api.CLIQUE_BUILDING_DONATE_RESULT_FAILED
		session.Send(out)
		return
	}
	baseInfo := clique_dat.GetCenterBuildingLevelInfo(cliqueInfo.CenterBuildingLevel)
	if !baseInfo.CanUpgrade {
		out.Result = clique_building_api.CLIQUE_BUILDING_DONATE_RESULT_FAILED
		session.Send(out)
		return
	}

	// 验证玩家今日捐钱是否超过上限，然后rpc扣钱，回调函数更新捐钱 判断是否升级
	donatedCoins := playerCliqueInfo.DailyDonateCoins
	if !time.IsToday(playerCliqueInfo.DonateCoinsTime) {
		donatedCoins = 0
	}
	if donatedCoins+money > baseInfo.DailyMaxCoins {
		out.Result = clique_building_api.CLIQUE_BUILDING_DONATE_RESULT_FAILED
		session.Send(out)
		return
	}
	oldCliqueId := cliqueInfo.Id
	rpc.RemoteDecMoney(state.PlayerId, money, player_dat.COINS, tlog.MFR_CLIQUE_DONATE_BASE, xdlog.ET_CLIQUE_DONATE, func() {
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(state.PlayerId, func(agentDB *mdb.Database) {
				playerCliqueInfo = agentDB.Lookup.PlayerGlobalCliqueInfo(agentDB.PlayerId())
				if playerCliqueInfo == nil || playerCliqueInfo.CliqueId <= 0 || playerCliqueInfo.CliqueId != oldCliqueId {
					return
				}
				cliqueInfo := module.CliqueRPC.CliqueInfoLookUp(agentDB, playerCliqueInfo.CliqueId)
				baseInfo := clique_dat.GetCenterBuildingLevelInfo(cliqueInfo.CenterBuildingLevel)
				// 判断是否需要升级总舵，以及更新玩家表
				if cliqueInfo.CenterBuildingCoins+money >= baseInfo.LevelupCoins && baseInfo.CanUpgrade {
					cliqueInfo.CenterBuildingLevel += 1
					cliqueInfo.CenterBuildingCoins = cliqueInfo.CenterBuildingCoins + money - baseInfo.LevelupCoins
					//升级就发动态
					upgradeLevelNotice(clique_dat.CLIQUE_BUILDING_ZONGDUO, cliqueInfo.Id, cliqueInfo.CenterBuildingLevel)

				} else {
					cliqueInfo.CenterBuildingCoins += money
				}
				agentDB.Update.GlobalClique(cliqueInfo)

				playerCliqueBuilding := agentDB.Lookup.PlayerGlobalCliqueBuilding(state.PlayerId)
				playerCliqueBuilding.DonateCoinsCenterBuilding += money
				agentDB.Update.PlayerGlobalCliqueBuilding(playerCliqueBuilding)

				out.PlayerDonateCoins = donateBuilding(agentDB, money, tlog.MFR_CLIQUE_DONATE_BASE)

				out.Result = clique_building_api.CLIQUE_BUILDING_DONATE_RESULT_SUCCESS
				out.CliqueBuildingBaseDonateCoins = int32(cliqueInfo.CenterBuildingCoins)
				out.CliqueBuildingBaseLevel = cliqueInfo.CenterBuildingLevel

				//更新帮派每日任务 ，铜钱捐赠
				module.CliqueQuestRPC.UpDatePlayerCliqueDailyQuest(agentDB, clique_dat.CLIQUE_DAYLY_QUEST_DONATE)
				session.Send(out)
			})
		})
	})
}

func cliqueBuildingStatus(session *net.Session) {
	state := module.State(session)
	out := &clique_building_api.CliqueBuildingStatus_Out{}
	playerCliqueInfo := state.Database.Lookup.PlayerGlobalCliqueInfo(state.PlayerId)
	if playerCliqueInfo == nil || playerCliqueInfo.CliqueId <= 0 {
		session.Send(out)
		return
	}
	playerCliqueBuilding := state.Database.Lookup.PlayerGlobalCliqueBuilding(state.PlayerId)
	cliqueInfo := module.CliqueRPC.CliqueInfoLookUp(state.Database, playerCliqueInfo.CliqueId)
	if cliqueInfo == nil {
		session.Send(out)
		return
	}
	// 验证玩家今日捐钱是否超过上限，然后rpc扣钱，回调函数更新捐钱 判断是否升级
	donatedCoins := playerCliqueInfo.DailyDonateCoins
	if !time.IsToday(playerCliqueInfo.DonateCoinsTime) {
		donatedCoins = 0
	}

	out.DailyTotalDonatedCoins = donatedCoins

	out.Base = clique_building_api.CliqueBuildingStatusBase{
		Level:       cliqueInfo.CenterBuildingLevel,
		DonateCoins: cliqueInfo.CenterBuildingCoins,
	}

	out.Bank = clique_building_api.CliqueBuildingStatusBank{
		Level:                cliqueInfo.BankBuildingLevel,
		DonateCoins:          int32(cliqueInfo.BankBuildingCoins),
		SilverCouponNum:      playerCliqueBuilding.SilverExchangeNum,
		SilverCouponTimespan: playerCliqueBuilding.SilverExchangeTime,
		GoldCouponNum:        playerCliqueBuilding.GoldExchangeNum,
		GoldCouponTimespan:   playerCliqueBuilding.GlodExchangeTime,
	}

	out.AttackBuilding = clique_building_api.CliqueBuildingStatusAttack{
		Level:       cliqueInfo.AttackBuildingLevel,
		DonateCoins: int32(cliqueInfo.AttackBuildingCoins),
	}

	out.HealthBuilding = clique_building_api.CliqueBuildingStatusHealth{
		Level:       cliqueInfo.HealthBuildingLevel,
		DonateCoins: int32(cliqueInfo.HealthBuildingCoins),
	}

	out.DefenceBuilding = clique_building_api.CliqueBuildingStatusDefence{
		Level:       cliqueInfo.DefenseBuildingLevel,
		DonateCoins: int32(cliqueInfo.DefenseBuildingCoins),
	}
	out.TempleBuilding = clique_building_api.CliqueBuildingStatusTemple{
		Level:       cliqueInfo.TempleBuildingLevel,
		DonateCoins: int32(cliqueInfo.TempleBuildingCoins),
	}

	out.Success = true
	session.Send(out)
}
