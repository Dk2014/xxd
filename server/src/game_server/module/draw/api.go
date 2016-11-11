package draw

import (
	"core/net"
	"game_server/api/protocol/draw_api"
	"game_server/dat/heart_dat"
	"game_server/module"
	"game_server/xdlog"
)

func init() {
	draw_api.SetInHandler(DrawAPI{})
}

type DrawAPI struct {
}

func (api DrawAPI) GetHeartDrawInfo(session *net.Session, in *draw_api.GetHeartDrawInfo_In) {
	getHeartDrawInfo(session, in.DrawType, in.AwardRecord)
}

func (api DrawAPI) HeartDraw(session *net.Session, in *draw_api.HeartDraw_In) {
	heartDraw(session, in.DrawType, xdlog.ET_HEART_DRAW)
}

func (api DrawAPI) GetChestInfo(session *net.Session, in *draw_api.GetChestInfo_In) {
	//out := &draw_api.GetChestInfo_Out{}
	//GetChestInfo(session, out)
	//session.Send(out)
}

func (api DrawAPI) DrawChest(session *net.Session, in *draw_api.DrawChest_In) {
	//out := &draw_api.DrawChest_Out{}
	//out.Items = []draw_api.DrawChest_Out_Items{}
	////DrawChest(session, in, out)
	//session.Send(out)
}

func (api DrawAPI) HeartInfo(session *net.Session, in *draw_api.HeartInfo_In) {
	state := module.State(session)
	playerHeart := state.Database.Lookup.PlayerHeart(state.PlayerId)
	session.Send(&draw_api.HeartInfo_Out{
		RecoverToday: playerHeart.RecoverDayCount,
		Timestamp:    playerHeart.UpdateTime + heart_dat.HEART_RECOVER_TIME_INTERVAL,
	})
}

func (api DrawAPI) GetFateBoxInfo(session *net.Session, in *draw_api.GetFateBoxInfo_In) {
	state := module.State(session)
	out := &draw_api.GetFateBoxInfo_Out{}
	getFateBoxInfo(state, out)
	session.Send(out)
}

func (api DrawAPI) OpenFateBox(session *net.Session, in *draw_api.OpenFateBox_In) {
	out := &draw_api.OpenFateBox_Out{}
	openFateBox(session, in.BoxType, uint8(in.Times), out)
	session.Send(out)
}

func (this DrawAPI) ExchangeGiftCode(session *net.Session, in *draw_api.ExchangeGiftCode_In) {
	exchangeGiftCode(session, string(in.Code))
}
