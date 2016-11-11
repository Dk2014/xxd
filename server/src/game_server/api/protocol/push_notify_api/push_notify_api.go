package push_notify_api

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
	PushInfo(*net.Session, *PushInfo_In)
	PushNotificationSwitch(*net.Session, *PushNotificationSwitch_In)
}

type OutHandler interface {
	PushInfo(*net.Session, *PushInfo_Out)
	PushNotificationSwitch(*net.Session, *PushNotificationSwitch_Out)
}

func DecodeIn(message []byte) Request {
	var actionId = message[0]
	var buffer = net.NewBuffer(message[1:])
	switch actionId {
	case 1:
		request := new(PushInfo_In)
		request.Decode(buffer)
		return request
	case 2:
		request := new(PushNotificationSwitch_In)
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
		request := new(PushInfo_Out)
		request.Decode(buffer)
		return request
	case 2:
		request := new(PushNotificationSwitch_Out)
		request.Decode(buffer)
		return request
	}
	_ = buffer
	panic("unsupported response")
}

type PushInfo_In struct {
}

func (this *PushInfo_In) Process(session *net.Session) {
	g_InHandler.PushInfo(session, this)
}

func (this *PushInfo_In) TypeName() string {
	return "push_notify.push_info.in"
}

func (this *PushInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 26, 1
}

type PushInfo_Out struct {
	EffectNotification []PushInfo_Out_EffectNotification `json:"effect_notification"`
}

type PushInfo_Out_EffectNotification struct {
	NotificationId int32 `json:"notification_id"`
}

func (this *PushInfo_Out) Process(session *net.Session) {
	g_OutHandler.PushInfo(session, this)
}

func (this *PushInfo_Out) TypeName() string {
	return "push_notify.push_info.out"
}

func (this *PushInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 26, 1
}

func (this *PushInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type PushNotificationSwitch_In struct {
	NotificationId int32 `json:"notification_id"`
	TurnOn         bool  `json:"turn_on"`
}

func (this *PushNotificationSwitch_In) Process(session *net.Session) {
	g_InHandler.PushNotificationSwitch(session, this)
}

func (this *PushNotificationSwitch_In) TypeName() string {
	return "push_notify.push_notification_switch.in"
}

func (this *PushNotificationSwitch_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 26, 2
}

type PushNotificationSwitch_Out struct {
}

func (this *PushNotificationSwitch_Out) Process(session *net.Session) {
	g_OutHandler.PushNotificationSwitch(session, this)
}

func (this *PushNotificationSwitch_Out) TypeName() string {
	return "push_notify.push_notification_switch.out"
}

func (this *PushNotificationSwitch_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 26, 2
}

func (this *PushNotificationSwitch_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

func (this *PushInfo_In) Decode(buffer *net.Buffer) {
}

func (this *PushInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(26)
	buffer.WriteUint8(1)
}

func (this *PushInfo_In) ByteSize() int {
	size := 2
	return size
}

func (this *PushInfo_Out) Decode(buffer *net.Buffer) {
	this.EffectNotification = make([]PushInfo_Out_EffectNotification, buffer.ReadUint8())
	for i := 0; i < len(this.EffectNotification); i++ {
		this.EffectNotification[i].Decode(buffer)
	}
}

func (this *PushInfo_Out_EffectNotification) Decode(buffer *net.Buffer) {
	this.NotificationId = int32(buffer.ReadUint32LE())
}

func (this *PushInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(26)
	buffer.WriteUint8(1)
	buffer.WriteUint8(uint8(len(this.EffectNotification)))
	for i := 0; i < len(this.EffectNotification); i++ {
		this.EffectNotification[i].Encode(buffer)
	}
}

func (this *PushInfo_Out_EffectNotification) Encode(buffer *net.Buffer) {
	buffer.WriteUint32LE(uint32(this.NotificationId))
}

func (this *PushInfo_Out) ByteSize() int {
	size := 3
	size += len(this.EffectNotification) * 4
	return size
}

func (this *PushNotificationSwitch_In) Decode(buffer *net.Buffer) {
	this.NotificationId = int32(buffer.ReadUint32LE())
	this.TurnOn = buffer.ReadUint8() == 1
}

func (this *PushNotificationSwitch_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(26)
	buffer.WriteUint8(2)
	buffer.WriteUint32LE(uint32(this.NotificationId))
	if this.TurnOn {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *PushNotificationSwitch_In) ByteSize() int {
	size := 7
	return size
}

func (this *PushNotificationSwitch_Out) Decode(buffer *net.Buffer) {
}

func (this *PushNotificationSwitch_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(26)
	buffer.WriteUint8(2)
}

func (this *PushNotificationSwitch_Out) ByteSize() int {
	size := 2
	return size
}
