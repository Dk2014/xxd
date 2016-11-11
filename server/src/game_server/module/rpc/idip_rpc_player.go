package rpc

import (
	"game_server/dat/physical_dat"
	"game_server/dat/player_dat"
	"game_server/mdb"
	"game_server/module"
	. "game_server/rpc"
	"game_server/tlog"
	"math"
	"time"
)

/*

	修改角色铜钱&&元宝

*/
type Args_IdipUpdateMoney struct {
	RPCArgTag
	OpenId string
	RoleId int64 //角色ID
	Value  int64 /* 数量（正加负减） */
	Mtype  int   //货币类型
}

type Reply_IdipUpdateMoney struct {
	Result int32  // 结果：（0）成功，（其他）失败
	RetMsg string // 返回消息
	ErrMsg string // 错误信息
}

func (this *RemoteServe) IdipUpdateMoney(args *Args_IdipUpdateMoney, reply *Reply_IdipUpdateMoney) error {
	return Remote.Serve(mdb.RPC_Remote_IdipUpdateMoney, args, mdb.TRANS_TAG_RPC_Serve_IdipUpdateMoney, func() error {
		pid, ok := module.Player.GetPlayerByUsername(args.OpenId)
		if !ok {
			reply.ErrMsg = "query empty"
			return nil
		}
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(pid, func(db *mdb.Database) {
				if args.Mtype == tlog.MT_COIN {
					if args.Value > 0 {
						module.Player.IncMoney(db, nil, int64(math.Abs(float64(args.Value))), args.Mtype, tlog.MFR_IDIP, 0, "")
					} else {
						module.Player.DecMoney(db, nil, int64(math.Abs(float64(args.Value))), args.Mtype, tlog.MFR_IDIP, 0)
					}
					reply.Result = 0
					reply.RetMsg = "success"
				} else {
					reply.Result = 2
					reply.RetMsg = "ingot operation not supported"
				}
			})
		})
		return nil
	})
}

/*
	修改角色经验
*/
type Args_IdipUpdateExp struct {
	RPCArgTag
	OpenId string
	RoleId int64 //角色ID
	Value  int64 /* 数量（正加负减） */
}

type Reply_IdipUpdateExp struct {
	Result int32  // 结果：（0）成功，（其他）失败
	RetMsg string // 返回消息
	ErrMsg string // 错误信息
}

func (this *RemoteServe) IdipUpdateExp(args *Args_IdipUpdateExp, reply *Reply_IdipUpdateExp) error {
	return Remote.Serve(mdb.RPC_Remote_IdipUpdateExp, args, mdb.TRANS_TAG_RPC_Serve_IdipUpdateExp, func() error {
		pid, ok := module.Player.GetPlayerByUsername(args.OpenId)
		if !ok {
			reply.ErrMsg = "query empty"
			return nil
		}
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(pid, func(db *mdb.Database) {
				mainrole := module.Role.GetMainRole(db)
				if args.Value >= 0 {
					module.Role.AddRoleExp(db, int8(args.RoleId), args.Value, mainrole.RoleId, tlog.EFT_IDIP_ADD)
					reply.Result = 0
					reply.RetMsg = "success"
				} else {
					reply.Result = 2
					reply.RetMsg = "negative exp value not supported"
				}
			})
		})
		return nil
	})
}

/*
	增加角色体力
*/
type Args_IdipUpdatePhysical struct {
	RPCArgTag
	OpenId string
	RoleId int64 //角色ID
	Value  int64 /* 数量（正加负减） */
}

type Reply_IdipUpdatePhysical struct {
	Result int32  // 结果：（0）成功，（其他）失败
	RetMsg string // 返回消息
	ErrMsg string // 错误信息
}

func (this *RemoteServe) IdipUpdatePhysical(args *Args_IdipUpdatePhysical, reply *Reply_IdipUpdatePhysical) error {
	return Remote.Serve(mdb.RPC_Remote_IdipUpdateHeart, args, mdb.TRANS_TAG_RPC_Serve_IdipUpdatePhysical, func() error {
		pid, ok := module.Player.GetPlayerByUsername(args.OpenId)
		if !ok {
			reply.ErrMsg = "query empty"
			return nil
		}
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(pid, func(db *mdb.Database) {
				playerPhysical := db.Lookup.PlayerPhysical(db.PlayerId())
				if playerPhysical.Value >= physical_dat.MAX_PHYSICAL_VALUE && args.Value >= 0 {
					reply.Result = 2
					reply.RetMsg = "physical is full"
					return
				}
				if args.Value >= 0 {
					playerPhysical.Value += int16(math.Abs(float64(args.Value)))
				} else {
					playerPhysical.Value -= int16(math.Abs(float64(args.Value)))
					if playerPhysical.Value < 0 {
						playerPhysical.Value = 0
					}
				}
				db.Update.PlayerPhysical(playerPhysical)
				reply.Result = 0
				reply.RetMsg = "success"
			})
		})
		return nil
	})
}

/*
	增加|减少角色爱心
*/
type Args_IdipUpdateHeart struct {
	RPCArgTag
	OpenId string
	RoleId int64 //角色ID
	Value  int64 /* 数量（正加负减） */
}

type Reply_IdipUpdateHeart struct {
	Result int32  // 结果：（0）成功，（其他）失败
	RetMsg string // 返回消息
	ErrMsg string // 错误信息
}

func (this *RemoteServe) IdipUpdateHeart(args *Args_IdipUpdateHeart, reply *Reply_IdipUpdateHeart) error {
	return Remote.Serve(mdb.RPC_Remote_IdipUpdateHeart, args, mdb.TRANS_TAG_RPC_Serve_IdipUpdateHeart, func() error {
		pid, ok := module.Player.GetPlayerByUsername(args.OpenId)
		if !ok {
			reply.ErrMsg = "query empty"
			return nil
		}
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(pid, func(db *mdb.Database) {
				playerHeart := db.Lookup.PlayerHeart(db.PlayerId())
				if args.Value >= 0 {
					playerHeart.Value += int16(math.Abs(float64(args.Value)))
				} else {
					playerHeart.Value -= int16(math.Abs(float64(args.Value)))
					if playerHeart.Value < 0 {
						playerHeart.Value = 0
					}
				}
				db.Update.PlayerHeart(playerHeart)
				reply.Result = 0
				reply.RetMsg = "success"
			})
		})
		return nil
	})
}

/*
	封号
*/
type Args_IdipBanUser struct {
	RPCArgTag
	OpenId  string // openid
	BanTime uint32 // 封号时长0 永久，**秒）
}

type Reply_IdipBanUser struct {
	Result int32  // 结果：（0）成功，（其他）失败
	RetMsg string // 返回消息
	ErrMsg string // 错误信息
}

func (this *RemoteServe) IdipBanUser(args *Args_IdipBanUser, reply *Reply_IdipBanUser) error {
	return Remote.Serve(mdb.RPC_Remote_IdipBanUser, args, mdb.TRANS_TAG_RPC_Serve_IdipBanUser, func() error {
		pid, ok := module.Player.GetPlayerByUsername(args.OpenId)
		if !ok {
			reply.ErrMsg = "query empty"
			return nil
		}
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(pid, func(db *mdb.Database) {
				var row = &mdb.PlayerState{}
				data := db.Lookup.PlayerState(db.PlayerId())
				row.Pid = db.PlayerId()
				row.BanStartTime = int64(time.Now().Unix())
				if args.BanTime == 0 {
					row.BanEndTime = 0
				} else {
					row.BanEndTime = row.BanStartTime + int64(args.BanTime)
				}
				if data == nil {
					db.Insert.PlayerState(row)
					reply.Result = 0
					reply.RetMsg = "success"
					return
				}
				db.Update.PlayerState(row)
				reply.Result = 0
				reply.RetMsg = "success"
			})
		})
		return nil
	})
}

/*
	解封
*/
type Args_IdipUnBanUser struct {
	RPCArgTag
	OpenId string // openid
}

type Reply_IdipUnBanUser struct {
	Result int32  // 结果：（0）成功，（其他）失败
	RetMsg string // 返回消息
	ErrMsg string // 错误信息
}

func (this *RemoteServe) IdipUnBanUser(args *Args_IdipUnBanUser, reply *Reply_IdipUnBanUser) error {
	return Remote.Serve(mdb.RPC_Remote_IdipUnBanUser, args, mdb.TRANS_TAG_RPC_Serve_IdipUnBanUser, func() error {
		pid, ok := module.Player.GetPlayerByUsername(args.OpenId)
		if !ok {
			reply.ErrMsg = "query empty"
			return nil
		}
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(pid, func(db *mdb.Database) {
				data := db.Lookup.PlayerState(db.PlayerId())
				if data != nil {
					data.BanEndTime = -1
					db.Update.PlayerState(data)
				}
				reply.Result = 0
				reply.RetMsg = "success"
			})
		})
		return nil
	})
}

/*
	封号(AQ)
*/
type Args_IdipAqBanUser struct {
	RPCArgTag
	OpenId string // openid
	Time   uint32 // 封号时长0 永久，**秒）
}

type Reply_IdipAqBanUser struct {
	Result int32  // 结果：（0）成功，（其他）失败
	RetMsg string // 返回消息
	ErrMsg string // 错误信息
}

func (this *RemoteServe) IdipAqBanUser(args *Args_IdipAqBanUser, reply *Reply_IdipAqBanUser) error {
	return Remote.Serve(mdb.RPC_Remote_IdipAqBanUser, args, mdb.TRANS_TAG_RPC_Serve_IdipAqBanUser, func() error {
		pid, ok := module.Player.GetPlayerByUsername(args.OpenId)
		if !ok {
			reply.ErrMsg = "query empty"
			return nil
		}
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(pid, func(db *mdb.Database) {
				var row = &mdb.PlayerState{}
				data := db.Lookup.PlayerState(db.PlayerId())
				row.Pid = db.PlayerId()
				row.BanStartTime = int64(time.Now().Unix())
				if args.Time == 0 {
					row.BanEndTime = 0
				} else {
					row.BanEndTime = row.BanStartTime + int64(args.Time)
				}
				if data == nil {
					db.Insert.PlayerState(row)
					reply.Result = 0
					reply.RetMsg = "success"
					return
				}
				db.Update.PlayerState(row)
				p_session, ok := module.Player.GetPlayerOnline(pid)
				if ok {
					p_session.Close()
				}
				reply.Result = 0
				reply.RetMsg = "success"
			})
		})
		return nil
	})
}

/*
	解除惩罚(AQ)
*/
type Args_IdipAqRelievePunish struct {
	RPCArgTag
	OpenId     string // openid
	RelieveBan uint8  // 解除封号（0 否，1 是）
}

type Reply_IdipAqRelievePunish struct {
	Result int32  // 结果：（0）成功，（其他）失败
	RetMsg string // 返回消息
	ErrMsg string // 错误信息
}

func (this *RemoteServe) IdipAqRelievePunish(args *Args_IdipAqRelievePunish, reply *Reply_IdipAqRelievePunish) error {
	return Remote.Serve(mdb.RPC_Remote_IdipAqRelievePunish, args, mdb.TRANS_TAG_RPC_Serve_IdipAqRelievePunish, func() error {
		pid, ok := module.Player.GetPlayerByUsername(args.OpenId)
		if !ok {
			reply.ErrMsg = "query empty"
			return nil
		}
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(pid, func(db *mdb.Database) {
				if args.RelieveBan == 1 {
					data := db.Lookup.PlayerState(db.PlayerId())
					if data != nil {
						data.BanEndTime = -1
						db.Update.PlayerState(data)
					}
				}
				reply.Result = 0
				reply.RetMsg = "success"
			})
		})
		return nil
	})
}

/*
	清零数据
*/
type Args_IdipCleanData struct {
	RPCArgTag
	OpenId        string //openid
	ClearCoin     uint8  //铜钱清零：否（0），是（1）
	ClearPhysical uint8  //体力清零：否（0），是（1）
	ClearHeart    uint8  //爱心清零：否（0），是（1）
}

type Reply_IdipCleanData struct {
	Result int32  // 结果：（0）成功，（其他）失败
	RetMsg string // 返回消息
	ErrMsg string // 错误信息
}

func (this *RemoteServe) IdipCleanData(args *Args_IdipCleanData, reply *Reply_IdipCleanData) error {
	return Remote.Serve(mdb.RPC_Remote_IdipCleanData, args, mdb.TRANS_TAG_RPC_Serve_IdipCleanData, func() error {
		pid, ok := module.Player.GetPlayerByUsername(args.OpenId)
		if !ok {
			reply.ErrMsg = "query empty"
			return nil
		}
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(pid, func(db *mdb.Database) {
				if args.ClearCoin == 1 {
					playerCoinInfo := db.Lookup.PlayerInfo(db.PlayerId())
					coinNum := playerCoinInfo.Coins
					module.Player.DecMoney(db, nil, coinNum, player_dat.COINS, tlog.MFR_IDIP_CLEAN, 0)
				}
				if args.ClearHeart == 1 {
					playerHeartInfo := db.Lookup.PlayerHeart(db.PlayerId())
					heartNum := playerHeartInfo.Value
					module.Heart.DecHeart(db, heartNum)
				}
				if args.ClearPhysical == 1 {
					playerPhysicalInfo := db.Lookup.PlayerPhysical(db.PlayerId())
					physicalNum := playerPhysicalInfo.Value
					module.Physical.Decrease(db, physicalNum, tlog.PFR_IDIP_CLEAN)
				}
				reply.Result = 0
				reply.RetMsg = "success"
			})
		})
		return nil
	})
}

/*
	设置宠物激活请求
*/
type Args_IdipAddPet struct {
	RPCArgTag
	OpenId string //openid
	PetId  int32  //宠物ID
}

type Reply_IdipAddPet struct {
	Result int32  // 结果：（0）成功，（其他）失败
	RetMsg string // 返回消息
	ErrMsg string // 错误信息
}

func (this *RemoteServe) IdipAddPet(args *Args_IdipAddPet, reply *Reply_IdipAddPet) error {
	return Remote.Serve(mdb.RPC_Remote_IdipAddPet, args, mdb.TRANS_TAG_RPC_Serve_IdipAddPet, func() error {
		pid, ok := module.Player.GetPlayerByUsername(args.OpenId)
		if !ok {
			reply.ErrMsg = "query empty"
			return nil
		}
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(pid, func(db *mdb.Database) {
				//AddPet 暂时不支持返回错误，有错误即 panic
				module.BattlePet.AddPet(db, args.PetId, tlog.IFR_IDIP_ACTIVE_PET, 0)
				result := 1
				if result == 1 {
					reply.Result = 0
					reply.RetMsg = "success"
				} else {
					reply.Result = 2
					reply.RetMsg = "add pet false"
				}
			})
		})
		return nil
	})
}
