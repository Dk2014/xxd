package multi_level

import (
	"core/fail"
	"core/log"
	"core/net"
	"core/time"
	"game_server/api/protocol/multi_level_api"
	"game_server/config"
	"game_server/dat/item_dat"
	"game_server/dat/multi_level_dat"
	"game_server/dat/player_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/module/rpc"
	"game_server/tlog"
	"sync"
	goTime "time"
)

type Partner struct {
	Session *net.Session

	Nick      []byte
	Level     int16
	FashionId int16

	MainRoleId  int8
	BuddyRoleId int8

	Col          int8 // 玩家所在列
	MainRoleRow  int8 // 主角所在行
	BuddyRoleRow int8 // 伙伴所在行
	Tactical     int8
	AssistOnly   bool // 仅仅是援助（当玩家当日关卡次数用完，只能加入房间或者被邀请，此时为助阵，胜利后不会获得奖励和经验)
}

type MultiLevelRoom struct {
	Id        int64 // 房间ID
	LevelId   int32 // 关卡ID
	LevelInfo *multi_level_dat.MultiLevel

	channel       *net.Channel
	offlinePlayer map[int64]*Partner // 战斗中离线的玩家

	LeaderPid int64              // 队长
	Partners  map[int64]*Partner // 合伙人

	OnFighting bool // 是否已开启战斗
}

const (
	MAX_ROOM_NUM = 3 // 房间人数上限
)

type dataTable struct {
	roomMutex      sync.Mutex
	autoTableMutex sync.Mutex

	idleRoomCount int64
	curRoomId     int64

	// 房间列表
	roomList map[int64]*MultiLevelRoom
	// 参与随机关卡的玩家队列
	autoTable map[int32]map[int64]*net.Session
}

const MaxRoomID = 9999

var multiDataTable *dataTable = &dataTable{
	idleRoomCount: MaxRoomID,
	curRoomId:     0,

	roomList:  make(map[int64]*MultiLevelRoom, 1000),
	autoTable: make(map[int32]map[int64]*net.Session),
}

func (this *dataTable) newRoomId() int64 {
	this.roomMutex.Lock()
	defer this.roomMutex.Unlock()

	// 没有空闲房间
	if this.idleRoomCount < 1 {
		return -1
	}

	// 房间号分配到上限就重置
	//if this.curRoomId == MaxRoomID {
	//	this.curRoomId = 0
	//}

	for this.curRoomId = (this.curRoomId + 1) % MaxRoomID; this.isRoomIDUsed(this.curRoomId); this.curRoomId = (this.curRoomId + 1) % MaxRoomID {
	}

	//this.curRoomId += 1
	this.idleRoomCount -= 1

	return this.curRoomId
}

func (this *dataTable) setRoom(id int64, room *MultiLevelRoom) {
	this.roomMutex.Lock()
	defer this.roomMutex.Unlock()
	this.roomList[id] = room
}

func (this *dataTable) getRoom(id int64) (*MultiLevelRoom, bool) {
	this.roomMutex.Lock()
	defer this.roomMutex.Unlock()
	room, ok := this.roomList[id]
	return room, ok
}

func (this *dataTable) isRoomIDUsed(id int64) bool {
	_, ok := this.roomList[id]
	return ok
}

func (this *dataTable) delRoom(id int64) {
	this.roomMutex.Lock()
	defer this.roomMutex.Unlock()

	delete(this.roomList, id)
	// 统计空闲房间
	this.idleRoomCount += 1
}

func (this *dataTable) setAutoLevel(levelId int32) {
	this.autoTableMutex.Lock()
	defer this.autoTableMutex.Unlock()
	if _, ok := this.autoTable[levelId]; !ok {
		this.autoTable[levelId] = make(map[int64]*net.Session)
	}
}

func (this *dataTable) getAutoPlayerLevelCount(levelId int32) int64 {
	this.autoTableMutex.Lock()
	defer this.autoTableMutex.Unlock()
	return int64(len(this.autoTable[levelId]))
}

func (this *dataTable) setAutoPlayerLevel(levelId int32, playerId int64, session *net.Session) {
	this.autoTableMutex.Lock()
	defer this.autoTableMutex.Unlock()
	this.autoTable[levelId][playerId] = session
}

func (this *dataTable) delAutoPlayerLevel(levelId int32, playerId int64) {
	this.autoTableMutex.Lock()
	defer this.autoTableMutex.Unlock()
	delete(this.autoTable[levelId], playerId)
}

func (this *dataTable) fetchAutoPlayerLevel(levelId int32, callback func(playerId int64, session *net.Session) bool) {
	this.autoTableMutex.Lock()
	defer this.autoTableMutex.Unlock()

	for pid, s := range this.autoTable[levelId] {
		if false == callback(pid, s) {
			break
		}
	}
}

/*
	备注：
	玩家在进出房间时，不计算上阵排列位置，当开始战斗的时候才进行计算
*/

func createRoom(session *net.Session, levelId int32) multi_level_api.RoomStatus {
	state := module.State(session)

	tlog.PlayerSystemModuleFlowLog(state.Database, tlog.SMT_MULITI_MISSION)
	if state.MultiLevelRoomId > 0 {
		return multi_level_api.ROOM_STATUS_FAILED_FIGHTING
	}

	levelInfo := multi_level_dat.GetMultiLevelById(int16(levelId))
	playerMultiLevel := state.Database.Lookup.PlayerMultiLevelInfo(state.PlayerId)
	if !time.IsInPointHour(player_dat.RESET_MULTI_LEVEL_TIMES_IN_HOUR, playerMultiLevel.BattleTime) {
		playerMultiLevel.DailyNum = 0
		state.Database.Update.PlayerMultiLevelInfo(playerMultiLevel)
	}
	enterChecker(state, playerMultiLevel, levelInfo.Lock, true)

	role := module.Role.GetMainRole(state.Database)
	fail.When(levelInfo.RequireLevel > role.Level, "require main role level")

	fashionState := state.Database.Lookup.PlayerFashionState(state.PlayerId)

	// 申请一个房间
	roomId := multiDataTable.newRoomId()
	if roomId <= 0 {
		// 房间已满
		return multi_level_api.ROOM_STATUS_FAILED_FULL
	}

	room := &MultiLevelRoom{
		Id:            roomId,
		LevelId:       levelId,
		LevelInfo:     levelInfo,
		channel:       net.NewChannel(),
		LeaderPid:     state.PlayerId,
		Partners:      make(map[int64]*Partner, MAX_ROOM_NUM),
		OnFighting:    false,
		offlinePlayer: make(map[int64]*Partner),
	}

	mainRoleLevel := module.Role.GetMainRole(state.Database).Level
	room.Partners[state.PlayerId] = newPartner(session, mainRoleLevel, playerMultiLevel, fashionState.DressedFashionId)
	room.channel.Join(session)

	state.MultiLevelRoomId = room.Id
	state.MultiLevelState = &module.MultiLevelState{
		BattlePetInfo:  module.BattlePet.GetAvailableBattlePet(state.Database),
		BattleItemInfo: module.Item.CountItemNumByType(state.Database, item_dat.TYPE_BATTLE_PROPS),
	}

	multiDataTable.setRoom(room.Id, room)

	notifyRoomInfo(session, room)

	return multi_level_api.ROOM_STATUS_SUCCESS
}

func enterRoom(session *net.Session, roomId int64) {
	state := module.State(session)
	fail.When(state.MultiLevelRoomId > 0, "[enterRoom] already enter room")
	fail.When(state.CrossInfo == nil, "[enterRoom] cross state is nil")

	rsp := &multi_level_api.EnterRoom_Out{}
	room, ok := multiDataTable.getRoom(roomId)
	if !ok {
		rsp.Status = multi_level_api.ROOM_STATUS_FAILED_NOT_EXIST
		session.Send(rsp)
		return
	}

	if room.OnFighting {
		rsp.Status = multi_level_api.ROOM_STATUS_FAILED_FIGHTING
		session.Send(rsp)
		return
	}

	if room.LevelInfo.RequireLevel > state.CrossInfo.RoleLevel {
		rsp.Status = multi_level_api.ROOM_STATUS_FAILED_REQUIRE_LEVEL
		session.Send(rsp)
		return
	}

	rpc.RemoteGetPlayerMultiLevelInfo(state.PlayerId, func(playerMultiLevel *mdb.PlayerMultiLevelInfo, petInfo map[int8]int32, itemInfo map[int16]int32, fashionId int16) {
		if room.LevelInfo.Lock > playerMultiLevel.Lock {
			rsp.Status = multi_level_api.ROOM_STATUS_FAILED_REQUIRE_LOCK
			session.Send(rsp)
			return
		}

		enterChecker(state, playerMultiLevel, room.LevelInfo.Lock, false)

		_, ok = room.Partners[state.PlayerId]
		fail.When(ok, "[enterRoom] already in room")
		if len(room.Partners) >= MAX_ROOM_NUM {
			rsp.Status = multi_level_api.ROOM_STATUS_FAILED_FULL
			session.Send(rsp)
			return
		}

		partner := newPartner(session, state.CrossInfo.RoleLevel, playerMultiLevel, fashionId)
		// 关卡次数超过的玩家可以加入房间，以助阵行为进行多人关卡
		partner.AssistOnly = (playerMultiLevel.DailyNum >= multi_level_dat.MULTI_LEVEL_DAILY_NUM_MAX)

		room.Partners[state.PlayerId] = partner

		// 通知房间队友有新进来的玩家
		module.API.Broadcast(room.channel, &multi_level_api.NotifyJoinPartner_Out{
			Partner: multi_level_api.PartnerInfo{
				Pid:         state.PlayerId,
				Nick:        partner.Nick,
				RoleId:      partner.MainRoleId,
				Level:       partner.Level,
				BuddyRoleId: partner.BuddyRoleId,
				FashionId:   partner.FashionId,
			},
		})

		room.channel.Join(session)
		state.MultiLevelRoomId = room.Id
		state.MultiLevelState = &module.MultiLevelState{
			BattlePetInfo:  petInfo,
			BattleItemInfo: itemInfo,
		}

		rsp.Status = multi_level_api.ROOM_STATUS_SUCCESS
		session.Send(rsp)
		notifyRoomInfo(session, room)

	})
}

func notifyRoomInfo(session *net.Session, room *MultiLevelRoom) {
	// 通知登陆房间的玩家，当前房间的信息
	rsp := &multi_level_api.NotifyRoomInfo_Out{}
	rsp.RoomId = room.Id
	rsp.LeaderPid = room.LeaderPid
	rsp.LevelId = room.LevelId
	rsp.Partners = []multi_level_api.NotifyRoomInfo_Out_Partners{}

	for pid, partner := range room.Partners {
		rsp.Partners = append(rsp.Partners, multi_level_api.NotifyRoomInfo_Out_Partners{
			Partner: multi_level_api.PartnerInfo{
				Pid:         pid,
				Nick:        partner.Nick,
				RoleId:      partner.MainRoleId,
				Level:       partner.Level,
				FashionId:   partner.FashionId,
				BuddyRoleId: partner.BuddyRoleId,
			},
		})
	}
	session.Send(rsp)
}

func leaveRoom(session *net.Session) {
	state := module.State(session)
	room, ok := multiDataTable.getRoom(state.MultiLevelRoomId)
	if !ok {
		log.Debugf("[leaveRoom] incorrect pid %d, gsid %d multi-level room id %d, crossInfo %#v",
			state.PlayerId, config.ServerCfg.ServerId, state.MultiLevelRoomId, state.CrossInfo)
		state.MultiLevelRoomId = 0
		state.MultiLevelState = nil
		return
	}

	if room.OnFighting && state.Battle != nil {
		state.Battle.LeaveBattle(session)
		room.offlinePlayer[state.PlayerId] = room.Partners[state.PlayerId]
	}

	state.MultiLevelRoomId = 0
	state.MultiLevelState = nil

	room.channel.Exit(session)
	delete(room.Partners, state.PlayerId)

	// 队长离开
	if room.LeaderPid == state.PlayerId {
		// 清理空房间
		if len(room.Partners) == 0 {
			multiDataTable.delRoom(room.Id)
			return
		}

		// 重选队长
		for pid, _ := range room.Partners {
			room.LeaderPid = pid
			break
		}
	}

	// 通知房间内其他队友
	module.API.Broadcast(room.channel, &multi_level_api.LeaveRoom_Out{
		Pid:       state.PlayerId,
		LeaderPid: room.LeaderPid,
	})
}

func newPartner(session *net.Session, mainRoleLevel int16, playerMultiLevel *mdb.PlayerMultiLevelInfo, fashionId int16) *Partner {
	state := module.State(session)

	var mainRoleRow int8 = 1
	if playerMultiLevel.BuddyRow == 1 {
		mainRoleRow = 2
	}

	return &Partner{
		Session: session,

		Nick:      state.PlayerNick,
		Level:     mainRoleLevel,
		FashionId: fashionId,

		MainRoleId:  state.RoleId,
		BuddyRoleId: playerMultiLevel.BuddyRoleId,

		Col:          0,
		MainRoleRow:  mainRoleRow,
		BuddyRoleRow: playerMultiLevel.BuddyRow,

		Tactical: playerMultiLevel.TacticalGrid,
	}
}

func autoEnterRoom(session *net.Session, levelId int32) {
	state := module.State(session)
	fail.When(state.MultiLevelRoomId > 0, "[autoEnterRoom] already enter room")

	levelInfo := multi_level_dat.GetMultiLevelById(int16(levelId))
	playerMultiLevel := state.Database.Lookup.PlayerMultiLevelInfo(state.PlayerId)
	if !time.IsInPointHour(player_dat.RESET_MULTI_LEVEL_TIMES_IN_HOUR, playerMultiLevel.BattleTime) {
		playerMultiLevel.DailyNum = 0
		state.Database.Update.PlayerMultiLevelInfo(playerMultiLevel)
	}

	enterChecker(state, playerMultiLevel, levelInfo.Lock, true)

	multiDataTable.setAutoLevel(levelId)

	// 人数不足
	if multiDataTable.getAutoPlayerLevelCount(levelId) < MAX_ROOM_NUM-1 {
		multiDataTable.setAutoPlayerLevel(levelId, state.PlayerId, session)

		state.CleanMultiLevelAuto = func() {
			multiDataTable.delAutoPlayerLevel(levelId, state.PlayerId)
			state.MultiLevelState = nil
			state.MultiLevelRoomId = 0
			state.CleanMultiLevelAuto = nil
		}

		// 随机超时清理
		state.TimerMgr.Start(module.TIMER_MULTI_LEVEL_AUTO, multi_level_dat.AUTO_BATTLE_TIMEOUT_SECONDS*goTime.Second, func(state *module.SessionState) {
			state.CleanMultiLevelAuto()
		})
		session.Send(&multi_level_api.AutoEnterRoom_Out{Success: false})
		return
	}

	roomId := multiDataTable.newRoomId()
	if roomId <= 0 {
		session.Send(&multi_level_api.AutoEnterRoom_Out{Success: false})
		return
	}

	room := &MultiLevelRoom{
		Id:            roomId,
		LevelId:       levelId,
		LevelInfo:     levelInfo,
		channel:       net.NewChannel(),
		Partners:      make(map[int64]*Partner, MAX_ROOM_NUM),
		OnFighting:    false,
		offlinePlayer: make(map[int64]*Partner),
	}

	var maxLevel int16
	var leaderPid int64
	var partner *Partner

	notifyRsp := &multi_level_api.NotifyPlayersInfo_Out{
		Players: []multi_level_api.NotifyPlayersInfo_Out_Players{
			{
				PlayerId: state.PlayerId,
				Nickname: state.PlayerNick,
			},
		},
	}

	var toDel []int64
	// 随机找出两个等待玩家
	multiDataTable.fetchAutoPlayerLevel(levelId, func(pid int64, s *net.Session) bool {

		state := module.State(s)
		state.TimerMgr.Stop(module.TIMER_MULTI_LEVEL_AUTO)

		mainRoleLevel := module.Role.GetMainRole(state.Database).Level
		multiLevel := state.Database.Lookup.PlayerMultiLevelInfo(state.PlayerId)
		fashionState := state.Database.Lookup.PlayerFashionState(state.PlayerId)

		partner = newPartner(s, mainRoleLevel, multiLevel, fashionState.DressedFashionId)
		if partner.Level > maxLevel {
			maxLevel = partner.Level
			leaderPid = pid
		}

		room.Partners[pid] = partner
		room.channel.Join(s)
		state.MultiLevelRoomId = room.Id
		state.MultiLevelState = &module.MultiLevelState{
			BattlePetInfo:  module.BattlePet.GetAvailableBattlePet(state.Database),
			BattleItemInfo: module.Item.CountItemNumByType(state.Database, item_dat.TYPE_BATTLE_PROPS),
		}

		toDel = append(toDel, pid)

		notifyRsp.Players = append(notifyRsp.Players, multi_level_api.NotifyPlayersInfo_Out_Players{
			PlayerId: pid,
			Nickname: partner.Nick,
		})

		if len(room.Partners) >= MAX_ROOM_NUM-1 {
			return false
		}

		return true
	})

	// 从随机队列移除
	for _, pid := range toDel {
		multiDataTable.delAutoPlayerLevel(levelId, pid)
	}

	// 加入当前玩家
	mainRoleLevel := module.Role.GetMainRole(state.Database).Level
	fashionState := state.Database.Lookup.PlayerFashionState(state.PlayerId)

	room.channel.Join(session)
	partner = newPartner(session, mainRoleLevel, playerMultiLevel, fashionState.DressedFashionId)
	state.MultiLevelRoomId = room.Id
	state.MultiLevelState = &module.MultiLevelState{
		BattlePetInfo:  module.BattlePet.GetAvailableBattlePet(state.Database),
		BattleItemInfo: module.Item.CountItemNumByType(state.Database, item_dat.TYPE_BATTLE_PROPS),
	}
	room.Partners[state.PlayerId] = partner
	if partner.Level > maxLevel {
		leaderPid = state.PlayerId
	}

	// 等级最高的为队长
	room.LeaderPid = leaderPid
	multiDataTable.setRoom(room.Id, room)

	session.Send(&multi_level_api.AutoEnterRoom_Out{Success: true})
	module.API.Broadcast(room.channel, notifyRsp)

	newBattle(room)
}

func cancelAutoMatch(state *module.SessionState) {
	if state.CleanMultiLevelAuto != nil {
		state.TimerMgr.Stop(module.TIMER_MULTI_LEVEL_AUTO)
		state.CleanMultiLevelAuto()
	}
}
