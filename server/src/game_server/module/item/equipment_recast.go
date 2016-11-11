package item

import (
	"core/fail"
	"core/net"
	"game_server/api/protocol/item_api"
	"game_server/dat/item_dat"
	"game_server/dat/player_dat"
	"game_server/dat/quest_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/tlog"
	"game_server/xdlog"
	"math"
	"math/rand"
)

const allAttrTypeBits = 1<<uint(item_api.ATTRIBUTE_HEALTH) |
	1<<uint(item_api.ATTRIBUTE_CULTIVATION) |
	1<<uint(item_api.ATTRIBUTE_SPEED) |
	1<<uint(item_api.ATTRIBUTE_ATTACK) |
	1<<uint(item_api.ATTRIBUTE_DEFENCE) |
	1<<uint(item_api.ATTRIBUTE_DODGE_LEVEL) |
	1<<uint(item_api.ATTRIBUTE_HIT_LEVEL) |
	1<<uint(item_api.ATTRIBUTE_BLOCK_LEVEL) |
	1<<uint(item_api.ATTRIBUTE_CRITICAL_LEVEL) |
	1<<uint(item_api.ATTRIBUTE_TENACITY_LEVEL) |
	1<<uint(item_api.ATTRIBUTE_DESTROY_LEVEL)

func GetRecastInfo(session *net.Session, in *item_api.GetRecastInfo_In, out *item_api.GetRecastInfo_Out) {
	state := module.State(session)
	db := state.Database
	tlog.PlayerSystemModuleFlowLog(db, tlog.SMT_EQUIP_RECAST)

	fail.When(db.Lookup.PlayerItemRecastState(state.PlayerId) != nil, "RecastState is exist")

	playerItemId := in.Id
	playerItem := db.Lookup.PlayerItem(playerItemId)
	playerItemAppendix := db.Lookup.PlayerItemAppendix(playerItem.AppendixId)
	item := item_dat.GetItem(playerItem.ItemId)

	attrType := in.Attr

	if playerItemAppendix.RecastAttr != int8(item_api.ATTRIBUTE_NULL) {
		// 如果已重铸过，属性必须是重铸过的属性
		fail.When(playerItemAppendix.RecastAttr != int8(attrType), "recast attr type no match")
	}

	attrs := []item_api.GetRecastInfo_Out_Attrs{}

	attrTypeBit := 1 << uint(attrType)

	// 获取当前装备的附加属性类型
	curAttrTypeBits := getAttrTypeBits(playerItemAppendix)

	sampleAttrTypeBits := allAttrTypeBits - (curAttrTypeBits - attrTypeBit)

	randomAttrCount := item_dat.EQUIPMENT_RECAST_RANDOM_ATTR_NUM
	for {

		randomAttrTypeBit := 1 << uint(rand.Intn(int(item_api.ATTRIBUTE_NUM))+1) // 类型从 1 开始

		if sampleAttrTypeBits&randomAttrTypeBit == randomAttrTypeBit {
			// 随机命中

			randomAttrType := int(math.Log2(float64(randomAttrTypeBit)))

			// 武器，不出现防御属性
			if item.TypeId == item_dat.TYPE_WEAPON && randomAttrType == int(item_api.ATTRIBUTE_DEFENCE) {
				continue
			}
			// 饰品，不出现防御属性
			if item.TypeId == item_dat.TYPE_ACCESSORIES && randomAttrType == int(item_api.ATTRIBUTE_DEFENCE) {
				continue
			}

			attr := getRandomAttr(randomAttrType, item.AppendixLevel)

			attrs = append(attrs, item_api.GetRecastInfo_Out_Attrs{
				Attr:  item_api.Attribute(randomAttrType),
				Value: attr,
			})

			// 被选属性中去掉此属性
			sampleAttrTypeBits -= randomAttrTypeBit

			randomAttrCount--
		}

		if randomAttrCount == 0 {
			break
		}
	}

	// 消耗
	equipmentRecast := item_dat.GetEquipmentRecast(item.Quality, item.Level)

	// 消耗部位碎片
	fragmentNum := equipmentRecast.FragmentNum
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
	}
	module.Item.DelItemByItemId(db, fragmentId, fragmentNum, tlog.IFR_EQUIPMENT_RECAST, xdlog.ET_EQUIPMENT_RECAST)

	// 消耗结晶
	// 蓝色
	if equipmentRecast.BlueCrystalNum > 0 {
		module.Item.DelItemByItemId(db, item_dat.ITEM_BLUE_CRYSTAL, equipmentRecast.BlueCrystalNum, tlog.IFR_EQUIPMENT_RECAST, xdlog.ET_EQUIPMENT_RECAST)
	}
	// 紫色
	if equipmentRecast.PurpleCrystalNum > 0 {
		module.Item.DelItemByItemId(db, item_dat.ITEM_PURPLE_CRYSTAL, equipmentRecast.PurpleCrystalNum, tlog.IFR_EQUIPMENT_RECAST, xdlog.ET_EQUIPMENT_RECAST)
	}
	// 金色
	if equipmentRecast.GoldenCrystalNum > 0 {
		module.Item.DelItemByItemId(db, item_dat.ITEM_GOLDEN_CRYSTAL, equipmentRecast.GoldenCrystalNum, tlog.IFR_EQUIPMENT_RECAST, xdlog.ET_EQUIPMENT_RECAST)
	}
	// 橙色
	if equipmentRecast.OrangeCrystalNum > 0 {
		module.Item.DelItemByItemId(db, item_dat.ITEM_ORANGE_CRYSTAL, equipmentRecast.OrangeCrystalNum, tlog.IFR_EQUIPMENT_RECAST, xdlog.ET_EQUIPMENT_RECAST)
	}

	// 消耗铜钱
	module.Player.DecMoney(db, state.MoneyState, equipmentRecast.ConsumeCoin, player_dat.COINS, tlog.MFR_EQUIPMENT_RECAST, xdlog.ET_EQUIPMENT_RECAST)

	// 写入ItemRecastState
	playerItemRecastState := &mdb.PlayerItemRecastState{
		Pid:          state.PlayerId,
		PlayerItemId: playerItemId,
		SelectedAttr: int8(attrType),
		Attr1Type:    int8(attrs[0].Attr),
		Attr1Value:   attrs[0].Value,
		Attr2Type:    int8(attrs[1].Attr),
		Attr2Value:   attrs[1].Value,
		Attr3Type:    int8(attrs[2].Attr),
		Attr3Value:   attrs[2].Value,
	}
	db.Insert.PlayerItemRecastState(playerItemRecastState)

	//消耗玩家资源就更新每日任务
	module.Quest.RefreshDailyQuest(state, quest_dat.DAILY_QUEST_CLASS_RECAST_EQ)
	// 返回
	out.Attrs = attrs
}

func getRandomAttr(attrType int, itemAppendixLevel int32) (attr int32) {
	lower, upper := item_dat.GetEquipmentRecastAppendix(itemAppendixLevel)

	var upperAttr int32
	var lowerAttr int32
	var randomUnit int32
	switch item_api.Attribute(attrType) {
	case item_api.ATTRIBUTE_ATTACK:
		upperAttr = upper.Attack
		lowerAttr = lower.Attack
		randomUnit = item_dat.ATTRIBUTE_RANDOM_UNIT_ATTACK
	case item_api.ATTRIBUTE_DEFENCE:
		upperAttr = upper.Defence
		lowerAttr = lower.Defence
		randomUnit = item_dat.ATTRIBUTE_RANDOM_UNIT_DEFENCE
	case item_api.ATTRIBUTE_HEALTH:
		upperAttr = upper.Health
		lowerAttr = lower.Health
		randomUnit = item_dat.ATTRIBUTE_RANDOM_UNIT_HEALTH
	case item_api.ATTRIBUTE_SPEED:
		upperAttr = upper.Speed
		lowerAttr = lower.Speed
		randomUnit = item_dat.ATTRIBUTE_RANDOM_UNIT_SPEED
	case item_api.ATTRIBUTE_CULTIVATION:
		upperAttr = upper.Cultivation
		lowerAttr = lower.Cultivation
		randomUnit = item_dat.ATTRIBUTE_RANDOM_UNIT_CULTIVATION
	case item_api.ATTRIBUTE_HIT_LEVEL:
		upperAttr = upper.HitLevel
		lowerAttr = lower.HitLevel
		randomUnit = item_dat.ATTRIBUTE_RANDOM_UNIT_HIT_LEVEL
	case item_api.ATTRIBUTE_CRITICAL_LEVEL:
		upperAttr = upper.CriticalLevel
		lowerAttr = lower.CriticalLevel
		randomUnit = item_dat.ATTRIBUTE_RANDOM_UNIT_CRITICAL_LEVEL
	case item_api.ATTRIBUTE_BLOCK_LEVEL:
		upperAttr = upper.BlockLevel
		lowerAttr = lower.BlockLevel
		randomUnit = item_dat.ATTRIBUTE_RANDOM_UNIT_BLOCK_LEVEL
	case item_api.ATTRIBUTE_DESTROY_LEVEL:
		upperAttr = upper.DestroyLevel
		lowerAttr = lower.DestroyLevel
		randomUnit = item_dat.ATTRIBUTE_RANDOM_UNIT_DESTROY_LEVEL
	case item_api.ATTRIBUTE_TENACITY_LEVEL:
		upperAttr = upper.TenacityLevel
		lowerAttr = lower.TenacityLevel
		randomUnit = item_dat.ATTRIBUTE_RANDOM_UNIT_TENACITY_LEVEL
	case item_api.ATTRIBUTE_DODGE_LEVEL:
		upperAttr = upper.DodgeLevel
		lowerAttr = lower.DodgeLevel
		randomUnit = item_dat.ATTRIBUTE_RANDOM_UNIT_DODGE_LEVEL
	}

	attr = randomAttr(upperAttr, lowerAttr, randomUnit)
	return attr
}

func randomAttr(upperAttr int32, lowerAttr int32, randomUnit int32) (attr int32) {
	interal := (upperAttr - lowerAttr) / randomUnit
	attr = lowerAttr + int32(rand.Intn(int(interal)+1))*randomUnit
	return attr
}

func Recast(session *net.Session, in *item_api.Recast_In) {
	state := module.State(session)
	db := state.Database
	tlog.PlayerSystemModuleFlowLog(db, tlog.SMT_EQUIP_RECAST)
	playerItemRecastState := db.Lookup.PlayerItemRecastState(state.PlayerId)
	fail.When(playerItemRecastState == nil, "no player item recast state")

	playerItemId := playerItemRecastState.PlayerItemId
	playerItem := db.Lookup.PlayerItem(playerItemId)
	playerItemAppendix := db.Lookup.PlayerItemAppendix(playerItem.AppendixId)
	attrType := int8(in.Attr)

	if attrType == int8(item_api.ATTRIBUTE_NULL) {
		// 当选中的重铸属性为 NULL 时，即保留以前属性，但需要标记此属性为重铸属性
		playerItemAppendix.RecastAttr = playerItemRecastState.SelectedAttr
		db.Update.PlayerItemAppendix(playerItemAppendix)

		// 销毁重铸状态
		db.Delete.PlayerItemRecastState(playerItemRecastState)
		return
	}

	var attrValue int32
	switch attrType {
	case playerItemRecastState.Attr1Type:
		attrValue = playerItemRecastState.Attr1Value
	case playerItemRecastState.Attr2Type:
		attrValue = playerItemRecastState.Attr2Value
	case playerItemRecastState.Attr3Type:
		attrValue = playerItemRecastState.Attr3Value
	default:
		fail.When(true, "no this attr type in item recast state")
	}

	// 去掉原有属性
	setAppendixAttr(playerItemAppendix, playerItemRecastState.SelectedAttr, 0)
	// 添加新属性
	setAppendixAttr(playerItemAppendix, attrType, attrValue)
	// 设置重铸属性
	playerItemAppendix.RecastAttr = attrType

	// 更新装备附加属性数据
	db.Update.PlayerItemAppendix(playerItemAppendix)

	// 销毁重铸状态
	db.Delete.PlayerItemRecastState(playerItemRecastState)

	return
}

func setAppendixAttr(playerItemAppendix *mdb.PlayerItemAppendix, attrType int8, attrValue int32) {
	switch item_api.Attribute(attrType) {
	case item_api.ATTRIBUTE_ATTACK:
		playerItemAppendix.Attack = attrValue
	case item_api.ATTRIBUTE_DEFENCE:
		playerItemAppendix.Defence = attrValue
	case item_api.ATTRIBUTE_HEALTH:
		playerItemAppendix.Health = attrValue
	case item_api.ATTRIBUTE_SPEED:
		playerItemAppendix.Speed = attrValue
	case item_api.ATTRIBUTE_CULTIVATION:
		playerItemAppendix.Cultivation = attrValue
	case item_api.ATTRIBUTE_HIT_LEVEL:
		playerItemAppendix.HitLevel = attrValue
	case item_api.ATTRIBUTE_CRITICAL_LEVEL:
		playerItemAppendix.CriticalLevel = attrValue
	case item_api.ATTRIBUTE_BLOCK_LEVEL:
		playerItemAppendix.BlockLevel = attrValue
	case item_api.ATTRIBUTE_DESTROY_LEVEL:
		playerItemAppendix.DestroyLevel = attrValue
	case item_api.ATTRIBUTE_TENACITY_LEVEL:
		playerItemAppendix.TenacityLevel = attrValue
	case item_api.ATTRIBUTE_DODGE_LEVEL:
		playerItemAppendix.DodgeLevel = attrValue
	}
}
