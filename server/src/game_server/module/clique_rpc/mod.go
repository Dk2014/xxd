package clique_rpc

import (
	"core/net"
	"core/time"
	"game_server/api/protocol/channel_api"
	"game_server/dat/channel_dat"
	"game_server/mdb"
	"game_server/module"
)

func init() {
	module.CliqueRPC = CliqueMod{}
}

type CliqueMod struct{}

//部分实现在 xxx_mod.go 文件中

func (mod CliqueMod) GetLatestMessages(cliqueId int64) []channel_api.CliqueMessage {
	return CacheGetLatestMessage(cliqueId)
}

func (mod CliqueMod) AddCliqueNews(cliqueId int64, msgTpl channel_dat.MessageTpl) {
	msg := channel_api.CliqueMessage{
		MsgType: channel_api.MESSAGE_TYPE_CLIQUE_NEWS,
		//Pid:        pid,
		//Nickanme:   nick,
		TplId:      msgTpl.GetTplId(),
		Parameters: msgTpl.GetParameters(),
		Timestamp:  time.GetNowTime(),
	}
	module.CliqueRPC.Broadcast(cliqueId, &channel_api.SendCliqueMessage_Out{
		Message: msg,
	})
	CacheAddMessage(cliqueId, msg)
}

func (mod CliqueMod) AddCliqueChat(cliqueId int64, pid int64, nick []byte, msgTpl channel_dat.MessageTpl) {
	msg := channel_api.CliqueMessage{
		MsgType:    channel_api.MESSAGE_TYPE_CLIQUE_CHAT,
		Pid:        pid,
		Nickname:   nick,
		TplId:      msgTpl.GetTplId(),
		Parameters: msgTpl.GetParameters(),
		Timestamp:  time.GetNowTime(),
	}
	module.CliqueRPC.Broadcast(cliqueId, &channel_api.SendCliqueMessage_Out{
		Message: msg,
	})
	CacheAddMessage(cliqueId, msg)
}

func (mod CliqueMod) LoginUpdateCliqueInfo(session *net.Session) {
	state := module.State(session)
	//等了更新帮派信息

	//1. 帮主需要更新登陆时间
	playerCliqueInfo := state.Database.Lookup.PlayerGlobalCliqueInfo(state.PlayerId)
	if playerCliqueInfo == nil || playerCliqueInfo.CliqueId <= 0 {
		return
	}
	cliqueInfo := cliqueInfoLookUp(state.Database, playerCliqueInfo.CliqueId)
	if cliqueInfo.OwnerPid == state.PlayerId {
		cliqueInfo.OwnerLoginTime = time.GetNowTime()
	}
	//加入帮派频道
	JoinCliqueChannel(cliqueInfo.Id, session)
	state.Database.Update.GlobalClique(cliqueInfo)
}

func (mod CliqueMod) Logout(session *net.Session) {
	state := module.State(session)
	playerCliqueInfo := state.Database.Lookup.PlayerGlobalCliqueInfo(state.PlayerId)
	if playerCliqueInfo == nil || playerCliqueInfo.CliqueId <= 0 {
		return
	}
	leaveClubhouse(session)
	module.CliqueRPC.LeaveCliqueChannel(playerCliqueInfo.CliqueId, session)
}

func (this CliqueMod) GetMemberNum(cliqueId int64) int16 {
	return CacheGetMembeNum(cliqueId)
}

//增加贡献
func (this CliqueMod) AddPlayerCliqueContrib(db *mdb.Database, contrib int64) {
	addCliqueContrib(db, contrib)
	//TODO notify
}

//扣除贡献
func (this CliqueMod) DecPlayerCliqueContrib(db *mdb.Database, contrib int64) bool {
	return delCliqueContrib(db, contrib)
	//TODO notify
}

//lookup global_clique表，并在周一清空
func (this CliqueMod) CliqueInfoLookUp(db *mdb.Database, cliqueId int64) *mdb.GlobalClique {
	return cliqueInfoLookUp(db, cliqueId)
}

func (this CliqueMod) CliqueInfoListPid(cliqueId int64) []int64 {
	return GetCliquePid(cliqueId)
}

func (this CliqueMod) ClearCliqueDailyDonate(session *net.Session) {
	state := module.State(session)
	playerCliqueInfo := state.Database.Lookup.PlayerGlobalCliqueInfo(state.PlayerId)
	if playerCliqueInfo == nil || playerCliqueInfo.CliqueId <= 0 {
		return
	}
	playerCliqueInfo.DonateCoinsTime = 1
	state.Database.Update.PlayerGlobalCliqueInfo(playerCliqueInfo)

}

func (this CliqueMod) IsCliqueMember(cliqueId int64, pid int64) bool {
	return CacheIsCliqueMember(cliqueId, pid)
}

func (this CliqueMod) GetCliqueNameById(cliqueId int64) string {
	cliqueCacheInfo := CacheGetCliqueInfo(cliqueId)
	if cliqueCacheInfo != nil {
		return cliqueCacheInfo.Name
	}
	return ""
}
