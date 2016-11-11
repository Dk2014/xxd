package main

import (
	"core/log"
	"core/redis"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"platform_server"
	"sort"
	"strconv"
	"strings"
)

func getServerList(resp http.ResponseWriter, req *http.Request) {
	app := req.URL.Query().Get("App")
	sType := req.URL.Query().Get("Type")
	if strings.Trim(app, " ") == "" || strings.Trim(sType, " ") == "" {
		httpResponse(resp, 10001, nil, "app or type cant be empty!")
		return
	}

	iType, err := strconv.Atoi(sType)
	if err != nil {
		httpResponse(resp, 10002, nil, "type error")
		return
	}
	serverList := ServerList(getServerListFromRedis(app, uint8(iType)))
	sort.Sort(serverList)
	httpResponse(resp, 0, serverList, "")
}

func editServer(resp http.ResponseWriter, req *http.Request) {
	app := req.URL.Query().Get("App")
	if strings.Trim(app, " ") == "" {
		httpResponse(resp, 10001, nil, "app cant be empty!")
		return
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		httpResponse(resp, 10002, nil, "body error")
		return
	}

	server := platform_server.Server{}
	err = json.Unmarshal(body, &server)
	if err != nil {
		log.Infof("json parse error, body: %s", body)
		httpResponse(resp, 10003, nil, "json parse body error")
		return
	}

	c := platform_server.GetDBConn(app)
	defer c.Close()

	rediskey := platform_server.RedisKey_GameServerList(c)
	_, err = c.Do("HSET", rediskey, server.Id, body)
	if err != nil {
		httpResponse(resp, 10004, nil, "save server error")
		return
	}
	httpResponse(resp, 0, nil, "ok")
}

func isCloseType(resp http.ResponseWriter, req *http.Request) {
	app := req.URL.Query().Get("App")
	sType := req.URL.Query().Get("Type")
	if strings.Trim(app, " ") == "" || strings.Trim(sType, " ") == "" {
		httpResponse(resp, 10001, false, "app or type cant be empty!")
		return
	}

	iType, err := strconv.Atoi(sType)
	if err != nil {
		httpResponse(resp, 10002, false, "type error")
		return
	}

	c := platform_server.GetDBConn(app)
	defer c.Close()

	rediskey := platform_server.RedisKey_CloseTypeServers(uint8(iType))
	isClose, err := redis.Bool(c.Do("GET", rediskey))
	if err != nil {
		httpResponse(resp, 10004, false, "server inner error")
		return
	}
	httpResponse(resp, 0, isClose, "ok")
}

func setIsCloseType(resp http.ResponseWriter, req *http.Request) {
	app := req.URL.Query().Get("app")
	sType := req.URL.Query().Get("type")
	if strings.Trim(app, " ") == "" || strings.Trim(sType, " ") == "" {
		httpResponse(resp, 10001, nil, "app or type cant be empty!")
		return
	}

	iType, err := strconv.Atoi(sType)
	if err != nil {
		httpResponse(resp, 10002, nil, "type error")
		return
	}

	c := platform_server.GetDBConn(app)
	defer c.Close()

	isClose := req.URL.Query().Get("isCloseType")

	rediskey := platform_server.RedisKey_CloseTypeServers(uint8(iType))
	_, err = c.Do("SET", rediskey, isClose)
	if err != nil {
		httpResponse(resp, 10004, nil, "server inner error")
		return
	}

	bIsClose, err := redis.Bool(c.Do("GET", rediskey))
	if err != nil {
		httpResponse(resp, 10004, false, "server inner error")
		return
	}
	httpResponse(resp, 0, bIsClose, "ok")
}

type ServerList []platform_server.Server

func (s ServerList) Len() int           { return len(s) }
func (s ServerList) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ServerList) Less(i, j int) bool { return s[i].Id > s[j].Id }

func getServerListFromRedis(app string, iTyep uint8) []platform_server.Server {
	c := platform_server.GetDBConn(app)
	defer c.Close()
	serverList := make([]platform_server.Server, 0)

	v, err := redis.Values(c.Do("HGETALL", platform_server.RedisKey_GameServerList(c)))
	if err != nil {
		log.Info(err.Error())
		return serverList
	}

	for len(v) > 0 {
		var k string
		var b []byte
		v, err = redis.Scan(v, &k, &b)
		if err != nil {
			log.Info(err.Error())
			continue
		}

		var svr platform_server.Server
		err = json.Unmarshal(b, &svr)
		if err != nil {
			log.Info(err.Error())
			continue
		}
		if svr.Type == iTyep || svr.Type == platform_server.TYPE_MOBILE_AUDIT_WEIXIN || svr.Type == platform_server.TYPE_MOBILE_SANDBOX {
			serverList = append(serverList, svr)
		}
	}

	return serverList
}
