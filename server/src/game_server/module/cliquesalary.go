package module

import (
	"core/debug"
	"core/log"
	"core/time"
	"game_server/dat/mail_dat"
	"game_server/mdb"
	"strconv"
	gotime "time"
)

//每天12点发工资
func Cliquesalary() {
	timeNow := time.GetNowTime()
	timeNextDayZero := time.GetNextTodayPointHour(0, gotime.Now())
	needSeconds := timeNextDayZero - timeNow
	gotime.AfterFunc(gotime.Second*gotime.Duration(needSeconds), func() {
		defer func() {
			if err := recover(); err != nil {
				log.Errorf("Cliquesalary() error: %v\n Stack %s \n", err, debug.Stack(1, "    "))
			}
		}()
		defer Cliquesalary()
		mdb.Transaction(mdb.TRANS_TAG_GlobalCliquesalary, func() {
			mdb.GlobalExecute(func(globalDB *mdb.Database) {
				globalDB.Select.GlobalClique(func(row *mdb.GlobalCliqueRow) {
					memberNum := CliqueRPC.GetMemberNum(row.Id())
					if memberNum <= 0 || memberNum > 40 {
						log.Errorf("CacheError CliqueRPC.GetMemberNum(%d) return %d \n", row.Id(), memberNum)
						memberNum = 1
					}
					var attachments []*mail_dat.Attachment
					attachments = append(attachments, &mail_dat.Attachment{
						AttachmentType: mail_dat.ATTACHMENT_COINS,
						ItemId:         0,
						ItemNum:        mail_dat.CLIQUE_SALARY_MAIL_OWNER_COIN_BASE * int64(memberNum),
					})
					MailRPC.SendSalaryMail(row.OwnerPid(), &mail_dat.Mailcliquesalaryowner{
						Name:        row.Name(),
						Num:         strconv.FormatInt(int64(memberNum), 10),
						Coins:       strconv.FormatInt(mail_dat.CLIQUE_SALARY_MAIL_OWNER_COIN_BASE*int64(memberNum), 10),
						Attachments: attachments,
					})
					if row.MangerPid1() > 0 {
						var attachments []*mail_dat.Attachment
						attachments = append(attachments, &mail_dat.Attachment{
							AttachmentType: mail_dat.ATTACHMENT_COINS,
							ItemId:         0,
							ItemNum:        mail_dat.CLIQUE_SALARY_MAIL_MANAGER_COIN_BASE * int64(memberNum),
						})
						MailRPC.SendSalaryMail(row.MangerPid1(), &mail_dat.Mailcliquesalarymanager{
							Name:        row.Name(),
							Num:         strconv.FormatInt(int64(memberNum), 10),
							Coins:       strconv.FormatInt(mail_dat.CLIQUE_SALARY_MAIL_MANAGER_COIN_BASE*int64(memberNum), 10),
							Attachments: attachments,
						})
					}
					if row.MangerPid2() > 0 {
						var attachments []*mail_dat.Attachment
						attachments = append(attachments, &mail_dat.Attachment{
							AttachmentType: mail_dat.ATTACHMENT_COINS,
							ItemId:         0,
							ItemNum:        mail_dat.CLIQUE_SALARY_MAIL_MANAGER_COIN_BASE * int64(memberNum),
						})
						MailRPC.SendSalaryMail(row.MangerPid2(), &mail_dat.Mailcliquesalarymanager{
							Name:        row.Name(),
							Num:         strconv.FormatInt(int64(memberNum), 10),
							Coins:       strconv.FormatInt(mail_dat.CLIQUE_SALARY_MAIL_MANAGER_COIN_BASE*int64(memberNum), 10),
							Attachments: attachments,
						})
					}
				})
			})
		})
	})
}
