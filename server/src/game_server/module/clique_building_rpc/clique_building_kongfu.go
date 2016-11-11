package clique_building_rpc

import (
	"core/fail"
	"core/net"
	"core/time"
	"fmt"
	"game_server/api/protocol/clique_building_api"
	"game_server/dat/clique_dat"
	"game_server/dat/player_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/module/rpc"
	"game_server/tlog"
	"game_server/xdlog"
)

func kongfuBuildingDonate(session *net.Session, buildingId int32, money int64) {
	fail.When(money <= 0, "参数 money <= 0 ")
	kongfuBuildingCheck(buildingId)

	state := module.State(session)
	playerCliqueInfo := state.Database.Lookup.PlayerGlobalCliqueInfo(state.PlayerId)
	out := &clique_building_api.CliqueKongfuDonate_Out{
		Building: buildingId,
	}
	if playerCliqueInfo == nil || playerCliqueInfo.CliqueId <= 0 {
		out.Result = clique_building_api.CLIQUE_BUILDING_DONATE_RESULT_FAILED
		session.Send(out)
		return
	}
	cliqueInfo := module.CliqueRPC.CliqueInfoLookUp(state.Database, playerCliqueInfo.CliqueId)
	baseInfo := clique_dat.GetCenterBuildingLevelInfo(cliqueInfo.CenterBuildingLevel)
	oldCliqueId := cliqueInfo.Id

	//假设 总部等级和xx堂等级一样
	//假设 帮派还差 10 铜币升级， 玩家捐了11
	//这里不会返回捐款失败， 因为没有什么好办法避免 *101/100 捐的钱够了但是没有升级* 这种情况
	//这个时候在让他钱累加但是等级不要上去

	donatedCoins := playerCliqueInfo.DailyDonateCoins
	if !time.IsToday(playerCliqueInfo.DonateCoinsTime) {
		donatedCoins = 0
	}
	if donatedCoins+money > baseInfo.DailyMaxCoins {
		out.Result = clique_building_api.CLIQUE_BUILDING_DONATE_RESULT_FAILED
		session.Send(out)
		return
	}

	var flowID int32
	switch buildingId {
	case clique_dat.CLIQUE_BUILDING_HUICHUNTANG:
		buildingDat := clique_dat.GetKongfuBuildingDat(clique_dat.CLIQUE_BUILDING_HUICHUNTANG, cliqueInfo.HealthBuildingLevel)
		cliqueInfo.HealthBuildingCoins += money
		if !buildingDat.CanUpgrade ||
			cliqueInfo.HealthBuildingLevel+1 > cliqueInfo.CenterBuildingLevel && cliqueInfo.HealthBuildingCoins >= buildingDat.UpgradeCoins {
			out.Result = clique_building_api.CLIQUE_BUILDING_DONATE_RESULT_FAILED
			session.Send(out)
			return
		}
		flowID = tlog.MFR_CLIQUE_DONATE_KONGFU_HUICHUNTANG
	case clique_dat.CLIQUE_BUILDING_SHENBINGTANG:
		buildingDat := clique_dat.GetKongfuBuildingDat(buildingId, cliqueInfo.AttackBuildingLevel)
		cliqueInfo.AttackBuildingCoins += money
		if !buildingDat.CanUpgrade ||
			cliqueInfo.AttackBuildingLevel+1 > cliqueInfo.CenterBuildingLevel && cliqueInfo.AttackBuildingCoins >= buildingDat.UpgradeCoins {

			out.Result = clique_building_api.CLIQUE_BUILDING_DONATE_RESULT_FAILED
			session.Send(out)
			return
		}
		flowID = tlog.MFR_CLIQUE_DONATE_KONGFU_SHENBINGTANG
	case clique_dat.CLIQUE_BUILDING_JINGANGTANG:
		buildingDat := clique_dat.GetKongfuBuildingDat(clique_dat.CLIQUE_BUILDING_HUICHUNTANG, cliqueInfo.DefenseBuildingLevel)
		cliqueInfo.DefenseBuildingCoins += money
		if !buildingDat.CanUpgrade ||
			cliqueInfo.DefenseBuildingLevel+1 > cliqueInfo.CenterBuildingLevel && cliqueInfo.DefenseBuildingCoins >= buildingDat.UpgradeCoins {

			out.Result = clique_building_api.CLIQUE_BUILDING_DONATE_RESULT_FAILED
			session.Send(out)
			return
		}
		flowID = tlog.MFR_CLIQUE_DONATE_KONGFU_JINGANGTANG
	default:
		fail.When(true, fmt.Sprintf("undefind kongfu building %d", buildingId))
	}

	rpc.RemoteDecMoney(state.PlayerId, money, player_dat.COINS, flowID, xdlog.ET_CLIQUE_DONATE, func() {
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(state.PlayerId, func(agentDB *mdb.Database) {
				playerCliqueInfo := agentDB.Lookup.PlayerGlobalCliqueInfo(state.PlayerId)
				if playerCliqueInfo.CliqueId != oldCliqueId {
					//玩家没有帮派 或者 换了帮派
					return
				}
				cliqueInfo := module.CliqueRPC.CliqueInfoLookUp(agentDB, playerCliqueInfo.CliqueId)
				playerCliqueBuilding := agentDB.Lookup.PlayerGlobalCliqueBuilding(state.PlayerId)
				upgrade := false

				switch buildingId {
				case clique_dat.CLIQUE_BUILDING_HUICHUNTANG:
					//回春堂
					buildingDat := clique_dat.GetKongfuBuildingDat(clique_dat.CLIQUE_BUILDING_HUICHUNTANG, cliqueInfo.HealthBuildingLevel)
					cliqueInfo.HealthBuildingCoins += money
					if buildingDat.CanUpgrade &&
						cliqueInfo.HealthBuildingLevel+1 <= cliqueInfo.CenterBuildingLevel &&
						cliqueInfo.HealthBuildingCoins >= buildingDat.UpgradeCoins {

						cliqueInfo.HealthBuildingCoins -= buildingDat.UpgradeCoins
						cliqueInfo.HealthBuildingLevel += 1
						upgrade = true
					}
					playerCliqueBuilding.DonateCoinsHealthBuilding += money
					out.TotalDonateCoins = int32(cliqueInfo.HealthBuildingCoins)
					out.Level = cliqueInfo.HealthBuildingLevel

				case clique_dat.CLIQUE_BUILDING_SHENBINGTANG:
					//神兵堂
					buildingDat := clique_dat.GetKongfuBuildingDat(buildingId, cliqueInfo.AttackBuildingLevel)
					cliqueInfo.AttackBuildingCoins += money
					if buildingDat.CanUpgrade &&
						cliqueInfo.AttackBuildingLevel+1 <= cliqueInfo.CenterBuildingLevel &&
						cliqueInfo.AttackBuildingCoins >= buildingDat.UpgradeCoins {

						cliqueInfo.AttackBuildingCoins -= buildingDat.UpgradeCoins
						cliqueInfo.AttackBuildingLevel += 1
						upgrade = true
					}
					playerCliqueBuilding.DonateCoinsAttackBuilding += money
					out.TotalDonateCoins = int32(cliqueInfo.AttackBuildingCoins)
					out.Level = cliqueInfo.AttackBuildingLevel
				case clique_dat.CLIQUE_BUILDING_JINGANGTANG:
					//金刚堂
					buildingDat := clique_dat.GetKongfuBuildingDat(clique_dat.CLIQUE_BUILDING_HUICHUNTANG, cliqueInfo.DefenseBuildingLevel)
					cliqueInfo.DefenseBuildingCoins += money
					if buildingDat.CanUpgrade &&
						cliqueInfo.DefenseBuildingLevel+1 <= cliqueInfo.CenterBuildingLevel &&
						cliqueInfo.DefenseBuildingCoins >= buildingDat.UpgradeCoins {

						cliqueInfo.DefenseBuildingCoins -= buildingDat.UpgradeCoins
						cliqueInfo.DefenseBuildingLevel += 1
						upgrade = true
					}
					playerCliqueBuilding.DonateCoinsDefenseBuilding += money
					out.TotalDonateCoins = int32(cliqueInfo.DefenseBuildingCoins)
					out.Level = cliqueInfo.DefenseBuildingLevel
				default:
					fail.When(true, fmt.Sprintf("undefind kongfu building %d", buildingId))
				}
				if upgrade {
					//升级就发动态
					upgradeLevelNotice(buildingId, cliqueInfo.Id, out.Level)
				}

				cliqueInfo.Contrib += money / clique_dat.CLIQUE_CONTRIB_RATE
				cliqueInfo.TotalDonateCoins += money

				agentDB.Update.GlobalClique(cliqueInfo)
				agentDB.Update.PlayerGlobalCliqueBuilding(playerCliqueBuilding)

				//捐赠统一接口
				out.PlayerDonateCoins = donateBuilding(agentDB, money, flowID)

				out.Result = clique_building_api.CLIQUE_BUILDING_DONATE_RESULT_SUCCESS

				//更新帮派每日任务 ，铜钱捐赠
				module.CliqueQuestRPC.UpDatePlayerCliqueDailyQuest(agentDB, clique_dat.CLIQUE_DAYLY_QUEST_DONATE)

				session.Send(out)
			})
		})
	})
}

func kongfuInfo(session *net.Session, buildingId int32) {
	kongfuBuildingCheck(buildingId)
	state := module.State(session)
	out := &clique_building_api.CliqueKongfuInfo_Out{}
	state.Database.Select.PlayerGlobalCliqueKongfu(func(row *mdb.PlayerGlobalCliqueKongfuRow) {
		kongfuDat := clique_dat.GetKongfuById(row.KongfuId())
		if kongfuDat.Building == buildingId {
			out.KongfuList = append(out.KongfuList, clique_building_api.CliqueKongfuInfo_Out_KongfuList{
				KongfuId: row.KongfuId(),
				Level:    row.Level(),
			})
		}
	})
	session.Send(out)

}

func kongfuTrain(session *net.Session, kongfuId int32) {

	kongfuDat := clique_dat.GetKongfuById(kongfuId)

	state := module.State(session)
	out := &clique_building_api.CliqueKongfuTrain_Out{
		KongfuId: kongfuId,
	}
	playerCliqueInfo := state.Database.Lookup.PlayerGlobalCliqueInfo(state.PlayerId)
	if playerCliqueInfo == nil || playerCliqueInfo.CliqueId <= 0 {
		out.Result = clique_building_api.CLIQUE_KONGFU_TRAIN_RESULT_NO_CLIQUE
		session.Send(out)
		return
	}

	var newKongfu bool
	var kongfu *mdb.PlayerGlobalCliqueKongfu
	state.Database.Select.PlayerGlobalCliqueKongfu(func(row *mdb.PlayerGlobalCliqueKongfuRow) {
		if row.KongfuId() == kongfuId {
			kongfu = row.GoObject()
			row.Break()
		}
	})
	if kongfu == nil {
		kongfu = &mdb.PlayerGlobalCliqueKongfu{
			Pid:      state.PlayerId,
			KongfuId: kongfuDat.Id,
			Level:    1,
		}
		newKongfu = true
	} else {
		out.Level = kongfu.Level
		kongfu.Level++
	}
	cliqueInfo := module.CliqueRPC.CliqueInfoLookUp(state.Database, playerCliqueInfo.CliqueId)
	switch kongfuDat.Building {
	case clique_dat.CLIQUE_BUILDING_HUICHUNTANG:
		//回春堂
		if kongfuDat.RequireBuildingLevel > cliqueInfo.HealthBuildingLevel ||
			kongfu.Level > cliqueInfo.HealthBuildingLevel*clique_dat.CLIQUE_KONGFU_LEVEL_LIIMT_FACTOR {

			out.Result = clique_building_api.CLIQUE_KONGFU_TRAIN_RESULT_MAX_LEVEL
			session.Send(out)
			return
		}
	case clique_dat.CLIQUE_BUILDING_SHENBINGTANG:
		//神兵堂
		if kongfuDat.RequireBuildingLevel > cliqueInfo.AttackBuildingLevel ||
			kongfu.Level > cliqueInfo.AttackBuildingLevel*clique_dat.CLIQUE_KONGFU_LEVEL_LIIMT_FACTOR {
			out.Result = clique_building_api.CLIQUE_KONGFU_TRAIN_RESULT_MAX_LEVEL
			session.Send(out)
			return
		}
	case clique_dat.CLIQUE_BUILDING_JINGANGTANG:
		//金刚堂
		if kongfuDat.RequireBuildingLevel > cliqueInfo.DefenseBuildingLevel ||
			kongfu.Level > cliqueInfo.DefenseBuildingLevel*clique_dat.CLIQUE_KONGFU_LEVEL_LIIMT_FACTOR {
			out.Result = clique_building_api.CLIQUE_KONGFU_TRAIN_RESULT_MAX_LEVEL
			session.Send(out)
			return
		}
	default:
	}
	//帮派贡献
	kongfuLevelUpConsume, ok := clique_dat.GetKongfuLevelConsume(kongfuId, kongfu.Level)
	if !ok {
		//没有配置
		out.Result = clique_building_api.CLIQUE_KONGFU_TRAIN_RESULT_MAX_LEVEL
		session.Send(out)
		return
	}
	if !module.CliqueRPC.DecPlayerCliqueContrib(state.Database, kongfuLevelUpConsume) {
		//贡献不足
		out.Result = clique_building_api.CLIQUE_KONGFU_TRAIN_RESULT_LACK_CONTRIB
		session.Send(out)
		return
	}

	if newKongfu {
		state.Database.Insert.PlayerGlobalCliqueKongfu(kongfu)
	} else {
		state.Database.Update.PlayerGlobalCliqueKongfu(kongfu)
	}
	out.Result = clique_building_api.CLIQUE_KONGFU_TRAIN_RESULT_SUCCESS
	out.Level = kongfu.Level
	session.Send(out)
	attribMap := map[int8]int32{}
	state.Database.Select.PlayerGlobalCliqueKongfu(func(row *mdb.PlayerGlobalCliqueKongfuRow) {
		kongfuDat := clique_dat.GetKongfuById(row.KongfuId())
		attribMap[kongfuDat.AttribType] = attribMap[kongfuDat.AttribType] + int32(row.Level())*kongfuDat.AttribValue
	})
	rpc.RemoteUpdateCliqueKongfuAttrib(state.PlayerId, attribMap)

}

func kongfuBuildingCheck(buildingId int32) {
	fail.When(buildingId != clique_dat.CLIQUE_BUILDING_HUICHUNTANG &&
		buildingId != clique_dat.CLIQUE_BUILDING_SHENBINGTANG &&
		buildingId != clique_dat.CLIQUE_BUILDING_JINGANGTANG, fmt.Sprintf("unsupport kongfu building %d", buildingId))
}
