package platform_server

import (
	"core/log"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"testing"
)

func TestPatchBefore(t *testing.T) {
	log.Setup("./log", false)
	InitRedis("114.112.58.162:56379", Config{}.Apps)
}

func TestPatch(t *testing.T) {
	fmt.Println("==/patch==")
	jsonPatch := `{"Type": 1, "Version": 100, "ServerVersion": 123}`
	route("/patch", strings.NewReader(jsonPatch), os.Stdout, "127.0.0.1")

	fmt.Println("")
}

func TestAnnounce(t *testing.T) {
	fmt.Println("==/announce==")
	jsonPatch := `{"Type": 1, "AnnounceRevision": 1}`
	route("/announce", strings.NewReader(jsonPatch), os.Stdout, "127.0.0.1")

	fmt.Println("")
}

func TestSetAnnounce(t *testing.T) {
	fmt.Println("==/set announce==")

	result := SetAnnounce(1, "xxd_qq", fmt.Sprintf("新闻\n 内容%v", rand.Int()), "新闻公告", "2014-08-20")

	fmt.Println("set announce:", result)
}

func TestSystemInfo(t *testing.T) {
	fmt.Println("==/systemInfo==")

	jsonSystemInfo := `{"Type": 1, "Version": 1}`
	route("/systemInfo", strings.NewReader(jsonSystemInfo), os.Stdout, "127.0.0.1")

	fmt.Println("")
}
