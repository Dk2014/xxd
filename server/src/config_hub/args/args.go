package args

import (
	"encoding/json"
	"flag"
	"io/ioutil"
)

var Conf = &CmdConf{}

func Parse() error {
	var conf_file string

	flag.StringVar(&conf_file, "conf", "", "config file")
	flag.Parse()

	conf, err := ioutil.ReadFile(conf_file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(conf, Conf)
	if err != nil {
		return err
	}
	return nil
}

type CmdConf struct {
	Host      string     `json:"host"`
	GiftCode  *GiftCode  `json:"gift_code"`
	GameConf  *GameConf  `json:"game_conf"`
	EventConf *EventConf `json:"event_conf"`
}

type GiftCode struct {
	SqlUser string `json:"sql_user"`
	SqlPass string `json:"sql_pass"`
	SqlPort string `json:"sql_port"`
	SqlAddr string `json:"sql_addr"`
	SqlName string `json:"sql_name"`
}

type GameConf struct {
	FileRoot string `json:"file_root"`
}

type EventConf struct {
	FileRoot string `json:"file_root"`
}
