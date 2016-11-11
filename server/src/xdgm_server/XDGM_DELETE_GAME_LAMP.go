package xdgm_server

import (
	"core/log"
	"fmt"
	"xdgm_server/rpc"
)

func (this *XDGM_DEL_GAME_LAMP_REQ) Process(platid uint8, sid int) (XDGM_RSP, error) {
	rsp := XDGM_RSP{}
	gs, exists := GetGsServerByHttp(platid, sid)
	if exists {
		hasHd := false
		for _, v := range gs {
			if v.HD == true {
				hasHd = true
				rpc.RemoteXdgmDelGameLamp(int(v.GSID), this.NoticeId, func(Reply *rpc.Reply_XdgmDelGameLamp, err error) {
					if err != nil {
						rsp.Status = -1
						rsp.Message = fmt.Sprintf("%v", err)
						log.Errorf("XDGM_DEL_GAME_LAMP_REQ rpc error:", err)
					} else {
						rsp.Status = 1
						rsp.Message = "success"
					}
				})
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
