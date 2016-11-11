package quest

import (
	"core/net"
	"game_server/api/protocol/quest_api"
	"game_server/dat/quest_dat"
	"game_server/module"
	"game_server/xdlog"
)

func init() {
	quest_api.SetInHandler(QuestAPI{})
}

type QuestAPI struct {
}

func (api QuestAPI) UpdateQuest(session *net.Session, in *quest_api.UpdateQuest_In) {
	module.Quest.RefreshQuest(module.State(session), quest_dat.QUEST_TYPE_NONE, xdlog.ET_QUEST)
}

func (api QuestAPI) GetDailyInfo(session *net.Session, in *quest_api.GetDailyInfo_In) {
	getDailyInfo(session)
}

func (api QuestAPI) AwardDaily(session *net.Session, in *quest_api.AwardDaily_In) {
	awardDailyQuest(session, in.Id, xdlog.ET_DAILY_QUEST)
}

func (api QuestAPI) Guide(session *net.Session, in *quest_api.Guide_In) {
	state := module.State(session)
	if state.Database.PlayerId() > 0 {
		guide(state.Database, int32(in.GuideType), int32(in.Action))
	}
}

func (api QuestAPI) GetExtendQuestInfoByNpcId(session *net.Session, in *quest_api.GetExtendQuestInfoByNpcId_In) {
	state := module.State(session)
	out := &quest_api.GetExtendQuestInfoByNpcId_Out{}
	getExtendQuestInfo(state, in.NpcId, out)
	session.Send(out)
}

func (api QuestAPI) TakeExtendQuestAward(session *net.Session, in *quest_api.TakeExtendQuestAward_In) {
	state := module.State(session)
	takeExtendQuestAward(state, in.QuestId, xdlog.ET_EXTEND_QUEST)
	out := &quest_api.TakeExtendQuestAward_Out{
		QuestId: in.QuestId,
	}
	session.Send(out)
}

func (api QuestAPI) GetPannelQuestInfo(session *net.Session, in *quest_api.GetPannelQuestInfo_In) {
	getPannelQuestInfo(session)
}

func (api QuestAPI) GiveUpAdditionQuest(session *net.Session, in *quest_api.GiveUpAdditionQuest_In) {
	giveUpAdditionQuest(session, in.QuestId)
}

func (api QuestAPI) TakeAdditionQuest(session *net.Session, in *quest_api.TakeAdditionQuest_In) {
	takeAdditionQuest(session, in.QuestId)
}

func (api QuestAPI) TakeAdditionQuestAward(session *net.Session, in *quest_api.TakeAdditionQuestAward_In) {
	takeAdditionQuestAward(session, in.QuestId, xdlog.ET_ADDITION_QUEST)
}

func (api QuestAPI) GetAdditionQuest(session *net.Session, in *quest_api.GetAdditionQuest_In) {
	getAdditionQuest(session)
}

func (api QuestAPI) RefreshAdditionQuest(session *net.Session, in *quest_api.RefreshAdditionQuest_In) {
	refreshAdditionQuest(session, in.QuestId)
}

func (api QuestAPI) TakeQuestStarsAwaded(session *net.Session, in *quest_api.TakeQuestStarsAwaded_In) {
	takeQuestStarsAwaded(session, in.StarsLevel)
}
