package clique_war_api

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
	WarInfo(*net.Session, *WarInfo_In)
	SignUpCliqueWar(*net.Session, *SignUpCliqueWar_In)
	StartCliqueWarBattle(*net.Session, *StartCliqueWarBattle_In)
	SignUpCliqueWarInfo(*net.Session, *SignUpCliqueWarInfo_In)
}

type OutHandler interface {
	WarInfo(*net.Session, *WarInfo_Out)
	NotifyWarReady(*net.Session, *NotifyWarReady_Out)
	NotifyWarStart(*net.Session, *NotifyWarStart_Out)
	NotifyWarEnd(*net.Session, *NotifyWarEnd_Out)
	SignUpCliqueWar(*net.Session, *SignUpCliqueWar_Out)
	StartCliqueWarBattle(*net.Session, *StartCliqueWarBattle_Out)
	SignUpCliqueWarInfo(*net.Session, *SignUpCliqueWarInfo_Out)
}

func DecodeIn(message []byte) Request {
	var actionId = message[0]
	var buffer = net.NewBuffer(message[1:])
	switch actionId {
	case 1:
		request := new(WarInfo_In)
		request.Decode(buffer)
		return request
	case 5:
		request := new(SignUpCliqueWar_In)
		request.Decode(buffer)
		return request
	case 6:
		request := new(StartCliqueWarBattle_In)
		request.Decode(buffer)
		return request
	case 7:
		request := new(SignUpCliqueWarInfo_In)
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
		request := new(WarInfo_Out)
		request.Decode(buffer)
		return request
	case 2:
		request := new(NotifyWarReady_Out)
		request.Decode(buffer)
		return request
	case 3:
		request := new(NotifyWarStart_Out)
		request.Decode(buffer)
		return request
	case 4:
		request := new(NotifyWarEnd_Out)
		request.Decode(buffer)
		return request
	case 5:
		request := new(SignUpCliqueWar_Out)
		request.Decode(buffer)
		return request
	case 6:
		request := new(StartCliqueWarBattle_Out)
		request.Decode(buffer)
		return request
	case 7:
		request := new(SignUpCliqueWarInfo_Out)
		request.Decode(buffer)
		return request
	}
	_ = buffer
	panic("unsupported response")
}

type SignUpCliqueWarResult int8

const (
	SIGN_UP_CLIQUE_WAR_RESULT_FAILED            SignUpCliqueWarResult = 0
	SIGN_UP_CLIQUE_WAR_RESULT_SUCCEED           SignUpCliqueWarResult = 1
	SIGN_UP_CLIQUE_WAR_RESULT_REPEAT            SignUpCliqueWarResult = 2
	SIGN_UP_CLIQUE_WAR_RESULT_NO_CLIQUE         SignUpCliqueWarResult = 3
	SIGN_UP_CLIQUE_WAR_RESULT_NOT_MANAGER       SignUpCliqueWarResult = 4
	SIGN_UP_CLIQUE_WAR_RESULT_MEMBER_NOT_ENOUGH SignUpCliqueWarResult = 5
	SIGN_UP_CLIQUE_WAR_RESULT_OUT_OF_TIME       SignUpCliqueWarResult = 6
)

type WarInfo_In struct {
}

func (this *WarInfo_In) Process(session *net.Session) {
	g_InHandler.WarInfo(session, this)
}

func (this *WarInfo_In) TypeName() string {
	return "clique_war.war_info.in"
}

func (this *WarInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 42, 1
}

type WarInfo_Out struct {
	StartTimestamp       int64
	EndTimestamp         int64
	RemainFightTimes     int64
	CliqueId             int64
	CliqueName           []byte
	CliqueServerId       int64
	BattleCliqueId       int64
	BattleCliqueName     []byte
	BattleCliqueServerId int64
	TotalPoint           int16
	BattleTotalPoint     int16
	MemberList           []WarInfo_Out_MemberList
	BattleMemerList      []WarInfo_Out_BattleMemerList
}

type WarInfo_Out_MemberList struct {
	GotPoint   int16
	Pid        int64
	Name       []byte
	Level      int16
	RoleId     int8
	FightNum   int64
	ServerId   int64
	HasBattled int32
}

type WarInfo_Out_BattleMemerList struct {
	Point    int16
	Pid      int64
	Name     []byte
	Level    int16
	RoleId   int8
	FightNum int64
	ServerId int64
	HasWin   int32
}

func (this *WarInfo_Out) Process(session *net.Session) {
	g_OutHandler.WarInfo(session, this)
}

func (this *WarInfo_Out) TypeName() string {
	return "clique_war.war_info.out"
}

func (this *WarInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 42, 1
}

func (this *WarInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type NotifyWarReady_Out struct {
	RemainTime int64
}

func (this *NotifyWarReady_Out) Process(session *net.Session) {
	g_OutHandler.NotifyWarReady(session, this)
}

func (this *NotifyWarReady_Out) TypeName() string {
	return "clique_war.notify_war_ready.out"
}

func (this *NotifyWarReady_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 42, 2
}

func (this *NotifyWarReady_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type NotifyWarStart_Out struct {
}

func (this *NotifyWarStart_Out) Process(session *net.Session) {
	g_OutHandler.NotifyWarStart(session, this)
}

func (this *NotifyWarStart_Out) TypeName() string {
	return "clique_war.notify_war_start.out"
}

func (this *NotifyWarStart_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 42, 3
}

func (this *NotifyWarStart_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type NotifyWarEnd_Out struct {
}

func (this *NotifyWarEnd_Out) Process(session *net.Session) {
	g_OutHandler.NotifyWarEnd(session, this)
}

func (this *NotifyWarEnd_Out) TypeName() string {
	return "clique_war.notify_war_end.out"
}

func (this *NotifyWarEnd_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 42, 4
}

func (this *NotifyWarEnd_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type SignUpCliqueWar_In struct {
}

func (this *SignUpCliqueWar_In) Process(session *net.Session) {
	g_InHandler.SignUpCliqueWar(session, this)
}

func (this *SignUpCliqueWar_In) TypeName() string {
	return "clique_war.sign_up_clique_war.in"
}

func (this *SignUpCliqueWar_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 42, 5
}

type SignUpCliqueWar_Out struct {
	Result SignUpCliqueWarResult
}

func (this *SignUpCliqueWar_Out) Process(session *net.Session) {
	g_OutHandler.SignUpCliqueWar(session, this)
}

func (this *SignUpCliqueWar_Out) TypeName() string {
	return "clique_war.sign_up_clique_war.out"
}

func (this *SignUpCliqueWar_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 42, 5
}

func (this *SignUpCliqueWar_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type StartCliqueWarBattle_In struct {
	TargetPlayerid int64
}

func (this *StartCliqueWarBattle_In) Process(session *net.Session) {
	g_InHandler.StartCliqueWarBattle(session, this)
}

func (this *StartCliqueWarBattle_In) TypeName() string {
	return "clique_war.start_clique_war_battle.in"
}

func (this *StartCliqueWarBattle_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 42, 6
}

type StartCliqueWarBattle_Out struct {
}

func (this *StartCliqueWarBattle_Out) Process(session *net.Session) {
	g_OutHandler.StartCliqueWarBattle(session, this)
}

func (this *StartCliqueWarBattle_Out) TypeName() string {
	return "clique_war.start_clique_war_battle.out"
}

func (this *StartCliqueWarBattle_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 42, 6
}

func (this *StartCliqueWarBattle_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type SignUpCliqueWarInfo_In struct {
}

func (this *SignUpCliqueWarInfo_In) Process(session *net.Session) {
	g_InHandler.SignUpCliqueWarInfo(session, this)
}

func (this *SignUpCliqueWarInfo_In) TypeName() string {
	return "clique_war.sign_up_clique_war_info.in"
}

func (this *SignUpCliqueWarInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 42, 7
}

type SignUpCliqueWarInfo_Out struct {
	Result bool
}

func (this *SignUpCliqueWarInfo_Out) Process(session *net.Session) {
	g_OutHandler.SignUpCliqueWarInfo(session, this)
}

func (this *SignUpCliqueWarInfo_Out) TypeName() string {
	return "clique_war.sign_up_clique_war_info.out"
}

func (this *SignUpCliqueWarInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 42, 7
}

func (this *SignUpCliqueWarInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

func (this *WarInfo_In) Decode(buffer *net.Buffer) {
}

func (this *WarInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(42)
	buffer.WriteUint8(1)
}

func (this *WarInfo_In) ByteSize() int {
	size := 2
	return size
}

func (this *WarInfo_Out) Decode(buffer *net.Buffer) {
	this.StartTimestamp = int64(buffer.ReadUint64LE())
	this.EndTimestamp = int64(buffer.ReadUint64LE())
	this.RemainFightTimes = int64(buffer.ReadUint64LE())
	this.CliqueId = int64(buffer.ReadUint64LE())
	this.CliqueName = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.CliqueServerId = int64(buffer.ReadUint64LE())
	this.BattleCliqueId = int64(buffer.ReadUint64LE())
	this.BattleCliqueName = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.BattleCliqueServerId = int64(buffer.ReadUint64LE())
	this.TotalPoint = int16(buffer.ReadUint16LE())
	this.BattleTotalPoint = int16(buffer.ReadUint16LE())
	this.MemberList = make([]WarInfo_Out_MemberList, buffer.ReadUint8())
	for i := 0; i < len(this.MemberList); i++ {
		this.MemberList[i].Decode(buffer)
	}
	this.BattleMemerList = make([]WarInfo_Out_BattleMemerList, buffer.ReadUint8())
	for i := 0; i < len(this.BattleMemerList); i++ {
		this.BattleMemerList[i].Decode(buffer)
	}
}

func (this *WarInfo_Out_MemberList) Decode(buffer *net.Buffer) {
	this.GotPoint = int16(buffer.ReadUint16LE())
	this.Pid = int64(buffer.ReadUint64LE())
	this.Name = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Level = int16(buffer.ReadUint16LE())
	this.RoleId = int8(buffer.ReadUint8())
	this.FightNum = int64(buffer.ReadUint64LE())
	this.ServerId = int64(buffer.ReadUint64LE())
	this.HasBattled = int32(buffer.ReadUint32LE())
}

func (this *WarInfo_Out_BattleMemerList) Decode(buffer *net.Buffer) {
	this.Point = int16(buffer.ReadUint16LE())
	this.Pid = int64(buffer.ReadUint64LE())
	this.Name = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Level = int16(buffer.ReadUint16LE())
	this.RoleId = int8(buffer.ReadUint8())
	this.FightNum = int64(buffer.ReadUint64LE())
	this.ServerId = int64(buffer.ReadUint64LE())
	this.HasWin = int32(buffer.ReadUint32LE())
}

func (this *WarInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(42)
	buffer.WriteUint8(1)
	buffer.WriteUint64LE(uint64(this.StartTimestamp))
	buffer.WriteUint64LE(uint64(this.EndTimestamp))
	buffer.WriteUint64LE(uint64(this.RemainFightTimes))
	buffer.WriteUint64LE(uint64(this.CliqueId))
	buffer.WriteUint16LE(uint16(len(this.CliqueName)))
	buffer.WriteBytes(this.CliqueName)
	buffer.WriteUint64LE(uint64(this.CliqueServerId))
	buffer.WriteUint64LE(uint64(this.BattleCliqueId))
	buffer.WriteUint16LE(uint16(len(this.BattleCliqueName)))
	buffer.WriteBytes(this.BattleCliqueName)
	buffer.WriteUint64LE(uint64(this.BattleCliqueServerId))
	buffer.WriteUint16LE(uint16(this.TotalPoint))
	buffer.WriteUint16LE(uint16(this.BattleTotalPoint))
	buffer.WriteUint8(uint8(len(this.MemberList)))
	for i := 0; i < len(this.MemberList); i++ {
		this.MemberList[i].Encode(buffer)
	}
	buffer.WriteUint8(uint8(len(this.BattleMemerList)))
	for i := 0; i < len(this.BattleMemerList); i++ {
		this.BattleMemerList[i].Encode(buffer)
	}
}

func (this *WarInfo_Out_MemberList) Encode(buffer *net.Buffer) {
	buffer.WriteUint16LE(uint16(this.GotPoint))
	buffer.WriteUint64LE(uint64(this.Pid))
	buffer.WriteUint16LE(uint16(len(this.Name)))
	buffer.WriteBytes(this.Name)
	buffer.WriteUint16LE(uint16(this.Level))
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint64LE(uint64(this.FightNum))
	buffer.WriteUint64LE(uint64(this.ServerId))
	buffer.WriteUint32LE(uint32(this.HasBattled))
}

func (this *WarInfo_Out_BattleMemerList) Encode(buffer *net.Buffer) {
	buffer.WriteUint16LE(uint16(this.Point))
	buffer.WriteUint64LE(uint64(this.Pid))
	buffer.WriteUint16LE(uint16(len(this.Name)))
	buffer.WriteBytes(this.Name)
	buffer.WriteUint16LE(uint16(this.Level))
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint64LE(uint64(this.FightNum))
	buffer.WriteUint64LE(uint64(this.ServerId))
	buffer.WriteUint32LE(uint32(this.HasWin))
}

func (this *WarInfo_Out) ByteSize() int {
	size := 68
	size += len(this.CliqueName)
	size += len(this.BattleCliqueName)
	for i := 0; i < len(this.MemberList); i++ {
		size += this.MemberList[i].ByteSize()
	}
	for i := 0; i < len(this.BattleMemerList); i++ {
		size += this.BattleMemerList[i].ByteSize()
	}
	return size
}

func (this *WarInfo_Out_MemberList) ByteSize() int {
	size := 35
	size += len(this.Name)
	return size
}

func (this *WarInfo_Out_BattleMemerList) ByteSize() int {
	size := 35
	size += len(this.Name)
	return size
}

func (this *NotifyWarReady_Out) Decode(buffer *net.Buffer) {
	this.RemainTime = int64(buffer.ReadUint64LE())
}

func (this *NotifyWarReady_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(42)
	buffer.WriteUint8(2)
	buffer.WriteUint64LE(uint64(this.RemainTime))
}

func (this *NotifyWarReady_Out) ByteSize() int {
	size := 10
	return size
}

func (this *NotifyWarStart_Out) Decode(buffer *net.Buffer) {
}

func (this *NotifyWarStart_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(42)
	buffer.WriteUint8(3)
}

func (this *NotifyWarStart_Out) ByteSize() int {
	size := 2
	return size
}

func (this *NotifyWarEnd_Out) Decode(buffer *net.Buffer) {
}

func (this *NotifyWarEnd_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(42)
	buffer.WriteUint8(4)
}

func (this *NotifyWarEnd_Out) ByteSize() int {
	size := 2
	return size
}

func (this *SignUpCliqueWar_In) Decode(buffer *net.Buffer) {
}

func (this *SignUpCliqueWar_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(42)
	buffer.WriteUint8(5)
}

func (this *SignUpCliqueWar_In) ByteSize() int {
	size := 2
	return size
}

func (this *SignUpCliqueWar_Out) Decode(buffer *net.Buffer) {
	this.Result = SignUpCliqueWarResult(buffer.ReadUint8())
}

func (this *SignUpCliqueWar_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(42)
	buffer.WriteUint8(5)
	buffer.WriteUint8(uint8(this.Result))
}

func (this *SignUpCliqueWar_Out) ByteSize() int {
	size := 3
	return size
}

func (this *StartCliqueWarBattle_In) Decode(buffer *net.Buffer) {
	this.TargetPlayerid = int64(buffer.ReadUint64LE())
}

func (this *StartCliqueWarBattle_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(42)
	buffer.WriteUint8(6)
	buffer.WriteUint64LE(uint64(this.TargetPlayerid))
}

func (this *StartCliqueWarBattle_In) ByteSize() int {
	size := 10
	return size
}

func (this *StartCliqueWarBattle_Out) Decode(buffer *net.Buffer) {
}

func (this *StartCliqueWarBattle_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(42)
	buffer.WriteUint8(6)
}

func (this *StartCliqueWarBattle_Out) ByteSize() int {
	size := 2
	return size
}

func (this *SignUpCliqueWarInfo_In) Decode(buffer *net.Buffer) {
}

func (this *SignUpCliqueWarInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(42)
	buffer.WriteUint8(7)
}

func (this *SignUpCliqueWarInfo_In) ByteSize() int {
	size := 2
	return size
}

func (this *SignUpCliqueWarInfo_Out) Decode(buffer *net.Buffer) {
	this.Result = buffer.ReadUint8() == 1
}

func (this *SignUpCliqueWarInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(42)
	buffer.WriteUint8(7)
	if this.Result {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *SignUpCliqueWarInfo_Out) ByteSize() int {
	size := 3
	return size
}
