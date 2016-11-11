package payments_rule_dat

import (
	"core/fail"
	"core/mysql"
	"core/time"
	"encoding/json"
	"strconv"
)

var (
	mapPaymentsPresentConfig *PhysicalBuyCostConfig
)

func Load(db *mysql.Connection) {
	loadPaymentsPresentConfig(db)
}

type PhysicalBuyCostConfig struct {
	Rule      string // 充值返利规则
	Begintime int64  // 活动开始时间
	Endtime   int64  // 活动结束时间
}

func loadPaymentsPresentConfig(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT `rule`,`begin_time`,`end_time` FROM payments_rule"), -1)
	if err != nil {
		panic(err)
	}

	iRule := res.Map("rule")
	iBegintime := res.Map("begin_time")
	iEndtime := res.Map("end_time")

	mapPaymentsPresentConfig = &PhysicalBuyCostConfig{}
	for _, row := range res.Rows {
		mapPaymentsPresentConfig = &PhysicalBuyCostConfig{
			Rule:      row.Str(iRule),
			Begintime: row.Int64(iBegintime),
			Endtime:   row.Int64(iEndtime),
		}
		break
	}
}

// ############### 对外接口实现 coding here ###############
//按照当天购买体力次数获取所需消耗元宝
func GetPresent(money int64) int64 {
	nowtime := time.GetNowTime()
	var present, key int64
	rule := make(map[string]int64)
	if mapPaymentsPresentConfig.Begintime <= nowtime && nowtime <= mapPaymentsPresentConfig.Endtime && mapPaymentsPresentConfig.Rule != "" {
		err := json.Unmarshal([]byte(mapPaymentsPresentConfig.Rule), &rule)
		fail.When(err != nil, err)
		if len(rule) > 0 {
			for k, v := range rule {
				k1, _ := strconv.ParseInt(k, 10, 64)
				if k1 <= money && k1 > key {
					present = v
					key = k1
				}
			}
		}
	}
	return present
}

func ModifyRule(rule string, begintime, endtime int64) {
	mapPaymentsPresentConfig = &PhysicalBuyCostConfig{
		Rule:      rule,
		Begintime: begintime,
		Endtime:   endtime,
	}
}
