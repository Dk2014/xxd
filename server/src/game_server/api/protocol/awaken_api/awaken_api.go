package awaken_api

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
	AwakenInfo(*net.Session, *AwakenInfo_In)
	LevelupAttr(*net.Session, *LevelupAttr_In)
}

type OutHandler interface {
	AwakenInfo(*net.Session, *AwakenInfo_Out)
	LevelupAttr(*net.Session, *LevelupAttr_Out)
}

func DecodeIn(message []byte) Request {
	var actionId = message[0]
	var buffer = net.NewBuffer(message[1:])
	switch actionId {
	case 0:
		request := new(AwakenInfo_In)
		request.Decode(buffer)
		return request
	case 1:
		request := new(LevelupAttr_In)
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
		request := new(AwakenInfo_Out)
		request.Decode(buffer)
		return request
	case 1:
		request := new(LevelupAttr_Out)
		request.Decode(buffer)
		return request
	}
	_ = buffer
	panic("unsupported response")
}

type AwakenInfo_In struct {
	RoleId int8 `json:"role_id"`
}

func (this *AwakenInfo_In) Process(session *net.Session) {
	g_InHandler.AwakenInfo(session, this)
}

func (this *AwakenInfo_In) TypeName() string {
	return "awaken.awaken_info.in"
}

func (this *AwakenInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 38, 0
}

type AwakenInfo_Out struct {
	RoleId int8                   `json:"role_id"`
	Attrs  []AwakenInfo_Out_Attrs `json:"attrs"`
}

type AwakenInfo_Out_Attrs struct {
	AttrImpl int16 `json:"attr_impl"`
	Level    int8  `json:"level"`
}

func (this *AwakenInfo_Out) Process(session *net.Session) {
	g_OutHandler.AwakenInfo(session, this)
}

func (this *AwakenInfo_Out) TypeName() string {
	return "awaken.awaken_info.out"
}

func (this *AwakenInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 38, 0
}

func (this *AwakenInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type LevelupAttr_In struct {
	RoleId   int8  `json:"role_id"`
	AttrImpl int16 `json:"attr_impl"`
}

func (this *LevelupAttr_In) Process(session *net.Session) {
	g_InHandler.LevelupAttr(session, this)
}

func (this *LevelupAttr_In) TypeName() string {
	return "awaken.levelup_attr.in"
}

func (this *LevelupAttr_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 38, 1
}

type LevelupAttr_Out struct {
	RoleId   int8  `json:"role_id"`
	AttrImpl int16 `json:"attr_impl"`
}

func (this *LevelupAttr_Out) Process(session *net.Session) {
	g_OutHandler.LevelupAttr(session, this)
}

func (this *LevelupAttr_Out) TypeName() string {
	return "awaken.levelup_attr.out"
}

func (this *LevelupAttr_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 38, 1
}

func (this *LevelupAttr_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

func (this *AwakenInfo_In) Decode(buffer *net.Buffer) {
	this.RoleId = int8(buffer.ReadUint8())
}

func (this *AwakenInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(38)
	buffer.WriteUint8(0)
	buffer.WriteUint8(uint8(this.RoleId))
}

func (this *AwakenInfo_In) ByteSize() int {
	size := 3
	return size
}

func (this *AwakenInfo_Out) Decode(buffer *net.Buffer) {
	this.RoleId = int8(buffer.ReadUint8())
	this.Attrs = make([]AwakenInfo_Out_Attrs, buffer.ReadUint8())
	for i := 0; i < len(this.Attrs); i++ {
		this.Attrs[i].Decode(buffer)
	}
}

func (this *AwakenInfo_Out_Attrs) Decode(buffer *net.Buffer) {
	this.AttrImpl = int16(buffer.ReadUint16LE())
	this.Level = int8(buffer.ReadUint8())
}

func (this *AwakenInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(38)
	buffer.WriteUint8(0)
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint8(uint8(len(this.Attrs)))
	for i := 0; i < len(this.Attrs); i++ {
		this.Attrs[i].Encode(buffer)
	}
}

func (this *AwakenInfo_Out_Attrs) Encode(buffer *net.Buffer) {
	buffer.WriteUint16LE(uint16(this.AttrImpl))
	buffer.WriteUint8(uint8(this.Level))
}

func (this *AwakenInfo_Out) ByteSize() int {
	size := 4
	size += len(this.Attrs) * 3
	return size
}

func (this *LevelupAttr_In) Decode(buffer *net.Buffer) {
	this.RoleId = int8(buffer.ReadUint8())
	this.AttrImpl = int16(buffer.ReadUint16LE())
}

func (this *LevelupAttr_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(38)
	buffer.WriteUint8(1)
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint16LE(uint16(this.AttrImpl))
}

func (this *LevelupAttr_In) ByteSize() int {
	size := 5
	return size
}

func (this *LevelupAttr_Out) Decode(buffer *net.Buffer) {
	this.RoleId = int8(buffer.ReadUint8())
	this.AttrImpl = int16(buffer.ReadUint16LE())
}

func (this *LevelupAttr_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(38)
	buffer.WriteUint8(1)
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint16LE(uint16(this.AttrImpl))
}

func (this *LevelupAttr_Out) ByteSize() int {
	size := 5
	return size
}
