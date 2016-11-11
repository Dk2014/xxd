package web

import (
	"core/gozd"
	"net"
	"payment/config"
	"reflect"
	"syscall"
	//"runtime"
	"core/debug"
	"core/fail"
	"core/log"
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
	"payment/module"
	"time"
)

type HTTPServer struct{}

func (s HTTPServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("Error request: err %v \n%s \nstack\n%s\n",
				err,
				debug.Print(0, false, true, "    ", nil, req.PostForm),
				//debug.Print(0, false, true, "    ", nil, req.Form),
				debug.Stack(1, "    "))

		}
	}()
	log.Infof(" url %s \n", req.RequestURI)
	switch req.URL.Path {
	case "/debug/dummy_ok":
		dummyOk(resp, req)
	case "/debug/retry":
		retryHandler(resp, req)
	case "/ios/retry":
		//TODO
	case "/ios/verify_purchase":
		appStorePaymentHandler(resp, req)
	case "/google_play/retry":
		//TODO
	case "/google_play/verify_purchase":
		googlePlayPaymentHandler(resp, req)
	case "/wegames":
		err := req.ParseForm()
		fail.When(err != nil, err)
		wegamePlatformPaymentHandler(resp, req)
	default:
		fail.When(req.Body == nil, "request body is nil")
		paymentHandler(resp, req)
	}
}

func dummyOk(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "dummy ok")
}

func jsonResponse(w http.ResponseWriter, data interface{}) {
	var writeN int
	rawData, err := json.Marshal(data)
	fail.When(err != nil, err)
	for writeN < len(rawData) {
		n, err := w.Write(rawData[writeN:])
		fail.When(err != nil, err)
		writeN += n
	}
}

func init() {
	module.Web = WebMod{}
}

type WebMod struct{}

func (mod WebMod) Start(addr, command string, tlsCfg config.TlsConfig) {
	ctx := gozd.Context{
		Hash:       "xxd_payment",
		Logfile:    "",
		Maxfds:     syscall.Rlimit{Cur: 300000, Max: 300000},
		Command:    command,
		Directives: map[string]gozd.Server{},
	}
	if len(tlsCfg.TlsCert) > 0 && len(tlsCfg.TlsKey) > 0 && len(tlsCfg.TlsAddr) > 0 {
		ctx.Directives["https"] = gozd.Server{
			Network: "tcp4",
			Address: tlsCfg.TlsAddr,
		}
	} else {
		ctx.Directives["http"] = gozd.Server{
			Network: "tcp4",
			Address: addr,
		}
	}
	cl := make(chan net.Listener, 1)
	go handleListners(cl, tlsCfg)
	sig, err := gozd.Daemonize(ctx, cl) // returns channel that connects with daemon
	if err != nil {
		log.Errorf("Daemonize error [%v]", err)
		return
	}
	for s := range sig {
		log.Infof("recive signal %v", s)
		log.Flush()
		switch s {
		case syscall.SIGUSR1:
		case syscall.SIGHUP, syscall.SIGUSR2:
		case syscall.SIGTERM:
			return
		}
	}
}

func handleListners(cl chan net.Listener, tlsCfg config.TlsConfig) {
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
				config.Certificates[0], err = tls.LoadX509KeyPair(tlsCfg.TlsCert, tlsCfg.TlsKey)
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
