package friend_api

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
	GetFriendList(*net.Session, *GetFriendList_In)
	ListenByNick(*net.Session, *ListenByNick_In)
	CancelListen(*net.Session, *CancelListen_In)
	SendHeart(*net.Session, *SendHeart_In)
	Chat(*net.Session, *Chat_In)
	GetChatHistory(*net.Session, *GetChatHistory_In)
	Block(*net.Session, *Block_In)
	CancelBlock(*net.Session, *CancelBlock_In)
	CleanChatHistory(*net.Session, *CleanChatHistory_In)
	CurrentPlatformFriendNum(*net.Session, *CurrentPlatformFriendNum_In)
	GetSendHeartList(*net.Session, *GetSendHeartList_In)
	SendHeartToAllFriends(*net.Session, *SendHeartToAllFriends_In)
}

type OutHandler interface {
	GetFriendList(*net.Session, *GetFriendList_Out)
	ListenByNick(*net.Session, *ListenByNick_Out)
	CancelListen(*net.Session, *CancelListen_Out)
	SendHeart(*net.Session, *SendHeart_Out)
	Chat(*net.Session, *Chat_Out)
	GetChatHistory(*net.Session, *GetChatHistory_Out)
	GetOfflineMessages(*net.Session, *GetOfflineMessages_Out)
	Block(*net.Session, *Block_Out)
	CancelBlock(*net.Session, *CancelBlock_Out)
	CleanChatHistory(*net.Session, *CleanChatHistory_Out)
	NotifyListenedState(*net.Session, *NotifyListenedState_Out)
	CurrentPlatformFriendNum(*net.Session, *CurrentPlatformFriendNum_Out)
	GetSendHeartList(*net.Session, *GetSendHeartList_Out)
	SendHeartToAllFriends(*net.Session, *SendHeartToAllFriends_Out)
}

func DecodeIn(message []byte) Request {
	var actionId = message[0]
	var buffer = net.NewBuffer(message[1:])
	switch actionId {
	case 0:
		request := new(GetFriendList_In)
		request.Decode(buffer)
		return request
	case 1:
		request := new(ListenByNick_In)
		request.Decode(buffer)
		return request
	case 2:
		request := new(CancelListen_In)
		request.Decode(buffer)
		return request
	case 3:
		request := new(SendHeart_In)
		request.Decode(buffer)
		return request
	case 4:
		request := new(Chat_In)
		request.Decode(buffer)
		return request
	case 5:
		request := new(GetChatHistory_In)
		request.Decode(buffer)
		return request
	case 7:
		request := new(Block_In)
		request.Decode(buffer)
		return request
	case 8:
		request := new(CancelBlock_In)
		request.Decode(buffer)
		return request
	case 9:
		request := new(CleanChatHistory_In)
		request.Decode(buffer)
		return request
	case 11:
		request := new(CurrentPlatformFriendNum_In)
		request.Decode(buffer)
		return request
	case 12:
		request := new(GetSendHeartList_In)
		request.Decode(buffer)
		return request
	case 13:
		request := new(SendHeartToAllFriends_In)
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
		request := new(GetFriendList_Out)
		request.Decode(buffer)
		return request
	case 1:
		request := new(ListenByNick_Out)
		request.Decode(buffer)
		return request
	case 2:
		request := new(CancelListen_Out)
		request.Decode(buffer)
		return request
	case 3:
		request := new(SendHeart_Out)
		request.Decode(buffer)
		return request
	case 4:
		request := new(Chat_Out)
		request.Decode(buffer)
		return request
	case 5:
		request := new(GetChatHistory_Out)
		request.Decode(buffer)
		return request
	case 6:
		request := new(GetOfflineMessages_Out)
		request.Decode(buffer)
		return request
	case 7:
		request := new(Block_Out)
		request.Decode(buffer)
		return request
	case 8:
		request := new(CancelBlock_Out)
		request.Decode(buffer)
		return request
	case 9:
		request := new(CleanChatHistory_Out)
		request.Decode(buffer)
		return request
	case 10:
		request := new(NotifyListenedState_Out)
		request.Decode(buffer)
		return request
	case 11:
		request := new(CurrentPlatformFriendNum_Out)
		request.Decode(buffer)
		return request
	case 12:
		request := new(GetSendHeartList_Out)
		request.Decode(buffer)
		return request
	case 13:
		request := new(SendHeartToAllFriends_Out)
		request.Decode(buffer)
		return request
	}
	_ = buffer
	panic("unsupported response")
}

type FriendType int8

const (
	FRIEND_TYPE_GAME_FRIEND     FriendType = 1
	FRIEND_TYPE_PLATFORM_FRIEND FriendType = 2
)

type FriendMode int8

const (
	FRIEND_MODE_STRANGE       FriendMode = 0
	FRIEND_MODE_LISTEN_ONLY   FriendMode = 1
	FRIEND_MODE_LISTENED_ONLY FriendMode = 2
	FRIEND_MODE_FRIEND        FriendMode = 3
)

type AddResult int8

const (
	ADD_RESULT_SUCCEED              AddResult = 0
	ADD_RESULT_FAILED_ADD_SELF      AddResult = 1
	ADD_RESULT_FAILED_ADD_NOT_EXIST AddResult = 2
	ADD_RESULT_FAILED_ADD_FOLLOW    AddResult = 3
	ADD_RESULT_FAILED_ADD_FULL      AddResult = 4
	ADD_RESULT_FAILED_TARGET_FULL   AddResult = 5
)

type ListendState int8

const (
	LISTEND_STATE_CANCEL  ListendState = 1
	LISTEND_STATE_LISTEND ListendState = 2
)

type MsgMode int8

const (
	MSG_MODE_SEND    MsgMode = 1
	MSG_MODE_RECEIVE MsgMode = 2
)

type GetFriendList_In struct {
}

func (this *GetFriendList_In) Process(session *net.Session) {
	g_InHandler.GetFriendList(session, this)
}

func (this *GetFriendList_In) TypeName() string {
	return "friend.get_friend_list.in"
}

func (this *GetFriendList_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 14, 0
}

type GetFriendList_Out struct {
	CancelListenCount int32                       `json:"cancel_listen_count"`
	PlatformFriendNum int32                       `json:"platform_friend_num"`
	Friends           []GetFriendList_Out_Friends `json:"friends"`
}

type GetFriendList_Out_Friends struct {
	Pid          int64      `json:"pid"`
	RoleId       int8       `json:"role_id"`
	Nickname     []byte     `json:"nickname"`
	Level        int16      `json:"level"`
	FightNum     int32      `json:"fight_num"`
	FriendMode   FriendMode `json:"friend_mode"`
	BlockMode    int8       `json:"block_mode"`
	LastChatTime int64      `json:"last_chat_time"`
	FriendTime   int64      `json:"friend_time"`
}

func (this *GetFriendList_Out) Process(session *net.Session) {
	g_OutHandler.GetFriendList(session, this)
}

func (this *GetFriendList_Out) TypeName() string {
	return "friend.get_friend_list.out"
}

func (this *GetFriendList_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 14, 0
}

func (this *GetFriendList_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ListenByNick_In struct {
	Nick []byte `json:"nick"`
}

func (this *ListenByNick_In) Process(session *net.Session) {
	g_InHandler.ListenByNick(session, this)
}

func (this *ListenByNick_In) TypeName() string {
	return "friend.listen_by_nick.in"
}

func (this *ListenByNick_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 14, 1
}

type ListenByNick_Out struct {
	Result   AddResult `json:"result"`
	RoleId   int8      `json:"role_id"`
	Nickname []byte    `json:"nickname"`
	Level    int16     `json:"level"`
	FightNum int32     `json:"fight_num"`
}

func (this *ListenByNick_Out) Process(session *net.Session) {
	g_OutHandler.ListenByNick(session, this)
}

func (this *ListenByNick_Out) TypeName() string {
	return "friend.listen_by_nick.out"
}

func (this *ListenByNick_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 14, 1
}

func (this *ListenByNick_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type CancelListen_In struct {
	Pid int64 `json:"pid"`
}

func (this *CancelListen_In) Process(session *net.Session) {
	g_InHandler.CancelListen(session, this)
}

func (this *CancelListen_In) TypeName() string {
	return "friend.cancel_listen.in"
}

func (this *CancelListen_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 14, 2
}

type CancelListen_Out struct {
	Result bool `json:"result"`
}

func (this *CancelListen_Out) Process(session *net.Session) {
	g_OutHandler.CancelListen(session, this)
}

func (this *CancelListen_Out) TypeName() string {
	return "friend.cancel_listen.out"
}

func (this *CancelListen_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 14, 2
}

func (this *CancelListen_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type SendHeart_In struct {
	Nickname   []byte     `json:"nickname"`
	FriendType FriendType `json:"friend_type"`
	Pid        int64      `json:"pid"`
}

func (this *SendHeart_In) Process(session *net.Session) {
	g_InHandler.SendHeart(session, this)
}

func (this *SendHeart_In) TypeName() string {
	return "friend.send_heart.in"
}

func (this *SendHeart_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 14, 3
}

type SendHeart_Out struct {
}

func (this *SendHeart_Out) Process(session *net.Session) {
	g_OutHandler.SendHeart(session, this)
}

func (this *SendHeart_Out) TypeName() string {
	return "friend.send_heart.out"
}

func (this *SendHeart_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 14, 3
}

func (this *SendHeart_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type Chat_In struct {
	Pid     int64  `json:"pid"`
	Message []byte `json:"message"`
}

func (this *Chat_In) Process(session *net.Session) {
	g_InHandler.Chat(session, this)
}

func (this *Chat_In) TypeName() string {
	return "friend.chat.in"
}

func (this *Chat_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 14, 4
}

type Chat_Out struct {
	Banned bool `json:"banned"`
}

func (this *Chat_Out) Process(session *net.Session) {
	g_OutHandler.Chat(session, this)
}

func (this *Chat_Out) TypeName() string {
	return "friend.chat.out"
}

func (this *Chat_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 14, 4
}

func (this *Chat_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetChatHistory_In struct {
	Pid int64 `json:"pid"`
}

func (this *GetChatHistory_In) Process(session *net.Session) {
	g_InHandler.GetChatHistory(session, this)
}

func (this *GetChatHistory_In) TypeName() string {
	return "friend.get_chat_history.in"
}

func (this *GetChatHistory_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 14, 5
}

type GetChatHistory_Out struct {
	Messages []GetChatHistory_Out_Messages `json:"messages"`
}

type GetChatHistory_Out_Messages struct {
	Mode     MsgMode `json:"mode"`
	Id       int64   `json:"id"`
	SendTime int64   `json:"send_time"`
	Message  []byte  `json:"message"`
}

func (this *GetChatHistory_Out) Process(session *net.Session) {
	g_OutHandler.GetChatHistory(session, this)
}

func (this *GetChatHistory_Out) TypeName() string {
	return "friend.get_chat_history.out"
}

func (this *GetChatHistory_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 14, 5
}

func (this *GetChatHistory_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetOfflineMessages_Out struct {
	Chats    []GetOfflineMessages_Out_Chats    `json:"chats"`
	Listener []GetOfflineMessages_Out_Listener `json:"listener"`
}

type GetOfflineMessages_Out_Chats struct {
	Pid      int64  `json:"pid"`
	Nickname []byte `json:"nickname"`
	RoleId   int8   `json:"role_id"`
	Level    int16  `json:"level"`
}

type GetOfflineMessages_Out_Listener struct {
	Pid  int64  `json:"pid"`
	Nick []byte `json:"nick"`
}

func (this *GetOfflineMessages_Out) Process(session *net.Session) {
	g_OutHandler.GetOfflineMessages(session, this)
}

func (this *GetOfflineMessages_Out) TypeName() string {
	return "friend.get_offline_messages.out"
}

func (this *GetOfflineMessages_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 14, 6
}

func (this *GetOfflineMessages_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type Block_In struct {
	Pid int64 `json:"pid"`
}

func (this *Block_In) Process(session *net.Session) {
	g_InHandler.Block(session, this)
}

func (this *Block_In) TypeName() string {
	return "friend.block.in"
}

func (this *Block_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 14, 7
}

type Block_Out struct {
}

func (this *Block_Out) Process(session *net.Session) {
	g_OutHandler.Block(session, this)
}

func (this *Block_Out) TypeName() string {
	return "friend.block.out"
}

func (this *Block_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 14, 7
}

func (this *Block_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type CancelBlock_In struct {
	Pid int64 `json:"pid"`
}

func (this *CancelBlock_In) Process(session *net.Session) {
	g_InHandler.CancelBlock(session, this)
}

func (this *CancelBlock_In) TypeName() string {
	return "friend.cancel_block.in"
}

func (this *CancelBlock_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 14, 8
}

type CancelBlock_Out struct {
}

func (this *CancelBlock_Out) Process(session *net.Session) {
	g_OutHandler.CancelBlock(session, this)
}

func (this *CancelBlock_Out) TypeName() string {
	return "friend.cancel_block.out"
}

func (this *CancelBlock_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 14, 8
}

func (this *CancelBlock_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type CleanChatHistory_In struct {
	Pid int64 `json:"pid"`
}

func (this *CleanChatHistory_In) Process(session *net.Session) {
	g_InHandler.CleanChatHistory(session, this)
}

func (this *CleanChatHistory_In) TypeName() string {
	return "friend.clean_chat_history.in"
}

func (this *CleanChatHistory_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 14, 9
}

type CleanChatHistory_Out struct {
}

func (this *CleanChatHistory_Out) Process(session *net.Session) {
	g_OutHandler.CleanChatHistory(session, this)
}

func (this *CleanChatHistory_Out) TypeName() string {
	return "friend.clean_chat_history.out"
}

func (this *CleanChatHistory_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 14, 9
}

func (this *CleanChatHistory_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type NotifyListenedState_Out struct {
	Pid   int64        `json:"pid"`
	Nick  []byte       `json:"nick"`
	State ListendState `json:"state"`
}

func (this *NotifyListenedState_Out) Process(session *net.Session) {
	g_OutHandler.NotifyListenedState(session, this)
}

func (this *NotifyListenedState_Out) TypeName() string {
	return "friend.notify_listened_state.out"
}

func (this *NotifyListenedState_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 14, 10
}

func (this *NotifyListenedState_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type CurrentPlatformFriendNum_In struct {
	Num int32 `json:"num"`
}

func (this *CurrentPlatformFriendNum_In) Process(session *net.Session) {
	g_InHandler.CurrentPlatformFriendNum(session, this)
}

func (this *CurrentPlatformFriendNum_In) TypeName() string {
	return "friend.current_platform_friend_num.in"
}

func (this *CurrentPlatformFriendNum_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 14, 11
}

type CurrentPlatformFriendNum_Out struct {
}

func (this *CurrentPlatformFriendNum_Out) Process(session *net.Session) {
	g_OutHandler.CurrentPlatformFriendNum(session, this)
}

func (this *CurrentPlatformFriendNum_Out) TypeName() string {
	return "friend.current_platform_friend_num.out"
}

func (this *CurrentPlatformFriendNum_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 14, 11
}

func (this *CurrentPlatformFriendNum_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetSendHeartList_In struct {
}

func (this *GetSendHeartList_In) Process(session *net.Session) {
	g_InHandler.GetSendHeartList(session, this)
}

func (this *GetSendHeartList_In) TypeName() string {
	return "friend.get_send_heart_list.in"
}

func (this *GetSendHeartList_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 14, 12
}

type GetSendHeartList_Out struct {
	Friends []GetSendHeartList_Out_Friends `json:"friends"`
}

type GetSendHeartList_Out_Friends struct {
	Pid           int64 `json:"pid"`
	SendHeartTime int64 `json:"send_heart_time"`
}

func (this *GetSendHeartList_Out) Process(session *net.Session) {
	g_OutHandler.GetSendHeartList(session, this)
}

func (this *GetSendHeartList_Out) TypeName() string {
	return "friend.get_send_heart_list.out"
}

func (this *GetSendHeartList_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 14, 12
}

func (this *GetSendHeartList_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type SendHeartToAllFriends_In struct {
	FriendType FriendType                         `json:"friend_type"`
	Friends    []SendHeartToAllFriends_In_Friends `json:"friends"`
}

type SendHeartToAllFriends_In_Friends struct {
	Nickname []byte `json:"nickname"`
	Pid      int64  `json:"pid"`
}

func (this *SendHeartToAllFriends_In) Process(session *net.Session) {
	g_InHandler.SendHeartToAllFriends(session, this)
}

func (this *SendHeartToAllFriends_In) TypeName() string {
	return "friend.send_heart_to_all_friends.in"
}

func (this *SendHeartToAllFriends_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 14, 13
}

type SendHeartToAllFriends_Out struct {
}

func (this *SendHeartToAllFriends_Out) Process(session *net.Session) {
	g_OutHandler.SendHeartToAllFriends(session, this)
}

func (this *SendHeartToAllFriends_Out) TypeName() string {
	return "friend.send_heart_to_all_friends.out"
}

func (this *SendHeartToAllFriends_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 14, 13
}

func (this *SendHeartToAllFriends_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

func (this *GetFriendList_In) Decode(buffer *net.Buffer) {
}

func (this *GetFriendList_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(14)
	buffer.WriteUint8(0)
}

func (this *GetFriendList_In) ByteSize() int {
	size := 2
	return size
}

func (this *GetFriendList_Out) Decode(buffer *net.Buffer) {
	this.CancelListenCount = int32(buffer.ReadUint32LE())
	this.PlatformFriendNum = int32(buffer.ReadUint32LE())
	this.Friends = make([]GetFriendList_Out_Friends, buffer.ReadUint8())
	for i := 0; i < len(this.Friends); i++ {
		this.Friends[i].Decode(buffer)
	}
}

func (this *GetFriendList_Out_Friends) Decode(buffer *net.Buffer) {
	this.Pid = int64(buffer.ReadUint64LE())
	this.RoleId = int8(buffer.ReadUint8())
	this.Nickname = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Level = int16(buffer.ReadUint16LE())
	this.FightNum = int32(buffer.ReadUint32LE())
	this.FriendMode = FriendMode(buffer.ReadUint8())
	this.BlockMode = int8(buffer.ReadUint8())
	this.LastChatTime = int64(buffer.ReadUint64LE())
	this.FriendTime = int64(buffer.ReadUint64LE())
}

func (this *GetFriendList_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(14)
	buffer.WriteUint8(0)
	buffer.WriteUint32LE(uint32(this.CancelListenCount))
	buffer.WriteUint32LE(uint32(this.PlatformFriendNum))
	buffer.WriteUint8(uint8(len(this.Friends)))
	for i := 0; i < len(this.Friends); i++ {
		this.Friends[i].Encode(buffer)
	}
}

func (this *GetFriendList_Out_Friends) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.Pid))
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint16LE(uint16(len(this.Nickname)))
	buffer.WriteBytes(this.Nickname)
	buffer.WriteUint16LE(uint16(this.Level))
	buffer.WriteUint32LE(uint32(this.FightNum))
	buffer.WriteUint8(uint8(this.FriendMode))
	buffer.WriteUint8(uint8(this.BlockMode))
	buffer.WriteUint64LE(uint64(this.LastChatTime))
	buffer.WriteUint64LE(uint64(this.FriendTime))
}

func (this *GetFriendList_Out) ByteSize() int {
	size := 11
	for i := 0; i < len(this.Friends); i++ {
		size += this.Friends[i].ByteSize()
	}
	return size
}

func (this *GetFriendList_Out_Friends) ByteSize() int {
	size := 35
	size += len(this.Nickname)
	return size
}

func (this *ListenByNick_In) Decode(buffer *net.Buffer) {
	this.Nick = buffer.ReadBytes(int(buffer.ReadUint16LE()))
}

func (this *ListenByNick_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(14)
	buffer.WriteUint8(1)
	buffer.WriteUint16LE(uint16(len(this.Nick)))
	buffer.WriteBytes(this.Nick)
}

func (this *ListenByNick_In) ByteSize() int {
	size := 4
	size += len(this.Nick)
	return size
}

func (this *ListenByNick_Out) Decode(buffer *net.Buffer) {
	this.Result = AddResult(buffer.ReadUint8())
	this.RoleId = int8(buffer.ReadUint8())
	this.Nickname = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Level = int16(buffer.ReadUint16LE())
	this.FightNum = int32(buffer.ReadUint32LE())
}

func (this *ListenByNick_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(14)
	buffer.WriteUint8(1)
	buffer.WriteUint8(uint8(this.Result))
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint16LE(uint16(len(this.Nickname)))
	buffer.WriteBytes(this.Nickname)
	buffer.WriteUint16LE(uint16(this.Level))
	buffer.WriteUint32LE(uint32(this.FightNum))
}

func (this *ListenByNick_Out) ByteSize() int {
	size := 12
	size += len(this.Nickname)
	return size
}

func (this *CancelListen_In) Decode(buffer *net.Buffer) {
	this.Pid = int64(buffer.ReadUint64LE())
}

func (this *CancelListen_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(14)
	buffer.WriteUint8(2)
	buffer.WriteUint64LE(uint64(this.Pid))
}

func (this *CancelListen_In) ByteSize() int {
	size := 10
	return size
}

func (this *CancelListen_Out) Decode(buffer *net.Buffer) {
	this.Result = buffer.ReadUint8() == 1
}

func (this *CancelListen_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(14)
	buffer.WriteUint8(2)
	if this.Result {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *CancelListen_Out) ByteSize() int {
	size := 3
	return size
}

func (this *SendHeart_In) Decode(buffer *net.Buffer) {
	this.Nickname = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.FriendType = FriendType(buffer.ReadUint8())
	this.Pid = int64(buffer.ReadUint64LE())
}

func (this *SendHeart_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(14)
	buffer.WriteUint8(3)
	buffer.WriteUint16LE(uint16(len(this.Nickname)))
	buffer.WriteBytes(this.Nickname)
	buffer.WriteUint8(uint8(this.FriendType))
	buffer.WriteUint64LE(uint64(this.Pid))
}

func (this *SendHeart_In) ByteSize() int {
	size := 13
	size += len(this.Nickname)
	return size
}

func (this *SendHeart_Out) Decode(buffer *net.Buffer) {
}

func (this *SendHeart_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(14)
	buffer.WriteUint8(3)
}

func (this *SendHeart_Out) ByteSize() int {
	size := 2
	return size
}

func (this *Chat_In) Decode(buffer *net.Buffer) {
	this.Pid = int64(buffer.ReadUint64LE())
	this.Message = buffer.ReadBytes(int(buffer.ReadUint16LE()))
}

func (this *Chat_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(14)
	buffer.WriteUint8(4)
	buffer.WriteUint64LE(uint64(this.Pid))
	buffer.WriteUint16LE(uint16(len(this.Message)))
	buffer.WriteBytes(this.Message)
}

func (this *Chat_In) ByteSize() int {
	size := 12
	size += len(this.Message)
	return size
}

func (this *Chat_Out) Decode(buffer *net.Buffer) {
	this.Banned = buffer.ReadUint8() == 1
}

func (this *Chat_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(14)
	buffer.WriteUint8(4)
	if this.Banned {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *Chat_Out) ByteSize() int {
	size := 3
	return size
}

func (this *GetChatHistory_In) Decode(buffer *net.Buffer) {
	this.Pid = int64(buffer.ReadUint64LE())
}

func (this *GetChatHistory_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(14)
	buffer.WriteUint8(5)
	buffer.WriteUint64LE(uint64(this.Pid))
}

func (this *GetChatHistory_In) ByteSize() int {
	size := 10
	return size
}

func (this *GetChatHistory_Out) Decode(buffer *net.Buffer) {
	this.Messages = make([]GetChatHistory_Out_Messages, buffer.ReadUint8())
	for i := 0; i < len(this.Messages); i++ {
		this.Messages[i].Decode(buffer)
	}
}

func (this *GetChatHistory_Out_Messages) Decode(buffer *net.Buffer) {
	this.Mode = MsgMode(buffer.ReadUint8())
	this.Id = int64(buffer.ReadUint64LE())
	this.SendTime = int64(buffer.ReadUint64LE())
	this.Message = buffer.ReadBytes(int(buffer.ReadUint16LE()))
}

func (this *GetChatHistory_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(14)
	buffer.WriteUint8(5)
	buffer.WriteUint8(uint8(len(this.Messages)))
	for i := 0; i < len(this.Messages); i++ {
		this.Messages[i].Encode(buffer)
	}
}

func (this *GetChatHistory_Out_Messages) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(uint8(this.Mode))
	buffer.WriteUint64LE(uint64(this.Id))
	buffer.WriteUint64LE(uint64(this.SendTime))
	buffer.WriteUint16LE(uint16(len(this.Message)))
	buffer.WriteBytes(this.Message)
}

func (this *GetChatHistory_Out) ByteSize() int {
	size := 3
	for i := 0; i < len(this.Messages); i++ {
		size += this.Messages[i].ByteSize()
	}
	return size
}

func (this *GetChatHistory_Out_Messages) ByteSize() int {
	size := 19
	size += len(this.Message)
	return size
}

func (this *GetOfflineMessages_Out) Decode(buffer *net.Buffer) {
	this.Chats = make([]GetOfflineMessages_Out_Chats, buffer.ReadUint8())
	for i := 0; i < len(this.Chats); i++ {
		this.Chats[i].Decode(buffer)
	}
	this.Listener = make([]GetOfflineMessages_Out_Listener, buffer.ReadUint8())
	for i := 0; i < len(this.Listener); i++ {
		this.Listener[i].Decode(buffer)
	}
}

func (this *GetOfflineMessages_Out_Chats) Decode(buffer *net.Buffer) {
	this.Pid = int64(buffer.ReadUint64LE())
	this.Nickname = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.RoleId = int8(buffer.ReadUint8())
	this.Level = int16(buffer.ReadUint16LE())
}

func (this *GetOfflineMessages_Out_Listener) Decode(buffer *net.Buffer) {
	this.Pid = int64(buffer.ReadUint64LE())
	this.Nick = buffer.ReadBytes(int(buffer.ReadUint16LE()))
}

func (this *GetOfflineMessages_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(14)
	buffer.WriteUint8(6)
	buffer.WriteUint8(uint8(len(this.Chats)))
	for i := 0; i < len(this.Chats); i++ {
		this.Chats[i].Encode(buffer)
	}
	buffer.WriteUint8(uint8(len(this.Listener)))
	for i := 0; i < len(this.Listener); i++ {
		this.Listener[i].Encode(buffer)
	}
}

func (this *GetOfflineMessages_Out_Chats) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.Pid))
	buffer.WriteUint16LE(uint16(len(this.Nickname)))
	buffer.WriteBytes(this.Nickname)
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint16LE(uint16(this.Level))
}

func (this *GetOfflineMessages_Out_Listener) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.Pid))
	buffer.WriteUint16LE(uint16(len(this.Nick)))
	buffer.WriteBytes(this.Nick)
}

func (this *GetOfflineMessages_Out) ByteSize() int {
	size := 4
	for i := 0; i < len(this.Chats); i++ {
		size += this.Chats[i].ByteSize()
	}
	for i := 0; i < len(this.Listener); i++ {
		size += this.Listener[i].ByteSize()
	}
	return size
}

func (this *GetOfflineMessages_Out_Chats) ByteSize() int {
	size := 13
	size += len(this.Nickname)
	return size
}

func (this *GetOfflineMessages_Out_Listener) ByteSize() int {
	size := 10
	size += len(this.Nick)
	return size
}

func (this *Block_In) Decode(buffer *net.Buffer) {
	this.Pid = int64(buffer.ReadUint64LE())
}

func (this *Block_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(14)
	buffer.WriteUint8(7)
	buffer.WriteUint64LE(uint64(this.Pid))
}

func (this *Block_In) ByteSize() int {
	size := 10
	return size
}

func (this *Block_Out) Decode(buffer *net.Buffer) {
}

func (this *Block_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(14)
	buffer.WriteUint8(7)
}

func (this *Block_Out) ByteSize() int {
	size := 2
	return size
}

func (this *CancelBlock_In) Decode(buffer *net.Buffer) {
	this.Pid = int64(buffer.ReadUint64LE())
}

func (this *CancelBlock_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(14)
	buffer.WriteUint8(8)
	buffer.WriteUint64LE(uint64(this.Pid))
}

func (this *CancelBlock_In) ByteSize() int {
	size := 10
	return size
}

func (this *CancelBlock_Out) Decode(buffer *net.Buffer) {
}

func (this *CancelBlock_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(14)
	buffer.WriteUint8(8)
}

func (this *CancelBlock_Out) ByteSize() int {
	size := 2
	return size
}

func (this *CleanChatHistory_In) Decode(buffer *net.Buffer) {
	this.Pid = int64(buffer.ReadUint64LE())
}

func (this *CleanChatHistory_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(14)
	buffer.WriteUint8(9)
	buffer.WriteUint64LE(uint64(this.Pid))
}

func (this *CleanChatHistory_In) ByteSize() int {
	size := 10
	return size
}

func (this *CleanChatHistory_Out) Decode(buffer *net.Buffer) {
}

func (this *CleanChatHistory_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(14)
	buffer.WriteUint8(9)
}

func (this *CleanChatHistory_Out) ByteSize() int {
	size := 2
	return size
}

func (this *NotifyListenedState_Out) Decode(buffer *net.Buffer) {
	this.Pid = int64(buffer.ReadUint64LE())
	this.Nick = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.State = ListendState(buffer.ReadUint8())
}

func (this *NotifyListenedState_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(14)
	buffer.WriteUint8(10)
	buffer.WriteUint64LE(uint64(this.Pid))
	buffer.WriteUint16LE(uint16(len(this.Nick)))
	buffer.WriteBytes(this.Nick)
	buffer.WriteUint8(uint8(this.State))
}

func (this *NotifyListenedState_Out) ByteSize() int {
	size := 13
	size += len(this.Nick)
	return size
}

func (this *CurrentPlatformFriendNum_In) Decode(buffer *net.Buffer) {
	this.Num = int32(buffer.ReadUint32LE())
}

func (this *CurrentPlatformFriendNum_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(14)
	buffer.WriteUint8(11)
	buffer.WriteUint32LE(uint32(this.Num))
}

func (this *CurrentPlatformFriendNum_In) ByteSize() int {
	size := 6
	return size
}

func (this *CurrentPlatformFriendNum_Out) Decode(buffer *net.Buffer) {
}

func (this *CurrentPlatformFriendNum_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(14)
	buffer.WriteUint8(11)
}

func (this *CurrentPlatformFriendNum_Out) ByteSize() int {
	size := 2
	return size
}

func (this *GetSendHeartList_In) Decode(buffer *net.Buffer) {
}

func (this *GetSendHeartList_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(14)
	buffer.WriteUint8(12)
}

func (this *GetSendHeartList_In) ByteSize() int {
	size := 2
	return size
}

func (this *GetSendHeartList_Out) Decode(buffer *net.Buffer) {
	this.Friends = make([]GetSendHeartList_Out_Friends, buffer.ReadUint16LE())
	for i := 0; i < len(this.Friends); i++ {
		this.Friends[i].Decode(buffer)
	}
}

func (this *GetSendHeartList_Out_Friends) Decode(buffer *net.Buffer) {
	this.Pid = int64(buffer.ReadUint64LE())
	this.SendHeartTime = int64(buffer.ReadUint64LE())
}

func (this *GetSendHeartList_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(14)
	buffer.WriteUint8(12)
	buffer.WriteUint16LE(uint16(len(this.Friends)))
	for i := 0; i < len(this.Friends); i++ {
		this.Friends[i].Encode(buffer)
	}
}

func (this *GetSendHeartList_Out_Friends) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.Pid))
	buffer.WriteUint64LE(uint64(this.SendHeartTime))
}

func (this *GetSendHeartList_Out) ByteSize() int {
	size := 4
	size += len(this.Friends) * 16
	return size
}

func (this *SendHeartToAllFriends_In) Decode(buffer *net.Buffer) {
	this.FriendType = FriendType(buffer.ReadUint8())
	this.Friends = make([]SendHeartToAllFriends_In_Friends, buffer.ReadUint16LE())
	for i := 0; i < len(this.Friends); i++ {
		this.Friends[i].Decode(buffer)
	}
}

func (this *SendHeartToAllFriends_In_Friends) Decode(buffer *net.Buffer) {
	this.Nickname = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Pid = int64(buffer.ReadUint64LE())
}

func (this *SendHeartToAllFriends_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(14)
	buffer.WriteUint8(13)
	buffer.WriteUint8(uint8(this.FriendType))
	buffer.WriteUint16LE(uint16(len(this.Friends)))
	for i := 0; i < len(this.Friends); i++ {
		this.Friends[i].Encode(buffer)
	}
}

func (this *SendHeartToAllFriends_In_Friends) Encode(buffer *net.Buffer) {
	buffer.WriteUint16LE(uint16(len(this.Nickname)))
	buffer.WriteBytes(this.Nickname)
	buffer.WriteUint64LE(uint64(this.Pid))
}

func (this *SendHeartToAllFriends_In) ByteSize() int {
	size := 5
	for i := 0; i < len(this.Friends); i++ {
		size += this.Friends[i].ByteSize()
	}
	return size
}

func (this *SendHeartToAllFriends_In_Friends) ByteSize() int {
	size := 10
	size += len(this.Nickname)
	return size
}

func (this *SendHeartToAllFriends_Out) Decode(buffer *net.Buffer) {
}

func (this *SendHeartToAllFriends_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(14)
	buffer.WriteUint8(13)
}

func (this *SendHeartToAllFriends_Out) ByteSize() int {
	size := 2
	return size
}
