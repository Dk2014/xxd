package idip_server

import (
	"fmt"
	"idip_server/rpc"
	"platform_server"
)

func (req *IDIP_AQ_DO_SEND_MAIL_REQ) Process() (IDIP_AQ_DO_SEND_MAIL_RSP, error) {
	st := &platform_server.ServerType{AreaId: uint8(req.Body.AreaId), PlatId: req.Body.PlatId}
	rsp := IDIP_AQ_DO_SEND_MAIL_RSP{}
	rsp.Head.Cmdid = req.Head.Cmdid + 1
	rsp.Body.Result = 2
	rsp.Body.RetMsg = ""
	gs, exist := platform_server.GetGServerInfoByOpenIdSid(req.Body.OpenId, st.GetType(), int32(req.Body.Partition), "xxd_qq")
	if exist {
		rpc.RemoteIdipAqSendMail(req.Body.OpenId, int(gs.GSId), req.Body.MailContent, func(Reply *rpc.Reply_IdipAqSendMail, err error) {
			if err != nil {
				rsp.Body.Result = 2
				rsp.Body.RetMsg = fmt.Sprintf("%v", err)
				rsp.Head.Result = -101
				rsp.Head.RetErrMsg = fmt.Sprintf("%v", err)
			} else {
				if Reply.ErrMsg == "query empty" {
					rsp.Body.Result = 1
					rsp.Body.RetMsg = "query empty"
					rsp.Head.Result = 1
					rsp.Head.RetErrMsg = "query empty"
				} else {
					rsp.Body.Result = Reply.Result
					rsp.Body.RetMsg = Reply.RetMsg
				}
			}
		})
	} else {
		rsp.Body.Result = 1
		rsp.Body.RetMsg = "query empty"
		rsp.Head.Result = 1
		rsp.Head.RetErrMsg = "query empty"
	}
	return rsp, nil
}
