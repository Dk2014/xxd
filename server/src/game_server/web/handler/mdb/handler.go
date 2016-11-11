package mdb_handler

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

import (
	"core/debug"
	"game_server/mdb"
)

func init() {
	http.HandleFunc("/debug/mdb/query", queryHandler)
	http.HandleFunc("/debug/mdb/update", updateHandler)
}

type QueryRequest struct {
	Table string `json:"table"`
	Pid   int64  `json:"pid"`
}

type QueryResponse struct {
	Status  int           `json:"status"`
	Message string        `json:"msg"`
	Result  []interface{} `json:"result"`
}

func queryHandler(w http.ResponseWriter, req *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			io.WriteString(w, fmt.Sprintf(`{status:-1, msg:"%v"}`, err))
			fmt.Println(string(debug.Stack(1, "    ")))
		}
	}()

	w.Header().Add("Content-Type", "application/json")

	rawReq, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	reqStruct := &QueryRequest{}
	if err := json.Unmarshal(rawReq, reqStruct); err != nil {
		panic(err)
	}
	reqStruct.Table = formatName(reqStruct.Table)

	fmt.Println(reqStruct.Pid, reqStruct.Table)

	respStruct := &QueryResponse{}
	mdb.Transaction(mdb.TRANS_TAG_WebDebugger, func() {
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(reqStruct.Pid, func(agentDB *mdb.Database) {
				allOp := agentDB.All
				allOpType := reflect.TypeOf(allOp)
				if _, ok := allOpType.MethodByName(reqStruct.Table); ok {
					results := reflect.ValueOf(allOp).MethodByName(reqStruct.Table).Call(nil)
					if len(results) > 0 {
						respStruct.Result = results[0].Interface().([]interface{})
					}
				} else if _, ok := reflect.TypeOf(agentDB.Lookup).MethodByName(reqStruct.Table); ok {
					args := make([]reflect.Value, 0)
					args = append(args, reflect.ValueOf(reqStruct.Pid))
					results := reflect.ValueOf(agentDB.Lookup).MethodByName(reqStruct.Table).Call(args)
					if len(results) > 0 {
						respStruct.Result = append(respStruct.Result, results[0].Elem().Interface())
					}
				} else {
					panic(fmt.Sprint("表 [%s] 不存在", reqStruct.Table))
				}
			})
		})
	})

	rspBytes, err := json.Marshal(respStruct)
	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(rspBytes)
}

type UpdateRequest struct {
	Table string                 `json:"table"`
	Key   int64                  `json:"key"`
	Pid   int64                  `json:"pid"`
	Set   map[string]interface{} `json:"set"`
}

type UpdateResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"msg"`
	Result  interface{} `json:"result"`
}

func updateHandler(w http.ResponseWriter, req *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			io.WriteString(w, fmt.Sprintf(`{status:-1, msg:"%v"}`, err))
			fmt.Println(string(debug.Stack(1, "    ")))
		}
	}()

	w.Header().Add("Content-Type", "application/json")

	rawReq, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	reqStruct := &UpdateRequest{}
	if err := json.Unmarshal(rawReq, reqStruct); err != nil {
		panic(err)
	}
	reqStruct.Table = formatName(reqStruct.Table)
	fmt.Println(reqStruct.Key, reqStruct.Table)

	respStruct := &UpdateResponse{}
	mdb.Transaction(mdb.TRANS_TAG_WebDebugger, func() {
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(reqStruct.Pid, func(agentDB *mdb.Database) {
				if _, ok := reflect.TypeOf(agentDB.Lookup).MethodByName(reqStruct.Table); ok {
					args := make([]reflect.Value, 0)
					args = append(args, reflect.ValueOf(reqStruct.Key))
					recordPtr := reflect.ValueOf(agentDB.Lookup).MethodByName(reqStruct.Table).Call(args)[0].Interface()
					if reflect.ValueOf(recordPtr).IsNil() {
						panic(fmt.Sprintf("record not found table[%s] key [%d], pid [%d]", reqStruct.Table, reqStruct.Key, reqStruct.Pid))
					}
					record := reflect.ValueOf(recordPtr).Elem()

					for field, value := range reqStruct.Set {
						field = formatName(field)
						fmt.Println(field, "=>", value)
						if structField := record.FieldByName(field); structField.IsValid() && structField.CanSet() {
							fmt.Println("结构体成员", structField.CanAddr(), structField.IsValid(), structField.CanSet())
							fmt.Println(field, reflect.TypeOf(structField.Interface()))
							targetValue := reflect.ValueOf(value).Convert(reflect.TypeOf(structField.Interface()))
							structField.Set(targetValue)
						} else {
							respStruct.Message += fmt.Sprintf("unknow field [%s]|", field)
						}
					}
					//fmt.Printf("%#v\n", record.Interface())
					args2 := make([]reflect.Value, 0)
					args2 = append(args2, reflect.ValueOf(recordPtr))
					reflect.ValueOf(agentDB.Update).MethodByName(reqStruct.Table).Call(args2)
					respStruct.Result = record.Interface()

				} else {
					panic(fmt.Sprint("表 [%s] 不存在", reqStruct.Table))
				}
			})
		})
	})

	rspBytes, err := json.Marshal(respStruct)
	if err != nil {
		panic(err)
	}

	w.Write(rspBytes)
}

func formatName(name string) (formatedName string) {
	parts := strings.Split(name, "_")
	for _, part := range parts {
		formatedName += strings.Title(part)
	}
	return formatedName
}
