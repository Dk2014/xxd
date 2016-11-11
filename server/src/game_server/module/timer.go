package module

import (
	"core/debug"
	"core/log"
	"game_server/mdb"
	"time"
)

type Timer struct {
	timers map[int]*time.Timer
	state  *SessionState
}

func NewTimer(state *SessionState) *Timer {
	return &Timer{timers: make(map[int]*time.Timer), state: state}
}

func (this *Timer) Start(tag int, delay time.Duration, worker func(state *SessionState)) {
	this.Stop(tag)
	this.timers[tag] = time.AfterFunc(delay, func() {

		defer func() {
			if err := recover(); err != nil {
				log.Errorf(`SessionState Timer
Error   = %v
Session = {PlayerId:%d, Nickname:"%s", RoleId:"%d", TownId:%d}
TimerTag = %d
Stack   =
%s`,
					err,
					this.state.PlayerId,
					this.state.PlayerNick, //state.Nickname,
					this.state.RoleId,     //state.RoleId,
					this.state.TownId,     //state.TownId,
					tag,
					debug.Stack(1, "    "),
				)
			}
		}()

		Transaction(this.state, mdb.TRANS_TAG_SESSION_STATE_TIMER, func() {
			worker(this.state)
		})
	})
}

func (this *Timer) Stop(tag int) {
	if t, ok := this.timers[tag]; ok {
		t.Stop()
		delete(this.timers, tag)
	}
}

func (this *Timer) StopAll() {
	for _, t := range this.timers {
		t.Stop()
	}
	this.timers = make(map[int]*time.Timer)
}
