package sword_soul_api

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
	Draw(*net.Session, *Draw_In)
	Upgrade(*net.Session, *Upgrade_In)
	Exchange(*net.Session, *Exchange_In)
	Equip(*net.Session, *Equip_In)
	Unequip(*net.Session, *Unequip_In)
	EquipPosChange(*net.Session, *EquipPosChange_In)
	Swap(*net.Session, *Swap_In)
	IsBagFull(*net.Session, *IsBagFull_In)
	EmptyPosNum(*net.Session, *EmptyPosNum_In)
}

type OutHandler interface {
	Info(*net.Session, *Info_Out)
	Draw(*net.Session, *Draw_Out)
	Upgrade(*net.Session, *Upgrade_Out)
	Exchange(*net.Session, *Exchange_Out)
	Equip(*net.Session, *Equip_Out)
	Unequip(*net.Session, *Unequip_Out)
	EquipPosChange(*net.Session, *EquipPosChange_Out)
	Swap(*net.Session, *Swap_Out)
	IsBagFull(*net.Session, *IsBagFull_Out)
	EmptyPosNum(*net.Session, *EmptyPosNum_Out)
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
		request := new(Draw_In)
		request.Decode(buffer)
		return request
	case 2:
		request := new(Upgrade_In)
		request.Decode(buffer)
		return request
	case 3:
		request := new(Exchange_In)
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
	case 10:
		request := new(EmptyPosNum_In)
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
		request := new(Draw_Out)
		request.Decode(buffer)
		return request
	case 2:
		request := new(Upgrade_Out)
		request.Decode(buffer)
		return request
	case 3:
		request := new(Exchange_Out)
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
	case 10:
		request := new(EmptyPosNum_Out)
		request.Decode(buffer)
		return request
	}
	_ = buffer
	panic("unsupported response")
}

type DrawType int8

const (
	DRAW_TYPE_COIN  DrawType = 0
	DRAW_TYPE_INGOT DrawType = 1
)

type EquipPos int8

const (
	EQUIP_POS_POS1 EquipPos = 0
	EQUIP_POS_POS2 EquipPos = 1
	EQUIP_POS_POS3 EquipPos = 2
	EQUIP_POS_POS4 EquipPos = 3
	EQUIP_POS_POS5 EquipPos = 4
	EQUIP_POS_POS6 EquipPos = 5
	EQUIP_POS_POS7 EquipPos = 6
	EQUIP_POS_POS8 EquipPos = 7
	EQUIP_POS_POS9 EquipPos = 8
	EQUIP_POS_NUM  EquipPos = 9
)

type Box int8

const (
	BOX_A Box = 0
	BOX_B Box = 1
	BOX_C Box = 2
	BOX_D Box = 3
	BOX_E Box = 4
)

type Info_In struct {
}

func (this *Info_In) Process(session *net.Session) {
	g_InHandler.Info(session, this)
}

func (this *Info_In) TypeName() string {
	return "sword_soul.info.in"
}

func (this *Info_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 10, 0
}

type Info_Out struct {
	SwordSouls []Info_Out_SwordSouls `json:"sword_souls"`
	RoleEquip  []Info_Out_RoleEquip  `json:"role_equip"`
	BoxState   int8                  `json:"box_state"`
	Num        int16                 `json:"num"`
	IngotNum   int64                 `json:"ingot_num"`
	CdTime     int64                 `json:"cd_time"`
}

type Info_Out_SwordSouls struct {
	Id          int64 `json:"id"`
	SwordSoulId int16 `json:"sword_soul_id"`
	Exp         int32 `json:"exp"`
	Level       int8  `json:"level"`
}

type Info_Out_RoleEquip struct {
	RoleId int8  `json:"role_id"`
	Pos1Id int64 `json:"pos1_id"`
	Pos2Id int64 `json:"pos2_id"`
	Pos3Id int64 `json:"pos3_id"`
	Pos4Id int64 `json:"pos4_id"`
	Pos5Id int64 `json:"pos5_id"`
	Pos6Id int64 `json:"pos6_id"`
	Pos7Id int64 `json:"pos7_id"`
	Pos8Id int64 `json:"pos8_id"`
	Pos9Id int64 `json:"pos9_id"`
}

func (this *Info_Out) Process(session *net.Session) {
	g_OutHandler.Info(session, this)
}

func (this *Info_Out) TypeName() string {
	return "sword_soul.info.out"
}

func (this *Info_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 10, 0
}

func (this *Info_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type Draw_In struct {
	Box      Box      `json:"box"`
	DrawType DrawType `json:"draw_type"`
}

func (this *Draw_In) Process(session *net.Session) {
	g_InHandler.Draw(session, this)
}

func (this *Draw_In) TypeName() string {
	return "sword_soul.draw.in"
}

func (this *Draw_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 10, 1
}

type Draw_Out struct {
	Id          int64 `json:"id"`
	SwordSoulId int16 `json:"sword_soul_id"`
	Coins       int64 `json:"coins"`
	BoxState    int8  `json:"box_state"`
	Fragments   int64 `json:"fragments"`
}

func (this *Draw_Out) Process(session *net.Session) {
	g_OutHandler.Draw(session, this)
}

func (this *Draw_Out) TypeName() string {
	return "sword_soul.draw.out"
}

func (this *Draw_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 10, 1
}

func (this *Draw_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type Upgrade_In struct {
	TargetId   int64                   `json:"target_id"`
	SwordSouls []Upgrade_In_SwordSouls `json:"sword_souls"`
}

type Upgrade_In_SwordSouls struct {
	Id int64 `json:"id"`
}

func (this *Upgrade_In) Process(session *net.Session) {
	g_InHandler.Upgrade(session, this)
}

func (this *Upgrade_In) TypeName() string {
	return "sword_soul.upgrade.in"
}

func (this *Upgrade_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 10, 2
}

type Upgrade_Out struct {
	Id          int64 `json:"id"`
	SwordSoulId int16 `json:"sword_soul_id"`
	Exp         int32 `json:"exp"`
	Level       int8  `json:"level"`
}

func (this *Upgrade_Out) Process(session *net.Session) {
	g_OutHandler.Upgrade(session, this)
}

func (this *Upgrade_Out) TypeName() string {
	return "sword_soul.upgrade.out"
}

func (this *Upgrade_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 10, 2
}

func (this *Upgrade_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type Exchange_In struct {
	SwordSoulId int16 `json:"sword_soul_id"`
}

func (this *Exchange_In) Process(session *net.Session) {
	g_InHandler.Exchange(session, this)
}

func (this *Exchange_In) TypeName() string {
	return "sword_soul.exchange.in"
}

func (this *Exchange_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 10, 3
}

type Exchange_Out struct {
	Id int64 `json:"id"`
}

func (this *Exchange_Out) Process(session *net.Session) {
	g_OutHandler.Exchange(session, this)
}

func (this *Exchange_Out) TypeName() string {
	return "sword_soul.exchange.out"
}

func (this *Exchange_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 10, 3
}

func (this *Exchange_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type Equip_In struct {
	FromId   int64    `json:"from_id"`
	RoleId   int8     `json:"role_id"`
	EquipPos EquipPos `json:"equip_pos"`
}

func (this *Equip_In) Process(session *net.Session) {
	g_InHandler.Equip(session, this)
}

func (this *Equip_In) TypeName() string {
	return "sword_soul.equip.in"
}

func (this *Equip_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 10, 4
}

type Equip_Out struct {
}

func (this *Equip_Out) Process(session *net.Session) {
	g_OutHandler.Equip(session, this)
}

func (this *Equip_Out) TypeName() string {
	return "sword_soul.equip.out"
}

func (this *Equip_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 10, 4
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
	return "sword_soul.unequip.in"
}

func (this *Unequip_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 10, 5
}

type Unequip_Out struct {
}

func (this *Unequip_Out) Process(session *net.Session) {
	g_OutHandler.Unequip(session, this)
}

func (this *Unequip_Out) TypeName() string {
	return "sword_soul.unequip.out"
}

func (this *Unequip_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 10, 5
}

func (this *Unequip_Out) Bytes() []byte {
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
	return "sword_soul.equip_pos_change.in"
}

func (this *EquipPosChange_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 10, 6
}

type EquipPosChange_Out struct {
}

func (this *EquipPosChange_Out) Process(session *net.Session) {
	g_OutHandler.EquipPosChange(session, this)
}

func (this *EquipPosChange_Out) TypeName() string {
	return "sword_soul.equip_pos_change.out"
}

func (this *EquipPosChange_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 10, 6
}

func (this *EquipPosChange_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type Swap_In struct {
	RoleId int8  `json:"role_id"`
	FromId int64 `json:"from_id"`
	ToId   int64 `json:"to_id"`
}

func (this *Swap_In) Process(session *net.Session) {
	g_InHandler.Swap(session, this)
}

func (this *Swap_In) TypeName() string {
	return "sword_soul.swap.in"
}

func (this *Swap_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 10, 8
}

type Swap_Out struct {
}

func (this *Swap_Out) Process(session *net.Session) {
	g_OutHandler.Swap(session, this)
}

func (this *Swap_Out) TypeName() string {
	return "sword_soul.swap.out"
}

func (this *Swap_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 10, 8
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
	return "sword_soul.is_bag_full.in"
}

func (this *IsBagFull_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 10, 9
}

type IsBagFull_Out struct {
	IsFull bool `json:"is_full"`
}

func (this *IsBagFull_Out) Process(session *net.Session) {
	g_OutHandler.IsBagFull(session, this)
}

func (this *IsBagFull_Out) TypeName() string {
	return "sword_soul.is_bag_full.out"
}

func (this *IsBagFull_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 10, 9
}

func (this *IsBagFull_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type EmptyPosNum_In struct {
}

func (this *EmptyPosNum_In) Process(session *net.Session) {
	g_InHandler.EmptyPosNum(session, this)
}

func (this *EmptyPosNum_In) TypeName() string {
	return "sword_soul.empty_pos_num.in"
}

func (this *EmptyPosNum_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 10, 10
}

type EmptyPosNum_Out struct {
	EmptyPosNum int16 `json:"empty_pos_num"`
}

func (this *EmptyPosNum_Out) Process(session *net.Session) {
	g_OutHandler.EmptyPosNum(session, this)
}

func (this *EmptyPosNum_Out) TypeName() string {
	return "sword_soul.empty_pos_num.out"
}

func (this *EmptyPosNum_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 10, 10
}

func (this *EmptyPosNum_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

func (this *Info_In) Decode(buffer *net.Buffer) {
}

func (this *Info_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(10)
	buffer.WriteUint8(0)
}

func (this *Info_In) ByteSize() int {
	size := 2
	return size
}

func (this *Info_Out) Decode(buffer *net.Buffer) {
	this.SwordSouls = make([]Info_Out_SwordSouls, buffer.ReadUint8())
	for i := 0; i < len(this.SwordSouls); i++ {
		this.SwordSouls[i].Decode(buffer)
	}
	this.RoleEquip = make([]Info_Out_RoleEquip, buffer.ReadUint8())
	for i := 0; i < len(this.RoleEquip); i++ {
		this.RoleEquip[i].Decode(buffer)
	}
	this.BoxState = int8(buffer.ReadUint8())
	this.Num = int16(buffer.ReadUint16LE())
	this.IngotNum = int64(buffer.ReadUint64LE())
	this.CdTime = int64(buffer.ReadUint64LE())
}

func (this *Info_Out_SwordSouls) Decode(buffer *net.Buffer) {
	this.Id = int64(buffer.ReadUint64LE())
	this.SwordSoulId = int16(buffer.ReadUint16LE())
	this.Exp = int32(buffer.ReadUint32LE())
	this.Level = int8(buffer.ReadUint8())
}

func (this *Info_Out_RoleEquip) Decode(buffer *net.Buffer) {
	this.RoleId = int8(buffer.ReadUint8())
	this.Pos1Id = int64(buffer.ReadUint64LE())
	this.Pos2Id = int64(buffer.ReadUint64LE())
	this.Pos3Id = int64(buffer.ReadUint64LE())
	this.Pos4Id = int64(buffer.ReadUint64LE())
	this.Pos5Id = int64(buffer.ReadUint64LE())
	this.Pos6Id = int64(buffer.ReadUint64LE())
	this.Pos7Id = int64(buffer.ReadUint64LE())
	this.Pos8Id = int64(buffer.ReadUint64LE())
	this.Pos9Id = int64(buffer.ReadUint64LE())
}

func (this *Info_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(10)
	buffer.WriteUint8(0)
	buffer.WriteUint8(uint8(len(this.SwordSouls)))
	for i := 0; i < len(this.SwordSouls); i++ {
		this.SwordSouls[i].Encode(buffer)
	}
	buffer.WriteUint8(uint8(len(this.RoleEquip)))
	for i := 0; i < len(this.RoleEquip); i++ {
		this.RoleEquip[i].Encode(buffer)
	}
	buffer.WriteUint8(uint8(this.BoxState))
	buffer.WriteUint16LE(uint16(this.Num))
	buffer.WriteUint64LE(uint64(this.IngotNum))
	buffer.WriteUint64LE(uint64(this.CdTime))
}

func (this *Info_Out_SwordSouls) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.Id))
	buffer.WriteUint16LE(uint16(this.SwordSoulId))
	buffer.WriteUint32LE(uint32(this.Exp))
	buffer.WriteUint8(uint8(this.Level))
}

func (this *Info_Out_RoleEquip) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint64LE(uint64(this.Pos1Id))
	buffer.WriteUint64LE(uint64(this.Pos2Id))
	buffer.WriteUint64LE(uint64(this.Pos3Id))
	buffer.WriteUint64LE(uint64(this.Pos4Id))
	buffer.WriteUint64LE(uint64(this.Pos5Id))
	buffer.WriteUint64LE(uint64(this.Pos6Id))
	buffer.WriteUint64LE(uint64(this.Pos7Id))
	buffer.WriteUint64LE(uint64(this.Pos8Id))
	buffer.WriteUint64LE(uint64(this.Pos9Id))
}

func (this *Info_Out) ByteSize() int {
	size := 23
	size += len(this.SwordSouls) * 15
	size += len(this.RoleEquip) * 73
	return size
}

func (this *Draw_In) Decode(buffer *net.Buffer) {
	this.Box = Box(buffer.ReadUint8())
	this.DrawType = DrawType(buffer.ReadUint8())
}

func (this *Draw_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(10)
	buffer.WriteUint8(1)
	buffer.WriteUint8(uint8(this.Box))
	buffer.WriteUint8(uint8(this.DrawType))
}

func (this *Draw_In) ByteSize() int {
	size := 4
	return size
}

func (this *Draw_Out) Decode(buffer *net.Buffer) {
	this.Id = int64(buffer.ReadUint64LE())
	this.SwordSoulId = int16(buffer.ReadUint16LE())
	this.Coins = int64(buffer.ReadUint64LE())
	this.BoxState = int8(buffer.ReadUint8())
	this.Fragments = int64(buffer.ReadUint64LE())
}

func (this *Draw_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(10)
	buffer.WriteUint8(1)
	buffer.WriteUint64LE(uint64(this.Id))
	buffer.WriteUint16LE(uint16(this.SwordSoulId))
	buffer.WriteUint64LE(uint64(this.Coins))
	buffer.WriteUint8(uint8(this.BoxState))
	buffer.WriteUint64LE(uint64(this.Fragments))
}

func (this *Draw_Out) ByteSize() int {
	size := 29
	return size
}

func (this *Upgrade_In) Decode(buffer *net.Buffer) {
	this.TargetId = int64(buffer.ReadUint64LE())
	this.SwordSouls = make([]Upgrade_In_SwordSouls, buffer.ReadUint8())
	for i := 0; i < len(this.SwordSouls); i++ {
		this.SwordSouls[i].Decode(buffer)
	}
}

func (this *Upgrade_In_SwordSouls) Decode(buffer *net.Buffer) {
	this.Id = int64(buffer.ReadUint64LE())
}

func (this *Upgrade_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(10)
	buffer.WriteUint8(2)
	buffer.WriteUint64LE(uint64(this.TargetId))
	buffer.WriteUint8(uint8(len(this.SwordSouls)))
	for i := 0; i < len(this.SwordSouls); i++ {
		this.SwordSouls[i].Encode(buffer)
	}
}

func (this *Upgrade_In_SwordSouls) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.Id))
}

func (this *Upgrade_In) ByteSize() int {
	size := 11
	size += len(this.SwordSouls) * 8
	return size
}

func (this *Upgrade_Out) Decode(buffer *net.Buffer) {
	this.Id = int64(buffer.ReadUint64LE())
	this.SwordSoulId = int16(buffer.ReadUint16LE())
	this.Exp = int32(buffer.ReadUint32LE())
	this.Level = int8(buffer.ReadUint8())
}

func (this *Upgrade_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(10)
	buffer.WriteUint8(2)
	buffer.WriteUint64LE(uint64(this.Id))
	buffer.WriteUint16LE(uint16(this.SwordSoulId))
	buffer.WriteUint32LE(uint32(this.Exp))
	buffer.WriteUint8(uint8(this.Level))
}

func (this *Upgrade_Out) ByteSize() int {
	size := 17
	return size
}

func (this *Exchange_In) Decode(buffer *net.Buffer) {
	this.SwordSoulId = int16(buffer.ReadUint16LE())
}

func (this *Exchange_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(10)
	buffer.WriteUint8(3)
	buffer.WriteUint16LE(uint16(this.SwordSoulId))
}

func (this *Exchange_In) ByteSize() int {
	size := 4
	return size
}

func (this *Exchange_Out) Decode(buffer *net.Buffer) {
	this.Id = int64(buffer.ReadUint64LE())
}

func (this *Exchange_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(10)
	buffer.WriteUint8(3)
	buffer.WriteUint64LE(uint64(this.Id))
}

func (this *Exchange_Out) ByteSize() int {
	size := 10
	return size
}

func (this *Equip_In) Decode(buffer *net.Buffer) {
	this.FromId = int64(buffer.ReadUint64LE())
	this.RoleId = int8(buffer.ReadUint8())
	this.EquipPos = EquipPos(buffer.ReadUint8())
}

func (this *Equip_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(10)
	buffer.WriteUint8(4)
	buffer.WriteUint64LE(uint64(this.FromId))
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint8(uint8(this.EquipPos))
}

func (this *Equip_In) ByteSize() int {
	size := 12
	return size
}

func (this *Equip_Out) Decode(buffer *net.Buffer) {
}

func (this *Equip_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(10)
	buffer.WriteUint8(4)
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
	buffer.WriteUint8(10)
	buffer.WriteUint8(5)
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
	buffer.WriteUint8(10)
	buffer.WriteUint8(5)
}

func (this *Unequip_Out) ByteSize() int {
	size := 2
	return size
}

func (this *EquipPosChange_In) Decode(buffer *net.Buffer) {
	this.RoleId = int8(buffer.ReadUint8())
	this.FromId = int64(buffer.ReadUint64LE())
	this.ToPos = EquipPos(buffer.ReadUint8())
}

func (this *EquipPosChange_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(10)
	buffer.WriteUint8(6)
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
	buffer.WriteUint8(10)
	buffer.WriteUint8(6)
}

func (this *EquipPosChange_Out) ByteSize() int {
	size := 2
	return size
}

func (this *Swap_In) Decode(buffer *net.Buffer) {
	this.RoleId = int8(buffer.ReadUint8())
	this.FromId = int64(buffer.ReadUint64LE())
	this.ToId = int64(buffer.ReadUint64LE())
}

func (this *Swap_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(10)
	buffer.WriteUint8(8)
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint64LE(uint64(this.FromId))
	buffer.WriteUint64LE(uint64(this.ToId))
}

func (this *Swap_In) ByteSize() int {
	size := 19
	return size
}

func (this *Swap_Out) Decode(buffer *net.Buffer) {
}

func (this *Swap_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(10)
	buffer.WriteUint8(8)
}

func (this *Swap_Out) ByteSize() int {
	size := 2
	return size
}

func (this *IsBagFull_In) Decode(buffer *net.Buffer) {
}

func (this *IsBagFull_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(10)
	buffer.WriteUint8(9)
}

func (this *IsBagFull_In) ByteSize() int {
	size := 2
	return size
}

func (this *IsBagFull_Out) Decode(buffer *net.Buffer) {
	this.IsFull = buffer.ReadUint8() == 1
}

func (this *IsBagFull_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(10)
	buffer.WriteUint8(9)
	if this.IsFull {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *IsBagFull_Out) ByteSize() int {
	size := 3
	return size
}

func (this *EmptyPosNum_In) Decode(buffer *net.Buffer) {
}

func (this *EmptyPosNum_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(10)
	buffer.WriteUint8(10)
}

func (this *EmptyPosNum_In) ByteSize() int {
	size := 2
	return size
}

func (this *EmptyPosNum_Out) Decode(buffer *net.Buffer) {
	this.EmptyPosNum = int16(buffer.ReadUint16LE())
}

func (this *EmptyPosNum_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(10)
	buffer.WriteUint8(10)
	buffer.WriteUint16LE(uint16(this.EmptyPosNum))
}

func (this *EmptyPosNum_Out) ByteSize() int {
	size := 4
	return size
}
