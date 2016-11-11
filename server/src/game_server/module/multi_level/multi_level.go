package multi_level

import (
	"core/fail"
	"core/time"
	"encoding/json"
	"game_server/api/protocol/multi_level_api"
	"game_server/battle"
	"game_server/dat/multi_level_dat"
	"game_server/dat/player_dat"
	"game_server/dat/team_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/module/rpc"
	goTime "time"
)

func getInfo(state *module.SessionState) (int8, int32) {
	playerMultiLevel := state.Database.Lookup.PlayerMultiLevelInfo(state.PlayerId)
	fail.When(playerMultiLevel == nil, "multi-level function not open")

	if !time.IsInPointHour(player_dat.RESET_MULTI_LEVEL_TIMES_IN_HOUR, playerMultiLevel.BattleTime) {
		playerMultiLevel.DailyNum = 0
		state.Database.Update.PlayerMultiLevelInfo(playerMultiLevel)
	}

	return playerMultiLevel.DailyNum, playerMultiLevel.Lock
}

func getFormInfo(state *module.SessionState, rsp *multi_level_api.GetFormInfo_Out) {
	playerMultiLevel := state.Database.Lookup.PlayerMultiLevelInfo(state.PlayerId)
	rsp.BuddyRoleId = playerMultiLevel.BuddyRoleId
	rsp.BuddyRoleRow = playerMultiLevel.BuddyRow
	rsp.Tactical = playerMultiLevel.TacticalGrid
}

func changeBuddy(state *module.SessionState, roleId int8) bool {

	fail.When(roleId == state.RoleId, "wrong buddy id")
	rpc.RemoteMultiLevelChangeBuddy(state.PlayerId, roleId, func() {
		//如果玩家已经在多人关卡房间里面了需要发通知
		if state.MultiLevelRoomId > 0 {
			if room, ok := multiDataTable.getRoom(state.MultiLevelRoomId); ok {
				partner, ok := room.Partners[state.PlayerId]
				if ok {
					//加锁?
					partner.BuddyRoleId = roleId
					module.API.Broadcast(room.channel, &multi_level_api.NotifyUpdatePartner_Out{
						Partner: multi_level_api.PartnerInfo{
							Pid:         state.PlayerId,
							Nick:        partner.Nick,
							FashionId:   partner.FashionId,
							RoleId:      partner.MainRoleId,
							Level:       partner.Level,
							BuddyRoleId: partner.BuddyRoleId,
						},
					})
				}
			}
		}
	})
	return true
}

func changeForm(state *module.SessionState, buddyRow int8) bool {
	fail.When(buddyRow != multi_level_dat.MULTI_LEVEL_DEPLOY_ROW_1 && buddyRow != multi_level_dat.MULTI_LEVEL_DEPLOY_ROW_2, "[changeForm] error deploy row")
	rpc.RemoteMultiLevelChangeForm(state.PlayerId, buddyRow, func() {
		if state.MultiLevelRoomId > 0 {
			if room, ok := multiDataTable.getRoom(state.MultiLevelRoomId); ok {
				partner, ok := room.Partners[state.PlayerId]
				if ok {
					//加锁?
					partner.BuddyRoleRow = buddyRow

					if partner.BuddyRoleRow == 1 {
						partner.MainRoleRow = 2
					} else {
						partner.MainRoleRow = 1
					}
				}
			}
		}
	})
	return true
}

func enterChecker(state *module.SessionState, playerMultiLevel *mdb.PlayerMultiLevelInfo, levelLock int32, checkDailyNum bool) {
	fail.When(levelLock > playerMultiLevel.Lock, "multi-level not opened")

	cancelAutoMatch(state)

	if checkDailyNum {
		fail.When(playerMultiLevel.DailyNum >= multi_level_dat.MULTI_LEVEL_DAILY_NUM_MAX, "can't enter multi-level in today")
	}
}

func newPartnerInFormRoleInfos(partner *Partner) []*module.InFormRoleInfo {
	// 构造上阵角色信息
	inFormRoleInfos := []*module.InFormRoleInfo{
		/* 0 */
		&module.InFormRoleInfo{
			RoleId: partner.MainRoleId,
			Pos:    int((partner.MainRoleRow-1)*5 + partner.Col),
		},
	}
	if partner.BuddyRoleId != team_dat.POS_NO_ROLE {
		inFormRoleInfos = append(inFormRoleInfos, &module.InFormRoleInfo{
			RoleId: partner.BuddyRoleId,
			Pos:    int((partner.BuddyRoleRow-1)*5 + partner.Col),
		})
	}
	return inFormRoleInfos
}

func newBattle(room *MultiLevelRoom) {
	num := len(room.Partners)
	fail.When(num < 2, "[MultiLevel StartBattle] player to less")

	setcol2 := false

	var args []*rpc.Args_NewPlayerFighter
	var battleTotemInfo [5]*battle.TotemInfo
	for pid, partner := range room.Partners {
		arg := &rpc.Args_NewPlayerFighter{
			Pid:      pid,
			Tactical: partner.Tactical,
		}

		if pid == room.LeaderPid {
			partner.Col = 3 // 队长在中间
		} else {
			if !setcol2 {
				partner.Col = 2
				setcol2 = true
			} else {
				partner.Col = 1
			}
		}

		arg.AutoFight = true
		arg.FormRoleInfo = newPartnerInFormRoleInfos(partner)
		args = append(args, arg)
	}

	fighters := make([]*battle.Fighter, module.ALL_FIGHTER_POS_NUM)
	attackFighters := make([]*battle.BattlePlayerInfo, 0, num)
	rpc.RemoteNewPlayerFighter(args, func(replys []*rpc.Reply_NewPlayerFighter, errs []error) {
		for _, err := range errs {
			fail.When(err != nil, errs)
		}

		for _, reply := range replys {
			attackFighters = append(attackFighters, reply.BattlePlayerInfo)
			if reply.BattlePlayerInfo.PlayerId == room.LeaderPid {
				err := json.Unmarshal(reply.TotemInfo, &battleTotemInfo)
				fail.When(err != nil, err)
			}
			for _, raw := range reply.Fighters {
				fighter := new(battle.Fighter)
				err := json.Unmarshal(raw, fighter)
				fail.When(err != nil, err)
				fighters[fighter.Position-1] = fighter
			}
		}

		attackSide := &battle.SideInfo{
			Groups:    [][]*battle.Fighter{fighters},
			Fighters:  fighters,
			Players:   attackFighters,
			TotemInfo: battleTotemInfo,
		}

		defendSide := module.NewBattleSideWithEnemyDeployForm(battle.BT_MULTI_LEVEL, room.LevelId)

		room.OnFighting = true
		//module.Battle.NewBattleRoom(battle.BT_MULTI_LEVEL, attackSide, defendSide, room.channel, room, 15*goTime.Second)
		module.Battle.NewBattleRoom(battle.BT_MULTI_LEVEL, attackSide, defendSide, room.channel, room, 30*goTime.Second)
	})
}
