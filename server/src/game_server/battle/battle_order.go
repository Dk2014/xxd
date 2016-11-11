package battle

// 战斗顺序
type FightOrder struct {
	count    int // 剩余出手角色数量
	current  int // 当前出手到谁，用于减少遍历数量，重新排序时被重置
	fighters []*Fighter
}

func (this *FightOrder) Len() int {
	return len(this.fighters)
}

func (this *FightOrder) Swap(i, j int) {
	this.fighters[i], this.fighters[j] = this.fighters[j], this.fighters[i]
}

func (this *FightOrder) Less(i, j int) bool {
	if this.fighters[i].Speed > this.fighters[j].Speed {
		return true
	}

	if this.fighters[i].Speed == this.fighters[j].Speed {
		if this.fighters[i].Position < this.fighters[j].Position {
			return true
		}

		if this.fighters[i].Position == this.fighters[j].Position {
			return this.fighters[i].Ftype < this.fighters[j].Ftype
		}
	}

	return false
}

func (this *FightOrder) Reset() {
	this.current = 0
	this.fighters = this.fighters[0:0]
}

func (this *FightOrder) Append(fighters []*Fighter, resetFighted bool) {
	for _, fighter := range fighters {
		if fighter == nil {
			continue
		}
		//如果怪物或灵宠死亡则排除再 oder 之外，如果是玩家或者伙伴则保留，因为可能复活
		if fighter.Health <= MIN_HEALTH && (fighter.Kind == FK_ENEMY || fighter.IsBattlePet) {
			continue
		}

		this.fighters = append(this.fighters, fighter)

		if resetFighted {
			this.count += 1
			fighter.round += 1
			fighter.fighted = false
		}
	}
}
