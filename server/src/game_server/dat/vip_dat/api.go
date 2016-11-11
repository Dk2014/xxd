package vip_dat

import (
	"core/fail"
	"core/mysql"
	"fmt"
)

const (
	YINJIESHUAXIN = 26
)

var (
	mapVipLevel map[int16]int64

	// [等级][特权ID] => 特权次数
	mapVipPrivilegeTime map[int16]map[int32]int16
	mapVipLevelupGift   map[int16][]*VipLevelupGift
)

func Load(db *mysql.Connection) {
	loadVipLevel(db)
	loadVipPrivilegeConfig(db)
	loadVipLevelupGift(db)
}

//加载VIP特权使用次数
func loadVipPrivilegeConfig(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM vip_privilege_config ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iPrivilegeId := res.Map("privilege_id")
	iLevel := res.Map("level")
	iTimes := res.Map("times")

	var level int16
	var privilegeId int32
	mapVipPrivilegeTime = make(map[int16]map[int32]int16, MAX_VIP_LEVEL)
	for _, row := range res.Rows {
		level = row.Int16(iLevel)
		privilegeId = row.Int32(iPrivilegeId)
		config, ok := mapVipPrivilegeTime[level]
		if !ok || config == nil {
			mapVipPrivilegeTime[level] = make(map[int32]int16, 0)
		}
		mapVipPrivilegeTime[level][privilegeId] = row.Int16(iTimes)
	}
}

type VipLevelupGift struct {
	ItemId  int16 // 物品ID
	ItemNum int16 // 物品数量
}

func loadVipLevelupGift(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM vip_levelup_gift ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iVipLevel := res.Map("vip_level")
	iItemId := res.Map("item_id")
	iItemNum := res.Map("item_num")

	mapVipLevelupGift = make(map[int16][]*VipLevelupGift)
	for _, row := range res.Rows {
		mapVipLevelupGift[row.Int16(iVipLevel)] = append(mapVipLevelupGift[row.Int16(iVipLevel)], &VipLevelupGift{
			ItemId:  row.Int16(iItemId),
			ItemNum: row.Int16(iItemNum),
		})
	}
}

//加载VIP等级元宝充值要求对应关系
func loadVipLevel(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM vip_level ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iLevel := res.Map("level")
	iIngot := res.Map("ingot")

	var pri_id int16
	mapVipLevel = make(map[int16]int64)
	for _, row := range res.Rows {
		pri_id = row.Int16(iLevel)
		mapVipLevel[pri_id] = row.Int64(iIngot)
	}
}

func GetVIPLevelInfo(level int16) (ingot int64, exist bool) {
	ingot, exist = mapVipLevel[level]
	return
}

func GetVIPPrivilegeTime(level int16, privilegeId int32) int16 {
	vipConfig, exist := mapVipPrivilegeTime[level]
	if !exist {
		return 0
	}
	//如果没有则必定是返回0
	return vipConfig[privilegeId]
}

func HaveVIPPrivilege(level int16, privilegeId int32) bool {
	if level == 0 {
		return false
	}
	vipConfig, exist := mapVipPrivilegeTime[level]
	fail.When(!exist, fmt.Sprintf("not config for vip level [%d]", level))

	_, exist = vipConfig[privilegeId]
	return exist
}

func GetVIPLevelUpGift(vipLevel int16) []*VipLevelupGift {
	gifts, exist := mapVipLevelupGift[vipLevel]
	if !exist {
		return nil
	}

	return gifts
}
