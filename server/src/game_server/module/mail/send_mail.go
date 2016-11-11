package mail

import (
	"core/time"
	"game_server/dat/mail_dat"
	"game_server/mdb"
	"game_server/module"
)

func sendNewMail(db *mdb.Database, mail mail_dat.Mailer) {
	nowTime := time.GetNowTime()

	pMail := &mdb.PlayerMail{
		Pid:        db.PlayerId(),
		Title:      mail.GetTitle(),
		Content:    mail.GetContent(),
		MailId:     mail.GetMailId(),
		Parameters: mail.GetParameters(),
		State:      mail_dat.UN_READ,
		Priority:   mail.GetPriority(),
		SendTime:   nowTime,
		ExpireTime: mail.GetExpireTime(),
	}

	attachments := mail.GetAttachments()
	if len(attachments) > 0 {
		pMail.HaveAttachment = mail_dat.HAVE_ATTACHMENT
	} else {
		//fail.When(pMail.ExpireTime == mail_dat.AUTO_DELETE_AFTER_READ_WITHOUT_ATTACHMENT, "没有附件的邮件不能设置*领取附件后删除")
		if pMail.ExpireTime == mail_dat.AUTO_DELETE_AFTER_READ_WITHOUT_ATTACHMENT {
			pMail.ExpireTime = mail_dat.AUTO_DELETE_AFTER_EXPIRED
		}
		pMail.HaveAttachment = mail_dat.NO_ATTACHMENT
	}

	db.Insert.PlayerMail(pMail)

	//附件操作
	var newAttachments []*mdb.PlayerMailAttachment
	if pMail.HaveAttachment == mail_dat.HAVE_ATTACHMENT {
		for _, attachment := range attachments {
			attach := addAttachment(db, pMail.Id, attachment)
			newAttachments = append(newAttachments, attach)
		}
	}

	//Notify
	session, ok := module.Player.GetPlayerOnline(db.PlayerId())
	if ok {
		module.Notify.NewMail(session)
	} else {
		playerInfo := db.Lookup.PlayerInfo(db.PlayerId())
		playerInfo.NewMailNum++
		db.Update.PlayerInfo(playerInfo)
	}
}

// 为一封邮件添加附件,mailId为player_mail 主键
func addAttachment(db *mdb.Database, mailId int64, attachment *mail_dat.Attachment) *mdb.PlayerMailAttachment {
	attach := &mdb.PlayerMailAttachment{
		Pid:            db.PlayerId(),
		PlayerMailId:   mailId,
		AttachmentType: attachment.AttachmentType,
		ItemId:         attachment.ItemId,
		ItemNum:        attachment.ItemNum,
	}
	db.Insert.PlayerMailAttachment(attach)
	return attach
}
