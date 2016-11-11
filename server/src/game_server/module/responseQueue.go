package module

import (
	"core/log"
	"core/net"
	"sync"
)

const (
	RESPONSE_QUEUE_MAX = 10
)

type ResponseQueue struct {
	sync.Mutex

	ReqCounter int64
	RspCounter int64

	queueRsp []interface{}
}

func NewResponseQueue() *ResponseQueue {
	return &ResponseQueue{
		queueRsp: make([]interface{}, 0, RESPONSE_QUEUE_MAX),

		ReqCounter: 0,
		RspCounter: 0,
	}
}

func (this *ResponseQueue) addRsp(rsp interface{}) {
	this.Lock()
	defer this.Unlock()

	this.RspCounter += 1

	this.queueRsp = append(this.queueRsp, rsp)
	currLen := len(this.queueRsp)
	if currLen > RESPONSE_QUEUE_MAX {
		this.queueRsp = this.queueRsp[currLen-RESPONSE_QUEUE_MAX:]
	}
}

func (this *ResponseQueue) ReSend(s *net.Session, clientCounter int64) bool {
	var sendQ []interface{}

	defer func() {
		for _, rsp := range sendQ {
			if rsp == nil {
				continue
			}
			switch rsp.(type) {
			case net.Response:
				s.Send(rsp.(net.Response))
			default:
				s.SendRaw(rsp.([]byte))
			}
		}
	}()

	this.Lock()
	defer this.Unlock()

	log.Infof("reSend Client(%d) Srv(%d)\n", clientCounter, this.RspCounter)
	if this.RspCounter == clientCounter {
		return true
	}

	// 客户端接收到的协议数超过服务端发送的数量
	if this.RspCounter < clientCounter {
		return false
	}

	c := len(this.queueRsp)
	diff := int(this.RspCounter - clientCounter)
	// 客户端和服务端各自统计协议数差异太大就直接断开客户端
	if c > RESPONSE_QUEUE_MAX || diff > c {
		return false
	}

	start := c - diff - 1
	if diff == c {
		start = 0
	}

	sendQ = make([]interface{}, 0, diff)
	copy(sendQ, this.queueRsp[start:])

	return true
}
