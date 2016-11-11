package clique_escort_api

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
	EscortInfo(*net.Session, *EscortInfo_In)
	GetIngotBoat(*net.Session, *GetIngotBoat_In)
	StartEscort(*net.Session, *StartEscort_In)
	HijackBoat(*net.Session, *HijackBoat_In)
	RecoverBoat(*net.Session, *RecoverBoat_In)
	ListBoats(*net.Session, *ListBoats_In)
	GetRandomBoat(*net.Session, *GetRandomBoat_In)
	TakeHijackAward(*net.Session, *TakeHijackAward_In)
	TakeEscortAward(*net.Session, *TakeEscortAward_In)
	GetCliqueBoatMessages(*net.Session, *GetCliqueBoatMessages_In)
	ReadCliqueBoatMessage(*net.Session, *ReadCliqueBoatMessage_In)
}

type OutHandler interface {
	EscortInfo(*net.Session, *EscortInfo_Out)
	GetIngotBoat(*net.Session, *GetIngotBoat_Out)
	StartEscort(*net.Session, *StartEscort_Out)
	HijackBoat(*net.Session, *HijackBoat_Out)
	RecoverBoat(*net.Session, *RecoverBoat_Out)
	ListBoats(*net.Session, *ListBoats_Out)
	GetRandomBoat(*net.Session, *GetRandomBoat_Out)
	NotifyEscortFinished(*net.Session, *NotifyEscortFinished_Out)
	NotifyHijackFinished(*net.Session, *NotifyHijackFinished_Out)
	NotifyRecoverBattleWin(*net.Session, *NotifyRecoverBattleWin_Out)
	NotifyHijackBattleWin(*net.Session, *NotifyHijackBattleWin_Out)
	TakeHijackAward(*net.Session, *TakeHijackAward_Out)
	TakeEscortAward(*net.Session, *TakeEscortAward_Out)
	GetCliqueBoatMessages(*net.Session, *GetCliqueBoatMessages_Out)
	SendCliqueBoatMessage(*net.Session, *SendCliqueBoatMessage_Out)
	ReadCliqueBoatMessage(*net.Session, *ReadCliqueBoatMessage_Out)
	NotifyBoatStatusChange(*net.Session, *NotifyBoatStatusChange_Out)
}

func DecodeIn(message []byte) Request {
	var actionId = message[0]
	var buffer = net.NewBuffer(message[1:])
	switch actionId {
	case 0:
		request := new(EscortInfo_In)
		request.Decode(buffer)
		return request
	case 1:
		request := new(GetIngotBoat_In)
		request.Decode(buffer)
		return request
	case 2:
		request := new(StartEscort_In)
		request.Decode(buffer)
		return request
	case 3:
		request := new(HijackBoat_In)
		request.Decode(buffer)
		return request
	case 4:
		request := new(RecoverBoat_In)
		request.Decode(buffer)
		return request
	case 5:
		request := new(ListBoats_In)
		request.Decode(buffer)
		return request
	case 6:
		request := new(GetRandomBoat_In)
		request.Decode(buffer)
		return request
	case 11:
		request := new(TakeHijackAward_In)
		request.Decode(buffer)
		return request
	case 12:
		request := new(TakeEscortAward_In)
		request.Decode(buffer)
		return request
	case 13:
		request := new(GetCliqueBoatMessages_In)
		request.Decode(buffer)
		return request
	case 15:
		request := new(ReadCliqueBoatMessage_In)
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
		request := new(EscortInfo_Out)
		request.Decode(buffer)
		return request
	case 1:
		request := new(GetIngotBoat_Out)
		request.Decode(buffer)
		return request
	case 2:
		request := new(StartEscort_Out)
		request.Decode(buffer)
		return request
	case 3:
		request := new(HijackBoat_Out)
		request.Decode(buffer)
		return request
	case 4:
		request := new(RecoverBoat_Out)
		request.Decode(buffer)
		return request
	case 5:
		request := new(ListBoats_Out)
		request.Decode(buffer)
		return request
	case 6:
		request := new(GetRandomBoat_Out)
		request.Decode(buffer)
		return request
	case 7:
		request := new(NotifyEscortFinished_Out)
		request.Decode(buffer)
		return request
	case 8:
		request := new(NotifyHijackFinished_Out)
		request.Decode(buffer)
		return request
	case 9:
		request := new(NotifyRecoverBattleWin_Out)
		request.Decode(buffer)
		return request
	case 10:
		request := new(NotifyHijackBattleWin_Out)
		request.Decode(buffer)
		return request
	case 11:
		request := new(TakeHijackAward_Out)
		request.Decode(buffer)
		return request
	case 12:
		request := new(TakeEscortAward_Out)
		request.Decode(buffer)
		return request
	case 13:
		request := new(GetCliqueBoatMessages_Out)
		request.Decode(buffer)
		return request
	case 14:
		request := new(SendCliqueBoatMessage_Out)
		request.Decode(buffer)
		return request
	case 15:
		request := new(ReadCliqueBoatMessage_Out)
		request.Decode(buffer)
		return request
	case 16:
		request := new(NotifyBoatStatusChange_Out)
		request.Decode(buffer)
		return request
	}
	_ = buffer
	panic("unsupported response")
}

type BoatStatusChange int8

const (
	BOAT_STATUS_CHANGE_MY_BOAT_HIJACKING       BoatStatusChange = 0
	BOAT_STATUS_CHANGE_MY_BOAT_HIJACKED        BoatStatusChange = 1
	BOAT_STATUS_CHANGE_MY_BOAT_RECOVERED       BoatStatusChange = 2
	BOAT_STATUS_CHANGE_HIJACKED_BOAT_RECOVERED BoatStatusChange = 3
	BOAT_STATUS_CHANGE_ESCORT_FINISHED         BoatStatusChange = 4
	BOAT_STATUS_CHANGE_HIJACK_FINISHED         BoatStatusChange = 5
)

type RecoverBattleWinResult int8

const (
	RECOVER_BATTLE_WIN_RESULT_SUCCESS       RecoverBattleWinResult = 0
	RECOVER_BATTLE_WIN_RESULT_BOAT_EXPIRE   RecoverBattleWinResult = 1
	RECOVER_BATTLE_WIN_RESULT_NO_PERMISSION RecoverBattleWinResult = 2
)

type HijackBattleWinResult int8

const (
	HIJACK_BATTLE_WIN_RESULT_SUCCESS         HijackBattleWinResult = 0
	HIJACK_BATTLE_WIN_RESULT_ESCORT_FINISHED HijackBattleWinResult = 1
	HIJACK_BATTLE_WIN_RESULT_HIJACKED        HijackBattleWinResult = 2
	HIJACK_BATTLE_WIN_RESULT_NO_PERMISSION   HijackBattleWinResult = 3
)

type StartEscortResult int8

const (
	START_ESCORT_RESULT_SUCCESS        StartEscortResult = 0
	START_ESCORT_RESULT_ILLEGAL_TIME   StartEscortResult = 1
	START_ESCORT_RESULT_ESCORT_NOT_END StartEscortResult = 2
	START_ESCORT_RESULT_NO_CLIQUE      StartEscortResult = 3
	START_ESCORT_RESULT_RUN_OUT        StartEscortResult = 4
	START_ESCORT_RESULT_NO_BOAT        StartEscortResult = 5
)

type HijackBoatResult int8

const (
	HIJACK_BOAT_RESULT_START_BATTLE   HijackBoatResult = 0
	HIJACK_BOAT_RESULT_CLIQUE_MEMBER  HijackBoatResult = 1
	HIJACK_BOAT_RESULT_HIJACK_NOT_END HijackBoatResult = 2
	HIJACK_BOAT_RESULT_NO_CLIQUE      HijackBoatResult = 3
	HIJACK_BOAT_RESULT_RUN_OUT        HijackBoatResult = 4
	HIJACK_BOAT_RESULT_NO_BOAT        HijackBoatResult = 5
	HIJACK_BOAT_RESULT_CAN_NOT_HIJACK HijackBoatResult = 6
)

type RecoverBoatResult int8

const (
	RECOVER_BOAT_RESULT_START_BATTLE    RecoverBoatResult = 0
	RECOVER_BOAT_RESULT_NO_PERMISSION   RecoverBoatResult = 1
	RECOVER_BOAT_RESULT_RECOVERING      RecoverBoatResult = 2
	RECOVER_BOAT_RESULT_NO_BOAT         RecoverBoatResult = 3
	RECOVER_BOAT_RESULT_CAN_NOT_RECOVER RecoverBoatResult = 4
)

type EscortStatus int8

const (
	ESCORT_STATUS_NONE     EscortStatus = 0
	ESCORT_STATUS_ESCORT   EscortStatus = 1
	ESCORT_STATUS_FINISHED EscortStatus = 2
)

type BoatStatus int8

const (
	BOAT_STATUS_ESCORT        BoatStatus = 0
	BOAT_STATUS_HIJACK        BoatStatus = 1
	BOAT_STATUS_RECOVER       BoatStatus = 2
	BOAT_STATUS_HIJACK_FINISH BoatStatus = 3
)

type HijackStatus int8

const (
	HIJACK_STATUS_NONE     HijackStatus = 0
	HIJACK_STATUS_HIJACK   HijackStatus = 1
	HIJACK_STATUS_FINISHED HijackStatus = 2
)

type CliqueBoatMessage struct {
	Id         int64  `json:"id"`
	TplId      int16  `json:"tpl_id"`
	Parameters []byte `json:"parameters"`
}

type CliqueBoat struct {
	BoatId               int64      `json:"boat_id"`
	BoatType             int16      `json:"boat_type"`
	Status               BoatStatus `json:"status"`
	EscortTime           int64      `json:"escort_time"`
	StartTimestamp       int64      `json:"start_timestamp"`
	OwnerPid             int64      `json:"owner_pid"`
	OwnerNick            []byte     `json:"owner_nick"`
	OwnerLevel           int16      `json:"owner_level"`
	HijackerNick         []byte     `json:"hijacker_nick"`
	HijackerPid          int64      `json:"hijacker_pid"`
	CliqueId             int64      `json:"clique_id"`
	CliqueName           []byte     `json:"clique_name"`
	HijackStartTimestamp int64      `json:"hijack_start_timestamp"`
}

func (this *CliqueBoatMessage) Decode(buffer *net.Buffer) {
	this.Id = int64(buffer.ReadUint64LE())
	this.TplId = int16(buffer.ReadUint16LE())
	this.Parameters = buffer.ReadBytes(int(buffer.ReadUint16LE()))
}

func (this *CliqueBoatMessage) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.Id))
	buffer.WriteUint16LE(uint16(this.TplId))
	buffer.WriteUint16LE(uint16(len(this.Parameters)))
	buffer.WriteBytes(this.Parameters)
}

func (this *CliqueBoatMessage) ByteSize() int {
	size := 12
	size += len(this.Parameters)
	return size
}

func (this *CliqueBoat) Decode(buffer *net.Buffer) {
	this.BoatId = int64(buffer.ReadUint64LE())
	this.BoatType = int16(buffer.ReadUint16LE())
	this.Status = BoatStatus(buffer.ReadUint8())
	this.EscortTime = int64(buffer.ReadUint64LE())
	this.StartTimestamp = int64(buffer.ReadUint64LE())
	this.OwnerPid = int64(buffer.ReadUint64LE())
	this.OwnerNick = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.OwnerLevel = int16(buffer.ReadUint16LE())
	this.HijackerNick = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.HijackerPid = int64(buffer.ReadUint64LE())
	this.CliqueId = int64(buffer.ReadUint64LE())
	this.CliqueName = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.HijackStartTimestamp = int64(buffer.ReadUint64LE())
}

func (this *CliqueBoat) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.BoatId))
	buffer.WriteUint16LE(uint16(this.BoatType))
	buffer.WriteUint8(uint8(this.Status))
	buffer.WriteUint64LE(uint64(this.EscortTime))
	buffer.WriteUint64LE(uint64(this.StartTimestamp))
	buffer.WriteUint64LE(uint64(this.OwnerPid))
	buffer.WriteUint16LE(uint16(len(this.OwnerNick)))
	buffer.WriteBytes(this.OwnerNick)
	buffer.WriteUint16LE(uint16(this.OwnerLevel))
	buffer.WriteUint16LE(uint16(len(this.HijackerNick)))
	buffer.WriteBytes(this.HijackerNick)
	buffer.WriteUint64LE(uint64(this.HijackerPid))
	buffer.WriteUint64LE(uint64(this.CliqueId))
	buffer.WriteUint16LE(uint16(len(this.CliqueName)))
	buffer.WriteBytes(this.CliqueName)
	buffer.WriteUint64LE(uint64(this.HijackStartTimestamp))
}

func (this *CliqueBoat) ByteSize() int {
	size := 67
	size += len(this.OwnerNick)
	size += len(this.HijackerNick)
	size += len(this.CliqueName)
	return size
}

type EscortInfo_In struct {
}

func (this *EscortInfo_In) Process(session *net.Session) {
	g_InHandler.EscortInfo(session, this)
}

func (this *EscortInfo_In) TypeName() string {
	return "clique_escort.escort_info.in"
}

func (this *EscortInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 36, 0
}

type EscortInfo_Out struct {
	DailyEscortNum  int16        `json:"daily_escort_num"`
	DailyHijackNum  int16        `json:"daily_hijack_num"`
	BoughtHijackNum int16        `json:"bought_hijack_num"`
	EscortStatus    EscortStatus `json:"escort_status"`
	HijackStatus    HijackStatus `json:"hijack_status"`
	BoatType        int16        `json:"boat_type"`
}

func (this *EscortInfo_Out) Process(session *net.Session) {
	g_OutHandler.EscortInfo(session, this)
}

func (this *EscortInfo_Out) TypeName() string {
	return "clique_escort.escort_info.out"
}

func (this *EscortInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 36, 0
}

func (this *EscortInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetIngotBoat_In struct {
}

func (this *GetIngotBoat_In) Process(session *net.Session) {
	g_InHandler.GetIngotBoat(session, this)
}

func (this *GetIngotBoat_In) TypeName() string {
	return "clique_escort.get_ingot_boat.in"
}

func (this *GetIngotBoat_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 36, 1
}

type GetIngotBoat_Out struct {
	Ok bool `json:"ok"`
}

func (this *GetIngotBoat_Out) Process(session *net.Session) {
	g_OutHandler.GetIngotBoat(session, this)
}

func (this *GetIngotBoat_Out) TypeName() string {
	return "clique_escort.get_ingot_boat.out"
}

func (this *GetIngotBoat_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 36, 1
}

func (this *GetIngotBoat_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type StartEscort_In struct {
}

func (this *StartEscort_In) Process(session *net.Session) {
	g_InHandler.StartEscort(session, this)
}

func (this *StartEscort_In) TypeName() string {
	return "clique_escort.start_escort.in"
}

func (this *StartEscort_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 36, 2
}

type StartEscort_Out struct {
	Result StartEscortResult `json:"result"`
	Boat   CliqueBoat        `json:"boat"`
}

func (this *StartEscort_Out) Process(session *net.Session) {
	g_OutHandler.StartEscort(session, this)
}

func (this *StartEscort_Out) TypeName() string {
	return "clique_escort.start_escort.out"
}

func (this *StartEscort_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 36, 2
}

func (this *StartEscort_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type HijackBoat_In struct {
	BoatId int64 `json:"boat_id"`
	Pay    bool  `json:"pay"`
}

func (this *HijackBoat_In) Process(session *net.Session) {
	g_InHandler.HijackBoat(session, this)
}

func (this *HijackBoat_In) TypeName() string {
	return "clique_escort.hijack_boat.in"
}

func (this *HijackBoat_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 36, 3
}

type HijackBoat_Out struct {
	Result        HijackBoatResult `json:"result"`
	BoatId        int64            `json:"boat_id"`
	BoatOwnerPid  int64            `json:"boat_owner_pid"`
	BoatOwnerNick []byte           `json:"boat_owner_nick"`
}

func (this *HijackBoat_Out) Process(session *net.Session) {
	g_OutHandler.HijackBoat(session, this)
}

func (this *HijackBoat_Out) TypeName() string {
	return "clique_escort.hijack_boat.out"
}

func (this *HijackBoat_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 36, 3
}

func (this *HijackBoat_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type RecoverBoat_In struct {
	BoatId int64 `json:"boat_id"`
}

func (this *RecoverBoat_In) Process(session *net.Session) {
	g_InHandler.RecoverBoat(session, this)
}

func (this *RecoverBoat_In) TypeName() string {
	return "clique_escort.recover_boat.in"
}

func (this *RecoverBoat_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 36, 4
}

type RecoverBoat_Out struct {
	BoatId       int64             `json:"boat_id"`
	Result       RecoverBoatResult `json:"result"`
	HijackerPid  int64             `json:"hijacker_pid"`
	HijackerNick []byte            `json:"hijacker_nick"`
}

func (this *RecoverBoat_Out) Process(session *net.Session) {
	g_OutHandler.RecoverBoat(session, this)
}

func (this *RecoverBoat_Out) TypeName() string {
	return "clique_escort.recover_boat.out"
}

func (this *RecoverBoat_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 36, 4
}

func (this *RecoverBoat_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ListBoats_In struct {
}

func (this *ListBoats_In) Process(session *net.Session) {
	g_InHandler.ListBoats(session, this)
}

func (this *ListBoats_In) TypeName() string {
	return "clique_escort.list_boats.in"
}

func (this *ListBoats_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 36, 5
}

type ListBoats_Out struct {
	Boats []ListBoats_Out_Boats `json:"boats"`
}

type ListBoats_Out_Boats struct {
	Boat CliqueBoat `json:"boat"`
}

func (this *ListBoats_Out) Process(session *net.Session) {
	g_OutHandler.ListBoats(session, this)
}

func (this *ListBoats_Out) TypeName() string {
	return "clique_escort.list_boats.out"
}

func (this *ListBoats_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 36, 5
}

func (this *ListBoats_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetRandomBoat_In struct {
}

func (this *GetRandomBoat_In) Process(session *net.Session) {
	g_InHandler.GetRandomBoat(session, this)
}

func (this *GetRandomBoat_In) TypeName() string {
	return "clique_escort.get_random_boat.in"
}

func (this *GetRandomBoat_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 36, 6
}

type GetRandomBoat_Out struct {
	BoatType int16 `json:"boat_type"`
}

func (this *GetRandomBoat_Out) Process(session *net.Session) {
	g_OutHandler.GetRandomBoat(session, this)
}

func (this *GetRandomBoat_Out) TypeName() string {
	return "clique_escort.get_random_boat.out"
}

func (this *GetRandomBoat_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 36, 6
}

func (this *GetRandomBoat_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type NotifyEscortFinished_Out struct {
}

func (this *NotifyEscortFinished_Out) Process(session *net.Session) {
	g_OutHandler.NotifyEscortFinished(session, this)
}

func (this *NotifyEscortFinished_Out) TypeName() string {
	return "clique_escort.notify_escort_finished.out"
}

func (this *NotifyEscortFinished_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 36, 7
}

func (this *NotifyEscortFinished_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type NotifyHijackFinished_Out struct {
}

func (this *NotifyHijackFinished_Out) Process(session *net.Session) {
	g_OutHandler.NotifyHijackFinished(session, this)
}

func (this *NotifyHijackFinished_Out) TypeName() string {
	return "clique_escort.notify_hijack_finished.out"
}

func (this *NotifyHijackFinished_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 36, 8
}

func (this *NotifyHijackFinished_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type NotifyRecoverBattleWin_Out struct {
	BoatId    int64                  `json:"boat_id"`
	Result    RecoverBattleWinResult `json:"result"`
	OwnerNick []byte                 `json:"owner_nick"`
}

func (this *NotifyRecoverBattleWin_Out) Process(session *net.Session) {
	g_OutHandler.NotifyRecoverBattleWin(session, this)
}

func (this *NotifyRecoverBattleWin_Out) TypeName() string {
	return "clique_escort.notify_recover_battle_win.out"
}

func (this *NotifyRecoverBattleWin_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 36, 9
}

func (this *NotifyRecoverBattleWin_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type NotifyHijackBattleWin_Out struct {
	BoatId    int64                 `json:"boat_id"`
	Result    HijackBattleWinResult `json:"result"`
	OwnerNick []byte                `json:"owner_nick"`
}

func (this *NotifyHijackBattleWin_Out) Process(session *net.Session) {
	g_OutHandler.NotifyHijackBattleWin(session, this)
}

func (this *NotifyHijackBattleWin_Out) TypeName() string {
	return "clique_escort.notify_hijack_battle_win.out"
}

func (this *NotifyHijackBattleWin_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 36, 10
}

func (this *NotifyHijackBattleWin_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type TakeHijackAward_In struct {
}

func (this *TakeHijackAward_In) Process(session *net.Session) {
	g_InHandler.TakeHijackAward(session, this)
}

func (this *TakeHijackAward_In) TypeName() string {
	return "clique_escort.take_hijack_award.in"
}

func (this *TakeHijackAward_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 36, 11
}

type TakeHijackAward_Out struct {
	Ok bool `json:"ok"`
}

func (this *TakeHijackAward_Out) Process(session *net.Session) {
	g_OutHandler.TakeHijackAward(session, this)
}

func (this *TakeHijackAward_Out) TypeName() string {
	return "clique_escort.take_hijack_award.out"
}

func (this *TakeHijackAward_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 36, 11
}

func (this *TakeHijackAward_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type TakeEscortAward_In struct {
}

func (this *TakeEscortAward_In) Process(session *net.Session) {
	g_InHandler.TakeEscortAward(session, this)
}

func (this *TakeEscortAward_In) TypeName() string {
	return "clique_escort.take_escort_award.in"
}

func (this *TakeEscortAward_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 36, 12
}

type TakeEscortAward_Out struct {
	Ok bool `json:"ok"`
}

func (this *TakeEscortAward_Out) Process(session *net.Session) {
	g_OutHandler.TakeEscortAward(session, this)
}

func (this *TakeEscortAward_Out) TypeName() string {
	return "clique_escort.take_escort_award.out"
}

func (this *TakeEscortAward_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 36, 12
}

func (this *TakeEscortAward_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type GetCliqueBoatMessages_In struct {
}

func (this *GetCliqueBoatMessages_In) Process(session *net.Session) {
	g_InHandler.GetCliqueBoatMessages(session, this)
}

func (this *GetCliqueBoatMessages_In) TypeName() string {
	return "clique_escort.get_clique_boat_messages.in"
}

func (this *GetCliqueBoatMessages_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 36, 13
}

type GetCliqueBoatMessages_Out struct {
	Messages []GetCliqueBoatMessages_Out_Messages `json:"messages"`
}

type GetCliqueBoatMessages_Out_Messages struct {
	Message CliqueBoatMessage `json:"message"`
}

func (this *GetCliqueBoatMessages_Out) Process(session *net.Session) {
	g_OutHandler.GetCliqueBoatMessages(session, this)
}

func (this *GetCliqueBoatMessages_Out) TypeName() string {
	return "clique_escort.get_clique_boat_messages.out"
}

func (this *GetCliqueBoatMessages_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 36, 13
}

func (this *GetCliqueBoatMessages_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type SendCliqueBoatMessage_Out struct {
	Message CliqueBoatMessage `json:"message"`
}

func (this *SendCliqueBoatMessage_Out) Process(session *net.Session) {
	g_OutHandler.SendCliqueBoatMessage(session, this)
}

func (this *SendCliqueBoatMessage_Out) TypeName() string {
	return "clique_escort.send_clique_boat_message.out"
}

func (this *SendCliqueBoatMessage_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 36, 14
}

func (this *SendCliqueBoatMessage_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ReadCliqueBoatMessage_In struct {
	Id int64 `json:"id"`
}

func (this *ReadCliqueBoatMessage_In) Process(session *net.Session) {
	g_InHandler.ReadCliqueBoatMessage(session, this)
}

func (this *ReadCliqueBoatMessage_In) TypeName() string {
	return "clique_escort.read_clique_boat_message.in"
}

func (this *ReadCliqueBoatMessage_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 36, 15
}

type ReadCliqueBoatMessage_Out struct {
}

func (this *ReadCliqueBoatMessage_Out) Process(session *net.Session) {
	g_OutHandler.ReadCliqueBoatMessage(session, this)
}

func (this *ReadCliqueBoatMessage_Out) TypeName() string {
	return "clique_escort.read_clique_boat_message.out"
}

func (this *ReadCliqueBoatMessage_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 36, 15
}

func (this *ReadCliqueBoatMessage_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type NotifyBoatStatusChange_Out struct {
	Boat   CliqueBoat       `json:"boat"`
	Change BoatStatusChange `json:"change"`
}

func (this *NotifyBoatStatusChange_Out) Process(session *net.Session) {
	g_OutHandler.NotifyBoatStatusChange(session, this)
}

func (this *NotifyBoatStatusChange_Out) TypeName() string {
	return "clique_escort.notify_boat_status_change.out"
}

func (this *NotifyBoatStatusChange_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 36, 16
}

func (this *NotifyBoatStatusChange_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

func (this *EscortInfo_In) Decode(buffer *net.Buffer) {
}

func (this *EscortInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(36)
	buffer.WriteUint8(0)
}

func (this *EscortInfo_In) ByteSize() int {
	size := 2
	return size
}

func (this *EscortInfo_Out) Decode(buffer *net.Buffer) {
	this.DailyEscortNum = int16(buffer.ReadUint16LE())
	this.DailyHijackNum = int16(buffer.ReadUint16LE())
	this.BoughtHijackNum = int16(buffer.ReadUint16LE())
	this.EscortStatus = EscortStatus(buffer.ReadUint8())
	this.HijackStatus = HijackStatus(buffer.ReadUint8())
	this.BoatType = int16(buffer.ReadUint16LE())
}

func (this *EscortInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(36)
	buffer.WriteUint8(0)
	buffer.WriteUint16LE(uint16(this.DailyEscortNum))
	buffer.WriteUint16LE(uint16(this.DailyHijackNum))
	buffer.WriteUint16LE(uint16(this.BoughtHijackNum))
	buffer.WriteUint8(uint8(this.EscortStatus))
	buffer.WriteUint8(uint8(this.HijackStatus))
	buffer.WriteUint16LE(uint16(this.BoatType))
}

func (this *EscortInfo_Out) ByteSize() int {
	size := 12
	return size
}

func (this *GetIngotBoat_In) Decode(buffer *net.Buffer) {
}

func (this *GetIngotBoat_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(36)
	buffer.WriteUint8(1)
}

func (this *GetIngotBoat_In) ByteSize() int {
	size := 2
	return size
}

func (this *GetIngotBoat_Out) Decode(buffer *net.Buffer) {
	this.Ok = buffer.ReadUint8() == 1
}

func (this *GetIngotBoat_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(36)
	buffer.WriteUint8(1)
	if this.Ok {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *GetIngotBoat_Out) ByteSize() int {
	size := 3
	return size
}

func (this *StartEscort_In) Decode(buffer *net.Buffer) {
}

func (this *StartEscort_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(36)
	buffer.WriteUint8(2)
}

func (this *StartEscort_In) ByteSize() int {
	size := 2
	return size
}

func (this *StartEscort_Out) Decode(buffer *net.Buffer) {
	this.Result = StartEscortResult(buffer.ReadUint8())
	this.Boat.Decode(buffer)
}

func (this *StartEscort_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(36)
	buffer.WriteUint8(2)
	buffer.WriteUint8(uint8(this.Result))
	this.Boat.Encode(buffer)
}

func (this *StartEscort_Out) ByteSize() int {
	size := 3
	size += this.Boat.ByteSize()
	return size
}

func (this *HijackBoat_In) Decode(buffer *net.Buffer) {
	this.BoatId = int64(buffer.ReadUint64LE())
	this.Pay = buffer.ReadUint8() == 1
}

func (this *HijackBoat_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(36)
	buffer.WriteUint8(3)
	buffer.WriteUint64LE(uint64(this.BoatId))
	if this.Pay {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *HijackBoat_In) ByteSize() int {
	size := 11
	return size
}

func (this *HijackBoat_Out) Decode(buffer *net.Buffer) {
	this.Result = HijackBoatResult(buffer.ReadUint8())
	this.BoatId = int64(buffer.ReadUint64LE())
	this.BoatOwnerPid = int64(buffer.ReadUint64LE())
	this.BoatOwnerNick = buffer.ReadBytes(int(buffer.ReadUint16LE()))
}

func (this *HijackBoat_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(36)
	buffer.WriteUint8(3)
	buffer.WriteUint8(uint8(this.Result))
	buffer.WriteUint64LE(uint64(this.BoatId))
	buffer.WriteUint64LE(uint64(this.BoatOwnerPid))
	buffer.WriteUint16LE(uint16(len(this.BoatOwnerNick)))
	buffer.WriteBytes(this.BoatOwnerNick)
}

func (this *HijackBoat_Out) ByteSize() int {
	size := 21
	size += len(this.BoatOwnerNick)
	return size
}

func (this *RecoverBoat_In) Decode(buffer *net.Buffer) {
	this.BoatId = int64(buffer.ReadUint64LE())
}

func (this *RecoverBoat_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(36)
	buffer.WriteUint8(4)
	buffer.WriteUint64LE(uint64(this.BoatId))
}

func (this *RecoverBoat_In) ByteSize() int {
	size := 10
	return size
}

func (this *RecoverBoat_Out) Decode(buffer *net.Buffer) {
	this.BoatId = int64(buffer.ReadUint64LE())
	this.Result = RecoverBoatResult(buffer.ReadUint8())
	this.HijackerPid = int64(buffer.ReadUint64LE())
	this.HijackerNick = buffer.ReadBytes(int(buffer.ReadUint16LE()))
}

func (this *RecoverBoat_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(36)
	buffer.WriteUint8(4)
	buffer.WriteUint64LE(uint64(this.BoatId))
	buffer.WriteUint8(uint8(this.Result))
	buffer.WriteUint64LE(uint64(this.HijackerPid))
	buffer.WriteUint16LE(uint16(len(this.HijackerNick)))
	buffer.WriteBytes(this.HijackerNick)
}

func (this *RecoverBoat_Out) ByteSize() int {
	size := 21
	size += len(this.HijackerNick)
	return size
}

func (this *ListBoats_In) Decode(buffer *net.Buffer) {
}

func (this *ListBoats_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(36)
	buffer.WriteUint8(5)
}

func (this *ListBoats_In) ByteSize() int {
	size := 2
	return size
}

func (this *ListBoats_Out) Decode(buffer *net.Buffer) {
	this.Boats = make([]ListBoats_Out_Boats, buffer.ReadUint8())
	for i := 0; i < len(this.Boats); i++ {
		this.Boats[i].Decode(buffer)
	}
}

func (this *ListBoats_Out_Boats) Decode(buffer *net.Buffer) {
	this.Boat.Decode(buffer)
}

func (this *ListBoats_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(36)
	buffer.WriteUint8(5)
	buffer.WriteUint8(uint8(len(this.Boats)))
	for i := 0; i < len(this.Boats); i++ {
		this.Boats[i].Encode(buffer)
	}
}

func (this *ListBoats_Out_Boats) Encode(buffer *net.Buffer) {
	this.Boat.Encode(buffer)
}

func (this *ListBoats_Out) ByteSize() int {
	size := 3
	for i := 0; i < len(this.Boats); i++ {
		size += this.Boats[i].ByteSize()
	}
	return size
}

func (this *ListBoats_Out_Boats) ByteSize() int {
	size := 0
	size += this.Boat.ByteSize()
	return size
}

func (this *GetRandomBoat_In) Decode(buffer *net.Buffer) {
}

func (this *GetRandomBoat_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(36)
	buffer.WriteUint8(6)
}

func (this *GetRandomBoat_In) ByteSize() int {
	size := 2
	return size
}

func (this *GetRandomBoat_Out) Decode(buffer *net.Buffer) {
	this.BoatType = int16(buffer.ReadUint16LE())
}

func (this *GetRandomBoat_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(36)
	buffer.WriteUint8(6)
	buffer.WriteUint16LE(uint16(this.BoatType))
}

func (this *GetRandomBoat_Out) ByteSize() int {
	size := 4
	return size
}

func (this *NotifyEscortFinished_Out) Decode(buffer *net.Buffer) {
}

func (this *NotifyEscortFinished_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(36)
	buffer.WriteUint8(7)
}

func (this *NotifyEscortFinished_Out) ByteSize() int {
	size := 2
	return size
}

func (this *NotifyHijackFinished_Out) Decode(buffer *net.Buffer) {
}

func (this *NotifyHijackFinished_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(36)
	buffer.WriteUint8(8)
}

func (this *NotifyHijackFinished_Out) ByteSize() int {
	size := 2
	return size
}

func (this *NotifyRecoverBattleWin_Out) Decode(buffer *net.Buffer) {
	this.BoatId = int64(buffer.ReadUint64LE())
	this.Result = RecoverBattleWinResult(buffer.ReadUint8())
	this.OwnerNick = buffer.ReadBytes(int(buffer.ReadUint16LE()))
}

func (this *NotifyRecoverBattleWin_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(36)
	buffer.WriteUint8(9)
	buffer.WriteUint64LE(uint64(this.BoatId))
	buffer.WriteUint8(uint8(this.Result))
	buffer.WriteUint16LE(uint16(len(this.OwnerNick)))
	buffer.WriteBytes(this.OwnerNick)
}

func (this *NotifyRecoverBattleWin_Out) ByteSize() int {
	size := 13
	size += len(this.OwnerNick)
	return size
}

func (this *NotifyHijackBattleWin_Out) Decode(buffer *net.Buffer) {
	this.BoatId = int64(buffer.ReadUint64LE())
	this.Result = HijackBattleWinResult(buffer.ReadUint8())
	this.OwnerNick = buffer.ReadBytes(int(buffer.ReadUint16LE()))
}

func (this *NotifyHijackBattleWin_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(36)
	buffer.WriteUint8(10)
	buffer.WriteUint64LE(uint64(this.BoatId))
	buffer.WriteUint8(uint8(this.Result))
	buffer.WriteUint16LE(uint16(len(this.OwnerNick)))
	buffer.WriteBytes(this.OwnerNick)
}

func (this *NotifyHijackBattleWin_Out) ByteSize() int {
	size := 13
	size += len(this.OwnerNick)
	return size
}

func (this *TakeHijackAward_In) Decode(buffer *net.Buffer) {
}

func (this *TakeHijackAward_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(36)
	buffer.WriteUint8(11)
}

func (this *TakeHijackAward_In) ByteSize() int {
	size := 2
	return size
}

func (this *TakeHijackAward_Out) Decode(buffer *net.Buffer) {
	this.Ok = buffer.ReadUint8() == 1
}

func (this *TakeHijackAward_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(36)
	buffer.WriteUint8(11)
	if this.Ok {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *TakeHijackAward_Out) ByteSize() int {
	size := 3
	return size
}

func (this *TakeEscortAward_In) Decode(buffer *net.Buffer) {
}

func (this *TakeEscortAward_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(36)
	buffer.WriteUint8(12)
}

func (this *TakeEscortAward_In) ByteSize() int {
	size := 2
	return size
}

func (this *TakeEscortAward_Out) Decode(buffer *net.Buffer) {
	this.Ok = buffer.ReadUint8() == 1
}

func (this *TakeEscortAward_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(36)
	buffer.WriteUint8(12)
	if this.Ok {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *TakeEscortAward_Out) ByteSize() int {
	size := 3
	return size
}

func (this *GetCliqueBoatMessages_In) Decode(buffer *net.Buffer) {
}

func (this *GetCliqueBoatMessages_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(36)
	buffer.WriteUint8(13)
}

func (this *GetCliqueBoatMessages_In) ByteSize() int {
	size := 2
	return size
}

func (this *GetCliqueBoatMessages_Out) Decode(buffer *net.Buffer) {
	this.Messages = make([]GetCliqueBoatMessages_Out_Messages, buffer.ReadUint8())
	for i := 0; i < len(this.Messages); i++ {
		this.Messages[i].Decode(buffer)
	}
}

func (this *GetCliqueBoatMessages_Out_Messages) Decode(buffer *net.Buffer) {
	this.Message.Decode(buffer)
}

func (this *GetCliqueBoatMessages_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(36)
	buffer.WriteUint8(13)
	buffer.WriteUint8(uint8(len(this.Messages)))
	for i := 0; i < len(this.Messages); i++ {
		this.Messages[i].Encode(buffer)
	}
}

func (this *GetCliqueBoatMessages_Out_Messages) Encode(buffer *net.Buffer) {
	this.Message.Encode(buffer)
}

func (this *GetCliqueBoatMessages_Out) ByteSize() int {
	size := 3
	for i := 0; i < len(this.Messages); i++ {
		size += this.Messages[i].ByteSize()
	}
	return size
}

func (this *GetCliqueBoatMessages_Out_Messages) ByteSize() int {
	size := 0
	size += this.Message.ByteSize()
	return size
}

func (this *SendCliqueBoatMessage_Out) Decode(buffer *net.Buffer) {
	this.Message.Decode(buffer)
}

func (this *SendCliqueBoatMessage_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(36)
	buffer.WriteUint8(14)
	this.Message.Encode(buffer)
}

func (this *SendCliqueBoatMessage_Out) ByteSize() int {
	size := 2
	size += this.Message.ByteSize()
	return size
}

func (this *ReadCliqueBoatMessage_In) Decode(buffer *net.Buffer) {
	this.Id = int64(buffer.ReadUint64LE())
}

func (this *ReadCliqueBoatMessage_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(36)
	buffer.WriteUint8(15)
	buffer.WriteUint64LE(uint64(this.Id))
}

func (this *ReadCliqueBoatMessage_In) ByteSize() int {
	size := 10
	return size
}

func (this *ReadCliqueBoatMessage_Out) Decode(buffer *net.Buffer) {
}

func (this *ReadCliqueBoatMessage_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(36)
	buffer.WriteUint8(15)
}

func (this *ReadCliqueBoatMessage_Out) ByteSize() int {
	size := 2
	return size
}

func (this *NotifyBoatStatusChange_Out) Decode(buffer *net.Buffer) {
	this.Boat.Decode(buffer)
	this.Change = BoatStatusChange(buffer.ReadUint8())
}

func (this *NotifyBoatStatusChange_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(36)
	buffer.WriteUint8(16)
	this.Boat.Encode(buffer)
	buffer.WriteUint8(uint8(this.Change))
}

func (this *NotifyBoatStatusChange_Out) ByteSize() int {
	size := 3
	size += this.Boat.ByteSize()
	return size
}
