package mail

import (
	"game_server/dat/mail_dat"
	"game_server/mdb"
	"game_server/module"
)

func init() {
	module.Mail = MailMod{}
}

type MailMod struct{}

//仅供RPC调用，产生一条本地的 mail 数据
func (mailMod MailMod) SendMail(db *mdb.Database, mail mail_dat.Mailer) {
	sendNewMail(db, mail)
}
