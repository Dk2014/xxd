package rpc

import (
	"C"
	"core/mysql"
	"encoding/json"
	"errors"
	"game_server/config"
	"game_server/dat/event_dat"
	"game_server/dat/mail_dat"
	"game_server/mdb"
	. "game_server/rpc"
	"unsafe"
)

type Args_XdgmGetTextEvents struct {
	RPCArgTag
}

type EventAnnoce struct {
	Ltitle        string
	Rtitle        string
	Content       string
	StartUnixTime int64
	EndUnixTime   int64
	Weight        int16
	Tag           int8
	Jump          int8
	IsRelative    bool
}
type Reply_XdgmGetTextEvents struct {
	Events []*EventAnnoce `json:"events"`
}

func (this *RemoteServe) XdgmGetTextEvents(args *Args_XdgmGetTextEvents, reply *Reply_XdgmGetTextEvents) error {
	return Remote.Serve(mdb.RPC_Remote_XdgmGetTextEvents, args, mdb.TRANS_TAG_RPC_Serve_XdgmGetTextEvents, func() error {
		for _, textEvent := range event_dat.TextEvents {
			var isRelative bool
			if textEvent.IsRelative > 0 {
				isRelative = true
			}
			reply.Events = append(reply.Events, &EventAnnoce{
				Ltitle:        textEvent.Ltitle,
				Rtitle:        textEvent.Rtitle,
				Content:       textEvent.Content,
				StartUnixTime: textEvent.StartUnixTime,
				EndUnixTime:   textEvent.EndUnixTime,
				Weight:        textEvent.Weight,
				Tag:           textEvent.Tag,
				Jump:          textEvent.Jump,
				IsRelative:    isRelative,
			})
		}
		return nil
	})
}

type Args_XdgmGetNormalEventInfo struct {
	RPCArgTag
	Limit  int16
	Offset int16
}

type EventInfo struct {
	EventSign       int16
	Page            int32
	LTitle          string
	RTitle          string
	Content         string
	StartUnixTime   int64
	EndUnixTime     int64
	DisposeUnixTime int64
	IsRelative      bool
	Weight          int16
	Tag             int8
}

type Reply_XdgmGetNormalEventInfo struct {
	Total  int          `json:"total"`
	Events []*EventInfo `json:"events"`
}

func (this *RemoteServe) XdgmGetNormalEventInfo(args *Args_XdgmGetNormalEventInfo, reply *Reply_XdgmGetNormalEventInfo) error {
	var total int16
	expireEvents := map[int16]bool{
		event_dat.EVENT_FIRST_RECHARGE: true, event_dat.EVENT_SHARE_AWARDS: true, event_dat.EVENT_RECHARGE_AWARD: true, event_dat.EVENT_QQVIP_GIFT_AWARDS: true, event_dat.EVENT_GROUP_BUY: true,
		event_dat.EVENT_QQVIP_ADDITION: true, event_dat.EVENT_RICHMAN_CLUB: true, event_dat.EVENT_BUY_PARTNER: true, event_dat.EVENT_LEVEL_RECHARGE: true, event_dat.EVENT_MULTIPY_CONFIG: true,
		event_dat.EVENT_LOGIN_AWARD: true,
	}
	return Remote.Serve(mdb.RPC_Remote_XdgmGetNormalEventInfo, args, mdb.TRANS_TAG_RPC_Serve_XdgmGetNormalEventInfo, func() error {

		tempEvents := event_dat.GetEventsInfo()
		reply.Total = len(tempEvents)
		for _, event := range tempEvents {
			// 过滤过期活动
			if _, ok := expireEvents[event.Id]; ok {
				reply.Total--
			}
		}
		for _, event := range tempEvents {

			if _, ok := expireEvents[event.Id]; ok {
				continue
			}

			total++
			if total < args.Offset || total > args.Offset+args.Limit {
				continue
			}

			var isRelative bool
			if event.IsRelative > 0 {
				isRelative = true
			}
			reply.Events = append(reply.Events, &EventInfo{
				EventSign:       event.Id,
				LTitle:          event.LTitle,
				RTitle:          event.RTitle,
				Content:         event.Content,
				StartUnixTime:   event.Start,
				EndUnixTime:     event.End,
				DisposeUnixTime: event.Dispose,
				IsRelative:      isRelative,
				Weight:          event.Weight,
				Tag:             event.Tag,
			})
		}
		return nil
	})
	return nil
}

type Args_XdgmUpdateEventAwards struct {
	RPCArgTag
	ServerId       int
	EventId        int16
	EventAwardsRaw string
}
type Reply_XdgmUpdateEventAwards struct {
}

func (this *RemoteServe) XdgmUpdateEventAwards(args *Args_XdgmUpdateEventAwards, reply *Reply_XdgmUpdateEventAwards) error {
	eventAwardsList := make([]*event_dat.JsonEventAward, 0)
	err := json.Unmarshal([]byte(args.EventAwardsRaw), &eventAwardsList)
	if err != nil {
		return err
	}
	var eventInfoRaw string
	var eventInfo event_dat.EventConfigExtend
	db, err := mysql.Connect(config.GetDBConfig())
	if err != nil {
		return err
	}
	defer db.Close()
	res, err := db.ExecuteFetch([]byte(`select * from server_info`), -1)
	if err != nil {
		return err
	}

	iInfo := res.Map("event_json_info")

	for _, row := range res.Rows {
		eventInfoRaw = row.Str(iInfo)
	}
	err = json.Unmarshal([]byte(eventInfoRaw), &eventInfo)
	if err != nil {
		return err
	}

	// 持久化到数据库
	switch args.EventId {
	// TODO 标示特殊活动不能处理的 为：团购活动
	case event_dat.EVENT_DINNER_AWARDS:
		// 午餐活动
		dinnerList := make([]*event_dat.EventsDinner, 0)
		for _, award := range eventAwardsList {
			dinnerList = append(dinnerList, &event_dat.EventsDinner{
				RequireDinner: award.Grade,
				Ingot:         award.Award.Ingot,
				Coin:          award.Award.Coin,
				Item1Id:       award.Award.Item1Id,
				Item1Num:      award.Award.Item1Num,
				Item2Id:       award.Award.Item2Id,
				Item2Num:      award.Award.Item2Num,
				Item3Id:       award.Award.Item3Id,
				Item3Num:      award.Award.Item3Num,
				Item4Id:       award.Award.Item4Id,
				Item4Num:      award.Award.Item4Num,
				Item5Id:       award.Award.Item5Id,
				Item5Num:      award.Award.Item5Num,
			})
		}
		eventInfo.Dinner.List = dinnerList
		event_dat.LoadEventsDinner(dinnerList)
	case event_dat.EVENT_SUPPER_AWARDS:
		// 晚餐活动
		supperList := make([]*event_dat.EventsSupper, 0)
		for _, award := range eventAwardsList {
			supperList = append(supperList, &event_dat.EventsSupper{
				RequireSupper: award.Grade,
				Ingot:         award.Award.Ingot,
				Coin:          award.Award.Coin,
				Item1Id:       award.Award.Item1Id,
				Item1Num:      award.Award.Item1Num,
				Item2Id:       award.Award.Item2Id,
				Item2Num:      award.Award.Item2Num,
				Item3Id:       award.Award.Item3Id,
				Item3Num:      award.Award.Item3Num,
				Item4Id:       award.Award.Item4Id,
				Item4Num:      award.Award.Item4Num,
				Item5Id:       award.Award.Item5Id,
				Item5Num:      award.Award.Item5Num,
			})
		}
		eventInfo.Supper.List = supperList
		event_dat.LoadEventsSupper(supperList)
	case event_dat.EVENT_PHYSICAL_AWARDS:
		// 活跃度活动
		physicalList := make([]*event_dat.EventsPhysical, 0)
		for _, award := range eventAwardsList {
			physicalList = append(physicalList, &event_dat.EventsPhysical{
				RequirePhysical: award.Grade,
				Ingot:           award.Award.Ingot,
				Coin:            award.Award.Coin,
				Item1Id:         award.Award.Item1Id,
				Item1Num:        award.Award.Item1Num,
				Item2Id:         award.Award.Item2Id,
				Item2Num:        award.Award.Item2Num,
				Item3Id:         award.Award.Item3Id,
				Item3Num:        award.Award.Item3Num,
				Item4Id:         award.Award.Item4Id,
				Item4Num:        award.Award.Item4Num,
				Item5Id:         award.Award.Item5Id,
				Item5Num:        award.Award.Item5Num,
			})
		}
		eventInfo.Physical.List = physicalList
		event_dat.LoadEventsPhysical(physicalList)
	case event_dat.EVENT_MONTH_CARD_AWARDS:
		// 月卡活动
		monthcardList := make([]*event_dat.EventsMonthCard, 0)
		for _, award := range eventAwardsList {
			monthcardList = append(monthcardList, &event_dat.EventsMonthCard{
				RequireMonthCard: award.Grade,
				Ingot:            award.Award.Ingot,
				Coin:             award.Award.Coin,
				Item1Id:          award.Award.Item1Id,
				Item1Num:         award.Award.Item1Num,
				Item2Id:          award.Award.Item2Id,
				Item2Num:         award.Award.Item2Num,
				Item3Id:          award.Award.Item3Id,
				Item3Num:         award.Award.Item3Num,
				Item4Id:          award.Award.Item4Id,
				Item4Num:         award.Award.Item4Num,
				Item5Id:          award.Award.Item5Id,
				Item5Num:         award.Award.Item5Num,
			})
		}
		eventInfo.MonthCard.List = monthcardList
		event_dat.LoadEventsMonthCard(monthcardList)

	case event_dat.EVENT_SEVEN_DAY_AWARDS:
		// 新手七天登陆
		sevendayList := make([]*event_dat.EventsSevenDay, 0)
		for _, award := range eventAwardsList {
			sevendayList = append(sevendayList, &event_dat.EventsSevenDay{
				RequireDay: award.Grade,
				Ingot:      award.Award.Ingot,
				Coin:       award.Award.Coin,
				Item1Id:    award.Award.Item1Id,
				Item1Num:   award.Award.Item1Num,
				Item2Id:    award.Award.Item2Id,
				Item2Num:   award.Award.Item2Num,
				Item3Id:    award.Award.Item3Id,
				Item3Num:   award.Award.Item3Num,
				Item4Id:    award.Award.Item4Id,
				Item4Num:   award.Award.Item4Num,
				Item5Id:    award.Award.Item5Id,
				Item5Num:   award.Award.Item5Num,
			})
		}
		eventInfo.SevenDay.List = sevendayList
		event_dat.LoadEventsSevenDay(sevendayList)

	case event_dat.EVENT_LEVEL_AWARD:
		levelList := make([]*event_dat.EventsLevelUp, 0)
		for _, award := range eventAwardsList {
			levelList = append(levelList, &event_dat.EventsLevelUp{
				RequireLevel: int16(award.Grade),
				Ingot:        award.Award.Ingot,
				Coin:         int64(award.Award.Coin),
				Item1Id:      award.Award.Item1Id,
				Item1Num:     award.Award.Item1Num,
				Item2Id:      award.Award.Item2Id,
				Item2Num:     award.Award.Item2Num,
				Item3Id:      award.Award.Item3Id,
				Item3Num:     award.Award.Item3Num,
			})
		}
		eventInfo.LevelUp.List = levelList
		event_dat.LoadEventsLevelUp(levelList)
	case event_dat.EVENT_STRONG_AWARD:
		// 战力活动 TODO 不统一 要单独处理
		// fightpowerList := make([]*event_dat.EventFightPowerAward, 0)
		// for _, award := range eventAwardsList {
		// 	fightpowerList = append(fightpowerList, &event_dat.EventFightPowerAward{
		// 		FightPower: award.Grade,
		// 		Ingot:      award.Award.Ingot,
		// 		Coin:       award.Award.Coin,
		// 		Item1Id:    award.Award.Item1Id,
		// 		Item1Num:   award.Award.Item1Num,
		// 		Item2Id:    award.Award.Item2Id,
		// 		Item2Num:   award.Award.Item2Num,
		// 		Item3Id:    award.Award.Item3Id,
		// 		Item3Num:   award.Award.Item3Num,
		// 		Item4Id:    award.Award.Item4Id,
		// 		Item4Num:   award.Award.Item4Num,
		// 		Item5Id:    award.Award.Item5Id,
		// 		Item5Num:   award.Award.Item5Num,
		// 	})
		// }
		// eventInfo.FightPower.List = fightpowerList
		// event_dat.LoadEventFightPower(fightpowerList)
	case event_dat.EVENT_VIP_CLUB:
		// VIP俱乐部
		vipclubList := make([]*event_dat.EventsVipClub, 0)
		for _, award := range eventAwardsList {
			vipclubList = append(vipclubList, &event_dat.EventsVipClub{
				RequireVipCount: award.Grade,
				Ingot:           award.Award.Ingot,
				Coin:            award.Award.Coin,
				Item1Id:         award.Award.Item1Id,
				Item1Num:        award.Award.Item1Num,
				Item2Id:         award.Award.Item2Id,
				Item2Num:        award.Award.Item2Num,
				Item3Id:         award.Award.Item3Id,
				Item3Num:        award.Award.Item3Num,
				Item4Id:         award.Award.Item4Id,
				Item4Num:        award.Award.Item4Num,
				Item5Id:         award.Award.Item5Id,
				Item5Num:        award.Award.Item5Num,
			})
		}
		eventInfo.VipClub.List = vipclubList
		event_dat.LoadEventsVipClub(vipclubList)
	case event_dat.EVENT_ARENA_RANK_AWARDS:
		// 比武场活动
		arearankList := make([]*event_dat.EventsArenaRank, 0)
		for _, award := range eventAwardsList {
			arearankList = append(arearankList, &event_dat.EventsArenaRank{
				RequireArenaRank: award.Grade,
				Ingot:            award.Award.Ingot,
				Coin:             award.Award.Coin,
				Item1Id:          award.Award.Item1Id,
				Item1Num:         award.Award.Item1Num,
				Item2Id:          award.Award.Item2Id,
				Item2Num:         award.Award.Item2Num,
				Item3Id:          award.Award.Item3Id,
				Item3Num:         award.Award.Item3Num,
				Item4Id:          award.Award.Item4Id,
				Item4Num:         award.Award.Item4Num,
				Item5Id:          award.Award.Item5Id,
				Item5Num:         award.Award.Item5Num,
			})
		}
		eventInfo.ArenaRank.List = arearankList
		event_dat.LoadEventsArenaRank(arearankList)
	case event_dat.EVENT_SHARE_AWARDS:
		// 分享送好礼活动
		shareList := make([]*event_dat.EventsShare, 0)
		for _, award := range eventAwardsList {
			shareList = append(shareList, &event_dat.EventsShare{
				RequireTimes: award.Grade,
				Ingot:        award.Award.Ingot,
				Coin:         award.Award.Coin,
				Item1Id:      award.Award.Item1Id,
				Item1Num:     award.Award.Item1Num,
				Item2Id:      award.Award.Item2Id,
				Item2Num:     award.Award.Item2Num,
				Item3Id:      award.Award.Item3Id,
				Item3Num:     award.Award.Item3Num,
				Item4Id:      award.Award.Item4Id,
				Item4Num:     award.Award.Item4Num,
				Item5Id:      award.Award.Item5Id,
				Item5Num:     award.Award.Item5Num,
			})
		}
		eventInfo.Share.List = shareList
		event_dat.LoadEventsShare(shareList)
	case event_dat.EVENT_TOTAL_RECHARGE:
		// 累计充值活动
		totalRechargeList := make([]*event_dat.EventsTotalRechargeAwardsExt, 0)
		for _, award := range eventAwardsList {
			totalRechargeList = append(totalRechargeList, &event_dat.EventsTotalRechargeAwardsExt{
				RequireTotal: int16(award.Grade),
				Ingot:        award.Award.Ingot,
				Coin:         award.Award.Coin,
				Heart:        award.Award.Heart,
				Item1Id:      award.Award.Item1Id,
				Item1Num:     award.Award.Item1Num,
				Item2Id:      award.Award.Item2Id,
				Item2Num:     award.Award.Item2Num,
				Item3Id:      award.Award.Item3Id,
				Item3Num:     award.Award.Item3Num,
				Item4Id:      award.Award.Item4Id,
				Item4Num:     award.Award.Item4Num,
				Item5Id:      award.Award.Item5Id,
				Item5Num:     award.Award.Item5Num,
			})
		}
		eventInfo.TotalRecharge.List = totalRechargeList
		event_dat.LoadEventsTotalRecharge(totalRechargeList)
	case event_dat.EVENT_TEN_DRAW:
		// 宝箱十连抽活动
		tendrawList := make([]*event_dat.EventsTenDraw, 0)
		for _, award := range eventAwardsList {
			tendrawList = append(tendrawList, &event_dat.EventsTenDraw{
				RequireTimes: award.Grade,
				Ingot:        award.Award.Ingot,
				Coin:         award.Award.Coin,
				Item1Id:      award.Award.Item1Id,
				Item1Num:     award.Award.Item1Num,
				Item2Id:      award.Award.Item2Id,
				Item2Num:     award.Award.Item2Num,
				Item3Id:      award.Award.Item3Id,
				Item3Num:     award.Award.Item3Num,
				Item4Id:      award.Award.Item4Id,
				Item4Num:     award.Award.Item4Num,
				Item5Id:      award.Award.Item5Id,
				Item5Num:     award.Award.Item5Num,
			})
		}
		eventInfo.TenDraw.List = tendrawList
		event_dat.LoadEventsTenDraw(tendrawList)

	case event_dat.EVENT_TOTAL_CONSUME:
		// 累计消费活动
		totalConsumeList := make([]*event_dat.EventsTotalConsumeAward, 0)
		for _, award := range eventAwardsList {
			totalConsumeList = append(totalConsumeList, &event_dat.EventsTotalConsumeAward{
				RequireCost: int16(award.Grade),
				Ingot:       award.Award.Ingot,
				Coin:        award.Award.Coin,
				Heart:       award.Award.Heart,
				Item1Id:     award.Award.Item1Id,
				Item1Num:    award.Award.Item1Num,
				Item2Id:     award.Award.Item2Id,
				Item2Num:    award.Award.Item2Num,
				Item3Id:     award.Award.Item3Id,
				Item3Num:    award.Award.Item3Num,
				Item4Id:     award.Award.Item4Id,
				Item4Num:    award.Award.Item4Num,
				Item5Id:     award.Award.Item5Id,
				Item5Num:    award.Award.Item5Num,
			})
		}
		eventInfo.TotalConsume.List = totalConsumeList
		event_dat.LoadEventTotalConsume(totalConsumeList)
	case event_dat.EVENT_FIRST_RECHARGE_DAILY:
		// 每日首冲活动
		firstRechargeDailyList := make([]*event_dat.EventsFRDAward, 0)
		for _, award := range eventAwardsList {
			firstRechargeDailyList = append(firstRechargeDailyList, &event_dat.EventsFRDAward{
				RequireDay: int16(award.Grade),
				Ingot:      award.Award.Ingot,
				Coin:       award.Award.Coin,
				Heart:      award.Award.Heart,
				Item1Id:    award.Award.Item1Id,
				Item1Num:   award.Award.Item1Num,
				Item2Id:    award.Award.Item2Id,
				Item2Num:   award.Award.Item2Num,
				Item3Id:    award.Award.Item3Id,
				Item3Num:   award.Award.Item3Num,
				Item4Id:    award.Award.Item4Id,
				Item4Num:   award.Award.Item4Num,
				Item5Id:    award.Award.Item5Id,
				Item5Num:   award.Award.Item5Num,
			})
		}
		eventInfo.FirstRechargeDaily.List = firstRechargeDailyList
		event_dat.LoadEventsFirstRechargeDaily(firstRechargeDailyList)
	case event_dat.EVENT_TOTAL_LOGIN:
		// 连续登陆活动
		totalLoginList := make([]*event_dat.EventsTotalLogin, 0)
		for _, award := range eventAwardsList {
			totalLoginList = append(totalLoginList, &event_dat.EventsTotalLogin{
				RequireLoginDays: award.Grade,
				Ingot:            award.Award.Ingot,
				Coin:             award.Award.Coin,
				Heart:            award.Award.Heart,
				Item1Id:          award.Award.Item1Id,
				Item1Num:         award.Award.Item1Num,
				Item2Id:          award.Award.Item2Id,
				Item2Num:         award.Award.Item2Num,
				Item3Id:          award.Award.Item3Id,
				Item3Num:         award.Award.Item3Num,
				Item4Id:          award.Award.Item4Id,
				Item4Num:         award.Award.Item4Num,
				Item5Id:          award.Award.Item5Id,
				Item5Num:         award.Award.Item5Num,
			})
		}
		eventInfo.TotalLogin.List = totalLoginList
		event_dat.LoadEventsTotalLogin(totalLoginList)
	case event_dat.EVENT_RICHMAN_CLUB:
		// 土豪俱乐部活动
	}
	eventInfo.Version++
	eventInfoBytes, err := json.Marshal(eventInfo)
	stmt, err := db.Prepare([]byte("UPDATE `server_info` SET `event_version`=?, `event_json_info`=?;"))
	if err != nil {
		return err
	}
	stmt.BindInt(unsafe.Pointer(&eventInfo.Version))
	str := C.CString(string(eventInfoBytes))
	stmt.BindVarchar(unsafe.Pointer(str), len(string(eventInfoBytes)))
	stmt.Execute()
	stmt.Close()
	return nil
}

type Args_XdgmUpdateNormalEventInfo struct {
	RPCArgTag
	ServerId      int
	EventsInfoRaw string
}
type Reply_XdgmUpdateNormalEventInfo struct {
}

func (this *RemoteServe) XdgmUpdateNormalEventInfo(args *Args_XdgmUpdateNormalEventInfo, reply *Reply_XdgmUpdateNormalEventInfo) error {
	eventsList := make([]*EventInfo, 0)
	err := json.Unmarshal([]byte(args.EventsInfoRaw), &eventsList)
	if err != nil {
		return err
	}
	// server := GetServerBySId(args.ServerId, args.Platid, args.ChannelId)
	// for _, s := range server.GServers {
	var eventInfoRaw string
	var eventInfo event_dat.EventConfigExtend
	db, err := mysql.Connect(config.GetDBConfig())
	if err != nil {
		return err
	}
	defer db.Close()
	res, err := db.ExecuteFetch([]byte(`select * from server_info`), -1)
	if err != nil {
		return err
	}

	iInfo := res.Map("event_json_info")

	for _, row := range res.Rows {
		eventInfoRaw = row.Str(iInfo)
	}
	err = json.Unmarshal([]byte(eventInfoRaw), &eventInfo)
	if err != nil {
		return err
	}
	for _, event := range eventsList {
		var isRelative int8
		if event.IsRelative {
			isRelative = 1
		}
		event_dat.LoadEventCenterExt(event.StartUnixTime, event.EndUnixTime, event.DisposeUnixTime, event.EventSign, isRelative, event.LTitle, event.RTitle, event.Content, event.Weight, event.Tag)
		// 持久化到数据库
		switch event.EventSign {
		// TODO 标示特殊活动不能处理的 为：团购活动
		case event_dat.EVENT_DINNER_AWARDS:
			// 午餐活动
			eventInfo.Dinner = &event_dat.EventsDinnerExt{
				StartUnixTime:   event.StartUnixTime,
				EndUnixTime:     event.EndUnixTime,
				DisposeUnixTime: event.DisposeUnixTime,
				IsRelative:      isRelative,
				LTitle:          event.LTitle,
				RTitle:          event.RTitle,
				Content:         event.Content,
				Tag:             event.Tag,
				Weight:          event.Weight,
				List:            eventInfo.Dinner.List,
			}
		case event_dat.EVENT_SUPPER_AWARDS:
			// 晚餐活动
			eventInfo.Supper = &event_dat.EventsSupperExt{
				StartUnixTime:   event.StartUnixTime,
				EndUnixTime:     event.EndUnixTime,
				DisposeUnixTime: event.DisposeUnixTime,
				IsRelative:      isRelative,
				LTitle:          event.LTitle,
				RTitle:          event.RTitle,
				Content:         event.Content,
				Tag:             event.Tag,
				Weight:          event.Weight,
				List:            eventInfo.Supper.List,
			}
		case event_dat.EVENT_PHYSICAL_AWARDS:
			// 活跃度活动
			eventInfo.Physical = &event_dat.EventsPhysicalExt{
				StartUnixTime:   event.StartUnixTime,
				EndUnixTime:     event.EndUnixTime,
				DisposeUnixTime: event.DisposeUnixTime,
				IsRelative:      isRelative,
				LTitle:          event.LTitle,
				RTitle:          event.RTitle,
				Content:         event.Content,
				Tag:             event.Tag,
				Weight:          event.Weight,
				List:            eventInfo.Physical.List,
			}
		case event_dat.EVENT_MONTH_CARD_AWARDS:
			// 月卡活动
			eventInfo.MonthCard = &event_dat.EventsMonthCardExt{
				StartUnixTime:   event.StartUnixTime,
				EndUnixTime:     event.EndUnixTime,
				DisposeUnixTime: event.DisposeUnixTime,
				IsRelative:      isRelative,
				LTitle:          event.LTitle,
				RTitle:          event.RTitle,
				Content:         event.Content,
				Tag:             event.Tag,
				Weight:          event.Weight,
				List:            eventInfo.MonthCard.List,
			}

		case event_dat.EVENT_SEVEN_DAY_AWARDS:
			// 新手七天登陆
			eventInfo.SevenDay = &event_dat.EventsSevenDayExt{
				StartUnixTime:   event.StartUnixTime,
				EndUnixTime:     event.EndUnixTime,
				DisposeUnixTime: event.DisposeUnixTime,
				IsRelative:      isRelative,
				LTitle:          event.LTitle,
				RTitle:          event.RTitle,
				Content:         event.Content,
				Tag:             event.Tag,
				Weight:          event.Weight,
				List:            eventInfo.SevenDay.List,
			}
		case event_dat.EVENT_LEVEL_AWARD:
			eventInfo.LevelUp = &event_dat.EventsLevelUpExt{
				StartUnixTime:   event.StartUnixTime,
				EndUnixTime:     event.EndUnixTime,
				DisposeUnixTime: event.DisposeUnixTime,
				IsRelative:      isRelative,
				LTitle:          event.LTitle,
				RTitle:          event.RTitle,
				Content:         event.Content,
				Tag:             event.Tag,
				Weight:          event.Weight,
				List:            eventInfo.LevelUp.List,
			}
		case event_dat.EVENT_STRONG_AWARD:
			// 战力活动
			eventInfo.FightPower = &event_dat.EventFightPower{
				StartUnixTime:   event.StartUnixTime,
				EndUnixTime:     event.EndUnixTime,
				DisposeUnixTime: event.DisposeUnixTime,
				IsRelative:      isRelative,
				LTitle:          event.LTitle,
				RTitle:          event.RTitle,
				Content:         event.Content,
				Tag:             event.Tag,
				Weight:          event.Weight,
				List:            eventInfo.FightPower.List,
			}
		case event_dat.EVENT_VIP_CLUB:
			// VIP俱乐部
			eventInfo.VipClub = &event_dat.EventsVipClubExt{
				StartUnixTime:   event.StartUnixTime,
				EndUnixTime:     event.EndUnixTime,
				DisposeUnixTime: event.DisposeUnixTime,
				IsRelative:      isRelative,
				LTitle:          event.LTitle,
				RTitle:          event.RTitle,
				Content:         event.Content,
				Tag:             event.Tag,
				Weight:          event.Weight,
				List:            eventInfo.VipClub.List,
			}
		case event_dat.EVENT_ARENA_RANK_AWARDS:
			// 比武场活动
			eventInfo.ArenaRank = &event_dat.EventsArenaRankExt{
				StartUnixTime:   event.StartUnixTime,
				EndUnixTime:     event.EndUnixTime,
				DisposeUnixTime: event.DisposeUnixTime,
				IsRelative:      isRelative,
				LTitle:          event.LTitle,
				RTitle:          event.RTitle,
				Content:         event.Content,
				Tag:             event.Tag,
				Weight:          event.Weight,
				List:            eventInfo.ArenaRank.List,
			}
		case event_dat.EVENT_SHARE_AWARDS:
			// 分享送好礼活动
			eventInfo.Share = &event_dat.EventsShareExt{
				StartUnixTime:   event.StartUnixTime,
				EndUnixTime:     event.EndUnixTime,
				DisposeUnixTime: event.DisposeUnixTime,
				IsRelative:      isRelative,
				LTitle:          event.LTitle,
				RTitle:          event.RTitle,
				Content:         event.Content,
				Tag:             event.Tag,
				Weight:          event.Weight,
				List:            eventInfo.Share.List,
			}
		case event_dat.EVENT_TOTAL_RECHARGE:
			// 累计充值活动
			eventInfo.TotalRecharge = &event_dat.EventTotalRechargeExt{
				StartUnixTime:   event.StartUnixTime,
				EndUnixTime:     event.EndUnixTime,
				DisposeUnixTime: event.DisposeUnixTime,
				IsRelative:      isRelative,
				LTitle:          event.LTitle,
				RTitle:          event.RTitle,
				Content:         event.Content,
				Tag:             event.Tag,
				Weight:          event.Weight,
				List:            eventInfo.TotalRecharge.List,
			}
		case event_dat.EVENT_TEN_DRAW:
			// 宝箱十连抽活动
			eventInfo.TenDraw = &event_dat.EventsTenDrawExt{
				StartUnixTime:   event.StartUnixTime,
				EndUnixTime:     event.EndUnixTime,
				DisposeUnixTime: event.DisposeUnixTime,
				IsRelative:      isRelative,
				LTitle:          event.LTitle,
				RTitle:          event.RTitle,
				Content:         event.Content,
				Tag:             event.Tag,
				Weight:          event.Weight,
				List:            eventInfo.TenDraw.List,
			}
		case event_dat.EVENT_TOTAL_CONSUME:
			// 累计消费活动
			eventInfo.TotalConsume = &event_dat.EventsTotalConsumeExt{
				StartUnixTime:   event.StartUnixTime,
				EndUnixTime:     event.EndUnixTime,
				DisposeUnixTime: event.DisposeUnixTime,
				IsRelative:      isRelative,
				LTitle:          event.LTitle,
				RTitle:          event.RTitle,
				Content:         event.Content,
				Tag:             event.Tag,
				Weight:          event.Weight,
				List:            eventInfo.TotalConsume.List,
			}
		case event_dat.EVENT_FIRST_RECHARGE_DAILY:
			// 每日首冲活动
			eventInfo.FirstRechargeDaily = &event_dat.EventsFirstRechargeDailyExt{
				StartUnixTime:   event.StartUnixTime,
				EndUnixTime:     event.EndUnixTime,
				DisposeUnixTime: event.DisposeUnixTime,
				IsRelative:      isRelative,
				LTitle:          event.LTitle,
				RTitle:          event.RTitle,
				Content:         event.Content,
				Tag:             event.Tag,
				Weight:          event.Weight,
				List:            eventInfo.FirstRechargeDaily.List,
			}
		case event_dat.EVENT_TOTAL_LOGIN:
			// 连续登陆活动
			eventInfo.TotalLogin = &event_dat.EventsTotalLoginExt{
				StartUnixTime:   event.StartUnixTime,
				EndUnixTime:     event.EndUnixTime,
				DisposeUnixTime: event.DisposeUnixTime,
				IsRelative:      isRelative,
				LTitle:          event.LTitle,
				RTitle:          event.RTitle,
				Content:         event.Content,
				Tag:             event.Tag,
				Weight:          event.Weight,
				List:            eventInfo.TotalLogin.List,
			}
		case event_dat.EVENT_RICHMAN_CLUB:
			// 土豪俱乐部活动
			eventInfo.Richman = &event_dat.EventsRichmanExt{
				StartUnixTime:   event.StartUnixTime,
				EndUnixTime:     event.EndUnixTime,
				DisposeUnixTime: event.DisposeUnixTime,
				IsRelative:      isRelative,
				LTitle:          event.LTitle,
				RTitle:          event.RTitle,
				Content:         event.Content,
				Tag:             event.Tag,
				Weight:          event.Weight,
			}
		}
	}
	eventInfo.Version++
	eventInfoBytes, err := json.Marshal(eventInfo)
	stmt, err := db.Prepare([]byte("UPDATE `server_info` SET `event_version`=?, `event_json_info`=?;"))
	if err != nil {
		return err
	}
	stmt.BindInt(unsafe.Pointer(&eventInfo.Version))
	str := C.CString(string(eventInfoBytes))
	stmt.BindVarchar(unsafe.Pointer(str), len(string(eventInfoBytes)))
	stmt.Execute()
	stmt.Close()
	return nil
}

type Args_XdgmGetJsonEventInfo struct {
	RPCArgTag
	Limit  int16
	Offset int16
}

type Reply_XdgmGetJsonEventInfo struct {
	Total  int          `json:"total"`
	Events []*EventInfo `json:"events"`
}

func (this *RemoteServe) XdgmGetJsonEventInfo(args *Args_XdgmGetJsonEventInfo, reply *Reply_XdgmGetJsonEventInfo) error {
	var total int16
	return Remote.Serve(mdb.RPC_Remote_XdgmGetJsonEventInfo, args, mdb.TRANS_TAG_RPC_Serve_XdgmGetJsonEventInfo, func() error {
		// 接下来加载json配置活动 可重复开启的那一类
		jsonEvents := event_dat.GetJsonEventsInfo()
		reply.Total = len(jsonEvents)
		for _, event := range jsonEvents {
			total++
			if total < args.Offset || total > args.Offset+args.Limit {
				continue
			}
			reply.Events = append(reply.Events, &EventInfo{
				EventSign:       event.Type,
				Page:            event.Page,
				LTitle:          event.LTitle,
				RTitle:          event.RTitle,
				Content:         event.Content,
				StartUnixTime:   event.StartUnixTime,
				EndUnixTime:     event.EndUnixTime,
				DisposeUnixTime: event.DisposeUnixTime,
				IsRelative:      event.IsRelative,
				Weight:          event.Weight,
				Tag:             event.Tag,
			})
		}
		return nil
	})
	return nil
}

type Args_XdgmGetEventAwardInfo struct {
	RPCArgTag
	EventSign int16
	Page      int32
}

type Reply_XdgmGetEventAwardInfo struct {
	Awards []*event_dat.JsonEventAward
}

func (this *RemoteServe) XdgmGetEventAwardInfo(args *Args_XdgmGetEventAwardInfo, reply *Reply_XdgmGetEventAwardInfo) error {
	return Remote.Serve(mdb.RPC_Remote_XdgmGetEventAwardInfo, args, mdb.TRANS_TAG_RPC_Serve_XdgmGetEventAwardInfo, func() error {
		if args.Page == 0 {
			// 正常活动
			events := event_dat.GetEventsInfo()
			for _, event := range events {
				if event.Id == args.EventSign {
					switch event.Id {
					// TODO 标示特殊活动不能处理的 为：团购活动
					case event_dat.EVENT_DINNER_AWARDS:
						// 午餐活动
						awards := event_dat.GetEventDinnerAwards()
						for day, award := range awards {
							reply.Awards = append(reply.Awards, &event_dat.JsonEventAward{
								Grade: day,
								Award: award,
							})
						}
					case event_dat.EVENT_SUPPER_AWARDS:
						// 晚餐活动
						awards := event_dat.GetEventSupperAwards()
						for day, award := range awards {
							reply.Awards = append(reply.Awards, &event_dat.JsonEventAward{
								Grade: day,
								Award: award,
							})
						}
					case event_dat.EVENT_PHYSICAL_AWARDS:
						// 活跃度活动
						awards := event_dat.GetEventPhysicalAwards()
						for _, award := range awards {
							reply.Awards = append(reply.Awards, &event_dat.JsonEventAward{
								Grade: award.Require_physical,
								Award: award.Awards,
							})

						}
					case event_dat.EVENT_MONTH_CARD_AWARDS:
						// 月卡活动
						awards := event_dat.GetEventMonthCardAwards()
						for _, award := range awards {
							reply.Awards = append(reply.Awards, &event_dat.JsonEventAward{
								Grade: award.Require_month_card,
								Award: award.Awards,
							})
						}
					case event_dat.EVENT_SEVEN_DAY_AWARDS:
						// 新手七天登陆
						awards := event_dat.GetEventSevenDayAwards()
						for _, award := range awards {
							reply.Awards = append(reply.Awards, &event_dat.JsonEventAward{
								Grade: award.Require_day,
								Award: award.Awards,
							})
						}
					case event_dat.EVENT_LEVEL_AWARD:
						// 等级活动 返回奖励内容
						awards := event_dat.GetEventLevelUpAwards()
						for _, award := range awards {
							index := 1
							var realAward *event_dat.EventDefaultAward
							for _, mailAward := range award.Awards {
								if mailAward.AttachmentType == mail_dat.ATTACHMENT_INGOT {
									realAward.Ingot = int16(mailAward.ItemNum)
								} else if mailAward.AttachmentType == mail_dat.ATTACHMENT_COINS {
									realAward.Coin = int32(mailAward.ItemNum)
								} else {
									switch index {
									case 1:
										realAward.Item1Id = mailAward.ItemId
										realAward.Item1Num = int16(mailAward.ItemNum)
									case 2:
										realAward.Item2Id = mailAward.ItemId
										realAward.Item2Num = int16(mailAward.ItemNum)
									case 3:
										realAward.Item3Id = mailAward.ItemId
										realAward.Item3Num = int16(mailAward.ItemNum)
									case 4:
										realAward.Item4Id = mailAward.ItemId
										realAward.Item4Num = int16(mailAward.ItemNum)
									case 5:
										realAward.Item5Id = mailAward.ItemId
										realAward.Item5Num = int16(mailAward.ItemNum)
									}
									index++
								}
							}
							reply.Awards = append(reply.Awards, &event_dat.JsonEventAward{
								Grade: int32(award.Require_level),
								Award: realAward,
							})
						}
					case event_dat.EVENT_STRONG_AWARD:
						// 战力活动
						awards := event_dat.GetEventFightPowerAwards()
						for _, award := range awards {
							reply.Awards = append(reply.Awards, &event_dat.JsonEventAward{
								Grade: award.FightPower,
								Award: &award.EventDefaultAward,
							})
						}
					case event_dat.EVENT_VIP_CLUB:
						// VIP俱乐部
						awards := event_dat.GetEventVipClubAwards()
						for _, award := range awards {
							reply.Awards = append(reply.Awards, &event_dat.JsonEventAward{
								Grade: award.Require_vip_count,
								Award: award.Awards,
							})
						}
					case event_dat.EVENT_ARENA_RANK_AWARDS:
						// 比武场活动
						awards := event_dat.GetEventArenaRankAwards()
						for grade, award := range awards {
							reply.Awards = append(reply.Awards, &event_dat.JsonEventAward{
								Grade: grade,
								Award: award,
							})
						}
					case event_dat.EVENT_SHARE_AWARDS:
						// 分享送好礼活动
						awards := event_dat.GetEventShareAwards()
						for _, award := range awards {
							reply.Awards = append(reply.Awards, &event_dat.JsonEventAward{
								Grade: award.Require_times,
								Award: award.Awards,
							})
						}
					case event_dat.EVENT_TOTAL_RECHARGE:
						// 累计充值活动
						awards := event_dat.GetEventTotalRechargeAwards()
						for _, award := range awards {
							reply.Awards = append(reply.Awards, &event_dat.JsonEventAward{
								Grade: int32(award.RequireTotal),
								Award: award.Awards,
							})
						}
					case event_dat.EVENT_TEN_DRAW:
						// 宝箱十连抽活动
						awards := event_dat.GetEventTenDrawAwards()
						for _, awardInfo := range awards {
							reply.Awards = append(reply.Awards, &event_dat.JsonEventAward{
								Grade: int32(awardInfo.RequireTimes),
								Award: awardInfo.Award,
							})
						}
					case event_dat.EVENT_TOTAL_CONSUME:
						// 累计消费活动
						awards := event_dat.GetTotalConsumeAwards()
						reply.Awards = append(reply.Awards, &event_dat.JsonEventAward{
							Grade: int32(awards.RequireCost),
							Award: awards.Award,
						})
					case event_dat.EVENT_FIRST_RECHARGE_DAILY:
						// 每日首冲活动
						awards := event_dat.GetFirstRechargeDailyRewards()
						for _, award := range awards {
							reply.Awards = append(reply.Awards, &event_dat.JsonEventAward{
								Grade: int32(award.RequireDay),
								Award: award.EventAward,
							})
						}
					case event_dat.EVENT_TOTAL_LOGIN:
						// 连续登陆活动
						awards := event_dat.GetEventTotalLoginAwards()
						for _, award := range awards {
							reply.Awards = append(reply.Awards, &event_dat.JsonEventAward{
								Grade: award.Require_login_days,
								Award: award.Awards,
							})
						}
					case event_dat.EVENT_RICHMAN_CLUB:
						// 土豪俱乐部活动 TOD 不统一
					}
					return nil
				}
			}
		} else if args.Page > 0 {
			// json配置活动
			award, exists := event_dat.GetJsonEventInfoById(args.EventSign, args.Page)
			if !exists {
				return errors.New("json event not found")
			}
			reply.Awards = award.List
			return nil
		}
		return errors.New("event not found")
	})
	return nil
}
