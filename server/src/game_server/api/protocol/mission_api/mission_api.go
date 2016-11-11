package mission_api

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
	Open(*net.Session, *Open_In)
	GetMissionLevel(*net.Session, *GetMissionLevel_In)
	EnterLevel(*net.Session, *EnterLevel_In)
	OpenLevelBox(*net.Session, *OpenLevelBox_In)
	UseIngotRelive(*net.Session, *UseIngotRelive_In)
	UseItem(*net.Session, *UseItem_In)
	EnterExtendLevel(*net.Session, *EnterExtendLevel_In)
	GetExtendLevelInfo(*net.Session, *GetExtendLevelInfo_In)
	OpenSmallBox(*net.Session, *OpenSmallBox_In)
	EnterHardLevel(*net.Session, *EnterHardLevel_In)
	GetHardLevel(*net.Session, *GetHardLevel_In)
	GetBuddyLevelRoleId(*net.Session, *GetBuddyLevelRoleId_In)
	SetBuddyLevelTeam(*net.Session, *SetBuddyLevelTeam_In)
	AutoFightLevel(*net.Session, *AutoFightLevel_In)
	EnterRainbowLevel(*net.Session, *EnterRainbowLevel_In)
	OpenMengYao(*net.Session, *OpenMengYao_In)
	GetMissionLevelByItemId(*net.Session, *GetMissionLevelByItemId_In)
	BuyHardLevelTimes(*net.Session, *BuyHardLevelTimes_In)
	OpenRandomAwardBox(*net.Session, *OpenRandomAwardBox_In)
	EvaluateLevel(*net.Session, *EvaluateLevel_In)
	OpenShadedBox(*net.Session, *OpenShadedBox_In)
	GetMissionTotalStarInfo(*net.Session, *GetMissionTotalStarInfo_In)
	GetMissionTotalStarAwards(*net.Session, *GetMissionTotalStarAwards_In)
	ClearMissionLevelState(*net.Session, *ClearMissionLevelState_In)
	BuyResourceMissionLevelTimes(*net.Session, *BuyResourceMissionLevelTimes_In)
	GetDragonBall(*net.Session, *GetDragonBall_In)
	GetEventItem(*net.Session, *GetEventItem_In)
}

type OutHandler interface {
	Open(*net.Session, *Open_Out)
	GetMissionLevel(*net.Session, *GetMissionLevel_Out)
	EnterLevel(*net.Session, *EnterLevel_Out)
	OpenLevelBox(*net.Session, *OpenLevelBox_Out)
	UseIngotRelive(*net.Session, *UseIngotRelive_Out)
	UseItem(*net.Session, *UseItem_Out)
	Rebuild(*net.Session, *Rebuild_Out)
	EnterExtendLevel(*net.Session, *EnterExtendLevel_Out)
	GetExtendLevelInfo(*net.Session, *GetExtendLevelInfo_Out)
	OpenSmallBox(*net.Session, *OpenSmallBox_Out)
	NotifyCatchBattlePet(*net.Session, *NotifyCatchBattlePet_Out)
	EnterHardLevel(*net.Session, *EnterHardLevel_Out)
	GetHardLevel(*net.Session, *GetHardLevel_Out)
	GetBuddyLevelRoleId(*net.Session, *GetBuddyLevelRoleId_Out)
	SetBuddyLevelTeam(*net.Session, *SetBuddyLevelTeam_Out)
	AutoFightLevel(*net.Session, *AutoFightLevel_Out)
	EnterRainbowLevel(*net.Session, *EnterRainbowLevel_Out)
	OpenMengYao(*net.Session, *OpenMengYao_Out)
	GetMissionLevelByItemId(*net.Session, *GetMissionLevelByItemId_Out)
	BuyHardLevelTimes(*net.Session, *BuyHardLevelTimes_Out)
	OpenRandomAwardBox(*net.Session, *OpenRandomAwardBox_Out)
	EvaluateLevel(*net.Session, *EvaluateLevel_Out)
	OpenShadedBox(*net.Session, *OpenShadedBox_Out)
	GetMissionTotalStarInfo(*net.Session, *GetMissionTotalStarInfo_Out)
	GetMissionTotalStarAwards(*net.Session, *GetMissionTotalStarAwards_Out)
	ClearMissionLevelState(*net.Session, *ClearMissionLevelState_Out)
	BuyResourceMissionLevelTimes(*net.Session, *BuyResourceMissionLevelTimes_Out)
	GetDragonBall(*net.Session, *GetDragonBall_Out)
	GetEventItem(*net.Session, *GetEventItem_Out)
}

func DecodeIn(message []byte) Request {
	var actionId = message[0]
	var buffer = net.NewBuffer(message[1:])
	switch actionId {
	case 0:
		request := new(Open_In)
		request.Decode(buffer)
		return request
	case 1:
		request := new(GetMissionLevel_In)
		request.Decode(buffer)
		return request
	case 2:
		request := new(EnterLevel_In)
		request.Decode(buffer)
		return request
	case 3:
		request := new(OpenLevelBox_In)
		request.Decode(buffer)
		return request
	case 4:
		request := new(UseIngotRelive_In)
		request.Decode(buffer)
		return request
	case 5:
		request := new(UseItem_In)
		request.Decode(buffer)
		return request
	case 7:
		request := new(EnterExtendLevel_In)
		request.Decode(buffer)
		return request
	case 8:
		request := new(GetExtendLevelInfo_In)
		request.Decode(buffer)
		return request
	case 9:
		request := new(OpenSmallBox_In)
		request.Decode(buffer)
		return request
	case 11:
		request := new(EnterHardLevel_In)
		request.Decode(buffer)
		return request
	case 12:
		request := new(GetHardLevel_In)
		request.Decode(buffer)
		return request
	case 13:
		request := new(GetBuddyLevelRoleId_In)
		request.Decode(buffer)
		return request
	case 14:
		request := new(SetBuddyLevelTeam_In)
		request.Decode(buffer)
		return request
	case 15:
		request := new(AutoFightLevel_In)
		request.Decode(buffer)
		return request
	case 16:
		request := new(EnterRainbowLevel_In)
		request.Decode(buffer)
		return request
	case 17:
		request := new(OpenMengYao_In)
		request.Decode(buffer)
		return request
	case 18:
		request := new(GetMissionLevelByItemId_In)
		request.Decode(buffer)
		return request
	case 19:
		request := new(BuyHardLevelTimes_In)
		request.Decode(buffer)
		return request
	case 20:
		request := new(OpenRandomAwardBox_In)
		request.Decode(buffer)
		return request
	case 21:
		request := new(EvaluateLevel_In)
		request.Decode(buffer)
		return request
	case 22:
		request := new(OpenShadedBox_In)
		request.Decode(buffer)
		return request
	case 23:
		request := new(GetMissionTotalStarInfo_In)
		request.Decode(buffer)
		return request
	case 24:
		request := new(GetMissionTotalStarAwards_In)
		request.Decode(buffer)
		return request
	case 25:
		request := new(ClearMissionLevelState_In)
		request.Decode(buffer)
		return request
	case 26:
		request := new(BuyResourceMissionLevelTimes_In)
		request.Decode(buffer)
		return request
	case 27:
		request := new(GetDragonBall_In)
		request.Decode(buffer)
		return request
	case 28:
		request := new(GetEventItem_In)
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
	case 0:
		request := new(Open_Out)
		request.Decode(buffer)
		return request
	case 1:
		request := new(GetMissionLevel_Out)
		request.Decode(buffer)
		return request
	case 2:
		request := new(EnterLevel_Out)
		request.Decode(buffer)
		return request
	case 3:
		request := new(OpenLevelBox_Out)
		request.Decode(buffer)
		return request
	case 4:
		request := new(UseIngotRelive_Out)
		request.Decode(buffer)
		return request
	case 5:
		request := new(UseItem_Out)
		request.Decode(buffer)
		return request
	case 6:
		request := new(Rebuild_Out)
		request.Decode(buffer)
		return request
	case 7:
		request := new(EnterExtendLevel_Out)
		request.Decode(buffer)
		return request
	case 8:
		request := new(GetExtendLevelInfo_Out)
		request.Decode(buffer)
		return request
	case 9:
		request := new(OpenSmallBox_Out)
		request.Decode(buffer)
		return request
	case 10:
		request := new(NotifyCatchBattlePet_Out)
		request.Decode(buffer)
		return request
	case 11:
		request := new(EnterHardLevel_Out)
		request.Decode(buffer)
		return request
	case 12:
		request := new(GetHardLevel_Out)
		request.Decode(buffer)
		return request
	case 13:
		request := new(GetBuddyLevelRoleId_Out)
		request.Decode(buffer)
		return request
	case 14:
		request := new(SetBuddyLevelTeam_Out)
		request.Decode(buffer)
		return request
	case 15:
		request := new(AutoFightLevel_Out)
		request.Decode(buffer)
		return request
	case 16:
		request := new(EnterRainbowLevel_Out)
		request.Decode(buffer)
		return request
	case 17:
		request := new(OpenMengYao_Out)
		request.Decode(buffer)
		return request
	case 18:
		request := new(GetMissionLevelByItemId_Out)
		request.Decode(buffer)
		return request
	case 19:
		request := new(BuyHardLevelTimes_Out)
		request.Decode(buffer)
		return request
	case 20:
		request := new(OpenRandomAwardBox_Out)
		request.Decode(buffer)
		return request
	case 21:
		request := new(EvaluateLevel_Out)
		request.Decode(buffer)
		return request
	case 22:
		request := new(OpenShadedBox_Out)
		request.Decode(buffer)
		return request
	case 23:
		request := new(GetMissionTotalStarInfo_Out)
		request.Decode(buffer)
		return request
	case 24:
		request := new(GetMissionTotalStarAwards_Out)
		request.Decode(buffer)
		return request
	case 25:
		request := new(ClearMissionLevelState_Out)
		request.Decode(buffer)
		return request
	case 26:
		request := new(BuyResourceMissionLevelTimes_Out)
		request.Decode(buffer)
		return request
	case 27:
		request := new(GetDragonBall_Out)
		request.Decode(buffer)
		return request
	case 28:
		request := new(GetEventItem_Out)
		request.Decode(buffer)
		return request
	}
	_ = buffer
	panic("unsupported response")
}

type OutResult int8

const (
	OUT_RESULT_FAILED  OutResult = 0
	OUT_RESULT_SUCCEED OutResult = 1
)

type ExtendType int8

const (
	EXTEND_TYPE_RESOURCE ExtendType = 1
	EXTEND_TYPE_ACTIVITY ExtendType = 2
)

type ExtendLevelType int8

const (
	EXTEND_LEVEL_TYPE_RESOURCE ExtendLevelType = 1
	EXTEND_LEVEL_TYPE_BUDDY    ExtendLevelType = 9
	EXTEND_LEVEL_TYPE_PET      ExtendLevelType = 10
	EXTEND_LEVEL_TYPE_GHOST    ExtendLevelType = 11
)

type ExtendLevelSubType int8

const (
	EXTEND_LEVEL_SUB_TYPE_NONE ExtendLevelSubType = 0
	EXTEND_LEVEL_SUB_TYPE_COIN ExtendLevelSubType = 1
	EXTEND_LEVEL_SUB_TYPE_EXP  ExtendLevelSubType = 2
)

type BattleType int8

const (
	BATTLE_TYPE_MISSION    BattleType = 0
	BATTLE_TYPE_RESOURCE   BattleType = 1
	BATTLE_TYPE_TOWER      BattleType = 2
	BATTLE_TYPE_MULTILEVEL BattleType = 3
	BATTLE_TYPE_HARD       BattleType = 8
	BATTLE_TYPE_BUDDY      BattleType = 9
	BATTLE_TYPE_PET        BattleType = 10
	BATTLE_TYPE_GHOST      BattleType = 11
	BATTLE_TYPE_RAINBOW    BattleType = 12
	BATTLE_TYPE_PVE        BattleType = 13
	BATTLE_TYPE_DESPAIR    BattleType = 19
)

type Open_In struct {
	MissionId int16 `json:"mission_id"`
}

func (this *Open_In) Process(session *net.Session) {
	g_InHandler.Open(session, this)
}

func (this *Open_In) TypeName() string {
	return "mission.open.in"
}

func (this *Open_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 0
}

type Open_Out struct {
	Result OutResult `json:"result"`
}

func (this *Open_Out) Process(session *net.Session) {
	g_OutHandler.Open(session, this)
}

func (this *Open_Out) TypeName() string {
	return "mission.open.out"
}

func (this *Open_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 0
}

func (this *Open_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetMissionLevel_In struct {
	MissionId int16 `json:"mission_id"`
}

func (this *GetMissionLevel_In) Process(session *net.Session) {
	g_InHandler.GetMissionLevel(session, this)
}

func (this *GetMissionLevel_In) TypeName() string {
	return "mission.get_mission_level.in"
}

func (this *GetMissionLevel_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 1
}

type GetMissionLevel_Out struct {
	Result OutResult                    `json:"result"`
	Levels []GetMissionLevel_Out_Levels `json:"levels"`
}

type GetMissionLevel_Out_Levels struct {
	LevelId        int32 `json:"level_id"`
	RoundNum       int8  `json:"round_num"`
	DailyNum       int8  `json:"daily_num"`
	WaitingShadows int8  `json:"waiting_shadows"`
}

func (this *GetMissionLevel_Out) Process(session *net.Session) {
	g_OutHandler.GetMissionLevel(session, this)
}

func (this *GetMissionLevel_Out) TypeName() string {
	return "mission.get_mission_level.out"
}

func (this *GetMissionLevel_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 1
}

func (this *GetMissionLevel_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type EnterLevel_In struct {
	MissionLevelId int32 `json:"mission_level_id"`
}

func (this *EnterLevel_In) Process(session *net.Session) {
	g_InHandler.EnterLevel(session, this)
}

func (this *EnterLevel_In) TypeName() string {
	return "mission.enter_level.in"
}

func (this *EnterLevel_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 2
}

type EnterLevel_Out struct {
	Result   OutResult                 `json:"result"`
	Smallbox []EnterLevel_Out_Smallbox `json:"smallbox"`
	MengYao  []EnterLevel_Out_MengYao  `json:"meng_yao"`
	Shadow   []EnterLevel_Out_Shadow   `json:"shadow"`
}

type EnterLevel_Out_Smallbox struct {
	BoxId int32 `json:"box_id"`
}

type EnterLevel_Out_MengYao struct {
	MyId int32 `json:"my_id"`
}

type EnterLevel_Out_Shadow struct {
	ShadedId int32 `json:"shaded_id"`
}

func (this *EnterLevel_Out) Process(session *net.Session) {
	g_OutHandler.EnterLevel(session, this)
}

func (this *EnterLevel_Out) TypeName() string {
	return "mission.enter_level.out"
}

func (this *EnterLevel_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 2
}

func (this *EnterLevel_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type OpenLevelBox_In struct {
	Times int32 `json:"times"`
}

func (this *OpenLevelBox_In) Process(session *net.Session) {
	g_InHandler.OpenLevelBox(session, this)
}

func (this *OpenLevelBox_In) TypeName() string {
	return "mission.open_level_box.in"
}

func (this *OpenLevelBox_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 3
}

type OpenLevelBox_Out struct {
	AllBoxId []OpenLevelBox_Out_AllBoxId `json:"all_box_id"`
}

type OpenLevelBox_Out_AllBoxId struct {
	BoxId int64 `json:"box_id"`
}

func (this *OpenLevelBox_Out) Process(session *net.Session) {
	g_OutHandler.OpenLevelBox(session, this)
}

func (this *OpenLevelBox_Out) TypeName() string {
	return "mission.open_level_box.out"
}

func (this *OpenLevelBox_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 3
}

func (this *OpenLevelBox_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type UseIngotRelive_In struct {
}

func (this *UseIngotRelive_In) Process(session *net.Session) {
	g_InHandler.UseIngotRelive(session, this)
}

func (this *UseIngotRelive_In) TypeName() string {
	return "mission.use_ingot_relive.in"
}

func (this *UseIngotRelive_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 4
}

type UseIngotRelive_Out struct {
	Ingot int32 `json:"ingot"`
}

func (this *UseIngotRelive_Out) Process(session *net.Session) {
	g_OutHandler.UseIngotRelive(session, this)
}

func (this *UseIngotRelive_Out) TypeName() string {
	return "mission.use_ingot_relive.out"
}

func (this *UseIngotRelive_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 4
}

func (this *UseIngotRelive_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type UseItem_In struct {
	RoleId int8  `json:"role_id"`
	ItemId int16 `json:"item_id"`
}

func (this *UseItem_In) Process(session *net.Session) {
	g_InHandler.UseItem(session, this)
}

func (this *UseItem_In) TypeName() string {
	return "mission.use_item.in"
}

func (this *UseItem_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 5
}

type UseItem_Out struct {
}

func (this *UseItem_Out) Process(session *net.Session) {
	g_OutHandler.UseItem(session, this)
}

func (this *UseItem_Out) TypeName() string {
	return "mission.use_item.out"
}

func (this *UseItem_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 5
}

func (this *UseItem_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type Rebuild_Out struct {
	LevelType       BattleType             `json:"level_type"`
	LevelId         int32                  `json:"level_id"`
	ReliveIngot     int32                  `json:"relive_ingot"`
	TotalRound      int16                  `json:"total_round"`
	BuddyRoleId     int8                   `json:"buddy_role_id"`
	MainRolePos     int8                   `json:"main_role_pos"`
	BuddyPos        int8                   `json:"buddy_pos"`
	LastFightingMap int32                  `json:"last_fighting_map"`
	Pass            []Rebuild_Out_Pass     `json:"pass"`
	Smallbox        []Rebuild_Out_Smallbox `json:"smallbox"`
	MengYao         []Rebuild_Out_MengYao  `json:"meng_yao"`
	Shadow          []Rebuild_Out_Shadow   `json:"shadow"`
}

type Rebuild_Out_Pass struct {
	EnemyId int32 `json:"enemy_id"`
}

type Rebuild_Out_Smallbox struct {
	BoxId int32 `json:"box_id"`
}

type Rebuild_Out_MengYao struct {
	MyId int32 `json:"my_id"`
}

type Rebuild_Out_Shadow struct {
	ShadedId int32 `json:"shaded_id"`
}

func (this *Rebuild_Out) Process(session *net.Session) {
	g_OutHandler.Rebuild(session, this)
}

func (this *Rebuild_Out) TypeName() string {
	return "mission.rebuild.out"
}

func (this *Rebuild_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 6
}

func (this *Rebuild_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type EnterExtendLevel_In struct {
	LevelType ExtendLevelType `json:"level_type"`
	LevelId   int32           `json:"level_id"`
}

func (this *EnterExtendLevel_In) Process(session *net.Session) {
	g_InHandler.EnterExtendLevel(session, this)
}

func (this *EnterExtendLevel_In) TypeName() string {
	return "mission.enter_extend_level.in"
}

func (this *EnterExtendLevel_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 7
}

type EnterExtendLevel_Out struct {
	Result OutResult `json:"result"`
}

func (this *EnterExtendLevel_Out) Process(session *net.Session) {
	g_OutHandler.EnterExtendLevel(session, this)
}

func (this *EnterExtendLevel_Out) TypeName() string {
	return "mission.enter_extend_level.out"
}

func (this *EnterExtendLevel_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 7
}

func (this *EnterExtendLevel_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetExtendLevelInfo_In struct {
	LevelType ExtendType `json:"level_type"`
}

func (this *GetExtendLevelInfo_In) Process(session *net.Session) {
	g_InHandler.GetExtendLevelInfo(session, this)
}

func (this *GetExtendLevelInfo_In) TypeName() string {
	return "mission.get_extend_level_info.in"
}

func (this *GetExtendLevelInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 8
}

type GetExtendLevelInfo_Out struct {
	Info []GetExtendLevelInfo_Out_Info `json:"info"`
}

type GetExtendLevelInfo_Out_Info struct {
	LevelType    ExtendLevelType    `json:"level_type"`
	LevelSubType ExtendLevelSubType `json:"level_sub_type"`
	DailyNum     int8               `json:"daily_num"`
	MaxLevel     int16              `json:"max_level"`
	BuyNum       int16              `json:"buy_num"`
}

func (this *GetExtendLevelInfo_Out) Process(session *net.Session) {
	g_OutHandler.GetExtendLevelInfo(session, this)
}

func (this *GetExtendLevelInfo_Out) TypeName() string {
	return "mission.get_extend_level_info.out"
}

func (this *GetExtendLevelInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 8
}

func (this *GetExtendLevelInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type OpenSmallBox_In struct {
	BoxId int32 `json:"box_id"`
}

func (this *OpenSmallBox_In) Process(session *net.Session) {
	g_InHandler.OpenSmallBox(session, this)
}

func (this *OpenSmallBox_In) TypeName() string {
	return "mission.open_small_box.in"
}

func (this *OpenSmallBox_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 9
}

type OpenSmallBox_Out struct {
	Items []OpenSmallBox_Out_Items `json:"items"`
}

type OpenSmallBox_Out_Items struct {
	BoxItemId int32 `json:"box_item_id"`
}

func (this *OpenSmallBox_Out) Process(session *net.Session) {
	g_OutHandler.OpenSmallBox(session, this)
}

func (this *OpenSmallBox_Out) TypeName() string {
	return "mission.open_small_box.out"
}

func (this *OpenSmallBox_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 9
}

func (this *OpenSmallBox_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type NotifyCatchBattlePet_Out struct {
	PetId int32 `json:"petId"`
}

func (this *NotifyCatchBattlePet_Out) Process(session *net.Session) {
	g_OutHandler.NotifyCatchBattlePet(session, this)
}

func (this *NotifyCatchBattlePet_Out) TypeName() string {
	return "mission.notify_catch_battle_pet.out"
}

func (this *NotifyCatchBattlePet_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 10
}

func (this *NotifyCatchBattlePet_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type EnterHardLevel_In struct {
	LevelId int32 `json:"level_id"`
}

func (this *EnterHardLevel_In) Process(session *net.Session) {
	g_InHandler.EnterHardLevel(session, this)
}

func (this *EnterHardLevel_In) TypeName() string {
	return "mission.enter_hard_level.in"
}

func (this *EnterHardLevel_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 11
}

type EnterHardLevel_Out struct {
	Result OutResult `json:"result"`
}

func (this *EnterHardLevel_Out) Process(session *net.Session) {
	g_OutHandler.EnterHardLevel(session, this)
}

func (this *EnterHardLevel_Out) TypeName() string {
	return "mission.enter_hard_level.out"
}

func (this *EnterHardLevel_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 11
}

func (this *EnterHardLevel_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetHardLevel_In struct {
}

func (this *GetHardLevel_In) Process(session *net.Session) {
	g_InHandler.GetHardLevel(session, this)
}

func (this *GetHardLevel_In) TypeName() string {
	return "mission.get_hard_level.in"
}

func (this *GetHardLevel_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 12
}

type GetHardLevel_Out struct {
	Levels []GetHardLevel_Out_Levels `json:"levels"`
}

type GetHardLevel_Out_Levels struct {
	LevelId  int32 `json:"level_id"`
	DailyNum int8  `json:"daily_num"`
	RoundNum int8  `json:"round_num"`
	BuyTimes int16 `json:"buy_times"`
}

func (this *GetHardLevel_Out) Process(session *net.Session) {
	g_OutHandler.GetHardLevel(session, this)
}

func (this *GetHardLevel_Out) TypeName() string {
	return "mission.get_hard_level.out"
}

func (this *GetHardLevel_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 12
}

func (this *GetHardLevel_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetBuddyLevelRoleId_In struct {
}

func (this *GetBuddyLevelRoleId_In) Process(session *net.Session) {
	g_InHandler.GetBuddyLevelRoleId(session, this)
}

func (this *GetBuddyLevelRoleId_In) TypeName() string {
	return "mission.get_buddy_level_role_id.in"
}

func (this *GetBuddyLevelRoleId_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 13
}

type GetBuddyLevelRoleId_Out struct {
	RoleId int8 `json:"role_id"`
}

func (this *GetBuddyLevelRoleId_Out) Process(session *net.Session) {
	g_OutHandler.GetBuddyLevelRoleId(session, this)
}

func (this *GetBuddyLevelRoleId_Out) TypeName() string {
	return "mission.get_buddy_level_role_id.out"
}

func (this *GetBuddyLevelRoleId_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 13
}

func (this *GetBuddyLevelRoleId_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type SetBuddyLevelTeam_In struct {
	RolePos  int8 `json:"role_pos"`
	BuddyPos int8 `json:"buddy_pos"`
	Tactical int8 `json:"tactical"`
}

func (this *SetBuddyLevelTeam_In) Process(session *net.Session) {
	g_InHandler.SetBuddyLevelTeam(session, this)
}

func (this *SetBuddyLevelTeam_In) TypeName() string {
	return "mission.set_buddy_level_team.in"
}

func (this *SetBuddyLevelTeam_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 14
}

type SetBuddyLevelTeam_Out struct {
	Result bool `json:"result"`
}

func (this *SetBuddyLevelTeam_Out) Process(session *net.Session) {
	g_OutHandler.SetBuddyLevelTeam(session, this)
}

func (this *SetBuddyLevelTeam_Out) TypeName() string {
	return "mission.set_buddy_level_team.out"
}

func (this *SetBuddyLevelTeam_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 14
}

func (this *SetBuddyLevelTeam_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type AutoFightLevel_In struct {
	LevelType BattleType `json:"level_type"`
	LevelId   int32      `json:"level_id"`
	Times     int8       `json:"times"`
}

func (this *AutoFightLevel_In) Process(session *net.Session) {
	g_InHandler.AutoFightLevel(session, this)
}

func (this *AutoFightLevel_In) TypeName() string {
	return "mission.auto_fight_level.in"
}

func (this *AutoFightLevel_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 15
}

type AutoFightLevel_Out struct {
	Result []AutoFightLevel_Out_Result `json:"result"`
}

type AutoFightLevel_Out_Result struct {
	LevelBox          []AutoFightLevel_Out_Result_LevelBox          `json:"level_box"`
	BattlePet         []AutoFightLevel_Out_Result_BattlePet         `json:"battle_pet"`
	SmallBox          []AutoFightLevel_Out_Result_SmallBox          `json:"small_box"`
	RandomAwardBox    []AutoFightLevel_Out_Result_RandomAwardBox    `json:"random_award_box"`
	AdditionQuestItem []AutoFightLevel_Out_Result_AdditionQuestItem `json:"addition_quest_item"`
}

type AutoFightLevel_Out_Result_LevelBox struct {
	BoxId int64 `json:"box_id"`
}

type AutoFightLevel_Out_Result_BattlePet struct {
	PetId        int32 `json:"pet_id"`
	Catched      bool  `json:"catched"`
	ConsumeBalls int8  `json:"consume_balls"`
}

type AutoFightLevel_Out_Result_SmallBox struct {
	BoxItemId int32 `json:"box_item_id"`
}

type AutoFightLevel_Out_Result_RandomAwardBox struct {
	BoxId int64 `json:"box_id"`
}

type AutoFightLevel_Out_Result_AdditionQuestItem struct {
	ItemId  int16 `json:"item_id"`
	ItemNum int16 `json:"item_num"`
}

func (this *AutoFightLevel_Out) Process(session *net.Session) {
	g_OutHandler.AutoFightLevel(session, this)
}

func (this *AutoFightLevel_Out) TypeName() string {
	return "mission.auto_fight_level.out"
}

func (this *AutoFightLevel_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 15
}

func (this *AutoFightLevel_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type EnterRainbowLevel_In struct {
	MissionLevelId int32 `json:"mission_level_id"`
}

func (this *EnterRainbowLevel_In) Process(session *net.Session) {
	g_InHandler.EnterRainbowLevel(session, this)
}

func (this *EnterRainbowLevel_In) TypeName() string {
	return "mission.enter_rainbow_level.in"
}

func (this *EnterRainbowLevel_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 16
}

type EnterRainbowLevel_Out struct {
	UsedGhost []EnterRainbowLevel_Out_UsedGhost `json:"used_ghost"`
	CalledPet []EnterRainbowLevel_Out_CalledPet `json:"called_pet"`
}

type EnterRainbowLevel_Out_UsedGhost struct {
	GhostId int16 `json:"ghost_id"`
}

type EnterRainbowLevel_Out_CalledPet struct {
	PetId int32 `json:"pet_id"`
}

func (this *EnterRainbowLevel_Out) Process(session *net.Session) {
	g_OutHandler.EnterRainbowLevel(session, this)
}

func (this *EnterRainbowLevel_Out) TypeName() string {
	return "mission.enter_rainbow_level.out"
}

func (this *EnterRainbowLevel_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 16
}

func (this *EnterRainbowLevel_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type OpenMengYao_In struct {
	MengYaoId int32 `json:"meng_yao_id"`
}

func (this *OpenMengYao_In) Process(session *net.Session) {
	g_InHandler.OpenMengYao(session, this)
}

func (this *OpenMengYao_In) TypeName() string {
	return "mission.open_meng_yao.in"
}

func (this *OpenMengYao_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 17
}

type OpenMengYao_Out struct {
	MengYaoId int32 `json:"meng_yao_id"`
}

func (this *OpenMengYao_Out) Process(session *net.Session) {
	g_OutHandler.OpenMengYao(session, this)
}

func (this *OpenMengYao_Out) TypeName() string {
	return "mission.open_meng_yao.out"
}

func (this *OpenMengYao_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 17
}

func (this *OpenMengYao_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetMissionLevelByItemId_In struct {
	ItemId int16 `json:"item_id"`
}

func (this *GetMissionLevelByItemId_In) Process(session *net.Session) {
	g_InHandler.GetMissionLevelByItemId(session, this)
}

func (this *GetMissionLevelByItemId_In) TypeName() string {
	return "mission.get_mission_level_by_item_id.in"
}

func (this *GetMissionLevelByItemId_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 18
}

type GetMissionLevelByItemId_Out struct {
	Levels []GetMissionLevelByItemId_Out_Levels `json:"levels"`
}

type GetMissionLevelByItemId_Out_Levels struct {
	LevelId  int32 `json:"level_id"`
	RoundNum int8  `json:"round_num"`
	DailyNum int8  `json:"daily_num"`
}

func (this *GetMissionLevelByItemId_Out) Process(session *net.Session) {
	g_OutHandler.GetMissionLevelByItemId(session, this)
}

func (this *GetMissionLevelByItemId_Out) TypeName() string {
	return "mission.get_mission_level_by_item_id.out"
}

func (this *GetMissionLevelByItemId_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 18
}

func (this *GetMissionLevelByItemId_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type BuyHardLevelTimes_In struct {
	LevelId int32 `json:"level_id"`
}

func (this *BuyHardLevelTimes_In) Process(session *net.Session) {
	g_InHandler.BuyHardLevelTimes(session, this)
}

func (this *BuyHardLevelTimes_In) TypeName() string {
	return "mission.buy_hard_level_times.in"
}

func (this *BuyHardLevelTimes_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 19
}

type BuyHardLevelTimes_Out struct {
	Result OutResult `json:"result"`
}

func (this *BuyHardLevelTimes_Out) Process(session *net.Session) {
	g_OutHandler.BuyHardLevelTimes(session, this)
}

func (this *BuyHardLevelTimes_Out) TypeName() string {
	return "mission.buy_hard_level_times.out"
}

func (this *BuyHardLevelTimes_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 19
}

func (this *BuyHardLevelTimes_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type OpenRandomAwardBox_In struct {
	LevelId int32 `json:"level_id"`
}

func (this *OpenRandomAwardBox_In) Process(session *net.Session) {
	g_InHandler.OpenRandomAwardBox(session, this)
}

func (this *OpenRandomAwardBox_In) TypeName() string {
	return "mission.open_random_award_box.in"
}

func (this *OpenRandomAwardBox_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 20
}

type OpenRandomAwardBox_Out struct {
	BoxId int64 `json:"box_id"`
}

func (this *OpenRandomAwardBox_Out) Process(session *net.Session) {
	g_OutHandler.OpenRandomAwardBox(session, this)
}

func (this *OpenRandomAwardBox_Out) TypeName() string {
	return "mission.open_random_award_box.out"
}

func (this *OpenRandomAwardBox_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 20
}

func (this *OpenRandomAwardBox_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type EvaluateLevel_In struct {
}

func (this *EvaluateLevel_In) Process(session *net.Session) {
	g_InHandler.EvaluateLevel(session, this)
}

func (this *EvaluateLevel_In) TypeName() string {
	return "mission.evaluate_level.in"
}

func (this *EvaluateLevel_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 21
}

type EvaluateLevel_Out struct {
	AdditionalQuestItems []EvaluateLevel_Out_AdditionalQuestItems `json:"additional_quest_items"`
}

type EvaluateLevel_Out_AdditionalQuestItems struct {
	ItemId  int16 `json:"item_id"`
	ItemCnt int16 `json:"item_cnt"`
}

func (this *EvaluateLevel_Out) Process(session *net.Session) {
	g_OutHandler.EvaluateLevel(session, this)
}

func (this *EvaluateLevel_Out) TypeName() string {
	return "mission.evaluate_level.out"
}

func (this *EvaluateLevel_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 21
}

func (this *EvaluateLevel_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type OpenShadedBox_In struct {
	ShadedId int32 `json:"shaded_id"`
}

func (this *OpenShadedBox_In) Process(session *net.Session) {
	g_InHandler.OpenShadedBox(session, this)
}

func (this *OpenShadedBox_In) TypeName() string {
	return "mission.open_shaded_box.in"
}

func (this *OpenShadedBox_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 22
}

type OpenShadedBox_Out struct {
}

func (this *OpenShadedBox_Out) Process(session *net.Session) {
	g_OutHandler.OpenShadedBox(session, this)
}

func (this *OpenShadedBox_Out) TypeName() string {
	return "mission.open_shaded_box.out"
}

func (this *OpenShadedBox_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 22
}

func (this *OpenShadedBox_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetMissionTotalStarInfo_In struct {
	TownId int16 `json:"town_id"`
}

func (this *GetMissionTotalStarInfo_In) Process(session *net.Session) {
	g_InHandler.GetMissionTotalStarInfo(session, this)
}

func (this *GetMissionTotalStarInfo_In) TypeName() string {
	return "mission.get_mission_total_star_info.in"
}

func (this *GetMissionTotalStarInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 23
}

type GetMissionTotalStarInfo_Out struct {
	TownId       int16                                      `json:"town_id"`
	TotalStar    int16                                      `json:"total_star"`
	BoxTypeList  []GetMissionTotalStarInfo_Out_BoxTypeList  `json:"box_type_list"`
	Missionstars []GetMissionTotalStarInfo_Out_Missionstars `json:"missionstars"`
}

type GetMissionTotalStarInfo_Out_BoxTypeList struct {
	BoxType int8 `json:"box_type"`
}

type GetMissionTotalStarInfo_Out_Missionstars struct {
	Missionid int16 `json:"missionid"`
	Stars     int16 `json:"stars"`
}

func (this *GetMissionTotalStarInfo_Out) Process(session *net.Session) {
	g_OutHandler.GetMissionTotalStarInfo(session, this)
}

func (this *GetMissionTotalStarInfo_Out) TypeName() string {
	return "mission.get_mission_total_star_info.out"
}

func (this *GetMissionTotalStarInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 23
}

func (this *GetMissionTotalStarInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetMissionTotalStarAwards_In struct {
	TownId  int16 `json:"town_id"`
	BoxType int8  `json:"box_type"`
}

func (this *GetMissionTotalStarAwards_In) Process(session *net.Session) {
	g_InHandler.GetMissionTotalStarAwards(session, this)
}

func (this *GetMissionTotalStarAwards_In) TypeName() string {
	return "mission.get_mission_total_star_awards.in"
}

func (this *GetMissionTotalStarAwards_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 24
}

type GetMissionTotalStarAwards_Out struct {
	TownId  int16 `json:"town_id"`
	BoxType int8  `json:"box_type"`
}

func (this *GetMissionTotalStarAwards_Out) Process(session *net.Session) {
	g_OutHandler.GetMissionTotalStarAwards(session, this)
}

func (this *GetMissionTotalStarAwards_Out) TypeName() string {
	return "mission.get_mission_total_star_awards.out"
}

func (this *GetMissionTotalStarAwards_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 24
}

func (this *GetMissionTotalStarAwards_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ClearMissionLevelState_In struct {
}

func (this *ClearMissionLevelState_In) Process(session *net.Session) {
	g_InHandler.ClearMissionLevelState(session, this)
}

func (this *ClearMissionLevelState_In) TypeName() string {
	return "mission.clear_mission_level_state.in"
}

func (this *ClearMissionLevelState_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 25
}

type ClearMissionLevelState_Out struct {
}

func (this *ClearMissionLevelState_Out) Process(session *net.Session) {
	g_OutHandler.ClearMissionLevelState(session, this)
}

func (this *ClearMissionLevelState_Out) TypeName() string {
	return "mission.clear_mission_level_state.out"
}

func (this *ClearMissionLevelState_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 25
}

func (this *ClearMissionLevelState_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type BuyResourceMissionLevelTimes_In struct {
	SubType int8 `json:"sub_type"`
}

func (this *BuyResourceMissionLevelTimes_In) Process(session *net.Session) {
	g_InHandler.BuyResourceMissionLevelTimes(session, this)
}

func (this *BuyResourceMissionLevelTimes_In) TypeName() string {
	return "mission.buy_resource_mission_level_times.in"
}

func (this *BuyResourceMissionLevelTimes_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 26
}

type BuyResourceMissionLevelTimes_Out struct {
}

func (this *BuyResourceMissionLevelTimes_Out) Process(session *net.Session) {
	g_OutHandler.BuyResourceMissionLevelTimes(session, this)
}

func (this *BuyResourceMissionLevelTimes_Out) TypeName() string {
	return "mission.buy_resource_mission_level_times.out"
}

func (this *BuyResourceMissionLevelTimes_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 26
}

func (this *BuyResourceMissionLevelTimes_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetDragonBall_In struct {
}

func (this *GetDragonBall_In) Process(session *net.Session) {
	g_InHandler.GetDragonBall(session, this)
}

func (this *GetDragonBall_In) TypeName() string {
	return "mission.get_dragon_ball.in"
}

func (this *GetDragonBall_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 27
}

type GetDragonBall_Out struct {
	Ball int16 `json:"ball"`
}

func (this *GetDragonBall_Out) Process(session *net.Session) {
	g_OutHandler.GetDragonBall(session, this)
}

func (this *GetDragonBall_Out) TypeName() string {
	return "mission.get_dragon_ball.out"
}

func (this *GetDragonBall_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 27
}

func (this *GetDragonBall_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetEventItem_In struct {
}

func (this *GetEventItem_In) Process(session *net.Session) {
	g_InHandler.GetEventItem(session, this)
}

func (this *GetEventItem_In) TypeName() string {
	return "mission.get_event_item.in"
}

func (this *GetEventItem_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 28
}

type GetEventItem_Out struct {
	Items []GetEventItem_Out_Items `json:"items"`
}

type GetEventItem_Out_Items struct {
	ItemId  int16 `json:"item_id"`
	ItemNum int16 `json:"item_num"`
}

func (this *GetEventItem_Out) Process(session *net.Session) {
	g_OutHandler.GetEventItem(session, this)
}

func (this *GetEventItem_Out) TypeName() string {
	return "mission.get_event_item.out"
}

func (this *GetEventItem_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 4, 28
}

func (this *GetEventItem_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

func (this *Open_In) Decode(buffer *net.Buffer) {
	this.MissionId = int16(buffer.ReadUint16LE())
}

func (this *Open_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(0)
	buffer.WriteUint16LE(uint16(this.MissionId))
}

func (this *Open_In) ByteSize() int {
	size := 4
	return size
}

func (this *Open_Out) Decode(buffer *net.Buffer) {
	this.Result = OutResult(buffer.ReadUint8())
}

func (this *Open_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(0)
	buffer.WriteUint8(uint8(this.Result))
}

func (this *Open_Out) ByteSize() int {
	size := 3
	return size
}

func (this *GetMissionLevel_In) Decode(buffer *net.Buffer) {
	this.MissionId = int16(buffer.ReadUint16LE())
}

func (this *GetMissionLevel_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(1)
	buffer.WriteUint16LE(uint16(this.MissionId))
}

func (this *GetMissionLevel_In) ByteSize() int {
	size := 4
	return size
}

func (this *GetMissionLevel_Out) Decode(buffer *net.Buffer) {
	this.Result = OutResult(buffer.ReadUint8())
	this.Levels = make([]GetMissionLevel_Out_Levels, buffer.ReadUint8())
	for i := 0; i < len(this.Levels); i++ {
		this.Levels[i].Decode(buffer)
	}
}

func (this *GetMissionLevel_Out_Levels) Decode(buffer *net.Buffer) {
	this.LevelId = int32(buffer.ReadUint32LE())
	this.RoundNum = int8(buffer.ReadUint8())
	this.DailyNum = int8(buffer.ReadUint8())
	this.WaitingShadows = int8(buffer.ReadUint8())
}

func (this *GetMissionLevel_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(1)
	buffer.WriteUint8(uint8(this.Result))
	buffer.WriteUint8(uint8(len(this.Levels)))
	for i := 0; i < len(this.Levels); i++ {
		this.Levels[i].Encode(buffer)
	}
}

func (this *GetMissionLevel_Out_Levels) Encode(buffer *net.Buffer) {
	buffer.WriteUint32LE(uint32(this.LevelId))
	buffer.WriteUint8(uint8(this.RoundNum))
	buffer.WriteUint8(uint8(this.DailyNum))
	buffer.WriteUint8(uint8(this.WaitingShadows))
}

func (this *GetMissionLevel_Out) ByteSize() int {
	size := 4
	size += len(this.Levels) * 7
	return size
}

func (this *EnterLevel_In) Decode(buffer *net.Buffer) {
	this.MissionLevelId = int32(buffer.ReadUint32LE())
}

func (this *EnterLevel_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(2)
	buffer.WriteUint32LE(uint32(this.MissionLevelId))
}

func (this *EnterLevel_In) ByteSize() int {
	size := 6
	return size
}

func (this *EnterLevel_Out) Decode(buffer *net.Buffer) {
	this.Result = OutResult(buffer.ReadUint8())
	this.Smallbox = make([]EnterLevel_Out_Smallbox, buffer.ReadUint8())
	for i := 0; i < len(this.Smallbox); i++ {
		this.Smallbox[i].Decode(buffer)
	}
	this.MengYao = make([]EnterLevel_Out_MengYao, buffer.ReadUint8())
	for i := 0; i < len(this.MengYao); i++ {
		this.MengYao[i].Decode(buffer)
	}
	this.Shadow = make([]EnterLevel_Out_Shadow, buffer.ReadUint8())
	for i := 0; i < len(this.Shadow); i++ {
		this.Shadow[i].Decode(buffer)
	}
}

func (this *EnterLevel_Out_Smallbox) Decode(buffer *net.Buffer) {
	this.BoxId = int32(buffer.ReadUint32LE())
}

func (this *EnterLevel_Out_MengYao) Decode(buffer *net.Buffer) {
	this.MyId = int32(buffer.ReadUint32LE())
}

func (this *EnterLevel_Out_Shadow) Decode(buffer *net.Buffer) {
	this.ShadedId = int32(buffer.ReadUint32LE())
}

func (this *EnterLevel_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(2)
	buffer.WriteUint8(uint8(this.Result))
	buffer.WriteUint8(uint8(len(this.Smallbox)))
	for i := 0; i < len(this.Smallbox); i++ {
		this.Smallbox[i].Encode(buffer)
	}
	buffer.WriteUint8(uint8(len(this.MengYao)))
	for i := 0; i < len(this.MengYao); i++ {
		this.MengYao[i].Encode(buffer)
	}
	buffer.WriteUint8(uint8(len(this.Shadow)))
	for i := 0; i < len(this.Shadow); i++ {
		this.Shadow[i].Encode(buffer)
	}
}

func (this *EnterLevel_Out_Smallbox) Encode(buffer *net.Buffer) {
	buffer.WriteUint32LE(uint32(this.BoxId))
}

func (this *EnterLevel_Out_MengYao) Encode(buffer *net.Buffer) {
	buffer.WriteUint32LE(uint32(this.MyId))
}

func (this *EnterLevel_Out_Shadow) Encode(buffer *net.Buffer) {
	buffer.WriteUint32LE(uint32(this.ShadedId))
}

func (this *EnterLevel_Out) ByteSize() int {
	size := 6
	size += len(this.Smallbox) * 4
	size += len(this.MengYao) * 4
	size += len(this.Shadow) * 4
	return size
}

func (this *OpenLevelBox_In) Decode(buffer *net.Buffer) {
	this.Times = int32(buffer.ReadUint32LE())
}

func (this *OpenLevelBox_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(3)
	buffer.WriteUint32LE(uint32(this.Times))
}

func (this *OpenLevelBox_In) ByteSize() int {
	size := 6
	return size
}

func (this *OpenLevelBox_Out) Decode(buffer *net.Buffer) {
	this.AllBoxId = make([]OpenLevelBox_Out_AllBoxId, buffer.ReadUint8())
	for i := 0; i < len(this.AllBoxId); i++ {
		this.AllBoxId[i].Decode(buffer)
	}
}

func (this *OpenLevelBox_Out_AllBoxId) Decode(buffer *net.Buffer) {
	this.BoxId = int64(buffer.ReadUint64LE())
}

func (this *OpenLevelBox_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(3)
	buffer.WriteUint8(uint8(len(this.AllBoxId)))
	for i := 0; i < len(this.AllBoxId); i++ {
		this.AllBoxId[i].Encode(buffer)
	}
}

func (this *OpenLevelBox_Out_AllBoxId) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.BoxId))
}

func (this *OpenLevelBox_Out) ByteSize() int {
	size := 3
	size += len(this.AllBoxId) * 8
	return size
}

func (this *UseIngotRelive_In) Decode(buffer *net.Buffer) {
}

func (this *UseIngotRelive_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(4)
}

func (this *UseIngotRelive_In) ByteSize() int {
	size := 2
	return size
}

func (this *UseIngotRelive_Out) Decode(buffer *net.Buffer) {
	this.Ingot = int32(buffer.ReadUint32LE())
}

func (this *UseIngotRelive_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(4)
	buffer.WriteUint32LE(uint32(this.Ingot))
}

func (this *UseIngotRelive_Out) ByteSize() int {
	size := 6
	return size
}

func (this *UseItem_In) Decode(buffer *net.Buffer) {
	this.RoleId = int8(buffer.ReadUint8())
	this.ItemId = int16(buffer.ReadUint16LE())
}

func (this *UseItem_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(5)
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint16LE(uint16(this.ItemId))
}

func (this *UseItem_In) ByteSize() int {
	size := 5
	return size
}

func (this *UseItem_Out) Decode(buffer *net.Buffer) {
}

func (this *UseItem_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(5)
}

func (this *UseItem_Out) ByteSize() int {
	size := 2
	return size
}

func (this *Rebuild_Out) Decode(buffer *net.Buffer) {
	this.LevelType = BattleType(buffer.ReadUint8())
	this.LevelId = int32(buffer.ReadUint32LE())
	this.ReliveIngot = int32(buffer.ReadUint32LE())
	this.TotalRound = int16(buffer.ReadUint16LE())
	this.BuddyRoleId = int8(buffer.ReadUint8())
	this.MainRolePos = int8(buffer.ReadUint8())
	this.BuddyPos = int8(buffer.ReadUint8())
	this.LastFightingMap = int32(buffer.ReadUint32LE())
	this.Pass = make([]Rebuild_Out_Pass, buffer.ReadUint8())
	for i := 0; i < len(this.Pass); i++ {
		this.Pass[i].Decode(buffer)
	}
	this.Smallbox = make([]Rebuild_Out_Smallbox, buffer.ReadUint8())
	for i := 0; i < len(this.Smallbox); i++ {
		this.Smallbox[i].Decode(buffer)
	}
	this.MengYao = make([]Rebuild_Out_MengYao, buffer.ReadUint8())
	for i := 0; i < len(this.MengYao); i++ {
		this.MengYao[i].Decode(buffer)
	}
	this.Shadow = make([]Rebuild_Out_Shadow, buffer.ReadUint8())
	for i := 0; i < len(this.Shadow); i++ {
		this.Shadow[i].Decode(buffer)
	}
}

func (this *Rebuild_Out_Pass) Decode(buffer *net.Buffer) {
	this.EnemyId = int32(buffer.ReadUint32LE())
}

func (this *Rebuild_Out_Smallbox) Decode(buffer *net.Buffer) {
	this.BoxId = int32(buffer.ReadUint32LE())
}

func (this *Rebuild_Out_MengYao) Decode(buffer *net.Buffer) {
	this.MyId = int32(buffer.ReadUint32LE())
}

func (this *Rebuild_Out_Shadow) Decode(buffer *net.Buffer) {
	this.ShadedId = int32(buffer.ReadUint32LE())
}

func (this *Rebuild_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(6)
	buffer.WriteUint8(uint8(this.LevelType))
	buffer.WriteUint32LE(uint32(this.LevelId))
	buffer.WriteUint32LE(uint32(this.ReliveIngot))
	buffer.WriteUint16LE(uint16(this.TotalRound))
	buffer.WriteUint8(uint8(this.BuddyRoleId))
	buffer.WriteUint8(uint8(this.MainRolePos))
	buffer.WriteUint8(uint8(this.BuddyPos))
	buffer.WriteUint32LE(uint32(this.LastFightingMap))
	buffer.WriteUint8(uint8(len(this.Pass)))
	for i := 0; i < len(this.Pass); i++ {
		this.Pass[i].Encode(buffer)
	}
	buffer.WriteUint8(uint8(len(this.Smallbox)))
	for i := 0; i < len(this.Smallbox); i++ {
		this.Smallbox[i].Encode(buffer)
	}
	buffer.WriteUint8(uint8(len(this.MengYao)))
	for i := 0; i < len(this.MengYao); i++ {
		this.MengYao[i].Encode(buffer)
	}
	buffer.WriteUint8(uint8(len(this.Shadow)))
	for i := 0; i < len(this.Shadow); i++ {
		this.Shadow[i].Encode(buffer)
	}
}

func (this *Rebuild_Out_Pass) Encode(buffer *net.Buffer) {
	buffer.WriteUint32LE(uint32(this.EnemyId))
}

func (this *Rebuild_Out_Smallbox) Encode(buffer *net.Buffer) {
	buffer.WriteUint32LE(uint32(this.BoxId))
}

func (this *Rebuild_Out_MengYao) Encode(buffer *net.Buffer) {
	buffer.WriteUint32LE(uint32(this.MyId))
}

func (this *Rebuild_Out_Shadow) Encode(buffer *net.Buffer) {
	buffer.WriteUint32LE(uint32(this.ShadedId))
}

func (this *Rebuild_Out) ByteSize() int {
	size := 24
	size += len(this.Pass) * 4
	size += len(this.Smallbox) * 4
	size += len(this.MengYao) * 4
	size += len(this.Shadow) * 4
	return size
}

func (this *EnterExtendLevel_In) Decode(buffer *net.Buffer) {
	this.LevelType = ExtendLevelType(buffer.ReadUint8())
	this.LevelId = int32(buffer.ReadUint32LE())
}

func (this *EnterExtendLevel_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(7)
	buffer.WriteUint8(uint8(this.LevelType))
	buffer.WriteUint32LE(uint32(this.LevelId))
}

func (this *EnterExtendLevel_In) ByteSize() int {
	size := 7
	return size
}

func (this *EnterExtendLevel_Out) Decode(buffer *net.Buffer) {
	this.Result = OutResult(buffer.ReadUint8())
}

func (this *EnterExtendLevel_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(7)
	buffer.WriteUint8(uint8(this.Result))
}

func (this *EnterExtendLevel_Out) ByteSize() int {
	size := 3
	return size
}

func (this *GetExtendLevelInfo_In) Decode(buffer *net.Buffer) {
	this.LevelType = ExtendType(buffer.ReadUint8())
}

func (this *GetExtendLevelInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(8)
	buffer.WriteUint8(uint8(this.LevelType))
}

func (this *GetExtendLevelInfo_In) ByteSize() int {
	size := 3
	return size
}

func (this *GetExtendLevelInfo_Out) Decode(buffer *net.Buffer) {
	this.Info = make([]GetExtendLevelInfo_Out_Info, buffer.ReadUint8())
	for i := 0; i < len(this.Info); i++ {
		this.Info[i].Decode(buffer)
	}
}

func (this *GetExtendLevelInfo_Out_Info) Decode(buffer *net.Buffer) {
	this.LevelType = ExtendLevelType(buffer.ReadUint8())
	this.LevelSubType = ExtendLevelSubType(buffer.ReadUint8())
	this.DailyNum = int8(buffer.ReadUint8())
	this.MaxLevel = int16(buffer.ReadUint16LE())
	this.BuyNum = int16(buffer.ReadUint16LE())
}

func (this *GetExtendLevelInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(8)
	buffer.WriteUint8(uint8(len(this.Info)))
	for i := 0; i < len(this.Info); i++ {
		this.Info[i].Encode(buffer)
	}
}

func (this *GetExtendLevelInfo_Out_Info) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(this.LevelType))
	buffer.WriteUint8(uint8(this.LevelSubType))
	buffer.WriteUint8(uint8(this.DailyNum))
	buffer.WriteUint16LE(uint16(this.MaxLevel))
	buffer.WriteUint16LE(uint16(this.BuyNum))
}

func (this *GetExtendLevelInfo_Out) ByteSize() int {
	size := 3
	size += len(this.Info) * 7
	return size
}

func (this *OpenSmallBox_In) Decode(buffer *net.Buffer) {
	this.BoxId = int32(buffer.ReadUint32LE())
}

func (this *OpenSmallBox_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(9)
	buffer.WriteUint32LE(uint32(this.BoxId))
}

func (this *OpenSmallBox_In) ByteSize() int {
	size := 6
	return size
}

func (this *OpenSmallBox_Out) Decode(buffer *net.Buffer) {
	this.Items = make([]OpenSmallBox_Out_Items, buffer.ReadUint8())
	for i := 0; i < len(this.Items); i++ {
		this.Items[i].Decode(buffer)
	}
}

func (this *OpenSmallBox_Out_Items) Decode(buffer *net.Buffer) {
	this.BoxItemId = int32(buffer.ReadUint32LE())
}

func (this *OpenSmallBox_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(9)
	buffer.WriteUint8(uint8(len(this.Items)))
	for i := 0; i < len(this.Items); i++ {
		this.Items[i].Encode(buffer)
	}
}

func (this *OpenSmallBox_Out_Items) Encode(buffer *net.Buffer) {
	buffer.WriteUint32LE(uint32(this.BoxItemId))
}

func (this *OpenSmallBox_Out) ByteSize() int {
	size := 3
	size += len(this.Items) * 4
	return size
}

func (this *NotifyCatchBattlePet_Out) Decode(buffer *net.Buffer) {
	this.PetId = int32(buffer.ReadUint32LE())
}

func (this *NotifyCatchBattlePet_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(10)
	buffer.WriteUint32LE(uint32(this.PetId))
}

func (this *NotifyCatchBattlePet_Out) ByteSize() int {
	size := 6
	return size
}

func (this *EnterHardLevel_In) Decode(buffer *net.Buffer) {
	this.LevelId = int32(buffer.ReadUint32LE())
}

func (this *EnterHardLevel_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(11)
	buffer.WriteUint32LE(uint32(this.LevelId))
}

func (this *EnterHardLevel_In) ByteSize() int {
	size := 6
	return size
}

func (this *EnterHardLevel_Out) Decode(buffer *net.Buffer) {
	this.Result = OutResult(buffer.ReadUint8())
}

func (this *EnterHardLevel_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(11)
	buffer.WriteUint8(uint8(this.Result))
}

func (this *EnterHardLevel_Out) ByteSize() int {
	size := 3
	return size
}

func (this *GetHardLevel_In) Decode(buffer *net.Buffer) {
}

func (this *GetHardLevel_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(12)
}

func (this *GetHardLevel_In) ByteSize() int {
	size := 2
	return size
}

func (this *GetHardLevel_Out) Decode(buffer *net.Buffer) {
	this.Levels = make([]GetHardLevel_Out_Levels, buffer.ReadUint8())
	for i := 0; i < len(this.Levels); i++ {
		this.Levels[i].Decode(buffer)
	}
}

func (this *GetHardLevel_Out_Levels) Decode(buffer *net.Buffer) {
	this.LevelId = int32(buffer.ReadUint32LE())
	this.DailyNum = int8(buffer.ReadUint8())
	this.RoundNum = int8(buffer.ReadUint8())
	this.BuyTimes = int16(buffer.ReadUint16LE())
}

func (this *GetHardLevel_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(12)
	buffer.WriteUint8(uint8(len(this.Levels)))
	for i := 0; i < len(this.Levels); i++ {
		this.Levels[i].Encode(buffer)
	}
}

func (this *GetHardLevel_Out_Levels) Encode(buffer *net.Buffer) {
	buffer.WriteUint32LE(uint32(this.LevelId))
	buffer.WriteUint8(uint8(this.DailyNum))
	buffer.WriteUint8(uint8(this.RoundNum))
	buffer.WriteUint16LE(uint16(this.BuyTimes))
}

func (this *GetHardLevel_Out) ByteSize() int {
	size := 3
	size += len(this.Levels) * 8
	return size
}

func (this *GetBuddyLevelRoleId_In) Decode(buffer *net.Buffer) {
}

func (this *GetBuddyLevelRoleId_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(13)
}

func (this *GetBuddyLevelRoleId_In) ByteSize() int {
	size := 2
	return size
}

func (this *GetBuddyLevelRoleId_Out) Decode(buffer *net.Buffer) {
	this.RoleId = int8(buffer.ReadUint8())
}

func (this *GetBuddyLevelRoleId_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(13)
	buffer.WriteUint8(uint8(this.RoleId))
}

func (this *GetBuddyLevelRoleId_Out) ByteSize() int {
	size := 3
	return size
}

func (this *SetBuddyLevelTeam_In) Decode(buffer *net.Buffer) {
	this.RolePos = int8(buffer.ReadUint8())
	this.BuddyPos = int8(buffer.ReadUint8())
	this.Tactical = int8(buffer.ReadUint8())
}

func (this *SetBuddyLevelTeam_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(14)
	buffer.WriteUint8(uint8(this.RolePos))
	buffer.WriteUint8(uint8(this.BuddyPos))
	buffer.WriteUint8(uint8(this.Tactical))
}

func (this *SetBuddyLevelTeam_In) ByteSize() int {
	size := 5
	return size
}

func (this *SetBuddyLevelTeam_Out) Decode(buffer *net.Buffer) {
	this.Result = buffer.ReadUint8() == 1
}

func (this *SetBuddyLevelTeam_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(14)
	if this.Result {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *SetBuddyLevelTeam_Out) ByteSize() int {
	size := 3
	return size
}

func (this *AutoFightLevel_In) Decode(buffer *net.Buffer) {
	this.LevelType = BattleType(buffer.ReadUint8())
	this.LevelId = int32(buffer.ReadUint32LE())
	this.Times = int8(buffer.ReadUint8())
}

func (this *AutoFightLevel_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(15)
	buffer.WriteUint8(uint8(this.LevelType))
	buffer.WriteUint32LE(uint32(this.LevelId))
	buffer.WriteUint8(uint8(this.Times))
}

func (this *AutoFightLevel_In) ByteSize() int {
	size := 8
	return size
}

func (this *AutoFightLevel_Out) Decode(buffer *net.Buffer) {
	this.Result = make([]AutoFightLevel_Out_Result, buffer.ReadUint8())
	for i := 0; i < len(this.Result); i++ {
		this.Result[i].Decode(buffer)
	}
}

func (this *AutoFightLevel_Out_Result) Decode(buffer *net.Buffer) {
	this.LevelBox = make([]AutoFightLevel_Out_Result_LevelBox, buffer.ReadUint8())
	for i := 0; i < len(this.LevelBox); i++ {
		this.LevelBox[i].Decode(buffer)
	}
	this.BattlePet = make([]AutoFightLevel_Out_Result_BattlePet, buffer.ReadUint8())
	for i := 0; i < len(this.BattlePet); i++ {
		this.BattlePet[i].Decode(buffer)
	}
	this.SmallBox = make([]AutoFightLevel_Out_Result_SmallBox, buffer.ReadUint8())
	for i := 0; i < len(this.SmallBox); i++ {
		this.SmallBox[i].Decode(buffer)
	}
	this.RandomAwardBox = make([]AutoFightLevel_Out_Result_RandomAwardBox, buffer.ReadUint8())
	for i := 0; i < len(this.RandomAwardBox); i++ {
		this.RandomAwardBox[i].Decode(buffer)
	}
	this.AdditionQuestItem = make([]AutoFightLevel_Out_Result_AdditionQuestItem, buffer.ReadUint8())
	for i := 0; i < len(this.AdditionQuestItem); i++ {
		this.AdditionQuestItem[i].Decode(buffer)
	}
}

func (this *AutoFightLevel_Out_Result_LevelBox) Decode(buffer *net.Buffer) {
	this.BoxId = int64(buffer.ReadUint64LE())
}

func (this *AutoFightLevel_Out_Result_BattlePet) Decode(buffer *net.Buffer) {
	this.PetId = int32(buffer.ReadUint32LE())
	this.Catched = buffer.ReadUint8() == 1
	this.ConsumeBalls = int8(buffer.ReadUint8())
}

func (this *AutoFightLevel_Out_Result_SmallBox) Decode(buffer *net.Buffer) {
	this.BoxItemId = int32(buffer.ReadUint32LE())
}

func (this *AutoFightLevel_Out_Result_RandomAwardBox) Decode(buffer *net.Buffer) {
	this.BoxId = int64(buffer.ReadUint64LE())
}

func (this *AutoFightLevel_Out_Result_AdditionQuestItem) Decode(buffer *net.Buffer) {
	this.ItemId = int16(buffer.ReadUint16LE())
	this.ItemNum = int16(buffer.ReadUint16LE())
}

func (this *AutoFightLevel_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(15)
	buffer.WriteUint8(uint8(len(this.Result)))
	for i := 0; i < len(this.Result); i++ {
		this.Result[i].Encode(buffer)
	}
}

func (this *AutoFightLevel_Out_Result) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(len(this.LevelBox)))
	for i := 0; i < len(this.LevelBox); i++ {
		this.LevelBox[i].Encode(buffer)
	}
	buffer.WriteUint8(uint8(len(this.BattlePet)))
	for i := 0; i < len(this.BattlePet); i++ {
		this.BattlePet[i].Encode(buffer)
	}
	buffer.WriteUint8(uint8(len(this.SmallBox)))
	for i := 0; i < len(this.SmallBox); i++ {
		this.SmallBox[i].Encode(buffer)
	}
	buffer.WriteUint8(uint8(len(this.RandomAwardBox)))
	for i := 0; i < len(this.RandomAwardBox); i++ {
		this.RandomAwardBox[i].Encode(buffer)
	}
	buffer.WriteUint8(uint8(len(this.AdditionQuestItem)))
	for i := 0; i < len(this.AdditionQuestItem); i++ {
		this.AdditionQuestItem[i].Encode(buffer)
	}
}

func (this *AutoFightLevel_Out_Result_LevelBox) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.BoxId))
}

func (this *AutoFightLevel_Out_Result_BattlePet) Encode(buffer *net.Buffer) {
	buffer.WriteUint32LE(uint32(this.PetId))
	if this.Catched {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
	buffer.WriteUint8(uint8(this.ConsumeBalls))
}

func (this *AutoFightLevel_Out_Result_SmallBox) Encode(buffer *net.Buffer) {
	buffer.WriteUint32LE(uint32(this.BoxItemId))
}

func (this *AutoFightLevel_Out_Result_RandomAwardBox) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.BoxId))
}

func (this *AutoFightLevel_Out_Result_AdditionQuestItem) Encode(buffer *net.Buffer) {
	buffer.WriteUint16LE(uint16(this.ItemId))
	buffer.WriteUint16LE(uint16(this.ItemNum))
}

func (this *AutoFightLevel_Out) ByteSize() int {
	size := 3
	for i := 0; i < len(this.Result); i++ {
		size += this.Result[i].ByteSize()
	}
	return size
}

func (this *AutoFightLevel_Out_Result) ByteSize() int {
	size := 5
	size += len(this.LevelBox) * 8
	size += len(this.BattlePet) * 6
	size += len(this.SmallBox) * 4
	size += len(this.RandomAwardBox) * 8
	size += len(this.AdditionQuestItem) * 4
	return size
}

func (this *EnterRainbowLevel_In) Decode(buffer *net.Buffer) {
	this.MissionLevelId = int32(buffer.ReadUint32LE())
}

func (this *EnterRainbowLevel_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(16)
	buffer.WriteUint32LE(uint32(this.MissionLevelId))
}

func (this *EnterRainbowLevel_In) ByteSize() int {
	size := 6
	return size
}

func (this *EnterRainbowLevel_Out) Decode(buffer *net.Buffer) {
	this.UsedGhost = make([]EnterRainbowLevel_Out_UsedGhost, buffer.ReadUint8())
	for i := 0; i < len(this.UsedGhost); i++ {
		this.UsedGhost[i].Decode(buffer)
	}
	this.CalledPet = make([]EnterRainbowLevel_Out_CalledPet, buffer.ReadUint8())
	for i := 0; i < len(this.CalledPet); i++ {
		this.CalledPet[i].Decode(buffer)
	}
}

func (this *EnterRainbowLevel_Out_UsedGhost) Decode(buffer *net.Buffer) {
	this.GhostId = int16(buffer.ReadUint16LE())
}

func (this *EnterRainbowLevel_Out_CalledPet) Decode(buffer *net.Buffer) {
	this.PetId = int32(buffer.ReadUint32LE())
}

func (this *EnterRainbowLevel_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(16)
	buffer.WriteUint8(uint8(len(this.UsedGhost)))
	for i := 0; i < len(this.UsedGhost); i++ {
		this.UsedGhost[i].Encode(buffer)
	}
	buffer.WriteUint8(uint8(len(this.CalledPet)))
	for i := 0; i < len(this.CalledPet); i++ {
		this.CalledPet[i].Encode(buffer)
	}
}

func (this *EnterRainbowLevel_Out_UsedGhost) Encode(buffer *net.Buffer) {
	buffer.WriteUint16LE(uint16(this.GhostId))
}

func (this *EnterRainbowLevel_Out_CalledPet) Encode(buffer *net.Buffer) {
	buffer.WriteUint32LE(uint32(this.PetId))
}

func (this *EnterRainbowLevel_Out) ByteSize() int {
	size := 4
	size += len(this.UsedGhost) * 2
	size += len(this.CalledPet) * 4
	return size
}

func (this *OpenMengYao_In) Decode(buffer *net.Buffer) {
	this.MengYaoId = int32(buffer.ReadUint32LE())
}

func (this *OpenMengYao_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(17)
	buffer.WriteUint32LE(uint32(this.MengYaoId))
}

func (this *OpenMengYao_In) ByteSize() int {
	size := 6
	return size
}

func (this *OpenMengYao_Out) Decode(buffer *net.Buffer) {
	this.MengYaoId = int32(buffer.ReadUint32LE())
}

func (this *OpenMengYao_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(17)
	buffer.WriteUint32LE(uint32(this.MengYaoId))
}

func (this *OpenMengYao_Out) ByteSize() int {
	size := 6
	return size
}

func (this *GetMissionLevelByItemId_In) Decode(buffer *net.Buffer) {
	this.ItemId = int16(buffer.ReadUint16LE())
}

func (this *GetMissionLevelByItemId_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(18)
	buffer.WriteUint16LE(uint16(this.ItemId))
}

func (this *GetMissionLevelByItemId_In) ByteSize() int {
	size := 4
	return size
}

func (this *GetMissionLevelByItemId_Out) Decode(buffer *net.Buffer) {
	this.Levels = make([]GetMissionLevelByItemId_Out_Levels, buffer.ReadUint8())
	for i := 0; i < len(this.Levels); i++ {
		this.Levels[i].Decode(buffer)
	}
}

func (this *GetMissionLevelByItemId_Out_Levels) Decode(buffer *net.Buffer) {
	this.LevelId = int32(buffer.ReadUint32LE())
	this.RoundNum = int8(buffer.ReadUint8())
	this.DailyNum = int8(buffer.ReadUint8())
}

func (this *GetMissionLevelByItemId_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(18)
	buffer.WriteUint8(uint8(len(this.Levels)))
	for i := 0; i < len(this.Levels); i++ {
		this.Levels[i].Encode(buffer)
	}
}

func (this *GetMissionLevelByItemId_Out_Levels) Encode(buffer *net.Buffer) {
	buffer.WriteUint32LE(uint32(this.LevelId))
	buffer.WriteUint8(uint8(this.RoundNum))
	buffer.WriteUint8(uint8(this.DailyNum))
}

func (this *GetMissionLevelByItemId_Out) ByteSize() int {
	size := 3
	size += len(this.Levels) * 6
	return size
}

func (this *BuyHardLevelTimes_In) Decode(buffer *net.Buffer) {
	this.LevelId = int32(buffer.ReadUint32LE())
}

func (this *BuyHardLevelTimes_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(19)
	buffer.WriteUint32LE(uint32(this.LevelId))
}

func (this *BuyHardLevelTimes_In) ByteSize() int {
	size := 6
	return size
}

func (this *BuyHardLevelTimes_Out) Decode(buffer *net.Buffer) {
	this.Result = OutResult(buffer.ReadUint8())
}

func (this *BuyHardLevelTimes_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(19)
	buffer.WriteUint8(uint8(this.Result))
}

func (this *BuyHardLevelTimes_Out) ByteSize() int {
	size := 3
	return size
}

func (this *OpenRandomAwardBox_In) Decode(buffer *net.Buffer) {
	this.LevelId = int32(buffer.ReadUint32LE())
}

func (this *OpenRandomAwardBox_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(20)
	buffer.WriteUint32LE(uint32(this.LevelId))
}

func (this *OpenRandomAwardBox_In) ByteSize() int {
	size := 6
	return size
}

func (this *OpenRandomAwardBox_Out) Decode(buffer *net.Buffer) {
	this.BoxId = int64(buffer.ReadUint64LE())
}

func (this *OpenRandomAwardBox_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(20)
	buffer.WriteUint64LE(uint64(this.BoxId))
}

func (this *OpenRandomAwardBox_Out) ByteSize() int {
	size := 10
	return size
}

func (this *EvaluateLevel_In) Decode(buffer *net.Buffer) {
}

func (this *EvaluateLevel_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(21)
}

func (this *EvaluateLevel_In) ByteSize() int {
	size := 2
	return size
}

func (this *EvaluateLevel_Out) Decode(buffer *net.Buffer) {
	this.AdditionalQuestItems = make([]EvaluateLevel_Out_AdditionalQuestItems, buffer.ReadUint8())
	for i := 0; i < len(this.AdditionalQuestItems); i++ {
		this.AdditionalQuestItems[i].Decode(buffer)
	}
}

func (this *EvaluateLevel_Out_AdditionalQuestItems) Decode(buffer *net.Buffer) {
	this.ItemId = int16(buffer.ReadUint16LE())
	this.ItemCnt = int16(buffer.ReadUint16LE())
}

func (this *EvaluateLevel_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(21)
	buffer.WriteUint8(uint8(len(this.AdditionalQuestItems)))
	for i := 0; i < len(this.AdditionalQuestItems); i++ {
		this.AdditionalQuestItems[i].Encode(buffer)
	}
}

func (this *EvaluateLevel_Out_AdditionalQuestItems) Encode(buffer *net.Buffer) {
	buffer.WriteUint16LE(uint16(this.ItemId))
	buffer.WriteUint16LE(uint16(this.ItemCnt))
}

func (this *EvaluateLevel_Out) ByteSize() int {
	size := 3
	size += len(this.AdditionalQuestItems) * 4
	return size
}

func (this *OpenShadedBox_In) Decode(buffer *net.Buffer) {
	this.ShadedId = int32(buffer.ReadUint32LE())
}

func (this *OpenShadedBox_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(22)
	buffer.WriteUint32LE(uint32(this.ShadedId))
}

func (this *OpenShadedBox_In) ByteSize() int {
	size := 6
	return size
}

func (this *OpenShadedBox_Out) Decode(buffer *net.Buffer) {
}

func (this *OpenShadedBox_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(22)
}

func (this *OpenShadedBox_Out) ByteSize() int {
	size := 2
	return size
}

func (this *GetMissionTotalStarInfo_In) Decode(buffer *net.Buffer) {
	this.TownId = int16(buffer.ReadUint16LE())
}

func (this *GetMissionTotalStarInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(23)
	buffer.WriteUint16LE(uint16(this.TownId))
}

func (this *GetMissionTotalStarInfo_In) ByteSize() int {
	size := 4
	return size
}

func (this *GetMissionTotalStarInfo_Out) Decode(buffer *net.Buffer) {
	this.TownId = int16(buffer.ReadUint16LE())
	this.TotalStar = int16(buffer.ReadUint16LE())
	this.BoxTypeList = make([]GetMissionTotalStarInfo_Out_BoxTypeList, buffer.ReadUint8())
	for i := 0; i < len(this.BoxTypeList); i++ {
		this.BoxTypeList[i].Decode(buffer)
	}
	this.Missionstars = make([]GetMissionTotalStarInfo_Out_Missionstars, buffer.ReadUint8())
	for i := 0; i < len(this.Missionstars); i++ {
		this.Missionstars[i].Decode(buffer)
	}
}

func (this *GetMissionTotalStarInfo_Out_BoxTypeList) Decode(buffer *net.Buffer) {
	this.BoxType = int8(buffer.ReadUint8())
}

func (this *GetMissionTotalStarInfo_Out_Missionstars) Decode(buffer *net.Buffer) {
	this.Missionid = int16(buffer.ReadUint16LE())
	this.Stars = int16(buffer.ReadUint16LE())
}

func (this *GetMissionTotalStarInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(23)
	buffer.WriteUint16LE(uint16(this.TownId))
	buffer.WriteUint16LE(uint16(this.TotalStar))
	buffer.WriteUint8(uint8(len(this.BoxTypeList)))
	for i := 0; i < len(this.BoxTypeList); i++ {
		this.BoxTypeList[i].Encode(buffer)
	}
	buffer.WriteUint8(uint8(len(this.Missionstars)))
	for i := 0; i < len(this.Missionstars); i++ {
		this.Missionstars[i].Encode(buffer)
	}
}

func (this *GetMissionTotalStarInfo_Out_BoxTypeList) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(this.BoxType))
}

func (this *GetMissionTotalStarInfo_Out_Missionstars) Encode(buffer *net.Buffer) {
	buffer.WriteUint16LE(uint16(this.Missionid))
	buffer.WriteUint16LE(uint16(this.Stars))
}

func (this *GetMissionTotalStarInfo_Out) ByteSize() int {
	size := 8
	size += len(this.BoxTypeList) * 1
	size += len(this.Missionstars) * 4
	return size
}

func (this *GetMissionTotalStarAwards_In) Decode(buffer *net.Buffer) {
	this.TownId = int16(buffer.ReadUint16LE())
	this.BoxType = int8(buffer.ReadUint8())
}

func (this *GetMissionTotalStarAwards_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(24)
	buffer.WriteUint16LE(uint16(this.TownId))
	buffer.WriteUint8(uint8(this.BoxType))
}

func (this *GetMissionTotalStarAwards_In) ByteSize() int {
	size := 5
	return size
}

func (this *GetMissionTotalStarAwards_Out) Decode(buffer *net.Buffer) {
	this.TownId = int16(buffer.ReadUint16LE())
	this.BoxType = int8(buffer.ReadUint8())
}

func (this *GetMissionTotalStarAwards_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(24)
	buffer.WriteUint16LE(uint16(this.TownId))
	buffer.WriteUint8(uint8(this.BoxType))
}

func (this *GetMissionTotalStarAwards_Out) ByteSize() int {
	size := 5
	return size
}

func (this *ClearMissionLevelState_In) Decode(buffer *net.Buffer) {
}

func (this *ClearMissionLevelState_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(25)
}

func (this *ClearMissionLevelState_In) ByteSize() int {
	size := 2
	return size
}

func (this *ClearMissionLevelState_Out) Decode(buffer *net.Buffer) {
}

func (this *ClearMissionLevelState_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(25)
}

func (this *ClearMissionLevelState_Out) ByteSize() int {
	size := 2
	return size
}

func (this *BuyResourceMissionLevelTimes_In) Decode(buffer *net.Buffer) {
	this.SubType = int8(buffer.ReadUint8())
}

func (this *BuyResourceMissionLevelTimes_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(26)
	buffer.WriteUint8(uint8(this.SubType))
}

func (this *BuyResourceMissionLevelTimes_In) ByteSize() int {
	size := 3
	return size
}

func (this *BuyResourceMissionLevelTimes_Out) Decode(buffer *net.Buffer) {
}

func (this *BuyResourceMissionLevelTimes_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(26)
}

func (this *BuyResourceMissionLevelTimes_Out) ByteSize() int {
	size := 2
	return size
}

func (this *GetDragonBall_In) Decode(buffer *net.Buffer) {
}

func (this *GetDragonBall_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(27)
}

func (this *GetDragonBall_In) ByteSize() int {
	size := 2
	return size
}

func (this *GetDragonBall_Out) Decode(buffer *net.Buffer) {
	this.Ball = int16(buffer.ReadUint16LE())
}

func (this *GetDragonBall_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(27)
	buffer.WriteUint16LE(uint16(this.Ball))
}

func (this *GetDragonBall_Out) ByteSize() int {
	size := 4
	return size
}

func (this *GetEventItem_In) Decode(buffer *net.Buffer) {
}

func (this *GetEventItem_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(28)
}

func (this *GetEventItem_In) ByteSize() int {
	size := 2
	return size
}

func (this *GetEventItem_Out) Decode(buffer *net.Buffer) {
	this.Items = make([]GetEventItem_Out_Items, buffer.ReadUint8())
	for i := 0; i < len(this.Items); i++ {
		this.Items[i].Decode(buffer)
	}
}

func (this *GetEventItem_Out_Items) Decode(buffer *net.Buffer) {
	this.ItemId = int16(buffer.ReadUint16LE())
	this.ItemNum = int16(buffer.ReadUint16LE())
}

func (this *GetEventItem_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(4)
	buffer.WriteUint8(28)
	buffer.WriteUint8(uint8(len(this.Items)))
	for i := 0; i < len(this.Items); i++ {
		this.Items[i].Encode(buffer)
	}
}

func (this *GetEventItem_Out_Items) Encode(buffer *net.Buffer) {
	buffer.WriteUint16LE(uint16(this.ItemId))
	buffer.WriteUint16LE(uint16(this.ItemNum))
}

func (this *GetEventItem_Out) ByteSize() int {
	size := 3
	size += len(this.Items) * 4
	return size
}
