package mdb

import (
	"core/log"
	"fmt"
	"os"
	"time"
)

// MySql同步进程
type tlogCommitter struct {
	commitChan     chan []TransLog // 事务被发送到这个频道
	commitCount    int64           // 发送到队列的次数
	completeCount  int64           // 同步完成次数
	waitStopChan   chan int        // 用于等待同步循环结束
	switchFileChan chan bool       // 用于切换同步文件
}

func newTlogCommitter(logDir string) *tlogCommitter {
	this := &tlogCommitter{
		commitChan:     make(chan []TransLog, TLOG_COMMIT_CHAN_BUF_MAX),
		waitStopChan:   make(chan int),
		switchFileChan: make(chan bool),
	}

	this.startFileTimer()
	go this.commitLoop(logDir)
	return this
}

// 同步循环
func (this *tlogCommitter) commitLoop(logDir string) {
	var (
		ok       bool
		transLog []TransLog
		tlogFile *SyncFile
	)

	tlogFile = NewTLogFile(logDir)
L:
	for {
		select {
		case transLog, ok = <-this.commitChan:
			if !ok {
				break L
			}

			defer func() {
				if err := recover(); err != nil {
					panic(TransError{err})
				}
			}()

			for _, tran := range transLog {
				if tran.GetTransType() == TRANS_TYPE_TLOG {
					tran.CommitToFile(tlogFile)
				}
			}

			this.completeCount += 1

			// 文件切换
		case _, ok := <-this.switchFileChan:
			if ok {
				tlogFile.Close()
				tlogFile = NewTLogFile(logDir)
			}
		}
	}

	tlogFile.Close()
	close(this.waitStopChan)
}

// 停止同步进程
// 等待，直到同步队列中所有事务同步完成
func (this *tlogCommitter) Stop() {
	close(this.commitChan)
	<-this.waitStopChan
}

func (this *tlogCommitter) Commit(trans []TransLog) {
	this.commitCount += 1
	this.commitChan <- trans
}

// 获取同步队列当前长度
func (this *tlogCommitter) QueueLength() int64 {
	return this.commitCount - this.completeCount
}

// 仅供tlog使用
func (file *SyncFile) WriteTLog(log []byte) {
	file.f.Write(log)
}

func (this *tlogCommitter) startFileTimer() {
	go func() {
		for {
			now := time.Now()
			nextHour := time.Date(now.Year(), now.Month(), now.Day(), now.Hour()+1, 0, 0, 0, now.Location())
			duration := nextHour.Sub(now)
			time.Sleep(duration)

			this.switchFileChan <- true
		}
	}()
}

func (this *tlogCommitter) ProfilePrintToLog() {
	log.Infof("[Profile] tlogCommitter::commitChan len/cap: %d/%d, completeCount/commitCount: %d/%d",
		len(this.commitChan), cap(this.commitChan),
		this.completeCount, this.commitCount)
}

func (this *tlogCommitter) ProfileWriteToFile(fp *os.File) {
	fp.WriteString(fmt.Sprintf("[Profile] tlogCommitter::commitChan len/cap: %d/%d, completeCount/commitCount: %d/%d \n",
		len(this.commitChan), cap(this.commitChan),
		this.completeCount, this.commitCount))
}

func NewTLogFile(fileDir string) *SyncFile {
	now := time.Now()
	logDir := fileDir + now.Format("2006-01-02")

	if _, errDir := os.Stat(logDir); errDir != nil {
		os.MkdirAll(logDir, 0777)
	}

	var err error
	var file *os.File
	var fileName string

	index := 0
	for {
		if index == 0 {
			fileName = fmt.Sprintf("%s/%d.log", logDir, now.Hour())
		} else {
			fileName = fmt.Sprintf("%s/%d-%d.log", logDir, now.Hour(), index)
		}

		if _, errExist := os.Stat(fileName); errExist != nil {
			file, err = os.OpenFile(fileName, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0777)
			if err != nil {
				panic(err)
			}

			break
		}

		index++
	}

	syncFile := &SyncFile{
		f: file,
	}

	return syncFile
}
