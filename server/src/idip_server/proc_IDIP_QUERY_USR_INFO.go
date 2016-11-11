package idip_server

import (
	"fmt"
	"idip_server/rpc"
	"platform_server"
)

func (req *IDIP_QUERY_USR_INFO_REQ) Process() (IDIP_QUERY_USR_INFO_RSP, error) {
	st := &platform_server.ServerType{AreaId: uint8(req.Body.AreaId), PlatId: req.Body.PlatId}
	rsp := IDIP_QUERY_USR_INFO_RSP{}
	rsp.Head.Cmdid = req.Head.Cmdid + 1
	gs, exist := platform_server.GetGServerInfoByOpenIdSid(req.Body.OpenId, st.GetType(), int32(req.Body.Partition), "xxd_qq")
	if exist {
		rpc.RemoteIdipGetUserInfo(req.Body.OpenId, int(gs.GSId), func(Reply *rpc.Reply_IdipGetUserinfo, err error) {
			if err != nil {
				rsp.Head.Result = -101
				rsp.Head.RetErrMsg = fmt.Sprintf("%v", err)
			}
			if Reply.ErrMsg == "query empty" {
				rsp.Head.Result = 1
				rsp.Head.RetErrMsg = "query empty"
			}
			rsp.Body.Level = Reply.Level
			rsp.Body.Vip = Reply.Vip
			rsp.Body.Exp = Reply.Exp
			rsp.Body.Coin = Reply.Coin
			rsp.Body.Gold = Reply.Gold
			rsp.Body.Physical = Reply.Physical
			rsp.Body.MaxPhysical = Reply.MaxPhysical
			rsp.Body.MaxBag = Reply.MaxBag
			rsp.Body.RegisterTime = Reply.RegisterTime
			rsp.Body.IsOnline = Reply.IsOnline
			rsp.Body.AccStatus = Reply.AccStatus
			rsp.Body.BanEndTime = Reply.BanEndTime
			rsp.Body.ArmyId = Reply.ArmyId
			rsp.Body.RankInArmy = Reply.RankInArmy
			rsp.Body.ArmyRank = Reply.ArmyRank
			rsp.Body.PassProgress = Reply.PassProgress
			rsp.Body.PvpRank = Reply.PvpRank
			rsp.Body.PvpScore = Reply.PvpScore
			rsp.Body.LastLoginTime = Reply.LastLoginTime
			rsp.Body.RoleName = Reply.RoleName
		})
	} else {
		rsp.Head.Result = 1
		rsp.Head.RetErrMsg = "query empty"
	}
	return rsp, nil
}
