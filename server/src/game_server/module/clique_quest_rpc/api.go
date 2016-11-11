package clique_quest_rpc

import (
	"core/net"
	"game_server/api/protocol/clique_quest_api"
)

func init() {
	clique_quest_api.SetInHandler(CliqueQuestRPCAPI{})
}

type CliqueQuestRPCAPI struct{}

func (this CliqueQuestRPCAPI) GetCliqueDailyQuest(session *net.Session, in *clique_quest_api.GetCliqueDailyQuest_In) {
	getCliqueDailyQuest(session)
}
func (this CliqueQuestRPCAPI) AwardCliqueDailyQuest(session *net.Session, in *clique_quest_api.AwardCliqueDailyQuest_In) {
	awardCliqueDailyQuest(session, in.Id)
}

func (this CliqueQuestRPCAPI) GetCliqueBuildingQuest(session *net.Session, in *clique_quest_api.GetCliqueBuildingQuest_In) {
	getCliqueBuildingQuest(session)
}

func (this CliqueQuestRPCAPI) AwardCliqueBuildingQuest(session *net.Session, in *clique_quest_api.AwardCliqueBuildingQuest_In) {
	awardCliqueBuildingQuest(session, in.Id)
}
