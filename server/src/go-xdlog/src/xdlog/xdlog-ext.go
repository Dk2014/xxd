// 本文件由 xdlog-parser 根据日志xml规范生成
//请勿手动改写！
package xdlog

import (
	"game_server/mdb"
	"encoding/json"
	gotime "core/time"
	//"strconv"
	"fmt"
	//"core/fail"
	"time"
	//."game_server/config"
)

var layout = "2006-01-02 15:04:05"


// 此结构体由工具按XML描述自动生成
type LoginFlow struct {
	Tag string `json:"tag"`
	Eid uint64 `json:"eid"`
	Microtime string `json:"microtime"`
	Cid	int32 `json:"cid"` // (必填)平台渠道ID
	Sid	int32 `json:"sid"` // (必填)区唯一ID,如果不分区则为0
	Account	string `json:"account"` // (必填)平台帐号
	Pid	int64 `json:"pid"` // (必填)玩家角色ID
	Name	string `json:"name"` // (必填)角色名
	Time	int64 `json:"time"` // (必填)登陆时间戳
	Ip	string `json:"ip"` // (必填)登陆IP
	Type	int32 `json:"type"` // (必填)类型0:登出；1:登入
	Ispay	int32 `json:"ispay"` // (必填)是否付费玩家1:付费;0:非付费
	Level	int32 `json:"level"` // (必填)等级
	Vip	int32 `json:"vip"` // (必填)VIP等级

}

// Log接口要求的Packet方法
func (log *LoginFlow) Packet() []byte {
	log.Tag = "login"
	js, err := json.Marshal(*log)
	if err != nil {
	    panic(err.Error())
	}
	js = append(js, '\n')
	return js
	//return []byte(string(js)+"\n")
}

func (log *LoginFlow) InvokeHook() {}
func (log *LoginFlow) CommitToTLog() {}
func (log *LoginFlow) CommitToXdLog() {}
func (log *LoginFlow) CommitToMySql() error { return nil }
func (log *LoginFlow) Rollback() {}
func (log *LoginFlow) Free() {}

func (log *LoginFlow) GetTransType() int8 {
	return mdb.TRANS_TYPE_XDLOG
}

func (log *LoginFlow) SetEventId(nowtime time.Time) {
	log.Microtime = fmt.Sprintf("%d.0000",nowtime.Unix())
	log.Eid = createeid(nowtime)
}

func (log *LoginFlow) CommitToFile(f *mdb.SyncFile) {
	buff:=log.Packet()
	f.WriteXdLog(buff)
}

// 严格的 NewStucture functions 防止松散的创建方式带来未初始化的字段
func NewLoginFlow(Cid int32, Account string, Pid int64, Name string, Ip string, Type int32, Ispay int32, Level int32, Vip int32) *LoginFlow {
	s_LoginFlow:=new(LoginFlow)
	
		s_LoginFlow.Sid = g_GameSrvID
		s_LoginFlow.Time = gotime.GetNowTime()
	
	s_LoginFlow.Cid=Cid
	s_LoginFlow.Account=Account
	s_LoginFlow.Pid=Pid
	s_LoginFlow.Name=Name
	s_LoginFlow.Ip=Ip
	s_LoginFlow.Type=Type
	s_LoginFlow.Ispay=Ispay
	s_LoginFlow.Level=Level
	s_LoginFlow.Vip=Vip
	return s_LoginFlow
}

// 此结构体由工具按XML描述自动生成
type CreateFlow struct {
	Tag string `json:"tag"`
	Eid uint64 `json:"eid"`
	Microtime string `json:"microtime"`
	Cid	int32 `json:"cid"` // (必填)平台渠道ID
	Sid	int32 `json:"sid"` // (必填)区唯一ID,如果不分区则为0
	Account	string `json:"account"` // (必填)平台帐号
	Pid	int64 `json:"pid"` // (必填)玩家角色ID
	Name	string `json:"name"` // (必填)角色名
	Time	int64 `json:"time"` // (必填)登陆时间戳
	Ip	string `json:"ip"` // (必填)登陆IP
	Guest	int32 `json:"guest"` // (必填)是否游客

}

// Log接口要求的Packet方法
func (log *CreateFlow) Packet() []byte {
	log.Tag = "create"
	js, err := json.Marshal(*log)
	if err != nil {
	    panic(err.Error())
	}
	js = append(js, '\n')
	return js
	//return []byte(string(js)+"\n")
}

func (log *CreateFlow) InvokeHook() {}
func (log *CreateFlow) CommitToTLog() {}
func (log *CreateFlow) CommitToXdLog() {}
func (log *CreateFlow) CommitToMySql() error { return nil }
func (log *CreateFlow) Rollback() {}
func (log *CreateFlow) Free() {}

func (log *CreateFlow) GetTransType() int8 {
	return mdb.TRANS_TYPE_XDLOG
}

func (log *CreateFlow) SetEventId(nowtime time.Time) {
	log.Microtime = fmt.Sprintf("%d.0000",nowtime.Unix())
	log.Eid = createeid(nowtime)
}

func (log *CreateFlow) CommitToFile(f *mdb.SyncFile) {
	buff:=log.Packet()
	f.WriteXdLog(buff)
}

// 严格的 NewStucture functions 防止松散的创建方式带来未初始化的字段
func NewCreateFlow(Cid int32, Account string, Pid int64, Name string, Ip string, Guest int32) *CreateFlow {
	s_CreateFlow:=new(CreateFlow)
	
		s_CreateFlow.Sid = g_GameSrvID
		s_CreateFlow.Time = gotime.GetNowTime()
	
	s_CreateFlow.Cid=Cid
	s_CreateFlow.Account=Account
	s_CreateFlow.Pid=Pid
	s_CreateFlow.Name=Name
	s_CreateFlow.Ip=Ip
	s_CreateFlow.Guest=Guest
	return s_CreateFlow
}

// 此结构体由工具按XML描述自动生成
type ChargeFlow struct {
	Tag string `json:"tag"`
	Eid uint64 `json:"eid"`
	Microtime string `json:"microtime"`
	Cid	int32 `json:"cid"` // (必填)平台渠道ID
	Sid	int32 `json:"sid"` // (必填)区唯一ID,如果不分区则为0
	Account	string `json:"account"` // (必填)平台帐号
	Pid	int64 `json:"pid"` // (必填)玩家角色ID
	Name	string `json:"name"` // (必填)角色名
	Time	int64 `json:"time"` // (必填)登陆时间戳
	Ip	string `json:"ip"` // (必填)登陆IP
	Oid	string `json:"oid"` // (必填)􏴥􏴯订单ID
	Poid	string `json:"poid"` // (必填)渠道方订单ID
	Type	string `json:"type"` // (必填)充值渠道类型
	Level	int32 `json:"level"` // (必填)充值时等级
	Amount	float64 `json:"amount"` // (必填)充值货币数
	Coins	int64 `json:"coins"` // (必填)兑换的游戏币数
	Currency	string `json:"currency"` // (必填)货币类型

}

// Log接口要求的Packet方法
func (log *ChargeFlow) Packet() []byte {
	log.Tag = "charge"
	js, err := json.Marshal(*log)
	if err != nil {
	    panic(err.Error())
	}
	js = append(js, '\n')
	return js
	//return []byte(string(js)+"\n")
}

func (log *ChargeFlow) InvokeHook() {}
func (log *ChargeFlow) CommitToTLog() {}
func (log *ChargeFlow) CommitToXdLog() {}
func (log *ChargeFlow) CommitToMySql() error { return nil }
func (log *ChargeFlow) Rollback() {}
func (log *ChargeFlow) Free() {}

func (log *ChargeFlow) GetTransType() int8 {
	return mdb.TRANS_TYPE_XDLOG
}

func (log *ChargeFlow) SetEventId(nowtime time.Time) {
	log.Microtime = fmt.Sprintf("%d.0000",nowtime.Unix())
	log.Eid = createeid(nowtime)
}

func (log *ChargeFlow) CommitToFile(f *mdb.SyncFile) {
	buff:=log.Packet()
	f.WriteXdLog(buff)
}

// 严格的 NewStucture functions 防止松散的创建方式带来未初始化的字段
func NewChargeFlow(Cid int32, Account string, Pid int64, Name string, Ip string, Oid string, Poid string, Type string, Level int32, Amount float64, Coins int64, Currency string) *ChargeFlow {
	s_ChargeFlow:=new(ChargeFlow)
	
		s_ChargeFlow.Sid = g_GameSrvID
		s_ChargeFlow.Time = gotime.GetNowTime()
	
	s_ChargeFlow.Cid=Cid
	s_ChargeFlow.Account=Account
	s_ChargeFlow.Pid=Pid
	s_ChargeFlow.Name=Name
	s_ChargeFlow.Ip=Ip
	s_ChargeFlow.Oid=Oid
	s_ChargeFlow.Poid=Poid
	s_ChargeFlow.Type=Type
	s_ChargeFlow.Level=Level
	s_ChargeFlow.Amount=Amount
	s_ChargeFlow.Coins=Coins
	s_ChargeFlow.Currency=Currency
	return s_ChargeFlow
}

// 此结构体由工具按XML描述自动生成
type EventFlow struct {
	Tag string `json:"tag"`
	Eid uint64 `json:"eid"`
	Microtime string `json:"microtime"`
	Cid	int32 `json:"cid"` // (必填)平台渠道ID
	Sid	int32 `json:"sid"` // (必填)区唯一ID,如果不分区则为0
	Account	string `json:"account"` // (必填)平台帐号
	Pid	int64 `json:"pid"` // (必填)玩家角色ID
	Name	string `json:"name"` // (必填)角色名
	Time	int64 `json:"time"` // (必填)操作时间戳
	Ispay	int32 `json:"ispay"` // (必填)是否付费玩家1:付费;0:非付费
	Type	int32 `json:"type"` // (必填)事件类型ID
	Count	int32 `json:"count"` // (必填)操作次数

}

// Log接口要求的Packet方法
func (log *EventFlow) Packet() []byte {
	log.Tag = "event"
	js, err := json.Marshal(*log)
	if err != nil {
	    panic(err.Error())
	}
	js = append(js, '\n')
	return js
	//return []byte(string(js)+"\n")
}

func (log *EventFlow) InvokeHook() {}
func (log *EventFlow) CommitToTLog() {}
func (log *EventFlow) CommitToXdLog() {}
func (log *EventFlow) CommitToMySql() error { return nil }
func (log *EventFlow) Rollback() {}
func (log *EventFlow) Free() {}

func (log *EventFlow) GetTransType() int8 {
	return mdb.TRANS_TYPE_XDLOG
}

func (log *EventFlow) SetEventId(nowtime time.Time) {
	log.Microtime = fmt.Sprintf("%d.0000",nowtime.Unix())
	log.Eid = createeid(nowtime)
}

func (log *EventFlow) CommitToFile(f *mdb.SyncFile) {
	buff:=log.Packet()
	f.WriteXdLog(buff)
}

// 严格的 NewStucture functions 防止松散的创建方式带来未初始化的字段
func NewEventFlow(Cid int32, Account string, Pid int64, Name string, Ispay int32, Type int32, Count int32) *EventFlow {
	s_EventFlow:=new(EventFlow)
	
		s_EventFlow.Sid = g_GameSrvID
		s_EventFlow.Time = gotime.GetNowTime()
	
	s_EventFlow.Cid=Cid
	s_EventFlow.Account=Account
	s_EventFlow.Pid=Pid
	s_EventFlow.Name=Name
	s_EventFlow.Ispay=Ispay
	s_EventFlow.Type=Type
	s_EventFlow.Count=Count
	return s_EventFlow
}

// 此结构体由工具按XML描述自动生成
type IncomeFlow struct {
	Tag string `json:"tag"`
	Eid uint64 `json:"eid"`
	Microtime string `json:"microtime"`
	Cid	int32 `json:"cid"` // (必填)平台渠道ID
	Sid	int32 `json:"sid"` // (必填)区唯一ID,如果不分区则为0
	Account	string `json:"account"` // (必填)平台帐号
	Pid	int64 `json:"pid"` // (必填)玩家角色ID
	Name	string `json:"name"` // (必填)角色名
	Time	int64 `json:"time"` // (必填)操作时间戳
	Ispay	int32 `json:"ispay"` // (必填)是否付费玩家1:付费;0:非付费
	Value	int32 `json:"value"` // (必填)获取数额
	Coin_type	int32 `json:"coin_type"` // (必填)游戏币类型0:铜钱,1:元宝
	Type	int32 `json:"type"` // (必填)事件类型ID
	After	int32 `json:"after"` // (必填)变动后的游戏币数
	Param	string `json:"param"` // (必填)其他参数，备用

}

// Log接口要求的Packet方法
func (log *IncomeFlow) Packet() []byte {
	log.Tag = "income"
	js, err := json.Marshal(*log)
	if err != nil {
	    panic(err.Error())
	}
	js = append(js, '\n')
	return js
	//return []byte(string(js)+"\n")
}

func (log *IncomeFlow) InvokeHook() {}
func (log *IncomeFlow) CommitToTLog() {}
func (log *IncomeFlow) CommitToXdLog() {}
func (log *IncomeFlow) CommitToMySql() error { return nil }
func (log *IncomeFlow) Rollback() {}
func (log *IncomeFlow) Free() {}

func (log *IncomeFlow) GetTransType() int8 {
	return mdb.TRANS_TYPE_XDLOG
}

func (log *IncomeFlow) SetEventId(nowtime time.Time) {
	log.Microtime = fmt.Sprintf("%d.0000",nowtime.Unix())
	log.Eid = createeid(nowtime)
}

func (log *IncomeFlow) CommitToFile(f *mdb.SyncFile) {
	buff:=log.Packet()
	f.WriteXdLog(buff)
}

// 严格的 NewStucture functions 防止松散的创建方式带来未初始化的字段
func NewIncomeFlow(Cid int32, Account string, Pid int64, Name string, Ispay int32, Value int32, Coin_type int32, Type int32, After int32, Param string) *IncomeFlow {
	s_IncomeFlow:=new(IncomeFlow)
	
		s_IncomeFlow.Sid = g_GameSrvID
		s_IncomeFlow.Time = gotime.GetNowTime()
	
	s_IncomeFlow.Cid=Cid
	s_IncomeFlow.Account=Account
	s_IncomeFlow.Pid=Pid
	s_IncomeFlow.Name=Name
	s_IncomeFlow.Ispay=Ispay
	s_IncomeFlow.Value=Value
	s_IncomeFlow.Coin_type=Coin_type
	s_IncomeFlow.Type=Type
	s_IncomeFlow.After=After
	s_IncomeFlow.Param=Param
	return s_IncomeFlow
}

// 此结构体由工具按XML描述自动生成
type ConsumeFlow struct {
	Tag string `json:"tag"`
	Eid uint64 `json:"eid"`
	Microtime string `json:"microtime"`
	Cid	int32 `json:"cid"` // (必填)平台渠道ID
	Sid	int32 `json:"sid"` // (必填)区唯一ID,如果不分区则为0
	Account	string `json:"account"` // (必填)平台帐号
	Pid	int64 `json:"pid"` // (必填)玩家角色ID
	Name	string `json:"name"` // (必填)角色名
	Time	int64 `json:"time"` // (必填)操作时间戳
	Ispay	int32 `json:"ispay"` // (必填)是否付费玩家1:付费;0:非付费
	Value	int32 `json:"value"` // (必填)获取数额
	Coin_type	int32 `json:"coin_type"` // (必填)游戏币类型0:铜钱,1:元宝
	Type	int32 `json:"type"` // (必填)事件类型ID
	Kind	int32 `json:"kind"` // (必填)消耗类型
	After	int32 `json:"after"` // (必填)变动后的游戏币数

}

// Log接口要求的Packet方法
func (log *ConsumeFlow) Packet() []byte {
	log.Tag = "consume"
	js, err := json.Marshal(*log)
	if err != nil {
	    panic(err.Error())
	}
	js = append(js, '\n')
	return js
	//return []byte(string(js)+"\n")
}

func (log *ConsumeFlow) InvokeHook() {}
func (log *ConsumeFlow) CommitToTLog() {}
func (log *ConsumeFlow) CommitToXdLog() {}
func (log *ConsumeFlow) CommitToMySql() error { return nil }
func (log *ConsumeFlow) Rollback() {}
func (log *ConsumeFlow) Free() {}

func (log *ConsumeFlow) GetTransType() int8 {
	return mdb.TRANS_TYPE_XDLOG
}

func (log *ConsumeFlow) SetEventId(nowtime time.Time) {
	log.Microtime = fmt.Sprintf("%d.0000",nowtime.Unix())
	log.Eid = createeid(nowtime)
}

func (log *ConsumeFlow) CommitToFile(f *mdb.SyncFile) {
	buff:=log.Packet()
	f.WriteXdLog(buff)
}

// 严格的 NewStucture functions 防止松散的创建方式带来未初始化的字段
func NewConsumeFlow(Cid int32, Account string, Pid int64, Name string, Ispay int32, Value int32, Coin_type int32, Type int32, Kind int32, After int32) *ConsumeFlow {
	s_ConsumeFlow:=new(ConsumeFlow)
	
		s_ConsumeFlow.Sid = g_GameSrvID
		s_ConsumeFlow.Time = gotime.GetNowTime()
	
	s_ConsumeFlow.Cid=Cid
	s_ConsumeFlow.Account=Account
	s_ConsumeFlow.Pid=Pid
	s_ConsumeFlow.Name=Name
	s_ConsumeFlow.Ispay=Ispay
	s_ConsumeFlow.Value=Value
	s_ConsumeFlow.Coin_type=Coin_type
	s_ConsumeFlow.Type=Type
	s_ConsumeFlow.Kind=Kind
	s_ConsumeFlow.After=After
	return s_ConsumeFlow
}

// 此结构体由工具按XML描述自动生成
type ItemFlow struct {
	Tag string `json:"tag"`
	Eid uint64 `json:"eid"`
	Microtime string `json:"microtime"`
	Cid	int32 `json:"cid"` // (必填)平台渠道ID
	Sid	int32 `json:"sid"` // (必填)区唯一ID,如果不分区则为0
	Account	string `json:"account"` // (必填)平台帐号
	Pid	int64 `json:"pid"` // (必填)玩家角色ID
	Name	string `json:"name"` // (必填)角色名
	Time	int64 `json:"time"` // (必填)操作时间戳
	Ispay	int32 `json:"ispay"` // (必填)是否付费玩家1:付费;0:非付费
	Itemid	int32 `json:"itemid"` // (必填)物品ID
	Value	int32 `json:"value"` // (必填)获取数额
	From_pid	int32 `json:"from_pid"` // (必填)来源玩家ID,系统为0
	Type	int32 `json:"type"` // (必填)事件类型ID
	After	int32 `json:"after"` // (必填)变动后的游戏币数
	Param	string `json:"param"` // (必填)其他参数，备用

}

// Log接口要求的Packet方法
func (log *ItemFlow) Packet() []byte {
	log.Tag = "item"
	js, err := json.Marshal(*log)
	if err != nil {
	    panic(err.Error())
	}
	js = append(js, '\n')
	return js
	//return []byte(string(js)+"\n")
}

func (log *ItemFlow) InvokeHook() {}
func (log *ItemFlow) CommitToTLog() {}
func (log *ItemFlow) CommitToXdLog() {}
func (log *ItemFlow) CommitToMySql() error { return nil }
func (log *ItemFlow) Rollback() {}
func (log *ItemFlow) Free() {}

func (log *ItemFlow) GetTransType() int8 {
	return mdb.TRANS_TYPE_XDLOG
}

func (log *ItemFlow) SetEventId(nowtime time.Time) {
	log.Microtime = fmt.Sprintf("%d.0000",nowtime.Unix())
	log.Eid = createeid(nowtime)
}

func (log *ItemFlow) CommitToFile(f *mdb.SyncFile) {
	buff:=log.Packet()
	f.WriteXdLog(buff)
}

// 严格的 NewStucture functions 防止松散的创建方式带来未初始化的字段
func NewItemFlow(Cid int32, Account string, Pid int64, Name string, Ispay int32, Itemid int32, Value int32, From_pid int32, Type int32, After int32, Param string) *ItemFlow {
	s_ItemFlow:=new(ItemFlow)
	
		s_ItemFlow.Sid = g_GameSrvID
		s_ItemFlow.Time = gotime.GetNowTime()
	
	s_ItemFlow.Cid=Cid
	s_ItemFlow.Account=Account
	s_ItemFlow.Pid=Pid
	s_ItemFlow.Name=Name
	s_ItemFlow.Ispay=Ispay
	s_ItemFlow.Itemid=Itemid
	s_ItemFlow.Value=Value
	s_ItemFlow.From_pid=From_pid
	s_ItemFlow.Type=Type
	s_ItemFlow.After=After
	s_ItemFlow.Param=Param
	return s_ItemFlow
}

// 此结构体由工具按XML描述自动生成
type PropsFlow struct {
	Tag string `json:"tag"`
	Eid uint64 `json:"eid"`
	Microtime string `json:"microtime"`
	Cid	int32 `json:"cid"` // (必填)平台渠道ID
	Sid	int32 `json:"sid"` // (必填)区唯一ID,如果不分区则为0
	Account	string `json:"account"` // (必填)平台帐号
	Pid	int64 `json:"pid"` // (必填)玩家角色ID
	Name	string `json:"name"` // (必填)角色名
	Time	int64 `json:"time"` // (必填)操作时间戳
	Ispay	int32 `json:"ispay"` // (必填)是否付费玩家1:付费;0:非付费
	Itemid	int32 `json:"itemid"` // (必填)物品ID
	Value	int32 `json:"value"` // (必填)消费数额
	Type	int32 `json:"type"` // (必填)事件类型ID
	After	int32 `json:"after"` // (必填)变动后的游戏币数

}

// Log接口要求的Packet方法
func (log *PropsFlow) Packet() []byte {
	log.Tag = "props"
	js, err := json.Marshal(*log)
	if err != nil {
	    panic(err.Error())
	}
	js = append(js, '\n')
	return js
	//return []byte(string(js)+"\n")
}

func (log *PropsFlow) InvokeHook() {}
func (log *PropsFlow) CommitToTLog() {}
func (log *PropsFlow) CommitToXdLog() {}
func (log *PropsFlow) CommitToMySql() error { return nil }
func (log *PropsFlow) Rollback() {}
func (log *PropsFlow) Free() {}

func (log *PropsFlow) GetTransType() int8 {
	return mdb.TRANS_TYPE_XDLOG
}

func (log *PropsFlow) SetEventId(nowtime time.Time) {
	log.Microtime = fmt.Sprintf("%d.0000",nowtime.Unix())
	log.Eid = createeid(nowtime)
}

func (log *PropsFlow) CommitToFile(f *mdb.SyncFile) {
	buff:=log.Packet()
	f.WriteXdLog(buff)
}

// 严格的 NewStucture functions 防止松散的创建方式带来未初始化的字段
func NewPropsFlow(Cid int32, Account string, Pid int64, Name string, Ispay int32, Itemid int32, Value int32, Type int32, After int32) *PropsFlow {
	s_PropsFlow:=new(PropsFlow)
	
		s_PropsFlow.Sid = g_GameSrvID
		s_PropsFlow.Time = gotime.GetNowTime()
	
	s_PropsFlow.Cid=Cid
	s_PropsFlow.Account=Account
	s_PropsFlow.Pid=Pid
	s_PropsFlow.Name=Name
	s_PropsFlow.Ispay=Ispay
	s_PropsFlow.Itemid=Itemid
	s_PropsFlow.Value=Value
	s_PropsFlow.Type=Type
	s_PropsFlow.After=After
	return s_PropsFlow
}

// 此结构体由工具按XML描述自动生成
type LevelFlow struct {
	Tag string `json:"tag"`
	Eid uint64 `json:"eid"`
	Microtime string `json:"microtime"`
	Cid	int32 `json:"cid"` // (必填)平台渠道ID
	Sid	int32 `json:"sid"` // (必填)区唯一ID,如果不分区则为0
	Account	string `json:"account"` // (必填)平台帐号
	Pid	int64 `json:"pid"` // (必填)玩家角色ID
	Name	string `json:"name"` // (必填)角色名
	Time	int64 `json:"time"` // (必填)时间戳
	Exp	int32 `json:"exp"` // (必填)获得经验
	From	int32 `json:"from"` // (必填)升级前等级
	To	int32 `json:"to"` // (必填)升级后等级
	Ispay	int32 `json:"ispay"` // (必填)是否付费玩家1:付费;0:非付费

}

// Log接口要求的Packet方法
func (log *LevelFlow) Packet() []byte {
	log.Tag = "level"
	js, err := json.Marshal(*log)
	if err != nil {
	    panic(err.Error())
	}
	js = append(js, '\n')
	return js
	//return []byte(string(js)+"\n")
}

func (log *LevelFlow) InvokeHook() {}
func (log *LevelFlow) CommitToTLog() {}
func (log *LevelFlow) CommitToXdLog() {}
func (log *LevelFlow) CommitToMySql() error { return nil }
func (log *LevelFlow) Rollback() {}
func (log *LevelFlow) Free() {}

func (log *LevelFlow) GetTransType() int8 {
	return mdb.TRANS_TYPE_XDLOG
}

func (log *LevelFlow) SetEventId(nowtime time.Time) {
	log.Microtime = fmt.Sprintf("%d.0000",nowtime.Unix())
	log.Eid = createeid(nowtime)
}

func (log *LevelFlow) CommitToFile(f *mdb.SyncFile) {
	buff:=log.Packet()
	f.WriteXdLog(buff)
}

// 严格的 NewStucture functions 防止松散的创建方式带来未初始化的字段
func NewLevelFlow(Cid int32, Account string, Pid int64, Name string, Exp int32, From int32, To int32, Ispay int32) *LevelFlow {
	s_LevelFlow:=new(LevelFlow)
	
		s_LevelFlow.Sid = g_GameSrvID
		s_LevelFlow.Time = gotime.GetNowTime()
	
	s_LevelFlow.Cid=Cid
	s_LevelFlow.Account=Account
	s_LevelFlow.Pid=Pid
	s_LevelFlow.Name=Name
	s_LevelFlow.Exp=Exp
	s_LevelFlow.From=From
	s_LevelFlow.To=To
	s_LevelFlow.Ispay=Ispay
	return s_LevelFlow
}

// 此结构体由工具按XML描述自动生成
type VipFlow struct {
	Tag string `json:"tag"`
	Eid uint64 `json:"eid"`
	Microtime string `json:"microtime"`
	Cid	int32 `json:"cid"` // (必填)平台渠道ID
	Sid	int32 `json:"sid"` // (必填)区唯一ID,如果不分区则为0
	Account	string `json:"account"` // (必填)平台帐号
	Pid	int64 `json:"pid"` // (必填)玩家角色ID
	Name	string `json:"name"` // (必填)角色名
	Time	int64 `json:"time"` // (必填)时间戳
	Exp	int32 `json:"exp"` // (必填)获得VIP经验
	From	int32 `json:"from"` // (必填)升级前VIP等级
	To	int32 `json:"to"` // (必填)升级后VIP等级
	Ispay	int32 `json:"ispay"` // (必填)是否付费玩家1:付费;0:非付费

}

// Log接口要求的Packet方法
func (log *VipFlow) Packet() []byte {
	log.Tag = "vip"
	js, err := json.Marshal(*log)
	if err != nil {
	    panic(err.Error())
	}
	js = append(js, '\n')
	return js
	//return []byte(string(js)+"\n")
}

func (log *VipFlow) InvokeHook() {}
func (log *VipFlow) CommitToTLog() {}
func (log *VipFlow) CommitToXdLog() {}
func (log *VipFlow) CommitToMySql() error { return nil }
func (log *VipFlow) Rollback() {}
func (log *VipFlow) Free() {}

func (log *VipFlow) GetTransType() int8 {
	return mdb.TRANS_TYPE_XDLOG
}

func (log *VipFlow) SetEventId(nowtime time.Time) {
	log.Microtime = fmt.Sprintf("%d.0000",nowtime.Unix())
	log.Eid = createeid(nowtime)
}

func (log *VipFlow) CommitToFile(f *mdb.SyncFile) {
	buff:=log.Packet()
	f.WriteXdLog(buff)
}

// 严格的 NewStucture functions 防止松散的创建方式带来未初始化的字段
func NewVipFlow(Cid int32, Account string, Pid int64, Name string, Exp int32, From int32, To int32, Ispay int32) *VipFlow {
	s_VipFlow:=new(VipFlow)
	
		s_VipFlow.Sid = g_GameSrvID
		s_VipFlow.Time = gotime.GetNowTime()
	
	s_VipFlow.Cid=Cid
	s_VipFlow.Account=Account
	s_VipFlow.Pid=Pid
	s_VipFlow.Name=Name
	s_VipFlow.Exp=Exp
	s_VipFlow.From=From
	s_VipFlow.To=To
	s_VipFlow.Ispay=Ispay
	return s_VipFlow
}

// 此结构体由工具按XML描述自动生成
type OnlineFlow struct {
	Tag string `json:"tag"`
	Eid uint64 `json:"eid"`
	Microtime string `json:"microtime"`
	Cid	int32 `json:"cid"` // (必填)平台渠道ID
	Sid	int32 `json:"sid"` // (必填)区唯一ID,如果不分区则为0
	Time	int64 `json:"time"` // (必填)时间戳
	Count_all	int32 `json:"count_all"` // (必填)总在线人数

}

// Log接口要求的Packet方法
func (log *OnlineFlow) Packet() []byte {
	log.Tag = "online"
	js, err := json.Marshal(*log)
	if err != nil {
	    panic(err.Error())
	}
	js = append(js, '\n')
	return js
	//return []byte(string(js)+"\n")
}

func (log *OnlineFlow) InvokeHook() {}
func (log *OnlineFlow) CommitToTLog() {}
func (log *OnlineFlow) CommitToXdLog() {}
func (log *OnlineFlow) CommitToMySql() error { return nil }
func (log *OnlineFlow) Rollback() {}
func (log *OnlineFlow) Free() {}

func (log *OnlineFlow) GetTransType() int8 {
	return mdb.TRANS_TYPE_XDLOG
}

func (log *OnlineFlow) SetEventId(nowtime time.Time) {
	log.Microtime = fmt.Sprintf("%d.0000",nowtime.Unix())
	log.Eid = createeid(nowtime)
}

func (log *OnlineFlow) CommitToFile(f *mdb.SyncFile) {
	buff:=log.Packet()
	f.WriteXdLog(buff)
}

// 严格的 NewStucture functions 防止松散的创建方式带来未初始化的字段
func NewOnlineFlow(Cid int32, Count_all int32) *OnlineFlow {
	s_OnlineFlow:=new(OnlineFlow)
	
		s_OnlineFlow.Sid = g_GameSrvID
		s_OnlineFlow.Time = gotime.GetNowTime()
	
	s_OnlineFlow.Cid=Cid
	s_OnlineFlow.Count_all=Count_all
	return s_OnlineFlow
}

// 此结构体由工具按XML描述自动生成
type SnapshotFlow struct {
	Tag string `json:"tag"`
	Eid uint64 `json:"eid"`
	Microtime string `json:"microtime"`
	Cid	int32 `json:"cid"` // (必填)平台渠道ID
	Sid	int32 `json:"sid"` // (必填)区唯一ID,如果不分区则为0
	Account	string `json:"account"` // (必填)平台帐号
	Pid	int64 `json:"pid"` // (必填)玩家角色ID
	Name	string `json:"name"` // (必填)角色名
	Coins	int64 `json:"coins"` // (必填)当前剩余的全部游戏币数量
	Charge_coins	int64 `json:"charge_coins"` // (必填)剩余的充值的游戏币数量
	Login_time	int64 `json:"login_time"` // (必填)最后登录时间
	Item	string `json:"item"` // (必填)对应物品id的json数据
	Level	int32 `json:"level"` // (必填)对应角色等级
	Vip_level	int32 `json:"vip_level"` // (必填)VIP等级
	FriendsLevel	int32 `json:"friendslevel"` // (必填)伙伴等级的json数据

}

// Log接口要求的Packet方法
func (log *SnapshotFlow) Packet() []byte {
	log.Tag = "snapshot"
	js, err := json.Marshal(*log)
	if err != nil {
	    panic(err.Error())
	}
	js = append(js, '\n')
	return js
	//return []byte(string(js)+"\n")
}

func (log *SnapshotFlow) InvokeHook() {}
func (log *SnapshotFlow) CommitToTLog() {}
func (log *SnapshotFlow) CommitToXdLog() {}
func (log *SnapshotFlow) CommitToMySql() error { return nil }
func (log *SnapshotFlow) Rollback() {}
func (log *SnapshotFlow) Free() {}

func (log *SnapshotFlow) GetTransType() int8 {
	return mdb.TRANS_TYPE_XDLOG
}

func (log *SnapshotFlow) SetEventId(nowtime time.Time) {
	log.Microtime = fmt.Sprintf("%d.0000",nowtime.Unix())
	log.Eid = createeid(nowtime)
}

func (log *SnapshotFlow) CommitToFile(f *mdb.SyncFile) {
	buff:=log.Packet()
	f.WriteXdLog(buff)
}

// 严格的 NewStucture functions 防止松散的创建方式带来未初始化的字段
func NewSnapshotFlow(Cid int32, Account string, Pid int64, Name string, Coins int64, Charge_coins int64, Login_time int64, Item string, Level int32, Vip_level int32, FriendsLevel int32) *SnapshotFlow {
	s_SnapshotFlow:=new(SnapshotFlow)
	
		s_SnapshotFlow.Sid = g_GameSrvID
	
	s_SnapshotFlow.Cid=Cid
	s_SnapshotFlow.Account=Account
	s_SnapshotFlow.Pid=Pid
	s_SnapshotFlow.Name=Name
	s_SnapshotFlow.Coins=Coins
	s_SnapshotFlow.Charge_coins=Charge_coins
	s_SnapshotFlow.Login_time=Login_time
	s_SnapshotFlow.Item=Item
	s_SnapshotFlow.Level=Level
	s_SnapshotFlow.Vip_level=Vip_level
	s_SnapshotFlow.FriendsLevel=FriendsLevel
	return s_SnapshotFlow
}

// 此结构体由工具按XML描述自动生成
type MissionFlow struct {
	Tag string `json:"tag"`
	Eid uint64 `json:"eid"`
	Microtime string `json:"microtime"`
	Cid	int32 `json:"cid"` // (必填)平台渠道ID
	Sid	int32 `json:"sid"` // (必填)区唯一ID,如果不分区则为0
	Account	string `json:"account"` // (必填)平台帐号
	Pid	int64 `json:"pid"` // (必填)玩家角色ID
	Name	string `json:"name"` // (必填)角色名
	Time	int64 `json:"time"` // (必填)操作时间戳
	MissionID	int32 `json:"missionid"` // (必填)副本ID
	MissionType	int32 `json:"missiontype"` // (必填)副本类型
	MissionLock	int32 `json:"missionlock"` // (必填)副本权值
	MaxLock	int32 `json:"maxlock"` // (必填)人物最大权值
	Ispay	int32 `json:"ispay"` // (必填)是否付费玩家1:付费;0:非付费
	Level	int32 `json:"level"` // (必填)等级
	Vip	int32 `json:"vip"` // (必填)VIP等级
	Action	int32 `json:"action"` // (必填)进入0/完成1/扫荡2

}

// Log接口要求的Packet方法
func (log *MissionFlow) Packet() []byte {
	log.Tag = "mission"
	js, err := json.Marshal(*log)
	if err != nil {
	    panic(err.Error())
	}
	js = append(js, '\n')
	return js
	//return []byte(string(js)+"\n")
}

func (log *MissionFlow) InvokeHook() {}
func (log *MissionFlow) CommitToTLog() {}
func (log *MissionFlow) CommitToXdLog() {}
func (log *MissionFlow) CommitToMySql() error { return nil }
func (log *MissionFlow) Rollback() {}
func (log *MissionFlow) Free() {}

func (log *MissionFlow) GetTransType() int8 {
	return mdb.TRANS_TYPE_XDLOG
}

func (log *MissionFlow) SetEventId(nowtime time.Time) {
	log.Microtime = fmt.Sprintf("%d.0000",nowtime.Unix())
	log.Eid = createeid(nowtime)
}

func (log *MissionFlow) CommitToFile(f *mdb.SyncFile) {
	buff:=log.Packet()
	f.WriteXdLog(buff)
}

// 严格的 NewStucture functions 防止松散的创建方式带来未初始化的字段
func NewMissionFlow(Cid int32, Account string, Pid int64, Name string, MissionID int32, MissionType int32, MissionLock int32, MaxLock int32, Ispay int32, Level int32, Vip int32, Action int32) *MissionFlow {
	s_MissionFlow:=new(MissionFlow)
	
		s_MissionFlow.Sid = g_GameSrvID
		s_MissionFlow.Time = gotime.GetNowTime()
	
	s_MissionFlow.Cid=Cid
	s_MissionFlow.Account=Account
	s_MissionFlow.Pid=Pid
	s_MissionFlow.Name=Name
	s_MissionFlow.MissionID=MissionID
	s_MissionFlow.MissionType=MissionType
	s_MissionFlow.MissionLock=MissionLock
	s_MissionFlow.MaxLock=MaxLock
	s_MissionFlow.Ispay=Ispay
	s_MissionFlow.Level=Level
	s_MissionFlow.Vip=Vip
	s_MissionFlow.Action=Action
	return s_MissionFlow
}

// 此结构体由工具按XML描述自动生成
type GiftcodeFlow struct {
	Tag string `json:"tag"`
	Eid uint64 `json:"eid"`
	Microtime string `json:"microtime"`
	Cid	int32 `json:"cid"` // (必填)平台渠道ID
	Sid	int32 `json:"sid"` // (必填)区唯一ID,如果不分区则为0
	Account	string `json:"account"` // (必填)平台帐号
	Pid	int64 `json:"pid"` // (必填)玩家角色ID
	Name	string `json:"name"` // (必填)角色名
	Time	int64 `json:"time"` // (必填)操作时间戳
	Code	string `json:"code"` // (必填)兑换码
	CodeType	int32 `json:"codetype"` // (必填)兑换码类型
	Ispay	int32 `json:"ispay"` // (必填)是否付费玩家1:付费;0:非付费
	Level	int32 `json:"level"` // (必填)等级

}

// Log接口要求的Packet方法
func (log *GiftcodeFlow) Packet() []byte {
	log.Tag = "giftcode"
	js, err := json.Marshal(*log)
	if err != nil {
	    panic(err.Error())
	}
	js = append(js, '\n')
	return js
	//return []byte(string(js)+"\n")
}

func (log *GiftcodeFlow) InvokeHook() {}
func (log *GiftcodeFlow) CommitToTLog() {}
func (log *GiftcodeFlow) CommitToXdLog() {}
func (log *GiftcodeFlow) CommitToMySql() error { return nil }
func (log *GiftcodeFlow) Rollback() {}
func (log *GiftcodeFlow) Free() {}

func (log *GiftcodeFlow) GetTransType() int8 {
	return mdb.TRANS_TYPE_XDLOG
}

func (log *GiftcodeFlow) SetEventId(nowtime time.Time) {
	log.Microtime = fmt.Sprintf("%d.0000",nowtime.Unix())
	log.Eid = createeid(nowtime)
}

func (log *GiftcodeFlow) CommitToFile(f *mdb.SyncFile) {
	buff:=log.Packet()
	f.WriteXdLog(buff)
}

// 严格的 NewStucture functions 防止松散的创建方式带来未初始化的字段
func NewGiftcodeFlow(Cid int32, Account string, Pid int64, Name string, Code string, CodeType int32, Ispay int32, Level int32) *GiftcodeFlow {
	s_GiftcodeFlow:=new(GiftcodeFlow)
	
		s_GiftcodeFlow.Sid = g_GameSrvID
		s_GiftcodeFlow.Time = gotime.GetNowTime()
	
	s_GiftcodeFlow.Cid=Cid
	s_GiftcodeFlow.Account=Account
	s_GiftcodeFlow.Pid=Pid
	s_GiftcodeFlow.Name=Name
	s_GiftcodeFlow.Code=Code
	s_GiftcodeFlow.CodeType=CodeType
	s_GiftcodeFlow.Ispay=Ispay
	s_GiftcodeFlow.Level=Level
	return s_GiftcodeFlow
}

// 此结构体由工具按XML描述自动生成
type FightnumFlow struct {
	Tag string `json:"tag"`
	Eid uint64 `json:"eid"`
	Microtime string `json:"microtime"`
	Cid	int32 `json:"cid"` // (必填)平台渠道ID
	Sid	int32 `json:"sid"` // (必填)区唯一ID,如果不分区则为0
	Account	string `json:"account"` // (必填)平台帐号
	Pid	int64 `json:"pid"` // (必填)玩家角色ID
	Name	string `json:"name"` // (必填)角色名
	Time	int64 `json:"time"` // (必填)操作时间戳
	FightNum	int64 `json:"fightnum"` // (必填)战力
	Ispay	int32 `json:"ispay"` // (必填)是否付费玩家1:付费;0:非付费
	Level	int32 `json:"level"` // (必填)等级

}

// Log接口要求的Packet方法
func (log *FightnumFlow) Packet() []byte {
	log.Tag = "fightnum"
	js, err := json.Marshal(*log)
	if err != nil {
	    panic(err.Error())
	}
	js = append(js, '\n')
	return js
	//return []byte(string(js)+"\n")
}

func (log *FightnumFlow) InvokeHook() {}
func (log *FightnumFlow) CommitToTLog() {}
func (log *FightnumFlow) CommitToXdLog() {}
func (log *FightnumFlow) CommitToMySql() error { return nil }
func (log *FightnumFlow) Rollback() {}
func (log *FightnumFlow) Free() {}

func (log *FightnumFlow) GetTransType() int8 {
	return mdb.TRANS_TYPE_XDLOG
}

func (log *FightnumFlow) SetEventId(nowtime time.Time) {
	log.Microtime = fmt.Sprintf("%d.0000",nowtime.Unix())
	log.Eid = createeid(nowtime)
}

func (log *FightnumFlow) CommitToFile(f *mdb.SyncFile) {
	buff:=log.Packet()
	f.WriteXdLog(buff)
}

// 严格的 NewStucture functions 防止松散的创建方式带来未初始化的字段
func NewFightnumFlow(Cid int32, Account string, Pid int64, Name string, FightNum int64, Ispay int32, Level int32) *FightnumFlow {
	s_FightnumFlow:=new(FightnumFlow)
	
		s_FightnumFlow.Sid = g_GameSrvID
		s_FightnumFlow.Time = gotime.GetNowTime()
	
	s_FightnumFlow.Cid=Cid
	s_FightnumFlow.Account=Account
	s_FightnumFlow.Pid=Pid
	s_FightnumFlow.Name=Name
	s_FightnumFlow.FightNum=FightNum
	s_FightnumFlow.Ispay=Ispay
	s_FightnumFlow.Level=Level
	return s_FightnumFlow
}


/*
func createeid(time time.Time) uint64 {
	timenanoid := time.UnixNano()
	if ServerCfg.EnableGlobalServer {
		eid, err := strconv.ParseUint(fmt.Sprintf("%v", timenanoid)+"1", 10, 64)
		fail.When(err != nil, err)
		return eid
	}
	eid, err := strconv.ParseUint(fmt.Sprintf("%v", timenanoid)+"0", 10, 64)
	fail.When(err != nil, err)
	return eid
}
*/
