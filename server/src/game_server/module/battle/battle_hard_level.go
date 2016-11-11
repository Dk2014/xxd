package battle

/*
import (
	"core/time"

	"game_server/battle"
	"game_server/mdb"
	"game_server/module"
)

func initHardLevelEnemy(state *module.SessionState, defendSide *battle.SideInfo) {
	var hardLevelState *mdb.PlayerHardLevelState
	state.Database.Select.PlayerHardLevelState(func(row *mdb.PlayerHardLevelStateRow) {
		//查找出这个关卡的状态
		if row.MissionLevelId() == state.MissionLevelState.LevelId {
			hardLevelState = row.GoObject()
			row.Break()
		}
	})
	if hardLevelState != nil && time.IsToday(hardLevelState.BattleTime) {
		//从上次失败的状态设置defend side怪物血量
		loadHardLevelEnemyHealth(defendSide, hardLevelState)
	}
}

//把side里面怪物血量存放到hardLevelState里面
func dumpHardLevelEnemyHealth(side *battle.SideInfo, hardLevelState *mdb.PlayerHardLevelState) {
	hardLevelState.Pos1 = 0
	hardLevelState.Pos2 = 0
	hardLevelState.Pos3 = 0
	hardLevelState.Pos4 = 0
	hardLevelState.Pos5 = 0
	hardLevelState.Pos6 = 0
	hardLevelState.Pos7 = 0
	hardLevelState.Pos8 = 0
	hardLevelState.Pos9 = 0
	hardLevelState.Pos10 = 0
	hardLevelState.Pos11 = 0
	hardLevelState.Pos12 = 0
	hardLevelState.Pos13 = 0
	hardLevelState.Pos14 = 0
	hardLevelState.Pos15 = 0
	for _, fighter := range side.Fighters {
		if fighter == nil {
			continue
		}
		health := int32(fighter.Health)
		switch fighter.Position {
		case 1:
			hardLevelState.Pos1 = health
		case 2:
			hardLevelState.Pos2 = health
		case 3:
			hardLevelState.Pos3 = health
		case 4:
			hardLevelState.Pos4 = health
		case 5:
			hardLevelState.Pos5 = health
		case 6:
			hardLevelState.Pos6 = health
		case 7:
			hardLevelState.Pos7 = health
		case 8:
			hardLevelState.Pos8 = health
		case 9:
			hardLevelState.Pos9 = health
		case 10:
			hardLevelState.Pos10 = health
		case 11:
			hardLevelState.Pos11 = health
		case 12:
			hardLevelState.Pos12 = health
		case 13:
			hardLevelState.Pos13 = health
		case 14:
			hardLevelState.Pos14 = health
		case 15:
			hardLevelState.Pos15 = health
		default:
			panic("undefine fighter position")
		}
	}
}

//把hardLevelState里面怪物血量存放到side里面
func loadHardLevelEnemyHealth(side *battle.SideInfo, hardLevelState *mdb.PlayerHardLevelState) {
	for i, fighter := range side.Fighters {
		if fighter == nil {
			continue
		}
		var health int32
		switch fighter.Position {
		case 1:
			health = hardLevelState.Pos1
		case 2:
			health = hardLevelState.Pos2
		case 3:
			health = hardLevelState.Pos3
		case 4:
			health = hardLevelState.Pos4
		case 5:
			health = hardLevelState.Pos5
		case 6:
			health = hardLevelState.Pos6
		case 7:
			health = hardLevelState.Pos7
		case 8:
			health = hardLevelState.Pos8
		case 9:
			health = hardLevelState.Pos9
		case 10:
			health = hardLevelState.Pos10
		case 11:
			health = hardLevelState.Pos11
		case 12:
			health = hardLevelState.Pos12
		case 13:
			health = hardLevelState.Pos13
		case 14:
			health = hardLevelState.Pos14
		case 15:
			health = hardLevelState.Pos15
		default:
			panic("undefine emeny position")
		}
		if health <= 0 {
			side.Fighters[i] = nil
		} else {
			side.Fighters[i].Health = int(health)
		}
	}
}
*/
