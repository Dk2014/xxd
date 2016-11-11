package platform_server

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestStatusBefore(t *testing.T) {
	go Tick(3)
}

func TestStatus(t *testing.T) {
	fmt.Println("==/stat==")
	jsonStat := `{"OpenId": "_test_id", "Type": 1, "Version": 1, "App": "xxd_qq"}`
	route("/stat", strings.NewReader(jsonStat), os.Stdout, "127.0.0.1")

	fmt.Println("")
}
