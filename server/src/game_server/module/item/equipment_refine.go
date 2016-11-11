package item

import (
	//"core/fail"
	"core/net"
	"game_server/api/protocol/item_api"
	"game_server/dat/item_dat"
	"game_server/dat/player_dat"
	"game_server/dat/quest_dat"
	"game_server/module"
	"game_server/tlog"
	"game_server/xdlog"
	//"math/rand"
)

func Refine(session *net.Session, in *item_api.Refine_In, out *item_api.Refine_Out) {
	state := module.State(session)
	db := state.Database
	player := db.Lookup.Player(db.PlayerId())
	role := db.Lookup.PlayerRole(player.MainRoleId)
	var times int16 = 1 // 强化次数
	if in.IsBatch {
		times = item_dat.BATCH_REFINE_TIMES
	}
	tlog.PlayerSystemModuleFlowLog(db, tlog.SMT_EQUIP_REFINE)
	costinfo := tlog.EquipRefineList{}
	playerItemId := in.Id
	playerItem := db.Lookup.PlayerItem(playerItemId)

	//如果装备等级已达到满级则直接return
	if playerItem.RefineLevel >= role.Level {
		return
	}
	item := item_dat.GetItem(playerItem.ItemId)
	if playerItem.RefineLevel < role.Level && playerItem.RefineLevel+times > role.Level {
		times = role.Level - playerItem.RefineLevel
	}
	dstRefineLevel := playerItem.RefineLevel + times // 目标等级

	//fail.When(dstRefineLevel > role.Level, "refine level is full")

	// // 消耗
	// equipmentRefine := item_dat.GetEquipmentRefine(item.Quality, item.Level)

	// // 消耗部位碎片
	// fragmentNum := equipmentRefine.FragmentNum
	// var fragmentId int16
	// switch item.TypeId {
	// case item_dat.TYPE_WEAPON:
	// 	fragmentId = item_dat.ITEM_WEAPON_FRAGMENT
	// case item_dat.TYPE_CLOTHES:
	// 	fragmentId = item_dat.ITEM_CLOTHES_FRAGMENT
	// case item_dat.TYPE_SHOE:
	// 	fragmentId = item_dat.ITEM_SHOE_FRAGMENT
	// case item_dat.TYPE_ACCESSORIES:
	// 	fragmentId = item_dat.ITEM_ACCESSORIES_FRAGMENT
	// }

	// db.Select.PlayerItem(func(row *mdb.PlayerItemRow) {
	// 	if fragmentNum > 0 {
	// 		if row.ItemId() == fragmentId && row.BuyBackState() == 0 {
	// 			costinfo.Fragmentcount += int32(row.Num() * times)
	// 		}
	// 	}
	// 	if equipmentRefine.BlueCrystalNum > 0 {
	// 		if row.ItemId() == item_dat.ITEM_BLUE_CRYSTAL && row.BuyBackState() == 0 {
	// 			costinfo.Bluecount += int32(row.Num() * times)
	// 		}
	// 	}
	// 	if equipmentRefine.PurpleCrystalNum > 0 {
	// 		if row.ItemId() == item_dat.ITEM_PURPLE_CRYSTAL && row.BuyBackState() == 0 {
	// 			costinfo.Purplecount += int32(row.Num() * times)
	// 		}
	// 	}
	// 	if equipmentRefine.GoldenCrystalNum > 0 {
	// 		if row.ItemId() == item_dat.ITEM_GOLDEN_CRYSTAL && row.BuyBackState() == 0 {
	// 			costinfo.Goldcount += int32(row.Num() * times)
	// 		}
	// 	}
	// 	if equipmentRefine.OrangeCrystalNum > 0 {
	// 		if row.ItemId() == item_dat.ITEM_ORANGE_CRYSTAL && row.BuyBackState() == 0 {
	// 			costinfo.Orangecount += int32(row.Num() * times)
	// 		}
	// 	}
	// })
	// if fragmentNum > 0 {
	// 	costinfo.Fragmentid = int32(fragmentId)
	// 	costinfo.Fragmentaftercount = costinfo.Fragmentcount - int32(fragmentNum)
	// }
	// module.Item.DelItemByItemId(db, fragmentId, fragmentNum, tlog.IFR_EQUIPMENT_REFINE)

	// // 消耗结晶
	// // 蓝色
	// if equipmentRefine.BlueCrystalNum > 0 {
	// 	costinfo.Blueaftercount = costinfo.Bluecount - int32(equipmentRefine.BlueCrystalNum)
	// 	module.Item.DelItemByItemId(db, item_dat.ITEM_BLUE_CRYSTAL, equipmentRefine.BlueCrystalNum, tlog.IFR_EQUIPMENT_REFINE)
	// }
	// // 紫色
	// if equipmentRefine.PurpleCrystalNum > 0 {
	// 	costinfo.Purpleaftercount = costinfo.Purplecount - int32(equipmentRefine.PurpleCrystalNum)
	// 	module.Item.DelItemByItemId(db, item_dat.ITEM_PURPLE_CRYSTAL, equipmentRefine.PurpleCrystalNum, tlog.IFR_EQUIPMENT_REFINE)
	// }
	// // 金色
	// if equipmentRefine.GoldenCrystalNum > 0 {
	// 	costinfo.Goldaftercount = costinfo.Goldcount - int32(equipmentRefine.GoldenCrystalNum)
	// 	module.Item.DelItemByItemId(db, item_dat.ITEM_GOLDEN_CRYSTAL, equipmentRefine.GoldenCrystalNum, tlog.IFR_EQUIPMENT_REFINE)
	// }
	// // 橙色
	// if equipmentRefine.OrangeCrystalNum > 0 {
	// 	costinfo.Orangeaftercount = costinfo.Orangecount - int32(equipmentRefine.OrangeCrystalNum)
	// 	module.Item.DelItemByItemId(db, item_dat.ITEM_ORANGE_CRYSTAL, equipmentRefine.OrangeCrystalNum, tlog.IFR_EQUIPMENT_REFINE)
	// }

	// 消耗铜钱
	var consumeCoin int64
	for i := playerItem.RefineLevel + 1; i <= dstRefineLevel; i++ {
		refineDat := item_dat.GetEquipmentRefineDat(int8(item.TypeId), int8(item.EquipTypeId))
		consumeCoin += int64(refineDat.BasePrice + int32(i)*refineDat.IncrePrice)
	}
	playerInfo := db.Lookup.PlayerInfo(db.PlayerId())
	imoney := int32(playerInfo.Coins)
	afterimoney := imoney - int32(consumeCoin)
	if afterimoney < 0 {
		out.Code = 1
		return //防止由于操作过快导致的钱不够的情况
	}
	module.Player.DecMoney(db, state.MoneyState, consumeCoin, player_dat.COINS, tlog.MFR_EQUIPMENT_REFINE, xdlog.ET_EQUIPMENT_REFINE)
	playerItem.Price += int32(consumeCoin)

	beforeRefineLevel := int32(playerItem.RefineLevel)

	// 更新玩家物品数据，精炼等级为目标等级
	playerItem.RefineLevel = dstRefineLevel
	playerItem.RefineFailTimes = 0

	out.Level = dstRefineLevel
	out.Id = playerItem.Id

	// 更新玩家物品数据
	db.Update.PlayerItem(playerItem)

	tlog.PlayerEquipRefineFlowLog(db, int32(playerItem.ItemId), tlog.MT_COIN, imoney, afterimoney, beforeRefineLevel, int32(playerItem.RefineLevel), costinfo)

	for i := int16(1); i <= times; i++ {
		module.Quest.RefreshDailyQuest(state, quest_dat.DAILY_QUEST_CLASS_REFINE_EQ)
	}
	return
}
