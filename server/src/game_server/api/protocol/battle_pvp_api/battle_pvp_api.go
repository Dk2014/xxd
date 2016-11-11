package battle_pvp_api

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
	InviteFight(*net.Session, *InviteFight_In)
	ReplyInvite(*net.Session, *ReplyInvite_In)
	EnterPvp(*net.Session, *EnterPvp_In)
	CancelInvite(*net.Session, *CancelInvite_In)
}

type OutHandler interface {
	InviteFight(*net.Session, *InviteFight_Out)
	ReplyInvite(*net.Session, *ReplyInvite_Out)
	EnterPvp(*net.Session, *EnterPvp_Out)
	CancelInvite(*net.Session, *CancelInvite_Out)
}

func DecodeIn(message []byte) Request {
	var actionId = message[0]
	var buffer = net.NewBuffer(message[1:])
	switch actionId {
	case 0:
		request := new(InviteFight_In)
		request.Decode(buffer)
		return request
	case 1:
		request := new(ReplyInvite_In)
		request.Decode(buffer)
		return request
	case 2:
		request := new(EnterPvp_In)
		request.Decode(buffer)
		return request
	case 3:
		request := new(CancelInvite_In)
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
		request := new(InviteFight_Out)
		request.Decode(buffer)
		return request
	case 1:
		request := new(ReplyInvite_Out)
		request.Decode(buffer)
		return request
	case 2:
		request := new(EnterPvp_Out)
		request.Decode(buffer)
		return request
	case 3:
		request := new(CancelInvite_Out)
		request.Decode(buffer)
		return request
	}
	_ = buffer
	panic("unsupported response")
}

type InviteFight_In struct {
	TargetPid int64
}

func (this *InviteFight_In) Process(session *net.Session) {
	g_InHandler.InviteFight(session, this)
}

func (this *InviteFight_In) TypeName() string {
	return "battle_pvp.invite_fight.in"
}

func (this *InviteFight_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 40, 0
}

type InviteFight_Out struct {
	Result int8
}

func (this *InviteFight_Out) Process(session *net.Session) {
	g_OutHandler.InviteFight(session, this)
}

func (this *InviteFight_Out) TypeName() string {
	return "battle_pvp.invite_fight.out"
}

func (this *InviteFight_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 40, 0
}

func (this *InviteFight_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ReplyInvite_In struct {
	TargetPid int64
	ReplyType int8
}

func (this *ReplyInvite_In) Process(session *net.Session) {
	g_InHandler.ReplyInvite(session, this)
}

func (this *ReplyInvite_In) TypeName() string {
	return "battle_pvp.reply_invite.in"
}

func (this *ReplyInvite_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 40, 1
}

type ReplyInvite_Out struct {
}

func (this *ReplyInvite_Out) Process(session *net.Session) {
	g_OutHandler.ReplyInvite(session, this)
}

func (this *ReplyInvite_Out) TypeName() string {
	return "battle_pvp.reply_invite.out"
}

func (this *ReplyInvite_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 40, 1
}

func (this *ReplyInvite_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type EnterPvp_In struct {
	InviterPid int64
}

func (this *EnterPvp_In) Process(session *net.Session) {
	g_InHandler.EnterPvp(session, this)
}

func (this *EnterPvp_In) TypeName() string {
	return "battle_pvp.enter_pvp.in"
}

func (this *EnterPvp_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 40, 2
}

type EnterPvp_Out struct {
	Result int8
}

func (this *EnterPvp_Out) Process(session *net.Session) {
	g_OutHandler.EnterPvp(session, this)
}

func (this *EnterPvp_Out) TypeName() string {
	return "battle_pvp.enter_pvp.out"
}

func (this *EnterPvp_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 40, 2
}

func (this *EnterPvp_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type CancelInvite_In struct {
}

func (this *CancelInvite_In) Process(session *net.Session) {
	g_InHandler.CancelInvite(session, this)
}

func (this *CancelInvite_In) TypeName() string {
	return "battle_pvp.cancel_invite.in"
}

func (this *CancelInvite_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 40, 3
}

type CancelInvite_Out struct {
}

func (this *CancelInvite_Out) Process(session *net.Session) {
	g_OutHandler.CancelInvite(session, this)
}

func (this *CancelInvite_Out) TypeName() string {
	return "battle_pvp.cancel_invite.out"
}

func (this *CancelInvite_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 40, 3
}

func (this *CancelInvite_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

func (this *InviteFight_In) Decode(buffer *net.Buffer) {
	this.TargetPid = int64(buffer.ReadUint64LE())
}

func (this *InviteFight_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(40)
	buffer.WriteUint8(0)
	buffer.WriteUint64LE(uint64(this.TargetPid))
}

func (this *InviteFight_In) ByteSize() int {
	size := 10
	return size
}

func (this *InviteFight_Out) Decode(buffer *net.Buffer) {
	this.Result = int8(buffer.ReadUint8())
}

func (this *InviteFight_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(40)
	buffer.WriteUint8(0)
	buffer.WriteUint8(uint8(this.Result))
}

func (this *InviteFight_Out) ByteSize() int {
	size := 3
	return size
}

func (this *ReplyInvite_In) Decode(buffer *net.Buffer) {
	this.TargetPid = int64(buffer.ReadUint64LE())
	this.ReplyType = int8(buffer.ReadUint8())
}

func (this *ReplyInvite_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(40)
	buffer.WriteUint8(1)
	buffer.WriteUint64LE(uint64(this.TargetPid))
	buffer.WriteUint8(uint8(this.ReplyType))
}

func (this *ReplyInvite_In) ByteSize() int {
	size := 11
	return size
}

func (this *ReplyInvite_Out) Decode(buffer *net.Buffer) {
}

func (this *ReplyInvite_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(40)
	buffer.WriteUint8(1)
}

func (this *ReplyInvite_Out) ByteSize() int {
	size := 2
	return size
}

func (this *EnterPvp_In) Decode(buffer *net.Buffer) {
	this.InviterPid = int64(buffer.ReadUint64LE())
}

func (this *EnterPvp_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(40)
	buffer.WriteUint8(2)
	buffer.WriteUint64LE(uint64(this.InviterPid))
}

func (this *EnterPvp_In) ByteSize() int {
	size := 10
	return size
}

func (this *EnterPvp_Out) Decode(buffer *net.Buffer) {
	this.Result = int8(buffer.ReadUint8())
}

func (this *EnterPvp_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(40)
	buffer.WriteUint8(2)
	buffer.WriteUint8(uint8(this.Result))
}

func (this *EnterPvp_Out) ByteSize() int {
	size := 3
	return size
}

func (this *CancelInvite_In) Decode(buffer *net.Buffer) {
}

func (this *CancelInvite_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(40)
	buffer.WriteUint8(3)
}

func (this *CancelInvite_In) ByteSize() int {
	size := 2
	return size
}

func (this *CancelInvite_Out) Decode(buffer *net.Buffer) {
}

func (this *CancelInvite_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(40)
	buffer.WriteUint8(3)
}

func (this *CancelInvite_Out) ByteSize() int {
	size := 2
	return size
}
