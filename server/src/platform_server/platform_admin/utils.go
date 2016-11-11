package main

import (
	"core/log"
	"encoding/json"
	"net/http"
)

type ResponseObj struct {
	Code int         `json:"code"`
	Data interface{} `json:"data,omitempty"`
	Msg  string      `json:"msg,omitempty"`
}

func httpResponse(resp http.ResponseWriter, code int, data interface{}, msg string) {
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Content-Type", "application/json")

	obj := ResponseObj{
		Code: code,
		Data: data,
		Msg:  msg,
	}
	b, err := json.Marshal(obj)
	if err != nil {
		log.Errorf("http response json marshal error, obj:%v\n", obj)
		resp.Write([]byte(`{"code": 10004, "msg":"server inner error"}`))
		return
	}

	log.Infof("response: " + string(b))
	resp.Write(b)
}
