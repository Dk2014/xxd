package battle

import (
	"math/rand"
)

func mainRoleAI_2(f *Fighter) (*Skill, *SkillInfo) {
	var skill = f.normalAttack
	var skillInfo = normalSkillInfo
	// 灵宠关卡中主角只能使用特殊技
	if f.battle.BatType == BT_PET_LEVEL {
		skill = createRoleSkill(PET_LEVEL_MAIN_ROLE_SKILL_ID)
		return skill, skillInfo
	}
	if !f.playerInfo.Auto {
		if f.useSkillIndex >= 0 && f.useSkillIndex < len(f.Skills) && f.Skills[f.useSkillIndex] != nil {
			var skill2 = f.Skills[f.useSkillIndex]
			if f.Power >= skill2.DecPower {
				skill = skill2
				skillInfo = f.SkillInfos[f.useSkillIndex]
			}
			f.useSkillIndex = 0
		}
	} else {
		return pickSkillForPlayerRole(f, f.Skills, f.SkillInfos)
	}
	return skill, skillInfo
}

func buddyAI(f *Fighter) (*Skill, *SkillInfo) {
	var skill = f.normalAttack
	var skillInfo = normalSkillInfo

	// 召唤的灵宠释放技能
	if f.IsBattlePet {
		return getEnemySkill(f)
	}

	//自动模式暂时默认使用第一个基本技能，没有第一个技能则使用普通攻击
	skillIndex := f.useSkillIndex
	if f.playerInfo.Auto {
		skillIndex = rand.Intn(2)
		if f.Skills[skillIndex] != nil && f.SkillInfos[skillIndex].MaxReleaseNum > 0 && f.SkillInfos[skillIndex].ReleaseNum <= 0 {
			skillIndex = 0
		}

		return pickSkillForPlayerRole(f, f.Skills, f.SkillInfos)
	}

	//伙伴
	if f.Kind == FK_BUDDY {
		for i := 0; i < len(f.Skills); i++ {
			if f.Skills[i] != nil && i == skillIndex &&
				(f.SkillInfos[i].MaxReleaseNum <= 0 || f.SkillInfos[i].ReleaseNum > 0) {
				//如果伙伴技能有限制最大使用次数，则检查当前使用次数
				skill = f.Skills[skillIndex]
				skillInfo = f.SkillInfos[skillIndex]
			}
		}
	}
	return skill, skillInfo
}

// 精气满就放绝招的AI，伙伴用（废弃）
func normalAI(f *Fighter) (*Skill, *SkillInfo) {
	var skill = f.normalAttack
	var skillInfo = normalSkillInfo

	// 召唤的灵宠释放技能
	if f.IsBattlePet {
		return getEnemySkill(f)
	}

	// 有绝招的怪和伙伴才需要
	if f.useSkillIndex >= 0 {
		var selectSkill = f.Skills[f.useSkillIndex]
		var selectSkillInfo = f.SkillInfos[f.useSkillIndex]

		// 恢复完才能使用
		if selectSkillInfo.Rhythm > selectSkillInfo.RecoverRhythm {
			skill = selectSkill
			skillInfo = selectSkillInfo
		}

		selectSkillInfo.Rhythm += 1

		// 用完重新开始恢复
		if selectSkillInfo.Rhythm > selectSkillInfo.UseRhythm {
			selectSkillInfo.Rhythm = 1
		}

		// 伙伴没用到的绝招也需要恢复
		if f.Kind == FK_BUDDY {
			for i, s := range f.SkillInfos {
				// 跳过正在用的绝招
				if i == f.useSkillIndex {
					continue
				}

				if s == nil {
					break
				}

				// 还没恢复满的恢复一次
				if s.Rhythm <= s.RecoverRhythm {
					s.Rhythm += 1
				}
			}
		}
	}

	return skill, skillInfo
}

func pickSkillForPlayerRole(fighter *Fighter, skills [4]*Skill, skillInfos []*SkillInfo) (*Skill, *SkillInfo) {
	var (
		cureIndex               int = -1 //治疗技能
		aoeIndex                int = -1 //群体攻击
		suppotrIndex            int = -1 //辅助技能
		attackIndex             int = -1 //单体攻击
		skill, oldSkill         *Skill
		skillInfo, oldSkillInfo *SkillInfo
		aoeTargetNum            int = 1
	)
	for index := 0; index < len(skills); index++ {
		skill = skills[index]
		if skill == nil {
			continue
		}
		skillInfo = skillInfos[index] //因为 skillInfo 可能不足4个
		if fighter.Kind == FK_PLAYER {
			if fighter.Power < skill.DecPower {
				continue
			}
		} else if fighter.Kind == FK_BUDDY {
			if skillInfo.MaxReleaseNum > 0 && skillInfo.ReleaseNum <= 0 {
				continue
			}
		}
		//治疗技能
		switch skill.ChildType {
		case SKILL_KIND_CURE: //治疗
			if cureIndex >= 0 {
				oldSkill = skills[cureIndex]
				oldSkillInfo = skillInfos[cureIndex]
			} else {
				oldSkill = nil
				oldSkillInfo = nil
			}
			if oldSkill == nil {
				cureIndex = index
			} else {
				if fighter.Kind == FK_PLAYER { //主角技能看精气，有限使用消耗精气多的技能
					if oldSkill.DecPower < skill.DecPower {
						cureIndex = index
					}
				} else if fighter.Kind == FK_BUDDY { //伙伴技能看释放次数，伙伴只有2个技能，有限使用有释放技能限制的
					if oldSkillInfo.MaxReleaseNum < skillInfo.MaxReleaseNum {
						cureIndex = index
					}
				}
			}
		case SKILL_KIND_DEFEND, SKILL_KIND_SUPPORT: //辅助：
			if oldSkill == nil {
				suppotrIndex = index
			} else {
				if fighter.Kind == FK_PLAYER {
					if oldSkill.DecPower < skill.DecPower {
						suppotrIndex = index
					}
				} else if fighter.Kind == FK_BUDDY {
					if oldSkillInfo.MaxReleaseNum < skillInfo.MaxReleaseNum {
						suppotrIndex = index
					}
				}
			}
		case SKILL_KIND_ATTACK, SKILL_KIND_REDUCE_SUNDER: //进攻
			if skill.AttackMode == SKILL_ATTACK_MODE_AOE {
				if skill._FindTargets != nil {
					newAoeTargetNum := len(skill._FindTargets(fighter))
					if newAoeTargetNum > aoeTargetNum {
						aoeTargetNum = newAoeTargetNum
						aoeIndex = index
					}
				}
			} else {
				if oldSkill == nil {
					attackIndex = index
				} else {
					if fighter.Kind == FK_PLAYER {
						if oldSkill.DecPower < skill.DecPower {
							attackIndex = index
						}
					} else if fighter.Kind == FK_BUDDY {
						if oldSkillInfo.MaxReleaseNum < skillInfo.MaxReleaseNum {
							attackIndex = index
						}
					}
				}

			}
		}
	}
	if cureIndex >= 0 {
		if fighter.side.AnyRoleBadlyHurt() {
			return skills[cureIndex], skillInfos[cureIndex]
		}
	}
	if aoeIndex >= 0 {
		return skills[aoeIndex], skillInfos[aoeIndex]
	}
	if suppotrIndex >= 0 {
		if rand.Intn(2) == 0 {
			return skills[suppotrIndex], skillInfos[suppotrIndex]
		}
	}
	if attackIndex >= 0 {
		return skills[attackIndex], skillInfos[attackIndex]
	}
	return fighter.normalAttack, normalSkillInfo
}

func enemyAI(f *Fighter) (*Skill, *SkillInfo) {
	return getEnemySkill(f)
}
