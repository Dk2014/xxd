package net

import "sync"

type Channel struct {
	sync.RWMutex
	sessions map[uint64]*Session
}

func NewChannel() *Channel {
	return &Channel{
		sessions: make(map[uint64]*Session),
	}
}

func (channel *Channel) Join(session *Session) {
	channel.Lock()
	defer channel.Unlock()
	channel.sessions[session.Id()] = session
}

func (channel *Channel) Exit(session *Session) {
	channel.Lock()
	defer channel.Unlock()
	delete(channel.sessions, session.Id())
}

func (channel *Channel) Fetch(callback func(*Session)) {
	channel.RLock()
	defer channel.RUnlock()
	for _, sesssion := range channel.sessions {
		callback(sesssion)
	}
}

func (channel *Channel) FetchOne(callback func(*Session), pid uint64) {
	channel.RLock()
	defer channel.RUnlock()
	session := channel.sessions[pid]
	callback(session)
}

func (channel *Channel) Count() int {
	channel.RLock()
	defer channel.RUnlock()

	return len(channel.sessions)
}
