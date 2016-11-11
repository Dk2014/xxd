package mail

import (
	"core/net"
	"game_server/api/protocol/mail_api"
	"game_server/mdb"
	"game_server/module"
	"game_server/tlog"
)

type MailAPI struct{}

func init() {
	mail_api.SetInHandler(MailAPI{})
}

func (m MailAPI) GetList(session *net.Session, in *mail_api.GetList_In) {
	state := module.State(session)
	tlog.PlayerSystemModuleFlowLog(state.Database, tlog.SMT_MAIL)
	out := &mail_api.GetList_Out{}
	out.GetHeartNum = module.Heart.GetHeartDailyNum(state)

	mailIter(state.Database, func(row *mdb.PlayerMailRow, attachments []*mdb.PlayerMailAttachment) {
		mail := mail_api.GetList_Out_Mails{
			Id:         row.Id(),
			MailId:     row.MailId(),
			State:      row.State(),
			SendTime:   row.SendTime(),
			ExpireTime: row.ExpireTime(),
			Priority:   row.Priority(),
			Content:    []byte(row.Content()),
			Title:      []byte(row.Title()),
			Parameters: []byte(row.Parameters()),
		}
		for _, attachment := range attachments {
			mail.Attachments = append(mail.Attachments, mail_api.GetList_Out_Mails_Attachments{
				Id:             attachment.Id,
				AttachmentType: attachment.AttachmentType,
				ItemId:         attachment.ItemId,
				ItemNum:        attachment.ItemNum,
			})
		}
		out.Mails = append(out.Mails, mail)
	})

	session.Send(out)
}

func (m MailAPI) Read(session *net.Session, in *mail_api.Read_In) {
	state := module.State(session)
	readMail(state.Database, in.Id)
	session.Send(&mail_api.Read_Out{})
}

func (m MailAPI) TakeAttachment(session *net.Session, in *mail_api.TakeAttachment_In) {
	state := module.State(session)
	takeAttachment(state, in.AttachmentId)
	session.Send(&mail_api.TakeAttachment_Out{})
}

func (m MailAPI) GetInfos(session *net.Session, in *mail_api.GetInfos_In) {
	state := module.State(session)
	total, unread := getMailCount(state.Database)
	playerInfo := state.Database.Lookup.PlayerInfo(state.PlayerId)
	newMailNum := playerInfo.NewMailNum
	if playerInfo.NewMailNum > 0 {
		playerInfo.NewMailNum = 0
		state.Database.Update.PlayerInfo(playerInfo)
	}
	session.Send(&mail_api.GetInfos_Out{
		UnreadNum:  unread,
		Total:      total,
		NewMailNum: newMailNum,
	})
}

func (m MailAPI) RequestGlobalMail(session *net.Session, in *mail_api.RequestGlobalMail_In) {
	state := module.State(session)
	module.MailRPC.SendGlobaldMail(state.Database, state.PlayerId)
}
