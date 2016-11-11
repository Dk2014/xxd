package rainbow_api

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
	Info(*net.Session, *Info_In)
	Reset(*net.Session, *Reset_In)
	AwardInfo(*net.Session, *AwardInfo_In)
	TakeAward(*net.Session, *TakeAward_In)
	JumpToSegment(*net.Session, *JumpToSegment_In)
	AutoFight(*net.Session, *AutoFight_In)
}

type OutHandler interface {
	Info(*net.Session, *Info_Out)
	Reset(*net.Session, *Reset_Out)
	AwardInfo(*net.Session, *AwardInfo_Out)
	TakeAward(*net.Session, *TakeAward_Out)
	JumpToSegment(*net.Session, *JumpToSegment_Out)
	AutoFight(*net.Session, *AutoFight_Out)
}

func DecodeIn(message []byte) Request {
	var actionId = message[0]
	var buffer = net.NewBuffer(message[1:])
	switch actionId {
	case 1:
		request := new(Info_In)
		request.Decode(buffer)
		return request
	case 2:
		request := new(Reset_In)
		request.Decode(buffer)
		return request
	case 3:
		request := new(AwardInfo_In)
		request.Decode(buffer)
		return request
	case 4:
		request := new(TakeAward_In)
		request.Decode(buffer)
		return request
	case 5:
		request := new(JumpToSegment_In)
		request.Decode(buffer)
		return request
	case 6:
		request := new(AutoFight_In)
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
		request := new(Info_Out)
		request.Decode(buffer)
		return request
	case 2:
		request := new(Reset_Out)
		request.Decode(buffer)
		return request
	case 3:
		request := new(AwardInfo_Out)
		request.Decode(buffer)
		return request
	case 4:
		request := new(TakeAward_Out)
		request.Decode(buffer)
		return request
	case 5:
		request := new(JumpToSegment_Out)
		request.Decode(buffer)
		return request
	case 6:
		request := new(AutoFight_Out)
		request.Decode(buffer)
		return request
	}
	_ = buffer
	panic("unsupported response")
}

type Info_In struct {
}

func (this *Info_In) Process(session *net.Session) {
	g_InHandler.Info(session, this)
}

func (this *Info_In) TypeName() string {
	return "rainbow.info.in"
}

func (this *Info_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 23, 1
}

type Info_Out struct {
	SegmentNum        int16 `json:"segment_num"`
	LevelOrder        int8  `json:"level_order"`
	Status            int8  `json:"status"`
	ResetNum          int8  `json:"reset_num"`
	MaxSegmentCanJump int16 `json:"max_segment_can_jump"`
	MaxPassSegment    int16 `json:"max_pass_segment"`
	AutoFightNum      int8  `json:"auto_fight_num"`
	BuyTimes          int16 `json:"buy_times"`
}

func (this *Info_Out) Process(session *net.Session) {
	g_OutHandler.Info(session, this)
}

func (this *Info_Out) TypeName() string {
	return "rainbow.info.out"
}

func (this *Info_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 23, 1
}

func (this *Info_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type Reset_In struct {
}

func (this *Reset_In) Process(session *net.Session) {
	g_InHandler.Reset(session, this)
}

func (this *Reset_In) TypeName() string {
	return "rainbow.reset.in"
}

func (this *Reset_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 23, 2
}

type Reset_Out struct {
}

func (this *Reset_Out) Process(session *net.Session) {
	g_OutHandler.Reset(session, this)
}

func (this *Reset_Out) TypeName() string {
	return "rainbow.reset.out"
}

func (this *Reset_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 23, 2
}

func (this *Reset_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type AwardInfo_In struct {
}

func (this *AwardInfo_In) Process(session *net.Session) {
	g_InHandler.AwardInfo(session, this)
}

func (this *AwardInfo_In) TypeName() string {
	return "rainbow.award_info.in"
}

func (this *AwardInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 23, 3
}

type AwardInfo_Out struct {
	Award []AwardInfo_Out_Award `json:"award"`
}

type AwardInfo_Out_Award struct {
	Order int8 `json:"order"`
}

func (this *AwardInfo_Out) Process(session *net.Session) {
	g_OutHandler.AwardInfo(session, this)
}

func (this *AwardInfo_Out) TypeName() string {
	return "rainbow.award_info.out"
}

func (this *AwardInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 23, 3
}

func (this *AwardInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type TakeAward_In struct {
	Pos1 int8 `json:"pos1"`
	Pos2 int8 `json:"pos2"`
}

func (this *TakeAward_In) Process(session *net.Session) {
	g_InHandler.TakeAward(session, this)
}

func (this *TakeAward_In) TypeName() string {
	return "rainbow.take_award.in"
}

func (this *TakeAward_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 23, 4
}

type TakeAward_Out struct {
	NextLevel bool `json:"next_level"`
}

func (this *TakeAward_Out) Process(session *net.Session) {
	g_OutHandler.TakeAward(session, this)
}

func (this *TakeAward_Out) TypeName() string {
	return "rainbow.take_award.out"
}

func (this *TakeAward_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 23, 4
}

func (this *TakeAward_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type JumpToSegment_In struct {
	Segment int16 `json:"segment"`
}

func (this *JumpToSegment_In) Process(session *net.Session) {
	g_InHandler.JumpToSegment(session, this)
}

func (this *JumpToSegment_In) TypeName() string {
	return "rainbow.jump_to_segment.in"
}

func (this *JumpToSegment_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 23, 5
}

type JumpToSegment_Out struct {
}

func (this *JumpToSegment_Out) Process(session *net.Session) {
	g_OutHandler.JumpToSegment(session, this)
}

func (this *JumpToSegment_Out) TypeName() string {
	return "rainbow.jump_to_segment.out"
}

func (this *JumpToSegment_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 23, 5
}

func (this *JumpToSegment_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type AutoFight_In struct {
	Segment int16 `json:"segment"`
}

func (this *AutoFight_In) Process(session *net.Session) {
	g_InHandler.AutoFight(session, this)
}

func (this *AutoFight_In) TypeName() string {
	return "rainbow.auto_fight.in"
}

func (this *AutoFight_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 23, 6
}

type AutoFight_Out struct {
	AwardCoin    int32 `json:"awardCoin"`
	AwardExp     int32 `json:"awardExp"`
	AwardBoxPos1 int8  `json:"awardBoxPos1"`
	AwardBoxPos2 int8  `json:"awardBoxPos2"`
}

func (this *AutoFight_Out) Process(session *net.Session) {
	g_OutHandler.AutoFight(session, this)
}

func (this *AutoFight_Out) TypeName() string {
	return "rainbow.auto_fight.out"
}

func (this *AutoFight_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 23, 6
}

func (this *AutoFight_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

func (this *Info_In) Decode(buffer *net.Buffer) {
}

func (this *Info_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(23)
	buffer.WriteUint8(1)
}

func (this *Info_In) ByteSize() int {
	size := 2
	return size
}

func (this *Info_Out) Decode(buffer *net.Buffer) {
	this.SegmentNum = int16(buffer.ReadUint16LE())
	this.LevelOrder = int8(buffer.ReadUint8())
	this.Status = int8(buffer.ReadUint8())
	this.ResetNum = int8(buffer.ReadUint8())
	this.MaxSegmentCanJump = int16(buffer.ReadUint16LE())
	this.MaxPassSegment = int16(buffer.ReadUint16LE())
	this.AutoFightNum = int8(buffer.ReadUint8())
	this.BuyTimes = int16(buffer.ReadUint16LE())
}

func (this *Info_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(23)
	buffer.WriteUint8(1)
	buffer.WriteUint16LE(uint16(this.SegmentNum))
	buffer.WriteUint8(uint8(this.LevelOrder))
	buffer.WriteUint8(uint8(this.Status))
	buffer.WriteUint8(uint8(this.ResetNum))
	buffer.WriteUint16LE(uint16(this.MaxSegmentCanJump))
	buffer.WriteUint16LE(uint16(this.MaxPassSegment))
	buffer.WriteUint8(uint8(this.AutoFightNum))
	buffer.WriteUint16LE(uint16(this.BuyTimes))
}

func (this *Info_Out) ByteSize() int {
	size := 14
	return size
}

func (this *Reset_In) Decode(buffer *net.Buffer) {
}

func (this *Reset_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(23)
	buffer.WriteUint8(2)
}

func (this *Reset_In) ByteSize() int {
	size := 2
	return size
}

func (this *Reset_Out) Decode(buffer *net.Buffer) {
}

func (this *Reset_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(23)
	buffer.WriteUint8(2)
}

func (this *Reset_Out) ByteSize() int {
	size := 2
	return size
}

func (this *AwardInfo_In) Decode(buffer *net.Buffer) {
}

func (this *AwardInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(23)
	buffer.WriteUint8(3)
}

func (this *AwardInfo_In) ByteSize() int {
	size := 2
	return size
}

func (this *AwardInfo_Out) Decode(buffer *net.Buffer) {
	this.Award = make([]AwardInfo_Out_Award, buffer.ReadUint8())
	for i := 0; i < len(this.Award); i++ {
		this.Award[i].Decode(buffer)
	}
}

func (this *AwardInfo_Out_Award) Decode(buffer *net.Buffer) {
	this.Order = int8(buffer.ReadUint8())
}

func (this *AwardInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(23)
	buffer.WriteUint8(3)
	buffer.WriteUint8(uint8(len(this.Award)))
	for i := 0; i < len(this.Award); i++ {
		this.Award[i].Encode(buffer)
	}
}

func (this *AwardInfo_Out_Award) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(this.Order))
}

func (this *AwardInfo_Out) ByteSize() int {
	size := 3
	size += len(this.Award) * 1
	return size
}

func (this *TakeAward_In) Decode(buffer *net.Buffer) {
	this.Pos1 = int8(buffer.ReadUint8())
	this.Pos2 = int8(buffer.ReadUint8())
}

func (this *TakeAward_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(23)
	buffer.WriteUint8(4)
	buffer.WriteUint8(uint8(this.Pos1))
	buffer.WriteUint8(uint8(this.Pos2))
}

func (this *TakeAward_In) ByteSize() int {
	size := 4
	return size
}

func (this *TakeAward_Out) Decode(buffer *net.Buffer) {
	this.NextLevel = buffer.ReadUint8() == 1
}

func (this *TakeAward_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(23)
	buffer.WriteUint8(4)
	if this.NextLevel {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *TakeAward_Out) ByteSize() int {
	size := 3
	return size
}

func (this *JumpToSegment_In) Decode(buffer *net.Buffer) {
	this.Segment = int16(buffer.ReadUint16LE())
}

func (this *JumpToSegment_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(23)
	buffer.WriteUint8(5)
	buffer.WriteUint16LE(uint16(this.Segment))
}

func (this *JumpToSegment_In) ByteSize() int {
	size := 4
	return size
}

func (this *JumpToSegment_Out) Decode(buffer *net.Buffer) {
}

func (this *JumpToSegment_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(23)
	buffer.WriteUint8(5)
}

func (this *JumpToSegment_Out) ByteSize() int {
	size := 2
	return size
}

func (this *AutoFight_In) Decode(buffer *net.Buffer) {
	this.Segment = int16(buffer.ReadUint16LE())
}

func (this *AutoFight_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(23)
	buffer.WriteUint8(6)
	buffer.WriteUint16LE(uint16(this.Segment))
}

func (this *AutoFight_In) ByteSize() int {
	size := 4
	return size
}

func (this *AutoFight_Out) Decode(buffer *net.Buffer) {
	this.AwardCoin = int32(buffer.ReadUint32LE())
	this.AwardExp = int32(buffer.ReadUint32LE())
	this.AwardBoxPos1 = int8(buffer.ReadUint8())
	this.AwardBoxPos2 = int8(buffer.ReadUint8())
}

func (this *AutoFight_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(23)
	buffer.WriteUint8(6)
	buffer.WriteUint32LE(uint32(this.AwardCoin))
	buffer.WriteUint32LE(uint32(this.AwardExp))
	buffer.WriteUint8(uint8(this.AwardBoxPos1))
	buffer.WriteUint8(uint8(this.AwardBoxPos2))
}

func (this *AutoFight_Out) ByteSize() int {
	size := 12
	return size
}
