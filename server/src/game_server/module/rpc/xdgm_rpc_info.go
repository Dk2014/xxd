package rpc

import (
	"core/fail"
	"encoding/json"
	"fmt"
	"game_server/api/protocol/role_api"
	"game_server/dat/item_dat"
	"game_server/dat/mail_dat"
	"game_server/mdb"
	"game_server/module"
	. "game_server/rpc"
	"strings"
)

/*
	玩家当前个人信息
*/
type FriendsInfo struct {
	Name       string
	Level      int16
	Exp        int64
	FriendShip int32
}

type MailsInfo struct {
	MailTitle   string // 邮件标题
	MailContent string // 邮件内容
	SendTime    int64  // 邮件发送时间
	ItemDetail  []MailItemxdgm
}

type MailItemxdgm struct {
	ItemId  int16 // 邮件赠送道具ID
	ItemNum int64 // 邮件赠送道具数量
}

type Args_XdgmGetUserinfo struct {
	RPCArgTag
	Pid int64
}

type Reply_XdgmGetUserinfo struct {
	RoleName      string                      // 角色名称
	Pid           int64                       // 玩家ID
	Level         int16                       // 当前等级
	Vip           int16                       // 当前VIP等级
	Exp           int64                       // 当前经验
	Coin          int64                       // 当前铜钱
	Ingot         int64                       // 当前元宝数量
	Physical      int16                       // 当前体力值
	RegisterTime  int64                       // 注册时间
	LastLoginTime int64                       // 玩家最后登录时间
	Friends       []FriendsInfo               //伙伴信息
	Mails         []MailsInfo                 //邮件信息
	PlayerInfo    []role_api.PlayerInfo_Roles //角色信息
}

func (this *RemoteServe) XdgmGetUserInfo(args *Args_XdgmGetUserinfo, reply *Reply_XdgmGetUserinfo) error {
	return Remote.Serve(mdb.RPC_Remote_XdgmGetUserInfo, args, mdb.TRANS_TAG_RPC_Serve_XdgmGetUserInfo, func() error {
		if ok := mdb.CheckPlayer(args.Pid); !ok {
			return nil
		}
		mdb.GlobalExecute(func(globalDB *mdb.Database) {

			globalDB.AgentExecute(args.Pid, func(db *mdb.Database) {
				reply.Pid = args.Pid
				//角色信息
				roleInfo := module.Role.GetMainRole(db)
				reply.Level = roleInfo.Level
				reply.Exp = roleInfo.Exp

				//vip信息
				playerVip := module.VIP.VIPInfo(db)
				reply.Vip = playerVip.Level

				//玩家信息
				playerInfo := db.Lookup.PlayerInfo(db.PlayerId())
				reply.Coin = playerInfo.Coins
				reply.Ingot = playerInfo.Ingot
				reply.RegisterTime = playerInfo.FirstLoginTime
				reply.LastLoginTime = playerInfo.LastLoginTime
				playerinfo2 := module.Player.GetPlayer(db)
				reply.RoleName = playerinfo2.Nick

				//体力信息
				playerPhysical := db.Lookup.PlayerPhysical(db.PlayerId())
				reply.Physical = playerPhysical.Value

				// //玩家伙伴信息
				// db.Select.PlayerRole(func(row *mdb.PlayerRoleRow) {
				// 	if row.RoleId() != 1 && row.RoleId() != 2 {
				// 		role_info := role_dat.GetRoleInfo(int8(row.RoleId()))
				// 		reply.Friends = append(reply.Friends, FriendsInfo{
				// 			Name:       role_info.Name,
				// 			Level:      row.Level(),
				// 			Exp:        row.Exp(),
				// 			FriendShip: row.FriendshipLevel(),
				// 		})
				// 	}
				// })

				//玩家邮件信息
				maiitem := make(map[int64][]MailItemxdgm)
				db.Select.PlayerMailAttachment(func(row *mdb.PlayerMailAttachmentRow) {
					maiitem[row.PlayerMailId()] = append(maiitem[row.PlayerMailId()], MailItemxdgm{
						ItemId:  row.ItemId(),
						ItemNum: row.ItemNum(),
					})
				})
				db.Select.PlayerMailAttachmentLg(func(row *mdb.PlayerMailAttachmentLgRow) {
					maiitem[row.PlayerMailId()] = append(maiitem[row.PlayerMailId()], MailItemxdgm{
						ItemId:  row.ItemId(),
						ItemNum: row.ItemNum(),
					})
				})
				db.Select.PlayerMail(func(row *mdb.PlayerMailRow) {
					var mailtitle, mailcontent string
					if row.MailId() > 0 {
						paraslice := make([]string, 0)
						mailtitle, mailcontent = mail_dat.GetMail(row.MailId())
						err := json.Unmarshal([]byte(row.Parameters()), &paraslice)
						fail.When(err != nil, "jsondecode error")
						i := 0
						for _, v := range paraslice {
							mailcontent = strings.Replace(mailcontent, "{"+fmt.Sprintf("%v", i)+"}", v, -1)
							i++
						}
					} else {
						mailtitle, mailcontent = row.Title(), row.Content()
					}
					reply.Mails = append(reply.Mails, MailsInfo{
						MailTitle:   mailtitle,
						MailContent: mailcontent,
						SendTime:    row.SendTime(),
						ItemDetail:  maiitem[row.Id()],
					})
				})
				db.Select.PlayerMailLg(func(row *mdb.PlayerMailLgRow) {
					var mailtitle, mailcontent string
					if row.MailId() > 0 {
						paraslice := make([]string, 0)
						mailtitle, mailcontent = mail_dat.GetMail(row.MailId())
						err := json.Unmarshal([]byte(row.Parameters()), &paraslice)
						fail.When(err != nil, "jsondecode error")
						i := 0
						for _, v := range paraslice {
							mailcontent = strings.Replace(mailcontent, "{"+fmt.Sprintf("%v", i)+"}", v, -1)
							i++
						}
					} else {
						mailtitle, mailcontent = row.Title(), row.Content()
					}
					reply.Mails = append(reply.Mails, MailsInfo{
						MailTitle:   mailtitle,
						MailContent: mailcontent,
						SendTime:    row.SendTime(),
						ItemDetail:  maiitem[row.Pmid()],
					})
				})
				PlayerInfo := &role_api.GetPlayerInfo_Out{}
				module.Role.GetOtherPlayerInfo(args.Pid, db, &PlayerInfo.Player)
				for k, v := range PlayerInfo.Player.Roles {
					for k2, v2 := range v.Equips {
						itemInfo := item_dat.GetItem(v2.ItemId)
						if itemInfo.Health > 0 {
							basehealth := itemInfo.RefineBase*int32(v2.RefineLevel) + itemInfo.Health
							v2.Health += basehealth
						}
						if itemInfo.Speed > 0 {
							basespeed := itemInfo.RefineBase*int32(v2.RefineLevel) + itemInfo.Speed
							v2.Speed += basespeed
						}
						if itemInfo.Attack > 0 {
							baseattack := itemInfo.RefineBase*int32(v2.RefineLevel) + itemInfo.Attack
							v2.Attack += baseattack
						}
						if itemInfo.Defence > 0 {
							basedefence := itemInfo.RefineBase*int32(v2.RefineLevel) + itemInfo.Defence
							v2.Defence += basedefence
						}
						v.Equips[k2] = v2
					}
					PlayerInfo.Player.Roles[k] = v
				}
				reply.PlayerInfo = PlayerInfo.Player.Roles
			})
		})
		return nil
	})
}

/*
	玩家比武场排名
*/
type Args_XdgmgetUserRank struct {
	RPCArgTag
	Pid int64
}

type Reply_XdgmgetUserRank struct {
	Rank int32
}

func (this *RemoteServe) XdgmGetUserRank(args *Args_XdgmgetUserRank, reply *Reply_XdgmgetUserRank) error {
	return Remote.Serve(mdb.RPC_Remote_XdgmGetUserRank, args, mdb.TRANS_TAG_RPC_Serve_XdgmGetUserRank, func() error {
		if ok := mdb.CheckPlayer(args.Pid); !ok {
			return nil
		}
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			rank := module.ArenaRPC.GetPlayerRank(args.Pid)
			reply.Rank = rank
		})
		return nil
	})
}
