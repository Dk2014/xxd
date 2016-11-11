package item

import (
	"game_server/api/protocol/item_api"
	"game_server/dat/item_dat"
	"game_server/mdb"
	"game_server/tlog"
	"game_server/xdlog"
	"math/rand"
)

//给装备添加附加属性
func addItemAppendix(db *mdb.Database, playerItem *mdb.PlayerItem) (playerItemAppendix *mdb.PlayerItemAppendix) {

	item := item_dat.GetItem(playerItem.ItemId)
	if item.TypeId != item_dat.TYPE_WEAPON && item.TypeId != item_dat.TYPE_ACCESSORIES &&
		item.TypeId != item_dat.TYPE_CLOTHES && item.TypeId != item_dat.TYPE_SHOE {
		return
	}

	if item.AppendixNum == 0 {
		return
	}

	playerItemAppendix = &mdb.PlayerItemAppendix{Pid: db.PlayerId()}
	randArray1 := []*int32{&playerItemAppendix.Health, &playerItemAppendix.Cultivation, &playerItemAppendix.Speed,
		&playerItemAppendix.Attack, &playerItemAppendix.DodgeLevel, &playerItemAppendix.HitLevel,
		&playerItemAppendix.BlockLevel, &playerItemAppendix.CriticalLevel, &playerItemAppendix.TenacityLevel,
		&playerItemAppendix.DestroyLevel, &playerItemAppendix.Defence}

	if item.Quality == item_dat.ITEM_QUALITY_ORANGE {
		//橙色装备特殊处理
		appendix := item_dat.GetSpecialEquipmentAppendix(int32(item.Id))
		randArray2 := []int32{appendix.Health, appendix.Cultivation, appendix.Speed, appendix.Attack, appendix.DodgeLevel,
			appendix.HitLevel, appendix.BlockLevel, appendix.CriticalLevel, appendix.TenacityLevel, appendix.DestroyLevel, appendix.Defence}
		for i := 0; i < len(randArray1); i++ {
			*randArray1[i] = randArray2[i]
		}
	} else {
		lower, upper := item_dat.GetEquipmentAppendix(item.AppendixLevel)
		//随机取值下限
		randArray2 := []int32{lower.Health, lower.Cultivation, lower.Speed, lower.Attack, lower.DodgeLevel,
			lower.HitLevel, lower.BlockLevel, lower.CriticalLevel, lower.TenacityLevel, lower.DestroyLevel, lower.Defence}
		//随机取值上限
		randArray3 := []int32{upper.Health, upper.Cultivation, upper.Speed, upper.Attack, upper.DodgeLevel,
			upper.HitLevel, upper.BlockLevel, upper.CriticalLevel, upper.TenacityLevel, upper.DestroyLevel, upper.Defence}
		//每个对应属性的随机单位
		randArray4 := []int32{50, 10, 10, 10, 5, 5, 5, 5, 5, 5, 5}

		//武器和饰品不会随机到防御属性
		num := 0
		if item.TypeId == item_dat.TYPE_WEAPON || item.TypeId == item_dat.TYPE_ACCESSORIES {
			num = 1
		}
		//随机出指定数量的附加属性
		for i := 0; i < int(item.AppendixNum); i++ {
			size := len(randArray1) - i - num
			index := rand.Int() % size
			*randArray1[index] = randArray2[index] + randArray4[index]*(rand.Int31()%((randArray3[index]-randArray2[index])/randArray4[index]+1))

			randArray1[index], randArray1[size-1] = randArray1[size-1], randArray1[index]
			randArray2[index], randArray2[size-1] = randArray2[size-1], randArray2[index]
			randArray3[index], randArray3[size-1] = randArray3[size-1], randArray3[index]
			randArray4[index], randArray4[size-1] = randArray4[size-1], randArray4[index]
		}
	}

	db.Insert.PlayerItemAppendix(playerItemAppendix)
	playerItem.AppendixId = playerItemAppendix.Id
	db.Update.PlayerItem(playerItem)
	return
}

func delItemFromDB(db *mdb.Database, playerItem *mdb.PlayerItem, itemFlowReason, xdlogEventType int32) {
	if playerItem.AppendixId > 0 {
		playerItemAppendix := db.Lookup.PlayerItemAppendix(playerItem.AppendixId)
		db.Delete.PlayerItemAppendix(playerItemAppendix)
	}

	if playerItem.BuyBackState == 0 {
		//买回状态删除，不记录itemflow
		item := item_dat.GetItem(playerItem.ItemId)
		tlog.PlayerItemFlowLog(db, item.Id, item.TypeId, playerItem.Num, tlog.REDUCE, itemFlowReason)
		xdlog.PropsLog(db, item.Id, playerItem.Num, xdlogEventType)
	}

	db.Delete.PlayerItem(playerItem)
}

func getAttrTypeBits(playerItemAppendix *mdb.PlayerItemAppendix) (attrTypeBits int) {

	if playerItemAppendix.Health > 0 {
		attrTypeBits |= 1 << uint(item_api.ATTRIBUTE_HEALTH)
	}

	if playerItemAppendix.Cultivation > 0 {
		attrTypeBits |= 1 << uint(item_api.ATTRIBUTE_CULTIVATION)
	}

	if playerItemAppendix.Speed > 0 {
		attrTypeBits |= 1 << uint(item_api.ATTRIBUTE_SPEED)
	}

	if playerItemAppendix.Attack > 0 {
		attrTypeBits |= 1 << uint(item_api.ATTRIBUTE_ATTACK)
	}

	if playerItemAppendix.Defence > 0 {
		attrTypeBits |= 1 << uint(item_api.ATTRIBUTE_DEFENCE)
	}

	if playerItemAppendix.DodgeLevel > 0 {
		attrTypeBits |= 1 << uint(item_api.ATTRIBUTE_DODGE_LEVEL)
	}

	if playerItemAppendix.HitLevel > 0 {
		attrTypeBits |= 1 << uint(item_api.ATTRIBUTE_HIT_LEVEL)
	}

	if playerItemAppendix.BlockLevel > 0 {
		attrTypeBits |= 1 << uint(item_api.ATTRIBUTE_BLOCK_LEVEL)
	}

	if playerItemAppendix.CriticalLevel > 0 {
		attrTypeBits |= 1 << uint(item_api.ATTRIBUTE_CRITICAL_LEVEL)
	}

	if playerItemAppendix.TenacityLevel > 0 {
		attrTypeBits |= 1 << uint(item_api.ATTRIBUTE_TENACITY_LEVEL)
	}

	if playerItemAppendix.DestroyLevel > 0 {
		attrTypeBits |= 1 << uint(item_api.ATTRIBUTE_DESTROY_LEVEL)
	}
	return attrTypeBits
}
