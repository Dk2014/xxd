package town

import (
	"core/fail"
	"core/net"
	"game_server/module"
	"sync"
)

var g_TownChannelMgr *townChannelMgr = &townChannelMgr{
	townChans: make(map[int16]*townChannel),
}

const (
	TownChannelMaxCount = 10
)

type townChannel struct {
	branches []*net.Channel
}

func (this *townChannel) getBranch() *net.Channel {
	// 找出还未满的通道
	for _, branch := range this.branches {
		if branch.Count() < TownChannelMaxCount {
			return branch
		}
	}

	// 所有通道都满了，就新建一条
	branch := net.NewChannel()
	this.branches = append(this.branches, branch)
	return branch
}

type townChannelMgr struct {
	sync.Mutex
	townChans map[int16]*townChannel
}

func (this *townChannelMgr) addPlayer(session *net.Session, townId int16) {
	this.Lock()
	defer this.Unlock()

	townChan, ok := this.townChans[townId]
	if !ok {
		// 为城镇新建一个channel
		townChan = &townChannel{}
		this.townChans[townId] = townChan
	}

	state := module.State(session)
	state.TownChannel = townChan.getBranch()
	state.TownChannel.Join(session)
}

func (this *townChannelMgr) removePlayer(session *net.Session) {
	state := module.State(session)
	fail.When(state.TownChannel == nil, "state.TownChannel is nil")

	state.TownChannel.Exit(session)
	state.TownChannel = nil
}
