package trader_dat

import (
	"core/fail"
)

type TraderSchedule struct {
	Expire    int64 //过期时间 0则永远有效
	Showup    int64 //出现时间，一天中的滴 X 秒
	Disappear int64 //消失时间，一天中的第  X秒
}

func TraderScheduleInfo(traderId int16) []*TraderSchedule {
	schedules, ok := MapTraderSchedule[traderId]
	fail.When(!ok, "trader is not configured")
	return schedules
}
