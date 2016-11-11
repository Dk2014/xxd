// 本文件由 xdlog-parser 根据日志xml规范生成
//请勿手动改写！
package xdlog

import (
	"time"
	"strconv"
	"game_server/mdb"
	glog "core/log"
	"encoding/json"
	gotime "core/time"
	"strings"
)

import (
	"testing"
	"os"
	"time"
)

var logs []Log

func Test(t *testing.T) {
	
	logs=append(logs,NewLoginFlow(123, "test string", 12345, "test string", "test string", 123, 123, 123, 123))
	
	logs=append(logs,NewCreateFlow(123, "test string", 12345, "test string", "test string", 123))
	
	logs=append(logs,NewChargeFlow(123, "test string", 12345, "test string", "test string", "test string", "test string", "test string", 123, 123.456, 12345, "test string"))
	
	logs=append(logs,NewEventFlow(123, "test string", 12345, "test string", 123, 123, 123))
	
	logs=append(logs,NewIncomeFlow(123, "test string", 12345, "test string", 123, 123, 123, 123, 123, "test string"))
	
	logs=append(logs,NewConsumeFlow(123, "test string", 12345, "test string", 123, 123, 123, 123, 123, 123))
	
	logs=append(logs,NewItemFlow(123, "test string", 12345, "test string", 123, 123, 123, 123, 123, 123, "test string"))
	
	logs=append(logs,NewPropsFlow(123, "test string", 12345, "test string", 123, 123, 123, 123, 123))
	
	logs=append(logs,NewLevelFlow(123, "test string", 12345, "test string", 123, 123, 123, 123))
	
	logs=append(logs,NewVipFlow(123, "test string", 12345, "test string", 123, 123, 123, 123))
	
	logs=append(logs,NewOnlineFlow(123, 123))
	
	logs=append(logs,NewSnapshotFlow(123, "test string", 12345, "test string", 12345, 12345, 12345, "test string", 123, 123, 123))
	
	logs=append(logs,NewMissionFlow(123, "test string", 12345, "test string", 123, 123, 123, 123, 123, 123, 123, 123))
	
	logs=append(logs,NewGiftcodeFlow(123, "test string", 12345, "test string", "test string", 123, 123, 123))
	
	logs=append(logs,NewFightnumFlow(123, "test string", 12345, "test string", 12345, 123, 123))
	
	f, err := os.Create("../../testdata/log.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for _, log := range logs {
		debugWriteToFile(log, f)
	}
}