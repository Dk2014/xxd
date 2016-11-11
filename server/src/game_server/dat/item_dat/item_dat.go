package item_dat

import (
	"core/fail"
	"core/mysql"
)

var (
	mapItem          map[int16]*Item
	mapItemByType    map[int32][]*Item
	mapItemCostprops map[int16]*ItemCostprops
)

func Load(db *mysql.Connection) {
	loadItem(db)
	loadItemType(db)
	loadItemExchange(db)
	loadItemBoxContent(db)
	loadEquipmentAppendix(db)
	loadEquipmentDecompose(db)
	loadEquipmentRecast(db)
	//loadEquipmentRefine(db)
	//loadEquipmentRefineLevel(db)
	loadEquipmentRefineNew(db)
	loadItemCostprops(db)
	loadPurchaseLimit(db)
	loadDragonBallConfig(db)
	loadItemReflectConfig(db)
	loadEquipmentResonance(db)
	loadSealedBook(db)
}

type Item struct {
	Id            int16  // 物品ID
	TypeId        int32  // 类型ID
	Quality       int8   // 品质
	Name          string // 物品名称
	Level         int32  // 需求等级
	Desc          string // 物品描述
	Price         int64  // 物品售价
	Sign          string // 资源标识
	CanUse        int8   // 是否可在格子中使用，0不能，1反之
	FuncId        int32  // 使用的功能限制
	RenewIngot    int32  // 续费的元宝价格
	UseIngot      int32  // 使用的元宝价格
	ValidHours    int32  // 有效小时数
	EquipTypeId   int32  // 装备等级类型
	Health        int32  // 生命
	Speed         int32  // 速度
	Attack        int32  // 攻击
	Defence       int32  // 防御
	EquipRoleId   int8   // 可装备角色ID
	AppendixNum   int8   // 追加属性数
	AppendixLevel int32  // 追加属性等级
	CanBatch      int8   // 是否可以批量使用,0:非，1:是
	RefineBase    int32  // 精炼基础值
	ActId         int32  // 触发功能
	CanSell       bool   // 是否可以出售
}

func (item *Item) IsRareItem() bool {
	return item.Quality == ITEM_QUALITY_GOLD || item.Quality == ITEM_QUALITY_ORANGE
}

func loadItem(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM item ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iId := res.Map("id")
	iTypeId := res.Map("type_id")
	iQuality := res.Map("quality")
	iName := res.Map("name")
	iLevel := res.Map("level")
	iDesc := res.Map("desc")
	iPrice := res.Map("price")
	iSign := res.Map("sign")
	iCanUse := res.Map("can_use")
	iFuncId := res.Map("func_id")
	iRenewIngot := res.Map("renew_ingot")
	iUseIngot := res.Map("use_ingot")
	iValidHours := res.Map("valid_hours")
	iEquipTypeId := res.Map("equip_type_id")
	iHealth := res.Map("health")
	iSpeed := res.Map("speed")
	iAttack := res.Map("attack")
	iDefence := res.Map("defence")
	iEquipRoleId := res.Map("equip_role_id")
	iAppendixNum := res.Map("appendix_num")
	iAppendixLevel := res.Map("appendix_level")
	iCanBatch := res.Map("can_batch")
	iRefineBase := res.Map("refine_base")
	iActId := res.Map("act_id")
	iCanSell := res.Map("can_sell")

	var pri_id int16
	mapItem = map[int16]*Item{}
	mapItemByType = map[int32][]*Item{}
	for _, row := range res.Rows {
		pri_id = row.Int16(iId)

		item := &Item{
			Id:            pri_id,
			TypeId:        row.Int32(iTypeId),
			Quality:       row.Int8(iQuality),
			Name:          row.Str(iName),
			Level:         row.Int32(iLevel),
			Desc:          row.Str(iDesc),
			Price:         row.Int64(iPrice),
			Sign:          row.Str(iSign),
			CanUse:        row.Int8(iCanUse),
			FuncId:        row.Int32(iFuncId),
			RenewIngot:    row.Int32(iRenewIngot),
			UseIngot:      row.Int32(iUseIngot),
			ValidHours:    row.Int32(iValidHours),
			EquipTypeId:   row.Int32(iEquipTypeId),
			Health:        row.Int32(iHealth),
			Speed:         row.Int32(iSpeed),
			Attack:        row.Int32(iAttack),
			Defence:       row.Int32(iDefence),
			EquipRoleId:   row.Int8(iEquipRoleId),
			AppendixNum:   row.Int8(iAppendixNum),
			AppendixLevel: row.Int32(iAppendixLevel),
			CanBatch:      row.Int8(iCanBatch),
			RefineBase:    row.Int32(iRefineBase),
			ActId:         row.Int32(iActId),
			CanSell:       row.Int8(iCanSell) == 1,
		}

		mapItem[pri_id] = item

		if mapItemByType[row.Int32(iTypeId)] == nil {
			mapItemByType[row.Int32(iTypeId)] = []*Item{}
		}

		mapItemByType[row.Int32(iTypeId)] = append(mapItemByType[row.Int32(iTypeId)], item)
	}
}

type ItemCostprops struct {
	Type  int8  // 消耗类型； 0 - 经验； 1 - 体力; 2 - 友情
	Value int32 // 值
}

func loadItemCostprops(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM item_costprops ORDER BY `item_id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iItemId := res.Map("item_id")
	iType := res.Map("type")
	iValue := res.Map("value")

	var pri_item_id int16
	mapItemCostprops = map[int16]*ItemCostprops{}
	for _, row := range res.Rows {
		pri_item_id = row.Int16(iItemId)
		mapItemCostprops[pri_item_id] = &ItemCostprops{
			Type:  row.Int8(iType),
			Value: row.Int32(iValue),
		}
	}
}

// ############### 对外接口实现 coding here ###############

func GetItem(id int16) *Item {
	v, ok := mapItem[id]
	fail.When(!ok, "item wrong id")
	return v
}

func GetItemsByType(typeId int32) (items []*Item) {
	items = mapItemByType[typeId]
	fail.When(items == nil, "wrong items id")
	return items
}

func GetCostpropsWithItemId(id int16) *ItemCostprops {
	v, ok := mapItemCostprops[id]
	fail.When(!ok, "GetCostpropsWithItemId id error")
	return v
}
