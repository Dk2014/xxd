package battle


	/*
	目标对象
	1 -- 道具使用者
	2 -- 我方全体
	3 -- 我方血量最少的
	4 -- 敌方全体
	5 -- 敌方血量最少的
	6 -- 我方阵亡角色
	7 -- 灵宠
	8 -- 我方伙伴
	9 -- 主角和伙伴
	10 -- 敌方单体


	产生效果
	1 -- 基础属性
	2 -- 捕捉球
	3 -- Buff增益


	基础属性类型
	0 -- 精气
	1 -- 生命
	3 -- 复活
	4 -- 捕捉概率(%)


	Buff增益类型(对应buff类型)
	0 -- (BUFF_POWER)精气
	4 -- (BUFF_HEALTH)生命
	7 -- (BUFF_CLEAN_BAD)清除负面状态
	13 -- (BUFF_DODGE_LEVEL)闪避等级
	16 -- (BUFF_HURT_ADD)伤害加成(%)
	29 -- (BUFF_BUDDY_SKILL)伙伴技能
	36 -- (BUFF_RECOVER_BUDDY_SKILL)回复伙伴技能
	37 -- (BUFF_MAKE_POWER_FULL)增加精气到满

	*/

func createBattleItem(id int32) *Item {
	switch id {
	case 209:
		return battle_item_209 // 止血草
	case 210:
		return battle_item_210 // 金创药
	case 211:
		return battle_item_211 // 大还丹
	case 212:
		return battle_item_212 // 凤凰羽毛
	case 213:
		return battle_item_213 // 飞刀
	case 214:
		return battle_item_214 // 暗影飞刀
	case 215:
		return battle_item_215 // 风的种子
	case 216:
		return battle_item_216 // 火的种子
	case 217:
		return battle_item_217 // 水的种子
	case 250:
		return battle_item_250 // 契约球
	case 251:
		return battle_item_251 // 至尊契约球
	case 327:
		return battle_item_327 // 仙豆
	case 353:
		return battle_item_353 // BUG飞刀
	case 354:
		return battle_item_354 // 高级契约球
	case 730:
		return battle_item_730 // BUG血药
		
	}
	return nil
}


//止血草
var battle_item_209 = &Item{
	ItemId:      209,
	Target:      2,
	EffectType:  1,
	Keep:        0,
	CostPower: 	 1,
	MaxOverride: 0,
	Effect: []*ItemEffect{
		&ItemEffect{
			Mode:  1,
			Value: 3000,
		},
	},
	ApplyItemForNormal: func(attr *FighterAttribute) {
		attr.Health += (3000)
	},
	_GetTargets: func(user *Fighter) []*Fighter {
		return findSideLiveFighter(user.battle.Attackers.Fighters)
	},
	_applyFunc: func(item *Item, user *Fighter, target *Fighter) []*Buff {
		target.Health += (3000)
		if target.Health > target.MaxHealth {
			target.Health = target.MaxHealth
		}

		return nil
	},
}

//金创药
var battle_item_210 = &Item{
	ItemId:      210,
	Target:      2,
	EffectType:  1,
	Keep:        0,
	CostPower: 	 2,
	MaxOverride: 0,
	Effect: []*ItemEffect{
		&ItemEffect{
			Mode:  1,
			Value: 10000,
		},
	},
	ApplyItemForNormal: func(attr *FighterAttribute) {
		attr.Health += (10000)
	},
	_GetTargets: func(user *Fighter) []*Fighter {
		return findSideLiveFighter(user.battle.Attackers.Fighters)
	},
	_applyFunc: func(item *Item, user *Fighter, target *Fighter) []*Buff {
		target.Health += (10000)
		if target.Health > target.MaxHealth {
			target.Health = target.MaxHealth
		}

		return nil
	},
}

//大还丹
var battle_item_211 = &Item{
	ItemId:      211,
	Target:      2,
	EffectType:  1,
	Keep:        0,
	CostPower: 	 3,
	MaxOverride: 0,
	Effect: []*ItemEffect{
		&ItemEffect{
			Mode:  1,
			Value: 30000,
		},
	},
	ApplyItemForNormal: func(attr *FighterAttribute) {
		attr.Health += (30000)
	},
	_GetTargets: func(user *Fighter) []*Fighter {
		return findSideLiveFighter(user.battle.Attackers.Fighters)
	},
	_applyFunc: func(item *Item, user *Fighter, target *Fighter) []*Buff {
		target.Health += (30000)
		if target.Health > target.MaxHealth {
			target.Health = target.MaxHealth
		}

		return nil
	},
}

//凤凰羽毛
var battle_item_212 = &Item{
	ItemId:      212,
	Target:      6,
	EffectType:  1,
	Keep:        0,
	CostPower: 	 4,
	MaxOverride: 0,
	Effect: []*ItemEffect{
		&ItemEffect{
			Mode:  3,
			Value: 100,
		},
	},

	_GetTargets: func(user *Fighter) []*Fighter {
		return []*Fighter{findDeadFighter(user.battle.Attackers.Fighters)}
	},
	_applyFunc: func(item *Item, user *Fighter, target *Fighter) []*Buff {
		reliveFighter(user, target, item.Effect[0])
		return nil
	},
}

//飞刀
var battle_item_213 = &Item{
	ItemId:      213,
	Target:      10,
	EffectType:  1,
	Keep:        0,
	CostPower: 	 4,
	MaxOverride: 0,
	Effect: []*ItemEffect{
		&ItemEffect{
			Mode:  1,
			Value: -10000,
		},
	},

	_GetTargets: func(user *Fighter) []*Fighter {
		return findOneTarget(user)
	},
	_applyFunc: func(item *Item, user *Fighter, target *Fighter) []*Buff {
		target.Health += (-10000)
		if target.Health < MIN_HEALTH {
			target.Health = MIN_HEALTH
		}

		return nil
	},
}

//暗影飞刀
var battle_item_214 = &Item{
	ItemId:      214,
	Target:      10,
	EffectType:  1,
	Keep:        0,
	CostPower: 	 4,
	MaxOverride: 0,
	Effect: []*ItemEffect{
		&ItemEffect{
			Mode:  1,
			Value: -30000,
		},
	},

	_GetTargets: func(user *Fighter) []*Fighter {
		return findOneTarget(user)
	},
	_applyFunc: func(item *Item, user *Fighter, target *Fighter) []*Buff {
		target.Health += (-30000)
		if target.Health < MIN_HEALTH {
			target.Health = MIN_HEALTH
		}

		return nil
	},
}

//风的种子
var battle_item_215 = &Item{
	ItemId:      215,
	Target:      2,
	EffectType:  3,
	Keep:        2,
	CostPower: 	 4,
	MaxOverride: 1,
	Effect: []*ItemEffect{
		&ItemEffect{
			Mode:  13,
			Value: 1000,
		},
	},

	_GetTargets: func(user *Fighter) []*Fighter {
		return findSideLiveFighter(user.battle.Attackers.Fighters)
	},
	_applyFunc: func(item *Item, user *Fighter, target *Fighter) []*Buff {
		return []*Buff{target.addBuff(user, 13, 1000, item.Keep, 0, item.MaxOverride, false)}
	},
}

//火的种子
var battle_item_216 = &Item{
	ItemId:      216,
	Target:      1,
	EffectType:  3,
	Keep:        3,
	CostPower: 	 2,
	MaxOverride: 1,
	Effect: []*ItemEffect{
		&ItemEffect{
			Mode:  16,
			Value: 20,
		},
	},

	_GetTargets: func(user *Fighter) []*Fighter {
		return []*Fighter{user}
	},
	_applyFunc: func(item *Item, user *Fighter, target *Fighter) []*Buff {
		return []*Buff{target.addBuff(user, 16, 20, item.Keep, 0, item.MaxOverride, false)}
	},
}

//水的种子
var battle_item_217 = &Item{
	ItemId:      217,
	Target:      2,
	EffectType:  3,
	Keep:        1,
	CostPower: 	 2,
	MaxOverride: 1,
	Effect: []*ItemEffect{
		&ItemEffect{
			Mode:  7,
			Value: 0,
		},
	},

	_GetTargets: func(user *Fighter) []*Fighter {
		return findSideLiveFighter(user.battle.Attackers.Fighters)
	},
	_applyFunc: func(item *Item, user *Fighter, target *Fighter) []*Buff {
		return []*Buff{target.addBuff(user, 7, 0, item.Keep, 0, item.MaxOverride, false)}
	},
}

//契约球
var battle_item_250 = &Item{
	ItemId:      250,
	Target:      7,
	EffectType:  2,
	Keep:        0,
	CostPower: 	 1,
	MaxOverride: 0,
	Effect: []*ItemEffect{
		&ItemEffect{
			Mode:  4,
			Value: 30,
		},
	},

	_GetTargets: func(user *Fighter) []*Fighter {
		return []*Fighter{findBattlePet(user.battle.Defenders.Fighters)}
	},
	_applyFunc: func(item *Item, user *Fighter, target *Fighter) []*Buff {
		catchPet(user, target, item.Effect[0])
		return nil
	},
}

//至尊契约球
var battle_item_251 = &Item{
	ItemId:      251,
	Target:      7,
	EffectType:  2,
	Keep:        0,
	CostPower: 	 1,
	MaxOverride: 0,
	Effect: []*ItemEffect{
		&ItemEffect{
			Mode:  4,
			Value: 100,
		},
	},

	_GetTargets: func(user *Fighter) []*Fighter {
		return []*Fighter{findBattlePet(user.battle.Defenders.Fighters)}
	},
	_applyFunc: func(item *Item, user *Fighter, target *Fighter) []*Buff {
		catchPet(user, target, item.Effect[0])
		return nil
	},
}

//仙豆
var battle_item_327 = &Item{
	ItemId:      327,
	Target:      9,
	EffectType:  3,
	Keep:        1,
	CostPower: 	 0,
	MaxOverride: 1,
	Effect: []*ItemEffect{
		&ItemEffect{
			Mode:  36,
			Value: 1,
		},		&ItemEffect{
			Mode:  37,
			Value: 1,
		},
	},

	_GetTargets: func(user *Fighter) []*Fighter {
		return findOurSideWithoutPet(user)
	},
	_applyFunc: func(item *Item, user *Fighter, target *Fighter) []*Buff {
		return []*Buff{target.addBuff(user, 36, 1, item.Keep, 0, item.MaxOverride, false),target.addBuff(user, 37, 1, item.Keep, 0, item.MaxOverride, false)}
	},
}

//BUG飞刀
var battle_item_353 = &Item{
	ItemId:      353,
	Target:      4,
	EffectType:  1,
	Keep:        0,
	CostPower: 	 0,
	MaxOverride: 0,
	Effect: []*ItemEffect{
		&ItemEffect{
			Mode:  1,
			Value: -999999,
		},
	},

	_GetTargets: func(user *Fighter) []*Fighter {
		return findSideLiveFighter(user.battle.Defenders.Fighters)
	},
	_applyFunc: func(item *Item, user *Fighter, target *Fighter) []*Buff {
		target.Health += (-999999)
		if target.Health < MIN_HEALTH {
			target.Health = MIN_HEALTH
		}

		return nil
	},
}

//高级契约球
var battle_item_354 = &Item{
	ItemId:      354,
	Target:      7,
	EffectType:  2,
	Keep:        0,
	CostPower: 	 1,
	MaxOverride: 0,
	Effect: []*ItemEffect{
		&ItemEffect{
			Mode:  4,
			Value: 60,
		},
	},

	_GetTargets: func(user *Fighter) []*Fighter {
		return []*Fighter{findBattlePet(user.battle.Defenders.Fighters)}
	},
	_applyFunc: func(item *Item, user *Fighter, target *Fighter) []*Buff {
		catchPet(user, target, item.Effect[0])
		return nil
	},
}

//BUG血药
var battle_item_730 = &Item{
	ItemId:      730,
	Target:      2,
	EffectType:  1,
	Keep:        0,
	CostPower: 	 0,
	MaxOverride: 0,
	Effect: []*ItemEffect{
		&ItemEffect{
			Mode:  1,
			Value: 999999,
		},
	},
	ApplyItemForNormal: func(attr *FighterAttribute) {
		attr.Health += (999999)
	},
	_GetTargets: func(user *Fighter) []*Fighter {
		return findSideLiveFighter(user.battle.Attackers.Fighters)
	},
	_applyFunc: func(item *Item, user *Fighter, target *Fighter) []*Buff {
		target.Health += (999999)
		if target.Health > target.MaxHealth {
			target.Health = target.MaxHealth
		}

		return nil
	},
}

