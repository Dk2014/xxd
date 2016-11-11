package clique_quest_rpc

import (
	"game_server/mdb"
	"game_server/module"
)

func init() {
	module.CliqueQuestRPC = CliqueQuestMod{}
}

type CliqueQuestMod struct{}

func (mod CliqueQuestMod) UpDatePlayerCliqueDailyQuest(db *mdb.Database, class int16) {
	upDatePlayerCliqueDailyQuest(db, class)
}

func (mod CliqueQuestMod) AddCliqueBuildingQuest(cliqueInfo *mdb.GlobalClique, db *mdb.Database, pid int64) {
	addCliqueBuildingQuest(cliqueInfo, db, pid)
}

func (mod CliqueQuestMod) CleanCliqueBuildingQuest(db *mdb.Database, pid int64) {
	cleanCliqueBuildingQuest(db, pid)
}
