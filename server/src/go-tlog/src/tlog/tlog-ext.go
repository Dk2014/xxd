// 本文件由 tlog-parser 根据 日志xml规范生成
// 请勿手动改写！

package tlog
import (
	"time"
	"strconv"
	"game_server/mdb"
	glog "core/log"
)

var layout = "2006-01-02 15:04:05"

// 此结构体由工具按XML描述自动生成
type GameSvrState struct {
	DtEventTime time.Time // (必填) 格式 YYYY-MM-DD HH:MM:SS
	VGameIP string // (必填)服务器IP

}

// Log接口要求的Packet方法
func (log *GameSvrState) Packet() []byte {
	struct_name := "GameSvrState"
	buff := make([]byte, 0, 66)
	buff = append(buff, []byte(struct_name)...)
	
	buff = append(buff, '|')
	buff = append(buff, []byte(log.DtEventTime.Format(layout))...)
	
	buff = append(buff, '|')
	if len(log.VGameIP) > 32 {
		buff = append(buff, []byte(log.VGameIP[:32])...)
	} else {
		buff = append(buff, []byte(log.VGameIP)...)
	}
	
	buff = append(buff, '\n')
	return buff[:len(buff)]
}

func (log *GameSvrState) InvokeHook() {}
func (log *GameSvrState) CommitToTLog() {}
func (log *GameSvrState) CommitToXdLog() {}
func (log *GameSvrState) CommitToMySql() error { return nil }
func (log *GameSvrState) Rollback() {}
func (log *GameSvrState) Free() {}
func (log *GameSvrState) SetEventId(time.Time) {}

func (log *GameSvrState) GetTransType() int8 {
	return mdb.TRANS_TYPE_TLOG
}

func (log *GameSvrState) CommitToFile(f *mdb.SyncFile) {
	buff:=log.Packet()
	f.WriteTLog(buff)
	if _, err := SendRaw(buff); err != nil {
		glog.Errorf("[tlog] GameSvrState error. %v\n info: %v", err, log)
	}
}

// 严格的 NewStucture functions 防止松散的创建方式带来未初始化的字段
func NewGameSvrState(VGameIP string) *GameSvrState {
	s_GameSvrState:=new(GameSvrState)
	
	s_GameSvrState.DtEventTime = time.Now()
	
	
	s_GameSvrState.VGameIP=VGameIP
	return s_GameSvrState
}

// 此结构体由工具按XML描述自动生成
type PlayerRegister struct {
	GameSvrId string // (必填)登录的游戏服务器编号
	DtEventTime time.Time // (必填)游戏事件的时间, 格式 YYYY-MM-DD HH:MM:SS
	VGameAppid string // (必填)游戏APPID
	PlatID int32 // (必填)ios 0 /android 1
	Vopenid string // (必填)用户OPENID号
	ClientVersion string // (可选)客户端版本
	SystemSoftware string // (可选)移动终端操作系统版本
	SystemHardware string // (可选)移动终端机型
	TelecomOper string // (必填)运营商
	Network string // (可选)3G/WIFI/2G
	ScreenWidth int32 // (可选)显示屏宽度
	ScreenHight int32 // (可选)显示屏高度
	Density float64 // (可选)像素密度
	RegChannel int32 // (必填)注册渠道
	CpuHardware string // (可选)cpu类型|频率|核数
	Memory int32 // (可选)内存信息单位M
	GLRender string // (可选)opengl render信息
	GLVersion string // (可选)opengl版本信息
	DeviceId string // (可选)设备ID

}

// Log接口要求的Packet方法
func (log *PlayerRegister) Packet() []byte {
	struct_name := "PlayerRegister"
	buff := make([]byte, 0, 798)
	buff = append(buff, []byte(struct_name)...)
	
	buff = append(buff, '|')
	if len(log.GameSvrId) > 25 {
		buff = append(buff, []byte(log.GameSvrId[:25])...)
	} else {
		buff = append(buff, []byte(log.GameSvrId)...)
	}
	
	buff = append(buff, '|')
	buff = append(buff, []byte(log.DtEventTime.Format(layout))...)
	
	buff = append(buff, '|')
	if len(log.VGameAppid) > 32 {
		buff = append(buff, []byte(log.VGameAppid[:32])...)
	} else {
		buff = append(buff, []byte(log.VGameAppid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.PlatID), 10)
	
	buff = append(buff, '|')
	if len(log.Vopenid) > 64 {
		buff = append(buff, []byte(log.Vopenid[:64])...)
	} else {
		buff = append(buff, []byte(log.Vopenid)...)
	}
	
	buff = append(buff, '|')
	if len(log.ClientVersion) > 64 {
		buff = append(buff, []byte(log.ClientVersion[:64])...)
	} else {
		buff = append(buff, []byte(log.ClientVersion)...)
	}
	
	buff = append(buff, '|')
	if len(log.SystemSoftware) > 64 {
		buff = append(buff, []byte(log.SystemSoftware[:64])...)
	} else {
		buff = append(buff, []byte(log.SystemSoftware)...)
	}
	
	buff = append(buff, '|')
	if len(log.SystemHardware) > 64 {
		buff = append(buff, []byte(log.SystemHardware[:64])...)
	} else {
		buff = append(buff, []byte(log.SystemHardware)...)
	}
	
	buff = append(buff, '|')
	if len(log.TelecomOper) > 64 {
		buff = append(buff, []byte(log.TelecomOper[:64])...)
	} else {
		buff = append(buff, []byte(log.TelecomOper)...)
	}
	
	buff = append(buff, '|')
	if len(log.Network) > 64 {
		buff = append(buff, []byte(log.Network[:64])...)
	} else {
		buff = append(buff, []byte(log.Network)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.ScreenWidth), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.ScreenHight), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendFloat(buff, float64(log.Density), 'g', 10, 64)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.RegChannel), 10)
	
	buff = append(buff, '|')
	if len(log.CpuHardware) > 64 {
		buff = append(buff, []byte(log.CpuHardware[:64])...)
	} else {
		buff = append(buff, []byte(log.CpuHardware)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Memory), 10)
	
	buff = append(buff, '|')
	if len(log.GLRender) > 64 {
		buff = append(buff, []byte(log.GLRender[:64])...)
	} else {
		buff = append(buff, []byte(log.GLRender)...)
	}
	
	buff = append(buff, '|')
	if len(log.GLVersion) > 64 {
		buff = append(buff, []byte(log.GLVersion[:64])...)
	} else {
		buff = append(buff, []byte(log.GLVersion)...)
	}
	
	buff = append(buff, '|')
	if len(log.DeviceId) > 64 {
		buff = append(buff, []byte(log.DeviceId[:64])...)
	} else {
		buff = append(buff, []byte(log.DeviceId)...)
	}
	
	buff = append(buff, '\n')
	return buff[:len(buff)]
}

func (log *PlayerRegister) InvokeHook() {}
func (log *PlayerRegister) CommitToTLog() {}
func (log *PlayerRegister) CommitToXdLog() {}
func (log *PlayerRegister) CommitToMySql() error { return nil }
func (log *PlayerRegister) Rollback() {}
func (log *PlayerRegister) Free() {}
func (log *PlayerRegister) SetEventId(time.Time) {}

func (log *PlayerRegister) GetTransType() int8 {
	return mdb.TRANS_TYPE_TLOG
}

func (log *PlayerRegister) CommitToFile(f *mdb.SyncFile) {
	buff:=log.Packet()
	f.WriteTLog(buff)
	if _, err := SendRaw(buff); err != nil {
		glog.Errorf("[tlog] PlayerRegister error. %v\n info: %v", err, log)
	}
}

// 严格的 NewStucture functions 防止松散的创建方式带来未初始化的字段
func NewPlayerRegister(Vopenid string, ClientVersion string, SystemSoftware string, SystemHardware string, TelecomOper string, Network string, ScreenWidth int32, ScreenHight int32, Density float64, RegChannel int32, CpuHardware string, Memory int32, GLRender string, GLVersion string, DeviceId string) *PlayerRegister {
	s_PlayerRegister:=new(PlayerRegister)
	
	s_PlayerRegister.GameSvrId = g_GameSrvID
	s_PlayerRegister.VGameAppid = g_GameAppID
	s_PlayerRegister.PlatID = g_GamePlatID
	s_PlayerRegister.DtEventTime = time.Now()
	
	
	s_PlayerRegister.Vopenid=Vopenid
	s_PlayerRegister.ClientVersion=ClientVersion
	s_PlayerRegister.SystemSoftware=SystemSoftware
	s_PlayerRegister.SystemHardware=SystemHardware
	s_PlayerRegister.TelecomOper=TelecomOper
	s_PlayerRegister.Network=Network
	s_PlayerRegister.ScreenWidth=ScreenWidth
	s_PlayerRegister.ScreenHight=ScreenHight
	s_PlayerRegister.Density=Density
	s_PlayerRegister.RegChannel=RegChannel
	s_PlayerRegister.CpuHardware=CpuHardware
	s_PlayerRegister.Memory=Memory
	s_PlayerRegister.GLRender=GLRender
	s_PlayerRegister.GLVersion=GLVersion
	s_PlayerRegister.DeviceId=DeviceId
	return s_PlayerRegister
}

// 此结构体由工具按XML描述自动生成
type PlayerLogin struct {
	GameSvrId string // (必填)登录的游戏服务器编号
	DtEventTime time.Time // (必填)游戏事件的时间, 格式 YYYY-MM-DD HH:MM:SS
	VGameAppid string // (必填)游戏APPID
	PlatID int32 // (必填)ios 0/android 1
	Vopenid string // (必填)用户OPENID号
	Level int32 // (必填)等级
	PlayerFriendsNum int32 // (必填)玩家好友数量
	ClientVersion string // (必填)客户端版本
	SystemSoftware string // (可选)移动终端操作系统版本
	SystemHardware string // (必填)移动终端机型
	TelecomOper string // (必填)运营商
	Network string // (必填)3G/WIFI/2G
	ScreenWidth int32 // (可选)显示屏宽度
	ScreenHight int32 // (可选)显示屏高度
	Density float64 // (可选)像素密度
	LoginChannel int32 // (必填)登录渠道
	CpuHardware string // (可选)cpu类型|频率|核数
	Memory int32 // (可选)内存信息单位M
	GLRender string // (可选)opengl render信息
	GLVersion string // (可选)opengl版本信息
	DeviceId string // (可选)设备ID

}

// Log接口要求的Packet方法
func (log *PlayerLogin) Packet() []byte {
	struct_name := "PlayerLogin"
	buff := make([]byte, 0, 813)
	buff = append(buff, []byte(struct_name)...)
	
	buff = append(buff, '|')
	if len(log.GameSvrId) > 25 {
		buff = append(buff, []byte(log.GameSvrId[:25])...)
	} else {
		buff = append(buff, []byte(log.GameSvrId)...)
	}
	
	buff = append(buff, '|')
	buff = append(buff, []byte(log.DtEventTime.Format(layout))...)
	
	buff = append(buff, '|')
	if len(log.VGameAppid) > 32 {
		buff = append(buff, []byte(log.VGameAppid[:32])...)
	} else {
		buff = append(buff, []byte(log.VGameAppid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.PlatID), 10)
	
	buff = append(buff, '|')
	if len(log.Vopenid) > 64 {
		buff = append(buff, []byte(log.Vopenid[:64])...)
	} else {
		buff = append(buff, []byte(log.Vopenid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Level), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.PlayerFriendsNum), 10)
	
	buff = append(buff, '|')
	if len(log.ClientVersion) > 64 {
		buff = append(buff, []byte(log.ClientVersion[:64])...)
	} else {
		buff = append(buff, []byte(log.ClientVersion)...)
	}
	
	buff = append(buff, '|')
	if len(log.SystemSoftware) > 64 {
		buff = append(buff, []byte(log.SystemSoftware[:64])...)
	} else {
		buff = append(buff, []byte(log.SystemSoftware)...)
	}
	
	buff = append(buff, '|')
	if len(log.SystemHardware) > 64 {
		buff = append(buff, []byte(log.SystemHardware[:64])...)
	} else {
		buff = append(buff, []byte(log.SystemHardware)...)
	}
	
	buff = append(buff, '|')
	if len(log.TelecomOper) > 64 {
		buff = append(buff, []byte(log.TelecomOper[:64])...)
	} else {
		buff = append(buff, []byte(log.TelecomOper)...)
	}
	
	buff = append(buff, '|')
	if len(log.Network) > 64 {
		buff = append(buff, []byte(log.Network[:64])...)
	} else {
		buff = append(buff, []byte(log.Network)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.ScreenWidth), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.ScreenHight), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendFloat(buff, float64(log.Density), 'g', 10, 64)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.LoginChannel), 10)
	
	buff = append(buff, '|')
	if len(log.CpuHardware) > 64 {
		buff = append(buff, []byte(log.CpuHardware[:64])...)
	} else {
		buff = append(buff, []byte(log.CpuHardware)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Memory), 10)
	
	buff = append(buff, '|')
	if len(log.GLRender) > 64 {
		buff = append(buff, []byte(log.GLRender[:64])...)
	} else {
		buff = append(buff, []byte(log.GLRender)...)
	}
	
	buff = append(buff, '|')
	if len(log.GLVersion) > 64 {
		buff = append(buff, []byte(log.GLVersion[:64])...)
	} else {
		buff = append(buff, []byte(log.GLVersion)...)
	}
	
	buff = append(buff, '|')
	if len(log.DeviceId) > 64 {
		buff = append(buff, []byte(log.DeviceId[:64])...)
	} else {
		buff = append(buff, []byte(log.DeviceId)...)
	}
	
	buff = append(buff, '\n')
	return buff[:len(buff)]
}

func (log *PlayerLogin) InvokeHook() {}
func (log *PlayerLogin) CommitToTLog() {}
func (log *PlayerLogin) CommitToXdLog() {}
func (log *PlayerLogin) CommitToMySql() error { return nil }
func (log *PlayerLogin) Rollback() {}
func (log *PlayerLogin) Free() {}
func (log *PlayerLogin) SetEventId(time.Time) {}

func (log *PlayerLogin) GetTransType() int8 {
	return mdb.TRANS_TYPE_TLOG
}

func (log *PlayerLogin) CommitToFile(f *mdb.SyncFile) {
	buff:=log.Packet()
	f.WriteTLog(buff)
	if _, err := SendRaw(buff); err != nil {
		glog.Errorf("[tlog] PlayerLogin error. %v\n info: %v", err, log)
	}
}

// 严格的 NewStucture functions 防止松散的创建方式带来未初始化的字段
func NewPlayerLogin(Vopenid string, Level int32, PlayerFriendsNum int32, ClientVersion string, SystemSoftware string, SystemHardware string, TelecomOper string, Network string, ScreenWidth int32, ScreenHight int32, Density float64, LoginChannel int32, CpuHardware string, Memory int32, GLRender string, GLVersion string, DeviceId string) *PlayerLogin {
	s_PlayerLogin:=new(PlayerLogin)
	
	s_PlayerLogin.GameSvrId = g_GameSrvID
	s_PlayerLogin.VGameAppid = g_GameAppID
	s_PlayerLogin.PlatID = g_GamePlatID
	s_PlayerLogin.DtEventTime = time.Now()
	
	
	s_PlayerLogin.Vopenid=Vopenid
	s_PlayerLogin.Level=Level
	s_PlayerLogin.PlayerFriendsNum=PlayerFriendsNum
	s_PlayerLogin.ClientVersion=ClientVersion
	s_PlayerLogin.SystemSoftware=SystemSoftware
	s_PlayerLogin.SystemHardware=SystemHardware
	s_PlayerLogin.TelecomOper=TelecomOper
	s_PlayerLogin.Network=Network
	s_PlayerLogin.ScreenWidth=ScreenWidth
	s_PlayerLogin.ScreenHight=ScreenHight
	s_PlayerLogin.Density=Density
	s_PlayerLogin.LoginChannel=LoginChannel
	s_PlayerLogin.CpuHardware=CpuHardware
	s_PlayerLogin.Memory=Memory
	s_PlayerLogin.GLRender=GLRender
	s_PlayerLogin.GLVersion=GLVersion
	s_PlayerLogin.DeviceId=DeviceId
	return s_PlayerLogin
}

// 此结构体由工具按XML描述自动生成
type PlayerLogout struct {
	GameSvrId string // (必填)登录的游戏服务器编号
	DtEventTime time.Time // (必填)游戏事件的时间, 格式 YYYY-MM-DD HH:MM:SS
	VGameAppid string // (必填)游戏APPID
	PlatID int32 // (必填)ios 0/android 1
	Vopenid string // (必填)用户OPENID号
	OnlineTime int32 // (必填)本次登录在线时间(秒)
	Level int32 // (必填)等级
	PlayerFriendsNum int32 // (必填)玩家好友数量
	ClientVersion string // (必填)客户端版本
	SystemSoftware string // (可选)移动终端操作系统版本
	SystemHardware string // (必填)移动终端机型
	TelecomOper string // (必填)运营商
	Network string // (必填)3G/WIFI/2G
	ScreenWidth int32 // (可选)显示屏宽度
	ScreenHight int32 // (可选)显示高度
	Density float64 // (可选)像素密度
	LoginChannel int32 // (可选)登录渠道
	CpuHardware string // (可选)cpu类型|频率|核数
	Memory int32 // (可选)内存信息单位M
	GLRender string // (可选)opengl render信息
	GLVersion string // (可选)opengl版本信息
	DeviceId string // (可选)设备ID
	Vip int32 // (可选)VIP等级

}

// Log接口要求的Packet方法
func (log *PlayerLogout) Packet() []byte {
	struct_name := "PlayerLogout"
	buff := make([]byte, 0, 832)
	buff = append(buff, []byte(struct_name)...)
	
	buff = append(buff, '|')
	if len(log.GameSvrId) > 25 {
		buff = append(buff, []byte(log.GameSvrId[:25])...)
	} else {
		buff = append(buff, []byte(log.GameSvrId)...)
	}
	
	buff = append(buff, '|')
	buff = append(buff, []byte(log.DtEventTime.Format(layout))...)
	
	buff = append(buff, '|')
	if len(log.VGameAppid) > 32 {
		buff = append(buff, []byte(log.VGameAppid[:32])...)
	} else {
		buff = append(buff, []byte(log.VGameAppid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.PlatID), 10)
	
	buff = append(buff, '|')
	if len(log.Vopenid) > 64 {
		buff = append(buff, []byte(log.Vopenid[:64])...)
	} else {
		buff = append(buff, []byte(log.Vopenid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.OnlineTime), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Level), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.PlayerFriendsNum), 10)
	
	buff = append(buff, '|')
	if len(log.ClientVersion) > 64 {
		buff = append(buff, []byte(log.ClientVersion[:64])...)
	} else {
		buff = append(buff, []byte(log.ClientVersion)...)
	}
	
	buff = append(buff, '|')
	if len(log.SystemSoftware) > 64 {
		buff = append(buff, []byte(log.SystemSoftware[:64])...)
	} else {
		buff = append(buff, []byte(log.SystemSoftware)...)
	}
	
	buff = append(buff, '|')
	if len(log.SystemHardware) > 64 {
		buff = append(buff, []byte(log.SystemHardware[:64])...)
	} else {
		buff = append(buff, []byte(log.SystemHardware)...)
	}
	
	buff = append(buff, '|')
	if len(log.TelecomOper) > 64 {
		buff = append(buff, []byte(log.TelecomOper[:64])...)
	} else {
		buff = append(buff, []byte(log.TelecomOper)...)
	}
	
	buff = append(buff, '|')
	if len(log.Network) > 64 {
		buff = append(buff, []byte(log.Network[:64])...)
	} else {
		buff = append(buff, []byte(log.Network)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.ScreenWidth), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.ScreenHight), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendFloat(buff, float64(log.Density), 'g', 10, 64)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.LoginChannel), 10)
	
	buff = append(buff, '|')
	if len(log.CpuHardware) > 64 {
		buff = append(buff, []byte(log.CpuHardware[:64])...)
	} else {
		buff = append(buff, []byte(log.CpuHardware)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Memory), 10)
	
	buff = append(buff, '|')
	if len(log.GLRender) > 64 {
		buff = append(buff, []byte(log.GLRender[:64])...)
	} else {
		buff = append(buff, []byte(log.GLRender)...)
	}
	
	buff = append(buff, '|')
	if len(log.GLVersion) > 64 {
		buff = append(buff, []byte(log.GLVersion[:64])...)
	} else {
		buff = append(buff, []byte(log.GLVersion)...)
	}
	
	buff = append(buff, '|')
	if len(log.DeviceId) > 64 {
		buff = append(buff, []byte(log.DeviceId[:64])...)
	} else {
		buff = append(buff, []byte(log.DeviceId)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Vip), 10)
	
	buff = append(buff, '\n')
	return buff[:len(buff)]
}

func (log *PlayerLogout) InvokeHook() {}
func (log *PlayerLogout) CommitToTLog() {}
func (log *PlayerLogout) CommitToXdLog() {}
func (log *PlayerLogout) CommitToMySql() error { return nil }
func (log *PlayerLogout) Rollback() {}
func (log *PlayerLogout) Free() {}
func (log *PlayerLogout) SetEventId(time.Time) {}

func (log *PlayerLogout) GetTransType() int8 {
	return mdb.TRANS_TYPE_TLOG
}

func (log *PlayerLogout) CommitToFile(f *mdb.SyncFile) {
	buff:=log.Packet()
	f.WriteTLog(buff)
	if _, err := SendRaw(buff); err != nil {
		glog.Errorf("[tlog] PlayerLogout error. %v\n info: %v", err, log)
	}
}

// 严格的 NewStucture functions 防止松散的创建方式带来未初始化的字段
func NewPlayerLogout(Vopenid string, OnlineTime int32, Level int32, PlayerFriendsNum int32, ClientVersion string, SystemSoftware string, SystemHardware string, TelecomOper string, Network string, ScreenWidth int32, ScreenHight int32, Density float64, LoginChannel int32, CpuHardware string, Memory int32, GLRender string, GLVersion string, DeviceId string, Vip int32) *PlayerLogout {
	s_PlayerLogout:=new(PlayerLogout)
	
	s_PlayerLogout.GameSvrId = g_GameSrvID
	s_PlayerLogout.VGameAppid = g_GameAppID
	s_PlayerLogout.PlatID = g_GamePlatID
	s_PlayerLogout.DtEventTime = time.Now()
	
	
	s_PlayerLogout.Vopenid=Vopenid
	s_PlayerLogout.OnlineTime=OnlineTime
	s_PlayerLogout.Level=Level
	s_PlayerLogout.PlayerFriendsNum=PlayerFriendsNum
	s_PlayerLogout.ClientVersion=ClientVersion
	s_PlayerLogout.SystemSoftware=SystemSoftware
	s_PlayerLogout.SystemHardware=SystemHardware
	s_PlayerLogout.TelecomOper=TelecomOper
	s_PlayerLogout.Network=Network
	s_PlayerLogout.ScreenWidth=ScreenWidth
	s_PlayerLogout.ScreenHight=ScreenHight
	s_PlayerLogout.Density=Density
	s_PlayerLogout.LoginChannel=LoginChannel
	s_PlayerLogout.CpuHardware=CpuHardware
	s_PlayerLogout.Memory=Memory
	s_PlayerLogout.GLRender=GLRender
	s_PlayerLogout.GLVersion=GLVersion
	s_PlayerLogout.DeviceId=DeviceId
	s_PlayerLogout.Vip=Vip
	return s_PlayerLogout
}

// 此结构体由工具按XML描述自动生成
type MoneyFlow struct {
	GameSvrId string // (必填)登录的游戏服务器编号
	DtEventTime time.Time // (必填)游戏事件的时间, 格式 YYYY-MM-DD HH:MM:SS
	Sequence int32 // (可选)用于关联一次动作产生多条不同类型的货币流动日志
	VGameAppid string // (必填)游戏APPID
	PlatID int32 // (必填)ios 0/android 1
	Vopenid string // (必填)用户OPENID号
	Level int32 // (必填)玩家等级
	AfterMoney int32 // (可选)动作后的金钱数
	IMoney int32 // (必填)动作涉及的金钱数
	Reason int32 // (必填)货币流动一级原因
	SubReason int32 // (可选)货币流动二级原因
	AddOrReduce int32 // (必填)增加 0/减少 1
	IMoneyType int32 // (必填)钱的类型MONEYTYPE,其它货币类型参考FAQ文档
	Vip int32 // (必填)Vip等级

}

// Log接口要求的Packet方法
func (log *MoneyFlow) Packet() []byte {
	struct_name := "MoneyFlow"
	buff := make([]byte, 0, 244)
	buff = append(buff, []byte(struct_name)...)
	
	buff = append(buff, '|')
	if len(log.GameSvrId) > 25 {
		buff = append(buff, []byte(log.GameSvrId[:25])...)
	} else {
		buff = append(buff, []byte(log.GameSvrId)...)
	}
	
	buff = append(buff, '|')
	buff = append(buff, []byte(log.DtEventTime.Format(layout))...)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Sequence), 10)
	
	buff = append(buff, '|')
	if len(log.VGameAppid) > 32 {
		buff = append(buff, []byte(log.VGameAppid[:32])...)
	} else {
		buff = append(buff, []byte(log.VGameAppid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.PlatID), 10)
	
	buff = append(buff, '|')
	if len(log.Vopenid) > 64 {
		buff = append(buff, []byte(log.Vopenid[:64])...)
	} else {
		buff = append(buff, []byte(log.Vopenid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Level), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.AfterMoney), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.IMoney), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Reason), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.SubReason), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.AddOrReduce), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.IMoneyType), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Vip), 10)
	
	buff = append(buff, '\n')
	return buff[:len(buff)]
}

func (log *MoneyFlow) InvokeHook() {}
func (log *MoneyFlow) CommitToTLog() {}
func (log *MoneyFlow) CommitToXdLog() {}
func (log *MoneyFlow) CommitToMySql() error { return nil }
func (log *MoneyFlow) Rollback() {}
func (log *MoneyFlow) Free() {}
func (log *MoneyFlow) SetEventId(time.Time) {}

func (log *MoneyFlow) GetTransType() int8 {
	return mdb.TRANS_TYPE_TLOG
}

func (log *MoneyFlow) CommitToFile(f *mdb.SyncFile) {
	buff:=log.Packet()
	f.WriteTLog(buff)
	if _, err := SendRaw(buff); err != nil {
		glog.Errorf("[tlog] MoneyFlow error. %v\n info: %v", err, log)
	}
}

// 严格的 NewStucture functions 防止松散的创建方式带来未初始化的字段
func NewMoneyFlow(Sequence int32, Vopenid string, Level int32, AfterMoney int32, IMoney int32, Reason int32, SubReason int32, AddOrReduce int32, IMoneyType int32, Vip int32) *MoneyFlow {
	s_MoneyFlow:=new(MoneyFlow)
	
	s_MoneyFlow.GameSvrId = g_GameSrvID
	s_MoneyFlow.VGameAppid = g_GameAppID
	s_MoneyFlow.PlatID = g_GamePlatID
	s_MoneyFlow.DtEventTime = time.Now()
	
	
	s_MoneyFlow.Sequence=Sequence
	s_MoneyFlow.Vopenid=Vopenid
	s_MoneyFlow.Level=Level
	s_MoneyFlow.AfterMoney=AfterMoney
	s_MoneyFlow.IMoney=IMoney
	s_MoneyFlow.Reason=Reason
	s_MoneyFlow.SubReason=SubReason
	s_MoneyFlow.AddOrReduce=AddOrReduce
	s_MoneyFlow.IMoneyType=IMoneyType
	s_MoneyFlow.Vip=Vip
	return s_MoneyFlow
}

// 此结构体由工具按XML描述自动生成
type ItemFlow struct {
	GameSvrId string // (必填)登录的游戏服务器编号
	DtEventTime time.Time // (必填)游戏事件的时间, 格式 YYYY-MM-DD HH:MM:SS
	Sequence int32 // (可选)用于关联一次动作产生多条不同类型的道具流动日志
	VGameAppid string // (必填)游戏APPID
	PlatID int32 // (必填)ios 0/android 1
	Vopenid string // (必填)玩家
	IGoodsType int32 // (必填)道具类型
	IGoodsId int32 // (必填)道具ID
	AfterCount int32 // (必填)动作后的物品存量
	Count int32 // (必填)动作涉及的物品数量
	Reason int32 // (必填)道具流动一级原因
	SubReason int32 // (必填)道具流动二级原因
	AddOrReduce int32 // (必填)增加 0/减少 1
	Vip int32 // (必填)Vip等级

}

// Log接口要求的Packet方法
func (log *ItemFlow) Packet() []byte {
	struct_name := "ItemFlow"
	buff := make([]byte, 0, 243)
	buff = append(buff, []byte(struct_name)...)
	
	buff = append(buff, '|')
	if len(log.GameSvrId) > 25 {
		buff = append(buff, []byte(log.GameSvrId[:25])...)
	} else {
		buff = append(buff, []byte(log.GameSvrId)...)
	}
	
	buff = append(buff, '|')
	buff = append(buff, []byte(log.DtEventTime.Format(layout))...)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Sequence), 10)
	
	buff = append(buff, '|')
	if len(log.VGameAppid) > 32 {
		buff = append(buff, []byte(log.VGameAppid[:32])...)
	} else {
		buff = append(buff, []byte(log.VGameAppid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.PlatID), 10)
	
	buff = append(buff, '|')
	if len(log.Vopenid) > 64 {
		buff = append(buff, []byte(log.Vopenid[:64])...)
	} else {
		buff = append(buff, []byte(log.Vopenid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.IGoodsType), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.IGoodsId), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.AfterCount), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Count), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Reason), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.SubReason), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.AddOrReduce), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Vip), 10)
	
	buff = append(buff, '\n')
	return buff[:len(buff)]
}

func (log *ItemFlow) InvokeHook() {}
func (log *ItemFlow) CommitToTLog() {}
func (log *ItemFlow) CommitToXdLog() {}
func (log *ItemFlow) CommitToMySql() error { return nil }
func (log *ItemFlow) Rollback() {}
func (log *ItemFlow) Free() {}
func (log *ItemFlow) SetEventId(time.Time) {}

func (log *ItemFlow) GetTransType() int8 {
	return mdb.TRANS_TYPE_TLOG
}

func (log *ItemFlow) CommitToFile(f *mdb.SyncFile) {
	buff:=log.Packet()
	f.WriteTLog(buff)
	if _, err := SendRaw(buff); err != nil {
		glog.Errorf("[tlog] ItemFlow error. %v\n info: %v", err, log)
	}
}

// 严格的 NewStucture functions 防止松散的创建方式带来未初始化的字段
func NewItemFlow(Sequence int32, Vopenid string, IGoodsType int32, IGoodsId int32, AfterCount int32, Count int32, Reason int32, SubReason int32, AddOrReduce int32, Vip int32) *ItemFlow {
	s_ItemFlow:=new(ItemFlow)
	
	s_ItemFlow.GameSvrId = g_GameSrvID
	s_ItemFlow.VGameAppid = g_GameAppID
	s_ItemFlow.PlatID = g_GamePlatID
	s_ItemFlow.DtEventTime = time.Now()
	
	
	s_ItemFlow.Sequence=Sequence
	s_ItemFlow.Vopenid=Vopenid
	s_ItemFlow.IGoodsType=IGoodsType
	s_ItemFlow.IGoodsId=IGoodsId
	s_ItemFlow.AfterCount=AfterCount
	s_ItemFlow.Count=Count
	s_ItemFlow.Reason=Reason
	s_ItemFlow.SubReason=SubReason
	s_ItemFlow.AddOrReduce=AddOrReduce
	s_ItemFlow.Vip=Vip
	return s_ItemFlow
}

// 此结构体由工具按XML描述自动生成
type ItemMoneyFlow struct {
	GameSvrId string // (必填)登录的游戏服务器编号
	DtEventTime time.Time // (必填)游戏事件的时间, 格式 YYYY-MM-DD HH:MM:SS
	Sequence int32 // (必填)用于关联一次购买产生多条不同类型的货币日志
	VGameAppid string // (必填)游戏APPID
	PlatID int32 // (必填)ios 0/android 1
	Vopenid string // (必填)玩家
	IGoodsType int32 // (必填)道具类型
	IGoodsId int32 // (必填)道具ID
	Count int32 // (必填)数量
	IMoney int32 // (必填)钱
	Level int32 // (必填)玩家等级
	IMoneyType int32 // (必填)钱的类型MONEYTYPE,其它货币类型参考FAQ文档
	Vip int32 // (必填)Vip等级

}

// Log接口要求的Packet方法
func (log *ItemMoneyFlow) Packet() []byte {
	struct_name := "ItemMoneyFlow"
	buff := make([]byte, 0, 239)
	buff = append(buff, []byte(struct_name)...)
	
	buff = append(buff, '|')
	if len(log.GameSvrId) > 25 {
		buff = append(buff, []byte(log.GameSvrId[:25])...)
	} else {
		buff = append(buff, []byte(log.GameSvrId)...)
	}
	
	buff = append(buff, '|')
	buff = append(buff, []byte(log.DtEventTime.Format(layout))...)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Sequence), 10)
	
	buff = append(buff, '|')
	if len(log.VGameAppid) > 32 {
		buff = append(buff, []byte(log.VGameAppid[:32])...)
	} else {
		buff = append(buff, []byte(log.VGameAppid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.PlatID), 10)
	
	buff = append(buff, '|')
	if len(log.Vopenid) > 64 {
		buff = append(buff, []byte(log.Vopenid[:64])...)
	} else {
		buff = append(buff, []byte(log.Vopenid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.IGoodsType), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.IGoodsId), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Count), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.IMoney), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Level), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.IMoneyType), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Vip), 10)
	
	buff = append(buff, '\n')
	return buff[:len(buff)]
}

func (log *ItemMoneyFlow) InvokeHook() {}
func (log *ItemMoneyFlow) CommitToTLog() {}
func (log *ItemMoneyFlow) CommitToXdLog() {}
func (log *ItemMoneyFlow) CommitToMySql() error { return nil }
func (log *ItemMoneyFlow) Rollback() {}
func (log *ItemMoneyFlow) Free() {}
func (log *ItemMoneyFlow) SetEventId(time.Time) {}

func (log *ItemMoneyFlow) GetTransType() int8 {
	return mdb.TRANS_TYPE_TLOG
}

func (log *ItemMoneyFlow) CommitToFile(f *mdb.SyncFile) {
	buff:=log.Packet()
	f.WriteTLog(buff)
	if _, err := SendRaw(buff); err != nil {
		glog.Errorf("[tlog] ItemMoneyFlow error. %v\n info: %v", err, log)
	}
}

// 严格的 NewStucture functions 防止松散的创建方式带来未初始化的字段
func NewItemMoneyFlow(Sequence int32, Vopenid string, IGoodsType int32, IGoodsId int32, Count int32, IMoney int32, Level int32, IMoneyType int32, Vip int32) *ItemMoneyFlow {
	s_ItemMoneyFlow:=new(ItemMoneyFlow)
	
	s_ItemMoneyFlow.GameSvrId = g_GameSrvID
	s_ItemMoneyFlow.VGameAppid = g_GameAppID
	s_ItemMoneyFlow.PlatID = g_GamePlatID
	s_ItemMoneyFlow.DtEventTime = time.Now()
	
	
	s_ItemMoneyFlow.Sequence=Sequence
	s_ItemMoneyFlow.Vopenid=Vopenid
	s_ItemMoneyFlow.IGoodsType=IGoodsType
	s_ItemMoneyFlow.IGoodsId=IGoodsId
	s_ItemMoneyFlow.Count=Count
	s_ItemMoneyFlow.IMoney=IMoney
	s_ItemMoneyFlow.Level=Level
	s_ItemMoneyFlow.IMoneyType=IMoneyType
	s_ItemMoneyFlow.Vip=Vip
	return s_ItemMoneyFlow
}

// 此结构体由工具按XML描述自动生成
type PlayerExpFlow struct {
	GameSvrId string // (必填)登录的游戏服务器编号
	DtEventTime time.Time // (必填)游戏事件的时间, 格式 YYYY-MM-DD HH:MM:SS
	VGameAppid string // (必填)游戏APPID
	PlatID int32 // (必填)ios 0/android 1
	Vopenid string // (必填)玩家
	ExpChange int32 // (必填)经验变化
	BeforeLevel int32 // (可选)动作前等级
	AfterLevel int32 // (必填)动作后等级
	Time int32 // (必填)升级所用时间(秒)
	Reason int32 // (必填)经验流动一级原因
	SubReason int32 // (必填)经验流动二级原因
	Vip int32 // (必填)Vip等级

}

// Log接口要求的Packet方法
func (log *PlayerExpFlow) Packet() []byte {
	struct_name := "PlayerExpFlow"
	buff := make([]byte, 0, 230)
	buff = append(buff, []byte(struct_name)...)
	
	buff = append(buff, '|')
	if len(log.GameSvrId) > 25 {
		buff = append(buff, []byte(log.GameSvrId[:25])...)
	} else {
		buff = append(buff, []byte(log.GameSvrId)...)
	}
	
	buff = append(buff, '|')
	buff = append(buff, []byte(log.DtEventTime.Format(layout))...)
	
	buff = append(buff, '|')
	if len(log.VGameAppid) > 32 {
		buff = append(buff, []byte(log.VGameAppid[:32])...)
	} else {
		buff = append(buff, []byte(log.VGameAppid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.PlatID), 10)
	
	buff = append(buff, '|')
	if len(log.Vopenid) > 64 {
		buff = append(buff, []byte(log.Vopenid[:64])...)
	} else {
		buff = append(buff, []byte(log.Vopenid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.ExpChange), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.BeforeLevel), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.AfterLevel), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Time), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Reason), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.SubReason), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Vip), 10)
	
	buff = append(buff, '\n')
	return buff[:len(buff)]
}

func (log *PlayerExpFlow) InvokeHook() {}
func (log *PlayerExpFlow) CommitToTLog() {}
func (log *PlayerExpFlow) CommitToXdLog() {}
func (log *PlayerExpFlow) CommitToMySql() error { return nil }
func (log *PlayerExpFlow) Rollback() {}
func (log *PlayerExpFlow) Free() {}
func (log *PlayerExpFlow) SetEventId(time.Time) {}

func (log *PlayerExpFlow) GetTransType() int8 {
	return mdb.TRANS_TYPE_TLOG
}

func (log *PlayerExpFlow) CommitToFile(f *mdb.SyncFile) {
	buff:=log.Packet()
	f.WriteTLog(buff)
	if _, err := SendRaw(buff); err != nil {
		glog.Errorf("[tlog] PlayerExpFlow error. %v\n info: %v", err, log)
	}
}

// 严格的 NewStucture functions 防止松散的创建方式带来未初始化的字段
func NewPlayerExpFlow(Vopenid string, ExpChange int32, BeforeLevel int32, AfterLevel int32, Time int32, Reason int32, SubReason int32, Vip int32) *PlayerExpFlow {
	s_PlayerExpFlow:=new(PlayerExpFlow)
	
	s_PlayerExpFlow.GameSvrId = g_GameSrvID
	s_PlayerExpFlow.VGameAppid = g_GameAppID
	s_PlayerExpFlow.PlatID = g_GamePlatID
	s_PlayerExpFlow.DtEventTime = time.Now()
	
	
	s_PlayerExpFlow.Vopenid=Vopenid
	s_PlayerExpFlow.ExpChange=ExpChange
	s_PlayerExpFlow.BeforeLevel=BeforeLevel
	s_PlayerExpFlow.AfterLevel=AfterLevel
	s_PlayerExpFlow.Time=Time
	s_PlayerExpFlow.Reason=Reason
	s_PlayerExpFlow.SubReason=SubReason
	s_PlayerExpFlow.Vip=Vip
	return s_PlayerExpFlow
}

// 此结构体由工具按XML描述自动生成
type PlayerMissionFlow struct {
	GameSvrId string // (必填)登录的游戏服务器编号
	DtEventTime time.Time // (必填)游戏事件的时间, 格式 YYYY-MM-DD HH:MM:SS
	VGameAppid string // (必填)游戏APPID
	PlatID int32 // (必填)ios 0/android 1
	Vopenid string // (必填)玩家
	Level int32 // (必填)人物等级
	MissionType int32 // (必填)副本类型
	MissionID int32 // (必填)副本ID
	MissionLock int32 // (必填)副本权值
	MaxLock int32 // (必填)人物最大权值
	Vip int32 // (必填)Vip等级
	Action int32 // (必填)进入0/完成1/扫荡2

}

// Log接口要求的Packet方法
func (log *PlayerMissionFlow) Packet() []byte {
	struct_name := "PlayerMissionFlow"
	buff := make([]byte, 0, 234)
	buff = append(buff, []byte(struct_name)...)
	
	buff = append(buff, '|')
	if len(log.GameSvrId) > 25 {
		buff = append(buff, []byte(log.GameSvrId[:25])...)
	} else {
		buff = append(buff, []byte(log.GameSvrId)...)
	}
	
	buff = append(buff, '|')
	buff = append(buff, []byte(log.DtEventTime.Format(layout))...)
	
	buff = append(buff, '|')
	if len(log.VGameAppid) > 32 {
		buff = append(buff, []byte(log.VGameAppid[:32])...)
	} else {
		buff = append(buff, []byte(log.VGameAppid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.PlatID), 10)
	
	buff = append(buff, '|')
	if len(log.Vopenid) > 64 {
		buff = append(buff, []byte(log.Vopenid[:64])...)
	} else {
		buff = append(buff, []byte(log.Vopenid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Level), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.MissionType), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.MissionID), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.MissionLock), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.MaxLock), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Vip), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Action), 10)
	
	buff = append(buff, '\n')
	return buff[:len(buff)]
}

func (log *PlayerMissionFlow) InvokeHook() {}
func (log *PlayerMissionFlow) CommitToTLog() {}
func (log *PlayerMissionFlow) CommitToXdLog() {}
func (log *PlayerMissionFlow) CommitToMySql() error { return nil }
func (log *PlayerMissionFlow) Rollback() {}
func (log *PlayerMissionFlow) Free() {}
func (log *PlayerMissionFlow) SetEventId(time.Time) {}

func (log *PlayerMissionFlow) GetTransType() int8 {
	return mdb.TRANS_TYPE_TLOG
}

func (log *PlayerMissionFlow) CommitToFile(f *mdb.SyncFile) {
	buff:=log.Packet()
	f.WriteTLog(buff)
	if _, err := SendRaw(buff); err != nil {
		glog.Errorf("[tlog] PlayerMissionFlow error. %v\n info: %v", err, log)
	}
}

// 严格的 NewStucture functions 防止松散的创建方式带来未初始化的字段
func NewPlayerMissionFlow(Vopenid string, Level int32, MissionType int32, MissionID int32, MissionLock int32, MaxLock int32, Vip int32, Action int32) *PlayerMissionFlow {
	s_PlayerMissionFlow:=new(PlayerMissionFlow)
	
	s_PlayerMissionFlow.GameSvrId = g_GameSrvID
	s_PlayerMissionFlow.VGameAppid = g_GameAppID
	s_PlayerMissionFlow.PlatID = g_GamePlatID
	s_PlayerMissionFlow.DtEventTime = time.Now()
	
	
	s_PlayerMissionFlow.Vopenid=Vopenid
	s_PlayerMissionFlow.Level=Level
	s_PlayerMissionFlow.MissionType=MissionType
	s_PlayerMissionFlow.MissionID=MissionID
	s_PlayerMissionFlow.MissionLock=MissionLock
	s_PlayerMissionFlow.MaxLock=MaxLock
	s_PlayerMissionFlow.Vip=Vip
	s_PlayerMissionFlow.Action=Action
	return s_PlayerMissionFlow
}

// 此结构体由工具按XML描述自动生成
type PlayerPhysicalFlow struct {
	GameSvrId string // (必填)登录的游戏服务器编号
	DtEventTime time.Time // (必填)游戏事件的时间, 格式 YYYY-MM-DD HH:MM:SS
	VGameAppid string // (必填)游戏APPID
	PlatID int32 // (必填)ios 0/android 1
	Vopenid string // (必填)玩家
	Level int32 // (必填)人物等级
	Atfercount int32 // (必填)动作后体力
	Count int32 // (必填)体力变化
	Vip int32 // (必填)Vip等级
	Reason int32 // (必填)经验流动一级原因
	AddOrReduce int32 // (必填)增加0/减少1

}

// Log接口要求的Packet方法
func (log *PlayerPhysicalFlow) Packet() []byte {
	struct_name := "PlayerPhysicalFlow"
	buff := make([]byte, 0, 226)
	buff = append(buff, []byte(struct_name)...)
	
	buff = append(buff, '|')
	if len(log.GameSvrId) > 25 {
		buff = append(buff, []byte(log.GameSvrId[:25])...)
	} else {
		buff = append(buff, []byte(log.GameSvrId)...)
	}
	
	buff = append(buff, '|')
	buff = append(buff, []byte(log.DtEventTime.Format(layout))...)
	
	buff = append(buff, '|')
	if len(log.VGameAppid) > 32 {
		buff = append(buff, []byte(log.VGameAppid[:32])...)
	} else {
		buff = append(buff, []byte(log.VGameAppid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.PlatID), 10)
	
	buff = append(buff, '|')
	if len(log.Vopenid) > 64 {
		buff = append(buff, []byte(log.Vopenid[:64])...)
	} else {
		buff = append(buff, []byte(log.Vopenid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Level), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Atfercount), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Count), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Vip), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Reason), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.AddOrReduce), 10)
	
	buff = append(buff, '\n')
	return buff[:len(buff)]
}

func (log *PlayerPhysicalFlow) InvokeHook() {}
func (log *PlayerPhysicalFlow) CommitToTLog() {}
func (log *PlayerPhysicalFlow) CommitToXdLog() {}
func (log *PlayerPhysicalFlow) CommitToMySql() error { return nil }
func (log *PlayerPhysicalFlow) Rollback() {}
func (log *PlayerPhysicalFlow) Free() {}
func (log *PlayerPhysicalFlow) SetEventId(time.Time) {}

func (log *PlayerPhysicalFlow) GetTransType() int8 {
	return mdb.TRANS_TYPE_TLOG
}

func (log *PlayerPhysicalFlow) CommitToFile(f *mdb.SyncFile) {
	buff:=log.Packet()
	f.WriteTLog(buff)
	if _, err := SendRaw(buff); err != nil {
		glog.Errorf("[tlog] PlayerPhysicalFlow error. %v\n info: %v", err, log)
	}
}

// 严格的 NewStucture functions 防止松散的创建方式带来未初始化的字段
func NewPlayerPhysicalFlow(Vopenid string, Level int32, Atfercount int32, Count int32, Vip int32, Reason int32, AddOrReduce int32) *PlayerPhysicalFlow {
	s_PlayerPhysicalFlow:=new(PlayerPhysicalFlow)
	
	s_PlayerPhysicalFlow.GameSvrId = g_GameSrvID
	s_PlayerPhysicalFlow.VGameAppid = g_GameAppID
	s_PlayerPhysicalFlow.PlatID = g_GamePlatID
	s_PlayerPhysicalFlow.DtEventTime = time.Now()
	
	
	s_PlayerPhysicalFlow.Vopenid=Vopenid
	s_PlayerPhysicalFlow.Level=Level
	s_PlayerPhysicalFlow.Atfercount=Atfercount
	s_PlayerPhysicalFlow.Count=Count
	s_PlayerPhysicalFlow.Vip=Vip
	s_PlayerPhysicalFlow.Reason=Reason
	s_PlayerPhysicalFlow.AddOrReduce=AddOrReduce
	return s_PlayerPhysicalFlow
}

// 此结构体由工具按XML描述自动生成
type PlayerPvpFlow struct {
	GameSvrId string // (必填)登录的游戏服务器编号
	DtEventTime time.Time // (必填)游戏事件的时间, 格式 YYYY-MM-DD HH:MM:SS
	VGameAppid string // (必填)游戏APPID
	PlatID int32 // (必填)ios 0/android 1
	Vopenid string // (必填)玩家
	Level int32 // (必填)人物等级
	Beforerank int32 // (必填)动作前排名
	Afterrank int32 // (必填)动作后排名
	Vip int32 // (必填)Vip等级
	Action int32 // (必填)增加0/减少1

}

// Log接口要求的Packet方法
func (log *PlayerPvpFlow) Packet() []byte {
	struct_name := "PlayerPvpFlow"
	buff := make([]byte, 0, 212)
	buff = append(buff, []byte(struct_name)...)
	
	buff = append(buff, '|')
	if len(log.GameSvrId) > 25 {
		buff = append(buff, []byte(log.GameSvrId[:25])...)
	} else {
		buff = append(buff, []byte(log.GameSvrId)...)
	}
	
	buff = append(buff, '|')
	buff = append(buff, []byte(log.DtEventTime.Format(layout))...)
	
	buff = append(buff, '|')
	if len(log.VGameAppid) > 32 {
		buff = append(buff, []byte(log.VGameAppid[:32])...)
	} else {
		buff = append(buff, []byte(log.VGameAppid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.PlatID), 10)
	
	buff = append(buff, '|')
	if len(log.Vopenid) > 64 {
		buff = append(buff, []byte(log.Vopenid[:64])...)
	} else {
		buff = append(buff, []byte(log.Vopenid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Level), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Beforerank), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Afterrank), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Vip), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Action), 10)
	
	buff = append(buff, '\n')
	return buff[:len(buff)]
}

func (log *PlayerPvpFlow) InvokeHook() {}
func (log *PlayerPvpFlow) CommitToTLog() {}
func (log *PlayerPvpFlow) CommitToXdLog() {}
func (log *PlayerPvpFlow) CommitToMySql() error { return nil }
func (log *PlayerPvpFlow) Rollback() {}
func (log *PlayerPvpFlow) Free() {}
func (log *PlayerPvpFlow) SetEventId(time.Time) {}

func (log *PlayerPvpFlow) GetTransType() int8 {
	return mdb.TRANS_TYPE_TLOG
}

func (log *PlayerPvpFlow) CommitToFile(f *mdb.SyncFile) {
	buff:=log.Packet()
	f.WriteTLog(buff)
	if _, err := SendRaw(buff); err != nil {
		glog.Errorf("[tlog] PlayerPvpFlow error. %v\n info: %v", err, log)
	}
}

// 严格的 NewStucture functions 防止松散的创建方式带来未初始化的字段
func NewPlayerPvpFlow(Vopenid string, Level int32, Beforerank int32, Afterrank int32, Vip int32, Action int32) *PlayerPvpFlow {
	s_PlayerPvpFlow:=new(PlayerPvpFlow)
	
	s_PlayerPvpFlow.GameSvrId = g_GameSrvID
	s_PlayerPvpFlow.VGameAppid = g_GameAppID
	s_PlayerPvpFlow.PlatID = g_GamePlatID
	s_PlayerPvpFlow.DtEventTime = time.Now()
	
	
	s_PlayerPvpFlow.Vopenid=Vopenid
	s_PlayerPvpFlow.Level=Level
	s_PlayerPvpFlow.Beforerank=Beforerank
	s_PlayerPvpFlow.Afterrank=Afterrank
	s_PlayerPvpFlow.Vip=Vip
	s_PlayerPvpFlow.Action=Action
	return s_PlayerPvpFlow
}

// 此结构体由工具按XML描述自动生成
type PlayerGuideFlow struct {
	GameSvrId string // (必填)登录的游戏服务器编号
	DtEventTime time.Time // (必填)游戏事件的时间, 格式 YYYY-MM-DD HH:MM:SS
	VGameAppid string // (必填)游戏APPID
	PlatID int32 // (必填)ios 0/android 1
	Vopenid string // (必填)玩家
	Level int32 // (必填)人物等级
	Typeid int32 // (必填)引导id
	Vip int32 // (必填)Vip等级
	Action int32 // (必填)接受0/完成1

}

// Log接口要求的Packet方法
func (log *PlayerGuideFlow) Packet() []byte {
	struct_name := "PlayerGuideFlow"
	buff := make([]byte, 0, 205)
	buff = append(buff, []byte(struct_name)...)
	
	buff = append(buff, '|')
	if len(log.GameSvrId) > 25 {
		buff = append(buff, []byte(log.GameSvrId[:25])...)
	} else {
		buff = append(buff, []byte(log.GameSvrId)...)
	}
	
	buff = append(buff, '|')
	buff = append(buff, []byte(log.DtEventTime.Format(layout))...)
	
	buff = append(buff, '|')
	if len(log.VGameAppid) > 32 {
		buff = append(buff, []byte(log.VGameAppid[:32])...)
	} else {
		buff = append(buff, []byte(log.VGameAppid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.PlatID), 10)
	
	buff = append(buff, '|')
	if len(log.Vopenid) > 64 {
		buff = append(buff, []byte(log.Vopenid[:64])...)
	} else {
		buff = append(buff, []byte(log.Vopenid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Level), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Typeid), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Vip), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Action), 10)
	
	buff = append(buff, '\n')
	return buff[:len(buff)]
}

func (log *PlayerGuideFlow) InvokeHook() {}
func (log *PlayerGuideFlow) CommitToTLog() {}
func (log *PlayerGuideFlow) CommitToXdLog() {}
func (log *PlayerGuideFlow) CommitToMySql() error { return nil }
func (log *PlayerGuideFlow) Rollback() {}
func (log *PlayerGuideFlow) Free() {}
func (log *PlayerGuideFlow) SetEventId(time.Time) {}

func (log *PlayerGuideFlow) GetTransType() int8 {
	return mdb.TRANS_TYPE_TLOG
}

func (log *PlayerGuideFlow) CommitToFile(f *mdb.SyncFile) {
	buff:=log.Packet()
	f.WriteTLog(buff)
	if _, err := SendRaw(buff); err != nil {
		glog.Errorf("[tlog] PlayerGuideFlow error. %v\n info: %v", err, log)
	}
}

// 严格的 NewStucture functions 防止松散的创建方式带来未初始化的字段
func NewPlayerGuideFlow(Vopenid string, Level int32, Typeid int32, Vip int32, Action int32) *PlayerGuideFlow {
	s_PlayerGuideFlow:=new(PlayerGuideFlow)
	
	s_PlayerGuideFlow.GameSvrId = g_GameSrvID
	s_PlayerGuideFlow.VGameAppid = g_GameAppID
	s_PlayerGuideFlow.PlatID = g_GamePlatID
	s_PlayerGuideFlow.DtEventTime = time.Now()
	
	
	s_PlayerGuideFlow.Vopenid=Vopenid
	s_PlayerGuideFlow.Level=Level
	s_PlayerGuideFlow.Typeid=Typeid
	s_PlayerGuideFlow.Vip=Vip
	s_PlayerGuideFlow.Action=Action
	return s_PlayerGuideFlow
}

// 此结构体由工具按XML描述自动生成
type PlayerSystemModuleFlow struct {
	GameSvrId string // (必填)登录的游戏服务器编号
	DtEventTime time.Time // (必填)游戏事件的时间, 格式 YYYY-MM-DD HH:MM:SS
	VGameAppid string // (必填)游戏APPID
	PlatID int32 // (必填)ios 0/android 1
	Vopenid string // (必填)玩家
	Level int32 // (必填)人物等级
	Typeid int32 // (必填)模块ID
	Vip int32 // (必填)Vip等级

}

// Log接口要求的Packet方法
func (log *PlayerSystemModuleFlow) Packet() []byte {
	struct_name := "PlayerSystemModuleFlow"
	buff := make([]byte, 0, 203)
	buff = append(buff, []byte(struct_name)...)
	
	buff = append(buff, '|')
	if len(log.GameSvrId) > 25 {
		buff = append(buff, []byte(log.GameSvrId[:25])...)
	} else {
		buff = append(buff, []byte(log.GameSvrId)...)
	}
	
	buff = append(buff, '|')
	buff = append(buff, []byte(log.DtEventTime.Format(layout))...)
	
	buff = append(buff, '|')
	if len(log.VGameAppid) > 32 {
		buff = append(buff, []byte(log.VGameAppid[:32])...)
	} else {
		buff = append(buff, []byte(log.VGameAppid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.PlatID), 10)
	
	buff = append(buff, '|')
	if len(log.Vopenid) > 64 {
		buff = append(buff, []byte(log.Vopenid[:64])...)
	} else {
		buff = append(buff, []byte(log.Vopenid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Level), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Typeid), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Vip), 10)
	
	buff = append(buff, '\n')
	return buff[:len(buff)]
}

func (log *PlayerSystemModuleFlow) InvokeHook() {}
func (log *PlayerSystemModuleFlow) CommitToTLog() {}
func (log *PlayerSystemModuleFlow) CommitToXdLog() {}
func (log *PlayerSystemModuleFlow) CommitToMySql() error { return nil }
func (log *PlayerSystemModuleFlow) Rollback() {}
func (log *PlayerSystemModuleFlow) Free() {}
func (log *PlayerSystemModuleFlow) SetEventId(time.Time) {}

func (log *PlayerSystemModuleFlow) GetTransType() int8 {
	return mdb.TRANS_TYPE_TLOG
}

func (log *PlayerSystemModuleFlow) CommitToFile(f *mdb.SyncFile) {
	buff:=log.Packet()
	f.WriteTLog(buff)
	if _, err := SendRaw(buff); err != nil {
		glog.Errorf("[tlog] PlayerSystemModuleFlow error. %v\n info: %v", err, log)
	}
}

// 严格的 NewStucture functions 防止松散的创建方式带来未初始化的字段
func NewPlayerSystemModuleFlow(Vopenid string, Level int32, Typeid int32, Vip int32) *PlayerSystemModuleFlow {
	s_PlayerSystemModuleFlow:=new(PlayerSystemModuleFlow)
	
	s_PlayerSystemModuleFlow.GameSvrId = g_GameSrvID
	s_PlayerSystemModuleFlow.VGameAppid = g_GameAppID
	s_PlayerSystemModuleFlow.PlatID = g_GamePlatID
	s_PlayerSystemModuleFlow.DtEventTime = time.Now()
	
	
	s_PlayerSystemModuleFlow.Vopenid=Vopenid
	s_PlayerSystemModuleFlow.Level=Level
	s_PlayerSystemModuleFlow.Typeid=Typeid
	s_PlayerSystemModuleFlow.Vip=Vip
	return s_PlayerSystemModuleFlow
}

// 此结构体由工具按XML描述自动生成
type PlayerSwordDrawFlow struct {
	GameSvrId string // (必填)登录的游戏服务器编号
	DtEventTime time.Time // (必填)游戏事件的时间, 格式 YYYY-MM-DD HH:MM:SS
	VGameAppid string // (必填)游戏APPID
	PlatID int32 // (必填)ios 0/android 1
	Vopenid string // (必填)玩家
	Level int32 // (必填)人物等级
	Vip int32 // (必填)Vip等级
	IMoney int32 // (必填)动作涉及的金钱数
	IMoneyType int32 // (必填)钱的类型MONEYTYPE,其它货币类型参考FAQ文档
	Swordid int32 // 剑心id

}

// Log接口要求的Packet方法
func (log *PlayerSwordDrawFlow) Packet() []byte {
	struct_name := "PlayerSwordDrawFlow"
	buff := make([]byte, 0, 218)
	buff = append(buff, []byte(struct_name)...)
	
	buff = append(buff, '|')
	if len(log.GameSvrId) > 25 {
		buff = append(buff, []byte(log.GameSvrId[:25])...)
	} else {
		buff = append(buff, []byte(log.GameSvrId)...)
	}
	
	buff = append(buff, '|')
	buff = append(buff, []byte(log.DtEventTime.Format(layout))...)
	
	buff = append(buff, '|')
	if len(log.VGameAppid) > 32 {
		buff = append(buff, []byte(log.VGameAppid[:32])...)
	} else {
		buff = append(buff, []byte(log.VGameAppid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.PlatID), 10)
	
	buff = append(buff, '|')
	if len(log.Vopenid) > 64 {
		buff = append(buff, []byte(log.Vopenid[:64])...)
	} else {
		buff = append(buff, []byte(log.Vopenid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Level), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Vip), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.IMoney), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.IMoneyType), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Swordid), 10)
	
	buff = append(buff, '\n')
	return buff[:len(buff)]
}

func (log *PlayerSwordDrawFlow) InvokeHook() {}
func (log *PlayerSwordDrawFlow) CommitToTLog() {}
func (log *PlayerSwordDrawFlow) CommitToXdLog() {}
func (log *PlayerSwordDrawFlow) CommitToMySql() error { return nil }
func (log *PlayerSwordDrawFlow) Rollback() {}
func (log *PlayerSwordDrawFlow) Free() {}
func (log *PlayerSwordDrawFlow) SetEventId(time.Time) {}

func (log *PlayerSwordDrawFlow) GetTransType() int8 {
	return mdb.TRANS_TYPE_TLOG
}

func (log *PlayerSwordDrawFlow) CommitToFile(f *mdb.SyncFile) {
	buff:=log.Packet()
	f.WriteTLog(buff)
	if _, err := SendRaw(buff); err != nil {
		glog.Errorf("[tlog] PlayerSwordDrawFlow error. %v\n info: %v", err, log)
	}
}

// 严格的 NewStucture functions 防止松散的创建方式带来未初始化的字段
func NewPlayerSwordDrawFlow(Vopenid string, Level int32, Vip int32, IMoney int32, IMoneyType int32, Swordid int32) *PlayerSwordDrawFlow {
	s_PlayerSwordDrawFlow:=new(PlayerSwordDrawFlow)
	
	s_PlayerSwordDrawFlow.GameSvrId = g_GameSrvID
	s_PlayerSwordDrawFlow.VGameAppid = g_GameAppID
	s_PlayerSwordDrawFlow.PlatID = g_GamePlatID
	s_PlayerSwordDrawFlow.DtEventTime = time.Now()
	
	
	s_PlayerSwordDrawFlow.Vopenid=Vopenid
	s_PlayerSwordDrawFlow.Level=Level
	s_PlayerSwordDrawFlow.Vip=Vip
	s_PlayerSwordDrawFlow.IMoney=IMoney
	s_PlayerSwordDrawFlow.IMoneyType=IMoneyType
	s_PlayerSwordDrawFlow.Swordid=Swordid
	return s_PlayerSwordDrawFlow
}

// 此结构体由工具按XML描述自动生成
type PlayerChestDrawFlow struct {
	GameSvrId string // (必填)登录的游戏服务器编号
	DtEventTime time.Time // (必填)游戏事件的时间, 格式 YYYY-MM-DD HH:MM:SS
	VGameAppid string // (必填)游戏APPID
	PlatID int32 // (必填)ios 0/android 1
	Vopenid string // (必填)玩家
	Level int32 // (必填)人物等级
	Vip int32 // (必填)Vip等级
	IMoney int32 // (必填)动作涉及的金钱数
	IMoneyType int32 // (必填)钱的类型MONEYTYPE
	ItemList string // (必填)抽取得到的物品列表（数组：id type num 由id 类型和数量组成）
	Chesttype int32 // 宝箱类别

}

// Log接口要求的Packet方法
func (log *PlayerChestDrawFlow) Packet() []byte {
	struct_name := "PlayerChestDrawFlow"
	buff := make([]byte, 0, 475)
	buff = append(buff, []byte(struct_name)...)
	
	buff = append(buff, '|')
	if len(log.GameSvrId) > 25 {
		buff = append(buff, []byte(log.GameSvrId[:25])...)
	} else {
		buff = append(buff, []byte(log.GameSvrId)...)
	}
	
	buff = append(buff, '|')
	buff = append(buff, []byte(log.DtEventTime.Format(layout))...)
	
	buff = append(buff, '|')
	if len(log.VGameAppid) > 32 {
		buff = append(buff, []byte(log.VGameAppid[:32])...)
	} else {
		buff = append(buff, []byte(log.VGameAppid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.PlatID), 10)
	
	buff = append(buff, '|')
	if len(log.Vopenid) > 64 {
		buff = append(buff, []byte(log.Vopenid[:64])...)
	} else {
		buff = append(buff, []byte(log.Vopenid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Level), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Vip), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.IMoney), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.IMoneyType), 10)
	
	buff = append(buff, '|')
	if len(log.ItemList) > 256 {
		buff = append(buff, []byte(log.ItemList[:256])...)
	} else {
		buff = append(buff, []byte(log.ItemList)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Chesttype), 10)
	
	buff = append(buff, '\n')
	return buff[:len(buff)]
}

func (log *PlayerChestDrawFlow) InvokeHook() {}
func (log *PlayerChestDrawFlow) CommitToTLog() {}
func (log *PlayerChestDrawFlow) CommitToXdLog() {}
func (log *PlayerChestDrawFlow) CommitToMySql() error { return nil }
func (log *PlayerChestDrawFlow) Rollback() {}
func (log *PlayerChestDrawFlow) Free() {}
func (log *PlayerChestDrawFlow) SetEventId(time.Time) {}

func (log *PlayerChestDrawFlow) GetTransType() int8 {
	return mdb.TRANS_TYPE_TLOG
}

func (log *PlayerChestDrawFlow) CommitToFile(f *mdb.SyncFile) {
	buff:=log.Packet()
	f.WriteTLog(buff)
	if _, err := SendRaw(buff); err != nil {
		glog.Errorf("[tlog] PlayerChestDrawFlow error. %v\n info: %v", err, log)
	}
}

// 严格的 NewStucture functions 防止松散的创建方式带来未初始化的字段
func NewPlayerChestDrawFlow(Vopenid string, Level int32, Vip int32, IMoney int32, IMoneyType int32, ItemList string, Chesttype int32) *PlayerChestDrawFlow {
	s_PlayerChestDrawFlow:=new(PlayerChestDrawFlow)
	
	s_PlayerChestDrawFlow.GameSvrId = g_GameSrvID
	s_PlayerChestDrawFlow.VGameAppid = g_GameAppID
	s_PlayerChestDrawFlow.PlatID = g_GamePlatID
	s_PlayerChestDrawFlow.DtEventTime = time.Now()
	
	
	s_PlayerChestDrawFlow.Vopenid=Vopenid
	s_PlayerChestDrawFlow.Level=Level
	s_PlayerChestDrawFlow.Vip=Vip
	s_PlayerChestDrawFlow.IMoney=IMoney
	s_PlayerChestDrawFlow.IMoneyType=IMoneyType
	s_PlayerChestDrawFlow.ItemList=ItemList
	s_PlayerChestDrawFlow.Chesttype=Chesttype
	return s_PlayerChestDrawFlow
}

// 此结构体由工具按XML描述自动生成
type PlayerFightNumFlow struct {
	GameSvrId string // (必填)登录的游戏服务器编号
	DtEventTime time.Time // (必填)游戏事件的时间, 格式 YYYY-MM-DD HH:MM:SS
	VGameAppid string // (必填)游戏APPID
	PlatID int32 // (必填)ios 0/android 1
	Vopenid string // (必填)玩家
	Level int32 // (必填)人物等级
	Vip int32 // (必填)Vip等级
	Fightnum int32 // (必填)人物战力

}

// Log接口要求的Packet方法
func (log *PlayerFightNumFlow) Packet() []byte {
	struct_name := "PlayerFightNumFlow"
	buff := make([]byte, 0, 199)
	buff = append(buff, []byte(struct_name)...)
	
	buff = append(buff, '|')
	if len(log.GameSvrId) > 25 {
		buff = append(buff, []byte(log.GameSvrId[:25])...)
	} else {
		buff = append(buff, []byte(log.GameSvrId)...)
	}
	
	buff = append(buff, '|')
	buff = append(buff, []byte(log.DtEventTime.Format(layout))...)
	
	buff = append(buff, '|')
	if len(log.VGameAppid) > 32 {
		buff = append(buff, []byte(log.VGameAppid[:32])...)
	} else {
		buff = append(buff, []byte(log.VGameAppid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.PlatID), 10)
	
	buff = append(buff, '|')
	if len(log.Vopenid) > 64 {
		buff = append(buff, []byte(log.Vopenid[:64])...)
	} else {
		buff = append(buff, []byte(log.Vopenid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Level), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Vip), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Fightnum), 10)
	
	buff = append(buff, '\n')
	return buff[:len(buff)]
}

func (log *PlayerFightNumFlow) InvokeHook() {}
func (log *PlayerFightNumFlow) CommitToTLog() {}
func (log *PlayerFightNumFlow) CommitToXdLog() {}
func (log *PlayerFightNumFlow) CommitToMySql() error { return nil }
func (log *PlayerFightNumFlow) Rollback() {}
func (log *PlayerFightNumFlow) Free() {}
func (log *PlayerFightNumFlow) SetEventId(time.Time) {}

func (log *PlayerFightNumFlow) GetTransType() int8 {
	return mdb.TRANS_TYPE_TLOG
}

func (log *PlayerFightNumFlow) CommitToFile(f *mdb.SyncFile) {
	buff:=log.Packet()
	f.WriteTLog(buff)
	if _, err := SendRaw(buff); err != nil {
		glog.Errorf("[tlog] PlayerFightNumFlow error. %v\n info: %v", err, log)
	}
}

// 严格的 NewStucture functions 防止松散的创建方式带来未初始化的字段
func NewPlayerFightNumFlow(Vopenid string, Level int32, Vip int32, Fightnum int32) *PlayerFightNumFlow {
	s_PlayerFightNumFlow:=new(PlayerFightNumFlow)
	
	s_PlayerFightNumFlow.GameSvrId = g_GameSrvID
	s_PlayerFightNumFlow.VGameAppid = g_GameAppID
	s_PlayerFightNumFlow.PlatID = g_GamePlatID
	s_PlayerFightNumFlow.DtEventTime = time.Now()
	
	
	s_PlayerFightNumFlow.Vopenid=Vopenid
	s_PlayerFightNumFlow.Level=Level
	s_PlayerFightNumFlow.Vip=Vip
	s_PlayerFightNumFlow.Fightnum=Fightnum
	return s_PlayerFightNumFlow
}

// 此结构体由工具按XML描述自动生成
type PlayerQuestFlow struct {
	GameSvrId string // (必填)登录的游戏服务器编号
	DtEventTime time.Time // (必填)游戏事件的时间, 格式 YYYY-MM-DD HH:MM:SS
	VGameAppid string // (必填)游戏APPID
	PlatID int32 // (必填)ios 0/android 1
	Vopenid string // (必填)玩家
	Level int32 // (必填)人物等级
	Vip int32 // (必填)Vip等级
	Questid int32 // (必填)任务id
	QuestType int32 // (必填)任务类型
	Questaction int32 // (必填)任务完成情况(0接收/1完成)

}

// Log接口要求的Packet方法
func (log *PlayerQuestFlow) Packet() []byte {
	struct_name := "PlayerQuestFlow"
	buff := make([]byte, 0, 214)
	buff = append(buff, []byte(struct_name)...)
	
	buff = append(buff, '|')
	if len(log.GameSvrId) > 25 {
		buff = append(buff, []byte(log.GameSvrId[:25])...)
	} else {
		buff = append(buff, []byte(log.GameSvrId)...)
	}
	
	buff = append(buff, '|')
	buff = append(buff, []byte(log.DtEventTime.Format(layout))...)
	
	buff = append(buff, '|')
	if len(log.VGameAppid) > 32 {
		buff = append(buff, []byte(log.VGameAppid[:32])...)
	} else {
		buff = append(buff, []byte(log.VGameAppid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.PlatID), 10)
	
	buff = append(buff, '|')
	if len(log.Vopenid) > 64 {
		buff = append(buff, []byte(log.Vopenid[:64])...)
	} else {
		buff = append(buff, []byte(log.Vopenid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Level), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Vip), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Questid), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.QuestType), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Questaction), 10)
	
	buff = append(buff, '\n')
	return buff[:len(buff)]
}

func (log *PlayerQuestFlow) InvokeHook() {}
func (log *PlayerQuestFlow) CommitToTLog() {}
func (log *PlayerQuestFlow) CommitToXdLog() {}
func (log *PlayerQuestFlow) CommitToMySql() error { return nil }
func (log *PlayerQuestFlow) Rollback() {}
func (log *PlayerQuestFlow) Free() {}
func (log *PlayerQuestFlow) SetEventId(time.Time) {}

func (log *PlayerQuestFlow) GetTransType() int8 {
	return mdb.TRANS_TYPE_TLOG
}

func (log *PlayerQuestFlow) CommitToFile(f *mdb.SyncFile) {
	buff:=log.Packet()
	f.WriteTLog(buff)
	if _, err := SendRaw(buff); err != nil {
		glog.Errorf("[tlog] PlayerQuestFlow error. %v\n info: %v", err, log)
	}
}

// 严格的 NewStucture functions 防止松散的创建方式带来未初始化的字段
func NewPlayerQuestFlow(Vopenid string, Level int32, Vip int32, Questid int32, QuestType int32, Questaction int32) *PlayerQuestFlow {
	s_PlayerQuestFlow:=new(PlayerQuestFlow)
	
	s_PlayerQuestFlow.GameSvrId = g_GameSrvID
	s_PlayerQuestFlow.VGameAppid = g_GameAppID
	s_PlayerQuestFlow.PlatID = g_GamePlatID
	s_PlayerQuestFlow.DtEventTime = time.Now()
	
	
	s_PlayerQuestFlow.Vopenid=Vopenid
	s_PlayerQuestFlow.Level=Level
	s_PlayerQuestFlow.Vip=Vip
	s_PlayerQuestFlow.Questid=Questid
	s_PlayerQuestFlow.QuestType=QuestType
	s_PlayerQuestFlow.Questaction=Questaction
	return s_PlayerQuestFlow
}

// 此结构体由工具按XML描述自动生成
type BusinessManFlow struct {
	GameSvrId string // (必填)登录的游戏服务器编号
	DtEventTime time.Time // (必填)游戏事件的时间, 格式 YYYY-MM-DD HH:MM:SS
	VGameAppid string // (必填)游戏APPID
	PlatID int32 // (必填)ios 0/android 1
	Vopenid string // (必填)玩家
	IGoodsId int32 // (必填)道具ID
	Money int32 // (必填)涉及的铜币数
	Count int32 // (必填)涉及的物品数
	Level int32 // (必填)人物等级
	AddOrReduce int32 // (必填)增加 0/减少 1 针对物品
	Vip int32 // (必填)Vip等级

}

// Log接口要求的Packet方法
func (log *BusinessManFlow) Packet() []byte {
	struct_name := "BusinessManFlow"
	buff := make([]byte, 0, 223)
	buff = append(buff, []byte(struct_name)...)
	
	buff = append(buff, '|')
	if len(log.GameSvrId) > 25 {
		buff = append(buff, []byte(log.GameSvrId[:25])...)
	} else {
		buff = append(buff, []byte(log.GameSvrId)...)
	}
	
	buff = append(buff, '|')
	buff = append(buff, []byte(log.DtEventTime.Format(layout))...)
	
	buff = append(buff, '|')
	if len(log.VGameAppid) > 32 {
		buff = append(buff, []byte(log.VGameAppid[:32])...)
	} else {
		buff = append(buff, []byte(log.VGameAppid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.PlatID), 10)
	
	buff = append(buff, '|')
	if len(log.Vopenid) > 64 {
		buff = append(buff, []byte(log.Vopenid[:64])...)
	} else {
		buff = append(buff, []byte(log.Vopenid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.IGoodsId), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Money), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Count), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Level), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.AddOrReduce), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Vip), 10)
	
	buff = append(buff, '\n')
	return buff[:len(buff)]
}

func (log *BusinessManFlow) InvokeHook() {}
func (log *BusinessManFlow) CommitToTLog() {}
func (log *BusinessManFlow) CommitToXdLog() {}
func (log *BusinessManFlow) CommitToMySql() error { return nil }
func (log *BusinessManFlow) Rollback() {}
func (log *BusinessManFlow) Free() {}
func (log *BusinessManFlow) SetEventId(time.Time) {}

func (log *BusinessManFlow) GetTransType() int8 {
	return mdb.TRANS_TYPE_TLOG
}

func (log *BusinessManFlow) CommitToFile(f *mdb.SyncFile) {
	buff:=log.Packet()
	f.WriteTLog(buff)
	if _, err := SendRaw(buff); err != nil {
		glog.Errorf("[tlog] BusinessManFlow error. %v\n info: %v", err, log)
	}
}

// 严格的 NewStucture functions 防止松散的创建方式带来未初始化的字段
func NewBusinessManFlow(Vopenid string, IGoodsId int32, Money int32, Count int32, Level int32, AddOrReduce int32, Vip int32) *BusinessManFlow {
	s_BusinessManFlow:=new(BusinessManFlow)
	
	s_BusinessManFlow.GameSvrId = g_GameSrvID
	s_BusinessManFlow.VGameAppid = g_GameAppID
	s_BusinessManFlow.PlatID = g_GamePlatID
	s_BusinessManFlow.DtEventTime = time.Now()
	
	
	s_BusinessManFlow.Vopenid=Vopenid
	s_BusinessManFlow.IGoodsId=IGoodsId
	s_BusinessManFlow.Money=Money
	s_BusinessManFlow.Count=Count
	s_BusinessManFlow.Level=Level
	s_BusinessManFlow.AddOrReduce=AddOrReduce
	s_BusinessManFlow.Vip=Vip
	return s_BusinessManFlow
}

// 此结构体由工具按XML描述自动生成
type TradeBuyFlow struct {
	GameSvrId string // (必填)登录的游戏服务器编号
	DtEventTime time.Time // (必填)游戏事件的时间, 格式 YYYY-MM-DD HH:MM:SS
	VGameAppid string // (必填)游戏APPID
	PlatID int32 // (必填)ios 0/android 1
	Vopenid string // (必填)玩家
	IGoodsId int32 // (必填)道具ID
	Count int32 // (必填)涉及的物品数
	Money int32 // (必填)涉及的铜币数
	IMoneyType int32 // (必填)钱的类型MONEYTYPE,其它货币类型参考FAQ文档
	Level int32 // (必填)人物等级
	Vip int32 // (必填)Vip等级

}

// Log接口要求的Packet方法
func (log *TradeBuyFlow) Packet() []byte {
	struct_name := "TradeBuyFlow"
	buff := make([]byte, 0, 220)
	buff = append(buff, []byte(struct_name)...)
	
	buff = append(buff, '|')
	if len(log.GameSvrId) > 25 {
		buff = append(buff, []byte(log.GameSvrId[:25])...)
	} else {
		buff = append(buff, []byte(log.GameSvrId)...)
	}
	
	buff = append(buff, '|')
	buff = append(buff, []byte(log.DtEventTime.Format(layout))...)
	
	buff = append(buff, '|')
	if len(log.VGameAppid) > 32 {
		buff = append(buff, []byte(log.VGameAppid[:32])...)
	} else {
		buff = append(buff, []byte(log.VGameAppid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.PlatID), 10)
	
	buff = append(buff, '|')
	if len(log.Vopenid) > 64 {
		buff = append(buff, []byte(log.Vopenid[:64])...)
	} else {
		buff = append(buff, []byte(log.Vopenid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.IGoodsId), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Count), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Money), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.IMoneyType), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Level), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Vip), 10)
	
	buff = append(buff, '\n')
	return buff[:len(buff)]
}

func (log *TradeBuyFlow) InvokeHook() {}
func (log *TradeBuyFlow) CommitToTLog() {}
func (log *TradeBuyFlow) CommitToXdLog() {}
func (log *TradeBuyFlow) CommitToMySql() error { return nil }
func (log *TradeBuyFlow) Rollback() {}
func (log *TradeBuyFlow) Free() {}
func (log *TradeBuyFlow) SetEventId(time.Time) {}

func (log *TradeBuyFlow) GetTransType() int8 {
	return mdb.TRANS_TYPE_TLOG
}

func (log *TradeBuyFlow) CommitToFile(f *mdb.SyncFile) {
	buff:=log.Packet()
	f.WriteTLog(buff)
	if _, err := SendRaw(buff); err != nil {
		glog.Errorf("[tlog] TradeBuyFlow error. %v\n info: %v", err, log)
	}
}

// 严格的 NewStucture functions 防止松散的创建方式带来未初始化的字段
func NewTradeBuyFlow(Vopenid string, IGoodsId int32, Count int32, Money int32, IMoneyType int32, Level int32, Vip int32) *TradeBuyFlow {
	s_TradeBuyFlow:=new(TradeBuyFlow)
	
	s_TradeBuyFlow.GameSvrId = g_GameSrvID
	s_TradeBuyFlow.VGameAppid = g_GameAppID
	s_TradeBuyFlow.PlatID = g_GamePlatID
	s_TradeBuyFlow.DtEventTime = time.Now()
	
	
	s_TradeBuyFlow.Vopenid=Vopenid
	s_TradeBuyFlow.IGoodsId=IGoodsId
	s_TradeBuyFlow.Count=Count
	s_TradeBuyFlow.Money=Money
	s_TradeBuyFlow.IMoneyType=IMoneyType
	s_TradeBuyFlow.Level=Level
	s_TradeBuyFlow.Vip=Vip
	return s_TradeBuyFlow
}

// 此结构体由工具按XML描述自动生成
type EquipRefineFlow struct {
	GameSvrId string // (必填)登录的游戏服务器编号
	DtEventTime time.Time // (必填)游戏事件的时间, 格式 YYYY-MM-DD HH:MM:SS
	VGameAppid string // (必填)游戏APPID
	PlatID int32 // (必填)ios 0/android 1
	Vopenid string // (必填)玩家
	Itemid int32 // (必填)物品id
	Bluecount int32 // (必填)原有蓝色结晶数量
	Blueaftercount int32 // (必填)精炼后蓝色结晶数量
	Purplecount int32 // (必填)原有紫色结晶数量
	Purpleaftercount int32 // (必填)精炼后紫色结晶数量
	Goldcount int32 // (必填)原有金色结晶数量
	Goldaftercount int32 // (必填)精炼后金色结晶数量
	Orangecount int32 // (必填)原有橙色结晶数量
	Orangeaftercount int32 // (必填)精炼后橙色结晶数量
	Fragmentid int32 // (必填)消耗碎片Id
	Fragmentcount int32 // (必填)原有碎片数量
	Fragmentaftercount int32 // (必填)精炼后碎片数量
	IMoneyType int32 // (必填)钱的类型MONEYTYPE,其它货币类型参考FAQ文档
	Imoney int32 // (必填)原有货币数量
	Afterimoney int32 // (必填)动作后的货币数量
	Refinelevel int32 // (必填)动作前装备强化等级
	Afterrefinelevel int32 // (必填)动作后装备强化等级
	Level int32 // (必填)人物等级
	Vip int32 // (必填)Vip等级

}

// Log接口要求的Packet方法
func (log *EquipRefineFlow) Packet() []byte {
	struct_name := "EquipRefineFlow"
	buff := make([]byte, 0, 340)
	buff = append(buff, []byte(struct_name)...)
	
	buff = append(buff, '|')
	if len(log.GameSvrId) > 25 {
		buff = append(buff, []byte(log.GameSvrId[:25])...)
	} else {
		buff = append(buff, []byte(log.GameSvrId)...)
	}
	
	buff = append(buff, '|')
	buff = append(buff, []byte(log.DtEventTime.Format(layout))...)
	
	buff = append(buff, '|')
	if len(log.VGameAppid) > 32 {
		buff = append(buff, []byte(log.VGameAppid[:32])...)
	} else {
		buff = append(buff, []byte(log.VGameAppid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.PlatID), 10)
	
	buff = append(buff, '|')
	if len(log.Vopenid) > 64 {
		buff = append(buff, []byte(log.Vopenid[:64])...)
	} else {
		buff = append(buff, []byte(log.Vopenid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Itemid), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Bluecount), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Blueaftercount), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Purplecount), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Purpleaftercount), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Goldcount), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Goldaftercount), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Orangecount), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Orangeaftercount), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Fragmentid), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Fragmentcount), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Fragmentaftercount), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.IMoneyType), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Imoney), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Afterimoney), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Refinelevel), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Afterrefinelevel), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Level), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Vip), 10)
	
	buff = append(buff, '\n')
	return buff[:len(buff)]
}

func (log *EquipRefineFlow) InvokeHook() {}
func (log *EquipRefineFlow) CommitToTLog() {}
func (log *EquipRefineFlow) CommitToXdLog() {}
func (log *EquipRefineFlow) CommitToMySql() error { return nil }
func (log *EquipRefineFlow) Rollback() {}
func (log *EquipRefineFlow) Free() {}
func (log *EquipRefineFlow) SetEventId(time.Time) {}

func (log *EquipRefineFlow) GetTransType() int8 {
	return mdb.TRANS_TYPE_TLOG
}

func (log *EquipRefineFlow) CommitToFile(f *mdb.SyncFile) {
	buff:=log.Packet()
	f.WriteTLog(buff)
	if _, err := SendRaw(buff); err != nil {
		glog.Errorf("[tlog] EquipRefineFlow error. %v\n info: %v", err, log)
	}
}

// 严格的 NewStucture functions 防止松散的创建方式带来未初始化的字段
func NewEquipRefineFlow(Vopenid string, Itemid int32, Bluecount int32, Blueaftercount int32, Purplecount int32, Purpleaftercount int32, Goldcount int32, Goldaftercount int32, Orangecount int32, Orangeaftercount int32, Fragmentid int32, Fragmentcount int32, Fragmentaftercount int32, IMoneyType int32, Imoney int32, Afterimoney int32, Refinelevel int32, Afterrefinelevel int32, Level int32, Vip int32) *EquipRefineFlow {
	s_EquipRefineFlow:=new(EquipRefineFlow)
	
	s_EquipRefineFlow.GameSvrId = g_GameSrvID
	s_EquipRefineFlow.VGameAppid = g_GameAppID
	s_EquipRefineFlow.PlatID = g_GamePlatID
	s_EquipRefineFlow.DtEventTime = time.Now()
	
	
	s_EquipRefineFlow.Vopenid=Vopenid
	s_EquipRefineFlow.Itemid=Itemid
	s_EquipRefineFlow.Bluecount=Bluecount
	s_EquipRefineFlow.Blueaftercount=Blueaftercount
	s_EquipRefineFlow.Purplecount=Purplecount
	s_EquipRefineFlow.Purpleaftercount=Purpleaftercount
	s_EquipRefineFlow.Goldcount=Goldcount
	s_EquipRefineFlow.Goldaftercount=Goldaftercount
	s_EquipRefineFlow.Orangecount=Orangecount
	s_EquipRefineFlow.Orangeaftercount=Orangeaftercount
	s_EquipRefineFlow.Fragmentid=Fragmentid
	s_EquipRefineFlow.Fragmentcount=Fragmentcount
	s_EquipRefineFlow.Fragmentaftercount=Fragmentaftercount
	s_EquipRefineFlow.IMoneyType=IMoneyType
	s_EquipRefineFlow.Imoney=Imoney
	s_EquipRefineFlow.Afterimoney=Afterimoney
	s_EquipRefineFlow.Refinelevel=Refinelevel
	s_EquipRefineFlow.Afterrefinelevel=Afterrefinelevel
	s_EquipRefineFlow.Level=Level
	s_EquipRefineFlow.Vip=Vip
	return s_EquipRefineFlow
}

// 此结构体由工具按XML描述自动生成
type EquipDecomposeFlow struct {
	GameSvrId string // (必填)登录的游戏服务器编号
	DtEventTime time.Time // (必填)游戏事件的时间, 格式 YYYY-MM-DD HH:MM:SS
	VGameAppid string // (必填)游戏APPID
	PlatID int32 // (必填)ios 0/android 1
	Vopenid string // (必填)玩家
	Itemid int32 // (必填)物品id
	Crystaid int32 // (必填)结晶id
	Crystacount int32 // (必填)原有结晶数量
	Crystaaftercount int32 // (必填)分解后结晶数量
	Fragmentid int32 // (必填)碎片id
	Fragmentcount int32 // (必填)原有碎片数量
	Fragmentaftercount int32 // (必填)分解后碎片数量
	Level int32 // (必填)人物等级
	Vip int32 // (必填)Vip等级

}

// Log接口要求的Packet方法
func (log *EquipDecomposeFlow) Packet() []byte {
	struct_name := "EquipDecomposeFlow"
	buff := make([]byte, 0, 253)
	buff = append(buff, []byte(struct_name)...)
	
	buff = append(buff, '|')
	if len(log.GameSvrId) > 25 {
		buff = append(buff, []byte(log.GameSvrId[:25])...)
	} else {
		buff = append(buff, []byte(log.GameSvrId)...)
	}
	
	buff = append(buff, '|')
	buff = append(buff, []byte(log.DtEventTime.Format(layout))...)
	
	buff = append(buff, '|')
	if len(log.VGameAppid) > 32 {
		buff = append(buff, []byte(log.VGameAppid[:32])...)
	} else {
		buff = append(buff, []byte(log.VGameAppid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.PlatID), 10)
	
	buff = append(buff, '|')
	if len(log.Vopenid) > 64 {
		buff = append(buff, []byte(log.Vopenid[:64])...)
	} else {
		buff = append(buff, []byte(log.Vopenid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Itemid), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Crystaid), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Crystacount), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Crystaaftercount), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Fragmentid), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Fragmentcount), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Fragmentaftercount), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Level), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Vip), 10)
	
	buff = append(buff, '\n')
	return buff[:len(buff)]
}

func (log *EquipDecomposeFlow) InvokeHook() {}
func (log *EquipDecomposeFlow) CommitToTLog() {}
func (log *EquipDecomposeFlow) CommitToXdLog() {}
func (log *EquipDecomposeFlow) CommitToMySql() error { return nil }
func (log *EquipDecomposeFlow) Rollback() {}
func (log *EquipDecomposeFlow) Free() {}
func (log *EquipDecomposeFlow) SetEventId(time.Time) {}

func (log *EquipDecomposeFlow) GetTransType() int8 {
	return mdb.TRANS_TYPE_TLOG
}

func (log *EquipDecomposeFlow) CommitToFile(f *mdb.SyncFile) {
	buff:=log.Packet()
	f.WriteTLog(buff)
	if _, err := SendRaw(buff); err != nil {
		glog.Errorf("[tlog] EquipDecomposeFlow error. %v\n info: %v", err, log)
	}
}

// 严格的 NewStucture functions 防止松散的创建方式带来未初始化的字段
func NewEquipDecomposeFlow(Vopenid string, Itemid int32, Crystaid int32, Crystacount int32, Crystaaftercount int32, Fragmentid int32, Fragmentcount int32, Fragmentaftercount int32, Level int32, Vip int32) *EquipDecomposeFlow {
	s_EquipDecomposeFlow:=new(EquipDecomposeFlow)
	
	s_EquipDecomposeFlow.GameSvrId = g_GameSrvID
	s_EquipDecomposeFlow.VGameAppid = g_GameAppID
	s_EquipDecomposeFlow.PlatID = g_GamePlatID
	s_EquipDecomposeFlow.DtEventTime = time.Now()
	
	
	s_EquipDecomposeFlow.Vopenid=Vopenid
	s_EquipDecomposeFlow.Itemid=Itemid
	s_EquipDecomposeFlow.Crystaid=Crystaid
	s_EquipDecomposeFlow.Crystacount=Crystacount
	s_EquipDecomposeFlow.Crystaaftercount=Crystaaftercount
	s_EquipDecomposeFlow.Fragmentid=Fragmentid
	s_EquipDecomposeFlow.Fragmentcount=Fragmentcount
	s_EquipDecomposeFlow.Fragmentaftercount=Fragmentaftercount
	s_EquipDecomposeFlow.Level=Level
	s_EquipDecomposeFlow.Vip=Vip
	return s_EquipDecomposeFlow
}

// 此结构体由工具按XML描述自动生成
type GhostTrainFlow struct {
	GameSvrId string // (必填)登录的游戏服务器编号
	DtEventTime time.Time // (必填)游戏事件的时间, 格式 YYYY-MM-DD HH:MM:SS
	VGameAppid string // (必填)游戏APPID
	PlatID int32 // (必填)ios 0/android 1
	Vopenid string // (必填)玩家
	Ghostid int32 // (必填)魂侍id
	Fruitcount int32 // (必填)原有影界果实数量
	Fruitaftercount int32 // (必填)强化后影界果实数量
	IMoneyType int32 // (必填)货币类型
	IMoney int32 // (必填)消耗货币数量
	Ghostlevel int32 // (必填)原有魂侍等级
	Afterghostlevel int32 // (必填)强化后魂侍等级
	Level int32 // (必填)人物等级
	Vip int32 // (必填)Vip等级

}

// Log接口要求的Packet方法
func (log *GhostTrainFlow) Packet() []byte {
	struct_name := "GhostTrainFlow"
	buff := make([]byte, 0, 249)
	buff = append(buff, []byte(struct_name)...)
	
	buff = append(buff, '|')
	if len(log.GameSvrId) > 25 {
		buff = append(buff, []byte(log.GameSvrId[:25])...)
	} else {
		buff = append(buff, []byte(log.GameSvrId)...)
	}
	
	buff = append(buff, '|')
	buff = append(buff, []byte(log.DtEventTime.Format(layout))...)
	
	buff = append(buff, '|')
	if len(log.VGameAppid) > 32 {
		buff = append(buff, []byte(log.VGameAppid[:32])...)
	} else {
		buff = append(buff, []byte(log.VGameAppid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.PlatID), 10)
	
	buff = append(buff, '|')
	if len(log.Vopenid) > 64 {
		buff = append(buff, []byte(log.Vopenid[:64])...)
	} else {
		buff = append(buff, []byte(log.Vopenid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Ghostid), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Fruitcount), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Fruitaftercount), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.IMoneyType), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.IMoney), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Ghostlevel), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Afterghostlevel), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Level), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Vip), 10)
	
	buff = append(buff, '\n')
	return buff[:len(buff)]
}

func (log *GhostTrainFlow) InvokeHook() {}
func (log *GhostTrainFlow) CommitToTLog() {}
func (log *GhostTrainFlow) CommitToXdLog() {}
func (log *GhostTrainFlow) CommitToMySql() error { return nil }
func (log *GhostTrainFlow) Rollback() {}
func (log *GhostTrainFlow) Free() {}
func (log *GhostTrainFlow) SetEventId(time.Time) {}

func (log *GhostTrainFlow) GetTransType() int8 {
	return mdb.TRANS_TYPE_TLOG
}

func (log *GhostTrainFlow) CommitToFile(f *mdb.SyncFile) {
	buff:=log.Packet()
	f.WriteTLog(buff)
	if _, err := SendRaw(buff); err != nil {
		glog.Errorf("[tlog] GhostTrainFlow error. %v\n info: %v", err, log)
	}
}

// 严格的 NewStucture functions 防止松散的创建方式带来未初始化的字段
func NewGhostTrainFlow(Vopenid string, Ghostid int32, Fruitcount int32, Fruitaftercount int32, IMoneyType int32, IMoney int32, Ghostlevel int32, Afterghostlevel int32, Level int32, Vip int32) *GhostTrainFlow {
	s_GhostTrainFlow:=new(GhostTrainFlow)
	
	s_GhostTrainFlow.GameSvrId = g_GameSrvID
	s_GhostTrainFlow.VGameAppid = g_GameAppID
	s_GhostTrainFlow.PlatID = g_GamePlatID
	s_GhostTrainFlow.DtEventTime = time.Now()
	
	
	s_GhostTrainFlow.Vopenid=Vopenid
	s_GhostTrainFlow.Ghostid=Ghostid
	s_GhostTrainFlow.Fruitcount=Fruitcount
	s_GhostTrainFlow.Fruitaftercount=Fruitaftercount
	s_GhostTrainFlow.IMoneyType=IMoneyType
	s_GhostTrainFlow.IMoney=IMoney
	s_GhostTrainFlow.Ghostlevel=Ghostlevel
	s_GhostTrainFlow.Afterghostlevel=Afterghostlevel
	s_GhostTrainFlow.Level=Level
	s_GhostTrainFlow.Vip=Vip
	return s_GhostTrainFlow
}

// 此结构体由工具按XML描述自动生成
type GhostUpgradeFlow struct {
	GameSvrId string // (必填)登录的游戏服务器编号
	DtEventTime time.Time // (必填)游戏事件的时间, 格式 YYYY-MM-DD HH:MM:SS
	VGameAppid string // (必填)游戏APPID
	PlatID int32 // (必填)ios 0/android 1
	Vopenid string // (必填)玩家
	Ghostid int32 // (必填)魂侍id
	Fragmentid int32 // (必填)魂侍碎片id
	Fragmentcount int32 // (必填)原有魂侍碎片数量
	Fragmentaftercount int32 // (必填)升星后魂侍碎片数量
	IMoneyType int32 // (必填)货币类型
	IMoney int32 // (必填)消耗货币数量
	Ghoststar int32 // (必填)原有魂侍星级
	Afterghoststar int32 // (必填)升星后魂侍星级
	Level int32 // (必填)人物等级
	Vip int32 // (必填)Vip等级

}

// Log接口要求的Packet方法
func (log *GhostUpgradeFlow) Packet() []byte {
	struct_name := "GhostUpgradeFlow"
	buff := make([]byte, 0, 260)
	buff = append(buff, []byte(struct_name)...)
	
	buff = append(buff, '|')
	if len(log.GameSvrId) > 25 {
		buff = append(buff, []byte(log.GameSvrId[:25])...)
	} else {
		buff = append(buff, []byte(log.GameSvrId)...)
	}
	
	buff = append(buff, '|')
	buff = append(buff, []byte(log.DtEventTime.Format(layout))...)
	
	buff = append(buff, '|')
	if len(log.VGameAppid) > 32 {
		buff = append(buff, []byte(log.VGameAppid[:32])...)
	} else {
		buff = append(buff, []byte(log.VGameAppid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.PlatID), 10)
	
	buff = append(buff, '|')
	if len(log.Vopenid) > 64 {
		buff = append(buff, []byte(log.Vopenid[:64])...)
	} else {
		buff = append(buff, []byte(log.Vopenid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Ghostid), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Fragmentid), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Fragmentcount), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Fragmentaftercount), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.IMoneyType), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.IMoney), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Ghoststar), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Afterghoststar), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Level), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Vip), 10)
	
	buff = append(buff, '\n')
	return buff[:len(buff)]
}

func (log *GhostUpgradeFlow) InvokeHook() {}
func (log *GhostUpgradeFlow) CommitToTLog() {}
func (log *GhostUpgradeFlow) CommitToXdLog() {}
func (log *GhostUpgradeFlow) CommitToMySql() error { return nil }
func (log *GhostUpgradeFlow) Rollback() {}
func (log *GhostUpgradeFlow) Free() {}
func (log *GhostUpgradeFlow) SetEventId(time.Time) {}

func (log *GhostUpgradeFlow) GetTransType() int8 {
	return mdb.TRANS_TYPE_TLOG
}

func (log *GhostUpgradeFlow) CommitToFile(f *mdb.SyncFile) {
	buff:=log.Packet()
	f.WriteTLog(buff)
	if _, err := SendRaw(buff); err != nil {
		glog.Errorf("[tlog] GhostUpgradeFlow error. %v\n info: %v", err, log)
	}
}

// 严格的 NewStucture functions 防止松散的创建方式带来未初始化的字段
func NewGhostUpgradeFlow(Vopenid string, Ghostid int32, Fragmentid int32, Fragmentcount int32, Fragmentaftercount int32, IMoneyType int32, IMoney int32, Ghoststar int32, Afterghoststar int32, Level int32, Vip int32) *GhostUpgradeFlow {
	s_GhostUpgradeFlow:=new(GhostUpgradeFlow)
	
	s_GhostUpgradeFlow.GameSvrId = g_GameSrvID
	s_GhostUpgradeFlow.VGameAppid = g_GameAppID
	s_GhostUpgradeFlow.PlatID = g_GamePlatID
	s_GhostUpgradeFlow.DtEventTime = time.Now()
	
	
	s_GhostUpgradeFlow.Vopenid=Vopenid
	s_GhostUpgradeFlow.Ghostid=Ghostid
	s_GhostUpgradeFlow.Fragmentid=Fragmentid
	s_GhostUpgradeFlow.Fragmentcount=Fragmentcount
	s_GhostUpgradeFlow.Fragmentaftercount=Fragmentaftercount
	s_GhostUpgradeFlow.IMoneyType=IMoneyType
	s_GhostUpgradeFlow.IMoney=IMoney
	s_GhostUpgradeFlow.Ghoststar=Ghoststar
	s_GhostUpgradeFlow.Afterghoststar=Afterghoststar
	s_GhostUpgradeFlow.Level=Level
	s_GhostUpgradeFlow.Vip=Vip
	return s_GhostUpgradeFlow
}

// 此结构体由工具按XML描述自动生成
type GhostSwordUpgradeFlow struct {
	GameSvrId string // (必填)登录的游戏服务器编号
	DtEventTime time.Time // (必填)游戏事件的时间, 格式 YYYY-MM-DD HH:MM:SS
	VGameAppid string // (必填)游戏APPID
	PlatID int32 // (必填)ios 0/android 1
	Vopenid string // (必填)玩家
	Swordid int32 // (必填)剑心id
	Swordlist string // (必填)消耗剑心列表|id 数量 id 数量|这样的数字
	Swordlevel int32 // (必填)原有剑心等级
	Afterswordlevel int32 // (必填)强化后剑心等级
	Level int32 // (必填)人物等级
	Vip int32 // (必填)Vip等级

}

// Log接口要求的Packet方法
func (log *GhostSwordUpgradeFlow) Packet() []byte {
	struct_name := "GhostSwordUpgradeFlow"
	buff := make([]byte, 0, 477)
	buff = append(buff, []byte(struct_name)...)
	
	buff = append(buff, '|')
	if len(log.GameSvrId) > 25 {
		buff = append(buff, []byte(log.GameSvrId[:25])...)
	} else {
		buff = append(buff, []byte(log.GameSvrId)...)
	}
	
	buff = append(buff, '|')
	buff = append(buff, []byte(log.DtEventTime.Format(layout))...)
	
	buff = append(buff, '|')
	if len(log.VGameAppid) > 32 {
		buff = append(buff, []byte(log.VGameAppid[:32])...)
	} else {
		buff = append(buff, []byte(log.VGameAppid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.PlatID), 10)
	
	buff = append(buff, '|')
	if len(log.Vopenid) > 64 {
		buff = append(buff, []byte(log.Vopenid[:64])...)
	} else {
		buff = append(buff, []byte(log.Vopenid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Swordid), 10)
	
	buff = append(buff, '|')
	if len(log.Swordlist) > 256 {
		buff = append(buff, []byte(log.Swordlist[:256])...)
	} else {
		buff = append(buff, []byte(log.Swordlist)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Swordlevel), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Afterswordlevel), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Level), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Vip), 10)
	
	buff = append(buff, '\n')
	return buff[:len(buff)]
}

func (log *GhostSwordUpgradeFlow) InvokeHook() {}
func (log *GhostSwordUpgradeFlow) CommitToTLog() {}
func (log *GhostSwordUpgradeFlow) CommitToXdLog() {}
func (log *GhostSwordUpgradeFlow) CommitToMySql() error { return nil }
func (log *GhostSwordUpgradeFlow) Rollback() {}
func (log *GhostSwordUpgradeFlow) Free() {}
func (log *GhostSwordUpgradeFlow) SetEventId(time.Time) {}

func (log *GhostSwordUpgradeFlow) GetTransType() int8 {
	return mdb.TRANS_TYPE_TLOG
}

func (log *GhostSwordUpgradeFlow) CommitToFile(f *mdb.SyncFile) {
	buff:=log.Packet()
	f.WriteTLog(buff)
	if _, err := SendRaw(buff); err != nil {
		glog.Errorf("[tlog] GhostSwordUpgradeFlow error. %v\n info: %v", err, log)
	}
}

// 严格的 NewStucture functions 防止松散的创建方式带来未初始化的字段
func NewGhostSwordUpgradeFlow(Vopenid string, Swordid int32, Swordlist string, Swordlevel int32, Afterswordlevel int32, Level int32, Vip int32) *GhostSwordUpgradeFlow {
	s_GhostSwordUpgradeFlow:=new(GhostSwordUpgradeFlow)
	
	s_GhostSwordUpgradeFlow.GameSvrId = g_GameSrvID
	s_GhostSwordUpgradeFlow.VGameAppid = g_GameAppID
	s_GhostSwordUpgradeFlow.PlatID = g_GamePlatID
	s_GhostSwordUpgradeFlow.DtEventTime = time.Now()
	
	
	s_GhostSwordUpgradeFlow.Vopenid=Vopenid
	s_GhostSwordUpgradeFlow.Swordid=Swordid
	s_GhostSwordUpgradeFlow.Swordlist=Swordlist
	s_GhostSwordUpgradeFlow.Swordlevel=Swordlevel
	s_GhostSwordUpgradeFlow.Afterswordlevel=Afterswordlevel
	s_GhostSwordUpgradeFlow.Level=Level
	s_GhostSwordUpgradeFlow.Vip=Vip
	return s_GhostSwordUpgradeFlow
}

// 此结构体由工具按XML描述自动生成
type SendHeartFlow struct {
	GameSvrId string // (必填)登录的游戏服务器编号
	DtEventTime time.Time // (必填)游戏事件的时间, 格式 YYYY-MM-DD HH:MM:SS
	VGameAppid string // (必填)游戏APPID
	PlatID int32 // (必填)ios 0/android 1
	Vopenid string // (必填)玩家
	FriendID uint64 // (必填)接受爱心的好友ID
	Level int32 // (必填)人物等级
	Vip int32 // (必填)Vip等级

}

// Log接口要求的Packet方法
func (log *SendHeartFlow) Packet() []byte {
	struct_name := "SendHeartFlow"
	buff := make([]byte, 0, 194)
	buff = append(buff, []byte(struct_name)...)
	
	buff = append(buff, '|')
	if len(log.GameSvrId) > 25 {
		buff = append(buff, []byte(log.GameSvrId[:25])...)
	} else {
		buff = append(buff, []byte(log.GameSvrId)...)
	}
	
	buff = append(buff, '|')
	buff = append(buff, []byte(log.DtEventTime.Format(layout))...)
	
	buff = append(buff, '|')
	if len(log.VGameAppid) > 32 {
		buff = append(buff, []byte(log.VGameAppid[:32])...)
	} else {
		buff = append(buff, []byte(log.VGameAppid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.PlatID), 10)
	
	buff = append(buff, '|')
	if len(log.Vopenid) > 64 {
		buff = append(buff, []byte(log.Vopenid[:64])...)
	} else {
		buff = append(buff, []byte(log.Vopenid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendUint(buff, uint64(log.FriendID), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Level), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Vip), 10)
	
	buff = append(buff, '\n')
	return buff[:len(buff)]
}

func (log *SendHeartFlow) InvokeHook() {}
func (log *SendHeartFlow) CommitToTLog() {}
func (log *SendHeartFlow) CommitToXdLog() {}
func (log *SendHeartFlow) CommitToMySql() error { return nil }
func (log *SendHeartFlow) Rollback() {}
func (log *SendHeartFlow) Free() {}
func (log *SendHeartFlow) SetEventId(time.Time) {}

func (log *SendHeartFlow) GetTransType() int8 {
	return mdb.TRANS_TYPE_TLOG
}

func (log *SendHeartFlow) CommitToFile(f *mdb.SyncFile) {
	buff:=log.Packet()
	f.WriteTLog(buff)
	if _, err := SendRaw(buff); err != nil {
		glog.Errorf("[tlog] SendHeartFlow error. %v\n info: %v", err, log)
	}
}

// 严格的 NewStucture functions 防止松散的创建方式带来未初始化的字段
func NewSendHeartFlow(Vopenid string, FriendID uint64, Level int32, Vip int32) *SendHeartFlow {
	s_SendHeartFlow:=new(SendHeartFlow)
	
	s_SendHeartFlow.GameSvrId = g_GameSrvID
	s_SendHeartFlow.VGameAppid = g_GameAppID
	s_SendHeartFlow.PlatID = g_GamePlatID
	s_SendHeartFlow.DtEventTime = time.Now()
	
	
	s_SendHeartFlow.Vopenid=Vopenid
	s_SendHeartFlow.FriendID=FriendID
	s_SendHeartFlow.Level=Level
	s_SendHeartFlow.Vip=Vip
	return s_SendHeartFlow
}

// 此结构体由工具按XML描述自动生成
type AddPetFlow struct {
	GameSvrId string // (必填)登录的游戏服务器编号
	DtEventTime time.Time // (必填)游戏事件的时间, 格式 YYYY-MM-DD HH:MM:SS
	VGameAppid string // (必填)游戏APPID
	PlatID int32 // (必填)ios 0/android 1
	Vopenid string // (必填)玩家
	PetID int32 // (必填)灵宠ID
	RoundAttack int32 // (可选)单回合行动次数
	CostPower int32 // (可选)召唤时消耗精气
	LiveRound int32 // (可选)召唤后存活回合数
	LivePos int32 // (可选)召唤后出现的位置(1-前排；2-后排；3-左侧)
	Quality int32 // (可选)品质
	ParentPetId int32 // (可选)父灵宠ID
	Health int32 // (可选)生命
	Attack int32 // (可选)攻击
	Defence int32 // (可选)防御
	Speed int32 // (可选)速度
	ForceNum int32 // (可选)绝招威力
	Star int32 // (可选)星级
	Level int32 // (必填)人物等级
	Vip int32 // (必填)Vip等级

}

// Log接口要求的Packet方法
func (log *AddPetFlow) Packet() []byte {
	struct_name := "AddPetFlow"
	buff := make([]byte, 0, 299)
	buff = append(buff, []byte(struct_name)...)
	
	buff = append(buff, '|')
	if len(log.GameSvrId) > 25 {
		buff = append(buff, []byte(log.GameSvrId[:25])...)
	} else {
		buff = append(buff, []byte(log.GameSvrId)...)
	}
	
	buff = append(buff, '|')
	buff = append(buff, []byte(log.DtEventTime.Format(layout))...)
	
	buff = append(buff, '|')
	if len(log.VGameAppid) > 32 {
		buff = append(buff, []byte(log.VGameAppid[:32])...)
	} else {
		buff = append(buff, []byte(log.VGameAppid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.PlatID), 10)
	
	buff = append(buff, '|')
	if len(log.Vopenid) > 64 {
		buff = append(buff, []byte(log.Vopenid[:64])...)
	} else {
		buff = append(buff, []byte(log.Vopenid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.PetID), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.RoundAttack), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.CostPower), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.LiveRound), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.LivePos), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Quality), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.ParentPetId), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Health), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Attack), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Defence), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Speed), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.ForceNum), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Star), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Level), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Vip), 10)
	
	buff = append(buff, '\n')
	return buff[:len(buff)]
}

func (log *AddPetFlow) InvokeHook() {}
func (log *AddPetFlow) CommitToTLog() {}
func (log *AddPetFlow) CommitToXdLog() {}
func (log *AddPetFlow) CommitToMySql() error { return nil }
func (log *AddPetFlow) Rollback() {}
func (log *AddPetFlow) Free() {}
func (log *AddPetFlow) SetEventId(time.Time) {}

func (log *AddPetFlow) GetTransType() int8 {
	return mdb.TRANS_TYPE_TLOG
}

func (log *AddPetFlow) CommitToFile(f *mdb.SyncFile) {
	buff:=log.Packet()
	f.WriteTLog(buff)
	if _, err := SendRaw(buff); err != nil {
		glog.Errorf("[tlog] AddPetFlow error. %v\n info: %v", err, log)
	}
}

// 严格的 NewStucture functions 防止松散的创建方式带来未初始化的字段
func NewAddPetFlow(Vopenid string, PetID int32, RoundAttack int32, CostPower int32, LiveRound int32, LivePos int32, Quality int32, ParentPetId int32, Health int32, Attack int32, Defence int32, Speed int32, ForceNum int32, Star int32, Level int32, Vip int32) *AddPetFlow {
	s_AddPetFlow:=new(AddPetFlow)
	
	s_AddPetFlow.GameSvrId = g_GameSrvID
	s_AddPetFlow.VGameAppid = g_GameAppID
	s_AddPetFlow.PlatID = g_GamePlatID
	s_AddPetFlow.DtEventTime = time.Now()
	
	
	s_AddPetFlow.Vopenid=Vopenid
	s_AddPetFlow.PetID=PetID
	s_AddPetFlow.RoundAttack=RoundAttack
	s_AddPetFlow.CostPower=CostPower
	s_AddPetFlow.LiveRound=LiveRound
	s_AddPetFlow.LivePos=LivePos
	s_AddPetFlow.Quality=Quality
	s_AddPetFlow.ParentPetId=ParentPetId
	s_AddPetFlow.Health=Health
	s_AddPetFlow.Attack=Attack
	s_AddPetFlow.Defence=Defence
	s_AddPetFlow.Speed=Speed
	s_AddPetFlow.ForceNum=ForceNum
	s_AddPetFlow.Star=Star
	s_AddPetFlow.Level=Level
	s_AddPetFlow.Vip=Vip
	return s_AddPetFlow
}

// 此结构体由工具按XML描述自动生成
type UpgradePetGridFlow struct {
	GameSvrId string // (必填)登录的游戏服务器编号
	DtEventTime time.Time // (必填)游戏事件的时间, 格式 YYYY-MM-DD HH:MM:SS
	VGameAppid string // (必填)游戏APPID
	PlatID int32 // (必填)ios 0/android 1
	Vopenid string // (必填)玩家
	GridNum int32 // (必填)格子号
	MatrialPreAmount int32 // (必填)升级前灵魄数
	MatrialPostAmount int32 // (必填)升级后灵魄数
	Level int32 // (必填)人物等级
	Vip int32 // (必填)Vip等级

}

// Log接口要求的Packet方法
func (log *UpgradePetGridFlow) Packet() []byte {
	struct_name := "UpgradePetGridFlow"
	buff := make([]byte, 0, 217)
	buff = append(buff, []byte(struct_name)...)
	
	buff = append(buff, '|')
	if len(log.GameSvrId) > 25 {
		buff = append(buff, []byte(log.GameSvrId[:25])...)
	} else {
		buff = append(buff, []byte(log.GameSvrId)...)
	}
	
	buff = append(buff, '|')
	buff = append(buff, []byte(log.DtEventTime.Format(layout))...)
	
	buff = append(buff, '|')
	if len(log.VGameAppid) > 32 {
		buff = append(buff, []byte(log.VGameAppid[:32])...)
	} else {
		buff = append(buff, []byte(log.VGameAppid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.PlatID), 10)
	
	buff = append(buff, '|')
	if len(log.Vopenid) > 64 {
		buff = append(buff, []byte(log.Vopenid[:64])...)
	} else {
		buff = append(buff, []byte(log.Vopenid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.GridNum), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.MatrialPreAmount), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.MatrialPostAmount), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Level), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendInt(buff, int64(log.Vip), 10)
	
	buff = append(buff, '\n')
	return buff[:len(buff)]
}

func (log *UpgradePetGridFlow) InvokeHook() {}
func (log *UpgradePetGridFlow) CommitToTLog() {}
func (log *UpgradePetGridFlow) CommitToXdLog() {}
func (log *UpgradePetGridFlow) CommitToMySql() error { return nil }
func (log *UpgradePetGridFlow) Rollback() {}
func (log *UpgradePetGridFlow) Free() {}
func (log *UpgradePetGridFlow) SetEventId(time.Time) {}

func (log *UpgradePetGridFlow) GetTransType() int8 {
	return mdb.TRANS_TYPE_TLOG
}

func (log *UpgradePetGridFlow) CommitToFile(f *mdb.SyncFile) {
	buff:=log.Packet()
	f.WriteTLog(buff)
	if _, err := SendRaw(buff); err != nil {
		glog.Errorf("[tlog] UpgradePetGridFlow error. %v\n info: %v", err, log)
	}
}

// 严格的 NewStucture functions 防止松散的创建方式带来未初始化的字段
func NewUpgradePetGridFlow(Vopenid string, GridNum int32, MatrialPreAmount int32, MatrialPostAmount int32, Level int32, Vip int32) *UpgradePetGridFlow {
	s_UpgradePetGridFlow:=new(UpgradePetGridFlow)
	
	s_UpgradePetGridFlow.GameSvrId = g_GameSrvID
	s_UpgradePetGridFlow.VGameAppid = g_GameAppID
	s_UpgradePetGridFlow.PlatID = g_GamePlatID
	s_UpgradePetGridFlow.DtEventTime = time.Now()
	
	
	s_UpgradePetGridFlow.Vopenid=Vopenid
	s_UpgradePetGridFlow.GridNum=GridNum
	s_UpgradePetGridFlow.MatrialPreAmount=MatrialPreAmount
	s_UpgradePetGridFlow.MatrialPostAmount=MatrialPostAmount
	s_UpgradePetGridFlow.Level=Level
	s_UpgradePetGridFlow.Vip=Vip
	return s_UpgradePetGridFlow
}

// 此结构体由工具按XML描述自动生成
type IDIPFLOW struct {
	DtEventTime time.Time // (必填)游戏事件的时间, 格式 YYYY-MM-DD HH:MM:SS
	Area_id uint32 // 服务器：微信（1），手Q（2），游客（5）
	Plat_id uint32 // 平台：IOS（0），安卓（1）
	Partition_id uint32 // 小区ID
	Vopenid string // (必填)用户OPENID号，如果是端游改为iuin填QQ号
	Item_id uint32 // (必填)道具id，非道具填0
	Item_num uint32 // (必填)涉及的物品数量
	Item_type uint32 // (必填)涉及的类型(参考IDIPTYPE表)
	Serial string // (必填)流水号前端传入，默认0
	Source uint32 // (必填)渠道id前端传入，默认0
	Cmd uint32 // (必填)接口指令id，需要转为十进制

}

// Log接口要求的Packet方法
func (log *IDIPFLOW) Packet() []byte {
	struct_name := "IDIPFLOW"
	buff := make([]byte, 0, 231)
	buff = append(buff, []byte(struct_name)...)
	
	buff = append(buff, '|')
	buff = append(buff, []byte(log.DtEventTime.Format(layout))...)
	
	buff = append(buff, '|')
	buff = strconv.AppendUint(buff, uint64(log.Area_id), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendUint(buff, uint64(log.Plat_id), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendUint(buff, uint64(log.Partition_id), 10)
	
	buff = append(buff, '|')
	if len(log.Vopenid) > 64 {
		buff = append(buff, []byte(log.Vopenid[:64])...)
	} else {
		buff = append(buff, []byte(log.Vopenid)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendUint(buff, uint64(log.Item_id), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendUint(buff, uint64(log.Item_num), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendUint(buff, uint64(log.Item_type), 10)
	
	buff = append(buff, '|')
	if len(log.Serial) > 64 {
		buff = append(buff, []byte(log.Serial[:64])...)
	} else {
		buff = append(buff, []byte(log.Serial)...)
	}
	
	buff = append(buff, '|')
	buff = strconv.AppendUint(buff, uint64(log.Source), 10)
	
	buff = append(buff, '|')
	buff = strconv.AppendUint(buff, uint64(log.Cmd), 10)
	
	buff = append(buff, '\n')
	return buff[:len(buff)]
}

func (log *IDIPFLOW) InvokeHook() {}
func (log *IDIPFLOW) CommitToTLog() {}
func (log *IDIPFLOW) CommitToXdLog() {}
func (log *IDIPFLOW) CommitToMySql() error { return nil }
func (log *IDIPFLOW) Rollback() {}
func (log *IDIPFLOW) Free() {}
func (log *IDIPFLOW) SetEventId(time.Time) {}

func (log *IDIPFLOW) GetTransType() int8 {
	return mdb.TRANS_TYPE_TLOG
}

func (log *IDIPFLOW) CommitToFile(f *mdb.SyncFile) {
	buff:=log.Packet()
	f.WriteTLog(buff)
	if _, err := SendRaw(buff); err != nil {
		glog.Errorf("[tlog] IDIPFLOW error. %v\n info: %v", err, log)
	}
}

// 严格的 NewStucture functions 防止松散的创建方式带来未初始化的字段
func NewIDIPFLOW(Area_id uint32, Plat_id uint32, Partition_id uint32, Vopenid string, Item_id uint32, Item_num uint32, Item_type uint32, Serial string, Source uint32, Cmd uint32) *IDIPFLOW {
	s_IDIPFLOW:=new(IDIPFLOW)
	
	s_IDIPFLOW.DtEventTime = time.Now()
	
	
	s_IDIPFLOW.Area_id=Area_id
	s_IDIPFLOW.Plat_id=Plat_id
	s_IDIPFLOW.Partition_id=Partition_id
	s_IDIPFLOW.Vopenid=Vopenid
	s_IDIPFLOW.Item_id=Item_id
	s_IDIPFLOW.Item_num=Item_num
	s_IDIPFLOW.Item_type=Item_type
	s_IDIPFLOW.Serial=Serial
	s_IDIPFLOW.Source=Source
	s_IDIPFLOW.Cmd=Cmd
	return s_IDIPFLOW
}

