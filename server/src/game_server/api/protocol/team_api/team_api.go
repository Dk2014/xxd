package team_api

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
	GetFormationInfo(*net.Session, *GetFormationInfo_In)
	UpFormation(*net.Session, *UpFormation_In)
	DownFormation(*net.Session, *DownFormation_In)
	SwapFormation(*net.Session, *SwapFormation_In)
	ReplaceFormation(*net.Session, *ReplaceFormation_In)
	TrainingTeamship(*net.Session, *TrainingTeamship_In)
}

type OutHandler interface {
	GetFormationInfo(*net.Session, *GetFormationInfo_Out)
	UpFormation(*net.Session, *UpFormation_Out)
	DownFormation(*net.Session, *DownFormation_Out)
	SwapFormation(*net.Session, *SwapFormation_Out)
	ReplaceFormation(*net.Session, *ReplaceFormation_Out)
	TrainingTeamship(*net.Session, *TrainingTeamship_Out)
}

func DecodeIn(message []byte) Request {
	var actionId = message[0]
	var buffer = net.NewBuffer(message[1:])
	switch actionId {
	case 0:
		request := new(GetFormationInfo_In)
		request.Decode(buffer)
		return request
	case 2:
		request := new(UpFormation_In)
		request.Decode(buffer)
		return request
	case 3:
		request := new(DownFormation_In)
		request.Decode(buffer)
		return request
	case 4:
		request := new(SwapFormation_In)
		request.Decode(buffer)
		return request
	case 5:
		request := new(ReplaceFormation_In)
		request.Decode(buffer)
		return request
	case 6:
		request := new(TrainingTeamship_In)
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
		request := new(GetFormationInfo_Out)
		request.Decode(buffer)
		return request
	case 2:
		request := new(UpFormation_Out)
		request.Decode(buffer)
		return request
	case 3:
		request := new(DownFormation_Out)
		request.Decode(buffer)
		return request
	case 4:
		request := new(SwapFormation_Out)
		request.Decode(buffer)
		return request
	case 5:
		request := new(ReplaceFormation_Out)
		request.Decode(buffer)
		return request
	case 6:
		request := new(TrainingTeamship_Out)
		request.Decode(buffer)
		return request
	}
	_ = buffer
	panic("unsupported response")
}

type GetFormationInfo_In struct {
}

func (this *GetFormationInfo_In) Process(session *net.Session) {
	g_InHandler.GetFormationInfo(session, this)
}

func (this *GetFormationInfo_In) TypeName() string {
	return "team.get_formation_info.in"
}

func (this *GetFormationInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 2, 0
}

type GetFormationInfo_Out struct {
	Pos0Role     int8  `json:"pos0_role"`
	Pos1Role     int8  `json:"pos1_role"`
	Pos2Role     int8  `json:"pos2_role"`
	Pos3Role     int8  `json:"pos3_role"`
	Pos4Role     int8  `json:"pos4_role"`
	Pos5Role     int8  `json:"pos5_role"`
	Pos6Role     int8  `json:"pos6_role"`
	Pos7Role     int8  `json:"pos7_role"`
	Pos8Role     int8  `json:"pos8_role"`
	Relationship int32 `json:"relationship"`
	HealthLv     int16 `json:"health_lv"`
	AttackLv     int16 `json:"attack_lv"`
	DefenceLv    int16 `json:"defence_lv"`
}

func (this *GetFormationInfo_Out) Process(session *net.Session) {
	g_OutHandler.GetFormationInfo(session, this)
}

func (this *GetFormationInfo_Out) TypeName() string {
	return "team.get_formation_info.out"
}

func (this *GetFormationInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 2, 0
}

func (this *GetFormationInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type UpFormation_In struct {
	RoleId int8 `json:"role_id"`
	Pos    int8 `json:"pos"`
}

func (this *UpFormation_In) Process(session *net.Session) {
	g_InHandler.UpFormation(session, this)
}

func (this *UpFormation_In) TypeName() string {
	return "team.up_formation.in"
}

func (this *UpFormation_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 2, 2
}

type UpFormation_Out struct {
}

func (this *UpFormation_Out) Process(session *net.Session) {
	g_OutHandler.UpFormation(session, this)
}

func (this *UpFormation_Out) TypeName() string {
	return "team.up_formation.out"
}

func (this *UpFormation_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 2, 2
}

func (this *UpFormation_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type DownFormation_In struct {
	Pos int8 `json:"pos"`
}

func (this *DownFormation_In) Process(session *net.Session) {
	g_InHandler.DownFormation(session, this)
}

func (this *DownFormation_In) TypeName() string {
	return "team.down_formation.in"
}

func (this *DownFormation_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 2, 3
}

type DownFormation_Out struct {
}

func (this *DownFormation_Out) Process(session *net.Session) {
	g_OutHandler.DownFormation(session, this)
}

func (this *DownFormation_Out) TypeName() string {
	return "team.down_formation.out"
}

func (this *DownFormation_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 2, 3
}

func (this *DownFormation_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type SwapFormation_In struct {
	PosFrom int8 `json:"pos_from"`
	PosTo   int8 `json:"pos_to"`
}

func (this *SwapFormation_In) Process(session *net.Session) {
	g_InHandler.SwapFormation(session, this)
}

func (this *SwapFormation_In) TypeName() string {
	return "team.swap_formation.in"
}

func (this *SwapFormation_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 2, 4
}

type SwapFormation_Out struct {
}

func (this *SwapFormation_Out) Process(session *net.Session) {
	g_OutHandler.SwapFormation(session, this)
}

func (this *SwapFormation_Out) TypeName() string {
	return "team.swap_formation.out"
}

func (this *SwapFormation_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 2, 4
}

func (this *SwapFormation_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ReplaceFormation_In struct {
	RoleId int8 `json:"role_id"`
	Pos    int8 `json:"pos"`
}

func (this *ReplaceFormation_In) Process(session *net.Session) {
	g_InHandler.ReplaceFormation(session, this)
}

func (this *ReplaceFormation_In) TypeName() string {
	return "team.replace_formation.in"
}

func (this *ReplaceFormation_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 2, 5
}

type ReplaceFormation_Out struct {
}

func (this *ReplaceFormation_Out) Process(session *net.Session) {
	g_OutHandler.ReplaceFormation(session, this)
}

func (this *ReplaceFormation_Out) TypeName() string {
	return "team.replace_formation.out"
}

func (this *ReplaceFormation_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 2, 5
}

func (this *ReplaceFormation_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type TrainingTeamship_In struct {
	AttrInd int8 `json:"attr_ind"`
}

func (this *TrainingTeamship_In) Process(session *net.Session) {
	g_InHandler.TrainingTeamship(session, this)
}

func (this *TrainingTeamship_In) TypeName() string {
	return "team.training_teamship.in"
}

func (this *TrainingTeamship_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 2, 6
}

type TrainingTeamship_Out struct {
}

func (this *TrainingTeamship_Out) Process(session *net.Session) {
	g_OutHandler.TrainingTeamship(session, this)
}

func (this *TrainingTeamship_Out) TypeName() string {
	return "team.training_teamship.out"
}

func (this *TrainingTeamship_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 2, 6
}

func (this *TrainingTeamship_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

func (this *GetFormationInfo_In) Decode(buffer *net.Buffer) {
}

func (this *GetFormationInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(2)
	buffer.WriteUint8(0)
}

func (this *GetFormationInfo_In) ByteSize() int {
	size := 2
	return size
}

func (this *GetFormationInfo_Out) Decode(buffer *net.Buffer) {
	this.Pos0Role = int8(buffer.ReadUint8())
	this.Pos1Role = int8(buffer.ReadUint8())
	this.Pos2Role = int8(buffer.ReadUint8())
	this.Pos3Role = int8(buffer.ReadUint8())
	this.Pos4Role = int8(buffer.ReadUint8())
	this.Pos5Role = int8(buffer.ReadUint8())
	this.Pos6Role = int8(buffer.ReadUint8())
	this.Pos7Role = int8(buffer.ReadUint8())
	this.Pos8Role = int8(buffer.ReadUint8())
	this.Relationship = int32(buffer.ReadUint32LE())
	this.HealthLv = int16(buffer.ReadUint16LE())
	this.AttackLv = int16(buffer.ReadUint16LE())
	this.DefenceLv = int16(buffer.ReadUint16LE())
}

func (this *GetFormationInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(2)
	buffer.WriteUint8(0)
	buffer.WriteUint8(uint8(this.Pos0Role))
	buffer.WriteUint8(uint8(this.Pos1Role))
	buffer.WriteUint8(uint8(this.Pos2Role))
	buffer.WriteUint8(uint8(this.Pos3Role))
	buffer.WriteUint8(uint8(this.Pos4Role))
	buffer.WriteUint8(uint8(this.Pos5Role))
	buffer.WriteUint8(uint8(this.Pos6Role))
	buffer.WriteUint8(uint8(this.Pos7Role))
	buffer.WriteUint8(uint8(this.Pos8Role))
	buffer.WriteUint32LE(uint32(this.Relationship))
	buffer.WriteUint16LE(uint16(this.HealthLv))
	buffer.WriteUint16LE(uint16(this.AttackLv))
	buffer.WriteUint16LE(uint16(this.DefenceLv))
}

func (this *GetFormationInfo_Out) ByteSize() int {
	size := 21
	return size
}

func (this *UpFormation_In) Decode(buffer *net.Buffer) {
	this.RoleId = int8(buffer.ReadUint8())
	this.Pos = int8(buffer.ReadUint8())
}

func (this *UpFormation_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(2)
	buffer.WriteUint8(2)
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint8(uint8(this.Pos))
}

func (this *UpFormation_In) ByteSize() int {
	size := 4
	return size
}

func (this *UpFormation_Out) Decode(buffer *net.Buffer) {
}

func (this *UpFormation_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(2)
	buffer.WriteUint8(2)
}

func (this *UpFormation_Out) ByteSize() int {
	size := 2
	return size
}

func (this *DownFormation_In) Decode(buffer *net.Buffer) {
	this.Pos = int8(buffer.ReadUint8())
}

func (this *DownFormation_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(2)
	buffer.WriteUint8(3)
	buffer.WriteUint8(uint8(this.Pos))
}

func (this *DownFormation_In) ByteSize() int {
	size := 3
	return size
}

func (this *DownFormation_Out) Decode(buffer *net.Buffer) {
}

func (this *DownFormation_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(2)
	buffer.WriteUint8(3)
}

func (this *DownFormation_Out) ByteSize() int {
	size := 2
	return size
}

func (this *SwapFormation_In) Decode(buffer *net.Buffer) {
	this.PosFrom = int8(buffer.ReadUint8())
	this.PosTo = int8(buffer.ReadUint8())
}

func (this *SwapFormation_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(2)
	buffer.WriteUint8(4)
	buffer.WriteUint8(uint8(this.PosFrom))
	buffer.WriteUint8(uint8(this.PosTo))
}

func (this *SwapFormation_In) ByteSize() int {
	size := 4
	return size
}

func (this *SwapFormation_Out) Decode(buffer *net.Buffer) {
}

func (this *SwapFormation_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(2)
	buffer.WriteUint8(4)
}

func (this *SwapFormation_Out) ByteSize() int {
	size := 2
	return size
}

func (this *ReplaceFormation_In) Decode(buffer *net.Buffer) {
	this.RoleId = int8(buffer.ReadUint8())
	this.Pos = int8(buffer.ReadUint8())
}

func (this *ReplaceFormation_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(2)
	buffer.WriteUint8(5)
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint8(uint8(this.Pos))
}

func (this *ReplaceFormation_In) ByteSize() int {
	size := 4
	return size
}

func (this *ReplaceFormation_Out) Decode(buffer *net.Buffer) {
}

func (this *ReplaceFormation_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(2)
	buffer.WriteUint8(5)
}

func (this *ReplaceFormation_Out) ByteSize() int {
	size := 2
	return size
}

func (this *TrainingTeamship_In) Decode(buffer *net.Buffer) {
	this.AttrInd = int8(buffer.ReadUint8())
}

func (this *TrainingTeamship_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(2)
	buffer.WriteUint8(6)
	buffer.WriteUint8(uint8(this.AttrInd))
}

func (this *TrainingTeamship_In) ByteSize() int {
	size := 3
	return size
}

func (this *TrainingTeamship_Out) Decode(buffer *net.Buffer) {
}

func (this *TrainingTeamship_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(2)
	buffer.WriteUint8(6)
}

func (this *TrainingTeamship_Out) ByteSize() int {
	size := 2
	return size
}
