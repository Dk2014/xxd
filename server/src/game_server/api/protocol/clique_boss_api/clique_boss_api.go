package clique_boss_api

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
	BossInfo(*net.Session, *BossInfo_In)
	CliqueBossRefreshChallenge(*net.Session, *CliqueBossRefreshChallenge_In)
}

type OutHandler interface {
	BossInfo(*net.Session, *BossInfo_Out)
	CliqueNotifyBossOpen(*net.Session, *CliqueNotifyBossOpen_Out)
	CliqueNotifyBossDead(*net.Session, *CliqueNotifyBossDead_Out)
	CliqueBattleBossAwardInfo(*net.Session, *CliqueBattleBossAwardInfo_Out)
	CliqueBossRefreshChallenge(*net.Session, *CliqueBossRefreshChallenge_Out)
	CliqueNotifyBossClose(*net.Session, *CliqueNotifyBossClose_Out)
}

func DecodeIn(message []byte) Request {
	var actionId = message[0]
	var buffer = net.NewBuffer(message[1:])
	switch actionId {
	case 1:
		request := new(BossInfo_In)
		request.Decode(buffer)
		return request
	case 5:
		request := new(CliqueBossRefreshChallenge_In)
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
		request := new(BossInfo_Out)
		request.Decode(buffer)
		return request
	case 2:
		request := new(CliqueNotifyBossOpen_Out)
		request.Decode(buffer)
		return request
	case 3:
		request := new(CliqueNotifyBossDead_Out)
		request.Decode(buffer)
		return request
	case 4:
		request := new(CliqueBattleBossAwardInfo_Out)
		request.Decode(buffer)
		return request
	case 5:
		request := new(CliqueBossRefreshChallenge_Out)
		request.Decode(buffer)
		return request
	case 6:
		request := new(CliqueNotifyBossClose_Out)
		request.Decode(buffer)
		return request
	}
	_ = buffer
	panic("unsupported response")
}

type RefreshChallengeResult int8

const (
	REFRESH_CHALLENGE_RESULT_FAILED              RefreshChallengeResult = 0
	REFRESH_CHALLENGE_RESULT_SUCCEED             RefreshChallengeResult = 1
	REFRESH_CHALLENGE_RESULT_DO_NOT_NEED_REFRESH RefreshChallengeResult = 2
)

type BossInfo_In struct {
}

func (this *BossInfo_In) Process(session *net.Session) {
	g_InHandler.BossInfo(session, this)
}

func (this *BossInfo_In) TypeName() string {
	return "clique_boss.boss_info.in"
}

func (this *BossInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 41, 1
}

type BossInfo_Out struct {
	Hp              int64
	MaxHp           int64
	BossLevel       int16
	StartTimestamp  int64
	DeadTimestamp   int64
	BattleTimestamp int64
	BattleNum       int32
	DeadNum         int64
	RanksPlayer     []BossInfo_Out_RanksPlayer
	RanksClique     []BossInfo_Out_RanksClique
}

type BossInfo_Out_RanksPlayer struct {
	Rank     int64
	Pid      int64
	Name     []byte
	Hurt     int64
	ServerId int8
}

type BossInfo_Out_RanksClique struct {
	Rank     int64
	Cliqueid int64
	Name     []byte
	Hurt     int64
	ServerId int8
}

func (this *BossInfo_Out) Process(session *net.Session) {
	g_OutHandler.BossInfo(session, this)
}

func (this *BossInfo_Out) TypeName() string {
	return "clique_boss.boss_info.out"
}

func (this *BossInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 41, 1
}

func (this *BossInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type CliqueNotifyBossOpen_Out struct {
}

func (this *CliqueNotifyBossOpen_Out) Process(session *net.Session) {
	g_OutHandler.CliqueNotifyBossOpen(session, this)
}

func (this *CliqueNotifyBossOpen_Out) TypeName() string {
	return "clique_boss.clique_notify_boss_open.out"
}

func (this *CliqueNotifyBossOpen_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 41, 2
}

func (this *CliqueNotifyBossOpen_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type CliqueNotifyBossDead_Out struct {
	Ranks       []CliqueNotifyBossDead_Out_Ranks
	CliqueRanks []CliqueNotifyBossDead_Out_CliqueRanks
}

type CliqueNotifyBossDead_Out_Ranks struct {
	Rank int64
	Pid  int64
	Name []byte
	Sid  int8
}

type CliqueNotifyBossDead_Out_CliqueRanks struct {
	Rank     int64
	CliqueId int64
	Name     []byte
	Sid      int8
}

func (this *CliqueNotifyBossDead_Out) Process(session *net.Session) {
	g_OutHandler.CliqueNotifyBossDead(session, this)
}

func (this *CliqueNotifyBossDead_Out) TypeName() string {
	return "clique_boss.clique_notify_boss_dead.out"
}

func (this *CliqueNotifyBossDead_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 41, 3
}

func (this *CliqueNotifyBossDead_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type CliqueBattleBossAwardInfo_Out struct {
	Hurt int64
}

func (this *CliqueBattleBossAwardInfo_Out) Process(session *net.Session) {
	g_OutHandler.CliqueBattleBossAwardInfo(session, this)
}

func (this *CliqueBattleBossAwardInfo_Out) TypeName() string {
	return "clique_boss.clique_battle_boss_award_info.out"
}

func (this *CliqueBattleBossAwardInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 41, 4
}

func (this *CliqueBattleBossAwardInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type CliqueBossRefreshChallenge_In struct {
}

func (this *CliqueBossRefreshChallenge_In) Process(session *net.Session) {
	g_InHandler.CliqueBossRefreshChallenge(session, this)
}

func (this *CliqueBossRefreshChallenge_In) TypeName() string {
	return "clique_boss.clique_boss_refresh_challenge.in"
}

func (this *CliqueBossRefreshChallenge_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 41, 5
}

type CliqueBossRefreshChallenge_Out struct {
	Result RefreshChallengeResult
}

func (this *CliqueBossRefreshChallenge_Out) Process(session *net.Session) {
	g_OutHandler.CliqueBossRefreshChallenge(session, this)
}

func (this *CliqueBossRefreshChallenge_Out) TypeName() string {
	return "clique_boss.clique_boss_refresh_challenge.out"
}

func (this *CliqueBossRefreshChallenge_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 41, 5
}

func (this *CliqueBossRefreshChallenge_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type CliqueNotifyBossClose_Out struct {
}

func (this *CliqueNotifyBossClose_Out) Process(session *net.Session) {
	g_OutHandler.CliqueNotifyBossClose(session, this)
}

func (this *CliqueNotifyBossClose_Out) TypeName() string {
	return "clique_boss.clique_notify_boss_close.out"
}

func (this *CliqueNotifyBossClose_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 41, 6
}

func (this *CliqueNotifyBossClose_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

func (this *BossInfo_In) Decode(buffer *net.Buffer) {
}

func (this *BossInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(41)
	buffer.WriteUint8(1)
}

func (this *BossInfo_In) ByteSize() int {
	size := 2
	return size
}

func (this *BossInfo_Out) Decode(buffer *net.Buffer) {
	this.Hp = int64(buffer.ReadUint64LE())
	this.MaxHp = int64(buffer.ReadUint64LE())
	this.BossLevel = int16(buffer.ReadUint16LE())
	this.StartTimestamp = int64(buffer.ReadUint64LE())
	this.DeadTimestamp = int64(buffer.ReadUint64LE())
	this.BattleTimestamp = int64(buffer.ReadUint64LE())
	this.BattleNum = int32(buffer.ReadUint32LE())
	this.DeadNum = int64(buffer.ReadUint64LE())
	this.RanksPlayer = make([]BossInfo_Out_RanksPlayer, buffer.ReadUint8())
	for i := 0; i < len(this.RanksPlayer); i++ {
		this.RanksPlayer[i].Decode(buffer)
	}
	this.RanksClique = make([]BossInfo_Out_RanksClique, buffer.ReadUint8())
	for i := 0; i < len(this.RanksClique); i++ {
		this.RanksClique[i].Decode(buffer)
	}
}

func (this *BossInfo_Out_RanksPlayer) Decode(buffer *net.Buffer) {
	this.Rank = int64(buffer.ReadUint64LE())
	this.Pid = int64(buffer.ReadUint64LE())
	this.Name = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Hurt = int64(buffer.ReadUint64LE())
	this.ServerId = int8(buffer.ReadUint8())
}

func (this *BossInfo_Out_RanksClique) Decode(buffer *net.Buffer) {
	this.Rank = int64(buffer.ReadUint64LE())
	this.Cliqueid = int64(buffer.ReadUint64LE())
	this.Name = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Hurt = int64(buffer.ReadUint64LE())
	this.ServerId = int8(buffer.ReadUint8())
}

func (this *BossInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(41)
	buffer.WriteUint8(1)
	buffer.WriteUint64LE(uint64(this.Hp))
	buffer.WriteUint64LE(uint64(this.MaxHp))
	buffer.WriteUint16LE(uint16(this.BossLevel))
	buffer.WriteUint64LE(uint64(this.StartTimestamp))
	buffer.WriteUint64LE(uint64(this.DeadTimestamp))
	buffer.WriteUint64LE(uint64(this.BattleTimestamp))
	buffer.WriteUint32LE(uint32(this.BattleNum))
	buffer.WriteUint64LE(uint64(this.DeadNum))
	buffer.WriteUint8(uint8(len(this.RanksPlayer)))
	for i := 0; i < len(this.RanksPlayer); i++ {
		this.RanksPlayer[i].Encode(buffer)
	}
	buffer.WriteUint8(uint8(len(this.RanksClique)))
	for i := 0; i < len(this.RanksClique); i++ {
		this.RanksClique[i].Encode(buffer)
	}
}

func (this *BossInfo_Out_RanksPlayer) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.Rank))
	buffer.WriteUint64LE(uint64(this.Pid))
	buffer.WriteUint16LE(uint16(len(this.Name)))
	buffer.WriteBytes(this.Name)
	buffer.WriteUint64LE(uint64(this.Hurt))
	buffer.WriteUint8(uint8(this.ServerId))
}

func (this *BossInfo_Out_RanksClique) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.Rank))
	buffer.WriteUint64LE(uint64(this.Cliqueid))
	buffer.WriteUint16LE(uint16(len(this.Name)))
	buffer.WriteBytes(this.Name)
	buffer.WriteUint64LE(uint64(this.Hurt))
	buffer.WriteUint8(uint8(this.ServerId))
}

func (this *BossInfo_Out) ByteSize() int {
	size := 58
	for i := 0; i < len(this.RanksPlayer); i++ {
		size += this.RanksPlayer[i].ByteSize()
	}
	for i := 0; i < len(this.RanksClique); i++ {
		size += this.RanksClique[i].ByteSize()
	}
	return size
}

func (this *BossInfo_Out_RanksPlayer) ByteSize() int {
	size := 27
	size += len(this.Name)
	return size
}

func (this *BossInfo_Out_RanksClique) ByteSize() int {
	size := 27
	size += len(this.Name)
	return size
}

func (this *CliqueNotifyBossOpen_Out) Decode(buffer *net.Buffer) {
}

func (this *CliqueNotifyBossOpen_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(41)
	buffer.WriteUint8(2)
}

func (this *CliqueNotifyBossOpen_Out) ByteSize() int {
	size := 2
	return size
}

func (this *CliqueNotifyBossDead_Out) Decode(buffer *net.Buffer) {
	this.Ranks = make([]CliqueNotifyBossDead_Out_Ranks, buffer.ReadUint8())
	for i := 0; i < len(this.Ranks); i++ {
		this.Ranks[i].Decode(buffer)
	}
	this.CliqueRanks = make([]CliqueNotifyBossDead_Out_CliqueRanks, buffer.ReadUint8())
	for i := 0; i < len(this.CliqueRanks); i++ {
		this.CliqueRanks[i].Decode(buffer)
	}
}

func (this *CliqueNotifyBossDead_Out_Ranks) Decode(buffer *net.Buffer) {
	this.Rank = int64(buffer.ReadUint64LE())
	this.Pid = int64(buffer.ReadUint64LE())
	this.Name = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Sid = int8(buffer.ReadUint8())
}

func (this *CliqueNotifyBossDead_Out_CliqueRanks) Decode(buffer *net.Buffer) {
	this.Rank = int64(buffer.ReadUint64LE())
	this.CliqueId = int64(buffer.ReadUint64LE())
	this.Name = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Sid = int8(buffer.ReadUint8())
}

func (this *CliqueNotifyBossDead_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(41)
	buffer.WriteUint8(3)
	buffer.WriteUint8(uint8(len(this.Ranks)))
	for i := 0; i < len(this.Ranks); i++ {
		this.Ranks[i].Encode(buffer)
	}
	buffer.WriteUint8(uint8(len(this.CliqueRanks)))
	for i := 0; i < len(this.CliqueRanks); i++ {
		this.CliqueRanks[i].Encode(buffer)
	}
}

func (this *CliqueNotifyBossDead_Out_Ranks) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.Rank))
	buffer.WriteUint64LE(uint64(this.Pid))
	buffer.WriteUint16LE(uint16(len(this.Name)))
	buffer.WriteBytes(this.Name)
	buffer.WriteUint8(uint8(this.Sid))
}

func (this *CliqueNotifyBossDead_Out_CliqueRanks) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.Rank))
	buffer.WriteUint64LE(uint64(this.CliqueId))
	buffer.WriteUint16LE(uint16(len(this.Name)))
	buffer.WriteBytes(this.Name)
	buffer.WriteUint8(uint8(this.Sid))
}

func (this *CliqueNotifyBossDead_Out) ByteSize() int {
	size := 4
	for i := 0; i < len(this.Ranks); i++ {
		size += this.Ranks[i].ByteSize()
	}
	for i := 0; i < len(this.CliqueRanks); i++ {
		size += this.CliqueRanks[i].ByteSize()
	}
	return size
}

func (this *CliqueNotifyBossDead_Out_Ranks) ByteSize() int {
	size := 19
	size += len(this.Name)
	return size
}

func (this *CliqueNotifyBossDead_Out_CliqueRanks) ByteSize() int {
	size := 19
	size += len(this.Name)
	return size
}

func (this *CliqueBattleBossAwardInfo_Out) Decode(buffer *net.Buffer) {
	this.Hurt = int64(buffer.ReadUint64LE())
}

func (this *CliqueBattleBossAwardInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(41)
	buffer.WriteUint8(4)
	buffer.WriteUint64LE(uint64(this.Hurt))
}

func (this *CliqueBattleBossAwardInfo_Out) ByteSize() int {
	size := 10
	return size
}

func (this *CliqueBossRefreshChallenge_In) Decode(buffer *net.Buffer) {
}

func (this *CliqueBossRefreshChallenge_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(41)
	buffer.WriteUint8(5)
}

func (this *CliqueBossRefreshChallenge_In) ByteSize() int {
	size := 2
	return size
}

func (this *CliqueBossRefreshChallenge_Out) Decode(buffer *net.Buffer) {
	this.Result = RefreshChallengeResult(buffer.ReadUint8())
}

func (this *CliqueBossRefreshChallenge_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(41)
	buffer.WriteUint8(5)
	buffer.WriteUint8(uint8(this.Result))
}

func (this *CliqueBossRefreshChallenge_Out) ByteSize() int {
	size := 3
	return size
}

func (this *CliqueNotifyBossClose_Out) Decode(buffer *net.Buffer) {
}

func (this *CliqueNotifyBossClose_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(41)
	buffer.WriteUint8(6)
}

func (this *CliqueNotifyBossClose_Out) ByteSize() int {
	size := 2
	return size
}
