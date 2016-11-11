package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"syscall"
	"time"
)

var (
	out           *bufio.Writer
	_LOG_PATH_    string
	_LOG_PATH_FD_ string
	mutex         *sync.Mutex
	logFile       *os.File
	logFileFd     *os.File
	logger        *log.Logger
)

func main() {
	var (
		LogPath      = flag.String("l", "./log", "log path")
		XdlogLogPath = flag.String("xl", "./xdlog", "xdlog path")
		UdpPort      = flag.Int("p", 6667, "upd port")
	)
	flag.Parse()

	socket, err := net.ListenUDP("udp4", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: *UdpPort,
	})

	if err != nil {
		fmt.Println("Listen失败", err)
		return
	}

	Setup(*XdlogLogPath, *LogPath, true)

	defer socket.Close()

	for {
		data := make([]byte, 2048)
		length, _, err := socket.ReadFromUDP(data)
		if err != nil {
			fmt.Println("读取数据失败", err)
			continue
		}
		switch err {
		case nil:
			handleMsg(length, err, data, logFile)
		default:
			goto DISCONNECT
		}
	}
DISCONNECT:
	err = socket.Close()
	checkError(err, "Close:")

}

func handleMsg(length int, err error, msg []byte, f *os.File) {
	if length > 0 {
		for i := 0; ; i++ {
			if msg[i] == 0 {
				break
			}
		}
		f.Write(msg[:length])
	}

}

func checkError(error error, info string) {
	if error != nil {
		panic("ERROR: " + info + " " + error.Error()) // terminate program
	}
}

func Setup(logDir string, logDirFd string, isRedirect bool) {
	mutex = new(sync.Mutex)
	_LOG_PATH_ = logDir
	_LOG_PATH_FD_ = logDirFd

	ensureLogDir()

	// 每天切换文件
	go func() {
		defer func() { recover() }()
		for {
			now := time.Now()
			switchLogFile(now, isRedirect)
			// 计算此刻到第二天零点的时间
			t := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, now.Nanosecond(), now.Location())
			duration := t.Sub(now)
			time.Sleep(duration)
		}
	}()
}

func ensureLogDir() {
	dir, _ := os.Stat(_LOG_PATH_)
	if dir == nil {
		os.Mkdir(_LOG_PATH_, 0777)
	}
	dir, _ = os.Stat(_LOG_PATH_FD_)
	if dir == nil {
		os.Mkdir(_LOG_PATH_FD_, 0777)
	}
}

func switchLogFile(now time.Time, isRedirect bool) {
	// file.Fd()==18446744073709551615为true  文件已经close 或者 没有打开
	// 目前正在查找是否有更加直观的写法
	mutex.Lock()
	defer mutex.Unlock()

	if logFile != nil {
		logFile.Close()
	}

	if logFileFd != nil {
		logFileFd.Close()
	}
	var err error
	logName := _LOG_PATH_ + "/" + now.Format("2006-01-02") + ".log"
	logNameFd := _LOG_PATH_FD_ + "/" + now.Format("2006-01-02") + ".log"
	logFile, err = os.OpenFile(logName, os.O_WRONLY|os.O_CREATE|os.O_APPEND|os.O_SYNC, 0777)
	if err != nil {
		panic(err)
	}
	logFileFd, err = os.OpenFile(logNameFd, os.O_WRONLY|os.O_CREATE|os.O_APPEND|os.O_SYNC, 0777)
	if err != nil {
		panic(err)
	}
	if isRedirect {
		syscall.Dup2(int(logFileFd.Fd()), 1)
		syscall.Dup2(int(logFileFd.Fd()), 2)
	}

	out = bufio.NewWriterSize(logFileFd, 1024000)
	logger = log.New(out, "", log.Ldate|log.Ltime)
}
