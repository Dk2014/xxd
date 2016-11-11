package platform_tool

import (
	"bufio"
	"core/redis"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"platform_server"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func exportServerListByRev(in string, app string) {
	apps := make([]string, 0)
	if strings.Contains(app, ":") {
		apps = strings.Split(app, ":")
	} else {
		apps = append(apps, app)
	}

	for _, app := range apps {
		path_rev := strings.Split(in, ",")
		var rev int = -1
		path := path_rev[0]
		if len(path_rev) == 2 {
			rev, _ = strconv.Atoi(path_rev[1])
		}

		c := platform_server.GetDBConn(app)
		defer c.Close()

		if rev < 0 {
			rev, _ = redis.Int(c.Do("GET", platform_server.RedisKey_GameServerListCurrentRev()))
		}

		export_key := platform_server.RedisKey_GameServerListByRevision(rev)

		v, err := redis.Values(c.Do("HGETALL", export_key))
		if err != nil {
			fmt.Println("fail to get game server list")
		}

		var list = []platform_server.Server{}

		for len(v) > 0 {
			var k string
			var b []byte
			v, err = redis.Scan(v, &k, &b)
			if err != nil {
				fmt.Printf("fail to parse game server list: %v", v)
			}

			var svr platform_server.Server
			// validate serverlist is proper json format
			err = json.Unmarshal(b, &svr)
			if err != nil {
				fmt.Printf("fail to parse game server info: %v", string(b))
				continue
			}

			list = append(list, svr)
		}

		// write marshaled json of rsp
		b, err := json.MarshalIndent(list, "", "	")
		if err != nil {
			fmt.Printf("server list marshal error: %v\n", v)
			return
		}

		fmt.Println(string(b))

		err = ioutil.WriteFile(path, b, 0777)
		if err != nil {
			fmt.Println("Write file error:", err)
		}
	}
}

// 检查每个server id是否唯一,每个server是否都有game server,每个game server的ip和port是否非空,
// 每个game server id是否唯一,每个game server的ip:port是否唯一,每个server是否有且仅有一个互动的
// game server返回值为0则表示通过检查,1则表示检查不通过
// 此函数已在commit函数中调用以确保最终提交的版本符合要求
func verify(local_list []platform_server.Server) int {
	var flag int
	var serverId []int32
	var gameserverId []int32
	var ipAndPort []string
	for _, value := range local_list {
		serverId = append(serverId, value.Id)
		for _, value1 := range value.GServers {
			gameserverId = append(gameserverId, value1.GSId)
			if value1.Ip == "" {
				fmt.Println("The game server", value1, "is with empty ip.")
				flag = 1
			}
			if value1.Port == "" {
				fmt.Println("The game server", value1, "is with empty port.")
				flag = 1
			}

			ipAndPort = append(ipAndPort, value1.Ip+":"+value1.Port)
		}
		flag = flag | duplicate_check(serverId, "server id")
	}
	return flag
}

//内部调用，重复值检验
func duplicate_check(check []int32, mtype string) int {
	var flag int
	for i := 0; i < len(check); i++ {
		for j := i; j < len(check); j++ {
			if i != j && check[j] == check[i] {
				fmt.Println("These is duplicate", mtype, ".Please check and modify it.")
				flag = 1
				break
			}
		}
	}
	return flag
}

func importServerList(local_list []platform_server.Server, app string) (int, error) {
	if verify(local_list) == 1 {
		fmt.Println("Cannot pass the verify.Please check.")
		return 0, errors.New("Verify check fail.")
	}
	c := platform_server.GetDBConn(app)
	defer c.Close()

	rev := platform_server.NextServerListRevsion(app)
	rediskey := platform_server.RedisKey_GameServerListByRevision(rev)

	// importing
	for _, v := range local_list {
		b, err := json.Marshal(v)
		if err != nil {
			fmt.Println("Svrlists marshal error:", err)
			return -1, err
		}

		c.Do("HSET", rediskey, v.Id, b)
	}

	return rev, nil
}

func importServerListByFile(path string, app string) {
	apps := make([]string, 0)
	if strings.Contains(app, ":") {
		apps = strings.Split(app, ":")
	} else {
		apps = append(apps, app)
	}

	for _, app := range apps {
		file, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Println("Error when read file:", err)
			return
		}

		var local_list []platform_server.Server
		err = json.Unmarshal(file, &local_list)
		if err != nil {
			fmt.Println("Error when unmarshaling: ", err)
			return
		}

		rev, err := importServerList(local_list, app)

		if err != nil {
			fmt.Println("Error when import server list:", err)
		} else {
			fmt.Println("import server list successfully!")
			for {
				fmt.Printf("Are you sure to commit revision:%v? y/n\n", rev)
				answer := getInput()
				if answer == "y" || answer == "Y" {
					commitServerListByRevsion(rev, app)
					break
				} else if answer == "n" || answer == "N" {
					break
				} else {
					fmt.Println("Wrong input！Please input the right format.")
				}
			}
		}
	}
}

func addServerListByFile(path string, app string) {
	apps := make([]string, 0)
	if strings.Contains(app, ":") {
		apps = strings.Split(app, ":")
	} else {
		apps = append(apps, app)
	}

	for _, app := range apps {
		file, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Println("Error when read file:", err)
			return
		}

		var local_list []platform_server.Server
		err = json.Unmarshal(file, &local_list)
		if err != nil {
			fmt.Println("Error when unmarshaling: ", err)
			return
		}

		rev, err := addServerList(local_list, app)

		if err != nil {
			fmt.Println("Error when import server list:", err)
		} else {
			fmt.Println("import server list successfully!")
			for {
				fmt.Printf("Are you sure to commit revision:%v? y/n\n", rev)
				answer := getInput()
				if answer == "y" || answer == "Y" {
					commitServerListByRevsion(rev, app)
					break
				} else if answer == "n" || answer == "N" {
					break
				} else {
					fmt.Println("Wrong input！Please input the right format.")
				}
			}
		}
	}
}

func addServerList(local_list []platform_server.Server, app string) (int, error) {
	if verify(local_list) == 1 {
		fmt.Println("Cannot pass the verify.Please check.")
		return 0, errors.New("Verify check fail.")
	}
	c := platform_server.GetDBConn(app)
	defer c.Close()

	currentRev, err := redis.Int(c.Do("GET", platform_server.RedisKey_GameServerListCurrentRev()))
	var currentList []platform_server.Server
	var currentIds []int32

	if err == nil {
		currentkey := platform_server.RedisKey_GameServerListByRevision(currentRev)
		v, err := redis.Values(c.Do("HGETALL", currentkey))
		if err != nil {
			fmt.Println("fail to get game server list")
		}

		for len(v) > 0 {
			var k string
			var b []byte
			v, err = redis.Scan(v, &k, &b)
			if err != nil {
				fmt.Printf("fail to parse game server list: %v", v)
			}

			var svr platform_server.Server
			// validate serverlist is proper json format
			err = json.Unmarshal(b, &svr)
			if err != nil {
				fmt.Printf("fail to parse game server info: %v", string(b))
				continue
			}

			currentList = append(currentList, svr)
			currentIds = append(currentIds, svr.Id)
		}
	}

	nextRev := platform_server.NextServerListRevsion(app)
	nextKey := platform_server.RedisKey_GameServerListByRevision(nextRev)
	//写入新加入的服务器配置
	totalAdd := 0
	for _, server := range local_list {
		if arrayContains(currentIds, server.Id) {
			continue
		}

		b, err := json.Marshal(server)
		if err != nil {
			fmt.Println("Svrlists marshal error:", err)
			return -1, err
		}
		c.Do("HSET", nextKey, server.Id, b)
		totalAdd++
	}

	fmt.Printf("add new server count: %d\n", totalAdd)

	//写入原有的服务器配置
	for _, server := range currentList {
		b, err := json.Marshal(server)
		if err != nil {
			fmt.Println("Svrlists marshal error:", err)
			return -1, err
		}
		c.Do("HSET", nextKey, server.Id, b)
	}
	return nextRev, err
}

func arrayContains(arr []int32, val int32) bool {
	if arr == nil || len(arr) == 0 {
		return false
	}

	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

func commitServerListByRevsion(rev int, app string) {
	apps := make([]string, 0)
	if strings.Contains(app, ":") {
		apps = strings.Split(app, ":")
	} else {
		apps = append(apps, app)
	}

	for _, app := range apps {
		c := platform_server.GetDBConn(app)
		defer c.Close()

		exist, err := redis.Bool(c.Do("EXISTS", platform_server.RedisKey_GameServerListByRevision(rev)))
		if err != nil {
			fmt.Println("Revsion not exist", platform_server.RedisKey_GameServerListByRevision(rev), ":", err)
			return
		}
		if !exist {
			fmt.Println("Revsion not exist")
			return
		}

		old_rev, err := redis.Int(c.Do("GET", platform_server.RedisKey_GameServerListCurrentRev()))
		if err != nil {
			fmt.Println("Error on geting current revision")
			old_rev = -1
		}

		c.Do("SET", platform_server.RedisKey_GameServerListCurrentRev(), rev)

		new_rev, err := redis.Int(c.Do("GET", platform_server.RedisKey_GameServerListCurrentRev()))
		if err != nil {
			fmt.Println("Error on geting current revision")
			return
		}

		fmt.Println("Server list has been set to revision:", new_rev, ". (previous revsion:", old_rev, ")")

		return
	}
}

//关闭服务器(即将状态为开放(非0)的服务器的状态置为维护(0))
func closeServer(id, app string) {
	tmp, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Please input the right id of the server you want to close!")
		return
	}
	err = platform_server.InitServerList()
	if err != nil {
		fmt.Println("Server list error:", err)
		return
	}
	serverlist := platform_server.ServerList(app)
	var listkey string
	var server platform_server.Server
	var flag int
	for key, value := range serverlist {
		if value.Id == int32(tmp) {
			if value.Status == 0 {
				fmt.Println("Server with id", id, "already closed.You can't close a closed server.")
				return
			}
			flag = 1
			listkey = key
			server = value
			server.Status = 0
			break
		}
	}
	if flag == 0 {
		fmt.Println("Cannot find server with id", id, ",please check if you have input the right id.")
		return
	}
	delete(serverlist, listkey)
	serverlist[listkey] = server
	var local_list []platform_server.Server
	for _, value := range serverlist {
		local_list = append(local_list, value)
	}

	rev, err := importServerList(local_list, app)
	if err != nil {
		fmt.Println("Error when closing server:", err)
	} else {
		fmt.Println("Server with id", id, "closed successfully!")
		for {
			fmt.Printf("Are you sure to commit revision:%v? y/n\n", rev)
			answer := getInput()
			if answer == "y" || answer == "Y" {
				commitServerListByRevsion(rev, app)
				break
			} else if answer == "n" || answer == "N" {
				break
			} else {
				fmt.Println("Wrong input！Please input the right format.")
			}
		}
	}
}

//开放服务器(即将状态为关闭(0)的服务器的状态更改为空闲(1))
func openServer(id, app string) {
	tmp, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Please input the right id of the server you want to open!")
		return
	}
	err = platform_server.InitServerList()
	if err != nil {
		fmt.Println("Server list error:", err)
		return
	}
	serverlist := platform_server.ServerList(app)
	var listkey string
	var server platform_server.Server
	var flag int
	for key, value := range serverlist {
		if value.Id == int32(tmp) {
			if value.Status != 0 {
				fmt.Println("Server with id", id, "already opened.You can't open a opened server.")
				return
			}
			flag = 1
			listkey = key
			server = value
			server.Status = 1
			break
		}
	}
	if flag == 0 {
		fmt.Println("Cannot find server with id", id, ",please check if you have input the right id.")
		return
	}
	delete(serverlist, listkey)
	serverlist[listkey] = server
	var local_list []platform_server.Server
	for _, value := range serverlist {
		local_list = append(local_list, value)
	}

	rev, err := importServerList(local_list, app)
	if err != nil {
		fmt.Println("Error when opening server:", err)
	} else {
		fmt.Println("Server with id", id, "open successfully!")
		for {
			fmt.Printf("Are you sure to commit revision:%v? y/n\n", rev)
			answer := getInput()
			if answer == "y" || answer == "Y" {
				commitServerListByRevsion(rev, app)
				break
			} else if answer == "n" || answer == "N" {
				break
			} else {
				fmt.Println("Wrong input！Please input the right format.")
			}
		}
	}
}

//设置server的状态
func setServerStatus(id string, status int8, statusMsg string, app string) {
	tmp, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Please input the right id of the server you want to set status!")
		return
	}
	err = platform_server.InitServerList()
	if err != nil {
		fmt.Println("Server list error:", err)
		return
	}
	serverlist := platform_server.ServerList(app)
	var listkey string
	var server platform_server.Server
	var flag int
	for key, value := range serverlist {
		if value.Id == int32(tmp) {
			if value.Status == status && server.StatusMsg == statusMsg {
				fmt.Println("Server with id", id, "'s status is already ", status, ".You can't set a server status if it is already in such status.")
				return
			}
			flag = 1
			listkey = key
			server = value
			server.Status = status
			server.StatusMsg = statusMsg
			break
		}
	}
	if flag == 0 {
		fmt.Println("Cannot find server with id", id, ",please check if you have input the right id.")
		return
	}
	delete(serverlist, listkey)
	serverlist[listkey] = server
	var local_list []platform_server.Server
	for _, value := range serverlist {
		local_list = append(local_list, value)
	}
	rev, err := importServerList(local_list, app)
	if err != nil {
		fmt.Println("Error when setting server status:", err)
	} else {
		fmt.Println("Set server with id", id, "'s status successfully!")
		for {
			fmt.Printf("Are you sure to commit revision:%v? y/n\n", rev)
			answer := getInput()
			if answer == "y" || answer == "Y" {
				commitServerListByRevsion(rev, app)
				break
			} else if answer == "n" || answer == "N" {
				break
			} else {
				fmt.Println("Wrong input！Please input the right format.")
			}
		}
	}
}

//设置server为推荐服(即将IsHot属性原本不为true的设置为true)
func operateHotServer(id string, isHot bool, app string) {
	tmp, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Please input the right id of the server you want to set hot!")
		return
	}
	err = platform_server.InitServerList()
	if err != nil {
		fmt.Println("Server list error:", err)
		return
	}
	serverlist := platform_server.ServerList(app)
	var listkey string
	var server platform_server.Server
	var flag int
	for key, value := range serverlist {
		if value.Id == int32(tmp) {
			if value.IsHot == isHot {
				fmt.Println("no need to set")
				return
			}
			flag = 1
			listkey = key
			server = value
			server.IsHot = isHot
			break
		}
	}
	if flag == 0 {
		fmt.Println("Cannot find server with id", id, ",please check if you have input the right id.")
		return
	}
	delete(serverlist, listkey)
	serverlist[listkey] = server
	var local_list []platform_server.Server
	for _, value := range serverlist {
		local_list = append(local_list, value)
	}
	rev, err := importServerList(local_list, app)
	if err != nil {
		fmt.Println("Error when setting server hot:", err)
	} else {
		for {
			fmt.Printf("Are you sure to commit revision:%v? y/n\n", rev)
			answer := getInput()
			if answer == "y" || answer == "Y" {
				commitServerListByRevsion(rev, app)
				break
			} else if answer == "n" || answer == "N" {
				break
			} else {
				fmt.Println("Wrong input！Please input the right format.")
			}
		}
	}
}

//设置server为新服(即将IsNew属性原本不为true的设置为true)
func operateNewServer(id string, isNew bool, app string) {
	tmp, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Please input the right id of the server you want to set new!")
		return
	}
	err = platform_server.InitServerList()
	if err != nil {
		fmt.Println("Server list error:", err)
		return
	}
	serverlist := platform_server.ServerList(app)
	var listkey string
	var server platform_server.Server
	var flag int
	for key, value := range serverlist {
		if value.Id == int32(tmp) {
			if value.IsNew == isNew {
				fmt.Println("no need to set")
				return
			}
			flag = 1
			listkey = key
			server = value
			server.IsNew = isNew
			break
		}
	}
	if flag == 0 {
		fmt.Println("Cannot find server with id", id, ",please check if you have input the right id.")
		return
	}
	delete(serverlist, listkey)
	serverlist[listkey] = server
	var local_list []platform_server.Server
	for _, value := range serverlist {
		local_list = append(local_list, value)
	}

	rev, err := importServerList(local_list, app)
	if err != nil {
		fmt.Println("Error when setting server new:", err)
	} else {
		for {
			fmt.Printf("Are you sure to commit revision:%v? y/n\n", rev)
			answer := getInput()
			if answer == "y" || answer == "Y" {
				commitServerListByRevsion(rev, app)
				break
			} else if answer == "n" || answer == "N" {
				break
			} else {
				fmt.Println("Wrong input！Please input the right format.")
			}
		}
	}
}

//更改服务器名
func modifyServerName(id string, app string) {
	tmp, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Please input the right id of the server you want to modify name!")
		return
	}
	err = platform_server.InitServerList()
	if err != nil {
		fmt.Println("Server list error:", err)
		return
	}
	serverlist := platform_server.ServerList(app)
	var listkey string
	var server platform_server.Server
	var flag int
	for key, value := range serverlist {
		if value.Id == int32(tmp) {
			fmt.Println("The server's name is", server.Name, ".Please input a new name for it:")
			tmp := getInput()
			flag = 1
			listkey = key
			server = value
			server.Name = tmp
			break
		}
	}
	if flag == 0 {
		fmt.Println("Cannot find server with id", id, ",please check if you have input the right id.")
		return
	}
	delete(serverlist, listkey)
	serverlist[listkey] = server
	var local_list []platform_server.Server
	for _, value := range serverlist {
		local_list = append(local_list, value)
	}

	rev, err := importServerList(local_list, app)
	if err != nil {
		fmt.Println("Error when modifying server name:", err)
	} else {
		fmt.Println("Set server with id", id, "name modefied successfully!")
		for {
			fmt.Printf("Are you sure to commit revision:%v? y/n\n", rev)
			answer := getInput()
			if answer == "y" || answer == "Y" {
				commitServerListByRevsion(rev, app)
				break
			} else if answer == "n" || answer == "N" {
				break
			} else {
				fmt.Println("Wrong input！Please input the right format.")
			}
		}
	}
}

//读取当前服务器列表
func showList(app string) {
	apps := make([]string, 0)
	if strings.Contains(app, ":") {
		apps = strings.Split(app, ":")
	} else {
		apps = append(apps, app)
	}

	for _, app := range apps {
		err := platform_server.InitServerList()
		if err != nil {
			fmt.Println("Server list error:", err)
			return
		}
		serverlist := platform_server.ServerList(app)

		var list = []platform_server.Server{}
		for _, value := range serverlist {
			list = append(list, value)
		}

		b, err := json.MarshalIndent(list, "", "	")
		if err != nil {
			fmt.Printf("server list marshal error: %v\n", list)
			return
		}

		fmt.Printf("%s", b)
	}
}

func addServer(server platform_server.Server, app string) []platform_server.Server {
	err := platform_server.InitServerList()
	if err != nil {
		fmt.Println("Server list error:", err)
		return nil
	}
	serverlist := platform_server.ServerList(app)
	for _, value := range serverlist {
		fmt.Println(value)
	}
	var local_list []platform_server.Server
	for _, value := range serverlist {
		local_list = append(local_list, value)
	}
	local_list = append(local_list, server)
	return local_list
}

// 删除逻辑区服
func delServer(sid string, app string) {
	err := platform_server.InitServerList()
	if err != nil {
		fmt.Println("Server list error:", err)
		return
	}
	serverlist := platform_server.ServerList(app)

	i, err := strconv.Atoi(sid)
	if err != nil {
		fmt.Println("Please input the right id!\nErr:", err)
	} else {
		var flag int
		for key, value := range serverlist {
			if value.Id == int32(i) {
				flag = 1
				delete(serverlist, key)
				var local_list []platform_server.Server
				for _, value = range serverlist {
					local_list = append(local_list, value)
				}

				rev, err := importServerList(local_list, app)
				if err != nil {
					fmt.Println("Error when delete server:", err)
				} else {
					fmt.Println("Server delete successfully!")
					for {
						fmt.Printf("Are you sure to commit revision:%v? y/n\n", rev)
						answer := getInput()
						if answer == "y" || answer == "Y" {
							commitServerListByRevsion(rev, app)
							break
						} else if answer == "n" || answer == "N" {
							break
						} else {
							fmt.Println("Wrong input！Please input the right format.")
						}
					}
				}
				break
			}
		}
		if flag == 0 {
			fmt.Println("Cannot find the server with id:", i, ",please check if you've input the right server id.")
		}
	}
}

//获取键盘输入
func getInput() string {
	stdin := bufio.NewReader(os.Stdin)
	input, _ := stdin.ReadString('\n')
	tmp := input[0 : len(input)-1]
	return tmp
}

// 添加新逻辑区服配置（human readable）
func serverInfoAdd(app string) {
	apps := make([]string, 0)
	if strings.Contains(app, ":") {
		apps = strings.Split(app, ":")
	} else {
		apps = append(apps, app)
	}

	for _, app := range apps {
		server := new(platform_server.Server)
		for {
			fmt.Println("Please input the id of server you want to add:")
			tmp := getInput()
			i, err := strconv.Atoi(tmp)
			if err == nil {
				server.Id = int32(i)
				break
			}
			fmt.Println("Please input the right id!\nErr:", err)
		}
		for {
			fmt.Println("Please input the Type of server you want to add:")
			tmp := getInput()
			i, err := strconv.Atoi(tmp)
			if err != nil {
				fmt.Println("Please input the right type!Err:", err)
			} else {
				if platform_server.ValidateServerType(uint8(i)) {
					server.Type = uint8(i)
					break
				}
			}
			fmt.Println("Please input the right type!")
		}
		fmt.Println("Please input the name of server you want to add:")
		server.Name = getInput()
		server.Status = 0
		server.IsNew = true
		for {
			fmt.Println("Is recommanded server?y/n")
			tmp := getInput()
			if tmp == "y" || tmp == "Y" {
				server.IsHot = true
				break
			} else if tmp == "n" || tmp == "N" {
				server.IsHot = false
				break
			} else {
				fmt.Println("Wrong input！Please input the right format.")
			}
		}
		for {
			fmt.Println("Please input the opentime(format like:2014-06-12 08:00:00)")
			tmp := getInput()
			regexpTest := "^\\d{4}-\\d{1,2}-\\d{1,2}\\s\\d{1,2}:\\d{1,2}:\\d{1,2}$"
			pass, err := regexp.MatchString(regexpTest, tmp)
			if err != nil {
				fmt.Println("Regexp match fail!")
				return
			}
			if !pass {
				fmt.Println("Please input the right format of opentime")
				continue
			}
			temp := regexp.MustCompile(`[\d]+`)
			mytime := temp.FindAllString(tmp, -1)
			var myTime [6]int
			for i := 0; i < 6; i++ {
				var err error
				myTime[i], err = strconv.Atoi(mytime[i])
				if err != nil {
					fmt.Println("Error when get time:", err)
					return
				}
			}
			if myTime[1] == 0 || myTime[1] > 12 {
				fmt.Println("Please input the right month.")
				continue
			}
			t := time.Date(myTime[0], time.Month(myTime[1]), myTime[2], myTime[3], myTime[4], myTime[5], 0, time.Local)
			fmt.Println("The opentime is:", t.Format("2006-01-02 15:04:05.999999 -07:00"), ".Press y to comfirm, others to cancel.")
			tmp = getInput()
			if tmp == "Y" || tmp == "y" {
				server.OpenTime = t.Unix()
				break
			}
		}
		fmt.Println("Please add a GameServer for this new server:")
		gameServer := gameServerSet()
		server.GServers = append(server.GServers, gameServer)
		fmt.Println("The server you want to add is:", server, "\npress y to comfirm and others to cancel")
		tmp := getInput()
		if tmp == "Y" || tmp == "y" {
			local_list := addServer(*server, app)

			rev, err := importServerList(local_list, app)
			if err != nil {
				fmt.Println("Error when adding server:", err)
			} else {
				fmt.Println("Server added successfully!")
				for {
					fmt.Printf("Are you sure to commit revision:%v? y/n\n", rev)
					answer := getInput()
					if answer == "y" || answer == "Y" {
						commitServerListByRevsion(rev, app)
						break
					} else if answer == "n" || answer == "N" {
						break
					} else {
						fmt.Println("Wrong input！Please input the right format.")
					}
				}
			}
		} else {
			fmt.Println("Server add canceled.No server added.")
		}
	}
}

// 添加 game_server
func addGameServer(app string) {
	apps := make([]string, 0)
	if strings.Contains(app, ":") {
		apps = strings.Split(app, ":")
	} else {
		apps = append(apps, app)
	}

	for _, app := range apps {
		err := platform_server.InitServerList()
		if err != nil {
			fmt.Println("Server list error:", err)
			return
		}
		serverlist := platform_server.ServerList(app)
		for _, value := range serverlist {
			fmt.Println(value)
		}
		listkey := getListKey(serverlist)
		server := serverlist[listkey]
		gameServer := gameServerSet()
		server.GServers = append(server.GServers, gameServer)
		delete(serverlist, listkey)
		serverlist[listkey] = server
		fmt.Println("Game server add successfully.Plaese commit to let it work!")
		var local_list []platform_server.Server
		for _, value := range serverlist {
			local_list = append(local_list, value)
		}

		rev, err := importServerList(local_list, app)
		if err != nil {
			fmt.Println("Error when adding gameserver:", err)
		} else {
			fmt.Println("Gameserver added successfully!")
			for {
				fmt.Printf("Are you sure to commit revision:%v? y/n\n", rev)
				answer := getInput()
				if answer == "y" || answer == "Y" {
					commitServerListByRevsion(rev, app)
					break
				} else if answer == "n" || answer == "N" {
					break
				} else {
					fmt.Println("Wrong input！Please input the right format.")
				}
			}
		}
	}
}

//获取所选择服务器的key
func getListKey(serverlist map[string]platform_server.Server) string {
	var listkey string
	for {
		fmt.Println("Please input the id of server that you want to operate.")
		tmp := getInput()
		id, err := strconv.Atoi(tmp)
		if err != nil {
			fmt.Println("Please input the right id.Err:", err)
		} else {
			flag := 0
			for key, value := range serverlist {
				if value.Id == int32(id) {
					listkey = key
					flag = 1
					fmt.Println("Server with id", id, "successfully chosen.")
					break
				}
			}
			if flag == 0 {
				fmt.Println("Cannot find server with this id.Please input the right id.")
			} else {
				break
			}
		}
	}
	return listkey
}

// 设置gameserver各项属性,仅供内部调用
func gameServerSet() platform_server.GServer {
	var gameServer platform_server.GServer
	for {
		fmt.Println("Please input GameServer id:")
		tmp := getInput()
		gameServerId, err := strconv.Atoi(tmp)
		if err == nil {
			gameServer.GSId = int32(gameServerId)
			break
		}
		fmt.Println("Input Error.Please input the right id.")
	}
	fmt.Println("Please input GameServer ip:")
	tmp := getInput()
	gameServer.Ip = tmp
	fmt.Println("Please input GameServer port:")
	tmp = getInput()
	gameServer.Port = tmp
	fmt.Println("Please input GameServer Rpc ip:")
	tmp = getInput()
	gameServer.RPCIp = tmp
	fmt.Println("Please input GameServer Rpc port:")
	tmp = getInput()
	gameServer.RPCPort = tmp
	fmt.Println("Is this a hd server? please input y/n")
	tmp = getInput()
	if tmp == "y" || tmp == "Y" {
		gameServer.HD = true
	} else {
		gameServer.HD = false
	}

	return gameServer
}

func setOpenTime(openTimeAndId string, app string) {
	var mynumber int
	var position int
	for i := 0; i < len(openTimeAndId); i++ {
		if openTimeAndId[i] == ',' {
			if i == 0 || i == len(openTimeAndId) {
				fmt.Println("Format Error.Please input the right format.")
				return
			}
			mynumber++
			position = i
		}
	}
	if mynumber != 1 {
		fmt.Println("Format Error.Please input the right format.")
		return
	}

	openTime := openTimeAndId[:position]
	id := openTimeAndId[position+1:]
	myid, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Id format error.Please input the right format of id.")
		return
	}
	regexpTest := "^\\d{4}-\\d{1,2}-\\d{1,2}_\\d{1,2}:\\d{1,2}:\\d{1,2}$"
	pass, err := regexp.MatchString(regexpTest, openTime)
	if err != nil {
		fmt.Println("Regexp match fail!")
		return
	}
	if !pass {
		fmt.Println("Open time format error.Please input the right format of opentime.")
		fmt.Println(openTime)
		return
	}
	temp := regexp.MustCompile(`[\d]+`)
	mytime := temp.FindAllString(openTime, -1)
	var myTime [6]int
	for i := 0; i < 6; i++ {
		var err error
		myTime[i], err = strconv.Atoi(mytime[i])
		if err != nil {
			fmt.Println("Error when get time:", err)
			return
		}
	}
	if myTime[1] == 0 || myTime[1] > 12 {
		fmt.Println("Please input the right month.")
		return
	}
	t := time.Date(myTime[0], time.Month(myTime[1]), myTime[2], myTime[3], myTime[4], myTime[5], 0, time.Local)
	fmt.Println("The opentime is:", t.Format("2006-01-02 15:04:05.999999 -07:00"))

	err = platform_server.InitServerList()
	if err != nil {
		fmt.Println("Serverlist init error:", err)
		return
	}
	serverlist := platform_server.ServerList(app)
	for key, value := range serverlist {
		if value.Id == int32(myid) {
			value.OpenTime = t.Unix()
			if value.OpenTime <= time.Now().Unix() {
				value.Status = 1
			} else {
				value.Status = 0
			}
			delete(serverlist, key)
			serverlist[key] = value
			var local_list []platform_server.Server
			for _, values := range serverlist {
				local_list = append(local_list, values)
			}

			rev, err := importServerList(local_list, app)
			if err != nil {
				fmt.Println("Error when setting open time:", err)
				return
			} else {
				fmt.Println("Open time set successfully.")
				commitServerListByRevsion(rev, app)
				return
			}
		}
	}
	fmt.Println("Can not find server with id", myid, ".Please check if you have input the right server id.")
}

// 设置 hd_server 互动服务器
func setHDServer(id string, app string) {
	tmp, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Please input the right id of the server you want to open!")
		return
	}
	err = platform_server.InitServerList()
	if err != nil {
		fmt.Println("Server list error:", err)
		return
	}
	serverlist := platform_server.ServerList(app)
	var server platform_server.Server
	var flag int
	for key, value := range serverlist {
		for key1, value1 := range value.GServers {
			if value1.GSId == int32(tmp) {
				if value1.HD == true {
					fmt.Println("The game server is already a HD server.You can't set a server that is already HD server to be a HD Server.")
					fmt.Println("HD server set canceled.")
					return
				}
				flag = 1
				server = value
				server.GServers[key1].HD = true
				for key2, value2 := range server.GServers {
					if value2.HD == true && key2 != key1 {
						server.GServers[key2].HD = false
						break
					}
				}
				delete(serverlist, key)
				serverlist[key] = server
				break
			}
			if flag == 1 {
				break
			}
		}
	}
	if flag == 0 {
		fmt.Println("Cannot find server with id", id, ",please check if you have input the right id.")
		return
	}
	var local_list []platform_server.Server
	for _, value := range serverlist {
		local_list = append(local_list, value)
	}

	rev, err := importServerList(local_list, app)
	if err != nil {
		fmt.Println("Error when setting HD server:", err)
	} else {
		fmt.Println("Set game server with id", id, "HD server successfully!")
		for {
			fmt.Printf("Are you sure to commit revision:%v? y/n\n", rev)
			answer := getInput()
			if answer == "y" || answer == "Y" {
				commitServerListByRevsion(rev, app)
				break
			} else if answer == "n" || answer == "N" {
				break
			} else {
				fmt.Println("Wrong input！Please input the right format.")
			}
		}
	}
}

//修改GSId为id的game server的Ip和Port
func modifyIpAndPort(in string, app string) {
	arrs := strings.Split(in, ",")

	if len(arrs) != 5 {
		fmt.Println("please input right format such as -modifyIpAndPort=gsid,ip,port,rpcip,rpcport")
		return
	}

	err := platform_server.InitServerList()
	if err != nil {
		fmt.Println("Server list error:", err)
		return
	}
	serverlist := platform_server.ServerList(app)
	var server platform_server.Server
	var isFound = false
	gsid, err := strconv.Atoi(arrs[0])
	if err != nil {
		fmt.Println("gsid format error")
		return
	}
	for key, value := range serverlist {
		for key1, value1 := range value.GServers {
			if value1.GSId == int32(gsid) {
				server = value

				if strings.Trim(arrs[1], " ") != "" {
					server.GServers[key1].Ip = arrs[1]
				}
				if strings.Trim(arrs[2], " ") != "" {
					server.GServers[key1].Port = arrs[2]
				}
				if strings.Trim(arrs[3], " ") != "" {
					server.GServers[key1].RPCIp = arrs[3]
				}
				if strings.Trim(arrs[4], " ") != "" {
					server.GServers[key1].RPCPort = arrs[4]
				}

				delete(serverlist, key)
				serverlist[key] = server

				isFound = true
				break
			}
		}
	}
	if !isFound {
		fmt.Println("Cannot find game server with id", gsid, ",please check if you have input the right id.")
		return
	}
	var local_list []platform_server.Server
	for _, value := range serverlist {
		local_list = append(local_list, value)
	}

	rev, err := importServerList(local_list, app)
	if err != nil {
		fmt.Println("Error when setting HD server:", err)
	} else {
		commitServerListByRevsion(rev, app)
	}
}
