package idip_server

import (
	"fmt"
	"idip_server/rpc"
	"platform_server"
)

func (req *IDIP_QUERY_BAG_INFO_REQ) Process() (IDIP_QUERY_BAG_INFO_RSP, error) {
	st := &platform_server.ServerType{AreaId: uint8(req.Body.AreaId), PlatId: req.Body.PlatId}
	rsp := IDIP_QUERY_BAG_INFO_RSP{}
	rsp.Head.Cmdid = req.Head.Cmdid + 1
	gs, exist := platform_server.GetGServerInfoByOpenIdSid(req.Body.OpenId, st.GetType(), int32(req.Body.Partition), "xxd_qq")
	if exist {
		rpc.RemoteIdipGetBagInfo(req.Body.OpenId, int(gs.GSId), uint64(req.Body.BeginTime), uint64(req.Body.EndTime), func(Reply *rpc.Reply_IdipGetBaginfo, err error) {
			if err != nil {
				rsp.Head.Result = -101
				rsp.Head.RetErrMsg = fmt.Sprintf("%v", err)
			}
			if Reply.ErrMsg == "query empty" {
				rsp.Head.Result = 1
				rsp.Head.RetErrMsg = "query empty"
			}
			if Reply.BagList_count == 0 {
				rsp.Head.Result = 1
				rsp.Head.RetErrMsg = "query empty"
			}
			rsp.Body.BagList_count = Reply.BagList_count
			for _, v := range Reply.BagList {
				rsp.Body.BagList = append(rsp.Body.BagList, SBagInfo{
					ItemName: v.ItemName,
					ItemId:   v.ItemId,
					ItemNum:  v.ItemNum,
				})
			}
		})
	} else {
		rsp.Head.Result = 1
		rsp.Head.RetErrMsg = "query empty"
	}
	return rsp, nil
}
