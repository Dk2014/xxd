package httpsrv

import (
	"config_hub/args"
	"net/http"
)

var actionList = map[string]http.Handler{}

func RegistAction(pattern string, handler http.Handler) http.Handler {
	_, exists := actionList[pattern]
	if exists {
		return handler
	}
	actionList[pattern] = handler
	return nil
}

func Start() error {
	// construct http handler
	srvHandler := http.NewServeMux()
	for k, v := range actionList {
		srvHandler.Handle(k, v)
	}

	// open server and work
	srv := http.Server{
		Addr:    args.Conf.Host,
		Handler: srvHandler,
	}
	err := srv.ListenAndServe()
	return err
}
