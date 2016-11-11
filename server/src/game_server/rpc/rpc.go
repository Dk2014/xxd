package rpc

import (
	"core/debug"
	"core/log"
	"errors"
	"fmt"
	"game_server/config"
	"game_server/mdb"
	"io"
	"net"
	"net/rpc"
	//"net/url"
	//"strings"
	"sync"
	"time"
	"utils/rpc_conf"
)

type RPC struct {
	serviceAPIs []interface{}

	clients      map[int]*rpc.Client
	clientsMutex sync.Mutex

	servers map[int]*config.RPCServer
}

func NewRPC() *RPC {
	return &RPC{clients: make(map[int]*rpc.Client)}
}

func (this *RPC) GetRPCServerIds() []int {
	var list []int
	for id, _ := range this.servers {
		list = append(list, id)
	}

	return list
}

func (this *RPC) Start(addr string) {
	this.RefreshRPCServerList()

	listener, err := net.Listen("tcp4", addr)
	if err != nil {
		panic(err)
	}

	rpcServer := rpc.NewServer()

	for _, api := range this.serviceAPIs {
		if err = rpcServer.Register(api); err != nil {
			panic(err)
		}
	}

	go rpcServer.Accept(listener)
}

func (this *RPC) Register(api interface{}) {
	this.serviceAPIs = append(this.serviceAPIs, api)
}

func (this *RPC) Serve(serviceAPI string, args RPCArg, transId uint32, callback func() error) (err error) {
	defer func() {
		if err2 := recover(); err2 != nil {

			log.Errorf(`RPC-Serve panic
Error    = %v
ClientId = %d
serviceAPI  = %v
Args     = %s
Stack    = %s`,
				err2,
				args.GetClientServerId(),
				serviceAPI,
				debug.Print(0, false, true, "    ", nil, args),
				debug.Stack(1, "    "),
			)

			err = errors.New(fmt.Sprintf("%v", err2))
		}
	}()

	if callback != nil {
		mdb.Transaction(transId, func() {
			err = callback()
		})
	}

	return
}

func (this *RPC) Call(serverId int, serviceAPI string, args RPCArg, reply interface{}, transId uint32, callback func(error)) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Errorf(`RPC-Call panic
Error   = %v
serviceAPI = %v
Args    = %s
Result  = %s
Stack   =
%s`,
					err,
					serviceAPI,
					debug.Print(0, false, true, "    ", nil, args),
					debug.Print(0, false, true, "    ", nil, reply),
					debug.Stack(1, "    "),
				)
			}
		}()

		client, err := this.getClient(serverId)
		if err == nil {
			args.SetClientServerId(config.ServerCfg.ServerId)
			err = client.Call(serviceAPI, args, reply)
			if err == io.ErrUnexpectedEOF || err == rpc.ErrShutdown {
				client.Close()
				this.removeClient(serverId)
			}
		}

		if err != nil {
			log.Errorf(`RPC-Call found error when excute client.Call
			Error   	= %v
			serviceAPI 	= %v
			Args    = %s
			Result  = %s`,
				err,
				serviceAPI,
				debug.Print(0, false, true, "    ", nil, args),
				debug.Print(0, false, true, "    ", nil, reply),
			)
		}

		if callback != nil {
			mdb.Transaction(transId, func() {
				callback(err)
			})
		}
	}()
}

func (this *RPC) BatchCall(serviceAPI string, serverIds []int, args []RPCArg, replys []interface{}, transId uint32, callback func([]error)) {
	wg := new(sync.WaitGroup)

	errors := make([]error, len(serverIds))

	for i, serverId := range serverIds {
		wg.Add(1)
		this.Call(serverId, serviceAPI, args[i], replys[i], transId, func(err error) {
			errors[i] = err
			wg.Done()
		})
	}

	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Errorf(`RPC-BatchCall failed
Error     	= %v
serviceAPI  = %v
ServerIds = %s
Args      = %s
Result    = %s
Stack     =
%s`,
					err,
					serviceAPI,
					debug.Print(0, false, true, "    ", nil, serverIds),
					debug.Print(0, false, true, "    ", nil, args),
					debug.Print(0, false, true, "    ", nil, replys),
					debug.Stack(1, "    "),
				)
			}
		}()

		wg.Wait()

		if callback != nil {
			mdb.Transaction(transId, func() {
				callback(errors)
			})

		}
	}()
}

func (this *RPC) getClient(serverId int) (*rpc.Client, error) {
	this.clientsMutex.Lock()
	client, ok := this.clients[serverId]
	this.clientsMutex.Unlock()

	if ok {
		return client, nil
	}

	this.clientsMutex.Lock()
	info, ok := this.servers[serverId]
	this.clientsMutex.Unlock()

	if !ok {
		return nil, errors.New(fmt.Sprintf(`RPC::getClient can't found rpc-server %d`, serverId))
	}

	conn, err := net.DialTimeout("tcp", info.Addr, time.Second*3)
	if err != nil {
		return nil, errors.New(fmt.Sprintf(`RPC::getClient can't connect rpc-server %d`, serverId))
	}

	client = rpc.NewClient(conn)

	// 成功连上的时候做第二遍检查，防止重复连接
	this.clientsMutex.Lock()
	if oldClient, ok := this.clients[serverId]; ok {
		client.Close()
		client = oldClient
	} else {
		this.clients[serverId] = client
	}
	this.clientsMutex.Unlock()

	return client, nil
}

func (this *RPC) removeClient(serverId int) {
	this.clientsMutex.Lock()
	defer this.clientsMutex.Unlock()
	delete(this.clients, serverId)
}

func (this *RPC) RefreshRPCServerList() {
	var servers []*config.RPCServer

	if len(config.ServerCfg.RPCServerList) > 0 {
		for _, s := range config.ServerCfg.RPCServerList {
			servers = append(servers, s)
		}
	} else {
		// 默认没有配置文件就把自己作为rpc服务
		servers = append(servers, &config.RPCServer{
			Id:   config.ServerCfg.ServerId,
			Name: "",
			Addr: fmt.Sprintf("127.0.0.1:%d", config.ServerCfg.GamePort+1),
		})
	}

	if config.ServerCfg.RPCServerConfEnable {
		url := config.ServerCfg.RPCServerConfHttp.Url
		app := config.ServerCfg.RPCServerConfHttp.App
		for _, rPCServerConf := range rpc_conf.RPCList(url, app) {
			servers = append(servers, &config.RPCServer{
				Id:   rPCServerConf.GSID,
				Name: "",
				Addr: rPCServerConf.RPCIp + ":" + rPCServerConf.RPCPort,
			})
		}
	}

	this.clientsMutex.Lock()
	defer this.clientsMutex.Unlock()

	newServers := make(map[int]*config.RPCServer)
	// 服务器地址发生变化
	for _, s := range servers {
		newServers[s.Id] = s

		if oldS, ok := this.servers[s.Id]; !ok || oldS.Addr == s.Addr {
			continue
		}

		if c, ok := this.clients[s.Id]; ok {
			c.Close()
			delete(this.clients, s.Id)
		}
	}

	// 服务器发生变化
	for id, _ := range this.servers {
		if _, ok := newServers[id]; ok {
			continue
		}

		if c, ok := this.clients[id]; ok {
			c.Close()
			delete(this.clients, id)
		}
	}

	// 更新服务器列表
	this.servers = newServers
}
