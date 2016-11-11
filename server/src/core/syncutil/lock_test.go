package syncutil

import (
	"sync"
	"testing"
)

func Test_Lock(t *testing.T) {
	a := NewMutex(new(sync.Mutex))
	b := NewMutex(new(sync.Mutex))

	t.Logf("\n%s", Dump())

	a.Lock()
	a.Unlock()

	b.Lock()
	b.Unlock()

	x := func() {
		a.Lock()
		defer func() {
			if err := recover(); err != nil {
				t.Log(err)
			}
			a.Unlock()
		}()
		a.Lock()
	}

	x()

	a.Lock()
	a.Unlock()
}
