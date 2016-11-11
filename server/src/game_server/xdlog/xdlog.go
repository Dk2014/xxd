package xdlog

import (
	. "game_server/config"
	"io"
	"sync/atomic"
	"time"
)

var (
	g_GameSrvID int32
	g_EID       uint32
)

// xdlog模块要求所有日志类型都实现Log接口
type Log interface {
	Packet() []byte
}

// xdlog模块对外只提供一个统一的Send方法
// func Send(log Log) (int, error) {
// 	return s_tcpConn.Write(log.Packet())
// }

// func SendRaw(raw []byte) (int, error) {
// 	return s_tcpConn.Write(raw)
// }

// 参数初始化
func Init(srvID int32) {

	g_GameSrvID = srvID

	//logHeart()
	// if ServerCfg.EnableGlobalServer == false {
	// 	snapshot()
	// }

	// return err
}

// func Stop() {
// 	if s_tcpConn != nil {
// 		s_tcpConn.Close()
// 	}
// }

// func snapshot() {
// 	timeNow := coretime.GetNowTime()
// 	timeNextDayZero := coretime.GetNextMonthPointHour(0, 0, time.Now())
// 	needSeconds := timeNextDayZero - timeNow
// 	needSeconds = 30
// 	time.AfterFunc(time.Second*time.Duration(needSeconds), func() {
// 		mdb.Transaction(mdb.TRANS_TAG_XdlogSnaps, func() {
// 			mdb.GlobalExecute(func(globaldb *mdb.Database) {
// 				SnapshotLog(db)
// 			})
// 		})
// 		snapshot()
// 	})
// }

// func logHeart() {
// 	time.AfterFunc(time.Minute*1, func() {
// 		mdb.Transaction(mdb.TRANS_TAG_XdlogHeart, func() {
// 			mdb.GlobalExecute(func(db *mdb.Database) {
// 				db.AddXdLog(NewGameSvrState(s_hostAddr))
// 			})
// 		})

// 		logHeart()
// 	})
// }

// 测试用，packet写入文件，使用工具检查是否正确
func debugWriteToFile(log Log, f io.Writer) {
	buf := log.Packet()
	buf = append(buf, []byte("\n")...)
	f.Write(buf)
}

func createeid(ts time.Time) uint64 {
	var hdserver uint64
	if ServerCfg.EnableGlobalServer {
		hdserver = 1
	}
	eid := uint16(atomic.AddUint32(&g_EID, 1))
	return (uint64(ts.Unix()) << 16) | (hdserver << 47) | uint64(eid)

}
