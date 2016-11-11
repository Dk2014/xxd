package gameconf

import (
	"config_hub/args"
	"config_hub/httpsrv"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func Load() {
	conf := args.Conf.EventConf

	if conf == nil {
		return
	}

	httpsrv.RegistAction("/gameconf/", &gameConfAct{
		FileRoot: conf.FileRoot,
	})
}

func Unload() {
}

type gameConfAct struct {
	FileRoot string
	pattern  *regexp.Regexp
}

func (this *gameConfAct) getPattern() *regexp.Regexp {
	if this.pattern == nil {
		this.pattern = regexp.MustCompile("gs\\_([0-9]+)\\.conf")
	}
	return this.pattern
}

func (this *gameConfAct) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	match := this.getPattern().FindStringSubmatch(req.URL.Path)
	if len(match) != 2 {
		http.Error(resp, "url didn't match the format while requiring event config", http.StatusNotFound)
		return
	}
	srv_id := match[1]

	content, err := ioutil.ReadFile(this.FileRoot + srv_id + ".conf")
	if err != nil {
		log.Println(err)
		http.Error(resp, "error while reading server config file", http.StatusInternalServerError)
		return
	}

	resp.Write(content)
}
