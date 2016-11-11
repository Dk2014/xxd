package xdgm_server

import (
	"core/log"
	"net/http"
	"strconv"
	"xdgm_server/rpc"
)

func (this *XDGM_GET_TEXT_EVENTS_REQ) Load(req *http.Request) (err error) {
	server_id, err := strconv.ParseInt(req.FormValue("server_id"), 10, 32)
	if err != nil {
		return err
	}
	this.ServerId = int(server_id)
	return nil
}

func (this *XDGM_GET_TEXT_EVENTS_REQ) Process() (rsp XDGM_RSP, err error) {
	rpc.XdgmGetTextEvents(Sid2GlobalServerId(this.ServerId), func(reply *rpc.Reply_XdgmGetTextEvents, err1 error) {
		if err1 != nil {
			log.Errorf("XDGM_GET_TEXT_EVENTS_REQ rpc error:", err1)
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

	return
}
