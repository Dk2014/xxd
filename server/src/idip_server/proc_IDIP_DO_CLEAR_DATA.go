package idip_server

import (
	"fmt"
	"idip_server/rpc"
	"platform_server"
)

func (req *IDIP_DO_CLEAR_DATA_REQ) Process() (IDIP_DO_CLEAR_DATA_RSP, error) {
	st := &platform_server.ServerType{AreaId: uint8(req.Body.AreaId), PlatId: req.Body.PlatId}
	rsp := IDIP_DO_CLEAR_DATA_RSP{}
	rsp.Head.Cmdid = req.Head.Cmdid + 1
	rsp.Body.Result = 2
	rsp.Body.RetMsg = ""
	gs, exist := platform_server.GetGServerInfoByOpenIdSid(req.Body.OpenId, st.GetType(), int32(req.Body.Partition), "xxd_qq")
	if exist {
		rpc.RemoteIdipCleanData(req.Body.OpenId, int(gs.GSId), uint8(req.Body.ClearCoin), uint8(req.Body.ClearHeart), uint8(req.Body.ClearPhysical), func(Reply *rpc.Reply_IdipCleanData, err error) {
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
