package multi_level_api

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
	CreateRoom(*net.Session, *CreateRoom_In)
	EnterRoom(*net.Session, *EnterRoom_In)
	LeaveRoom(*net.Session, *LeaveRoom_In)
	ChangeBuddy(*net.Session, *ChangeBuddy_In)
	GetBaseInfo(*net.Session, *GetBaseInfo_In)
	GetOnlineFriend(*net.Session, *GetOnlineFriend_In)
	InviteIntoRoom(*net.Session, *InviteIntoRoom_In)
	RefuseRoomInvite(*net.Session, *RefuseRoomInvite_In)
	MatchPlayer(*net.Session, *MatchPlayer_In)
	MatchRoom(*net.Session, *MatchRoom_In)
	CancelMatchRoom(*net.Session, *CancelMatchRoom_In)
	GetFormInfo(*net.Session, *GetFormInfo_In)
	SetForm(*net.Session, *SetForm_In)
	GetBattleRoleInfo(*net.Session, *GetBattleRoleInfo_In)
	GetMatchInfo(*net.Session, *GetMatchInfo_In)
	CancelMatchPlayer(*net.Session, *CancelMatchPlayer_In)
}

type OutHandler interface {
	CreateRoom(*net.Session, *CreateRoom_Out)
	EnterRoom(*net.Session, *EnterRoom_Out)
	NotifyRoomInfo(*net.Session, *NotifyRoomInfo_Out)
	LeaveRoom(*net.Session, *LeaveRoom_Out)
	NotifyJoinPartner(*net.Session, *NotifyJoinPartner_Out)
	ChangeBuddy(*net.Session, *ChangeBuddy_Out)
	GetBaseInfo(*net.Session, *GetBaseInfo_Out)
	GetOnlineFriend(*net.Session, *GetOnlineFriend_Out)
	InviteIntoRoom(*net.Session, *InviteIntoRoom_Out)
	NotifyRoomInvite(*net.Session, *NotifyRoomInvite_Out)
	NotifyPlayersInfo(*net.Session, *NotifyPlayersInfo_Out)
	RefuseRoomInvite(*net.Session, *RefuseRoomInvite_Out)
	NotifyRoomInviteRefuse(*net.Session, *NotifyRoomInviteRefuse_Out)
	NotifyUpdatePartner(*net.Session, *NotifyUpdatePartner_Out)
	MatchPlayer(*net.Session, *MatchPlayer_Out)
	NotifyMatchPlayerSuccess(*net.Session, *NotifyMatchPlayerSuccess_Out)
	MatchRoom(*net.Session, *MatchRoom_Out)
	CancelMatchRoom(*net.Session, *CancelMatchRoom_Out)
	GetFormInfo(*net.Session, *GetFormInfo_Out)
	SetForm(*net.Session, *SetForm_Out)
	GetBattleRoleInfo(*net.Session, *GetBattleRoleInfo_Out)
	NotifyMatchPlayerInfo(*net.Session, *NotifyMatchPlayerInfo_Out)
	GetMatchInfo(*net.Session, *GetMatchInfo_Out)
	NotifyFormInfo(*net.Session, *NotifyFormInfo_Out)
	CancelMatchPlayer(*net.Session, *CancelMatchPlayer_Out)
}

func DecodeIn(message []byte) Request {
	var actionId = message[0]
	var buffer = net.NewBuffer(message[1:])
	switch actionId {
	case 1:
		request := new(CreateRoom_In)
		request.Decode(buffer)
		return request
	case 2:
		request := new(EnterRoom_In)
		request.Decode(buffer)
		return request
	case 5:
		request := new(LeaveRoom_In)
		request.Decode(buffer)
		return request
	case 7:
		request := new(ChangeBuddy_In)
		request.Decode(buffer)
		return request
	case 10:
		request := new(GetBaseInfo_In)
		request.Decode(buffer)
		return request
	case 12:
		request := new(GetOnlineFriend_In)
		request.Decode(buffer)
		return request
	case 13:
		request := new(InviteIntoRoom_In)
		request.Decode(buffer)
		return request
	case 16:
		request := new(RefuseRoomInvite_In)
		request.Decode(buffer)
		return request
	case 19:
		request := new(MatchPlayer_In)
		request.Decode(buffer)
		return request
	case 22:
		request := new(MatchRoom_In)
		request.Decode(buffer)
		return request
	case 23:
		request := new(CancelMatchRoom_In)
		request.Decode(buffer)
		return request
	case 24:
		request := new(GetFormInfo_In)
		request.Decode(buffer)
		return request
	case 25:
		request := new(SetForm_In)
		request.Decode(buffer)
		return request
	case 26:
		request := new(GetBattleRoleInfo_In)
		request.Decode(buffer)
		return request
	case 28:
		request := new(GetMatchInfo_In)
		request.Decode(buffer)
		return request
	case 30:
		request := new(CancelMatchPlayer_In)
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
		request := new(CreateRoom_Out)
		request.Decode(buffer)
		return request
	case 2:
		request := new(EnterRoom_Out)
		request.Decode(buffer)
		return request
	case 4:
		request := new(NotifyRoomInfo_Out)
		request.Decode(buffer)
		return request
	case 5:
		request := new(LeaveRoom_Out)
		request.Decode(buffer)
		return request
	case 6:
		request := new(NotifyJoinPartner_Out)
		request.Decode(buffer)
		return request
	case 7:
		request := new(ChangeBuddy_Out)
		request.Decode(buffer)
		return request
	case 10:
		request := new(GetBaseInfo_Out)
		request.Decode(buffer)
		return request
	case 12:
		request := new(GetOnlineFriend_Out)
		request.Decode(buffer)
		return request
	case 13:
		request := new(InviteIntoRoom_Out)
		request.Decode(buffer)
		return request
	case 14:
		request := new(NotifyRoomInvite_Out)
		request.Decode(buffer)
		return request
	case 15:
		request := new(NotifyPlayersInfo_Out)
		request.Decode(buffer)
		return request
	case 16:
		request := new(RefuseRoomInvite_Out)
		request.Decode(buffer)
		return request
	case 17:
		request := new(NotifyRoomInviteRefuse_Out)
		request.Decode(buffer)
		return request
	case 18:
		request := new(NotifyUpdatePartner_Out)
		request.Decode(buffer)
		return request
	case 19:
		request := new(MatchPlayer_Out)
		request.Decode(buffer)
		return request
	case 21:
		request := new(NotifyMatchPlayerSuccess_Out)
		request.Decode(buffer)
		return request
	case 22:
		request := new(MatchRoom_Out)
		request.Decode(buffer)
		return request
	case 23:
		request := new(CancelMatchRoom_Out)
		request.Decode(buffer)
		return request
	case 24:
		request := new(GetFormInfo_Out)
		request.Decode(buffer)
		return request
	case 25:
		request := new(SetForm_Out)
		request.Decode(buffer)
		return request
	case 26:
		request := new(GetBattleRoleInfo_Out)
		request.Decode(buffer)
		return request
	case 27:
		request := new(NotifyMatchPlayerInfo_Out)
		request.Decode(buffer)
		return request
	case 28:
		request := new(GetMatchInfo_Out)
		request.Decode(buffer)
		return request
	case 29:
		request := new(NotifyFormInfo_Out)
		request.Decode(buffer)
		return request
	case 30:
		request := new(CancelMatchPlayer_Out)
		request.Decode(buffer)
		return request
	}
	_ = buffer
	panic("unsupported response")
}

type RoomStatus int8

const (
	ROOM_STATUS_SUCCESS              RoomStatus = 0
	ROOM_STATUS_FAILED_FULL          RoomStatus = 1
	ROOM_STATUS_FAILED_NOT_EXIST     RoomStatus = 2
	ROOM_STATUS_FAILED_FIGHTING      RoomStatus = 3
	ROOM_STATUS_FAILED_REQUIRE_LEVEL RoomStatus = 4
	ROOM_STATUS_FAILED_REQUIRE_LOCK  RoomStatus = 5
)

type PartnerInfo struct {
	Pid                  int64  `json:"pid"`
	Pos                  int8   `json:"pos"`
	Nick                 []byte `json:"nick"`
	RoleId               int8   `json:"role_id"`
	FashionId            int16  `json:"fashion_id"`
	Level                int16  `json:"level"`
	BuddyLevel           int16  `json:"buddy_level"`
	BuddyRoleId          int8   `json:"buddy_role_id"`
	FriendshipLevel      int32  `json:"friendship_level"`
	BuddyFriendshipLevel int32  `json:"buddy_friendship_level"`
}

func (this *PartnerInfo) Decode(buffer *net.Buffer) {
	this.Pid = int64(buffer.ReadUint64LE())
	this.Pos = int8(buffer.ReadUint8())
	this.Nick = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.RoleId = int8(buffer.ReadUint8())
	this.FashionId = int16(buffer.ReadUint16LE())
	this.Level = int16(buffer.ReadUint16LE())
	this.BuddyLevel = int16(buffer.ReadUint16LE())
	this.BuddyRoleId = int8(buffer.ReadUint8())
	this.FriendshipLevel = int32(buffer.ReadUint32LE())
	this.BuddyFriendshipLevel = int32(buffer.ReadUint32LE())
}

func (this *PartnerInfo) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.Pid))
	buffer.WriteUint8(uint8(this.Pos))
	buffer.WriteUint16LE(uint16(len(this.Nick)))
	buffer.WriteBytes(this.Nick)
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint16LE(uint16(this.FashionId))
	buffer.WriteUint16LE(uint16(this.Level))
	buffer.WriteUint16LE(uint16(this.BuddyLevel))
	buffer.WriteUint8(uint8(this.BuddyRoleId))
	buffer.WriteUint32LE(uint32(this.FriendshipLevel))
	buffer.WriteUint32LE(uint32(this.BuddyFriendshipLevel))
}

func (this *PartnerInfo) ByteSize() int {
	size := 27
	size += len(this.Nick)
	return size
}

type CreateRoom_In struct {
	LevelId int32 `json:"level_id"`
}

func (this *CreateRoom_In) Process(session *net.Session) {
	g_InHandler.CreateRoom(session, this)
}

func (this *CreateRoom_In) TypeName() string {
	return "multi_level.create_room.in"
}

func (this *CreateRoom_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 16, 1
}

type CreateRoom_Out struct {
	Status RoomStatus `json:"status"`
}

func (this *CreateRoom_Out) Process(session *net.Session) {
	g_OutHandler.CreateRoom(session, this)
}

func (this *CreateRoom_Out) TypeName() string {
	return "multi_level.create_room.out"
}

func (this *CreateRoom_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 16, 1
}

func (this *CreateRoom_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type EnterRoom_In struct {
	RoomId    int64  `json:"room_id"`
	RoomToken uint64 `json:"room_token"`
}

func (this *EnterRoom_In) Process(session *net.Session) {
	g_InHandler.EnterRoom(session, this)
}

func (this *EnterRoom_In) TypeName() string {
	return "multi_level.enter_room.in"
}

func (this *EnterRoom_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 16, 2
}

type EnterRoom_Out struct {
	Status RoomStatus `json:"status"`
}

func (this *EnterRoom_Out) Process(session *net.Session) {
	g_OutHandler.EnterRoom(session, this)
}

func (this *EnterRoom_Out) TypeName() string {
	return "multi_level.enter_room.out"
}

func (this *EnterRoom_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 16, 2
}

func (this *EnterRoom_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type NotifyRoomInfo_Out struct {
	RoomId    int64                         `json:"room_id"`
	LeaderPid int64                         `json:"leader_pid"`
	LevelId   int32                         `json:"level_id"`
	Partners  []NotifyRoomInfo_Out_Partners `json:"partners"`
}

type NotifyRoomInfo_Out_Partners struct {
	Partner PartnerInfo `json:"partner"`
}

func (this *NotifyRoomInfo_Out) Process(session *net.Session) {
	g_OutHandler.NotifyRoomInfo(session, this)
}

func (this *NotifyRoomInfo_Out) TypeName() string {
	return "multi_level.notify_room_info.out"
}

func (this *NotifyRoomInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 16, 4
}

func (this *NotifyRoomInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type LeaveRoom_In struct {
}

func (this *LeaveRoom_In) Process(session *net.Session) {
	g_InHandler.LeaveRoom(session, this)
}

func (this *LeaveRoom_In) TypeName() string {
	return "multi_level.leave_room.in"
}

func (this *LeaveRoom_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 16, 5
}

type LeaveRoom_Out struct {
	Pid       int64                  `json:"pid"`
	LeaderPid int64                  `json:"leader_pid"`
	Tokens    []LeaveRoom_Out_Tokens `json:"tokens"`
}

type LeaveRoom_Out_Tokens struct {
	Pos   int8   `json:"pos"`
	Token uint64 `json:"token"`
}

func (this *LeaveRoom_Out) Process(session *net.Session) {
	g_OutHandler.LeaveRoom(session, this)
}

func (this *LeaveRoom_Out) TypeName() string {
	return "multi_level.leave_room.out"
}

func (this *LeaveRoom_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 16, 5
}

func (this *LeaveRoom_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type NotifyJoinPartner_Out struct {
	Partner PartnerInfo `json:"partner"`
}

func (this *NotifyJoinPartner_Out) Process(session *net.Session) {
	g_OutHandler.NotifyJoinPartner(session, this)
}

func (this *NotifyJoinPartner_Out) TypeName() string {
	return "multi_level.notify_join_partner.out"
}

func (this *NotifyJoinPartner_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 16, 6
}

func (this *NotifyJoinPartner_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ChangeBuddy_In struct {
	BuddyRoleId int8 `json:"buddy_role_id"`
}

func (this *ChangeBuddy_In) Process(session *net.Session) {
	g_InHandler.ChangeBuddy(session, this)
}

func (this *ChangeBuddy_In) TypeName() string {
	return "multi_level.change_buddy.in"
}

func (this *ChangeBuddy_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 16, 7
}

type ChangeBuddy_Out struct {
	Success bool `json:"success"`
}

func (this *ChangeBuddy_Out) Process(session *net.Session) {
	g_OutHandler.ChangeBuddy(session, this)
}

func (this *ChangeBuddy_Out) TypeName() string {
	return "multi_level.change_buddy.out"
}

func (this *ChangeBuddy_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 16, 7
}

func (this *ChangeBuddy_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetBaseInfo_In struct {
}

func (this *GetBaseInfo_In) Process(session *net.Session) {
	g_InHandler.GetBaseInfo(session, this)
}

func (this *GetBaseInfo_In) TypeName() string {
	return "multi_level.get_base_info.in"
}

func (this *GetBaseInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 16, 10
}

type GetBaseInfo_Out struct {
	DailyNum int8  `json:"daily_num"`
	Lock     int32 `json:"lock"`
}

func (this *GetBaseInfo_Out) Process(session *net.Session) {
	g_OutHandler.GetBaseInfo(session, this)
}

func (this *GetBaseInfo_Out) TypeName() string {
	return "multi_level.get_base_info.out"
}

func (this *GetBaseInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 16, 10
}

func (this *GetBaseInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetOnlineFriend_In struct {
}

func (this *GetOnlineFriend_In) Process(session *net.Session) {
	g_InHandler.GetOnlineFriend(session, this)
}

func (this *GetOnlineFriend_In) TypeName() string {
	return "multi_level.get_online_friend.in"
}

func (this *GetOnlineFriend_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 16, 12
}

type GetOnlineFriend_Out struct {
	Friends []GetOnlineFriend_Out_Friends `json:"friends"`
}

type GetOnlineFriend_Out_Friends struct {
	Pid      int64  `json:"pid"`
	RoleId   int8   `json:"role_id"`
	Nickname []byte `json:"nickname"`
	Level    int16  `json:"level"`
	FightNum int32  `json:"fight_num"`
	DailyNum int8   `json:"daily_num"`
	Lock     int32  `json:"lock"`
}

func (this *GetOnlineFriend_Out) Process(session *net.Session) {
	g_OutHandler.GetOnlineFriend(session, this)
}

func (this *GetOnlineFriend_Out) TypeName() string {
	return "multi_level.get_online_friend.out"
}

func (this *GetOnlineFriend_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 16, 12
}

func (this *GetOnlineFriend_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type InviteIntoRoom_In struct {
	Pid int64 `json:"pid"`
}

func (this *InviteIntoRoom_In) Process(session *net.Session) {
	g_InHandler.InviteIntoRoom(session, this)
}

func (this *InviteIntoRoom_In) TypeName() string {
	return "multi_level.invite_into_room.in"
}

func (this *InviteIntoRoom_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 16, 13
}

type InviteIntoRoom_Out struct {
	IsOffline bool `json:"isOffline"`
}

func (this *InviteIntoRoom_Out) Process(session *net.Session) {
	g_OutHandler.InviteIntoRoom(session, this)
}

func (this *InviteIntoRoom_Out) TypeName() string {
	return "multi_level.invite_into_room.out"
}

func (this *InviteIntoRoom_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 16, 13
}

func (this *InviteIntoRoom_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type NotifyRoomInvite_Out struct {
	RoomId       int64  `json:"room_id"`
	LevelId      int32  `json:"level_id"`
	Nickname     []byte `json:"nickname"`
	InviterId    int64  `json:"inviter_id"`
	DailyNum     int8   `json:"daily_num"`
	GameServerId int64  `json:"game_server_id"`
}

func (this *NotifyRoomInvite_Out) Process(session *net.Session) {
	g_OutHandler.NotifyRoomInvite(session, this)
}

func (this *NotifyRoomInvite_Out) TypeName() string {
	return "multi_level.notify_room_invite.out"
}

func (this *NotifyRoomInvite_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 16, 14
}

func (this *NotifyRoomInvite_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type NotifyPlayersInfo_Out struct {
	Players []NotifyPlayersInfo_Out_Players `json:"players"`
}

type NotifyPlayersInfo_Out_Players struct {
	PlayerId int64  `json:"player_id"`
	Nickname []byte `json:"nickname"`
}

func (this *NotifyPlayersInfo_Out) Process(session *net.Session) {
	g_OutHandler.NotifyPlayersInfo(session, this)
}

func (this *NotifyPlayersInfo_Out) TypeName() string {
	return "multi_level.notify_players_info.out"
}

func (this *NotifyPlayersInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 16, 15
}

func (this *NotifyPlayersInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type RefuseRoomInvite_In struct {
	RoomId    int64 `json:"room_id"`
	InviterId int64 `json:"inviter_id"`
}

func (this *RefuseRoomInvite_In) Process(session *net.Session) {
	g_InHandler.RefuseRoomInvite(session, this)
}

func (this *RefuseRoomInvite_In) TypeName() string {
	return "multi_level.refuse_room_invite.in"
}

func (this *RefuseRoomInvite_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 16, 16
}

type RefuseRoomInvite_Out struct {
}

func (this *RefuseRoomInvite_Out) Process(session *net.Session) {
	g_OutHandler.RefuseRoomInvite(session, this)
}

func (this *RefuseRoomInvite_Out) TypeName() string {
	return "multi_level.refuse_room_invite.out"
}

func (this *RefuseRoomInvite_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 16, 16
}

func (this *RefuseRoomInvite_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type NotifyRoomInviteRefuse_Out struct {
	Nickname []byte `json:"nickname"`
}

func (this *NotifyRoomInviteRefuse_Out) Process(session *net.Session) {
	g_OutHandler.NotifyRoomInviteRefuse(session, this)
}

func (this *NotifyRoomInviteRefuse_Out) TypeName() string {
	return "multi_level.notify_room_invite_refuse.out"
}

func (this *NotifyRoomInviteRefuse_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 16, 17
}

func (this *NotifyRoomInviteRefuse_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type NotifyUpdatePartner_Out struct {
	Partner PartnerInfo `json:"partner"`
}

func (this *NotifyUpdatePartner_Out) Process(session *net.Session) {
	g_OutHandler.NotifyUpdatePartner(session, this)
}

func (this *NotifyUpdatePartner_Out) TypeName() string {
	return "multi_level.notify_update_partner.out"
}

func (this *NotifyUpdatePartner_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 16, 18
}

func (this *NotifyUpdatePartner_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type MatchPlayer_In struct {
	Level int32 `json:"level"`
}

func (this *MatchPlayer_In) Process(session *net.Session) {
	g_InHandler.MatchPlayer(session, this)
}

func (this *MatchPlayer_In) TypeName() string {
	return "multi_level.match_player.in"
}

func (this *MatchPlayer_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 16, 19
}

type MatchPlayer_Out struct {
	Token uint64 `json:"token"`
	Ok    bool   `json:"ok"`
}

func (this *MatchPlayer_Out) Process(session *net.Session) {
	g_OutHandler.MatchPlayer(session, this)
}

func (this *MatchPlayer_Out) TypeName() string {
	return "multi_level.match_player.out"
}

func (this *MatchPlayer_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 16, 19
}

func (this *MatchPlayer_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type NotifyMatchPlayerSuccess_Out struct {
	ServerId         int32  `json:"server_id"`
	RoomId           int64  `json:"room_id"`
	PlayerMatchToken uint64 `json:"player_match_token"`
	RoomMatchToken   uint64 `json:"room_match_token"`
}

func (this *NotifyMatchPlayerSuccess_Out) Process(session *net.Session) {
	g_OutHandler.NotifyMatchPlayerSuccess(session, this)
}

func (this *NotifyMatchPlayerSuccess_Out) TypeName() string {
	return "multi_level.notify_match_player_success.out"
}

func (this *NotifyMatchPlayerSuccess_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 16, 21
}

func (this *NotifyMatchPlayerSuccess_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type MatchRoom_In struct {
	Pos int8 `json:"pos"`
}

func (this *MatchRoom_In) Process(session *net.Session) {
	g_InHandler.MatchRoom(session, this)
}

func (this *MatchRoom_In) TypeName() string {
	return "multi_level.match_room.in"
}

func (this *MatchRoom_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 16, 22
}

type MatchRoom_Out struct {
	MatchToken uint64 `json:"match_token"`
	Ok         bool   `json:"ok"`
	Pos        int8   `json:"pos"`
}

func (this *MatchRoom_Out) Process(session *net.Session) {
	g_OutHandler.MatchRoom(session, this)
}

func (this *MatchRoom_Out) TypeName() string {
	return "multi_level.match_room.out"
}

func (this *MatchRoom_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 16, 22
}

func (this *MatchRoom_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type CancelMatchRoom_In struct {
	MatchToken uint64 `json:"match_token"`
}

func (this *CancelMatchRoom_In) Process(session *net.Session) {
	g_InHandler.CancelMatchRoom(session, this)
}

func (this *CancelMatchRoom_In) TypeName() string {
	return "multi_level.cancel_match_room.in"
}

func (this *CancelMatchRoom_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 16, 23
}

type CancelMatchRoom_Out struct {
}

func (this *CancelMatchRoom_Out) Process(session *net.Session) {
	g_OutHandler.CancelMatchRoom(session, this)
}

func (this *CancelMatchRoom_Out) TypeName() string {
	return "multi_level.cancel_match_room.out"
}

func (this *CancelMatchRoom_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 16, 23
}

func (this *CancelMatchRoom_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetFormInfo_In struct {
}

func (this *GetFormInfo_In) Process(session *net.Session) {
	g_InHandler.GetFormInfo(session, this)
}

func (this *GetFormInfo_In) TypeName() string {
	return "multi_level.get_form_info.in"
}

func (this *GetFormInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 16, 24
}

type GetFormInfo_Out struct {
	Formation []GetFormInfo_Out_Formation `json:"formation"`
}

type GetFormInfo_Out_Formation struct {
	Pos    int8  `json:"pos"`
	Pid    int64 `json:"pid"`
	RoleId int8  `json:"role_id"`
}

func (this *GetFormInfo_Out) Process(session *net.Session) {
	g_OutHandler.GetFormInfo(session, this)
}

func (this *GetFormInfo_Out) TypeName() string {
	return "multi_level.get_form_info.out"
}

func (this *GetFormInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 16, 24
}

func (this *GetFormInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type SetForm_In struct {
	Formation []SetForm_In_Formation `json:"formation"`
}

type SetForm_In_Formation struct {
	Pos        int8  `json:"pos"`
	Pid        int64 `json:"pid"`
	IsMainRole bool  `json:"is_main_role"`
}

func (this *SetForm_In) Process(session *net.Session) {
	g_InHandler.SetForm(session, this)
}

func (this *SetForm_In) TypeName() string {
	return "multi_level.set_form.in"
}

func (this *SetForm_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 16, 25
}

type SetForm_Out struct {
}

func (this *SetForm_Out) Process(session *net.Session) {
	g_OutHandler.SetForm(session, this)
}

func (this *SetForm_Out) TypeName() string {
	return "multi_level.set_form.out"
}

func (this *SetForm_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 16, 25
}

func (this *SetForm_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetBattleRoleInfo_In struct {
}

func (this *GetBattleRoleInfo_In) Process(session *net.Session) {
	g_InHandler.GetBattleRoleInfo(session, this)
}

func (this *GetBattleRoleInfo_In) TypeName() string {
	return "multi_level.get_battle_role_info.in"
}

func (this *GetBattleRoleInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 16, 26
}

type GetBattleRoleInfo_Out struct {
	MainRoleId  int8 `json:"main_role_id"`
	BuddyRoleId int8 `json:"buddy_role_id"`
}

func (this *GetBattleRoleInfo_Out) Process(session *net.Session) {
	g_OutHandler.GetBattleRoleInfo(session, this)
}

func (this *GetBattleRoleInfo_Out) TypeName() string {
	return "multi_level.get_battle_role_info.out"
}

func (this *GetBattleRoleInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 16, 26
}

func (this *GetBattleRoleInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type NotifyMatchPlayerInfo_Out struct {
	RoomMatchToken uint64 `json:"room_match_token"`
}

func (this *NotifyMatchPlayerInfo_Out) Process(session *net.Session) {
	g_OutHandler.NotifyMatchPlayerInfo(session, this)
}

func (this *NotifyMatchPlayerInfo_Out) TypeName() string {
	return "multi_level.notify_match_player_info.out"
}

func (this *NotifyMatchPlayerInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 16, 27
}

func (this *NotifyMatchPlayerInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetMatchInfo_In struct {
}

func (this *GetMatchInfo_In) Process(session *net.Session) {
	g_InHandler.GetMatchInfo(session, this)
}

func (this *GetMatchInfo_In) TypeName() string {
	return "multi_level.get_match_info.in"
}

func (this *GetMatchInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 16, 28
}

type GetMatchInfo_Out struct {
	MatchStatus []GetMatchInfo_Out_MatchStatus `json:"match_status"`
}

type GetMatchInfo_Out_MatchStatus struct {
	LevelId   int32 `json:"level_id"`
	PosNum    int32 `json:"pos_num"`
	PlayerNum int32 `json:"player_num"`
}

func (this *GetMatchInfo_Out) Process(session *net.Session) {
	g_OutHandler.GetMatchInfo(session, this)
}

func (this *GetMatchInfo_Out) TypeName() string {
	return "multi_level.get_match_info.out"
}

func (this *GetMatchInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 16, 28
}

func (this *GetMatchInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type NotifyFormInfo_Out struct {
	Formation []NotifyFormInfo_Out_Formation `json:"formation"`
}

type NotifyFormInfo_Out_Formation struct {
	Pos    int8  `json:"pos"`
	Pid    int64 `json:"pid"`
	RoleId int8  `json:"role_id"`
}

func (this *NotifyFormInfo_Out) Process(session *net.Session) {
	g_OutHandler.NotifyFormInfo(session, this)
}

func (this *NotifyFormInfo_Out) TypeName() string {
	return "multi_level.notify_form_info.out"
}

func (this *NotifyFormInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 16, 29
}

func (this *NotifyFormInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type CancelMatchPlayer_In struct {
}

func (this *CancelMatchPlayer_In) Process(session *net.Session) {
	g_InHandler.CancelMatchPlayer(session, this)
}

func (this *CancelMatchPlayer_In) TypeName() string {
	return "multi_level.cancel_match_player.in"
}

func (this *CancelMatchPlayer_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 16, 30
}

type CancelMatchPlayer_Out struct {
}

func (this *CancelMatchPlayer_Out) Process(session *net.Session) {
	g_OutHandler.CancelMatchPlayer(session, this)
}

func (this *CancelMatchPlayer_Out) TypeName() string {
	return "multi_level.cancel_match_player.out"
}

func (this *CancelMatchPlayer_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 16, 30
}

func (this *CancelMatchPlayer_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

func (this *CreateRoom_In) Decode(buffer *net.Buffer) {
	this.LevelId = int32(buffer.ReadUint32LE())
}

func (this *CreateRoom_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(16)
	buffer.WriteUint8(1)
	buffer.WriteUint32LE(uint32(this.LevelId))
}

func (this *CreateRoom_In) ByteSize() int {
	size := 6
	return size
}

func (this *CreateRoom_Out) Decode(buffer *net.Buffer) {
	this.Status = RoomStatus(buffer.ReadUint8())
}

func (this *CreateRoom_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(16)
	buffer.WriteUint8(1)
	buffer.WriteUint8(uint8(this.Status))
}

func (this *CreateRoom_Out) ByteSize() int {
	size := 3
	return size
}

func (this *EnterRoom_In) Decode(buffer *net.Buffer) {
	this.RoomId = int64(buffer.ReadUint64LE())
	this.RoomToken = buffer.ReadUint64LE()
}

func (this *EnterRoom_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(16)
	buffer.WriteUint8(2)
	buffer.WriteUint64LE(uint64(this.RoomId))
	buffer.WriteUint64LE(this.RoomToken)
}

func (this *EnterRoom_In) ByteSize() int {
	size := 18
	return size
}

func (this *EnterRoom_Out) Decode(buffer *net.Buffer) {
	this.Status = RoomStatus(buffer.ReadUint8())
}

func (this *EnterRoom_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(16)
	buffer.WriteUint8(2)
	buffer.WriteUint8(uint8(this.Status))
}

func (this *EnterRoom_Out) ByteSize() int {
	size := 3
	return size
}

func (this *NotifyRoomInfo_Out) Decode(buffer *net.Buffer) {
	this.RoomId = int64(buffer.ReadUint64LE())
	this.LeaderPid = int64(buffer.ReadUint64LE())
	this.LevelId = int32(buffer.ReadUint32LE())
	this.Partners = make([]NotifyRoomInfo_Out_Partners, buffer.ReadUint8())
	for i := 0; i < len(this.Partners); i++ {
		this.Partners[i].Decode(buffer)
	}
}

func (this *NotifyRoomInfo_Out_Partners) Decode(buffer *net.Buffer) {
	this.Partner.Decode(buffer)
}

func (this *NotifyRoomInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(16)
	buffer.WriteUint8(4)
	buffer.WriteUint64LE(uint64(this.RoomId))
	buffer.WriteUint64LE(uint64(this.LeaderPid))
	buffer.WriteUint32LE(uint32(this.LevelId))
	buffer.WriteUint8(uint8(len(this.Partners)))
	for i := 0; i < len(this.Partners); i++ {
		this.Partners[i].Encode(buffer)
	}
}

func (this *NotifyRoomInfo_Out_Partners) Encode(buffer *net.Buffer) {
	this.Partner.Encode(buffer)
}

func (this *NotifyRoomInfo_Out) ByteSize() int {
	size := 23
	for i := 0; i < len(this.Partners); i++ {
		size += this.Partners[i].ByteSize()
	}
	return size
}

func (this *NotifyRoomInfo_Out_Partners) ByteSize() int {
	size := 0
	size += this.Partner.ByteSize()
	return size
}

func (this *LeaveRoom_In) Decode(buffer *net.Buffer) {
}

func (this *LeaveRoom_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(16)
	buffer.WriteUint8(5)
}

func (this *LeaveRoom_In) ByteSize() int {
	size := 2
	return size
}

func (this *LeaveRoom_Out) Decode(buffer *net.Buffer) {
	this.Pid = int64(buffer.ReadUint64LE())
	this.LeaderPid = int64(buffer.ReadUint64LE())
	this.Tokens = make([]LeaveRoom_Out_Tokens, buffer.ReadUint8())
	for i := 0; i < len(this.Tokens); i++ {
		this.Tokens[i].Decode(buffer)
	}
}

func (this *LeaveRoom_Out_Tokens) Decode(buffer *net.Buffer) {
	this.Pos = int8(buffer.ReadUint8())
	this.Token = buffer.ReadUint64LE()
}

func (this *LeaveRoom_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(16)
	buffer.WriteUint8(5)
	buffer.WriteUint64LE(uint64(this.Pid))
	buffer.WriteUint64LE(uint64(this.LeaderPid))
	buffer.WriteUint8(uint8(len(this.Tokens)))
	for i := 0; i < len(this.Tokens); i++ {
		this.Tokens[i].Encode(buffer)
	}
}

func (this *LeaveRoom_Out_Tokens) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(this.Pos))
	buffer.WriteUint64LE(this.Token)
}

func (this *LeaveRoom_Out) ByteSize() int {
	size := 19
	size += len(this.Tokens) * 9
	return size
}

func (this *NotifyJoinPartner_Out) Decode(buffer *net.Buffer) {
	this.Partner.Decode(buffer)
}

func (this *NotifyJoinPartner_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(16)
	buffer.WriteUint8(6)
	this.Partner.Encode(buffer)
}

func (this *NotifyJoinPartner_Out) ByteSize() int {
	size := 2
	size += this.Partner.ByteSize()
	return size
}

func (this *ChangeBuddy_In) Decode(buffer *net.Buffer) {
	this.BuddyRoleId = int8(buffer.ReadUint8())
}

func (this *ChangeBuddy_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(16)
	buffer.WriteUint8(7)
	buffer.WriteUint8(uint8(this.BuddyRoleId))
}

func (this *ChangeBuddy_In) ByteSize() int {
	size := 3
	return size
}

func (this *ChangeBuddy_Out) Decode(buffer *net.Buffer) {
	this.Success = buffer.ReadUint8() == 1
}

func (this *ChangeBuddy_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(16)
	buffer.WriteUint8(7)
	if this.Success {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *ChangeBuddy_Out) ByteSize() int {
	size := 3
	return size
}

func (this *GetBaseInfo_In) Decode(buffer *net.Buffer) {
}

func (this *GetBaseInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(16)
	buffer.WriteUint8(10)
}

func (this *GetBaseInfo_In) ByteSize() int {
	size := 2
	return size
}

func (this *GetBaseInfo_Out) Decode(buffer *net.Buffer) {
	this.DailyNum = int8(buffer.ReadUint8())
	this.Lock = int32(buffer.ReadUint32LE())
}

func (this *GetBaseInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(16)
	buffer.WriteUint8(10)
	buffer.WriteUint8(uint8(this.DailyNum))
	buffer.WriteUint32LE(uint32(this.Lock))
}

func (this *GetBaseInfo_Out) ByteSize() int {
	size := 7
	return size
}

func (this *GetOnlineFriend_In) Decode(buffer *net.Buffer) {
}

func (this *GetOnlineFriend_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(16)
	buffer.WriteUint8(12)
}

func (this *GetOnlineFriend_In) ByteSize() int {
	size := 2
	return size
}

func (this *GetOnlineFriend_Out) Decode(buffer *net.Buffer) {
	this.Friends = make([]GetOnlineFriend_Out_Friends, buffer.ReadUint8())
	for i := 0; i < len(this.Friends); i++ {
		this.Friends[i].Decode(buffer)
	}
}

func (this *GetOnlineFriend_Out_Friends) Decode(buffer *net.Buffer) {
	this.Pid = int64(buffer.ReadUint64LE())
	this.RoleId = int8(buffer.ReadUint8())
	this.Nickname = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Level = int16(buffer.ReadUint16LE())
	this.FightNum = int32(buffer.ReadUint32LE())
	this.DailyNum = int8(buffer.ReadUint8())
	this.Lock = int32(buffer.ReadUint32LE())
}

func (this *GetOnlineFriend_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(16)
	buffer.WriteUint8(12)
	buffer.WriteUint8(uint8(len(this.Friends)))
	for i := 0; i < len(this.Friends); i++ {
		this.Friends[i].Encode(buffer)
	}
}

func (this *GetOnlineFriend_Out_Friends) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.Pid))
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint16LE(uint16(len(this.Nickname)))
	buffer.WriteBytes(this.Nickname)
	buffer.WriteUint16LE(uint16(this.Level))
	buffer.WriteUint32LE(uint32(this.FightNum))
	buffer.WriteUint8(uint8(this.DailyNum))
	buffer.WriteUint32LE(uint32(this.Lock))
}

func (this *GetOnlineFriend_Out) ByteSize() int {
	size := 3
	for i := 0; i < len(this.Friends); i++ {
		size += this.Friends[i].ByteSize()
	}
	return size
}

func (this *GetOnlineFriend_Out_Friends) ByteSize() int {
	size := 22
	size += len(this.Nickname)
	return size
}

func (this *InviteIntoRoom_In) Decode(buffer *net.Buffer) {
	this.Pid = int64(buffer.ReadUint64LE())
}

func (this *InviteIntoRoom_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(16)
	buffer.WriteUint8(13)
	buffer.WriteUint64LE(uint64(this.Pid))
}

func (this *InviteIntoRoom_In) ByteSize() int {
	size := 10
	return size
}

func (this *InviteIntoRoom_Out) Decode(buffer *net.Buffer) {
	this.IsOffline = buffer.ReadUint8() == 1
}

func (this *InviteIntoRoom_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(16)
	buffer.WriteUint8(13)
	if this.IsOffline {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *InviteIntoRoom_Out) ByteSize() int {
	size := 3
	return size
}

func (this *NotifyRoomInvite_Out) Decode(buffer *net.Buffer) {
	this.RoomId = int64(buffer.ReadUint64LE())
	this.LevelId = int32(buffer.ReadUint32LE())
	this.Nickname = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.InviterId = int64(buffer.ReadUint64LE())
	this.DailyNum = int8(buffer.ReadUint8())
	this.GameServerId = int64(buffer.ReadUint64LE())
}

func (this *NotifyRoomInvite_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(16)
	buffer.WriteUint8(14)
	buffer.WriteUint64LE(uint64(this.RoomId))
	buffer.WriteUint32LE(uint32(this.LevelId))
	buffer.WriteUint16LE(uint16(len(this.Nickname)))
	buffer.WriteBytes(this.Nickname)
	buffer.WriteUint64LE(uint64(this.InviterId))
	buffer.WriteUint8(uint8(this.DailyNum))
	buffer.WriteUint64LE(uint64(this.GameServerId))
}

func (this *NotifyRoomInvite_Out) ByteSize() int {
	size := 33
	size += len(this.Nickname)
	return size
}

func (this *NotifyPlayersInfo_Out) Decode(buffer *net.Buffer) {
	this.Players = make([]NotifyPlayersInfo_Out_Players, buffer.ReadUint8())
	for i := 0; i < len(this.Players); i++ {
		this.Players[i].Decode(buffer)
	}
}

func (this *NotifyPlayersInfo_Out_Players) Decode(buffer *net.Buffer) {
	this.PlayerId = int64(buffer.ReadUint64LE())
	this.Nickname = buffer.ReadBytes(int(buffer.ReadUint16LE()))
}

func (this *NotifyPlayersInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(16)
	buffer.WriteUint8(15)
	buffer.WriteUint8(uint8(len(this.Players)))
	for i := 0; i < len(this.Players); i++ {
		this.Players[i].Encode(buffer)
	}
}

func (this *NotifyPlayersInfo_Out_Players) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.PlayerId))
	buffer.WriteUint16LE(uint16(len(this.Nickname)))
	buffer.WriteBytes(this.Nickname)
}

func (this *NotifyPlayersInfo_Out) ByteSize() int {
	size := 3
	for i := 0; i < len(this.Players); i++ {
		size += this.Players[i].ByteSize()
	}
	return size
}

func (this *NotifyPlayersInfo_Out_Players) ByteSize() int {
	size := 10
	size += len(this.Nickname)
	return size
}

func (this *RefuseRoomInvite_In) Decode(buffer *net.Buffer) {
	this.RoomId = int64(buffer.ReadUint64LE())
	this.InviterId = int64(buffer.ReadUint64LE())
}

func (this *RefuseRoomInvite_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(16)
	buffer.WriteUint8(16)
	buffer.WriteUint64LE(uint64(this.RoomId))
	buffer.WriteUint64LE(uint64(this.InviterId))
}

func (this *RefuseRoomInvite_In) ByteSize() int {
	size := 18
	return size
}

func (this *RefuseRoomInvite_Out) Decode(buffer *net.Buffer) {
}

func (this *RefuseRoomInvite_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(16)
	buffer.WriteUint8(16)
}

func (this *RefuseRoomInvite_Out) ByteSize() int {
	size := 2
	return size
}

func (this *NotifyRoomInviteRefuse_Out) Decode(buffer *net.Buffer) {
	this.Nickname = buffer.ReadBytes(int(buffer.ReadUint16LE()))
}

func (this *NotifyRoomInviteRefuse_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(16)
	buffer.WriteUint8(17)
	buffer.WriteUint16LE(uint16(len(this.Nickname)))
	buffer.WriteBytes(this.Nickname)
}

func (this *NotifyRoomInviteRefuse_Out) ByteSize() int {
	size := 4
	size += len(this.Nickname)
	return size
}

func (this *NotifyUpdatePartner_Out) Decode(buffer *net.Buffer) {
	this.Partner.Decode(buffer)
}

func (this *NotifyUpdatePartner_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(16)
	buffer.WriteUint8(18)
	this.Partner.Encode(buffer)
}

func (this *NotifyUpdatePartner_Out) ByteSize() int {
	size := 2
	size += this.Partner.ByteSize()
	return size
}

func (this *MatchPlayer_In) Decode(buffer *net.Buffer) {
	this.Level = int32(buffer.ReadUint32LE())
}

func (this *MatchPlayer_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(16)
	buffer.WriteUint8(19)
	buffer.WriteUint32LE(uint32(this.Level))
}

func (this *MatchPlayer_In) ByteSize() int {
	size := 6
	return size
}

func (this *MatchPlayer_Out) Decode(buffer *net.Buffer) {
	this.Token = buffer.ReadUint64LE()
	this.Ok = buffer.ReadUint8() == 1
}

func (this *MatchPlayer_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(16)
	buffer.WriteUint8(19)
	buffer.WriteUint64LE(this.Token)
	if this.Ok {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *MatchPlayer_Out) ByteSize() int {
	size := 11
	return size
}

func (this *NotifyMatchPlayerSuccess_Out) Decode(buffer *net.Buffer) {
	this.ServerId = int32(buffer.ReadUint32LE())
	this.RoomId = int64(buffer.ReadUint64LE())
	this.PlayerMatchToken = buffer.ReadUint64LE()
	this.RoomMatchToken = buffer.ReadUint64LE()
}

func (this *NotifyMatchPlayerSuccess_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(16)
	buffer.WriteUint8(21)
	buffer.WriteUint32LE(uint32(this.ServerId))
	buffer.WriteUint64LE(uint64(this.RoomId))
	buffer.WriteUint64LE(this.PlayerMatchToken)
	buffer.WriteUint64LE(this.RoomMatchToken)
}

func (this *NotifyMatchPlayerSuccess_Out) ByteSize() int {
	size := 30
	return size
}

func (this *MatchRoom_In) Decode(buffer *net.Buffer) {
	this.Pos = int8(buffer.ReadUint8())
}

func (this *MatchRoom_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(16)
	buffer.WriteUint8(22)
	buffer.WriteUint8(uint8(this.Pos))
}

func (this *MatchRoom_In) ByteSize() int {
	size := 3
	return size
}

func (this *MatchRoom_Out) Decode(buffer *net.Buffer) {
	this.MatchToken = buffer.ReadUint64LE()
	this.Ok = buffer.ReadUint8() == 1
	this.Pos = int8(buffer.ReadUint8())
}

func (this *MatchRoom_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(16)
	buffer.WriteUint8(22)
	buffer.WriteUint64LE(this.MatchToken)
	if this.Ok {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
	buffer.WriteUint8(uint8(this.Pos))
}

func (this *MatchRoom_Out) ByteSize() int {
	size := 12
	return size
}

func (this *CancelMatchRoom_In) Decode(buffer *net.Buffer) {
	this.MatchToken = buffer.ReadUint64LE()
}

func (this *CancelMatchRoom_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(16)
	buffer.WriteUint8(23)
	buffer.WriteUint64LE(this.MatchToken)
}

func (this *CancelMatchRoom_In) ByteSize() int {
	size := 10
	return size
}

func (this *CancelMatchRoom_Out) Decode(buffer *net.Buffer) {
}

func (this *CancelMatchRoom_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(16)
	buffer.WriteUint8(23)
}

func (this *CancelMatchRoom_Out) ByteSize() int {
	size := 2
	return size
}

func (this *GetFormInfo_In) Decode(buffer *net.Buffer) {
}

func (this *GetFormInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(16)
	buffer.WriteUint8(24)
}

func (this *GetFormInfo_In) ByteSize() int {
	size := 2
	return size
}

func (this *GetFormInfo_Out) Decode(buffer *net.Buffer) {
	this.Formation = make([]GetFormInfo_Out_Formation, buffer.ReadUint8())
	for i := 0; i < len(this.Formation); i++ {
		this.Formation[i].Decode(buffer)
	}
}

func (this *GetFormInfo_Out_Formation) Decode(buffer *net.Buffer) {
	this.Pos = int8(buffer.ReadUint8())
	this.Pid = int64(buffer.ReadUint64LE())
	this.RoleId = int8(buffer.ReadUint8())
}

func (this *GetFormInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(16)
	buffer.WriteUint8(24)
	buffer.WriteUint8(uint8(len(this.Formation)))
	for i := 0; i < len(this.Formation); i++ {
		this.Formation[i].Encode(buffer)
	}
}

func (this *GetFormInfo_Out_Formation) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(this.Pos))
	buffer.WriteUint64LE(uint64(this.Pid))
	buffer.WriteUint8(uint8(this.RoleId))
}

func (this *GetFormInfo_Out) ByteSize() int {
	size := 3
	size += len(this.Formation) * 10
	return size
}

func (this *SetForm_In) Decode(buffer *net.Buffer) {
	this.Formation = make([]SetForm_In_Formation, buffer.ReadUint8())
	for i := 0; i < len(this.Formation); i++ {
		this.Formation[i].Decode(buffer)
	}
}

func (this *SetForm_In_Formation) Decode(buffer *net.Buffer) {
	this.Pos = int8(buffer.ReadUint8())
	this.Pid = int64(buffer.ReadUint64LE())
	this.IsMainRole = buffer.ReadUint8() == 1
}

func (this *SetForm_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(16)
	buffer.WriteUint8(25)
	buffer.WriteUint8(uint8(len(this.Formation)))
	for i := 0; i < len(this.Formation); i++ {
		this.Formation[i].Encode(buffer)
	}
}

func (this *SetForm_In_Formation) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(this.Pos))
	buffer.WriteUint64LE(uint64(this.Pid))
	if this.IsMainRole {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *SetForm_In) ByteSize() int {
	size := 3
	size += len(this.Formation) * 10
	return size
}

func (this *SetForm_Out) Decode(buffer *net.Buffer) {
}

func (this *SetForm_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(16)
	buffer.WriteUint8(25)
}

func (this *SetForm_Out) ByteSize() int {
	size := 2
	return size
}

func (this *GetBattleRoleInfo_In) Decode(buffer *net.Buffer) {
}

func (this *GetBattleRoleInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(16)
	buffer.WriteUint8(26)
}

func (this *GetBattleRoleInfo_In) ByteSize() int {
	size := 2
	return size
}

func (this *GetBattleRoleInfo_Out) Decode(buffer *net.Buffer) {
	this.MainRoleId = int8(buffer.ReadUint8())
	this.BuddyRoleId = int8(buffer.ReadUint8())
}

func (this *GetBattleRoleInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(16)
	buffer.WriteUint8(26)
	buffer.WriteUint8(uint8(this.MainRoleId))
	buffer.WriteUint8(uint8(this.BuddyRoleId))
}

func (this *GetBattleRoleInfo_Out) ByteSize() int {
	size := 4
	return size
}

func (this *NotifyMatchPlayerInfo_Out) Decode(buffer *net.Buffer) {
	this.RoomMatchToken = buffer.ReadUint64LE()
}

func (this *NotifyMatchPlayerInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(16)
	buffer.WriteUint8(27)
	buffer.WriteUint64LE(this.RoomMatchToken)
}

func (this *NotifyMatchPlayerInfo_Out) ByteSize() int {
	size := 10
	return size
}

func (this *GetMatchInfo_In) Decode(buffer *net.Buffer) {
}

func (this *GetMatchInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(16)
	buffer.WriteUint8(28)
}

func (this *GetMatchInfo_In) ByteSize() int {
	size := 2
	return size
}

func (this *GetMatchInfo_Out) Decode(buffer *net.Buffer) {
	this.MatchStatus = make([]GetMatchInfo_Out_MatchStatus, buffer.ReadUint8())
	for i := 0; i < len(this.MatchStatus); i++ {
		this.MatchStatus[i].Decode(buffer)
	}
}

func (this *GetMatchInfo_Out_MatchStatus) Decode(buffer *net.Buffer) {
	this.LevelId = int32(buffer.ReadUint32LE())
	this.PosNum = int32(buffer.ReadUint32LE())
	this.PlayerNum = int32(buffer.ReadUint32LE())
}

func (this *GetMatchInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(16)
	buffer.WriteUint8(28)
	buffer.WriteUint8(uint8(len(this.MatchStatus)))
	for i := 0; i < len(this.MatchStatus); i++ {
		this.MatchStatus[i].Encode(buffer)
	}
}

func (this *GetMatchInfo_Out_MatchStatus) Encode(buffer *net.Buffer) {
	buffer.WriteUint32LE(uint32(this.LevelId))
	buffer.WriteUint32LE(uint32(this.PosNum))
	buffer.WriteUint32LE(uint32(this.PlayerNum))
}

func (this *GetMatchInfo_Out) ByteSize() int {
	size := 3
	size += len(this.MatchStatus) * 12
	return size
}

func (this *NotifyFormInfo_Out) Decode(buffer *net.Buffer) {
	this.Formation = make([]NotifyFormInfo_Out_Formation, buffer.ReadUint8())
	for i := 0; i < len(this.Formation); i++ {
		this.Formation[i].Decode(buffer)
	}
}

func (this *NotifyFormInfo_Out_Formation) Decode(buffer *net.Buffer) {
	this.Pos = int8(buffer.ReadUint8())
	this.Pid = int64(buffer.ReadUint64LE())
	this.RoleId = int8(buffer.ReadUint8())
}

func (this *NotifyFormInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(16)
	buffer.WriteUint8(29)
	buffer.WriteUint8(uint8(len(this.Formation)))
	for i := 0; i < len(this.Formation); i++ {
		this.Formation[i].Encode(buffer)
	}
}

func (this *NotifyFormInfo_Out_Formation) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(this.Pos))
	buffer.WriteUint64LE(uint64(this.Pid))
	buffer.WriteUint8(uint8(this.RoleId))
}

func (this *NotifyFormInfo_Out) ByteSize() int {
	size := 3
	size += len(this.Formation) * 10
	return size
}

func (this *CancelMatchPlayer_In) Decode(buffer *net.Buffer) {
}

func (this *CancelMatchPlayer_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(16)
	buffer.WriteUint8(30)
}

func (this *CancelMatchPlayer_In) ByteSize() int {
	size := 2
	return size
}

func (this *CancelMatchPlayer_Out) Decode(buffer *net.Buffer) {
}

func (this *CancelMatchPlayer_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(16)
	buffer.WriteUint8(30)
}

func (this *CancelMatchPlayer_Out) ByteSize() int {
	size := 2
	return size
}
