package idip_server

import (
	"fmt"
	"idip_server/rpc"
	"platform_server"
)

func (req *IDIP_QUERY_GAME_LAMP_NOTICE_REQ) Process() (IDIP_QUERY_GAME_LAMP_NOTICE_RSP, error) {
	st := &platform_server.ServerType{AreaId: uint8(req.Body.AreaId), PlatId: req.Body.PlatId}
	rsp := IDIP_QUERY_GAME_LAMP_NOTICE_RSP{}
	rsp.Head.Cmdid = req.Head.Cmdid + 1
	gs, exist := platform_server.GetServerInfo(int32(req.Body.Partition), st.GetType(), "xxd_qq")
	if exist {
		hasHd := false
		for _, v := range gs.GServers {
			if v.HD == true {
				hasHd = true
				rpc.RemoteIdipGameLampInfo(int(v.GSId), func(Reply *rpc.Reply_IdipGameLampInfo, err error) {
					if err != nil {
						rsp.Head.Result = -101
						rsp.Head.RetErrMsg = fmt.Sprintf("%v", err)
					}
					if Reply.GameLampNoticeList_count == 0 {
						rsp.Head.Result = 1
						rsp.Head.RetErrMsg = "query empty"
						return
					}
					rsp.Body.GameLampNoticeList_count = Reply.GameLampNoticeList_count
					for _, v := range Reply.GameLampNoticeList {
						rsp.Body.GameLampNoticeList = append(rsp.Body.GameLampNoticeList, SGameLampNoticeInfo{
							BeginTime:     v.BeginTime,
							EndTime:       v.EndTime,
							Freq:          v.Freq,
							NoticeId:      v.NoticeId,
							NoticeContent: v.NoticeContent,
						})
					}
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
	return rsp, nil
}
