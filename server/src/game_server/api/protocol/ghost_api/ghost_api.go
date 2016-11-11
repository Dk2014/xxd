package ghost_api

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
	Info(*net.Session, *Info_In)
	Equip(*net.Session, *Equip_In)
	Unequip(*net.Session, *Unequip_In)
	Swap(*net.Session, *Swap_In)
	EquipPosChange(*net.Session, *EquipPosChange_In)
	Train(*net.Session, *Train_In)
	Upgrade(*net.Session, *Upgrade_In)
	Baptize(*net.Session, *Baptize_In)
	Compose(*net.Session, *Compose_In)
	TrainSkill(*net.Session, *TrainSkill_In)
	FlushAttr(*net.Session, *FlushAttr_In)
	RelationGhost(*net.Session, *RelationGhost_In)
}

type OutHandler interface {
	Info(*net.Session, *Info_Out)
	Equip(*net.Session, *Equip_Out)
	Unequip(*net.Session, *Unequip_Out)
	Swap(*net.Session, *Swap_Out)
	EquipPosChange(*net.Session, *EquipPosChange_Out)
	Train(*net.Session, *Train_Out)
	Upgrade(*net.Session, *Upgrade_Out)
	Baptize(*net.Session, *Baptize_Out)
	Compose(*net.Session, *Compose_Out)
	TrainSkill(*net.Session, *TrainSkill_Out)
	FlushAttr(*net.Session, *FlushAttr_Out)
	RelationGhost(*net.Session, *RelationGhost_Out)
}

func DecodeIn(message []byte) Request {
	var actionId = message[0]
	var buffer = net.NewBuffer(message[1:])
	switch actionId {
	case 0:
		request := new(Info_In)
		request.Decode(buffer)
		return request
	case 1:
		request := new(Equip_In)
		request.Decode(buffer)
		return request
	case 2:
		request := new(Unequip_In)
		request.Decode(buffer)
		return request
	case 3:
		request := new(Swap_In)
		request.Decode(buffer)
		return request
	case 5:
		request := new(EquipPosChange_In)
		request.Decode(buffer)
		return request
	case 6:
		request := new(Train_In)
		request.Decode(buffer)
		return request
	case 7:
		request := new(Upgrade_In)
		request.Decode(buffer)
		return request
	case 8:
		request := new(Baptize_In)
		request.Decode(buffer)
		return request
	case 11:
		request := new(Compose_In)
		request.Decode(buffer)
		return request
	case 12:
		request := new(TrainSkill_In)
		request.Decode(buffer)
		return request
	case 13:
		request := new(FlushAttr_In)
		request.Decode(buffer)
		return request
	case 14:
		request := new(RelationGhost_In)
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
		request := new(Info_Out)
		request.Decode(buffer)
		return request
	case 1:
		request := new(Equip_Out)
		request.Decode(buffer)
		return request
	case 2:
		request := new(Unequip_Out)
		request.Decode(buffer)
		return request
	case 3:
		request := new(Swap_Out)
		request.Decode(buffer)
		return request
	case 5:
		request := new(EquipPosChange_Out)
		request.Decode(buffer)
		return request
	case 6:
		request := new(Train_Out)
		request.Decode(buffer)
		return request
	case 7:
		request := new(Upgrade_Out)
		request.Decode(buffer)
		return request
	case 8:
		request := new(Baptize_Out)
		request.Decode(buffer)
		return request
	case 11:
		request := new(Compose_Out)
		request.Decode(buffer)
		return request
	case 12:
		request := new(TrainSkill_Out)
		request.Decode(buffer)
		return request
	case 13:
		request := new(FlushAttr_Out)
		request.Decode(buffer)
		return request
	case 14:
		request := new(RelationGhost_Out)
		request.Decode(buffer)
		return request
	}
	_ = buffer
	panic("unsupported response")
}

type EquipPos int8

const (
	EQUIP_POS_POS1 EquipPos = 0
	EQUIP_POS_POS2 EquipPos = 1
	EQUIP_POS_POS3 EquipPos = 2
	EQUIP_POS_POS4 EquipPos = 3
)

type ResultType int8

const (
	RESULT_TYPE_GHOST    ResultType = 0
	RESULT_TYPE_FRAGMENT ResultType = 1
	RESULT_TYPE_FRUIT    ResultType = 2
)

type Info_In struct {
}

func (this *Info_In) Process(session *net.Session) {
	g_InHandler.Info(session, this)
}

func (this *Info_In) TypeName() string {
	return "ghost.info.in"
}

func (this *Info_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 9, 0
}

type Info_Out struct {
	TrainTimes int32                `json:"train_times"`
	FlushTime  int64                `json:"flush_time"`
	Ghosts     []Info_Out_Ghosts    `json:"ghosts"`
	RoleEquip  []Info_Out_RoleEquip `json:"role_equip"`
}

type Info_Out_Ghosts struct {
	Id         int64 `json:"id"`
	GhostId    int16 `json:"ghost_id"`
	Star       int8  `json:"star"`
	Level      int16 `json:"level"`
	SkillLevel int16 `json:"skill_level"`
	Exp        int64 `json:"exp"`
	Pos        int16 `json:"pos"`
	Health     int32 `json:"health"`
	Attack     int32 `json:"attack"`
	Defence    int32 `json:"defence"`
	Speed      int32 `json:"speed"`
	AddGrowth  int16 `json:"add_growth"`
	RelationId int64 `json:"relation_id"`
	Used       bool  `json:"used"`
}

type Info_Out_RoleEquip struct {
	RoleId          int8  `json:"role_id"`
	AlreadyUseGhost bool  `json:"already_use_ghost"`
	GhostPower      int32 `json:"ghost_power"`
	Pos1Id          int64 `json:"pos1_id"`
	Pos2Id          int64 `json:"pos2_id"`
	Pos3Id          int64 `json:"pos3_id"`
	Pos4Id          int64 `json:"pos4_id"`
}

func (this *Info_Out) Process(session *net.Session) {
	g_OutHandler.Info(session, this)
}

func (this *Info_Out) TypeName() string {
	return "ghost.info.out"
}

func (this *Info_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 9, 0
}

func (this *Info_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type Equip_In struct {
	FromId int64    `json:"from_id"`
	RoleId int8     `json:"role_id"`
	ToPos  EquipPos `json:"to_pos"`
}

func (this *Equip_In) Process(session *net.Session) {
	g_InHandler.Equip(session, this)
}

func (this *Equip_In) TypeName() string {
	return "ghost.equip.in"
}

func (this *Equip_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 9, 1
}

type Equip_Out struct {
}

func (this *Equip_Out) Process(session *net.Session) {
	g_OutHandler.Equip(session, this)
}

func (this *Equip_Out) TypeName() string {
	return "ghost.equip.out"
}

func (this *Equip_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 9, 1
}

func (this *Equip_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type Unequip_In struct {
	RoleId int8  `json:"role_id"`
	FromId int64 `json:"from_id"`
}

func (this *Unequip_In) Process(session *net.Session) {
	g_InHandler.Unequip(session, this)
}

func (this *Unequip_In) TypeName() string {
	return "ghost.unequip.in"
}

func (this *Unequip_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 9, 2
}

type Unequip_Out struct {
}

func (this *Unequip_Out) Process(session *net.Session) {
	g_OutHandler.Unequip(session, this)
}

func (this *Unequip_Out) TypeName() string {
	return "ghost.unequip.out"
}

func (this *Unequip_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 9, 2
}

func (this *Unequip_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type Swap_In struct {
	RoleId  int8  `json:"role_id"`
	IdBag   int64 `json:"id_bag"`
	IdEquip int64 `json:"id_equip"`
}

func (this *Swap_In) Process(session *net.Session) {
	g_InHandler.Swap(session, this)
}

func (this *Swap_In) TypeName() string {
	return "ghost.swap.in"
}

func (this *Swap_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 9, 3
}

type Swap_Out struct {
}

func (this *Swap_Out) Process(session *net.Session) {
	g_OutHandler.Swap(session, this)
}

func (this *Swap_Out) TypeName() string {
	return "ghost.swap.out"
}

func (this *Swap_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 9, 3
}

func (this *Swap_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type EquipPosChange_In struct {
	RoleId int8     `json:"role_id"`
	FromId int64    `json:"from_id"`
	ToPos  EquipPos `json:"to_pos"`
}

func (this *EquipPosChange_In) Process(session *net.Session) {
	g_InHandler.EquipPosChange(session, this)
}

func (this *EquipPosChange_In) TypeName() string {
	return "ghost.equip_pos_change.in"
}

func (this *EquipPosChange_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 9, 5
}

type EquipPosChange_Out struct {
}

func (this *EquipPosChange_Out) Process(session *net.Session) {
	g_OutHandler.EquipPosChange(session, this)
}

func (this *EquipPosChange_Out) TypeName() string {
	return "ghost.equip_pos_change.out"
}

func (this *EquipPosChange_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 9, 5
}

func (this *EquipPosChange_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type Train_In struct {
	Id int64 `json:"id"`
}

func (this *Train_In) Process(session *net.Session) {
	g_InHandler.Train(session, this)
}

func (this *Train_In) TypeName() string {
	return "ghost.train.in"
}

func (this *Train_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 9, 6
}

type Train_Out struct {
	AddExp int64 `json:"add_exp"`
}

func (this *Train_Out) Process(session *net.Session) {
	g_OutHandler.Train(session, this)
}

func (this *Train_Out) TypeName() string {
	return "ghost.train.out"
}

func (this *Train_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 9, 6
}

func (this *Train_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type Upgrade_In struct {
	Id int64 `json:"id"`
}

func (this *Upgrade_In) Process(session *net.Session) {
	g_InHandler.Upgrade(session, this)
}

func (this *Upgrade_In) TypeName() string {
	return "ghost.upgrade.in"
}

func (this *Upgrade_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 9, 7
}

type Upgrade_Out struct {
}

func (this *Upgrade_Out) Process(session *net.Session) {
	g_OutHandler.Upgrade(session, this)
}

func (this *Upgrade_Out) TypeName() string {
	return "ghost.upgrade.out"
}

func (this *Upgrade_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 9, 7
}

func (this *Upgrade_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type Baptize_In struct {
	Id int64 `json:"id"`
}

func (this *Baptize_In) Process(session *net.Session) {
	g_InHandler.Baptize(session, this)
}

func (this *Baptize_In) TypeName() string {
	return "ghost.baptize.in"
}

func (this *Baptize_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 9, 8
}

type Baptize_Out struct {
	AddGrowth int8 `json:"add_growth"`
}

func (this *Baptize_Out) Process(session *net.Session) {
	g_OutHandler.Baptize(session, this)
}

func (this *Baptize_Out) TypeName() string {
	return "ghost.baptize.out"
}

func (this *Baptize_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 9, 8
}

func (this *Baptize_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type Compose_In struct {
	GhostId int16 `json:"ghost_id"`
}

func (this *Compose_In) Process(session *net.Session) {
	g_InHandler.Compose(session, this)
}

func (this *Compose_In) TypeName() string {
	return "ghost.compose.in"
}

func (this *Compose_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 9, 11
}

type Compose_Out struct {
	Id      int64 `json:"id"`
	GhostId int16 `json:"ghost_id"`
}

func (this *Compose_Out) Process(session *net.Session) {
	g_OutHandler.Compose(session, this)
}

func (this *Compose_Out) TypeName() string {
	return "ghost.compose.out"
}

func (this *Compose_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 9, 11
}

func (this *Compose_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type TrainSkill_In struct {
	Id int64 `json:"id"`
}

func (this *TrainSkill_In) Process(session *net.Session) {
	g_InHandler.TrainSkill(session, this)
}

func (this *TrainSkill_In) TypeName() string {
	return "ghost.train_skill.in"
}

func (this *TrainSkill_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 9, 12
}

type TrainSkill_Out struct {
}

func (this *TrainSkill_Out) Process(session *net.Session) {
	g_OutHandler.TrainSkill(session, this)
}

func (this *TrainSkill_Out) TypeName() string {
	return "ghost.train_skill.out"
}

func (this *TrainSkill_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 9, 12
}

func (this *TrainSkill_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type FlushAttr_In struct {
	Id int64 `json:"id"`
}

func (this *FlushAttr_In) Process(session *net.Session) {
	g_InHandler.FlushAttr(session, this)
}

func (this *FlushAttr_In) TypeName() string {
	return "ghost.flush_attr.in"
}

func (this *FlushAttr_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 9, 13
}

type FlushAttr_Out struct {
	FlushTime int64 `json:"flush_time"`
}

func (this *FlushAttr_Out) Process(session *net.Session) {
	g_OutHandler.FlushAttr(session, this)
}

func (this *FlushAttr_Out) TypeName() string {
	return "ghost.flush_attr.out"
}

func (this *FlushAttr_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 9, 13
}

func (this *FlushAttr_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type RelationGhost_In struct {
	Master int64 `json:"master"`
	Slave  int64 `json:"slave"`
}

func (this *RelationGhost_In) Process(session *net.Session) {
	g_InHandler.RelationGhost(session, this)
}

func (this *RelationGhost_In) TypeName() string {
	return "ghost.relation_ghost.in"
}

func (this *RelationGhost_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 9, 14
}

type RelationGhost_Out struct {
}

func (this *RelationGhost_Out) Process(session *net.Session) {
	g_OutHandler.RelationGhost(session, this)
}

func (this *RelationGhost_Out) TypeName() string {
	return "ghost.relation_ghost.out"
}

func (this *RelationGhost_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 9, 14
}

func (this *RelationGhost_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

func (this *Info_In) Decode(buffer *net.Buffer) {
}

func (this *Info_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(9)
	buffer.WriteUint8(0)
}

func (this *Info_In) ByteSize() int {
	size := 2
	return size
}

func (this *Info_Out) Decode(buffer *net.Buffer) {
	this.TrainTimes = int32(buffer.ReadUint32LE())
	this.FlushTime = int64(buffer.ReadUint64LE())
	this.Ghosts = make([]Info_Out_Ghosts, buffer.ReadUint8())
	for i := 0; i < len(this.Ghosts); i++ {
		this.Ghosts[i].Decode(buffer)
	}
	this.RoleEquip = make([]Info_Out_RoleEquip, buffer.ReadUint8())
	for i := 0; i < len(this.RoleEquip); i++ {
		this.RoleEquip[i].Decode(buffer)
	}
}

func (this *Info_Out_Ghosts) Decode(buffer *net.Buffer) {
	this.Id = int64(buffer.ReadUint64LE())
	this.GhostId = int16(buffer.ReadUint16LE())
	this.Star = int8(buffer.ReadUint8())
	this.Level = int16(buffer.ReadUint16LE())
	this.SkillLevel = int16(buffer.ReadUint16LE())
	this.Exp = int64(buffer.ReadUint64LE())
	this.Pos = int16(buffer.ReadUint16LE())
	this.Health = int32(buffer.ReadUint32LE())
	this.Attack = int32(buffer.ReadUint32LE())
	this.Defence = int32(buffer.ReadUint32LE())
	this.Speed = int32(buffer.ReadUint32LE())
	this.AddGrowth = int16(buffer.ReadUint16LE())
	this.RelationId = int64(buffer.ReadUint64LE())
	this.Used = buffer.ReadUint8() == 1
}

func (this *Info_Out_RoleEquip) Decode(buffer *net.Buffer) {
	this.RoleId = int8(buffer.ReadUint8())
	this.AlreadyUseGhost = buffer.ReadUint8() == 1
	this.GhostPower = int32(buffer.ReadUint32LE())
	this.Pos1Id = int64(buffer.ReadUint64LE())
	this.Pos2Id = int64(buffer.ReadUint64LE())
	this.Pos3Id = int64(buffer.ReadUint64LE())
	this.Pos4Id = int64(buffer.ReadUint64LE())
}

func (this *Info_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(9)
	buffer.WriteUint8(0)
	buffer.WriteUint32LE(uint32(this.TrainTimes))
	buffer.WriteUint64LE(uint64(this.FlushTime))
	buffer.WriteUint8(uint8(len(this.Ghosts)))
	for i := 0; i < len(this.Ghosts); i++ {
		this.Ghosts[i].Encode(buffer)
	}
	buffer.WriteUint8(uint8(len(this.RoleEquip)))
	for i := 0; i < len(this.RoleEquip); i++ {
		this.RoleEquip[i].Encode(buffer)
	}
}

func (this *Info_Out_Ghosts) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.Id))
	buffer.WriteUint16LE(uint16(this.GhostId))
	buffer.WriteUint8(uint8(this.Star))
	buffer.WriteUint16LE(uint16(this.Level))
	buffer.WriteUint16LE(uint16(this.SkillLevel))
	buffer.WriteUint64LE(uint64(this.Exp))
	buffer.WriteUint16LE(uint16(this.Pos))
	buffer.WriteUint32LE(uint32(this.Health))
	buffer.WriteUint32LE(uint32(this.Attack))
	buffer.WriteUint32LE(uint32(this.Defence))
	buffer.WriteUint32LE(uint32(this.Speed))
	buffer.WriteUint16LE(uint16(this.AddGrowth))
	buffer.WriteUint64LE(uint64(this.RelationId))
	if this.Used {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *Info_Out_RoleEquip) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(this.RoleId))
	if this.AlreadyUseGhost {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
	buffer.WriteUint32LE(uint32(this.GhostPower))
	buffer.WriteUint64LE(uint64(this.Pos1Id))
	buffer.WriteUint64LE(uint64(this.Pos2Id))
	buffer.WriteUint64LE(uint64(this.Pos3Id))
	buffer.WriteUint64LE(uint64(this.Pos4Id))
}

func (this *Info_Out) ByteSize() int {
	size := 16
	size += len(this.Ghosts) * 52
	size += len(this.RoleEquip) * 38
	return size
}

func (this *Equip_In) Decode(buffer *net.Buffer) {
	this.FromId = int64(buffer.ReadUint64LE())
	this.RoleId = int8(buffer.ReadUint8())
	this.ToPos = EquipPos(buffer.ReadUint8())
}

func (this *Equip_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(9)
	buffer.WriteUint8(1)
	buffer.WriteUint64LE(uint64(this.FromId))
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint8(uint8(this.ToPos))
}

func (this *Equip_In) ByteSize() int {
	size := 12
	return size
}

func (this *Equip_Out) Decode(buffer *net.Buffer) {
}

func (this *Equip_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(9)
	buffer.WriteUint8(1)
}

func (this *Equip_Out) ByteSize() int {
	size := 2
	return size
}

func (this *Unequip_In) Decode(buffer *net.Buffer) {
	this.RoleId = int8(buffer.ReadUint8())
	this.FromId = int64(buffer.ReadUint64LE())
}

func (this *Unequip_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(9)
	buffer.WriteUint8(2)
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint64LE(uint64(this.FromId))
}

func (this *Unequip_In) ByteSize() int {
	size := 11
	return size
}

func (this *Unequip_Out) Decode(buffer *net.Buffer) {
}

func (this *Unequip_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(9)
	buffer.WriteUint8(2)
}

func (this *Unequip_Out) ByteSize() int {
	size := 2
	return size
}

func (this *Swap_In) Decode(buffer *net.Buffer) {
	this.RoleId = int8(buffer.ReadUint8())
	this.IdBag = int64(buffer.ReadUint64LE())
	this.IdEquip = int64(buffer.ReadUint64LE())
}

func (this *Swap_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(9)
	buffer.WriteUint8(3)
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint64LE(uint64(this.IdBag))
	buffer.WriteUint64LE(uint64(this.IdEquip))
}

func (this *Swap_In) ByteSize() int {
	size := 19
	return size
}

func (this *Swap_Out) Decode(buffer *net.Buffer) {
}

func (this *Swap_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(9)
	buffer.WriteUint8(3)
}

func (this *Swap_Out) ByteSize() int {
	size := 2
	return size
}

func (this *EquipPosChange_In) Decode(buffer *net.Buffer) {
	this.RoleId = int8(buffer.ReadUint8())
	this.FromId = int64(buffer.ReadUint64LE())
	this.ToPos = EquipPos(buffer.ReadUint8())
}

func (this *EquipPosChange_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(9)
	buffer.WriteUint8(5)
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint64LE(uint64(this.FromId))
	buffer.WriteUint8(uint8(this.ToPos))
}

func (this *EquipPosChange_In) ByteSize() int {
	size := 12
	return size
}

func (this *EquipPosChange_Out) Decode(buffer *net.Buffer) {
}

func (this *EquipPosChange_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(9)
	buffer.WriteUint8(5)
}

func (this *EquipPosChange_Out) ByteSize() int {
	size := 2
	return size
}

func (this *Train_In) Decode(buffer *net.Buffer) {
	this.Id = int64(buffer.ReadUint64LE())
}

func (this *Train_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(9)
	buffer.WriteUint8(6)
	buffer.WriteUint64LE(uint64(this.Id))
}

func (this *Train_In) ByteSize() int {
	size := 10
	return size
}

func (this *Train_Out) Decode(buffer *net.Buffer) {
	this.AddExp = int64(buffer.ReadUint64LE())
}

func (this *Train_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(9)
	buffer.WriteUint8(6)
	buffer.WriteUint64LE(uint64(this.AddExp))
}

func (this *Train_Out) ByteSize() int {
	size := 10
	return size
}

func (this *Upgrade_In) Decode(buffer *net.Buffer) {
	this.Id = int64(buffer.ReadUint64LE())
}

func (this *Upgrade_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(9)
	buffer.WriteUint8(7)
	buffer.WriteUint64LE(uint64(this.Id))
}

func (this *Upgrade_In) ByteSize() int {
	size := 10
	return size
}

func (this *Upgrade_Out) Decode(buffer *net.Buffer) {
}

func (this *Upgrade_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(9)
	buffer.WriteUint8(7)
}

func (this *Upgrade_Out) ByteSize() int {
	size := 2
	return size
}

func (this *Baptize_In) Decode(buffer *net.Buffer) {
	this.Id = int64(buffer.ReadUint64LE())
}

func (this *Baptize_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(9)
	buffer.WriteUint8(8)
	buffer.WriteUint64LE(uint64(this.Id))
}

func (this *Baptize_In) ByteSize() int {
	size := 10
	return size
}

func (this *Baptize_Out) Decode(buffer *net.Buffer) {
	this.AddGrowth = int8(buffer.ReadUint8())
}

func (this *Baptize_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(9)
	buffer.WriteUint8(8)
	buffer.WriteUint8(uint8(this.AddGrowth))
}

func (this *Baptize_Out) ByteSize() int {
	size := 3
	return size
}

func (this *Compose_In) Decode(buffer *net.Buffer) {
	this.GhostId = int16(buffer.ReadUint16LE())
}

func (this *Compose_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(9)
	buffer.WriteUint8(11)
	buffer.WriteUint16LE(uint16(this.GhostId))
}

func (this *Compose_In) ByteSize() int {
	size := 4
	return size
}

func (this *Compose_Out) Decode(buffer *net.Buffer) {
	this.Id = int64(buffer.ReadUint64LE())
	this.GhostId = int16(buffer.ReadUint16LE())
}

func (this *Compose_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(9)
	buffer.WriteUint8(11)
	buffer.WriteUint64LE(uint64(this.Id))
	buffer.WriteUint16LE(uint16(this.GhostId))
}

func (this *Compose_Out) ByteSize() int {
	size := 12
	return size
}

func (this *TrainSkill_In) Decode(buffer *net.Buffer) {
	this.Id = int64(buffer.ReadUint64LE())
}

func (this *TrainSkill_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(9)
	buffer.WriteUint8(12)
	buffer.WriteUint64LE(uint64(this.Id))
}

func (this *TrainSkill_In) ByteSize() int {
	size := 10
	return size
}

func (this *TrainSkill_Out) Decode(buffer *net.Buffer) {
}

func (this *TrainSkill_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(9)
	buffer.WriteUint8(12)
}

func (this *TrainSkill_Out) ByteSize() int {
	size := 2
	return size
}

func (this *FlushAttr_In) Decode(buffer *net.Buffer) {
	this.Id = int64(buffer.ReadUint64LE())
}

func (this *FlushAttr_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(9)
	buffer.WriteUint8(13)
	buffer.WriteUint64LE(uint64(this.Id))
}

func (this *FlushAttr_In) ByteSize() int {
	size := 10
	return size
}

func (this *FlushAttr_Out) Decode(buffer *net.Buffer) {
	this.FlushTime = int64(buffer.ReadUint64LE())
}

func (this *FlushAttr_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(9)
	buffer.WriteUint8(13)
	buffer.WriteUint64LE(uint64(this.FlushTime))
}

func (this *FlushAttr_Out) ByteSize() int {
	size := 10
	return size
}

func (this *RelationGhost_In) Decode(buffer *net.Buffer) {
	this.Master = int64(buffer.ReadUint64LE())
	this.Slave = int64(buffer.ReadUint64LE())
}

func (this *RelationGhost_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(9)
	buffer.WriteUint8(14)
	buffer.WriteUint64LE(uint64(this.Master))
	buffer.WriteUint64LE(uint64(this.Slave))
}

func (this *RelationGhost_In) ByteSize() int {
	size := 18
	return size
}

func (this *RelationGhost_Out) Decode(buffer *net.Buffer) {
}

func (this *RelationGhost_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(9)
	buffer.WriteUint8(14)
}

func (this *RelationGhost_Out) ByteSize() int {
	size := 2
	return size
}
