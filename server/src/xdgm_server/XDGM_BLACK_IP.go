package xdgm_server

import (
	"fmt"
)

func (this *XDGM_BLACK_IP_REQ) Process() (XDGM_RSP, error) {
	rsp := XDGM_RSP{}
	errorCode, message, list := BlackIp(this.Ip, this.Mode)
	fmt.Println(errorCode)
	if errorCode == 0 {
		rsp.Status = 1
		rsp.Message = "success"
	} else if errorCode == 1 {
		rsp.Status = -1
		rsp.Message = "sign error"
	} else if errorCode == 2 {
		rsp.Status = -1
		rsp.Message = "server error"
	} else if errorCode == -1 {
		rsp.Status = -1
		rsp.Message = message
	} else {
		rsp.Status = -1
		rsp.Message = "some other error"
	}
	rsp.Data = list
	return rsp, nil
}
