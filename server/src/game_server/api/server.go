package api

import (
	"core/debug"
	"core/log"
	"core/net"
	util "core/time"
	"core/tprof"
	"fmt"
	"game_server/api/protocol"
	"game_server/config"
	"game_server/mdb"
	"game_server/module"
	"os"
	"time"
)

const (
	_API_SESSION_SEND_CHAN_BUFF_ = 1024              // Send chan buffer size
	_API_SESSION_MAX_R_SIZE_     = 500 * 1024        // Max read size
	_API_SESSION_MAX_W_SIZE_     = 1024 * 1024       // Max write size
	_API_SESSION_R_TIMEOUT_      = 600 * time.Second // Read timeout
	_API_SESSION_W_TIMEOUT_      = 10 * time.Second  // Write timeout
)

// The API server.
var g_ApiServer *net.Server

var g_TimeProf *tprof.P

func init() {
	protocol.TheDebugConfig = &config.ServerCfg
}

// Start API server
func Start(network, address string) {
	var err error
	g_ApiServer, err = net.Serve(
		network, address, net.PacketN(config.API_PACKENT_HEAD_LENGTH, net.LittleEndian),
	)

	if err != nil {
		panic(err)
	}

	g_ApiServer.SetMaxReadSize(_API_SESSION_MAX_R_SIZE_)
	g_ApiServer.SetMaxWriteSize(_API_SESSION_MAX_W_SIZE_)
	g_ApiServer.SetReadTimeout(_API_SESSION_R_TIMEOUT_)
	g_ApiServer.SetWriteTimeout(_API_SESSION_W_TIMEOUT_)

	g_ApiServer.SetSessionStartHook(onSessionStart)
	g_ApiServer.SetSessionCloseHook(onSessionClose)

	g_ApiServer.Start()

	g_TimeProf = tprof.New()
}

func GetAPIProfile() {
	profFile := fmt.Sprintf("game_server_%d_%d_tprof.csv", config.ServerCfg.ServerId, util.GetNowTime())
	if file, err := os.Create(profFile); err == nil {
		g_TimeProf.Save(file)
		file.Close()
	} else {
		log.Errorf("Could't save time profile file. Error: %v", err)
	}
}

func Stop() {
	GetAPIProfile()
	g_ApiServer.Stop()
}

func onSessionStart(session *net.Session) {
	session.State = module.NewSessionState()
	session.SetRequestHandlerFunc(requestHandler)
}

func onSessionClose(session *net.Session) {
	module.DeleteSessionState(session)
}

func requestHandler(session *net.Session, msg []byte) {
	var request protocol.Request

	defer func() {
		if err := recover(); err != nil {
			state := module.State(session)
			log.Errorf(`Request Failed
Error   = %v
Session = {PlayerId:%d, Nickname:"%s", RoleId:"%d", TownId:%d}
Request = %s
Stack   =
%s`,
				err,
				state.PlayerId,
				state.PlayerNick, //state.Nickname,
				state.RoleId,     //state.RoleId,
				state.TownId,     //state.TownId,
				debug.Print(0, false, true, "    ", nil, request),
				debug.Stack(1, "    "),
			)

			/*
				// 事务内出错不断开客户端连接。可以继续游戏
				switch err.(type) {
				case mdb.TransError:
						session.Send(&notify_api.TransError_Out{})
						return
				default:
				}
			*/

			// 标记会话状态出现异常
			module.State(session).MarkPanic()
			session.Close()

			switch err.(type) {
			case mdb.TransError:
				// 事务内出错才保护
			case protocol.ProtoError:
				// 协议解析错误保护
			default:
				panic(err)
			}
		}
	}()

	request = protocol.DecodeIn(msg)

	var (
		tProcBegin time.Time
		tProcEnd   time.Time
	)

	tTransBegin := time.Now()
	state := module.State(session)
	// 统计接口成功请求次数
	state.CountAPIRequest(request)

	module.Transaction(state, (uint32(msg[0])<<8)|(uint32(msg[1])), func() {
		tProcBegin = time.Now()
		request.Process(session)
		tProcEnd = time.Now()
	})

	tTransEnd := time.Now()

	tRequest := tProcEnd.Sub(tProcBegin)

	g_TimeProf.Record(request.TypeName(), tRequest)
	g_TimeProf.Record("all request", tRequest)
	g_TimeProf.Record("all transaction", tTransEnd.Sub(tTransBegin))
}
