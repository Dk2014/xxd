package driving_sword_api

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
	CloudMapInfo(*net.Session, *CloudMapInfo_In)
	CloudClimb(*net.Session, *CloudClimb_In)
	CloudTeleport(*net.Session, *CloudTeleport_In)
	AreaTeleport(*net.Session, *AreaTeleport_In)
	AreaMove(*net.Session, *AreaMove_In)
	ExplorerStartBattle(*net.Session, *ExplorerStartBattle_In)
	ExplorerAward(*net.Session, *ExplorerAward_In)
	ExplorerGarrison(*net.Session, *ExplorerGarrison_In)
	VisitMountain(*net.Session, *VisitMountain_In)
	VisiterStartBattle(*net.Session, *VisiterStartBattle_In)
	VisiterAward(*net.Session, *VisiterAward_In)
	MountainTreasureOpen(*net.Session, *MountainTreasureOpen_In)
	ListGarrisons(*net.Session, *ListGarrisons_In)
	AwardGarrison(*net.Session, *AwardGarrison_In)
	EndGarrison(*net.Session, *EndGarrison_In)
	BuyAllowedAction(*net.Session, *BuyAllowedAction_In)
}

type OutHandler interface {
	CloudMapInfo(*net.Session, *CloudMapInfo_Out)
	CloudClimb(*net.Session, *CloudClimb_Out)
	CloudTeleport(*net.Session, *CloudTeleport_Out)
	AreaTeleport(*net.Session, *AreaTeleport_Out)
	AreaMove(*net.Session, *AreaMove_Out)
	ExplorerStartBattle(*net.Session, *ExplorerStartBattle_Out)
	ExplorerAward(*net.Session, *ExplorerAward_Out)
	ExplorerGarrison(*net.Session, *ExplorerGarrison_Out)
	VisitMountain(*net.Session, *VisitMountain_Out)
	VisiterStartBattle(*net.Session, *VisiterStartBattle_Out)
	VisiterAward(*net.Session, *VisiterAward_Out)
	MountainTreasureOpen(*net.Session, *MountainTreasureOpen_Out)
	ListGarrisons(*net.Session, *ListGarrisons_Out)
	AwardGarrison(*net.Session, *AwardGarrison_Out)
	EndGarrison(*net.Session, *EndGarrison_Out)
	BuyAllowedAction(*net.Session, *BuyAllowedAction_Out)
}

func DecodeIn(message []byte) Request {
	var actionId = message[0]
	var buffer = net.NewBuffer(message[1:])
	switch actionId {
	case 0:
		request := new(CloudMapInfo_In)
		request.Decode(buffer)
		return request
	case 1:
		request := new(CloudClimb_In)
		request.Decode(buffer)
		return request
	case 2:
		request := new(CloudTeleport_In)
		request.Decode(buffer)
		return request
	case 3:
		request := new(AreaTeleport_In)
		request.Decode(buffer)
		return request
	case 4:
		request := new(AreaMove_In)
		request.Decode(buffer)
		return request
	case 5:
		request := new(ExplorerStartBattle_In)
		request.Decode(buffer)
		return request
	case 6:
		request := new(ExplorerAward_In)
		request.Decode(buffer)
		return request
	case 7:
		request := new(ExplorerGarrison_In)
		request.Decode(buffer)
		return request
	case 8:
		request := new(VisitMountain_In)
		request.Decode(buffer)
		return request
	case 9:
		request := new(VisiterStartBattle_In)
		request.Decode(buffer)
		return request
	case 10:
		request := new(VisiterAward_In)
		request.Decode(buffer)
		return request
	case 11:
		request := new(MountainTreasureOpen_In)
		request.Decode(buffer)
		return request
	case 12:
		request := new(ListGarrisons_In)
		request.Decode(buffer)
		return request
	case 13:
		request := new(AwardGarrison_In)
		request.Decode(buffer)
		return request
	case 14:
		request := new(EndGarrison_In)
		request.Decode(buffer)
		return request
	case 15:
		request := new(BuyAllowedAction_In)
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
		request := new(CloudMapInfo_Out)
		request.Decode(buffer)
		return request
	case 1:
		request := new(CloudClimb_Out)
		request.Decode(buffer)
		return request
	case 2:
		request := new(CloudTeleport_Out)
		request.Decode(buffer)
		return request
	case 3:
		request := new(AreaTeleport_Out)
		request.Decode(buffer)
		return request
	case 4:
		request := new(AreaMove_Out)
		request.Decode(buffer)
		return request
	case 5:
		request := new(ExplorerStartBattle_Out)
		request.Decode(buffer)
		return request
	case 6:
		request := new(ExplorerAward_Out)
		request.Decode(buffer)
		return request
	case 7:
		request := new(ExplorerGarrison_Out)
		request.Decode(buffer)
		return request
	case 8:
		request := new(VisitMountain_Out)
		request.Decode(buffer)
		return request
	case 9:
		request := new(VisiterStartBattle_Out)
		request.Decode(buffer)
		return request
	case 10:
		request := new(VisiterAward_Out)
		request.Decode(buffer)
		return request
	case 11:
		request := new(MountainTreasureOpen_Out)
		request.Decode(buffer)
		return request
	case 12:
		request := new(ListGarrisons_Out)
		request.Decode(buffer)
		return request
	case 13:
		request := new(AwardGarrison_Out)
		request.Decode(buffer)
		return request
	case 14:
		request := new(EndGarrison_Out)
		request.Decode(buffer)
		return request
	case 15:
		request := new(BuyAllowedAction_Out)
		request.Decode(buffer)
		return request
	}
	_ = buffer
	panic("unsupported response")
}

type ExploringMountainStatus int8

const (
	EXPLORING_MOUNTAIN_STATUS_UNEXPLORED     ExploringMountainStatus = 0
	EXPLORING_MOUNTAIN_STATUS_TREASURE_EMPTY ExploringMountainStatus = 1
	EXPLORING_MOUNTAIN_STATUS_IN_GARRISON    ExploringMountainStatus = 2
	EXPLORING_MOUNTAIN_STATUS_BROKEN         ExploringMountainStatus = 3
)

type VisitingMountainStatus int8

const (
	VISITING_MOUNTAIN_STATUS_UNWIN   VisitingMountainStatus = 0
	VISITING_MOUNTAIN_STATUS_WIN     VisitingMountainStatus = 1
	VISITING_MOUNTAIN_STATUS_AWARDED VisitingMountainStatus = 2
)

type CommonEvent int8

const (
	COMMON_EVENT_HOLE            CommonEvent = 0
	COMMON_EVENT_TELEPORT        CommonEvent = 1
	COMMON_EVENT_OBSTACLE        CommonEvent = 2
	COMMON_EVENT_UNKNOW_TELEPORT CommonEvent = 3
)

type MovingDirection int8

const (
	MOVING_DIRECTION_NORTH MovingDirection = 0
	MOVING_DIRECTION_SOUTH MovingDirection = 1
	MOVING_DIRECTION_WEST  MovingDirection = 2
	MOVING_DIRECTION_EAST  MovingDirection = 3
)

type EventAreas struct {
	Common           []EventAreas_Common           `json:"common"`
	ExploringStatus  []EventAreas_ExploringStatus  `json:"exploring_status"`
	VisitingStatus   []EventAreas_VisitingStatus   `json:"visiting_status"`
	TreasureProgress []EventAreas_TreasureProgress `json:"treasure_progress"`
}

type EventAreas_Common struct {
	X     uint8       `json:"x"`
	Y     uint8       `json:"y"`
	Id    int8        `json:"id"`
	Event CommonEvent `json:"event"`
}

type EventAreas_ExploringStatus struct {
	X            uint8                   `json:"x"`
	Y            uint8                   `json:"y"`
	Id           int8                    `json:"id"`
	Status       ExploringMountainStatus `json:"status"`
	GarrisonTime int64                   `json:"garrison_time"`
}

type EventAreas_VisitingStatus struct {
	X               uint8                  `json:"x"`
	Y               uint8                  `json:"y"`
	Id              int8                   `json:"id"`
	Status          VisitingMountainStatus `json:"status"`
	Pid             int64                  `json:"pid"`
	Nick            []byte                 `json:"nick"`
	RoleId          int8                   `json:"role_id"`
	Level           int16                  `json:"level"`
	FightNum        int32                  `json:"fight_num"`
	FashionId       int16                  `json:"fashion_id"`
	FriendshipLevel int16                  `json:"friendship_level"`
}

type EventAreas_TreasureProgress struct {
	X        uint8 `json:"x"`
	Y        uint8 `json:"y"`
	Id       int8  `json:"id"`
	Progress int8  `json:"progress"`
}

type CloudMap struct {
	Shadows []byte     `json:"shadows"`
	Events  EventAreas `json:"events"`
}

func (this *EventAreas) Decode(buffer *net.Buffer) {
	this.Common = make([]EventAreas_Common, buffer.ReadUint8())
	for i := 0; i < len(this.Common); i++ {
		this.Common[i].Decode(buffer)
	}
	this.ExploringStatus = make([]EventAreas_ExploringStatus, buffer.ReadUint8())
	for i := 0; i < len(this.ExploringStatus); i++ {
		this.ExploringStatus[i].Decode(buffer)
	}
	this.VisitingStatus = make([]EventAreas_VisitingStatus, buffer.ReadUint8())
	for i := 0; i < len(this.VisitingStatus); i++ {
		this.VisitingStatus[i].Decode(buffer)
	}
	this.TreasureProgress = make([]EventAreas_TreasureProgress, buffer.ReadUint8())
	for i := 0; i < len(this.TreasureProgress); i++ {
		this.TreasureProgress[i].Decode(buffer)
	}
}

func (this *EventAreas_Common) Decode(buffer *net.Buffer) {
	this.X = buffer.ReadUint8()
	this.Y = buffer.ReadUint8()
	this.Id = int8(buffer.ReadUint8())
	this.Event = CommonEvent(buffer.ReadUint8())
}

func (this *EventAreas_ExploringStatus) Decode(buffer *net.Buffer) {
	this.X = buffer.ReadUint8()
	this.Y = buffer.ReadUint8()
	this.Id = int8(buffer.ReadUint8())
	this.Status = ExploringMountainStatus(buffer.ReadUint8())
	this.GarrisonTime = int64(buffer.ReadUint64LE())
}

func (this *EventAreas_VisitingStatus) Decode(buffer *net.Buffer) {
	this.X = buffer.ReadUint8()
	this.Y = buffer.ReadUint8()
	this.Id = int8(buffer.ReadUint8())
	this.Status = VisitingMountainStatus(buffer.ReadUint8())
	this.Pid = int64(buffer.ReadUint64LE())
	this.Nick = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.RoleId = int8(buffer.ReadUint8())
	this.Level = int16(buffer.ReadUint16LE())
	this.FightNum = int32(buffer.ReadUint32LE())
	this.FashionId = int16(buffer.ReadUint16LE())
	this.FriendshipLevel = int16(buffer.ReadUint16LE())
}

func (this *EventAreas_TreasureProgress) Decode(buffer *net.Buffer) {
	this.X = buffer.ReadUint8()
	this.Y = buffer.ReadUint8()
	this.Id = int8(buffer.ReadUint8())
	this.Progress = int8(buffer.ReadUint8())
}

func (this *EventAreas) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(len(this.Common)))
	for i := 0; i < len(this.Common); i++ {
		this.Common[i].Encode(buffer)
	}
	buffer.WriteUint8(uint8(len(this.ExploringStatus)))
	for i := 0; i < len(this.ExploringStatus); i++ {
		this.ExploringStatus[i].Encode(buffer)
	}
	buffer.WriteUint8(uint8(len(this.VisitingStatus)))
	for i := 0; i < len(this.VisitingStatus); i++ {
		this.VisitingStatus[i].Encode(buffer)
	}
	buffer.WriteUint8(uint8(len(this.TreasureProgress)))
	for i := 0; i < len(this.TreasureProgress); i++ {
		this.TreasureProgress[i].Encode(buffer)
	}
}

func (this *EventAreas_Common) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(this.X)
	buffer.WriteUint8(this.Y)
	buffer.WriteUint8(uint8(this.Id))
	buffer.WriteUint8(uint8(this.Event))
}

func (this *EventAreas_ExploringStatus) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(this.X)
	buffer.WriteUint8(this.Y)
	buffer.WriteUint8(uint8(this.Id))
	buffer.WriteUint8(uint8(this.Status))
	buffer.WriteUint64LE(uint64(this.GarrisonTime))
}

func (this *EventAreas_VisitingStatus) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(this.X)
	buffer.WriteUint8(this.Y)
	buffer.WriteUint8(uint8(this.Id))
	buffer.WriteUint8(uint8(this.Status))
	buffer.WriteUint64LE(uint64(this.Pid))
	buffer.WriteUint16LE(uint16(len(this.Nick)))
	buffer.WriteBytes(this.Nick)
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint16LE(uint16(this.Level))
	buffer.WriteUint32LE(uint32(this.FightNum))
	buffer.WriteUint16LE(uint16(this.FashionId))
	buffer.WriteUint16LE(uint16(this.FriendshipLevel))
}

func (this *EventAreas_TreasureProgress) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(this.X)
	buffer.WriteUint8(this.Y)
	buffer.WriteUint8(uint8(this.Id))
	buffer.WriteUint8(uint8(this.Progress))
}

func (this *EventAreas) ByteSize() int {
	size := 4
	size += len(this.Common) * 4
	size += len(this.ExploringStatus) * 12
	for i := 0; i < len(this.VisitingStatus); i++ {
		size += this.VisitingStatus[i].ByteSize()
	}
	size += len(this.TreasureProgress) * 4
	return size
}

func (this *EventAreas_VisitingStatus) ByteSize() int {
	size := 25
	size += len(this.Nick)
	return size
}

func (this *CloudMap) Decode(buffer *net.Buffer) {
	this.Shadows = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Events.Decode(buffer)
}

func (this *CloudMap) Encode(buffer *net.Buffer) {
	buffer.WriteUint16LE(uint16(len(this.Shadows)))
	buffer.WriteBytes(this.Shadows)
	this.Events.Encode(buffer)
}

func (this *CloudMap) ByteSize() int {
	size := 2
	size += len(this.Shadows)
	size += this.Events.ByteSize()
	return size
}

type CloudMapInfo_In struct {
}

func (this *CloudMapInfo_In) Process(session *net.Session) {
	g_InHandler.CloudMapInfo(session, this)
}

func (this *CloudMapInfo_In) TypeName() string {
	return "driving_sword.cloud_map_info.in"
}

func (this *CloudMapInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 30, 0
}

type CloudMapInfo_Out struct {
	CurrentCloud      int16    `json:"current_cloud"`
	HighestCloud      int16    `json:"highest_cloud"`
	CurrentX          uint8    `json:"current_x"`
	CurrentY          uint8    `json:"current_y"`
	AllowedAction     int16    `json:"allowed_action"`
	DailyActionBought int8     `json:"daily_action_bought"`
	Map               CloudMap `json:"map"`
}

func (this *CloudMapInfo_Out) Process(session *net.Session) {
	g_OutHandler.CloudMapInfo(session, this)
}

func (this *CloudMapInfo_Out) TypeName() string {
	return "driving_sword.cloud_map_info.out"
}

func (this *CloudMapInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 30, 0
}

func (this *CloudMapInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type CloudClimb_In struct {
}

func (this *CloudClimb_In) Process(session *net.Session) {
	g_InHandler.CloudClimb(session, this)
}

func (this *CloudClimb_In) TypeName() string {
	return "driving_sword.cloud_climb.in"
}

func (this *CloudClimb_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 30, 1
}

type CloudClimb_Out struct {
}

func (this *CloudClimb_Out) Process(session *net.Session) {
	g_OutHandler.CloudClimb(session, this)
}

func (this *CloudClimb_Out) TypeName() string {
	return "driving_sword.cloud_climb.out"
}

func (this *CloudClimb_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 30, 1
}

func (this *CloudClimb_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type CloudTeleport_In struct {
	Cloud int16 `json:"cloud"`
}

func (this *CloudTeleport_In) Process(session *net.Session) {
	g_InHandler.CloudTeleport(session, this)
}

func (this *CloudTeleport_In) TypeName() string {
	return "driving_sword.cloud_teleport.in"
}

func (this *CloudTeleport_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 30, 2
}

type CloudTeleport_Out struct {
	Map CloudMap `json:"map"`
}

func (this *CloudTeleport_Out) Process(session *net.Session) {
	g_OutHandler.CloudTeleport(session, this)
}

func (this *CloudTeleport_Out) TypeName() string {
	return "driving_sword.cloud_teleport.out"
}

func (this *CloudTeleport_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 30, 2
}

func (this *CloudTeleport_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type AreaTeleport_In struct {
}

func (this *AreaTeleport_In) Process(session *net.Session) {
	g_InHandler.AreaTeleport(session, this)
}

func (this *AreaTeleport_In) TypeName() string {
	return "driving_sword.area_teleport.in"
}

func (this *AreaTeleport_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 30, 3
}

type AreaTeleport_Out struct {
	Events EventAreas `json:"events"`
}

func (this *AreaTeleport_Out) Process(session *net.Session) {
	g_OutHandler.AreaTeleport(session, this)
}

func (this *AreaTeleport_Out) TypeName() string {
	return "driving_sword.area_teleport.out"
}

func (this *AreaTeleport_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 30, 3
}

func (this *AreaTeleport_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type AreaMove_In struct {
	Direction MovingDirection `json:"direction"`
}

func (this *AreaMove_In) Process(session *net.Session) {
	g_InHandler.AreaMove(session, this)
}

func (this *AreaMove_In) TypeName() string {
	return "driving_sword.area_move.in"
}

func (this *AreaMove_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 30, 4
}

type AreaMove_Out struct {
	Events EventAreas `json:"events"`
}

func (this *AreaMove_Out) Process(session *net.Session) {
	g_OutHandler.AreaMove(session, this)
}

func (this *AreaMove_Out) TypeName() string {
	return "driving_sword.area_move.out"
}

func (this *AreaMove_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 30, 4
}

func (this *AreaMove_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ExplorerStartBattle_In struct {
}

func (this *ExplorerStartBattle_In) Process(session *net.Session) {
	g_InHandler.ExplorerStartBattle(session, this)
}

func (this *ExplorerStartBattle_In) TypeName() string {
	return "driving_sword.explorer_start_battle.in"
}

func (this *ExplorerStartBattle_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 30, 5
}

type ExplorerStartBattle_Out struct {
}

func (this *ExplorerStartBattle_Out) Process(session *net.Session) {
	g_OutHandler.ExplorerStartBattle(session, this)
}

func (this *ExplorerStartBattle_Out) TypeName() string {
	return "driving_sword.explorer_start_battle.out"
}

func (this *ExplorerStartBattle_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 30, 5
}

func (this *ExplorerStartBattle_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ExplorerAward_In struct {
}

func (this *ExplorerAward_In) Process(session *net.Session) {
	g_InHandler.ExplorerAward(session, this)
}

func (this *ExplorerAward_In) TypeName() string {
	return "driving_sword.explorer_award.in"
}

func (this *ExplorerAward_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 30, 6
}

type ExplorerAward_Out struct {
}

func (this *ExplorerAward_Out) Process(session *net.Session) {
	g_OutHandler.ExplorerAward(session, this)
}

func (this *ExplorerAward_Out) TypeName() string {
	return "driving_sword.explorer_award.out"
}

func (this *ExplorerAward_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 30, 6
}

func (this *ExplorerAward_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ExplorerGarrison_In struct {
	RoleId int8 `json:"role_id"`
}

func (this *ExplorerGarrison_In) Process(session *net.Session) {
	g_InHandler.ExplorerGarrison(session, this)
}

func (this *ExplorerGarrison_In) TypeName() string {
	return "driving_sword.explorer_garrison.in"
}

func (this *ExplorerGarrison_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 30, 7
}

type ExplorerGarrison_Out struct {
}

func (this *ExplorerGarrison_Out) Process(session *net.Session) {
	g_OutHandler.ExplorerGarrison(session, this)
}

func (this *ExplorerGarrison_Out) TypeName() string {
	return "driving_sword.explorer_garrison.out"
}

func (this *ExplorerGarrison_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 30, 7
}

func (this *ExplorerGarrison_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type VisitMountain_In struct {
}

func (this *VisitMountain_In) Process(session *net.Session) {
	g_InHandler.VisitMountain(session, this)
}

func (this *VisitMountain_In) TypeName() string {
	return "driving_sword.visit_mountain.in"
}

func (this *VisitMountain_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 30, 8
}

type VisitMountain_Out struct {
	Status          VisitingMountainStatus `json:"status"`
	Pid             int64                  `json:"pid"`
	Nick            []byte                 `json:"nick"`
	RoleId          int8                   `json:"role_id"`
	Level           int16                  `json:"level"`
	FightNum        int32                  `json:"fight_num"`
	FashionId       int16                  `json:"fashion_id"`
	FriendshipLevel int16                  `json:"friendship_level"`
}

func (this *VisitMountain_Out) Process(session *net.Session) {
	g_OutHandler.VisitMountain(session, this)
}

func (this *VisitMountain_Out) TypeName() string {
	return "driving_sword.visit_mountain.out"
}

func (this *VisitMountain_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 30, 8
}

func (this *VisitMountain_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type VisiterStartBattle_In struct {
}

func (this *VisiterStartBattle_In) Process(session *net.Session) {
	g_InHandler.VisiterStartBattle(session, this)
}

func (this *VisiterStartBattle_In) TypeName() string {
	return "driving_sword.visiter_start_battle.in"
}

func (this *VisiterStartBattle_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 30, 9
}

type VisiterStartBattle_Out struct {
}

func (this *VisiterStartBattle_Out) Process(session *net.Session) {
	g_OutHandler.VisiterStartBattle(session, this)
}

func (this *VisiterStartBattle_Out) TypeName() string {
	return "driving_sword.visiter_start_battle.out"
}

func (this *VisiterStartBattle_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 30, 9
}

func (this *VisiterStartBattle_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type VisiterAward_In struct {
}

func (this *VisiterAward_In) Process(session *net.Session) {
	g_InHandler.VisiterAward(session, this)
}

func (this *VisiterAward_In) TypeName() string {
	return "driving_sword.visiter_award.in"
}

func (this *VisiterAward_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 30, 10
}

type VisiterAward_Out struct {
}

func (this *VisiterAward_Out) Process(session *net.Session) {
	g_OutHandler.VisiterAward(session, this)
}

func (this *VisiterAward_Out) TypeName() string {
	return "driving_sword.visiter_award.out"
}

func (this *VisiterAward_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 30, 10
}

func (this *VisiterAward_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type MountainTreasureOpen_In struct {
}

func (this *MountainTreasureOpen_In) Process(session *net.Session) {
	g_InHandler.MountainTreasureOpen(session, this)
}

func (this *MountainTreasureOpen_In) TypeName() string {
	return "driving_sword.mountain_treasure_open.in"
}

func (this *MountainTreasureOpen_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 30, 11
}

type MountainTreasureOpen_Out struct {
}

func (this *MountainTreasureOpen_Out) Process(session *net.Session) {
	g_OutHandler.MountainTreasureOpen(session, this)
}

func (this *MountainTreasureOpen_Out) TypeName() string {
	return "driving_sword.mountain_treasure_open.out"
}

func (this *MountainTreasureOpen_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 30, 11
}

func (this *MountainTreasureOpen_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ListGarrisons_In struct {
}

func (this *ListGarrisons_In) Process(session *net.Session) {
	g_InHandler.ListGarrisons(session, this)
}

func (this *ListGarrisons_In) TypeName() string {
	return "driving_sword.list_garrisons.in"
}

func (this *ListGarrisons_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 30, 12
}

type ListGarrisons_Out struct {
	Garrisons []ListGarrisons_Out_Garrisons `json:"garrisons"`
}

type ListGarrisons_Out_Garrisons struct {
	RoleId       int8                    `json:"role_id"`
	GarrisonTime int64                   `json:"garrison_time"`
	AwardedTime  int64                   `json:"awarded_time"`
	Cloud        int16                   `json:"cloud"`
	EventId      int8                    `json:"event_id"`
	Status       ExploringMountainStatus `json:"status"`
}

func (this *ListGarrisons_Out) Process(session *net.Session) {
	g_OutHandler.ListGarrisons(session, this)
}

func (this *ListGarrisons_Out) TypeName() string {
	return "driving_sword.list_garrisons.out"
}

func (this *ListGarrisons_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 30, 12
}

func (this *ListGarrisons_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type AwardGarrison_In struct {
	RoleId int8 `json:"role_id"`
}

func (this *AwardGarrison_In) Process(session *net.Session) {
	g_InHandler.AwardGarrison(session, this)
}

func (this *AwardGarrison_In) TypeName() string {
	return "driving_sword.award_garrison.in"
}

func (this *AwardGarrison_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 30, 13
}

type AwardGarrison_Out struct {
}

func (this *AwardGarrison_Out) Process(session *net.Session) {
	g_OutHandler.AwardGarrison(session, this)
}

func (this *AwardGarrison_Out) TypeName() string {
	return "driving_sword.award_garrison.out"
}

func (this *AwardGarrison_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 30, 13
}

func (this *AwardGarrison_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type EndGarrison_In struct {
	RoleId int8 `json:"role_id"`
}

func (this *EndGarrison_In) Process(session *net.Session) {
	g_InHandler.EndGarrison(session, this)
}

func (this *EndGarrison_In) TypeName() string {
	return "driving_sword.end_garrison.in"
}

func (this *EndGarrison_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 30, 14
}

type EndGarrison_Out struct {
	X       uint8                   `json:"x"`
	Y       uint8                   `json:"y"`
	Status  ExploringMountainStatus `json:"status"`
	CloudId int16                   `json:"cloud_id"`
}

func (this *EndGarrison_Out) Process(session *net.Session) {
	g_OutHandler.EndGarrison(session, this)
}

func (this *EndGarrison_Out) TypeName() string {
	return "driving_sword.end_garrison.out"
}

func (this *EndGarrison_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 30, 14
}

func (this *EndGarrison_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type BuyAllowedAction_In struct {
}

func (this *BuyAllowedAction_In) Process(session *net.Session) {
	g_InHandler.BuyAllowedAction(session, this)
}

func (this *BuyAllowedAction_In) TypeName() string {
	return "driving_sword.buy_allowed_action.in"
}

func (this *BuyAllowedAction_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 30, 15
}

type BuyAllowedAction_Out struct {
}

func (this *BuyAllowedAction_Out) Process(session *net.Session) {
	g_OutHandler.BuyAllowedAction(session, this)
}

func (this *BuyAllowedAction_Out) TypeName() string {
	return "driving_sword.buy_allowed_action.out"
}

func (this *BuyAllowedAction_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 30, 15
}

func (this *BuyAllowedAction_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

func (this *CloudMapInfo_In) Decode(buffer *net.Buffer) {
}

func (this *CloudMapInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(30)
	buffer.WriteUint8(0)
}

func (this *CloudMapInfo_In) ByteSize() int {
	size := 2
	return size
}

func (this *CloudMapInfo_Out) Decode(buffer *net.Buffer) {
	this.CurrentCloud = int16(buffer.ReadUint16LE())
	this.HighestCloud = int16(buffer.ReadUint16LE())
	this.CurrentX = buffer.ReadUint8()
	this.CurrentY = buffer.ReadUint8()
	this.AllowedAction = int16(buffer.ReadUint16LE())
	this.DailyActionBought = int8(buffer.ReadUint8())
	this.Map.Decode(buffer)
}

func (this *CloudMapInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(30)
	buffer.WriteUint8(0)
	buffer.WriteUint16LE(uint16(this.CurrentCloud))
	buffer.WriteUint16LE(uint16(this.HighestCloud))
	buffer.WriteUint8(this.CurrentX)
	buffer.WriteUint8(this.CurrentY)
	buffer.WriteUint16LE(uint16(this.AllowedAction))
	buffer.WriteUint8(uint8(this.DailyActionBought))
	this.Map.Encode(buffer)
}

func (this *CloudMapInfo_Out) ByteSize() int {
	size := 11
	size += this.Map.ByteSize()
	return size
}

func (this *CloudClimb_In) Decode(buffer *net.Buffer) {
}

func (this *CloudClimb_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(30)
	buffer.WriteUint8(1)
}

func (this *CloudClimb_In) ByteSize() int {
	size := 2
	return size
}

func (this *CloudClimb_Out) Decode(buffer *net.Buffer) {
}

func (this *CloudClimb_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(30)
	buffer.WriteUint8(1)
}

func (this *CloudClimb_Out) ByteSize() int {
	size := 2
	return size
}

func (this *CloudTeleport_In) Decode(buffer *net.Buffer) {
	this.Cloud = int16(buffer.ReadUint16LE())
}

func (this *CloudTeleport_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(30)
	buffer.WriteUint8(2)
	buffer.WriteUint16LE(uint16(this.Cloud))
}

func (this *CloudTeleport_In) ByteSize() int {
	size := 4
	return size
}

func (this *CloudTeleport_Out) Decode(buffer *net.Buffer) {
	this.Map.Decode(buffer)
}

func (this *CloudTeleport_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(30)
	buffer.WriteUint8(2)
	this.Map.Encode(buffer)
}

func (this *CloudTeleport_Out) ByteSize() int {
	size := 2
	size += this.Map.ByteSize()
	return size
}

func (this *AreaTeleport_In) Decode(buffer *net.Buffer) {
}

func (this *AreaTeleport_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(30)
	buffer.WriteUint8(3)
}

func (this *AreaTeleport_In) ByteSize() int {
	size := 2
	return size
}

func (this *AreaTeleport_Out) Decode(buffer *net.Buffer) {
	this.Events.Decode(buffer)
}

func (this *AreaTeleport_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(30)
	buffer.WriteUint8(3)
	this.Events.Encode(buffer)
}

func (this *AreaTeleport_Out) ByteSize() int {
	size := 2
	size += this.Events.ByteSize()
	return size
}

func (this *AreaMove_In) Decode(buffer *net.Buffer) {
	this.Direction = MovingDirection(buffer.ReadUint8())
}

func (this *AreaMove_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(30)
	buffer.WriteUint8(4)
	buffer.WriteUint8(uint8(this.Direction))
}

func (this *AreaMove_In) ByteSize() int {
	size := 3
	return size
}

func (this *AreaMove_Out) Decode(buffer *net.Buffer) {
	this.Events.Decode(buffer)
}

func (this *AreaMove_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(30)
	buffer.WriteUint8(4)
	this.Events.Encode(buffer)
}

func (this *AreaMove_Out) ByteSize() int {
	size := 2
	size += this.Events.ByteSize()
	return size
}

func (this *ExplorerStartBattle_In) Decode(buffer *net.Buffer) {
}

func (this *ExplorerStartBattle_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(30)
	buffer.WriteUint8(5)
}

func (this *ExplorerStartBattle_In) ByteSize() int {
	size := 2
	return size
}

func (this *ExplorerStartBattle_Out) Decode(buffer *net.Buffer) {
}

func (this *ExplorerStartBattle_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(30)
	buffer.WriteUint8(5)
}

func (this *ExplorerStartBattle_Out) ByteSize() int {
	size := 2
	return size
}

func (this *ExplorerAward_In) Decode(buffer *net.Buffer) {
}

func (this *ExplorerAward_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(30)
	buffer.WriteUint8(6)
}

func (this *ExplorerAward_In) ByteSize() int {
	size := 2
	return size
}

func (this *ExplorerAward_Out) Decode(buffer *net.Buffer) {
}

func (this *ExplorerAward_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(30)
	buffer.WriteUint8(6)
}

func (this *ExplorerAward_Out) ByteSize() int {
	size := 2
	return size
}

func (this *ExplorerGarrison_In) Decode(buffer *net.Buffer) {
	this.RoleId = int8(buffer.ReadUint8())
}

func (this *ExplorerGarrison_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(30)
	buffer.WriteUint8(7)
	buffer.WriteUint8(uint8(this.RoleId))
}

func (this *ExplorerGarrison_In) ByteSize() int {
	size := 3
	return size
}

func (this *ExplorerGarrison_Out) Decode(buffer *net.Buffer) {
}

func (this *ExplorerGarrison_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(30)
	buffer.WriteUint8(7)
}

func (this *ExplorerGarrison_Out) ByteSize() int {
	size := 2
	return size
}

func (this *VisitMountain_In) Decode(buffer *net.Buffer) {
}

func (this *VisitMountain_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(30)
	buffer.WriteUint8(8)
}

func (this *VisitMountain_In) ByteSize() int {
	size := 2
	return size
}

func (this *VisitMountain_Out) Decode(buffer *net.Buffer) {
	this.Status = VisitingMountainStatus(buffer.ReadUint8())
	this.Pid = int64(buffer.ReadUint64LE())
	this.Nick = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.RoleId = int8(buffer.ReadUint8())
	this.Level = int16(buffer.ReadUint16LE())
	this.FightNum = int32(buffer.ReadUint32LE())
	this.FashionId = int16(buffer.ReadUint16LE())
	this.FriendshipLevel = int16(buffer.ReadUint16LE())
}

func (this *VisitMountain_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(30)
	buffer.WriteUint8(8)
	buffer.WriteUint8(uint8(this.Status))
	buffer.WriteUint64LE(uint64(this.Pid))
	buffer.WriteUint16LE(uint16(len(this.Nick)))
	buffer.WriteBytes(this.Nick)
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint16LE(uint16(this.Level))
	buffer.WriteUint32LE(uint32(this.FightNum))
	buffer.WriteUint16LE(uint16(this.FashionId))
	buffer.WriteUint16LE(uint16(this.FriendshipLevel))
}

func (this *VisitMountain_Out) ByteSize() int {
	size := 24
	size += len(this.Nick)
	return size
}

func (this *VisiterStartBattle_In) Decode(buffer *net.Buffer) {
}

func (this *VisiterStartBattle_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(30)
	buffer.WriteUint8(9)
}

func (this *VisiterStartBattle_In) ByteSize() int {
	size := 2
	return size
}

func (this *VisiterStartBattle_Out) Decode(buffer *net.Buffer) {
}

func (this *VisiterStartBattle_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(30)
	buffer.WriteUint8(9)
}

func (this *VisiterStartBattle_Out) ByteSize() int {
	size := 2
	return size
}

func (this *VisiterAward_In) Decode(buffer *net.Buffer) {
}

func (this *VisiterAward_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(30)
	buffer.WriteUint8(10)
}

func (this *VisiterAward_In) ByteSize() int {
	size := 2
	return size
}

func (this *VisiterAward_Out) Decode(buffer *net.Buffer) {
}

func (this *VisiterAward_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(30)
	buffer.WriteUint8(10)
}

func (this *VisiterAward_Out) ByteSize() int {
	size := 2
	return size
}

func (this *MountainTreasureOpen_In) Decode(buffer *net.Buffer) {
}

func (this *MountainTreasureOpen_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(30)
	buffer.WriteUint8(11)
}

func (this *MountainTreasureOpen_In) ByteSize() int {
	size := 2
	return size
}

func (this *MountainTreasureOpen_Out) Decode(buffer *net.Buffer) {
}

func (this *MountainTreasureOpen_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(30)
	buffer.WriteUint8(11)
}

func (this *MountainTreasureOpen_Out) ByteSize() int {
	size := 2
	return size
}

func (this *ListGarrisons_In) Decode(buffer *net.Buffer) {
}

func (this *ListGarrisons_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(30)
	buffer.WriteUint8(12)
}

func (this *ListGarrisons_In) ByteSize() int {
	size := 2
	return size
}

func (this *ListGarrisons_Out) Decode(buffer *net.Buffer) {
	this.Garrisons = make([]ListGarrisons_Out_Garrisons, buffer.ReadUint8())
	for i := 0; i < len(this.Garrisons); i++ {
		this.Garrisons[i].Decode(buffer)
	}
}

func (this *ListGarrisons_Out_Garrisons) Decode(buffer *net.Buffer) {
	this.RoleId = int8(buffer.ReadUint8())
	this.GarrisonTime = int64(buffer.ReadUint64LE())
	this.AwardedTime = int64(buffer.ReadUint64LE())
	this.Cloud = int16(buffer.ReadUint16LE())
	this.EventId = int8(buffer.ReadUint8())
	this.Status = ExploringMountainStatus(buffer.ReadUint8())
}

func (this *ListGarrisons_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(30)
	buffer.WriteUint8(12)
	buffer.WriteUint8(uint8(len(this.Garrisons)))
	for i := 0; i < len(this.Garrisons); i++ {
		this.Garrisons[i].Encode(buffer)
	}
}

func (this *ListGarrisons_Out_Garrisons) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint64LE(uint64(this.GarrisonTime))
	buffer.WriteUint64LE(uint64(this.AwardedTime))
	buffer.WriteUint16LE(uint16(this.Cloud))
	buffer.WriteUint8(uint8(this.EventId))
	buffer.WriteUint8(uint8(this.Status))
}

func (this *ListGarrisons_Out) ByteSize() int {
	size := 3
	size += len(this.Garrisons) * 21
	return size
}

func (this *AwardGarrison_In) Decode(buffer *net.Buffer) {
	this.RoleId = int8(buffer.ReadUint8())
}

func (this *AwardGarrison_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(30)
	buffer.WriteUint8(13)
	buffer.WriteUint8(uint8(this.RoleId))
}

func (this *AwardGarrison_In) ByteSize() int {
	size := 3
	return size
}

func (this *AwardGarrison_Out) Decode(buffer *net.Buffer) {
}

func (this *AwardGarrison_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(30)
	buffer.WriteUint8(13)
}

func (this *AwardGarrison_Out) ByteSize() int {
	size := 2
	return size
}

func (this *EndGarrison_In) Decode(buffer *net.Buffer) {
	this.RoleId = int8(buffer.ReadUint8())
}

func (this *EndGarrison_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(30)
	buffer.WriteUint8(14)
	buffer.WriteUint8(uint8(this.RoleId))
}

func (this *EndGarrison_In) ByteSize() int {
	size := 3
	return size
}

func (this *EndGarrison_Out) Decode(buffer *net.Buffer) {
	this.X = buffer.ReadUint8()
	this.Y = buffer.ReadUint8()
	this.Status = ExploringMountainStatus(buffer.ReadUint8())
	this.CloudId = int16(buffer.ReadUint16LE())
}

func (this *EndGarrison_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(30)
	buffer.WriteUint8(14)
	buffer.WriteUint8(this.X)
	buffer.WriteUint8(this.Y)
	buffer.WriteUint8(uint8(this.Status))
	buffer.WriteUint16LE(uint16(this.CloudId))
}

func (this *EndGarrison_Out) ByteSize() int {
	size := 7
	return size
}

func (this *BuyAllowedAction_In) Decode(buffer *net.Buffer) {
}

func (this *BuyAllowedAction_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(30)
	buffer.WriteUint8(15)
}

func (this *BuyAllowedAction_In) ByteSize() int {
	size := 2
	return size
}

func (this *BuyAllowedAction_Out) Decode(buffer *net.Buffer) {
}

func (this *BuyAllowedAction_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(30)
	buffer.WriteUint8(15)
}

func (this *BuyAllowedAction_Out) ByteSize() int {
	size := 2
	return size
}
