package daily_sign_in_api

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
	Sign(*net.Session, *Sign_In)
	SignPastDay(*net.Session, *SignPastDay_In)
}

type OutHandler interface {
	Info(*net.Session, *Info_Out)
	Sign(*net.Session, *Sign_Out)
	SignPastDay(*net.Session, *SignPastDay_Out)
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
		request := new(Sign_In)
		request.Decode(buffer)
		return request
	case 3:
		request := new(SignPastDay_In)
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
		request := new(Sign_Out)
		request.Decode(buffer)
		return request
	case 3:
		request := new(SignPastDay_Out)
		request.Decode(buffer)
		return request
	}
	_ = buffer
	panic("unsupported response")
}

type AwardType int8

const (
	AWARD_TYPE_NEW_PLAYER_AWARD AwardType = 0
	AWARD_TYPE_REGULAR_AWARD    AwardType = 1
)

type Info_In struct {
}

func (this *Info_In) Process(session *net.Session) {
	g_InHandler.Info(session, this)
}

func (this *Info_In) TypeName() string {
	return "daily_sign_in.info.in"
}

func (this *Info_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 22, 1
}

type Info_Out struct {
	Index   int8               `json:"index"`
	Records []Info_Out_Records `json:"records"`
}

type Info_Out_Records struct {
	AwardType  AwardType `json:"award_type"`
	AwardIndex int8      `json:"award_index"`
	Signed     bool      `json:"signed"`
}

func (this *Info_Out) Process(session *net.Session) {
	g_OutHandler.Info(session, this)
}

func (this *Info_Out) TypeName() string {
	return "daily_sign_in.info.out"
}

func (this *Info_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 22, 1
}

func (this *Info_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type Sign_In struct {
}

func (this *Sign_In) Process(session *net.Session) {
	g_InHandler.Sign(session, this)
}

func (this *Sign_In) TypeName() string {
	return "daily_sign_in.sign.in"
}

func (this *Sign_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 22, 2
}

type Sign_Out struct {
	Expired bool               `json:"expired"`
	Index   int8               `json:"index"`
	Records []Sign_Out_Records `json:"records"`
}

type Sign_Out_Records struct {
	AwardType  AwardType `json:"award_type"`
	AwardIndex int8      `json:"award_index"`
	Signed     bool      `json:"signed"`
}

func (this *Sign_Out) Process(session *net.Session) {
	g_OutHandler.Sign(session, this)
}

func (this *Sign_Out) TypeName() string {
	return "daily_sign_in.sign.out"
}

func (this *Sign_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 22, 2
}

func (this *Sign_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type SignPastDay_In struct {
	Index int8 `json:"index"`
}

func (this *SignPastDay_In) Process(session *net.Session) {
	g_InHandler.SignPastDay(session, this)
}

func (this *SignPastDay_In) TypeName() string {
	return "daily_sign_in.sign_past_day.in"
}

func (this *SignPastDay_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 22, 3
}

type SignPastDay_Out struct {
	Expired bool `json:"expired"`
}

func (this *SignPastDay_Out) Process(session *net.Session) {
	g_OutHandler.SignPastDay(session, this)
}

func (this *SignPastDay_Out) TypeName() string {
	return "daily_sign_in.sign_past_day.out"
}

func (this *SignPastDay_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 22, 3
}

func (this *SignPastDay_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

func (this *Info_In) Decode(buffer *net.Buffer) {
}

func (this *Info_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(22)
	buffer.WriteUint8(1)
}

func (this *Info_In) ByteSize() int {
	size := 2
	return size
}

func (this *Info_Out) Decode(buffer *net.Buffer) {
	this.Index = int8(buffer.ReadUint8())
	this.Records = make([]Info_Out_Records, buffer.ReadUint8())
	for i := 0; i < len(this.Records); i++ {
		this.Records[i].Decode(buffer)
	}
}

func (this *Info_Out_Records) Decode(buffer *net.Buffer) {
	this.AwardType = AwardType(buffer.ReadUint8())
	this.AwardIndex = int8(buffer.ReadUint8())
	this.Signed = buffer.ReadUint8() == 1
}

func (this *Info_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(22)
	buffer.WriteUint8(1)
	buffer.WriteUint8(uint8(this.Index))
	buffer.WriteUint8(uint8(len(this.Records)))
	for i := 0; i < len(this.Records); i++ {
		this.Records[i].Encode(buffer)
	}
}

func (this *Info_Out_Records) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(this.AwardType))
	buffer.WriteUint8(uint8(this.AwardIndex))
	if this.Signed {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *Info_Out) ByteSize() int {
	size := 4
	size += len(this.Records) * 3
	return size
}

func (this *Sign_In) Decode(buffer *net.Buffer) {
}

func (this *Sign_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(22)
	buffer.WriteUint8(2)
}

func (this *Sign_In) ByteSize() int {
	size := 2
	return size
}

func (this *Sign_Out) Decode(buffer *net.Buffer) {
	this.Expired = buffer.ReadUint8() == 1
	this.Index = int8(buffer.ReadUint8())
	this.Records = make([]Sign_Out_Records, buffer.ReadUint8())
	for i := 0; i < len(this.Records); i++ {
		this.Records[i].Decode(buffer)
	}
}

func (this *Sign_Out_Records) Decode(buffer *net.Buffer) {
	this.AwardType = AwardType(buffer.ReadUint8())
	this.AwardIndex = int8(buffer.ReadUint8())
	this.Signed = buffer.ReadUint8() == 1
}

func (this *Sign_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(22)
	buffer.WriteUint8(2)
	if this.Expired {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
	buffer.WriteUint8(uint8(this.Index))
	buffer.WriteUint8(uint8(len(this.Records)))
	for i := 0; i < len(this.Records); i++ {
		this.Records[i].Encode(buffer)
	}
}

func (this *Sign_Out_Records) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(this.AwardType))
	buffer.WriteUint8(uint8(this.AwardIndex))
	if this.Signed {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *Sign_Out) ByteSize() int {
	size := 5
	size += len(this.Records) * 3
	return size
}

func (this *SignPastDay_In) Decode(buffer *net.Buffer) {
	this.Index = int8(buffer.ReadUint8())
}

func (this *SignPastDay_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(22)
	buffer.WriteUint8(3)
	buffer.WriteUint8(uint8(this.Index))
}

func (this *SignPastDay_In) ByteSize() int {
	size := 3
	return size
}

func (this *SignPastDay_Out) Decode(buffer *net.Buffer) {
	this.Expired = buffer.ReadUint8() == 1
}

func (this *SignPastDay_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(22)
	buffer.WriteUint8(3)
	if this.Expired {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *SignPastDay_Out) ByteSize() int {
	size := 3
	return size
}
