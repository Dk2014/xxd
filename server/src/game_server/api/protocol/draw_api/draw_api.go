package draw_api

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
	GetHeartDrawInfo(*net.Session, *GetHeartDrawInfo_In)
	HeartDraw(*net.Session, *HeartDraw_In)
	GetChestInfo(*net.Session, *GetChestInfo_In)
	DrawChest(*net.Session, *DrawChest_In)
	HeartInfo(*net.Session, *HeartInfo_In)
	GetFateBoxInfo(*net.Session, *GetFateBoxInfo_In)
	OpenFateBox(*net.Session, *OpenFateBox_In)
	ExchangeGiftCode(*net.Session, *ExchangeGiftCode_In)
}

type OutHandler interface {
	GetHeartDrawInfo(*net.Session, *GetHeartDrawInfo_Out)
	HeartDraw(*net.Session, *HeartDraw_Out)
	GetChestInfo(*net.Session, *GetChestInfo_Out)
	DrawChest(*net.Session, *DrawChest_Out)
	HeartInfo(*net.Session, *HeartInfo_Out)
	GetFateBoxInfo(*net.Session, *GetFateBoxInfo_Out)
	OpenFateBox(*net.Session, *OpenFateBox_Out)
	ExchangeGiftCode(*net.Session, *ExchangeGiftCode_Out)
}

func DecodeIn(message []byte) Request {
	var actionId = message[0]
	var buffer = net.NewBuffer(message[1:])
	switch actionId {
	case 1:
		request := new(GetHeartDrawInfo_In)
		request.Decode(buffer)
		return request
	case 2:
		request := new(HeartDraw_In)
		request.Decode(buffer)
		return request
	case 3:
		request := new(GetChestInfo_In)
		request.Decode(buffer)
		return request
	case 4:
		request := new(DrawChest_In)
		request.Decode(buffer)
		return request
	case 5:
		request := new(HeartInfo_In)
		request.Decode(buffer)
		return request
	case 6:
		request := new(GetFateBoxInfo_In)
		request.Decode(buffer)
		return request
	case 7:
		request := new(OpenFateBox_In)
		request.Decode(buffer)
		return request
	case 8:
		request := new(ExchangeGiftCode_In)
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
		request := new(GetHeartDrawInfo_Out)
		request.Decode(buffer)
		return request
	case 2:
		request := new(HeartDraw_Out)
		request.Decode(buffer)
		return request
	case 3:
		request := new(GetChestInfo_Out)
		request.Decode(buffer)
		return request
	case 4:
		request := new(DrawChest_Out)
		request.Decode(buffer)
		return request
	case 5:
		request := new(HeartInfo_Out)
		request.Decode(buffer)
		return request
	case 6:
		request := new(GetFateBoxInfo_Out)
		request.Decode(buffer)
		return request
	case 7:
		request := new(OpenFateBox_Out)
		request.Decode(buffer)
		return request
	case 8:
		request := new(ExchangeGiftCode_Out)
		request.Decode(buffer)
		return request
	}
	_ = buffer
	panic("unsupported response")
}

type ExchangeGiftResult int8

const (
	EXCHANGE_GIFT_RESULT_SUCCESS      ExchangeGiftResult = 0
	EXCHANGE_GIFT_RESULT_EXPIRE       ExchangeGiftResult = 1
	EXCHANGE_GIFT_RESULT_DUP_EXCHANGE ExchangeGiftResult = 2
)

type ChestType int8

const (
	CHEST_TYPE_COIN       ChestType = 0
	CHEST_TYPE_COIN_FREE  ChestType = 1
	CHEST_TYPE_COIN_TEN   ChestType = 2
	CHEST_TYPE_INGOT      ChestType = 3
	CHEST_TYPE_INGOT_FREE ChestType = 4
	CHEST_TYPE_INGOT_TEN  ChestType = 5
	CHEST_TYPE_PET        ChestType = 6
	CHEST_TYPE_PET_FREE   ChestType = 7
	CHEST_TYPE_PET_TEN    ChestType = 8
)

type ItemType int8

const (
	ITEM_TYPE_COIN           ItemType = 1
	ITEM_TYPE_INGOT          ItemType = 2
	ITEM_TYPE_ITEM           ItemType = 3
	ITEM_TYPE_GHOST          ItemType = 4
	ITEM_TYPE_SWORD_SOUL     ItemType = 5
	ITEM_TYPE_PET            ItemType = 6
	ITEM_TYPE_GHOST_FRAGMENT ItemType = 7
	ITEM_TYPE_PREFERENCE     ItemType = 8
	ITEM_TYPE_EQUIPMENT      ItemType = 9
)

type AwardInfo struct {
	AwardType  int8  `json:"award_type"`
	AwardNum   int16 `json:"award_num"`
	ItemId     int16 `json:"item_id"`
	DrawTime   int64 `json:"draw_time"`
	AwardIndex int16 `json:"award_index"`
}

func (this *AwardInfo) Decode(buffer *net.Buffer) {
	this.AwardType = int8(buffer.ReadUint8())
	this.AwardNum = int16(buffer.ReadUint16LE())
	this.ItemId = int16(buffer.ReadUint16LE())
	this.DrawTime = int64(buffer.ReadUint64LE())
	this.AwardIndex = int16(buffer.ReadUint16LE())
}

func (this *AwardInfo) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(this.AwardType))
	buffer.WriteUint16LE(uint16(this.AwardNum))
	buffer.WriteUint16LE(uint16(this.ItemId))
	buffer.WriteUint64LE(uint64(this.DrawTime))
	buffer.WriteUint16LE(uint16(this.AwardIndex))
}

func (this *AwardInfo) ByteSize() int {
	size := 15
	return size
}

type GetHeartDrawInfo_In struct {
	DrawType    int8 `json:"draw_type"`
	AwardRecord bool `json:"award_record"`
}

func (this *GetHeartDrawInfo_In) Process(session *net.Session) {
	g_InHandler.GetHeartDrawInfo(session, this)
}

func (this *GetHeartDrawInfo_In) TypeName() string {
	return "draw.get_heart_draw_info.in"
}

func (this *GetHeartDrawInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 50, 1
}

type GetHeartDrawInfo_Out struct {
	Hearts      int16                              `json:"hearts"`
	DailyNum    int8                               `json:"daily_num"`
	AwardRecord []GetHeartDrawInfo_Out_AwardRecord `json:"award_record"`
}

type GetHeartDrawInfo_Out_AwardRecord struct {
	Award AwardInfo `json:"award"`
}

func (this *GetHeartDrawInfo_Out) Process(session *net.Session) {
	g_OutHandler.GetHeartDrawInfo(session, this)
}

func (this *GetHeartDrawInfo_Out) TypeName() string {
	return "draw.get_heart_draw_info.out"
}

func (this *GetHeartDrawInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 50, 1
}

func (this *GetHeartDrawInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type HeartDraw_In struct {
	DrawType int8 `json:"draw_type"`
}

func (this *HeartDraw_In) Process(session *net.Session) {
	g_InHandler.HeartDraw(session, this)
}

func (this *HeartDraw_In) TypeName() string {
	return "draw.heart_draw.in"
}

func (this *HeartDraw_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 50, 2
}

type HeartDraw_Out struct {
	Award AwardInfo `json:"award"`
}

func (this *HeartDraw_Out) Process(session *net.Session) {
	g_OutHandler.HeartDraw(session, this)
}

func (this *HeartDraw_Out) TypeName() string {
	return "draw.heart_draw.out"
}

func (this *HeartDraw_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 50, 2
}

func (this *HeartDraw_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetChestInfo_In struct {
}

func (this *GetChestInfo_In) Process(session *net.Session) {
	g_InHandler.GetChestInfo(session, this)
}

func (this *GetChestInfo_In) TypeName() string {
	return "draw.get_chest_info.in"
}

func (this *GetChestInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 50, 3
}

type GetChestInfo_Out struct {
	FreeCoinNum       int32 `json:"free_coin_num"`
	NextFreeCoinLeft  int64 `json:"next_free_coin_left"`
	NextFreeIngotLeft int64 `json:"next_free_ingot_left"`
	NextFreePetLeft   int64 `json:"next_free_pet_left"`
}

func (this *GetChestInfo_Out) Process(session *net.Session) {
	g_OutHandler.GetChestInfo(session, this)
}

func (this *GetChestInfo_Out) TypeName() string {
	return "draw.get_chest_info.out"
}

func (this *GetChestInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 50, 3
}

func (this *GetChestInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type DrawChest_In struct {
	ChestType ChestType `json:"chest_type"`
}

func (this *DrawChest_In) Process(session *net.Session) {
	g_InHandler.DrawChest(session, this)
}

func (this *DrawChest_In) TypeName() string {
	return "draw.draw_chest.in"
}

func (this *DrawChest_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 50, 4
}

type DrawChest_Out struct {
	Items []DrawChest_Out_Items `json:"items"`
}

type DrawChest_Out_Items struct {
	ItemType ItemType `json:"item_type"`
	ItemId   int16    `json:"item_id"`
	Num      int32    `json:"num"`
}

func (this *DrawChest_Out) Process(session *net.Session) {
	g_OutHandler.DrawChest(session, this)
}

func (this *DrawChest_Out) TypeName() string {
	return "draw.draw_chest.out"
}

func (this *DrawChest_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 50, 4
}

func (this *DrawChest_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type HeartInfo_In struct {
}

func (this *HeartInfo_In) Process(session *net.Session) {
	g_InHandler.HeartInfo(session, this)
}

func (this *HeartInfo_In) TypeName() string {
	return "draw.heart_info.in"
}

func (this *HeartInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 50, 5
}

type HeartInfo_Out struct {
	RecoverToday int16 `json:"recover_today"`
	Timestamp    int64 `json:"timestamp"`
}

func (this *HeartInfo_Out) Process(session *net.Session) {
	g_OutHandler.HeartInfo(session, this)
}

func (this *HeartInfo_Out) TypeName() string {
	return "draw.heart_info.out"
}

func (this *HeartInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 50, 5
}

func (this *HeartInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetFateBoxInfo_In struct {
}

func (this *GetFateBoxInfo_In) Process(session *net.Session) {
	g_InHandler.GetFateBoxInfo(session, this)
}

func (this *GetFateBoxInfo_In) TypeName() string {
	return "draw.get_fate_box_info.in"
}

func (this *GetFateBoxInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 50, 6
}

type GetFateBoxInfo_Out struct {
	Lock               int32 `json:"lock"`
	NextFreeStarBox    int32 `json:"next_free_star_box"`
	NextFreeMoonBox    int32 `json:"next_free_moon_box"`
	NextFreeSunBox     int32 `json:"next_free_sun_box"`
	NextFreeHunyuanBox int32 `json:"next_free_hunyuan_box"`
}

func (this *GetFateBoxInfo_Out) Process(session *net.Session) {
	g_OutHandler.GetFateBoxInfo(session, this)
}

func (this *GetFateBoxInfo_Out) TypeName() string {
	return "draw.get_fate_box_info.out"
}

func (this *GetFateBoxInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 50, 6
}

func (this *GetFateBoxInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type OpenFateBox_In struct {
	BoxType int32 `json:"box_type"`
	Times   int8  `json:"times"`
}

func (this *OpenFateBox_In) Process(session *net.Session) {
	g_InHandler.OpenFateBox(session, this)
}

func (this *OpenFateBox_In) TypeName() string {
	return "draw.open_fate_box.in"
}

func (this *OpenFateBox_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 50, 7
}

type OpenFateBox_Out struct {
	Items []OpenFateBox_Out_Items `json:"items"`
}

type OpenFateBox_Out_Items struct {
	ItemType ItemType `json:"item_type"`
	ItemId   int16    `json:"item_id"`
	Num      int32    `json:"num"`
}

func (this *OpenFateBox_Out) Process(session *net.Session) {
	g_OutHandler.OpenFateBox(session, this)
}

func (this *OpenFateBox_Out) TypeName() string {
	return "draw.open_fate_box.out"
}

func (this *OpenFateBox_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 50, 7
}

func (this *OpenFateBox_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ExchangeGiftCode_In struct {
	Code []byte `json:"code"`
}

func (this *ExchangeGiftCode_In) Process(session *net.Session) {
	g_InHandler.ExchangeGiftCode(session, this)
}

func (this *ExchangeGiftCode_In) TypeName() string {
	return "draw.exchange_gift_code.in"
}

func (this *ExchangeGiftCode_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 50, 8
}

type ExchangeGiftCode_Out struct {
	Result ExchangeGiftResult `json:"result"`
}

func (this *ExchangeGiftCode_Out) Process(session *net.Session) {
	g_OutHandler.ExchangeGiftCode(session, this)
}

func (this *ExchangeGiftCode_Out) TypeName() string {
	return "draw.exchange_gift_code.out"
}

func (this *ExchangeGiftCode_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 50, 8
}

func (this *ExchangeGiftCode_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

func (this *GetHeartDrawInfo_In) Decode(buffer *net.Buffer) {
	this.DrawType = int8(buffer.ReadUint8())
	this.AwardRecord = buffer.ReadUint8() == 1
}

func (this *GetHeartDrawInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(50)
	buffer.WriteUint8(1)
	buffer.WriteUint8(uint8(this.DrawType))
	if this.AwardRecord {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *GetHeartDrawInfo_In) ByteSize() int {
	size := 4
	return size
}

func (this *GetHeartDrawInfo_Out) Decode(buffer *net.Buffer) {
	this.Hearts = int16(buffer.ReadUint16LE())
	this.DailyNum = int8(buffer.ReadUint8())
	this.AwardRecord = make([]GetHeartDrawInfo_Out_AwardRecord, buffer.ReadUint8())
	for i := 0; i < len(this.AwardRecord); i++ {
		this.AwardRecord[i].Decode(buffer)
	}
}

func (this *GetHeartDrawInfo_Out_AwardRecord) Decode(buffer *net.Buffer) {
	this.Award.Decode(buffer)
}

func (this *GetHeartDrawInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(50)
	buffer.WriteUint8(1)
	buffer.WriteUint16LE(uint16(this.Hearts))
	buffer.WriteUint8(uint8(this.DailyNum))
	buffer.WriteUint8(uint8(len(this.AwardRecord)))
	for i := 0; i < len(this.AwardRecord); i++ {
		this.AwardRecord[i].Encode(buffer)
	}
}

func (this *GetHeartDrawInfo_Out_AwardRecord) Encode(buffer *net.Buffer) {
	this.Award.Encode(buffer)
}

func (this *GetHeartDrawInfo_Out) ByteSize() int {
	size := 6
	for i := 0; i < len(this.AwardRecord); i++ {
		size += this.AwardRecord[i].ByteSize()
	}
	return size
}

func (this *GetHeartDrawInfo_Out_AwardRecord) ByteSize() int {
	size := 0
	size += this.Award.ByteSize()
	return size
}

func (this *HeartDraw_In) Decode(buffer *net.Buffer) {
	this.DrawType = int8(buffer.ReadUint8())
}

func (this *HeartDraw_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(50)
	buffer.WriteUint8(2)
	buffer.WriteUint8(uint8(this.DrawType))
}

func (this *HeartDraw_In) ByteSize() int {
	size := 3
	return size
}

func (this *HeartDraw_Out) Decode(buffer *net.Buffer) {
	this.Award.Decode(buffer)
}

func (this *HeartDraw_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(50)
	buffer.WriteUint8(2)
	this.Award.Encode(buffer)
}

func (this *HeartDraw_Out) ByteSize() int {
	size := 2
	size += this.Award.ByteSize()
	return size
}

func (this *GetChestInfo_In) Decode(buffer *net.Buffer) {
}

func (this *GetChestInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(50)
	buffer.WriteUint8(3)
}

func (this *GetChestInfo_In) ByteSize() int {
	size := 2
	return size
}

func (this *GetChestInfo_Out) Decode(buffer *net.Buffer) {
	this.FreeCoinNum = int32(buffer.ReadUint32LE())
	this.NextFreeCoinLeft = int64(buffer.ReadUint64LE())
	this.NextFreeIngotLeft = int64(buffer.ReadUint64LE())
	this.NextFreePetLeft = int64(buffer.ReadUint64LE())
}

func (this *GetChestInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(50)
	buffer.WriteUint8(3)
	buffer.WriteUint32LE(uint32(this.FreeCoinNum))
	buffer.WriteUint64LE(uint64(this.NextFreeCoinLeft))
	buffer.WriteUint64LE(uint64(this.NextFreeIngotLeft))
	buffer.WriteUint64LE(uint64(this.NextFreePetLeft))
}

func (this *GetChestInfo_Out) ByteSize() int {
	size := 30
	return size
}

func (this *DrawChest_In) Decode(buffer *net.Buffer) {
	this.ChestType = ChestType(buffer.ReadUint8())
}

func (this *DrawChest_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(50)
	buffer.WriteUint8(4)
	buffer.WriteUint8(uint8(this.ChestType))
}

func (this *DrawChest_In) ByteSize() int {
	size := 3
	return size
}

func (this *DrawChest_Out) Decode(buffer *net.Buffer) {
	this.Items = make([]DrawChest_Out_Items, buffer.ReadUint8())
	for i := 0; i < len(this.Items); i++ {
		this.Items[i].Decode(buffer)
	}
}

func (this *DrawChest_Out_Items) Decode(buffer *net.Buffer) {
	this.ItemType = ItemType(buffer.ReadUint8())
	this.ItemId = int16(buffer.ReadUint16LE())
	this.Num = int32(buffer.ReadUint32LE())
}

func (this *DrawChest_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(50)
	buffer.WriteUint8(4)
	buffer.WriteUint8(uint8(len(this.Items)))
	for i := 0; i < len(this.Items); i++ {
		this.Items[i].Encode(buffer)
	}
}

func (this *DrawChest_Out_Items) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(this.ItemType))
	buffer.WriteUint16LE(uint16(this.ItemId))
	buffer.WriteUint32LE(uint32(this.Num))
}

func (this *DrawChest_Out) ByteSize() int {
	size := 3
	size += len(this.Items) * 7
	return size
}

func (this *HeartInfo_In) Decode(buffer *net.Buffer) {
}

func (this *HeartInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(50)
	buffer.WriteUint8(5)
}

func (this *HeartInfo_In) ByteSize() int {
	size := 2
	return size
}

func (this *HeartInfo_Out) Decode(buffer *net.Buffer) {
	this.RecoverToday = int16(buffer.ReadUint16LE())
	this.Timestamp = int64(buffer.ReadUint64LE())
}

func (this *HeartInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(50)
	buffer.WriteUint8(5)
	buffer.WriteUint16LE(uint16(this.RecoverToday))
	buffer.WriteUint64LE(uint64(this.Timestamp))
}

func (this *HeartInfo_Out) ByteSize() int {
	size := 12
	return size
}

func (this *GetFateBoxInfo_In) Decode(buffer *net.Buffer) {
}

func (this *GetFateBoxInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(50)
	buffer.WriteUint8(6)
}

func (this *GetFateBoxInfo_In) ByteSize() int {
	size := 2
	return size
}

func (this *GetFateBoxInfo_Out) Decode(buffer *net.Buffer) {
	this.Lock = int32(buffer.ReadUint32LE())
	this.NextFreeStarBox = int32(buffer.ReadUint32LE())
	this.NextFreeMoonBox = int32(buffer.ReadUint32LE())
	this.NextFreeSunBox = int32(buffer.ReadUint32LE())
	this.NextFreeHunyuanBox = int32(buffer.ReadUint32LE())
}

func (this *GetFateBoxInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(50)
	buffer.WriteUint8(6)
	buffer.WriteUint32LE(uint32(this.Lock))
	buffer.WriteUint32LE(uint32(this.NextFreeStarBox))
	buffer.WriteUint32LE(uint32(this.NextFreeMoonBox))
	buffer.WriteUint32LE(uint32(this.NextFreeSunBox))
	buffer.WriteUint32LE(uint32(this.NextFreeHunyuanBox))
}

func (this *GetFateBoxInfo_Out) ByteSize() int {
	size := 22
	return size
}

func (this *OpenFateBox_In) Decode(buffer *net.Buffer) {
	this.BoxType = int32(buffer.ReadUint32LE())
	this.Times = int8(buffer.ReadUint8())
}

func (this *OpenFateBox_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(50)
	buffer.WriteUint8(7)
	buffer.WriteUint32LE(uint32(this.BoxType))
	buffer.WriteUint8(uint8(this.Times))
}

func (this *OpenFateBox_In) ByteSize() int {
	size := 7
	return size
}

func (this *OpenFateBox_Out) Decode(buffer *net.Buffer) {
	this.Items = make([]OpenFateBox_Out_Items, buffer.ReadUint8())
	for i := 0; i < len(this.Items); i++ {
		this.Items[i].Decode(buffer)
	}
}

func (this *OpenFateBox_Out_Items) Decode(buffer *net.Buffer) {
	this.ItemType = ItemType(buffer.ReadUint8())
	this.ItemId = int16(buffer.ReadUint16LE())
	this.Num = int32(buffer.ReadUint32LE())
}

func (this *OpenFateBox_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(50)
	buffer.WriteUint8(7)
	buffer.WriteUint8(uint8(len(this.Items)))
	for i := 0; i < len(this.Items); i++ {
		this.Items[i].Encode(buffer)
	}
}

func (this *OpenFateBox_Out_Items) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(this.ItemType))
	buffer.WriteUint16LE(uint16(this.ItemId))
	buffer.WriteUint32LE(uint32(this.Num))
}

func (this *OpenFateBox_Out) ByteSize() int {
	size := 3
	size += len(this.Items) * 7
	return size
}

func (this *ExchangeGiftCode_In) Decode(buffer *net.Buffer) {
	this.Code = buffer.ReadBytes(int(buffer.ReadUint16LE()))
}

func (this *ExchangeGiftCode_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(50)
	buffer.WriteUint8(8)
	buffer.WriteUint16LE(uint16(len(this.Code)))
	buffer.WriteBytes(this.Code)
}

func (this *ExchangeGiftCode_In) ByteSize() int {
	size := 4
	size += len(this.Code)
	return size
}

func (this *ExchangeGiftCode_Out) Decode(buffer *net.Buffer) {
	this.Result = ExchangeGiftResult(buffer.ReadUint8())
}

func (this *ExchangeGiftCode_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(50)
	buffer.WriteUint8(8)
	buffer.WriteUint8(uint8(this.Result))
}

func (this *ExchangeGiftCode_Out) ByteSize() int {
	size := 3
	return size
}
