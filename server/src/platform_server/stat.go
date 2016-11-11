package platform_server

import (
	"core/debug"
	"core/log"
	"core/redis"
	"fmt"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// 更好的服务器状态输出
var (
	_t_start = time.Now()
)

type ReqStatus struct {
	ReqBase
}

type procStatus struct {
}

func (proc procStatus) Req() interface{} {
	return &ReqStatus{}
}

func (proc procStatus) Validate(in interface{}) {
}

func (proc procStatus) Process(in interface{}) (interface{}, error) {
	req := in.(*ReqStatus)

	memStat := new(runtime.MemStats)
	runtime.ReadMemStats(memStat)
	tnow := time.Now()
	c := GetDBConn(req.App)
	defer c.Close()
	server_list_rev, _ := redis.String(c.Do("GET", RedisKey_GameServerListCurrentRev()))
	stats := getStatsInfo()
	return &map[string]interface{}{
		"mem":                memStat.Alloc,
		"goroutine":          runtime.NumGoroutine(),
		"gc":                 memStat.NumGC,
		"pause":              memStat.PauseTotalNs / uint64(time.Millisecond),
		"up":                 strings.TrimRight(tnow.Sub(_t_start).String(), "s0123456789."),
		"server_list_rev":    server_list_rev,
		"server_update_time": _server_update_time,
		"stats":              stats,
	}, nil
}

var (
	_msg   int64 = 0
	_err   int64 = 0
	_secs  int64 = 0
	_mutex sync.Mutex
)

type Printfer func(string, ...interface{})

func IncrMsgBy(n int64) {
	atomic.AddInt64(&(_msg), n)
}

func IncrErrBy(n int64) {
	atomic.AddInt64(&(_err), n)
}

//d: 间隔时间，单位s
func Tick(d int64) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Errorf(`plarmore_server.stat.Tick
Error = %v
Stack = 
%s`,
					err,
					debug.Stack(1, "    "),
				)
			}
		}()
		for {
			time.Sleep(time.Second)
			atomic.AddInt64(&(_secs), 1)
			if atomic.CompareAndSwapInt64(&(_secs), d, 0) {
				atomic.StoreInt64(&(_msg), 0)
				atomic.StoreInt64(&(_err), 0)
			}
		}
	}()
}

//输出统计信息
//d: 间隔时间，单位s
func TickTo(d int64, p Printfer) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Errorf(`plarmore_server.stat.TickTo
Error = %v
Stack = 
%s`,
					err,
					debug.Stack(1, "    "),
				)
			}
		}()
		for {
			time.Sleep(time.Second)
			atomic.AddInt64(&(_secs), 1)
			if atomic.CompareAndSwapInt64(&(_secs), d, 0) {
				write(p, d)
				atomic.StoreInt64(&(_msg), 0)
				atomic.StoreInt64(&(_err), 0)
			}
		}
	}()
}

func getStatsInfo() (info string) {
	_mutex.Lock()
	defer _mutex.Unlock()

	var msg_avg, err_avg int64
	if _secs == 0 {
		msg_avg = 0
		err_avg = 0
	} else {
		msg_avg = _msg / _secs
		err_avg = _err / _secs
	}

	out_temp := "duration:%ds msg_count:%d msg_avg:%d err_count:%d err_avg:%d"
	return fmt.Sprintf(out_temp, _secs, _msg, msg_avg, _err, err_avg)
}

func write(p Printfer, d int64) {
	_mutex.Lock()
	defer _mutex.Unlock()

	var msg_avg, err_avg int64
	if d == 0 {
		msg_avg = 0
		err_avg = 0
	} else {
		msg_avg = _msg / d
		err_avg = _err / d
	}
	out_temp := "[stats]duration:%ds msg_count:%d msg_avg:%d err_count:%d err_avg:%d"
	p(out_temp, d, _msg, msg_avg, _err, err_avg)
}
