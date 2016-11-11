package clique_escort_rpc

import (
	"core/net"
	"core/time"
	"game_server/api/protocol/clique_escort_api"
	"game_server/dat/channel_dat"
	"game_server/dat/clique_dat"
	"game_server/mdb"
	"game_server/module"
	gotime "time"
)

func init() {
	module.CliqueEscortRPC = CliqueEscortMod{}
}

type CliqueEscortMod struct{}

func (mod CliqueEscortMod) RecoverBattleWin(db *mdb.Database, fighterPid, boatId int64) {
	recoverBattleWin(db, fighterPid, boatId)
}

func (mod CliqueEscortMod) RecoverBattleLose(db *mdb.Database, fighterPid, boatId int64) {
	recoverBattleLose(db, fighterPid, boatId)
}

func (mod CliqueEscortMod) HijackBattleWin(db *mdb.Database, fighterPid, boatId int64) {
	hijackBattleWin(db, fighterPid, boatId)
}

//登陆时候检查 劫持完成 或者 护送完成 通知
func (mod CliqueEscortMod) LoginEscortNotify(session *net.Session) {
	state := module.State(session)

	playerEscortInfo := lookupEscortInfo(state.Database)
	playerHijackInfo := lookupHijackInfo(state.Database)

	if playerEscortInfo.Status == int8(clique_escort_api.ESCORT_STATUS_FINISHED) {
		session.Send(&clique_escort_api.NotifyEscortFinished_Out{})
	}
	if playerHijackInfo.Status == int8(clique_escort_api.HIJACK_STATUS_FINISHED) {
		session.Send(&clique_escort_api.NotifyHijackFinished_Out{})
	}
}

func (mod CliqueEscortMod) NewBoatMessage(db *mdb.Database, msg channel_dat.MessageTpl) {
	newMessage(db, msg)
}

func (mod CliqueEscortMod) MigrateBoatToNewClique(db *mdb.Database, newCliqueId int64) {
	playerEscortInfo := db.Lookup.PlayerGlobalCliqueEscort(db.PlayerId())
	if playerEscortInfo == nil {
		return
	}
	if playerEscortInfo.Status == int8(clique_escort_api.ESCORT_STATUS_NONE) || playerEscortInfo.Status == int8(clique_escort_api.ESCORT_STATUS_FINISHED) {
		//没有运送 或者 运送结束带领取奖励
		return
	}
	var boat *mdb.GlobalCliqueBoat
	db.Select.GlobalCliqueBoat(func(row *mdb.GlobalCliqueBoatRow) {
		if row.OwnerPid() == db.PlayerId() {
			boat = row.GoObject()
			row.Break()
		}
	})
	if boat != nil {
		boat.CliqueId = newCliqueId
		db.Update.GlobalCliqueBoat(boat)
		//A 劫持了 B 的船，B进入A的帮派，这个时候如果船还在劫持中
		if boat.HijackerPid > 0 && module.CliqueRPC.IsCliqueMember(newCliqueId, boat.HijackerPid) && isHijacking(time.GetNowTime(), boat) {
			db.AgentExecute(boat.HijackerPid, func(agentDB *mdb.Database) {
				playerHijackInfo := lookupHijackInfo(agentDB)
				playerHijackInfo.Status = int8(clique_escort_api.HIJACK_STATUS_NONE)
				playerHijackInfo.HijackedBoatType = 0
				agentDB.Update.PlayerGlobalCliqueHijack(playerHijackInfo)
			})
			boat.Status = int8(clique_escort_api.BOAT_STATUS_ESCORT)
			boat.HijackerPid = 0
			db.Update.GlobalCliqueBoat(boat)
			escortTimerCallback(db, boat)
		}
	}
}

//运镖测试负责接口
func (mod CliqueEscortMod) ForceEscort(db *mdb.Database) {
	playerCliqueInfo := db.Lookup.PlayerGlobalCliqueInfo(db.PlayerId())
	if playerCliqueInfo == nil || playerCliqueInfo.CliqueId <= 0 {
		return
	}
	playerEscortInfo := lookupEscortInfo(db)
	if playerEscortInfo.EscortBoatType <= 0 {
		playerEscortInfo.EscortBoatType = clique_dat.GetRandomBoatId()
		db.Update.PlayerGlobalCliqueEscort(playerEscortInfo)
	}
	if playerEscortInfo.Status == int8(clique_escort_api.ESCORT_STATUS_FINISHED) ||
		playerEscortInfo.Status == int8(clique_escort_api.ESCORT_STATUS_NONE) {
		boat := &mdb.GlobalCliqueBoat{
			CliqueId:             playerCliqueInfo.CliqueId,
			OwnerPid:             db.PlayerId(),
			BoatType:             playerEscortInfo.EscortBoatType,
			EscortStartTimestamp: time.GetNowTime(),
			Status:               int8(clique_escort_api.BOAT_STATUS_ESCORT),
		}
		db.Insert.GlobalCliqueBoat(boat)
		playerEscortInfo.Status = int8(clique_escort_api.ESCORT_STATUS_ESCORT)
		playerEscortInfo.DailyEscortNum++
		playerEscortInfo.DailyEscortTimestamp = time.GetNowTime()
		db.Update.PlayerGlobalCliqueEscort(playerEscortInfo)
		StartTimer(boat.Id, gotime.Duration(clique_dat.ESCORT_TIME)*gotime.Second)
	}
}
