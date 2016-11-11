package xdgm_server

import (
	"core/log"
	"fmt"
	"xdgm_server/rpc"
)

func (this *XDGM_SEARCH_GAME_LAMP_REQ) Process(platid uint8, sid, cid int) (XDGM_RSP, error) {
	rsp := XDGM_RSP{}
	gs, exists := GetGsServerByHttp(platid, sid)
	if exists {
		hasHd := false
		for _, v := range gs {
			if v.HD == true {
				hasHd = true
				rpc.RemoteXdgmSearchGameLamp(int(v.GSID), func(Reply *rpc.Reply_XdgmSearchGameLamp, err error) {
					if err != nil {
						log.Errorf("XDGM_SEARCH_GAME_LAMP_REQ rpc error:", err)
						rsp.Status = -1
						rsp.Message = fmt.Sprintf("%v", err)
					} else {
						rsp.Status = 1
						rsp.Message = "success"
						Reply.Channel_id = cid
						Reply.Server_id = sid
						rsp.Data = Reply
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
