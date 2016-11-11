package idip_server

import (
	"platform_server"
	"time"
)

func (req *IDIP_DO_UPDATE_NOTICE_REQ) Process() (IDIP_DO_UPDATE_NOTICE_RSP, error) {
	st := &platform_server.ServerType{AreaId: uint8(req.Body.AreaId), PlatId: req.Body.PlatId}
	rsp := IDIP_DO_UPDATE_NOTICE_RSP{}
	rsp.Head.Cmdid = req.Head.Cmdid + 1
	rsp.Body.Result = 2
	rsp.Body.RetMsg = ""
	isSuccess := platform_server.SetAnnounce(st.GetType(), "xxd_qq", req.Body.Content, req.Body.Title, time.Unix(int64(req.Body.Time), 0).Format("2006-01-02 15:04:05"))
	if isSuccess {
		rsp.Body.Result = 0
		rsp.Body.RetMsg = "success"
	}
	return rsp, nil
}
