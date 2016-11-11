package daily_sign_in

import (
	"core/fail"
	"core/net"
	"core/time"
	"game_server/api/protocol/daily_sign_in_api"
	//"game_server/dat/battle_pet_dat"
	"game_server/dat/daily_sign_in_dat"
	"game_server/dat/event_dat"
	"game_server/dat/player_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/tlog"
)

func info(session *net.Session, out *daily_sign_in_api.Info_Out) {
	state := module.State(session)

	playerInfo := state.Database.Lookup.PlayerInfo(state.PlayerId)
	firstLoginDate := time.GetNowDayFromUnix(playerInfo.FirstLoginTime)
	nowDate := time.GetNowDay()

	playerSignInState, _ := updateSignInState(state.Database)
	index := currentIndex(state.Database, playerSignInState) //[0, DISPLAY_RECORD_NUM)

	//当天格子索引是 7 时，假如当天尚未领取，则放回给客户的的索引是 7
	//如果已经领取，由于此时需要翻页，即最后一个格子变成第一个格子，此时返回0。
	if index == daily_sign_in_dat.DISPLAY_RECORD_NUM-1 && playerSignInState.SignedToday == 1 {
		playerSignInState.Record = playerSignInState.Record >> uint(daily_sign_in_dat.DISPLAY_RECORD_NUM-1)
		index = 0
	}
	indexInt := int(index)

	for i := uint(0); i < daily_sign_in_dat.DISPLAY_RECORD_NUM; i++ {
		//fmt.Println("hit", nowDate, firstLoginDate, indexInt, i)
		targetDate := nowDate - indexInt + int(i)
		//fmt.Println(targetDate, "targetData")
		var awardIndex int8
		var awardType daily_sign_in_api.AwardType
		if isNewPlayer(firstLoginDate, targetDate) {
			awardIndex = daily_sign_in_dat.NewPlayerSignInAwardIndex(targetDate - firstLoginDate + 1)
			awardType = daily_sign_in_api.AWARD_TYPE_NEW_PLAYER_AWARD
		} else {
			awardIndex = daily_sign_in_dat.RegularDailySignInAwardIndex(targetDate)
			awardType = daily_sign_in_api.AWARD_TYPE_REGULAR_AWARD
		}
		signed := ((playerSignInState.Record >> i) & 1) == 1
		out.Records = append(out.Records, daily_sign_in_api.Info_Out_Records{
			AwardType:  awardType,
			AwardIndex: awardIndex + 1,
			Signed:     signed,
		})
	}
	out.Index = index
}

func sign(session *net.Session, out *daily_sign_in_api.Sign_Out, xdEventType int32) {
	state := module.State(session)
	now := time.GetNowTime()
	playerSignInState, modify := updateSignInState(state.Database)
	if modify {
		out.Expired = true
		return
	}

	index := currentIndex(state.Database, playerSignInState)
	realSign(state, index, now, playerSignInState, xdEventType)
	tlog.PlayerSystemModuleFlowLog(state.Database, tlog.SMT_DAILY_SIGN)

	playerSignInState.SignedToday = 1
	state.Database.Update.PlayerDailySignInState(playerSignInState)

	if index == daily_sign_in_dat.DISPLAY_RECORD_NUM-1 {
		playerSignInState.Record = playerSignInState.Record >> uint(daily_sign_in_dat.DISPLAY_RECORD_NUM-1)
	}

	out.Index = index
	nowDate := time.GetNowDayFromUnix(now)

	playerInfo := state.Database.Lookup.PlayerInfo(state.PlayerId)
	firstLoginDate := time.GetNowDayFromUnix(playerInfo.FirstLoginTime)

	for i := uint(0); i < daily_sign_in_dat.DISPLAY_RECORD_NUM; i++ {
		targetDate := nowDate - int(index) + int(i)
		var awardIndex int8
		var awardType daily_sign_in_api.AwardType
		if isNewPlayer(firstLoginDate, targetDate) {
			awardIndex = daily_sign_in_dat.NewPlayerSignInAwardIndex(targetDate - firstLoginDate + 1)
			awardType = daily_sign_in_api.AWARD_TYPE_NEW_PLAYER_AWARD
		} else {
			awardIndex = daily_sign_in_dat.RegularDailySignInAwardIndex(targetDate)
			awardType = daily_sign_in_api.AWARD_TYPE_REGULAR_AWARD
		}
		signed := ((playerSignInState.Record >> i) & 1) == 1
		out.Records = append(out.Records, daily_sign_in_api.Sign_Out_Records{
			AwardType:  awardType,
			AwardIndex: awardIndex + 1,
			Signed:     signed,
		})
	}
}

func signPastDay(session *net.Session, index int8, out *daily_sign_in_api.SignPastDay_Out, xdEventType int32) {
	state := module.State(session)
	now := time.GetNowTime()
	playerSignInState, modify := updateSignInState(state.Database)
	if modify {
		out.Expired = true
		return
	}
	realSign(state, index, now, playerSignInState, xdEventType)
	tlog.PlayerSystemModuleFlowLog(state.Database, tlog.SMT_DAILY_SIGN_PAST)
}

func currentIndex(db *mdb.Database, playerSignInState *mdb.PlayerDailySignInState) int8 {
	playerInfo := db.Lookup.PlayerInfo(db.PlayerId())
	firstLoginDate := time.GetNowDayFromUnix(playerInfo.FirstLoginTime)
	nowDate := time.GetNowDay()
	if nowDate == firstLoginDate {
		return 0
	}

	nowRound := (nowDate - firstLoginDate - 1) / (daily_sign_in_dat.DISPLAY_RECORD_NUM - 1)
	index := nowDate - firstLoginDate - (daily_sign_in_dat.DISPLAY_RECORD_NUM-1)*nowRound

	return int8(index)
}

//更新时间戳和签到记录
func updateSignInState(db *mdb.Database) (*mdb.PlayerDailySignInState, bool) {
	playerSignInState := db.Lookup.PlayerDailySignInState(db.PlayerId())

	if time.IsToday(playerSignInState.UpdateTime) {
		return playerSignInState, false
	}

	playerInfo := db.Lookup.PlayerInfo(db.PlayerId())
	firstLoginDate := time.GetNowDayFromUnix(playerInfo.FirstLoginTime) //like 20140303
	now := time.GetNowTime()
	nowDate := time.GetNowDayFromUnix(now)
	lastestSignDate := time.GetNowDayFromUnix(playerSignInState.UpdateTime)

	//避免 首轮的前几天不需要刷新 Record，这里避免preGroup 变为负数
	if nowDate >= firstLoginDate+2 {

		newGroup := (nowDate - firstLoginDate - 1) / (daily_sign_in_dat.DISPLAY_RECORD_NUM - 1)
		preGroup := (lastestSignDate - firstLoginDate - 1) / (daily_sign_in_dat.DISPLAY_RECORD_NUM - 1)
		if newGroup > preGroup {
			//需要刷新Record字段
			playerSignInState.Record = playerSignInState.Record >> uint(((newGroup - preGroup) * (daily_sign_in_dat.DISPLAY_RECORD_NUM - 1)))
			db.Update.PlayerDailySignInState(playerSignInState)
		}
	}

	index := currentIndex(db, playerSignInState) //[0, DISPLAY_RECORD_NUM)
	bitMask := uint16(((1 << uint16(index+1)) - 1))
	dirtyRecord := uint16(playerSignInState.Record)
	playerSignInState.Record = int16(dirtyRecord & bitMask)

	playerSignInState.SignedToday = 0
	playerSignInState.UpdateTime = now
	db.Update.PlayerDailySignInState(playerSignInState)
	return playerSignInState, true
}

//index [0, DISPLAY_RECORD_NUM)
//now 时间戳
func realSign(state *module.SessionState, index int8, now int64, playerSignInState *mdb.PlayerDailySignInState, xdEventType int32) {
	fail.When(index < 0 || index >= daily_sign_in_dat.DISPLAY_RECORD_NUM, "error index")
	signed := ((playerSignInState.Record >> uint(index)) & 1) == 1
	fail.When(signed, "已签到")

	nowIndex := currentIndex(state.Database, playerSignInState)
	fail.When(index > nowIndex, "领取未来的奖励")
	if index < nowIndex {
		module.Player.DecMoney(state.Database, state.MoneyState, daily_sign_in_dat.SIGN_PAST_DAY_PRICE, player_dat.INGOT, tlog.MFR_DAILY_SIGN, xdEventType)
	}

	playerSignInState.Record += 1 << uint(index)
	state.Database.Update.PlayerDailySignInState(playerSignInState)
	state.Database.Insert.PlayerDailySignInRecord(&mdb.PlayerDailySignInRecord{
		Pid:        state.PlayerId,
		SignInTime: now - int64(nowIndex-index)*86400,
	})

	nowDate := time.GetNowDayFromUnix(now)
	targetDate := nowDate - int(nowIndex-index)

	playerInfo := state.Database.Lookup.PlayerInfo(state.PlayerId)
	firstLoginDate := time.GetNowDayFromUnix(playerInfo.FirstLoginTime) //like 20140303

	var awardConfig *daily_sign_in_dat.DailySignInAward
	if isNewPlayer(firstLoginDate, targetDate) {
		awardConfig = daily_sign_in_dat.NewPlayerSignInAward(targetDate - firstLoginDate + 1)
	} else {
		awardConfig = daily_sign_in_dat.RegularPlayerSignInAward(targetDate)
	}
	award(state, awardConfig, xdEventType)
}

func award(state *module.SessionState, awardConfig *daily_sign_in_dat.DailySignInAward, xdEventType int32) {
	vipDouble := int32(1)
	if awardConfig.VIPDouble && module.VIP.IsVIP(state.Database) {
		vipDouble = 2
	}
	switch awardConfig.AwardType {
	case daily_sign_in_dat.ITEM:
		module.Item.AddItem(state.Database, awardConfig.AwardId, int16(awardConfig.Num*vipDouble), tlog.IFR_DAILY_SIGN, xdEventType, "")
	case daily_sign_in_dat.SWORD_SOUL:
		module.SwordSoul.AddSwordSoul(state, awardConfig.AwardId, tlog.IFR_DAILY_SIGN)
	case daily_sign_in_dat.GHOST:
		module.Ghost.AddGhost(state, awardConfig.AwardId, tlog.IFR_DAILY_SIGN, xdEventType)
	case daily_sign_in_dat.HEART:
		module.Heart.IncHeart(state, int16(awardConfig.Num*vipDouble))
	case daily_sign_in_dat.COINS:
		module.Player.IncMoney(state.Database, state.MoneyState, int64(awardConfig.Num*vipDouble), player_dat.COINS, tlog.MFR_DAILY_SIGN, xdEventType, "")
	case daily_sign_in_dat.INGOT:
		module.Player.IncMoney(state.Database, state.MoneyState, int64(awardConfig.Num*vipDouble), player_dat.INGOT, tlog.MFR_DAILY_SIGN, xdEventType, "")
	case daily_sign_in_dat.PET:
		//petId := int16(battle_pet_dat.GetBattlePetWithItemId(int32(awardConfig.AwardId)).PetId)
		//module.BattlePet.AddBattlePetBall(state.Database, petId, int16(awardConfig.Num*vipDouble))
	default:
		panic("unsupport award type")
	}

	// 对于qq会员的在qq会员活动期间会多获得一个钱袋
	eventInfo, _ := event_dat.GetEventInfoById(event_dat.EVENT_QQVIP_GIFT_AWARDS)

	if event_dat.CheckEventTime(eventInfo, event_dat.NOT_END) {
		if (state.TencentState.QqVipStatus & 2) > 0 {
			module.Item.AddItem(state.Database, daily_sign_in_dat.COINS_BAG, 2, tlog.IFR_DAILY_SIGN, xdEventType, "")
		} else if (state.TencentState.QqVipStatus & 1) > 0 {
			module.Item.AddItem(state.Database, daily_sign_in_dat.COINS_BAG, 1, tlog.IFR_DAILY_SIGN, xdEventType, "")
		}
	}
}
