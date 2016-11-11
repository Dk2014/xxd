package main

import (
	"core/debug"
	"core/log"
	"core/time"
	"cron_server/config"
	"cron_server/database"
	"flag"
	"runtime"
	gotime "time"
)

var (
	configFilePath string
)

func init() {
	flag.StringVar(&configFilePath, "conf", "cron_conf.json", "server config file")
	flag.Parse()
}

func main() {
	done := make(chan bool, 1)
	runtime.GOMAXPROCS(runtime.NumCPU())
	if err := config.Load(configFilePath); err != nil {
		panic(err)
	}
	log.Setup(config.SCFG.LogDir, false)

	if err := database.OnlineInit(config.SCFG.OnlineDB.ToConnectStrOnline()); err != nil {
		panic(err)
	}
	go SnapShot()
	go OnlineCount()
	<-done
}

func SnapShot() {
	timeNow := time.GetNowTime()
	nextMonth := time.GetNextMonthPointHour(0, 0, gotime.Now())
	needSeconds := nextMonth - timeNow
	gotime.AfterFunc(gotime.Second*gotime.Duration(needSeconds), func() {
		defer func() {
			if err := recover(); err != nil {
				log.Errorf("SnapShot() error: %v\n Stack %s \n", err, debug.Stack(1, "    "))
			}
		}()
		defer SnapShot()
		for _, v := range config.SCFG.ServerInfo {
			if err := database.Init(v.Database.ToConnectStr()); err != nil {
				panic(err)
			}
			result := database.FetchPlayer(int32(v.Sid))
			for _, v2 := range result {
				js := PacketPlayer(v2)
				SnapWriteToLog(js, config.SCFG.XdlogFileDir, v.Sid)
			}
		}
	})
}

func OnlineCount() {
	gotime.AfterFunc(gotime.Second*300, func() {
		defer func() {
			if err := recover(); err != nil {
				log.Errorf("OnlineCount() error: %v\n Stack %s \n", err, debug.Stack(1, "    "))
			}
		}()
		defer OnlineCount()
		for _, v := range config.SCFG.ServerInfo {
			result := database.FetchOnline(int32(v.Sid), config.SCFG.OnlineDB.TableName)
			for _, v := range result {
				js := PacketOnline(v)
				OnlineWriteToLog(js, config.SCFG.XdlogFileDir, v.Sid)
			}
		}
	})
}
