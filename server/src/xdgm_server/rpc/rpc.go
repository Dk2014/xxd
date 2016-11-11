package rpc

import (
	"bytes"
	"core/debug"
	"core/log"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/rpc"
	"sync"
	"time"
	"xdgm_server/config"
)

var Remote *RPC = NewRPC()

type RPC struct {
	serviceAPIs []interface{}

	clients      map[int]*rpc.Client
	clientsMutex sync.Mutex

	servers map[int]*config.RPCServer
}

func NewRPC() *RPC {
	return &RPC{clients: make(map[int]*rpc.Client)}
}

func (this *RPC) Start(url, app string) {
	time.AfterFunc(time.Second*60, func() {
		defer func() {
			if err := recover(); err != nil {
				log.Errorf(`xdgm_server.Main
Error = %v
Stack = 
%s`,
					err,
					debug.Stack(1, "    "),
				)
			}
		}()
		defer this.Start(url, app)
		this.httpUpdateServerInfo(url, app)
	})
}

func (this *RPC) Call(serverId int, serviceAPI string, args RPCArg, reply interface{}, callback func(error)) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf(`RPC failed
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

	err := this.callclient(0, serverId, serviceAPI, args, reply)

	if err != nil {
		log.Errorf(`RPC error
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

	callback(err)
}

func (this *RPC) BatchCall(serviceAPI string, serverIds []int, args []RPCArg, replys []interface{}, callback func([]error)) {
	wg := new(sync.WaitGroup)

	errors := make([]error, len(serverIds))

	for i, serverId := range serverIds {
		wg.Add(1)
		this.Call(serverId, serviceAPI, args[i], replys[i], func(err error) {
			errors[i] = err
			wg.Done()
		})
	}

	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Errorf(`RPC BatchCall failed
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
		callback(errors)
		wg.Wait()
	}()
}

func (this *RPC) GetClient(serverId int) (*rpc.Client, error) {
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
		return nil, errors.New(fmt.Sprintf(`RPC server %d not exists`, serverId))
	}

	conn, err := net.DialTimeout("tcp", info.Addr, time.Second*3)
	if err != nil {
		return nil, errors.New(fmt.Sprintf(`RPC server %d connect failed`, serverId))
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

func (this *RPC) httpUpdateServerInfo(url, app string) {
	reqURL := url + "/gserverall"
	// Add rpc server list from http response
	postValues := make(map[string]interface{})
	postValues["App"] = app
	postDataBytes, _ := json.Marshal(postValues)
	postBytesReader := bytes.NewReader(postDataBytes)
	req, _ := http.NewRequest("POST", reqURL, postBytesReader)
	req.Header.Add("Content-Type", "application/json")

	DefaultClient := http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				deadline := time.Now().Add(10 * time.Second)
				c, err := net.DialTimeout(netw, addr, time.Second*5)
				if err != nil {
					return nil, err
				}
				c.SetDeadline(deadline)
				return c, nil
			},
		},
	}
	resp, err := DefaultClient.Do(req)
	if err != nil {
		log.Error(fmt.Sprintf("error: %v, request: %s\n", err, reqURL))
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(fmt.Sprintf("error: %v, request: %s\n", err, reqURL))
		panic(err)
	}
	list := make(map[string]config.RPCServerConf)
	err = json.Unmarshal(body, &list)
	if err != nil {
		log.Error(fmt.Sprintf("解析json失败: %v [%v], %s\n", err, body, reqURL))
		panic(err)
	}
	newServers := make(map[int]*config.RPCServer)
	// 服务器地址发生变化
	for _, s := range list {
		addr := fmt.Sprintf("%v:%v", s.RPCIp, s.RPCPort)
		newServers[s.GSID] = &config.RPCServer{
			Id:   s.GSID,
			Name: app,
			Addr: addr,
		}

		if oldS, ok := this.servers[s.GSID]; !ok || oldS.Addr == addr {
			continue
		}

		if c, ok := this.clients[s.GSID]; ok {
			c.Close()
			delete(this.clients, s.GSID)
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

func (this *RPC) callclient(times, serverId int, serviceAPI string, args RPCArg, reply interface{}) error {
	client, err := this.GetClient(serverId)
	if err == nil {
		args.SetClientServerId(config.XDGM_SERVER_ID)
		err = client.Call(serviceAPI, args, reply)
		if err == io.ErrUnexpectedEOF || err == rpc.ErrShutdown {
			client.Close()
			this.removeClient(serverId)
			if err == rpc.ErrShutdown && times < 3 {
				this.callclient(times, serverId, serviceAPI, args, reply)
			}
		}
	}
	return err
}
