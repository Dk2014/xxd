package item

import (
	"core/fail"
	"core/net"
	"core/time"
	"encoding/json"
	"game_server/api/protocol/item_api"
	"game_server/api/protocol/notify_api"
	//"game_server/dat/battle_pet_dat"
	"game_server/dat/fashion_dat"
	"game_server/dat/item_dat"
	"game_server/dat/physical_dat"
	"game_server/dat/player_dat"
	"game_server/dat/role_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/tlog"
	"game_server/xdlog"
	"math/rand"
)

func useItem(session *net.Session, playerItemId int64, num int32, isBatchUse bool) {
	state := module.State(session)
	db := state.Database

	playerItem := db.Lookup.PlayerItem(playerItemId)
	item := item_dat.GetItem(playerItem.ItemId)
	var changedItemId []int64

	switch item.TypeId {
	case item_dat.TYPE_CHEST:
		// 物品宝箱
		fail.When(item.FuncId > 0 && !module.Player.IsOpenFunc(db, int32(item.FuncId)), "needed func is not open")

		// 消耗宝箱
		if !isBatchUse {
			openItemChest(state, playerItem.ItemId, 1)
			changedItemId = delItemByItemId(db, item.Id, 1, tlog.IFR_OPEN_ITEM_CHEST, xdlog.ET_OPEN_ITEM_CHEST)
		} else {
			openItemChest(state, playerItem.ItemId, num)
		}
	case item_dat.TYPE_FASHION:
		fashionId := fashion_dat.GetFashionIdByItemId(item.Id)
		mainRoleLevel := int32(module.Role.GetMainRole(db).Level)
		fail.When(mainRoleLevel < item.Level, "未达到时装穿戴等级")
		playerFashion := module.Fashion.AddFashion(db, fashionId, int64(item.ValidHours))
		module.Notify.SendFashionChange(session, playerFashion)
		changedItemId = delItemByItemId(db, item.Id, 1, tlog.IFR_GET_FASHION, xdlog.ET_USE_ITEM)
	case item_dat.TYPE_DRAGON_BALL:
		dragonBallConfigs := item_dat.GetDragonBallConfig(item.Id)
		//TODO 随即获得物品受概率配置的概率影响
		index := rand.Intn(len(dragonBallConfigs))
		awardItem := dragonBallConfigs[index]
		addItem(db, awardItem.ItemId, awardItem.ItemNum, tlog.IFR_USE_DRAGON_BALL, xdlog.ET_USE_ITEM, "")
		changedItemId = delItemByItemId(db, item.Id, item_dat.CONSUME_DRAGON_BALL_NUM, tlog.IFR_USE_DRAGON_BALL, xdlog.ET_USE_ITEM)
		session.Send(&item_api.DragonBallExchangeNotify_Out{
			ItemId:  awardItem.ItemId,
			ItemNum: awardItem.ItemNum,
		})
	case item_dat.TYPE_ACT_REFLECT:
		changedItemId = delItemByItemId(db, item.Id, 1, 0, xdlog.ET_USE_ITEM)
		magicItem(state, item.Id)
	default:
		fail.When(true, "wrong item type")
	}
	if !isBatchUse {
		changed := false
		for _, id := range changedItemId {
			if id == playerItemId {
				changed = true
				break
			}
		}
		session.Send(&item_api.UseItem_Out{
			Origin:  playerItemId,
			Changed: changed,
		})
	}
}

// 开启物品宝箱
func openItemChest(state *module.SessionState, itemId int16, total_num int32) {
	db := state.Database
	out := &notify_api.ItemChange_Out{}
	out.Items = make([]notify_api.ItemChange_Out_Items, 0, item_dat.MAX_BAG_NUM)

	itemBoxContents := item_dat.GetItemBoxContent(itemId)
	for _, itemBoxContent := range itemBoxContents {
		switch itemBoxContent.Mode {
		case ITEM_BOX_CONTENT_MODE_DIRECT:
			//直接获得
			typeId := itemBoxContent.Type

			switch typeId {
			case ITEM_BOX_CONTENT_TYPE_COIN:
				num := int64(itemBoxContent.MinNum)
				module.Player.IncMoney(db, state.MoneyState, num*int64(total_num), player_dat.COINS, tlog.MFR_OPEN_ITEM_CHEST, xdlog.ET_OPEN_ITEM_CHEST, "")
			case ITEM_BOX_CONTENT_TYPE_INGOT:
				num := int64(itemBoxContent.MinNum)
				module.Player.IncMoney(db, state.MoneyState, num*int64(total_num), player_dat.INGOT, tlog.MFR_OPEN_ITEM_CHEST, xdlog.ET_OPEN_ITEM_CHEST, "")
			case ITEM_BOX_CONTENT_TYPE_ITEM:
				getItemId := itemBoxContent.GetItemId
				num := int16(itemBoxContent.MinNum)
				addItemAndSetResp(db, getItemId, num*int16(total_num), out, tlog.IFR_OPEN_ITEM_CHEST, xdlog.ET_OPEN_ITEM_CHEST, "")
			case ITEM_BOX_CONTENT_TYPE_EXP:
				num := int64(itemBoxContent.MinNum)
				mainRole := module.Role.GetMainRole(db)
				module.Role.AddRoleExp(db, mainRole.RoleId, num*int64(total_num), mainRole.RoleId, tlog.EFT_OPEN_ITEM_CHEST)
			case ITEM_BOX_CONTENT_TYPE_HEART:
				num := int16(itemBoxContent.MinNum)
				module.Heart.IncHeart(state, num*int16(total_num))
			case ITEM_BOX_CONTENT_TYPE_PHYSICAL:
				num := int16(itemBoxContent.MinNum)
				module.Physical.AwardIncrease(state, num*int16(total_num), tlog.PFR_OPEN_ITEM_CHEST)
			case ITEM_BOX_CONTENT_TYPE_GHOST:
				for index := 0; index < int(total_num); index++ {
					module.Ghost.AddGhost(state, itemBoxContent.GetItemId, tlog.IFR_OPEN_ITEM_CHEST, xdlog.ET_OPEN_ITEM_CHEST)
				}
			case ITEM_BOX_CONTENT_TYPE_SWORD_SOUL:
				for index := 0; index < int(total_num); index++ {
					module.SwordSoul.AddSwordSoul(state, itemBoxContent.GetItemId, tlog.IFR_OPEN_ITEM_CHEST)
				}
			case ITEM_BOX_CONTENT_TYPE_PET:
				module.BattlePet.AddPet(state.Database, int32(itemBoxContent.GetItemId), tlog.IFR_OPEN_ITEM_CHEST, xdlog.ET_OPEN_ITEM_CHEST)
			case ITEM_BOX_CONTENT_TYPE_FRAGMENT:
				num := int64(itemBoxContent.MinNum)
				module.Player.IncSwordSoulFragment(state.Database, num*int64(total_num), player_dat.SWORDSOULFRAGMENT, tlog.MFR_OPEN_ITEM_CHEST, xdlog.ET_OPEN_ITEM_CHEST)
			}

		case ITEM_BOX_CONTENT_MODE_RANDOM_NUM:
			// 随机数目
			typeId := itemBoxContent.Type

			min := int(itemBoxContent.MinNum)
			max := int(itemBoxContent.MaxNum)

			total := 0
			for index := 0; index < int(total_num); index++ {
				total += rand.Intn(max-min+1) + min
			}
			if total <= 0 {
				continue
			}

			switch typeId {
			case ITEM_BOX_CONTENT_TYPE_COIN:
				module.Player.IncMoney(db, state.MoneyState, int64(total), player_dat.COINS, tlog.MFR_OPEN_ITEM_CHEST, xdlog.ET_OPEN_ITEM_CHEST, "")
			case ITEM_BOX_CONTENT_TYPE_INGOT:
				module.Player.IncMoney(db, state.MoneyState, int64(total), player_dat.INGOT, tlog.MFR_OPEN_ITEM_CHEST, xdlog.ET_OPEN_ITEM_CHEST, "")
			case ITEM_BOX_CONTENT_TYPE_ITEM:
				getItemId := itemBoxContent.GetItemId
				addItemAndSetResp(db, getItemId, int16(total), out, tlog.IFR_OPEN_ITEM_CHEST, xdlog.ET_OPEN_ITEM_CHEST, "")
			case ITEM_BOX_CONTENT_TYPE_EXP:
				mainRole := module.Role.GetMainRole(db)
				module.Role.AddRoleExp(db, mainRole.RoleId, int64(total), mainRole.RoleId, tlog.EFT_OPEN_ITEM_CHEST)
			case ITEM_BOX_CONTENT_TYPE_FRAGMENT:
				module.Player.IncSwordSoulFragment(state.Database, int64(total), player_dat.SWORDSOULFRAGMENT, tlog.MFR_OPEN_ITEM_CHEST, xdlog.ET_OPEN_ITEM_CHEST)
			}
		case ITEM_BOX_CONTENT_MODE_MAY_GET:
			// 随机获取

			total := 0
			for index := 0; index < int(total_num); index++ {
				probability := int8(rand.Intn(100))

				if probability < itemBoxContent.Probability {
					total++
				}
			}
			if total <= 0 {
				continue
			}
			typeId := itemBoxContent.Type

			switch typeId {
			case ITEM_BOX_CONTENT_TYPE_COIN:
				num := int64(itemBoxContent.MinNum)
				module.Player.IncMoney(db, state.MoneyState, num*int64(total), player_dat.COINS, tlog.MFR_OPEN_ITEM_CHEST, xdlog.ET_OPEN_ITEM_CHEST, "")
			case ITEM_BOX_CONTENT_TYPE_INGOT:
				num := int64(itemBoxContent.MinNum)
				module.Player.IncMoney(db, state.MoneyState, num*int64(total), player_dat.INGOT, tlog.MFR_OPEN_ITEM_CHEST, xdlog.ET_OPEN_ITEM_CHEST, "")
			case ITEM_BOX_CONTENT_TYPE_ITEM:
				getItemId := itemBoxContent.GetItemId
				num := int16(itemBoxContent.MinNum)
				addItemAndSetResp(db, getItemId, num*int16(total), out, tlog.IFR_OPEN_ITEM_CHEST, xdlog.ET_OPEN_ITEM_CHEST, "")
			case ITEM_BOX_CONTENT_TYPE_EXP:
				num := int64(itemBoxContent.MinNum)
				mainRole := module.Role.GetMainRole(db)
				module.Role.AddRoleExp(db, mainRole.RoleId, num*int64(total), mainRole.RoleId, tlog.EFT_OPEN_ITEM_CHEST)
			case ITEM_BOX_CONTENT_TYPE_HEART:
				num := int16(itemBoxContent.MinNum)
				module.Heart.IncHeart(state, num*int16(total))
			case ITEM_BOX_CONTENT_TYPE_PHYSICAL:
				num := int16(itemBoxContent.MinNum)
				module.Physical.AwardIncrease(state, num*int16(total), tlog.PFR_OPEN_ITEM_CHEST)
			case ITEM_BOX_CONTENT_TYPE_GHOST:
				for index := 0; index < int(total); index++ {
					module.Ghost.AddGhost(state, itemBoxContent.GetItemId, tlog.IFR_OPEN_ITEM_CHEST, xdlog.ET_OPEN_ITEM_CHEST)
				}
			case ITEM_BOX_CONTENT_TYPE_SWORD_SOUL:
				//TODO 潜龙剑心一次获得20个
				for index := 0; index < int(total); index++ {
					module.SwordSoul.AddSwordSoul(state, itemBoxContent.GetItemId, tlog.IFR_OPEN_ITEM_CHEST)
				}
			case ITEM_BOX_CONTENT_TYPE_PET:
				module.BattlePet.AddPet(state.Database, int32(itemBoxContent.GetItemId), tlog.IFR_OPEN_ITEM_CHEST, xdlog.ET_OPEN_ITEM_CHEST)
			case ITEM_BOX_CONTENT_TYPE_FRAGMENT:
				num := int64(itemBoxContent.MinNum)
				module.Player.IncSwordSoulFragment(state.Database, num*int64(total), player_dat.SWORDSOULFRAGMENT, tlog.MFR_OPEN_ITEM_CHEST, xdlog.ET_OPEN_ITEM_CHEST)
			}

		case ITEM_BOX_CONTENT_MODE_RANDOM_ITEM:
			// 随机物品

			type Item struct {
				ItemId int16
				Num    int16
			}

			var items []*Item

			json.Unmarshal([]byte(itemBoxContent.ItemIdSet), &items)

			fail.When(len(items) == 0, "wrong item_id_set")
			for index := 0; index < int(total_num); index++ {
				index := rand.Intn(len(items))
				randItem := items[index]

				item := item_dat.GetItem(randItem.ItemId)
				fail.When(item.TypeId == item_dat.TYPE_RESOURCE, "wrong item type")

				addItemAndSetResp(db, randItem.ItemId, randItem.Num, out, tlog.IFR_OPEN_ITEM_CHEST, xdlog.ET_OPEN_ITEM_CHEST, "")
			}
		}
	}
	if session, ok := module.Player.GetPlayerOnline(db.PlayerId()); ok {
		module.Notify.SendItemChange(session, out)
	}
}

func roleUseItem(state *module.SessionState, roleId int8, id int64, isBatchUse bool) {
	playerItem := state.Database.Lookup.PlayerItem(id)
	fail.When(playerItem == nil, "roleUseItem. id error")

	item := item_dat.GetItem(playerItem.ItemId)
	fail.When(item.TypeId != item_dat.TYPE_COST_PROPS, "MUST cost item")
	if !isBatchUse {
		delItemByItemId(state.Database, item.Id, 1, tlog.IFR_ITEM_ROLE_USE, xdlog.ET_ITEM_ROLE_USE)
	}

	attr := item_dat.GetCostpropsWithItemId(item.Id)
	switch attr.Type {
	case item_dat.COST_ITEM_EXP:
		module.Role.AddRoleExp(state.Database, roleId, int64(attr.Value), state.RoleId, tlog.EFT_ITEM_ROLE_USE)

	case item_dat.COST_ITEM_PHY:
		module.Physical.AwardIncrease(state, int16(attr.Value), tlog.PFR_USE_ITEM)

	case item_dat.COST_ITEM_SHP:
		module.Team.IncRelationship(state.Database, attr.Value)

	default:
		fail.When(true, "error cost-item type")
	}
}

func BatchUseItem(session *net.Session, in *item_api.BatchUseItem_In) {
	state := module.State(session)
	playerItem := state.Database.Lookup.PlayerItem(in.Id)
	fail.When(playerItem == nil, "fail to found player item in batch use")
	item := item_dat.GetItem(playerItem.ItemId)
	fail.When(item.CanBatch == 0, "该物品无法使用批量使用")
	total := playerItem.Num
	if total > item_dat.BATCH_USE_MAX_NUM {
		//最大使用数量限制
		total = item_dat.BATCH_USE_MAX_NUM
	}
	var (
		flowReason  int32
		xdEventType int32
		used        int16 = 0 //使用掉得个数
	)

	switch item.TypeId {
	case item_dat.TYPE_CHEST:
		flowReason = tlog.IFR_OPEN_ITEM_CHEST
		xdEventType = xdlog.ET_OPEN_ITEM_CHEST
		useItem(session, in.Id, int32(total), true)
		used = total
	case item_dat.TYPE_COST_PROPS:
		var (
			role_id       int8  = in.RoleId
			mainRoleLevel int16 = 0                            //主角等级，经验消耗道具验证时使用
			isMainRole    bool  = role_dat.IsMainRole(role_id) //使用对象是否为主角
		)
		if !isMainRole {
			mainRoleLevel = module.Role.GetMainRole(state.Database).Level
		}
		attr := item_dat.GetCostpropsWithItemId(item.Id)
		flowReason = tlog.IFR_ITEM_ROLE_USE
		xdEventType = xdlog.ET_ITEM_ROLE_USE
		//添加判断限定
		//1.对伙伴增加经验，判断伙伴等级不能超过主角
		role := module.Role.GetBuddyRole(state.Database, in.RoleId)
		level, exp := role.Level, role.Exp
		if !isMainRole && attr.Type == item_dat.COST_ITEM_EXP {
			fail.When(level > mainRoleLevel, "伙伴等级越界")
			for index := 0; index < int(total); index++ {
				value := int64(attr.Value)
				if index == 0 {
					value = value + exp
				}
				level, exp = module.Role.GetNewLevelAndExp(level, value)
				used++
				if level >= mainRoleLevel {
					break
				}
			}
			module.Role.AddRoleExp(state.Database, role_id, int64(used)*int64(attr.Value), state.RoleId, tlog.EFT_ITEM_ROLE_USE)
		}

		//2.增加的是体力，不能超过最大值
		if attr.Type == item_dat.COST_ITEM_PHY {
			playerPhysical := state.Database.Lookup.PlayerPhysical(state.PlayerId)
			count := (physical_dat.MAX_PHYSICAL_VALUE - playerPhysical.Value) / int16(attr.Value)
			if count < total {
				used = count
				if (physical_dat.MAX_PHYSICAL_VALUE-playerPhysical.Value)%int16(attr.Value) != 0 {
					used++
				}
			} else {
				used = total
			}

			module.Physical.AwardIncrease(state, used*int16(attr.Value), tlog.PFR_USE_ITEM)
		}
		//ToAdd 其他更多条件限定添加....

	default:
		fail.When(true, "错误的批量使用道具类型")
	}
	delItemByItemId(state.Database, item.Id, used, flowReason, xdEventType)
}

func awardCornucopia(state *module.SessionState, pItem_id int64) (award_coins int64) {
	//首先检查，这家伙有没有聚宝盆类型功能道具
	db := state.Database
	pItem := db.Lookup.PlayerItem(pItem_id)
	fail.When(pItem == nil, "this guy require awarding money but don't have any item")
	item := item_dat.GetItem(pItem.ItemId)
	fail.When(item.TypeId != item_dat.TYPE_ACT_REFLECT || item.ActId != item_dat.ITEM_ACTREF_OPEN_CORNUCOPIA, "trying to require cornucopia with wrong item")

	//查出玩家的聚宝盆数据，没有则初始化一个
	cornucopiaInfo := db.Lookup.PlayerCornucopia(db.PlayerId())
	if cornucopiaInfo == nil {
		db.Insert.PlayerCornucopia(&mdb.PlayerCornucopia{
			Pid:        db.PlayerId(),
			OpenTime:   0,
			DailyCount: 0,
		})
		cornucopiaInfo = db.Lookup.PlayerCornucopia(db.PlayerId())
	}

	//检查当天使用次数
	if !time.IsToday(cornucopiaInfo.OpenTime) {
		cornucopiaInfo.DailyCount = 0
	}
	fail.When(cornucopiaInfo.DailyCount >= item_dat.CORNUCOPIA_MAX_DAILY_USING_AMOUNT, "cornucopia daily using time has reached the max limit")

	//奖励铜钱儿
	config := item_dat.GetItemReflectConfig(item.Id)[0]
	min := int(config.AwardCoinMin)
	max := int(config.AwardCoinMax)
	award_coins = int64(rand.Intn(max-min+1) + min)
	module.Player.IncMoney(db, state.MoneyState, award_coins, player_dat.COINS, tlog.MFR_OPEN_ITEM_CHEST, xdlog.ET_CORNUCOPIA, "")

	//更新使用次数和使用时间
	cornucopiaInfo.DailyCount++
	cornucopiaInfo.OpenTime = time.GetNowTime()
	db.Update.PlayerCornucopia(cornucopiaInfo)

	return
}

func cornucopiaCount(db *mdb.Database) int16 {
	cornucopiaInfo := db.Lookup.PlayerCornucopia(db.PlayerId())
	if cornucopiaInfo == nil {
		return 0
	} else {
		if !time.IsToday(cornucopiaInfo.OpenTime) {
			cornucopiaInfo.DailyCount = 0
			db.Update.PlayerCornucopia(cornucopiaInfo)
		}
		return cornucopiaInfo.DailyCount
	}
}

func magicItem(state *module.SessionState, itemId int16) {
	switch itemId {
	case item_dat.ITEM_SWORD_DRIVING_FLAG:
		drivingInfo := state.Database.Lookup.PlayerDrivingSwordInfo(state.PlayerId)
		if drivingInfo != nil {
			module.DrivingSword.ResetCloudData(state, drivingInfo.CurrentCloud)
		}
	default:
		fail.When(true, "undefind magic item")
	}
}
