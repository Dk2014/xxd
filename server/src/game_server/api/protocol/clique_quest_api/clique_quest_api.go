package clique_quest_api

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
	GetCliqueDailyQuest(*net.Session, *GetCliqueDailyQuest_In)
	AwardCliqueDailyQuest(*net.Session, *AwardCliqueDailyQuest_In)
	GetCliqueBuildingQuest(*net.Session, *GetCliqueBuildingQuest_In)
	AwardCliqueBuildingQuest(*net.Session, *AwardCliqueBuildingQuest_In)
}

type OutHandler interface {
	GetCliqueDailyQuest(*net.Session, *GetCliqueDailyQuest_Out)
	AwardCliqueDailyQuest(*net.Session, *AwardCliqueDailyQuest_Out)
	NotifyCliqueDailyChange(*net.Session, *NotifyCliqueDailyChange_Out)
	GetCliqueBuildingQuest(*net.Session, *GetCliqueBuildingQuest_Out)
	AwardCliqueBuildingQuest(*net.Session, *AwardCliqueBuildingQuest_Out)
}

func DecodeIn(message []byte) Request {
	var actionId = message[0]
	var buffer = net.NewBuffer(message[1:])
	switch actionId {
	case 1:
		request := new(GetCliqueDailyQuest_In)
		request.Decode(buffer)
		return request
	case 2:
		request := new(AwardCliqueDailyQuest_In)
		request.Decode(buffer)
		return request
	case 4:
		request := new(GetCliqueBuildingQuest_In)
		request.Decode(buffer)
		return request
	case 5:
		request := new(AwardCliqueBuildingQuest_In)
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
		request := new(GetCliqueDailyQuest_Out)
		request.Decode(buffer)
		return request
	case 2:
		request := new(AwardCliqueDailyQuest_Out)
		request.Decode(buffer)
		return request
	case 3:
		request := new(NotifyCliqueDailyChange_Out)
		request.Decode(buffer)
		return request
	case 4:
		request := new(GetCliqueBuildingQuest_Out)
		request.Decode(buffer)
		return request
	case 5:
		request := new(AwardCliqueBuildingQuest_Out)
		request.Decode(buffer)
		return request
	}
	_ = buffer
	panic("unsupported response")
}

type GetCliqueDailyQuest_In struct {
}

func (this *GetCliqueDailyQuest_In) Process(session *net.Session) {
	g_InHandler.GetCliqueDailyQuest(session, this)
}

func (this *GetCliqueDailyQuest_In) TypeName() string {
	return "clique_quest.get_clique_daily_Quest.in"
}

func (this *GetCliqueDailyQuest_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 35, 1
}

type GetCliqueDailyQuest_Out struct {
	Quest []GetCliqueDailyQuest_Out_Quest `json:"quest"`
}

type GetCliqueDailyQuest_Out_Quest struct {
	Id          int16 `json:"id"`
	FinishCount int16 `json:"finish_count"`
	AwardState  int8  `json:"award_state"`
}

func (this *GetCliqueDailyQuest_Out) Process(session *net.Session) {
	g_OutHandler.GetCliqueDailyQuest(session, this)
}

func (this *GetCliqueDailyQuest_Out) TypeName() string {
	return "clique_quest.get_clique_daily_Quest.out"
}

func (this *GetCliqueDailyQuest_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 35, 1
}

func (this *GetCliqueDailyQuest_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type AwardCliqueDailyQuest_In struct {
	Id int16 `json:"id"`
}

func (this *AwardCliqueDailyQuest_In) Process(session *net.Session) {
	g_InHandler.AwardCliqueDailyQuest(session, this)
}

func (this *AwardCliqueDailyQuest_In) TypeName() string {
	return "clique_quest.award_clique_daily_Quest.in"
}

func (this *AwardCliqueDailyQuest_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 35, 2
}

type AwardCliqueDailyQuest_Out struct {
	Result int8 `json:"result"`
}

func (this *AwardCliqueDailyQuest_Out) Process(session *net.Session) {
	g_OutHandler.AwardCliqueDailyQuest(session, this)
}

func (this *AwardCliqueDailyQuest_Out) TypeName() string {
	return "clique_quest.award_clique_daily_Quest.out"
}

func (this *AwardCliqueDailyQuest_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 35, 2
}

func (this *AwardCliqueDailyQuest_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type NotifyCliqueDailyChange_Out struct {
	Id          int16 `json:"id"`
	FinishCount int16 `json:"finish_count"`
	AwardState  int8  `json:"award_state"`
}

func (this *NotifyCliqueDailyChange_Out) Process(session *net.Session) {
	g_OutHandler.NotifyCliqueDailyChange(session, this)
}

func (this *NotifyCliqueDailyChange_Out) TypeName() string {
	return "clique_quest.notify_clique_daily_change.out"
}

func (this *NotifyCliqueDailyChange_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 35, 3
}

func (this *NotifyCliqueDailyChange_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetCliqueBuildingQuest_In struct {
}

func (this *GetCliqueBuildingQuest_In) Process(session *net.Session) {
	g_InHandler.GetCliqueBuildingQuest(session, this)
}

func (this *GetCliqueBuildingQuest_In) TypeName() string {
	return "clique_quest.get_clique_building_quest.in"
}

func (this *GetCliqueBuildingQuest_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 35, 4
}

type GetCliqueBuildingQuest_Out struct {
	Quest []GetCliqueBuildingQuest_Out_Quest `json:"quest"`
}

type GetCliqueBuildingQuest_Out_Quest struct {
	Id          int16 `json:"id"`
	AwardState  int8  `json:"award_state"`
	DonateCoins int64 `json:"donateCoins"`
}

func (this *GetCliqueBuildingQuest_Out) Process(session *net.Session) {
	g_OutHandler.GetCliqueBuildingQuest(session, this)
}

func (this *GetCliqueBuildingQuest_Out) TypeName() string {
	return "clique_quest.get_clique_building_quest.out"
}

func (this *GetCliqueBuildingQuest_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 35, 4
}

func (this *GetCliqueBuildingQuest_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type AwardCliqueBuildingQuest_In struct {
	Id int16 `json:"id"`
}

func (this *AwardCliqueBuildingQuest_In) Process(session *net.Session) {
	g_InHandler.AwardCliqueBuildingQuest(session, this)
}

func (this *AwardCliqueBuildingQuest_In) TypeName() string {
	return "clique_quest.award_clique_building_Quest.in"
}

func (this *AwardCliqueBuildingQuest_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 35, 5
}

type AwardCliqueBuildingQuest_Out struct {
	Result int8 `json:"result"`
}

func (this *AwardCliqueBuildingQuest_Out) Process(session *net.Session) {
	g_OutHandler.AwardCliqueBuildingQuest(session, this)
}

func (this *AwardCliqueBuildingQuest_Out) TypeName() string {
	return "clique_quest.award_clique_building_Quest.out"
}

func (this *AwardCliqueBuildingQuest_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 35, 5
}

func (this *AwardCliqueBuildingQuest_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

func (this *GetCliqueDailyQuest_In) Decode(buffer *net.Buffer) {
}

func (this *GetCliqueDailyQuest_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(35)
	buffer.WriteUint8(1)
}

func (this *GetCliqueDailyQuest_In) ByteSize() int {
	size := 2
	return size
}

func (this *GetCliqueDailyQuest_Out) Decode(buffer *net.Buffer) {
	this.Quest = make([]GetCliqueDailyQuest_Out_Quest, buffer.ReadUint8())
	for i := 0; i < len(this.Quest); i++ {
		this.Quest[i].Decode(buffer)
	}
}

func (this *GetCliqueDailyQuest_Out_Quest) Decode(buffer *net.Buffer) {
	this.Id = int16(buffer.ReadUint16LE())
	this.FinishCount = int16(buffer.ReadUint16LE())
	this.AwardState = int8(buffer.ReadUint8())
}

func (this *GetCliqueDailyQuest_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(35)
	buffer.WriteUint8(1)
	buffer.WriteUint8(uint8(len(this.Quest)))
	for i := 0; i < len(this.Quest); i++ {
		this.Quest[i].Encode(buffer)
	}
}

func (this *GetCliqueDailyQuest_Out_Quest) Encode(buffer *net.Buffer) {
	buffer.WriteUint16LE(uint16(this.Id))
	buffer.WriteUint16LE(uint16(this.FinishCount))
	buffer.WriteUint8(uint8(this.AwardState))
}

func (this *GetCliqueDailyQuest_Out) ByteSize() int {
	size := 3
	size += len(this.Quest) * 5
	return size
}

func (this *AwardCliqueDailyQuest_In) Decode(buffer *net.Buffer) {
	this.Id = int16(buffer.ReadUint16LE())
}

func (this *AwardCliqueDailyQuest_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(35)
	buffer.WriteUint8(2)
	buffer.WriteUint16LE(uint16(this.Id))
}

func (this *AwardCliqueDailyQuest_In) ByteSize() int {
	size := 4
	return size
}

func (this *AwardCliqueDailyQuest_Out) Decode(buffer *net.Buffer) {
	this.Result = int8(buffer.ReadUint8())
}

func (this *AwardCliqueDailyQuest_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(35)
	buffer.WriteUint8(2)
	buffer.WriteUint8(uint8(this.Result))
}

func (this *AwardCliqueDailyQuest_Out) ByteSize() int {
	size := 3
	return size
}

func (this *NotifyCliqueDailyChange_Out) Decode(buffer *net.Buffer) {
	this.Id = int16(buffer.ReadUint16LE())
	this.FinishCount = int16(buffer.ReadUint16LE())
	this.AwardState = int8(buffer.ReadUint8())
}

func (this *NotifyCliqueDailyChange_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(35)
	buffer.WriteUint8(3)
	buffer.WriteUint16LE(uint16(this.Id))
	buffer.WriteUint16LE(uint16(this.FinishCount))
	buffer.WriteUint8(uint8(this.AwardState))
}

func (this *NotifyCliqueDailyChange_Out) ByteSize() int {
	size := 7
	return size
}

func (this *GetCliqueBuildingQuest_In) Decode(buffer *net.Buffer) {
}

func (this *GetCliqueBuildingQuest_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(35)
	buffer.WriteUint8(4)
}

func (this *GetCliqueBuildingQuest_In) ByteSize() int {
	size := 2
	return size
}

func (this *GetCliqueBuildingQuest_Out) Decode(buffer *net.Buffer) {
	this.Quest = make([]GetCliqueBuildingQuest_Out_Quest, buffer.ReadUint8())
	for i := 0; i < len(this.Quest); i++ {
		this.Quest[i].Decode(buffer)
	}
}

func (this *GetCliqueBuildingQuest_Out_Quest) Decode(buffer *net.Buffer) {
	this.Id = int16(buffer.ReadUint16LE())
	this.AwardState = int8(buffer.ReadUint8())
	this.DonateCoins = int64(buffer.ReadUint64LE())
}

func (this *GetCliqueBuildingQuest_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(35)
	buffer.WriteUint8(4)
	buffer.WriteUint8(uint8(len(this.Quest)))
	for i := 0; i < len(this.Quest); i++ {
		this.Quest[i].Encode(buffer)
	}
}

func (this *GetCliqueBuildingQuest_Out_Quest) Encode(buffer *net.Buffer) {
	buffer.WriteUint16LE(uint16(this.Id))
	buffer.WriteUint8(uint8(this.AwardState))
	buffer.WriteUint64LE(uint64(this.DonateCoins))
}

func (this *GetCliqueBuildingQuest_Out) ByteSize() int {
	size := 3
	size += len(this.Quest) * 11
	return size
}

func (this *AwardCliqueBuildingQuest_In) Decode(buffer *net.Buffer) {
	this.Id = int16(buffer.ReadUint16LE())
}

func (this *AwardCliqueBuildingQuest_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(35)
	buffer.WriteUint8(5)
	buffer.WriteUint16LE(uint16(this.Id))
}

func (this *AwardCliqueBuildingQuest_In) ByteSize() int {
	size := 4
	return size
}

func (this *AwardCliqueBuildingQuest_Out) Decode(buffer *net.Buffer) {
	this.Result = int8(buffer.ReadUint8())
}

func (this *AwardCliqueBuildingQuest_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(35)
	buffer.WriteUint8(5)
	buffer.WriteUint8(uint8(this.Result))
}

func (this *AwardCliqueBuildingQuest_Out) ByteSize() int {
	size := 3
	return size
}
