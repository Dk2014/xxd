package player

import (
	"core/time"
	"game_server/api/protocol/notify_api"
	"game_server/dat/player_dat"
	"game_server/mdb"
	"game_server/module"
)

func activeMonthcard(db *mdb.Database) bool {
	playerMonthCard := db.Lookup.PlayerMonthCard(db.PlayerId())
	now := time.GetNowTime()
	needNotify := false
	ok := true //是否激活成功
	if playerMonthCard == nil {
		playerMonthCard = &mdb.PlayerMonthCard{
			Pid:             db.PlayerId(),
			ExpireTimestamp: time.GetTodayZero() + 30*86400,
		}
		needNotify = true
		db.Insert.PlayerMonthCard(playerMonthCard)
	} else {
		if playerMonthCard.ExpireTimestamp < now {
			//已过期
			playerMonthCard.ExpireTimestamp = time.GetTodayZero() + 30*86400
			needNotify = true
		} else {
			//未过期
			if playerMonthCard.ExpireTimestamp-player_dat.MONTH_CARD_RENEW_TIME_BEFORE_EXPIRE < now {
				//出于可续买时间
				playerMonthCard.ExpireTimestamp += 30 * 86400
				needNotify = true
			} else {
				//激活失败
				ok = false
			}
		}
		db.Update.PlayerMonthCard(playerMonthCard)
	}
	if needNotify {
		if session, online := module.Player.GetPlayerOnline(db.PlayerId()); online {
			session.Send(&notify_api.NotifyMonthCardChange_Out{
				ExpireAt: playerMonthCard.ExpireTimestamp,
			})
		}
	}
	return ok
}

func getMontCardExpireTimestamp(db *mdb.Database) int64 {
	if playerMonthCard := db.Lookup.PlayerMonthCard(db.PlayerId()); playerMonthCard != nil {
		return playerMonthCard.ExpireTimestamp
	}
	return 0
}

/*
func checkMonthCardAward(db *mdb.Database) {

	now := time.GetNowTime()
	playerMonthCard := db.Lookup.PlayerMonthCard(db.PlayerId())
	if playerMonthCard != nil && playerMonthCard.ExpireTimestamp < now {
		if !time.IsToday(playerMonthCard.AwardTimestamp) {
			//TODO 发送奖励
			playerMonthCard.AwardTimestamp = now
			db.Update.PlayerMonthCard(playerMonthCard)
		}
	}
}
*/
