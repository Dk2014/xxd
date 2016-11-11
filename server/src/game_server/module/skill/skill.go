package skill

import (
	"core/fail"
	"core/net"
	"core/time"
	"game_server/api/protocol/skill_api"
	"game_server/dat/item_dat"
	"game_server/dat/player_dat"
	"game_server/dat/quest_dat"
	"game_server/dat/role_dat"
	"game_server/dat/skill_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/tlog"
	"game_server/xdlog"
)

func GetAllSkillsInfo(session *net.Session, out *skill_api.GetAllSkillsInfo_Out) {
	state := module.State(session)

	playerInfo := state.Database.Lookup.PlayerInfo(state.PlayerId)
	out.FlushTime = playerInfo.LastSkillFlush

	out.Roles = make([]skill_api.GetAllSkillsInfo_Out_Roles, 0, role_dat.MAX_ROLE_NUM)
	out.Skills = make([]skill_api.GetAllSkillsInfo_Out_Skills, 0, skill_dat.MAX_SKILL_NUM)

	state.Database.Select.PlayerSkill(func(row *mdb.PlayerSkillRow) {
		out.Skills = append(out.Skills,
			skill_api.GetAllSkillsInfo_Out_Skills{
				RoleId:        row.RoleId(),
				SkillId:       row.SkillId(),
				TrainingLevel: int16(row.SkillTrnlv()),
			})
	})

	state.Database.Select.PlayerUseSkill(func(row *mdb.PlayerUseSkillRow) {
		var status int8
		state.Database.Select.PlayerRole(func(pr_row *mdb.PlayerRoleRow) {
			if pr_row.RoleId() == row.RoleId() {
				status = int8(pr_row.Status())
			}
		})
		out.Roles = append(out.Roles,
			skill_api.GetAllSkillsInfo_Out_Roles{
				RoleId:   row.RoleId(),
				Status:   status,
				SkillId1: row.SkillId1(),
				SkillId2: row.SkillId2(),
				SkillId3: row.SkillId3(),
				SkillId4: row.SkillId4(),
			})
	})
}

func changeSkill(session *net.Session, roleId int8, pos int8, skillId int16) {
	state := module.State(session)
	tlog.PlayerSystemModuleFlowLog(state.Database, tlog.SMT_SKILL)

	//主角可以装备1-4共4个位置，伙伴可以装备1和4共2个位置
	fail.When(pos < 1 || pos > 4 || (pos != 1 && pos != 4 && !role_dat.IsMainRole(roleId)), "wrong order number")

	if skillId != skill_dat.SKILL_IS_NULL {
		//判定绝招是否可装备
		//if !role_dat.IsMainRole(roleId) {
		//	fail.When(roleId != skill.RoleId, "wrong role id")
		//} else {
		//	fail.When(skill_dat.SKILL_MAIN_ROLE_ID != skill.RoleId, "wrong role id")
		//}
		//the previous logic with the detection of whether this role learned this skill should be equally to the below one -- DHZ 20150107

		//角色是否拥有此绝招
		hasThisSkill := false
		state.Database.Select.PlayerSkill(func(row *mdb.PlayerSkillRow) {
			if row.RoleId() == roleId && row.SkillId() == skillId {
				hasThisSkill = true
				row.Break()
			}
		})
		fail.When(!hasThisSkill, "this role hasn't learned this skill yet")

		//获取技能槽可装备技能类型
		skillSlotType := skill_dat.GetSkillTypeByPos(pos)
		//获取技能信息
		skill := skill_dat.GetSkillInfo(skillId)

		fail.When(skillSlotType != skill.ChildKind, "技能槽不可装备该类型技能")
	}

	var data *mdb.PlayerUseSkill = nil
	state.Database.Select.PlayerUseSkill(func(row *mdb.PlayerUseSkillRow) {
		if row.RoleId() == roleId {
			data = row.GoObject()
			row.Break()
		}
	})

	fail.When(data == nil, "角色技能未初始化")

	switch pos {
	case skill_dat.POS_SKILL_1:
		data.SkillId1 = skillId
	case skill_dat.POS_SKILL_2:
		data.SkillId2 = skillId
	case skill_dat.POS_SKILL_3:
		data.SkillId3 = skillId
	case skill_dat.POS_SKILL_4:
		data.SkillId4 = skillId
	}

	state.Database.Update.PlayerUseSkill(data)
}

func studyskill(session *net.Session, roleId int8, itemId int16, xdEventType int32) (result skill_api.CheatResult) {
	state := module.State(session)
	item_info := item_dat.GetItem(itemId)
	var skillRoleId int8
	if roleId == 1 || roleId == 2 {
		skillRoleId = -2
	} else {
		skillRoleId = roleId
	}
	//对比是否为武功秘籍类别
	if item_info.TypeId != item_dat.TYPE_CHEAT {
		return skill_api.CHEAT_RESULT_NOT_CHEAT_TYPE
	}
	skill := skill_dat.GetSkillByCheatId(itemId)
	//查看对应itemid的秘籍是否已配置
	if skill == nil {
		return skill_api.CHEAT_RESULT_NO_DEPAND_SKILL
	} else if skill.RoleId != skillRoleId { //看角色是否可学习绝技
		return skill_api.CHEAT_RESULT_SKILL_NOT_MATCH_ROLE
	}
	parentSkill := skill.ParentSkillId
	var has_study = false
	var study_before_skill = false
	var has_role = false
	var level_enough = false
	//看角色是否存在,如果存在判断等级是否打到学习要求
	state.Database.Select.PlayerRole(func(row *mdb.PlayerRoleRow) {
		if row.RoleId() == roleId {
			has_role = true
			if row.Level() >= skill.RequiredLevel {
				level_enough = true
			}
			row.Break()
		}
	})
	//如果角色不存在或达不到等级要求，直接return
	if !level_enough {
		return skill_api.CHEAT_RESULT_LEVEL_NOT_REACHED
	}
	if !has_role {
		return skill_api.CHEAT_RESULT_ROLE_DOES_NOT_EXISTS
	}
	state.Database.Select.PlayerSkill(func(row *mdb.PlayerSkillRow) {
		//看是否已学习秘籍
		if row.SkillId() == skill.Id {
			has_study = true
			row.Break()
		} else if parentSkill > 0 && parentSkill != row.SkillId() { //如果有父技能并且已学的技能不含父技能，而是之后的子技能，则对比学习等级，如果已学的等级大于想学的，则不允许，返回错误
			skill_info := skill_dat.GetSkillInfo(row.SkillId())
			if skill_info.ParentSkillId == parentSkill && skill_info.RequiredLevel > skill.RequiredLevel {
				study_before_skill = true
				row.Break()
			}
		} else if parentSkill > 0 && parentSkill == row.SkillId() { //已学习父技能，可以直接学习秘籍,直接break
			row.Break()
		}
	})
	if has_study {
		return skill_api.CHEAT_RESULT_ALREADY_STUDY
	}
	if study_before_skill {
		return skill_api.CHEAT_RESULT_CAN_NOT_STUDY_BEFORE
	}
	//删除秘籍
	module.Item.DelItemByItemId(state.Database, itemId, 1, tlog.IFR_CHEAT, xdEventType)
	//学习技能
	state.Database.Insert.PlayerSkill(&mdb.PlayerSkill{
		Pid:        state.PlayerId,
		RoleId:     roleId,
		SkillId:    skill.Id,
		SkillTrnlv: 1,
	})
	//通知玩家习得技能
	if session != nil {
		module.Notify.SendSkillAdd(session, roleId, skill.Id)
	}
	return skill_api.CHEAT_RESULT_SUCCESS
}

func trainSkill(session *net.Session, roleId int8, skillId int16) {
	state := module.State(session)
	db := state.Database

	//找出玩家的指定技能
	var playerOneSkill *mdb.PlayerSkill
	db.Select.PlayerSkill(func(row *mdb.PlayerSkillRow) {
		if row.SkillId() == skillId && row.RoleId() == roleId {
			playerOneSkill = row.GoObject()
			row.Break()
		}
	})
	fail.When(playerOneSkill == nil, "try to train a player skill which doesn't exist")

	//检查技能最大训练等级
	fail.When(playerOneSkill.SkillTrnlv >= skill_dat.MAX_SKILL_LEVEL, "try to train a skill already is highest level")

	//确保技能等级不高于角色等级
	var playerRole *mdb.PlayerRole
	if role_dat.IsMainRole(roleId) {
		playerRole = module.Role.GetMainRole(db)
	} else {
		playerRole = module.Role.GetBuddyRole(db, roleId)
	}
	fail.When(int16(playerOneSkill.SkillTrnlv) >= playerRole.Level, "try to train a skill which level will come over role level")

	//扣钱
	cost := skill_dat.GetSkillTrainingCost(int16(playerOneSkill.SkillTrnlv))
	if module.Player.CheckMoney(state, cost, player_dat.COINS) {
		module.Player.DecMoney(db, state.MoneyState, cost, player_dat.COINS, tlog.MFR_ROLE_SKILL_LEVELUP, xdlog.ET_ROLE_SKILL)

		playerOneSkill.SkillTrnlv++
		db.Update.PlayerSkill(playerOneSkill)
		module.Quest.RefreshDailyQuest(state, quest_dat.DAILY_QUEST_CLASS_SKILL_TRAINING)
	}
}

func flushSkill(session *net.Session, roleId int8) int64 {
	state := module.State(session)
	db := state.Database

	playerInfo := db.Lookup.PlayerInfo(db.PlayerId())
	nowTime := time.GetNowTime()
	fail.When(nowTime-playerInfo.LastSkillFlush < skill_dat.FLUSH_SKILL_CD, "the skill flushing is in CD time")

	has_flushed_anything := false

	db.Select.PlayerSkill(func(row *mdb.PlayerSkillRow) {
		if row.RoleId() == roleId && row.SkillTrnlv() > 1 {
			has_flushed_anything = true
			obj := row.GoObject()
			price := float64(skill_dat.GetSkillTrainingTotalCost(int16(obj.SkillTrnlv) - int16(1)))
			price *= skill_dat.FLUSH_SKILL_BACK_PERCENT
			module.Player.IncMoney(db, state.MoneyState, int64(price), player_dat.COINS, 0 /*TODO*/, xdlog.ET_ROLE_SKILL, "")
			obj.SkillTrnlv = 1
			db.Update.PlayerSkill(obj)
		}
	})

	//这里需要重新获取一次 playerInfo 否则会覆盖 IncMoney 增加的铜钱
	playerInfo = db.Lookup.PlayerInfo(db.PlayerId())
	if has_flushed_anything {
		playerInfo.LastSkillFlush = nowTime
		db.Update.PlayerInfo(playerInfo)
	}

	return playerInfo.LastSkillFlush
}
