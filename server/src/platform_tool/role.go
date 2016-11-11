package platform_tool

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"core/redis"
	"platform_server"
)

//获取玩家在平台服的信息
func getUserRoles(flag string, app string) {
	c := platform_server.GetDBConn(app)
	defer c.Close()

	strs := strings.Split(flag, ",")
	if len(strs) != 2 {
		fmt.Println("please input right format, such as 1,openid")
		return
	}

	sType := strs[0]
	openid := strs[1]

	key := fmt.Sprintf("account_openid_type_%v_%v", openid, sType)
	bytes, err := redis.Values(c.Do("hgetall", key))

	if err != nil {
		fmt.Printf("hgetall error:%v\n", err)
	}

	fmt.Printf("roles:%s\n", bytes)
}

// 删除角色
func delUserRole(flag string, app string) {
	c := platform_server.GetDBConn(app)
	defer c.Close()

	strs := strings.Split(flag, ",")
	if len(strs) != 3 {
		fmt.Println("please input right format, such as 1,openid,sid")
		return
	}

	sType := strs[0]
	openid := strs[1]
	sid := strs[2]

	key := fmt.Sprintf("account_openid_type_%v_%v", openid, sType)
	count, err := redis.Int(c.Do("hdel", key, sid))

	if err != nil {
		fmt.Printf("hdel error:%v\n", err)
	}

	fmt.Println(count)
}

//从文件导入角色信息
func importRoles(flag string, app string) {
	arr := strings.Split(flag, ",")
	if len(arr) != 2 {
		fmt.Println("please input such as: -importRoles=roles.json,18")
		return
	}

	filePath := arr[0]
	sType := arr[1]
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	var roles map[string]platform_server.RoleInfo
	err = json.Unmarshal(b, &roles)
	if err != nil {
		panic(err)
	}

	c := platform_server.GetDBConn(app)
	defer c.Close()
	for k, v := range roles {
		key := fmt.Sprintf("account_openid_type_%s_%s", k, sType)
		sid := strconv.Itoa(int(v.Sid))
		value, err := json.Marshal(v)
		if err != nil {
			fmt.Println(err)
			continue
		}
		_, err = c.Do("HSET", key, sid, value)
		if err != nil {
			fmt.Printf("hset error:%v\n", err)
		}
	}
}

//获取游戏服在平台服注册玩家数
func showPlayerCount(flag string, app string) {
	sid, err := strconv.Atoi(flag)
	if err != nil {
		fmt.Println("wrong sid!!")
		return
	}
	err = platform_server.InitServerList()
	if err != nil {
		fmt.Println("Server list error:", err)
		return
	}
	serverlist := platform_server.ServerList(app)
	count := 0
	for _, s := range serverlist {
		if s.Id == int32(sid) {
			c := platform_server.GetDBConn(app)
			defer c.Close()

			gServers := s.GServers

			for _, v := range gServers {
				if v.HD {
					continue
				}
				gsid := v.GSId

				counterKey := platform_server.RedisKey_CounterBySidGsid(int32(sid), uint8(s.Type), gsid)
				tem_count, err := redis.Int(c.Do("GET", counterKey))
				if err != nil {
					continue
				}
				count = count + tem_count
			}
		}
	}

	fmt.Println(count)
}
