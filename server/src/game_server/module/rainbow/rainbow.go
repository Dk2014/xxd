package rainbow

import (
	"core/fail"
	"core/time"
	"game_server/api/protocol/rainbow_api"
	"game_server/battle"
	"game_server/dat/mission_dat"
	"game_server/dat/player_dat"
	"game_server/dat/quest_dat"
	"game_server/dat/rainbow_dat"
	"game_server/dat/role_dat"
	"game_server/dat/skill_dat"
	"game_server/module"
	"game_server/tlog"
	"game_server/xdlog"
	"math/rand"
)

func info(state *module.SessionState, out *rainbow_api.Info_Out) {
	fail.When(!module.Player.IsOpenFunc(state.Database, player_dat.FUNC_RAINBOW_LEVEL), "未开彩虹关卡")

	rainbowLevel := state.Database.Lookup.PlayerRainbowLevel(state.PlayerId)
	needUpdate := false
	if !time.IsInPointHour(player_dat.RESET_RAINBOW_LEVEL_TIMES_IN_HOUR, rainbowLevel.ResetTimestamp) {
		rainbowLevel.ResetNum = 0
		needUpdate = true
	}

	if !time.IsInPointHour(player_dat.RESET_RAINBOW_LEVEL_TIMES_IN_HOUR, rainbowLevel.AutoFightTime) {
		rainbowLevel.AutoFightNum = 0
		needUpdate = true
	}
	if !time.IsInPointHour(player_dat.RESET_RAINBOW_LEVEL_TIMES_IN_HOUR, rainbowLevel.BuyTimestamp) {
		rainbowLevel.BuyTimes = 0
		needUpdate = true
	}

	if needUpdate {
		state.Database.Update.PlayerRainbowLevel(rainbowLevel)
	}

	out.SegmentNum = rainbowLevel.Segment
	out.LevelOrder = rainbowLevel.Order
	out.Status = rainbowLevel.Status
	out.AutoFightNum = rainbowLevel.AutoFightNum
	out.BuyTimes = rainbowLevel.BuyTimes
	out.ResetNum = rainbow_dat.DAILY_RESET_NUM - int8(rainbowLevel.ResetNum)
	out.MaxSegmentCanJump = rainbowLevel.MaxOpenSegment
	out.MaxPassSegment = rainbowLevel.MaxPassSegment
}

func resetRainbowLevel(state *module.SessionState) {
	rainbowLevel := state.Database.Lookup.PlayerRainbowLevel(state.PlayerId)
	if !time.IsInPointHour(player_dat.RESET_RAINBOW_LEVEL_TIMES_IN_HOUR, rainbowLevel.ResetTimestamp) {
		rainbowLevel.ResetNum = 0
	}
	fail.When(rainbowLevel.ResetNum >= rainbow_dat.DAILY_RESET_NUM, "重置次数不足")
	rainbowLevel.ResetNum++
	rainbowLevel.Status = rainbow_dat.RAINBOW_LEVEL_STATUS_NEVER_PASS
	rainbowLevel.Order = rainbow_dat.INIT_LEVEL_ORDER
	rainbowLevel.ResetTimestamp = time.GetNowTime()
	state.Database.Update.PlayerRainbowLevel(rainbowLevel)

	state.RainbowLevelState = module.NewRainbowLevelState()
	state.SaveRainbowLevelState()

}

func awardInfo(state *module.SessionState, out *rainbow_api.AwardInfo_Out) {
	for _, order := range state.RainbowLevelState.AwardBoxIndex {
		out.Award = append(out.Award, rainbow_api.AwardInfo_Out_Award{
			Order: order,
		})
	}
}

func takeAward(state *module.SessionState, pos1, pos2 int8) (haveNextLevel bool) {
	rainbowLevel := state.Database.Lookup.PlayerRainbowLevel(state.PlayerId)

	fail.When(rainbowLevel.Status != rainbow_dat.RAINBOW_LEVEL_STATUS_NEVER_AWARD, "未通关不能领取奖励")

	levelId := rainbow_dat.GetRainbowLevelId(rainbowLevel.Segment, rainbowLevel.Order)
	awardConfigs := rainbow_dat.GetRainbowLevelAward(levelId)

	awardConis := state.RainbowLevelState.AwardCoin
	awardExp := state.RainbowLevelState.AwardExp
	awardRelationship := state.RainbowLevelState.AwardRelationship

	doAward(state, awardConis, awardExp, awardRelationship, []*rainbow_dat.RainbowLevelAward{awardConfigs[pos1-1], awardConfigs[pos2-1]})

	if rainbowLevel.Order == rainbow_dat.LEVEL_NUM_PER_SEGMENT {
		//领取本段最后一关的奖励后重置状态
		state.RainbowLevelState = module.NewRainbowLevelState()
	}

	//设置下一个关卡
	if rainbowLevel.Segment == rainbow_dat.MaxRainbowLevelSegment && rainbowLevel.Order == rainbow_dat.LEVEL_NUM_PER_SEGMENT { //已打通最后一个关卡？
		haveNextLevel = false
		rainbowLevel.Status = rainbow_dat.RAINBOW_LEVEL_STATUS_NO_MORE
	} else {
		haveNextLevel = true
		rainbowLevel.Status = rainbow_dat.RAINBOW_LEVEL_STATUS_NEVER_PASS
		if rainbowLevel.Order < rainbow_dat.LEVEL_NUM_PER_SEGMENT { //当前段内还有关卡？
			rainbowLevel.Order += 1
		} else {
			rainbowLevel.Segment += 1
			rainbowLevel.Order = 1
		}
	}

	state.Database.Update.PlayerRainbowLevel(rainbowLevel)

	//保存状态
	state.SaveRainbowLevelState()

	return haveNextLevel
}

func doAward(state *module.SessionState, awardConis int64, awardExp int32, awardRelationship int32, awardList []*rainbow_dat.RainbowLevelAward) {
	var multExp, multCoins bool

	for _, award := range awardList {
		switch award.AwardType {
		case rainbow_dat.AWARD_TYPE_COIN:
			awardConis += int64(award.AwardNum)
		case rainbow_dat.AWARD_TYPE_ITEM:
			module.Item.AddItem(state.Database, int16(award.ItemId), int16(award.AwardNum), tlog.IFR_RAINBOW, xdlog.ET_RAINBOW, "")
		case rainbow_dat.AWARD_TYPE_EQUIMENT:
			module.Item.AddItem(state.Database, int16(award.ItemId), int16(award.AwardNum), tlog.IFR_RAINBOW, xdlog.ET_RAINBOW, "")
		case rainbow_dat.AWARD_TYPE_EXP:
			awardExp += int32(award.AwardNum)
		case rainbow_dat.AWARD_TYPE_MULTIPLE_EXP:
			multExp = true
		case rainbow_dat.AWARD_TYPE_MULTIPLE_COIN:
			multCoins = true
		case rainbow_dat.AWARD_TYPE_RECOVER_BUDDY_SKILL:
			for skillId, releaseNum := range state.RainbowLevelState.SkillReleaseNum {
				if skillId > skill_dat.SKILL_IS_NULL {
					skillContent := skill_dat.GetSkillContent(int16(skillId))
					if releaseNum < int(skillContent.ReleaseNum) {
						//技能可用次数少于技能默认的使用次数？则恢复
						state.RainbowLevelState.SkillReleaseNum[skillId] = int(skillContent.ReleaseNum)
					}
				}
			}
		case rainbow_dat.AWARD_TYPE_RECOVER_GHOST:
			state.RainbowLevelState.UsedGhost = make(map[int16]int8)
			for _, roleId := range module.Role.GetFormRoleId(state) {
				fighterAttr := state.RainbowLevelState.AttackerInfo.Fighters[int(roleId)]
				fighterAttr.UsedGhostSkill = false
				state.RainbowLevelState.AttackerInfo.Fighters[int(roleId)] = fighterAttr
			}
		case rainbow_dat.AWARD_TYPE_RECOVER_PET:
			state.RainbowLevelState.CalledBattlePet = make(map[int32]int8)
		case rainbow_dat.AWARD_TYPE_ADD_POWER:
			mailRole := module.Role.GetMainRole(state.Database)

			//主角增加精气
			fighterAttr := state.RainbowLevelState.AttackerInfo.Fighters[int(state.RoleId)]
			fighterAttr.Power += int(award.AwardNum)

			roleInfo := role_dat.GetRoleLevelInfo(state.RoleId, mailRole.Level)
			if fighterAttr.Power > int(roleInfo.MaxPower) {
				fighterAttr.Power = int(roleInfo.MaxPower)
			}
			state.RainbowLevelState.AttackerInfo.Fighters[int(state.RoleId)] = fighterAttr
		case rainbow_dat.AWARD_TYPE_ADD_HEALTH_BY_PERSENT:
			for _, roleId := range module.Role.GetFormRoleId(state) {
				fighterAttr := state.RainbowLevelState.AttackerInfo.Fighters[int(roleId)]
				if fighterAttr.Health < 0 {
					fighterAttr.Health = 0
				}
				fighterAttr.Health += int(float64(fighterAttr.MaxHealth) * float64(award.AwardNum) / 100.0)
				if fighterAttr.Health > fighterAttr.MaxHealth {
					fighterAttr.Health = fighterAttr.MaxHealth
				}
				state.RainbowLevelState.AttackerInfo.Fighters[int(roleId)] = fighterAttr
			}
		case rainbow_dat.AWARD_TYPE_AAD_GHOST_POWER:
			ghostInfo := state.RainbowLevelState.AttackerInfo.GhostInfo[state.PlayerId]
			ghostInfo.GhostPower += int(award.AwardNum)
			if ghostInfo.GhostPower > battle.FULL_GHOST_POWER {
				ghostInfo.GhostPower = battle.FULL_GHOST_POWER
			}
			state.RainbowLevelState.AttackerInfo.GhostInfo[state.PlayerId] = ghostInfo
		default:
			fail.When(true, "彩虹关卡：未知奖励类型")
		}
	}

	if awardConis > 0 {
		if multCoins {
			awardConis *= 2
		}
		module.Player.IncMoney(state.Database, state.MoneyState, awardConis, player_dat.COINS, tlog.MFR_RAINBOW, xdlog.ET_RAINBOW, "")
	}

	if awardExp > 0 {
		if multExp {
			awardExp *= 2
		}
		module.Role.AddFormRoleExp(state, int64(awardExp), tlog.EFT_RAINBOW)
	}

	if awardRelationship > 0 {
		module.Team.IncRelationship(state.Database, awardRelationship)
	}

}

func jumpToSegment(state *module.SessionState, segment int16) {
	rainbowLevel := state.Database.Lookup.PlayerRainbowLevel(state.PlayerId)

	fail.When(segment < rainbow_dat.MIN_SEGMENT_NUM_CAN_JUMP || segment > rainbow_dat.MAX_RAINBOW_SEGMENT_NUM, "彩虹关卡：跳到非法段")
	fail.When(rainbowLevel.Segment >= segment, "只能跳到在当前段后的彩虹段")
	fail.When(rainbowLevel.MaxOpenSegment < segment, "不能跳到尚未开启的彩虹关卡")

	rainbowLevel.Segment = segment
	rainbowLevel.Status = rainbow_dat.RAINBOW_LEVEL_STATUS_NEVER_PASS
	rainbowLevel.Order = rainbow_dat.INIT_LEVEL_ORDER
	state.Database.Update.PlayerRainbowLevel(rainbowLevel)

	state.RainbowLevelState = module.NewRainbowLevelState()
	state.SaveRainbowLevelState()
}

func rainbowLevelAutoFight(state *module.SessionState, segment int16) *rainbow_api.AutoFight_Out {
	fail.When(segment < 1 || segment > rainbow_dat.MAX_RAINBOW_SEGMENT_NUM, "彩虹关卡：扫荡非法的段")

	rainbowLevel := state.Database.Lookup.PlayerRainbowLevel(state.PlayerId)
	fail.When(rainbowLevel.AutoFightNum >= rainbow_dat.DAILY_AUTO_FIGHT_NUM+int8(rainbowLevel.BuyTimes), "彩虹关卡扫荡次数已满")
	//fail.When(segment > rainbowLevel.Segment, "彩虹关卡扫荡段数不正确")
	fail.When(segment > rainbowLevel.MaxPassSegment, "彩虹关卡扫荡段未通关")

	rainbowLevel.AutoFightNum++
	rainbowLevel.AutoFightTime = time.GetNowTime()
	state.Database.Update.PlayerRainbowLevel(rainbowLevel)

	//关卡的铜钱以及经验奖励扫荡后立即给玩家，宝箱则需要玩家点开的时候发放
	var awardCoin int64
	var awardExp int32
	for order := int8(rainbow_dat.INIT_LEVEL_ORDER); order <= rainbow_dat.LEVEL_NUM_PER_SEGMENT; order++ {
		levelId := rainbow_dat.GetRainbowLevelId(segment, order)
		levelInfo := mission_dat.GetMissionLevelById(levelId)
		awardCoin += levelInfo.AwardCoin
		awardExp += levelInfo.AwardExp
	}

	awardBox := rainbow_dat.GetRainbowAutoFightBox(int32(segment))
	awardList := []*rainbow_dat.RainbowLevelAward{}
	var randChance, awardedChance, awardIdx, maxChance int
	for _, award := range awardBox {
		maxChance += int(award.AwardChance)
	}

	awardIdx = -1
	for i := 1; i <= 2; /*rainbow_dat.AWARD_NUM 下面的算法只能随即获得两个奖励*/ i++ {
		for j, award := range awardBox {
			randChance = rand.Intn(maxChance-awardedChance) + 1
			if j != awardIdx {
				awardedChance += int(award.AwardChance)
			}
			if j != awardIdx && randChance <= int(award.AwardChance) {
				awardIdx = j
				awardList = append(awardList, award)
				awardedChance = int(award.AwardChance)
				break
			}
		}
	}

	doAward(state, awardCoin, awardExp, 0, awardList)

	// 彩虹管卡扫荡增加刷新任务
	module.Quest.RefreshDailyQuest(state, quest_dat.DAILY_QUEST_CLASS_LIMIT_LEVEL)

	autofightOut := &rainbow_api.AutoFight_Out{
		AwardCoin:    int32(awardCoin),
		AwardExp:     int32(awardExp),
		AwardBoxPos1: awardList[0].Order,
		AwardBoxPos2: awardList[1].Order,
	}

	return autofightOut
}
