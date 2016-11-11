package tss_sdk_go

/*
#include "tss_sdk_go.h"
#include "tss_sdk_anti.h"
#include <stdio.h>
#include <string.h>

#if defined(WIN32) || defined(_WIN64)
#include <windows.h>
#else
#include <stdlib.h>
#endif

TssSdkProcResult on_send_data_to_client_cgo(TssSdkAntiSendDataInfoEx *anti_data);
TssSdkProcResult add_user_cgo(TssSdkAntiInterfEx *anti_interf_, char *openid, int plat_id, char *client_ip, int client_ver,char *role_name);
TssSdkProcResult del_user_cgo(TssSdkAntiInterfEx *anti_interf_, char *openid, int plat_id);
TssSdkProcResult recv_anti_data_cgo(TssSdkAntiInterfEx *anti_interf_,char *openid, int plat_id, unsigned char *anti_data, unsigned int data_len);

*/
import "C"
import (
	"core/debug"
	clog "core/log"
	"log"
	"time"
	"unsafe"
)

const (
	TSS_SDK_PROC_OK           = int(C.TSS_SDK_PROC_OK)
	TSS_SDK_PROC_INVALID_ARG  = int(C.TSS_SDK_PROC_INVALID_ARG)
	TSS_SDK_PROC_INTERNAL_ERR = int(C.TSS_SDK_PROC_INTERNAL_ERR)
	TSS_SDK_PROC_FAIL         = int(C.TSS_SDK_PROC_FAIL)

	TSS_ACTION_USERLOGIN  = 1
	TSS_ACTION_USERLOGOUT = 2
	TSS_ACTION_RECV       = 3
	TSS_ACTION_PROC       = 4
)

var (
	busi_interf_ *C.TssSdkBusiInterf
	init_data    C.TssSdkInitInfo
	init_info    C.TssSdkAntiInitInfoEx
	anti_interf_ *C.TssSdkAntiInterfEx
	CallBack     CallBackInterf
)

type SyncUnit struct {
	Act  int
	Data interface{}
}

type CallBackInterf interface {
	Send_Data_To_Client(openid string, plat_id int, anti_data []byte) int
}

//export on_send_data_to_client
func on_send_data_to_client(openid_ *C.char, plat_id_ C.int, anti_data_ *C.char, anti_data_len_ C.int) (ret C.TssSdkProcResult) {
	//CHECK_POINT_3 PASSED.
	defer func() {
		if r := recover(); r != nil {
			log.Println("on_send_data_to_client recover:", r)
			ret = C.TSS_SDK_PROC_INTERNAL_ERR
		}
	}()
	ret = C.TssSdkProcResult(CallBack.Send_Data_To_Client(C.GoString(openid_), int(plat_id_), C.GoBytes(unsafe.Pointer(anti_data_), anti_data_len_)))
	return
}

//load library, get anti interface and run cron
func Init(conf_path, lib_path string, t_interval int64, instance_id uint) chan SyncUnit {
	if CallBack == nil {
		log.Fatal("CallBack nil.")
	}
	//load library
	//CHECK_POINT_1 PASSED.
	init_data.unique_instance_id_ = C.uint(instance_id)
	init_data.tss_sdk_conf_ = C.CString(conf_path) //path not effect int test
	busi_interf_ = C.tss_sdk_load((*C.TSS_TCHAR)(C.CString(lib_path)), &init_data)

	//get anti_interf_
	init_info.send_data_to_client_ = (C.TssSdkSendDataToClientEx)(unsafe.Pointer(C.on_send_data_to_client_cgo))
	anti_interf_ = C.M_TSS_SDK_GET_ANTI_INTERF_EX(&init_info)

	syncChan := make(chan SyncUnit, 100000)

	//cron
	go func() {
		defer func() {
			if err := recover(); err != nil {
				clog.Errorf(`tss_server.tss_sdk.Init
Error = %v
Stack = 
%s`,
					err,
					debug.Stack(1, "    "),
				)
			}
		}()
		ticker := time.NewTicker(time.Millisecond * time.Duration(t_interval))
		for _ = range ticker.C {
			syncChan <- SyncUnit{Act: TSS_ACTION_PROC}
		}
	}()

	return syncChan
}

//user login
func UserLogin(open_id string, plat_id int, client_ip string, client_ver int, role_name string) int {
	//CHECK_POINT_4 PASSED.
	c_open_id := C.CString(open_id)
	defer C.free(unsafe.Pointer(c_open_id))
	c_client_ip := C.CString(client_ip)
	defer C.free(unsafe.Pointer(c_client_ip))
	c_role_name := C.CString(role_name)
	defer C.free(unsafe.Pointer(c_role_name))

	ret := C.add_user_cgo(anti_interf_, c_open_id, C.int(plat_id), c_client_ip, C.int(client_ver), c_role_name)
	return int(ret)
}

//user logout
func UserLogout(open_id string, plat_id int) int {
	//CHECK_POINT_5 PASSED.
	c_open_id := C.CString(open_id)
	defer C.free(unsafe.Pointer(c_open_id))

	ret := C.del_user_cgo(anti_interf_, c_open_id, C.int(plat_id))
	return int(ret)
}

//receive data from client
func RecvData(open_id string, plat_id int, anti_data []byte) int {
	//CHECK_POINT_6 PASSED.
	c_open_id := C.CString(open_id)
	defer C.free(unsafe.Pointer(c_open_id))

	ret := C.recv_anti_data_cgo(anti_interf_, c_open_id, C.int(plat_id), (*C.uchar)(unsafe.Pointer(&anti_data[0])), C.uint(len(anti_data)))
	return int(ret)
}

//cron call proc
func CallProc() {
	//CHECK_POINT_7 PASSED.
	C.call_proc(busi_interf_)
}
