package module

import (
	//"core/fail"
	"game_server/battle"
	"math/rand"
)

const (
	BATTLE_PET_FRONT_ROW    = 1 // 前排
	BATTLE_PET_BACK_ROW     = 2 // 后排
	BATTLE_PET_LEFT_COLUMN  = 3 // 左侧
	BATTLE_PET_RIGHT_COLUMN = 4 // 右侧
)

var deployFrontRow = []int{
	1, 2, 3, 4, 5,
	6, 7, 8, 9, 10,
	12, 13, 14, 15, 11,
}

var deployBackRow = []int{
	12, 13, 14,
	6, 7, 8, 9, 10,
	1, 2, 3, 4, 5,
	11, 15,
}

var deployLeftColumn = []int{
	1, 6,
	2, 7, 12,
	3, 8, 13,
	4, 9, 14,
	5, 10, 11, 15,
}

var deployRightColumn = []int{
	5, 10,
	4, 9, 14,
	3, 8, 13,
	2, 7, 12,
	1, 6, 15, 11,
}

func GetBattlePetPos(petPosType int8, fighter []*battle.Fighter) (pos int) {
	var deploy []int
	var max int

	switch petPosType {
	case BATTLE_PET_FRONT_ROW:
		max = 5
		deploy = deployFrontRow

	case BATTLE_PET_BACK_ROW:
		max = 3
		deploy = deployBackRow

	case BATTLE_PET_LEFT_COLUMN:
		max = 2
		deploy = deployLeftColumn

	case BATTLE_PET_RIGHT_COLUMN:
		max = 2
		deploy = deployRightColumn
	}

	idx, num := 0, max
	randList := make([]int, max)
	for _, pos = range deploy {
		if fighter[pos-1] != nil &&
			(!fighter[pos-1].IsBattlePet || fighter[pos-1].Health > 0) {
			num -= 1
			continue
		}

		randList[idx] = pos
		if idx++; idx >= max {
			break
		}
	}

	// 灵宠首选位置已经满了
	if num < 1 {
		num = max
	}

	idx = rand.Intn(num)
	if pos = randList[idx]; pos > 0 {
		return
	}

	//多人关卡内可能A B 两个玩家同事召唤灵宠，不过其中一个召唤失败
	//fail.When(true, "can't get pos for battle pet")
	return -1
}
