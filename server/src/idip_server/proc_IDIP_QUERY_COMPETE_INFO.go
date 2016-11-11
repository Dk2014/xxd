package idip_server

import (
	"fmt"
	"idip_server/rpc"
	"platform_server"
	"strconv"
)

func (req *IDIP_QUERY_COMPETE_INFO_REQ) Process() (IDIP_QUERY_COMPETE_INFO_RSP, error) {
	st := &platform_server.ServerType{AreaId: uint8(req.Body.AreaId), PlatId: req.Body.PlatId}
	rsp := IDIP_QUERY_COMPETE_INFO_RSP{}
	rsp.Head.Cmdid = req.Head.Cmdid + 1
	var roleid int8
	var pid int64
	gs, exist := platform_server.GetGServerInfoByOpenIdSid(req.Body.OpenId, st.GetType(), int32(req.Body.Partition), "xxd_qq")
	if exist {
		rpc.RemoteIdipGetRankInfoGs(req.Body.OpenId, int(gs.GSId), func(Reply *rpc.Reply_IdipGetRankinfoGs, err error) {
			if err != nil {
				rsp.Head.Result = -101
				rsp.Head.RetErrMsg = fmt.Sprintf("%v", err)
				return
			}
			if Reply.ErrMsg != "" && Reply.ErrMsg == "query empty" {
				rsp.Head.Result = 1
				rsp.Head.RetErrMsg = "query empty"
				return
			} else if Reply.ErrMsg != "" {
				rsp.Head.Result = 1
				rsp.Head.RetErrMsg = Reply.ErrMsg
				return
			}
			pid = Reply.Pid
			roleid = Reply.RoleId
		})
	} else {
		rsp.Head.Result = 1
		rsp.Head.RetErrMsg = "query empty"
	}
	if pid > 0 && roleid > 0 {
		gs, exist := platform_server.GetServerInfo(int32(req.Body.Partition), st.GetType(), "xxd_qq")
		if exist {
			hasHd := false
			for _, v := range gs.GServers {
				if v.HD == true {
					hasHd = true
					rpc.RemoteIdipGetRankInfoHd(pid, int(v.GSId), func(Reply *rpc.Reply_IdipGetRankinfoHd, err error) {
						if err != nil {
							rsp.Head.Result = -101
							rsp.Head.RetErrMsg = fmt.Sprintf("%v", err)
							return
						}
						if Reply.ErrMsg != "" {
							rsp.Head.Result = 1
							rsp.Head.RetErrMsg = Reply.ErrMsg
							return
						}
						rsp.Body.Rank = Reply.Rank
						rsp.Body.RoleId = strconv.Itoa(int(roleid))
					})
					break
				}
			}
			if !hasHd {
				rsp.Head.Result = -4000
				rsp.Head.RetErrMsg = "Hd server not found"
			}
		} else {
			rsp.Head.Result = 1
			rsp.Head.RetErrMsg = "query empty"
		}
	}
	return rsp, nil
}
