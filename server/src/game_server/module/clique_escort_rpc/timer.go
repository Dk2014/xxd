package clique_escort_rpc

import (
	"core/debug"
	"core/log"
	"core/time"
	"fmt"
	"game_server/api/protocol/clique_escort_api"
	"game_server/dat/channel_dat"
	"game_server/dat/clique_dat"
	"game_server/global"
	"game_server/mdb"
	"game_server/module"
	"sync"
	gotime "time"
)

var g_escort_timer *TimerMgr

type TimerMgr struct {
	sync.Mutex
	//sync.RWMutex
	timers map[int64]*gotime.Timer //boat_id -> timer
}

func init() {
	g_escort_timer = &TimerMgr{}
	g_escort_timer.timers = make(map[int64]*gotime.Timer, 200)
}

func getBoatTImer(boatId int64) *gotime.Timer {
	g_escort_timer.Lock()
	defer g_escort_timer.Unlock()
	return g_escort_timer.timers[boatId]
}

//确保船只有一个timer
func EnsureBoatTimer(db *mdb.Database, boat *mdb.GlobalCliqueBoat) {
	if getBoatTImer(boat.Id) != nil {
		return
	}

	delaySecond := nextBoatAction(boat)
	if delaySecond <= 0 {
		escortTimerCallback(db, boat)
	} else {
		StartTimer(boat.Id, gotime.Duration(delaySecond)*gotime.Second)
	}

}

func StartTimer(boatId int64, delay gotime.Duration) {
	fmt.Printf("StartTimer callstack %v \n%s\n", delay, debug.Stack(1, "    "))
	g_escort_timer.Lock()
	defer g_escort_timer.Unlock()
	if delay < 0 {
		delay = 0
	}
	old_tm, has_old_timer := g_escort_timer.timers[boatId]
	if has_old_timer {
		fmt.Println("reset StartTimer for ", boatId, "delay", delay)
		old_tm.Reset(delay)
		return
	}
	g_escort_timer.timers[boatId] = gotime.AfterFunc(delay, func() {
		defer func() {
			if err := recover(); err != nil {
				log.Errorf(`CliqueBoat Timer
Error = %v
BoatId = %d
Stack = 
%s`,
					err,
					boatId,
					debug.Stack(1, "    "))
			}
		}()
		fmt.Println("set StartTimer for ", boatId, "delay", delay)
		mdb.Transaction(mdb.TRANS_TAG_ESCORT_BOAT_TIMER, func() {
			//TODO 这里会发生一些边界情况导致上一个timer无法取消，需要再想一下

			mdb.GlobalExecute(func(globalDB *mdb.Database) {
				boat := globalDB.Lookup.GlobalCliqueBoat(boatId)
				escortTimerCallback(globalDB, boat)
			})
		})
	})

}

func StopTimer(boatId int64) {
	g_escort_timer.Lock()
	defer g_escort_timer.Unlock()
	old_tm, has_old_timer := g_escort_timer.timers[boatId]
	if has_old_timer {
		old_tm.Stop()
		delete(g_escort_timer.timers, boatId)
	}
}

//func escortTimerCallback(globalDB *mdb.Database, boatId int64) {
func escortTimerCallback(globalDB *mdb.Database, boat *mdb.GlobalCliqueBoat) {
	//boat := globalDB.Lookup.GlobalCliqueBoat(boatId)
	if boat == nil {
		return
	}
	now := time.GetNowTime()
	fmt.Println("escortTimerCallback", boat.Id, boat.BoatType, now)

	//完成劫持 ?
	if boat.Status == int8(clique_escort_api.BOAT_STATUS_HIJACK) ||
		boat.Status == int8(clique_escort_api.BOAT_STATUS_RECOVER) {
		var delt int64 //预计剩余劫持时间
		delt = boat.HijackTimestamp + clique_dat.HIJACK_TIME - now
		if delt <= 0 {
			fmt.Println("escortTimerCallback", boat.Id, "hijack finished")
			boat.Status = int8(clique_escort_api.BOAT_STATUS_HIJACK_FINISH)
			boat.EscortStartTimestamp = boat.HijackTimestamp + clique_dat.HIJACK_TIME
			globalDB.Update.GlobalCliqueBoat(boat)
			boatDat := clique_dat.GetBoatById(boat.BoatType)
			globalDB.AgentExecute(boat.HijackerPid, func(agentDB *mdb.Database) {
				playerHijackInfo := lookupHijackInfo(agentDB)
				fmt.Println("劫持信息：", playerHijackInfo.Pid, playerHijackInfo.HijackedBoatType, playerHijackInfo.Status)
				playerHijackInfo.Status = int8(clique_escort_api.HIJACK_STATUS_FINISHED)
				agentDB.Update.PlayerGlobalCliqueHijack(playerHijackInfo)
				awardCliqueContrib := boatDat.HijackLoseCliqueContrib
				playerCliqueInfo := agentDB.Lookup.PlayerGlobalCliqueInfo(agentDB.PlayerId())
				if playerCliqueInfo == nil || playerCliqueInfo.CliqueId <= 0 {
					awardCliqueContrib = 0
				}

				//完成劫持动态
				newMessage(agentDB, channel_dat.MessageBoatHijackFinished{
					Coins:   channel_dat.ParamString{fmt.Sprintf("%d", boatDat.HijackLoseCoins)},
					Fame:    channel_dat.ParamString{fmt.Sprintf("%d", boatDat.HijackLoseFame)},
					Contrib: channel_dat.ParamString{fmt.Sprintf("%d", awardCliqueContrib)},
				})
			})
			globalDB.AgentExecute(boat.OwnerPid, func(agentDB *mdb.Database) {
				playerEscortInfo := lookupEscortInfo(agentDB)
				playerEscortInfo.Hijacked = clique_dat.HIJACKED
				agentDB.Update.PlayerGlobalCliqueEscort(playerEscortInfo)
				//被劫持完成动态
				hijacker := global.GetPlayerInfo(boat.HijackerPid)
				awardCliqueContrib := boatDat.HijackLoseCliqueContrib
				playerCliqueInfo := agentDB.Lookup.PlayerGlobalCliqueInfo(agentDB.PlayerId())
				if playerCliqueInfo == nil || playerCliqueInfo.CliqueId <= 0 {
					awardCliqueContrib = 0
				}
				newMessage(agentDB, channel_dat.MessageBoatHijackingFinished{
					Hijacker: channel_dat.ParamPlayer{Nick: hijacker.PlayerNick, Pid: hijacker.PlayerId},
					Coins:    channel_dat.ParamString{fmt.Sprintf("%d", boatDat.HijackLoseCoins)},
					Fame:     channel_dat.ParamString{fmt.Sprintf("%d", boatDat.HijackLoseFame)},
					Contrib:  channel_dat.ParamString{fmt.Sprintf("%d", awardCliqueContrib)},
				})
			})
			if session, online := module.Player.GetPlayerOnline(boat.HijackerPid); online {
				session.Send(&clique_escort_api.NotifyHijackFinished_Out{})
				session.Send(&clique_escort_api.NotifyBoatStatusChange_Out{
					Boat:   mdbBoatToRespBoat(boat),
					Change: clique_escort_api.BOAT_STATUS_CHANGE_HIJACK_FINISHED,
				})
			}
			if session, online := module.Player.GetPlayerOnline(boat.OwnerPid); online {
				session.Send(&clique_escort_api.NotifyBoatStatusChange_Out{
					Boat:   mdbBoatToRespBoat(boat),
					Change: clique_escort_api.BOAT_STATUS_CHANGE_MY_BOAT_HIJACKED,
				})
			}

		} else {
			//劫持中
			//设置一个 timer 在劫持结束的时候启动
			fmt.Println("劫持: start timer after", delt)
			StartTimer(boat.Id, gotime.Duration(delt)*gotime.Second)
			return
		}
	}

	//运送完？
	if isEscortFinished(now, boat) {
		// update player escort info
		fmt.Println("escortTimerCallback", boat.Id, "escort finished")
		globalDB.AgentExecute(boat.OwnerPid, func(agentDB *mdb.Database) {
			playerEscortInfo := lookupEscortInfo(agentDB)
			playerEscortInfo.Status = int8(clique_escort_api.ESCORT_STATUS_FINISHED)
			agentDB.Update.PlayerGlobalCliqueEscort(playerEscortInfo)

			//完成护送动态
			boatDat := clique_dat.GetBoatById(boat.BoatType)
			awardCoins, awardFame, awardCliqueContrib := boatDat.AwardCoins, boatDat.AwardFame, boatDat.AwardCliqueContrib
			if playerEscortInfo.Hijacked == clique_dat.HIJACKED {
				awardCoins -= boatDat.HijackLoseCoins
				awardFame -= boatDat.HijackLoseFame
				awardCliqueContrib -= boatDat.HijackLoseCliqueContrib
			}
			playerCliqueInfo := agentDB.Lookup.PlayerGlobalCliqueInfo(agentDB.PlayerId())
			if playerCliqueInfo == nil || playerCliqueInfo.CliqueId <= 0 {
				awardCliqueContrib = 0
			}
			newMessage(agentDB, channel_dat.MessageBoatEscortFinished{
				Coins:   channel_dat.ParamString{fmt.Sprintf("%d", awardCoins)},
				Fame:    channel_dat.ParamString{fmt.Sprintf("%d", awardFame)},
				Contrib: channel_dat.ParamString{fmt.Sprintf("%d", awardCliqueContrib)},
			})
		})

		// delete boat
		globalDB.Delete.GlobalCliqueBoat(boat)
		StopTimer(boat.Id)

		//notify
		if session, online := module.Player.GetPlayerOnline(boat.OwnerPid); online {
			session.Send(&clique_escort_api.NotifyEscortFinished_Out{})
			session.Send(&clique_escort_api.NotifyBoatStatusChange_Out{
				Boat:   mdbBoatToRespBoat(boat),
				Change: clique_escort_api.BOAT_STATUS_CHANGE_ESCORT_FINISHED,
			})
		}
	} else {
		//运送中
		var delt int64 //预计剩余运送时间
		delt = now + clique_dat.ESCORT_TIME - int64(boat.TotalEscortTime) - boat.EscortStartTimestamp
		StartTimer(boat.Id, gotime.Duration(delt)*gotime.Second)
		return
	}
}
