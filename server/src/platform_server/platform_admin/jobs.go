package main

import (
	"core/log"
	"core/redis"
	"encoding/json"
	"platform_server"
	"runtime/debug"
	"strconv"
	"sync"
	"time"
)

const (
	OPEN_SERVER_TICKER_TIME    = time.Minute * 10
	OPEN_SERVER_TICKER_SECONDS = 600
)

var (
	_server_list_lock sync.RWMutex
)

func openNewServer(app, sid string, delayTime int) error {
	defer func() {
		if r := recover(); r != nil {
			log.Errorf("Recovered in %v: %v", r, string(debug.Stack()))
		}
	}()

	time.Sleep(time.Duration(delayTime) * time.Second)

	c := platform_server.GetDBConn(app)
	defer c.Close()

	iSid, err := strconv.Atoi(sid)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	lastSid := strconv.Itoa(iSid - 1)

	rediskey := platform_server.RedisKey_GameServerList(c)
	bNewServer, err := redis.Bytes(c.Do("HGET", rediskey, sid))
	bLastServer, err := redis.Bytes(c.Do("HGET", rediskey, lastSid))
	if err != nil {
		log.Error(err.Error())
		return err
	}

	newServer := new(platform_server.Server)
	lastServer := new(platform_server.Server)
	err = json.Unmarshal(bNewServer, newServer)
	err = json.Unmarshal(bLastServer, lastServer)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	newServer.IsNew = true
	newServer.IsHot = true

	lastServer.IsNew = false
	lastServer.IsHot = false
	lastServer.Status = 2

	bNewServer, err = json.Marshal(newServer)
	bLastServer, err = json.Marshal(lastServer)
	_, err = redis.Bytes(c.Do("HMSET", rediskey, sid, bNewServer, lastSid, bLastServer))

	if err != nil {
		log.Error(err.Error())
	}
	return err
}

func startJobs() {
	go startOpenServerJobs()
}

func startOpenServerJobs() {
	defer func() {
		if r := recover(); r != nil {
			log.Errorf("Recovered in %v: %v", r, string(debug.Stack()))
		}
	}()

	ticker := time.NewTicker(OPEN_SERVER_TICKER_TIME)
	for {
		select {
		case <-ticker.C:
			doOpenServerJobs()
		}
	}
}

func doOpenServerJobs() {
	serverLists := getAllServerList()
	for app, serverlist := range serverLists {
		for sid, server := range serverlist {
			openTime := server.OpenTime
			curTime := time.Now().Unix()
			if openTime > curTime && openTime <= curTime+OPEN_SERVER_TICKER_SECONDS {
				go openNewServer(app, sid, int(openTime-curTime))
			}
		}
	}
}

func getAllServerList() map[string]map[string]platform_server.Server {
	serverLists := make(map[string]map[string]platform_server.Server)
	dbmap := platform_server.GetDBMap()

	for app, _ := range dbmap {
		c := platform_server.GetDBConn(app)
		defer c.Close()

		v, err := redis.Values(c.Do("HGETALL", platform_server.RedisKey_GameServerList(c)))
		if err != nil {
			log.Error(err.Error())
			break
		}
		var list_map = make(map[string]platform_server.Server)

		for len(v) > 0 {
			var k string
			var b []byte
			v, err = redis.Scan(v, &k, &b)
			if err != nil {
				log.Error(err.Error())
				break
			}

			var svr platform_server.Server
			err = json.Unmarshal(b, &svr)
			if err != nil {
				log.Error(err.Error())
				break
			}
			list_map[strconv.Itoa(int(svr.Id))] = svr
		}

		serverLists[app] = list_map
	}
	return serverLists
}
