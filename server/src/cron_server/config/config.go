package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var SCFG *CronServerConfig

type DatabaseConfig struct {
	Protocol string
	DBName   string
	Host     string
	Port     int
	User     string
	Password string
}

type OnlineDBCofing struct {
	Protocol  string
	DBName    string
	TableName string
	Host      string
	Port      int
	User      string
	Password  string
}

type ServerConfig struct {
	Sid      int
	Database DatabaseConfig
}

func (scfg *DatabaseConfig) ToConnectStr() string {
	return fmt.Sprintf("%s:%s@%s(%s:%d)/%s", scfg.User, scfg.Password, scfg.Protocol, scfg.Host, scfg.Port, scfg.DBName)
}

func (scfg *OnlineDBCofing) ToConnectStrOnline() string {
	return fmt.Sprintf("%s:%s@%s(%s:%d)/%s", scfg.User, scfg.Password, scfg.Protocol, scfg.Host, scfg.Port, scfg.DBName)
}

type CronServerConfig struct {
	LogDir       string
	XdlogFileDir string
	OnlineDB     OnlineDBCofing
	ServerInfo   []ServerConfig
}

func Load(configFilePath string) error {
	raw, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return err
	}
	SCFG = &CronServerConfig{}
	err = json.Unmarshal(raw, SCFG)
	if err != nil {
		return err
	}
	return nil
}
