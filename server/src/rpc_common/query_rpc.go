package rpc_common

type Args_QueryPid struct {
	RPCArgTag
	OpenId string
}

type Reply_QueryPid struct {
	Pid int64
}
