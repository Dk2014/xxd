package clique_dat

import (
	"core/fail"
	"core/mysql"
	"fmt"
)

var (
	mapCliqueKongfu                  map[int32]*CliqueKongfu
	mapCliqueKongfuBuildingLevelInfo map[int32][]*CliqueBuildingKongfu //[building][level]
	mapCliqueKongfuLevelupContrib    map[int32][]int64
)

type CliqueKongfu struct {
	Id                   int32  // 标识ID
	Building             int32  // 所属建筑等级
	Name                 string // 武功名称
	RequireBuildingLevel int16  // 要求建筑等级
	AttribType           int8   // 属性类型
	AttribValue          int32  // 属性值
}

func loadCliqueKongfu(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM clique_kongfu ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iId := res.Map("id")
	iBuilding := res.Map("building")
	iName := res.Map("name")
	iRequireBuildingLevel := res.Map("require_building_level")
	iAttribType := res.Map("attrib_type")
	iAttribValue := res.Map("attrib_value")

	var pri_id int32
	mapCliqueKongfu = map[int32]*CliqueKongfu{}
	for _, row := range res.Rows {
		pri_id = row.Int32(iId)
		mapCliqueKongfu[pri_id] = &CliqueKongfu{
			Id:                   pri_id,
			Building:             row.Int32(iBuilding),
			Name:                 row.Str(iName),
			RequireBuildingLevel: row.Int16(iRequireBuildingLevel),
			AttribType:           row.Int8(iAttribType),
			AttribValue:          row.Int32(iAttribValue),
		}
	}
}

type CliqueBuildingKongfu struct {
	//Building     int32  // 所属建筑等级
	//Level        int16  // 等级
	UpgradeCoins int64 // 升级所需铜钱
	CanUpgrade   bool  //是否可以游泳下一级
}

func loadCliqueBuildingKongfuLevelInfo(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM clique_building_kongfu ORDER BY `building`, `level` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iBuilding := res.Map("building")
	//iLevel := res.Map("level")
	iUpgradeCoins := res.Map("upgrade_coins")

	var buildingId int32
	mapCliqueKongfuBuildingLevelInfo = map[int32][]*CliqueBuildingKongfu{}
	for _, row := range res.Rows {
		buildingId = row.Int32(iBuilding)
		if len(mapCliqueKongfuBuildingLevelInfo[buildingId]) > 0 {
			mapCliqueKongfuBuildingLevelInfo[buildingId][len(mapCliqueKongfuBuildingLevelInfo[buildingId])-1].CanUpgrade = true
		}
		mapCliqueKongfuBuildingLevelInfo[buildingId] = append(mapCliqueKongfuBuildingLevelInfo[buildingId], &CliqueBuildingKongfu{
			UpgradeCoins: row.Int64(iUpgradeCoins),
		})
	}
}

type CliqueKongfuLevel struct {
	KongfuId       int32 // 心法ID
	ConsumeContrib int64 // 消耗帮派贡献
	Level          int16 // 等级
}

func loadCliqueKongfuLevel(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM clique_kongfu_level ORDER BY `kongfu_id`,`level` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iKongfuId := res.Map("kongfu_id")
	iConsumeContrib := res.Map("consume_contrib")
	//iLevel := res.Map("level")

	var kongfuId int32
	mapCliqueKongfuLevelupContrib = map[int32][]int64{}
	for _, row := range res.Rows {
		kongfuId = row.Int32(iKongfuId)
		mapCliqueKongfuLevelupContrib[kongfuId] = append(mapCliqueKongfuLevelupContrib[kongfuId], row.Int64(iConsumeContrib))
	}
}

func GetKongfuById(kongfuId int32) *CliqueKongfu {
	kongfu := mapCliqueKongfu[kongfuId]
	fail.When(kongfu == nil, "undefined kongfu")
	return kongfu
}

func GetKongfuBuildingDat(buildingId int32, level int16) *CliqueBuildingKongfu {
	dat := mapCliqueKongfuBuildingLevelInfo[buildingId]
	//index 从0起
	fail.When(int(level+1) > len(dat), fmt.Sprintf(" kongfu building level dat not found: buildingId %d, level %d, len(dat) %d", buildingId, level, len(dat)))
	return dat[level]
}

func GetKongfuLevelConsume(kongfuId int32, level int16) (int64, bool) {
	dat := mapCliqueKongfuLevelupContrib[kongfuId]
	//index 从1起
	if int(level) <= len(dat) {
		return dat[level-1], true
	}

	return 0, false
}
