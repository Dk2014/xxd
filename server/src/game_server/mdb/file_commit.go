package mdb

import "os"
import "math"
import "encoding/binary"
import "time"
import util "core/time"
import "fmt"
import "core/log"

// 数据库同步文件结构说明:
//
// *. 每个同步文件的头部，通过init()函数写入一份文件头
// *. 文件头包含的信息是每个表的表名字段名，字段顺序，字段类型等信息
// *. 文件头信息用四个字节起始，这4个字节类似于协议的包长度信息，是一个uint32值，表明头信息字节长度
// *. 头信息中，每个表有两条[]byte类型的数据，自带长度信息
// *. 第一条数据为表名+字段列表，内容按逗号分割，第一个值是表名，之后为字段列表
// *. 第二条数据为表字段类型，每个字段类型占用一个字节
// *. 文件头之后的内容为事务内容
// *. 每个事务用4个字节开头，算法给文件头长度信息一致
// *. 通过内容长度可以验证内容是否完整，如果出现损失，则损失最后一个事务的内容
// *. 事务中每一笔操作，都一个字节的开头用来标明具体是什么类型操作，1-插入，2-删除，3-修改
// *. 操作类型之后是4个字节的表序号，对应同步文件头中给出的表顺序，从0开始
// *. 表名之后的内容根据操作类型不同，分为以下几种情况：插入操作写入新数据内容，删除操作写入旧数据内容，修改操作先写入新数据内容，再写入旧数据内容
// *. 新旧数据的写入按字段顺序写入操作，对应同步文件头中表信息给出的字段顺序
// *. 由于字段顺序、字段名、字段类型是已知的，所以事务内容中不再包含这部分信息，读取时按字段顺序和类型顺序读取即可

/*
	文件头：							|uint32|
	表名+字段名：第一条[]byte:			|tableName,field1,field2,field3,...|
	字段类型：第二条[]byte:			|uint8|uint8|uint8|...|		(1-char, 2-varchar, 3-int, 4-bigint)
	协议数据： 第三条[]byte: 			|模块ID, 模块名, actionID, action名, action...|

	协议号：							|uint32|  低16为是网络协议uint8|uint8,分别是模块id和接口id；高16位自定义常量标识
	事务发生时间：					|int64|

	事务头：							|uint32|
	操作标识：						|uint8| 					(1-插入，2-删除，3-修改)
	表序号：							|uint32|					(序号对应init()中的表信息记录序列，记录事务处理所属哪个表)
	内容：							|bytes|
*/

type syncTrans struct {
	transId uint32
	trans   []TransLog
}

type fileCommitter struct {
	commitChan     chan syncTrans  // 事务被发送到这个频道
	mysqlCommitter *mysqlCommitter // 对应的数据库同步器
	tlogCommitter  *tlogCommitter  // 腾讯tlog同步
	xdlogCommitter *xdlogCommitter // 心动xdlog同步
	commitCount    int64           // 发送到队列的次数
	completeCount  int64           // 同步完成次数
	waitStopChan   chan int        // 用于等待同步循环结束
	switchFileChan chan bool       // 用于切换同步文件
}

func newFileCommitter(mysqlCommitter *mysqlCommitter, syncDir string, tlogDir string, xdlogDir string, enablexdlog, enabletlog bool) *fileCommitter {
	this := &fileCommitter{
		commitChan:     make(chan syncTrans, FILE_COMMIT_CHAN_BUF_MAX),
		waitStopChan:   make(chan int),
		switchFileChan: make(chan bool),
		mysqlCommitter: mysqlCommitter,
	}
	if enablexdlog {
		this.xdlogCommitter = newXdlogCommitter(xdlogDir)
	}
	if enabletlog {
		this.tlogCommitter = newTlogCommitter(tlogDir)
	}

	this.startFileTimer()
	go this.commitLoop(syncDir, enablexdlog, enabletlog)

	return this
}

func (this *fileCommitter) commitLoop(syncDir string, enablexdlog, enabletlog bool) {

	var syncFile *SyncFile

	syncFile = NewSyncFile(syncDir)
L:
	for {
		select {
		// 提交事务
		case syncTrans, ok := <-this.commitChan:
			if !ok {
				break L
			}

			haveMysqlTrans := false
			for _, tran := range syncTrans.trans {
				if tran.GetTransType() == TRANS_TYPE_MYSQL {
					haveMysqlTrans = true
					break
				}
			}

			// 提交mysql事务
			if haveMysqlTrans {
				syncFile.BeginTrans()

				syncFile.WriteUint32(syncTrans.transId)
				syncFile.WriteInt64(util.GetNowTime())

				for _, tran := range syncTrans.trans {
					if tran.GetTransType() == TRANS_TYPE_MYSQL {
						tran.CommitToFile(syncFile)
					}
				}
				syncFile.EndTrans()

				this.mysqlCommitter.Commit(syncTrans.trans)
			}

			// 提交心动log事务
			if enablexdlog {
				this.xdlogCommitter.Commit(syncTrans.trans)
			}
			if enabletlog {
				// 提交tlog事务
				this.tlogCommitter.Commit(syncTrans.trans)
			}

			this.completeCount += 1

		// 文件切换
		case _, ok := <-this.switchFileChan:
			if ok {
				syncFile.Close()
				syncFile = NewSyncFile(syncDir)
			}
		}
	}

	syncFile.Close()
	close(this.waitStopChan)
}

func (this *fileCommitter) startFileTimer() {
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

func (this *fileCommitter) Commit(transId uint32, trans []TransLog) {
	this.commitCount += 1
	this.commitChan <- syncTrans{transId, trans}
}

func (this *fileCommitter) Stop(enablexdlog, enabletlog bool) {
	close(this.commitChan)
	<-this.waitStopChan
	if enablexdlog {
		this.xdlogCommitter.Stop()
	}
	if enabletlog {
		this.tlogCommitter.Stop()
	}
	this.mysqlCommitter.Stop()
}

func (this *fileCommitter) ProfilePrintToLog(enablexdlog, enabletlog bool) {
	log.Infof("[Profile] fileCommitter::commitChan len/cap: %d/%d, completeCount/commitCount: %d/%d",
		len(this.commitChan), cap(this.commitChan),
		this.completeCount, this.commitCount)

	if enablexdlog {
		this.xdlogCommitter.ProfilePrintToLog()
	}
	if enabletlog {
		this.tlogCommitter.ProfilePrintToLog()
	}
	this.mysqlCommitter.ProfilePrintToLog()
}

func (this *fileCommitter) ProfileWriteToFile(fp *os.File, enablexdlog, enabletlog bool) {
	fp.WriteString(fmt.Sprintf("[Profile] fileCommitter::commitChan len/cap: %d/%d, completeCount/commitCount: %d/%d \n",
		len(this.commitChan), cap(this.commitChan),
		this.completeCount, this.commitCount))

	if enabletlog {
		this.tlogCommitter.ProfileWriteToFile(fp)
	}
	this.mysqlCommitter.ProfileWriteToFile(fp)
	if enablexdlog {
		this.xdlogCommitter.ProfileWriteToFile(fp)
	}
}

type SyncFile struct {
	f         *os.File
	transHead []byte
	buffer    []byte
}

/*
   	按时间获取文件名，并防止冲突
	比如按当前时间获得文件名：2014-07-10-15-00.sync
	发现2014-07-10-15-00.sync已存在时，则命名为2014-07-10-15-01.sync
	递归直到发现可用的文件名，用来防止服务器重启后同步文件被覆盖
*/
func NewSyncFile(syncFileDir string) *SyncFile {
	now := time.Now()
	logDir := syncFileDir + now.Format("2006-01-02")

	if _, errDir := os.Stat(logDir); errDir != nil {
		os.MkdirAll(logDir, 0777)
	}

	var err error
	var file *os.File
	var fileName string

	index := 0
	for {
		if index == 0 {
			fileName = fmt.Sprintf("%s/%d.bsql", logDir, now.Hour())
		} else {
			fileName = fmt.Sprintf("%s/%d-%d.bsql", logDir, now.Hour(), index)
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
		f:         file,
		transHead: make([]byte, 4),
		buffer:    make([]byte, 0, 1024),
	}

	err = syncFile.initDB()
	if err != nil {
		panic(err)
	}

	err = syncFile.initProtocol()
	if err != nil {
		panic(err)
	}

	err = syncFile.initRPCTransInfo()
	if err != nil {
		panic(err)
	}

	return syncFile
}

func (file *SyncFile) BeginTrans() {
	file.buffer = file.buffer[0:0]
}

func (file *SyncFile) EndTrans() error {
	binary.BigEndian.PutUint32(file.transHead, uint32(len(file.buffer)))

	_, err := file.f.Write(file.transHead)
	if err != nil {
		return err
	}

	_, err = file.f.Write(file.buffer)

	return err
}

func (file *SyncFile) Close() error {
	return file.f.Close()
}

func (file *SyncFile) growsBuffer(n int) []byte {
	oldLen := len(file.buffer)
	newLen := oldLen + n
	if newLen <= cap(file.buffer) {
		file.buffer = file.buffer[0:newLen]
	} else {
		b := make([]byte, newLen, newLen+512)
		copy(b, file.buffer)
		file.buffer = b
	}
	return file.buffer[oldLen:newLen]
}

func (file *SyncFile) WriteUint8(v uint8) {
	file.growsBuffer(1)[0] = v
}

func (file *SyncFile) WriteInt8(v int8) {
	file.WriteUint8(uint8(v))
}

func (file *SyncFile) WriteUint16(v uint16) {
	binary.BigEndian.PutUint16(file.growsBuffer(2), v)
}

func (file *SyncFile) WriteInt16(v int16) {
	file.WriteUint16(uint16(v))
}

func (file *SyncFile) WriteUint32(v uint32) {
	binary.BigEndian.PutUint32(file.growsBuffer(4), v)
}

func (file *SyncFile) WriteInt32(v int32) {
	file.WriteUint32(uint32(v))
}

func (file *SyncFile) WriteUint64(v uint64) {
	binary.BigEndian.PutUint64(file.growsBuffer(8), v)
}

func (file *SyncFile) WriteInt64(v int64) {
	file.WriteUint64(uint64(v))
}

func (file *SyncFile) WriteFloat32(v float32) {
	file.WriteUint32(math.Float32bits(v))
}

func (file *SyncFile) WriteFloat64(v float64) {
	file.WriteUint64(math.Float64bits(v))
}

func (file *SyncFile) WriteBytes(v []byte) {
	n := len(v)
	file.WriteUint32(uint32(n))
	copy(file.growsBuffer(n), v)
}

func (file *SyncFile) WriteString(v string) {
	file.WriteBytes([]byte(v))
}
