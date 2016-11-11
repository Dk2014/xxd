package event_api

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
	LoginAwardInfo(*net.Session, *LoginAwardInfo_In)
	TakeLoginAward(*net.Session, *TakeLoginAward_In)
	GetEvents(*net.Session, *GetEvents_In)
	GetEventAward(*net.Session, *GetEventAward_In)
	PlayerEventPhysicalCost(*net.Session, *PlayerEventPhysicalCost_In)
	PlayerMonthCardInfo(*net.Session, *PlayerMonthCardInfo_In)
	GetSevenInfo(*net.Session, *GetSevenInfo_In)
	GetSevenAward(*net.Session, *GetSevenAward_In)
	GetRichmanClubInfo(*net.Session, *GetRichmanClubInfo_In)
	GetRichmanClubAward(*net.Session, *GetRichmanClubAward_In)
	InfoShare(*net.Session, *InfoShare_In)
	InfoGroupBuy(*net.Session, *InfoGroupBuy_In)
	GetIngotChangeTotal(*net.Session, *GetIngotChangeTotal_In)
	GetEventTotalAward(*net.Session, *GetEventTotalAward_In)
	GetEventArenaRank(*net.Session, *GetEventArenaRank_In)
	GetEventTenDrawTimes(*net.Session, *GetEventTenDrawTimes_In)
	GetEventRechargeAward(*net.Session, *GetEventRechargeAward_In)
	GetEventNewYear(*net.Session, *GetEventNewYear_In)
	QqVipContinue(*net.Session, *QqVipContinue_In)
	DailyOnlineInfo(*net.Session, *DailyOnlineInfo_In)
	TakeDailyOnlineAward(*net.Session, *TakeDailyOnlineAward_In)
}

type OutHandler interface {
	LoginAwardInfo(*net.Session, *LoginAwardInfo_Out)
	TakeLoginAward(*net.Session, *TakeLoginAward_Out)
	GetEvents(*net.Session, *GetEvents_Out)
	GetEventAward(*net.Session, *GetEventAward_Out)
	PlayerEventPhysicalCost(*net.Session, *PlayerEventPhysicalCost_Out)
	PlayerMonthCardInfo(*net.Session, *PlayerMonthCardInfo_Out)
	GetSevenInfo(*net.Session, *GetSevenInfo_Out)
	GetSevenAward(*net.Session, *GetSevenAward_Out)
	GetRichmanClubInfo(*net.Session, *GetRichmanClubInfo_Out)
	GetRichmanClubAward(*net.Session, *GetRichmanClubAward_Out)
	InfoShare(*net.Session, *InfoShare_Out)
	InfoGroupBuy(*net.Session, *InfoGroupBuy_Out)
	GetIngotChangeTotal(*net.Session, *GetIngotChangeTotal_Out)
	GetEventTotalAward(*net.Session, *GetEventTotalAward_Out)
	GetEventArenaRank(*net.Session, *GetEventArenaRank_Out)
	GetEventTenDrawTimes(*net.Session, *GetEventTenDrawTimes_Out)
	GetEventRechargeAward(*net.Session, *GetEventRechargeAward_Out)
	GetEventNewYear(*net.Session, *GetEventNewYear_Out)
	QqVipContinue(*net.Session, *QqVipContinue_Out)
	DailyOnlineInfo(*net.Session, *DailyOnlineInfo_Out)
	TakeDailyOnlineAward(*net.Session, *TakeDailyOnlineAward_Out)
}

func DecodeIn(message []byte) Request {
	var actionId = message[0]
	var buffer = net.NewBuffer(message[1:])
	switch actionId {
	case 1:
		request := new(LoginAwardInfo_In)
		request.Decode(buffer)
		return request
	case 2:
		request := new(TakeLoginAward_In)
		request.Decode(buffer)
		return request
	case 3:
		request := new(GetEvents_In)
		request.Decode(buffer)
		return request
	case 4:
		request := new(GetEventAward_In)
		request.Decode(buffer)
		return request
	case 6:
		request := new(PlayerEventPhysicalCost_In)
		request.Decode(buffer)
		return request
	case 7:
		request := new(PlayerMonthCardInfo_In)
		request.Decode(buffer)
		return request
	case 8:
		request := new(GetSevenInfo_In)
		request.Decode(buffer)
		return request
	case 9:
		request := new(GetSevenAward_In)
		request.Decode(buffer)
		return request
	case 10:
		request := new(GetRichmanClubInfo_In)
		request.Decode(buffer)
		return request
	case 11:
		request := new(GetRichmanClubAward_In)
		request.Decode(buffer)
		return request
	case 12:
		request := new(InfoShare_In)
		request.Decode(buffer)
		return request
	case 13:
		request := new(InfoGroupBuy_In)
		request.Decode(buffer)
		return request
	case 14:
		request := new(GetIngotChangeTotal_In)
		request.Decode(buffer)
		return request
	case 15:
		request := new(GetEventTotalAward_In)
		request.Decode(buffer)
		return request
	case 16:
		request := new(GetEventArenaRank_In)
		request.Decode(buffer)
		return request
	case 17:
		request := new(GetEventTenDrawTimes_In)
		request.Decode(buffer)
		return request
	case 18:
		request := new(GetEventRechargeAward_In)
		request.Decode(buffer)
		return request
	case 19:
		request := new(GetEventNewYear_In)
		request.Decode(buffer)
		return request
	case 20:
		request := new(QqVipContinue_In)
		request.Decode(buffer)
		return request
	case 21:
		request := new(DailyOnlineInfo_In)
		request.Decode(buffer)
		return request
	case 22:
		request := new(TakeDailyOnlineAward_In)
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
		request := new(LoginAwardInfo_Out)
		request.Decode(buffer)
		return request
	case 2:
		request := new(TakeLoginAward_Out)
		request.Decode(buffer)
		return request
	case 3:
		request := new(GetEvents_Out)
		request.Decode(buffer)
		return request
	case 4:
		request := new(GetEventAward_Out)
		request.Decode(buffer)
		return request
	case 6:
		request := new(PlayerEventPhysicalCost_Out)
		request.Decode(buffer)
		return request
	case 7:
		request := new(PlayerMonthCardInfo_Out)
		request.Decode(buffer)
		return request
	case 8:
		request := new(GetSevenInfo_Out)
		request.Decode(buffer)
		return request
	case 9:
		request := new(GetSevenAward_Out)
		request.Decode(buffer)
		return request
	case 10:
		request := new(GetRichmanClubInfo_Out)
		request.Decode(buffer)
		return request
	case 11:
		request := new(GetRichmanClubAward_Out)
		request.Decode(buffer)
		return request
	case 12:
		request := new(InfoShare_Out)
		request.Decode(buffer)
		return request
	case 13:
		request := new(InfoGroupBuy_Out)
		request.Decode(buffer)
		return request
	case 14:
		request := new(GetIngotChangeTotal_Out)
		request.Decode(buffer)
		return request
	case 15:
		request := new(GetEventTotalAward_Out)
		request.Decode(buffer)
		return request
	case 16:
		request := new(GetEventArenaRank_Out)
		request.Decode(buffer)
		return request
	case 17:
		request := new(GetEventTenDrawTimes_Out)
		request.Decode(buffer)
		return request
	case 18:
		request := new(GetEventRechargeAward_Out)
		request.Decode(buffer)
		return request
	case 19:
		request := new(GetEventNewYear_Out)
		request.Decode(buffer)
		return request
	case 20:
		request := new(QqVipContinue_Out)
		request.Decode(buffer)
		return request
	case 21:
		request := new(DailyOnlineInfo_Out)
		request.Decode(buffer)
		return request
	case 22:
		request := new(TakeDailyOnlineAward_Out)
		request.Decode(buffer)
		return request
	}
	_ = buffer
	panic("unsupported response")
}

type LoginAwardInfo_In struct {
}

func (this *LoginAwardInfo_In) Process(session *net.Session) {
	g_InHandler.LoginAwardInfo(session, this)
}

func (this *LoginAwardInfo_In) TypeName() string {
	return "event.login_award_info.in"
}

func (this *LoginAwardInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 24, 1
}

type LoginAwardInfo_Out struct {
	Record         int32 `json:"record"`
	TotalLoginDays int32 `json:"total_login_days"`
}

func (this *LoginAwardInfo_Out) Process(session *net.Session) {
	g_OutHandler.LoginAwardInfo(session, this)
}

func (this *LoginAwardInfo_Out) TypeName() string {
	return "event.login_award_info.out"
}

func (this *LoginAwardInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 24, 1
}

func (this *LoginAwardInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type TakeLoginAward_In struct {
	Day int32 `json:"day"`
}

func (this *TakeLoginAward_In) Process(session *net.Session) {
	g_InHandler.TakeLoginAward(session, this)
}

func (this *TakeLoginAward_In) TypeName() string {
	return "event.take_login_award.in"
}

func (this *TakeLoginAward_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 24, 2
}

type TakeLoginAward_Out struct {
}

func (this *TakeLoginAward_Out) Process(session *net.Session) {
	g_OutHandler.TakeLoginAward(session, this)
}

func (this *TakeLoginAward_Out) TypeName() string {
	return "event.take_login_award.out"
}

func (this *TakeLoginAward_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 24, 2
}

func (this *TakeLoginAward_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetEvents_In struct {
}

func (this *GetEvents_In) Process(session *net.Session) {
	g_InHandler.GetEvents(session, this)
}

func (this *GetEvents_In) TypeName() string {
	return "event.get_events.in"
}

func (this *GetEvents_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 24, 3
}

type GetEvents_Out struct {
	Events   []GetEvents_Out_Events   `json:"events"`
	Specials []GetEvents_Out_Specials `json:"specials"`
}

type GetEvents_Out_Events struct {
	EventId       int16 `json:"event_id"`
	Process       int32 `json:"process"`
	PlayerProcess int32 `json:"player_process"`
	Page          int32 `json:"page"`
	IsAward       bool  `json:"is_award"`
}

type GetEvents_Out_Specials struct {
	Sign   []byte                          `json:"sign"`
	Params []GetEvents_Out_Specials_Params `json:"params"`
}

type GetEvents_Out_Specials_Params struct {
	Key []byte `json:"key"`
	Val []byte `json:"val"`
}

func (this *GetEvents_Out) Process(session *net.Session) {
	g_OutHandler.GetEvents(session, this)
}

func (this *GetEvents_Out) TypeName() string {
	return "event.get_events.out"
}

func (this *GetEvents_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 24, 3
}

func (this *GetEvents_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetEventAward_In struct {
	EventId int16 `json:"event_id"`
	Page    int32 `json:"page"`
	Param1  int32 `json:"param1"`
	Param2  int32 `json:"param2"`
	Param3  int32 `json:"param3"`
}

func (this *GetEventAward_In) Process(session *net.Session) {
	g_InHandler.GetEventAward(session, this)
}

func (this *GetEventAward_In) TypeName() string {
	return "event.get_event_award.in"
}

func (this *GetEventAward_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 24, 4
}

type GetEventAward_Out struct {
	Result int8  `json:"result"`
	Award  int32 `json:"award"`
}

func (this *GetEventAward_Out) Process(session *net.Session) {
	g_OutHandler.GetEventAward(session, this)
}

func (this *GetEventAward_Out) TypeName() string {
	return "event.get_event_award.out"
}

func (this *GetEventAward_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 24, 4
}

func (this *GetEventAward_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type PlayerEventPhysicalCost_In struct {
}

func (this *PlayerEventPhysicalCost_In) Process(session *net.Session) {
	g_InHandler.PlayerEventPhysicalCost(session, this)
}

func (this *PlayerEventPhysicalCost_In) TypeName() string {
	return "event.player_event_physical_cost.in"
}

func (this *PlayerEventPhysicalCost_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 24, 6
}

type PlayerEventPhysicalCost_Out struct {
	Value int32 `json:"value"`
}

func (this *PlayerEventPhysicalCost_Out) Process(session *net.Session) {
	g_OutHandler.PlayerEventPhysicalCost(session, this)
}

func (this *PlayerEventPhysicalCost_Out) TypeName() string {
	return "event.player_event_physical_cost.out"
}

func (this *PlayerEventPhysicalCost_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 24, 6
}

func (this *PlayerEventPhysicalCost_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type PlayerMonthCardInfo_In struct {
}

func (this *PlayerMonthCardInfo_In) Process(session *net.Session) {
	g_InHandler.PlayerMonthCardInfo(session, this)
}

func (this *PlayerMonthCardInfo_In) TypeName() string {
	return "event.player_month_card_info.in"
}

func (this *PlayerMonthCardInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 24, 7
}

type PlayerMonthCardInfo_Out struct {
	BeginTime int64 `json:"BeginTime"`
	EndTime   int64 `json:"EndTime"`
}

func (this *PlayerMonthCardInfo_Out) Process(session *net.Session) {
	g_OutHandler.PlayerMonthCardInfo(session, this)
}

func (this *PlayerMonthCardInfo_Out) TypeName() string {
	return "event.player_month_card_info.out"
}

func (this *PlayerMonthCardInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 24, 7
}

func (this *PlayerMonthCardInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetSevenInfo_In struct {
}

func (this *GetSevenInfo_In) Process(session *net.Session) {
	g_InHandler.GetSevenInfo(session, this)
}

func (this *GetSevenInfo_In) TypeName() string {
	return "event.get_seven_info.in"
}

func (this *GetSevenInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 24, 8
}

type GetSevenInfo_Out struct {
	Status   int8  `json:"status"`
	Day      int16 `json:"day"`
	Schedule int32 `json:"schedule"`
}

func (this *GetSevenInfo_Out) Process(session *net.Session) {
	g_OutHandler.GetSevenInfo(session, this)
}

func (this *GetSevenInfo_Out) TypeName() string {
	return "event.get_seven_info.out"
}

func (this *GetSevenInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 24, 8
}

func (this *GetSevenInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetSevenAward_In struct {
}

func (this *GetSevenAward_In) Process(session *net.Session) {
	g_InHandler.GetSevenAward(session, this)
}

func (this *GetSevenAward_In) TypeName() string {
	return "event.get_seven_award.in"
}

func (this *GetSevenAward_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 24, 9
}

type GetSevenAward_Out struct {
	Result int8 `json:"result"`
}

func (this *GetSevenAward_Out) Process(session *net.Session) {
	g_OutHandler.GetSevenAward(session, this)
}

func (this *GetSevenAward_Out) TypeName() string {
	return "event.get_seven_award.out"
}

func (this *GetSevenAward_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 24, 9
}

func (this *GetSevenAward_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetRichmanClubInfo_In struct {
}

func (this *GetRichmanClubInfo_In) Process(session *net.Session) {
	g_InHandler.GetRichmanClubInfo(session, this)
}

func (this *GetRichmanClubInfo_In) TypeName() string {
	return "event.get_richman_club_info.in"
}

func (this *GetRichmanClubInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 24, 10
}

type GetRichmanClubInfo_Out struct {
	Status   int8  `json:"status"`
	Schedule int32 `json:"schedule"`
}

func (this *GetRichmanClubInfo_Out) Process(session *net.Session) {
	g_OutHandler.GetRichmanClubInfo(session, this)
}

func (this *GetRichmanClubInfo_Out) TypeName() string {
	return "event.get_richman_club_info.out"
}

func (this *GetRichmanClubInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 24, 10
}

func (this *GetRichmanClubInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetRichmanClubAward_In struct {
	Column   int8 `json:"column"`
	Sequence int8 `json:"sequence"`
}

func (this *GetRichmanClubAward_In) Process(session *net.Session) {
	g_InHandler.GetRichmanClubAward(session, this)
}

func (this *GetRichmanClubAward_In) TypeName() string {
	return "event.get_richman_club_award.in"
}

func (this *GetRichmanClubAward_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 24, 11
}

type GetRichmanClubAward_Out struct {
	Result int8 `json:"result"`
}

func (this *GetRichmanClubAward_Out) Process(session *net.Session) {
	g_OutHandler.GetRichmanClubAward(session, this)
}

func (this *GetRichmanClubAward_Out) TypeName() string {
	return "event.get_richman_club_award.out"
}

func (this *GetRichmanClubAward_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 24, 11
}

func (this *GetRichmanClubAward_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type InfoShare_In struct {
	IsShare bool `json:"is_share"`
}

func (this *InfoShare_In) Process(session *net.Session) {
	g_InHandler.InfoShare(session, this)
}

func (this *InfoShare_In) TypeName() string {
	return "event.info_share.in"
}

func (this *InfoShare_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 24, 12
}

type InfoShare_Out struct {
	Count int16 `json:"count"`
}

func (this *InfoShare_Out) Process(session *net.Session) {
	g_OutHandler.InfoShare(session, this)
}

func (this *InfoShare_Out) TypeName() string {
	return "event.info_share.out"
}

func (this *InfoShare_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 24, 12
}

func (this *InfoShare_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type InfoGroupBuy_In struct {
}

func (this *InfoGroupBuy_In) Process(session *net.Session) {
	g_InHandler.InfoGroupBuy(session, this)
}

func (this *InfoGroupBuy_In) TypeName() string {
	return "event.info_group_buy.in"
}

func (this *InfoGroupBuy_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 24, 13
}

type InfoGroupBuy_Out struct {
	Count int32 `json:"count"`
}

func (this *InfoGroupBuy_Out) Process(session *net.Session) {
	g_OutHandler.InfoGroupBuy(session, this)
}

func (this *InfoGroupBuy_Out) TypeName() string {
	return "event.info_group_buy.out"
}

func (this *InfoGroupBuy_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 24, 13
}

func (this *InfoGroupBuy_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetIngotChangeTotal_In struct {
	IsIn bool `json:"is_in"`
}

func (this *GetIngotChangeTotal_In) Process(session *net.Session) {
	g_InHandler.GetIngotChangeTotal(session, this)
}

func (this *GetIngotChangeTotal_In) TypeName() string {
	return "event.get_ingot_change_total.in"
}

func (this *GetIngotChangeTotal_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 24, 14
}

type GetIngotChangeTotal_Out struct {
	Count int64 `json:"count"`
}

func (this *GetIngotChangeTotal_Out) Process(session *net.Session) {
	g_OutHandler.GetIngotChangeTotal(session, this)
}

func (this *GetIngotChangeTotal_Out) TypeName() string {
	return "event.get_ingot_change_total.out"
}

func (this *GetIngotChangeTotal_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 24, 14
}

func (this *GetIngotChangeTotal_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetEventTotalAward_In struct {
	Order int16 `json:"order"`
}

func (this *GetEventTotalAward_In) Process(session *net.Session) {
	g_InHandler.GetEventTotalAward(session, this)
}

func (this *GetEventTotalAward_In) TypeName() string {
	return "event.get_event_total_award.in"
}

func (this *GetEventTotalAward_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 24, 15
}

type GetEventTotalAward_Out struct {
	Result int8 `json:"result"`
}

func (this *GetEventTotalAward_Out) Process(session *net.Session) {
	g_OutHandler.GetEventTotalAward(session, this)
}

func (this *GetEventTotalAward_Out) TypeName() string {
	return "event.get_event_total_award.out"
}

func (this *GetEventTotalAward_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 24, 15
}

func (this *GetEventTotalAward_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetEventArenaRank_In struct {
}

func (this *GetEventArenaRank_In) Process(session *net.Session) {
	g_InHandler.GetEventArenaRank(session, this)
}

func (this *GetEventArenaRank_In) TypeName() string {
	return "event.get_event_arena_rank.in"
}

func (this *GetEventArenaRank_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 24, 16
}

type GetEventArenaRank_Out struct {
	Rank int32 `json:"rank"`
}

func (this *GetEventArenaRank_Out) Process(session *net.Session) {
	g_OutHandler.GetEventArenaRank(session, this)
}

func (this *GetEventArenaRank_Out) TypeName() string {
	return "event.get_event_arena_rank.out"
}

func (this *GetEventArenaRank_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 24, 16
}

func (this *GetEventArenaRank_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetEventTenDrawTimes_In struct {
}

func (this *GetEventTenDrawTimes_In) Process(session *net.Session) {
	g_InHandler.GetEventTenDrawTimes(session, this)
}

func (this *GetEventTenDrawTimes_In) TypeName() string {
	return "event.get_event_ten_draw_times.in"
}

func (this *GetEventTenDrawTimes_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 24, 17
}

type GetEventTenDrawTimes_Out struct {
	Times int32 `json:"times"`
}

func (this *GetEventTenDrawTimes_Out) Process(session *net.Session) {
	g_OutHandler.GetEventTenDrawTimes(session, this)
}

func (this *GetEventTenDrawTimes_Out) TypeName() string {
	return "event.get_event_ten_draw_times.out"
}

func (this *GetEventTenDrawTimes_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 24, 17
}

func (this *GetEventTenDrawTimes_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetEventRechargeAward_In struct {
	Page      int32 `json:"page"`
	Requireid int32 `json:"requireid"`
}

func (this *GetEventRechargeAward_In) Process(session *net.Session) {
	g_InHandler.GetEventRechargeAward(session, this)
}

func (this *GetEventRechargeAward_In) TypeName() string {
	return "event.get_event_recharge_award.in"
}

func (this *GetEventRechargeAward_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 24, 18
}

type GetEventRechargeAward_Out struct {
	IsRechage bool `json:"is_rechage"`
	IsAward   bool `json:"is_award"`
}

func (this *GetEventRechargeAward_Out) Process(session *net.Session) {
	g_OutHandler.GetEventRechargeAward(session, this)
}

func (this *GetEventRechargeAward_Out) TypeName() string {
	return "event.get_event_recharge_award.out"
}

func (this *GetEventRechargeAward_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 24, 18
}

func (this *GetEventRechargeAward_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetEventNewYear_In struct {
}

func (this *GetEventNewYear_In) Process(session *net.Session) {
	g_InHandler.GetEventNewYear(session, this)
}

func (this *GetEventNewYear_In) TypeName() string {
	return "event.get_event_new_year.in"
}

func (this *GetEventNewYear_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 24, 19
}

type GetEventNewYear_Out struct {
	Processes []GetEventNewYear_Out_Processes `json:"processes"`
}

type GetEventNewYear_Out_Processes struct {
	Day       int8  `json:"day"`
	Ingot     int32 `json:"ingot"`
	IsAwarded bool  `json:"is_awarded"`
}

func (this *GetEventNewYear_Out) Process(session *net.Session) {
	g_OutHandler.GetEventNewYear(session, this)
}

func (this *GetEventNewYear_Out) TypeName() string {
	return "event.get_event_new_year.out"
}

func (this *GetEventNewYear_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 24, 19
}

func (this *GetEventNewYear_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type QqVipContinue_In struct {
	Kind int8 `json:"kind"`
}

func (this *QqVipContinue_In) Process(session *net.Session) {
	g_InHandler.QqVipContinue(session, this)
}

func (this *QqVipContinue_In) TypeName() string {
	return "event.qq_vip_continue.in"
}

func (this *QqVipContinue_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 24, 20
}

type QqVipContinue_Out struct {
	Status int16 `json:"status"`
}

func (this *QqVipContinue_Out) Process(session *net.Session) {
	g_OutHandler.QqVipContinue(session, this)
}

func (this *QqVipContinue_Out) TypeName() string {
	return "event.qq_vip_continue.out"
}

func (this *QqVipContinue_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 24, 20
}

func (this *QqVipContinue_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type DailyOnlineInfo_In struct {
}

func (this *DailyOnlineInfo_In) Process(session *net.Session) {
	g_InHandler.DailyOnlineInfo(session, this)
}

func (this *DailyOnlineInfo_In) TypeName() string {
	return "event.daily_online_info.in"
}

func (this *DailyOnlineInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 24, 21
}

type DailyOnlineInfo_Out struct {
	TotalOnlineTime   int32 `json:"total_online_time"`
	AwardedOnlineTime int32 `json:"awarded_online_time"`
}

func (this *DailyOnlineInfo_Out) Process(session *net.Session) {
	g_OutHandler.DailyOnlineInfo(session, this)
}

func (this *DailyOnlineInfo_Out) TypeName() string {
	return "event.daily_online_info.out"
}

func (this *DailyOnlineInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 24, 21
}

func (this *DailyOnlineInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type TakeDailyOnlineAward_In struct {
}

func (this *TakeDailyOnlineAward_In) Process(session *net.Session) {
	g_InHandler.TakeDailyOnlineAward(session, this)
}

func (this *TakeDailyOnlineAward_In) TypeName() string {
	return "event.take_daily_online_award.in"
}

func (this *TakeDailyOnlineAward_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 24, 22
}

type TakeDailyOnlineAward_Out struct {
	Ok                bool  `json:"ok"`
	TotalOnlineTime   int32 `json:"total_online_time"`
	AwardedOnlineTime int32 `json:"awarded_online_time"`
}

func (this *TakeDailyOnlineAward_Out) Process(session *net.Session) {
	g_OutHandler.TakeDailyOnlineAward(session, this)
}

func (this *TakeDailyOnlineAward_Out) TypeName() string {
	return "event.take_daily_online_award.out"
}

func (this *TakeDailyOnlineAward_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 24, 22
}

func (this *TakeDailyOnlineAward_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

func (this *LoginAwardInfo_In) Decode(buffer *net.Buffer) {
}

func (this *LoginAwardInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(24)
	buffer.WriteUint8(1)
}

func (this *LoginAwardInfo_In) ByteSize() int {
	size := 2
	return size
}

func (this *LoginAwardInfo_Out) Decode(buffer *net.Buffer) {
	this.Record = int32(buffer.ReadUint32LE())
	this.TotalLoginDays = int32(buffer.ReadUint32LE())
}

func (this *LoginAwardInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(24)
	buffer.WriteUint8(1)
	buffer.WriteUint32LE(uint32(this.Record))
	buffer.WriteUint32LE(uint32(this.TotalLoginDays))
}

func (this *LoginAwardInfo_Out) ByteSize() int {
	size := 10
	return size
}

func (this *TakeLoginAward_In) Decode(buffer *net.Buffer) {
	this.Day = int32(buffer.ReadUint32LE())
}

func (this *TakeLoginAward_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(24)
	buffer.WriteUint8(2)
	buffer.WriteUint32LE(uint32(this.Day))
}

func (this *TakeLoginAward_In) ByteSize() int {
	size := 6
	return size
}

func (this *TakeLoginAward_Out) Decode(buffer *net.Buffer) {
}

func (this *TakeLoginAward_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(24)
	buffer.WriteUint8(2)
}

func (this *TakeLoginAward_Out) ByteSize() int {
	size := 2
	return size
}

func (this *GetEvents_In) Decode(buffer *net.Buffer) {
}

func (this *GetEvents_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(24)
	buffer.WriteUint8(3)
}

func (this *GetEvents_In) ByteSize() int {
	size := 2
	return size
}

func (this *GetEvents_Out) Decode(buffer *net.Buffer) {
	this.Events = make([]GetEvents_Out_Events, buffer.ReadUint8())
	for i := 0; i < len(this.Events); i++ {
		this.Events[i].Decode(buffer)
	}
	this.Specials = make([]GetEvents_Out_Specials, buffer.ReadUint8())
	for i := 0; i < len(this.Specials); i++ {
		this.Specials[i].Decode(buffer)
	}
}

func (this *GetEvents_Out_Events) Decode(buffer *net.Buffer) {
	this.EventId = int16(buffer.ReadUint16LE())
	this.Process = int32(buffer.ReadUint32LE())
	this.PlayerProcess = int32(buffer.ReadUint32LE())
	this.Page = int32(buffer.ReadUint32LE())
	this.IsAward = buffer.ReadUint8() == 1
}

func (this *GetEvents_Out_Specials) Decode(buffer *net.Buffer) {
	this.Sign = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Params = make([]GetEvents_Out_Specials_Params, buffer.ReadUint8())
	for i := 0; i < len(this.Params); i++ {
		this.Params[i].Decode(buffer)
	}
}

func (this *GetEvents_Out_Specials_Params) Decode(buffer *net.Buffer) {
	this.Key = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Val = buffer.ReadBytes(int(buffer.ReadUint16LE()))
}

func (this *GetEvents_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(24)
	buffer.WriteUint8(3)
	buffer.WriteUint8(uint8(len(this.Events)))
	for i := 0; i < len(this.Events); i++ {
		this.Events[i].Encode(buffer)
	}
	buffer.WriteUint8(uint8(len(this.Specials)))
	for i := 0; i < len(this.Specials); i++ {
		this.Specials[i].Encode(buffer)
	}
}

func (this *GetEvents_Out_Events) Encode(buffer *net.Buffer) {
	buffer.WriteUint16LE(uint16(this.EventId))
	buffer.WriteUint32LE(uint32(this.Process))
	buffer.WriteUint32LE(uint32(this.PlayerProcess))
	buffer.WriteUint32LE(uint32(this.Page))
	if this.IsAward {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *GetEvents_Out_Specials) Encode(buffer *net.Buffer) {
	buffer.WriteUint16LE(uint16(len(this.Sign)))
	buffer.WriteBytes(this.Sign)
	buffer.WriteUint8(uint8(len(this.Params)))
	for i := 0; i < len(this.Params); i++ {
		this.Params[i].Encode(buffer)
	}
}

func (this *GetEvents_Out_Specials_Params) Encode(buffer *net.Buffer) {
	buffer.WriteUint16LE(uint16(len(this.Key)))
	buffer.WriteBytes(this.Key)
	buffer.WriteUint16LE(uint16(len(this.Val)))
	buffer.WriteBytes(this.Val)
}

func (this *GetEvents_Out) ByteSize() int {
	size := 4
	size += len(this.Events) * 15
	for i := 0; i < len(this.Specials); i++ {
		size += this.Specials[i].ByteSize()
	}
	return size
}

func (this *GetEvents_Out_Specials) ByteSize() int {
	size := 3
	size += len(this.Sign)
	for i := 0; i < len(this.Params); i++ {
		size += this.Params[i].ByteSize()
	}
	return size
}

func (this *GetEvents_Out_Specials_Params) ByteSize() int {
	size := 4
	size += len(this.Key)
	size += len(this.Val)
	return size
}

func (this *GetEventAward_In) Decode(buffer *net.Buffer) {
	this.EventId = int16(buffer.ReadUint16LE())
	this.Page = int32(buffer.ReadUint32LE())
	this.Param1 = int32(buffer.ReadUint32LE())
	this.Param2 = int32(buffer.ReadUint32LE())
	this.Param3 = int32(buffer.ReadUint32LE())
}

func (this *GetEventAward_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(24)
	buffer.WriteUint8(4)
	buffer.WriteUint16LE(uint16(this.EventId))
	buffer.WriteUint32LE(uint32(this.Page))
	buffer.WriteUint32LE(uint32(this.Param1))
	buffer.WriteUint32LE(uint32(this.Param2))
	buffer.WriteUint32LE(uint32(this.Param3))
}

func (this *GetEventAward_In) ByteSize() int {
	size := 20
	return size
}

func (this *GetEventAward_Out) Decode(buffer *net.Buffer) {
	this.Result = int8(buffer.ReadUint8())
	this.Award = int32(buffer.ReadUint32LE())
}

func (this *GetEventAward_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(24)
	buffer.WriteUint8(4)
	buffer.WriteUint8(uint8(this.Result))
	buffer.WriteUint32LE(uint32(this.Award))
}

func (this *GetEventAward_Out) ByteSize() int {
	size := 7
	return size
}

func (this *PlayerEventPhysicalCost_In) Decode(buffer *net.Buffer) {
}

func (this *PlayerEventPhysicalCost_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(24)
	buffer.WriteUint8(6)
}

func (this *PlayerEventPhysicalCost_In) ByteSize() int {
	size := 2
	return size
}

func (this *PlayerEventPhysicalCost_Out) Decode(buffer *net.Buffer) {
	this.Value = int32(buffer.ReadUint32LE())
}

func (this *PlayerEventPhysicalCost_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(24)
	buffer.WriteUint8(6)
	buffer.WriteUint32LE(uint32(this.Value))
}

func (this *PlayerEventPhysicalCost_Out) ByteSize() int {
	size := 6
	return size
}

func (this *PlayerMonthCardInfo_In) Decode(buffer *net.Buffer) {
}

func (this *PlayerMonthCardInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(24)
	buffer.WriteUint8(7)
}

func (this *PlayerMonthCardInfo_In) ByteSize() int {
	size := 2
	return size
}

func (this *PlayerMonthCardInfo_Out) Decode(buffer *net.Buffer) {
	this.BeginTime = int64(buffer.ReadUint64LE())
	this.EndTime = int64(buffer.ReadUint64LE())
}

func (this *PlayerMonthCardInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(24)
	buffer.WriteUint8(7)
	buffer.WriteUint64LE(uint64(this.BeginTime))
	buffer.WriteUint64LE(uint64(this.EndTime))
}

func (this *PlayerMonthCardInfo_Out) ByteSize() int {
	size := 18
	return size
}

func (this *GetSevenInfo_In) Decode(buffer *net.Buffer) {
}

func (this *GetSevenInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(24)
	buffer.WriteUint8(8)
}

func (this *GetSevenInfo_In) ByteSize() int {
	size := 2
	return size
}

func (this *GetSevenInfo_Out) Decode(buffer *net.Buffer) {
	this.Status = int8(buffer.ReadUint8())
	this.Day = int16(buffer.ReadUint16LE())
	this.Schedule = int32(buffer.ReadUint32LE())
}

func (this *GetSevenInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(24)
	buffer.WriteUint8(8)
	buffer.WriteUint8(uint8(this.Status))
	buffer.WriteUint16LE(uint16(this.Day))
	buffer.WriteUint32LE(uint32(this.Schedule))
}

func (this *GetSevenInfo_Out) ByteSize() int {
	size := 9
	return size
}

func (this *GetSevenAward_In) Decode(buffer *net.Buffer) {
}

func (this *GetSevenAward_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(24)
	buffer.WriteUint8(9)
}

func (this *GetSevenAward_In) ByteSize() int {
	size := 2
	return size
}

func (this *GetSevenAward_Out) Decode(buffer *net.Buffer) {
	this.Result = int8(buffer.ReadUint8())
}

func (this *GetSevenAward_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(24)
	buffer.WriteUint8(9)
	buffer.WriteUint8(uint8(this.Result))
}

func (this *GetSevenAward_Out) ByteSize() int {
	size := 3
	return size
}

func (this *GetRichmanClubInfo_In) Decode(buffer *net.Buffer) {
}

func (this *GetRichmanClubInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(24)
	buffer.WriteUint8(10)
}

func (this *GetRichmanClubInfo_In) ByteSize() int {
	size := 2
	return size
}

func (this *GetRichmanClubInfo_Out) Decode(buffer *net.Buffer) {
	this.Status = int8(buffer.ReadUint8())
	this.Schedule = int32(buffer.ReadUint32LE())
}

func (this *GetRichmanClubInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(24)
	buffer.WriteUint8(10)
	buffer.WriteUint8(uint8(this.Status))
	buffer.WriteUint32LE(uint32(this.Schedule))
}

func (this *GetRichmanClubInfo_Out) ByteSize() int {
	size := 7
	return size
}

func (this *GetRichmanClubAward_In) Decode(buffer *net.Buffer) {
	this.Column = int8(buffer.ReadUint8())
	this.Sequence = int8(buffer.ReadUint8())
}

func (this *GetRichmanClubAward_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(24)
	buffer.WriteUint8(11)
	buffer.WriteUint8(uint8(this.Column))
	buffer.WriteUint8(uint8(this.Sequence))
}

func (this *GetRichmanClubAward_In) ByteSize() int {
	size := 4
	return size
}

func (this *GetRichmanClubAward_Out) Decode(buffer *net.Buffer) {
	this.Result = int8(buffer.ReadUint8())
}

func (this *GetRichmanClubAward_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(24)
	buffer.WriteUint8(11)
	buffer.WriteUint8(uint8(this.Result))
}

func (this *GetRichmanClubAward_Out) ByteSize() int {
	size := 3
	return size
}

func (this *InfoShare_In) Decode(buffer *net.Buffer) {
	this.IsShare = buffer.ReadUint8() == 1
}

func (this *InfoShare_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(24)
	buffer.WriteUint8(12)
	if this.IsShare {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *InfoShare_In) ByteSize() int {
	size := 3
	return size
}

func (this *InfoShare_Out) Decode(buffer *net.Buffer) {
	this.Count = int16(buffer.ReadUint16LE())
}

func (this *InfoShare_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(24)
	buffer.WriteUint8(12)
	buffer.WriteUint16LE(uint16(this.Count))
}

func (this *InfoShare_Out) ByteSize() int {
	size := 4
	return size
}

func (this *InfoGroupBuy_In) Decode(buffer *net.Buffer) {
}

func (this *InfoGroupBuy_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(24)
	buffer.WriteUint8(13)
}

func (this *InfoGroupBuy_In) ByteSize() int {
	size := 2
	return size
}

func (this *InfoGroupBuy_Out) Decode(buffer *net.Buffer) {
	this.Count = int32(buffer.ReadUint32LE())
}

func (this *InfoGroupBuy_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(24)
	buffer.WriteUint8(13)
	buffer.WriteUint32LE(uint32(this.Count))
}

func (this *InfoGroupBuy_Out) ByteSize() int {
	size := 6
	return size
}

func (this *GetIngotChangeTotal_In) Decode(buffer *net.Buffer) {
	this.IsIn = buffer.ReadUint8() == 1
}

func (this *GetIngotChangeTotal_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(24)
	buffer.WriteUint8(14)
	if this.IsIn {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *GetIngotChangeTotal_In) ByteSize() int {
	size := 3
	return size
}

func (this *GetIngotChangeTotal_Out) Decode(buffer *net.Buffer) {
	this.Count = int64(buffer.ReadUint64LE())
}

func (this *GetIngotChangeTotal_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(24)
	buffer.WriteUint8(14)
	buffer.WriteUint64LE(uint64(this.Count))
}

func (this *GetIngotChangeTotal_Out) ByteSize() int {
	size := 10
	return size
}

func (this *GetEventTotalAward_In) Decode(buffer *net.Buffer) {
	this.Order = int16(buffer.ReadUint16LE())
}

func (this *GetEventTotalAward_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(24)
	buffer.WriteUint8(15)
	buffer.WriteUint16LE(uint16(this.Order))
}

func (this *GetEventTotalAward_In) ByteSize() int {
	size := 4
	return size
}

func (this *GetEventTotalAward_Out) Decode(buffer *net.Buffer) {
	this.Result = int8(buffer.ReadUint8())
}

func (this *GetEventTotalAward_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(24)
	buffer.WriteUint8(15)
	buffer.WriteUint8(uint8(this.Result))
}

func (this *GetEventTotalAward_Out) ByteSize() int {
	size := 3
	return size
}

func (this *GetEventArenaRank_In) Decode(buffer *net.Buffer) {
}

func (this *GetEventArenaRank_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(24)
	buffer.WriteUint8(16)
}

func (this *GetEventArenaRank_In) ByteSize() int {
	size := 2
	return size
}

func (this *GetEventArenaRank_Out) Decode(buffer *net.Buffer) {
	this.Rank = int32(buffer.ReadUint32LE())
}

func (this *GetEventArenaRank_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(24)
	buffer.WriteUint8(16)
	buffer.WriteUint32LE(uint32(this.Rank))
}

func (this *GetEventArenaRank_Out) ByteSize() int {
	size := 6
	return size
}

func (this *GetEventTenDrawTimes_In) Decode(buffer *net.Buffer) {
}

func (this *GetEventTenDrawTimes_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(24)
	buffer.WriteUint8(17)
}

func (this *GetEventTenDrawTimes_In) ByteSize() int {
	size := 2
	return size
}

func (this *GetEventTenDrawTimes_Out) Decode(buffer *net.Buffer) {
	this.Times = int32(buffer.ReadUint32LE())
}

func (this *GetEventTenDrawTimes_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(24)
	buffer.WriteUint8(17)
	buffer.WriteUint32LE(uint32(this.Times))
}

func (this *GetEventTenDrawTimes_Out) ByteSize() int {
	size := 6
	return size
}

func (this *GetEventRechargeAward_In) Decode(buffer *net.Buffer) {
	this.Page = int32(buffer.ReadUint32LE())
	this.Requireid = int32(buffer.ReadUint32LE())
}

func (this *GetEventRechargeAward_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(24)
	buffer.WriteUint8(18)
	buffer.WriteUint32LE(uint32(this.Page))
	buffer.WriteUint32LE(uint32(this.Requireid))
}

func (this *GetEventRechargeAward_In) ByteSize() int {
	size := 10
	return size
}

func (this *GetEventRechargeAward_Out) Decode(buffer *net.Buffer) {
	this.IsRechage = buffer.ReadUint8() == 1
	this.IsAward = buffer.ReadUint8() == 1
}

func (this *GetEventRechargeAward_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(24)
	buffer.WriteUint8(18)
	if this.IsRechage {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
	if this.IsAward {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *GetEventRechargeAward_Out) ByteSize() int {
	size := 4
	return size
}

func (this *GetEventNewYear_In) Decode(buffer *net.Buffer) {
}

func (this *GetEventNewYear_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(24)
	buffer.WriteUint8(19)
}

func (this *GetEventNewYear_In) ByteSize() int {
	size := 2
	return size
}

func (this *GetEventNewYear_Out) Decode(buffer *net.Buffer) {
	this.Processes = make([]GetEventNewYear_Out_Processes, buffer.ReadUint8())
	for i := 0; i < len(this.Processes); i++ {
		this.Processes[i].Decode(buffer)
	}
}

func (this *GetEventNewYear_Out_Processes) Decode(buffer *net.Buffer) {
	this.Day = int8(buffer.ReadUint8())
	this.Ingot = int32(buffer.ReadUint32LE())
	this.IsAwarded = buffer.ReadUint8() == 1
}

func (this *GetEventNewYear_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(24)
	buffer.WriteUint8(19)
	buffer.WriteUint8(uint8(len(this.Processes)))
	for i := 0; i < len(this.Processes); i++ {
		this.Processes[i].Encode(buffer)
	}
}

func (this *GetEventNewYear_Out_Processes) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(this.Day))
	buffer.WriteUint32LE(uint32(this.Ingot))
	if this.IsAwarded {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *GetEventNewYear_Out) ByteSize() int {
	size := 3
	size += len(this.Processes) * 6
	return size
}

func (this *QqVipContinue_In) Decode(buffer *net.Buffer) {
	this.Kind = int8(buffer.ReadUint8())
}

func (this *QqVipContinue_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(24)
	buffer.WriteUint8(20)
	buffer.WriteUint8(uint8(this.Kind))
}

func (this *QqVipContinue_In) ByteSize() int {
	size := 3
	return size
}

func (this *QqVipContinue_Out) Decode(buffer *net.Buffer) {
	this.Status = int16(buffer.ReadUint16LE())
}

func (this *QqVipContinue_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(24)
	buffer.WriteUint8(20)
	buffer.WriteUint16LE(uint16(this.Status))
}

func (this *QqVipContinue_Out) ByteSize() int {
	size := 4
	return size
}

func (this *DailyOnlineInfo_In) Decode(buffer *net.Buffer) {
}

func (this *DailyOnlineInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(24)
	buffer.WriteUint8(21)
}

func (this *DailyOnlineInfo_In) ByteSize() int {
	size := 2
	return size
}

func (this *DailyOnlineInfo_Out) Decode(buffer *net.Buffer) {
	this.TotalOnlineTime = int32(buffer.ReadUint32LE())
	this.AwardedOnlineTime = int32(buffer.ReadUint32LE())
}

func (this *DailyOnlineInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(24)
	buffer.WriteUint8(21)
	buffer.WriteUint32LE(uint32(this.TotalOnlineTime))
	buffer.WriteUint32LE(uint32(this.AwardedOnlineTime))
}

func (this *DailyOnlineInfo_Out) ByteSize() int {
	size := 10
	return size
}

func (this *TakeDailyOnlineAward_In) Decode(buffer *net.Buffer) {
}

func (this *TakeDailyOnlineAward_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(24)
	buffer.WriteUint8(22)
}

func (this *TakeDailyOnlineAward_In) ByteSize() int {
	size := 2
	return size
}

func (this *TakeDailyOnlineAward_Out) Decode(buffer *net.Buffer) {
	this.Ok = buffer.ReadUint8() == 1
	this.TotalOnlineTime = int32(buffer.ReadUint32LE())
	this.AwardedOnlineTime = int32(buffer.ReadUint32LE())
}

func (this *TakeDailyOnlineAward_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(24)
	buffer.WriteUint8(22)
	if this.Ok {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
	buffer.WriteUint32LE(uint32(this.TotalOnlineTime))
	buffer.WriteUint32LE(uint32(this.AwardedOnlineTime))
}

func (this *TakeDailyOnlineAward_Out) ByteSize() int {
	size := 11
	return size
}
