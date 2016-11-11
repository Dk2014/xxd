package item

import (
	"core/debug"
	"core/fail"
	"core/i18l"
	"core/log"
	"core/net"
	"core/time"
	"fmt"
	"game_server/api/protocol/item_api"
	"game_server/api/protocol/notify_api"
	"game_server/dat/item_dat"
	"game_server/dat/mail_dat"
	"game_server/dat/player_dat"
	"game_server/dat/role_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/module/rpc"
	"game_server/tlog"
	"game_server/xdlog"
)

func GetAllItems(session *net.Session, out *item_api.GetAllItems_Out) {
	state := module.State(session)
	db := state.Database

	out.Items = make([]item_api.GetAllItems_Out_Items, 0, item_dat.MAX_ITEM_NUM)
	out.Equipments = make([]item_api.GetAllItems_Out_Equipments, 0, role_dat.MAX_ROLE_NUM)
	out.Buybacks = make([]item_api.GetAllItems_Out_Buybacks, 0, item_dat.BUY_BACK_NUM)
	equipments := make(map[int64]*mdb.PlayerItem)
	buybacks := make(map[int64]*mdb.PlayerItem)

	//获取玩家背包，不包括被装备的和回购栏的
	db.Select.PlayerItem(func(row *mdb.PlayerItemRow) {
		if row.IsDressed() == 0 && row.BuyBackState() == 0 {

			playerItemAppendix := tryGetPlayerItemAppendix(db, row.AppendixId())

			out.Items = append(out.Items,
				item_api.GetAllItems_Out_Items{
					Id:     row.Id(),
					ItemId: row.ItemId(),
					Num:    row.Num(),

					Attack:        playerItemAppendix.Attack,
					Defence:       playerItemAppendix.Defence,
					Health:        playerItemAppendix.Health,
					Speed:         playerItemAppendix.Speed,
					Cultivation:   playerItemAppendix.Cultivation,
					HitLevel:      playerItemAppendix.HitLevel,
					CriticalLevel: playerItemAppendix.CriticalLevel,
					BlockLevel:    playerItemAppendix.BlockLevel,
					DestroyLevel:  playerItemAppendix.DestroyLevel,
					TenacityLevel: playerItemAppendix.TenacityLevel,
					DodgeLevel:    playerItemAppendix.DodgeLevel,

					RefineLevel:     row.RefineLevel(),
					RefineFailTimes: row.RefineFailTimes(),
					RecastAttr:      item_api.Attribute(playerItemAppendix.RecastAttr),
				})
		} else if row.IsDressed() == 1 {
			equipments[row.Id()] = row.GoObject()
		} else if row.BuyBackState() == 1 {
			buybacks[row.Id()] = row.GoObject()
		}
	})

	//获取玩家所有角色装备
	if len(equipments) > 0 {
		db.Select.PlayerEquipment(func(row *mdb.PlayerEquipmentRow) {
			equips := make([]item_api.GetAllItems_Out_Equipments_Equips, 0, item_dat.EQUIPMENT_NUM_EVERYBODY)
			for _, eqId := range []int64{row.WeaponId(), row.ClothesId(), row.AccessoriesId(), row.ShoeId()} {
				if equipment, ok := equipments[eqId]; ok {

					playerItemAppendix := tryGetPlayerItemAppendix(db, equipment.AppendixId)

					equips = append(equips, item_api.GetAllItems_Out_Equipments_Equips{
						Id:              eqId,
						ItemId:          equipment.ItemId,
						Attack:          playerItemAppendix.Attack,
						Defence:         playerItemAppendix.Defence,
						Health:          playerItemAppendix.Health,
						Speed:           playerItemAppendix.Speed,
						Cultivation:     playerItemAppendix.Cultivation,
						HitLevel:        playerItemAppendix.HitLevel,
						CriticalLevel:   playerItemAppendix.CriticalLevel,
						BlockLevel:      playerItemAppendix.BlockLevel,
						DestroyLevel:    playerItemAppendix.DestroyLevel,
						TenacityLevel:   playerItemAppendix.TenacityLevel,
						DodgeLevel:      playerItemAppendix.DodgeLevel,
						RefineLevel:     equipment.RefineLevel,
						RefineFailTimes: equipment.RefineFailTimes,
						RecastAttr:      item_api.Attribute(playerItemAppendix.RecastAttr),
					})
				}
			}

			out.Equipments = append(out.Equipments,
				item_api.GetAllItems_Out_Equipments{
					RoleId: row.RoleId(),
					Equips: equips,
				})
		})
	}

	//获取玩家回购栏物品
	if len(buybacks) > 0 {
		playerItemBuyback := db.Lookup.PlayerItemBuyback(state.PlayerId)

		for _, backId := range []int64{playerItemBuyback.BackId1, playerItemBuyback.BackId2, playerItemBuyback.BackId3, playerItemBuyback.BackId4, playerItemBuyback.BackId5, playerItemBuyback.BackId6} {
			if buyback, ok := buybacks[backId]; ok {
				buybackItem := item_api.GetAllItems_Out_Buybacks{
					Id:     buyback.Id,
					ItemId: buyback.ItemId,
					Num:    buyback.Num,
				}
				playerItemAppendix := state.Database.Lookup.PlayerItemAppendix(buyback.AppendixId)
				if playerItemAppendix != nil {
					buybackItem.RecastAttr = item_api.Attribute(playerItemAppendix.RecastAttr)
				}
				buybackItem.RefineLevel = buyback.RefineLevel
				out.Buybacks = append(out.Buybacks, buybackItem)

			}
		}
	}

	db.Select.PlayerPurchaseRecord(func(row *mdb.PlayerPurchaseRecordRow) {
		record := item_api.GetAllItems_Out_BuyRecords{
			ItemId: row.ItemId(),
			Num:    row.Num(),
		}
		if !time.IsToday(row.Timestamp()) {
			record.Num = 0
		}
		out.BuyRecords = append(out.BuyRecords, record)

	})
}

func BuyItem(session *net.Session, in *item_api.BuyItem_In) (id int64) {
	state := module.State(session)
	//TODO 物品兑换和物品购买略有不同，待和策划确认。
	fail.When(!canAddItem(state.Database, in.ItemId, 1), "背包已满不能购买")

	//先查找兑换
	var itemId int16
	var itemNum int16
	res := item_dat.FetchItemExchange(in.ItemId, func(itemExchange *item_dat.ItemExchange) {
		itemId = itemExchange.ItemId
		itemNum = itemExchange.ItemNum
	})

	if res {
		delItemByItemId(state.Database, itemId, itemNum, tlog.IFR_BUY_ITEM, xdlog.ET_BUY_ITEM)
	} else {
		buyLimit := item_dat.PurchaseLimitation(in.ItemId)
		buyRecord := updateBuyRecord(state, in.ItemId)
		buyRecord.Num++
		fail.When(buyRecord.Num > buyLimit, "购买次数到达上限")
		state.Database.Update.PlayerPurchaseRecord(buyRecord)
		item := item_dat.GetItem(in.ItemId)
		module.Player.DecMoney(state.Database, state.MoneyState, item.Price, player_dat.COINS, tlog.MFR_BUY_ITEM, xdlog.ET_BUY_ITEM)
		itemMoneyFlow(state.Database, item, 1, int32(item.Price), tlog.MT_COIN)
		tlog.PlayerBusinessManFlowLog(state.Database, int32(in.ItemId), 1, int32(item.Price), tlog.ADD)
	}

	id = addItem(state.Database, in.ItemId, 1, tlog.IFR_BUY_ITEM, xdlog.ET_BUY_ITEM, "")

	return
}

func SellItem(session *net.Session, in *item_api.SellItem_In) {
	state := module.State(session)

	playerItem := state.Database.Lookup.PlayerItem(in.Id)
	fail.When(playerItem == nil, "SellItem wrong id")
	fail.When(playerItem.BuyBackState == item_dat.ITEM_STATE_SOLD, "物品已出售")

	item := item_dat.GetItem(playerItem.ItemId)
	fail.When(!item.CanSell, "物品禁止出售")

	var chance float64 = item_dat.SELL_DISCOUNT
	if item.TypeId == item_dat.TYPE_WEAPON || item.TypeId == item_dat.TYPE_CLOTHES || item.TypeId == item_dat.TYPE_SHOE || item.TypeId == item_dat.TYPE_ACCESSORIES {
		chance = item_dat.SELL_EQUIP_DISCOUNT
	}
	// 物品的销售价格为原价加上精炼总价加上一次重铸价格的30%
	total_price := item.Price

	// 是否精炼
	if playerItem.RefineLevel > 0 {
		equipmentRefine := item_dat.GetEquipmentRefineDat(int8(item.TypeId), int8(item.EquipTypeId))
		for i := 1; i <= int(playerItem.RefineLevel); i++ {
			total_price += int64(equipmentRefine.BasePrice + int32(i)*equipmentRefine.IncrePrice)
		}
	}
	total_price = int64(float64(total_price) * chance)

	// 是否重铸
	if playerItemAppendix := state.Database.Lookup.PlayerItemAppendix(playerItem.AppendixId); playerItemAppendix != nil && playerItemAppendix.RecastAttr != int8(item_api.ATTRIBUTE_NULL) {
		total_price += int64(float64(item_dat.GetEquipmentRecast(item.Quality, item.Level).ConsumeCoin) * float64(item_dat.SELL_EQUIP_RECAST_DISCOUNT))
	}

	module.Player.IncMoney(state.Database, state.MoneyState, total_price*int64(playerItem.Num), player_dat.COINS, tlog.MFR_SELL_ITEM, xdlog.ET_SELL_ITEM, "")

	//如果装备被穿戴，先脱下
	if playerItem.IsDressed == 1 {
		undress(state.Database, playerItem)
	}

	//加入回购栏
	playerItem.IsDressed = 0
	playerItem.BuyBackState = 1
	state.Database.Update.PlayerItem(playerItem)
	tlog.PlayerItemFlowLog(state.Database, item.Id, item.TypeId, playerItem.Num, tlog.REDUCE, tlog.IFR_SELL_ITEM)
	xdlog.PropsLog(state.Database, item.Id, playerItem.Num, xdlog.ET_SELL_ITEM)

	playerItemBuyback := state.Database.Lookup.PlayerItemBuyback(state.PlayerId)
	fail.When(playerItemBuyback == nil, "SellItem wrong pid")

	ids := []*int64{&playerItemBuyback.BackId1, &playerItemBuyback.BackId2, &playerItemBuyback.BackId3,
		&playerItemBuyback.BackId4, &playerItemBuyback.BackId5, &playerItemBuyback.BackId6}

	//回购栏有空位置则放入
	added := false
	for i := 0; i < 6; i++ {
		if *ids[i] == 0 {
			*ids[i] = in.Id
			added = true
			break
		}
	}

	//没空位置则删除最早的,然后放入,保持回购栏按时间有序
	var delId int64 = 0
	if !added {
		delId = *ids[0]
		*ids[0] = 0
		for i := 0; i < 5; i++ {
			*ids[i], *ids[i+1] = *ids[i+1], *ids[i]
		}
		*ids[5] = in.Id
	}

	if delId != 0 {
		delItem := state.Database.Lookup.PlayerItem(delId)
		//fail.When(delItem == nil, "SellItem del wrong id")
		if delItem != nil {
			delItemFromDB(state.Database, delItem, tlog.IFR_SELL_ITEM, xdlog.ET_SELL_ITEM)
		} else {
			//因为之前有bug导致用户重复出售物品这里纪录一下
			//TODO 修复数据并且恢复严格校验
			log.Debugf("物品购状态不一致\n %d %#v \n %s\n", state.PlayerId, playerItemBuyback, debug.Stack(1, "    "))
		}

	}

	state.Database.Update.PlayerItemBuyback(playerItemBuyback)
	tlog.PlayerBusinessManFlowLog(state.Database, int32(item.Id), int32(playerItem.Num), int32(float64(item.Price*int64(playerItem.Num))*item_dat.SELL_DISCOUNT), tlog.REDUCE)
}

func Dress(session *net.Session, in *item_api.Dress_In) {
	state := module.State(session)

	playerItem := state.Database.Lookup.PlayerItem(in.Id)
	fail.When(playerItem == nil, "Dress wrong id")
	fail.When(playerItem.IsDressed == 1, "装备已经被穿戴")

	item := item_dat.GetItem(playerItem.ItemId)
	fail.When(item.TypeId != item_dat.TYPE_WEAPON && item.TypeId != item_dat.TYPE_CLOTHES && item.TypeId != item_dat.TYPE_ACCESSORIES && item.TypeId != item_dat.TYPE_SHOE, "wrong item type")

	// 检查装备角色
	if item.EquipRoleId != 0 {
		equipRoleId := in.RoleId

		if role_dat.IsMainRole(equipRoleId) {
			equipRoleId = item_dat.EQUIPMENT_MAIN_ROLE_ID
		}

		fail.When(equipRoleId != item.EquipRoleId, "Dress wrong equipRoleId")
	}

	// 当前角色等级是否可穿戴
	playerRole := module.Role.GetBuddyRoleInTeam(state.Database, in.RoleId)
	fail.When(int32(playerRole.Level) < item.Level, "Dress wrong level")

	var playerEquipment *mdb.PlayerEquipment
	state.Database.Select.PlayerEquipment(func(row *mdb.PlayerEquipmentRow) {
		if row.RoleId() == in.RoleId {
			playerEquipment = row.GoObject()
			row.Break()
		}
	})
	fail.When(playerEquipment == nil, "角色装备数据未初始化")

	var unDressEquipId int64

	switch item.TypeId {
	case item_dat.TYPE_WEAPON:
		unDressEquipId = playerEquipment.WeaponId
		playerEquipment.WeaponId = playerItem.Id

	case item_dat.TYPE_CLOTHES:
		unDressEquipId = playerEquipment.ClothesId
		playerEquipment.ClothesId = playerItem.Id

	case item_dat.TYPE_ACCESSORIES:
		unDressEquipId = playerEquipment.AccessoriesId
		playerEquipment.AccessoriesId = playerItem.Id

	case item_dat.TYPE_SHOE:
		unDressEquipId = playerEquipment.ShoeId
		playerEquipment.ShoeId = playerItem.Id

	default:
		fail.When(true, "无效的装备类型")
	}

	state.Database.Update.PlayerEquipment(playerEquipment)

	playerItem.IsDressed = 1
	state.Database.Update.PlayerItem(playerItem)

	//原位置有装备的话，将原装备脱下
	if unDressEquipId > 0 {
		if playerItem2 := state.Database.Lookup.PlayerItem(unDressEquipId); playerItem2 != nil {
			playerItem2.IsDressed = 0
			state.Database.Update.PlayerItem(playerItem2)
		}
	}
}

func Undress(session *net.Session, in *item_api.Undress_In) {
	fail.When(in.Pos < 0 || in.Pos > 3, "wrong pos")
	state := module.State(session)
	fail.When(isBagFull(state.Database), "Undress bag full")

	var playerEquipment *mdb.PlayerEquipment
	state.Database.Select.PlayerEquipment(func(row *mdb.PlayerEquipmentRow) {
		if row.RoleId() == in.RoleId {
			playerEquipment = row.GoObject()
			row.Break()
		}
	})
	fail.When(playerEquipment == nil, "不存在已穿戴的装备")

	var unDressEquipId int64
	switch in.Pos {
	case item_api.EQUIPMENT_POS_WEAPON:
		unDressEquipId = playerEquipment.WeaponId
		playerEquipment.WeaponId = 0

	case item_api.EQUIPMENT_POS_CLOTHES:
		unDressEquipId = playerEquipment.ClothesId
		playerEquipment.ClothesId = 0

	case item_api.EQUIPMENT_POS_ACCESSORIES:
		unDressEquipId = playerEquipment.AccessoriesId
		playerEquipment.AccessoriesId = 0

	case item_api.EQUIPMENT_POS_SHOE:
		unDressEquipId = playerEquipment.ShoeId
		playerEquipment.ShoeId = 0

	default:
		fail.When(true, "无效的装备位置")
	}

	state.Database.Update.PlayerEquipment(playerEquipment)
	if playerItem := state.Database.Lookup.PlayerItem(unDressEquipId); playerItem != nil {
		playerItem.IsDressed = 0
		state.Database.Update.PlayerItem(playerItem)
	}
}

func BuyItemBack(session *net.Session, in *item_api.BuyItemBack_In, out *item_api.BuyItemBack_Out) {
	state := module.State(session)

	playerItem := state.Database.Lookup.PlayerItem(in.Id)
	fail.When(playerItem == nil, "BuyItemBack playerItem nil")

	//检查背包是否有空间装回购的物品
	fail.When(!canAddItem(state.Database, playerItem.ItemId, playerItem.Num), "BuyItemBack bag full")

	playerItemBuyback := state.Database.Lookup.PlayerItemBuyback(state.PlayerId)
	fail.When(playerItemBuyback == nil, "BuyItemBack wrong pid")

	ids := []*int64{&playerItemBuyback.BackId1, &playerItemBuyback.BackId2, &playerItemBuyback.BackId3,
		&playerItemBuyback.BackId4, &playerItemBuyback.BackId5, &playerItemBuyback.BackId6}
	index := -1
	for i := 0; i < 6; i++ {
		if *ids[i] == in.Id {
			*ids[i] = 0
			index = i
		}
	}
	fail.When(index == -1, "BuyItemBack wrong id")

	//保持回购栏位置按时间有序
	for i := index; i < 5; i++ {
		*ids[i], *ids[i+1] = *ids[i+1], *ids[i]
	}
	state.Database.Update.PlayerItemBuyback(playerItemBuyback)

	// 消耗
	item := item_dat.GetItem(playerItem.ItemId)
	var chance float64 = item_dat.SELL_DISCOUNT
	if item.TypeId == item_dat.TYPE_WEAPON || item.TypeId == item_dat.TYPE_CLOTHES || item.TypeId == item_dat.TYPE_SHOE || item.TypeId == item_dat.TYPE_ACCESSORIES {
		chance = item_dat.SELL_EQUIP_DISCOUNT
	}
	// 物品的销售价格为原价加上精炼总价加上一次重铸价格的30%
	total_price := item.Price

	// 是否精炼
	if playerItem.RefineLevel > 0 {
		equipmentRefine := item_dat.GetEquipmentRefineDat(int8(item.TypeId), int8(item.EquipTypeId))
		for i := 1; i <= int(playerItem.RefineLevel); i++ {
			total_price += int64(equipmentRefine.BasePrice + int32(i)*equipmentRefine.IncrePrice)
		}
	}
	total_price = int64(float64(total_price) * chance)

	// 是否重铸
	if playerItemAppendix := state.Database.Lookup.PlayerItemAppendix(playerItem.AppendixId); playerItemAppendix != nil && playerItemAppendix.RecastAttr != int8(item_api.ATTRIBUTE_NULL) {
		total_price += int64(float64(item_dat.GetEquipmentRecast(item.Quality, item.Level).ConsumeCoin) * float64(item_dat.SELL_EQUIP_RECAST_DISCOUNT))
	}
	module.Player.DecMoney(state.Database, state.MoneyState, total_price*int64(playerItem.Num), player_dat.COINS, tlog.MFR_BUY_ITEM_BACK, xdlog.ET_BUY_BACK_ITEM)
	itemMoneyFlow(state.Database, item, int32(playerItem.Num), int32(total_price*int64(playerItem.Num)), tlog.MT_COIN)

	//涉及到的物品ID
	var newIds []int64
	itemType := item_dat.GetItemType(item.TypeId)
	if itemType.MaxNumInPos == 1 {
		//堆叠次数为1的物品直接把买回状态设置为0
		playerItem.BuyBackState = 0
		state.Database.Update.PlayerItem(playerItem)
		tlog.PlayerItemFlowLog(state.Database, item.Id, item.TypeId, 1, tlog.ADD, tlog.IFR_BUY_ITEM_BACK)
		xdlog.ItemLog(state.Database, item.Id, 1, xdlog.ET_BUY_BACK_ITEM, "")
		newIds = append(newIds, playerItem.Id)
	} else {
		//删除原先的playerItem 然后增加物品
		num := playerItem.Num
		delItemFromDB(state.Database, playerItem, tlog.IFR_BUY_ITEM_BACK, xdlog.ET_BUY_BACK_ITEM)
		useLess := &notify_api.ItemChange_Out{}
		useLess.Items = make([]notify_api.ItemChange_Out_Items, 0, item_dat.MAX_BAG_NUM)
		newIds = addItemAndSetResp(state.Database, item.Id, num, useLess, tlog.IFR_BUY_ITEM_BACK, xdlog.ET_BUY_BACK_ITEM, "")
	}
	setBuyBackItemChange(session, newIds, out)
	tlog.PlayerBusinessManFlowLog(state.Database, int32(item.Id), int32(playerItem.Num), int32(float64(item.Price*int64(playerItem.Num))*item_dat.SELL_DISCOUNT), tlog.ADD)
}

//返回最后个ID
func addItem(db *mdb.Database, itemId int16, num int16, itemFlowReason, xdEventType int32, xdEventParam string) int64 {

	out := &notify_api.ItemChange_Out{}
	out.Items = make([]notify_api.ItemChange_Out_Items, 0, item_dat.MAX_BAG_NUM)
	ids := addItemAndSetResp(db, itemId, num, out, itemFlowReason, xdEventType, xdEventParam)
	//通知玩家物品改变
	if session, ok := module.Player.GetPlayerOnline(db.PlayerId()); ok {
		module.Notify.SendItemChange(session, out)
	}
	if len(ids) > 0 {
		return ids[len(ids)-1]
	}
	return 0
}

func batchAddItem(db *mdb.Database, items map[int16]int16, itemFlowReason, xdEventType int32) {
	fail.When(len(items) == 0, "batchAddItem: add nothing")
	if isBagFull(db) {
		attach := make([]*mail_dat.Attachment, 0, len(items))
		for itemId, num := range items {
			attach = append(attach, &mail_dat.Attachment{mail_dat.ATTACHMENT_ITEM, itemId, int64(num)})
		}
		rpc.RemoteMailSend(db.PlayerId(), mail_dat.MailBagFull{i18l.T.Tran("获得"), attach})
		return
	}

	out := &notify_api.ItemChange_Out{}
	out.Items = make([]notify_api.ItemChange_Out_Items, 0, item_dat.MAX_BAG_NUM)

	for itemId, num := range items {
		addItemAndSetResp(db, itemId, num, out, itemFlowReason, xdEventType, "")
	}
	if session, ok := module.Player.GetPlayerOnline(db.PlayerId()); ok {
		module.Notify.SendItemChange(session, out)
	}
}

//返回最后一条插入的id
func addItemAndSetResp(db *mdb.Database, itemId int16, num int16, items *notify_api.ItemChange_Out, itemFlowReason, xdEventType int32, xdEventParam string) []int64 {
	ids := make([]int64, 0, 1)

	fail.When(num > 30000 || num <= 0, "wrong number")
	item := item_dat.GetItem(itemId)
	fail.When(item.TypeId == item_dat.TYPE_RESOURCE, "资源类物品不能添加")

	if !canAddItem(db, itemId, num) {
		rpc.RemoteMailSend(db.PlayerId(), mail_dat.MailBagFull{i18l.T.Tran(item.Name), []*mail_dat.Attachment{&mail_dat.Attachment{0, itemId, int64(num)}}})
		return ids
	}

	updateSealedBookRecord(db, item)

	itemType := item_dat.GetItemType(item.TypeId)

	var data *mdb.PlayerItem = nil
	var minNum int16 = itemType.MaxNumInPos
	var bagSize int16 = 0
	var heapNum int16 = num / itemType.MaxNumInPos
	var rest int16 = num % itemType.MaxNumInPos

	//寻找对应物品数量最少的格子
	db.Select.PlayerItem(func(row *mdb.PlayerItemRow) {
		if row.IsDressed() == 0 && row.BuyBackState() == 0 {
			bagSize++
		}
		if row.ItemId() == itemId && row.BuyBackState() == 0 {
			if row.Num() < minNum {
				data = row.GoObject()
				minNum = row.Num()
			}
		}
	})
	if heapNum > item_dat.MAX_BAG_NUM-bagSize {
		rpc.RemoteMailSend(db.PlayerId(), mail_dat.MailBagFull{i18l.T.Tran(item.Name), []*mail_dat.Attachment{&mail_dat.Attachment{0, itemId, int64(itemType.MaxNumInPos * (heapNum - item_dat.MAX_BAG_NUM + bagSize))}}})

		heapNum = item_dat.MAX_BAG_NUM - bagSize
	}

	//插入满堆的物品
	for i := 0; i < int(heapNum); i++ {
		insertItem := &mdb.PlayerItem{
			Pid:       db.PlayerId(),
			ItemId:    itemId,
			Num:       itemType.MaxNumInPos,
			IsDressed: 0,
		}
		modifyPlayerItem(db, insertItem, items, insertItem.Num, item_dat.MODIFY_INSERT, tlog.ADD, itemFlowReason, xdEventType, xdEventParam)

		ids = append(ids, insertItem.Id)
	}

	bagSize += heapNum

	//插入剩余数目的物品
	//装备物品格子不能叠加数目,所以不用考虑搜索到装备格子的问题,装备附加属性也不需要考虑
	if rest > 0 {
		if data == nil {
			insertItem := &mdb.PlayerItem{
				Pid:       db.PlayerId(),
				ItemId:    itemId,
				Num:       rest,
				IsDressed: 0,
			}
			modifyPlayerItem(db, insertItem, items, insertItem.Num, item_dat.MODIFY_INSERT, tlog.ADD, itemFlowReason, xdEventType, xdEventParam)
			ids = append(ids, insertItem.Id)
		} else if data.Num >= itemType.MaxNumInPos {
			if bagSize < item_dat.MAX_BAG_NUM {
				insertItem := &mdb.PlayerItem{
					Pid:       db.PlayerId(),
					ItemId:    itemId,
					Num:       rest,
					IsDressed: 0,
				}
				modifyPlayerItem(db, insertItem, items, insertItem.Num, item_dat.MODIFY_INSERT, tlog.ADD, itemFlowReason, xdEventType, xdEventParam)
				ids = append(ids, insertItem.Id)
			} else {
				rpc.RemoteMailSend(db.PlayerId(), mail_dat.MailBagFull{item.Name, []*mail_dat.Attachment{&mail_dat.Attachment{0, itemId, int64(rest)}}})
			}
		} else {
			overNum := rest + data.Num - itemType.MaxNumInPos

			if overNum <= 0 {
				data.Num += rest
				modifyPlayerItem(db, data, items, rest, item_dat.MODIFY_UPDATE, tlog.ADD, itemFlowReason, xdEventType, xdEventParam)
				ids = append(ids, data.Id)
			} else {
				beforeNum := data.Num
				data.Num = itemType.MaxNumInPos
				modifyPlayerItem(db, data, items, itemType.MaxNumInPos-beforeNum, item_dat.MODIFY_UPDATE, tlog.ADD, itemFlowReason, xdEventType, xdEventParam)
				ids = append(ids, data.Id)
				if bagSize < item_dat.MAX_BAG_NUM {
					insertItem := &mdb.PlayerItem{
						Pid:       db.PlayerId(),
						ItemId:    itemId,
						Num:       overNum,
						IsDressed: 0,
					}
					modifyPlayerItem(db, insertItem, items, insertItem.Num, item_dat.MODIFY_INSERT, tlog.ADD, itemFlowReason, xdEventType, xdEventParam)
					ids = append(ids, insertItem.Id, data.Id)
				} else {
					rpc.RemoteMailSend(db.PlayerId(), mail_dat.MailBagFull{item.Name, []*mail_dat.Attachment{&mail_dat.Attachment{0, itemId, int64(overNum)}}})
				}
			}
		}
	}
	return ids
}

//删除某格子的道具
func delItemById(db *mdb.Database, id int64, num int16, itemFlowReason, xdEventType int32) {
	data := db.Lookup.PlayerItem(id)
	fail.When(data == nil, "delItemById wrong id")
	fail.When(num > data.Num, "delItemById wrong num")
	fail.When(data.BuyBackState == 1, "delItemById cant delete")

	item := item_dat.GetItem(data.ItemId)

	//删除已穿戴的装备需要将装备脱下
	if data.IsDressed == 1 {
		undress(db, data)
	}

	data.Num -= num

	if num == 0 || data.Num == 0 {
		delItemFromDB(db, data, itemFlowReason, xdEventType)
	} else {
		db.Update.PlayerItem(data)
		tlog.PlayerItemFlowLog(db, item.Id, item.TypeId, num, tlog.REDUCE, itemFlowReason)
		xdlog.PropsLog(db, item.Id, num, xdEventType)
	}
}

//根据物品ID删除道具
func delItemByItemId(db *mdb.Database, itemId int16, num int16, itemFlowReason, xdlogEventType int32) []int64 {
	item := item_dat.GetItem(itemId)
	itemType := item_dat.GetItemType(item.TypeId)
	var heaps []*mdb.PlayerItem = make([]*mdb.PlayerItem, 0, item_dat.MAX_BAG_NUM)
	var minHeapIndex int = 0
	var maxHeapIndex int = 0 //取出除最少列外的一列（都是最大堆叠，随便取一个就行），用于做满堆叠扣除
	var maxHeapNum int16 = itemType.MaxNumInPos
	var minHeapNum int16 = itemType.MaxNumInPos
	var heapNum int16 = num / itemType.MaxNumInPos
	var rest int16 = num % itemType.MaxNumInPos
	var itemNum int64 = 0
	var changedItemsId []int64

	items := &notify_api.ItemChange_Out{}
	items.Items = make([]notify_api.ItemChange_Out_Items, 0, item_dat.MAX_BAG_NUM)

	db.Select.PlayerItem(func(row *mdb.PlayerItemRow) {
		if row.ItemId() == itemId && row.BuyBackState() == 0 {
			itemNum += int64(row.Num())
			heaps = append(heaps, row.GoObject())
			if maxHeapIndex == 0 && row.Num() == maxHeapNum {
				maxHeapIndex = len(heaps) - 1
			}
			if row.Num() < minHeapNum {
				minHeapIndex = len(heaps) - 1
				minHeapNum = row.Num()
			}
		}
	})

	fail.When(itemNum < int64(num), fmt.Sprintf("delItem wrong num, delete num %d, cur num %d", num, itemNum))

	//删除整堆的物品
	gap := 0
	for i := 0; i < int(heapNum); i++ {
		if i == minHeapIndex && minHeapNum < itemType.MaxNumInPos {
			gap = 1
		}
		changedItemsId = append(changedItemsId, heaps[i+gap].Id)
		modifyPlayerItem(db, heaps[i+gap], items, 0, item_dat.MODIFY_DELETE, tlog.REDUCE, itemFlowReason, xdlogEventType, "")
	}

	//删除不满堆的物品
	if rest > 0 {
		if heaps[minHeapIndex].Num > rest {
			heaps[minHeapIndex].Num -= rest
			changedItemsId = append(changedItemsId, heaps[minHeapIndex].Id)
			modifyPlayerItem(db, heaps[minHeapIndex], items, rest, item_dat.MODIFY_UPDATE, tlog.REDUCE, itemFlowReason, xdlogEventType, "")
		} else {
			//rest就是删除最小堆叠之后，还需向满堆叠中删除的数量
			rest = rest - heaps[minHeapIndex].Num
			changedItemsId = append(changedItemsId, heaps[minHeapIndex].Id)
			modifyPlayerItem(db, heaps[minHeapIndex], items, 0, item_dat.MODIFY_DELETE, tlog.REDUCE, itemFlowReason, xdlogEventType, "")
			if rest > 0 {
				//如果删除后还有多余，从最大堆叠中扣除
				changedItemsId = append(changedItemsId, heaps[maxHeapIndex].Id)
				heaps[maxHeapIndex].Num -= rest
				modifyPlayerItem(db, heaps[maxHeapIndex], items, rest, item_dat.MODIFY_UPDATE, tlog.REDUCE, itemFlowReason, xdlogEventType, "")
			}
		}
	}
	//通知玩家物品改变
	if session, ok := module.Player.GetPlayerOnline(db.PlayerId()); ok {
		module.Notify.SendItemChange(session, items)
	}
	return changedItemsId
}

func checkItemNum(db *mdb.Database, itemId int16, num int16) bool {
	return getItemNum(db, itemId) >= num
}

func getItemNum(db *mdb.Database, itemId int16) int16 {
	var itemNum int16 = 0
	db.Select.PlayerItem(func(row *mdb.PlayerItemRow) {
		if row.ItemId() == itemId && row.BuyBackState() == 0 {
			itemNum += row.Num()
		}
	})

	return itemNum
}

func isBagFull(db *mdb.Database) bool {
	bagSize := 0
	db.Select.PlayerItem(func(row *mdb.PlayerItemRow) {
		if row.IsDressed() == 0 && row.BuyBackState() == 0 {
			bagSize++
		}
	})
	return bagSize >= item_dat.MAX_BAG_NUM
}

//检查背包是否可以插入 num 个 物品
//isBagFull 仅检查所有格子都用完了
func canAddItem(db *mdb.Database, itemId int16, num int16) bool {
	bagSize := 0
	itemDat := item_dat.GetItem(itemId)
	itemTypeDat := item_dat.GetItemType(itemDat.TypeId)
	var bagRoom int32 = 0
	db.Select.PlayerItem(func(row *mdb.PlayerItemRow) {
		if row.IsDressed() == 0 && row.BuyBackState() == 0 {
			bagSize++
			if row.ItemId() == itemId {
				bagRoom += int32(itemTypeDat.MaxNumInPos - row.Num())
			}
		}
	})
	bagRoom += (int32(item_dat.MAX_BAG_NUM-bagSize) * int32(itemTypeDat.MaxNumInPos))
	return bagRoom >= int32(num)
}

func undress(db *mdb.Database, playerItem *mdb.PlayerItem) {
	fail.When(playerItem == nil, "undress wrong id")
	fail.When(isBagFull(db), "undress bag full")

	item := item_dat.GetItem(playerItem.ItemId)
	itemType := item_dat.GetItemType(item.TypeId)

	var playerEquipment *mdb.PlayerEquipment
	db.Select.PlayerEquipment(func(row *mdb.PlayerEquipmentRow) {
		switch itemType.Id {
		case item_dat.TYPE_WEAPON:
			if row.WeaponId() == playerItem.Id {
				playerEquipment = row.GoObject()
				playerEquipment.WeaponId = 0
				row.Break()
			}
		case item_dat.TYPE_CLOTHES:
			if row.ClothesId() == playerItem.Id {
				playerEquipment = row.GoObject()
				playerEquipment.ClothesId = 0
				row.Break()
			}
		case item_dat.TYPE_ACCESSORIES:
			if row.AccessoriesId() == playerItem.Id {
				playerEquipment = row.GoObject()
				playerEquipment.AccessoriesId = 0
				row.Break()
			}
		case item_dat.TYPE_SHOE:
			if row.ShoeId() == playerItem.Id {
				playerEquipment = row.GoObject()
				playerEquipment.ShoeId = 0
				row.Break()
			}
		default:
			fail.When(true, "undress wrong equipment type")
		}
	})

	db.Update.PlayerEquipment(playerEquipment)
}

func setBuyBackItemChange(session *net.Session, ids []int64, out *item_api.BuyItemBack_Out) {
	state := module.State(session)
	db := state.Database
	out.Items = make([]item_api.BuyItemBack_Out_Items, 0, len(ids))

	for _, id := range ids {
		item := db.Lookup.PlayerItem(id)
		playerItemAppendix := tryGetPlayerItemAppendix(db, item.AppendixId)
		out.Items = append(out.Items, item_api.BuyItemBack_Out_Items{
			Id:            item.Id,
			ItemId:        item.ItemId,
			Num:           item.Num,
			Attack:        playerItemAppendix.Attack,
			Defence:       playerItemAppendix.Defence,
			Health:        playerItemAppendix.Health,
			Speed:         playerItemAppendix.Speed,
			Cultivation:   playerItemAppendix.Cultivation,
			HitLevel:      playerItemAppendix.HitLevel,
			CriticalLevel: playerItemAppendix.CriticalLevel,
			BlockLevel:    playerItemAppendix.BlockLevel,
			DestroyLevel:  playerItemAppendix.DestroyLevel,
			TenacityLevel: playerItemAppendix.TenacityLevel,
			DodgeLevel:    playerItemAppendix.DodgeLevel,
			RefineLevel:   item.RefineLevel,
			RecastAttr:    item_api.Attribute(playerItemAppendix.RecastAttr),
		})
	}
}

//道具消费流水表
func itemMoneyFlow(db *mdb.Database, item *item_dat.Item, num, Price, iMoneyType int32) {
	player_info := module.Player.GetPlayer(db)
	role_info := module.Role.GetMainRole(db)
	role_vip := module.VIP.VIPInfo(db)
	db.AddTLog(tlog.NewItemMoneyFlow(0, player_info.User, int32(item.TypeId), int32(item.Id), num, Price, int32(role_info.Level), iMoneyType, int32(role_vip.Level)))
}

func modifyPlayerItem(db *mdb.Database, heap *mdb.PlayerItem, items *notify_api.ItemChange_Out, rest, modifyType int16, AddOrReduce, itemFlowReason, xdlogEventType int32, xdlogEventParam string) {
	var num int16
	var itemAppendix *mdb.PlayerItemAppendix
	item := item_dat.GetItem(heap.ItemId)
	switch modifyType {
	case item_dat.MODIFY_DELETE:
		//删除已穿戴的装备需要将装备脱下
		if heap.IsDressed == 1 {
			undress(db, heap)
		}
		num = 0
		itemAppendix = tryGetPlayerItemAppendix(db, heap.AppendixId)
		delItemFromDB(db, heap, itemFlowReason, xdlogEventType)
	case item_dat.MODIFY_UPDATE:
		db.Update.PlayerItem(heap)
		tlog.PlayerItemFlowLog(db, item.Id, item.TypeId, rest, AddOrReduce, itemFlowReason)
		if AddOrReduce == tlog.ADD {
			xdlog.ItemLog(db, item.Id, rest, xdlogEventType, xdlogEventParam)
		} else if AddOrReduce == tlog.REDUCE {
			xdlog.PropsLog(db, item.Id, rest, xdlogEventType)
		}
		num = heap.Num
		itemAppendix = tryGetPlayerItemAppendix(db, heap.AppendixId)
	case item_dat.MODIFY_INSERT:
		db.Insert.PlayerItem(heap)
		tlog.PlayerItemFlowLog(db, item.Id, item.TypeId, rest, AddOrReduce, itemFlowReason)
		if AddOrReduce == tlog.ADD {
			xdlog.ItemLog(db, item.Id, rest, xdlogEventType, xdlogEventParam)
		} else if AddOrReduce == tlog.REDUCE {
			xdlog.PropsLog(db, item.Id, rest, xdlogEventType)
		}
		num = heap.Num
		itemAppendix = addItemAppendix(db, heap)
	}

	outItem := notify_api.ItemChange_Out_Items{
		Id:          heap.Id,
		ItemId:      heap.ItemId,
		Num:         num,
		RefineLevel: heap.RefineLevel,
	}
	if itemAppendix != nil {
		outItem.Attack = itemAppendix.Attack
		outItem.Defence = itemAppendix.Defence
		outItem.Health = itemAppendix.Health
		outItem.Speed = itemAppendix.Speed
		outItem.Cultivation = itemAppendix.Cultivation
		outItem.HitLevel = itemAppendix.HitLevel
		outItem.CriticalLevel = itemAppendix.CriticalLevel
		outItem.BlockLevel = itemAppendix.BlockLevel
		outItem.DestroyLevel = itemAppendix.DestroyLevel
		outItem.TenacityLevel = itemAppendix.TenacityLevel
		outItem.DodgeLevel = itemAppendix.DodgeLevel
		outItem.RecastAttr = notify_api.Attribute(itemAppendix.RecastAttr)
	}

	items.Items = append(items.Items, outItem)
}

func updateBuyRecord(state *module.SessionState, itemId int16) (record *mdb.PlayerPurchaseRecord) {
	state.Database.Select.PlayerPurchaseRecord(func(row *mdb.PlayerPurchaseRecordRow) {
		if row.ItemId() == itemId {
			record = row.GoObject()
			row.Break()
		}
	})
	if record == nil {
		record = &mdb.PlayerPurchaseRecord{
			ItemId:    itemId,
			Pid:       state.PlayerId,
			Timestamp: time.GetNowTime(),
		}
		state.Database.Insert.PlayerPurchaseRecord(record)
	} else if !time.IsToday(record.Timestamp) {
		record.Timestamp = time.GetNowTime()
		record.Num = 0
		state.Database.Update.PlayerPurchaseRecord(record)
	}
	return record
}

func updateSealedBookRecord(db *mdb.Database, item *item_dat.Item) {
	var sealedbookItemType int8
	switch item.TypeId {
	case item_dat.TYPE_WEAPON:
		sealedbookItemType = item_dat.STEALDBOOK_TYPE_WEAPON
	case item_dat.TYPE_CLOTHES, item_dat.TYPE_ACCESSORIES, item_dat.TYPE_SHOE:
		sealedbookItemType = item_dat.STEALDBOOK_TYPE_CLOTHES
	/*case item_dat.TYPE_ACCESSORIES:
	sealedbookItemType = item_dat.STEALDBOOK_TYPE_ACCESSORIES
	case item_dat.TYPE_SHOE:
		sealedbookItemType = item_dat.STEALDBOOK_TYPE_SHOE
	*/
	case item_dat.TYPE_BATTLE_PROPS:
		sealedbookItemType = item_dat.STEALDBOOK_TYPE_BATTLETOOLS
	default:
		sealedbookItemType = item_dat.STEALDBOOK_TYPE_NOEQUIPMENTS
	}

	var record *module.SealedBookRecord
	if sealedBook := item_dat.GetSealedBookInfo(sealedbookItemType, item.Id); sealedBook != nil {
		if session, exist := module.Player.GetPlayerOnline(db.PlayerId()); exist {
			state := module.State(session)
			record = state.GetSealedBookRecord()
		} else {

			record = &module.SealedBookRecord{}
			record.Load(db)
		}
		if _, result := record.FindRecord(sealedbookItemType, item.Id); !result {
			record.AddRecord(sealedbookItemType, item.Id, item_dat.STEALDBOOK_HAVING, db)
		}
	}
}
