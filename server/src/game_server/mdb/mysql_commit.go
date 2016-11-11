package mdb

import (
	"core/log"
	"core/mysql"
	"fmt"
	"os"
	"time"
)

// MySql同步进程
type mysqlCommitter struct {
	mysqlInfo     map[string]interface{} // MySql连接参数
	commitConn    *mysql.Connection      // 同步用的MySql连接
	commitChan    chan []TransLog        // 事务被发送到这个频道
	commitCount   int64                  // 发送到队列的次数
	completeCount int64                  // 同步完成次数
	waitStopChan  chan int               // 用于等待同步循环结束
}

// 连接到MySql，并创建一个MySql同步进程
func newMySqlCommitter(mysqlInfo map[string]interface{}) *mysqlCommitter {
	this := &mysqlCommitter{
		mysqlInfo:    mysqlInfo,
		commitChan:   make(chan []TransLog, MYSQL_COMMIT_CHAN_BUF_MAX),
		waitStopChan: make(chan int),
	}

	conn, err := this.connMySql()
	if err != nil {
		panic(err)
	}

	prepareSQL(conn)
	this.commitConn = conn

	go this.commitLoop()
	return this
}

func (this *mysqlCommitter) connMySql() (*mysql.Connection, error) {
	// 尝试连接数据库
	conn, err1 := mysql.Connect(this.mysqlInfo)
	if err1 != nil {
		return nil, err1
	}
	// 禁止事务自动提交
	// 才能在同步时手工控制数据库事务提交，避免部分提交的情况发生
	err2 := conn.ForbidAutoCommit()
	if err2 != nil {
		conn.Close()
		return nil, err2
	}
	return conn, nil
}

// 同步循环
func (this *mysqlCommitter) commitLoop() {
	var (
		ok       bool
		err      error
		transLog []TransLog
	)
L:
	for {
		select {
		case transLog, ok = <-this.commitChan:
			if !ok {
				break L
			}

		RESTRY:
			for _, item := range transLog {
				if item.GetTransType() != TRANS_TYPE_MYSQL {
					continue
				}

				err = item.CommitToMySql()
				if err != nil {
					break
				}
			}

			if err == nil {
				err = this.commitConn.Commit()
			} else {
				this.commitConn.RollBack()

				log.Errorf("mysql commit failed: %v", err)
				log.Flush()

				// 如果是连接断开，尝试重连MySql
				if stmtError, ok := err.(*mysql.StmtError); ok && stmtError.Num == 2013 {
					reconnected := false

					for i := 0; i < 100; i++ {
						newConn, reconnError := this.connMySql()
						if reconnError == nil {
							reconnected = true
							prepareSQLClean()
							prepareSQL(newConn)
							this.commitConn = newConn
							log.Info("mysql reconnect succeed")
							log.Flush()
							break
						}
						log.Errorf("mysql reconnect failed: %v", reconnError)
						log.Flush()
						time.Sleep(time.Second * 6)
					}

					if reconnected {
						err = nil
						goto RESTRY
					}
				}

				panic(err)
			}

			// 释放内存
			for _, item := range transLog {
				item.Free()
			}

			this.completeCount += 1
		}
	}

	close(this.waitStopChan)
}

// 停止mysql同步进程
// 等待，直到同步队列中所有事务同步完成
func (this *mysqlCommitter) Stop() {
	close(this.commitChan)
	<-this.waitStopChan
	this.commitConn.Close()
}

// 将事务同步到mysql
func (this *mysqlCommitter) Commit(trans []TransLog) {
	this.commitCount += 1
	this.commitChan <- trans
}

// 获取同步队列当前长度
func (this *mysqlCommitter) QueueLength() int64 {
	return this.commitCount - this.completeCount
}

func (this *mysqlCommitter) ProfilePrintToLog() {
	log.Infof("[Profile] mysqlCommitter::commitChan len/cap: %d/%d, completeCount/commitCount: %d/%d",
		len(this.commitChan), cap(this.commitChan),
		this.completeCount, this.commitCount)
}

func (this *mysqlCommitter) ProfileWriteToFile(fp *os.File) {
	fp.WriteString(fmt.Sprintf("[Profile] mysqlCommitter::commitChan len/cap: %d/%d, completeCount/commitCount: %d/%d \n",
		len(this.commitChan), cap(this.commitChan),
		this.completeCount, this.commitCount))
}
