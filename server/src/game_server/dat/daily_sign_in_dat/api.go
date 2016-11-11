package daily_sign_in_dat

import (
	"core/fail"
	"core/mysql"
	"core/time"
	"game_server/config"
)

const (
	DISPLAY_RECORD_NUM = 8 //展示数量
)

var (
	NewPlayerDailySignInAwardConfig []*DailySignInAward //新玩家签到奖励配置
	RegularDailySignInAwardConfig   []*DailySignInAward //常规签到奖励配置
	ServerOpenTime                  int                 //开服时间(unix_timestamp/86400)
	NewPlayerSignInDuration         int                 //新玩家签到时间（单位天）
	RegularSignInDuration           int                 //常规签到周期
)

func Load(db *mysql.Connection) {
	loadDailySignInAwardConfig(db)
}

type DailySignInAward struct {
	AwardType int8  // 奖励类型
	AwardId   int16 // 奖励物品ID
	Num       int32 // 奖励数量
	VIPDouble bool  //VIP用户双倍奖励
}

func loadDailySignInAwardConfig(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM daily_sign_in_award ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iType := res.Map("type")
	iAwardType := res.Map("award_type")
	iAwardId := res.Map("award_id")
	iNum := res.Map("num")
	iVIPDouble := res.Map("vip_double")

	var signType int8
	NewPlayerDailySignInAwardConfig = make([]*DailySignInAward, 0, 14)
	RegularDailySignInAwardConfig = make([]*DailySignInAward, 0, 28)
	for _, row := range res.Rows {
		signType = row.Int8(iType)
		switch signType {
		case NEW_PLAYER_SIGN_IN_AWARD:
			NewPlayerDailySignInAwardConfig = append(NewPlayerDailySignInAwardConfig, &DailySignInAward{
				AwardType: row.Int8(iAwardType),
				AwardId:   row.Int16(iAwardId),
				Num:       row.Int32(iNum),
				VIPDouble: row.Bool(iVIPDouble),
			})
		case REGULAR_SIGN_IN_AWARD:
			RegularDailySignInAwardConfig = append(RegularDailySignInAwardConfig, &DailySignInAward{
				AwardType: row.Int8(iAwardType),
				AwardId:   row.Int16(iAwardId),
				Num:       row.Int32(iNum),
				VIPDouble: row.Bool(iVIPDouble),
			})
		default:
			panic("未定义得签到类型")
		}
	}
	ServerOpenTime = time.GetNowDayFromUnix(config.ServerCfg.ServerOpenTime)
	NewPlayerSignInDuration = len(NewPlayerDailySignInAwardConfig)
	RegularSignInDuration = len(RegularDailySignInAwardConfig)
}

//传新手签到天数，登录第一天则传参数day＝1
func NewPlayerSignInAward(day int) *DailySignInAward {
	fail.When(day < 1 || day > NewPlayerSignInDuration, "非法每日新手签到日期索引")
	return NewPlayerDailySignInAwardConfig[day-1]
}

func NewPlayerSignInAwardIndex(day int) int8 {
	fail.When(day < 1 || day > NewPlayerSignInDuration, "非法每日新手签到日期索引")
	return int8(day - 1)
}

func RegularPlayerSignInAward(day int) *DailySignInAward {
	/*
		例如开服是 1 日， 有 7 天新手奖励， 8日开始常规签到
		开启： 	8 - 1 < 7 false
		未开启：7 - 1 < 7 true
	*/
	fail.When(day-ServerOpenTime < NewPlayerSignInDuration, "此日期常规签到尚未开启")
	index := (day - ServerOpenTime - NewPlayerSignInDuration) % RegularSignInDuration
	return RegularDailySignInAwardConfig[index]
}

func RegularDailySignInAwardIndex(day int) int8 {
	fail.When(day-ServerOpenTime < NewPlayerSignInDuration, "此日期常规签到尚未开启")
	index := int8((day - ServerOpenTime - NewPlayerSignInDuration) % RegularSignInDuration)
	return index
}

type EventDailySign struct {
	SignType int // 签到类型
	Day      int
	DailySignInAward
}

// 加载每日签到的运营配置数据
func LoadEventDailySignData(list []*EventDailySign) {
	for _, item := range list {
		if item.Day-1 >= 0 {
			if item.SignType == NEW_PLAYER_SIGN_IN_AWARD {
				NewPlayerDailySignInAwardConfig[item.Day-1] = &item.DailySignInAward
			} else if item.SignType == REGULAR_SIGN_IN_AWARD {
				RegularDailySignInAwardConfig[item.Day-1] = &item.DailySignInAward
			}
		}
	}
}
