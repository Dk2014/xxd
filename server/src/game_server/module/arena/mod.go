package arena

import (
	"core/net"
	"core/time"
	"game_server/api/protocol/arena_api"
	"game_server/dat/arena_award_box_dat"
	"game_server/dat/event_dat"
	"game_server/dat/player_dat"
	"game_server/dat/quest_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/module/rpc"
	"game_server/tlog"
)

func init() {
	module.Arena = ArenaMod{}
	module.PrepareStoreEvent.Regisiter(PerpareStoreHandler)
}

func PerpareStoreHandler(session *net.Session) {
	state := module.State(session)

	if state.ArenaState != nil {
		if state.Battle != nil {
			state.Battle.LeaveBattle(session)
		}
		state.ArenaState = nil
		state.Battle = nil
	}
}

type ArenaMod struct {
}

func (mod ArenaMod) OpenFunc(db *mdb.Database) {
	db.Insert.PlayerArena(&mdb.PlayerArena{
		Pid:            db.PlayerId(),
		DailyNum:       0,
		FailedCdTime:   0,
		BattleTime:     0,
		WinTimes:       0,
		DailyAwardItem: 0,
		NewRecordCount: 0,
	})
	//比武场活动初始值
	if session, ok := module.Player.GetPlayerOnline(db.PlayerId()); ok {
		state := module.State(session)
		rpc.RemoteGetArenaRank(state.PlayerId, func(reply *rpc.Reply_GlobalArenaGetArenaRank) {
			mdb.GlobalExecute(func(globalDB *mdb.Database) {
				globalDB.AgentExecute(state.PlayerId, func(agentDB *mdb.Database) {
					module.Arena.InitPlayerArenaRank(agentDB, state.PlayerId)
					state.EventsState.AddEventAwardState(agentDB, event_dat.EVENT_ARENA_RANK_AWARDS, reply.PlayerRank, 0)
				})
			})
		})
	}
}

func (mod ArenaMod) BattleDoWin(session *net.Session) {
	state := module.State(session)

	//module.Item.AddItem(state.Database, item_dat.ITEM_YINGJIEGUOSHI, awardItemNum, tlog.IFR_ARENA_BATTLE_WIN)

	playerArena := state.Database.Lookup.PlayerArena(state.PlayerId)
	if playerArena.WinTimes > 0 {
		playerArena.WinTimes += 1
	} else {
		playerArena.WinTimes = 1
	}
	//playerArena.DailyAwardItem += int32(awardItemNum)
	playerArena.DailyFame += player_dat.ARENA_WIN_AWARD_FAME
	module.Player.AddFame(state.Database, player_dat.ARENA_WIN_AWARD_FAME)
	state.Database.Update.PlayerArena(playerArena)

	// 同步玩家比武场趋势到互动服
	rpc.RemoteUpdatePlayerArenaTrendWin(playerArena.Pid, playerArena.WinTimes)

	arenaState := state.ArenaState
	newRank, targetNewRank := arenaState.PlayerRank, arenaState.TargetPlayerRank

	// 玩家自己排名落后于对手排名，交换排名
	if arenaState.PlayerRank > arenaState.TargetPlayerRank {

		rpc.RemoteSwapPlayerRank(state.PlayerId, arenaState.TargetInfo.PlayerId, arenaState.PlayerRank, arenaState.TargetPlayerRank,
			func(reply *rpc.Reply_GlobalArenaSwapPlayerRank) {
				if reply.Result == false {
					// 通知排名已被抢
					session.Send(&arena_api.NotifyLoseRank_Out{})
				} else {
					newRank, targetNewRank = arenaState.TargetPlayerRank, arenaState.PlayerRank
					//比武场活动更新
					module.UpdateEventArenaRank(state, newRank)
					//JSON 比武场活动更新
					module.Event.UpdateJsonEventArenaRank(state.Database, newRank)
				}
				mdb.GlobalExecute(func(globalDB *mdb.Database) {
					globalDB.AgentExecute(state.PlayerId, func(agentDB *mdb.Database) {
						addRecord(state, agentDB, arenaState, arena_award_box_dat.ARENA_ATTACK_SUCC, newRank, targetNewRank)
					})
				})

			})
	} else {
		addRecord(state, state.Database, arenaState, arena_award_box_dat.ARENA_ATTACK_SUCC, newRank, targetNewRank)
	}

	state.ArenaState = nil

	module.Quest.RefreshDailyQuest(state, quest_dat.DAILY_QUEST_CLASS_AREAN)
}

func (mod ArenaMod) BattleDoLose(session *net.Session) {
	state := module.State(session)

	//module.Item.AddItem(state.Database, item_dat.ITEM_YINGJIEGUOSHI, arena_award_box_dat.AWARD_ANYINGGUOSHI_LOSE_LOW_RANK, tlog.IFR_ARENA_BATTLE_LOSE)

	playerArena := state.Database.Lookup.PlayerArena(state.PlayerId)

	playerArena.FailedCdTime = time.GetNowTime() + arena_award_box_dat.LOSE_CD_TIME_SECONDS
	//playerArena.DailyAwardItem += arena_award_box_dat.AWARD_ANYINGGUOSHI_LOSE_LOW_RANK
	playerArena.DailyFame += player_dat.ARENA_LOSE_AWARD_FAME
	module.Player.AddFame(state.Database, player_dat.ARENA_LOSE_AWARD_FAME)

	if playerArena.WinTimes > 0 {
		playerArena.WinTimes = 0
	}
	state.Database.Update.PlayerArena(playerArena)

	// 同步玩家比武场趋势到互动服; 失败是下降趋势(-1)
	rpc.RemoteUpdatePlayerArenaTrendWin(playerArena.Pid, -1)

	session.Send(&arena_api.NotifyFailedCdTime_Out{CdTime: playerArena.FailedCdTime})
	addRecord(state, state.Database, state.ArenaState, arena_award_box_dat.ARENA_ATTACK_FAILED, state.ArenaState.PlayerRank, state.ArenaState.TargetPlayerRank)

	state.ArenaState = nil

	module.Quest.RefreshDailyQuest(state, quest_dat.DAILY_QUEST_CLASS_AREAN)
}

func (mod ArenaMod) GetPlayerArenaRank(db *mdb.Database, callback func(int32)) {
	if db.PlayerId() > 0 {
		rpc.RemoteGetArenaRank(db.PlayerId(), func(reply *rpc.Reply_GlobalArenaGetArenaRank) {
			mdb.GlobalExecute(func(globalDB *mdb.Database) {
				globalDB.AgentExecute(db.PlayerId(), func(agentDB *mdb.Database) {
					module.Arena.InitPlayerArenaRank(agentDB, db.PlayerId())
					callback(reply.PlayerRank)
				})
			})
		})
	}
}

func addRecord(state *module.SessionState, db *mdb.Database, arenaState *module.PlayerArenaState, playerMod int8, newRank, targetNewRank int32) {
	nowTime := time.GetNowTime()
	// 玩家增加战报
	playerRecord := &mdb.PlayerArenaRecord{
		Pid:            state.PlayerId,
		Mode:           playerMod,
		OldRank:        arenaState.PlayerRank,
		NewRank:        newRank,
		FightNum:       arenaState.PlayerFightNum,
		TargetPid:      arenaState.TargetInfo.PlayerId,
		TargetNick:     string(arenaState.TargetInfo.PlayerNick),
		TargetOldRank:  arenaState.TargetPlayerRank,
		TargetNewRank:  targetNewRank,
		TargetFightNum: arenaState.TargetInfo.FightNum,
		RecordTime:     nowTime,
	}

	db.Insert.PlayerArenaRecord(playerRecord)

	tlog.PlayerPvpFlowLog(db, int32(arenaState.PlayerRank), int32(newRank), int32(playerMod))

	var targetMod int8
	if playerMod == arena_award_box_dat.ARENA_ATTACK_FAILED {
		targetMod = arena_award_box_dat.ARENA_ATTACKED_SUCC
	} else {
		targetMod = arena_award_box_dat.ARENA_ATTACKED_FAILED
	}

	// 对手增加战报
	rpc.RemoteArenaAddRecord(&mdb.PlayerArenaRecord{
		Pid:            arenaState.TargetInfo.PlayerId,
		Mode:           targetMod,
		OldRank:        arenaState.TargetPlayerRank,
		NewRank:        targetNewRank,
		FightNum:       arenaState.TargetInfo.FightNum,
		TargetPid:      state.PlayerId,
		TargetNick:     string(state.PlayerNick),
		TargetOldRank:  arenaState.PlayerRank,
		TargetNewRank:  newRank,
		TargetFightNum: arenaState.PlayerFightNum,
		RecordTime:     nowTime,
	})
}

func (mod ArenaMod) InitPlayerArenaRank(db *mdb.Database, pid int64) {
	if db.Lookup.PlayerArenaRank(pid) == nil {
		db.Insert.PlayerArenaRank(&mdb.PlayerArenaRank{
			Pid:   pid,
			Rank:  0,
			Rank1: 0,
			Rank2: 0,
			Rank3: 0,
			Time:  time.GetNowTime(),
		})
	}
}

/*
	更新宝箱状态
		如果当前排名的更新时间超过21点（每日21点结算排名）：
			- 前天排名保持到今天(rank)，重置排名时间；
			- 如果前天奖励未领取，则移到前一天的奖励宝箱中
			- 如果奖励已领取，rank1为0
	@playerRank - mdb PlayerArenaRank struct
	@rank	    - 当前排名 from global_arena_rank
*/
func (mod ArenaMod) RefreshAwardBox(playerRank *mdb.PlayerArenaRank, rank int32) bool {
	if !time.IsInPointHour(player_dat.RESET_ARENA_AWARD_BOX_IN_HOUR, playerRank.Time) {
		// rank2 move to rank3, rank1 move to rank2, rank move to rank1
		playerRank.Rank3, playerRank.Rank2, playerRank.Rank1 = playerRank.Rank2, playerRank.Rank1, playerRank.Rank
		playerRank.Rank = rank
		playerRank.Time = time.GetNowTime()
		return true
	}
	return false
}
