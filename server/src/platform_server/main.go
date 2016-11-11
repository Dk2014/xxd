package platform_server

import (
	"core/gozd"
	"core/log"
	"crypto/tls"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"reflect"
	"runtime"
	"runtime/debug"
	"runtime/pprof"
	"strings"
	"syscall"
	"time"
)

/*
 * usage:
 * platform_server -redis=<redis_server_ip:port>
 * platform_server -redis=<redis_server_ip:port> -cert=<https cert file> -key=<https key file>
 */

type HTTPServer struct{}

func (s HTTPServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	IncrMsgBy(1)

	defer func() {
		if r := recover(); r != nil {
			IncrErrBy(1)
			resp.Write([]byte(fmt.Sprintf(`{"error":400, "msg":"访问出错"}`)))
			log.Errorf("Recovered in %v: %v", r, string(debug.Stack()))
		}
	}()

	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("content-type", "application/json")

	if req.Method != "POST" {
		resp.Write([]byte(fmt.Sprintf(`{"error":400, "msg":"not allow GET Method"}`)))
		return
	}

	if req.Body == nil {
		resp.Write([]byte(fmt.Sprintf(`{"error":400, "msg":"request body is nil"}`)))
		return
	}

	clientIp := getClientIp(req)

	err := route(req.RequestURI, req.Body, resp, clientIp)
	if err != nil {
		IncrErrBy(1)
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
				srv := &http.Server{
					Handler:      handler,
					ReadTimeout:  10 * time.Second,
					WriteTimeout: 10 * time.Second,
				}

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
		optConfig     = flag.String("config", "", "redis server config")
		optCommand    = flag.String("s", "", "send signal to a master process: stop, kill, reopen, reload")
		optHelp       = flag.Bool("h", false, "this help")
		optForeground = flag.Bool("f", false, "run in foreground")
		optPort       = flag.String("p", "8888", "port")
		optSSLPort    = flag.String("sslport", "8833", "sslport")
		optTlsCert    = flag.String("cert", "", "https cert")
		optTlsKey     = flag.String("key", "", "https key")
	)

	// parse arguments
	flag.Parse()

	if *optHelp {
		usage()
		return
	}

	runtime.GOMAXPROCS(runtime.NumCPU())
	os.Setenv("GOTRACEBACK", "crash")

	config := LoadConfig(*optConfig)
	if len(config.Log) > 0 {
		if isDir(config.Log) {
			log.Setup(config.Log, false)
		} else {
			fmt.Println("the log path does not exist, so use default log path: ./log")
			log.Setup("./log", false)
		}
	} else {
		log.Setup("./log", false)
	}

	var err error
	if len(config.Redis) > 0 {
		if err = InitRedis(config.Redis, config.Apps); err != nil {
			panic(err)
		}

		if err = InitServerList(); err != nil {
			panic(err)
		}

		if err = initSetting(); err != nil {
			panic(err)
		}
	}

	InitXdlog(config.XDLog)
	go TickTo(60, log.Infof)

	ctx := gozd.Context{
		Hash:    "xxd_platform_server_redis",
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
		case syscall.SIGUSR1:
			p := pprof.Lookup("goroutine")
			if out, err := os.Create("pl.goroutine"); err == nil {
				p.WriteTo(out, 2)
			} else {
				log.Errorf("[pl.goroutine] error. %v", err)
			}

		case syscall.SIGHUP, syscall.SIGUSR2:
			// do some custom jobs while reload/hotupdate

		case syscall.SIGTERM:
			// do some clean up and exit
			return
		}
	}
}

func getClientIp(req *http.Request) string {
	real_ip := req.Header.Get("X-Forwarded-For")
	if real_ip == "" {
		real_ip = strings.Split(req.RemoteAddr, ":")[0]
	}
	return real_ip
}

func isDir(path string) bool {
	file, err := os.Stat(path)
	if (err == nil || os.IsExist(err)) && file.IsDir() {
		return true
	}
	return false
}
