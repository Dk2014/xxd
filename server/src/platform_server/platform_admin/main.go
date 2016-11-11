package main

import (
	"core/gozd"
	"core/log"
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

	"platform_server"
)

type MyServeMux struct {
	*http.ServeMux
}

func (this MyServeMux) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Errorf("recover on MyServeHTTP: %v\n%v", r, string(debug.Stack()))
		}
	}()

	uri := req.URL.RequestURI()
	real_ip := req.Header.Get("x-real-ip")
	if real_ip == "" {
		real_ip = strings.Split(req.RemoteAddr, ":")[0]
	}
	log.Infof("[access] -- %s -- %s", real_ip, uri)

	if !isSkipCheckCookie(req.Method, uri) && !check_cookie(req) {
		resp.Header().Set("Cache-Control", "no-cache")
		setClientRedirect(resp, req, "/login")
	} else {
		//get网页请求处理
		if req.Method == "GET" &&
			uri != "/" &&
			uri != "/favicon.ico" &&
			!strings.HasPrefix(uri, "/dist") && !strings.HasPrefix(uri, "/api/") {
			setClientRedirect(resp, req, uri)
		}
	}

	this.ServeMux.ServeHTTP(resp, req)
}

func setClientRedirect(resp http.ResponseWriter, req *http.Request, uri string) {
	cookie := http.Cookie{}
	cookie.Name = "requestRoute"
	cookie.Value = uri
	cookie.Path = "/"
	http.SetCookie(resp, &cookie)

	//redirect to index, but route to uri
	req.URL.Path = "/"
}

func isSkipCheckCookie(method, uri string) bool {
	//libs
	if method == "GET" {
		if strings.HasPrefix(uri, "/dist") || uri == "/favicon.ico" {
			return true
		}
	}

	if method == "POST" {
		if strings.HasPrefix(uri, "/api/register") || strings.HasPrefix(uri, "/api/login") {
			return true
		}
	}

	return false
}

func handleListners(cl chan net.Listener) {
	for v := range cl {
		go func(l net.Listener) {
			switch reflect.ValueOf(l).Elem().FieldByName("Name").String() {
			case "http":
				handler := MyServeMux{http.NewServeMux()}
				initRouter(handler)

				srv := &http.Server{
					Handler:      handler,
					ReadTimeout:  10 * time.Second,
					WriteTimeout: 10 * time.Second,
				}

				srv.SetKeepAlivesEnabled(false)
				srv.Serve(l)
			}
		}(v)
	}
}

func usage() {
	flag.PrintDefaults()
}

func main() {
	var (
		optConfig  = flag.String("config", "", "redis server config")
		optCommand = flag.String("s", "", "send signal to a master process: stop, kill, reopen, reload")
		optHelp    = flag.Bool("h", false, "this help")
		optPort    = flag.String("p", "8889", "port")
	)

	flag.Parse()

	if *optHelp {
		usage()
		return
	}

	runtime.GOMAXPROCS(runtime.NumCPU())

	config := platform_server.LoadConfig(*optConfig)
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
		if err = platform_server.InitRedis(config.Redis, config.Apps); err != nil {
			panic(err)
		}
	}

	startJobs()

	ctx := gozd.Context{
		Hash:    "xxd_platform_admin_redis",
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

	cl := make(chan net.Listener, 1)
	go handleListners(cl)

	sig, err := gozd.Daemonize(ctx, cl) // returns channel that connects with daemon
	if err != nil {
		fmt.Println("error: ", err)
		return
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

func isDir(path string) bool {
	file, err := os.Stat(path)
	if (err == nil || os.IsExist(err)) && file.IsDir() {
		return true
	}
	return false
}
