package despair_land_api

import "core/net"

type Request interface {
	Process(*net.Session)
	TypeName() string
	GetModuleIdAndActionId() (int8, int8)
}

var (
	g_InHandler  InHandler
	g_OutHandler OutHandler
)

func SetInHandler(handler InHandler) {
	g_InHandler = handler
}

func SetOutHandler(handler OutHandler) {
	g_OutHandler = handler
}

type InHandler interface {
	DespairLandInfo(*net.Session, *DespairLandInfo_In)
	DespairLandCampInfo(*net.Session, *DespairLandCampInfo_In)
	DespairLandCampPlayerInfo(*net.Session, *DespairLandCampPlayerInfo_In)
	DespairLandPickBox(*net.Session, *DespairLandPickBox_In)
	DespairLandWatchRecord(*net.Session, *DespairLandWatchRecord_In)
	DespairLandBuyBattleNum(*net.Session, *DespairLandBuyBattleNum_In)
	DespairLandBossInfo(*net.Session, *DespairLandBossInfo_In)
	DespairLandBuyBossBattleNum(*net.Session, *DespairLandBuyBossBattleNum_In)
	DespairLandPickThreeStarBox(*net.Session, *DespairLandPickThreeStarBox_In)
	DespairLandBattleBossAwardInfo(*net.Session, *DespairLandBattleBossAwardInfo_In)
	DespairLandBossOpenInfo(*net.Session, *DespairLandBossOpenInfo_In)
}

type OutHandler interface {
	DespairLandInfo(*net.Session, *DespairLandInfo_Out)
	DespairLandCampInfo(*net.Session, *DespairLandCampInfo_Out)
	DespairLandCampPlayerInfo(*net.Session, *DespairLandCampPlayerInfo_Out)
	DespairLandPickBox(*net.Session, *DespairLandPickBox_Out)
	DespairLandWatchRecord(*net.Session, *DespairLandWatchRecord_Out)
	DespairLandBuyBattleNum(*net.Session, *DespairLandBuyBattleNum_Out)
	DespairLandBossInfo(*net.Session, *DespairLandBossInfo_Out)
	DespairLandNotifyBossOpen(*net.Session, *DespairLandNotifyBossOpen_Out)
	DespairLandNotifyBossDead(*net.Session, *DespairLandNotifyBossDead_Out)
	DespairLandBuyBossBattleNum(*net.Session, *DespairLandBuyBossBattleNum_Out)
	DespairLandNotifyPass(*net.Session, *DespairLandNotifyPass_Out)
	DespairLandPickThreeStarBox(*net.Session, *DespairLandPickThreeStarBox_Out)
	DespairLandBattleBossAwardInfo(*net.Session, *DespairLandBattleBossAwardInfo_Out)
	DespairLandBossOpenInfo(*net.Session, *DespairLandBossOpenInfo_Out)
	DespairLandNotifyWeeklyRefresh(*net.Session, *DespairLandNotifyWeeklyRefresh_Out)
}

func DecodeIn(message []byte) Request {
	var actionId = message[0]
	var buffer = net.NewBuffer(message[1:])
	switch actionId {
	case 1:
		request := new(DespairLandInfo_In)
		request.Decode(buffer)
		return request
	case 2:
		request := new(DespairLandCampInfo_In)
		request.Decode(buffer)
		return request
	case 3:
		request := new(DespairLandCampPlayerInfo_In)
		request.Decode(buffer)
		return request
	case 4:
		request := new(DespairLandPickBox_In)
		request.Decode(buffer)
		return request
	case 5:
		request := new(DespairLandWatchRecord_In)
		request.Decode(buffer)
		return request
	case 6:
		request := new(DespairLandBuyBattleNum_In)
		request.Decode(buffer)
		return request
	case 7:
		request := new(DespairLandBossInfo_In)
		request.Decode(buffer)
		return request
	case 10:
		request := new(DespairLandBuyBossBattleNum_In)
		request.Decode(buffer)
		return request
	case 12:
		request := new(DespairLandPickThreeStarBox_In)
		request.Decode(buffer)
		return request
	case 13:
		request := new(DespairLandBattleBossAwardInfo_In)
		request.Decode(buffer)
		return request
	case 14:
		request := new(DespairLandBossOpenInfo_In)
		request.Decode(buffer)
		return request
	}
	_ = buffer
	panic("unsupported request")
}

func DecodeOut(message []byte) Request {
	var actionId = message[0]
	var buffer = net.NewBuffer(message[1:])
	switch actionId {
	case 1:
		request := new(DespairLandInfo_Out)
		request.Decode(buffer)
		return request
	case 2:
		request := new(DespairLandCampInfo_Out)
		request.Decode(buffer)
		return request
	case 3:
		request := new(DespairLandCampPlayerInfo_Out)
		request.Decode(buffer)
		return request
	case 4:
		request := new(DespairLandPickBox_Out)
		request.Decode(buffer)
		return request
	case 5:
		request := new(DespairLandWatchRecord_Out)
		request.Decode(buffer)
		return request
	case 6:
		request := new(DespairLandBuyBattleNum_Out)
		request.Decode(buffer)
		return request
	case 7:
		request := new(DespairLandBossInfo_Out)
		request.Decode(buffer)
		return request
	case 8:
		request := new(DespairLandNotifyBossOpen_Out)
		request.Decode(buffer)
		return request
	case 9:
		request := new(DespairLandNotifyBossDead_Out)
		request.Decode(buffer)
		return request
	case 10:
		request := new(DespairLandBuyBossBattleNum_Out)
		request.Decode(buffer)
		return request
	case 11:
		request := new(DespairLandNotifyPass_Out)
		request.Decode(buffer)
		return request
	case 12:
		request := new(DespairLandPickThreeStarBox_Out)
		request.Decode(buffer)
		return request
	case 13:
		request := new(DespairLandBattleBossAwardInfo_Out)
		request.Decode(buffer)
		return request
	case 14:
		request := new(DespairLandBossOpenInfo_Out)
		request.Decode(buffer)
		return request
	case 15:
		request := new(DespairLandNotifyWeeklyRefresh_Out)
		request.Decode(buffer)
		return request
	}
	_ = buffer
	panic("unsupported response")
}

type DespairLandCamp int8

const (
	DESPAIR_LAND_CAMP_BEAST        DespairLandCamp = 1
	DESPAIR_LAND_CAMP_WALKING_DEAD DespairLandCamp = 2
	DESPAIR_LAND_CAMP_DEVIL        DespairLandCamp = 3
)

type CampBossStatus int8

const (
	CAMP_BOSS_STATUS_CLOSING  CampBossStatus = 0
	CAMP_BOSS_STATUS_OPENING  CampBossStatus = 1
	CAMP_BOSS_STATUS_FINISHED CampBossStatus = 2
)

type DespairLandBattleRecordType int8

const (
	DESPAIR_LAND_BATTLE_RECORD_TYPE_FIRST         DespairLandBattleRecordType = 1
	DESPAIR_LAND_BATTLE_RECORD_TYPE_LATEST        DespairLandBattleRecordType = 2
	DESPAIR_LAND_BATTLE_RECORD_TYPE_LOW_FIGHT_NUM DespairLandBattleRecordType = 3
)

type DespairLandInfo_In struct {
}

func (this *DespairLandInfo_In) Process(session *net.Session) {
	g_InHandler.DespairLandInfo(session, this)
}

func (this *DespairLandInfo_In) TypeName() string {
	return "despair_land.despair_land_info.in"
}

func (this *DespairLandInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 37, 1
}

type DespairLandInfo_Out struct {
	CampInfos []DespairLandInfo_Out_CampInfos `json:"camp_infos"`
	KillNum   int64                           `json:"kill_num"`
	DeadNum   int64                           `json:"dead_num"`
	Ranks     []DespairLandInfo_Out_Ranks     `json:"ranks"`
}

type DespairLandInfo_Out_CampInfos struct {
	CampType        DespairLandCamp `json:"camp_type"`
	BattlePlayerNum int64           `json:"battle_player_num"`
	BattlePoint     int64           `json:"battle_point"`
	BattleLevel     int16           `json:"battle_level"`
}

type DespairLandInfo_Out_Ranks struct {
	Rank  int64  `json:"rank"`
	Pid   int64  `json:"pid"`
	Name  []byte `json:"name"`
	Point int64  `json:"point"`
}

func (this *DespairLandInfo_Out) Process(session *net.Session) {
	g_OutHandler.DespairLandInfo(session, this)
}

func (this *DespairLandInfo_Out) TypeName() string {
	return "despair_land.despair_land_info.out"
}

func (this *DespairLandInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 37, 1
}

func (this *DespairLandInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type DespairLandCampInfo_In struct {
	CampType DespairLandCamp `json:"camp_type"`
}

func (this *DespairLandCampInfo_In) Process(session *net.Session) {
	g_InHandler.DespairLandCampInfo(session, this)
}

func (this *DespairLandCampInfo_In) TypeName() string {
	return "despair_land.despair_land_camp_info.in"
}

func (this *DespairLandCampInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 37, 2
}

type DespairLandCampInfo_Out struct {
	BattlePoint            int64                                `json:"battle_point"`
	DailyBattleNum         int16                                `json:"daily_battle_num"`
	DailyBoughtBattleNum   int16                                `json:"daily_bought_battle_num"`
	TokenMilestoneAwardNum int16                                `json:"token_milestone_award_num"`
	LevelInfos             []DespairLandCampInfo_Out_LevelInfos `json:"level_infos"`
}

type DespairLandCampInfo_Out_LevelInfos struct {
	LevelId          int32 `json:"level_id"`
	Round            int8  `json:"round"`
	MilestoneAwarded bool  `json:"milestone_awarded"`
	PerfectAwarded   bool  `json:"perfect_awarded"`
}

func (this *DespairLandCampInfo_Out) Process(session *net.Session) {
	g_OutHandler.DespairLandCampInfo(session, this)
}

func (this *DespairLandCampInfo_Out) TypeName() string {
	return "despair_land.despair_land_camp_info.out"
}

func (this *DespairLandCampInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 37, 2
}

func (this *DespairLandCampInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type DespairLandCampPlayerInfo_In struct {
	LevelId int32 `json:"level_id"`
}

func (this *DespairLandCampPlayerInfo_In) Process(session *net.Session) {
	g_InHandler.DespairLandCampPlayerInfo(session, this)
}

func (this *DespairLandCampPlayerInfo_In) TypeName() string {
	return "despair_land.despair_land_camp_player_info.in"
}

func (this *DespairLandCampPlayerInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 37, 3
}

type DespairLandCampPlayerInfo_Out struct {
	EarliestPid        int64  `json:"earliest_pid"`
	EarliestName       []byte `json:"earliest_name"`
	ClosestPid         int64  `json:"closest_pid"`
	ClosestName        []byte `json:"closest_name"`
	LowestFightnumPid  int64  `json:"lowest_fightnum_pid"`
	LowestFightnumName []byte `json:"lowest_fightnum_name"`
}

func (this *DespairLandCampPlayerInfo_Out) Process(session *net.Session) {
	g_OutHandler.DespairLandCampPlayerInfo(session, this)
}

func (this *DespairLandCampPlayerInfo_Out) TypeName() string {
	return "despair_land.despair_land_camp_player_info.out"
}

func (this *DespairLandCampPlayerInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 37, 3
}

func (this *DespairLandCampPlayerInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type DespairLandPickBox_In struct {
	LevelId  int32           `json:"level_id"`
	CampType DespairLandCamp `json:"camp_type"`
}

func (this *DespairLandPickBox_In) Process(session *net.Session) {
	g_InHandler.DespairLandPickBox(session, this)
}

func (this *DespairLandPickBox_In) TypeName() string {
	return "despair_land.despair_land_pick_box.in"
}

func (this *DespairLandPickBox_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 37, 4
}

type DespairLandPickBox_Out struct {
	Result bool `json:"result"`
}

func (this *DespairLandPickBox_Out) Process(session *net.Session) {
	g_OutHandler.DespairLandPickBox(session, this)
}

func (this *DespairLandPickBox_Out) TypeName() string {
	return "despair_land.despair_land_pick_box.out"
}

func (this *DespairLandPickBox_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 37, 4
}

func (this *DespairLandPickBox_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type DespairLandWatchRecord_In struct {
	RecordType DespairLandBattleRecordType `json:"record_type"`
	CampType   DespairLandCamp             `json:"camp_type"`
	LevelId    int32                       `json:"level_id"`
}

func (this *DespairLandWatchRecord_In) Process(session *net.Session) {
	g_InHandler.DespairLandWatchRecord(session, this)
}

func (this *DespairLandWatchRecord_In) TypeName() string {
	return "despair_land.despair_land_watch_record.in"
}

func (this *DespairLandWatchRecord_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 37, 5
}

type DespairLandWatchRecord_Out struct {
	Ok      bool   `json:"ok"`
	Pid     int64  `json:"pid"`
	Nick    []byte `json:"nick"`
	LevelId int32  `json:"level_id"`
}

func (this *DespairLandWatchRecord_Out) Process(session *net.Session) {
	g_OutHandler.DespairLandWatchRecord(session, this)
}

func (this *DespairLandWatchRecord_Out) TypeName() string {
	return "despair_land.despair_land_watch_record.out"
}

func (this *DespairLandWatchRecord_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 37, 5
}

func (this *DespairLandWatchRecord_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type DespairLandBuyBattleNum_In struct {
}

func (this *DespairLandBuyBattleNum_In) Process(session *net.Session) {
	g_InHandler.DespairLandBuyBattleNum(session, this)
}

func (this *DespairLandBuyBattleNum_In) TypeName() string {
	return "despair_land.despair_land_buy_battle_num.in"
}

func (this *DespairLandBuyBattleNum_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 37, 6
}

type DespairLandBuyBattleNum_Out struct {
	Result bool `json:"result"`
}

func (this *DespairLandBuyBattleNum_Out) Process(session *net.Session) {
	g_OutHandler.DespairLandBuyBattleNum(session, this)
}

func (this *DespairLandBuyBattleNum_Out) TypeName() string {
	return "despair_land.despair_land_buy_battle_num.out"
}

func (this *DespairLandBuyBattleNum_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 37, 6
}

func (this *DespairLandBuyBattleNum_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type DespairLandBossInfo_In struct {
	CampType DespairLandCamp `json:"camp_type"`
}

func (this *DespairLandBossInfo_In) Process(session *net.Session) {
	g_InHandler.DespairLandBossInfo(session, this)
}

func (this *DespairLandBossInfo_In) TypeName() string {
	return "despair_land.despair_land_boss_info.in"
}

func (this *DespairLandBossInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 37, 7
}

type DespairLandBossInfo_Out struct {
	Point                   int64                           `json:"point"`
	Hp                      int64                           `json:"hp"`
	MaxHp                   int64                           `json:"max_hp"`
	BattleNum               int64                           `json:"battle_num"`
	DeadNum                 int64                           `json:"dead_num"`
	BossOpenStamp           int64                           `json:"boss_open_stamp"`
	BossLevel               int16                           `json:"boss_level"`
	DailyBossBattleNum      int16                           `json:"daily_boss_battle_num"`
	DailyCampBossBoughtNum  int16                           `json:"daily_camp_boss_bought_num"`
	DailyTotalBossBoughtNum int16                           `json:"daily_total_boss_bought_num"`
	Ranks                   []DespairLandBossInfo_Out_Ranks `json:"ranks"`
}

type DespairLandBossInfo_Out_Ranks struct {
	Rank int64  `json:"rank"`
	Pid  int64  `json:"pid"`
	Name []byte `json:"name"`
	Hurt int64  `json:"hurt"`
}

func (this *DespairLandBossInfo_Out) Process(session *net.Session) {
	g_OutHandler.DespairLandBossInfo(session, this)
}

func (this *DespairLandBossInfo_Out) TypeName() string {
	return "despair_land.despair_land_boss_info.out"
}

func (this *DespairLandBossInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 37, 7
}

func (this *DespairLandBossInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type DespairLandNotifyBossOpen_Out struct {
	CampType      DespairLandCamp `json:"camp_type"`
	BossOpenStamp int64           `json:"boss_open_stamp"`
}

func (this *DespairLandNotifyBossOpen_Out) Process(session *net.Session) {
	g_OutHandler.DespairLandNotifyBossOpen(session, this)
}

func (this *DespairLandNotifyBossOpen_Out) TypeName() string {
	return "despair_land.despair_land_notify_boss_open.out"
}

func (this *DespairLandNotifyBossOpen_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 37, 8
}

func (this *DespairLandNotifyBossOpen_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type DespairLandNotifyBossDead_Out struct {
	CampType DespairLandCamp                       `json:"camp_type"`
	Ranks    []DespairLandNotifyBossDead_Out_Ranks `json:"ranks"`
}

type DespairLandNotifyBossDead_Out_Ranks struct {
	Rank int64  `json:"rank"`
	Pid  int64  `json:"pid"`
	Name []byte `json:"name"`
}

func (this *DespairLandNotifyBossDead_Out) Process(session *net.Session) {
	g_OutHandler.DespairLandNotifyBossDead(session, this)
}

func (this *DespairLandNotifyBossDead_Out) TypeName() string {
	return "despair_land.despair_land_notify_boss_dead.out"
}

func (this *DespairLandNotifyBossDead_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 37, 9
}

func (this *DespairLandNotifyBossDead_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type DespairLandBuyBossBattleNum_In struct {
	CampType DespairLandCamp `json:"camp_type"`
}

func (this *DespairLandBuyBossBattleNum_In) Process(session *net.Session) {
	g_InHandler.DespairLandBuyBossBattleNum(session, this)
}

func (this *DespairLandBuyBossBattleNum_In) TypeName() string {
	return "despair_land.despair_land_buy_boss_battle_num.in"
}

func (this *DespairLandBuyBossBattleNum_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 37, 10
}

type DespairLandBuyBossBattleNum_Out struct {
	Result bool `json:"result"`
}

func (this *DespairLandBuyBossBattleNum_Out) Process(session *net.Session) {
	g_OutHandler.DespairLandBuyBossBattleNum(session, this)
}

func (this *DespairLandBuyBossBattleNum_Out) TypeName() string {
	return "despair_land.despair_land_buy_boss_battle_num.out"
}

func (this *DespairLandBuyBossBattleNum_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 37, 10
}

func (this *DespairLandBuyBossBattleNum_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type DespairLandNotifyPass_Out struct {
	LevelId int32  `json:"level_id"`
	Round   int8   `json:"round"`
	Pid     int64  `json:"pid"`
	Name    []byte `json:"name"`
}

func (this *DespairLandNotifyPass_Out) Process(session *net.Session) {
	g_OutHandler.DespairLandNotifyPass(session, this)
}

func (this *DespairLandNotifyPass_Out) TypeName() string {
	return "despair_land.despair_land_notify_pass.out"
}

func (this *DespairLandNotifyPass_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 37, 11
}

func (this *DespairLandNotifyPass_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type DespairLandPickThreeStarBox_In struct {
	LevelId  int32           `json:"level_id"`
	CampType DespairLandCamp `json:"camp_type"`
}

func (this *DespairLandPickThreeStarBox_In) Process(session *net.Session) {
	g_InHandler.DespairLandPickThreeStarBox(session, this)
}

func (this *DespairLandPickThreeStarBox_In) TypeName() string {
	return "despair_land.despair_land_pick_three_star_box.in"
}

func (this *DespairLandPickThreeStarBox_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 37, 12
}

type DespairLandPickThreeStarBox_Out struct {
	Result bool `json:"result"`
}

func (this *DespairLandPickThreeStarBox_Out) Process(session *net.Session) {
	g_OutHandler.DespairLandPickThreeStarBox(session, this)
}

func (this *DespairLandPickThreeStarBox_Out) TypeName() string {
	return "despair_land.despair_land_pick_three_star_box.out"
}

func (this *DespairLandPickThreeStarBox_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 37, 12
}

func (this *DespairLandPickThreeStarBox_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type DespairLandBattleBossAwardInfo_In struct {
	CampType DespairLandCamp `json:"camp_type"`
}

func (this *DespairLandBattleBossAwardInfo_In) Process(session *net.Session) {
	g_InHandler.DespairLandBattleBossAwardInfo(session, this)
}

func (this *DespairLandBattleBossAwardInfo_In) TypeName() string {
	return "despair_land.despair_land_battle_boss_award_info.in"
}

func (this *DespairLandBattleBossAwardInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 37, 13
}

type DespairLandBattleBossAwardInfo_Out struct {
	Hurt  int64 `json:"hurt"`
	Point int8  `json:"point"`
}

func (this *DespairLandBattleBossAwardInfo_Out) Process(session *net.Session) {
	g_OutHandler.DespairLandBattleBossAwardInfo(session, this)
}

func (this *DespairLandBattleBossAwardInfo_Out) TypeName() string {
	return "despair_land.despair_land_battle_boss_award_info.out"
}

func (this *DespairLandBattleBossAwardInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 37, 13
}

func (this *DespairLandBattleBossAwardInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type DespairLandBossOpenInfo_In struct {
}

func (this *DespairLandBossOpenInfo_In) Process(session *net.Session) {
	g_InHandler.DespairLandBossOpenInfo(session, this)
}

func (this *DespairLandBossOpenInfo_In) TypeName() string {
	return "despair_land.despair_land_boss_open_info.in"
}

func (this *DespairLandBossOpenInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 37, 14
}

type DespairLandBossOpenInfo_Out struct {
	Info []DespairLandBossOpenInfo_Out_Info `json:"info"`
}

type DespairLandBossOpenInfo_Out_Info struct {
	CampType      DespairLandCamp `json:"camp_type"`
	BossOpenStamp int64           `json:"boss_open_stamp"`
	Status        int8            `json:"status"`
}

func (this *DespairLandBossOpenInfo_Out) Process(session *net.Session) {
	g_OutHandler.DespairLandBossOpenInfo(session, this)
}

func (this *DespairLandBossOpenInfo_Out) TypeName() string {
	return "despair_land.despair_land_boss_open_info.out"
}

func (this *DespairLandBossOpenInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 37, 14
}

func (this *DespairLandBossOpenInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type DespairLandNotifyWeeklyRefresh_Out struct {
}

func (this *DespairLandNotifyWeeklyRefresh_Out) Process(session *net.Session) {
	g_OutHandler.DespairLandNotifyWeeklyRefresh(session, this)
}

func (this *DespairLandNotifyWeeklyRefresh_Out) TypeName() string {
	return "despair_land.despair_land_notify_weekly_refresh.out"
}

func (this *DespairLandNotifyWeeklyRefresh_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 37, 15
}

func (this *DespairLandNotifyWeeklyRefresh_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

func (this *DespairLandInfo_In) Decode(buffer *net.Buffer) {
}

func (this *DespairLandInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(37)
	buffer.WriteUint8(1)
}

func (this *DespairLandInfo_In) ByteSize() int {
	size := 2
	return size
}

func (this *DespairLandInfo_Out) Decode(buffer *net.Buffer) {
	this.CampInfos = make([]DespairLandInfo_Out_CampInfos, buffer.ReadUint8())
	for i := 0; i < len(this.CampInfos); i++ {
		this.CampInfos[i].Decode(buffer)
	}
	this.KillNum = int64(buffer.ReadUint64LE())
	this.DeadNum = int64(buffer.ReadUint64LE())
	this.Ranks = make([]DespairLandInfo_Out_Ranks, buffer.ReadUint8())
	for i := 0; i < len(this.Ranks); i++ {
		this.Ranks[i].Decode(buffer)
	}
}

func (this *DespairLandInfo_Out_CampInfos) Decode(buffer *net.Buffer) {
	this.CampType = DespairLandCamp(buffer.ReadUint8())
	this.BattlePlayerNum = int64(buffer.ReadUint64LE())
	this.BattlePoint = int64(buffer.ReadUint64LE())
	this.BattleLevel = int16(buffer.ReadUint16LE())
}

func (this *DespairLandInfo_Out_Ranks) Decode(buffer *net.Buffer) {
	this.Rank = int64(buffer.ReadUint64LE())
	this.Pid = int64(buffer.ReadUint64LE())
	this.Name = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Point = int64(buffer.ReadUint64LE())
}

func (this *DespairLandInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(37)
	buffer.WriteUint8(1)
	buffer.WriteUint8(uint8(len(this.CampInfos)))
	for i := 0; i < len(this.CampInfos); i++ {
		this.CampInfos[i].Encode(buffer)
	}
	buffer.WriteUint64LE(uint64(this.KillNum))
	buffer.WriteUint64LE(uint64(this.DeadNum))
	buffer.WriteUint8(uint8(len(this.Ranks)))
	for i := 0; i < len(this.Ranks); i++ {
		this.Ranks[i].Encode(buffer)
	}
}

func (this *DespairLandInfo_Out_CampInfos) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(this.CampType))
	buffer.WriteUint64LE(uint64(this.BattlePlayerNum))
	buffer.WriteUint64LE(uint64(this.BattlePoint))
	buffer.WriteUint16LE(uint16(this.BattleLevel))
}

func (this *DespairLandInfo_Out_Ranks) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.Rank))
	buffer.WriteUint64LE(uint64(this.Pid))
	buffer.WriteUint16LE(uint16(len(this.Name)))
	buffer.WriteBytes(this.Name)
	buffer.WriteUint64LE(uint64(this.Point))
}

func (this *DespairLandInfo_Out) ByteSize() int {
	size := 20
	size += len(this.CampInfos) * 19
	for i := 0; i < len(this.Ranks); i++ {
		size += this.Ranks[i].ByteSize()
	}
	return size
}

func (this *DespairLandInfo_Out_Ranks) ByteSize() int {
	size := 26
	size += len(this.Name)
	return size
}

func (this *DespairLandCampInfo_In) Decode(buffer *net.Buffer) {
	this.CampType = DespairLandCamp(buffer.ReadUint8())
}

func (this *DespairLandCampInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(37)
	buffer.WriteUint8(2)
	buffer.WriteUint8(uint8(this.CampType))
}

func (this *DespairLandCampInfo_In) ByteSize() int {
	size := 3
	return size
}

func (this *DespairLandCampInfo_Out) Decode(buffer *net.Buffer) {
	this.BattlePoint = int64(buffer.ReadUint64LE())
	this.DailyBattleNum = int16(buffer.ReadUint16LE())
	this.DailyBoughtBattleNum = int16(buffer.ReadUint16LE())
	this.TokenMilestoneAwardNum = int16(buffer.ReadUint16LE())
	this.LevelInfos = make([]DespairLandCampInfo_Out_LevelInfos, buffer.ReadUint8())
	for i := 0; i < len(this.LevelInfos); i++ {
		this.LevelInfos[i].Decode(buffer)
	}
}

func (this *DespairLandCampInfo_Out_LevelInfos) Decode(buffer *net.Buffer) {
	this.LevelId = int32(buffer.ReadUint32LE())
	this.Round = int8(buffer.ReadUint8())
	this.MilestoneAwarded = buffer.ReadUint8() == 1
	this.PerfectAwarded = buffer.ReadUint8() == 1
}

func (this *DespairLandCampInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(37)
	buffer.WriteUint8(2)
	buffer.WriteUint64LE(uint64(this.BattlePoint))
	buffer.WriteUint16LE(uint16(this.DailyBattleNum))
	buffer.WriteUint16LE(uint16(this.DailyBoughtBattleNum))
	buffer.WriteUint16LE(uint16(this.TokenMilestoneAwardNum))
	buffer.WriteUint8(uint8(len(this.LevelInfos)))
	for i := 0; i < len(this.LevelInfos); i++ {
		this.LevelInfos[i].Encode(buffer)
	}
}

func (this *DespairLandCampInfo_Out_LevelInfos) Encode(buffer *net.Buffer) {
	buffer.WriteUint32LE(uint32(this.LevelId))
	buffer.WriteUint8(uint8(this.Round))
	if this.MilestoneAwarded {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
	if this.PerfectAwarded {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *DespairLandCampInfo_Out) ByteSize() int {
	size := 17
	size += len(this.LevelInfos) * 7
	return size
}

func (this *DespairLandCampPlayerInfo_In) Decode(buffer *net.Buffer) {
	this.LevelId = int32(buffer.ReadUint32LE())
}

func (this *DespairLandCampPlayerInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(37)
	buffer.WriteUint8(3)
	buffer.WriteUint32LE(uint32(this.LevelId))
}

func (this *DespairLandCampPlayerInfo_In) ByteSize() int {
	size := 6
	return size
}

func (this *DespairLandCampPlayerInfo_Out) Decode(buffer *net.Buffer) {
	this.EarliestPid = int64(buffer.ReadUint64LE())
	this.EarliestName = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.ClosestPid = int64(buffer.ReadUint64LE())
	this.ClosestName = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.LowestFightnumPid = int64(buffer.ReadUint64LE())
	this.LowestFightnumName = buffer.ReadBytes(int(buffer.ReadUint16LE()))
}

func (this *DespairLandCampPlayerInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(37)
	buffer.WriteUint8(3)
	buffer.WriteUint64LE(uint64(this.EarliestPid))
	buffer.WriteUint16LE(uint16(len(this.EarliestName)))
	buffer.WriteBytes(this.EarliestName)
	buffer.WriteUint64LE(uint64(this.ClosestPid))
	buffer.WriteUint16LE(uint16(len(this.ClosestName)))
	buffer.WriteBytes(this.ClosestName)
	buffer.WriteUint64LE(uint64(this.LowestFightnumPid))
	buffer.WriteUint16LE(uint16(len(this.LowestFightnumName)))
	buffer.WriteBytes(this.LowestFightnumName)
}

func (this *DespairLandCampPlayerInfo_Out) ByteSize() int {
	size := 32
	size += len(this.EarliestName)
	size += len(this.ClosestName)
	size += len(this.LowestFightnumName)
	return size
}

func (this *DespairLandPickBox_In) Decode(buffer *net.Buffer) {
	this.LevelId = int32(buffer.ReadUint32LE())
	this.CampType = DespairLandCamp(buffer.ReadUint8())
}

func (this *DespairLandPickBox_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(37)
	buffer.WriteUint8(4)
	buffer.WriteUint32LE(uint32(this.LevelId))
	buffer.WriteUint8(uint8(this.CampType))
}

func (this *DespairLandPickBox_In) ByteSize() int {
	size := 7
	return size
}

func (this *DespairLandPickBox_Out) Decode(buffer *net.Buffer) {
	this.Result = buffer.ReadUint8() == 1
}

func (this *DespairLandPickBox_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(37)
	buffer.WriteUint8(4)
	if this.Result {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *DespairLandPickBox_Out) ByteSize() int {
	size := 3
	return size
}

func (this *DespairLandWatchRecord_In) Decode(buffer *net.Buffer) {
	this.RecordType = DespairLandBattleRecordType(buffer.ReadUint8())
	this.CampType = DespairLandCamp(buffer.ReadUint8())
	this.LevelId = int32(buffer.ReadUint32LE())
}

func (this *DespairLandWatchRecord_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(37)
	buffer.WriteUint8(5)
	buffer.WriteUint8(uint8(this.RecordType))
	buffer.WriteUint8(uint8(this.CampType))
	buffer.WriteUint32LE(uint32(this.LevelId))
}

func (this *DespairLandWatchRecord_In) ByteSize() int {
	size := 8
	return size
}

func (this *DespairLandWatchRecord_Out) Decode(buffer *net.Buffer) {
	this.Ok = buffer.ReadUint8() == 1
	this.Pid = int64(buffer.ReadUint64LE())
	this.Nick = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.LevelId = int32(buffer.ReadUint32LE())
}

func (this *DespairLandWatchRecord_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(37)
	buffer.WriteUint8(5)
	if this.Ok {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
	buffer.WriteUint64LE(uint64(this.Pid))
	buffer.WriteUint16LE(uint16(len(this.Nick)))
	buffer.WriteBytes(this.Nick)
	buffer.WriteUint32LE(uint32(this.LevelId))
}

func (this *DespairLandWatchRecord_Out) ByteSize() int {
	size := 17
	size += len(this.Nick)
	return size
}

func (this *DespairLandBuyBattleNum_In) Decode(buffer *net.Buffer) {
}

func (this *DespairLandBuyBattleNum_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(37)
	buffer.WriteUint8(6)
}

func (this *DespairLandBuyBattleNum_In) ByteSize() int {
	size := 2
	return size
}

func (this *DespairLandBuyBattleNum_Out) Decode(buffer *net.Buffer) {
	this.Result = buffer.ReadUint8() == 1
}

func (this *DespairLandBuyBattleNum_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(37)
	buffer.WriteUint8(6)
	if this.Result {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *DespairLandBuyBattleNum_Out) ByteSize() int {
	size := 3
	return size
}

func (this *DespairLandBossInfo_In) Decode(buffer *net.Buffer) {
	this.CampType = DespairLandCamp(buffer.ReadUint8())
}

func (this *DespairLandBossInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(37)
	buffer.WriteUint8(7)
	buffer.WriteUint8(uint8(this.CampType))
}

func (this *DespairLandBossInfo_In) ByteSize() int {
	size := 3
	return size
}

func (this *DespairLandBossInfo_Out) Decode(buffer *net.Buffer) {
	this.Point = int64(buffer.ReadUint64LE())
	this.Hp = int64(buffer.ReadUint64LE())
	this.MaxHp = int64(buffer.ReadUint64LE())
	this.BattleNum = int64(buffer.ReadUint64LE())
	this.DeadNum = int64(buffer.ReadUint64LE())
	this.BossOpenStamp = int64(buffer.ReadUint64LE())
	this.BossLevel = int16(buffer.ReadUint16LE())
	this.DailyBossBattleNum = int16(buffer.ReadUint16LE())
	this.DailyCampBossBoughtNum = int16(buffer.ReadUint16LE())
	this.DailyTotalBossBoughtNum = int16(buffer.ReadUint16LE())
	this.Ranks = make([]DespairLandBossInfo_Out_Ranks, buffer.ReadUint8())
	for i := 0; i < len(this.Ranks); i++ {
		this.Ranks[i].Decode(buffer)
	}
}

func (this *DespairLandBossInfo_Out_Ranks) Decode(buffer *net.Buffer) {
	this.Rank = int64(buffer.ReadUint64LE())
	this.Pid = int64(buffer.ReadUint64LE())
	this.Name = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Hurt = int64(buffer.ReadUint64LE())
}

func (this *DespairLandBossInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(37)
	buffer.WriteUint8(7)
	buffer.WriteUint64LE(uint64(this.Point))
	buffer.WriteUint64LE(uint64(this.Hp))
	buffer.WriteUint64LE(uint64(this.MaxHp))
	buffer.WriteUint64LE(uint64(this.BattleNum))
	buffer.WriteUint64LE(uint64(this.DeadNum))
	buffer.WriteUint64LE(uint64(this.BossOpenStamp))
	buffer.WriteUint16LE(uint16(this.BossLevel))
	buffer.WriteUint16LE(uint16(this.DailyBossBattleNum))
	buffer.WriteUint16LE(uint16(this.DailyCampBossBoughtNum))
	buffer.WriteUint16LE(uint16(this.DailyTotalBossBoughtNum))
	buffer.WriteUint8(uint8(len(this.Ranks)))
	for i := 0; i < len(this.Ranks); i++ {
		this.Ranks[i].Encode(buffer)
	}
}

func (this *DespairLandBossInfo_Out_Ranks) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.Rank))
	buffer.WriteUint64LE(uint64(this.Pid))
	buffer.WriteUint16LE(uint16(len(this.Name)))
	buffer.WriteBytes(this.Name)
	buffer.WriteUint64LE(uint64(this.Hurt))
}

func (this *DespairLandBossInfo_Out) ByteSize() int {
	size := 59
	for i := 0; i < len(this.Ranks); i++ {
		size += this.Ranks[i].ByteSize()
	}
	return size
}

func (this *DespairLandBossInfo_Out_Ranks) ByteSize() int {
	size := 26
	size += len(this.Name)
	return size
}

func (this *DespairLandNotifyBossOpen_Out) Decode(buffer *net.Buffer) {
	this.CampType = DespairLandCamp(buffer.ReadUint8())
	this.BossOpenStamp = int64(buffer.ReadUint64LE())
}

func (this *DespairLandNotifyBossOpen_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(37)
	buffer.WriteUint8(8)
	buffer.WriteUint8(uint8(this.CampType))
	buffer.WriteUint64LE(uint64(this.BossOpenStamp))
}

func (this *DespairLandNotifyBossOpen_Out) ByteSize() int {
	size := 11
	return size
}

func (this *DespairLandNotifyBossDead_Out) Decode(buffer *net.Buffer) {
	this.CampType = DespairLandCamp(buffer.ReadUint8())
	this.Ranks = make([]DespairLandNotifyBossDead_Out_Ranks, buffer.ReadUint8())
	for i := 0; i < len(this.Ranks); i++ {
		this.Ranks[i].Decode(buffer)
	}
}

func (this *DespairLandNotifyBossDead_Out_Ranks) Decode(buffer *net.Buffer) {
	this.Rank = int64(buffer.ReadUint64LE())
	this.Pid = int64(buffer.ReadUint64LE())
	this.Name = buffer.ReadBytes(int(buffer.ReadUint16LE()))
}

func (this *DespairLandNotifyBossDead_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(37)
	buffer.WriteUint8(9)
	buffer.WriteUint8(uint8(this.CampType))
	buffer.WriteUint8(uint8(len(this.Ranks)))
	for i := 0; i < len(this.Ranks); i++ {
		this.Ranks[i].Encode(buffer)
	}
}

func (this *DespairLandNotifyBossDead_Out_Ranks) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.Rank))
	buffer.WriteUint64LE(uint64(this.Pid))
	buffer.WriteUint16LE(uint16(len(this.Name)))
	buffer.WriteBytes(this.Name)
}

func (this *DespairLandNotifyBossDead_Out) ByteSize() int {
	size := 4
	for i := 0; i < len(this.Ranks); i++ {
		size += this.Ranks[i].ByteSize()
	}
	return size
}

func (this *DespairLandNotifyBossDead_Out_Ranks) ByteSize() int {
	size := 18
	size += len(this.Name)
	return size
}

func (this *DespairLandBuyBossBattleNum_In) Decode(buffer *net.Buffer) {
	this.CampType = DespairLandCamp(buffer.ReadUint8())
}

func (this *DespairLandBuyBossBattleNum_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(37)
	buffer.WriteUint8(10)
	buffer.WriteUint8(uint8(this.CampType))
}

func (this *DespairLandBuyBossBattleNum_In) ByteSize() int {
	size := 3
	return size
}

func (this *DespairLandBuyBossBattleNum_Out) Decode(buffer *net.Buffer) {
	this.Result = buffer.ReadUint8() == 1
}

func (this *DespairLandBuyBossBattleNum_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(37)
	buffer.WriteUint8(10)
	if this.Result {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *DespairLandBuyBossBattleNum_Out) ByteSize() int {
	size := 3
	return size
}

func (this *DespairLandNotifyPass_Out) Decode(buffer *net.Buffer) {
	this.LevelId = int32(buffer.ReadUint32LE())
	this.Round = int8(buffer.ReadUint8())
	this.Pid = int64(buffer.ReadUint64LE())
	this.Name = buffer.ReadBytes(int(buffer.ReadUint16LE()))
}

func (this *DespairLandNotifyPass_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(37)
	buffer.WriteUint8(11)
	buffer.WriteUint32LE(uint32(this.LevelId))
	buffer.WriteUint8(uint8(this.Round))
	buffer.WriteUint64LE(uint64(this.Pid))
	buffer.WriteUint16LE(uint16(len(this.Name)))
	buffer.WriteBytes(this.Name)
}

func (this *DespairLandNotifyPass_Out) ByteSize() int {
	size := 17
	size += len(this.Name)
	return size
}

func (this *DespairLandPickThreeStarBox_In) Decode(buffer *net.Buffer) {
	this.LevelId = int32(buffer.ReadUint32LE())
	this.CampType = DespairLandCamp(buffer.ReadUint8())
}

func (this *DespairLandPickThreeStarBox_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(37)
	buffer.WriteUint8(12)
	buffer.WriteUint32LE(uint32(this.LevelId))
	buffer.WriteUint8(uint8(this.CampType))
}

func (this *DespairLandPickThreeStarBox_In) ByteSize() int {
	size := 7
	return size
}

func (this *DespairLandPickThreeStarBox_Out) Decode(buffer *net.Buffer) {
	this.Result = buffer.ReadUint8() == 1
}

func (this *DespairLandPickThreeStarBox_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(37)
	buffer.WriteUint8(12)
	if this.Result {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *DespairLandPickThreeStarBox_Out) ByteSize() int {
	size := 3
	return size
}

func (this *DespairLandBattleBossAwardInfo_In) Decode(buffer *net.Buffer) {
	this.CampType = DespairLandCamp(buffer.ReadUint8())
}

func (this *DespairLandBattleBossAwardInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(37)
	buffer.WriteUint8(13)
	buffer.WriteUint8(uint8(this.CampType))
}

func (this *DespairLandBattleBossAwardInfo_In) ByteSize() int {
	size := 3
	return size
}

func (this *DespairLandBattleBossAwardInfo_Out) Decode(buffer *net.Buffer) {
	this.Hurt = int64(buffer.ReadUint64LE())
	this.Point = int8(buffer.ReadUint8())
}

func (this *DespairLandBattleBossAwardInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(37)
	buffer.WriteUint8(13)
	buffer.WriteUint64LE(uint64(this.Hurt))
	buffer.WriteUint8(uint8(this.Point))
}

func (this *DespairLandBattleBossAwardInfo_Out) ByteSize() int {
	size := 11
	return size
}

func (this *DespairLandBossOpenInfo_In) Decode(buffer *net.Buffer) {
}

func (this *DespairLandBossOpenInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(37)
	buffer.WriteUint8(14)
}

func (this *DespairLandBossOpenInfo_In) ByteSize() int {
	size := 2
	return size
}

func (this *DespairLandBossOpenInfo_Out) Decode(buffer *net.Buffer) {
	this.Info = make([]DespairLandBossOpenInfo_Out_Info, buffer.ReadUint8())
	for i := 0; i < len(this.Info); i++ {
		this.Info[i].Decode(buffer)
	}
}

func (this *DespairLandBossOpenInfo_Out_Info) Decode(buffer *net.Buffer) {
	this.CampType = DespairLandCamp(buffer.ReadUint8())
	this.BossOpenStamp = int64(buffer.ReadUint64LE())
	this.Status = int8(buffer.ReadUint8())
}

func (this *DespairLandBossOpenInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(37)
	buffer.WriteUint8(14)
	buffer.WriteUint8(uint8(len(this.Info)))
	for i := 0; i < len(this.Info); i++ {
		this.Info[i].Encode(buffer)
	}
}

func (this *DespairLandBossOpenInfo_Out_Info) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(this.CampType))
	buffer.WriteUint64LE(uint64(this.BossOpenStamp))
	buffer.WriteUint8(uint8(this.Status))
}

func (this *DespairLandBossOpenInfo_Out) ByteSize() int {
	size := 3
	size += len(this.Info) * 10
	return size
}

func (this *DespairLandNotifyWeeklyRefresh_Out) Decode(buffer *net.Buffer) {
}

func (this *DespairLandNotifyWeeklyRefresh_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(37)
	buffer.WriteUint8(15)
}

func (this *DespairLandNotifyWeeklyRefresh_Out) ByteSize() int {
	size := 2
	return size
}
