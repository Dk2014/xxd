package taoyuan_api

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
	TaoyuanInfo(*net.Session, *TaoyuanInfo_In)
	Bless(*net.Session, *Bless_In)
	ShopBuy(*net.Session, *ShopBuy_In)
	ShopSell(*net.Session, *ShopSell_In)
	GetAllItems(*net.Session, *GetAllItems_In)
	FiledsInfo(*net.Session, *FiledsInfo_In)
	GrowPlant(*net.Session, *GrowPlant_In)
	TakePlant(*net.Session, *TakePlant_In)
	UpgradeFiled(*net.Session, *UpgradeFiled_In)
	OpenFiled(*net.Session, *OpenFiled_In)
	StudySkill(*net.Session, *StudySkill_In)
	MakeProduct(*net.Session, *MakeProduct_In)
	TakeProduct(*net.Session, *TakeProduct_In)
	ProductMakeQueue(*net.Session, *ProductMakeQueue_In)
	QuestInfo(*net.Session, *QuestInfo_In)
	QuestFinish(*net.Session, *QuestFinish_In)
	QuestRefresh(*net.Session, *QuestRefresh_In)
	FriendTaoyuanInfo(*net.Session, *FriendTaoyuanInfo_In)
	SkillInfo(*net.Session, *SkillInfo_In)
	OpenQueue(*net.Session, *OpenQueue_In)
	PlantQuicklyMaturity(*net.Session, *PlantQuicklyMaturity_In)
	TaoyuanMessageInfo(*net.Session, *TaoyuanMessageInfo_In)
	TaoyuanMessageRead(*net.Session, *TaoyuanMessageRead_In)
	OpenProductBuilding(*net.Session, *OpenProductBuilding_In)
}

type OutHandler interface {
	TaoyuanInfo(*net.Session, *TaoyuanInfo_Out)
	Bless(*net.Session, *Bless_Out)
	ShopBuy(*net.Session, *ShopBuy_Out)
	ShopSell(*net.Session, *ShopSell_Out)
	GetAllItems(*net.Session, *GetAllItems_Out)
	FiledsInfo(*net.Session, *FiledsInfo_Out)
	GrowPlant(*net.Session, *GrowPlant_Out)
	TakePlant(*net.Session, *TakePlant_Out)
	UpgradeFiled(*net.Session, *UpgradeFiled_Out)
	OpenFiled(*net.Session, *OpenFiled_Out)
	StudySkill(*net.Session, *StudySkill_Out)
	MakeProduct(*net.Session, *MakeProduct_Out)
	TakeProduct(*net.Session, *TakeProduct_Out)
	ProductMakeQueue(*net.Session, *ProductMakeQueue_Out)
	QuestInfo(*net.Session, *QuestInfo_Out)
	QuestFinish(*net.Session, *QuestFinish_Out)
	QuestRefresh(*net.Session, *QuestRefresh_Out)
	FriendTaoyuanInfo(*net.Session, *FriendTaoyuanInfo_Out)
	SkillInfo(*net.Session, *SkillInfo_Out)
	OpenQueue(*net.Session, *OpenQueue_Out)
	PlantQuicklyMaturity(*net.Session, *PlantQuicklyMaturity_Out)
	TaoyuanMessageInfo(*net.Session, *TaoyuanMessageInfo_Out)
	TaoyuanMessageRead(*net.Session, *TaoyuanMessageRead_Out)
	OpenProductBuilding(*net.Session, *OpenProductBuilding_Out)
}

func DecodeIn(message []byte) Request {
	var actionId = message[0]
	var buffer = net.NewBuffer(message[1:])
	switch actionId {
	case 0:
		request := new(TaoyuanInfo_In)
		request.Decode(buffer)
		return request
	case 1:
		request := new(Bless_In)
		request.Decode(buffer)
		return request
	case 2:
		request := new(ShopBuy_In)
		request.Decode(buffer)
		return request
	case 3:
		request := new(ShopSell_In)
		request.Decode(buffer)
		return request
	case 4:
		request := new(GetAllItems_In)
		request.Decode(buffer)
		return request
	case 5:
		request := new(FiledsInfo_In)
		request.Decode(buffer)
		return request
	case 6:
		request := new(GrowPlant_In)
		request.Decode(buffer)
		return request
	case 7:
		request := new(TakePlant_In)
		request.Decode(buffer)
		return request
	case 8:
		request := new(UpgradeFiled_In)
		request.Decode(buffer)
		return request
	case 9:
		request := new(OpenFiled_In)
		request.Decode(buffer)
		return request
	case 10:
		request := new(StudySkill_In)
		request.Decode(buffer)
		return request
	case 11:
		request := new(MakeProduct_In)
		request.Decode(buffer)
		return request
	case 12:
		request := new(TakeProduct_In)
		request.Decode(buffer)
		return request
	case 13:
		request := new(ProductMakeQueue_In)
		request.Decode(buffer)
		return request
	case 14:
		request := new(QuestInfo_In)
		request.Decode(buffer)
		return request
	case 15:
		request := new(QuestFinish_In)
		request.Decode(buffer)
		return request
	case 16:
		request := new(QuestRefresh_In)
		request.Decode(buffer)
		return request
	case 17:
		request := new(FriendTaoyuanInfo_In)
		request.Decode(buffer)
		return request
	case 18:
		request := new(SkillInfo_In)
		request.Decode(buffer)
		return request
	case 19:
		request := new(OpenQueue_In)
		request.Decode(buffer)
		return request
	case 20:
		request := new(PlantQuicklyMaturity_In)
		request.Decode(buffer)
		return request
	case 21:
		request := new(TaoyuanMessageInfo_In)
		request.Decode(buffer)
		return request
	case 22:
		request := new(TaoyuanMessageRead_In)
		request.Decode(buffer)
		return request
	case 23:
		request := new(OpenProductBuilding_In)
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
		request := new(TaoyuanInfo_Out)
		request.Decode(buffer)
		return request
	case 1:
		request := new(Bless_Out)
		request.Decode(buffer)
		return request
	case 2:
		request := new(ShopBuy_Out)
		request.Decode(buffer)
		return request
	case 3:
		request := new(ShopSell_Out)
		request.Decode(buffer)
		return request
	case 4:
		request := new(GetAllItems_Out)
		request.Decode(buffer)
		return request
	case 5:
		request := new(FiledsInfo_Out)
		request.Decode(buffer)
		return request
	case 6:
		request := new(GrowPlant_Out)
		request.Decode(buffer)
		return request
	case 7:
		request := new(TakePlant_Out)
		request.Decode(buffer)
		return request
	case 8:
		request := new(UpgradeFiled_Out)
		request.Decode(buffer)
		return request
	case 9:
		request := new(OpenFiled_Out)
		request.Decode(buffer)
		return request
	case 10:
		request := new(StudySkill_Out)
		request.Decode(buffer)
		return request
	case 11:
		request := new(MakeProduct_Out)
		request.Decode(buffer)
		return request
	case 12:
		request := new(TakeProduct_Out)
		request.Decode(buffer)
		return request
	case 13:
		request := new(ProductMakeQueue_Out)
		request.Decode(buffer)
		return request
	case 14:
		request := new(QuestInfo_Out)
		request.Decode(buffer)
		return request
	case 15:
		request := new(QuestFinish_Out)
		request.Decode(buffer)
		return request
	case 16:
		request := new(QuestRefresh_Out)
		request.Decode(buffer)
		return request
	case 17:
		request := new(FriendTaoyuanInfo_Out)
		request.Decode(buffer)
		return request
	case 18:
		request := new(SkillInfo_Out)
		request.Decode(buffer)
		return request
	case 19:
		request := new(OpenQueue_Out)
		request.Decode(buffer)
		return request
	case 20:
		request := new(PlantQuicklyMaturity_Out)
		request.Decode(buffer)
		return request
	case 21:
		request := new(TaoyuanMessageInfo_Out)
		request.Decode(buffer)
		return request
	case 22:
		request := new(TaoyuanMessageRead_Out)
		request.Decode(buffer)
		return request
	case 23:
		request := new(OpenProductBuilding_Out)
		request.Decode(buffer)
		return request
	}
	_ = buffer
	panic("unsupported response")
}

type BlessResult int8

const (
	BLESS_RESULT_SUCCESS                    BlessResult = 0
	BLESS_RESULT_DAILY_BLESS_TIMES_LIMIT    BlessResult = 1
	BLESS_RESULT_DAILY_BE_BLESS_TIMES_LIMIT BlessResult = 2
	BLESS_RESULT_BLESS_SAME_FRIENDS         BlessResult = 3
	BLESS_RESULT_CAN_NOT_BLESS_SELF         BlessResult = 4
)

type GrowResult int8

const (
	GROW_RESULT_SUCCESS                GrowResult = 0
	GROW_RESULT_HAS_GROWED             GrowResult = 1
	GROW_RESULT_FILED_NOT_OPEN         GrowResult = 2
	GROW_RESULT_HEART_LEVEL_NOT_ENOUGH GrowResult = 3
)

type TakeResult int8

const (
	TAKE_RESULT_SUCCESS        TakeResult = 0
	TAKE_RESULT_NOT_RIPE       TakeResult = 1
	TAKE_RESULT_FILED_NOT_OPEN TakeResult = 2
)

type UpgradeFiledResult int8

const (
	UPGRADE_FILED_RESULT_SUCCESS                UpgradeFiledResult = 0
	UPGRADE_FILED_RESULT_IS_BLACK               UpgradeFiledResult = 1
	UPGRADE_FILED_RESULT_FILED_NOT_OPEN         UpgradeFiledResult = 2
	UPGRADE_FILED_RESULT_HEART_LEVEL_NOT_ENOUGH UpgradeFiledResult = 3
)

type OpenFiledResult int8

const (
	OPEN_FILED_RESULT_SUCCESS                OpenFiledResult = 0
	OPEN_FILED_RESULT_FILED_HAS_OPEN         OpenFiledResult = 1
	OPEN_FILED_RESULT_HEART_LEVEL_NOT_ENOUGH OpenFiledResult = 2
)

type QuestFinishResult int8

const (
	QUEST_FINISH_RESULT_SUCCESS             QuestFinishResult = 0
	QUEST_FINISH_RESULT_QUEST_HAS_REFRESHED QuestFinishResult = 1
)

type QuestStateResult int8

const (
	QUEST_STATE_RESULT_CAN_NOT_FINISH QuestStateResult = 0
	QUEST_STATE_RESULT_CAN_FINISH     QuestStateResult = 1
	QUEST_STATE_RESULT_HAS_FINISHED   QuestStateResult = 2
)

type TaoyuanInfo_In struct {
}

func (this *TaoyuanInfo_In) Process(session *net.Session) {
	g_InHandler.TaoyuanInfo(session, this)
}

func (this *TaoyuanInfo_In) TypeName() string {
	return "taoyuan.taoyuan_info.in"
}

func (this *TaoyuanInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 39, 0
}

type TaoyuanInfo_Out struct {
	Level              int8                    `json:"level"`
	Exp                int64                   `json:"exp"`
	DailyBlessTimes    int32                   `json:"daily_bless_times"`
	DailyBeBlessTime   int32                   `json:"daily_be_bless_time"`
	QuestState         QuestStateResult        `json:"quest_state"`
	IsOpenWineBuilding int8                    `json:"is_open_wine_building"`
	IsOpenFoodBuilding int8                    `json:"is_open_food_building"`
	Queue              []TaoyuanInfo_Out_Queue `json:"queue"`
}

type TaoyuanInfo_Out_Queue struct {
	QueueId     int16 `json:"queue_id"`
	ItemId      int16 `json:"item_id"`
	ProductType int8  `json:"product_type"`
	StartTime   int64 `json:"start_time"`
	EndTime     int64 `json:"end_time"`
}

func (this *TaoyuanInfo_Out) Process(session *net.Session) {
	g_OutHandler.TaoyuanInfo(session, this)
}

func (this *TaoyuanInfo_Out) TypeName() string {
	return "taoyuan.taoyuan_info.out"
}

func (this *TaoyuanInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 39, 0
}

func (this *TaoyuanInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type Bless_In struct {
	OtherPid int64 `json:"other_pid"`
}

func (this *Bless_In) Process(session *net.Session) {
	g_InHandler.Bless(session, this)
}

func (this *Bless_In) TypeName() string {
	return "taoyuan.bless.in"
}

func (this *Bless_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 39, 1
}

type Bless_Out struct {
	Result BlessResult `json:"result"`
}

func (this *Bless_Out) Process(session *net.Session) {
	g_OutHandler.Bless(session, this)
}

func (this *Bless_Out) TypeName() string {
	return "taoyuan.bless.out"
}

func (this *Bless_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 39, 1
}

func (this *Bless_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ShopBuy_In struct {
	ItemId int16 `json:"item_id"`
	Num    int16 `json:"num"`
}

func (this *ShopBuy_In) Process(session *net.Session) {
	g_InHandler.ShopBuy(session, this)
}

func (this *ShopBuy_In) TypeName() string {
	return "taoyuan.shop_buy.in"
}

func (this *ShopBuy_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 39, 2
}

type ShopBuy_Out struct {
	Id int64 `json:"id"`
}

func (this *ShopBuy_Out) Process(session *net.Session) {
	g_OutHandler.ShopBuy(session, this)
}

func (this *ShopBuy_Out) TypeName() string {
	return "taoyuan.shop_buy.out"
}

func (this *ShopBuy_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 39, 2
}

func (this *ShopBuy_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ShopSell_In struct {
	Id  int64 `json:"id"`
	Num int16 `json:"num"`
}

func (this *ShopSell_In) Process(session *net.Session) {
	g_InHandler.ShopSell(session, this)
}

func (this *ShopSell_In) TypeName() string {
	return "taoyuan.shop_sell.in"
}

func (this *ShopSell_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 39, 3
}

type ShopSell_Out struct {
}

func (this *ShopSell_Out) Process(session *net.Session) {
	g_OutHandler.ShopSell(session, this)
}

func (this *ShopSell_Out) TypeName() string {
	return "taoyuan.shop_sell.out"
}

func (this *ShopSell_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 39, 3
}

func (this *ShopSell_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetAllItems_In struct {
}

func (this *GetAllItems_In) Process(session *net.Session) {
	g_InHandler.GetAllItems(session, this)
}

func (this *GetAllItems_In) TypeName() string {
	return "taoyuan.get_all_items.in"
}

func (this *GetAllItems_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 39, 4
}

type GetAllItems_Out struct {
	Items      []GetAllItems_Out_Items      `json:"items"`
	BuyRecords []GetAllItems_Out_BuyRecords `json:"buy_records"`
}

type GetAllItems_Out_Items struct {
	Id     int64 `json:"id"`
	ItemId int16 `json:"item_id"`
	Num    int16 `json:"num"`
}

type GetAllItems_Out_BuyRecords struct {
	ItemId int16 `json:"item_id"`
	Num    int16 `json:"num"`
}

func (this *GetAllItems_Out) Process(session *net.Session) {
	g_OutHandler.GetAllItems(session, this)
}

func (this *GetAllItems_Out) TypeName() string {
	return "taoyuan.get_all_items.out"
}

func (this *GetAllItems_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 39, 4
}

func (this *GetAllItems_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type FiledsInfo_In struct {
}

func (this *FiledsInfo_In) Process(session *net.Session) {
	g_InHandler.FiledsInfo(session, this)
}

func (this *FiledsInfo_In) TypeName() string {
	return "taoyuan.fileds_info.in"
}

func (this *FiledsInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 39, 5
}

type FiledsInfo_Out struct {
	Fileds []FiledsInfo_Out_Fileds `json:"fileds"`
}

type FiledsInfo_Out_Fileds struct {
	FiledId     int16 `json:"filed_id"`
	FiledStatus int16 `json:"filed_status"`
	PlantId     int16 `json:"plant_id"`
	GrowTime    int64 `json:"grow_time"`
	CanTake     bool  `json:"can_take"`
}

func (this *FiledsInfo_Out) Process(session *net.Session) {
	g_OutHandler.FiledsInfo(session, this)
}

func (this *FiledsInfo_Out) TypeName() string {
	return "taoyuan.fileds_info.out"
}

func (this *FiledsInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 39, 5
}

func (this *FiledsInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GrowPlant_In struct {
	FiledId int16 `json:"filed_id"`
	PlantId int16 `json:"plant_id"`
}

func (this *GrowPlant_In) Process(session *net.Session) {
	g_InHandler.GrowPlant(session, this)
}

func (this *GrowPlant_In) TypeName() string {
	return "taoyuan.grow_plant.in"
}

func (this *GrowPlant_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 39, 6
}

type GrowPlant_Out struct {
	FiledId int16      `json:"filed_id"`
	Result  GrowResult `json:"result"`
}

func (this *GrowPlant_Out) Process(session *net.Session) {
	g_OutHandler.GrowPlant(session, this)
}

func (this *GrowPlant_Out) TypeName() string {
	return "taoyuan.grow_plant.out"
}

func (this *GrowPlant_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 39, 6
}

func (this *GrowPlant_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type TakePlant_In struct {
	FiledId int16 `json:"filed_id"`
}

func (this *TakePlant_In) Process(session *net.Session) {
	g_InHandler.TakePlant(session, this)
}

func (this *TakePlant_In) TypeName() string {
	return "taoyuan.take_plant.in"
}

func (this *TakePlant_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 39, 7
}

type TakePlant_Out struct {
	FiledId int16      `json:"filed_id"`
	Result  TakeResult `json:"result"`
}

func (this *TakePlant_Out) Process(session *net.Session) {
	g_OutHandler.TakePlant(session, this)
}

func (this *TakePlant_Out) TypeName() string {
	return "taoyuan.take_plant.out"
}

func (this *TakePlant_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 39, 7
}

func (this *TakePlant_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type UpgradeFiled_In struct {
	FiledId int16 `json:"filed_id"`
}

func (this *UpgradeFiled_In) Process(session *net.Session) {
	g_InHandler.UpgradeFiled(session, this)
}

func (this *UpgradeFiled_In) TypeName() string {
	return "taoyuan.upgrade_filed.in"
}

func (this *UpgradeFiled_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 39, 8
}

type UpgradeFiled_Out struct {
	FiledId int16              `json:"filed_id"`
	Result  UpgradeFiledResult `json:"result"`
}

func (this *UpgradeFiled_Out) Process(session *net.Session) {
	g_OutHandler.UpgradeFiled(session, this)
}

func (this *UpgradeFiled_Out) TypeName() string {
	return "taoyuan.upgrade_filed.out"
}

func (this *UpgradeFiled_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 39, 8
}

func (this *UpgradeFiled_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type OpenFiled_In struct {
	FiledId int16 `json:"filed_id"`
}

func (this *OpenFiled_In) Process(session *net.Session) {
	g_InHandler.OpenFiled(session, this)
}

func (this *OpenFiled_In) TypeName() string {
	return "taoyuan.open_filed.in"
}

func (this *OpenFiled_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 39, 9
}

type OpenFiled_Out struct {
	FiledId int16           `json:"filed_id"`
	Result  OpenFiledResult `json:"result"`
}

func (this *OpenFiled_Out) Process(session *net.Session) {
	g_OutHandler.OpenFiled(session, this)
}

func (this *OpenFiled_Out) TypeName() string {
	return "taoyuan.open_filed.out"
}

func (this *OpenFiled_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 39, 9
}

func (this *OpenFiled_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type StudySkill_In struct {
	SkillId int16 `json:"skill_id"`
}

func (this *StudySkill_In) Process(session *net.Session) {
	g_InHandler.StudySkill(session, this)
}

func (this *StudySkill_In) TypeName() string {
	return "taoyuan.study_skill.in"
}

func (this *StudySkill_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 39, 10
}

type StudySkill_Out struct {
}

func (this *StudySkill_Out) Process(session *net.Session) {
	g_OutHandler.StudySkill(session, this)
}

func (this *StudySkill_Out) TypeName() string {
	return "taoyuan.study_skill.out"
}

func (this *StudySkill_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 39, 10
}

func (this *StudySkill_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type MakeProduct_In struct {
	ItemId      int16 `json:"item_id"`
	ProductType int8  `json:"product_type"`
}

func (this *MakeProduct_In) Process(session *net.Session) {
	g_InHandler.MakeProduct(session, this)
}

func (this *MakeProduct_In) TypeName() string {
	return "taoyuan.make_product.in"
}

func (this *MakeProduct_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 39, 11
}

type MakeProduct_Out struct {
	Queue []MakeProduct_Out_Queue `json:"queue"`
}

type MakeProduct_Out_Queue struct {
	QueueId     int16 `json:"queue_id"`
	ItemId      int16 `json:"item_id"`
	ProductType int8  `json:"product_type"`
	StartTime   int64 `json:"start_time"`
	EndTime     int64 `json:"end_time"`
}

func (this *MakeProduct_Out) Process(session *net.Session) {
	g_OutHandler.MakeProduct(session, this)
}

func (this *MakeProduct_Out) TypeName() string {
	return "taoyuan.make_product.out"
}

func (this *MakeProduct_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 39, 11
}

func (this *MakeProduct_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type TakeProduct_In struct {
	ProductType int8 `json:"product_type"`
	IsIngot     bool `json:"is_ingot"`
}

func (this *TakeProduct_In) Process(session *net.Session) {
	g_InHandler.TakeProduct(session, this)
}

func (this *TakeProduct_In) TypeName() string {
	return "taoyuan.take_product.in"
}

func (this *TakeProduct_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 39, 12
}

type TakeProduct_Out struct {
}

func (this *TakeProduct_Out) Process(session *net.Session) {
	g_OutHandler.TakeProduct(session, this)
}

func (this *TakeProduct_Out) TypeName() string {
	return "taoyuan.take_product.out"
}

func (this *TakeProduct_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 39, 12
}

func (this *TakeProduct_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ProductMakeQueue_In struct {
	ProductType int8 `json:"product_type"`
}

func (this *ProductMakeQueue_In) Process(session *net.Session) {
	g_InHandler.ProductMakeQueue(session, this)
}

func (this *ProductMakeQueue_In) TypeName() string {
	return "taoyuan.product_make_queue.in"
}

func (this *ProductMakeQueue_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 39, 13
}

type ProductMakeQueue_Out struct {
	SkillId int16                        `json:"skill_id"`
	Queue   []ProductMakeQueue_Out_Queue `json:"queue"`
}

type ProductMakeQueue_Out_Queue struct {
	QueueId     int16 `json:"queue_id"`
	ItemId      int16 `json:"item_id"`
	ProductType int8  `json:"product_type"`
	StartTime   int64 `json:"start_time"`
	EndTime     int64 `json:"end_time"`
}

func (this *ProductMakeQueue_Out) Process(session *net.Session) {
	g_OutHandler.ProductMakeQueue(session, this)
}

func (this *ProductMakeQueue_Out) TypeName() string {
	return "taoyuan.product_make_queue.out"
}

func (this *ProductMakeQueue_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 39, 13
}

func (this *ProductMakeQueue_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type QuestInfo_In struct {
}

func (this *QuestInfo_In) Process(session *net.Session) {
	g_InHandler.QuestInfo(session, this)
}

func (this *QuestInfo_In) TypeName() string {
	return "taoyuan.quest_info.in"
}

func (this *QuestInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 39, 14
}

type QuestInfo_Out struct {
	Quset []QuestInfo_Out_Quset `json:"quset"`
}

type QuestInfo_Out_Quset struct {
	QuestId    int16 `json:"quest_id"`
	ItemId     int16 `json:"item_id"`
	ItemNum    int16 `json:"item_num"`
	Exp        int64 `json:"exp"`
	Coins      int64 `json:"coins"`
	FinishTime int64 `json:"finish_time"`
}

func (this *QuestInfo_Out) Process(session *net.Session) {
	g_OutHandler.QuestInfo(session, this)
}

func (this *QuestInfo_Out) TypeName() string {
	return "taoyuan.quest_info.out"
}

func (this *QuestInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 39, 14
}

func (this *QuestInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type QuestFinish_In struct {
	QusetId int16 `json:"quset_id"`
}

func (this *QuestFinish_In) Process(session *net.Session) {
	g_InHandler.QuestFinish(session, this)
}

func (this *QuestFinish_In) TypeName() string {
	return "taoyuan.quest_finish.in"
}

func (this *QuestFinish_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 39, 15
}

type QuestFinish_Out struct {
	QuestId int16             `json:"quest_id"`
	Result  QuestFinishResult `json:"result"`
}

func (this *QuestFinish_Out) Process(session *net.Session) {
	g_OutHandler.QuestFinish(session, this)
}

func (this *QuestFinish_Out) TypeName() string {
	return "taoyuan.quest_finish.out"
}

func (this *QuestFinish_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 39, 15
}

func (this *QuestFinish_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type QuestRefresh_In struct {
}

func (this *QuestRefresh_In) Process(session *net.Session) {
	g_InHandler.QuestRefresh(session, this)
}

func (this *QuestRefresh_In) TypeName() string {
	return "taoyuan.quest_refresh.in"
}

func (this *QuestRefresh_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 39, 16
}

type QuestRefresh_Out struct {
}

func (this *QuestRefresh_Out) Process(session *net.Session) {
	g_OutHandler.QuestRefresh(session, this)
}

func (this *QuestRefresh_Out) TypeName() string {
	return "taoyuan.quest_refresh.out"
}

func (this *QuestRefresh_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 39, 16
}

func (this *QuestRefresh_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type FriendTaoyuanInfo_In struct {
	Pid int64 `json:"pid"`
}

func (this *FriendTaoyuanInfo_In) Process(session *net.Session) {
	g_InHandler.FriendTaoyuanInfo(session, this)
}

func (this *FriendTaoyuanInfo_In) TypeName() string {
	return "taoyuan.friend_taoyuan_info.in"
}

func (this *FriendTaoyuanInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 39, 17
}

type FriendTaoyuanInfo_Out struct {
	Level          int16                          `json:"level"`
	BeBlessedTimes int32                          `json:"be_blessed_times"`
	Pid            int64                          `json:"pid"`
	Nick           []byte                         `json:"nick"`
	IsBlessed      bool                           `json:"is_blessed"`
	Fileds         []FriendTaoyuanInfo_Out_Fileds `json:"fileds"`
}

type FriendTaoyuanInfo_Out_Fileds struct {
	FiledId     int16 `json:"filed_id"`
	FiledStatus int16 `json:"filed_status"`
	PlantId     int16 `json:"plant_id"`
	GrowTime    int64 `json:"grow_time"`
	CanTake     bool  `json:"can_take"`
}

func (this *FriendTaoyuanInfo_Out) Process(session *net.Session) {
	g_OutHandler.FriendTaoyuanInfo(session, this)
}

func (this *FriendTaoyuanInfo_Out) TypeName() string {
	return "taoyuan.friend_taoyuan_info.out"
}

func (this *FriendTaoyuanInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 39, 17
}

func (this *FriendTaoyuanInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type SkillInfo_In struct {
}

func (this *SkillInfo_In) Process(session *net.Session) {
	g_InHandler.SkillInfo(session, this)
}

func (this *SkillInfo_In) TypeName() string {
	return "taoyuan.skill_info.in"
}

func (this *SkillInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 39, 18
}

type SkillInfo_Out struct {
	SkillInfo []SkillInfo_Out_SkillInfo `json:"skill_info"`
}

type SkillInfo_Out_SkillInfo struct {
	SkillId   int16 `json:"skill_id"`
	SkillType int8  `json:"skill_type"`
	MakeTimes int64 `json:"make_times"`
}

func (this *SkillInfo_Out) Process(session *net.Session) {
	g_OutHandler.SkillInfo(session, this)
}

func (this *SkillInfo_Out) TypeName() string {
	return "taoyuan.skill_info.out"
}

func (this *SkillInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 39, 18
}

func (this *SkillInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type OpenQueue_In struct {
	ProductType int8 `json:"product_type"`
}

func (this *OpenQueue_In) Process(session *net.Session) {
	g_InHandler.OpenQueue(session, this)
}

func (this *OpenQueue_In) TypeName() string {
	return "taoyuan.open_queue.in"
}

func (this *OpenQueue_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 39, 19
}

type OpenQueue_Out struct {
}

func (this *OpenQueue_Out) Process(session *net.Session) {
	g_OutHandler.OpenQueue(session, this)
}

func (this *OpenQueue_Out) TypeName() string {
	return "taoyuan.open_queue.out"
}

func (this *OpenQueue_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 39, 19
}

func (this *OpenQueue_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type PlantQuicklyMaturity_In struct {
	FiledId int16 `json:"filed_id"`
}

func (this *PlantQuicklyMaturity_In) Process(session *net.Session) {
	g_InHandler.PlantQuicklyMaturity(session, this)
}

func (this *PlantQuicklyMaturity_In) TypeName() string {
	return "taoyuan.plant_quickly_maturity.in"
}

func (this *PlantQuicklyMaturity_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 39, 20
}

type PlantQuicklyMaturity_Out struct {
	FiledId int16 `json:"filed_id"`
}

func (this *PlantQuicklyMaturity_Out) Process(session *net.Session) {
	g_OutHandler.PlantQuicklyMaturity(session, this)
}

func (this *PlantQuicklyMaturity_Out) TypeName() string {
	return "taoyuan.plant_quickly_maturity.out"
}

func (this *PlantQuicklyMaturity_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 39, 20
}

func (this *PlantQuicklyMaturity_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type TaoyuanMessageInfo_In struct {
}

func (this *TaoyuanMessageInfo_In) Process(session *net.Session) {
	g_InHandler.TaoyuanMessageInfo(session, this)
}

func (this *TaoyuanMessageInfo_In) TypeName() string {
	return "taoyuan.taoyuan_message_info.in"
}

func (this *TaoyuanMessageInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 39, 21
}

type TaoyuanMessageInfo_Out struct {
	MessageInfo []TaoyuanMessageInfo_Out_MessageInfo `json:"message_info"`
}

type TaoyuanMessageInfo_Out_MessageInfo struct {
	Id   int64  `json:"id"`
	Nick []byte `json:"nick"`
	Exp  int32  `json:"exp"`
}

func (this *TaoyuanMessageInfo_Out) Process(session *net.Session) {
	g_OutHandler.TaoyuanMessageInfo(session, this)
}

func (this *TaoyuanMessageInfo_Out) TypeName() string {
	return "taoyuan.taoyuan_message_info.out"
}

func (this *TaoyuanMessageInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 39, 21
}

func (this *TaoyuanMessageInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type TaoyuanMessageRead_In struct {
	Id int64 `json:"id"`
}

func (this *TaoyuanMessageRead_In) Process(session *net.Session) {
	g_InHandler.TaoyuanMessageRead(session, this)
}

func (this *TaoyuanMessageRead_In) TypeName() string {
	return "taoyuan.taoyuan_message_read.in"
}

func (this *TaoyuanMessageRead_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 39, 22
}

type TaoyuanMessageRead_Out struct {
}

func (this *TaoyuanMessageRead_Out) Process(session *net.Session) {
	g_OutHandler.TaoyuanMessageRead(session, this)
}

func (this *TaoyuanMessageRead_Out) TypeName() string {
	return "taoyuan.taoyuan_message_read.out"
}

func (this *TaoyuanMessageRead_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 39, 22
}

func (this *TaoyuanMessageRead_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type OpenProductBuilding_In struct {
	ProductType int8 `json:"product_type"`
}

func (this *OpenProductBuilding_In) Process(session *net.Session) {
	g_InHandler.OpenProductBuilding(session, this)
}

func (this *OpenProductBuilding_In) TypeName() string {
	return "taoyuan.open_product_building.in"
}

func (this *OpenProductBuilding_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 39, 23
}

type OpenProductBuilding_Out struct {
}

func (this *OpenProductBuilding_Out) Process(session *net.Session) {
	g_OutHandler.OpenProductBuilding(session, this)
}

func (this *OpenProductBuilding_Out) TypeName() string {
	return "taoyuan.open_product_building.out"
}

func (this *OpenProductBuilding_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 39, 23
}

func (this *OpenProductBuilding_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

func (this *TaoyuanInfo_In) Decode(buffer *net.Buffer) {
}

func (this *TaoyuanInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(39)
	buffer.WriteUint8(0)
}

func (this *TaoyuanInfo_In) ByteSize() int {
	size := 2
	return size
}

func (this *TaoyuanInfo_Out) Decode(buffer *net.Buffer) {
	this.Level = int8(buffer.ReadUint8())
	this.Exp = int64(buffer.ReadUint64LE())
	this.DailyBlessTimes = int32(buffer.ReadUint32LE())
	this.DailyBeBlessTime = int32(buffer.ReadUint32LE())
	this.QuestState = QuestStateResult(buffer.ReadUint8())
	this.IsOpenWineBuilding = int8(buffer.ReadUint8())
	this.IsOpenFoodBuilding = int8(buffer.ReadUint8())
	this.Queue = make([]TaoyuanInfo_Out_Queue, buffer.ReadUint8())
	for i := 0; i < len(this.Queue); i++ {
		this.Queue[i].Decode(buffer)
	}
}

func (this *TaoyuanInfo_Out_Queue) Decode(buffer *net.Buffer) {
	this.QueueId = int16(buffer.ReadUint16LE())
	this.ItemId = int16(buffer.ReadUint16LE())
	this.ProductType = int8(buffer.ReadUint8())
	this.StartTime = int64(buffer.ReadUint64LE())
	this.EndTime = int64(buffer.ReadUint64LE())
}

func (this *TaoyuanInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(39)
	buffer.WriteUint8(0)
	buffer.WriteUint8(uint8(this.Level))
	buffer.WriteUint64LE(uint64(this.Exp))
	buffer.WriteUint32LE(uint32(this.DailyBlessTimes))
	buffer.WriteUint32LE(uint32(this.DailyBeBlessTime))
	buffer.WriteUint8(uint8(this.QuestState))
	buffer.WriteUint8(uint8(this.IsOpenWineBuilding))
	buffer.WriteUint8(uint8(this.IsOpenFoodBuilding))
	buffer.WriteUint8(uint8(len(this.Queue)))
	for i := 0; i < len(this.Queue); i++ {
		this.Queue[i].Encode(buffer)
	}
}

func (this *TaoyuanInfo_Out_Queue) Encode(buffer *net.Buffer) {
	buffer.WriteUint16LE(uint16(this.QueueId))
	buffer.WriteUint16LE(uint16(this.ItemId))
	buffer.WriteUint8(uint8(this.ProductType))
	buffer.WriteUint64LE(uint64(this.StartTime))
	buffer.WriteUint64LE(uint64(this.EndTime))
}

func (this *TaoyuanInfo_Out) ByteSize() int {
	size := 23
	size += len(this.Queue) * 21
	return size
}

func (this *Bless_In) Decode(buffer *net.Buffer) {
	this.OtherPid = int64(buffer.ReadUint64LE())
}

func (this *Bless_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(39)
	buffer.WriteUint8(1)
	buffer.WriteUint64LE(uint64(this.OtherPid))
}

func (this *Bless_In) ByteSize() int {
	size := 10
	return size
}

func (this *Bless_Out) Decode(buffer *net.Buffer) {
	this.Result = BlessResult(buffer.ReadUint8())
}

func (this *Bless_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(39)
	buffer.WriteUint8(1)
	buffer.WriteUint8(uint8(this.Result))
}

func (this *Bless_Out) ByteSize() int {
	size := 3
	return size
}

func (this *ShopBuy_In) Decode(buffer *net.Buffer) {
	this.ItemId = int16(buffer.ReadUint16LE())
	this.Num = int16(buffer.ReadUint16LE())
}

func (this *ShopBuy_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(39)
	buffer.WriteUint8(2)
	buffer.WriteUint16LE(uint16(this.ItemId))
	buffer.WriteUint16LE(uint16(this.Num))
}

func (this *ShopBuy_In) ByteSize() int {
	size := 6
	return size
}

func (this *ShopBuy_Out) Decode(buffer *net.Buffer) {
	this.Id = int64(buffer.ReadUint64LE())
}

func (this *ShopBuy_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(39)
	buffer.WriteUint8(2)
	buffer.WriteUint64LE(uint64(this.Id))
}

func (this *ShopBuy_Out) ByteSize() int {
	size := 10
	return size
}

func (this *ShopSell_In) Decode(buffer *net.Buffer) {
	this.Id = int64(buffer.ReadUint64LE())
	this.Num = int16(buffer.ReadUint16LE())
}

func (this *ShopSell_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(39)
	buffer.WriteUint8(3)
	buffer.WriteUint64LE(uint64(this.Id))
	buffer.WriteUint16LE(uint16(this.Num))
}

func (this *ShopSell_In) ByteSize() int {
	size := 12
	return size
}

func (this *ShopSell_Out) Decode(buffer *net.Buffer) {
}

func (this *ShopSell_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(39)
	buffer.WriteUint8(3)
}

func (this *ShopSell_Out) ByteSize() int {
	size := 2
	return size
}

func (this *GetAllItems_In) Decode(buffer *net.Buffer) {
}

func (this *GetAllItems_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(39)
	buffer.WriteUint8(4)
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
	this.BuyRecords = make([]GetAllItems_Out_BuyRecords, buffer.ReadUint8())
	for i := 0; i < len(this.BuyRecords); i++ {
		this.BuyRecords[i].Decode(buffer)
	}
}

func (this *GetAllItems_Out_Items) Decode(buffer *net.Buffer) {
	this.Id = int64(buffer.ReadUint64LE())
	this.ItemId = int16(buffer.ReadUint16LE())
	this.Num = int16(buffer.ReadUint16LE())
}

func (this *GetAllItems_Out_BuyRecords) Decode(buffer *net.Buffer) {
	this.ItemId = int16(buffer.ReadUint16LE())
	this.Num = int16(buffer.ReadUint16LE())
}

func (this *GetAllItems_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(39)
	buffer.WriteUint8(4)
	buffer.WriteUint8(uint8(len(this.Items)))
	for i := 0; i < len(this.Items); i++ {
		this.Items[i].Encode(buffer)
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
}

func (this *GetAllItems_Out_BuyRecords) Encode(buffer *net.Buffer) {
	buffer.WriteUint16LE(uint16(this.ItemId))
	buffer.WriteUint16LE(uint16(this.Num))
}

func (this *GetAllItems_Out) ByteSize() int {
	size := 4
	size += len(this.Items) * 12
	size += len(this.BuyRecords) * 4
	return size
}

func (this *FiledsInfo_In) Decode(buffer *net.Buffer) {
}

func (this *FiledsInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(39)
	buffer.WriteUint8(5)
}

func (this *FiledsInfo_In) ByteSize() int {
	size := 2
	return size
}

func (this *FiledsInfo_Out) Decode(buffer *net.Buffer) {
	this.Fileds = make([]FiledsInfo_Out_Fileds, buffer.ReadUint8())
	for i := 0; i < len(this.Fileds); i++ {
		this.Fileds[i].Decode(buffer)
	}
}

func (this *FiledsInfo_Out_Fileds) Decode(buffer *net.Buffer) {
	this.FiledId = int16(buffer.ReadUint16LE())
	this.FiledStatus = int16(buffer.ReadUint16LE())
	this.PlantId = int16(buffer.ReadUint16LE())
	this.GrowTime = int64(buffer.ReadUint64LE())
	this.CanTake = buffer.ReadUint8() == 1
}

func (this *FiledsInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(39)
	buffer.WriteUint8(5)
	buffer.WriteUint8(uint8(len(this.Fileds)))
	for i := 0; i < len(this.Fileds); i++ {
		this.Fileds[i].Encode(buffer)
	}
}

func (this *FiledsInfo_Out_Fileds) Encode(buffer *net.Buffer) {
	buffer.WriteUint16LE(uint16(this.FiledId))
	buffer.WriteUint16LE(uint16(this.FiledStatus))
	buffer.WriteUint16LE(uint16(this.PlantId))
	buffer.WriteUint64LE(uint64(this.GrowTime))
	if this.CanTake {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *FiledsInfo_Out) ByteSize() int {
	size := 3
	size += len(this.Fileds) * 15
	return size
}

func (this *GrowPlant_In) Decode(buffer *net.Buffer) {
	this.FiledId = int16(buffer.ReadUint16LE())
	this.PlantId = int16(buffer.ReadUint16LE())
}

func (this *GrowPlant_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(39)
	buffer.WriteUint8(6)
	buffer.WriteUint16LE(uint16(this.FiledId))
	buffer.WriteUint16LE(uint16(this.PlantId))
}

func (this *GrowPlant_In) ByteSize() int {
	size := 6
	return size
}

func (this *GrowPlant_Out) Decode(buffer *net.Buffer) {
	this.FiledId = int16(buffer.ReadUint16LE())
	this.Result = GrowResult(buffer.ReadUint8())
}

func (this *GrowPlant_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(39)
	buffer.WriteUint8(6)
	buffer.WriteUint16LE(uint16(this.FiledId))
	buffer.WriteUint8(uint8(this.Result))
}

func (this *GrowPlant_Out) ByteSize() int {
	size := 5
	return size
}

func (this *TakePlant_In) Decode(buffer *net.Buffer) {
	this.FiledId = int16(buffer.ReadUint16LE())
}

func (this *TakePlant_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(39)
	buffer.WriteUint8(7)
	buffer.WriteUint16LE(uint16(this.FiledId))
}

func (this *TakePlant_In) ByteSize() int {
	size := 4
	return size
}

func (this *TakePlant_Out) Decode(buffer *net.Buffer) {
	this.FiledId = int16(buffer.ReadUint16LE())
	this.Result = TakeResult(buffer.ReadUint8())
}

func (this *TakePlant_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(39)
	buffer.WriteUint8(7)
	buffer.WriteUint16LE(uint16(this.FiledId))
	buffer.WriteUint8(uint8(this.Result))
}

func (this *TakePlant_Out) ByteSize() int {
	size := 5
	return size
}

func (this *UpgradeFiled_In) Decode(buffer *net.Buffer) {
	this.FiledId = int16(buffer.ReadUint16LE())
}

func (this *UpgradeFiled_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(39)
	buffer.WriteUint8(8)
	buffer.WriteUint16LE(uint16(this.FiledId))
}

func (this *UpgradeFiled_In) ByteSize() int {
	size := 4
	return size
}

func (this *UpgradeFiled_Out) Decode(buffer *net.Buffer) {
	this.FiledId = int16(buffer.ReadUint16LE())
	this.Result = UpgradeFiledResult(buffer.ReadUint8())
}

func (this *UpgradeFiled_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(39)
	buffer.WriteUint8(8)
	buffer.WriteUint16LE(uint16(this.FiledId))
	buffer.WriteUint8(uint8(this.Result))
}

func (this *UpgradeFiled_Out) ByteSize() int {
	size := 5
	return size
}

func (this *OpenFiled_In) Decode(buffer *net.Buffer) {
	this.FiledId = int16(buffer.ReadUint16LE())
}

func (this *OpenFiled_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(39)
	buffer.WriteUint8(9)
	buffer.WriteUint16LE(uint16(this.FiledId))
}

func (this *OpenFiled_In) ByteSize() int {
	size := 4
	return size
}

func (this *OpenFiled_Out) Decode(buffer *net.Buffer) {
	this.FiledId = int16(buffer.ReadUint16LE())
	this.Result = OpenFiledResult(buffer.ReadUint8())
}

func (this *OpenFiled_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(39)
	buffer.WriteUint8(9)
	buffer.WriteUint16LE(uint16(this.FiledId))
	buffer.WriteUint8(uint8(this.Result))
}

func (this *OpenFiled_Out) ByteSize() int {
	size := 5
	return size
}

func (this *StudySkill_In) Decode(buffer *net.Buffer) {
	this.SkillId = int16(buffer.ReadUint16LE())
}

func (this *StudySkill_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(39)
	buffer.WriteUint8(10)
	buffer.WriteUint16LE(uint16(this.SkillId))
}

func (this *StudySkill_In) ByteSize() int {
	size := 4
	return size
}

func (this *StudySkill_Out) Decode(buffer *net.Buffer) {
}

func (this *StudySkill_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(39)
	buffer.WriteUint8(10)
}

func (this *StudySkill_Out) ByteSize() int {
	size := 2
	return size
}

func (this *MakeProduct_In) Decode(buffer *net.Buffer) {
	this.ItemId = int16(buffer.ReadUint16LE())
	this.ProductType = int8(buffer.ReadUint8())
}

func (this *MakeProduct_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(39)
	buffer.WriteUint8(11)
	buffer.WriteUint16LE(uint16(this.ItemId))
	buffer.WriteUint8(uint8(this.ProductType))
}

func (this *MakeProduct_In) ByteSize() int {
	size := 5
	return size
}

func (this *MakeProduct_Out) Decode(buffer *net.Buffer) {
	this.Queue = make([]MakeProduct_Out_Queue, buffer.ReadUint8())
	for i := 0; i < len(this.Queue); i++ {
		this.Queue[i].Decode(buffer)
	}
}

func (this *MakeProduct_Out_Queue) Decode(buffer *net.Buffer) {
	this.QueueId = int16(buffer.ReadUint16LE())
	this.ItemId = int16(buffer.ReadUint16LE())
	this.ProductType = int8(buffer.ReadUint8())
	this.StartTime = int64(buffer.ReadUint64LE())
	this.EndTime = int64(buffer.ReadUint64LE())
}

func (this *MakeProduct_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(39)
	buffer.WriteUint8(11)
	buffer.WriteUint8(uint8(len(this.Queue)))
	for i := 0; i < len(this.Queue); i++ {
		this.Queue[i].Encode(buffer)
	}
}

func (this *MakeProduct_Out_Queue) Encode(buffer *net.Buffer) {
	buffer.WriteUint16LE(uint16(this.QueueId))
	buffer.WriteUint16LE(uint16(this.ItemId))
	buffer.WriteUint8(uint8(this.ProductType))
	buffer.WriteUint64LE(uint64(this.StartTime))
	buffer.WriteUint64LE(uint64(this.EndTime))
}

func (this *MakeProduct_Out) ByteSize() int {
	size := 3
	size += len(this.Queue) * 21
	return size
}

func (this *TakeProduct_In) Decode(buffer *net.Buffer) {
	this.ProductType = int8(buffer.ReadUint8())
	this.IsIngot = buffer.ReadUint8() == 1
}

func (this *TakeProduct_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(39)
	buffer.WriteUint8(12)
	buffer.WriteUint8(uint8(this.ProductType))
	if this.IsIngot {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *TakeProduct_In) ByteSize() int {
	size := 4
	return size
}

func (this *TakeProduct_Out) Decode(buffer *net.Buffer) {
}

func (this *TakeProduct_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(39)
	buffer.WriteUint8(12)
}

func (this *TakeProduct_Out) ByteSize() int {
	size := 2
	return size
}

func (this *ProductMakeQueue_In) Decode(buffer *net.Buffer) {
	this.ProductType = int8(buffer.ReadUint8())
}

func (this *ProductMakeQueue_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(39)
	buffer.WriteUint8(13)
	buffer.WriteUint8(uint8(this.ProductType))
}

func (this *ProductMakeQueue_In) ByteSize() int {
	size := 3
	return size
}

func (this *ProductMakeQueue_Out) Decode(buffer *net.Buffer) {
	this.SkillId = int16(buffer.ReadUint16LE())
	this.Queue = make([]ProductMakeQueue_Out_Queue, buffer.ReadUint8())
	for i := 0; i < len(this.Queue); i++ {
		this.Queue[i].Decode(buffer)
	}
}

func (this *ProductMakeQueue_Out_Queue) Decode(buffer *net.Buffer) {
	this.QueueId = int16(buffer.ReadUint16LE())
	this.ItemId = int16(buffer.ReadUint16LE())
	this.ProductType = int8(buffer.ReadUint8())
	this.StartTime = int64(buffer.ReadUint64LE())
	this.EndTime = int64(buffer.ReadUint64LE())
}

func (this *ProductMakeQueue_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(39)
	buffer.WriteUint8(13)
	buffer.WriteUint16LE(uint16(this.SkillId))
	buffer.WriteUint8(uint8(len(this.Queue)))
	for i := 0; i < len(this.Queue); i++ {
		this.Queue[i].Encode(buffer)
	}
}

func (this *ProductMakeQueue_Out_Queue) Encode(buffer *net.Buffer) {
	buffer.WriteUint16LE(uint16(this.QueueId))
	buffer.WriteUint16LE(uint16(this.ItemId))
	buffer.WriteUint8(uint8(this.ProductType))
	buffer.WriteUint64LE(uint64(this.StartTime))
	buffer.WriteUint64LE(uint64(this.EndTime))
}

func (this *ProductMakeQueue_Out) ByteSize() int {
	size := 5
	size += len(this.Queue) * 21
	return size
}

func (this *QuestInfo_In) Decode(buffer *net.Buffer) {
}

func (this *QuestInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(39)
	buffer.WriteUint8(14)
}

func (this *QuestInfo_In) ByteSize() int {
	size := 2
	return size
}

func (this *QuestInfo_Out) Decode(buffer *net.Buffer) {
	this.Quset = make([]QuestInfo_Out_Quset, buffer.ReadUint8())
	for i := 0; i < len(this.Quset); i++ {
		this.Quset[i].Decode(buffer)
	}
}

func (this *QuestInfo_Out_Quset) Decode(buffer *net.Buffer) {
	this.QuestId = int16(buffer.ReadUint16LE())
	this.ItemId = int16(buffer.ReadUint16LE())
	this.ItemNum = int16(buffer.ReadUint16LE())
	this.Exp = int64(buffer.ReadUint64LE())
	this.Coins = int64(buffer.ReadUint64LE())
	this.FinishTime = int64(buffer.ReadUint64LE())
}

func (this *QuestInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(39)
	buffer.WriteUint8(14)
	buffer.WriteUint8(uint8(len(this.Quset)))
	for i := 0; i < len(this.Quset); i++ {
		this.Quset[i].Encode(buffer)
	}
}

func (this *QuestInfo_Out_Quset) Encode(buffer *net.Buffer) {
	buffer.WriteUint16LE(uint16(this.QuestId))
	buffer.WriteUint16LE(uint16(this.ItemId))
	buffer.WriteUint16LE(uint16(this.ItemNum))
	buffer.WriteUint64LE(uint64(this.Exp))
	buffer.WriteUint64LE(uint64(this.Coins))
	buffer.WriteUint64LE(uint64(this.FinishTime))
}

func (this *QuestInfo_Out) ByteSize() int {
	size := 3
	size += len(this.Quset) * 30
	return size
}

func (this *QuestFinish_In) Decode(buffer *net.Buffer) {
	this.QusetId = int16(buffer.ReadUint16LE())
}

func (this *QuestFinish_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(39)
	buffer.WriteUint8(15)
	buffer.WriteUint16LE(uint16(this.QusetId))
}

func (this *QuestFinish_In) ByteSize() int {
	size := 4
	return size
}

func (this *QuestFinish_Out) Decode(buffer *net.Buffer) {
	this.QuestId = int16(buffer.ReadUint16LE())
	this.Result = QuestFinishResult(buffer.ReadUint8())
}

func (this *QuestFinish_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(39)
	buffer.WriteUint8(15)
	buffer.WriteUint16LE(uint16(this.QuestId))
	buffer.WriteUint8(uint8(this.Result))
}

func (this *QuestFinish_Out) ByteSize() int {
	size := 5
	return size
}

func (this *QuestRefresh_In) Decode(buffer *net.Buffer) {
}

func (this *QuestRefresh_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(39)
	buffer.WriteUint8(16)
}

func (this *QuestRefresh_In) ByteSize() int {
	size := 2
	return size
}

func (this *QuestRefresh_Out) Decode(buffer *net.Buffer) {
}

func (this *QuestRefresh_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(39)
	buffer.WriteUint8(16)
}

func (this *QuestRefresh_Out) ByteSize() int {
	size := 2
	return size
}

func (this *FriendTaoyuanInfo_In) Decode(buffer *net.Buffer) {
	this.Pid = int64(buffer.ReadUint64LE())
}

func (this *FriendTaoyuanInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(39)
	buffer.WriteUint8(17)
	buffer.WriteUint64LE(uint64(this.Pid))
}

func (this *FriendTaoyuanInfo_In) ByteSize() int {
	size := 10
	return size
}

func (this *FriendTaoyuanInfo_Out) Decode(buffer *net.Buffer) {
	this.Level = int16(buffer.ReadUint16LE())
	this.BeBlessedTimes = int32(buffer.ReadUint32LE())
	this.Pid = int64(buffer.ReadUint64LE())
	this.Nick = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.IsBlessed = buffer.ReadUint8() == 1
	this.Fileds = make([]FriendTaoyuanInfo_Out_Fileds, buffer.ReadUint8())
	for i := 0; i < len(this.Fileds); i++ {
		this.Fileds[i].Decode(buffer)
	}
}

func (this *FriendTaoyuanInfo_Out_Fileds) Decode(buffer *net.Buffer) {
	this.FiledId = int16(buffer.ReadUint16LE())
	this.FiledStatus = int16(buffer.ReadUint16LE())
	this.PlantId = int16(buffer.ReadUint16LE())
	this.GrowTime = int64(buffer.ReadUint64LE())
	this.CanTake = buffer.ReadUint8() == 1
}

func (this *FriendTaoyuanInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(39)
	buffer.WriteUint8(17)
	buffer.WriteUint16LE(uint16(this.Level))
	buffer.WriteUint32LE(uint32(this.BeBlessedTimes))
	buffer.WriteUint64LE(uint64(this.Pid))
	buffer.WriteUint16LE(uint16(len(this.Nick)))
	buffer.WriteBytes(this.Nick)
	if this.IsBlessed {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
	buffer.WriteUint8(uint8(len(this.Fileds)))
	for i := 0; i < len(this.Fileds); i++ {
		this.Fileds[i].Encode(buffer)
	}
}

func (this *FriendTaoyuanInfo_Out_Fileds) Encode(buffer *net.Buffer) {
	buffer.WriteUint16LE(uint16(this.FiledId))
	buffer.WriteUint16LE(uint16(this.FiledStatus))
	buffer.WriteUint16LE(uint16(this.PlantId))
	buffer.WriteUint64LE(uint64(this.GrowTime))
	if this.CanTake {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *FriendTaoyuanInfo_Out) ByteSize() int {
	size := 20
	size += len(this.Nick)
	size += len(this.Fileds) * 15
	return size
}

func (this *SkillInfo_In) Decode(buffer *net.Buffer) {
}

func (this *SkillInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(39)
	buffer.WriteUint8(18)
}

func (this *SkillInfo_In) ByteSize() int {
	size := 2
	return size
}

func (this *SkillInfo_Out) Decode(buffer *net.Buffer) {
	this.SkillInfo = make([]SkillInfo_Out_SkillInfo, buffer.ReadUint8())
	for i := 0; i < len(this.SkillInfo); i++ {
		this.SkillInfo[i].Decode(buffer)
	}
}

func (this *SkillInfo_Out_SkillInfo) Decode(buffer *net.Buffer) {
	this.SkillId = int16(buffer.ReadUint16LE())
	this.SkillType = int8(buffer.ReadUint8())
	this.MakeTimes = int64(buffer.ReadUint64LE())
}

func (this *SkillInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(39)
	buffer.WriteUint8(18)
	buffer.WriteUint8(uint8(len(this.SkillInfo)))
	for i := 0; i < len(this.SkillInfo); i++ {
		this.SkillInfo[i].Encode(buffer)
	}
}

func (this *SkillInfo_Out_SkillInfo) Encode(buffer *net.Buffer) {
	buffer.WriteUint16LE(uint16(this.SkillId))
	buffer.WriteUint8(uint8(this.SkillType))
	buffer.WriteUint64LE(uint64(this.MakeTimes))
}

func (this *SkillInfo_Out) ByteSize() int {
	size := 3
	size += len(this.SkillInfo) * 11
	return size
}

func (this *OpenQueue_In) Decode(buffer *net.Buffer) {
	this.ProductType = int8(buffer.ReadUint8())
}

func (this *OpenQueue_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(39)
	buffer.WriteUint8(19)
	buffer.WriteUint8(uint8(this.ProductType))
}

func (this *OpenQueue_In) ByteSize() int {
	size := 3
	return size
}

func (this *OpenQueue_Out) Decode(buffer *net.Buffer) {
}

func (this *OpenQueue_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(39)
	buffer.WriteUint8(19)
}

func (this *OpenQueue_Out) ByteSize() int {
	size := 2
	return size
}

func (this *PlantQuicklyMaturity_In) Decode(buffer *net.Buffer) {
	this.FiledId = int16(buffer.ReadUint16LE())
}

func (this *PlantQuicklyMaturity_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(39)
	buffer.WriteUint8(20)
	buffer.WriteUint16LE(uint16(this.FiledId))
}

func (this *PlantQuicklyMaturity_In) ByteSize() int {
	size := 4
	return size
}

func (this *PlantQuicklyMaturity_Out) Decode(buffer *net.Buffer) {
	this.FiledId = int16(buffer.ReadUint16LE())
}

func (this *PlantQuicklyMaturity_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(39)
	buffer.WriteUint8(20)
	buffer.WriteUint16LE(uint16(this.FiledId))
}

func (this *PlantQuicklyMaturity_Out) ByteSize() int {
	size := 4
	return size
}

func (this *TaoyuanMessageInfo_In) Decode(buffer *net.Buffer) {
}

func (this *TaoyuanMessageInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(39)
	buffer.WriteUint8(21)
}

func (this *TaoyuanMessageInfo_In) ByteSize() int {
	size := 2
	return size
}

func (this *TaoyuanMessageInfo_Out) Decode(buffer *net.Buffer) {
	this.MessageInfo = make([]TaoyuanMessageInfo_Out_MessageInfo, buffer.ReadUint8())
	for i := 0; i < len(this.MessageInfo); i++ {
		this.MessageInfo[i].Decode(buffer)
	}
}

func (this *TaoyuanMessageInfo_Out_MessageInfo) Decode(buffer *net.Buffer) {
	this.Id = int64(buffer.ReadUint64LE())
	this.Nick = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Exp = int32(buffer.ReadUint32LE())
}

func (this *TaoyuanMessageInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(39)
	buffer.WriteUint8(21)
	buffer.WriteUint8(uint8(len(this.MessageInfo)))
	for i := 0; i < len(this.MessageInfo); i++ {
		this.MessageInfo[i].Encode(buffer)
	}
}

func (this *TaoyuanMessageInfo_Out_MessageInfo) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.Id))
	buffer.WriteUint16LE(uint16(len(this.Nick)))
	buffer.WriteBytes(this.Nick)
	buffer.WriteUint32LE(uint32(this.Exp))
}

func (this *TaoyuanMessageInfo_Out) ByteSize() int {
	size := 3
	for i := 0; i < len(this.MessageInfo); i++ {
		size += this.MessageInfo[i].ByteSize()
	}
	return size
}

func (this *TaoyuanMessageInfo_Out_MessageInfo) ByteSize() int {
	size := 14
	size += len(this.Nick)
	return size
}

func (this *TaoyuanMessageRead_In) Decode(buffer *net.Buffer) {
	this.Id = int64(buffer.ReadUint64LE())
}

func (this *TaoyuanMessageRead_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(39)
	buffer.WriteUint8(22)
	buffer.WriteUint64LE(uint64(this.Id))
}

func (this *TaoyuanMessageRead_In) ByteSize() int {
	size := 10
	return size
}

func (this *TaoyuanMessageRead_Out) Decode(buffer *net.Buffer) {
}

func (this *TaoyuanMessageRead_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(39)
	buffer.WriteUint8(22)
}

func (this *TaoyuanMessageRead_Out) ByteSize() int {
	size := 2
	return size
}

func (this *OpenProductBuilding_In) Decode(buffer *net.Buffer) {
	this.ProductType = int8(buffer.ReadUint8())
}

func (this *OpenProductBuilding_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(39)
	buffer.WriteUint8(23)
	buffer.WriteUint8(uint8(this.ProductType))
}

func (this *OpenProductBuilding_In) ByteSize() int {
	size := 3
	return size
}

func (this *OpenProductBuilding_Out) Decode(buffer *net.Buffer) {
}

func (this *OpenProductBuilding_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(39)
	buffer.WriteUint8(23)
}

func (this *OpenProductBuilding_Out) ByteSize() int {
	size := 2
	return size
}
