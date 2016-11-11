package syncutil

import (
	"bytes"
	"container/list"
	"core/debug"
	"core/goid"
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var (
	lockerId        int64
	globalMutex     sync.Mutex
	allLockers      = list.New()
	enableLockDebug = false
)

func EnableLockDebug() {
	enableLockDebug = true
}

func DisableLockDebug() {
	enableLockDebug = false
}

type Mutex interface {
	Lock()
	Unlock()
}

type RWMutex interface {
	Mutex
	RLock()
	RUnlock()
}

func NewMutex(l *sync.Mutex) Mutex {
	if !enableLockDebug {
		return l
	}

	w := &debugMutex{
		l:        l,
		debugger: newLockDebugger(),
	}

	globalMutex.Lock()
	element := allLockers.PushBack(w.debugger)
	globalMutex.Unlock()

	runtime.SetFinalizer(w, func(w *debugMutex) {
		globalMutex.Lock()
		allLockers.Remove(element)
		globalMutex.Unlock()
	})

	return w
}

func NewRWMutex(l *sync.RWMutex) RWMutex {
	if !enableLockDebug {
		return l
	}

	w := &debugRWMutex{
		l:        l,
		debugger: newLockDebugger(),
	}

	globalMutex.Lock()
	element := allLockers.PushBack(w.debugger)
	globalMutex.Unlock()

	runtime.SetFinalizer(w, func(w *debugRWMutex) {
		globalMutex.Lock()
		allLockers.Remove(element)
		globalMutex.Unlock()
	})

	return w
}

type debugMutex struct {
	l        *sync.Mutex
	debugger *lockDebugger
}

func (mutex *debugMutex) Lock() {
	mutex.debugger.Check()
	mutex.l.Lock()
	mutex.debugger.Update()
}

func (mutex *debugMutex) Unlock() {
	mutex.debugger.Reset()
	mutex.l.Unlock()
}

type debugRWMutex struct {
	l        *sync.RWMutex
	debugger *lockDebugger
}

func (rwmutex *debugRWMutex) Lock() {
	rwmutex.debugger.Check()
	rwmutex.l.Lock()
	rwmutex.debugger.Update()
}

func (rwmutex *debugRWMutex) Unlock() {
	rwmutex.debugger.Reset()
	rwmutex.l.Unlock()
}

func (rwmutex *debugRWMutex) RLock() {
	rwmutex.debugger.Check()
	rwmutex.l.RLock()
	rwmutex.debugger.Update()
}

func (rwmutex *debugRWMutex) RUnlock() {
	rwmutex.debugger.Reset()
	rwmutex.l.RUnlock()
}

type lockDebugger struct {
	element  *list.Element
	id       int64
	createBy int64
	createAt time.Time
	creator  string
	lockBy   int64
	lockAt   time.Time
}

func newLockDebugger() *lockDebugger {
	return &lockDebugger{
		id:       atomic.AddInt64(&lockerId, 1),
		createBy: goid.GetGoroutineId(),
		createAt: time.Now(),
		creator:  string(debug.StackPointBeginAndEnd(3, 0, "")),
	}
}

func (debugger *lockDebugger) Check() {
	if debugger.lockBy != 0 && debugger.lockBy == goid.GetGoroutineId() {
		panic("duplicate lock")
	}
}

func (debugger *lockDebugger) Reset() {
	debugger.lockAt = time.Time{}
	debugger.lockBy = 0
}

func (debugger *lockDebugger) Update() {
	debugger.lockAt = time.Now()
	debugger.lockBy = goid.GetGoroutineId()
}

func Dump() []byte {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	var buffer = new(bytes.Buffer)

	for e := allLockers.Front(); e != nil; e = e.Next() {
		l := e.Value.(*lockDebugger)
		if l.lockAt.IsZero() {
			fmt.Fprintf(
				buffer,
				"lock %d [create by goroutine: %d, create at: %s] creator:\n%s\n",
				l.id, l.createBy, l.createAt, l.creator,
			)
		} else {
			fmt.Fprintf(
				buffer,
				"lock %d [create by goroutine: %d, create at: %s, lock by goroutine: %d, lock at: %s] creator:\n%s\n",
				l.id, l.createBy, l.createAt, l.lockBy, l.lockAt, l.creator,
			)
		}
	}

	return buffer.Bytes()
}
