package xdgm_server

import (
	"core/log"
	"net/http"
	"strconv"
	"xdgm_server/rpc"
)

func (this *XDGM_EVENT_AWARD_INFO_REQ) Process() (rsp XDGM_RSP, err error) {
	rpc.XdgmGetEventAwardInfo(Sid2GlobalServerId(int(this.ServerId)), this.EventSign, this.Page, func(reply *rpc.Reply_XdgmGetEventAwardInfo, err1 error) {
		if err1 != nil {
			log.Errorf("XDGM_EVENT_AWARD_INFO_REQ rpc error:", err1)
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

func (this *XDGM_EVENT_AWARD_INFO_REQ) Load(req *http.Request) (err error) {
	server_id, err := strconv.ParseInt(req.FormValue("server_id"), 10, 32)
	if err != nil {
		return err
	}
	this.ServerId = int(server_id)

	sign, err := strconv.ParseInt(req.PostFormValue("event_sign"), 10, 16)
	if err != nil {
		return err
	}
	this.EventSign = int16(sign)
	page, err := strconv.ParseInt(req.PostFormValue("page"), 10, 32)
	if err != nil {
		return err
	}
	this.Page = int32(page)
	return nil
}
