package event

import (
	"core/fail"
	"core/net"
	"core/time"
	"game_server/api/protocol/event_api"
	"game_server/dat/event_dat"
	"game_server/dat/player_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/tlog"
	"game_server/xdlog"
)

func login_award_info(session *net.Session, out *event_api.LoginAwardInfo_Out) {
	state := module.State(session)
	playerLoginAwardInfo := updateLoginAwardInfo(state.Database)
	out.Record = playerLoginAwardInfo.Record
	out.TotalLoginDays = playerLoginAwardInfo.ActiveDays
}

func updateLoginAwardInfo(db *mdb.Database) *mdb.PlayerLoginAwardRecord {
	playerAwardInfo := db.Lookup.PlayerLoginAwardRecord(db.PlayerId())
	if time.IsToday(playerAwardInfo.UpdateTimestamp) {
		return playerAwardInfo
	}
	playerAwardInfo.ActiveDays += 1
	playerAwardInfo.UpdateTimestamp = time.GetNowTime()
	db.Update.PlayerLoginAwardRecord(playerAwardInfo)
	return playerAwardInfo
}

func takeLoginAward(session *net.Session, day int32) {
	state := module.State(session)
	playerLoginAwardInfo := updateLoginAwardInfo(state.Database)
	fail.When(day > playerLoginAwardInfo.ActiveDays, "累计登录天数不足")
	currentBit := uint32(1) << uint32(day-1)
	fail.When(currentBit&uint32(playerLoginAwardInfo.Record) != 0, "登录奖励已领取")

	awards := event_dat.GetLoginAwards(day)
	for _, award := range awards {
		loginAward(state, award)
	}

	playerLoginAwardInfo.Record += int32(currentBit)
	state.Database.Update.PlayerLoginAwardRecord(playerLoginAwardInfo)
}

func loginAward(state *module.SessionState, awardConfig *event_dat.LoginAward) {
	switch awardConfig.AwardType {
	case event_dat.LOGIN_AWARD_TYPE_ITEM:
		module.Item.AddItem(state.Database, int16(awardConfig.AwardId), int16(awardConfig.AwardNum), tlog.IFR_LOGIN_AWARD, xdlog.ET_EVENT_CENTER, "")
	case event_dat.LOGIN_AWARD_TYPE_SWORD_SOUL:
		module.SwordSoul.AddSwordSoul(state, int16(awardConfig.AwardId), tlog.IFR_LOGIN_AWARD)
	case event_dat.LOGIN_AWARD_TYPE_GHOST:
		module.Ghost.AddGhost(state, int16(awardConfig.AwardId), tlog.IFR_LOGIN_AWARD, xdlog.ET_EVENT_CENTER)
	case event_dat.LOGIN_AWARD_TYPE_HEART:
		module.Heart.IncHeart(state, int16(awardConfig.AwardNum))
	case event_dat.LOGIN_AWARD_TYPE_COINS:
		module.Player.IncMoney(state.Database, state.MoneyState, int64(awardConfig.AwardNum), player_dat.COINS, tlog.MFR_LOGIN_AWARD, xdlog.ET_EVENT_CENTER, "")
	case event_dat.LOGIN_AWARD_TYPE_INGOT:
		module.Player.IncMoney(state.Database, state.MoneyState, int64(awardConfig.AwardNum), player_dat.INGOT, tlog.MFR_LOGIN_AWARD, xdlog.ET_EVENT_CENTER, "")
	case event_dat.LOGIN_AWARD_TYPE_PET:
		//TODO 不支持赠送灵宠契约球
		//petId := int16(battle_pet_dat.GetBattlePetWithItemId(int32(awardConfig.AwardId)).PetId)
		//module.BattlePet.AddBattlePetBall(state.Database, petId, int16(awardConfig.AwardNum))
	default:
		panic("unsupport award type")
	}
}
