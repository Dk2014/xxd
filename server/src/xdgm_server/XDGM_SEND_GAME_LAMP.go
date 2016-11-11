package xdgm_server

import (
	"core/log"
	"fmt"
	"xdgm_server/rpc"
)

func (this *XDGM_SEND_GAME_LAMP_REQ) Process(platid uint8, sid int) (XDGM_RSP, error) {
	rsp := XDGM_RSP{}
	gs, exists := GetGsServerByHttp(platid, sid)
	var hasHd bool
	if exists {
		for _, v := range gs {
			if v.HD == true {
				hasHd = true
				var success bool
				rpc.RemoteXdgmSendGameLamp(int(v.GSID), this.Content, this.BeginTime, this.EndTime, this.Interval, func(Reply *rpc.Reply_XdgmSendGameLamp, err error) {
					if err != nil {
						log.Errorf("XDGM_SEND_GAME_LAMP_REQ rpc error:", err)
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
				break
			}
		}
		if !hasHd {
			rsp.Status = -1
			rsp.Message = "Hd Server Not Found"
		}
	} else {
		rsp.Status = -1
		rsp.Message = "Sid Not Found"
	}
	return rsp, nil
}
