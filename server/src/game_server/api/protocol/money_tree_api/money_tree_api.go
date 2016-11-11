package money_tree_api

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
	GetTreeStatus(*net.Session, *GetTreeStatus_In)
	GetTreeMoney(*net.Session, *GetTreeMoney_In)
	WaveTree(*net.Session, *WaveTree_In)
}

type OutHandler interface {
	GetTreeStatus(*net.Session, *GetTreeStatus_Out)
	GetTreeMoney(*net.Session, *GetTreeMoney_Out)
	WaveTree(*net.Session, *WaveTree_Out)
}

func DecodeIn(message []byte) Request {
	var actionId = message[0]
	var buffer = net.NewBuffer(message[1:])
	switch actionId {
	case 0:
		request := new(GetTreeStatus_In)
		request.Decode(buffer)
		return request
	case 1:
		request := new(GetTreeMoney_In)
		request.Decode(buffer)
		return request
	case 2:
		request := new(WaveTree_In)
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
		request := new(GetTreeStatus_Out)
		request.Decode(buffer)
		return request
	case 1:
		request := new(GetTreeMoney_Out)
		request.Decode(buffer)
		return request
	case 2:
		request := new(WaveTree_Out)
		request.Decode(buffer)
		return request
	}
	_ = buffer
	panic("unsupported response")
}

type GetTreeStatus_In struct {
}

func (this *GetTreeStatus_In) Process(session *net.Session) {
	g_InHandler.GetTreeStatus(session, this)
}

func (this *GetTreeStatus_In) TypeName() string {
	return "money_tree.get_tree_status.in"
}

func (this *GetTreeStatus_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 32, 0
}

type GetTreeStatus_Out struct {
	Times    int8  `json:"times"`
	Money    int32 `json:"money"`
	LastTime int64 `json:"last_time"`
	Remind   int8  `json:"remind"`
	Status   int8  `json:"status"`
}

func (this *GetTreeStatus_Out) Process(session *net.Session) {
	g_OutHandler.GetTreeStatus(session, this)
}

func (this *GetTreeStatus_Out) TypeName() string {
	return "money_tree.get_tree_status.out"
}

func (this *GetTreeStatus_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 32, 0
}

func (this *GetTreeStatus_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetTreeMoney_In struct {
}

func (this *GetTreeMoney_In) Process(session *net.Session) {
	g_InHandler.GetTreeMoney(session, this)
}

func (this *GetTreeMoney_In) TypeName() string {
	return "money_tree.get_tree_money.in"
}

func (this *GetTreeMoney_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 32, 1
}

type GetTreeMoney_Out struct {
	Code int8 `json:"code"`
}

func (this *GetTreeMoney_Out) Process(session *net.Session) {
	g_OutHandler.GetTreeMoney(session, this)
}

func (this *GetTreeMoney_Out) TypeName() string {
	return "money_tree.get_tree_money.out"
}

func (this *GetTreeMoney_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 32, 1
}

func (this *GetTreeMoney_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type WaveTree_In struct {
}

func (this *WaveTree_In) Process(session *net.Session) {
	g_InHandler.WaveTree(session, this)
}

func (this *WaveTree_In) TypeName() string {
	return "money_tree.wave_tree.in"
}

func (this *WaveTree_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 32, 2
}

type WaveTree_Out struct {
	Status  int8  `json:"status"`
	Money   int32 `json:"money"`
	Remaind int8  `json:"remaind"`
}

func (this *WaveTree_Out) Process(session *net.Session) {
	g_OutHandler.WaveTree(session, this)
}

func (this *WaveTree_Out) TypeName() string {
	return "money_tree.wave_tree.out"
}

func (this *WaveTree_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 32, 2
}

func (this *WaveTree_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

func (this *GetTreeStatus_In) Decode(buffer *net.Buffer) {
}

func (this *GetTreeStatus_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(32)
	buffer.WriteUint8(0)
}

func (this *GetTreeStatus_In) ByteSize() int {
	size := 2
	return size
}

func (this *GetTreeStatus_Out) Decode(buffer *net.Buffer) {
	this.Times = int8(buffer.ReadUint8())
	this.Money = int32(buffer.ReadUint32LE())
	this.LastTime = int64(buffer.ReadUint64LE())
	this.Remind = int8(buffer.ReadUint8())
	this.Status = int8(buffer.ReadUint8())
}

func (this *GetTreeStatus_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(32)
	buffer.WriteUint8(0)
	buffer.WriteUint8(uint8(this.Times))
	buffer.WriteUint32LE(uint32(this.Money))
	buffer.WriteUint64LE(uint64(this.LastTime))
	buffer.WriteUint8(uint8(this.Remind))
	buffer.WriteUint8(uint8(this.Status))
}

func (this *GetTreeStatus_Out) ByteSize() int {
	size := 17
	return size
}

func (this *GetTreeMoney_In) Decode(buffer *net.Buffer) {
}

func (this *GetTreeMoney_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(32)
	buffer.WriteUint8(1)
}

func (this *GetTreeMoney_In) ByteSize() int {
	size := 2
	return size
}

func (this *GetTreeMoney_Out) Decode(buffer *net.Buffer) {
	this.Code = int8(buffer.ReadUint8())
}

func (this *GetTreeMoney_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(32)
	buffer.WriteUint8(1)
	buffer.WriteUint8(uint8(this.Code))
}

func (this *GetTreeMoney_Out) ByteSize() int {
	size := 3
	return size
}

func (this *WaveTree_In) Decode(buffer *net.Buffer) {
}

func (this *WaveTree_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(32)
	buffer.WriteUint8(2)
}

func (this *WaveTree_In) ByteSize() int {
	size := 2
	return size
}

func (this *WaveTree_Out) Decode(buffer *net.Buffer) {
	this.Status = int8(buffer.ReadUint8())
	this.Money = int32(buffer.ReadUint32LE())
	this.Remaind = int8(buffer.ReadUint8())
}

func (this *WaveTree_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(32)
	buffer.WriteUint8(2)
	buffer.WriteUint8(uint8(this.Status))
	buffer.WriteUint32LE(uint32(this.Money))
	buffer.WriteUint8(uint8(this.Remaind))
}

func (this *WaveTree_Out) ByteSize() int {
	size := 8
	return size
}
