package main

import (
	"core/fail"
	"encoding/json"
	"io/ioutil"
)

type RPCServer struct {
	Id   int
	Name string
	Addr string
}

type ServerConfig struct {
	Id     int
	Port   int
	LogDir string // 进程log路径

	InstanceId   int    // sdk初始化要求的
	ConfPath     string // xml配置文件路径
	LibPath      string //  安全sdk库目录
	TimeInterval int64  // proc_定时调用的间隔

	RPCServerList []*RPCServer // rpc服务配置列表
}

func loadJSONConfig(file string, mapstruct interface{}) []byte {
	raw, err := ioutil.ReadFile(file)
	fail.When(err != nil, "can't load file "+file)

	err = json.Unmarshal(raw, mapstruct)
	fail.When(err != nil, err)

	return raw
}
