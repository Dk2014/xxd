package friend_rpc

import (
	"core/fail"
	"core/net"
	"core/time"
	"game_server/api/protocol/multi_level_api"
	"game_server/dat/friend_dat"
	"game_server/global"
	"game_server/mdb"
	"game_server/module"
)

func init() {
	module.Friend = FriendMod{}
}

type FriendMod struct{}

func (this FriendMod) ListenByName(state *module.SessionState, name string) {
	pid, ok := global.GetPlayerIdWithNick(name)
	fail.When(!ok, "can't found player by nickname with ListenByName")
	ListenByPid(state, pid)
}

func (this FriendMod) SendOfflineMessages(session *net.Session) {
	if out, ok := GetOfflineMessages(module.State(session)); ok {
		session.Send(out)
	}
}

func (this FriendMod) GetMultiLevelOnlineFriends(db *mdb.Database) net.Response {
	rsp := &multi_level_api.GetOnlineFriend_Out{}
	db.Select.PlayerGlobalFriend(func(row *mdb.PlayerGlobalFriendRow) {
		if row.FriendMode() != friend_dat.FRIEND_MODE_FRIEND {
			return
		}

		if _, ok := module.Player.GetPlayerOnline(row.FriendPid()); !ok {
			return
		}

		friendInfo := global.GetPlayerInfo(row.FriendPid())
		rsp.Friends = append(rsp.Friends,
			multi_level_api.GetOnlineFriend_Out_Friends{
				Pid:      friendInfo.PlayerId,
				RoleId:   friendInfo.RoleId,
				Nickname: friendInfo.PlayerNick,
				Level:    friendInfo.RoleLevel,
				FightNum: friendInfo.FightNum,
				DailyNum: friendInfo.MultiLevelDailyNum,
				Lock:     friendInfo.MultiLevelLock,
			})
	})
	return rsp
}

func (this FriendMod) GetSendHeartRecord(db *mdb.Database, friendPid int64) (sendHeartRecord *mdb.PlayerSendHeartRecord) {
	db.Select.PlayerSendHeartRecord(func(row *mdb.PlayerSendHeartRecordRow) {
		if row.FriendPid() == friendPid {
			sendHeartRecord = row.GoObject()
			row.Break()
		}
	})
	return
}

//近供RPC回调使用，不直接调用
func (this FriendMod) SendHeart(db *mdb.Database, friendPid int64) {
	sendHeartRecord := module.Friend.GetSendHeartRecord(db, friendPid)
	nowUnix := time.GetNowTime()
	if sendHeartRecord == nil {
		sendHeartRecord = &mdb.PlayerSendHeartRecord{
			Pid:           db.PlayerId(),
			FriendPid:     friendPid,
			SendHeartTime: nowUnix,
		}
		db.Insert.PlayerSendHeartRecord(sendHeartRecord)
	} else {
		fail.When(sendHeartRecord.SendHeartTime+friend_dat.FRIEND_SEND_HEART_INTERVAL > nowUnix, "赠送爱心CD中")
		sendHeartRecord.SendHeartTime = nowUnix
		db.Update.PlayerSendHeartRecord(sendHeartRecord)
	}
}
