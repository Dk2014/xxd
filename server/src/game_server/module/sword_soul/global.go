package sword_soul

import (
	"core/net"
	"core/time"
	"game_server/dat/player_dat"
	"game_server/dat/sword_soul_dat"
	"game_server/mdb"
	"game_server/module"
)

func init() {
	module.SwordSoul = SwordSoulMod{}
}

type SwordSoulMod struct {
}

func (mod SwordSoulMod) OpenFuncForSwordSoul(db *mdb.Database) {

	playerId := db.PlayerId()

	if db.Lookup.PlayerSwordSoulState(playerId) != nil {
		return
	}

	// 初始化玩家剑心状态
	db.Insert.PlayerSwordSoulState(&mdb.PlayerSwordSoulState{
		Pid:        playerId,
		BoxState:   1,
		Num:        sword_soul_dat.MAX_DRAW_NUM,
		UpdateTime: time.GetNowTime(),
	})

	// 初始角色化装备位
	db.Select.PlayerRole(func(row *mdb.PlayerRoleRow) {
		db.Insert.PlayerSwordSoulEquipment(&mdb.PlayerSwordSoulEquipment{
			Pid:    playerId,
			RoleId: row.RoleId(),
		})
	})
}

// 初始化角色剑心装备数据
func (mod SwordSoulMod) InitRoleSwordSoulEquipment(state *module.SessionState, roleId int8) {

	db := state.Database
	playerId := state.PlayerId

	db.Select.PlayerSwordSoulEquipment(func(row *mdb.PlayerSwordSoulEquipmentRow) {
		if row.RoleId() == roleId {
			row.Break()
			return
		}
	})

	db.Insert.PlayerSwordSoulEquipment(&mdb.PlayerSwordSoulEquipment{
		Pid:    playerId,
		RoleId: roleId,
	})
	return
}

// 增加剑心，放入背包
func (mod SwordSoulMod) AddSwordSoul(state *module.SessionState, swordSoulId int16, itemReason int32) {
	addSwordSoul(state, swordSoulId, itemReason)
}

//登录时如果拔剑次数满则发送通知
func (mod SwordSoulMod) SendSwordSoulWhenMaxDrawNum(session *net.Session) {
	state := module.State(session)
	if !module.Player.IsOpenFunc(state.Database, player_dat.FUNC_SWORD_SOUL) {
		return
	}
	playerSwordSoulState := state.Database.Lookup.PlayerSwordSoulState(state.PlayerId)
	updateSwordSoulNum(state, playerSwordSoulState)
	if playerSwordSoulState.Num == sword_soul_dat.MAX_DRAW_NUM {
		module.Notify.SendSwordSoulDrawNumChange(session, playerSwordSoulState.Num, 0)
	}
}

//设置剑心等级

func (mod SwordSoulMod) SetSwordSoulLevel(db *mdb.Database, roleid int8, swordpos int8, swordlevel int8) {
	setswordsoullevel(db, roleid, swordpos, swordlevel)
}
