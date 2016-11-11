package idip_server

import (
	"fmt"
	"idip_server/rpc"
	"platform_server"
)

func (req *IDIP_QUERY_SOUL_INFO_REQ) Process() (IDIP_QUERY_SOUL_INFO_RSP, error) {
	st := &platform_server.ServerType{AreaId: uint8(req.Body.AreaId), PlatId: req.Body.PlatId}
	rsp := IDIP_QUERY_SOUL_INFO_RSP{}
	rsp.Head.Cmdid = req.Head.Cmdid + 1
	gs, exist := platform_server.GetGServerInfoByOpenIdSid(req.Body.OpenId, st.GetType(), int32(req.Body.Partition), "xxd_qq")
	if exist {
		rpc.RemoteIdipGetSoulInfo(req.Body.OpenId, int(gs.GSId), int16(req.Body.RoleId), func(Reply *rpc.Reply_IdipGetSoulinfo, err error) {
			if err != nil {
				rsp.Head.Result = -101
				rsp.Head.RetErrMsg = fmt.Sprintf("%v", err)
			}
			if Reply.ErrMsg == "query empty" {
				rsp.Head.Result = 1
				rsp.Head.RetErrMsg = "query empty"
			}
			if Reply.SoulList_count == 0 {
				rsp.Head.Result = 1
				rsp.Head.RetErrMsg = "query empty"
			}
			rsp.Body.SoulList_count = Reply.SoulList_count
			for _, v := range Reply.SoulList {
				rsp.Body.SoulList = append(rsp.Body.SoulList, SSoulInfo{
					Slot:            v.Slot,
					RoleName:        v.RoleName,
					MajorSoul:       v.MajorSoul,
					MajorSoulLevel:  v.MajorSoulLevel,
					Minor1Soul:      v.Minor1Soul,
					MinorSoul1Level: v.MinorSoul1Level,
					Minor2Soul:      v.Minor2Soul,
					MinorSoul3Level: v.MinorSoul3Level,
					Minor3Soul:      v.Minor3Soul,
					MinorSoul2Level: v.MinorSoul2Level,
					IsBattle:        v.IsBattle,
				})
			}
		})
	} else {
		rsp.Head.Result = 1
		rsp.Head.RetErrMsg = "query empty"
	}
	return rsp, nil
}
