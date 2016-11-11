package event_dat

import (
	"core/mysql"
	"core/time"
	"sort"
)

var (
	eventFightPower *EventFightPower
)

type EventFightPowerAward struct {
	//Lock       int32 // 档位
	FightPower int32 // 战力
	EventDefaultAward
}

type EventFightPower struct {
	StartUnixTime   int64
	EndUnixTime     int64
	DisposeUnixTime int64
	IsRelative      int8
	LTitle          string
	RTitle          string
	Content         string
	Tag             int8
	Weight          int16
	List            []*EventFightPowerAward
}

type byFightPower []*EventFightPowerAward

func (a byFightPower) Len() int           { return len(a) }
func (a byFightPower) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byFightPower) Less(i, j int) bool { return a[i].FightPower < a[j].FightPower }

func loadEventsFightPower(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM events_fight_power ORDER BY `fight` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iFight := res.Map("fight")
	iIngot := res.Map("ingot")
	iItem1Id := res.Map("item1_id")
	iItem1Num := res.Map("item1_num")
	iItem2Id := res.Map("item2_id")
	iItem2Num := res.Map("item2_num")
	iItem3Id := res.Map("item3_id")
	iItem3Num := res.Map("item3_num")

	if eventFightPower == nil {
		eventFightPower = &EventFightPower{}
	}

	for _, row := range res.Rows {
		award := &EventFightPowerAward{}
		award.FightPower = row.Int32(iFight)
		award.Ingot = row.Int16(iIngot)
		award.Item1Id = row.Int16(iItem1Id)
		award.Item1Num = row.Int16(iItem1Num)
		award.Item2Id = row.Int16(iItem2Id)
		award.Item2Num = row.Int16(iItem2Num)
		award.Item3Id = row.Int16(iItem3Id)
		award.Item3Num = row.Int16(iItem3Num)
		eventFightPower.List = append(eventFightPower.List, award)
	}
}

func GetEventFightPowerAwards() []*EventFightPowerAward {
	return eventFightPower.List
}

func EnableEventFightPower() (ok bool) {
	if eventFightPower != nil {
		nowTime := time.GetNowTime()
		ok = nowTime >= eventFightPower.StartUnixTime && nowTime < eventFightPower.EndUnixTime
	}
	return
}

func GetEventFightAwardPower(fightPower int32) (power int32) {
	// 找到fightpower对应的奖励
	for _, item := range eventFightPower.List {
		if item.FightPower > fightPower {
			break
		}
		if power <= item.FightPower {
			power = item.FightPower
		}
	}
	return
}

// 战力活动数据加载
func LoadEventFightPower(data *EventFightPower) {
	if eventFightPower == nil {
		eventFightPower = data
	} else {
		eventFightPower.StartUnixTime = data.StartUnixTime
		eventFightPower.EndUnixTime = data.EndUnixTime
		eventFightPower.DisposeUnixTime = data.DisposeUnixTime

		var doUpdate bool
		for _, newItem := range data.List {
			doUpdate = false
			for idx, item := range eventFightPower.List {
				if item.FightPower == newItem.FightPower { // 更新
					eventFightPower.List[idx] = newItem
					doUpdate = true
					break
				}
			}
			if !doUpdate {
				eventFightPower.List = append(eventFightPower.List, newItem)
			}
		}
	}

	if eventFightPower != nil {
		sort.Sort(byFightPower(eventFightPower.List))
	}
}

func GetMaxWeightInStrong() (max int32) {
	max = 0
	for _, item := range eventFightPower.List {
		if item.FightPower > max {
			max = item.FightPower
		}
	}
	return
}
func GetNextStrong(now int32) (next int32) {
	for _, item := range eventFightPower.List {
		if item.FightPower > now {
			next = item.FightPower
			break
		}
	}
	return
}

func GetPlayerStrongAward(awarded, maxAward int32) (newAwarded, newAward int32, awards *EventFightPowerAward) {
	if awarded > maxAward {
		return //战力活动可能因战力减少而发生这种情况
	}
	found := false
	for _, item := range eventFightPower.List {
		if item.FightPower <= awarded {
			continue
		} else {
			if found == true {
				newAward = item.FightPower
				break
			} else {
				newAwarded = item.FightPower
				awards = item
				found = true
			}
		}
	}
	return
}
