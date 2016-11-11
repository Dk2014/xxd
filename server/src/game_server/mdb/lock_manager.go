package mdb

import (
	"core/syncutil"
	"sync"
)

type lockManager struct {
	mutex sync.RWMutex
	locks map[int64]syncutil.Mutex
}

func newLockManager() *lockManager {
	return &lockManager{
		locks: make(map[int64]syncutil.Mutex),
	}
}

func (lm *lockManager) Get(pid int64) syncutil.Mutex {
	lm.mutex.RLock()
	lock, exists := lm.locks[pid]
	lm.mutex.RUnlock()

	if exists {
		return lock
	}

	lm.mutex.Lock()
	defer lm.mutex.Unlock()

	// 二次检查，防止锁升级的时候边界情况发生，重复创建lock
	lock, exists = lm.locks[pid]

	if exists {
		return lock
	}

	lock = syncutil.NewMutex(new(sync.Mutex))
	lm.locks[pid] = lock

	return lock
}
