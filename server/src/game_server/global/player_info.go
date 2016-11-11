package global

import (
	"core/fail"
	"core/mysql"
	"core/net"
	"game_server/dat/event_dat"
	"runtime"
	"sync"
	"sync/atomic"
)

type PlayerInfo struct {
	Openid     []byte
	PlayerId   int64
	CId        int64
	PlayerNick []byte
	RoleId     int8

	FightNum  int32
	RoleLevel int16

	MultiLevelLock     int32
	MultiLevelDailyNum int8

	LastLoginTime   int64
	LastOfflineTime int64

	ArenaTrendWin           int16 // 竞技场趋势 -1下降；0不变；>0上升
	PushNotificationOptions int64 //推送通知开关
	VIPLevel                int16 //vip等级
	FashionId               int16 //时装

	WegamesPlatformUid []byte //台湾wegames 平台ID 详细说明见 module/player/player.go
}

const (
	PLAYER_INFO_TYPE_FIGHT_NUM = iota + 1
	PLAYER_INFO_TYPE_ROLE_LEVEL
	PLAYER_INFO_TYPE_LOGIN_TIME
	PLAYER_INFO_TYPE_OFFLINE_TIME

	PLAYER_INFO_TYPE_MULTI_LEVEL_LOCK
	PLAYER_INFO_TYPE_MULTI_DAILY_NUM

	PLAYER_INFO_TYPE_ARENA_TREND_WIN

	PLAYER_INFO_TYPE_PUSH_NOTIFICATION_OPTION

	PLAYER_INFO_TYPE_VIP_LEVEL

	PLAYER_INFO_TYPE_FASHION
)

type playerInfoTable struct {
	sync.RWMutex
	tables            map[int64]PlayerInfo
	nickMap           map[string]int64
	openIdMap         map[string]int64
	VIP               []int32 //该slice保存所有的vip总数和分阶段总数数据 0代表所有的总数 其他对应相应阶级的vip个数
	GroupBuyInfoTable map[int16]int32
}

func (this *playerInfoTable) ExistPlayerInfo(pid int64) bool {
	this.RLock()
	defer this.RUnlock()

	_, ok := this.tables[pid]
	return ok
}

func (this *playerInfoTable) GetPlayerInfo(pid int64) *PlayerInfo {
	this.RLock()
	defer this.RUnlock()

	p, ok := this.tables[pid]
	fail.When(!ok, "[GetPlayerInfo] can't get player info")
	return &p
}

/*
	避免update整个playerInfo产生脏数据，所以这里提供对字段原子操作的update方法
*/

func (this *playerInfoTable) GetPlayerIdWithNick(nick string) (pid int64, ok bool) {
	this.RLock()
	defer this.RUnlock()

	pid, ok = this.nickMap[nick]
	return
}

func (this *playerInfoTable) GetPlayerIdWithOpenId(openId string) (pid int64, ok bool) {
	this.RLock()
	defer this.RUnlock()

	pid, ok = this.openIdMap[openId]
	return pid, ok
}

func (this *playerInfoTable) UpdateLastLoginTime(playerId int64, value int64) {
	this.Lock()
	defer this.Unlock()

	info, ok := this.tables[playerId]
	fail.When(!ok, "[UpdateLastLoginTime] can't get player info")

	info.LastLoginTime = value
	this.tables[playerId] = info
}

func (this *playerInfoTable) addPlayerInfo(info *PlayerInfo) {
	this.Lock()
	defer this.Unlock()
	this.tables[info.PlayerId] = *info
	this.nickMap[string(info.PlayerNick)] = info.PlayerId
	this.openIdMap[string(info.Openid)] = info.PlayerId
	if len(info.WegamesPlatformUid) > 0 {
		this.openIdMap[string(info.WegamesPlatformUid)] = info.PlayerId
	}
}

func (this *playerInfoTable) UpdatePlayerInfoSets(sets map[int64]map[int]int64) {
	this.Lock()
	defer this.Unlock()

	for playerId, values := range sets {
		info, ok := this.tables[playerId]
		if !ok {
			continue
		}

		for valueType, value := range values {
			switch valueType {
			case PLAYER_INFO_TYPE_FIGHT_NUM:
				info.FightNum = int32(value)

			case PLAYER_INFO_TYPE_ROLE_LEVEL:
				info.RoleLevel = int16(value)

			case PLAYER_INFO_TYPE_LOGIN_TIME:
				info.LastLoginTime = value

			case PLAYER_INFO_TYPE_OFFLINE_TIME:
				info.LastOfflineTime = value

			case PLAYER_INFO_TYPE_MULTI_LEVEL_LOCK:
				info.MultiLevelLock = int32(value)

			case PLAYER_INFO_TYPE_MULTI_DAILY_NUM:
				info.MultiLevelDailyNum = int8(value)

			case PLAYER_INFO_TYPE_ARENA_TREND_WIN:
				info.ArenaTrendWin = int16(value)
			case PLAYER_INFO_TYPE_PUSH_NOTIFICATION_OPTION:
				info.PushNotificationOptions = value
			case PLAYER_INFO_TYPE_FASHION:
				info.FashionId = int16(value)
			case PLAYER_INFO_TYPE_VIP_LEVEL:
				if value > 0 {
					if info.VIPLevel == 0 {
						this.VIP[0] = atomic.AddInt32(&this.VIP[0], 1) //新加入vip的人加1
					} else {
						this.VIP[int(info.VIPLevel)] = atomic.AddInt32(&this.VIP[int(info.VIPLevel)], -1) //老阶段的vip人数减1
					}
					this.VIP[int(value)] = atomic.AddInt32(&this.VIP[int(value)], 1) //新进阶的vip人数加1
				}
				info.VIPLevel = int16(value)
			}

		}
		this.tables[playerId] = info
	}
}

func (this *playerInfoTable) UpdatePlayerInfo(pid int64, info *PlayerInfo) {
	this.tables[pid] = *info
}

func (this *playerInfoTable) load(mysqlInfo map[string]interface{}) {
	db, err1 := mysql.Connect(mysqlInfo)
	fail.When(err1 != nil, err1)

	defer db.Close()
	res, err := db.ExecuteFetch([]byte(`select p.id,p.nick,p.user,pi.last_offline_time,pr.role_id,pr.level,pf.fight_num,mul.lock,mul.daily_num, push.options,vip.present_exp,vip.level as vip_level, fashion.dressed_fashion_id as fashion_id, 
	 pwu.wegames_platform_uid as wegames_platform_uid
	 from player p
	 left join player_info pi on pi.pid = p.id
	 left join player_wegames_uid pwu on pwu.pid = p.id
	 left join player_fight_num pf on pf.pid = p.id
	 left join player_role pr on p.main_role_id = pr.id
	 left join player_multi_level_info mul on mul.pid = p.id 
	 left join player_push_notify_switch push on push.pid = p.id
	 left join player_fashion_state fashion on fashion.pid = p.id
	left join player_vip vip on vip.pid = p.id `), -1)
	fail.When(err != nil, err)

	iId := res.Map("id")
	iUser := res.Map("user")
	iNick := res.Map("nick")
	iWegamesUid := res.Map("nick")
	iRoleId := res.Map("role_id")
	iLevel := res.Map("level")
	iFightNum := res.Map("fight_num")
	iLock := res.Map("lock")
	iDailyNum := res.Map("daily_num")
	iOfflineTime := res.Map("last_offline_time")
	iPushNotifyOptions := res.Map("options")
	iVIPLevel := res.Map("vip_level")
	iVIPPresentExp := res.Map("present_exp")
	iFashionId := res.Map("fashion_id")

	var pid int64
	var nick []byte
	var wegameUid []byte
	for _, row := range res.Rows {
		pid = row.Int64(iId)
		nick = row.Bin(iNick)
		wegameUid = row.Bin(iWegamesUid)

		this.nickMap[string(nick)] = pid
		this.openIdMap[string(row.Bin(iUser))] = pid
		if len(wegameUid) > 0 {
			this.openIdMap[string(wegameUid)] = pid
		}
		this.tables[pid] = PlayerInfo{
			Openid:                  row.Bin(iUser),
			PlayerId:                pid,
			PlayerNick:              nick,
			FightNum:                row.Int32(iFightNum),
			RoleId:                  row.Int8(iRoleId),
			RoleLevel:               row.Int16(iLevel),
			MultiLevelLock:          row.Int32(iLock),
			MultiLevelDailyNum:      row.Int8(iDailyNum),
			LastOfflineTime:         row.Int64(iOfflineTime),
			PushNotificationOptions: row.Int64(iPushNotifyOptions),
			VIPLevel:                row.Int16(iVIPLevel),
			FashionId:               row.Int16(iFashionId),
			WegamesPlatformUid:      wegameUid,
		}
		vip_level := this.tables[pid].VIPLevel
		if vip_level > 0 {
			if row.Int64(iVIPPresentExp) == 0 {
				this.VIP[0] = atomic.AddInt32(&this.VIP[0], 1)                           //vip总数加1
				this.VIP[int(vip_level)] = atomic.AddInt32(&this.VIP[int(vip_level)], 1) //相应vip阶段的加1
			}
		}
	}

	this.loadGroupBuyInfo(db)
}

func (this *playerInfoTable) loadGroupBuyInfo(db *mysql.Connection) {
	eventId := int16(event_dat.EVENT_GROUP_BUY) //目前团购活动就只有一个 在活动中心
	eventInfo, _ := event_dat.GetEventInfoById(eventId)
	if !event_dat.CheckEventTime(eventInfo, event_dat.NOT_DISPOSE) {
		return
	}

	res, err := db.ExecuteFetch([]byte("select `record_bytes` from `player_event_award_record`"), -1)
	fail.When(err != nil, err)

	iRecordBytes := res.Map("record_bytes")

	wg := new(sync.WaitGroup)
	wgn := runtime.GOMAXPROCS(-1)
	pages := len(res.Rows) / wgn

	var start, end int
	for i := 0; i < wgn; i++ {
		wg.Add(1)
		start = i * pages
		if i == wgn-1 {
			end = len(res.Rows)
		} else {
			end = (i + 1) * pages
		}
		go func(start, end int) {
			for ; start < end; start++ {
				raw := res.Rows[start].Bin(iRecordBytes)
				if len(raw) > 0 && isJoinGroupBuy(raw, eventInfo) {
					this.Lock()
					this.GroupBuyInfoTable[eventId]++
					this.Unlock()
				}
			}
			wg.Done()
		}(start, end)
	}
	wg.Wait()
}

func isJoinGroupBuy(recordBytes []byte, eventInfo *event_dat.EventCenter) (isJoin bool) {
	buffer := net.NewBuffer(recordBytes)
	eventCount := len(recordBytes) / 26 //一个活动记录的长度为26
	for i := 0; i < eventCount; i++ {
		eventId := buffer.ReadUint16LE()
		endUnixTime := int64(buffer.ReadUint64LE())
		if eventId == event_dat.EVENT_GROUP_BUY && endUnixTime == eventInfo.End {

			buffer.SetReadPosition(buffer.GetReadPosition() + 4)
			maxAward := buffer.ReadUint32LE()
			if maxAward > 0 {
				isJoin = true
			}
			buffer.SetReadPosition(buffer.GetReadPosition() + 8)
			break
		} else {
			buffer.SetReadPosition(buffer.GetReadPosition() + 16)
		}
	}
	return
}
