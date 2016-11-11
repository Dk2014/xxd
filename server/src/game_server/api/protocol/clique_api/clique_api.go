package clique_api

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
	CreateClique(*net.Session, *CreateClique_In)
	ApplyJoinClique(*net.Session, *ApplyJoinClique_In)
	CancelApplyClique(*net.Session, *CancelApplyClique_In)
	ProcessJoinApply(*net.Session, *ProcessJoinApply_In)
	ElectOwner(*net.Session, *ElectOwner_In)
	MangeMember(*net.Session, *MangeMember_In)
	DestoryClique(*net.Session, *DestoryClique_In)
	UpdateAnnounce(*net.Session, *UpdateAnnounce_In)
	LeaveClique(*net.Session, *LeaveClique_In)
	ListClique(*net.Session, *ListClique_In)
	CliquePublicInfo(*net.Session, *CliquePublicInfo_In)
	CliqueInfo(*net.Session, *CliqueInfo_In)
	ListApply(*net.Session, *ListApply_In)
	EnterClubhouse(*net.Session, *EnterClubhouse_In)
	LeaveClubhouse(*net.Session, *LeaveClubhouse_In)
	ClubMove(*net.Session, *ClubMove_In)
	CliquePublicInfoSummary(*net.Session, *CliquePublicInfoSummary_In)
	CliqueAutoAudit(*net.Session, *CliqueAutoAudit_In)
	CliqueBaseDonate(*net.Session, *CliqueBaseDonate_In)
	CliqueRecruitment(*net.Session, *CliqueRecruitment_In)
	QuickApply(*net.Session, *QuickApply_In)
	OtherClique(*net.Session, *OtherClique_In)
}

type OutHandler interface {
	CreateClique(*net.Session, *CreateClique_Out)
	ApplyJoinClique(*net.Session, *ApplyJoinClique_Out)
	CancelApplyClique(*net.Session, *CancelApplyClique_Out)
	ProcessJoinApply(*net.Session, *ProcessJoinApply_Out)
	ElectOwner(*net.Session, *ElectOwner_Out)
	MangeMember(*net.Session, *MangeMember_Out)
	DestoryClique(*net.Session, *DestoryClique_Out)
	UpdateAnnounce(*net.Session, *UpdateAnnounce_Out)
	LeaveClique(*net.Session, *LeaveClique_Out)
	ListClique(*net.Session, *ListClique_Out)
	CliquePublicInfo(*net.Session, *CliquePublicInfo_Out)
	CliqueInfo(*net.Session, *CliqueInfo_Out)
	ListApply(*net.Session, *ListApply_Out)
	OperaErrorNotify(*net.Session, *OperaErrorNotify_Out)
	EnterClubhouse(*net.Session, *EnterClubhouse_Out)
	LeaveClubhouse(*net.Session, *LeaveClubhouse_Out)
	ClubMove(*net.Session, *ClubMove_Out)
	NotifyClubhousePlayers(*net.Session, *NotifyClubhousePlayers_Out)
	NotifyNewPlayer(*net.Session, *NotifyNewPlayer_Out)
	NotifyUpdatePlayer(*net.Session, *NotifyUpdatePlayer_Out)
	CliquePublicInfoSummary(*net.Session, *CliquePublicInfoSummary_Out)
	CliqueAutoAudit(*net.Session, *CliqueAutoAudit_Out)
	CliqueBaseDonate(*net.Session, *CliqueBaseDonate_Out)
	NotifyLeaveClique(*net.Session, *NotifyLeaveClique_Out)
	NotifyJoincliqueSuccess(*net.Session, *NotifyJoincliqueSuccess_Out)
	NotifyCliqueMangerAction(*net.Session, *NotifyCliqueMangerAction_Out)
	CliqueRecruitment(*net.Session, *CliqueRecruitment_Out)
	NotifyCliqueAnnounce(*net.Session, *NotifyCliqueAnnounce_Out)
	NotifyCliqueElectowner(*net.Session, *NotifyCliqueElectowner_Out)
	QuickApply(*net.Session, *QuickApply_Out)
	NotifyContribChange(*net.Session, *NotifyContribChange_Out)
	OtherClique(*net.Session, *OtherClique_Out)
}

func DecodeIn(message []byte) Request {
	var actionId = message[0]
	var buffer = net.NewBuffer(message[1:])
	switch actionId {
	case 0:
		request := new(CreateClique_In)
		request.Decode(buffer)
		return request
	case 1:
		request := new(ApplyJoinClique_In)
		request.Decode(buffer)
		return request
	case 2:
		request := new(CancelApplyClique_In)
		request.Decode(buffer)
		return request
	case 3:
		request := new(ProcessJoinApply_In)
		request.Decode(buffer)
		return request
	case 4:
		request := new(ElectOwner_In)
		request.Decode(buffer)
		return request
	case 5:
		request := new(MangeMember_In)
		request.Decode(buffer)
		return request
	case 6:
		request := new(DestoryClique_In)
		request.Decode(buffer)
		return request
	case 7:
		request := new(UpdateAnnounce_In)
		request.Decode(buffer)
		return request
	case 8:
		request := new(LeaveClique_In)
		request.Decode(buffer)
		return request
	case 9:
		request := new(ListClique_In)
		request.Decode(buffer)
		return request
	case 10:
		request := new(CliquePublicInfo_In)
		request.Decode(buffer)
		return request
	case 11:
		request := new(CliqueInfo_In)
		request.Decode(buffer)
		return request
	case 12:
		request := new(ListApply_In)
		request.Decode(buffer)
		return request
	case 14:
		request := new(EnterClubhouse_In)
		request.Decode(buffer)
		return request
	case 15:
		request := new(LeaveClubhouse_In)
		request.Decode(buffer)
		return request
	case 16:
		request := new(ClubMove_In)
		request.Decode(buffer)
		return request
	case 20:
		request := new(CliquePublicInfoSummary_In)
		request.Decode(buffer)
		return request
	case 21:
		request := new(CliqueAutoAudit_In)
		request.Decode(buffer)
		return request
	case 22:
		request := new(CliqueBaseDonate_In)
		request.Decode(buffer)
		return request
	case 26:
		request := new(CliqueRecruitment_In)
		request.Decode(buffer)
		return request
	case 29:
		request := new(QuickApply_In)
		request.Decode(buffer)
		return request
	case 31:
		request := new(OtherClique_In)
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
		request := new(CreateClique_Out)
		request.Decode(buffer)
		return request
	case 1:
		request := new(ApplyJoinClique_Out)
		request.Decode(buffer)
		return request
	case 2:
		request := new(CancelApplyClique_Out)
		request.Decode(buffer)
		return request
	case 3:
		request := new(ProcessJoinApply_Out)
		request.Decode(buffer)
		return request
	case 4:
		request := new(ElectOwner_Out)
		request.Decode(buffer)
		return request
	case 5:
		request := new(MangeMember_Out)
		request.Decode(buffer)
		return request
	case 6:
		request := new(DestoryClique_Out)
		request.Decode(buffer)
		return request
	case 7:
		request := new(UpdateAnnounce_Out)
		request.Decode(buffer)
		return request
	case 8:
		request := new(LeaveClique_Out)
		request.Decode(buffer)
		return request
	case 9:
		request := new(ListClique_Out)
		request.Decode(buffer)
		return request
	case 10:
		request := new(CliquePublicInfo_Out)
		request.Decode(buffer)
		return request
	case 11:
		request := new(CliqueInfo_Out)
		request.Decode(buffer)
		return request
	case 12:
		request := new(ListApply_Out)
		request.Decode(buffer)
		return request
	case 13:
		request := new(OperaErrorNotify_Out)
		request.Decode(buffer)
		return request
	case 14:
		request := new(EnterClubhouse_Out)
		request.Decode(buffer)
		return request
	case 15:
		request := new(LeaveClubhouse_Out)
		request.Decode(buffer)
		return request
	case 16:
		request := new(ClubMove_Out)
		request.Decode(buffer)
		return request
	case 17:
		request := new(NotifyClubhousePlayers_Out)
		request.Decode(buffer)
		return request
	case 18:
		request := new(NotifyNewPlayer_Out)
		request.Decode(buffer)
		return request
	case 19:
		request := new(NotifyUpdatePlayer_Out)
		request.Decode(buffer)
		return request
	case 20:
		request := new(CliquePublicInfoSummary_Out)
		request.Decode(buffer)
		return request
	case 21:
		request := new(CliqueAutoAudit_Out)
		request.Decode(buffer)
		return request
	case 22:
		request := new(CliqueBaseDonate_Out)
		request.Decode(buffer)
		return request
	case 23:
		request := new(NotifyLeaveClique_Out)
		request.Decode(buffer)
		return request
	case 24:
		request := new(NotifyJoincliqueSuccess_Out)
		request.Decode(buffer)
		return request
	case 25:
		request := new(NotifyCliqueMangerAction_Out)
		request.Decode(buffer)
		return request
	case 26:
		request := new(CliqueRecruitment_Out)
		request.Decode(buffer)
		return request
	case 27:
		request := new(NotifyCliqueAnnounce_Out)
		request.Decode(buffer)
		return request
	case 28:
		request := new(NotifyCliqueElectowner_Out)
		request.Decode(buffer)
		return request
	case 29:
		request := new(QuickApply_Out)
		request.Decode(buffer)
		return request
	case 30:
		request := new(NotifyContribChange_Out)
		request.Decode(buffer)
		return request
	case 31:
		request := new(OtherClique_Out)
		request.Decode(buffer)
		return request
	}
	_ = buffer
	panic("unsupported response")
}

type CreateCliqueResult int8

const (
	CREATE_CLIQUE_RESULT_SUCCESS     CreateCliqueResult = 0
	CREATE_CLIQUE_RESULT_DUP_NAME    CreateCliqueResult = 1
	CREATE_CLIQUE_RESULT_HAVE_CLIQUE CreateCliqueResult = 2
)

type ApplyCliqueResult int8

const (
	APPLY_CLIQUE_RESULT_SUCCESS        ApplyCliqueResult = 0
	APPLY_CLIQUE_RESULT_ALREADY_JOIN   ApplyCliqueResult = 1
	APPLY_CLIQUE_RESULT_NOT_EXIST      ApplyCliqueResult = 2
	APPLY_CLIQUE_RESULT_REFUSE         ApplyCliqueResult = 3
	APPLY_CLIQUE_RESULT_TOO_MUCH_APPLY ApplyCliqueResult = 4
)

type CancelApplyCliqueResult int8

const (
	CANCEL_APPLY_CLIQUE_RESULT_SUCCESS   CancelApplyCliqueResult = 0
	CANCEL_APPLY_CLIQUE_RESULT_EXPIRED   CancelApplyCliqueResult = 1
	CANCEL_APPLY_CLIQUE_RESULT_NOT_EXIST CancelApplyCliqueResult = 2
)

type ProcessJoinApplyResult int8

const (
	PROCESS_JOIN_APPLY_RESULT_SUCCESS       ProcessJoinApplyResult = 0
	PROCESS_JOIN_APPLY_RESULT_EXPIRED       ProcessJoinApplyResult = 1
	PROCESS_JOIN_APPLY_RESULT_NO_ROOM       ProcessJoinApplyResult = 2
	PROCESS_JOIN_APPLY_RESULT_NO_PERMISSION ProcessJoinApplyResult = 3
	PROCESS_JOIN_APPLY_RESULT_CANCEL_APPLY  ProcessJoinApplyResult = 4
)

type MangeMemberAction int8

const (
	MANGE_MEMBER_ACTION_SET_OWNER  MangeMemberAction = 0
	MANGE_MEMBER_ACTION_SET_MANGER MangeMemberAction = 1
	MANGE_MEMBER_ACTION_SET_MEMBER MangeMemberAction = 2
	MANGE_MEMBER_ACTION_KICKOFF    MangeMemberAction = 3
)

type MangeMemberResult int8

const (
	MANGE_MEMBER_RESULT_SUCCESS       MangeMemberResult = 0
	MANGE_MEMBER_RESULT_NOT_EXIST     MangeMemberResult = 1
	MANGE_MEMBER_RESULT_NO_PERMISSION MangeMemberResult = 2
)

type CliqueOperaError int8

const (
	CLIQUE_OPERA_ERROR_SUCCESS          CliqueOperaError = 0
	CLIQUE_OPERA_ERROR_CLIQUE_NOT_EXIST CliqueOperaError = 1
	CLIQUE_OPERA_ERROR_NO_PERMISSION    CliqueOperaError = 2
	CLIQUE_OPERA_ERROR_MEMBER_NOT_EXIST CliqueOperaError = 3
	CLIQUE_OPERA_ERROR_ALREADY_JOIN     CliqueOperaError = 4
)

type NotifyLeaveCliqueReason int8

const (
	NOTIFY_LEAVE_CLIQUE_REASON_KICKOFF  NotifyLeaveCliqueReason = 0
	NOTIFY_LEAVE_CLIQUE_REASON_COLLAPSE NotifyLeaveCliqueReason = 1
	NOTIFY_LEAVE_CLIQUE_REASON_LEAVE    NotifyLeaveCliqueReason = 2
)

type NotifyJoincliqueFailedReason int8

const (
	NOTIFY_JOINCLIQUE_FAILED_REASON_REFUSE  NotifyJoincliqueFailedReason = 0
	NOTIFY_JOINCLIQUE_FAILED_REASON_EXPIRED NotifyJoincliqueFailedReason = 1
	NOTIFY_JOINCLIQUE_FAILED_REASON_NOROOM  NotifyJoincliqueFailedReason = 2
)

type CliqueRecuitmentResult int8

const (
	CLIQUE_RECUITMENT_RESULT_SUCCESS       CliqueRecuitmentResult = 0
	CLIQUE_RECUITMENT_RESULT_NO_PERMISSION CliqueRecuitmentResult = 1
	CLIQUE_RECUITMENT_RESULT_CD            CliqueRecuitmentResult = 2
)

type Player struct {
	PlayerId          int64  `json:"player_id"`
	Nickname          []byte `json:"nickname"`
	RoleId            int8   `json:"role_id"`
	AtX               int16  `json:"at_x"`
	AtY               int16  `json:"at_y"`
	FashionId         int16  `json:"fashion_id"`
	InMeditationState bool   `json:"in_meditation_state"`
	Level             int16  `json:"level"`
}

func (this *Player) Decode(buffer *net.Buffer) {
	this.PlayerId = int64(buffer.ReadUint64LE())
	this.Nickname = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.RoleId = int8(buffer.ReadUint8())
	this.AtX = int16(buffer.ReadUint16LE())
	this.AtY = int16(buffer.ReadUint16LE())
	this.FashionId = int16(buffer.ReadUint16LE())
	this.InMeditationState = buffer.ReadUint8() == 1
	this.Level = int16(buffer.ReadUint16LE())
}

func (this *Player) Encode(buffer *net.Buffer) {
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

func (this *Player) ByteSize() int {
	size := 20
	size += len(this.Nickname)
	return size
}

type CreateClique_In struct {
	Name     []byte `json:"name"`
	Announce []byte `json:"announce"`
}

func (this *CreateClique_In) Process(session *net.Session) {
	g_InHandler.CreateClique(session, this)
}

func (this *CreateClique_In) TypeName() string {
	return "clique.create_clique.in"
}

func (this *CreateClique_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 0
}

type CreateClique_Out struct {
	Result CreateCliqueResult `json:"result"`
}

func (this *CreateClique_Out) Process(session *net.Session) {
	g_OutHandler.CreateClique(session, this)
}

func (this *CreateClique_Out) TypeName() string {
	return "clique.create_clique.out"
}

func (this *CreateClique_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 0
}

func (this *CreateClique_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ApplyJoinClique_In struct {
	CliqueId int64 `json:"clique_id"`
}

func (this *ApplyJoinClique_In) Process(session *net.Session) {
	g_InHandler.ApplyJoinClique(session, this)
}

func (this *ApplyJoinClique_In) TypeName() string {
	return "clique.apply_join_clique.in"
}

func (this *ApplyJoinClique_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 1
}

type ApplyJoinClique_Out struct {
	CliqueId int64             `json:"clique_id"`
	Result   ApplyCliqueResult `json:"result"`
}

func (this *ApplyJoinClique_Out) Process(session *net.Session) {
	g_OutHandler.ApplyJoinClique(session, this)
}

func (this *ApplyJoinClique_Out) TypeName() string {
	return "clique.apply_join_clique.out"
}

func (this *ApplyJoinClique_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 1
}

func (this *ApplyJoinClique_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type CancelApplyClique_In struct {
	CliqueId int64 `json:"clique_id"`
}

func (this *CancelApplyClique_In) Process(session *net.Session) {
	g_InHandler.CancelApplyClique(session, this)
}

func (this *CancelApplyClique_In) TypeName() string {
	return "clique.cancel_apply_clique.in"
}

func (this *CancelApplyClique_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 2
}

type CancelApplyClique_Out struct {
	Result   CancelApplyCliqueResult `json:"result"`
	CliqueId int64                   `json:"clique_id"`
}

func (this *CancelApplyClique_Out) Process(session *net.Session) {
	g_OutHandler.CancelApplyClique(session, this)
}

func (this *CancelApplyClique_Out) TypeName() string {
	return "clique.cancel_apply_clique.out"
}

func (this *CancelApplyClique_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 2
}

func (this *CancelApplyClique_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ProcessJoinApply_In struct {
	Agree   bool                          `json:"agree"`
	Pidlist []ProcessJoinApply_In_Pidlist `json:"pidlist"`
}

type ProcessJoinApply_In_Pidlist struct {
	Pid int64 `json:"pid"`
}

func (this *ProcessJoinApply_In) Process(session *net.Session) {
	g_InHandler.ProcessJoinApply(session, this)
}

func (this *ProcessJoinApply_In) TypeName() string {
	return "clique.process_join_apply.in"
}

func (this *ProcessJoinApply_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 3
}

type ProcessJoinApply_Out struct {
	Applylist []ProcessJoinApply_Out_Applylist `json:"applylist"`
}

type ProcessJoinApply_Out_Applylist struct {
	Pid    int64                  `json:"pid"`
	Result ProcessJoinApplyResult `json:"result"`
}

func (this *ProcessJoinApply_Out) Process(session *net.Session) {
	g_OutHandler.ProcessJoinApply(session, this)
}

func (this *ProcessJoinApply_Out) TypeName() string {
	return "clique.process_join_apply.out"
}

func (this *ProcessJoinApply_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 3
}

func (this *ProcessJoinApply_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ElectOwner_In struct {
}

func (this *ElectOwner_In) Process(session *net.Session) {
	g_InHandler.ElectOwner(session, this)
}

func (this *ElectOwner_In) TypeName() string {
	return "clique.elect_owner.in"
}

func (this *ElectOwner_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 4
}

type ElectOwner_Out struct {
	Success bool `json:"success"`
}

func (this *ElectOwner_Out) Process(session *net.Session) {
	g_OutHandler.ElectOwner(session, this)
}

func (this *ElectOwner_Out) TypeName() string {
	return "clique.elect_owner.out"
}

func (this *ElectOwner_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 4
}

func (this *ElectOwner_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type MangeMember_In struct {
	Pid    int64             `json:"pid"`
	Action MangeMemberAction `json:"action"`
}

func (this *MangeMember_In) Process(session *net.Session) {
	g_InHandler.MangeMember(session, this)
}

func (this *MangeMember_In) TypeName() string {
	return "clique.mange_member.in"
}

func (this *MangeMember_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 5
}

type MangeMember_Out struct {
	Action MangeMemberAction `json:"action"`
	Pid    int64             `json:"pid"`
	Result MangeMemberResult `json:"result"`
}

func (this *MangeMember_Out) Process(session *net.Session) {
	g_OutHandler.MangeMember(session, this)
}

func (this *MangeMember_Out) TypeName() string {
	return "clique.mange_member.out"
}

func (this *MangeMember_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 5
}

func (this *MangeMember_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type DestoryClique_In struct {
}

func (this *DestoryClique_In) Process(session *net.Session) {
	g_InHandler.DestoryClique(session, this)
}

func (this *DestoryClique_In) TypeName() string {
	return "clique.destory_clique.in"
}

func (this *DestoryClique_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 6
}

type DestoryClique_Out struct {
}

func (this *DestoryClique_Out) Process(session *net.Session) {
	g_OutHandler.DestoryClique(session, this)
}

func (this *DestoryClique_Out) TypeName() string {
	return "clique.destory_clique.out"
}

func (this *DestoryClique_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 6
}

func (this *DestoryClique_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type UpdateAnnounce_In struct {
	Content []byte `json:"content"`
}

func (this *UpdateAnnounce_In) Process(session *net.Session) {
	g_InHandler.UpdateAnnounce(session, this)
}

func (this *UpdateAnnounce_In) TypeName() string {
	return "clique.update_announce.in"
}

func (this *UpdateAnnounce_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 7
}

type UpdateAnnounce_Out struct {
}

func (this *UpdateAnnounce_Out) Process(session *net.Session) {
	g_OutHandler.UpdateAnnounce(session, this)
}

func (this *UpdateAnnounce_Out) TypeName() string {
	return "clique.update_announce.out"
}

func (this *UpdateAnnounce_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 7
}

func (this *UpdateAnnounce_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type LeaveClique_In struct {
}

func (this *LeaveClique_In) Process(session *net.Session) {
	g_InHandler.LeaveClique(session, this)
}

func (this *LeaveClique_In) TypeName() string {
	return "clique.leave_clique.in"
}

func (this *LeaveClique_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 8
}

type LeaveClique_Out struct {
	Success bool `json:"success"`
}

func (this *LeaveClique_Out) Process(session *net.Session) {
	g_OutHandler.LeaveClique(session, this)
}

func (this *LeaveClique_Out) TypeName() string {
	return "clique.leave_clique.out"
}

func (this *LeaveClique_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 8
}

func (this *LeaveClique_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ListClique_In struct {
	Offset int16 `json:"offset"`
	Limit  int16 `json:"limit"`
}

func (this *ListClique_In) Process(session *net.Session) {
	g_InHandler.ListClique(session, this)
}

func (this *ListClique_In) TypeName() string {
	return "clique.list_clique.in"
}

func (this *ListClique_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 9
}

type ListClique_Out struct {
	AppliedCliques []ListClique_Out_AppliedCliques `json:"applied_cliques"`
	Total          int16                           `json:"total"`
	Cliques        []ListClique_Out_Cliques        `json:"cliques"`
}

type ListClique_Out_AppliedCliques struct {
	CliqueId int64 `json:"clique_id"`
}

type ListClique_Out_Cliques struct {
	Id        int64  `json:"id"`
	Name      []byte `json:"name"`
	Level     int16  `json:"level"`
	MemberNum int16  `json:"member_num"`
	OwnerNick []byte `json:"owner_nick"`
	OwnerPid  int64  `json:"owner_pid"`
	Announce  []byte `json:"announce"`
}

func (this *ListClique_Out) Process(session *net.Session) {
	g_OutHandler.ListClique(session, this)
}

func (this *ListClique_Out) TypeName() string {
	return "clique.list_clique.out"
}

func (this *ListClique_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 9
}

func (this *ListClique_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type CliquePublicInfo_In struct {
	CliqueId int64 `json:"clique_id"`
}

func (this *CliquePublicInfo_In) Process(session *net.Session) {
	g_InHandler.CliquePublicInfo(session, this)
}

func (this *CliquePublicInfo_In) TypeName() string {
	return "clique.clique_public_info.in"
}

func (this *CliquePublicInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 10
}

type CliquePublicInfo_Out struct {
	CliqueId             int64                                 `json:"clique_id"`
	Exist                bool                                  `json:"exist"`
	Name                 []byte                                `json:"name"`
	OwnerNick            []byte                                `json:"owner_nick"`
	OwnerPid             int64                                 `json:"owner_pid"`
	MemberNum            int16                                 `json:"member_num"`
	Level                int16                                 `json:"level"`
	Announce             []byte                                `json:"announce"`
	CenterBuildingLevel  int16                                 `json:"center_building_level"`
	TempleBuildingLevel  int16                                 `json:"temple_building_level"`
	BankBuildingLevel    int16                                 `json:"bank_building_level"`
	HealthBuildingLevel  int16                                 `json:"health_building_level"`
	AttackBuildingLevel  int16                                 `json:"attack_building_level"`
	DefenseBuildingLevel int16                                 `json:"defense_building_level"`
	AppliedCliques       []CliquePublicInfo_Out_AppliedCliques `json:"applied_cliques"`
}

type CliquePublicInfo_Out_AppliedCliques struct {
	CliqueId int64 `json:"clique_id"`
}

func (this *CliquePublicInfo_Out) Process(session *net.Session) {
	g_OutHandler.CliquePublicInfo(session, this)
}

func (this *CliquePublicInfo_Out) TypeName() string {
	return "clique.clique_public_info.out"
}

func (this *CliquePublicInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 10
}

func (this *CliquePublicInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type CliqueInfo_In struct {
}

func (this *CliqueInfo_In) Process(session *net.Session) {
	g_InHandler.CliqueInfo(session, this)
}

func (this *CliqueInfo_In) TypeName() string {
	return "clique.clique_info.in"
}

func (this *CliqueInfo_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 11
}

type CliqueInfo_Out struct {
	CliqueId             int64                    `json:"clique_id"`
	Name                 []byte                   `json:"name"`
	Announce             []byte                   `json:"announce"`
	TotalDonateCoins     int64                    `json:"total_donate_coins"`
	Contrib              int64                    `json:"contrib"`
	OwnerLoginTime       int64                    `json:"owner_login_time"`
	OwnerPid             int64                    `json:"owner_pid"`
	MangerPid1           int64                    `json:"manger_pid1"`
	MangerPid2           int64                    `json:"manger_pid2"`
	CenterBuildingCoins  int64                    `json:"center_building_coins"`
	TempleBuildingCoins  int64                    `json:"temple_building_coins"`
	BankBuildingCoins    int64                    `json:"bank_building_coins"`
	HealthBuildingCoins  int64                    `json:"health_building_coins"`
	AttackBuildingCoins  int64                    `json:"attack_building_coins"`
	DefenseBuildingCoins int64                    `json:"defense_building_coins"`
	CenterBuildingLevel  int16                    `json:"center_building_level"`
	TempleBuildingLevel  int16                    `json:"temple_building_level"`
	BankBuildingLevel    int16                    `json:"bank_building_level"`
	HealthBuildingLevel  int16                    `json:"health_building_level"`
	AttackBuildingLevel  int16                    `json:"attack_building_level"`
	DefenseBuildingLevel int16                    `json:"defense_building_level"`
	RecruitTimestamp     int64                    `json:"recruit_timestamp"`
	Members              []CliqueInfo_Out_Members `json:"members"`
}

type CliqueInfo_Out_Members struct {
	Pid     int64  `json:"pid"`
	RoleId  int8   `json:"role_id"`
	Level   int16  `json:"level"`
	Nick    []byte `json:"nick"`
	Contrib int64  `json:"contrib"`
}

func (this *CliqueInfo_Out) Process(session *net.Session) {
	g_OutHandler.CliqueInfo(session, this)
}

func (this *CliqueInfo_Out) TypeName() string {
	return "clique.clique_info.out"
}

func (this *CliqueInfo_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 11
}

func (this *CliqueInfo_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ListApply_In struct {
	Offset int16 `json:"offset"`
	Limit  int16 `json:"limit"`
}

func (this *ListApply_In) Process(session *net.Session) {
	g_InHandler.ListApply(session, this)
}

func (this *ListApply_In) TypeName() string {
	return "clique.list_apply.in"
}

func (this *ListApply_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 12
}

type ListApply_Out struct {
	AutoAudit bool                    `json:"auto_audit"`
	Level     int16                   `json:"level"`
	Players   []ListApply_Out_Players `json:"players"`
}

type ListApply_Out_Players struct {
	Pid       int64  `json:"pid"`
	Nick      []byte `json:"nick"`
	Level     int16  `json:"level"`
	ArenaRank int64  `json:"arena_rank"`
	Timestamp int64  `json:"timestamp"`
}

func (this *ListApply_Out) Process(session *net.Session) {
	g_OutHandler.ListApply(session, this)
}

func (this *ListApply_Out) TypeName() string {
	return "clique.list_apply.out"
}

func (this *ListApply_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 12
}

func (this *ListApply_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type OperaErrorNotify_Out struct {
	Resutl CliqueOperaError `json:"resutl"`
}

func (this *OperaErrorNotify_Out) Process(session *net.Session) {
	g_OutHandler.OperaErrorNotify(session, this)
}

func (this *OperaErrorNotify_Out) TypeName() string {
	return "clique.opera_error_notify.out"
}

func (this *OperaErrorNotify_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 13
}

func (this *OperaErrorNotify_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type EnterClubhouse_In struct {
}

func (this *EnterClubhouse_In) Process(session *net.Session) {
	g_InHandler.EnterClubhouse(session, this)
}

func (this *EnterClubhouse_In) TypeName() string {
	return "clique.enter_clubhouse.in"
}

func (this *EnterClubhouse_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 14
}

type EnterClubhouse_Out struct {
	Ok     bool   `json:"ok"`
	Player Player `json:"player"`
}

func (this *EnterClubhouse_Out) Process(session *net.Session) {
	g_OutHandler.EnterClubhouse(session, this)
}

func (this *EnterClubhouse_Out) TypeName() string {
	return "clique.enter_clubhouse.out"
}

func (this *EnterClubhouse_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 14
}

func (this *EnterClubhouse_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type LeaveClubhouse_In struct {
}

func (this *LeaveClubhouse_In) Process(session *net.Session) {
	g_InHandler.LeaveClubhouse(session, this)
}

func (this *LeaveClubhouse_In) TypeName() string {
	return "clique.leave_clubhouse.in"
}

func (this *LeaveClubhouse_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 15
}

type LeaveClubhouse_Out struct {
	PlayerId int64 `json:"player_id"`
}

func (this *LeaveClubhouse_Out) Process(session *net.Session) {
	g_OutHandler.LeaveClubhouse(session, this)
}

func (this *LeaveClubhouse_Out) TypeName() string {
	return "clique.leave_clubhouse.out"
}

func (this *LeaveClubhouse_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 15
}

func (this *LeaveClubhouse_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type ClubMove_In struct {
	ToX int16 `json:"to_x"`
	ToY int16 `json:"to_y"`
}

func (this *ClubMove_In) Process(session *net.Session) {
	g_InHandler.ClubMove(session, this)
}

func (this *ClubMove_In) TypeName() string {
	return "clique.club_move.in"
}

func (this *ClubMove_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 16
}

type ClubMove_Out struct {
	PlayerId int64 `json:"player_id"`
	ToX      int16 `json:"to_x"`
	ToY      int16 `json:"to_y"`
}

func (this *ClubMove_Out) Process(session *net.Session) {
	g_OutHandler.ClubMove(session, this)
}

func (this *ClubMove_Out) TypeName() string {
	return "clique.club_move.out"
}

func (this *ClubMove_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 16
}

func (this *ClubMove_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type NotifyClubhousePlayers_Out struct {
	Players []NotifyClubhousePlayers_Out_Players `json:"players"`
}

type NotifyClubhousePlayers_Out_Players struct {
	Player Player `json:"player"`
}

func (this *NotifyClubhousePlayers_Out) Process(session *net.Session) {
	g_OutHandler.NotifyClubhousePlayers(session, this)
}

func (this *NotifyClubhousePlayers_Out) TypeName() string {
	return "clique.notify_clubhouse_players.out"
}

func (this *NotifyClubhousePlayers_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 17
}

func (this *NotifyClubhousePlayers_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type NotifyNewPlayer_Out struct {
	Player Player `json:"player"`
}

func (this *NotifyNewPlayer_Out) Process(session *net.Session) {
	g_OutHandler.NotifyNewPlayer(session, this)
}

func (this *NotifyNewPlayer_Out) TypeName() string {
	return "clique.notify_new_player.out"
}

func (this *NotifyNewPlayer_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 18
}

func (this *NotifyNewPlayer_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type NotifyUpdatePlayer_Out struct {
	PlayerId          int64 `json:"player_id"`
	FashionId         int16 `json:"fashion_id"`
	InMeditationState bool  `json:"in_meditation_state"`
}

func (this *NotifyUpdatePlayer_Out) Process(session *net.Session) {
	g_OutHandler.NotifyUpdatePlayer(session, this)
}

func (this *NotifyUpdatePlayer_Out) TypeName() string {
	return "clique.notify_update_player.out"
}

func (this *NotifyUpdatePlayer_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 19
}

func (this *NotifyUpdatePlayer_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type CliquePublicInfoSummary_In struct {
	CliqueId int64 `json:"clique_id"`
}

func (this *CliquePublicInfoSummary_In) Process(session *net.Session) {
	g_InHandler.CliquePublicInfoSummary(session, this)
}

func (this *CliquePublicInfoSummary_In) TypeName() string {
	return "clique.clique_public_info_summary.in"
}

func (this *CliquePublicInfoSummary_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 20
}

type CliquePublicInfoSummary_Out struct {
	CliqueId  int64                                 `json:"CliqueId"`
	Name      []byte                                `json:"name"`
	Level     int16                                 `json:"level"`
	MemberNum int16                                 `json:"member_num"`
	OwnerNick []byte                                `json:"owner_nick"`
	OwnerPid  int64                                 `json:"owner_pid"`
	Announce  []byte                                `json:"announce"`
	Cliques   []CliquePublicInfoSummary_Out_Cliques `json:"cliques"`
}

type CliquePublicInfoSummary_Out_Cliques struct {
	CliqueId int64 `json:"clique_id"`
}

func (this *CliquePublicInfoSummary_Out) Process(session *net.Session) {
	g_OutHandler.CliquePublicInfoSummary(session, this)
}

func (this *CliquePublicInfoSummary_Out) TypeName() string {
	return "clique.clique_public_info_summary.out"
}

func (this *CliquePublicInfoSummary_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 20
}

func (this *CliquePublicInfoSummary_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type CliqueAutoAudit_In struct {
	Level  int16 `json:"level"`
	Enable bool  `json:"enable"`
}

func (this *CliqueAutoAudit_In) Process(session *net.Session) {
	g_InHandler.CliqueAutoAudit(session, this)
}

func (this *CliqueAutoAudit_In) TypeName() string {
	return "clique.clique_auto_audit.in"
}

func (this *CliqueAutoAudit_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 21
}

type CliqueAutoAudit_Out struct {
}

func (this *CliqueAutoAudit_Out) Process(session *net.Session) {
	g_OutHandler.CliqueAutoAudit(session, this)
}

func (this *CliqueAutoAudit_Out) TypeName() string {
	return "clique.clique_auto_audit.out"
}

func (this *CliqueAutoAudit_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 21
}

func (this *CliqueAutoAudit_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type CliqueBaseDonate_In struct {
}

func (this *CliqueBaseDonate_In) Process(session *net.Session) {
	g_InHandler.CliqueBaseDonate(session, this)
}

func (this *CliqueBaseDonate_In) TypeName() string {
	return "clique.clique_base_donate.in"
}

func (this *CliqueBaseDonate_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 22
}

type CliqueBaseDonate_Out struct {
}

func (this *CliqueBaseDonate_Out) Process(session *net.Session) {
	g_OutHandler.CliqueBaseDonate(session, this)
}

func (this *CliqueBaseDonate_Out) TypeName() string {
	return "clique.clique_base_donate.out"
}

func (this *CliqueBaseDonate_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 22
}

func (this *CliqueBaseDonate_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type NotifyLeaveClique_Out struct {
	Pid    int64                   `json:"pid"`
	Reason NotifyLeaveCliqueReason `json:"reason"`
}

func (this *NotifyLeaveClique_Out) Process(session *net.Session) {
	g_OutHandler.NotifyLeaveClique(session, this)
}

func (this *NotifyLeaveClique_Out) TypeName() string {
	return "clique.notify_leave_clique.out"
}

func (this *NotifyLeaveClique_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 23
}

func (this *NotifyLeaveClique_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type NotifyJoincliqueSuccess_Out struct {
	Pidlist []NotifyJoincliqueSuccess_Out_Pidlist `json:"pidlist"`
}

type NotifyJoincliqueSuccess_Out_Pidlist struct {
	Pid    int64  `json:"pid"`
	RoleId int8   `json:"role_id"`
	Level  int16  `json:"level"`
	Nick   []byte `json:"nick"`
}

func (this *NotifyJoincliqueSuccess_Out) Process(session *net.Session) {
	g_OutHandler.NotifyJoincliqueSuccess(session, this)
}

func (this *NotifyJoincliqueSuccess_Out) TypeName() string {
	return "clique.notify_joinclique_success.out"
}

func (this *NotifyJoincliqueSuccess_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 24
}

func (this *NotifyJoincliqueSuccess_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type NotifyCliqueMangerAction_Out struct {
	Actiontype MangeMemberAction `json:"actiontype"`
	Pid        int64             `json:"pid"`
}

func (this *NotifyCliqueMangerAction_Out) Process(session *net.Session) {
	g_OutHandler.NotifyCliqueMangerAction(session, this)
}

func (this *NotifyCliqueMangerAction_Out) TypeName() string {
	return "clique.notify_clique_manger_action.out"
}

func (this *NotifyCliqueMangerAction_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 25
}

func (this *NotifyCliqueMangerAction_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type CliqueRecruitment_In struct {
}

func (this *CliqueRecruitment_In) Process(session *net.Session) {
	g_InHandler.CliqueRecruitment(session, this)
}

func (this *CliqueRecruitment_In) TypeName() string {
	return "clique.clique_recruitment.in"
}

func (this *CliqueRecruitment_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 26
}

type CliqueRecruitment_Out struct {
	Result    CliqueRecuitmentResult `json:"result"`
	Timestamp int64                  `json:"timestamp"`
}

func (this *CliqueRecruitment_Out) Process(session *net.Session) {
	g_OutHandler.CliqueRecruitment(session, this)
}

func (this *CliqueRecruitment_Out) TypeName() string {
	return "clique.clique_recruitment.out"
}

func (this *CliqueRecruitment_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 26
}

func (this *CliqueRecruitment_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type NotifyCliqueAnnounce_Out struct {
	Announce []byte `json:"announce"`
}

func (this *NotifyCliqueAnnounce_Out) Process(session *net.Session) {
	g_OutHandler.NotifyCliqueAnnounce(session, this)
}

func (this *NotifyCliqueAnnounce_Out) TypeName() string {
	return "clique.notify_clique_announce.out"
}

func (this *NotifyCliqueAnnounce_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 27
}

func (this *NotifyCliqueAnnounce_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type NotifyCliqueElectowner_Out struct {
	Ownerid int64 `json:"ownerid"`
}

func (this *NotifyCliqueElectowner_Out) Process(session *net.Session) {
	g_OutHandler.NotifyCliqueElectowner(session, this)
}

func (this *NotifyCliqueElectowner_Out) TypeName() string {
	return "clique.notify_clique_electowner.out"
}

func (this *NotifyCliqueElectowner_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 28
}

func (this *NotifyCliqueElectowner_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type QuickApply_In struct {
}

func (this *QuickApply_In) Process(session *net.Session) {
	g_InHandler.QuickApply(session, this)
}

func (this *QuickApply_In) TypeName() string {
	return "clique.quick_apply.in"
}

func (this *QuickApply_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 29
}

type QuickApply_Out struct {
	Success bool `json:"success"`
}

func (this *QuickApply_Out) Process(session *net.Session) {
	g_OutHandler.QuickApply(session, this)
}

func (this *QuickApply_Out) TypeName() string {
	return "clique.quick_apply.out"
}

func (this *QuickApply_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 29
}

func (this *QuickApply_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type NotifyContribChange_Out struct {
	Value int64 `json:"value"`
}

func (this *NotifyContribChange_Out) Process(session *net.Session) {
	g_OutHandler.NotifyContribChange(session, this)
}

func (this *NotifyContribChange_Out) TypeName() string {
	return "clique.notify_contrib_change.out"
}

func (this *NotifyContribChange_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 30
}

func (this *NotifyContribChange_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

type OtherClique_In struct {
	Page int16 `json:"page"`
}

func (this *OtherClique_In) Process(session *net.Session) {
	g_InHandler.OtherClique(session, this)
}

func (this *OtherClique_In) TypeName() string {
	return "clique.other_clique.in"
}

func (this *OtherClique_In) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 31
}

type OtherClique_Out struct {
	TotalNum   int16                        `json:"total_num"`
	CliqueList []OtherClique_Out_CliqueList `json:"clique_list"`
}

type OtherClique_Out_CliqueList struct {
	Rank            int16  `json:"rank"`
	Name            []byte `json:"name"`
	CliqueId        int64  `json:"clique_id"`
	CliqueLevel     int16  `json:"clique_level"`
	OwnerName       []byte `json:"owner_name"`
	OwnerPid        int64  `json:"owner_pid"`
	CliqueMemberNum int16  `json:"clique_member_num"`
}

func (this *OtherClique_Out) Process(session *net.Session) {
	g_OutHandler.OtherClique(session, this)
}

func (this *OtherClique_Out) TypeName() string {
	return "clique.other_clique.out"
}

func (this *OtherClique_Out) GetModuleIdAndActionId() (moduleId, actionId int8) {
	return 33, 31
}

func (this *OtherClique_Out) Bytes() []byte {
	data := make([]byte, this.ByteSize()+2)
	this.Encode(net.NewBuffer(data))
	return data
}

func (this *CreateClique_In) Decode(buffer *net.Buffer) {
	this.Name = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Announce = buffer.ReadBytes(int(buffer.ReadUint16LE()))
}

func (this *CreateClique_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(0)
	buffer.WriteUint16LE(uint16(len(this.Name)))
	buffer.WriteBytes(this.Name)
	buffer.WriteUint16LE(uint16(len(this.Announce)))
	buffer.WriteBytes(this.Announce)
}

func (this *CreateClique_In) ByteSize() int {
	size := 6
	size += len(this.Name)
	size += len(this.Announce)
	return size
}

func (this *CreateClique_Out) Decode(buffer *net.Buffer) {
	this.Result = CreateCliqueResult(buffer.ReadUint8())
}

func (this *CreateClique_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(0)
	buffer.WriteUint8(uint8(this.Result))
}

func (this *CreateClique_Out) ByteSize() int {
	size := 3
	return size
}

func (this *ApplyJoinClique_In) Decode(buffer *net.Buffer) {
	this.CliqueId = int64(buffer.ReadUint64LE())
}

func (this *ApplyJoinClique_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(1)
	buffer.WriteUint64LE(uint64(this.CliqueId))
}

func (this *ApplyJoinClique_In) ByteSize() int {
	size := 10
	return size
}

func (this *ApplyJoinClique_Out) Decode(buffer *net.Buffer) {
	this.CliqueId = int64(buffer.ReadUint64LE())
	this.Result = ApplyCliqueResult(buffer.ReadUint8())
}

func (this *ApplyJoinClique_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(1)
	buffer.WriteUint64LE(uint64(this.CliqueId))
	buffer.WriteUint8(uint8(this.Result))
}

func (this *ApplyJoinClique_Out) ByteSize() int {
	size := 11
	return size
}

func (this *CancelApplyClique_In) Decode(buffer *net.Buffer) {
	this.CliqueId = int64(buffer.ReadUint64LE())
}

func (this *CancelApplyClique_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(2)
	buffer.WriteUint64LE(uint64(this.CliqueId))
}

func (this *CancelApplyClique_In) ByteSize() int {
	size := 10
	return size
}

func (this *CancelApplyClique_Out) Decode(buffer *net.Buffer) {
	this.Result = CancelApplyCliqueResult(buffer.ReadUint8())
	this.CliqueId = int64(buffer.ReadUint64LE())
}

func (this *CancelApplyClique_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(2)
	buffer.WriteUint8(uint8(this.Result))
	buffer.WriteUint64LE(uint64(this.CliqueId))
}

func (this *CancelApplyClique_Out) ByteSize() int {
	size := 11
	return size
}

func (this *ProcessJoinApply_In) Decode(buffer *net.Buffer) {
	this.Agree = buffer.ReadUint8() == 1
	this.Pidlist = make([]ProcessJoinApply_In_Pidlist, buffer.ReadUint8())
	for i := 0; i < len(this.Pidlist); i++ {
		this.Pidlist[i].Decode(buffer)
	}
}

func (this *ProcessJoinApply_In_Pidlist) Decode(buffer *net.Buffer) {
	this.Pid = int64(buffer.ReadUint64LE())
}

func (this *ProcessJoinApply_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(3)
	if this.Agree {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
	buffer.WriteUint8(uint8(len(this.Pidlist)))
	for i := 0; i < len(this.Pidlist); i++ {
		this.Pidlist[i].Encode(buffer)
	}
}

func (this *ProcessJoinApply_In_Pidlist) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.Pid))
}

func (this *ProcessJoinApply_In) ByteSize() int {
	size := 4
	size += len(this.Pidlist) * 8
	return size
}

func (this *ProcessJoinApply_Out) Decode(buffer *net.Buffer) {
	this.Applylist = make([]ProcessJoinApply_Out_Applylist, buffer.ReadUint8())
	for i := 0; i < len(this.Applylist); i++ {
		this.Applylist[i].Decode(buffer)
	}
}

func (this *ProcessJoinApply_Out_Applylist) Decode(buffer *net.Buffer) {
	this.Pid = int64(buffer.ReadUint64LE())
	this.Result = ProcessJoinApplyResult(buffer.ReadUint8())
}

func (this *ProcessJoinApply_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(3)
	buffer.WriteUint8(uint8(len(this.Applylist)))
	for i := 0; i < len(this.Applylist); i++ {
		this.Applylist[i].Encode(buffer)
	}
}

func (this *ProcessJoinApply_Out_Applylist) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.Pid))
	buffer.WriteUint8(uint8(this.Result))
}

func (this *ProcessJoinApply_Out) ByteSize() int {
	size := 3
	size += len(this.Applylist) * 9
	return size
}

func (this *ElectOwner_In) Decode(buffer *net.Buffer) {
}

func (this *ElectOwner_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(4)
}

func (this *ElectOwner_In) ByteSize() int {
	size := 2
	return size
}

func (this *ElectOwner_Out) Decode(buffer *net.Buffer) {
	this.Success = buffer.ReadUint8() == 1
}

func (this *ElectOwner_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(4)
	if this.Success {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *ElectOwner_Out) ByteSize() int {
	size := 3
	return size
}

func (this *MangeMember_In) Decode(buffer *net.Buffer) {
	this.Pid = int64(buffer.ReadUint64LE())
	this.Action = MangeMemberAction(buffer.ReadUint8())
}

func (this *MangeMember_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(5)
	buffer.WriteUint64LE(uint64(this.Pid))
	buffer.WriteUint8(uint8(this.Action))
}

func (this *MangeMember_In) ByteSize() int {
	size := 11
	return size
}

func (this *MangeMember_Out) Decode(buffer *net.Buffer) {
	this.Action = MangeMemberAction(buffer.ReadUint8())
	this.Pid = int64(buffer.ReadUint64LE())
	this.Result = MangeMemberResult(buffer.ReadUint8())
}

func (this *MangeMember_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(5)
	buffer.WriteUint8(uint8(this.Action))
	buffer.WriteUint64LE(uint64(this.Pid))
	buffer.WriteUint8(uint8(this.Result))
}

func (this *MangeMember_Out) ByteSize() int {
	size := 12
	return size
}

func (this *DestoryClique_In) Decode(buffer *net.Buffer) {
}

func (this *DestoryClique_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(6)
}

func (this *DestoryClique_In) ByteSize() int {
	size := 2
	return size
}

func (this *DestoryClique_Out) Decode(buffer *net.Buffer) {
}

func (this *DestoryClique_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(6)
}

func (this *DestoryClique_Out) ByteSize() int {
	size := 2
	return size
}

func (this *UpdateAnnounce_In) Decode(buffer *net.Buffer) {
	this.Content = buffer.ReadBytes(int(buffer.ReadUint16LE()))
}

func (this *UpdateAnnounce_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(7)
	buffer.WriteUint16LE(uint16(len(this.Content)))
	buffer.WriteBytes(this.Content)
}

func (this *UpdateAnnounce_In) ByteSize() int {
	size := 4
	size += len(this.Content)
	return size
}

func (this *UpdateAnnounce_Out) Decode(buffer *net.Buffer) {
}

func (this *UpdateAnnounce_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(7)
}

func (this *UpdateAnnounce_Out) ByteSize() int {
	size := 2
	return size
}

func (this *LeaveClique_In) Decode(buffer *net.Buffer) {
}

func (this *LeaveClique_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(8)
}

func (this *LeaveClique_In) ByteSize() int {
	size := 2
	return size
}

func (this *LeaveClique_Out) Decode(buffer *net.Buffer) {
	this.Success = buffer.ReadUint8() == 1
}

func (this *LeaveClique_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(8)
	if this.Success {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *LeaveClique_Out) ByteSize() int {
	size := 3
	return size
}

func (this *ListClique_In) Decode(buffer *net.Buffer) {
	this.Offset = int16(buffer.ReadUint16LE())
	this.Limit = int16(buffer.ReadUint16LE())
}

func (this *ListClique_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(9)
	buffer.WriteUint16LE(uint16(this.Offset))
	buffer.WriteUint16LE(uint16(this.Limit))
}

func (this *ListClique_In) ByteSize() int {
	size := 6
	return size
}

func (this *ListClique_Out) Decode(buffer *net.Buffer) {
	this.AppliedCliques = make([]ListClique_Out_AppliedCliques, buffer.ReadUint8())
	for i := 0; i < len(this.AppliedCliques); i++ {
		this.AppliedCliques[i].Decode(buffer)
	}
	this.Total = int16(buffer.ReadUint16LE())
	this.Cliques = make([]ListClique_Out_Cliques, buffer.ReadUint16LE())
	for i := 0; i < len(this.Cliques); i++ {
		this.Cliques[i].Decode(buffer)
	}
}

func (this *ListClique_Out_AppliedCliques) Decode(buffer *net.Buffer) {
	this.CliqueId = int64(buffer.ReadUint64LE())
}

func (this *ListClique_Out_Cliques) Decode(buffer *net.Buffer) {
	this.Id = int64(buffer.ReadUint64LE())
	this.Name = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Level = int16(buffer.ReadUint16LE())
	this.MemberNum = int16(buffer.ReadUint16LE())
	this.OwnerNick = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.OwnerPid = int64(buffer.ReadUint64LE())
	this.Announce = buffer.ReadBytes(int(buffer.ReadUint16LE()))
}

func (this *ListClique_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(9)
	buffer.WriteUint8(uint8(len(this.AppliedCliques)))
	for i := 0; i < len(this.AppliedCliques); i++ {
		this.AppliedCliques[i].Encode(buffer)
	}
	buffer.WriteUint16LE(uint16(this.Total))
	buffer.WriteUint16LE(uint16(len(this.Cliques)))
	for i := 0; i < len(this.Cliques); i++ {
		this.Cliques[i].Encode(buffer)
	}
}

func (this *ListClique_Out_AppliedCliques) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.CliqueId))
}

func (this *ListClique_Out_Cliques) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.Id))
	buffer.WriteUint16LE(uint16(len(this.Name)))
	buffer.WriteBytes(this.Name)
	buffer.WriteUint16LE(uint16(this.Level))
	buffer.WriteUint16LE(uint16(this.MemberNum))
	buffer.WriteUint16LE(uint16(len(this.OwnerNick)))
	buffer.WriteBytes(this.OwnerNick)
	buffer.WriteUint64LE(uint64(this.OwnerPid))
	buffer.WriteUint16LE(uint16(len(this.Announce)))
	buffer.WriteBytes(this.Announce)
}

func (this *ListClique_Out) ByteSize() int {
	size := 7
	size += len(this.AppliedCliques) * 8
	for i := 0; i < len(this.Cliques); i++ {
		size += this.Cliques[i].ByteSize()
	}
	return size
}

func (this *ListClique_Out_Cliques) ByteSize() int {
	size := 26
	size += len(this.Name)
	size += len(this.OwnerNick)
	size += len(this.Announce)
	return size
}

func (this *CliquePublicInfo_In) Decode(buffer *net.Buffer) {
	this.CliqueId = int64(buffer.ReadUint64LE())
}

func (this *CliquePublicInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(10)
	buffer.WriteUint64LE(uint64(this.CliqueId))
}

func (this *CliquePublicInfo_In) ByteSize() int {
	size := 10
	return size
}

func (this *CliquePublicInfo_Out) Decode(buffer *net.Buffer) {
	this.CliqueId = int64(buffer.ReadUint64LE())
	this.Exist = buffer.ReadUint8() == 1
	this.Name = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.OwnerNick = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.OwnerPid = int64(buffer.ReadUint64LE())
	this.MemberNum = int16(buffer.ReadUint16LE())
	this.Level = int16(buffer.ReadUint16LE())
	this.Announce = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.CenterBuildingLevel = int16(buffer.ReadUint16LE())
	this.TempleBuildingLevel = int16(buffer.ReadUint16LE())
	this.BankBuildingLevel = int16(buffer.ReadUint16LE())
	this.HealthBuildingLevel = int16(buffer.ReadUint16LE())
	this.AttackBuildingLevel = int16(buffer.ReadUint16LE())
	this.DefenseBuildingLevel = int16(buffer.ReadUint16LE())
	this.AppliedCliques = make([]CliquePublicInfo_Out_AppliedCliques, buffer.ReadUint8())
	for i := 0; i < len(this.AppliedCliques); i++ {
		this.AppliedCliques[i].Decode(buffer)
	}
}

func (this *CliquePublicInfo_Out_AppliedCliques) Decode(buffer *net.Buffer) {
	this.CliqueId = int64(buffer.ReadUint64LE())
}

func (this *CliquePublicInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(10)
	buffer.WriteUint64LE(uint64(this.CliqueId))
	if this.Exist {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
	buffer.WriteUint16LE(uint16(len(this.Name)))
	buffer.WriteBytes(this.Name)
	buffer.WriteUint16LE(uint16(len(this.OwnerNick)))
	buffer.WriteBytes(this.OwnerNick)
	buffer.WriteUint64LE(uint64(this.OwnerPid))
	buffer.WriteUint16LE(uint16(this.MemberNum))
	buffer.WriteUint16LE(uint16(this.Level))
	buffer.WriteUint16LE(uint16(len(this.Announce)))
	buffer.WriteBytes(this.Announce)
	buffer.WriteUint16LE(uint16(this.CenterBuildingLevel))
	buffer.WriteUint16LE(uint16(this.TempleBuildingLevel))
	buffer.WriteUint16LE(uint16(this.BankBuildingLevel))
	buffer.WriteUint16LE(uint16(this.HealthBuildingLevel))
	buffer.WriteUint16LE(uint16(this.AttackBuildingLevel))
	buffer.WriteUint16LE(uint16(this.DefenseBuildingLevel))
	buffer.WriteUint8(uint8(len(this.AppliedCliques)))
	for i := 0; i < len(this.AppliedCliques); i++ {
		this.AppliedCliques[i].Encode(buffer)
	}
}

func (this *CliquePublicInfo_Out_AppliedCliques) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.CliqueId))
}

func (this *CliquePublicInfo_Out) ByteSize() int {
	size := 42
	size += len(this.Name)
	size += len(this.OwnerNick)
	size += len(this.Announce)
	size += len(this.AppliedCliques) * 8
	return size
}

func (this *CliqueInfo_In) Decode(buffer *net.Buffer) {
}

func (this *CliqueInfo_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(11)
}

func (this *CliqueInfo_In) ByteSize() int {
	size := 2
	return size
}

func (this *CliqueInfo_Out) Decode(buffer *net.Buffer) {
	this.CliqueId = int64(buffer.ReadUint64LE())
	this.Name = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Announce = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.TotalDonateCoins = int64(buffer.ReadUint64LE())
	this.Contrib = int64(buffer.ReadUint64LE())
	this.OwnerLoginTime = int64(buffer.ReadUint64LE())
	this.OwnerPid = int64(buffer.ReadUint64LE())
	this.MangerPid1 = int64(buffer.ReadUint64LE())
	this.MangerPid2 = int64(buffer.ReadUint64LE())
	this.CenterBuildingCoins = int64(buffer.ReadUint64LE())
	this.TempleBuildingCoins = int64(buffer.ReadUint64LE())
	this.BankBuildingCoins = int64(buffer.ReadUint64LE())
	this.HealthBuildingCoins = int64(buffer.ReadUint64LE())
	this.AttackBuildingCoins = int64(buffer.ReadUint64LE())
	this.DefenseBuildingCoins = int64(buffer.ReadUint64LE())
	this.CenterBuildingLevel = int16(buffer.ReadUint16LE())
	this.TempleBuildingLevel = int16(buffer.ReadUint16LE())
	this.BankBuildingLevel = int16(buffer.ReadUint16LE())
	this.HealthBuildingLevel = int16(buffer.ReadUint16LE())
	this.AttackBuildingLevel = int16(buffer.ReadUint16LE())
	this.DefenseBuildingLevel = int16(buffer.ReadUint16LE())
	this.RecruitTimestamp = int64(buffer.ReadUint64LE())
	this.Members = make([]CliqueInfo_Out_Members, buffer.ReadUint16LE())
	for i := 0; i < len(this.Members); i++ {
		this.Members[i].Decode(buffer)
	}
}

func (this *CliqueInfo_Out_Members) Decode(buffer *net.Buffer) {
	this.Pid = int64(buffer.ReadUint64LE())
	this.RoleId = int8(buffer.ReadUint8())
	this.Level = int16(buffer.ReadUint16LE())
	this.Nick = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Contrib = int64(buffer.ReadUint64LE())
}

func (this *CliqueInfo_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(11)
	buffer.WriteUint64LE(uint64(this.CliqueId))
	buffer.WriteUint16LE(uint16(len(this.Name)))
	buffer.WriteBytes(this.Name)
	buffer.WriteUint16LE(uint16(len(this.Announce)))
	buffer.WriteBytes(this.Announce)
	buffer.WriteUint64LE(uint64(this.TotalDonateCoins))
	buffer.WriteUint64LE(uint64(this.Contrib))
	buffer.WriteUint64LE(uint64(this.OwnerLoginTime))
	buffer.WriteUint64LE(uint64(this.OwnerPid))
	buffer.WriteUint64LE(uint64(this.MangerPid1))
	buffer.WriteUint64LE(uint64(this.MangerPid2))
	buffer.WriteUint64LE(uint64(this.CenterBuildingCoins))
	buffer.WriteUint64LE(uint64(this.TempleBuildingCoins))
	buffer.WriteUint64LE(uint64(this.BankBuildingCoins))
	buffer.WriteUint64LE(uint64(this.HealthBuildingCoins))
	buffer.WriteUint64LE(uint64(this.AttackBuildingCoins))
	buffer.WriteUint64LE(uint64(this.DefenseBuildingCoins))
	buffer.WriteUint16LE(uint16(this.CenterBuildingLevel))
	buffer.WriteUint16LE(uint16(this.TempleBuildingLevel))
	buffer.WriteUint16LE(uint16(this.BankBuildingLevel))
	buffer.WriteUint16LE(uint16(this.HealthBuildingLevel))
	buffer.WriteUint16LE(uint16(this.AttackBuildingLevel))
	buffer.WriteUint16LE(uint16(this.DefenseBuildingLevel))
	buffer.WriteUint64LE(uint64(this.RecruitTimestamp))
	buffer.WriteUint16LE(uint16(len(this.Members)))
	for i := 0; i < len(this.Members); i++ {
		this.Members[i].Encode(buffer)
	}
}

func (this *CliqueInfo_Out_Members) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.Pid))
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint16LE(uint16(this.Level))
	buffer.WriteUint16LE(uint16(len(this.Nick)))
	buffer.WriteBytes(this.Nick)
	buffer.WriteUint64LE(uint64(this.Contrib))
}

func (this *CliqueInfo_Out) ByteSize() int {
	size := 132
	size += len(this.Name)
	size += len(this.Announce)
	for i := 0; i < len(this.Members); i++ {
		size += this.Members[i].ByteSize()
	}
	return size
}

func (this *CliqueInfo_Out_Members) ByteSize() int {
	size := 21
	size += len(this.Nick)
	return size
}

func (this *ListApply_In) Decode(buffer *net.Buffer) {
	this.Offset = int16(buffer.ReadUint16LE())
	this.Limit = int16(buffer.ReadUint16LE())
}

func (this *ListApply_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(12)
	buffer.WriteUint16LE(uint16(this.Offset))
	buffer.WriteUint16LE(uint16(this.Limit))
}

func (this *ListApply_In) ByteSize() int {
	size := 6
	return size
}

func (this *ListApply_Out) Decode(buffer *net.Buffer) {
	this.AutoAudit = buffer.ReadUint8() == 1
	this.Level = int16(buffer.ReadUint16LE())
	this.Players = make([]ListApply_Out_Players, buffer.ReadUint16LE())
	for i := 0; i < len(this.Players); i++ {
		this.Players[i].Decode(buffer)
	}
}

func (this *ListApply_Out_Players) Decode(buffer *net.Buffer) {
	this.Pid = int64(buffer.ReadUint64LE())
	this.Nick = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Level = int16(buffer.ReadUint16LE())
	this.ArenaRank = int64(buffer.ReadUint64LE())
	this.Timestamp = int64(buffer.ReadUint64LE())
}

func (this *ListApply_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(12)
	if this.AutoAudit {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
	buffer.WriteUint16LE(uint16(this.Level))
	buffer.WriteUint16LE(uint16(len(this.Players)))
	for i := 0; i < len(this.Players); i++ {
		this.Players[i].Encode(buffer)
	}
}

func (this *ListApply_Out_Players) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.Pid))
	buffer.WriteUint16LE(uint16(len(this.Nick)))
	buffer.WriteBytes(this.Nick)
	buffer.WriteUint16LE(uint16(this.Level))
	buffer.WriteUint64LE(uint64(this.ArenaRank))
	buffer.WriteUint64LE(uint64(this.Timestamp))
}

func (this *ListApply_Out) ByteSize() int {
	size := 7
	for i := 0; i < len(this.Players); i++ {
		size += this.Players[i].ByteSize()
	}
	return size
}

func (this *ListApply_Out_Players) ByteSize() int {
	size := 28
	size += len(this.Nick)
	return size
}

func (this *OperaErrorNotify_Out) Decode(buffer *net.Buffer) {
	this.Resutl = CliqueOperaError(buffer.ReadUint8())
}

func (this *OperaErrorNotify_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(13)
	buffer.WriteUint8(uint8(this.Resutl))
}

func (this *OperaErrorNotify_Out) ByteSize() int {
	size := 3
	return size
}

func (this *EnterClubhouse_In) Decode(buffer *net.Buffer) {
}

func (this *EnterClubhouse_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(14)
}

func (this *EnterClubhouse_In) ByteSize() int {
	size := 2
	return size
}

func (this *EnterClubhouse_Out) Decode(buffer *net.Buffer) {
	this.Ok = buffer.ReadUint8() == 1
	this.Player.Decode(buffer)
}

func (this *EnterClubhouse_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(14)
	if this.Ok {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
	this.Player.Encode(buffer)
}

func (this *EnterClubhouse_Out) ByteSize() int {
	size := 3
	size += this.Player.ByteSize()
	return size
}

func (this *LeaveClubhouse_In) Decode(buffer *net.Buffer) {
}

func (this *LeaveClubhouse_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(15)
}

func (this *LeaveClubhouse_In) ByteSize() int {
	size := 2
	return size
}

func (this *LeaveClubhouse_Out) Decode(buffer *net.Buffer) {
	this.PlayerId = int64(buffer.ReadUint64LE())
}

func (this *LeaveClubhouse_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(15)
	buffer.WriteUint64LE(uint64(this.PlayerId))
}

func (this *LeaveClubhouse_Out) ByteSize() int {
	size := 10
	return size
}

func (this *ClubMove_In) Decode(buffer *net.Buffer) {
	this.ToX = int16(buffer.ReadUint16LE())
	this.ToY = int16(buffer.ReadUint16LE())
}

func (this *ClubMove_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(16)
	buffer.WriteUint16LE(uint16(this.ToX))
	buffer.WriteUint16LE(uint16(this.ToY))
}

func (this *ClubMove_In) ByteSize() int {
	size := 6
	return size
}

func (this *ClubMove_Out) Decode(buffer *net.Buffer) {
	this.PlayerId = int64(buffer.ReadUint64LE())
	this.ToX = int16(buffer.ReadUint16LE())
	this.ToY = int16(buffer.ReadUint16LE())
}

func (this *ClubMove_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(16)
	buffer.WriteUint64LE(uint64(this.PlayerId))
	buffer.WriteUint16LE(uint16(this.ToX))
	buffer.WriteUint16LE(uint16(this.ToY))
}

func (this *ClubMove_Out) ByteSize() int {
	size := 14
	return size
}

func (this *NotifyClubhousePlayers_Out) Decode(buffer *net.Buffer) {
	this.Players = make([]NotifyClubhousePlayers_Out_Players, buffer.ReadUint8())
	for i := 0; i < len(this.Players); i++ {
		this.Players[i].Decode(buffer)
	}
}

func (this *NotifyClubhousePlayers_Out_Players) Decode(buffer *net.Buffer) {
	this.Player.Decode(buffer)
}

func (this *NotifyClubhousePlayers_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(17)
	buffer.WriteUint8(uint8(len(this.Players)))
	for i := 0; i < len(this.Players); i++ {
		this.Players[i].Encode(buffer)
	}
}

func (this *NotifyClubhousePlayers_Out_Players) Encode(buffer *net.Buffer) {
	this.Player.Encode(buffer)
}

func (this *NotifyClubhousePlayers_Out) ByteSize() int {
	size := 3
	for i := 0; i < len(this.Players); i++ {
		size += this.Players[i].ByteSize()
	}
	return size
}

func (this *NotifyClubhousePlayers_Out_Players) ByteSize() int {
	size := 0
	size += this.Player.ByteSize()
	return size
}

func (this *NotifyNewPlayer_Out) Decode(buffer *net.Buffer) {
	this.Player.Decode(buffer)
}

func (this *NotifyNewPlayer_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(18)
	this.Player.Encode(buffer)
}

func (this *NotifyNewPlayer_Out) ByteSize() int {
	size := 2
	size += this.Player.ByteSize()
	return size
}

func (this *NotifyUpdatePlayer_Out) Decode(buffer *net.Buffer) {
	this.PlayerId = int64(buffer.ReadUint64LE())
	this.FashionId = int16(buffer.ReadUint16LE())
	this.InMeditationState = buffer.ReadUint8() == 1
}

func (this *NotifyUpdatePlayer_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(19)
	buffer.WriteUint64LE(uint64(this.PlayerId))
	buffer.WriteUint16LE(uint16(this.FashionId))
	if this.InMeditationState {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *NotifyUpdatePlayer_Out) ByteSize() int {
	size := 13
	return size
}

func (this *CliquePublicInfoSummary_In) Decode(buffer *net.Buffer) {
	this.CliqueId = int64(buffer.ReadUint64LE())
}

func (this *CliquePublicInfoSummary_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(20)
	buffer.WriteUint64LE(uint64(this.CliqueId))
}

func (this *CliquePublicInfoSummary_In) ByteSize() int {
	size := 10
	return size
}

func (this *CliquePublicInfoSummary_Out) Decode(buffer *net.Buffer) {
	this.CliqueId = int64(buffer.ReadUint64LE())
	this.Name = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Level = int16(buffer.ReadUint16LE())
	this.MemberNum = int16(buffer.ReadUint16LE())
	this.OwnerNick = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.OwnerPid = int64(buffer.ReadUint64LE())
	this.Announce = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.Cliques = make([]CliquePublicInfoSummary_Out_Cliques, buffer.ReadUint16LE())
	for i := 0; i < len(this.Cliques); i++ {
		this.Cliques[i].Decode(buffer)
	}
}

func (this *CliquePublicInfoSummary_Out_Cliques) Decode(buffer *net.Buffer) {
	this.CliqueId = int64(buffer.ReadUint64LE())
}

func (this *CliquePublicInfoSummary_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(20)
	buffer.WriteUint64LE(uint64(this.CliqueId))
	buffer.WriteUint16LE(uint16(len(this.Name)))
	buffer.WriteBytes(this.Name)
	buffer.WriteUint16LE(uint16(this.Level))
	buffer.WriteUint16LE(uint16(this.MemberNum))
	buffer.WriteUint16LE(uint16(len(this.OwnerNick)))
	buffer.WriteBytes(this.OwnerNick)
	buffer.WriteUint64LE(uint64(this.OwnerPid))
	buffer.WriteUint16LE(uint16(len(this.Announce)))
	buffer.WriteBytes(this.Announce)
	buffer.WriteUint16LE(uint16(len(this.Cliques)))
	for i := 0; i < len(this.Cliques); i++ {
		this.Cliques[i].Encode(buffer)
	}
}

func (this *CliquePublicInfoSummary_Out_Cliques) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.CliqueId))
}

func (this *CliquePublicInfoSummary_Out) ByteSize() int {
	size := 30
	size += len(this.Name)
	size += len(this.OwnerNick)
	size += len(this.Announce)
	size += len(this.Cliques) * 8
	return size
}

func (this *CliqueAutoAudit_In) Decode(buffer *net.Buffer) {
	this.Level = int16(buffer.ReadUint16LE())
	this.Enable = buffer.ReadUint8() == 1
}

func (this *CliqueAutoAudit_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(21)
	buffer.WriteUint16LE(uint16(this.Level))
	if this.Enable {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *CliqueAutoAudit_In) ByteSize() int {
	size := 5
	return size
}

func (this *CliqueAutoAudit_Out) Decode(buffer *net.Buffer) {
}

func (this *CliqueAutoAudit_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(21)
}

func (this *CliqueAutoAudit_Out) ByteSize() int {
	size := 2
	return size
}

func (this *CliqueBaseDonate_In) Decode(buffer *net.Buffer) {
}

func (this *CliqueBaseDonate_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(22)
}

func (this *CliqueBaseDonate_In) ByteSize() int {
	size := 2
	return size
}

func (this *CliqueBaseDonate_Out) Decode(buffer *net.Buffer) {
}

func (this *CliqueBaseDonate_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(22)
}

func (this *CliqueBaseDonate_Out) ByteSize() int {
	size := 2
	return size
}

func (this *NotifyLeaveClique_Out) Decode(buffer *net.Buffer) {
	this.Pid = int64(buffer.ReadUint64LE())
	this.Reason = NotifyLeaveCliqueReason(buffer.ReadUint8())
}

func (this *NotifyLeaveClique_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(23)
	buffer.WriteUint64LE(uint64(this.Pid))
	buffer.WriteUint8(uint8(this.Reason))
}

func (this *NotifyLeaveClique_Out) ByteSize() int {
	size := 11
	return size
}

func (this *NotifyJoincliqueSuccess_Out) Decode(buffer *net.Buffer) {
	this.Pidlist = make([]NotifyJoincliqueSuccess_Out_Pidlist, buffer.ReadUint8())
	for i := 0; i < len(this.Pidlist); i++ {
		this.Pidlist[i].Decode(buffer)
	}
}

func (this *NotifyJoincliqueSuccess_Out_Pidlist) Decode(buffer *net.Buffer) {
	this.Pid = int64(buffer.ReadUint64LE())
	this.RoleId = int8(buffer.ReadUint8())
	this.Level = int16(buffer.ReadUint16LE())
	this.Nick = buffer.ReadBytes(int(buffer.ReadUint16LE()))
}

func (this *NotifyJoincliqueSuccess_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(24)
	buffer.WriteUint8(uint8(len(this.Pidlist)))
	for i := 0; i < len(this.Pidlist); i++ {
		this.Pidlist[i].Encode(buffer)
	}
}

func (this *NotifyJoincliqueSuccess_Out_Pidlist) Encode(buffer *net.Buffer) {
	buffer.WriteUint64LE(uint64(this.Pid))
	buffer.WriteUint8(uint8(this.RoleId))
	buffer.WriteUint16LE(uint16(this.Level))
	buffer.WriteUint16LE(uint16(len(this.Nick)))
	buffer.WriteBytes(this.Nick)
}

func (this *NotifyJoincliqueSuccess_Out) ByteSize() int {
	size := 3
	for i := 0; i < len(this.Pidlist); i++ {
		size += this.Pidlist[i].ByteSize()
	}
	return size
}

func (this *NotifyJoincliqueSuccess_Out_Pidlist) ByteSize() int {
	size := 13
	size += len(this.Nick)
	return size
}

func (this *NotifyCliqueMangerAction_Out) Decode(buffer *net.Buffer) {
	this.Actiontype = MangeMemberAction(buffer.ReadUint8())
	this.Pid = int64(buffer.ReadUint64LE())
}

func (this *NotifyCliqueMangerAction_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(25)
	buffer.WriteUint8(uint8(this.Actiontype))
	buffer.WriteUint64LE(uint64(this.Pid))
}

func (this *NotifyCliqueMangerAction_Out) ByteSize() int {
	size := 11
	return size
}

func (this *CliqueRecruitment_In) Decode(buffer *net.Buffer) {
}

func (this *CliqueRecruitment_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(26)
}

func (this *CliqueRecruitment_In) ByteSize() int {
	size := 2
	return size
}

func (this *CliqueRecruitment_Out) Decode(buffer *net.Buffer) {
	this.Result = CliqueRecuitmentResult(buffer.ReadUint8())
	this.Timestamp = int64(buffer.ReadUint64LE())
}

func (this *CliqueRecruitment_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(26)
	buffer.WriteUint8(uint8(this.Result))
	buffer.WriteUint64LE(uint64(this.Timestamp))
}

func (this *CliqueRecruitment_Out) ByteSize() int {
	size := 11
	return size
}

func (this *NotifyCliqueAnnounce_Out) Decode(buffer *net.Buffer) {
	this.Announce = buffer.ReadBytes(int(buffer.ReadUint16LE()))
}

func (this *NotifyCliqueAnnounce_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(27)
	buffer.WriteUint16LE(uint16(len(this.Announce)))
	buffer.WriteBytes(this.Announce)
}

func (this *NotifyCliqueAnnounce_Out) ByteSize() int {
	size := 4
	size += len(this.Announce)
	return size
}

func (this *NotifyCliqueElectowner_Out) Decode(buffer *net.Buffer) {
	this.Ownerid = int64(buffer.ReadUint64LE())
}

func (this *NotifyCliqueElectowner_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(28)
	buffer.WriteUint64LE(uint64(this.Ownerid))
}

func (this *NotifyCliqueElectowner_Out) ByteSize() int {
	size := 10
	return size
}

func (this *QuickApply_In) Decode(buffer *net.Buffer) {
}

func (this *QuickApply_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(29)
}

func (this *QuickApply_In) ByteSize() int {
	size := 2
	return size
}

func (this *QuickApply_Out) Decode(buffer *net.Buffer) {
	this.Success = buffer.ReadUint8() == 1
}

func (this *QuickApply_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(29)
	if this.Success {
		buffer.WriteUint8(1)
	} else {
		buffer.WriteUint8(0)
	}
}

func (this *QuickApply_Out) ByteSize() int {
	size := 3
	return size
}

func (this *NotifyContribChange_Out) Decode(buffer *net.Buffer) {
	this.Value = int64(buffer.ReadUint64LE())
}

func (this *NotifyContribChange_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(30)
	buffer.WriteUint64LE(uint64(this.Value))
}

func (this *NotifyContribChange_Out) ByteSize() int {
	size := 10
	return size
}

func (this *OtherClique_In) Decode(buffer *net.Buffer) {
	this.Page = int16(buffer.ReadUint16LE())
}

func (this *OtherClique_In) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(31)
	buffer.WriteUint16LE(uint16(this.Page))
}

func (this *OtherClique_In) ByteSize() int {
	size := 4
	return size
}

func (this *OtherClique_Out) Decode(buffer *net.Buffer) {
	this.TotalNum = int16(buffer.ReadUint16LE())
	this.CliqueList = make([]OtherClique_Out_CliqueList, buffer.ReadUint8())
	for i := 0; i < len(this.CliqueList); i++ {
		this.CliqueList[i].Decode(buffer)
	}
}

func (this *OtherClique_Out_CliqueList) Decode(buffer *net.Buffer) {
	this.Rank = int16(buffer.ReadUint16LE())
	this.Name = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.CliqueId = int64(buffer.ReadUint64LE())
	this.CliqueLevel = int16(buffer.ReadUint16LE())
	this.OwnerName = buffer.ReadBytes(int(buffer.ReadUint16LE()))
	this.OwnerPid = int64(buffer.ReadUint64LE())
	this.CliqueMemberNum = int16(buffer.ReadUint16LE())
}

func (this *OtherClique_Out) Encode(buffer *net.Buffer) {
	buffer.WriteUint8(33)
	buffer.WriteUint8(31)
	buffer.WriteUint16LE(uint16(this.TotalNum))
	buffer.WriteUint8(uint8(len(this.CliqueList)))
	for i := 0; i < len(this.CliqueList); i++ {
		this.CliqueList[i].Encode(buffer)
	}
}

func (this *OtherClique_Out_CliqueList) Encode(buffer *net.Buffer) {
	buffer.WriteUint16LE(uint16(this.Rank))
	buffer.WriteUint16LE(uint16(len(this.Name)))
	buffer.WriteBytes(this.Name)
	buffer.WriteUint64LE(uint64(this.CliqueId))
	buffer.WriteUint16LE(uint16(this.CliqueLevel))
	buffer.WriteUint16LE(uint16(len(this.OwnerName)))
	buffer.WriteBytes(this.OwnerName)
	buffer.WriteUint64LE(uint64(this.OwnerPid))
	buffer.WriteUint16LE(uint16(this.CliqueMemberNum))
}

func (this *OtherClique_Out) ByteSize() int {
	size := 5
	for i := 0; i < len(this.CliqueList); i++ {
		size += this.CliqueList[i].ByteSize()
	}
	return size
}

func (this *OtherClique_Out_CliqueList) ByteSize() int {
	size := 26
	size += len(this.Name)
	size += len(this.OwnerName)
	return size
}
