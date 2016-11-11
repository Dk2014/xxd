package town

import (
	"core/net"
	"game_server/api/protocol/town_api"
	"game_server/module"
	"time"
)

func init() {
	town_api.SetInHandler(TownAPI{})
}

type TownAPI struct {
}

func (api TownAPI) Enter(session *net.Session, in *town_api.Enter_In) {
	EnterTown(session, in.TownId)
}

func (api TownAPI) Leave(session *net.Session, in *town_api.Leave_In) {
	module.Town.LeaveTown(session)
}

func (api TownAPI) Move(session *net.Session, in *town_api.Move_In) {
	state := module.State(session)

	if state.TownChannel == nil {
		return
	}

	now := time.Now().UnixNano() / 1e6
	state.LastTownX, state.LastTownY = in.ToX, in.ToY

	if now-state.LastMoveInTown > 4000 {
		module.API.Broadcast(state.TownChannel, &town_api.Move_Out{
			PlayerId: state.PlayerId,
			ToX:      in.ToX,
			ToY:      in.ToY,
		})

		state.LastMoveInTown = now
	}
}

func (api TownAPI) TalkedNpcList(session *net.Session, in *town_api.TalkedNpcList_In) {
	state := module.State(session)
	out := &town_api.TalkedNpcList_Out{}
	talkedNpcList(state, in.TownId, out)
	session.Send(out)
}

func (api TownAPI) NpcTalk(session *net.Session, in *town_api.NpcTalk_In) {
	state := module.State(session)
	npcTalkAward(state, in.NpcId)
}

func (api TownAPI) ListOpenedTownTreasures(session *net.Session, in *town_api.ListOpenedTownTreasures_In) {
	state := module.State(session)
	db := state.Database
	out := &town_api.ListOpenedTownTreasures_Out{}
	findoutOpenedTownTreasures(db, out)
	session.Send(out)
}

func (api TownAPI) TakeTownTreasures(session *net.Session, in *town_api.TakeTownTreasures_In) {
	state := module.State(session)
	awardTownTreasure(state, in.TownId)
	out := &town_api.TakeTownTreasures_Out{}
	session.Send(out)
}
