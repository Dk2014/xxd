package tlog

import (
	"io"
	"net"
)

var (
	s_addr *net.UDPAddr

	g_GamePlatID int32 // 0 ios; 1 android
	g_GameAppID  string
	g_GameSrvID  string
)

// tlog模块要求所有日志类型都实现Log接口
type Log interface {
	Packet() []byte
}

// tlog模块对外只提供一个统一的Send方法
func Send(log Log) (int, error) {
	packet := log.Packet()
	//UDP
	conn, e := net.DialUDP("udp4", nil, s_addr)
	if e != nil {
		return 0, e
	}
	defer conn.Close()
	return conn.Write(packet)
}

// 初始化服务器地址
func Init(server, appID, srvID string, platID int32) error {
	var err error
	s_addr, err = net.ResolveUDPAddr("udp4", server)

	g_GamePlatID = platID
	g_GameAppID = appID
	g_GameSrvID = srvID

	return err
}

// 测试用，packet写入文件，使用工具检查是否正确
func debugWriteToFile(log Log, f io.Writer) {
	buf := log.Packet()
	buf = append(buf, []byte("\n")...)
	f.Write(buf)
}
