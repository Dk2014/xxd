package rpc

import (
	"game_server/dat/mail_dat"
	"game_server/mdb"
	"game_server/module"
	. "game_server/rpc"
)

//----------------------------------------------------------------------------------------------------
//					    发送邮件
//----------------------------------------------------------------------------------------------------
type Args_MailSend struct {
	RPCArgTag
	Pid         int64
	MailId      int32
	Title       string
	Content     string
	Parameters  string
	Attachments []*mail_dat.Attachment
	SendTime    int64
	ExpireTime  int64
	Priority    int8
	MinLevel    int16 //0不限制
	MaxLevel    int16 //0不限制
	MinVIPLevel int16 //0不限制
	MaxVIPLevel int16 //0不限制
}

//实现 mail_dat.Mailer 接口
func (this *Args_MailSend) GetMailId() int32 {
	return this.MailId
}

func (this *Args_MailSend) GetSendTime() int64 {
	return this.SendTime
}

func (this *Args_MailSend) GetExpireTime() int64 {
	return this.ExpireTime
}

func (this *Args_MailSend) GetTitle() string {
	return this.Title
}

func (this *Args_MailSend) GetContent() string {
	return this.Content
}

func (this *Args_MailSend) GetParameters() string {
	return this.Parameters
}

func (this *Args_MailSend) GetPriority() int8 {
	return this.Priority
}

func (this *Args_MailSend) GetMinLevel() int16 {
	return this.MinLevel
}

func (this *Args_MailSend) GetMaxLevel() int16 {
	return this.MaxLevel
}

func (this *Args_MailSend) GetMinVIPLevel() int16 {
	return this.MinVIPLevel
}

func (this *Args_MailSend) GetMaxVIPLevel() int16 {
	return this.MaxVIPLevel
}

func (this *Args_MailSend) GetAttachments() []*mail_dat.Attachment {
	return this.Attachments
}

type Reply_MailSend struct {
}

//游戏服务器作为RPC服务器，本服玩家通过该服务增加邮件
func (mail *RemoteServe) MailSend(args *Args_MailSend, reply *Reply_MailSend) error {
	return Remote.Serve(mdb.RPC_Remote_MailSend, args, mdb.TRANS_TAG_RPC_Serve_MailSend, func() error {
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(args.Pid, func(db *mdb.Database) {
				vipInfo := module.VIP.VIPInfo(db)
				mainRole := module.Role.GetMainRole(db)
				if args.MinLevel > 0 && mainRole.Level < args.MinLevel {
					return
				}
				if args.MaxLevel > 0 && mainRole.Level > args.MaxLevel {
					return
				}
				if args.MinVIPLevel > 0 && vipInfo.Level < args.MinVIPLevel {
					return
				}
				if args.MaxVIPLevel > 0 && vipInfo.Level > args.MaxVIPLevel {
					return
				}
				module.Mail.SendMail(db, args)
			})
		})
		return nil
	})
}

func RemoteMailSend(pid int64, mailer mail_dat.Mailer) {
	args := &Args_MailSend{
		Pid:         pid,
		MailId:      mailer.GetMailId(),
		Title:       mailer.GetTitle(),
		Content:     mailer.GetContent(),
		Parameters:  mailer.GetParameters(),
		Attachments: mailer.GetAttachments(),
		ExpireTime:  mailer.GetExpireTime(),
		Priority:    mailer.GetPriority(),
	}
	serverId, _ := module.GetServerIdWithPlayerId(pid)
	Remote.Call(serverId, mdb.RPC_Remote_MailSend, args, &Reply_MailSend{}, mdb.TRANS_TAG_RPC_Call_MailSend, nil)
}

//----------------------------------------------------------------------------------------------------
//					批量发送邮件
//----------------------------------------------------------------------------------------------------

type Args_MailBatchSend struct {
	RPCArgTag
	Pid   int64
	Mails []*mail_dat.EmptyMail
}

type Reply_MailBatchSend struct {
}

//游戏服务器作为RPC服务器，本服玩家通过该服务增加邮件
func (mail *RemoteServe) MailBatchSend(args *Args_MailBatchSend, reply *Reply_MailBatchSend) error {
	return Remote.Serve(mdb.RPC_Remote_MailBatchSend, args, mdb.TRANS_TAG_RPC_Serve_MailBatchSend, func() error {
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(args.Pid, func(db *mdb.Database) {
				vipInfo := module.VIP.VIPInfo(db)
				mainRole := module.Role.GetMainRole(db)
				for i := 0; i < len(args.Mails); i++ {
					if args.Mails[i].GetMinLevel() > 0 && mainRole.Level < args.Mails[i].GetMinLevel() {
						continue
					}
					if args.Mails[i].GetMaxLevel() > 0 && mainRole.Level > args.Mails[i].GetMaxLevel() {
						continue
					}
					if args.Mails[i].GetMinVIPLevel() > 0 && vipInfo.Level < args.Mails[i].GetMinVIPLevel() {
						continue
					}
					if args.Mails[i].GetMaxVIPLevel() > 0 && vipInfo.Level > args.Mails[i].GetMaxVIPLevel() {
						continue
					}
					module.Mail.SendMail(db, args.Mails[i])
				}
			})
		})
		return nil
	})
}

func RemoteMailBatchSend(pid int64, mails []*mail_dat.EmptyMail) {
	if mails == nil || len(mails) == 0 {
		return
	}
	args := &Args_MailBatchSend{
		Pid:   pid,
		Mails: mails,
	}
	serverId, _ := module.GetServerIdWithPlayerId(pid)
	Remote.Call(serverId, mdb.RPC_Remote_MailBatchSend, args, &Reply_MailBatchSend{}, mdb.TRANS_TAG_RPC_Call_MailBatchSend, func(err error) {
		//发送邮件成功后，通过这个回调函数更新玩家的最大邮件时间戳
		if err == nil {
			var maxTimestamp int64
			for i := 0; i < len(mails); i++ {
				if mails[i].SendTime > maxTimestamp {
					maxTimestamp = mails[i].SendTime
				}
			}
			mdb.GlobalExecute(func(globalDB *mdb.Database) {
				globalDB.AgentExecute(pid, func(db *mdb.Database) {
					playerMailState := db.Lookup.PlayerGlobalMailState(pid)
					if playerMailState.MaxTimestamp < maxTimestamp {
						playerMailState.MaxTimestamp = maxTimestamp
						db.Update.PlayerGlobalMailState(playerMailState)
					}
				})
			})
		}
	})
}
