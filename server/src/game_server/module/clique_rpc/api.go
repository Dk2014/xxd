package clique_rpc

import (
	"core/net"
	"game_server/api/protocol/clique_api"
	//"game_server/module"
)

func init() {
	clique_api.SetInHandler(CliqueAPI{})
}

type CliqueAPI struct{}

//创建帮派
func (this CliqueAPI) CreateClique(session *net.Session, in *clique_api.CreateClique_In) {
	createClique(session, string(in.Name), string(in.Announce))
}

//申请加入
func (this CliqueAPI) ApplyJoinClique(session *net.Session, in *clique_api.ApplyJoinClique_In) {
	applyJoinClique(session, in.CliqueId)
}

//取消申请
func (this CliqueAPI) CancelApplyClique(session *net.Session, in *clique_api.CancelApplyClique_In) {
	cancelApplyClique(session, in.CliqueId)
}

//处理帮派加入申请
func (this CliqueAPI) ProcessJoinApply(session *net.Session, in *clique_api.ProcessJoinApply_In) {
	processJoinApply(session, in)
}

//解散帮派
func (this CliqueAPI) DestoryClique(session *net.Session, in *clique_api.DestoryClique_In) {
	destoryClique(session, in)
}

//弹劾帮主
func (this CliqueAPI) ElectOwner(session *net.Session, in *clique_api.ElectOwner_In) {
	electOwner(session)
}

//离开帮派
func (this CliqueAPI) LeaveClique(session *net.Session, in *clique_api.LeaveClique_In) {
	leaveClique(session)
}

//成员管理
func (this CliqueAPI) MangeMember(sessoin *net.Session, in *clique_api.MangeMember_In) {
	mangeMember(sessoin, in)
}

//更新公告
func (this CliqueAPI) UpdateAnnounce(session *net.Session, in *clique_api.UpdateAnnounce_In) {
	updateAnnounce(session, in)
}

//帮派列表
func (this CliqueAPI) ListClique(session *net.Session, in *clique_api.ListClique_In) {
	listClique(session, in)
}

//帮派信息（对公）
func (this CliqueAPI) CliquePublicInfo(session *net.Session, in *clique_api.CliquePublicInfo_In) {
	cliquePublicInfo(session, in)
}

//帮派基础信息（对公）搜索用
func (this CliqueAPI) CliquePublicInfoSummary(session *net.Session, in *clique_api.CliquePublicInfoSummary_In) {
	cliquePublicInfoSummary(session, in.CliqueId)
}

//帮派信息
func (this CliqueAPI) CliqueInfo(session *net.Session, in *clique_api.CliqueInfo_In) {
	cliqueInfo(session)
}

//帮派申请列表
func (this CliqueAPI) ListApply(session *net.Session, in *clique_api.ListApply_In) {
	listApply(session, in)
}

//进入帮派集会场所
func (this CliqueAPI) EnterClubhouse(session *net.Session, in *clique_api.EnterClubhouse_In) {
	enterClubhouse(session)
}

//离开集会场所
func (this CliqueAPI) LeaveClubhouse(session *net.Session, in *clique_api.LeaveClubhouse_In) {
	leaveClubhouse(session)
}

//集会场所内移动
func (this CliqueAPI) ClubMove(session *net.Session, in *clique_api.ClubMove_In) {
	clubhouseMove(session, in.ToX, in.ToY)
}

//自动审核
func (this CliqueAPI) CliqueAutoAudit(session *net.Session, in *clique_api.CliqueAutoAudit_In) {
	cliqueAutoAudit(session, in.Level, in.Enable)
}

func (this CliqueAPI) CliqueBaseDonate(session *net.Session, in *clique_api.CliqueBaseDonate_In) {
	// 作废
}

//帮派招募公告
func (this CliqueAPI) CliqueRecruitment(session *net.Session, in *clique_api.CliqueRecruitment_In) {
	cliqueRecruitment(session)
}

//快速加入
func (this CliqueAPI) QuickApply(session *net.Session, in *clique_api.QuickApply_In) {
	quickApply(session)
}

func (this CliqueAPI) OtherClique(session *net.Session, in *clique_api.OtherClique_In) {
	otherClique(session, int(in.Page))
}
