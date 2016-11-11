package platform_tool

import (
	"encoding/json"
	"fmt"
	"platform_server"
	"testing"
)

const (
	__redis_server = "114.112.58.162:56379"
)

func TestAddServerList(t *testing.T) {
	err := platform_server.InitRedis(__redis_server, platform_server.Config{}.Apps)
	if err != nil {
		t.Errorf("redis %v", err)
	}

	const jsonStream = `
  [
    {
      "Id": 1,
      "Type": 1,
      "Name": "内网测试",
      "Status": 1,
      "IsNew": false,
      "IsHot": false,
      "OpenTime": 1029,
      "GServers": [
        {
          "GSId": 1,
          "Ip": "svn.vvv.io",
          "Port": "8080",
          "HD": false
        }
      ]
    },
    {
      "Id": 2,
      "Type": 1,
      "Name": "策划专用",
      "Status": 1,
      "IsNew": false,
      "IsHot": false,
      "OpenTime": 1029,
      "GServers": [
        {
          "GSId": 2,
          "Ip": "cehua.vvv.io",
          "Port": "8080",
          "HD": false
        }
      ]
    }
    ,
    {
      "Id": 3,
      "Type": 1,
      "Name": "S0外网备用",
      "Status": 1,
      "IsNew": false,
      "IsHot": false,
      "OpenTime": 1029,
      "GServers": [
        {
          "GSId": 3,
          "Ip": "s0.xxd.pinidea.co",
          "Port": "8081",
          "HD": false
        }
      ]
    }
    ,
    {
      "Id": 4,
      "Type": 1,
      "Name": "S1外网备用",
      "Status": 1,
      "IsNew": true,
      "IsHot": true,
      "OpenTime": 1029,
      "GServers": [
        {
          "GSId": 4,
          "Ip": "s1.xxd.pinidea.co",
          "Port": "8081",
          "HD": false
        }
      ]
    }
    ,
    {
      "Id": 5,
      "Type": 1,
      "Name": "内网测试组专用",
      "Status": 1,
      "IsNew": false,
      "IsHot": false,
      "OpenTime": 1029,
      "GServers": [
        {
          "GSId": 5,
          "Ip": "172.26.160.11",
          "Port": "8081",
          "HD": false
        }
      ]
    }
  ]`

	fmt.Println("========== test platform import server list ============")
	var local_list []platform_server.Server
	err = json.Unmarshal([]byte(jsonStream), &local_list)
	if err != nil {
		fmt.Println("Error when unmarshaling: ", err)
		return
	}

	rev, err := importServerList(local_list, "xxd_qq")
	if err != nil || rev <= 0 {
		t.Errorf("importing server list error: %v", err)
	}

	if rev > 0 {
		commitServerListByRevsion(rev, "xxd_qq")
	}
}
