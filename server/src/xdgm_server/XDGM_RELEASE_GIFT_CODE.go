package xdgm_server

import (
	"core/log"
	"fmt"
	"xdgm_server/rpc"
)

func (this *XDGM_RELEASE_GIFT_CODE_REQ) Process(platid uint8, sid int) (XDGM_RSP, error) {
	rsp := XDGM_RSP{}
	var found bool
	var err error
	if gs, exists := GetGsServerByHttp(platid, sid); exists {
		for _, cfg := range gs {
			if cfg.HD {
				rpc.RemoteXdgmReloadGiftCode(int(cfg.GSID), func(reply *rpc.Reply_XdgmReloadGiftCode, err1 error) {
					if err1 != nil {
						log.Errorf("XDGM_RELEASE_GIFT_CODE_REQ rpc error:", err1)
						err = err1
						rsp.Status = -1
						rsp.Message = fmt.Sprintf("%v", err)
					}
				})
				break
			}
		}
	}
	if err != nil {
		return rsp, err
	}
	if !found {
		rsp.Status = -1
		rsp.Message = "not server config"
		return rsp, nil
	}

	rsp.Status = 1
	rsp.Message = "success"
	//rsp.Data = failpids
	return rsp, nil
}
