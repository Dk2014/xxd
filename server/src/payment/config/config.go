package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var CFG *ServerConfig

type ProductConfig struct {
	GameMoney        int64
	PresentGameMoney int64
	IsMonthCard      bool
	Cost             float64
}

type WegamesConfig struct {
	Enable       bool
	GameCode     string
	IsProduction bool
	PrivateKey   string
	ProductUrl   string
	DevUrl       string
}

type GooglePlayVeriyConfig struct {
	Enable              bool
	DefaultPkg          string
	TimeoutSecond       int
	PublicKeyPath       string
	NewPaymentQueueSize int64
}

type AppStoreVerifyConfig struct {
	Enable              bool
	TimeoutSecond       int
	IsProduction        bool
	NewPaymentQueueSize int64
	SandServerList      []int
}

type DatabaseConfig struct {
	Protocol string
	Name     string
	Address  string
	User     string
	Password string
	Params   map[string]string
}

func (cfg *DatabaseConfig) ToConnectStr() string {
	return fmt.Sprintf("%s:%s@%s(%s)/%s", cfg.User, cfg.Password, cfg.Protocol, cfg.Address, cfg.Name)
}

type RPCServer struct {
	Id   int
	Name string
	Addr string
}

type TlsConfig struct {
	TlsAddr string
	TlsCert string
	TlsKey  string
}

type ServerConfig struct {
	ServerId           int64
	LogDir             string
	RPCPort            string
	Port               string
	Debug              bool
	DisableSoha        bool
	DisableValidate    bool
	PrivateKey         string
	EnhancedValidate   bool
	EnhancedPrivateKey string
	ProductInfo        map[string]*ProductConfig
	Database           DatabaseConfig
	GooglePlayVerify   GooglePlayVeriyConfig
	AppStoreVerify     AppStoreVerifyConfig
	Wegames            WegamesConfig
	Tls                TlsConfig

	RPCServerList []*RPCServer

	EnableHTTPConfig   bool
	RPCServerConfigURL string
	RPCServerConfigApp string
}

func Load(configFilePath string) error {
	raw, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return err
	}
	CFG = &ServerConfig{}
	err = json.Unmarshal(raw, CFG)
	if err != nil {
		return err
	}
	return nil
}
