package log

import (
	"bufio"
	golog "log"
	"os"
	"reflect"
	"strconv"
	"sync"
	"time"
)

var (
	out     *bufio.Writer
	logger  *golog.Logger
	mutex   *sync.Mutex
	logFile *os.File
)

func WriteLog(path, struct_name string, Cmdid uint32, val map[int]interface{}) {
	buff := make([]byte, 0, 1484)
	buff = append(buff, []byte(struct_name)...)
	buff = append(buff, '|')
	buff = strconv.AppendUint(buff, uint64(Cmdid), 10)
	buff = append(buff, '|')
	nowTime := time.Now().Format("2006-01-02 15:04:05")
	if len(nowTime) > 256 {
		buff = append(buff, []byte(nowTime[:256])...)
	} else {
		buff = append(buff, []byte(nowTime)...)
	}
	for i := 0; i < len(val); i++ {
		switch val[i].(type) {
		case int64:
			buff = append(buff, '|')
			buff = strconv.AppendInt(buff, int64(val[i].(int64)), 10)
		case int32:
			buff = append(buff, '|')
			buff = strconv.AppendInt(buff, int64(val[i].(int32)), 10)
		case int16:
			buff = append(buff, '|')
			buff = strconv.AppendInt(buff, int64(val[i].(int16)), 10)
		case int8:
			buff = append(buff, '|')
			buff = strconv.AppendInt(buff, int64(val[i].(int8)), 10)
		case int:
			buff = append(buff, '|')
			buff = strconv.AppendInt(buff, int64(val[i].(int)), 10)
		case uint64:
			buff = append(buff, '|')
			buff = strconv.AppendUint(buff, uint64(val[i].(uint64)), 10)
		case uint32:
			buff = append(buff, '|')
			buff = strconv.AppendUint(buff, uint64(val[i].(uint32)), 10)
		case uint16:
			buff = append(buff, '|')
			buff = strconv.AppendUint(buff, uint64(val[i].(uint16)), 10)
		case uint8:
			buff = append(buff, '|')
			buff = strconv.AppendUint(buff, uint64(val[i].(uint8)), 10)
		case uint:
			buff = append(buff, '|')
			buff = strconv.AppendUint(buff, uint64(val[i].(uint)), 10)
		case string:
			buff = append(buff, '|')
			if len(val[i].(string)) > 256 {
				buff = append(buff, []byte(val[i].(string)[:256])...)
			} else {
				buff = append(buff, []byte(val[i].(string))...)
			}
		default:
			continue
		}
	}
	buff = append(buff, '\n')
	switchAndWrite(path, buff[:len(buff)])
}

func switchAndWrite(path string, content []byte) {
	mutex = new(sync.Mutex)
	mutex.Lock()
	defer mutex.Unlock()

	if logFile != nil {
		logFile.Close()
	}

	dir, _ := os.Stat(path)
	if dir == nil {
		os.Mkdir(path, 0777)
	}

	var err error
	logName := path + "/" + time.Now().Format("2006-01-02") + ".log"
	logFile, err = os.OpenFile(logName, os.O_WRONLY|os.O_CREATE|os.O_APPEND|os.O_SYNC, 0777)
	defer logFile.Close()
	if err != nil {
		panic(err)
	}

	logFile.Write(content)
}

func Getdatamap(obj interface{}, data map[int]interface{}) map[int]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	dataLen := len(data)
	for i := 0; i < t.NumField(); i++ {
		data[i+dataLen] = v.Field(i).Interface()
	}
	return data
}
