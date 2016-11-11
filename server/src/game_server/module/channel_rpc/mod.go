package channel_rpc

import (
	"core/time"
	"game_server/dat/channel_dat"
	"game_server/module"
)

type ChanMod struct{}

func init() {
	module.Chan = ChanMod{}
}

//仅仅在互动服调用
func (mod ChanMod) AddWorldTplMessage(pid int64, nick []byte, msgType int8, msgTpl channel_dat.MessageTpl) {
	module.AddWorldChannelMessage(&module.Message{
		Pid:        pid,
		TplId:      msgTpl.GetTplId(),
		Timestamp:  time.GetNowTime(),
		Parameters: msgTpl.GetParameters(),
		MsgType:    msgType,
	})
}
