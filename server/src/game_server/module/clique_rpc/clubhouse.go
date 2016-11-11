package clique_rpc

import (
	"core/fail"
	"core/net"
	"game_server/global"
	//"game_server/mdb"
	"game_server/api/protocol/clique_api"
	//"game_server/api/protocol/town_api"
	"game_server/dat/town_dat"
	"game_server/module"
)

func enterClubhouse(session *net.Session) {
	state := module.State(session)
	playerCliqueInfo := state.Database.Lookup.PlayerGlobalCliqueInfo(state.PlayerId)
	if playerCliqueInfo == nil || playerCliqueInfo.CliqueId <= 0 {
		session.Send(&clique_api.EnterClubhouse_Out{
			Ok:     false,
			Player: clique_api.Player{},
		})
		return
	}
	state.InCliqueClubhouse = true
	//集会所初始坐标
	townInfo := town_dat.GetJiHuiSuo()
	state.ClubhouseX = townInfo.StartX
	state.ClubhouseY = townInfo.StartY

	//1. 通知帮派内成员信息
	otherPlayers := &clique_api.NotifyClubhousePlayers_Out{}
	channel := GetCliqueChannel(playerCliqueInfo.CliqueId)
	channel.Fetch(func(otherSession *net.Session) {
		otherState := module.State(otherSession)
		if !otherState.InCliqueClubhouse {
			return
		}
		playerInfo := global.GetPlayerInfo(otherState.PlayerId)
		if otherState.InCliqueClubhouse == true {
			otherPlayers.Players = append(otherPlayers.Players, clique_api.NotifyClubhousePlayers_Out_Players{
				Player: clique_api.Player{
					PlayerId:          otherState.PlayerId,
					RoleId:            otherState.RoleId,
					Nickname:          otherState.PlayerNick,
					AtX:               otherState.ClubhouseX,
					AtY:               otherState.ClubhouseY,
					Level:             playerInfo.RoleLevel,
					InMeditationState: otherState.IsMeditation,
					FashionId:         playerInfo.FashionId,
				},
			})
		}
	})
	session.Send(otherPlayers)

	//2. 通知其他成员该玩家的信息
	playerInfo := global.GetPlayerInfo(state.PlayerId)
	module.CliqueRPC.BroadcastClubhouse(playerCliqueInfo.CliqueId, &clique_api.NotifyNewPlayer_Out{
		Player: clique_api.Player{
			PlayerId:          state.PlayerId,
			Level:             playerInfo.RoleLevel,
			RoleId:            state.RoleId,
			Nickname:          state.PlayerNick,
			AtX:               state.ClubhouseX,
			AtY:               state.ClubhouseY,
			InMeditationState: state.IsMeditation,
			FashionId:         playerInfo.FashionId,
		},
	})
	session.Send(&clique_api.EnterClubhouse_Out{
		Ok: true,
		Player: clique_api.Player{
			PlayerId:          state.PlayerId,
			Level:             playerInfo.RoleLevel,
			RoleId:            state.RoleId,
			Nickname:          state.PlayerNick,
			AtX:               state.ClubhouseX,
			AtY:               state.ClubhouseY,
			InMeditationState: state.IsMeditation,
			FashionId:         playerInfo.FashionId,
		},
	})
}

func leaveClubhouse(session *net.Session) {
	state := module.State(session)
	playerCliqueInfo := state.Database.Lookup.PlayerGlobalCliqueInfo(state.PlayerId)
	if playerCliqueInfo == nil || playerCliqueInfo.CliqueId <= 0 {
		return
	}
	if !state.InCliqueClubhouse {
		return
	}
	state.InCliqueClubhouse = false
	state.IsMeditation = false
	//1. 通知其他成员该玩家离开
	module.CliqueRPC.BroadcastClubhouse(playerCliqueInfo.CliqueId, &clique_api.LeaveClubhouse_Out{
		PlayerId: state.PlayerId,
	})
	//module.CliqueRPC.BroadcastClubhouse(playerCliqueInfo.CliqueId, &town_api.UpdateTownPlayerMeditationState_Out{
	//	PlayerId:        state.PlayerId,
	//	MeditationState: false,
	//})
}

func clubhouseMove(session *net.Session, x, y int16) {
	state := module.State(session)
	playerCliqueInfo := state.Database.Lookup.PlayerGlobalCliqueInfo(state.PlayerId)
	if playerCliqueInfo == nil || playerCliqueInfo.CliqueId <= 0 {
		return
	}
	fail.When(!state.InCliqueClubhouse, "不在帮派集会场所中")
	state.ClubhouseX = x
	state.ClubhouseY = y
	module.CliqueRPC.BroadcastClubhouse(playerCliqueInfo.CliqueId, &clique_api.ClubMove_Out{
		PlayerId: state.PlayerId,
		ToX:      x,
		ToY:      y,
	})

}
