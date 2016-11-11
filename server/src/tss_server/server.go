package main

import (
	"core/debug"
	"core/fail"
	"core/log"
	"flag"
	"fmt"
	"net"
	"net/rpc"
	"runtime"
	"tss_server/tss_sdk_go"
)

var (
	serverConf string
	serverCfg  ServerConfig

	g_players  *playerTable = NewPlayerTable()
	g_rpc      *rpcClient   = NewRPCClient()
	g_syncChan chan tss_sdk_go.SyncUnit
)

//callback
type callback struct {
}

//send data to channel, let rpc return it
func (this *callback) Send_Data_To_Client(openid string, plat_id int, anti_data []byte) int {
	playerId := g_players.GetPlayerId(openid)
	if playerId == 0 {
		log.Errorf("Send_Data_To_Client:%v can't found player\n", openid)
		return tss_sdk_go.TSS_SDK_PROC_INTERNAL_ERR
	}

	// todo get gsid call rpc
	var reply int
	g_rpc.Call(int(playerId>>32), "RemoteServe.TssSendData", &PkgInfo{PlayerId: playerId, OpenId: openid, Plat_Id: plat_id, Anti_Data: anti_data}, &reply)

	return tss_sdk_go.TSS_SDK_PROC_OK
}

func init() {
	flag.StringVar(&serverConf, "conf", "tss.conf", "tss server config file")
	flag.Parse()
	loadJSONConfig(serverConf, &serverCfg)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.Setup(serverCfg.LogDir, true)

	go func() {
		l, e := net.Listen("tcp", fmt.Sprintf(":%d", serverCfg.Port))
		fail.When(e != nil, e)

		rpcServer := rpc.NewServer()
		rpcServer.Register(new(RpcOper))
		go func() {
			defer func() {
				if err := recover(); err != nil {
					log.Errorf(`tss_server.main
	Error = %v
	Stack = 
	%s`,
						err,
						debug.Stack(1, "    "),
					)
				}
			}()
			rpcServer.Accept(l)
		}()

		log.Info("start rpc server")
	}()

	runtime.LockOSThread()

	//init tss_sdk
	tss_sdk_go.CallBack = &callback{}
	g_syncChan = tss_sdk_go.Init(serverCfg.ConfPath, serverCfg.LibPath, serverCfg.TimeInterval, uint(serverCfg.InstanceId))
EXIT:
	for {
		select {
		case unit, ok := <-g_syncChan:
			if !ok {
				break EXIT
			}
			switch unit.Act {
			case tss_sdk_go.TSS_ACTION_PROC:
				tss_sdk_go.CallProc()

			case tss_sdk_go.TSS_ACTION_USERLOGIN:
				user_info := unit.Data.(*UserInfo)
				tss_sdk_go.UserLogin(user_info.OpenId, user_info.Plat_Id, user_info.Client_Ip, user_info.Client_Ver, user_info.Role_Name)

			case tss_sdk_go.TSS_ACTION_USERLOGOUT:
				user_info := unit.Data.(*UserInfo)
				tss_sdk_go.UserLogout(user_info.OpenId, user_info.Plat_Id)

			case tss_sdk_go.TSS_ACTION_RECV:
				pkg_info := unit.Data.(*PkgInfo)
				tss_sdk_go.RecvData(pkg_info.OpenId, pkg_info.Plat_Id, pkg_info.Anti_Data)
			}
		}
	}

	log.Info("tss server shutdown.")
}
