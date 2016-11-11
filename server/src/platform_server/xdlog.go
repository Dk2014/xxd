package platform_server

import (
	"core/log"
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"strings"
	"sync"
	"time"
)

var (
	_LOG_PATH_ string
	fLocker    sync.Mutex
	fLog       *os.File

	flushDuartion = time.Second * 2

	cacheLogs = make([]string, 0)
	logLocker sync.Mutex
)

type UserCreateLog struct {
	Type    string
	Cid     int
	Sid     int32
	Account string
	Time    int64
	Ip      string
}

func LogAccountCreate(logObj *UserCreateLog) {
	log := fmt.Sprintf(`{"tag":"%s","cid":%d,"sid":%d,"account":"%s","time":%d,"ip":"%s"}`,
		logObj.Type, logObj.Cid, logObj.Sid,
		logObj.Account, logObj.Time, logObj.Ip)

	logLocker.Lock()
	defer logLocker.Unlock()
	cacheLogs = append(cacheLogs, log)
}

func flushLog() {
	if len(cacheLogs) == 0 {
		return
	}

	logLocker.Lock()
	defer logLocker.Unlock()

	content := strings.Join(cacheLogs, "\n")
	_, err := fLog.Write([]byte(content + "\n"))
	if err != nil {
		log.Error(err.Error())
		return
	}
	cacheLogs = make([]string, 0)
}

func doTick() {
	defer func() {
		if r := recover(); r != nil {
			log.Errorf("Recovered in %v: %v", r, string(debug.Stack()))
		}
	}()
	ticker := time.NewTicker(flushDuartion)

	for {
		select {
		case <-ticker.C:
			flushLog()
		}
	}
}

func InitXdlog(logPath string) {
	dir, _ := os.Stat(logPath)
	if dir == nil {
		os.Mkdir(logPath, 0777)
	}
	_LOG_PATH_ = logPath

	go func() {
		defer func() { recover() }()
		for {
			now := time.Now()
			switchLogFile(now)
			// 计算此刻到第二天零点的时间
			t := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, now.Nanosecond(), now.Location())
			duration := t.Sub(now)
			time.Sleep(duration)
			runtime.Gosched()
		}
	}()

	go doTick()
}

func switchLogFile(now time.Time) {
	// file.Fd()==18446744073709551615为true  文件已经close 或者 没有打开
	// 目前正在查找是否有更加直观的写法
	fLocker.Lock()
	defer fLocker.Unlock()

	if fLog != nil {
		fLog.Close()
	}

	var err error
	logName := _LOG_PATH_ + "/" + now.Format("2006-01-02") + ".log"
	fLog, err = os.OpenFile(logName, os.O_WRONLY|os.O_CREATE|os.O_APPEND|os.O_SYNC, 0777)
	if err != nil {
		log.Error(err.Error())
	}
}

func testSendLog() {
	ticker := time.NewTicker(time.Second)

	for {
		select {
		case <-ticker.C:
			logObj := UserCreateLog{
				Type:    "account",
				Cid:     1,
				Sid:     123,
				Account: "daizong_test_openid",
				Time:    time.Now().Unix(),
				Ip:      "127.0.0.1",
			}
			LogAccountCreate(&logObj)
		}
	}
}
