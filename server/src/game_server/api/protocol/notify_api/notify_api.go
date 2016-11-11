package notify_api

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
	PlayerKeyChanged(*net.Session, *PlayerKeyChanged_In)
}

type OutHandler interface {
	PlayerKeyChanged(*net.Session, *PlayerKeyChanged_Out)
	MissionLevelLockChanged(*net.Session, *MissionLevelLockChanged_Out)
	RoleExpChange(*net.Session, *RoleExpChange_Out)
	PhysicalChange(*net.Session, *PhysicalChange_Out)
	MoneyChange(*net.Session, *MoneyChange_Out)
	SkillAdd(*net.Session, *SkillAdd_Out)
	ItemChange(*net.Session, *ItemChange_Out)
	RoleBattleStatusChange(*net.Session, *RoleBattleStatusChange_Out)
	NewMail(*net.Session, *NewMail_Out)
	HeartChange(*net.Session, *HeartChange_Out)
	QuestChange(*net.Session, *QuestChange_Out)
	TownLockChange(*net.Session, *TownLockChange_Out)
	Chat(*net.Session, *Chat_Out)
	FuncKeyChange(*net.Session, *FuncKeyChange_Out)
	ItemRecastStateRebuild(*net.Session, *ItemRecastStateRebuild_Out)
	SendAnnouncement(*net.Session, *SendAnnouncement_Out)
	VipLevelChange(*net.Session, *VipLevelChange_Out)
	NotifyNewBuddy(*net.Session, *NotifyNewBuddy_Out)
	HardLevelLockChanged(*net.Session, *HardLevelLockChanged_Out)
	SendSwordSoulDrawNumChange(*net.Session, *SendSwordSoulDrawNumChange_Out)
	SendHaveNewGhost(*net.Session, *SendHaveNewGhost_Out)
	SendHeartRecoverTime(*net.Session, *SendHeartRecoverTime_Out)
	SendGlobalMail(*net.Session, *SendGlobalMail_Out)
	SendPhysicalRecoverTime(*net.Session, *SendPhysicalRecoverTime_Out)
	SendFashionChange(*net.Session, *SendFashionChange_Out)
	TransError(*net.Session, *TransError_Out)
	SendEventCenterChange(*net.Session, *SendEventCenterChange_Out)
	MeditationState(*net.Session, *MeditationState_Out)
	DeleteAnnouncement(*net.Session, *DeleteAnnouncement_Out)
	SendHaveNewPet(*net.Session, *SendHaveNewPet_Out)
	SendLogout(*net.Session, *SendLogout_Out)
	FameChange(*net.Session, *FameChange_Out)
	NotifyMonthCardOpen(*net.Session, *NotifyMonthCardOpen_Out)
	NotifyMonthCardRenewal(*net.Session, *NotifyMonthCardRenewal_Out)
	NotifyNewTotem(*net.Session, *NotifyNewTotem_Out)
	NotifyRuneChange(*net.Session, *NotifyRuneChange_Out)
	TaoyuanItemChange(*net.Session, *TaoyuanItemChange_Out)
	TaoyuanMessageRefresh(*net.Session, *TaoyuanMessageRefresh_Out)
	TaoyuanQuestCanFinish(*net.Session, *TaoyuanQuestCanFinish_Out)
	TaoyuanExpRefresh(*net.Session, *TaoyuanExpRefresh_Out)
}

func DecodeIn(message []byte) Request {
	var actionId = message[0]
	var buffer = net.NewBuffer(message[1:])
	switch actionId {
	case 0:
		request := new(PlayerKeyChanged_In)
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
		request := new(PlayerKeyChanged_Out)
		request.Decode(buffer)
		return request
	case 1:
		request := new(MissionLevelLockChanged_Out)
		request.Decode(buffer)
		return request
	case 2:
		request := new(RoleExpChange_Out)
		request.Decode(buffer)
		return request
	case 3:
		request := new(PhysicalChange_Out)
		request.Decode(buffer)
		return request
	case 4:
		request := new(MoneyChange_Out)
		request.Decode(buffer)
		return request
	case 5:
		request := new(SkillAdd_Out)
		request.Decode(buffer)
		return request
	case 6:
		request := new(ItemChange_Out)
		request.Decode(buffer)
		return request
	case 7:
		request := new(RoleBattleStatusChange_Out)
		request.Decode(buffer)
		return request
	case 8:
		request := new(NewMail_Out)
		request.Decode(buffer)
		return request
	case 9:
		request := new(HeartChange_Out)
		request.Decode(buffer)
		return request
	case 10:
		request := new(QuestChange_Out)
		request.Decode(buffer)
		return request
	case 11:
		request := new(TownLockChange_Out)
		request.Decode(buffer)
		return request
	case 12:
		request := new(Chat_Out)
		request.Decode(buffer)
		return request
	case 13:
		request := new(FuncKeyChange_Out)
		request.Decode(buffer)
		return request
	case 14:
		request := new(ItemRecastStateRebuild_Out)
		request.Decode(buffer)
		return request
	case 15:
		request := new(SendAnnouncement_Out)
		request.Decode(buffer)
		return request
	case 16:
		request := new(VipLevelChange_Out)
		request.Decode(buffer)
		return request
	case 17:
		request := new(NotifyNewBuddy_Out)
		request.Decode(buffer)
		return request
	case 18:
		request := new(HardLevelLockChanged_Out)
		request.Decode(buffer)
		return request
	case 19:
		request := new(SendSwordSoulDrawNumChange_Out)
		request.Decode(buffer)
		return request
	case 21:
		request := new(SendHaveNewGhost_Out)
		request.Decode(buffer)
		return request
	case 22:
		request := new(SendHeartRecoverTime_Out)
		request.Decode(buffer)
		return request
	case 23:
		request := new(SendGlobalMail_Out)
		request.Decode(buffer)
		return request
	case 24:
		request := new(SendPhysicalRecoverTime_Out)
		request.Decode(buffer)
		return request
	case 25:
		request := new(SendFashionChange_Out)
		request.Decode(buffer)
		return request
	case 26:
		request := new(TransError_Out)
		request.Decode(buffer)
		return request
	case 27:
		request := new(SendEventCenterChange_Out)
		request.Decode(buffer)
		return request
	case 29:
		request := new(MeditationState_Out)
		request.Decode(buffer)
		return request
	case 31:
		request := new(DeleteAnnouncement_Out)
		request.Decode(buffer)
		return request
	case 32:
		request := new(SendHaveNewPet_Out)
		request.Decode(buffer)
		return request
	case 33:
		request := new(SendLogout_Out)
		request.Decode(buffer)
		return request
	case 34:
		request := new(FameChange_Out)
		request.Decode(buffer)
		return request
	case 36:
		request := new(NotifyMonthCardOpen_Out)
		request.Decode(buffer)
		return request
	case 37:
		request := new(NotifyMonthCardRenewal_Out)
		request.Decode(buffer)
		return request
	case 38:
		request := new(NotifyNewTotem_Out)
		request.Decode(buffer)
		return request
	case 39:
		request := new(NotifyRuneChange_Out)
		request.Decode(buffer)
		return request
	case 40:
		request := new(TaoyuanItemChange_Out)
		request.Decode(buffer)
		return request
	case 41:
		request := new(TaoyuanMessageRefresh_Out)
		request.Decode(buffer)
		return request
	case 42:
		request := new(TaoyuanQuestCanFinish_Out)
		request.Decode(buffer)
		return request
	case 43:
		request := new(TaoyuanExpRefresh_Out)
		request.Decode(buffer)
		return request
	}
	_ = buffer
	panic("unsupported response")
}

type BuffMode int8

const (
	BUFF_MODE_POWER                    BuffMode = 0
	BUFF_MODE_SPEED                    BuffMode = 1
	BUFF_MODE_ATTACK                   BuffMode = 2
	BUFF_MODE_DEFEND                   BuffMode = 3
	BUFF_MODE_HEALTH                   BuffMode = 4
	BUFF_MODE_DIZZINESS                BuffMode = 5
	BUFF_MODE_POISONING                BuffMode = 6
	BUFF_MODE_CLEAN_BAD                BuffMode = 7
	BUFF_MODE_CLEAN_GOOD               BuffMode = 8
	BUFF_MODE_REDUCE_HURT              BuffMode = 9
	BUFF_MODE_RANDOM                   BuffMode = 10
	BUFF_MODE_BLOCK                    BuffMode = 11
	BUFF_MODE_BLOCK_LEVEL              BuffMode = 12
	BUFF_MODE_DODGE_LEVEL              BuffMode = 13
	BUFF_MODE_CRITIAL_LEVEL            BuffMode = 14
	BUFF_MODE_HIT_LEVEL                BuffMode = 15
	BUFF_MODE_HURT_ADD                 BuffMode = 16
	BUFF_MODE_MAX_HEALTH               BuffMode = 17
	BUFF_MODE_KEEPER_REDUCE_HURT       BuffMode = 18
	BUFF_MODE_ATTRACT_FIRE             BuffMode = 19
	BUFF_MODE_DESTROY_LEVEL            BuffMode = 20
	BUFF_MODE_TENACITY_LEVEL           BuffMode = 21
	BUFF_MODE_SUNDER                   BuffMode = 22
	BUFF_MODE_SLEEP                    BuffMode = 23
	BUFF_MODE_DISABLE_SKILL            BuffMode = 24
	BUFF_MODE_DIRECT_REDUCE_HURT       BuffMode = 25
	BUFF_MODE_ABSORB_HURT              BuffMode = 26
	BUFF_MODE_GHOST_POWER              BuffMode = 27
	BUFF_MODE_PET_LIVE_ROUND           BuffMode = 28
	BUFF_MODE_BUDDY_SKILL              BuffMode = 29
	BUFF_MODE_CLEAR_ABSORB_HURT        BuffMode = 30
	BUFF_MODE_SLEEP_LEVEL              BuffMode = 31
	BUFF_MODE_DIZZINESS_LEVEL          BuffMode = 32
	BUFF_MODE_RANDOM_LEVEL             BuffMode = 33
	BUFF_MODE_DISABLE_SKILL_LEVEL      BuffMode = 34
	BUFF_MODE_BUFF_POISONING_LEVEL     BuffMode = 35
	BUFF_MODE_BUFF_RECOVER_BUDDY_SKILL BuffMode = 36
	BUFF_MODE_BUFF_MAKE_POWER_FULL     BuffMode = 37
	BUFF_MODE_BUFF_DOGE                BuffMode = 38
	BUFF_MODE_BUFF_HIT                 BuffMode = 39
	BUFF_MODE_BUFF_CRITIAL             BuffMode = 40
	BUFF_MODE_BUFF_TENACITY            BuffMode = 41
	BUFF_MODE_BUFF_TAKE_SUNSER         BuffMode = 42
	BUFF_MODE_BUFF_DEFEND_PERSENT      BuffMode = 43
	BUFF_MODE_BUFF_SUNDER_STATE        BuffMode = 44
)

type BufferInfo struct {
	Mode        BuffMode `json:"mode"`
	Keep        int8     `json:"keep"`
	Value       int32    `json:"value"`
	SkillId     int16    `json:"skill_id"`
	MaxOverride int8     `json:"max_override"`
	OverrideNum int8     `json:"override_num"`
}

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

type ChestType int8

const (
	CHEST_TYPE_COIN_FREE  ChestType = 0
	CHEST_TYPE_INGOT_FREE ChestType = 1
)

func (this *BufferInfo) Decode(buffer *net.Buffer) {
	this.Mode = BuffMode(buffer.ReadUint8())
	this.Keep = int8(buffer.ReadUint8())
	this.Value = int32(buffer.ReadUint32LE())
	this.SkillId = int16(buffer.ReadUint16LE())
	this.MaxOverride = int8(buffer.ReadUint8())
	this.OverrideNum = int8(buffer.ReadUint8())
}

func (this *BufferInfo) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(this.Mode))
	buffer.WriteUint8(uint8(this.Keep))
	buffer.WriteUint32LE(uint32(this.Value))
	buffer.WriteUint16LE(uint16(this.SkillId))
	buffer.WriteUint8(uint8(this.MaxOverride))
	buffer.WriteUint8(uint8(this.OverrideNum))
}

func (this *BufferInfo) ByteSize() int {
	size := 10
	return size
}

type PlayerKeyChanged_In struct {
}

func (this *PlayerKeyChanged_In) Process(session *net.Session) {
	g_InHandler.PlayerKeyChanged(session, this)
}

func (this *PlayerKeyChanged_In) TypeName() string {
	return "notify.player_key_changed.in"
}

func (this *PlayerKeyChanged_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 8, 0
}

type PlayerKeyChanged_Out struct {
	Key      int32 `json:"key"`
	MaxOrder int8  `json:"max_order"`
}

func (this *PlayerKeyChanged_Out) Process(session *net.Session) {
	g_OutHandler.PlayerKeyChanged(session, this)
}

func (this *PlayerKeyChanged_Out) TypeName() string {
	return "notify.player_key_changed.out"
}

func (this *PlayerKeyChanged_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 8, 0
}

func (this *PlayerKeyChanged_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type MissionLevelLockChanged_Out struct {
	MaxLock   int32 `json:"max_lock"`
	AwardLock int32 `json:"award_lock"`
}

func (this *MissionLevelLockChanged_Out) Process(session *net.Session) {
	g_OutHandler.MissionLevelLockChanged(session, this)
}

func (this *MissionLevelLockChanged_Out) TypeName() string {
	return "notify.mission_level_lock_changed.out"
}

func (this *MissionLevelLockChanged_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 8, 1
}

func (this *MissionLevelLockChanged_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type RoleExpChange_Out struct {
	RoleId int8  `json:"role_id"`
	AddExp int64 `json:"add_exp"`
	Exp    int64 `json:"exp"`
	Level  int16 `json:"level"`
}

func (this *RoleExpChange_Out) Process(session *net.Session) {
	g_OutHandler.RoleExpChange(session, this)
}

func (this *RoleExpChange_Out) TypeName() string {
	return "notify.role_exp_change.out"
}

func (this *RoleExpChange_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 8, 2
}

func (this *RoleExpChange_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type PhysicalChange_Out struct {
	Value int16 `json:"value"`
}

func (this *PhysicalChange_Out) Process(session *net.Session) {
	g_OutHandler.PhysicalChange(session, this)
}

func (this *PhysicalChange_Out) TypeName() string {
	return "notify.physical_change.out"
}

func (this *PhysicalChange_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 8, 3
}

func (this *PhysicalChange_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type MoneyChange_Out struct {
	Moneytype int8  `json:"moneytype"`
	Value     int64 `json:"value"`
	Timestamp int64 `json:"timestamp"`
}

func (this *MoneyChange_Out) Process(session *net.Session) {
	g_OutHandler.MoneyChange(session, this)
}

func (this *MoneyChange_Out) TypeName() string {
	return "notify.money_change.out"
}

func (this *MoneyChange_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 8, 4
}

func (this *MoneyChange_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type SkillAdd_Out struct {
	RoleId  int8  `json:"role_id"`
	SkillId int16 `json:"skill_id"`
}

func (this *SkillAdd_Out) Process(session *net.Session) {
	g_OutHandler.SkillAdd(session, this)
}

func (this *SkillAdd_Out) TypeName() string {
	return "notify.skill_add.out"
}

func (this *SkillAdd_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 8, 5
}

func (this *SkillAdd_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ItemChange_Out struct {
	Items []ItemChange_Out_Items `json:"items"`
}

type ItemChange_Out_Items struct {
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

func (this *ItemChange_Out) Process(session *net.Session) {
	g_OutHandler.ItemChange(session, this)
}

func (this *ItemChange_Out) TypeName() string {
	return "notify.item_change.out"
}

func (this *ItemChange_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 8, 6
}

func (this *ItemChange_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type RoleBattleStatusChange_Out struct {
	Roles []RoleBattleStatusChange_Out_Roles `json:"roles"`
}

type RoleBattleStatusChange_Out_Roles struct {
	RoleId int8                                     `json:"role_id"`
	Health int32                                    `json:"health"`
	Buffs  []RoleBattleStatusChange_Out_Roles_Buffs `json:"buffs"`
}

type RoleBattleStatusChange_Out_Roles_Buffs struct {
	Buffer BufferInfo `json:"buffer"`
}

func (this *RoleBattleStatusChange_Out) Process(session *net.Session) {
	g_OutHandler.RoleBattleStatusChange(session, this)
}

func (this *RoleBattleStatusChange_Out) TypeName() string {
	return "notify.role_battle_status_change.out"
}

func (this *RoleBattleStatusChange_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 8, 7
}

func (this *RoleBattleStatusChange_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type NewMail_Out struct {
}

func (this *NewMail_Out) Process(session *net.Session) {
	g_OutHandler.NewMail(session, this)
}

func (this *NewMail_Out) TypeName() string {
	return "notify.new_mail.out"
}

func (this *NewMail_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 8, 8
}

func (this *NewMail_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type HeartChange_Out struct {
	Value int16 `json:"value"`
}

func (this *HeartChange_Out) Process(session *net.Session) {
	g_OutHandler.HeartChange(session, this)
}

func (this *HeartChange_Out) TypeName() string {
	return "notify.heart_change.out"
}

func (this *HeartChange_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 8, 9
}

func (this *HeartChange_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type QuestChange_Out struct {
	QuestId int16 `json:"quest_id"`
	State   int8  `json:"state"`
}

func (this *QuestChange_Out) Process(session *net.Session) {
	g_OutHandler.QuestChange(session, this)
}

func (this *QuestChange_Out) TypeName() string {
	return "notify.quest_change.out"
}

func (this *QuestChange_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 8, 10
}

func (this *QuestChange_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type TownLockChange_Out struct {
	Lock int32 `json:"lock"`
}

func (this *TownLockChange_Out) Process(session *net.Session) {
	g_OutHandler.TownLockChange(session, this)
}

func (this *TownLockChange_Out) TypeName() string {
	return "notify.townLock_change.out"
}

func (this *TownLockChange_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 8, 11
}

func (this *TownLockChange_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type Chat_Out struct {
	Pid      int64  `json:"pid"`
	RoleId   int8   `json:"role_id"`
	Nickname []byte `json:"nickname"`
	Level    int16  `json:"level"`
	FightNum int32  `json:"fight_num"`
	Message  []byte `json:"message"`
}

func (this *Chat_Out) Process(session *net.Session) {
	g_OutHandler.Chat(session, this)
}

func (this *Chat_Out) TypeName() string {
	return "notify.chat.out"
}

func (this *Chat_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 8, 12
}

func (this *Chat_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type FuncKeyChange_Out struct {
	FuncKey int16 `json:"func_key"`
}

func (this *FuncKeyChange_Out) Process(session *net.Session) {
	g_OutHandler.FuncKeyChange(session, this)
}

func (this *FuncKeyChange_Out) TypeName() string {
	return "notify.func_key_change.out"
}

func (this *FuncKeyChange_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 8, 13
}

func (this *FuncKeyChange_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ItemRecastStateRebuild_Out struct {
	Id           int64                              `json:"id"`
	SelectedAttr Attribute                          `json:"selected_attr"`
	Attrs        []ItemRecastStateRebuild_Out_Attrs `json:"attrs"`
}

type ItemRecastStateRebuild_Out_Attrs struct {
	Attr  Attribute `json:"attr"`
	Value int32     `json:"value"`
}

func (this *ItemRecastStateRebuild_Out) Process(session *net.Session) {
	g_OutHandler.ItemRecastStateRebuild(session, this)
}

func (this *ItemRecastStateRebuild_Out) TypeName() string {
	return "notify.item_recast_state_rebuild.out"
}

func (this *ItemRecastStateRebuild_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 8, 14
}

func (this *ItemRecastStateRebuild_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type SendAnnouncement_Out struct {
	Id          int64  `json:"id"`
	TplId       int32  `json:"tpl_id"`
	ExpireTime  int64  `json:"expire_time"`
	Parameters  []byte `json:"parameters"`
	Content     []byte `json:"content"`
	SpacingTime int32  `json:"spacing_time"`
}

func (this *SendAnnouncement_Out) Process(session *net.Session) {
	g_OutHandler.SendAnnouncement(session, this)
}

func (this *SendAnnouncement_Out) TypeName() string {
	return "notify.send_announcement.out"
}

func (this *SendAnnouncement_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 8, 15
}

func (this *SendAnnouncement_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type VipLevelChange_Out struct {
	Level int16 `json:"level"`
}

func (this *VipLevelChange_Out) Process(session *net.Session) {
	g_OutHandler.VipLevelChange(session, this)
}

func (this *VipLevelChange_Out) TypeName() string {
	return "notify.vip_level_change.out"
}

func (this *VipLevelChange_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 8, 16
}

func (this *VipLevelChange_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type NotifyNewBuddy_Out struct {
	RoleId    int8  `json:"role_id"`
	RoleLevel int16 `json:"role_level"`
}

func (this *NotifyNewBuddy_Out) Process(session *net.Session) {
	g_OutHandler.NotifyNewBuddy(session, this)
}

func (this *NotifyNewBuddy_Out) TypeName() string {
	return "notify.notify_new_buddy.out"
}

func (this *NotifyNewBuddy_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 8, 17
}

func (this *NotifyNewBuddy_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type HardLevelLockChanged_Out struct {
	Lock int32 `json:"lock"`
}

func (this *HardLevelLockChanged_Out) Process(session *net.Session) {
	g_OutHandler.HardLevelLockChanged(session, this)
}

func (this *HardLevelLockChanged_Out) TypeName() string {
	return "notify.hard_level_lock_changed.out"
}

func (this *HardLevelLockChanged_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 8, 18
}

func (this *HardLevelLockChanged_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type SendSwordSoulDrawNumChange_Out struct {
	Num    int16 `json:"num"`
	CdTime int64 `json:"cd_time"`
}

func (this *SendSwordSoulDrawNumChange_Out) Process(session *net.Session) {
	g_OutHandler.SendSwordSoulDrawNumChange(session, this)
}

func (this *SendSwordSoulDrawNumChange_Out) TypeName() string {
	return "notify.send_sword_soul_draw_num_change.out"
}

func (this *SendSwordSoulDrawNumChange_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 8, 19
}

func (this *SendSwordSoulDrawNumChange_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type SendHaveNewGhost_Out struct {
	PlayerGhostId int64 `json:"player_ghost_id"`
}

func (this *SendHaveNewGhost_Out) Process(session *net.Session) {
	g_OutHandler.SendHaveNewGhost(session, this)
}

func (this *SendHaveNewGhost_Out) TypeName() string {
	return "notify.send_have_new_ghost.out"
}

func (this *SendHaveNewGhost_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 8, 21
}

func (this *SendHaveNewGhost_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type SendHeartRecoverTime_Out struct {
	Timestamp int64 `json:"timestamp"`
}

func (this *SendHeartRecoverTime_Out) Process(session *net.Session) {
	g_OutHandler.SendHeartRecoverTime(session, this)
}

func (this *SendHeartRecoverTime_Out) TypeName() string {
	return "notify.send_heart_recover_time.out"
}

func (this *SendHeartRecoverTime_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 8, 22
}

func (this *SendHeartRecoverTime_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type SendGlobalMail_Out struct {
}

func (this *SendGlobalMail_Out) Process(session *net.Session) {
	g_OutHandler.SendGlobalMail(session, this)
}

func (this *SendGlobalMail_Out) TypeName() string {
	return "notify.send_global_mail.out"
}

func (this *SendGlobalMail_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 8, 23
}

func (this *SendGlobalMail_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type SendPhysicalRecoverTime_Out struct {
	Timestamp int64 `json:"timestamp"`
}

func (this *SendPhysicalRecoverTime_Out) Process(session *net.Session) {
	g_OutHandler.SendPhysicalRecoverTime(session, this)
}

func (this *SendPhysicalRecoverTime_Out) TypeName() string {
	return "notify.send_physical_recover_time.out"
}

func (this *SendPhysicalRecoverTime_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 8, 24
}

func (this *SendPhysicalRecoverTime_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type SendFashionChange_Out struct {
	FashionId  int16 `json:"fashion_id"`
	ExpireTime int64 `json:"expire_time"`
}

func (this *SendFashionChange_Out) Process(session *net.Session) {
	g_OutHandler.SendFashionChange(session, this)
}

func (this *SendFashionChange_Out) TypeName() string {
	return "notify.send_fashion_change.out"
}

func (this *SendFashionChange_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 8, 25
}

func (this *SendFashionChange_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type TransError_Out struct {
}

func (this *TransError_Out) Process(session *net.Session) {
	g_OutHandler.TransError(session, this)
}

func (this *TransError_Out) TypeName() string {
	return "notify.trans_error.out"
}

func (this *TransError_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 8, 26
}

func (this *TransError_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type SendEventCenterChange_Out struct {
}

func (this *SendEventCenterChange_Out) Process(session *net.Session) {
	g_OutHandler.SendEventCenterChange(session, this)
}

func (this *SendEventCenterChange_Out) TypeName() string {
	return "notify.send_event_center_change.out"
}

func (this *SendEventCenterChange_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 8, 27
}

func (this *SendEventCenterChange_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type MeditationState_Out struct {
}

func (this *MeditationState_Out) Process(session *net.Session) {
	g_OutHandler.MeditationState(session, this)
}

func (this *MeditationState_Out) TypeName() string {
	return "notify.meditation_state.out"
}

func (this *MeditationState_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 8, 29
}

func (this *MeditationState_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type DeleteAnnouncement_Out struct {
	Id int64 `json:"id"`
}

func (this *DeleteAnnouncement_Out) Process(session *net.Session) {
	g_OutHandler.DeleteAnnouncement(session, this)
}

func (this *DeleteAnnouncement_Out) TypeName() string {
	return "notify.delete_announcement.out"
}

func (this *DeleteAnnouncement_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 8, 31
}

func (this *DeleteAnnouncement_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type SendHaveNewPet_Out struct {
	PetId int32 `json:"pet_id"`
}

func (this *SendHaveNewPet_Out) Process(session *net.Session) {
	g_OutHandler.SendHaveNewPet(session, this)
}

func (this *SendHaveNewPet_Out) TypeName() string {
	return "notify.send_have_new_pet.out"
}

func (this *SendHaveNewPet_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 8, 32
}

func (this *SendHaveNewPet_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type SendLogout_Out struct {
}

func (this *SendLogout_Out) Process(session *net.Session) {
	g_OutHandler.SendLogout(session, this)
}

func (this *SendLogout_Out) TypeName() string {
	return "notify.send_logout.out"
}

func (this *SendLogout_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 8, 33
}

func (this *SendLogout_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type FameChange_Out struct {
	Fame int32 `json:"fame"`
}

func (this *FameChange_Out) Process(session *net.Session) {
	g_OutHandler.FameChange(session, this)
}

func (this *FameChange_Out) TypeName() string {
	return "notify.fame_change.out"
}

func (this *FameChange_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 8, 34
}

func (this *FameChange_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type NotifyMonthCardOpen_Out struct {
}

func (this *NotifyMonthCardOpen_Out) Process(session *net.Session) {
	g_OutHandler.NotifyMonthCardOpen(session, this)
}

func (this *NotifyMonthCardOpen_Out) TypeName() string {
	return "notify.notify_month_card_open.out"
}

func (this *NotifyMonthCardOpen_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 8, 36
}

func (this *NotifyMonthCardOpen_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type NotifyMonthCardRenewal_Out struct {
}

func (this *NotifyMonthCardRenewal_Out) Process(session *net.Session) {
	g_OutHandler.NotifyMonthCardRenewal(session, this)
}

func (this *NotifyMonthCardRenewal_Out) TypeName() string {
	return "notify.notify_month_card_renewal.out"
}

func (this *NotifyMonthCardRenewal_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 8, 37
}

func (this *NotifyMonthCardRenewal_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type NotifyNewTotem_Out struct {
	Id      int64 `json:"id"`
	TotemId int16 `json:"totem_id"`
	Skill   int16 `json:"skill"`
}

func (this *NotifyNewTotem_Out) Process(session *net.Session) {
	g_OutHandler.NotifyNewTotem(session, this)
}

func (this *NotifyNewTotem_Out) TypeName() string {
	return "notify.notify_new_totem.out"
}

func (this *NotifyNewTotem_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 8, 38
}

func (this *NotifyNewTotem_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type NotifyRuneChange_Out struct {
	RockRuneNum int32 `json:"rock_rune_num"`
	JadeRuneNum int32 `json:"jade_rune_num"`
}

func (this *NotifyRuneChange_Out) Process(session *net.Session) {
	g_OutHandler.NotifyRuneChange(session, this)
}

func (this *NotifyRuneChange_Out) TypeName() string {
	return "notify.notify_rune_change.out"
}

func (this *NotifyRuneChange_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 8, 39
}

func (this *NotifyRuneChange_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type TaoyuanItemChange_Out struct {
	Items []TaoyuanItemChange_Out_Items `json:"items"`
}

type TaoyuanItemChange_Out_Items struct {
	Id     int64 `json:"id"`
	ItemId int16 `json:"item_id"`
	Num    int16 `json:"num"`
}

func (this *TaoyuanItemChange_Out) Process(session *net.Session) {
	g_OutHandler.TaoyuanItemChange(session, this)
}

func (this *TaoyuanItemChange_Out) TypeName() string {
	return "notify.taoyuan_item_change.out"
}

func (this *TaoyuanItemChange_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 8, 40
}

func (this *TaoyuanItemChange_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type TaoyuanMessageRefresh_Out struct {
}

func (this *TaoyuanMessageRefresh_Out) Process(session *net.Session) {
	g_OutHandler.TaoyuanMessageRefresh(session, this)
}

func (this *TaoyuanMessageRefresh_Out) TypeName() string {
	return "notify.taoyuan_message_refresh.out"
}

func (this *TaoyuanMessageRefresh_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 8, 41
}

func (this *TaoyuanMessageRefresh_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type TaoyuanQuestCanFinish_Out struct {
}

func (this *TaoyuanQuestCanFinish_Out) Process(session *net.Session) {
	g_OutHandler.TaoyuanQuestCanFinish(session, this)
}

func (this *TaoyuanQuestCanFinish_Out) TypeName() string {
	return "notify.taoyuan_quest_can_finish.out"
}

func (this *TaoyuanQuestCanFinish_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 8, 42
}

func (this *TaoyuanQuestCanFinish_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type TaoyuanExpRefresh_Out struct {
	LevelChange bool  `json:"level_change"`
	Exp         int64 `json:"exp"`
	Level       int16 `json:"level"`
}

func (this *TaoyuanExpRefresh_Out) Process(session *net.Session) {
	g_OutHandler.TaoyuanExpRefresh(session, this)
}

func (this *TaoyuanExpRefresh_Out) TypeName() string {
	return "notify.taoyuan_exp_refresh.out"
}

func (this *TaoyuanExpRefresh_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 8, 43
}

func (this *TaoyuanExpRefresh_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

func (this *PlayerKeyChanged_In) Decode(buffer *net.Buffer) {
}

func (this *PlayerKeyChanged_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(8)
	buffer.WriteUint8(0)
}

func (this *PlayerKeyChanged_In) ByteSize() int {
	size := 2
	return size
}

func (this *PlayerKeyChanged_Out) Decode(buffer *net.Buffer) {
	this.Key = int32(buffer.ReadUint32LE())
	this.MaxOrder = int8(buffer.ReadUint8())
}

func (this *PlayerKeyChanged_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(8)
	buffer.WriteUint8(0)
	buffer.WriteUint32LE(uint32(this.Key))
	buffer.WriteUint8(uint8(this.MaxOrder))
}

func (this *PlayerKeyChanged_Out) ByteSize() int {
	size := 7
	return size
}

func (this *MissionLevelLockChanged_Out) Decode(buffer *net.Buffer) {
	this.MaxLock = int32(buffer.ReadUint32LE())
	this.AwardLock = int32(buffer.ReadUint32LE())
}

func (this *MissionLevelLockChanged_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(8)
	buffer.WriteUint8(1)
	buffer.WriteUint32LE(uint32(this.MaxLock))
	buffer.WriteUint32LE(uint32(this.AwardLock))
}

func (this *MissionLevelLockChanged_Out) ByteSize() int {
	size := 10
	return size
}

func (this *RoleExpChange_Out) Decode(buffer *net.Buffer) {
	this.RoleId = int8(buffer.ReadUint8())
	this.AddExp = int64(buffer.ReadUint64LE())
	this.Exp = int64(buffer.ReadUint64LE())
	this.Level = int16(buffer.ReadUint16LE())
}

func (this *RoleExpChange_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(8)
	buffer.WriteUint8(2)
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint64LE(uint64(this.AddExp))
	buffer.WriteUint64LE(uint64(this.Exp))
	buffer.WriteUint16LE(uint16(this.Level))
}

func (this *RoleExpChange_Out) ByteSize() int {
	size := 21
	return size
}

func (this *PhysicalChange_Out) Decode(buffer *net.Buffer) {
	this.Value = int16(buffer.ReadUint16LE())
}

func (this *PhysicalChange_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(8)
	buffer.WriteUint8(3)
	buffer.WriteUint16LE(uint16(this.Value))
}

func (this *PhysicalChange_Out) ByteSize() int {
	size := 4
	return size
}

func (this *MoneyChange_Out) Decode(buffer *net.Buffer) {
	this.Moneytype = int8(buffer.ReadUint8())
	this.Value = int64(buffer.ReadUint64LE())
	this.Timestamp = int64(buffer.ReadUint64LE())
}

func (this *MoneyChange_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(8)
	buffer.WriteUint8(4)
	buffer.WriteUint8(uint8(this.Moneytype))
	buffer.WriteUint64LE(uint64(this.Value))
	buffer.WriteUint64LE(uint64(this.Timestamp))
}

func (this *MoneyChange_Out) ByteSize() int {
	size := 19
	return size
}

func (this *SkillAdd_Out) Decode(buffer *net.Buffer) {
	this.RoleId = int8(buffer.ReadUint8())
	this.SkillId = int16(buffer.ReadUint16LE())
}

func (this *SkillAdd_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(8)
	buffer.WriteUint8(5)
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint16LE(uint16(this.SkillId))
}

func (this *SkillAdd_Out) ByteSize() int {
	size := 5
	return size
}

func (this *ItemChange_Out) Decode(buffer *net.Buffer) {
	this.Items = make([]ItemChange_Out_Items, buffer.ReadUint8())
	for i := 0; i < len(this.Items); i++ {
		this.Items[i].Decode(buffer)
	}
}

func (this *ItemChange_Out_Items) Decode(buffer *net.Buffer) {
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

func (this *ItemChange_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(8)
	buffer.WriteUint8(6)
	buffer.WriteUint8(uint8(len(this.Items)))
	for i := 0; i < len(this.Items); i++ {
		this.Items[i].Encode(buffer)
	}
}

func (this *ItemChange_Out_Items) Encode(buffer *net.Buffer) {
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

func (this *ItemChange_Out) ByteSize() int {
	size := 3
	size += len(this.Items) * 59
	return size
}

func (this *RoleBattleStatusChange_Out) Decode(buffer *net.Buffer) {
	this.Roles = make([]RoleBattleStatusChange_Out_Roles, buffer.ReadUint8())
	for i := 0; i < len(this.Roles); i++ {
		this.Roles[i].Decode(buffer)
	}
}

func (this *RoleBattleStatusChange_Out_Roles) Decode(buffer *net.Buffer) {
	this.RoleId = int8(buffer.ReadUint8())
	this.Health = int32(buffer.ReadUint32LE())
	this.Buffs = make([]RoleBattleStatusChange_Out_Roles_Buffs, buffer.ReadUint8())
	for i := 0; i < len(this.Buffs); i++ {
		this.Buffs[i].Decode(buffer)
	}
}

func (this *RoleBattleStatusChange_Out_Roles_Buffs) Decode(buffer *net.Buffer) {
	this.Buffer.Decode(buffer)
}

func (this *RoleBattleStatusChange_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(8)
	buffer.WriteUint8(7)
	buffer.WriteUint8(uint8(len(this.Roles)))
	for i := 0; i < len(this.Roles); i++ {
		this.Roles[i].Encode(buffer)
	}
}

func (this *RoleBattleStatusChange_Out_Roles) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint32LE(uint32(this.Health))
	buffer.WriteUint8(uint8(len(this.Buffs)))
	for i := 0; i < len(this.Buffs); i++ {
		this.Buffs[i].Encode(buffer)
	}
}

func (this *RoleBattleStatusChange_Out_Roles_Buffs) Encode(buffer *net.Buffer) {
	this.Buffer.Encode(buffer)
}

func (this *RoleBattleStatusChange_Out) ByteSize() int {
	size := 3
	for i := 0; i < len(this.Roles); i++ {
		size += this.Roles[i].ByteSize()
	}
	return size
}

func (this *RoleBattleStatusChange_Out_Roles) ByteSize() int {
	size := 6
	for i := 0; i < len(this.Buffs); i++ {
		size += this.Buffs[i].ByteSize()
	}
	return size
}

func (this *RoleBattleStatusChange_Out_Roles_Buffs) ByteSize() int {
	size := 0
	size += this.Buffer.ByteSize()
	return size
}

func (this *NewMail_Out) Decode(buffer *net.Buffer) {
}

func (this *NewMail_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(8)
	buffer.WriteUint8(8)
}

func (this *NewMail_Out) ByteSize() int {
	size := 2
	return size
}

func (this *HeartChange_Out) Decode(buffer *net.Buffer) {
	this.Value = int16(buffer.ReadUint16LE())
}

func (this *HeartChange_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(8)
	buffer.WriteUint8(9)
	buffer.WriteUint16LE(uint16(this.Value))
}

func (this *HeartChange_Out) ByteSize() int {
	size := 4
	return size
}

func (this *QuestChange_Out) Decode(buffer *net.Buffer) {
	this.QuestId = int16(buffer.ReadUint16LE())
	this.State = int8(buffer.ReadUint8())
}

func (this *QuestChange_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(8)
	buffer.WriteUint8(10)
	buffer.WriteUint16LE(uint16(this.QuestId))
	buffer.WriteUint8(uint8(this.State))
}

func (this *QuestChange_Out) ByteSize() int {
	size := 5
	return size
}

func (this *TownLockChange_Out) Decode(buffer *net.Buffer) {
	this.Lock = int32(buffer.ReadUint32LE())
}

func (this *TownLockChange_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(8)
	buffer.WriteUint8(11)
	buffer.WriteUint32LE(uint32(this.Lock))
}

func (this *TownLockChange_Out) ByteSize() int {
	size := 6
	return size
}

func (this *Chat_Out) Decode(buffer *net.Buffer) {
	this.Pid = int64(buffer.ReadUint64LE())
	this.RoleId = int8(buffer.ReadUint8())
	this.Nickname = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Level = int16(buffer.ReadUint16LE())
	this.FightNum = int32(buffer.ReadUint32LE())
	this.Message = buffer.ReadBytes(int(buffer.ReadUint16LE()))
}

func (this *Chat_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(8)
	buffer.WriteUint8(12)
	buffer.WriteUint64LE(uint64(this.Pid))
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint16LE(uint16(len(this.Nickname)))
	buffer.WriteBytes(this.Nickname)
	buffer.WriteUint16LE(uint16(this.Level))
	buffer.WriteUint32LE(uint32(this.FightNum))
	buffer.WriteUint16LE(uint16(len(this.Message)))
	buffer.WriteBytes(this.Message)
}

func (this *Chat_Out) ByteSize() int {
	size := 21
	size += len(this.Nickname)
	size += len(this.Message)
	return size
}

func (this *FuncKeyChange_Out) Decode(buffer *net.Buffer) {
	this.FuncKey = int16(buffer.ReadUint16LE())
}

func (this *FuncKeyChange_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(8)
	buffer.WriteUint8(13)
	buffer.WriteUint16LE(uint16(this.FuncKey))
}

func (this *FuncKeyChange_Out) ByteSize() int {
	size := 4
	return size
}

func (this *ItemRecastStateRebuild_Out) Decode(buffer *net.Buffer) {
	this.Id = int64(buffer.ReadUint64LE())
	this.SelectedAttr = Attribute(buffer.ReadUint8())
	this.Attrs = make([]ItemRecastStateRebuild_Out_Attrs, buffer.ReadUint8())
	for i := 0; i < len(this.Attrs); i++ {
		this.Attrs[i].Decode(buffer)
	}
}

func (this *ItemRecastStateRebuild_Out_Attrs) Decode(buffer *net.Buffer) {
	this.Attr = Attribute(buffer.ReadUint8())
	this.Value = int32(buffer.ReadUint32LE())
}

func (this *ItemRecastStateRebuild_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(8)
	buffer.WriteUint8(14)
	buffer.WriteUint64LE(uint64(this.Id))
	buffer.WriteUint8(uint8(this.SelectedAttr))
	buffer.WriteUint8(uint8(len(this.Attrs)))
	for i := 0; i < len(this.Attrs); i++ {
		this.Attrs[i].Encode(buffer)
	}
}

func (this *ItemRecastStateRebuild_Out_Attrs) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(this.Attr))
	buffer.WriteUint32LE(uint32(this.Value))
}

func (this *ItemRecastStateRebuild_Out) ByteSize() int {
	size := 12
	size += len(this.Attrs) * 5
	return size
}

func (this *SendAnnouncement_Out) Decode(buffer *net.Buffer) {
	this.Id = int64(buffer.ReadUint64LE())
	this.TplId = int32(buffer.ReadUint32LE())
	this.ExpireTime = int64(buffer.ReadUint64LE())
	this.Parameters = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Content = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.SpacingTime = int32(buffer.ReadUint32LE())
}

func (this *SendAnnouncement_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(8)
	buffer.WriteUint8(15)
	buffer.WriteUint64LE(uint64(this.Id))
	buffer.WriteUint32LE(uint32(this.TplId))
	buffer.WriteUint64LE(uint64(this.ExpireTime))
	buffer.WriteUint16LE(uint16(len(this.Parameters)))
	buffer.WriteBytes(this.Parameters)
	buffer.WriteUint16LE(uint16(len(this.Content)))
	buffer.WriteBytes(this.Content)
	buffer.WriteUint32LE(uint32(this.SpacingTime))
}

func (this *SendAnnouncement_Out) ByteSize() int {
	size := 30
	size += len(this.Parameters)
	size += len(this.Content)
	return size
}

func (this *VipLevelChange_Out) Decode(buffer *net.Buffer) {
	this.Level = int16(buffer.ReadUint16LE())
}

func (this *VipLevelChange_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(8)
	buffer.WriteUint8(16)
	buffer.WriteUint16LE(uint16(this.Level))
}

func (this *VipLevelChange_Out) ByteSize() int {
	size := 4
	return size
}

func (this *NotifyNewBuddy_Out) Decode(buffer *net.Buffer) {
	this.RoleId = int8(buffer.ReadUint8())
	this.RoleLevel = int16(buffer.ReadUint16LE())
}

func (this *NotifyNewBuddy_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(8)
	buffer.WriteUint8(17)
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint16LE(uint16(this.RoleLevel))
}

func (this *NotifyNewBuddy_Out) ByteSize() int {
	size := 5
	return size
}

func (this *HardLevelLockChanged_Out) Decode(buffer *net.Buffer) {
	this.Lock = int32(buffer.ReadUint32LE())
}

func (this *HardLevelLockChanged_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(8)
	buffer.WriteUint8(18)
	buffer.WriteUint32LE(uint32(this.Lock))
}

func (this *HardLevelLockChanged_Out) ByteSize() int {
	size := 6
	return size
}

func (this *SendSwordSoulDrawNumChange_Out) Decode(buffer *net.Buffer) {
	this.Num = int16(buffer.ReadUint16LE())
	this.CdTime = int64(buffer.ReadUint64LE())
}

func (this *SendSwordSoulDrawNumChange_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(8)
	buffer.WriteUint8(19)
	buffer.WriteUint16LE(uint16(this.Num))
	buffer.WriteUint64LE(uint64(this.CdTime))
}

func (this *SendSwordSoulDrawNumChange_Out) ByteSize() int {
	size := 12
	return size
}

func (this *SendHaveNewGhost_Out) Decode(buffer *net.Buffer) {
	this.PlayerGhostId = int64(buffer.ReadUint64LE())
}

func (this *SendHaveNewGhost_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(8)
	buffer.WriteUint8(21)
	buffer.WriteUint64LE(uint64(this.PlayerGhostId))
}

func (this *SendHaveNewGhost_Out) ByteSize() int {
	size := 10
	return size
}

func (this *SendHeartRecoverTime_Out) Decode(buffer *net.Buffer) {
	this.Timestamp = int64(buffer.ReadUint64LE())
}

func (this *SendHeartRecoverTime_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(8)
	buffer.WriteUint8(22)
	buffer.WriteUint64LE(uint64(this.Timestamp))
}

func (this *SendHeartRecoverTime_Out) ByteSize() int {
	size := 10
	return size
}

func (this *SendGlobalMail_Out) Decode(buffer *net.Buffer) {
}

func (this *SendGlobalMail_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(8)
	buffer.WriteUint8(23)
}

func (this *SendGlobalMail_Out) ByteSize() int {
	size := 2
	return size
}

func (this *SendPhysicalRecoverTime_Out) Decode(buffer *net.Buffer) {
	this.Timestamp = int64(buffer.ReadUint64LE())
}

func (this *SendPhysicalRecoverTime_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(8)
	buffer.WriteUint8(24)
	buffer.WriteUint64LE(uint64(this.Timestamp))
}

func (this *SendPhysicalRecoverTime_Out) ByteSize() int {
	size := 10
	return size
}

func (this *SendFashionChange_Out) Decode(buffer *net.Buffer) {
	this.FashionId = int16(buffer.ReadUint16LE())
	this.ExpireTime = int64(buffer.ReadUint64LE())
}

func (this *SendFashionChange_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(8)
	buffer.WriteUint8(25)
	buffer.WriteUint16LE(uint16(this.FashionId))
	buffer.WriteUint64LE(uint64(this.ExpireTime))
}

func (this *SendFashionChange_Out) ByteSize() int {
	size := 12
	return size
}

func (this *TransError_Out) Decode(buffer *net.Buffer) {
}

func (this *TransError_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(8)
	buffer.WriteUint8(26)
}

func (this *TransError_Out) ByteSize() int {
	size := 2
	return size
}

func (this *SendEventCenterChange_Out) Decode(buffer *net.Buffer) {
}

func (this *SendEventCenterChange_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(8)
	buffer.WriteUint8(27)
}

func (this *SendEventCenterChange_Out) ByteSize() int {
	size := 2
	return size
}

func (this *MeditationState_Out) Decode(buffer *net.Buffer) {
}

func (this *MeditationState_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(8)
	buffer.WriteUint8(29)
}

func (this *MeditationState_Out) ByteSize() int {
	size := 2
	return size
}

func (this *DeleteAnnouncement_Out) Decode(buffer *net.Buffer) {
	this.Id = int64(buffer.ReadUint64LE())
}

func (this *DeleteAnnouncement_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(8)
	buffer.WriteUint8(31)
	buffer.WriteUint64LE(uint64(this.Id))
}

func (this *DeleteAnnouncement_Out) ByteSize() int {
	size := 10
	return size
}

func (this *SendHaveNewPet_Out) Decode(buffer *net.Buffer) {
	this.PetId = int32(buffer.ReadUint32LE())
}

func (this *SendHaveNewPet_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(8)
	buffer.WriteUint8(32)
	buffer.WriteUint32LE(uint32(this.PetId))
}

func (this *SendHaveNewPet_Out) ByteSize() int {
	size := 6
	return size
}

func (this *SendLogout_Out) Decode(buffer *net.Buffer) {
}

func (this *SendLogout_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(8)
	buffer.WriteUint8(33)
}

func (this *SendLogout_Out) ByteSize() int {
	size := 2
	return size
}

func (this *FameChange_Out) Decode(buffer *net.Buffer) {
	this.Fame = int32(buffer.ReadUint32LE())
}

func (this *FameChange_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(8)
	buffer.WriteUint8(34)
	buffer.WriteUint32LE(uint32(this.Fame))
}

func (this *FameChange_Out) ByteSize() int {
	size := 6
	return size
}

func (this *NotifyMonthCardOpen_Out) Decode(buffer *net.Buffer) {
}

func (this *NotifyMonthCardOpen_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(8)
	buffer.WriteUint8(36)
}

func (this *NotifyMonthCardOpen_Out) ByteSize() int {
	size := 2
	return size
}

func (this *NotifyMonthCardRenewal_Out) Decode(buffer *net.Buffer) {
}

func (this *NotifyMonthCardRenewal_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(8)
	buffer.WriteUint8(37)
}

func (this *NotifyMonthCardRenewal_Out) ByteSize() int {
	size := 2
	return size
}

func (this *NotifyNewTotem_Out) Decode(buffer *net.Buffer) {
	this.Id = int64(buffer.ReadUint64LE())
	this.TotemId = int16(buffer.ReadUint16LE())
	this.Skill = int16(buffer.ReadUint16LE())
}

func (this *NotifyNewTotem_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(8)
	buffer.WriteUint8(38)
	buffer.WriteUint64LE(uint64(this.Id))
	buffer.WriteUint16LE(uint16(this.TotemId))
	buffer.WriteUint16LE(uint16(this.Skill))
}

func (this *NotifyNewTotem_Out) ByteSize() int {
	size := 14
	return size
}

func (this *NotifyRuneChange_Out) Decode(buffer *net.Buffer) {
	this.RockRuneNum = int32(buffer.ReadUint32LE())
	this.JadeRuneNum = int32(buffer.ReadUint32LE())
}

func (this *NotifyRuneChange_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(8)
	buffer.WriteUint8(39)
	buffer.WriteUint32LE(uint32(this.RockRuneNum))
	buffer.WriteUint32LE(uint32(this.JadeRuneNum))
}

func (this *NotifyRuneChange_Out) ByteSize() int {
	size := 10
	return size
}

func (this *TaoyuanItemChange_Out) Decode(buffer *net.Buffer) {
	this.Items = make([]TaoyuanItemChange_Out_Items, buffer.ReadUint8())
	for i := 0; i < len(this.Items); i++ {
		this.Items[i].Decode(buffer)
	}
}

func (this *TaoyuanItemChange_Out_Items) Decode(buffer *net.Buffer) {
	this.Id = int64(buffer.ReadUint64LE())
	this.ItemId = int16(buffer.ReadUint16LE())
	this.Num = int16(buffer.ReadUint16LE())
}

func (this *TaoyuanItemChange_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(8)
	buffer.WriteUint8(40)
	buffer.WriteUint8(uint8(len(this.Items)))
	for i := 0; i < len(this.Items); i++ {
		this.Items[i].Encode(buffer)
	}
}

func (this *TaoyuanItemChange_Out_Items) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.Id))
	buffer.WriteUint16LE(uint16(this.ItemId))
	buffer.WriteUint16LE(uint16(this.Num))
}

func (this *TaoyuanItemChange_Out) ByteSize() int {
	size := 3
	size += len(this.Items) * 12
	return size
}

func (this *TaoyuanMessageRefresh_Out) Decode(buffer *net.Buffer) {
}

func (this *TaoyuanMessageRefresh_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(8)
	buffer.WriteUint8(41)
}

func (this *TaoyuanMessageRefresh_Out) ByteSize() int {
	size := 2
	return size
}

func (this *TaoyuanQuestCanFinish_Out) Decode(buffer *net.Buffer) {
}

func (this *TaoyuanQuestCanFinish_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(8)
	buffer.WriteUint8(42)
}

func (this *TaoyuanQuestCanFinish_Out) ByteSize() int {
	size := 2
	return size
}

func (this *TaoyuanExpRefresh_Out) Decode(buffer *net.Buffer) {
	this.LevelChange = buffer.ReadUint8() == 1
	this.Exp = int64(buffer.ReadUint64LE())
	this.Level = int16(buffer.ReadUint16LE())
}

func (this *TaoyuanExpRefresh_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(8)
	buffer.WriteUint8(43)
	if this.LevelChange {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
	buffer.WriteUint64LE(uint64(this.Exp))
	buffer.WriteUint16LE(uint16(this.Level))
}

func (this *TaoyuanExpRefresh_Out) ByteSize() int {
	size := 13
	return size
}
