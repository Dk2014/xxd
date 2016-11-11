package driving_sword_dat

import (
	"core/fail"
	"core/mysql"
)

func Load(db *mysql.Connection) {
	loadDrivingSword(db)
	loadDrivingSwordBuyCostConfig(db)
	loadDrivingSwordTreasureContent(db)
	loadDrivingSwordTreasureCount(db)
	loadTeleportCount(db)
	loadDrivingSwordVistingAwards(db)
	loadDrivingSwordExploring(db)

	//务必在加载完所有地图后执行
	genDrivingSwordCloudMap()
}

var (
	mapDrivingSwordBuyCostConfig map[int32]int32
)

func loadDrivingSwordBuyCostConfig(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM driving_sword_buy_cost_config ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iTimes := res.Map("times")
	iCost := res.Map("cost")

	mapDrivingSwordBuyCostConfig = make(map[int32]int32)
	for _, row := range res.Rows {
		mapDrivingSwordBuyCostConfig[row.Int32(iTimes)] = row.Int32(iCost)
	}
}

//获取当天当次购买云海御剑活动点次数所需元宝
func GetCost(times int32) int32 {
	times += 1 //次数从1,times为已购买次数
	cost, ok := mapDrivingSwordBuyCostConfig[times]
	if !ok {
		for max_times, max_cost := range mapDrivingSwordBuyCostConfig {
			if times > max_times { //超过设置上限，则按最大消耗元宝来
				cost = max_cost
			} else {
				fail.When(true, "arena buy cost config missing value ")
			}
		}
	}
	return cost
}

//获取云层地图 -- may be deprecated
func GetCloudMap(cloud int16, rand int) [][]int8 {
	return driving_sword_cloud_map[cloud-1][rand%len(driving_sword_cloud_map[cloud-1])]
}

//获取坐标事件
func GetCloudMapEvent(cloud int16, rand int, x, y uint8) int {
	return int(GetCloudMap(cloud, rand)[x][y])
}
