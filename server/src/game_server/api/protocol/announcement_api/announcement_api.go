package announcement_api

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
	GetList(*net.Session, *GetList_In)
}

type OutHandler interface {
	GetList(*net.Session, *GetList_Out)
}

func DecodeIn(message []byte) Request {
	var actionId = message[0]
	var buffer = net.NewBuffer(message[1:])
	switch actionId {
	case 0:
		request := new(GetList_In)
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
		request := new(GetList_Out)
		request.Decode(buffer)
		return request
	}
	_ = buffer
	panic("unsupported response")
}

type GetList_In struct {
}

func (this *GetList_In) Process(session *net.Session) {
	g_InHandler.GetList(session, this)
}

func (this *GetList_In) TypeName() string {
	return "announcement.get_list.in"
}

func (this *GetList_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 18, 0
}

type GetList_Out struct {
	Announcements []GetList_Out_Announcements `json:"announcements"`
}

type GetList_Out_Announcements struct {
	Id          int64  `json:"id"`
	TplId       int32  `json:"tpl_id"`
	ExpireTime  int64  `json:"expire_time"`
	Parameters  []byte `json:"parameters"`
	Content     []byte `json:"content"`
	SpacingTime int32  `json:"spacing_time"`
}

func (this *GetList_Out) Process(session *net.Session) {
	g_OutHandler.GetList(session, this)
}

func (this *GetList_Out) TypeName() string {
	return "announcement.get_list.out"
}

func (this *GetList_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 18, 0
}

func (this *GetList_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

func (this *GetList_In) Decode(buffer *net.Buffer) {
}

func (this *GetList_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(18)
	buffer.WriteUint8(0)
}

func (this *GetList_In) ByteSize() int {
	size := 2
	return size
}

func (this *GetList_Out) Decode(buffer *net.Buffer) {
	this.Announcements = make([]GetList_Out_Announcements, buffer.ReadUint8())
	for i := 0; i < len(this.Announcements); i++ {
		this.Announcements[i].Decode(buffer)
	}
}

func (this *GetList_Out_Announcements) Decode(buffer *net.Buffer) {
	this.Id = int64(buffer.ReadUint64LE())
	this.TplId = int32(buffer.ReadUint32LE())
	this.ExpireTime = int64(buffer.ReadUint64LE())
	this.Parameters = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Content = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.SpacingTime = int32(buffer.ReadUint32LE())
}

func (this *GetList_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(18)
	buffer.WriteUint8(0)
	buffer.WriteUint8(uint8(len(this.Announcements)))
	for i := 0; i < len(this.Announcements); i++ {
		this.Announcements[i].Encode(buffer)
	}
}

func (this *GetList_Out_Announcements) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.Id))
	buffer.WriteUint32LE(uint32(this.TplId))
	buffer.WriteUint64LE(uint64(this.ExpireTime))
	buffer.WriteUint16LE(uint16(len(this.Parameters)))
	buffer.WriteBytes(this.Parameters)
	buffer.WriteUint16LE(uint16(len(this.Content)))
	buffer.WriteBytes(this.Content)
	buffer.WriteUint32LE(uint32(this.SpacingTime))
}

func (this *GetList_Out) ByteSize() int {
	size := 3
	for i := 0; i < len(this.Announcements); i++ {
		size += this.Announcements[i].ByteSize()
	}
	return size
}

func (this *GetList_Out_Announcements) ByteSize() int {
	size := 28
	size += len(this.Parameters)
	size += len(this.Content)
	return size
}
