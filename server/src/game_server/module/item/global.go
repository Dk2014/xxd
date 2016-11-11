package item

import (
	"core/net"
	"game_server/api/protocol/item_api"
	"game_server/dat/item_dat"
	"game_server/dat/player_dat"
	"game_server/mdb"
	"game_server/module"
)

func init() {
	module.Item = ItemMod{}
}

type ItemMod struct {
}

//增加物品
func (this ItemMod) AddItem(db *mdb.Database, itemId int16, num int16, itemFLowReason, xdEventType int32, xdEventParam string) {
	addItem(db, itemId, num, itemFLowReason, xdEventType, xdEventParam)
}

//通过物品ID删除物品
func (this ItemMod) DelItemByItemId(db *mdb.Database, itemId int16, num int16, itemFLowReason, xdEventType int32) {
	delItemByItemId(db, itemId, num, itemFLowReason, xdEventType)
}

//检查物品数量
func (this ItemMod) CheckItemNum(state *module.SessionState, itemId int16, num int16) bool {
	return checkItemNum(state.Database, itemId, num)
}

//返回物品数量
func (this ItemMod) GetItemNum(db *mdb.Database, itemId int16) int16 {
	return getItemNum(db, itemId)
}

func (this ItemMod) ItemMoneyFlow(db *mdb.Database, item *item_dat.Item, num, Price, iMoneyType int32) {
	itemMoneyFlow(db, item, num, Price, iMoneyType)
}

//初始化角色装备
func (this ItemMod) InitRoleEquipment(state *module.SessionState, roleId int8) {
	state.Database.Insert.PlayerEquipment(&mdb.PlayerEquipment{
		Pid:           state.PlayerId,
		RoleId:        roleId,
		WeaponId:      0,
		ClothesId:     0,
		AccessoriesId: 0,
		ShoeId:        0,
	})
}

// 主动发送物品信息
func (this ItemMod) SendAllItems(session *net.Session) {
	out := &item_api.GetAllItems_Out{}
	GetAllItems(session, out)
	session.Send(out)
}

//批量增加物品 items itemId -> num
func (this ItemMod) BatchAddItem(db *mdb.Database, items map[int16]int16, itemFLowReason, xdEventType int32) {
	batchAddItem(db, items, itemFLowReason, xdEventType)
}

//统计物品数量
func (this ItemMod) CountItemNumByType(db *mdb.Database, itemType int32) map[int16]int32 {
	count := make(map[int16]int32)
	var itemDat *item_dat.Item
	db.Select.PlayerItem(func(row *mdb.PlayerItemRow) {
		itemDat = item_dat.GetItem(row.ItemId())
		if itemDat.TypeId == itemType {
			count[row.ItemId()] += int32(row.Num())
		}
	})
	return count
}

func (this ItemMod) Award(state *module.SessionState, awarder module.Awarder, itemFlowReason, moneyFlowReason, expFlowReason, xdEventType int32) {
	if awarder.Coins() > 0 {
		module.Player.IncMoney(state.Database, state.MoneyState, int64(awarder.Coins()), player_dat.COINS, moneyFlowReason, xdEventType, "")
	}
	if awarder.Ingot() > 0 {
		module.Player.IncMoney(state.Database, state.MoneyState, int64(awarder.Ingot()), player_dat.INGOT, moneyFlowReason, xdEventType, "")
	}
	if awarder.Exp() > 0 {
		module.Role.AddFormRoleExp(state, int64(awarder.Exp()), expFlowReason)
	}
	if len(awarder.Item()) > 0 {
		module.Item.BatchAddItem(state.Database, awarder.Item(), itemFlowReason, xdEventType)
	}
}

func (this ItemMod) CornucopiaCount(db *mdb.Database) int16 {
	return cornucopiaCount(db)
}
