package tower_api

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
	GetInfo(*net.Session, *GetInfo_In)
	UseLadder(*net.Session, *UseLadder_In)
}

type OutHandler interface {
	GetInfo(*net.Session, *GetInfo_Out)
	UseLadder(*net.Session, *UseLadder_Out)
}

func DecodeIn(message []byte) Request {
	var actionId = message[0]
	var buffer = net.NewBuffer(message[1:])
	switch actionId {
	case 1:
		request := new(GetInfo_In)
		request.Decode(buffer)
		return request
	case 2:
		request := new(UseLadder_In)
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
		request := new(GetInfo_Out)
		request.Decode(buffer)
		return request
	case 2:
		request := new(UseLadder_Out)
		request.Decode(buffer)
		return request
	}
	_ = buffer
	panic("unsupported response")
}

type GetInfo_In struct {
}

func (this *GetInfo_In) Process(session *net.Session) {
	g_InHandler.GetInfo(session, this)
}

func (this *GetInfo_In) TypeName() string {
	return "tower.get_info.in"
}

func (this *GetInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 15, 1
}

type GetInfo_Out struct {
	FloorNum int16                 `json:"floor_num"`
	Friends  []GetInfo_Out_Friends `json:"friends"`
}

type GetInfo_Out_Friends struct {
	RoleId   int8   `json:"role_id"`
	Nickname []byte `json:"nickname"`
	Level    int16  `json:"level"`
	Floor    int16  `json:"floor"`
}

func (this *GetInfo_Out) Process(session *net.Session) {
	g_OutHandler.GetInfo(session, this)
}

func (this *GetInfo_Out) TypeName() string {
	return "tower.get_info.out"
}

func (this *GetInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 15, 1
}

func (this *GetInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type UseLadder_In struct {
}

func (this *UseLadder_In) Process(session *net.Session) {
	g_InHandler.UseLadder(session, this)
}

func (this *UseLadder_In) TypeName() string {
	return "tower.use_ladder.in"
}

func (this *UseLadder_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 15, 2
}

type UseLadder_Out struct {
	FloorNum int16 `json:"floor_num"`
}

func (this *UseLadder_Out) Process(session *net.Session) {
	g_OutHandler.UseLadder(session, this)
}

func (this *UseLadder_Out) TypeName() string {
	return "tower.use_ladder.out"
}

func (this *UseLadder_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 15, 2
}

func (this *UseLadder_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

func (this *GetInfo_In) Decode(buffer *net.Buffer) {
}

func (this *GetInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(15)
	buffer.WriteUint8(1)
}

func (this *GetInfo_In) ByteSize() int {
	size := 2
	return size
}

func (this *GetInfo_Out) Decode(buffer *net.Buffer) {
	this.FloorNum = int16(buffer.ReadUint16LE())
	this.Friends = make([]GetInfo_Out_Friends, buffer.ReadUint8())
	for i := 0; i < len(this.Friends); i++ {
		this.Friends[i].Decode(buffer)
	}
}

func (this *GetInfo_Out_Friends) Decode(buffer *net.Buffer) {
	this.RoleId = int8(buffer.ReadUint8())
	this.Nickname = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Level = int16(buffer.ReadUint16LE())
	this.Floor = int16(buffer.ReadUint16LE())
}

func (this *GetInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(15)
	buffer.WriteUint8(1)
	buffer.WriteUint16LE(uint16(this.FloorNum))
	buffer.WriteUint8(uint8(len(this.Friends)))
	for i := 0; i < len(this.Friends); i++ {
		this.Friends[i].Encode(buffer)
	}
}

func (this *GetInfo_Out_Friends) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint16LE(uint16(len(this.Nickname)))
	buffer.WriteBytes(this.Nickname)
	buffer.WriteUint16LE(uint16(this.Level))
	buffer.WriteUint16LE(uint16(this.Floor))
}

func (this *GetInfo_Out) ByteSize() int {
	size := 5
	for i := 0; i < len(this.Friends); i++ {
		size += this.Friends[i].ByteSize()
	}
	return size
}

func (this *GetInfo_Out_Friends) ByteSize() int {
	size := 7
	size += len(this.Nickname)
	return size
}

func (this *UseLadder_In) Decode(buffer *net.Buffer) {
}

func (this *UseLadder_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(15)
	buffer.WriteUint8(2)
}

func (this *UseLadder_In) ByteSize() int {
	size := 2
	return size
}

func (this *UseLadder_Out) Decode(buffer *net.Buffer) {
	this.FloorNum = int16(buffer.ReadUint16LE())
}

func (this *UseLadder_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(15)
	buffer.WriteUint8(2)
	buffer.WriteUint16LE(uint16(this.FloorNum))
}

func (this *UseLadder_Out) ByteSize() int {
	size := 4
	return size
}
