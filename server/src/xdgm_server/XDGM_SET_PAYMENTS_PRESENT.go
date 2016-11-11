package xdgm_server

import (
	"core/log"
	"encoding/json"
	"xdgm_server/rpc"
)

func (this *XDGM_SET_PAYMENTS_PRESENT_REQ) Process(platid uint8) (XDGM_RSP, error) {
	var global bool
	rsp := XDGM_RSP{}
	serverids := make([]int, 0)
	failserverids := make([]int, 0)
	err := json.Unmarshal([]byte(this.Servers), &serverids)
	if err != nil {
		rsp.Status = -1
		rsp.Message = "players json wrong"
		return rsp, err
	}
	for _, v := range serverids {
		if v == 0 {
			global = true
			break
		}
	}
	if global {
		serverMap := GetAllGsServerByHttp(platid)
		for _, gs := range serverMap {
			if gs.HD != true {
				rpc.RemoteXdgmSetPaymentsPresent(int(gs.GSID/10), this.Rule, this.BeginTime, this.EndTime, func(Reply *rpc.Reply_XdgmSetPaymentsPresent, err error) {
					if err != nil {
						log.Errorf("XDGM_SET_PAYMENTS_PRESENT_REQ rpc error:", err)
						failserverids = append(failserverids, int(gs.GSID/10))
					}
				})
			}
		}
	} else {
		for _, v := range serverids {
			server, ok := GetGsServerByHttp(platid, v)
			if ok {
				for _, gs := range server {
					if gs.HD != true {
						rpc.RemoteXdgmSetPaymentsPresent(int(gs.GSID), this.Rule, this.BeginTime, this.EndTime, func(Reply *rpc.Reply_XdgmSetPaymentsPresent, err error) {
							if err != nil {
								log.Errorf("XDGM_SET_PAYMENTS_PRESENT_REQ rpc error:", err)
								failserverids = append(failserverids, v)
							}
						})
					}
				}
			}
		}
	}
	rsp.Status = 1
	rsp.Message = "success"
	rsp.Data = failserverids
	return rsp, nil
}
