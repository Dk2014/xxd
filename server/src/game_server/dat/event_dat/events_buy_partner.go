package event_dat

import (
	"core/mysql"
)

var (
	listEventsBuyPartner []*EventsBuyPartner
)

type EventsBuyPartner struct {
	//Id         int32 //
	PatnerId   int16 // 伙伴ID
	BuddyLevel int16 // 伙伴等级
	Cost       int64 // 价格
	SkillId1   int16 // 招式名称1
	SkillId2   int16 // 招式名称2
}

type EventsBuyPartnerpExt struct {
	StartUnixTime   int64
	EndUnixTime     int64
	DisposeUnixTime int64
	IsRelative      int8
	LTitle          string
	RTitle          string
	Content         string
	Tag             int8
	Weight          int16
	List            []*EventsBuyPartner
}

func loadEventsBuyPartner(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM events_buy_partner ORDER BY `id` DESC"), -1)
	if err != nil {
		panic(err)
	}

	iPatnerId := res.Map("patner_id")
	iBuddyLevel := res.Map("buddy_level")
	iCost := res.Map("cost")
	iSkillId1 := res.Map("skill_id1")
	iSkillId2 := res.Map("skill_id2")

	listEventsBuyPartner = make([]*EventsBuyPartner, 0)
	for _, row := range res.Rows {

		listEventsBuyPartner = append(listEventsBuyPartner, &EventsBuyPartner{
			//Id:         pri_id,
			PatnerId:   row.Int16(iPatnerId),
			BuddyLevel: row.Int16(iBuddyLevel),
			Cost:       row.Int64(iCost),
			SkillId1:   row.Int16(iSkillId1),
			SkillId2:   row.Int16(iSkillId2),
		})

	}

}

// 买卖伙伴活动运营活动数据配置
func LoadEventsBuyPartner(list []*EventsBuyPartner) {

	if len(listEventsBuyPartner) > 0 && len(list) > 0 {
		listEventsBuyPartner[0] = &EventsBuyPartner{
			PatnerId:   list[0].PatnerId,
			BuddyLevel: list[0].BuddyLevel,
			Cost:       list[0].Cost,
			SkillId1:   list[0].SkillId1,
			SkillId2:   list[0].SkillId2,
		}
	}

}

func GetPartnersCosts() int64 {
	return listEventsBuyPartner[0].Cost
}

func GetPartneInfo() (patnerid, buddyLevel int16, consts int64) {
	return listEventsBuyPartner[0].PatnerId, listEventsBuyPartner[0].BuddyLevel, listEventsBuyPartner[0].Cost
}
