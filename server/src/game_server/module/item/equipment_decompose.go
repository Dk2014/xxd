package item

import (
	"core/fail"
	"game_server/api/protocol/item_api"
	"game_server/dat/item_dat"
	"game_server/dat/player_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/tlog"
	"game_server/xdlog"
)

func decompose(state *module.SessionState, playerItemId int64) {
	db := state.Database

	playerItem := db.Lookup.PlayerItem(playerItemId)
	item := item_dat.GetItem(playerItem.ItemId)
	playerItemAppendix := db.Lookup.PlayerItemAppendix(playerItem.AppendixId)

	equipmentDecompose := item_dat.GetEquipmentDecompose(item.Quality, item.Level)

	//计算精炼返还铜钱
	coinsRefine := float64(calcRefund(item, playerItem.RefineLevel)) * item_dat.DECOMPOSE_COINS_REFINE_RATE
	refundCoins := int64(coinsRefine)

	// 附加
	var /*addNumForRefine, */ addNumForRecast int16
	// 精炼附加
	//addNumForRefine += int16(playerItem.RefineLevel)

	// 重铸附加
	if playerItemAppendix != nil && playerItemAppendix.RecastAttr != int8(item_api.ATTRIBUTE_NULL) {
		addNumForRecast = 1
		//重铸需要增加返还金钱
		recastDat := item_dat.GetEquipmentRecast(item.Quality, item.Level)
		refundCoins += int64(float64(recastDat.ConsumeCoin) * item_dat.DECOMPOSE_COINS_REFUND_RATE)
	}

	// 移除此装备
	delItemById(db, playerItemId, 0, tlog.IFR_ITEM_DECOMPOSE, xdlog.ET_ITEM_DECOMPOSE)

	//分解获得物品集合
	decomposeItems := make(map[int16]int16)

	//获得结晶
	//var crystalId int16 = equipmentDecompose.CrystalId
	//decomposeItems[crystalId] += equipmentDecompose.CrystalNum // + addNumForRefine

	//获得龙珠
	var dragonBallId int16 = equipmentDecompose.DragonBall
	if dragonBallId > 0 {
		decomposeItems[dragonBallId] += equipmentDecompose.DragonBallNum
	}

	//获得碎片
	var fragmentId int16
	switch item.TypeId {
	case item_dat.TYPE_WEAPON:
		fragmentId = item_dat.ITEM_WEAPON_FRAGMENT
	case item_dat.TYPE_CLOTHES:
		fragmentId = item_dat.ITEM_CLOTHES_FRAGMENT
	case item_dat.TYPE_SHOE:
		fragmentId = item_dat.ITEM_SHOE_FRAGMENT
	case item_dat.TYPE_ACCESSORIES:
		fragmentId = item_dat.ITEM_ACCESSORIES_FRAGMENT
	default:
		fail.When(true, "wrong type")
	}

	decomposeItems[fragmentId] += equipmentDecompose.FragmentNum + addNumForRecast

	var /*crystacount, */ fragmentcount int32
	db.Select.PlayerItem(func(row *mdb.PlayerItemRow) {
		if fragmentId > 0 {
			if row.ItemId() == fragmentId && row.BuyBackState() == 0 {
				fragmentcount += int32(row.Num())
			}
		}
		//if crystalId > 0 {
		//	if row.ItemId() == crystalId && row.BuyBackState() == 0 {
		//		crystacount += int32(row.Num())
		//	}
		//}
	})

	batchAddItem(db, decomposeItems, tlog.IFR_ITEM_DECOMPOSE, xdlog.ET_ITEM_DECOMPOSE)
	module.Player.IncMoney(db, state.MoneyState, refundCoins, player_dat.COINS, tlog.MFR_EQUIPMENT_REFINE, xdlog.ET_ITEM_DECOMPOSE, "")

	tlog.PlayerSystemModuleFlowLog(db, tlog.SMT_EQUIP_DECOMPOSE)
	tlog.PlayerEquipDecomposeFlowLog(db, int32(playerItem.ItemId), 0 /*int32(crystalId)*/, 0 /*crystacount*/, 0 /*int32(decomposeItems[crystalId])+crystacount*/, int32(fragmentId), fragmentcount, int32(decomposeItems[fragmentId])+fragmentcount)

}

func calcRefund(item *item_dat.Item, refineLv int16) (refundCoins int64) {
	equipmentRefine := item_dat.GetEquipmentRefineDat(int8(item.TypeId), int8(item.EquipTypeId))

	//计算装备附加属返还金钱
	for i := 1; i <= int(refineLv); i++ {
		refundCoins += int64(equipmentRefine.BasePrice + int32(i)*equipmentRefine.IncrePrice)
	}

	// switch refineLv {
	// case item_dat.EQUIPMENT_REFINE_LEVEL10:
	// 	refundCoins += int64(float64(equipmentRefine.Level10ConsumeCoin) * item_dat.DECOMPOSE_COINS_REFUND_RATE)
	// 	fallthrough
	// case item_dat.EQUIPMENT_REFINE_LEVEL9:
	// 	refundCoins += int64(float64(equipmentRefine.Level9ConsumeCoin) * item_dat.DECOMPOSE_COINS_REFUND_RATE)
	// 	fallthrough
	// case item_dat.EQUIPMENT_REFINE_LEVEL8:
	// 	refundCoins += int64(float64(equipmentRefine.Level8ConsumeCoin) * item_dat.DECOMPOSE_COINS_REFUND_RATE)
	// 	fallthrough
	// case item_dat.EQUIPMENT_REFINE_LEVEL7:
	// 	refundCoins += int64(float64(equipmentRefine.Level7ConsumeCoin) * item_dat.DECOMPOSE_COINS_REFUND_RATE)
	// 	fallthrough
	// case item_dat.EQUIPMENT_REFINE_LEVEL6:
	// 	refundCoins += int64(float64(equipmentRefine.Level6ConsumeCoin) * item_dat.DECOMPOSE_COINS_REFUND_RATE)
	// 	fallthrough
	// case item_dat.EQUIPMENT_REFINE_LEVEL5:
	// 	refundCoins += int64(float64(equipmentRefine.Level5ConsumeCoin) * item_dat.DECOMPOSE_COINS_REFUND_RATE)
	// 	fallthrough
	// case item_dat.EQUIPMENT_REFINE_LEVEL4:
	// 	refundCoins += int64(float64(equipmentRefine.Level4ConsumeCoin) * item_dat.DECOMPOSE_COINS_REFUND_RATE)
	// 	fallthrough
	// case item_dat.EQUIPMENT_REFINE_LEVEL3:
	// 	refundCoins += int64(float64(equipmentRefine.Level3ConsumeCoin) * item_dat.DECOMPOSE_COINS_REFUND_RATE)
	// 	fallthrough
	// case item_dat.EQUIPMENT_REFINE_LEVEL2:
	// 	refundCoins += int64(float64(equipmentRefine.Level2ConsumeCoin) * item_dat.DECOMPOSE_COINS_REFUND_RATE)
	// 	fallthrough
	// case item_dat.EQUIPMENT_REFINE_LEVEL1:
	// 	refundCoins += int64(float64(equipmentRefine.Level1ConsumeCoin) * item_dat.DECOMPOSE_COINS_REFUND_RATE)
	// }
	return refundCoins
}
