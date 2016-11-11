package meditation

import (
	"core/net"
	"core/time"
	"game_server/api/protocol/meditation_api"
	"game_server/api/protocol/town_api"
	"game_server/dat/meditation_dat"
	"game_server/dat/player_dat"
	"game_server/module"
	"game_server/module/rpc"
)

type MeditationAPI struct{}

func init() {
	meditation_api.SetInHandler(MeditationAPI{})
}

func (this MeditationAPI) MeditationInfo(session *net.Session, in *meditation_api.MeditationInfo_In) {
	state := module.State(session)
	expTime, keyTime := meditationInfo(state)
	out := &meditation_api.MeditationInfo_Out{
		KeyAccumulateTime: keyTime,
		ExpAccumulateTime: expTime,
	}
	session.Send(out)
}

func (this MeditationAPI) StartMeditation(session *net.Session, in *meditation_api.StartMeditation_In) {
	state := module.State(session)
	module.Player.MustOpenFunc(state.Database, player_dat.FUNC_MEDITATION)

	startMeditation(state.Database, time.GetNowTime(), in.InClubhouse)

	expTime, keyTime := meditationInfo(state)
	out := &meditation_api.StartMeditation_Out{
		KeyAccumulateTime: keyTime,
		ExpAccumulateTime: expTime,
	}
	if !in.InClubhouse {
		if state.TownChannel != nil {
			module.API.Broadcast(state.TownChannel, &town_api.UpdateTownPlayerMeditationState_Out{
				PlayerId:        state.PlayerId,
				MeditationState: true,
			})
		}
	}

	session.Send(out)
}

func (this MeditationAPI) StopMeditation(session *net.Session, in *meditation_api.StopMeditation_In) {
	state := module.State(session)
	stopMeditation(state.Database, in.InClubhouse)
	rpc.RemoteGetClubHouseMeditation(state.PlayerId, meditation_dat.CLUBHOUSE_MEDITATION_STOP)
	if !in.InClubhouse && state.TownChannel != nil {
		module.API.Broadcast(state.TownChannel, &town_api.UpdateTownPlayerMeditationState_Out{
			PlayerId:        state.PlayerId,
			MeditationState: false,
		})
	}
}
