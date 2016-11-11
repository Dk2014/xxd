package platform_tool

import (
	"fmt"
	"strconv"
	"strings"

	"core/redis"
	"platform_server"
)

//操作patch的url前缀
func operatePatchUrl(patchUrl string, app string) {
	apps := make([]string, 0)
	if strings.Contains(app, ":") {
		apps = strings.Split(app, ":")
	} else {
		apps = append(apps, app)
	}

	for _, app := range apps {
		c := platform_server.GetDBConn(app)
		defer c.Close()
		patchUrlKey := platform_server.RedisKey_ClientPatchUrl()
		if patchUrl == "show" {
			cur_patch_url, err := redis.String(c.Do("GET", patchUrlKey))
			if err != nil {
				fmt.Println("patch url not found")
				return
			}
			fmt.Println("current patch url:" + cur_patch_url)
			return
		}

		_, err := c.Do("SET", patchUrlKey, patchUrl)
		if err != nil {
			fmt.Println(err)
		}
	}
}

//操作resource的url前缀
func operateResourceUrl(resourceUrl string, app string) {
	apps := make([]string, 0)
	if strings.Contains(app, ":") {
		apps = strings.Split(app, ":")
	} else {
		apps = append(apps, app)
	}

	for _, app := range apps {
		c := platform_server.GetDBConn(app)
		defer c.Close()
		resourceUrlKey := platform_server.RedisKey_ClientResourceUrl()
		if resourceUrl == "show" {
			cur_resource_url, err := redis.String(c.Do("GET", resourceUrlKey))
			if err != nil {
				fmt.Println("resource url not found")
				return
			}
			fmt.Println("current resource url:" + cur_resource_url)
			return
		}

		_, err := c.Do("SET", resourceUrlKey, resourceUrl)
		if err != nil {
			fmt.Println(err)
		}
	}
}

/*
	操作patch
	example:
	-patch=1
	-patch=1:2:5:17:18,version1_version2_patch.dat:version3_version2_patch.dat
*/
func operatePatch(flag string, app string) {
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
		if len(strs) == 1 {
			//列出当前type的所有patch
			sType := strs[0]
			iType, err := strconv.Atoi(sType)
			if err != nil {
				fmt.Println("wrong type!")
				return
			}
			key_patch_hash := platform_server.RedisKey_ClientPatchHash(uint8(iType))
			patches, err := redis.Values(c.Do("HGETALL", key_patch_hash))
			if err != nil {
				fmt.Println("patches not found")
				return
			}
			fmt.Printf("%s\n", patches)
		} else if len(strs) == 2 {
			types := strings.Split(strs[0], ":")
			for _, val := range types {
				iType, err := strconv.Atoi(val)
				if err != nil {
					fmt.Println("wrong type!")
					continue
				}
				key_patch_hash := platform_server.RedisKey_ClientPatchHash(uint8(iType))

				patches := strings.Split(strs[1], ":")

				for _, patch := range patches {
					arrs := strings.Split(patch, "_")
					if len(arrs) != 3 {
						continue
					}

					version := arrs[0] + "_" + arrs[1]
					patch := arrs[2]
					//如果patch为空，删除此版本patch配置
					if strings.Trim(patch, " ") == "" {
						count, err := c.Do("HDEL", key_patch_hash, version)

						if err != nil {
							fmt.Println(err)
						} else {
							fmt.Println(fmt.Sprintf("del count: %d", count))
						}
						continue
					}
					_, err = c.Do("HSET", key_patch_hash, version, patch)
					if err != nil {
						fmt.Println(err)
					}
				}
			}
		} else {
			fmt.Println("wroing input")
		}
	}
}

/*
	城镇等资源的相关操作
	-totalResource=list
	-totalResource=19000
	-totalResource=19000,2_dat_size:3_dat_size:4_dat_size
*/
func operateTotalResource(flag string, app string) {
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
		if len(strs) == 1 {
			//列出所有TotalResource
			if strings.EqualFold(strs[0], "list") {
				versions, err := redis.Strings(c.Do("SMEMBERS", "total_resource_file_version"))
				if err != nil {
					fmt.Println(err)
				}
				for _, v := range versions {
					version, err := strconv.Atoi(v)
					if err != nil {
						fmt.Println(err)
					}
					hash_key := platform_server.RedisKey_TotalResourceHash(int32(version))
					resource, err := redis.Values(c.Do("HGETALL", hash_key))
					if err != nil {
						fmt.Println("resource not found")
						return
					}
					fmt.Println(v + ":")
					fmt.Printf("%s\n", resource)
				}
			} else {
				version, err := strconv.Atoi(strs[0])
				if err != nil {
					fmt.Println(err)
				}
				hash_key := platform_server.RedisKey_TotalResourceHash(int32(version))
				resource, err := redis.Values(c.Do("HGETALL", hash_key))
				if err != nil {
					fmt.Println("resource not found")
					return
				}
				fmt.Println(strs[0] + ":")
				fmt.Printf("%s\n", resource)
			}
		} else if len(strs) == 2 {
			version, err := strconv.Atoi(strs[0])
			if err != nil {
				fmt.Println(err)
			}

			resources := strings.Split(strs[1], ":")
			for _, resource := range resources {
				arrs := strings.Split(resource, "_")
				if len(arrs) != 3 {
					fmt.Println("wrong resource config:" + resource)
					continue
				}
				id := arrs[0]
				hash_key := platform_server.RedisKey_TotalResourceHash(int32(version))
				path := arrs[1]
				size := arrs[2]
				//如果resource为空，删除此resource配置
				if strings.Trim(path, " ") == "" {
					count, err := c.Do("HDEL", hash_key, id)

					if err != nil {
						fmt.Println(err)
					} else {
						fmt.Println(fmt.Sprintf("del count: %d", count))
					}
					continue
				}

				_, err = c.Do("HSET", hash_key, id, path+","+size)
				if err != nil {
					fmt.Println(err)
				}
				_, err = c.Do("SADD", "total_resource_file_version", version)
				if err != nil {
					fmt.Println(err)
				}
			}
		} else {
			fmt.Println("wroing input, please input string such as -totalResource=list, -totalResource=6553,1, -totalResource=6553,1,-totalResource.dat")
		}
	}
}
