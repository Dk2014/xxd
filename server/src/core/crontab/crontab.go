package crontab

import (
	debug "core/dbgutil"
	"core/log"
	"sync"
	"time"
)

const EVERY_WEEKDAY = -1

type Time struct {
	Weekday time.Weekday
	Hour    int
	Minute  int

	next    *Time
	minutes int
}

var (
	crontabs     map[string]*time.Timer
	crontabMutex *sync.Mutex
)

func init() {
	crontabs = make(map[string]*time.Timer, 100)
	crontabMutex = new(sync.Mutex)
}

// 新增crontab项目
func Add(name string, config []*Time, callback func()) {
	sortConfig := sort(config)

	crontabMutex.Lock()
	defer crontabMutex.Unlock()

	newTimer(name, sortConfig, callback, false)
}

// 删除crontab项目
func Del(name string) {
	if timer, ok := crontabs[name]; ok {
		timer.Stop()
		delete(crontabs, name)
	}
}

// @TODO 顺序执行有可能因某个方法执行时间过长，而错过了下个执行时间
func newTimer(name string, sortConfig *Time, cb func(), resetAtZero bool) {
	now := time.Now()
	minutes := now.Hour()*60 + now.Minute()

	var (
		d       time.Duration = -1
		execute bool
	)

	for p := sortConfig; p != nil; p = p.next {
		weekday := p.Weekday
		if weekday != EVERY_WEEKDAY && weekday != now.Weekday() {
			continue
		}

		if resetAtZero && minutes == p.minutes {
			// 零点重置时刻，如果有00:00任务，则需要立即执行
			d = 0
			execute = true
			break
		}

		if minutes < p.minutes {
			d = time.Date(now.Year(), now.Month(), now.Day(), p.Hour, p.Minute, 0, 0, now.Location()).Sub(now)

			execute = true

			break
		}
	}

	resetAtZero = false

	if d == -1 {
		// 没找到任何执行任务，则在第二天零点空跑定时器（不执行任务）
		d = time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location()).Sub(now)
		resetAtZero = true
	}

	crontabs[name] = time.AfterFunc(d, func() {
		defer func() {
			if err := recover(); err != nil {
				log.Errorf(`crontab
Error = %v
Stack = 
%s`,
					err,
					debug.Stack(1, "    "),
				)
			}
		}()

		if execute {
			log.Infof("Crontab Execute %s", name)
			cb()
		}

		newTimer(name, sortConfig, cb, resetAtZero)
	})

	if t, exist := crontabs[name]; exist && t != nil {
		log.Infof("Crontab Start %s", name)
	} else {
		log.Infof("Crontab Start Fail %s", name)
	}
}

func sort(config []*Time) (sortConfig *Time) {
	for _, ct := range config {
		ct.minutes = ct.Hour*60 + ct.Minute

		if sortConfig == nil {
			sortConfig = ct

			continue
		}

		isIns := false
		pp := &sortConfig

		for *pp != nil {
			if (*pp).minutes > ct.minutes {
				isIns = true

				ct.next = *pp
				*pp = ct

				break
			}

			pp = &((*pp).next)
		}

		if !isIns {
			*pp = ct
		}
	}

	return
}
