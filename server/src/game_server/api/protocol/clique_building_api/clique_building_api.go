package clique_building_api

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
	CliqueBaseDonate(*net.Session, *CliqueBaseDonate_In)
	CliqueBuildingStatus(*net.Session, *CliqueBuildingStatus_In)
	CliqueBankDonate(*net.Session, *CliqueBankDonate_In)
	CliqueBankBuy(*net.Session, *CliqueBankBuy_In)
	CliqueBankSold(*net.Session, *CliqueBankSold_In)
	CliqueKongfuDonate(*net.Session, *CliqueKongfuDonate_In)
	CliqueKongfuInfo(*net.Session, *CliqueKongfuInfo_In)
	CliqueKongfuTrain(*net.Session, *CliqueKongfuTrain_In)
	CliqueTempleWorship(*net.Session, *CliqueTempleWorship_In)
	CliqueTempleDonate(*net.Session, *CliqueTempleDonate_In)
	CliqueTempleInfo(*net.Session, *CliqueTempleInfo_In)
	CliqueStoreDonate(*net.Session, *CliqueStoreDonate_In)
	CliqueStoreInfo(*net.Session, *CliqueStoreInfo_In)
	CliqueStoreSendChest(*net.Session, *CliqueStoreSendChest_In)
}

type OutHandler interface {
	CliqueBaseDonate(*net.Session, *CliqueBaseDonate_Out)
	CliqueBuildingStatus(*net.Session, *CliqueBuildingStatus_Out)
	CliqueBankDonate(*net.Session, *CliqueBankDonate_Out)
	CliqueBankBuy(*net.Session, *CliqueBankBuy_Out)
	CliqueBankSold(*net.Session, *CliqueBankSold_Out)
	CliqueKongfuDonate(*net.Session, *CliqueKongfuDonate_Out)
	CliqueKongfuInfo(*net.Session, *CliqueKongfuInfo_Out)
	CliqueKongfuTrain(*net.Session, *CliqueKongfuTrain_Out)
	CliqueTempleWorship(*net.Session, *CliqueTempleWorship_Out)
	CliqueTempleDonate(*net.Session, *CliqueTempleDonate_Out)
	CliqueTempleInfo(*net.Session, *CliqueTempleInfo_Out)
	CliqueStoreDonate(*net.Session, *CliqueStoreDonate_Out)
	CliqueStoreInfo(*net.Session, *CliqueStoreInfo_Out)
	CliqueStoreSendChest(*net.Session, *CliqueStoreSendChest_Out)
}

func DecodeIn(message []byte) Request {
	var actionId = message[0]
	var buffer = net.NewBuffer(message[1:])
	switch actionId {
	case 1:
		request := new(CliqueBaseDonate_In)
		request.Decode(buffer)
		return request
	case 2:
		request := new(CliqueBuildingStatus_In)
		request.Decode(buffer)
		return request
	case 3:
		request := new(CliqueBankDonate_In)
		request.Decode(buffer)
		return request
	case 4:
		request := new(CliqueBankBuy_In)
		request.Decode(buffer)
		return request
	case 5:
		request := new(CliqueBankSold_In)
		request.Decode(buffer)
		return request
	case 6:
		request := new(CliqueKongfuDonate_In)
		request.Decode(buffer)
		return request
	case 7:
		request := new(CliqueKongfuInfo_In)
		request.Decode(buffer)
		return request
	case 8:
		request := new(CliqueKongfuTrain_In)
		request.Decode(buffer)
		return request
	case 9:
		request := new(CliqueTempleWorship_In)
		request.Decode(buffer)
		return request
	case 10:
		request := new(CliqueTempleDonate_In)
		request.Decode(buffer)
		return request
	case 11:
		request := new(CliqueTempleInfo_In)
		request.Decode(buffer)
		return request
	case 12:
		request := new(CliqueStoreDonate_In)
		request.Decode(buffer)
		return request
	case 13:
		request := new(CliqueStoreInfo_In)
		request.Decode(buffer)
		return request
	case 14:
		request := new(CliqueStoreSendChest_In)
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
		request := new(CliqueBaseDonate_Out)
		request.Decode(buffer)
		return request
	case 2:
		request := new(CliqueBuildingStatus_Out)
		request.Decode(buffer)
		return request
	case 3:
		request := new(CliqueBankDonate_Out)
		request.Decode(buffer)
		return request
	case 4:
		request := new(CliqueBankBuy_Out)
		request.Decode(buffer)
		return request
	case 5:
		request := new(CliqueBankSold_Out)
		request.Decode(buffer)
		return request
	case 6:
		request := new(CliqueKongfuDonate_Out)
		request.Decode(buffer)
		return request
	case 7:
		request := new(CliqueKongfuInfo_Out)
		request.Decode(buffer)
		return request
	case 8:
		request := new(CliqueKongfuTrain_Out)
		request.Decode(buffer)
		return request
	case 9:
		request := new(CliqueTempleWorship_Out)
		request.Decode(buffer)
		return request
	case 10:
		request := new(CliqueTempleDonate_Out)
		request.Decode(buffer)
		return request
	case 11:
		request := new(CliqueTempleInfo_Out)
		request.Decode(buffer)
		return request
	case 12:
		request := new(CliqueStoreDonate_Out)
		request.Decode(buffer)
		return request
	case 13:
		request := new(CliqueStoreInfo_Out)
		request.Decode(buffer)
		return request
	case 14:
		request := new(CliqueStoreSendChest_Out)
		request.Decode(buffer)
		return request
	}
	_ = buffer
	panic("unsupported response")
}

type CliqueKongfuTrainResult int8

const (
	CLIQUE_KONGFU_TRAIN_RESULT_SUCCESS      CliqueKongfuTrainResult = 0
	CLIQUE_KONGFU_TRAIN_RESULT_LACK_CONTRIB CliqueKongfuTrainResult = 1
	CLIQUE_KONGFU_TRAIN_RESULT_NO_CLIQUE    CliqueKongfuTrainResult = 2
	CLIQUE_KONGFU_TRAIN_RESULT_MAX_LEVEL    CliqueKongfuTrainResult = 3
)

type CliqueBankSoldResult int8

const (
	CLIQUE_BANK_SOLD_RESULT_SUCCESS   CliqueBankSoldResult = 0
	CLIQUE_BANK_SOLD_RESULT_CD        CliqueBankSoldResult = 1
	CLIQUE_BANK_SOLD_RESULT_NO_CLIQUE CliqueBankSoldResult = 2
	CLIQUE_BANK_SOLD_RESULT_MAX_LEVEL CliqueBankSoldResult = 3
)

type CliqueBuildingDonateResult int8

const (
	CLIQUE_BUILDING_DONATE_RESULT_SUCCESS CliqueBuildingDonateResult = 0
	CLIQUE_BUILDING_DONATE_RESULT_FAILED  CliqueBuildingDonateResult = 1
)

type CliqueStoreChest int8

const (
	CLIQUE_STORE_CHEST_JUNLIANG CliqueStoreChest = 1
	CLIQUE_STORE_CHEST_BAOXIANG CliqueStoreChest = 2
)

type CliqueStoreSendResult int8

const (
	CLIQUE_STORE_SEND_RESULT_SUCCESS                CliqueStoreSendResult = 0
	CLIQUE_STORE_SEND_RESULT_LACK_COINS             CliqueStoreSendResult = 1
	CLIQUE_STORE_SEND_RESULT_NO_CLIQUE              CliqueStoreSendResult = 2
	CLIQUE_STORE_SEND_RESULT_NOT_MANAGER            CliqueStoreSendResult = 3
	CLIQUE_STORE_SEND_RESULT_TIMES_NOT_ENOUGH       CliqueStoreSendResult = 4
	CLIQUE_STORE_SEND_RESULT_CLIQUE_NOT_FOUND       CliqueStoreSendResult = 5
	CLIQUE_STORE_SEND_RESULT_CLIQUE_CHEST_NOT_FOUND CliqueStoreSendResult = 6
)

type CliqueBuildingStatusBase struct {
	Level       int16 `json:"level"`
	DonateCoins int64 `json:"donate_coins"`
}

type CliqueBuildingStatusBank struct {
	Level                int16 `json:"level"`
	DonateCoins          int32 `json:"donate_coins"`
	SilverCouponNum      int16 `json:"silver_coupon_num"`
	SilverCouponTimespan int64 `json:"silver_coupon_timespan"`
	GoldCouponNum        int16 `json:"gold_coupon_num"`
	GoldCouponTimespan   int64 `json:"gold_coupon_timespan"`
}

type CliqueBuildingStatusAttack struct {
	Level       int16 `json:"level"`
	DonateCoins int32 `json:"donate_coins"`
}

type CliqueBuildingStatusHealth struct {
	Level       int16 `json:"level"`
	DonateCoins int32 `json:"donate_coins"`
}

type CliqueBuildingStatusDefence struct {
	Level       int16 `json:"level"`
	DonateCoins int32 `json:"donate_coins"`
}

type CliqueBuildingStatusTemple struct {
	Level       int16 `json:"level"`
	DonateCoins int32 `json:"donate_coins"`
}

type CliqueBuildingStatusStore struct {
	DonateCoins int32 `json:"donate_coins"`
}

type AncestralHallWorship int8

const (
	ANCESTRAL_HALL_WORSHIP_WHITESANDALWOOD AncestralHallWorship = 1
	ANCESTRAL_HALL_WORSHIP_STORAX          AncestralHallWorship = 2
	ANCESTRAL_HALL_WORSHIP_DAYS            AncestralHallWorship = 3
)

func (this *CliqueBuildingStatusBase) Decode(buffer *net.Buffer) {
	this.Level = int16(buffer.ReadUint16LE())
	this.DonateCoins = int64(buffer.ReadUint64LE())
}

func (this *CliqueBuildingStatusBase) Encode(buffer *net.Buffer) {
	buffer.WriteUint16LE(uint16(this.Level))
	buffer.WriteUint64LE(uint64(this.DonateCoins))
}

func (this *CliqueBuildingStatusBase) ByteSize() int {
	size := 10
	return size
}

func (this *CliqueBuildingStatusBank) Decode(buffer *net.Buffer) {
	this.Level = int16(buffer.ReadUint16LE())
	this.DonateCoins = int32(buffer.ReadUint32LE())
	this.SilverCouponNum = int16(buffer.ReadUint16LE())
	this.SilverCouponTimespan = int64(buffer.ReadUint64LE())
	this.GoldCouponNum = int16(buffer.ReadUint16LE())
	this.GoldCouponTimespan = int64(buffer.ReadUint64LE())
}

func (this *CliqueBuildingStatusBank) Encode(buffer *net.Buffer) {
	buffer.WriteUint16LE(uint16(this.Level))
	buffer.WriteUint32LE(uint32(this.DonateCoins))
	buffer.WriteUint16LE(uint16(this.SilverCouponNum))
	buffer.WriteUint64LE(uint64(this.SilverCouponTimespan))
	buffer.WriteUint16LE(uint16(this.GoldCouponNum))
	buffer.WriteUint64LE(uint64(this.GoldCouponTimespan))
}

func (this *CliqueBuildingStatusBank) ByteSize() int {
	size := 26
	return size
}

func (this *CliqueBuildingStatusAttack) Decode(buffer *net.Buffer) {
	this.Level = int16(buffer.ReadUint16LE())
	this.DonateCoins = int32(buffer.ReadUint32LE())
}

func (this *CliqueBuildingStatusAttack) Encode(buffer *net.Buffer) {
	buffer.WriteUint16LE(uint16(this.Level))
	buffer.WriteUint32LE(uint32(this.DonateCoins))
}

func (this *CliqueBuildingStatusAttack) ByteSize() int {
	size := 6
	return size
}

func (this *CliqueBuildingStatusHealth) Decode(buffer *net.Buffer) {
	this.Level = int16(buffer.ReadUint16LE())
	this.DonateCoins = int32(buffer.ReadUint32LE())
}

func (this *CliqueBuildingStatusHealth) Encode(buffer *net.Buffer) {
	buffer.WriteUint16LE(uint16(this.Level))
	buffer.WriteUint32LE(uint32(this.DonateCoins))
}

func (this *CliqueBuildingStatusHealth) ByteSize() int {
	size := 6
	return size
}

func (this *CliqueBuildingStatusDefence) Decode(buffer *net.Buffer) {
	this.Level = int16(buffer.ReadUint16LE())
	this.DonateCoins = int32(buffer.ReadUint32LE())
}

func (this *CliqueBuildingStatusDefence) Encode(buffer *net.Buffer) {
	buffer.WriteUint16LE(uint16(this.Level))
	buffer.WriteUint32LE(uint32(this.DonateCoins))
}

func (this *CliqueBuildingStatusDefence) ByteSize() int {
	size := 6
	return size
}

func (this *CliqueBuildingStatusTemple) Decode(buffer *net.Buffer) {
	this.Level = int16(buffer.ReadUint16LE())
	this.DonateCoins = int32(buffer.ReadUint32LE())
}

func (this *CliqueBuildingStatusTemple) Encode(buffer *net.Buffer) {
	buffer.WriteUint16LE(uint16(this.Level))
	buffer.WriteUint32LE(uint32(this.DonateCoins))
}

func (this *CliqueBuildingStatusTemple) ByteSize() int {
	size := 6
	return size
}

func (this *CliqueBuildingStatusStore) Decode(buffer *net.Buffer) {
	this.DonateCoins = int32(buffer.ReadUint32LE())
}

func (this *CliqueBuildingStatusStore) Encode(buffer *net.Buffer) {
	buffer.WriteUint32LE(uint32(this.DonateCoins))
}

func (this *CliqueBuildingStatusStore) ByteSize() int {
	size := 4
	return size
}

type CliqueBaseDonate_In struct {
	Money int64 `json:"money"`
}

func (this *CliqueBaseDonate_In) Process(session *net.Session) {
	g_InHandler.CliqueBaseDonate(session, this)
}

func (this *CliqueBaseDonate_In) TypeName() string {
	return "clique_building.clique_base_donate.in"
}

func (this *CliqueBaseDonate_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 34, 1
}

type CliqueBaseDonate_Out struct {
	Result                        CliqueBuildingDonateResult `json:"result"`
	CliqueBuildingBaseLevel       int16                      `json:"clique_building_base_level"`
	CliqueBuildingBaseDonateCoins int32                      `json:"clique_building_base_donate_coins"`
	PlayerDonateCoins             int64                      `json:"player_donate_coins"`
}

func (this *CliqueBaseDonate_Out) Process(session *net.Session) {
	g_OutHandler.CliqueBaseDonate(session, this)
}

func (this *CliqueBaseDonate_Out) TypeName() string {
	return "clique_building.clique_base_donate.out"
}

func (this *CliqueBaseDonate_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 34, 1
}

func (this *CliqueBaseDonate_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type CliqueBuildingStatus_In struct {
}

func (this *CliqueBuildingStatus_In) Process(session *net.Session) {
	g_InHandler.CliqueBuildingStatus(session, this)
}

func (this *CliqueBuildingStatus_In) TypeName() string {
	return "clique_building.clique_building_status.in"
}

func (this *CliqueBuildingStatus_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 34, 2
}

type CliqueBuildingStatus_Out struct {
	Success                bool                        `json:"success"`
	DailyTotalDonatedCoins int64                       `json:"daily_total_donated_coins"`
	Base                   CliqueBuildingStatusBase    `json:"base"`
	Bank                   CliqueBuildingStatusBank    `json:"bank"`
	AttackBuilding         CliqueBuildingStatusAttack  `json:"attack_building"`
	HealthBuilding         CliqueBuildingStatusHealth  `json:"health_building"`
	DefenceBuilding        CliqueBuildingStatusDefence `json:"defence_building"`
	TempleBuilding         CliqueBuildingStatusTemple  `json:"temple_building"`
	StoreBuilding          CliqueBuildingStatusStore   `json:"store_building"`
}

func (this *CliqueBuildingStatus_Out) Process(session *net.Session) {
	g_OutHandler.CliqueBuildingStatus(session, this)
}

func (this *CliqueBuildingStatus_Out) TypeName() string {
	return "clique_building.clique_building_status.out"
}

func (this *CliqueBuildingStatus_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 34, 2
}

func (this *CliqueBuildingStatus_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type CliqueBankDonate_In struct {
	Money int64 `json:"money"`
}

func (this *CliqueBankDonate_In) Process(session *net.Session) {
	g_InHandler.CliqueBankDonate(session, this)
}

func (this *CliqueBankDonate_In) TypeName() string {
	return "clique_building.clique_bank_donate.in"
}

func (this *CliqueBankDonate_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 34, 3
}

type CliqueBankDonate_Out struct {
	Result                        CliqueBuildingDonateResult `json:"result"`
	CliqueBuildingBankLevel       int16                      `json:"clique_building_bank_level"`
	CliqueBuildingBankDonateCoins int32                      `json:"clique_building_bank_donate_coins"`
	PlayerDonateCoins             int64                      `json:"player_donate_coins"`
}

func (this *CliqueBankDonate_Out) Process(session *net.Session) {
	g_OutHandler.CliqueBankDonate(session, this)
}

func (this *CliqueBankDonate_Out) TypeName() string {
	return "clique_building.clique_bank_donate.out"
}

func (this *CliqueBankDonate_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 34, 3
}

func (this *CliqueBankDonate_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type CliqueBankBuy_In struct {
	Kind int8 `json:"kind"`
	Num  int8 `json:"num"`
}

func (this *CliqueBankBuy_In) Process(session *net.Session) {
	g_InHandler.CliqueBankBuy(session, this)
}

func (this *CliqueBankBuy_In) TypeName() string {
	return "clique_building.clique_bank_buy.in"
}

func (this *CliqueBankBuy_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 34, 4
}

type CliqueBankBuy_Out struct {
	Success bool `json:"success"`
}

func (this *CliqueBankBuy_Out) Process(session *net.Session) {
	g_OutHandler.CliqueBankBuy(session, this)
}

func (this *CliqueBankBuy_Out) TypeName() string {
	return "clique_building.clique_bank_buy.out"
}

func (this *CliqueBankBuy_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 34, 4
}

func (this *CliqueBankBuy_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type CliqueBankSold_In struct {
	Kind int8 `json:"kind"`
}

func (this *CliqueBankSold_In) Process(session *net.Session) {
	g_InHandler.CliqueBankSold(session, this)
}

func (this *CliqueBankSold_In) TypeName() string {
	return "clique_building.clique_bank_sold.in"
}

func (this *CliqueBankSold_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 34, 5
}

type CliqueBankSold_Out struct {
	Result CliqueBankSoldResult `json:"result"`
}

func (this *CliqueBankSold_Out) Process(session *net.Session) {
	g_OutHandler.CliqueBankSold(session, this)
}

func (this *CliqueBankSold_Out) TypeName() string {
	return "clique_building.clique_bank_sold.out"
}

func (this *CliqueBankSold_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 34, 5
}

func (this *CliqueBankSold_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type CliqueKongfuDonate_In struct {
	Building int32 `json:"building"`
	Money    int64 `json:"money"`
}

func (this *CliqueKongfuDonate_In) Process(session *net.Session) {
	g_InHandler.CliqueKongfuDonate(session, this)
}

func (this *CliqueKongfuDonate_In) TypeName() string {
	return "clique_building.clique_kongfu_donate.in"
}

func (this *CliqueKongfuDonate_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 34, 6
}

type CliqueKongfuDonate_Out struct {
	Building          int32                      `json:"building"`
	Result            CliqueBuildingDonateResult `json:"result"`
	TotalDonateCoins  int32                      `json:"total_donate_coins"`
	PlayerDonateCoins int64                      `json:"player_donate_coins"`
	Level             int16                      `json:"level"`
}

func (this *CliqueKongfuDonate_Out) Process(session *net.Session) {
	g_OutHandler.CliqueKongfuDonate(session, this)
}

func (this *CliqueKongfuDonate_Out) TypeName() string {
	return "clique_building.clique_kongfu_donate.out"
}

func (this *CliqueKongfuDonate_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 34, 6
}

func (this *CliqueKongfuDonate_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type CliqueKongfuInfo_In struct {
	Building int32 `json:"building"`
}

func (this *CliqueKongfuInfo_In) Process(session *net.Session) {
	g_InHandler.CliqueKongfuInfo(session, this)
}

func (this *CliqueKongfuInfo_In) TypeName() string {
	return "clique_building.clique_kongfu_info.in"
}

func (this *CliqueKongfuInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 34, 7
}

type CliqueKongfuInfo_Out struct {
	KongfuList []CliqueKongfuInfo_Out_KongfuList `json:"kongfu_list"`
}

type CliqueKongfuInfo_Out_KongfuList struct {
	KongfuId int32 `json:"kongfu_id"`
	Level    int16 `json:"level"`
}

func (this *CliqueKongfuInfo_Out) Process(session *net.Session) {
	g_OutHandler.CliqueKongfuInfo(session, this)
}

func (this *CliqueKongfuInfo_Out) TypeName() string {
	return "clique_building.clique_kongfu_info.out"
}

func (this *CliqueKongfuInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 34, 7
}

func (this *CliqueKongfuInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type CliqueKongfuTrain_In struct {
	KongfuId int32 `json:"kongfu_id"`
}

func (this *CliqueKongfuTrain_In) Process(session *net.Session) {
	g_InHandler.CliqueKongfuTrain(session, this)
}

func (this *CliqueKongfuTrain_In) TypeName() string {
	return "clique_building.clique_kongfu_train.in"
}

func (this *CliqueKongfuTrain_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 34, 8
}

type CliqueKongfuTrain_Out struct {
	KongfuId int32                   `json:"kongfu_id"`
	Result   CliqueKongfuTrainResult `json:"result"`
	Level    int16                   `json:"level"`
}

func (this *CliqueKongfuTrain_Out) Process(session *net.Session) {
	g_OutHandler.CliqueKongfuTrain(session, this)
}

func (this *CliqueKongfuTrain_Out) TypeName() string {
	return "clique_building.clique_kongfu_train.out"
}

func (this *CliqueKongfuTrain_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 34, 8
}

func (this *CliqueKongfuTrain_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type CliqueTempleWorship_In struct {
	WorshipType AncestralHallWorship `json:"worship_type"`
}

func (this *CliqueTempleWorship_In) Process(session *net.Session) {
	g_InHandler.CliqueTempleWorship(session, this)
}

func (this *CliqueTempleWorship_In) TypeName() string {
	return "clique_building.clique_temple_worship.in"
}

func (this *CliqueTempleWorship_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 34, 9
}

type CliqueTempleWorship_Out struct {
	Success bool `json:"success"`
}

func (this *CliqueTempleWorship_Out) Process(session *net.Session) {
	g_OutHandler.CliqueTempleWorship(session, this)
}

func (this *CliqueTempleWorship_Out) TypeName() string {
	return "clique_building.clique_temple_worship.out"
}

func (this *CliqueTempleWorship_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 34, 9
}

func (this *CliqueTempleWorship_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type CliqueTempleDonate_In struct {
	Money int64 `json:"money"`
}

func (this *CliqueTempleDonate_In) Process(session *net.Session) {
	g_InHandler.CliqueTempleDonate(session, this)
}

func (this *CliqueTempleDonate_In) TypeName() string {
	return "clique_building.clique_temple_donate.in"
}

func (this *CliqueTempleDonate_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 34, 10
}

type CliqueTempleDonate_Out struct {
	Result              CliqueBuildingDonateResult `json:"result"`
	TempleBuildingLevel int16                      `json:"temple_building_level"`
	TempleBuildingCoins int32                      `json:"temple_building_coins"`
	Totaldonatecoins    int64                      `json:"totaldonatecoins"`
}

func (this *CliqueTempleDonate_Out) Process(session *net.Session) {
	g_OutHandler.CliqueTempleDonate(session, this)
}

func (this *CliqueTempleDonate_Out) TypeName() string {
	return "clique_building.clique_temple_donate.out"
}

func (this *CliqueTempleDonate_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 34, 10
}

func (this *CliqueTempleDonate_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type CliqueTempleInfo_In struct {
}

func (this *CliqueTempleInfo_In) Process(session *net.Session) {
	g_InHandler.CliqueTempleInfo(session, this)
}

func (this *CliqueTempleInfo_In) TypeName() string {
	return "clique_building.clique_temple_info.in"
}

func (this *CliqueTempleInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 34, 11
}

type CliqueTempleInfo_Out struct {
	Totaldonatecoins    int64 `json:"totaldonatecoins"`
	Isworship           bool  `json:"isworship"`
	WorshipType         int8  `json:"worship_type"`
	Worshipcnt          int8  `json:"worshipcnt"`
	TempleBuildingLevel int16 `json:"temple_building_level"`
	TempleBuildingCoins int32 `json:"temple_building_coins"`
}

func (this *CliqueTempleInfo_Out) Process(session *net.Session) {
	g_OutHandler.CliqueTempleInfo(session, this)
}

func (this *CliqueTempleInfo_Out) TypeName() string {
	return "clique_building.clique_temple_info.out"
}

func (this *CliqueTempleInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 34, 11
}

func (this *CliqueTempleInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type CliqueStoreDonate_In struct {
	Money int64 `json:"money"`
}

func (this *CliqueStoreDonate_In) Process(session *net.Session) {
	g_InHandler.CliqueStoreDonate(session, this)
}

func (this *CliqueStoreDonate_In) TypeName() string {
	return "clique_building.clique_store_donate.in"
}

func (this *CliqueStoreDonate_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 34, 12
}

type CliqueStoreDonate_Out struct {
	Result             CliqueBuildingDonateResult `json:"result"`
	StoreBuildingCoins int32                      `json:"store_building_coins"`
	PlayerDonateCoins  int64                      `json:"player_donate_coins"`
}

func (this *CliqueStoreDonate_Out) Process(session *net.Session) {
	g_OutHandler.CliqueStoreDonate(session, this)
}

func (this *CliqueStoreDonate_Out) TypeName() string {
	return "clique_building.clique_store_donate.out"
}

func (this *CliqueStoreDonate_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 34, 12
}

func (this *CliqueStoreDonate_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type CliqueStoreInfo_In struct {
}

func (this *CliqueStoreInfo_In) Process(session *net.Session) {
	g_InHandler.CliqueStoreInfo(session, this)
}

func (this *CliqueStoreInfo_In) TypeName() string {
	return "clique_building.clique_store_info.in"
}

func (this *CliqueStoreInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 34, 13
}

type CliqueStoreInfo_Out struct {
	OwnerPid           int64 `json:"owner_pid"`
	Manager1Pid        int64 `json:"manager1_pid"`
	Manager2Pid        int64 `json:"manager2_pid"`
	StoreBuildingCoins int32 `json:"store_building_coins"`
	SendTimes          int16 `json:"send_times"`
}

func (this *CliqueStoreInfo_Out) Process(session *net.Session) {
	g_OutHandler.CliqueStoreInfo(session, this)
}

func (this *CliqueStoreInfo_Out) TypeName() string {
	return "clique_building.clique_store_info.out"
}

func (this *CliqueStoreInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 34, 13
}

func (this *CliqueStoreInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type CliqueStoreSendChest_In struct {
	StoreChestType CliqueStoreChest `json:"store_chest_type"`
}

func (this *CliqueStoreSendChest_In) Process(session *net.Session) {
	g_InHandler.CliqueStoreSendChest(session, this)
}

func (this *CliqueStoreSendChest_In) TypeName() string {
	return "clique_building.clique_store_send_chest.in"
}

func (this *CliqueStoreSendChest_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 34, 14
}

type CliqueStoreSendChest_Out struct {
	Result CliqueStoreSendResult `json:"result"`
}

func (this *CliqueStoreSendChest_Out) Process(session *net.Session) {
	g_OutHandler.CliqueStoreSendChest(session, this)
}

func (this *CliqueStoreSendChest_Out) TypeName() string {
	return "clique_building.clique_store_send_chest.out"
}

func (this *CliqueStoreSendChest_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 34, 14
}

func (this *CliqueStoreSendChest_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

func (this *CliqueBaseDonate_In) Decode(buffer *net.Buffer) {
	this.Money = int64(buffer.ReadUint64LE())
}

func (this *CliqueBaseDonate_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(34)
	buffer.WriteUint8(1)
	buffer.WriteUint64LE(uint64(this.Money))
}

func (this *CliqueBaseDonate_In) ByteSize() int {
	size := 10
	return size
}

func (this *CliqueBaseDonate_Out) Decode(buffer *net.Buffer) {
	this.Result = CliqueBuildingDonateResult(buffer.ReadUint8())
	this.CliqueBuildingBaseLevel = int16(buffer.ReadUint16LE())
	this.CliqueBuildingBaseDonateCoins = int32(buffer.ReadUint32LE())
	this.PlayerDonateCoins = int64(buffer.ReadUint64LE())
}

func (this *CliqueBaseDonate_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(34)
	buffer.WriteUint8(1)
	buffer.WriteUint8(uint8(this.Result))
	buffer.WriteUint16LE(uint16(this.CliqueBuildingBaseLevel))
	buffer.WriteUint32LE(uint32(this.CliqueBuildingBaseDonateCoins))
	buffer.WriteUint64LE(uint64(this.PlayerDonateCoins))
}

func (this *CliqueBaseDonate_Out) ByteSize() int {
	size := 17
	return size
}

func (this *CliqueBuildingStatus_In) Decode(buffer *net.Buffer) {
}

func (this *CliqueBuildingStatus_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(34)
	buffer.WriteUint8(2)
}

func (this *CliqueBuildingStatus_In) ByteSize() int {
	size := 2
	return size
}

func (this *CliqueBuildingStatus_Out) Decode(buffer *net.Buffer) {
	this.Success = buffer.ReadUint8() == 1
	this.DailyTotalDonatedCoins = int64(buffer.ReadUint64LE())
	this.Base.Decode(buffer)
	this.Bank.Decode(buffer)
	this.AttackBuilding.Decode(buffer)
	this.HealthBuilding.Decode(buffer)
	this.DefenceBuilding.Decode(buffer)
	this.TempleBuilding.Decode(buffer)
	this.StoreBuilding.Decode(buffer)
}

func (this *CliqueBuildingStatus_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(34)
	buffer.WriteUint8(2)
	if this.Success {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
	buffer.WriteUint64LE(uint64(this.DailyTotalDonatedCoins))
	this.Base.Encode(buffer)
	this.Bank.Encode(buffer)
	this.AttackBuilding.Encode(buffer)
	this.HealthBuilding.Encode(buffer)
	this.DefenceBuilding.Encode(buffer)
	this.TempleBuilding.Encode(buffer)
	this.StoreBuilding.Encode(buffer)
}

func (this *CliqueBuildingStatus_Out) ByteSize() int {
	size := 11
	size += this.Base.ByteSize()
	size += this.Bank.ByteSize()
	size += this.AttackBuilding.ByteSize()
	size += this.HealthBuilding.ByteSize()
	size += this.DefenceBuilding.ByteSize()
	size += this.TempleBuilding.ByteSize()
	size += this.StoreBuilding.ByteSize()
	return size
}

func (this *CliqueBankDonate_In) Decode(buffer *net.Buffer) {
	this.Money = int64(buffer.ReadUint64LE())
}

func (this *CliqueBankDonate_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(34)
	buffer.WriteUint8(3)
	buffer.WriteUint64LE(uint64(this.Money))
}

func (this *CliqueBankDonate_In) ByteSize() int {
	size := 10
	return size
}

func (this *CliqueBankDonate_Out) Decode(buffer *net.Buffer) {
	this.Result = CliqueBuildingDonateResult(buffer.ReadUint8())
	this.CliqueBuildingBankLevel = int16(buffer.ReadUint16LE())
	this.CliqueBuildingBankDonateCoins = int32(buffer.ReadUint32LE())
	this.PlayerDonateCoins = int64(buffer.ReadUint64LE())
}

func (this *CliqueBankDonate_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(34)
	buffer.WriteUint8(3)
	buffer.WriteUint8(uint8(this.Result))
	buffer.WriteUint16LE(uint16(this.CliqueBuildingBankLevel))
	buffer.WriteUint32LE(uint32(this.CliqueBuildingBankDonateCoins))
	buffer.WriteUint64LE(uint64(this.PlayerDonateCoins))
}

func (this *CliqueBankDonate_Out) ByteSize() int {
	size := 17
	return size
}

func (this *CliqueBankBuy_In) Decode(buffer *net.Buffer) {
	this.Kind = int8(buffer.ReadUint8())
	this.Num = int8(buffer.ReadUint8())
}

func (this *CliqueBankBuy_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(34)
	buffer.WriteUint8(4)
	buffer.WriteUint8(uint8(this.Kind))
	buffer.WriteUint8(uint8(this.Num))
}

func (this *CliqueBankBuy_In) ByteSize() int {
	size := 4
	return size
}

func (this *CliqueBankBuy_Out) Decode(buffer *net.Buffer) {
	this.Success = buffer.ReadUint8() == 1
}

func (this *CliqueBankBuy_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(34)
	buffer.WriteUint8(4)
	if this.Success {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *CliqueBankBuy_Out) ByteSize() int {
	size := 3
	return size
}

func (this *CliqueBankSold_In) Decode(buffer *net.Buffer) {
	this.Kind = int8(buffer.ReadUint8())
}

func (this *CliqueBankSold_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(34)
	buffer.WriteUint8(5)
	buffer.WriteUint8(uint8(this.Kind))
}

func (this *CliqueBankSold_In) ByteSize() int {
	size := 3
	return size
}

func (this *CliqueBankSold_Out) Decode(buffer *net.Buffer) {
	this.Result = CliqueBankSoldResult(buffer.ReadUint8())
}

func (this *CliqueBankSold_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(34)
	buffer.WriteUint8(5)
	buffer.WriteUint8(uint8(this.Result))
}

func (this *CliqueBankSold_Out) ByteSize() int {
	size := 3
	return size
}

func (this *CliqueKongfuDonate_In) Decode(buffer *net.Buffer) {
	this.Building = int32(buffer.ReadUint32LE())
	this.Money = int64(buffer.ReadUint64LE())
}

func (this *CliqueKongfuDonate_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(34)
	buffer.WriteUint8(6)
	buffer.WriteUint32LE(uint32(this.Building))
	buffer.WriteUint64LE(uint64(this.Money))
}

func (this *CliqueKongfuDonate_In) ByteSize() int {
	size := 14
	return size
}

func (this *CliqueKongfuDonate_Out) Decode(buffer *net.Buffer) {
	this.Building = int32(buffer.ReadUint32LE())
	this.Result = CliqueBuildingDonateResult(buffer.ReadUint8())
	this.TotalDonateCoins = int32(buffer.ReadUint32LE())
	this.PlayerDonateCoins = int64(buffer.ReadUint64LE())
	this.Level = int16(buffer.ReadUint16LE())
}

func (this *CliqueKongfuDonate_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(34)
	buffer.WriteUint8(6)
	buffer.WriteUint32LE(uint32(this.Building))
	buffer.WriteUint8(uint8(this.Result))
	buffer.WriteUint32LE(uint32(this.TotalDonateCoins))
	buffer.WriteUint64LE(uint64(this.PlayerDonateCoins))
	buffer.WriteUint16LE(uint16(this.Level))
}

func (this *CliqueKongfuDonate_Out) ByteSize() int {
	size := 21
	return size
}

func (this *CliqueKongfuInfo_In) Decode(buffer *net.Buffer) {
	this.Building = int32(buffer.ReadUint32LE())
}

func (this *CliqueKongfuInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(34)
	buffer.WriteUint8(7)
	buffer.WriteUint32LE(uint32(this.Building))
}

func (this *CliqueKongfuInfo_In) ByteSize() int {
	size := 6
	return size
}

func (this *CliqueKongfuInfo_Out) Decode(buffer *net.Buffer) {
	this.KongfuList = make([]CliqueKongfuInfo_Out_KongfuList, buffer.ReadUint8())
	for i := 0; i < len(this.KongfuList); i++ {
		this.KongfuList[i].Decode(buffer)
	}
}

func (this *CliqueKongfuInfo_Out_KongfuList) Decode(buffer *net.Buffer) {
	this.KongfuId = int32(buffer.ReadUint32LE())
	this.Level = int16(buffer.ReadUint16LE())
}

func (this *CliqueKongfuInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(34)
	buffer.WriteUint8(7)
	buffer.WriteUint8(uint8(len(this.KongfuList)))
	for i := 0; i < len(this.KongfuList); i++ {
		this.KongfuList[i].Encode(buffer)
	}
}

func (this *CliqueKongfuInfo_Out_KongfuList) Encode(buffer *net.Buffer) {
	buffer.WriteUint32LE(uint32(this.KongfuId))
	buffer.WriteUint16LE(uint16(this.Level))
}

func (this *CliqueKongfuInfo_Out) ByteSize() int {
	size := 3
	size += len(this.KongfuList) * 6
	return size
}

func (this *CliqueKongfuTrain_In) Decode(buffer *net.Buffer) {
	this.KongfuId = int32(buffer.ReadUint32LE())
}

func (this *CliqueKongfuTrain_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(34)
	buffer.WriteUint8(8)
	buffer.WriteUint32LE(uint32(this.KongfuId))
}

func (this *CliqueKongfuTrain_In) ByteSize() int {
	size := 6
	return size
}

func (this *CliqueKongfuTrain_Out) Decode(buffer *net.Buffer) {
	this.KongfuId = int32(buffer.ReadUint32LE())
	this.Result = CliqueKongfuTrainResult(buffer.ReadUint8())
	this.Level = int16(buffer.ReadUint16LE())
}

func (this *CliqueKongfuTrain_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(34)
	buffer.WriteUint8(8)
	buffer.WriteUint32LE(uint32(this.KongfuId))
	buffer.WriteUint8(uint8(this.Result))
	buffer.WriteUint16LE(uint16(this.Level))
}

func (this *CliqueKongfuTrain_Out) ByteSize() int {
	size := 9
	return size
}

func (this *CliqueTempleWorship_In) Decode(buffer *net.Buffer) {
	this.WorshipType = AncestralHallWorship(buffer.ReadUint8())
}

func (this *CliqueTempleWorship_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(34)
	buffer.WriteUint8(9)
	buffer.WriteUint8(uint8(this.WorshipType))
}

func (this *CliqueTempleWorship_In) ByteSize() int {
	size := 3
	return size
}

func (this *CliqueTempleWorship_Out) Decode(buffer *net.Buffer) {
	this.Success = buffer.ReadUint8() == 1
}

func (this *CliqueTempleWorship_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(34)
	buffer.WriteUint8(9)
	if this.Success {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *CliqueTempleWorship_Out) ByteSize() int {
	size := 3
	return size
}

func (this *CliqueTempleDonate_In) Decode(buffer *net.Buffer) {
	this.Money = int64(buffer.ReadUint64LE())
}

func (this *CliqueTempleDonate_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(34)
	buffer.WriteUint8(10)
	buffer.WriteUint64LE(uint64(this.Money))
}

func (this *CliqueTempleDonate_In) ByteSize() int {
	size := 10
	return size
}

func (this *CliqueTempleDonate_Out) Decode(buffer *net.Buffer) {
	this.Result = CliqueBuildingDonateResult(buffer.ReadUint8())
	this.TempleBuildingLevel = int16(buffer.ReadUint16LE())
	this.TempleBuildingCoins = int32(buffer.ReadUint32LE())
	this.Totaldonatecoins = int64(buffer.ReadUint64LE())
}

func (this *CliqueTempleDonate_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(34)
	buffer.WriteUint8(10)
	buffer.WriteUint8(uint8(this.Result))
	buffer.WriteUint16LE(uint16(this.TempleBuildingLevel))
	buffer.WriteUint32LE(uint32(this.TempleBuildingCoins))
	buffer.WriteUint64LE(uint64(this.Totaldonatecoins))
}

func (this *CliqueTempleDonate_Out) ByteSize() int {
	size := 17
	return size
}

func (this *CliqueTempleInfo_In) Decode(buffer *net.Buffer) {
}

func (this *CliqueTempleInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(34)
	buffer.WriteUint8(11)
}

func (this *CliqueTempleInfo_In) ByteSize() int {
	size := 2
	return size
}

func (this *CliqueTempleInfo_Out) Decode(buffer *net.Buffer) {
	this.Totaldonatecoins = int64(buffer.ReadUint64LE())
	this.Isworship = buffer.ReadUint8() == 1
	this.WorshipType = int8(buffer.ReadUint8())
	this.Worshipcnt = int8(buffer.ReadUint8())
	this.TempleBuildingLevel = int16(buffer.ReadUint16LE())
	this.TempleBuildingCoins = int32(buffer.ReadUint32LE())
}

func (this *CliqueTempleInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(34)
	buffer.WriteUint8(11)
	buffer.WriteUint64LE(uint64(this.Totaldonatecoins))
	if this.Isworship {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
	buffer.WriteUint8(uint8(this.WorshipType))
	buffer.WriteUint8(uint8(this.Worshipcnt))
	buffer.WriteUint16LE(uint16(this.TempleBuildingLevel))
	buffer.WriteUint32LE(uint32(this.TempleBuildingCoins))
}

func (this *CliqueTempleInfo_Out) ByteSize() int {
	size := 19
	return size
}

func (this *CliqueStoreDonate_In) Decode(buffer *net.Buffer) {
	this.Money = int64(buffer.ReadUint64LE())
}

func (this *CliqueStoreDonate_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(34)
	buffer.WriteUint8(12)
	buffer.WriteUint64LE(uint64(this.Money))
}

func (this *CliqueStoreDonate_In) ByteSize() int {
	size := 10
	return size
}

func (this *CliqueStoreDonate_Out) Decode(buffer *net.Buffer) {
	this.Result = CliqueBuildingDonateResult(buffer.ReadUint8())
	this.StoreBuildingCoins = int32(buffer.ReadUint32LE())
	this.PlayerDonateCoins = int64(buffer.ReadUint64LE())
}

func (this *CliqueStoreDonate_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(34)
	buffer.WriteUint8(12)
	buffer.WriteUint8(uint8(this.Result))
	buffer.WriteUint32LE(uint32(this.StoreBuildingCoins))
	buffer.WriteUint64LE(uint64(this.PlayerDonateCoins))
}

func (this *CliqueStoreDonate_Out) ByteSize() int {
	size := 15
	return size
}

func (this *CliqueStoreInfo_In) Decode(buffer *net.Buffer) {
}

func (this *CliqueStoreInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(34)
	buffer.WriteUint8(13)
}

func (this *CliqueStoreInfo_In) ByteSize() int {
	size := 2
	return size
}

func (this *CliqueStoreInfo_Out) Decode(buffer *net.Buffer) {
	this.OwnerPid = int64(buffer.ReadUint64LE())
	this.Manager1Pid = int64(buffer.ReadUint64LE())
	this.Manager2Pid = int64(buffer.ReadUint64LE())
	this.StoreBuildingCoins = int32(buffer.ReadUint32LE())
	this.SendTimes = int16(buffer.ReadUint16LE())
}

func (this *CliqueStoreInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(34)
	buffer.WriteUint8(13)
	buffer.WriteUint64LE(uint64(this.OwnerPid))
	buffer.WriteUint64LE(uint64(this.Manager1Pid))
	buffer.WriteUint64LE(uint64(this.Manager2Pid))
	buffer.WriteUint32LE(uint32(this.StoreBuildingCoins))
	buffer.WriteUint16LE(uint16(this.SendTimes))
}

func (this *CliqueStoreInfo_Out) ByteSize() int {
	size := 32
	return size
}

func (this *CliqueStoreSendChest_In) Decode(buffer *net.Buffer) {
	this.StoreChestType = CliqueStoreChest(buffer.ReadUint8())
}

func (this *CliqueStoreSendChest_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(34)
	buffer.WriteUint8(14)
	buffer.WriteUint8(uint8(this.StoreChestType))
}

func (this *CliqueStoreSendChest_In) ByteSize() int {
	size := 3
	return size
}

func (this *CliqueStoreSendChest_Out) Decode(buffer *net.Buffer) {
	this.Result = CliqueStoreSendResult(buffer.ReadUint8())
}

func (this *CliqueStoreSendChest_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(34)
	buffer.WriteUint8(14)
	buffer.WriteUint8(uint8(this.Result))
}

func (this *CliqueStoreSendChest_Out) ByteSize() int {
	size := 3
	return size
}
