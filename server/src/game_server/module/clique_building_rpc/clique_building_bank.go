package clique_building_rpc

// 帮派建筑物 钱庄

import (
	"core/fail"
	"core/i18l"
	"core/net"
	"core/time"
	"fmt"
	"game_server/api/protocol/clique_building_api"
	"game_server/dat/clique_dat"
	"game_server/dat/mail_dat"
	"game_server/dat/player_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/module/rpc"
	"game_server/tlog"
	"game_server/xdlog"
)

func cliqueBuildingBankDonate(session *net.Session, money int64) {
	state := module.State(session)
	out := &clique_building_api.CliqueBankDonate_Out{}
	playerCliqueInfo := state.Database.Lookup.PlayerGlobalCliqueInfo(state.PlayerId)
	if playerCliqueInfo == nil || playerCliqueInfo.CliqueId <= 0 {
		out.Result = clique_building_api.CLIQUE_BUILDING_DONATE_RESULT_FAILED
		session.Send(out)
		return
	}
	cliqueInfo := module.CliqueRPC.CliqueInfoLookUp(state.Database, playerCliqueInfo.CliqueId)

	// 帮派总舵信息
	baseInfo := clique_dat.GetCenterBuildingLevelInfo(cliqueInfo.CenterBuildingLevel)

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
	oldCliqueId := playerCliqueInfo.CliqueId
	rpc.RemoteDecMoney(state.PlayerId, money, player_dat.COINS, tlog.MFR_CLIQUE_DONATE_BANK, xdlog.ET_CLIQUE_DONATE, func() {
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(state.PlayerId, func(agentDB *mdb.Database) {
				playerCliqueInfo = agentDB.Lookup.PlayerGlobalCliqueInfo(agentDB.PlayerId())
				if playerCliqueInfo == nil || playerCliqueInfo.CliqueId <= 0 || playerCliqueInfo.CliqueId != oldCliqueId {
					return
				}
				cliqueInfo := module.CliqueRPC.CliqueInfoLookUp(agentDB, playerCliqueInfo.CliqueId)

				cliqueInfo.BankBuildingCoins += money

				// 钱庄模板信息
				bankInfo := clique_dat.GetCliqueBuildingBankInfo(cliqueInfo.BankBuildingLevel)

				// 判断是否需要升级总舵，以及更新玩家表
				if cliqueInfo.BankBuildingCoins >= bankInfo.Upgrade && cliqueInfo.BankBuildingLevel < clique_dat.CLIQUE_BUILDING_MAX_LEVEL && cliqueInfo.BankBuildingLevel < cliqueInfo.CenterBuildingLevel {
					cliqueInfo.BankBuildingLevel += 1
					cliqueInfo.BankBuildingCoins = cliqueInfo.BankBuildingCoins - bankInfo.Upgrade
					//升级就发动态
					upgradeLevelNotice(clique_dat.CLIQUE_BUILDING_QIANZHUANG, cliqueInfo.Id, cliqueInfo.BankBuildingLevel)
				}

				agentDB.Update.GlobalClique(cliqueInfo)

				playerCliqueBuilding := agentDB.Lookup.PlayerGlobalCliqueBuilding(state.PlayerId)
				playerCliqueBuilding.DonateCoinsBankBuilding += money
				agentDB.Update.PlayerGlobalCliqueBuilding(playerCliqueBuilding)

				out.PlayerDonateCoins = donateBuilding(agentDB, money, tlog.MFR_CLIQUE_DONATE_BANK)

				out.Result = clique_building_api.CLIQUE_BUILDING_DONATE_RESULT_SUCCESS
				out.CliqueBuildingBankDonateCoins = int32(cliqueInfo.BankBuildingCoins)
				out.CliqueBuildingBankLevel = cliqueInfo.BankBuildingLevel
				//更新帮派每日任务 ，铜钱捐赠
				module.CliqueQuestRPC.UpDatePlayerCliqueDailyQuest(agentDB, clique_dat.CLIQUE_DAYLY_QUEST_DONATE)
				session.Send(out)
			})
		})
	})
}

func cliqueBankBuy(session *net.Session, kind int8, num int8) {
	fail.When(kind != clique_dat.CLIQUE_BUILDING_BANK_GOLD && kind != clique_dat.CLIQUE_BUILDING_BANK_SILVER, "非法购买类型")
	fail.When(num <= 0, "非法购买数量")
	state := module.State(session)
	out := &clique_building_api.CliqueBankBuy_Out{}

	var cost int64
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
	bankDatInfo := clique_dat.GetCliqueBuildingBankInfo(cliqueInfo.BankBuildingLevel)
	if bankDatInfo == nil ||
		(kind == clique_dat.CLIQUE_BUILDING_BANK_GOLD && int16(num) > bankDatInfo.GoldCouponNum) ||
		(kind == clique_dat.CLIQUE_BUILDING_BANK_SILVER && int16(num) > bankDatInfo.SilverCouponNum) {
		session.Send(out)
		return
	}

	playerCliqueBuilding := state.Database.Lookup.PlayerGlobalCliqueBuilding(state.PlayerId)
	money_type := player_dat.COINS
	switch kind {
	case clique_dat.CLIQUE_BUILDING_BANK_GOLD:
		if playerCliqueBuilding.GoldExchangeNum > 0 {
			session.Send(out)
			return
		}
		cost = int64(num) * bankDatInfo.GoldCouponCoins
		money_type = player_dat.INGOT
	case clique_dat.CLIQUE_BUILDING_BANK_SILVER:
		if playerCliqueBuilding.SilverExchangeNum > 0 {
			session.Send(out)
			return
		}
		cost = int64(num) * bankDatInfo.SilverCouponCoins
	default:
		session.Send(out)
		return
	}
	rpc.RemoteDecMoney(state.PlayerId, cost, money_type, tlog.MFR_CLIQUE_DONATE_BANK_TRADE, xdlog.ET_CLIQUE_DONATE_BANK_TRADE, func() {
		// 检查玩家是否还在原来的帮派
		var tempCost int64
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(state.PlayerId, func(agentDB *mdb.Database) {
				playerCliqueInfo := agentDB.Lookup.PlayerGlobalCliqueInfo(state.PlayerId)
				if playerCliqueInfo == nil {
					session.Send(out)
					return
				}
				cliqueInfo := module.CliqueRPC.CliqueInfoLookUp(agentDB, playerCliqueInfo.CliqueId)
				if cliqueInfo == nil {
					session.Send(out)
					return
				}
				bankDatInfo := clique_dat.GetCliqueBuildingBankInfo(cliqueInfo.BankBuildingLevel)
				if bankDatInfo == nil ||
					(kind == clique_dat.CLIQUE_BUILDING_BANK_GOLD && int16(num) > bankDatInfo.GoldCouponNum) ||
					(kind == clique_dat.CLIQUE_BUILDING_BANK_SILVER && int16(num) > bankDatInfo.SilverCouponNum) {
					session.Send(out)
					return
				}
				playerCliqueBuilding := state.Database.Lookup.PlayerGlobalCliqueBuilding(state.PlayerId)
				temp_money_type := player_dat.COINS
				switch kind {
				case clique_dat.CLIQUE_BUILDING_BANK_GOLD:
					if playerCliqueBuilding.GoldExchangeNum > 0 {
						session.Send(out)
						return
					}
					tempCost = int64(num) * bankDatInfo.GoldCouponCoins
					temp_money_type = player_dat.INGOT
				case clique_dat.CLIQUE_BUILDING_BANK_SILVER:
					if playerCliqueBuilding.SilverExchangeNum > 0 {
						session.Send(out)
						return
					}
					tempCost = int64(num) * bankDatInfo.SilverCouponCoins
				default:
					session.Send(out)
					return
				}

				if tempCost != cost || temp_money_type != money_type {
					session.Send(out)
					return
				}

				if money_type == player_dat.INGOT {
					playerCliqueBuilding.GoldExchangeNum = int16(num)
					playerCliqueBuilding.GlodExchangeTime = time.GetNowTime()
				}
				if money_type == player_dat.COINS {
					playerCliqueBuilding.SilverExchangeNum = int16(num)
					playerCliqueBuilding.SilverExchangeTime = time.GetNowTime()
				}
				agentDB.Update.PlayerGlobalCliqueBuilding(playerCliqueBuilding)
				out.Success = true
				session.Send(out)
			})
		})
	})

}

func cliqueBankSold(session *net.Session, kind int8) {
	state := module.State(session)
	out := &clique_building_api.CliqueBankSold_Out{}
	playerCliqueInfo := state.Database.Lookup.PlayerGlobalCliqueInfo(state.PlayerId)
	if playerCliqueInfo == nil || playerCliqueInfo.CliqueId <= 0 {
		out.Result = clique_building_api.CLIQUE_BANK_SOLD_RESULT_NO_CLIQUE
		session.Send(out)
		return
	}
	cliqueInfo := module.CliqueRPC.CliqueInfoLookUp(state.Database, playerCliqueInfo.CliqueId)

	bankDatInfo := clique_dat.GetCliqueBuildingBankInfo(cliqueInfo.BankBuildingLevel)

	playerCliqueBuilding := state.Database.Lookup.PlayerGlobalCliqueBuilding(state.PlayerId)

	if kind == clique_dat.CLIQUE_BUILDING_BANK_GOLD && playerCliqueBuilding.GoldExchangeNum == 0 || (kind == clique_dat.CLIQUE_BUILDING_BANK_SILVER && playerCliqueBuilding.SilverExchangeNum == 0) {
		return
	}

	money_type := player_dat.COINS
	var price int64
	switch kind {
	case clique_dat.CLIQUE_BUILDING_BANK_GOLD:
		if time.GetNowTime()-playerCliqueBuilding.GlodExchangeTime < clique_dat.CLIQUE_BUILDING_BANK_SOLD_TIMESPAN {
			out.Result = clique_building_api.CLIQUE_BANK_SOLD_RESULT_CD
			session.Send(out)
			return
		}
		money_type = player_dat.INGOT
		price = int64(playerCliqueBuilding.GoldExchangeNum) * bankDatInfo.GoldCouponSold
	case clique_dat.CLIQUE_BUILDING_BANK_SILVER:
		if time.GetNowTime()-playerCliqueBuilding.SilverExchangeTime < clique_dat.CLIQUE_BUILDING_BANK_SOLD_TIMESPAN {
			out.Result = clique_building_api.CLIQUE_BANK_SOLD_RESULT_CD
			session.Send(out)
			return
		}
		price = int64(playerCliqueBuilding.SilverExchangeNum) * bankDatInfo.SilverCouponSold
	default:
		fail.When(true, "unknow type")
	}
	content1 := "您在帮派钱庄的投资得到了回报，连本带利共获得"
	var content2 string
	var attach []*mail_dat.Attachment
	if money_type == player_dat.INGOT {
		playerCliqueBuilding.GoldExchangeNum = 0
		playerCliqueBuilding.GlodExchangeTime = 0
		content2 = "元宝"
		attach = []*mail_dat.Attachment{&mail_dat.Attachment{
			AttachmentType: mail_dat.ATTACHMENT_INGOT,
			ItemNum:        price,
		}}
	}
	if money_type == player_dat.COINS {
		playerCliqueBuilding.SilverExchangeNum = 0
		playerCliqueBuilding.SilverExchangeTime = 0
		content2 = "铜钱"
		attach = []*mail_dat.Attachment{&mail_dat.Attachment{
			AttachmentType: mail_dat.ATTACHMENT_COINS,
			ItemNum:        price,
		}}
	}
	state.Database.Update.PlayerGlobalCliqueBuilding(playerCliqueBuilding)

	// 修改为通过邮件发送
	mail := &mail_dat.EmptyMail{
		Title:       i18l.T.Tran("帮派钱庄返利"),
		Content:     fmt.Sprintf("%s%d%s", i18l.T.Tran(content1), price, i18l.T.Tran(content2)),
		ExpireTime:  1,
		Attachments: attach,
	}
	rpc.RemoteMailSend(state.PlayerId, mail)
	out.Result = clique_building_api.CLIQUE_BANK_SOLD_RESULT_SUCCESS
	session.Send(out)
}
