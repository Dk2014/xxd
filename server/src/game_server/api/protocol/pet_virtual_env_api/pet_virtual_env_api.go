package pet_virtual_env_api

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
	TakeAward(*net.Session, *TakeAward_In)
	AutoFight(*net.Session, *AutoFight_In)
}

type OutHandler interface {
	Info(*net.Session, *Info_Out)
	TakeAward(*net.Session, *TakeAward_Out)
	AutoFight(*net.Session, *AutoFight_Out)
	PveKills(*net.Session, *PveKills_Out)
}

func DecodeIn(message []byte) Request {
	var actionId = message[0]
	var buffer = net.NewBuffer(message[1:])
	switch actionId {
	case 1:
		request := new(Info_In)
		request.Decode(buffer)
		return request
	case 2:
		request := new(TakeAward_In)
		request.Decode(buffer)
		return request
	case 3:
		request := new(AutoFight_In)
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
		request := new(Info_Out)
		request.Decode(buffer)
		return request
	case 2:
		request := new(TakeAward_Out)
		request.Decode(buffer)
		return request
	case 3:
		request := new(AutoFight_Out)
		request.Decode(buffer)
		return request
	case 4:
		request := new(PveKills_Out)
		request.Decode(buffer)
		return request
	}
	_ = buffer
	panic("unsupported response")
}

type Info_In struct {
}

func (this *Info_In) Process(session *net.Session) {
	g_InHandler.Info(session, this)
}

func (this *Info_In) TypeName() string {
	return "pet_virtual_env.info.in"
}

func (this *Info_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 28, 1
}

type Info_Out struct {
	DailyNum              int8  `json:"daily_num"`
	MaxFloor              int16 `json:"max_floor"`
	MaxAwardedFloor       int16 `json:"max_awarded_floor"`
	UnpassedFloorEnemyNum int16 `json:"unpassed_floor_enemy_num"`
}

func (this *Info_Out) Process(session *net.Session) {
	g_OutHandler.Info(session, this)
}

func (this *Info_Out) TypeName() string {
	return "pet_virtual_env.info.out"
}

func (this *Info_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 28, 1
}

func (this *Info_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type TakeAward_In struct {
}

func (this *TakeAward_In) Process(session *net.Session) {
	g_InHandler.TakeAward(session, this)
}

func (this *TakeAward_In) TypeName() string {
	return "pet_virtual_env.take_award.in"
}

func (this *TakeAward_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 28, 2
}

type TakeAward_Out struct {
}

func (this *TakeAward_Out) Process(session *net.Session) {
	g_OutHandler.TakeAward(session, this)
}

func (this *TakeAward_Out) TypeName() string {
	return "pet_virtual_env.take_award.out"
}

func (this *TakeAward_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 28, 2
}

func (this *TakeAward_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type AutoFight_In struct {
	Floor int16 `json:"floor"`
}

func (this *AutoFight_In) Process(session *net.Session) {
	g_InHandler.AutoFight(session, this)
}

func (this *AutoFight_In) TypeName() string {
	return "pet_virtual_env.auto_fight.in"
}

func (this *AutoFight_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 28, 3
}

type AutoFight_Out struct {
}

func (this *AutoFight_Out) Process(session *net.Session) {
	g_OutHandler.AutoFight(session, this)
}

func (this *AutoFight_Out) TypeName() string {
	return "pet_virtual_env.auto_fight.out"
}

func (this *AutoFight_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 28, 3
}

func (this *AutoFight_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type PveKills_Out struct {
	Num int16 `json:"num"`
}

func (this *PveKills_Out) Process(session *net.Session) {
	g_OutHandler.PveKills(session, this)
}

func (this *PveKills_Out) TypeName() string {
	return "pet_virtual_env.pve_kills.out"
}

func (this *PveKills_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 28, 4
}

func (this *PveKills_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

func (this *Info_In) Decode(buffer *net.Buffer) {
}

func (this *Info_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(28)
	buffer.WriteUint8(1)
}

func (this *Info_In) ByteSize() int {
	size := 2
	return size
}

func (this *Info_Out) Decode(buffer *net.Buffer) {
	this.DailyNum = int8(buffer.ReadUint8())
	this.MaxFloor = int16(buffer.ReadUint16LE())
	this.MaxAwardedFloor = int16(buffer.ReadUint16LE())
	this.UnpassedFloorEnemyNum = int16(buffer.ReadUint16LE())
}

func (this *Info_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(28)
	buffer.WriteUint8(1)
	buffer.WriteUint8(uint8(this.DailyNum))
	buffer.WriteUint16LE(uint16(this.MaxFloor))
	buffer.WriteUint16LE(uint16(this.MaxAwardedFloor))
	buffer.WriteUint16LE(uint16(this.UnpassedFloorEnemyNum))
}

func (this *Info_Out) ByteSize() int {
	size := 9
	return size
}

func (this *TakeAward_In) Decode(buffer *net.Buffer) {
}

func (this *TakeAward_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(28)
	buffer.WriteUint8(2)
}

func (this *TakeAward_In) ByteSize() int {
	size := 2
	return size
}

func (this *TakeAward_Out) Decode(buffer *net.Buffer) {
}

func (this *TakeAward_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(28)
	buffer.WriteUint8(2)
}

func (this *TakeAward_Out) ByteSize() int {
	size := 2
	return size
}

func (this *AutoFight_In) Decode(buffer *net.Buffer) {
	this.Floor = int16(buffer.ReadUint16LE())
}

func (this *AutoFight_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(28)
	buffer.WriteUint8(3)
	buffer.WriteUint16LE(uint16(this.Floor))
}

func (this *AutoFight_In) ByteSize() int {
	size := 4
	return size
}

func (this *AutoFight_Out) Decode(buffer *net.Buffer) {
}

func (this *AutoFight_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(28)
	buffer.WriteUint8(3)
}

func (this *AutoFight_Out) ByteSize() int {
	size := 2
	return size
}

func (this *PveKills_Out) Decode(buffer *net.Buffer) {
	this.Num = int16(buffer.ReadUint16LE())
}

func (this *PveKills_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(28)
	buffer.WriteUint8(4)
	buffer.WriteUint16LE(uint16(this.Num))
}

func (this *PveKills_Out) ByteSize() int {
	size := 4
	return size
}
