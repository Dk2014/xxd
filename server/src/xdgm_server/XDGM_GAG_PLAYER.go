package xdgm_server

import (
	"core/log"
	"encoding/json"
	"xdgm_server/rpc"
)

func (this *XDGM_GAG_PLAYER_REQ) Process() (XDGM_RSP, error) {
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
		rpc.RemoteXdgmGagPlyaer(v, Pid2GameServerId(v), this.BlockTime, func(Reply *rpc.Reply_XdgmGagPlayer, err error) {
			if err != nil {
				failpids = append(failpids, v)
				log.Errorf("XDGM_GAG_PLAYER_REQ rpc error:", err)
			}
		})
	}
	rsp.Status = 1
	rsp.Message = "success"
	rsp.Data = failpids
	return rsp, nil
}
