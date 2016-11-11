package rpc

import (
	"game_server/config"
	"game_server/dat/channel_dat"
	"game_server/global"
	"game_server/mdb"
	"game_server/module"
	. "game_server/rpc"
)

type Message struct {
	TplId      int16
	Parameters []byte
}

type Args_AddWorldChannelTplMessage struct {
	RPCArgTag
	Pid      int64
	Messages []*Message
}

func (this *Message) GetTplId() int16 {
	return this.TplId
}

func (this *Message) GetParameters() []byte {
	return this.Parameters
}

type Reply_AddWorldChannelTplMessage struct {
}

//游戏服务器作为RPC服务器，本服玩家通过该服务增加邮件
func (mail *RemoteServe) AddWorldChannelTplMessage(args *Args_AddWorldChannelTplMessage, reply *Reply_AddWorldChannelTplMessage) error {
	return Remote.Serve(mdb.RPC_Remote_AddWorldChannelTplMessage, args, mdb.TRANS_TAG_RPC_Serve_AddWorldChannelTplMessage, func() error {
		playerInfo := global.GetPlayerInfo(args.Pid)
		for _, msg := range args.Messages {
			module.Chan.AddWorldTplMessage(args.Pid, playerInfo.PlayerNick, module.CHANNEL_MESSAGE, msg)
		}
		return nil
	})
}

func RemoteAddWorldChannelTplMessage(pid int64, msgTpls []channel_dat.MessageTpl) {
	args := &Args_AddWorldChannelTplMessage{
		Pid: pid,
	}
	for _, msgTpl := range msgTpls {
		args.Messages = append(args.Messages, &Message{
			TplId:      msgTpl.GetTplId(),
			Parameters: msgTpl.GetParameters(),
		})
	}
	Remote.Call(config.ServerCfg.GlobalServerId, mdb.RPC_Remote_AddWorldChannelTplMessage, args, &Reply_AddWorldChannelTplMessage{}, mdb.TRANS_TAG_RPC_Call_AddWorldChannelTplMessage, nil)
}
