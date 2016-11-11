package module

import (
	"core/debug"
	"core/log"
	"core/net"
	"game_server/api/protocol"
	"game_server/config"
	"game_server/mdb"
	"sync/atomic"
)

var (
	traceUpModule   int8
	traceUpAPI      int8
	traceDownModule int8
	traceDownAPI    int8
)

func SetTraceUp(mid, aid int8) {
	traceUpModule, traceUpAPI = mid, aid
}

func SetTraceDown(mid, aid int8) {
	traceDownModule, traceDownAPI = mid, aid
}

type CrossPlayerInfo struct {
	RoleLevel int16
}

type MultiLevelState struct {
	BattlePetInfo  map[int8]int32  //玩家灵宠信息 格子ID -> 灵宠模版ID
	BattleItemInfo map[int16]int32 //玩家战斗道具信息 物品ID -> 数量
}

type PetVirtualEnvState struct {
	Floor          int16   //关卡层数
	PveLevelId     int16   //灵宠幻境关卡数据ID
	EnemyNum       int16   //杀怪数量
	MaxLoadedIndex int     //已加载的怪物组最大索引
	EnemyIds       []int32 //
}
type DrivingSwordState struct {
	//Floor          int16   //关卡层数
	DrivingLevelId int16 //关卡数据ID
	//EnemyNum       int16   //杀怪数量
	MaxLoadedIndex int     //已加载的怪物组最大索引
	EnemyIds       []int32 //
}
type DrivingSwordMapCache struct {
	UnknowTeleports map[uint16] /* (x << 8) | y */ *mdb.PlayerDrivingSwordEvent //未展示的传送点
}
type TencentState struct {
	QqVipStatus int16
}

type SessionState struct {
	occurPanic bool // 发生panic异常

	SendReLoginCallback func() // 发送重新登陆指令（当会话状态关闭和丢弃时）

	Database        *mdb.Database
	sandBoxDatabase *mdb.Database //沙箱数据，用来获取一个不起作用的事务控制器
	MoneyState      *MoneyState

	TimerMgr *Timer

	PlayerId           int64
	CId                int32
	PlayerNick         []byte
	RoleId             int8  // 玩家主角ID
	TownId             int16 // 玩家当前所在城镇
	WegamesPlatformUid []byte

	LastMoveInTown int64
	LastTownX      int16
	LastTownY      int16

	TownChannel *net.Channel

	//帮派
	InCliqueClubhouse bool  //互动服用 是否在帮派集会所标记
	IsMeditation      bool  //是否打坐
	ClubhouseX        int16 //帮派内Y坐标
	ClubhouseY        int16 //帮派内X坐标

	// 战场信息
	Battle IBattle
	// 玩家区域关卡状态
	MissionLevelState *MissionLevelState
	// 比武场状态
	ArenaState *PlayerArenaState
	//彩虹关卡状态
	RainbowLevelState *RainbowLevelState

	// 多人关卡房间ID
	MultiLevelRoomId    int64
	CleanMultiLevelAuto func()           // 清理随机多人关卡的接口
	MultiLevelState     *MultiLevelState //多人关卡状态
	//游戏活动记录状态
	EventsState *EventInfoList
	//游戏json配置活动记录状态
	JsonEventsState *JsonEventRecord

	CrossInfo *CrossPlayerInfo
	RspQ      *ResponseQueue

	// 客户端登出tlog用
	TlogClientVersion  string
	TlogNetwork        string
	TlogTelecomOper    string
	TlogSystemHardware string

	RefreshDailyQuestAtWeeHours bool // 凌晨更新每日任务

	TencentState       *TencentState
	PetVirtualEnvState *PetVirtualEnvState //灵宠幻境状态
	DrivingSwordState  *DrivingSwordState  //仙山御剑状态

	DrivingSwordMapCache *DrivingSwordMapCache //云海驭剑地图相关缓存
	sealedBook           *SealedBookRecord     //天书记录
	Ip                   string                //登陆时ip
}

func NewSessionState() *SessionState {
	state := &SessionState{}
	state.TimerMgr = NewTimer(state)
	state.Database = mdb.NewDatabase()
	state.RspQ = NewResponseQueue()
	return state
}

func (s *SessionState) MarkPanic() {
	s.occurPanic = true
}

func checkIsCountAPI(mid, aid int8) bool {
	if mid == 88 /* server_info */ {
		return false
	} else if mid == 0 && (aid == 4 || aid == 7 || aid == 8) /* player::get_time || player::login */ {
		return false
	}

	return true
}

func (s *SessionState) RecordResponse(rsp interface{}) {
	var mid, aid int8

	switch rsp.(type) {
	case net.Response:
		mid, aid = rsp.(net.Response).GetModuleIdAndActionId()
	default:
		buf := rsp.([]byte)
		mid, aid = int8(buf[config.API_PACKENT_HEAD_LENGTH]), int8(buf[config.API_PACKENT_HEAD_LENGTH+1])
	}
	if mid == traceDownModule && (aid == traceDownAPI || traceDownAPI == -1) {
		log.Debugf("[Trace Down] pid %d, protocol %d %d\n %s \n %s\n", s.PlayerId, mid, aid,
			debug.Print(0, false, true, "    ", nil, rsp),
			debug.Stack(2, "    "))
	}

	// 只在gs服上记录下行数据。跨服和互动服不处理
	if config.ServerCfg.EnableGlobalServer || s.CrossInfo != nil || !checkIsCountAPI(mid, aid) {
		return
	}

	s.RspQ.addRsp(rsp)
}

func (s *SessionState) CountAPIRequest(request protocol.Request) {
	mid, aid := request.GetModuleIdAndActionId()
	if mid == traceUpModule && (aid == traceUpAPI || traceUpAPI == -1) {
		log.Debugf("[Trace Up] pid %d, protocol %s \n %s\n", s.PlayerId, request.TypeName(),
			debug.Print(0, false, true, "    ", nil, request))
	}
	// 只统计gs上收到的上行数据。跨服和互动服不处理
	if config.ServerCfg.EnableGlobalServer || s.CrossInfo != nil || !checkIsCountAPI(mid, aid) {
		return
	}

	atomic.AddInt64(&s.RspQ.ReqCounter, 1)
}

func (s *SessionState) SandBoxMode() {
	if s.sandBoxDatabase == nil {
		s.sandBoxDatabase = mdb.NewDatabase()
		log.Debugf("enter snad box mode %p\n", s.sandBoxDatabase)
		s.sandBoxDatabase.Mount(s.PlayerId, s.CId)
	}
}

func (s *SessionState) ExitSandbox() {
	if s.sandBoxDatabase != nil {
		s.sandBoxDatabase = nil
	}
}

func (s *SessionState) cleanupSandbox() {
	if s.sandBoxDatabase != nil {
		s.Database.Rollback()
		s.sandBoxDatabase = nil
		s.MissionLevelState = nil
		s.Battle = nil

		playerTown := s.Database.Lookup.PlayerTown(s.PlayerId)
		if playerTown != nil {
			s.TownId = playerTown.TownId
			s.LastTownX = playerTown.AtPosX
			s.LastTownY = playerTown.AtPosY
		}
	}
}

// 会话关闭时将SessionState保存起来
func DeleteSessionState(session *net.Session) {
	state := State(session)

	defer func() {
		if err := recover(); err != nil {
			log.Errorf(`DeleteSessionState
Error   = %v
Session = {PlayerId:%d, Nickname:"%s", TownId:%d, RoleId:%d}
Stack   =
%s`,
				err,
				state.PlayerId,
				state.PlayerNick,
				state.TownId,
				state.RoleId,
				debug.Stack(1, "    "),
			)
		}
	}()

	state.TimerMgr.StopAll()

	state.cleanupSandbox()

	if state.PlayerId <= 0 || state.RoleId == 0 {
		return
	}

	if state.SendReLoginCallback != nil {
		defer state.SendReLoginCallback()
	}

	// 跨服登陆
	if state.CrossInfo != nil {
		Player.LogoutFromCross(state.PlayerId, session)
		state.CrossInfo = nil
		return
	}

	// 如果是互动服务就不需要处理session状态
	if config.ServerCfg.EnableGlobalServer {
		Player.LogoutFromGlobal(state.CId, state.PlayerId, session)
		return
	}

	PrepareStoreEvent.SafeExecute(session)

	// 会话期间未发生异常才保存会话状态
	if !state.occurPanic {
		StoreSessionState(state)
	}
}

// 丢弃保存的SessionState（当登陆不需要restore会话时）
func TryDropSessionState(playerId int64, newState *SessionState) {
	state := ReStoreSessionState(playerId)
	if state == nil {
		return
	}
	state.cleanupSandbox()

	state.MoneyState = newState.MoneyState

	// 玩家上一次的sessionState需要保存数据，但是现在没有启动事务，所以用玩家新的sessionState起的事务来处理
	// 替换之前要保证玩家新的sessionState已经mount了玩家db
	state.Database = newState.Database

	// 回写状态数据到存储设备. TryDropSessionState已是在事务中执行。不需要在SafeExecute里再启动事务
	DisposeEvent.SafeExecute(state)
	return
}

func State(session *net.Session) *SessionState {
	return session.State.(*SessionState)
}

func Transaction(state *SessionState, transId uint32, work func()) {
	if state.sandBoxDatabase == nil {
		state.Database.Transaction(transId, work)
	} else {
		log.Debugf("player %d transaction: use sandbox db\n", state.PlayerId)
		state.sandBoxDatabase.Transaction(transId, work)
	}
}

func (s *SessionState) GetSealedBookRecord() *SealedBookRecord {
	if s.sealedBook == nil {
		s.sealedBook = &SealedBookRecord{}
		s.sealedBook.Load(s.Database)
	}
	return s.sealedBook
}
