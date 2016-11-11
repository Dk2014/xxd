package quest

import (
	"core/fail"
	"game_server/mdb"
	"game_server/tlog"
)

func guide(db *mdb.Database, typeid, action int32) {
	fail.When(action != 0 && action != 1, "action not allowed (only allow 0,1)")
	tlog.PlayerGuideFlowLog(db, typeid, action)
}
