package event_dat

import (
	"core/mysql"
	"sort"
)

type TotalLoginAward struct {
	Require_login_days int32
	Awards             *EventDefaultAward
}

var (
	listEventsTotalLogin []*TotalLoginAward
)

type bySortEventsTotalLoginAward []*TotalLoginAward

func (bySort bySortEventsTotalLoginAward) Len() int      { return len(bySort) }
func (bySort bySortEventsTotalLoginAward) Swap(i, j int) { bySort[i], bySort[j] = bySort[j], bySort[i] }
func (bySort bySortEventsTotalLoginAward) Less(i, j int) bool {
	return bySort[i].Require_login_days < bySort[j].Require_login_days
}

func loadEventsTotalLoginAward(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM events_total_login ORDER BY `require_login_days` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iRequireLoginDays := res.Map("require_login_days")
	iIngot := res.Map("ingot")
	iCoins := res.Map("coins")
	iHeart := res.Map("heart")
	iItem1Id := res.Map("item1_id")
	iItem1Num := res.Map("item1_num")
	iItem2Id := res.Map("item2_id")
	iItem2Num := res.Map("item2_num")
	iItem3Id := res.Map("item3_id")
	iItem3Num := res.Map("item3_num")
	iItem4Id := res.Map("item4_id")
	iItem4Num := res.Map("item4_num")
	iItem5Id := res.Map("item5_id")
	iItem5Num := res.Map("item5_num")

	var pri_require_login_days int32
	listEventsTotalLogin = make([]*TotalLoginAward, 0)

	for _, row := range res.Rows {
		pri_require_login_days = row.Int32(iRequireLoginDays)
		listEventsTotalLogin = append(listEventsTotalLogin, &TotalLoginAward{
			Require_login_days: pri_require_login_days,
			Awards: &EventDefaultAward{
				Ingot:    row.Int16(iIngot),
				Coin:     row.Int32(iCoins),
				Heart:    row.Int16(iHeart),
				Item1Id:  row.Int16(iItem1Id),
				Item1Num: row.Int16(iItem1Num),
				Item2Id:  row.Int16(iItem2Id),
				Item2Num: row.Int16(iItem2Num),
				Item3Id:  row.Int16(iItem3Id),
				Item3Num: row.Int16(iItem3Num),
				Item4Id:  row.Int16(iItem4Id),
				Item4Num: row.Int16(iItem4Num),
				Item5Id:  row.Int16(iItem5Id),
				Item5Num: row.Int16(iItem5Num),
			},
		})

	}
}

// ############### 对外接口实现 coding here ###############
func GetEventTotalLoginAward(login_days int32) (days int32, awards *EventDefaultAward) {
	for index := 0; index < len(listEventsTotalLogin); index++ {
		login_days_awards := listEventsTotalLogin[index]
		awards = login_days_awards.Awards
		days = login_days_awards.Require_login_days
		if login_days_awards.Require_login_days == login_days {
			//查到就不往下找了
			break
		} else if login_days_awards.Require_login_days > login_days {
			//第一次超越了则返回上一个，而且不再继续找下去
			if index > 0 { // 开始领奖天数从大于1的开始
				awards = listEventsTotalLogin[index-1].Awards
				days = listEventsTotalLogin[index-1].Require_login_days
			} else {
				days = listEventsTotalLogin[0].Require_login_days
				awards = nil
			}
			break
		}
	}
	return
}

func GetEventTotalLoginAwards() []*TotalLoginAward {
	return listEventsTotalLogin
}

type EventsTotalLogin struct {
	RequireLoginDays int32 // 需要QQ特权赠送物品
	Ingot            int16 // 奖励元宝
	Coin             int32 // 奖励铜钱
	Heart            int16 //奖励爱心
	Item1Id          int16 // 物品1
	Item1Num         int16 // 物品1数量
	Item2Id          int16 // 物品2
	Item2Num         int16 // 物品2数量
	Item3Id          int16 // 物品3
	Item3Num         int16 // 物品3数量
	Item4Id          int16 // 物品4
	Item4Num         int16 // 物品4数量
	Item5Id          int16 // 物品5
	Item5Num         int16 // 物品5数量
}

type EventsTotalLoginExt struct {
	StartUnixTime   int64
	EndUnixTime     int64
	DisposeUnixTime int64
	IsRelative      int8
	LTitle          string
	RTitle          string
	Content         string
	Tag             int8
	Weight          int16
	List            []*EventsTotalLogin
}

// 连续登陆运营活动数据配置
func LoadEventsTotalLogin(list []*EventsTotalLogin) {
	var isAdd bool
	for _, item := range list {
		if item.RequireLoginDays > 0 {
			isAdd = true
			for index := 0; index < len(listEventsTotalLogin); index++ {
				if listEventsTotalLogin[index].Require_login_days == item.RequireLoginDays {
					listEventsTotalLogin[index].Awards = &EventDefaultAward{
						Ingot:    item.Ingot,
						Coin:     item.Coin,
						Heart:    item.Heart,
						Item1Id:  item.Item1Id,
						Item1Num: item.Item1Num,
						Item2Id:  item.Item2Id,
						Item2Num: item.Item2Num,
						Item3Id:  item.Item3Id,
						Item3Num: item.Item3Num,
						Item4Id:  item.Item4Id,
						Item4Num: item.Item4Num,
						Item5Id:  item.Item5Id,
						Item5Num: item.Item5Num,
					}
					isAdd = false
				}
			}
			if isAdd {
				listEventsTotalLogin = append(listEventsTotalLogin, &TotalLoginAward{
					Require_login_days: item.RequireLoginDays,
					Awards: &EventDefaultAward{
						Ingot:    item.Ingot,
						Coin:     item.Coin,
						Heart:    item.Heart,
						Item1Id:  item.Item1Id,
						Item1Num: item.Item1Num,
						Item2Id:  item.Item2Id,
						Item2Num: item.Item2Num,
						Item3Id:  item.Item3Id,
						Item3Num: item.Item3Num,
						Item4Id:  item.Item4Id,
						Item4Num: item.Item4Num,
						Item5Id:  item.Item5Id,
						Item5Num: item.Item5Num,
					},
				})
			}
		}

	}

	tempList := make([]*TotalLoginAward, 0)
	for _, item := range listEventsTotalLogin {
		if IsVaildEventDefaultAward(item.Awards) {
			tempList = append(tempList, item)
		}
	}
	listEventsTotalLogin = tempList
	sort.Sort(bySortEventsTotalLoginAward(listEventsTotalLogin))
}

func GetMaxWeightInTotalLogin() int32 {
	return listEventsTotalLogin[len(listEventsTotalLogin)-1].Require_login_days
}

func GetNextTotalLogin(now int32) (next int32) {
	for index := 0; index < len(listEventsTotalLogin); index++ {
		if listEventsTotalLogin[index].Require_login_days > now {
			next = listEventsTotalLogin[index].Require_login_days
			break
		}
	}
	return
}

func GetPlayerTotalLoginAward(awarded int32) (newAwarded, nextAward int32, awards *EventDefaultAward) {
	var index int
	for index = 0; index < len(listEventsTotalLogin); index++ {
		if listEventsTotalLogin[index].Require_login_days > awarded {
			newAwarded = listEventsTotalLogin[index].Require_login_days
			awards = listEventsTotalLogin[index].Awards
			break
		}
	}
	if index < len(listEventsTotalLogin)-1 {
		nextAward = int32(listEventsTotalLogin[index+1].Require_login_days)
	}
	return
}
