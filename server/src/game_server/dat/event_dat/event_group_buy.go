package event_dat

import (
	"core/mysql"
)

var (
	listEventsGroupBuy []*EventsGroupBuy
)

type EventsGroupBuy struct {
	Id          int16   // ID标识
	ItemId      int16   // 参与团购得物品id
	BasePrice   float32 // 团购物品底价
	BuyTimes1   int16   // 购买次数1
	BuyPercent1 float32 // 购买折扣1
	BuyTimes2   int16   // 购买次数2
	BuyPercent2 float32 // 购买折扣2
	BuyTimes3   int16   // 购买次数3
	BuyPercent3 float32 // 购买折扣3
	BuyTimes4   int16   // 购买次数4
	BuyPercent4 float32 // 购买折扣4
	BuyTimes5   int16   // 购买次数5
	BuyPercent5 float32 // 购买折扣5
	BuyTimes6   int16   // 购买次数6
	BuyPercent6 float32 // 购买折扣6
}

func loadEventsGroupBuy(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM events_group_buy ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iId := res.Map("id")
	iItemId := res.Map("item_id")
	iBasePrice := res.Map("base_price")
	iBuyTimes1 := res.Map("buy_times1")
	iBuyPercent1 := res.Map("buy_percent1")
	iBuyTimes2 := res.Map("buy_times2")
	iBuyPercent2 := res.Map("buy_percent2")
	iBuyTimes3 := res.Map("buy_times3")
	iBuyPercent3 := res.Map("buy_percent3")
	iBuyTimes4 := res.Map("buy_times4")
	iBuyPercent4 := res.Map("buy_percent4")
	iBuyTimes5 := res.Map("buy_times5")
	iBuyPercent5 := res.Map("buy_percent5")
	iBuyTimes6 := res.Map("buy_times6")
	iBuyPercent6 := res.Map("buy_percent6")

	var pri_id int16
	listEventsGroupBuy = make([]*EventsGroupBuy, 0)
	for _, row := range res.Rows {
		pri_id = row.Int16(iId)
		listEventsGroupBuy = append(listEventsGroupBuy, &EventsGroupBuy{
			Id:          pri_id,
			ItemId:      row.Int16(iItemId),
			BasePrice:   row.Float32(iBasePrice),
			BuyTimes1:   row.Int16(iBuyTimes1),
			BuyPercent1: row.Float32(iBuyPercent1),
			BuyTimes2:   row.Int16(iBuyTimes2),
			BuyPercent2: row.Float32(iBuyPercent2),
			BuyTimes3:   row.Int16(iBuyTimes3),
			BuyPercent3: row.Float32(iBuyPercent3),
			BuyTimes4:   row.Int16(iBuyTimes4),
			BuyPercent4: row.Float32(iBuyPercent4),
			BuyTimes5:   row.Int16(iBuyTimes5),
			BuyPercent5: row.Float32(iBuyPercent5),
			BuyTimes6:   row.Int16(iBuyTimes6),
			BuyPercent6: row.Float32(iBuyPercent6),
		})
	}
	eventGroupBuyItem = listEventsGroupBuy[0]
}

type EventGroupBuyExt struct {
	StartUnixTime   int64
	EndUnixTime     int64
	DisposeUnixTime int64
	IsRelative      int8
	LTitle          string
	RTitle          string
	Content         string
	Tag             int8
	Weight          int16
	ItemId          int16   // 参与团购得物品id
	BasePrice       float32 // 团购物品底价
	BuyTimes1       int16   // 购买次数1
	BuyPercent1     float32 // 购买折扣1
	BuyTimes2       int16   // 购买次数2
	BuyPercent2     float32 // 购买折扣2
	BuyTimes3       int16   // 购买次数3
	BuyPercent3     float32 // 购买折扣3
	BuyTimes4       int16   // 购买次数4
	BuyPercent4     float32 // 购买折扣4
	BuyTimes5       int16   // 购买次数5
	BuyPercent5     float32 // 购买折扣5
	BuyTimes6       int16   // 购买次数6
	BuyPercent6     float32 // 购买折扣6
}

// ############### 对外接口实现 coding here ###############
var eventGroupBuyItem *EventsGroupBuy

func LoadEventGroupBuyInfo(groupBuyInfo *EventGroupBuyExt) {
	if groupBuyInfo.ItemId > 0 && groupBuyInfo.BasePrice > 0 {
		eventGroupBuyItem.ItemId = groupBuyInfo.ItemId
		eventGroupBuyItem.BasePrice = groupBuyInfo.BasePrice
		eventGroupBuyItem.BuyTimes1 = groupBuyInfo.BuyTimes1
		eventGroupBuyItem.BuyPercent1 = groupBuyInfo.BuyPercent1
		eventGroupBuyItem.BuyTimes2 = groupBuyInfo.BuyTimes2
		eventGroupBuyItem.BuyPercent2 = groupBuyInfo.BuyPercent2
		eventGroupBuyItem.BuyTimes3 = groupBuyInfo.BuyTimes3
		eventGroupBuyItem.BuyPercent3 = groupBuyInfo.BuyPercent3
		eventGroupBuyItem.BuyTimes4 = groupBuyInfo.BuyTimes4
		eventGroupBuyItem.BuyPercent4 = groupBuyInfo.BuyPercent4
		eventGroupBuyItem.BuyTimes5 = groupBuyInfo.BuyTimes5
		eventGroupBuyItem.BuyPercent5 = groupBuyInfo.BuyPercent5
		eventGroupBuyItem.BuyTimes6 = groupBuyInfo.BuyTimes6
		eventGroupBuyItem.BuyPercent6 = groupBuyInfo.BuyPercent6
	}
}

func GetCostByGroupBuyCount(count int32) int32 {
	listOfTimes := []int16{eventGroupBuyItem.BuyTimes2, eventGroupBuyItem.BuyTimes3, eventGroupBuyItem.BuyTimes4, eventGroupBuyItem.BuyTimes5, eventGroupBuyItem.BuyTimes6}
	listOfPercent := []float32{eventGroupBuyItem.BuyPercent2, eventGroupBuyItem.BuyPercent3, eventGroupBuyItem.BuyPercent4, eventGroupBuyItem.BuyPercent5, eventGroupBuyItem.BuyPercent6}
	var percent float32 = eventGroupBuyItem.BuyPercent1
	for key, times := range listOfTimes {
		if times > 0 && int32(times) <= count {
			percent = listOfPercent[key]
		} else {
			break
		}
	}
	return int32(float32(eventGroupBuyItem.BasePrice) * percent)
}

func GetEventGroupBuyInfo() *EventsGroupBuy {
	return eventGroupBuyItem
}
