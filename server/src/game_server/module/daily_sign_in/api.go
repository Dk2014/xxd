package daily_sign_in

import (
	"core/net"
	"game_server/api/protocol/daily_sign_in_api"
	"game_server/xdlog"
)

func init() {
	daily_sign_in_api.SetInHandler(DailySignInAPI{})
}

type DailySignInAPI struct {
}

func (this DailySignInAPI) Info(session *net.Session, in *daily_sign_in_api.Info_In) {
	out := new(daily_sign_in_api.Info_Out)
	info(session, out)
	session.Send(out)
}

func (this DailySignInAPI) Sign(session *net.Session, in *daily_sign_in_api.Sign_In) {
	out := new(daily_sign_in_api.Sign_Out)
	sign(session, out, xdlog.ET_DAILY_SIGN)
	session.Send(out)
}

func (this DailySignInAPI) SignPastDay(session *net.Session, in *daily_sign_in_api.SignPastDay_In) {
	out := new(daily_sign_in_api.SignPastDay_Out)
	signPastDay(session, in.Index, out, xdlog.ET_DAILY_SIGN)
	session.Send(out)
}
