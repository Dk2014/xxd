package xdgm_server

import (
	"core/log"
	"fmt"
	"xdgm_server/rpc"
)

func (this *XDGM_UPDATE_EVENT_AWARDS_REQ) Process(platid uint8, sid int, event_id int16, event_awards string) (XDGM_RSP, error) {
	rsp := XDGM_RSP{}
	gs, exists := GetGsServerByHttp(platid, sid)
	if exists {
		for _, v := range gs {
			var success bool
			rpc.XdgmUpdateEventAwards(int(v.GSID), event_id, event_awards, func(Reply *rpc.Reply_XdgmUpdateEventAwards, err error) {
				if err != nil {
					log.Errorf("XDGM_UPDATE_EVENT_AWARDS_REQ rpc error:", err)
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
