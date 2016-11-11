package battle

import (
	"core/fail"
	"math/rand"
)

/*
	战斗道具
*/

type itemApplyFunc func(*Item, *Fighter, *Fighter) []*Buff
type itemGetTargets func(*Fighter) []*Fighter
type itemForNormal func(attr *FighterAttribute)
type itemForBuff func(*Fighter, *Item)

// 已捕捉到的灵宠ID
var catchBattlePetId int32 = -1

type ItemEffect struct {
	Mode  int     // 影响属性类型
	Value float64 // 产生效果值
}

type Item struct {
	ItemId      int32 // 道具ID
	Target      int   // 道具作用对象,like ITEM_EFFECT_TARGET_XXXX
	EffectType  int   // 在战场中产生的效果类型
	Keep        int   // 持续回合数
	MaxOverride int   // 最大叠加数
	CostPower   int
	Effect      []*ItemEffect // 待作用的buff

	ApplyItemForNormal itemForNormal // for 关卡内使用
	ApplyItemForBuff   itemForBuff   // for 关卡内使用

	_applyFunc  itemApplyFunc
	_GetTargets itemGetTargets
}

// 战场道具使用结果
type UseItemResult struct {
	ItemId int32
	Result []EffectResult
}

type EffectResult struct {
	Type      int     // 被攻击方类型，FT_开头的常量
	TargetPos int     // 目标站位ID
	Health    int     // 使用后的生命值
	Power     int     // 使用后的精气值
	Hurt      int     // 受到的伤害
	Buffs     []*Buff // 对目标产生的buff
}

func (user *Fighter) UseBattleItem(item_id int32) (result *UseItemResult, ok bool) {
	item := createBattleItem(item_id)
	fail.When(item == nil, "use incorrect battle item id")

	// 使用道具消耗精力，精力不够
	if user.Power-item.CostPower < 0 {
		return
	}

	user.Power = user.Power - item.CostPower

	result = &UseItemResult{}
	result.ItemId = item.ItemId

	for _, fighter := range item._GetTargets(user) {
		if fighter != nil {
			result.Result = append(result.Result, applyItem(item, user, fighter))
		}
	}

	ok = true
	return
}

func UseLevelItem(getFighter func() *Fighter, item_id int32, attr *FighterAttribute, fighterBuffs []Buff) (buffs []Buff) {
	item := createBattleItem(item_id)
	fail.When(item == nil, "use incorrect level item id")

	if item.ApplyItemForNormal != nil {
		item.ApplyItemForNormal(attr)
	} else if item.ApplyItemForBuff != nil {
		fighter := getFighter()
		fighter.applyFighterAttribute(attr)
		fighter.applyBuffs(fighterBuffs)
		item.ApplyItemForBuff(fighter, item)
		*attr = fighter.exportAttribute()
		buffs = fighter.exportBuff()
	}

	return
}

func applyItem(item *Item, user *Fighter, target *Fighter) EffectResult {
	var result EffectResult
	var buffs []*Buff
	var oldHurt, realHurt int

	oldHurt = target.Health

	buffs = item._applyFunc(item, user, target)

	if oldHurt > target.Health {
		realHurt = oldHurt - target.Health
		// 累计角色造成的伤害
		user.TotalHurt += realHurt

		if target.Health <= MIN_HEALTH {
			target.side.live -= 1

			if !target.IsBattlePet {
				target.side.deadNums += 1
			}

			if target.IsBoss {
				target.side.live = 0
			}
		}
	}

	result.Type = target.Ftype
	result.TargetPos = target.Position
	result.Health = target.Health
	result.Power = target.Power
	result.Hurt = realHurt
	result.Buffs = buffs
	return result
}

// 复活
func reliveFighter(user *Fighter, target *Fighter, effect *ItemEffect) {
	target.Health = 0
	target.Health += int(effect.Value / 100 * float64(target.MaxHealth))
	if target.Health > target.MaxHealth {
		target.Health = target.MaxHealth
	}
	target.side.live++
	target.Buffs.CleanAllBuff(target)
	target.sunderValue = target.SunderMaxValue
}

func (b *BattleState) GetCatchBattlePetId() (petId int32) {
	petId = catchBattlePetId
	catchBattlePetId = -1
	return
}

// 捕捉
func catchPet(user *Fighter, target *Fighter, effect *ItemEffect) {
	randNum := rand.Intn(100) + 1
	catchBattlePetId = int32(0)
	if randNum <= int(effect.Value)+getBattlePetCatchRate(target) {
		// 成功捕捉后灵宠消失
		target.Health = MIN_HEALTH
		catchBattlePetId = int32(target.RoleId)
	}
}

// 全体存活着的
func findSideLiveFighter(sideFighter []*Fighter) (fighters []*Fighter) {
	for _, f := range sideFighter {
		if f != nil && f.Health > MIN_HEALTH {
			fighters = append(fighters, f)
		}
	}

	return
}

// 生命最少的角色(按百分比)
func findLeastHealthFighter(fighters []*Fighter) *Fighter {
	var targetFighter *Fighter
	var targetPercent float64

	for _, fighter := range fighters {
		if fighter != nil && fighter.Health > MIN_HEALTH {
			percent := float64(fighter.MaxHealth-fighter.Health) / float64(fighter.MaxHealth)

			if targetFighter == nil || percent > targetPercent || (percent == targetPercent && fighter.Health < targetFighter.Health) {
				targetFighter = fighter
				targetPercent = percent
			}
		}
	}

	return targetFighter
}

// 死亡的角色不包含灵宠
func findDeadFighter(fighters []*Fighter) *Fighter {
	for _, fighter := range fighters {
		if fighter != nil && fighter.Health <= MIN_HEALTH && !fighter.IsBattlePet {
			return fighter
		}
	}
	return nil
}

// 找灵宠
func findBattlePet(fighters []*Fighter) *Fighter {
	for _, fighter := range fighters {
		if fighter != nil && fighter.Health > MIN_HEALTH && fighter.IsBattlePet {
			return fighter
		}
	}
	return nil
}
