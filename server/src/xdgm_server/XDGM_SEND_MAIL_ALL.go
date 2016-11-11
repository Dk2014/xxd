package xdgm_server

import (
	"core/log"
	"fmt"
	"xdgm_server/rpc"
)

func (this *XDGM_SEND_MAIL_ALL_REQ) Process(sid int) (XDGM_RSP, error) {
	var success bool
	rsp := XDGM_RSP{}
	rpc.RemoteXdgmSendMailAll(Sid2GlobalServerId(sid), this.Title, this.Content, this.Attach, this.BeginTime, this.Endtime, this.MinLevel, this.MaxLevel, this.MinVipLevel, this.MaxVipLevel, func(Reply *rpc.Reply_XdgmSendMailAll, err error) {
		if err != nil {
			log.Errorf("XDGM_SEND_MAIL_ALL_REQ rpc error:", err)
			rsp.Status = -1
			rsp.Message = fmt.Sprintf("%v", err)
		} else {
			success = true
		}
		if success {
			rsp.Status = 1
			rsp.Message = "Success"
		}
	})
	return rsp, nil
}
