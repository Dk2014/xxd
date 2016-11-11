package battle

// import (
// 	"encoding/json"
// 	"fmt"
// 	debug "github.com/realint/dbgutil"
// 	"github.com/realint/netutil"
// 	"io"
// 	"lms/lib/config"
// 	"lms/lib/log"
// 	"lms/lib/sqlite"
// 	"lms_client/lib/util"
// 	"reflect"
// 	"sync/atomic"
// 	"time"
// 	"unsafe"
// )

// 战报记录
type iRecord interface {
	// Pack(data *netutil.Output)
	Pack(data interface{})
	Size() int
}

/*
// 增加返回
func (b *BattleState) Record(record iRecord) {
	if b.needRecord {
		b.records = append(b.records, record)
	}
}

// 封包
func (b *BattleState) RecordPack() []byte {
	var (
		recordBuf = make([]byte, 0, 32*1024)
		freeLen   = len(recordBuf)
		totalLen  = 0
	)

	for _, r := range b.records {
		packetLen := r.Size()

		realLen := packetLen + 4

		// 剩余空间不足
		if freeLen < realLen {
			newBuf := make([]byte, len(recordBuf)+32*1024)

			copy(newBuf, recordBuf)

			freeLen = len(newBuf) - totalLen + freeLen

			recordBuf = newBuf
		}

		// 总长度 ＝ 消息长度 + 消息头长度（4）
		endIndex := totalLen + realLen
		buff := recordBuf[totalLen:endIndex]

		// 伪造一个output
		output := &netutil.Output{Data: buff}

		// 写入消息头
		output.WriteUint32(uint32(packetLen))

		r.Pack(output)

		totalLen += realLen
		freeLen -= realLen
	}

	return recordBuf[:totalLen]
}

// 储存
func (b *BattleState) RecordSave(details interface{}) (recordId, version int64) {
	db := defaultRecordDb
	recordId = atomic.AddInt64(&(db.initRecordId), 1)
	version = db.srvVersion
	detailsBytes, err := json.Marshal(details)
	if err != nil {
		panic(err)
	}

	db.recordCh <- &recordData{b, recordId, detailsBytes}
	return
}

type BattleRecordDb struct {
	conn         *sqlite.Conn
	tableName    string
	initRecordId int64
	insertStmt   *sqlite.Stmt
	selectStmt   *sqlite.Stmt
	deleteStmt   *sqlite.Stmt
	err          error

	srvVersion int64 // 服务版本

	recordCh chan *recordData       // 战报纪录管道
	selectCh chan *getRecordRequest // 取战报管道
	stopCh   chan int               // 关闭战报管道
}

type recordData struct {
	battle  *BattleState
	id      int64
	details []byte
}

type getRecordRequest struct {
	id       int64
	callback func(batType int, details, v []byte) // 回调函数
}

var defaultRecordDb *BattleRecordDb

func StartDefaultRecordDb(dbName string) {
	defaultRecordDb = NewBattleRecordDb(dbName)
}

func GetDefaultRecordDb() *BattleRecordDb {
	return defaultRecordDb
}

func StopDefaultRecordDb() {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf(`Battle Record Release
Error    = %v
Stack    =
%s`,
				err,
				debug.Stack(1, "    "),
			)
		}
	}()
	// 等待数据写完
	for {
		if len(defaultRecordDb.recordCh) == 0 {
			defaultRecordDb.stopCh <- 1
			break
		} else {
			// 每次执行时间
			time.Sleep(5e7)
		}
	}

	defaultRecordDb.Release()
}

func NewBattleRecordDb(dbName string) *BattleRecordDb {
	db := &BattleRecordDb{
		recordCh:   make(chan *recordData, 1000),       // 战报纪录管道
		selectCh:   make(chan *getRecordRequest, 1000), // 取战报管道
		stopCh:     make(chan int),
		srvVersion: config.Server.Version,
	}

	db.conn, db.err = sqlite.Open(dbName)
	db.e()

	db.init()

	go db.loop()

	return db
}

func (b *BattleRecordDb) loop() {
L:
	for {
		select {
		case r := <-b.recordCh:
			b.Insert(int(r.battle.BatType), r.details, r.battle.RecordPack(), r.id)
		case req := <-b.selectCh:
			b.getRecord(req)
		case <-b.stopCh:
			break L
		}
	}
}

func (b *BattleRecordDb) init() {
	isExist := b.tableIsExist()
	if !isExist {
		sql := "CREATE TABLE record (" +
			" `id` BIGINT PRIMARY KEY," +
			" `v` BLOB COMMENT '战报数据'," +
			" `version` BIGINT COMMENT '版本号'," +
			" `time` BIGINT COMMENT '添加时间'," +
			" `type` TINYINT COMMENT '战报类型'," +
			" `details` BLOB COMMENT '战报信息'" +
			");"
		b.err = b.conn.Exec(sql)
		b.e()
	} else {
		b.initRecordId = b.getMaxId()
	}

	b.insertStmt, b.err = b.conn.Prepare("insert into record (`id`, `v`, `version`,`time`, `type`, `details`) values (?,?,?,?,?,?);")
	b.e()
	b.selectStmt, b.err = b.conn.Prepare("SELECT `v`, `version`, `details`, `type`  FROM record WHERE id = ?;")
	b.e()
	b.deleteStmt, b.err = b.conn.Prepare("DELETE FROM record WHERE id = ?;")
	b.e()
}

func (b *BattleRecordDb) Insert(batType int, details, v []byte, newId int64) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf(`Battle Record Insert
Error      = %v
BattleId   = %v
BattleType = %v
Stack      =
%s`,
				err,
				newId,
				batType,
				debug.Stack(1, "    "),
			)
		}
	}()

	b.insertStmt.Reset()
	b.err = b.insertStmt.Exec(newId, v, b.srvVersion, time.Now().Unix(), batType, details)
	b.e()

	b.insertStmt.Next()
	b.err = b.insertStmt.Error()
	b.e()
}

func (b *BattleRecordDb) getRecord(req *getRecordRequest) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf(`Battle Record Get
Error    = %v
BattleId = %v
Stack    =
%s`,
				err,
				req.id,
				debug.Stack(1, "    "),
			)
		}
	}()

	var (
		v, details []byte
		version    int64
		batType    int
	)

	b.selectStmt.Reset()
	b.err = b.selectStmt.Exec(req.id)
	b.e()

	b.selectStmt.Next()
	b.err = b.selectStmt.Error()
	b.e()

	b.err = b.selectStmt.Scan(&v, &version, &details, &batType)
	b.e()

	req.callback(batType, details, v)
	return
}

func (b *BattleRecordDb) getMaxId() (maxId int64) {
	var stmt *sqlite.Stmt
	stmt, b.err = b.conn.Prepare("SELECT max(id) FROM record")
	b.e()
	defer func() {
		b.err = stmt.Finalize()
		b.e()
	}()
	b.err = stmt.Exec()
	b.e()

	stmt.Next()
	b.err = stmt.Error()
	b.e()

	stmt.Scan(&maxId)

	return
}

func (b *BattleRecordDb) tableIsExist() bool {
	var stmt *sqlite.Stmt
	stmt, b.err = b.conn.Prepare("SELECT tbl_name FROM sqlite_master WHERE type='table' AND tbl_name='record'")
	b.e()
	defer func() {
		b.err = stmt.Finalize()
		b.e()
	}()
	b.err = stmt.Exec()
	b.e()

	stmt.Next()
	b.err = stmt.Error()
	b.e()

	var (
		tableName string
	)
	stmt.Scan(&tableName)
	if tableName != "record" {
		return false
	}
	return true
}

func (b *BattleRecordDb) Delete(id int64) {
	b.deleteStmt.Reset()
	b.err = b.deleteStmt.Exec(id)
	b.e()
	b.deleteStmt.Next()
	return
}

func (b *BattleRecordDb) GetRecord(id int64, callback func(batType int, details, v []byte)) {
	req := &getRecordRequest{
		id:       id,
		callback: callback,
	}

	// 避免管道阻塞 引起数据库transaction锁太久～
	select {
	case b.selectCh <- req:
	default:
		go func() {
			b.selectCh <- req
		}()
	}
}

// TODO 战斗可以使用全局对象，但是需要Release
func (b *BattleRecordDb) Release() {
	b.err = b.insertStmt.Finalize()
	b.e()
	b.err = b.selectStmt.Finalize()
	b.e()
	b.err = b.deleteStmt.Finalize()
	b.e()
	b.err = b.conn.Close()
	b.e()
}

func (b *BattleRecordDb) e() {
	if b.err != nil {
		panic(b.err)
	}
}

type IRecordData interface {
	UnPack(*uintptr)
}

// 测试打印
func PrintBattleRecord(reader io.Reader, newIRecordStruct func(uint8, uint8) IRecordData) {
	for {
		headBuf := make([]byte, 4)
		recvBuf := make([]byte, 1024)
		// 接收消息头
		if _, err := io.ReadFull(reader, headBuf); err != nil {
			break
		}
		fmt.Println("head", headBuf)

		// 获取消息长度
		var packetLen uint32 = *(*uint32)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&headBuf)).Data))
		if len(recvBuf) < int(packetLen) {
			recvBuf = make([]byte, packetLen)
		}

		// 接收消息主体
		bytes := recvBuf[0:packetLen]
		if _, err := io.ReadFull(reader, bytes); err != nil {
			break
		}
		fmt.Println("body", bytes)

		d := (*reflect.SliceHeader)(unsafe.Pointer(&bytes)).Data

		// 模块ID和操作ID
		moduleId := *(*uint8)(unsafe.Pointer(d))
		d += 1
		actionId := *(*uint8)(unsafe.Pointer(d))
		d += 1

		rsp := newIRecordStruct(moduleId, actionId)

		rsp.UnPack(&d)

		util.PrintN(rsp)
	}
}
*/
