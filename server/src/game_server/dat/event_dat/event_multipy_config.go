package event_dat

import (
	"core/mysql"
)

var (
	mapEventMultipyConfig map[int32]float32
)

//奖励翻倍活动配置
func loadEventMultiplyConfig(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM event_multiply_config ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}
	iCondition := res.Map("condition")
	iTimes := res.Map("times")

	mapEventMultipyConfig = map[int32]float32{}
	for _, row := range res.Rows {
		mapEventMultipyConfig[row.Int32(iCondition)] = row.Float32(iTimes)
	}
}

// ############### 对外接口实现 coding here ###############
type EventMultipy struct {
	StartUnixTime   int64
	EndUnixTime     int64
	DisposeUnixTime int64
	IsRelative      int8
	LTitle          string
	RTitle          string
	Content         string
	Tag             int8
	Weight          int16
	List            map[string]float32
}

func GetMultipyByKey(key int32) float32 {
	var result float32 = 1
	if times, ok := mapEventMultipyConfig[key]; ok {
		result = times
	}
	return result
}

//加载扩展配置
func LoadEventMultipy(data map[string]float32) {
	for key, times := range data {
		if times > 0 && len(key) > 0 {
			var realKey int32
			switch key {
			case "MISSION_COINS":
				realKey = MISSION_COINS
			case "MISSION_EXPS":
				realKey = MISSION_EXPS
			case "QQVIP_COINS":
				realKey = QQVIP_COINS
			case "QQVIP_EXPS":
				realKey = QQVIP_EXPS
			case "SUPERQQ_COINS":
				realKey = SUPERQQ_COINS
			case "SUPERQQ_EXPS":
				realKey = SUPERQQ_EXPS
			}
			if realKey > 0 {
				mapEventMultipyConfig[realKey] = times
			}
		}
	}
}
