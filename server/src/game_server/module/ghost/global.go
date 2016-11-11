package ghost

import (
	"game_server/dat/ghost_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/tlog"
	"game_server/xdlog"
)

func init() {
	module.Ghost = GhostMod{}
}

type GhostMod struct {
}

// 魂侍功能开放初始化
func (this GhostMod) OpenFuncForGhost(db *mdb.Database) {
	// 初始化用户装备
	db.Select.PlayerRole(func(row *mdb.PlayerRoleRow) {
		db.Insert.PlayerGhostEquipment(&mdb.PlayerGhostEquipment{
			Pid:    db.PlayerId(),
			RoleId: row.RoleId(),
		})
	})

	//初始给一个勇气之卫魂侍
	addGhost(db, GHOST_YONGQIZHIWEI, tlog.IFR_OPEN_FUNC, xdlog.ET_OPEN_FUNC)
	db.Insert.PlayerGhostState(&mdb.PlayerGhostState{
		Pid: db.PlayerId(),
	})

}

// 角色魂侍装备表初始化
func (this GhostMod) InitRoleGhostEquipment(db *mdb.Database, roleId int8) {
	db.Insert.PlayerGhostEquipment(&mdb.PlayerGhostEquipment{
		Pid:    db.PlayerId(),
		RoleId: roleId,
	})
}

//添加魂侍
func (this GhostMod) AddGhost(state *module.SessionState, ghostId int16, itemFlowReason, xdEventType int32) {
	addGhost(state.Database, ghostId, itemFlowReason, xdEventType)
}

// 获取魂侍加成数据
func (this GhostMod) GetGhostAddData(playerGhost *mdb.PlayerGhost) (ghostAddData *ghost_dat.GhostAddData) {
	return getGhostAddData(playerGhost)
}

// 获得魂侍护盾比例值
//设置魂侍等级
func (this GhostMod) SetLevel(db *mdb.Database, playerGhostId int64, ghostLevel int16) {
	setlevel(db, playerGhostId, ghostLevel)
}
