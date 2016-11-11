package rpc_common

type RPCArg interface {
	SetClientServerId(int)
	GetClientServerId() int
}

type RPCArgTag struct {
	ClientId int // 请求RPC的服务端ID
}

func (this *RPCArgTag) SetClientServerId(id int) {
	this.ClientId = id
}

func (this *RPCArgTag) GetClientServerId() int {
	return this.ClientId
}
