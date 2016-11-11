package xdgm_server

import (
	"core/log"
	"fmt"
	"xdgm_server/gift_code"
)

func (this *XDGM_QUERY_GIFT_CODE_REQ) Process() (XDGM_RSP, error) {
	rsp := XDGM_RSP{}

	data, err := gift_code.QueryCode(this.Version, this.ServerId, this.Offset, this.Limit)
	if err != nil {
		log.Errorf("XDGM_QUERY_GIFT_CODE_REQ rpc error:", err)
		rsp.Status = 0
		rsp.Message = fmt.Sprintf("%v", err)
		return rsp, err
	}

	rsp.Status = 1
	rsp.Message = "success"
	rsp.Data = data
	return rsp, nil
}
