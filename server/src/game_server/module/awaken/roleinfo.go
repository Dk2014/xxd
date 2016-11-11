package awaken

import (
	"core/fail"
	"core/net"
	"game_server/api/protocol/awaken_api"
	"game_server/dat/awaken_dat"
	"game_server/dat/player_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/tlog"
	"game_server/xdlog"
)

func awakenInfo(state *module.SessionState, role_id int8) *awaken_api.AwakenInfo_Out {
	out := &awaken_api.AwakenInfo_Out{}
	db := state.Database

	db.Select.PlayerAwakenGraphic(func(row *mdb.PlayerAwakenGraphicRow) {
		if row.RoleId() == role_id {
			out.Attrs = append(out.Attrs, awaken_api.AwakenInfo_Out_Attrs{
				AttrImpl: row.AttrImpl(),
				Level:    row.Level(),
			})
		}
	})

	return out
}

func levelupAttr(session *net.Session, role_id int8, attr_impl int16) bool {
	state := module.State(session)
	db := state.Database

	fail.When(!module.Player.IsOpenFunc(db, player_dat.FUNC_AWAKEN), "觉醒功能未开启")

	attr := seekAttr(db, role_id, attr_impl)

	if attr != nil {
		// 觉醒属性已存在

		// 检查等级
		if awaken_dat.IsSkillAttr(role_id, attr_impl) {
			//fail.When(attr.Level >= 1, "already learned this awaken skill")
			if attr.Level >= 1 {
				return false
			}
		} else {
			//fail.When(attr.Level >= awaken_dat.ATTR_LEVEL_LIMIT, "awaken attribute is already the maximum level")
			if attr.Level >= awaken_dat.ATTR_LEVEL_LIMIT {
				return false
			}
		}

		// 扣除希望之光
		attr_dat := awaken_dat.GetRoleAttr(role_id, attr_impl)
		if !module.Item.CheckItemNum(state, awaken_dat.ITEM_HOPES_LIGHT, int16(attr_dat.Lights)) {
			return false
		}
		module.Item.DelItemByItemId(state.Database, awaken_dat.ITEM_HOPES_LIGHT, int16(attr_dat.Lights), tlog.IFR_AWAKEN, xdlog.ET_ROLE_AWAKEN)

		// 升级
		attr.Level++
		db.Update.PlayerAwakenGraphic(attr)
	} else {
		// 新觉醒属性

		// 检查依赖
		dep_id := awaken_dat.GetDepImplId(role_id, attr_impl)
		if dep_id > 0 {
			dep := seekAttr(db, role_id, dep_id)
			fail.When(dep == nil || dep.Level <= 0, "broken awaken attribute dependency")
		}

		// 扣除希望之光
		attr_dat := awaken_dat.GetRoleAttr(role_id, attr_impl)
		if !module.Item.CheckItemNum(state, awaken_dat.ITEM_HOPES_LIGHT, int16(attr_dat.Lights)) {
			return false
		}
		module.Item.DelItemByItemId(state.Database, awaken_dat.ITEM_HOPES_LIGHT, int16(attr_dat.Lights), tlog.IFR_AWAKEN, xdlog.ET_ROLE_AWAKEN)

		if attr_dat.IsSkill {
			skill_id := attr_dat.SkillId
			//学习技能
			db.Insert.PlayerSkill(&mdb.PlayerSkill{
				Pid:        state.PlayerId,
				RoleId:     role_id,
				SkillId:    skill_id,
				SkillTrnlv: 1,
			})
			//通知玩家习得技能
			module.Notify.SendSkillAdd(session, role_id, skill_id)
		}

		// 学习
		attr = &mdb.PlayerAwakenGraphic{
			Pid:      state.PlayerId,
			RoleId:   role_id,
			AttrImpl: attr_impl,
			Level:    1,
		}
		db.Insert.PlayerAwakenGraphic(attr)
	}
	return true
}

// 查找角色觉醒属性
func seekAttr(db *mdb.Database, role_id int8, attr_impl int16) *mdb.PlayerAwakenGraphic {
	var attr *mdb.PlayerAwakenGraphic
	db.Select.PlayerAwakenGraphic(func(row *mdb.PlayerAwakenGraphicRow) {
		if row.RoleId() == role_id && row.AttrImpl() == attr_impl {
			attr = row.GoObject()
			row.Break()
		}
	})
	return attr
}
