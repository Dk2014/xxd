package channel_api

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
	GetLatestWorldChannelMessage(*net.Session, *GetLatestWorldChannelMessage_In)
	AddWorldChat(*net.Session, *AddWorldChat_In)
	WorldChannelInfo(*net.Session, *WorldChannelInfo_In)
	AddCliqueChat(*net.Session, *AddCliqueChat_In)
	GetLatestCliqueMessages(*net.Session, *GetLatestCliqueMessages_In)
}

type OutHandler interface {
	GetLatestWorldChannelMessage(*net.Session, *GetLatestWorldChannelMessage_Out)
	AddWorldChat(*net.Session, *AddWorldChat_Out)
	WorldChannelInfo(*net.Session, *WorldChannelInfo_Out)
	SendGlobalMessages(*net.Session, *SendGlobalMessages_Out)
	AddCliqueChat(*net.Session, *AddCliqueChat_Out)
	GetLatestCliqueMessages(*net.Session, *GetLatestCliqueMessages_Out)
	SendCliqueMessage(*net.Session, *SendCliqueMessage_Out)
}

func DecodeIn(message []byte) Request {
	var actionId = message[0]
	var buffer = net.NewBuffer(message[1:])
	switch actionId {
	case 0:
		request := new(GetLatestWorldChannelMessage_In)
		request.Decode(buffer)
		return request
	case 1:
		request := new(AddWorldChat_In)
		request.Decode(buffer)
		return request
	case 2:
		request := new(WorldChannelInfo_In)
		request.Decode(buffer)
		return request
	case 4:
		request := new(AddCliqueChat_In)
		request.Decode(buffer)
		return request
	case 5:
		request := new(GetLatestCliqueMessages_In)
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
		request := new(GetLatestWorldChannelMessage_Out)
		request.Decode(buffer)
		return request
	case 1:
		request := new(AddWorldChat_Out)
		request.Decode(buffer)
		return request
	case 2:
		request := new(WorldChannelInfo_Out)
		request.Decode(buffer)
		return request
	case 3:
		request := new(SendGlobalMessages_Out)
		request.Decode(buffer)
		return request
	case 4:
		request := new(AddCliqueChat_Out)
		request.Decode(buffer)
		return request
	case 5:
		request := new(GetLatestCliqueMessages_Out)
		request.Decode(buffer)
		return request
	case 6:
		request := new(SendCliqueMessage_Out)
		request.Decode(buffer)
		return request
	}
	_ = buffer
	panic("unsupported response")
}

type MessageType int8

const (
	MESSAGE_TYPE_CHAT                     MessageType = 0
	MESSAGE_TYPE_RARE_PROPS               MessageType = 1
	MESSAGE_TYPE_CLIQUE_MESSAGE           MessageType = 2
	MESSAGE_TYPE_CLIQUE_CHAT              MessageType = 3
	MESSAGE_TYPE_CLIQUE_NEWS              MessageType = 4
	MESSAGE_TYPE_DESPAIR_LAND_BOSS_KILLED MessageType = 5
)

type CliqueMessage struct {
	TplId      int16       `json:"tpl_id"`
	Pid        int64       `json:"pid"`
	MsgType    MessageType `json:"msg_type"`
	Nickname   []byte      `json:"nickname"`
	Timestamp  int64       `json:"timestamp"`
	Parameters []byte      `json:"parameters"`
}

func (this *CliqueMessage) Decode(buffer *net.Buffer) {
	this.TplId = int16(buffer.ReadUint16LE())
	this.Pid = int64(buffer.ReadUint64LE())
	this.MsgType = MessageType(buffer.ReadUint8())
	this.Nickname = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Timestamp = int64(buffer.ReadUint64LE())
	this.Parameters = buffer.ReadBytes(int(buffer.ReadUint16LE()))
}

func (this *CliqueMessage) Encode(buffer *net.Buffer) {
	buffer.WriteUint16LE(uint16(this.TplId))
	buffer.WriteUint64LE(uint64(this.Pid))
	buffer.WriteUint8(uint8(this.MsgType))
	buffer.WriteUint16LE(uint16(len(this.Nickname)))
	buffer.WriteBytes(this.Nickname)
	buffer.WriteUint64LE(uint64(this.Timestamp))
	buffer.WriteUint16LE(uint16(len(this.Parameters)))
	buffer.WriteBytes(this.Parameters)
}

func (this *CliqueMessage) ByteSize() int {
	size := 23
	size += len(this.Nickname)
	size += len(this.Parameters)
	return size
}

type GetLatestWorldChannelMessage_In struct {
}

func (this *GetLatestWorldChannelMessage_In) Process(session *net.Session) {
	g_InHandler.GetLatestWorldChannelMessage(session, this)
}

func (this *GetLatestWorldChannelMessage_In) TypeName() string {
	return "channel.get_latest_world_channel_message.in"
}

func (this *GetLatestWorldChannelMessage_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 29, 0
}

type GetLatestWorldChannelMessage_Out struct {
	Messages []GetLatestWorldChannelMessage_Out_Messages `json:"messages"`
}

type GetLatestWorldChannelMessage_Out_Messages struct {
	Pid        int64       `json:"pid"`
	MsgType    MessageType `json:"msg_type"`
	Nickname   []byte      `json:"nickname"`
	Timestamp  int64       `json:"timestamp"`
	Parameters []byte      `json:"parameters"`
	TplId      int16       `json:"tpl_id"`
}

func (this *GetLatestWorldChannelMessage_Out) Process(session *net.Session) {
	g_OutHandler.GetLatestWorldChannelMessage(session, this)
}

func (this *GetLatestWorldChannelMessage_Out) TypeName() string {
	return "channel.get_latest_world_channel_message.out"
}

func (this *GetLatestWorldChannelMessage_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 29, 0
}

func (this *GetLatestWorldChannelMessage_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type AddWorldChat_In struct {
	Content []byte `json:"content"`
}

func (this *AddWorldChat_In) Process(session *net.Session) {
	g_InHandler.AddWorldChat(session, this)
}

func (this *AddWorldChat_In) TypeName() string {
	return "channel.add_world_chat.in"
}

func (this *AddWorldChat_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 29, 1
}

type AddWorldChat_Out struct {
	Banned bool `json:"banned"`
}

func (this *AddWorldChat_Out) Process(session *net.Session) {
	g_OutHandler.AddWorldChat(session, this)
}

func (this *AddWorldChat_Out) TypeName() string {
	return "channel.add_world_chat.out"
}

func (this *AddWorldChat_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 29, 1
}

func (this *AddWorldChat_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type WorldChannelInfo_In struct {
}

func (this *WorldChannelInfo_In) Process(session *net.Session) {
	g_InHandler.WorldChannelInfo(session, this)
}

func (this *WorldChannelInfo_In) TypeName() string {
	return "channel.world_channel_info.in"
}

func (this *WorldChannelInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 29, 2
}

type WorldChannelInfo_Out struct {
	Timestamp int64 `json:"timestamp"`
	DailyNum  int16 `json:"daily_num"`
}

func (this *WorldChannelInfo_Out) Process(session *net.Session) {
	g_OutHandler.WorldChannelInfo(session, this)
}

func (this *WorldChannelInfo_Out) TypeName() string {
	return "channel.world_channel_info.out"
}

func (this *WorldChannelInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 29, 2
}

func (this *WorldChannelInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type SendGlobalMessages_Out struct {
	Messages []SendGlobalMessages_Out_Messages `json:"messages"`
}

type SendGlobalMessages_Out_Messages struct {
	Pid        int64       `json:"pid"`
	MsgType    MessageType `json:"msg_type"`
	Nickname   []byte      `json:"nickname"`
	Timestamp  int64       `json:"timestamp"`
	Parameters []byte      `json:"parameters"`
	TplId      int16       `json:"tpl_id"`
}

func (this *SendGlobalMessages_Out) Process(session *net.Session) {
	g_OutHandler.SendGlobalMessages(session, this)
}

func (this *SendGlobalMessages_Out) TypeName() string {
	return "channel.send_global_messages.out"
}

func (this *SendGlobalMessages_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 29, 3
}

func (this *SendGlobalMessages_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type AddCliqueChat_In struct {
	Content []byte `json:"content"`
}

func (this *AddCliqueChat_In) Process(session *net.Session) {
	g_InHandler.AddCliqueChat(session, this)
}

func (this *AddCliqueChat_In) TypeName() string {
	return "channel.add_clique_chat.in"
}

func (this *AddCliqueChat_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 29, 4
}

type AddCliqueChat_Out struct {
	Banned bool `json:"banned"`
}

func (this *AddCliqueChat_Out) Process(session *net.Session) {
	g_OutHandler.AddCliqueChat(session, this)
}

func (this *AddCliqueChat_Out) TypeName() string {
	return "channel.add_clique_chat.out"
}

func (this *AddCliqueChat_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 29, 4
}

func (this *AddCliqueChat_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetLatestCliqueMessages_In struct {
}

func (this *GetLatestCliqueMessages_In) Process(session *net.Session) {
	g_InHandler.GetLatestCliqueMessages(session, this)
}

func (this *GetLatestCliqueMessages_In) TypeName() string {
	return "channel.get_latest_clique_messages.in"
}

func (this *GetLatestCliqueMessages_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 29, 5
}

type GetLatestCliqueMessages_Out struct {
	Messages []GetLatestCliqueMessages_Out_Messages `json:"messages"`
}

type GetLatestCliqueMessages_Out_Messages struct {
	Message CliqueMessage `json:"message"`
}

func (this *GetLatestCliqueMessages_Out) Process(session *net.Session) {
	g_OutHandler.GetLatestCliqueMessages(session, this)
}

func (this *GetLatestCliqueMessages_Out) TypeName() string {
	return "channel.get_latest_clique_messages.out"
}

func (this *GetLatestCliqueMessages_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 29, 5
}

func (this *GetLatestCliqueMessages_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type SendCliqueMessage_Out struct {
	Message CliqueMessage `json:"message"`
}

func (this *SendCliqueMessage_Out) Process(session *net.Session) {
	g_OutHandler.SendCliqueMessage(session, this)
}

func (this *SendCliqueMessage_Out) TypeName() string {
	return "channel.send_clique_message.out"
}

func (this *SendCliqueMessage_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 29, 6
}

func (this *SendCliqueMessage_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

func (this *GetLatestWorldChannelMessage_In) Decode(buffer *net.Buffer) {
}

func (this *GetLatestWorldChannelMessage_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(29)
	buffer.WriteUint8(0)
}

func (this *GetLatestWorldChannelMessage_In) ByteSize() int {
	size := 2
	return size
}

func (this *GetLatestWorldChannelMessage_Out) Decode(buffer *net.Buffer) {
	this.Messages = make([]GetLatestWorldChannelMessage_Out_Messages, buffer.ReadUint8())
	for i := 0; i < len(this.Messages); i++ {
		this.Messages[i].Decode(buffer)
	}
}

func (this *GetLatestWorldChannelMessage_Out_Messages) Decode(buffer *net.Buffer) {
	this.Pid = int64(buffer.ReadUint64LE())
	this.MsgType = MessageType(buffer.ReadUint8())
	this.Nickname = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Timestamp = int64(buffer.ReadUint64LE())
	this.Parameters = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.TplId = int16(buffer.ReadUint16LE())
}

func (this *GetLatestWorldChannelMessage_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(29)
	buffer.WriteUint8(0)
	buffer.WriteUint8(uint8(len(this.Messages)))
	for i := 0; i < len(this.Messages); i++ {
		this.Messages[i].Encode(buffer)
	}
}

func (this *GetLatestWorldChannelMessage_Out_Messages) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.Pid))
	buffer.WriteUint8(uint8(this.MsgType))
	buffer.WriteUint16LE(uint16(len(this.Nickname)))
	buffer.WriteBytes(this.Nickname)
	buffer.WriteUint64LE(uint64(this.Timestamp))
	buffer.WriteUint16LE(uint16(len(this.Parameters)))
	buffer.WriteBytes(this.Parameters)
	buffer.WriteUint16LE(uint16(this.TplId))
}

func (this *GetLatestWorldChannelMessage_Out) ByteSize() int {
	size := 3
	for i := 0; i < len(this.Messages); i++ {
		size += this.Messages[i].ByteSize()
	}
	return size
}

func (this *GetLatestWorldChannelMessage_Out_Messages) ByteSize() int {
	size := 23
	size += len(this.Nickname)
	size += len(this.Parameters)
	return size
}

func (this *AddWorldChat_In) Decode(buffer *net.Buffer) {
	this.Content = buffer.ReadBytes(int(buffer.ReadUint16LE()))
}

func (this *AddWorldChat_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(29)
	buffer.WriteUint8(1)
	buffer.WriteUint16LE(uint16(len(this.Content)))
	buffer.WriteBytes(this.Content)
}

func (this *AddWorldChat_In) ByteSize() int {
	size := 4
	size += len(this.Content)
	return size
}

func (this *AddWorldChat_Out) Decode(buffer *net.Buffer) {
	this.Banned = buffer.ReadUint8() == 1
}

func (this *AddWorldChat_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(29)
	buffer.WriteUint8(1)
	if this.Banned {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *AddWorldChat_Out) ByteSize() int {
	size := 3
	return size
}

func (this *WorldChannelInfo_In) Decode(buffer *net.Buffer) {
}

func (this *WorldChannelInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(29)
	buffer.WriteUint8(2)
}

func (this *WorldChannelInfo_In) ByteSize() int {
	size := 2
	return size
}

func (this *WorldChannelInfo_Out) Decode(buffer *net.Buffer) {
	this.Timestamp = int64(buffer.ReadUint64LE())
	this.DailyNum = int16(buffer.ReadUint16LE())
}

func (this *WorldChannelInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(29)
	buffer.WriteUint8(2)
	buffer.WriteUint64LE(uint64(this.Timestamp))
	buffer.WriteUint16LE(uint16(this.DailyNum))
}

func (this *WorldChannelInfo_Out) ByteSize() int {
	size := 12
	return size
}

func (this *SendGlobalMessages_Out) Decode(buffer *net.Buffer) {
	this.Messages = make([]SendGlobalMessages_Out_Messages, buffer.ReadUint8())
	for i := 0; i < len(this.Messages); i++ {
		this.Messages[i].Decode(buffer)
	}
}

func (this *SendGlobalMessages_Out_Messages) Decode(buffer *net.Buffer) {
	this.Pid = int64(buffer.ReadUint64LE())
	this.MsgType = MessageType(buffer.ReadUint8())
	this.Nickname = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Timestamp = int64(buffer.ReadUint64LE())
	this.Parameters = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.TplId = int16(buffer.ReadUint16LE())
}

func (this *SendGlobalMessages_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(29)
	buffer.WriteUint8(3)
	buffer.WriteUint8(uint8(len(this.Messages)))
	for i := 0; i < len(this.Messages); i++ {
		this.Messages[i].Encode(buffer)
	}
}

func (this *SendGlobalMessages_Out_Messages) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.Pid))
	buffer.WriteUint8(uint8(this.MsgType))
	buffer.WriteUint16LE(uint16(len(this.Nickname)))
	buffer.WriteBytes(this.Nickname)
	buffer.WriteUint64LE(uint64(this.Timestamp))
	buffer.WriteUint16LE(uint16(len(this.Parameters)))
	buffer.WriteBytes(this.Parameters)
	buffer.WriteUint16LE(uint16(this.TplId))
}

func (this *SendGlobalMessages_Out) ByteSize() int {
	size := 3
	for i := 0; i < len(this.Messages); i++ {
		size += this.Messages[i].ByteSize()
	}
	return size
}

func (this *SendGlobalMessages_Out_Messages) ByteSize() int {
	size := 23
	size += len(this.Nickname)
	size += len(this.Parameters)
	return size
}

func (this *AddCliqueChat_In) Decode(buffer *net.Buffer) {
	this.Content = buffer.ReadBytes(int(buffer.ReadUint16LE()))
}

func (this *AddCliqueChat_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(29)
	buffer.WriteUint8(4)
	buffer.WriteUint16LE(uint16(len(this.Content)))
	buffer.WriteBytes(this.Content)
}

func (this *AddCliqueChat_In) ByteSize() int {
	size := 4
	size += len(this.Content)
	return size
}

func (this *AddCliqueChat_Out) Decode(buffer *net.Buffer) {
	this.Banned = buffer.ReadUint8() == 1
}

func (this *AddCliqueChat_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(29)
	buffer.WriteUint8(4)
	if this.Banned {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *AddCliqueChat_Out) ByteSize() int {
	size := 3
	return size
}

func (this *GetLatestCliqueMessages_In) Decode(buffer *net.Buffer) {
}

func (this *GetLatestCliqueMessages_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(29)
	buffer.WriteUint8(5)
}

func (this *GetLatestCliqueMessages_In) ByteSize() int {
	size := 2
	return size
}

func (this *GetLatestCliqueMessages_Out) Decode(buffer *net.Buffer) {
	this.Messages = make([]GetLatestCliqueMessages_Out_Messages, buffer.ReadUint8())
	for i := 0; i < len(this.Messages); i++ {
		this.Messages[i].Decode(buffer)
	}
}

func (this *GetLatestCliqueMessages_Out_Messages) Decode(buffer *net.Buffer) {
	this.Message.Decode(buffer)
}

func (this *GetLatestCliqueMessages_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(29)
	buffer.WriteUint8(5)
	buffer.WriteUint8(uint8(len(this.Messages)))
	for i := 0; i < len(this.Messages); i++ {
		this.Messages[i].Encode(buffer)
	}
}

func (this *GetLatestCliqueMessages_Out_Messages) Encode(buffer *net.Buffer) {
	this.Message.Encode(buffer)
}

func (this *GetLatestCliqueMessages_Out) ByteSize() int {
	size := 3
	for i := 0; i < len(this.Messages); i++ {
		size += this.Messages[i].ByteSize()
	}
	return size
}

func (this *GetLatestCliqueMessages_Out_Messages) ByteSize() int {
	size := 0
	size += this.Message.ByteSize()
	return size
}

func (this *SendCliqueMessage_Out) Decode(buffer *net.Buffer) {
	this.Message.Decode(buffer)
}

func (this *SendCliqueMessage_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(29)
	buffer.WriteUint8(6)
	this.Message.Encode(buffer)
}

func (this *SendCliqueMessage_Out) ByteSize() int {
	size := 2
	size += this.Message.ByteSize()
	return size
}
