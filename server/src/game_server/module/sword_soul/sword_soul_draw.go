package sword_soul

import (
	"core/fail"
	"core/net"
	util "core/time"
	"fmt"
	"game_server/api/protocol/sword_soul_api"
	"game_server/dat/channel_dat"
	"game_server/dat/mail_dat"
	"game_server/dat/player_dat"
	"game_server/dat/quest_dat"
	"game_server/dat/sword_soul_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/module/rpc"
	"game_server/tlog"
	"game_server/xdlog"
	"math/rand"
	"time"
)

func Draw(session *net.Session, in *sword_soul_api.Draw_In, out *sword_soul_api.Draw_Out) {
	state := module.State(session)
	db := state.Database
	drawBox := int8(in.Box)
	drawType := in.DrawType

	tlog.PlayerSystemModuleFlowLog(db, tlog.SMT_SWORD_SOUL)

	var (
		drawOutType       int
		playerSwordSoulId int64
		swordSoulId       int16
		IngotYelloOne     int16
	)
	switch drawType {
	case sword_soul_api.DRAW_TYPE_COIN:
		// 铜钱拔剑

		drawOutType, playerSwordSoulId, swordSoulId, IngotYelloOne = drawByCoin(session, drawBox, tlog.IFR_DRAW_SWORD, xdlog.ET_DRAW_SWORD_SOUL)

		switch drawOutType {
		case DRAW_OUT_NULL:
			// 没有拔到剑心
			out.Coins = sword_soul_dat.RUBBISH_AWARD_COIN
			//返还铜币
			module.Player.IncMoney(db, state.MoneyState, sword_soul_dat.RUBBISH_AWARD_COIN, player_dat.COINS, tlog.MFR_DRAW_SWORD, xdlog.ET_DRAW_SWORD_SOUL, "")
			tlog.PlayerSwordDrawFlowLog(db, tlog.MT_COIN, int32(sword_soul_dat.BoxConfigs[drawBox].Coin), 0)

		case DRAW_OUT_SWORD_SOUL:
			// 拔到剑心
			out.Id = playerSwordSoulId
			out.SwordSoulId = swordSoulId
			tlog.PlayerSwordDrawFlowLog(db, tlog.MT_COIN, int32(sword_soul_dat.BoxConfigs[drawBox].Coin), int32(swordSoulId))

		case DRAW_OUT_QIANLONG:

			// 增加潜龙剑心
			if IngotYelloOne == 1 && drawBox == sword_soul_dat.BOX_E {
				newPlayerSwordSoulId := addSwordSoul(state, sword_soul_dat.QIAN_LONG_ID, tlog.IFR_DRAW_SWORD)
				out.Id = newPlayerSwordSoulId
				out.SwordSoulId = sword_soul_dat.QIAN_LONG_ID
			}
			out.Fragments = 1
			module.Player.IncSwordSoulFragment(db, 1, player_dat.SWORDSOULFRAGMENT, tlog.MFR_DRAW_SWORD, xdlog.ET_DRAW_SWORD_SOUL)
			tlog.PlayerSwordDrawFlowLog(db, tlog.MT_COIN, int32(sword_soul_dat.BoxConfigs[drawBox].Coin), sword_soul_dat.QIAN_LONG_ID)
		}

	case sword_soul_api.DRAW_TYPE_INGOT:
		// 元宝拔剑
		drawOutType, playerSwordSoulId, swordSoulId = drawByIngot(session, drawBox, tlog.IFR_DRAW_SWORD, xdlog.ET_DRAW_SWORD_SOUL)
		var swordId int32

		switch drawOutType {
		case DRAW_OUT_NULL:
			// 没有拔到剑心
			out.Coins = sword_soul_dat.RUBBISH_AWARD_COIN
			//返还铜币
			module.Player.IncMoney(db, state.MoneyState, sword_soul_dat.RUBBISH_AWARD_COIN, player_dat.COINS, tlog.MFR_DRAW_SWORD, xdlog.ET_DRAW_SWORD_SOUL, "")
			swordId = 0

		case DRAW_OUT_SWORD_SOUL:
			// 拔到剑心
			out.Id = playerSwordSoulId
			out.SwordSoulId = swordSoulId
			swordId = int32(swordSoulId)

		case DRAW_OUT_QIANLONG:

			// 增加潜龙剑心
			out.Fragments = 1
			newPlayerSwordSoulId := addSwordSoul(state, sword_soul_dat.QIAN_LONG_ID, tlog.IFR_DRAW_SWORD)
			out.Id = newPlayerSwordSoulId
			out.SwordSoulId = sword_soul_dat.QIAN_LONG_ID
			swordId = sword_soul_dat.QIAN_LONG_ID
			module.Player.IncSwordSoulFragment(db, 1, player_dat.SWORDSOULFRAGMENT, tlog.MFR_DRAW_SWORD, xdlog.ET_DRAW_SWORD_SOUL)
		}

		playerSwordSoulState := db.Lookup.PlayerSwordSoulState(state.PlayerId)
		tlog.PlayerSwordDrawFlowLog(db, tlog.MT_INGOT, int32(sword_soul_dat.GetSwordDrawPriceInIngot(playerSwordSoulState.IngotNum)), swordId)
	}

	if drawOutType == DRAW_OUT_SWORD_SOUL {
		swordSoulDat := sword_soul_dat.GetSwordSoul(swordSoulId)
		if swordSoulDat.Quality == sword_soul_dat.QUALITY_ARTIFACT || swordSoulDat.Quality == sword_soul_dat.QUALITY_ONLY {
			//module.Chan.AddWorldTplMessage(state.PlayerId, state.PlayerNick,
			rpc.RemoteAddWorldChannelTplMessage(state.PlayerId, []channel_dat.MessageTpl{
				channel_dat.MessageDrawSwordSoul{
					Player: channel_dat.ParamPlayer{state.PlayerNick, state.PlayerId},
					Item:   channel_dat.ParamItem{mail_dat.ATTACHMENT_SWORD_SOUL, swordSoulDat.Id},
				},
			})
		}
	}

	playerSwordSoulState := db.Lookup.PlayerSwordSoulState(state.PlayerId)
	out.BoxState = playerSwordSoulState.BoxState

	module.Quest.RefreshDailyQuest(state, quest_dat.DAILY_QUEST_CLASS_DRAW_SWORD)
}

func drawByCoin(session *net.Session, drawBox int8, itemReason, xdEventType int32) (drawOutType int, playerSwordSoulId int64, swordSoulId int16, IngotYelloOne int16) {
	state := module.State(session)
	db := state.Database

	playerSwordSoulState := db.Lookup.PlayerSwordSoulState(state.PlayerId)

	IngotYelloOne = playerSwordSoulState.IngotYelloOne
	//如果当前拔剑次数已满的，那么就设置级定时器定时恢复
	//shoulSetTimmer := playerSwordSoulState.Num == sword_soul_dat.MAX_DRAW_NUM

	fail.When(!boxIsAvailable(playerSwordSoulState.BoxState, drawBox), "box is not available")

	//检查拔剑次数
	//fail.When(playerSwordSoulState.Num <= 0, "no draw num")
	// 次数减一
	//playerSwordSoulState.Num--

	// 消耗铜钱
	module.Player.DecMoney(state.Database, state.MoneyState, sword_soul_dat.BoxConfigs[drawBox].Coin, player_dat.COINS, tlog.MFR_DRAW_SWORD, xdEventType)

	// 随机开启下一个箱子
	nextBoxProbability := rand.Float64()
	if nextBoxProbability < sword_soul_dat.NextBoxProbability[drawBox] {
		playerSwordSoulState.BoxState = playerSwordSoulState.BoxState | sword_soul_dat.BoxConfigs[drawBox+1].Key
	}
	if drawBox > 0 {
		playerSwordSoulState.BoxState = playerSwordSoulState.BoxState ^ sword_soul_dat.BoxConfigs[drawBox].Key
	}

	// 如果第一次拔剑 必的一个剑心 必开第二箱子
	if module.CheckPlayerFirstOperate(state.Database, module.PLAYER_OPERATE_SWORD_DRAW_1) {
		playerSwordSoulState.BoxState = 1 | 2
		db.Update.PlayerSwordSoulState(playerSwordSoulState)
		//module.Notify.SendSwordSoulDrawNumChange(session, playerSwordSoulState.Num, playerSwordSoulState.UpdateTime+sword_soul_dat.RECOVERY_TIME)
		playerSwordSoulId = addSwordSoul(state, FIRST_DRAW_AWARD_SWORD_SOUL, itemReason)
		drawOutType = DRAW_OUT_SWORD_SOUL
		module.SetPlayerFirstOperated(state.Database, module.PLAYER_OPERATE_SWORD_DRAW_1)
		swordSoulId = FIRST_DRAW_AWARD_SWORD_SOUL
		return drawOutType, playerSwordSoulId, swordSoulId, IngotYelloOne
	}

	//原本拔剑次满，消耗了一次
	//if shoulSetTimmer {
	//	playerSwordSoulState.UpdateTime = util.GetNowTime()
	//	startTimer(state, playerSwordSoulState, sword_soul_dat.RECOVERY_TIME)
	//}

	//如果第2次拔剑 必的一个剑心
	if module.CheckPlayerFirstOperate(state.Database, module.PLAYER_OPERATE_SWORD_DRAW_2) {
		db.Update.PlayerSwordSoulState(playerSwordSoulState)
		playerSwordSoulId = addSwordSoul(state, FIRST_DRAW_AWARD_SWORD_SOUL, itemReason)
		drawOutType = DRAW_OUT_SWORD_SOUL
		module.SetPlayerFirstOperated(state.Database, module.PLAYER_OPERATE_SWORD_DRAW_2)
		swordSoulId = FIRST_DRAW_AWARD_SWORD_SOUL
		return drawOutType, playerSwordSoulId, swordSoulId, IngotYelloOne
	}

	// 拔剑
	var swordSoulQuality int8
	if drawBox == sword_soul_dat.BOX_E && playerSwordSoulState.IngotYelloOne > 0 {
		// RMB player tricking time
		trick_prob := float64(playerSwordSoulState.SupersoulAdditionalPossible) / 100
		// 每次增加1%，10次必爆
		if trick_prob >= float64(10)/100 {
			// ten times, bad lucky, poor guy
			swordSoulQuality = sword_soul_dat.QUALITY_ARTIFACT
		} else {
			drawProbability := rand.Float64()
			// replace the possibility of the lowest quality sword soul
			if drawProbability < trick_prob {
				swordSoulQuality = sword_soul_dat.QUALITY_ARTIFACT
			} else {
				// 准备拔剑概率模板, VIP Box E
				SwordSoulProbabilityWithQuality := sword_soul_dat.SwordSoulProbability_VIP[1]
				// now normally iterate, first from the lowest quality one
				for quality, probability := range SwordSoulProbabilityWithQuality {
					if drawProbability < probability {
						swordSoulQuality = int8(quality)
						break
					}
					drawProbability -= probability
				}
			}
		}
		// this RMB got his lucky back, so we clean the count.
		if swordSoulQuality == sword_soul_dat.QUALITY_ARTIFACT {
			playerSwordSoulState.SupersoulAdditionalPossible = 0
		} else {
			playerSwordSoulState.SupersoulAdditionalPossible += 1
		}
		playerSwordSoulState.IngotYelloOne = 0
	} else {
		// 穷人拔剑流程

		// 准备拔剑概率模板
		SwordSoulProbabilityWithQuality := sword_soul_dat.SwordSoulProbability[drawBox]

		drawProbability := rand.Float64()
		for quality, probability := range SwordSoulProbabilityWithQuality {
			if drawProbability < probability {
				swordSoulQuality = int8(quality)
				break
			}
			drawProbability -= probability
		}
	}
	// 更新用户剑心状态
	db.Update.PlayerSwordSoulState(playerSwordSoulState)
	//module.Notify.SendSwordSoulDrawNumChange(session, playerSwordSoulState.Num, playerSwordSoulState.UpdateTime+sword_soul_dat.RECOVERY_TIME)

	if swordSoulQuality == sword_soul_dat.QUALITY_NULL {
		// 拔到垃圾
		drawOutType = DRAW_OUT_NULL
		return drawOutType, playerSwordSoulId, swordSoulId, IngotYelloOne
	} else if swordSoulQuality == sword_soul_dat.QUALITY_SPECIAL {
		// 拔到潜龙
		drawOutType = DRAW_OUT_QIANLONG
		return drawOutType, playerSwordSoulId, swordSoulId, IngotYelloOne
	} else {
		// 拔到剑心
		drawOutType = DRAW_OUT_SWORD_SOUL
		swordSoulId = randomSwordSoulIdByQuality(swordSoulQuality)
		playerSwordSoulId = addSwordSoul(state, swordSoulId, itemReason)
		return drawOutType, playerSwordSoulId, swordSoulId, IngotYelloOne
	}
}

func drawByIngot(session *net.Session, drawBox int8, itemReason, xdEventType int32) (drawOutType int, playerSwordSoulId int64, swordSoulId int16) {
	state := module.State(session)
	db := state.Database

	fail.When(drawBox != sword_soul_dat.BOX_D, "wrong draw box")

	playerSwordSoulState := db.Lookup.PlayerSwordSoulState(state.PlayerId)

	fmt.Sprintf("%d", playerSwordSoulState.IngotNum)
	fmt.Sprintf("%d", sword_soul_dat.MAX_INGOT_DRAW_NUM)
	fmt.Sprintf("sword soul draw by ingot %s times limited", sword_soul_dat.MAX_INGOT_DRAW_NUM)

	fail.When(playerSwordSoulState.IngotNum >= int64(sword_soul_dat.MAX_INGOT_DRAW_NUM), "sword soul draw by ingot times limited")
	//如果当前拔剑次数已满的，那么就设置级定时器定时恢复
	//shoulSetTimmer := playerSwordSoulState.Num == sword_soul_dat.MAX_DRAW_NUM

	fixSwordSoulIngotNum(playerSwordSoulState)
	increaseSwordSoulIngotNum(playerSwordSoulState)

	//检查拔剑次数
	//fail.When(playerSwordSoulState.Num <= 0, "no draw num")
	// 次数减一
	//playerSwordSoulState.Num--

	// 消耗元宝
	module.Player.DecMoney(state.Database, state.MoneyState, sword_soul_dat.GetSwordDrawPriceInIngot(playerSwordSoulState.IngotNum), player_dat.INGOT, tlog.MFR_DRAW_SWORD, xdEventType)
	//消耗铜钱
	module.Player.DecMoney(state.Database, state.MoneyState, sword_soul_dat.GetSwordDrawPriceInIngotCoin(playerSwordSoulState.IngotNum), player_dat.COINS, tlog.MFR_DRAW_SWORD, xdEventType)

	// 随机开启下一个箱子
	nextBoxProbability := rand.Float64()
	if nextBoxProbability < sword_soul_dat.NextBoxProbability_VIP[0] {
		playerSwordSoulState.BoxState = playerSwordSoulState.BoxState | sword_soul_dat.BoxConfigs[drawBox+1].Key
		playerSwordSoulState.IngotYelloOne = 1
	}

	// 准备拔剑概率模板
	SwordSoulProbabilityWithQuality := sword_soul_dat.SwordSoulProbability_VIP[0]

	//原本拔剑次满，消耗了一次
	//if shoulSetTimmer {
	//	playerSwordSoulState.UpdateTime = util.GetNowTime()
	//	startTimer(state, playerSwordSoulState, sword_soul_dat.RECOVERY_TIME)
	//}

	// 更新用户剑心状态
	db.Update.PlayerSwordSoulState(playerSwordSoulState)
	//module.Notify.SendSwordSoulDrawNumChange(session, playerSwordSoulState.Num, playerSwordSoulState.UpdateTime+sword_soul_dat.RECOVERY_TIME)

	// 拔剑
	drawProbability := rand.Float64()
	var swordSoulQuality int8
	for quality, probability := range SwordSoulProbabilityWithQuality {
		if drawProbability < probability {
			swordSoulQuality = int8(quality)
			break
		}
		drawProbability -= probability
	}

	if swordSoulQuality == sword_soul_dat.QUALITY_NULL {
		drawOutType = DRAW_OUT_NULL
		return drawOutType, playerSwordSoulId, swordSoulId
	} else if swordSoulQuality == sword_soul_dat.QUALITY_SPECIAL {
		drawOutType = DRAW_OUT_QIANLONG
		return drawOutType, playerSwordSoulId, swordSoulId
	} else {
		drawOutType = DRAW_OUT_SWORD_SOUL
		swordSoulId = randomSwordSoulIdByQuality(swordSoulQuality)
		playerSwordSoulId = addSwordSoul(state, swordSoulId, itemReason)
		return drawOutType, playerSwordSoulId, swordSoulId
	}
}

func boxIsAvailable(boxState int8, box int8) bool {
	return (boxState & sword_soul_dat.BoxConfigs[box].Key) == sword_soul_dat.BoxConfigs[box].Key
}

func randomSwordSoulIdByQuality(quality int8) (swordSoulId int16) {
	swordSoulIds := sword_soul_dat.GetDrawSwordSoulsByQuality(quality)
	key := rand.Intn(len(swordSoulIds))
	swordSoulId = swordSoulIds[key]
	return swordSoulId
}

// 更新元宝拔剑次数
func fixSwordSoulIngotNum(playerSwordSoulState *mdb.PlayerSwordSoulState) (needUpdate bool) {
	needUpdate = false
	if playerSwordSoulState == nil {
		return needUpdate
	}
	if playerSwordSoulState.LastIngotDrawTime < util.GetTodayZero() {
		// 上次元宝拔剑不在今天
		playerSwordSoulState.IngotNum = 0
		needUpdate = true
	}
	return needUpdate
}

// 累加元宝拔剑次数
func increaseSwordSoulIngotNum(playerSwordSoulState *mdb.PlayerSwordSoulState) {
	if playerSwordSoulState == nil {
		return
	}
	now := util.GetNowTime()
	playerSwordSoulState.LastIngotDrawTime = now
	playerSwordSoulState.IngotNum++
	return
}

//  更新拔剑次数
func updateSwordSoulNum(state *module.SessionState, playerSwordSoulState *mdb.PlayerSwordSoulState) {
	if playerSwordSoulState == nil || playerSwordSoulState.Num > sword_soul_dat.MAX_DRAW_NUM {
		return
	}
	now := util.GetNowTime()
	passTime := now - playerSwordSoulState.UpdateTime
	deltaT := passTime % sword_soul_dat.RECOVERY_TIME

	playerSwordSoulState.Num += int16(passTime / sword_soul_dat.RECOVERY_TIME * sword_soul_dat.RECOVERY_VALUE)
	playerSwordSoulState.UpdateTime = int64(now - deltaT)

	if playerSwordSoulState.Num > sword_soul_dat.MAX_DRAW_NUM {
		playerSwordSoulState.Num = sword_soul_dat.MAX_DRAW_NUM
	}
	state.Database.Update.PlayerSwordSoulState(playerSwordSoulState)
	startTimer(state, playerSwordSoulState, time.Duration(sword_soul_dat.RECOVERY_TIME-deltaT))
}

func increase(state *module.SessionState, playerSwordSoulState *mdb.PlayerSwordSoulState, incVal int16) {
	playerSwordSoulState.Num += incVal
	playerSwordSoulState.UpdateTime = util.GetNowTime()
	cdTime := playerSwordSoulState.UpdateTime + sword_soul_dat.RECOVERY_TIME

	// 最终拔剑次数恢复满了则停止定时器
	if playerSwordSoulState.Num >= sword_soul_dat.MAX_DRAW_NUM {
		stopTimer(state)
		cdTime = 0
	}

	if session, ok := module.Player.GetPlayerOnline(state.PlayerId); ok {
		module.Notify.SendSwordSoulDrawNumChange(session, playerSwordSoulState.Num, cdTime)
	}
}

func increaseByTimer(state *module.SessionState) {
	playerSwordSoulState := state.Database.Lookup.PlayerSwordSoulState(state.PlayerId)

	//恢复拔剑次数
	increase(state, playerSwordSoulState, sword_soul_dat.RECOVERY_VALUE)

	if playerSwordSoulState.Num > sword_soul_dat.MAX_DRAW_NUM {
		playerSwordSoulState.Num = sword_soul_dat.MAX_DRAW_NUM
	}

	state.Database.Update.PlayerSwordSoulState(playerSwordSoulState)

	startTimer(state, playerSwordSoulState, sword_soul_dat.RECOVERY_TIME)
}

// 在线玩家定时恢复拔剑次数
func startTimer(state *module.SessionState, playerSwordSoulState *mdb.PlayerSwordSoulState, delaySecs time.Duration) {
	// 当前可用拔剑次数
	if playerSwordSoulState.Num >= sword_soul_dat.MAX_DRAW_NUM {
		return
	}

	state.TimerMgr.Start(module.TIMER_SWORD_SOUL, delaySecs*time.Second, increaseByTimer)
}

func stopTimer(state *module.SessionState) {
	state.TimerMgr.Stop(module.TIMER_SWORD_SOUL)
}
