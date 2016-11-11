package idip_server

import (
	"fmt"
	"idip_server/rpc"
	"platform_server"
)

func (req *IDIP_QUERY_ROLE_PET_INFO_REQ) Process() (IDIP_QUERY_ROLE_PET_INFO_RSP, error) {
	st := &platform_server.ServerType{AreaId: uint8(req.Body.AreaId), PlatId: req.Body.PlatId}
	rsp := IDIP_QUERY_ROLE_PET_INFO_RSP{}
	rsp.Head.Cmdid = req.Head.Cmdid + 1
	gs, exist := platform_server.GetGServerInfoByOpenIdSid(req.Body.OpenId, st.GetType(), int32(req.Body.Partition), "xxd_qq")
	if exist {
		rpc.RemoteIdipGetPetInfo(req.Body.OpenId, int(gs.GSId), int8(req.Body.RoleId), func(Reply *rpc.Reply_IdipGetPetinfo, err error) {
			if err != nil {
				rsp.Head.Result = -101
				rsp.Head.RetErrMsg = fmt.Sprintf("%v", err)
			}
			if Reply.ErrMsg == "query empty" {
				rsp.Head.Result = 1
				rsp.Head.RetErrMsg = "query empty"
			}
			if Reply.RolePetList_count == 0 {
				rsp.Head.Result = 1
				rsp.Head.RetErrMsg = "query empty"
			}
			rsp.Body.RolePetList_count = Reply.RolePetList_count
			for _, v := range Reply.RolePetList {
				rsp.Body.RolePetList = append(rsp.Body.RolePetList, SRolePetInfo{
					PetId:    v.PetId,
					PetName:  v.PetName,
					Level:    v.Level,
					IsBattle: v.IsBattle,
				})
			}
		})
	} else {
		rsp.Head.Result = 1
		rsp.Head.RetErrMsg = "query empty"
	}
	return rsp, nil
}
