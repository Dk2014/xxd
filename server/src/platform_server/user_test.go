package platform_server

import (
	"bytes"
	"core/log"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"
)

func TestUser(t *testing.T) {
	log.Setup("./log", false)
	InitRedis(__redis_server, Config{}.Apps)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	id := r.Int63()
	fmt.Println("id:", id)

	// test user gsinfo
	fmt.Println("==/user/gsinfo==")
	jsonGSinfo := fmt.Sprintf(`
  {
    "OpenId": "_test_id%v",
    "Type": 1,
    "Sid": 1,
	"Token": "_SANDBOX_",
	"App": "xxd_qq"
  }`, id)

	matched := true
	var last bytes.Buffer
	for i := 0; i < 10; i++ {
		var this bytes.Buffer
		route("/user/gsinfo", strings.NewReader(jsonGSinfo), &this, "127.0.0.1")
		fmt.Println(this.String())

		if i > 0 && this.String() != last.String() {
			matched = false
			break
		}

		last = this
	}

	if !matched {
		t.Errorf("/user/gsinfo is not consistence")
	}

	// test user creation
	fmt.Println("==/user/create==")
	jsonCreation := fmt.Sprintf(`
  {
    "OpenId": "_test_id%v",
    "Type": 1,
    "Sid": 1,
    "RoleId": 1,
    "Nick": "_test_nick%v",
		"Token": "_SANDBOX_"
  }`, id, id)

	route("/user/create", strings.NewReader(jsonCreation), os.Stdout, "127.0.0.1")
	fmt.Println("")

	fmt.Println("==/user/logon==")

	// test exist user logon
	jsonLogon := fmt.Sprintf(`
  {
    "OpenId": "_test_id%v",
    "Type": 1,
    "Sid": 1,
		"Token": "_SANDBOX_"
  }`, id)

	route("/user/logon", strings.NewReader(jsonLogon), os.Stdout, "127.0.0.1")

	fmt.Println("")
	fmt.Println("==/user/list==")

	// test user role list
	jsonList := fmt.Sprintf(`
  {
    "OpenId": "_test_id%v",
    "Type": 1
  }`, id)
	route("/user/list", strings.NewReader(jsonList), os.Stdout, "127.0.0.1")

	fmt.Println("")
	fmt.Println("==/user/update==")

	// test user update and display diff
	jsonUpdate := fmt.Sprintf(`
  {
    "OpenId": "_test_id%v",
    "Type": 1,
    "Sid": 1,
    "RoleLevel": 3,
		"Token": "_SANDBOX_"
  }`, id)
	route("/user/update", strings.NewReader(jsonUpdate), os.Stdout, "127.0.0.1")

	fmt.Println("")
	fmt.Println("==/user/list==")

	// test user role list
	jsonList = fmt.Sprintf(`
  {
    "OpenId": "_test_id%v",
    "Type": 1
  }`, id)
	route("/user/list", strings.NewReader(jsonList), os.Stdout, "127.0.0.1")

	fmt.Println("")
	// test none-exist user logon
	// test exist user logon
	fmt.Println("==/user/logon  none-exist user==")
	jsonLogon = `
  {
    "OpenId": "_test_id_not_exist",
    "Type": 1,
    "Sid": 1,
		"Token": "_SANDBOX_"
  }`

	route("/user/logon", strings.NewReader(jsonLogon), os.Stdout, "127.0.0.1")
	fmt.Println("")

	fmt.Println("==GetGServerInfoByOpenIdSid==")
	v, ok := GetGServerInfoByOpenIdSid(fmt.Sprintf("_test_id%v", id), 1, 1, "xxd_qq")
	fmt.Println(v)
	if !ok {
		t.Errorf("Test GetGServerInfoByOpenIdSid Failed")
	}
	fmt.Println("")
}
