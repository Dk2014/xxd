package rpc

import (
	"core/fail"
	"core/time"
	"encoding/json"
	"fmt"
	"game_server/dat/item_dat"
	"game_server/dat/mail_dat"
	"game_server/mdb"
	"game_server/module"
	. "game_server/rpc"
)

/*
	发送单人邮件
*/
type Args_XdgmSendMail struct {
	RPCArgTag
	Pid     int64
	Title   string
	Content string
	Attach  string
	Endtime int64
}

type Reply_XdgmSendMail struct {
}

func (this *RemoteServe) XdgmSendMail(args *Args_XdgmSendMail, reply *Reply_XdgmSendMail) error {
	return Remote.Serve(mdb.RPC_Remote_XdgmSendMail, args, mdb.TRANS_TAG_RPC_Serve_XdgmSendMail, func() error {
		fail.When(!mdb.CheckPlayer(args.Pid), "pid does not exists")
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(args.Pid, func(db *mdb.Database) {
				attachments := xdgmgetattachments(args.Attach)
				RemoteMailSend(args.Pid, &mail_dat.EmptyMail{
					MailId:      0,
					Title:       args.Title,
					Content:     args.Content,
					Parameters:  "",
					Attachments: attachments,
					SendTime:    time.GetNowTime(),
					ExpireTime:  args.Endtime,
				})
			})
		})
		return nil
	})
}

/*
	发放全区邮件
*/
type Args_XdgmSendMailAll struct {
	RPCArgTag
	Title       string
	Content     string
	Attach      string
	Begintime   int64
	Endtime     int64
	MinLevel    int16
	MaxLevel    int16
	MinVipLevel int16
	MaxVipLevel int16
}

type Reply_XdgmSendMailAll struct {
}

func (this *RemoteServe) XdgmSendMailAll(args *Args_XdgmSendMailAll, reply *Reply_XdgmSendMailAll) error {
	return Remote.Serve(mdb.RPC_Remote_XdgmSendMailAll, args, mdb.TRANS_TAG_RPC_Serve_XdgmSendMailAll, func() error {
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			attachments := xdgmgetattachments(args.Attach)
			module.MailRPC.AddGlobalMail(globalDB, args.Title, args.Content, attachments, args.Begintime, args.Endtime, 0, args.MinLevel, args.MinVipLevel, args.MaxLevel, args.MaxVipLevel)
		})
		return nil
	})
}

func xdgmgetattachments(attach string) []*mail_dat.Attachment {
	attachments := []*mail_dat.Attachment{}
	attachinterface := make([]map[string]int64, 0)
	json.Unmarshal([]byte(attach), &attachinterface)
	i := 0
	for _, v := range attachinterface {
		if i > 4 {
			break
		}
		i++
		if v["itemnum"] > 0 {
			attachments = xdgmappendToAttchment(int16(v["itemid"]), v["itemnum"], attachments)
		}
	}
	return attachments
}

func xdgmappendToAttchment(itemid int16, itemnum int64, attachments []*mail_dat.Attachment) []*mail_dat.Attachment {
	itemInfo := item_dat.GetItem(itemid)
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
	case 247:
		attachments = append(attachments, &mail_dat.Attachment{
			AttachmentType: mail_dat.ATTACHMENT_HEART,
			ItemId:         0,
			ItemNum:        itemnum,
		})
	//铜钱
	case 244:
		attachments = append(attachments, &mail_dat.Attachment{
			AttachmentType: mail_dat.ATTACHMENT_COINS,
			ItemId:         0,
			ItemNum:        itemnum,
		})
	case 242:
		attachments = append(attachments, &mail_dat.Attachment{
			AttachmentType: mail_dat.ATTACHMENT_INGOT,
			ItemId:         0,
			ItemNum:        itemnum,
		})
	case 243:
		break
	case 245:
		break
	case 246:
		break
	case 248:
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
