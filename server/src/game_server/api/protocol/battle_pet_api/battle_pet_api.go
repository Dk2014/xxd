package battle_pet_api

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
	GetPetInfo(*net.Session, *GetPetInfo_In)
	SetPet(*net.Session, *SetPet_In)
	SetPetSwap(*net.Session, *SetPetSwap_In)
	UpgradePet(*net.Session, *UpgradePet_In)
	TrainingPetSkill(*net.Session, *TrainingPetSkill_In)
}

type OutHandler interface {
	GetPetInfo(*net.Session, *GetPetInfo_Out)
	SetPet(*net.Session, *SetPet_Out)
	SetPetSwap(*net.Session, *SetPetSwap_Out)
	UpgradePet(*net.Session, *UpgradePet_Out)
	TrainingPetSkill(*net.Session, *TrainingPetSkill_Out)
}

func DecodeIn(message []byte) Request {
	var actionId = message[0]
	var buffer = net.NewBuffer(message[1:])
	switch actionId {
	case 1:
		request := new(GetPetInfo_In)
		request.Decode(buffer)
		return request
	case 3:
		request := new(SetPet_In)
		request.Decode(buffer)
		return request
	case 4:
		request := new(SetPetSwap_In)
		request.Decode(buffer)
		return request
	case 6:
		request := new(UpgradePet_In)
		request.Decode(buffer)
		return request
	case 7:
		request := new(TrainingPetSkill_In)
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
		request := new(GetPetInfo_Out)
		request.Decode(buffer)
		return request
	case 3:
		request := new(SetPet_Out)
		request.Decode(buffer)
		return request
	case 4:
		request := new(SetPetSwap_Out)
		request.Decode(buffer)
		return request
	case 6:
		request := new(UpgradePet_Out)
		request.Decode(buffer)
		return request
	case 7:
		request := new(TrainingPetSkill_Out)
		request.Decode(buffer)
		return request
	}
	_ = buffer
	panic("unsupported response")
}

type GetPetInfo_In struct {
}

func (this *GetPetInfo_In) Process(session *net.Session) {
	g_InHandler.GetPetInfo(session, this)
}

func (this *GetPetInfo_In) TypeName() string {
	return "battle_pet.get_pet_info.in"
}

func (this *GetPetInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 17, 1
}

type GetPetInfo_Out struct {
	Pet []GetPetInfo_Out_Pet `json:"pet"`
	Set []GetPetInfo_Out_Set `json:"set"`
}

type GetPetInfo_Out_Pet struct {
	PetId      int32 `json:"pet_id"`
	Level      int16 `json:"level"`
	Exp        int64 `json:"exp"`
	SkillLevel int16 `json:"skill_level"`
	Called     bool  `json:"called"`
}

type GetPetInfo_Out_Set struct {
	GridNum int8  `json:"grid_num"`
	PetId   int32 `json:"pet_id"`
}

func (this *GetPetInfo_Out) Process(session *net.Session) {
	g_OutHandler.GetPetInfo(session, this)
}

func (this *GetPetInfo_Out) TypeName() string {
	return "battle_pet.get_pet_info.out"
}

func (this *GetPetInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 17, 1
}

func (this *GetPetInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type SetPet_In struct {
	GridNum int8  `json:"grid_num"`
	PetId   int32 `json:"pet_id"`
}

func (this *SetPet_In) Process(session *net.Session) {
	g_InHandler.SetPet(session, this)
}

func (this *SetPet_In) TypeName() string {
	return "battle_pet.set_pet.in"
}

func (this *SetPet_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 17, 3
}

type SetPet_Out struct {
}

func (this *SetPet_Out) Process(session *net.Session) {
	g_OutHandler.SetPet(session, this)
}

func (this *SetPet_Out) TypeName() string {
	return "battle_pet.set_pet.out"
}

func (this *SetPet_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 17, 3
}

func (this *SetPet_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type SetPetSwap_In struct {
	FromGridNum int8 `json:"from_grid_num"`
	ToGridNum   int8 `json:"to_grid_num"`
}

func (this *SetPetSwap_In) Process(session *net.Session) {
	g_InHandler.SetPetSwap(session, this)
}

func (this *SetPetSwap_In) TypeName() string {
	return "battle_pet.set_pet_swap.in"
}

func (this *SetPetSwap_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 17, 4
}

type SetPetSwap_Out struct {
}

func (this *SetPetSwap_Out) Process(session *net.Session) {
	g_OutHandler.SetPetSwap(session, this)
}

func (this *SetPetSwap_Out) TypeName() string {
	return "battle_pet.set_pet_swap.out"
}

func (this *SetPetSwap_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 17, 4
}

func (this *SetPetSwap_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type UpgradePet_In struct {
	PetId int32 `json:"pet_id"`
}

func (this *UpgradePet_In) Process(session *net.Session) {
	g_InHandler.UpgradePet(session, this)
}

func (this *UpgradePet_In) TypeName() string {
	return "battle_pet.upgrade_pet.in"
}

func (this *UpgradePet_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 17, 6
}

type UpgradePet_Out struct {
	Exp   int64 `json:"exp"`
	Level int16 `json:"level"`
}

func (this *UpgradePet_Out) Process(session *net.Session) {
	g_OutHandler.UpgradePet(session, this)
}

func (this *UpgradePet_Out) TypeName() string {
	return "battle_pet.upgrade_pet.out"
}

func (this *UpgradePet_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 17, 6
}

func (this *UpgradePet_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type TrainingPetSkill_In struct {
	PetId int32 `json:"pet_id"`
}

func (this *TrainingPetSkill_In) Process(session *net.Session) {
	g_InHandler.TrainingPetSkill(session, this)
}

func (this *TrainingPetSkill_In) TypeName() string {
	return "battle_pet.training_pet_skill.in"
}

func (this *TrainingPetSkill_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 17, 7
}

type TrainingPetSkill_Out struct {
}

func (this *TrainingPetSkill_Out) Process(session *net.Session) {
	g_OutHandler.TrainingPetSkill(session, this)
}

func (this *TrainingPetSkill_Out) TypeName() string {
	return "battle_pet.training_pet_skill.out"
}

func (this *TrainingPetSkill_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 17, 7
}

func (this *TrainingPetSkill_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

func (this *GetPetInfo_In) Decode(buffer *net.Buffer) {
}

func (this *GetPetInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(17)
	buffer.WriteUint8(1)
}

func (this *GetPetInfo_In) ByteSize() int {
	size := 2
	return size
}

func (this *GetPetInfo_Out) Decode(buffer *net.Buffer) {
	this.Pet = make([]GetPetInfo_Out_Pet, buffer.ReadUint8())
	for i := 0; i < len(this.Pet); i++ {
		this.Pet[i].Decode(buffer)
	}
	this.Set = make([]GetPetInfo_Out_Set, buffer.ReadUint8())
	for i := 0; i < len(this.Set); i++ {
		this.Set[i].Decode(buffer)
	}
}

func (this *GetPetInfo_Out_Pet) Decode(buffer *net.Buffer) {
	this.PetId = int32(buffer.ReadUint32LE())
	this.Level = int16(buffer.ReadUint16LE())
	this.Exp = int64(buffer.ReadUint64LE())
	this.SkillLevel = int16(buffer.ReadUint16LE())
	this.Called = buffer.ReadUint8() == 1
}

func (this *GetPetInfo_Out_Set) Decode(buffer *net.Buffer) {
	this.GridNum = int8(buffer.ReadUint8())
	this.PetId = int32(buffer.ReadUint32LE())
}

func (this *GetPetInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(17)
	buffer.WriteUint8(1)
	buffer.WriteUint8(uint8(len(this.Pet)))
	for i := 0; i < len(this.Pet); i++ {
		this.Pet[i].Encode(buffer)
	}
	buffer.WriteUint8(uint8(len(this.Set)))
	for i := 0; i < len(this.Set); i++ {
		this.Set[i].Encode(buffer)
	}
}

func (this *GetPetInfo_Out_Pet) Encode(buffer *net.Buffer) {
	buffer.WriteUint32LE(uint32(this.PetId))
	buffer.WriteUint16LE(uint16(this.Level))
	buffer.WriteUint64LE(uint64(this.Exp))
	buffer.WriteUint16LE(uint16(this.SkillLevel))
	if this.Called {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *GetPetInfo_Out_Set) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(this.GridNum))
	buffer.WriteUint32LE(uint32(this.PetId))
}

func (this *GetPetInfo_Out) ByteSize() int {
	size := 4
	size += len(this.Pet) * 17
	size += len(this.Set) * 5
	return size
}

func (this *SetPet_In) Decode(buffer *net.Buffer) {
	this.GridNum = int8(buffer.ReadUint8())
	this.PetId = int32(buffer.ReadUint32LE())
}

func (this *SetPet_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(17)
	buffer.WriteUint8(3)
	buffer.WriteUint8(uint8(this.GridNum))
	buffer.WriteUint32LE(uint32(this.PetId))
}

func (this *SetPet_In) ByteSize() int {
	size := 7
	return size
}

func (this *SetPet_Out) Decode(buffer *net.Buffer) {
}

func (this *SetPet_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(17)
	buffer.WriteUint8(3)
}

func (this *SetPet_Out) ByteSize() int {
	size := 2
	return size
}

func (this *SetPetSwap_In) Decode(buffer *net.Buffer) {
	this.FromGridNum = int8(buffer.ReadUint8())
	this.ToGridNum = int8(buffer.ReadUint8())
}

func (this *SetPetSwap_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(17)
	buffer.WriteUint8(4)
	buffer.WriteUint8(uint8(this.FromGridNum))
	buffer.WriteUint8(uint8(this.ToGridNum))
}

func (this *SetPetSwap_In) ByteSize() int {
	size := 4
	return size
}

func (this *SetPetSwap_Out) Decode(buffer *net.Buffer) {
}

func (this *SetPetSwap_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(17)
	buffer.WriteUint8(4)
}

func (this *SetPetSwap_Out) ByteSize() int {
	size := 2
	return size
}

func (this *UpgradePet_In) Decode(buffer *net.Buffer) {
	this.PetId = int32(buffer.ReadUint32LE())
}

func (this *UpgradePet_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(17)
	buffer.WriteUint8(6)
	buffer.WriteUint32LE(uint32(this.PetId))
}

func (this *UpgradePet_In) ByteSize() int {
	size := 6
	return size
}

func (this *UpgradePet_Out) Decode(buffer *net.Buffer) {
	this.Exp = int64(buffer.ReadUint64LE())
	this.Level = int16(buffer.ReadUint16LE())
}

func (this *UpgradePet_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(17)
	buffer.WriteUint8(6)
	buffer.WriteUint64LE(uint64(this.Exp))
	buffer.WriteUint16LE(uint16(this.Level))
}

func (this *UpgradePet_Out) ByteSize() int {
	size := 12
	return size
}

func (this *TrainingPetSkill_In) Decode(buffer *net.Buffer) {
	this.PetId = int32(buffer.ReadUint32LE())
}

func (this *TrainingPetSkill_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(17)
	buffer.WriteUint8(7)
	buffer.WriteUint32LE(uint32(this.PetId))
}

func (this *TrainingPetSkill_In) ByteSize() int {
	size := 6
	return size
}

func (this *TrainingPetSkill_Out) Decode(buffer *net.Buffer) {
}

func (this *TrainingPetSkill_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(17)
	buffer.WriteUint8(7)
}

func (this *TrainingPetSkill_Out) ByteSize() int {
	size := 2
	return size
}
