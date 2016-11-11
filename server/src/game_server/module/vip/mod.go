package vip

import (
	"core/fail"
	"core/time"
	"game_server/api/protocol/notify_api"
	"game_server/api/protocol/vip_api"
	"game_server/dat/arena_award_box_dat"
	"game_server/dat/arena_buy_cost_config_dat"
	// "game_server/dat/event_dat"
	"fmt"
	"game_server/dat/event_dat"
	"game_server/dat/mail_dat"
	"game_server/dat/player_dat"
	"game_server/dat/quest_dat"
	"game_server/dat/rainbow_buy_cost_config_dat"
	"game_server/dat/rainbow_dat"
	"game_server/dat/vip_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/module/rpc"
	"game_server/tlog"
	"game_server/xdlog"
	gotime "time"
)

func init() {
	module.VIP = VIPMod{}
}

type VIPMod struct{}

func (mod VIPMod) UpdateIngot(db *mdb.Database, balance int64) {
	//module.Player.IncMoney(state.Database, state.MoneyState, balance, player_dat.INGOT, tlog.MFR_BUY_INGOT)

	playerVIPInfo := db.Lookup.PlayerVip(db.PlayerId())

	//元宝未到账，无需操作
	if playerVIPInfo.Ingot == balance {
		return
	}

	// var (
	// 	isOpen   bool
	// 	isGoOn   bool
	// 	isChange bool
	// )
	chargevalue := balance - playerVIPInfo.Ingot
	oldlevel := playerVIPInfo.Level
	//月卡
	// eventMonthCardInfo, ok := event_dat.GetEventInfoById(event_dat.EVENT_MONTH_CARD_AWARDS)
	// if ok && event_dat.CheckEventTime(eventMonthCardInfo, event_dat.NOT_END) {
	// 	monthcardIngot := module.Event.GetMonthCardPresentIngot()
	// 	if balance-playerVIPInfo.Ingot >= monthcardIngot {
	// 		//充值元宝数量达到条件
	// 		playerMonthCard := db.Lookup.PlayerMonthCardInfo(db.PlayerId())
	// 		start := time.GetDayFirstTime(gotime.Now(), 1)
	// 		if playerMonthCard == nil {
	// 			db.Insert.PlayerMonthCardInfo(&mdb.PlayerMonthCardInfo{
	// 				Pid:       db.PlayerId(),
	// 				Starttime: start,
	// 				Endtime:   start + 30*3600*24,
	// 				Buytimes:  1,
	// 			})
	// 			isOpen = true
	// 			isChange = true

	// 		} else if playerMonthCard.Endtime-start <= 3*3600*24 {
	// 			playerMonthCard.Buytimes++
	// 			if playerMonthCard.Starttime < start {
	// 				playerMonthCard.Starttime = start
	// 			}
	// 			if playerMonthCard.Endtime >= start {
	// 				playerMonthCard.Endtime += 30 * 3600 * 24
	// 			} else {
	// 				playerMonthCard.Endtime = start + 30*3600*24
	// 			}
	// 			db.Update.PlayerMonthCardInfo(playerMonthCard)
	// 			isGoOn = true
	// 			isChange = true
	// 		}

	// 		//发送通知
	// 		if isOpen || isGoOn || isChange {
	// 			if session, ok := module.Player.GetPlayerOnline(db.PlayerId()); ok {
	// 				if isOpen {
	// 					session.Send(&notify_api.NotifyMonthCardOpen_Out{})
	// 				}
	// 				if isGoOn {
	// 					session.Send(&notify_api.NotifyMonthCardRenewal_Out{})
	// 				}
	// 				if isChange {
	// 					session.Send(&notify_api.SendEventCenterChange_Out{})
	// 				}
	// 			}
	// 		}

	// 	}
	// }
	eventInGotRankConfig := event_dat.GetEventIngotRankConfig()
	inGotRankOpenTime := eventInGotRankConfig.StartUnixTime
	inGotRankEndTime := eventInGotRankConfig.EndUnixTime
	nowTime := time.GetNowTime()
	if (nowTime >= inGotRankOpenTime && nowTime <= inGotRankEndTime) && int64(balance-playerVIPInfo.Ingot) > 0 {
		updatePlayerAddInGotRecord(db, int64(balance-playerVIPInfo.Ingot), nowTime)
		module.Player.AddUpdateInGot(db, nowTime)
	}
	//累计充值活动
	module.Event.UpdateEventsIngot(db, balance-playerVIPInfo.Ingot, true)
	module.Event.UpdateJsonEventTotalRecharge(db, int32(balance-playerVIPInfo.Ingot))

	//充值就刷新首冲奖励表
	module.Event.UpdateFirstRecharge(db)
	//首冲奖励json
	module.Event.UpdateJsonFirstRecharge(db)

	if balance-playerVIPInfo.Ingot > 0 {
		if session, online := module.Player.GetPlayerOnline(db.PlayerId()); online {
			state := module.State(session)
			module.Quest.RefreshDailyQuest(state, quest_dat.DAILY_QUEST_CLASS_RECHARGE)
		}
	}

	//首次充值
	if playerVIPInfo.Ingot == 0 {
		if session, online := module.Player.GetPlayerOnline(db.PlayerId()); online {
			session.Send(&notify_api.NotifyFirstRechargeState_Out{
				NeverRecharge: false,
			})
			module.Notify.VIPLevelChange(session, playerVIPInfo.Level)
		}
		//  首充豪礼邮件
		module.Mail.SendMail(db, &mail_dat.MailShouChongHaoHuaLi{})

		now := gotime.Now()
		playerVIPInfo.CardId = fmt.Sprintf("%d%02d%02d", now.Year(), now.Month(), now.Day())
		//首充成为VIP后，若今日签到奖励有VIP双倍，并且玩家在成为VIP之前已经把签到了，那么把一份奖励通过邮件发送给玩家
		if module.DailySignIn.SignedToday(db) {
			award := module.DailySignIn.GetAwardForToday(db)
			if award.VIPDouble {
				module.DailySignIn.AwardByMail(db, award)
			}
		}

	}
	playerVIPInfo.Ingot = balance

	var vipLevleChanged bool
	oldLevel := playerVIPInfo.Level
	newLevelRequire, levelExist := vip_dat.GetVIPLevelInfo(playerVIPInfo.Level + 1)
	for levelExist && playerVIPInfo.Ingot+playerVIPInfo.PresentExp >= newLevelRequire {
		playerVIPInfo.Level += 1
		vipLevleChanged = true
		newLevelRequire, levelExist = vip_dat.GetVIPLevelInfo(playerVIPInfo.Level + 1)

		// vip升级奖励
		gifts := vip_dat.GetVIPLevelUpGift(playerVIPInfo.Level)
		if gifts != nil {
			mail := &mail_dat.MailXianZun{}
			mail.VipNum = fmt.Sprintf("%d", playerVIPInfo.Level)
			mail.Attachments = make([]*mail_dat.Attachment, len(gifts))
			for index, item := range gifts {
				if item.ItemId > 0 && item.ItemNum > 0 {
					mail.Attachments[index] = &mail_dat.Attachment{
						AttachmentType: mail_dat.ATTACHMENT_ITEM,
						ItemId:         item.ItemId,
						ItemNum:        int64(item.ItemNum),
					}
				}
			}
			rpc.RemoteMailSend(db.PlayerId(), mail)
		}
	}

	if vipLevleChanged {
		if session, online := module.Player.GetPlayerOnline(db.PlayerId()); online {
			module.Notify.VIPLevelChange(session, playerVIPInfo.Level)
		}
		//首次成为vip则补发爱心奖励邮件
		if oldLevel == 0 {
			rpc.RemoteMailSend(db.PlayerId(), mail_dat.MailVIPHeart{
				Attachments: []*mail_dat.Attachment{
					&mail_dat.Attachment{
						AttachmentType: mail_dat.ATTACHMENT_HEART,
						ItemId:         0,
						ItemNum:        int64(vip_dat.GetVIPPrivilegeTime(playerVIPInfo.Level, int32(vip_dat.AIXINFULI))),
					},
				},
			})
		}
		//rpc方式更新hd服上的玩家vip信息和总的vip人数
		if playerVIPInfo.PresentExp == 0 {
			rpc.RemoteUpdatePlayerVIPLevel(db.PlayerId(), int64(playerVIPInfo.Level))
		}
	}

	db.Update.PlayerVip(playerVIPInfo)
	xdlog.VipLog(db, int32(chargevalue), xdlog.ISPAY, int32(oldlevel), int32(playerVIPInfo.Level))
}

func (mod VIPMod) VIPInfo(db *mdb.Database) *mdb.PlayerVip {
	playerVIPInfo := db.Lookup.PlayerVip(db.PlayerId())
	return playerVIPInfo
}

//增加vip经验，暂时仅限idip
func (mod VIPMod) AddVipExp(db *mdb.Database, presentExp int32) {
	fail.When(presentExp <= 0, "vip exp must greate than zero")
	playerVIPInfo := db.Lookup.PlayerVip(db.PlayerId())
	oldviplevel := playerVIPInfo.Level
	playerVIPInfo.PresentExp += int64(presentExp)
	var vipLevleChanged bool
	newLevelRequire, levelExist := vip_dat.GetVIPLevelInfo(playerVIPInfo.Level + 1)
	for levelExist && playerVIPInfo.Ingot+playerVIPInfo.PresentExp >= newLevelRequire {
		playerVIPInfo.Level += 1
		vipLevleChanged = true
		newLevelRequire, levelExist = vip_dat.GetVIPLevelInfo(playerVIPInfo.Level + 1)
	}
	if vipLevleChanged {
		if session, online := module.Player.GetPlayerOnline(db.PlayerId()); online {
			module.Notify.VIPLevelChange(session, playerVIPInfo.Level)
		}
	}
	db.Update.PlayerVip(playerVIPInfo)
	xdlog.VipLog(db, presentExp, xdlog.NOTPAY, int32(oldviplevel), int32(playerVIPInfo.Level))
}

func (mod VIPMod) PrivilegeTimes(state *module.SessionState, privilege int16) int16 {
	playerVIPInfo := state.Database.Lookup.PlayerVip(state.PlayerId)
	return vip_dat.GetVIPPrivilegeTime(playerVIPInfo.Level, int32(privilege))
}

//如果当前VIP特权次数未0则检查非VIP用户是否有次数，有的话使用非VIP用户的次数。
//目前仅在爱心抽奖初使用。
func (mod VIPMod) GetPrivilegeTimesByDB(db *mdb.Database, privilege int16) int16 {
	playerVIPInfo := db.Lookup.PlayerVip(db.PlayerId())
	times := vip_dat.GetVIPPrivilegeTime(playerVIPInfo.Level, int32(privilege))
	if times <= 0 {
		times = vip_dat.GetVIPPrivilegeTime(0, int32(privilege))
	}
	return times
}

func (mod VIPMod) HavePrivileve(state *module.SessionState, privilege int16) bool {
	playerVIPInfo := state.Database.Lookup.PlayerVip(state.PlayerId)
	return vip_dat.HaveVIPPrivilege(playerVIPInfo.Level, int32(privilege))
}

func (mod VIPMod) CheckPrivilege(state *module.SessionState, privilege int16) {
	playerVIPInfo := state.Database.Lookup.PlayerVip(state.PlayerId)
	fail.When(!vip_dat.HaveVIPPrivilege(playerVIPInfo.Level, int32(privilege)), "no privilege")
}

func (mod VIPMod) IsVIP(db *mdb.Database) bool {
	vipInfo := module.VIP.VIPInfo(db)
	if vipInfo.Level > 0 {
		return true
	}
	return false
}

//vip购买对应功能次数
func buytimes(state *module.SessionState, privilege vip_api.BuyTimesType) bool {
	playerVIPInfo := state.Database.Lookup.PlayerVip(state.PlayerId)
	switch privilege {
	case vip_api.BUY_TIMES_TYPE_BIWUCHANGCISHU:
		playerArena := state.Database.Lookup.PlayerArena(state.PlayerId)
		fail.When(playerArena == nil || playerArena.DailyNum < arena_award_box_dat.MAX_DAILY_NUM, "daily free num not use up")
		times := vip_dat.GetVIPPrivilegeTime(playerVIPInfo.Level, vip_dat.BIWUCHANGCISHU)
		fail.When(times == 0 || times <= playerArena.BuyTimes, "can not buy times")
		costIngot := arena_buy_cost_config_dat.GetCost(int32(playerArena.BuyTimes))
		module.Player.DecMoney(state.Database, state.MoneyState, int64(costIngot), player_dat.INGOT, tlog.MFR_ARENA_TIMES_BUY, xdlog.ET_MONTH_CARD)
		playerArena.BuyTimes += 1
		state.Database.Update.PlayerArena(playerArena)
		return true
	case vip_api.BUY_TIMES_TYPE_RAINBOWSAODANG:
		playerRainbow := state.Database.Lookup.PlayerRainbowLevel(state.PlayerId)
		fail.When(playerRainbow == nil || playerRainbow.AutoFightNum < rainbow_dat.DAILY_AUTO_FIGHT_NUM, "daily free num not use up")
		times := vip_dat.GetVIPPrivilegeTime(playerVIPInfo.Level, vip_dat.CAIHONGTEQUAN)
		fail.When(times == 0 || times <= playerRainbow.BuyTimes, "can not buy times")
		costIngot := rainbow_buy_cost_config_dat.GetCost(int32(playerRainbow.BuyTimes))
		module.Player.DecMoney(state.Database, state.MoneyState, int64(costIngot), player_dat.INGOT, tlog.MFR_RAINBOW_TIMES_BUY, xdlog.ET_MONTH_CARD)
		playerRainbow.BuyTimes += 1
		playerRainbow.BuyTimestamp = time.GetNowTime()
		state.Database.Update.PlayerRainbowLevel(playerRainbow)
		return true
	default:
		fail.When(true, "unknow buy times type")
	}
	return false
}

func updatePlayerAddInGotRecord(db *mdb.Database, inGot, timestamp int64) {
	playerAddIngotRecord := db.Lookup.PlayerAddIngotRecord(db.PlayerId())
	if playerAddIngotRecord == nil {
		playerAddIngotRecord := &mdb.PlayerAddIngotRecord{}
		playerAddIngotRecord.Pid = db.PlayerId()
		playerAddIngotRecord.Ingot = inGot
		playerAddIngotRecord.Timestamp = timestamp
		db.Insert.PlayerAddIngotRecord(playerAddIngotRecord)
	} else {
		playerAddIngotRecord.Ingot += inGot
		playerAddIngotRecord.Timestamp = timestamp
		db.Update.PlayerAddIngotRecord(playerAddIngotRecord)
	}
}
