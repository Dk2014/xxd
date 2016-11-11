package town

import (
	"core/fail"
	"core/net"
	"game_server/api/protocol/town_api"
	"game_server/dat/player_dat"
	"game_server/dat/quest_dat"
	"game_server/dat/town_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/tlog"
	"game_server/xdlog"
)

func EnterTown(session *net.Session, townId int16) {
	state := module.State(session)
	town := town_dat.GetTownWithTownId(townId)
	playerTown := state.Database.Lookup.PlayerTown(state.PlayerId)
	playerFashionState := state.Database.Lookup.PlayerFashionState(state.PlayerId)

	fail.When(playerTown.Lock < town.Lock, "town is not open")

	// 每次进入城镇，清理关卡数据
	module.Mission.LeaveMissionLevel(state)

	state.TownId = townId
	module.Quest.RefreshQuest(state, quest_dat.QUEST_TYPE_TOWN, xdlog.ET_QUEST)

	if state.TownChannel != nil {
		g_TownChannelMgr.removePlayer(session)
	}
	g_TownChannelMgr.addPlayer(session, townId)

	state.LastTownX, state.LastTownY = playerTown.AtPosX, playerTown.AtPosY
	if townId != playerTown.TownId {
		state.LastTownX, state.LastTownY = town.StartX, town.StartY
	}

	module.API.Broadcast(state.TownChannel, &town_api.Enter_Out{
		Player: town_api.TownPlayer{
			PlayerId:          state.PlayerId,
			Nickname:          state.PlayerNick,
			RoleId:            state.RoleId,
			AtX:               state.LastTownX,
			AtY:               state.LastTownY,
			FashionId:         playerFashionState.DressedFashionId,
			InMeditationState: module.Meditation.InMeditationState(state.Database),
			Level:             module.Role.GetMainRole(state.Database).Level,
		},
	})

	townPlayers := &town_api.NotifyTownPlayers_Out{}
	state.TownChannel.Fetch(func(playerSession *net.Session) {
		otherState := module.State(playerSession)
		playerFashionState := otherState.Database.Lookup.PlayerFashionState(otherState.PlayerId)
		townPlayers.Players = append(townPlayers.Players, town_api.NotifyTownPlayers_Out_Players{
			Player: town_api.TownPlayer{
				PlayerId:          otherState.PlayerId,
				Nickname:          otherState.PlayerNick,
				RoleId:            otherState.RoleId,
				AtX:               otherState.LastTownX,
				AtY:               otherState.LastTownY,
				FashionId:         playerFashionState.DressedFashionId,
				InMeditationState: module.Meditation.InMeditationState(otherState.Database),
				Level:             module.Role.GetMainRole(otherState.Database).Level,
			},
		})
	})

	session.Send(townPlayers)
}

func findoutOpenedTownTreasures(db *mdb.Database, out *town_api.ListOpenedTownTreasures_Out) {
	db.Select.PlayerOpenedTownTreasure(func(row *mdb.PlayerOpenedTownTreasureRow) {
		out.Treasures = append(out.Treasures, town_api.ListOpenedTownTreasures_Out_Treasures{
			TownId: row.TownId(),
		})
	})
}

func awardTownTreasure(state *module.SessionState, town_id int16) {
	xdEventType := int32(xdlog.ET_TOWN_TREASURES)
	//town := town_dat.GetTownWithTownId(town_id)
	//playerTown := state.Database.Lookup.PlayerTown(state.PlayerId)
	//fail.When(playerTown.Lock < town.Lock, "town is not open")

	state.Database.Select.PlayerOpenedTownTreasure(func(row *mdb.PlayerOpenedTownTreasureRow) {
		fail.When(row.TownId() == town_id, "town treasure has been opened once")
	})
	treasure, ok := town_dat.GetTownTreasure(town_id)
	fail.When(!ok, "this town treasure doesn't exist")

	// 所有上阵角色增加经验
	if treasure.AwardExp > int32(0) {
		module.Role.AddFormRoleExp(state, int64(treasure.AwardExp), tlog.EFT_TOWN_TREASURE)
	}

	if treasure.AwardCoins > 0 {
		module.Player.IncMoney(state.Database, state.MoneyState, treasure.AwardCoins, player_dat.COINS, tlog.MFR_TOWN_TREASURE, xdEventType, "")
	}

	if treasure.AwardPhysical > 0 {
		module.Physical.AwardIncrease(state, int16(treasure.AwardPhysical), tlog.PFR_TOWN_TREASURE)
	}

	if treasure.AwardIngot > 0 {
		module.Player.IncMoney(state.Database, state.MoneyState, int64(treasure.AwardIngot), player_dat.INGOT, tlog.MFR_TOWN_TREASURE, xdEventType, "")
	}

	if treasure.AwardItem1Id > int16(0) {
		module.Item.AddItem(state.Database, treasure.AwardItem1Id, treasure.AwardItem1Num, tlog.IFR_TOWN_TREASURE, xdEventType, "")
	}

	if treasure.AwardItem2Id > int16(0) {
		module.Item.AddItem(state.Database, treasure.AwardItem2Id, treasure.AwardItem2Num, tlog.IFR_TOWN_TREASURE, xdEventType, "")
	}

	if treasure.AwardItem3Id > int16(0) {
		module.Item.AddItem(state.Database, treasure.AwardItem3Id, treasure.AwardItem3Num, tlog.IFR_TOWN_TREASURE, xdEventType, "")
	}

	if treasure.AwardItem4Id > int16(0) {
		module.Item.AddItem(state.Database, treasure.AwardItem4Id, treasure.AwardItem4Num, tlog.IFR_TOWN_TREASURE, xdEventType, "")
	}

	state.Database.Insert.PlayerOpenedTownTreasure(&mdb.PlayerOpenedTownTreasure{
		Pid:    state.PlayerId,
		TownId: town_id,
	})
}
