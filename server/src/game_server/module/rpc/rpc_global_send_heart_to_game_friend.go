package rpc

import (
	"core/fail"
	"game_server/api/protocol/friend_api"
	"game_server/config"
	"game_server/mdb"
	. "game_server/rpc"
)

type Args_SendHeartToGameFriend struct {
	RPCArgTag
	Pid       int64
	FriendPid int64
}

type Reply_SendHeartToGameFriend struct {
	IsFriend bool
}

type Args_SendHeartToAllGameFriends struct {
	RPCArgTag
	Pid              int64
	ExcludeFriendPid []int64
}

type Reply_SendHeartToAllGameFriends struct {
	ValuedFriendsListPid []int64
}

//RPC服务端。在互动服务器检查好友关系并设置RPC返回结构体
func (this *RemoteServe) SendHeartToGameFriend(args *Args_SendHeartToGameFriend, reply *Reply_SendHeartToGameFriend) error {
	return Remote.Serve(mdb.RPC_Remote_SendHeartToGameFriend, args, mdb.TRANS_TAG_RPC_Serve_SendHeartToGameFriend, func() error {
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(args.Pid, func(db *mdb.Database) {
				db.Select.PlayerGlobalFriend(func(row *mdb.PlayerGlobalFriendRow) {
					if row.FriendPid() == args.FriendPid && row.FriendMode() == int8(friend_api.FRIEND_MODE_FRIEND) {
						reply.IsFriend = true
						row.Break()
					}
				})
			})

		})
		return nil
	})
}

//RPC客户端，暴露给游戏服内部使用的API
func RemoteSendHeartToGameFriend(pid int64, friendPid int64, cb func(bool)) {
	args := &Args_SendHeartToGameFriend{
		Pid:       pid,
		FriendPid: friendPid,
	}
	reply := new(Reply_SendHeartToGameFriend)
	Remote.Call(config.ServerCfg.GlobalServerId, mdb.RPC_Remote_SendHeartToGameFriend, args, reply, mdb.TRANS_TAG_RPC_Call_SendHeartToGameFriend, func(err error) {
		fail.When(err != nil, err)
		cb(reply.IsFriend)
	})
}

//RPC服务端。在互动服务器检查好友关系并设置RPC返回结构体
func (this *RemoteServe) SendHeartToAllGameFriends(args *Args_SendHeartToAllGameFriends, reply *Reply_SendHeartToAllGameFriends) error {
	return Remote.Serve(mdb.RPC_Remote_SendHeartToAllGameFriends, args, mdb.TRANS_TAG_RPC_Serve_SendHeartToAllGameFriends, func() error {
		mdb.GlobalExecute(func(globalDB *mdb.Database) {
			globalDB.AgentExecute(args.Pid, func(db *mdb.Database) {
				var (
					allFriendListPid         = []int64{}
					valuedFrindslistPid      = []int64{}
					isValue             bool = true
				)
				db.Select.PlayerGlobalFriend(func(row *mdb.PlayerGlobalFriendRow) {
					if row.FriendMode() == int8(friend_api.FRIEND_MODE_FRIEND) {
						allFriendListPid = append(allFriendListPid, row.FriendPid())
					}
				})
				for _, pid := range allFriendListPid {
					for _, invaluePid := range args.ExcludeFriendPid {
						if pid == invaluePid {
							isValue = false
						}
					}
					if isValue {
						valuedFrindslistPid = append(valuedFrindslistPid, pid)
					}
					isValue = true
				}
				reply.ValuedFriendsListPid = valuedFrindslistPid

			})
		})
		return nil
	})
}

//RPC客户端，暴露给游戏服内部使用的API
func RemoteSendHeartToAllGameFriends(playerPid int64, excludePid []int64, cb func([]int64)) {
	args := &Args_SendHeartToAllGameFriends{
		Pid:              playerPid,
		ExcludeFriendPid: excludePid,
	}
	reply := new(Reply_SendHeartToAllGameFriends)
	Remote.Call(config.ServerCfg.GlobalServerId, mdb.RPC_Remote_SendHeartToAllGameFriends, args, reply, mdb.TRANS_TAG_RPC_Call_SendHeartToAllGameFriends, func(err error) {
		fail.When(err != nil, err)
		cb(reply.ValuedFriendsListPid)
	})
}
