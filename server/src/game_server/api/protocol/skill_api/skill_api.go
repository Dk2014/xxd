package skill_api

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
	GetAllSkillsInfo(*net.Session, *GetAllSkillsInfo_In)
	EquipSkill(*net.Session, *EquipSkill_In)
	UnequipSkill(*net.Session, *UnequipSkill_In)
	StudySkillByCheat(*net.Session, *StudySkillByCheat_In)
	TrainSkill(*net.Session, *TrainSkill_In)
	FlushSkill(*net.Session, *FlushSkill_In)
}

type OutHandler interface {
	GetAllSkillsInfo(*net.Session, *GetAllSkillsInfo_Out)
	EquipSkill(*net.Session, *EquipSkill_Out)
	UnequipSkill(*net.Session, *UnequipSkill_Out)
	StudySkillByCheat(*net.Session, *StudySkillByCheat_Out)
	TrainSkill(*net.Session, *TrainSkill_Out)
	FlushSkill(*net.Session, *FlushSkill_Out)
}

func DecodeIn(message []byte) Request {
	var actionId = message[0]
	var buffer = net.NewBuffer(message[1:])
	switch actionId {
	case 1:
		request := new(GetAllSkillsInfo_In)
		request.Decode(buffer)
		return request
	case 2:
		request := new(EquipSkill_In)
		request.Decode(buffer)
		return request
	case 3:
		request := new(UnequipSkill_In)
		request.Decode(buffer)
		return request
	case 4:
		request := new(StudySkillByCheat_In)
		request.Decode(buffer)
		return request
	case 5:
		request := new(TrainSkill_In)
		request.Decode(buffer)
		return request
	case 6:
		request := new(FlushSkill_In)
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
		request := new(GetAllSkillsInfo_Out)
		request.Decode(buffer)
		return request
	case 2:
		request := new(EquipSkill_Out)
		request.Decode(buffer)
		return request
	case 3:
		request := new(UnequipSkill_Out)
		request.Decode(buffer)
		return request
	case 4:
		request := new(StudySkillByCheat_Out)
		request.Decode(buffer)
		return request
	case 5:
		request := new(TrainSkill_Out)
		request.Decode(buffer)
		return request
	case 6:
		request := new(FlushSkill_Out)
		request.Decode(buffer)
		return request
	}
	_ = buffer
	panic("unsupported response")
}

type CheatResult int8

const (
	CHEAT_RESULT_SUCCESS              CheatResult = 0
	CHEAT_RESULT_NO_DEPAND_SKILL      CheatResult = 1
	CHEAT_RESULT_ALREADY_STUDY        CheatResult = 2
	CHEAT_RESULT_CAN_NOT_STUDY_BEFORE CheatResult = 3
	CHEAT_RESULT_NOT_ROLE_SKILL       CheatResult = 4
	CHEAT_RESULT_NOT_CHEAT_TYPE       CheatResult = 5
	CHEAT_RESULT_ROLE_DOES_NOT_EXISTS CheatResult = 6
	CHEAT_RESULT_SKILL_NOT_MATCH_ROLE CheatResult = 7
	CHEAT_RESULT_LEVEL_NOT_REACHED    CheatResult = 8
)

type GetAllSkillsInfo_In struct {
}

func (this *GetAllSkillsInfo_In) Process(session *net.Session) {
	g_InHandler.GetAllSkillsInfo(session, this)
}

func (this *GetAllSkillsInfo_In) TypeName() string {
	return "skill.get_all_skills_info.in"
}

func (this *GetAllSkillsInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 5, 1
}

type GetAllSkillsInfo_Out struct {
	FlushTime int64                         `json:"flush_time"`
	Roles     []GetAllSkillsInfo_Out_Roles  `json:"roles"`
	Skills    []GetAllSkillsInfo_Out_Skills `json:"skills"`
}

type GetAllSkillsInfo_Out_Roles struct {
	RoleId   int8  `json:"role_id"`
	Status   int8  `json:"status"`
	SkillId1 int16 `json:"skill_id1"`
	SkillId2 int16 `json:"skill_id2"`
	SkillId3 int16 `json:"skill_id3"`
	SkillId4 int16 `json:"skill_id4"`
}

type GetAllSkillsInfo_Out_Skills struct {
	RoleId        int8  `json:"role_id"`
	SkillId       int16 `json:"skill_id"`
	TrainingLevel int16 `json:"training_level"`
}

func (this *GetAllSkillsInfo_Out) Process(session *net.Session) {
	g_OutHandler.GetAllSkillsInfo(session, this)
}

func (this *GetAllSkillsInfo_Out) TypeName() string {
	return "skill.get_all_skills_info.out"
}

func (this *GetAllSkillsInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 5, 1
}

func (this *GetAllSkillsInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type EquipSkill_In struct {
	RoleId      int8  `json:"role_id"`
	OrderNumber int8  `json:"order_number"`
	SkillId     int16 `json:"skill_id"`
}

func (this *EquipSkill_In) Process(session *net.Session) {
	g_InHandler.EquipSkill(session, this)
}

func (this *EquipSkill_In) TypeName() string {
	return "skill.equip_skill.in"
}

func (this *EquipSkill_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 5, 2
}

type EquipSkill_Out struct {
}

func (this *EquipSkill_Out) Process(session *net.Session) {
	g_OutHandler.EquipSkill(session, this)
}

func (this *EquipSkill_Out) TypeName() string {
	return "skill.equip_skill.out"
}

func (this *EquipSkill_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 5, 2
}

func (this *EquipSkill_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type UnequipSkill_In struct {
	RoleId      int8 `json:"role_id"`
	OrderNumber int8 `json:"order_number"`
}

func (this *UnequipSkill_In) Process(session *net.Session) {
	g_InHandler.UnequipSkill(session, this)
}

func (this *UnequipSkill_In) TypeName() string {
	return "skill.unequip_skill.in"
}

func (this *UnequipSkill_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 5, 3
}

type UnequipSkill_Out struct {
}

func (this *UnequipSkill_Out) Process(session *net.Session) {
	g_OutHandler.UnequipSkill(session, this)
}

func (this *UnequipSkill_Out) TypeName() string {
	return "skill.unequip_skill.out"
}

func (this *UnequipSkill_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 5, 3
}

func (this *UnequipSkill_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type StudySkillByCheat_In struct {
	RoleId int8  `json:"role_id"`
	ItemId int16 `json:"item_id"`
}

func (this *StudySkillByCheat_In) Process(session *net.Session) {
	g_InHandler.StudySkillByCheat(session, this)
}

func (this *StudySkillByCheat_In) TypeName() string {
	return "skill.study_skill_by_cheat.in"
}

func (this *StudySkillByCheat_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 5, 4
}

type StudySkillByCheat_Out struct {
	Result CheatResult `json:"result"`
}

func (this *StudySkillByCheat_Out) Process(session *net.Session) {
	g_OutHandler.StudySkillByCheat(session, this)
}

func (this *StudySkillByCheat_Out) TypeName() string {
	return "skill.study_skill_by_cheat.out"
}

func (this *StudySkillByCheat_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 5, 4
}

func (this *StudySkillByCheat_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type TrainSkill_In struct {
	RoleId  int8  `json:"role_id"`
	SkillId int16 `json:"skill_id"`
}

func (this *TrainSkill_In) Process(session *net.Session) {
	g_InHandler.TrainSkill(session, this)
}

func (this *TrainSkill_In) TypeName() string {
	return "skill.train_skill.in"
}

func (this *TrainSkill_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 5, 5
}

type TrainSkill_Out struct {
}

func (this *TrainSkill_Out) Process(session *net.Session) {
	g_OutHandler.TrainSkill(session, this)
}

func (this *TrainSkill_Out) TypeName() string {
	return "skill.train_skill.out"
}

func (this *TrainSkill_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 5, 5
}

func (this *TrainSkill_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type FlushSkill_In struct {
	RoleId int8 `json:"role_id"`
}

func (this *FlushSkill_In) Process(session *net.Session) {
	g_InHandler.FlushSkill(session, this)
}

func (this *FlushSkill_In) TypeName() string {
	return "skill.flush_skill.in"
}

func (this *FlushSkill_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 5, 6
}

type FlushSkill_Out struct {
	FlushTime int64 `json:"flush_time"`
}

func (this *FlushSkill_Out) Process(session *net.Session) {
	g_OutHandler.FlushSkill(session, this)
}

func (this *FlushSkill_Out) TypeName() string {
	return "skill.flush_skill.out"
}

func (this *FlushSkill_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 5, 6
}

func (this *FlushSkill_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

func (this *GetAllSkillsInfo_In) Decode(buffer *net.Buffer) {
}

func (this *GetAllSkillsInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(5)
	buffer.WriteUint8(1)
}

func (this *GetAllSkillsInfo_In) ByteSize() int {
	size := 2
	return size
}

func (this *GetAllSkillsInfo_Out) Decode(buffer *net.Buffer) {
	this.FlushTime = int64(buffer.ReadUint64LE())
	this.Roles = make([]GetAllSkillsInfo_Out_Roles, buffer.ReadUint8())
	for i := 0; i < len(this.Roles); i++ {
		this.Roles[i].Decode(buffer)
	}
	this.Skills = make([]GetAllSkillsInfo_Out_Skills, buffer.ReadUint8())
	for i := 0; i < len(this.Skills); i++ {
		this.Skills[i].Decode(buffer)
	}
}

func (this *GetAllSkillsInfo_Out_Roles) Decode(buffer *net.Buffer) {
	this.RoleId = int8(buffer.ReadUint8())
	this.Status = int8(buffer.ReadUint8())
	this.SkillId1 = int16(buffer.ReadUint16LE())
	this.SkillId2 = int16(buffer.ReadUint16LE())
	this.SkillId3 = int16(buffer.ReadUint16LE())
	this.SkillId4 = int16(buffer.ReadUint16LE())
}

func (this *GetAllSkillsInfo_Out_Skills) Decode(buffer *net.Buffer) {
	this.RoleId = int8(buffer.ReadUint8())
	this.SkillId = int16(buffer.ReadUint16LE())
	this.TrainingLevel = int16(buffer.ReadUint16LE())
}

func (this *GetAllSkillsInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(5)
	buffer.WriteUint8(1)
	buffer.WriteUint64LE(uint64(this.FlushTime))
	buffer.WriteUint8(uint8(len(this.Roles)))
	for i := 0; i < len(this.Roles); i++ {
		this.Roles[i].Encode(buffer)
	}
	buffer.WriteUint8(uint8(len(this.Skills)))
	for i := 0; i < len(this.Skills); i++ {
		this.Skills[i].Encode(buffer)
	}
}

func (this *GetAllSkillsInfo_Out_Roles) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint8(uint8(this.Status))
	buffer.WriteUint16LE(uint16(this.SkillId1))
	buffer.WriteUint16LE(uint16(this.SkillId2))
	buffer.WriteUint16LE(uint16(this.SkillId3))
	buffer.WriteUint16LE(uint16(this.SkillId4))
}

func (this *GetAllSkillsInfo_Out_Skills) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint16LE(uint16(this.SkillId))
	buffer.WriteUint16LE(uint16(this.TrainingLevel))
}

func (this *GetAllSkillsInfo_Out) ByteSize() int {
	size := 12
	size += len(this.Roles) * 10
	size += len(this.Skills) * 5
	return size
}

func (this *EquipSkill_In) Decode(buffer *net.Buffer) {
	this.RoleId = int8(buffer.ReadUint8())
	this.OrderNumber = int8(buffer.ReadUint8())
	this.SkillId = int16(buffer.ReadUint16LE())
}

func (this *EquipSkill_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(5)
	buffer.WriteUint8(2)
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint8(uint8(this.OrderNumber))
	buffer.WriteUint16LE(uint16(this.SkillId))
}

func (this *EquipSkill_In) ByteSize() int {
	size := 6
	return size
}

func (this *EquipSkill_Out) Decode(buffer *net.Buffer) {
}

func (this *EquipSkill_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(5)
	buffer.WriteUint8(2)
}

func (this *EquipSkill_Out) ByteSize() int {
	size := 2
	return size
}

func (this *UnequipSkill_In) Decode(buffer *net.Buffer) {
	this.RoleId = int8(buffer.ReadUint8())
	this.OrderNumber = int8(buffer.ReadUint8())
}

func (this *UnequipSkill_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(5)
	buffer.WriteUint8(3)
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint8(uint8(this.OrderNumber))
}

func (this *UnequipSkill_In) ByteSize() int {
	size := 4
	return size
}

func (this *UnequipSkill_Out) Decode(buffer *net.Buffer) {
}

func (this *UnequipSkill_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(5)
	buffer.WriteUint8(3)
}

func (this *UnequipSkill_Out) ByteSize() int {
	size := 2
	return size
}

func (this *StudySkillByCheat_In) Decode(buffer *net.Buffer) {
	this.RoleId = int8(buffer.ReadUint8())
	this.ItemId = int16(buffer.ReadUint16LE())
}

func (this *StudySkillByCheat_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(5)
	buffer.WriteUint8(4)
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint16LE(uint16(this.ItemId))
}

func (this *StudySkillByCheat_In) ByteSize() int {
	size := 5
	return size
}

func (this *StudySkillByCheat_Out) Decode(buffer *net.Buffer) {
	this.Result = CheatResult(buffer.ReadUint8())
}

func (this *StudySkillByCheat_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(5)
	buffer.WriteUint8(4)
	buffer.WriteUint8(uint8(this.Result))
}

func (this *StudySkillByCheat_Out) ByteSize() int {
	size := 3
	return size
}

func (this *TrainSkill_In) Decode(buffer *net.Buffer) {
	this.RoleId = int8(buffer.ReadUint8())
	this.SkillId = int16(buffer.ReadUint16LE())
}

func (this *TrainSkill_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(5)
	buffer.WriteUint8(5)
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint16LE(uint16(this.SkillId))
}

func (this *TrainSkill_In) ByteSize() int {
	size := 5
	return size
}

func (this *TrainSkill_Out) Decode(buffer *net.Buffer) {
}

func (this *TrainSkill_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(5)
	buffer.WriteUint8(5)
}

func (this *TrainSkill_Out) ByteSize() int {
	size := 2
	return size
}

func (this *FlushSkill_In) Decode(buffer *net.Buffer) {
	this.RoleId = int8(buffer.ReadUint8())
}

func (this *FlushSkill_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(5)
	buffer.WriteUint8(6)
	buffer.WriteUint8(uint8(this.RoleId))
}

func (this *FlushSkill_In) ByteSize() int {
	size := 3
	return size
}

func (this *FlushSkill_Out) Decode(buffer *net.Buffer) {
	this.FlushTime = int64(buffer.ReadUint64LE())
}

func (this *FlushSkill_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(5)
	buffer.WriteUint8(6)
	buffer.WriteUint64LE(uint64(this.FlushTime))
}

func (this *FlushSkill_Out) ByteSize() int {
	size := 10
	return size
}
