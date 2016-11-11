package rpc

type RPCArg interface {
	SetClientServerId(int)
}

type RPCArgTag struct {
	ClientId int // 请求RPC的服务端ID
}

func (this *RPCArgTag) SetClientServerId(id int) {
	this.ClientId = id
}
