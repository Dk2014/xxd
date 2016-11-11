package event

import (
	"core/fail"
	"core/net"
	"core/time"
	"game_server/api/protocol/event_api"
	"game_server/dat/event_dat"
	"game_server/dat/mail_dat"
	"game_server/dat/player_dat"
	"game_server/dat/which_branch_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/module/event/json_handlers"
	"game_server/module/rpc"
	"game_server/tlog"
	"game_server/xdlog"
	gotime "time"
)

func getEvents(session *net.Session) {
	state := module.State(session)
	eventsInfo := event_dat.GetEventsInfo()
	out := &event_api.GetEvents_Out{}
	out.Events = make([]event_api.GetEvents_Out_Events, 0)
	module.CheckForEvents(session) // 增加每次点开活动中心都去检查活动状态
	for _, eventInfo := range eventsInfo {
		if !event_dat.CheckEventTime(eventInfo, event_dat.NOT_DISPOSE) {
			continue
		}
		var (
			nextAward   int32
			needAwarded bool
			isAppend    bool = true
		)
		eventRecord := state.EventsState.GetPlayerEventInfoById(eventInfo.Id)
		if eventRecord.MaxAward > eventRecord.Awarded {
			needAwarded = true
		}
		switch eventInfo.Id {
		case event_dat.EVENT_LEVEL_AWARD:
			if eventRecord.MaxAward == 0 {
				nowLevel := module.Role.GetMainRole(state.Database).Level
				module.Event.LevelUp(state.Database, 0, nowLevel)
			}
			nextAward = event_dat.GetNextLevel(eventRecord.Awarded)
			if nextAward <= eventRecord.Awarded {
				isAppend = false
			}
		case event_dat.EVENT_STRONG_AWARD:
			nextAward = event_dat.GetNextStrong(eventRecord.Awarded)
			if nextAward <= eventRecord.Awarded {
				isAppend = false
			}
		case event_dat.EVENT_ARENA_RANK_AWARDS:
			if !event_dat.CheckEventTime(eventInfo, event_dat.NOT_END) {
				//领奖阶段
				if eventRecord.Awarded == 0 {
					nextAward = event_dat.GetAwardRankByRank(eventRecord.MaxAward)
					needAwarded = true
					if nextAward == 0 {
						//没有获得奖励
						isAppend = false
						state.EventsState.UpdateAwarded(state.Database, eventInfo.Id, eventRecord.MaxAward)
					}
				} else {
					isAppend = false
				}
			} else {
				needAwarded = false
			}
		case event_dat.EVENT_PHYSICAL_AWARDS:
			if !time.IsInPointHour(player_dat.RESET_PLAYER_ACTIVITY_IN_HOUR, eventRecord.LastUpdated) {
				//隔天了 清空活跃度状态
				state.EventsState.ClearState(state.Database, eventInfo.Id)
				eventRecord = state.EventsState.GetPlayerEventInfoById(eventInfo.Id)
			}
			if eventRecord.MaxAward == player_dat.MAX_PLAYER_ACTIVITY+1 {
				nextAward = event_dat.GetMaxWeightInPhysical()
				needAwarded = false
			} else {
				nextAward = event_dat.GetNextPhysical(eventRecord.Awarded)
			}
		case event_dat.EVENT_MONTH_CARD_AWARDS:
			needAwarded = false
			if which_branch_dat.WHICH_BRANCH == which_branch_dat.TW {
				// 台湾月卡部分
				cardInfo := state.Database.Lookup.PlayerMonthCard(state.PlayerId)
				if cardInfo != nil && cardInfo.ExpireTimestamp > time.GetNowTime() && !time.IsToday(cardInfo.AwardTimestamp) {
					nextAward = event_dat.GetNextMonthCard(0)
					needAwarded = true
				}
			} else {
				monthCardInfo := state.Database.Lookup.PlayerMonthCardInfo(state.PlayerId)
				if monthCardInfo != nil && monthCardInfo.Endtime > time.GetNowTime() { // 有效的月卡用户
					monthCardRecord := state.Database.Lookup.PlayerMonthCardAwardRecord(state.PlayerId)
					nextAward = event_dat.GetNextMonthCard(eventRecord.Awarded)

					if monthCardRecord != nil && time.IsToday(monthCardRecord.LastUpdate) { // 已领取过奖励
						needAwarded = false
					}
				} else {
					//isAppend = false
					needAwarded = false
				}
			}

		case event_dat.EVENT_DINNER_AWARDS:
			if time.IsInRangeHour(event_dat.DINNER_START_IN_HOUR, event_dat.DINNER_END_IN_HOUR, eventRecord.LastUpdated) {
				isAppend = false
			} else {
				nextAward = 1 //午餐活动固定为1号奖励
				needAwarded = true
			}
		case event_dat.EVENT_SUPPER_AWARDS:
			if time.IsInRangeHour(event_dat.SUPPER_START_IN_HOUR, event_dat.SUPPER_END_IN_HOUR, eventRecord.LastUpdated) {
				isAppend = false
			} else {
				nextAward = 1 //晚餐活动固定为1号奖励
				needAwarded = true
			}
		case event_dat.EVENT_QQVIP_GIFT_AWARDS:
			// if state.TencentState.IsQQVIP {
			// 	//上次操作时间跟今天不同
			// 	gap_days := int32(time.GetNowDay()) - int32(time.GetNowDayFromUnix(eventRecord.LastUpdated))
			// 	if gap_days > 0 {
			// 		state.EventsState.UpdateMax(state.Database, eventRecord.EventId, eventRecord.MaxAward+gap_days)
			// 	}
			// 	nextAward, _ = event_dat.GetEventQQVIPGiftAward(eventRecord.MaxAward)
			// } else {
			// 	isAppend = false
			// }
			// 注 qq会员活动不在活动中心显示
			isAppend = false
		case event_dat.EVENT_TOTAL_LOGIN:
			day_zero := time.GetTodayZero()
			if day_zero > eventRecord.LastUpdated || eventRecord.MaxAward == 0 {
				if eventRecord.MaxAward == 0 || eventRecord.LastUpdated > day_zero-24*3600 {
					//连续登录 天数加1
					state.EventsState.UpdateMax(state.Database, eventInfo.Id, eventRecord.MaxAward+1)
				} else if playerInfo := state.Database.Lookup.PlayerInfo(state.PlayerId); playerInfo.LastOfflineTime > day_zero-24*3600 {
					//上次是跨天玩游戏
					gap_days := gotime.Unix(playerInfo.LastOfflineTime, 0).Day() - gotime.Now().Day()
					state.EventsState.UpdateMax(state.Database, eventInfo.Id, eventRecord.MaxAward+int32(gap_days))
				} else {
					//隔天登录清空累计登录天数
					state.EventsState.List[eventInfo.Id] = &module.EventInfo{
						EventId:     eventInfo.Id,
						EndUnixTime: eventInfo.End,
						Awarded:     0,
						MaxAward:    1,
						LastUpdated: time.GetNowTime(),
					}
					state.Database.Update.PlayerEventAwardRecord(&mdb.PlayerEventAwardRecord{
						Pid:             state.PlayerId,
						RecordBytes:     state.EventsState.Encode(),
						JsonEventRecord: state.JsonEventsState.Encode(),
					})

				}
				eventRecord = state.EventsState.GetPlayerEventInfoById(eventInfo.Id)
				needAwarded = true
			}
			//上次操作时间跟今天不同
			// gap_days := int32(time.GetNowDay()) - int32(time.GetNowDayFromUnix(eventRecord.LastUpdated))
			// if gap_days > 0 {
			// 	state.EventsState.UpdateMax(state.Database, eventRecord.EventId, eventRecord.MaxAward+gap_days)
			// }
			var awardContent *event_dat.EventDefaultAward
			if needAwarded {
				nextAward, awardContent = event_dat.GetEventTotalLoginAward(eventRecord.MaxAward)
			} else {
				nextAward, awardContent = event_dat.GetEventTotalLoginAward(eventRecord.MaxAward + 1)
			}
			if awardContent == nil { // 累计登陆奖励天数从大于1的开始
				needAwarded = false
			}
		case event_dat.EVENT_QQVIP_ADDITION:
			isAppend = false
		case event_dat.EVENT_VIP_CLUB:
			nextAward = eventRecord.Awarded
			needAwarded = false
		case event_dat.EVENT_SEVEN_DAY_AWARDS:
			if eventRecord.LastUpdated == 0 {
				isAppend = false
			} else {
				gap := int32(time.GetNowDay()) - int32(time.GetNowDayFromUnix(eventRecord.LastUpdated)) + 1
				if gap > event_dat.GetMaxWeightInSevenDay() {
					isAppend = false
				} else {
					nextAward = gap
					if eventRecord.Awarded >= 1<<uint(gap-1) {
						needAwarded = false
					}
				}
			}
		case event_dat.EVENT_SHARE_AWARDS:
			nextAward = eventRecord.Awarded
			nextLevel := event_dat.GetNextShare(eventRecord.Awarded)
			if nextLevel == 0 {
				isAppend = false
			} else if eventRecord.MaxAward >= nextLevel {
				needAwarded = true
			} else {
				needAwarded = false
			}
		case event_dat.EVENT_RICHMAN_CLUB:
			nextAward = eventRecord.Awarded
			needAwarded = false
		case event_dat.EVENT_GROUP_BUY:
			if event_dat.CheckEventTime(eventInfo, event_dat.NOT_END) {
				//认购阶段
				nextAward = eventRecord.MaxAward
				if nextAward == 0 {
					needAwarded = true
				} else {
					needAwarded = false
				}
			} else {
				//结算阶段
				if eventRecord.MaxAward == 0 {
					//没有参与认购
					isAppend = false
				} else if eventRecord.Awarded == 0 {
					//认购了但没有结算
					needAwarded = true
					nextAward = eventRecord.MaxAward
				} else {
					//认购且结算完毕
					isAppend = false
				}
			}
		case event_dat.EVENT_BUY_PARTNER:
			if eventRecord.Awarded == 0 {
				nextAward = 1
				needAwarded = false
			} else {
				isAppend = false
			}
		case event_dat.EVENT_TOTAL_RECHARGE:
			if eventRecord.Awarded == (1<<uint(event_dat.CountEventTotalRecharge()) - 1) {
				isAppend = false
			} else {
				nextAward = eventRecord.Awarded
				_, level := event_dat.GetEventTotalRechargeLevel(int32(module.Event.GetEventsIngot(state.Database, true)))
				if level >= 0 && eventRecord.Awarded != (1<<uint(level+1)-1) {
					needAwarded = true
				} else {
					needAwarded = false
				}
			}

		case event_dat.EVENT_TEN_DRAW:
			nextAward = eventRecord.Awarded
			nextLevel := event_dat.GetNextTenDraw(eventRecord.Awarded)
			if nextLevel == 0 {
				isAppend = false
			} else if eventRecord.MaxAward >= nextLevel {
				needAwarded = true
			} else {
				needAwarded = false
			}
		case event_dat.EVENT_TOTAL_CONSUME:
			nextAward = eventRecord.Awarded
			if eventRecord.Awarded < int32(module.Event.GetEventsIngot(state.Database, false)/int64(event_dat.GetTotalConsumeRadix())) {
				needAwarded = true
			}
		case event_dat.EVENT_FIRST_RECHARGE:
			vipInfo := state.Database.Lookup.PlayerVip(state.PlayerId)
			if vipInfo.Ingot > 0 {
				continue
			}

		case event_dat.EVENT_FIRST_RECHARGE_DAILY:

			nextAward = int32(event_dat.EventIntervalDays(eventInfo.Id))

			//得不到奖励索引，无法显示奖品
			if nextAward == event_dat.INVALIDINDEX || event_dat.FirstRechargeDailyReward(int(nextAward)) == nil {
				isAppend = false
				break
			}

			//今天充值了
			if time.IsToday(eventRecord.LastUpdated) {
				//没有领取
				if eventRecord.Awarded != nextAward {
					needAwarded = true
				}
			}
		default:
			if handler, ok := module.ExtraEvents[eventInfo.Id]; ok {
				isAppend, nextAward, _, needAwarded = handler.GetEventStatus()
			}
		}

		if isAppend == true {
			out.Events = append(out.Events, event_api.GetEvents_Out_Events{
				EventId: eventInfo.Id,
				Process: nextAward,
				IsAward: needAwarded,
			})
		}
	}

	// 处理json文件中配置的活动
	json_handlers.HandleJsonEventsStatus(session, out)

	session.Send(out)
}

func getEventAward(session *net.Session, eventId int16, param1, param2, param3 int32) {
	var newAward int32 //领奖励之后得下一阶段奖励
	var status int32   //奖励领取到得状态
	var awardeContent *event_dat.EventDefaultAward

	state := module.State(session)
	db := state.Database
	now := time.GetNowTime()
	eventInfo, ok := event_dat.GetEventInfoById(eventId)
	out := &event_api.GetEventAward_Out{}
	out.Result = -1
	//验证
	fail.When(!ok, "event is not found")
	if !event_dat.CheckEventTime(eventInfo, event_dat.NOT_DISPOSE) {
		out.Result = 3
		session.Send(out)
		return
	}

	eventRecord := state.EventsState.GetPlayerEventInfoById(eventId)
	var xdEventType int32 = 0
	//要更新领取奖励得有等级奖励活动和战力奖励活动
	switch eventId {
	case event_dat.EVENT_LEVEL_AWARD:
		var awards []*mail_dat.Attachment
		status, newAward, awards = event_dat.GetPlayerLevelAward(eventRecord.Awarded)
		if status == 0 || status > int32(module.Role.GetMainRole(state.Database).Level) {
			out.Result = 3
			session.Send(out)
			return
		}

		//获取奖励
		for _, attachment := range awards {
			switch attachment.AttachmentType {
			case mail_dat.ATTACHMENT_ITEM:
				//普通物品
				module.Item.AddItem(state.Database, attachment.ItemId, int16(attachment.ItemNum), tlog.IFR_EVENT_CENTER, xdlog.ET_EVENT_CENTER_LEVEL_AWARD, "")
			case mail_dat.ATTACHMENT_COINS:
				//铜钱
				module.Player.IncMoney(state.Database, state.MoneyState, attachment.ItemNum, player_dat.COINS, tlog.MFR_EVENT_CENTER, xdlog.ET_EVENT_CENTER_LEVEL_AWARD, "")
			case mail_dat.ATTACHMENT_INGOT:
				// 元宝
				module.Player.IncMoney(state.Database, state.MoneyState, attachment.ItemNum, player_dat.INGOT, tlog.MFR_EVENT_CENTER, xdlog.ET_EVENT_CENTER_LEVEL_AWARD, "")
			case mail_dat.ATTACHMENT_HEART:
				// 爱心
				module.Heart.IncHeart(state, int16(attachment.ItemNum))
			}
		}
		tlog.PlayerSystemModuleFlowLog(db, tlog.SMT_LEVEL_EVENT)
	case event_dat.EVENT_STRONG_AWARD:
		var award *event_dat.EventFightPowerAward
		status, newAward, award = event_dat.GetPlayerStrongAward(eventRecord.Awarded, eventRecord.MaxAward)
		if status == 0 || status > state.Database.Lookup.PlayerFightNum(state.PlayerId).FightNum {
			out.Result = 3
			session.Send(out)
			return
		}

		//addAwardContents(state, &award.EventDefaultAward)
		awardeContent = &award.EventDefaultAward
		xdEventType = xdlog.ET_EVENT_CENTER_STRONG_AWARD
		tlog.PlayerSystemModuleFlowLog(db, tlog.SMT_STRONG_EVENT)
	case event_dat.EVENT_ARENA_RANK_AWARDS:

		if eventRecord.Awarded == 0 || (eventRecord.Awarded > 0 && !event_dat.IsInPointEventTime(eventInfo, event_dat.NOT_DISPOSE, eventRecord.LastUpdated)) {
			if eventRecord.Awarded > 0 || eventRecord.MaxAward == 0 {
				out.Result = 3
				session.Send(out)
				return
			}
			award_rank := event_dat.GetAwardRankByRank(eventRecord.MaxAward)
			award, ok := event_dat.GetEventArenaRankAward(award_rank)
			if !ok {
				out.Result = 3
				session.Send(out)
				return
			}
			status = eventRecord.MaxAward
			awardeContent = award
			//addAwardContents(state, award)
			xdEventType = xdlog.ET_EVENT_CENTER_ARENA_RANK_AWARDS
		} else {
			out.Result = 3
			session.Send(out)
			return
		}
		tlog.PlayerSystemModuleFlowLog(db, tlog.SMT_ARENA_EVENT)
	case event_dat.EVENT_PHYSICAL_AWARDS:
		var award *event_dat.EventDefaultAward
		status, newAward, award = event_dat.GetPlayerPhysicalAward(eventRecord.Awarded)
		if status == 0 {
			out.Result = 3
			session.Send(out)
			return
		}
		if nowStatus, _ := module.GetPlayerActivity(state.Database); nowStatus < status {
			out.Result = 3
			session.Send(out)
			return
		}
		if newAward == 0 {
			eventRecord.MaxAward = player_dat.MAX_PLAYER_ACTIVITY + 1
			newAward = event_dat.GetMaxWeightInPhysical()
			out.Result = 0
		}
		//addAwardContents(state, award)
		awardeContent = award
		xdEventType = xdlog.ET_EVENT_CENTER_PHYSICAL_AWARDS
		tlog.PlayerSystemModuleFlowLog(db, tlog.SMT_PHYSICAL_EVENT)
	case event_dat.EVENT_MONTH_CARD_AWARDS:
		if which_branch_dat.WHICH_BRANCH == which_branch_dat.TW {
			//台湾月卡领取逻辑
			_, _, award := event_dat.GetPlayerMonthCardAward(0, 0)
			monthCardInfo := state.Database.Lookup.PlayerMonthCard(state.PlayerId)
			fail.When(monthCardInfo == nil || monthCardInfo.ExpireTimestamp < time.GetNowTime() || time.IsToday(monthCardInfo.AwardTimestamp), "monthcard award is not avalid")
			monthCardInfo.AwardTimestamp = time.GetNowTime()
			state.Database.Update.PlayerMonthCard(monthCardInfo)
			awardeContent = award
			xdEventType = xdlog.ET_MONTH_CARD
			out.Result = 0
			newAward = 0
			status = 0 // 不更新玩家月卡数据记录，月卡只有一个返利
			tlog.PlayerSystemModuleFlowLog(db, tlog.SMT_MONTCH_CARD_EVENT)
		} else {
			var award *event_dat.EventDefaultAward
			status, newAward, award = event_dat.GetPlayerMonthCardAward(0, 0) //BUG!!!!!(eventRecord.Awarded, eventRecord.MaxAward)
			fail.When(status == 0, "no award need to get")
			monthCardInfo := state.Database.Lookup.PlayerMonthCardInfo(state.PlayerId)
			fail.When(monthCardInfo == nil, "no tss info")
			ensureMonthCard := false
			row := state.Database.Lookup.PlayerMonthCardAwardRecord(state.PlayerId)
			fail.When(row != nil && time.IsToday(row.LastUpdate), "month card has recevied today")
			if now <= int64(monthCardInfo.Endtime) && now >= int64(monthCardInfo.Starttime) {
				ensureMonthCard = true
			}
			fail.When(!ensureMonthCard, "month card expire")
			//addAwardContents(state, award)
			awardeContent = award
			xdEventType = xdlog.ET_MONTH_CARD
			if row == nil { //之前未领取过
				state.Database.Insert.PlayerMonthCardAwardRecord(&mdb.PlayerMonthCardAwardRecord{
					Pid:        state.PlayerId,
					LastUpdate: now,
				})
			} else { //之前领取过
				row.LastUpdate = now
				state.Database.Update.PlayerMonthCardAwardRecord(row)
			}
			out.Result = 0
			newAward = 0
			status = 0 // 不更新玩家月卡数据记录，月卡只有一个返利
			tlog.PlayerSystemModuleFlowLog(db, tlog.SMT_MONTCH_CARD_EVENT)
		}
	case event_dat.EVENT_DINNER_AWARDS:
		if eventRecord.LastUpdated == 0 || (time.IsInRangeHour(event_dat.DINNER_START_IN_HOUR, event_dat.DINNER_END_IN_HOUR, time.GetNowTime()) && eventRecord.LastUpdated < time.GetTodayZero()) {
			dayLevel := gotime.Now().Weekday()
			if dayLevel == gotime.Sunday {
				dayLevel = 7 //星期天修整为第7天奖励
			}
			award, ok := event_dat.GetEventDinnerAward(int32(dayLevel)) //根据星期几获取奖励
			fail.When(!ok, "午餐奖励配置错误")
			//addAwardContents(state, award)
			awardeContent = award
			xdEventType = xdlog.ET_EVENT_CENTER_DINNER_AWARDS
			state.EventsState.UpdateAwardedTime(state.Database, eventId)
		}
		newAward = 0
		tlog.PlayerSystemModuleFlowLog(db, tlog.SMT_DINNER_EVENT)
	case event_dat.EVENT_SUPPER_AWARDS:
		if eventRecord.LastUpdated == 0 || (time.IsInRangeHour(event_dat.SUPPER_START_IN_HOUR, event_dat.SUPPER_END_IN_HOUR, time.GetNowTime()) && eventRecord.LastUpdated < time.GetTodayZero()) {
			dayLevel := gotime.Now().Weekday()
			if dayLevel == gotime.Sunday {
				dayLevel = 7 //星期天修整为第7天奖励
			}
			award, ok := event_dat.GetEventSupperAward(int32(dayLevel)) //根据星期几获取奖励
			fail.When(!ok, "晚餐奖励配置错误")
			//addAwardContents(state, award)
			awardeContent = award
			xdEventType = xdlog.ET_EVENT_CENTER_SUPPER_AWARDS
			state.EventsState.UpdateAwardedTime(state.Database, eventId)
		}
		newAward = 0
		tlog.PlayerSystemModuleFlowLog(db, tlog.SMT_SUPPER_EVENT)
	case event_dat.EVENT_QQVIP_GIFT_AWARDS:
		if (state.TencentState.QqVipStatus & 1) > 0 {
			fail.When(eventRecord.Awarded >= eventRecord.MaxAward, "player event record error")
			_, award := event_dat.GetEventQQVIPGiftAward(eventRecord.MaxAward)
			status = eventRecord.MaxAward
			//addAwardContents(state, award)
			awardeContent = award
			xdEventType = xdlog.ET_EVENT_CENTER_QQVIP_GIFT_AWARDS
			newAward = eventRecord.MaxAward + 1 //领完即不能领取了
		} else {
			fail.When(true, "not qq normal vip")
		}
		tlog.PlayerSystemModuleFlowLog(db, tlog.SMT_QQVIP_GIFT)
	case event_dat.EVENT_TOTAL_LOGIN:
		if eventRecord.Awarded >= eventRecord.MaxAward {
			out.Result = 3
			session.Send(out)
			return
		}
		_, award := event_dat.GetEventTotalLoginAward(eventRecord.MaxAward)
		status = eventRecord.MaxAward
		//addAwardContents(state, award)
		awardeContent = award
		xdEventType = xdlog.ET_EVENT_CENTER_TOTAL_LOGIN
		newAward, _ = event_dat.GetEventTotalLoginAward(eventRecord.MaxAward + 1)
		out.Result = 0 //领完即不能领取了
		tlog.PlayerSystemModuleFlowLog(db, tlog.SMT_TOTAL_LOGIN_EVENT)
	case event_dat.EVENT_VIP_CLUB:
		out := &event_api.GetEventAward_Out{}
		if state.Database.Lookup.PlayerVip(state.PlayerId).Level == 0 {
			//不是会员
			out.Result = 3
			session.Send(out)
			return
		}
		//rpc验证是否可领取
		rpc.RemoteGetGlobalVipCount(func(vipTotal int64) {

			if _, ok := state.EventsState.List[eventId]; !ok {
				state.EventsState.List[eventId] = new(module.EventInfo)
				state.EventsState.List[eventId].EventId = eventId
			}
			newAward = event_dat.GetNextVipCount(eventRecord.Awarded)
			var next int64
			if newAward == 0 || int64(newAward) > vipTotal {
				next = 0
			} else {
				award, _ := event_dat.GetEventVipClubAward(newAward)
				addAwardContents(state, award, xdlog.ET_EVENT_CENTER_VIP_CLUB)
				out.Award = newAward
				mdb.GlobalExecute(func(globalDB *mdb.Database) {
					globalDB.AgentExecute(state.PlayerId, func(agenDB *mdb.Database) {
						state.EventsState.AddEventAwardState(agenDB, eventId, eventRecord.MaxAward, newAward)
					})
				})
				next = int64(event_dat.GetNextVipCount(newAward))
			}

			if next == 0 { //已全部领完
				out.Result = 2
			} else if next <= vipTotal {
				out.Result = 1
			} else {
				out.Result = 0
			}
			session.Send(out)
		})
		tlog.PlayerSystemModuleFlowLog(db, tlog.SMT_VIP_CLUB_EVENT)
		return
	case event_dat.EVENT_SEVEN_DAY_AWARDS:
		return //七天乐活动领奖单独处理
	case event_dat.EVENT_RICHMAN_CLUB:
		return //土豪俱乐部活动已单独处理
	case event_dat.EVENT_SHARE_AWARDS:
		newAward = event_dat.GetNextShare(eventRecord.Awarded)
		if newAward == 0 || newAward > eventRecord.MaxAward {
			out.Result = 3
			session.Send(out)
			return
		}
		award, _ := event_dat.GetEventShareAward(newAward)
		//addAwardContents(state, award)
		awardeContent = award
		xdEventType = xdlog.ET_EVENT_CENTER_SHARE_AWARDS
		status = newAward
		nextAward := event_dat.GetNextShare(newAward)
		if nextAward == 0 {
			out.Result = 2
			newAward = event_dat.MAX_SHARE_TIMES + 1
			status = newAward
		} else if nextAward <= eventRecord.MaxAward {
			out.Result = 1
		} else {
			out.Result = 0
		}
		tlog.PlayerSystemModuleFlowLog(db, tlog.SMT_SHARE_EVENT)
	case event_dat.EVENT_GROUP_BUY:
		if event_dat.CheckEventTime(eventInfo, event_dat.NOT_END) {
			//认购阶段
			if eventRecord.MaxAward > 0 {
				out.Result = 3
				session.Send(out)
				return
			}
			rpc.RemoteOperateGroupBuyCount(true, func(count int32) {
				cost := event_dat.GetCostByGroupBuyCount(count - 1)
				groupBuyInfo := event_dat.GetEventGroupBuyInfo()
				mdb.GlobalExecute(func(globalDB *mdb.Database) {
					globalDB.AgentExecute(state.PlayerId, func(agenDB *mdb.Database) {
						module.Player.DecMoney(agenDB, state.MoneyState, int64(cost), player_dat.INGOT, tlog.MFR_EVENT_CENTER, xdlog.ET_EVENT_CENTER) //扣除元宝
						module.Item.AddItem(agenDB, groupBuyInfo.ItemId, 1, tlog.IFR_EVENT_CENTER, xdlog.ET_EVENT_CENTER, "")                         //添加物品
						state.EventsState.AddEventAwardState(agenDB, eventId, cost, 0)
					})
				})
				out.Result = 0
				out.Award = cost
				session.Send(out)
			})

		} else {
			//结算阶段
			if eventRecord.Awarded > 0 || eventRecord.MaxAward == 0 {
				out.Result = 3
				session.Send(out)
				return
			}
			rpc.RemoteOperateGroupBuyCount(false, func(count int32) {
				cost := event_dat.GetCostByGroupBuyCount(count)
				oldCost := eventRecord.MaxAward
				differ := oldCost - cost
				mdb.GlobalExecute(func(globalDB *mdb.Database) {
					globalDB.AgentExecute(state.PlayerId, func(agenDB *mdb.Database) {
						state.EventsState.UpdateAwarded(agenDB, eventId, eventRecord.MaxAward)
					})
				})
				out.Result = 2
				if differ > 0 {
					out.Award = differ
					mdb.GlobalExecute(func(globalDB *mdb.Database) {
						globalDB.AgentExecute(state.PlayerId, func(agenDB *mdb.Database) {
							module.Player.IncMoney(agenDB, state.MoneyState, int64(differ), player_dat.INGOT, tlog.MFR_EVENT_CENTER, xdlog.ET_EVENT_CENTER, "") //返还元宝
						})
					})
				}
				session.Send(out)
			})
		}
		tlog.PlayerSystemModuleFlowLog(db, tlog.SMT_GROUP_BUY_EVENT)
		return
	case event_dat.EVENT_BUY_PARTNER:

		// 1.检查团队中是否已有此人
		patnerid, buddyLevel, costs := event_dat.GetPartneInfo()
		patnerExist := module.Role.GetBuddyRole(state.Database, int8(patnerid))

		if patnerExist != nil {
			out.Result = 0
			break
		}

		// 2.检查是否有过买卖记录
		if eventRecord.Awarded != 0 {
			//已经买卖过
			out.Result = 3
			break
		}
		// 3.扣元宝

		module.Player.DecMoney(state.Database, state.MoneyState, costs, player_dat.INGOT, tlog.MFR_EVENT_CENTER, xdlog.ET_EVENT_CENTER)

		// 4.插入购买记录
		state.EventsState.AddEventAwardState(state.Database, eventId, 1, 1)
		// 5.增加角色
		module.Role.AddBuddyRole(state, int8(patnerid), int16(buddyLevel))
		out.Result = 2
		// 6.发邮件送装备
		var attachments []*mail_dat.Attachment
		attachments = append(attachments, &mail_dat.Attachment{AttachmentType: 0,
			ItemId:  418,
			ItemNum: 1})
		attachments = append(attachments, &mail_dat.Attachment{AttachmentType: 0,
			ItemId:  423,
			ItemNum: 10})

		module.Mail.SendMail(state.Database, &mail_dat.EmptyMail{
			MailId:      0,
			Title:       "胧月参上",
			Content:     "胧月很高兴能成为您的伙伴与您共闯仙侠道，在此为您奉上胧月的专属武器及魂侍，赶快装备起来让胧月变得更强吧！",
			Parameters:  "",
			Attachments: attachments,
			ExpireTime:  1,
			Priority:    0,
		})
		tlog.PlayerSystemModuleFlowLog(db, tlog.SMT_BUY_PARTNER_EVENT)
	case event_dat.EVENT_TEN_DRAW:
		newAward = event_dat.GetNextTenDraw(eventRecord.Awarded)
		if newAward == 0 || newAward > eventRecord.MaxAward {
			out.Result = 3
			session.Send(out)
			return
		}
		award, _ := event_dat.GetEventTenDrawAward(newAward)
		//addAwardContents(state, award)
		awardeContent = award
		xdEventType = xdlog.ET_EVENT_CENTER_TEN_DRAW
		status = newAward
		nextAward := event_dat.GetNextTenDraw(newAward)
		if nextAward == 0 {
			out.Result = 2
			newAward = event_dat.GetMaxTenDrawAwardTimes() + 1
			status = newAward
		} else if nextAward <= eventRecord.MaxAward {
			out.Result = 1
		} else {
			out.Result = 0
		}
		tlog.PlayerSystemModuleFlowLog(db, tlog.SMT_TEN_DRAW_EVENT)
	case event_dat.EVENT_TOTAL_CONSUME:
		canAwardTimes := int32(module.Event.GetEventsIngot(state.Database, false) / int64(event_dat.GetTotalConsumeRadix()))
		if canAwardTimes <= eventRecord.Awarded {
			out.Result = 3
			session.Send(out)
			return
		}
		if canAwardTimes-eventRecord.Awarded == 1 {
			out.Result = 0
		} else {
			out.Result = 1
		}
		//addAwardContents(state, event_dat.GetTotalConsumeAward())
		awardeContent = event_dat.GetTotalConsumeAward()
		eventRecord = state.EventsState.AddEventAwardState(state.Database, eventId, eventRecord.MaxAward, eventRecord.Awarded+1)
		newAward = eventRecord.Awarded
		tlog.PlayerSystemModuleFlowLog(db, tlog.SMT_TOTAL_CONSUME_EVENT)

	case event_dat.EVENT_FIRST_RECHARGE_DAILY:
		//跨天不能领奖
		if !time.IsToday(eventRecord.LastUpdated) {
			out.Result = 3
			session.Send(out)
			return
		}
		//获取天数索引
		day := event_dat.EventIntervalDays(eventId)
		awarded := event_dat.FirstRechargeDailyReward(day)
		if eventRecord.Awarded == int32(awarded.RequireDay) || day == event_dat.INVALIDINDEX {
			//已经领取
			out.Result = 3
			session.Send(out)
			return
		}

		//addAwardContents(state, awarded.EventAward)
		awardeContent = awarded.EventAward
		xdEventType = xdlog.ET_EVENT_CENTER_FIRST_RECHARGE_DAILY

		status = int32(awarded.RequireDay)
		newAward = status
		out.Result = event_dat.CONTINUE //继续领取奖励
		tlog.PlayerSystemModuleFlowLog(db, tlog.SMT_FRIST_RECHARGE_DAILY_EVENT)

	default:
		if handler, ok := module.ExtraEvents[eventInfo.Id]; ok {
			handler.GetEventAward(param1, param2, param3)
			return
		}
		fail.When(true, "wrong eventId")
	}
	//更新状态
	if status > 0 {
		state.EventsState.UpdateAwarded(state.Database, eventId, status)
	}

	if awardeContent != nil {
		// 将领奖操作放到更新状态之后
		addAwardContents(state, awardeContent, xdEventType)
	}

	if out.Result == -1 { //-1表示没有被预先处理过result
		if newAward == 0 {
			out.Result = 2 //活动已结束
		} else if eventRecord.MaxAward >= newAward {
			out.Result = 1 //继续领取奖励
		} else {
			out.Result = 0 //无法领取奖励，活动继续
		}
	}
	out.Award = newAward
	session.Send(out)
}

func addAwardContents(state *module.SessionState, award *event_dat.EventDefaultAward, xdEventType int32) {
	if award.Ingot > 0 {
		module.Player.IncMoney(state.Database, state.MoneyState, int64(award.Ingot), player_dat.INGOT, tlog.MFR_EVENT_CENTER, xdEventType, "")
	}
	if award.Coin > 0 {
		module.Player.IncMoney(state.Database, state.MoneyState, int64(award.Coin), player_dat.COINS, tlog.MFR_EVENT_CENTER, xdEventType, "")
	}
	if award.Heart > 0 {
		module.Heart.IncHeart(state, award.Heart)
	}
	table := map[int16]int16{
		award.Item1Id: award.Item1Num,
		award.Item2Id: award.Item2Num,
		award.Item3Id: award.Item3Num,
		award.Item4Id: award.Item4Num,
		award.Item5Id: award.Item5Num,
	}
	for id, n := range table {
		if id > 0 && n > 0 {
			module.Item.AddItem(state.Database, id, int16(n), tlog.IFR_EVENT_CENTER, xdEventType, "")
		}
	}
}
func GetRechargeAward(session *net.Session, page int32, requireid int) {

	state := module.State(session)
	out := &event_api.GetEventRechargeAward_Out{}
	if page > 0 {
		_, ok := event_dat.GetJsonEventInfoById(event_dat.JSON_EVENT_FIRST_RECHAGE_DAILY, page)
		if ok {
			jsonEventRecord, exists := state.JsonEventsState.GetStatus(event_dat.JSON_EVENT_FIRST_RECHAGE_DAILY, page)
			if exists {
				if time.IsToday(jsonEventRecord.LastUpdated) {
					out.IsRechage = true
					if jsonEventRecord.Awarded != int32(requireid) {
						out.IsAward = true
					}
				}
			}
		}
	} else {
		eventRecord := state.EventsState.GetPlayerEventInfoById(event_dat.EVENT_FIRST_RECHARGE_DAILY)
		// 是否充值了
		if time.IsToday(eventRecord.LastUpdated) {
			out.IsRechage = true
			//是否领取了
			nextAward := int32(event_dat.EventIntervalDays(event_dat.EVENT_FIRST_RECHARGE_DAILY))
			if eventRecord.Awarded != nextAward {
				out.IsAward = true
			}
		}
	}

	session.Send(out)
}
