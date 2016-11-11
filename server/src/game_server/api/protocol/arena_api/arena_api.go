package arena_api

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
	Enter(*net.Session, *Enter_In)
	GetRecords(*net.Session, *GetRecords_In)
	AwardBox(*net.Session, *AwardBox_In)
	StartBattle(*net.Session, *StartBattle_In)
	GetTopRank(*net.Session, *GetTopRank_In)
	CleanFailedCdTime(*net.Session, *CleanFailedCdTime_In)
}

type OutHandler interface {
	Enter(*net.Session, *Enter_Out)
	GetRecords(*net.Session, *GetRecords_Out)
	AwardBox(*net.Session, *AwardBox_Out)
	NotifyFailedCdTime(*net.Session, *NotifyFailedCdTime_Out)
	StartBattle(*net.Session, *StartBattle_Out)
	NotifyLoseRank(*net.Session, *NotifyLoseRank_Out)
	NotifyArenaInfo(*net.Session, *NotifyArenaInfo_Out)
	GetTopRank(*net.Session, *GetTopRank_Out)
	CleanFailedCdTime(*net.Session, *CleanFailedCdTime_Out)
}

func DecodeIn(message []byte) Request {
	var actionId = message[0]
	var buffer = net.NewBuffer(message[1:])
	switch actionId {
	case 1:
		request := new(Enter_In)
		request.Decode(buffer)
		return request
	case 2:
		request := new(GetRecords_In)
		request.Decode(buffer)
		return request
	case 3:
		request := new(AwardBox_In)
		request.Decode(buffer)
		return request
	case 5:
		request := new(StartBattle_In)
		request.Decode(buffer)
		return request
	case 8:
		request := new(GetTopRank_In)
		request.Decode(buffer)
		return request
	case 9:
		request := new(CleanFailedCdTime_In)
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
		request := new(Enter_Out)
		request.Decode(buffer)
		return request
	case 2:
		request := new(GetRecords_Out)
		request.Decode(buffer)
		return request
	case 3:
		request := new(AwardBox_Out)
		request.Decode(buffer)
		return request
	case 4:
		request := new(NotifyFailedCdTime_Out)
		request.Decode(buffer)
		return request
	case 5:
		request := new(StartBattle_Out)
		request.Decode(buffer)
		return request
	case 6:
		request := new(NotifyLoseRank_Out)
		request.Decode(buffer)
		return request
	case 7:
		request := new(NotifyArenaInfo_Out)
		request.Decode(buffer)
		return request
	case 8:
		request := new(GetTopRank_Out)
		request.Decode(buffer)
		return request
	case 9:
		request := new(CleanFailedCdTime_Out)
		request.Decode(buffer)
		return request
	}
	_ = buffer
	panic("unsupported response")
}

type NotifyArenaMode int8

const (
	NOTIFY_ARENA_MODE_ATTACKED_SUCC   NotifyArenaMode = 3
	NOTIFY_ARENA_MODE_ATTACKED_FAILED NotifyArenaMode = 4
)

type Enter_In struct {
}

func (this *Enter_In) Process(session *net.Session) {
	g_InHandler.Enter(session, this)
}

func (this *Enter_In) TypeName() string {
	return "arena.enter.in"
}

func (this *Enter_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 19, 1
}

type Enter_Out struct {
	Rank           int32                `json:"rank"`
	AwardBoxTime   int64                `json:"award_box_time"`
	DailyNum       int16                `json:"daily_num"`
	BuyTimes       int16                `json:"buy_times"`
	WinTimes       int16                `json:"win_times"`
	NewRecordNum   int16                `json:"new_record_num"`
	FailedCdTime   int64                `json:"failed_cd_time"`
	DailyAwardItem int32                `json:"daily_award_item"`
	DailyFame      int32                `json:"daily_fame"`
	AwardBox       []Enter_Out_AwardBox `json:"award_box"`
	Ranks          []Enter_Out_Ranks    `json:"ranks"`
}

type Enter_Out_AwardBox struct {
	Num  int8  `json:"num"`
	Rank int32 `json:"rank"`
}

type Enter_Out_Ranks struct {
	Pid      int64  `json:"pid"`
	Nick     []byte `json:"nick"`
	RoleId   int8   `json:"role_id"`
	Rank     int32  `json:"rank"`
	Level    int16  `json:"level"`
	FightNum int32  `json:"fight_num"`
}

func (this *Enter_Out) Process(session *net.Session) {
	g_OutHandler.Enter(session, this)
}

func (this *Enter_Out) TypeName() string {
	return "arena.enter.out"
}

func (this *Enter_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 19, 1
}

func (this *Enter_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetRecords_In struct {
}

func (this *GetRecords_In) Process(session *net.Session) {
	g_InHandler.GetRecords(session, this)
}

func (this *GetRecords_In) TypeName() string {
	return "arena.get_records.in"
}

func (this *GetRecords_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 19, 2
}

type GetRecords_Out struct {
	Records []GetRecords_Out_Records `json:"records"`
}

type GetRecords_Out_Records struct {
	Mode           int8   `json:"mode"`
	Time           int64  `json:"time"`
	OldRank        int32  `json:"old_rank"`
	NewRank        int32  `json:"new_rank"`
	FightNum       int32  `json:"fight_num"`
	TargetPid      int64  `json:"target_pid"`
	TargetNick     []byte `json:"target_nick"`
	TargetOldRank  int32  `json:"target_old_rank"`
	TargetNewRank  int32  `json:"target_new_rank"`
	TargetFightNum int32  `json:"target_fight_num"`
}

func (this *GetRecords_Out) Process(session *net.Session) {
	g_OutHandler.GetRecords(session, this)
}

func (this *GetRecords_Out) TypeName() string {
	return "arena.get_records.out"
}

func (this *GetRecords_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 19, 2
}

func (this *GetRecords_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type AwardBox_In struct {
	Num int8 `json:"num"`
}

func (this *AwardBox_In) Process(session *net.Session) {
	g_InHandler.AwardBox(session, this)
}

func (this *AwardBox_In) TypeName() string {
	return "arena.award_box.in"
}

func (this *AwardBox_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 19, 3
}

type AwardBox_Out struct {
	Result bool `json:"result"`
}

func (this *AwardBox_Out) Process(session *net.Session) {
	g_OutHandler.AwardBox(session, this)
}

func (this *AwardBox_Out) TypeName() string {
	return "arena.award_box.out"
}

func (this *AwardBox_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 19, 3
}

func (this *AwardBox_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type NotifyFailedCdTime_Out struct {
	CdTime int64 `json:"cd_time"`
}

func (this *NotifyFailedCdTime_Out) Process(session *net.Session) {
	g_OutHandler.NotifyFailedCdTime(session, this)
}

func (this *NotifyFailedCdTime_Out) TypeName() string {
	return "arena.notify_failed_cd_time.out"
}

func (this *NotifyFailedCdTime_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 19, 4
}

func (this *NotifyFailedCdTime_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type StartBattle_In struct {
	PlayerId   int64 `json:"player_id"`
	PlayerRank int32 `json:"player_rank"`
}

func (this *StartBattle_In) Process(session *net.Session) {
	g_InHandler.StartBattle(session, this)
}

func (this *StartBattle_In) TypeName() string {
	return "arena.start_battle.in"
}

func (this *StartBattle_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 19, 5
}

type StartBattle_Out struct {
}

func (this *StartBattle_Out) Process(session *net.Session) {
	g_OutHandler.StartBattle(session, this)
}

func (this *StartBattle_Out) TypeName() string {
	return "arena.start_battle.out"
}

func (this *StartBattle_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 19, 5
}

func (this *StartBattle_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type NotifyLoseRank_Out struct {
}

func (this *NotifyLoseRank_Out) Process(session *net.Session) {
	g_OutHandler.NotifyLoseRank(session, this)
}

func (this *NotifyLoseRank_Out) TypeName() string {
	return "arena.notify_lose_rank.out"
}

func (this *NotifyLoseRank_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 19, 6
}

func (this *NotifyLoseRank_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type NotifyArenaInfo_Out struct {
	NotifyType NotifyArenaMode `json:"notify_type"`
	Pid        int64           `json:"pid"`
	Nick       []byte          `json:"nick"`
	Num        int32           `json:"num"`
}

func (this *NotifyArenaInfo_Out) Process(session *net.Session) {
	g_OutHandler.NotifyArenaInfo(session, this)
}

func (this *NotifyArenaInfo_Out) TypeName() string {
	return "arena.notify_arena_info.out"
}

func (this *NotifyArenaInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 19, 7
}

func (this *NotifyArenaInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetTopRank_In struct {
}

func (this *GetTopRank_In) Process(session *net.Session) {
	g_InHandler.GetTopRank(session, this)
}

func (this *GetTopRank_In) TypeName() string {
	return "arena.get_top_rank.in"
}

func (this *GetTopRank_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 19, 8
}

type GetTopRank_Out struct {
	Top3  []GetTopRank_Out_Top3  `json:"top3"`
	Ranks []GetTopRank_Out_Ranks `json:"ranks"`
}

type GetTopRank_Out_Top3 struct {
	Openid []byte `json:"openid"`
	RoleId int8   `json:"role_id"`
}

type GetTopRank_Out_Ranks struct {
	Pid   int64  `json:"pid"`
	Rank  int32  `json:"rank"`
	Nick  []byte `json:"nick"`
	Level int16  `json:"level"`
	Trend int8   `json:"trend"`
}

func (this *GetTopRank_Out) Process(session *net.Session) {
	g_OutHandler.GetTopRank(session, this)
}

func (this *GetTopRank_Out) TypeName() string {
	return "arena.get_top_rank.out"
}

func (this *GetTopRank_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 19, 8
}

func (this *GetTopRank_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type CleanFailedCdTime_In struct {
}

func (this *CleanFailedCdTime_In) Process(session *net.Session) {
	g_InHandler.CleanFailedCdTime(session, this)
}

func (this *CleanFailedCdTime_In) TypeName() string {
	return "arena.clean_failed_cd_time.in"
}

func (this *CleanFailedCdTime_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 19, 9
}

type CleanFailedCdTime_Out struct {
	FailedCdTime int64 `json:"failed_cd_time"`
}

func (this *CleanFailedCdTime_Out) Process(session *net.Session) {
	g_OutHandler.CleanFailedCdTime(session, this)
}

func (this *CleanFailedCdTime_Out) TypeName() string {
	return "arena.clean_failed_cd_time.out"
}

func (this *CleanFailedCdTime_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 19, 9
}

func (this *CleanFailedCdTime_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

func (this *Enter_In) Decode(buffer *net.Buffer) {
}

func (this *Enter_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(19)
	buffer.WriteUint8(1)
}

func (this *Enter_In) ByteSize() int {
	size := 2
	return size
}

func (this *Enter_Out) Decode(buffer *net.Buffer) {
	this.Rank = int32(buffer.ReadUint32LE())
	this.AwardBoxTime = int64(buffer.ReadUint64LE())
	this.DailyNum = int16(buffer.ReadUint16LE())
	this.BuyTimes = int16(buffer.ReadUint16LE())
	this.WinTimes = int16(buffer.ReadUint16LE())
	this.NewRecordNum = int16(buffer.ReadUint16LE())
	this.FailedCdTime = int64(buffer.ReadUint64LE())
	this.DailyAwardItem = int32(buffer.ReadUint32LE())
	this.DailyFame = int32(buffer.ReadUint32LE())
	this.AwardBox = make([]Enter_Out_AwardBox, buffer.ReadUint8())
	for i := 0; i < len(this.AwardBox); i++ {
		this.AwardBox[i].Decode(buffer)
	}
	this.Ranks = make([]Enter_Out_Ranks, buffer.ReadUint8())
	for i := 0; i < len(this.Ranks); i++ {
		this.Ranks[i].Decode(buffer)
	}
}

func (this *Enter_Out_AwardBox) Decode(buffer *net.Buffer) {
	this.Num = int8(buffer.ReadUint8())
	this.Rank = int32(buffer.ReadUint32LE())
}

func (this *Enter_Out_Ranks) Decode(buffer *net.Buffer) {
	this.Pid = int64(buffer.ReadUint64LE())
	this.Nick = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.RoleId = int8(buffer.ReadUint8())
	this.Rank = int32(buffer.ReadUint32LE())
	this.Level = int16(buffer.ReadUint16LE())
	this.FightNum = int32(buffer.ReadUint32LE())
}

func (this *Enter_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(19)
	buffer.WriteUint8(1)
	buffer.WriteUint32LE(uint32(this.Rank))
	buffer.WriteUint64LE(uint64(this.AwardBoxTime))
	buffer.WriteUint16LE(uint16(this.DailyNum))
	buffer.WriteUint16LE(uint16(this.BuyTimes))
	buffer.WriteUint16LE(uint16(this.WinTimes))
	buffer.WriteUint16LE(uint16(this.NewRecordNum))
	buffer.WriteUint64LE(uint64(this.FailedCdTime))
	buffer.WriteUint32LE(uint32(this.DailyAwardItem))
	buffer.WriteUint32LE(uint32(this.DailyFame))
	buffer.WriteUint8(uint8(len(this.AwardBox)))
	for i := 0; i < len(this.AwardBox); i++ {
		this.AwardBox[i].Encode(buffer)
	}
	buffer.WriteUint8(uint8(len(this.Ranks)))
	for i := 0; i < len(this.Ranks); i++ {
		this.Ranks[i].Encode(buffer)
	}
}

func (this *Enter_Out_AwardBox) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(this.Num))
	buffer.WriteUint32LE(uint32(this.Rank))
}

func (this *Enter_Out_Ranks) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.Pid))
	buffer.WriteUint16LE(uint16(len(this.Nick)))
	buffer.WriteBytes(this.Nick)
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint32LE(uint32(this.Rank))
	buffer.WriteUint16LE(uint16(this.Level))
	buffer.WriteUint32LE(uint32(this.FightNum))
}

func (this *Enter_Out) ByteSize() int {
	size := 40
	size += len(this.AwardBox) * 5
	for i := 0; i < len(this.Ranks); i++ {
		size += this.Ranks[i].ByteSize()
	}
	return size
}

func (this *Enter_Out_Ranks) ByteSize() int {
	size := 21
	size += len(this.Nick)
	return size
}

func (this *GetRecords_In) Decode(buffer *net.Buffer) {
}

func (this *GetRecords_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(19)
	buffer.WriteUint8(2)
}

func (this *GetRecords_In) ByteSize() int {
	size := 2
	return size
}

func (this *GetRecords_Out) Decode(buffer *net.Buffer) {
	this.Records = make([]GetRecords_Out_Records, buffer.ReadUint8())
	for i := 0; i < len(this.Records); i++ {
		this.Records[i].Decode(buffer)
	}
}

func (this *GetRecords_Out_Records) Decode(buffer *net.Buffer) {
	this.Mode = int8(buffer.ReadUint8())
	this.Time = int64(buffer.ReadUint64LE())
	this.OldRank = int32(buffer.ReadUint32LE())
	this.NewRank = int32(buffer.ReadUint32LE())
	this.FightNum = int32(buffer.ReadUint32LE())
	this.TargetPid = int64(buffer.ReadUint64LE())
	this.TargetNick = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.TargetOldRank = int32(buffer.ReadUint32LE())
	this.TargetNewRank = int32(buffer.ReadUint32LE())
	this.TargetFightNum = int32(buffer.ReadUint32LE())
}

func (this *GetRecords_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(19)
	buffer.WriteUint8(2)
	buffer.WriteUint8(uint8(len(this.Records)))
	for i := 0; i < len(this.Records); i++ {
		this.Records[i].Encode(buffer)
	}
}

func (this *GetRecords_Out_Records) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(this.Mode))
	buffer.WriteUint64LE(uint64(this.Time))
	buffer.WriteUint32LE(uint32(this.OldRank))
	buffer.WriteUint32LE(uint32(this.NewRank))
	buffer.WriteUint32LE(uint32(this.FightNum))
	buffer.WriteUint64LE(uint64(this.TargetPid))
	buffer.WriteUint16LE(uint16(len(this.TargetNick)))
	buffer.WriteBytes(this.TargetNick)
	buffer.WriteUint32LE(uint32(this.TargetOldRank))
	buffer.WriteUint32LE(uint32(this.TargetNewRank))
	buffer.WriteUint32LE(uint32(this.TargetFightNum))
}

func (this *GetRecords_Out) ByteSize() int {
	size := 3
	for i := 0; i < len(this.Records); i++ {
		size += this.Records[i].ByteSize()
	}
	return size
}

func (this *GetRecords_Out_Records) ByteSize() int {
	size := 43
	size += len(this.TargetNick)
	return size
}

func (this *AwardBox_In) Decode(buffer *net.Buffer) {
	this.Num = int8(buffer.ReadUint8())
}

func (this *AwardBox_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(19)
	buffer.WriteUint8(3)
	buffer.WriteUint8(uint8(this.Num))
}

func (this *AwardBox_In) ByteSize() int {
	size := 3
	return size
}

func (this *AwardBox_Out) Decode(buffer *net.Buffer) {
	this.Result = buffer.ReadUint8() == 1
}

func (this *AwardBox_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(19)
	buffer.WriteUint8(3)
	if this.Result {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *AwardBox_Out) ByteSize() int {
	size := 3
	return size
}

func (this *NotifyFailedCdTime_Out) Decode(buffer *net.Buffer) {
	this.CdTime = int64(buffer.ReadUint64LE())
}

func (this *NotifyFailedCdTime_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(19)
	buffer.WriteUint8(4)
	buffer.WriteUint64LE(uint64(this.CdTime))
}

func (this *NotifyFailedCdTime_Out) ByteSize() int {
	size := 10
	return size
}

func (this *StartBattle_In) Decode(buffer *net.Buffer) {
	this.PlayerId = int64(buffer.ReadUint64LE())
	this.PlayerRank = int32(buffer.ReadUint32LE())
}

func (this *StartBattle_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(19)
	buffer.WriteUint8(5)
	buffer.WriteUint64LE(uint64(this.PlayerId))
	buffer.WriteUint32LE(uint32(this.PlayerRank))
}

func (this *StartBattle_In) ByteSize() int {
	size := 14
	return size
}

func (this *StartBattle_Out) Decode(buffer *net.Buffer) {
}

func (this *StartBattle_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(19)
	buffer.WriteUint8(5)
}

func (this *StartBattle_Out) ByteSize() int {
	size := 2
	return size
}

func (this *NotifyLoseRank_Out) Decode(buffer *net.Buffer) {
}

func (this *NotifyLoseRank_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(19)
	buffer.WriteUint8(6)
}

func (this *NotifyLoseRank_Out) ByteSize() int {
	size := 2
	return size
}

func (this *NotifyArenaInfo_Out) Decode(buffer *net.Buffer) {
	this.NotifyType = NotifyArenaMode(buffer.ReadUint8())
	this.Pid = int64(buffer.ReadUint64LE())
	this.Nick = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Num = int32(buffer.ReadUint32LE())
}

func (this *NotifyArenaInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(19)
	buffer.WriteUint8(7)
	buffer.WriteUint8(uint8(this.NotifyType))
	buffer.WriteUint64LE(uint64(this.Pid))
	buffer.WriteUint16LE(uint16(len(this.Nick)))
	buffer.WriteBytes(this.Nick)
	buffer.WriteUint32LE(uint32(this.Num))
}

func (this *NotifyArenaInfo_Out) ByteSize() int {
	size := 17
	size += len(this.Nick)
	return size
}

func (this *GetTopRank_In) Decode(buffer *net.Buffer) {
}

func (this *GetTopRank_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(19)
	buffer.WriteUint8(8)
}

func (this *GetTopRank_In) ByteSize() int {
	size := 2
	return size
}

func (this *GetTopRank_Out) Decode(buffer *net.Buffer) {
	this.Top3 = make([]GetTopRank_Out_Top3, buffer.ReadUint8())
	for i := 0; i < len(this.Top3); i++ {
		this.Top3[i].Decode(buffer)
	}
	this.Ranks = make([]GetTopRank_Out_Ranks, buffer.ReadUint8())
	for i := 0; i < len(this.Ranks); i++ {
		this.Ranks[i].Decode(buffer)
	}
}

func (this *GetTopRank_Out_Top3) Decode(buffer *net.Buffer) {
	this.Openid = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.RoleId = int8(buffer.ReadUint8())
}

func (this *GetTopRank_Out_Ranks) Decode(buffer *net.Buffer) {
	this.Pid = int64(buffer.ReadUint64LE())
	this.Rank = int32(buffer.ReadUint32LE())
	this.Nick = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Level = int16(buffer.ReadUint16LE())
	this.Trend = int8(buffer.ReadUint8())
}

func (this *GetTopRank_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(19)
	buffer.WriteUint8(8)
	buffer.WriteUint8(uint8(len(this.Top3)))
	for i := 0; i < len(this.Top3); i++ {
		this.Top3[i].Encode(buffer)
	}
	buffer.WriteUint8(uint8(len(this.Ranks)))
	for i := 0; i < len(this.Ranks); i++ {
		this.Ranks[i].Encode(buffer)
	}
}

func (this *GetTopRank_Out_Top3) Encode(buffer *net.Buffer) {
	buffer.WriteUint16LE(uint16(len(this.Openid)))
	buffer.WriteBytes(this.Openid)
	buffer.WriteUint8(uint8(this.RoleId))
}

func (this *GetTopRank_Out_Ranks) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.Pid))
	buffer.WriteUint32LE(uint32(this.Rank))
	buffer.WriteUint16LE(uint16(len(this.Nick)))
	buffer.WriteBytes(this.Nick)
	buffer.WriteUint16LE(uint16(this.Level))
	buffer.WriteUint8(uint8(this.Trend))
}

func (this *GetTopRank_Out) ByteSize() int {
	size := 4
	for i := 0; i < len(this.Top3); i++ {
		size += this.Top3[i].ByteSize()
	}
	for i := 0; i < len(this.Ranks); i++ {
		size += this.Ranks[i].ByteSize()
	}
	return size
}

func (this *GetTopRank_Out_Top3) ByteSize() int {
	size := 3
	size += len(this.Openid)
	return size
}

func (this *GetTopRank_Out_Ranks) ByteSize() int {
	size := 17
	size += len(this.Nick)
	return size
}

func (this *CleanFailedCdTime_In) Decode(buffer *net.Buffer) {
}

func (this *CleanFailedCdTime_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(19)
	buffer.WriteUint8(9)
}

func (this *CleanFailedCdTime_In) ByteSize() int {
	size := 2
	return size
}

func (this *CleanFailedCdTime_Out) Decode(buffer *net.Buffer) {
	this.FailedCdTime = int64(buffer.ReadUint64LE())
}

func (this *CleanFailedCdTime_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(19)
	buffer.WriteUint8(9)
	buffer.WriteUint64LE(uint64(this.FailedCdTime))
}

func (this *CleanFailedCdTime_Out) ByteSize() int {
	size := 10
	return size
}
