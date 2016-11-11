package draw

import (
	"core/fail"
	"core/net"
	coretime "core/time"
	"game_server/api/protocol/draw_api"
	"game_server/dat/heart_draw_dat"
	"game_server/dat/player_dat"
	"game_server/dat/quest_dat"
	"game_server/dat/vip_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/tlog"
	"math/rand"
)

func getPlayerHeartDraw(db *mdb.Database, drawType int8) (playerHeartDraw *mdb.PlayerHeartDraw) {
	db.Select.PlayerHeartDraw(func(row *mdb.PlayerHeartDrawRow) {
		if row.DrawType() == drawType {
			playerHeartDraw = row.GoObject()
			row.Break()
		}
	})

	fail.When(playerHeartDraw == nil, "can't get playerHeartDraw")

	return
}

func getHeartDrawInfo(session *net.Session, drawType int8, awardRecord bool) {
	fail.When(drawType != heart_draw_dat.DRAW_TYPE_WHEEL, "error heart draw type")
	state := module.State(session)
	playerHeartDraw := getPlayerHeartDraw(state.Database, drawType)

	// 隔天重置抽奖次数
	if !coretime.IsInPointHour(player_dat.RESET_HEART_DRAW_TIMES_IN_HOUR, playerHeartDraw.DrawTime) {
		playerHeartDraw.DailyNum = 0
		state.Database.Update.PlayerHeartDraw(playerHeartDraw)
	}

	playerHeart := state.Database.Lookup.PlayerHeart(state.PlayerId)
	rsp := &draw_api.GetHeartDrawInfo_Out{
		Hearts:   playerHeart.Value,
		DailyNum: playerHeartDraw.DailyNum,
	}

	if awardRecord {
		switch drawType {
		case heart_draw_dat.DRAW_TYPE_CARD:
			state.Database.Select.PlayerHeartDrawCardRecord(func(row *mdb.PlayerHeartDrawCardRecordRow) {
				rsp.AwardRecord = append(rsp.AwardRecord, draw_api.GetHeartDrawInfo_Out_AwardRecord{
					Award: draw_api.AwardInfo{
						AwardType: row.AwardType(),
						AwardNum:  row.AwardNum(),
						ItemId:    row.ItemId(),
						DrawTime:  row.DrawTime(),
					},
				})
			})

		case heart_draw_dat.DRAW_TYPE_WHEEL:
			state.Database.Select.PlayerHeartDrawWheelRecord(func(row *mdb.PlayerHeartDrawWheelRecordRow) {
				rsp.AwardRecord = append(rsp.AwardRecord, draw_api.GetHeartDrawInfo_Out_AwardRecord{
					Award: draw_api.AwardInfo{
						AwardType: row.AwardType(),
						AwardNum:  row.AwardNum(),
						ItemId:    row.ItemId(),
						DrawTime:  row.DrawTime(),
					},
				})
			})
		}
	}

	session.Send(rsp)
}

func heartDraw(session *net.Session, drawType int8, xdEventType int32) {
	state := module.State(session)
	heartDraw := heart_draw_dat.GetDrawInfo(drawType)
	playerHeartDraw := getPlayerHeartDraw(state.Database, drawType)
	// 隔天重置抽奖次数
	if !coretime.IsInPointHour(player_dat.RESET_HEART_DRAW_TIMES_IN_HOUR, playerHeartDraw.DrawTime) {
		playerHeartDraw.DailyNum = 0
	}

	//非VIP用户也可以从VIP特权配置表读出次数
	maxDrawNum := int8(module.VIP.GetPrivilegeTimesByDB(state.Database, vip_dat.AIXINCHOUJIANG))

	// 今日抽奖次数已满
	fail.When(playerHeartDraw.DailyNum >= maxDrawNum, "[heartDraw] today is not draw")

	playerHeart := state.Database.Lookup.PlayerHeart(state.PlayerId)
	fail.When(playerHeart.Value < int16(heartDraw.CostHeart), "[heartDraw] player heart is not enough")

	var award *heart_draw_dat.HeartDrawAward
	var drawedIndex int
	chance := int8(0)
	randNum := int8(rand.Intn(100) + 1)

	for drawedIndex, award = range heartDraw.Awards {
		if randNum > chance && randNum <= chance+award.Chance {
			break
		}
		chance += award.Chance
	}

	// 默认第一次大转盘抽奖奖励元宝20，第一次抽奖，玩家的抽奖时间为0
	if module.CheckPlayerFirstOperate(state.Database, module.PLAYER_OPERATE_DRAW_WHEEL) {
		module.SetPlayerFirstOperated(state.Database, module.PLAYER_OPERATE_DRAW_WHEEL)
		for index, newAward := range heartDraw.Awards {
			// 一定要找到奖励20元宝的配置才进行奖励
			if newAward.AwardType == heart_draw_dat.DRAW_AWARD_TYPE_INGOT && newAward.AwardNum == 20 {
				award = newAward
				drawedIndex = index
				break
			}
		}
	}

	switch award.AwardType {
	case heart_draw_dat.DRAW_AWARD_TYPE_COIN:
		module.Player.IncMoney(state.Database, state.MoneyState, int64(award.AwardNum), player_dat.COINS, tlog.MFR_DRAW_HEART, xdEventType, "")

	case heart_draw_dat.DRAW_AWARD_TYPE_INGOT:
		module.Player.IncMoney(state.Database, state.MoneyState, int64(award.AwardNum), player_dat.INGOT, tlog.MFR_DRAW_HEART, xdEventType, "")

	case heart_draw_dat.DRAW_AWARD_TYPE_ITEM:
		module.Item.AddItem(state.Database, award.ItemId, award.AwardNum, tlog.IFR_DRAW_HEART, xdEventType, "")
	}

	// 累计抽奖次数
	playerHeartDraw.DailyNum += 1
	playerHeartDraw.DrawTime = coretime.GetNowTime()
	state.Database.Update.PlayerHeartDraw(playerHeartDraw)
	// 扣爱心
	module.Heart.DecHeart(state.Database, int16(heartDraw.CostHeart))

	module.Quest.RefreshDailyQuest(state, quest_dat.DAILY_QUEST_CLASS_DRAW_HEART_BOX)

	// 记录玩家抽奖奖品
	switch drawType {
	case heart_draw_dat.DRAW_TYPE_CARD:
		// 策划更改要求只显示玩家最新抽奖历史
		var playerDrawCardRecord *mdb.PlayerHeartDrawCardRecord
		state.Database.Select.PlayerHeartDrawCardRecord(func(row *mdb.PlayerHeartDrawCardRecordRow) {
			playerDrawCardRecord = row.GoObject()
			row.Break()
		})

		if playerDrawCardRecord == nil {
			state.Database.Insert.PlayerHeartDrawCardRecord(&mdb.PlayerHeartDrawCardRecord{
				Pid:       state.PlayerId,
				AwardType: award.AwardType,
				AwardNum:  award.AwardNum,
				ItemId:    award.ItemId,
				DrawTime:  playerHeartDraw.DrawTime,
			})
		} else {
			playerDrawCardRecord.AwardType = award.AwardType
			playerDrawCardRecord.AwardNum = award.AwardNum
			playerDrawCardRecord.ItemId = award.ItemId
			playerDrawCardRecord.DrawTime = playerHeartDraw.DrawTime
			state.Database.Update.PlayerHeartDrawCardRecord(playerDrawCardRecord)
		}

	case heart_draw_dat.DRAW_TYPE_WHEEL:
		// 只记录一条抽奖记录
		var playerDrawWheelRecord *mdb.PlayerHeartDrawWheelRecord
		state.Database.Select.PlayerHeartDrawWheelRecord(func(row *mdb.PlayerHeartDrawWheelRecordRow) {
			playerDrawWheelRecord = row.GoObject()
			row.Break()
		})

		if playerDrawWheelRecord == nil {
			state.Database.Insert.PlayerHeartDrawWheelRecord(&mdb.PlayerHeartDrawWheelRecord{
				Pid:       state.PlayerId,
				AwardType: award.AwardType,
				AwardNum:  award.AwardNum,
				ItemId:    award.ItemId,
				DrawTime:  playerHeartDraw.DrawTime,
			})
		} else {
			playerDrawWheelRecord.AwardType = award.AwardType
			playerDrawWheelRecord.AwardNum = award.AwardNum
			playerDrawWheelRecord.ItemId = award.ItemId
			playerDrawWheelRecord.DrawTime = playerHeartDraw.DrawTime
			state.Database.Update.PlayerHeartDrawWheelRecord(playerDrawWheelRecord)
		}
	}

	tlog.PlayerSystemModuleFlowLog(state.Database, tlog.SMT_HEART_DRAW)
	session.Send(&draw_api.HeartDraw_Out{
		Award: draw_api.AwardInfo{
			AwardType:  award.AwardType,
			AwardNum:   award.AwardNum,
			ItemId:     award.ItemId,
			DrawTime:   playerHeartDraw.DrawTime,
			AwardIndex: int16(drawedIndex),
		},
	})
}
