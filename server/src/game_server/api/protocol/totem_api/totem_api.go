package totem_api

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
	CallTotem(*net.Session, *CallTotem_In)
	Upgrade(*net.Session, *Upgrade_In)
	Decompose(*net.Session, *Decompose_In)
	Equip(*net.Session, *Equip_In)
	Unequip(*net.Session, *Unequip_In)
	EquipPosChange(*net.Session, *EquipPosChange_In)
	Swap(*net.Session, *Swap_In)
	IsBagFull(*net.Session, *IsBagFull_In)
}

type OutHandler interface {
	Info(*net.Session, *Info_Out)
	CallTotem(*net.Session, *CallTotem_Out)
	Upgrade(*net.Session, *Upgrade_Out)
	Decompose(*net.Session, *Decompose_Out)
	Equip(*net.Session, *Equip_Out)
	Unequip(*net.Session, *Unequip_Out)
	EquipPosChange(*net.Session, *EquipPosChange_Out)
	Swap(*net.Session, *Swap_Out)
	IsBagFull(*net.Session, *IsBagFull_Out)
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
		request := new(CallTotem_In)
		request.Decode(buffer)
		return request
	case 2:
		request := new(Upgrade_In)
		request.Decode(buffer)
		return request
	case 3:
		request := new(Decompose_In)
		request.Decode(buffer)
		return request
	case 4:
		request := new(Equip_In)
		request.Decode(buffer)
		return request
	case 5:
		request := new(Unequip_In)
		request.Decode(buffer)
		return request
	case 6:
		request := new(EquipPosChange_In)
		request.Decode(buffer)
		return request
	case 8:
		request := new(Swap_In)
		request.Decode(buffer)
		return request
	case 9:
		request := new(IsBagFull_In)
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
		request := new(CallTotem_Out)
		request.Decode(buffer)
		return request
	case 2:
		request := new(Upgrade_Out)
		request.Decode(buffer)
		return request
	case 3:
		request := new(Decompose_Out)
		request.Decode(buffer)
		return request
	case 4:
		request := new(Equip_Out)
		request.Decode(buffer)
		return request
	case 5:
		request := new(Unequip_Out)
		request.Decode(buffer)
		return request
	case 6:
		request := new(EquipPosChange_Out)
		request.Decode(buffer)
		return request
	case 8:
		request := new(Swap_Out)
		request.Decode(buffer)
		return request
	case 9:
		request := new(IsBagFull_Out)
		request.Decode(buffer)
		return request
	}
	_ = buffer
	panic("unsupported response")
}

type CallType int8

const (
	CALL_TYPE_BONE   CallType = 0
	CALL_TYPE_HALLOW CallType = 1
	CALL_TYPE_INGOT  CallType = 2
)

type EquipPos int8

const (
	EQUIP_POS_POS1 EquipPos = 0
	EQUIP_POS_POS2 EquipPos = 1
	EQUIP_POS_POS3 EquipPos = 2
	EQUIP_POS_POS4 EquipPos = 3
	EQUIP_POS_POS5 EquipPos = 4
)

type Info_In struct {
}

func (this *Info_In) Process(session *net.Session) {
	g_InHandler.Info(session, this)
}

func (this *Info_In) TypeName() string {
	return "totem.info.in"
}

func (this *Info_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 31, 0
}

type Info_Out struct {
	Totem       []Info_Out_Totem `json:"totem"`
	RockRuneNum int32            `json:"rock_rune_num"`
	JadeRuneNum int32            `json:"jade_rune_num"`
	CallNum     int16            `json:"call_num"`
	Pos1Id      int64            `json:"pos1_id"`
	Pos2Id      int64            `json:"pos2_id"`
	Pos3Id      int64            `json:"pos3_id"`
	Pos4Id      int64            `json:"pos4_id"`
	Pos5Id      int64            `json:"pos5_id"`
}

type Info_Out_Totem struct {
	Id      int64 `json:"id"`
	TotemId int16 `json:"totem_id"`
	Skill   int16 `json:"skill"`
	Level   int8  `json:"level"`
}

func (this *Info_Out) Process(session *net.Session) {
	g_OutHandler.Info(session, this)
}

func (this *Info_Out) TypeName() string {
	return "totem.info.out"
}

func (this *Info_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 31, 0
}

func (this *Info_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type CallTotem_In struct {
	CallType CallType `json:"call_type"`
}

func (this *CallTotem_In) Process(session *net.Session) {
	g_InHandler.CallTotem(session, this)
}

func (this *CallTotem_In) TypeName() string {
	return "totem.call_totem.in"
}

func (this *CallTotem_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 31, 1
}

type CallTotem_Out struct {
}

func (this *CallTotem_Out) Process(session *net.Session) {
	g_OutHandler.CallTotem(session, this)
}

func (this *CallTotem_Out) TypeName() string {
	return "totem.call_totem.out"
}

func (this *CallTotem_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 31, 1
}

func (this *CallTotem_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type Upgrade_In struct {
	TargetId int64 `json:"target_id"`
}

func (this *Upgrade_In) Process(session *net.Session) {
	g_InHandler.Upgrade(session, this)
}

func (this *Upgrade_In) TypeName() string {
	return "totem.upgrade.in"
}

func (this *Upgrade_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 31, 2
}

type Upgrade_Out struct {
	Ok    bool  `json:"ok"`
	Id    int64 `json:"id"`
	Skill int16 `json:"skill"`
	Level int8  `json:"level"`
}

func (this *Upgrade_Out) Process(session *net.Session) {
	g_OutHandler.Upgrade(session, this)
}

func (this *Upgrade_Out) TypeName() string {
	return "totem.upgrade.out"
}

func (this *Upgrade_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 31, 2
}

func (this *Upgrade_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type Decompose_In struct {
	TotemId int64 `json:"totem_id"`
}

func (this *Decompose_In) Process(session *net.Session) {
	g_InHandler.Decompose(session, this)
}

func (this *Decompose_In) TypeName() string {
	return "totem.decompose.in"
}

func (this *Decompose_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 31, 3
}

type Decompose_Out struct {
}

func (this *Decompose_Out) Process(session *net.Session) {
	g_OutHandler.Decompose(session, this)
}

func (this *Decompose_Out) TypeName() string {
	return "totem.decompose.out"
}

func (this *Decompose_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 31, 3
}

func (this *Decompose_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type Equip_In struct {
	TotemId  int64    `json:"totem_id"`
	EquipPos EquipPos `json:"equip_pos"`
}

func (this *Equip_In) Process(session *net.Session) {
	g_InHandler.Equip(session, this)
}

func (this *Equip_In) TypeName() string {
	return "totem.equip.in"
}

func (this *Equip_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 31, 4
}

type Equip_Out struct {
}

func (this *Equip_Out) Process(session *net.Session) {
	g_OutHandler.Equip(session, this)
}

func (this *Equip_Out) TypeName() string {
	return "totem.equip.out"
}

func (this *Equip_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 31, 4
}

func (this *Equip_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type Unequip_In struct {
	EquipPos EquipPos `json:"equip_pos"`
}

func (this *Unequip_In) Process(session *net.Session) {
	g_InHandler.Unequip(session, this)
}

func (this *Unequip_In) TypeName() string {
	return "totem.unequip.in"
}

func (this *Unequip_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 31, 5
}

type Unequip_Out struct {
}

func (this *Unequip_Out) Process(session *net.Session) {
	g_OutHandler.Unequip(session, this)
}

func (this *Unequip_Out) TypeName() string {
	return "totem.unequip.out"
}

func (this *Unequip_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 31, 5
}

func (this *Unequip_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type EquipPosChange_In struct {
	FromPos EquipPos `json:"from_pos"`
	ToPos   EquipPos `json:"to_pos"`
}

func (this *EquipPosChange_In) Process(session *net.Session) {
	g_InHandler.EquipPosChange(session, this)
}

func (this *EquipPosChange_In) TypeName() string {
	return "totem.equip_pos_change.in"
}

func (this *EquipPosChange_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 31, 6
}

type EquipPosChange_Out struct {
}

func (this *EquipPosChange_Out) Process(session *net.Session) {
	g_OutHandler.EquipPosChange(session, this)
}

func (this *EquipPosChange_Out) TypeName() string {
	return "totem.equip_pos_change.out"
}

func (this *EquipPosChange_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 31, 6
}

func (this *EquipPosChange_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type Swap_In struct {
	EquipedId int64 `json:"equiped_id"`
	InbagId   int64 `json:"inbag_id"`
}

func (this *Swap_In) Process(session *net.Session) {
	g_InHandler.Swap(session, this)
}

func (this *Swap_In) TypeName() string {
	return "totem.swap.in"
}

func (this *Swap_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 31, 8
}

type Swap_Out struct {
}

func (this *Swap_Out) Process(session *net.Session) {
	g_OutHandler.Swap(session, this)
}

func (this *Swap_Out) TypeName() string {
	return "totem.swap.out"
}

func (this *Swap_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 31, 8
}

func (this *Swap_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type IsBagFull_In struct {
}

func (this *IsBagFull_In) Process(session *net.Session) {
	g_InHandler.IsBagFull(session, this)
}

func (this *IsBagFull_In) TypeName() string {
	return "totem.is_bag_full.in"
}

func (this *IsBagFull_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 31, 9
}

type IsBagFull_Out struct {
	Full bool `json:"full"`
}

func (this *IsBagFull_Out) Process(session *net.Session) {
	g_OutHandler.IsBagFull(session, this)
}

func (this *IsBagFull_Out) TypeName() string {
	return "totem.is_bag_full.out"
}

func (this *IsBagFull_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 31, 9
}

func (this *IsBagFull_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

func (this *Info_In) Decode(buffer *net.Buffer) {
}

func (this *Info_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(31)
	buffer.WriteUint8(0)
}

func (this *Info_In) ByteSize() int {
	size := 2
	return size
}

func (this *Info_Out) Decode(buffer *net.Buffer) {
	this.Totem = make([]Info_Out_Totem, buffer.ReadUint8())
	for i := 0; i < len(this.Totem); i++ {
		this.Totem[i].Decode(buffer)
	}
	this.RockRuneNum = int32(buffer.ReadUint32LE())
	this.JadeRuneNum = int32(buffer.ReadUint32LE())
	this.CallNum = int16(buffer.ReadUint16LE())
	this.Pos1Id = int64(buffer.ReadUint64LE())
	this.Pos2Id = int64(buffer.ReadUint64LE())
	this.Pos3Id = int64(buffer.ReadUint64LE())
	this.Pos4Id = int64(buffer.ReadUint64LE())
	this.Pos5Id = int64(buffer.ReadUint64LE())
}

func (this *Info_Out_Totem) Decode(buffer *net.Buffer) {
	this.Id = int64(buffer.ReadUint64LE())
	this.TotemId = int16(buffer.ReadUint16LE())
	this.Skill = int16(buffer.ReadUint16LE())
	this.Level = int8(buffer.ReadUint8())
}

func (this *Info_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(31)
	buffer.WriteUint8(0)
	buffer.WriteUint8(uint8(len(this.Totem)))
	for i := 0; i < len(this.Totem); i++ {
		this.Totem[i].Encode(buffer)
	}
	buffer.WriteUint32LE(uint32(this.RockRuneNum))
	buffer.WriteUint32LE(uint32(this.JadeRuneNum))
	buffer.WriteUint16LE(uint16(this.CallNum))
	buffer.WriteUint64LE(uint64(this.Pos1Id))
	buffer.WriteUint64LE(uint64(this.Pos2Id))
	buffer.WriteUint64LE(uint64(this.Pos3Id))
	buffer.WriteUint64LE(uint64(this.Pos4Id))
	buffer.WriteUint64LE(uint64(this.Pos5Id))
}

func (this *Info_Out_Totem) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.Id))
	buffer.WriteUint16LE(uint16(this.TotemId))
	buffer.WriteUint16LE(uint16(this.Skill))
	buffer.WriteUint8(uint8(this.Level))
}

func (this *Info_Out) ByteSize() int {
	size := 53
	size += len(this.Totem) * 13
	return size
}

func (this *CallTotem_In) Decode(buffer *net.Buffer) {
	this.CallType = CallType(buffer.ReadUint8())
}

func (this *CallTotem_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(31)
	buffer.WriteUint8(1)
	buffer.WriteUint8(uint8(this.CallType))
}

func (this *CallTotem_In) ByteSize() int {
	size := 3
	return size
}

func (this *CallTotem_Out) Decode(buffer *net.Buffer) {
}

func (this *CallTotem_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(31)
	buffer.WriteUint8(1)
}

func (this *CallTotem_Out) ByteSize() int {
	size := 2
	return size
}

func (this *Upgrade_In) Decode(buffer *net.Buffer) {
	this.TargetId = int64(buffer.ReadUint64LE())
}

func (this *Upgrade_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(31)
	buffer.WriteUint8(2)
	buffer.WriteUint64LE(uint64(this.TargetId))
}

func (this *Upgrade_In) ByteSize() int {
	size := 10
	return size
}

func (this *Upgrade_Out) Decode(buffer *net.Buffer) {
	this.Ok = buffer.ReadUint8() == 1
	this.Id = int64(buffer.ReadUint64LE())
	this.Skill = int16(buffer.ReadUint16LE())
	this.Level = int8(buffer.ReadUint8())
}

func (this *Upgrade_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(31)
	buffer.WriteUint8(2)
	if this.Ok {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
	buffer.WriteUint64LE(uint64(this.Id))
	buffer.WriteUint16LE(uint16(this.Skill))
	buffer.WriteUint8(uint8(this.Level))
}

func (this *Upgrade_Out) ByteSize() int {
	size := 14
	return size
}

func (this *Decompose_In) Decode(buffer *net.Buffer) {
	this.TotemId = int64(buffer.ReadUint64LE())
}

func (this *Decompose_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(31)
	buffer.WriteUint8(3)
	buffer.WriteUint64LE(uint64(this.TotemId))
}

func (this *Decompose_In) ByteSize() int {
	size := 10
	return size
}

func (this *Decompose_Out) Decode(buffer *net.Buffer) {
}

func (this *Decompose_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(31)
	buffer.WriteUint8(3)
}

func (this *Decompose_Out) ByteSize() int {
	size := 2
	return size
}

func (this *Equip_In) Decode(buffer *net.Buffer) {
	this.TotemId = int64(buffer.ReadUint64LE())
	this.EquipPos = EquipPos(buffer.ReadUint8())
}

func (this *Equip_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(31)
	buffer.WriteUint8(4)
	buffer.WriteUint64LE(uint64(this.TotemId))
	buffer.WriteUint8(uint8(this.EquipPos))
}

func (this *Equip_In) ByteSize() int {
	size := 11
	return size
}

func (this *Equip_Out) Decode(buffer *net.Buffer) {
}

func (this *Equip_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(31)
	buffer.WriteUint8(4)
}

func (this *Equip_Out) ByteSize() int {
	size := 2
	return size
}

func (this *Unequip_In) Decode(buffer *net.Buffer) {
	this.EquipPos = EquipPos(buffer.ReadUint8())
}

func (this *Unequip_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(31)
	buffer.WriteUint8(5)
	buffer.WriteUint8(uint8(this.EquipPos))
}

func (this *Unequip_In) ByteSize() int {
	size := 3
	return size
}

func (this *Unequip_Out) Decode(buffer *net.Buffer) {
}

func (this *Unequip_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(31)
	buffer.WriteUint8(5)
}

func (this *Unequip_Out) ByteSize() int {
	size := 2
	return size
}

func (this *EquipPosChange_In) Decode(buffer *net.Buffer) {
	this.FromPos = EquipPos(buffer.ReadUint8())
	this.ToPos = EquipPos(buffer.ReadUint8())
}

func (this *EquipPosChange_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(31)
	buffer.WriteUint8(6)
	buffer.WriteUint8(uint8(this.FromPos))
	buffer.WriteUint8(uint8(this.ToPos))
}

func (this *EquipPosChange_In) ByteSize() int {
	size := 4
	return size
}

func (this *EquipPosChange_Out) Decode(buffer *net.Buffer) {
}

func (this *EquipPosChange_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(31)
	buffer.WriteUint8(6)
}

func (this *EquipPosChange_Out) ByteSize() int {
	size := 2
	return size
}

func (this *Swap_In) Decode(buffer *net.Buffer) {
	this.EquipedId = int64(buffer.ReadUint64LE())
	this.InbagId = int64(buffer.ReadUint64LE())
}

func (this *Swap_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(31)
	buffer.WriteUint8(8)
	buffer.WriteUint64LE(uint64(this.EquipedId))
	buffer.WriteUint64LE(uint64(this.InbagId))
}

func (this *Swap_In) ByteSize() int {
	size := 18
	return size
}

func (this *Swap_Out) Decode(buffer *net.Buffer) {
}

func (this *Swap_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(31)
	buffer.WriteUint8(8)
}

func (this *Swap_Out) ByteSize() int {
	size := 2
	return size
}

func (this *IsBagFull_In) Decode(buffer *net.Buffer) {
}

func (this *IsBagFull_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(31)
	buffer.WriteUint8(9)
}

func (this *IsBagFull_In) ByteSize() int {
	size := 2
	return size
}

func (this *IsBagFull_Out) Decode(buffer *net.Buffer) {
	this.Full = buffer.ReadUint8() == 1
}

func (this *IsBagFull_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(31)
	buffer.WriteUint8(9)
	if this.Full {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *IsBagFull_Out) ByteSize() int {
	size := 3
	return size
}
