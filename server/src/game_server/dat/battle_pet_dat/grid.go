package battle_pet_dat

/*

import (
	"core/mysql"
)

var (
	mapBattlePetGridLevel        map[int16]*BattlePetGridLevel              //level -> data
	mapBattlePetGridAttribute    map[int8]map[int16]*BattlePetGridAttribute //gridId level  -> data
	mapBattlePetGridUpgradePrice []*BattlePetGridUpgradePrice
)

type BattlePetGridUpgradePrice struct {
	Times int32
	Price int64
}

func loadBattlePetGridUpgradePrice(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM battle_pet_grid_upgrade_price ORDER BY `times` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iTimes := res.Map("times")
	iCost := res.Map("cost")

	for _, row := range res.Rows {
		mapBattlePetGridUpgradePrice = append(mapBattlePetGridUpgradePrice, &BattlePetGridUpgradePrice{
			Times: row.Int32(iTimes),
			Price: int64(row.Int64(iCost)),
		})
	}
}

type BattlePetGridLevel struct {
	Exp          int64 // 升级到下一级需要的经验
	CostSoulNum  int16 // 升级消耗灵魂
	MinAddExp    int64 // 最小经验加值
	MaxAddExp    int64 // 最大经验加值
	RequireLevel int16 // 培养要求等级（即玩家达到这个等级才能升级格子)
}

func loadBattlePetGridLevel(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM battle_pet_grid_level ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iLevel := res.Map("level")
	iExp := res.Map("exp")
	iCostSoulNum := res.Map("cost_soul_num")
	iMinAddExp := res.Map("min_add_exp")
	iMaxAddExp := res.Map("max_add_exp")
	iRequireLevel := res.Map("require_level")

	var pri_id int16
	mapBattlePetGridLevel = map[int16]*BattlePetGridLevel{}
	for _, row := range res.Rows {
		pri_id = row.Int16(iLevel)
		mapBattlePetGridLevel[pri_id] = &BattlePetGridLevel{
			Exp:          row.Int64(iExp),
			CostSoulNum:  row.Int16(iCostSoulNum),
			MinAddExp:    row.Int64(iMinAddExp),
			MaxAddExp:    row.Int64(iMaxAddExp),
			RequireLevel: row.Int16(iRequireLevel),
		}
	}
}

//根据格子等级获取相关信息
func GetBattleGridLevelInfo(level int16) *BattlePetGridLevel {
	return mapBattlePetGridLevel[level]
}

func loadBattlePetGridAttribute(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM battle_pet_grid_attribute ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iGridId := res.Map("grid_id")
	iLevel := res.Map("level")
	iHealth := res.Map("health")
	iAttack := res.Map("attack")
	iDefence := res.Map("defence")
	iSpeed := res.Map("speed")

	var gridId int8
	var level int16
	mapBattlePetGridAttribute = map[int8]map[int16]*BattlePetGridAttribute{}
	for _, row := range res.Rows {
		gridId = row.Int8(iGridId)
		level = row.Int16(iLevel)
		if mapBattlePetGridAttribute[gridId] == nil {
			mapBattlePetGridAttribute[gridId] = map[int16]*BattlePetGridAttribute{}
		}
		mapBattlePetGridAttribute[gridId][level] = &BattlePetGridAttribute{
			Health:  row.Int32(iHealth),
			Attack:  row.Int32(iAttack),
			Defence: row.Int32(iDefence),
			Speed:   row.Int32(iSpeed),
		}
	}
}

type BattlePetGridAttribute struct {
	//GridId  int8  // 格子ID
	//Level   int16 // 格子等级
	Health  int32 // 生命
	Attack  int32 // 攻击
	Defence int32 // 防御
	Speed   int32 // 速度
}

func GetBattlePetGridAttribute(gridId int8, level int16) *BattlePetGridAttribute {
	return mapBattlePetGridAttribute[gridId][level]
}

//获取魂魄单价
func GetBattlePetGhostPrice(times int32) (price int64) {
	for _, data := range mapBattlePetGridUpgradePrice {
		if data.Times > times {
			break
		}
		price = data.Price
	}
	return price
}
*/
