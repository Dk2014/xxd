package skill

import (
	"core/net"
	"game_server/api/protocol/skill_api"
	"game_server/xdlog"
)

func init() {
	skill_api.SetInHandler(SkillAPI{})
}

type SkillAPI struct {
}

func (this SkillAPI) GetAllSkillsInfo(session *net.Session, in *skill_api.GetAllSkillsInfo_In) {
	out := &skill_api.GetAllSkillsInfo_Out{}
	GetAllSkillsInfo(session, out)
	session.Send(out)
}

func (this SkillAPI) EquipSkill(session *net.Session, in *skill_api.EquipSkill_In) {
	changeSkill(session, in.RoleId, in.OrderNumber, in.SkillId)
	session.Send(&skill_api.EquipSkill_Out{})
}

func (this SkillAPI) UnequipSkill(session *net.Session, in *skill_api.UnequipSkill_In) {
	changeSkill(session, in.RoleId, in.OrderNumber, 0)
	session.Send(&skill_api.UnequipSkill_Out{})
}

func (this SkillAPI) StudySkillByCheat(session *net.Session, in *skill_api.StudySkillByCheat_In) {
	result := studyskill(session, in.RoleId, in.ItemId, xdlog.ET_ROLE_SKILL)
	session.Send(&skill_api.StudySkillByCheat_Out{
		Result: result,
	})
}

func (this SkillAPI) TrainSkill(session *net.Session, in *skill_api.TrainSkill_In) {
	trainSkill(session, in.RoleId, in.SkillId)
	session.Send(&skill_api.TrainSkill_Out{})
}

func (api SkillAPI) FlushSkill(session *net.Session, in *skill_api.FlushSkill_In) {
	flush_time := flushSkill(session, in.RoleId)
	session.Send(&skill_api.FlushSkill_Out{
		FlushTime: flush_time,
	})
}
