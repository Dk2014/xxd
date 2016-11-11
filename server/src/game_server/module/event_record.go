package module

import (
	"bytes"
	"core/fail"
	"core/net"
	"core/time"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"game_server/dat/event_dat"
	"game_server/dat/player_dat"
	"game_server/mdb"
	gotime "time"
)

// NOTE: 如果EventInfo结构进行变更和调整，请找到isJoinGroupBuy接口，进行相应更新变更
type EventInfo struct {
	EventId     int16
	EndUnixTime int64
	Awarded     int32
	MaxAward    int32
	LastUpdated int64
}

type IPlayerEventOperate interface {
	GetEventStatus() (IsAppend bool, Process, PlayerProcess int32, isAward bool)
	GetEventAward(param1, param2, param3 int32)
}

var ExtraEvents map[int16]IPlayerEventOperate = make(map[int16]IPlayerEventOperate) //后续修改的活动添加到这个map中

var EventInfoStructSize = binary.Size(&EventInfo{})

type EventInfoList struct {
	List map[int16]*EventInfo
}

func NewEventInfoList() (ev *EventInfoList) {
	ev = &EventInfoList{}
	ev.List = make(map[int16]*EventInfo)
	return
}
func (list *EventInfoList) Decode(bin []byte) {
	buffer := net.NewBuffer(bin)
	count := len(bin) / EventInfoStructSize //一个活动的记录总长度为26
	for index := 0; index < count; index++ {
		eventId := int16(buffer.ReadUint16LE())
		list.List[eventId] = &EventInfo{
			EventId:     eventId,
			EndUnixTime: int64(buffer.ReadUint64LE()),
			Awarded:     int32(buffer.ReadUint32LE()),
			MaxAward:    int32(buffer.ReadUint32LE()),
			LastUpdated: int64(buffer.ReadUint64LE()),
		}
	}
}

func (list *EventInfoList) Encode() []byte {
	count := len(list.List)
	bytes := make([]byte, count*EventInfoStructSize)
	buffer := net.NewBuffer(bytes)
	for _, info := range list.List {
		buffer.WriteUint16LE(uint16(info.EventId))
		buffer.WriteUint64LE(uint64(info.EndUnixTime))
		buffer.WriteUint32LE(uint32(info.Awarded))
		buffer.WriteUint32LE(uint32(info.MaxAward))
		buffer.WriteUint64LE(uint64(info.LastUpdated))
	}
	return buffer.Get()
}

func (list *EventInfoList) GetPlayerEventInfoById(eventId int16) *EventInfo {
	eventInfo, ok1 := event_dat.GetEventInfoById(eventId)
	fail.When(!ok1, "wrong event ID")
	info, ok2 := list.List[eventId]
	if eventInfo.Id == event_dat.EVENT_FIRST_RECHARGE_DAILY { // 单独正对本次活动上线问题，有待明天修复
		if ok2 && eventInfo.End == info.EndUnixTime && event_dat.CheckEventTime(eventInfo, event_dat.NOT_DISPOSE) {
			return info
		}
	} else {
		if ok2 {
			return info
		}
	}
	return &EventInfo{
		EventId:     eventId,
		EndUnixTime: eventInfo.End,
		Awarded:     0,
		MaxAward:    0,
		LastUpdated: 0,
	}
}

//更新已领取的奖励，不存在的话就忽略
func (list *EventInfoList) UpdateAwarded(db *mdb.Database, eventId int16, awarded int32) {
	if info, ok := list.List[eventId]; ok && info.Awarded < awarded {
		info.Awarded = awarded
		info.LastUpdated = time.GetNowTime()

		//更新至数据库
		oldRecord := db.Lookup.PlayerEventAwardRecord(db.PlayerId())
		oldRecord.RecordBytes = list.Encode()
		db.Update.PlayerEventAwardRecord(oldRecord)
	} else {
		fail.When(true, "can't not update awarded event record")
	}
}

//更新最大可领取奖励，如果没找到的话就插入
func (list *EventInfoList) UpdateMax(db *mdb.Database, eventId int16, maxAward int32) {
	if info, ok := list.List[eventId]; ok {
		info.MaxAward = maxAward
	} else {
		list.List[eventId] = &EventInfo{
			EventId:  eventId,
			Awarded:  0,
			MaxAward: maxAward,
			//EndUnixTime: info.EndUnixTime,
		}
	}
	list.List[eventId].LastUpdated = time.GetNowTime()
	event_info, _ := event_dat.GetEventInfoById(eventId)
	list.List[eventId].EndUnixTime = event_info.End
	//更新至数据库
	oldRecord := db.Lookup.PlayerEventAwardRecord(db.PlayerId())
	oldRecord.RecordBytes = list.Encode()
	db.Update.PlayerEventAwardRecord(oldRecord)
}

//单独更新最近领取奖励得时间点
func (list *EventInfoList) UpdateAwardedTime(db *mdb.Database, eventId int16) {
	if _, ok := event_dat.GetEventInfoById(eventId); ok {
		info := list.GetPlayerEventInfoById(eventId)
		info.LastUpdated = time.GetNowTime()
		list.List[eventId] = info

		//更新至数据库
		oldRecord := db.Lookup.PlayerEventAwardRecord(db.PlayerId())
		oldRecord.RecordBytes = list.Encode()
		db.Update.PlayerEventAwardRecord(oldRecord)
	}
}

//清空玩家相应事件的状态
func (list *EventInfoList) ClearState(db *mdb.Database, eventId int16) {
	info, ok := event_dat.GetEventInfoById(eventId) //活动存在与否
	fail.When(!ok, "wrong event ID")
	list.List[eventId] = &EventInfo{
		EventId:     eventId,
		EndUnixTime: info.End,
		Awarded:     0,
		MaxAward:    0,
		LastUpdated: time.GetNowTime(),
	}
	oldRecord := db.Lookup.PlayerEventAwardRecord(db.PlayerId())
	oldRecord.RecordBytes = list.Encode()
	db.Update.PlayerEventAwardRecord(oldRecord)
}

//初始化一个玩家的活动状态，并指定初识max
func (list *EventInfoList) AddEventAwardState(db *mdb.Database, eventId int16, firstMax, firstAwarded int32) *EventInfo {
	info, ok := event_dat.GetEventInfoById(eventId)
	fail.When(!ok, "wrong event ID")

	var exist bool
	var event *EventInfo
	nowTime := time.GetNowTime()

	if event, exist = list.List[eventId]; !exist {
		event = &EventInfo{}
		list.List[eventId] = event
	}

	event.EventId = eventId
	event.EndUnixTime = info.End
	event.Awarded = firstAwarded
	event.MaxAward = firstMax
	event.LastUpdated = nowTime

	oldRecord := db.Lookup.PlayerEventAwardRecord(db.PlayerId())
	oldRecord.RecordBytes = list.Encode()
	db.Update.PlayerEventAwardRecord(oldRecord)

	return event
}

func InitPlayerEventsStatus(session *net.Session) {
	state := State(session)
	record := state.Database.Lookup.PlayerEventAwardRecord(state.PlayerId)
	eventInfoList := NewEventInfoList()
	eventInfoList.Decode(record.RecordBytes)
	state.EventsState = eventInfoList

	//登陆的时候必然会调用活动列表接口，顾不再在这里调用活动检查，放到获取活动中心列表接口去检查
	//CheckForEvents(session)

}

// 比武场战斗结束调用时传实际名次的rank
func UpdateEventArenaRank(state *SessionState, rank int32) {
	eventId := int16(event_dat.EVENT_ARENA_RANK_AWARDS)
	eventInfo, _ := event_dat.GetEventInfoById(eventId)
	singleState := state.EventsState.GetPlayerEventInfoById(eventId)
	if singleState.MaxAward != rank {
		//名次有变动则更新
		if event_dat.CheckEventTime(eventInfo, event_dat.NOT_END) {
			mdb.GlobalExecute(func(globalDB *mdb.Database) {
				globalDB.AgentExecute(state.PlayerId, func(agentDB *mdb.Database) {
					state.EventsState.AddEventAwardState(agentDB, event_dat.EVENT_ARENA_RANK_AWARDS, rank, 0)
				})
			})
		}

	}
}

func CheckForEvents(session *net.Session) {
	state := State(session)
	events := event_dat.GetEventsInfo()
	eventState := state.EventsState
	for _, event := range events {
		needClear := true
		singleState := eventState.GetPlayerEventInfoById(event.Id)
		//开始处理活动检查
		switch event.Id {
		case event_dat.EVENT_ARENA_RANK_AWARDS:
			//比武场活动
			if Player.IsOpenFunc(state.Database, player_dat.FUNC_ARENA) {
				if event_dat.CheckEventTime(event, event_dat.NOT_DISPOSE) && !event_dat.IsInPointEventTime(event, event_dat.NOT_DISPOSE, singleState.LastUpdated) {
					Arena.GetPlayerArenaRank(state.Database, func(rank int32) {
						mdb.GlobalExecute(func(globalDB *mdb.Database) {
							globalDB.AgentExecute(state.PlayerId, func(agenDB *mdb.Database) {
								state.EventsState.AddEventAwardState(agenDB, event_dat.EVENT_ARENA_RANK_AWARDS, rank, 0)
							})
						})
					})
				}
			}
			needClear = false
		case event_dat.EVENT_MULTIPY_CONFIG:
			//ADD 翻倍奖励活动
		case event_dat.EVENT_LEVEL_AWARD:
			//ADD 等级活动处理
		case event_dat.EVENT_LOGIN_AWARD:
			//ADD 累计登录活动处理
		case event_dat.EVENT_STRONG_AWARD:
			//ADD 战力活动处理
		case event_dat.EVENT_RECHARGE_AWARD:
			//ADD 充值返利活动处理
		case event_dat.EVENT_PHYSICAL_AWARDS:
			//ADD 活跃度活动
		case event_dat.EVENT_TOTAL_LOGIN:
			//连续登录活动奖励
			if event_dat.CheckEventTime(event, event_dat.NOT_END) {
				totalLoginState := state.EventsState.GetPlayerEventInfoById(event_dat.EVENT_TOTAL_LOGIN)
				day_zero := time.GetTodayZero()
				if day_zero > totalLoginState.LastUpdated || totalLoginState.MaxAward == 0 {
					if totalLoginState.MaxAward == 0 || totalLoginState.LastUpdated > day_zero-24*3600 {
						//连续登录 天数加1
						state.EventsState.UpdateMax(state.Database, event.Id, totalLoginState.MaxAward+1)
					} else if playerInfo := state.Database.Lookup.PlayerInfo(state.PlayerId); playerInfo.LastOfflineTime > day_zero-24*3600 {
						//上次是跨天玩游戏
						gap_days := gotime.Unix(playerInfo.LastOfflineTime, 0).Day() - gotime.Now().Day()
						state.EventsState.UpdateMax(state.Database, event.Id, totalLoginState.MaxAward+int32(gap_days))
					} else {
						//隔天登录清空累计登录天数
						state.EventsState.List[event.Id] = &EventInfo{
							EventId:     event.Id,
							EndUnixTime: event.End,
							Awarded:     0,
							MaxAward:    1,
							LastUpdated: time.GetNowTime(),
						}
						state.Database.Update.PlayerEventAwardRecord(&mdb.PlayerEventAwardRecord{
							Pid:             state.Database.PlayerId(),
							RecordBytes:     state.EventsState.Encode(),
							JsonEventRecord: state.JsonEventsState.Encode(),
						})

					}
				}
			}
		case event_dat.EVENT_FIRST_RECHARGE_DAILY:

		}
		if needClear && (!event_dat.CheckEventTime(event, event_dat.NOT_DISPOSE) || event.End != singleState.EndUnixTime) {
			state.EventsState.ClearState(state.Database, event.Id)
		}
	}
}

//翻倍奖励配置获取
func GetMulitipyByKey(key int32) float32 {
	var result float32 = 1
	result = event_dat.GetMultipyByKey(key)
	return result
}

//####################后续的json配置的活动的相关代码#######################
type JsonEventRecordItem struct {
	Type        int16
	Page        int32
	Process     int32
	Awarded     int32
	LastUpdated int64
}

type JsonEventRecord map[string]*JsonEventRecordItem

func (record *JsonEventRecord) Encode() []byte {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(record)
	if err != nil {
		fail.When(true, err.Error())
	}
	return buffer.Bytes()
}

func (record *JsonEventRecord) Decode(buf []byte) {
	if len(buf) > 0 {
		buffer := bytes.NewBuffer(buf)
		decoder := gob.NewDecoder(buffer)
		err := decoder.Decode(record)
		if err != nil {
			fail.When(true, err.Error())
		}
	}
}

func (record *JsonEventRecord) Save(db *mdb.Database) {
	recordBytes := record.Encode()
	oldItem := db.Lookup.PlayerEventAwardRecord(db.PlayerId())
	oldItem.JsonEventRecord = recordBytes
	db.Update.PlayerEventAwardRecord(oldItem)
}

func (record *JsonEventRecord) Load(db *mdb.Database) {
	buffer := db.Lookup.PlayerEventAwardRecord(db.PlayerId()).JsonEventRecord
	record.Decode(buffer)
}

func (record *JsonEventRecord) GetStatus(kind int16, page int32) (status *JsonEventRecordItem, exists bool) {
	idenity := fmt.Sprintf("%d-%d", kind, page)
	status, exists = (*record)[idenity]
	return
}

//注意 调用ChangeStatus尽心一系列删除之后应立即调用Save方法来保存到数据库
func (record *JsonEventRecord) ChangeStatus(kind int16, page int32, process, awarded int32) {
	idenity := fmt.Sprintf("%d-%d", kind, page)
	status, exists := (*record)[idenity]
	if !exists {
		status = &JsonEventRecordItem{
			Type: kind,
			Page: page,
		}
		(*record)[idenity] = status // 新建的写入到json活动状态中
	}
	status.Process = process
	status.Awarded = awarded
	status.LastUpdated = time.GetNowTime()
}

//注意 调用DeleteStatus尽心一系列删除之后应立即调用Save方法来保存到数据库
func (record *JsonEventRecord) DeleteStatus(kind int16, page int32) (changed bool) {
	idenity := fmt.Sprintf("%d-%d", kind, page)
	_, changed = (*record)[idenity]
	if changed {
		delete(*record, idenity)
	}
	return
}

func InitPlayerJsonEventsStatus(session *net.Session) {
	state := State(session)
	state.JsonEventsState = &JsonEventRecord{}
	state.JsonEventsState.Load(state.Database)
}
