package clique_rpc

import (
	"core/net"
	"game_server/module"
	"sync"
)

var g_CliqueChannel *cliqueChannel = &cliqueChannel{
	channelTable: map[int64]*net.Channel{},
}

type cliqueChannel struct {
	sync.RWMutex
	channelTable map[int64]*net.Channel
}

func JoinCliqueChannel(cliqueId int64, session *net.Session) {
	if cliqueId <= 0 {
		return
	}
	g_CliqueChannel.Lock()
	defer g_CliqueChannel.Unlock()
	channel, ok := g_CliqueChannel.channelTable[cliqueId]
	if !ok {
		channel = net.NewChannel()
		g_CliqueChannel.channelTable[cliqueId] = channel
	}

	channel.Join(session)
}

func GetCliqueChannel(cliqueId int64) *net.Channel {
	g_CliqueChannel.RLock()
	defer g_CliqueChannel.RUnlock()
	return g_CliqueChannel.channelTable[cliqueId]
}

func (mod CliqueMod) DeleteCliqueChannel(cliqueId int64) {
	g_CliqueChannel.Lock()
	defer g_CliqueChannel.Unlock()
	delete(g_CliqueChannel.channelTable, cliqueId)
}

func (mod CliqueMod) LeaveCliqueChannel(cliqueId int64, session *net.Session) {
	g_CliqueChannel.Lock()
	defer g_CliqueChannel.Unlock()
	if channel, ok := g_CliqueChannel.channelTable[cliqueId]; ok {
		channel.Exit(session)
	}
}

func (mod CliqueMod) Broadcast(cliqueId int64, response net.Response) {
	g_CliqueChannel.RLock()
	defer g_CliqueChannel.RUnlock()
	if channel, ok := g_CliqueChannel.channelTable[cliqueId]; ok {
		channel.Fetch(func(session *net.Session) {
			session.Send(response)
		})
	}
}

func (mod CliqueMod) BroadcastClubhouse(cliqueId int64, response net.Response) {
	g_CliqueChannel.RLock()
	defer g_CliqueChannel.RUnlock()
	if channel, ok := g_CliqueChannel.channelTable[cliqueId]; ok {
		channel.Fetch(func(session *net.Session) {
			state := module.State(session)
			if state.InCliqueClubhouse {
				session.Send(response)
			}
		})
	}
}
