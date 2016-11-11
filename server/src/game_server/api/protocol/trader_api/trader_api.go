package trader_api

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
	StoreState(*net.Session, *StoreState_In)
	Buy(*net.Session, *Buy_In)
	Refresh(*net.Session, *Refresh_In)
}

type OutHandler interface {
	Info(*net.Session, *Info_Out)
	StoreState(*net.Session, *StoreState_Out)
	Buy(*net.Session, *Buy_Out)
	Refresh(*net.Session, *Refresh_Out)
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
		request := new(StoreState_In)
		request.Decode(buffer)
		return request
	case 3:
		request := new(Buy_In)
		request.Decode(buffer)
		return request
	case 4:
		request := new(Refresh_In)
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
		request := new(StoreState_Out)
		request.Decode(buffer)
		return request
	case 3:
		request := new(Buy_Out)
		request.Decode(buffer)
		return request
	case 4:
		request := new(Refresh_Out)
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
	return "trader.info.in"
}

func (this *Info_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 21, 1
}

type Info_Out struct {
	During []Info_Out_During `json:"during"`
}

type Info_Out_During struct {
	Expire    int64 `json:"expire"`
	Showup    int64 `json:"showup"`
	Disappear int64 `json:"disappear"`
}

func (this *Info_Out) Process(session *net.Session) {
	g_OutHandler.Info(session, this)
}

func (this *Info_Out) TypeName() string {
	return "trader.info.out"
}

func (this *Info_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 21, 1
}

func (this *Info_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type StoreState_In struct {
	TraderId int16 `json:"trader_id"`
}

func (this *StoreState_In) Process(session *net.Session) {
	g_InHandler.StoreState(session, this)
}

func (this *StoreState_In) TypeName() string {
	return "trader.store_state.in"
}

func (this *StoreState_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 21, 2
}

type StoreState_Out struct {
	RefreshNum int16                  `json:"refresh_num"`
	Goods      []StoreState_Out_Goods `json:"goods"`
}

type StoreState_Out_Goods struct {
	GridId    int32 `json:"grid_id"`
	Cost      int64 `json:"cost"`
	GoodsType int8  `json:"goods_type"`
	ItemId    int16 `json:"item_id"`
	Num       int16 `json:"num"`
	Stock     int8  `json:"stock"`
}

func (this *StoreState_Out) Process(session *net.Session) {
	g_OutHandler.StoreState(session, this)
}

func (this *StoreState_Out) TypeName() string {
	return "trader.store_state.out"
}

func (this *StoreState_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 21, 2
}

func (this *StoreState_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type Buy_In struct {
	GridId int32 `json:"grid_id"`
}

func (this *Buy_In) Process(session *net.Session) {
	g_InHandler.Buy(session, this)
}

func (this *Buy_In) TypeName() string {
	return "trader.buy.in"
}

func (this *Buy_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 21, 3
}

type Buy_Out struct {
	Expired bool `json:"expired"`
}

func (this *Buy_Out) Process(session *net.Session) {
	g_OutHandler.Buy(session, this)
}

func (this *Buy_Out) TypeName() string {
	return "trader.buy.out"
}

func (this *Buy_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 21, 3
}

func (this *Buy_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type Refresh_In struct {
	TraderId int16 `json:"trader_id"`
}

func (this *Refresh_In) Process(session *net.Session) {
	g_InHandler.Refresh(session, this)
}

func (this *Refresh_In) TypeName() string {
	return "trader.refresh.in"
}

func (this *Refresh_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 21, 4
}

type Refresh_Out struct {
	Goods []Refresh_Out_Goods `json:"goods"`
}

type Refresh_Out_Goods struct {
	GridId    int32 `json:"grid_id"`
	GoodsType int8  `json:"goods_type"`
	Cost      int64 `json:"cost"`
	ItemId    int16 `json:"item_id"`
	Num       int16 `json:"num"`
	Stock     int8  `json:"stock"`
}

func (this *Refresh_Out) Process(session *net.Session) {
	g_OutHandler.Refresh(session, this)
}

func (this *Refresh_Out) TypeName() string {
	return "trader.refresh.out"
}

func (this *Refresh_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 21, 4
}

func (this *Refresh_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

func (this *Info_In) Decode(buffer *net.Buffer) {
}

func (this *Info_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(21)
	buffer.WriteUint8(1)
}

func (this *Info_In) ByteSize() int {
	size := 2
	return size
}

func (this *Info_Out) Decode(buffer *net.Buffer) {
	this.During = make([]Info_Out_During, buffer.ReadUint8())
	for i := 0; i < len(this.During); i++ {
		this.During[i].Decode(buffer)
	}
}

func (this *Info_Out_During) Decode(buffer *net.Buffer) {
	this.Expire = int64(buffer.ReadUint64LE())
	this.Showup = int64(buffer.ReadUint64LE())
	this.Disappear = int64(buffer.ReadUint64LE())
}

func (this *Info_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(21)
	buffer.WriteUint8(1)
	buffer.WriteUint8(uint8(len(this.During)))
	for i := 0; i < len(this.During); i++ {
		this.During[i].Encode(buffer)
	}
}

func (this *Info_Out_During) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.Expire))
	buffer.WriteUint64LE(uint64(this.Showup))
	buffer.WriteUint64LE(uint64(this.Disappear))
}

func (this *Info_Out) ByteSize() int {
	size := 3
	size += len(this.During) * 24
	return size
}

func (this *StoreState_In) Decode(buffer *net.Buffer) {
	this.TraderId = int16(buffer.ReadUint16LE())
}

func (this *StoreState_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(21)
	buffer.WriteUint8(2)
	buffer.WriteUint16LE(uint16(this.TraderId))
}

func (this *StoreState_In) ByteSize() int {
	size := 4
	return size
}

func (this *StoreState_Out) Decode(buffer *net.Buffer) {
	this.RefreshNum = int16(buffer.ReadUint16LE())
	this.Goods = make([]StoreState_Out_Goods, buffer.ReadUint8())
	for i := 0; i < len(this.Goods); i++ {
		this.Goods[i].Decode(buffer)
	}
}

func (this *StoreState_Out_Goods) Decode(buffer *net.Buffer) {
	this.GridId = int32(buffer.ReadUint32LE())
	this.Cost = int64(buffer.ReadUint64LE())
	this.GoodsType = int8(buffer.ReadUint8())
	this.ItemId = int16(buffer.ReadUint16LE())
	this.Num = int16(buffer.ReadUint16LE())
	this.Stock = int8(buffer.ReadUint8())
}

func (this *StoreState_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(21)
	buffer.WriteUint8(2)
	buffer.WriteUint16LE(uint16(this.RefreshNum))
	buffer.WriteUint8(uint8(len(this.Goods)))
	for i := 0; i < len(this.Goods); i++ {
		this.Goods[i].Encode(buffer)
	}
}

func (this *StoreState_Out_Goods) Encode(buffer *net.Buffer) {
	buffer.WriteUint32LE(uint32(this.GridId))
	buffer.WriteUint64LE(uint64(this.Cost))
	buffer.WriteUint8(uint8(this.GoodsType))
	buffer.WriteUint16LE(uint16(this.ItemId))
	buffer.WriteUint16LE(uint16(this.Num))
	buffer.WriteUint8(uint8(this.Stock))
}

func (this *StoreState_Out) ByteSize() int {
	size := 5
	size += len(this.Goods) * 18
	return size
}

func (this *Buy_In) Decode(buffer *net.Buffer) {
	this.GridId = int32(buffer.ReadUint32LE())
}

func (this *Buy_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(21)
	buffer.WriteUint8(3)
	buffer.WriteUint32LE(uint32(this.GridId))
}

func (this *Buy_In) ByteSize() int {
	size := 6
	return size
}

func (this *Buy_Out) Decode(buffer *net.Buffer) {
	this.Expired = buffer.ReadUint8() == 1
}

func (this *Buy_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(21)
	buffer.WriteUint8(3)
	if this.Expired {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *Buy_Out) ByteSize() int {
	size := 3
	return size
}

func (this *Refresh_In) Decode(buffer *net.Buffer) {
	this.TraderId = int16(buffer.ReadUint16LE())
}

func (this *Refresh_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(21)
	buffer.WriteUint8(4)
	buffer.WriteUint16LE(uint16(this.TraderId))
}

func (this *Refresh_In) ByteSize() int {
	size := 4
	return size
}

func (this *Refresh_Out) Decode(buffer *net.Buffer) {
	this.Goods = make([]Refresh_Out_Goods, buffer.ReadUint8())
	for i := 0; i < len(this.Goods); i++ {
		this.Goods[i].Decode(buffer)
	}
}

func (this *Refresh_Out_Goods) Decode(buffer *net.Buffer) {
	this.GridId = int32(buffer.ReadUint32LE())
	this.GoodsType = int8(buffer.ReadUint8())
	this.Cost = int64(buffer.ReadUint64LE())
	this.ItemId = int16(buffer.ReadUint16LE())
	this.Num = int16(buffer.ReadUint16LE())
	this.Stock = int8(buffer.ReadUint8())
}

func (this *Refresh_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(21)
	buffer.WriteUint8(4)
	buffer.WriteUint8(uint8(len(this.Goods)))
	for i := 0; i < len(this.Goods); i++ {
		this.Goods[i].Encode(buffer)
	}
}

func (this *Refresh_Out_Goods) Encode(buffer *net.Buffer) {
	buffer.WriteUint32LE(uint32(this.GridId))
	buffer.WriteUint8(uint8(this.GoodsType))
	buffer.WriteUint64LE(uint64(this.Cost))
	buffer.WriteUint16LE(uint16(this.ItemId))
	buffer.WriteUint16LE(uint16(this.Num))
	buffer.WriteUint8(uint8(this.Stock))
}

func (this *Refresh_Out) ByteSize() int {
	size := 3
	size += len(this.Goods) * 18
	return size
}
