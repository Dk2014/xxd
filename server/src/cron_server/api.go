package main

import (
	. "cron_server/database"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

func PacketPlayer(player *Player) []byte {
	js, err := json.Marshal(player)
	if err != nil {
		panic(err.Error())
	}
	js = append(js, '\n')
	return js
}

func PacketOnline(online *Online) []byte {
	js, err := json.Marshal(online)
	if err != nil {
		panic(err.Error())
	}
	js = append(js, '\n')
	return js
}

//
func SnapWriteToLog(data []byte, filepath string, sid int) {
	var f *os.File
	var err error
	nowDate := time.Now().Format("2006-01")
	dirpath := filepath + "/snapshot/" + fmt.Sprintf("srv-%d", sid)
	filename := nowDate + ".log"

	if !checkFileIsExist(dirpath) {
		os.MkdirAll(dirpath, 0777)
	}
	f, err = os.OpenFile(dirpath+"/"+filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.WriteString(f, string(data))
	if err != nil {
		panic(err)
	}
}

func OnlineWriteToLog(data []byte, filepath string, sid int32) {
	var f *os.File
	var err error
	nowDate := time.Now().Format("2006-01-02")
	dirpath := filepath + "/online/" + fmt.Sprintf("srv-%d", sid)
	filename := nowDate + ".log"

	if !checkFileIsExist(dirpath) {
		os.MkdirAll(dirpath, 0777)
	}
	f, err = os.OpenFile(dirpath+"/"+filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.WriteString(f, string(data))
	if err != nil {
		panic(err)
	}
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
