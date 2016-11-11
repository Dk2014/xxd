package platform_server

import (
	"core/log"
	"fmt"
	"os"
	"strings"
	"testing"
)

const (
	__redis_server = "114.112.58.162:56379"
)

func TestServerBefore(t *testing.T) {
	log.Setup("./log", false)
	InitRedis(__redis_server, Config{}.Apps)
}

func TestUpdateServerList(t *testing.T) {
	err := updateServerList()
	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestReadServerList(t *testing.T) {
	fmt.Println("==/server/list==")
	jsonServers := fmt.Sprintf(`
{
    "OpenId": "_test_id%v",
    "Type": 1,
    "Version": 123
}`, 12312313)
	route("/server/list", strings.NewReader(jsonServers), os.Stdout, "127.0.0.1")
	fmt.Println("")
}
