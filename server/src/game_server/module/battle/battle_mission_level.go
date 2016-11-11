package battle

/*
	城镇区域关卡战场

	玩法要求：

	战斗失败
		* 直接回城镇，不允许留在关卡内

	战斗胜利或逃跑后继续留在关卡内：
		* 保存玩家上阵角色的基本属性值
		* 保存玩家在战场中未消失的buf状态

	退出关卡后，保存的玩家状态失效
*/

import (
	"core/fail"
	"core/net"
	"game_server/api/protocol/battle_api"
	"game_server/battle"
	"game_server/dat/battle_pet_dat"
	"game_server/dat/mission_dat"
	"game_server/dat/player_dat"
	"game_server/dat/quest_dat"
	"game_server/dat/role_dat"
	"game_server/dat/skill_dat"
	"game_server/module"
	"game_server/module/rpc"
	"game_server/tlog"
	"game_server/xdlog"
	"math"
	"math/rand"
)

type Battle_MissionLevel struct {
	module.BasicBattle
	battleState *battle.BattleState
	notAward    bool // 对战斗胜利和失败不作处理

	bestRound    int8 // 通关的最好回合数
	battleStatus int  // 战场状态
	isBoss       bool // 当前怪物组是否是boss
	boatId       int64
}

func (this *Battle_MissionLevel) GetBattle() *battle.BattleState {
	return this.battleState
}

func (this *Battle_MissionLevel) handlerBattleResult(session *net.Session, results []*battle.FightResult, status, nowRound int) {
	var (
		state       = module.State(session)
		battleState = this.battleState
	)

	// 检查是否有捕获灵宠, -1表示没有捕捉灵宠
	catchPetId := battleState.GetCatchBattlePetId()
	if catchPetId != -1 {
		state.MissionLevelState.CatchedBattlePetId = catchPetId
		module.Mission.NotifyCatchBattlePet(session, catchPetId)
	}
	if state.MissionLevelState.MaxRound > 0 && nowRound+1 >= state.MissionLevelState.MaxRound && status != battle.ATK_WIN {
		status = battle.DEF_WIN
	}
	this.battleStatus = status
	rsp := GetNextRoundResponse(results, status, nowRound, battleState, nil, state.MissionLevelState.SkillReleaseNum)
	session.Send(rsp)
	if (status == battle.ATK_WIN || status == battle.DEF_WIN) && this.notAward {
		//this.LeaveBattle(session)
		return
	}

	switch status {
	case battle.ATK_WIN:
		this.DoWin(session)
	case battle.DEF_WIN:
		this.DoLose(session)
	}
}

func (this *Battle_MissionLevel) InitRound(session *net.Session) {
	var (
		state        = module.State(session)
		nextPid      = this.getNextPlayer(this.battleState.GetRounds())
		oldRound int = this.battleState.GetRounds()
	)
	if this.battleState.BatType == battle.BT_PVE_LEVEL {
		//动态增加怪物组
		fighterGroup := module.PetVirtualEnv.WarnupNextEnemyGroup(state, this.battleState)
		if fighterGroup != nil {
			this.battleState.RunTimeAddFightersGroup(battle.FT_DEF, fighterGroup)
			newGroupRsp := GetNewFighterGroupResponse(0, battle.FT_DEF, fighterGroup, nil)
			session.Send(newGroupRsp)
		}
	}

	if this.battleState.BatType == battle.BT_DRIVING_LEVEL {
		//仙山拜访动态增加怪物组
		fighterGroup := module.DrivingSword.WarnupNextEnemyGroup(state, this.battleState)
		if fighterGroup != nil {
			this.battleState.RunTimeAddFightersGroup(battle.FT_DEF, fighterGroup)
			newGroupRsp := GetNewFighterGroupResponse(0, battle.FT_DEF, fighterGroup, nil)
			session.Send(newGroupRsp)
		}
	}
	results, status, nowRound, _ := this.battleState.RequireTotem()
	if len(results) > 0 {
		this.handlerBattleResult(session, results, status, nowRound)
	}
	for nextPid > 0 && nextPid != state.PlayerId {
		results, status, nowRound = this.battleState.AutoUseGhost(nextPid)
		if len(results) > 0 {
			this.handlerBattleResult(session, results, status, nowRound)
			nextPid = this.getNextPlayer(oldRound)
			continue
		}

		results, status, nowRound = this.battleState.PrepareReady(nextPid)
		if len(results) > 0 {
			this.handlerBattleResult(session, results, status, nowRound)
		}
		nextPid = this.getNextPlayer(oldRound)
	}
	if nextPid == state.PlayerId {
		session.Send(&battle_api.NotifyReady_Out{
			Pid: nextPid,
		})
	}
}

func (this *Battle_MissionLevel) UseGhost(session *net.Session, isAttacker bool, posIdx int8) {
	state := module.State(session)
	results, status, nowRound := this.battleState.UseGhost(state.PlayerId, isAttacker, posIdx)
	if len(results) > 0 {
		this.handlerBattleResult(session, results, status, nowRound)
	}
}

func (this *Battle_MissionLevel) UseItem(session *net.Session, isAttacker bool, posIdx int8, itemId int16) {
	state := module.State(session)
	results, status, nowRound := this.battleState.UseItem(state.PlayerId, isAttacker, posIdx, int32(itemId))
	if len(results) > 0 {
		this.handlerBattleResult(session, results, status, nowRound)
	}
}

func (this *Battle_MissionLevel) PrepareReady(session *net.Session, isAuto bool) {
	var (
		state    = module.State(session)
		results  []*battle.FightResult
		status   int
		nowRound int = this.battleState.GetRounds()
	)
	oldRound := nowRound
	nextPid := this.battleState.GetNextPlayer(nowRound)
	for nextPid > 0 {
		//for oldRound == nowRound && nextPid > 0 {
		if nextPid != state.PlayerId || isAuto {
			results, status, nowRound = this.battleState.AutoUseGhost(nextPid)
			if len(results) > 0 {
				this.handlerBattleResult(session, results, status, nowRound)
				nextPid = this.getNextPlayer(oldRound)
				continue
			}
		}
		results, status, nowRound = this.battleState.PrepareReady(nextPid)
		if len(results) > 0 {
			this.handlerBattleResult(session, results, status, nowRound)
		}

		//results, status, nowRound = this.battleState.AutoUseGhost(nextPid)
		nextPid = this.getNextPlayer(oldRound)
	}
}

func (this *Battle_MissionLevel) SetAuto(session *net.Session) {
	state := module.State(session)
	this.battleState.SetAuto_v2(state.PlayerId)
}

func (this *Battle_MissionLevel) CancelAuto(session *net.Session) {
	state := module.State(session)
	joinRound := this.battleState.CancelAuto_v2(state.PlayerId)
	session.Send(&battle_api.CancelAuto_Out{
		Round: int16(joinRound),
	})
}

func (this *Battle_MissionLevel) getNextPlayer(round int) int64 {
	if this.battleState == nil {
		return 0
	}
	return this.battleState.GetNextPlayer(round)
}

func (this *Battle_MissionLevel) NextRound(params *module.NextRoundParams) {
	if this.battleStatus == battle.ATK_WIN || this.battleStatus == battle.DEF_WIN {
		return
	}

	battleState := this.battleState
	session := params.Session
	state := module.State(session)

	// 灵宠关卡禁用道具和魂侍
	if battleState.BatType == battle.BT_PET_LEVEL {
		fail.When(params.UseItemId > 0, "pet-level can't use battle-item")
		fail.When(params.UseGhostSkillId > 0, "pet-level can't release ghost")
	}

	// 检查是否使用战斗道具
	if params.UseItemId > 0 {
		module.Item.DelItemByItemId(state.Database, params.UseItemId, 1, tlog.IFR_BATTLE_ITEM, xdlog.ET_MISSION_LEVEL_USE_ITEM)
	}

	//使用魂侍
	if params.UseGhostSkillPosition > 0 {
		valid, reason, _, targetGhost := checkGhost(state, params, battleState)
		fail.When(!valid, reason)
		if battleState.BatType != battle.BT_GHOST_LEVEL {
			_, ghostUsed := state.MissionLevelState.UsedGhost[targetGhost.GhostId]

			fail.When(ghostUsed, "魂侍已使用")
			state.MissionLevelState.UsedGhost[targetGhost.GhostId] = 1
		}
	}

	if battleState.BatType == battle.BT_PVE_LEVEL {
		//动态增加怪物组
		fighterGroup := module.PetVirtualEnv.WarnupNextEnemyGroup(state, battleState)
		if fighterGroup != nil {
			battleState.RunTimeAddFightersGroup(battle.FT_DEF, fighterGroup)
			newGroupRsp := GetNewFighterGroupResponse(0, battle.FT_DEF, fighterGroup, nil)
			session.Send(newGroupRsp)
		}
	}

	if battleState.BatType == battle.BT_DRIVING_LEVEL {
		//仙山拜访动态增加怪物组
		fighterGroup := module.DrivingSword.WarnupNextEnemyGroup(state, battleState)
		if fighterGroup != nil {
			battleState.RunTimeAddFightersGroup(battle.FT_DEF, fighterGroup)
			newGroupRsp := GetNewFighterGroupResponse(0, battle.FT_DEF, fighterGroup, nil)
			session.Send(newGroupRsp)
		}
	}

	battleState.SetPlayerState(params.SkillIndex, int32(params.UseItemId), params.AutoFight, params.PlayerId, params.IsAttacker, params.Position, params.JobIndex, params.UseSwordSoul, params.UseGhostSkillPosition, params.UseGhostSkillId, params.LaunchBattleTotem)

	results, status, nowRound := battleState.NextRound()

	// 检查是否有捕获灵宠, -1表示没有捕捉灵宠
	catchPetId := battleState.GetCatchBattlePetId()
	if catchPetId != -1 {
		state.MissionLevelState.CatchedBattlePetId = catchPetId
		module.Mission.NotifyCatchBattlePet(session, catchPetId)
	}

	// 伙伴关卡如果伙伴死亡战斗结束
	if this.battleState.BatType == battle.BT_BUDDY_LEVEL {
		if this.battleState.HaveBuddyDeadExcludePet() {
			status = battle.DEF_WIN
		}
	}

	if state.MissionLevelState.MaxRound > 0 && nowRound+1 >= state.MissionLevelState.MaxRound && status != battle.ATK_WIN {
		status = battle.DEF_WIN
	}

	this.battleStatus = status

	//检查是否有召唤怪物操作
	/*
		if status == battle.TRIGGER_CALL_ENEMYS {
			result := results[len(results)-1]
			callEnemysResponse := &battle_api.CallNewEnemys_Out{}
			callEnemysResponse.CallInfo = make([]battle_api.CallNewEnemys_Out_CallInfo, 0)

			for index, calls := range result.CallEnemys {
				real_call_infos := CallEnemys(session, calls)
				callEnemysResponse.CallInfo = append(callEnemysResponse.CallInfo, battle_api.CallNewEnemys_Out_CallInfo{
					Ftype:       int8(result.Type),
					Position:    int8(result.FighterPos),
					AttackIndex: int8(index),
					Enemys:      real_call_infos,
				})
			}
			session.Send(callEnemysResponse)
		}
	*/
	rsp := GetNextRoundResponse(results, status, nowRound, battleState, nil, state.MissionLevelState.SkillReleaseNum)
	session.Send(rsp)

	if (status == battle.ATK_WIN || status == battle.DEF_WIN) && this.notAward {
		this.LeaveBattle(session)
		return
	}

	switch status {
	case battle.ATK_WIN:
		this.DoWin(session)
	case battle.DEF_WIN:
		this.DoLose(session)
	}
}

func (this *Battle_MissionLevel) Escape(session *net.Session) {
	state := module.State(session)
	// 难度关卡逃跑处理
	switch this.battleState.BatType {
	case battle.BT_HARD_LEVEL:
		module.Mission.HardLevelLose(state, this.battleState.Defenders)
		return
	case battle.BT_RAINBOW_LEVEL:
		this.saveAttackerInfo(state)
		module.Rainbow.BattleLose(state)
		return
	case battle.BT_PVE_LEVEL:
		state.PetVirtualEnvState.EnemyNum = int16(this.battleState.Defenders.GetDead())
		module.PetVirtualEnv.BattleLose(session, xdlog.ET_PET_VIRTUAL_LOSE)
	case battle.BT_RECOVER_BOAT:
		//rpc 通知互动服 夺回战斗失败
		rpc.RemoteEscortBoatRecoverBattleLose(state.PlayerId, this.boatId)
	}

	// 副本区域关卡逃跑处理
	this.saveAttackerInfo(state)
}

// 召唤灵宠
func (this *Battle_MissionLevel) CallBattlePet(session *net.Session, petId int32, petLevel int16, petSkillLv int16) bool {
	state := module.State(session)
	battlePet := battle_pet_dat.GetBattlePetWithEnemyId(petId)

	// 灵宠关卡判定：同一个灵宠只能在死亡后才能继续召唤
	if this.GetBattle().BatType == battle.BT_PET_LEVEL {
		for _, f := range this.battleState.Attackers.Fighters {
			fail.When((f != nil && f.RoleId == int(petId) && f.IsBattlePetAlive()), "can't call battle-pet. pet is alive now")
		}
	}

	var playerFighter *battle.Fighter
	// 找玩家，扣精气
	for _, f := range this.battleState.Attackers.Fighters {
		if f != nil && f.Kind == battle.FK_PLAYER {
			fail.When(f.Power < int(battlePet.CostPower), "[power] can't CallBattlePet.")
			playerFighter = f
			break
		}
	}

	fail.When(playerFighter == nil, "player fighter not found")
	playerFighter.Power -= int(battlePet.CostPower)

	pos := module.GetBattlePetPos(battlePet.LivePos, this.battleState.Attackers.Fighters)
	if pos < 0 {
		return false
	}
	fighter := module.NewFighterForBattlePet(state.PlayerId, battlePet.PetId, pos, battlePet.RoundAttack, battlePet.LiveRound, this.battleState.GetRounds(), petLevel, petSkillLv)
	// 添加灵宠到战场
	this.battleState.RuntimeAddFighter(battle.FT_ATK, fighter)

	rsp := &battle_api.CallBattlePet_Out{}
	rsp.PlayerPower = int16(playerFighter.Power)
	rsp.Skills = make([]battle_api.CallBattlePet_Out_Skills, len(fighter.Skills))

	ghosts := []battle_api.BattleRole_Ghosts{}
	for _, roleGhost := range fighter.Ghosts {
		ghosts = append(ghosts, battle_api.BattleRole_Ghosts{
			GhostId:      roleGhost.GhostId,
			GhostStar:    roleGhost.GhostStar,
			GhostSkillId: int32(roleGhost.GhostSkillId),
		})
	}

	rsp.Role = battle_api.BattleRole{
		Kind:                battle_api.FighterKind(fighter.Kind),
		PlayerId:            fighter.PlayerId,
		RoleId:              int32(fighter.RoleId),
		RoleLevel:           int16(fighter.Level),
		Position:            int32(fighter.Position),
		FashionId:           fighter.FashionId,
		Health:              int32(fighter.Health),
		MaxHealth:           int32(fighter.MaxHealth),
		Power:               int16(fighter.Power),
		MaxPower:            int16(fighter.MaxPower),
		SunderValue:         int16(fighter.GetSunderValue()),
		SunderMaxValue:      int16(fighter.SunderMaxValue),
		SunderMinHurtRate:   int16(fighter.SunderMinHurtRate),
		SunderEndHurtRate:   int16(fighter.SunderEndHurtRate),
		SunderEndDefendRate: int16(fighter.SunderEndDefendRate),

		Speed: int32(fighter.Speed),

		GhostPower:        int32(fighter.GetGhostPower()),
		Ghosts:            ghosts,
		GhostShieldValue:  int32(fighter.GhostShieldValue),
		CouldUseSwordSoul: fighter.SwordSoulValue > 0,
	}

	for i, skill := range fighter.Skills {
		if skill == nil {
			rsp.Skills[i] = battle_api.CallBattlePet_Out_Skills{
				Skill: battle_api.SkillInfo{
					SkillId: -2,
				},
			}
			continue
		}
		rsp.Skills[i] = battle_api.CallBattlePet_Out_Skills{
			Skill: battle_api.SkillInfo{
				SkillId:  int16(skill.SkillId),
				IncPower: int8(skill.IncPower),
				DecPower: int8(skill.DecPower),
			},
		}
	}
	rsp.Success = true

	session.Send(rsp)
	return true
}

// 复活
func (this *Battle_MissionLevel) Relive(session *net.Session) {
	state := module.State(session)
	fail.When(state.MissionLevelState.LevelType == battle.BT_HARD_LEVEL, "难度关卡不能复活")
	fail.When(state.MissionLevelState.LevelType == battle.BT_RAINBOW_LEVEL, "彩虹关卡不能复活")

	roleLevel := module.Role.GetMainRole(state.Database).Level
	// 如果角色当前等级大于了免费复活要求的等级，就要对元宝进行检查
	if roleLevel > mission_dat.FREE_RELIVE_LEVEL {
		var costIngot int32 = mission_dat.FIRST_RELIVE_INGOT
		if state.MissionLevelState.NextReliveCostIngot > 0 {
			costIngot = state.MissionLevelState.NextReliveCostIngot
		}

		// 扣元宝
		module.Player.DecMoney(state.Database, state.MoneyState, int64(costIngot), player_dat.INGOT, tlog.MFR_MISSION_LEVEL_RELIVE, xdlog.ET_MISSION_LEVEL_RELIVE)
		state.MissionLevelState.NextReliveCostIngot = costIngot * 2
	}

	state.MissionLevelState.ReliveState = module.LEVEL_RELIVE_STATE_RELIVING

	// 保存怪物状态，为复活后导入
	this.saveDefendInfoForEscape(state)
	//倒入上复活前的回合数
	round := this.battleState.GetRounds()
	round++

	state.MissionLevelState.ReliveRound += round

	// 移除战场关联，假如断线可以重建关卡
	state.Battle = nil
}

func (this *Battle_MissionLevel) SetSkill(session *net.Session, posIdx int8, skillIdx int8) {
	state := module.State(session)
	side := this.battleState.Attackers
	fighter := side.Fighters[posIdx-1]
	if fighter != nil && fighter.PlayerId == state.PlayerId {
		fighter.SetBuddySkill(int(skillIdx))
		return
	}
	side = this.battleState.Defenders
	fighter = side.Fighters[posIdx-1]
	if fighter != nil && fighter.PlayerId == state.PlayerId {
		fighter.SetBuddySkill(int(skillIdx))
		return
	}
	fail.When(true, "找不到伙伴")
}

func (this *Battle_MissionLevel) UseBuddySkill(session *net.Session, posIndex int8, skillIndex int8) {
	state := module.State(session)

	side := this.battleState.Attackers
	buddy := side.Fighters[posIndex-1]
	fail.When(buddy == nil, "指定位置不存在伙伴")
	fail.When(buddy.PlayerId != state.PlayerId, "操作不属于自己的伙伴")

	buddy.SetBuddySkill(int(skillIndex))
}

/*
流程： 计算回合 [计算boss战得分] [发放关卡内特殊奖励 消耗关卡次数] 准备或发放通关奖励 更新任务 保存玩家战斗状态 清理战场
*/
func (this *Battle_MissionLevel) DoWin(session *net.Session) {
	state := module.State(session)

	round := this.battleState.GetRounds()

	//round++

	// 统计本次关卡的回合数
	state.MissionLevelState.TotalRound += round
	//本次怪物组回合
	state.MissionLevelState.ReliveRound += round

	// 标记当前怪物组已通过
	state.MissionLevelState.MarkPassEnemy()

	if this.isBoss {
		//state.MissionLevelState.BossScore = this.countBossBattleScore(state)
	}

	//ReliveRound记录与一组怪物战斗的回合数，在此重置
	state.MissionLevelState.ReliveRound = 0

	// 策划规则：没有通关，但获得了贵重道具如：灵宠契约球，则消耗关卡次数
	// 如果是区域关卡中捕捉到灵宠，就先累计关卡进入次数，因为在退出区域关卡时有很多种情况，不好判断，所以在这里提前处理。通关后根据状态来判断是否累计了次数
	if state.MissionLevelState.CatchedBattlePetId > 0 {
		module.Mission.AwardCatchedBattlePet(state, xdlog.ET_MISSION_LEVEL_WIN)

		if this.battleState.BatType == battle.BT_MISSION_LEVEL && !state.MissionLevelState.HaveReduceLevelDialyNum {
			state.MissionLevelState.HaveReduceLevelDialyNum = true

			levelInfo := mission_dat.GetMissionLevelById(state.MissionLevelState.LevelId)
			if levelInfo.DailyNum > 0 {
				levelRecord := module.Mission.GetMissionLevelRecord(state.Database, levelInfo.Id)
				levelRecord.DailyNum += 1
				state.Database.Update.PlayerMissionLevelRecord(levelRecord)
			}
		}
	}

	//保存攻击者信息需要在彩虹关卡结算之前
	this.saveAttackerInfo(state)

	//准备奖励内容或者直接奖励
	if this.battleState.BatType == battle.BT_DRIVING_SWORD_BF_LEVEL {
		module.DrivingSword.VisitingAward(session, xdlog.ET_DRIVING_VISIT_AWARD)
	} else if this.battleState.BatType == battle.BT_RAINBOW_LEVEL {
		//彩虹关卡奖励走独立的奖励逻辑
		// 计算奖励，并写入数据库。写入数据库是为了防止推出重新进入刷新奖励内容
		module.Rainbow.BattleWin(state, xdlog.ET_RAINBOW)
	} else if this.battleState.BatType == battle.BT_PVE_LEVEL {
		state.PetVirtualEnvState.EnemyNum = int16(this.battleState.Defenders.GetDead())
		module.PetVirtualEnv.BattleWin(session, xdlog.ET_PET_VIRTUAL_WIN)
	} else if this.battleState.BatType == battle.BT_DRIVING_LEVEL {
		//state.DrivingSword.EnemyNum = int16(this.battleState.Defenders.GetDead())
		module.DrivingSword.BattleWin(session, xdlog.ET_DRIVING_SWORD_WIN)
	} else if this.battleState.BatType == battle.BT_HIJACK_BOAT {
		rpc.RemoteEscortBoatHijackBattleWin(state.PlayerId, this.boatId)
	} else if this.battleState.BatType == battle.BT_RECOVER_BOAT {
		rpc.RemoteEscortBoatRecoverBattleWin(state.PlayerId, this.boatId)
	} else {
		// 关卡固定奖励是在开过宝箱后，如果没有宝箱就直接进行固定奖励
		if state.MissionLevelState.AwardBox {
			// 奖励宝箱抽取次数
			var awardCount int8 = mission_dat.LEVEL_BOX_AWARD_COUNT
			if this.battleState.BatType == battle.BT_MISSION_LEVEL || this.battleState.BatType == battle.BT_HARD_LEVEL {
				awardCount = mission_dat.CalLevelStarByRound(state.MissionLevelState.LevelId, int8(state.MissionLevelState.TotalRound))
			}
			state.MissionLevelState.BoxState.AwardOpenCount = awardCount

			if this.battleState.BatType == battle.BT_HARD_LEVEL {
				state.MissionLevelState.RandomBox.RandomBoxOpenCount = mission_dat.RANDDOM_BOX_AWARD_COUNT
			}

		} else {
			module.Mission.AwardLevel(state, xdlog.ET_MISSION_LEVEL_AWARD)
		}
	}

	//更新任务状态
	this.updateQuest(state, this.battleState.BatType, this.isBoss, xdlog.ET_QUEST)

	// 战斗成功后销毁战场，表明当前已在关卡中，不在战场内
	this.LeaveBattle(session)
	// 移除战场关联，假如断线可以重建关卡
	state.Battle = nil
	// 清理逃跑时记录的怪物数据
	state.MissionLevelState.DefendEnemy = nil
}

func (this *Battle_MissionLevel) DoLose(session *net.Session) {
	/*
		对抗失败后，玩家可以元宝买活，接着战斗；要么直接回到城镇
	*/

	state := module.State(session)

	// 记录玩家失败的战场
	state.MissionLevelState.ReliveState = module.LEVEL_RELIVE_STATE_CAN_RELIVE
	// 清理逃跑时记录的怪物数据
	state.MissionLevelState.DefendEnemy = nil
	// 统计本次关卡的回合数
	round := this.battleState.GetRounds()
	round++
	state.MissionLevelState.TotalRound += round

	switch this.battleState.BatType {
	case battle.BT_TOWER_LEVEL:
		//module.Tower.BattleLose(state)
	case battle.BT_HARD_LEVEL:
		module.Mission.HardLevelLose(state, this.battleState.Defenders)
		module.Mission.LeaveMissionLevel(state)
	case battle.BT_RAINBOW_LEVEL:
		//保存攻击者信息需要在彩虹关卡结算之前
		this.saveAttackerInfo(state)
		module.Rainbow.BattleLose(state)
	case battle.BT_PVE_LEVEL:
		state.PetVirtualEnvState.EnemyNum = int16(this.battleState.Defenders.GetDead())
		module.PetVirtualEnv.BattleLose(session, xdlog.ET_PET_VIRTUAL_LOSE)
	case battle.BT_MISSION_LEVEL:
		this.saveAttackerInfo(state)
	case battle.BT_DRIVING_LEVEL:
		//state.DrivingSword.EnemyNum = int16(this.battleState.Defenders.GetDead())
		module.DrivingSword.BattleLose(session)
	case battle.BT_RECOVER_BOAT:
		rpc.RemoteEscortBoatRecoverBattleLose(state.PlayerId, this.boatId)
	}

	/*
		ps: 这里不设置state.Battle = nil是因为，战败后会提示是否复活，如果选否则回到城镇（战场就没有断线重建的必要），选是会在复活后进行设置battle=nil
	*/
}

func (this *Battle_MissionLevel) LeaveBattle(session *net.Session) {
	//fmt.Println("session ", session)
	//state := module.State(session)
	//if this.battleState.BatType == battle.BT_RECOVER_BOAT {
	//	//rpc 通知互动服 夺回战斗失败
	//	rpc.RemoteEscortBoatRecoverBattleLose(state.PlayerId, this.boatId)
	//}
	// 清理战场
	this.battleState = nil
}

func (this *Battle_MissionLevel) enemyDefendSide(battleType int8, enemyId int32, defendEnemy map[int]module.EnemyState) *battle.SideInfo {
	// boss关卡 使用阵型
	// 没有配置(1~5)怪物 使用阵型
	levelEnemy := mission_dat.GetMissionLevelEnemyById(enemyId)
	this.bestRound = levelEnemy.BestRound
	this.isBoss = levelEnemy.IsBoss

	var fighters []*battle.Fighter
	// 有记录关卡怪物战场数据(逃跑)
	if defendEnemy != nil {
		fighters = make([]*battle.Fighter, module.ALL_FIGHTER_POS_NUM)
		for pos, enemy := range defendEnemy {
			//如果要对怪物进行属性加成，请修改NewFighterForEnemy第二个参数
			f := module.NewFighterForEnemy(enemy.EnemyId, 0, pos)
			f.Health = enemy.Health
			fighters[pos-1] = f
		}
	} else {
		// 如果是boss战或者怪物组没有配置随机的怪物，直接从怪物上阵表找关卡怪物
		if levelEnemy.IsBoss ||
			(levelEnemy.Monster1Id == 0 && levelEnemy.Monster2Id == 0 && levelEnemy.Monster3Id == 0 && levelEnemy.Monster4Id == 0 && levelEnemy.Monster5Id == 0) {
			battleSide := module.NewBattleSideWithEnemyDeployForm(battleType, enemyId)
			this.loadLevelBattlePet(levelEnemy.Id, &battleSide.Fighters)
			return battleSide
		}
		fighters = module.NewEnemyFighterGroup(enemyId)
	}

	this.loadLevelBattlePet(levelEnemy.Id, &fighters)

	return &battle.SideInfo{
		Groups:   [][]*battle.Fighter{fighters},
		Fighters: fighters,
	}
}

// 加载关卡灵宠
func (this *Battle_MissionLevel) loadLevelBattlePet(levelEnemyId int32, fighters *[]*battle.Fighter) {
	levelPet, ok := battle_pet_dat.GetLevelBattlePet(levelEnemyId)
	if !ok {
		return
	}

	randNum := rand.Intn(100) + 1
	if int8(randNum) > levelPet.Rate {
		return
	}

	pos := module.GetBattlePetPos(levelPet.Pet.LivePos, *fighters)
	(*fighters)[pos-1] = module.NewFighterForBattlePet(0, levelPet.Pet.PetId, pos, levelPet.Pet.RoundAttack, levelPet.LiveRound, 0, 0, 0)
}

// 当逃跑时保存战场上出现怪物信息，以便玩家再次战斗时相遇
func (this *Battle_MissionLevel) saveDefendInfoForEscape(state *module.SessionState) {
	state.MissionLevelState.DefendEnemy = make(map[int]module.EnemyState)
	for _, f := range this.battleState.Defenders.Fighters {
		if f != nil && !f.IsBattlePet { // 不导出随机出现的灵宠
			state.MissionLevelState.DefendEnemy[f.Position] = module.EnemyState{int32(f.RoleId), f.Health}
		}
	}
}

// 保存攻方战场状态，延续至下一组怪
func (this *Battle_MissionLevel) saveAttackerInfo(state *module.SessionState) {
	// 通天塔关卡无需保存角色状态
	if this.battleState.BatType == battle.BT_TOWER_LEVEL {
		return
	}

	attackerInfo := &state.MissionLevelState.AttackerInfo
	newAttackInfo := this.battleState.ExportAttackersInfo()

	// 保存新上阵角色状态，覆盖已有角色状态
	for roleId, buffs := range newAttackInfo.Buffs {
		attackerInfo.Buffs[roleId] = buffs
	}

	for roleId, skills := range newAttackInfo.Skills {
		attackerInfo.Skills[roleId] = skills
		//伙伴技能使用次数可能在战斗中变化，故需要导出保存
		if !role_dat.IsMainRole(int8(roleId)) {
			for _, skill := range skills {
				if skill.SkillId > skill_dat.SKILL_IS_NULL {
					state.MissionLevelState.SkillReleaseNum[skill.SkillId] = skill.ReleaseNum
				}
			}
		}
	}

	for roleId, f := range newAttackInfo.Fighters {
		attackerInfo.Fighters[roleId] = f
	}

	for pid, ghostInfo := range newAttackInfo.GhostInfo {
		attackerInfo.GhostInfo[pid] = ghostInfo
	}
}

func (this *Battle_MissionLevel) importGhostInfo(state *module.SessionState) {
	currentAttackInfo := this.battleState.ExportAttackersInfo()
	attackerInfo := &state.MissionLevelState.AttackerInfo
	for pid, ghostInfo := range attackerInfo.GhostInfo {
		currentAttackInfo.GhostInfo[pid] = ghostInfo
	}
	this.battleState.ImportAttackersInfo(currentAttackInfo)
}

// 导入上一场战斗玩家角色数据
func (this *Battle_MissionLevel) importAttackInfo(state *module.SessionState) {
	// 通天塔关卡无需导入角色状态
	if this.battleState.BatType == battle.BT_TOWER_LEVEL {
		return
	}

	//获取当前技能信息
	newAttackInfo := this.battleState.ExportAttackersInfo()
	state.MissionLevelState.AttackerInfo.Skills = newAttackInfo.Skills

	//技能次数可能在战场外改变，所以需要用用 mission level state 中的技能数据 覆盖 attackerInfo 的数据
	for roleId, skills := range state.MissionLevelState.AttackerInfo.Skills {
		if !role_dat.IsMainRole(int8(roleId)) {
			for i, skill := range skills {
				if releaseNum, ok := state.MissionLevelState.SkillReleaseNum[skill.SkillId]; ok {
					skills[i].ReleaseNum = releaseNum
				}
			}
			state.MissionLevelState.AttackerInfo.Skills[roleId] = skills
		}
	}

	this.battleState.ImportAttackersInfo(&state.MissionLevelState.AttackerInfo)
}

//计算Boss战 战斗得分
func (this *Battle_MissionLevel) countBossBattleScore(state *module.SessionState) int32 {
	// 标记已打败boss
	state.MissionLevelState.PassBoss = true
	// boss战计算评分
	var totalHurt int // 所有角色输出总伤害
	for _, fighter := range this.battleState.Attackers.Fighters {
		if fighter != nil {
			totalHurt += fighter.TotalHurt
		}
	}

	// 得分 ＝ ceil(总伤害/回合数)
	return int32(math.Ceil(float64(totalHurt) / float64(state.MissionLevelState.ReliveRound)))
}

//更新任务状态
func (this *Battle_MissionLevel) updateQuest(state *module.SessionState, batType int, isBoss bool, xdEventType int32) {
	switch batType {
	case battle.BT_MISSION_LEVEL:
		module.Quest.RefreshQuest(state, quest_dat.QUEST_TYPE_MISSION_ENEMY, xdEventType)
		module.Quest.RefreshQuestForBeatEnemyGroup(state.Database, state.MissionLevelState.EnemyId)
		//消灭所有怪物后 结算通关任务
		if state.MissionLevelState.HasKilledThoseMustDie() {
			module.Quest.RefreshQuest(state, quest_dat.QUEST_TYPE_WIN_MISSION, xdEventType)
		}

		if isBoss {
			module.Quest.RefreshDailyQuest(state, quest_dat.DAILY_QUEST_CLASS_MISSION_LEVEL_BOSS)
		}
	}
}

func checkGhost(state *module.SessionState, params *module.NextRoundParams, battleState *battle.BattleState) (ok bool, reason string, fighter *battle.Fighter, ghost *battle.FightGhost) {
	fighter = battleState.Attackers.Fighters[params.UseGhostSkillPosition-1]
	if fighter.Health <= battle.MIN_HEALTH {
		return false, "角色阵亡", nil, nil
	}
	if fighter.GhostInfo.GhostPower < battle.FULL_GHOST_POWER {
		return false, "魂力不足", nil, nil
	}

	targetGhost := fighter.GetMainGhost()
	if targetGhost == nil || targetGhost.GhostSkillId != int(params.UseGhostSkillId) {
		return false, "没有可用的魂侍技能", nil, nil
	}

	return true, "", fighter, targetGhost
}
