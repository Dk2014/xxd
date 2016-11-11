package xdgm_server

import (
	"core/log"
	"fmt"
	"xdgm_server/gift_code"
	"xdgm_server/rpc"
)

func (this *XDGM_CANCEL_GIFT_CODE_REQ) Process() (XDGM_RSP, error) {
	rsp := XDGM_RSP{}

	err := gift_code.CancelCode(this.Version, this.ServerId)
	if err != nil {
		rsp.Status = 0
		rsp.Message = fmt.Sprintf("%v", err)
		return rsp, err
	}

	// 取消完之后rpc通知相应服务器重新加载激活码数据
	defer func() {
		servers := GetAllGsServerByHttp(0)
		for _, server := range servers {
			if server.HD == true {
				rpc.RemoteXdgmReloadGiftCode(server.GSID, func(*rpc.Reply_XdgmReloadGiftCode, error) {
					if err != nil {
						log.Errorf("XDGM_CANCEL_GIFT_CODE_REQ rpc error:", err)
					}
				})
			}
		}
	}()
	rsp.Status = 1
	rsp.Message = "success"
	//rsp.Data = failpids
	return rsp, nil
}
