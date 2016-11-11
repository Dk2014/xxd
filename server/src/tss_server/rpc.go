package main

import (
	"core/log"
	"io"
	"net"
	"net/rpc"
	"sync"
	"time"
)

type rpcClient struct {
	sync.RWMutex
	clients map[int]*rpc.Client
}

func NewRPCClient() *rpcClient {
	return &rpcClient{clients: make(map[int]*rpc.Client)}
}

func (this *rpcClient) get(gsId int) *rpc.Client {
	this.RLock()
	c, ok := this.clients[gsId]
	this.RUnlock()

	if ok {
		return c
	}

	this.Lock()
	defer this.Unlock()
	for _, v := range serverCfg.RPCServerList {
		if v.Id == gsId {
			conn, err := net.DialTimeout("tcp", v.Addr, time.Second*3)
			if err != nil {
				log.Errorf("can't connect to gsid = %v, err: %v", gsId, err)
				return nil
			}

			client := rpc.NewClient(conn)
			this.clients[v.Id] = client

			return client
		}
	}

	return nil
}

func (this *rpcClient) Call(serverId int, remoteAPI string, args, reply interface{}) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Errorf(`RPC-Call panic
Error   = %v
remoteAPI = %v
`,
					err,
					remoteAPI,
				)
			}
		}()

		client := this.get(serverId)
		if client == nil {
			return
		}

		err := client.Call(remoteAPI, args, reply)
		if err == io.ErrUnexpectedEOF || err == rpc.ErrShutdown {
			log.Errorf("rpc-call %s error. %v", remoteAPI, err)
			client.Close()

			this.Lock()
			defer this.Unlock()
			delete(this.clients, serverId)
		}
	}()
}
