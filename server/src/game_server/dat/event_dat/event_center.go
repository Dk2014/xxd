package event_dat

import (
	"core/fail"
	"core/mysql"
	"core/time"
	. "game_server/config"
	"game_server/dat/mail_dat"
)

var (
	mapEventCenter  map[int16]*EventCenter
	listEventsIndex []int16
)

type EventCenter struct {
	Id          int16  //
	Relative    int16  // 关联的活动
	Start       int64  // 活动开始时间戳
	End         int64  // 活动结束时间戳
	Dispose     int64  // 活动销毁时间戳
	IsRelative  int8   //是否为相对时间戳
	IsGo        int8   // 是否前往
	IsMail      int8   // 活动结束是否补发奖励
	MailTitle   string // 补发奖励邮件标题
	MailContent string // 补发奖励邮件内容,{val}对应权值
	LTitle      string // 左标题
	RTitle      string // 右标题
	Content     string // 活动描述
	Tag         int8
	Weight      int16
}

func loadEventCenter(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM quest_activity_center ORDER BY `weight` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iId := res.Map("id")
	iRelative := res.Map("relative")
	iStart := res.Map("start")
	iEnd := res.Map("end")
	iDispose := res.Map("dispose")
	iIsGo := res.Map("is_go")
	iIsRelative := res.Map("is_relative")
	iIsMail := res.Map("is_mail")
	iMailTitle := res.Map("mail_title")
	iMailContent := res.Map("mail_content")
	iLTitle := res.Map("name")
	iRTitle := res.Map("title")
	iContent := res.Map("content")
	iTag := res.Map("tag")
	iWeight := res.Map("weight")

	var pri_id int16
	mapEventCenter = map[int16]*EventCenter{}
	listEventsIndex = make([]int16, 0)
	for _, row := range res.Rows {
		pri_id = row.Int16(iId)
		mapEventCenter[pri_id] = &EventCenter{
			Id:          pri_id,
			Relative:    row.Int16(iRelative),
			Start:       row.Int64(iStart),
			End:         row.Int64(iEnd),
			Dispose:     row.Int64(iDispose),
			IsRelative:  row.Int8(iIsRelative),
			IsGo:        row.Int8(iIsGo),
			IsMail:      row.Int8(iIsMail),
			MailTitle:   row.Str(iMailTitle),
			MailContent: row.Str(iMailContent),
			LTitle:      row.Str(iLTitle),
			RTitle:      row.Str(iRTitle),
			Content:     row.Str(iContent),
			Tag:         row.Int8(iTag),
			Weight:      row.Int16(iWeight),
		}
		listEventsIndex = append(listEventsIndex, pri_id)
	}
}

// ############### 对外接口实现 coding here ###############
func LoadEventCenterExt(start, end, dispose int64, typeId int16, isRelative int8, LTitle string, RTitle string, Content string, Weight int16, Tag int8) {
	if event, ok := mapEventCenter[typeId]; ok {
		if start > -1 {
			event.Start = start
		}
		if end > -1 {
			event.End = end
		}
		if dispose > -1 {
			event.Dispose = dispose
		}
		if isRelative > -1 {
			event.IsRelative = isRelative
		}
		if LTitle != "" {
			event.LTitle = LTitle
		}
		if RTitle != "" {
			event.RTitle = RTitle
		}
		if Content != "" {
			event.Content = Content
		}
		if Weight > -1 {
			event.Weight = Weight
		}
		if Tag > -1 {
			event.Tag = Tag
		}
	} else {
		fail.When(true, "wrong event type")
	}
}

func GetEventsInfo() []*EventCenter {
	events := make([]*EventCenter, 0)
	for _, index := range listEventsIndex {
		if event, ok := mapEventCenter[index]; ok {
			events = append(events, event)
		}
	}
	return events
}

func GetEventInfoById(eventId int16) (event *EventCenter, ok bool) {
	event, ok = mapEventCenter[eventId]
	return
}

//验证活动操作是否还在活动有效期范围内 （分两种，活动进行期，活动存在期）
func CheckEventTime(eventInfo *EventCenter, checkType int) bool {
	dt := int64(0)
	if eventInfo.IsRelative == 1 { //是相对时间
		dt = ServerCfg.ServerOpenTime
	}
	now := time.GetNowTime()

	if eventInfo.Start+dt > 0 && eventInfo.Start+dt > now {
		return false //活动还没开始的状况
	}
	if checkType == 0 { //领取奖励时验证是否过期
		if eventInfo.Dispose+dt > 0 && now >= eventInfo.Dispose+dt {
			return false
		}
	} else if checkType == 1 { //获取活动时验证是否结束
		if eventInfo.End+dt > 0 && now >= eventInfo.End+dt {
			return false
		}
	}
	return true
}

func IsInPointEventTime(eventInfo *EventCenter, checkType int, timeSpan int64) bool {
	dt := int64(0)
	if eventInfo.IsRelative == 1 { //是相对时间
		dt = ServerCfg.ServerOpenTime
	}
	if eventInfo.Start+dt > 0 && eventInfo.Start+dt > timeSpan {
		return false //活动还没开始的状况
	}
	if checkType == 0 { //领取奖励时验证是否过期
		if eventInfo.Dispose+dt > 0 && timeSpan >= eventInfo.Dispose+dt {
			return false
		}
	} else if checkType == 1 { //获取活动时验证是否结束
		if eventInfo.End+dt > 0 && timeSpan >= eventInfo.End+dt {
			return false
		}
	}
	return true
}

////验证活动操作是否还在活动有效期范围内 （分两种，活动进行期，活动存在期）
//func VaildEvent(eventId int16, kind int8) bool {
//	eventInfo, ok := mapEventCenter[eventId]
//	fail.When(!ok, "wrong event ID")
//	result := true
//	now := time.GetNowTime()
//	if eventInfo.Start > 0 && eventInfo.Start > now {
//		result = false //活动还没开始的状况
//	}
//	if kind == 0 { //领取奖励时验证是否过期
//		if eventInfo.Dispose > 0 && now >= eventInfo.Dispose {
//			result = false
//		}
//	} else if kind == 1 { //获取活动时验证是否结束
//		if eventInfo.End > 0 && now >= eventInfo.End {
//			result = false
//		}
//	}
//	return result
//}

//过期活动需要补发邮件的，补发奖励邮件给玩家
func AwardMailAttachments(eventId int16, awarded, maxAward int32) (attachments map[int32][]*mail_dat.Attachment) {
	//TODO 添加判断逻辑
	return
}

type JsonEventAward struct {
	Grade int32 //奖励所需的档位
	Award *EventDefaultAward
}

type JsonEvent struct {
	Type            int16
	Page            int32 //期数
	StartUnixTime   int64
	EndUnixTime     int64
	DisposeUnixTime int64
	IsRelative      bool
	LTitle          string
	RTitle          string
	Content         string
	Weight          int16
	Tag             int8
	List            []*JsonEventAward
}

type JsonEventAwardItemInJsonFile struct {
	Require  int32
	Ingot    int16 // 奖励元宝
	Coin     int32 // 奖励铜钱
	Heart    int16 // 奖励爱心
	Item1Id  int16 // 物品1
	Item1Num int16 // 物品1数量
	Item2Id  int16 // 物品2
	Item2Num int16 // 物品2数量
	Item3Id  int16 // 物品3
	Item3Num int16 // 物品3数量
	Item4Id  int16 // 物品4
	Item4Num int16 // 物品4数量
	Item5Id  int16 // 物品5
	Item5Num int16 // 物品5数量
}

type JsonEventStructInJsonFile struct {
	Type            int16
	Page            int32 //期数
	StartUnixTime   int64
	EndUnixTime     int64
	DisposeUnixTime int64
	IsRelative      bool
	LTitle          string
	RTitle          string
	Content         string
	Weight          int16
	Tag             int8
	AwardList       []*JsonEventAwardItemInJsonFile
}

func (jsonEvent *JsonEvent) CheckStatus(kind int) bool {
	if jsonEvent == nil {
		return false
	}
	var dt int64
	if jsonEvent.IsRelative {
		dt = ServerCfg.ServerOpenTime
	}
	now := time.GetNowTime()
	if jsonEvent.StartUnixTime+dt > 0 && jsonEvent.StartUnixTime+dt > now {
		return false //活动还没开始的状况
	}
	if kind == NOT_DISPOSE { //领取奖励时验证是否过期
		if jsonEvent.DisposeUnixTime+dt > 0 && now >= jsonEvent.DisposeUnixTime+dt {
			return false
		}
	} else if kind == NOT_END { //获取活动时验证是否结束
		if jsonEvent.EndUnixTime+dt > 0 && now >= jsonEvent.EndUnixTime+dt {
			return false
		}
	}
	return true
}

func (jsonEvent *JsonEvent) GetMaxGrade() *JsonEventAward {
	if jsonEvent != nil {
		list := jsonEvent.List
		if len(list) > 0 {
			return list[len(list)-1]
		}
	}
	return nil
}

/*
 * 查找下个领奖点得索引和奖励内容
 * 注意判断是否找到用next==nil 来判断，index=0也有可能是没找到 也有可能是第一个
 */
func (jsonEvent *JsonEvent) GetNextGrade(now int32) (index int, next *JsonEventAward) {
	if jsonEvent == nil {
		return
	}
	max := jsonEvent.GetMaxGrade()
	if now < max.Grade {
		for tempIndex, item := range jsonEvent.List {
			if item.Grade > now {
				index = tempIndex
				next = item
				break
			}
		}
	}
	return
}

func (jsonEvent *JsonEvent) GetGradeByIndex(now int32) (index int32, next *JsonEventAward) {
	if jsonEvent == nil {
		return
	}
	max := jsonEvent.GetMaxGrade()
	if now <= max.Grade {
		for rank, v := range jsonEvent.List {
			if now > int32(v.Grade) {
				continue
			} else {
				index = int32(rank)
				next = v
				break
			}
		}
	}

	return
}

/*
 * 根据活动进度权值获取最大可领奖的经度
 * 同样不能通过index==0来判断是否能找到,通过maxAward==nil来判断
 */
func (jsonEvent *JsonEvent) GetMaxCanAwardGrade(now int32) (index int, maxAward *JsonEventAward) {
	if jsonEvent == nil {
		return
	}
	for tempIndex, item := range jsonEvent.List {
		if item.Grade <= now {
			index = tempIndex
			maxAward = item
		} else {
			break
		}
	}
	return
}

var JsonEvents map[int16]map[int32]*JsonEvent
var ListJsonEvents []*JsonEvent

func LoadJsonEventInfo(list []*JsonEventStructInJsonFile) {
	if len(list) > 0 {
		JsonEvents = make(map[int16]map[int32]*JsonEvent)

		for _, eventItem := range list {
			if eventItem.Type > 0 && eventItem.Page > 0 {
				typeEvents := JsonEvents[eventItem.Type]
				if typeEvents == nil {
					typeEvents = make(map[int32]*JsonEvent)
				}
				typeEvents[eventItem.Page] = &JsonEvent{
					Type:            eventItem.Type,
					Page:            eventItem.Page,
					StartUnixTime:   eventItem.StartUnixTime,
					EndUnixTime:     eventItem.EndUnixTime,
					DisposeUnixTime: eventItem.DisposeUnixTime,
					IsRelative:      eventItem.IsRelative,
					LTitle:          eventItem.LTitle,
					RTitle:          eventItem.RTitle,
					Content:         eventItem.Content,
					Weight:          eventItem.Weight,
					Tag:             eventItem.Tag,
				}
				ListJsonEvents = append(ListJsonEvents, typeEvents[eventItem.Page])
				for _, awardItem := range eventItem.AwardList {
					typeEvents[eventItem.Page].List = append(typeEvents[eventItem.Page].List, &JsonEventAward{
						Grade: awardItem.Require,
						Award: &EventDefaultAward{
							Ingot:    awardItem.Ingot,
							Coin:     awardItem.Coin,
							Heart:    awardItem.Heart,
							Item1Id:  awardItem.Item1Id,
							Item1Num: awardItem.Item1Num,
							Item2Id:  awardItem.Item2Id,
							Item2Num: awardItem.Item2Num,
							Item3Id:  awardItem.Item3Id,
							Item3Num: awardItem.Item3Num,
							Item4Id:  awardItem.Item4Id,
							Item4Num: awardItem.Item4Num,
							Item5Id:  awardItem.Item5Id,
							Item5Num: awardItem.Item5Num,
						},
					})
				}
				JsonEvents[eventItem.Type] = typeEvents
			}
		}
	}
}

func GetJsonEventInfoById(kind int16, page int32) (info *JsonEvent, exists bool) {

	if jsonEventInfoByKind, ok1 := JsonEvents[kind]; ok1 {
		if jsonEventInfo, ok2 := jsonEventInfoByKind[page]; ok2 {
			info = jsonEventInfo
			exists = true
		}
	}
	return
}

func GetJsonEventsInfo() []*JsonEvent {
	return ListJsonEvents
}

func EventIntervalDays(eventId int16) (index int) {
	if event, ok := mapEventCenter[eventId]; ok {

		var dt = event.Start
		if event.IsRelative == 1 { //是相对时间
			dt += ServerCfg.ServerOpenTime
		}

		now := time.GetNowTime()
		index = time.GetNowDayFromUnix(now) - time.GetNowDayFromUnix(dt)
		var timespan = time.GetNowDayFromUnix(event.End) - time.GetNowDayFromUnix(event.Start)

		if index > timespan {
			index = INVALIDINDEX
		} else {
			index += 1
		}

	} else { //没有找到
		index = INVALIDINDEX
	}

	return
}

type EventAnnounce struct {
	Ltitle        string
	Rtitle        string
	Content       string
	StartUnixTime int64
	EndUnixTime   int64
	Weight        int16
	Tag           int8
	Jump          int8
	IsRelative    int8
}

var TextEvents []*EventAnnounce
