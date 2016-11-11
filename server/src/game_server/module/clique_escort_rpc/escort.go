package clique_escort_rpc

import (
	"core/net"
	"core/time"
	"game_server/api/protocol/clique_escort_api"
	"game_server/dat/clique_dat"
	"game_server/dat/player_dat"
	"game_server/global"
	"game_server/mdb"
	"game_server/module"
	"game_server/module/rpc"
	"game_server/xdlog"
	"math/rand"
	gotime "time"
)

func escortInfo(session *net.Session) {
	state := module.State(session)

	playerEscortInfo := lookupEscortInfo(state.Database)
	playerHijackInfo := lookupHijackInfo(state.Database)
	out := &clique_escort_api.EscortInfo_Out{
		DailyEscortNum: playerEscortInfo.DailyEscortNum,
		EscortStatus:   clique_escort_api.EscortStatus(playerEscortInfo.Status),
		DailyHijackNum: playerHijackInfo.DailyHijackNum,
		HijackStatus:   clique_escort_api.HijackStatus(playerHijackInfo.Status),
		BoatType:       playerEscortInfo.EscortBoatType,
	}

	session.Send(out)
}

func getIngotBoat(session *net.Session) {
	state := module.State(session)
	playerCliqueInfo := state.Database.Lookup.PlayerGlobalCliqueInfo(state.PlayerId)
	if playerCliqueInfo == nil || playerCliqueInfo.CliqueId <= 0 {
		//没有帮派
		session.Send(&clique_escort_api.GetIngotBoat_Out{Ok: false})
		return
	}

	playerEscortInfo := lookupEscortInfo(state.Database)
	if playerEscortInfo.DailyEscortNum >= clique_dat.ESCORT_DAILY_NUM {
		//没有次数
		session.Send(&clique_escort_api.GetIngotBoat_Out{Ok: false})
		return
	}
	if playerEscortInfo.Status != int8(clique_escort_api.ESCORT_STATUS_NONE) {
		//上次护送尚未结束
		session.Send(&clique_escort_api.GetIngotBoat_Out{Ok: false})
		return
	}
	ingotBoat := clique_dat.GetIngotBoat()
	rpc.RemoteDecMoney(state.PlayerId, ingotBoat.SelectCostIngot, player_dat.INGOT, 0 /*TODO tlog */, xdlog.ET_CLIQUE_BOAT, func() {
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(state.PlayerId, func(agentDB *mdb.Database) {
				playerCliqueInfo := agentDB.Lookup.PlayerGlobalCliqueInfo(agentDB.PlayerId())
				if playerCliqueInfo == nil || playerCliqueInfo.CliqueId <= 0 {
					//notify 已经被提出帮派
					return
				}
				playerEscortInfo := lookupEscortInfo(agentDB)
				if playerEscortInfo.Status != int8(clique_escort_api.ESCORT_STATUS_NONE) {
					//notify 上次护送尚未结束
					return
				}
				playerEscortInfo.EscortBoatType = ingotBoat.Id
				agentDB.Update.PlayerGlobalCliqueEscort(playerEscortInfo)
				session.Send(&clique_escort_api.GetIngotBoat_Out{Ok: true})
			})
		})
	})
}

func startEscort(session *net.Session) {
	state := module.State(session)
	out := &clique_escort_api.StartEscort_Out{}

	//检查时间段是否合法
	todayZero := time.GetTodayZero()
	now := time.GetNowTime()
	start := todayZero + clique_dat.ESCORT_START_HOUR
	end := todayZero + clique_dat.ESCORT_END_HOUR

	if now < start || now >= end {
		out.Result = clique_escort_api.START_ESCORT_RESULT_ILLEGAL_TIME
		session.Send(out)
		return
	}

	playerCliqueInfo := state.Database.Lookup.PlayerGlobalCliqueInfo(state.PlayerId)
	if playerCliqueInfo == nil || playerCliqueInfo.CliqueId <= 0 {
		out.Result = clique_escort_api.START_ESCORT_RESULT_NO_CLIQUE
		session.Send(out)
		return
	}

	playerEscortInfo := lookupEscortInfo(state.Database)
	if playerEscortInfo.DailyEscortNum >= clique_dat.ESCORT_DAILY_NUM {
		out.Result = clique_escort_api.START_ESCORT_RESULT_RUN_OUT
		session.Send(out)
		return
	}
	if playerEscortInfo.EscortBoatType <= 0 {
		out.Result = clique_escort_api.START_ESCORT_RESULT_NO_BOAT
		session.Send(out)
		return
	}
	if playerEscortInfo.Status != int8(clique_escort_api.ESCORT_STATUS_NONE) {
		out.Result = clique_escort_api.START_ESCORT_RESULT_ESCORT_NOT_END
		session.Send(out)
		return
	}
	boat := &mdb.GlobalCliqueBoat{
		CliqueId:             playerCliqueInfo.CliqueId,
		OwnerPid:             state.PlayerId,
		BoatType:             playerEscortInfo.EscortBoatType,
		EscortStartTimestamp: time.GetNowTime(),
		Status:               int8(clique_escort_api.BOAT_STATUS_ESCORT),
	}
	state.Database.Insert.GlobalCliqueBoat(boat)
	playerEscortInfo.Status = int8(clique_escort_api.ESCORT_STATUS_ESCORT)
	playerEscortInfo.DailyEscortNum++
	playerEscortInfo.DailyEscortTimestamp = now
	state.Database.Update.PlayerGlobalCliqueEscort(playerEscortInfo)

	out.Result = clique_escort_api.START_ESCORT_RESULT_SUCCESS
	out.Boat = mdbBoatToRespBoat(boat)

	//设置定时器
	StartTimer(boat.Id, gotime.Duration(clique_dat.ESCORT_TIME)*gotime.Second)
	module.CliqueQuestRPC.UpDatePlayerCliqueDailyQuest(state.Database, clique_dat.CLIQUE_DAYLY_QUEST_SEVERAL)
	session.Send(out)
}

func hijackBoat(session *net.Session, boatId int64) {
	state := module.State(session)
	out := &clique_escort_api.HijackBoat_Out{}
	now := time.GetNowTime()

	playerCliqueInfo := state.Database.Lookup.PlayerGlobalCliqueInfo(state.PlayerId)
	if playerCliqueInfo == nil || playerCliqueInfo.CliqueId <= 0 {
		out.Result = clique_escort_api.HIJACK_BOAT_RESULT_NO_CLIQUE
		session.Send(out)
		return
	}

	playerHijackInfo := lookupHijackInfo(state.Database)
	if playerHijackInfo.DailyHijackNum >= clique_dat.HIJACK_DAILY_NUM {
		out.Result = clique_escort_api.HIJACK_BOAT_RESULT_RUN_OUT
		session.Send(out)
		return
	}
	if playerHijackInfo.Status != int8(clique_escort_api.HIJACK_STATUS_NONE) {
		out.Result = clique_escort_api.HIJACK_BOAT_RESULT_HIJACK_NOT_END
		session.Send(out)
		return
	}
	boat := lookupBoat(state.Database, boatId)
	if boat == nil {
		//boat 不存在
		out.Result = clique_escort_api.HIJACK_BOAT_RESULT_NO_CLIQUE
		session.Send(out)
		return
	}
	if boat.CliqueId == playerCliqueInfo.CliqueId || boat.OwnerPid == state.PlayerId {
		//不能劫持本帮的船
		out.Result = clique_escort_api.HIJACK_BOAT_RESULT_CLIQUE_MEMBER
		session.Send(out)
		return
	}
	if boat.Status != int8(clique_escort_api.BOAT_STATUS_ESCORT) || isEscortFinished(now, boat) {
		//船不在可劫持的状态
		out.Result = clique_escort_api.HIJACK_BOAT_RESULT_CAN_NOT_HIJACK
		session.Send(out)
		return
	}

	out.Result = clique_escort_api.HIJACK_BOAT_RESULT_START_BATTLE
	out.BoatId = boat.Id
	boatOwnerInfo := global.GetPlayerInfo(boat.OwnerPid)
	out.BoatOwnerNick = boatOwnerInfo.PlayerNick
	out.BoatOwnerPid = boat.OwnerPid
	session.Send(out)
}

func recoverBoat(session *net.Session, boatId int64) {
	state := module.State(session)
	out := &clique_escort_api.RecoverBoat_Out{}

	boat := lookupBoat(state.Database, boatId)
	if boat == nil {
		out.Result = clique_escort_api.RECOVER_BOAT_RESULT_NO_BOAT
		session.Send(out)
		return
	}
	if boat.RecoverPid > 0 {
		out.Result = clique_escort_api.RECOVER_BOAT_RESULT_RECOVERING
		session.Send(out)
		return
	}
	if boat.Status != int8(clique_escort_api.BOAT_STATUS_HIJACK) || time.GetNowTime()-boat.HijackTimestamp >= clique_dat.HIJACK_TIME {
		//不在劫持中 或者 劫持已结束
		out.Result = clique_escort_api.RECOVER_BOAT_RESULT_CAN_NOT_RECOVER
		session.Send(out)
		return
	}
	var boatOwnerCliqueInfo *mdb.PlayerGlobalCliqueInfo
	state.Database.AgentExecute(boat.OwnerPid, func(agentDB *mdb.Database) {
		boatOwnerCliqueInfo = agentDB.Lookup.PlayerGlobalCliqueInfo(boat.OwnerPid)
	})

	playerCliqueInfo := state.Database.Lookup.PlayerGlobalCliqueInfo(state.PlayerId)
	if playerCliqueInfo == nil || boatOwnerCliqueInfo == nil || boatOwnerCliqueInfo.CliqueId != playerCliqueInfo.CliqueId {
		//自己所在帮派成员才能夺回 没有帮派就不能夺回
		out.Result = clique_escort_api.RECOVER_BOAT_RESULT_NO_PERMISSION
		session.Send(out)
		return
	}

	//船的劫持中状态仅用于客户端返回，数据库中 recover_pid + 劫持中 代表着这个状态
	//boat.Status = int8(clique_escort_api.BOAT_STATUS_RECOVER)
	boat.RecoverPid = state.PlayerId
	state.Database.Update.GlobalCliqueBoat(boat)
	out.Result = clique_escort_api.RECOVER_BOAT_RESULT_START_BATTLE
	out.BoatId = boat.Id
	out.HijackerPid = boat.HijackerPid
	hijackerInfo := global.GetPlayerInfo(boat.HijackerPid)
	out.HijackerNick = hijackerInfo.PlayerNick
	session.Send(out)
}

func listBoats(session *net.Session) {
	state := module.State(session)
	playerCliqueInfo := state.Database.Lookup.PlayerGlobalCliqueInfo(state.PlayerId)
	if playerCliqueInfo == nil || playerCliqueInfo.CliqueId <= 0 {
		return
	}

	var relatedBoats, randomBoats []*mdb.GlobalCliqueBoat
	now := time.GetNowTime()
	state.Database.Select.GlobalCliqueBoat(func(row *mdb.GlobalCliqueBoatRow) {
		//检查船只是否拥有对应的 timer 如果没有就设置一个
		//boat = updateBoatHook(row.GoObject()) //修正状态

		boat := row.GoObject()
		if isEscortFinished(now, row.GoObject()) {
			escortTimerCallback(state.Database, boat)
			return
		}
		EnsureBoatTimer(state.Database, boat)
		boat = updateBoatHook(boat)

		if boat.OwnerPid == state.PlayerId {
			relatedBoats = append(relatedBoats, boat)
			return
		}
		if boat.HijackerPid == state.PlayerId {
			relatedBoats = append(relatedBoats, boat)
			return
		}

		//处理我帮的船
		if module.CliqueRPC.IsCliqueMember(playerCliqueInfo.CliqueId, boat.OwnerPid) {
			relatedBoats = append(relatedBoats, boat)
			return
		}

		//其他可劫持的船
		if boat.Status == int8(clique_escort_api.BOAT_STATUS_ESCORT) {
			randomBoats = append(randomBoats, boat)
			return
		}
	})

	out := &clique_escort_api.ListBoats_Out{}
	for _, cboat := range relatedBoats {
		out.Boats = append(out.Boats, clique_escort_api.ListBoats_Out_Boats{
			Boat: mdbBoatToRespBoat(cboat),
		})
	}
	randomBoatNum := 50 - len(relatedBoats)
	if randomBoatNum > 0 && len(randomBoats) > 0 {
		idxs := rand.Perm(len(randomBoats))
		for _, idx := range idxs {
			out.Boats = append(out.Boats, clique_escort_api.ListBoats_Out_Boats{
				Boat: mdbBoatToRespBoat(randomBoats[idx]),
			})
			randomBoatNum--
			if randomBoatNum <= 0 {
				break
			}

		}
	}
	session.Send(out)
}

func getRandomBoat(session *net.Session) {
	state := module.State(session)
	playerEscortInfo := lookupEscortInfo(state.Database)
	if playerEscortInfo.EscortBoatType <= 0 {
		playerEscortInfo.EscortBoatType = clique_dat.GetRandomBoatId()
		state.Database.Update.PlayerGlobalCliqueEscort(playerEscortInfo)
	}
	session.Send(&clique_escort_api.GetRandomBoat_Out{
		BoatType: playerEscortInfo.EscortBoatType,
	})
}

func takeHijackAward(session *net.Session) {
	state := module.State(session)
	playerHijackInfo := lookupHijackInfo(state.Database)
	out := &clique_escort_api.TakeHijackAward_Out{}
	if playerHijackInfo.Status != int8(clique_escort_api.HIJACK_STATUS_FINISHED) {
		out.Ok = false
		session.Send(out)
		return
	}
	//发送奖励
	boatDat := clique_dat.GetBoatById(playerHijackInfo.HijackedBoatType)
	if boatDat.HijackLoseCoins > 0 {
		rpc.RemoteIncMoney(state.PlayerId, boatDat.HijackLoseCoins, player_dat.COINS, 0 /*TODO tlog*/, xdlog.ET_CLIQUE_BOAT, nil)
	}
	if boatDat.HijackLoseFame > 0 {
		rpc.RemoteAddFame(state.PlayerId, int64(boatDat.HijackLoseFame), player_dat.ARENA_SYSTEM, nil)
	}
	if boatDat.HijackLoseCliqueContrib > 0 {
		module.CliqueRPC.AddPlayerCliqueContrib(state.Database, boatDat.HijackLoseCliqueContrib)
	}
	playerHijackInfo.HijackedBoatType = 0
	playerHijackInfo.Status = int8(clique_escort_api.HIJACK_STATUS_NONE)
	state.Database.Update.PlayerGlobalCliqueHijack(playerHijackInfo)
	out.Ok = true
	session.Send(out)
}

func takeEscortAward(session *net.Session) {
	state := module.State(session)
	playerEscortInfo := lookupEscortInfo(state.Database)
	out := &clique_escort_api.TakeEscortAward_Out{}
	if playerEscortInfo.Status != int8(clique_escort_api.ESCORT_STATUS_FINISHED) {
		out.Ok = false
		session.Send(out)
		return
	}

	boatDat := clique_dat.GetBoatById(playerEscortInfo.EscortBoatType)
	awardCoins := boatDat.AwardCoins
	awardFame := boatDat.AwardFame
	awardContrib := boatDat.AwardCliqueContrib
	if playerEscortInfo.Hijacked == clique_dat.HIJACKED {
		awardCoins -= boatDat.HijackLoseCoins
		awardFame -= boatDat.HijackLoseFame
		awardContrib -= boatDat.HijackLoseCliqueContrib
	}
	if awardCoins > 0 {
		rpc.RemoteIncMoney(state.PlayerId, awardCoins, player_dat.COINS, 0 /*TODO tlog*/, xdlog.ET_CLIQUE_BOAT, nil)
	}
	if awardFame > 0 {
		rpc.RemoteAddFame(state.PlayerId, int64(awardFame), player_dat.ARENA_SYSTEM, nil)
	}
	if awardContrib > 0 {
		module.CliqueRPC.AddPlayerCliqueContrib(state.Database, awardContrib)
	}

	playerEscortInfo.EscortBoatType = 0
	playerEscortInfo.Hijacked = 0
	playerEscortInfo.Status = int8(clique_escort_api.ESCORT_STATUS_NONE)
	state.Database.Update.PlayerGlobalCliqueEscort(playerEscortInfo)
	out.Ok = true
	session.Send(out)
}
