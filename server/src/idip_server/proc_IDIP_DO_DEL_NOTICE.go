package idip_server

import (
	"fmt"
	"idip_server/rpc"
	"platform_server"
)

func (req *IDIP_DO_DEL_NOTICE_REQ) Process() (IDIP_DO_DEL_NOTICE_RSP, error) {
	st := &platform_server.ServerType{AreaId: uint8(req.Body.AreaId), PlatId: req.Body.PlatId}
	rsp := IDIP_DO_DEL_NOTICE_RSP{}
	rsp.Head.Cmdid = req.Head.Cmdid + 1
	gs, exist := platform_server.GetServerInfo(int32(req.Body.Partition), st.GetType(), "xxd_qq")
	var hasRpc = false
	if exist {
		for _, v := range gs.GServers {
			if v.HD == true {
				hasRpc = true
				rpc.RemoteIdipDelNotice(int(v.GSId), int64(req.Body.NoticeId), func(Reply *rpc.Reply_IdipDelNotice, err error) {
					if err != nil {
						rsp.Body.Result = 2
						rsp.Body.RetMsg = fmt.Sprintf("%v", err)
						rsp.Head.Result = -101
						rsp.Head.RetErrMsg = fmt.Sprintf("%v", err)
					} else {
						rsp.Body.Result = Reply.Result
						rsp.Body.RetMsg = Reply.RetMsg
					}
				})
				break
			}
		}
		if !hasRpc {
			rsp.Body.Result = -4000
			rsp.Body.RetMsg = "Hd server not found"
			rsp.Head.Result = -4000
			rsp.Head.RetErrMsg = "Hd server not found"
		}
	} else {
		rsp.Body.Result = -4000
		rsp.Body.RetMsg = "area|partition|platid at least one error"
		rsp.Head.Result = -4000
		rsp.Head.RetErrMsg = "area|partition|platid at least one error"
	}
	return rsp, nil
}
