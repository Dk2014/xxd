package config

import (
	"core/fail"
	"encoding/json"
	"flag"
	"io/ioutil"
	"net/http"
	"strings"
)

type OnlineDBConf struct {
	TableName string
	DBConfig
}

type DBConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
}

type RPCServer struct {
	Id   int
	Name string
	Addr string
}

type RPCServerConf struct {
	GSID    int
	HD      bool
	RPCIp   string
	RPCPort string
}

type RPCServerConfHttp struct {
	TimeOutSec int
	Url        string
	App        string
}

type MoneySdkConf struct {
	MoneySdkAddr       string
	MoneySdkAppId      string
	MoneySdkAppKey     string
	MoneySdkTimeOutSec int
}

type XGSDKConf struct {
	Enable       bool   //是否开启信鸽功能
	Addr         string // 信鸽地址
	AccessId     string //应用唯一标识，在提交应用时管理系统返回
	SecretKey    string
	Timeout      int
	PlatformType int // 平台类型 0-安卓 1-iOS
}

type ServerConfig struct {
	ServerId int // 服务器ID

	EnableDebugAPI     bool // 开启debug接口
	EnableGlobalServer bool //  是否是互动服务器
	GlobalServerId     int  // 互动服务器ID
	Channel            int  // 运行渠道，默认0代表腾讯 1代表越南

	GlobalServerAddr   string // 互动服务器连接地址 127.0.0.1:8080(用于玩家成功登陆游戏服后游戏服返回给客户端)
	TlogServerAddr     string // tlog服务器地址 platform.xxd.pinidea.co:6667(用于记录玩家游戏中的行为流水)
	TlogPlatformId     int32  // 平台ID，ios 0 /android 1
	DebugWebServerAddr string // RESTful 调试接口

	GameAppID  string
	GameAppKey string
	GamePort   int          // 服务端端口号
	GameDB     DBConfig     // 玩家数据库配置
	OnlineDB   OnlineDBConf // 在线表数据库配置

	ServerOpenTime int64 // 游戏开服时间戳

	LogDir      string // 进程log路径
	DBFileDir   string // mysql同步文件存储路径
	TlogFileDir string // tlog日志文件存储路径

	MoneySdk MoneySdkConf //腾讯游戏币相关接口地址

	MSDKApiAddr string //msdk api地址

	XGSdk XGSDKConf //信鸽SDK相关配置

	TssServerId         int               // tss 安全服务器id
	RPCServerList       []*RPCServer      // rpc服务配置列表
	RPCServerConfEnable bool              // rpc是否读取平台服配置
	RPCServerConfHttp   RPCServerConfHttp // rpc请求数据

	EnableXDlog  bool   //是否开启心动log
	EnableTlog   bool   //是否开启tlog
	XdlogFileDir string // xdlog日志文件存储路径

	LocalConf string // 语言配置文件地址
	CurrLocal string // 当前语言

	EnableGiftCode bool     //启用兑换码
	GiftCodeDB     DBConfig // 兑换码数据库
}

var ServerCfg ServerConfig
var ServerConf, Eventconf string

const (
	API_PACKENT_HEAD_LENGTH = 8 // RandKey(4byte)+PackSize(4byte)
)

func init() {
	flag.StringVar(&ServerConf, "conf", "xxd.conf", "server config file")
	flag.StringVar(&Eventconf, "event", "", "event config file")
	flag.Parse()

	LoadJSONConfig(ServerConf, &ServerCfg)
}

func LoadJSONConfig(file string, mapstruct interface{}) []byte {
	raw, err := []byte(""), error(nil)
	// 支持http获取配置文件
	if strings.HasPrefix(strings.ToLower(file), "http://") {
		// 配置文件读取协议为http
		resp, err := http.Get(file)
		fail.When(err != nil, "can't load file throw http get, file "+file)
		defer resp.Body.Close()
		raw, err = ioutil.ReadAll(resp.Body)
	} else {
		// 配置文件为本地文件
		raw, err = ioutil.ReadFile(file)
	}
	fail.When(err != nil, "can't load file "+file)

	err = json.Unmarshal(raw, mapstruct)
	fail.When(err != nil, err)

	return raw
}

func GetDBConfig() map[string]interface{} {
	dbcfg := make(map[string]interface{})

	dbcfg["unix_socket"] = "tcp"
	dbcfg["charset"] = "utf8mb4"

	dbcfg["host"] = ServerCfg.GameDB.Host
	dbcfg["port"] = ServerCfg.GameDB.Port
	dbcfg["uname"] = ServerCfg.GameDB.Username
	dbcfg["pass"] = ServerCfg.GameDB.Password
	dbcfg["dbname"] = ServerCfg.GameDB.DBName

	return dbcfg
}

func GetOnlineDBConfig() map[string]interface{} {
	dbcfg := make(map[string]interface{})

	dbcfg["unix_socket"] = "tcp"
	dbcfg["charset"] = "utf8mb4"

	dbcfg["host"] = ServerCfg.OnlineDB.Host
	dbcfg["port"] = ServerCfg.OnlineDB.Port
	dbcfg["uname"] = ServerCfg.OnlineDB.Username
	dbcfg["pass"] = ServerCfg.OnlineDB.Password
	dbcfg["dbname"] = ServerCfg.OnlineDB.DBName

	return dbcfg
}

func GetGiftCodeDBConfig() map[string]interface{} {
	dbcfg := make(map[string]interface{})

	dbcfg["unix_socket"] = "tcp"
	dbcfg["charset"] = "utf8mb4"

	dbcfg["host"] = ServerCfg.GiftCodeDB.Host
	dbcfg["port"] = ServerCfg.GiftCodeDB.Port
	dbcfg["uname"] = ServerCfg.GiftCodeDB.Username
	dbcfg["pass"] = ServerCfg.GiftCodeDB.Password
	dbcfg["dbname"] = ServerCfg.GiftCodeDB.DBName

	return dbcfg
}

// 实现api路由模块的接口，避免api路由模块依赖config包
func (cfg *ServerConfig) IsEnableDebugAPI() bool {
	return cfg.EnableDebugAPI
}
