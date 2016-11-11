package battle

import "math/rand"

// 站位编号：
// ----------------
// |15|14|13|12|11|
// ----------------
// |10| 9| 8| 7| 6|  [守方]
// ----------------
// | 5| 4| 3| 2| 1|
// ----------------
//
// ----------------
// | 1| 2| 3| 4| 5|
// ----------------
// | 6| 7| 8| 9|10|  [攻方]
// ----------------
// |11|12|13|14|15|
// ----------------

//
// 单体攻击目标索引
//
// 说明：
// 站位总共3行5列，不同行的同一列的三个位置，按出手规则攻击目标选取顺序其实是相同的
// 所以这里用位置编号跟5取模的方式得到1 ~ 5列的攻击目标列表
//
// [站位类型][]int
var attackOneIndex = [][]int{
	{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}, // 0 对应 5、10、15
	{4, 3, 2, 1, 0, 9, 8, 7, 6, 5, 14, 13, 12, 11, 10}, // 1 对应 1、 6、11
	{3, 4, 2, 1, 0, 8, 9, 7, 6, 5, 13, 14, 12, 11, 10}, // 2 对应 2、 7、12
	{2, 3, 1, 4, 0, 7, 8, 6, 9, 5, 12, 13, 11, 14, 10}, // 3 对应 3、 8、13
	{1, 2, 0, 3, 4, 6, 7, 5, 8, 9, 11, 12, 10, 13, 14}, // 4 对应 4、 9、14
}

// 横向攻击目标
var attackRowIndex = [][]int{
	{0, 1, 2, 3, 4},
	{5, 6, 7, 8, 9},
	{10, 11, 12, 13, 14},
}

// 横向攻击目标(从后向前)
var attackLastRowIndex = [][]int{
	{10, 11, 12, 13, 14},
	{5, 6, 7, 8, 9},
	{0, 1, 2, 3, 4},
}

// 纵向攻击目标
var attackColIndex = [][]int{
	{4, 9, 14}, // 0 对应 5、10、15
	{0, 5, 10}, // 1 对应 1、 6、11
	{1, 6, 11}, // 2 对应 2、 7、12
	{2, 7, 12}, // 3 对应 3、 8、13
	{3, 8, 13}, // 4 对应 4、 9、14
}

// V字攻击索引
// 使用单体攻击所选中的目标的位置减1作为索引
var vAttackIndex = [][]int{
	{0, 6, 12},
	{1, 5, 7, 13},
	{2, 6, 8, 10, 14},
	{3, 7, 9, 11},
	{4, 8, 12},
	{5, 11},
	{6, 10, 12},
	{7, 11, 13},
	{8, 12, 14},
	{9, 13},
	{10},
	{11},
	{12},
	{13},
	{14},
}

// 十字攻击索引
var crossAttackIndex = [][]int{
	{0, 1, 2, 3, 4, 5, 10},
	{0, 1, 2, 3, 4, 6, 11},
	{0, 1, 2, 3, 4, 7, 12},
	{0, 1, 2, 3, 4, 8, 13},
	{0, 1, 2, 3, 4, 9, 14},
	{5, 6, 7, 8, 9, 10},
	{5, 6, 7, 8, 9, 11},
	{5, 6, 7, 8, 9, 12},
	{5, 6, 7, 8, 9, 13},
	{5, 6, 7, 8, 9, 14},
	{10, 11, 12, 13, 14},
	{10, 11, 12, 13, 14},
	{10, 11, 12, 13, 14},
	{10, 11, 12, 13, 14},
	{10, 11, 12, 13, 14},
}

// 固定十字攻击索引
var fixCrossAttackIndex = []int{
	2, 5, 6, 7, 8, 9, 12,
}

//前排固定攻击
var fixFrontendRowIndex = []int{
	0, 1, 2, 3, 4,
}

//中排固定攻击
var fixMiddleRowIndex = []int{
	5, 6, 7, 8, 9,
}
var fixBackendRowIndex = []int{
	10, 11, 12, 13, 14,
}

// 后排单体攻击索引
var attackOneFromBackIndex = [][]int{
	{10, 11, 12, 13, 14, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4}, // 0 对应 5、10、15
	{14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, // 1 对应 1、 6、11
	{13, 14, 12, 11, 10, 8, 9, 7, 6, 5, 3, 4, 2, 1, 0}, // 2 对应 2、 7、12
	{12, 13, 11, 14, 10, 7, 8, 6, 9, 5, 2, 3, 1, 4, 0}, // 3 对应 3、 8、13
	{11, 12, 10, 13, 14, 6, 7, 5, 8, 9, 1, 2, 0, 3, 4}, // 4 对应 4、 9、14
}

func isSameRow(pos1, pos2 int) bool {
	switch {
	case pos1 >= 0 && pos1 <= 4:
		return pos2 >= 0 && pos2 <= 4
	case pos1 >= 5 && pos1 <= 9:
		return pos2 >= 5 && pos2 <= 9
	case pos1 >= 10 && pos1 <= 14:
		return pos2 >= 10 && pos2 <= 14
	}

	return false
}

// 单体攻击
func findOneTarget(f *Fighter) []*Fighter {
	targets := f.getTargets()

	otherSide := f.getOtherSide()
	if otherSide.live <= 0 {
		return nil
	}

	result := f.battle.targets[0:0]

	normalTarget := -1

	for _, i := range attackOneIndex[f.ptype] {
		// 已经有普通目标的时候只遍历到本行结束
		if normalTarget != -1 && !isSameRow(normalTarget, i) {
			break
		}

		target := targets[i]

		if target != nil && target.Health > MIN_HEALTH {
			// 本行中未被破甲的守卫者优先
			if target.attractFire > 0 && target.sunderValue > 0 {
				return append(result, target)
			}

			if normalTarget == -1 {
				normalTarget = i
			}
		}
	}

	if normalTarget != -1 {
		return append(result, targets[normalTarget])
	}

	return nil
}

// 单体攻击（2次）
func findOneTarget2(f *Fighter) []*Fighter {
	targets := findOneTarget(f)

	if targets == nil {
		return nil
	}

	return []*Fighter{targets[0], targets[0]}
}

// 横向攻击基础逻辑
func findRowTarget2(f *Fighter, indexs [][]int) []*Fighter {
	targets := f.getTargets()

	result := f.battle.targets[0:0]

	for _, index := range indexs {
		for _, i := range index {
			t := targets[i]

			if t != nil && t.Health > MIN_HEALTH {
				result = append(result, t)
			}
		}

		if len(result) > 0 {
			return result
		}
	}

	return nil
}

// 前排横向攻击
func findRowTarget(f *Fighter) []*Fighter {
	return findRowTarget2(f, attackRowIndex)
}

// 后排横向攻击
func findLastRowTargets(f *Fighter) []*Fighter {
	return findRowTarget2(f, attackLastRowIndex)
}

// 纵向攻击
func findColTarget(f *Fighter) []*Fighter {
	// 先找到单体攻击目标
	colTarget := findOneTarget(f)

	if colTarget == nil {
		return nil
	}

	targets := f.getTargets()

	result := f.battle.targets[0:0]

	// 找到目标所属的列上的其它目标
	for _, j := range attackColIndex[colTarget[0].ptype] {
		t := targets[j]

		if t != nil && t.Health > MIN_HEALTH {
			result = append(result, t)
		}
	}

	return result
}

// 全体攻击
func findAllTargets(f *Fighter) []*Fighter {
	targets := f.getTargets()

	result := f.battle.targets[0:0]

	for _, t := range targets {
		if t != nil && t.Health > MIN_HEALTH {
			result = append(result, t)
		}
	}

	return result
}

// 最少血攻击
func findLeastHealthTarget(f *Fighter) []*Fighter {
	targets := f.getTargets()

	var last *Fighter

	for _, i := range attackOneIndex[f.ptype] {
		t := targets[i]
		if t != nil && t.Health > MIN_HEALTH {
			if last == nil || t.Health < last.Health {
				last = t
			}
		}
	}

	if last == nil {
		return nil
	}

	return []*Fighter{last}
}

// 最少血攻击（两次）
func findLeastHealthTarget2(f *Fighter) []*Fighter {
	targets := f.getTargets()

	var last *Fighter

	for _, i := range attackOneIndex[f.ptype] {
		t := targets[i]
		if t != nil && t.Health > MIN_HEALTH {
			if last == nil || t.Health < last.Health {
				last = t
			}
		}
	}

	if last == nil {
		return nil
	}

	return []*Fighter{last, last}
}

// 最大血攻击
func findMostHealthTarget(f *Fighter) []*Fighter {
	targets := f.getTargets()

	var last *Fighter

	for _, i := range attackOneIndex[f.ptype] {
		t := targets[i]
		if t != nil && t.Health > MIN_HEALTH {
			if last == nil || t.Health > last.Health {
				last = t
			}
		}
	}

	if last == nil {
		return nil
	}

	return []*Fighter{last}
}

// V字攻击
func findVAttackTarget(f *Fighter) []*Fighter {
	// 先找到单体攻击目标
	colTarget := findOneTarget(f)

	if colTarget == nil {
		return nil
	}

	targets := f.getTargets()

	result := f.battle.targets[0:0]

	// 找到目标所属的列上的其它目标
	for _, j := range vAttackIndex[colTarget[0].Position-1] {
		t := targets[j]

		if t != nil && t.Health > MIN_HEALTH {
			result = append(result, t)
		}
	}

	return result
}

// 后排单体攻击
func findOneTargetFromBack(f *Fighter) []*Fighter {
	targets := f.getTargets()

	result := f.battle.targets[0:0]

	normalTarget := -1

	for _, i := range attackOneFromBackIndex[f.ptype] {
		// 已经有普通目标的时候只遍历到本行结束
		if normalTarget != -1 && !isSameRow(normalTarget, i) {
			break
		}

		target := targets[i]

		if target != nil && target.Health > MIN_HEALTH {
			// 本行中未被破甲的守卫者优先
			if target.attractFire > 0 && target.sunderValue > 0 {
				return append(result, target)
			}

			if normalTarget == -1 {
				normalTarget = i
			}
		}
	}

	if normalTarget != -1 {
		return append(result, targets[normalTarget])
	}

	return nil
}

// 十字攻击
func findCrossTarget(f *Fighter) []*Fighter {
	// 先找到单体攻击目标
	colTarget := findOneTarget(f)

	if colTarget == nil {
		return nil
	}

	targets := f.getTargets()

	result := f.battle.targets[0:0]

	// 找到目标所属的列上的其它目标
	for _, j := range crossAttackIndex[colTarget[0].Position-1] {
		t := targets[j]

		if t != nil && t.Health > MIN_HEALTH {
			result = append(result, t)
		}
	}

	return result
}

// 十字攻击(固定)
func findFixCrossTarget(f *Fighter) []*Fighter {
	targets := f.getTargets()

	result := f.battle.targets[0:0]

	// 找到目标所属的列上的其它目标
	for _, j := range fixCrossAttackIndex {
		t := targets[j]

		if t != nil && t.Health > MIN_HEALTH {
			result = append(result, t)
		}
	}

	return result
}

// 横排穿透攻击
func findRowPenetrateTarget(f *Fighter) []*Fighter {
	targets := f.getTargets()

	result := f.battle.targets[0:0]

	for _, col := range attackColIndex {
		for _, i := range col {
			t := targets[i]

			if t != nil && t.Health > MIN_HEALTH {
				result = append(result, t)
				break
			}
		}
	}

	return result
}

// 随机目标
func findRandomTarget(num int) func(f *Fighter) []*Fighter {
	return func(f *Fighter) []*Fighter {
		targets := findAllTargets(f)

		if len(targets) == 0 {
			return nil
		}

		result := f.battle.targets[0:num]

		for i := 0; i < num; i++ {
			result[i] = targets[rand.Intn(len(targets))]
		}

		return result
	}
}

func getRandomOne(f *Fighter, targets []*Fighter) *Fighter {
	if targets == nil {
		return nil
	}

	otherTargets := make([]*Fighter, 0, 5)

	for _, target := range f.getTargets() {
		if target != nil && target.Health > MIN_HEALTH {
			duplicated := false
			for _, target2 := range targets {
				if target2 == target {
					duplicated = true
					break
				}
			}
			if !duplicated {
				otherTargets = append(otherTargets, target)
			}
		}
	}

	if len(otherTargets) == 0 {
		return nil
	}

	if len(otherTargets) == 1 {
		return otherTargets[0]
	}

	return otherTargets[rand.Intn(len(otherTargets))]
}

func appendRandomOne(f *Fighter, targets []*Fighter) []*Fighter {
	if targets == nil {
		return nil
	}

	randomOne := getRandomOne(f, targets)

	if randomOne == nil {
		return targets
	}

	return append(targets, randomOne)
}

func appendRandomCol(f *Fighter, targets []*Fighter) []*Fighter {
	if targets == nil {
		return nil
	}

	randomOne := getRandomOne(f, targets)

	if randomOne == nil {
		return targets
	}

	allTargets := f.getTargets()

	// 找到目标所属的列上的其它目标
	for _, j := range attackColIndex[randomOne.ptype] {
		t := allTargets[j]

		if t != nil && t.Health > MIN_HEALTH {
			targets = append(targets, t)
		}
	}

	return targets
}

// 纵向攻击附带随机攻击另一列
func findColTargetWithOneColRandom(f *Fighter) []*Fighter {
	return appendRandomCol(f, findColTarget(f))
}

// 随机攻击两个目标
func findTowRandomTargets(f *Fighter) []*Fighter {
	if findOneTarget(f) == nil {
		return nil
	}

	return appendRandomOne(f,
		appendRandomOne(f, f.battle.targets[0:0]),
	)
}

// 单体攻击，并额外随机攻击另一人
func findOneTargetWithOneRandom(f *Fighter) []*Fighter {
	return appendRandomOne(f, findOneTarget(f))
}

// 最少血的伙伴(按百分比)
func findLeastHealthBuddy(f *Fighter) *Fighter {
	buddies := f.getBuddies()

	var targetBuddy *Fighter
	var targetPercent float64

	for _, buddy := range buddies {
		if buddy != nil && buddy.Health > MIN_HEALTH {
			percent := float64(buddy.MaxHealth-buddy.Health) / float64(buddy.MaxHealth)

			if targetBuddy == nil || percent > targetPercent || (percent == targetPercent && buddy.Health < targetBuddy.Health) {
				targetBuddy = buddy
				targetPercent = percent
			}
		}
	}

	return targetBuddy
}

// 主角优先
func findBuddyMainRoleFirst(f *Fighter) *Fighter {
	targets := f.getBuddies()

	var last *Fighter

	for _, i := range attackOneIndex[f.ptype] {
		t := targets[i]
		if t != nil && t.Health > MIN_HEALTH {
			if last == nil || t.Kind == FK_PLAYER {
				last = t
			}
		}
	}

	return last
}

// 伙伴优先，并且不包括自己
func findBuddyWithoutSelf(f *Fighter) []*Fighter {
	targets := f.getBuddies()

	var last *Fighter

	for _, i := range attackOneIndex[f.ptype] {
		t := targets[i]
		if t != nil && t != f && t.Health > MIN_HEALTH {
			last = t
			break
		}
	}

	if last == nil {
		return nil
	}

	return []*Fighter{last}
}

func buddyFindTarget(realFunc func(*Fighter) []*Fighter) func(*Fighter) []*Fighter {
	return func(f *Fighter) []*Fighter {
		fastest := f
		buddies := f.getBuddies()

		for _, buddy := range buddies {
			if buddy == nil || buddy.Health <= MIN_HEALTH {
				continue
			}

			if buddy.Kind == FK_PLAYER {
				return realFunc(buddy)
			}

			if fastest.Speed < buddy.Speed {
				fastest = buddy
			}
		}

		return realFunc(fastest)
	}
}

// 我方灵宠
func findBattlePetBuddy(f *Fighter) (pets []*Fighter) {
	for _, fighter := range f.side.Fighters {
		if fighter != nil && fighter.IsBattlePet && fighter.Health > MIN_HEALTH {
			pets = append(pets, fighter)
		}
	}
	return pets
}

// 我方全体伙伴
func findOurBuddies(f *Fighter) (fighters []*Fighter) {
	for _, fighter := range f.side.Fighters {
		if fighter != nil && fighter.Health > MIN_HEALTH && fighter.Kind == FK_BUDDY {
			fighters = append(fighters, fighter)
		}
	}
	return fighters
}

func findOurSideWithoutPet(f *Fighter) (fighters []*Fighter) {
	for _, fighter := range f.side.Fighters {
		if fighter != nil && fighter.Health > MIN_HEALTH && !fighter.IsBattlePet {
			fighters = append(fighters, fighter)
		}
	}
	return fighters
}

func findOurRoleByFixRow(f *Fighter, index []int) (fighters []*Fighter) {
	targets := f.getBuddies()
	for _, i := range index {
		t := targets[i]
		if t != nil && t.Health > MIN_HEALTH {
			fighters = append(fighters, t)
		}
	}
	return fighters
}

func findTargetByFixRow(f *Fighter, index []int) []*Fighter {
	result := f.battle.targets[0:0]
	targets := f.getTargets()
	for _, i := range index {
		t := targets[i]
		if t != nil && t.Health > MIN_HEALTH {
			result = append(result, t)
		}
	}
	return result
}

func findFixFrontendRow(f *Fighter) []*Fighter {
	return findTargetByFixRow(f, fixFrontendRowIndex)
}

func findFixMiddleRow(f *Fighter) []*Fighter {
	return findTargetByFixRow(f, fixMiddleRowIndex)
}

func findFixBackendRow(f *Fighter) []*Fighter {
	return findTargetByFixRow(f, fixBackendRowIndex)
}
