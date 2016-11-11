package idip_server

import (
	"fmt"
	"idip_server/rpc"
	"platform_server"
)

func (req *IDIP_DO_UPDATE_GAME_LAMP_REQ) Process() (IDIP_DO_UPDATE_GAME_LAMP_RSP, error) {
	st := &platform_server.ServerType{AreaId: uint8(req.Body.AreaId), PlatId: req.Body.PlatId}
	rsp := IDIP_DO_UPDATE_GAME_LAMP_RSP{}
	rsp.Head.Cmdid = req.Head.Cmdid + 1
	rsp.Body.Result = 2
	rsp.Body.RetMsg = ""
	gs, exist := platform_server.GetServerInfo(int32(req.Body.Partition), st.GetType(), "xxd_qq")
	if exist {
		for _, v := range gs.GServers {
			rpc.RemoteIdipUpdateGameLamp(int(v.GSId), req.Body.LampContent, uint64(req.Body.BeginTime), uint64(req.Body.EndTime), int32(req.Body.Freq), func(Reply *rpc.Reply_IdipUpdateGameLamp, err error) {
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
	} else {
		rsp.Head.Result = -4000
		rsp.Head.RetErrMsg = "area|partition|platid at least one error"
		rsp.Body.Result = -4000
		rsp.Body.RetMsg = "area|partition|platid at least one error"
	}
	return rsp, nil
}
