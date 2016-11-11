package event_dat

import (
	"core/fail"
	"core/mysql"
)

var (
	mapEvents{tf_key} map[int32]*EventDefaultAward
)

func loadEvents{tf_key}(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM events_{pt_val}_awards ORDER BY `require_{pt_val}` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iRequire{tf_key} := res.Map("require_{pt_val}")
	iIngot := res.Map("ingot")
	iCoins := res.Map("coins")
	iItem1Id := res.Map("item1_id")
	iItem1Num := res.Map("item1_num")
	iItem2Id := res.Map("item2_id")
	iItem2Num := res.Map("item2_num")
	iItem3Id := res.Map("item3_id")
	iItem3Num := res.Map("item3_num")
	iItem4Id := res.Map("item4_id")
	iItem4Num := res.Map("item4_num")
	iItem5Id := res.Map("item5_id")
	iItem5Num := res.Map("item5_num")

	var pri_require_{pt_val} int32
	mapEvents{tf_key} = map[int32]*EventDefaultAward{}

	for _, row := range res.Rows {
		pri_require_{pt_val} = row.Int32(iRequire{tf_key})
		mapEvents{tf_key}[pri_require_{pt_val}] = &EventDefaultAward{
			Ingot : row.Int16(iIngot),
			Coin  : row.Int32(iCoins),
			Item1Id : row.Int16(iItem1Id),
			Item1Num : row.Int16(iItem1Num),
			Item2Id : row.Int16(iItem2Id),
			Item2Num : row.Int16(iItem2Num),
			Item3Id : row.Int16(iItem3Id),
			Item3Num : row.Int16(iItem3Num),
			Item4Id : row.Int16(iItem4Id),
			Item4Num : row.Int16(iItem4Num),
			Item5Id : row.Int16(iItem5Id),
			Item5Num : row.Int16(iItem5Num),

	}

	}
}

// ############### 对外接口实现 coding here ###############
func GetEvent{tf_key}Award({pt_val} int32) (awards *EventDefaultAward, ok bool) {
	awards, ok = mapEvents{tf_key}[{pt_val}]
	return
}

type Events{tf_key} struct {
	Require{tf_key}    int32 // 需要{name_val}
	Ingot           int16 // 奖励元宝
	Coin            int32 // 奖励铜钱
	Item1Id         int16 // 物品1
	Item1Num        int16 // 物品1数量
	Item2Id         int16 // 物品2
	Item2Num        int16 // 物品2数量
	Item3Id         int16 // 物品3
	Item3Num        int16 // 物品3数量
	Item4Id         int16 // 物品4
	Item4Num        int16 // 物品4数量
	Item5Id         int16 // 物品5
	Item5Num        int16 // 物品5数量
}

type Events{tf_key}Ext struct {
	StartUnixTime   int64
	EndUnixTime     int64
	DisposeUnixTime int64
	List            []*Events{tf_key}
}

// {name_val}运营活动数据配置
func LoadEvents{tf_key}(list []*Events{tf_key}) {
	for _, item := range list {
			mapEvents{tf_key}[item.Require{tf_key}] = &EventDefaultAward{
			Ingot : item.Ingot,
			Coin  : item.Coin,
			Item1Id : item.Item1Id,
			Item1Num : item.Item1Num,
			Item2Id : item.Item2Id,
			Item2Num : item.Item2Num,
			Item3Id : item.Item3Id,
			Item3Num : item.Item3Num,
			Item4Id : item.Item4Id,
			Item4Num : item.Item4Num,
			Item5Id : item.Item5Id,
			Item5Num : item.Item5Num,

	}
	}
}

func GetMaxWeightIn{tf_key}() int32 {
	var max int32 = 0
	for {pt_val}, _ := range mapEvents{tf_key} {
		if {pt_val} >= max {
			max = {pt_val}
		}
	}
	return max
}

func GetNext{tf_key}(now int32) (next int32) {
	for {pt_val}, _ := range mapEvents{tf_key} {
		if {pt_val} <= now {
			continue
		}
		if next == 0 {
			next = {pt_val}
		} else if next > {pt_val} {
			next = {pt_val}
		}
	}
	return
}

func GetPlayer{tf_key}Award(awarded, maxAward int32) (newAwarded, nextAward int32, awards *EventDefaultAward) {
	fail.When(awarded > maxAward, "player event awarded status error")
	newAwarded = GetNext{tf_key}(awarded)
	if newAwarded <= maxAward {
		awards = mapEvents{tf_key}[newAwarded]
		nextAward = GetNext{tf_key}(newAwarded)
	}
	return
}
