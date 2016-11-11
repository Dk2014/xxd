package battle_api

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
	StartBattle(*net.Session, *StartBattle_In)
	NextRound(*net.Session, *NextRound_In)
	Escape(*net.Session, *Escape_In)
	StartReady(*net.Session, *StartReady_In)
	CallBattlePet(*net.Session, *CallBattlePet_In)
	UseBuddySkill(*net.Session, *UseBuddySkill_In)
	StartBattleForHijackBoat(*net.Session, *StartBattleForHijackBoat_In)
	StartBattleForRecoverBoat(*net.Session, *StartBattleForRecoverBoat_In)
	RoundReady(*net.Session, *RoundReady_In)
	InitRound(*net.Session, *InitRound_In)
	SetAuto(*net.Session, *SetAuto_In)
	CancelAuto(*net.Session, *CancelAuto_In)
	SetSkill(*net.Session, *SetSkill_In)
	UseItem(*net.Session, *UseItem_In)
	UseGhost(*net.Session, *UseGhost_In)
	BattleReconnect(*net.Session, *BattleReconnect_In)
}

type OutHandler interface {
	StartBattle(*net.Session, *StartBattle_Out)
	NextRound(*net.Session, *NextRound_Out)
	End(*net.Session, *End_Out)
	Escape(*net.Session, *Escape_Out)
	Fightnum(*net.Session, *Fightnum_Out)
	StartReadyTimeout(*net.Session, *StartReadyTimeout_Out)
	StartReady(*net.Session, *StartReady_Out)
	StateChange(*net.Session, *StateChange_Out)
	CallBattlePet(*net.Session, *CallBattlePet_Out)
	UseBuddySkill(*net.Session, *UseBuddySkill_Out)
	CallNewEnemys(*net.Session, *CallNewEnemys_Out)
	NewFighterGroup(*net.Session, *NewFighterGroup_Out)
	StartBattleForHijackBoat(*net.Session, *StartBattleForHijackBoat_Out)
	StartBattleForRecoverBoat(*net.Session, *StartBattleForRecoverBoat_Out)
	RoundReady(*net.Session, *RoundReady_Out)
	InitRound(*net.Session, *InitRound_Out)
	SetAuto(*net.Session, *SetAuto_Out)
	CancelAuto(*net.Session, *CancelAuto_Out)
	SetSkill(*net.Session, *SetSkill_Out)
	UseItem(*net.Session, *UseItem_Out)
	UseGhost(*net.Session, *UseGhost_Out)
	NotifyReady(*net.Session, *NotifyReady_Out)
	BattleReconnect(*net.Session, *BattleReconnect_Out)
}

func DecodeIn(message []byte) Request {
	var actionId = message[0]
	var buffer = net.NewBuffer(message[1:])
	switch actionId {
	case 0:
		request := new(StartBattle_In)
		request.Decode(buffer)
		return request
	case 1:
		request := new(NextRound_In)
		request.Decode(buffer)
		return request
	case 3:
		request := new(Escape_In)
		request.Decode(buffer)
		return request
	case 6:
		request := new(StartReady_In)
		request.Decode(buffer)
		return request
	case 8:
		request := new(CallBattlePet_In)
		request.Decode(buffer)
		return request
	case 9:
		request := new(UseBuddySkill_In)
		request.Decode(buffer)
		return request
	case 12:
		request := new(StartBattleForHijackBoat_In)
		request.Decode(buffer)
		return request
	case 13:
		request := new(StartBattleForRecoverBoat_In)
		request.Decode(buffer)
		return request
	case 14:
		request := new(RoundReady_In)
		request.Decode(buffer)
		return request
	case 15:
		request := new(InitRound_In)
		request.Decode(buffer)
		return request
	case 17:
		request := new(SetAuto_In)
		request.Decode(buffer)
		return request
	case 18:
		request := new(CancelAuto_In)
		request.Decode(buffer)
		return request
	case 19:
		request := new(SetSkill_In)
		request.Decode(buffer)
		return request
	case 20:
		request := new(UseItem_In)
		request.Decode(buffer)
		return request
	case 21:
		request := new(UseGhost_In)
		request.Decode(buffer)
		return request
	case 23:
		request := new(BattleReconnect_In)
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
		request := new(StartBattle_Out)
		request.Decode(buffer)
		return request
	case 1:
		request := new(NextRound_Out)
		request.Decode(buffer)
		return request
	case 2:
		request := new(End_Out)
		request.Decode(buffer)
		return request
	case 3:
		request := new(Escape_Out)
		request.Decode(buffer)
		return request
	case 4:
		request := new(Fightnum_Out)
		request.Decode(buffer)
		return request
	case 5:
		request := new(StartReadyTimeout_Out)
		request.Decode(buffer)
		return request
	case 6:
		request := new(StartReady_Out)
		request.Decode(buffer)
		return request
	case 7:
		request := new(StateChange_Out)
		request.Decode(buffer)
		return request
	case 8:
		request := new(CallBattlePet_Out)
		request.Decode(buffer)
		return request
	case 9:
		request := new(UseBuddySkill_Out)
		request.Decode(buffer)
		return request
	case 10:
		request := new(CallNewEnemys_Out)
		request.Decode(buffer)
		return request
	case 11:
		request := new(NewFighterGroup_Out)
		request.Decode(buffer)
		return request
	case 12:
		request := new(StartBattleForHijackBoat_Out)
		request.Decode(buffer)
		return request
	case 13:
		request := new(StartBattleForRecoverBoat_Out)
		request.Decode(buffer)
		return request
	case 14:
		request := new(RoundReady_Out)
		request.Decode(buffer)
		return request
	case 15:
		request := new(InitRound_Out)
		request.Decode(buffer)
		return request
	case 17:
		request := new(SetAuto_Out)
		request.Decode(buffer)
		return request
	case 18:
		request := new(CancelAuto_Out)
		request.Decode(buffer)
		return request
	case 19:
		request := new(SetSkill_Out)
		request.Decode(buffer)
		return request
	case 20:
		request := new(UseItem_Out)
		request.Decode(buffer)
		return request
	case 21:
		request := new(UseGhost_Out)
		request.Decode(buffer)
		return request
	case 22:
		request := new(NotifyReady_Out)
		request.Decode(buffer)
		return request
	case 23:
		request := new(BattleReconnect_Out)
		request.Decode(buffer)
		return request
	}
	_ = buffer
	panic("unsupported response")
}

type FighterType int8

const (
	FIGHTER_TYPE_ATK FighterType = 0
	FIGHTER_TYPE_DEF FighterType = 1
)

type FighterKind int8

const (
	FIGHTER_KIND_PLAYER FighterKind = 0
	FIGHTER_KIND_BUDDY  FighterKind = 1
	FIGHTER_KIND_ENEMY  FighterKind = 2
)

type RoundEvent int8

const (
	ROUND_EVENT_NONE    RoundEvent = 0
	ROUND_EVENT_DODGE   RoundEvent = 1
	ROUND_EVENT_CRIT    RoundEvent = 2
	ROUND_EVENT_BLOCK   RoundEvent = 3
	ROUND_EVENT_SQUELCH RoundEvent = 4
)

type RoundStatus int8

const (
	ROUND_STATUS_NOT_END             RoundStatus = 0
	ROUND_STATUS_ATK_WIN             RoundStatus = 1
	ROUND_STATUS_DEF_WIN             RoundStatus = 2
	ROUND_STATUS_ATK_NEXT            RoundStatus = 3
	ROUND_STATUS_DEF_NEXT            RoundStatus = 4
	ROUND_STATUS_TRIGGER_CALL_ENEMYS RoundStatus = 5
	ROUND_STATUS_WAITING             RoundStatus = 6
)

type BuffMode int8

const (
	BUFF_MODE_POWER                      BuffMode = 0
	BUFF_MODE_SPEED                      BuffMode = 1
	BUFF_MODE_ATTACK                     BuffMode = 2
	BUFF_MODE_DEFEND                     BuffMode = 3
	BUFF_MODE_HEALTH                     BuffMode = 4
	BUFF_MODE_DIZZINESS                  BuffMode = 5
	BUFF_MODE_POISONING                  BuffMode = 6
	BUFF_MODE_CLEAN_BAD                  BuffMode = 7
	BUFF_MODE_CLEAN_GOOD                 BuffMode = 8
	BUFF_MODE_REDUCE_HURT                BuffMode = 9
	BUFF_MODE_RANDOM                     BuffMode = 10
	BUFF_MODE_BLOCK                      BuffMode = 11
	BUFF_MODE_BLOCK_LEVEL                BuffMode = 12
	BUFF_MODE_DODGE_LEVEL                BuffMode = 13
	BUFF_MODE_CRITIAL_LEVEL              BuffMode = 14
	BUFF_MODE_HIT_LEVEL                  BuffMode = 15
	BUFF_MODE_HURT_ADD                   BuffMode = 16
	BUFF_MODE_MAX_HEALTH                 BuffMode = 17
	BUFF_MODE_KEEPER_REDUCE_HURT         BuffMode = 18
	BUFF_MODE_ATTRACT_FIRE               BuffMode = 19
	BUFF_MODE_DESTROY_LEVEL              BuffMode = 20
	BUFF_MODE_TENACITY_LEVEL             BuffMode = 21
	BUFF_MODE_SUNDER                     BuffMode = 22
	BUFF_MODE_SLEEP                      BuffMode = 23
	BUFF_MODE_DISABLE_SKILL              BuffMode = 24
	BUFF_MODE_BUFF_DIRECT_REDUCE_HURT    BuffMode = 25
	BUFF_MODE_BUFF_ABSORB_HURT           BuffMode = 26
	BUFF_MODE_BUFF_GHOST_POWER           BuffMode = 27
	BUFF_MODE_BUFF_PET_LIVE_ROUND        BuffMode = 28
	BUFF_MODE_BUFF_BUDDY_SKILL           BuffMode = 29
	BUFF_MODE_BUFF_CLEAR_ABSORB_HURT     BuffMode = 30
	BUFF_MODE_BUFF_SLEEP_LEVEL           BuffMode = 31
	BUFF_MODE_BUFF_DIZZINESS_LEVEL       BuffMode = 32
	BUFF_MODE_BUFF_RANDOM_LEVEL          BuffMode = 33
	BUFF_MODE_BUFF_DISABLE_SKILL_LEVEL   BuffMode = 34
	BUFF_MODE_BUFF_POISONING_LEVEL       BuffMode = 35
	BUFF_MODE_BUFF_RECOVER_BUDDY_SKILL   BuffMode = 36
	BUFF_MODE_BUFF_MAKE_POWER_FULL       BuffMode = 37
	BUFF_MODE_BUFF_DOGE                  BuffMode = 38
	BUFF_MODE_BUFF_HIT                   BuffMode = 39
	BUFF_MODE_BUFF_CRITIAL               BuffMode = 40
	BUFF_MODE_BUFF_TENACITY              BuffMode = 41
	BUFF_MODE_BUFF_TAKE_SUNSER           BuffMode = 42
	BUFF_MODE_BUFF_DEFEND_PERSENT        BuffMode = 43
	BUFF_MODE_BUFF_SUNDER_STATE          BuffMode = 44
	BUFF_MODE_BUFF_HEALTH_PERCENT        BuffMode = 45
	BUFF_MODE_BUFF_ALL_RESIST            BuffMode = 46
	BUFF_MODE_BUFF_REBOTH_HEALTH         BuffMode = 47
	BUFF_MODE_BUFF_REBOTH_HEALTH_PERCENT BuffMode = 48
)

type BattleType int8

const (
	BATTLE_TYPE_MISSION                BattleType = 0
	BATTLE_TYPE_RESOURCE               BattleType = 1
	BATTLE_TYPE_TOWER                  BattleType = 2
	BATTLE_TYPE_MULTILEVEL             BattleType = 3
	BATTLE_TYPE_ARENA                  BattleType = 4
	BATTLE_TYPE_HARD                   BattleType = 8
	BATTLE_TYPE_BUDDY                  BattleType = 9
	BATTLE_TYPE_PET                    BattleType = 10
	BATTLE_TYPE_GHOST                  BattleType = 11
	BATTLE_TYPE_RAINBOW                BattleType = 12
	BATTLE_TYPE_PVE                    BattleType = 13
	BATTLE_TYPE_FATE_BOX               BattleType = 14
	BATTLE_TYPE_DRIVING_EXPLORING      BattleType = 15
	BATTLE_TYPE_DRIVING_SWORD_BF_LEVEL BattleType = 16
	BATTLE_TYPE_HIJACK_BOAT            BattleType = 17
	BATTLE_TYPE_RECOVER_BOAT           BattleType = 18
	BATTLE_TYPE_DESPAIR                BattleType = 19
	BATTLE_TYPE_DESPAIR_BOSS           BattleType = 20
)

type JobType int8

const (
	JOB_TYPE_NONE       JobType = 0
	JOB_TYPE_ATTACKER   JobType = 1
	JOB_TYPE_DESTROYER  JobType = 2
	JOB_TYPE_DEFENDER   JobType = 3
	JOB_TYPE_TREATER    JobType = 4
	JOB_TYPE_SUPPORTER  JobType = 5
	JOB_TYPE_OBSTRUCTOR JobType = 6
)

type BattleRole struct {
	Kind                FighterKind         `json:"kind"`
	PlayerId            int64               `json:"player_id"`
	RoleId              int32               `json:"role_id"`
	RoleLevel           int16               `json:"role_level"`
	Position            int32               `json:"position"`
	FashionId           int16               `json:"fashion_id"`
	FriendshipLevel     int16               `json:"friendship_level"`
	Health              int32               `json:"health"`
	MaxHealth           int32               `json:"max_health"`
	Power               int16               `json:"power"`
	MaxPower            int16               `json:"max_power"`
	SunderValue         int16               `json:"sunder_value"`
	SunderMaxValue      int16               `json:"sunder_max_value"`
	SunderMinHurtRate   int16               `json:"sunder_min_hurt_rate"`
	SunderEndHurtRate   int16               `json:"sunder_end_hurt_rate"`
	SunderEndDefendRate int16               `json:"sunder_end_defend_rate"`
	Speed               int32               `json:"speed"`
	GhostShieldValue    int32               `json:"ghost_shield_value"`
	GhostPower          int32               `json:"ghost_power"`
	CanUseGhost         bool                `json:"can_use_ghost"`
	Ghosts              []BattleRole_Ghosts `json:"ghosts"`
	Auras               []BattleRole_Auras  `json:"auras"`
	CouldUseSwordSoul   bool                `json:"could_use_sword_soul"`
}

type BattleRole_Ghosts struct {
	GhostId      int16 `json:"ghost_id"`
	GhostStar    int8  `json:"ghost_star"`
	GhostLevel   int16 `json:"ghost_level"`
	GhostSkillId int32 `json:"ghost_skill_id"`
	RelatedGhost int16 `json:"related_ghost"`
	Used         bool  `json:"used"`
}

type BattleRole_Auras struct {
	AuraId int16 `json:"aura_id"`
	Value  int32 `json:"value"`
}

type BufferInfo struct {
	Mode        BuffMode `json:"mode"`
	Keep        int8     `json:"keep"`
	Value       int32    `json:"value"`
	SkillId     int16    `json:"skill_id"`
	MaxOverride int8     `json:"max_override"`
	OverrideNum int8     `json:"override_num"`
	Uncleanable bool     `json:"uncleanable"`
}

type SkillInfo struct {
	SkillId  int16 `json:"skill_id"`
	IncPower int8  `json:"inc_power"`
	DecPower int8  `json:"dec_power"`
}

func (this *BattleRole) Decode(buffer *net.Buffer) {
	this.Kind = FighterKind(buffer.ReadUint8())
	this.PlayerId = int64(buffer.ReadUint64LE())
	this.RoleId = int32(buffer.ReadUint32LE())
	this.RoleLevel = int16(buffer.ReadUint16LE())
	this.Position = int32(buffer.ReadUint32LE())
	this.FashionId = int16(buffer.ReadUint16LE())
	this.FriendshipLevel = int16(buffer.ReadUint16LE())
	this.Health = int32(buffer.ReadUint32LE())
	this.MaxHealth = int32(buffer.ReadUint32LE())
	this.Power = int16(buffer.ReadUint16LE())
	this.MaxPower = int16(buffer.ReadUint16LE())
	this.SunderValue = int16(buffer.ReadUint16LE())
	this.SunderMaxValue = int16(buffer.ReadUint16LE())
	this.SunderMinHurtRate = int16(buffer.ReadUint16LE())
	this.SunderEndHurtRate = int16(buffer.ReadUint16LE())
	this.SunderEndDefendRate = int16(buffer.ReadUint16LE())
	this.Speed = int32(buffer.ReadUint32LE())
	this.GhostShieldValue = int32(buffer.ReadUint32LE())
	this.GhostPower = int32(buffer.ReadUint32LE())
	this.CanUseGhost = buffer.ReadUint8() == 1
	this.Ghosts = make([]BattleRole_Ghosts, buffer.ReadUint8())
	for i := 0; i < len(this.Ghosts); i++ {
		this.Ghosts[i].Decode(buffer)
	}
	this.Auras = make([]BattleRole_Auras, buffer.ReadUint8())
	for i := 0; i < len(this.Auras); i++ {
		this.Auras[i].Decode(buffer)
	}
	this.CouldUseSwordSoul = buffer.ReadUint8() == 1
}

func (this *BattleRole_Ghosts) Decode(buffer *net.Buffer) {
	this.GhostId = int16(buffer.ReadUint16LE())
	this.GhostStar = int8(buffer.ReadUint8())
	this.GhostLevel = int16(buffer.ReadUint16LE())
	this.GhostSkillId = int32(buffer.ReadUint32LE())
	this.RelatedGhost = int16(buffer.ReadUint16LE())
	this.Used = buffer.ReadUint8() == 1
}

func (this *BattleRole_Auras) Decode(buffer *net.Buffer) {
	this.AuraId = int16(buffer.ReadUint16LE())
	this.Value = int32(buffer.ReadUint32LE())
}

func (this *BattleRole) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(this.Kind))
	buffer.WriteUint64LE(uint64(this.PlayerId))
	buffer.WriteUint32LE(uint32(this.RoleId))
	buffer.WriteUint16LE(uint16(this.RoleLevel))
	buffer.WriteUint32LE(uint32(this.Position))
	buffer.WriteUint16LE(uint16(this.FashionId))
	buffer.WriteUint16LE(uint16(this.FriendshipLevel))
	buffer.WriteUint32LE(uint32(this.Health))
	buffer.WriteUint32LE(uint32(this.MaxHealth))
	buffer.WriteUint16LE(uint16(this.Power))
	buffer.WriteUint16LE(uint16(this.MaxPower))
	buffer.WriteUint16LE(uint16(this.SunderValue))
	buffer.WriteUint16LE(uint16(this.SunderMaxValue))
	buffer.WriteUint16LE(uint16(this.SunderMinHurtRate))
	buffer.WriteUint16LE(uint16(this.SunderEndHurtRate))
	buffer.WriteUint16LE(uint16(this.SunderEndDefendRate))
	buffer.WriteUint32LE(uint32(this.Speed))
	buffer.WriteUint32LE(uint32(this.GhostShieldValue))
	buffer.WriteUint32LE(uint32(this.GhostPower))
	if this.CanUseGhost {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
	buffer.WriteUint8(uint8(len(this.Ghosts)))
	for i := 0; i < len(this.Ghosts); i++ {
		this.Ghosts[i].Encode(buffer)
	}
	buffer.WriteUint8(uint8(len(this.Auras)))
	for i := 0; i < len(this.Auras); i++ {
		this.Auras[i].Encode(buffer)
	}
	if this.CouldUseSwordSoul {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *BattleRole_Ghosts) Encode(buffer *net.Buffer) {
	buffer.WriteUint16LE(uint16(this.GhostId))
	buffer.WriteUint8(uint8(this.GhostStar))
	buffer.WriteUint16LE(uint16(this.GhostLevel))
	buffer.WriteUint32LE(uint32(this.GhostSkillId))
	buffer.WriteUint16LE(uint16(this.RelatedGhost))
	if this.Used {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *BattleRole_Auras) Encode(buffer *net.Buffer) {
	buffer.WriteUint16LE(uint16(this.AuraId))
	buffer.WriteUint32LE(uint32(this.Value))
}

func (this *BattleRole) ByteSize() int {
	size := 61
	size += len(this.Ghosts) * 12
	size += len(this.Auras) * 6
	return size
}

func (this *BufferInfo) Decode(buffer *net.Buffer) {
	this.Mode = BuffMode(buffer.ReadUint8())
	this.Keep = int8(buffer.ReadUint8())
	this.Value = int32(buffer.ReadUint32LE())
	this.SkillId = int16(buffer.ReadUint16LE())
	this.MaxOverride = int8(buffer.ReadUint8())
	this.OverrideNum = int8(buffer.ReadUint8())
	this.Uncleanable = buffer.ReadUint8() == 1
}

func (this *BufferInfo) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(this.Mode))
	buffer.WriteUint8(uint8(this.Keep))
	buffer.WriteUint32LE(uint32(this.Value))
	buffer.WriteUint16LE(uint16(this.SkillId))
	buffer.WriteUint8(uint8(this.MaxOverride))
	buffer.WriteUint8(uint8(this.OverrideNum))
	if this.Uncleanable {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *BufferInfo) ByteSize() int {
	size := 11
	return size
}

func (this *SkillInfo) Decode(buffer *net.Buffer) {
	this.SkillId = int16(buffer.ReadUint16LE())
	this.IncPower = int8(buffer.ReadUint8())
	this.DecPower = int8(buffer.ReadUint8())
}

func (this *SkillInfo) Encode(buffer *net.Buffer) {
	buffer.WriteUint16LE(uint16(this.SkillId))
	buffer.WriteUint8(uint8(this.IncPower))
	buffer.WriteUint8(uint8(this.DecPower))
}

func (this *SkillInfo) ByteSize() int {
	size := 4
	return size
}

type StartBattle_In struct {
	BattleType BattleType `json:"battle_type"`
	BattleId   int64      `json:"battle_id"`
}

func (this *StartBattle_In) Process(session *net.Session) {
	g_InHandler.StartBattle(session, this)
}

func (this *StartBattle_In) TypeName() string {
	return "battle.start_battle.in"
}

func (this *StartBattle_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 6, 0
}

type StartBattle_Out struct {
	TotalGroup        int8                                `json:"total_group"`
	AttackerPlayerIds []StartBattle_Out_AttackerPlayerIds `json:"attacker_player_ids"`
	IsMainRoleFirst   bool                                `json:"is_main_role_first"`
	IsAttackerFirst   bool                                `json:"is_attacker_first"`
	AllAttackers      []StartBattle_Out_AllAttackers      `json:"all_attackers"`
	AllDefenders      []StartBattle_Out_AllDefenders      `json:"all_defenders"`
	AttackerTotems    []StartBattle_Out_AttackerTotems    `json:"attacker_totems"`
	DefenderTotems    []StartBattle_Out_DefenderTotems    `json:"defender_totems"`
	AttackerGroups    []StartBattle_Out_AttackerGroups    `json:"attacker_groups"`
	DefenderGroups    []StartBattle_Out_DefenderGroups    `json:"defender_groups"`
}

type StartBattle_Out_AttackerPlayerIds struct {
	PlayerId int64 `json:"player_id"`
}

type StartBattle_Out_AllAttackers struct {
	PlayerId        int64 `json:"player_id"`
	GhostSkillIndex int8  `json:"ghost_skill_index"`
}

type StartBattle_Out_AllDefenders struct {
	PlayerId        int64 `json:"player_id"`
	GhostSkillIndex int8  `json:"ghost_skill_index"`
}

type StartBattle_Out_AttackerTotems struct {
	Round   int16 `json:"round"`
	TotemId int16 `json:"totem_id"`
}

type StartBattle_Out_DefenderTotems struct {
	Round   int16 `json:"round"`
	TotemId int16 `json:"totem_id"`
}

type StartBattle_Out_AttackerGroups struct {
	Attackers  []StartBattle_Out_AttackerGroups_Attackers  `json:"attackers"`
	SelfBuffs  []StartBattle_Out_AttackerGroups_SelfBuffs  `json:"self_buffs"`
	BuddyBuffs []StartBattle_Out_AttackerGroups_BuddyBuffs `json:"buddy_buffs"`
}

type StartBattle_Out_AttackerGroups_Attackers struct {
	Role   BattleRole                                        `json:"role"`
	Skills []StartBattle_Out_AttackerGroups_Attackers_Skills `json:"skills"`
}

type StartBattle_Out_AttackerGroups_Attackers_Skills struct {
	Skill          SkillInfo `json:"skill"`
	RestReleaseNum int16     `json:"rest_release_num"`
}

type StartBattle_Out_AttackerGroups_SelfBuffs struct {
	Buffer BufferInfo `json:"buffer"`
}

type StartBattle_Out_AttackerGroups_BuddyBuffs struct {
	Pos    int8       `json:"pos"`
	Buffer BufferInfo `json:"buffer"`
}

type StartBattle_Out_DefenderGroups struct {
	Defenders []StartBattle_Out_DefenderGroups_Defenders `json:"defenders"`
}

type StartBattle_Out_DefenderGroups_Defenders struct {
	Role   BattleRole                                        `json:"role"`
	Skills []StartBattle_Out_DefenderGroups_Defenders_Skills `json:"skills"`
}

type StartBattle_Out_DefenderGroups_Defenders_Skills struct {
	Skill          SkillInfo `json:"skill"`
	SkillId2       int16     `json:"skill_id2"`
	RestReleaseNum int16     `json:"rest_release_num"`
}

func (this *StartBattle_Out) Process(session *net.Session) {
	g_OutHandler.StartBattle(session, this)
}

func (this *StartBattle_Out) TypeName() string {
	return "battle.start_battle.out"
}

func (this *StartBattle_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 6, 0
}

func (this *StartBattle_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type NextRound_In struct {
	UseSkill              int8  `json:"use_skill"`
	UseItem               int16 `json:"use_item"`
	AutoFight             bool  `json:"auto_fight"`
	IsAttacker            bool  `json:"is_attacker"`
	Position              int8  `json:"position"`
	JobIndex              int8  `json:"job_index"`
	SendNum               int16 `json:"send_num"`
	UseSwordSoul          bool  `json:"use_sword_soul"`
	UseGhostSkillPosition int8  `json:"use_ghost_skill_position"`
	UseGhostSkillId       int32 `json:"use_ghost_skill_id"`
	UseTotem              bool  `json:"use_totem"`
}

func (this *NextRound_In) Process(session *net.Session) {
	g_InHandler.NextRound(session, this)
}

func (this *NextRound_In) TypeName() string {
	return "battle.next_round.in"
}

func (this *NextRound_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 6, 1
}

type NextRound_Out struct {
	Status       RoundStatus                  `json:"status"`
	NowRound     int16                        `json:"now_round"`
	AllAttackers []NextRound_Out_AllAttackers `json:"all_attackers"`
	AllDefenders []NextRound_Out_AllDefenders `json:"all_defenders"`
	Results      []NextRound_Out_Results      `json:"results"`
	Autos        []NextRound_Out_Autos        `json:"autos"`
}

type NextRound_Out_AllAttackers struct {
	PlayerId int64 `json:"player_id"`
}

type NextRound_Out_AllDefenders struct {
	PlayerId int64 `json:"player_id"`
}

type NextRound_Out_Results struct {
	Ftype         FighterType                     `json:"ftype"`
	Event         RoundEvent                      `json:"event"`
	Position      int8                            `json:"position"`
	Power         int16                           `json:"power"`
	Health        int32                           `json:"health"`
	SunderValue   int16                           `json:"sunder_value"`
	UseGhostSkill bool                            `json:"use_ghost_skill"`
	TotemId       int16                           `json:"totem_id"`
	GhostId       int16                           `json:"ghost_id"`
	GhostShieldOn bool                            `json:"ghost_shield_on"`
	ShieldGhostId int16                           `json:"shield_ghost_id"`
	GhostPower    int32                           `json:"ghost_power"`
	AddPower      int32                           `json:"add_power"`
	Attacks       []NextRound_Out_Results_Attacks `json:"attacks"`
	Item          []NextRound_Out_Results_Item    `json:"item"`
}

type NextRound_Out_Results_Attacks struct {
	SkillId        int32                                      `json:"skill_id"`
	RestReleaseNum int16                                      `json:"rest_release_num"`
	Targets        []NextRound_Out_Results_Attacks_Targets    `json:"targets"`
	SelfBuffs      []NextRound_Out_Results_Attacks_SelfBuffs  `json:"self_buffs"`
	BuddyBuffs     []NextRound_Out_Results_Attacks_BuddyBuffs `json:"buddy_buffs"`
}

type NextRound_Out_Results_Attacks_Targets struct {
	Ftype            FighterType                                           `json:"ftype"`
	Hurt             int32                                                 `json:"hurt"`
	Event            RoundEvent                                            `json:"event"`
	Position         int8                                                  `json:"position"`
	TakeSunder       int16                                                 `json:"take_sunder"`
	TakeGhostShield  int32                                                 `json:"take_ghost_shield"`
	DirectReductHurt int32                                                 `json:"direct_reduct_hurt"`
	GhostShieldOn    bool                                                  `json:"ghost_shield_on"`
	ShieldGhostId    int16                                                 `json:"shield_ghost_id"`
	GhostPower       int32                                                 `json:"ghost_power"`
	Buffs            []NextRound_Out_Results_Attacks_Targets_Buffs         `json:"buffs"`
	PassiveAttack    []NextRound_Out_Results_Attacks_Targets_PassiveAttack `json:"passive_attack"`
}

type NextRound_Out_Results_Attacks_Targets_Buffs struct {
	Buffer BufferInfo `json:"buffer"`
}

type NextRound_Out_Results_Attacks_Targets_PassiveAttack struct {
	SkillId      int32                                                              `json:"skill_id"`
	TargetsBuffs []NextRound_Out_Results_Attacks_Targets_PassiveAttack_TargetsBuffs `json:"targets_buffs"`
	TeamBuffs    []NextRound_Out_Results_Attacks_Targets_PassiveAttack_TeamBuffs    `json:"team_buffs"`
}

type NextRound_Out_Results_Attacks_Targets_PassiveAttack_TargetsBuffs struct {
	Pos    int8       `json:"pos"`
	Buffer BufferInfo `json:"buffer"`
}

type NextRound_Out_Results_Attacks_Targets_PassiveAttack_TeamBuffs struct {
	Pos    int8       `json:"pos"`
	Buffer BufferInfo `json:"buffer"`
}

type NextRound_Out_Results_Attacks_SelfBuffs struct {
	Buffer BufferInfo `json:"buffer"`
}

type NextRound_Out_Results_Attacks_BuddyBuffs struct {
	Pos    int8       `json:"pos"`
	Buffer BufferInfo `json:"buffer"`
}

type NextRound_Out_Results_Item struct {
	ItemId  int32                                `json:"item_id"`
	Targets []NextRound_Out_Results_Item_Targets `json:"targets"`
}

type NextRound_Out_Results_Item_Targets struct {
	Ftype    FighterType                                `json:"ftype"`
	Health   int32                                      `json:"health"`
	Power    int16                                      `json:"power"`
	Hurt     int32                                      `json:"hurt"`
	Position int8                                       `json:"position"`
	Buffs    []NextRound_Out_Results_Item_Targets_Buffs `json:"buffs"`
}

type NextRound_Out_Results_Item_Targets_Buffs struct {
	Buffer BufferInfo `json:"buffer"`
}

type NextRound_Out_Autos struct {
	PlayerId int64 `json:"player_id"`
}

func (this *NextRound_Out) Process(session *net.Session) {
	g_OutHandler.NextRound(session, this)
}

func (this *NextRound_Out) TypeName() string {
	return "battle.next_round.out"
}

func (this *NextRound_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 6, 1
}

func (this *NextRound_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type End_Out struct {
	Status RoundStatus `json:"status"`
}

func (this *End_Out) Process(session *net.Session) {
	g_OutHandler.End(session, this)
}

func (this *End_Out) TypeName() string {
	return "battle.end.out"
}

func (this *End_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 6, 2
}

func (this *End_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type Escape_In struct {
}

func (this *Escape_In) Process(session *net.Session) {
	g_InHandler.Escape(session, this)
}

func (this *Escape_In) TypeName() string {
	return "battle.escape.in"
}

func (this *Escape_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 6, 3
}

type Escape_Out struct {
}

func (this *Escape_Out) Process(session *net.Session) {
	g_OutHandler.Escape(session, this)
}

func (this *Escape_Out) TypeName() string {
	return "battle.escape.out"
}

func (this *Escape_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 6, 3
}

func (this *Escape_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type Fightnum_Out struct {
	Attacker int32 `json:"attacker"`
	Defender int32 `json:"defender"`
}

func (this *Fightnum_Out) Process(session *net.Session) {
	g_OutHandler.Fightnum(session, this)
}

func (this *Fightnum_Out) TypeName() string {
	return "battle.fightnum.out"
}

func (this *Fightnum_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 6, 4
}

func (this *Fightnum_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type StartReadyTimeout_Out struct {
}

func (this *StartReadyTimeout_Out) Process(session *net.Session) {
	g_OutHandler.StartReadyTimeout(session, this)
}

func (this *StartReadyTimeout_Out) TypeName() string {
	return "battle.start_ready_timeout.out"
}

func (this *StartReadyTimeout_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 6, 5
}

func (this *StartReadyTimeout_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type StartReady_In struct {
	Ok bool `json:"ok"`
}

func (this *StartReady_In) Process(session *net.Session) {
	g_InHandler.StartReady(session, this)
}

func (this *StartReady_In) TypeName() string {
	return "battle.start_ready.in"
}

func (this *StartReady_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 6, 6
}

type StartReady_Out struct {
	ReadyPid int64 `json:"ready_pid"`
}

func (this *StartReady_Out) Process(session *net.Session) {
	g_OutHandler.StartReady(session, this)
}

func (this *StartReady_Out) TypeName() string {
	return "battle.start_ready.out"
}

func (this *StartReady_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 6, 6
}

func (this *StartReady_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type StateChange_Out struct {
	PlayerId int64  `json:"player_id"`
	Auto     bool   `json:"auto"`
	Desc     []byte `json:"desc"`
}

func (this *StateChange_Out) Process(session *net.Session) {
	g_OutHandler.StateChange(session, this)
}

func (this *StateChange_Out) TypeName() string {
	return "battle.state_change.out"
}

func (this *StateChange_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 6, 7
}

func (this *StateChange_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type CallBattlePet_In struct {
	GridNum int8 `json:"grid_num"`
}

func (this *CallBattlePet_In) Process(session *net.Session) {
	g_InHandler.CallBattlePet(session, this)
}

func (this *CallBattlePet_In) TypeName() string {
	return "battle.call_battle_pet.in"
}

func (this *CallBattlePet_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 6, 8
}

type CallBattlePet_Out struct {
	Success     bool                       `json:"success"`
	PlayerPower int16                      `json:"player_power"`
	Role        BattleRole                 `json:"role"`
	Skills      []CallBattlePet_Out_Skills `json:"skills"`
}

type CallBattlePet_Out_Skills struct {
	Skill SkillInfo `json:"skill"`
}

func (this *CallBattlePet_Out) Process(session *net.Session) {
	g_OutHandler.CallBattlePet(session, this)
}

func (this *CallBattlePet_Out) TypeName() string {
	return "battle.call_battle_pet.out"
}

func (this *CallBattlePet_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 6, 8
}

func (this *CallBattlePet_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type UseBuddySkill_In struct {
	Pos      int8 `json:"pos"`
	UseSkill int8 `json:"use_skill"`
}

func (this *UseBuddySkill_In) Process(session *net.Session) {
	g_InHandler.UseBuddySkill(session, this)
}

func (this *UseBuddySkill_In) TypeName() string {
	return "battle.use_buddy_skill.in"
}

func (this *UseBuddySkill_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 6, 9
}

type UseBuddySkill_Out struct {
	Pos      int8 `json:"pos"`
	UseSkill int8 `json:"use_skill"`
}

func (this *UseBuddySkill_Out) Process(session *net.Session) {
	g_OutHandler.UseBuddySkill(session, this)
}

func (this *UseBuddySkill_Out) TypeName() string {
	return "battle.use_buddy_skill.out"
}

func (this *UseBuddySkill_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 6, 9
}

func (this *UseBuddySkill_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type CallNewEnemys_Out struct {
	CallInfo []CallNewEnemys_Out_CallInfo `json:"call_info"`
}

type CallNewEnemys_Out_CallInfo struct {
	Ftype       int8                                `json:"ftype"`
	Position    int8                                `json:"position"`
	AttackIndex int8                                `json:"attack_index"`
	Enemys      []CallNewEnemys_Out_CallInfo_Enemys `json:"enemys"`
}

type CallNewEnemys_Out_CallInfo_Enemys struct {
	Role   BattleRole                                 `json:"role"`
	Skills []CallNewEnemys_Out_CallInfo_Enemys_Skills `json:"skills"`
}

type CallNewEnemys_Out_CallInfo_Enemys_Skills struct {
	Skill          SkillInfo `json:"skill"`
	SkillId2       int16     `json:"skill_id2"`
	RestReleaseNum int16     `json:"rest_release_num"`
}

func (this *CallNewEnemys_Out) Process(session *net.Session) {
	g_OutHandler.CallNewEnemys(session, this)
}

func (this *CallNewEnemys_Out) TypeName() string {
	return "battle.call_new_enemys.out"
}

func (this *CallNewEnemys_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 6, 10
}

func (this *CallNewEnemys_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type NewFighterGroup_Out struct {
	Ftype           int8                           `json:"ftype"`
	PlayerId        int64                          `json:"player_id"`
	GhostSkillIndex int8                           `json:"ghost_skill_index"`
	Fighters        []NewFighterGroup_Out_Fighters `json:"fighters"`
}

type NewFighterGroup_Out_Fighters struct {
	Role   BattleRole                            `json:"role"`
	Skills []NewFighterGroup_Out_Fighters_Skills `json:"skills"`
}

type NewFighterGroup_Out_Fighters_Skills struct {
	Skill          SkillInfo `json:"skill"`
	RestReleaseNum int16     `json:"rest_release_num"`
}

func (this *NewFighterGroup_Out) Process(session *net.Session) {
	g_OutHandler.NewFighterGroup(session, this)
}

func (this *NewFighterGroup_Out) TypeName() string {
	return "battle.new_fighter_group.out"
}

func (this *NewFighterGroup_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 6, 11
}

func (this *NewFighterGroup_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type StartBattleForHijackBoat_In struct {
	Pid    int64 `json:"pid"`
	BoatId int64 `json:"boat_id"`
}

func (this *StartBattleForHijackBoat_In) Process(session *net.Session) {
	g_InHandler.StartBattleForHijackBoat(session, this)
}

func (this *StartBattleForHijackBoat_In) TypeName() string {
	return "battle.start_battle_for_hijack_boat.in"
}

func (this *StartBattleForHijackBoat_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 6, 12
}

type StartBattleForHijackBoat_Out struct {
}

func (this *StartBattleForHijackBoat_Out) Process(session *net.Session) {
	g_OutHandler.StartBattleForHijackBoat(session, this)
}

func (this *StartBattleForHijackBoat_Out) TypeName() string {
	return "battle.start_battle_for_hijack_boat.out"
}

func (this *StartBattleForHijackBoat_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 6, 12
}

func (this *StartBattleForHijackBoat_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type StartBattleForRecoverBoat_In struct {
	Pid    int64 `json:"pid"`
	BoatId int64 `json:"boat_id"`
}

func (this *StartBattleForRecoverBoat_In) Process(session *net.Session) {
	g_InHandler.StartBattleForRecoverBoat(session, this)
}

func (this *StartBattleForRecoverBoat_In) TypeName() string {
	return "battle.start_battle_for_recover_boat.in"
}

func (this *StartBattleForRecoverBoat_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 6, 13
}

type StartBattleForRecoverBoat_Out struct {
}

func (this *StartBattleForRecoverBoat_Out) Process(session *net.Session) {
	g_OutHandler.StartBattleForRecoverBoat(session, this)
}

func (this *StartBattleForRecoverBoat_Out) TypeName() string {
	return "battle.start_battle_for_recover_boat.out"
}

func (this *StartBattleForRecoverBoat_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 6, 13
}

func (this *StartBattleForRecoverBoat_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type RoundReady_In struct {
	IsAuto bool `json:"is_auto"`
}

func (this *RoundReady_In) Process(session *net.Session) {
	g_InHandler.RoundReady(session, this)
}

func (this *RoundReady_In) TypeName() string {
	return "battle.round_ready.in"
}

func (this *RoundReady_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 6, 14
}

type RoundReady_Out struct {
}

func (this *RoundReady_Out) Process(session *net.Session) {
	g_OutHandler.RoundReady(session, this)
}

func (this *RoundReady_Out) TypeName() string {
	return "battle.round_ready.out"
}

func (this *RoundReady_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 6, 14
}

func (this *RoundReady_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type InitRound_In struct {
}

func (this *InitRound_In) Process(session *net.Session) {
	g_InHandler.InitRound(session, this)
}

func (this *InitRound_In) TypeName() string {
	return "battle.init_round.in"
}

func (this *InitRound_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 6, 15
}

type InitRound_Out struct {
}

func (this *InitRound_Out) Process(session *net.Session) {
	g_OutHandler.InitRound(session, this)
}

func (this *InitRound_Out) TypeName() string {
	return "battle.init_round.out"
}

func (this *InitRound_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 6, 15
}

func (this *InitRound_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type SetAuto_In struct {
}

func (this *SetAuto_In) Process(session *net.Session) {
	g_InHandler.SetAuto(session, this)
}

func (this *SetAuto_In) TypeName() string {
	return "battle.set_auto.in"
}

func (this *SetAuto_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 6, 17
}

type SetAuto_Out struct {
}

func (this *SetAuto_Out) Process(session *net.Session) {
	g_OutHandler.SetAuto(session, this)
}

func (this *SetAuto_Out) TypeName() string {
	return "battle.set_auto.out"
}

func (this *SetAuto_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 6, 17
}

func (this *SetAuto_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type CancelAuto_In struct {
}

func (this *CancelAuto_In) Process(session *net.Session) {
	g_InHandler.CancelAuto(session, this)
}

func (this *CancelAuto_In) TypeName() string {
	return "battle.cancel_auto.in"
}

func (this *CancelAuto_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 6, 18
}

type CancelAuto_Out struct {
	Round int16 `json:"round"`
}

func (this *CancelAuto_Out) Process(session *net.Session) {
	g_OutHandler.CancelAuto(session, this)
}

func (this *CancelAuto_Out) TypeName() string {
	return "battle.cancel_auto.out"
}

func (this *CancelAuto_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 6, 18
}

func (this *CancelAuto_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type SetSkill_In struct {
	IsAttacker bool `json:"is_attacker"`
	PosIdx     int8 `json:"pos_idx"`
	SkillIdx   int8 `json:"skill_idx"`
}

func (this *SetSkill_In) Process(session *net.Session) {
	g_InHandler.SetSkill(session, this)
}

func (this *SetSkill_In) TypeName() string {
	return "battle.set_skill.in"
}

func (this *SetSkill_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 6, 19
}

type SetSkill_Out struct {
	PosIdx   int8 `json:"pos_idx"`
	SkillIdx int8 `json:"skill_idx"`
}

func (this *SetSkill_Out) Process(session *net.Session) {
	g_OutHandler.SetSkill(session, this)
}

func (this *SetSkill_Out) TypeName() string {
	return "battle.set_skill.out"
}

func (this *SetSkill_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 6, 19
}

func (this *SetSkill_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type UseItem_In struct {
	IsAttacker bool  `json:"is_attacker"`
	Position   int8  `json:"position"`
	ItemId     int16 `json:"item_id"`
}

func (this *UseItem_In) Process(session *net.Session) {
	g_InHandler.UseItem(session, this)
}

func (this *UseItem_In) TypeName() string {
	return "battle.use_item.in"
}

func (this *UseItem_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 6, 20
}

type UseItem_Out struct {
}

func (this *UseItem_Out) Process(session *net.Session) {
	g_OutHandler.UseItem(session, this)
}

func (this *UseItem_Out) TypeName() string {
	return "battle.use_item.out"
}

func (this *UseItem_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 6, 20
}

func (this *UseItem_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type UseGhost_In struct {
	IsAttacker bool `json:"is_attacker"`
	Position   int8 `json:"position"`
}

func (this *UseGhost_In) Process(session *net.Session) {
	g_InHandler.UseGhost(session, this)
}

func (this *UseGhost_In) TypeName() string {
	return "battle.use_ghost.in"
}

func (this *UseGhost_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 6, 21
}

type UseGhost_Out struct {
}

func (this *UseGhost_Out) Process(session *net.Session) {
	g_OutHandler.UseGhost(session, this)
}

func (this *UseGhost_Out) TypeName() string {
	return "battle.use_ghost.out"
}

func (this *UseGhost_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 6, 21
}

func (this *UseGhost_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type NotifyReady_Out struct {
	Pid int64 `json:"Pid"`
}

func (this *NotifyReady_Out) Process(session *net.Session) {
	g_OutHandler.NotifyReady(session, this)
}

func (this *NotifyReady_Out) TypeName() string {
	return "battle.notify_ready.out"
}

func (this *NotifyReady_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 6, 22
}

func (this *NotifyReady_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type BattleReconnect_In struct {
}

func (this *BattleReconnect_In) Process(session *net.Session) {
	g_InHandler.BattleReconnect(session, this)
}

func (this *BattleReconnect_In) TypeName() string {
	return "battle.battle_reconnect.in"
}

func (this *BattleReconnect_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 6, 23
}

type BattleReconnect_Out struct {
}

func (this *BattleReconnect_Out) Process(session *net.Session) {
	g_OutHandler.BattleReconnect(session, this)
}

func (this *BattleReconnect_Out) TypeName() string {
	return "battle.battle_reconnect.out"
}

func (this *BattleReconnect_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 6, 23
}

func (this *BattleReconnect_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

func (this *StartBattle_In) Decode(buffer *net.Buffer) {
	this.BattleType = BattleType(buffer.ReadUint8())
	this.BattleId = int64(buffer.ReadUint64LE())
}

func (this *StartBattle_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(6)
	buffer.WriteUint8(0)
	buffer.WriteUint8(uint8(this.BattleType))
	buffer.WriteUint64LE(uint64(this.BattleId))
}

func (this *StartBattle_In) ByteSize() int {
	size := 11
	return size
}

func (this *StartBattle_Out) Decode(buffer *net.Buffer) {
	this.TotalGroup = int8(buffer.ReadUint8())
	this.AttackerPlayerIds = make([]StartBattle_Out_AttackerPlayerIds, buffer.ReadUint8())
	for i := 0; i < len(this.AttackerPlayerIds); i++ {
		this.AttackerPlayerIds[i].Decode(buffer)
	}
	this.IsMainRoleFirst = buffer.ReadUint8() == 1
	this.IsAttackerFirst = buffer.ReadUint8() == 1
	this.AllAttackers = make([]StartBattle_Out_AllAttackers, buffer.ReadUint8())
	for i := 0; i < len(this.AllAttackers); i++ {
		this.AllAttackers[i].Decode(buffer)
	}
	this.AllDefenders = make([]StartBattle_Out_AllDefenders, buffer.ReadUint8())
	for i := 0; i < len(this.AllDefenders); i++ {
		this.AllDefenders[i].Decode(buffer)
	}
	this.AttackerTotems = make([]StartBattle_Out_AttackerTotems, buffer.ReadUint8())
	for i := 0; i < len(this.AttackerTotems); i++ {
		this.AttackerTotems[i].Decode(buffer)
	}
	this.DefenderTotems = make([]StartBattle_Out_DefenderTotems, buffer.ReadUint8())
	for i := 0; i < len(this.DefenderTotems); i++ {
		this.DefenderTotems[i].Decode(buffer)
	}
	this.AttackerGroups = make([]StartBattle_Out_AttackerGroups, buffer.ReadUint8())
	for i := 0; i < len(this.AttackerGroups); i++ {
		this.AttackerGroups[i].Decode(buffer)
	}
	this.DefenderGroups = make([]StartBattle_Out_DefenderGroups, buffer.ReadUint8())
	for i := 0; i < len(this.DefenderGroups); i++ {
		this.DefenderGroups[i].Decode(buffer)
	}
}

func (this *StartBattle_Out_AttackerPlayerIds) Decode(buffer *net.Buffer) {
	this.PlayerId = int64(buffer.ReadUint64LE())
}

func (this *StartBattle_Out_AllAttackers) Decode(buffer *net.Buffer) {
	this.PlayerId = int64(buffer.ReadUint64LE())
	this.GhostSkillIndex = int8(buffer.ReadUint8())
}

func (this *StartBattle_Out_AllDefenders) Decode(buffer *net.Buffer) {
	this.PlayerId = int64(buffer.ReadUint64LE())
	this.GhostSkillIndex = int8(buffer.ReadUint8())
}

func (this *StartBattle_Out_AttackerTotems) Decode(buffer *net.Buffer) {
	this.Round = int16(buffer.ReadUint16LE())
	this.TotemId = int16(buffer.ReadUint16LE())
}

func (this *StartBattle_Out_DefenderTotems) Decode(buffer *net.Buffer) {
	this.Round = int16(buffer.ReadUint16LE())
	this.TotemId = int16(buffer.ReadUint16LE())
}

func (this *StartBattle_Out_AttackerGroups) Decode(buffer *net.Buffer) {
	this.Attackers = make([]StartBattle_Out_AttackerGroups_Attackers, buffer.ReadUint8())
	for i := 0; i < len(this.Attackers); i++ {
		this.Attackers[i].Decode(buffer)
	}
	this.SelfBuffs = make([]StartBattle_Out_AttackerGroups_SelfBuffs, buffer.ReadUint8())
	for i := 0; i < len(this.SelfBuffs); i++ {
		this.SelfBuffs[i].Decode(buffer)
	}
	this.BuddyBuffs = make([]StartBattle_Out_AttackerGroups_BuddyBuffs, buffer.ReadUint8())
	for i := 0; i < len(this.BuddyBuffs); i++ {
		this.BuddyBuffs[i].Decode(buffer)
	}
}

func (this *StartBattle_Out_AttackerGroups_Attackers) Decode(buffer *net.Buffer) {
	this.Role.Decode(buffer)
	this.Skills = make([]StartBattle_Out_AttackerGroups_Attackers_Skills, buffer.ReadUint8())
	for i := 0; i < len(this.Skills); i++ {
		this.Skills[i].Decode(buffer)
	}
}

func (this *StartBattle_Out_AttackerGroups_Attackers_Skills) Decode(buffer *net.Buffer) {
	this.Skill.Decode(buffer)
	this.RestReleaseNum = int16(buffer.ReadUint16LE())
}

func (this *StartBattle_Out_AttackerGroups_SelfBuffs) Decode(buffer *net.Buffer) {
	this.Buffer.Decode(buffer)
}

func (this *StartBattle_Out_AttackerGroups_BuddyBuffs) Decode(buffer *net.Buffer) {
	this.Pos = int8(buffer.ReadUint8())
	this.Buffer.Decode(buffer)
}

func (this *StartBattle_Out_DefenderGroups) Decode(buffer *net.Buffer) {
	this.Defenders = make([]StartBattle_Out_DefenderGroups_Defenders, buffer.ReadUint8())
	for i := 0; i < len(this.Defenders); i++ {
		this.Defenders[i].Decode(buffer)
	}
}

func (this *StartBattle_Out_DefenderGroups_Defenders) Decode(buffer *net.Buffer) {
	this.Role.Decode(buffer)
	this.Skills = make([]StartBattle_Out_DefenderGroups_Defenders_Skills, buffer.ReadUint8())
	for i := 0; i < len(this.Skills); i++ {
		this.Skills[i].Decode(buffer)
	}
}

func (this *StartBattle_Out_DefenderGroups_Defenders_Skills) Decode(buffer *net.Buffer) {
	this.Skill.Decode(buffer)
	this.SkillId2 = int16(buffer.ReadUint16LE())
	this.RestReleaseNum = int16(buffer.ReadUint16LE())
}

func (this *StartBattle_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(6)
	buffer.WriteUint8(0)
	buffer.WriteUint8(uint8(this.TotalGroup))
	buffer.WriteUint8(uint8(len(this.AttackerPlayerIds)))
	for i := 0; i < len(this.AttackerPlayerIds); i++ {
		this.AttackerPlayerIds[i].Encode(buffer)
	}
	if this.IsMainRoleFirst {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
	if this.IsAttackerFirst {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
	buffer.WriteUint8(uint8(len(this.AllAttackers)))
	for i := 0; i < len(this.AllAttackers); i++ {
		this.AllAttackers[i].Encode(buffer)
	}
	buffer.WriteUint8(uint8(len(this.AllDefenders)))
	for i := 0; i < len(this.AllDefenders); i++ {
		this.AllDefenders[i].Encode(buffer)
	}
	buffer.WriteUint8(uint8(len(this.AttackerTotems)))
	for i := 0; i < len(this.AttackerTotems); i++ {
		this.AttackerTotems[i].Encode(buffer)
	}
	buffer.WriteUint8(uint8(len(this.DefenderTotems)))
	for i := 0; i < len(this.DefenderTotems); i++ {
		this.DefenderTotems[i].Encode(buffer)
	}
	buffer.WriteUint8(uint8(len(this.AttackerGroups)))
	for i := 0; i < len(this.AttackerGroups); i++ {
		this.AttackerGroups[i].Encode(buffer)
	}
	buffer.WriteUint8(uint8(len(this.DefenderGroups)))
	for i := 0; i < len(this.DefenderGroups); i++ {
		this.DefenderGroups[i].Encode(buffer)
	}
}

func (this *StartBattle_Out_AttackerPlayerIds) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.PlayerId))
}

func (this *StartBattle_Out_AllAttackers) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.PlayerId))
	buffer.WriteUint8(uint8(this.GhostSkillIndex))
}

func (this *StartBattle_Out_AllDefenders) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.PlayerId))
	buffer.WriteUint8(uint8(this.GhostSkillIndex))
}

func (this *StartBattle_Out_AttackerTotems) Encode(buffer *net.Buffer) {
	buffer.WriteUint16LE(uint16(this.Round))
	buffer.WriteUint16LE(uint16(this.TotemId))
}

func (this *StartBattle_Out_DefenderTotems) Encode(buffer *net.Buffer) {
	buffer.WriteUint16LE(uint16(this.Round))
	buffer.WriteUint16LE(uint16(this.TotemId))
}

func (this *StartBattle_Out_AttackerGroups) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(len(this.Attackers)))
	for i := 0; i < len(this.Attackers); i++ {
		this.Attackers[i].Encode(buffer)
	}
	buffer.WriteUint8(uint8(len(this.SelfBuffs)))
	for i := 0; i < len(this.SelfBuffs); i++ {
		this.SelfBuffs[i].Encode(buffer)
	}
	buffer.WriteUint8(uint8(len(this.BuddyBuffs)))
	for i := 0; i < len(this.BuddyBuffs); i++ {
		this.BuddyBuffs[i].Encode(buffer)
	}
}

func (this *StartBattle_Out_AttackerGroups_Attackers) Encode(buffer *net.Buffer) {
	this.Role.Encode(buffer)
	buffer.WriteUint8(uint8(len(this.Skills)))
	for i := 0; i < len(this.Skills); i++ {
		this.Skills[i].Encode(buffer)
	}
}

func (this *StartBattle_Out_AttackerGroups_Attackers_Skills) Encode(buffer *net.Buffer) {
	this.Skill.Encode(buffer)
	buffer.WriteUint16LE(uint16(this.RestReleaseNum))
}

func (this *StartBattle_Out_AttackerGroups_SelfBuffs) Encode(buffer *net.Buffer) {
	this.Buffer.Encode(buffer)
}

func (this *StartBattle_Out_AttackerGroups_BuddyBuffs) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(this.Pos))
	this.Buffer.Encode(buffer)
}

func (this *StartBattle_Out_DefenderGroups) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(len(this.Defenders)))
	for i := 0; i < len(this.Defenders); i++ {
		this.Defenders[i].Encode(buffer)
	}
}

func (this *StartBattle_Out_DefenderGroups_Defenders) Encode(buffer *net.Buffer) {
	this.Role.Encode(buffer)
	buffer.WriteUint8(uint8(len(this.Skills)))
	for i := 0; i < len(this.Skills); i++ {
		this.Skills[i].Encode(buffer)
	}
}

func (this *StartBattle_Out_DefenderGroups_Defenders_Skills) Encode(buffer *net.Buffer) {
	this.Skill.Encode(buffer)
	buffer.WriteUint16LE(uint16(this.SkillId2))
	buffer.WriteUint16LE(uint16(this.RestReleaseNum))
}

func (this *StartBattle_Out) ByteSize() int {
	size := 12
	size += len(this.AttackerPlayerIds) * 8
	size += len(this.AllAttackers) * 9
	size += len(this.AllDefenders) * 9
	size += len(this.AttackerTotems) * 4
	size += len(this.DefenderTotems) * 4
	for i := 0; i < len(this.AttackerGroups); i++ {
		size += this.AttackerGroups[i].ByteSize()
	}
	for i := 0; i < len(this.DefenderGroups); i++ {
		size += this.DefenderGroups[i].ByteSize()
	}
	return size
}

func (this *StartBattle_Out_AttackerGroups) ByteSize() int {
	size := 3
	for i := 0; i < len(this.Attackers); i++ {
		size += this.Attackers[i].ByteSize()
	}
	for i := 0; i < len(this.SelfBuffs); i++ {
		size += this.SelfBuffs[i].ByteSize()
	}
	for i := 0; i < len(this.BuddyBuffs); i++ {
		size += this.BuddyBuffs[i].ByteSize()
	}
	return size
}

func (this *StartBattle_Out_AttackerGroups_Attackers) ByteSize() int {
	size := 1
	size += this.Role.ByteSize()
	for i := 0; i < len(this.Skills); i++ {
		size += this.Skills[i].ByteSize()
	}
	return size
}

func (this *StartBattle_Out_AttackerGroups_Attackers_Skills) ByteSize() int {
	size := 2
	size += this.Skill.ByteSize()
	return size
}

func (this *StartBattle_Out_AttackerGroups_SelfBuffs) ByteSize() int {
	size := 0
	size += this.Buffer.ByteSize()
	return size
}

func (this *StartBattle_Out_AttackerGroups_BuddyBuffs) ByteSize() int {
	size := 1
	size += this.Buffer.ByteSize()
	return size
}

func (this *StartBattle_Out_DefenderGroups) ByteSize() int {
	size := 1
	for i := 0; i < len(this.Defenders); i++ {
		size += this.Defenders[i].ByteSize()
	}
	return size
}

func (this *StartBattle_Out_DefenderGroups_Defenders) ByteSize() int {
	size := 1
	size += this.Role.ByteSize()
	for i := 0; i < len(this.Skills); i++ {
		size += this.Skills[i].ByteSize()
	}
	return size
}

func (this *StartBattle_Out_DefenderGroups_Defenders_Skills) ByteSize() int {
	size := 4
	size += this.Skill.ByteSize()
	return size
}

func (this *NextRound_In) Decode(buffer *net.Buffer) {
	this.UseSkill = int8(buffer.ReadUint8())
	this.UseItem = int16(buffer.ReadUint16LE())
	this.AutoFight = buffer.ReadUint8() == 1
	this.IsAttacker = buffer.ReadUint8() == 1
	this.Position = int8(buffer.ReadUint8())
	this.JobIndex = int8(buffer.ReadUint8())
	this.SendNum = int16(buffer.ReadUint16LE())
	this.UseSwordSoul = buffer.ReadUint8() == 1
	this.UseGhostSkillPosition = int8(buffer.ReadUint8())
	this.UseGhostSkillId = int32(buffer.ReadUint32LE())
	this.UseTotem = buffer.ReadUint8() == 1
}

func (this *NextRound_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(6)
	buffer.WriteUint8(1)
	buffer.WriteUint8(uint8(this.UseSkill))
	buffer.WriteUint16LE(uint16(this.UseItem))
	if this.AutoFight {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
	if this.IsAttacker {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
	buffer.WriteUint8(uint8(this.Position))
	buffer.WriteUint8(uint8(this.JobIndex))
	buffer.WriteUint16LE(uint16(this.SendNum))
	if this.UseSwordSoul {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
	buffer.WriteUint8(uint8(this.UseGhostSkillPosition))
	buffer.WriteUint32LE(uint32(this.UseGhostSkillId))
	if this.UseTotem {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *NextRound_In) ByteSize() int {
	size := 18
	return size
}

func (this *NextRound_Out) Decode(buffer *net.Buffer) {
	this.Status = RoundStatus(buffer.ReadUint8())
	this.NowRound = int16(buffer.ReadUint16LE())
	this.AllAttackers = make([]NextRound_Out_AllAttackers, buffer.ReadUint8())
	for i := 0; i < len(this.AllAttackers); i++ {
		this.AllAttackers[i].Decode(buffer)
	}
	this.AllDefenders = make([]NextRound_Out_AllDefenders, buffer.ReadUint8())
	for i := 0; i < len(this.AllDefenders); i++ {
		this.AllDefenders[i].Decode(buffer)
	}
	this.Results = make([]NextRound_Out_Results, buffer.ReadUint8())
	for i := 0; i < len(this.Results); i++ {
		this.Results[i].Decode(buffer)
	}
	this.Autos = make([]NextRound_Out_Autos, buffer.ReadUint8())
	for i := 0; i < len(this.Autos); i++ {
		this.Autos[i].Decode(buffer)
	}
}

func (this *NextRound_Out_AllAttackers) Decode(buffer *net.Buffer) {
	this.PlayerId = int64(buffer.ReadUint64LE())
}

func (this *NextRound_Out_AllDefenders) Decode(buffer *net.Buffer) {
	this.PlayerId = int64(buffer.ReadUint64LE())
}

func (this *NextRound_Out_Results) Decode(buffer *net.Buffer) {
	this.Ftype = FighterType(buffer.ReadUint8())
	this.Event = RoundEvent(buffer.ReadUint8())
	this.Position = int8(buffer.ReadUint8())
	this.Power = int16(buffer.ReadUint16LE())
	this.Health = int32(buffer.ReadUint32LE())
	this.SunderValue = int16(buffer.ReadUint16LE())
	this.UseGhostSkill = buffer.ReadUint8() == 1
	this.TotemId = int16(buffer.ReadUint16LE())
	this.GhostId = int16(buffer.ReadUint16LE())
	this.GhostShieldOn = buffer.ReadUint8() == 1
	this.ShieldGhostId = int16(buffer.ReadUint16LE())
	this.GhostPower = int32(buffer.ReadUint32LE())
	this.AddPower = int32(buffer.ReadUint32LE())
	this.Attacks = make([]NextRound_Out_Results_Attacks, buffer.ReadUint8())
	for i := 0; i < len(this.Attacks); i++ {
		this.Attacks[i].Decode(buffer)
	}
	this.Item = make([]NextRound_Out_Results_Item, buffer.ReadUint8())
	for i := 0; i < len(this.Item); i++ {
		this.Item[i].Decode(buffer)
	}
}

func (this *NextRound_Out_Results_Attacks) Decode(buffer *net.Buffer) {
	this.SkillId = int32(buffer.ReadUint32LE())
	this.RestReleaseNum = int16(buffer.ReadUint16LE())
	this.Targets = make([]NextRound_Out_Results_Attacks_Targets, buffer.ReadUint8())
	for i := 0; i < len(this.Targets); i++ {
		this.Targets[i].Decode(buffer)
	}
	this.SelfBuffs = make([]NextRound_Out_Results_Attacks_SelfBuffs, buffer.ReadUint8())
	for i := 0; i < len(this.SelfBuffs); i++ {
		this.SelfBuffs[i].Decode(buffer)
	}
	this.BuddyBuffs = make([]NextRound_Out_Results_Attacks_BuddyBuffs, buffer.ReadUint8())
	for i := 0; i < len(this.BuddyBuffs); i++ {
		this.BuddyBuffs[i].Decode(buffer)
	}
}

func (this *NextRound_Out_Results_Attacks_Targets) Decode(buffer *net.Buffer) {
	this.Ftype = FighterType(buffer.ReadUint8())
	this.Hurt = int32(buffer.ReadUint32LE())
	this.Event = RoundEvent(buffer.ReadUint8())
	this.Position = int8(buffer.ReadUint8())
	this.TakeSunder = int16(buffer.ReadUint16LE())
	this.TakeGhostShield = int32(buffer.ReadUint32LE())
	this.DirectReductHurt = int32(buffer.ReadUint32LE())
	this.GhostShieldOn = buffer.ReadUint8() == 1
	this.ShieldGhostId = int16(buffer.ReadUint16LE())
	this.GhostPower = int32(buffer.ReadUint32LE())
	this.Buffs = make([]NextRound_Out_Results_Attacks_Targets_Buffs, buffer.ReadUint8())
	for i := 0; i < len(this.Buffs); i++ {
		this.Buffs[i].Decode(buffer)
	}
	this.PassiveAttack = make([]NextRound_Out_Results_Attacks_Targets_PassiveAttack, buffer.ReadUint8())
	for i := 0; i < len(this.PassiveAttack); i++ {
		this.PassiveAttack[i].Decode(buffer)
	}
}

func (this *NextRound_Out_Results_Attacks_Targets_Buffs) Decode(buffer *net.Buffer) {
	this.Buffer.Decode(buffer)
}

func (this *NextRound_Out_Results_Attacks_Targets_PassiveAttack) Decode(buffer *net.Buffer) {
	this.SkillId = int32(buffer.ReadUint32LE())
	this.TargetsBuffs = make([]NextRound_Out_Results_Attacks_Targets_PassiveAttack_TargetsBuffs, buffer.ReadUint8())
	for i := 0; i < len(this.TargetsBuffs); i++ {
		this.TargetsBuffs[i].Decode(buffer)
	}
	this.TeamBuffs = make([]NextRound_Out_Results_Attacks_Targets_PassiveAttack_TeamBuffs, buffer.ReadUint8())
	for i := 0; i < len(this.TeamBuffs); i++ {
		this.TeamBuffs[i].Decode(buffer)
	}
}

func (this *NextRound_Out_Results_Attacks_Targets_PassiveAttack_TargetsBuffs) Decode(buffer *net.Buffer) {
	this.Pos = int8(buffer.ReadUint8())
	this.Buffer.Decode(buffer)
}

func (this *NextRound_Out_Results_Attacks_Targets_PassiveAttack_TeamBuffs) Decode(buffer *net.Buffer) {
	this.Pos = int8(buffer.ReadUint8())
	this.Buffer.Decode(buffer)
}

func (this *NextRound_Out_Results_Attacks_SelfBuffs) Decode(buffer *net.Buffer) {
	this.Buffer.Decode(buffer)
}

func (this *NextRound_Out_Results_Attacks_BuddyBuffs) Decode(buffer *net.Buffer) {
	this.Pos = int8(buffer.ReadUint8())
	this.Buffer.Decode(buffer)
}

func (this *NextRound_Out_Results_Item) Decode(buffer *net.Buffer) {
	this.ItemId = int32(buffer.ReadUint32LE())
	this.Targets = make([]NextRound_Out_Results_Item_Targets, buffer.ReadUint8())
	for i := 0; i < len(this.Targets); i++ {
		this.Targets[i].Decode(buffer)
	}
}

func (this *NextRound_Out_Results_Item_Targets) Decode(buffer *net.Buffer) {
	this.Ftype = FighterType(buffer.ReadUint8())
	this.Health = int32(buffer.ReadUint32LE())
	this.Power = int16(buffer.ReadUint16LE())
	this.Hurt = int32(buffer.ReadUint32LE())
	this.Position = int8(buffer.ReadUint8())
	this.Buffs = make([]NextRound_Out_Results_Item_Targets_Buffs, buffer.ReadUint8())
	for i := 0; i < len(this.Buffs); i++ {
		this.Buffs[i].Decode(buffer)
	}
}

func (this *NextRound_Out_Results_Item_Targets_Buffs) Decode(buffer *net.Buffer) {
	this.Buffer.Decode(buffer)
}

func (this *NextRound_Out_Autos) Decode(buffer *net.Buffer) {
	this.PlayerId = int64(buffer.ReadUint64LE())
}

func (this *NextRound_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(6)
	buffer.WriteUint8(1)
	buffer.WriteUint8(uint8(this.Status))
	buffer.WriteUint16LE(uint16(this.NowRound))
	buffer.WriteUint8(uint8(len(this.AllAttackers)))
	for i := 0; i < len(this.AllAttackers); i++ {
		this.AllAttackers[i].Encode(buffer)
	}
	buffer.WriteUint8(uint8(len(this.AllDefenders)))
	for i := 0; i < len(this.AllDefenders); i++ {
		this.AllDefenders[i].Encode(buffer)
	}
	buffer.WriteUint8(uint8(len(this.Results)))
	for i := 0; i < len(this.Results); i++ {
		this.Results[i].Encode(buffer)
	}
	buffer.WriteUint8(uint8(len(this.Autos)))
	for i := 0; i < len(this.Autos); i++ {
		this.Autos[i].Encode(buffer)
	}
}

func (this *NextRound_Out_AllAttackers) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.PlayerId))
}

func (this *NextRound_Out_AllDefenders) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.PlayerId))
}

func (this *NextRound_Out_Results) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(this.Ftype))
	buffer.WriteUint8(uint8(this.Event))
	buffer.WriteUint8(uint8(this.Position))
	buffer.WriteUint16LE(uint16(this.Power))
	buffer.WriteUint32LE(uint32(this.Health))
	buffer.WriteUint16LE(uint16(this.SunderValue))
	if this.UseGhostSkill {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
	buffer.WriteUint16LE(uint16(this.TotemId))
	buffer.WriteUint16LE(uint16(this.GhostId))
	if this.GhostShieldOn {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
	buffer.WriteUint16LE(uint16(this.ShieldGhostId))
	buffer.WriteUint32LE(uint32(this.GhostPower))
	buffer.WriteUint32LE(uint32(this.AddPower))
	buffer.WriteUint8(uint8(len(this.Attacks)))
	for i := 0; i < len(this.Attacks); i++ {
		this.Attacks[i].Encode(buffer)
	}
	buffer.WriteUint8(uint8(len(this.Item)))
	for i := 0; i < len(this.Item); i++ {
		this.Item[i].Encode(buffer)
	}
}

func (this *NextRound_Out_Results_Attacks) Encode(buffer *net.Buffer) {
	buffer.WriteUint32LE(uint32(this.SkillId))
	buffer.WriteUint16LE(uint16(this.RestReleaseNum))
	buffer.WriteUint8(uint8(len(this.Targets)))
	for i := 0; i < len(this.Targets); i++ {
		this.Targets[i].Encode(buffer)
	}
	buffer.WriteUint8(uint8(len(this.SelfBuffs)))
	for i := 0; i < len(this.SelfBuffs); i++ {
		this.SelfBuffs[i].Encode(buffer)
	}
	buffer.WriteUint8(uint8(len(this.BuddyBuffs)))
	for i := 0; i < len(this.BuddyBuffs); i++ {
		this.BuddyBuffs[i].Encode(buffer)
	}
}

func (this *NextRound_Out_Results_Attacks_Targets) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(this.Ftype))
	buffer.WriteUint32LE(uint32(this.Hurt))
	buffer.WriteUint8(uint8(this.Event))
	buffer.WriteUint8(uint8(this.Position))
	buffer.WriteUint16LE(uint16(this.TakeSunder))
	buffer.WriteUint32LE(uint32(this.TakeGhostShield))
	buffer.WriteUint32LE(uint32(this.DirectReductHurt))
	if this.GhostShieldOn {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
	buffer.WriteUint16LE(uint16(this.ShieldGhostId))
	buffer.WriteUint32LE(uint32(this.GhostPower))
	buffer.WriteUint8(uint8(len(this.Buffs)))
	for i := 0; i < len(this.Buffs); i++ {
		this.Buffs[i].Encode(buffer)
	}
	buffer.WriteUint8(uint8(len(this.PassiveAttack)))
	for i := 0; i < len(this.PassiveAttack); i++ {
		this.PassiveAttack[i].Encode(buffer)
	}
}

func (this *NextRound_Out_Results_Attacks_Targets_Buffs) Encode(buffer *net.Buffer) {
	this.Buffer.Encode(buffer)
}

func (this *NextRound_Out_Results_Attacks_Targets_PassiveAttack) Encode(buffer *net.Buffer) {
	buffer.WriteUint32LE(uint32(this.SkillId))
	buffer.WriteUint8(uint8(len(this.TargetsBuffs)))
	for i := 0; i < len(this.TargetsBuffs); i++ {
		this.TargetsBuffs[i].Encode(buffer)
	}
	buffer.WriteUint8(uint8(len(this.TeamBuffs)))
	for i := 0; i < len(this.TeamBuffs); i++ {
		this.TeamBuffs[i].Encode(buffer)
	}
}

func (this *NextRound_Out_Results_Attacks_Targets_PassiveAttack_TargetsBuffs) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(this.Pos))
	this.Buffer.Encode(buffer)
}

func (this *NextRound_Out_Results_Attacks_Targets_PassiveAttack_TeamBuffs) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(this.Pos))
	this.Buffer.Encode(buffer)
}

func (this *NextRound_Out_Results_Attacks_SelfBuffs) Encode(buffer *net.Buffer) {
	this.Buffer.Encode(buffer)
}

func (this *NextRound_Out_Results_Attacks_BuddyBuffs) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(this.Pos))
	this.Buffer.Encode(buffer)
}

func (this *NextRound_Out_Results_Item) Encode(buffer *net.Buffer) {
	buffer.WriteUint32LE(uint32(this.ItemId))
	buffer.WriteUint8(uint8(len(this.Targets)))
	for i := 0; i < len(this.Targets); i++ {
		this.Targets[i].Encode(buffer)
	}
}

func (this *NextRound_Out_Results_Item_Targets) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(this.Ftype))
	buffer.WriteUint32LE(uint32(this.Health))
	buffer.WriteUint16LE(uint16(this.Power))
	buffer.WriteUint32LE(uint32(this.Hurt))
	buffer.WriteUint8(uint8(this.Position))
	buffer.WriteUint8(uint8(len(this.Buffs)))
	for i := 0; i < len(this.Buffs); i++ {
		this.Buffs[i].Encode(buffer)
	}
}

func (this *NextRound_Out_Results_Item_Targets_Buffs) Encode(buffer *net.Buffer) {
	this.Buffer.Encode(buffer)
}

func (this *NextRound_Out_Autos) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.PlayerId))
}

func (this *NextRound_Out) ByteSize() int {
	size := 9
	size += len(this.AllAttackers) * 8
	size += len(this.AllDefenders) * 8
	for i := 0; i < len(this.Results); i++ {
		size += this.Results[i].ByteSize()
	}
	size += len(this.Autos) * 8
	return size
}

func (this *NextRound_Out_Results) ByteSize() int {
	size := 29
	for i := 0; i < len(this.Attacks); i++ {
		size += this.Attacks[i].ByteSize()
	}
	for i := 0; i < len(this.Item); i++ {
		size += this.Item[i].ByteSize()
	}
	return size
}

func (this *NextRound_Out_Results_Attacks) ByteSize() int {
	size := 9
	for i := 0; i < len(this.Targets); i++ {
		size += this.Targets[i].ByteSize()
	}
	for i := 0; i < len(this.SelfBuffs); i++ {
		size += this.SelfBuffs[i].ByteSize()
	}
	for i := 0; i < len(this.BuddyBuffs); i++ {
		size += this.BuddyBuffs[i].ByteSize()
	}
	return size
}

func (this *NextRound_Out_Results_Attacks_Targets) ByteSize() int {
	size := 26
	for i := 0; i < len(this.Buffs); i++ {
		size += this.Buffs[i].ByteSize()
	}
	for i := 0; i < len(this.PassiveAttack); i++ {
		size += this.PassiveAttack[i].ByteSize()
	}
	return size
}

func (this *NextRound_Out_Results_Attacks_Targets_Buffs) ByteSize() int {
	size := 0
	size += this.Buffer.ByteSize()
	return size
}

func (this *NextRound_Out_Results_Attacks_Targets_PassiveAttack) ByteSize() int {
	size := 6
	for i := 0; i < len(this.TargetsBuffs); i++ {
		size += this.TargetsBuffs[i].ByteSize()
	}
	for i := 0; i < len(this.TeamBuffs); i++ {
		size += this.TeamBuffs[i].ByteSize()
	}
	return size
}

func (this *NextRound_Out_Results_Attacks_Targets_PassiveAttack_TargetsBuffs) ByteSize() int {
	size := 1
	size += this.Buffer.ByteSize()
	return size
}

func (this *NextRound_Out_Results_Attacks_Targets_PassiveAttack_TeamBuffs) ByteSize() int {
	size := 1
	size += this.Buffer.ByteSize()
	return size
}

func (this *NextRound_Out_Results_Attacks_SelfBuffs) ByteSize() int {
	size := 0
	size += this.Buffer.ByteSize()
	return size
}

func (this *NextRound_Out_Results_Attacks_BuddyBuffs) ByteSize() int {
	size := 1
	size += this.Buffer.ByteSize()
	return size
}

func (this *NextRound_Out_Results_Item) ByteSize() int {
	size := 5
	for i := 0; i < len(this.Targets); i++ {
		size += this.Targets[i].ByteSize()
	}
	return size
}

func (this *NextRound_Out_Results_Item_Targets) ByteSize() int {
	size := 13
	for i := 0; i < len(this.Buffs); i++ {
		size += this.Buffs[i].ByteSize()
	}
	return size
}

func (this *NextRound_Out_Results_Item_Targets_Buffs) ByteSize() int {
	size := 0
	size += this.Buffer.ByteSize()
	return size
}

func (this *End_Out) Decode(buffer *net.Buffer) {
	this.Status = RoundStatus(buffer.ReadUint8())
}

func (this *End_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(6)
	buffer.WriteUint8(2)
	buffer.WriteUint8(uint8(this.Status))
}

func (this *End_Out) ByteSize() int {
	size := 3
	return size
}

func (this *Escape_In) Decode(buffer *net.Buffer) {
}

func (this *Escape_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(6)
	buffer.WriteUint8(3)
}

func (this *Escape_In) ByteSize() int {
	size := 2
	return size
}

func (this *Escape_Out) Decode(buffer *net.Buffer) {
}

func (this *Escape_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(6)
	buffer.WriteUint8(3)
}

func (this *Escape_Out) ByteSize() int {
	size := 2
	return size
}

func (this *Fightnum_Out) Decode(buffer *net.Buffer) {
	this.Attacker = int32(buffer.ReadUint32LE())
	this.Defender = int32(buffer.ReadUint32LE())
}

func (this *Fightnum_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(6)
	buffer.WriteUint8(4)
	buffer.WriteUint32LE(uint32(this.Attacker))
	buffer.WriteUint32LE(uint32(this.Defender))
}

func (this *Fightnum_Out) ByteSize() int {
	size := 10
	return size
}

func (this *StartReadyTimeout_Out) Decode(buffer *net.Buffer) {
}

func (this *StartReadyTimeout_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(6)
	buffer.WriteUint8(5)
}

func (this *StartReadyTimeout_Out) ByteSize() int {
	size := 2
	return size
}

func (this *StartReady_In) Decode(buffer *net.Buffer) {
	this.Ok = buffer.ReadUint8() == 1
}

func (this *StartReady_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(6)
	buffer.WriteUint8(6)
	if this.Ok {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *StartReady_In) ByteSize() int {
	size := 3
	return size
}

func (this *StartReady_Out) Decode(buffer *net.Buffer) {
	this.ReadyPid = int64(buffer.ReadUint64LE())
}

func (this *StartReady_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(6)
	buffer.WriteUint8(6)
	buffer.WriteUint64LE(uint64(this.ReadyPid))
}

func (this *StartReady_Out) ByteSize() int {
	size := 10
	return size
}

func (this *StateChange_Out) Decode(buffer *net.Buffer) {
	this.PlayerId = int64(buffer.ReadUint64LE())
	this.Auto = buffer.ReadUint8() == 1
	this.Desc = buffer.ReadBytes(int(buffer.ReadUint16LE()))
}

func (this *StateChange_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(6)
	buffer.WriteUint8(7)
	buffer.WriteUint64LE(uint64(this.PlayerId))
	if this.Auto {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
	buffer.WriteUint16LE(uint16(len(this.Desc)))
	buffer.WriteBytes(this.Desc)
}

func (this *StateChange_Out) ByteSize() int {
	size := 13
	size += len(this.Desc)
	return size
}

func (this *CallBattlePet_In) Decode(buffer *net.Buffer) {
	this.GridNum = int8(buffer.ReadUint8())
}

func (this *CallBattlePet_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(6)
	buffer.WriteUint8(8)
	buffer.WriteUint8(uint8(this.GridNum))
}

func (this *CallBattlePet_In) ByteSize() int {
	size := 3
	return size
}

func (this *CallBattlePet_Out) Decode(buffer *net.Buffer) {
	this.Success = buffer.ReadUint8() == 1
	this.PlayerPower = int16(buffer.ReadUint16LE())
	this.Role.Decode(buffer)
	this.Skills = make([]CallBattlePet_Out_Skills, buffer.ReadUint8())
	for i := 0; i < len(this.Skills); i++ {
		this.Skills[i].Decode(buffer)
	}
}

func (this *CallBattlePet_Out_Skills) Decode(buffer *net.Buffer) {
	this.Skill.Decode(buffer)
}

func (this *CallBattlePet_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(6)
	buffer.WriteUint8(8)
	if this.Success {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
	buffer.WriteUint16LE(uint16(this.PlayerPower))
	this.Role.Encode(buffer)
	buffer.WriteUint8(uint8(len(this.Skills)))
	for i := 0; i < len(this.Skills); i++ {
		this.Skills[i].Encode(buffer)
	}
}

func (this *CallBattlePet_Out_Skills) Encode(buffer *net.Buffer) {
	this.Skill.Encode(buffer)
}

func (this *CallBattlePet_Out) ByteSize() int {
	size := 6
	size += this.Role.ByteSize()
	for i := 0; i < len(this.Skills); i++ {
		size += this.Skills[i].ByteSize()
	}
	return size
}

func (this *CallBattlePet_Out_Skills) ByteSize() int {
	size := 0
	size += this.Skill.ByteSize()
	return size
}

func (this *UseBuddySkill_In) Decode(buffer *net.Buffer) {
	this.Pos = int8(buffer.ReadUint8())
	this.UseSkill = int8(buffer.ReadUint8())
}

func (this *UseBuddySkill_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(6)
	buffer.WriteUint8(9)
	buffer.WriteUint8(uint8(this.Pos))
	buffer.WriteUint8(uint8(this.UseSkill))
}

func (this *UseBuddySkill_In) ByteSize() int {
	size := 4
	return size
}

func (this *UseBuddySkill_Out) Decode(buffer *net.Buffer) {
	this.Pos = int8(buffer.ReadUint8())
	this.UseSkill = int8(buffer.ReadUint8())
}

func (this *UseBuddySkill_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(6)
	buffer.WriteUint8(9)
	buffer.WriteUint8(uint8(this.Pos))
	buffer.WriteUint8(uint8(this.UseSkill))
}

func (this *UseBuddySkill_Out) ByteSize() int {
	size := 4
	return size
}

func (this *CallNewEnemys_Out) Decode(buffer *net.Buffer) {
	this.CallInfo = make([]CallNewEnemys_Out_CallInfo, buffer.ReadUint8())
	for i := 0; i < len(this.CallInfo); i++ {
		this.CallInfo[i].Decode(buffer)
	}
}

func (this *CallNewEnemys_Out_CallInfo) Decode(buffer *net.Buffer) {
	this.Ftype = int8(buffer.ReadUint8())
	this.Position = int8(buffer.ReadUint8())
	this.AttackIndex = int8(buffer.ReadUint8())
	this.Enemys = make([]CallNewEnemys_Out_CallInfo_Enemys, buffer.ReadUint8())
	for i := 0; i < len(this.Enemys); i++ {
		this.Enemys[i].Decode(buffer)
	}
}

func (this *CallNewEnemys_Out_CallInfo_Enemys) Decode(buffer *net.Buffer) {
	this.Role.Decode(buffer)
	this.Skills = make([]CallNewEnemys_Out_CallInfo_Enemys_Skills, buffer.ReadUint8())
	for i := 0; i < len(this.Skills); i++ {
		this.Skills[i].Decode(buffer)
	}
}

func (this *CallNewEnemys_Out_CallInfo_Enemys_Skills) Decode(buffer *net.Buffer) {
	this.Skill.Decode(buffer)
	this.SkillId2 = int16(buffer.ReadUint16LE())
	this.RestReleaseNum = int16(buffer.ReadUint16LE())
}

func (this *CallNewEnemys_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(6)
	buffer.WriteUint8(10)
	buffer.WriteUint8(uint8(len(this.CallInfo)))
	for i := 0; i < len(this.CallInfo); i++ {
		this.CallInfo[i].Encode(buffer)
	}
}

func (this *CallNewEnemys_Out_CallInfo) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(this.Ftype))
	buffer.WriteUint8(uint8(this.Position))
	buffer.WriteUint8(uint8(this.AttackIndex))
	buffer.WriteUint8(uint8(len(this.Enemys)))
	for i := 0; i < len(this.Enemys); i++ {
		this.Enemys[i].Encode(buffer)
	}
}

func (this *CallNewEnemys_Out_CallInfo_Enemys) Encode(buffer *net.Buffer) {
	this.Role.Encode(buffer)
	buffer.WriteUint8(uint8(len(this.Skills)))
	for i := 0; i < len(this.Skills); i++ {
		this.Skills[i].Encode(buffer)
	}
}

func (this *CallNewEnemys_Out_CallInfo_Enemys_Skills) Encode(buffer *net.Buffer) {
	this.Skill.Encode(buffer)
	buffer.WriteUint16LE(uint16(this.SkillId2))
	buffer.WriteUint16LE(uint16(this.RestReleaseNum))
}

func (this *CallNewEnemys_Out) ByteSize() int {
	size := 3
	for i := 0; i < len(this.CallInfo); i++ {
		size += this.CallInfo[i].ByteSize()
	}
	return size
}

func (this *CallNewEnemys_Out_CallInfo) ByteSize() int {
	size := 4
	for i := 0; i < len(this.Enemys); i++ {
		size += this.Enemys[i].ByteSize()
	}
	return size
}

func (this *CallNewEnemys_Out_CallInfo_Enemys) ByteSize() int {
	size := 1
	size += this.Role.ByteSize()
	for i := 0; i < len(this.Skills); i++ {
		size += this.Skills[i].ByteSize()
	}
	return size
}

func (this *CallNewEnemys_Out_CallInfo_Enemys_Skills) ByteSize() int {
	size := 4
	size += this.Skill.ByteSize()
	return size
}

func (this *NewFighterGroup_Out) Decode(buffer *net.Buffer) {
	this.Ftype = int8(buffer.ReadUint8())
	this.PlayerId = int64(buffer.ReadUint64LE())
	this.GhostSkillIndex = int8(buffer.ReadUint8())
	this.Fighters = make([]NewFighterGroup_Out_Fighters, buffer.ReadUint8())
	for i := 0; i < len(this.Fighters); i++ {
		this.Fighters[i].Decode(buffer)
	}
}

func (this *NewFighterGroup_Out_Fighters) Decode(buffer *net.Buffer) {
	this.Role.Decode(buffer)
	this.Skills = make([]NewFighterGroup_Out_Fighters_Skills, buffer.ReadUint8())
	for i := 0; i < len(this.Skills); i++ {
		this.Skills[i].Decode(buffer)
	}
}

func (this *NewFighterGroup_Out_Fighters_Skills) Decode(buffer *net.Buffer) {
	this.Skill.Decode(buffer)
	this.RestReleaseNum = int16(buffer.ReadUint16LE())
}

func (this *NewFighterGroup_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(6)
	buffer.WriteUint8(11)
	buffer.WriteUint8(uint8(this.Ftype))
	buffer.WriteUint64LE(uint64(this.PlayerId))
	buffer.WriteUint8(uint8(this.GhostSkillIndex))
	buffer.WriteUint8(uint8(len(this.Fighters)))
	for i := 0; i < len(this.Fighters); i++ {
		this.Fighters[i].Encode(buffer)
	}
}

func (this *NewFighterGroup_Out_Fighters) Encode(buffer *net.Buffer) {
	this.Role.Encode(buffer)
	buffer.WriteUint8(uint8(len(this.Skills)))
	for i := 0; i < len(this.Skills); i++ {
		this.Skills[i].Encode(buffer)
	}
}

func (this *NewFighterGroup_Out_Fighters_Skills) Encode(buffer *net.Buffer) {
	this.Skill.Encode(buffer)
	buffer.WriteUint16LE(uint16(this.RestReleaseNum))
}

func (this *NewFighterGroup_Out) ByteSize() int {
	size := 13
	for i := 0; i < len(this.Fighters); i++ {
		size += this.Fighters[i].ByteSize()
	}
	return size
}

func (this *NewFighterGroup_Out_Fighters) ByteSize() int {
	size := 1
	size += this.Role.ByteSize()
	for i := 0; i < len(this.Skills); i++ {
		size += this.Skills[i].ByteSize()
	}
	return size
}

func (this *NewFighterGroup_Out_Fighters_Skills) ByteSize() int {
	size := 2
	size += this.Skill.ByteSize()
	return size
}

func (this *StartBattleForHijackBoat_In) Decode(buffer *net.Buffer) {
	this.Pid = int64(buffer.ReadUint64LE())
	this.BoatId = int64(buffer.ReadUint64LE())
}

func (this *StartBattleForHijackBoat_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(6)
	buffer.WriteUint8(12)
	buffer.WriteUint64LE(uint64(this.Pid))
	buffer.WriteUint64LE(uint64(this.BoatId))
}

func (this *StartBattleForHijackBoat_In) ByteSize() int {
	size := 18
	return size
}

func (this *StartBattleForHijackBoat_Out) Decode(buffer *net.Buffer) {
}

func (this *StartBattleForHijackBoat_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(6)
	buffer.WriteUint8(12)
}

func (this *StartBattleForHijackBoat_Out) ByteSize() int {
	size := 2
	return size
}

func (this *StartBattleForRecoverBoat_In) Decode(buffer *net.Buffer) {
	this.Pid = int64(buffer.ReadUint64LE())
	this.BoatId = int64(buffer.ReadUint64LE())
}

func (this *StartBattleForRecoverBoat_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(6)
	buffer.WriteUint8(13)
	buffer.WriteUint64LE(uint64(this.Pid))
	buffer.WriteUint64LE(uint64(this.BoatId))
}

func (this *StartBattleForRecoverBoat_In) ByteSize() int {
	size := 18
	return size
}

func (this *StartBattleForRecoverBoat_Out) Decode(buffer *net.Buffer) {
}

func (this *StartBattleForRecoverBoat_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(6)
	buffer.WriteUint8(13)
}

func (this *StartBattleForRecoverBoat_Out) ByteSize() int {
	size := 2
	return size
}

func (this *RoundReady_In) Decode(buffer *net.Buffer) {
	this.IsAuto = buffer.ReadUint8() == 1
}

func (this *RoundReady_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(6)
	buffer.WriteUint8(14)
	if this.IsAuto {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *RoundReady_In) ByteSize() int {
	size := 3
	return size
}

func (this *RoundReady_Out) Decode(buffer *net.Buffer) {
}

func (this *RoundReady_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(6)
	buffer.WriteUint8(14)
}

func (this *RoundReady_Out) ByteSize() int {
	size := 2
	return size
}

func (this *InitRound_In) Decode(buffer *net.Buffer) {
}

func (this *InitRound_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(6)
	buffer.WriteUint8(15)
}

func (this *InitRound_In) ByteSize() int {
	size := 2
	return size
}

func (this *InitRound_Out) Decode(buffer *net.Buffer) {
}

func (this *InitRound_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(6)
	buffer.WriteUint8(15)
}

func (this *InitRound_Out) ByteSize() int {
	size := 2
	return size
}

func (this *SetAuto_In) Decode(buffer *net.Buffer) {
}

func (this *SetAuto_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(6)
	buffer.WriteUint8(17)
}

func (this *SetAuto_In) ByteSize() int {
	size := 2
	return size
}

func (this *SetAuto_Out) Decode(buffer *net.Buffer) {
}

func (this *SetAuto_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(6)
	buffer.WriteUint8(17)
}

func (this *SetAuto_Out) ByteSize() int {
	size := 2
	return size
}

func (this *CancelAuto_In) Decode(buffer *net.Buffer) {
}

func (this *CancelAuto_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(6)
	buffer.WriteUint8(18)
}

func (this *CancelAuto_In) ByteSize() int {
	size := 2
	return size
}

func (this *CancelAuto_Out) Decode(buffer *net.Buffer) {
	this.Round = int16(buffer.ReadUint16LE())
}

func (this *CancelAuto_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(6)
	buffer.WriteUint8(18)
	buffer.WriteUint16LE(uint16(this.Round))
}

func (this *CancelAuto_Out) ByteSize() int {
	size := 4
	return size
}

func (this *SetSkill_In) Decode(buffer *net.Buffer) {
	this.IsAttacker = buffer.ReadUint8() == 1
	this.PosIdx = int8(buffer.ReadUint8())
	this.SkillIdx = int8(buffer.ReadUint8())
}

func (this *SetSkill_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(6)
	buffer.WriteUint8(19)
	if this.IsAttacker {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
	buffer.WriteUint8(uint8(this.PosIdx))
	buffer.WriteUint8(uint8(this.SkillIdx))
}

func (this *SetSkill_In) ByteSize() int {
	size := 5
	return size
}

func (this *SetSkill_Out) Decode(buffer *net.Buffer) {
	this.PosIdx = int8(buffer.ReadUint8())
	this.SkillIdx = int8(buffer.ReadUint8())
}

func (this *SetSkill_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(6)
	buffer.WriteUint8(19)
	buffer.WriteUint8(uint8(this.PosIdx))
	buffer.WriteUint8(uint8(this.SkillIdx))
}

func (this *SetSkill_Out) ByteSize() int {
	size := 4
	return size
}

func (this *UseItem_In) Decode(buffer *net.Buffer) {
	this.IsAttacker = buffer.ReadUint8() == 1
	this.Position = int8(buffer.ReadUint8())
	this.ItemId = int16(buffer.ReadUint16LE())
}

func (this *UseItem_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(6)
	buffer.WriteUint8(20)
	if this.IsAttacker {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
	buffer.WriteUint8(uint8(this.Position))
	buffer.WriteUint16LE(uint16(this.ItemId))
}

func (this *UseItem_In) ByteSize() int {
	size := 6
	return size
}

func (this *UseItem_Out) Decode(buffer *net.Buffer) {
}

func (this *UseItem_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(6)
	buffer.WriteUint8(20)
}

func (this *UseItem_Out) ByteSize() int {
	size := 2
	return size
}

func (this *UseGhost_In) Decode(buffer *net.Buffer) {
	this.IsAttacker = buffer.ReadUint8() == 1
	this.Position = int8(buffer.ReadUint8())
}

func (this *UseGhost_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(6)
	buffer.WriteUint8(21)
	if this.IsAttacker {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
	buffer.WriteUint8(uint8(this.Position))
}

func (this *UseGhost_In) ByteSize() int {
	size := 4
	return size
}

func (this *UseGhost_Out) Decode(buffer *net.Buffer) {
}

func (this *UseGhost_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(6)
	buffer.WriteUint8(21)
}

func (this *UseGhost_Out) ByteSize() int {
	size := 2
	return size
}

func (this *NotifyReady_Out) Decode(buffer *net.Buffer) {
	this.Pid = int64(buffer.ReadUint64LE())
}

func (this *NotifyReady_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(6)
	buffer.WriteUint8(22)
	buffer.WriteUint64LE(uint64(this.Pid))
}

func (this *NotifyReady_Out) ByteSize() int {
	size := 10
	return size
}

func (this *BattleReconnect_In) Decode(buffer *net.Buffer) {
}

func (this *BattleReconnect_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(6)
	buffer.WriteUint8(23)
}

func (this *BattleReconnect_In) ByteSize() int {
	size := 2
	return size
}

func (this *BattleReconnect_Out) Decode(buffer *net.Buffer) {
}

func (this *BattleReconnect_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(6)
	buffer.WriteUint8(23)
}

func (this *BattleReconnect_Out) ByteSize() int {
	size := 2
	return size
}
