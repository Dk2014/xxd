package idip_server

import (
	"fmt"
	"idip_server/obj"
	"idip_server/rpc"
	"platform_server"
)

func (req *IDIP_DO_UPDATE_VIP_REQ) Process() (IDIP_DO_UPDATE_VIP_RSP, error) {
	st := &platform_server.ServerType{AreaId: uint8(req.Body.AreaId), PlatId: req.Body.PlatId}
	rsp := IDIP_DO_UPDATE_VIP_RSP{}
	rsp.Head.Cmdid = req.Head.Cmdid + 1
	gs, exist := platform_server.GetGServerInfoByOpenIdSid(req.Body.OpenId, st.GetType(), int32(req.Body.Partition), "xxd_qq")
	if exist {
		var data = make(map[string]interface{})
		dataTotal := obj.Getdatamap(req.Body, data)
		rpc.RemoteIdipAddVipExp(dataTotal, req.Head.Cmdid, int(gs.GSId), func(Reply *rpc.Reply_IdipAddVipExp, err error) {
			if err != nil {
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
