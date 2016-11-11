package town_api

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
	Leave(*net.Session, *Leave_In)
	Move(*net.Session, *Move_In)
	TalkedNpcList(*net.Session, *TalkedNpcList_In)
	NpcTalk(*net.Session, *NpcTalk_In)
	ListOpenedTownTreasures(*net.Session, *ListOpenedTownTreasures_In)
	TakeTownTreasures(*net.Session, *TakeTownTreasures_In)
}

type OutHandler interface {
	Enter(*net.Session, *Enter_Out)
	Leave(*net.Session, *Leave_Out)
	Move(*net.Session, *Move_Out)
	TalkedNpcList(*net.Session, *TalkedNpcList_Out)
	NpcTalk(*net.Session, *NpcTalk_Out)
	NotifyTownPlayers(*net.Session, *NotifyTownPlayers_Out)
	UpdateTownPlayer(*net.Session, *UpdateTownPlayer_Out)
	UpdateTownPlayerMeditationState(*net.Session, *UpdateTownPlayerMeditationState_Out)
	ListOpenedTownTreasures(*net.Session, *ListOpenedTownTreasures_Out)
	TakeTownTreasures(*net.Session, *TakeTownTreasures_Out)
}

func DecodeIn(message []byte) Request {
	var actionId = message[0]
	var buffer = net.NewBuffer(message[1:])
	switch actionId {
	case 0:
		request := new(Enter_In)
		request.Decode(buffer)
		return request
	case 1:
		request := new(Leave_In)
		request.Decode(buffer)
		return request
	case 2:
		request := new(Move_In)
		request.Decode(buffer)
		return request
	case 3:
		request := new(TalkedNpcList_In)
		request.Decode(buffer)
		return request
	case 4:
		request := new(NpcTalk_In)
		request.Decode(buffer)
		return request
	case 8:
		request := new(ListOpenedTownTreasures_In)
		request.Decode(buffer)
		return request
	case 9:
		request := new(TakeTownTreasures_In)
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
		request := new(Enter_Out)
		request.Decode(buffer)
		return request
	case 1:
		request := new(Leave_Out)
		request.Decode(buffer)
		return request
	case 2:
		request := new(Move_Out)
		request.Decode(buffer)
		return request
	case 3:
		request := new(TalkedNpcList_Out)
		request.Decode(buffer)
		return request
	case 4:
		request := new(NpcTalk_Out)
		request.Decode(buffer)
		return request
	case 5:
		request := new(NotifyTownPlayers_Out)
		request.Decode(buffer)
		return request
	case 6:
		request := new(UpdateTownPlayer_Out)
		request.Decode(buffer)
		return request
	case 7:
		request := new(UpdateTownPlayerMeditationState_Out)
		request.Decode(buffer)
		return request
	case 8:
		request := new(ListOpenedTownTreasures_Out)
		request.Decode(buffer)
		return request
	case 9:
		request := new(TakeTownTreasures_Out)
		request.Decode(buffer)
		return request
	}
	_ = buffer
	panic("unsupported response")
}

type TownPlayer struct {
	PlayerId          int64  `json:"player_id"`
	Nickname          []byte `json:"nickname"`
	RoleId            int8   `json:"role_id"`
	AtX               int16  `json:"at_x"`
	AtY               int16  `json:"at_y"`
	FashionId         int16  `json:"fashion_id"`
	InMeditationState bool   `json:"in_meditation_state"`
	Level             int16  `json:"level"`
}

func (this *TownPlayer) Decode(buffer *net.Buffer) {
	this.PlayerId = int64(buffer.ReadUint64LE())
	this.Nickname = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.RoleId = int8(buffer.ReadUint8())
	this.AtX = int16(buffer.ReadUint16LE())
	this.AtY = int16(buffer.ReadUint16LE())
	this.FashionId = int16(buffer.ReadUint16LE())
	this.InMeditationState = buffer.ReadUint8() == 1
	this.Level = int16(buffer.ReadUint16LE())
}

func (this *TownPlayer) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.PlayerId))
	buffer.WriteUint16LE(uint16(len(this.Nickname)))
	buffer.WriteBytes(this.Nickname)
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint16LE(uint16(this.AtX))
	buffer.WriteUint16LE(uint16(this.AtY))
	buffer.WriteUint16LE(uint16(this.FashionId))
	if this.InMeditationState {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
	buffer.WriteUint16LE(uint16(this.Level))
}

func (this *TownPlayer) ByteSize() int {
	size := 20
	size += len(this.Nickname)
	return size
}

type Enter_In struct {
	TownId int16 `json:"town_id"`
}

func (this *Enter_In) Process(session *net.Session) {
	g_InHandler.Enter(session, this)
}

func (this *Enter_In) TypeName() string {
	return "town.enter.in"
}

func (this *Enter_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 1, 0
}

type Enter_Out struct {
	Player TownPlayer `json:"player"`
}

func (this *Enter_Out) Process(session *net.Session) {
	g_OutHandler.Enter(session, this)
}

func (this *Enter_Out) TypeName() string {
	return "town.enter.out"
}

func (this *Enter_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 1, 0
}

func (this *Enter_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type Leave_In struct {
}

func (this *Leave_In) Process(session *net.Session) {
	g_InHandler.Leave(session, this)
}

func (this *Leave_In) TypeName() string {
	return "town.leave.in"
}

func (this *Leave_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 1, 1
}

type Leave_Out struct {
	PlayerId int64 `json:"player_id"`
}

func (this *Leave_Out) Process(session *net.Session) {
	g_OutHandler.Leave(session, this)
}

func (this *Leave_Out) TypeName() string {
	return "town.leave.out"
}

func (this *Leave_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 1, 1
}

func (this *Leave_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type Move_In struct {
	ToX int16 `json:"to_x"`
	ToY int16 `json:"to_y"`
}

func (this *Move_In) Process(session *net.Session) {
	g_InHandler.Move(session, this)
}

func (this *Move_In) TypeName() string {
	return "town.move.in"
}

func (this *Move_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 1, 2
}

type Move_Out struct {
	PlayerId int64 `json:"player_id"`
	ToX      int16 `json:"to_x"`
	ToY      int16 `json:"to_y"`
}

func (this *Move_Out) Process(session *net.Session) {
	g_OutHandler.Move(session, this)
}

func (this *Move_Out) TypeName() string {
	return "town.move.out"
}

func (this *Move_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 1, 2
}

func (this *Move_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type TalkedNpcList_In struct {
	TownId int16 `json:"town_id"`
}

func (this *TalkedNpcList_In) Process(session *net.Session) {
	g_InHandler.TalkedNpcList(session, this)
}

func (this *TalkedNpcList_In) TypeName() string {
	return "town.talked_npc_list.in"
}

func (this *TalkedNpcList_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 1, 3
}

type TalkedNpcList_Out struct {
	NpcList []TalkedNpcList_Out_NpcList `json:"npc_list"`
}

type TalkedNpcList_Out_NpcList struct {
	NpcId   int32 `json:"npc_id"`
	QuestId int16 `json:"quest_id"`
}

func (this *TalkedNpcList_Out) Process(session *net.Session) {
	g_OutHandler.TalkedNpcList(session, this)
}

func (this *TalkedNpcList_Out) TypeName() string {
	return "town.talked_npc_list.out"
}

func (this *TalkedNpcList_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 1, 3
}

func (this *TalkedNpcList_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type NpcTalk_In struct {
	NpcId int32 `json:"npc_id"`
}

func (this *NpcTalk_In) Process(session *net.Session) {
	g_InHandler.NpcTalk(session, this)
}

func (this *NpcTalk_In) TypeName() string {
	return "town.npc_talk.in"
}

func (this *NpcTalk_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 1, 4
}

type NpcTalk_Out struct {
}

func (this *NpcTalk_Out) Process(session *net.Session) {
	g_OutHandler.NpcTalk(session, this)
}

func (this *NpcTalk_Out) TypeName() string {
	return "town.npc_talk.out"
}

func (this *NpcTalk_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 1, 4
}

func (this *NpcTalk_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type NotifyTownPlayers_Out struct {
	Players []NotifyTownPlayers_Out_Players `json:"players"`
}

type NotifyTownPlayers_Out_Players struct {
	Player TownPlayer `json:"player"`
}

func (this *NotifyTownPlayers_Out) Process(session *net.Session) {
	g_OutHandler.NotifyTownPlayers(session, this)
}

func (this *NotifyTownPlayers_Out) TypeName() string {
	return "town.notify_town_players.out"
}

func (this *NotifyTownPlayers_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 1, 5
}

func (this *NotifyTownPlayers_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type UpdateTownPlayer_Out struct {
	PlayerId  int64 `json:"player_id"`
	FashionId int16 `json:"fashion_id"`
}

func (this *UpdateTownPlayer_Out) Process(session *net.Session) {
	g_OutHandler.UpdateTownPlayer(session, this)
}

func (this *UpdateTownPlayer_Out) TypeName() string {
	return "town.update_town_player.out"
}

func (this *UpdateTownPlayer_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 1, 6
}

func (this *UpdateTownPlayer_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type UpdateTownPlayerMeditationState_Out struct {
	PlayerId        int64 `json:"player_id"`
	MeditationState bool  `json:"meditation_state"`
}

func (this *UpdateTownPlayerMeditationState_Out) Process(session *net.Session) {
	g_OutHandler.UpdateTownPlayerMeditationState(session, this)
}

func (this *UpdateTownPlayerMeditationState_Out) TypeName() string {
	return "town.update_town_player_meditation_state.out"
}

func (this *UpdateTownPlayerMeditationState_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 1, 7
}

func (this *UpdateTownPlayerMeditationState_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ListOpenedTownTreasures_In struct {
}

func (this *ListOpenedTownTreasures_In) Process(session *net.Session) {
	g_InHandler.ListOpenedTownTreasures(session, this)
}

func (this *ListOpenedTownTreasures_In) TypeName() string {
	return "town.list_opened_town_treasures.in"
}

func (this *ListOpenedTownTreasures_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 1, 8
}

type ListOpenedTownTreasures_Out struct {
	Treasures []ListOpenedTownTreasures_Out_Treasures `json:"treasures"`
}

type ListOpenedTownTreasures_Out_Treasures struct {
	TownId int16 `json:"town_id"`
}

func (this *ListOpenedTownTreasures_Out) Process(session *net.Session) {
	g_OutHandler.ListOpenedTownTreasures(session, this)
}

func (this *ListOpenedTownTreasures_Out) TypeName() string {
	return "town.list_opened_town_treasures.out"
}

func (this *ListOpenedTownTreasures_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 1, 8
}

func (this *ListOpenedTownTreasures_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type TakeTownTreasures_In struct {
	TownId int16 `json:"town_id"`
}

func (this *TakeTownTreasures_In) Process(session *net.Session) {
	g_InHandler.TakeTownTreasures(session, this)
}

func (this *TakeTownTreasures_In) TypeName() string {
	return "town.take_town_treasures.in"
}

func (this *TakeTownTreasures_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 1, 9
}

type TakeTownTreasures_Out struct {
}

func (this *TakeTownTreasures_Out) Process(session *net.Session) {
	g_OutHandler.TakeTownTreasures(session, this)
}

func (this *TakeTownTreasures_Out) TypeName() string {
	return "town.take_town_treasures.out"
}

func (this *TakeTownTreasures_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 1, 9
}

func (this *TakeTownTreasures_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

func (this *Enter_In) Decode(buffer *net.Buffer) {
	this.TownId = int16(buffer.ReadUint16LE())
}

func (this *Enter_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(1)
	buffer.WriteUint8(0)
	buffer.WriteUint16LE(uint16(this.TownId))
}

func (this *Enter_In) ByteSize() int {
	size := 4
	return size
}

func (this *Enter_Out) Decode(buffer *net.Buffer) {
	this.Player.Decode(buffer)
}

func (this *Enter_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(1)
	buffer.WriteUint8(0)
	this.Player.Encode(buffer)
}

func (this *Enter_Out) ByteSize() int {
	size := 2
	size += this.Player.ByteSize()
	return size
}

func (this *Leave_In) Decode(buffer *net.Buffer) {
}

func (this *Leave_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(1)
	buffer.WriteUint8(1)
}

func (this *Leave_In) ByteSize() int {
	size := 2
	return size
}

func (this *Leave_Out) Decode(buffer *net.Buffer) {
	this.PlayerId = int64(buffer.ReadUint64LE())
}

func (this *Leave_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(1)
	buffer.WriteUint8(1)
	buffer.WriteUint64LE(uint64(this.PlayerId))
}

func (this *Leave_Out) ByteSize() int {
	size := 10
	return size
}

func (this *Move_In) Decode(buffer *net.Buffer) {
	this.ToX = int16(buffer.ReadUint16LE())
	this.ToY = int16(buffer.ReadUint16LE())
}

func (this *Move_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(1)
	buffer.WriteUint8(2)
	buffer.WriteUint16LE(uint16(this.ToX))
	buffer.WriteUint16LE(uint16(this.ToY))
}

func (this *Move_In) ByteSize() int {
	size := 6
	return size
}

func (this *Move_Out) Decode(buffer *net.Buffer) {
	this.PlayerId = int64(buffer.ReadUint64LE())
	this.ToX = int16(buffer.ReadUint16LE())
	this.ToY = int16(buffer.ReadUint16LE())
}

func (this *Move_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(1)
	buffer.WriteUint8(2)
	buffer.WriteUint64LE(uint64(this.PlayerId))
	buffer.WriteUint16LE(uint16(this.ToX))
	buffer.WriteUint16LE(uint16(this.ToY))
}

func (this *Move_Out) ByteSize() int {
	size := 14
	return size
}

func (this *TalkedNpcList_In) Decode(buffer *net.Buffer) {
	this.TownId = int16(buffer.ReadUint16LE())
}

func (this *TalkedNpcList_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(1)
	buffer.WriteUint8(3)
	buffer.WriteUint16LE(uint16(this.TownId))
}

func (this *TalkedNpcList_In) ByteSize() int {
	size := 4
	return size
}

func (this *TalkedNpcList_Out) Decode(buffer *net.Buffer) {
	this.NpcList = make([]TalkedNpcList_Out_NpcList, buffer.ReadUint8())
	for i := 0; i < len(this.NpcList); i++ {
		this.NpcList[i].Decode(buffer)
	}
}

func (this *TalkedNpcList_Out_NpcList) Decode(buffer *net.Buffer) {
	this.NpcId = int32(buffer.ReadUint32LE())
	this.QuestId = int16(buffer.ReadUint16LE())
}

func (this *TalkedNpcList_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(1)
	buffer.WriteUint8(3)
	buffer.WriteUint8(uint8(len(this.NpcList)))
	for i := 0; i < len(this.NpcList); i++ {
		this.NpcList[i].Encode(buffer)
	}
}

func (this *TalkedNpcList_Out_NpcList) Encode(buffer *net.Buffer) {
	buffer.WriteUint32LE(uint32(this.NpcId))
	buffer.WriteUint16LE(uint16(this.QuestId))
}

func (this *TalkedNpcList_Out) ByteSize() int {
	size := 3
	size += len(this.NpcList) * 6
	return size
}

func (this *NpcTalk_In) Decode(buffer *net.Buffer) {
	this.NpcId = int32(buffer.ReadUint32LE())
}

func (this *NpcTalk_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(1)
	buffer.WriteUint8(4)
	buffer.WriteUint32LE(uint32(this.NpcId))
}

func (this *NpcTalk_In) ByteSize() int {
	size := 6
	return size
}

func (this *NpcTalk_Out) Decode(buffer *net.Buffer) {
}

func (this *NpcTalk_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(1)
	buffer.WriteUint8(4)
}

func (this *NpcTalk_Out) ByteSize() int {
	size := 2
	return size
}

func (this *NotifyTownPlayers_Out) Decode(buffer *net.Buffer) {
	this.Players = make([]NotifyTownPlayers_Out_Players, buffer.ReadUint8())
	for i := 0; i < len(this.Players); i++ {
		this.Players[i].Decode(buffer)
	}
}

func (this *NotifyTownPlayers_Out_Players) Decode(buffer *net.Buffer) {
	this.Player.Decode(buffer)
}

func (this *NotifyTownPlayers_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(1)
	buffer.WriteUint8(5)
	buffer.WriteUint8(uint8(len(this.Players)))
	for i := 0; i < len(this.Players); i++ {
		this.Players[i].Encode(buffer)
	}
}

func (this *NotifyTownPlayers_Out_Players) Encode(buffer *net.Buffer) {
	this.Player.Encode(buffer)
}

func (this *NotifyTownPlayers_Out) ByteSize() int {
	size := 3
	for i := 0; i < len(this.Players); i++ {
		size += this.Players[i].ByteSize()
	}
	return size
}

func (this *NotifyTownPlayers_Out_Players) ByteSize() int {
	size := 0
	size += this.Player.ByteSize()
	return size
}

func (this *UpdateTownPlayer_Out) Decode(buffer *net.Buffer) {
	this.PlayerId = int64(buffer.ReadUint64LE())
	this.FashionId = int16(buffer.ReadUint16LE())
}

func (this *UpdateTownPlayer_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(1)
	buffer.WriteUint8(6)
	buffer.WriteUint64LE(uint64(this.PlayerId))
	buffer.WriteUint16LE(uint16(this.FashionId))
}

func (this *UpdateTownPlayer_Out) ByteSize() int {
	size := 12
	return size
}

func (this *UpdateTownPlayerMeditationState_Out) Decode(buffer *net.Buffer) {
	this.PlayerId = int64(buffer.ReadUint64LE())
	this.MeditationState = buffer.ReadUint8() == 1
}

func (this *UpdateTownPlayerMeditationState_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(1)
	buffer.WriteUint8(7)
	buffer.WriteUint64LE(uint64(this.PlayerId))
	if this.MeditationState {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *UpdateTownPlayerMeditationState_Out) ByteSize() int {
	size := 11
	return size
}

func (this *ListOpenedTownTreasures_In) Decode(buffer *net.Buffer) {
}

func (this *ListOpenedTownTreasures_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(1)
	buffer.WriteUint8(8)
}

func (this *ListOpenedTownTreasures_In) ByteSize() int {
	size := 2
	return size
}

func (this *ListOpenedTownTreasures_Out) Decode(buffer *net.Buffer) {
	this.Treasures = make([]ListOpenedTownTreasures_Out_Treasures, buffer.ReadUint16LE())
	for i := 0; i < len(this.Treasures); i++ {
		this.Treasures[i].Decode(buffer)
	}
}

func (this *ListOpenedTownTreasures_Out_Treasures) Decode(buffer *net.Buffer) {
	this.TownId = int16(buffer.ReadUint16LE())
}

func (this *ListOpenedTownTreasures_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(1)
	buffer.WriteUint8(8)
	buffer.WriteUint16LE(uint16(len(this.Treasures)))
	for i := 0; i < len(this.Treasures); i++ {
		this.Treasures[i].Encode(buffer)
	}
}

func (this *ListOpenedTownTreasures_Out_Treasures) Encode(buffer *net.Buffer) {
	buffer.WriteUint16LE(uint16(this.TownId))
}

func (this *ListOpenedTownTreasures_Out) ByteSize() int {
	size := 4
	size += len(this.Treasures) * 2
	return size
}

func (this *TakeTownTreasures_In) Decode(buffer *net.Buffer) {
	this.TownId = int16(buffer.ReadUint16LE())
}

func (this *TakeTownTreasures_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(1)
	buffer.WriteUint8(9)
	buffer.WriteUint16LE(uint16(this.TownId))
}

func (this *TakeTownTreasures_In) ByteSize() int {
	size := 4
	return size
}

func (this *TakeTownTreasures_Out) Decode(buffer *net.Buffer) {
}

func (this *TakeTownTreasures_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(1)
	buffer.WriteUint8(9)
}

func (this *TakeTownTreasures_Out) ByteSize() int {
	size := 2
	return size
}
