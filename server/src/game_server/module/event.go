package module

import (
	"core/net"
	"game_server/mdb"
)

var (
	PrepareStoreEvent *SessionEvent = &SessionEvent{eventList: make([]func(*net.Session), 0)}

	DisposeEvent *StateEvent = &StateEvent{eventList: make([]func(*SessionState), 0)}
)

type SessionEvent struct {
	eventList []func(*net.Session)
}

func (this *SessionEvent) Regisiter(fn func(*net.Session)) {
	this.eventList = append(this.eventList, fn)
}

func (this *SessionEvent) execute(session *net.Session) {
	for _, fn := range this.eventList {
		fn(session)
	}
}

func (this *SessionEvent) SafeExecute(session *net.Session) {
	Transaction(State(session), mdb.TRANS_TAG_PREPARE_STORE_EVENT, func() {
		this.execute(session)
	})
}

type StateEvent struct {
	eventList []func(*SessionState)
}

func (this *StateEvent) Regisiter(fn func(*SessionState)) {
	this.eventList = append(this.eventList, fn)
}

func (this *StateEvent) execute(state *SessionState) {
	for _, fn := range this.eventList {
		fn(state)
	}
}

func (this *StateEvent) SafeExecute(state *SessionState) {
	this.execute(state)
}
