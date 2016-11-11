package role

import (
	"core/fail"
	"core/net"
	"game_server/api/protocol/role_api"
	"game_server/dat/event_dat"
	"game_server/dat/player_dat"
	"game_server/dat/role_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/module/rpc"
	"game_server/tlog"
	"game_server/xdlog"
)

func init() {
	role_api.SetInHandler(RoleAPI{})
}

type RoleAPI struct{}

// 获取所有角色
func (api RoleAPI) GetAllRoles(session *net.Session, in *role_api.GetAllRoles_In) {
	out := &role_api.GetAllRoles_Out{}
	state := module.State(session)

	roleIds := module.Role.GetFormRoleId(state)
	inFormSet := make(map[int8]bool)
	for _, rid := range roleIds {
		if !role_dat.IsMainRole(rid) {
			inFormSet[rid] = true
		}
	}

	state.Database.Select.PlayerRole(func(row *mdb.PlayerRoleRow) {
		out.Roles = append(out.Roles, role_api.GetAllRoles_Out_Roles{
			RoleId:          row.RoleId(),
			Level:           row.Level(),
			Exp:             row.Exp(),
			FriendshipLevel: int16(row.FriendshipLevel()),
			InForm:          role_dat.IsMainRole(row.RoleId()) || inFormSet[row.RoleId()],
			Status:          int8(row.Status()),
			CoopId:          row.CoopId(),
		})
	})

	session.Send(out)
}

func (apt RoleAPI) GetRoleFightNum(session *net.Session, in *role_api.GetRoleFightNum_In) {
	state := module.State(session)
	out := &role_api.GetRoleFightNum_Out{}
	state.Database.Select.PlayerRole(func(row *mdb.PlayerRoleRow) {
		if in.RoleId > 0 {
			if row.RoleId() == in.RoleId {
				_, _, totalFighterNum := getFightNum(state.Database, row.RoleId(), false, role_dat.FIGHT_FOR_ALL)
				out.FightNums = append(out.FightNums, role_api.GetRoleFightNum_Out_FightNums{
					RoleId:   row.RoleId(),
					FightNum: totalFighterNum,
				})
				row.Break()
			}
		} else {
			_, _, totalFighterNum := getFightNum(state.Database, row.RoleId(), false, role_dat.FIGHT_FOR_ALL)
			out.FightNums = append(out.FightNums, role_api.GetRoleFightNum_Out_FightNums{
				RoleId:   row.RoleId(),
				FightNum: totalFighterNum,
			})
		}
	})
	session.Send(out)
}

func (api RoleAPI) GetRoleInfo(session *net.Session, in *role_api.GetRoleInfo_In) {
	state := module.State(session)

	role := findRoleByRoleId(state.Database, in.RoleId, false)
	fail.When(role == nil, "can't find role in db")

	fighter, _, totalFighterNum := getFightNum(state.Database, role.RoleId, false, role_dat.FIGHT_FOR_ALL)

	hitLevel, critialLevel, sleepLevel, dizzinessLevel, randomLevel, disableSkillLevel, poisoningLevel, blockLevel, destroyLevel, critialHurtLevel, tenacityLevel, dodgeLevel := calcFighterLevel(fighter)

	session.Send(&role_api.GetRoleInfo_Out{
		RoleId:          role.RoleId,
		Level:           role.Level,
		Exp:             role.Exp,
		FriendshipLevel: int16(role.FriendshipLevel), // 羁绊等级
		FightNum:        totalFighterNum,             //  战力
		Status:          int8(role.Status),           //伙伴状态
		CoopId:          role.CoopId,

		// 基础属性
		Attack:      int32(fighter.Attack),
		Defence:     int32(fighter.Defend), // 防御
		Health:      int32(fighter.MaxHealth),
		Speed:       int32(fighter.Speed),
		Cultivation: int32(fighter.Cultivation),    // 内力
		Sunder:      int32(fighter.SunderMaxValue), // 护甲值
		GhostPower:  int32(fighter.GetGhostPower()),

		// 概率属性
		HitLevel:      int32(fighter.HitLevel + hitLevel),
		CriticalLevel: int32(fighter.CritialLevel + critialLevel),

		SleepLevel:        int32(fighter.SleepLevel + sleepLevel),
		DizzinessLevel:    int32(fighter.DizzinessLevel + dizzinessLevel),
		RandomLevel:       int32(fighter.RandomLevel + randomLevel),
		DisableSkillLevel: int32(fighter.DisableSkillLevel + disableSkillLevel),
		PoisoningLevel:    int32(fighter.PoisoningLevel + poisoningLevel),

		BlockLevel:        int32(fighter.BlockLevel + blockLevel),
		DestroyLevel:      int32(fighter.DestroyLevel + destroyLevel),
		CriticalHurtLevel: int32(fighter.CritialHurtLevel + critialHurtLevel),
		TenacityLevel:     int32(fighter.TenacityLevel + tenacityLevel),
		DodgeLevel:        int32(fighter.DodgeLevel + dodgeLevel),
	})
}

func (api RoleAPI) GetPlayerInfo(session *net.Session, in *role_api.GetPlayerInfo_In) {
	rpc.RemoteGetInfo(module.State(session), in.Pid, func(Reply *rpc.Reply_GetInfo, err error) {
		fail.When(err != nil, err)
		session.Send(Reply.PlayerInfo)
	})
}

func (api RoleAPI) GetFightNum(session *net.Session, in *role_api.GetFightNum_In) {
	state := module.State(session)
	_, fightNums, totalFight := getFightNum(state.Database, 0, true, in.FightType)
	out := &role_api.GetFightNum_Out{}
	out.FightNums = []role_api.GetFightNum_Out_FightNums{}
	for k, v := range fightNums {
		out.FightNums = append(out.FightNums, role_api.GetFightNum_Out_FightNums{
			FightType: k,
			FightNum:  v,
		})
	}

	session.Send(out)

	if in.FightType == module.FIGHT_FOR_ALL {
		eventInfo, _ := event_dat.GetEventInfoById(event_dat.EVENT_STRONG_AWARD)
		if event_dat.CheckEventTime(eventInfo, event_dat.NOT_END) {
			//更新最大可领取状态
			if power := event_dat.GetEventFightAwardPower(totalFight); power > 0 {
				state.EventsState.UpdateMax(state.Database, event_dat.EVENT_STRONG_AWARD, power)
			}
		}
	}

	if in.FightType == module.FIGHT_FOR_ALL {
		tlog.PlayerFightNumFlowLog(state.Database, int32(totalFight))
	}
}

func (api RoleAPI) GetPlayerInfoWithOpenid(session *net.Session, in *role_api.GetPlayerInfoWithOpenid_In) {
	state := module.State(session)
	fail.When(state.PlayerId == 0, "未登陆,不能调用GetPlayerInfoWithOpenid接口")
	rpc.RemoteGetInfoWithOpenId(state, in.Openid, int(in.GameServerId), func(Reply *rpc.Reply_GetInfoWithOpenId, err error) {
		if err == nil && Reply.PlayerInfo != nil {
			session.Send(Reply.PlayerInfo)
		}
	})
}

func (api RoleAPI) LevelupRoleFriendship(session *net.Session, in *role_api.LevelupRoleFriendship_In) {
	state := module.State(session)
	db := state.Database
	role_id := in.RoleId

	// 检查开启等级
	fail.When(!module.Player.IsOpenFunc(db, player_dat.FUNC_FRIENDSHIP), "未开启伙伴羁绊")

	role := findRoleByRoleId(db, role_id, true)
	friendshipStuff := role_dat.GetRoleFriendship(role_id, int16(role.FriendshipLevel+1))

	// 检查羁绊等级是否达到最大
	fail.When(role.FriendshipLevel >= role_dat.MAX_FRISHP_LEVEL, "the friendship level of this role has been up to top")
	// 检查角色等级是否满足羁绊升级要求
	fail.When(role.Level < friendshipStuff.RequiredRoleLevel, "role level is not enough to levelup friendship level")

	// 尝试扣除喜好品（也不知道能不能成功，好担心啊）
	module.Item.DelItemByItemId(db, friendshipStuff.FavouriteItem, int16(friendshipStuff.FavouriteCount), tlog.IFR_FRISHP_LEVELUP, xdlog.ET_FRIEND_SHIP)

	role.FriendshipLevel++
	db.Update.PlayerRole(role)

	// 检查是否可学习新技能
	player_fame := db.Lookup.PlayerFame(db.PlayerId())
	module.Skill.UpdateSkill(db, role_id, int16(role.FriendshipLevel), player_fame.Level, role.Level)

	out := &role_api.LevelupRoleFriendship_Out{}
	session.Send(out)
}

func (api RoleAPI) RecruitBuddy(session *net.Session, in *role_api.RecruitBuddy_In) {
	state := module.State(session)
	status := recruitBuddy(state, in.RoleId)
	session.Send(&role_api.RecruitBuddy_Out{
		Result: status,
	})
}

//change_role_status
func (api RoleAPI) ChangeRoleStatus(session *net.Session, in *role_api.ChangeRoleStatus_In) {
	state := module.State(session)
	result := changeRoleStatus(state.Database, in.RoleId, in.Status)
	session.Send(&role_api.ChangeRoleStatus_Out{
		Result: result,
		RoleId: in.RoleId,
		Status: in.Status,
	})
}

// get_inn_role_list
func (api RoleAPI) GetInnRoleList(session *net.Session, in *role_api.GetInnRoleList_In) {
	state := module.State(session)
	buddyList := getInnRoleList(state.Database)
	session.Send(&role_api.GetInnRoleList_Out{
		RoleList: buddyList,
	})
}

func (api RoleAPI) BuddyCoop(session *net.Session, in *role_api.BuddyCoop_In) {
	state := module.State(session)
	buddyCoop(state.Database, in.CoopId)
}
