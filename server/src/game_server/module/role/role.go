package role

import (
	"core/fail"
	"core/time"
	"game_server/api/protocol/role_api"
	"game_server/battle"
	"game_server/dat/fight_num_dat"
	"game_server/dat/player_dat"
	"game_server/dat/quest_dat"
	"game_server/dat/role_dat"
	"game_server/dat/team_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/module/rpc"
	"game_server/tlog"
	"game_server/xdlog"
)

// 通过rold id 查找角色
func findRoleByRoleId(db *mdb.Database, roleId int8, requireInTeam bool) *mdb.PlayerRole {
	var role *mdb.PlayerRole = nil
	db.Select.PlayerRole(func(row *mdb.PlayerRoleRow) {
		if row.RoleId() == roleId {
			role = row.GoObject()
			fail.When(requireInTeam && row.Status() == role_dat.ROLE_STATUS_ININN, "角色不在队伍中")
			row.Break()
		}
	})
	return role
}

// 为指定角色加经验
func addRoleExp(db *mdb.Database, roleId int8, addExp int64, mainRoleId int8, playerExpFlowReason int32) {
	fail.When(addExp < 1 || mainRoleId < 1, "addRoleExp error")

	var newLevel int16
	var newExp int64

	role := findRoleByRoleId(db, roleId, false)

	if !role_dat.IsMainRole(roleId) {
		mainRoleLevel := findRoleByRoleId(db, mainRoleId, false).Level
		// 已达到主角等级，且经验溢出，则本次不记录增加的经验
		if role.Level >= mainRoleLevel && role.Exp >= role_dat.GetRoleLevelExp(role.Level) {
			return
		}

		newLevel, newExp = getNewLevelAndExp(role.Level, role.Exp+addExp)

		// 伙伴等级大于主角等级 level 的情况
		if newLevel > mainRoleLevel {
			// 如果伙伴等级是从小于主角的等级开始升级的，如果发现升级后比主角等级还大，那么需要设置成主角相同的等级
			if role.Level < mainRoleLevel {
				newLevel = mainRoleLevel
			} else {
				// 目前运营有活动直接可买高等级伙伴，所以对高等级的伙伴保持等级和经验
				newLevel = role.Level
			}
			newExp = role_dat.GetRoleLevelExp(newLevel)
		}

	} else {
		newLevel, newExp = getNewLevelAndExp(role.Level, role.Exp+addExp)
		// 主角等级变化
		if newLevel > role.Level {
			//  玩家等级变化会开启功能
			module.Player.RefreshPlayerLevelFuncKey(db, newLevel)

			// 更新玩家每日任务
			module.Quest.UpdatePlayerDailyQuest(db, newLevel)
			// 向互动服务器更新主角等级
			rpc.RemoteUpdateMainRoleLevel(db.PlayerId(), newLevel)
			// 升级运营活动
			module.Event.LevelUp(db, role.Level, newLevel)
		}
		playerExpFlow(db, int32(addExp), int32(role.Level), int32(newLevel), playerExpFlowReason)
		xdlog.LevelLog(db, int32(addExp), int32(role.Level), int32(newLevel))
	}

	// 角色升级更新绝招
	if newLevel > role.Level {
		playerFame := db.Lookup.PlayerFame(db.PlayerId())
		module.Skill.UpdateSkill(db, roleId, int16(role.FriendshipLevel), playerFame.Level, newLevel)
	}

	oldLevel := role.Level
	if newLevel > role.Level {
		role.Timestamp = time.GetNowTime()
	}
	role.Level, role.Exp = newLevel, newExp
	db.Update.PlayerRole(role)

	// 等级排行榜数据更新
	if newLevel > oldLevel && newLevel >= player_dat.RANK_OPEN_MIN_LEVEL && (roleId == role_dat.PLAYER_BOY_ROLE_ID || roleId == role_dat.PLAYER_GIRL_ROLE_ID) {
		// 更新战力排行榜
		playerFightNum := db.Lookup.PlayerFightNum(db.PlayerId())
		if playerFightNum == nil {
			module.Player.AddUpdateFightNum(db, 0)
		} else {
			module.Player.AddUpdateFightNum(db, playerFightNum.FightNum)
		}
		module.Player.ForceUpdateFightNum()
		// 向互动服务器更新战力
		rpc.RemoteUpdatePlayerFightNum(db.PlayerId(), playerFightNum.FightNum)
		module.Player.AddUpdateLevel(db, newLevel, role.Timestamp)
		module.Player.ForceUpdateLevel()
	}

	//触发支线任务 因为支线任务有等级限制，必须再保存了玩家等级之后再刷新
	if role_dat.IsMainRole(roleId) && newLevel > oldLevel {
		module.Quest.RefreshQuestOnLevelUp(db, oldLevel, newLevel)
	}

	if session, ok := module.Player.GetPlayerOnline(db.PlayerId()); ok {
		module.Notify.SendRoleExpChanged(session, role.RoleId, addExp, role.Exp, role.Level)
	}
}

// 返回新的level, exp, 返回参数为 level int16, exp int64
func getNewLevelAndExp(level int16, exp int64) (int16, int64) {
	mostExp := role_dat.GetRoleLevelExp(role_dat.MAX_ROLE_LEVEL)

	for {
		if level == role_dat.MAX_ROLE_LEVEL && exp >= mostExp {
			exp = mostExp // 满级 经验溢出
			break
		}

		needExp := role_dat.GetRoleLevelExp(level)
		if exp >= needExp {
			level = level + 1
			exp = exp - needExp
		} else {
			break // 不能升级就直接加经验 跳出循环
		}
	}

	return level, exp
}

func getFightNum(db *mdb.Database, roleId int8, total bool, fightNumType role_api.FightnumType) (fighter *battle.Fighter, fightNums map[int16]int32, fightNum int32) {
	playerForm := db.Lookup.PlayerFormation(db.PlayerId())
	var inFormRoleInfos []*module.InFormRoleInfo
	if !total {
		inFormRoleInfos = []*module.InFormRoleInfo{&module.InFormRoleInfo{RoleId: roleId}}
	} else {
		inFormRoleInfos, _ = module.GetBattleBiz.GetInFormRoleForBattle(db, playerForm, false)
	}

	fighters := make([]*battle.Fighter, module.ALL_FIGHTER_POS_NUM)

	fightNums, fightNum = module.GetBattleBiz.SetFighters(db, inFormRoleInfos, fighters, false, true, int8(fightNumType))
	fighter = fighters[0]

	if total && fightNumType == module.FIGHT_FOR_ALL {
		playerFightNum := db.Lookup.PlayerFightNum(db.PlayerId())

		//新战斗力比之前大50以上则更新数据库
		if fightNum-playerFightNum.FightNum >= 50 {
			playerFightNum.FightNum = fightNum
			db.Update.PlayerFightNum(playerFightNum)

			// 更新战力排行榜
			module.Player.AddUpdateFightNum(db, playerFightNum.FightNum)

			// 向互动服务器更新战力
			rpc.RemoteUpdatePlayerFightNum(db.PlayerId(), fightNum)
		}
	}

	return
}

// 把几率百分比转换成几率等级
func calcFighterLevel(fighter *battle.Fighter) (
	hitLevel, critialLevel, sleepLevel, dizzinessLevel, randomLevel, disableSkillLevel, poisoningLevel, blockLevel, destroyLevel, critialHurtLevel, tenacityLevel, dodgeLevel float64) {

	// 公式：几率百分比 = 几率等级*调整系数（0.025）/目标等级
	var probLevel float64
	if fighter.Level < 50 {
		probLevel = float64(50) / 100
	} else {
		probLevel = float64(fighter.Level) / 100
	}

	hitLevel = fighter.Hit * probLevel / battle.HIT_LEVEL_ARG
	critialLevel = fighter.Critial * probLevel / battle.CRITIAL_LEVEL_ARG
	sleepLevel = fighter.Sleep * probLevel / battle.SLEEP_LEVEL_ARG
	dizzinessLevel = fighter.Dizziness * probLevel / battle.DIZZINESS_LEVEL_ARG
	randomLevel = fighter.Random * probLevel / battle.RANDOM_LEVEL_ARG
	disableSkillLevel = fighter.DisableSkill * probLevel / battle.DISABLE_SKILL_LEVEL_ARG
	poisoningLevel = fighter.Poisoning * probLevel / battle.POISONING_LEVEL_ARG
	blockLevel = fighter.Block * probLevel / battle.BLOCK_LEVEL_ARG
	destroyLevel = fighter.Destroy * probLevel / battle.DESTROY_LEVEL_ARG
	critialHurtLevel = fighter.CritialHurt * probLevel / battle.CRITIAL_HURT_LEVEL_ARG
	tenacityLevel = fighter.Tenacity * probLevel / battle.TENACITY_LEVEL_ARG
	dodgeLevel = fighter.Dodge * probLevel / battle.DODGE_LEVEL_ARG

	return
}

func playerExpFlow(db *mdb.Database, expchange, beforelevel, afterlevel, playerExpFlowReason int32) {
	openid := module.Player.GetPlayer(db).User
	role_vip := module.VIP.VIPInfo(db)
	db.AddTLog(tlog.NewPlayerExpFlow(openid, expchange, beforelevel, afterlevel, 0, playerExpFlowReason, 0, int32(role_vip.Level)))
}

func recruitBuddy(state *module.SessionState, roleId int8) int8 {
	recruitBuddyDat := role_dat.GetRecruitBuddyDat(roleId)
	playerQuest := state.Database.Lookup.PlayerQuest(state.PlayerId)
	playerQuestDat := quest_dat.GetQuestById(playerQuest.QuestId)
	requireQuestDat := quest_dat.GetQuestById(recruitBuddyDat.QuestId)

	fail.When(playerQuestDat.Order < requireQuestDat.Order, "未达到开启任务")

	role := findRoleByRoleId(state.Database, roleId, false)
	fail.When(role != nil, "已获伙伴无法招募")
	module.Item.DelItemByItemId(state.Database, recruitBuddyDat.FavouriteItem, recruitBuddyDat.FavouriteCount, tlog.IFR_RECIT_BUDDY, xdlog.ET_FRIEND_SHIP)
	return module.Role.AddBuddyRole(state, roleId, recruitBuddyDat.InitLevel)
}

func changeRoleStatus(db *mdb.Database, roleId int8, status int8) bool {
	fail.When(roleId == 1 || roleId == 2, "main role can not operate")
	fail.When(status != role_dat.ROLE_STATUS_ININN && status != role_dat.ROLE_STATUS_NOMAL, "非法操作")
	var role *mdb.PlayerRole
	var num int
	result := true
	db.Select.PlayerRole(func(row *mdb.PlayerRoleRow) {
		if row.RoleId() == roleId {
			role = row.GoObject()
			if status != role_dat.ROLE_STATUS_ININN {
				num++
			}
		} else if row.Status() != role_dat.ROLE_STATUS_ININN {
			num++
		}
	})
	if num > role_dat.MAX_ROLE_NUM_USING {
		return false
	}

	fail.When(role == nil, "can not find the role")

	if status == role_dat.ROLE_STATUS_NOMAL {
		// 入列玩家
		role.Status = role_dat.ROLE_STATUS_NOMAL //置为入队的状态
		db.Update.PlayerRole(role)
	} else {

		if playerMultiLevel := db.Lookup.PlayerMultiLevelInfo(db.PlayerId()); playerMultiLevel != nil && playerMultiLevel.BuddyRoleId == role.RoleId {
			playerMultiLevel.BuddyRoleId = team_dat.POS_NO_ROLE
			db.Update.PlayerMultiLevelInfo(playerMultiLevel)
		}

		if module.Team.IsRoleInForm(db, roleId) {
			module.Team.DownFormRole(db, roleId)
		}

		role.Status = role_dat.ROLE_STATUS_ININN //置为在客栈的状态
		db.Update.PlayerRole(role)
	}

	return result
}

func getInnRoleList(db *mdb.Database) []role_api.GetInnRoleList_Out_RoleList {
	role_list := make([]role_api.GetInnRoleList_Out_RoleList, 0)
	own_roles := make(map[int8]int8)
	// 客栈中的已有角色
	db.Select.PlayerRole(func(row *mdb.PlayerRoleRow) {
		own_roles[row.RoleId()] = row.RoleId()
		if row.Status() == role_dat.ROLE_STATUS_ININN {
			_, _, fightNum := getFightNum(db, row.RoleId(), false, role_dat.FIGHT_FOR_ALL)
			role_list = append(role_list, role_api.GetInnRoleList_Out_RoleList{
				RoleId:          row.RoleId(),
				FriendshipLevel: int16(row.FriendshipLevel()),
				FightNum:        fightNum,
				RoleLevel:       row.Level(),
				Operate:         role_dat.ROLE_STATUS_NOMAL,
			})
		}
	})

	// 可以招募的角色的信息
	recruitBuddys := role_dat.GetRecruitBuddys()
	playerQuest := db.Lookup.PlayerQuest(db.PlayerId())
	quest := quest_dat.GetQuestById(playerQuest.QuestId)

	for roleid, buddyInfo := range recruitBuddys {
		recruitBuddyDat := role_dat.GetRecruitBuddyDat(roleid)
		if _, ok := own_roles[roleid]; !ok {
			needQuest := quest_dat.GetQuestById(recruitBuddyDat.QuestId)
			if quest.Order >= needQuest.Order {
				//还没招募然后可以招募的伙伴
				role_list = append(role_list, role_api.GetInnRoleList_Out_RoleList{
					RoleId:          roleid,
					FriendshipLevel: 1,
					FightNum:        int32(getRoleFightNum(roleid, buddyInfo.InitLevel)),
					RoleLevel:       buddyInfo.InitLevel,
					Operate:         role_dat.ROLE_STATUS_ININN,
				})
			}
		}
	}

	return role_list
}

func getRoleFightNum(roleId int8, initLevel int16) float64 {
	roleInfo := role_dat.GetRoleLevelInfo(roleId, initLevel)
	fighter := &battle.Fighter{
		Kind:                battle.FK_BUDDY,                // 是否是主角
		PlayerId:            0,                              // 玩家ID
		RoleId:              int(roleId),                    // 角色ID
		Level:               int(initLevel),                 // 等级
		FriendshipLevel:     1,                              // 羁绊等级
		Position:            1,                              // 站位ID
		Health:              int(roleInfo.Health),           // 生命
		MaxHealth:           int(roleInfo.Health),           // 生命值
		Speed:               roleInfo.Speed,                 // 速度
		Attack:              roleInfo.Attack,                // 普通攻击
		Defend:              roleInfo.Defence,               // 普通防御
		Block:               roleInfo.Block,                 // 格挡
		Critial:             roleInfo.Critial,               // 暴击
		CritialHurt:         roleInfo.CritialHurt,           // 必杀
		Dodge:               roleInfo.Dodge,                 // 闪避
		Hit:                 roleInfo.Hit,                   // 命中
		Cultivation:         roleInfo.Cultivation,           // 修为
		Sleep:               float64(roleInfo.Sleep),        // 睡眠抗性
		Dizziness:           float64(roleInfo.Dizziness),    // 眩晕抗性
		Random:              float64(roleInfo.Random),       // 混乱抗性
		DisableSkill:        float64(roleInfo.DisableSkill), // 封魔抗性
		Poisoning:           float64(roleInfo.Poisoning),    // 中毒抗性
		SunderMaxValue:      int(roleInfo.SunderMaxValue),
		SunderMinHurtRate:   int(roleInfo.SunderHurtRate),
		SunderEndHurtRate:   int(roleInfo.SunderEndHurtRate),
		SunderEndDefendRate: int(roleInfo.SunderEndDefendRate), // 破甲后减防
	}
	var fightNum float64
	fightNum += float64(fighter.Level * fight_num_dat.LEVEL_RATE)                   // 等级
	fightNum += float64(fighter.MaxHealth) * float64(fight_num_dat.MAX_HEALTH_RATE) // 生命
	fightNum += fighter.Cultivation * fight_num_dat.CULTIVATION_RATE                // 内力
	fightNum += fighter.Speed * fight_num_dat.SPEED_RATE                            // 速度
	fightNum += fighter.Attack * fight_num_dat.ATTACK_RATE                          // 攻击
	fightNum += fighter.Defend * fight_num_dat.DEFEND_RATE                          // 防御
	fightNum += float64(fighter.SunderMaxValue * fight_num_dat.SUNDER_RATE)         // 破甲
	fightNum += fighter.CritialLevel * fight_num_dat.CRITIAL_LEVEL_RATE             // 暴击等级
	fightNum += fighter.TenacityLevel * fight_num_dat.TENACITY_LEVEL_RATE           // 韧性等级
	fightNum += fighter.HitLevel * fight_num_dat.HIT_LEVEL_RATE                     // 命中等级
	fightNum += fighter.DodgeLevel * fight_num_dat.DODGE_LEVEL_RATE                 // 闪避等级
	fightNum += fighter.DestroyLevel * fight_num_dat.DESTROY_LEVEL_RATE             // 破击等级
	fightNum += fighter.BlockLevel * fight_num_dat.BLOCK_LEVEL_RATE                 // 格挡等级

	fightNum += fighter.SleepLevel * fight_num_dat.SLEEP_LEVEL_RATE
	fightNum += fighter.DizzinessLevel * fight_num_dat.DIZZINESS_LEVEL_RATE
	fightNum += fighter.RandomLevel * fight_num_dat.RANDOM_LEVEL_RATE
	fightNum += fighter.DisableSkillLevel * fight_num_dat.DISABLE_SKILL_LEVEL_RATE
	fightNum += fighter.PoisoningLevel * fight_num_dat.POISONING_LEVEL_RATE

	fightNum += fighter.CritialHurtLevel * fight_num_dat.CRITIAL_HURT_LEVEL_RATE // 必杀等级
	fightNum += fighter.Critial * fight_num_dat.CRITIAL_RATE                     // 暴击百分比
	fightNum += fighter.Tenacity * fight_num_dat.TENACITY_RATE                   // 韧性百分比
	fightNum += fighter.Hit * fight_num_dat.HIT_RATE                             // 命中百分比
	fightNum += fighter.Dodge * fight_num_dat.DODGE_RATE                         // 闪避百分比
	fightNum += fighter.Destroy * fight_num_dat.DESTROY_RATE                     // 破击百分比
	fightNum += fighter.Block * fight_num_dat.BLOCK_RATE                         // 格挡百分比

	fightNum += fighter.Sleep * fight_num_dat.SLEEP_RATE
	fightNum += fighter.Dizziness * fight_num_dat.DIZZINESS_RATE
	fightNum += fighter.Random * fight_num_dat.RANDOM_RATE
	fightNum += fighter.DisableSkill * fight_num_dat.DISABLE_SKILL_RATE
	fightNum += fighter.Poisoning * fight_num_dat.POISONING_RATE

	fightNum += fighter.CritialHurt * fight_num_dat.CRITIAL_HURT_RATE // 必杀百分比
	fightNum -= 1000
	return fightNum
}

func buddyCoop(db *mdb.Database, coopId int16) {
	var (
		role1, role2 *mdb.PlayerRole
		find         int8
	)

	coopDat := role_dat.GetBuddyCooperation(coopId)

	//确保角色上阵中
	teamPosArray := module.Team.GetFormPosArray(db)
	for _, roleIdPtr := range teamPosArray {
		if *roleIdPtr == coopDat.RoleId1 {
			find++
		}
		if *roleIdPtr == coopDat.RoleId2 {
			find++
		}
	}
	fail.When(find != 2, "目标伙伴不在上阵状态")

	//涉及此次协助的伙伴解除所有组合
	breakBuddyCoop(db, coopDat.RoleId1)
	breakBuddyCoop(db, coopDat.RoleId2)

	db.Select.PlayerRole(func(row *mdb.PlayerRoleRow) {
		if row.RoleId() == coopDat.RoleId1 {
			role1 = row.GoObject()
			return
		}
		if row.RoleId() == coopDat.RoleId2 {
			role2 = row.GoObject()
			return
		}
	})
	fail.When(role1.FriendshipLevel < coopDat.RequireFriendshipLevel || role2.FriendshipLevel < coopDat.RequireFriendshipLevel, "伙伴羁绊等级不足")

	role1.CoopId, role2.CoopId = coopDat.Id, coopDat.Id
	db.Update.PlayerRole(role1)
	db.Update.PlayerRole(role2)
}

//取消伙伴协力关系
func breakBuddyCoop(db *mdb.Database, roleId int8) {
	brokenCoopIds := role_dat.GetBuddyCooperationGroup(roleId)
	var role *mdb.PlayerRole
	for _, coopId := range brokenCoopIds {
		db.Select.PlayerRole(func(row *mdb.PlayerRoleRow) {
			if row.CoopId() == coopId {
				role = row.GoObject()
				role.CoopId = 0
				db.Update.PlayerRole(role)
			}
		})
	}
}
