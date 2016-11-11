package fashion_api

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
	FashionInfo(*net.Session, *FashionInfo_In)
	DressFashion(*net.Session, *DressFashion_In)
}

type OutHandler interface {
	FashionInfo(*net.Session, *FashionInfo_Out)
	DressFashion(*net.Session, *DressFashion_Out)
}

func DecodeIn(message []byte) Request {
	var actionId = message[0]
	var buffer = net.NewBuffer(message[1:])
	switch actionId {
	case 1:
		request := new(FashionInfo_In)
		request.Decode(buffer)
		return request
	case 2:
		request := new(DressFashion_In)
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
		request := new(FashionInfo_Out)
		request.Decode(buffer)
		return request
	case 2:
		request := new(DressFashion_Out)
		request.Decode(buffer)
		return request
	}
	_ = buffer
	panic("unsupported response")
}

type FashionInfo_In struct {
}

func (this *FashionInfo_In) Process(session *net.Session) {
	g_InHandler.FashionInfo(session, this)
}

func (this *FashionInfo_In) TypeName() string {
	return "fashion.fashion_info.in"
}

func (this *FashionInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 25, 1
}

type FashionInfo_Out struct {
	DressCdTime      int64                      `json:"dress_cd_time"`
	DressedFashionId int16                      `json:"dressed_fashion_id"`
	Fashions         []FashionInfo_Out_Fashions `json:"fashions"`
}

type FashionInfo_Out_Fashions struct {
	FashionId  int16 `json:"fashion_id"`
	ExpireTime int64 `json:"expire_time"`
}

func (this *FashionInfo_Out) Process(session *net.Session) {
	g_OutHandler.FashionInfo(session, this)
}

func (this *FashionInfo_Out) TypeName() string {
	return "fashion.fashion_info.out"
}

func (this *FashionInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 25, 1
}

func (this *FashionInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type DressFashion_In struct {
	FashionId   int16 `json:"fashion_id"`
	InClubhouse bool  `json:"in_clubhouse"`
}

func (this *DressFashion_In) Process(session *net.Session) {
	g_InHandler.DressFashion(session, this)
}

func (this *DressFashion_In) TypeName() string {
	return "fashion.dress_fashion.in"
}

func (this *DressFashion_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 25, 2
}

type DressFashion_Out struct {
	DressCdTime int64 `json:"dress_cd_time"`
}

func (this *DressFashion_Out) Process(session *net.Session) {
	g_OutHandler.DressFashion(session, this)
}

func (this *DressFashion_Out) TypeName() string {
	return "fashion.dress_fashion.out"
}

func (this *DressFashion_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 25, 2
}

func (this *DressFashion_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

func (this *FashionInfo_In) Decode(buffer *net.Buffer) {
}

func (this *FashionInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(25)
	buffer.WriteUint8(1)
}

func (this *FashionInfo_In) ByteSize() int {
	size := 2
	return size
}

func (this *FashionInfo_Out) Decode(buffer *net.Buffer) {
	this.DressCdTime = int64(buffer.ReadUint64LE())
	this.DressedFashionId = int16(buffer.ReadUint16LE())
	this.Fashions = make([]FashionInfo_Out_Fashions, buffer.ReadUint8())
	for i := 0; i < len(this.Fashions); i++ {
		this.Fashions[i].Decode(buffer)
	}
}

func (this *FashionInfo_Out_Fashions) Decode(buffer *net.Buffer) {
	this.FashionId = int16(buffer.ReadUint16LE())
	this.ExpireTime = int64(buffer.ReadUint64LE())
}

func (this *FashionInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(25)
	buffer.WriteUint8(1)
	buffer.WriteUint64LE(uint64(this.DressCdTime))
	buffer.WriteUint16LE(uint16(this.DressedFashionId))
	buffer.WriteUint8(uint8(len(this.Fashions)))
	for i := 0; i < len(this.Fashions); i++ {
		this.Fashions[i].Encode(buffer)
	}
}

func (this *FashionInfo_Out_Fashions) Encode(buffer *net.Buffer) {
	buffer.WriteUint16LE(uint16(this.FashionId))
	buffer.WriteUint64LE(uint64(this.ExpireTime))
}

func (this *FashionInfo_Out) ByteSize() int {
	size := 13
	size += len(this.Fashions) * 10
	return size
}

func (this *DressFashion_In) Decode(buffer *net.Buffer) {
	this.FashionId = int16(buffer.ReadUint16LE())
	this.InClubhouse = buffer.ReadUint8() == 1
}

func (this *DressFashion_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(25)
	buffer.WriteUint8(2)
	buffer.WriteUint16LE(uint16(this.FashionId))
	if this.InClubhouse {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *DressFashion_In) ByteSize() int {
	size := 5
	return size
}

func (this *DressFashion_Out) Decode(buffer *net.Buffer) {
	this.DressCdTime = int64(buffer.ReadUint64LE())
}

func (this *DressFashion_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(25)
	buffer.WriteUint8(2)
	buffer.WriteUint64LE(uint64(this.DressCdTime))
}

func (this *DressFashion_Out) ByteSize() int {
	size := 10
	return size
}
