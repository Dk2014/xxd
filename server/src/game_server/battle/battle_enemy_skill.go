package battle

type triggerSkill struct {
	index   int8
	skillId int
	force   int
}

type funcBossTriggerSkill func(*Fighter) []triggerSkill

func getEnemySkill(f *Fighter) (skill *Skill, skillInfo *SkillInfo) {
	skill1_id := 0
	skill1_force := 0

	skill2_id := 0
	skill2_force := 0

	skills := []triggerSkill{}
	// boss有触发技
	if f.IsBoss {
		skills = getTriggerSkill(f)
		c := len(skills)
		switch c {
		case 0:
			skill1_id, skill1_force = createEnemySkill(f)
			skill2_id, skill2_force = createEnemySkill(f)
		case 1:
			skill1_id = skills[0].skillId
			skill1_force = skills[0].force

			skill2_id, skill2_force = createEnemySkill(f)
		default:
			skill1_id = skills[0].skillId
			skill1_force = skills[0].force

			skill2_id = skills[1].skillId
			skill2_force = skills[1].force
		}

		if skill1_id == 0 {
			skill1_id = DEFAULT_SKILL_ID
		}

		if skill2_id == 0 {
			skill2_id = DEFAULT_SKILL_ID
		}
	} else {
		skill1_id, skill1_force = createEnemySkill(f)
	}

	skill = f.normalAttack
	skillInfo = &SkillInfo{}

	if skill1_id > 0 {
		skill = createRoleSkill(skill1_id)
	}

	skillInfo.SkillId = int(skill1_id)
	skillInfo.SkillId2 = int(skill2_id)
	if f.IsBattlePet && f.PlayerId > 0 {
		skillInfo.SkillTrnLv = int16(skill1_force)
		skillInfo.SkillTrnLv2 = int16(skill2_force)
	} else {
		skillInfo.SkillForce = float64(skill1_force)
		skillInfo.SkillForce2 = float64(skill2_force)
	}
	return
}

func getTriggerSkill(f *Fighter) []triggerSkill {
	trigger := createTriggerSkill(f.RoleId)
	newSkills := trigger(f)
	if len(newSkills) == 0 {
		return newSkills
	}

	skills := []triggerSkill{}

	haveUsed := false
	for _, info := range newSkills {
		for _, idx := range f.triggerIdxRecord {
			if idx == info.index {
				haveUsed = true
				break
			}
		}

		if !haveUsed {
			skills = append(skills, info)
			f.triggerIdxRecord = append(f.triggerIdxRecord, info.index)
			if len(skills) >= 2 {
				break
			}
		}
	}

	return skills
}
