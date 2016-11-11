package draw

import (
	"core/fail"
	"core/net"
	"core/time"
	"game_server/api/protocol/draw_api"
	"game_server/api/protocol/notify_api"
	"game_server/dat/channel_dat"
	"game_server/dat/chest_dat"
	"game_server/dat/event_dat"
	"game_server/dat/item_dat"
	"game_server/dat/mail_dat"
	"game_server/dat/player_dat"
	"game_server/dat/which_branch_dat"
	"game_server/module"
	"game_server/module/rpc"
	"game_server/tlog"
	"game_server/xdlog"
)

func _getTlogInfoByFateBoxType(boxType int32, is_ten, free bool) (moneyreason, itemreason, xdEventType, tlogChestType, moneyNum int32) {
	switch boxType {
	case chest_dat.StarBox:
		if free {
			return tlog.MFR_STAR_FATE_BOX_FREE, tlog.IFR_STAR_FATE_BOX_FREE, xdlog.ET_STAR_FATE_BOX_FREE, tlog.PCD_STAR_FATE_BOX_FREE, 0
		}
		if !is_ten {
			return tlog.MFR_STAR_FATE_BOX_ONE, tlog.IFR_STAR_FATE_BOX_ONE, xdlog.ET_STAR_FATE_BOX_ONE, tlog.PCD_STAR_FATE_BOX_ONE, chest_dat.FATE_BOX_PRICE
		}
		return tlog.MFR_STAR_FATE_BOX_TEN, tlog.IFR_STAR_FATE_BOX_TEN, xdlog.ET_STAR_FATE_BOX_TEN, tlog.PCD_STAR_FATE_BOX_TEN, chest_dat.FATE_BOX_TEN_TIME_PRICE
	case chest_dat.MoonBox:
		if free {
			return tlog.MFR_MOON_FATE_BOX_FREE, tlog.IFR_MOON_FATE_BOX_FREE, xdlog.ET_MOON_FATE_BOX_FREE, tlog.PCD_MOON_FATE_BOX_FREE, 0
		}
		if !is_ten {
			return tlog.MFR_MOON_FATE_BOX_ONE, tlog.IFR_MOON_FATE_BOX_ONE, xdlog.ET_MOON_FATE_BOX_ONE, tlog.PCD_MOON_FATE_BOX_ONE, chest_dat.FATE_BOX_PRICE
		}
		return tlog.MFR_MOON_FATE_BOX_TEN, tlog.IFR_MOON_FATE_BOX_TEN, xdlog.ET_MOON_FATE_BOX_TEN, tlog.PCD_MOON_FATE_BOX_TEN, chest_dat.FATE_BOX_TEN_TIME_PRICE
	case chest_dat.SunBox:
		if free {
			return tlog.MFR_SUN_FATE_BOX_FREE, tlog.IFR_SUN_FATE_BOX_FREE, xdlog.ET_SUN_FATE_BOX_FREE, tlog.PCD_SUN_FATE_BOX_FREE, 0
		}
		if !is_ten {
			return tlog.MFR_SUN_FATE_BOX_ONE, tlog.IFR_SUN_FATE_BOX_ONE, xdlog.ET_SUN_FATE_BOX_ONE, tlog.PCD_SUN_FATE_BOX_ONE, chest_dat.FATE_BOX_PRICE
		}
		return tlog.MFR_SUN_FATE_BOX_TEN, tlog.IFR_SUN_FATE_BOX_TEN, xdlog.ET_SUN_FATE_BOX_TEN, tlog.PCD_SUN_FATE_BOX_TEN, chest_dat.FATE_BOX_TEN_TIME_PRICE
	case chest_dat.HunyuanBox:
		if free {
			return tlog.MFR_HUNYUAN_FATE_BOX_FREE, tlog.IFR_HUNYUAN_FATE_BOX_FREE, xdlog.ET_HUNYUAN_FATE_BOX_FREE, tlog.PCD_HUNYUAN_FATE_BOX_FREE, 0
		}
		if !is_ten {
			return tlog.MFR_HUNYUAN_FATE_BOX_ONE, tlog.IFR_HUNYUAN_FATE_BOX_ONE, xdlog.ET_HUNYUAN_FATE_BOX_ONE, tlog.PCD_HUNYUAN_FATE_BOX_ONE, chest_dat.FATE_BOX_PRICE
		}
		return tlog.MFR_HUNYUAN_FATE_BOX_TEN, tlog.IFR_HUNYUAN_FATE_BOX_TEN, xdlog.ET_HUNYUAN_FATE_BOX_TEN, tlog.PCD_HUNYUAN_FATE_BOX_TEN, chest_dat.FATE_BOX_TEN_TIME_PRICE
	}

	panic("undefine fate box type")
}

func getFateBoxInfo(state *module.SessionState, out *draw_api.GetFateBoxInfo_Out) {
	fail.When(!module.Player.IsOpenFunc(state.Database, player_dat.FUNC_FATE_BOX), "命锁宝箱功能未开启")

	playerFateBoxState := state.Database.Lookup.PlayerFateBoxState(state.PlayerId)
	out.Lock = playerFateBoxState.Lock
	nowTime := time.GetNowTime()
	//已冷却时间
	freeBoxCD := nowTime - playerFateBoxState.StarBoxFreeDrawTimestamp
	if freeBoxCD >= chest_dat.FREE_FATE_BOX_CD {
		out.NextFreeStarBox = 0
		if playerFateBoxState.StarBoxFreeDrawTimestamp == 0 {
			out.NextFreeStarBox = chest_dat.FATE_BOX_INIT_CD_TIME
		}
	} else {
		out.NextFreeStarBox = chest_dat.FREE_FATE_BOX_CD - int32(freeBoxCD)
	}

	freeBoxCD = nowTime - playerFateBoxState.MoonBoxFreeDrawTimestamp
	if freeBoxCD >= chest_dat.FREE_FATE_BOX_CD {
		out.NextFreeMoonBox = 0
		if playerFateBoxState.MoonBoxFreeDrawTimestamp == 0 {
			out.NextFreeMoonBox = chest_dat.FATE_BOX_INIT_CD_TIME
		}
	} else {
		out.NextFreeMoonBox = chest_dat.FREE_FATE_BOX_CD - int32(freeBoxCD)
	}

	freeBoxCD = nowTime - playerFateBoxState.SunBoxFreeDrawTimestamp
	if freeBoxCD >= chest_dat.FREE_FATE_BOX_CD {
		out.NextFreeSunBox = 0
		if playerFateBoxState.SunBoxFreeDrawTimestamp == 0 {
			out.NextFreeSunBox = chest_dat.FATE_BOX_INIT_CD_TIME
		}
	} else {
		out.NextFreeSunBox = chest_dat.FREE_FATE_BOX_CD - int32(freeBoxCD)
	}

	freeBoxCD = nowTime - playerFateBoxState.HunyuanBoxFreeDrawTimestamp
	if freeBoxCD >= chest_dat.FREE_FATE_BOX_CD {
		out.NextFreeHunyuanBox = 0
		if playerFateBoxState.HunyuanBoxFreeDrawTimestamp == 0 {
			out.NextFreeHunyuanBox = chest_dat.FATE_BOX_INIT_CD_TIME
		}
	} else {
		out.NextFreeHunyuanBox = chest_dat.FREE_FATE_BOX_CD - int32(freeBoxCD)
	}
}

func openFateBox(session *net.Session, boxType int32, times uint8, out *draw_api.OpenFateBox_Out) {
	state := module.State(session)
	//功能开启检查
	fail.When(!module.Player.IsOpenFunc(state.Database, player_dat.FUNC_FATE_BOX), "命锁宝箱功能未开启")

	playerFateBoxState := state.Database.Lookup.PlayerFateBoxState(state.PlayerId)
	fateBoxDat := chest_dat.GetFateBoxById(boxType)
	//宝箱开启检查
	fail.When(playerFateBoxState.Lock < fateBoxDat.AwardLock, "未通关命锁关卡")
	var totalDrawTime *int32
	var lastestDrawTimestamp *int64
	var tlogMoneyFlowType, xdlogMoneyEventType int32
	switch boxType {
	case chest_dat.StarBox:
		//star_box_draw_count
		totalDrawTime = &playerFateBoxState.StarBoxDrawCount
		lastestDrawTimestamp = &playerFateBoxState.StarBoxFreeDrawTimestamp
		if times == 1 {
			tlogMoneyFlowType = tlog.MFR_STAR_FATE_BOX_ONE
			xdlogMoneyEventType = xdlog.ET_STAR_FATE_BOX_ONE
		} else {
			tlogMoneyFlowType = tlog.MFR_STAR_FATE_BOX_TEN
			xdlogMoneyEventType = xdlog.ET_STAR_FATE_BOX_TEN
		}
	case chest_dat.MoonBox:
		totalDrawTime = &playerFateBoxState.MoonBoxDrawCount
		lastestDrawTimestamp = &playerFateBoxState.MoonBoxFreeDrawTimestamp
		if times == 1 {
			tlogMoneyFlowType = tlog.MFR_MOON_FATE_BOX_ONE
			xdlogMoneyEventType = xdlog.ET_MOON_FATE_BOX_ONE
		} else {
			tlogMoneyFlowType = tlog.MFR_MOON_FATE_BOX_TEN
			xdlogMoneyEventType = xdlog.ET_MOON_FATE_BOX_TEN
		}
	case chest_dat.SunBox:
		totalDrawTime = &playerFateBoxState.SunBoxDrawCount
		lastestDrawTimestamp = &playerFateBoxState.SunBoxFreeDrawTimestamp
		if times == 1 {
			tlogMoneyFlowType = tlog.MFR_SUN_FATE_BOX_ONE
			xdlogMoneyEventType = xdlog.ET_SUN_FATE_BOX_ONE
		} else {
			tlogMoneyFlowType = tlog.MFR_SUN_FATE_BOX_TEN
			xdlogMoneyEventType = xdlog.ET_SUN_FATE_BOX_TEN
		}
	case chest_dat.HunyuanBox:
		totalDrawTime = &playerFateBoxState.HunyuanBoxDrawCount
		lastestDrawTimestamp = &playerFateBoxState.HunyuanBoxFreeDrawTimestamp
		if times == 1 {
			tlogMoneyFlowType = tlog.MFR_HUNYUAN_FATE_BOX_ONE
			xdlogMoneyEventType = xdlog.ET_HUNYUAN_FATE_BOX_ONE
		} else {
			tlogMoneyFlowType = tlog.MFR_HUNYUAN_FATE_BOX_TEN
			xdlogMoneyEventType = xdlog.ET_HUNYUAN_FATE_BOX_TEN
		}
	default:
		panic("undefine fate box")
	}
	nowTime := time.GetNowTime()
	var isTenDraw, isFree bool
	if times == 1 {
		freeBoxCD := nowTime - *lastestDrawTimestamp
		if freeBoxCD >= chest_dat.FREE_FATE_BOX_CD { //免费？
			isFree = true
			if *lastestDrawTimestamp == 0 {
				//首次抽奖CD时间固定为18小时，那么把时间戳设置为当前时间的30个小时以前
				*lastestDrawTimestamp = (nowTime - chest_dat.FREE_FATE_BOX_CD + chest_dat.FREE_FATE_BOX_FIRST_TIME_CD)
			} else {
				*lastestDrawTimestamp = nowTime
			}
		} else {
			module.Player.DecMoney(state.Database, state.MoneyState, chest_dat.FATE_BOX_PRICE, player_dat.INGOT, tlogMoneyFlowType, xdlogMoneyEventType)
		}
	} else if times == 10 {
		isTenDraw = true
		module.Player.DecMoney(state.Database, state.MoneyState, chest_dat.FATE_BOX_TEN_TIME_PRICE, player_dat.INGOT, tlogMoneyFlowType, xdlogMoneyEventType)

	} else {
		fail.When(true, "unsupport")
	}

	var awardItems []*chest_dat.ChestItem
	FixedProp_Num := times
	for times > 0 {
		times--
		if *totalDrawTime < chest_dat.SPECIAL_DRAW_COUNT {
			chests := chest_dat.GetChestFixAward(-int8(boxType), *totalDrawTime+1)
			awardItems = append(awardItems, drawChest(chests))
		} else {
			chests := chest_dat.GetChestsByType(int8(boxType))
			awardItems = append(awardItems, drawChest(chests))
		}
		*totalDrawTime++
	}
	tlog.PlayerSystemModuleFlowLog(state.Database, tlog.SMT_FATE_BOX)
	var chest_list tlog.ItemList
	moneyReason, itemReason, xdEventType, tlogChestType, moneyNum := _getTlogInfoByFateBoxType(boxType, isTenDraw, isFree)
	rareItem := make(map[int8]map[int16]bool)
	if which_branch_dat.WHICH_BRANCH == which_branch_dat.TENCENT {
		db := state.Database
		if fateBoxDat.FixedProp > 0 && FixedProp_Num > 0 {
			module.Item.AddItem(db, fateBoxDat.FixedProp, int16(FixedProp_Num), itemReason, xdEventType, "")
		}
	}
	for _, award := range awardItems {
		//1.奖励物品
		awardItemForFateBox(session, award, moneyReason, itemReason, xdEventType)
		chest_list.ChestItem = append(chest_list.ChestItem, tlog.TChestItem{
			Id:       award.ItemId,
			Itemtype: award.Type,
			Num:      award.ItemNum,
		})

		// 2.设置返回结果
		out.Items = append(out.Items, draw_api.OpenFateBox_Out_Items{
			ItemType: draw_api.ItemType(award.Type),
			ItemId:   award.ItemId,
			Num:      award.ItemNum,
		})

		//3. 设置稀有物品
		setRareItem(award, rareItem)
	}
	broadcastRareItem(state, rareItem)
	tlog.PlayerChestDrawFlowLog(state.Database, tlog.MT_INGOT, moneyNum, tlogChestType, chest_list)
	state.Database.Update.PlayerFateBoxState(playerFateBoxState)
	if isTenDraw {
		eventId := int16(event_dat.EVENT_TEN_DRAW)
		eventInfo, ok := event_dat.GetEventInfoById(eventId)
		if ok && event_dat.CheckEventTime(eventInfo, event_dat.NOT_END) {
			//活动进行中
			eventRecord := state.EventsState.GetPlayerEventInfoById(eventId)
			state.EventsState.UpdateMax(state.Database, eventId, eventRecord.MaxAward+1)
			if event_dat.IsAccessNewTenDrawAward(int16(eventRecord.MaxAward)) {
				session.Send(&notify_api.SendEventCenterChange_Out{})
			}
		}
		module.Event.UpdateJsonEventTenDraw(session)
	}
}

func awardItemForFateBox(session *net.Session, awardItem *chest_dat.ChestItem, moneyReason, itemReason, xdEventType int32) {
	state := module.State(session)
	db := state.Database
	switch draw_api.ItemType(awardItem.Type) {
	case draw_api.ITEM_TYPE_COIN:
		module.Player.IncMoney(db, state.MoneyState, int64(awardItem.ItemNum), player_dat.COINS, moneyReason, xdEventType, "")

	case draw_api.ITEM_TYPE_INGOT:
		module.Player.IncMoney(db, state.MoneyState, int64(awardItem.ItemNum), player_dat.INGOT, moneyReason, xdEventType, "")

	case draw_api.ITEM_TYPE_GHOST:
		addGhost(state, awardItem.ItemId, itemReason, xdEventType)

	case draw_api.ITEM_TYPE_SWORD_SOUL:
		addSwordSoul(state, awardItem.ItemId, itemReason)

	case draw_api.ITEM_TYPE_PET:
		addPet(state, int32(awardItem.ItemId), awardItem.ItemNum, itemReason, xdEventType)

	case draw_api.ITEM_TYPE_ITEM, draw_api.ITEM_TYPE_GHOST_FRAGMENT, draw_api.ITEM_TYPE_PREFERENCE, draw_api.ITEM_TYPE_EQUIPMENT:
		module.Item.AddItem(db, awardItem.ItemId, int16(awardItem.ItemNum), itemReason, xdEventType, "")

	default:
		fail.When(true, "wrong item type")
	}
}

func setRareItem(award *chest_dat.ChestItem, rareItem map[int8]map[int16]bool) {
	set := false
	var typeId int8 = -1
	switch award.Type {
	case int8(draw_api.ITEM_TYPE_EQUIPMENT):
		itemDat := item_dat.GetItem(award.ItemId)
		if itemDat.IsRareItem() {
			set = true
			typeId = int8(draw_api.ITEM_TYPE_EQUIPMENT)
		}
	case int8(draw_api.ITEM_TYPE_GHOST_FRAGMENT):
		itemDat := item_dat.GetItem(award.ItemId)
		if itemDat.IsRareItem() {
			set = true
			typeId = int8(draw_api.ITEM_TYPE_GHOST_FRAGMENT)
		}
	}
	if set {
		if rareItem == nil {
			rareItem = make(map[int8]map[int16]bool)
		}
		if rareItem[typeId] == nil {
			rareItem[typeId] = make(map[int16]bool)
		}
		rareItem[typeId][award.ItemId] = true
	}
}

func broadcastRareItem(state *module.SessionState, rareItem map[int8]map[int16]bool) {
	var msgTpls []channel_dat.MessageTpl
	for typeId, awardList := range rareItem {
		for itemId, _ := range awardList {
			switch typeId {
			case int8(draw_api.ITEM_TYPE_EQUIPMENT):
				itemDat := item_dat.GetItem(itemId)
				msgTpls = append(msgTpls, channel_dat.MessageFateBoxEquipment{
					Player: channel_dat.ParamPlayer{state.PlayerNick, state.PlayerId},
					Item:   channel_dat.ParamItem{mail_dat.ATTACHMENT_ITEM, itemDat.Id},
				})
			case int8(draw_api.ITEM_TYPE_GHOST_FRAGMENT):
				itemDat := item_dat.GetItem(itemId)
				msgTpls = append(msgTpls, channel_dat.MessageFateBoxGhostFrame{
					Player: channel_dat.ParamPlayer{state.PlayerNick, state.PlayerId},
					Item:   channel_dat.ParamItem{mail_dat.ATTACHMENT_ITEM, itemDat.Id},
				})
			}
		}
	}
	rpc.RemoteAddWorldChannelTplMessage(state.PlayerId, msgTpls)
}
