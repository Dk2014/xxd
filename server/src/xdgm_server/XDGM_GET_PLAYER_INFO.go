package xdgm_server

import (
	"core/log"
	"encoding/json"
	"fmt"
	"xdgm_server/rpc"
)

func (this *XDGM_GET_PLAYER_INFO_REQ) Process() (XDGM_RSP, error) {
	rsp := XDGM_RSP{}
	player := make(map[string]*XDGM_PLAYER_INFO)
	pids := make([]int64, 0)
	err := json.Unmarshal([]byte(this.Players), &pids)
	if err != nil {
		rsp.Status = -1
		rsp.Message = "players json wrong"
		return rsp, nil
	}
	for _, v := range pids {
		pid := fmt.Sprintf("%v", v)
		rpc.RemoteXdgmGetUserInfo(v, Pid2GameServerId(v), func(Reply *rpc.Reply_XdgmGetUserinfo, err error) {
			if err == nil {
				mails := make([]XDGM_MAIL, 0)
				if err == nil {
					for _, v := range Reply.Mails {
						itemdetail := make([]MailItem, 0)
						for _, v2 := range v.ItemDetail {
							itemdetail = append(itemdetail, MailItem{
								ItemId:  v2.ItemId,
								ItemNum: v2.ItemNum,
							})
						}
						mails = append(mails, XDGM_MAIL{
							MailTitle:   v.MailTitle,
							MailContent: v.MailContent,
							SendTime:    v.SendTime,
							ItemDetail:  itemdetail,
						})
					}
					player[pid] = &XDGM_PLAYER_INFO{
						RoleName:      Reply.RoleName,
						Pid:           Reply.Pid,
						Level:         Reply.Level,
						Vip:           Reply.Vip,
						Exp:           Reply.Exp,
						Coin:          Reply.Coin,
						Ingot:         Reply.Ingot,
						Physical:      Reply.Physical,
						RegisterTime:  Reply.RegisterTime,
						LastLoginTime: Reply.LastLoginTime,
						Mails:         mails,
						PlayerInfo:    Reply.PlayerInfo,
					}
				}
			} else {
				log.Errorf("XDGM_GET_PLAYER_INFO_REQ rpc error:", err)
			}
		})
		rpc.RemoteXdgmGetUserRank(v, Pid2GlobalServerId(v), func(Reply *rpc.Reply_XdgmGetUserRank, err error) {
			if err == nil {
				player[pid].Rank = Reply.Rank
			} else {
				log.Errorf("XDGM_GET_PLAYER_INFO_REQ Rank rpc error:", err)
			}
		})
		rsp.Status = 1
		rsp.Message = "success"
		rsp.Data = player
	}
	return rsp, nil
}
