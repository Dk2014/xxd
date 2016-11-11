package platform_tool

import (
	"core/redis"
	"fmt"
	"platform_server"
	"strconv"
	"strings"
)

func currentRevision(app string) {
	apps := make([]string, 0)
	if strings.Contains(app, ":") {
		apps = strings.Split(app, ":")
	} else {
		apps = append(apps, app)
	}

	for _, app := range apps {
		c := platform_server.GetDBConn(app)
		defer c.Close()

		cur_rev, err := redis.Int(c.Do("GET", platform_server.RedisKey_GameServerListCurrentRev()))
		if err != nil {
			fmt.Println("Error on geting current revision")
			return
		}
		fmt.Println("current revison:", cur_rev)
	}
}

//操作平台服公告
func operateAnnounce(flag string, app string) {
	arr := strings.Split(flag, ",")
	if len(arr) != 2 && len(arr) != 1 {
		fmt.Println("please intput right format such as -announce=1,'公告', -announce=1")
		return
	}

	itype, err := strconv.Atoi(arr[0])
	if err != nil {
		fmt.Println("please intput right type:1-ios微信，2-ios手Q，17-android微信，18-android手Q")
		return
	}

	c := platform_server.GetDBConn(app)
	defer c.Close()

	//get current announce
	if len(arr) == 1 {
		rev_key := platform_server.RedisKey_AnnounceCurrentRevByType(uint8(itype))
		rev, err := redis.Int(c.Do("GET", rev_key))
		if err != nil {
			fmt.Println("type is wrong or announce of type has been not setted")
			return
		}

		announce_key := platform_server.RedisKey_AnnounceByTypeRevision(uint8(itype), int32(rev))
		announce, err := redis.String(c.Do("GET", announce_key))
		if err != nil {
			fmt.Printf("current announce of type %v is not found, error:%v\n", itype, err)
			return
		}

		fmt.Printf("current announce of type %v is %v\n", itype, announce)
		return
	}

	//set current announce
	announce := arr[1]
	if announce == "" {
		fmt.Println("please input announce content")
		return
	}

	for {
		fmt.Printf("Are you sure to set content of type %v? y/n\n", itype)
		answer := getInput()
		if answer == "y" || answer == "Y" {
			rev_key := platform_server.RedisKey_AnnounceCurrentRevByType(uint8(itype))
			rev, err := redis.Int(c.Do("GET", rev_key))
			if err != nil {
				fmt.Println("rev not found, so use 0")
				rev = 0
			}

			next_rev := rev + 1
			announce_key := platform_server.RedisKey_AnnounceByTypeRevision(uint8(itype), int32(next_rev))
			announce = strings.Replace(announce, "\\n", "\n", -1)
			c.Do("SET", rev_key, next_rev)
			c.Do("SET", announce_key, announce)

			break
		} else if answer == "n" || answer == "N" {
			break
		} else {
			fmt.Println("Wrong input！Please input the right format.")
		}
	}
}

//操作openid白名单
func operateOpenidWhiteTable(flag string, app string) {
	arr := strings.Split(flag, ",")
	if len(arr) != 2 && len(arr) != 1 {
		fmt.Println("please intput right format: -openidWhite=<list/add/rm, openid>")
		return
	}

	c := platform_server.GetDBConn(app)
	defer c.Close()

	cmd := arr[0]
	openid_white_table_key := platform_server.RedisKey_OpenidWhiteTable()
	switch cmd {
	case "list":
		white_tables, err := redis.Strings(c.Do("SMEMBERS", openid_white_table_key))
		if err != nil {
			fmt.Printf("run smembers error, openid_white_table_key:%v\n", openid_white_table_key)
			return
		}

		fmt.Println("current white talbes:", white_tables)
	case "add":
		openid := arr[1]
		if openid == "" {
			fmt.Println("openid is empty!!")
			return
		}

		_, err := c.Do("SADD", openid_white_table_key, openid)
		if err != nil {
			fmt.Println(err)
		}

		white_tables, err := redis.Strings(c.Do("SMEMBERS", openid_white_table_key))
		if err != nil {
			fmt.Printf("run smembers error, openid_white_table_key:%v\n", openid_white_table_key)
			return
		}

		fmt.Println("current white talbes:", white_tables)
	case "rm":
		openid := arr[1]
		if openid == "" {
			fmt.Println("openid is empty!!")
			return
		}
		_, err := c.Do("SREM", openid_white_table_key, openid)
		if err != nil {
			fmt.Println(err)
		}

		white_tables, err := redis.Strings(c.Do("SMEMBERS", openid_white_table_key))
		if err != nil {
			fmt.Printf("run smembers error, openid_white_table_key:%v\n", openid_white_table_key)
			return
		}

		fmt.Println("current white talbes:", white_tables)
	default:
		fmt.Println("wrong cmd")
	}
}

//开启或者关闭活动
func setActionAudit(flag string, app string) {
	apps := make([]string, 0)
	if strings.Contains(app, ":") {
		apps = strings.Split(app, ":")
	} else {
		apps = append(apps, app)
	}

	for _, app := range apps {
		c := platform_server.GetDBConn(app)
		defer c.Close()

		strs := strings.Split(flag, ",")
		if len(strs) != 2 {
			fmt.Println("please input right format, such as 1,0")
			return
		}

		isAudit := strs[1]

		itype, err := strconv.Atoi(strs[0])
		if err != nil {
			fmt.Println("please intput right type:1-ios微信，2-ios手Q，17-android微信，18-android手Q")
			return
		}

		key := platform_server.RedisKey_ActionAudit(uint8(itype))
		_, err = c.Do("set", key, isAudit)

		if err != nil {
			fmt.Printf("set %s error:%v\n", key, err)
		}
	}
}

//开启或者关闭维护
func setCloseStatus(sType string, iClose int, app string) {
	apps := make([]string, 0)
	if strings.Contains(app, ":") {
		apps = strings.Split(app, ":")
	} else {
		apps = append(apps, app)
	}

	for _, app := range apps {
		c := platform_server.GetDBConn(app)
		defer c.Close()

		itype, err := strconv.Atoi(sType)
		if err != nil {
			fmt.Println("please intput right type:1-ios微信，2-ios手Q，17-android微信，18-android手Q")
			return
		}

		key := platform_server.RedisKey_CloseTypeServers(uint8(itype))
		_, err = c.Do("set", key, iClose)

		if err != nil {
			fmt.Printf("set %s error:%v\n", key, err)
		}
	}
}

//灰度服设置
func setDarkLaunch(flag string, app string) {
	arrs := strings.Split(flag, ",")

	c := platform_server.GetDBConn(app)
	defer c.Close()

	if flag == "show" {
		sids := []uint8{1, 2, 17, 18}
		for _, v := range sids {
			fmt.Printf("type:%d\n", v)
			key := platform_server.RedisKey_DarkLaunch(v)
			val, err := redis.String(c.Do("get", key))
			if err != nil {
				fmt.Println("empty")
			} else {
				fmt.Println(val)
			}
		}

		return
	}

	if len(arrs) == 2 {
		if arrs[0] != "del" {
			fmt.Println("wrong input format, such -darkLaunch=del,type")
			return
		}

		itype, err := strconv.Atoi(arrs[1])
		if err != nil {
			fmt.Println("please intput right type:1-ios微信，2-ios手Q，17-android微信，18-android手Q")
			return
		}

		key := platform_server.RedisKey_DarkLaunch(uint8(itype))
		_, err = c.Do("del", key)
		if err != nil {
			fmt.Println("del error")
		} else {
			fmt.Println("ok")
		}

		return
	}

	if len(arrs) != 5 {
		fmt.Println("please intput right format: -darkLaunch=<type,sid1:sid2,version,0,url>)")
		return
	}

	itype, err := strconv.Atoi(arrs[0])
	if err != nil {
		fmt.Println("please intput right type:1-ios微信，2-ios手Q，17-android微信，18-android手Q")
		return
	}

	key := platform_server.RedisKey_DarkLaunch(uint8(itype))
	val := strings.Join(arrs[1:], ",")
	_, err = c.Do("set", key, val)

	if err != nil {
		fmt.Printf("set %s error:%v\n", key, err)
	}
}

//操作屏蔽活动图片
func operateDisableActionPic(flag string, app string) {
	arr := strings.Split(flag, ",")
	if len(arr) < 1 || len(arr) > 3 {
		fmt.Println("please intput right format: -disableActionPic=<type, list/add/rm, pic>")
		return
	}

	c := platform_server.GetDBConn(app)
	defer c.Close()

	itype, err := strconv.Atoi(arr[0])
	if err != nil {
		fmt.Println("please intput right type:1-ios微信，2-ios手Q，5-ios游客，17-android微信，18-android手Q")
		return
	}

	cmd := arr[1]
	disable_picture_key := platform_server.RedisKey_DisableActionPic(uint8(itype))
	switch cmd {
	case "list":
		pictures, err := redis.Strings(c.Do("SMEMBERS", disable_picture_key))
		if err != nil {
			fmt.Printf("run smembers error, disable_picture_key:%v\n", disable_picture_key)
			return
		}

		fmt.Println("current disabled pictures:", pictures)
	case "add":
		pic := arr[2]
		if pic == "" {
			fmt.Println("openid is empty!!")
			return
		}

		_, err := c.Do("SADD", disable_picture_key, pic)
		if err != nil {
			fmt.Println(err)
		}

		pics, err := redis.Strings(c.Do("SMEMBERS", disable_picture_key))
		if err != nil {
			fmt.Printf("run smembers error, disable_picture_key:%v\n", disable_picture_key)
			return
		}

		fmt.Println("current disabled pictures:", pics)
	case "rm":
		pic := arr[2]
		if pic == "" {
			fmt.Println("openid is empty!!")
			return
		}
		_, err := c.Do("SREM", disable_picture_key, pic)
		if err != nil {
			fmt.Println(err)
		}

		pictures, err := redis.Strings(c.Do("SMEMBERS", disable_picture_key))
		if err != nil {
			fmt.Printf("run smembers error, disable_picture_key:%v\n", disable_picture_key)
			return
		}

		fmt.Println("current disabled pictures:", pictures)
	default:
		fmt.Println("wrong cmd")
	}
}

//操作ip黑名单
func operateIpBlackTable(flag string, app string) {
	arr := strings.Split(flag, ",")
	if len(arr) > 2 || len(arr) < 1 {
		fmt.Println("please intput right format: -ipBlack=<list/add/rm, ip>")
		return
	}

	c := platform_server.GetDBConn(app)
	defer c.Close()

	cmd := arr[0]
	ip_black_table_key := platform_server.RedisKey_IpBlackTable()
	switch cmd {
	case "list":
		tables, err := redis.Strings(c.Do("SMEMBERS", ip_black_table_key))
		if err != nil {
			fmt.Printf("run smembers error, ip_black_table_key:%v\n", ip_black_table_key)
			return
		}
		fmt.Println("current black talbes:", tables)
	case "add":
		if arr[1] == "" {
			fmt.Println("error: ip is empty!!")
			return
		}
		_, err := c.Do("SADD", ip_black_table_key, arr[1])
		if err != nil {
			fmt.Println(err)
		}

		tables, err := redis.Strings(c.Do("SMEMBERS", ip_black_table_key))
		if err != nil {
			fmt.Printf("run smembers error, ip_black_table_key:%v\n", ip_black_table_key)
			return
		}
		fmt.Println("current black talbes:", tables)
	case "rm":
		if arr[1] == "" {
			fmt.Println("error: ip is empty!!")
			return
		}
		_, err := c.Do("SREM", ip_black_table_key, arr[1])
		if err != nil {
			fmt.Println(err)
		}
		tables, err := redis.Strings(c.Do("SMEMBERS", ip_black_table_key))
		if err != nil {
			fmt.Printf("run smembers error, ip_black_table_key:%v\n", ip_black_table_key)
			return
		}
		fmt.Println("current black talbes:", tables)
	default:
		fmt.Println("wrong cmd")
	}
}

func initDB(app string) {
	c := platform_server.GetDBConn(app)
	defer c.Close()

	//初始化patch url
	patchUrlKey := platform_server.RedisKey_ClientPatchUrl()
	_, err := c.Do("SET", patchUrlKey, "http://please.set.url/")
	if err != nil {
		fmt.Println(err)
		return
	}

	//初始化town url
	resourceUrlKey := platform_server.RedisKey_ClientResourceUrl()
	_, err = c.Do("SET", resourceUrlKey, "http://please.set.url/")
	if err != nil {
		fmt.Println(err)
		return
	}
}
