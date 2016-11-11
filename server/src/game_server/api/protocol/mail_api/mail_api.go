package mail_api

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
	Read(*net.Session, *Read_In)
	TakeAttachment(*net.Session, *TakeAttachment_In)
	GetInfos(*net.Session, *GetInfos_In)
	RequestGlobalMail(*net.Session, *RequestGlobalMail_In)
}

type OutHandler interface {
	GetList(*net.Session, *GetList_Out)
	Read(*net.Session, *Read_Out)
	TakeAttachment(*net.Session, *TakeAttachment_Out)
	GetInfos(*net.Session, *GetInfos_Out)
	RequestGlobalMail(*net.Session, *RequestGlobalMail_Out)
}

func DecodeIn(message []byte) Request {
	var actionId = message[0]
	var buffer = net.NewBuffer(message[1:])
	switch actionId {
	case 0:
		request := new(GetList_In)
		request.Decode(buffer)
		return request
	case 1:
		request := new(Read_In)
		request.Decode(buffer)
		return request
	case 2:
		request := new(TakeAttachment_In)
		request.Decode(buffer)
		return request
	case 3:
		request := new(GetInfos_In)
		request.Decode(buffer)
		return request
	case 4:
		request := new(RequestGlobalMail_In)
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
	case 1:
		request := new(Read_Out)
		request.Decode(buffer)
		return request
	case 2:
		request := new(TakeAttachment_Out)
		request.Decode(buffer)
		return request
	case 3:
		request := new(GetInfos_Out)
		request.Decode(buffer)
		return request
	case 4:
		request := new(RequestGlobalMail_Out)
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
	return "mail.get_list.in"
}

func (this *GetList_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 12, 0
}

type GetList_Out struct {
	GetHeartNum int32               `json:"get_heart_num"`
	Mails       []GetList_Out_Mails `json:"mails"`
}

type GetList_Out_Mails struct {
	Id          int64                           `json:"id"`
	MailId      int32                           `json:"mail_id"`
	State       int8                            `json:"state"`
	Priority    int8                            `json:"priority"`
	SendTime    int64                           `json:"send_time"`
	ExpireTime  int64                           `json:"expire_time"`
	Parameters  []byte                          `json:"parameters"`
	Title       []byte                          `json:"title"`
	Content     []byte                          `json:"content"`
	Attachments []GetList_Out_Mails_Attachments `json:"attachments"`
}

type GetList_Out_Mails_Attachments struct {
	Id             int64 `json:"id"`
	AttachmentType int8  `json:"attachment_type"`
	ItemId         int16 `json:"item_id"`
	ItemNum        int64 `json:"item_num"`
}

func (this *GetList_Out) Process(session *net.Session) {
	g_OutHandler.GetList(session, this)
}

func (this *GetList_Out) TypeName() string {
	return "mail.get_list.out"
}

func (this *GetList_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 12, 0
}

func (this *GetList_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type Read_In struct {
	Id int64 `json:"id"`
}

func (this *Read_In) Process(session *net.Session) {
	g_InHandler.Read(session, this)
}

func (this *Read_In) TypeName() string {
	return "mail.read.in"
}

func (this *Read_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 12, 1
}

type Read_Out struct {
}

func (this *Read_Out) Process(session *net.Session) {
	g_OutHandler.Read(session, this)
}

func (this *Read_Out) TypeName() string {
	return "mail.read.out"
}

func (this *Read_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 12, 1
}

func (this *Read_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type TakeAttachment_In struct {
	AttachmentId int64 `json:"attachment_id"`
}

func (this *TakeAttachment_In) Process(session *net.Session) {
	g_InHandler.TakeAttachment(session, this)
}

func (this *TakeAttachment_In) TypeName() string {
	return "mail.take_attachment.in"
}

func (this *TakeAttachment_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 12, 2
}

type TakeAttachment_Out struct {
}

func (this *TakeAttachment_Out) Process(session *net.Session) {
	g_OutHandler.TakeAttachment(session, this)
}

func (this *TakeAttachment_Out) TypeName() string {
	return "mail.take_attachment.out"
}

func (this *TakeAttachment_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 12, 2
}

func (this *TakeAttachment_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetInfos_In struct {
}

func (this *GetInfos_In) Process(session *net.Session) {
	g_InHandler.GetInfos(session, this)
}

func (this *GetInfos_In) TypeName() string {
	return "mail.get_infos.in"
}

func (this *GetInfos_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 12, 3
}

type GetInfos_Out struct {
	NewMailNum int16 `json:"new_mail_num"`
	UnreadNum  int16 `json:"unread_num"`
	Total      int16 `json:"total"`
}

func (this *GetInfos_Out) Process(session *net.Session) {
	g_OutHandler.GetInfos(session, this)
}

func (this *GetInfos_Out) TypeName() string {
	return "mail.get_infos.out"
}

func (this *GetInfos_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 12, 3
}

func (this *GetInfos_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type RequestGlobalMail_In struct {
}

func (this *RequestGlobalMail_In) Process(session *net.Session) {
	g_InHandler.RequestGlobalMail(session, this)
}

func (this *RequestGlobalMail_In) TypeName() string {
	return "mail.request_global_mail.in"
}

func (this *RequestGlobalMail_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 12, 4
}

type RequestGlobalMail_Out struct {
}

func (this *RequestGlobalMail_Out) Process(session *net.Session) {
	g_OutHandler.RequestGlobalMail(session, this)
}

func (this *RequestGlobalMail_Out) TypeName() string {
	return "mail.request_global_mail.out"
}

func (this *RequestGlobalMail_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 12, 4
}

func (this *RequestGlobalMail_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

func (this *GetList_In) Decode(buffer *net.Buffer) {
}

func (this *GetList_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(12)
	buffer.WriteUint8(0)
}

func (this *GetList_In) ByteSize() int {
	size := 2
	return size
}

func (this *GetList_Out) Decode(buffer *net.Buffer) {
	this.GetHeartNum = int32(buffer.ReadUint32LE())
	this.Mails = make([]GetList_Out_Mails, buffer.ReadUint16LE())
	for i := 0; i < len(this.Mails); i++ {
		this.Mails[i].Decode(buffer)
	}
}

func (this *GetList_Out_Mails) Decode(buffer *net.Buffer) {
	this.Id = int64(buffer.ReadUint64LE())
	this.MailId = int32(buffer.ReadUint32LE())
	this.State = int8(buffer.ReadUint8())
	this.Priority = int8(buffer.ReadUint8())
	this.SendTime = int64(buffer.ReadUint64LE())
	this.ExpireTime = int64(buffer.ReadUint64LE())
	this.Parameters = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Title = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Content = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Attachments = make([]GetList_Out_Mails_Attachments, buffer.ReadUint8())
	for i := 0; i < len(this.Attachments); i++ {
		this.Attachments[i].Decode(buffer)
	}
}

func (this *GetList_Out_Mails_Attachments) Decode(buffer *net.Buffer) {
	this.Id = int64(buffer.ReadUint64LE())
	this.AttachmentType = int8(buffer.ReadUint8())
	this.ItemId = int16(buffer.ReadUint16LE())
	this.ItemNum = int64(buffer.ReadUint64LE())
}

func (this *GetList_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(12)
	buffer.WriteUint8(0)
	buffer.WriteUint32LE(uint32(this.GetHeartNum))
	buffer.WriteUint16LE(uint16(len(this.Mails)))
	for i := 0; i < len(this.Mails); i++ {
		this.Mails[i].Encode(buffer)
	}
}

func (this *GetList_Out_Mails) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.Id))
	buffer.WriteUint32LE(uint32(this.MailId))
	buffer.WriteUint8(uint8(this.State))
	buffer.WriteUint8(uint8(this.Priority))
	buffer.WriteUint64LE(uint64(this.SendTime))
	buffer.WriteUint64LE(uint64(this.ExpireTime))
	buffer.WriteUint16LE(uint16(len(this.Parameters)))
	buffer.WriteBytes(this.Parameters)
	buffer.WriteUint16LE(uint16(len(this.Title)))
	buffer.WriteBytes(this.Title)
	buffer.WriteUint16LE(uint16(len(this.Content)))
	buffer.WriteBytes(this.Content)
	buffer.WriteUint8(uint8(len(this.Attachments)))
	for i := 0; i < len(this.Attachments); i++ {
		this.Attachments[i].Encode(buffer)
	}
}

func (this *GetList_Out_Mails_Attachments) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.Id))
	buffer.WriteUint8(uint8(this.AttachmentType))
	buffer.WriteUint16LE(uint16(this.ItemId))
	buffer.WriteUint64LE(uint64(this.ItemNum))
}

func (this *GetList_Out) ByteSize() int {
	size := 8
	for i := 0; i < len(this.Mails); i++ {
		size += this.Mails[i].ByteSize()
	}
	return size
}

func (this *GetList_Out_Mails) ByteSize() int {
	size := 37
	size += len(this.Parameters)
	size += len(this.Title)
	size += len(this.Content)
	size += len(this.Attachments) * 19
	return size
}

func (this *Read_In) Decode(buffer *net.Buffer) {
	this.Id = int64(buffer.ReadUint64LE())
}

func (this *Read_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(12)
	buffer.WriteUint8(1)
	buffer.WriteUint64LE(uint64(this.Id))
}

func (this *Read_In) ByteSize() int {
	size := 10
	return size
}

func (this *Read_Out) Decode(buffer *net.Buffer) {
}

func (this *Read_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(12)
	buffer.WriteUint8(1)
}

func (this *Read_Out) ByteSize() int {
	size := 2
	return size
}

func (this *TakeAttachment_In) Decode(buffer *net.Buffer) {
	this.AttachmentId = int64(buffer.ReadUint64LE())
}

func (this *TakeAttachment_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(12)
	buffer.WriteUint8(2)
	buffer.WriteUint64LE(uint64(this.AttachmentId))
}

func (this *TakeAttachment_In) ByteSize() int {
	size := 10
	return size
}

func (this *TakeAttachment_Out) Decode(buffer *net.Buffer) {
}

func (this *TakeAttachment_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(12)
	buffer.WriteUint8(2)
}

func (this *TakeAttachment_Out) ByteSize() int {
	size := 2
	return size
}

func (this *GetInfos_In) Decode(buffer *net.Buffer) {
}

func (this *GetInfos_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(12)
	buffer.WriteUint8(3)
}

func (this *GetInfos_In) ByteSize() int {
	size := 2
	return size
}

func (this *GetInfos_Out) Decode(buffer *net.Buffer) {
	this.NewMailNum = int16(buffer.ReadUint16LE())
	this.UnreadNum = int16(buffer.ReadUint16LE())
	this.Total = int16(buffer.ReadUint16LE())
}

func (this *GetInfos_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(12)
	buffer.WriteUint8(3)
	buffer.WriteUint16LE(uint16(this.NewMailNum))
	buffer.WriteUint16LE(uint16(this.UnreadNum))
	buffer.WriteUint16LE(uint16(this.Total))
}

func (this *GetInfos_Out) ByteSize() int {
	size := 8
	return size
}

func (this *RequestGlobalMail_In) Decode(buffer *net.Buffer) {
}

func (this *RequestGlobalMail_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(12)
	buffer.WriteUint8(4)
}

func (this *RequestGlobalMail_In) ByteSize() int {
	size := 2
	return size
}

func (this *RequestGlobalMail_Out) Decode(buffer *net.Buffer) {
}

func (this *RequestGlobalMail_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(12)
	buffer.WriteUint8(4)
}

func (this *RequestGlobalMail_Out) ByteSize() int {
	size := 2
	return size
}
