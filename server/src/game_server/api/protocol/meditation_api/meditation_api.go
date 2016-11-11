package meditation_api

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
	MeditationInfo(*net.Session, *MeditationInfo_In)
	StartMeditation(*net.Session, *StartMeditation_In)
	StopMeditation(*net.Session, *StopMeditation_In)
}

type OutHandler interface {
	MeditationInfo(*net.Session, *MeditationInfo_Out)
	StartMeditation(*net.Session, *StartMeditation_Out)
	StopMeditation(*net.Session, *StopMeditation_Out)
}

func DecodeIn(message []byte) Request {
	var actionId = message[0]
	var buffer = net.NewBuffer(message[1:])
	switch actionId {
	case 1:
		request := new(MeditationInfo_In)
		request.Decode(buffer)
		return request
	case 2:
		request := new(StartMeditation_In)
		request.Decode(buffer)
		return request
	case 3:
		request := new(StopMeditation_In)
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
		request := new(MeditationInfo_Out)
		request.Decode(buffer)
		return request
	case 2:
		request := new(StartMeditation_Out)
		request.Decode(buffer)
		return request
	case 3:
		request := new(StopMeditation_Out)
		request.Decode(buffer)
		return request
	}
	_ = buffer
	panic("unsupported response")
}

type MeditationInfo_In struct {
}

func (this *MeditationInfo_In) Process(session *net.Session) {
	g_InHandler.MeditationInfo(session, this)
}

func (this *MeditationInfo_In) TypeName() string {
	return "meditation.meditation_info.in"
}

func (this *MeditationInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 27, 1
}

type MeditationInfo_Out struct {
	ExpAccumulateTime int32 `json:"exp_accumulate_time"`
	KeyAccumulateTime int32 `json:"key_accumulate_time"`
}

func (this *MeditationInfo_Out) Process(session *net.Session) {
	g_OutHandler.MeditationInfo(session, this)
}

func (this *MeditationInfo_Out) TypeName() string {
	return "meditation.meditation_info.out"
}

func (this *MeditationInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 27, 1
}

func (this *MeditationInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type StartMeditation_In struct {
	InClubhouse bool `json:"in_clubhouse"`
}

func (this *StartMeditation_In) Process(session *net.Session) {
	g_InHandler.StartMeditation(session, this)
}

func (this *StartMeditation_In) TypeName() string {
	return "meditation.start_meditation.in"
}

func (this *StartMeditation_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 27, 2
}

type StartMeditation_Out struct {
	ExpAccumulateTime int32 `json:"exp_accumulate_time"`
	KeyAccumulateTime int32 `json:"key_accumulate_time"`
}

func (this *StartMeditation_Out) Process(session *net.Session) {
	g_OutHandler.StartMeditation(session, this)
}

func (this *StartMeditation_Out) TypeName() string {
	return "meditation.start_meditation.out"
}

func (this *StartMeditation_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 27, 2
}

func (this *StartMeditation_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type StopMeditation_In struct {
	InClubhouse bool `json:"in_clubhouse"`
}

func (this *StopMeditation_In) Process(session *net.Session) {
	g_InHandler.StopMeditation(session, this)
}

func (this *StopMeditation_In) TypeName() string {
	return "meditation.stop_meditation.in"
}

func (this *StopMeditation_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 27, 3
}

type StopMeditation_Out struct {
}

func (this *StopMeditation_Out) Process(session *net.Session) {
	g_OutHandler.StopMeditation(session, this)
}

func (this *StopMeditation_Out) TypeName() string {
	return "meditation.stop_meditation.out"
}

func (this *StopMeditation_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 27, 3
}

func (this *StopMeditation_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

func (this *MeditationInfo_In) Decode(buffer *net.Buffer) {
}

func (this *MeditationInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(27)
	buffer.WriteUint8(1)
}

func (this *MeditationInfo_In) ByteSize() int {
	size := 2
	return size
}

func (this *MeditationInfo_Out) Decode(buffer *net.Buffer) {
	this.ExpAccumulateTime = int32(buffer.ReadUint32LE())
	this.KeyAccumulateTime = int32(buffer.ReadUint32LE())
}

func (this *MeditationInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(27)
	buffer.WriteUint8(1)
	buffer.WriteUint32LE(uint32(this.ExpAccumulateTime))
	buffer.WriteUint32LE(uint32(this.KeyAccumulateTime))
}

func (this *MeditationInfo_Out) ByteSize() int {
	size := 10
	return size
}

func (this *StartMeditation_In) Decode(buffer *net.Buffer) {
	this.InClubhouse = buffer.ReadUint8() == 1
}

func (this *StartMeditation_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(27)
	buffer.WriteUint8(2)
	if this.InClubhouse {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *StartMeditation_In) ByteSize() int {
	size := 3
	return size
}

func (this *StartMeditation_Out) Decode(buffer *net.Buffer) {
	this.ExpAccumulateTime = int32(buffer.ReadUint32LE())
	this.KeyAccumulateTime = int32(buffer.ReadUint32LE())
}

func (this *StartMeditation_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(27)
	buffer.WriteUint8(2)
	buffer.WriteUint32LE(uint32(this.ExpAccumulateTime))
	buffer.WriteUint32LE(uint32(this.KeyAccumulateTime))
}

func (this *StartMeditation_Out) ByteSize() int {
	size := 10
	return size
}

func (this *StopMeditation_In) Decode(buffer *net.Buffer) {
	this.InClubhouse = buffer.ReadUint8() == 1
}

func (this *StopMeditation_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(27)
	buffer.WriteUint8(3)
	if this.InClubhouse {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *StopMeditation_In) ByteSize() int {
	size := 3
	return size
}

func (this *StopMeditation_Out) Decode(buffer *net.Buffer) {
}

func (this *StopMeditation_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(27)
	buffer.WriteUint8(3)
}

func (this *StopMeditation_Out) ByteSize() int {
	size := 2
	return size
}
