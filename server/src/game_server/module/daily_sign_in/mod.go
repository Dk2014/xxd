package daily_sign_in

import (
	"core/fail"
	"core/time"
	"game_server/dat/daily_sign_in_dat"
	"game_server/dat/mail_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/module/rpc"
)

type DailySignInMod struct {
}

func init() {
	module.DailySignIn = DailySignInMod{}
}

func (mod DailySignInMod) GetAwardForToday(db *mdb.Database) (awardConfig *daily_sign_in_dat.DailySignInAward) {
	playerInfo := db.Lookup.PlayerInfo(db.PlayerId())
	firstLoginDate := time.GetNowDayFromUnix(playerInfo.FirstLoginTime)
	nowDate := time.GetNowDay()
	if isNewPlayer(firstLoginDate, nowDate) {
		awardConfig = daily_sign_in_dat.NewPlayerSignInAward(nowDate - firstLoginDate + 1)
	} else {
		awardConfig = daily_sign_in_dat.RegularPlayerSignInAward(nowDate)
	}
	return awardConfig
}

func (mod DailySignInMod) SignedToday(db *mdb.Database) (isSigned bool) {
	playerSignInState, _ := updateSignInState(db)
	index := currentIndex(db, playerSignInState)
	isSigned = ((playerSignInState.Record >> uint(index)) & 1) == 1
	return isSigned
}

//发送当天奖励：仅用于发送签到后成为VIP用户的需要补偿奖励的特殊情况，故不会检查发送双倍奖励
func (mod DailySignInMod) AwardByMail(db *mdb.Database, awardConfig *daily_sign_in_dat.DailySignInAward) {
	mail := mail_dat.MailVIPSignAward{}
	switch awardConfig.AwardType {
	case daily_sign_in_dat.ITEM:
		mail.Attachments = append(mail.Attachments, &mail_dat.Attachment{
			AttachmentType: mail_dat.ATTACHMENT_ITEM,
			ItemId:         awardConfig.AwardId,
			ItemNum:        int64(awardConfig.Num),
		})
	case daily_sign_in_dat.HEART:
		mail.Attachments = append(mail.Attachments, &mail_dat.Attachment{
			AttachmentType: mail_dat.ATTACHMENT_HEART,
			ItemId:         awardConfig.AwardId,
			ItemNum:        int64(awardConfig.Num),
		})
	case daily_sign_in_dat.COINS:
		mail.Attachments = append(mail.Attachments, &mail_dat.Attachment{
			AttachmentType: mail_dat.ATTACHMENT_COINS,
			ItemId:         awardConfig.AwardId,
			ItemNum:        int64(awardConfig.Num),
		})
	case daily_sign_in_dat.INGOT:
		mail.Attachments = append(mail.Attachments, &mail_dat.Attachment{
			AttachmentType: mail_dat.ATTACHMENT_INGOT,
			ItemId:         awardConfig.AwardId,
			ItemNum:        int64(awardConfig.Num),
		})
	case daily_sign_in_dat.SWORD_SOUL:
		mail.Attachments = append(mail.Attachments, &mail_dat.Attachment{
			AttachmentType: mail_dat.ATTACHMENT_SWORD_SOUL,
			ItemId:         awardConfig.AwardId,
			ItemNum:        int64(awardConfig.Num),
		})
	case daily_sign_in_dat.GHOST:
		mail.Attachments = append(mail.Attachments, &mail_dat.Attachment{
			AttachmentType: mail_dat.ATTACHMENT_GHOST,
			ItemId:         awardConfig.AwardId,
			ItemNum:        int64(awardConfig.Num),
		})
	default:
		fail.When(true, "[每日签到] 未定义的奖励类型")

	}
	rpc.RemoteMailSend(db.PlayerId(), mail)
}
