package role

import (
	"core/fail"
	"game_server/api/protocol/role_api"
	"game_server/dat/item_dat"
	"game_server/dat/player_dat"
	"game_server/mdb"
	"game_server/module"
)

func init() {
	module.Role = RoleMod{}
}

type RoleMod struct{}

// 给角色加经验
func (roleMod RoleMod) AddRoleExp(db *mdb.Database, roleId int8, addExp int64, mainRoleId int8, playerExpFlowReason int32) {
	addRoleExp(db, roleId, addExp, mainRoleId, playerExpFlowReason)
}

// 获取主角数据
func (roleMod RoleMod) GetMainRole(db *mdb.Database) *mdb.PlayerRole {
	player := db.Lookup.Player(db.PlayerId())
	return db.Lookup.PlayerRole(player.MainRoleId)
}

// 获取伙伴角色数据
func (roleMod RoleMod) GetBuddyRole(db *mdb.Database, roleId int8) *mdb.PlayerRole {
	return findRoleByRoleId(db, roleId, false)
}

// 获取在队伍中的伙伴
func (roleMod RoleMod) GetBuddyRoleInTeam(db *mdb.Database, roleId int8) *mdb.PlayerRole {
	return findRoleByRoleId(db, roleId, true)
}

// 添加伙伴
func (roleMod RoleMod) AddBuddyRole(state *module.SessionState, roleId int8, level int16) (status int8) {
	module.Quest.RefreshQuestForNewBuddy(state.Database, roleId)
	status = addNewBuddyRole(state, roleId, level)

	// 通知玩家增加新伙伴
	if s, ok := module.Player.GetPlayerOnline(state.PlayerId); ok {
		module.Notify.SendNewBuddy(s, roleId, level)
	}

	//for sealedbook
	if _, exist := state.GetSealedBookRecord().FindRecord(item_dat.STEALDBOOK_TYPE_ROLES, int16(roleId)); !exist {
		sealedbook := item_dat.GetSealedBookInfo(item_dat.STEALDBOOK_TYPE_ROLES, int16(roleId))
		if sealedbook != nil {
			state.GetSealedBookRecord().AddRecord(item_dat.STEALDBOOK_TYPE_ROLES, int16(roleId), item_dat.STEALDBOOK_HAVING, state.Database)
		}
	}

	return
}

//获取上阵角色ID
func (roleMod RoleMod) GetFormRoleId(state *module.SessionState) (roleIds []int8) {
	playerForm := state.Database.Lookup.PlayerFormation(state.PlayerId)
	if playerForm.Pos0 > 0 {
		roleIds = append(roleIds, playerForm.Pos0)
	}

	if playerForm.Pos1 > 0 {
		roleIds = append(roleIds, playerForm.Pos1)
	}

	if playerForm.Pos2 > 0 {
		roleIds = append(roleIds, playerForm.Pos2)
	}
	if playerForm.Pos3 > 0 {
		roleIds = append(roleIds, playerForm.Pos3)
	}
	if playerForm.Pos4 > 0 {
		roleIds = append(roleIds, playerForm.Pos4)
	}
	if playerForm.Pos5 > 0 {
		roleIds = append(roleIds, playerForm.Pos5)
	}
	if playerForm.Pos6 > 0 {
		roleIds = append(roleIds, playerForm.Pos6)
	}
	if playerForm.Pos7 > 0 {
		roleIds = append(roleIds, playerForm.Pos7)
	}
	if playerForm.Pos8 > 0 {
		roleIds = append(roleIds, playerForm.Pos8)
	}
	return roleIds
}

// 为上阵角色加经验（包括主角）
func (roleMod RoleMod) AddFormRoleExp(state *module.SessionState, addExp int64, playerExpFlowReason int32) {
	playerForm := state.Database.Lookup.PlayerFormation(state.PlayerId)

	// 先加主角经验，伙伴等级不能超过主角
	addRoleExp(state.Database, state.RoleId, addExp, state.RoleId, playerExpFlowReason)

	if playerForm.Pos0 > 0 && state.RoleId != playerForm.Pos0 {
		addRoleExp(state.Database, playerForm.Pos0, addExp, state.RoleId, playerExpFlowReason)
	}

	if playerForm.Pos1 > 0 && state.RoleId != playerForm.Pos1 {
		addRoleExp(state.Database, playerForm.Pos1, addExp, state.RoleId, playerExpFlowReason)
	}

	if playerForm.Pos2 > 0 && state.RoleId != playerForm.Pos2 {
		addRoleExp(state.Database, playerForm.Pos2, addExp, state.RoleId, playerExpFlowReason)
	}

	if playerForm.Pos3 > 0 && state.RoleId != playerForm.Pos3 {
		addRoleExp(state.Database, playerForm.Pos3, addExp, state.RoleId, playerExpFlowReason)
	}

	if playerForm.Pos4 > 0 && state.RoleId != playerForm.Pos4 {
		addRoleExp(state.Database, playerForm.Pos4, addExp, state.RoleId, playerExpFlowReason)
	}

	if playerForm.Pos5 > 0 && state.RoleId != playerForm.Pos5 {
		addRoleExp(state.Database, playerForm.Pos5, addExp, state.RoleId, playerExpFlowReason)
	}
	if playerForm.Pos6 > 0 && state.RoleId != playerForm.Pos6 {
		addRoleExp(state.Database, playerForm.Pos6, addExp, state.RoleId, playerExpFlowReason)
	}
	if playerForm.Pos7 > 0 && state.RoleId != playerForm.Pos7 {
		addRoleExp(state.Database, playerForm.Pos7, addExp, state.RoleId, playerExpFlowReason)
	}
	if playerForm.Pos8 > 0 && state.RoleId != playerForm.Pos8 {
		addRoleExp(state.Database, playerForm.Pos8, addExp, state.RoleId, playerExpFlowReason)
	}
}

// 返回新的level, exp, 返回参数为 level int16, exp int64
func (roleMod RoleMod) GetNewLevelAndExp(level int16, exp int64) (int16, int64) {
	return getNewLevelAndExp(level, exp)
}

func (roleMod RoleMod) GetOtherPlayerInfo(playerId int64, db *mdb.Database, playerInfo *role_api.PlayerInfo) {
	player := db.Lookup.Player(db.PlayerId())
	fail.When(player == nil, "GetPlayerInfo player nil")

	playerInfo.Openid = []byte(player.User)
	playerInfo.Pid = db.PlayerId()
	playerInfo.Name = []byte(player.Nick)
	playerInfo.Roles = []role_api.PlayerInfo_Roles{}

	playerFashionState := db.Lookup.PlayerFashionState(db.PlayerId())
	playerInfo.FashionId = playerFashionState.DressedFashionId

	if module.Player.IsOpenFunc(db, player_dat.FUNC_RAINBOW_LEVEL) {
		playerRainbow := db.Lookup.PlayerRainbowLevel(db.PlayerId())
		playerInfo.BestSegment = playerRainbow.BestSegment
		playerInfo.BestOrder = playerRainbow.BestOrder
		playerInfo.BestRecordTimestamp = playerRainbow.BestRecordTimestamp
	}

	PlayerFormation := db.Lookup.PlayerFormation(db.PlayerId())

	db.Select.PlayerRole(func(row *mdb.PlayerRoleRow) {
		role := role_api.PlayerInfo_Roles{}
		role.RoleId = row.RoleId()
		role.Level = row.Level()
		role.FriendshipLevel = int16(row.FriendshipLevel())
		role.IsDeploy = false
		role.Status = int8(row.Status())
		role.CoopId = row.CoopId()

		switch role.RoleId {
		case PlayerFormation.Pos0:
			role.IsDeploy = true
		case PlayerFormation.Pos1:
			role.IsDeploy = true
		case PlayerFormation.Pos2:
			role.IsDeploy = true
		case PlayerFormation.Pos3:
			role.IsDeploy = true
		case PlayerFormation.Pos4:
			role.IsDeploy = true
		case PlayerFormation.Pos5:
			role.IsDeploy = true
		case PlayerFormation.Pos6:
			role.IsDeploy = true
		case PlayerFormation.Pos7:
			role.IsDeploy = true
		case PlayerFormation.Pos8:
			role.IsDeploy = true
		}

		var playerEquipment *mdb.PlayerEquipment
		db.Select.PlayerEquipment(func(row2 *mdb.PlayerEquipmentRow) {
			if row2.RoleId() == role.RoleId {
				playerEquipment = row2.GoObject()
				row2.Break()
			}
		})

		role.Equips = []role_api.PlayerInfo_Roles_Equips{}

		setEqInfo := func(pos int, eq *mdb.PlayerItem, eqAppendix *mdb.PlayerItemAppendix) {
			if eqAppendix == nil {
				role.Equips = append(role.Equips, role_api.PlayerInfo_Roles_Equips{
					Pos:         int8(pos),
					ItemId:      eq.ItemId,
					RefineLevel: eq.RefineLevel,
				})
			} else {
				role.Equips = append(role.Equips, role_api.PlayerInfo_Roles_Equips{
					Pos:           int8(pos),
					ItemId:        eq.ItemId,
					RefineLevel:   eq.RefineLevel,
					Attack:        eqAppendix.Attack,
					Defence:       eqAppendix.Defence,
					Health:        eqAppendix.Health,
					Speed:         eqAppendix.Speed,
					Cultivation:   eqAppendix.Cultivation,
					HitLevel:      eqAppendix.HitLevel,
					CriticalLevel: eqAppendix.CriticalLevel,
					BlockLevel:    eqAppendix.BlockLevel,
					DestroyLevel:  eqAppendix.DestroyLevel,
					TenacityLevel: eqAppendix.TenacityLevel,
					DodgeLevel:    eqAppendix.DodgeLevel,
					RecastAttr:    role_api.Attribute(eqAppendix.RecastAttr),
				})
			}
		}

		if playerEquipment.WeaponId > 0 {
			eq := db.Lookup.PlayerItem(playerEquipment.WeaponId)
			if eq != nil {
				eqAppendix := db.Lookup.PlayerItemAppendix(eq.AppendixId)
				setEqInfo(item_dat.TYPE_WEAPON, eq, eqAppendix)
			}
		}

		if playerEquipment.ClothesId > 0 {
			eq := db.Lookup.PlayerItem(playerEquipment.ClothesId)
			if eq != nil {
				eqAppendix := db.Lookup.PlayerItemAppendix(eq.AppendixId)
				setEqInfo(item_dat.TYPE_CLOTHES, eq, eqAppendix)
			}
		}

		if playerEquipment.AccessoriesId > 0 {
			eq := db.Lookup.PlayerItem(playerEquipment.AccessoriesId)
			if eq != nil {
				eqAppendix := db.Lookup.PlayerItemAppendix(eq.AppendixId)
				setEqInfo(item_dat.TYPE_ACCESSORIES, eq, eqAppendix)
			}
		}

		if playerEquipment.ShoeId > 0 {
			eq := db.Lookup.PlayerItem(playerEquipment.ShoeId)
			if eq != nil {
				eqAppendix := db.Lookup.PlayerItemAppendix(eq.AppendixId)
				setEqInfo(item_dat.TYPE_SHOE, eq, eqAppendix)
			}
		}

		fighter, _, totalFighterNum := getFightNum(db, role.RoleId, false, module.FIGHT_FOR_ALL)

		hitLevel, critialLevel, sleepLevel, dizzinessLevel, randomLevel, disableSkillLevel, poisoningLevel, blockLevel, destroyLevel, critialHurtLevel, tenacityLevel, dodgeLevel := calcFighterLevel(fighter)

		role.FightNum = totalFighterNum
		role.Attack = int32(fighter.Attack)
		role.Defence = int32(fighter.Defend)
		role.Health = int32(fighter.MaxHealth)
		role.Speed = int32(fighter.Speed)
		role.Cultivation = int32(fighter.Cultivation)
		role.Sunder = int32(fighter.SunderMaxValue)

		role.HitLevel = int32(fighter.HitLevel + hitLevel)
		role.CriticalLevel = int32(fighter.CritialLevel + critialLevel)

		role.SleepLevel = int32(fighter.Sleep + sleepLevel)
		role.DizzinessLevel = int32(fighter.Dizziness + dizzinessLevel)
		role.RandomLevel = int32(fighter.Random + randomLevel)
		role.DisableSkillLevel = int32(fighter.DisableSkill + disableSkillLevel)
		role.PoisoningLevel = int32(fighter.Poisoning + poisoningLevel)

		role.BlockLevel = int32(fighter.BlockLevel + blockLevel)
		role.DestroyLevel = int32(fighter.DestroyLevel + destroyLevel)
		role.CriticalHurtLevel = int32(fighter.CritialHurtLevel + critialHurtLevel)
		role.TenacityLevel = int32(fighter.TenacityLevel + tenacityLevel)
		role.DodgeLevel = int32(fighter.DodgeLevel + dodgeLevel)
		playerInfo.Roles = append(playerInfo.Roles, role)
	})
}

func (roleMod RoleMod) GetFormFightNum(db *mdb.Database) (fightNum int32) {
	_, _, fightNum = getFightNum(db, 1, true, module.FIGHT_FOR_ALL)
	return
}

func (roleMod RoleMod) BreakBuddyCoop(db *mdb.Database, roleId int8) {
	breakBuddyCoop(db, roleId)
}
