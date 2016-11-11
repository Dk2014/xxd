package role_api

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
	GetAllRoles(*net.Session, *GetAllRoles_In)
	GetRoleInfo(*net.Session, *GetRoleInfo_In)
	GetPlayerInfo(*net.Session, *GetPlayerInfo_In)
	GetFightNum(*net.Session, *GetFightNum_In)
	GetPlayerInfoWithOpenid(*net.Session, *GetPlayerInfoWithOpenid_In)
	LevelupRoleFriendship(*net.Session, *LevelupRoleFriendship_In)
	RecruitBuddy(*net.Session, *RecruitBuddy_In)
	GetRoleFightNum(*net.Session, *GetRoleFightNum_In)
	ChangeRoleStatus(*net.Session, *ChangeRoleStatus_In)
	GetInnRoleList(*net.Session, *GetInnRoleList_In)
	BuddyCoop(*net.Session, *BuddyCoop_In)
}

type OutHandler interface {
	GetAllRoles(*net.Session, *GetAllRoles_Out)
	GetRoleInfo(*net.Session, *GetRoleInfo_Out)
	GetPlayerInfo(*net.Session, *GetPlayerInfo_Out)
	GetFightNum(*net.Session, *GetFightNum_Out)
	GetPlayerInfoWithOpenid(*net.Session, *GetPlayerInfoWithOpenid_Out)
	LevelupRoleFriendship(*net.Session, *LevelupRoleFriendship_Out)
	RecruitBuddy(*net.Session, *RecruitBuddy_Out)
	GetRoleFightNum(*net.Session, *GetRoleFightNum_Out)
	ChangeRoleStatus(*net.Session, *ChangeRoleStatus_Out)
	GetInnRoleList(*net.Session, *GetInnRoleList_Out)
	BuddyCoop(*net.Session, *BuddyCoop_Out)
}

func DecodeIn(message []byte) Request {
	var actionId = message[0]
	var buffer = net.NewBuffer(message[1:])
	switch actionId {
	case 0:
		request := new(GetAllRoles_In)
		request.Decode(buffer)
		return request
	case 1:
		request := new(GetRoleInfo_In)
		request.Decode(buffer)
		return request
	case 2:
		request := new(GetPlayerInfo_In)
		request.Decode(buffer)
		return request
	case 3:
		request := new(GetFightNum_In)
		request.Decode(buffer)
		return request
	case 4:
		request := new(GetPlayerInfoWithOpenid_In)
		request.Decode(buffer)
		return request
	case 5:
		request := new(LevelupRoleFriendship_In)
		request.Decode(buffer)
		return request
	case 6:
		request := new(RecruitBuddy_In)
		request.Decode(buffer)
		return request
	case 7:
		request := new(GetRoleFightNum_In)
		request.Decode(buffer)
		return request
	case 8:
		request := new(ChangeRoleStatus_In)
		request.Decode(buffer)
		return request
	case 9:
		request := new(GetInnRoleList_In)
		request.Decode(buffer)
		return request
	case 10:
		request := new(BuddyCoop_In)
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
		request := new(GetAllRoles_Out)
		request.Decode(buffer)
		return request
	case 1:
		request := new(GetRoleInfo_Out)
		request.Decode(buffer)
		return request
	case 2:
		request := new(GetPlayerInfo_Out)
		request.Decode(buffer)
		return request
	case 3:
		request := new(GetFightNum_Out)
		request.Decode(buffer)
		return request
	case 4:
		request := new(GetPlayerInfoWithOpenid_Out)
		request.Decode(buffer)
		return request
	case 5:
		request := new(LevelupRoleFriendship_Out)
		request.Decode(buffer)
		return request
	case 6:
		request := new(RecruitBuddy_Out)
		request.Decode(buffer)
		return request
	case 7:
		request := new(GetRoleFightNum_Out)
		request.Decode(buffer)
		return request
	case 8:
		request := new(ChangeRoleStatus_Out)
		request.Decode(buffer)
		return request
	case 9:
		request := new(GetInnRoleList_Out)
		request.Decode(buffer)
		return request
	case 10:
		request := new(BuddyCoop_Out)
		request.Decode(buffer)
		return request
	}
	_ = buffer
	panic("unsupported response")
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

type FightnumType int8

const (
	FIGHTNUM_TYPE_ALL           FightnumType = 0
	FIGHTNUM_TYPE_ROLE_LEVEL    FightnumType = 1
	FIGHTNUM_TYPE_SKILL         FightnumType = 2
	FIGHTNUM_TYPE_EQUIP         FightnumType = 3
	FIGHTNUM_TYPE_GHOST         FightnumType = 4
	FIGHTNUM_TYPE_REALM         FightnumType = 5
	FIGHTNUM_TYPE_FASHION       FightnumType = 8
	FIGHTNUM_TYPE_FRIENDSHIP    FightnumType = 9
	FIGHTNUM_TYPE_TEAMSHIP      FightnumType = 10
	FIGHTNUM_TYPE_CLIQUE_KONGFU FightnumType = 11
)

type PlayerInfo struct {
	Openid              []byte             `json:"openid"`
	Pid                 int64              `json:"pid"`
	Name                []byte             `json:"name"`
	BestSegment         int16              `json:"best_segment"`
	BestOrder           int8               `json:"best_order"`
	BestRecordTimestamp int64              `json:"best_record_timestamp"`
	FashionId           int16              `json:"fashion_id"`
	Roles               []PlayerInfo_Roles `json:"roles"`
}

type PlayerInfo_Roles struct {
	RoleId            int8                      `json:"role_id"`
	Level             int16                     `json:"level"`
	FriendshipLevel   int16                     `json:"friendship_level"`
	FightNum          int32                     `json:"fight_num"`
	IsDeploy          bool                      `json:"is_deploy"`
	Status            int8                      `json:"status"`
	CoopId            int16                     `json:"coop_id"`
	Attack            int32                     `json:"attack"`
	Defence           int32                     `json:"defence"`
	Health            int32                     `json:"health"`
	Speed             int32                     `json:"speed"`
	Cultivation       int32                     `json:"cultivation"`
	Sunder            int32                     `json:"sunder"`
	HitLevel          int32                     `json:"hit_level"`
	CriticalLevel     int32                     `json:"critical_level"`
	SleepLevel        int32                     `json:"sleep_level"`
	DizzinessLevel    int32                     `json:"dizziness_level"`
	RandomLevel       int32                     `json:"random_level"`
	DisableSkillLevel int32                     `json:"disable_skill_level"`
	PoisoningLevel    int32                     `json:"poisoning_level"`
	BlockLevel        int32                     `json:"block_level"`
	DestroyLevel      int32                     `json:"destroy_level"`
	CriticalHurtLevel int32                     `json:"critical_hurt_level"`
	TenacityLevel     int32                     `json:"tenacity_level"`
	DodgeLevel        int32                     `json:"dodge_level"`
	Equips            []PlayerInfo_Roles_Equips `json:"equips"`
}

type PlayerInfo_Roles_Equips struct {
	Pos           int8      `json:"pos"`
	ItemId        int16     `json:"item_id"`
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

func (this *PlayerInfo) Decode(buffer *net.Buffer) {
	this.Openid = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Pid = int64(buffer.ReadUint64LE())
	this.Name = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.BestSegment = int16(buffer.ReadUint16LE())
	this.BestOrder = int8(buffer.ReadUint8())
	this.BestRecordTimestamp = int64(buffer.ReadUint64LE())
	this.FashionId = int16(buffer.ReadUint16LE())
	this.Roles = make([]PlayerInfo_Roles, buffer.ReadUint8())
	for i := 0; i < len(this.Roles); i++ {
		this.Roles[i].Decode(buffer)
	}
}

func (this *PlayerInfo_Roles) Decode(buffer *net.Buffer) {
	this.RoleId = int8(buffer.ReadUint8())
	this.Level = int16(buffer.ReadUint16LE())
	this.FriendshipLevel = int16(buffer.ReadUint16LE())
	this.FightNum = int32(buffer.ReadUint32LE())
	this.IsDeploy = buffer.ReadUint8() == 1
	this.Status = int8(buffer.ReadUint8())
	this.CoopId = int16(buffer.ReadUint16LE())
	this.Attack = int32(buffer.ReadUint32LE())
	this.Defence = int32(buffer.ReadUint32LE())
	this.Health = int32(buffer.ReadUint32LE())
	this.Speed = int32(buffer.ReadUint32LE())
	this.Cultivation = int32(buffer.ReadUint32LE())
	this.Sunder = int32(buffer.ReadUint32LE())
	this.HitLevel = int32(buffer.ReadUint32LE())
	this.CriticalLevel = int32(buffer.ReadUint32LE())
	this.SleepLevel = int32(buffer.ReadUint32LE())
	this.DizzinessLevel = int32(buffer.ReadUint32LE())
	this.RandomLevel = int32(buffer.ReadUint32LE())
	this.DisableSkillLevel = int32(buffer.ReadUint32LE())
	this.PoisoningLevel = int32(buffer.ReadUint32LE())
	this.BlockLevel = int32(buffer.ReadUint32LE())
	this.DestroyLevel = int32(buffer.ReadUint32LE())
	this.CriticalHurtLevel = int32(buffer.ReadUint32LE())
	this.TenacityLevel = int32(buffer.ReadUint32LE())
	this.DodgeLevel = int32(buffer.ReadUint32LE())
	this.Equips = make([]PlayerInfo_Roles_Equips, buffer.ReadUint8())
	for i := 0; i < len(this.Equips); i++ {
		this.Equips[i].Decode(buffer)
	}
}

func (this *PlayerInfo_Roles_Equips) Decode(buffer *net.Buffer) {
	this.Pos = int8(buffer.ReadUint8())
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
	this.RecastAttr = Attribute(buffer.ReadUint8())
}

func (this *PlayerInfo) Encode(buffer *net.Buffer) {
	buffer.WriteUint16LE(uint16(len(this.Openid)))
	buffer.WriteBytes(this.Openid)
	buffer.WriteUint64LE(uint64(this.Pid))
	buffer.WriteUint16LE(uint16(len(this.Name)))
	buffer.WriteBytes(this.Name)
	buffer.WriteUint16LE(uint16(this.BestSegment))
	buffer.WriteUint8(uint8(this.BestOrder))
	buffer.WriteUint64LE(uint64(this.BestRecordTimestamp))
	buffer.WriteUint16LE(uint16(this.FashionId))
	buffer.WriteUint8(uint8(len(this.Roles)))
	for i := 0; i < len(this.Roles); i++ {
		this.Roles[i].Encode(buffer)
	}
}

func (this *PlayerInfo_Roles) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint16LE(uint16(this.Level))
	buffer.WriteUint16LE(uint16(this.FriendshipLevel))
	buffer.WriteUint32LE(uint32(this.FightNum))
	if this.IsDeploy {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
	buffer.WriteUint8(uint8(this.Status))
	buffer.WriteUint16LE(uint16(this.CoopId))
	buffer.WriteUint32LE(uint32(this.Attack))
	buffer.WriteUint32LE(uint32(this.Defence))
	buffer.WriteUint32LE(uint32(this.Health))
	buffer.WriteUint32LE(uint32(this.Speed))
	buffer.WriteUint32LE(uint32(this.Cultivation))
	buffer.WriteUint32LE(uint32(this.Sunder))
	buffer.WriteUint32LE(uint32(this.HitLevel))
	buffer.WriteUint32LE(uint32(this.CriticalLevel))
	buffer.WriteUint32LE(uint32(this.SleepLevel))
	buffer.WriteUint32LE(uint32(this.DizzinessLevel))
	buffer.WriteUint32LE(uint32(this.RandomLevel))
	buffer.WriteUint32LE(uint32(this.DisableSkillLevel))
	buffer.WriteUint32LE(uint32(this.PoisoningLevel))
	buffer.WriteUint32LE(uint32(this.BlockLevel))
	buffer.WriteUint32LE(uint32(this.DestroyLevel))
	buffer.WriteUint32LE(uint32(this.CriticalHurtLevel))
	buffer.WriteUint32LE(uint32(this.TenacityLevel))
	buffer.WriteUint32LE(uint32(this.DodgeLevel))
	buffer.WriteUint8(uint8(len(this.Equips)))
	for i := 0; i < len(this.Equips); i++ {
		this.Equips[i].Encode(buffer)
	}
}

func (this *PlayerInfo_Roles_Equips) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(this.Pos))
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
	buffer.WriteUint8(uint8(this.RecastAttr))
}

func (this *PlayerInfo) ByteSize() int {
	size := 26
	size += len(this.Openid)
	size += len(this.Name)
	for i := 0; i < len(this.Roles); i++ {
		size += this.Roles[i].ByteSize()
	}
	return size
}

func (this *PlayerInfo_Roles) ByteSize() int {
	size := 86
	size += len(this.Equips) * 50
	return size
}

type GetAllRoles_In struct {
}

func (this *GetAllRoles_In) Process(session *net.Session) {
	g_InHandler.GetAllRoles(session, this)
}

func (this *GetAllRoles_In) TypeName() string {
	return "role.get_all_roles.in"
}

func (this *GetAllRoles_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 3, 0
}

type GetAllRoles_Out struct {
	Roles []GetAllRoles_Out_Roles `json:"roles"`
}

type GetAllRoles_Out_Roles struct {
	RoleId          int8  `json:"role_id"`
	Level           int16 `json:"level"`
	Exp             int64 `json:"exp"`
	FriendshipLevel int16 `json:"friendship_level"`
	InForm          bool  `json:"in_form"`
	Status          int8  `json:"status"`
	CoopId          int16 `json:"coop_id"`
}

func (this *GetAllRoles_Out) Process(session *net.Session) {
	g_OutHandler.GetAllRoles(session, this)
}

func (this *GetAllRoles_Out) TypeName() string {
	return "role.get_all_roles.out"
}

func (this *GetAllRoles_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 3, 0
}

func (this *GetAllRoles_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetRoleInfo_In struct {
	RoleId int8 `json:"role_id"`
}

func (this *GetRoleInfo_In) Process(session *net.Session) {
	g_InHandler.GetRoleInfo(session, this)
}

func (this *GetRoleInfo_In) TypeName() string {
	return "role.get_role_info.in"
}

func (this *GetRoleInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 3, 1
}

type GetRoleInfo_Out struct {
	RoleId            int8  `json:"role_id"`
	Level             int16 `json:"level"`
	Exp               int64 `json:"exp"`
	FriendshipLevel   int16 `json:"friendship_level"`
	FightNum          int32 `json:"fight_num"`
	Status            int8  `json:"status"`
	CoopId            int16 `json:"coop_id"`
	Attack            int32 `json:"attack"`
	Defence           int32 `json:"defence"`
	Health            int32 `json:"health"`
	Speed             int32 `json:"speed"`
	Cultivation       int32 `json:"cultivation"`
	Sunder            int32 `json:"sunder"`
	GhostPower        int32 `json:"ghost_power"`
	HitLevel          int32 `json:"hit_level"`
	CriticalLevel     int32 `json:"critical_level"`
	SleepLevel        int32 `json:"sleep_level"`
	DizzinessLevel    int32 `json:"dizziness_level"`
	RandomLevel       int32 `json:"random_level"`
	DisableSkillLevel int32 `json:"disable_skill_level"`
	PoisoningLevel    int32 `json:"poisoning_level"`
	BlockLevel        int32 `json:"block_level"`
	DestroyLevel      int32 `json:"destroy_level"`
	CriticalHurtLevel int32 `json:"critical_hurt_level"`
	TenacityLevel     int32 `json:"tenacity_level"`
	DodgeLevel        int32 `json:"dodge_level"`
}

func (this *GetRoleInfo_Out) Process(session *net.Session) {
	g_OutHandler.GetRoleInfo(session, this)
}

func (this *GetRoleInfo_Out) TypeName() string {
	return "role.get_role_info.out"
}

func (this *GetRoleInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 3, 1
}

func (this *GetRoleInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetPlayerInfo_In struct {
	Pid int64 `json:"pid"`
}

func (this *GetPlayerInfo_In) Process(session *net.Session) {
	g_InHandler.GetPlayerInfo(session, this)
}

func (this *GetPlayerInfo_In) TypeName() string {
	return "role.get_player_info.in"
}

func (this *GetPlayerInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 3, 2
}

type GetPlayerInfo_Out struct {
	Player PlayerInfo `json:"player"`
}

func (this *GetPlayerInfo_Out) Process(session *net.Session) {
	g_OutHandler.GetPlayerInfo(session, this)
}

func (this *GetPlayerInfo_Out) TypeName() string {
	return "role.get_player_info.out"
}

func (this *GetPlayerInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 3, 2
}

func (this *GetPlayerInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetFightNum_In struct {
	FightType FightnumType `json:"fight_type"`
}

func (this *GetFightNum_In) Process(session *net.Session) {
	g_InHandler.GetFightNum(session, this)
}

func (this *GetFightNum_In) TypeName() string {
	return "role.get_fight_num.in"
}

func (this *GetFightNum_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 3, 3
}

type GetFightNum_Out struct {
	FightNums []GetFightNum_Out_FightNums `json:"fight_nums"`
}

type GetFightNum_Out_FightNums struct {
	FightType int16 `json:"fight_type"`
	FightNum  int32 `json:"fight_num"`
}

func (this *GetFightNum_Out) Process(session *net.Session) {
	g_OutHandler.GetFightNum(session, this)
}

func (this *GetFightNum_Out) TypeName() string {
	return "role.get_fight_num.out"
}

func (this *GetFightNum_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 3, 3
}

func (this *GetFightNum_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetPlayerInfoWithOpenid_In struct {
	Openid       []byte `json:"openid"`
	GameServerId int32  `json:"game_server_id"`
}

func (this *GetPlayerInfoWithOpenid_In) Process(session *net.Session) {
	g_InHandler.GetPlayerInfoWithOpenid(session, this)
}

func (this *GetPlayerInfoWithOpenid_In) TypeName() string {
	return "role.get_player_info_with_openid.in"
}

func (this *GetPlayerInfoWithOpenid_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 3, 4
}

type GetPlayerInfoWithOpenid_Out struct {
	Player PlayerInfo `json:"player"`
}

func (this *GetPlayerInfoWithOpenid_Out) Process(session *net.Session) {
	g_OutHandler.GetPlayerInfoWithOpenid(session, this)
}

func (this *GetPlayerInfoWithOpenid_Out) TypeName() string {
	return "role.get_player_info_with_openid.out"
}

func (this *GetPlayerInfoWithOpenid_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 3, 4
}

func (this *GetPlayerInfoWithOpenid_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type LevelupRoleFriendship_In struct {
	RoleId int8 `json:"role_id"`
}

func (this *LevelupRoleFriendship_In) Process(session *net.Session) {
	g_InHandler.LevelupRoleFriendship(session, this)
}

func (this *LevelupRoleFriendship_In) TypeName() string {
	return "role.levelup_role_friendship.in"
}

func (this *LevelupRoleFriendship_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 3, 5
}

type LevelupRoleFriendship_Out struct {
}

func (this *LevelupRoleFriendship_Out) Process(session *net.Session) {
	g_OutHandler.LevelupRoleFriendship(session, this)
}

func (this *LevelupRoleFriendship_Out) TypeName() string {
	return "role.levelup_role_friendship.out"
}

func (this *LevelupRoleFriendship_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 3, 5
}

func (this *LevelupRoleFriendship_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type RecruitBuddy_In struct {
	RoleId int8 `json:"role_id"`
}

func (this *RecruitBuddy_In) Process(session *net.Session) {
	g_InHandler.RecruitBuddy(session, this)
}

func (this *RecruitBuddy_In) TypeName() string {
	return "role.recruit_buddy.in"
}

func (this *RecruitBuddy_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 3, 6
}

type RecruitBuddy_Out struct {
	Result int8 `json:"result"`
}

func (this *RecruitBuddy_Out) Process(session *net.Session) {
	g_OutHandler.RecruitBuddy(session, this)
}

func (this *RecruitBuddy_Out) TypeName() string {
	return "role.recruit_buddy.out"
}

func (this *RecruitBuddy_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 3, 6
}

func (this *RecruitBuddy_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetRoleFightNum_In struct {
	RoleId int8 `json:"role_id"`
}

func (this *GetRoleFightNum_In) Process(session *net.Session) {
	g_InHandler.GetRoleFightNum(session, this)
}

func (this *GetRoleFightNum_In) TypeName() string {
	return "role.get_role_fight_num.in"
}

func (this *GetRoleFightNum_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 3, 7
}

type GetRoleFightNum_Out struct {
	FightNums []GetRoleFightNum_Out_FightNums `json:"fight_nums"`
}

type GetRoleFightNum_Out_FightNums struct {
	RoleId   int8  `json:"role_id"`
	FightNum int32 `json:"fight_num"`
}

func (this *GetRoleFightNum_Out) Process(session *net.Session) {
	g_OutHandler.GetRoleFightNum(session, this)
}

func (this *GetRoleFightNum_Out) TypeName() string {
	return "role.get_role_fight_num.out"
}

func (this *GetRoleFightNum_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 3, 7
}

func (this *GetRoleFightNum_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ChangeRoleStatus_In struct {
	RoleId int8 `json:"role_id"`
	Status int8 `json:"status"`
}

func (this *ChangeRoleStatus_In) Process(session *net.Session) {
	g_InHandler.ChangeRoleStatus(session, this)
}

func (this *ChangeRoleStatus_In) TypeName() string {
	return "role.change_role_status.in"
}

func (this *ChangeRoleStatus_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 3, 8
}

type ChangeRoleStatus_Out struct {
	Result bool `json:"result"`
	RoleId int8 `json:"role_id"`
	Status int8 `json:"status"`
}

func (this *ChangeRoleStatus_Out) Process(session *net.Session) {
	g_OutHandler.ChangeRoleStatus(session, this)
}

func (this *ChangeRoleStatus_Out) TypeName() string {
	return "role.change_role_status.out"
}

func (this *ChangeRoleStatus_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 3, 8
}

func (this *ChangeRoleStatus_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetInnRoleList_In struct {
}

func (this *GetInnRoleList_In) Process(session *net.Session) {
	g_InHandler.GetInnRoleList(session, this)
}

func (this *GetInnRoleList_In) TypeName() string {
	return "role.get_inn_role_list.in"
}

func (this *GetInnRoleList_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 3, 9
}

type GetInnRoleList_Out struct {
	RoleList []GetInnRoleList_Out_RoleList `json:"role_list"`
}

type GetInnRoleList_Out_RoleList struct {
	RoleId          int8  `json:"role_id"`
	FriendshipLevel int16 `json:"friendship_level"`
	FightNum        int32 `json:"fight_num"`
	RoleLevel       int16 `json:"role_level"`
	Operate         int8  `json:"operate"`
}

func (this *GetInnRoleList_Out) Process(session *net.Session) {
	g_OutHandler.GetInnRoleList(session, this)
}

func (this *GetInnRoleList_Out) TypeName() string {
	return "role.get_inn_role_list.out"
}

func (this *GetInnRoleList_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 3, 9
}

func (this *GetInnRoleList_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type BuddyCoop_In struct {
	CoopId int16 `json:"coop_id"`
}

func (this *BuddyCoop_In) Process(session *net.Session) {
	g_InHandler.BuddyCoop(session, this)
}

func (this *BuddyCoop_In) TypeName() string {
	return "role.buddy_coop.in"
}

func (this *BuddyCoop_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 3, 10
}

type BuddyCoop_Out struct {
}

func (this *BuddyCoop_Out) Process(session *net.Session) {
	g_OutHandler.BuddyCoop(session, this)
}

func (this *BuddyCoop_Out) TypeName() string {
	return "role.buddy_coop.out"
}

func (this *BuddyCoop_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 3, 10
}

func (this *BuddyCoop_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

func (this *GetAllRoles_In) Decode(buffer *net.Buffer) {
}

func (this *GetAllRoles_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(3)
	buffer.WriteUint8(0)
}

func (this *GetAllRoles_In) ByteSize() int {
	size := 2
	return size
}

func (this *GetAllRoles_Out) Decode(buffer *net.Buffer) {
	this.Roles = make([]GetAllRoles_Out_Roles, buffer.ReadUint8())
	for i := 0; i < len(this.Roles); i++ {
		this.Roles[i].Decode(buffer)
	}
}

func (this *GetAllRoles_Out_Roles) Decode(buffer *net.Buffer) {
	this.RoleId = int8(buffer.ReadUint8())
	this.Level = int16(buffer.ReadUint16LE())
	this.Exp = int64(buffer.ReadUint64LE())
	this.FriendshipLevel = int16(buffer.ReadUint16LE())
	this.InForm = buffer.ReadUint8() == 1
	this.Status = int8(buffer.ReadUint8())
	this.CoopId = int16(buffer.ReadUint16LE())
}

func (this *GetAllRoles_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(3)
	buffer.WriteUint8(0)
	buffer.WriteUint8(uint8(len(this.Roles)))
	for i := 0; i < len(this.Roles); i++ {
		this.Roles[i].Encode(buffer)
	}
}

func (this *GetAllRoles_Out_Roles) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint16LE(uint16(this.Level))
	buffer.WriteUint64LE(uint64(this.Exp))
	buffer.WriteUint16LE(uint16(this.FriendshipLevel))
	if this.InForm {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
	buffer.WriteUint8(uint8(this.Status))
	buffer.WriteUint16LE(uint16(this.CoopId))
}

func (this *GetAllRoles_Out) ByteSize() int {
	size := 3
	size += len(this.Roles) * 17
	return size
}

func (this *GetRoleInfo_In) Decode(buffer *net.Buffer) {
	this.RoleId = int8(buffer.ReadUint8())
}

func (this *GetRoleInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(3)
	buffer.WriteUint8(1)
	buffer.WriteUint8(uint8(this.RoleId))
}

func (this *GetRoleInfo_In) ByteSize() int {
	size := 3
	return size
}

func (this *GetRoleInfo_Out) Decode(buffer *net.Buffer) {
	this.RoleId = int8(buffer.ReadUint8())
	this.Level = int16(buffer.ReadUint16LE())
	this.Exp = int64(buffer.ReadUint64LE())
	this.FriendshipLevel = int16(buffer.ReadUint16LE())
	this.FightNum = int32(buffer.ReadUint32LE())
	this.Status = int8(buffer.ReadUint8())
	this.CoopId = int16(buffer.ReadUint16LE())
	this.Attack = int32(buffer.ReadUint32LE())
	this.Defence = int32(buffer.ReadUint32LE())
	this.Health = int32(buffer.ReadUint32LE())
	this.Speed = int32(buffer.ReadUint32LE())
	this.Cultivation = int32(buffer.ReadUint32LE())
	this.Sunder = int32(buffer.ReadUint32LE())
	this.GhostPower = int32(buffer.ReadUint32LE())
	this.HitLevel = int32(buffer.ReadUint32LE())
	this.CriticalLevel = int32(buffer.ReadUint32LE())
	this.SleepLevel = int32(buffer.ReadUint32LE())
	this.DizzinessLevel = int32(buffer.ReadUint32LE())
	this.RandomLevel = int32(buffer.ReadUint32LE())
	this.DisableSkillLevel = int32(buffer.ReadUint32LE())
	this.PoisoningLevel = int32(buffer.ReadUint32LE())
	this.BlockLevel = int32(buffer.ReadUint32LE())
	this.DestroyLevel = int32(buffer.ReadUint32LE())
	this.CriticalHurtLevel = int32(buffer.ReadUint32LE())
	this.TenacityLevel = int32(buffer.ReadUint32LE())
	this.DodgeLevel = int32(buffer.ReadUint32LE())
}

func (this *GetRoleInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(3)
	buffer.WriteUint8(1)
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint16LE(uint16(this.Level))
	buffer.WriteUint64LE(uint64(this.Exp))
	buffer.WriteUint16LE(uint16(this.FriendshipLevel))
	buffer.WriteUint32LE(uint32(this.FightNum))
	buffer.WriteUint8(uint8(this.Status))
	buffer.WriteUint16LE(uint16(this.CoopId))
	buffer.WriteUint32LE(uint32(this.Attack))
	buffer.WriteUint32LE(uint32(this.Defence))
	buffer.WriteUint32LE(uint32(this.Health))
	buffer.WriteUint32LE(uint32(this.Speed))
	buffer.WriteUint32LE(uint32(this.Cultivation))
	buffer.WriteUint32LE(uint32(this.Sunder))
	buffer.WriteUint32LE(uint32(this.GhostPower))
	buffer.WriteUint32LE(uint32(this.HitLevel))
	buffer.WriteUint32LE(uint32(this.CriticalLevel))
	buffer.WriteUint32LE(uint32(this.SleepLevel))
	buffer.WriteUint32LE(uint32(this.DizzinessLevel))
	buffer.WriteUint32LE(uint32(this.RandomLevel))
	buffer.WriteUint32LE(uint32(this.DisableSkillLevel))
	buffer.WriteUint32LE(uint32(this.PoisoningLevel))
	buffer.WriteUint32LE(uint32(this.BlockLevel))
	buffer.WriteUint32LE(uint32(this.DestroyLevel))
	buffer.WriteUint32LE(uint32(this.CriticalHurtLevel))
	buffer.WriteUint32LE(uint32(this.TenacityLevel))
	buffer.WriteUint32LE(uint32(this.DodgeLevel))
}

func (this *GetRoleInfo_Out) ByteSize() int {
	size := 98
	return size
}

func (this *GetPlayerInfo_In) Decode(buffer *net.Buffer) {
	this.Pid = int64(buffer.ReadUint64LE())
}

func (this *GetPlayerInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(3)
	buffer.WriteUint8(2)
	buffer.WriteUint64LE(uint64(this.Pid))
}

func (this *GetPlayerInfo_In) ByteSize() int {
	size := 10
	return size
}

func (this *GetPlayerInfo_Out) Decode(buffer *net.Buffer) {
	this.Player.Decode(buffer)
}

func (this *GetPlayerInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(3)
	buffer.WriteUint8(2)
	this.Player.Encode(buffer)
}

func (this *GetPlayerInfo_Out) ByteSize() int {
	size := 2
	size += this.Player.ByteSize()
	return size
}

func (this *GetFightNum_In) Decode(buffer *net.Buffer) {
	this.FightType = FightnumType(buffer.ReadUint8())
}

func (this *GetFightNum_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(3)
	buffer.WriteUint8(3)
	buffer.WriteUint8(uint8(this.FightType))
}

func (this *GetFightNum_In) ByteSize() int {
	size := 3
	return size
}

func (this *GetFightNum_Out) Decode(buffer *net.Buffer) {
	this.FightNums = make([]GetFightNum_Out_FightNums, buffer.ReadUint8())
	for i := 0; i < len(this.FightNums); i++ {
		this.FightNums[i].Decode(buffer)
	}
}

func (this *GetFightNum_Out_FightNums) Decode(buffer *net.Buffer) {
	this.FightType = int16(buffer.ReadUint16LE())
	this.FightNum = int32(buffer.ReadUint32LE())
}

func (this *GetFightNum_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(3)
	buffer.WriteUint8(3)
	buffer.WriteUint8(uint8(len(this.FightNums)))
	for i := 0; i < len(this.FightNums); i++ {
		this.FightNums[i].Encode(buffer)
	}
}

func (this *GetFightNum_Out_FightNums) Encode(buffer *net.Buffer) {
	buffer.WriteUint16LE(uint16(this.FightType))
	buffer.WriteUint32LE(uint32(this.FightNum))
}

func (this *GetFightNum_Out) ByteSize() int {
	size := 3
	size += len(this.FightNums) * 6
	return size
}

func (this *GetPlayerInfoWithOpenid_In) Decode(buffer *net.Buffer) {
	this.Openid = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.GameServerId = int32(buffer.ReadUint32LE())
}

func (this *GetPlayerInfoWithOpenid_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(3)
	buffer.WriteUint8(4)
	buffer.WriteUint16LE(uint16(len(this.Openid)))
	buffer.WriteBytes(this.Openid)
	buffer.WriteUint32LE(uint32(this.GameServerId))
}

func (this *GetPlayerInfoWithOpenid_In) ByteSize() int {
	size := 8
	size += len(this.Openid)
	return size
}

func (this *GetPlayerInfoWithOpenid_Out) Decode(buffer *net.Buffer) {
	this.Player.Decode(buffer)
}

func (this *GetPlayerInfoWithOpenid_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(3)
	buffer.WriteUint8(4)
	this.Player.Encode(buffer)
}

func (this *GetPlayerInfoWithOpenid_Out) ByteSize() int {
	size := 2
	size += this.Player.ByteSize()
	return size
}

func (this *LevelupRoleFriendship_In) Decode(buffer *net.Buffer) {
	this.RoleId = int8(buffer.ReadUint8())
}

func (this *LevelupRoleFriendship_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(3)
	buffer.WriteUint8(5)
	buffer.WriteUint8(uint8(this.RoleId))
}

func (this *LevelupRoleFriendship_In) ByteSize() int {
	size := 3
	return size
}

func (this *LevelupRoleFriendship_Out) Decode(buffer *net.Buffer) {
}

func (this *LevelupRoleFriendship_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(3)
	buffer.WriteUint8(5)
}

func (this *LevelupRoleFriendship_Out) ByteSize() int {
	size := 2
	return size
}

func (this *RecruitBuddy_In) Decode(buffer *net.Buffer) {
	this.RoleId = int8(buffer.ReadUint8())
}

func (this *RecruitBuddy_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(3)
	buffer.WriteUint8(6)
	buffer.WriteUint8(uint8(this.RoleId))
}

func (this *RecruitBuddy_In) ByteSize() int {
	size := 3
	return size
}

func (this *RecruitBuddy_Out) Decode(buffer *net.Buffer) {
	this.Result = int8(buffer.ReadUint8())
}

func (this *RecruitBuddy_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(3)
	buffer.WriteUint8(6)
	buffer.WriteUint8(uint8(this.Result))
}

func (this *RecruitBuddy_Out) ByteSize() int {
	size := 3
	return size
}

func (this *GetRoleFightNum_In) Decode(buffer *net.Buffer) {
	this.RoleId = int8(buffer.ReadUint8())
}

func (this *GetRoleFightNum_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(3)
	buffer.WriteUint8(7)
	buffer.WriteUint8(uint8(this.RoleId))
}

func (this *GetRoleFightNum_In) ByteSize() int {
	size := 3
	return size
}

func (this *GetRoleFightNum_Out) Decode(buffer *net.Buffer) {
	this.FightNums = make([]GetRoleFightNum_Out_FightNums, buffer.ReadUint8())
	for i := 0; i < len(this.FightNums); i++ {
		this.FightNums[i].Decode(buffer)
	}
}

func (this *GetRoleFightNum_Out_FightNums) Decode(buffer *net.Buffer) {
	this.RoleId = int8(buffer.ReadUint8())
	this.FightNum = int32(buffer.ReadUint32LE())
}

func (this *GetRoleFightNum_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(3)
	buffer.WriteUint8(7)
	buffer.WriteUint8(uint8(len(this.FightNums)))
	for i := 0; i < len(this.FightNums); i++ {
		this.FightNums[i].Encode(buffer)
	}
}

func (this *GetRoleFightNum_Out_FightNums) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint32LE(uint32(this.FightNum))
}

func (this *GetRoleFightNum_Out) ByteSize() int {
	size := 3
	size += len(this.FightNums) * 5
	return size
}

func (this *ChangeRoleStatus_In) Decode(buffer *net.Buffer) {
	this.RoleId = int8(buffer.ReadUint8())
	this.Status = int8(buffer.ReadUint8())
}

func (this *ChangeRoleStatus_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(3)
	buffer.WriteUint8(8)
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint8(uint8(this.Status))
}

func (this *ChangeRoleStatus_In) ByteSize() int {
	size := 4
	return size
}

func (this *ChangeRoleStatus_Out) Decode(buffer *net.Buffer) {
	this.Result = buffer.ReadUint8() == 1
	this.RoleId = int8(buffer.ReadUint8())
	this.Status = int8(buffer.ReadUint8())
}

func (this *ChangeRoleStatus_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(3)
	buffer.WriteUint8(8)
	if this.Result {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint8(uint8(this.Status))
}

func (this *ChangeRoleStatus_Out) ByteSize() int {
	size := 5
	return size
}

func (this *GetInnRoleList_In) Decode(buffer *net.Buffer) {
}

func (this *GetInnRoleList_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(3)
	buffer.WriteUint8(9)
}

func (this *GetInnRoleList_In) ByteSize() int {
	size := 2
	return size
}

func (this *GetInnRoleList_Out) Decode(buffer *net.Buffer) {
	this.RoleList = make([]GetInnRoleList_Out_RoleList, buffer.ReadUint8())
	for i := 0; i < len(this.RoleList); i++ {
		this.RoleList[i].Decode(buffer)
	}
}

func (this *GetInnRoleList_Out_RoleList) Decode(buffer *net.Buffer) {
	this.RoleId = int8(buffer.ReadUint8())
	this.FriendshipLevel = int16(buffer.ReadUint16LE())
	this.FightNum = int32(buffer.ReadUint32LE())
	this.RoleLevel = int16(buffer.ReadUint16LE())
	this.Operate = int8(buffer.ReadUint8())
}

func (this *GetInnRoleList_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(3)
	buffer.WriteUint8(9)
	buffer.WriteUint8(uint8(len(this.RoleList)))
	for i := 0; i < len(this.RoleList); i++ {
		this.RoleList[i].Encode(buffer)
	}
}

func (this *GetInnRoleList_Out_RoleList) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint16LE(uint16(this.FriendshipLevel))
	buffer.WriteUint32LE(uint32(this.FightNum))
	buffer.WriteUint16LE(uint16(this.RoleLevel))
	buffer.WriteUint8(uint8(this.Operate))
}

func (this *GetInnRoleList_Out) ByteSize() int {
	size := 3
	size += len(this.RoleList) * 10
	return size
}

func (this *BuddyCoop_In) Decode(buffer *net.Buffer) {
	this.CoopId = int16(buffer.ReadUint16LE())
}

func (this *BuddyCoop_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(3)
	buffer.WriteUint8(10)
	buffer.WriteUint16LE(uint16(this.CoopId))
}

func (this *BuddyCoop_In) ByteSize() int {
	size := 4
	return size
}

func (this *BuddyCoop_Out) Decode(buffer *net.Buffer) {
}

func (this *BuddyCoop_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(3)
	buffer.WriteUint8(10)
}

func (this *BuddyCoop_Out) ByteSize() int {
	size := 2
	return size
}
