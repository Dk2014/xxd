package main

import (
	"core/redis"
	"net/http"
	"strconv"
	"strings"

	"platform_server"
)

func getVersionInfo(resp http.ResponseWriter, req *http.Request) {
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

	min_key := platform_server.RedisKey_ClientMinVersionByType(uint8(iType))
	audit_key := platform_server.RedisKey_ClientAuditVersionByType(uint8(iType))
	url_key := platform_server.RedisKey_ClientUpgradeUrl(uint8(iType))

	minVersion, err := redis.String(c.Do("GET", min_key))
	auditVersion, err := redis.String(c.Do("GET", audit_key))
	upgradeUrl, err := redis.String(c.Do("GET", url_key))

	if err != nil {
		httpResponse(resp, 10004, nil, "server inner error")
		return
	}

	versinInfo := make(map[string]string)
	versinInfo["MinVersion"] = minVersion
	versinInfo["AuditVersion"] = auditVersion
	versinInfo["UpgradeUrl"] = upgradeUrl

	httpResponse(resp, 0, versinInfo, "ok")
}

func setVersionInfo(resp http.ResponseWriter, req *http.Request) {
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

	minVersion, err := strconv.Atoi(req.URL.Query().Get("minVersion"))
	auditVersion, err := strconv.Atoi(req.URL.Query().Get("auditVersion"))
	upgradeUrl := req.URL.Query().Get("upgradeUrl")
	if err != nil || strings.Trim(upgradeUrl, " ") == "" {
		httpResponse(resp, 10003, nil, "wrong minVersion or auditVersion or upgradeUrl.")
		return
	}

	min_key := platform_server.RedisKey_ClientMinVersionByType(uint8(iType))
	audit_key := platform_server.RedisKey_ClientAuditVersionByType(uint8(iType))
	url_key := platform_server.RedisKey_ClientUpgradeUrl(uint8(iType))
	_, err = c.Do("SET", min_key, minVersion)
	_, err = c.Do("SET", audit_key, auditVersion)
	_, err = c.Do("SET", url_key, upgradeUrl)

	if err != nil {
		httpResponse(resp, 10004, nil, "server inner error")
		return
	}

	httpResponse(resp, 0, nil, "ok")
}

//灰度服
func getDarkInfo(resp http.ResponseWriter, req *http.Request) {
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

	key := platform_server.RedisKey_DarkLaunch(uint8(iType))
	val, err := redis.String(c.Do("GET", key))
	if err != nil {
		httpResponse(resp, 10004, nil, "server inner error")
		return
	}

	arrs := strings.Split(val, ",")
	if len(arrs) != 4 {
		httpResponse(resp, 10004, nil, "server inner error")
		return
	}
	version, err := strconv.Atoi(arrs[1])
	isCompatible, err := strconv.Atoi(arrs[2])
	upgradeUrl := arrs[3]

	darkInfo := make(map[string]interface{})
	darkInfo["Sids"] = arrs[0]
	darkInfo["IsCompatible"] = isCompatible
	darkInfo["Version"] = version
	darkInfo["UpgradeUrl"] = upgradeUrl

	httpResponse(resp, 0, darkInfo, "ok")
}

func setDarkInfo(resp http.ResponseWriter, req *http.Request) {
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

	key := platform_server.RedisKey_DarkLaunch(uint8(iType))
	sids := req.URL.Query().Get("sids")
	version := req.URL.Query().Get("version")
	isCompatible := req.URL.Query().Get("isCompatible")
	upgradeUrl := req.URL.Query().Get("upgradeUrl")

	_, err = redis.String(c.Do("SET", key, sids+","+version+","+isCompatible+","+upgradeUrl))
	if err != nil {
		httpResponse(resp, 10004, nil, "server inner error")
		return
	}

	httpResponse(resp, 0, nil, "ok")
}

func delDarkInfo(resp http.ResponseWriter, req *http.Request) {
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

	key := platform_server.RedisKey_DarkLaunch(uint8(iType))
	_, err = c.Do("del", key)
	if err != nil {
		httpResponse(resp, 10004, nil, "server inner error")
		return
	}

	httpResponse(resp, 0, nil, "ok")
}
