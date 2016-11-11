package config

type RPCServer struct {
	Id   int
	Name string
	Addr string
}

type RPCServerConf struct {
	GSID    int
	HD      bool
	RPCIp   string
	RPCPort string
	Type    int
}
