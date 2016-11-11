package idip_server

import (
	"fmt"
	"idip_server/rpc"
	"net/url"
	"platform_server"
)

func (req *IDIP_DO_SEND_MAIL_REQ) Process() (IDIP_DO_SEND_MAIL_RSP, error) {
	st := &platform_server.ServerType{AreaId: uint8(req.Body.AreaId), PlatId: req.Body.PlatId}
	rsp := IDIP_DO_SEND_MAIL_RSP{}
	rsp.Head.Cmdid = req.Head.Cmdid + 1
	rsp.Body.Result = 2
	rsp.Body.RetMsg = ""
	itemdetail := rpc.MailItem{
		ItemId1:  req.Body.ItemId1,
		ItemNum1: req.Body.ItemNum1,
		ItemId2:  req.Body.ItemId2,
		ItemNum2: req.Body.ItemNum2,
		ItemId3:  req.Body.ItemId3,
		ItemNum3: req.Body.ItemNum3,
		ItemId4:  req.Body.ItemId4,
		ItemNum4: req.Body.ItemNum4,
	}
	mailtitle, _ := url.QueryUnescape(req.Body.MailTitle)
	mailcontent, _ := url.QueryUnescape(req.Body.MailContent)
	if req.Body.OpenId == "" {
		if req.Body.Partition == 0 {
			server_list := platform_server.ServerList("xxd_qq")
			hasHd := false
			for _, v := range server_list {
				if v.Type == st.GetType() {
					for _, v2 := range v.GServers {
						if v2.HD == true {
							hasHd = true
							rpc.RemoteIdipSendMail(req.Body.OpenId, int(v2.GSId), mailtitle, mailcontent, req.Body.SendTime, req.Body.EndTime, itemdetail, func(Reply *rpc.Reply_IdipSendMail, err error) {
								if err != nil {
									rsp.Body.Result = 2
									rsp.Body.RetMsg = fmt.Sprintf("%v", err)
									rsp.Head.Result = -101
									rsp.Head.RetErrMsg = fmt.Sprintf("%v", err)
								} else {
									rsp.Body.Result = Reply.Result
									rsp.Body.RetMsg = Reply.RetMsg
								}
							})
							break
						}
					}
				}
			}
			if !hasHd {
				rsp.Body.Result = -4000
				rsp.Body.RetMsg = "Hd server not found"
				rsp.Head.Result = -4000
				rsp.Head.RetErrMsg = "Hd server not found"
			}
		} else {
			gs, exist := platform_server.GetServerInfo(int32(req.Body.Partition), st.GetType(), "xxd_qq")
			if exist {
				hasHd := false
				for _, v := range gs.GServers {
					if v.HD == true {
						hasHd = true
						rpc.RemoteIdipSendMail(req.Body.OpenId, int(v.GSId), mailtitle, mailcontent, req.Body.SendTime, req.Body.EndTime, itemdetail, func(Reply *rpc.Reply_IdipSendMail, err error) {
							if err != nil {
								rsp.Body.Result = 2
								rsp.Body.RetMsg = fmt.Sprintf("%v", err)
								rsp.Head.Result = -101
								rsp.Head.RetErrMsg = fmt.Sprintf("%v", err)
							} else {
								rsp.Body.Result = Reply.Result
								rsp.Body.RetMsg = Reply.RetMsg
							}
						})
						break
					}
				}
				if !hasHd {
					rsp.Body.Result = -4000
					rsp.Body.RetMsg = "Hd server not found"
					rsp.Head.Result = -4000
					rsp.Head.RetErrMsg = "Hd server not found"
				}
			} else {
				rsp.Body.Result = -4000
				rsp.Body.RetMsg = "area|partition|platid at least one error"
				rsp.Head.Result = -4000
				rsp.Head.RetErrMsg = "area|partition|platid at least one error"
			}
		}
		return rsp, nil
	}
	gs, exist := platform_server.GetGServerInfoByOpenIdSid(req.Body.OpenId, st.GetType(), int32(req.Body.Partition), "xxd_qq")
	if exist {
		rpc.RemoteIdipSendMail(req.Body.OpenId, int(gs.GSId), mailtitle, mailcontent, req.Body.SendTime, req.Body.EndTime, itemdetail, func(Reply *rpc.Reply_IdipSendMail, err error) {
			if err != nil {
				rsp.Body.Result = 2
				rsp.Body.RetMsg = fmt.Sprint("%v", err)
				rsp.Head.Result = -101
				rsp.Head.RetErrMsg = fmt.Sprint("%v", err)
			} else {
				if Reply.ErrMsg == "query empty" {
					rsp.Body.Result = 1
					rsp.Body.RetMsg = "query empty"
					rsp.Head.Result = 1
					rsp.Head.RetErrMsg = "query empty"
				} else {
					rsp.Body.Result = Reply.Result
					rsp.Body.RetMsg = Reply.RetMsg
				}
			}
		})
	} else {
		rsp.Body.Result = 1
		rsp.Body.RetMsg = "query empty"
		rsp.Head.Result = 1
		rsp.Head.RetErrMsg = "query empty"
	}
	return rsp, nil
}

func HashServerKey(Sid int32, iType uint8) string {
	return fmt.Sprintf("%08v_%03v", Sid, iType)
}
