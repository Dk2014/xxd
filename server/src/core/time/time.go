package time

import (
	"sync"
	"sync/atomic"
	"time"
)

var (
	offset            int64
	_atomic_timestamp int64 // 原子钟
)

func init() {
	var once sync.Once
	initOnce := func() {
		_, _offset := time.Now().Zone()
		offset = int64(_offset)

		atomic.StoreInt64(&_atomic_timestamp, time.Now().Unix())

		go func() {
			c := time.Tick(1 * time.Second)
			for now := range c {
				atomic.StoreInt64(&_atomic_timestamp, now.Unix())
			}
		}()

	}

	once.Do(initOnce)
}

// 获取当天零点时间戳
func GetTodayZero() int64 {
	nowTime := time.Now()

	return time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), 0, 0, 0, 0, nowTime.Location()).Unix()
}

// 获取当前时间戳
func GetNowTime() int64 {
	return _atomic_timestamp // cost 0.33ns/op , but may not safe
	// return atomic.LoadInt64(&_atomic_timestamp) // cost 5.6ns/op , safe
	// return time.Now().Unix() // cost 13.4ns/op , accurate
}

// 获取当前时间的上一个整点时间 当前时间 2013-10-17 14:14:34 返回 2013-10-17 14:00:00
func GetLastHourTime(nowTime time.Time) (lastHourTimeUnix int64) {
	lastHourTime := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), nowTime.Hour(), 0, 0, 0, time.Local)
	lastHourTimeUnix = lastHourTime.Unix()
	return
}

// 获取当前时间的下一个整点时间
func GetNextHourTime(nowTime time.Time) (nextHourTimeUnix int64) {
	nextHourTime := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), nowTime.Hour()+1, 0, 0, 0, time.Local) // Hour 可传入25
	nextHourTimeUnix = nextHourTime.Unix()
	return
}

// 获取当前时间的23点59分59秒 当前时间 2013-12-24 14:14:34 返回 2013-12-24 23:59:59
func GetTodayLastTime(nowTime time.Time) int64 {
	return time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), 23, 59, 59, 0, time.Local).Unix()
}

// 获取下一个指定的时间点
func GetNextTodayPointHour(expectHour int, nowTime time.Time) int64 {
	nextDay := nowTime.Day()
	if nowTime.Hour() >= expectHour {
		nextDay += 1
	}
	return time.Date(nowTime.Year(), nowTime.Month(), nextDay, expectHour, 0, 0, 0, time.Local).Unix()
}

// 获取下一个指定的时间点
func GetNextMonthPointHour(expectDay int, expectHour int, nowTime time.Time) int64 {
	nextMonth := nowTime.Month() + 1
	return time.Date(nowTime.Year(), nextMonth, expectDay, expectHour, 0, 0, 0, time.Local).Unix()
}

/*
获取当前时间的第X天的23点59分59秒 第一天的话 就传1
比如 当前时间 2013-12-24 14:14:34 当前时间第一天的23点59分59秒
传入1
返回 2013-12-27 23:59:59
*/
func GetDayLastTime(nowTime time.Time, dayNum int) int64 {
	return time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day()+(dayNum-1), 23, 59, 59, 0, time.Local).Unix()
}

func GetDayFirstTime(nowTime time.Time, dayNum int) int64 {
	return time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day()+(dayNum-1), 0, 0, 0, 0, time.Local).Unix()
}

// 传入的时间是否今天
func IsToday(timeUnixTmp int64) (isToday bool) {
	timeTmp := time.Unix(timeUnixTmp, 0)
	timeToday := time.Now()
	if timeTmp.Year() == timeToday.Year() && timeTmp.Month() == timeToday.Month() &&
		timeTmp.Day() == timeToday.Day() {
		isToday = true
	}
	return
}

// 传入的时间是否本周
func IsThisWeek(timeUnixTmp int64) (isThisWeek bool) {
	timeTmp := time.Unix(timeUnixTmp, 0)
	timeToday := time.Now()
	if timeTmp.Year() == timeToday.Year() && timeTmp.Month() == timeToday.Month() && timeTmp.Day()/7 == timeToday.Day()/7 {
		isThisWeek = true
	}
	return
}

func GetNowDay() int {
	unixNow := GetNowTime() + offset
	return int(unixNow / 86400)
}

func GetNowDayFromUnix(unixTime int64) int {
	return int((unixTime + offset) / 86400)
}

// 指定时间段（小时）,判断某一时间戳是否在[0~23]周期的指定时间段内
func IsInPointHour(expectHour int, specTime int64) bool {
	now := time.Now()
	specUnix := time.Unix(specTime, 0)

	st1 := time.Date(now.Year(), now.Month(), now.Day(), expectHour, 0, 0, 0, time.Local).Unix()
	st2 := time.Date(specUnix.Year(), specUnix.Month(), specUnix.Day(), expectHour, 0, 0, 0, time.Local).Unix()

	if now.Hour() < expectHour {
		st1 -= 86400
	}

	if specUnix.Hour() < expectHour {
		st2 -= 86400
	}

	return st1 == st2
}

//指定开始和结束时间(小时), 判断给定得时间戳是否在当天的开始小时和结束小时段内
func IsInRangeHour(expectStartHour int, expectEndHour int, specTime int64) bool {
	dt := specTime - GetTodayZero()
	return int64(expectStartHour)*3600 <= dt && dt <= int64(expectEndHour)*3600
}

func GetNowWeekByExpectHour(expectHour int) (weekDay time.Weekday) {
	now := time.Now()
	weekDay = now.Weekday()

	if now.Hour() < expectHour {
		if weekDay == time.Sunday {
			weekDay = time.Saturday
		} else {
			weekDay -= 1
		}
	}

	return
}
