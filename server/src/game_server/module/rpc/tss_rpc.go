package rpc

import (
	"game_server/api/protocol/server_info_api"
	"game_server/config"
	"game_server/mdb"
	"game_server/module"
	. "game_server/rpc"
)

type Args_UserInfo struct {
	RPCArgTag
	OpenId     string
	PlayerId   int64
	Plat_Id    int
	Client_Ip  string
	Client_Ver int
	Role_Name  string
}

//rpc pkg_info
type Args_PkgInfo struct {
	RPCArgTag
	OpenId    string
	PlayerId  int64
	Plat_Id   int
	Anti_Data []byte
}

func (this *RemoteServe) TssSendData(args *Args_PkgInfo, reply *int) error {
	return Remote.Serve(mdb.RPC_Remote_TssSendData, args, mdb.TRANS_TAG_RPC_Serve_TssSendData, func() error {
		session, ok := module.Player.GetPlayerOnline(args.PlayerId)
		if ok {
			session.Send(&server_info_api.TssData_Out{Data: args.Anti_Data})
		}
		return nil
	})
}

func RemoteTssRecvData(openId string, pid int64, platId int, data []byte) {
	if config.ServerCfg.TssServerId == 0 {
		return
	}

	var reply int
	Remote.Call(config.ServerCfg.TssServerId, "RpcOper.RecvData", &Args_PkgInfo{OpenId: openId, PlayerId: pid, Plat_Id: platId, Anti_Data: data},
		&reply, mdb.TRANS_TAG_RPC_Call_TssRecvData, nil)
}

func RemoteTssUserLogin(openId, roleName, clientIP string, pid int64, platId, clientVer int) {
	if config.ServerCfg.TssServerId == 0 {
		return
	}

	var reply int
	Remote.Call(config.ServerCfg.TssServerId, "RpcOper.UserLogin", &Args_UserInfo{OpenId: openId, PlayerId: pid, Plat_Id: platId, Role_Name: roleName},
		&reply, mdb.TRANS_TAG_RPC_Call_TssUserLogin, nil)
}

func RemoteTssUserLogout(openId string, platId int) {
	if config.ServerCfg.TssServerId == 0 {
		return
	}

	var reply int
	Remote.Call(config.ServerCfg.TssServerId, "RpcOper.UserLogout", &Args_UserInfo{OpenId: openId, Plat_Id: platId},
		&reply, mdb.TRANS_TAG_RPC_Call_TssUserLogout, nil)
}
