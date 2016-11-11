package rpc

import (
	"game_server/mdb"
)

// 角色装备信息对象
type SEquipInfo struct {
	EquipId   uint64 // 装备ID
	EquipName string // 装备名称
	Level     uint32 // 装备精炼等级
	IsBattle  uint8  // 上阵1/未上阵0

}

/*

	玩家当前个人信息

*/
type Args_IdipGetUserinfo struct {
	RPCArgTag
	OpenId string
}

type Reply_IdipGetUserinfo struct {
	Level         uint32 // 当前等级
	Vip           uint32 // 当前VIP等级
	Exp           uint32 // 当前经验
	Coin          uint32 // 当前铜钱
	Gold          uint32 // 当前元宝数量
	Physical      uint32 // 当前体力值
	MaxPhysical   uint32 // 体力值上限
	MaxBag        uint32 // 背包上限值
	RegisterTime  uint64 // 注册时间
	IsOnline      uint8  // 是否在线（0在线，1离线）
	AccStatus     uint8  // 帐号状态（0 正常，1封号）
	BanEndTime    uint64 // 封号截至时间
	ArmyId        uint64 // 所在公会
	RankInArmy    uint32 // 在公会中的排名
	ArmyRank      uint32 // 公会排名
	PassProgress  uint32 // 当前关卡进度
	PvpRank       uint32 // 个人PVP排名
	PvpScore      uint32 // PVP积分数量
	LastLoginTime string //玩家最后登录时间
	RoleName      string //角色名称
	ErrMsg        string // 错误信息
}

func RemoteIdipGetUserInfo(openId string, sid int, callback func(*Reply_IdipGetUserinfo, error)) {
	reply := &Reply_IdipGetUserinfo{}

	args := &Args_IdipGetUserinfo{OpenId: openId}
	Remote.Call(sid, mdb.RPC_Remote_IdipGetUserInfo, args, reply, func(err error) {
		callback(reply, err)
	})

}

/*
	查询装备信息
*/
type Args_IdipGetEquipinfo struct {
	RPCArgTag
	OpenId string
	RoleId int16
}

type Reply_IdipGetEquipinfo struct {
	EquipList_count uint32       // 角色装备信息列表的最大数量
	EquipList       []SEquipInfo // 角色装备信息列表
	ErrMsg          string       // 错误信息
}

func RemoteIdipGetEquipInfo(openId string, sid int, roleId int16, callback func(*Reply_IdipGetEquipinfo, error)) {
	reply := &Reply_IdipGetEquipinfo{}
	args := &Args_IdipGetEquipinfo{OpenId: openId, RoleId: roleId}
	Remote.Call(sid, mdb.RPC_Remote_IdipGetEquipInfo, args, reply, func(err error) {
		callback(reply, err)
	})
}

/*
	查询背包信息
*/

// 背包存量信息对象
type SBagInfo struct {
	ItemName string // 道具名称
	ItemId   uint64 // 道具ID（包括灵宠、魂侍、剑心、消耗品道具等）
	ItemNum  uint32 // 道具存量

}
type Args_IdipGetBaginfo struct {
	RPCArgTag
	OpenId    string
	BeginTime uint64 // 开始时间
	EndTime   uint64 // 结束时间
}

type Reply_IdipGetBaginfo struct {
	BagList_count uint32     // 背包存量信息列表的最大数量
	BagList       []SBagInfo // 背包存量信息列表
	ErrMsg        string     // 错误信息
}

func RemoteIdipGetBagInfo(openId string, sid int, beginTime, endTime uint64, callback func(*Reply_IdipGetBaginfo, error)) {
	reply := &Reply_IdipGetBaginfo{}
	args := &Args_IdipGetBaginfo{OpenId: openId, BeginTime: beginTime, EndTime: endTime}
	Remote.Call(sid, mdb.RPC_Remote_IdipGetBagInfo, args, reply, func(err error) {
		callback(reply, err)
	})
}

/*
	魂侍查询
*/

// 魂侍信息对象
type SSoulInfo struct {
	Slot            uint32 // 开启槽数
	RoleName        string // 角色名称
	MajorSoul       uint32 // 主魂侍
	MajorSoulLevel  uint32 // 魂侍等级
	Minor1Soul      uint32 // 辅魂侍1
	MinorSoul1Level uint32 // 辅魂侍1等级
	Minor2Soul      uint32 // 辅魂侍2
	MinorSoul2Level uint32 // 辅魂侍2等级
	Minor3Soul      uint32 // 辅魂侍2
	MinorSoul3Level uint32 // 辅魂侍2等级
	IsBattle        uint8  // 上阵1/未上阵0

}
type Args_IdipGetSoulinfo struct {
	RPCArgTag
	OpenId string // openid
	RoleId int16  // 角色ID
}

type Reply_IdipGetSoulinfo struct {
	SoulList_count uint32      // 魂侍信息列表的最大数量
	SoulList       []SSoulInfo // 魂侍信息列表
	ErrMsg         string      // 错误信息
}

func RemoteIdipGetSoulInfo(openId string, sid int, roleId int16, callback func(*Reply_IdipGetSoulinfo, error)) {
	reply := &Reply_IdipGetSoulinfo{}
	args := &Args_IdipGetSoulinfo{OpenId: openId, RoleId: roleId}
	Remote.Call(sid, mdb.RPC_Remote_IdipGetSoulInfo, args, reply, func(err error) {
		callback(reply, err)
	})
}

/*
	剑心查询
*/

// 剑心信息对象
type SSwordInfo struct {
	Slot      uint32 // 开启槽数
	SwordName string // 剑心名
	Level     uint32 // 剑心等级
	IsBattle  uint8  // 上阵1/未上阵0

}
type Args_IdipGetSwordinfo struct {
	RPCArgTag
	OpenId string // openid
	RoleId int8   // 角色ID
}

type Reply_IdipGetSwordinfo struct {
	SwordList_count uint32       // 剑心信息列表的最大数量
	SwordList       []SSwordInfo // 剑心信息列表
	ErrMsg          string       // 错误信息
}

func RemoteIdipGetSwordInfo(openId string, sid int, roleId int8, callback func(*Reply_IdipGetSwordinfo, error)) {
	reply := &Reply_IdipGetSwordinfo{}
	args := &Args_IdipGetSwordinfo{OpenId: openId, RoleId: roleId}
	Remote.Call(sid, mdb.RPC_Remote_IdipGetSwordInfo, args, reply, func(err error) {
		callback(reply, err)
	})
}

/*
	灵宠
*/

// 角色灵宠信息对象
type SRolePetInfo struct {
	PetId    uint64 // 灵宠ID
	PetName  string // 灵宠名称
	Level    uint32 // 灵宠等级
	IsBattle uint8  // 上阵1/未上阵0

}

type Args_IdipGetPetinfo struct {
	RPCArgTag
	OpenId string // openid
	RoleId int8   // 角色ID
}

type Reply_IdipGetPetinfo struct {
	RolePetList_count uint32         // 角色灵宠信息列表的最大数量
	RolePetList       []SRolePetInfo // 角色灵宠信息列表
	ErrMsg            string         // 错误信息
}

func RemoteIdipGetPetInfo(openId string, sid int, roleId int8, callback func(*Reply_IdipGetPetinfo, error)) {
	reply := &Reply_IdipGetPetinfo{}
	args := &Args_IdipGetPetinfo{OpenId: openId, RoleId: roleId}
	Remote.Call(sid, mdb.RPC_Remote_IdipGetPetInfo, args, reply, func(err error) {
		callback(reply, err)
	})
}

/*
	比武场排名查询游戏服
*/
type Args_IdipGetRankinfoGs struct {
	RPCArgTag
	OpenId string // 当前排名
}

type Reply_IdipGetRankinfoGs struct {
	Pid    int64  //玩家playerId
	RoleId int8   //玩家roleId
	ErrMsg string // 错误信息
}

func RemoteIdipGetRankInfoGs(openId string, sid int, callback func(*Reply_IdipGetRankinfoGs, error)) {
	reply := &Reply_IdipGetRankinfoGs{}
	args := &Args_IdipGetRankinfoGs{OpenId: openId}
	Remote.Call(sid, mdb.RPC_Remote_IdipGetRankInfoGs, args, reply, func(err error) {
		callback(reply, err)
	})
}

/*
	比武场排名查询互动富
*/
type Args_IdipGetRankinfoHd struct {
	RPCArgTag
	Pid int64 //玩家playerId
}

type Reply_IdipGetRankinfoHd struct {
	Rank   int32  //玩家排名
	ErrMsg string // 错误信息
}

func RemoteIdipGetRankInfoHd(pid int64, sid int, callback func(*Reply_IdipGetRankinfoHd, error)) {
	reply := &Reply_IdipGetRankinfoHd{}
	args := &Args_IdipGetRankinfoHd{Pid: pid}
	Remote.Call(sid, mdb.RPC_Remote_IdipGetRankInfoHd, args, reply, func(err error) {
		callback(reply, err)
	})
}

/*
	玩家战力
*/
type Args_IdipGetUserFight struct {
	RPCArgTag
	OpenId string
}

type Reply_IdipGetUserFight struct {
	Fight  int32  //玩家战力
	ErrMsg string // 错误信息
}

func RemoteIdipGetUserFight(openid string, sid int, callback func(*Reply_IdipGetUserFight, error)) {
	reply := &Reply_IdipGetUserFight{}
	args := &Args_IdipGetUserFight{OpenId: openid}
	Remote.Call(sid, mdb.RPC_Remote_IdipGetUserFight, args, reply, func(err error) {
		callback(reply, err)
	})
}

/*
	玩家深渊关卡进度
*/
type Args_IdipGetHardLevelStatus struct {
	RPCArgTag
	OpenId string
}

type Reply_IdipGetHardLevelStatus struct {
	Status int32  //进度
	ErrMsg string // 错误信息
}

func RemoteIdipGetHardLevelStatus(openid string, sid int, callback func(*Reply_IdipGetHardLevelStatus, error)) {
	reply := &Reply_IdipGetHardLevelStatus{}
	args := &Args_IdipGetHardLevelStatus{OpenId: openid}
	Remote.Call(sid, mdb.RPC_Remote_IdipGetHardLevelStatus, args, reply, func(err error) {
		callback(reply, err)
	})
}
