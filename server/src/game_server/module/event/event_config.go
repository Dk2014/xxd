package event

import (
	"C"
	"core/fail"
	"core/log"
	"core/mysql"
	"encoding/json"
	"game_server/config"
	. "game_server/config"
	"game_server/dat/daily_sign_in_dat"
	"game_server/dat/event_dat"
	"game_server/global"
	"unsafe"
)

var EventConfigRawData []byte
var eventExt event_dat.EventConfigExtend

func getVersion(db *mysql.Connection) (version int32, eventConfigRaw string) {
	res, err := db.ExecuteFetch([]byte(`select * from server_info`), -1)
	fail.When(err != nil, err)

	iVersion := res.Map("event_version")
	iEventConfig := res.Map("event_json_info")

	for _, row := range res.Rows {
		version = row.Int32(iVersion)
		eventConfigRaw = row.Str(iEventConfig)
	}

	return
}

func updateVersion(db *mysql.Connection, version int32, eventConfigStr string) {
	stmt, err := db.Prepare([]byte("UPDATE `server_info` SET `event_version`=?, `event_json_info`=?"))
	fail.When(err != nil, err)
	stmt.BindInt(unsafe.Pointer(&version))
	str := C.CString(eventConfigStr)
	stmt.BindVarchar(unsafe.Pointer(str), len(eventConfigStr))
	stmt.Execute()
	stmt.Close()
}

func LoadEventExtend(conf string) {
	fail.When(conf == "", "event.json not given")

	EventConfigRawData = config.LoadJSONConfig(conf, &eventExt)

	db, err := mysql.Connect(config.GetDBConfig())
	fail.When(err != nil, err)
	defer db.Close()

	// 不能更新比当前json版本低的配置文件。允许更新的数据版本和当前的版本相同（多个gs进程）
	// 2015-06-02 修改 因为rpc可能修改活动配置，导致老的event.json版本低于数据库的server_info记录的版本
	version, EventConfigRaw := getVersion(db)
	if version > eventExt.Version {
		log.Warnf("event.json version(%d) less than event db version(%d), ", eventExt.Version, version)
		EventConfigRawData = []byte(EventConfigRaw)
		err = json.Unmarshal(EventConfigRawData, &eventExt)
		fail.When(err != nil, err)
	}
	//fail.When(eventExt.Version < getVersion(db), "更新的运营数据版本不正确. (<当前版本)")

	// 根据配置数据初始化对应的运营功能数据
	// event_dat.LoadTotalPhysicalCost(eventExt.TotalPhysicalCost)
	daily_sign_in_dat.LoadEventDailySignData(eventExt.DailySignIn)
	event_dat.LoadEventsLevelUp(eventExt.LevelUp.List)
	event_dat.LoadEventFightPower(eventExt.FightPower)
	event_dat.LoadEventsLogin(eventExt.LoginIn.List)
	event_dat.LoadEventsArenaRank(eventExt.ArenaRank.List)
	event_dat.LoadEventMultipy(eventExt.Multipy.List)
	event_dat.LoadEventsPhysical(eventExt.Physical.List)
	event_dat.LoadEventsDinner(eventExt.Dinner.List)
	event_dat.LoadEventsMonthCard(eventExt.MonthCard.List)
	event_dat.LoadEventsSupper(eventExt.Supper.List)
	event_dat.LoadEventsQQVIPGift(eventExt.QQVIPGift.List)
	event_dat.LoadEventMultipy(eventExt.QQVIPGift.Addition)
	event_dat.LoadEventsVipClub(eventExt.VipClub.List)
	event_dat.LoadEventsSevenDay(eventExt.SevenDay.List)
	event_dat.LoadEventsShare(eventExt.Share.List)
	event_dat.LoadEventsBuyPartner(eventExt.BuyPartner.List)
	event_dat.LoadEventsTotalRecharge(eventExt.TotalRecharge.List)
	event_dat.LoadEventsTenDraw(eventExt.TenDraw.List)
	event_dat.LoadEventTotalConsume(eventExt.TotalConsume.List)
	event_dat.LoadEventsFirstRechargeDaily(eventExt.FirstRechargeDaily.List)
	event_dat.LoadEventsTotalLogin(eventExt.TotalLogin.List)
	event_dat.LoadEventGroupBuyInfo(eventExt.GroupBuy)
	//根据配置数据初始化对应的活动中心配置信息的运营数据
	event_dat.LoadEventCenterExt(eventExt.TotalConsume.StartUnixTime, eventExt.TotalConsume.EndUnixTime, eventExt.TotalConsume.DisposeUnixTime, event_dat.EVENT_TOTAL_CONSUME, eventExt.TotalConsume.IsRelative, eventExt.TotalConsume.LTitle, eventExt.TotalConsume.RTitle, eventExt.TotalConsume.Content, eventExt.TotalConsume.Weight, eventExt.TotalConsume.Tag)           //累计消费活动
	event_dat.LoadEventCenterExt(eventExt.ArenaRank.StartUnixTime, eventExt.ArenaRank.EndUnixTime, eventExt.ArenaRank.DisposeUnixTime, event_dat.EVENT_ARENA_RANK_AWARDS, eventExt.ArenaRank.IsRelative, eventExt.ArenaRank.LTitle, eventExt.ArenaRank.RTitle, eventExt.ArenaRank.Content, eventExt.ArenaRank.Weight, eventExt.ArenaRank.Tag)                                  //比武场排名活动
	event_dat.LoadEventCenterExt(eventExt.Dinner.StartUnixTime, eventExt.Dinner.EndUnixTime, eventExt.Dinner.DisposeUnixTime, event_dat.EVENT_DINNER_AWARDS, eventExt.Dinner.IsRelative, eventExt.Dinner.LTitle, eventExt.Dinner.RTitle, eventExt.Dinner.Content, eventExt.Dinner.Weight, eventExt.Dinner.Tag)                                                                 //午餐活动
	event_dat.LoadEventCenterExt(eventExt.LoginIn.StartUnixTime, eventExt.LoginIn.EndUnixTime, eventExt.LoginIn.DisposeUnixTime, event_dat.EVENT_LOGIN_AWARD, eventExt.LoginIn.IsRelative, eventExt.LoginIn.LTitle, eventExt.LoginIn.RTitle, eventExt.LoginIn.Content, eventExt.LoginIn.Weight, eventExt.LoginIn.Tag)                                                          //登录活动
	event_dat.LoadEventCenterExt(eventExt.LevelUp.StartUnixTime, eventExt.LevelUp.EndUnixTime, eventExt.LevelUp.DisposeUnixTime, event_dat.EVENT_LEVEL_AWARD, eventExt.LevelUp.IsRelative, eventExt.LevelUp.LTitle, eventExt.LevelUp.RTitle, eventExt.LevelUp.Content, eventExt.LevelUp.Weight, eventExt.LevelUp.Tag)                                                          //升级活动
	event_dat.LoadEventCenterExt(eventExt.FightPower.StartUnixTime, eventExt.FightPower.EndUnixTime, eventExt.FightPower.DisposeUnixTime, event_dat.EVENT_STRONG_AWARD, eventExt.FightPower.IsRelative, eventExt.FightPower.LTitle, eventExt.FightPower.RTitle, eventExt.FightPower.Content, eventExt.FightPower.Weight, eventExt.FightPower.Tag)                              //战力活动
	event_dat.LoadEventCenterExt(eventExt.Recharge.StartUnixTime, eventExt.Recharge.EndUnixTime, eventExt.Recharge.DisposeUnixTime, event_dat.EVENT_RECHARGE_AWARD, eventExt.Recharge.IsRelative, eventExt.Recharge.LTitle, eventExt.Recharge.RTitle, eventExt.Recharge.Content, eventExt.Recharge.Weight, eventExt.Recharge.Tag)                                              //充值返利活动
	event_dat.LoadEventCenterExt(eventExt.TotalRecharge.StartUnixTime, eventExt.TotalRecharge.EndUnixTime, eventExt.TotalRecharge.DisposeUnixTime, event_dat.EVENT_TOTAL_RECHARGE, eventExt.TotalRecharge.IsRelative, eventExt.TotalRecharge.LTitle, eventExt.TotalRecharge.RTitle, eventExt.TotalRecharge.Content, eventExt.TotalRecharge.Weight, eventExt.TotalRecharge.Tag) //累计充值元宝活动
	event_dat.LoadEventCenterExt(eventExt.Multipy.StartUnixTime, eventExt.Multipy.EndUnixTime, eventExt.Multipy.DisposeUnixTime, event_dat.EVENT_MULTIPY_CONFIG, eventExt.Multipy.IsRelative, eventExt.Multipy.LTitle, eventExt.Multipy.RTitle, eventExt.Multipy.Content, eventExt.Multipy.Weight, eventExt.Multipy.Tag)                                                       //翻倍奖励活动
	event_dat.LoadEventCenterExt(eventExt.Physical.StartUnixTime, eventExt.Physical.EndUnixTime, eventExt.Physical.DisposeUnixTime, event_dat.EVENT_PHYSICAL_AWARDS, eventExt.Physical.IsRelative, eventExt.Physical.LTitle, eventExt.Physical.RTitle, eventExt.Physical.Content, eventExt.Physical.Weight, eventExt.Physical.Tag)                                             //活跃度活动
	event_dat.LoadEventCenterExt(eventExt.MonthCard.StartUnixTime, eventExt.MonthCard.EndUnixTime, eventExt.MonthCard.DisposeUnixTime, event_dat.EVENT_MONTH_CARD_AWARDS, eventExt.MonthCard.IsRelative, eventExt.MonthCard.LTitle, eventExt.MonthCard.RTitle, eventExt.MonthCard.Content, eventExt.MonthCard.Weight, eventExt.MonthCard.Tag)                                  //月卡活动
	event_dat.LoadEventCenterExt(eventExt.Supper.StartUnixTime, eventExt.Supper.EndUnixTime, eventExt.Supper.DisposeUnixTime, event_dat.EVENT_SUPPER_AWARDS, eventExt.Supper.IsRelative, eventExt.Supper.LTitle, eventExt.Supper.RTitle, eventExt.Supper.Content, eventExt.Supper.Weight, eventExt.Supper.Tag)                                                                 //晚餐活动
	event_dat.LoadEventCenterExt(eventExt.QQVIPGift.StartUnixTime, eventExt.QQVIPGift.EndUnixTime, eventExt.QQVIPGift.DisposeUnixTime, event_dat.EVENT_QQVIP_GIFT_AWARDS, eventExt.QQVIPGift.IsRelative, eventExt.QQVIPGift.LTitle, eventExt.QQVIPGift.RTitle, eventExt.QQVIPGift.Content, eventExt.QQVIPGift.Weight, eventExt.QQVIPGift.Tag)                                  //QQ特权赠送物品
	event_dat.LoadEventCenterExt(eventExt.VipClub.StartUnixTime, eventExt.VipClub.EndUnixTime, eventExt.VipClub.DisposeUnixTime, event_dat.EVENT_VIP_CLUB, eventExt.VipClub.IsRelative, eventExt.VipClub.LTitle, eventExt.VipClub.RTitle, eventExt.VipClub.Content, eventExt.VipClub.Weight, eventExt.VipClub.Tag)                                                             //VIP俱乐部
	event_dat.LoadEventCenterExt(eventExt.LevelRecharge.StartUnixTime, eventExt.LevelRecharge.EndUnixTime, eventExt.LevelRecharge.DisposeUnixTime, event_dat.EVENT_LEVEL_RECHARGE, eventExt.LevelRecharge.IsRelative, eventExt.LevelRecharge.LTitle, eventExt.LevelRecharge.RTitle, eventExt.LevelRecharge.Content, eventExt.LevelRecharge.Weight, eventExt.LevelRecharge.Tag) //冲级返利
	event_dat.LoadEventCenterExt(eventExt.FirstRecharge.StartUnixTime, eventExt.FirstRecharge.EndUnixTime, eventExt.FirstRecharge.DisposeUnixTime, event_dat.EVENT_FIRST_RECHARGE, eventExt.FirstRecharge.IsRelative, eventExt.FirstRecharge.LTitle, eventExt.FirstRecharge.RTitle, eventExt.FirstRecharge.Content, eventExt.FirstRecharge.Weight, eventExt.FirstRecharge.Tag) //首充返利
	event_dat.LoadEventCenterExt(eventExt.SevenDay.StartUnixTime, eventExt.SevenDay.EndUnixTime, eventExt.SevenDay.DisposeUnixTime, event_dat.EVENT_SEVEN_DAY_AWARDS, eventExt.SevenDay.IsRelative, eventExt.SevenDay.LTitle, eventExt.SevenDay.RTitle, eventExt.SevenDay.Content, eventExt.SevenDay.Weight, eventExt.SevenDay.Tag)                                            //新手七天乐
	event_dat.LoadEventCenterExt(eventExt.Richman.StartUnixTime, eventExt.Richman.EndUnixTime, eventExt.Richman.DisposeUnixTime, event_dat.EVENT_RICHMAN_CLUB, eventExt.Richman.IsRelative, eventExt.Richman.LTitle, eventExt.Richman.RTitle, eventExt.Richman.Content, eventExt.Richman.Weight, eventExt.Richman.Tag)                                                         //土豪俱乐部活动
	event_dat.LoadEventCenterExt(eventExt.Share.StartUnixTime, eventExt.Share.EndUnixTime, eventExt.Share.DisposeUnixTime, event_dat.EVENT_SHARE_AWARDS, eventExt.Share.IsRelative, eventExt.Share.LTitle, eventExt.Share.RTitle, eventExt.Share.Content, eventExt.Share.Weight, eventExt.Share.Tag)                                                                           //分享活动
	event_dat.LoadEventCenterExt(eventExt.GroupBuy.StartUnixTime, eventExt.GroupBuy.EndUnixTime, eventExt.GroupBuy.DisposeUnixTime, event_dat.EVENT_GROUP_BUY, eventExt.GroupBuy.IsRelative, eventExt.GroupBuy.LTitle, eventExt.GroupBuy.RTitle, eventExt.GroupBuy.Content, eventExt.GroupBuy.Weight, eventExt.GroupBuy.Tag)
	event_dat.LoadEventCenterExt(eventExt.BuyPartner.StartUnixTime, eventExt.BuyPartner.EndUnixTime, eventExt.BuyPartner.DisposeUnixTime, event_dat.EVENT_BUY_PARTNER, eventExt.BuyPartner.IsRelative, eventExt.BuyPartner.LTitle, eventExt.BuyPartner.RTitle, eventExt.BuyPartner.Content, eventExt.BuyPartner.Weight, eventExt.BuyPartner.Tag)                                                                                  //买卖伙伴
	event_dat.LoadEventCenterExt(eventExt.TenDraw.StartUnixTime, eventExt.TenDraw.EndUnixTime, eventExt.TenDraw.DisposeUnixTime, event_dat.EVENT_TEN_DRAW, eventExt.TenDraw.IsRelative, eventExt.TenDraw.LTitle, eventExt.TenDraw.RTitle, eventExt.TenDraw.Content, eventExt.TenDraw.Weight, eventExt.TenDraw.Tag)                                                                                                                //神龙宝箱十连抽
	event_dat.LoadEventCenterExt(eventExt.FirstRechargeDaily.StartUnixTime, eventExt.FirstRechargeDaily.EndUnixTime, eventExt.FirstRechargeDaily.DisposeUnixTime, event_dat.EVENT_FIRST_RECHARGE_DAILY, eventExt.FirstRechargeDaily.IsRelative, eventExt.FirstRechargeDaily.LTitle, eventExt.FirstRechargeDaily.RTitle, eventExt.FirstRechargeDaily.Content, eventExt.FirstRechargeDaily.Weight, eventExt.FirstRechargeDaily.Tag) //每日首冲奖励
	event_dat.LoadEventCenterExt(eventExt.TotalLogin.StartUnixTime, eventExt.TotalLogin.EndUnixTime, eventExt.TotalLogin.DisposeUnixTime, event_dat.EVENT_TOTAL_LOGIN, eventExt.TotalLogin.IsRelative, eventExt.TotalLogin.LTitle, eventExt.TotalLogin.RTitle, eventExt.TotalLogin.Content, eventExt.TotalLogin.Weight, eventExt.TotalLogin.Tag)                                                                                  //连续登陆奖励活动                                        //QQ特权赠送物品

	event_dat.LoadSpecialEvent(&eventExt) // 加载特殊的活动配置信息
	// LOAD JSON EVENT CONFIG
	event_dat.LoadJsonEventInfo(eventExt.JsonEvents)

	event_dat.TextEvents = eventExt.EventAnnounce

	if ServerCfg.EnableGlobalServer {
		// load互动服上面的一些全局活动信息 比方说团购活动等
		global.Start(GetDBConfig())
	}
	updateVersion(db, eventExt.Version, string(EventConfigRawData))

}
