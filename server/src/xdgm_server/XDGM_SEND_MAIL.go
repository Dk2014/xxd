package xdgm_server

import (
	"core/log"
	"encoding/json"
	"xdgm_server/rpc"
)

func (this *XDGM_SEND_MAIL_REQ) Process() (XDGM_RSP, error) {
	rsp := XDGM_RSP{}
	pids := make([]int64, 0)
	failpids := make([]int64, 0)
	err := json.Unmarshal([]byte(this.Players), &pids)
	if err != nil {
		rsp.Status = -1
		rsp.Message = "players json wrong"
		return rsp, err
	}
	for _, v := range pids {
		rpc.RemoteXdgmSendMail(v, Pid2GameServerId(v), this.Title, this.Content, this.Attach, this.Endtime, func(Reply *rpc.Reply_XdgmSendMail, err error) {
			if err != nil {
				log.Errorf("XDGM_SEND_MAIL_REQ rpc error:", err)
				failpids = append(failpids, v)
			}
		})
	}
	rsp.Status = 1
	rsp.Message = "success"
	rsp.Data = failpids
	return rsp, nil
}
