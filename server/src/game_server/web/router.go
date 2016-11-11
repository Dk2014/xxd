package web

import (
	"net"
	"net/http"
	"time"
)

import (
	_ "game_server/web/handler/dummy"
	_ "game_server/web/handler/mdb"
)

func StartDebugServer(addr string) {
	srv := http.Server{
		Addr:         addr,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	err = srv.Serve(ln)
	if err != nil {
		panic(err)
	}
}
