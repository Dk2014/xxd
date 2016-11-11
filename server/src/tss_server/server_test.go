package main

import (
	"bitbucket.org/PinIdea/xxd-tss-server/tss_sdk_go"
	"bytes"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"testing"
)

func Test(t *testing.T) {
	//init test data
	t_open_id := "oShxat0VfPkOaZq3F4G4eicN2YXo"
	t_plat_id := 1
	t_client_ip := "127.0.0.1"
	t_client_ver := 1234
	t_role_name := "nick name"
	t_anti_data := []byte("test data, received")

	//init tss_sdk
	tss_sdk_go.CallBack = &callback{}
	tss_sdk_go.Init(*optConfPath, *optLibPath, *optTimeInterval, uint(*optInstanceId))

	//rpc
	rpc_oper := new(RpcOper)
	rpc.Register(rpc_oper)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", *optPort)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)

	//client
	client, err := rpc.DialHTTP("tcp", "127.0.0.1"+*optPort)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var reply int

	//test user operation
	user_info := &UserInfo{OpenId: t_open_id, Plat_Id: t_plat_id, Client_Ip: t_client_ip, Client_Ver: t_client_ver, Role_Name: t_role_name}
	err = client.Call("RpcOper.UserLogin", user_info, &reply)
	if err != nil {
		log.Fatal("RpcOper error:", err)
	}
	fmt.Println("RpcOper.UserLogin:", reply)
	if reply != tss_sdk_go.TSS_SDK_PROC_OK {
		t.Error("RpcOper.UserLogin test fail.")
	}
	err = client.Call("RpcOper.UserLogout", user_info, &reply)
	if err != nil {
		log.Fatal("RpcOper error:", err)
	}
	fmt.Println("RpcOper.UserLogout:", reply)
	if reply != tss_sdk_go.TSS_SDK_PROC_OK {
		t.Error("RpcOper.UserLogout test fail.")
	}

	//test recvdata
	var b_reply []byte
	pkg_info := &PkgInfo{OpenId: t_open_id, Plat_Id: t_plat_id, Anti_Data: t_anti_data}
	err = client.Call("RpcOper.RecvData", pkg_info, &b_reply)
	if err != nil {
		log.Fatal("RpcOper error:", err)
	}
	fmt.Println("RpcOper.RecvData:", string(b_reply))
	if bytes.Compare(t_anti_data, b_reply) == 0 {
		log.Println("anti_data PASSED")
	} else {
		log.Println("want:", len(t_anti_data), t_anti_data)
		log.Println("get:", len(b_reply), b_reply)
		log.Fatal("anti_data ERROR")
	}
}
