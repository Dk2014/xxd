package idip_server

import (
	"core/debug"
	"core/fail"
	"core/gozd"
	"core/log"
	"crypto/tls"
	"flag"
	"fmt"
	"idip_server/rpc"
	"net"
	"net/http"
	"os"
	"platform_server"
	"reflect"
	"runtime"
	"syscall"
	"time"
)

/*
 * usage:
 * idip_server -redis=<redis_server_ip:port>
 * idip_server -redis=<redis_server_ip:port> -cert=<https cert file> -key=<https key file>
 */

type HTTPServer struct{}

func (s HTTPServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			resp.Write([]byte(fmt.Sprintf(`{"error":400,"msg":"访问出错：%v"}`, r)))
			log.Errorf("Recovered in %v: %s", r, debug.Stack(1, "    "))
		}
	}()

	fail.When(req.Method != "POST", "not allow GET Method")
	fail.When(req.Body == nil, "request body is nil")
	err := route(req.Body, resp)
	if err != nil {
		log.Errorf("ServeHTTP return: %v", err)
	}
}

func usage() {
	flag.PrintDefaults()
}

func handleListners(cl chan net.Listener, cert string, key string) {
	for v := range cl {
		go func(l net.Listener) {

			switch reflect.ValueOf(l).Elem().FieldByName("Name").String() {

			case "http":
				handler := new(HTTPServer)
				srv := &http.Server{Handler: handler}
				srv.SetKeepAlivesEnabled(false)
				srv.Serve(l)

			case "https":
				handler := new(HTTPServer)

				srv := &http.Server{Handler: handler}

				config := &tls.Config{}
				if srv.TLSConfig != nil {
					*config = *srv.TLSConfig
				}
				if config.NextProtos == nil {
					config.NextProtos = []string{"http/1.1"}
				}

				var err error
				config.Certificates = make([]tls.Certificate, 1)
				config.Certificates[0], err = tls.LoadX509KeyPair(cert, key)
				if err != nil {
					log.Errorf("ssl cert has problem: %v", err)
					return
				}

				tlsListener := tls.NewListener(l, config)
				srv.Serve(tlsListener)
			}
		}(v)
	}
}

// exectuble entry point
func Main() {
	var (
		optCommand    = flag.String("s", "", "send signal to a master process: stop, kill, reopen, reload")
		optRedis      = flag.String("redis", "", "redis server ip:port")
		optPort       = flag.String("p", "9999", "port")
		optSSLPort    = flag.String("sslport", "9933", "sslport")
		optTlsCert    = flag.String("cert", "", "https cert")
		optTlsKey     = flag.String("key", "", "https key")
		optForeground = flag.Bool("f", false, "run in foreground")
		optLogPath    = flag.String("logpath", "", "log path")
		optHelp       = flag.Bool("h", false, "this help")
	)

	// parse arguments
	flag.Parse()
	if *optHelp {
		usage()
		return
	}

	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Errorf(`idip_server.Main
Error = %v
Stack = 
%s`,
					err,
					debug.Stack(1, "    "),
				)
			}
		}()
		ticker := time.NewTicker(time.Second * 60)
		for _ = range ticker.C {
			rpc.Remote.Start()
		}
	}()

	runtime.GOMAXPROCS(runtime.NumCPU())
	os.Setenv("GOTRACEBACK", "crash")

	if len(*optLogPath) > 0 {
		if isDir(*optLogPath) {
			log.Setup(*optLogPath, false)
		} else {
			fmt.Println("the log path does not exist, so use default log path: ./log")
			log.Setup("./log", false)
		}
	} else {
		log.Setup("./log", false)
	}

	var err error
	if len(*optRedis) > 0 {
		if err = platform_server.InitRedis(*optRedis, platform_server.Config{}.Apps); err != nil {
			fmt.Println("redis connection error:", err)
			return
		}
		if err = platform_server.InitServerList(); err != nil {
			fmt.Println("server list error:", err)
			return
		}
	}

	ctx := gozd.Context{
		Hash:    "xxd_idip_server_redis",
		Logfile: "",
		Maxfds:  syscall.Rlimit{Cur: 300000, Max: 300000},
		Command: *optCommand,
		Directives: map[string]gozd.Server{
			"http": gozd.Server{
				Network: "tcp4",
				Address: "0.0.0.0:" + *optPort,
			},
		},
	}

	if (len(*optTlsCert) + len(*optTlsKey)) > 1 {
		ctx.Directives["https"] = gozd.Server{
			Network: "tcp4",
			Address: "0.0.0.0:" + *optSSLPort,
		}
	}

	cl := make(chan net.Listener, 1)
	go handleListners(cl, *optTlsCert, *optTlsKey)

	var sig chan os.Signal

	if *optForeground {
		sig = make(chan os.Signal, 1)
	} else {
		sig, err = gozd.Daemonize(ctx, cl) // returns channel that connects with daemon
		if err != nil {
			fmt.Println("error: ", err)
			return
		}
	}

	// other initializations or config setting

	for s := range sig {
		switch s {
		case syscall.SIGHUP, syscall.SIGUSR2:
			// do some custom jobs while reload/hotupdate

		case syscall.SIGTERM:
			// do some clean up and exit
			return
		}
	}
}

func isDir(path string) bool {
	file, err := os.Stat(path)
	if (err == nil || os.IsExist(err)) && file.IsDir() {
		return true
	}
	return false
}
