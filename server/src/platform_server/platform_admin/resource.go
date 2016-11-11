package main

import (
	"core/log"
	"core/redis"
	"net/http"
	"platform_server"
	"strconv"
	"strings"
)

func getPatches(resp http.ResponseWriter, req *http.Request) {
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

	key_patch_hash := platform_server.RedisKey_ClientPatchHash(uint8(iType))
	bytes, err := redis.Values(c.Do("HGETALL", key_patch_hash))
	if err != nil {
		httpResponse(resp, 10003, nil, "patches not found")
		return
	}

	patches := make(map[string]string, 0)
	for len(bytes) > 0 {
		var k string
		var v string
		bytes, err = redis.Scan(bytes, &k, &v)
		if err != nil {
			continue
		}
		patches[k] = v
	}

	httpResponse(resp, 0, patches, "ok")
}

func editPatch(resp http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	version := params.Get("version")
	client_version := params.Get("client_version")
	server_version := params.Get("server_version")
	patch_path := params.Get("path")
	app := params.Get("app")
	sType := params.Get("type")

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

	key_patch_hash := platform_server.RedisKey_ClientPatchHash(uint8(iType))

	if strings.Trim(version, " ") == "" {
		version = client_version + "_" + server_version
	}

	log.Info(version + "_" + patch_path)
	//如果patch为空，删除此版本patch配置
	if strings.Trim(patch_path, " ") == "" {
		_, err := c.Do("HDEL", key_patch_hash, version)
		if err != nil {
			log.Error(err.Error())
			httpResponse(resp, 10003, nil, "server inner error!")
			return
		}
		httpResponse(resp, 0, nil, "ok")
		return
	}

	_, err = c.Do("HSET", key_patch_hash, version, patch_path)
	if err != nil {
		log.Error(err.Error())
		httpResponse(resp, 10003, nil, "server inner error!")
		return
	}
	httpResponse(resp, 0, nil, "ok")
}

func getPatchUrl(resp http.ResponseWriter, req *http.Request) {
	app := req.URL.Query().Get("app")
	if strings.Trim(app, " ") == "" {
		httpResponse(resp, 10001, nil, "app or cant be empty!")
		return
	}

	c := platform_server.GetDBConn(app)
	defer c.Close()

	key_patch_url := platform_server.RedisKey_ClientPatchUrl()
	patchUrl, err := redis.String(c.Do("GET", key_patch_url))
	if err != nil {
		log.Error(err.Error())
		httpResponse(resp, 0, "", "not found")
		return
	}
	httpResponse(resp, 0, patchUrl, "ok")
}

func editPatchUrl(resp http.ResponseWriter, req *http.Request) {
	app := req.URL.Query().Get("app")
	patchUrl := req.URL.Query().Get("patchurl")
	if strings.Trim(app, " ") == "" || strings.Trim(patchUrl, " ") == "" {
		httpResponse(resp, 10001, nil, "app or cant be empty!")
		return
	}

	c := platform_server.GetDBConn(app)
	defer c.Close()

	key_patch_url := platform_server.RedisKey_ClientPatchUrl()
	_, err := c.Do("SET", key_patch_url, patchUrl)
	if err != nil {
		log.Error(err.Error())
		httpResponse(resp, 10003, nil, "server inner error!")
		return
	}
	httpResponse(resp, 0, nil, "ok")
}

func getTowns(resp http.ResponseWriter, req *http.Request) {
	app := req.URL.Query().Get("app")

	c := platform_server.GetDBConn(app)
	defer c.Close()

	townsMap := make(map[string]map[string]string)
	versions, err := redis.Strings(c.Do("SMEMBERS", "total_resource_file_version"))
	if err != nil {
		httpResponse(resp, 0, townsMap, "ok")
		return
	}

	for _, v := range versions {
		version, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		hash_key := platform_server.RedisKey_TotalResourceHash(int32(version))
		bytes, err := redis.Values(c.Do("HGETALL", hash_key))
		if err != nil {
			continue
		}

		towns := make(map[string]string, 0)
		for len(bytes) > 0 {
			var k string
			var v string
			bytes, err = redis.Scan(bytes, &k, &v)
			if err != nil {
				continue
			}
			towns[k] = v
		}

		townsMap[strconv.Itoa(version)] = towns
	}

	httpResponse(resp, 0, townsMap, "ok")
}

func editTown(resp http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	version := params.Get("version")
	name := params.Get("name")
	path := params.Get("path")
	size := params.Get("size")
	app := params.Get("app")

	if strings.Trim(app, " ") == "" || strings.Trim(version, " ") == "" {
		httpResponse(resp, 10001, nil, "app or version cant be empty!")
		return
	}

	iVersion, err := strconv.Atoi(version)
	if err != nil {
		httpResponse(resp, 10002, nil, "version error")
		return
	}

	c := platform_server.GetDBConn(app)
	defer c.Close()

	hash_key := platform_server.RedisKey_TotalResourceHash(int32(iVersion))
	//如果resource为空，删除此resource配置
	if strings.Trim(path, " ") == "" {
		_, err := c.Do("HDEL", hash_key, name)
		if err != nil {
			log.Error(err.Error())
			httpResponse(resp, 10003, nil, "server inner error!")
			return
		}
		httpResponse(resp, 0, nil, "ok")
		return
	}

	_, err = c.Do("HSET", hash_key, name, path+","+size)
	if err != nil {
		httpResponse(resp, 10003, nil, "server inner error!")
		return
	}

	_, err = c.Do("SADD", "total_resource_file_version", iVersion)
	if err != nil {
		log.Error(err.Error())
	}

	httpResponse(resp, 0, nil, "ok")
}

func getTownUrl(resp http.ResponseWriter, req *http.Request) {
	app := req.URL.Query().Get("app")
	if strings.Trim(app, " ") == "" {
		httpResponse(resp, 10001, nil, "app can't be empty!")
		return
	}

	c := platform_server.GetDBConn(app)
	defer c.Close()

	key_town_url := platform_server.RedisKey_ClientResourceUrl()
	townUrl, err := redis.String(c.Do("GET", key_town_url))
	if err != nil {
		log.Error(err.Error())
		httpResponse(resp, 0, "", "not found")
		return
	}
	httpResponse(resp, 0, townUrl, "ok")
}

func editTownUrl(resp http.ResponseWriter, req *http.Request) {
	app := req.URL.Query().Get("app")
	townUrl := req.URL.Query().Get("townurl")
	if strings.Trim(app, " ") == "" || strings.Trim(townUrl, " ") == "" {
		httpResponse(resp, 10001, nil, "app or cant be empty!")
		return
	}

	c := platform_server.GetDBConn(app)
	defer c.Close()

	key_town_url := platform_server.RedisKey_ClientResourceUrl()
	_, err := c.Do("SET", key_town_url, townUrl)
	if err != nil {
		log.Error(err.Error())
		httpResponse(resp, 10003, nil, "server inner error!")
		return
	}
	httpResponse(resp, 0, nil, "ok")
}
