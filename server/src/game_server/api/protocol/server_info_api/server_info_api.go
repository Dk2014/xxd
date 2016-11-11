package server_info_api

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
	GetVersion(*net.Session, *GetVersion_In)
	GetApiCount(*net.Session, *GetApiCount_In)
	SearchPlayerRole(*net.Session, *SearchPlayerRole_In)
	UpdateAccessToken(*net.Session, *UpdateAccessToken_In)
	UpdateEventData(*net.Session, *UpdateEventData_In)
	TssData(*net.Session, *TssData_In)
}

type OutHandler interface {
	GetVersion(*net.Session, *GetVersion_Out)
	GetApiCount(*net.Session, *GetApiCount_Out)
	SearchPlayerRole(*net.Session, *SearchPlayerRole_Out)
	UpdateAccessToken(*net.Session, *UpdateAccessToken_Out)
	UpdateEventData(*net.Session, *UpdateEventData_Out)
	TssData(*net.Session, *TssData_Out)
}

func DecodeIn(message []byte) Request {
	var actionId = message[0]
	var buffer = net.NewBuffer(message[1:])
	switch actionId {
	case 0:
		request := new(GetVersion_In)
		request.Decode(buffer)
		return request
	case 1:
		request := new(GetApiCount_In)
		request.Decode(buffer)
		return request
	case 2:
		request := new(SearchPlayerRole_In)
		request.Decode(buffer)
		return request
	case 3:
		request := new(UpdateAccessToken_In)
		request.Decode(buffer)
		return request
	case 4:
		request := new(UpdateEventData_In)
		request.Decode(buffer)
		return request
	case 5:
		request := new(TssData_In)
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
		request := new(GetVersion_Out)
		request.Decode(buffer)
		return request
	case 1:
		request := new(GetApiCount_Out)
		request.Decode(buffer)
		return request
	case 2:
		request := new(SearchPlayerRole_Out)
		request.Decode(buffer)
		return request
	case 3:
		request := new(UpdateAccessToken_Out)
		request.Decode(buffer)
		return request
	case 4:
		request := new(UpdateEventData_Out)
		request.Decode(buffer)
		return request
	case 5:
		request := new(TssData_Out)
		request.Decode(buffer)
		return request
	}
	_ = buffer
	panic("unsupported response")
}

type GetVersion_In struct {
}

func (this *GetVersion_In) Process(session *net.Session) {
	g_InHandler.GetVersion(session, this)
}

func (this *GetVersion_In) TypeName() string {
	return "server_info.get_version.in"
}

func (this *GetVersion_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 88, 0
}

type GetVersion_Out struct {
	Version []byte `json:"version"`
}

func (this *GetVersion_Out) Process(session *net.Session) {
	g_OutHandler.GetVersion(session, this)
}

func (this *GetVersion_Out) TypeName() string {
	return "server_info.get_version.out"
}

func (this *GetVersion_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 88, 0
}

func (this *GetVersion_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetApiCount_In struct {
}

func (this *GetApiCount_In) Process(session *net.Session) {
	g_InHandler.GetApiCount(session, this)
}

func (this *GetApiCount_In) TypeName() string {
	return "server_info.get_api_count.in"
}

func (this *GetApiCount_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 88, 1
}

type GetApiCount_Out struct {
	CountIn  int64 `json:"count_in"`
	CountOut int64 `json:"count_out"`
}

func (this *GetApiCount_Out) Process(session *net.Session) {
	g_OutHandler.GetApiCount(session, this)
}

func (this *GetApiCount_Out) TypeName() string {
	return "server_info.get_api_count.out"
}

func (this *GetApiCount_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 88, 1
}

func (this *GetApiCount_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type SearchPlayerRole_In struct {
	Openid []byte `json:"openid"`
}

func (this *SearchPlayerRole_In) Process(session *net.Session) {
	g_InHandler.SearchPlayerRole(session, this)
}

func (this *SearchPlayerRole_In) TypeName() string {
	return "server_info.search_player_role.in"
}

func (this *SearchPlayerRole_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 88, 2
}

type SearchPlayerRole_Out struct {
	Result bool `json:"result"`
}

func (this *SearchPlayerRole_Out) Process(session *net.Session) {
	g_OutHandler.SearchPlayerRole(session, this)
}

func (this *SearchPlayerRole_Out) TypeName() string {
	return "server_info.search_player_role.out"
}

func (this *SearchPlayerRole_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 88, 2
}

func (this *SearchPlayerRole_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type UpdateAccessToken_In struct {
	Token []byte `json:"token"`
	Pfkey []byte `json:"pfkey"`
}

func (this *UpdateAccessToken_In) Process(session *net.Session) {
	g_InHandler.UpdateAccessToken(session, this)
}

func (this *UpdateAccessToken_In) TypeName() string {
	return "server_info.update_access_token.in"
}

func (this *UpdateAccessToken_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 88, 3
}

type UpdateAccessToken_Out struct {
}

func (this *UpdateAccessToken_Out) Process(session *net.Session) {
	g_OutHandler.UpdateAccessToken(session, this)
}

func (this *UpdateAccessToken_Out) TypeName() string {
	return "server_info.update_access_token.out"
}

func (this *UpdateAccessToken_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 88, 3
}

func (this *UpdateAccessToken_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type UpdateEventData_In struct {
	Version int32 `json:"version"`
}

func (this *UpdateEventData_In) Process(session *net.Session) {
	g_InHandler.UpdateEventData(session, this)
}

func (this *UpdateEventData_In) TypeName() string {
	return "server_info.update_event_data.in"
}

func (this *UpdateEventData_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 88, 4
}

type UpdateEventData_Out struct {
	Json []byte `json:"json"`
}

func (this *UpdateEventData_Out) Process(session *net.Session) {
	g_OutHandler.UpdateEventData(session, this)
}

func (this *UpdateEventData_Out) TypeName() string {
	return "server_info.update_event_data.out"
}

func (this *UpdateEventData_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 88, 4
}

func (this *UpdateEventData_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type TssData_In struct {
	Data []byte `json:"data"`
}

func (this *TssData_In) Process(session *net.Session) {
	g_InHandler.TssData(session, this)
}

func (this *TssData_In) TypeName() string {
	return "server_info.tss_data.in"
}

func (this *TssData_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 88, 5
}

type TssData_Out struct {
	Data []byte `json:"data"`
}

func (this *TssData_Out) Process(session *net.Session) {
	g_OutHandler.TssData(session, this)
}

func (this *TssData_Out) TypeName() string {
	return "server_info.tss_data.out"
}

func (this *TssData_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 88, 5
}

func (this *TssData_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

func (this *GetVersion_In) Decode(buffer *net.Buffer) {
}

func (this *GetVersion_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(88)
	buffer.WriteUint8(0)
}

func (this *GetVersion_In) ByteSize() int {
	size := 2
	return size
}

func (this *GetVersion_Out) Decode(buffer *net.Buffer) {
	this.Version = buffer.ReadBytes(int(buffer.ReadUint16LE()))
}

func (this *GetVersion_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(88)
	buffer.WriteUint8(0)
	buffer.WriteUint16LE(uint16(len(this.Version)))
	buffer.WriteBytes(this.Version)
}

func (this *GetVersion_Out) ByteSize() int {
	size := 4
	size += len(this.Version)
	return size
}

func (this *GetApiCount_In) Decode(buffer *net.Buffer) {
}

func (this *GetApiCount_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(88)
	buffer.WriteUint8(1)
}

func (this *GetApiCount_In) ByteSize() int {
	size := 2
	return size
}

func (this *GetApiCount_Out) Decode(buffer *net.Buffer) {
	this.CountIn = int64(buffer.ReadUint64LE())
	this.CountOut = int64(buffer.ReadUint64LE())
}

func (this *GetApiCount_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(88)
	buffer.WriteUint8(1)
	buffer.WriteUint64LE(uint64(this.CountIn))
	buffer.WriteUint64LE(uint64(this.CountOut))
}

func (this *GetApiCount_Out) ByteSize() int {
	size := 18
	return size
}

func (this *SearchPlayerRole_In) Decode(buffer *net.Buffer) {
	this.Openid = buffer.ReadBytes(int(buffer.ReadUint16LE()))
}

func (this *SearchPlayerRole_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(88)
	buffer.WriteUint8(2)
	buffer.WriteUint16LE(uint16(len(this.Openid)))
	buffer.WriteBytes(this.Openid)
}

func (this *SearchPlayerRole_In) ByteSize() int {
	size := 4
	size += len(this.Openid)
	return size
}

func (this *SearchPlayerRole_Out) Decode(buffer *net.Buffer) {
	this.Result = buffer.ReadUint8() == 1
}

func (this *SearchPlayerRole_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(88)
	buffer.WriteUint8(2)
	if this.Result {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *SearchPlayerRole_Out) ByteSize() int {
	size := 3
	return size
}

func (this *UpdateAccessToken_In) Decode(buffer *net.Buffer) {
	this.Token = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Pfkey = buffer.ReadBytes(int(buffer.ReadUint16LE()))
}

func (this *UpdateAccessToken_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(88)
	buffer.WriteUint8(3)
	buffer.WriteUint16LE(uint16(len(this.Token)))
	buffer.WriteBytes(this.Token)
	buffer.WriteUint16LE(uint16(len(this.Pfkey)))
	buffer.WriteBytes(this.Pfkey)
}

func (this *UpdateAccessToken_In) ByteSize() int {
	size := 6
	size += len(this.Token)
	size += len(this.Pfkey)
	return size
}

func (this *UpdateAccessToken_Out) Decode(buffer *net.Buffer) {
}

func (this *UpdateAccessToken_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(88)
	buffer.WriteUint8(3)
}

func (this *UpdateAccessToken_Out) ByteSize() int {
	size := 2
	return size
}

func (this *UpdateEventData_In) Decode(buffer *net.Buffer) {
	this.Version = int32(buffer.ReadUint32LE())
}

func (this *UpdateEventData_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(88)
	buffer.WriteUint8(4)
	buffer.WriteUint32LE(uint32(this.Version))
}

func (this *UpdateEventData_In) ByteSize() int {
	size := 6
	return size
}

func (this *UpdateEventData_Out) Decode(buffer *net.Buffer) {
	this.Json = buffer.ReadBytes(int(buffer.ReadUint16LE()))
}

func (this *UpdateEventData_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(88)
	buffer.WriteUint8(4)
	buffer.WriteUint16LE(uint16(len(this.Json)))
	buffer.WriteBytes(this.Json)
}

func (this *UpdateEventData_Out) ByteSize() int {
	size := 4
	size += len(this.Json)
	return size
}

func (this *TssData_In) Decode(buffer *net.Buffer) {
	this.Data = buffer.ReadBytes(int(buffer.ReadUint16LE()))
}

func (this *TssData_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(88)
	buffer.WriteUint8(5)
	buffer.WriteUint16LE(uint16(len(this.Data)))
	buffer.WriteBytes(this.Data)
}

func (this *TssData_In) ByteSize() int {
	size := 4
	size += len(this.Data)
	return size
}

func (this *TssData_Out) Decode(buffer *net.Buffer) {
	this.Data = buffer.ReadBytes(int(buffer.ReadUint16LE()))
}

func (this *TssData_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(88)
	buffer.WriteUint8(5)
	buffer.WriteUint16LE(uint16(len(this.Data)))
	buffer.WriteBytes(this.Data)
}

func (this *TssData_Out) ByteSize() int {
	size := 4
	size += len(this.Data)
	return size
}
