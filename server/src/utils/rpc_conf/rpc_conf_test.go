package rpc_conf

import (
	"fmt"
	"testing"
)

func TestRPCList(t *testing.T) {
	list := RPCList("http://platform.xxd.io:8888/gserverall", "xxd_vn")
	fmt.Printf("resutl : %#v\n", list)
}
