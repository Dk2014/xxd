package module

import (
	"core/log"
	"core/net"
	util "core/time"
	"game_server/api/protocol/channel_api"
	"sync"
	"time"
)

const (
	CHANNEL_CHAT           = iota //世界聊天
	CHANNEL_MESSAGE               //世界公告
	CHANNEL_CLIQUE_MESSAGE        //帮派世界公告
	//增加消息类型需要同步更新协议注释 channel.pd
)

const (
	BUFF_SIZE           = 100
	BUFF_MAX_CHAN       = 2000
	BUFF_FLUSH_INTERVAL = 5 // second
)

var (
	worldChatChannel *WorldChannelBuff
)

func BroadcastChannelStart() {
	worldChatChannel = &WorldChannelBuff{
		broadcastChan: make(chan net.Response, BUFF_MAX_CHAN),
	}
	worldChatChannel.loop()
}

type Message struct {
	Id      int64
	MsgType int8

	Pid        int64
	Nickname   []byte
	Timestamp  int64
	Content    []byte
	TplId      int16
	Parameters []byte
}

type WorldChannelBuff struct {
	sync.RWMutex
	buff          [BUFF_SIZE]*Message
	base          uint64
	top           uint64
	broadcastChan chan net.Response
}

func (this *WorldChannelBuff) loop() {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Errorf("Error = %v\n", err)
			}
		}()
	L:
		for {
			select {
			case response, ok := <-this.broadcastChan:
				if !ok {
					break L
				}
				API.Broadcast(Player, response)

			case _, ok := <-time.After(BUFF_FLUSH_INTERVAL * time.Second):
				if !ok {
					break L
				}
				this.doFlush()
			}
		}
	}()
}

func (this *WorldChannelBuff) doFlush() {
	this.Lock()
	defer this.Unlock()
	this.flush()
}

func (this *WorldChannelBuff) AddMsg(msg *Message) {
	this.Lock()
	defer this.Unlock()
	if this.top-this.base < BUFF_SIZE-1 {
		this.buff[this.top%BUFF_SIZE] = msg
		this.top++
	} else {
		this.flush()
		this.buff[this.top%BUFF_SIZE] = msg
		this.top++
	}
}

func (this *WorldChannelBuff) flush() {
	rsp := &channel_api.SendGlobalMessages_Out{}

	for this.base < this.top {
		msg := this.buff[this.base%BUFF_SIZE]
		rsp.Messages = append(rsp.Messages, channel_api.SendGlobalMessages_Out_Messages{
			Pid:        msg.Pid,
			Nickname:   []byte(msg.Nickname),
			MsgType:    channel_api.MessageType(msg.MsgType),
			Timestamp:  msg.Timestamp,
			Parameters: []byte(msg.Parameters),
			TplId:      msg.TplId,
		})
		this.base++
	}

	if len(rsp.Messages) <= 0 {
		return
	}

	this.broadcastChan <- rsp
}

func (this *WorldChannelBuff) GetLatest(num, secondDelt int64) (msgs []*Message) {
	this.RLock()
	defer this.RUnlock()

	start := this.base - uint64(num)
	if start < 0 {
		start = 0
	}
	now := util.GetNowTime()
	for x := start; x < this.base; x++ {
		if this.buff[x%BUFF_SIZE].Timestamp+secondDelt >= now {
			msgs = append(msgs, this.buff[x%BUFF_SIZE])
		}
	}
	return msgs
}
