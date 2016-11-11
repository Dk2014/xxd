package multi_level

import (
	"core/fail"
	"core/net"
	"game_server/dat/multi_level_dat"
	"game_server/dat/role_dat"
	"game_server/mdb"
	"game_server/module"
)

func init() {
	module.MultiLevel = ModMultiLevel{}
	module.PrepareStoreEvent.Regisiter(PerpareStoreHandler)
}

func PerpareStoreHandler(session *net.Session) {
	state := module.State(session)

	cancelAutoMatch(state)

	if state.MultiLevelRoomId > 0 {
		leaveRoom(session)
	}
}

type ModMultiLevel struct {
}

func (mod ModMultiLevel) OpenFunc(db *mdb.Database) {
	// 找出玩家最高等级伙伴,设置为上阵
	var maxLevel int16
	var buddyRoleId int8
	db.Select.PlayerRole(func(row *mdb.PlayerRoleRow) {
		if !role_dat.IsMainRole(row.RoleId()) && row.Level() > maxLevel {
			maxLevel = row.Level()
			buddyRoleId = row.RoleId()
		}
	})

	db.Insert.PlayerMultiLevelInfo(&mdb.PlayerMultiLevelInfo{
		Pid:          db.PlayerId(),
		BuddyRoleId:  buddyRoleId,
		BuddyRow:     multi_level_dat.MULTI_LEVEL_DEPLOY_ROW_2,
		TacticalGrid: 0,
		DailyNum:     0,
		BattleTime:   0,
		Lock:         multi_level_dat.GetFirstMultiLevelLock(),
	})
}

func (mod ModMultiLevel) StartBattle(state *module.SessionState, levelId int32) {
	fail.When(state.MultiLevelRoomId == 0, "[MultiLevel StartBattle] error room id")

	room, ok := multiDataTable.getRoom(state.MultiLevelRoomId)
	fail.When(!ok, "[MultiLevel StartBattle] not found room")
	fail.When(room.LevelId != levelId, "not match level id")
	fail.When(room.LeaderPid != state.PlayerId, "[MultiLevel StartBattle] player is not leader")
	fail.When(room.OnFighting, "[MultiLevel StartBattle] fighting")

	newBattle(room)
}

//获取指定房间的某一玩家是否合适接受通知
func (mod ModMultiLevel) IsMultiInvitePlayerSuitable(roomId int64, inviterId int64) bool {
	multiDataTable.roomMutex.Lock()
	defer multiDataTable.roomMutex.Unlock()
	room, ok := multiDataTable.roomList[roomId]
	_, is_still_in := room.Partners[inviterId]
	return !room.OnFighting && ok && is_still_in
}

func (mod ModMultiLevel) LeaveMultiLevel(session *net.Session) {
	leaveRoom(session)
}
