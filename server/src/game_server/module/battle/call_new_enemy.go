package battle

import (
	"core/fail"
	"game_server/api/protocol/battle_api"
	"game_server/battle"
	"game_server/module"
	"math/rand"
)

var indexArr [15]int

func initFree(fighters []*battle.Fighter) {
	for i := 0; i < 15; i++ {

		if fighters[i] == nil || fighters[i].Health <= 0 {
			indexArr[i] = 1
		} else {
			indexArr[i] = 0
		}

	}
}
func preCall(callInfos []battle.CallInfo, fighters []*battle.Fighter) []battle.CallInfo {
	initFree(fighters)

	result := make([]battle.CallInfo, 0)
	conflict_arr := make([]int, 0)           //开头固定位置的起冲突的怪物索引
	for index, callInfo := range callInfos { //callinfos中所有固定位置召唤的都在开头
		if callInfo.Position < 0 { //随机位置的操作
			position := callInfo.Position
			new_position := -1
			switch position {
			case FRONT:
				new_position = newLocateEnemy(5, FrontArr)
			case BEHIND:
				new_position = newLocateEnemy(5, BehindArr)
			case LEFT:
				new_position = newLocateEnemy(3, LeftArr)
			case RIGHT:
				new_position = newLocateEnemy(3, RightArr)
			case ALL:
				new_position = newLocateEnemy(15, FrontArr)
			default:
				fail.When(true, "wront range position")
			}

			if new_position < 0 {
				break //随机得到的位置小于0 表明位置已满，无需继续下去
			}
			result = append(result, battle.CallInfo{callInfo.Enemy, int8(new_position)})
		} else {
			if indexArr[int(callInfo.Position)] == 0 { //固定位置且起冲突了
				conflict_arr = append(conflict_arr, index)
				callInfos[index].Position = FRONT
			} else { //固定位置，没有起冲突的
				indexArr[int(callInfo.Position)] = 0
				result = append(result, battle.CallInfo{callInfo.Enemy, int8(callInfo.Position)})
			}
		}
	}

	if len(conflict_arr) > 0 {

		for _, index := range conflict_arr {
			new_position := -1
			new_position = newLocateEnemy(5, FrontArr)
			if new_position < 0 { //位置已满，无需再继续找下去
				break
			} else {
				indexArr[new_position] = 0
			}
			result = append(result, battle.CallInfo{callInfos[index].Enemy, int8(new_position)})
		}
	}
	return result
}

func CallNewEnemyHelper(battleState *battle.BattleState, result *battle.FightResult) (*battle_api.CallNewEnemys_Out, bool) {
	callEnemysResponse := &battle_api.CallNewEnemys_Out{}
	fighters := battleState.Defenders.Fighters
	for index, calls := range result.CallEnemys {
		//enemyId := int32(enemy.Enemy)
		//position := int(enemy.Position)
		//fail.When(fighters[position] != nil && fighters[position].Health > 0, "already have enemy where to load new enemy")
		//newFighter := module.NewFighterForEnemy(enemyId, position+1)
		//battleState.RuntimeAddFighter(battle.FT_DEF, newFighter)
		real_call_infos := CallEnemys(battleState, fighters, calls)
		callEnemysResponse.CallInfo = append(callEnemysResponse.CallInfo, battle_api.CallNewEnemys_Out_CallInfo{
			Ftype:       int8(result.Type),
			Position:    int8(result.FighterPos),
			AttackIndex: int8(index),
			Enemys:      real_call_infos,
		})
	}
	return callEnemysResponse, len(callEnemysResponse.CallInfo) > 0
}

//召唤新怪上阵
func CallEnemys(battleState *battle.BattleState, fighters []*battle.Fighter, callInfos []battle.CallInfo) []battle_api.CallNewEnemys_Out_CallInfo_Enemys {
	//status := module.State(session)
	//fighters := status.Battle.GetBattle().Defenders.Fighters
	result := make([]battle_api.CallNewEnemys_Out_CallInfo_Enemys, 0)
	realInfos := preCall(callInfos, fighters)
	//执行召唤
	for _, realInfo := range realInfos {
		if realInfo.Position >= 0 && realInfo.Enemy > 0 {
			enemyId := int32(realInfo.Enemy)
			position := int(realInfo.Position)
			fail.When(fighters[position] != nil && fighters[position].Health > 0, "already have enemy where to load new enemy")
			//如果需要为怪物进行属性加成，请修改NewFighterForEnemy第二个参数
			newFighter := module.NewFighterForEnemy(enemyId, 0, position+1)
			battleState.RuntimeAddFighter(battle.FT_DEF, newFighter)
			result = append(result, tarCallInfo(newFighter))
		}
	}
	return result
}

//const
const (
	FRONT  = -2
	BEHIND = -1
	LEFT   = -4
	RIGHT  = -3
	ALL    = -5
)

var FrontArr []int = []int{
	1, 2, 3, 4, 5,
	6, 7, 8, 9, 10,
	11, 12, 13, 14, 15,
}

var BehindArr []int = []int{
	11, 12, 13, 14, 15,
	6, 7, 8, 9, 10,
	1, 2, 3, 4, 5,
}

var LeftArr []int = []int{
	1, 6, 11,
	2, 7, 12,
	3, 8, 13,
	4, 9, 14,
	5, 10, 15,
}

var RightArr []int = []int{
	5, 10, 15,
	4, 9, 14,
	3, 8, 13,
	2, 7, 12,
	1, 6, 11,
}

func newLocateEnemy(max int, deploy []int) int {
	idx, num := 0, max
	var pos int
	randList := make([]int, max)
	for _, pos = range deploy {
		if indexArr[pos-1] == 0 {
			num -= 1
			continue
		}

		randList[idx] = pos
		if idx++; idx >= max {
			break
		}
	}
	if num < 1 {
		num = max
	}
	idx = rand.Intn(num)
	if pos = randList[idx]; pos > 0 {
		indexArr[pos-1] = 0
		return pos - 1
	}
	//没有空位了，返回－1
	return -1

}

func tarCallInfo(v *battle.Fighter) (result battle_api.CallNewEnemys_Out_CallInfo_Enemys) {
	skills := make([]battle_api.CallNewEnemys_Out_CallInfo_Enemys_Skills, len(v.Skills))

	for i, skill := range v.Skills {
		if skill == nil {
			skills[i] = battle_api.CallNewEnemys_Out_CallInfo_Enemys_Skills{
				Skill: battle_api.SkillInfo{
					SkillId: -2,
				},
			}
			continue
		}

		skills[i] = battle_api.CallNewEnemys_Out_CallInfo_Enemys_Skills{
			Skill: battle_api.SkillInfo{
				SkillId:  int16(skill.SkillId),
				IncPower: int8(skill.IncPower),
				DecPower: int8(skill.DecPower),
			},
			SkillId2: int16(v.SkillInfos[i].SkillId2),
		}
	}
	result = battle_api.CallNewEnemys_Out_CallInfo_Enemys{
		Role: battle_api.BattleRole{
			Kind:                battle_api.FIGHTER_KIND_ENEMY,
			PlayerId:            v.PlayerId,
			RoleId:              int32(v.RoleId),
			RoleLevel:           int16(v.Level),
			Position:            int32(v.Position),
			FashionId:           v.FashionId,
			Health:              int32(v.Health),
			MaxHealth:           int32(v.MaxHealth),
			Power:               int16(v.Power),
			MaxPower:            int16(v.MaxPower),
			SunderValue:         int16(v.GetSunderValue()),
			SunderMaxValue:      int16(v.SunderMaxValue),
			SunderMinHurtRate:   int16(v.SunderMinHurtRate),
			SunderEndHurtRate:   int16(v.SunderEndHurtRate),
			SunderEndDefendRate: int16(v.SunderEndDefendRate),

			Speed: int32(v.Speed),

			GhostShieldValue:  int32(v.GhostShieldValue),
			CouldUseSwordSoul: v.SwordSoulValue > 0,
		},
		Skills: skills,
	}
	return
}

func commomSummonHandlerFacotry(pids ...int64) func(*battle.BattleState, *battle.FightResult) bool {
	return func(battleState *battle.BattleState, result *battle.FightResult) bool {
		resp, ok := CallNewEnemyHelper(battleState, result)
		if !ok {
			return false
		}
		for _, pid := range pids {
			if session, online := module.Player.GetPlayerOnline(pid); online {
				session.Send(resp)
			}
		}
		return true

	}
}
