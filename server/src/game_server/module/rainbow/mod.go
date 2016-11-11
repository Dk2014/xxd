package rainbow

import (
	"core/fail"
	"core/net"
	"core/time"
	"game_server/dat/channel_dat"
	"game_server/dat/ghost_dat"
	"game_server/dat/mail_dat"
	"game_server/dat/mission_dat"
	"game_server/dat/quest_dat"
	"game_server/dat/rainbow_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/module/rpc"
	"game_server/tlog"
	"game_server/xdlog"
	"math/rand"
)

func init() {
	module.Rainbow = RainbowMod{}
}

type RainbowMod struct{}

//为玩家初始化彩虹关卡所需的数据
func (mod RainbowMod) OpenFunc(db *mdb.Database) {
	db.Insert.PlayerRainbowLevel(&mdb.PlayerRainbowLevel{
		Pid:      db.PlayerId(),
		Segment:  rainbow_dat.INIT_SEGMENT_NUM,
		Order:    rainbow_dat.INIT_LEVEL_ORDER,
		Status:   rainbow_dat.RAINBOW_LEVEL_STATUS_NEVER_PASS,
		ResetNum: rainbow_dat.DAILY_RESET_NUM,
	})

	db.Insert.PlayerRainbowLevelStateBin(&mdb.PlayerRainbowLevelStateBin{
		Pid: db.PlayerId(),
	})
}

func (mod RainbowMod) StartRainbowLevel(session *net.Session, enemyId int32) {
	//module.Battle.NewBattleForRainbowLevel(state)
	state := module.State(session)
	fail.When(state.MissionLevelState == nil, "MissionLevelState is nil")
	fail.When(state.RainbowLevelState == nil, "RainbowLevelState is nil")

	levelEnemy := mission_dat.GetMissionLevelEnemyById(int32(enemyId))
	fail.When(levelEnemy.MissionLevelId != state.MissionLevelState.LevelId, "incorrect enemyId")
	state.MissionLevelState.EnemyId = enemyId

	state.Battle = module.Battle.NewBattleForRainbowLevel(session, enemyId)
}

func (mod RainbowMod) BattleWin(state *module.SessionState, xdEventType int32) {
	//更新通关状态
	playerRainbow := state.Database.Lookup.PlayerRainbowLevel(state.PlayerId)
	playerRainbow.Status = rainbow_dat.RAINBOW_LEVEL_STATUS_NEVER_AWARD

	//首次打通一段
	if playerRainbow.Order == rainbow_dat.LEVEL_NUM_PER_SEGMENT && playerRainbow.MaxPassSegment < playerRainbow.Segment {
		playerRainbow.MaxPassSegment = playerRainbow.Segment
		if awardGhostId := rainbow_dat.GetRainbowSegmentAward(playerRainbow.Segment); awardGhostId > 0 {
			ghost := ghost_dat.GetGhost(awardGhostId)
			if ghost.Quality == ghost_dat.COLOR_GOLD {
				rpc.RemoteAddWorldChannelTplMessage(state.PlayerId, []channel_dat.MessageTpl{
					channel_dat.MessageRainbowLevelGhost{
						Level:  channel_dat.ParamString{rainbow_dat.SegmentNum2SegmentName(playerRainbow.Segment)},
						Player: channel_dat.ParamPlayer{state.PlayerNick, state.PlayerId},
						Item:   channel_dat.ParamItem{mail_dat.ATTACHMENT_GHOST, ghost.Id},
					},
				})
			}
			module.Ghost.AddGhost(state, awardGhostId, tlog.IFR_RAINBOW, xdEventType)
		}
	}

	if playerRainbow.BestSegment < playerRainbow.Segment ||
		(playerRainbow.Segment == playerRainbow.BestSegment && playerRainbow.BestOrder < playerRainbow.Order) {
		playerRainbow.BestSegment = playerRainbow.Segment
		playerRainbow.BestOrder = playerRainbow.Order
		playerRainbow.BestRecordTimestamp = time.GetNowTime()
	}

	state.Database.Update.PlayerRainbowLevel(playerRainbow)

	levelId := rainbow_dat.GetRainbowLevelId(playerRainbow.Segment, playerRainbow.Order)

	//计算奖励
	state.RainbowLevelState.AwardBoxIndex = calAward(levelId)

	//保存玩家战斗相关信息
	state.RainbowLevelState.SyncMissionLevelState(state)
	state.SaveRainbowLevelState()

	module.Quest.RefreshDailyQuest(state, quest_dat.DAILY_QUEST_CLASS_LIMIT_LEVEL)
	tlog.PlayerMissionFlowLog(state.Database, levelId, tlog.FINISH)
	xdlog.MissionLog(state.Database, levelId, xdlog.MA_FINISH)
}

func (mod RainbowMod) BattleLose(state *module.SessionState) {
	//把 state.MissionLevelState 的数据导到 state.RainbowLevelState 并写入数据库
	state.RainbowLevelState.SyncMissionLevelState(state)
	state.SaveRainbowLevelState()
}

//func (mod RainbowMod) NextLevel(state *module.SessionState) (int8, int8) {
//	rainbowLevel := state.Database.Lookup.PlayerRainbowLevel(state.PlayerId)
//
//	return rainbowLevel.Segment == rainbow_dat.MaxRainbowLevelSegment && rainbowLevel.Order == rainbow_dat.LEVEL_NUM_PER_SEGMENT
//}

func calAward(levelId int32) (awards []int8) {
	awardConfigs := rainbow_dat.GetRainbowLevelAward(levelId)
	takedAwardOrder := make(map[int8]int8, 4)

	var totalChance int32
	var random int32
	for _, awardConfig := range awardConfigs {
		totalChance += int32(awardConfig.AwardChance)
	}

	for i := 1; i <= 5; i++ {
		random = rand.Int31n(totalChance) + 1
		for _, config := range awardConfigs {
			if _, taked := takedAwardOrder[config.Order]; taked {
				continue
			}
			if random <= int32(config.AwardChance) {
				takedAwardOrder[config.Order] = 1
				awards = append(awards, config.Order)
				totalChance -= int32(config.AwardChance)
				break
			}
			random -= int32(config.AwardChance)
		}
	}
	return awards
}
