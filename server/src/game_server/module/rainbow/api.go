package rainbow

import (
	//"core/fail"
	"core/net"
	"game_server/api/protocol/rainbow_api"
	"game_server/module"
)

func init() {
	rainbow_api.SetInHandler(RainbowAPI{})
}

type RainbowAPI struct{}

func (api RainbowAPI) Info(session *net.Session, _ *rainbow_api.Info_In) {
	out := &rainbow_api.Info_Out{}
	state := module.State(session)
	info(state, out)
	session.Send(out)
}

func (api RainbowAPI) Reset(session *net.Session, _ *rainbow_api.Reset_In) {
	out := &rainbow_api.Reset_Out{}
	state := module.State(session)
	resetRainbowLevel(state)
	session.Send(out)
}

func (api RainbowAPI) TakeAward(session *net.Session, in *rainbow_api.TakeAward_In) {
	out := &rainbow_api.TakeAward_Out{}
	state := module.State(session)
	pos1 := state.RainbowLevelState.AwardBoxIndex[in.Pos1-1]
	pos2 := state.RainbowLevelState.AwardBoxIndex[in.Pos2-1]
	out.NextLevel = takeAward(state, pos1, pos2)
	session.Send(out)
}

func (api RainbowAPI) AwardInfo(session *net.Session, in *rainbow_api.AwardInfo_In) {
	state := module.State(session)
	out := &rainbow_api.AwardInfo_Out{}
	awardInfo(state, out)
	session.Send(out)
}

func (api RainbowAPI) JumpToSegment(session *net.Session, in *rainbow_api.JumpToSegment_In) {
	state := module.State(session)
	out := &rainbow_api.JumpToSegment_Out{}
	jumpToSegment(state, in.Segment)
	session.Send(out)
}

func (api RainbowAPI) AutoFight(session *net.Session, in *rainbow_api.AutoFight_In) {
	state := module.State(session)
	out := rainbowLevelAutoFight(state, in.Segment)
	session.Send(out)
}
