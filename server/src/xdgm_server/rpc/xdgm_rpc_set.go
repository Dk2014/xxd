package rpc

import (
	"game_server/mdb"
)

/*
	修改充值返利规则
*/
type Args_XdgmSetPaymentsPresent struct {
	RPCArgTag
	Rule      string
	BeginTime int64
	EndTime   int64
}

type Reply_XdgmSetPaymentsPresent struct {
}

func RemoteXdgmSetPaymentsPresent(sid int, rule string, begintime, endtime int64, callback func(*Reply_XdgmSetPaymentsPresent, error)) {
	reply := &Reply_XdgmSetPaymentsPresent{}
	args := &Args_XdgmSetPaymentsPresent{
		Rule:      rule,
		BeginTime: begintime,
		EndTime:   endtime,
	}
	Remote.Call(sid, mdb.RPC_Remote_XdgmSetPaymentsPresent, args, reply, func(err error) {
		callback(reply, err)
	})
}
