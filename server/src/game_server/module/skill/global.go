package skill

import (
	"game_server/dat/role_dat"
	"game_server/dat/skill_dat"
	"game_server/mdb"
	"game_server/module"
)

func init() {
	module.Skill = SkillMod{}
}

type SkillMod struct {
}

// 初始化玩家伙伴的绝招
func (this SkillMod) InitRoleSkill(state *module.SessionState, roleId int8) {
	var skillId1 int16 = 0
	var skillId2 int16 = 0

	db := state.Database
	roleInfo := role_dat.GetRoleInfo(roleId)

	if roleInfo.SkillID1 != 0 {
		skillInfo1 := skill_dat.GetSkillInfo(roleInfo.SkillID1)
		db.Insert.PlayerSkill(&mdb.PlayerSkill{
			Pid:        db.PlayerId(),
			RoleId:     roleId,
			SkillId:    skillInfo1.Id,
			SkillTrnlv: 1,
		})
		skillId1 = skillInfo1.Id
	}

	if roleInfo.SkillID2 != 0 {
		skillInfo2 := skill_dat.GetSkillInfo(roleInfo.SkillID2)
		db.Insert.PlayerSkill(&mdb.PlayerSkill{
			Pid:        db.PlayerId(),
			RoleId:     roleId,
			SkillId:    skillInfo2.Id,
			SkillTrnlv: 1,
		})
		skillId2 = skillInfo2.Id
	}

	state.Database.Insert.PlayerUseSkill(&mdb.PlayerUseSkill{
		Pid:      db.PlayerId(),
		RoleId:   roleId,
		SkillId1: skillId1,
		SkillId2: 0,
		SkillId3: 0,
		SkillId4: skillId2,
	})
}

func (this SkillMod) UpdateSkill(db *mdb.Database, roleId int8, friendshipLevel, fameLevel, newLevel int16) {
	skillRoleId := roleId
	if role_dat.IsMainRole(roleId) {
		skillRoleId = -2
	}

	haveSkills := map[int16]*mdb.PlayerSkill{}
	canAddSkills := skill_dat.GetSkillByRoleIdWithLevel(skillRoleId, friendshipLevel, fameLevel, newLevel)

	db.Select.PlayerSkill(func(row *mdb.PlayerSkillRow) {
		if row.RoleId() == roleId {
			skill := skill_dat.GetSkillInfo(row.SkillId())
			haveSkills[skill.ParentSkillId] = row.GoObject()
		}
	})

	var playerUseSkill *mdb.PlayerUseSkill
	db.Select.PlayerUseSkill(func(row *mdb.PlayerUseSkillRow) {
		if row.RoleId() == roleId {
			playerUseSkill = row.GoObject()
			row.Break()
		}
	})

	needUpdateUseSkill := false
	session, _ := module.Player.GetPlayerOnline(db.PlayerId())

	var learnedSkill *mdb.PlayerSkill
	var ok bool
	for _, newSkillDat := range canAddSkills {
		// 不存在的绝招有两种情况：一种是直接学习；一种是原来技能的升级

		// 存在技能升级
		learnedSkill, ok = haveSkills[newSkillDat.ParentSkillId]
		if ok {

			learnedSkillDat := skill_dat.GetSkillInfo(learnedSkill.SkillId)

			/**
			 *  For now there is a few conditions to maste skill learning.
			 *  First, only the higher level skill can be learned.
			 *  Second, if we order the table on skill level, no one condition can be decreasing sequence, and there is at least one condition is increasing sequence.
			 */
			if newSkillDat.RequiredLevel <= learnedSkillDat.RequiredLevel && newSkillDat.RequiredFameLevel <= learnedSkillDat.RequiredFameLevel && newSkillDat.RequiredFriendshipLevel <= learnedSkillDat.RequiredFriendshipLevel {
				continue
			}
			// 如果要升级的技能已装配上也更新
			switch learnedSkill.SkillId {
			case playerUseSkill.SkillId1:
				playerUseSkill.SkillId1 = newSkillDat.Id
				needUpdateUseSkill = true

			case playerUseSkill.SkillId2:
				playerUseSkill.SkillId2 = newSkillDat.Id
				needUpdateUseSkill = true

			case playerUseSkill.SkillId3:
				playerUseSkill.SkillId3 = newSkillDat.Id
				needUpdateUseSkill = true

			case playerUseSkill.SkillId4:
				playerUseSkill.SkillId4 = newSkillDat.Id
				needUpdateUseSkill = true
			}

			learnedSkill.SkillId = newSkillDat.Id
			haveSkills[learnedSkillDat.ParentSkillId] = learnedSkill
			db.Update.PlayerSkill(learnedSkill)
		} else {
			learnedSkill = &mdb.PlayerSkill{
				Pid:        db.PlayerId(),
				RoleId:     roleId,
				SkillId:    newSkillDat.Id,
				SkillTrnlv: 1,
			}

			db.Insert.PlayerSkill(learnedSkill)
			haveSkills[newSkillDat.ParentSkillId] = learnedSkill
		}

		//通知玩家
		if session != nil {
			module.Notify.SendSkillAdd(session, roleId, newSkillDat.Id)
		}
	}

	if needUpdateUseSkill {
		db.Update.PlayerUseSkill(playerUseSkill)
	}
}
