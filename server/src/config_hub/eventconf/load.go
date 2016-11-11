package eventconf

import (
	"config_hub/args"
	"config_hub/httpsrv"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
)

const DEFAULT_EVENT_FILE = "default.json"

func Load() {
	conf := args.Conf.EventConf

	if conf == nil {
		return
	}

	httpsrv.RegistAction("/eventconf/", &eventConfAct{
		FileRoot: conf.FileRoot,
	})
}

func Unload() {
}

type eventConfAct struct {
	FileRoot string
	pattern  *regexp.Regexp
}

func (this *eventConfAct) getPattern() *regexp.Regexp {
	if this.pattern == nil {
		this.pattern = regexp.MustCompile("event\\_([0-9]+)\\.json")
	}
	return this.pattern
}

func (this *eventConfAct) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	match := this.getPattern().FindStringSubmatch(req.URL.Path)
	if len(match) != 2 {
		http.Error(resp, "url didn't match the format while requiring event config", http.StatusNotFound)
		return
	}
	srv_id := match[1]

	content, err := ioutil.ReadFile(this.FileRoot + srv_id + ".json")
	if err != nil {
		if os.IsNotExist(err) {
			content, err = ioutil.ReadFile(this.FileRoot + DEFAULT_EVENT_FILE)
			if err != nil {
				log.Println(err)
				http.Error(resp, "error while reading default event json", http.StatusInternalServerError)
				return
			}
		} else {
			log.Println(err)
			http.Error(resp, "error while reading event json", http.StatusInternalServerError)
			return
		}
	}

	resp.Write(content)
}
