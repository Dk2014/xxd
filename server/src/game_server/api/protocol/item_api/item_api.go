package item_api

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
	GetAllItems(*net.Session, *GetAllItems_In)
	DropItem(*net.Session, *DropItem_In)
	BuyItem(*net.Session, *BuyItem_In)
	SellItem(*net.Session, *SellItem_In)
	Dress(*net.Session, *Dress_In)
	Undress(*net.Session, *Undress_In)
	BuyItemBack(*net.Session, *BuyItemBack_In)
	IsBagFull(*net.Session, *IsBagFull_In)
	Decompose(*net.Session, *Decompose_In)
	Refine(*net.Session, *Refine_In)
	GetRecastInfo(*net.Session, *GetRecastInfo_In)
	Recast(*net.Session, *Recast_In)
	UseItem(*net.Session, *UseItem_In)
	RoleUseCostItem(*net.Session, *RoleUseCostItem_In)
	BatchUseItem(*net.Session, *BatchUseItem_In)
	OpenCornucopia(*net.Session, *OpenCornucopia_In)
	GetSealedbooks(*net.Session, *GetSealedbooks_In)
	ActivationSealedbook(*net.Session, *ActivationSealedbook_In)
	ExchangeGhostCrystal(*net.Session, *ExchangeGhostCrystal_In)
	ExchangeShopItem(*net.Session, *ExchangeShopItem_In)
	GetHoildayItemList(*net.Session, *GetHoildayItemList_In)
	ExchangeHoildayItem(*net.Session, *ExchangeHoildayItem_In)
	BatchExchangeDragonBall(*net.Session, *BatchExchangeDragonBall_In)
}

type OutHandler interface {
	GetAllItems(*net.Session, *GetAllItems_Out)
	DropItem(*net.Session, *DropItem_Out)
	BuyItem(*net.Session, *BuyItem_Out)
	SellItem(*net.Session, *SellItem_Out)
	Dress(*net.Session, *Dress_Out)
	Undress(*net.Session, *Undress_Out)
	BuyItemBack(*net.Session, *BuyItemBack_Out)
	IsBagFull(*net.Session, *IsBagFull_Out)
	Decompose(*net.Session, *Decompose_Out)
	Refine(*net.Session, *Refine_Out)
	GetRecastInfo(*net.Session, *GetRecastInfo_Out)
	Recast(*net.Session, *Recast_Out)
	UseItem(*net.Session, *UseItem_Out)
	RoleUseCostItem(*net.Session, *RoleUseCostItem_Out)
	BatchUseItem(*net.Session, *BatchUseItem_Out)
	DragonBallExchangeNotify(*net.Session, *DragonBallExchangeNotify_Out)
	OpenCornucopia(*net.Session, *OpenCornucopia_Out)
	GetSealedbooks(*net.Session, *GetSealedbooks_Out)
	ActivationSealedbook(*net.Session, *ActivationSealedbook_Out)
	ExchangeGhostCrystal(*net.Session, *ExchangeGhostCrystal_Out)
	ExchangeShopItem(*net.Session, *ExchangeShopItem_Out)
	GetHoildayItemList(*net.Session, *GetHoildayItemList_Out)
	ExchangeHoildayItem(*net.Session, *ExchangeHoildayItem_Out)
	BatchExchangeDragonBall(*net.Session, *BatchExchangeDragonBall_Out)
}

func DecodeIn(message []byte) Request {
	var actionId = message[0]
	var buffer = net.NewBuffer(message[1:])
	switch actionId {
	case 0:
		request := new(GetAllItems_In)
		request.Decode(buffer)
		return request
	case 1:
		request := new(DropItem_In)
		request.Decode(buffer)
		return request
	case 2:
		request := new(BuyItem_In)
		request.Decode(buffer)
		return request
	case 3:
		request := new(SellItem_In)
		request.Decode(buffer)
		return request
	case 4:
		request := new(Dress_In)
		request.Decode(buffer)
		return request
	case 5:
		request := new(Undress_In)
		request.Decode(buffer)
		return request
	case 6:
		request := new(BuyItemBack_In)
		request.Decode(buffer)
		return request
	case 7:
		request := new(IsBagFull_In)
		request.Decode(buffer)
		return request
	case 8:
		request := new(Decompose_In)
		request.Decode(buffer)
		return request
	case 9:
		request := new(Refine_In)
		request.Decode(buffer)
		return request
	case 10:
		request := new(GetRecastInfo_In)
		request.Decode(buffer)
		return request
	case 11:
		request := new(Recast_In)
		request.Decode(buffer)
		return request
	case 12:
		request := new(UseItem_In)
		request.Decode(buffer)
		return request
	case 13:
		request := new(RoleUseCostItem_In)
		request.Decode(buffer)
		return request
	case 14:
		request := new(BatchUseItem_In)
		request.Decode(buffer)
		return request
	case 16:
		request := new(OpenCornucopia_In)
		request.Decode(buffer)
		return request
	case 17:
		request := new(GetSealedbooks_In)
		request.Decode(buffer)
		return request
	case 18:
		request := new(ActivationSealedbook_In)
		request.Decode(buffer)
		return request
	case 19:
		request := new(ExchangeGhostCrystal_In)
		request.Decode(buffer)
		return request
	case 20:
		request := new(ExchangeShopItem_In)
		request.Decode(buffer)
		return request
	case 22:
		request := new(GetHoildayItemList_In)
		request.Decode(buffer)
		return request
	case 23:
		request := new(ExchangeHoildayItem_In)
		request.Decode(buffer)
		return request
	case 24:
		request := new(BatchExchangeDragonBall_In)
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
		request := new(GetAllItems_Out)
		request.Decode(buffer)
		return request
	case 1:
		request := new(DropItem_Out)
		request.Decode(buffer)
		return request
	case 2:
		request := new(BuyItem_Out)
		request.Decode(buffer)
		return request
	case 3:
		request := new(SellItem_Out)
		request.Decode(buffer)
		return request
	case 4:
		request := new(Dress_Out)
		request.Decode(buffer)
		return request
	case 5:
		request := new(Undress_Out)
		request.Decode(buffer)
		return request
	case 6:
		request := new(BuyItemBack_Out)
		request.Decode(buffer)
		return request
	case 7:
		request := new(IsBagFull_Out)
		request.Decode(buffer)
		return request
	case 8:
		request := new(Decompose_Out)
		request.Decode(buffer)
		return request
	case 9:
		request := new(Refine_Out)
		request.Decode(buffer)
		return request
	case 10:
		request := new(GetRecastInfo_Out)
		request.Decode(buffer)
		return request
	case 11:
		request := new(Recast_Out)
		request.Decode(buffer)
		return request
	case 12:
		request := new(UseItem_Out)
		request.Decode(buffer)
		return request
	case 13:
		request := new(RoleUseCostItem_Out)
		request.Decode(buffer)
		return request
	case 14:
		request := new(BatchUseItem_Out)
		request.Decode(buffer)
		return request
	case 15:
		request := new(DragonBallExchangeNotify_Out)
		request.Decode(buffer)
		return request
	case 16:
		request := new(OpenCornucopia_Out)
		request.Decode(buffer)
		return request
	case 17:
		request := new(GetSealedbooks_Out)
		request.Decode(buffer)
		return request
	case 18:
		request := new(ActivationSealedbook_Out)
		request.Decode(buffer)
		return request
	case 19:
		request := new(ExchangeGhostCrystal_Out)
		request.Decode(buffer)
		return request
	case 20:
		request := new(ExchangeShopItem_Out)
		request.Decode(buffer)
		return request
	case 22:
		request := new(GetHoildayItemList_Out)
		request.Decode(buffer)
		return request
	case 23:
		request := new(ExchangeHoildayItem_Out)
		request.Decode(buffer)
		return request
	case 24:
		request := new(BatchExchangeDragonBall_Out)
		request.Decode(buffer)
		return request
	}
	_ = buffer
	panic("unsupported response")
}

type EquipmentPos int8

const (
	EQUIPMENT_POS_WEAPON      EquipmentPos = 0
	EQUIPMENT_POS_CLOTHES     EquipmentPos = 1
	EQUIPMENT_POS_ACCESSORIES EquipmentPos = 2
	EQUIPMENT_POS_SHOE        EquipmentPos = 3
)

type Attribute int8

const (
	ATTRIBUTE_NULL           Attribute = 0
	ATTRIBUTE_ATTACK         Attribute = 1
	ATTRIBUTE_DEFENCE        Attribute = 2
	ATTRIBUTE_HEALTH         Attribute = 3
	ATTRIBUTE_SPEED          Attribute = 4
	ATTRIBUTE_CULTIVATION    Attribute = 5
	ATTRIBUTE_HIT_LEVEL      Attribute = 6
	ATTRIBUTE_CRITICAL_LEVEL Attribute = 7
	ATTRIBUTE_BLOCK_LEVEL    Attribute = 8
	ATTRIBUTE_DESTROY_LEVEL  Attribute = 9
	ATTRIBUTE_TENACITY_LEVEL Attribute = 10
	ATTRIBUTE_DODGE_LEVEL    Attribute = 11
	ATTRIBUTE_NUM            Attribute = 11
)

type GetAllItems_In struct {
}

func (this *GetAllItems_In) Process(session *net.Session) {
	g_InHandler.GetAllItems(session, this)
}

func (this *GetAllItems_In) TypeName() string {
	return "item.get_all_items.in"
}

func (this *GetAllItems_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 7, 0
}

type GetAllItems_Out struct {
	Items      []GetAllItems_Out_Items      `json:"items"`
	Equipments []GetAllItems_Out_Equipments `json:"equipments"`
	Buybacks   []GetAllItems_Out_Buybacks   `json:"buybacks"`
	BuyRecords []GetAllItems_Out_BuyRecords `json:"buy_records"`
}

type GetAllItems_Out_Items struct {
	Id              int64     `json:"id"`
	ItemId          int16     `json:"item_id"`
	Num             int16     `json:"num"`
	Attack          int32     `json:"attack"`
	Defence         int32     `json:"defence"`
	Health          int32     `json:"health"`
	Speed           int32     `json:"speed"`
	Cultivation     int32     `json:"cultivation"`
	HitLevel        int32     `json:"hit_level"`
	CriticalLevel   int32     `json:"critical_level"`
	BlockLevel      int32     `json:"block_level"`
	DestroyLevel    int32     `json:"destroy_level"`
	TenacityLevel   int32     `json:"tenacity_level"`
	DodgeLevel      int32     `json:"dodge_level"`
	RefineLevel     int16     `json:"refine_level"`
	RefineFailTimes int16     `json:"refine_fail_times"`
	RecastAttr      Attribute `json:"recast_attr"`
}

type GetAllItems_Out_Equipments struct {
	RoleId int8                                `json:"role_id"`
	Equips []GetAllItems_Out_Equipments_Equips `json:"equips"`
}

type GetAllItems_Out_Equipments_Equips struct {
	Id              int64     `json:"id"`
	ItemId          int16     `json:"item_id"`
	Attack          int32     `json:"attack"`
	Defence         int32     `json:"defence"`
	Health          int32     `json:"health"`
	Speed           int32     `json:"speed"`
	Cultivation     int32     `json:"cultivation"`
	HitLevel        int32     `json:"hit_level"`
	CriticalLevel   int32     `json:"critical_level"`
	BlockLevel      int32     `json:"block_level"`
	DestroyLevel    int32     `json:"destroy_level"`
	TenacityLevel   int32     `json:"tenacity_level"`
	DodgeLevel      int32     `json:"dodge_level"`
	RefineLevel     int16     `json:"refine_level"`
	RefineFailTimes int16     `json:"refine_fail_times"`
	RecastAttr      Attribute `json:"recast_attr"`
}

type GetAllItems_Out_Buybacks struct {
	Id          int64     `json:"id"`
	ItemId      int16     `json:"item_id"`
	Num         int16     `json:"num"`
	RefineLevel int16     `json:"refine_level"`
	RecastAttr  Attribute `json:"recast_attr"`
}

type GetAllItems_Out_BuyRecords struct {
	ItemId int16 `json:"item_id"`
	Num    int16 `json:"num"`
}

func (this *GetAllItems_Out) Process(session *net.Session) {
	g_OutHandler.GetAllItems(session, this)
}

func (this *GetAllItems_Out) TypeName() string {
	return "item.get_all_items.out"
}

func (this *GetAllItems_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 7, 0
}

func (this *GetAllItems_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type DropItem_In struct {
	Id int64 `json:"id"`
}

func (this *DropItem_In) Process(session *net.Session) {
	g_InHandler.DropItem(session, this)
}

func (this *DropItem_In) TypeName() string {
	return "item.drop_item.in"
}

func (this *DropItem_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 7, 1
}

type DropItem_Out struct {
}

func (this *DropItem_Out) Process(session *net.Session) {
	g_OutHandler.DropItem(session, this)
}

func (this *DropItem_Out) TypeName() string {
	return "item.drop_item.out"
}

func (this *DropItem_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 7, 1
}

func (this *DropItem_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type BuyItem_In struct {
	ItemId int16 `json:"item_id"`
}

func (this *BuyItem_In) Process(session *net.Session) {
	g_InHandler.BuyItem(session, this)
}

func (this *BuyItem_In) TypeName() string {
	return "item.buy_item.in"
}

func (this *BuyItem_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 7, 2
}

type BuyItem_Out struct {
	Id int64 `json:"id"`
}

func (this *BuyItem_Out) Process(session *net.Session) {
	g_OutHandler.BuyItem(session, this)
}

func (this *BuyItem_Out) TypeName() string {
	return "item.buy_item.out"
}

func (this *BuyItem_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 7, 2
}

func (this *BuyItem_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type SellItem_In struct {
	Id int64 `json:"id"`
}

func (this *SellItem_In) Process(session *net.Session) {
	g_InHandler.SellItem(session, this)
}

func (this *SellItem_In) TypeName() string {
	return "item.sell_item.in"
}

func (this *SellItem_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 7, 3
}

type SellItem_Out struct {
}

func (this *SellItem_Out) Process(session *net.Session) {
	g_OutHandler.SellItem(session, this)
}

func (this *SellItem_Out) TypeName() string {
	return "item.sell_item.out"
}

func (this *SellItem_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 7, 3
}

func (this *SellItem_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type Dress_In struct {
	RoleId int8  `json:"role_id"`
	Id     int64 `json:"id"`
}

func (this *Dress_In) Process(session *net.Session) {
	g_InHandler.Dress(session, this)
}

func (this *Dress_In) TypeName() string {
	return "item.dress.in"
}

func (this *Dress_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 7, 4
}

type Dress_Out struct {
}

func (this *Dress_Out) Process(session *net.Session) {
	g_OutHandler.Dress(session, this)
}

func (this *Dress_Out) TypeName() string {
	return "item.dress.out"
}

func (this *Dress_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 7, 4
}

func (this *Dress_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type Undress_In struct {
	RoleId int8         `json:"role_id"`
	Pos    EquipmentPos `json:"pos"`
}

func (this *Undress_In) Process(session *net.Session) {
	g_InHandler.Undress(session, this)
}

func (this *Undress_In) TypeName() string {
	return "item.undress.in"
}

func (this *Undress_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 7, 5
}

type Undress_Out struct {
}

func (this *Undress_Out) Process(session *net.Session) {
	g_OutHandler.Undress(session, this)
}

func (this *Undress_Out) TypeName() string {
	return "item.undress.out"
}

func (this *Undress_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 7, 5
}

func (this *Undress_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type BuyItemBack_In struct {
	Id int64 `json:"id"`
}

func (this *BuyItemBack_In) Process(session *net.Session) {
	g_InHandler.BuyItemBack(session, this)
}

func (this *BuyItemBack_In) TypeName() string {
	return "item.buy_item_back.in"
}

func (this *BuyItemBack_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 7, 6
}

type BuyItemBack_Out struct {
	Items []BuyItemBack_Out_Items `json:"items"`
}

type BuyItemBack_Out_Items struct {
	Id            int64     `json:"id"`
	ItemId        int16     `json:"item_id"`
	Num           int16     `json:"num"`
	Attack        int32     `json:"attack"`
	Defence       int32     `json:"defence"`
	Health        int32     `json:"health"`
	Speed         int32     `json:"speed"`
	Cultivation   int32     `json:"cultivation"`
	HitLevel      int32     `json:"hit_level"`
	CriticalLevel int32     `json:"critical_level"`
	BlockLevel    int32     `json:"block_level"`
	DestroyLevel  int32     `json:"destroy_level"`
	TenacityLevel int32     `json:"tenacity_level"`
	DodgeLevel    int32     `json:"dodge_level"`
	RefineLevel   int16     `json:"refine_level"`
	RecastAttr    Attribute `json:"recast_attr"`
}

func (this *BuyItemBack_Out) Process(session *net.Session) {
	g_OutHandler.BuyItemBack(session, this)
}

func (this *BuyItemBack_Out) TypeName() string {
	return "item.buy_item_back.out"
}

func (this *BuyItemBack_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 7, 6
}

func (this *BuyItemBack_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type IsBagFull_In struct {
}

func (this *IsBagFull_In) Process(session *net.Session) {
	g_InHandler.IsBagFull(session, this)
}

func (this *IsBagFull_In) TypeName() string {
	return "item.is_bag_full.in"
}

func (this *IsBagFull_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 7, 7
}

type IsBagFull_Out struct {
	IsFull bool `json:"is_full"`
}

func (this *IsBagFull_Out) Process(session *net.Session) {
	g_OutHandler.IsBagFull(session, this)
}

func (this *IsBagFull_Out) TypeName() string {
	return "item.is_bag_full.out"
}

func (this *IsBagFull_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 7, 7
}

func (this *IsBagFull_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type Decompose_In struct {
	Id int64 `json:"id"`
}

func (this *Decompose_In) Process(session *net.Session) {
	g_InHandler.Decompose(session, this)
}

func (this *Decompose_In) TypeName() string {
	return "item.decompose.in"
}

func (this *Decompose_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 7, 8
}

type Decompose_Out struct {
}

func (this *Decompose_Out) Process(session *net.Session) {
	g_OutHandler.Decompose(session, this)
}

func (this *Decompose_Out) TypeName() string {
	return "item.decompose.out"
}

func (this *Decompose_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 7, 8
}

func (this *Decompose_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type Refine_In struct {
	Id      int64 `json:"id"`
	IsBatch bool  `json:"is_batch"`
}

func (this *Refine_In) Process(session *net.Session) {
	g_InHandler.Refine(session, this)
}

func (this *Refine_In) TypeName() string {
	return "item.refine.in"
}

func (this *Refine_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 7, 9
}

type Refine_Out struct {
	Code  int8  `json:"code"`
	Id    int64 `json:"id"`
	Level int16 `json:"level"`
}

func (this *Refine_Out) Process(session *net.Session) {
	g_OutHandler.Refine(session, this)
}

func (this *Refine_Out) TypeName() string {
	return "item.refine.out"
}

func (this *Refine_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 7, 9
}

func (this *Refine_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetRecastInfo_In struct {
	Id   int64     `json:"id"`
	Attr Attribute `json:"attr"`
}

func (this *GetRecastInfo_In) Process(session *net.Session) {
	g_InHandler.GetRecastInfo(session, this)
}

func (this *GetRecastInfo_In) TypeName() string {
	return "item.get_recast_info.in"
}

func (this *GetRecastInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 7, 10
}

type GetRecastInfo_Out struct {
	Attrs []GetRecastInfo_Out_Attrs `json:"attrs"`
}

type GetRecastInfo_Out_Attrs struct {
	Attr  Attribute `json:"attr"`
	Value int32     `json:"value"`
}

func (this *GetRecastInfo_Out) Process(session *net.Session) {
	g_OutHandler.GetRecastInfo(session, this)
}

func (this *GetRecastInfo_Out) TypeName() string {
	return "item.get_recast_info.out"
}

func (this *GetRecastInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 7, 10
}

func (this *GetRecastInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type Recast_In struct {
	Attr Attribute `json:"attr"`
}

func (this *Recast_In) Process(session *net.Session) {
	g_InHandler.Recast(session, this)
}

func (this *Recast_In) TypeName() string {
	return "item.recast.in"
}

func (this *Recast_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 7, 11
}

type Recast_Out struct {
}

func (this *Recast_Out) Process(session *net.Session) {
	g_OutHandler.Recast(session, this)
}

func (this *Recast_Out) TypeName() string {
	return "item.recast.out"
}

func (this *Recast_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 7, 11
}

func (this *Recast_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type UseItem_In struct {
	Id int64 `json:"id"`
}

func (this *UseItem_In) Process(session *net.Session) {
	g_InHandler.UseItem(session, this)
}

func (this *UseItem_In) TypeName() string {
	return "item.use_item.in"
}

func (this *UseItem_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 7, 12
}

type UseItem_Out struct {
	Origin  int64 `json:"origin"`
	Changed bool  `json:"changed"`
}

func (this *UseItem_Out) Process(session *net.Session) {
	g_OutHandler.UseItem(session, this)
}

func (this *UseItem_Out) TypeName() string {
	return "item.use_item.out"
}

func (this *UseItem_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 7, 12
}

func (this *UseItem_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type RoleUseCostItem_In struct {
	RoleId int8  `json:"role_id"`
	ItemId int64 `json:"item_id"`
}

func (this *RoleUseCostItem_In) Process(session *net.Session) {
	g_InHandler.RoleUseCostItem(session, this)
}

func (this *RoleUseCostItem_In) TypeName() string {
	return "item.role_use_cost_item.in"
}

func (this *RoleUseCostItem_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 7, 13
}

type RoleUseCostItem_Out struct {
}

func (this *RoleUseCostItem_Out) Process(session *net.Session) {
	g_OutHandler.RoleUseCostItem(session, this)
}

func (this *RoleUseCostItem_Out) TypeName() string {
	return "item.role_use_cost_item.out"
}

func (this *RoleUseCostItem_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 7, 13
}

func (this *RoleUseCostItem_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type BatchUseItem_In struct {
	RoleId int8  `json:"role_id"`
	Id     int64 `json:"id"`
	Num    int32 `json:"num"`
}

func (this *BatchUseItem_In) Process(session *net.Session) {
	g_InHandler.BatchUseItem(session, this)
}

func (this *BatchUseItem_In) TypeName() string {
	return "item.batch_use_item.in"
}

func (this *BatchUseItem_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 7, 14
}

type BatchUseItem_Out struct {
	Id int64 `json:"id"`
}

func (this *BatchUseItem_Out) Process(session *net.Session) {
	g_OutHandler.BatchUseItem(session, this)
}

func (this *BatchUseItem_Out) TypeName() string {
	return "item.batch_use_item.out"
}

func (this *BatchUseItem_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 7, 14
}

func (this *BatchUseItem_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type DragonBallExchangeNotify_Out struct {
	ItemId  int16 `json:"item_id"`
	ItemNum int16 `json:"item_num"`
}

func (this *DragonBallExchangeNotify_Out) Process(session *net.Session) {
	g_OutHandler.DragonBallExchangeNotify(session, this)
}

func (this *DragonBallExchangeNotify_Out) TypeName() string {
	return "item.dragon_ball_exchange_notify.out"
}

func (this *DragonBallExchangeNotify_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 7, 15
}

func (this *DragonBallExchangeNotify_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type OpenCornucopia_In struct {
	Id int64 `json:"id"`
}

func (this *OpenCornucopia_In) Process(session *net.Session) {
	g_InHandler.OpenCornucopia(session, this)
}

func (this *OpenCornucopia_In) TypeName() string {
	return "item.open_cornucopia.in"
}

func (this *OpenCornucopia_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 7, 16
}

type OpenCornucopia_Out struct {
	Coins int64 `json:"coins"`
}

func (this *OpenCornucopia_Out) Process(session *net.Session) {
	g_OutHandler.OpenCornucopia(session, this)
}

func (this *OpenCornucopia_Out) TypeName() string {
	return "item.open_cornucopia.out"
}

func (this *OpenCornucopia_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 7, 16
}

func (this *OpenCornucopia_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetSealedbooks_In struct {
	ItemType int8 `json:"item_type"`
}

func (this *GetSealedbooks_In) Process(session *net.Session) {
	g_InHandler.GetSealedbooks(session, this)
}

func (this *GetSealedbooks_In) TypeName() string {
	return "item.get_sealedbooks.in"
}

func (this *GetSealedbooks_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 7, 17
}

type GetSealedbooks_Out struct {
	Items []GetSealedbooks_Out_Items `json:"items"`
}

type GetSealedbooks_Out_Items struct {
	ItemType int8  `json:"item_type"`
	ItemId   int64 `json:"item_id"`
	Status   int8  `json:"status"`
}

func (this *GetSealedbooks_Out) Process(session *net.Session) {
	g_OutHandler.GetSealedbooks(session, this)
}

func (this *GetSealedbooks_Out) TypeName() string {
	return "item.get_sealedbooks.out"
}

func (this *GetSealedbooks_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 7, 17
}

func (this *GetSealedbooks_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ActivationSealedbook_In struct {
	ItemType int8  `json:"item_type"`
	ItemId   int64 `json:"item_id"`
}

func (this *ActivationSealedbook_In) Process(session *net.Session) {
	g_InHandler.ActivationSealedbook(session, this)
}

func (this *ActivationSealedbook_In) TypeName() string {
	return "item.activation_sealedbook.in"
}

func (this *ActivationSealedbook_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 7, 18
}

type ActivationSealedbook_Out struct {
	Result bool `json:"result"`
}

func (this *ActivationSealedbook_Out) Process(session *net.Session) {
	g_OutHandler.ActivationSealedbook(session, this)
}

func (this *ActivationSealedbook_Out) TypeName() string {
	return "item.activation_sealedbook.out"
}

func (this *ActivationSealedbook_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 7, 18
}

func (this *ActivationSealedbook_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ExchangeGhostCrystal_In struct {
	ItemId       int16 `json:"item_id"`
	ExchangeType int8  `json:"exchange_type"`
}

func (this *ExchangeGhostCrystal_In) Process(session *net.Session) {
	g_InHandler.ExchangeGhostCrystal(session, this)
}

func (this *ExchangeGhostCrystal_In) TypeName() string {
	return "item.exchange_ghost_crystal.in"
}

func (this *ExchangeGhostCrystal_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 7, 19
}

type ExchangeGhostCrystal_Out struct {
}

func (this *ExchangeGhostCrystal_Out) Process(session *net.Session) {
	g_OutHandler.ExchangeGhostCrystal(session, this)
}

func (this *ExchangeGhostCrystal_Out) TypeName() string {
	return "item.exchange_ghost_crystal.out"
}

func (this *ExchangeGhostCrystal_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 7, 19
}

func (this *ExchangeGhostCrystal_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ExchangeShopItem_In struct {
	Kind int8 `json:"kind"`
}

func (this *ExchangeShopItem_In) Process(session *net.Session) {
	g_InHandler.ExchangeShopItem(session, this)
}

func (this *ExchangeShopItem_In) TypeName() string {
	return "item.exchange_shop_item.in"
}

func (this *ExchangeShopItem_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 7, 20
}

type ExchangeShopItem_Out struct {
	EventType int8 `json:"event_type"`
	Code      int8 `json:"code"`
}

func (this *ExchangeShopItem_Out) Process(session *net.Session) {
	g_OutHandler.ExchangeShopItem(session, this)
}

func (this *ExchangeShopItem_Out) TypeName() string {
	return "item.exchange_shop_item.out"
}

func (this *ExchangeShopItem_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 7, 20
}

func (this *ExchangeShopItem_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetHoildayItemList_In struct {
}

func (this *GetHoildayItemList_In) Process(session *net.Session) {
	g_InHandler.GetHoildayItemList(session, this)
}

func (this *GetHoildayItemList_In) TypeName() string {
	return "item.get_hoilday_item_list.in"
}

func (this *GetHoildayItemList_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 7, 22
}

type GetHoildayItemList_Out struct {
	Items []GetHoildayItemList_Out_Items `json:"items"`
}

type GetHoildayItemList_Out_Items struct {
	Id int32 `json:"id"`
}

func (this *GetHoildayItemList_Out) Process(session *net.Session) {
	g_OutHandler.GetHoildayItemList(session, this)
}

func (this *GetHoildayItemList_Out) TypeName() string {
	return "item.get_hoilday_item_list.out"
}

func (this *GetHoildayItemList_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 7, 22
}

func (this *GetHoildayItemList_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ExchangeHoildayItem_In struct {
	Id int32 `json:"id"`
}

func (this *ExchangeHoildayItem_In) Process(session *net.Session) {
	g_InHandler.ExchangeHoildayItem(session, this)
}

func (this *ExchangeHoildayItem_In) TypeName() string {
	return "item.exchange_hoilday_item.in"
}

func (this *ExchangeHoildayItem_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 7, 23
}

type ExchangeHoildayItem_Out struct {
	Code int8 `json:"code"`
}

func (this *ExchangeHoildayItem_Out) Process(session *net.Session) {
	g_OutHandler.ExchangeHoildayItem(session, this)
}

func (this *ExchangeHoildayItem_Out) TypeName() string {
	return "item.exchange_hoilday_item.out"
}

func (this *ExchangeHoildayItem_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 7, 23
}

func (this *ExchangeHoildayItem_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type BatchExchangeDragonBall_In struct {
	BallId int16 `json:"ball_id"`
}

func (this *BatchExchangeDragonBall_In) Process(session *net.Session) {
	g_InHandler.BatchExchangeDragonBall(session, this)
}

func (this *BatchExchangeDragonBall_In) TypeName() string {
	return "item.batch_exchange_dragon_ball.in"
}

func (this *BatchExchangeDragonBall_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 7, 24
}

type BatchExchangeDragonBall_Out struct {
	Code int8 `json:"code"`
}

func (this *BatchExchangeDragonBall_Out) Process(session *net.Session) {
	g_OutHandler.BatchExchangeDragonBall(session, this)
}

func (this *BatchExchangeDragonBall_Out) TypeName() string {
	return "item.batch_exchange_dragon_ball.out"
}

func (this *BatchExchangeDragonBall_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 7, 24
}

func (this *BatchExchangeDragonBall_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

func (this *GetAllItems_In) Decode(buffer *net.Buffer) {
}

func (this *GetAllItems_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(7)
	buffer.WriteUint8(0)
}

func (this *GetAllItems_In) ByteSize() int {
	size := 2
	return size
}

func (this *GetAllItems_Out) Decode(buffer *net.Buffer) {
	this.Items = make([]GetAllItems_Out_Items, buffer.ReadUint8())
	for i := 0; i < len(this.Items); i++ {
		this.Items[i].Decode(buffer)
	}
	this.Equipments = make([]GetAllItems_Out_Equipments, buffer.ReadUint8())
	for i := 0; i < len(this.Equipments); i++ {
		this.Equipments[i].Decode(buffer)
	}
	this.Buybacks = make([]GetAllItems_Out_Buybacks, buffer.ReadUint8())
	for i := 0; i < len(this.Buybacks); i++ {
		this.Buybacks[i].Decode(buffer)
	}
	this.BuyRecords = make([]GetAllItems_Out_BuyRecords, buffer.ReadUint8())
	for i := 0; i < len(this.BuyRecords); i++ {
		this.BuyRecords[i].Decode(buffer)
	}
}

func (this *GetAllItems_Out_Items) Decode(buffer *net.Buffer) {
	this.Id = int64(buffer.ReadUint64LE())
	this.ItemId = int16(buffer.ReadUint16LE())
	this.Num = int16(buffer.ReadUint16LE())
	this.Attack = int32(buffer.ReadUint32LE())
	this.Defence = int32(buffer.ReadUint32LE())
	this.Health = int32(buffer.ReadUint32LE())
	this.Speed = int32(buffer.ReadUint32LE())
	this.Cultivation = int32(buffer.ReadUint32LE())
	this.HitLevel = int32(buffer.ReadUint32LE())
	this.CriticalLevel = int32(buffer.ReadUint32LE())
	this.BlockLevel = int32(buffer.ReadUint32LE())
	this.DestroyLevel = int32(buffer.ReadUint32LE())
	this.TenacityLevel = int32(buffer.ReadUint32LE())
	this.DodgeLevel = int32(buffer.ReadUint32LE())
	this.RefineLevel = int16(buffer.ReadUint16LE())
	this.RefineFailTimes = int16(buffer.ReadUint16LE())
	this.RecastAttr = Attribute(buffer.ReadUint8())
}

func (this *GetAllItems_Out_Equipments) Decode(buffer *net.Buffer) {
	this.RoleId = int8(buffer.ReadUint8())
	this.Equips = make([]GetAllItems_Out_Equipments_Equips, buffer.ReadUint8())
	for i := 0; i < len(this.Equips); i++ {
		this.Equips[i].Decode(buffer)
	}
}

func (this *GetAllItems_Out_Equipments_Equips) Decode(buffer *net.Buffer) {
	this.Id = int64(buffer.ReadUint64LE())
	this.ItemId = int16(buffer.ReadUint16LE())
	this.Attack = int32(buffer.ReadUint32LE())
	this.Defence = int32(buffer.ReadUint32LE())
	this.Health = int32(buffer.ReadUint32LE())
	this.Speed = int32(buffer.ReadUint32LE())
	this.Cultivation = int32(buffer.ReadUint32LE())
	this.HitLevel = int32(buffer.ReadUint32LE())
	this.CriticalLevel = int32(buffer.ReadUint32LE())
	this.BlockLevel = int32(buffer.ReadUint32LE())
	this.DestroyLevel = int32(buffer.ReadUint32LE())
	this.TenacityLevel = int32(buffer.ReadUint32LE())
	this.DodgeLevel = int32(buffer.ReadUint32LE())
	this.RefineLevel = int16(buffer.ReadUint16LE())
	this.RefineFailTimes = int16(buffer.ReadUint16LE())
	this.RecastAttr = Attribute(buffer.ReadUint8())
}

func (this *GetAllItems_Out_Buybacks) Decode(buffer *net.Buffer) {
	this.Id = int64(buffer.ReadUint64LE())
	this.ItemId = int16(buffer.ReadUint16LE())
	this.Num = int16(buffer.ReadUint16LE())
	this.RefineLevel = int16(buffer.ReadUint16LE())
	this.RecastAttr = Attribute(buffer.ReadUint8())
}

func (this *GetAllItems_Out_BuyRecords) Decode(buffer *net.Buffer) {
	this.ItemId = int16(buffer.ReadUint16LE())
	this.Num = int16(buffer.ReadUint16LE())
}

func (this *GetAllItems_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(7)
	buffer.WriteUint8(0)
	buffer.WriteUint8(uint8(len(this.Items)))
	for i := 0; i < len(this.Items); i++ {
		this.Items[i].Encode(buffer)
	}
	buffer.WriteUint8(uint8(len(this.Equipments)))
	for i := 0; i < len(this.Equipments); i++ {
		this.Equipments[i].Encode(buffer)
	}
	buffer.WriteUint8(uint8(len(this.Buybacks)))
	for i := 0; i < len(this.Buybacks); i++ {
		this.Buybacks[i].Encode(buffer)
	}
	buffer.WriteUint8(uint8(len(this.BuyRecords)))
	for i := 0; i < len(this.BuyRecords); i++ {
		this.BuyRecords[i].Encode(buffer)
	}
}

func (this *GetAllItems_Out_Items) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.Id))
	buffer.WriteUint16LE(uint16(this.ItemId))
	buffer.WriteUint16LE(uint16(this.Num))
	buffer.WriteUint32LE(uint32(this.Attack))
	buffer.WriteUint32LE(uint32(this.Defence))
	buffer.WriteUint32LE(uint32(this.Health))
	buffer.WriteUint32LE(uint32(this.Speed))
	buffer.WriteUint32LE(uint32(this.Cultivation))
	buffer.WriteUint32LE(uint32(this.HitLevel))
	buffer.WriteUint32LE(uint32(this.CriticalLevel))
	buffer.WriteUint32LE(uint32(this.BlockLevel))
	buffer.WriteUint32LE(uint32(this.DestroyLevel))
	buffer.WriteUint32LE(uint32(this.TenacityLevel))
	buffer.WriteUint32LE(uint32(this.DodgeLevel))
	buffer.WriteUint16LE(uint16(this.RefineLevel))
	buffer.WriteUint16LE(uint16(this.RefineFailTimes))
	buffer.WriteUint8(uint8(this.RecastAttr))
}

func (this *GetAllItems_Out_Equipments) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint8(uint8(len(this.Equips)))
	for i := 0; i < len(this.Equips); i++ {
		this.Equips[i].Encode(buffer)
	}
}

func (this *GetAllItems_Out_Equipments_Equips) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.Id))
	buffer.WriteUint16LE(uint16(this.ItemId))
	buffer.WriteUint32LE(uint32(this.Attack))
	buffer.WriteUint32LE(uint32(this.Defence))
	buffer.WriteUint32LE(uint32(this.Health))
	buffer.WriteUint32LE(uint32(this.Speed))
	buffer.WriteUint32LE(uint32(this.Cultivation))
	buffer.WriteUint32LE(uint32(this.HitLevel))
	buffer.WriteUint32LE(uint32(this.CriticalLevel))
	buffer.WriteUint32LE(uint32(this.BlockLevel))
	buffer.WriteUint32LE(uint32(this.DestroyLevel))
	buffer.WriteUint32LE(uint32(this.TenacityLevel))
	buffer.WriteUint32LE(uint32(this.DodgeLevel))
	buffer.WriteUint16LE(uint16(this.RefineLevel))
	buffer.WriteUint16LE(uint16(this.RefineFailTimes))
	buffer.WriteUint8(uint8(this.RecastAttr))
}

func (this *GetAllItems_Out_Buybacks) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.Id))
	buffer.WriteUint16LE(uint16(this.ItemId))
	buffer.WriteUint16LE(uint16(this.Num))
	buffer.WriteUint16LE(uint16(this.RefineLevel))
	buffer.WriteUint8(uint8(this.RecastAttr))
}

func (this *GetAllItems_Out_BuyRecords) Encode(buffer *net.Buffer) {
	buffer.WriteUint16LE(uint16(this.ItemId))
	buffer.WriteUint16LE(uint16(this.Num))
}

func (this *GetAllItems_Out) ByteSize() int {
	size := 6
	size += len(this.Items) * 61
	for i := 0; i < len(this.Equipments); i++ {
		size += this.Equipments[i].ByteSize()
	}
	size += len(this.Buybacks) * 15
	size += len(this.BuyRecords) * 4
	return size
}

func (this *GetAllItems_Out_Equipments) ByteSize() int {
	size := 2
	size += len(this.Equips) * 59
	return size
}

func (this *DropItem_In) Decode(buffer *net.Buffer) {
	this.Id = int64(buffer.ReadUint64LE())
}

func (this *DropItem_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(7)
	buffer.WriteUint8(1)
	buffer.WriteUint64LE(uint64(this.Id))
}

func (this *DropItem_In) ByteSize() int {
	size := 10
	return size
}

func (this *DropItem_Out) Decode(buffer *net.Buffer) {
}

func (this *DropItem_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(7)
	buffer.WriteUint8(1)
}

func (this *DropItem_Out) ByteSize() int {
	size := 2
	return size
}

func (this *BuyItem_In) Decode(buffer *net.Buffer) {
	this.ItemId = int16(buffer.ReadUint16LE())
}

func (this *BuyItem_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(7)
	buffer.WriteUint8(2)
	buffer.WriteUint16LE(uint16(this.ItemId))
}

func (this *BuyItem_In) ByteSize() int {
	size := 4
	return size
}

func (this *BuyItem_Out) Decode(buffer *net.Buffer) {
	this.Id = int64(buffer.ReadUint64LE())
}

func (this *BuyItem_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(7)
	buffer.WriteUint8(2)
	buffer.WriteUint64LE(uint64(this.Id))
}

func (this *BuyItem_Out) ByteSize() int {
	size := 10
	return size
}

func (this *SellItem_In) Decode(buffer *net.Buffer) {
	this.Id = int64(buffer.ReadUint64LE())
}

func (this *SellItem_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(7)
	buffer.WriteUint8(3)
	buffer.WriteUint64LE(uint64(this.Id))
}

func (this *SellItem_In) ByteSize() int {
	size := 10
	return size
}

func (this *SellItem_Out) Decode(buffer *net.Buffer) {
}

func (this *SellItem_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(7)
	buffer.WriteUint8(3)
}

func (this *SellItem_Out) ByteSize() int {
	size := 2
	return size
}

func (this *Dress_In) Decode(buffer *net.Buffer) {
	this.RoleId = int8(buffer.ReadUint8())
	this.Id = int64(buffer.ReadUint64LE())
}

func (this *Dress_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(7)
	buffer.WriteUint8(4)
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint64LE(uint64(this.Id))
}

func (this *Dress_In) ByteSize() int {
	size := 11
	return size
}

func (this *Dress_Out) Decode(buffer *net.Buffer) {
}

func (this *Dress_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(7)
	buffer.WriteUint8(4)
}

func (this *Dress_Out) ByteSize() int {
	size := 2
	return size
}

func (this *Undress_In) Decode(buffer *net.Buffer) {
	this.RoleId = int8(buffer.ReadUint8())
	this.Pos = EquipmentPos(buffer.ReadUint8())
}

func (this *Undress_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(7)
	buffer.WriteUint8(5)
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint8(uint8(this.Pos))
}

func (this *Undress_In) ByteSize() int {
	size := 4
	return size
}

func (this *Undress_Out) Decode(buffer *net.Buffer) {
}

func (this *Undress_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(7)
	buffer.WriteUint8(5)
}

func (this *Undress_Out) ByteSize() int {
	size := 2
	return size
}

func (this *BuyItemBack_In) Decode(buffer *net.Buffer) {
	this.Id = int64(buffer.ReadUint64LE())
}

func (this *BuyItemBack_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(7)
	buffer.WriteUint8(6)
	buffer.WriteUint64LE(uint64(this.Id))
}

func (this *BuyItemBack_In) ByteSize() int {
	size := 10
	return size
}

func (this *BuyItemBack_Out) Decode(buffer *net.Buffer) {
	this.Items = make([]BuyItemBack_Out_Items, buffer.ReadUint8())
	for i := 0; i < len(this.Items); i++ {
		this.Items[i].Decode(buffer)
	}
}

func (this *BuyItemBack_Out_Items) Decode(buffer *net.Buffer) {
	this.Id = int64(buffer.ReadUint64LE())
	this.ItemId = int16(buffer.ReadUint16LE())
	this.Num = int16(buffer.ReadUint16LE())
	this.Attack = int32(buffer.ReadUint32LE())
	this.Defence = int32(buffer.ReadUint32LE())
	this.Health = int32(buffer.ReadUint32LE())
	this.Speed = int32(buffer.ReadUint32LE())
	this.Cultivation = int32(buffer.ReadUint32LE())
	this.HitLevel = int32(buffer.ReadUint32LE())
	this.CriticalLevel = int32(buffer.ReadUint32LE())
	this.BlockLevel = int32(buffer.ReadUint32LE())
	this.DestroyLevel = int32(buffer.ReadUint32LE())
	this.TenacityLevel = int32(buffer.ReadUint32LE())
	this.DodgeLevel = int32(buffer.ReadUint32LE())
	this.RefineLevel = int16(buffer.ReadUint16LE())
	this.RecastAttr = Attribute(buffer.ReadUint8())
}

func (this *BuyItemBack_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(7)
	buffer.WriteUint8(6)
	buffer.WriteUint8(uint8(len(this.Items)))
	for i := 0; i < len(this.Items); i++ {
		this.Items[i].Encode(buffer)
	}
}

func (this *BuyItemBack_Out_Items) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.Id))
	buffer.WriteUint16LE(uint16(this.ItemId))
	buffer.WriteUint16LE(uint16(this.Num))
	buffer.WriteUint32LE(uint32(this.Attack))
	buffer.WriteUint32LE(uint32(this.Defence))
	buffer.WriteUint32LE(uint32(this.Health))
	buffer.WriteUint32LE(uint32(this.Speed))
	buffer.WriteUint32LE(uint32(this.Cultivation))
	buffer.WriteUint32LE(uint32(this.HitLevel))
	buffer.WriteUint32LE(uint32(this.CriticalLevel))
	buffer.WriteUint32LE(uint32(this.BlockLevel))
	buffer.WriteUint32LE(uint32(this.DestroyLevel))
	buffer.WriteUint32LE(uint32(this.TenacityLevel))
	buffer.WriteUint32LE(uint32(this.DodgeLevel))
	buffer.WriteUint16LE(uint16(this.RefineLevel))
	buffer.WriteUint8(uint8(this.RecastAttr))
}

func (this *BuyItemBack_Out) ByteSize() int {
	size := 3
	size += len(this.Items) * 59
	return size
}

func (this *IsBagFull_In) Decode(buffer *net.Buffer) {
}

func (this *IsBagFull_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(7)
	buffer.WriteUint8(7)
}

func (this *IsBagFull_In) ByteSize() int {
	size := 2
	return size
}

func (this *IsBagFull_Out) Decode(buffer *net.Buffer) {
	this.IsFull = buffer.ReadUint8() == 1
}

func (this *IsBagFull_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(7)
	buffer.WriteUint8(7)
	if this.IsFull {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *IsBagFull_Out) ByteSize() int {
	size := 3
	return size
}

func (this *Decompose_In) Decode(buffer *net.Buffer) {
	this.Id = int64(buffer.ReadUint64LE())
}

func (this *Decompose_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(7)
	buffer.WriteUint8(8)
	buffer.WriteUint64LE(uint64(this.Id))
}

func (this *Decompose_In) ByteSize() int {
	size := 10
	return size
}

func (this *Decompose_Out) Decode(buffer *net.Buffer) {
}

func (this *Decompose_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(7)
	buffer.WriteUint8(8)
}

func (this *Decompose_Out) ByteSize() int {
	size := 2
	return size
}

func (this *Refine_In) Decode(buffer *net.Buffer) {
	this.Id = int64(buffer.ReadUint64LE())
	this.IsBatch = buffer.ReadUint8() == 1
}

func (this *Refine_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(7)
	buffer.WriteUint8(9)
	buffer.WriteUint64LE(uint64(this.Id))
	if this.IsBatch {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *Refine_In) ByteSize() int {
	size := 11
	return size
}

func (this *Refine_Out) Decode(buffer *net.Buffer) {
	this.Code = int8(buffer.ReadUint8())
	this.Id = int64(buffer.ReadUint64LE())
	this.Level = int16(buffer.ReadUint16LE())
}

func (this *Refine_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(7)
	buffer.WriteUint8(9)
	buffer.WriteUint8(uint8(this.Code))
	buffer.WriteUint64LE(uint64(this.Id))
	buffer.WriteUint16LE(uint16(this.Level))
}

func (this *Refine_Out) ByteSize() int {
	size := 13
	return size
}

func (this *GetRecastInfo_In) Decode(buffer *net.Buffer) {
	this.Id = int64(buffer.ReadUint64LE())
	this.Attr = Attribute(buffer.ReadUint8())
}

func (this *GetRecastInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(7)
	buffer.WriteUint8(10)
	buffer.WriteUint64LE(uint64(this.Id))
	buffer.WriteUint8(uint8(this.Attr))
}

func (this *GetRecastInfo_In) ByteSize() int {
	size := 11
	return size
}

func (this *GetRecastInfo_Out) Decode(buffer *net.Buffer) {
	this.Attrs = make([]GetRecastInfo_Out_Attrs, buffer.ReadUint8())
	for i := 0; i < len(this.Attrs); i++ {
		this.Attrs[i].Decode(buffer)
	}
}

func (this *GetRecastInfo_Out_Attrs) Decode(buffer *net.Buffer) {
	this.Attr = Attribute(buffer.ReadUint8())
	this.Value = int32(buffer.ReadUint32LE())
}

func (this *GetRecastInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(7)
	buffer.WriteUint8(10)
	buffer.WriteUint8(uint8(len(this.Attrs)))
	for i := 0; i < len(this.Attrs); i++ {
		this.Attrs[i].Encode(buffer)
	}
}

func (this *GetRecastInfo_Out_Attrs) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(this.Attr))
	buffer.WriteUint32LE(uint32(this.Value))
}

func (this *GetRecastInfo_Out) ByteSize() int {
	size := 3
	size += len(this.Attrs) * 5
	return size
}

func (this *Recast_In) Decode(buffer *net.Buffer) {
	this.Attr = Attribute(buffer.ReadUint8())
}

func (this *Recast_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(7)
	buffer.WriteUint8(11)
	buffer.WriteUint8(uint8(this.Attr))
}

func (this *Recast_In) ByteSize() int {
	size := 3
	return size
}

func (this *Recast_Out) Decode(buffer *net.Buffer) {
}

func (this *Recast_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(7)
	buffer.WriteUint8(11)
}

func (this *Recast_Out) ByteSize() int {
	size := 2
	return size
}

func (this *UseItem_In) Decode(buffer *net.Buffer) {
	this.Id = int64(buffer.ReadUint64LE())
}

func (this *UseItem_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(7)
	buffer.WriteUint8(12)
	buffer.WriteUint64LE(uint64(this.Id))
}

func (this *UseItem_In) ByteSize() int {
	size := 10
	return size
}

func (this *UseItem_Out) Decode(buffer *net.Buffer) {
	this.Origin = int64(buffer.ReadUint64LE())
	this.Changed = buffer.ReadUint8() == 1
}

func (this *UseItem_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(7)
	buffer.WriteUint8(12)
	buffer.WriteUint64LE(uint64(this.Origin))
	if this.Changed {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *UseItem_Out) ByteSize() int {
	size := 11
	return size
}

func (this *RoleUseCostItem_In) Decode(buffer *net.Buffer) {
	this.RoleId = int8(buffer.ReadUint8())
	this.ItemId = int64(buffer.ReadUint64LE())
}

func (this *RoleUseCostItem_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(7)
	buffer.WriteUint8(13)
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint64LE(uint64(this.ItemId))
}

func (this *RoleUseCostItem_In) ByteSize() int {
	size := 11
	return size
}

func (this *RoleUseCostItem_Out) Decode(buffer *net.Buffer) {
}

func (this *RoleUseCostItem_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(7)
	buffer.WriteUint8(13)
}

func (this *RoleUseCostItem_Out) ByteSize() int {
	size := 2
	return size
}

func (this *BatchUseItem_In) Decode(buffer *net.Buffer) {
	this.RoleId = int8(buffer.ReadUint8())
	this.Id = int64(buffer.ReadUint64LE())
	this.Num = int32(buffer.ReadUint32LE())
}

func (this *BatchUseItem_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(7)
	buffer.WriteUint8(14)
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint64LE(uint64(this.Id))
	buffer.WriteUint32LE(uint32(this.Num))
}

func (this *BatchUseItem_In) ByteSize() int {
	size := 15
	return size
}

func (this *BatchUseItem_Out) Decode(buffer *net.Buffer) {
	this.Id = int64(buffer.ReadUint64LE())
}

func (this *BatchUseItem_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(7)
	buffer.WriteUint8(14)
	buffer.WriteUint64LE(uint64(this.Id))
}

func (this *BatchUseItem_Out) ByteSize() int {
	size := 10
	return size
}

func (this *DragonBallExchangeNotify_Out) Decode(buffer *net.Buffer) {
	this.ItemId = int16(buffer.ReadUint16LE())
	this.ItemNum = int16(buffer.ReadUint16LE())
}

func (this *DragonBallExchangeNotify_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(7)
	buffer.WriteUint8(15)
	buffer.WriteUint16LE(uint16(this.ItemId))
	buffer.WriteUint16LE(uint16(this.ItemNum))
}

func (this *DragonBallExchangeNotify_Out) ByteSize() int {
	size := 6
	return size
}

func (this *OpenCornucopia_In) Decode(buffer *net.Buffer) {
	this.Id = int64(buffer.ReadUint64LE())
}

func (this *OpenCornucopia_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(7)
	buffer.WriteUint8(16)
	buffer.WriteUint64LE(uint64(this.Id))
}

func (this *OpenCornucopia_In) ByteSize() int {
	size := 10
	return size
}

func (this *OpenCornucopia_Out) Decode(buffer *net.Buffer) {
	this.Coins = int64(buffer.ReadUint64LE())
}

func (this *OpenCornucopia_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(7)
	buffer.WriteUint8(16)
	buffer.WriteUint64LE(uint64(this.Coins))
}

func (this *OpenCornucopia_Out) ByteSize() int {
	size := 10
	return size
}

func (this *GetSealedbooks_In) Decode(buffer *net.Buffer) {
	this.ItemType = int8(buffer.ReadUint8())
}

func (this *GetSealedbooks_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(7)
	buffer.WriteUint8(17)
	buffer.WriteUint8(uint8(this.ItemType))
}

func (this *GetSealedbooks_In) ByteSize() int {
	size := 3
	return size
}

func (this *GetSealedbooks_Out) Decode(buffer *net.Buffer) {
	this.Items = make([]GetSealedbooks_Out_Items, buffer.ReadUint8())
	for i := 0; i < len(this.Items); i++ {
		this.Items[i].Decode(buffer)
	}
}

func (this *GetSealedbooks_Out_Items) Decode(buffer *net.Buffer) {
	this.ItemType = int8(buffer.ReadUint8())
	this.ItemId = int64(buffer.ReadUint64LE())
	this.Status = int8(buffer.ReadUint8())
}

func (this *GetSealedbooks_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(7)
	buffer.WriteUint8(17)
	buffer.WriteUint8(uint8(len(this.Items)))
	for i := 0; i < len(this.Items); i++ {
		this.Items[i].Encode(buffer)
	}
}

func (this *GetSealedbooks_Out_Items) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(this.ItemType))
	buffer.WriteUint64LE(uint64(this.ItemId))
	buffer.WriteUint8(uint8(this.Status))
}

func (this *GetSealedbooks_Out) ByteSize() int {
	size := 3
	size += len(this.Items) * 10
	return size
}

func (this *ActivationSealedbook_In) Decode(buffer *net.Buffer) {
	this.ItemType = int8(buffer.ReadUint8())
	this.ItemId = int64(buffer.ReadUint64LE())
}

func (this *ActivationSealedbook_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(7)
	buffer.WriteUint8(18)
	buffer.WriteUint8(uint8(this.ItemType))
	buffer.WriteUint64LE(uint64(this.ItemId))
}

func (this *ActivationSealedbook_In) ByteSize() int {
	size := 11
	return size
}

func (this *ActivationSealedbook_Out) Decode(buffer *net.Buffer) {
	this.Result = buffer.ReadUint8() == 1
}

func (this *ActivationSealedbook_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(7)
	buffer.WriteUint8(18)
	if this.Result {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *ActivationSealedbook_Out) ByteSize() int {
	size := 3
	return size
}

func (this *ExchangeGhostCrystal_In) Decode(buffer *net.Buffer) {
	this.ItemId = int16(buffer.ReadUint16LE())
	this.ExchangeType = int8(buffer.ReadUint8())
}

func (this *ExchangeGhostCrystal_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(7)
	buffer.WriteUint8(19)
	buffer.WriteUint16LE(uint16(this.ItemId))
	buffer.WriteUint8(uint8(this.ExchangeType))
}

func (this *ExchangeGhostCrystal_In) ByteSize() int {
	size := 5
	return size
}

func (this *ExchangeGhostCrystal_Out) Decode(buffer *net.Buffer) {
}

func (this *ExchangeGhostCrystal_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(7)
	buffer.WriteUint8(19)
}

func (this *ExchangeGhostCrystal_Out) ByteSize() int {
	size := 2
	return size
}

func (this *ExchangeShopItem_In) Decode(buffer *net.Buffer) {
	this.Kind = int8(buffer.ReadUint8())
}

func (this *ExchangeShopItem_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(7)
	buffer.WriteUint8(20)
	buffer.WriteUint8(uint8(this.Kind))
}

func (this *ExchangeShopItem_In) ByteSize() int {
	size := 3
	return size
}

func (this *ExchangeShopItem_Out) Decode(buffer *net.Buffer) {
	this.EventType = int8(buffer.ReadUint8())
	this.Code = int8(buffer.ReadUint8())
}

func (this *ExchangeShopItem_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(7)
	buffer.WriteUint8(20)
	buffer.WriteUint8(uint8(this.EventType))
	buffer.WriteUint8(uint8(this.Code))
}

func (this *ExchangeShopItem_Out) ByteSize() int {
	size := 4
	return size
}

func (this *GetHoildayItemList_In) Decode(buffer *net.Buffer) {
}

func (this *GetHoildayItemList_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(7)
	buffer.WriteUint8(22)
}

func (this *GetHoildayItemList_In) ByteSize() int {
	size := 2
	return size
}

func (this *GetHoildayItemList_Out) Decode(buffer *net.Buffer) {
	this.Items = make([]GetHoildayItemList_Out_Items, buffer.ReadUint8())
	for i := 0; i < len(this.Items); i++ {
		this.Items[i].Decode(buffer)
	}
}

func (this *GetHoildayItemList_Out_Items) Decode(buffer *net.Buffer) {
	this.Id = int32(buffer.ReadUint32LE())
}

func (this *GetHoildayItemList_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(7)
	buffer.WriteUint8(22)
	buffer.WriteUint8(uint8(len(this.Items)))
	for i := 0; i < len(this.Items); i++ {
		this.Items[i].Encode(buffer)
	}
}

func (this *GetHoildayItemList_Out_Items) Encode(buffer *net.Buffer) {
	buffer.WriteUint32LE(uint32(this.Id))
}

func (this *GetHoildayItemList_Out) ByteSize() int {
	size := 3
	size += len(this.Items) * 4
	return size
}

func (this *ExchangeHoildayItem_In) Decode(buffer *net.Buffer) {
	this.Id = int32(buffer.ReadUint32LE())
}

func (this *ExchangeHoildayItem_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(7)
	buffer.WriteUint8(23)
	buffer.WriteUint32LE(uint32(this.Id))
}

func (this *ExchangeHoildayItem_In) ByteSize() int {
	size := 6
	return size
}

func (this *ExchangeHoildayItem_Out) Decode(buffer *net.Buffer) {
	this.Code = int8(buffer.ReadUint8())
}

func (this *ExchangeHoildayItem_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(7)
	buffer.WriteUint8(23)
	buffer.WriteUint8(uint8(this.Code))
}

func (this *ExchangeHoildayItem_Out) ByteSize() int {
	size := 3
	return size
}

func (this *BatchExchangeDragonBall_In) Decode(buffer *net.Buffer) {
	this.BallId = int16(buffer.ReadUint16LE())
}

func (this *BatchExchangeDragonBall_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(7)
	buffer.WriteUint8(24)
	buffer.WriteUint16LE(uint16(this.BallId))
}

func (this *BatchExchangeDragonBall_In) ByteSize() int {
	size := 4
	return size
}

func (this *BatchExchangeDragonBall_Out) Decode(buffer *net.Buffer) {
	this.Code = int8(buffer.ReadUint8())
}

func (this *BatchExchangeDragonBall_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(7)
	buffer.WriteUint8(24)
	buffer.WriteUint8(uint8(this.Code))
}

func (this *BatchExchangeDragonBall_Out) ByteSize() int {
	size := 3
	return size
}
