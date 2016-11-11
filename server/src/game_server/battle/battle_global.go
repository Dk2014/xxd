package battle

/* FighterAttribute 用户保存玩家进入关卡时所有角色的信息
   在关卡里操作背包等会影响这些数据
*/
type FighterAttribute struct {
	SunderValue    int  // 护甲
	Power          int  // 精气
	Health         int  // 血量
	UsedGhostSkill bool //是否是放过魂侍
	//GhostPower     int     // 魂力

	SunderMaxValue int //最大护甲值
	MaxHealth      int //最大血量
}

type ExportFighterInfo struct {
	// map[角色ID][]信息
	Buffs    map[int][]Buff
	Skills   map[int][]SkillInfo
	Fighters map[int]FighterAttribute

	//map[player_id]信息
	GhostInfo map[int64]PlayerGhostInfo
}

// 导出还未失效并作用在自己身上的Buff
func (this *Fighter) exportBuff() []Buff {
	//死亡角色不需要导出 buf
	if this.Health <= MIN_HEALTH {
		return nil
	}

	buffs := make([]Buff, 0)

	for buff := this.Buffs; buff != nil; buff = buff.Next {
		if buff.Keep == 0 || buff.IsBadBuff() || this.Kind != buff.Owner.Kind { // 保证是作用在fighter自己身上
			continue
		}
		showKeep := buff.Keep
		if showKeep < 0 {
			showKeep = 0
		}
		buffs = append(buffs, Buff{
			Skill:       buff.Skill,
			Value:       buff.Value,
			MaxOverride: buff.MaxOverride,
			Mode:        buff.Mode,
			Keep:        buff.Keep,
			ShowKeep:    buff.ShowKeep,
		})
	}
	return buffs
}

// 导出伙伴绝招信息
func (this *Fighter) exportSkill() []SkillInfo {
	skills := make([]SkillInfo, 0, len(this.SkillInfos))

	for _, info := range this.SkillInfos {
		if info.MaxReleaseNum > 0 {
			skills = append(skills, *info)
		}
	}

	return skills
}

func (this *Fighter) exportAttribute() FighterAttribute {
	health := this.Health
	if health <= MIN_HEALTH {
		health = MIN_HEALTH
	}
	return FighterAttribute{
		SunderValue:    this.sunderValue,
		Power:          this.Power,
		Health:         health,
		UsedGhostSkill: this.UsedGhostSkill,

		//彩虹关卡需要用到下面两个值。只需导出，不需要恢复
		SunderMaxValue: this.SunderMaxValue,
		MaxHealth:      this.MaxHealth,
	}
}

// 导出攻方信息
func (b *BattleState) ExportAttackersInfo() *ExportFighterInfo {
	exportInfo := &ExportFighterInfo{
		Buffs:     make(map[int][]Buff),
		Skills:    make(map[int][]SkillInfo),
		Fighters:  make(map[int]FighterAttribute),
		GhostInfo: make(map[int64]PlayerGhostInfo),
	}

	for _, f := range b.Attackers.Fighters {
		if f == nil || f.IsBattlePet { // 不导出灵宠数据
			continue
		}

		exportInfo.Buffs[f.RoleId] = f.exportBuff()
		exportInfo.Skills[f.RoleId] = f.exportSkill()
		exportInfo.Fighters[f.RoleId] = f.exportAttribute()
	}

	for _, player := range b.Attackers.Players {
		exportInfo.GhostInfo[player.PlayerId] = player.GhostInfo
	}

	return exportInfo
}

func (b *BattleState) ImportAttackersInfo(exportInfo *ExportFighterInfo) {
	for _, f := range b.Attackers.Fighters {
		if f == nil {
			continue
		}

		if fighterAttr, ok := exportInfo.Fighters[f.RoleId]; ok {
			f.applyFighterAttribute(&fighterAttr)
			// 在关卡内操作影响到角色上线血量，导致战斗后的血量可能大于上线血量
			if f.MaxHealth < f.Health {
				f.Health = f.MaxHealth
			}
		}

		if buffs, ok := exportInfo.Buffs[f.RoleId]; ok {
			f.applyBuffs(buffs)
		}

		if skills, ok := exportInfo.Skills[f.RoleId]; ok {
			f.applyBuddySkills(skills)
		}

		// 魂侍关卡初始满魂力
		if b.BatType == BT_GHOST_LEVEL {
			f.GhostInfo.GhostPower = FULL_GHOST_POWER
			f.ghostPowerIncMultiple = 2
		}
	}

	for _, player := range b.Attackers.Players {
		if ghostInfo, ok := exportInfo.GhostInfo[player.PlayerId]; ok {
			player.GhostInfo.GhostPower = ghostInfo.GhostPower
		}
	}
}

// 恢复buff到战场角色
func (this *Fighter) applyBuffs(buffs []Buff) {
	for _, buff := range buffs {
		this.addBuff(this, buff.Mode, buff.Value, buff.Keep, buff.Skill, buff.MaxOverride, buff.Uncleanable)
	}
}

// 恢复伙伴技能状态
func (this *Fighter) applyBuddySkills(skills []SkillInfo) {
	if this.Kind != FK_BUDDY {
		return
	}

	for _, skill := range this.SkillInfos {
		for i := 0; i < len(skills); i++ {
			if skill.SkillId == skills[i].SkillId {
				skill.Rhythm = skills[i].Rhythm
				skill.UseRhythm = skills[i].UseRhythm
				skill.RecoverRhythm = skills[i].RecoverRhythm
				skill.MaxReleaseNum = skills[i].MaxReleaseNum
				skill.ReleaseNum = skills[i].ReleaseNum
			}
		}
	}
}

// 恢复战场角色一些属性
func (this *Fighter) applyFighterAttribute(fighterAttr *FighterAttribute) {
	/*
	 角色恢复以下属性
	   魂侍魂侍是否使用过
	   血量
	   护甲
	   精气
	*/
	this.UsedGhostSkill = fighterAttr.UsedGhostSkill
	//this.GhostPower = fighterAttr.GhostPower
	this.Power = fighterAttr.Power
	if this.Power > this.MaxPower {
		this.Power = this.MaxPower
	}

	if fighterAttr.Health <= MIN_HEALTH {
		this.Health = 1
	} else {
		this.Health = fighterAttr.Health
		if this.Health > this.MaxHealth {
			this.Health = this.MaxHealth
		}
	}
	if fighterAttr.SunderValue < 1 {
		this.sunderValue = this.SunderMaxValue
	} else {
		this.sunderValue = fighterAttr.SunderValue
		if this.sunderValue > this.SunderMaxValue {
			this.sunderValue = this.SunderMaxValue
		}
	}
}
