package rpc

import (
	"game_server/dat/ghost_dat"
	"game_server/dat/item_dat"
	"game_server/dat/mission_dat"
	"game_server/dat/physical_dat"
	"game_server/dat/player_dat"
	"game_server/dat/role_dat"
	"game_server/dat/sword_soul_dat"
	"game_server/mdb"
	"game_server/module"
	. "game_server/rpc"
	"strconv"
)

// 角色装备信息对象
type SEquipInfo struct {
	EquipId   uint64 // 装备ID
	EquipName string // 装备名称
	Level     uint32 // 装备精炼等级
	IsBattle  uint8  // 上阵1/未上阵0

}

type SEquipInfoList []SEquipInfo

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

func (this *RemoteServe) IdipGetUserInfo(args *Args_IdipGetUserinfo, reply *Reply_IdipGetUserinfo) error {
	return Remote.Serve(mdb.RPC_Remote_IdipGetUserInfo, args, mdb.TRANS_TAG_RPC_Serve_IdipGetUserInfo, func() error {
		pid, ok := module.Player.GetPlayerByUsername(string(args.OpenId))
		if !ok {
			reply.ErrMsg = "query empty"
			return nil
		}
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(pid, func(db *mdb.Database) {
				//角色信息
				roleInfo := module.Role.GetMainRole(db)
				reply.Level = uint32(roleInfo.Level)
				reply.Exp = uint32(roleInfo.Exp)

				//vip信息
				playerVip := module.VIP.VIPInfo(db)
				reply.Vip = uint32(playerVip.Level)

				//玩家信息
				playerInfo := db.Lookup.PlayerInfo(db.PlayerId())
				reply.Coin = uint32(playerInfo.Coins)
				reply.Gold = 0
				reply.RegisterTime = uint64(playerInfo.FirstLoginTime)
				reply.LastLoginTime = strconv.FormatInt(playerInfo.LastLoginTime, 10)
				playerinfo2 := module.Player.GetPlayer(db)
				reply.RoleName = playerinfo2.Nick

				//体力信息
				playerPhysical := db.Lookup.PlayerPhysical(db.PlayerId())
				reply.Physical = uint32(playerPhysical.Value)

				//定义的常量
				reply.MaxPhysical = uint32(physical_dat.MAX_PHYSICAL_VALUE)
				reply.MaxBag = uint32(item_dat.MAX_BAG_NUM)

				//查询是否在线
				_, ok := module.Player.GetPlayerOnline(db.PlayerId())
				if ok {
					reply.IsOnline = uint8(0)
				} else {
					reply.IsOnline = uint8(1)
				}

				//暂无功能
				reply.AccStatus = 0
				reply.BanEndTime = 0
				reply.ArmyId = 0
				reply.RankInArmy = 0
				reply.ArmyRank = 0

				//查询关卡进度
				playerMissionLevel := db.Lookup.PlayerMissionLevel(db.PlayerId())
				reply.PassProgress = uint32(playerMissionLevel.MaxLock)

				//pvp相关
				reply.PvpRank = 0 //暂时0，需要rpc
				reply.PvpScore = 0

			})
		})
		return nil
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

func (this *RemoteServe) IdipGetEquipInfo(args *Args_IdipGetEquipinfo, reply *Reply_IdipGetEquipinfo) error {
	return Remote.Serve(mdb.RPC_Remote_IdipGetEquipInfo, args, mdb.TRANS_TAG_RPC_Serve_IdipGetEquipInfo, func() error {
		pid, ok := module.Player.GetPlayerByUsername(args.OpenId)
		if !ok {
			reply.ErrMsg = "query empty"
			return nil
		}
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(pid, func(db *mdb.Database) {
				var IsBattle uint8
				var equipcount uint32
				player_formation := db.Lookup.PlayerFormation(db.PlayerId())
				if player_formation.Pos0 == int8(args.RoleId) {
					IsBattle = 1
				}
				if player_formation.Pos1 == int8(args.RoleId) {
					IsBattle = 1
				}
				if player_formation.Pos2 == int8(args.RoleId) {
					IsBattle = 1
				}
				if player_formation.Pos3 == int8(args.RoleId) {
					IsBattle = 1
				}
				if player_formation.Pos4 == int8(args.RoleId) {
					IsBattle = 1
				}
				if player_formation.Pos5 == int8(args.RoleId) {
					IsBattle = 1
				}
				equip_list := make(SEquipInfoList, 0)
				db.Select.PlayerEquipment(func(row *mdb.PlayerEquipmentRow) {
					if row.RoleId() == int8(args.RoleId) {
						if row.ClothesId() != 0 {
							equipcount++
							clothe_info := db.Lookup.PlayerItem(row.ClothesId())
							item_info := item_dat.GetItem(clothe_info.ItemId)
							equip_list = append(equip_list, SEquipInfo{
								EquipId:   uint64(clothe_info.ItemId),
								EquipName: item_info.Name,
								IsBattle:  IsBattle,
								Level:     uint32(clothe_info.RefineLevel),
							})
						}
						if row.WeaponId() != 0 {
							equipcount++
							weapon_info := db.Lookup.PlayerItem(row.WeaponId())
							item_info := item_dat.GetItem(weapon_info.ItemId)
							equip_list = append(equip_list, SEquipInfo{
								EquipId:   uint64(weapon_info.ItemId),
								EquipName: item_info.Name,
								IsBattle:  IsBattle,
								Level:     uint32(weapon_info.RefineLevel),
							})
						}
						if row.AccessoriesId() != 0 {
							equipcount++
							accessories_info := db.Lookup.PlayerItem(row.AccessoriesId())
							item_info := item_dat.GetItem(accessories_info.ItemId)
							equip_list = append(equip_list, SEquipInfo{
								EquipId:   uint64(accessories_info.ItemId),
								EquipName: item_info.Name,
								IsBattle:  IsBattle,
								Level:     uint32(accessories_info.RefineLevel),
							})
						}
						if row.ShoeId() != 0 {
							equipcount++
							shoe_info := db.Lookup.PlayerItem(row.ShoeId())
							item_info := item_dat.GetItem(shoe_info.ItemId)
							equip_list = append(equip_list, SEquipInfo{
								EquipId:   uint64(shoe_info.ItemId),
								EquipName: item_info.Name,
								IsBattle:  IsBattle,
								Level:     uint32(shoe_info.RefineLevel),
							})
						}
					}
				})
				reply.EquipList_count = equipcount
				reply.EquipList = equip_list
			})
		})
		return nil
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

func (this *RemoteServe) IdipGetBagInfo(args *Args_IdipGetBaginfo, reply *Reply_IdipGetBaginfo) error {
	return Remote.Serve(mdb.RPC_Remote_IdipGetBagInfo, args, mdb.TRANS_TAG_RPC_Serve_IdipGetBagInfo, func() error {
		pid, ok := module.Player.GetPlayerByUsername(args.OpenId)
		if !ok {
			reply.ErrMsg = "query empty"
			return nil
		}
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(pid, func(db *mdb.Database) {
				bag_list := make(map[string]SBagInfo)
				db.Select.PlayerItem(func(row *mdb.PlayerItemRow) {
					item_info := item_dat.GetItem(row.ItemId())
					if row.IsDressed() == 0 && row.BuyBackState() == 0 {
						if v, ok := bag_list[item_info.Name]; ok {
							bag_list[item_info.Name] = SBagInfo{
								ItemName: item_info.Name,
								ItemId:   uint64(row.ItemId()),
								ItemNum:  uint32(row.Num()) + v.ItemNum,
							}
						} else {
							bag_list[item_info.Name] = SBagInfo{
								ItemName: item_info.Name,
								ItemId:   uint64(row.ItemId()),
								ItemNum:  uint32(row.Num()),
							}
						}
					}
				})
				for _, v := range bag_list {
					reply.BagList = append(reply.BagList, SBagInfo{
						ItemName: v.ItemName,
						ItemId:   v.ItemId,
						ItemNum:  v.ItemNum,
					})
				}
				reply.BagList_count = uint32(len(reply.BagList))
			})
		})
		return nil
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

const (
	POS2_NEED_ROLE_LEVEL = 25
	POS3_NEED_ROLE_LEVEL = 35
	POS4_NEED_ROLE_LEVEL = 45
)

func (this *RemoteServe) IdipGetSoulInfo(args *Args_IdipGetSoulinfo, reply *Reply_IdipGetSoulinfo) error {
	return Remote.Serve(mdb.RPC_Remote_IdipGetSoulInfo, args, mdb.TRANS_TAG_RPC_Serve_IdipGetSoulInfo, func() error {
		pid, ok := module.Player.GetPlayerByUsername(args.OpenId)
		if !ok {
			reply.ErrMsg = "query empty"
			return nil
		}
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(pid, func(db *mdb.Database) {
				var IsBattle uint8
				var Slotnum uint32
				var role *mdb.PlayerRole
				mainrole := module.Role.GetMainRole(db)
				if mainrole.RoleId == int8(args.RoleId) {
					role = mainrole
					IsBattle = 1
				} else {
					role = module.Role.GetBuddyRole(db, int8(args.RoleId))
					player_formation := db.Lookup.PlayerFormation(db.PlayerId())
					if player_formation.Pos0 == int8(args.RoleId) {
						IsBattle = 1
					}
					if player_formation.Pos1 == int8(args.RoleId) {
						IsBattle = 1
					}
					if player_formation.Pos2 == int8(args.RoleId) {
						IsBattle = 1
					}
					if player_formation.Pos3 == int8(args.RoleId) {
						IsBattle = 1
					}
					if player_formation.Pos4 == int8(args.RoleId) {
						IsBattle = 1
					}
					if player_formation.Pos5 == int8(args.RoleId) {
						IsBattle = 1
					}
				}
				if role.Level >= POS4_NEED_ROLE_LEVEL {
					Slotnum = 4
				} else if role.Level >= POS3_NEED_ROLE_LEVEL {
					Slotnum = 3
				} else if role.Level >= POS2_NEED_ROLE_LEVEL {
					Slotnum = 2
				} else {
					Slotnum = 1
				}
				var majorSoul, majorSoulLevel, minor1Soul, minorSoul1Level, minor2Soul, minorSoul2Level, minor3Soul, minorSoul3Level uint32
				var has = false
				//暂时查3个，因为接口问题
				db.Select.PlayerGhost(func(row *mdb.PlayerGhostRow) {
					switch row.Pos() {
					case ghost_dat.EQUIP_POS1:
						majorSoul = uint32(row.GhostId())
						majorSoulLevel = uint32(row.Level())
						has = true
					case ghost_dat.EQUIP_POS2:
						minor1Soul = uint32(row.GhostId())
						minorSoul1Level = uint32(row.Level())
						has = true
					case ghost_dat.EQUIP_POS3:
						minor2Soul = uint32(row.GhostId())
						minorSoul2Level = uint32(row.Level())
						has = true
					case ghost_dat.EQUIP_POS4:
						minor3Soul = uint32(row.GhostId())
						minorSoul3Level = uint32(row.Level())
						has = true
					}
				})
				if has == true {
					reply.SoulList_count = 1
					role_info := role_dat.GetRoleInfo(int8(args.RoleId))
					reply.SoulList = append(reply.SoulList, SSoulInfo{
						Slot:            Slotnum,
						RoleName:        role_info.Name,
						MajorSoul:       majorSoul,
						MajorSoulLevel:  majorSoulLevel,
						Minor1Soul:      minor1Soul,
						MinorSoul1Level: minorSoul1Level,
						Minor2Soul:      minor2Soul,
						MinorSoul2Level: minorSoul2Level,
						Minor3Soul:      minor3Soul,
						MinorSoul3Level: minorSoul3Level,
						IsBattle:        IsBattle,
					})
				} else {
					reply.SoulList_count = 0
				}
			})
		})
		return nil
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

func (this *RemoteServe) IdipGetSwordInfo(args *Args_IdipGetSwordinfo, reply *Reply_IdipGetSwordinfo) error {
	return Remote.Serve(mdb.RPC_Remote_IdipGetSwordInfo, args, mdb.TRANS_TAG_RPC_Serve_IdipGetSwordInfo, func() error {
		pid, ok := module.Player.GetPlayerByUsername(args.OpenId)
		if !ok {
			reply.ErrMsg = "query empty"
			return nil
		}
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(pid, func(db *mdb.Database) {
				var isBattle uint8
				mainrole := module.Role.GetMainRole(db)
				if mainrole.RoleId == int8(args.RoleId) {
					isBattle = 1
				} else {
					player_formation := db.Lookup.PlayerFormation(db.PlayerId())
					switch int8(args.RoleId) {
					case player_formation.Pos0:
						isBattle = 1
					case player_formation.Pos1:
						isBattle = 1
					case player_formation.Pos2:
						isBattle = 1
					case player_formation.Pos3:
						isBattle = 1
					case player_formation.Pos4:
						isBattle = 1
					case player_formation.Pos5:
						isBattle = 1
					}
				}
				db.Select.PlayerSwordSoulEquipment(func(row *mdb.PlayerSwordSoulEquipmentRow) {
					if row.RoleId() == args.RoleId {
						if row.Pos1() > 0 {
							reply.SwordList_count++
							player_sword_soul := db.Lookup.PlayerSwordSoul(row.Pos1())
							sword_info := sword_soul_dat.GetSwordSoul(player_sword_soul.SwordSoulId)
							reply.SwordList = append(reply.SwordList, SSwordInfo{
								Slot:      uint32(1),
								SwordName: sword_info.Name,
								Level:     uint32(player_sword_soul.Level),
								IsBattle:  isBattle,
							})
						}
						if row.Pos2() > 0 {
							reply.SwordList_count++
							player_sword_soul := db.Lookup.PlayerSwordSoul(row.Pos2())
							sword_info := sword_soul_dat.GetSwordSoul(player_sword_soul.SwordSoulId)
							reply.SwordList = append(reply.SwordList, SSwordInfo{
								Slot:      uint32(2),
								SwordName: sword_info.Name,
								Level:     uint32(player_sword_soul.Level),
								IsBattle:  isBattle,
							})
						}
						if row.Pos3() > 0 {
							reply.SwordList_count++
							player_sword_soul := db.Lookup.PlayerSwordSoul(row.Pos3())
							sword_info := sword_soul_dat.GetSwordSoul(player_sword_soul.SwordSoulId)
							reply.SwordList = append(reply.SwordList, SSwordInfo{
								Slot:      uint32(3),
								SwordName: sword_info.Name,
								Level:     uint32(player_sword_soul.Level),
								IsBattle:  isBattle,
							})
						}
						if row.Pos4() > 0 {
							reply.SwordList_count++
							player_sword_soul := db.Lookup.PlayerSwordSoul(row.Pos4())
							sword_info := sword_soul_dat.GetSwordSoul(player_sword_soul.SwordSoulId)
							reply.SwordList = append(reply.SwordList, SSwordInfo{
								Slot:      uint32(4),
								SwordName: sword_info.Name,
								Level:     uint32(player_sword_soul.Level),
								IsBattle:  isBattle,
							})
						}
						if row.Pos5() > 0 {
							reply.SwordList_count++
							player_sword_soul := db.Lookup.PlayerSwordSoul(row.Pos5())
							sword_info := sword_soul_dat.GetSwordSoul(player_sword_soul.SwordSoulId)
							reply.SwordList = append(reply.SwordList, SSwordInfo{
								Slot:      uint32(5),
								SwordName: sword_info.Name,
								Level:     uint32(player_sword_soul.Level),
								IsBattle:  isBattle,
							})
						}
						if row.Pos6() > 0 {
							reply.SwordList_count++
							player_sword_soul := db.Lookup.PlayerSwordSoul(row.Pos6())
							sword_info := sword_soul_dat.GetSwordSoul(player_sword_soul.SwordSoulId)
							reply.SwordList = append(reply.SwordList, SSwordInfo{
								Slot:      uint32(6),
								SwordName: sword_info.Name,
								Level:     uint32(player_sword_soul.Level),
								IsBattle:  isBattle,
							})
						}
						if row.Pos7() > 0 {
							reply.SwordList_count++
							player_sword_soul := db.Lookup.PlayerSwordSoul(row.Pos7())
							sword_info := sword_soul_dat.GetSwordSoul(player_sword_soul.SwordSoulId)
							reply.SwordList = append(reply.SwordList, SSwordInfo{
								Slot:      uint32(7),
								SwordName: sword_info.Name,
								Level:     uint32(player_sword_soul.Level),
								IsBattle:  isBattle,
							})
						}
						if row.Pos8() > 0 {
							reply.SwordList_count++
							player_sword_soul := db.Lookup.PlayerSwordSoul(row.Pos8())
							sword_info := sword_soul_dat.GetSwordSoul(player_sword_soul.SwordSoulId)
							reply.SwordList = append(reply.SwordList, SSwordInfo{
								Slot:      uint32(8),
								SwordName: sword_info.Name,
								Level:     uint32(player_sword_soul.Level),
								IsBattle:  isBattle,
							})
						}
						if row.Pos9() > 0 {
							reply.SwordList_count++
							player_sword_soul := db.Lookup.PlayerSwordSoul(row.Pos9())
							sword_info := sword_soul_dat.GetSwordSoul(player_sword_soul.SwordSoulId)
							reply.SwordList = append(reply.SwordList, SSwordInfo{
								Slot:      uint32(9),
								SwordName: sword_info.Name,
								Level:     uint32(player_sword_soul.Level),
								IsBattle:  isBattle,
							})
						}
					}
				})
			})
		})
		return nil
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

func (this *RemoteServe) IdipGetPetInfo(args *Args_IdipGetPetinfo, reply *Reply_IdipGetPetinfo) error {
	return Remote.Serve(mdb.RPC_Remote_IdipGetPetInfo, args, mdb.TRANS_TAG_RPC_Serve_IdipGetPetInfo, func() error {
		pid, ok := module.Player.GetPlayerByUsername(args.OpenId)
		if !ok {
			reply.ErrMsg = "query empty"
			return nil
		}
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(pid, func(db *mdb.Database) {
				var isBattle uint8
				mainrole := module.Role.GetMainRole(db)
				if mainrole.RoleId == int8(args.RoleId) {
					isBattle = 1
				} else {
					player_formation := db.Lookup.PlayerFormation(db.PlayerId())
					switch int8(args.RoleId) {
					case player_formation.Pos0:
						isBattle = 1
					case player_formation.Pos1:
						isBattle = 1
					case player_formation.Pos2:
						isBattle = 1
					case player_formation.Pos3:
						isBattle = 1
					case player_formation.Pos4:
						isBattle = 1
					case player_formation.Pos5:
						isBattle = 1
					}
				}
				db.Select.PlayerBattlePet(func(row *mdb.PlayerBattlePetRow) {
					pet := mission_dat.GetEnemyRole(int32(row.BattlePetId()))
					reply.RolePetList_count++
					reply.RolePetList = append(reply.RolePetList, SRolePetInfo{
						PetId:    uint64(row.BattlePetId()),
						PetName:  pet.Name,
						Level:    uint32(pet.Level),
						IsBattle: isBattle,
					})
				})
			})
		})
		return nil
	})
}

/*
	比武场排名查询（游戏服）
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

func (this *RemoteServe) IdipGetRankInfoGs(args *Args_IdipGetRankinfoGs, reply *Reply_IdipGetRankinfoGs) error {
	return Remote.Serve(mdb.RPC_Remote_IdipGetRankInfoGs, args, mdb.TRANS_TAG_RPC_Serve_IdipGetRankInfoGs, func() error {
		pid, ok := module.Player.GetPlayerByUsername(args.OpenId)
		if !ok {
			reply.ErrMsg = "query empty"
			return nil
		}
		mdb.GlobalExecute(func(globalDb *mdb.Database) {
			globalDb.AgentExecute(pid, func(db *mdb.Database) {
				reply.RoleId = module.Role.GetMainRole(db).RoleId
				reply.Pid = pid
				isOpenFunc := module.Player.IsOpenFunc(db, player_dat.FUNC_ARENA)
				if !isOpenFunc {
					reply.ErrMsg = "arena func not open"
				}
			})
		})
		return nil
	})
}

/*
	比武场排名查询
*/
type Args_IdipGetRankinfoHd struct {
	RPCArgTag
	Pid int64 // 玩家playerID
}

type Reply_IdipGetRankinfoHd struct {
	Rank   int32  //玩家排名
	ErrMsg string // 错误信息
}

func (this *RemoteServe) IdipGetRankInfoHd(args *Args_IdipGetRankinfoHd, reply *Reply_IdipGetRankinfoHd) error {
	return Remote.Serve(mdb.RPC_Remote_IdipGetRankInfoHd, args, mdb.TRANS_TAG_RPC_Serve_IdipGetRankInfoHd, func() error {
		mdb.GlobalExecute(func(globalDb *mdb.Database) {
			rank := module.ArenaRPC.GetPlayerRank(args.Pid)
			reply.Rank = rank
			if rank < 0 {
				reply.ErrMsg = "can't find player's rank"
			}
		})
		return nil
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

func (this *RemoteServe) IdipGetUserFight(args *Args_IdipGetUserFight, reply *Reply_IdipGetUserFight) error {
	return Remote.Serve(mdb.RPC_Remote_IdipGetUserFight, args, mdb.TRANS_TAG_RPC_Serve_IdipGetUserFight, func() error {
		pid, ok := module.Player.GetPlayerByUsername(args.OpenId)
		if !ok {
			reply.ErrMsg = "query empty"
			return nil
		}
		mdb.GlobalExecute(func(globalDb *mdb.Database) {
			globalDb.AgentExecute(pid, func(db *mdb.Database) {
				fightInfo := db.Lookup.PlayerFightNum(pid)
				reply.Fight = int32(fightInfo.FightNum)
			})
		})
		return nil
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

func (this *RemoteServe) IdipGetHardLevelStatus(args *Args_IdipGetHardLevelStatus, reply *Reply_IdipGetHardLevelStatus) error {
	return Remote.Serve(mdb.RPC_Remote_IdipGetHardLevelStatus, args, mdb.TRANS_TAG_RPC_Serve_IdipGetHardLevelStatus, func() error {
		pid, ok := module.Player.GetPlayerByUsername(args.OpenId)
		if !ok {
			reply.ErrMsg = "query empty"
			return nil
		}
		mdb.GlobalExecute(func(globalDb *mdb.Database) {
			globalDb.AgentExecute(pid, func(db *mdb.Database) {
				hardLevelInfo := db.Lookup.PlayerHardLevel(pid)
				if hardLevelInfo == nil {
					reply.Status = 0
				} else {
					reply.Status = int32(hardLevelInfo.Lock)
				}
			})
		})
		return nil
	})
}
