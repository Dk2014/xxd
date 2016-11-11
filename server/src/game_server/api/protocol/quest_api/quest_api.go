package quest_api

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
	UpdateQuest(*net.Session, *UpdateQuest_In)
	GetDailyInfo(*net.Session, *GetDailyInfo_In)
	AwardDaily(*net.Session, *AwardDaily_In)
	Guide(*net.Session, *Guide_In)
	GetExtendQuestInfoByNpcId(*net.Session, *GetExtendQuestInfoByNpcId_In)
	TakeExtendQuestAward(*net.Session, *TakeExtendQuestAward_In)
	GetPannelQuestInfo(*net.Session, *GetPannelQuestInfo_In)
	GiveUpAdditionQuest(*net.Session, *GiveUpAdditionQuest_In)
	TakeAdditionQuest(*net.Session, *TakeAdditionQuest_In)
	TakeAdditionQuestAward(*net.Session, *TakeAdditionQuestAward_In)
	GetAdditionQuest(*net.Session, *GetAdditionQuest_In)
	RefreshAdditionQuest(*net.Session, *RefreshAdditionQuest_In)
	TakeQuestStarsAwaded(*net.Session, *TakeQuestStarsAwaded_In)
}

type OutHandler interface {
	UpdateQuest(*net.Session, *UpdateQuest_Out)
	GetDailyInfo(*net.Session, *GetDailyInfo_Out)
	AwardDaily(*net.Session, *AwardDaily_Out)
	NotifyDailyChange(*net.Session, *NotifyDailyChange_Out)
	Guide(*net.Session, *Guide_Out)
	GetExtendQuestInfoByNpcId(*net.Session, *GetExtendQuestInfoByNpcId_Out)
	TakeExtendQuestAward(*net.Session, *TakeExtendQuestAward_Out)
	GetPannelQuestInfo(*net.Session, *GetPannelQuestInfo_Out)
	GiveUpAdditionQuest(*net.Session, *GiveUpAdditionQuest_Out)
	TakeAdditionQuest(*net.Session, *TakeAdditionQuest_Out)
	TakeAdditionQuestAward(*net.Session, *TakeAdditionQuestAward_Out)
	GetAdditionQuest(*net.Session, *GetAdditionQuest_Out)
	RefreshAdditionQuest(*net.Session, *RefreshAdditionQuest_Out)
	TakeQuestStarsAwaded(*net.Session, *TakeQuestStarsAwaded_Out)
}

func DecodeIn(message []byte) Request {
	var actionId = message[0]
	var buffer = net.NewBuffer(message[1:])
	switch actionId {
	case 1:
		request := new(UpdateQuest_In)
		request.Decode(buffer)
		return request
	case 2:
		request := new(GetDailyInfo_In)
		request.Decode(buffer)
		return request
	case 3:
		request := new(AwardDaily_In)
		request.Decode(buffer)
		return request
	case 5:
		request := new(Guide_In)
		request.Decode(buffer)
		return request
	case 6:
		request := new(GetExtendQuestInfoByNpcId_In)
		request.Decode(buffer)
		return request
	case 7:
		request := new(TakeExtendQuestAward_In)
		request.Decode(buffer)
		return request
	case 8:
		request := new(GetPannelQuestInfo_In)
		request.Decode(buffer)
		return request
	case 10:
		request := new(GiveUpAdditionQuest_In)
		request.Decode(buffer)
		return request
	case 11:
		request := new(TakeAdditionQuest_In)
		request.Decode(buffer)
		return request
	case 12:
		request := new(TakeAdditionQuestAward_In)
		request.Decode(buffer)
		return request
	case 13:
		request := new(GetAdditionQuest_In)
		request.Decode(buffer)
		return request
	case 14:
		request := new(RefreshAdditionQuest_In)
		request.Decode(buffer)
		return request
	case 15:
		request := new(TakeQuestStarsAwaded_In)
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
		request := new(UpdateQuest_Out)
		request.Decode(buffer)
		return request
	case 2:
		request := new(GetDailyInfo_Out)
		request.Decode(buffer)
		return request
	case 3:
		request := new(AwardDaily_Out)
		request.Decode(buffer)
		return request
	case 4:
		request := new(NotifyDailyChange_Out)
		request.Decode(buffer)
		return request
	case 5:
		request := new(Guide_Out)
		request.Decode(buffer)
		return request
	case 6:
		request := new(GetExtendQuestInfoByNpcId_Out)
		request.Decode(buffer)
		return request
	case 7:
		request := new(TakeExtendQuestAward_Out)
		request.Decode(buffer)
		return request
	case 8:
		request := new(GetPannelQuestInfo_Out)
		request.Decode(buffer)
		return request
	case 10:
		request := new(GiveUpAdditionQuest_Out)
		request.Decode(buffer)
		return request
	case 11:
		request := new(TakeAdditionQuest_Out)
		request.Decode(buffer)
		return request
	case 12:
		request := new(TakeAdditionQuestAward_Out)
		request.Decode(buffer)
		return request
	case 13:
		request := new(GetAdditionQuest_Out)
		request.Decode(buffer)
		return request
	case 14:
		request := new(RefreshAdditionQuest_Out)
		request.Decode(buffer)
		return request
	case 15:
		request := new(TakeQuestStarsAwaded_Out)
		request.Decode(buffer)
		return request
	}
	_ = buffer
	panic("unsupported response")
}

type GuideType int8

const (
	GUIDE_TYPE_GUIDE_ENTER_MISSION                    GuideType = 0
	GUIDE_TYPE_GUIDE_SKILL_USE                        GuideType = 1
	GUIDE_TYPE_GUIDE_ADVANCED_SKILL_USE               GuideType = 2
	GUIDE_TYPE_GUIDE_MAIN_ROLE_SKILL_EQUIP            GuideType = 3
	GUIDE_TYPE_GUIDE_MAIN_ROLE_SKILL_USE              GuideType = 4
	GUIDE_TYPE_GUIDE_EQUIP_REFINE                     GuideType = 5
	GUIDE_TYPE_GUIDE_BATTLE_ITEM_ZHIXUECAO_USE        GuideType = 6
	GUIDE_TYPE_GUIDE_BUDDY_ADVANCED_SKILL_EQUIP       GuideType = 7
	GUIDE_TYPE_GUIDE_BUDDY_ADVANCED_SKILL_USE         GuideType = 8
	GUIDE_TYPE_GUIDE_GHOST_EQUIP                      GuideType = 9
	GUIDE_TYPE_GUIDE_GHOST_POWER_LOOK                 GuideType = 10
	GUIDE_TYPE_GUIDE_GHOST_BATTLE_USE                 GuideType = 11
	GUIDE_TYPE_GUIDE_PET_EQUIP                        GuideType = 12
	GUIDE_TYPE_GUIDE_PET_BATTLE_USE                   GuideType = 13
	GUIDE_TYPE_GUIDE_PET_CATCH                        GuideType = 14
	GUIDE_TYPE_GUIDE_HARD_LEVEL                       GuideType = 15
	GUIDE_TYPE_GUIDE_TIANJIE                          GuideType = 16
	GUIDE_TYPE_GUIDE_GOTO_ZHUBAO                      GuideType = 17
	GUIDE_TYPE_GUIDE_SWORD_SOUL                       GuideType = 18
	GUIDE_TYPE_GUIDE_FIRST_BATTLE                     GuideType = 19
	GUIDE_TYPE_GUIDE_MAIN_ROLE_SKILL_EQUIP_FAILE      GuideType = 20
	GUIDE_TYPE_GUIDE_ENTER_MISSION_SECOND             GuideType = 21
	GUIDE_TYPE_GUIDE_BUDDY_EQUIP                      GuideType = 22
	GUIDE_TYPE_GUIDE_BUDDY_USE_SKILL_DAODUN           GuideType = 23
	GUIDE_TYPE_GUIDE_EQUIP_USE_SKILL_FENGJUANCANSHENG GuideType = 24
	GUIDE_TYPE_GUIDE_FRIENDSHIP                       GuideType = 25
	GUIDE_TYPE_GUIDE_EQUIP_ROLE_3                     GuideType = 26
	GUIDE_TYPE_GUIDE_EQUIP_ROLE_4                     GuideType = 27
)

type GuideAction int8

const (
	GUIDE_ACTION_GUIDE_ACCEPT GuideAction = 0
	GUIDE_ACTION_GUIDE_FINISH GuideAction = 1
)

type UpdateQuest_In struct {
}

func (this *UpdateQuest_In) Process(session *net.Session) {
	g_InHandler.UpdateQuest(session, this)
}

func (this *UpdateQuest_In) TypeName() string {
	return "quest.update_quest.in"
}

func (this *UpdateQuest_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 13, 1
}

type UpdateQuest_Out struct {
}

func (this *UpdateQuest_Out) Process(session *net.Session) {
	g_OutHandler.UpdateQuest(session, this)
}

func (this *UpdateQuest_Out) TypeName() string {
	return "quest.update_quest.out"
}

func (this *UpdateQuest_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 13, 1
}

func (this *UpdateQuest_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetDailyInfo_In struct {
}

func (this *GetDailyInfo_In) Process(session *net.Session) {
	g_InHandler.GetDailyInfo(session, this)
}

func (this *GetDailyInfo_In) TypeName() string {
	return "quest.get_daily_info.in"
}

func (this *GetDailyInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 13, 2
}

type GetDailyInfo_Out struct {
	Quest []GetDailyInfo_Out_Quest `json:"quest"`
}

type GetDailyInfo_Out_Quest struct {
	Id          int16 `json:"id"`
	FinishCount int16 `json:"finish_count"`
	AwardState  int8  `json:"award_state"`
}

func (this *GetDailyInfo_Out) Process(session *net.Session) {
	g_OutHandler.GetDailyInfo(session, this)
}

func (this *GetDailyInfo_Out) TypeName() string {
	return "quest.get_daily_info.out"
}

func (this *GetDailyInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 13, 2
}

func (this *GetDailyInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type AwardDaily_In struct {
	Id int16 `json:"id"`
}

func (this *AwardDaily_In) Process(session *net.Session) {
	g_InHandler.AwardDaily(session, this)
}

func (this *AwardDaily_In) TypeName() string {
	return "quest.award_daily.in"
}

func (this *AwardDaily_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 13, 3
}

type AwardDaily_Out struct {
}

func (this *AwardDaily_Out) Process(session *net.Session) {
	g_OutHandler.AwardDaily(session, this)
}

func (this *AwardDaily_Out) TypeName() string {
	return "quest.award_daily.out"
}

func (this *AwardDaily_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 13, 3
}

func (this *AwardDaily_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type NotifyDailyChange_Out struct {
	Id          int16 `json:"id"`
	FinishCount int16 `json:"finish_count"`
	AwardState  int8  `json:"award_state"`
}

func (this *NotifyDailyChange_Out) Process(session *net.Session) {
	g_OutHandler.NotifyDailyChange(session, this)
}

func (this *NotifyDailyChange_Out) TypeName() string {
	return "quest.notify_daily_change.out"
}

func (this *NotifyDailyChange_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 13, 4
}

func (this *NotifyDailyChange_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type Guide_In struct {
	GuideType GuideType   `json:"guide_type"`
	Action    GuideAction `json:"action"`
}

func (this *Guide_In) Process(session *net.Session) {
	g_InHandler.Guide(session, this)
}

func (this *Guide_In) TypeName() string {
	return "quest.guide.in"
}

func (this *Guide_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 13, 5
}

type Guide_Out struct {
}

func (this *Guide_Out) Process(session *net.Session) {
	g_OutHandler.Guide(session, this)
}

func (this *Guide_Out) TypeName() string {
	return "quest.guide.out"
}

func (this *Guide_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 13, 5
}

func (this *Guide_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetExtendQuestInfoByNpcId_In struct {
	NpcId int32 `json:"npc_id"`
}

func (this *GetExtendQuestInfoByNpcId_In) Process(session *net.Session) {
	g_InHandler.GetExtendQuestInfoByNpcId(session, this)
}

func (this *GetExtendQuestInfoByNpcId_In) TypeName() string {
	return "quest.get_extend_quest_info_by_npc_id.in"
}

func (this *GetExtendQuestInfoByNpcId_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 13, 6
}

type GetExtendQuestInfoByNpcId_Out struct {
	Quest []GetExtendQuestInfoByNpcId_Out_Quest `json:"quest"`
}

type GetExtendQuestInfoByNpcId_Out_Quest struct {
	Id       int32 `json:"id"`
	Progress int16 `json:"progress"`
	State    int8  `json:"state"`
}

func (this *GetExtendQuestInfoByNpcId_Out) Process(session *net.Session) {
	g_OutHandler.GetExtendQuestInfoByNpcId(session, this)
}

func (this *GetExtendQuestInfoByNpcId_Out) TypeName() string {
	return "quest.get_extend_quest_info_by_npc_id.out"
}

func (this *GetExtendQuestInfoByNpcId_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 13, 6
}

func (this *GetExtendQuestInfoByNpcId_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type TakeExtendQuestAward_In struct {
	QuestId int32 `json:"quest_id"`
}

func (this *TakeExtendQuestAward_In) Process(session *net.Session) {
	g_InHandler.TakeExtendQuestAward(session, this)
}

func (this *TakeExtendQuestAward_In) TypeName() string {
	return "quest.take_extend_quest_award.in"
}

func (this *TakeExtendQuestAward_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 13, 7
}

type TakeExtendQuestAward_Out struct {
	QuestId int32 `json:"quest_id"`
}

func (this *TakeExtendQuestAward_Out) Process(session *net.Session) {
	g_OutHandler.TakeExtendQuestAward(session, this)
}

func (this *TakeExtendQuestAward_Out) TypeName() string {
	return "quest.take_extend_quest_award.out"
}

func (this *TakeExtendQuestAward_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 13, 7
}

func (this *TakeExtendQuestAward_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetPannelQuestInfo_In struct {
}

func (this *GetPannelQuestInfo_In) Process(session *net.Session) {
	g_InHandler.GetPannelQuestInfo(session, this)
}

func (this *GetPannelQuestInfo_In) TypeName() string {
	return "quest.get_pannel_quest_info.in"
}

func (this *GetPannelQuestInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 13, 8
}

type GetPannelQuestInfo_Out struct {
	CurStars int32                          `json:"cur_stars"`
	Awarded  []byte                         `json:"awarded"`
	Quest    []GetPannelQuestInfo_Out_Quest `json:"quest"`
}

type GetPannelQuestInfo_Out_Quest struct {
	QuestClass int8  `json:"quest_class"`
	Id         int32 `json:"id"`
	Progress   int16 `json:"progress"`
	State      int8  `json:"state"`
}

func (this *GetPannelQuestInfo_Out) Process(session *net.Session) {
	g_OutHandler.GetPannelQuestInfo(session, this)
}

func (this *GetPannelQuestInfo_Out) TypeName() string {
	return "quest.get_pannel_quest_info.out"
}

func (this *GetPannelQuestInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 13, 8
}

func (this *GetPannelQuestInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GiveUpAdditionQuest_In struct {
	QuestId int32 `json:"quest_id"`
}

func (this *GiveUpAdditionQuest_In) Process(session *net.Session) {
	g_InHandler.GiveUpAdditionQuest(session, this)
}

func (this *GiveUpAdditionQuest_In) TypeName() string {
	return "quest.give_up_addition_quest.in"
}

func (this *GiveUpAdditionQuest_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 13, 10
}

type GiveUpAdditionQuest_Out struct {
}

func (this *GiveUpAdditionQuest_Out) Process(session *net.Session) {
	g_OutHandler.GiveUpAdditionQuest(session, this)
}

func (this *GiveUpAdditionQuest_Out) TypeName() string {
	return "quest.give_up_addition_quest.out"
}

func (this *GiveUpAdditionQuest_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 13, 10
}

func (this *GiveUpAdditionQuest_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type TakeAdditionQuest_In struct {
	QuestId int32 `json:"quest_id"`
}

func (this *TakeAdditionQuest_In) Process(session *net.Session) {
	g_InHandler.TakeAdditionQuest(session, this)
}

func (this *TakeAdditionQuest_In) TypeName() string {
	return "quest.take_addition_quest.in"
}

func (this *TakeAdditionQuest_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 13, 11
}

type TakeAdditionQuest_Out struct {
	Success bool `json:"success"`
}

func (this *TakeAdditionQuest_Out) Process(session *net.Session) {
	g_OutHandler.TakeAdditionQuest(session, this)
}

func (this *TakeAdditionQuest_Out) TypeName() string {
	return "quest.take_addition_quest.out"
}

func (this *TakeAdditionQuest_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 13, 11
}

func (this *TakeAdditionQuest_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type TakeAdditionQuestAward_In struct {
	QuestId int32 `json:"quest_id"`
}

func (this *TakeAdditionQuestAward_In) Process(session *net.Session) {
	g_InHandler.TakeAdditionQuestAward(session, this)
}

func (this *TakeAdditionQuestAward_In) TypeName() string {
	return "quest.take_addition_quest_award.in"
}

func (this *TakeAdditionQuestAward_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 13, 12
}

type TakeAdditionQuestAward_Out struct {
}

func (this *TakeAdditionQuestAward_Out) Process(session *net.Session) {
	g_OutHandler.TakeAdditionQuestAward(session, this)
}

func (this *TakeAdditionQuestAward_Out) TypeName() string {
	return "quest.take_addition_quest_award.out"
}

func (this *TakeAdditionQuestAward_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 13, 12
}

func (this *TakeAdditionQuestAward_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetAdditionQuest_In struct {
}

func (this *GetAdditionQuest_In) Process(session *net.Session) {
	g_InHandler.GetAdditionQuest(session, this)
}

func (this *GetAdditionQuest_In) TypeName() string {
	return "quest.get_addition_quest.in"
}

func (this *GetAdditionQuest_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 13, 13
}

type GetAdditionQuest_Out struct {
	Quest []GetAdditionQuest_Out_Quest `json:"quest"`
}

type GetAdditionQuest_Out_Quest struct {
	QuestId  int32 `json:"quest_id"`
	Progress int16 `json:"progress"`
	State    int8  `json:"state"`
}

func (this *GetAdditionQuest_Out) Process(session *net.Session) {
	g_OutHandler.GetAdditionQuest(session, this)
}

func (this *GetAdditionQuest_Out) TypeName() string {
	return "quest.get_addition_quest.out"
}

func (this *GetAdditionQuest_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 13, 13
}

func (this *GetAdditionQuest_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type RefreshAdditionQuest_In struct {
	QuestId int32 `json:"quest_id"`
}

func (this *RefreshAdditionQuest_In) Process(session *net.Session) {
	g_InHandler.RefreshAdditionQuest(session, this)
}

func (this *RefreshAdditionQuest_In) TypeName() string {
	return "quest.refresh_addition_quest.in"
}

func (this *RefreshAdditionQuest_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 13, 14
}

type RefreshAdditionQuest_Out struct {
}

func (this *RefreshAdditionQuest_Out) Process(session *net.Session) {
	g_OutHandler.RefreshAdditionQuest(session, this)
}

func (this *RefreshAdditionQuest_Out) TypeName() string {
	return "quest.refresh_addition_quest.out"
}

func (this *RefreshAdditionQuest_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 13, 14
}

func (this *RefreshAdditionQuest_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type TakeQuestStarsAwaded_In struct {
	StarsLevel int32 `json:"stars_level"`
}

func (this *TakeQuestStarsAwaded_In) Process(session *net.Session) {
	g_InHandler.TakeQuestStarsAwaded(session, this)
}

func (this *TakeQuestStarsAwaded_In) TypeName() string {
	return "quest.take_quest_stars_awaded.in"
}

func (this *TakeQuestStarsAwaded_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 13, 15
}

type TakeQuestStarsAwaded_Out struct {
	Result bool `json:"result"`
}

func (this *TakeQuestStarsAwaded_Out) Process(session *net.Session) {
	g_OutHandler.TakeQuestStarsAwaded(session, this)
}

func (this *TakeQuestStarsAwaded_Out) TypeName() string {
	return "quest.take_quest_stars_awaded.out"
}

func (this *TakeQuestStarsAwaded_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 13, 15
}

func (this *TakeQuestStarsAwaded_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

func (this *UpdateQuest_In) Decode(buffer *net.Buffer) {
}

func (this *UpdateQuest_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(13)
	buffer.WriteUint8(1)
}

func (this *UpdateQuest_In) ByteSize() int {
	size := 2
	return size
}

func (this *UpdateQuest_Out) Decode(buffer *net.Buffer) {
}

func (this *UpdateQuest_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(13)
	buffer.WriteUint8(1)
}

func (this *UpdateQuest_Out) ByteSize() int {
	size := 2
	return size
}

func (this *GetDailyInfo_In) Decode(buffer *net.Buffer) {
}

func (this *GetDailyInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(13)
	buffer.WriteUint8(2)
}

func (this *GetDailyInfo_In) ByteSize() int {
	size := 2
	return size
}

func (this *GetDailyInfo_Out) Decode(buffer *net.Buffer) {
	this.Quest = make([]GetDailyInfo_Out_Quest, buffer.ReadUint8())
	for i := 0; i < len(this.Quest); i++ {
		this.Quest[i].Decode(buffer)
	}
}

func (this *GetDailyInfo_Out_Quest) Decode(buffer *net.Buffer) {
	this.Id = int16(buffer.ReadUint16LE())
	this.FinishCount = int16(buffer.ReadUint16LE())
	this.AwardState = int8(buffer.ReadUint8())
}

func (this *GetDailyInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(13)
	buffer.WriteUint8(2)
	buffer.WriteUint8(uint8(len(this.Quest)))
	for i := 0; i < len(this.Quest); i++ {
		this.Quest[i].Encode(buffer)
	}
}

func (this *GetDailyInfo_Out_Quest) Encode(buffer *net.Buffer) {
	buffer.WriteUint16LE(uint16(this.Id))
	buffer.WriteUint16LE(uint16(this.FinishCount))
	buffer.WriteUint8(uint8(this.AwardState))
}

func (this *GetDailyInfo_Out) ByteSize() int {
	size := 3
	size += len(this.Quest) * 5
	return size
}

func (this *AwardDaily_In) Decode(buffer *net.Buffer) {
	this.Id = int16(buffer.ReadUint16LE())
}

func (this *AwardDaily_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(13)
	buffer.WriteUint8(3)
	buffer.WriteUint16LE(uint16(this.Id))
}

func (this *AwardDaily_In) ByteSize() int {
	size := 4
	return size
}

func (this *AwardDaily_Out) Decode(buffer *net.Buffer) {
}

func (this *AwardDaily_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(13)
	buffer.WriteUint8(3)
}

func (this *AwardDaily_Out) ByteSize() int {
	size := 2
	return size
}

func (this *NotifyDailyChange_Out) Decode(buffer *net.Buffer) {
	this.Id = int16(buffer.ReadUint16LE())
	this.FinishCount = int16(buffer.ReadUint16LE())
	this.AwardState = int8(buffer.ReadUint8())
}

func (this *NotifyDailyChange_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(13)
	buffer.WriteUint8(4)
	buffer.WriteUint16LE(uint16(this.Id))
	buffer.WriteUint16LE(uint16(this.FinishCount))
	buffer.WriteUint8(uint8(this.AwardState))
}

func (this *NotifyDailyChange_Out) ByteSize() int {
	size := 7
	return size
}

func (this *Guide_In) Decode(buffer *net.Buffer) {
	this.GuideType = GuideType(buffer.ReadUint8())
	this.Action = GuideAction(buffer.ReadUint8())
}

func (this *Guide_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(13)
	buffer.WriteUint8(5)
	buffer.WriteUint8(uint8(this.GuideType))
	buffer.WriteUint8(uint8(this.Action))
}

func (this *Guide_In) ByteSize() int {
	size := 4
	return size
}

func (this *Guide_Out) Decode(buffer *net.Buffer) {
}

func (this *Guide_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(13)
	buffer.WriteUint8(5)
}

func (this *Guide_Out) ByteSize() int {
	size := 2
	return size
}

func (this *GetExtendQuestInfoByNpcId_In) Decode(buffer *net.Buffer) {
	this.NpcId = int32(buffer.ReadUint32LE())
}

func (this *GetExtendQuestInfoByNpcId_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(13)
	buffer.WriteUint8(6)
	buffer.WriteUint32LE(uint32(this.NpcId))
}

func (this *GetExtendQuestInfoByNpcId_In) ByteSize() int {
	size := 6
	return size
}

func (this *GetExtendQuestInfoByNpcId_Out) Decode(buffer *net.Buffer) {
	this.Quest = make([]GetExtendQuestInfoByNpcId_Out_Quest, buffer.ReadUint8())
	for i := 0; i < len(this.Quest); i++ {
		this.Quest[i].Decode(buffer)
	}
}

func (this *GetExtendQuestInfoByNpcId_Out_Quest) Decode(buffer *net.Buffer) {
	this.Id = int32(buffer.ReadUint32LE())
	this.Progress = int16(buffer.ReadUint16LE())
	this.State = int8(buffer.ReadUint8())
}

func (this *GetExtendQuestInfoByNpcId_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(13)
	buffer.WriteUint8(6)
	buffer.WriteUint8(uint8(len(this.Quest)))
	for i := 0; i < len(this.Quest); i++ {
		this.Quest[i].Encode(buffer)
	}
}

func (this *GetExtendQuestInfoByNpcId_Out_Quest) Encode(buffer *net.Buffer) {
	buffer.WriteUint32LE(uint32(this.Id))
	buffer.WriteUint16LE(uint16(this.Progress))
	buffer.WriteUint8(uint8(this.State))
}

func (this *GetExtendQuestInfoByNpcId_Out) ByteSize() int {
	size := 3
	size += len(this.Quest) * 7
	return size
}

func (this *TakeExtendQuestAward_In) Decode(buffer *net.Buffer) {
	this.QuestId = int32(buffer.ReadUint32LE())
}

func (this *TakeExtendQuestAward_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(13)
	buffer.WriteUint8(7)
	buffer.WriteUint32LE(uint32(this.QuestId))
}

func (this *TakeExtendQuestAward_In) ByteSize() int {
	size := 6
	return size
}

func (this *TakeExtendQuestAward_Out) Decode(buffer *net.Buffer) {
	this.QuestId = int32(buffer.ReadUint32LE())
}

func (this *TakeExtendQuestAward_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(13)
	buffer.WriteUint8(7)
	buffer.WriteUint32LE(uint32(this.QuestId))
}

func (this *TakeExtendQuestAward_Out) ByteSize() int {
	size := 6
	return size
}

func (this *GetPannelQuestInfo_In) Decode(buffer *net.Buffer) {
}

func (this *GetPannelQuestInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(13)
	buffer.WriteUint8(8)
}

func (this *GetPannelQuestInfo_In) ByteSize() int {
	size := 2
	return size
}

func (this *GetPannelQuestInfo_Out) Decode(buffer *net.Buffer) {
	this.CurStars = int32(buffer.ReadUint32LE())
	this.Awarded = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Quest = make([]GetPannelQuestInfo_Out_Quest, buffer.ReadUint8())
	for i := 0; i < len(this.Quest); i++ {
		this.Quest[i].Decode(buffer)
	}
}

func (this *GetPannelQuestInfo_Out_Quest) Decode(buffer *net.Buffer) {
	this.QuestClass = int8(buffer.ReadUint8())
	this.Id = int32(buffer.ReadUint32LE())
	this.Progress = int16(buffer.ReadUint16LE())
	this.State = int8(buffer.ReadUint8())
}

func (this *GetPannelQuestInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(13)
	buffer.WriteUint8(8)
	buffer.WriteUint32LE(uint32(this.CurStars))
	buffer.WriteUint16LE(uint16(len(this.Awarded)))
	buffer.WriteBytes(this.Awarded)
	buffer.WriteUint8(uint8(len(this.Quest)))
	for i := 0; i < len(this.Quest); i++ {
		this.Quest[i].Encode(buffer)
	}
}

func (this *GetPannelQuestInfo_Out_Quest) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(this.QuestClass))
	buffer.WriteUint32LE(uint32(this.Id))
	buffer.WriteUint16LE(uint16(this.Progress))
	buffer.WriteUint8(uint8(this.State))
}

func (this *GetPannelQuestInfo_Out) ByteSize() int {
	size := 9
	size += len(this.Awarded)
	size += len(this.Quest) * 8
	return size
}

func (this *GiveUpAdditionQuest_In) Decode(buffer *net.Buffer) {
	this.QuestId = int32(buffer.ReadUint32LE())
}

func (this *GiveUpAdditionQuest_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(13)
	buffer.WriteUint8(10)
	buffer.WriteUint32LE(uint32(this.QuestId))
}

func (this *GiveUpAdditionQuest_In) ByteSize() int {
	size := 6
	return size
}

func (this *GiveUpAdditionQuest_Out) Decode(buffer *net.Buffer) {
}

func (this *GiveUpAdditionQuest_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(13)
	buffer.WriteUint8(10)
}

func (this *GiveUpAdditionQuest_Out) ByteSize() int {
	size := 2
	return size
}

func (this *TakeAdditionQuest_In) Decode(buffer *net.Buffer) {
	this.QuestId = int32(buffer.ReadUint32LE())
}

func (this *TakeAdditionQuest_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(13)
	buffer.WriteUint8(11)
	buffer.WriteUint32LE(uint32(this.QuestId))
}

func (this *TakeAdditionQuest_In) ByteSize() int {
	size := 6
	return size
}

func (this *TakeAdditionQuest_Out) Decode(buffer *net.Buffer) {
	this.Success = buffer.ReadUint8() == 1
}

func (this *TakeAdditionQuest_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(13)
	buffer.WriteUint8(11)
	if this.Success {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *TakeAdditionQuest_Out) ByteSize() int {
	size := 3
	return size
}

func (this *TakeAdditionQuestAward_In) Decode(buffer *net.Buffer) {
	this.QuestId = int32(buffer.ReadUint32LE())
}

func (this *TakeAdditionQuestAward_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(13)
	buffer.WriteUint8(12)
	buffer.WriteUint32LE(uint32(this.QuestId))
}

func (this *TakeAdditionQuestAward_In) ByteSize() int {
	size := 6
	return size
}

func (this *TakeAdditionQuestAward_Out) Decode(buffer *net.Buffer) {
}

func (this *TakeAdditionQuestAward_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(13)
	buffer.WriteUint8(12)
}

func (this *TakeAdditionQuestAward_Out) ByteSize() int {
	size := 2
	return size
}

func (this *GetAdditionQuest_In) Decode(buffer *net.Buffer) {
}

func (this *GetAdditionQuest_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(13)
	buffer.WriteUint8(13)
}

func (this *GetAdditionQuest_In) ByteSize() int {
	size := 2
	return size
}

func (this *GetAdditionQuest_Out) Decode(buffer *net.Buffer) {
	this.Quest = make([]GetAdditionQuest_Out_Quest, buffer.ReadUint8())
	for i := 0; i < len(this.Quest); i++ {
		this.Quest[i].Decode(buffer)
	}
}

func (this *GetAdditionQuest_Out_Quest) Decode(buffer *net.Buffer) {
	this.QuestId = int32(buffer.ReadUint32LE())
	this.Progress = int16(buffer.ReadUint16LE())
	this.State = int8(buffer.ReadUint8())
}

func (this *GetAdditionQuest_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(13)
	buffer.WriteUint8(13)
	buffer.WriteUint8(uint8(len(this.Quest)))
	for i := 0; i < len(this.Quest); i++ {
		this.Quest[i].Encode(buffer)
	}
}

func (this *GetAdditionQuest_Out_Quest) Encode(buffer *net.Buffer) {
	buffer.WriteUint32LE(uint32(this.QuestId))
	buffer.WriteUint16LE(uint16(this.Progress))
	buffer.WriteUint8(uint8(this.State))
}

func (this *GetAdditionQuest_Out) ByteSize() int {
	size := 3
	size += len(this.Quest) * 7
	return size
}

func (this *RefreshAdditionQuest_In) Decode(buffer *net.Buffer) {
	this.QuestId = int32(buffer.ReadUint32LE())
}

func (this *RefreshAdditionQuest_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(13)
	buffer.WriteUint8(14)
	buffer.WriteUint32LE(uint32(this.QuestId))
}

func (this *RefreshAdditionQuest_In) ByteSize() int {
	size := 6
	return size
}

func (this *RefreshAdditionQuest_Out) Decode(buffer *net.Buffer) {
}

func (this *RefreshAdditionQuest_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(13)
	buffer.WriteUint8(14)
}

func (this *RefreshAdditionQuest_Out) ByteSize() int {
	size := 2
	return size
}

func (this *TakeQuestStarsAwaded_In) Decode(buffer *net.Buffer) {
	this.StarsLevel = int32(buffer.ReadUint32LE())
}

func (this *TakeQuestStarsAwaded_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(13)
	buffer.WriteUint8(15)
	buffer.WriteUint32LE(uint32(this.StarsLevel))
}

func (this *TakeQuestStarsAwaded_In) ByteSize() int {
	size := 6
	return size
}

func (this *TakeQuestStarsAwaded_Out) Decode(buffer *net.Buffer) {
	this.Result = buffer.ReadUint8() == 1
}

func (this *TakeQuestStarsAwaded_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(13)
	buffer.WriteUint8(15)
	if this.Result {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *TakeQuestStarsAwaded_Out) ByteSize() int {
	size := 3
	return size
}
