package idip_server

import (
	"fmt"
	"idip_server/rpc"
	"platform_server"
)

func (req *IDIP_QUERY_ABYSSPASS_STATE_REQ) Process() (IDIP_QUERY_ABYSSPASS_STATE_RSP, error) {
	st := &platform_server.ServerType{AreaId: uint8(req.Body.AreaId), PlatId: req.Body.PlatId}
	rsp := IDIP_QUERY_ABYSSPASS_STATE_RSP{}
	rsp.Head.Cmdid = req.Head.Cmdid + 1

	gs, exist := platform_server.GetGServerInfoByOpenIdSid(req.Body.OpenId, st.GetType(), int32(req.Body.Partition), "xxd_qq")
	if exist {
		rpc.RemoteIdipGetHardLevelStatus(req.Body.OpenId, int(gs.GSId), func(Reply *rpc.Reply_IdipGetHardLevelStatus, err error) {
			if err != nil {
				rsp.Head.Result = -101
				rsp.Head.RetErrMsg = fmt.Sprintf("%v", err)
			}
			if Reply.ErrMsg == "query empty" {
				rsp.Head.Result = 1
				rsp.Head.RetErrMsg = "query empty"
			}
			rsp.Body.Status = Reply.Status
		})
	} else {
		rsp.Head.Result = 1
		rsp.Head.RetErrMsg = "query empty"
	}
	return rsp, nil
}
