package rpc

import (
	"core/fail"
	"core/time"
	"errors"
	"game_server/api/protocol/multi_level_api"
	"game_server/config"
	"game_server/dat/item_dat"
	"game_server/dat/player_dat"
	"game_server/mdb"
	"game_server/module"
	. "game_server/rpc"
	"game_server/tlog"
)

/*

 邀请多人关卡好友

*/
type Args_InviteFriendWithMultiLevel struct {
	RPCArgTag
	FriendId     int64
	PlayerNick   []byte
	GameServerId int64
	InviterId    int64
	RoomId       int64
	LevelId      int32
}

type Reply_InviteFriendWithMultiLevel struct {
	IsOnline bool
}

func (this *RemoteServe) InviteFriendWithMultiLevel(args *Args_InviteFriendWithMultiLevel, reply *Reply_InviteFriendWithMultiLevel) error {
	return Remote.Serve(mdb.RPC_Remote_InviteFriendWithMultiLevel, args, mdb.TRANS_TAG_RPC_Serve_InviteFriendWithMultiLevel, func() error {
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(args.FriendId, func(db *mdb.Database) {
				playerMultiLevel := db.Lookup.PlayerMultiLevelInfo(args.FriendId)
				if !time.IsInPointHour(player_dat.RESET_MULTI_LEVEL_TIMES_IN_HOUR, playerMultiLevel.BattleTime) {
					playerMultiLevel.DailyNum = 0
				}

				// 如果被邀请者在线，下发邀请通知
				if session, ok := module.Player.GetPlayerOnline(args.FriendId); ok {
					reply.IsOnline = true
					session.Send(&multi_level_api.NotifyRoomInvite_Out{
						RoomId:       args.RoomId,
						LevelId:      args.LevelId,
						Nickname:     args.PlayerNick,
						InviterId:    args.InviterId,
						DailyNum:     playerMultiLevel.DailyNum,
						GameServerId: args.GameServerId,
					})
				}
			})
		})
		return nil
	})
}

func RemoteInviteFriendWithMultiLevel(playerNick []byte, inviterId, friendId, roomId int64, levelId int32, callback func(*Reply_InviteFriendWithMultiLevel)) {
	reply := new(Reply_InviteFriendWithMultiLevel)
	args := &Args_InviteFriendWithMultiLevel{
		FriendId:     friendId,
		PlayerNick:   playerNick,
		RoomId:       roomId,
		LevelId:      levelId,
		InviterId:    inviterId,
		GameServerId: int64(config.ServerCfg.ServerId),
	}

	serverId, _ := module.GetServerIdWithPlayerId(friendId)
	Remote.Call(serverId, mdb.RPC_Remote_InviteFriendWithMultiLevel, args, reply, mdb.TRANS_TAG_RPC_Call_InviteFriendWithMultiLevel, func(err error) {
		fail.When(err != nil, err)
		callback(reply)
	})
}

type Args_StartMultiLevel struct {
	RPCArgTag
	PlayerId int64
}

type Reply_StartMultiLevel struct {
}

func (this *RemoteServe) StartMultiLevel(args *Args_StartMultiLevel, reply *Reply_StartMultiLevel) error {
	return Remote.Serve(mdb.RPC_Remote_StartMultiLevel, args, mdb.TRANS_TAG_RPC_Serve_StartMultiLevel, func() error {
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(args.PlayerId, func(db *mdb.Database) {
				module.Player.MustOpenFunc(db, player_dat.FUNC_MULTI_LEVEL)
				tlog.PlayerSystemModuleFlowLog(db, tlog.SMT_MULITI_MISSION)
			})
		})
		return nil
	})
}

func RemoteStartMultiLevel(playerId int64, callback func()) {
	reply := new(Reply_StartMultiLevel)
	args := &Args_StartMultiLevel{
		PlayerId: playerId,
	}

	serverId, _ := module.GetServerIdWithPlayerId(playerId)
	Remote.Call(serverId, mdb.RPC_Remote_StartMultiLevel, args, reply, mdb.TRANS_TAG_RPC_Call_StartMultiLevel, func(err error) {
		fail.When(err != nil, err)
		callback()
	})
}

/*

 多人关卡奖励

*/
type Args_AwardMultiLevel struct {
	RPCArgTag
	PlayerId    int64
	BuddyRoleId int8
	LevelId     int16
	Online      bool
}

type Reply_AwardMultiLevel struct {
}

func (this *RemoteServe) AwardMultiLevel(args *Args_AwardMultiLevel, reply *Reply_AwardMultiLevel) error {
	return Remote.Serve(mdb.RPC_Remote_AwardMultiLevel, args, mdb.TRANS_TAG_RPC_Serve_AwardMultiLevel, func() error {
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(args.PlayerId, func(db *mdb.Database) {
				module.MultiLevel.AwardWinner(db, args.BuddyRoleId, args.LevelId, args.Online)
			})
		})
		return nil
	})
}

func RemoteAwardMultiLevel(playerId int64, buddyRoleId int8, levelId int16, online bool) {
	reply := new(Reply_AwardMultiLevel)
	args := &Args_AwardMultiLevel{
		PlayerId:    playerId,
		BuddyRoleId: buddyRoleId,
		LevelId:     levelId,
		Online:      online,
	}

	serverId, _ := module.GetServerIdWithPlayerId(playerId)
	Remote.Call(serverId, mdb.RPC_Remote_AwardMultiLevel, args, reply, mdb.TRANS_TAG_RPC_Call_AwardMultiLevel, func(err error) {
		fail.When(err != nil, err)
	})
}

/*

 获取玩家多人关卡信息

*/
type Args_GetPlayerMultiLevelInfo struct {
	RPCArgTag
	PlayerId int64
}

type Reply_GetPlayerMultiLevelInfo struct {
	Info           *mdb.PlayerMultiLevelInfo
	FashionId      int16 //时装ID
	BattlePetInfo  map[int8]int32
	BattleItemInfo map[int16]int32
}

func (this *RemoteServe) GetPlayerMultiLevelInfo(args *Args_GetPlayerMultiLevelInfo, reply *Reply_GetPlayerMultiLevelInfo) error {
	return Remote.Serve(mdb.RPC_Remote_GetPlayerMultiLevelInfo, args, mdb.TRANS_TAG_RPC_Serve_GetPlayerMultiLevelInfo, func() error {
		var playerMultiLevel *mdb.PlayerMultiLevelInfo
		reply.BattlePetInfo = make(map[int8]int32)
		reply.BattleItemInfo = make(map[int16]int32)
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(args.PlayerId, func(db *mdb.Database) {
				if playerMultiLevel = db.Lookup.PlayerMultiLevelInfo(args.PlayerId); playerMultiLevel != nil {
					if !time.IsInPointHour(player_dat.RESET_MULTI_LEVEL_TIMES_IN_HOUR, playerMultiLevel.BattleTime) {
						playerMultiLevel.DailyNum = 0
						db.Update.PlayerMultiLevelInfo(playerMultiLevel)
					}
				}
				//获取时装信息
				playerFashionState := db.Lookup.PlayerFashionState(db.PlayerId())
				reply.FashionId = playerFashionState.DressedFashionId
				//获取装备的灵宠
				reply.BattlePetInfo = module.BattlePet.GetAvailableBattlePet(db)
				//获取战斗道具
				reply.BattleItemInfo = module.Item.CountItemNumByType(db, item_dat.TYPE_BATTLE_PROPS)
			})
		})

		if playerMultiLevel == nil {
			return errors.New("RemoteGetPlayerMultiLevelInfo is nil")
		}

		reply.Info = playerMultiLevel

		return nil
	})
}

func RemoteGetPlayerMultiLevelInfo(playerId int64, callback func(*mdb.PlayerMultiLevelInfo, map[int8]int32, map[int16]int32, int16)) {
	reply := new(Reply_GetPlayerMultiLevelInfo)
	args := &Args_GetPlayerMultiLevelInfo{
		PlayerId: playerId,
	}

	serverId, _ := module.GetServerIdWithPlayerId(playerId)
	Remote.Call(serverId, mdb.RPC_Remote_GetPlayerMultiLevelInfo, args, reply, mdb.TRANS_TAG_RPC_Call_GetPlayerMultiLevelInfo, func(err error) {
		fail.When(err != nil, err)
		callback(reply.Info, reply.BattlePetInfo, reply.BattleItemInfo, reply.FashionId)
	})
}

//拒绝多人关卡邀请
type Args_RefuseMultiLevelRoomInvite struct {
	RPCArgTag
	InviterId int64
	RoomId    int64
	NickName  []byte
}

type Reply_RefuseMultiLevelRoomInvite struct {
}

func (this *RemoteServe) RefuseMultiLevelRoomInvite(args *Args_RefuseMultiLevelRoomInvite, reply *Reply_RefuseMultiLevelRoomInvite) error {
	return Remote.Serve(mdb.RPC_Remote_RefuseMultiLevelRoomInvite, args, mdb.TRANS_TAG_RPC_Serve_RefuseMultiLevelRoomInvite, func() error {
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(args.InviterId, func(db *mdb.Database) {
				is_suitable := module.MultiLevel.IsMultiInvitePlayerSuitable(args.RoomId, args.InviterId)
				//玩家在线且房间没开始且玩家还在该房间内才下发通知
				if session, ok := module.Player.GetPlayerOnline(args.InviterId); is_suitable && ok {
					session.Send(&multi_level_api.NotifyRoomInviteRefuse_Out{
						Nickname: args.NickName,
					})
				}
			})
		})
		return nil
	})
}
func RemoteRefuseMultiLevelRoomInvite(roomId int64, inviterId int64, nickname []byte) {
	reply := new(Reply_RefuseMultiLevelRoomInvite)
	args := &Args_RefuseMultiLevelRoomInvite{
		InviterId: inviterId,
		RoomId:    roomId,
		NickName:  nickname,
	}
	serverId, _ := module.GetServerIdWithPlayerId(inviterId)
	Remote.Call(serverId, mdb.RPC_Remote_RefuseMultiLevelRoomInvite, args, reply, mdb.TRANS_TAG_RPC_Call_RefuseMultiLevelRoomInvite, func(err error) {
		fail.When(err != nil, err)
	})
}

//多人关卡设置阵形
type Args_MultiLevelChangeBuddy struct {
	RPCArgTag
	Pid         int64
	BuddyRoleId int8
}

type Reply_MultiLevelChangeBuddy struct {
}

func (this *RemoteServe) MultiLevelChangeBuddy(args *Args_MultiLevelChangeBuddy, reply *Reply_MultiLevelChangeBuddy) error {
	return Remote.Serve(mdb.RPC_Remote_MultiLevelChangeBuddy, args, mdb.TRANS_TAG_RPC_Serve_MultiLevelChangeBuddy, func() error {
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(args.Pid, func(db *mdb.Database) {
				playerMultiLevel := db.Lookup.PlayerMultiLevelInfo(args.Pid)
				buddyRole := module.Role.GetBuddyRoleInTeam(db, args.BuddyRoleId)
				fail.When(buddyRole == nil, "未获得角色")
				playerMultiLevel.BuddyRoleId = args.BuddyRoleId
				db.Update.PlayerMultiLevelInfo(playerMultiLevel)
			})
		})
		return nil
	})
}

func RemoteMultiLevelChangeBuddy(pid int64, buddyRoleId int8, callback func()) {
	reply := new(Reply_MultiLevelChangeBuddy)
	args := &Args_MultiLevelChangeBuddy{
		Pid:         pid,
		BuddyRoleId: buddyRoleId,
	}
	serverId, _ := module.GetServerIdWithPlayerId(pid)
	Remote.Call(serverId, mdb.RPC_Remote_MultiLevelChangeBuddy, args, reply, mdb.TRANS_TAG_RPC_Call_MultiLevelChangeBuddy, func(err error) {
		fail.When(err != nil, err)
		callback()
	})
}

//多人关卡设置阵形
type Args_MultiLevelChangeForm struct {
	RPCArgTag
	Pid      int64
	BuddyRow int8
}

type Reply_MultiLevelChangeForm struct {
}

func (this *RemoteServe) MultiLevelChangeForm(args *Args_MultiLevelChangeForm, reply *Reply_MultiLevelChangeForm) error {
	return Remote.Serve(mdb.RPC_Remote_MultiLevelChangeForm, args, mdb.TRANS_TAG_RPC_Serve_MultiLevelChangeForm, func() error {
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(args.Pid, func(db *mdb.Database) {
				playerMultiLevel := db.Lookup.PlayerMultiLevelInfo(args.Pid)
				playerMultiLevel.BuddyRow = args.BuddyRow
				db.Update.PlayerMultiLevelInfo(playerMultiLevel)
			})
		})
		return nil
	})
}

func RemoteMultiLevelChangeForm(pid int64, buddyRow int8, callback func()) {
	reply := new(Reply_MultiLevelChangeForm)
	args := &Args_MultiLevelChangeForm{
		Pid:      pid,
		BuddyRow: buddyRow,
	}
	serverId, _ := module.GetServerIdWithPlayerId(pid)
	Remote.Call(serverId, mdb.RPC_Remote_MultiLevelChangeForm, args, reply, mdb.TRANS_TAG_RPC_Call_MultiLevelChangeForm, func(err error) {
		fail.When(err != nil, err)
		callback()
	})
}
