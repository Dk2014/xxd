package clique_escort_rpc

import (
	"core/i18l"
	"core/time"
	"fmt"
	"game_server/api/protocol/clique_escort_api"
	"game_server/dat/channel_dat"
	"game_server/dat/clique_dat"
	"game_server/global"
	"game_server/mdb"
	"game_server/module"
	gotime "time"
)

func recoverBattleLose(globalDB *mdb.Database, fighterPid, boatId int64) {
	boat := lookupBoat(globalDB, boatId)
	if boat == nil {
		return
	}
	if boat.RecoverPid == fighterPid {
		boat.RecoverPid = 0
	}
	globalDB.Update.GlobalCliqueBoat(boat)
}

func recoverBattleWin(globalDB *mdb.Database, fighterPid, boatId int64) {
	boat := lookupBoat(globalDB, boatId)
	out := &clique_escort_api.NotifyRecoverBattleWin_Out{
		BoatId: boatId,
	}
	fighterSession, fighterOnline := module.Player.GetPlayerOnline(fighterPid)
	ownerSession, ownerOnline := module.Player.GetPlayerOnline(boat.OwnerPid)
	now := time.GetNowTime()

	if boat == nil || !isHijacking(now, boat) || boat.RecoverPid != fighterPid {
		if fighterOnline {
			out.Result = clique_escort_api.RECOVER_BATTLE_WIN_RESULT_BOAT_EXPIRE
			fighterSession.Send(out)
		}
		return
	}

	isCliqueMember := false
	globalDB.AgentExecute(fighterPid, func(agentDB *mdb.Database) {
		playerCliqueInfo := agentDB.Lookup.PlayerGlobalCliqueInfo(fighterPid)
		if playerCliqueInfo == nil || playerCliqueInfo.CliqueId <= 0 {
			return
		}
		isCliqueMember = module.CliqueRPC.IsCliqueMember(playerCliqueInfo.CliqueId, boat.OwnerPid)
	})
	if !isCliqueMember {
		if fighterOnline {
			out.Result = clique_escort_api.RECOVER_BATTLE_WIN_RESULT_NO_PERMISSION
			fighterSession.Send(out)
		}
		return
	}

	recoverFighter := global.GetPlayerInfo(fighterPid)
	owner := global.GetPlayerInfo(boat.OwnerPid)
	out.OwnerNick = owner.PlayerNick
	//hijacker := global.GetPlayerInfo(boat.HijackerPid)

	//通知劫持者
	globalDB.AgentExecute(boat.HijackerPid, func(agentDB *mdb.Database) {
		//恢复劫持状态
		hijackInfo := lookupHijackInfo(agentDB)
		hijackInfo.HijackedBoatType = 0
		hijackInfo.Status = int8(clique_escort_api.HIJACK_STATUS_NONE)
		agentDB.Update.PlayerGlobalCliqueHijack(hijackInfo)

		newMessage(agentDB, channel_dat.MessageBoatRecovered{
			Fighter: channel_dat.ParamPlayer{
				Nick: recoverFighter.PlayerNick,
				Pid:  recoverFighter.PlayerId,
			},
		})
	})
	//通知船主
	var ownerClique *mdb.GlobalClique
	globalDB.AgentExecute(boat.OwnerPid, func(agentDB *mdb.Database) {
		ownerCliqueInfo := agentDB.Lookup.PlayerGlobalCliqueInfo(boat.OwnerPid)
		ownerClique = agentDB.Lookup.GlobalClique(ownerCliqueInfo.CliqueId)

		newMessage(agentDB, channel_dat.MessageBoatRecoveredByHero{
			Fighter: channel_dat.ParamPlayer{
				Nick: recoverFighter.PlayerNick,
				Pid:  recoverFighter.PlayerId,
			},
		})
	})
	// 帮派内动态
	module.CliqueRPC.AddCliqueNews(ownerClique.Id, channel_dat.MessageCliqueBoatRecover{
		Fighter: channel_dat.ParamPlayer{
			Nick: recoverFighter.PlayerNick,
			Pid:  recoverFighter.PlayerId,
		},
		BoatOwner: channel_dat.ParamPlayer{
			Nick: owner.PlayerNick,
			Pid:  owner.PlayerId,
		},
	})

	boat.Status = int8(clique_escort_api.BOAT_STATUS_ESCORT)
	boat.EscortStartTimestamp = now
	boat.HijackerPid = 0
	boat.HijackTimestamp = 0
	boat.RecoverPid = 0
	globalDB.Update.GlobalCliqueBoat(boat)

	delt := now + clique_dat.ESCORT_TIME - int64(boat.TotalEscortTime) - boat.EscortStartTimestamp
	StartTimer(boat.Id, gotime.Duration(delt)*gotime.Second)

	if fighterOnline {
		out.Result = clique_escort_api.RECOVER_BATTLE_WIN_RESULT_SUCCESS
		fighterSession.Send(out)
		fighterSession.Send(&clique_escort_api.NotifyBoatStatusChange_Out{
			Boat:   mdbBoatToRespBoat(boat),
			Change: clique_escort_api.BOAT_STATUS_CHANGE_HIJACKED_BOAT_RECOVERED,
		})
	}
	if ownerOnline {
		ownerSession.Send(&clique_escort_api.NotifyBoatStatusChange_Out{
			Boat:   mdbBoatToRespBoat(boat),
			Change: clique_escort_api.BOAT_STATUS_CHANGE_MY_BOAT_RECOVERED,
		})
	}
}

func hijackBattleWin(globalDB *mdb.Database, fighterPid, boatId int64) {
	fmt.Println("hijackBattleWin", fighterPid, boatId)
	boat := lookupBoat(globalDB, boatId)

	fighterSession, fighterOnline := module.Player.GetPlayerOnline(fighterPid)
	out := &clique_escort_api.NotifyHijackBattleWin_Out{
		BoatId: boatId,
	}

	now := time.GetNowTime()
	if boat == nil || isEscortFinished(now, boat) {
		out.Result = clique_escort_api.HIJACK_BATTLE_WIN_RESULT_ESCORT_FINISHED
		if fighterOnline {
			fighterSession.Send(out)
		}
		return
	}

	ownerSession, ownerOnline := module.Player.GetPlayerOnline(boat.OwnerPid)

	if boat.Status != int8(clique_escort_api.BOAT_STATUS_ESCORT) {
		//不在可劫持状态
		if fighterOnline {
			out.Result = clique_escort_api.HIJACK_BATTLE_WIN_RESULT_HIJACKED
			fighterSession.Send(out)
		}
		return
	}

	canHijack := true
	globalDB.AgentExecute(fighterPid, func(agentDB *mdb.Database) {
		playerCliqueInfo := agentDB.Lookup.PlayerGlobalCliqueInfo(fighterPid)
		if playerCliqueInfo == nil || playerCliqueInfo.CliqueId <= 0 {
			canHijack = false
		}
		canHijack = !module.CliqueRPC.IsCliqueMember(playerCliqueInfo.CliqueId, boat.OwnerPid)
	})
	if !canHijack {
		//没有帮派不能劫持 或者 和劫持者同一个帮派
		if fighterOnline {
			out.Result = clique_escort_api.HIJACK_BATTLE_WIN_RESULT_NO_PERMISSION
			fighterSession.Send(out)
		}
		return
	}
	var hijackerClique *mdb.GlobalClique
	globalDB.AgentExecute(fighterPid, func(agentDB *mdb.Database) {
		hijackInfo := lookupHijackInfo(agentDB)
		hijackInfo.HijackedBoatType = boat.BoatType
		hijackInfo.Status = int8(clique_escort_api.HIJACK_STATUS_HIJACK)
		hijackInfo.DailyHijackNum++
		hijackInfo.DailyHijackTimestamp = now
		agentDB.Update.PlayerGlobalCliqueHijack(hijackInfo)

		hijackerCliqueInfo := agentDB.Lookup.PlayerGlobalCliqueInfo(fighterPid)
		hijackerClique = agentDB.Lookup.GlobalClique(hijackerCliqueInfo.CliqueId)
	})

	owner := global.GetPlayerInfo(boat.OwnerPid)
	hijacker := global.GetPlayerInfo(fighterPid)
	out.OwnerNick = owner.PlayerNick

	//通知被劫持玩家
	var ownerCliqueId int64
	globalDB.AgentExecute(boat.OwnerPid, func(agentDB *mdb.Database) {
		ownerCliqueInfo := agentDB.Lookup.PlayerGlobalCliqueInfo(boat.OwnerPid)
		ownerCliqueId = ownerCliqueInfo.CliqueId
		newMessage(agentDB, channel_dat.MessageBoatHijacking{
			HijackerClique: channel_dat.ParamClique{Name: hijackerClique.Name, Id: hijackerClique.Id},
			Hijacker:       channel_dat.ParamPlayer{Nick: hijacker.PlayerNick, Pid: hijacker.PlayerId},
		})
	})
	//帮派内动态
	module.CliqueRPC.AddCliqueNews(ownerCliqueId, channel_dat.MessageCliqueBoatHijacked{
		Clique:    channel_dat.ParamClique{Name: hijackerClique.Name, Id: hijackerClique.Id},
		Hijacker:  channel_dat.ParamPlayer{Nick: hijacker.PlayerNick, Pid: hijacker.PlayerId},
		DummyLink: channel_dat.ParamCliqueBoat{BoatId: boat.Id, OwnerPid: boat.OwnerPid, BoatName: i18l.T.Tran("立即夺回")},
	})

	boat.TotalEscortTime += int16(now - boat.EscortStartTimestamp)
	boat.HijackTimestamp = now
	boat.HijackerPid = fighterPid
	boat.Status = int8(clique_escort_api.BOAT_STATUS_HIJACK)
	globalDB.Update.GlobalCliqueBoat(boat)
	StartTimer(boat.Id, gotime.Duration(clique_dat.HIJACK_TIME)*gotime.Second)

	if fighterOnline {
		out.Result = clique_escort_api.HIJACK_BATTLE_WIN_RESULT_SUCCESS
		fighterSession.Send(out)
		fighterSession.Send(&clique_escort_api.NotifyBoatStatusChange_Out{
			Boat:   mdbBoatToRespBoat(boat),
			Change: clique_escort_api.BOAT_STATUS_CHANGE_MY_BOAT_HIJACKING,
		})
	}
	if ownerOnline {
		ownerSession.Send(&clique_escort_api.NotifyBoatStatusChange_Out{
			Boat:   mdbBoatToRespBoat(boat),
			Change: clique_escort_api.BOAT_STATUS_CHANGE_MY_BOAT_HIJACKING,
		})
	}
}
