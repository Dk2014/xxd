package module

import (
	"fmt"
	"time"
)

var Tracer *trace = &trace{}

type trace struct {
	playerId int64  // 追踪目标,玩家id
	tag      string // 追踪唯一标记

	sTime int64 // 开始追踪的时间
}

func (t *trace) SetTarget(playerId int64) {
	t.playerId = playerId
}

func (t *trace) Start(state *SessionState, tag string) {
	if t.playerId == state.PlayerId {
		t.tag = tag
		t.sTime = time.Now().UnixNano()
		fmt.Println("trace: ", tag)
	}
}

func (t *trace) Echo(format string, v ...interface{}) {
	if t.playerId > 0 {
		fmt.Println("  >>> ", fmt.Sprintf(format, v...))
	}
}

func (t *trace) Stop() {
	if t.playerId > 0 {
		costTime := float64(time.Now().UnixNano()-t.sTime) / float64(time.Microsecond)
		fmt.Printf("stop trace: %v μs\n\n", costTime)
	}
}
