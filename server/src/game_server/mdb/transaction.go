package mdb

import "core/syncutil"
import "sync/atomic"
import "core/goid"

import (
	"time"
	//"core/debug"
	//corelog "core/log"
)

const (
	TRANS_TYPE_MYSQL = 0
	TRANS_TYPE_TLOG  = 1
	TRANS_TYPE_XDLOG = 2
)

type TransLog interface {
	GetTransType() int8
	InvokeHook()
	CommitToFile(*SyncFile)
	CommitToMySql() error
	CommitToTLog()
	CommitToXdLog()
	Rollback()
	Free()
	SetEventId(time.Time)
}

type TransError struct {
	Message interface{}
}

type transController struct {
	lock        syncutil.Mutex
	lockOwnerId int64
	transLog    []TransLog
	commiter    *fileCommitter
}

func newTransController(lock syncutil.Mutex, commiter *fileCommitter) *transController {
	return &transController{
		lock:        lock,
		lockOwnerId: -1,
		commiter:    commiter,
	}
}

// 在一个内存数据库事务中执行一段代码，执行过程出错，内存数据将被回滚
// 所有事务被顺序执行，等同于MySql串行事务隔离界别
func (tc *transController) Transaction(transId uint32, work func()) {
	// 检查嵌套事务
	id := goid.GetGoroutineId()
	if atomic.CompareAndSwapInt64(&tc.lockOwnerId, id, id) {
		panic("nested transaction")
	}

	// 数据加锁，这个地方取出锁作为局部变量，目的是防止事务过程中锁被覆盖
	lock := tc.lock
	lock.Lock()
	atomic.StoreInt64(&tc.lockOwnerId, id)
	defer func() {
		atomic.StoreInt64(&tc.lockOwnerId, -1)
		lock.Unlock()
	}()

	// 事务控制
	defer func() {
		if err := recover(); err == nil {
			tc.Commit(transId)
		} else {
			tc.Rollback()
			panic(TransError{err})
		}
	}()

	work()
}

func (tc *transController) Commit(transId uint32) {
	if len(tc.transLog) > 0 {
		for _, log := range tc.transLog {
			log.InvokeHook()
		}

		tc.commiter.Commit(transId, tc.transLog)
		tc.transLog = nil
	}
}

func (tc *transController) Rollback() {
	if len(tc.transLog) > 0 {
		for i := len(tc.transLog) - 1; i >= 0; i-- {
			tc.transLog[i].Rollback()
		}
		tc.transLog = nil
	}
}

func (tc *transController) AddTransLog(log TransLog) {
	//corelog.Debugf("AddTransLog %p, uncommit translog size %d\n", tc, len(tc.transLog))
	tc.transLog = append(tc.transLog, log)
}
