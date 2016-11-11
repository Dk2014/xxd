package vip_api

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
	GetTotal(*net.Session, *GetTotal_In)
	VipLevelTotal(*net.Session, *VipLevelTotal_In)
	BuyTimes(*net.Session, *BuyTimes_In)
}

type OutHandler interface {
	Info(*net.Session, *Info_Out)
	GetTotal(*net.Session, *GetTotal_Out)
	VipLevelTotal(*net.Session, *VipLevelTotal_Out)
	BuyTimes(*net.Session, *BuyTimes_Out)
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
		request := new(GetTotal_In)
		request.Decode(buffer)
		return request
	case 3:
		request := new(VipLevelTotal_In)
		request.Decode(buffer)
		return request
	case 4:
		request := new(BuyTimes_In)
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
		request := new(GetTotal_Out)
		request.Decode(buffer)
		return request
	case 3:
		request := new(VipLevelTotal_Out)
		request.Decode(buffer)
		return request
	case 4:
		request := new(BuyTimes_Out)
		request.Decode(buffer)
		return request
	}
	_ = buffer
	panic("unsupported response")
}

type BuyTimesType int8

const (
	BUY_TIMES_TYPE_BIWUCHANGCISHU BuyTimesType = 0
	BUY_TIMES_TYPE_RAINBOWSAODANG BuyTimesType = 1
)

type Info_In struct {
	PlayerId int64 `json:"player_id"`
}

func (this *Info_In) Process(session *net.Session) {
	g_InHandler.Info(session, this)
}

func (this *Info_In) TypeName() string {
	return "vip.info.in"
}

func (this *Info_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 20, 1
}

type Info_Out struct {
	Level  int16  `json:"level"`
	Ingot  int64  `json:"ingot"`
	CardId []byte `json:"card_id"`
}

func (this *Info_Out) Process(session *net.Session) {
	g_OutHandler.Info(session, this)
}

func (this *Info_Out) TypeName() string {
	return "vip.info.out"
}

func (this *Info_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 20, 1
}

func (this *Info_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetTotal_In struct {
}

func (this *GetTotal_In) Process(session *net.Session) {
	g_InHandler.GetTotal(session, this)
}

func (this *GetTotal_In) TypeName() string {
	return "vip.get_total.in"
}

func (this *GetTotal_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 20, 2
}

type GetTotal_Out struct {
	Total int64 `json:"total"`
}

func (this *GetTotal_Out) Process(session *net.Session) {
	g_OutHandler.GetTotal(session, this)
}

func (this *GetTotal_Out) TypeName() string {
	return "vip.get_total.out"
}

func (this *GetTotal_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 20, 2
}

func (this *GetTotal_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type VipLevelTotal_In struct {
}

func (this *VipLevelTotal_In) Process(session *net.Session) {
	g_InHandler.VipLevelTotal(session, this)
}

func (this *VipLevelTotal_In) TypeName() string {
	return "vip.vip_level_total.in"
}

func (this *VipLevelTotal_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 20, 3
}

type VipLevelTotal_Out struct {
	VipLevelArr []VipLevelTotal_Out_VipLevelArr `json:"vip_level_arr"`
}

type VipLevelTotal_Out_VipLevelArr struct {
	VipLevel int16 `json:"vip_level"`
	Total    int32 `json:"total"`
}

func (this *VipLevelTotal_Out) Process(session *net.Session) {
	g_OutHandler.VipLevelTotal(session, this)
}

func (this *VipLevelTotal_Out) TypeName() string {
	return "vip.vip_level_total.out"
}

func (this *VipLevelTotal_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 20, 3
}

func (this *VipLevelTotal_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type BuyTimes_In struct {
	Buytype BuyTimesType `json:"buytype"`
}

func (this *BuyTimes_In) Process(session *net.Session) {
	g_InHandler.BuyTimes(session, this)
}

func (this *BuyTimes_In) TypeName() string {
	return "vip.buy_times.in"
}

func (this *BuyTimes_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 20, 4
}

type BuyTimes_Out struct {
	Result bool `json:"result"`
}

func (this *BuyTimes_Out) Process(session *net.Session) {
	g_OutHandler.BuyTimes(session, this)
}

func (this *BuyTimes_Out) TypeName() string {
	return "vip.buy_times.out"
}

func (this *BuyTimes_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 20, 4
}

func (this *BuyTimes_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

func (this *Info_In) Decode(buffer *net.Buffer) {
	this.PlayerId = int64(buffer.ReadUint64LE())
}

func (this *Info_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(20)
	buffer.WriteUint8(1)
	buffer.WriteUint64LE(uint64(this.PlayerId))
}

func (this *Info_In) ByteSize() int {
	size := 10
	return size
}

func (this *Info_Out) Decode(buffer *net.Buffer) {
	this.Level = int16(buffer.ReadUint16LE())
	this.Ingot = int64(buffer.ReadUint64LE())
	this.CardId = buffer.ReadBytes(int(buffer.ReadUint16LE()))
}

func (this *Info_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(20)
	buffer.WriteUint8(1)
	buffer.WriteUint16LE(uint16(this.Level))
	buffer.WriteUint64LE(uint64(this.Ingot))
	buffer.WriteUint16LE(uint16(len(this.CardId)))
	buffer.WriteBytes(this.CardId)
}

func (this *Info_Out) ByteSize() int {
	size := 14
	size += len(this.CardId)
	return size
}

func (this *GetTotal_In) Decode(buffer *net.Buffer) {
}

func (this *GetTotal_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(20)
	buffer.WriteUint8(2)
}

func (this *GetTotal_In) ByteSize() int {
	size := 2
	return size
}

func (this *GetTotal_Out) Decode(buffer *net.Buffer) {
	this.Total = int64(buffer.ReadUint64LE())
}

func (this *GetTotal_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(20)
	buffer.WriteUint8(2)
	buffer.WriteUint64LE(uint64(this.Total))
}

func (this *GetTotal_Out) ByteSize() int {
	size := 10
	return size
}

func (this *VipLevelTotal_In) Decode(buffer *net.Buffer) {
}

func (this *VipLevelTotal_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(20)
	buffer.WriteUint8(3)
}

func (this *VipLevelTotal_In) ByteSize() int {
	size := 2
	return size
}

func (this *VipLevelTotal_Out) Decode(buffer *net.Buffer) {
	this.VipLevelArr = make([]VipLevelTotal_Out_VipLevelArr, buffer.ReadUint8())
	for i := 0; i < len(this.VipLevelArr); i++ {
		this.VipLevelArr[i].Decode(buffer)
	}
}

func (this *VipLevelTotal_Out_VipLevelArr) Decode(buffer *net.Buffer) {
	this.VipLevel = int16(buffer.ReadUint16LE())
	this.Total = int32(buffer.ReadUint32LE())
}

func (this *VipLevelTotal_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(20)
	buffer.WriteUint8(3)
	buffer.WriteUint8(uint8(len(this.VipLevelArr)))
	for i := 0; i < len(this.VipLevelArr); i++ {
		this.VipLevelArr[i].Encode(buffer)
	}
}

func (this *VipLevelTotal_Out_VipLevelArr) Encode(buffer *net.Buffer) {
	buffer.WriteUint16LE(uint16(this.VipLevel))
	buffer.WriteUint32LE(uint32(this.Total))
}

func (this *VipLevelTotal_Out) ByteSize() int {
	size := 3
	size += len(this.VipLevelArr) * 6
	return size
}

func (this *BuyTimes_In) Decode(buffer *net.Buffer) {
	this.Buytype = BuyTimesType(buffer.ReadUint8())
}

func (this *BuyTimes_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(20)
	buffer.WriteUint8(4)
	buffer.WriteUint8(uint8(this.Buytype))
}

func (this *BuyTimes_In) ByteSize() int {
	size := 3
	return size
}

func (this *BuyTimes_Out) Decode(buffer *net.Buffer) {
	this.Result = buffer.ReadUint8() == 1
}

func (this *BuyTimes_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(20)
	buffer.WriteUint8(4)
	if this.Result {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *BuyTimes_Out) ByteSize() int {
	size := 3
	return size
}
