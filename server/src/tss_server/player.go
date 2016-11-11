package main

import (
	"sync"
)

type playerTable struct {
	sync.RWMutex
	tables map[string]int64 // openid => pid
}

func NewPlayerTable() *playerTable {
	return &playerTable{tables: make(map[string]int64)}
}

func (this *playerTable) GetPlayerId(openid string) int64 {
	this.RLock()
	defer this.RUnlock()

	pid, exist := this.tables[openid]
	if !exist {
		return 0
	}

	return pid
}

func (this *playerTable) Set(openid string, playerId int64) {
	if this.GetPlayerId(openid) > 0 {
		return
	}

	this.Lock()
	defer this.Unlock()

	this.tables[openid] = playerId
}
