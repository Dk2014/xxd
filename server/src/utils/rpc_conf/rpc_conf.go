package rpc_conf

import (
	"bytes"
	"core/fail"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type RPCServer struct {
	GSID    int
	HD      bool
	RPCIp   string
	RPCPort string
}

func RPCList(url, app string) map[string]RPCServer {
	postValues := make(map[string]interface{})
	postValues["App"] = app
	postDataBytes, err := json.Marshal(postValues)
	fail.When(err != nil, err)

	postBytesReader := bytes.NewReader(postDataBytes)
	resp, err := http.Post(url, "application/json", postBytesReader)
	fail.When(err != nil, err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fail.When(err != nil, err)
	list := make(map[string]RPCServer)
	err = json.Unmarshal(body, &list)
	fail.When(err != nil, err)
	return list
}
