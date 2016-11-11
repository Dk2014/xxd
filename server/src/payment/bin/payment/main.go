package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
)
import (
	"core/log"
	"payment/config"
	"payment/database"
	"payment/module"
)

import (
	_ "payment/module/appstore_verify"
	_ "payment/module/google_verify"
	"payment/module/postman"
	_ "payment/module/rpc"
	_ "payment/module/web"
	_ "payment/module/wegames_push"
)

var (
	configFilePath string
)

func init() {
	flag.StringVar(&configFilePath, "conf", "payment_conf.json", "server config file")
	flag.Parse()
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	if err := config.Load(configFilePath); err != nil {
		panic(err)
	}
	log.Setup(config.CFG.LogDir, true)
	if err := database.Init(config.CFG.Database.ToConnectStr()); err != nil {
		panic(err)
	}
	var postmanFlag uint
	if !config.CFG.DisableSoha {
		postmanFlag |= postman.POSTMAN_FLAG_SOHA
	}
	if config.CFG.GooglePlayVerify.Enable {
		postmanFlag |= postman.POSTMAN_FLAG_WEGAME_GOOGLE
		//module.GoogleVerify.Init(config.CFG.GooglePlayVerify.DefaultPkg,
		//	config.CFG.GooglePlayVerify.TimeoutSecond,
		//	config.CFG.GooglePlayVerify.PublicKeyPath)
		module.GoogleVerify.InitOpenSSL(config.CFG.GooglePlayVerify.PublicKeyPath)
	}
	if config.CFG.AppStoreVerify.Enable {
		postmanFlag |= postman.POSTMAN_FLAG_WEGAME_APPLE
		module.AppStoreVerify.Init(config.CFG.AppStoreVerify.TimeoutSecond)
	}
	if config.CFG.Wegames.Enable {
		postmanFlag |= postman.POSTMAN_FLAG_WEGAME_PLATFORM
	}
	module.RPC.Init()
	module.Postman.Init(10000, map[int8]int64{
		1: 60 * 1,
		2: 60 * 2,
		3: 60 * 4,
		4: 60 * 8,
		5: 60 * 16,
	}, 1000, postmanFlag)
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "stop", "reload", "kill", "reopen":
			module.Web.Start(fmt.Sprintf("0.0.0.0:%s", config.CFG.Port), os.Args[1], config.CFG.Tls)
		default:
			panic(fmt.Errorf("unknow web server command [%s]" + os.Args[1]))
		}
	} else {
		module.Web.Start(fmt.Sprintf("0.0.0.0:%s", config.CFG.Port), "", config.CFG.Tls)
	}
}
