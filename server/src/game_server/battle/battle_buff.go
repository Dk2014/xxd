package battle

import "math"

// buff信息
type Buff struct {
	Prev        *Buff    // 前一个buff
	Next        *Buff    // 下一个buff
	Owner       *Fighter // 被施加buff的对象
	Mode        int      // buff类型
	Keep        int      // buff持续回合数，每回合减1
	ShowKeep    int      // buff持续回合数，每回合减1
	Skill       int      // buff所属的绝招ID
	Value       float64  // buff产生的作用值
	MaxOverride int      // buff最大叠加数量
	OverrideNum int      // buff已叠加数量
	cleanNum    int      // 清除buff数量
	Uncleanable bool     //不可清除特性

	// updateRound 用来避免一个buff在一个回合被多次刷新次数
	// 初始化的时候需要取当前回合输减一，因为产生的回合可能会需要刷新，其他回合只需要刷新一次。
	updateRound int //更新回合
}

// buff作用时会用到原值百分比计算的属性
type rawdata struct {
	Speed        float64
	Attack       float64
	Defend       float64
	Block        float64
	DodgeLevel   float64
	CritialLevel float64
	BlockLevel   float64
	HitLevel     float64
}

// buff起作用
func (buff *Buff) Apply(f *Fighter) {
	buff.Value = math.Ceil(buff.Value)

	switch buff.Mode {
	case BUFF_DIZZINESS:
		f.dizziness += 1
	case BUFF_POISONING:
		f.DecHealth += int(buff.Value)
	case BUFF_KEEPER_REDUCE_HURT:
		fallthrough
	case BUFF_REDUCE_HURT:
		//减伤最多90%
		if f.reduceHurt+buff.Value > 90 {
			buff.Value = (90 - f.reduceHurt)
			f.reduceHurt = 90
		} else {
			f.reduceHurt += buff.Value
		}

	case BUFF_ABSORB_HURT:
		f.absorbHurt = buff.Value

	case BUFF_DIRECT_REDUCE_HURT:
		f.directReduceHurt = buff.Value

	case BUFF_BLOCK:
		if f.Block+buff.Value < 0 {
			buff.Value = -f.Block
			f.Block = 0
		} else {
			f.Block += buff.Value
		}
	case BUFF_RANDOM:
		f.random += 1
	case BUFF_SLEEP:
		f.sleepState = 1

	case BUFF_HIT_LEVEL:
		f.HitLevel += buff.Value
	case BUFF_SPEED:
		f.Speed += buff.Value
		// 角色还没出手被变更速度，需要重新排序
		if !f.fighted && f.battle != nil {
			f.battle.needResort = true
		}
	case BUFF_ATTACK:
		if f.Attack+buff.Value < 0 {
			buff.Value = -f.Attack
			f.Attack = 0
		} else {
			f.Attack += buff.Value
		}
	case BUFF_DEFEND:
		if f.Defend+buff.Value < 0 {
			buff.Value = -f.Defend
			f.Defend = 0
		} else {
			f.Defend += buff.Value
		}
	case BUFF_BLOCK_LEVEL:
		f.BlockLevel += buff.Value
	case BUFF_DODGE_LEVEL:
		f.DodgeLevel += buff.Value
	case BUFF_CRITIAL_LEVEL:
		f.CritialLevel += buff.Value
	case BUFF_HURT_ADD:
		if f.hurtAdd+buff.Value < 0 {
			buff.Value = -f.hurtAdd
			f.hurtAdd = 0
		} else {
			f.hurtAdd += buff.Value
		}
	case BUFF_ATTRACT_FIRE:
		f.attractFire += 1
	case BUFF_DESTROY_LEVEL:
		f.DestroyLevel += buff.Value
	case BUFF_TENACITY_LEVEL:
		f.TenacityLevel += buff.Value
	case BUFF_SLEEP_LEVEL:
		f.SleepLevel += buff.Value
	case BUFF_DIZZINESS_LEVEL:
		f.DizzinessLevel += buff.Value
	case BUFF_RANDOM_LEVEL:
		f.RandomLevel += buff.Value
	case BUFF_DISABLE_SKILL_LEVEL:
		f.DisableSkillLevel += buff.Value
	case BUFF_POISONING_LEVEL:
		f.PoisoningLevel += buff.Value
	case BUFF_DISABLE_SKILL:
		f.disableSkill += 1
	case BUFF_GHOST_SHIELD:
		f.ghostShieldState += 1

	case BUFF_TAKE_SUNDER:
		//在使用 f.SunderAttack 的地方避免负值
		f.SunderAttack += int(buff.Value)
	//一下四个技能可能把角色的属性设置为负数
	case BUFF_DOGE:
		f.Dodge += buff.Value
	case BUFF_HIT:
		f.Hit += buff.Value
	case BUFF_CRITIAL:
		f.Critial += buff.Value
	case BUFF_TENACITY:
		f.Tenacity += buff.Value

	case BUFF_DEFEND_PERSENT:
		f.defenceRate += int(buff.Value)

	case BUFF_SUNDER_STATE:
		f.sunderValue = -2

	//
	// 以下是不需要恢复的
	//
	case BUFF_POWER:
		if buff.Value > 0 {
			f.Power += int(buff.Value)

			if f.Power > f.MaxPower {
				f.Power = f.MaxPower
			}
		} else {
			if f.Power+int(buff.Value) < 0 {
				buff.Value = float64(f.Power)
				f.Power = 0
			}
		}
	case BUFF_HEALTH:
		f.Health += int(buff.Value)
		if f.Health > f.MaxHealth {
			f.Health = f.MaxHealth
		}
	case BUFF_MAX_HEALTH:
		f.Health += int(buff.Value)
		f.MaxHealth += int(buff.Value)
		if f.Health > f.MaxHealth {
			f.Health = f.MaxHealth
		}
	case BUFF_SUNDER:
		if f.sunderValue <= 0 {
			buff.Value = 0
			return
		}
		f.sunderValue += int(buff.Value)
		if f.sunderValue > f.SunderMaxValue {
			f.sunderValue = f.SunderMaxValue
		}
	case BUFF_CLEAN_BAD:
		for b := f.Buffs; b != nil; b = b.Next {
			if b.IsBadBuff() {
				b.Revert(f)
				f.Buffs.remove(b)
			}
		}
	case BUFF_CLEAN_GOOD:
		for b := f.Buffs; b != nil; b = b.Next {
			if !b.IsBadBuff() {
				b.Revert(f)
				f.Buffs.remove(b)
			}
		}
	case BUFF_GHOST_POWER:
		if len(f.Ghosts) > 0 {
			f.GhostInfo.GhostPower += int(buff.Value)
			if f.GhostInfo.GhostPower > FULL_GHOST_POWER {
				f.GhostInfo.GhostPower = FULL_GHOST_POWER
			}
		}
	case BUFF_PET_LIVE_ROUND:
		if f.IsBattlePet {
			f.BattlePetLiveRound += int(buff.Value)
		}
	case BUFF_BUDDY_SKILL:
		for i := 0; i < len(f.SkillInfos); i++ {
			if f.SkillInfos[i] != nil && f.SkillInfos[i].MaxReleaseNum > 0 {
				skillInfo := f.SkillInfos[i]

				if skillInfo.MaxReleaseNum > 0 {
					skillInfo.ReleaseNum += int(buff.Value)
					if skillInfo.ReleaseNum < 0 {
						skillInfo.ReleaseNum = 0
					}
				}
				//skillInfo.MaxReleaseNum += int(buff.Value)
			}
		}
	case BUFF_MAKE_POWER_FULL:
		f.Power = f.MaxPower
	case BUFF_RECOVER_BUDDY_SKILL:
		for i := 0; i < len(f.SkillInfos); i++ {
			if f.SkillInfos[i] != nil && f.SkillInfos[i].MaxReleaseNum > 0 {
				skillinfo := f.SkillInfos[i]
				skillinfo.ReleaseNum = skillinfo.MaxReleaseNum
			}
		}

	case BUFF_CLEAR_ABSORB_HURT:
		for b := f.Buffs; b != nil; b = b.Next {
			if b.Mode == BUFF_ABSORB_HURT {
				b.Revert(f)
				f.Buffs.remove(b)
			}
		}
	case BUFF_HEALTH_PERCENT:
		f.Health += int(float64(f.MaxHealth) * buff.Value / float64(100))
		if f.Health > f.MaxHealth {
			f.Health = f.MaxHealth
		}
	case BUFF_ALL_RESIST:
		f.SleepLevel += buff.Value        // 睡眠抗性
		f.DizzinessLevel += buff.Value    // 眩晕抗性
		f.RandomLevel += buff.Value       // 混乱抗性
		f.DisableSkillLevel += buff.Value // 封魔抗性
		f.PoisoningLevel += buff.Value    // 中毒抗性
	case BUFF_REBOTH_HEALTH:
		f.Health = int(buff.Value)
		if f.Health > f.MaxHealth {
			f.Health = f.MaxHealth
		}
		f.side.live++
		f.Buffs.CleanAllBuff(f)
		f.sunderValue = f.SunderMaxValue
	case BUFF_REBOTH_HEALTH_PERCENT:
		f.Health = 0
		f.Health += int(buff.Value / 100 * float64(f.MaxHealth))
		if f.Health > f.MaxHealth {
			f.Health = f.MaxHealth
		}
		f.side.live++
		f.Buffs.CleanAllBuff(f)
		f.sunderValue = f.SunderMaxValue
	}
}

func (buff *Buff) IsBadBuff() bool {
	switch buff.Mode {
	case BUFF_POWER:
		return buff.Value < 0 // 精气
	case BUFF_SPEED:
		return buff.Value < 0 // 速度
	case BUFF_ATTACK:
		return buff.Value < 0 // 攻击
	case BUFF_DEFEND:
		return buff.Value < 0 // 防御
	case BUFF_HEALTH:
		return buff.Value < 0 // 生命
	case BUFF_SLEEP:
		return true
	case BUFF_DIZZINESS:
		return true // 眩晕
	case BUFF_POISONING:
		return true // 中毒
	case BUFF_CLEAN_BAD:
		return false // 清除负面buff
	case BUFF_CLEAN_GOOD:
		return true // 清除增益buff
	case BUFF_REDUCE_HURT:
		return false // 伤害减免
	case BUFF_DIRECT_REDUCE_HURT:
		return false // 直接免伤
	case BUFF_ABSORB_HURT:
		return false // 伤害吸收
	case BUFF_RANDOM:
		return true // 混乱
	case BUFF_BLOCK:
		return buff.Value < 0 // 格挡概率
	case BUFF_BLOCK_LEVEL:
		return buff.Value < 0 // 格挡概率等级
	case BUFF_DODGE_LEVEL:
		return buff.Value < 0 // 闪避概率等级
	case BUFF_CRITIAL_LEVEL:
		return buff.Value < 0 // 暴击等级
	case BUFF_HIT_LEVEL:
		return buff.Value < 0 // 命中等级
	case BUFF_HURT_ADD:
		return buff.Value < 0 // 伤害加值（百分数）
	case BUFF_MAX_HEALTH:
		return buff.Value < 0 // 最大生命
	case BUFF_KEEPER_REDUCE_HURT:
		return buff.Value < 0 // 守卫者免伤
	case BUFF_ATTRACT_FIRE:
		return false // 吸引火力
	case BUFF_DESTROY_LEVEL:
		return buff.Value < 0 // 破击
	case BUFF_TENACITY_LEVEL:
		return buff.Value < 0 // 韧性
	case BUFF_SUNDER:
		return buff.Value < 0 // 护甲
	case BUFF_SLEEP_LEVEL:
		return buff.Value < 0
	case BUFF_DIZZINESS_LEVEL:
		return buff.Value < 0
	case BUFF_RANDOM_LEVEL:
		return buff.Value < 0
	case BUFF_DISABLE_SKILL_LEVEL:
		return buff.Value < 0
	case BUFF_POISONING_LEVEL:
		return buff.Value < 0
	case BUFF_DISABLE_SKILL:
		return true // 禁用绝招
	case BUFF_GHOST_SHIELD:
		return false
	case BUFF_DOGE:
		return buff.Value < 0
	case BUFF_HIT:
		return buff.Value < 0
	case BUFF_CRITIAL:
		return buff.Value < 0
	case BUFF_TENACITY:
		return buff.Value < 0
	case BUFF_TAKE_SUNDER:
		return buff.Value < 0
	case BUFF_DEFEND_PERSENT:
		return buff.Value < 0
	case BUFF_HEALTH_PERCENT:
		return buff.Value < 0 && buff.Value > 100
	case BUFF_ALL_RESIST:
		return buff.Value < 0
	case BUFF_REBOTH_HEALTH:
		return buff.Value < 0
	case BUFF_REBOTH_HEALTH_PERCENT:
		return buff.Value < 0 && buff.Value > 100
	}
	return false
}

// 撤销buff作用
func (buff *Buff) Revert(f *Fighter) {
	switch buff.Mode {
	case BUFF_DIZZINESS:
		f.dizziness -= 1
	case BUFF_POISONING:
		f.DecHealth -= int(buff.Value)
	case BUFF_BLOCK:
		f.Block -= buff.Value
	case BUFF_RANDOM:
		f.random -= 1
	case BUFF_SLEEP:
		f.sleepState = 0
	case BUFF_HIT_LEVEL:
		f.HitLevel -= buff.Value
	case BUFF_HURT_ADD:
		f.hurtAdd -= buff.Value
	case BUFF_KEEPER_REDUCE_HURT:
		fallthrough
	case BUFF_REDUCE_HURT:
		f.reduceHurt -= buff.Value
		if f.reduceHurt < 0 {
			f.reduceHurt = 0
		}

	case BUFF_DIRECT_REDUCE_HURT:
		f.directReduceHurt -= buff.Value
		if f.directReduceHurt < 0 {
			f.directReduceHurt = 0
		}

	case BUFF_ABSORB_HURT:
		f.absorbHurt -= buff.Value
		if f.absorbHurt < 0 {
			f.absorbHurt = 0
		}

	case BUFF_ATTRACT_FIRE:
		f.attractFire -= 1
	case BUFF_DESTROY_LEVEL:
		f.DestroyLevel -= buff.Value
	case BUFF_TENACITY_LEVEL:
		f.TenacityLevel -= buff.Value
	case BUFF_SLEEP_LEVEL:
		f.SleepLevel -= buff.Value
	case BUFF_DIZZINESS_LEVEL:
		f.DizzinessLevel -= buff.Value
	case BUFF_RANDOM_LEVEL:
		f.RandomLevel -= buff.Value
	case BUFF_DISABLE_SKILL_LEVEL:
		f.DisableSkillLevel -= buff.Value
	case BUFF_POISONING_LEVEL:
		f.PoisoningLevel -= buff.Value
	case BUFF_DISABLE_SKILL:
		f.disableSkill -= 1
	case BUFF_GHOST_SHIELD:
		f.ghostShieldState -= 1
		f.EnableGhostShield = false
		f.ShieldGhostId = 0

	case BUFF_TAKE_SUNDER:
		f.SunderAttack -= int(buff.Value)

	case BUFF_SUNDER_STATE:
		f.sunderValue = f.SunderMaxValue
	//
	// 以下字段需要防止职业加成来回切换后导致buff恢复后值超出上限
	//
	case BUFF_SPEED:
		f.Speed -= buff.Value
		//if f.isLastBuff(buff) && f.Speed > f.raw.Speed {
		//	f.Speed = f.raw.Speed
		//}
	case BUFF_ATTACK:
		f.Attack -= buff.Value
		//if f.isLastBuff(buff) && f.Attack > f.raw.Attack {
		//	f.Attack = f.raw.Attack
		//}
	case BUFF_DEFEND:
		f.Defend -= buff.Value
		//if f.isLastBuff(buff) && f.Defend > f.raw.Defend {
		//	f.Defend = f.raw.Defend
		//}
	case BUFF_BLOCK_LEVEL:
		f.BlockLevel -= buff.Value
		//if f.isLastBuff(buff) && f.BlockLevel > f.raw.BlockLevel {
		//	f.BlockLevel = f.raw.BlockLevel
		//}
	case BUFF_DODGE_LEVEL:
		f.DodgeLevel -= buff.Value
		//if f.isLastBuff(buff) && f.DodgeLevel > f.raw.DodgeLevel {
		//	f.DodgeLevel = f.raw.DodgeLevel
		//}
	case BUFF_CRITIAL_LEVEL:
		f.CritialLevel -= buff.Value
		//if f.isLastBuff(buff) && f.CritialLevel > f.raw.CritialLevel {
		//	f.CritialLevel = f.raw.CritialLevel
		//}

	//一下四个技能可能把角色的属性设置为负数
	case BUFF_DOGE:
		f.Dodge -= buff.Value
	case BUFF_HIT:
		f.Hit -= buff.Value
	case BUFF_CRITIAL:
		f.Critial -= buff.Value
	case BUFF_TENACITY:
		f.Tenacity -= buff.Value

	case BUFF_DEFEND_PERSENT:
		f.defenceRate -= int(buff.Value)
	case BUFF_HEALTH_PERCENT:
		f.Health -= int(float64(f.MaxHealth) * buff.Value / 100)
	case BUFF_ALL_RESIST:
		f.SleepLevel -= buff.Value        // 睡眠抗性
		f.DizzinessLevel -= buff.Value    // 眩晕抗性
		f.RandomLevel -= buff.Value       // 混乱抗性
		f.DisableSkillLevel -= buff.Value // 封魔抗性
		f.PoisoningLevel -= buff.Value    // 中毒抗性
		//case BUFF_REBOTH_HEALTH:
		//	f.Health = 0
		//	f.side.live--
		//case BUFF_REBOTH_HEALTH_PERCENT:
		//	f.Health = 0
		//	f.side.live--
	}
}

func (this *Buff) UpdateOnTurnEnd(f *Fighter) {
	for buff := this; buff != nil; buff = buff.Next {
		buff.Keep -= 1
		if buff.Keep == 0 {
			buff.Revert(f)
			this.remove(buff)
		}
	}
}

// 清除守卫者免伤
func (this *Buff) CleanKeeperReduceHurt(f *Fighter) {
	for buff := this; buff != nil; buff = buff.Next {
		if buff.Mode == BUFF_KEEPER_REDUCE_HURT {
			buff.Revert(f)
			this.remove(buff)
		}
	}
}

// 清除睡眠状态
func (this *Buff) CleanSleep(f *Fighter) {
	for buff := this; buff != nil; buff = buff.Next {
		if buff.Mode == BUFF_SLEEP {
			buff.Revert(f)
			this.remove(buff)
		}
	}
}

// 清除所有Buff
func (this *Buff) CleanAllBuff(f *Fighter) {
	for buff := this; buff != nil; buff = buff.Next {
		if !buff.Uncleanable {
			buff.Revert(f)
			this.remove(buff)
		}
	}
}

func (this *Buff) UpdateAbortHurtBuffValue(value float64, f *Fighter) {
	for buff := this; buff != nil; buff = buff.Next {
		if buff.Mode == BUFF_ABSORB_HURT {
			buff.Value -= value
			if buff.Value <= 0 {
				buff.Revert(f)
				this.remove(buff)
			}
			break
		}
	}
}

func (this *Buff) remove(b *Buff) {
	if b.Prev != nil {
		b.Prev.Next = b.Next
		if b.Next != nil {
			b.Next.Prev = b.Prev
		}
	} else {
		this.Owner.Buffs = b.Next
		if b.Next != nil {
			b.Next.Prev = nil
		}
	}
}

// 检查身上是否有某个绝招的buff
func (f *Fighter) HasBuffFromSkill(skill int) bool {
	for i := f.Buffs; i != nil; i = i.Next {
		if i.Skill == skill {
			return true
		}
	}

	return false
}

// 检查升上是否有某种 buff
func (f *Fighter) HasBuff(mode int) bool {
	for i := f.Buffs; i != nil; i = i.Next {
		if i.Mode == mode {
			return true
		}
	}
	return false
}

// 是否是次类型buff中最后一个
func (f *Fighter) isLastBuff(buff *Buff) bool {
	for i := f.Buffs; i != nil; i = i.Next {
		if i != buff && buff.Mode == i.Mode {
			return false
		}
	}

	return true
}

func buff_value(value float64) float64 {
	if value < 0 {
		return 0
	}
	return value
}

func append_buff(buffs []*Buff, buff *Buff) []*Buff {
	if buff == nil {
		return buffs
	}
	return append(buffs, buff)
}

func IsStateBuff(mode int) bool {
	switch mode {
	case BUFF_SLEEP, BUFF_DIZZINESS, BUFF_RANDOM, BUFF_DISABLE_SKILL:
		return true
	}
	return false
}

// 施加Buff
func (f *Fighter) addBuff(byWho *Fighter, mode int, value float64, keep, skill int, maxOverride int, uncleanable bool) *Buff {
	if value == 0 {
		value = 1
		//return nil
	}

	// 叠加判定
	if keep > 0 && (maxOverride > 0 || IsStateBuff(mode)) {
		for i := f.Buffs; i != nil; i = i.Next {
			// 状态类BUFF覆盖不需要判定是否同一绝招
			if i.Mode == mode && IsStateBuff(mode) {
				if keep > i.Keep { //避免有效时间更短的晕眩覆盖掉晕眩时间长的晕眩
					i.Keep = keep
					i.Skill = skill
				}
				i.ShowKeep = i.Keep
				return i
			}
			if i.Skill == skill && i.Mode == mode {
				if i.OverrideNum < maxOverride {
					// 撤销之前buff效果、刷新持续回合、叠加作用效果
					i.Revert(f)
					if keep > i.Keep { //避免有效时间更短的晕眩覆盖掉晕眩时间长的晕眩
						i.Keep = keep
						i.Skill = skill
					}
					i.Value += value
					i.OverrideNum += 1
					i.Apply(f)

					if i.OverrideNum > maxOverride {
						i.OverrideNum = maxOverride
					}
				} else {
					// 只刷新持续回合
					i.Keep = keep
					i.Skill = skill
					// 覆盖旧的伤害吸收
					if mode == BUFF_ABSORB_HURT {
						i.Revert(f)
						i.Value = value
						i.Apply(f)
					}
				}
				i.ShowKeep = i.Keep

				return i
			}
		}
	}

	b := &Buff{Mode: mode,
		Value:       value,
		Keep:        keep,
		ShowKeep:    keep,
		Skill:       skill,
		MaxOverride: maxOverride,
		OverrideNum: 1,
		Owner:       f,
		updateRound: f.battle.GetRounds() - 1,
		Uncleanable: uncleanable,
	}
	b.Apply(f)

	if keep != -1 {
		if f.Buffs != nil {
			f.Buffs.Prev = b
		}
		b.Next = f.Buffs
		f.Buffs = b
	}

	return b
}
