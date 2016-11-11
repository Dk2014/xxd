package xdgm_server

import (
	"core/debug"
	"core/fail"
	"core/gozd"
	"core/log"
	//"crypto/sha1"
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"platform_server"
	"reflect"
	"runtime"
	"syscall"
	"xdgm_server/gift_code"
	"xdgm_server/rpc"
)

type HTTPServer struct{}

var config *Config

type Cid struct {
	Name string `json:"Name"`
	Type uint8  `json:"Type"`
}

type DBConfig struct {
	Protocol string
	Name     string
	Address  string
	User     string
	Password string
}

func (cfg *DBConfig) ToConnectStr() string {
	return fmt.Sprintf("%s:%s@%s(%s)/%s", cfg.User, cfg.Password, cfg.Protocol, cfg.Address, cfg.Name)
}

type Config struct {
	Apps              map[string]*platform_server.App `json:"Apps"`
	Cids              map[string]*Cid                 `json:"Cids"`
	GiftCodeDB        DBConfig
	PlatformServerUrl string `json:"PlatformServerUrl"`
	DefaultApp        string `json:"DefaultApp"`
}

func (s HTTPServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	defer func() {
		if r := recover(); r != nil {
			resp.Write([]byte(fmt.Sprintf(`{"error":400,"msg":"访问出错：%v"}`, r)))
			log.Errorf("Recovered in %v: %s", r, debug.Stack(1, "    "))
		}
	}()

	fail.When(req.Body == nil, "request body is nil")
	err := route(req, resp, config)
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

func Main() {
	var (
		optCommand = flag.String("s", "", "send signal to a master process: stop, kill, reopen, reload")
		optRedis   = flag.String("redis", "", "redis server ip:port")
		optPort    = flag.String("port", "9999", "port")
		optSSLPort = flag.String("sslport", "9933", "sslport")
		optTlsCert = flag.String("cert", "", "https cert")
		optTlsKey  = flag.String("key", "", "https key")
		optLogPath = flag.String("logpath", "", "log path")
		optHelp    = flag.Bool("help", false, "this help")
	)
	flag.Parse()
	if *optHelp {
		usage()
		return
	}

	rpc.Remote.Start(config.PlatformServerUrl, config.DefaultApp)

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

	config = LoadConfig("./xdgm_config.json")

	var err error

	err = gift_code.Init(config.GiftCodeDB.ToConnectStr())
	if err != nil {
		fmt.Println("init gift code db error", err)
		return
	}

	if len(*optRedis) > 0 {
		if err = platform_server.InitRedis(*optRedis, config.Apps); err != nil {
			fmt.Println("redis connection error:", err)
			return
		}
		if err = platform_server.InitServerList(); err != nil {
			fmt.Println("server list error:", err)
			return
		}
	}

	ctx := gozd.Context{
		Hash:    "xxd_xdgm_server_redis" + *optPort,
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

	sig, err := gozd.Daemonize(ctx, cl) // returns channel that connects with daemon
	if err != nil {
		fmt.Println("error: ", err)
		return
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

func LoadConfig(file string) (config *Config) {
	if config != nil {
		return config
	}

	if file == "" {
		return &Config{}
	}

	b, err := ioutil.ReadFile(file)
	if err != nil {
		log.Errorf("read config error: %s", err.Error())
		return &Config{}
	}

	config = &Config{}
	err = json.Unmarshal(b, config)
	if err != nil {
		log.Errorf("json unmarshal config error: %s", err.Error())
		return &Config{}
	}
	return config
}
