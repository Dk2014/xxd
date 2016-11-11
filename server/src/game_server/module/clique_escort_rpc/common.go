package clique_escort_rpc

import (
	"core/time"
	"fmt"
	"game_server/api/protocol/clique_escort_api"
	"game_server/dat/clique_dat"
	"game_server/global"
	"game_server/mdb"
	"game_server/module"
)

func lookupEscortInfo(db *mdb.Database) (info *mdb.PlayerGlobalCliqueEscort) {
	info = db.Lookup.PlayerGlobalCliqueEscort(db.PlayerId())
	if info == nil {
		info = &mdb.PlayerGlobalCliqueEscort{
			Pid:                  db.PlayerId(),
			DailyEscortNum:       0,
			DailyEscortTimestamp: 0,
			EscortBoatType:       0,
		}
		db.Insert.PlayerGlobalCliqueEscort(info)
	}
	if !time.IsToday(info.DailyEscortTimestamp) {
		info.DailyEscortNum = 0
	}
	return info
}

func lookupHijackInfo(db *mdb.Database) (info *mdb.PlayerGlobalCliqueHijack) {
	info = db.Lookup.PlayerGlobalCliqueHijack(db.PlayerId())
	if info == nil {
		info = &mdb.PlayerGlobalCliqueHijack{
			Pid:                  db.PlayerId(),
			DailyHijackNum:       0,
			DailyHijackTimestamp: 0,
		}
		db.Insert.PlayerGlobalCliqueHijack(info)
	}
	if !time.IsToday(info.DailyHijackTimestamp) {
		info.DailyHijackNum = 0
	}
	return info
}

//把船的状态更正
func updateBoatHook(boat *mdb.GlobalCliqueBoat) *mdb.GlobalCliqueBoat {
	//更新船的状态
	//劫持中或者夺回中需要检查是否比变成劫持完成
	if boat.Status == int8(clique_escort_api.BOAT_STATUS_HIJACK) ||
		boat.Status == int8(clique_escort_api.BOAT_STATUS_RECOVER) {
		if time.GetNowTime()-boat.HijackTimestamp >= clique_dat.HIJACK_TIME {
			boat.Status = int8(clique_escort_api.BOAT_STATUS_HIJACK_FINISH)
			boat.EscortStartTimestamp = boat.HijackTimestamp + clique_dat.HIJACK_TIME
		}
	}
	//劫持完成和护送中由  escort_start_timestamp 和 total_escort_time 来决定是否处于完成中
	return boat
}

func lookupBoat(db *mdb.Database, boatId int64) *mdb.GlobalCliqueBoat {
	boat := db.Lookup.GlobalCliqueBoat(boatId)
	if boat != nil {
		boat = updateBoatHook(boat)
	}
	return boat
}

func mdbMsgToRespMsg(msg *mdb.PlayerGlobalCliqueEscortMessage) clique_escort_api.CliqueBoatMessage {
	return clique_escort_api.CliqueBoatMessage{
		Id:         msg.Id,
		TplId:      msg.TplId,
		Parameters: []byte(msg.Parameters),
	}
}

func mdbBoatToRespBoat(boat *mdb.GlobalCliqueBoat) clique_escort_api.CliqueBoat {
	now := time.GetNowTime()
	respBoat := clique_escort_api.CliqueBoat{
		BoatId:   boat.Id,
		BoatType: boat.BoatType,
		Status:   clique_escort_api.BoatStatus(boat.Status),
	}
	ownerInfo := global.GetPlayerInfo(boat.OwnerPid)
	respBoat.OwnerPid = boat.OwnerPid
	respBoat.OwnerNick = ownerInfo.PlayerNick
	respBoat.OwnerLevel = ownerInfo.RoleLevel
	respBoat.EscortTime = int64(boat.TotalEscortTime)
	if boat.EscortStartTimestamp > 0 {
		respBoat.StartTimestamp = boat.EscortStartTimestamp
	}
	if module.CliqueRPC.IsCliqueMember(boat.CliqueId, boat.OwnerPid) {
		respBoat.CliqueId = boat.CliqueId
		respBoat.CliqueName = []byte(module.CliqueRPC.GetCliqueNameById(boat.CliqueId))
	}
	if boat.HijackerPid > 0 {
		hijackerInfo := global.GetPlayerInfo(boat.HijackerPid)
		respBoat.HijackerNick = hijackerInfo.PlayerNick
		respBoat.HijackStartTimestamp = boat.HijackTimestamp
		respBoat.HijackerPid = boat.HijackerPid
	}
	if boat.RecoverPid > 0 && isHijacking(now, boat) {
		respBoat.Status = clique_escort_api.BOAT_STATUS_RECOVER
	}
	return respBoat
}

func isHijacking(timestamp int64, boat *mdb.GlobalCliqueBoat) bool {
	return boat.Status == int8(clique_escort_api.BOAT_STATUS_HIJACK) &&
		timestamp-boat.HijackTimestamp < clique_dat.HIJACK_TIME
}

func isEscortFinished(timestamp int64, boat *mdb.GlobalCliqueBoat) bool {
	status := boat.Status
	escortStart := boat.EscortStartTimestamp
	if status == int8(clique_escort_api.BOAT_STATUS_HIJACK) || status == int8(clique_escort_api.BOAT_STATUS_RECOVER) {
		if timestamp-boat.HijackTimestamp >= clique_dat.HIJACK_TIME {
			status = int8(clique_escort_api.BOAT_STATUS_HIJACK_FINISH)
			escortStart = boat.HijackTimestamp + clique_dat.HIJACK_TIME
		}
	}
	if (status == int8(clique_escort_api.BOAT_STATUS_ESCORT) || status == int8(clique_escort_api.BOAT_STATUS_HIJACK_FINISH)) &&
		int64(boat.TotalEscortTime)+timestamp-escortStart >= clique_dat.ESCORT_TIME {
		return true
	}
	return false
}

func nextBoatAction(boat *mdb.GlobalCliqueBoat) int64 {
	now := time.GetNowTime()
	if boat.Status == int8(clique_escort_api.BOAT_STATUS_HIJACK) ||
		boat.Status == int8(clique_escort_api.BOAT_STATUS_RECOVER) { //劫持中？
		if now-boat.HijackTimestamp >= clique_dat.HIJACK_TIME { //劫持结束？
			return 0
		}
		return boat.HijackTimestamp + clique_dat.HIJACK_TIME - now
	}
	if !isEscortFinished(now, boat) { //运送中？
		fmt.Println("next action", boat.EscortStartTimestamp+(clique_dat.ESCORT_TIME-int64(boat.EscortStartTimestamp))-now)
		return boat.EscortStartTimestamp + (clique_dat.ESCORT_TIME - int64(boat.EscortStartTimestamp)) - now
	}
	return 0 //运送结束
}
