package xdgm_server

import (
	"core/log"
	"net/http"
	"strconv"
	"xdgm_server/rpc"
)

func (this *XDGM_GET_JSON_EVENT_INFO_REQ) Process() (rsp XDGM_RSP, err error) {
	rpc.XdgmGetJsonEventInfo(Sid2GlobalServerId(int(this.ServerId)), this.Offset, this.Limit, func(reply *rpc.Reply_XdgmGetJsonEventInfo, err1 error) {
		if err1 != nil {
			log.Errorf("XDGM_GET_JSON_EVENT_INFO_REQ rpc error:", err1)
			err = err1
			return
		}
		rsp.Message = "success"
		rsp.Status = 1
		rsp.Data = reply
	})
	if err != nil {
		rsp.Message = err.Error()
		rsp.Status = -1
	}
	return rsp, err
}

func (this *XDGM_GET_JSON_EVENT_INFO_REQ) Load(req *http.Request) (err error) {
	server_id, err := strconv.ParseInt(req.FormValue("server_id"), 10, 32)
	if err != nil {
		return err
	}
	this.ServerId = int(server_id)

	limit, err := strconv.ParseInt(req.PostFormValue("limit"), 10, 16)
	if err != nil {
		return err
	}
	this.Limit = int16(limit)
	offset, err := strconv.ParseInt(req.PostFormValue("offset"), 10, 16)
	if err != nil {
		return err
	}
	this.Offset = int16(offset)
	return nil
}
