package player

import (
	"core/fail"
	"core/log"
	"core/net"
	"core/time"
	"game_server/api/protocol/player_api"
	"game_server/config"
	"game_server/dat/event_dat"
	"game_server/global"
	"game_server/mdb"
	"game_server/module"
	"game_server/tencent"
)

/*

 重点：在互动服的逻辑处理都是全局事务的

 处理互动服务器登陆
 	- 仅此接口混合在和游戏player模块中。
 	- 其他互动功能模块的逻辑都组织在如moduleName_rpc，后缀为_rpc的模块中

 互动服登陆流程：
 	1、首先会登陆游戏服。
 	2、成功登陆游戏服后，如果是新玩家，游戏服同时会向互动服通过rpc调用RemoteAddPlayerInfo，互动服创建新玩家数据，并为新玩家初始化互动服上的功能
 	3、游戏服收到新玩家在互动服初始化成功的通知后，再通知客户端，可以连接到互动服

 	这里规定: 新/老玩家连接互动服的时机在player::info请求成功后

 	所以互动服的登陆接口不用考虑在游戏服创建的新玩家数据是否同步到互动服
*/

func loginGlobal(session *net.Session, playerId, loginTime int64, openid, nick, hash []byte, roleId int8, roleLevel int16, fightNum int32, in *player_api.GlobalLogin_In) (ret bool, payToken bool) {
	fail.When(!config.ServerCfg.EnableGlobalServer, "forbidden global login")
	fail.When(time.GetNowTime()-loginTime > 20000, "loginGlobal timeout")
	fail.When(int(in.Gsid) != config.ServerCfg.ServerId, "登陆hd-srv请求的gsid错误")
	fail.When(!crossServerLoginTest(playerId, loginTime, openid, nick, hash, roleId, in.WegamesPlatformUid), "loginGlobal hash incorrect")

	state := module.State(session)
	state.RoleId = roleId
	state.PlayerId = playerId
	state.PlayerNick = nick
	state.CId = int32(in.ChannelId)
	state.WegamesPlatformUid = in.WegamesPlatformUid

	// 支付系统需要
	state.MoneyState = &module.MoneyState{
		Openid:       string(in.Username),
		Openkey:      string(in.Openkey),
		PayToken:     string(in.PayToken),
		Pfkey:        string(in.Pfkey),
		Zoneid:       int(in.Zoneid),
		Pf:           string(in.Pf),
		PlatformType: int8(in.PlatformType),
	}

	// 检查支付token是否有效
	TencentPayTest(state.MoneyState, func(rsp *tencent.GET_BALANCE_RSP) {
		if rsp.Ret != 0 {
			log.Errorf("(hd)TencentPayTest faild. %v, %v, %v ,%v ,%v, %v, %v\n",
				state.MoneyState.Openid,
				state.MoneyState.Openkey,
				state.MoneyState.PayToken,
				state.MoneyState.Pfkey,
				state.MoneyState.Zoneid,
				state.MoneyState.PlatformType,
				rsp,
			)
			session.Send(&player_api.GlobalLogin_Out{Paytoken: false, Result: false})
		}
	})

	// 如果玩家第一次登陆，初始化数据
	if !global.ExistPlayerInfo(playerId) {
		global.NewPlayerMount(state.Database, &global.PlayerInfo{
			Openid:             openid,
			PlayerId:           playerId,
			PlayerNick:         nick,
			RoleId:             roleId,
			RoleLevel:          roleLevel,
			CId:                int64(in.ChannelId),
			WegamesPlatformUid: in.WegamesPlatformUid,
		})
	} else {
		mdb.GlobalMount(state.Database, playerId)
		if !global.CheckPlayerHaveInitData(state.Database) {
			global.InitPlayerData(state.Database)
		}
	}

	SetPlayerOnline(state.CId, playerId, session)

	global.UpdatePlayerLoginTime(playerId, loginTime)

	module.Friend.SendOfflineMessages(session)
	module.MailRPC.SendGlobaldMail(state.Database, state.PlayerId)
	module.CliqueRPC.LoginUpdateCliqueInfo(session)
	module.CliqueEscortRPC.LoginEscortNotify(session)

	isOpenOrEndIngotRankTime(session)
	return true, true
}

func (mod PlayerMod) LogoutFromGlobal(cid int32, playerId int64, session *net.Session) {
	SetPlayerOffline(cid, playerId)
	module.CliqueRPC.Logout(session)
}

//判断当前时间是否是开放或者关闭充值排行榜的时间
func isOpenOrEndIngotRankTime(session *net.Session) {
	eventInGotRankConfig := event_dat.GetEventIngotRankConfig()
	openIngotRankTime := eventInGotRankConfig.StartUnixTime
	endIngotRankTime := eventInGotRankConfig.EndUnixTime
	nowTime := time.GetNowTime()
	isOpen := false
	if nowTime >= openIngotRankTime && nowTime <= endIngotRankTime {
		isOpen = true
	}
	session.Send(&player_api.OpenIngotRank_Out{
		IsOpenIngotRank: isOpen,
	})

	closeIngotRankPanelTime := eventInGotRankConfig.EndUnixTime + 24*60*60
	isClose := false
	if nowTime >= closeIngotRankPanelTime {
		isClose = true
	}

	session.Send(&player_api.CloseIngotRankPanel_Out{
		IsCloseIngotRankPanel: isClose,
	})
}
