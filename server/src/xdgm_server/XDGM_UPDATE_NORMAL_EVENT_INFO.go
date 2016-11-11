package xdgm_server

import (
	"core/log"
	"fmt"
	"xdgm_server/rpc"
)

func (this *XDGM_UPDATE_EVENTS_INFO_REQ) Process(platid uint8, sid int, events_info string) (XDGM_RSP, error) {
	rsp := XDGM_RSP{}
	gs, exists := GetGsServerByHttp(platid, sid)
	if exists {
		for _, v := range gs {
			var success bool
			rpc.XdgmUpdateNormalEventInfo(int(v.GSID), events_info, func(Reply *rpc.Reply_XdgmUpdateNormalEventInfo, err error) {
				if err != nil {
					log.Errorf("XDGM_UPDATE_EVENTS_INFO_REQ rpc error:", err)
					rsp.Status = -1
					rsp.Message = fmt.Sprintf("%v", err)
				} else {
					success = true
				}
			})
			if success {
				rsp.Status = 1
				rsp.Message = "Success"
			}
		}
	} else {
		rsp.Status = -1
		rsp.Message = "Sid Not Found"
	}
	return rsp, nil
}
