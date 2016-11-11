package team

import (
	"core/net"
	"game_server/api/protocol/team_api"
	"game_server/module"
)

func init() {
	team_api.SetInHandler(TeamAPI{})
}

type TeamAPI struct{}

// 获取布阵信息
func (team TeamAPI) GetFormationInfo(session *net.Session, in *team_api.GetFormationInfo_In) {
	state := module.State(session)
	formation := state.Database.Lookup.PlayerFormation(state.PlayerId)
	teamInfo := state.Database.Lookup.PlayerTeamInfo(state.PlayerId)
	session.Send(&team_api.GetFormationInfo_Out{
		Pos0Role: formation.Pos0,
		Pos1Role: formation.Pos1,
		Pos2Role: formation.Pos2,
		Pos3Role: formation.Pos3,
		Pos4Role: formation.Pos4,
		Pos5Role: formation.Pos5,
		Pos6Role: formation.Pos6,
		Pos7Role: formation.Pos7,
		Pos8Role: formation.Pos8,

		Relationship: teamInfo.Relationship,
		HealthLv:     teamInfo.HealthLv,
		AttackLv:     teamInfo.AttackLv,
		DefenceLv:    teamInfo.DefenceLv,
	})
}

// 角色上阵
func (team TeamAPI) UpFormation(session *net.Session, in *team_api.UpFormation_In) {
	state := module.State(session)
	upFormation(state, in.RoleId, in.Pos)
	session.Send(&team_api.UpFormation_Out{})
}

// 角色下阵
func (team TeamAPI) DownFormation(session *net.Session, in *team_api.DownFormation_In) {
	state := module.State(session)
	downFormation(state.Database, in.Pos)
	session.Send(&team_api.UpFormation_Out{})
}

// 交换两个角色布阵位置
func (team TeamAPI) SwapFormation(session *net.Session, in *team_api.SwapFormation_In) {
	state := module.State(session)
	swapFormation(state, in.PosFrom, in.PosTo)
	session.Send(&team_api.SwapFormation_Out{})
}

// 在阵上和不在阵上的两个角色交换
func (team TeamAPI) ReplaceFormation(session *net.Session, in *team_api.ReplaceFormation_In) {
	state := module.State(session)
	replaceFormation(state, in.RoleId, in.Pos)
	session.Send(&team_api.ReplaceFormation_Out{})
}

// 伙伴配合训练
func (team TeamAPI) TrainingTeamship(session *net.Session, in *team_api.TrainingTeamship_In) {
	state := module.State(session)
	trainingTeamship(state, in.AttrInd)
	session.Send(&team_api.TrainingTeamship_Out{})
}
