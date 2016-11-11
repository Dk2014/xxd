package platform_server

import (
	"core/log"
	"encoding/json"
	"io/ioutil"
)

var config *Config

type App struct {
	Name    string
	DBIndex int
	Active  bool
	Cid     []int //0-ios, 1-android
}

type Config struct {
	Redis string
	Apps  map[string]*App
	Log   string
	XDLog string
}

func LoadConfig(file string) *Config {
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

	for app, val := range config.Apps {
		if !val.Active {
			delete(config.Apps, app)
		}
	}

	return config
}

func getConfig() *Config {
	return config
}

func getCidOfApp(app string, iType uint8) int {
	if len(config.Apps) == 0 {
		return 0
	}
	if v, ok := config.Apps[app]; ok {
		if iType == TYPE_MOBILE_ANDROID_WEIXIN {
			return v.Cid[1]
		}
		return v.Cid[0]
	}
	return 0
}
