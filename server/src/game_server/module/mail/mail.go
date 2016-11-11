package mail

import (
	"core/time"

	"core/fail"
	"game_server/dat/mail_dat"
	"game_server/dat/player_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/tlog"
	"game_server/xdlog"
)

// 获取邮件附件
func getMailAttachMents(db *mdb.Database, id int64) []*mdb.PlayerMailAttachment {
	attachments := make([]*mdb.PlayerMailAttachment, 0, 4)

	db.Select.PlayerMailAttachment(func(row *mdb.PlayerMailAttachmentRow) {
		if row.PlayerMailId() == id {
			attachments = append(attachments, row.GoObject())
		}
	})

	return attachments
}

// 提取邮件附件
func takeAttachment(state *module.SessionState, attachmentId int64) {

	attachment := state.Database.Lookup.PlayerMailAttachment(attachmentId)
	fail.When(attachment == nil, "takeAttachment: attachment not found")
	fail.When(attachment.ItemNum <= 0, "takeAttachment: item num should > 0")

	switch attachment.AttachmentType {
	case mail_dat.ATTACHMENT_ITEM:
		//普通物品
		module.Item.AddItem(state.Database, attachment.ItemId, int16(attachment.ItemNum), tlog.IFR_MAIL_TAKE_ATTACHMENT, xdlog.ET_MAIL_TAKE_ATTACHMENT, "")
	case mail_dat.ATTACHMENT_COINS:
		//铜钱
		module.Player.IncMoney(state.Database, state.MoneyState, attachment.ItemNum, player_dat.COINS, tlog.MFR_MAIL_TAKE_ATTACHMENT, xdlog.ET_MAIL_TAKE_ATTACHMENT, "")
	case mail_dat.ATTACHMENT_RELATIONSHIP:
		//友情
		module.Team.IncRelationship(state.Database, int32(attachment.ItemNum))
	case mail_dat.ATTACHMENT_FAME:
		//声望
		module.Player.AddFame(state.Database, int32(attachment.ItemNum))
	case mail_dat.ATTACHMENT_INGOT:
		// 元宝
		module.Player.IncMoney(state.Database, state.MoneyState, attachment.ItemNum, player_dat.INGOT, tlog.MFR_MAIL_TAKE_ATTACHMENT, xdlog.ET_MAIL_TAKE_ATTACHMENT, "")
	case mail_dat.ATTACHMENT_HEART:
		// 爱心
		module.Heart.IncHeart(state, int16(attachment.ItemNum))
	case mail_dat.ATTACHMENT_HEART_FROM_FRIEND:
		//好友爱心
		module.Heart.IncHeartFromFriend(state, int16(attachment.ItemNum))
	case mail_dat.ATTACHMENT_FORMATION_EXP:
		//经验
		module.Role.AddFormRoleExp(state, int64(attachment.ItemNum), tlog.EFT_MAIL_TAKE_ATTACHMENT)
	case mail_dat.ATTACHMENT_SINGLE_ROLE_EXP:
		//单个角色加经验
		roleId := int8(attachment.ItemId)
		module.Role.AddRoleExp(state.Database, roleId, int64(attachment.ItemNum), state.RoleId, tlog.EFT_MAIL_TAKE_ATTACHMENT)
	case mail_dat.ATTACHMENT_GHOST:
		// 魂侍
		module.Ghost.AddGhost(state, attachment.ItemId, tlog.IFR_MAIL_TAKE_ATTACHMENT, xdlog.ET_MAIL_TAKE_ATTACHMENT)
	case mail_dat.ATTACHMENT_SWORD_SOUL:
		// 剑心
		module.SwordSoul.AddSwordSoul(state, attachment.ItemId, tlog.IFR_MAIL_TAKE_ATTACHMENT)
	case mail_dat.ATTACHMENT_TOTEM:
		//TODO tlog
		//TODO 检查背包?
		module.Totem.AddTotem(state.Database, attachment.ItemId)
	}
	deleteMailAttachment(state.Database, attachment, false)
}

func readMail(db *mdb.Database, id int64) {
	mail := db.Lookup.PlayerMail(id)
	fail.When(mail == nil, "readMail: mail not found")
	mail.State = mail_dat.HAS_BEEN_READ
	db.Update.PlayerMail(mail)
}

func getMailCount(db *mdb.Database) (total, unread int16) {
	mailIter(db, func(row *mdb.PlayerMailRow, attachments []*mdb.PlayerMailAttachment) {
		total++
		if row.State() == mail_dat.UN_READ {
			unread++
		}
	})
	return
}

func mailIter(db *mdb.Database, worker func(*mdb.PlayerMailRow, []*mdb.PlayerMailAttachment)) {
	now := time.GetNowTime()
	db.Select.PlayerMail(func(row *mdb.PlayerMailRow) {
		days := (now - row.SendTime()) / 86400
		var attachments []*mdb.PlayerMailAttachment
		if row.HaveAttachment() == mail_dat.HAVE_ATTACHMENT {
			//如果有附件
			attachments = getMailAttachMents(db, row.Id())
			//领取了附件马上删除
			if len(attachments) == 0 {
				deleteMail(db, row.GoObject())
				return
			}
			//过期自动删除 或者 到达指定删除时间 或者 爱心邮件使用独立的过期时间
			if row.MailId() == mail_dat.CLIQUE_SALARY_OWNER_MAIL_ID || row.MailId() == mail_dat.CLIQUE_SALARY_MANAGER_MAIL_ID {
				if days > mail_dat.CLIQUE_HAVE_ATTACHMENT_SAVE_DAYS {
					deleteMail(db, row.GoObject())
					for _, attachment := range attachments {
						deleteMailAttachment(db, attachment, true)
					}
					return
				}
			} else if (row.ExpireTime() <= mail_dat.AUTO_DELETE_MAGIC_NUM && days > mail_dat.HAVE_ATTACHMENT_SAVE_DAYS) ||
				(row.ExpireTime() > mail_dat.AUTO_DELETE_MAGIC_NUM && row.ExpireTime() < now ||
					row.MailId() == mail_dat.HEART_MAIL_ID && days > mail_dat.HEART_MAIL_SAVE_DAYS) {
				deleteMail(db, row.GoObject())
				for _, attachment := range attachments {
					deleteMailAttachment(db, attachment, true)
				}
				return
			}

		} else {
			if (row.ExpireTime() <= mail_dat.AUTO_DELETE_MAGIC_NUM && days > mail_dat.NO_ATTACHMENT_SAVE_DAYS) ||
				(row.ExpireTime() > mail_dat.AUTO_DELETE_MAGIC_NUM && row.ExpireTime() < now) {
				deleteMail(db, row.GoObject())
				return
			}
		}
		worker(row, attachments)

	})
}

func deleteMailAttachment(db *mdb.Database, attachment *mdb.PlayerMailAttachment, isExpired bool) {
	attachLg := &mdb.PlayerMailAttachmentLg{
		Pid:            attachment.Pid,
		PlayerMailId:   attachment.PlayerMailId,
		AttachmentType: attachment.AttachmentType,
		ItemId:         attachment.ItemId,
		ItemNum:        attachment.ItemNum,
	}
	if !isExpired {
		attachLg.TakeTimestamp = time.GetNowTime()
	}
	db.Insert.PlayerMailAttachmentLg(attachLg)
	db.Delete.PlayerMailAttachment(attachment)
}

func deleteMail(db *mdb.Database, mail *mdb.PlayerMail) {
	mailLg := &mdb.PlayerMailLg{
		Pmid:       mail.Id,     //原邮件ID
		Pid:        mail.Pid,    //玩家ID
		MailId:     mail.MailId, //模版ID
		Title:      mail.Title,
		Content:    mail.Content,
		Parameters: mail.Parameters,
		State:      mail.State,
		SendTime:   mail.SendTime,
	}
	db.Insert.PlayerMailLg(mailLg)
	db.Delete.PlayerMail(mail)
}
