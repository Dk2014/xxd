package platform_tool

import (
	"fmt"
	"strconv"
	"strings"

	"core/redis"
	"platform_server"
)

//操作release version
func operateReleaseVersion(typeAndVersion string, app string) {
	apps := make([]string, 0)
	if strings.Contains(app, ":") {
		apps = strings.Split(app, ":")
	} else {
		apps = append(apps, app)
	}

	for _, app := range apps {
		arr := strings.Split(typeAndVersion, ",")
		if len(arr) != 2 && len(arr) != 1 {
			fmt.Println("please intput right format such as -releaseVersion=1,1268, -releaseVersion=1")
			return
		}

		itype, err := strconv.Atoi(arr[0])
		if err != nil {
			fmt.Println("please intput right type:1-ios微信，2-ios手Q，17-android微信，18-android手Q")
			return
		}

		c := platform_server.GetDBConn(app)
		defer c.Close()

		release_key := platform_server.RedisKey_ClientReleaseVersionByType(uint8(itype))

		if len(arr) == 1 {
			version, err := redis.Int(c.Do("GET", release_key))
			if err != nil {
				fmt.Println("type is wrong or the release of type has been not setted")
				return
			}

			fmt.Printf("the release version of type %v is %v\n", itype, version)
			return
		}

		version, err := strconv.Atoi(arr[1])
		if err != nil {
			fmt.Println("please intput right release version, such as 1567")
			return
		}

		for {
			fmt.Printf("Are you sure to set release version of type %v: %v? y/n\n", itype, version)
			answer := getInput()
			if answer == "y" || answer == "Y" {
				c.Do("SET", release_key, version)
				break
			} else if answer == "n" || answer == "N" {
				break
			} else {
				fmt.Println("Wrong input！Please input the right format.")
			}
		}
	}
}

//操作audit version
func operateAuditVersion(typeAndVersion string, app string) {
	apps := make([]string, 0)
	if strings.Contains(app, ":") {
		apps = strings.Split(app, ":")
	} else {
		apps = append(apps, app)
	}

	for _, app := range apps {
		arr := strings.Split(typeAndVersion, ",")
		if len(arr) != 2 && len(arr) != 1 {
			fmt.Println("please intput right format such as -auditVersion=1,1268, -auditVersion=1")
			return
		}

		itype, err := strconv.Atoi(arr[0])
		if err != nil {
			fmt.Println("please intput right type:1-ios微信，2-ios手Q，17-android微信，18-android手Q")
			return
		}

		c := platform_server.GetDBConn(app)
		defer c.Close()

		audit_key := platform_server.RedisKey_ClientAuditVersionByType(uint8(itype))

		if len(arr) == 1 {
			version, err := redis.Int(c.Do("GET", audit_key))
			if err != nil {
				fmt.Println("type is wrong or the audit of type has been not setted")
				return
			}

			fmt.Printf("the audit version of type %v is %v\n", itype, version)
			return
		}

		version, err := strconv.Atoi(arr[1])
		if err != nil {
			fmt.Println("please intput right audit version, such as 1567")
			return
		}

		_, err = c.Do("SET", audit_key, version)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("ok")
		}

	}
}

//操作min version
func operateMinVersion(typeAndVersion string, app string) {
	apps := make([]string, 0)
	if strings.Contains(app, ":") {
		apps = strings.Split(app, ":")
	} else {
		apps = append(apps, app)
	}

	for _, app := range apps {
		arr := strings.Split(typeAndVersion, ",")
		if len(arr) != 2 && len(arr) != 1 {
			fmt.Println("please intput right format such as -minVersion=1,1268, -minVersion=1")
			return
		}

		itype, err := strconv.Atoi(arr[0])
		if err != nil {
			fmt.Println("please intput right type:1-ios微信，2-ios手Q，17-android微信，18-android手Q")
			return
		}

		c := platform_server.GetDBConn(app)
		defer c.Close()

		min_key := platform_server.RedisKey_ClientMinVersionByType(uint8(itype))

		if len(arr) == 1 {
			version, err := redis.Int(c.Do("GET", min_key))
			if err != nil {
				fmt.Println("type is wrong or the min version of type has been not setted")
				return
			}

			fmt.Printf("the min version of type %v is %v\n", itype, version)
			return
		}

		version, err := strconv.Atoi(arr[1])
		if err != nil {
			fmt.Println("please intput right release version, such as 1567")
			return
		}

		for {
			fmt.Printf("Are you sure to set min version of type %v: %v? y/n\n", itype, version)
			answer := getInput()
			if answer == "y" || answer == "Y" {
				c.Do("SET", min_key, version)
				break
			} else if answer == "n" || answer == "N" {
				break
			} else {
				fmt.Println("Wrong input！Please input the right format.")
			}
		}
	}
}

//操作版本黑名单
func operateVersionBlackTable(flag string, app string) {
	arr := strings.Split(flag, ",")
	if len(arr) > 3 || len(arr) < 2 {
		fmt.Println("please intput right format: -versionBlack=<list/add/rm, type, version>")
		return
	}

	c := platform_server.GetDBConn(app)
	defer c.Close()

	cmd := arr[0]
	itype, err := strconv.Atoi(arr[1])
	if err != nil {
		fmt.Println("wrong type format")
		return
	}

	version_black_table_key := platform_server.RedisKey_ClientVersionBlackTable(uint8(itype))
	switch cmd {
	case "list":
		tables, err := redis.Strings(c.Do("SMEMBERS", version_black_table_key))
		if err != nil {
			fmt.Printf("run smembers error, version_black_table_key:%v\n", version_black_table_key)
			return
		}
		fmt.Println("current black talbes:", tables)
	case "add":
		if arr[2] == "" {
			fmt.Println("error: version is empty!!")
			return
		}
		_, err := c.Do("SADD", version_black_table_key, arr[2])
		if err != nil {
			fmt.Println(err)
		}

		tables, err := redis.Strings(c.Do("SMEMBERS", version_black_table_key))
		if err != nil {
			fmt.Printf("run smembers error, version_black_table_key:%v\n", version_black_table_key)
			return
		}
		fmt.Println("current black talbes:", tables)
	case "rm":
		if arr[2] == "" {
			fmt.Println("error: version is empty!!")
			return
		}
		_, err := c.Do("SREM", version_black_table_key, arr[2])
		if err != nil {
			fmt.Println(err)
		}
		tables, err := redis.Strings(c.Do("SMEMBERS", version_black_table_key))
		if err != nil {
			fmt.Printf("run smembers error, version_black_table_key:%v\n", version_black_table_key)
			return
		}
		fmt.Println("current black talbes:", tables)
	default:
		fmt.Println("wrong cmd")
	}
}

//操作客户端升级地址
func operateUpgradeUrl(flag string, app string) {
	arr := strings.Split(flag, ",")
	if len(arr) != 2 && len(arr) != 1 {
		fmt.Println("please intput right format: -upgradeUrl=<type[,url]")
		return
	}

	itype, err := strconv.Atoi(arr[0])
	if err != nil {
		fmt.Println("please intput right type:1-ios微信，2-ios手Q，17-android微信，18-android手Q")
		return
	}

	c := platform_server.GetDBConn(app)
	defer c.Close()

	//get current upgrade url
	if len(arr) == 1 {
		url_key := platform_server.RedisKey_ClientUpgradeUrl(uint8(itype))
		url, err := redis.String(c.Do("GET", url_key))
		if err != nil {
			fmt.Println("type is wrong or upgrade url of type has been not setted")
			return
		}

		fmt.Printf("current upgrade url of type %v is %v\n", itype, url)
		return
	}

	//set current url
	url := arr[1]
	if url == "" {
		fmt.Println("please input url content")
		return
	}

	for {
		fmt.Printf("Are you sure to set content of type %v? y/n\n", itype)
		answer := getInput()
		if answer == "y" || answer == "Y" {
			url_key := platform_server.RedisKey_ClientUpgradeUrl(uint8(itype))
			c.Do("SET", url_key, url)

			url, err := redis.String(c.Do("GET", url_key))
			if err != nil {
				fmt.Println("type is wrong or upgrade url of type has been not setted")
				return
			}

			fmt.Printf("current upgrade url of type %v is %v\n", itype, url)
			break
		} else if answer == "n" || answer == "N" {
			break
		} else {
			fmt.Println("Wrong input！Please input the right format.")
		}
	}
}
