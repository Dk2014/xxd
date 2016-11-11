package player

import (
	"core/fail"
	"core/net"
	"core/time"
	"game_server/config"
	"game_server/module"
)

/*

	跨服登陆。
		- 多人关卡

*/

func CrossLogin(session *net.Session, playerId, loginTime int64, openid, nick, hash []byte, roleId int8, roleLevel int16, wegamesUid []byte) bool {
	fail.When(config.ServerCfg.EnableGlobalServer, "forbidden cross login")
	fail.When(time.GetNowTime()-loginTime > 20000, "CrossLogin timeout")
	fail.When(!crossServerLoginTest(playerId, loginTime, openid, nick, hash, roleId, wegamesUid), "loginGlobal hash incorrect")

	state := module.State(session)

	state.RoleId = roleId
	state.PlayerId = playerId
	state.PlayerNick = nick
	state.WegamesPlatformUid = wegamesUid
	state.CrossInfo = &module.CrossPlayerInfo{
		RoleLevel: roleLevel,
	}

	return true
}

func (mod PlayerMod) LogoutFromCross(playerId int64, session *net.Session) {
	module.MultiLevel.LeaveMultiLevel(session)
}
