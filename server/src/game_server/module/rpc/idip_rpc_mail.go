package rpc

import (
	"core/fail"
	"fmt"
	"game_server/dat/item_dat"
	"game_server/dat/mail_dat"
	"game_server/mdb"
	"game_server/module"
	. "game_server/rpc"
)

/*
	发放邮件
*/
type Args_IdipSendMail struct {
	RPCArgTag
	OpenId      string // openid
	MailTitle   string // 邮件标题
	MailContent string // 邮件内容
	SendTime    uint64 // 邮件发送时间
	EndTime     uint64 // 邮件自动删除时间
	ItemDetail  MailItem
}

type MailItem struct {
	ItemId1  uint64 // 邮件赠送道具1ID
	ItemNum1 uint32 // 邮件赠送道具1数量
	ItemId2  uint64 // 邮件赠送道具2ID
	ItemNum2 uint32 // 邮件赠送道具2数量
	ItemId3  uint64 // 邮件赠送道具3ID
	ItemNum3 uint32 // 邮件赠送道具3数量
	ItemId4  uint64 // 邮件赠送道具4ID
	ItemNum4 uint32 // 邮件赠送道具4数量
}

type Reply_IdipSendMail struct {
	Result int32  // 结果：（0）成功，（其他）失败
	RetMsg string // 返回消息
	ErrMsg string // 错误信息
}

func (this *RemoteServe) IdipSendMail(args *Args_IdipSendMail, reply *Reply_IdipSendMail) error {
	return Remote.Serve(mdb.RPC_Remote_IdipSendMail, args, mdb.TRANS_TAG_RPC_Serve_IdipSendMail, func() error {
		reply.Result = 2
		reply.RetMsg = "failure"
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			attachments := getattachments(args)
			if args.OpenId == "" {
				module.MailRPC.AddGlobalMail(globalDB, args.MailTitle, args.MailContent, attachments, int64(args.SendTime), int64(args.EndTime), 0, 0, 0, 0, 0)
				reply.Result = 0
				reply.RetMsg = "success"
				return
			}
			pid, ok := module.Player.GetPlayerByUsername(string(args.OpenId))
			if !ok {
				reply.ErrMsg = "query empty"
				return
			}
			RemoteMailSend(pid, &mail_dat.EmptyMail{
				MailId:      0,
				Title:       args.MailTitle,
				Content:     args.MailContent,
				Parameters:  "",
				Attachments: attachments,
				SendTime:    int64(args.SendTime),
				ExpireTime:  int64(args.EndTime),
			})
			reply.Result = 0
			reply.RetMsg = "success"
		})
		return nil
	})
}

/*
	发放邮件(AQ)
*/
type Args_IdipAqSendMail struct {
	RPCArgTag
	OpenId      string // openid
	MailContent string // 邮件内容
}

type Reply_IdipAqSendMail struct {
	Result int32  // 结果：（0）成功，（其他）失败
	RetMsg string // 返回消息
	ErrMsg string // 错误信息
}

func (this *RemoteServe) IdipAqSendMail(args *Args_IdipAqSendMail, reply *Reply_IdipAqSendMail) error {
	return Remote.Serve(mdb.RPC_Remote_IdipAqSendMail, args, mdb.TRANS_TAG_RPC_Serve_IdipAqSendMail, func() error {
		pid, ok := module.Player.GetPlayerByUsername(string(args.OpenId))
		if !ok {
			reply.ErrMsg = "query empty"
			return nil
		}
		RemoteMailSend(pid, &mail_dat.EmptyMail{
			MailId:      0,
			Title:       "系统消息",
			Content:     args.MailContent,
			Parameters:  "",
			Attachments: nil,
			ExpireTime:  0,
		})
		reply.Result = 0
		reply.RetMsg = "success"
		return nil
	})
}

/*
	发放全区邮件
*/
func (this *RemoteServe) IdipSendAllMail(args *Args_IdipSendMail, reply *Reply_IdipSendMail) error {
	return Remote.Serve(mdb.RPC_Remote_IdipSendAllMail, args, mdb.TRANS_TAG_RPC_Serve_IdipSendAllMail, func() error {
		reply.Result = 2
		reply.RetMsg = "failure"
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			attachments := getattachments(args)
			module.MailRPC.AddGlobalMail(globalDB, args.MailTitle, args.MailContent, attachments, int64(args.SendTime), int64(args.EndTime), 0, 0, 0, 0, 0)
			reply.Result = 0
			reply.RetMsg = "success"
			return
		})
		return nil
	})
}

func getattachments(args *Args_IdipSendMail) []*mail_dat.Attachment {
	attachments := []*mail_dat.Attachment{}
	if args.ItemDetail.ItemNum1 > 0 {
		attachments = appendToAttchment(args.ItemDetail.ItemId1, args.ItemDetail.ItemNum1, attachments)
	}
	if args.ItemDetail.ItemNum2 > 0 {
		attachments = appendToAttchment(args.ItemDetail.ItemId2, args.ItemDetail.ItemNum2, attachments)
	}
	if args.ItemDetail.ItemNum3 > 0 {
		attachments = appendToAttchment(args.ItemDetail.ItemId3, args.ItemDetail.ItemNum3, attachments)
	}
	if args.ItemDetail.ItemNum4 > 0 {
		attachments = appendToAttchment(args.ItemDetail.ItemId4, args.ItemDetail.ItemNum4, attachments)
	}
	return attachments
}

func appendToAttchment(itemid uint64, itemnum uint32, attachments []*mail_dat.Attachment) []*mail_dat.Attachment {
	itemInfo := item_dat.GetItem(int16(itemid))
	switch itemInfo.TypeId {
	case item_dat.TYPE_CLOTHES:
		fail.When(itemnum > 1, fmt.Sprintf("%d just allow 1 num", itemid))
	case item_dat.TYPE_WEAPON:
		fail.When(itemnum > 1, fmt.Sprintf("%d just allow 1 num", itemid))
	case item_dat.TYPE_SHOE:
		fail.When(itemnum > 1, fmt.Sprintf("%d just allow 1 num", itemid))
	case item_dat.TYPE_FASHION:
		fail.When(itemnum > 1, fmt.Sprintf("%d just allow 1 num", itemid))
	case item_dat.TYPE_CHEAT:
		fail.When(itemnum > 1, fmt.Sprintf("%d just allow 1 num", itemid))
	case item_dat.TYPE_ACCESSORIES:
		fail.When(itemnum > 1, fmt.Sprintf("%d just allow 1 num", itemid))
	case item_dat.TYPE_RESOURCE:
		break
	default:
		fail.When(itemnum > 999, fmt.Sprintf("%d just allow 999 num", itemid))
	}
	switch itemid {
	//爱心
	case uint64(247):
		attachments = append(attachments, &mail_dat.Attachment{
			AttachmentType: mail_dat.ATTACHMENT_HEART,
			ItemId:         0,
			ItemNum:        int64(itemnum),
		})
	//铜钱
	case uint64(244):
		attachments = append(attachments, &mail_dat.Attachment{
			AttachmentType: mail_dat.ATTACHMENT_COINS,
			ItemId:         0,
			ItemNum:        int64(itemnum),
		})
	case uint64(242):
		attachments = append(attachments, &mail_dat.Attachment{
			AttachmentType: mail_dat.ATTACHMENT_INGOT,
			ItemId:         0,
			ItemNum:        int64(itemnum),
		})
	case uint64(243):
		break
	case uint64(245):
		break
	case uint64(246):
		break
	case uint64(248):
		break
	default:
		attachments = append(attachments, &mail_dat.Attachment{
			AttachmentType: mail_dat.ATTACHMENT_ITEM,
			ItemId:         int16(itemid),
			ItemNum:        int64(itemnum),
		})
	}
	return attachments
}
