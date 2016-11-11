package arena_rpc

import (
	"game_server/mdb"
	"sync"
)

var (
	rankTable *globalRank = &globalRank{ranks: make(map[int64]int32), maxRank: 0}
)

type globalRank struct {
	sync.RWMutex
	ranks   map[int64]int32 // map[playerId]rank
	maxRank int32           // 当前最大排名
}

func (this *globalRank) loadPlayerRank(pid int64, rank int32) {
	this.Lock()
	defer func() {
		this.Unlock()
	}()

	this.maxRank += 1
	this.ranks[pid] = rank
}

func (this *globalRank) incRank() int32 {
	this.Lock()
	defer func() {
		this.Unlock()
	}()

	this.maxRank += 1
	return this.maxRank
}

func (this *globalRank) updatePlayerRank(pid int64, rank int32) {
	this.Lock()
	defer func() {
		this.Unlock()
	}()

	this.ranks[pid] = rank
}

func (this *globalRank) getPlayerRank(pid int64) int32 {
	this.RLock()
	defer func() {
		this.RUnlock()
	}()

	return this.ranks[pid]
}

// 玩家开启比武功能时调用
func addPlayerRank(db *mdb.Database, pid int64) (rank int32) {
	rank = rankTable.incRank()
	db.Insert.GlobalArenaRank(&mdb.GlobalArenaRank{
		Rank: rank,
		Pid:  pid,
	})
	return
}

func updatePlayerRank(db *mdb.Database, pid int64, rank int32) {
	db.Update.GlobalArenaRank(&mdb.GlobalArenaRank{
		Rank: rank,
		Pid:  pid,
	})
}

func getPlayerRank(pid int64) int32 {
	return rankTable.getPlayerRank(pid)
}
